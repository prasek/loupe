package testutil

import (
	"fmt"
	"github.com/prasek/go-testutil/internal"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	var a, b string
	var tests = internal.TestsA

	a = "aaabbbcccddd"
	b = "aaaccceeedddfff"

	m := Mock()
	AssertDeepEqual(m, a, b, "a, b not equal: '%v', '%v'", a, b)
	res := m.Results().Out
	fmt.Printf(res)

	return

	a = readFile(internal.File2A)
	b = readFile("../diff/test/fb.txt")
	AssertDeepEqual(t, a, b, "a, b not equal")

	a = readFile("../diff/test/log-a.txt")
	b = readFile("../diff/test/log-b.txt")
	AssertDeepEqual(t, a, b, "a, b not equal")

	AssertDeepEqual(t, tests[0], tests[1], "a, b tests not equal")

	for i := 0; i < 100; i++ {
		tests = append(tests, tests[1])
	}

	AssertDeepEqual(t, tests, tests[1], "a, b tests not equal")

}
