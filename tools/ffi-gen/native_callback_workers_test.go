package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const callbackFixture = `package main
import (
 r "gopurs/output/gopurs_runtime"
 "sync"
 "fmt"
)
var cache r.Value
var once sync.Once
func getter() r.Value {
 once.Do(func() { cache = r.Func2(func(k r.Value, x r.Value) r.Value { return worker(k,x) }) })
 return cache
}
func worker(k_loop r.Value, x r.Value) r.Value {
 var k r.Value = k_loop
 _ = k
 alias := k
 _ = alias
 value := struct { N int }{x.IntVal}
 return func(native func(struct { N int }) r.Value, ok bool) r.Value {
  if ok { return native(value) } else {}
  return r.Apply(alias, r.Int(value.N))
 }(r.FunctionData[func(struct { N int }) r.Value](alias))
}
var observations []int
func mark(x int) r.Value { observations=append(observations,x);return r.Int(x) }
func main() {
 seed := 7
 parser := r.Apply(getter(), r.WithFunctionData(r.Func(func(x r.Value) r.Value { return mark(seed+x.IntVal) }),func(x struct { N int }) r.Value { return mark(seed+x.N) }))
 seed = 9
 fmt.Println(len(observations),r.Apply(parser,r.Int(3)).IntVal,r.Apply(parser,r.Int(4)).IntVal)
 direct := worker(r.WithFunctionData(r.Func(func(x r.Value) r.Value { return mark(x.IntVal+10) }),func(x struct { N int }) r.Value { return mark(x.N+10) }),mark(5))
 fmt.Println(direct.IntVal,observations)
}
`

const callbackRuntime = `package gopurs_runtime
type Value struct { IntVal int; fn any; native any }
func Int(x int) Value { return Value{IntVal:x} }
func Func(f func(Value) Value) Value { return Value{fn:f} }
func Func2(f func(Value,Value) Value) Value { return Value{fn:f} }
func WithFunctionData(v Value, f any) Value { v.native=f; return v }
func FunctionData[T any](v Value) (T,bool) { f,ok:=v.native.(T);return f,ok }
func Apply(f,x Value) Value { switch fn:=f.fn.(type) { case func(Value)Value:return fn(x);case func(Value,Value)Value:return Func(func(y Value)Value{return fn(x,y)}) };panic("bad function") }
`

func TestNativeCallbackWorkersPreserveDeferredCalls(t *testing.T) {
	after, err := OptimizeNativeCallbacks(callbackFixture)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(after, "func GopursNativeCallback_worker(") != 1 || strings.Count(after, "GopursNativeCallback_worker(") != 3 {
		t.Fatalf("expected one worker and two selected calls:\n%s", after)
	}
	if _, err := parser.ParseFile(token.NewFileSet(), "after.go", after, 0); err != nil {
		t.Fatal(err)
	}
	run := func(source string) string {
		t.Helper()
		dir := t.TempDir()
		if err := os.Mkdir(filepath.Join(dir, "gopurs_runtime"), 0755); err != nil {
			t.Fatal(err)
		}
		for name, text := range map[string]string{"go.mod": "module gopurs/output\ngo 1.27.0\n", "main.go": source, "gopurs_runtime/runtime.go": callbackRuntime} {
			if err := os.WriteFile(filepath.Join(dir, name), []byte(text), 0644); err != nil {
				t.Fatal(err)
			}
		}
		command := exec.Command("go", "run", ".")
		command.Dir = dir
		result, err := command.CombinedOutput()
		if err != nil {
			t.Fatalf("run: %v\n%s", err, result)
		}
		return string(result)
	}
	beforeResult, afterResult := run(callbackFixture), run(after)
	if beforeResult != afterResult || afterResult != "0 12 13\n15 [12 13 5 15]\n" {
		t.Fatalf("evaluation or captures changed: before=%q after=%q", beforeResult, afterResult)
	}
	// The transformation is stable on cache hits / repeated frontend requests.
	again, err := OptimizeNativeCallbacks(after)
	if err != nil || again != after {
		t.Fatalf("not idempotent: %v", err)
	}
}

func TestNativeCallbackWorkersRejectObservableCallbacks(t *testing.T) {
	tests := []struct{ name, old, replacement string }{
		{"returned", "_ = alias", "if x.IntVal < 0 { return alias }; _ = alias"},
		{"stored", "_ = alias", "_ = []r.Value{alias}; _ = alias"},
		{"reassigned", "_ = alias", "alias = k; _ = alias"},
		{"address", "_ = alias", "_ = &alias; _ = alias"},
		{"unknown-call", "_ = alias", "observe(alias); _ = alias"},
		{"dispatcher-effect", "if ok { return native(value) }", "if ok { mark(99); return native(value) }"},
		{"dispatcher-extra-statement", "return r.Apply(alias, r.Int(value.N))", "mark(99); return r.Apply(alias, r.Int(value.N))"},
		{"callback-observed-in-conversion", "r.Int(value.N)", "observe(alias)"},
		{"type-shadow", "value := struct { N int }{x.IntVal}", "type int = int64; value := struct { N int }{int(x.IntVal)}"},
		{"parameter-write", "_ = k", "k_loop = k; _ = k"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source := strings.Replace(callbackFixture, test.old, test.replacement, 1)
			result, err := OptimizeNativeCallbacks(source)
			if err != nil {
				t.Fatal(err)
			}
			if result != source {
				t.Fatal("observable callback was specialized")
			}
		})
	}
}

