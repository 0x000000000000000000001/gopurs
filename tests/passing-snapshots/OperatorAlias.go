package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_what gopurs_runtime.Value
var once_Main_what sync.Once

func Get_Main_what() gopurs_runtime.Value {
	once_Main_what.Do(func() {
		cache_Main_what = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_what(a_0_box, v_1_box)
		})
	})
	return cache_Main_what
}

var cache_Main_what__641934996 gopurs_runtime.Value
var once_Main_what__641934996 sync.Once

func Get_Main_what__641934996() gopurs_runtime.Value {
	once_Main_what__641934996.Do(func() {
		cache_Main_what__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_what__641934996(a_0_box, v_1_box)
		})
	})
	return cache_Main_what__641934996
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Str("Done").StrVal()))
	})
	return cache_Main_main
}

func Call_Main_what(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return a_0
}

func Call_Main_what__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return a_0
}
