package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_squarePlusOne gopurs_runtime.Value
var once_Main_squarePlusOne sync.Once

func Get_Main_squarePlusOne() gopurs_runtime.Value {
	once_Main_squarePlusOne.Do(func() {
		cache_Main_squarePlusOne = gopurs_runtime.Func(func(value_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_squarePlusOne(value_0_box.IntVal))
		})
	})
	return cache_Main_squarePlusOne
}

var cache_Main_repeatEffects gopurs_runtime.Value
var once_Main_repeatEffects sync.Once

func Get_Main_repeatEffects() gopurs_runtime.Value {
	once_Main_repeatEffects.Do(func() {
		cache_Main_repeatEffects = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_repeatEffects(v_0_box.IntVal)
		})
	})
	return cache_Main_repeatEffects
}

var cache_Main_repeatApply gopurs_runtime.Value
var once_Main_repeatApply sync.Once

func Get_Main_repeatApply() gopurs_runtime.Value {
	once_Main_repeatApply.Do(func() {
		cache_Main_repeatApply = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_repeatApply(v_0_box.IntVal)
		})
	})
	return cache_Main_repeatApply
}

var cache_Main_makeThree gopurs_runtime.Value
var once_Main_makeThree sync.Once

func Get_Main_makeThree() gopurs_runtime.Value {
	once_Main_makeThree.Do(func() {
		cache_Main_makeThree = gopurs_runtime.Func(func(count_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_makeThree(count_0_box.IntVal)
		})
	})
	return cache_Main_makeThree
}

var cache_Main_makeSix gopurs_runtime.Value
var once_Main_makeSix sync.Once

func Get_Main_makeSix() gopurs_runtime.Value {
	once_Main_makeSix.Do(func() {
		cache_Main_makeSix = gopurs_runtime.Func(func(count_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_makeSix(count_0_box.IntVal)
		})
	})
	return cache_Main_makeSix
}

var cache_Main_composeRepeats gopurs_runtime.Value
var once_Main_composeRepeats sync.Once

func Get_Main_composeRepeats() gopurs_runtime.Value {
	once_Main_composeRepeats.Do(func() {
		cache_Main_composeRepeats = gopurs_runtime.Func2(func(leftCount_0_box gopurs_runtime.Value, rightCount_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_composeRepeats(leftCount_0_box.IntVal, rightCount_1_box.IntVal)
		})
	})
	return cache_Main_composeRepeats
}

var cache_Main_checkRepeat gopurs_runtime.Value
var once_Main_checkRepeat sync.Once

func Get_Main_checkRepeat() gopurs_runtime.Value {
	once_Main_checkRepeat.Do(func() {
		cache_Main_checkRepeat = gopurs_runtime.Func3(func(count_0_box gopurs_runtime.Value, step_1_box gopurs_runtime.Value, seed_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkRepeat(count_0_box.IntVal, step_1_box.IntVal, seed_2_box.IntVal)
		})
	})
	return cache_Main_checkRepeat
}

var cache_Main_checkPartialApplications gopurs_runtime.Value
var once_Main_checkPartialApplications sync.Once

func Get_Main_checkPartialApplications() gopurs_runtime.Value {
	once_Main_checkPartialApplications.Do(func() {
		cache_Main_checkPartialApplications = gopurs_runtime.Func(func(count_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkPartialApplications(count_0_box.IntVal)
		})
	})
	return cache_Main_checkPartialApplications
}

var cache_Main_checkEffectOrder gopurs_runtime.Value
var once_Main_checkEffectOrder sync.Once

func Get_Main_checkEffectOrder() gopurs_runtime.Value {
	once_Main_checkEffectOrder.Do(func() {
		cache_Main_checkEffectOrder = gopurs_runtime.Func(func(count_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkEffectOrder(count_0_box.IntVal)
		})
	})
	return cache_Main_checkEffectOrder
}

var cache_Main_checkArities gopurs_runtime.Value
var once_Main_checkArities sync.Once