func TestNativeCallbackWorkersLeaveUnknownArgumentEvaluation(t *testing.T) {
	// FunctionData and callbacks are not inspected through arbitrary calls.
	for _, replacement := range []string{
		"makeCallback()",
		"r.WithFunctionData(makeBoxed(),func(x struct { N int }) r.Value { return mark(x.N+10) })",
		"r.WithFunctionData(r.Func(func(x r.Value) r.Value { return mark(x.IntVal+10) }),makeNative())",
	} {
		source := strings.Replace(callbackFixture,
			"r.WithFunctionData(r.Func(func(x r.Value) r.Value { return mark(x.IntVal+10) }),func(x struct { N int }) r.Value { return mark(x.N+10) })",
			replacement, 1)
		result, err := OptimizeNativeCallbacks(source)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(result, "direct := worker("+replacement+",mark(5))") {
			t.Fatal("argument evaluation was moved")
		}
	}
}

func TestNativeCallbackWorkersRejectGetterEffects(t *testing.T) {
	for _, effect := range []string{
		"once.Do(func() { mark(90); cache =",
		"once.Do(func() { cache = mark(90); cache =",
	} {
		source := strings.Replace(callbackFixture, "once.Do(func() { cache =", effect, 1)
		result, err := OptimizeNativeCallbacks(source)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(result, "parser := r.Apply(getter(),") {
			t.Fatal("getter effects were bypassed")
		}
	}
}

func TestNativeCallbackWorkersTrackObjectsRatherThanNames(t *testing.T) {
	// A shadowed callback spelling is unrelated to the parameter under proof.
	source := strings.Replace(callbackFixture, "_ = alias", "{ alias := 1; _ = &alias }; _ = alias", 1)
	result, err := OptimizeNativeCallbacks(source)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(result, "func GopursNativeCallback_worker(") {
		t.Fatal("unrelated shadowing incorrectly rejected")
	}
	// Conversely, a local value named like the runtime import is never the
	// runtime intrinsic, even when its selector spelling happens to match.
	source = strings.Replace(callbackFixture, "var k r.Value", "r := otherRuntime; var k r.Value", 1)
	result, err = OptimizeNativeCallbacks(source)
	if err != nil {
		t.Fatal(err)
	}
	if result != source {
		t.Fatal("shadowed runtime intrinsic admitted")
	}
}

func TestNativeCallbackWorkersPreserveRecordABIThroughPackaging(t *testing.T) {
	// Source records retain their PureScript labels in producer annotations.
	// The callback pass runs before packaging exports these fields. Reprinting
	// its copied types without annotations used to produce inconsistent ABIs.
	const record = "struct { go__range int\n/*gopurs:record-v1:1:00720061006e00670065*/\n}"
	source := strings.ReplaceAll(callbackFixture, "struct { N int }", record)
	source = strings.ReplaceAll(source, ".N", ".go__range")
	source = strings.Replace(source, "package main", "package purescript", 1)
	source = strings.Replace(source, "func main()", "func Run()", 1)
	after, err := OptimizeNativeCallbacks(source)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(after, "func GopursNativeCallback_worker(") {
		t.Fatal("annotated callback did not specialize")
	}
	directory := packageEvidence(t)
	packagePut(t, directory, "go.mod", "module gopurs/output\n\ngo 1.27.0\n")
	packagePut(t, directory, "gopurs_runtime/runtime.go", callbackRuntime)
	request := PackageRequest{BindRequest: BindRequest{
		PackagePath: "gopurs/output/purescript", GoVersion: "go1.27", Arch: "arm64",
		Exports: packageExports(t, directory, "gopurs/output/gopurs_runtime", "sync", "fmt"),
		Files:   []SourceInput{{Owner: "generated/Callback", Module: "Callback", Path: "Callback.go", Content: after}},
	}, Entries: []SourceSymbol{{"generated/Callback", "Run"}}}
	result, err := packageSources(request, 1)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range result.Files {
		packagePut(t, directory, "purescript/"+file.Path, file.Content)
	}
	packagePut(t, directory, "main.go", fmt.Sprintf("package main\nimport p %q\nfunc main(){p.Run()}\n", result.Entries[0].Package))
	if got := string(packageRun(t, directory, "callback-record-ABI", "run", ".")); got != "0 12 13\n15 [12 13 5 15]\n" {
		t.Fatalf("packaged callback result: %q", got)
	}
}

func TestNativeCallbackWorkersDistinguishEqualGoShapesWithDifferentLabels(t *testing.T) {
	const record = "struct { N int\n/*gopurs:record-v1:1:004e*/\n}"
	source := strings.ReplaceAll(callbackFixture, "struct { N int }", record)
	// Only the direct callback's label differs. Its Go spelling alone matches
	// before packaging, but it is a different structural record ABI afterward.
	needle := "func(x " + record + ") r.Value { return mark(x.N+10) }"
	different := strings.Replace(needle, "004e", "006e", 1)
	source = strings.Replace(source, needle, different, 1)
	result, err := OptimizeNativeCallbacks(source)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(result, "direct := worker(") || !strings.Contains(result, different) {
		t.Fatal("different record label was conflated with native worker ABI")
	}
}
