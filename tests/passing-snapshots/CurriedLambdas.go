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

var cache_Main_eqRec gopurs_runtime.Value
var once_Main_eqRec sync.Once

func Get_Main_eqRec() gopurs_runtime.Value {
	once_Main_eqRec.Do(func() {
		cache_Main_eqRec = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3959995044_3790796878(Rebox_Main_3790796878_3959995044(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_rightIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_leftIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}))))))}
	})
	return cache_Main_eqRec
}

var cache_Main_showRecord gopurs_runtime.Value
var once_Main_showRecord sync.Once

func Get_Main_showRecord() gopurs_runtime.Value {
	once_Main_showRecord.Do(func() {
		cache_Main_showRecord = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1739864324_1386611502(Rebox_Main_1386611502_1739864324(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_leftIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_rightIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}))))))}
	})
	return cache_Main_showRecord
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
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkRepeat(int64(0), __local_var_6_6.IntVal, __local_var_5_5.IntVal), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkRepeat(int64(1), __local_var_6_6.IntVal, __local_var_5_5.IntVal), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkRepeat(int64(2), __local_var_6_6.IntVal, __local_var_5_5.IntVal), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkRepeat(__local_var_4_4.IntVal, __local_var_6_6.IntVal, __local_var_5_5.IntVal), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkRepeat(__local_var_4_4.IntVal, int64(0), __local_var_5_5.IntVal), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkRepeat(__local_var_4_4.IntVal, int64(2), int64(0)), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkRepeat(int64(3), int64(2), int64(-4)), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
											actual   int64
											expected int64
										}{gopurs_runtime.Apply2(Call_Main_composeRepeats(__local_var_4_4.IntVal, int64(2)), gopurs_runtime.Func(func(value_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Int((value_14.IntVal) + (__local_var_6_6.IntVal))
										}), gopurs_runtime.Int(__local_var_5_5.IntVal)).IntVal, (__local_var_5_5.IntVal) + (((__local_var_4_4.IntVal) + (int64(2))) * (__local_var_6_6.IntVal))}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
												actual   int64
												expected int64
											}{gopurs_runtime.Apply2(Call_Main_composeRepeats(int64(0), int64(0)), gopurs_runtime.Func(func(value_15 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Int((value_15.IntVal) + (__local_var_6_6.IntVal))
											}), gopurs_runtime.Int(__local_var_5_5.IntVal)).IntVal, __local_var_5_5.IntVal}), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													// TAST (Let): __local_var_16_7 shape=App(Var) bindingType=Any
													__local_var_16_7 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(3)))
													_ = __local_var_16_7
													__local_var_17_8 := gopurs_runtime.Apply(__local_var_16_7, gopurs_runtime.Value{})
													_ = __local_var_17_8
													__local_var_18_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_17_8), gopurs_runtime.Value{})
													_ = __local_var_18_9
													return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkPartialApplications(__local_var_18_9.IntVal), gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkEffectOrder(__local_var_18_9.IntVal), gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkArities(__local_var_18_9.IntVal), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkBoundaries(__local_var_18_9.IntVal), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkEffectBoundary(__local_var_18_9.IntVal), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
																	}))
																}))
															}))
														}))
													})), gopurs_runtime.Value{})
												})
											}))
										}))
									}))
								}))
							}))
						}))
					}))
				}))
			})), gopurs_runtime.Value{})
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

func Call_Main_makeUncurriedBoundary(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous__4066693242_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous__4066693242_1_0 := Call_Main_repeatApply(count_0)
	_ = previous__4066693242_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.Apply2(previous__4066693242_1_0, gopurs_runtime.Func(func(value_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((value_5.IntVal) + (int64(1)))
			}), gopurs_runtime.Int(((a_2.IntVal)+((int64(10))*(b_3.IntVal)))+((int64(100))*(__local_var_4.IntVal)))).IntVal)
		})
	})
}

