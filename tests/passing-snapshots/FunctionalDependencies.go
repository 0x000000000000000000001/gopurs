package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Proxy gopurs_runtime.Value
var once_Main_Proxy sync.Once

func Get_Main_Proxy() gopurs_runtime.Value {
	once_Main_Proxy.Do(func() {
		cache_Main_Proxy = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_Proxy
}

var cache_Main_Append_dollar_Dict gopurs_runtime.Value
var once_Main_Append_dollar_Dict sync.Once

func Get_Main_Append_dollar_Dict() gopurs_runtime.Value {
	once_Main_Append_dollar_Dict.Do(func() {
		cache_Main_Append_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Append_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Append_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_appendProxy gopurs_runtime.Value
var once_Main_appendProxy sync.Once

func Get_Main_appendProxy() gopurs_runtime.Value {
	once_Main_appendProxy.Do(func() {
		cache_Main_appendProxy = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_appendProxy(uint32(_dollar___unused_0_box.IntVal), uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_appendProxy
}

var cache_Main_appendProxy__494502715 gopurs_runtime.Value
var once_Main_appendProxy__494502715 sync.Once

func Get_Main_appendProxy__494502715() gopurs_runtime.Value {
	once_Main_appendProxy__494502715.Do(func() {
		cache_Main_appendProxy__494502715 = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_appendProxy__494502715(uint32(_dollar___unused_0_box.IntVal), uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_appendProxy__494502715
}

var cache_Main_appendNil gopurs_runtime.Value
var once_Main_appendNil sync.Once

func Get_Main_appendNil() gopurs_runtime.Value {
	once_Main_appendNil.Do(func() {
		cache_Main_appendNil = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_appendNil
}

var cache_Main_appendCons gopurs_runtime.Value
var once_Main_appendCons sync.Once

func Get_Main_appendCons() gopurs_runtime.Value {
	once_Main_appendCons.Do(func() {
		cache_Main_appendCons = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_appendCons(_dollar___unused_0_box)
		})
	})
	return cache_Main_appendCons
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test
}

type Constructor_Main_Proxy[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Append[T_a any, T_b any, T_c any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[1649103088] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Append[any, any, any])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Append: " + key)
		}
	}
}

func Call_Main_Append_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_appendProxy(_dollar___unused_0_loop uint32, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 227768594
}

func Call_Main_appendProxy__494502715(_dollar___unused_0_loop uint32, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 227768594
}

func Call_Main_appendCons(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordDict0().IntVal)), UnsafePtr: nil}
}
