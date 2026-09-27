package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func4(func(dictPartial_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_f(dictPartial_0_box, x_1_box.FloatVal(), y_2_box.FloatVal(), z_3_box.FloatVal()))
		})
	})
	return cache_Main_f
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Data_Show_ShowNumberImpl(1.0)))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_f(dictPartial_0_loop gopurs_runtime.Value, x_1_loop float64, y_2_loop float64, z_3_loop float64) float64 {
	var dictPartial_0 gopurs_runtime.Value = dictPartial_0_loop
	_ = dictPartial_0
	var x_1 float64 = x_1_loop
	_ = x_1
	var y_2 float64 = y_2_loop
	_ = y_2
	var z_3 float64 = z_3_loop
	_ = z_3
	var __t0 float64
	{
		if ((x_1) == (1.0)) && (((z_3) == (2.0)) && ((y_2) == (3.0))) {
			__t0 = 1.0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() float64 { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}
