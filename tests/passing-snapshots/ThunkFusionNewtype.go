package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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

var cache_Main_Lazy__609131035 gopurs_runtime.Value
var once_Main_Lazy__609131035 sync.Once

func Get_Main_Lazy__609131035() gopurs_runtime.Value {
	once_Main_Lazy__609131035.Do(func() {
		cache_Main_Lazy__609131035 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Lazy__609131035(x_0_box)
		})
	})
	return cache_Main_Lazy__609131035
}

var cache_Main_Lazy__3873826587 gopurs_runtime.Value
var once_Main_Lazy__3873826587 sync.Once

func Get_Main_Lazy__3873826587() gopurs_runtime.Value {
	once_Main_Lazy__3873826587.Do(func() {
		cache_Main_Lazy__3873826587 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Lazy__3873826587(x_0_box)
		})
	})
	return cache_Main_Lazy__3873826587
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

var cache_Main_force__421792098 gopurs_runtime.Value
var once_Main_force__421792098 sync.Once

func Get_Main_force__421792098() gopurs_runtime.Value {
	once_Main_force__421792098.Do(func() {
		cache_Main_force__421792098 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_force__421792098(v_0_box)
		})
	})
	return cache_Main_force__421792098
}

var cache_Main_force__3600449186 gopurs_runtime.Value
var once_Main_force__3600449186 sync.Once

func Get_Main_force__3600449186() gopurs_runtime.Value {
	once_Main_force__3600449186.Do(func() {
		cache_Main_force__3600449186 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_force__3600449186(v_0_box))
		})
	})
	return cache_Main_force__3600449186
}

var cache_Main_go__defer gopurs_runtime.Value
var once_Main_go__defer sync.Once

func Get_Main_go__defer() gopurs_runtime.Value {
	once_Main_go__defer.Do(func() {
		cache_Main_go__defer = Get_Main_Lazy()
	})
	return cache_Main_go__defer
}

var cache_Main_defer__609131035 gopurs_runtime.Value
var once_Main_defer__609131035 sync.Once

func Get_Main_defer__609131035() gopurs_runtime.Value {
	once_Main_defer__609131035.Do(func() {
		cache_Main_defer__609131035 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_defer__609131035(__eta_norm_0_0_box)
		})
	})
	return cache_Main_defer__609131035
}

var cache_Main_defer__3873826587 gopurs_runtime.Value
var once_Main_defer__3873826587 sync.Once

