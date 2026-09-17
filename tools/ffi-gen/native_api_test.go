package main

import (
	"encoding/json"
	"go/parser"
	"go/token"
	"strings"
	"sync"
	"testing"
)

func TestPreparedSourceRetainsPackageWordsOutsideItsClause(t *testing.T) {
	const declarations = "const Text = \"package literal\"\n" +
		"var Multiline = `first\npackage embedded\nlast`\n" +
		"// package trailing comment\nfunc Read() string { return Text + Multiline }\n"
	for _, clause := range []string{"", "package ffi\n", "\tpackage\tffi;", "package /* middle comment */ ffi /* trailing comment */ ;"} {
		source := "// package misleading comment\n/* package another comment */\n" + clause + declarations
		prepared, err := Prepare("Fixture_", source)
		if err != nil {
			t.Fatalf("clause %q: %v", clause, err)
		}
		var result struct{ Content string }
		if err := json.Unmarshal([]byte(prepared), &result); err != nil {
			t.Fatal(err)
		}
		for _, preserved := range []string{"\"package literal\"", "`first\npackage embedded\nlast`", "// package misleading comment", "/* package another comment */", "// package trailing comment"} {
			if !strings.Contains(result.Content, preserved) {
				t.Fatalf("clause %q lost %q: %s", clause, preserved, result.Content)
			}
		}
		if _, err := parser.ParseFile(token.NewFileSet(), "prepared.go", "package purescript\n"+result.Content, 0); err != nil {
			t.Fatalf("prepared source cannot be assembled: %v\n%s", err, result.Content)
		}
	}
}

func TestNativeParserMatchesContract(t *testing.T) {
	const source = "package ffi\nvar Value = 3\nfunc Apply(s string) string { return s }\n"
	wantExtract, err := parseFFI(source)
	if err != nil {
		t.Fatal(err)
	}
	wantPrepare, err := parseFFI(source, "Fixture_")
	if err != nil {
		t.Fatal(err)
	}
	var workers sync.WaitGroup
	for i := 0; i < 16; i++ {
		workers.Add(1)
		go func() {
			defer workers.Done()
			for i := 0; i < 4; i++ {
				if got, err := Extract(source); err != nil || got != wantExtract {
					t.Errorf("Extract = %q, %v; want %q", got, err, wantExtract)
				}
				if got, err := Prepare("Fixture_", source); err != nil || got != wantPrepare {
					t.Errorf("Prepare = %q, %v; want %q", got, err, wantPrepare)
				}
			}
		}()
	}
	workers.Wait()
	for _, parse := range []func(string) (string, error){Extract, func(source string) (string, error) {
		return Prepare("Fixture_", source)
	}} {
		got, err := parse("func Broken(")
		if got != "" || err == nil || !strings.Contains(err.Error(), "1:13") {
			t.Fatalf("invalid input returned %q, %v", got, err)
		}
	}
}
