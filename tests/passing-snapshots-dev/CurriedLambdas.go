package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_rightIsSymbol gopurs_runtime.Value
var once_Main_rightIsSymbol sync.Once

func Get_Main_rightIsSymbol() gopurs_runtime.Value {
	once_Main_rightIsSymbol.Do(func() {
		cache_Main_rightIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("right")
		}))
	})
	return cache_Main_rightIsSymbol
}

var cache_Main_leftIsSymbol gopurs_runtime.Value
var once_Main_leftIsSymbol sync.Once

func Get_Main_leftIsSymbol() gopurs_runtime.Value {
	once_Main_leftIsSymbol.Do(func() {
		cache_Main_leftIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("left")
		}))
	})
	return cache_Main_leftIsSymbol
}

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

var cache_Main_repeatApply__gopurs_counted_function_0 gopurs_runtime.Value
var once_Main_repeatApply__gopurs_counted_function_0 sync.Once

func Get_Main_repeatApply__gopurs_counted_function_0() gopurs_runtime.Value {
	once_Main_repeatApply__gopurs_counted_function_0.Do(func() {
		cache_Main_repeatApply__gopurs_counted_function_0 = gopurs_runtime.Func3(func(remaining_0_box gopurs_runtime.Value, callback_1_box gopurs_runtime.Value, result_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_repeatApply__gopurs_counted_function_0(remaining_0_box.IntVal, callback_1_box, result_2_box.IntVal))
		})
	})
	return cache_Main_repeatApply__gopurs_counted_function_0
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

var cache_Main_makeUncurriedBoundary gopurs_runtime.Value
var once_Main_makeUncurriedBoundary sync.Once

func Get_Main_makeUncurriedBoundary() gopurs_runtime.Value {
	once_Main_makeUncurriedBoundary.Do(func() {
		cache_Main_makeUncurriedBoundary = gopurs_runtime.Func(func(count_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_makeUncurriedBoundary(count_0_box.IntVal)
		})
	})
	return cache_Main_makeUncurriedBoundary
}

var cache_Main_makeTypeAppBoundary gopurs_runtime.Value
var once_Main_makeTypeAppBoundary sync.Once

