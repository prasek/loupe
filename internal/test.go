package internal

const (
	ValueInput InputType = "value"
	FileInput  InputType = "file"
)

type InputType string

type Test struct {
	Result    bool
	InputType InputType
	InputA    interface{}
	InputB    interface{}
	DiffFile  string
	DeEqFile  string
}

var Tests = []Test{
	Test{
		Result:    false,
		InputType: ValueInput,
		InputA:    "aaabbbcccddd",
		InputB:    "aaabbbeeecccdddfffggg",
		DiffFile:  "result/diff/1.txt",
		DeEqFile:  "result/de/1.txt",
	},
	Test{
		Result:    false,
		InputType: FileInput,
		InputA:    "input/2a.txt",
		InputB:    "input/2b.txt",
		DiffFile:  "result/diff/2.txt",
		DeEqFile:  "result/de/2.txt",
	},
	Test{
		Result:    false,
		InputType: FileInput,
		InputA:    "input/3a.txt",
		InputB:    "input/3b.txt",
		DiffFile:  "result/diff/3.txt",
		DeEqFile:  "result/de/3.txt",
	},
	Test{
		Result:    false,
		InputType: ValueInput,
		InputA:    StructsA[0],
		InputB:    StructsA[1],
		DiffFile:  "result/diff/4.txt",
		DeEqFile:  "result/de/4.txt",
	},
	Test{
		Result:    false,
		InputType: ValueInput,
		InputA:    StructsA,
		InputB:    StructsA[1],
		DiffFile:  "result/diff/5.txt",
		DeEqFile:  "result/de/5.txt",
	},
}

type NestedTestStruct struct {
	a string
	b string
}

type TestStruct struct {
	a string
	b int
	c bool
	d string
	e NestedTestStruct
}

var StructsA = []TestStruct{
	TestStruct{a: "foo", b: 5, c: false, d: "bar", e: NestedTestStruct{a: "zap", b: "pow"}},
	TestStruct{a: "bar", b: 5, c: true, d: "bar", e: NestedTestStruct{a: "zap", b: "pow"}},
}
