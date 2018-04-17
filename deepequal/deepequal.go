package deepequal

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

const line = "================================================================="

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

	if !reflect.DeepEqual(exp, act) {

		msg := fmt.Sprintf(format, args...)

		// debug heder
		_, file, ln, _ := runtime.Caller(2)
		base := filepath.Base(file)

		fmt.Println()
		red.Println(line)
		red.Printf("%s:%d: Not Equal (%T/%T)\n%s\n", base, ln, exp, act, msg)
		red.Println(line)

		Diff(exp, act).Print()

		if assert {
			t.FailNow()
		} else {
			t.Fail()
		}
	}
}