func Get_Main_checkArities() gopurs_runtime.Value {
	once_Main_checkArities.Do(func() {
		cache_Main_checkArities = gopurs_runtime.Func(func(count_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkArities(count_0_box.IntVal)
		})
	})
	return cache_Main_checkArities
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(5)))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(7))), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(-3))), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_2), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_3_3), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(Call_Main_checkRepeat(int64(0), __local_var_6_6.IntVal, __local_var_5_5.IntVal), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Main_checkRepeat(int64(1), __local_var_6_6.IntVal, __local_var_5_5.IntVal), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Main_checkRepeat(int64(2), __local_var_6_6.IntVal, __local_var_5_5.IntVal), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Call_Main_checkRepeat(__local_var_4_4.IntVal, __local_var_6_6.IntVal, __local_var_5_5.IntVal), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(Call_Main_checkRepeat(__local_var_4_4.IntVal, int64(0), __local_var_5_5.IntVal), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Main_checkRepeat(__local_var_4_4.IntVal, int64(2), int64(0)), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(Call_Main_checkRepeat(int64(3), int64(2), int64(-4)), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(Call_Main_composeRepeats(__local_var_4_4.IntVal, int64(2)), gopurs_runtime.Func(func(value_14 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((value_14.IntVal) + (__local_var_6_6.IntVal))
			}), gopurs_runtime.Int(__local_var_5_5.IntVal)).IntVal), gopurs_runtime.Int((__local_var_5_5.IntVal)+(((__local_var_4_4.IntVal)+(int64(2)))*(__local_var_6_6.IntVal))))), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(Call_Main_composeRepeats(int64(0), int64(0)), gopurs_runtime.Func(func(value_15 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((value_15.IntVal) + (__local_var_6_6.IntVal))
			}), gopurs_runtime.Int(__local_var_5_5.IntVal)).IntVal), gopurs_runtime.Int(__local_var_5_5.IntVal))), gopurs_runtime.Value{})
			_ = __local_var_15_15
			__local_var_16_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(3))), gopurs_runtime.Value{})
			_ = __local_var_16_16
			__local_var_17_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_16_16), gopurs_runtime.Value{})
			_ = __local_var_17_17
			__local_var_18_18 := gopurs_runtime.Apply(Call_Main_checkPartialApplications(__local_var_17_17.IntVal), gopurs_runtime.Value{})
			_ = __local_var_18_18
			__local_var_19_19 := gopurs_runtime.Apply(Call_Main_checkEffectOrder(__local_var_17_17.IntVal), gopurs_runtime.Value{})
			_ = __local_var_19_19
			__local_var_20_20 := gopurs_runtime.Apply(Call_Main_checkArities(__local_var_17_17.IntVal), gopurs_runtime.Value{})
			_ = __local_var_20_20
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_squarePlusOne(value_0_loop int64) int64 {
	var value_0 int64 = value_0_loop
	_ = value_0
	return ((value_0) * (value_0)) + (int64(1))
}

func Call_Main_repeatEffects(v_0_loop int64) gopurs_runtime.Value {
repeatEffects:
	for {
		if false {
			continue repeatEffects
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var __t3 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t3 = gopurs_runtime.Func2(func(v1_1 gopurs_runtime.Value, seed_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return seed_2
					})
				})
				goto end_branch_3
			} else {

			}
		}
		{
			// TAST (Let): previous__2632949050_1_0 shape=App(Var) bindingType=(Func [(Func [Int] (ADT ["Effect","Effect"] [Int])), Int] (ADT ["Effect","Effect"] [Int]))
			previous__2632949050_1_0 := Call_Main_repeatEffects((v_0) - (int64(1)))
			_ = previous__2632949050_1_0
			__t3 = gopurs_runtime.Func2(func(step_2 gopurs_runtime.Value, seed_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
					__local_var_4_1 := gopurs_runtime.Apply2(previous__2632949050_1_0, step_2, gopurs_runtime.Int(seed_3.IntVal))
					_ = __local_var_4_1
					__local_var_5_2 := gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Value{})
					_ = __local_var_5_2
					return gopurs_runtime.Apply(gopurs_runtime.Apply(step_2, __local_var_5_2), gopurs_runtime.Value{})
				})
			})
		}
	end_branch_3:
		return __t3
	}
}

