package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unicode/utf16"
)

func packageEvidence(t *testing.T) string {
	t.Helper()
	if parent := os.Getenv("GOPURS_TEST_EVIDENCE"); parent != "" {
		directory, err := os.MkdirTemp(parent, "packages-")
		if err != nil {
			t.Fatal(err)
		}
		t.Log("package evidence:", directory)
		return directory
	}
	return t.TempDir()
}

func packagePut(t *testing.T, root, path, content string) {
	t.Helper()
	file := filepath.Join(root, filepath.FromSlash(path))
	if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func packageRun(t *testing.T, directory, label string, args ...string) []byte {
	t.Helper()
	command := exec.Command("go", args...)
	command.Dir = directory
	command.Env = append(os.Environ(), "GOWORK=off", "GOMAXPROCS=2")
	output, err := command.CombinedOutput()
	packagePut(t, directory, label+".log", string(output))
	packagePut(t, directory, label+".command.json", identityJSON(args))
	if err != nil {
		t.Fatalf("go %v: %v\n%s", args, err, output)
	}
	return output
}

func packageExports(t *testing.T, directory string, packages ...string) map[string][]byte {
	t.Helper()
	args := append([]string{"list", "-deps", "-export", "-json"}, packages...)
	output := packageRun(t, directory, "exports", args...)
	decoder := json.NewDecoder(bytes.NewReader(output))
	exports := map[string][]byte{}
	for {
		var entry struct{ ImportPath, Export string }
		err := decoder.Decode(&entry)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		if entry.Export != "" {
			data, err := os.ReadFile(entry.Export)
			if err != nil {
				t.Fatal(err)
			}
			exports[entry.ImportPath] = data
		}
	}
	return exports
}

func TestPackagesPreserveScopesTypesAndInitialization(t *testing.T) {
	directory := packageEvidence(t)
	packagePut(t, directory, "go.mod", "module example\n\ngo 1.22\n")
	packagePut(t, directory, "events/events.go", `package events
var Trace string
func Add(s string) int { Trace += s; return 1 }
func init() { Trace += "external;" }
`)
	request := PackageRequest{BindRequest: BindRequest{PackagePath: "example/purescript", GoVersion: "go1.22", Arch: "arm64", Files: []SourceInput{
		{Owner: "generated/Model", Module: "Model", Path: "Model.go", Content: `package purescript
import "sync"
type Box[T any] struct { Key T }
func (renamed *Box[T]) Read() T { return renamed.Key }
var cache_Model_value *Box[int]
var once_Model_value sync.Once
var Evaluations int
func Get_Model_value() *Box[int] {
  once_Model_value.Do(func(){ Evaluations++; cache_Model_value = &Box[int]{Key:21} })
  return cache_Model_value
}
func Call_Model_value() *Box[int] { return Get_Model_value() }
type Readable[T any] interface { Read() T }
func Recursive(n int) int { if n==0 {return 1}; return Other(n-1) }
func Other(n int) int { if n==0 {return 1}; return Recursive(n-1) }
`},
		{Owner: "generated/Consumer", Module: "Consumer", Path: "Consumer.go", Content: `package purescript
type Wrapped struct { Box[int] }
type IntMap map[int]int
const Key = 3
func Use() int {
  value := Box[int]{Key: Get_Model_value().Read()}
  entries := IntMap{Key: 10}
  callback := func(Get_Model_value func() *Box[int]) int {return Get_Model_value().Key}
  var boxed any = value
  { Key := 99; _ = Box[int]{Key:Key} }
  return boxed.(Box[int]).Key + entries[Key] + callback(Get_Model_value) - 10
}
func ReadWrapped(value Wrapped) int {return value.Box.Read()}
func ReadBox(value Readable[int]) int {return value.Read()}
func Same() bool {return Get_Model_value() == Call_Model_value()}
func Result() int {return Evaluations}
`},
		{Owner: "ffi/Effects", Module: "Effects", Path: "Effects_ffi.go", Content: `package purescript
import "example/events"
var First = events.Add("A;")
var _ = events.Add("B;")
func init() {events.Add("C;")}
func NeverCalled() int {return First}
`},
		{Owner: "generated/Later", Module: "Later", Path: "Later.go", Initializers: []SourceSymbol{{"generated/Later", "probe/later"}}, Content: `package purescript
import "example/events"
var Later = events.Add("D;")
func init() {events.Add("E;")}
`},
	}}, Entries: []SourceSymbol{{"generated/Consumer", "Use"}, {"generated/Consumer", "Same"}, {"generated/Consumer", "Result"}}}
	request.Exports = packageExports(t, directory, "sync", "example/events")
	for _, source := range request.Files {
		packagePut(t, directory, "mono/"+source.Path, source.Content)
	}
	main := `package main
import("fmt"; "sync"; "example/events"; use %q; same %q; result %q; %s)
func main(){
  var group sync.WaitGroup
  for i:=0;i<64;i++ {group.Add(1);go func(){defer group.Done(); if !same.Same() {panic("distinct cells")}}()}
  group.Wait()
  fmt.Printf("%%d/%%d/%%s",use.Use(),result.Result(),events.Trace)
}
`
	packagePut(t, directory, "before/main.go", fmt.Sprintf(main, "example/mono", "example/mono", "example/mono", ""))
	want := packageRun(t, directory, "before", "run", "-race", "./before")
	if string(want) != "42/1/external;A;B;D;C;E;" {
		t.Fatalf("unexpected independent oracle: %s", want)
	}
	result, err := packageSources(request, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Packages) < 5 {
		t.Fatal("fixture did not exercise package boundaries")
	}
	data, _ := json.Marshal(result)
	packagePut(t, directory, "result.json", string(data))
	for _, file := range result.Files {
		packagePut(t, directory, "purescript/"+file.Path, file.Content)
	}
	packagePut(t, directory, "after/main.go", fmt.Sprintf(main, result.Entries[0].Package, result.Entries[1].Package, result.Entries[2].Package, `_ "example/purescript/init_all"`))
	got := packageRun(t, directory, "after", "run", "-race", "./after")
	if !bytes.Equal(got, want) {
		t.Fatalf("package execution differs: %s != %s", got, want)
	}
	packageRun(t, directory, "dag", "list", "-deps", "-json", "./after")
	for left, right := 0, len(request.Files)-1; left < right; left, right = left+1, right-1 {
		request.Files[left], request.Files[right] = request.Files[right], request.Files[left]
	}
	reordered, err := packageSources(request, 1)
	if err != nil {
		t.Fatal(err)
	}
	other, _ := json.Marshal(reordered)
	if !bytes.Equal(data, other) {
		t.Fatal("source arrival order changed package bytes")
	}
}

func TestPackageSCCMatchesReachability(t *testing.T) {
	random := rand.New(rand.NewSource(2303))
	for trial := 0; trial < 95; trial++ {
		n := 1 + random.Intn(30)
		edges, reach := make([][]int, n), make([][]bool, n)
		for i := 0; i < n; i++ {
			reach[i] = make([]bool, n)
			reach[i][i] = true
			for j := 0; j < n; j++ {
				if random.Intn(n) < 2 {
					edges[i] = append(edges[i], j)
					reach[i][j] = true
				}
			}
		}
		for k := 0; k < n; k++ {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					reach[i][j] = reach[i][j] || reach[i][k] && reach[k][j]
				}
			}
		}
		componentOf := make([]int, n)
		for i, component := range stronglyConnected(edges) {
			for _, node := range component {
				componentOf[node] = i
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if (componentOf[i] == componentOf[j]) != (reach[i][j] && reach[j][i]) {
					t.Fatalf("SCC differs from reachability on trial %d, %d/%d", trial, i, j)
				}
			}
		}
	}
}

