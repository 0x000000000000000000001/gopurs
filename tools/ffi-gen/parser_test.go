package main

import (
	"encoding/json"
	"go/scanner"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func TestParseFFIRenamesPackageReferences(t *testing.T) {
	const source = `package main
import "strings"
// Words, Callback and Count stay unchanged in comments.
var label = "Words Callback Count"
var Count = 0
var Callback = Words
var Key = "key"
var entries = map[string]int{Key: 1}
var object = struct{ Words func(string) []string }{Words: strings.Fields}
type reader struct{}
func (reader) Words(s string) []string { return strings.Fields(s) }
func Words(s string) []string {
	Count++
	if s == "again" { return Words("next") }
	return strings.Fields(s)
}
func CamelCase(s string) string { return strings.Join(Words(s), "") }
func Invoke(s string) []string { return Callback(s) }
func Shadow(Words func(string) []string) []string { return Words("parameter") }
func Local() []string {
	Words := object.Words
	return Words("local")
}
func Selectors() []string { return reader{}.Words("receiver") }
func FunctionValue() func(string) []string { return Words }
func hidden() string { return label }
func Lookup() int { return entries[Key] }
// --- Auto-generated FFI wrappers ---
invalid old generated wrapper
`
	const assertions = `package main
import (
	"strings"
	"testing"
)
func TestRenamedFFI(t *testing.T) {
	if Data_String_Extra_CamelCase("two words") != "twowords" { t.Fatal("internal call") }
	if Data_String_Extra_Invoke("again")[0] != "next" { t.Fatal("recursive callback") }
	if Data_String_Extra_FunctionValue()("value")[0] != "value" { t.Fatal("function value") }
	if Data_String_Extra_Count != 4 { t.Fatalf("count: %d", Data_String_Extra_Count) }
	if Data_String_Extra_Shadow(strings.Fields)[0] != "parameter" { t.Fatal("parameter shadowing") }
	if Data_String_Extra_Local()[0] != "local" { t.Fatal("local shadowing") }
	if Data_String_Extra_Selectors()[0] != "receiver" { t.Fatal("selector") }
	if Data_String_Extra_Lookup() != 1 { t.Fatal("map key reference") }
	if hidden() != "Words Callback Count" { t.Fatal("string literal or private declaration") }
}
`
	for _, withPackage := range []bool{true, false} {
		name := "with package"
		content := source
		if !withPackage {
			name = "without package"
			content = strings.TrimPrefix(content, "package main\n")
		}
		t.Run(name, func(t *testing.T) {
			got, err := parseFFI(content, "Data_String_Extra_")
			if err != nil {
				t.Fatal(err)
			}
			var result struct {
				Decls   []FFIDecl `json:"decls"`
				Content string    `json:"content"`
			}
			if err := json.Unmarshal([]byte(got), &result); err != nil {
				t.Fatal(err)
			}
			for _, decl := range result.Decls {
				// Methods are not package declarations and keep their selector names.
				if decl.Name != "Words" && !strings.HasPrefix(decl.Name, "Data_String_Extra_") {
					t.Errorf("declaration not renamed: %s", decl.Name)
				}
			}
			for _, preserved := range []string{
				"// Words, Callback and Count stay unchanged in comments.",
				`var label = "Words Callback Count"`,
				"Words := object.Words",
				"return Words(\"local\")",
				"func (reader) Words(s string)",
				"{Words: strings.Fields}",
			} {
				if !strings.Contains(result.Content, preserved) {
					t.Errorf("source fragment changed: %s", preserved)
				}
			}
			before, _, _ := strings.Cut(content, "// --- Auto-generated FFI wrappers ---")
			if strings.Count(before, "\n") != strings.Count(result.Content, "\n") {
				t.Error("source line count changed")
			}
			if strings.Contains(result.Content, "invalid old generated wrapper") {
				t.Error("generated wrappers were retained")
			}
			if strings.HasPrefix(result.Content, "package ") {
				t.Fatal("package retained in prepared source")
			}
			result.Content = "package main\n" + result.Content
			dir := t.TempDir()
			ffiPath := filepath.Join(dir, "ffi.go")
			testPath := filepath.Join(dir, "ffi_test.go")
			for path, contents := range map[string]string{ffiPath: result.Content, testPath: assertions} {
				if err := os.WriteFile(path, []byte(contents), 0600); err != nil {
					t.Fatal(err)
				}
			}
			cmd := exec.Command("go", "test", ffiPath, testPath)
			if output, err := cmd.CombinedOutput(); err != nil {
				t.Fatalf("renamed Go failed: %v\n%s\n%s", err, output, result.Content)
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
