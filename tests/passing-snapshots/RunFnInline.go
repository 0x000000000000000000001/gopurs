package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_runFn3 gopurs_runtime.Value
var once_Main_runFn3 sync.Once

func Get_Main_runFn3() gopurs_runtime.Value {
	once_Main_runFn3.Do(func() {
		cache_Main_runFn3 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, c_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runFn3(f_0_box, a_1_box, b_2_box, c_3_box)
		})
	})
	return cache_Main_runFn3
}

var cache_Main_runFn3__3178790991 gopurs_runtime.Value
var once_Main_runFn3__3178790991 sync.Once

func Get_Main_runFn3__3178790991() gopurs_runtime.Value {
	once_Main_runFn3__3178790991.Do(func() {
		cache_Main_runFn3__3178790991 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, c_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_runFn3__3178790991(f_0_box, a_1_box.IntVal, b_2_box.IntVal, c_3_box.StrVal()))
		})
	})
	return cache_Main_runFn3__3178790991
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Str("Done").StrVal()))
	})
	return cache_Main_main
}

func Call_Main_runFn3(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, c_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	var c_3 gopurs_runtime.Value = c_3_loop
	_ = c_3
	return gopurs_runtime.Apply3(f_0, a_1, b_2, c_3)
}

func Call_Main_runFn3__3178790991(f_0_loop gopurs_runtime.Value, a_1_loop int64, b_2_loop int64, c_3_loop string) string {
runFn3__3178790991:
	for {
		if false {
			continue runFn3__3178790991
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var a_1 int64 = a_1_loop
		_ = a_1
		var b_2 int64 = b_2_loop
		_ = b_2
		var c_3 string = c_3_loop
		_ = c_3
		return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(a_1), gopurs_runtime.Int(b_2), gopurs_runtime.Str(c_3)).StrVal()
	}
}