func Get_Main_makeTypeAppBoundary() gopurs_runtime.Value {
	once_Main_makeTypeAppBoundary.Do(func() {
		cache_Main_makeTypeAppBoundary = gopurs_runtime.Func(func(count_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_makeTypeAppBoundary(count_0_box.IntVal)
		})
	})
	return cache_Main_makeTypeAppBoundary
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

var cache_Main_makeRecursiveBoundary gopurs_runtime.Value
var once_Main_makeRecursiveBoundary sync.Once

func Get_Main_makeRecursiveBoundary() gopurs_runtime.Value {
	once_Main_makeRecursiveBoundary.Do(func() {
		cache_Main_makeRecursiveBoundary = gopurs_runtime.Func(func(count_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_makeRecursiveBoundary(count_0_box.IntVal)
		})
	})
	return cache_Main_makeRecursiveBoundary
}

var cache_Main_makeLetBoundary gopurs_runtime.Value
var once_Main_makeLetBoundary sync.Once

func Get_Main_makeLetBoundary() gopurs_runtime.Value {
	once_Main_makeLetBoundary.Do(func() {
		cache_Main_makeLetBoundary = gopurs_runtime.Func(func(count_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_makeLetBoundary(count_0_box.IntVal)
		})
	})
	return cache_Main_makeLetBoundary
}

var cache_Main_makeEffectBoundary gopurs_runtime.Value
var once_Main_makeEffectBoundary sync.Once

func Get_Main_makeEffectBoundary() gopurs_runtime.Value {
	once_Main_makeEffectBoundary.Do(func() {
		cache_Main_makeEffectBoundary = gopurs_runtime.Func2(func(count_0_box gopurs_runtime.Value, state_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_makeEffectBoundary(count_0_box.IntVal, state_1_box)
		})
	})
	return cache_Main_makeEffectBoundary
}

var cache_Main_makeCallBoundary gopurs_runtime.Value
var once_Main_makeCallBoundary sync.Once

func Get_Main_makeCallBoundary() gopurs_runtime.Value {
	once_Main_makeCallBoundary.Do(func() {
		cache_Main_makeCallBoundary = gopurs_runtime.Func2(func(count_0_box gopurs_runtime.Value, factory_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_makeCallBoundary(count_0_box.IntVal, factory_1_box)
		})
	})
	return cache_Main_makeCallBoundary
}

var cache_Main_makeBranchBoundary gopurs_runtime.Value
var once_Main_makeBranchBoundary sync.Once

func Get_Main_makeBranchBoundary() gopurs_runtime.Value {
	once_Main_makeBranchBoundary.Do(func() {
		cache_Main_makeBranchBoundary = gopurs_runtime.Func(func(count_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_makeBranchBoundary(count_0_box.IntVal)
		})
	})
	return cache_Main_makeBranchBoundary
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

var cache_Main_checkEffectBoundary gopurs_runtime.Value
var once_Main_checkEffectBoundary sync.Once

func Get_Main_checkEffectBoundary() gopurs_runtime.Value {
	once_Main_checkEffectBoundary.Do(func() {
		cache_Main_checkEffectBoundary = gopurs_runtime.Func(func(count_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkEffectBoundary(count_0_box.IntVal)
		})
	})
	return cache_Main_checkEffectBoundary
}

var cache_Main_checkBoundaries gopurs_runtime.Value
var once_Main_checkBoundaries sync.Once

func Get_Main_checkBoundaries() gopurs_runtime.Value {
	once_Main_checkBoundaries.Do(func() {
		cache_Main_checkBoundaries = gopurs_runtime.Func(func(count_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkBoundaries(count_0_box.IntVal)
		})
	})
	return cache_Main_checkBoundaries
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
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(5)))
			_ = __local_var_0_0
			depthRef_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = depthRef_1_1
			seedRef_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(7))), gopurs_runtime.Value{})
			_ = seedRef_2_2
			stepRef_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(-3))), gopurs_runtime.Value{})
			_ = stepRef_3_3
			depth_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), depthRef_1_1), gopurs_runtime.Value{})
			_ = depth_4_4
			seed_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), seedRef_2_2), gopurs_runtime.Value{})
			_ = seed_5_5
			step_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), stepRef_3_3), gopurs_runtime.Value{})
			_ = step_6_6
			_dollar___unused_7_7 := gopurs_runtime.Apply(Call_Main_checkRepeat(int64(0), step_6_6.IntVal, seed_5_5.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			_dollar___unused_8_8 := gopurs_runtime.Apply(Call_Main_checkRepeat(int64(1), step_6_6.IntVal, seed_5_5.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_8_8
			_dollar___unused_9_9 := gopurs_runtime.Apply(Call_Main_checkRepeat(int64(2), step_6_6.IntVal, seed_5_5.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_9_9
			_dollar___unused_10_10 := gopurs_runtime.Apply(Call_Main_checkRepeat(depth_4_4.IntVal, step_6_6.IntVal, seed_5_5.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_10_10
			_dollar___unused_11_11 := gopurs_runtime.Apply(Call_Main_checkRepeat(depth_4_4.IntVal, int64(0), seed_5_5.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_11_11
			_dollar___unused_12_12 := gopurs_runtime.Apply(Call_Main_checkRepeat(depth_4_4.IntVal, int64(2), int64(0)), gopurs_runtime.Value{})
			_ = _dollar___unused_12_12
			_dollar___unused_13_13 := gopurs_runtime.Apply(Call_Main_checkRepeat(int64(3), int64(2), int64(-4)), gopurs_runtime.Value{})
			_ = _dollar___unused_13_13
			_dollar___unused_14_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Apply2(Call_Main_composeRepeats(depth_4_4.IntVal, int64(2)), gopurs_runtime.Func(func(value_14 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_14.IntVal, step_6_6.IntVal))
				}), gopurs_runtime.Int(seed_5_5.IntVal)).IntVal, gopurs_runtime.IntAdd(seed_5_5.IntVal, gopurs_runtime.IntMul(gopurs_runtime.IntAdd(depth_4_4.IntVal, int64(2)), step_6_6.IntVal))}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_14_14
			_dollar___unused_15_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.Apply2(Call_Main_composeRepeats(int64(0), int64(0)), gopurs_runtime.Func(func(value_15 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_15.IntVal, step_6_6.IntVal))
				}), gopurs_runtime.Int(seed_5_5.IntVal)).IntVal, seed_5_5.IntVal}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_15_15
			smallDepthRef_16_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(3))), gopurs_runtime.Value{})
			_ = smallDepthRef_16_16
			smallDepth_17_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), smallDepthRef_16_16), gopurs_runtime.Value{})
			_ = smallDepth_17_17
			_dollar___unused_18_18 := gopurs_runtime.Apply(Call_Main_checkPartialApplications(smallDepth_17_17.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_18_18
			_dollar___unused_19_19 := gopurs_runtime.Apply(Call_Main_checkEffectOrder(smallDepth_17_17.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_19_19
			_dollar___unused_20_20 := gopurs_runtime.Apply(Call_Main_checkArities(smallDepth_17_17.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_20_20
			_dollar___unused_21_21 := gopurs_runtime.Apply(Call_Main_checkBoundaries(smallDepth_17_17.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_21_21
			_dollar___unused_22_22 := gopurs_runtime.Apply(Call_Main_checkEffectBoundary(smallDepth_17_17.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_22_22
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3959995044_3790796878(Rebox_Main_3790796878_3959995044(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_rightIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_leftIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))))))}
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1739864324_1386611502(Rebox_Main_1386611502_1739864324(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_leftIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_rightIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))))))}
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3959995044_3790796878(Rebox_Main_3790796878_3959995044(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_rightIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_leftIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))))))}
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1739864324_1386611502(Rebox_Main_1386611502_1739864324(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_leftIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_rightIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))))))}
	})
	return cache_Main___cached_dict_3
}

