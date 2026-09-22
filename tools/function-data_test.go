package gopurs_runtime_test

import (
	"fmt"
	rt "gopurs/output/gopurs_runtime"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

type functionDataMarker struct {
	seed    int64
	child   *int64
	padding [128]int64
}
type distinctFunctionDataMarker functionDataMarker

func TestFunctionDataArityMatrix(t *testing.T) {
	for arity := 1; arity <= 11; arity++ {
		for supplied := 1; supplied <= 11; supplied++ {
			t.Run(fmt.Sprintf("arity_%d_apply_%d", arity, supplied), func(t *testing.T) {
				total := max(arity, supplied)
				args := make([]rt.Value, total)
				want := make([]int64, total)
				for i := range args {
					args[i] = rt.Int(int64(i + 1))
					want[i] = int64(i + 1)
				}
				var seen []int64
				calls := 0
				var next func([]rt.Value) rt.Value
				next = func(values []rt.Value) rt.Value {
					calls++
					for _, v := range values {
						seen = append(seen, v.IntVal)
					}
					if len(seen) == total {
						return rt.Int(42)
					}
					return rt.Func(func(v rt.Value) rt.Value { return next([]rt.Value{v}) })
				}
				fn := rt.WithFunctionData(arityCallback(arity, next), "full-call")
				result := applyArguments(fn, args[:supplied])
				if supplied < arity {
					if calls != 0 {
						t.Fatalf("partial invoked %d times", calls)
					}
					if _, ok := rt.FunctionData[string](result); ok {
						t.Fatal("partial retained full-call metadata")
					}
					result = applyArguments(result, args[supplied:])
				}
				if !reflect.DeepEqual(seen, want) {
					t.Fatalf("argument order %v, want %v", seen, want)
				}
				if calls != 1+max(0, supplied-arity) {
					t.Fatalf("unexpected callback count %d", calls)
				}
				if result.Type != rt.TypeInt || result.IntVal != 42 {
					t.Fatalf("unexpected result %+v", result)
				}
				if value, ok := rt.FunctionData[string](fn); !ok || value != "full-call" {
					t.Fatal("call mutated source metadata")
				}
			})
		}
	}
}

func contextualUncurried(fn rt.Value, args []rt.Value) rt.Value {
	switch len(args) {
	case 2:
		return rt.UncurriedApp2(fn, args[0], args[1])
	case 3:
		return rt.UncurriedApp3(fn, args[0], args[1], args[2])
	case 4:
		return rt.UncurriedApp4(fn, args[0], args[1], args[2], args[3])
	case 5:
		return rt.UncurriedApp5(fn, args[0], args[1], args[2], args[3], args[4])
	case 6:
		return rt.UncurriedApp6(fn, args[0], args[1], args[2], args[3], args[4], args[5])
	case 7:
		return rt.UncurriedApp7(fn, args[0], args[1], args[2], args[3], args[4], args[5], args[6])
	case 8:
		return rt.UncurriedApp8(fn, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7])
	case 9:
		return rt.UncurriedApp9(fn, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8])
	case 10:
		return rt.UncurriedApp10(fn, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9])
	default:
		return rt.UncurriedApp(fn, args...)
	}
}
func TestFunctionDataUncurried(t *testing.T) {
	for arity := 1; arity <= 11; arity++ {
		t.Run(fmt.Sprint(arity), func(t *testing.T) {
			args := make([]rt.Value, arity)
			for i := range args {
				args[i] = rt.Int(int64(i + 1))
			}
			fn := rt.WithFunctionData(arityCallback(arity, func(values []rt.Value) rt.Value {
				sum := int64(0)
				for i, v := range values {
					sum += int64(i+1) * v.IntVal
				}
				return rt.Int(sum)
			}), arity)
			want := int64(arity * (arity + 1) * (2*arity + 1) / 6)
			for _, result := range []rt.Value{contextualUncurried(fn, args), rt.UncurriedApp(fn, args...)} {
				if result.IntVal != want {
					t.Fatalf("got %d, want %d", result.IntVal, want)
				}
			}
		})
	}
	unit := rt.WithFunctionData(rt.Wrap0(func() int64 { return 71 }), "unit")
	if rt.UncurriedApp(unit).IntVal != 71 {
		t.Fatal("zero-argument wrapper failed")
	}
}
func TestFunctionDataTypesAndReplacement(t *testing.T) {
	original := rt.Func(func(v rt.Value) rt.Value { return v })
	marker := &functionDataMarker{seed: 12}
	fn := rt.WithFunctionData(original, marker)
	if got, ok := rt.FunctionData[*functionDataMarker](fn); !ok || got != marker {
		t.Fatal("metadata identity lost")
	}
	if _, ok := rt.FunctionData[*distinctFunctionDataMarker](fn); ok {
		t.Fatal("accepted distinct metadata type")
	}
	if _, ok := rt.FunctionData[string](original); ok {
		t.Fatal("ordinary function has metadata")
	}
	if _, ok := rt.FunctionData[string](rt.Int(7)); ok {
		t.Fatal("scalar has metadata")
	}
	replaced := rt.WithFunctionData(fn, "replacement")
	if got, ok := rt.FunctionData[string](replaced); !ok || got != "replacement" {
		t.Fatal("metadata replacement failed")
	}
	if _, ok := rt.FunctionData[*functionDataMarker](replaced); ok {
		t.Fatal("old metadata still visible")
	}
	if got, ok := rt.FunctionData[*functionDataMarker](fn); !ok || got != marker {
		t.Fatal("replacement mutated original")
	}
	if rt.Apply(replaced, rt.Int(31)).IntVal != 31 {
		t.Fatal("replacement changed callback")
	}
	var nilMarker *functionDataMarker
	if got, ok := rt.FunctionData[*functionDataMarker](rt.WithFunctionData(original, nilMarker)); !ok || got != nil {
		t.Fatal("typed nil metadata lost")
	}
	if _, ok := rt.FunctionData[*functionDataMarker](rt.WithFunctionData(original, nil)); ok {
		t.Fatal("untyped nil falsely matched")
	}
}
func TestFunctionDataBoxing(t *testing.T) {
	fn := rt.WithFunctionData(rt.Func(func(v rt.Value) rt.Value { return rt.Int(v.IntVal + 3) }), "boxed")
	for _, roundtrip := range []rt.Value{rt.Box(fn), rt.Unbox[rt.Value](fn), rt.Box(fn.AnyVal())} {
		if got, ok := rt.FunctionData[string](roundtrip); !ok || got != "boxed" {
			t.Fatal("Value roundtrip lost metadata")
		}
		if rt.Apply(roundtrip, rt.Int(8)).IntVal != 11 {
			t.Fatal("Value roundtrip changed callback")
		}
	}
	native := rt.Unbox[func(any) any](fn)
	if got := rt.Unbox[int64](native(int64(9))); got != 12 {
		t.Fatalf("unboxed function returned %d", got)
	}
	if got := rt.Apply(rt.Box(native), rt.Int(10)).IntVal; got != 13 {
		t.Fatalf("function reboxing returned %d", got)
	}
}
func TestFunctionDataLifetime(t *testing.T) {
	var metadataFinalized []*atomic.Bool
	values := afterCreatorReturns(func() []closureCase {
		var out []closureCase
		for arity := 1; arity <= 11; arity++ {
			c := newClosureCase(arity, 0, int64(100+arity))
			marker := &functionDataMarker{seed: int64(arity)}
			child := int64(arity + 20)
			marker.child = &child
			marker.padding[127] = int64(arity + 30)
			finalized := new(atomic.Bool)
			metadataFinalized = append(metadataFinalized, finalized)
			runtime.SetFinalizer(marker, func(*functionDataMarker) { finalized.Store(true) })
			c.value = rt.WithFunctionData(c.value, marker)
			out = append(out, c)
		}
		return out
	})
	for i, c := range values {
		if err := verifyClosure(c); err != nil {
			t.Fatal(err)
		}
		marker, ok := rt.FunctionData[*functionDataMarker](c.value)
		if !ok || marker.seed != int64(c.arity) || *marker.child != int64(c.arity+20) || marker.padding[127] != int64(c.arity+30) {
			t.Fatal("metadata capture corrupted")
		}
		if metadataFinalized[i].Load() {
			t.Fatal("reachable metadata was finalized")
		}
		runtime.KeepAlive(c.value)
	}
}
func TestFunctionDataConcurrent(t *testing.T) {
	marker := &functionDataMarker{seed: 97}
	fn := rt.WithFunctionData(rt.Func2(func(a, b rt.Value) rt.Value { return rt.Int(a.IntVal*10 + b.IntVal) }), marker)
	var group sync.WaitGroup
	failures := make(chan string, 16)
	for worker := 0; worker < 8; worker++ {
		group.Add(1)
		go func(worker int) {
			defer group.Done()
			for i := 0; i < 100; i++ {
				if got := rt.Apply2(fn, rt.Int(int64(worker)), rt.Int(int64(i))).IntVal; got != int64(worker*10+i) {
					failures <- "wrong concurrent result"
					return
				}
				if data, ok := rt.FunctionData[*functionDataMarker](fn); !ok || data != marker || data.seed != 97 {
					failures <- "metadata changed"
					return
				}
				local := rt.WithFunctionData(fn, worker)
				if v, ok := rt.FunctionData[int](local); !ok || v != worker {
					failures <- "replacement race"
					return
				}
			}
		}(worker)
	}
	group.Add(1)
	go func() {
		defer group.Done()
		for i := 0; i < 4; i++ {
			runtime.GC()
		}
	}()
	group.Wait()
	close(failures)
	for err := range failures {
		t.Error(err)
	}
	runtime.KeepAlive(fn)
}

