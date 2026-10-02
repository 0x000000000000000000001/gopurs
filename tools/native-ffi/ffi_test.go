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

	"gopurs/output/gopurs_runtime"
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
			created := make(chan gopurs_runtime.Value, 1)
			var creator sync.WaitGroup
			creator.Add(1)
			prefix := fmt.Sprintf("%d:é😀\xed\xa0\x80\x00", seed)
			go func() {
				defer creator.Done()
				builder := NewBuilderImpl(gopurs_runtime.Value{})
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
			builder = gopurs_runtime.Value{}
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
	fallback := gopurs_runtime.Func(func(gopurs_runtime.Value) gopurs_runtime.Value {
		panic("native scanner invoked the PureScript fallback")
	})
	for _, item := range cases {
		text, err := hex.DecodeString(item.Input)
		if err != nil {
			t.Fatal(err)
		}
		result := ReferencedImportsImpl(fallback, gopurs_runtime.Str(string(text)))
		actual := make([]string, 0)
		for _, value := range gopurs_runtime.Unbox[[]gopurs_runtime.Value](result) {
			actual = append(actual, value.StrVal())
		}
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
		outer := make([]gopurs_runtime.Value, len(item.Input))
		for i, inner := range item.Input {
			values := make([]gopurs_runtime.Value, len(inner))
			for j, value := range inner {
				values[j] = gopurs_runtime.Str(value)
			}
			outer[i] = gopurs_runtime.Array(values)
		}
		input := gopurs_runtime.Array(outer)
		first := gopurs_runtime.Unbox[[]gopurs_runtime.Value](ConcatStringArrays(input))
		second := ConcatStringArrays(input)
		actual := make([]string, 0, len(first))
		for _, value := range first {
			actual = append(actual, value.StrVal())
		}
		if !reflect.DeepEqual(actual, item.Expected) {
			t.Fatalf("got %v, want %v", actual, item.Expected)
		}
		for i := range first {
			first[i] = gopurs_runtime.Str("changed")
		}
		runtime.GC()
		for _, values := range []gopurs_runtime.Value{second, ConcatStringArrays(input)} {
			for i, value := range gopurs_runtime.Unbox[[]gopurs_runtime.Value](values) {
				if value.StrVal() != item.Expected[i] {
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
	SetMemProfileRate(12345, gopurs_runtime.Value{})
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
