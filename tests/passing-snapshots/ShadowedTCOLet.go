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
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Call_Data_Show_show(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber())))), gopurs_runtime.Float(1.0)).StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
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

func Rebox_Main_1386611502_3263178038(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[float64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
