package tools

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var _ TestingT = (*TestMock)(nil)

//what we require from *t.Testing
type TestingT interface {
	Fail()
	FailNow()
	Errorf(format string, args ...interface{})
}

type TestResults struct {
	Fail    bool
	FailNow bool
	Err     string
	Out     string
}

//test mock that captures stdout
type TestMock struct {
	res  TestResults
	r    *os.File
	w    *os.File
	orig *os.File
	outc chan string
	err  bytes.Buffer
}

//must call results to close pipe and return stdout
func Mock() *TestMock {
	t := &TestMock{}
	var err error
	t.outc = make(chan string, 1)
	t.orig = os.Stdout
	t.r, t.w, err = os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = t.w
	go t.run()
	return t
}

func (t *TestMock) Fail() {
	t.res.Fail = true
}

func (t *TestMock) FailNow() {
	t.res.FailNow = true
}

func (t *TestMock) Errorf(format string, args ...interface{}) {
	fmt.Fprintf(&t.err, format, args...)
}

func (t *TestMock) run() {
	var buf bytes.Buffer
	io.Copy(&buf, t.r)
	t.outc <- buf.String()
}

func (t *TestMock) Results() *TestResults {
	os.Stdout = t.orig
	t.w.Close()
	t.res.Out = <-t.outc
	t.r.Close()
	close(t.outc)
	t.res.Err = t.err.String()
	return &t.res
}