func Call_Main_squarePlusOne(value_0_loop int64) int64 {
	var value_0 int64 = value_0_loop
	_ = value_0
	return gopurs_runtime.IntAdd(gopurs_runtime.IntMul(value_0, value_0), int64(1))
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
			// TAST (Let): previous_1_0 shape=App(Var) bindingType=(Func [(Func [Int] (ADT ["Effect","Effect"] [Int])), Int] (ADT ["Effect","Effect"] [Int]))
			previous_1_0 := Call_Main_repeatEffects(gopurs_runtime.IntSub(v_0, int64(1)))
			_ = previous_1_0
			__t3 = gopurs_runtime.Func2(func(step_2 gopurs_runtime.Value, seed_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
					__local_var_4_1 := gopurs_runtime.Apply2(previous_1_0, step_2, gopurs_runtime.Int(seed_3.IntVal))
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

func Call_Main_repeatApply__gopurs_counted_function_0(remaining_0_loop int64, callback_1_loop gopurs_runtime.Value, result_2_loop int64) int64 {
	if (remaining_0_loop) == (int64(int32(remaining_0_loop))) {
		return func() int64 {
			remaining_0_loop := int32(remaining_0_loop)
			_ = remaining_0_loop
		repeatApply__gopurs_counted_function_0:
			for {
				if false {
					continue repeatApply__gopurs_counted_function_0
				}
				remaining_0 := remaining_0_loop
				_ = remaining_0
				var callback_1 gopurs_runtime.Value = callback_1_loop
				_ = callback_1
				var result_2 int64 = result_2_loop
				_ = result_2
				var __t0 int64
				{
					if (remaining_0) == (int32(0)) {
						__t0 = result_2
						goto end_branch_0
					} else {

					}
				}
				{
					remaining_0_loop = (remaining_0) - (int32(1))
					callback_1_loop = callback_1
					result_2_loop = gopurs_runtime.Apply(callback_1, gopurs_runtime.Int(result_2)).IntVal
					continue repeatApply__gopurs_counted_function_0
					__t0 = func() int64 { panic("unreachable") }()
				}
			end_branch_0:
				return __t0
			}
		}()
	} else {

	}
repeatApply__gopurs_counted_function_0:
	for {
		if false {
			continue repeatApply__gopurs_counted_function_0
		}
		var remaining_0 int64 = remaining_0_loop
		_ = remaining_0
		var callback_1 gopurs_runtime.Value = callback_1_loop
		_ = callback_1
		var result_2 int64 = result_2_loop
		_ = result_2
		var __t0 int64
		{
			if (remaining_0) == (int64(0)) {
				__t0 = result_2
				goto end_branch_0
			} else {

			}
		}
		{
			remaining_0_loop = gopurs_runtime.IntSub(remaining_0, int64(1))
			callback_1_loop = callback_1
			result_2_loop = gopurs_runtime.Apply(callback_1, gopurs_runtime.Int(result_2)).IntVal
			continue repeatApply__gopurs_counted_function_0
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
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
		var __t2 gopurs_runtime.Value
		{
			if (v_0) >= (int64(0)) {
				__t2 = gopurs_runtime.Func2(func(step_2 gopurs_runtime.Value, seed_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_Main_repeatApply__gopurs_counted_function_0(v_0, step_2, seed_3.IntVal))
				})
				goto end_branch_2
			} else {

			}
		}
		{
			var __t1 gopurs_runtime.Value
			{
				if (v_0) == (int64(0)) {
					__t1 = gopurs_runtime.Func2(func(v1_1 gopurs_runtime.Value, seed_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return seed_2
					})
					goto end_branch_1
				} else {

				}
			}
			{
				// TAST (Let): previous_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
				previous_1_0 := Call_Main_repeatApply(gopurs_runtime.IntSub(v_0, int64(1)))
				_ = previous_1_0
				__t1 = gopurs_runtime.Func2(func(step_2 gopurs_runtime.Value, seed_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(step_2, gopurs_runtime.Int(gopurs_runtime.Apply2(previous_1_0, step_2, gopurs_runtime.Int(seed_3.IntVal)).IntVal))
				})
			}
		end_branch_1:
			__t2 = __t1
		}
	end_branch_2:
		return __t2
	}
}

func Call_Main_makeUncurriedBoundary(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous_1_0 := Call_Main_repeatApply(count_0)
	_ = previous_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(previous_1_0, gopurs_runtime.Func(func(value_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_5.IntVal, int64(1)))
			}), gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(a_2.IntVal, gopurs_runtime.IntMul(int64(10), b_3.IntVal)), gopurs_runtime.IntMul(int64(100), __local_var_4.IntVal))))
		})
	})
}

func Call_Main_makeTypeAppBoundary(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous_1_0 := Call_Main_repeatApply(count_0)
	_ = previous_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict2("left", "right", gopurs_runtime.Apply2(previous_1_0, gopurs_runtime.Func(func(value_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_4.IntVal, int64(1)))
			}), gopurs_runtime.Int(a_2.IntVal)), b_3)
		})
	})
}

