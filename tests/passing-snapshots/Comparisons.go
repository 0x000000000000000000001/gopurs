package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_checkOrdering gopurs_runtime.Value
var once_Main_checkOrdering sync.Once

func Get_Main_checkOrdering() gopurs_runtime.Value {
	once_Main_checkOrdering.Do(func() {
		cache_Main_checkOrdering = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Main_checkOrdering
}

var cache_Main_checkOrdering__39504560 gopurs_runtime.Value
var once_Main_checkOrdering__39504560 sync.Once

func Get_Main_checkOrdering__39504560() gopurs_runtime.Value {
	once_Main_checkOrdering__39504560.Do(func() {
		cache_Main_checkOrdering__39504560 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering__39504560(__eta_norm_2_0_box.StrVal(), __eta_norm_1_1_box.StrVal(), uint32(__eta_norm_0_unused_2_box.IntVal))
		})
	})
	return cache_Main_checkOrdering__39504560
}

var cache_Main_checkOrdering__4119744823 gopurs_runtime.Value
var once_Main_checkOrdering__4119744823 sync.Once

func Get_Main_checkOrdering__4119744823() gopurs_runtime.Value {
	once_Main_checkOrdering__4119744823.Do(func() {
		cache_Main_checkOrdering__4119744823 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering__4119744823(__eta_norm_2_0_box.StrVal(), __eta_norm_1_1_box.StrVal(), uint32(__eta_norm_0_unused_2_box.IntVal))
		})
	})
	return cache_Main_checkOrdering__4119744823
}

var cache_Main_checkOrdering__2650964060 gopurs_runtime.Value
var once_Main_checkOrdering__2650964060 sync.Once

func Get_Main_checkOrdering__2650964060() gopurs_runtime.Value {
	once_Main_checkOrdering__2650964060.Do(func() {
		cache_Main_checkOrdering__2650964060 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering__2650964060(__eta_norm_2_0_box.StrVal(), __eta_norm_1_1_box.StrVal(), uint32(__eta_norm_0_unused_2_box.IntVal))
		})
	})
	return cache_Main_checkOrdering__2650964060
}

var cache_Main_checkOrdering__325519965 gopurs_runtime.Value
var once_Main_checkOrdering__325519965 sync.Once

func Get_Main_checkOrdering__325519965() gopurs_runtime.Value {
	once_Main_checkOrdering__325519965.Do(func() {
		cache_Main_checkOrdering__325519965 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering__325519965(__eta_norm_2_0_box.StrVal(), __eta_norm_1_1_box.StrVal(), uint32(__eta_norm_0_unused_2_box.IntVal))
		})
	})
	return cache_Main_checkOrdering__325519965
}

var cache_Main_checkOrdering__2503009178 gopurs_runtime.Value
var once_Main_checkOrdering__2503009178 sync.Once

func Get_Main_checkOrdering__2503009178() gopurs_runtime.Value {
	once_Main_checkOrdering__2503009178.Do(func() {
		cache_Main_checkOrdering__2503009178 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering__2503009178(__eta_norm_2_0_box.StrVal(), __eta_norm_1_1_box.StrVal(), uint32(__eta_norm_0_unused_2_box.IntVal))
		})
	})
	return cache_Main_checkOrdering__2503009178
}

var cache_Main_checkOrdering__3515330801 gopurs_runtime.Value
var once_Main_checkOrdering__3515330801 sync.Once

