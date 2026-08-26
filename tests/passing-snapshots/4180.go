package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_C_dollar_Dict gopurs_runtime.Value
var once_Main_C_dollar_Dict sync.Once

func Get_Main_C_dollar_Dict() gopurs_runtime.Value {
	once_Main_C_dollar_Dict.Do(func() {
		cache_Main_C_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_C_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_C_dollar_Dict
}

var cache_Main_c gopurs_runtime.Value
var once_Main_c sync.Once

func Get_Main_c() gopurs_runtime.Value {
	once_Main_c.Do(func() {
		cache_Main_c = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(func() gopurs_runtime.Value {
			orig := func() *struct {
			} {
				orig := gopurs_runtime.RecordDict0()
				_ = orig
				clone := struct {
				}{}

				return &clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict([]string{}, []gopurs_runtime.Value{})
		}().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_c
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f(_dollar___unused_0_box)
		})
	})
	return cache_Main_f
}

var cache_Main_v gopurs_runtime.Value
var once_Main_v sync.Once

func Get_Main_v() gopurs_runtime.Value {
	once_Main_v.Do(func() {
		cache_Main_v = gopurs_runtime.Int(0)
	})
	return cache_Main_v
}

type Constructor_Main_C[T_t any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[2167983901] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_C[any])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_C: " + key)
		}
	}
}

func Call_Main_C_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_f(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Int(0)
}
