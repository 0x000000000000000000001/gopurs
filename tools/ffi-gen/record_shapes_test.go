package main

import (
	"encoding/hex"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestShareRecordShapesExecution(t *testing.T) {
	source := `package main
import ("fmt"; r "gopurs/output/gopurs_runtime")
var trace string
func value(n int64) r.Value { trace += fmt.Sprint(n); return r.Int(n) }
var first = r.RecordDict2("z", "a", value(1), value(2))
func main() {
 second := r.RecordDict2("z", "\x61", value(3), value(4))
 nested := r.RecordDict1("child", r.RecordDict2("z", "a", value(5), value(6)))
 keys1, _, _ := r.RecordFields2(first)
 keys2, _, _ := r.RecordFields2(second)
 child := r.RecordGet(nested, "child")
 fmt.Println(trace, keys1 == keys2, r.RecordAt2(first, 0).IntVal, r.RecordGet(child, "a").IntVal)
}
`
	output, err := ShareRecordShapes("Test/main.go", source)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(output, "var GopursRecordKeys_") != 2 {
		t.Fatalf("expected two schemas: %s", output)
	}
	for _, values := range []string{"value(1), value(2)", "value(3), value(4)", "value(5), value(6)"} {
		if !strings.Contains(output, values) {
			t.Fatalf("changed argument expression: %s", output)
		}
	}
	again, err := ShareRecordShapes("Test/main.go", output)
	if err != nil || again != output {
		t.Fatalf("pass is not idempotent: %v", err)
	}
	directory := t.TempDir()
	packagePut(t, directory, "go.mod", "module gopurs/output\n\ngo 1.22\n")
	runtime, err := os.ReadFile("../../runtime/runtime.go")
	if err != nil {
		t.Fatal(err)
	}
	packagePut(t, directory, "gopurs_runtime/runtime.go", string(runtime))
	packagePut(t, directory, "main.go", output)
	command := exec.Command("go", "run", ".")
	command.Dir = directory
	command.Env = append(os.Environ(), "GOWORK=off", "GOTOOLCHAIN=local")
	result, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("execute shapes: %v\n%s", err, result)
	}
	if string(result) != "123456 true 1 6\n" {
		t.Fatalf("unexpected evaluation/result: %s", result)
	}
}

func TestShareRecordShapesBoundaries(t *testing.T) {
	name := "GopursRecordKeys_" + hex.EncodeToString([]byte("scope")) + "_0"
	source := `package p
import r "gopurs/output/gopurs_runtime"
func valid() {
 ` + name + ` := 1; _ = ` + name + `
 _ = r.RecordDict2("\xed\xa0\x80", "a\x00:b", f(), g())
 _ = r.RecordDict2("a", "a", f(), g())
 _ = r.RecordDict2("a", "b", f(), g())
 _ = r.RecordDict2("b", "a", f(), g())
 _ = r.RecordDict1(key(), f())
 _ = r.RecordDict2("a", "b", args...)
}
func shadow(r SomeType) { _ = r.RecordDict1("a", f()) }
`
	output, err := ShareRecordShapes("scope", source)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(output, "var GopursRecordKeys_") != 4 {
		t.Fatalf("ordered/duplicate schemas were conflated: %s", output)
	}
	if strings.Contains(output, "&"+name+",") {
		t.Fatalf("introduced local shadowing: %s", output)
	}
	for _, preserved := range []string{`"\xed\xa0\x80"`, `"a\x00:b"`, `r.RecordDict1(key(), f())`, `r.RecordDict2("a", "b", args...)`, `r.RecordDict1("a", f())`} {
		if !strings.Contains(output, preserved) {
			t.Fatalf("missing unchanged bytes %s in %s", preserved, output)
		}
	}
	for _, source := range []string{
		`package p; import r "another/runtime"; var x = r.RecordDict1("a", f())`,
		`package p; import . "gopurs/output/gopurs_runtime"; var x = RecordDict1("a", f())`,
	} {
		output, err := ShareRecordShapes("other", source)
		if err != nil || output != source {
			t.Fatalf("changed unrelated source: %v %s", err, output)
		}
	}
	if _, err := ShareRecordShapes("bad", "not Go"); err == nil {
		t.Fatal("invalid source accepted")
	}
}
