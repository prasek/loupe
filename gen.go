//go:generate go run gen.go

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"

	"github.com/prasek/loupe/internal"
	"github.com/prasek/loupe/tools"
)

var regExFilename = regexp.MustCompile(`gen\.go\:[0-9]*\:`)

func main() {
	var res *tools.TestResults
	var m *tools.TestMock
	var a, b interface{}
	var d string
	var tests = internal.Tests

	for _, t := range tests {
		switch t.InputType {
		case internal.ValueInput:
			a = t.InputA
			b = t.InputB
		case internal.FileInput:
			a = readFile(t.InputA.(string))
			b = readFile(t.InputB.(string))
		default:
			panic(fmt.Sprintf("unkonwn file type %v", t.InputType))
		}
		d = tools.Diff(a, b).String()
		writeFile(t.DiffFile, []byte(d))

		m = tools.Mock()
		tools.AssertDeepEqual(m, a, b, "a, b not equal: see diff")
		res = m.Results()
		res.Out = regExFilename.ReplaceAllString(res.Out, ":")
		writeFile(t.DeEqFile, []byte(res.Out))
	}
}

func readFile(file string) string {
	file = path.Join("internal", file)
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(bs)
}

func writeFile(file string, data []byte) error {
	file = path.Join("internal", file)

	err := os.MkdirAll(path.Dir(file), os.FileMode(0777))
	if err != nil {
		return fmt.Errorf("Make dir failed: %v", err)
	}

	err = ioutil.WriteFile(file, data, os.FileMode(0666))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
