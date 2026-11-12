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

func parseDeclarations(t *testing.T, source string) []FFIDecl {
	t.Helper()
	encoded, err := Extract(source)
	if err != nil {
		t.Fatal(err)
	}
	var decls []FFIDecl
	if err := json.Unmarshal([]byte(encoded), &decls); err != nil {
		t.Fatal(err)
	}
	if decls == nil {
		t.Fatal("declarations must be an array, not null")
	}
	for _, decl := range decls {
		if decl.SchemaVersion != 2 || decl.Args == nil || decl.ArgNames == nil || decl.Results == nil || decl.ResultNames == nil {
			t.Fatalf("incomplete v2 declaration: %+v", decl)
		}
	}
	return decls
}

func TestParseFFIContract(t *testing.T) {
	decls := parseDeclarations(t, `package ffi
func Add(a, b int, label string) (first, second int, err error) { return 0, 0, nil }
func hidden() {}
func _Private() {}
func Étranger() {}
type Receiver struct{}
func (Receiver) Method() {}
var Exported, hiddenVariable, _Internal int
var Callback func(string) (bool, error)
var Inferred = 1
const Constant = 2
type Alias = int`)
	names := []string{}
	for _, decl := range decls {
		names = append(names, decl.Name)
	}
	if strings.Join(names, ",") != "Add,_Private,Method,Exported,_Internal,Callback,Inferred" {
		t.Fatal(names)
	}
	add := decls[0]
	if len(add.Args) != 3 || strings.Join(add.ArgNames, ",") != "a,b,label" {
		t.Fatal(add)
	}
	if len(add.Results) != 3 || strings.Join(add.ResultNames, ",") != "first,second,err" || add.Ret != nil {
		t.Fatalf("grouped results lost or truncated: %+v", add)
	}
	if add.Results[0].Name != "int" || add.Results[1].Name != "int" || add.Results[2].Identity != "builtin.error" {
		t.Fatal(add.Results)
	}
	if add.Line != 2 || add.Column != 6 || add.Signature != "func(a, b int, label string) (first, second int, err error)" {
		t.Fatal(add)
	}
	if decls[2].Receiver.Name != "Receiver" {
		t.Fatal(decls[2])
	}
	if !decls[3].IsVar || decls[3].ValueType.Name != "int" || len(decls[3].Results) != 0 {
		t.Fatal(decls[3])
	}
	if decls[5].ValueType.Type != "Func" || len(decls[5].ValueType.Results) != 2 || decls[5].ValueType.Ret != nil {
		t.Fatal(decls[5])
	}
	if decls[6].ValueType != nil {
		t.Fatal("inferred variables must not claim a guessed type")
	}
	for _, source := range []string{"", "package ffi\nfunc hidden() {}\nvar hiddenVariable int"} {
		if got := parseDeclarations(t, source); len(got) != 0 {
			t.Fatal(got)
		}
	}
	marker := parseDeclarations(t, "func Before() {}\n// --- Auto-generated FFI wrappers ---\ninvalid Go")
	if len(marker) != 1 || marker[0].Name != "Before" {
		t.Fatal(marker)
	}
}

