// Copied beside real FFI sources and their generated bridges by
// tools/compiler-ffi-bridge.test.mjs.
package main

import (
	"reflect"
	"runtime"
	"strings"
	"testing"

	cache "gopurs/output/gopurs_build_cache"
	rt "gopurs/output/gopurs_runtime"
)

func runEffect(effect rt.Value) rt.Value { return rt.Apply(effect, rt.Value{}) }

func stringArray(value rt.Value) []string {
	values := rt.Unbox[[]rt.Value](value)
	result := make([]string, len(values))
	for i, value := range values {
		result[i] = value.StrVal()
	}
	return result
}

func TestCompilerHandles(t *testing.T) {
	builder := runEffect(_Gopurs_Gopurs_Printer_Builder_NewBuilderImpl)
	other := runEffect(_Gopurs_Gopurs_Printer_Builder_NewBuilderImpl)
	pushed := rt.Apply2(_Gopurs_Gopurs_Printer_Builder_PushImpl, builder, rt.Str("é😀"))
	if rt.Unbox[*strings.Builder](pushed) != rt.Unbox[*strings.Builder](builder) {
		t.Fatal("push replaced the native builder handle")
	}
	completed := rt.Apply(_Gopurs_Gopurs_Printer_Builder_ToStringImpl, builder).StrVal()
	rt.Apply2(_Gopurs_Gopurs_Printer_Builder_PushImpl, builder, rt.Str(strings.Repeat("x", 8192)))
	builder, pushed = rt.Value{}, rt.Value{}
	runtime.GC()
	if completed != "é😀" || rt.Apply(_Gopurs_Gopurs_Printer_Builder_ToStringImpl, other).StrVal() != "" {
		t.Fatal("completed string or separate builder lost its private storage")
	}

	pool := runEffect(_Gopurs_Gopurs_Cache_Pool_Create)
	finish := rt.Apply(_Gopurs_Gopurs_Cache_Pool_Finish, pool)
	a := rt.Apply2(_Gopurs_Gopurs_Cache_Pool_Intern, pool, rt.Str("a"))
	b := rt.Apply2(_Gopurs_Gopurs_Cache_Pool_InternRef, pool, rt.Str("b"))
	if runEffect(finish).StrVal() != "" {
		t.Fatal("pool operations ran before their Effect thunks")
	}
	if runEffect(a).IntVal != 0 {
		t.Fatal("first pool index")
	}
	c := rt.Apply2(_Gopurs_Gopurs_Cache_Pool_Intern, pool, rt.Str("c"))
	if runEffect(c).IntVal != 1 || runEffect(b).StrVal() != `["int","2"]` || runEffect(b).StrVal() != `["int","2"]` {
		t.Fatal("numeric indices and references lost execution order or shared namespace")
	}
	if runEffect(finish).StrVal() != "a,c,b" || runEffect(a).IntVal != 0 {
		t.Fatal("pool deduplication changed")
	}
}

func TestCompilerImportArrays(t *testing.T) {
	for _, input := range [][][]string{nil, {{}}, {{}, {"math", ""}, {"math", "é😀"}, {}}} {
		outer := make([]rt.Value, len(input))
		want := []string{}
		for i, values := range input {
			inner := make([]rt.Value, len(values))
			for j, value := range values {
				inner[j] = rt.Str(value)
			}
			outer[i] = rt.Array(inner)
			want = append(want, values...)
		}
		boxed := rt.Array(outer)
		first := rt.Apply(_Gopurs_Gopurs_GoImports_ConcatStringArrays, boxed)
		if got := stringArray(first); !reflect.DeepEqual(got, want) {
			t.Fatalf("nested array: got %v, want %v", got, want)
		}
		for i := range rt.Unbox[[]rt.Value](first) {
			rt.Unbox[[]rt.Value](first)[i] = rt.Str("changed")
		}
		if got := stringArray(rt.Apply(_Gopurs_Gopurs_GoImports_ConcatStringArrays, boxed)); !reflect.DeepEqual(got, want) {
			t.Fatal("result mutation reached the borrowed input")
		}
	}
	fallback := rt.Func(func(rt.Value) rt.Value { panic("native scanner called PureScript fallback") })
	result := rt.Apply2(_Gopurs_Gopurs_GoCode_ReferencedImportsImpl, fallback, rt.Str("sync.Once{}; math.Abs(1)"))
	if got := stringArray(result); !reflect.DeepEqual(got, []string{"math", "sync"}) {
		t.Fatalf("native scanner result: %v", got)
	}
}

func TestCompilerHostEffects(t *testing.T) {
	cache.Events = nil
	onMiss := rt.Func(func(seed rt.Value) rt.Value {
		return rt.Func(func(rt.Value) rt.Value {
			cache.Events = append(cache.Events, "callback:"+seed.StrVal())
			return rt.Value{}
		})
	})
	miss := rt.Apply2(_Gopurs_Gopurs_BuildCache_SessionExchangeOnMiss, rt.Str("request"), onMiss)
	if len(cache.Events) != 0 {
		t.Fatal("cache request ran while constructing its effect")
	}
	if runEffect(miss).StrVal() != "request" || !reflect.DeepEqual(cache.Events, []string{"miss:request", "callback:seed", "returned"}) {
		t.Fatalf("cache callback ordering: %v", cache.Events)
	}
	run := rt.Func(func(rt.Value) rt.Value { cache.Events = append(cache.Events, "run"); return rt.Value{} })
	snapshot := rt.Apply(_Gopurs_Gopurs_Host_Snapshot_WithSnapshot, run)
	before := len(cache.Events)
	runEffect(snapshot)
	runEffect(_Gopurs_Gopurs_Host_Snapshot_Release)
	if !reflect.DeepEqual(cache.Events[before:], []string{"freeze", "run", "release"}) {
		t.Fatalf("snapshot callback ordering: %v", cache.Events)
	}

	previous := runtime.MemProfileRate
	defer func() { runtime.MemProfileRate = previous }()
	setRate := rt.Apply(_Gopurs_Gopurs_Host_Metrics_SetMemProfileRate, rt.Int(12345))
	if runtime.MemProfileRate != previous {
		t.Fatal("profile rate changed while constructing its effect")
	}
	runEffect(setRate)
	if runtime.MemProfileRate != 12345 {
		t.Fatal("profile rate effect did not run")
	}
	first := runEffect(_Gopurs_Gopurs_Host_Metrics_Now).FloatVal()
	if next := runEffect(_Gopurs_Gopurs_Host_Metrics_Now).FloatVal(); next < first {
		t.Fatal("monotonic timer moved backwards")
	}
}
