package main

import "testing"

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
		{
			name:    "invalid input discards partial declarations",
			content: "func Valid() {}\nfunc Broken(",
			want:    `[]`,
		},
		{
			name:    "package detection remains a literal substring check",
			content: "// package name omitted\nfunc Visible() {}",
			want:    `[]`,
		},
		{
			name:    "marker remains a literal substring check",
			content: `var Marker = "// --- Auto-generated FFI wrappers ---"`,
			want:    `[]`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := parseFFI(test.content); got != test.want {
				t.Errorf("parseFFI() = %s\nwant %s", got, test.want)
			}
		})
	}
}