func TestPackageAliasesAndBodyStability(t *testing.T) {
	paths := []string{"example/a", "example/b"}
	reserved := map[string]bool{"g" + sourceHash(paths[0])[:6]: true, "g" + sourceHash(paths[0])[:7]: true}
	aliases, err := packageAliases(paths, reserved)
	if err != nil || len(aliases[paths[0]]) != 9 {
		t.Fatalf("alias shadow not avoided: %v %v", aliases, err)
	}
	reordered, err := packageAliases([]string{paths[1], paths[0]}, reserved)
	if err != nil || !reflect.DeepEqual(aliases, reordered) {
		t.Fatal("order-dependent aliases")
	}
	fixture := bindingFixture()
	for i := range fixture.Files {
		_, module, _ := strings.Cut(fixture.Files[i].Owner, "/")
		fixture.Files[i].Module = module
	}
	before, err := packageSources(PackageRequest{BindRequest: fixture}, 1)
	if err != nil {
		t.Fatal(err)
	}
	fixture.Files[0].Content = strings.ReplaceAll(fixture.Files[0].Content, "return 21", "return 22")
	after, err := packageSources(PackageRequest{BindRequest: fixture}, 1)
	if err != nil || before.PlanHash != after.PlanHash {
		t.Fatalf("body-only edit changed placement: %v", err)
	}
}

