package diff

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/fatih/color"
	dmp "github.com/sergi/go-diff/diffmatchpatch"
)

var red = color.New(color.FgRed)
var green = color.New(color.FgGreen)

const (
	nle = "%0A"
	nl  = "\n"
)

var unescaper = strings.NewReplacer(
	"%21", "!", "%7E", "~", "%27", "'",
	"%28", "(", "%29", ")", "%3B", ";",
	"%2F", "/", "%3F", "?", "%3A", ":",
	"%40", "@", "%26", "&", "%3D", "=",
	"%2B", "+", "%24", "$", "%2C", ",",
	"%23", "#", "%2A", "*", "%0A", "",
	"%5B", "[", "%5D", "]", "%09", "^I",
	"%0D", "^M")

type Differ interface {
	Print()
	String() string
	WriteTo(w io.Writer) (int64, error)
}

func New(a, b interface{}) Differ {
	textA := getText(a)
	textB := getText(b)

	hasLines := false
	if strings.Contains(textA, nl) {
		hasLines = true
	} else if strings.Contains(textB, nl) {
		hasLines = true
	}

	var diff Differ

	switch hasLines {
	case false:
		diff = &wordDiff{
			a: textA,
			b: textB,
		}
	default:
		diff = &unifiedDiff{
			a: textA,
			b: textB,
		}
	}

	return diff

}

func getText(v interface{}) string {
	return fmt.Sprintf("%s", Value(v))
}

func Value(v interface{}) interface{} {
	vt := reflect.TypeOf(v)
	if vt.Kind() == reflect.Ptr {
		vv := reflect.ValueOf(v)
		v = vv.Elem().Interface()
	}
	return v
}

type wordDiff struct {
	a string
	b string
}

func (d *wordDiff) Print() {
	d.diff(os.Stdout)
	fmt.Println()
}

func (d *wordDiff) String() string {
	var buf bytes.Buffer
	d.diff(&buf)
	return buf.String()
}

func (d *wordDiff) WriteTo(w io.Writer) (int64, error) {
	var buf bytes.Buffer
	d.diff(&buf)
	return buf.WriteTo(w)
}

func (d *wordDiff) diff(w io.Writer) {
	gd := dmp.New()
	diffs := gd.DiffMain(d.a, d.b, false)

	for _, diff := range diffs {
		switch diff.Type {
		case dmp.DiffDelete:
			red.Fprint(w, diff.Text)

		case dmp.DiffInsert:
			green.Fprintf(w, diff.Text)

		case dmp.DiffEqual:
			fmt.Fprint(w, diff.Text)

		default:
			fmt.Fprintf(w, "ERROR: Unknown diff type: %v", diff.Type)
			return
		}
	}
}

type unifiedDiff struct {
	a string
	b string
}

func (d *unifiedDiff) Print() {
	d.diff(os.Stdout)
	fmt.Println()
}

func (d *unifiedDiff) String() string {
	var buf bytes.Buffer
	d.diff(&buf)
	return buf.String()
}

func (d *unifiedDiff) WriteTo(w io.Writer) (int64, error) {
	var buf bytes.Buffer
	d.diff(&buf)
	return buf.WriteTo(w)
}

func (d *unifiedDiff) diff(w io.Writer) {
	gd := dmp.New()
	a, b, lineArray := gd.DiffLinesToRunes(d.a, d.b)
	diffs := gd.DiffMainRunes(a, b, false)
	diffs = gd.DiffCharsToLines(diffs, lineArray)
	diffs = gd.DiffCleanupSemantic(diffs)

	patches := gd.PatchMake(diffs)
	for _, patch := range patches {
		lines := strings.Split(patch.String(), nl)
		for _, line := range lines {
			if len(line) == 0 {
				fmt.Fprintln(w)
				continue
			}
			prefix := line[0]
			switch prefix {
			case '+', '-':
				line = strings.TrimSuffix(line, nle)
				difflines := strings.Split(string(line[1:]), nle)
				for _, diffline := range difflines {
					diffline = unescaper.Replace(diffline)
					switch prefix {
					case '-':
						red.Fprintf(w, "-%s\n", diffline)
					case '+':
						green.Fprintf(w, "+%s\n", diffline)
					default:
						fmt.Fprintf(w, "ERROR: unknown prefix %v", prefix)
						return
					}
				}

			default:
				line = unescaper.Replace(line)
				fmt.Fprintln(w, line)
			}
		}
	}
}
