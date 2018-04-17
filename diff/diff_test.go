package diff

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	var a, b, c, d string
	a = readFile("test/fa.txt")
	b = readFile("test/fb.txt")
	d = readFile("test/fd.txt")
	c = New(a, b).String()

	if c != d {
		New(c, d).Print()
	}
	assert.Equal(t, d, c, "fd.txt mismatch")

	a = "aaabbbcccddd"
	b = "aaaccceeedddfff"
	d = readFile("test/word-d.txt")
	c = New(a, b).String()

	if c != d {
		New(c, d).Print()
	}
	assert.Equal(t, d, c, "word-d.txt mismatch")

	a = readFile("test/log-a.txt")
	b = readFile("test/log-b.txt")
	d = readFile("test/log-d.txt")
	var buf bytes.Buffer
	_, _ = New(a, b).WriteTo(&buf)
	c = buf.String()

	if c != d {
		New(c, d).Print()
	}
	assert.Equal(t, d, c, "log-d.txt mismatch")
}

func readFile(file string) string {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(bs)
}