func Get_Main_defer__3873826587() gopurs_runtime.Value {
	once_Main_defer__3873826587.Do(func() {
		cache_Main_defer__3873826587 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_defer__3873826587(__eta_norm_0_0_box)
		})
	})
	return cache_Main_defer__3873826587
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
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(1000)))
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
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_wrapIncrements__gopurs_strict_thunk_0(int64(0), __local_var_5_5.IntVal), __local_var_5_5.IntVal}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
					actual   int64
					expected int64
				}{Call_Main_wrapIncrements__gopurs_strict_thunk_0(int64(1), __local_var_5_5.IntVal), (__local_var_5_5.IntVal) + (int64(1))}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
						actual   int64
						expected int64
					}{Call_Main_wrapIncrements__gopurs_strict_thunk_0(__local_var_4_4.IntVal, __local_var_5_5.IntVal), (__local_var_5_5.IntVal) + (__local_var_4_4.IntVal)}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkAdds(int64(0), __local_var_5_5.IntVal, __local_var_6_6.IntVal), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkAdds(int64(1), __local_var_5_5.IntVal, __local_var_6_6.IntVal), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkAdds(int64(7), __local_var_5_5.IntVal, __local_var_6_6.IntVal), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkAdds(__local_var_4_4.IntVal, __local_var_5_5.IntVal, __local_var_6_6.IntVal), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkAdds(int64(7), __local_var_6_6.IntVal, __local_var_5_5.IntVal), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
												actual   int64
												expected int64
											}{Call_Main_wrapOrder__gopurs_strict_thunk_0(int64(0), __local_var_5_5.IntVal), __local_var_5_5.IntVal}), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
													actual   int64
													expected int64
												}{Call_Main_wrapOrder__gopurs_strict_thunk_0(int64(1), __local_var_5_5.IntVal), int64(15)}), gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
														actual   int64
														expected int64
													}{Call_Main_wrapOrder__gopurs_strict_thunk_0(int64(3), __local_var_5_5.IntVal), int64(73)}), gopurs_runtime.Func(func(_dollar___unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
															__local_var_18_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), Call_Main_wrapIncrements(__local_var_4_4.IntVal, gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Int(__local_var_5_5.IntVal)
															}))), gopurs_runtime.Value{})
															_ = __local_var_18_7
															__local_var_19_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_18_7), gopurs_runtime.Value{})
															_ = __local_var_19_8
															return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
																actual   int64
																expected int64
															}{(gopurs_runtime.Apply(__local_var_19_8, Get_Data_Unit_unit()).IntVal) + (gopurs_runtime.Apply(__local_var_19_8, Get_Data_Unit_unit()).IntVal), (int64(2)) * ((__local_var_5_5.IntVal) + (__local_var_4_4.IntVal))}), gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	// TAST (Let): __local_var_21_9 shape=App(Var) bindingType=Any
																	__local_var_21_9 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(0)))
																	_ = __local_var_21_9
																	__local_var_22_10 := gopurs_runtime.Apply(__local_var_21_9, gopurs_runtime.Value{})
																	_ = __local_var_22_10
																	// TAST (Let): pending_23_11 shape=App(Var) bindingType=(Func [Unit] (ADT ["Effect","Effect"] [Int]))
																	pending_23_11 := Call_Main_wrapEffects(int64(3), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(count_24 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Int((count_24.IntVal) + (int64(1)))
																		}), __local_var_22_10), gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																				return __local_var_5_5
																			})
																		}))
																	}))
																	_ = pending_23_11
																	__local_var_24_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_22_10), gopurs_runtime.Value{})
																	_ = __local_var_24_12
																	return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
																		actual   int64
																		expected int64
																	}{__local_var_24_12.IntVal, int64(0)}), gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																			// TAST (Let): __local_var_26_13 shape=App(Other) bindingType=Any
																			__local_var_26_13 := gopurs_runtime.Apply(pending_23_11, Get_Data_Unit_unit())
																			_ = __local_var_26_13
																			__local_var_27_14 := gopurs_runtime.Apply(__local_var_26_13, gopurs_runtime.Value{})
																			_ = __local_var_27_14
																			__local_var_28_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(pending_23_11, Get_Data_Unit_unit()), gopurs_runtime.Value{})
																			_ = __local_var_28_15
																			__local_var_29_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_22_10), gopurs_runtime.Value{})
																			_ = __local_var_29_16
																			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
																				actual   int64
																				expected int64
																			}{__local_var_27_14.IntVal, int64(10)}), gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
																					actual   int64
																					expected int64
																				}{__local_var_28_15.IntVal, int64(10)}), gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
																						actual   int64
																						expected int64
																					}{__local_var_29_16.IntVal, int64(2)}), gopurs_runtime.Func(func(_dollar___unused_32 gopurs_runtime.Value) gopurs_runtime.Value {
																						return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
																					}))
																				}))
																			})), gopurs_runtime.Value{})
																		})
																	})), gopurs_runtime.Value{})
																})
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
					}))
				}))
			})), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_Lazy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Lazy__609131035(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
Lazy__609131035:
	for {
		if false {
			continue Lazy__609131035
		}
		var x_0 gopurs_runtime.Value = x_0_loop
		_ = x_0
		return x_0
	}
}

func Call_Main_Lazy__3873826587(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
Lazy__3873826587:
	for {
		if false {
			continue Lazy__3873826587
		}
		var x_0 gopurs_runtime.Value = x_0_loop
		_ = x_0
		return x_0
	}
}

func Call_Main_force(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(v_0, Get_Data_Unit_unit())
}

func Call_Main_force__421792098(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
force__421792098:
	for {
		if false {
			continue force__421792098
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		return gopurs_runtime.Apply(v_0, Get_Data_Unit_unit())
	}
}

func Call_Main_force__3600449186(v_0_loop gopurs_runtime.Value) int64 {
force__3600449186:
	for {
		if false {
			continue force__3600449186
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		return gopurs_runtime.Apply(v_0, Get_Data_Unit_unit()).IntVal
	}
}

func Call_Main_defer__609131035(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
defer__609131035:
	for {
		if false {
			continue defer__609131035
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return __eta_norm_0_0
	}
}

func Call_Main_defer__3873826587(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
defer__3873826587:
	for {
		if false {
			continue defer__3873826587
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return __eta_norm_0_0
	}
}

func Call_Main_wrapAdds__gopurs_strict_thunk_0(v_0_loop int64, v1_1_loop int64, v2_2_loop int64) int64 {
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
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = v1_1
			v2_2_loop = (v2_2) + (v1_1)
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
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = v1_1
			v2_2_loop = gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((gopurs_runtime.Apply(v2_2, Get_Data_Unit_unit()).IntVal) + (v1_1))
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
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(ADT ["Effect","Effect"] [Int])
					__local_var_3_0 := gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit())
					_ = __local_var_3_0
					__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
					_ = __local_var_4_1
					return gopurs_runtime.Int((__local_var_4_1.IntVal) + (int64(1)))
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
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = (v1_1) + (int64(1))
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
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit()).IntVal) + (int64(1)))
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
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = ((int64(2)) * (v1_1)) + (v_0)
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
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(((int64(2)) * (gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit()).IntVal)) + (v_0))
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
	return Call_Test_Assert_assertEqual_prime___627669702("", struct {
		actual   int64
		expected int64
	}{Call_Main_wrapAdds__gopurs_strict_thunk_0(depth_0, step_2, seed_1), (seed_1) + ((depth_0) * (step_2))})
}