func Call_Main_makeTypeAppBoundary(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous__4066693242_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous__4066693242_1_0 := Call_Main_repeatApply(count_0)
	_ = previous__4066693242_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := struct {
					left  int64
					right gopurs_runtime.Value
				}{gopurs_runtime.Apply2(previous__4066693242_1_0, gopurs_runtime.Func(func(value_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int((value_4.IntVal) + (int64(1)))
				}), gopurs_runtime.Int(a_2.IntVal)).IntVal, b_3}
				_ = orig
				return gopurs_runtime.RecordDict2("left", "right", gopurs_runtime.Int(orig.left), orig.right)
			}()
		})
	})
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

func Call_Main_makeRecursiveBoundary(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous__4066693242_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous__4066693242_1_0 := Call_Main_repeatApply(count_0)
	_ = previous__4066693242_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var visit__3466805691_3_1_0 gopurs_runtime.Value
		_ = visit__3466805691_3_1_0
		var visit__3466805691_3_1_0_cell *gopurs_runtime.Value
		_ = visit__3466805691_3_1_0_cell
		// FALLBACK TCO: isLoop=false len=2
		var visit_3_2_1 gopurs_runtime.Value
		_ = visit_3_2_1
		var visit_3_2_1_cell *gopurs_runtime.Value
		_ = visit_3_2_1_cell
		// FALLBACK TCO: isLoop=false len=2
		visit__3466805691_3_1_0 = gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 int64
			{
				if (n_4.IntVal) == (int64(0)) {
					__t3 = gopurs_runtime.Apply2(previous__4066693242_1_0, gopurs_runtime.Func(func(value_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int((value_5.IntVal) + (int64(1)))
					}), gopurs_runtime.Int(a_2.IntVal)).IntVal
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = (gopurs_runtime.Apply((*visit__3466805691_3_1_0_cell), gopurs_runtime.Int((n_4.IntVal)-(int64(1)))).IntVal) + (int64(1))
			}
		end_branch_3:
			return gopurs_runtime.Int(__t3)
		})
		visit__3466805691_3_1_0_cell = &visit__3466805691_3_1_0
		visit_3_2_1 = gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 int64
			{
				if (n_4.IntVal) == (int64(0)) {
					__t4 = gopurs_runtime.Apply2(previous__4066693242_1_0, gopurs_runtime.Func(func(value_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int((value_5.IntVal) + (int64(1)))
					}), gopurs_runtime.Int(a_2.IntVal)).IntVal
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = (gopurs_runtime.Apply((*visit__3466805691_3_1_0_cell), gopurs_runtime.Int((n_4.IntVal)-(int64(1)))).IntVal) + (int64(1))
			}
		end_branch_4:
			return gopurs_runtime.Int(__t4)
		})
		visit_3_2_1_cell = &visit_3_2_1
		return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int((gopurs_runtime.Apply(visit__3466805691_3_1_0, gopurs_runtime.Int(b_4.IntVal)).IntVal) + (gopurs_runtime.Apply(visit__3466805691_3_1_0, gopurs_runtime.Int((b_4.IntVal)+(int64(1)))).IntVal))
		})
	})
}

func Call_Main_makeLetBoundary(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous__4066693242_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous__4066693242_1_0 := Call_Main_repeatApply(count_0)
	_ = previous__4066693242_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): cached_3_1 shape=App(Other) bindingType=Int
		cached_3_1 := gopurs_runtime.Apply2(previous__4066693242_1_0, Get_Main_squarePlusOne(), gopurs_runtime.Int(a_2.IntVal)).IntVal
		_ = cached_3_1
		return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(((cached_3_1) + (cached_3_1)) + (b_4.IntVal))
		})
	})
}

func Call_Main_makeEffectBoundary(count_0_loop int64, state_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	var state_1 gopurs_runtime.Value = state_1_loop
	_ = state_1
	// TAST (Let): previous__4066693242_2_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous__4066693242_2_0 := Call_Main_repeatApply(count_0)
	_ = previous__4066693242_2_0
	return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
			__local_var_4_1 := gopurs_runtime.Apply(Get_Effect_Ref_read(), state_1)
			_ = __local_var_4_1
			__local_var_5_2 := gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Value{})
			_ = __local_var_5_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Int((__local_var_5_2.IntVal)+(int64(1))), state_1), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int((gopurs_runtime.Apply2(previous__4066693242_2_0, gopurs_runtime.Func(func(value_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((value_8.IntVal) + (int64(1)))
						}), gopurs_runtime.Int((a_3.IntVal)+(__local_var_5_2.IntVal))).IntVal) + (b_7.IntVal))
					})
				})
			})), gopurs_runtime.Value{})
		})
	})
}

