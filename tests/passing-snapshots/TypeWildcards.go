package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_testTopLevel gopurs_runtime.Value
var once_Main_testTopLevel sync.Once

func Get_Main_testTopLevel() gopurs_runtime.Value {
	once_Main_testTopLevel.Do(func() {
		cache_Main_testTopLevel = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_testTopLevel(n_0_box.FloatVal()))
		})
	})
	return cache_Main_testTopLevel
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), f_1_box, a_2_box)
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

func Call_Main_testTopLevel(n_0_loop float64) float64 {
	var n_0 float64 = n_0_loop
	_ = n_0
	return (n_0) + (1.0)
}

func Call_Main_test(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
	_ = dictEq_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	var Call_local_Main_go__go_3_0_0 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
	_ = Call_local_Main_go__go_3_0_0
	var go__go_3_0_0 gopurs_runtime.Value
	_ = go__go_3_0_0
	Call_local_Main_go__go_3_0_0 = func(v_4_loop gopurs_runtime.Value, v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
	go__go_3_0_0:
		for {
			if false {
				continue go__go_3_0_0
			}
			var v_4 gopurs_runtime.Value = v_4_loop
			_ = v_4
			var v1_5 gopurs_runtime.Value = v1_5_loop
			_ = v1_5
			var __t1 gopurs_runtime.Value
			{
				if (gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_4, v1_5).IntVal) != (0) {
					__t1 = v_4
					goto end_branch_1
				} else {

				}
			}
			{
				v_4_loop = gopurs_runtime.Apply(f_1, v_4)
				v1_5_loop = v_4
				continue go__go_3_0_0
				__t1 = gopurs_runtime.Value{}
			}
		end_branch_1:
			return __t1
		}
	}
	go__go_3_0_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Main_go__go_3_0_0(v_4_loop_val, v1_5_loop_val)
		})
	})
	return Call_local_Main_go__go_3_0_0(gopurs_runtime.Apply(f_1, a_2), a_2)
}