func Get_Main_checkOrdering__3515330801() gopurs_runtime.Value {
	once_Main_checkOrdering__3515330801.Do(func() {
		cache_Main_checkOrdering__3515330801 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkOrdering__3515330801(__eta_norm_2_0_box.StrVal(), __eta_norm_1_1_box.StrVal(), uint32(__eta_norm_0_unused_2_box.IntVal))
		})
	})
	return cache_Main_checkOrdering__3515330801
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(1.0), gopurs_runtime.Float(2.0))
			_ = __t_tag_0
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_0.IntVal) == 1527465420))), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(3.0), gopurs_runtime.Float(1.0))
					_ = __t_tag_1
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_1.IntVal) == 380165415))), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str("a"), gopurs_runtime.Str("b"))
						_ = __t_tag_2
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_2.IntVal) == 1527465420))), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
								var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str("z"), gopurs_runtime.Str("a"))
								_ = __t_tag_3
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_3.IntVal) == 380165415))), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkOrdering__2650964060("a", "b", 1527465420), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkOrdering__39504560("é", "é", 902936544), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkOrdering__4119744823("Ā", "ÿ", 380165415), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkOrdering__2650964060("ÿ", "Ā", 1527465420), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkOrdering__3515330801("a", "b", 1527465420), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkOrdering__325519965("é", "é", 902936544), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkOrdering__2503009178("éa", "é", 380165415), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkOrdering__3515330801("é", "éa", 1527465420), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkOrdering__3515330801("😀", "😁", 1527465420), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkOrdering__2503009178("😁", "😀", 380165415), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
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
								}))
							}))
						}))
					}))
				}))
			}))
		}()
	})
	return cache_Main_main
}