func Call_Main_makeThree(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous_1_0 := Call_Main_repeatApply(count_0)
	_ = previous_1_0
	return gopurs_runtime.Func3(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, c_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(previous_1_0, gopurs_runtime.Func(func(value_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_5.IntVal, int64(1)))
		}), gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(a_2.IntVal, gopurs_runtime.IntMul(int64(10), b_3.IntVal)), gopurs_runtime.IntMul(int64(100), c_4.IntVal))))
	})
}

func Call_Main_makeSix(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous_1_0 := Call_Main_repeatApply(count_0)
	_ = previous_1_0
	return gopurs_runtime.Func5(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, c_4 gopurs_runtime.Value, d_5 gopurs_runtime.Value, e_6 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(previous_1_0, gopurs_runtime.Func(func(value_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_8.IntVal, int64(1)))
			}), gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(a_2.IntVal, gopurs_runtime.IntMul(int64(10), b_3.IntVal)), gopurs_runtime.IntMul(int64(100), c_4.IntVal)), gopurs_runtime.IntMul(int64(1000), d_5.IntVal)), gopurs_runtime.IntMul(int64(10000), e_6.IntVal)), gopurs_runtime.IntMul(int64(100000), f_7.IntVal))))
		})
	})
}

func Call_Main_makeRecursiveBoundary(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous_1_0 := Call_Main_repeatApply(count_0)
	_ = previous_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var visit_3_1_0 gopurs_runtime.Value
		_ = visit_3_1_0
		var visit_3_1_0_cell *gopurs_runtime.Value
		_ = visit_3_1_0_cell
		// FALLBACK TCO: isLoop=false len=1
		visit_3_1_0 = gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 gopurs_runtime.Value
			{
				if (n_4.IntVal) == (int64(0)) {
					__t2 = gopurs_runtime.Apply2(previous_1_0, gopurs_runtime.Func(func(value_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_5.IntVal, int64(1)))
					}), gopurs_runtime.Int(a_2.IntVal))
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.Apply((*visit_3_1_0_cell), gopurs_runtime.Int(gopurs_runtime.IntSub(n_4.IntVal, int64(1)))).IntVal, int64(1)))
			}
		end_branch_2:
			return __t2
		})
		visit_3_1_0_cell = &visit_3_1_0
		return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.Apply(visit_3_1_0, gopurs_runtime.Int(b_4.IntVal)).IntVal, gopurs_runtime.Apply(visit_3_1_0, gopurs_runtime.Int(gopurs_runtime.IntAdd(b_4.IntVal, int64(1)))).IntVal))
		})
	})
}

func Call_Main_makeLetBoundary(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous_1_0 := Call_Main_repeatApply(count_0)
	_ = previous_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): cached_3_1 shape=App(Other) bindingType=Int
		cached_3_1 := gopurs_runtime.Apply2(previous_1_0, Get_Main_squarePlusOne(), gopurs_runtime.Int(a_2.IntVal)).IntVal
		_ = cached_3_1
		return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(cached_3_1, cached_3_1), b_4.IntVal))
		})
	})
}

func Call_Main_makeEffectBoundary(count_0_loop int64, state_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	var state_1 gopurs_runtime.Value = state_1_loop
	_ = state_1
	// TAST (Let): previous_2_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous_2_0 := Call_Main_repeatApply(count_0)
	_ = previous_2_0
	return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
			__local_var_4_1 := gopurs_runtime.Apply(Get_Effect_Ref_read(), state_1)
			_ = __local_var_4_1
			offset_5_2 := gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Value{})
			_ = offset_5_2
			_dollar___unused_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Int(gopurs_runtime.IntAdd(offset_5_2.IntVal, int64(1))), state_1), gopurs_runtime.Value{})
			_ = _dollar___unused_6_3
			return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.Apply2(previous_2_0, gopurs_runtime.Func(func(value_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_8.IntVal, int64(1)))
				}), gopurs_runtime.Int(gopurs_runtime.IntAdd(a_3.IntVal, offset_5_2.IntVal))).IntVal, b_7.IntVal))
			})
		})
	})
}

func Call_Main_makeCallBoundary(count_0_loop int64, factory_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	var factory_1 gopurs_runtime.Value = factory_1_loop
	_ = factory_1
	// TAST (Let): previous_2_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous_2_0 := Call_Main_repeatApply(count_0)
	_ = previous_2_0
	return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(factory_1, gopurs_runtime.Int(gopurs_runtime.Apply2(previous_2_0, gopurs_runtime.Func(func(value_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_4.IntVal, int64(1)))
		}), gopurs_runtime.Int(a_3.IntVal)).IntVal))
	})
}

