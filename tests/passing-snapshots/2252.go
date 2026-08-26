package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_T gopurs_runtime.Value
var once_Main_T sync.Once

func Get_Main_T() gopurs_runtime.Value {
	once_Main_T.Do(func() {
		cache_Main_T = gopurs_runtime.Value{Type: 9, IntVal: int64(990467018), UnsafePtr: nil}
	})
	return cache_Main_T
}

var cache_Main_ti gopurs_runtime.Value
var once_Main_ti sync.Once

func Get_Main_ti() gopurs_runtime.Value {
	once_Main_ti.Do(func() {
		cache_Main_ti = gopurs_runtime.Value{Type: 9, IntVal: int64(990467018), UnsafePtr: nil}
	})
	return cache_Main_ti
}

var cache_Main_t gopurs_runtime.Value
var once_Main_t sync.Once

func Get_Main_t() gopurs_runtime.Value {
	once_Main_t.Do(func() {
		cache_Main_t = gopurs_runtime.Value{Type: 9, IntVal: int64(990467018), UnsafePtr: nil}
	})
	return cache_Main_t
}

var cache_Main_xs gopurs_runtime.Value
var once_Main_xs sync.Once

func Get_Main_xs() gopurs_runtime.Value {
	once_Main_xs.Do(func() {
		cache_Main_xs = func() gopurs_runtime.Value {
			arr := []uint32{990467018, 990467018, 990467018}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: int64(v), UnsafePtr: nil}
			}
			return gopurs_runtime.Array(boxed)
		}()
	})
	return cache_Main_xs
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_T[T_a any] struct {
	Rc uint32
}