func Call_Main_checkOrdering(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): Eq0_1_0 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeVar a$scope1)])
	Eq0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(dictOrd_0.V0, gopurs_runtime.Value{}))
	_ = Eq0_1_0
	return gopurs_runtime.Func3(func(left_2 gopurs_runtime.Value, right_3 gopurs_runtime.Value, expected_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=Any
			__local_var_5_1 := gopurs_runtime.Apply(Get_Effect_Ref__new(), left_2)
			_ = __local_var_5_1
			leftRef_6_2 := gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Value{})
			_ = leftRef_6_2
			rightRef_7_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), right_3), gopurs_runtime.Value{})
			_ = rightRef_7_3
			x_8_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), leftRef_6_2), gopurs_runtime.Value{})
			_ = x_8_4
			y_9_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), rightRef_7_3), gopurs_runtime.Value{})
			_ = y_9_5
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5).IntVal), uint32(expected_4.IntVal)}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t_tag_6 uint32 = uint32(gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5).IntVal)
				_ = __t_tag_6
				var __t_tag_7 uint32 = uint32(expected_4.IntVal)
				_ = __t_tag_7
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
					actual   bool
					expected bool
				}{(uint32(__t_tag_6) == 1527465420), (uint32(__t_tag_7) == 1527465420)}), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t_tag_8 uint32 = uint32(gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5).IntVal)
					_ = __t_tag_8
					var __t_tag_9 uint32 = uint32(expected_4.IntVal)
					_ = __t_tag_9
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
						actual   bool
						expected bool
					}{(uint32(__t_tag_8) == 380165415) != (true), (uint32(__t_tag_9) == 380165415) != (true)}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t_tag_10 uint32 = uint32(expected_4.IntVal)
						_ = __t_tag_10
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
							actual   bool
							expected bool
						}{(gopurs_runtime.Apply2(Eq0_1_0.V0, x_8_4, y_9_5).IntVal) != (0), (uint32(__t_tag_10) == 902936544)}), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
							var __t_tag_11 uint32 = uint32(expected_4.IntVal)
							_ = __t_tag_11
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
								actual   bool
								expected bool
							}{((gopurs_runtime.Apply2(Eq0_1_0.V0, x_8_4, y_9_5).IntVal) != (0)) != (true), (uint32(__t_tag_11) == 902936544) != (true)}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
								var __t_tag_12 uint32 = uint32(gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5).IntVal)
								_ = __t_tag_12
								var __t_tag_13 uint32 = uint32(expected_4.IntVal)
								_ = __t_tag_13
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
									actual   bool
									expected bool
								}{(uint32(__t_tag_12) == 1527465420) != (true), (uint32(__t_tag_13) == 1527465420) != (true)}), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
									var __t_tag_14 uint32 = uint32(gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, y_9_5).IntVal)
									_ = __t_tag_14
									var __t_tag_15 uint32 = uint32(expected_4.IntVal)
									_ = __t_tag_15
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
										actual   bool
										expected bool
									}{(uint32(__t_tag_14) == 380165415), (uint32(__t_tag_15) == 380165415)}), gopurs_runtime.Func(func(_dollar___unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											// TAST (Let): __local_var_17_16 shape=App(Var) bindingType=Any
											__local_var_17_16 := gopurs_runtime.Apply(Get_Effect_Ref__new(), dictOrd_0.V1)
											_ = __local_var_17_16
											comparatorRef_18_17 := gopurs_runtime.Apply(__local_var_17_16, gopurs_runtime.Value{})
											_ = comparatorRef_18_17
											comparator_19_18 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), comparatorRef_18_17), gopurs_runtime.Value{})
											_ = comparator_19_18
											return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
												actual   uint32
												expected uint32
											}{uint32(gopurs_runtime.Apply2(comparator_19_18, x_8_4, y_9_5).IntVal), uint32(expected_4.IntVal)}), gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													// TAST (Let): __local_var_21_19 shape=App(Var) bindingType=Any
													__local_var_21_19 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(comparator_19_18, x_8_4))
													_ = __local_var_21_19
													partialRef_22_20 := gopurs_runtime.Apply(__local_var_21_19, gopurs_runtime.Value{})
													_ = partialRef_22_20
													partial_23_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), partialRef_22_20), gopurs_runtime.Value{})
													_ = partial_23_21
													return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
														actual   uint32
														expected uint32
													}{uint32(gopurs_runtime.Apply(partial_23_21, y_9_5).IntVal), uint32(expected_4.IntVal)}), gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
															actual   uint32
															expected uint32
														}{uint32(gopurs_runtime.Apply(partial_23_21, x_8_4).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
																actual   uint32
																expected uint32
															}{uint32(gopurs_runtime.Apply(partial_23_21, y_9_5).IntVal), uint32(expected_4.IntVal)}), gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	// TAST (Let): __local_var_27_22 shape=App(Var) bindingType=Any
																	__local_var_27_22 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Func(func(a2_27 gopurs_runtime.Value) gopurs_runtime.Value {
																		var __t_tag_23 uint32 = uint32(gopurs_runtime.Apply2(dictOrd_0.V1, x_8_4, a2_27).IntVal)
																		_ = __t_tag_23
																		return gopurs_runtime.Bool((uint32(__t_tag_23) == 1527465420))
																	}))
																	_ = __local_var_27_22
																	lessThanRef_28_24 := gopurs_runtime.Apply(__local_var_27_22, gopurs_runtime.Value{})
																	_ = lessThanRef_28_24
																	lessThan_29_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), lessThanRef_28_24), gopurs_runtime.Value{})
																	_ = lessThan_29_25
																	var __t_tag_26 uint32 = uint32(expected_4.IntVal)
																	_ = __t_tag_26
																	return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																		actual   bool
																		expected bool
																	}{(gopurs_runtime.Apply(lessThan_29_25, y_9_5).IntVal) != (0), (uint32(__t_tag_26) == 1527465420)}), gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																			actual   bool
																			expected bool
																		}{(gopurs_runtime.Apply(lessThan_29_25, x_8_4).IntVal) != (0), false}), gopurs_runtime.Func(func(_dollar___unused_31 gopurs_runtime.Value) gopurs_runtime.Value {
																			var __t_tag_27 uint32 = uint32(expected_4.IntVal)
																			_ = __t_tag_27
																			return Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																				actual   bool
																				expected bool
																			}{(gopurs_runtime.Apply(lessThan_29_25, y_9_5).IntVal) != (0), (uint32(__t_tag_27) == 1527465420)})
																		}))
																	})), gopurs_runtime.Value{})
																})
															}))
														}))
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
			})), gopurs_runtime.Value{})
		})
	})
}