func Call_Main_makeCallBoundary(count_0_loop int64, factory_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	var factory_1 gopurs_runtime.Value = factory_1_loop
	_ = factory_1
	// TAST (Let): previous__4066693242_2_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous__4066693242_2_0 := Call_Main_repeatApply(count_0)
	_ = previous__4066693242_2_0
	return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(factory_1, gopurs_runtime.Int(gopurs_runtime.Apply2(previous__4066693242_2_0, gopurs_runtime.Func(func(value_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int((value_4.IntVal) + (int64(1)))
		}), gopurs_runtime.Int(a_3.IntVal)).IntVal))
	})
}

func Call_Main_makeBranchBoundary(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	// TAST (Let): previous__4066693242_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
	previous__4066693242_1_0 := Call_Main_repeatApply(count_0)
	_ = previous__4066693242_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 gopurs_runtime.Value
		{
			if (a_2.IntVal) < (int64(0)) {
				__t1 = gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(gopurs_runtime.Apply2(previous__4066693242_1_0, gopurs_runtime.Func(func(value_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int((value_4.IntVal) + (int64(1)))
					}), gopurs_runtime.Int((a_2.IntVal)-(b_3.IntVal))).IntVal)
				})
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.Apply2(previous__4066693242_1_0, gopurs_runtime.Func(func(value_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int((value_4.IntVal) + (int64(1)))
				}), gopurs_runtime.Int((a_2.IntVal)+((int64(2))*(b_3.IntVal)))).IntVal)
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
	return Call_Test_Assert_assertEqual_prime___627669702("", struct {
		actual   int64
		expected int64
	}{gopurs_runtime.Apply2(Call_Main_repeatApply(count_0), gopurs_runtime.Func(func(value_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int((value_3.IntVal) + (step_1))
	}), gopurs_runtime.Int(seed_2)).IntVal, (seed_2) + ((count_0) * (step_1))})
}

