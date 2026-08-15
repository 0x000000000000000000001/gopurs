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

var cache_Main_fn gopurs_runtime.Value
var once_Main_fn sync.Once

func Get_Main_fn() gopurs_runtime.Value {
	once_Main_fn.Do(func() {
		cache_Main_fn = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_fn(_dollar___unused_0_box, v_1_box.FloatVal()))
		})
	})
	return cache_Main_fn
}

func Call_Main_fn(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop float64) float64 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 float64 = v_1_loop
	_ = v_1
	var __t0 float64
	{
		if (v_1) == (0.0) {
			__t0 = 0.0
			goto end_branch_0
		} else {

		}
	}
	{
		if (v_1) == (1.0) {
			__t0 = 2.0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().FloatVal()
	}
end_branch_0:
	return gopurs_runtime.Float(__t0).FloatVal()
}