func Call_Main_repeatApply(v_0_loop int64) gopurs_runtime.Value {
repeatApply:
	for {
		if false {
			continue repeatApply
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var __t1 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t1 = gopurs_runtime.Func2(func(v1_1 gopurs_runtime.Value, seed_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(seed_2.IntVal)
				})
				goto end_branch_1
			} else {

			}
		}
		{
			// TAST (Let): previous__4066693242_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
			previous__4066693242_1_0 := Call_Main_repeatApply((v_0) - (int64(1)))
			_ = previous__4066693242_1_0
			__t1 = gopurs_runtime.Func2(func(step_2 gopurs_runtime.Value, seed_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.Apply(step_2, gopurs_runtime.Int(gopurs_runtime.Apply2(previous__4066693242_1_0, step_2, gopurs_runtime.Int(seed_3.IntVal)).IntVal)).IntVal)
			})
		}
	end_branch_1:
		return __t1
	}
}

func Call_Main_makeThree(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous__4066693242_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous__4066693242_1_0 := Call_Main_repeatApply(count_0)
	_ = previous__4066693242_1_0
	return gopurs_runtime.Func3(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, c_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int(gopurs_runtime.Apply2(previous__4066693242_1_0, gopurs_runtime.Func(func(value_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int((value_5.IntVal) + (int64(1)))
		}), gopurs_runtime.Int(((a_2.IntVal)+((int64(10))*(b_3.IntVal)))+((int64(100))*(c_4.IntVal)))).IntVal)
	})
}

func Call_Main_makeSix(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous__4066693242_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous__4066693242_1_0 := Call_Main_repeatApply(count_0)
	_ = previous__4066693242_1_0
	return gopurs_runtime.Func5(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, c_4 gopurs_runtime.Value, d_5 gopurs_runtime.Value, e_6 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.Apply2(previous__4066693242_1_0, gopurs_runtime.Func(func(value_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((value_8.IntVal) + (int64(1)))
			}), gopurs_runtime.Int((((((a_2.IntVal)+((int64(10))*(b_3.IntVal)))+((int64(100))*(c_4.IntVal)))+((int64(1000))*(d_5.IntVal)))+((int64(10000))*(e_6.IntVal)))+((int64(100000))*(f_7.IntVal)))).IntVal)
		})
	})
}

func Call_Main_composeRepeats(leftCount_0_loop int64, rightCount_1_loop int64) gopurs_runtime.Value {
	var leftCount_0 int64 = leftCount_0_loop
	_ = leftCount_0
	var rightCount_1 int64 = rightCount_1_loop
	_ = rightCount_1
	// TAST (Let): right__4066693242_2_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	right__4066693242_2_0 := Call_Main_repeatApply(rightCount_1)
	_ = right__4066693242_2_0
	// TAST (Let): left__4066693242_3_1 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	left__4066693242_3_1 := Call_Main_repeatApply(leftCount_0)
	_ = left__4066693242_3_1
	return gopurs_runtime.Func2(func(step_4 gopurs_runtime.Value, seed_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int(gopurs_runtime.Apply2(left__4066693242_3_1, step_4, gopurs_runtime.Int(gopurs_runtime.Apply2(right__4066693242_2_0, step_4, gopurs_runtime.Int(seed_5.IntVal)).IntVal)).IntVal)
	})
}

func Call_Main_checkRepeat(count_0_loop int64, step_1_loop int64, seed_2_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	var step_1 int64 = step_1_loop
	_ = step_1
	var seed_2 int64 = seed_2_loop
	_ = seed_2
	return gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(Call_Main_repeatApply(count_0), gopurs_runtime.Func(func(value_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int((value_3.IntVal) + (step_1))
	}), gopurs_runtime.Int(seed_2)).IntVal), gopurs_runtime.Int((seed_2)+((count_0)*(step_1)))))
}