func Call_Main_checkPartialApplications(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
		actual   int64
		expected int64
	}{gopurs_runtime.Apply2(Call_Main_repeatApply(count_0), Get_Main_squarePlusOne(), gopurs_runtime.Int(int64(1))).IntVal, int64(26)}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
			actual   int64
			expected int64
		}{gopurs_runtime.Apply2(Call_Main_repeatApply(count_0), Get_Main_squarePlusOne(), gopurs_runtime.Int(int64(2))).IntVal, int64(677)}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Call_Main_repeatApply(count_0), Get_Main_squarePlusOne())), gopurs_runtime.Value{})
				_ = __local_var_3_0
				__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_3_0), gopurs_runtime.Value{})
				_ = __local_var_4_1
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
					actual   int64
					expected int64
				}{gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Int(int64(1))).IntVal, int64(26)}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
						actual   int64
						expected int64
					}{gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Int(int64(2))).IntVal, int64(677)}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_Test_Assert_assertEqual_prime___627669702("", struct {
							actual   int64
							expected int64
						}{gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Int(int64(0))).IntVal, int64(5)})
					}))
				})), gopurs_runtime.Value{})
			})
		}))
	}))
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
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(seen_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str(((seen_4.StrVal()) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), value_3).StrVal())) + (","))
			}), __local_var_2_1), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_5_3 shape=Other bindingType=Int
					__local_var_5_3 := ((value_3.IntVal) * (value_3.IntVal)) + (int64(1))
					_ = __local_var_5_3
					return gopurs_runtime.Int(__local_var_5_3)
				})
			}))
		}
		step_3_2 = gopurs_runtime.Func(func(value_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Main_step_3_2(value_3_loop_val)
		})
		__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Main_repeatEffects(count_0), step_3_2, gopurs_runtime.Int(int64(1))), gopurs_runtime.Value{})
		_ = __local_var_4_4
		__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
		_ = __local_var_5_5
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
			actual   int64
			expected int64
		}{__local_var_4_4.IntVal, int64(26)}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
				actual   string
				expected string
			}{__local_var_5_5.StrVal(), "1,2,5,"}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Str(""), __local_var_2_1), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						__local_var_9_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Call_Main_repeatEffects(count_0), step_3_2)), gopurs_runtime.Value{})
						_ = __local_var_9_6
						__local_var_10_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_9_6), gopurs_runtime.Value{})
						_ = __local_var_10_7
						__local_var_11_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
						_ = __local_var_11_8
						return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
							actual   string
							expected string
						}{__local_var_11_8.StrVal(), ""}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								__local_var_13_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_10_7, gopurs_runtime.Int(int64(1)))), gopurs_runtime.Value{})
								_ = __local_var_13_9
								__local_var_14_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_13_9), gopurs_runtime.Value{})
								_ = __local_var_14_10
								__local_var_15_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
								_ = __local_var_15_11
								return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
									actual   string
									expected string
								}{__local_var_15_11.StrVal(), ""}), gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										__local_var_17_12 := gopurs_runtime.Apply(__local_var_14_10, gopurs_runtime.Value{})
										_ = __local_var_17_12
										__local_var_18_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
										_ = __local_var_18_13
										return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
											actual   int64
											expected int64
										}{__local_var_17_12.IntVal, int64(26)}), gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
												actual   string
												expected string
											}{__local_var_18_13.StrVal(), "1,2,5,"}), gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													// TAST (Let): __local_var_21_14 shape=App(Other) bindingType=(ADT ["Effect","Effect"] [Int])
													__local_var_21_14 := gopurs_runtime.Apply(__local_var_10_7, gopurs_runtime.Int(int64(2)))
													_ = __local_var_21_14
													__local_var_22_15 := gopurs_runtime.Apply(__local_var_21_14, gopurs_runtime.Value{})
													_ = __local_var_22_15
													__local_var_23_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
													_ = __local_var_23_16
													return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
														actual   int64
														expected int64
													}{__local_var_22_15.IntVal, int64(677)}), gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
														return Call_Test_Assert_assertEqual_prime___4176622598("", struct {
															actual   string
															expected string
														}{__local_var_23_16.StrVal(), "1,2,5,2,5,26,"})
													})), gopurs_runtime.Value{})
												})
											}))
										})), gopurs_runtime.Value{})
									})
								})), gopurs_runtime.Value{})
							})
						})), gopurs_runtime.Value{})
					})
				}))
			}))
		})), gopurs_runtime.Value{})
	})
}

func Call_Main_checkEffectBoundary(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(10)))
		_ = __local_var_1_0
		__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = __local_var_2_1
		__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Call_Main_makeEffectBoundary(count_0, __local_var_2_1)), gopurs_runtime.Value{})
		_ = __local_var_3_2
		__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_3_2), gopurs_runtime.Value{})
		_ = __local_var_4_3
		__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Int(int64(1)))), gopurs_runtime.Value{})
		_ = __local_var_5_4
		__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_4), gopurs_runtime.Value{})
		_ = __local_var_6_5
		__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
		_ = __local_var_7_6
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
			actual   int64
			expected int64
		}{__local_var_7_6.IntVal, int64(10)}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Int(int64(20)), __local_var_2_1), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					__local_var_10_7 := gopurs_runtime.Apply(__local_var_6_5, gopurs_runtime.Value{})
					_ = __local_var_10_7
					__local_var_11_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
					_ = __local_var_11_8
					return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
						actual   int64
						expected int64
					}{__local_var_11_8.IntVal, int64(21)}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Int(int64(100)), __local_var_2_1), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
								actual   int64
								expected int64
							}{gopurs_runtime.Apply(__local_var_10_7, gopurs_runtime.Int(int64(2))).IntVal, int64(26)}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
									actual   int64
									expected int64
								}{gopurs_runtime.Apply(__local_var_10_7, gopurs_runtime.Int(int64(5))).IntVal, int64(29)}), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										// TAST (Let): __local_var_16_9 shape=App(Var) bindingType=Any
										__local_var_16_9 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1)
										_ = __local_var_16_9
										__local_var_17_10 := gopurs_runtime.Apply(__local_var_16_9, gopurs_runtime.Value{})
										_ = __local_var_17_10
										return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
											actual   int64
											expected int64
										}{__local_var_17_10.IntVal, int64(100)}), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												__local_var_19_11 := gopurs_runtime.Apply(__local_var_6_5, gopurs_runtime.Value{})
												_ = __local_var_19_11
												__local_var_20_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_1), gopurs_runtime.Value{})
												_ = __local_var_20_12
												return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
													actual   int64
													expected int64
												}{__local_var_20_12.IntVal, int64(101)}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
													return Call_Test_Assert_assertEqual_prime___627669702("", struct {
														actual   int64
														expected int64
													}{gopurs_runtime.Apply(__local_var_19_11, gopurs_runtime.Int(int64(2))).IntVal, int64(106)})
												})), gopurs_runtime.Value{})
											})
										})), gopurs_runtime.Value{})
									})
								}))
							}))
						}))
					})), gopurs_runtime.Value{})
				})
			}))
		})), gopurs_runtime.Value{})
	})
}

