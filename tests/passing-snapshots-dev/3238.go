package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
			return gopurs_runtime.Str(Call_Main_fn1(_dollar___unused_0_box, _dollar___unused_1_box, v_2_box))
		})
	})
	return cache_Main_fn1
}

var cache_Main_fn2 gopurs_runtime.Value
var once_Main_fn2 sync.Once

func Get_Main_fn2() gopurs_runtime.Value {
	once_Main_fn2.Do(func() {
		cache_Main_fn2 = gopurs_runtime.Func3(func(dictFD_0_box gopurs_runtime.Value, dictC_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_fn2(dictFD_0_box, dictC_1_box, x_2_box))
		})
	})
	return cache_Main_fn2
}

type Constructor_Main_FD[T_a any, T_b any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00460044
func init() {
	gopurs_runtime.StructGetters[3330339132] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_FD[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_FD: " + key)
		}
	}
}

type Constructor_Main_C[T_a any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0043
func init() {
	gopurs_runtime.StructGetters[2167983901] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_C[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_C: " + key)
		}
	}
}

func Call_Main_fn1(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) string {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return ""
}

func Call_Main_fn2(dictFD_0_loop gopurs_runtime.Value, dictC_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) string {
	var dictFD_0 gopurs_runtime.Value = dictFD_0_loop
	_ = dictFD_0
	var dictC_1 gopurs_runtime.Value = dictC_1_loop
	_ = dictC_1
	var x_2 gopurs_runtime.Value = x_2_loop
	_ = x_2
	return Call_Main_fn1(dictFD_0, dictC_1, x_2)
}