func Call_Main_checkPartialApplications(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(Call_Main_repeatApply(count_0), Get_Main_squarePlusOne(), gopurs_runtime.Int(int64(1))).IntVal), gopurs_runtime.Int(int64(26))))
		_ = __local_var_1_0
		__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = __local_var_2_1
		__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(Call_Main_repeatApply(count_0), Get_Main_squarePlusOne(), gopurs_runtime.Int(int64(2))).IntVal), gopurs_runtime.Int(int64(677)))), gopurs_runtime.Value{})
		_ = __local_var_3_2
		__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Call_Main_repeatApply(count_0), Get_Main_squarePlusOne())), gopurs_runtime.Value{})
		_ = __local_var_4_3
		__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_3), gopurs_runtime.Value{})
		_ = __local_var_5_4
		__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Int(int64(1))).IntVal), gopurs_runtime.Int(int64(26)))), gopurs_runtime.Value{})
		_ = __local_var_6_5
		__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Int(int64(2))).IntVal), gopurs_runtime.Int(int64(677)))), gopurs_runtime.Value{})
		_ = __local_var_7_6
		return gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Int(int64(0))).IntVal), gopurs_runtime.Int(int64(5)))), gopurs_runtime.Value{})
	})
}

func Call_Main_checkEffectOrder(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(""))
		_ = __local_var_1_0
		__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = __local_var_2_1
		var Call_local_Main_step_3_2 func(gopurs_runtime.Value) gopurs_runtime.Value
		_ = Call_local_Main_step_3_2
		var step_3_2 gopurs_runtime.Value
		_ = step_3_2
		Call_local_Main_step_3_2 = func(value_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
			var value_3 gopurs_runtime.Value = value_3_loop
			_ = value_3
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_4_3 shape=App(Var) bindingType=Any
				__local_var_4_3 := gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(seen_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Str(((seen_4.StrVal()) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), value_3).StrVal())) + (","))
				}), __local_var_2_1)
				_ = __local_var_4_3
				__local_var_5_4 := gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Value{})
				_ = __local_var_5_4
				return gopurs_runtime.Int(((value_3.IntVal) * (value_3.IntVal)) + (int64(1)))
			})
		}
		step_3_2 = gopurs_runtime.Func(func(value_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Main_step_3_2(value_3_loop_val)
		})
		__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Main_repeatEffects(count_0), step_3_2, gopurs_runtime.Int(int64(1))), gopurs_runtime.Value{})
		_ = __local_var_4_5
		__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
		_ = __local_var_5_6
		__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(__local_var_4_5.IntVal), gopurs_runtime.Int(int64(26)))), gopurs_runtime.Value{})
		_ = __local_var_6_7
		__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(__local_var_5_6.StrVal()), gopurs_runtime.Str("1,2,5,"))), gopurs_runtime.Value{})
		_ = __local_var_7_8
		__local_var_8_9 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Str(""), __local_var_2_1), gopurs_runtime.Value{})
		_ = __local_var_8_9
		__local_var_9_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Call_Main_repeatEffects(count_0), step_3_2)), gopurs_runtime.Value{})
		_ = __local_var_9_10
		__local_var_10_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_9_10), gopurs_runtime.Value{})
		_ = __local_var_10_11
		__local_var_11_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
		_ = __local_var_11_12
		__local_var_12_13 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(__local_var_11_12.StrVal()), gopurs_runtime.Str(""))), gopurs_runtime.Value{})
		_ = __local_var_12_13
		__local_var_13_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_10_11, gopurs_runtime.Int(int64(1)))), gopurs_runtime.Value{})
		_ = __local_var_13_14
		__local_var_14_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_13_14), gopurs_runtime.Value{})
		_ = __local_var_14_15
		__local_var_15_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
		_ = __local_var_15_16
		__local_var_16_17 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(__local_var_15_16.StrVal()), gopurs_runtime.Str(""))), gopurs_runtime.Value{})
		_ = __local_var_16_17
		__local_var_17_18 := gopurs_runtime.Apply(__local_var_14_15, gopurs_runtime.Value{})
		_ = __local_var_17_18
		__local_var_18_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
		_ = __local_var_18_19
		__local_var_19_20 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(__local_var_17_18.IntVal), gopurs_runtime.Int(int64(26)))), gopurs_runtime.Value{})
		_ = __local_var_19_20
		__local_var_20_21 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(__local_var_18_19.StrVal()), gopurs_runtime.Str("1,2,5,"))), gopurs_runtime.Value{})
		_ = __local_var_20_21
		__local_var_21_22 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_10_11, gopurs_runtime.Int(int64(2))), gopurs_runtime.Value{})
		_ = __local_var_21_22
		__local_var_22_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
		_ = __local_var_22_23
		__local_var_23_24 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(__local_var_21_22.IntVal), gopurs_runtime.Int(int64(677)))), gopurs_runtime.Value{})
		_ = __local_var_23_24
		return gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(__local_var_22_23.StrVal()), gopurs_runtime.Str("1,2,5,2,5,26,"))), gopurs_runtime.Value{})
	})
}

