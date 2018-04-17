package deepequal

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestDiff(t *testing.T) {
	f1a := readFile("test/fa.txt")
	f1b := readFile("test/fb.txt")
	diff := Diff(f1a, f1b).String()
	fmt.Println(diff)
	fmt.Print("\n\n")

	a := "aaabbbcccddd"
	b := "aaaccceeedddfff"

	Diff(a, b).Print()
	fmt.Print("\n\n")

	l1a := readFile("test/log-a.txt")
	l1b := readFile("test/log-b.txt")
	Diff(l1a, l1b).Print()
	fmt.Print("\n\n")
}

func readFile(file string) string {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(bs)
}
