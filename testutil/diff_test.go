package testutil

import (
	"bytes"
	"testing"

	"github.com/prasek/go-testutil/internal"
	"github.com/stretchr/testify/assert"
)

func readFile(file string) string {
	file = path.Join("../internal", file)
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(bs)
}

func TestDiff(t *testing.T) {
	//test 1
	a = internal.Line1A
	b = internal.Line1B
	d = readFile(internal.Line1Diff)
	c = Diff(a, b).String()

	if c != d {
		Diff(c, d).Print()
	}
	assert.Equal(t, d, c, internal.Line1Diff)

	//test 2
	var a, b, c, d string
	a = readFile(internal.File2A)
	b = readFile(internal.File2B)
	d = readFile(internal.File2Diff)

	c = Diff(a, b).String()

	if c != d {
		Diff(c, d).Print()
	}
	assert.Equal(t, d, c, internal.File2Diff)

	//test 3
	a = readFile("test/log-a.txt")
	b = readFile("test/log-b.txt")
	d = readFile("test/log-d.txt")
	var buf bytes.Buffer
	_, _ = Diff(a, b).WriteTo(&buf)
	c = buf.String()

	if c != d {
		Diff(c, d).Print()
	}
	assert.Equal(t, d, c, "log-d.txt mismatch")
}
