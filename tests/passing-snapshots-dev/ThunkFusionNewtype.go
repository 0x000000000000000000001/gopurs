package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Lazy gopurs_runtime.Value
var once_Main_Lazy sync.Once

func Get_Main_Lazy() gopurs_runtime.Value {
	once_Main_Lazy.Do(func() {
		cache_Main_Lazy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Lazy(x_0_box)
		})
	})
	return cache_Main_Lazy
}

var cache_Main_force gopurs_runtime.Value
var once_Main_force sync.Once

func Get_Main_force() gopurs_runtime.Value {
	once_Main_force.Do(func() {
		cache_Main_force = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_force(v_0_box)
		})
	})
	return cache_Main_force
}

var cache_Main_go__defer gopurs_runtime.Value
var once_Main_go__defer sync.Once

func Get_Main_go__defer() gopurs_runtime.Value {
	once_Main_go__defer.Do(func() {
		cache_Main_go__defer = Get_Main_Lazy()
	})
	return cache_Main_go__defer
}

var cache_Main_wrapAdds__gopurs_strict_thunk_0 gopurs_runtime.Value
var once_Main_wrapAdds__gopurs_strict_thunk_0 sync.Once

func Get_Main_wrapAdds__gopurs_strict_thunk_0() gopurs_runtime.Value {
	once_Main_wrapAdds__gopurs_strict_thunk_0.Do(func() {
		cache_Main_wrapAdds__gopurs_strict_thunk_0 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_wrapAdds__gopurs_strict_thunk_0(v_0_box.IntVal, v1_1_box.IntVal, v2_2_box.IntVal))
		})
	})
	return cache_Main_wrapAdds__gopurs_strict_thunk_0
}

var cache_Main_wrapAdds gopurs_runtime.Value
var once_Main_wrapAdds sync.Once

func Get_Main_wrapAdds() gopurs_runtime.Value {
	once_Main_wrapAdds.Do(func() {
		cache_Main_wrapAdds = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_wrapAdds(v_0_box.IntVal, v1_1_box.IntVal, v2_2_box)
		})
	})
	return cache_Main_wrapAdds
}

var cache_Main_runAdds gopurs_runtime.Value
var once_Main_runAdds sync.Once

func Get_Main_runAdds() gopurs_runtime.Value {
	once_Main_runAdds.Do(func() {
		cache_Main_runAdds = gopurs_runtime.Func3(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value, step_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runAdds(depth_0_box.IntVal, seed_1_box.IntVal, step_2_box.IntVal))
		})
	})
	return cache_Main_runAdds
}

var cache_Main_wrapEffects gopurs_runtime.Value
var once_Main_wrapEffects sync.Once

func Get_Main_wrapEffects() gopurs_runtime.Value {
	once_Main_wrapEffects.Do(func() {
		cache_Main_wrapEffects = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_wrapEffects(v_0_box.IntVal, v1_1_box)
		})
	})
	return cache_Main_wrapEffects
}

var cache_Main_wrapIncrements__gopurs_strict_thunk_0 gopurs_runtime.Value
var once_Main_wrapIncrements__gopurs_strict_thunk_0 sync.Once

func Get_Main_wrapIncrements__gopurs_strict_thunk_0() gopurs_runtime.Value {
	once_Main_wrapIncrements__gopurs_strict_thunk_0.Do(func() {
		cache_Main_wrapIncrements__gopurs_strict_thunk_0 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_wrapIncrements__gopurs_strict_thunk_0(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Main_wrapIncrements__gopurs_strict_thunk_0
}

var cache_Main_wrapIncrements gopurs_runtime.Value
var once_Main_wrapIncrements sync.Once

func Get_Main_wrapIncrements() gopurs_runtime.Value {
	once_Main_wrapIncrements.Do(func() {
		cache_Main_wrapIncrements = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_wrapIncrements(v_0_box.IntVal, v1_1_box)
		})
	})
	return cache_Main_wrapIncrements
}

var cache_Main_keepWrapped gopurs_runtime.Value
var once_Main_keepWrapped sync.Once

func Get_Main_keepWrapped() gopurs_runtime.Value {
	once_Main_keepWrapped.Do(func() {
		cache_Main_keepWrapped = gopurs_runtime.Func2(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_keepWrapped(depth_0_box.IntVal, seed_1_box.IntVal)
		})
	})
	return cache_Main_keepWrapped
}