func TestRuntimeImportAliasPreservesShadowing(t *testing.T) {
	directory := packageEvidence(t)
	packagePut(t, directory, "go.mod", "module gopurs/output\n\ngo 1.22\n")
	packagePut(t, directory, "gopurs_runtime/runtime.go", "package gopurs_runtime\ntype Value int\nfunc Int(n int) Value {return Value(n)}\n")
	const header = `package purescript
import "gopurs/output/gopurs_runtime"
const r = 1
func Probe() int {
  r_ := 2
  value := gopurs_runtime.Int(0)
`
	const footer = `
  { gopurs_runtime := struct{Value gopurs_runtime.Value}{3}; value += gopurs_runtime.Value }
  if "gopurs_runtime.Value" != "gopurs_runtime.Value" { panic("literal changed") }
  return int(value)+r+r_
}
`
	content := header + strings.Repeat("  value += gopurs_runtime.Value(gopurs_runtime.Int(1))\n", 32) + footer
	request := PackageRequest{BindRequest: BindRequest{PackagePath: "gopurs/output/purescript", GoVersion: "go1.22", Arch: "arm64",
		Files: []SourceInput{{Owner: "ffi/Main", Module: "Main", Path: "Main_ffi.go", Content: content}}},
		Entries: []SourceSymbol{{"ffi/Main", "Probe"}}}
	request.Exports = packageExports(t, directory, "gopurs/output/gopurs_runtime")
	packagePut(t, directory, "mono/Main.go", content)
	main := "package main\nimport(\"fmt\"; probe %q)\nfunc main(){fmt.Println(probe.Probe())}\n"
	packagePut(t, directory, "before/main.go", fmt.Sprintf(main, "gopurs/output/mono"))
	want := packageRun(t, directory, "before", "run", "./before")
	if string(want) != "38\n" {
		t.Fatalf("independent oracle: %s", want)
	}
	result, err := packageSources(request, 1)
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(result)
	packagePut(t, directory, "result.json", string(data))
	for _, file := range result.Files {
		packagePut(t, directory, "purescript/"+file.Path, file.Content)
		if (strings.HasSuffix(file.Path, "/Main_ffi.go") || strings.HasSuffix(file.Path, "/"+packageBundleName)) && len(file.Content) >= len(content)-500 {
			t.Fatalf("runtime qualifiers still dominate fragment storage: %d -> %d", len(content), len(file.Content))
		}
	}
	packagePut(t, directory, "after/main.go", fmt.Sprintf(main, result.Entries[0].Package))
	if got := packageRun(t, directory, "after", "run", "./after"); !bytes.Equal(got, want) {
		t.Fatalf("alias changed shadowing semantics: %s != %s", got, want)
	}
}

