// Copied alongside the native FFI sources by tools/native-ffi.test.mjs.
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestEscaping(t *testing.T) {
	raw, err := os.ReadFile("cases.json")
	if err != nil {
		t.Fatal(err)
	}
	var cases []struct{ Input, Expected string }
	if err := json.Unmarshal(raw, &cases); err != nil {
		t.Fatal(err)
	}
	for _, item := range cases {
		raw, err := hex.DecodeString(item.Input)
		if err != nil {
			t.Fatal(err)
		}
		if got := EscapeGoStringImpl(string(raw)); got != item.Expected {
			t.Errorf("%x: got %q, want %q", raw, got, item.Expected)
		}
	}
}

func TestMemoization(t *testing.T) {
	calls := 0
	memo := MemoizeName(func(value string) string {
		calls++
		return strings.ToUpper(value)
	})
	var workers sync.WaitGroup
	for i := 0; i < 16; i++ {
		workers.Add(1)
		go func() {
			defer workers.Done()
			for i := 0; i < 32; i++ {
				if got := memo("abc"); got != "ABC" {
					t.Errorf("got %q", got)
				}
				if got := memo(""); got != "" {
					t.Errorf("got %q", got)
				}
			}
		}()
	}
	workers.Wait()
	if calls != 2 {
		t.Fatalf("callback calls: %d", calls)
	}
	second := MemoizeName(func(value string) string { return "other:" + value })
	if got := second("abc"); got != "other:abc" {
		t.Fatalf("caches shared: %q", got)
	}
}

func TestRuntimeBytes(t *testing.T) {
	source, err := os.ReadFile("runtime.txt")
	if err != nil {
		t.Fatal(err)
	}
	if RuntimeGoCode != string(source) {
		t.Fatal("native runtime differs from canonical source")
	}
}

func TestNativeBridge(t *testing.T) {
	source := "func Identity(value string) string { return value }"
	if got := ExtractFfiAstImpl(source, nil); !json.Valid([]byte(got)) || !strings.Contains(got, `"Identity"`) {
		t.Fatalf("invalid declarations: %s", got)
	}
	got := PrepareFfiAstImpl("Fixture_", source, nil)
	if !json.Valid([]byte(got)) || !strings.Contains(got, "Fixture_Identity") {
		t.Fatalf("invalid module: %s", got)
	}
	defer func() {
		failure := recover()
		if failure == nil || !strings.Contains(fmt.Sprint(failure), "Go FFI parse failed") {
			t.Fatalf("missing parse error: %v", failure)
		}
	}()
	ExtractFfiAstImpl("func Broken(", nil)
}
