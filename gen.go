//go:generate go run gen.go

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/prasek/go-testutil/internal"
	"github.com/prasek/go-testutil/testutil"
)

func main() {
	var m *testutil.TestMock
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
		d = testutil.Diff(a, b).String()
		writeFile(t.DiffFile, []byte(d))

		m = testutil.Mock()
		testutil.AssertDeepEqual(m, a, b, "a, b not equal: '%v', '%v'", a, b)
		writeFile(t.DeEqFile, []byte(m.Results().Out))
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
	err := ioutil.WriteFile(file, data, os.FileMode(0770))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