var cache_Main_runIncrements gopurs_runtime.Value
var once_Main_runIncrements sync.Once

func Get_Main_runIncrements() gopurs_runtime.Value {
	once_Main_runIncrements.Do(func() {
		cache_Main_runIncrements = gopurs_runtime.Func2(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runIncrements(depth_0_box.IntVal, seed_1_box.IntVal))
		})
	})
	return cache_Main_runIncrements
}

var cache_Main_wrapOrder__gopurs_strict_thunk_0 gopurs_runtime.Value
var once_Main_wrapOrder__gopurs_strict_thunk_0 sync.Once

func Get_Main_wrapOrder__gopurs_strict_thunk_0() gopurs_runtime.Value {
	once_Main_wrapOrder__gopurs_strict_thunk_0.Do(func() {
		cache_Main_wrapOrder__gopurs_strict_thunk_0 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_wrapOrder__gopurs_strict_thunk_0(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Main_wrapOrder__gopurs_strict_thunk_0
}

var cache_Main_wrapOrder gopurs_runtime.Value
var once_Main_wrapOrder sync.Once

func Get_Main_wrapOrder() gopurs_runtime.Value {
	once_Main_wrapOrder.Do(func() {
		cache_Main_wrapOrder = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_wrapOrder(v_0_box.IntVal, v1_1_box)
		})
	})
	return cache_Main_wrapOrder
}

var cache_Main_runOrder gopurs_runtime.Value
var once_Main_runOrder sync.Once

func Get_Main_runOrder() gopurs_runtime.Value {
	once_Main_runOrder.Do(func() {
		cache_Main_runOrder = gopurs_runtime.Func2(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runOrder(depth_0_box.IntVal, seed_1_box.IntVal))
		})
	})
	return cache_Main_runOrder
}

var cache_Main_checkAdds gopurs_runtime.Value
var once_Main_checkAdds sync.Once

