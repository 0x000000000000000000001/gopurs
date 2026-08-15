package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_FD_dollar_Dict gopurs_runtime.Value
var once_Main_FD_dollar_Dict sync.Once

func Get_Main_FD_dollar_Dict() gopurs_runtime.Value {
	once_Main_FD_dollar_Dict.Do(func() {
		cache_Main_FD_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_FD_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_FD_dollar_Dict
}

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

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_fn1 gopurs_runtime.Value
var once_Main_fn1 sync.Once

func Get_Main_fn1() gopurs_runtime.Value {
	once_Main_fn1.Do(func() {
		cache_Main_fn1 = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_fn1(uint32(_dollar___unused_0_box.IntVal), uint32(_dollar___unused_1_box.IntVal), v_2_box))
		})
	})
	return cache_Main_fn1
}

var cache_Main_fn2 gopurs_runtime.Value
var once_Main_fn2 sync.Once

func Get_Main_fn2() gopurs_runtime.Value {
	once_Main_fn2.Do(func() {
		cache_Main_fn2 = gopurs_runtime.Func3(func(dictFD_0_box gopurs_runtime.Value, dictC_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fn2(dictFD_0_box, dictC_1_box, x_2_box)
		})
	})
	return cache_Main_fn2
}

var cache_Main_fn1__913275691 gopurs_runtime.Value
var once_Main_fn1__913275691 sync.Once

func Get_Main_fn1__913275691() gopurs_runtime.Value {
	once_Main_fn1__913275691.Do(func() {
		cache_Main_fn1__913275691 = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_fn1__913275691(uint32(_dollar___unused_0_box.IntVal), uint32(_dollar___unused_1_box.IntVal), v_2_box))
		})
	})
	return cache_Main_fn1__913275691
}

type Constructor_Main_FD struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3330339132] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_FD)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_FD: " + key)
		}
	}
}

type Constructor_Main_C struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[2167983901] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_C)(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_C: " + key)
		}
	}
}

func Call_Main_FD_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_C_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_fn1(_dollar___unused_0_loop uint32, _dollar___unused_1_loop uint32, v_2_loop gopurs_runtime.Value) string {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 uint32 = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return ""
}

func Call_Main_fn2(dictFD_0_loop gopurs_runtime.Value, dictC_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFD_0 gopurs_runtime.Value = dictFD_0_loop
	_ = dictFD_0
	var dictC_1 gopurs_runtime.Value = dictC_1_loop
	_ = dictC_1
	var x_2 gopurs_runtime.Value = x_2_loop
	_ = x_2
	return gopurs_runtime.Str("")
}

func Call_Main_fn1__913275691(_dollar___unused_0_loop uint32, _dollar___unused_1_loop uint32, v_2_loop gopurs_runtime.Value) string {
	var _dollar___unused_0 uint32 = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 uint32 = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return ""
}
