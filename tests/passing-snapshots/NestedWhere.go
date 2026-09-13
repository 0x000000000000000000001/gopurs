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
	var Call_local_Main_go__4228559310_1_0_0 func(float64) gopurs_runtime.Value
	_ = Call_local_Main_go__4228559310_1_0_0
	var go__4228559310_1_0_0 gopurs_runtime.Value
	_ = go__4228559310_1_0_0
	var Call_local_Main_go1_1_1_1 func(float64) gopurs_runtime.Value
	_ = Call_local_Main_go1_1_1_1
	var go1_1_1_1 gopurs_runtime.Value
	_ = go1_1_1_1
	var Call_local_Main_go__go_1_2_2 func(float64) gopurs_runtime.Value
	_ = Call_local_Main_go__go_1_2_2
	var go__go_1_2_2 gopurs_runtime.Value
	_ = go__go_1_2_2
	Call_local_Main_go__4228559310_1_0_0 = func(x2_2_loop float64) gopurs_runtime.Value {
	go__4228559310_1_0_0:
		for {
			if false {
				continue go__4228559310_1_0_0
			}
			var x2_2 float64 = x2_2_loop
			_ = x2_2
			return Call_local_Main_go1_1_1_1((x2_2) - (1.0))
		}
	}
	go__4228559310_1_0_0 = gopurs_runtime.Func(func(x2_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_local_Main_go__4228559310_1_0_0(x2_2_loop_val.FloatVal())
	})
	Call_local_Main_go1_1_1_1 = func(x2_2_loop float64) gopurs_runtime.Value {
	go1_1_1_1:
		for {
			if false {
				continue go1_1_1_1
			}
			var x2_2 float64 = x2_2_loop
			_ = x2_2
			return Call_local_Main_go__4228559310_1_0_0(x2_2)
		}
	}
	go1_1_1_1 = gopurs_runtime.Func(func(x2_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_local_Main_go1_1_1_1(x2_2_loop_val.FloatVal())
	})
	Call_local_Main_go__go_1_2_2 = func(x2_2_loop float64) gopurs_runtime.Value {
	go__go_1_2_2:
		for {
			if false {
				continue go__go_1_2_2
			}
			var x2_2 float64 = x2_2_loop
			_ = x2_2
			return Call_local_Main_go1_1_1_1((x2_2) - (1.0))
		}
	}
	go__go_1_2_2 = gopurs_runtime.Func(func(x2_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_local_Main_go__go_1_2_2(x2_2_loop_val.FloatVal())
	})
	return Call_local_Main_go__4228559310_1_0_0(x_0)
}
