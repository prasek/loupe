//go:generate go run gen.go

package main

import (
	"io/ioutil"
	"log"
	"os"

	de "github.com/prasek/go-testutil/deepequal"
	"github.com/prasek/go-testutil/mock"
)

var tests = []Test{
	Test{a: "foo", b: 5, c: false, d: "bar", e: Nested{a: "zap", b: "pow"}},
	Test{a: "bar", b: 5, c: true, d: "bar", e: Nested{a: "zap", b: "pow"}},
}

type Nested struct {
	a string
	b string
}

type Test struct {
	a string
	b int
	c bool
	d string
	e Nested
}

func main() {
	var a, b string
	var m *mock.TestCaptureMock

	//case 1
	a = "aaabbbcccddd"
	b = "aaaccceeedddfff"

	m = mock.New()
	de.Assert(m, a, b, "a, b not equal: '%v', '%v'", a, b)
	writeFile("1.txt", []byte(m.Results().Out))

	//case 2
	a = readFile("../../diff/test/fa.txt")
	b = readFile("../../diff/test/fb.txt")

	m = mock.New()
	de.Assert(m, a, b, "a, b not equal")
	writeFile("2.txt", []byte(m.Results().Out))

	//case 3
	a = readFile("../../diff/test/log-a.txt")
	b = readFile("../../diff/test/log-b.txt")

	m = mock.New()
	de.Assert(m, a, b, "a, b not equal")
	writeFile("3.txt", []byte(m.Results().Out))

	//case 4
	m = mock.New()
	de.Assert(m, tests[0], tests[1], "a, b tests not equal")
	writeFile("4.txt", []byte(m.Results().Out))

	//case 5
	for i := 0; i < 100; i++ {
		tests = append(tests, tests[1])
	}

	m = mock.New()
	de.Assert(m, tests, tests[1], "a, b tests not equal")
	writeFile("5.txt", []byte(m.Results().Out))
}

func readFile(file string) string {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(bs)
}

func writeFile(file string, data []byte) error {
	err := ioutil.WriteFile(file, data, os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
