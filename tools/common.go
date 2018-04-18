package tools

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed)
var green = color.New(color.FgGreen)

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