func TestPackageCoalescingPreservesImportScopesAndInitOrder(t *testing.T) {
	directory := packageEvidence(t)
	packagePut(t, directory, "go.mod", "module example\n\ngo 1.22\n")
	packagePut(t, directory, "events/events.go", "package events\nvar Trace string\n")
	request := PackageRequest{BindRequest: BindRequest{PackagePath: "example/purescript", GoVersion: "go1.22", Arch: "arm64",
		Files: []SourceInput{
			{Owner: "ffi/A", Module: "A", Path: "A_ffi.go", Content: `package purescript
import(shared "fmt"; e "example/events")
func A(n int) string { if n == 0 { return shared.Sprint("a") }; return B(n-1) }
func init() { e.Trace += "A" }
`},
			{Owner: "ffi/B", Module: "B", Path: "B_ffi.go", Content: `package purescript
import(shared "strings"; . "math"; m "math"; e "example/events")
func B(n int) string { shared := func(s string) string {return s}; if n == 0 {return shared("b")}; return A(n-1) }
func Probe() string { if !IsNaN(m.NaN()) || Pi < 3 {panic("dot import")}; return e.Trace + shared.ToUpper(A(2)) }
func init() { e.Trace += "B" }
`},
		}}, Entries: []SourceSymbol{{"ffi/B", "Probe"}}}
	request.Exports = packageExports(t, directory, "fmt", "strings", "math", "example/events")
	for _, source := range request.Files {
		packagePut(t, directory, "mono/"+source.Path, source.Content)
	}
	main := "package main\nimport(\"fmt\"; probe %q)\nfunc main(){fmt.Println(probe.Probe())}\n"
	packagePut(t, directory, "before/main.go", fmt.Sprintf(main, "example/mono"))
	want := packageRun(t, directory, "before", "run", "./before")
	if string(want) != "ABA\n" {
		t.Fatalf("independent oracle: %s", want)
	}
	result, err := packageSources(request, 1)
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(result)
	packagePut(t, directory, "result.json", string(data))
	for _, file := range result.Files {
		packagePut(t, directory, "purescript/"+file.Path, file.Content)
	}
	packagePut(t, directory, "after/main.go", fmt.Sprintf(main, result.Entries[0].Package))
	if got := packageRun(t, directory, "after", "run", "./after"); !bytes.Equal(got, want) {
		t.Fatalf("coalesced scopes/order: %s != %s", got, want)
	}
	if len(result.Packages) != 1 || len(result.Files) != 4 {
		t.Fatalf("source fragmentation retained: %d packages, %d files", len(result.Packages), len(result.Files))
	}
	request.PackagePath = "example/directives"
	request.Files[0].Content = "//go:build darwin || linux\n\n" + request.Files[0].Content
	constrained, err := packageSources(request, 1)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range constrained.Files {
		if strings.HasSuffix(file.Path, "/"+packageBundleName) {
			t.Fatal("file directive was coalesced")
		}
		packagePut(t, directory, "directives/"+file.Path, file.Content)
	}
	packagePut(t, directory, "constrained/main.go", fmt.Sprintf(main, constrained.Entries[0].Package))
	if got := packageRun(t, directory, "constrained", "run", "./constrained"); !bytes.Equal(got, want) {
		t.Fatalf("file directive changed execution: %s", got)
	}
}