func TestParseFFICompleteTypes(t *testing.T) {
	decl := parseDeclarations(t, `import ctx "context"
func Convert[A, B any, K comparable](values []A, lookup map[K]B, fixed [2]int, context ctx.Context, arbitrary interface{}, pointer *A, values2 ...string) []B { return nil }`)[0]
	if strings.Join(decl.TypeParams, ",") != "A,B,K" {
		t.Fatal(decl.TypeParams)
	}
	args := decl.Args
	if len(args) != 7 || args[0].Type != "Array" || args[0].Elem.Name != "A" || args[1].Type != "Map" || args[1].Key.Name != "K" || args[1].Val.Name != "B" {
		t.Fatal(args)
	}
	if args[2].Type != "FixedArray" || args[2].Length != "2" || args[2].Elem.Name != "int" {
		t.Fatalf("fixed array collapsed into slice: %+v", args[2])
	}
	if args[3].Identity != "context.Context" || args[3].PackagePath != "context" || args[3].Name != "ctx.Context" {
		t.Fatal(args[3])
	}
	if args[4].Name != "any" || args[5].Type != "Pointer" || args[5].Elem.Name != "A" || args[6].Type != "Variadic" || args[6].Elem.Name != "string" {
		t.Fatal(args)
	}
	if len(decl.Results) != 1 || decl.Ret.Type != "Array" || decl.Ret.Elem.Name != "B" {
		t.Fatal(decl)
	}
	nested := parseDeclarations(t, `func Call(cb func(left, right int, next func() string) (func(func(bool)), error), empty func()) {}`)[0]
	callback := nested.Args[0]
	if callback.Type != "Func" || len(callback.Args) != 3 || strings.Join(callback.ArgNames, ",") != "left,right,next" || len(callback.Results) != 2 || callback.Ret != nil {
		t.Fatal(callback)
	}
	if callback.Args[2].Ret.Name != "string" || callback.Results[0].Type != "Func" || callback.Results[1].Identity != "builtin.error" {
		t.Fatal(callback)
	}
	if nested.Args[1].Results == nil || len(nested.Args[1].Results) != 0 || nested.Args[1].ResultNames == nil {
		t.Fatal(nested.Args[1])
	}
	unknown := parseDeclarations(t, `func Unknown(generic Box[int], contract interface{ Read([]byte) (int, error) }) chan int { return nil }`)[0]
	if unknown.Args[0].Type != "Unknown" || unknown.Args[0].Name != "Box[int]" || unknown.Args[1].Type != "Interface" || !strings.Contains(unknown.Args[1].Name, "Read") || unknown.Ret.Name != "chan int" {
		t.Fatal(unknown)
	}
}

