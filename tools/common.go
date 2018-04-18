package tools

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/fatih/color"
)

var regExColor = regexp.MustCompile(`\x1b\[[0-9]*m`)

var red = color.New(color.FgRed)
var green = color.New(color.FgGreen)

//Value returns the value of v
func Value(v interface{}) interface{} {
	vt := reflect.TypeOf(v)
	if vt.Kind() == reflect.Ptr {
		vv := reflect.ValueOf(v)
		v = vv.Elem().Interface()
	}
	return v
}

func getText(v interface{}) string {
	return fmt.Sprintf("%s", Value(v))
}
