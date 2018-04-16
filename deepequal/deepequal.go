package deepequal

import (
	"bytes"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/fatih/color"
	gdiff "github.com/sergi/go-diff/diffmatchpatch"
)

const line = "================================================================="

var red = color.New(color.FgRed)
var green = color.New(color.FgGreen)

// assert does t.FailNow() if not equal
func Assert(t *testing.T, exp, act interface{}, format string, args ...interface{}) {
	deepEqual(t, exp, act, true, format, args...)
}

// require does t.Fail() if not equal
func Require(t *testing.T, exp, act interface{}, format string, args ...interface{}) {
	deepEqual(t, exp, act, false, format, args...)
}

// returns a color code diff string
// grab is the number of context lines before and after the diff to grab
func Diff(exp, act interface{}, grab int) string {
	//min grab
	if grab < 0 {
		grab = 0
	}

	//adjust pointer to value if needed for compare
	expt := reflect.TypeOf(exp)
	if expt.Kind() == reflect.Ptr {
		expv := reflect.ValueOf(exp)
		exp = expv.Elem().Interface()
	}

	actt := reflect.TypeOf(act)
	if actt.Kind() == reflect.Ptr {
		actv := reflect.ValueOf(act)
		act = actv.Elem().Interface()
	}

	// diff output
	d := gdiff.New()
	expText := fmt.Sprintf("%s", exp)
	actText := fmt.Sprintf("%s", act)
	diffs := d.DiffMain(expText, actText, false)

	var top bytes.Buffer
	var bottom bytes.Buffer
	var out bytes.Buffer

	for _, diff := range diffs {
		switch diff.Type {
		case gdiff.DiffDelete:
			if top.Len() > 0 {
				out.WriteString(top.String())
				top.Reset()
			}
			if bottom.Len() > 0 {
				out.WriteString(fmt.Sprintf("\n\n...\n\n"))
				out.WriteString(bottom.String())
				bottom.Reset()
			}

			out.WriteString(red.Sprintf("%s", diff.Text))

		case gdiff.DiffInsert:
			if top.Len() > 0 {
				out.WriteString(top.String())
				top.Reset()
			}
			if bottom.Len() > 0 {
				out.WriteString(fmt.Sprintf("\n\n...\n\n"))
				out.WriteString(bottom.String())
				bottom.Reset()
			}

			out.WriteString(green.Sprintf("%s", diff.Text))

		case gdiff.DiffEqual:
			lines := strings.Split(diff.Text, "\n")

			if len(lines) <= 2*grab {
				top.WriteString(diff.Text)
			} else {
				top.WriteString(strings.Join(lines[:grab], "\n"))
				bottom.WriteString(strings.Join(lines[len(lines)-grab:], "\n"))
			}

		default:
			out.WriteString(fmt.Sprintf("Unknown diff type: %v", diff.Type))
			break
		}
	}

	if top.Len() > 0 {
		lines := strings.Split(top.String(), "\n")

		if len(lines) <= grab {
			out.WriteString(top.String())
		} else {
			out.WriteString(strings.Join(lines[:grab], "\n"))
		}
	}

	return out.String()

}

// verifies the underlying instances are equal with diff output
func deepEqual(t *testing.T, exp, act interface{}, assert bool, format string, args ...interface{}) {

	expt := reflect.TypeOf(exp)
	if expt.Kind() == reflect.Ptr {
		expv := reflect.ValueOf(exp)
		exp = expv.Elem().Interface()
	}

	actt := reflect.TypeOf(act)
	if actt.Kind() == reflect.Ptr {
		actv := reflect.ValueOf(act)
		act = actv.Elem().Interface()
	}

	msg := fmt.Sprintf(format, args...)

	if !reflect.DeepEqual(exp, act) {

		// debug heder
		_, file, ln, _ := runtime.Caller(2)
		base := filepath.Base(file)

		fmt.Println()
		red.Println(line)
		red.Printf("%s:%d: Not Equal (%T/%T)\n%s\n", base, ln, exp, act, msg)
		red.Println(line)

		//out

		if assert {
			t.FailNow()
		} else {
			t.Fail()
		}
	}
}
