package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_y gopurs_runtime.Value
var once_Main_y sync.Once

func Get_Main_y() gopurs_runtime.Value {
	once_Main_y.Do(func() {
		cache_Main_y = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_y(a_0_box))
		})
	})
	return cache_Main_y
}

var cache_Main_x gopurs_runtime.Value
var once_Main_x sync.Once

func Get_Main_x() gopurs_runtime.Value {
	once_Main_x.Do(func() {
		cache_Main_x = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_x(a_0_box))
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

func Call_Main_y(a_0_loop gopurs_runtime.Value) float64 {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return Call_Main_x(a_0)
}

func Call_Main_x(a_0_loop gopurs_runtime.Value) float64 {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return (1.0) + (Call_Main_x(a_0))
}
