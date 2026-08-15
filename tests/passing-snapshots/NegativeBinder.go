package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_test2(x_0_box.FloatVal(), y_1_box.FloatVal()))
		})
	})
	return cache_Main_test2
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_test(v_0_box.FloatVal()))
		})
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_test2(x_0_loop float64, y_1_loop float64) bool {
	var x_0 float64 = x_0_loop
	_ = x_0
	var y_1 float64 = y_1_loop
	_ = y_1
	var __t0 bool
	{
		if (x_0) == (-1.0) {
			__t0 = ((y_1) == (-1.0)) != (true)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = true
	}
end_branch_0:
	return __t0
}

func Call_Main_test(v_0_loop float64) bool {
	var v_0 float64 = v_0_loop
	_ = v_0
	return ((v_0) == (-1.0)) != (true)
}
