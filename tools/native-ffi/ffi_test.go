// Copied alongside the native FFI sources by tools/native-ffi.test.mjs.
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
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

// Each worker transfers its private builder after the creator goroutine exits.
// The published string must also outlive its builder and any later appends.
func TestBuilderLifetime(t *testing.T) {
	var workers sync.WaitGroup
	for i := 0; i < 8; i++ {
		workers.Add(1)
		go func(seed int) {
			defer workers.Done()
			created := make(chan *strings.Builder, 1)
			var creator sync.WaitGroup
			creator.Add(1)
			prefix := fmt.Sprintf("%d:é😀\xed\xa0\x80\x00", seed)
			go func() {
				defer creator.Done()
				builder := NewBuilderImpl()
				if ToStringImpl(builder) != "" {
					t.Error("new builder is not empty")
				}
				if PushImpl(builder, prefix) != builder {
					t.Error("push replaced the handle")
				}
				created <- PushImpl(builder, "")
			}()
			builder := <-created
			creator.Wait()
			runtime.GC()
			completed := ToStringImpl(builder)
			PushImpl(builder, strings.Repeat("x", 8192))
			if completed != prefix || ToStringImpl(builder) != prefix+strings.Repeat("x", 8192) {
				t.Error("builder contents or completed string changed")
			}
			builder = nil
			runtime.GC()
			if completed != prefix {
				t.Error("completed string did not retain its storage")
			}
		}(i)
	}
	workers.Wait()
}

func TestImportScanner(t *testing.T) {
	raw, err := os.ReadFile("imports.json")
	if err != nil {
		t.Fatal(err)
	}
	var cases []struct {
		Input    string
		Expected []string
	}
	if err := json.Unmarshal(raw, &cases); err != nil {
		t.Fatal(err)
	}
	fallback := func(string) []string {
		panic("native scanner invoked the PureScript fallback")
	}
	for _, item := range cases {
		text, err := hex.DecodeString(item.Input)
		if err != nil {
			t.Fatal(err)
		}
		actual := append([]string{}, ReferencedImportsImpl(fallback, string(text))...)
		if !reflect.DeepEqual(actual, item.Expected) {
			t.Errorf("%q: got %v, want %v", text, actual, item.Expected)
		}
	}
}

func TestImportArrayOwnership(t *testing.T) {
	raw, err := os.ReadFile("arrays.json")
	if err != nil {
		t.Fatal(err)
	}
	var cases []struct {
		Input    [][]string
		Expected []string
	}
	if err := json.Unmarshal(raw, &cases); err != nil {
		t.Fatal(err)
	}
	for _, item := range cases {
		first := ConcatStringArrays(item.Input)
		second := ConcatStringArrays(item.Input)
		actual := append([]string{}, first...)
		if !reflect.DeepEqual(actual, item.Expected) {
			t.Fatalf("got %v, want %v", actual, item.Expected)
		}
		for i := range first {
			first[i] = "changed"
		}
		runtime.GC()
		for _, values := range [][]string{second, ConcatStringArrays(item.Input)} {
			for i, value := range values {
				if value != item.Expected[i] {
					t.Fatal("output shares mutable storage with another output or input")
				}
			}
		}
	}
}

func TestMetrics(t *testing.T) {
	previous := Now()
	for i := 0; i < 100; i++ {
		current := Now()
		if current < previous {
			t.Fatal("clock moved backwards")
		}
		previous = current
	}
	rate := runtime.MemProfileRate
	defer func() { runtime.MemProfileRate = rate }()
	SetMemProfileRate(12345, nil)
	if runtime.MemProfileRate != 12345 {
		t.Fatal("profile rate was not applied")
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
