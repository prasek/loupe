package tools

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var _ TestingT = (*TestMock)(nil)

//TestingT defines what we require from *t.Testing
type TestingT interface {
	Fail()
	FailNow()
	Errorf(format string, args ...interface{})
}

//TestMock capture the output normally sent to *testing.T and os.Stdout
type TestMock struct {
	res  TestResults
	r    *os.File
	w    *os.File
	orig *os.File
	outc chan string
	err  bytes.Buffer
}

//TestResults contains the output normally sent to *testing.T and os.Stdout
type TestResults struct {
	Fail    bool
	FailNow bool
	Err     string
	Out     string
}

//Mock creates a new TestMock
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

//Fail marks the test as failed
func (t *TestMock) Fail() {
	t.res.Fail = true
}

//FailNow marks the test as failed and stops execution
func (t *TestMock) FailNow() {
	t.res.FailNow = true
}

//Errorf writes error info
func (t *TestMock) Errorf(format string, args ...interface{}) {
	fmt.Fprintf(&t.err, format, args...)
}

func (t *TestMock) run() {
	var buf bytes.Buffer
	io.Copy(&buf, t.r)
	t.outc <- buf.String()
}

//Results detaches from stdout and returns the TestResults that capture
//the info normally sent to testing.T and stdout
func (t *TestMock) Results() *TestResults {
	os.Stdout = t.orig
	t.w.Close()
	t.res.Out = <-t.outc
	t.r.Close()
	close(t.outc)
	t.res.Err = t.err.String()
	return &t.res
}