var contextualAllocationSink rt.Value

func TestFunctionDataSaturatedAllocations(t *testing.T) {
	functions := []rt.Value{
		rt.Func(func(a rt.Value) rt.Value { return a }),
		rt.Func2(func(a, b rt.Value) rt.Value { return b }),
		rt.Func3(func(a, b, c rt.Value) rt.Value { return c }),
		rt.Func4(func(a, b, c, d rt.Value) rt.Value { return d }),
		rt.Func5(func(a, b, c, d, e rt.Value) rt.Value { return e }),
		rt.Func6(func(a, b, c, d, e, f rt.Value) rt.Value { return f }),
		rt.Func7(func(a, b, c, d, e, f, g rt.Value) rt.Value { return g }),
		rt.Func8(func(a, b, c, d, e, f, g, h rt.Value) rt.Value { return h }),
		rt.Func9(func(a, b, c, d, e, f, g, h, i rt.Value) rt.Value { return i }),
		rt.Func10(func(a, b, c, d, e, f, g, h, i, j rt.Value) rt.Value { return j }),
	}
	args := make([]rt.Value, 10)
	for i := range args {
		args[i] = rt.Int(int64(i + 1))
	}
	for i, ordinary := range functions {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			fn := rt.WithFunctionData(ordinary, "allocation-marker")
			allocations := testing.AllocsPerRun(1000, func() { contextualAllocationSink = applyArguments(fn, args[:i+1]) })
			if allocations != 0 {
				t.Fatalf("saturated contextual call allocated %v times", allocations)
			}
			if contextualAllocationSink.IntVal != int64(i+1) {
				t.Fatal("wrong saturated result")
			}
		})
	}
	// Arity 11 has no saturated Apply11 helper; Apply10 + Apply necessarily creates a partial.
}
func TestFunctionDataInvalidFunction(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("non-function metadata attachment did not panic")
		}
	}()
	rt.WithFunctionData(rt.Int(1), "invalid")
}