func Call_Main_checkOrdering__39504560(__eta_norm_2_0_loop string, __eta_norm_1_1_loop string, __eta_norm_0_unused_2_loop uint32) gopurs_runtime.Value {
checkOrdering__39504560:
	for {
		if false {
			continue checkOrdering__39504560
		}
		var __eta_norm_2_0 string = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_1 string = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_unused_2 uint32 = __eta_norm_0_unused_2_loop
		_ = __eta_norm_0_unused_2
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_2_0))
			_ = __local_var_3_0
			__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
			_ = __local_var_4_1
			__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_1_1)), gopurs_runtime.Value{})
			_ = __local_var_5_2
			__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
			_ = __local_var_6_3
			__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_2), gopurs_runtime.Value{})
			_ = __local_var_7_4
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
				_ = __t_tag_5
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
					actual   bool
					expected bool
				}{(uint32(__t_tag_5.IntVal) == 1527465420), false}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t_tag_6 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
					_ = __t_tag_6
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
						actual   bool
						expected bool
					}{(uint32(__t_tag_6.IntVal) == 380165415) != (true), true}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
							actual   bool
							expected bool
						}{(__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal()), true}), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
								actual   bool
								expected bool
							}{((__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal())) != (true), false}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
								var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
								_ = __t_tag_7
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
									actual   bool
									expected bool
								}{(uint32(__t_tag_7.IntVal) == 1527465420) != (true), true}), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
									var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
									_ = __t_tag_8
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
										actual   bool
										expected bool
									}{(uint32(__t_tag_8.IntVal) == 380165415), false}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											// TAST (Let): __local_var_15_9 shape=App(Var) bindingType=Any
											__local_var_15_9 := gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Data_Ord_compare__2320923292())
											_ = __local_var_15_9
											__local_var_16_10 := gopurs_runtime.Apply(__local_var_15_9, gopurs_runtime.Value{})
											_ = __local_var_16_10
											__local_var_17_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_16_10), gopurs_runtime.Value{})
											_ = __local_var_17_11
											return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
												actual   uint32
												expected uint32
											}{uint32(gopurs_runtime.Apply2(__local_var_17_11, gopurs_runtime.Str(__local_var_6_3.StrVal()), gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													__local_var_19_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_17_11, gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
													_ = __local_var_19_12
													__local_var_20_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_19_12), gopurs_runtime.Value{})
													_ = __local_var_20_13
													return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
														actual   uint32
														expected uint32
													}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
															actual   uint32
															expected uint32
														}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
																actual   uint32
																expected uint32
															}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	__local_var_24_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Data_Ord_lessThan__3343604442(), gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
																	_ = __local_var_24_14
																	__local_var_25_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_24_14), gopurs_runtime.Value{})
																	_ = __local_var_25_15
																	return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																		actual   bool
																		expected bool
																	}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false}), gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																			actual   bool
																			expected bool
																		}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal) != (0), false}), gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																			return Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																				actual   bool
																				expected bool
																			}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false})
																		}))
																	})), gopurs_runtime.Value{})
																})
															}))
														}))
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
			})), gopurs_runtime.Value{})
		})
	}
}

