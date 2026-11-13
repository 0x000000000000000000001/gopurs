// Copied beside real FFI sources and their generated bridges by
// tools/compiler-ffi-bridge.test.mjs.
package main

import (
	"reflect"
	"runtime"
	"testing"
	"unsafe"
	"weak"

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

func TestCompilerTypeMemo(t *testing.T) {
	memo := runEffect(_Gopurs_Gopurs_Cache_TypeMemo_Create)
	other := runEffect(_Gopurs_Gopurs_Cache_TypeMemo_Create)
	owner := &memoPayload{label: "same fields"}
	separate := &memoPayload{label: "same fields"}
	keys := []rt.Value{
		{Type: rt.TypeConstructor, IntVal: 17},
		{Type: rt.TypeConstructor, IntVal: 18},
		{Type: rt.TypeConstructor, IntVal: 17, UnsafePtr: unsafe.Pointer(owner)},
		{Type: rt.TypeConstructor, IntVal: 18, UnsafePtr: unsafe.Pointer(owner)},
		{Type: rt.TypeConstructor, IntVal: 17, UnsafePtr: unsafe.Pointer(separate)},
	}
	lookup := func(cache, key rt.Value) int64 {
		return runEffect(rt.Apply2(_Gopurs_Gopurs_Cache_TypeMemo_Lookup, cache, key)).IntVal
	}
	for i, key := range keys {
		insert := rt.Apply3(_Gopurs_Gopurs_Cache_TypeMemo_Insert, memo, key, rt.Int(int64(i)))
		if lookup(memo, key) != -1 {
			t.Fatal("memo insertion ran before its Effect thunk")
		}
		runEffect(insert)
	}
	for i, key := range keys {
		if lookup(memo, key) != int64(i) || lookup(other, key) != -1 {
			t.Fatal("memo key lost its constructor tag, object identity or cache ownership")
		}
	}
	runEffect(rt.Apply3(_Gopurs_Gopurs_Cache_TypeMemo_Insert, memo, keys[2], rt.Int(42)))
	if lookup(memo, keys[2]) != 42 || lookup(memo, keys[3]) != 3 {
		t.Fatal("replacing a memo entry changed another constructor")
	}
}

type memoPayload struct {
	label   string
	padding [64]byte
}

//go:noinline
func memoWithRetainedOwner() (rt.Value, weak.Pointer[memoPayload]) {
	owner := &memoPayload{label: "retained by memo key"}
	key := rt.Value{Type: rt.TypeConstructor, IntVal: 17, UnsafePtr: unsafe.Pointer(owner)}
	memo := runEffect(_Gopurs_Gopurs_Cache_TypeMemo_Create)
	runEffect(rt.Apply3(_Gopurs_Gopurs_Cache_TypeMemo_Insert, memo, key, rt.Int(42)))
	return memo, weak.Make(owner)
}

func TestCompilerTypeMemoRetainsOwner(t *testing.T) {
	memo, reference := memoWithRetainedOwner()
	for i := 0; i < 3; i++ {
		runtime.GC()
	}
	owner := reference.Value()
	if owner == nil || owner.label != "retained by memo key" {
		t.Fatal("memo key did not keep its payload alive")
	}
	key := rt.Value{Type: rt.TypeConstructor, IntVal: 17, UnsafePtr: unsafe.Pointer(owner)}
	if runEffect(rt.Apply2(_Gopurs_Gopurs_Cache_TypeMemo_Lookup, memo, key)).IntVal != 42 {
		t.Fatal("memo lookup changed after collection")
	}
	runtime.KeepAlive(memo)
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