func Get_Main_checkAdds() gopurs_runtime.Value {
	once_Main_checkAdds.Do(func() {
		cache_Main_checkAdds = gopurs_runtime.Func3(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value, step_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkAdds(depth_0_box.IntVal, seed_1_box.IntVal, step_2_box.IntVal)
		})
	})
	return cache_Main_checkAdds
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(1000)))
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
			_dollar___unused_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_wrapIncrements__gopurs_strict_thunk_0(int64(0), seed_5_5.IntVal), seed_5_5.IntVal}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			_dollar___unused_8_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_wrapIncrements__gopurs_strict_thunk_0(int64(1), seed_5_5.IntVal), gopurs_runtime.IntAdd(seed_5_5.IntVal, int64(1))}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_8_8
			_dollar___unused_9_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_wrapIncrements__gopurs_strict_thunk_0(depth_4_4.IntVal, seed_5_5.IntVal), gopurs_runtime.IntAdd(seed_5_5.IntVal, depth_4_4.IntVal)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_9_9
			_dollar___unused_10_10 := gopurs_runtime.Apply(Call_Main_checkAdds(int64(0), seed_5_5.IntVal, step_6_6.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_10_10
			_dollar___unused_11_11 := gopurs_runtime.Apply(Call_Main_checkAdds(int64(1), seed_5_5.IntVal, step_6_6.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_11_11
			_dollar___unused_12_12 := gopurs_runtime.Apply(Call_Main_checkAdds(int64(7), seed_5_5.IntVal, step_6_6.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_12_12
			_dollar___unused_13_13 := gopurs_runtime.Apply(Call_Main_checkAdds(depth_4_4.IntVal, seed_5_5.IntVal, step_6_6.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_13_13
			_dollar___unused_14_14 := gopurs_runtime.Apply(Call_Main_checkAdds(int64(7), step_6_6.IntVal, seed_5_5.IntVal), gopurs_runtime.Value{})
			_ = _dollar___unused_14_14
			_dollar___unused_15_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_wrapOrder__gopurs_strict_thunk_0(int64(0), seed_5_5.IntVal), seed_5_5.IntVal}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_15_15
			_dollar___unused_16_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_wrapOrder__gopurs_strict_thunk_0(int64(1), seed_5_5.IntVal), int64(15)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_16_16
			_dollar___unused_17_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_wrapOrder__gopurs_strict_thunk_0(int64(3), seed_5_5.IntVal), int64(73)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_17_17
			escapedRef_18_18 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), Call_Main_wrapIncrements(depth_4_4.IntVal, gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
				return seed_5_5
			}))), gopurs_runtime.Value{})
			_ = escapedRef_18_18
			escaped_19_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), escapedRef_18_18), gopurs_runtime.Value{})
			_ = escaped_19_19
			_dollar___unused_20_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{gopurs_runtime.IntAdd(gopurs_runtime.Apply(escaped_19_19, Get_Data_Unit_unit()).IntVal, gopurs_runtime.Apply(escaped_19_19, Get_Data_Unit_unit()).IntVal), gopurs_runtime.IntMul(int64(2), gopurs_runtime.IntAdd(seed_5_5.IntVal, depth_4_4.IntVal))}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_20_20
			calls_21_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(0))), gopurs_runtime.Value{})
			_ = calls_21_21
			// TAST (Let): pending_22_22 shape=App(Var) bindingType=(Func [Unit] (ADT ["Effect","Effect"] [Int]))
			pending_22_22 := Call_Main_wrapEffects(int64(3), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_23_23 shape=App(Var) bindingType=Any
					__local_var_23_23 := gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(count_23 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int(gopurs_runtime.IntAdd(count_23.IntVal, int64(1)))
					}), calls_21_21)
					_ = __local_var_23_23
					_dollar___unused_24_24 := gopurs_runtime.Apply(__local_var_23_23, gopurs_runtime.Value{})
					_ = _dollar___unused_24_24
					return seed_5_5
				})
			}))
			_ = pending_22_22
			before_23_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), calls_21_21), gopurs_runtime.Value{})
			_ = before_23_25
			_dollar___unused_24_26 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{before_23_25.IntVal, int64(0)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_24_26
			first_25_27 := gopurs_runtime.Apply(gopurs_runtime.Apply(pending_22_22, Get_Data_Unit_unit()), gopurs_runtime.Value{})
			_ = first_25_27
			second_26_28 := gopurs_runtime.Apply(gopurs_runtime.Apply(pending_22_22, Get_Data_Unit_unit()), gopurs_runtime.Value{})
			_ = second_26_28
			after_27_29 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), calls_21_21), gopurs_runtime.Value{})
			_ = after_27_29
			_dollar___unused_28_30 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{first_25_27.IntVal, int64(10)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_28_30
			_dollar___unused_29_31 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{second_26_28.IntVal, int64(10)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_29_31
			_dollar___unused_30_32 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{after_27_29.IntVal, int64(2)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_30_32
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_Lazy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_force(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(v_0, Get_Data_Unit_unit())
}

func Call_Main_wrapAdds__gopurs_strict_thunk_0(v_0_loop int64, v1_1_loop int64, v2_2_loop int64) int64 {
	if ((v_0_loop) == (int64(int32(v_0_loop)))) && ((v2_2_loop) == (int64(int32(v2_2_loop)))) {
		return func() int64 {
			v_0_loop := int32(v_0_loop)
			_ = v_0_loop
			v2_2_loop := int32(v2_2_loop)
			_ = v2_2_loop
		wrapAdds__gopurs_strict_thunk_0:
			for {
				if false {
					continue wrapAdds__gopurs_strict_thunk_0
				}
				v_0 := v_0_loop
				_ = v_0
				var v1_1 int64 = v1_1_loop
				_ = v1_1
				v2_2 := v2_2_loop
				_ = v2_2
				var __t0 int64
				{
					if (v_0) == (int32(0)) {
						__t0 = int64(v2_2)
						goto end_branch_0
					} else {

					}
				}
				{
					v_0_loop = (v_0) - (int32(1))
					v1_1_loop = v1_1
					v2_2_loop = (v2_2) + (int32(gopurs_runtime.IntAdd(v1_1, int64(0))))
					continue wrapAdds__gopurs_strict_thunk_0
					__t0 = func() int64 { panic("unreachable") }()
				}
			end_branch_0:
				return __t0
			}
		}()
	} else {

	}
wrapAdds__gopurs_strict_thunk_0:
	for {
		if false {
			continue wrapAdds__gopurs_strict_thunk_0
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var v2_2 int64 = v2_2_loop
		_ = v2_2
		var __t0 int64
		{
			if (v_0) == (int64(0)) {
				__t0 = v2_2
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
			v1_1_loop = v1_1
			v2_2_loop = gopurs_runtime.IntAdd(v2_2, v1_1)
			continue wrapAdds__gopurs_strict_thunk_0
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_wrapAdds(v_0_loop int64, v1_1_loop int64, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
wrapAdds:
	for {
		if false {
			continue wrapAdds
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var v2_2 gopurs_runtime.Value = v2_2_loop
		_ = v2_2
		var __t0 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t0 = v2_2
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
			v1_1_loop = v1_1
			v2_2_loop = gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.Apply(v2_2, Get_Data_Unit_unit()).IntVal, v1_1))
			})
			continue wrapAdds
			__t0 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_runAdds(depth_0_loop int64, seed_1_loop int64, step_2_loop int64) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	var step_2 int64 = step_2_loop
	_ = step_2
	return Call_Main_wrapAdds__gopurs_strict_thunk_0(depth_0, step_2, seed_1)
}

func Call_Main_wrapEffects(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
wrapEffects:
	for {
		if false {
			continue wrapEffects
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var __t2 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t2 = v1_1
				goto end_branch_2
			} else {

			}
		}
		{
			v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
			v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(ADT ["Effect","Effect"] [Int])
					__local_var_3_0 := gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit())
					_ = __local_var_3_0
					value_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
					_ = value_4_1
					return gopurs_runtime.Int(gopurs_runtime.IntAdd(value_4_1.IntVal, int64(1)))
				})
			})
			continue wrapEffects
			__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_2:
		return __t2
	}
}

func Call_Main_wrapIncrements__gopurs_strict_thunk_0(v_0_loop int64, v1_1_loop int64) int64 {
	if ((v_0_loop) == (int64(int32(v_0_loop)))) && ((v1_1_loop) == (int64(int32(v1_1_loop)))) {
		return func() int64 {
			v_0_loop := int32(v_0_loop)
			_ = v_0_loop
			v1_1_loop := int32(v1_1_loop)
			_ = v1_1_loop
		wrapIncrements__gopurs_strict_thunk_0:
			for {
				if false {
					continue wrapIncrements__gopurs_strict_thunk_0
				}
				v_0 := v_0_loop
				_ = v_0
				v1_1 := v1_1_loop
				_ = v1_1
				var __t0 int64
				{
					if (v_0) == (int32(0)) {
						__t0 = int64(v1_1)
						goto end_branch_0
					} else {

					}
				}
				{
					v_0_loop = (v_0) - (int32(1))
					v1_1_loop = (v1_1) + (int32(1))
					continue wrapIncrements__gopurs_strict_thunk_0
					__t0 = func() int64 { panic("unreachable") }()
				}
			end_branch_0:
				return __t0
			}
		}()
	} else {

	}
wrapIncrements__gopurs_strict_thunk_0:
	for {
		if false {
			continue wrapIncrements__gopurs_strict_thunk_0
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var __t0 int64
		{
			if (v_0) == (int64(0)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
			v1_1_loop = gopurs_runtime.IntAdd(v1_1, int64(1))
			continue wrapIncrements__gopurs_strict_thunk_0
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_wrapIncrements(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
wrapIncrements:
	for {
		if false {
			continue wrapIncrements
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var __t0 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
			v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit()).IntVal, int64(1)))
			})
			continue wrapIncrements
			__t0 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_keepWrapped(depth_0_loop int64, seed_1_loop int64) gopurs_runtime.Value {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	return Call_Main_wrapIncrements(depth_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int(seed_1)
	}))
}

func Call_Main_runIncrements(depth_0_loop int64, seed_1_loop int64) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	return Call_Main_wrapIncrements__gopurs_strict_thunk_0(depth_0, seed_1)
}

func Call_Main_wrapOrder__gopurs_strict_thunk_0(v_0_loop int64, v1_1_loop int64) int64 {
	if ((v_0_loop) == (int64(int32(v_0_loop)))) && ((v1_1_loop) == (int64(int32(v1_1_loop)))) {
		return func() int64 {
			v_0_loop := int32(v_0_loop)
			_ = v_0_loop
			v1_1_loop := int32(v1_1_loop)
			_ = v1_1_loop
		wrapOrder__gopurs_strict_thunk_0:
			for {
				if false {
					continue wrapOrder__gopurs_strict_thunk_0
				}
				v_0 := v_0_loop
				_ = v_0
				v1_1 := v1_1_loop
				_ = v1_1
				var __t0 int64
				{
					if (v_0) == (int32(0)) {
						__t0 = int64(v1_1)
						goto end_branch_0
					} else {

					}
				}
				{
					v_0_loop = (v_0) - (int32(1))
					v1_1_loop = (int32(gopurs_runtime.IntAdd(gopurs_runtime.IntMul(int64(2), int64(v1_1)), int64(0)))) + (v_0)
					continue wrapOrder__gopurs_strict_thunk_0
					__t0 = func() int64 { panic("unreachable") }()
				}
			end_branch_0:
				return __t0
			}
		}()
	} else {

	}
wrapOrder__gopurs_strict_thunk_0:
	for {
		if false {
			continue wrapOrder__gopurs_strict_thunk_0
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var __t0 int64
		{
			if (v_0) == (int64(0)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
			v1_1_loop = gopurs_runtime.IntAdd(gopurs_runtime.IntMul(int64(2), v1_1), v_0)
			continue wrapOrder__gopurs_strict_thunk_0
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_wrapOrder(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
wrapOrder:
	for {
		if false {
			continue wrapOrder
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var __t0 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
			v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.IntMul(int64(2), gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit()).IntVal), v_0))
			})
			continue wrapOrder
			__t0 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_runOrder(depth_0_loop int64, seed_1_loop int64) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	return Call_Main_wrapOrder__gopurs_strict_thunk_0(depth_0, seed_1)
}

func Call_Main_checkAdds(depth_0_loop int64, seed_1_loop int64, step_2_loop int64) gopurs_runtime.Value {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	var step_2 int64 = step_2_loop
	_ = step_2
	return gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
		orig := struct {
			actual   int64
			expected int64
			/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
		}{Call_Main_wrapAdds__gopurs_strict_thunk_0(depth_0, step_2, seed_1), gopurs_runtime.IntAdd(seed_1, gopurs_runtime.IntMul(depth_0, step_2))}
		_ = orig
		return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
	}())
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
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
