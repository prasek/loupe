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

		// diff output
		d := gdiff.New()
		expText := fmt.Sprintf("%s", exp)
		actText := fmt.Sprintf("%s", act)
		diffs := d.DiffMain(expText, actText, false)

		var top bytes.Buffer
		var bottom bytes.Buffer

		for _, diff := range diffs {
			switch diff.Type {
			case gdiff.DiffDelete:
				if top.Len() > 0 {
					fmt.Printf(top.String())
					top.Reset()
				}
				if bottom.Len() > 0 {
					fmt.Printf("\n\n...\n\n")
					fmt.Printf(bottom.String())
					bottom.Reset()
				}

				red.Printf("%s", diff.Text)

			case gdiff.DiffInsert:
				if top.Len() > 0 {
					fmt.Printf(top.String())
					top.Reset()
				}
				if bottom.Len() > 0 {
					fmt.Printf("\n\n...\n\n")
					fmt.Printf(bottom.String())
					bottom.Reset()
				}

				green.Printf("%s", diff.Text)

			case gdiff.DiffEqual:
				grab := 10
				lines := strings.Split(diff.Text, "\n")

				if len(lines) <= 2*grab {
					top.WriteString(diff.Text)
				} else {
					top.WriteString(strings.Join(lines[:grab], "\n"))
					bottom.WriteString(strings.Join(lines[len(lines)-grab:], "\n"))
				}

			default:
				fmt.Printf("Unknown diff type: %v", diff.Type)
				break
			}
		}

		if top.Len() > 0 {
			fmt.Println(top.String())
		}

		if assert {
			t.FailNow()
		} else {
			t.Fail()
		}
	}
}
