package deepequal

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/fatih/color"
	"github.com/prasek/go-testutil/diff"
)

var red = color.New(color.FgRed)
var green = color.New(color.FgGreen)

//what we require from *t.Testing
type TestingT interface {
	Fail()
	FailNow()
	Errorf(format string, args ...interface{})
}

// verifies underlying values are deep equal and prints a diff if not
func Assert(t TestingT, exp, act interface{}, format string, args ...interface{}) bool {
	ok := testDeepEqual(t, exp, act, format, args...)
	if !ok {
		t.Fail()
	}
	return ok
}

// verifies underlying values are deep equal and prints a diff if not
func Require(t TestingT, exp, act interface{}, format string, args ...interface{}) bool {
	ok := testDeepEqual(t, exp, act, format, args...)
	if !ok {
		t.FailNow()
	}
	return ok
}

// deep compare of underlying values
func DeepEqual(a, b interface{}) bool {
	return reflect.DeepEqual(diff.Value(a), diff.Value(b))
}

// verifies the underlying instances are equal with diff output
func testDeepEqual(t TestingT, exp, act interface{}, format string, args ...interface{}) bool {
	if DeepEqual(exp, act) {
		return true
	}

	msg := fmt.Sprintf(format, args...)

	// debug heder
	_, file, ln, _ := runtime.Caller(2)
	base := filepath.Base(file)

	const line = "================================================================="

	var buf bytes.Buffer

	red.Fprintln(&buf, line)
	red.Fprintf(&buf, "%s:%d: Not Equal (%T/%T)\n%s\n", base, ln, exp, act, msg)
	red.Fprintln(&buf, line)

	diff.New(exp, act).WriteTo(&buf)
	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf)

	buf.WriteTo(os.Stdout)
	t.Errorf("%s:%d: Not Equal (%T/%T)\n%s\n", base, ln, exp, act, msg)

	return false
}