func TestParseFFITypeIdentity(t *testing.T) {
	cases := []struct{ name, source, argument, result string }{
		{"builtin and aliased import", `import c "context"; func Run(c c.Context) error { return nil }`, "context.Context", "builtin.error"},
		{"dot imports cannot shadow unexported universe error", `import . "example/sdk"; func Run(string) error { return nil }`, "", "builtin.error"},
		{"transitive aliases", `import c "context"; type C = c.Context; type D = C; type E = error; type F = E; func Run(D) F { return nil }`, "context.Context", "builtin.error"},
		{"parenthesized aliases", `import "context"; type C = (context.Context); type E = (error); func Run(C) E { return nil }`, "context.Context", "builtin.error"},
		{"shadowed error", `type error int; func Run(string) error { return 0 }`, "", ""},
		{"new named types", `import "context"; type C context.Context; type E error; func Run(C) E { return nil }`, "", ""},
		{"type parameter shadow", `func Run[error any](error) error { panic("unused") }`, "", ""},
		{"lookalike import alias", `import context "example/other"; func Run(context.Context) error { return nil }`, "", "builtin.error"},
		{"unresolved selector", `func Run(context.Context) error { return nil }`, "", "builtin.error"},
		{"external alias unproven", `import c "example/types"; func Run(c.Context) c.Error { return nil }`, "", ""},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			decl := parseDeclarations(t, test.source)[0]
			if decl.Args[0].Identity != test.argument || decl.Results[0].Identity != test.result {
				t.Fatalf("identity: argument %+v, result %+v", decl.Args[0], decl.Results[0])
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

func TestParseFFIStructLayouts(t *testing.T) {
	const source = "package ffi\n" + `
type Address struct { Street string }
type Payload struct {
 Name, Other string
 Count int ` + "`json:\"count,omitempty\" db:\"count\"`" + `
 Nested struct { Enabled bool; Labels []string }
 Addr *Address
 Address
 hidden bool
}
type Aliased = Payload
type Defined Payload
type Recursive struct { Next *Recursive }
func Echo(value Payload) Payload { return value }
func Alias(value Aliased) Defined { return Defined(value) }
func Node(value Recursive) Recursive { return value }
func Anonymous(value struct { Key string }) struct { Key string } { return value }
`
	decls := parseDeclarations(t, source)
	payload := decls[0].Args[0]
	if payload.Name != "Payload" || payload.Underlying == nil || payload.Underlying.Type != "Struct" {
		t.Fatalf("missing local layout: %+v", payload)
	}
	fields := payload.Underlying.Fields
	if len(fields) != 7 {
		t.Fatalf("fields: %+v", fields)
	}
	if fields[0].Name != "Name" || fields[1].Name != "Other" || fields[0].Type.Name != "string" || fields[1].Type.Name != "string" || !fields[0].Exported {
		t.Fatal(fields[:2])
	}
	if fields[2].Name != "Count" || fields[2].Tag != `json:"count,omitempty" db:"count"` || fields[2].Embedded {
		t.Fatal(fields[2])
	}
	if fields[3].Type.Type != "Struct" || len(fields[3].Type.Fields) != 2 || fields[3].Type.Fields[1].Type.Type != "Array" {
		t.Fatal(fields[3])
	}
	if fields[4].Type.Type != "Pointer" || fields[4].Type.Elem.Name != "Address" || fields[4].Type.Elem.Underlying.Fields[0].Name != "Street" {
		t.Fatal(fields[4])
	}
	if fields[5].Name != "Address" || !fields[5].Embedded || !fields[5].Exported {
		t.Fatal(fields[5])
	}
	if fields[6].Name != "hidden" || fields[6].Exported {
		t.Fatal(fields[6])
	}
	if decls[1].Args[0].Name != "Aliased" || !hasStructLayout(decls[1].Args[0]) || !hasStructLayout(decls[1].Results[0]) {
		t.Fatal(decls[1])
	}
	next := decls[2].Args[0].Underlying.Fields[0].Type
	if next.Type != "Pointer" || next.Elem.Name != "Recursive" || next.Elem.Underlying != nil {
		t.Fatal("recursive layout must terminate at the named reference")
	}
	if decls[3].Args[0].Type != "Struct" || decls[3].Results[0].Type != "Struct" {
		t.Fatal(decls[3])
	}

	prepared, err := Prepare("Fixture_", source)
	if err != nil {
		t.Fatal(err)
	}
	var result struct{ Content string }
	if err := json.Unmarshal([]byte(prepared), &result); err != nil {
		t.Fatal(err)
	}
	for _, fragment := range []string{"type Payload struct", "func Fixture_Echo(value Payload) Payload", "Name, Other string", "json:\"count,omitempty\""} {
		if !strings.Contains(result.Content, fragment) {
			t.Fatalf("struct source corrupted: %s", result.Content)
		}
	}
	directory := t.TempDir()
	path := filepath.Join(directory, "ffi.go")
	if err := os.WriteFile(path, []byte("package ffi\n"+result.Content), 0600); err != nil {
		t.Fatal(err)
	}
	if output, err := exec.Command("go", "test", path).CombinedOutput(); err != nil {
		t.Fatalf("prepared struct source failed to compile: %v\n%s", err, output)
	}
}

func TestParseFFIStructEmbeddingAndUnknownImports(t *testing.T) {
	decl := parseDeclarations(t, `import remote "example/remote"
type Container struct {
 *remote.Handle
 remote.Generic[int]
 Remote remote.Record
}
func Read(value Container) remote.Record { return value.Remote }`)[0]
	fields := decl.Args[0].Underlying.Fields
	if len(fields) != 3 || fields[0].Name != "Handle" || !fields[0].Embedded || fields[1].Name != "Generic" || !fields[1].Embedded {
		t.Fatal(fields)
	}
	if fields[2].Type.Underlying != nil || decl.Results[0].Underlying != nil {
		t.Fatal("external struct layouts must not be guessed")
	}
	empty := parseDeclarations(t, `func Empty(value struct{}) struct{} { return value }`)[0]
	if empty.Args[0].Type != "Struct" || len(empty.Args[0].Fields) != 0 {
		t.Fatal(empty)
	}
}
