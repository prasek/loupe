package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"testing"

	"github.com/prasek/loupe/internal"
	"github.com/stretchr/testify/assert"
)

var regExFilename = regexp.MustCompile(`deepequal_test\.go\:[0-9]*\:`)

func TestDeepEqual(t *testing.T) {
	var m *TestMock
	var res *TestResults
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

		// AssertDeepEqual
		// we're testing a test helper, so mock it and get the results
		c = readFile(test.AstDeEqFile)
		m = Mock()
		AssertDeepEqual(m, a, b, "a, b not equal: see diff")
		res = m.Results()
		d = regExFilename.ReplaceAllString(res.Out, ":")

		c = regExColor.ReplaceAllString(c, "")
		d = regExColor.ReplaceAllString(d, "")

		if !assert.Equal(t, c, d, "a, b not equal: see diff") {
			Diff(c, d).Print()
		}

		assert.Equal(t, !test.ExpResult, res.Fail, "res.Fail set incorrectly")
		assert.Equal(t, false, res.FailNow, "res.FailNow set incorrectly")

		// RequireDeepEqual
		// we're testing a test helper, so mock it and get the results
		c = readFile(test.ReqDeEqFile)
		m = Mock()
		RequireDeepEqual(m, a, b, "a, b not equal: see diff")
		res = m.Results()
		d = regExFilename.ReplaceAllString(res.Out, ":")

		c = regExColor.ReplaceAllString(c, "")
		d = regExColor.ReplaceAllString(d, "")

		if !assert.Equal(t, c, d, "a, b not equal: see diff") {
			Diff(c, d).Print()
		}
		assert.Equal(t, false, res.Fail, "res.Fail set incorrectly")
		assert.Equal(t, !test.ExpResult, res.FailNow, "res.FailNow set incorrectly")
	}
}

func readFile(file string) string {
	file = path.Join("../internal", file)
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(bs)
}