func Call_Main_checkOrdering__4119744823(__eta_norm_2_0_loop string, __eta_norm_1_1_loop string, __eta_norm_0_unused_2_loop uint32) gopurs_runtime.Value {
checkOrdering__4119744823:
	for {
		if false {
			continue checkOrdering__4119744823
		}
		var __eta_norm_2_0 string = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_1 string = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_unused_2 uint32 = __eta_norm_0_unused_2_loop
		_ = __eta_norm_0_unused_2
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_2_0))
			_ = __local_var_3_0
			__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
			_ = __local_var_4_1
			__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_1_1)), gopurs_runtime.Value{})
			_ = __local_var_5_2
			__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
			_ = __local_var_6_3
			__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_2), gopurs_runtime.Value{})
			_ = __local_var_7_4
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4).IntVal), 380165415}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
				_ = __t_tag_5
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
					actual   bool
					expected bool
				}{(uint32(__t_tag_5.IntVal) == 1527465420), false}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t_tag_6 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
					_ = __t_tag_6
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
						actual   bool
						expected bool
					}{(uint32(__t_tag_6.IntVal) == 380165415) != (true), false}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
							actual   bool
							expected bool
						}{(__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal()), false}), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
								actual   bool
								expected bool
							}{((__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal())) != (true), true}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
								var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
								_ = __t_tag_7
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
									actual   bool
									expected bool
								}{(uint32(__t_tag_7.IntVal) == 1527465420) != (true), true}), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
									var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
									_ = __t_tag_8
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
										actual   bool
										expected bool
									}{(uint32(__t_tag_8.IntVal) == 380165415), true}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											// TAST (Let): __local_var_15_9 shape=App(Var) bindingType=Any
											__local_var_15_9 := gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Data_Ord_compare__2320923292())
											_ = __local_var_15_9
											__local_var_16_10 := gopurs_runtime.Apply(__local_var_15_9, gopurs_runtime.Value{})
											_ = __local_var_16_10
											__local_var_17_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_16_10), gopurs_runtime.Value{})
											_ = __local_var_17_11
											return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
												actual   uint32
												expected uint32
											}{uint32(gopurs_runtime.Apply2(__local_var_17_11, gopurs_runtime.Str(__local_var_6_3.StrVal()), gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 380165415}), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													__local_var_19_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_17_11, gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
													_ = __local_var_19_12
													__local_var_20_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_19_12), gopurs_runtime.Value{})
													_ = __local_var_20_13
													return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
														actual   uint32
														expected uint32
													}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 380165415}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
															actual   uint32
															expected uint32
														}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
																actual   uint32
																expected uint32
															}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 380165415}), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	__local_var_24_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Data_Ord_lessThan__3343604442(), gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
																	_ = __local_var_24_14
																	__local_var_25_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_24_14), gopurs_runtime.Value{})
																	_ = __local_var_25_15
																	return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																		actual   bool
																		expected bool
																	}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false}), gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																			actual   bool
																			expected bool
																		}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal) != (0), false}), gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																			return Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																				actual   bool
																				expected bool
																			}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false})
																		}))
																	})), gopurs_runtime.Value{})
																})
															}))
														}))
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
			})), gopurs_runtime.Value{})
		})
	}
}

func Call_Main_checkOrdering__2650964060(__eta_norm_2_0_loop string, __eta_norm_1_1_loop string, __eta_norm_0_unused_2_loop uint32) gopurs_runtime.Value {
checkOrdering__2650964060:
	for {
		if false {
			continue checkOrdering__2650964060
		}
		var __eta_norm_2_0 string = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_1 string = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_unused_2 uint32 = __eta_norm_0_unused_2_loop
		_ = __eta_norm_0_unused_2
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_2_0))
			_ = __local_var_3_0
			__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
			_ = __local_var_4_1
			__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_1_1)), gopurs_runtime.Value{})
			_ = __local_var_5_2
			__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
			_ = __local_var_6_3
			__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_2), gopurs_runtime.Value{})
			_ = __local_var_7_4
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4).IntVal), 1527465420}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
				_ = __t_tag_5
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
					actual   bool
					expected bool
				}{(uint32(__t_tag_5.IntVal) == 1527465420), true}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t_tag_6 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
					_ = __t_tag_6
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
						actual   bool
						expected bool
					}{(uint32(__t_tag_6.IntVal) == 380165415) != (true), true}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
							actual   bool
							expected bool
						}{(__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal()), false}), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
								actual   bool
								expected bool
							}{((__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal())) != (true), true}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
								var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
								_ = __t_tag_7
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
									actual   bool
									expected bool
								}{(uint32(__t_tag_7.IntVal) == 1527465420) != (true), false}), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
									var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordCharImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
									_ = __t_tag_8
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
										actual   bool
										expected bool
									}{(uint32(__t_tag_8.IntVal) == 380165415), false}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											// TAST (Let): __local_var_15_9 shape=App(Var) bindingType=Any
											__local_var_15_9 := gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Data_Ord_compare__2320923292())
											_ = __local_var_15_9
											__local_var_16_10 := gopurs_runtime.Apply(__local_var_15_9, gopurs_runtime.Value{})
											_ = __local_var_16_10
											__local_var_17_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_16_10), gopurs_runtime.Value{})
											_ = __local_var_17_11
											return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
												actual   uint32
												expected uint32
											}{uint32(gopurs_runtime.Apply2(__local_var_17_11, gopurs_runtime.Str(__local_var_6_3.StrVal()), gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 1527465420}), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													__local_var_19_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_17_11, gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
													_ = __local_var_19_12
													__local_var_20_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_19_12), gopurs_runtime.Value{})
													_ = __local_var_20_13
													return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
														actual   uint32
														expected uint32
													}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 1527465420}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
															actual   uint32
															expected uint32
														}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
																actual   uint32
																expected uint32
															}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 1527465420}), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	__local_var_24_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Data_Ord_lessThan__3343604442(), gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
																	_ = __local_var_24_14
																	__local_var_25_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_24_14), gopurs_runtime.Value{})
																	_ = __local_var_25_15
																	return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																		actual   bool
																		expected bool
																	}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), true}), gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																			actual   bool
																			expected bool
																		}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal) != (0), false}), gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																			return Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																				actual   bool
																				expected bool
																			}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), true})
																		}))
																	})), gopurs_runtime.Value{})
																})
															}))
														}))
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
			})), gopurs_runtime.Value{})
		})
	}
}