func Call_Main_checkArities(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Call_Main_makeThree(count_0)), gopurs_runtime.Value{})
		_ = __local_var_1_0
		__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_0), gopurs_runtime.Value{})
		_ = __local_var_2_1
		__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply3(__local_var_2_1, gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3))).IntVal), gopurs_runtime.Int((count_0)+(int64(321))))), gopurs_runtime.Value{})
		_ = __local_var_3_2
		__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Int(int64(4)))), gopurs_runtime.Value{})
		_ = __local_var_4_3
		__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_3), gopurs_runtime.Value{})
		_ = __local_var_5_4
		__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(__local_var_5_4, gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(6))).IntVal), gopurs_runtime.Int((count_0)+(int64(654))))), gopurs_runtime.Value{})
		_ = __local_var_6_5
		__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(__local_var_5_4, gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(0))).IntVal), gopurs_runtime.Int((count_0)+(int64(34))))), gopurs_runtime.Value{})
		_ = __local_var_7_6
		__local_var_8_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Call_Main_makeSix(count_0)), gopurs_runtime.Value{})
		_ = __local_var_8_7
		__local_var_9_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_8_7), gopurs_runtime.Value{})
		_ = __local_var_9_8
		__local_var_10_9 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply6(__local_var_9_8, gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(6))).IntVal), gopurs_runtime.Int((count_0)+(int64(654321))))), gopurs_runtime.Value{})
		_ = __local_var_10_9
		__local_var_11_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply2(__local_var_9_8, gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2)))), gopurs_runtime.Value{})
		_ = __local_var_11_10
		__local_var_12_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_11_10), gopurs_runtime.Value{})
		_ = __local_var_12_11
		__local_var_13_12 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply4(__local_var_12_11, gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(6))).IntVal), gopurs_runtime.Int((count_0)+(int64(654321))))), gopurs_runtime.Value{})
		_ = __local_var_13_12
		__local_var_14_13 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply4(__local_var_12_11, gopurs_runtime.Int(int64(6)), gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(3))).IntVal), gopurs_runtime.Int((count_0)+(int64(345621))))), gopurs_runtime.Value{})
		_ = __local_var_14_13
		__local_var_15_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply5(__local_var_9_8, gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(5)))), gopurs_runtime.Value{})
		_ = __local_var_15_14
		__local_var_16_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_15_14), gopurs_runtime.Value{})
		_ = __local_var_16_15
		__local_var_17_16 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_16_15, gopurs_runtime.Int(int64(6))).IntVal), gopurs_runtime.Int((count_0)+(int64(654321))))), gopurs_runtime.Value{})
		_ = __local_var_17_16
		return gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(__local_var_16_15, gopurs_runtime.Int(int64(9))).IntVal), gopurs_runtime.Int((count_0)+(int64(954321))))), gopurs_runtime.Value{})
	})
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
