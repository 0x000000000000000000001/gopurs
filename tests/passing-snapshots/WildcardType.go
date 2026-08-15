package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_f2 gopurs_runtime.Value
var once_Main_f2 sync.Once

func Get_Main_f2() gopurs_runtime.Value {
	once_Main_f2.Do(func() {
		cache_Main_f2 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_f2(v_0_box))
		})
	})
	return cache_Main_f2
}

var cache_Main_f1 gopurs_runtime.Value
var once_Main_f1 sync.Once

func Get_Main_f1() gopurs_runtime.Value {
	once_Main_f1.Do(func() {
		cache_Main_f1 = gopurs_runtime.Func(func(g_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f1(g_0_box)
		})
	})
	return cache_Main_f1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_f2(v_0_loop gopurs_runtime.Value) string {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return "Done"
}

func Call_Main_f1(g_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var g_0 gopurs_runtime.Value = g_0_loop
	_ = g_0
	return gopurs_runtime.Apply(g_0, gopurs_runtime.Int(1))
}