func TestRecordABIChangesOnlyBoundFields(t *testing.T) {
	directory := packageEvidence(t)
	packagePut(t, directory, "go.mod", "module example\n\ngo 1.22\n")
	recordType := "struct{value int64; go__type string; gopurs_field_55296 int}"
	request := PackageRequest{BindRequest: BindRequest{PackagePath: "example/split", GoVersion: "go1.22", Arch: "arm64", Files: []SourceInput{
		{Owner: "generated/Provider", Module: "Provider", Path: "Provider.go", Content: `package purescript
func Make() ` + recordType + ` {return ` + recordType + `{value:21, go__type:"record", gopurs_field_55296:7}}
func Identity[T any](v T) T {return v}
`},
		{Owner: "generated/Consumer", Module: "Consumer", Path: "Consumer.go", Content: `package purescript
func Accept(v ` + recordType + `) int64 {return v.value}
func Use() int64 { v := Identity(Make()); return Accept(v) + ReadForeign(v) }
func Tuple() struct{V0 int} {return struct{V0 int}{V0:8}}
`},
		{Owner: "ffi/Foreign", Module: "Foreign", Path: "Foreign_ffi.go", Content: `package purescript
func ReadForeign(v any) int64 { r := v.(struct { value int64; go__type string; gopurs_field_55296 int }); if r.go__type!="record" || r.gopurs_field_55296!=7 {panic("labels/layout")}; return r.value }
type Native struct{value string}
func (v Native) Read() string {return v.value}
`},
	}}, Entries: []SourceSymbol{{"generated/Consumer", "Use"}}}
	for _, source := range request.Files {
		packagePut(t, directory, "mono/"+source.Path, source.Content)
	}
	packagePut(t, directory, "before/main.go", "package main\nimport(\"fmt\"; mono \"example/mono\")\nfunc main(){fmt.Print(mono.Use())}\n")
	want := packageRun(t, directory, "before", "run", "./before")
	naive, err := packageSources(request, 1)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range naive.Files {
		packagePut(t, directory, "split/"+file.Path, file.Content)
	}
	packagePut(t, directory, "rejected/main.go", "package main\nimport p "+strconv.Quote(naive.Entries[0].Package)+"\nfunc main(){p.Use()}\n")
	command := exec.Command("go", "build", "-o", filepath.Join(directory, "rejected-bin"), "./rejected")
	command.Dir, command.Env = directory, append(os.Environ(), "GOWORK=off")
	output, err := command.CombinedOutput()
	packagePut(t, directory, "rejected.log", string(output))
	if err == nil || !bytes.Contains(output, []byte("cannot use")) {
		t.Fatalf("unadapted private-record ABI must fail: %v\n%s", err, output)
	}
	request.PackagePath = "example/records"
	request.Files[0].Records = []SourceRecord{{recordType, []string{"00760061006c00750065", "0074007900700065", "d800"}}}
	request.Files[1].Records = request.Files[0].Records
	request.Files[2].Records = []SourceRecord{{"struct { value int64; go__type string; gopurs_field_55296 int }", request.Files[0].Records[0].Labels}}
	result, err := packageSources(request, 1)
	if err != nil {
		t.Fatal(err)
	}
	var all strings.Builder
	for _, file := range result.Files {
		packagePut(t, directory, "records/"+file.Path, file.Content)
		all.WriteString(file.Content)
	}
	for _, expected := range []string{"P_value", "P_type", "PX_d800", "type Native struct{value string}", "struct{V0 int}"} {
		if !strings.Contains(all.String(), expected) {
			t.Fatalf("lost field distinction: %s", expected)
		}
	}
	packagePut(t, directory, "after/main.go", "package main\nimport(\"fmt\"; p "+strconv.Quote(result.Entries[0].Package)+"; _ \"example/records/init_all\")\nfunc main(){fmt.Print(p.Use())}\n")
	got := packageRun(t, directory, "after", "run", "./after")
	if !bytes.Equal(got, want) || string(got) != "42" {
		t.Fatalf("record layout/boxing changed: %s != %s", got, want)
	}
	annotated := request
	annotated.Files = append([]SourceInput(nil), request.Files...)
	annotated.PackagePath = "example/annotated"
	annotated.Files[0].Records = nil
	annotated.Files[0].Content = strings.ReplaceAll(annotated.Files[0].Content, recordType,
		strings.TrimSuffix(recordType, "}")+"\n/*gopurs:record-v1:3:00760061006c00750065,0074007900700065,d800*/\n}")
	marked, err := packageSources(annotated, 1)
	if err != nil {
		t.Fatal(err)
	}
	if marked.PlanHash != result.PlanHash {
		t.Fatal("producer annotations changed logical placement")
	}
	for _, file := range marked.Files {
		if strings.Contains(file.Content, "/*gopurs:record-") {
			t.Fatal("private record annotation leaked into package output")
		}
		packagePut(t, directory, "annotated/"+file.Path, file.Content)
	}
	packagePut(t, directory, "annotations/main.go", "package main\nimport(\"fmt\"; p "+strconv.Quote(marked.Entries[0].Package)+"; _ \"example/annotated/init_all\")\nfunc main(){fmt.Print(p.Use())}\n")
	if got := packageRun(t, directory, "annotations", "run", "./annotations"); !bytes.Equal(got, want) {
		t.Fatalf("annotated record changed output: %s", got)
	}
	request.Files[1].Records = append([]SourceRecord{request.Files[1].Records[0]}, SourceRecord{recordType, []string{"00760061006c00750065", "0067006f005f005f0074007900700065", "d800"}})
	if _, err := packageSources(request, 1); err == nil || !strings.Contains(err.Error(), "ambiguous record labels") {
		t.Fatalf("ambiguous legacy record spellings were accepted: %v", err)
	}
}