func Call_Main_makeBranchBoundary(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous_1_0 := Call_Main_repeatApply(count_0)
	_ = previous_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			if (a_2.IntVal) < (int64(0)) {
				__t1 = gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(previous_1_0, gopurs_runtime.Func(func(value_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_4.IntVal, int64(1)))
					}), gopurs_runtime.Int(gopurs_runtime.IntSub(a_2.IntVal, b_3.IntVal)))
				})
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(previous_1_0, gopurs_runtime.Func(func(value_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_4.IntVal, int64(1)))
				}), gopurs_runtime.Int(gopurs_runtime.IntAdd(a_2.IntVal, gopurs_runtime.IntMul(int64(2), b_3.IntVal))))
			})
		}
	end_branch_1:
		return __t1
	})
}

func Call_Main_composeRepeats(leftCount_0_loop int64, rightCount_1_loop int64) gopurs_runtime.Value {
	var leftCount_0 int64 = leftCount_0_loop
	_ = leftCount_0
	var rightCount_1 int64 = rightCount_1_loop
	_ = rightCount_1
	// TAST (Let): right_2_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	right_2_0 := Call_Main_repeatApply(rightCount_1)
	_ = right_2_0
	// TAST (Let): left_3_1 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	left_3_1 := Call_Main_repeatApply(leftCount_0)
	_ = left_3_1
	return gopurs_runtime.Func2(func(step_4 gopurs_runtime.Value, seed_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(left_3_1, step_4, gopurs_runtime.Int(gopurs_runtime.Apply2(right_2_0, step_4, gopurs_runtime.Int(seed_5.IntVal)).IntVal))
	})
}

func Call_Main_checkRepeat(count_0_loop int64, step_1_loop int64, seed_2_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	var step_1 int64 = step_1_loop
	_ = step_1
	var seed_2 int64 = seed_2_loop
	_ = seed_2
	return gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
		orig := struct {
			actual   int64
			expected int64
			/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
		}{gopurs_runtime.Apply2(Call_Main_repeatApply(count_0), gopurs_runtime.Func(func(value_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_3.IntVal, step_1))
		}), gopurs_runtime.Int(seed_2)).IntVal, gopurs_runtime.IntAdd(seed_2, gopurs_runtime.IntMul(count_0, step_1))}
		_ = orig
		return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
	}())
}

func Call_Main_checkPartialApplications(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply2(Call_Main_repeatApply(count_0), Get_Main_squarePlusOne(), gopurs_runtime.Int(int64(1))).IntVal, int64(26)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}())
		_ = __local_var_1_0
		_dollar___unused_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = _dollar___unused_2_1
		_dollar___unused_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply2(Call_Main_repeatApply(count_0), Get_Main_squarePlusOne(), gopurs_runtime.Int(int64(2))).IntVal, int64(677)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_3_2
		partialRef_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(Call_Main_repeatApply(count_0), Get_Main_squarePlusOne())), gopurs_runtime.Value{})
		_ = partialRef_4_3
		partial_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), partialRef_4_3), gopurs_runtime.Value{})
		_ = partial_5_4
		_dollar___unused_6_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(partial_5_4, gopurs_runtime.Int(int64(1))).IntVal, int64(26)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_6_5
		_dollar___unused_7_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(partial_5_4, gopurs_runtime.Int(int64(2))).IntVal, int64(677)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_7_6
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(partial_5_4, gopurs_runtime.Int(int64(0))).IntVal, int64(5)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
	})
}

func Call_Main_checkEffectOrder(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Str(""))
		_ = __local_var_1_0
		trace_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = trace_2_1
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
					return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString(seen_4.StrVal(), Data_Show_ShowIntImpl(value_3.IntVal)), ","))
				}), trace_2_1)
				_ = __local_var_4_3
				_dollar___unused_5_4 := gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Value{})
				_ = _dollar___unused_5_4
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.IntMul(value_3.IntVal, value_3.IntVal), int64(1)))
			})
		}
		step_3_2 = gopurs_runtime.Func(func(value_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Main_step_3_2(value_3_loop_val)
		})
		saturated_4_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Main_repeatEffects(count_0), step_3_2, gopurs_runtime.Int(int64(1))), gopurs_runtime.Value{})
		_ = saturated_4_5
		saturatedTrace_5_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), trace_2_1), gopurs_runtime.Value{})
		_ = saturatedTrace_5_6
		_dollar___unused_6_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{saturated_4_5.IntVal, int64(26)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_6_7
		_dollar___unused_7_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{saturatedTrace_5_6.StrVal(), "1,2,5,"}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_7_8
		_dollar___unused_8_9 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Str(""), trace_2_1), gopurs_runtime.Value{})
		_ = _dollar___unused_8_9
		partialRef_9_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(Call_Main_repeatEffects(count_0), step_3_2)), gopurs_runtime.Value{})
		_ = partialRef_9_10
		partial_10_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), partialRef_9_10), gopurs_runtime.Value{})
		_ = partial_10_11
		before_11_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), trace_2_1), gopurs_runtime.Value{})
		_ = before_11_12
		_dollar___unused_12_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{before_11_12.StrVal(), ""}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_12_13
		pendingRef_13_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(partial_10_11, gopurs_runtime.Int(int64(1)))), gopurs_runtime.Value{})
		_ = pendingRef_13_14
		pending_14_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), pendingRef_13_14), gopurs_runtime.Value{})
		_ = pending_14_15
		beforeRun_15_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), trace_2_1), gopurs_runtime.Value{})
		_ = beforeRun_15_16
		_dollar___unused_16_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{beforeRun_15_16.StrVal(), ""}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_16_17
		first_17_18 := gopurs_runtime.Apply(pending_14_15, gopurs_runtime.Value{})
		_ = first_17_18
		firstTrace_18_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), trace_2_1), gopurs_runtime.Value{})
		_ = firstTrace_18_19
		_dollar___unused_19_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{first_17_18.IntVal, int64(26)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_19_20
		_dollar___unused_20_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{firstTrace_18_19.StrVal(), "1,2,5,"}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_20_21
		second_21_22 := gopurs_runtime.Apply(gopurs_runtime.Apply(partial_10_11, gopurs_runtime.Int(int64(2))), gopurs_runtime.Value{})
		_ = second_21_22
		secondTrace_22_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), trace_2_1), gopurs_runtime.Value{})
		_ = secondTrace_22_23
		_dollar___unused_23_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{second_21_22.IntVal, int64(677)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_23_24
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   string
				expected string
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{secondTrace_22_23.StrVal(), "1,2,5,2,5,26,"}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
		}()), gopurs_runtime.Value{})
	})
}

