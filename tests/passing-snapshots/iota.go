package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_s gopurs_runtime.Value
var once_Main_s sync.Once

func Get_Main_s() gopurs_runtime.Value {
	once_Main_s.Do(func() {
		cache_Main_s = gopurs_runtime.Func3(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_s(x_0_box, y_1_box, z_2_box)
		})
	})
	return cache_Main_s
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_k gopurs_runtime.Value
var once_Main_k sync.Once

func Get_Main_k() gopurs_runtime.Value {
	once_Main_k.Do(func() {
		cache_Main_k = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_k(x_0_box, y_1_box)
		})
	})
	return cache_Main_k
}

var cache_Main_iota gopurs_runtime.Value
var once_Main_iota sync.Once

func Get_Main_iota() gopurs_runtime.Value {
	once_Main_iota.Do(func() {
		cache_Main_iota = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_iota(x_0_box)
		})
	})
	return cache_Main_iota
}

func Call_Main_s(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var y_1 gopurs_runtime.Value = y_1_loop
	_ = y_1
	var z_2 gopurs_runtime.Value = z_2_loop
	_ = z_2
	return gopurs_runtime.Apply2(x_0, z_2, gopurs_runtime.Apply(y_1, z_2))
}

func Call_Main_k(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var y_1 gopurs_runtime.Value = y_1_loop
	_ = y_1
	return x_0
}

func Call_Main_iota(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return gopurs_runtime.Apply2(x_0, Get_Main_s(), Get_Main_k())
}
