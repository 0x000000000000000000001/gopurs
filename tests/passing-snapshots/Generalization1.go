package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_sum gopurs_runtime.Value
var once_Main_sum sync.Once

func Get_Main_sum() gopurs_runtime.Value {
	once_Main_sum.Do(func() {
		cache_Main_sum = gopurs_runtime.Func3(func(dictSemiring_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_sum(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](dictSemiring_0_box), x_1_box, y_2_box)
		})
	})
	return cache_Main_sum
}

var cache_Main_sum__1139587456 gopurs_runtime.Value
var once_Main_sum__1139587456 sync.Once

func Get_Main_sum__1139587456() gopurs_runtime.Value {
	once_Main_sum__1139587456.Do(func() {
		cache_Main_sum__1139587456 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_sum__1139587456(x_0_box.IntVal, y_1_box.IntVal))
		})
	})
	return cache_Main_sum__1139587456
}

var cache_Main_sum__1069076592 gopurs_runtime.Value
var once_Main_sum__1069076592 sync.Once

func Get_Main_sum__1069076592() gopurs_runtime.Value {
	once_Main_sum__1069076592.Do(func() {
		cache_Main_sum__1069076592 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_sum__1069076592(x_0_box.FloatVal(), y_1_box.FloatVal()))
		})
	})
	return cache_Main_sum__1069076592
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Data_Show_ShowNumberImpl(3.0))), gopurs_runtime.Value{})
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Data_Show_ShowIntImpl(int64(3)))), gopurs_runtime.Value{})
			_ = __local_var_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_sum(dictSemiring_0_loop *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemiring_0 *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] = dictSemiring_0_loop
	_ = dictSemiring_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	var y_2 gopurs_runtime.Value = y_2_loop
	_ = y_2
	return gopurs_runtime.Apply2(dictSemiring_0.V0, x_1, y_2)
}

func Call_Main_sum__1139587456(x_0_loop int64, y_1_loop int64) int64 {
sum__1139587456:
	for {
		if false {
			continue sum__1139587456
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		var y_1 int64 = y_1_loop
		_ = y_1
		return gopurs_runtime.Int(gopurs_runtime.IntAdd(x_0, y_1)).IntVal
	}
}

func Call_Main_sum__1069076592(x_0_loop float64, y_1_loop float64) float64 {
sum__1069076592:
	for {
		if false {
			continue sum__1069076592
		}
		var x_0 float64 = x_0_loop
		_ = x_0
		var y_1 float64 = y_1_loop
		_ = y_1
		return gopurs_runtime.Float((x_0) + (y_1)).FloatVal()
	}
}
