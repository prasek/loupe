//go:generate go run gen.go

package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/prasek/go-testutil/diff"
)

func main() {
	var a, b, d string

	a = readFile("fa.txt")
	b = readFile("fb.txt")
	d = diff.New(a, b).String()
	writeFile("fd.txt", []byte(d))

	a = "aaabbbcccddd"
	b = "aaaccceeedddfff"
	d = diff.New(a, b).String()
	writeFile("word-d.txt", []byte(d))

	a = readFile("log-a.txt")
	b = readFile("log-b.txt")
	d = diff.New(a, b).String()
	writeFile("log-d.txt", []byte(d))

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
