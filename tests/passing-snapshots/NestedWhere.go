package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

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
		cache_Main_f = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f(x_0_box.FloatVal())
		})
	})
	return cache_Main_f
}

func Call_Main_f(x_0_loop float64) gopurs_runtime.Value {
	var x_0 float64 = x_0_loop
	_ = x_0
	var go1_1_0_0 gopurs_runtime.Value
	_ = go1_1_0_0
	var go__go_1_1_1 gopurs_runtime.Value
	_ = go__go_1_1_1
	go1_1_0_0 = gopurs_runtime.Func(func(x2_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(go__go_1_1_1, gopurs_runtime.Float(x2_2.FloatVal()))
	})
	go__go_1_1_1 = gopurs_runtime.Func(func(x2_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(go1_1_0_0, gopurs_runtime.Float((x2_2.FloatVal())-(1.0)))
	})
	return gopurs_runtime.Apply(go__go_1_1_1, gopurs_runtime.Float(x_0))
}