func Call_Main_checkBoundaries(count_0_loop int64) gopurs_runtime.Value {
	var count_0 int64 = count_0_loop
	_ = count_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Call_Main_makeLetBoundary(count_0)), gopurs_runtime.Value{})
		_ = __local_var_1_0
		__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_0), gopurs_runtime.Value{})
		_ = __local_var_2_1
		__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Int(int64(1)))), gopurs_runtime.Value{})
		_ = __local_var_3_2
		__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_3_2), gopurs_runtime.Value{})
		_ = __local_var_4_3
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
			actual   int64
			expected int64
		}{gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Int(int64(2))).IntVal, int64(54)}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Int(int64(7))).IntVal, int64(59)}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
					actual   int64
					expected int64
				}{gopurs_runtime.Apply2(__local_var_2_1, gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(-1))).IntVal, int64(1353)}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						__local_var_8_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Call_Main_makeBranchBoundary(count_0)), gopurs_runtime.Value{})
						_ = __local_var_8_4
						__local_var_9_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_8_4), gopurs_runtime.Value{})
						_ = __local_var_9_5
						__local_var_10_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_9_5, gopurs_runtime.Int(int64(-2)))), gopurs_runtime.Value{})
						_ = __local_var_10_6
						__local_var_11_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_10_6), gopurs_runtime.Value{})
						_ = __local_var_11_7
						return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
							actual   int64
							expected int64
						}{gopurs_runtime.Apply(__local_var_11_7, gopurs_runtime.Int(int64(5))).IntVal, int64(-4)}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
								actual   int64
								expected int64
							}{gopurs_runtime.Apply(__local_var_11_7, gopurs_runtime.Int(int64(0))).IntVal, int64(1)}), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
									actual   int64
									expected int64
								}{gopurs_runtime.Apply2(__local_var_9_5, gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(5))).IntVal, int64(15)}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										__local_var_15_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Call_Main_makeRecursiveBoundary(count_0)), gopurs_runtime.Value{})
										_ = __local_var_15_8
										__local_var_16_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_15_8), gopurs_runtime.Value{})
										_ = __local_var_16_9
										__local_var_17_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_16_9, gopurs_runtime.Int(int64(7)))), gopurs_runtime.Value{})
										_ = __local_var_17_10
										__local_var_18_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_17_10), gopurs_runtime.Value{})
										_ = __local_var_18_11
										return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
											actual   int64
											expected int64
										}{gopurs_runtime.Apply(__local_var_18_11, gopurs_runtime.Int(int64(2))).IntVal, int64(25)}), gopurs_runtime.Func(func(_dollar___unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
												actual   int64
												expected int64
											}{gopurs_runtime.Apply(__local_var_18_11, gopurs_runtime.Int(int64(0))).IntVal, int64(21)}), gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													__local_var_21_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Func2(func(a_21 gopurs_runtime.Value, b_22 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Int(((int64(100)) * (a_21.IntVal)) + (b_22.IntVal))
													})), gopurs_runtime.Value{})
													_ = __local_var_21_12
													__local_var_22_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_21_12), gopurs_runtime.Value{})
													_ = __local_var_22_13
													__local_var_23_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Call_Main_makeCallBoundary(count_0, __local_var_22_13)), gopurs_runtime.Value{})
													_ = __local_var_23_14
													__local_var_24_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_23_14), gopurs_runtime.Value{})
													_ = __local_var_24_15
													__local_var_25_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_24_15, gopurs_runtime.Int(int64(2)))), gopurs_runtime.Value{})
													_ = __local_var_25_16
													__local_var_26_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_25_16), gopurs_runtime.Value{})
													_ = __local_var_26_17
													return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
														actual   int64
														expected int64
													}{gopurs_runtime.Apply(__local_var_26_17, gopurs_runtime.Int(int64(7))).IntVal, int64(507)}), gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
															actual   int64
															expected int64
														}{gopurs_runtime.Apply(__local_var_26_17, gopurs_runtime.Int(int64(-1))).IntVal, int64(499)}), gopurs_runtime.Func(func(_dollar___unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																__local_var_29_18 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Call_Main_makeUncurriedBoundary(count_0)), gopurs_runtime.Value{})
																_ = __local_var_29_18
																__local_var_30_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_29_18), gopurs_runtime.Value{})
																_ = __local_var_30_19
																__local_var_31_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_30_19, gopurs_runtime.Int(int64(1)))), gopurs_runtime.Value{})
																_ = __local_var_31_20
																__local_var_32_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_31_20), gopurs_runtime.Value{})
																_ = __local_var_32_21
																return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
																	actual   int64
																	expected int64
																}{gopurs_runtime.UncurriedApp2(__local_var_32_21, gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3))).IntVal, int64(324)}), gopurs_runtime.Func(func(_dollar___unused_33 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
																		actual   int64
																		expected int64
																	}{gopurs_runtime.UncurriedApp2(__local_var_32_21, gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(2))).IntVal, int64(234)}), gopurs_runtime.Func(func(_dollar___unused_34 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
																			actual   int64
																			expected int64
																		}{gopurs_runtime.UncurriedApp2(gopurs_runtime.Apply(__local_var_30_19, gopurs_runtime.Int(int64(1))), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(-2))).IntVal, int64(-176)}), gopurs_runtime.Func(func(_dollar___unused_35 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																				__local_var_36_22 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Call_Main_makeTypeAppBoundary(count_0)), gopurs_runtime.Value{})
																				_ = __local_var_36_22
																				__local_var_37_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_36_22), gopurs_runtime.Value{})
																				_ = __local_var_37_23
																				__local_var_38_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_37_23, gopurs_runtime.Int(int64(7)))), gopurs_runtime.Value{})
																				_ = __local_var_38_24
																				__local_var_39_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_38_24), gopurs_runtime.Value{})
																				_ = __local_var_39_25
																				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___101187400("", struct {
																					actual struct {
																						left  int64
																						right int64
																					}
																					expected struct {
																						left  int64
																						right int64
																					}
																				}{func() struct {
																					left  int64
																					right int64
																				} {
																					orig := gopurs_runtime.Apply(__local_var_39_25, gopurs_runtime.Int(int64(2)))
																					_ = orig
																					clone := struct {
																						left  int64
																						right int64
																					}{}
																					clone.left = gopurs_runtime.RecordGet(orig, "left").IntVal
																					clone.right = gopurs_runtime.RecordGet(orig, "right").IntVal
																					return clone
																				}(), struct {
																					left  int64
																					right int64
																				}{int64(10), int64(2)}}), gopurs_runtime.Func(func(_dollar___unused_40 gopurs_runtime.Value) gopurs_runtime.Value {
																					return Call_Test_Assert_assertEqual_prime___101187400("", struct {
																						actual struct {
																							left  int64
																							right int64
																						}
																						expected struct {
																							left  int64
																							right int64
																						}
																					}{func() struct {
																						left  int64
																						right int64
																					} {
																						orig := gopurs_runtime.Apply(__local_var_39_25, gopurs_runtime.Int(int64(-5)))
																						_ = orig
																						clone := struct {
																							left  int64
																							right int64
																						}{}
																						clone.left = gopurs_runtime.RecordGet(orig, "left").IntVal
																						clone.right = gopurs_runtime.RecordGet(orig, "right").IntVal
																						return clone
																					}(), struct {
																						left  int64
																						right int64
																					}{int64(10), int64(-5)}})
																				})), gopurs_runtime.Value{})
																			})
																		}))
																	}))
																})), gopurs_runtime.Value{})
															})
														}))
													})), gopurs_runtime.Value{})
												})
											}))
										})), gopurs_runtime.Value{})
									})
								}))
							}))
						})), gopurs_runtime.Value{})
					})
				}))
			}))
		})), gopurs_runtime.Value{})
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
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
			actual   int64
			expected int64
		}{gopurs_runtime.Apply3(__local_var_2_1, gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3))).IntVal, (count_0) + (int64(321))}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Int(int64(4)))), gopurs_runtime.Value{})
				_ = __local_var_4_2
				__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_2), gopurs_runtime.Value{})
				_ = __local_var_5_3
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
					actual   int64
					expected int64
				}{gopurs_runtime.Apply2(__local_var_5_3, gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(6))).IntVal, (count_0) + (int64(654))}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
						actual   int64
						expected int64
					}{gopurs_runtime.Apply2(__local_var_5_3, gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(0))).IntVal, (count_0) + (int64(34))}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							__local_var_8_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Call_Main_makeSix(count_0)), gopurs_runtime.Value{})
							_ = __local_var_8_4
							__local_var_9_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_8_4), gopurs_runtime.Value{})
							_ = __local_var_9_5
							return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
								actual   int64
								expected int64
							}{gopurs_runtime.Apply6(__local_var_9_5, gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(6))).IntVal, (count_0) + (int64(654321))}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									__local_var_11_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply2(__local_var_9_5, gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2)))), gopurs_runtime.Value{})
									_ = __local_var_11_6
									__local_var_12_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_11_6), gopurs_runtime.Value{})
									_ = __local_var_12_7
									return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
										actual   int64
										expected int64
									}{gopurs_runtime.Apply4(__local_var_12_7, gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(6))).IntVal, (count_0) + (int64(654321))}), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
											actual   int64
											expected int64
										}{gopurs_runtime.Apply4(__local_var_12_7, gopurs_runtime.Int(int64(6)), gopurs_runtime.Int(int64(5)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(3))).IntVal, (count_0) + (int64(345621))}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												__local_var_15_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply5(__local_var_9_5, gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2)), gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(int64(4)), gopurs_runtime.Int(int64(5)))), gopurs_runtime.Value{})
												_ = __local_var_15_8
												__local_var_16_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_15_8), gopurs_runtime.Value{})
												_ = __local_var_16_9
												return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
													actual   int64
													expected int64
												}{gopurs_runtime.Apply(__local_var_16_9, gopurs_runtime.Int(int64(6))).IntVal, (count_0) + (int64(654321))}), gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
													return Call_Test_Assert_assertEqual_prime___627669702("", struct {
														actual   int64
														expected int64
													}{gopurs_runtime.Apply(__local_var_16_9, gopurs_runtime.Int(int64(9))).IntVal, (count_0) + (int64(954321))})
												})), gopurs_runtime.Value{})
											})
										}))
									})), gopurs_runtime.Value{})
								})
							})), gopurs_runtime.Value{})
						})
					}))
				})), gopurs_runtime.Value{})
			})
		})), gopurs_runtime.Value{})
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

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1739864324(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	left  int64
	right int64
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[struct {
		left  int64
		right int64
	}]{}
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

func Rebox_Main_1739864324_1386611502(in *Constructor_Data_Show_Show[struct {
	left  int64
	right int64
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_3959995044(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	left  int64
	right int64
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[struct {
		left  int64
		right int64
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3959995044_3790796878(in *Constructor_Data_Eq_Eq[struct {
	left  int64
	right int64
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
