package tools

import (
	"fmt"
	"testing"

	"github.com/prasek/loupe/internal"
	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	var a, b interface{}
	var c, d string
	var tests = internal.Tests

	for _, test := range tests {
		switch test.InputType {
		case internal.ValueInput:
			a = test.InputA
			b = test.InputB
		case internal.FileInput:
			a = readFile(test.InputA.(string))
			b = readFile(test.InputB.(string))
		default:
			panic(fmt.Sprintf("unkonwn file type %v", test.InputType))
		}

		c = readFile(test.DiffFile)
		d = Diff(a, b).String()

		c = regExColor.ReplaceAllString(c, "")
		d = regExColor.ReplaceAllString(d, "")

		if !assert.Equal(t, c, d, "a, b not equal: see diff") {
			Diff(c, d).Print()
		}
	}
}