func Call_Main_checkOrdering__325519965(__eta_norm_2_0_loop string, __eta_norm_1_1_loop string, __eta_norm_0_unused_2_loop uint32) gopurs_runtime.Value {
checkOrdering__325519965:
	for {
		if false {
			continue checkOrdering__325519965
		}
		var __eta_norm_2_0 string = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_1 string = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_unused_2 uint32 = __eta_norm_0_unused_2_loop
		_ = __eta_norm_0_unused_2
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_2_0))
			_ = __local_var_3_0
			__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
			_ = __local_var_4_1
			__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_1_1)), gopurs_runtime.Value{})
			_ = __local_var_5_2
			__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
			_ = __local_var_6_3
			__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_2), gopurs_runtime.Value{})
			_ = __local_var_7_4
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
				_ = __t_tag_5
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
					actual   bool
					expected bool
				}{(uint32(__t_tag_5.IntVal) == 1527465420), false}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t_tag_6 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
					_ = __t_tag_6
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
						actual   bool
						expected bool
					}{(uint32(__t_tag_6.IntVal) == 380165415) != (true), true}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
							actual   bool
							expected bool
						}{(__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal()), true}), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
								actual   bool
								expected bool
							}{((__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal())) != (true), false}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
								var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
								_ = __t_tag_7
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
									actual   bool
									expected bool
								}{(uint32(__t_tag_7.IntVal) == 1527465420) != (true), true}), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
									var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
									_ = __t_tag_8
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
										actual   bool
										expected bool
									}{(uint32(__t_tag_8.IntVal) == 380165415), false}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											// TAST (Let): __local_var_15_9 shape=App(Var) bindingType=Any
											__local_var_15_9 := gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Data_Ord_compare__2340924753())
											_ = __local_var_15_9
											__local_var_16_10 := gopurs_runtime.Apply(__local_var_15_9, gopurs_runtime.Value{})
											_ = __local_var_16_10
											__local_var_17_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_16_10), gopurs_runtime.Value{})
											_ = __local_var_17_11
											return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
												actual   uint32
												expected uint32
											}{uint32(gopurs_runtime.Apply2(__local_var_17_11, gopurs_runtime.Str(__local_var_6_3.StrVal()), gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													__local_var_19_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_17_11, gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
													_ = __local_var_19_12
													__local_var_20_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_19_12), gopurs_runtime.Value{})
													_ = __local_var_20_13
													return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
														actual   uint32
														expected uint32
													}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
															actual   uint32
															expected uint32
														}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
																actual   uint32
																expected uint32
															}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	__local_var_24_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Data_Ord_lessThan__2428873527(), gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
																	_ = __local_var_24_14
																	__local_var_25_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_24_14), gopurs_runtime.Value{})
																	_ = __local_var_25_15
																	return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																		actual   bool
																		expected bool
																	}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false}), gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																			actual   bool
																			expected bool
																		}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal) != (0), false}), gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																			return Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																				actual   bool
																				expected bool
																			}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false})
																		}))
																	})), gopurs_runtime.Value{})
																})
															}))
														}))
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
			})), gopurs_runtime.Value{})
		})
	}
}