func Call_Main_checkEffectBoundary(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(10)))
		_ = __local_var_1_0
		state_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = state_2_1
		makerRef_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), Call_Main_makeEffectBoundary(count_0, state_2_1)), gopurs_runtime.Value{})
		_ = makerRef_3_2
		maker_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), makerRef_3_2), gopurs_runtime.Value{})
		_ = maker_4_3
		pendingRef_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(maker_4_3, gopurs_runtime.Int(int64(1)))), gopurs_runtime.Value{})
		_ = pendingRef_5_4
		pending_6_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), pendingRef_5_4), gopurs_runtime.Value{})
		_ = pending_6_5
		beforeRun_7_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), state_2_1), gopurs_runtime.Value{})
		_ = beforeRun_7_6
		_dollar___unused_8_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{beforeRun_7_6.IntVal, int64(10)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_8_7
		_dollar___unused_9_8 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Int(int64(20)), state_2_1), gopurs_runtime.Value{})
		_ = _dollar___unused_9_8
		first_10_9 := gopurs_runtime.Apply(pending_6_5, gopurs_runtime.Value{})
		_ = first_10_9
		afterRun_11_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), state_2_1), gopurs_runtime.Value{})
		_ = afterRun_11_10
		_dollar___unused_12_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{afterRun_11_10.IntVal, int64(21)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_12_11
		_dollar___unused_13_12 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Int(int64(100)), state_2_1), gopurs_runtime.Value{})
		_ = _dollar___unused_13_12
		_dollar___unused_14_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(first_10_9, gopurs_runtime.Int(int64(2))).IntVal, int64(26)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_14_13
		_dollar___unused_15_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(first_10_9, gopurs_runtime.Int(int64(5))).IntVal, int64(29)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_15_14
		afterCalls_16_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), state_2_1), gopurs_runtime.Value{})
		_ = afterCalls_16_15
		_dollar___unused_17_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{afterCalls_16_15.IntVal, int64(100)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_17_16
		second_18_17 := gopurs_runtime.Apply(pending_6_5, gopurs_runtime.Value{})
		_ = second_18_17
		afterSecondRun_19_18 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), state_2_1), gopurs_runtime.Value{})
		_ = afterSecondRun_19_18
		_dollar___unused_20_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{afterSecondRun_19_18.IntVal, int64(101)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_20_19
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(second_18_17, gopurs_runtime.Int(int64(2))).IntVal, int64(106)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
	})
}

