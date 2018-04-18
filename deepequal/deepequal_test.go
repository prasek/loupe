package deepequal

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/prasek/go-testutil/mock"
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

func TestDeepEqual(t *testing.T) {
	var a, b string

	a = "aaabbbcccddd"
	b = "aaaccceeedddfff"

	m := mock.New()
	Assert(m, a, b, "a, b not equal: '%v', '%v'", a, b)
	res := m.Results().Out
	fmt.Printf(res)

	return

	a = readFile("../diff/test/fa.txt")
	b = readFile("../diff/test/fb.txt")
	Assert(t, a, b, "a, b not equal")

	a = readFile("../diff/test/log-a.txt")
	b = readFile("../diff/test/log-b.txt")
	Assert(t, a, b, "a, b not equal")

	Assert(t, tests[0], tests[1], "a, b tests not equal")

	for i := 0; i < 100; i++ {
		tests = append(tests, tests[1])
	}

	Assert(t, tests, tests[1], "a, b tests not equal")

}

func readFile(file string) string {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(bs)
}
