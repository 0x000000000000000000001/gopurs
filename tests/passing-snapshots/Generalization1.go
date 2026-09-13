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

var cache_Main_sum__1135989556 gopurs_runtime.Value
var once_Main_sum__1135989556 sync.Once

func Get_Main_sum__1135989556() gopurs_runtime.Value {
	once_Main_sum__1135989556.Do(func() {
		cache_Main_sum__1135989556 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_sum__1135989556(x_0_box.IntVal, y_1_box.IntVal))
		})
	})
	return cache_Main_sum__1135989556
}

var cache_Main_sum__2506932 gopurs_runtime.Value
var once_Main_sum__2506932 sync.Once

func Get_Main_sum__2506932() gopurs_runtime.Value {
	once_Main_sum__2506932.Do(func() {
		cache_Main_sum__2506932 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_sum__2506932(x_0_box.FloatVal(), y_1_box.FloatVal()))
		})
	})
	return cache_Main_sum__2506932
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(3.0)).StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(int64(3))).StrVal())), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			}))
		}))
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

func Call_Main_sum__1135989556(x_0_loop int64, y_1_loop int64) int64 {
sum__1135989556:
	for {
		if false {
			continue sum__1135989556
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		var y_1 int64 = y_1_loop
		_ = y_1
		return gopurs_runtime.Int((x_0) + (y_1)).IntVal
	}
}

func Call_Main_sum__2506932(x_0_loop float64, y_1_loop float64) float64 {
sum__2506932:
	for {
		if false {
			continue sum__2506932
		}
		var x_0 float64 = x_0_loop
		_ = x_0
		var y_1 float64 = y_1_loop
		_ = y_1
		return gopurs_runtime.Float((x_0) + (y_1)).FloatVal()
	}
}
