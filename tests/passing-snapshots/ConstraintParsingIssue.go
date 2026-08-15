package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_X_dollar_Dict gopurs_runtime.Value
var once_Main_X_dollar_Dict sync.Once

func Get_Main_X_dollar_Dict() gopurs_runtime.Value {
	once_Main_X_dollar_Dict.Do(func() {
		cache_Main_X_dollar_Dict = gopurs_runtime.Func(func(x1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_X_dollar_Dict(x1_0_box)
		})
	})
	return cache_Main_X_dollar_Dict
}

var cache_Main_x gopurs_runtime.Value
var once_Main_x sync.Once

func Get_Main_x() gopurs_runtime.Value {
	once_Main_x.Do(func() {
		cache_Main_x = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_x(_dollar___unused_0_box)
		})
	})
	return cache_Main_x
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_X struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[1409933510] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_X)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_X: " + key)
		}
	}
}

func Call_Main_X_dollar_Dict(x1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x1_0 gopurs_runtime.Value = x1_0_loop
	_ = x1_0
	return x1_0
}

func Call_Main_x(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
}