func Call_Main_checkBoundaries(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), Call_Main_makeLetBoundary(count_0))
		_ = __local_var_1_0
		letRef_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = letRef_2_1
		withLet_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), letRef_2_1), gopurs_runtime.Value{})
		_ = withLet_3_2
		cachedRef_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(withLet_3_2, gopurs_runtime.Int(int64(1)))), gopurs_runtime.Value{})
		_ = cachedRef_4_3
		cached_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), cachedRef_4_3), gopurs_runtime.Value{})
		_ = cached_5_4
		_dollar___unused_6_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(cached_5_4, gopurs_runtime.Int(int64(2))).IntVal, int64(54)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_6_5
		_dollar___unused_7_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(cached_5_4, gopurs_runtime.Int(int64(7))).IntVal, int64(59)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_7_6
		_dollar___unused_8_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply2(withLet_3_2, gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(-1))).IntVal, int64(1353)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_8_7
		branchRef_9_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), Call_Main_makeBranchBoundary(count_0)), gopurs_runtime.Value{})
		_ = branchRef_9_8
		branch_10_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), branchRef_9_8), gopurs_runtime.Value{})
		_ = branch_10_9
		negativeRef_11_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(branch_10_9, gopurs_runtime.Int(int64(-2)))), gopurs_runtime.Value{})
		_ = negativeRef_11_10
		negative_12_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), negativeRef_11_10), gopurs_runtime.Value{})
		_ = negative_12_11
		_dollar___unused_13_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(negative_12_11, gopurs_runtime.Int(int64(5))).IntVal, int64(-4)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_13_12
		_dollar___unused_14_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(negative_12_11, gopurs_runtime.Int(int64(0))).IntVal, int64(1)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_14_13
		_dollar___unused_15_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply2(branch_10_9, gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(5))).IntVal, int64(15)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_15_14
		recursiveRef_16_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), Call_Main_makeRecursiveBoundary(count_0)), gopurs_runtime.Value{})
		_ = recursiveRef_16_15
		recursive_17_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), recursiveRef_16_15), gopurs_runtime.Value{})
		_ = recursive_17_16
		visitRef_18_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(recursive_17_16, gopurs_runtime.Int(int64(7)))), gopurs_runtime.Value{})
		_ = visitRef_18_17
		visit_19_18 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), visitRef_18_17), gopurs_runtime.Value{})
		_ = visit_19_18
		_dollar___unused_20_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(visit_19_18, gopurs_runtime.Int(int64(2))).IntVal, int64(25)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_20_19
		_dollar___unused_21_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(visit_19_18, gopurs_runtime.Int(int64(0))).IntVal, int64(21)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_21_20
		factoryRef_22_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Func2(func(a_22 gopurs_runtime.Value, b_23 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.IntMul(int64(100), a_22.IntVal), b_23.IntVal))
		})), gopurs_runtime.Value{})
		_ = factoryRef_22_21
		factory_23_22 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), factoryRef_22_21), gopurs_runtime.Value{})
		_ = factory_23_22
		callRef_24_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), Call_Main_makeCallBoundary(count_0, factory_23_22)), gopurs_runtime.Value{})
		_ = callRef_24_23
		call_25_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), callRef_24_23), gopurs_runtime.Value{})
		_ = call_25_24
		calledRef_26_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(call_25_24, gopurs_runtime.Int(int64(2)))), gopurs_runtime.Value{})
		_ = calledRef_26_25
		called_27_26 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), calledRef_26_25), gopurs_runtime.Value{})
		_ = called_27_26
		_dollar___unused_28_27 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(called_27_26, gopurs_runtime.Int(int64(7))).IntVal, int64(507)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_28_27
		_dollar___unused_29_28 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(called_27_26, gopurs_runtime.Int(int64(-1))).IntVal, int64(499)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_29_28
		uncurriedRef_30_29 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), Call_Main_makeUncurriedBoundary(count_0)), gopurs_runtime.Value{})
		_ = uncurriedRef_30_29
		uncurried_31_30 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), uncurriedRef_30_29), gopurs_runtime.Value{})
		_ = uncurried_31_30
		pairRef_32_31 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(uncurried_31_30, gopurs_runtime.Int(int64(1)))), gopurs_runtime.Value{})
		_ = pairRef_32_31
		pair_33_32 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), pairRef_32_31), gopurs_runtime.Value{})
		_ = pair_33_32
		_dollar___unused_34_33 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.UncurriedApp2(pair_33_32, gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3))).IntVal, int64(324)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_34_33
		_dollar___unused_35_34 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.UncurriedApp2(pair_33_32, gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(2))).IntVal, int64(234)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_35_34
		_dollar___unused_36_35 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.UncurriedApp2(gopurs_runtime.Apply(uncurried_31_30, gopurs_runtime.Int(int64(1))), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(-2))).IntVal, int64(-176)}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_36_35
		typeAppRef_37_36 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), Call_Main_makeTypeAppBoundary(count_0)), gopurs_runtime.Value{})
		_ = typeAppRef_37_36
		typeApp_38_37 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), typeAppRef_37_36), gopurs_runtime.Value{})
		_ = typeApp_38_37
		recordRef_39_38 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(typeApp_38_37, gopurs_runtime.Int(int64(7)))), gopurs_runtime.Value{})
		_ = recordRef_39_38
		record_40_39 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), recordRef_39_38), gopurs_runtime.Value{})
		_ = record_40_39
		_dollar___unused_41_40 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_3959995044_3790796878(Rebox_Main_3790796878_3959995044(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_0()))), Rebox_Main_1739864324_1386611502(Rebox_Main_1386611502_1739864324(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_1())))), func() gopurs_runtime.Value {
			orig := struct {
				actual struct {
					left  int64
					right int64
					/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
				}
				expected struct {
					left  int64
					right int64
					/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() struct {
				left  int64
				right int64
				/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
			} {
				orig := gopurs_runtime.Apply(record_40_39, gopurs_runtime.Int(int64(2)))
				_ = orig
				clone := struct {
					left  int64
					right int64
					/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
				}{}
				clone.left = gopurs_runtime.RecordGet(orig, "left").IntVal
				clone.right = gopurs_runtime.RecordGet(orig, "right").IntVal
				return clone
			}(), struct {
				left  int64
				right int64
				/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
			}{int64(10), int64(2)}}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
				orig := orig.actual
				_ = orig
				return gopurs_runtime.RecordDict2("left", "right", gopurs_runtime.Int(orig.left), gopurs_runtime.Int(orig.right))
			}(), func() gopurs_runtime.Value {
				orig := orig.expected
				_ = orig
				return gopurs_runtime.RecordDict2("left", "right", gopurs_runtime.Int(orig.left), gopurs_runtime.Int(orig.right))
			}())
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_41_40
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_3959995044_3790796878(Rebox_Main_3790796878_3959995044(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_2()))), Rebox_Main_1739864324_1386611502(Rebox_Main_1386611502_1739864324(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_3())))), func() gopurs_runtime.Value {
			orig := struct {
				actual struct {
					left  int64
					right int64
					/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
				}
				expected struct {
					left  int64
					right int64
					/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
				}
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{func() struct {
				left  int64
				right int64
				/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
			} {
				orig := gopurs_runtime.Apply(record_40_39, gopurs_runtime.Int(int64(-5)))
				_ = orig
				clone := struct {
					left  int64
					right int64
					/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
				}{}
				clone.left = gopurs_runtime.RecordGet(orig, "left").IntVal
				clone.right = gopurs_runtime.RecordGet(orig, "right").IntVal
				return clone
			}(), struct {
				left  int64
				right int64
				/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
			}{int64(10), int64(-5)}}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
				orig := orig.actual
				_ = orig
				return gopurs_runtime.RecordDict2("left", "right", gopurs_runtime.Int(orig.left), gopurs_runtime.Int(orig.right))
			}(), func() gopurs_runtime.Value {
				orig := orig.expected
				_ = orig
				return gopurs_runtime.RecordDict2("left", "right", gopurs_runtime.Int(orig.left), gopurs_runtime.Int(orig.right))
			}())
		}()), gopurs_runtime.Value{})
	})
}