func TestExportedRecordLabelsAreInjectiveUTF16(t *testing.T) {
	labels := []string{"", "_", "type", "go__type", "a-b", "a_minus_b", "P_value", "PX_", "1", "a1", "é", "💡", "*/"}
	seen := map[string]bool{}
	for _, label := range labels {
		var encoded strings.Builder
		for _, unit := range utf16.Encode([]rune(label)) {
			fmt.Fprintf(&encoded, "%04x", unit)
		}
		name, err := exportedRecordField(encoded.String())
		if err != nil || seen[name] || !strings.HasPrefix(name, "P") {
			t.Fatalf("invalid record label mapping: %q -> %q, %v", label, name, err)
		}
		seen[name] = true
	}
	for input, expected := range map[string]string{"d800": "PX_d800", "d801": "PX_d801", "d83ddca1": "PX_d83ddca1", "0074007900700065": "P_type"} {
		got, err := exportedRecordField(input)
		if err != nil || got != expected || seen[got] && input != "0074007900700065" && input != "d83ddca1" {
			t.Fatalf("UTF-16 identity changed: %s -> %s, %v", input, got, err)
		}
	}
	for _, input := range []string{"d80", "D800", "zzzz"} {
		if _, err := exportedRecordField(input); err == nil {
			t.Fatalf("invalid UTF-16 label accepted: %q", input)
		}
	}
}

func TestRecordAnnotationsDoNotClaimEqualTupleOrNativeLayouts(t *testing.T) {
	request := PackageRequest{BindRequest: BindRequest{PackagePath: "example/purescript", GoVersion: "go1.22", Arch: "arm64", Files: []SourceInput{
		{Owner: "generated/Record", Module: "Record", Path: "Record.go", Content: `package purescript
func MakeRecord() struct{V0 int64
/*gopurs:record-v1:1:00560030*/
} { return struct{V0 int64
/*gopurs:record-v1:1:00560030*/
}{V0:42} }
`},
		{Owner: "generated/Tuple", Module: "Tuple", Path: "Tuple.go", Content: `package purescript
func Tuple() struct{V0 int64} {return struct{V0 int64}{V0:42}}
`},
		{Owner: "ffi/Native", Module: "Native", Path: "Native_ffi.go", Content: `package purescript
type Native struct{V0 int64}
func MakeNative() Native {return Native{V0:42}}
`},
	}}, Entries: []SourceSymbol{{"generated/Tuple", "Tuple"}, {"ffi/Native", "MakeNative"}}}
	result, err := packageSources(request, 1)
	if err != nil {
		t.Fatal(err)
	}
	directory := packageEvidence(t)
	packagePut(t, directory, "go.mod", "module example\n\ngo 1.22\n")
	for _, file := range result.Files {
		packagePut(t, directory, "purescript/"+file.Path, file.Content)
	}
	packagePut(t, directory, "main.go", fmt.Sprintf(`package main
import("fmt"; tuples %q; native %q)
func main(){fmt.Print(tuples.Tuple().V0 + native.MakeNative().V0)}
`, result.Entries[0].Package, result.Entries[1].Package))
	if got := packageRun(t, directory, "unrelated-ABI", "run", "."); string(got) != "84" {
		t.Fatalf("unrelated layouts changed: %s", got)
	}
}
