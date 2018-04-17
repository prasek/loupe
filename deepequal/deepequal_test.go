package deepequal

import (
	"io/ioutil"
	"log"
	"testing"
)

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

var tests = []Test{
	Test{a: "foo", b: 5, c: false, d: "bar", e: Nested{a: "zap", b: "pow"}},
	Test{a: "bar", b: 5, c: true, d: "bar", e: Nested{a: "zap", b: "pow"}},
}

func TestDiff(t *testing.T) {
	var a, b string

	a = "aaabbbcccddd"
	b = "aaaccceeedddfff"

	Assert(t, a, b, "a, b not equal: '%v', '%v'", a, b)

	a = readFile("../diff/test/fa.txt")
	b = readFile("../diff/test/fb.txt")
	Assert(t, a, b, "a, b not equal")

	Assert(t, tests[0], tests[1], "a, b tests not equal")

}

func readFile(file string) string {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(bs)
}