func Call_Main_checkOrdering__2503009178(__eta_norm_2_0_loop string, __eta_norm_1_1_loop string, __eta_norm_0_unused_2_loop uint32) gopurs_runtime.Value {
checkOrdering__2503009178:
	for {
		if false {
			continue checkOrdering__2503009178
		}
		var __eta_norm_2_0 string = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_1 string = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_unused_2 uint32 = __eta_norm_0_unused_2_loop
		_ = __eta_norm_0_unused_2
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_2_0))
			_ = __local_var_3_0
			__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
			_ = __local_var_4_1
			__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_1_1)), gopurs_runtime.Value{})
			_ = __local_var_5_2
			__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
			_ = __local_var_6_3
			__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_2), gopurs_runtime.Value{})
			_ = __local_var_7_4
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4).IntVal), 380165415}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
				_ = __t_tag_5
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
					actual   bool
					expected bool
				}{(uint32(__t_tag_5.IntVal) == 1527465420), false}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t_tag_6 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
					_ = __t_tag_6
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
						actual   bool
						expected bool
					}{(uint32(__t_tag_6.IntVal) == 380165415) != (true), false}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
							actual   bool
							expected bool
						}{(__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal()), false}), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
								actual   bool
								expected bool
							}{((__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal())) != (true), true}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
								var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
								_ = __t_tag_7
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
									actual   bool
									expected bool
								}{(uint32(__t_tag_7.IntVal) == 1527465420) != (true), true}), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
									var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
									_ = __t_tag_8
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
										actual   bool
										expected bool
									}{(uint32(__t_tag_8.IntVal) == 380165415), true}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											// TAST (Let): __local_var_15_9 shape=App(Var) bindingType=Any
											__local_var_15_9 := gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Data_Ord_compare__2340924753())
											_ = __local_var_15_9
											__local_var_16_10 := gopurs_runtime.Apply(__local_var_15_9, gopurs_runtime.Value{})
											_ = __local_var_16_10
											__local_var_17_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_16_10), gopurs_runtime.Value{})
											_ = __local_var_17_11
											return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
												actual   uint32
												expected uint32
											}{uint32(gopurs_runtime.Apply2(__local_var_17_11, gopurs_runtime.Str(__local_var_6_3.StrVal()), gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 380165415}), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													__local_var_19_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_17_11, gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
													_ = __local_var_19_12
													__local_var_20_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_19_12), gopurs_runtime.Value{})
													_ = __local_var_20_13
													return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
														actual   uint32
														expected uint32
													}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 380165415}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
															actual   uint32
															expected uint32
														}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
																actual   uint32
																expected uint32
															}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 380165415}), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	__local_var_24_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Data_Ord_lessThan__2428873527(), gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
																	_ = __local_var_24_14
																	__local_var_25_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_24_14), gopurs_runtime.Value{})
																	_ = __local_var_25_15
																	return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																		actual   bool
																		expected bool
																	}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false}), gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																			actual   bool
																			expected bool
																		}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal) != (0), false}), gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																			return Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																				actual   bool
																				expected bool
																			}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), false})
																		}))
																	})), gopurs_runtime.Value{})
																})
															}))
														}))
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
			})), gopurs_runtime.Value{})
		})
	}
}

