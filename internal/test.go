package internal

const (
	//ValueInput allows interface{} value diff tests
	ValueInput InputType = "value"

	//FileInput pulls input from files
	FileInput InputType = "file"
)

//InputType is the type of test input
type InputType string

//Test is for table driven tests
type Test struct {
	ExpResult   bool
	InputType   InputType
	InputA      interface{}
	InputB      interface{}
	DiffFile    string
	AstDeEqFile string
	ReqDeEqFile string
}

//Tests is the  main set of tests
var Tests = []Test{
	{
		ExpResult:   false,
		InputType:   ValueInput,
		InputA:      "aaabbbcccddd",
		InputB:      "aaabbbeeecccdddfffggg",
		DiffFile:    "result/diff/1.txt",
		AstDeEqFile: "result/de/1.txt",
		ReqDeEqFile: "result/rde/1.txt",
	},
	{
		ExpResult:   false,
		InputType:   FileInput,
		InputA:      "input/2a.txt",
		InputB:      "input/2b.txt",
		DiffFile:    "result/diff/2.txt",
		AstDeEqFile: "result/de/2.txt",
		ReqDeEqFile: "result/rde/2.txt",
	},
	{
		ExpResult:   false,
		InputType:   FileInput,
		InputA:      "input/3a.txt",
		InputB:      "input/3b.txt",
		DiffFile:    "result/diff/3.txt",
		AstDeEqFile: "result/de/3.txt",
		ReqDeEqFile: "result/rde/3.txt",
	},
	{
		ExpResult:   false,
		InputType:   ValueInput,
		InputA:      &StructsA[0],
		InputB:      StructsA[1],
		DiffFile:    "result/diff/4.txt",
		AstDeEqFile: "result/de/4.txt",
		ReqDeEqFile: "result/rde/4.txt",
	},
	{
		ExpResult:   false,
		InputType:   ValueInput,
		InputA:      StructsA,
		InputB:      StructsA[1],
		DiffFile:    "result/diff/5.txt",
		AstDeEqFile: "result/de/5.txt",
		ReqDeEqFile: "result/rde/5.txt",
	},
	{
		ExpResult:   true,
		InputType:   ValueInput,
		InputA:      "aaaaaaaa",
		InputB:      "aaaaaaaa",
		DiffFile:    "result/diff/6.txt",
		AstDeEqFile: "result/de/6.txt",
		ReqDeEqFile: "result/rde/6.txt",
	},
}

//NestedTestStruct is a dummy struct for testing
type NestedTestStruct struct {
	a string
	b string
}

//TestStruct is a dummy struct for testing
type TestStruct struct {
	a string
	b int
	c bool
	d string
	e NestedTestStruct
}

//StructsA is an array of dummy structs for testing
var StructsA = []TestStruct{
	{a: "foo", b: 5, c: false, d: "bar", e: NestedTestStruct{a: "zap", b: "pow"}},
	{a: "bar", b: 5, c: true, d: "bar", e: NestedTestStruct{a: "zap", b: "pow"}},
}