func Call_Main_checkArities(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), Call_Main_makeThree(count_0))
		_ = __local_var_1_0
		threeRef_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = threeRef_2_1
		three_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), threeRef_2_1), gopurs_runtime.Value{})
		_ = three_3_2
		_dollar___unused_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply3(three_3_2, gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3))).IntVal, gopurs_runtime.IntAdd(count_0, int64(321))}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_4_3
		threePartialRef_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply(three_3_2, gopurs_runtime.Int(int64(4)))), gopurs_runtime.Value{})
		_ = threePartialRef_5_4
		threePartial_6_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), threePartialRef_5_4), gopurs_runtime.Value{})
		_ = threePartial_6_5
		_dollar___unused_7_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply2(threePartial_6_5, gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(6))).IntVal, gopurs_runtime.IntAdd(count_0, int64(654))}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_7_6
		_dollar___unused_8_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply2(threePartial_6_5, gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(0))).IntVal, gopurs_runtime.IntAdd(count_0, int64(34))}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_8_7
		sixRef_9_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), Call_Main_makeSix(count_0)), gopurs_runtime.Value{})
		_ = sixRef_9_8
		six_10_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), sixRef_9_8), gopurs_runtime.Value{})
		_ = six_10_9
		_dollar___unused_11_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply6(six_10_9, gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(6))).IntVal, gopurs_runtime.IntAdd(count_0, int64(654321))}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_11_10
		twoRef_12_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply2(six_10_9, gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2)))), gopurs_runtime.Value{})
		_ = twoRef_12_11
		two_13_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), twoRef_12_11), gopurs_runtime.Value{})
		_ = two_13_12
		_dollar___unused_14_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply4(two_13_12, gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(6))).IntVal, gopurs_runtime.IntAdd(count_0, int64(654321))}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_14_13
		_dollar___unused_15_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply4(two_13_12, gopurs_runtime.Int(int64(6)), gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(3))).IntVal, gopurs_runtime.IntAdd(count_0, int64(345621))}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_15_14
		fiveRef_16_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Apply5(six_10_9, gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(5)))), gopurs_runtime.Value{})
		_ = fiveRef_16_15
		five_17_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), fiveRef_16_15), gopurs_runtime.Value{})
		_ = five_17_16
		_dollar___unused_18_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(five_17_16, gopurs_runtime.Int(int64(6))).IntVal, gopurs_runtime.IntAdd(count_0, int64(654321))}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_18_17
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
			orig := struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{gopurs_runtime.Apply(five_17_16, gopurs_runtime.Int(int64(9))).IntVal, gopurs_runtime.IntAdd(count_0, int64(954321))}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
		}()), gopurs_runtime.Value{})
	})
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[string])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1739864324(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	left  int64
	right int64
	/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[struct {
		left  int64
		right int64
		/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1739864324_1386611502(in *Constructor_Data_Show_Show[struct {
	left  int64
	right int64
	/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[string])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_3959995044(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	left  int64
	right int64
	/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[struct {
		left  int64
		right int64
		/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_3959995044_3790796878(in *Constructor_Data_Eq_Eq[struct {
	left  int64
	right int64
	/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}