func Call_Main_checkOrdering__3515330801(__eta_norm_2_0_loop string, __eta_norm_1_1_loop string, __eta_norm_0_unused_2_loop uint32) gopurs_runtime.Value {
checkOrdering__3515330801:
	for {
		if false {
			continue checkOrdering__3515330801
		}
		var __eta_norm_2_0 string = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_1 string = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_unused_2 uint32 = __eta_norm_0_unused_2_loop
		_ = __eta_norm_0_unused_2
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
			__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_2_0))
			_ = __local_var_3_0
			__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
			_ = __local_var_4_1
			__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(__eta_norm_1_1)), gopurs_runtime.Value{})
			_ = __local_var_5_2
			__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
			_ = __local_var_6_3
			__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_2), gopurs_runtime.Value{})
			_ = __local_var_7_4
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
				actual   uint32
				expected uint32
			}{uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4).IntVal), 1527465420}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
				_ = __t_tag_5
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
					actual   bool
					expected bool
				}{(uint32(__t_tag_5.IntVal) == 1527465420), true}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t_tag_6 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
					_ = __t_tag_6
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
						actual   bool
						expected bool
					}{(uint32(__t_tag_6.IntVal) == 380165415) != (true), true}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
							actual   bool
							expected bool
						}{(__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal()), false}), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
								actual   bool
								expected bool
							}{((__local_var_6_3.StrVal()) == (__local_var_7_4.StrVal())) != (true), true}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
								var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
								_ = __t_tag_7
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
									actual   bool
									expected bool
								}{(uint32(__t_tag_7.IntVal) == 1527465420) != (true), false}), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
									var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __local_var_6_3, __local_var_7_4)
									_ = __t_tag_8
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
										actual   bool
										expected bool
									}{(uint32(__t_tag_8.IntVal) == 380165415), false}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											// TAST (Let): __local_var_15_9 shape=App(Var) bindingType=Any
											__local_var_15_9 := gopurs_runtime.Apply(Get_Effect_Ref__new(), Get_Data_Ord_compare__2340924753())
											_ = __local_var_15_9
											__local_var_16_10 := gopurs_runtime.Apply(__local_var_15_9, gopurs_runtime.Value{})
											_ = __local_var_16_10
											__local_var_17_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_16_10), gopurs_runtime.Value{})
											_ = __local_var_17_11
											return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
												actual   uint32
												expected uint32
											}{uint32(gopurs_runtime.Apply2(__local_var_17_11, gopurs_runtime.Str(__local_var_6_3.StrVal()), gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 1527465420}), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
													__local_var_19_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(__local_var_17_11, gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
													_ = __local_var_19_12
													__local_var_20_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_19_12), gopurs_runtime.Value{})
													_ = __local_var_20_13
													return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
														actual   uint32
														expected uint32
													}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 1527465420}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
															actual   uint32
															expected uint32
														}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal), 902936544}), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___162236433("", struct {
																actual   uint32
																expected uint32
															}{uint32(gopurs_runtime.Apply(__local_var_20_13, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal), 1527465420}), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																	__local_var_24_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Apply(Get_Data_Ord_lessThan__2428873527(), gopurs_runtime.Str(__local_var_6_3.StrVal()))), gopurs_runtime.Value{})
																	_ = __local_var_24_14
																	__local_var_25_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_24_14), gopurs_runtime.Value{})
																	_ = __local_var_25_15
																	return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																		actual   bool
																		expected bool
																	}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), true}), gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																			actual   bool
																			expected bool
																		}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_6_3.StrVal())).IntVal) != (0), false}), gopurs_runtime.Func(func(_dollar___unused_27 gopurs_runtime.Value) gopurs_runtime.Value {
																			return Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																				actual   bool
																				expected bool
																			}{(gopurs_runtime.Apply(__local_var_25_15, gopurs_runtime.Str(__local_var_7_4.StrVal())).IntVal) != (0), true})
																		}))
																	})), gopurs_runtime.Value{})
																})
															}))
														}))
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
			})), gopurs_runtime.Value{})
		})
	}
}
