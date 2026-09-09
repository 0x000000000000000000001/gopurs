package main

import (
	"go/scanner"
	"testing"
)

func TestParseFFIContract(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    string
	}{
		{
			name: "functions and first return value",
			content: `package ffi
func Add(a, b int, label string) (int, error) { return 0, nil }
func hidden() {}
func _Private() {}
func Étranger() {}
type Receiver struct{}
func (Receiver) Method() {}`,
			want: `[{"name":"Add","isVar":false,"typeParams":[],"args":[{"type":"Named","name":"int"},{"type":"Named","name":"int"},{"type":"Named","name":"string"}],"ret":{"type":"Named","name":"int"}},{"name":"_Private","isVar":false,"typeParams":[],"args":[],"ret":null},{"name":"Method","isVar":false,"typeParams":[],"args":[],"ret":null}]`,
		},
		{
			name: "variables keep declaration names but omit their types",
			content: `var Exported, hidden, _Internal int
var Callback func(string) bool
var Inferred = 1
const Constant = 2
type Alias = int`,
			want: `[{"name":"Exported","isVar":true,"typeParams":[],"args":[],"ret":null},{"name":"_Internal","isVar":true,"typeParams":[],"args":[],"ret":null},{"name":"Callback","isVar":true,"typeParams":[],"args":[],"ret":null},{"name":"Inferred","isVar":true,"typeParams":[],"args":[],"ret":null}]`,
		},
		{
			name:    "generic names and structured types",
			content: `func Convert[A, B any, K comparable](values []A, lookup map[K]B, fixed [2]int, ctx context.Context, arbitrary interface{}) []B { return nil }`,
			want:    `[{"name":"Convert","isVar":false,"typeParams":["A","B","K"],"args":[{"type":"Array","elem":{"type":"Named","name":"A"}},{"type":"Map","key":{"type":"Named","name":"K"},"val":{"type":"Named","name":"B"}},{"type":"Array","elem":{"type":"Named","name":"int"}},{"type":"Named","name":"context.Context"},{"type":"Named","name":"any"}],"ret":{"type":"Array","elem":{"type":"Named","name":"B"}}}]`,
		},
		{
			name:    "nested callbacks and omitted empty TypeNode fields",
			content: `func Call(cb func(left, right int, next func() string) (func(func(bool)), error), empty func()) {}`,
			want:    `[{"name":"Call","isVar":false,"typeParams":[],"args":[{"type":"Func","args":[{"type":"Named","name":"int"},{"type":"Named","name":"int"},{"type":"Func","ret":{"type":"Named","name":"string"}}],"ret":{"type":"Func","args":[{"type":"Func","args":[{"type":"Named","name":"bool"}]}]}},{"type":"Func"}],"ret":null}]`,
		},
		{
			name:    "unsupported expressions retain printer output",
			content: `func Unknown(pointer *Value, generic Box[int], values ...string) chan int { return nil }`,
			want:    `[{"name":"Unknown","isVar":false,"typeParams":[],"args":[{"type":"Unknown","name":"*Value"},{"type":"Unknown","name":"Box[int]"},{"type":"Unknown","name":"...string"}],"ret":{"type":"Unknown","name":"chan int"}}]`,
		},
		{
			name: "generated marker truncates before parsing",
			content: `func Before() {}
// --- Auto-generated FFI wrappers ---
this is invalid Go
func After() {}`,
			want: `[{"name":"Before","isVar":false,"typeParams":[],"args":[],"ret":null}]`,
		},
		{
			name:    "empty input",
			content: "",
			want:    `[]`,
		},
		{
			name:    "no retained declarations",
			content: "package ffi\nfunc hidden() {}\nvar hiddenVariable int",
			want:    `[]`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseFFI(test.content)
			if err != nil {
				t.Fatalf("parseFFI() failed: %v", err)
			}
			if got != test.want {
				t.Errorf("parseFFI() = %s\nwant %s", got, test.want)
			}
		})
	}
}

func TestParseFFIRejectsInvalidInput(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		line     int
		column   int
		offset   int
		filename string
	}{
		{
			name:    "invalid input discards partial declarations",
			content: "func Valid() {}\nfunc Broken(",
			line:    2,
			column:  13,
			offset:  28,
		},
		{
			name:    "explicit package keeps original positions",
			content: "package ffi\nfunc Valid() {}\nfunc Broken(",
			line:    3,
			column:  13,
			offset:  40,
		},
		{
			name:     "line directives retain their own positions",
			content:  "//line generated.go:40\nfunc Broken(",
			line:     40,
			column:   0,
			offset:   35,
			filename: "generated.go",
		},
		{
			name:    "package detection remains a literal substring check",
			content: "// package name omitted\nfunc Visible() {}",
			line:    2,
			column:  1,
			offset:  24,
		},
		{
			name:    "marker remains a literal substring check",
			content: `var Marker = "// --- Auto-generated FFI wrappers ---"`,
			line:    1,
			column:  14,
			offset:  13,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseFFI(test.content)
			if err == nil {
				t.Fatal("parseFFI() accepted invalid Go")
			}
			if got != "" {
				t.Errorf("parseFFI() returned declarations after a syntax error: %s", got)
			}
			parseErrors, ok := err.(scanner.ErrorList)
			if !ok || len(parseErrors) == 0 {
				t.Fatalf("parseFFI() did not return syntax errors: %v", err)
			}
			pos := parseErrors[0].Pos
			if pos.Filename != test.filename {
				t.Errorf("first error filename = %q, want %q", pos.Filename, test.filename)
			}
			if pos.Line != test.line || pos.Column != test.column || pos.Offset != test.offset {
				t.Errorf("first error position = %d:%d (offset %d), want %d:%d (offset %d)", pos.Line, pos.Column, pos.Offset, test.line, test.column, test.offset)
			}
		})
	}
}
