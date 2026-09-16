package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Crimson gopurs_runtime.Value
var once_Main_Crimson sync.Once

func Get_Main_Crimson() gopurs_runtime.Value {
	once_Main_Crimson.Do(func() {
		cache_Main_Crimson = gopurs_runtime.Value{Type: 9, IntVal: int64(2247809753), UnsafePtr: nil}
	})
	return cache_Main_Crimson
}

var cache_Main_Onyx gopurs_runtime.Value
var once_Main_Onyx sync.Once

func Get_Main_Onyx() gopurs_runtime.Value {
	once_Main_Onyx.Do(func() {
		cache_Main_Onyx = gopurs_runtime.Value{Type: 9, IntVal: int64(1685833310), UnsafePtr: nil}
	})
	return cache_Main_Onyx
}

var cache_Main_Tip gopurs_runtime.Value
var once_Main_Tip sync.Once

func Get_Main_Tip() gopurs_runtime.Value {
	once_Main_Tip.Do(func() {
		cache_Main_Tip = gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer((*Constructor_Main_Branch)(nil))}
	})
	return cache_Main_Tip
}

var cache_Main_Branch gopurs_runtime.Value
var once_Main_Branch sync.Once

func Get_Main_Branch() gopurs_runtime.Value {
	once_Main_Branch.Do(func() {
		cache_Main_Branch = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer((&Constructor_Main_Branch{1, uint32(value0.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](value1), value2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](value3)}))}
					})
				})
			})
		})
	})
	return cache_Main_Branch
}

var cache_Main_render gopurs_runtime.Value
var once_Main_render sync.Once

func Get_Main_render() gopurs_runtime.Value {
	once_Main_render.Do(func() {
		cache_Main_render = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v_0_box)))
		})
	})
	return cache_Main_render
}

var cache_Main_rebalance gopurs_runtime.Value
var once_Main_rebalance sync.Once

func Get_Main_rebalance() gopurs_runtime.Value {
	once_Main_rebalance.Do(func() {
		cache_Main_rebalance = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_rebalance(uint32(v_0_box.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v1_1_box), v2_2_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v3_3_box)))}
		})
	})
	return cache_Main_rebalance
}

var cache_Main_leftChild gopurs_runtime.Value
var once_Main_leftChild sync.Once

func Get_Main_leftChild() gopurs_runtime.Value {
	once_Main_leftChild.Do(func() {
		cache_Main_leftChild = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_leftChild(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v_0_box)))}
		})
	})
	return cache_Main_leftChild
}

var cache_Main_isRed gopurs_runtime.Value
var once_Main_isRed sync.Once

func Get_Main_isRed() gopurs_runtime.Value {
	once_Main_isRed.Do(func() {
		cache_Main_isRed = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_isRed(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v_0_box)))
		})
	})
	return cache_Main_isRed
}

var cache_Main_descend gopurs_runtime.Value
var once_Main_descend sync.Once

func Get_Main_descend() gopurs_runtime.Value {
	once_Main_descend.Do(func() {
		cache_Main_descend = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_descend(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v1_1_box)))}
		})
	})
	return cache_Main_descend
}

var cache_Main_blacken gopurs_runtime.Value
var once_Main_blacken sync.Once

func Get_Main_blacken() gopurs_runtime.Value {
	once_Main_blacken.Do(func() {
		cache_Main_blacken = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_blacken(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v_0_box)))}
		})
	})
	return cache_Main_blacken
}

var cache_Main_put gopurs_runtime.Value
var once_Main_put sync.Once

func Get_Main_put() gopurs_runtime.Value {
	once_Main_put.Do(func() {
		cache_Main_put = gopurs_runtime.Func2(func(value_0_box gopurs_runtime.Value, tree_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_put(value_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](tree_1_box)))}
		})
	})
	return cache_Main_put
}

var cache_Main_build gopurs_runtime.Value
var once_Main_build sync.Once

func Get_Main_build() gopurs_runtime.Value {
	once_Main_build.Do(func() {
		cache_Main_build = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_build(v_0_box.IntVal, v1_1_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v2_2_box)))}
		})
	})
	return cache_Main_build
}

var cache_Main_graftAndPut gopurs_runtime.Value
var once_Main_graftAndPut sync.Once

func Get_Main_graftAndPut() gopurs_runtime.Value {
	once_Main_graftAndPut.Do(func() {
		cache_Main_graftAndPut = gopurs_runtime.Func(func(child_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_graftAndPut(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](child_0_box)))}
		})
	})
	return cache_Main_graftAndPut
}

var cache_Main_mixed gopurs_runtime.Value
var once_Main_mixed sync.Once

func Get_Main_mixed() gopurs_runtime.Value {
	once_Main_mixed.Do(func() {
		cache_Main_mixed = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_mixed(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v1_1_box)))}
		})
	})
	return cache_Main_mixed
}

var cache_Main_audit gopurs_runtime.Value
var once_Main_audit sync.Once

func Get_Main_audit() gopurs_runtime.Value {
	once_Main_audit.Do(func() {
		cache_Main_audit = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_audit(v_0_box.IntVal, v1_1_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v2_2_box))
				_ = orig
				return gopurs_runtime.RecordDict3("blackHeight", "size", "valid", gopurs_runtime.Int(orig.blackHeight), gopurs_runtime.Int(orig.size), gopurs_runtime.Bool(orig.valid))
			}()
		})
	})
	return cache_Main_audit
}

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func2(func(expectedSize_0_box gopurs_runtime.Value, tree_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check(expectedSize_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](tree_1_box))
		})
	})
	return cache_Main_check
}

var cache_Main_snapshots gopurs_runtime.Value
var once_Main_snapshots sync.Once

func Get_Main_snapshots() gopurs_runtime.Value {
	once_Main_snapshots.Do(func() {
		cache_Main_snapshots = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_snapshots(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](v1_1_box))
		})
	})
	return cache_Main_snapshots
}

var cache_Main_three gopurs_runtime.Value
var once_Main_three sync.Once

func Get_Main_three() gopurs_runtime.Value {
	once_Main_three.Do(func() {
		cache_Main_three = gopurs_runtime.Func3(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_three(a_0_box.IntVal, b_1_box.IntVal, c_2_box.IntVal)
		})
	})
	return cache_Main_three
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_three(int64(3), int64(2), int64(1)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_three(int64(3), int64(1), int64(2)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_three(int64(1), int64(3), int64(2)), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_three(int64(1), int64(2), int64(3)), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=Any
							__local_var_4_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(127)))
							_ = __local_var_4_0
							__local_var_5_1 := gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Value{})
							_ = __local_var_5_1
							__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_1), gopurs_runtime.Value{})
							_ = __local_var_6_2
							__local_var_7_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main___gopurs_owned_build_0(__local_var_6_2.IntVal, int64(-1), (*Constructor_Main_Branch)(nil)))}), gopurs_runtime.Value{})
							_ = __local_var_7_3
							__local_var_8_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main___gopurs_owned_build_0(__local_var_6_2.IntVal, int64(1), (*Constructor_Main_Branch)(nil)))}), gopurs_runtime.Value{})
							_ = __local_var_8_4
							__local_var_9_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_mixed(__local_var_6_2.IntVal, (*Constructor_Main_Branch)(nil)))}), gopurs_runtime.Value{})
							_ = __local_var_9_5
							__local_var_10_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_7_3), gopurs_runtime.Value{})
							_ = __local_var_10_6
							__local_var_11_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_8_4), gopurs_runtime.Value{})
							_ = __local_var_11_7
							__local_var_12_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_9_5), gopurs_runtime.Value{})
							_ = __local_var_12_8
							return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(__local_var_6_2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_10_6)), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(__local_var_6_2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_11_7)), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(__local_var_6_2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_12_8)), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											__local_var_16_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_12_8)))), gopurs_runtime.Value{})
											_ = __local_var_16_9
											__local_var_17_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_16_9), gopurs_runtime.Value{})
											_ = __local_var_17_10
											__local_var_18_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_mixed(__local_var_6_2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_12_8)))}), gopurs_runtime.Value{})
											_ = __local_var_18_11
											__local_var_19_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_18_11), gopurs_runtime.Value{})
											_ = __local_var_19_12
											return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(__local_var_6_2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_19_12)), gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
													actual   string
													expected string
												}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_19_12)), __local_var_17_10.StrVal()}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
														actual   string
														expected string
													}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_12_8)), __local_var_17_10.StrVal()}), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_snapshots(int64(32), (*Constructor_Main_Branch)(nil)), gopurs_runtime.Func(func(_dollar___unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																// TAST (Let): __local_var_24_13 shape=App(Var) bindingType=Any
																__local_var_24_13 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer((&Constructor_Main_Branch{1, 1685833310, (&Constructor_Main_Branch{1, 1685833310, (*Constructor_Main_Branch)(nil), int64(10), (*Constructor_Main_Branch)(nil)}), int64(50), (&Constructor_Main_Branch{1, 1685833310, (*Constructor_Main_Branch)(nil), int64(90), (*Constructor_Main_Branch)(nil)})}))})
																_ = __local_var_24_13
																__local_var_25_14 := gopurs_runtime.Apply(__local_var_24_13, gopurs_runtime.Value{})
																_ = __local_var_25_14
																__local_var_26_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_25_14), gopurs_runtime.Value{})
																_ = __local_var_26_15
																var __t19 *Constructor_Main_Branch
																{
																	var __t_tag_17 *Constructor_Main_Branch = gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_26_15)
																	_ = __t_tag_17
																	if __t_tag_17 == nil {
																		__t19 = (*Constructor_Main_Branch)(nil)
																		goto end_branch_19
																	} else {

																	}
																}
																{
																	var __t_tag_18 *Constructor_Main_Branch = gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_26_15)
																	_ = __t_tag_18
																	if __t_tag_18 != nil {
																		__t19 = (*Constructor_Main_Branch)(__local_var_26_15.UnsafePtr).V1
																		goto end_branch_19
																	} else {

																	}
																}
																{
																	__t19 = func() *Constructor_Main_Branch { panic("Failed pattern match") }()
																}
															end_branch_19:
																__local_var_27_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(__t19)}), gopurs_runtime.Value{})
																_ = __local_var_27_16
																__local_var_28_20 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_27_16), gopurs_runtime.Value{})
																_ = __local_var_28_20
																__local_var_29_21 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_26_15)))), gopurs_runtime.Value{})
																_ = __local_var_29_21
																__local_var_30_22 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_28_20)))), gopurs_runtime.Value{})
																_ = __local_var_30_22
																__local_var_31_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_29_21), gopurs_runtime.Value{})
																_ = __local_var_31_23
																__local_var_32_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_30_22), gopurs_runtime.Value{})
																_ = __local_var_32_24
																__local_var_33_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_put(int64(5), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_26_15)))}), gopurs_runtime.Value{})
																_ = __local_var_33_25
																__local_var_34_26 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_33_25), gopurs_runtime.Value{})
																_ = __local_var_34_26
																return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(int64(4), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_34_26)), gopurs_runtime.Func(func(_dollar___unused_35 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
																		actual   string
																		expected string
																	}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_26_15)), __local_var_31_23.StrVal()}), gopurs_runtime.Func(func(_dollar___unused_36 gopurs_runtime.Value) gopurs_runtime.Value {
																		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
																			actual   string
																			expected string
																		}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_28_20)), __local_var_32_24.StrVal()}), gopurs_runtime.Func(func(_dollar___unused_37 gopurs_runtime.Value) gopurs_runtime.Value {
																			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																				__local_var_38_27 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_graftAndPut(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_28_20)))}), gopurs_runtime.Value{})
																				_ = __local_var_38_27
																				__local_var_39_28 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_38_27), gopurs_runtime.Value{})
																				_ = __local_var_39_28
																				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(int64(4), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_39_28)), gopurs_runtime.Func(func(_dollar___unused_40 gopurs_runtime.Value) gopurs_runtime.Value {
																					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
																						actual   string
																						expected string
																					}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_28_20)), __local_var_32_24.StrVal()}), gopurs_runtime.Func(func(_dollar___unused_41 gopurs_runtime.Value) gopurs_runtime.Value {
																						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
																							actual   string
																							expected string
																						}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_26_15)), __local_var_31_23.StrVal()}), gopurs_runtime.Func(func(_dollar___unused_42 gopurs_runtime.Value) gopurs_runtime.Value {
																							return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
																						}))
																					}))
																				})), gopurs_runtime.Value{})
																			})
																		}))
																	}))
																})), gopurs_runtime.Value{})
															})
														}))
													}))
												}))
											})), gopurs_runtime.Value{})
										})
									}))
								}))
							})), gopurs_runtime.Value{})
						})
					}))
				}))
			}))
		}))
	})
	return cache_Main_main
}

func Call_Main___gopurs_owned_blacken_0_consume(__arg0 *Constructor_Main_Branch, __donor *Constructor_Main_Branch) *Constructor_Main_Branch {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		if (__arg0) == (nil) {
			__donor_slot_1 := __donor
			_ = __donor_slot_1
			__dead_0 := __arg0
			_ = __dead_0
			return nil
		} else {
			if (__arg0) != (nil) {
				__scalar_2 := uint32(1685833310)
				_ = __scalar_2
				__read_3 := __arg0.V1
				_ = __read_3
				__scalar_4 := int64(__arg0.V2)
				_ = __scalar_4
				__read_5 := __arg0.V3
				_ = __read_5
				__donor_slot_7 := __donor
				_ = __donor_slot_7
				__dead_6 := __arg0
				_ = __dead_6
				var __cell_8 *Constructor_Main_Branch
				if (__donor_slot_7) != (nil) {
					__cell_8 = __donor_slot_7
					__donor_slot_7 = nil
				} else {
					if (__dead_6) != (nil) {
						__cell_8 = __dead_6
						__dead_6 = nil
					} else {

					}
				}
				if (__cell_8) == (nil) {
					__cell_8 = new(Constructor_Main_Branch)
				} else {

				}
				__cell_8.Rc = 1
				__cell_8.V0 = __scalar_2
				__cell_8.V1 = __read_3
				__cell_8.V2 = __scalar_4
				__cell_8.V3 = __read_5
				return __cell_8
			} else {
				panic("Failed pattern match")
			}
		}
	}
}

func Call_Main___gopurs_owned_blacken_0(__arg0 *Constructor_Main_Branch) *Constructor_Main_Branch {
	return Call_Main___gopurs_owned_blacken_0_consume(__arg0, nil)
}

func Call_Main___gopurs_owned_build_0_consume(__arg0 int64, __arg1 int64, __arg2 *Constructor_Main_Branch, __donor *Constructor_Main_Branch) *Constructor_Main_Branch {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		if (__arg0) == (int64(0)) {
			__read_7 := __arg2
			_ = __read_7
			__donor_slot_8 := __donor
			_ = __donor_slot_8
			return __read_7
		} else {
			__scalar_0 := int64((__arg0) - (int64(1)))
			_ = __scalar_0
			__scalar_1 := int64(__arg1)
			_ = __scalar_1
			__scalar_2 := int64((__arg0) * (__arg1))
			_ = __scalar_2
			__read_3 := __arg2
			_ = __read_3
			__donor_slot_4 := __donor
			_ = __donor_slot_4
			__result_5 := Call_Main___gopurs_owned_put_0_consume(__scalar_2, __read_3, nil)
			_ = __result_5
			var __cell_6 *Constructor_Main_Branch
			if (__donor_slot_4) != (nil) {
				__cell_6 = __donor_slot_4
				__donor_slot_4 = nil
			} else {

			}
			__arg0 = __scalar_0
			__arg1 = __scalar_1
			__arg2 = __result_5
			__donor = __cell_6
			continue __owned_loop
		}
	}
}

func Call_Main___gopurs_owned_build_0(__arg0 int64, __arg1 int64, __arg2 *Constructor_Main_Branch) *Constructor_Main_Branch {
	return Call_Main___gopurs_owned_build_0_consume(__arg0, __arg1, __arg2, nil)
}

func Call_Main___gopurs_owned_descend_0_consume(__arg0 int64, __arg1 *Constructor_Main_Branch, __donor *Constructor_Main_Branch) *Constructor_Main_Branch {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		if (__arg1) == (nil) {
			__scalar_0 := uint32(2247809753)
			_ = __scalar_0
			__scalar_1 := int64(__arg0)
			_ = __scalar_1
			__donor_slot_3 := __donor
			_ = __donor_slot_3
			__dead_2 := __arg1
			_ = __dead_2
			var __cell_4 *Constructor_Main_Branch
			if (__donor_slot_3) != (nil) {
				__cell_4 = __donor_slot_3
				__donor_slot_3 = nil
			} else {
				if (__dead_2) != (nil) {
					__cell_4 = __dead_2
					__dead_2 = nil
				} else {

				}
			}
			if (__cell_4) == (nil) {
				__cell_4 = new(Constructor_Main_Branch)
			} else {

			}
			__cell_4.Rc = 1
			__cell_4.V0 = __scalar_0
			__cell_4.V1 = nil
			__cell_4.V2 = __scalar_1
			__cell_4.V3 = nil
			return __cell_4
		} else {
			if (__arg1) != (nil) {
				if (__arg0) < (__arg1.V2) {
					__scalar_12 := uint32(__arg1.V0)
					_ = __scalar_12
					__scalar_13 := int64(__arg0)
					_ = __scalar_13
					__read_14 := __arg1.V1
					_ = __read_14
					__scalar_15 := int64(__arg1.V2)
					_ = __scalar_15
					__read_16 := __arg1.V3
					_ = __read_16
					__donor_slot_18 := __donor
					_ = __donor_slot_18
					__dead_17 := __arg1
					_ = __dead_17
					__result_19 := Call_Main___gopurs_owned_descend_0_consume(__scalar_13, __read_14, nil)
					_ = __result_19
					var __cell_20 *Constructor_Main_Branch
					if (__donor_slot_18) != (nil) {
						__cell_20 = __donor_slot_18
						__donor_slot_18 = nil
					} else {
						if (__dead_17) != (nil) {
							__cell_20 = __dead_17
							__dead_17 = nil
						} else {

						}
					}
					__result_21 := Call_Main___gopurs_owned_rebalance_0_consume(__scalar_12, __result_19, __scalar_15, __read_16, __cell_20)
					_ = __result_21
					return __result_21
				} else {
					if (__arg0) > (__arg1.V2) {
						__scalar_22 := uint32(__arg1.V0)
						_ = __scalar_22
						__read_23 := __arg1.V1
						_ = __read_23
						__scalar_24 := int64(__arg1.V2)
						_ = __scalar_24
						__scalar_25 := int64(__arg0)
						_ = __scalar_25
						__read_26 := __arg1.V3
						_ = __read_26
						__donor_slot_28 := __donor
						_ = __donor_slot_28
						__dead_27 := __arg1
						_ = __dead_27
						__result_29 := Call_Main___gopurs_owned_descend_0_consume(__scalar_25, __read_26, nil)
						_ = __result_29
						var __cell_30 *Constructor_Main_Branch
						if (__donor_slot_28) != (nil) {
							__cell_30 = __donor_slot_28
							__donor_slot_28 = nil
						} else {
							if (__dead_27) != (nil) {
								__cell_30 = __dead_27
								__dead_27 = nil
							} else {

							}
						}
						__result_31 := Call_Main___gopurs_owned_rebalance_0_consume(__scalar_22, __read_23, __scalar_24, __result_29, __cell_30)
						_ = __result_31
						return __result_31
					} else {
						__scalar_5 := uint32(__arg1.V0)
						_ = __scalar_5
						__read_6 := __arg1.V1
						_ = __read_6
						__scalar_7 := int64(__arg1.V2)
						_ = __scalar_7
						__read_8 := __arg1.V3
						_ = __read_8
						__donor_slot_10 := __donor
						_ = __donor_slot_10
						__dead_9 := __arg1
						_ = __dead_9
						var __cell_11 *Constructor_Main_Branch
						if (__donor_slot_10) != (nil) {
							__cell_11 = __donor_slot_10
							__donor_slot_10 = nil
						} else {
							if (__dead_9) != (nil) {
								__cell_11 = __dead_9
								__dead_9 = nil
							} else {

							}
						}
						if (__cell_11) == (nil) {
							__cell_11 = new(Constructor_Main_Branch)
						} else {

						}
						__cell_11.Rc = 1
						__cell_11.V0 = __scalar_5
						__cell_11.V1 = __read_6
						__cell_11.V2 = __scalar_7
						__cell_11.V3 = __read_8
						return __cell_11
					}
				}
			} else {
				panic("Failed pattern match")
			}
		}
	}
}

func Call_Main___gopurs_owned_descend_0(__arg0 int64, __arg1 *Constructor_Main_Branch) *Constructor_Main_Branch {
	return Call_Main___gopurs_owned_descend_0_consume(__arg0, __arg1, nil)
}

func Call_Main___gopurs_owned_graftAndPut_0_consume(__arg0 *Constructor_Main_Branch, __donor *Constructor_Main_Branch) *Constructor_Main_Branch {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		__scalar_0 := int64(int64(5))
		_ = __scalar_0
		__scalar_1 := uint32(1685833310)
		_ = __scalar_1
		__read_2 := __arg0
		_ = __read_2
		__scalar_3 := int64(int64(50))
		_ = __scalar_3
		__scalar_4 := uint32(1685833310)
		_ = __scalar_4
		__scalar_5 := int64(int64(90))
		_ = __scalar_5
		__donor_slot_6 := __donor
		_ = __donor_slot_6
		var __cell_7 *Constructor_Main_Branch
		if (__donor_slot_6) != (nil) {
			__cell_7 = __donor_slot_6
			__donor_slot_6 = nil
		} else {

		}
		if (__cell_7) == (nil) {
			__cell_7 = new(Constructor_Main_Branch)
		} else {

		}
		__cell_7.Rc = 1
		__cell_7.V0 = __scalar_4
		__cell_7.V1 = nil
		__cell_7.V2 = __scalar_5
		__cell_7.V3 = nil
		var __cell_8 *Constructor_Main_Branch
		if (__donor_slot_6) != (nil) {
			__cell_8 = __donor_slot_6
			__donor_slot_6 = nil
		} else {

		}
		if (__cell_8) == (nil) {
			__cell_8 = new(Constructor_Main_Branch)
		} else {

		}
		__cell_8.Rc = 1
		__cell_8.V0 = __scalar_1
		__cell_8.V1 = __read_2
		__cell_8.V2 = __scalar_3
		__cell_8.V3 = __cell_7
		var __cell_9 *Constructor_Main_Branch
		if (__donor_slot_6) != (nil) {
			__cell_9 = __donor_slot_6
			__donor_slot_6 = nil
		} else {

		}
		__result_10 := Call_Main___gopurs_owned_put_0_consume(__scalar_0, __cell_8, __cell_9)
		_ = __result_10
		return __result_10
	}
}

func Call_Main___gopurs_owned_graftAndPut_0(__arg0 *Constructor_Main_Branch) *Constructor_Main_Branch {
	return Call_Main___gopurs_owned_graftAndPut_0_consume(__arg0, nil)
}

func Call_Main___gopurs_owned_leftChild_0_consume(__arg0 *Constructor_Main_Branch, __donor *Constructor_Main_Branch) *Constructor_Main_Branch {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		if (__arg0) == (nil) {
			__donor_slot_1 := __donor
			_ = __donor_slot_1
			__dead_0 := __arg0
			_ = __dead_0
			return nil
		} else {
			if (__arg0) != (nil) {
				__read_2 := __arg0.V1
				_ = __read_2
				__donor_slot_4 := __donor
				_ = __donor_slot_4
				__dead_3 := __arg0
				_ = __dead_3
				return __read_2
			} else {
				panic("Failed pattern match")
			}
		}
	}
}

func Call_Main___gopurs_owned_leftChild_0(__arg0 *Constructor_Main_Branch) *Constructor_Main_Branch {
	return Call_Main___gopurs_owned_leftChild_0_consume(__arg0, nil)
}

func Call_Main___gopurs_owned_put_0_consume(__arg0 int64, __arg1 *Constructor_Main_Branch, __donor *Constructor_Main_Branch) *Constructor_Main_Branch {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		__scalar_0 := int64(__arg0)
		_ = __scalar_0
		__read_1 := __arg1
		_ = __read_1
		__donor_slot_2 := __donor
		_ = __donor_slot_2
		__result_3 := Call_Main___gopurs_owned_descend_0_consume(__scalar_0, __read_1, nil)
		_ = __result_3
		var __cell_4 *Constructor_Main_Branch
		if (__donor_slot_2) != (nil) {
			__cell_4 = __donor_slot_2
			__donor_slot_2 = nil
		} else {

		}
		__result_5 := Call_Main___gopurs_owned_blacken_0_consume(__result_3, __cell_4)
		_ = __result_5
		return __result_5
	}
}

func Call_Main___gopurs_owned_put_0(__arg0 int64, __arg1 *Constructor_Main_Branch) *Constructor_Main_Branch {
	return Call_Main___gopurs_owned_put_0_consume(__arg0, __arg1, nil)
}

func Call_Main___gopurs_owned_rebalance_0_consume(__arg0 uint32, __arg1 *Constructor_Main_Branch, __arg2 int64, __arg3 *Constructor_Main_Branch, __donor *Constructor_Main_Branch) *Constructor_Main_Branch {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		if (__arg0) == (1685833310) {
			if (__arg1) != (nil) {
				if (__arg1.V0) == (2247809753) {
					if (__arg1.V1) != (nil) {
						if (__arg1.V1.V0) == (2247809753) {
							__let_scalar_40 := int64(__arg1.V1.V2)
							_ = __let_scalar_40
							__let_scalar_41 := int64(__arg1.V2)
							_ = __let_scalar_41
							__let_scalar_42 := int64(__arg2)
							_ = __let_scalar_42
							__scalar_43 := uint32(2247809753)
							_ = __scalar_43
							__scalar_44 := uint32(1685833310)
							_ = __scalar_44
							__read_45 := __arg1.V1.V1
							_ = __read_45
							__scalar_46 := int64(__let_scalar_40)
							_ = __scalar_46
							__read_47 := __arg1.V1.V3
							_ = __read_47
							__scalar_48 := int64(__let_scalar_41)
							_ = __scalar_48
							__scalar_49 := uint32(1685833310)
							_ = __scalar_49
							__read_50 := __arg1.V3
							_ = __read_50
							__scalar_51 := int64(__let_scalar_42)
							_ = __scalar_51
							__read_52 := __arg3
							_ = __read_52
							__donor_slot_55 := __donor
							_ = __donor_slot_55
							__dead_53 := __arg1
							_ = __dead_53
							__dead_54 := __arg1.V1
							_ = __dead_54
							var __cell_56 *Constructor_Main_Branch
							if (__donor_slot_55) != (nil) {
								__cell_56 = __donor_slot_55
								__donor_slot_55 = nil
							} else {
								if (__dead_53) != (nil) {
									__cell_56 = __dead_53
									__dead_53 = nil
								} else {
									if (__dead_54) != (nil) {
										__cell_56 = __dead_54
										__dead_54 = nil
									} else {

									}
								}
							}
							if (__cell_56) == (nil) {
								__cell_56 = new(Constructor_Main_Branch)
							} else {

							}
							__cell_56.Rc = 1
							__cell_56.V0 = __scalar_44
							__cell_56.V1 = __read_45
							__cell_56.V2 = __scalar_46
							__cell_56.V3 = __read_47
							var __cell_57 *Constructor_Main_Branch
							if (__donor_slot_55) != (nil) {
								__cell_57 = __donor_slot_55
								__donor_slot_55 = nil
							} else {
								if (__dead_53) != (nil) {
									__cell_57 = __dead_53
									__dead_53 = nil
								} else {
									if (__dead_54) != (nil) {
										__cell_57 = __dead_54
										__dead_54 = nil
									} else {

									}
								}
							}
							if (__cell_57) == (nil) {
								__cell_57 = new(Constructor_Main_Branch)
							} else {

							}
							__cell_57.Rc = 1
							__cell_57.V0 = __scalar_49
							__cell_57.V1 = __read_50
							__cell_57.V2 = __scalar_51
							__cell_57.V3 = __read_52
							var __cell_58 *Constructor_Main_Branch
							if (__donor_slot_55) != (nil) {
								__cell_58 = __donor_slot_55
								__donor_slot_55 = nil
							} else {
								if (__dead_53) != (nil) {
									__cell_58 = __dead_53
									__dead_53 = nil
								} else {
									if (__dead_54) != (nil) {
										__cell_58 = __dead_54
										__dead_54 = nil
									} else {

									}
								}
							}
							if (__cell_58) == (nil) {
								__cell_58 = new(Constructor_Main_Branch)
							} else {

							}
							__cell_58.Rc = 1
							__cell_58.V0 = __scalar_43
							__cell_58.V1 = __cell_56
							__cell_58.V2 = __scalar_48
							__cell_58.V3 = __cell_57
							return __cell_58
						} else {
							if (__arg1.V3) != (nil) {
								if (__arg1.V3.V0) == (2247809753) {
									__let_scalar_67 := int64(__arg1.V2)
									_ = __let_scalar_67
									__let_scalar_68 := int64(__arg1.V3.V2)
									_ = __let_scalar_68
									__let_scalar_69 := int64(__arg2)
									_ = __let_scalar_69
									__scalar_70 := uint32(2247809753)
									_ = __scalar_70
									__scalar_71 := uint32(1685833310)
									_ = __scalar_71
									__read_72 := __arg1.V1
									_ = __read_72
									__scalar_73 := int64(__let_scalar_67)
									_ = __scalar_73
									__read_74 := __arg1.V3.V1
									_ = __read_74
									__scalar_75 := int64(__let_scalar_68)
									_ = __scalar_75
									__scalar_76 := uint32(1685833310)
									_ = __scalar_76
									__read_77 := __arg1.V3.V3
									_ = __read_77
									__scalar_78 := int64(__let_scalar_69)
									_ = __scalar_78
									__read_79 := __arg3
									_ = __read_79
									__donor_slot_82 := __donor
									_ = __donor_slot_82
									__dead_80 := __arg1
									_ = __dead_80
									__dead_81 := __arg1.V3
									_ = __dead_81
									var __cell_83 *Constructor_Main_Branch
									if (__donor_slot_82) != (nil) {
										__cell_83 = __donor_slot_82
										__donor_slot_82 = nil
									} else {
										if (__dead_80) != (nil) {
											__cell_83 = __dead_80
											__dead_80 = nil
										} else {
											if (__dead_81) != (nil) {
												__cell_83 = __dead_81
												__dead_81 = nil
											} else {

											}
										}
									}
									if (__cell_83) == (nil) {
										__cell_83 = new(Constructor_Main_Branch)
									} else {

									}
									__cell_83.Rc = 1
									__cell_83.V0 = __scalar_71
									__cell_83.V1 = __read_72
									__cell_83.V2 = __scalar_73
									__cell_83.V3 = __read_74
									var __cell_84 *Constructor_Main_Branch
									if (__donor_slot_82) != (nil) {
										__cell_84 = __donor_slot_82
										__donor_slot_82 = nil
									} else {
										if (__dead_80) != (nil) {
											__cell_84 = __dead_80
											__dead_80 = nil
										} else {
											if (__dead_81) != (nil) {
												__cell_84 = __dead_81
												__dead_81 = nil
											} else {

											}
										}
									}
									if (__cell_84) == (nil) {
										__cell_84 = new(Constructor_Main_Branch)
									} else {

									}
									__cell_84.Rc = 1
									__cell_84.V0 = __scalar_76
									__cell_84.V1 = __read_77
									__cell_84.V2 = __scalar_78
									__cell_84.V3 = __read_79
									var __cell_85 *Constructor_Main_Branch
									if (__donor_slot_82) != (nil) {
										__cell_85 = __donor_slot_82
										__donor_slot_82 = nil
									} else {
										if (__dead_80) != (nil) {
											__cell_85 = __dead_80
											__dead_80 = nil
										} else {
											if (__dead_81) != (nil) {
												__cell_85 = __dead_81
												__dead_81 = nil
											} else {

											}
										}
									}
									if (__cell_85) == (nil) {
										__cell_85 = new(Constructor_Main_Branch)
									} else {

									}
									__cell_85.Rc = 1
									__cell_85.V0 = __scalar_70
									__cell_85.V1 = __cell_83
									__cell_85.V2 = __scalar_75
									__cell_85.V3 = __cell_84
									return __cell_85
								} else {
									if ((__arg3) != (nil)) && ((__arg3.V0) == (2247809753)) {
										if (__arg3.V1) != (nil) {
											if (__arg3.V1.V0) == (2247809753) {
												__let_scalar_102 := int64(__arg2)
												_ = __let_scalar_102
												__let_scalar_103 := int64(__arg3.V1.V2)
												_ = __let_scalar_103
												__let_scalar_104 := int64(__arg3.V2)
												_ = __let_scalar_104
												__scalar_105 := uint32(2247809753)
												_ = __scalar_105
												__scalar_106 := uint32(1685833310)
												_ = __scalar_106
												__read_107 := __arg1
												_ = __read_107
												__scalar_108 := int64(__let_scalar_102)
												_ = __scalar_108
												__read_109 := __arg3.V1.V1
												_ = __read_109
												__scalar_110 := int64(__let_scalar_103)
												_ = __scalar_110
												__scalar_111 := uint32(1685833310)
												_ = __scalar_111
												__read_112 := __arg3.V1.V3
												_ = __read_112
												__scalar_113 := int64(__let_scalar_104)
												_ = __scalar_113
												__read_114 := __arg3.V3
												_ = __read_114
												__donor_slot_117 := __donor
												_ = __donor_slot_117
												__dead_115 := __arg3
												_ = __dead_115
												__dead_116 := __arg3.V1
												_ = __dead_116
												var __cell_118 *Constructor_Main_Branch
												if (__donor_slot_117) != (nil) {
													__cell_118 = __donor_slot_117
													__donor_slot_117 = nil
												} else {
													if (__dead_115) != (nil) {
														__cell_118 = __dead_115
														__dead_115 = nil
													} else {
														if (__dead_116) != (nil) {
															__cell_118 = __dead_116
															__dead_116 = nil
														} else {

														}
													}
												}
												if (__cell_118) == (nil) {
													__cell_118 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_118.Rc = 1
												__cell_118.V0 = __scalar_106
												__cell_118.V1 = __read_107
												__cell_118.V2 = __scalar_108
												__cell_118.V3 = __read_109
												var __cell_119 *Constructor_Main_Branch
												if (__donor_slot_117) != (nil) {
													__cell_119 = __donor_slot_117
													__donor_slot_117 = nil
												} else {
													if (__dead_115) != (nil) {
														__cell_119 = __dead_115
														__dead_115 = nil
													} else {
														if (__dead_116) != (nil) {
															__cell_119 = __dead_116
															__dead_116 = nil
														} else {

														}
													}
												}
												if (__cell_119) == (nil) {
													__cell_119 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_119.Rc = 1
												__cell_119.V0 = __scalar_111
												__cell_119.V1 = __read_112
												__cell_119.V2 = __scalar_113
												__cell_119.V3 = __read_114
												var __cell_120 *Constructor_Main_Branch
												if (__donor_slot_117) != (nil) {
													__cell_120 = __donor_slot_117
													__donor_slot_117 = nil
												} else {
													if (__dead_115) != (nil) {
														__cell_120 = __dead_115
														__dead_115 = nil
													} else {
														if (__dead_116) != (nil) {
															__cell_120 = __dead_116
															__dead_116 = nil
														} else {

														}
													}
												}
												if (__cell_120) == (nil) {
													__cell_120 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_120.Rc = 1
												__cell_120.V0 = __scalar_105
												__cell_120.V1 = __cell_118
												__cell_120.V2 = __scalar_110
												__cell_120.V3 = __cell_119
												return __cell_120
											} else {
												if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (2247809753)) {
													__let_scalar_121 := int64(__arg2)
													_ = __let_scalar_121
													__let_scalar_122 := int64(__arg3.V2)
													_ = __let_scalar_122
													__let_scalar_123 := int64(__arg3.V3.V2)
													_ = __let_scalar_123
													__scalar_124 := uint32(2247809753)
													_ = __scalar_124
													__scalar_125 := uint32(1685833310)
													_ = __scalar_125
													__read_126 := __arg1
													_ = __read_126
													__scalar_127 := int64(__let_scalar_121)
													_ = __scalar_127
													__read_128 := __arg3.V1
													_ = __read_128
													__scalar_129 := int64(__let_scalar_122)
													_ = __scalar_129
													__scalar_130 := uint32(1685833310)
													_ = __scalar_130
													__read_131 := __arg3.V3.V1
													_ = __read_131
													__scalar_132 := int64(__let_scalar_123)
													_ = __scalar_132
													__read_133 := __arg3.V3.V3
													_ = __read_133
													__donor_slot_136 := __donor
													_ = __donor_slot_136
													__dead_134 := __arg3
													_ = __dead_134
													__dead_135 := __arg3.V3
													_ = __dead_135
													var __cell_137 *Constructor_Main_Branch
													if (__donor_slot_136) != (nil) {
														__cell_137 = __donor_slot_136
														__donor_slot_136 = nil
													} else {
														if (__dead_134) != (nil) {
															__cell_137 = __dead_134
															__dead_134 = nil
														} else {
															if (__dead_135) != (nil) {
																__cell_137 = __dead_135
																__dead_135 = nil
															} else {

															}
														}
													}
													if (__cell_137) == (nil) {
														__cell_137 = new(Constructor_Main_Branch)
													} else {

													}
													__cell_137.Rc = 1
													__cell_137.V0 = __scalar_125
													__cell_137.V1 = __read_126
													__cell_137.V2 = __scalar_127
													__cell_137.V3 = __read_128
													var __cell_138 *Constructor_Main_Branch
													if (__donor_slot_136) != (nil) {
														__cell_138 = __donor_slot_136
														__donor_slot_136 = nil
													} else {
														if (__dead_134) != (nil) {
															__cell_138 = __dead_134
															__dead_134 = nil
														} else {
															if (__dead_135) != (nil) {
																__cell_138 = __dead_135
																__dead_135 = nil
															} else {

															}
														}
													}
													if (__cell_138) == (nil) {
														__cell_138 = new(Constructor_Main_Branch)
													} else {

													}
													__cell_138.Rc = 1
													__cell_138.V0 = __scalar_130
													__cell_138.V1 = __read_131
													__cell_138.V2 = __scalar_132
													__cell_138.V3 = __read_133
													var __cell_139 *Constructor_Main_Branch
													if (__donor_slot_136) != (nil) {
														__cell_139 = __donor_slot_136
														__donor_slot_136 = nil
													} else {
														if (__dead_134) != (nil) {
															__cell_139 = __dead_134
															__dead_134 = nil
														} else {
															if (__dead_135) != (nil) {
																__cell_139 = __dead_135
																__dead_135 = nil
															} else {

															}
														}
													}
													if (__cell_139) == (nil) {
														__cell_139 = new(Constructor_Main_Branch)
													} else {

													}
													__cell_139.Rc = 1
													__cell_139.V0 = __scalar_124
													__cell_139.V1 = __cell_137
													__cell_139.V2 = __scalar_129
													__cell_139.V3 = __cell_138
													return __cell_139
												} else {
													__let_scalar_94 := uint32(__arg0)
													_ = __let_scalar_94
													__let_scalar_95 := int64(__arg2)
													_ = __let_scalar_95
													__scalar_96 := uint32(__let_scalar_94)
													_ = __scalar_96
													__read_97 := __arg1
													_ = __read_97
													__scalar_98 := int64(__let_scalar_95)
													_ = __scalar_98
													__read_99 := __arg3
													_ = __read_99
													__donor_slot_100 := __donor
													_ = __donor_slot_100
													var __cell_101 *Constructor_Main_Branch
													if (__donor_slot_100) != (nil) {
														__cell_101 = __donor_slot_100
														__donor_slot_100 = nil
													} else {

													}
													if (__cell_101) == (nil) {
														__cell_101 = new(Constructor_Main_Branch)
													} else {

													}
													__cell_101.Rc = 1
													__cell_101.V0 = __scalar_96
													__cell_101.V1 = __read_97
													__cell_101.V2 = __scalar_98
													__cell_101.V3 = __read_99
													return __cell_101
												}
											}
										} else {
											if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (2247809753)) {
												__let_scalar_140 := int64(__arg2)
												_ = __let_scalar_140
												__let_scalar_141 := int64(__arg3.V2)
												_ = __let_scalar_141
												__let_scalar_142 := int64(__arg3.V3.V2)
												_ = __let_scalar_142
												__scalar_143 := uint32(2247809753)
												_ = __scalar_143
												__scalar_144 := uint32(1685833310)
												_ = __scalar_144
												__read_145 := __arg1
												_ = __read_145
												__scalar_146 := int64(__let_scalar_140)
												_ = __scalar_146
												__read_147 := __arg3.V1
												_ = __read_147
												__scalar_148 := int64(__let_scalar_141)
												_ = __scalar_148
												__scalar_149 := uint32(1685833310)
												_ = __scalar_149
												__read_150 := __arg3.V3.V1
												_ = __read_150
												__scalar_151 := int64(__let_scalar_142)
												_ = __scalar_151
												__read_152 := __arg3.V3.V3
												_ = __read_152
												__donor_slot_155 := __donor
												_ = __donor_slot_155
												__dead_153 := __arg3
												_ = __dead_153
												__dead_154 := __arg3.V3
												_ = __dead_154
												var __cell_156 *Constructor_Main_Branch
												if (__donor_slot_155) != (nil) {
													__cell_156 = __donor_slot_155
													__donor_slot_155 = nil
												} else {
													if (__dead_153) != (nil) {
														__cell_156 = __dead_153
														__dead_153 = nil
													} else {
														if (__dead_154) != (nil) {
															__cell_156 = __dead_154
															__dead_154 = nil
														} else {

														}
													}
												}
												if (__cell_156) == (nil) {
													__cell_156 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_156.Rc = 1
												__cell_156.V0 = __scalar_144
												__cell_156.V1 = __read_145
												__cell_156.V2 = __scalar_146
												__cell_156.V3 = __read_147
												var __cell_157 *Constructor_Main_Branch
												if (__donor_slot_155) != (nil) {
													__cell_157 = __donor_slot_155
													__donor_slot_155 = nil
												} else {
													if (__dead_153) != (nil) {
														__cell_157 = __dead_153
														__dead_153 = nil
													} else {
														if (__dead_154) != (nil) {
															__cell_157 = __dead_154
															__dead_154 = nil
														} else {

														}
													}
												}
												if (__cell_157) == (nil) {
													__cell_157 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_157.Rc = 1
												__cell_157.V0 = __scalar_149
												__cell_157.V1 = __read_150
												__cell_157.V2 = __scalar_151
												__cell_157.V3 = __read_152
												var __cell_158 *Constructor_Main_Branch
												if (__donor_slot_155) != (nil) {
													__cell_158 = __donor_slot_155
													__donor_slot_155 = nil
												} else {
													if (__dead_153) != (nil) {
														__cell_158 = __dead_153
														__dead_153 = nil
													} else {
														if (__dead_154) != (nil) {
															__cell_158 = __dead_154
															__dead_154 = nil
														} else {

														}
													}
												}
												if (__cell_158) == (nil) {
													__cell_158 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_158.Rc = 1
												__cell_158.V0 = __scalar_143
												__cell_158.V1 = __cell_156
												__cell_158.V2 = __scalar_148
												__cell_158.V3 = __cell_157
												return __cell_158
											} else {
												__let_scalar_86 := uint32(__arg0)
												_ = __let_scalar_86
												__let_scalar_87 := int64(__arg2)
												_ = __let_scalar_87
												__scalar_88 := uint32(__let_scalar_86)
												_ = __scalar_88
												__read_89 := __arg1
												_ = __read_89
												__scalar_90 := int64(__let_scalar_87)
												_ = __scalar_90
												__read_91 := __arg3
												_ = __read_91
												__donor_slot_92 := __donor
												_ = __donor_slot_92
												var __cell_93 *Constructor_Main_Branch
												if (__donor_slot_92) != (nil) {
													__cell_93 = __donor_slot_92
													__donor_slot_92 = nil
												} else {

												}
												if (__cell_93) == (nil) {
													__cell_93 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_93.Rc = 1
												__cell_93.V0 = __scalar_88
												__cell_93.V1 = __read_89
												__cell_93.V2 = __scalar_90
												__cell_93.V3 = __read_91
												return __cell_93
											}
										}
									} else {
										__let_scalar_59 := uint32(__arg0)
										_ = __let_scalar_59
										__let_scalar_60 := int64(__arg2)
										_ = __let_scalar_60
										__scalar_61 := uint32(__let_scalar_59)
										_ = __scalar_61
										__read_62 := __arg1
										_ = __read_62
										__scalar_63 := int64(__let_scalar_60)
										_ = __scalar_63
										__read_64 := __arg3
										_ = __read_64
										__donor_slot_65 := __donor
										_ = __donor_slot_65
										var __cell_66 *Constructor_Main_Branch
										if (__donor_slot_65) != (nil) {
											__cell_66 = __donor_slot_65
											__donor_slot_65 = nil
										} else {

										}
										if (__cell_66) == (nil) {
											__cell_66 = new(Constructor_Main_Branch)
										} else {

										}
										__cell_66.Rc = 1
										__cell_66.V0 = __scalar_61
										__cell_66.V1 = __read_62
										__cell_66.V2 = __scalar_63
										__cell_66.V3 = __read_64
										return __cell_66
									}
								}
							} else {
								if ((__arg3) != (nil)) && ((__arg3.V0) == (2247809753)) {
									if (__arg3.V1) != (nil) {
										if (__arg3.V1.V0) == (2247809753) {
											__let_scalar_175 := int64(__arg2)
											_ = __let_scalar_175
											__let_scalar_176 := int64(__arg3.V1.V2)
											_ = __let_scalar_176
											__let_scalar_177 := int64(__arg3.V2)
											_ = __let_scalar_177
											__scalar_178 := uint32(2247809753)
											_ = __scalar_178
											__scalar_179 := uint32(1685833310)
											_ = __scalar_179
											__read_180 := __arg1
											_ = __read_180
											__scalar_181 := int64(__let_scalar_175)
											_ = __scalar_181
											__read_182 := __arg3.V1.V1
											_ = __read_182
											__scalar_183 := int64(__let_scalar_176)
											_ = __scalar_183
											__scalar_184 := uint32(1685833310)
											_ = __scalar_184
											__read_185 := __arg3.V1.V3
											_ = __read_185
											__scalar_186 := int64(__let_scalar_177)
											_ = __scalar_186
											__read_187 := __arg3.V3
											_ = __read_187
											__donor_slot_190 := __donor
											_ = __donor_slot_190
											__dead_188 := __arg3
											_ = __dead_188
											__dead_189 := __arg3.V1
											_ = __dead_189
											var __cell_191 *Constructor_Main_Branch
											if (__donor_slot_190) != (nil) {
												__cell_191 = __donor_slot_190
												__donor_slot_190 = nil
											} else {
												if (__dead_188) != (nil) {
													__cell_191 = __dead_188
													__dead_188 = nil
												} else {
													if (__dead_189) != (nil) {
														__cell_191 = __dead_189
														__dead_189 = nil
													} else {

													}
												}
											}
											if (__cell_191) == (nil) {
												__cell_191 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_191.Rc = 1
											__cell_191.V0 = __scalar_179
											__cell_191.V1 = __read_180
											__cell_191.V2 = __scalar_181
											__cell_191.V3 = __read_182
											var __cell_192 *Constructor_Main_Branch
											if (__donor_slot_190) != (nil) {
												__cell_192 = __donor_slot_190
												__donor_slot_190 = nil
											} else {
												if (__dead_188) != (nil) {
													__cell_192 = __dead_188
													__dead_188 = nil
												} else {
													if (__dead_189) != (nil) {
														__cell_192 = __dead_189
														__dead_189 = nil
													} else {

													}
												}
											}
											if (__cell_192) == (nil) {
												__cell_192 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_192.Rc = 1
											__cell_192.V0 = __scalar_184
											__cell_192.V1 = __read_185
											__cell_192.V2 = __scalar_186
											__cell_192.V3 = __read_187
											var __cell_193 *Constructor_Main_Branch
											if (__donor_slot_190) != (nil) {
												__cell_193 = __donor_slot_190
												__donor_slot_190 = nil
											} else {
												if (__dead_188) != (nil) {
													__cell_193 = __dead_188
													__dead_188 = nil
												} else {
													if (__dead_189) != (nil) {
														__cell_193 = __dead_189
														__dead_189 = nil
													} else {

													}
												}
											}
											if (__cell_193) == (nil) {
												__cell_193 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_193.Rc = 1
											__cell_193.V0 = __scalar_178
											__cell_193.V1 = __cell_191
											__cell_193.V2 = __scalar_183
											__cell_193.V3 = __cell_192
											return __cell_193
										} else {
											if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (2247809753)) {
												__let_scalar_194 := int64(__arg2)
												_ = __let_scalar_194
												__let_scalar_195 := int64(__arg3.V2)
												_ = __let_scalar_195
												__let_scalar_196 := int64(__arg3.V3.V2)
												_ = __let_scalar_196
												__scalar_197 := uint32(2247809753)
												_ = __scalar_197
												__scalar_198 := uint32(1685833310)
												_ = __scalar_198
												__read_199 := __arg1
												_ = __read_199
												__scalar_200 := int64(__let_scalar_194)
												_ = __scalar_200
												__read_201 := __arg3.V1
												_ = __read_201
												__scalar_202 := int64(__let_scalar_195)
												_ = __scalar_202
												__scalar_203 := uint32(1685833310)
												_ = __scalar_203
												__read_204 := __arg3.V3.V1
												_ = __read_204
												__scalar_205 := int64(__let_scalar_196)
												_ = __scalar_205
												__read_206 := __arg3.V3.V3
												_ = __read_206
												__donor_slot_209 := __donor
												_ = __donor_slot_209
												__dead_207 := __arg3
												_ = __dead_207
												__dead_208 := __arg3.V3
												_ = __dead_208
												var __cell_210 *Constructor_Main_Branch
												if (__donor_slot_209) != (nil) {
													__cell_210 = __donor_slot_209
													__donor_slot_209 = nil
												} else {
													if (__dead_207) != (nil) {
														__cell_210 = __dead_207
														__dead_207 = nil
													} else {
														if (__dead_208) != (nil) {
															__cell_210 = __dead_208
															__dead_208 = nil
														} else {

														}
													}
												}
												if (__cell_210) == (nil) {
													__cell_210 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_210.Rc = 1
												__cell_210.V0 = __scalar_198
												__cell_210.V1 = __read_199
												__cell_210.V2 = __scalar_200
												__cell_210.V3 = __read_201
												var __cell_211 *Constructor_Main_Branch
												if (__donor_slot_209) != (nil) {
													__cell_211 = __donor_slot_209
													__donor_slot_209 = nil
												} else {
													if (__dead_207) != (nil) {
														__cell_211 = __dead_207
														__dead_207 = nil
													} else {
														if (__dead_208) != (nil) {
															__cell_211 = __dead_208
															__dead_208 = nil
														} else {

														}
													}
												}
												if (__cell_211) == (nil) {
													__cell_211 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_211.Rc = 1
												__cell_211.V0 = __scalar_203
												__cell_211.V1 = __read_204
												__cell_211.V2 = __scalar_205
												__cell_211.V3 = __read_206
												var __cell_212 *Constructor_Main_Branch
												if (__donor_slot_209) != (nil) {
													__cell_212 = __donor_slot_209
													__donor_slot_209 = nil
												} else {
													if (__dead_207) != (nil) {
														__cell_212 = __dead_207
														__dead_207 = nil
													} else {
														if (__dead_208) != (nil) {
															__cell_212 = __dead_208
															__dead_208 = nil
														} else {

														}
													}
												}
												if (__cell_212) == (nil) {
													__cell_212 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_212.Rc = 1
												__cell_212.V0 = __scalar_197
												__cell_212.V1 = __cell_210
												__cell_212.V2 = __scalar_202
												__cell_212.V3 = __cell_211
												return __cell_212
											} else {
												__let_scalar_167 := uint32(__arg0)
												_ = __let_scalar_167
												__let_scalar_168 := int64(__arg2)
												_ = __let_scalar_168
												__scalar_169 := uint32(__let_scalar_167)
												_ = __scalar_169
												__read_170 := __arg1
												_ = __read_170
												__scalar_171 := int64(__let_scalar_168)
												_ = __scalar_171
												__read_172 := __arg3
												_ = __read_172
												__donor_slot_173 := __donor
												_ = __donor_slot_173
												var __cell_174 *Constructor_Main_Branch
												if (__donor_slot_173) != (nil) {
													__cell_174 = __donor_slot_173
													__donor_slot_173 = nil
												} else {

												}
												if (__cell_174) == (nil) {
													__cell_174 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_174.Rc = 1
												__cell_174.V0 = __scalar_169
												__cell_174.V1 = __read_170
												__cell_174.V2 = __scalar_171
												__cell_174.V3 = __read_172
												return __cell_174
											}
										}
									} else {
										if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (2247809753)) {
											__let_scalar_213 := int64(__arg2)
											_ = __let_scalar_213
											__let_scalar_214 := int64(__arg3.V2)
											_ = __let_scalar_214
											__let_scalar_215 := int64(__arg3.V3.V2)
											_ = __let_scalar_215
											__scalar_216 := uint32(2247809753)
											_ = __scalar_216
											__scalar_217 := uint32(1685833310)
											_ = __scalar_217
											__read_218 := __arg1
											_ = __read_218
											__scalar_219 := int64(__let_scalar_213)
											_ = __scalar_219
											__read_220 := __arg3.V1
											_ = __read_220
											__scalar_221 := int64(__let_scalar_214)
											_ = __scalar_221
											__scalar_222 := uint32(1685833310)
											_ = __scalar_222
											__read_223 := __arg3.V3.V1
											_ = __read_223
											__scalar_224 := int64(__let_scalar_215)
											_ = __scalar_224
											__read_225 := __arg3.V3.V3
											_ = __read_225
											__donor_slot_228 := __donor
											_ = __donor_slot_228
											__dead_226 := __arg3
											_ = __dead_226
											__dead_227 := __arg3.V3
											_ = __dead_227
											var __cell_229 *Constructor_Main_Branch
											if (__donor_slot_228) != (nil) {
												__cell_229 = __donor_slot_228
												__donor_slot_228 = nil
											} else {
												if (__dead_226) != (nil) {
													__cell_229 = __dead_226
													__dead_226 = nil
												} else {
													if (__dead_227) != (nil) {
														__cell_229 = __dead_227
														__dead_227 = nil
													} else {

													}
												}
											}
											if (__cell_229) == (nil) {
												__cell_229 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_229.Rc = 1
											__cell_229.V0 = __scalar_217
											__cell_229.V1 = __read_218
											__cell_229.V2 = __scalar_219
											__cell_229.V3 = __read_220
											var __cell_230 *Constructor_Main_Branch
											if (__donor_slot_228) != (nil) {
												__cell_230 = __donor_slot_228
												__donor_slot_228 = nil
											} else {
												if (__dead_226) != (nil) {
													__cell_230 = __dead_226
													__dead_226 = nil
												} else {
													if (__dead_227) != (nil) {
														__cell_230 = __dead_227
														__dead_227 = nil
													} else {

													}
												}
											}
											if (__cell_230) == (nil) {
												__cell_230 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_230.Rc = 1
											__cell_230.V0 = __scalar_222
											__cell_230.V1 = __read_223
											__cell_230.V2 = __scalar_224
											__cell_230.V3 = __read_225
											var __cell_231 *Constructor_Main_Branch
											if (__donor_slot_228) != (nil) {
												__cell_231 = __donor_slot_228
												__donor_slot_228 = nil
											} else {
												if (__dead_226) != (nil) {
													__cell_231 = __dead_226
													__dead_226 = nil
												} else {
													if (__dead_227) != (nil) {
														__cell_231 = __dead_227
														__dead_227 = nil
													} else {

													}
												}
											}
											if (__cell_231) == (nil) {
												__cell_231 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_231.Rc = 1
											__cell_231.V0 = __scalar_216
											__cell_231.V1 = __cell_229
											__cell_231.V2 = __scalar_221
											__cell_231.V3 = __cell_230
											return __cell_231
										} else {
											__let_scalar_159 := uint32(__arg0)
											_ = __let_scalar_159
											__let_scalar_160 := int64(__arg2)
											_ = __let_scalar_160
											__scalar_161 := uint32(__let_scalar_159)
											_ = __scalar_161
											__read_162 := __arg1
											_ = __read_162
											__scalar_163 := int64(__let_scalar_160)
											_ = __scalar_163
											__read_164 := __arg3
											_ = __read_164
											__donor_slot_165 := __donor
											_ = __donor_slot_165
											var __cell_166 *Constructor_Main_Branch
											if (__donor_slot_165) != (nil) {
												__cell_166 = __donor_slot_165
												__donor_slot_165 = nil
											} else {

											}
											if (__cell_166) == (nil) {
												__cell_166 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_166.Rc = 1
											__cell_166.V0 = __scalar_161
											__cell_166.V1 = __read_162
											__cell_166.V2 = __scalar_163
											__cell_166.V3 = __read_164
											return __cell_166
										}
									}
								} else {
									__let_scalar_32 := uint32(__arg0)
									_ = __let_scalar_32
									__let_scalar_33 := int64(__arg2)
									_ = __let_scalar_33
									__scalar_34 := uint32(__let_scalar_32)
									_ = __scalar_34
									__read_35 := __arg1
									_ = __read_35
									__scalar_36 := int64(__let_scalar_33)
									_ = __scalar_36
									__read_37 := __arg3
									_ = __read_37
									__donor_slot_38 := __donor
									_ = __donor_slot_38
									var __cell_39 *Constructor_Main_Branch
									if (__donor_slot_38) != (nil) {
										__cell_39 = __donor_slot_38
										__donor_slot_38 = nil
									} else {

									}
									if (__cell_39) == (nil) {
										__cell_39 = new(Constructor_Main_Branch)
									} else {

									}
									__cell_39.Rc = 1
									__cell_39.V0 = __scalar_34
									__cell_39.V1 = __read_35
									__cell_39.V2 = __scalar_36
									__cell_39.V3 = __read_37
									return __cell_39
								}
							}
						}
					} else {
						if (__arg1.V3) != (nil) {
							if (__arg1.V3.V0) == (2247809753) {
								__let_scalar_240 := int64(__arg1.V2)
								_ = __let_scalar_240
								__let_scalar_241 := int64(__arg1.V3.V2)
								_ = __let_scalar_241
								__let_scalar_242 := int64(__arg2)
								_ = __let_scalar_242
								__scalar_243 := uint32(2247809753)
								_ = __scalar_243
								__scalar_244 := uint32(1685833310)
								_ = __scalar_244
								__read_245 := __arg1.V1
								_ = __read_245
								__scalar_246 := int64(__let_scalar_240)
								_ = __scalar_246
								__read_247 := __arg1.V3.V1
								_ = __read_247
								__scalar_248 := int64(__let_scalar_241)
								_ = __scalar_248
								__scalar_249 := uint32(1685833310)
								_ = __scalar_249
								__read_250 := __arg1.V3.V3
								_ = __read_250
								__scalar_251 := int64(__let_scalar_242)
								_ = __scalar_251
								__read_252 := __arg3
								_ = __read_252
								__donor_slot_255 := __donor
								_ = __donor_slot_255
								__dead_253 := __arg1
								_ = __dead_253
								__dead_254 := __arg1.V3
								_ = __dead_254
								var __cell_256 *Constructor_Main_Branch
								if (__donor_slot_255) != (nil) {
									__cell_256 = __donor_slot_255
									__donor_slot_255 = nil
								} else {
									if (__dead_253) != (nil) {
										__cell_256 = __dead_253
										__dead_253 = nil
									} else {
										if (__dead_254) != (nil) {
											__cell_256 = __dead_254
											__dead_254 = nil
										} else {

										}
									}
								}
								if (__cell_256) == (nil) {
									__cell_256 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_256.Rc = 1
								__cell_256.V0 = __scalar_244
								__cell_256.V1 = __read_245
								__cell_256.V2 = __scalar_246
								__cell_256.V3 = __read_247
								var __cell_257 *Constructor_Main_Branch
								if (__donor_slot_255) != (nil) {
									__cell_257 = __donor_slot_255
									__donor_slot_255 = nil
								} else {
									if (__dead_253) != (nil) {
										__cell_257 = __dead_253
										__dead_253 = nil
									} else {
										if (__dead_254) != (nil) {
											__cell_257 = __dead_254
											__dead_254 = nil
										} else {

										}
									}
								}
								if (__cell_257) == (nil) {
									__cell_257 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_257.Rc = 1
								__cell_257.V0 = __scalar_249
								__cell_257.V1 = __read_250
								__cell_257.V2 = __scalar_251
								__cell_257.V3 = __read_252
								var __cell_258 *Constructor_Main_Branch
								if (__donor_slot_255) != (nil) {
									__cell_258 = __donor_slot_255
									__donor_slot_255 = nil
								} else {
									if (__dead_253) != (nil) {
										__cell_258 = __dead_253
										__dead_253 = nil
									} else {
										if (__dead_254) != (nil) {
											__cell_258 = __dead_254
											__dead_254 = nil
										} else {

										}
									}
								}
								if (__cell_258) == (nil) {
									__cell_258 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_258.Rc = 1
								__cell_258.V0 = __scalar_243
								__cell_258.V1 = __cell_256
								__cell_258.V2 = __scalar_248
								__cell_258.V3 = __cell_257
								return __cell_258
							} else {
								if ((__arg3) != (nil)) && ((__arg3.V0) == (2247809753)) {
									if (__arg3.V1) != (nil) {
										if (__arg3.V1.V0) == (2247809753) {
											__let_scalar_275 := int64(__arg2)
											_ = __let_scalar_275
											__let_scalar_276 := int64(__arg3.V1.V2)
											_ = __let_scalar_276
											__let_scalar_277 := int64(__arg3.V2)
											_ = __let_scalar_277
											__scalar_278 := uint32(2247809753)
											_ = __scalar_278
											__scalar_279 := uint32(1685833310)
											_ = __scalar_279
											__read_280 := __arg1
											_ = __read_280
											__scalar_281 := int64(__let_scalar_275)
											_ = __scalar_281
											__read_282 := __arg3.V1.V1
											_ = __read_282
											__scalar_283 := int64(__let_scalar_276)
											_ = __scalar_283
											__scalar_284 := uint32(1685833310)
											_ = __scalar_284
											__read_285 := __arg3.V1.V3
											_ = __read_285
											__scalar_286 := int64(__let_scalar_277)
											_ = __scalar_286
											__read_287 := __arg3.V3
											_ = __read_287
											__donor_slot_290 := __donor
											_ = __donor_slot_290
											__dead_288 := __arg3
											_ = __dead_288
											__dead_289 := __arg3.V1
											_ = __dead_289
											var __cell_291 *Constructor_Main_Branch
											if (__donor_slot_290) != (nil) {
												__cell_291 = __donor_slot_290
												__donor_slot_290 = nil
											} else {
												if (__dead_288) != (nil) {
													__cell_291 = __dead_288
													__dead_288 = nil
												} else {
													if (__dead_289) != (nil) {
														__cell_291 = __dead_289
														__dead_289 = nil
													} else {

													}
												}
											}
											if (__cell_291) == (nil) {
												__cell_291 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_291.Rc = 1
											__cell_291.V0 = __scalar_279
											__cell_291.V1 = __read_280
											__cell_291.V2 = __scalar_281
											__cell_291.V3 = __read_282
											var __cell_292 *Constructor_Main_Branch
											if (__donor_slot_290) != (nil) {
												__cell_292 = __donor_slot_290
												__donor_slot_290 = nil
											} else {
												if (__dead_288) != (nil) {
													__cell_292 = __dead_288
													__dead_288 = nil
												} else {
													if (__dead_289) != (nil) {
														__cell_292 = __dead_289
														__dead_289 = nil
													} else {

													}
												}
											}
											if (__cell_292) == (nil) {
												__cell_292 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_292.Rc = 1
											__cell_292.V0 = __scalar_284
											__cell_292.V1 = __read_285
											__cell_292.V2 = __scalar_286
											__cell_292.V3 = __read_287
											var __cell_293 *Constructor_Main_Branch
											if (__donor_slot_290) != (nil) {
												__cell_293 = __donor_slot_290
												__donor_slot_290 = nil
											} else {
												if (__dead_288) != (nil) {
													__cell_293 = __dead_288
													__dead_288 = nil
												} else {
													if (__dead_289) != (nil) {
														__cell_293 = __dead_289
														__dead_289 = nil
													} else {

													}
												}
											}
											if (__cell_293) == (nil) {
												__cell_293 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_293.Rc = 1
											__cell_293.V0 = __scalar_278
											__cell_293.V1 = __cell_291
											__cell_293.V2 = __scalar_283
											__cell_293.V3 = __cell_292
											return __cell_293
										} else {
											if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (2247809753)) {
												__let_scalar_294 := int64(__arg2)
												_ = __let_scalar_294
												__let_scalar_295 := int64(__arg3.V2)
												_ = __let_scalar_295
												__let_scalar_296 := int64(__arg3.V3.V2)
												_ = __let_scalar_296
												__scalar_297 := uint32(2247809753)
												_ = __scalar_297
												__scalar_298 := uint32(1685833310)
												_ = __scalar_298
												__read_299 := __arg1
												_ = __read_299
												__scalar_300 := int64(__let_scalar_294)
												_ = __scalar_300
												__read_301 := __arg3.V1
												_ = __read_301
												__scalar_302 := int64(__let_scalar_295)
												_ = __scalar_302
												__scalar_303 := uint32(1685833310)
												_ = __scalar_303
												__read_304 := __arg3.V3.V1
												_ = __read_304
												__scalar_305 := int64(__let_scalar_296)
												_ = __scalar_305
												__read_306 := __arg3.V3.V3
												_ = __read_306
												__donor_slot_309 := __donor
												_ = __donor_slot_309
												__dead_307 := __arg3
												_ = __dead_307
												__dead_308 := __arg3.V3
												_ = __dead_308
												var __cell_310 *Constructor_Main_Branch
												if (__donor_slot_309) != (nil) {
													__cell_310 = __donor_slot_309
													__donor_slot_309 = nil
												} else {
													if (__dead_307) != (nil) {
														__cell_310 = __dead_307
														__dead_307 = nil
													} else {
														if (__dead_308) != (nil) {
															__cell_310 = __dead_308
															__dead_308 = nil
														} else {

														}
													}
												}
												if (__cell_310) == (nil) {
													__cell_310 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_310.Rc = 1
												__cell_310.V0 = __scalar_298
												__cell_310.V1 = __read_299
												__cell_310.V2 = __scalar_300
												__cell_310.V3 = __read_301
												var __cell_311 *Constructor_Main_Branch
												if (__donor_slot_309) != (nil) {
													__cell_311 = __donor_slot_309
													__donor_slot_309 = nil
												} else {
													if (__dead_307) != (nil) {
														__cell_311 = __dead_307
														__dead_307 = nil
													} else {
														if (__dead_308) != (nil) {
															__cell_311 = __dead_308
															__dead_308 = nil
														} else {

														}
													}
												}
												if (__cell_311) == (nil) {
													__cell_311 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_311.Rc = 1
												__cell_311.V0 = __scalar_303
												__cell_311.V1 = __read_304
												__cell_311.V2 = __scalar_305
												__cell_311.V3 = __read_306
												var __cell_312 *Constructor_Main_Branch
												if (__donor_slot_309) != (nil) {
													__cell_312 = __donor_slot_309
													__donor_slot_309 = nil
												} else {
													if (__dead_307) != (nil) {
														__cell_312 = __dead_307
														__dead_307 = nil
													} else {
														if (__dead_308) != (nil) {
															__cell_312 = __dead_308
															__dead_308 = nil
														} else {

														}
													}
												}
												if (__cell_312) == (nil) {
													__cell_312 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_312.Rc = 1
												__cell_312.V0 = __scalar_297
												__cell_312.V1 = __cell_310
												__cell_312.V2 = __scalar_302
												__cell_312.V3 = __cell_311
												return __cell_312
											} else {
												__let_scalar_267 := uint32(__arg0)
												_ = __let_scalar_267
												__let_scalar_268 := int64(__arg2)
												_ = __let_scalar_268
												__scalar_269 := uint32(__let_scalar_267)
												_ = __scalar_269
												__read_270 := __arg1
												_ = __read_270
												__scalar_271 := int64(__let_scalar_268)
												_ = __scalar_271
												__read_272 := __arg3
												_ = __read_272
												__donor_slot_273 := __donor
												_ = __donor_slot_273
												var __cell_274 *Constructor_Main_Branch
												if (__donor_slot_273) != (nil) {
													__cell_274 = __donor_slot_273
													__donor_slot_273 = nil
												} else {

												}
												if (__cell_274) == (nil) {
													__cell_274 = new(Constructor_Main_Branch)
												} else {

												}
												__cell_274.Rc = 1
												__cell_274.V0 = __scalar_269
												__cell_274.V1 = __read_270
												__cell_274.V2 = __scalar_271
												__cell_274.V3 = __read_272
												return __cell_274
											}
										}
									} else {
										if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (2247809753)) {
											__let_scalar_313 := int64(__arg2)
											_ = __let_scalar_313
											__let_scalar_314 := int64(__arg3.V2)
											_ = __let_scalar_314
											__let_scalar_315 := int64(__arg3.V3.V2)
											_ = __let_scalar_315
											__scalar_316 := uint32(2247809753)
											_ = __scalar_316
											__scalar_317 := uint32(1685833310)
											_ = __scalar_317
											__read_318 := __arg1
											_ = __read_318
											__scalar_319 := int64(__let_scalar_313)
											_ = __scalar_319
											__read_320 := __arg3.V1
											_ = __read_320
											__scalar_321 := int64(__let_scalar_314)
											_ = __scalar_321
											__scalar_322 := uint32(1685833310)
											_ = __scalar_322
											__read_323 := __arg3.V3.V1
											_ = __read_323
											__scalar_324 := int64(__let_scalar_315)
											_ = __scalar_324
											__read_325 := __arg3.V3.V3
											_ = __read_325
											__donor_slot_328 := __donor
											_ = __donor_slot_328
											__dead_326 := __arg3
											_ = __dead_326
											__dead_327 := __arg3.V3
											_ = __dead_327
											var __cell_329 *Constructor_Main_Branch
											if (__donor_slot_328) != (nil) {
												__cell_329 = __donor_slot_328
												__donor_slot_328 = nil
											} else {
												if (__dead_326) != (nil) {
													__cell_329 = __dead_326
													__dead_326 = nil
												} else {
													if (__dead_327) != (nil) {
														__cell_329 = __dead_327
														__dead_327 = nil
													} else {

													}
												}
											}
											if (__cell_329) == (nil) {
												__cell_329 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_329.Rc = 1
											__cell_329.V0 = __scalar_317
											__cell_329.V1 = __read_318
											__cell_329.V2 = __scalar_319
											__cell_329.V3 = __read_320
											var __cell_330 *Constructor_Main_Branch
											if (__donor_slot_328) != (nil) {
												__cell_330 = __donor_slot_328
												__donor_slot_328 = nil
											} else {
												if (__dead_326) != (nil) {
													__cell_330 = __dead_326
													__dead_326 = nil
												} else {
													if (__dead_327) != (nil) {
														__cell_330 = __dead_327
														__dead_327 = nil
													} else {

													}
												}
											}
											if (__cell_330) == (nil) {
												__cell_330 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_330.Rc = 1
											__cell_330.V0 = __scalar_322
											__cell_330.V1 = __read_323
											__cell_330.V2 = __scalar_324
											__cell_330.V3 = __read_325
											var __cell_331 *Constructor_Main_Branch
											if (__donor_slot_328) != (nil) {
												__cell_331 = __donor_slot_328
												__donor_slot_328 = nil
											} else {
												if (__dead_326) != (nil) {
													__cell_331 = __dead_326
													__dead_326 = nil
												} else {
													if (__dead_327) != (nil) {
														__cell_331 = __dead_327
														__dead_327 = nil
													} else {

													}
												}
											}
											if (__cell_331) == (nil) {
												__cell_331 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_331.Rc = 1
											__cell_331.V0 = __scalar_316
											__cell_331.V1 = __cell_329
											__cell_331.V2 = __scalar_321
											__cell_331.V3 = __cell_330
											return __cell_331
										} else {
											__let_scalar_259 := uint32(__arg0)
											_ = __let_scalar_259
											__let_scalar_260 := int64(__arg2)
											_ = __let_scalar_260
											__scalar_261 := uint32(__let_scalar_259)
											_ = __scalar_261
											__read_262 := __arg1
											_ = __read_262
											__scalar_263 := int64(__let_scalar_260)
											_ = __scalar_263
											__read_264 := __arg3
											_ = __read_264
											__donor_slot_265 := __donor
											_ = __donor_slot_265
											var __cell_266 *Constructor_Main_Branch
											if (__donor_slot_265) != (nil) {
												__cell_266 = __donor_slot_265
												__donor_slot_265 = nil
											} else {

											}
											if (__cell_266) == (nil) {
												__cell_266 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_266.Rc = 1
											__cell_266.V0 = __scalar_261
											__cell_266.V1 = __read_262
											__cell_266.V2 = __scalar_263
											__cell_266.V3 = __read_264
											return __cell_266
										}
									}
								} else {
									__let_scalar_232 := uint32(__arg0)
									_ = __let_scalar_232
									__let_scalar_233 := int64(__arg2)
									_ = __let_scalar_233
									__scalar_234 := uint32(__let_scalar_232)
									_ = __scalar_234
									__read_235 := __arg1
									_ = __read_235
									__scalar_236 := int64(__let_scalar_233)
									_ = __scalar_236
									__read_237 := __arg3
									_ = __read_237
									__donor_slot_238 := __donor
									_ = __donor_slot_238
									var __cell_239 *Constructor_Main_Branch
									if (__donor_slot_238) != (nil) {
										__cell_239 = __donor_slot_238
										__donor_slot_238 = nil
									} else {

									}
									if (__cell_239) == (nil) {
										__cell_239 = new(Constructor_Main_Branch)
									} else {

									}
									__cell_239.Rc = 1
									__cell_239.V0 = __scalar_234
									__cell_239.V1 = __read_235
									__cell_239.V2 = __scalar_236
									__cell_239.V3 = __read_237
									return __cell_239
								}
							}
						} else {
							if ((__arg3) != (nil)) && ((__arg3.V0) == (2247809753)) {
								if (__arg3.V1) != (nil) {
									if (__arg3.V1.V0) == (2247809753) {
										__let_scalar_348 := int64(__arg2)
										_ = __let_scalar_348
										__let_scalar_349 := int64(__arg3.V1.V2)
										_ = __let_scalar_349
										__let_scalar_350 := int64(__arg3.V2)
										_ = __let_scalar_350
										__scalar_351 := uint32(2247809753)
										_ = __scalar_351
										__scalar_352 := uint32(1685833310)
										_ = __scalar_352
										__read_353 := __arg1
										_ = __read_353
										__scalar_354 := int64(__let_scalar_348)
										_ = __scalar_354
										__read_355 := __arg3.V1.V1
										_ = __read_355
										__scalar_356 := int64(__let_scalar_349)
										_ = __scalar_356
										__scalar_357 := uint32(1685833310)
										_ = __scalar_357
										__read_358 := __arg3.V1.V3
										_ = __read_358
										__scalar_359 := int64(__let_scalar_350)
										_ = __scalar_359
										__read_360 := __arg3.V3
										_ = __read_360
										__donor_slot_363 := __donor
										_ = __donor_slot_363
										__dead_361 := __arg3
										_ = __dead_361
										__dead_362 := __arg3.V1
										_ = __dead_362
										var __cell_364 *Constructor_Main_Branch
										if (__donor_slot_363) != (nil) {
											__cell_364 = __donor_slot_363
											__donor_slot_363 = nil
										} else {
											if (__dead_361) != (nil) {
												__cell_364 = __dead_361
												__dead_361 = nil
											} else {
												if (__dead_362) != (nil) {
													__cell_364 = __dead_362
													__dead_362 = nil
												} else {

												}
											}
										}
										if (__cell_364) == (nil) {
											__cell_364 = new(Constructor_Main_Branch)
										} else {

										}
										__cell_364.Rc = 1
										__cell_364.V0 = __scalar_352
										__cell_364.V1 = __read_353
										__cell_364.V2 = __scalar_354
										__cell_364.V3 = __read_355
										var __cell_365 *Constructor_Main_Branch
										if (__donor_slot_363) != (nil) {
											__cell_365 = __donor_slot_363
											__donor_slot_363 = nil
										} else {
											if (__dead_361) != (nil) {
												__cell_365 = __dead_361
												__dead_361 = nil
											} else {
												if (__dead_362) != (nil) {
													__cell_365 = __dead_362
													__dead_362 = nil
												} else {

												}
											}
										}
										if (__cell_365) == (nil) {
											__cell_365 = new(Constructor_Main_Branch)
										} else {

										}
										__cell_365.Rc = 1
										__cell_365.V0 = __scalar_357
										__cell_365.V1 = __read_358
										__cell_365.V2 = __scalar_359
										__cell_365.V3 = __read_360
										var __cell_366 *Constructor_Main_Branch
										if (__donor_slot_363) != (nil) {
											__cell_366 = __donor_slot_363
											__donor_slot_363 = nil
										} else {
											if (__dead_361) != (nil) {
												__cell_366 = __dead_361
												__dead_361 = nil
											} else {
												if (__dead_362) != (nil) {
													__cell_366 = __dead_362
													__dead_362 = nil
												} else {

												}
											}
										}
										if (__cell_366) == (nil) {
											__cell_366 = new(Constructor_Main_Branch)
										} else {

										}
										__cell_366.Rc = 1
										__cell_366.V0 = __scalar_351
										__cell_366.V1 = __cell_364
										__cell_366.V2 = __scalar_356
										__cell_366.V3 = __cell_365
										return __cell_366
									} else {
										if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (2247809753)) {
											__let_scalar_367 := int64(__arg2)
											_ = __let_scalar_367
											__let_scalar_368 := int64(__arg3.V2)
											_ = __let_scalar_368
											__let_scalar_369 := int64(__arg3.V3.V2)
											_ = __let_scalar_369
											__scalar_370 := uint32(2247809753)
											_ = __scalar_370
											__scalar_371 := uint32(1685833310)
											_ = __scalar_371
											__read_372 := __arg1
											_ = __read_372
											__scalar_373 := int64(__let_scalar_367)
											_ = __scalar_373
											__read_374 := __arg3.V1
											_ = __read_374
											__scalar_375 := int64(__let_scalar_368)
											_ = __scalar_375
											__scalar_376 := uint32(1685833310)
											_ = __scalar_376
											__read_377 := __arg3.V3.V1
											_ = __read_377
											__scalar_378 := int64(__let_scalar_369)
											_ = __scalar_378
											__read_379 := __arg3.V3.V3
											_ = __read_379
											__donor_slot_382 := __donor
											_ = __donor_slot_382
											__dead_380 := __arg3
											_ = __dead_380
											__dead_381 := __arg3.V3
											_ = __dead_381
											var __cell_383 *Constructor_Main_Branch
											if (__donor_slot_382) != (nil) {
												__cell_383 = __donor_slot_382
												__donor_slot_382 = nil
											} else {
												if (__dead_380) != (nil) {
													__cell_383 = __dead_380
													__dead_380 = nil
												} else {
													if (__dead_381) != (nil) {
														__cell_383 = __dead_381
														__dead_381 = nil
													} else {

													}
												}
											}
											if (__cell_383) == (nil) {
												__cell_383 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_383.Rc = 1
											__cell_383.V0 = __scalar_371
											__cell_383.V1 = __read_372
											__cell_383.V2 = __scalar_373
											__cell_383.V3 = __read_374
											var __cell_384 *Constructor_Main_Branch
											if (__donor_slot_382) != (nil) {
												__cell_384 = __donor_slot_382
												__donor_slot_382 = nil
											} else {
												if (__dead_380) != (nil) {
													__cell_384 = __dead_380
													__dead_380 = nil
												} else {
													if (__dead_381) != (nil) {
														__cell_384 = __dead_381
														__dead_381 = nil
													} else {

													}
												}
											}
											if (__cell_384) == (nil) {
												__cell_384 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_384.Rc = 1
											__cell_384.V0 = __scalar_376
											__cell_384.V1 = __read_377
											__cell_384.V2 = __scalar_378
											__cell_384.V3 = __read_379
											var __cell_385 *Constructor_Main_Branch
											if (__donor_slot_382) != (nil) {
												__cell_385 = __donor_slot_382
												__donor_slot_382 = nil
											} else {
												if (__dead_380) != (nil) {
													__cell_385 = __dead_380
													__dead_380 = nil
												} else {
													if (__dead_381) != (nil) {
														__cell_385 = __dead_381
														__dead_381 = nil
													} else {

													}
												}
											}
											if (__cell_385) == (nil) {
												__cell_385 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_385.Rc = 1
											__cell_385.V0 = __scalar_370
											__cell_385.V1 = __cell_383
											__cell_385.V2 = __scalar_375
											__cell_385.V3 = __cell_384
											return __cell_385
										} else {
											__let_scalar_340 := uint32(__arg0)
											_ = __let_scalar_340
											__let_scalar_341 := int64(__arg2)
											_ = __let_scalar_341
											__scalar_342 := uint32(__let_scalar_340)
											_ = __scalar_342
											__read_343 := __arg1
											_ = __read_343
											__scalar_344 := int64(__let_scalar_341)
											_ = __scalar_344
											__read_345 := __arg3
											_ = __read_345
											__donor_slot_346 := __donor
											_ = __donor_slot_346
											var __cell_347 *Constructor_Main_Branch
											if (__donor_slot_346) != (nil) {
												__cell_347 = __donor_slot_346
												__donor_slot_346 = nil
											} else {

											}
											if (__cell_347) == (nil) {
												__cell_347 = new(Constructor_Main_Branch)
											} else {

											}
											__cell_347.Rc = 1
											__cell_347.V0 = __scalar_342
											__cell_347.V1 = __read_343
											__cell_347.V2 = __scalar_344
											__cell_347.V3 = __read_345
											return __cell_347
										}
									}
								} else {
									if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (2247809753)) {
										__let_scalar_386 := int64(__arg2)
										_ = __let_scalar_386
										__let_scalar_387 := int64(__arg3.V2)
										_ = __let_scalar_387
										__let_scalar_388 := int64(__arg3.V3.V2)
										_ = __let_scalar_388
										__scalar_389 := uint32(2247809753)
										_ = __scalar_389
										__scalar_390 := uint32(1685833310)
										_ = __scalar_390
										__read_391 := __arg1
										_ = __read_391
										__scalar_392 := int64(__let_scalar_386)
										_ = __scalar_392
										__read_393 := __arg3.V1
										_ = __read_393
										__scalar_394 := int64(__let_scalar_387)
										_ = __scalar_394
										__scalar_395 := uint32(1685833310)
										_ = __scalar_395
										__read_396 := __arg3.V3.V1
										_ = __read_396
										__scalar_397 := int64(__let_scalar_388)
										_ = __scalar_397
										__read_398 := __arg3.V3.V3
										_ = __read_398
										__donor_slot_401 := __donor
										_ = __donor_slot_401
										__dead_399 := __arg3
										_ = __dead_399
										__dead_400 := __arg3.V3
										_ = __dead_400
										var __cell_402 *Constructor_Main_Branch
										if (__donor_slot_401) != (nil) {
											__cell_402 = __donor_slot_401
											__donor_slot_401 = nil
										} else {
											if (__dead_399) != (nil) {
												__cell_402 = __dead_399
												__dead_399 = nil
											} else {
												if (__dead_400) != (nil) {
													__cell_402 = __dead_400
													__dead_400 = nil
												} else {

												}
											}
										}
										if (__cell_402) == (nil) {
											__cell_402 = new(Constructor_Main_Branch)
										} else {

										}
										__cell_402.Rc = 1
										__cell_402.V0 = __scalar_390
										__cell_402.V1 = __read_391
										__cell_402.V2 = __scalar_392
										__cell_402.V3 = __read_393
										var __cell_403 *Constructor_Main_Branch
										if (__donor_slot_401) != (nil) {
											__cell_403 = __donor_slot_401
											__donor_slot_401 = nil
										} else {
											if (__dead_399) != (nil) {
												__cell_403 = __dead_399
												__dead_399 = nil
											} else {
												if (__dead_400) != (nil) {
													__cell_403 = __dead_400
													__dead_400 = nil
												} else {

												}
											}
										}
										if (__cell_403) == (nil) {
											__cell_403 = new(Constructor_Main_Branch)
										} else {

										}
										__cell_403.Rc = 1
										__cell_403.V0 = __scalar_395
										__cell_403.V1 = __read_396
										__cell_403.V2 = __scalar_397
										__cell_403.V3 = __read_398
										var __cell_404 *Constructor_Main_Branch
										if (__donor_slot_401) != (nil) {
											__cell_404 = __donor_slot_401
											__donor_slot_401 = nil
										} else {
											if (__dead_399) != (nil) {
												__cell_404 = __dead_399
												__dead_399 = nil
											} else {
												if (__dead_400) != (nil) {
													__cell_404 = __dead_400
													__dead_400 = nil
												} else {

												}
											}
										}
										if (__cell_404) == (nil) {
											__cell_404 = new(Constructor_Main_Branch)
										} else {

										}
										__cell_404.Rc = 1
										__cell_404.V0 = __scalar_389
										__cell_404.V1 = __cell_402
										__cell_404.V2 = __scalar_394
										__cell_404.V3 = __cell_403
										return __cell_404
									} else {
										__let_scalar_332 := uint32(__arg0)
										_ = __let_scalar_332
										__let_scalar_333 := int64(__arg2)
										_ = __let_scalar_333
										__scalar_334 := uint32(__let_scalar_332)
										_ = __scalar_334
										__read_335 := __arg1
										_ = __read_335
										__scalar_336 := int64(__let_scalar_333)
										_ = __scalar_336
										__read_337 := __arg3
										_ = __read_337
										__donor_slot_338 := __donor
										_ = __donor_slot_338
										var __cell_339 *Constructor_Main_Branch
										if (__donor_slot_338) != (nil) {
											__cell_339 = __donor_slot_338
											__donor_slot_338 = nil
										} else {

										}
										if (__cell_339) == (nil) {
											__cell_339 = new(Constructor_Main_Branch)
										} else {

										}
										__cell_339.Rc = 1
										__cell_339.V0 = __scalar_334
										__cell_339.V1 = __read_335
										__cell_339.V2 = __scalar_336
										__cell_339.V3 = __read_337
										return __cell_339
									}
								}
							} else {
								__let_scalar_24 := uint32(__arg0)
								_ = __let_scalar_24
								__let_scalar_25 := int64(__arg2)
								_ = __let_scalar_25
								__scalar_26 := uint32(__let_scalar_24)
								_ = __scalar_26
								__read_27 := __arg1
								_ = __read_27
								__scalar_28 := int64(__let_scalar_25)
								_ = __scalar_28
								__read_29 := __arg3
								_ = __read_29
								__donor_slot_30 := __donor
								_ = __donor_slot_30
								var __cell_31 *Constructor_Main_Branch
								if (__donor_slot_30) != (nil) {
									__cell_31 = __donor_slot_30
									__donor_slot_30 = nil
								} else {

								}
								if (__cell_31) == (nil) {
									__cell_31 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_31.Rc = 1
								__cell_31.V0 = __scalar_26
								__cell_31.V1 = __read_27
								__cell_31.V2 = __scalar_28
								__cell_31.V3 = __read_29
								return __cell_31
							}
						}
					}
				} else {
					if ((__arg3) != (nil)) && ((__arg3.V0) == (2247809753)) {
						if (__arg3.V1) != (nil) {
							if (__arg3.V1.V0) == (2247809753) {
								__let_scalar_421 := int64(__arg2)
								_ = __let_scalar_421
								__let_scalar_422 := int64(__arg3.V1.V2)
								_ = __let_scalar_422
								__let_scalar_423 := int64(__arg3.V2)
								_ = __let_scalar_423
								__scalar_424 := uint32(2247809753)
								_ = __scalar_424
								__scalar_425 := uint32(1685833310)
								_ = __scalar_425
								__read_426 := __arg1
								_ = __read_426
								__scalar_427 := int64(__let_scalar_421)
								_ = __scalar_427
								__read_428 := __arg3.V1.V1
								_ = __read_428
								__scalar_429 := int64(__let_scalar_422)
								_ = __scalar_429
								__scalar_430 := uint32(1685833310)
								_ = __scalar_430
								__read_431 := __arg3.V1.V3
								_ = __read_431
								__scalar_432 := int64(__let_scalar_423)
								_ = __scalar_432
								__read_433 := __arg3.V3
								_ = __read_433
								__donor_slot_436 := __donor
								_ = __donor_slot_436
								__dead_434 := __arg3
								_ = __dead_434
								__dead_435 := __arg3.V1
								_ = __dead_435
								var __cell_437 *Constructor_Main_Branch
								if (__donor_slot_436) != (nil) {
									__cell_437 = __donor_slot_436
									__donor_slot_436 = nil
								} else {
									if (__dead_434) != (nil) {
										__cell_437 = __dead_434
										__dead_434 = nil
									} else {
										if (__dead_435) != (nil) {
											__cell_437 = __dead_435
											__dead_435 = nil
										} else {

										}
									}
								}
								if (__cell_437) == (nil) {
									__cell_437 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_437.Rc = 1
								__cell_437.V0 = __scalar_425
								__cell_437.V1 = __read_426
								__cell_437.V2 = __scalar_427
								__cell_437.V3 = __read_428
								var __cell_438 *Constructor_Main_Branch
								if (__donor_slot_436) != (nil) {
									__cell_438 = __donor_slot_436
									__donor_slot_436 = nil
								} else {
									if (__dead_434) != (nil) {
										__cell_438 = __dead_434
										__dead_434 = nil
									} else {
										if (__dead_435) != (nil) {
											__cell_438 = __dead_435
											__dead_435 = nil
										} else {

										}
									}
								}
								if (__cell_438) == (nil) {
									__cell_438 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_438.Rc = 1
								__cell_438.V0 = __scalar_430
								__cell_438.V1 = __read_431
								__cell_438.V2 = __scalar_432
								__cell_438.V3 = __read_433
								var __cell_439 *Constructor_Main_Branch
								if (__donor_slot_436) != (nil) {
									__cell_439 = __donor_slot_436
									__donor_slot_436 = nil
								} else {
									if (__dead_434) != (nil) {
										__cell_439 = __dead_434
										__dead_434 = nil
									} else {
										if (__dead_435) != (nil) {
											__cell_439 = __dead_435
											__dead_435 = nil
										} else {

										}
									}
								}
								if (__cell_439) == (nil) {
									__cell_439 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_439.Rc = 1
								__cell_439.V0 = __scalar_424
								__cell_439.V1 = __cell_437
								__cell_439.V2 = __scalar_429
								__cell_439.V3 = __cell_438
								return __cell_439
							} else {
								if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (2247809753)) {
									__let_scalar_440 := int64(__arg2)
									_ = __let_scalar_440
									__let_scalar_441 := int64(__arg3.V2)
									_ = __let_scalar_441
									__let_scalar_442 := int64(__arg3.V3.V2)
									_ = __let_scalar_442
									__scalar_443 := uint32(2247809753)
									_ = __scalar_443
									__scalar_444 := uint32(1685833310)
									_ = __scalar_444
									__read_445 := __arg1
									_ = __read_445
									__scalar_446 := int64(__let_scalar_440)
									_ = __scalar_446
									__read_447 := __arg3.V1
									_ = __read_447
									__scalar_448 := int64(__let_scalar_441)
									_ = __scalar_448
									__scalar_449 := uint32(1685833310)
									_ = __scalar_449
									__read_450 := __arg3.V3.V1
									_ = __read_450
									__scalar_451 := int64(__let_scalar_442)
									_ = __scalar_451
									__read_452 := __arg3.V3.V3
									_ = __read_452
									__donor_slot_455 := __donor
									_ = __donor_slot_455
									__dead_453 := __arg3
									_ = __dead_453
									__dead_454 := __arg3.V3
									_ = __dead_454
									var __cell_456 *Constructor_Main_Branch
									if (__donor_slot_455) != (nil) {
										__cell_456 = __donor_slot_455
										__donor_slot_455 = nil
									} else {
										if (__dead_453) != (nil) {
											__cell_456 = __dead_453
											__dead_453 = nil
										} else {
											if (__dead_454) != (nil) {
												__cell_456 = __dead_454
												__dead_454 = nil
											} else {

											}
										}
									}
									if (__cell_456) == (nil) {
										__cell_456 = new(Constructor_Main_Branch)
									} else {

									}
									__cell_456.Rc = 1
									__cell_456.V0 = __scalar_444
									__cell_456.V1 = __read_445
									__cell_456.V2 = __scalar_446
									__cell_456.V3 = __read_447
									var __cell_457 *Constructor_Main_Branch
									if (__donor_slot_455) != (nil) {
										__cell_457 = __donor_slot_455
										__donor_slot_455 = nil
									} else {
										if (__dead_453) != (nil) {
											__cell_457 = __dead_453
											__dead_453 = nil
										} else {
											if (__dead_454) != (nil) {
												__cell_457 = __dead_454
												__dead_454 = nil
											} else {

											}
										}
									}
									if (__cell_457) == (nil) {
										__cell_457 = new(Constructor_Main_Branch)
									} else {

									}
									__cell_457.Rc = 1
									__cell_457.V0 = __scalar_449
									__cell_457.V1 = __read_450
									__cell_457.V2 = __scalar_451
									__cell_457.V3 = __read_452
									var __cell_458 *Constructor_Main_Branch
									if (__donor_slot_455) != (nil) {
										__cell_458 = __donor_slot_455
										__donor_slot_455 = nil
									} else {
										if (__dead_453) != (nil) {
											__cell_458 = __dead_453
											__dead_453 = nil
										} else {
											if (__dead_454) != (nil) {
												__cell_458 = __dead_454
												__dead_454 = nil
											} else {

											}
										}
									}
									if (__cell_458) == (nil) {
										__cell_458 = new(Constructor_Main_Branch)
									} else {

									}
									__cell_458.Rc = 1
									__cell_458.V0 = __scalar_443
									__cell_458.V1 = __cell_456
									__cell_458.V2 = __scalar_448
									__cell_458.V3 = __cell_457
									return __cell_458
								} else {
									__let_scalar_413 := uint32(__arg0)
									_ = __let_scalar_413
									__let_scalar_414 := int64(__arg2)
									_ = __let_scalar_414
									__scalar_415 := uint32(__let_scalar_413)
									_ = __scalar_415
									__read_416 := __arg1
									_ = __read_416
									__scalar_417 := int64(__let_scalar_414)
									_ = __scalar_417
									__read_418 := __arg3
									_ = __read_418
									__donor_slot_419 := __donor
									_ = __donor_slot_419
									var __cell_420 *Constructor_Main_Branch
									if (__donor_slot_419) != (nil) {
										__cell_420 = __donor_slot_419
										__donor_slot_419 = nil
									} else {

									}
									if (__cell_420) == (nil) {
										__cell_420 = new(Constructor_Main_Branch)
									} else {

									}
									__cell_420.Rc = 1
									__cell_420.V0 = __scalar_415
									__cell_420.V1 = __read_416
									__cell_420.V2 = __scalar_417
									__cell_420.V3 = __read_418
									return __cell_420
								}
							}
						} else {
							if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (2247809753)) {
								__let_scalar_459 := int64(__arg2)
								_ = __let_scalar_459
								__let_scalar_460 := int64(__arg3.V2)
								_ = __let_scalar_460
								__let_scalar_461 := int64(__arg3.V3.V2)
								_ = __let_scalar_461
								__scalar_462 := uint32(2247809753)
								_ = __scalar_462
								__scalar_463 := uint32(1685833310)
								_ = __scalar_463
								__read_464 := __arg1
								_ = __read_464
								__scalar_465 := int64(__let_scalar_459)
								_ = __scalar_465
								__read_466 := __arg3.V1
								_ = __read_466
								__scalar_467 := int64(__let_scalar_460)
								_ = __scalar_467
								__scalar_468 := uint32(1685833310)
								_ = __scalar_468
								__read_469 := __arg3.V3.V1
								_ = __read_469
								__scalar_470 := int64(__let_scalar_461)
								_ = __scalar_470
								__read_471 := __arg3.V3.V3
								_ = __read_471
								__donor_slot_474 := __donor
								_ = __donor_slot_474
								__dead_472 := __arg3
								_ = __dead_472
								__dead_473 := __arg3.V3
								_ = __dead_473
								var __cell_475 *Constructor_Main_Branch
								if (__donor_slot_474) != (nil) {
									__cell_475 = __donor_slot_474
									__donor_slot_474 = nil
								} else {
									if (__dead_472) != (nil) {
										__cell_475 = __dead_472
										__dead_472 = nil
									} else {
										if (__dead_473) != (nil) {
											__cell_475 = __dead_473
											__dead_473 = nil
										} else {

										}
									}
								}
								if (__cell_475) == (nil) {
									__cell_475 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_475.Rc = 1
								__cell_475.V0 = __scalar_463
								__cell_475.V1 = __read_464
								__cell_475.V2 = __scalar_465
								__cell_475.V3 = __read_466
								var __cell_476 *Constructor_Main_Branch
								if (__donor_slot_474) != (nil) {
									__cell_476 = __donor_slot_474
									__donor_slot_474 = nil
								} else {
									if (__dead_472) != (nil) {
										__cell_476 = __dead_472
										__dead_472 = nil
									} else {
										if (__dead_473) != (nil) {
											__cell_476 = __dead_473
											__dead_473 = nil
										} else {

										}
									}
								}
								if (__cell_476) == (nil) {
									__cell_476 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_476.Rc = 1
								__cell_476.V0 = __scalar_468
								__cell_476.V1 = __read_469
								__cell_476.V2 = __scalar_470
								__cell_476.V3 = __read_471
								var __cell_477 *Constructor_Main_Branch
								if (__donor_slot_474) != (nil) {
									__cell_477 = __donor_slot_474
									__donor_slot_474 = nil
								} else {
									if (__dead_472) != (nil) {
										__cell_477 = __dead_472
										__dead_472 = nil
									} else {
										if (__dead_473) != (nil) {
											__cell_477 = __dead_473
											__dead_473 = nil
										} else {

										}
									}
								}
								if (__cell_477) == (nil) {
									__cell_477 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_477.Rc = 1
								__cell_477.V0 = __scalar_462
								__cell_477.V1 = __cell_475
								__cell_477.V2 = __scalar_467
								__cell_477.V3 = __cell_476
								return __cell_477
							} else {
								__let_scalar_405 := uint32(__arg0)
								_ = __let_scalar_405
								__let_scalar_406 := int64(__arg2)
								_ = __let_scalar_406
								__scalar_407 := uint32(__let_scalar_405)
								_ = __scalar_407
								__read_408 := __arg1
								_ = __read_408
								__scalar_409 := int64(__let_scalar_406)
								_ = __scalar_409
								__read_410 := __arg3
								_ = __read_410
								__donor_slot_411 := __donor
								_ = __donor_slot_411
								var __cell_412 *Constructor_Main_Branch
								if (__donor_slot_411) != (nil) {
									__cell_412 = __donor_slot_411
									__donor_slot_411 = nil
								} else {

								}
								if (__cell_412) == (nil) {
									__cell_412 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_412.Rc = 1
								__cell_412.V0 = __scalar_407
								__cell_412.V1 = __read_408
								__cell_412.V2 = __scalar_409
								__cell_412.V3 = __read_410
								return __cell_412
							}
						}
					} else {
						__let_scalar_16 := uint32(__arg0)
						_ = __let_scalar_16
						__let_scalar_17 := int64(__arg2)
						_ = __let_scalar_17
						__scalar_18 := uint32(__let_scalar_16)
						_ = __scalar_18
						__read_19 := __arg1
						_ = __read_19
						__scalar_20 := int64(__let_scalar_17)
						_ = __scalar_20
						__read_21 := __arg3
						_ = __read_21
						__donor_slot_22 := __donor
						_ = __donor_slot_22
						var __cell_23 *Constructor_Main_Branch
						if (__donor_slot_22) != (nil) {
							__cell_23 = __donor_slot_22
							__donor_slot_22 = nil
						} else {

						}
						if (__cell_23) == (nil) {
							__cell_23 = new(Constructor_Main_Branch)
						} else {

						}
						__cell_23.Rc = 1
						__cell_23.V0 = __scalar_18
						__cell_23.V1 = __read_19
						__cell_23.V2 = __scalar_20
						__cell_23.V3 = __read_21
						return __cell_23
					}
				}
			} else {
				if ((__arg3) != (nil)) && ((__arg3.V0) == (2247809753)) {
					if (__arg3.V1) != (nil) {
						if (__arg3.V1.V0) == (2247809753) {
							__let_scalar_494 := int64(__arg2)
							_ = __let_scalar_494
							__let_scalar_495 := int64(__arg3.V1.V2)
							_ = __let_scalar_495
							__let_scalar_496 := int64(__arg3.V2)
							_ = __let_scalar_496
							__scalar_497 := uint32(2247809753)
							_ = __scalar_497
							__scalar_498 := uint32(1685833310)
							_ = __scalar_498
							__read_499 := __arg1
							_ = __read_499
							__scalar_500 := int64(__let_scalar_494)
							_ = __scalar_500
							__read_501 := __arg3.V1.V1
							_ = __read_501
							__scalar_502 := int64(__let_scalar_495)
							_ = __scalar_502
							__scalar_503 := uint32(1685833310)
							_ = __scalar_503
							__read_504 := __arg3.V1.V3
							_ = __read_504
							__scalar_505 := int64(__let_scalar_496)
							_ = __scalar_505
							__read_506 := __arg3.V3
							_ = __read_506
							__donor_slot_509 := __donor
							_ = __donor_slot_509
							__dead_507 := __arg3
							_ = __dead_507
							__dead_508 := __arg3.V1
							_ = __dead_508
							var __cell_510 *Constructor_Main_Branch
							if (__donor_slot_509) != (nil) {
								__cell_510 = __donor_slot_509
								__donor_slot_509 = nil
							} else {
								if (__dead_507) != (nil) {
									__cell_510 = __dead_507
									__dead_507 = nil
								} else {
									if (__dead_508) != (nil) {
										__cell_510 = __dead_508
										__dead_508 = nil
									} else {

									}
								}
							}
							if (__cell_510) == (nil) {
								__cell_510 = new(Constructor_Main_Branch)
							} else {

							}
							__cell_510.Rc = 1
							__cell_510.V0 = __scalar_498
							__cell_510.V1 = __read_499
							__cell_510.V2 = __scalar_500
							__cell_510.V3 = __read_501
							var __cell_511 *Constructor_Main_Branch
							if (__donor_slot_509) != (nil) {
								__cell_511 = __donor_slot_509
								__donor_slot_509 = nil
							} else {
								if (__dead_507) != (nil) {
									__cell_511 = __dead_507
									__dead_507 = nil
								} else {
									if (__dead_508) != (nil) {
										__cell_511 = __dead_508
										__dead_508 = nil
									} else {

									}
								}
							}
							if (__cell_511) == (nil) {
								__cell_511 = new(Constructor_Main_Branch)
							} else {

							}
							__cell_511.Rc = 1
							__cell_511.V0 = __scalar_503
							__cell_511.V1 = __read_504
							__cell_511.V2 = __scalar_505
							__cell_511.V3 = __read_506
							var __cell_512 *Constructor_Main_Branch
							if (__donor_slot_509) != (nil) {
								__cell_512 = __donor_slot_509
								__donor_slot_509 = nil
							} else {
								if (__dead_507) != (nil) {
									__cell_512 = __dead_507
									__dead_507 = nil
								} else {
									if (__dead_508) != (nil) {
										__cell_512 = __dead_508
										__dead_508 = nil
									} else {

									}
								}
							}
							if (__cell_512) == (nil) {
								__cell_512 = new(Constructor_Main_Branch)
							} else {

							}
							__cell_512.Rc = 1
							__cell_512.V0 = __scalar_497
							__cell_512.V1 = __cell_510
							__cell_512.V2 = __scalar_502
							__cell_512.V3 = __cell_511
							return __cell_512
						} else {
							if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (2247809753)) {
								__let_scalar_513 := int64(__arg2)
								_ = __let_scalar_513
								__let_scalar_514 := int64(__arg3.V2)
								_ = __let_scalar_514
								__let_scalar_515 := int64(__arg3.V3.V2)
								_ = __let_scalar_515
								__scalar_516 := uint32(2247809753)
								_ = __scalar_516
								__scalar_517 := uint32(1685833310)
								_ = __scalar_517
								__read_518 := __arg1
								_ = __read_518
								__scalar_519 := int64(__let_scalar_513)
								_ = __scalar_519
								__read_520 := __arg3.V1
								_ = __read_520
								__scalar_521 := int64(__let_scalar_514)
								_ = __scalar_521
								__scalar_522 := uint32(1685833310)
								_ = __scalar_522
								__read_523 := __arg3.V3.V1
								_ = __read_523
								__scalar_524 := int64(__let_scalar_515)
								_ = __scalar_524
								__read_525 := __arg3.V3.V3
								_ = __read_525
								__donor_slot_528 := __donor
								_ = __donor_slot_528
								__dead_526 := __arg3
								_ = __dead_526
								__dead_527 := __arg3.V3
								_ = __dead_527
								var __cell_529 *Constructor_Main_Branch
								if (__donor_slot_528) != (nil) {
									__cell_529 = __donor_slot_528
									__donor_slot_528 = nil
								} else {
									if (__dead_526) != (nil) {
										__cell_529 = __dead_526
										__dead_526 = nil
									} else {
										if (__dead_527) != (nil) {
											__cell_529 = __dead_527
											__dead_527 = nil
										} else {

										}
									}
								}
								if (__cell_529) == (nil) {
									__cell_529 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_529.Rc = 1
								__cell_529.V0 = __scalar_517
								__cell_529.V1 = __read_518
								__cell_529.V2 = __scalar_519
								__cell_529.V3 = __read_520
								var __cell_530 *Constructor_Main_Branch
								if (__donor_slot_528) != (nil) {
									__cell_530 = __donor_slot_528
									__donor_slot_528 = nil
								} else {
									if (__dead_526) != (nil) {
										__cell_530 = __dead_526
										__dead_526 = nil
									} else {
										if (__dead_527) != (nil) {
											__cell_530 = __dead_527
											__dead_527 = nil
										} else {

										}
									}
								}
								if (__cell_530) == (nil) {
									__cell_530 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_530.Rc = 1
								__cell_530.V0 = __scalar_522
								__cell_530.V1 = __read_523
								__cell_530.V2 = __scalar_524
								__cell_530.V3 = __read_525
								var __cell_531 *Constructor_Main_Branch
								if (__donor_slot_528) != (nil) {
									__cell_531 = __donor_slot_528
									__donor_slot_528 = nil
								} else {
									if (__dead_526) != (nil) {
										__cell_531 = __dead_526
										__dead_526 = nil
									} else {
										if (__dead_527) != (nil) {
											__cell_531 = __dead_527
											__dead_527 = nil
										} else {

										}
									}
								}
								if (__cell_531) == (nil) {
									__cell_531 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_531.Rc = 1
								__cell_531.V0 = __scalar_516
								__cell_531.V1 = __cell_529
								__cell_531.V2 = __scalar_521
								__cell_531.V3 = __cell_530
								return __cell_531
							} else {
								__let_scalar_486 := uint32(__arg0)
								_ = __let_scalar_486
								__let_scalar_487 := int64(__arg2)
								_ = __let_scalar_487
								__scalar_488 := uint32(__let_scalar_486)
								_ = __scalar_488
								__read_489 := __arg1
								_ = __read_489
								__scalar_490 := int64(__let_scalar_487)
								_ = __scalar_490
								__read_491 := __arg3
								_ = __read_491
								__donor_slot_492 := __donor
								_ = __donor_slot_492
								var __cell_493 *Constructor_Main_Branch
								if (__donor_slot_492) != (nil) {
									__cell_493 = __donor_slot_492
									__donor_slot_492 = nil
								} else {

								}
								if (__cell_493) == (nil) {
									__cell_493 = new(Constructor_Main_Branch)
								} else {

								}
								__cell_493.Rc = 1
								__cell_493.V0 = __scalar_488
								__cell_493.V1 = __read_489
								__cell_493.V2 = __scalar_490
								__cell_493.V3 = __read_491
								return __cell_493
							}
						}
					} else {
						if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (2247809753)) {
							__let_scalar_532 := int64(__arg2)
							_ = __let_scalar_532
							__let_scalar_533 := int64(__arg3.V2)
							_ = __let_scalar_533
							__let_scalar_534 := int64(__arg3.V3.V2)
							_ = __let_scalar_534
							__scalar_535 := uint32(2247809753)
							_ = __scalar_535
							__scalar_536 := uint32(1685833310)
							_ = __scalar_536
							__read_537 := __arg1
							_ = __read_537
							__scalar_538 := int64(__let_scalar_532)
							_ = __scalar_538
							__read_539 := __arg3.V1
							_ = __read_539
							__scalar_540 := int64(__let_scalar_533)
							_ = __scalar_540
							__scalar_541 := uint32(1685833310)
							_ = __scalar_541
							__read_542 := __arg3.V3.V1
							_ = __read_542
							__scalar_543 := int64(__let_scalar_534)
							_ = __scalar_543
							__read_544 := __arg3.V3.V3
							_ = __read_544
							__donor_slot_547 := __donor
							_ = __donor_slot_547
							__dead_545 := __arg3
							_ = __dead_545
							__dead_546 := __arg3.V3
							_ = __dead_546
							var __cell_548 *Constructor_Main_Branch
							if (__donor_slot_547) != (nil) {
								__cell_548 = __donor_slot_547
								__donor_slot_547 = nil
							} else {
								if (__dead_545) != (nil) {
									__cell_548 = __dead_545
									__dead_545 = nil
								} else {
									if (__dead_546) != (nil) {
										__cell_548 = __dead_546
										__dead_546 = nil
									} else {

									}
								}
							}
							if (__cell_548) == (nil) {
								__cell_548 = new(Constructor_Main_Branch)
							} else {

							}
							__cell_548.Rc = 1
							__cell_548.V0 = __scalar_536
							__cell_548.V1 = __read_537
							__cell_548.V2 = __scalar_538
							__cell_548.V3 = __read_539
							var __cell_549 *Constructor_Main_Branch
							if (__donor_slot_547) != (nil) {
								__cell_549 = __donor_slot_547
								__donor_slot_547 = nil
							} else {
								if (__dead_545) != (nil) {
									__cell_549 = __dead_545
									__dead_545 = nil
								} else {
									if (__dead_546) != (nil) {
										__cell_549 = __dead_546
										__dead_546 = nil
									} else {

									}
								}
							}
							if (__cell_549) == (nil) {
								__cell_549 = new(Constructor_Main_Branch)
							} else {

							}
							__cell_549.Rc = 1
							__cell_549.V0 = __scalar_541
							__cell_549.V1 = __read_542
							__cell_549.V2 = __scalar_543
							__cell_549.V3 = __read_544
							var __cell_550 *Constructor_Main_Branch
							if (__donor_slot_547) != (nil) {
								__cell_550 = __donor_slot_547
								__donor_slot_547 = nil
							} else {
								if (__dead_545) != (nil) {
									__cell_550 = __dead_545
									__dead_545 = nil
								} else {
									if (__dead_546) != (nil) {
										__cell_550 = __dead_546
										__dead_546 = nil
									} else {

									}
								}
							}
							if (__cell_550) == (nil) {
								__cell_550 = new(Constructor_Main_Branch)
							} else {

							}
							__cell_550.Rc = 1
							__cell_550.V0 = __scalar_535
							__cell_550.V1 = __cell_548
							__cell_550.V2 = __scalar_540
							__cell_550.V3 = __cell_549
							return __cell_550
						} else {
							__let_scalar_478 := uint32(__arg0)
							_ = __let_scalar_478
							__let_scalar_479 := int64(__arg2)
							_ = __let_scalar_479
							__scalar_480 := uint32(__let_scalar_478)
							_ = __scalar_480
							__read_481 := __arg1
							_ = __read_481
							__scalar_482 := int64(__let_scalar_479)
							_ = __scalar_482
							__read_483 := __arg3
							_ = __read_483
							__donor_slot_484 := __donor
							_ = __donor_slot_484
							var __cell_485 *Constructor_Main_Branch
							if (__donor_slot_484) != (nil) {
								__cell_485 = __donor_slot_484
								__donor_slot_484 = nil
							} else {

							}
							if (__cell_485) == (nil) {
								__cell_485 = new(Constructor_Main_Branch)
							} else {

							}
							__cell_485.Rc = 1
							__cell_485.V0 = __scalar_480
							__cell_485.V1 = __read_481
							__cell_485.V2 = __scalar_482
							__cell_485.V3 = __read_483
							return __cell_485
						}
					}
				} else {
					__let_scalar_8 := uint32(__arg0)
					_ = __let_scalar_8
					__let_scalar_9 := int64(__arg2)
					_ = __let_scalar_9
					__scalar_10 := uint32(__let_scalar_8)
					_ = __scalar_10
					__read_11 := __arg1
					_ = __read_11
					__scalar_12 := int64(__let_scalar_9)
					_ = __scalar_12
					__read_13 := __arg3
					_ = __read_13
					__donor_slot_14 := __donor
					_ = __donor_slot_14
					var __cell_15 *Constructor_Main_Branch
					if (__donor_slot_14) != (nil) {
						__cell_15 = __donor_slot_14
						__donor_slot_14 = nil
					} else {

					}
					if (__cell_15) == (nil) {
						__cell_15 = new(Constructor_Main_Branch)
					} else {

					}
					__cell_15.Rc = 1
					__cell_15.V0 = __scalar_10
					__cell_15.V1 = __read_11
					__cell_15.V2 = __scalar_12
					__cell_15.V3 = __read_13
					return __cell_15
				}
			}
		} else {
			__let_scalar_0 := uint32(__arg0)
			_ = __let_scalar_0
			__let_scalar_1 := int64(__arg2)
			_ = __let_scalar_1
			__scalar_2 := uint32(__let_scalar_0)
			_ = __scalar_2
			__read_3 := __arg1
			_ = __read_3
			__scalar_4 := int64(__let_scalar_1)
			_ = __scalar_4
			__read_5 := __arg3
			_ = __read_5
			__donor_slot_6 := __donor
			_ = __donor_slot_6
			var __cell_7 *Constructor_Main_Branch
			if (__donor_slot_6) != (nil) {
				__cell_7 = __donor_slot_6
				__donor_slot_6 = nil
			} else {

			}
			if (__cell_7) == (nil) {
				__cell_7 = new(Constructor_Main_Branch)
			} else {

			}
			__cell_7.Rc = 1
			__cell_7.V0 = __scalar_2
			__cell_7.V1 = __read_3
			__cell_7.V2 = __scalar_4
			__cell_7.V3 = __read_5
			return __cell_7
		}
	}
}

func Call_Main___gopurs_owned_rebalance_0(__arg0 uint32, __arg1 *Constructor_Main_Branch, __arg2 int64, __arg3 *Constructor_Main_Branch) *Constructor_Main_Branch {
	return Call_Main___gopurs_owned_rebalance_0_consume(__arg0, __arg1, __arg2, __arg3, nil)
}

type Constructor_Main_Crimson struct {
	Rc uint32
}

type Constructor_Main_Onyx struct {
	Rc uint32
}

type Constructor_Main_Tip struct {
	Rc uint32
}

type Constructor_Main_Branch struct {
	Rc uint32
	V0 uint32
	V1 *Constructor_Main_Branch
	V2 int64
	V3 *Constructor_Main_Branch
}

func Call_Main_render(v_0_loop *Constructor_Main_Branch) string {
render:
	for {
		if false {
			continue render
		}
		var v_0 *Constructor_Main_Branch = v_0_loop
		_ = v_0
		var __t3 string
		{
			if v_0 == nil {
				__t3 = "E"
				goto end_branch_3
			} else {

			}
		}
		{
			if v_0 != nil {
				var __t2 string
				{
					var __t_tag_0 uint32 = (v_0).V0
					_ = __t_tag_0
					if uint32(__t_tag_0) == 2247809753 {
						__t2 = "R"
						goto end_branch_2
					} else {

					}
				}
				{
					var __t_tag_1 uint32 = (v_0).V0
					_ = __t_tag_1
					if uint32(__t_tag_1) == 1685833310 {
						__t2 = "B"
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
				}
			end_branch_2:
				__t3 = ((((((__t2) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((v_0).V2)).StrVal())) + ("(")) + (Call_Main_render((v_0).V1))) + (")(")) + (Call_Main_render((v_0).V3))) + (")")
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = func() string { panic("Failed pattern match") }()
		}
	end_branch_3:
		return __t3
	}
}

func Call_Main_rebalance(v_0_loop uint32, v1_1_loop *Constructor_Main_Branch, v2_2_loop int64, v3_3_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var v1_1 *Constructor_Main_Branch = v1_1_loop
	_ = v1_1
	var v2_2 int64 = v2_2_loop
	_ = v2_2
	var v3_3 *Constructor_Main_Branch = v3_3_loop
	_ = v3_3
	var __t308 *Constructor_Main_Branch
	{
		if v_0 == 1685833310 {
			var __t307 *Constructor_Main_Branch
			{
				if v1_1 != nil {
					var __t265 *Constructor_Main_Branch
					{
						var __t_tag_12 uint32 = (v1_1).V0
						_ = __t_tag_12
						if uint32(__t_tag_12) == 2247809753 {
							var __t223 *Constructor_Main_Branch
							{
								var __t_tag_17 *Constructor_Main_Branch = (v1_1).V1
								_ = __t_tag_17
								if __t_tag_17 != nil {
									var __t126 *Constructor_Main_Branch
									{
										var __t_tag_22 uint32 = ((v1_1).V1).V0
										_ = __t_tag_22
										if uint32(__t_tag_22) == 2247809753 {
											// TAST (Let): __local_var_4_23 shape=Other bindingType=Any
											__local_var_4_23 := ((v1_1).V1).V1
											_ = __local_var_4_23
											// TAST (Let): __local_var_5_24 shape=Other bindingType=Any
											__local_var_5_24 := ((v1_1).V1).V3
											_ = __local_var_5_24
											// TAST (Let): __local_var_6_25 shape=Other bindingType=Any
											__local_var_6_25 := (v1_1).V3
											_ = __local_var_6_25
											// TAST (Let): __local_var_7_26 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
											__local_var_7_26 := v3_3
											_ = __local_var_7_26
											// TAST (Let): __local_var_8_27 shape=Other bindingType=Any
											__local_var_8_27 := ((v1_1).V1).V2
											_ = __local_var_8_27
											// TAST (Let): __local_var_9_28 shape=Other bindingType=Any
											__local_var_9_28 := (v1_1).V2
											_ = __local_var_9_28
											// TAST (Let): __local_var_10_29 shape=Other bindingType=Int
											__local_var_10_29 := v2_2
											_ = __local_var_10_29
											__t126 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_23, __local_var_8_27, __local_var_5_24}), __local_var_9_28, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_25, __local_var_10_29, __local_var_7_26})})
											goto end_branch_126
										} else {

										}
									}
									{
										var __t_tag_30 *Constructor_Main_Branch = (v1_1).V3
										_ = __t_tag_30
										if __t_tag_30 != nil {
											var __t84 *Constructor_Main_Branch
											{
												var __t_tag_35 uint32 = ((v1_1).V3).V0
												_ = __t_tag_35
												if uint32(__t_tag_35) == 2247809753 {
													// TAST (Let): __local_var_4_36 shape=Other bindingType=Any
													__local_var_4_36 := (v1_1).V1
													_ = __local_var_4_36
													// TAST (Let): __local_var_5_37 shape=Other bindingType=Any
													__local_var_5_37 := ((v1_1).V3).V1
													_ = __local_var_5_37
													// TAST (Let): __local_var_6_38 shape=Other bindingType=Any
													__local_var_6_38 := ((v1_1).V3).V3
													_ = __local_var_6_38
													// TAST (Let): __local_var_7_39 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
													__local_var_7_39 := v3_3
													_ = __local_var_7_39
													// TAST (Let): __local_var_8_40 shape=Other bindingType=Any
													__local_var_8_40 := (v1_1).V2
													_ = __local_var_8_40
													// TAST (Let): __local_var_9_41 shape=Other bindingType=Any
													__local_var_9_41 := ((v1_1).V3).V2
													_ = __local_var_9_41
													// TAST (Let): __local_var_10_42 shape=Other bindingType=Int
													__local_var_10_42 := v2_2
													_ = __local_var_10_42
													__t84 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_36, __local_var_8_40, __local_var_5_37}), __local_var_9_41, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_38, __local_var_10_42, __local_var_7_39})})
													goto end_branch_84
												} else {

												}
											}
											{
												var __t_and_44 bool = false
												if v3_3 != nil {

													var __t_tag_43 uint32 = (v3_3).V0
													_ = __t_tag_43
													__t_and_44 = (uint32(__t_tag_43) == 2247809753)
												}
												if __t_and_44 {
													var __t83 *Constructor_Main_Branch
													{
														var __t_tag_49 *Constructor_Main_Branch = (v3_3).V1
														_ = __t_tag_49
														if __t_tag_49 != nil {
															var __t72 *Constructor_Main_Branch
															{
																var __t_tag_54 uint32 = ((v3_3).V1).V0
																_ = __t_tag_54
																if uint32(__t_tag_54) == 2247809753 {
																	// TAST (Let): __local_var_4_55 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
																	__local_var_4_55 := v1_1
																	_ = __local_var_4_55
																	// TAST (Let): __local_var_5_56 shape=Other bindingType=Any
																	__local_var_5_56 := ((v3_3).V1).V1
																	_ = __local_var_5_56
																	// TAST (Let): __local_var_6_57 shape=Other bindingType=Any
																	__local_var_6_57 := ((v3_3).V1).V3
																	_ = __local_var_6_57
																	// TAST (Let): __local_var_7_58 shape=Other bindingType=Any
																	__local_var_7_58 := (v3_3).V3
																	_ = __local_var_7_58
																	// TAST (Let): __local_var_8_59 shape=Other bindingType=Int
																	__local_var_8_59 := v2_2
																	_ = __local_var_8_59
																	// TAST (Let): __local_var_9_60 shape=Other bindingType=Any
																	__local_var_9_60 := ((v3_3).V1).V2
																	_ = __local_var_9_60
																	// TAST (Let): __local_var_10_61 shape=Other bindingType=Any
																	__local_var_10_61 := (v3_3).V2
																	_ = __local_var_10_61
																	__t72 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_55, __local_var_8_59, __local_var_5_56}), __local_var_9_60, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_57, __local_var_10_61, __local_var_7_58})})
																	goto end_branch_72
																} else {

																}
															}
															{
																var __t_tag_62 *Constructor_Main_Branch = (v3_3).V3
																_ = __t_tag_62
																var __t_and_64 bool = false
																if __t_tag_62 != nil {

																	var __t_tag_63 uint32 = ((v3_3).V3).V0
																	_ = __t_tag_63
																	__t_and_64 = (uint32(__t_tag_63) == 2247809753)
																}
																if __t_and_64 {
																	// TAST (Let): __local_var_4_65 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
																	__local_var_4_65 := v1_1
																	_ = __local_var_4_65
																	// TAST (Let): __local_var_5_66 shape=Other bindingType=Any
																	__local_var_5_66 := (v3_3).V1
																	_ = __local_var_5_66
																	// TAST (Let): __local_var_6_67 shape=Other bindingType=Any
																	__local_var_6_67 := ((v3_3).V3).V1
																	_ = __local_var_6_67
																	// TAST (Let): __local_var_7_68 shape=Other bindingType=Any
																	__local_var_7_68 := ((v3_3).V3).V3
																	_ = __local_var_7_68
																	// TAST (Let): __local_var_8_69 shape=Other bindingType=Int
																	__local_var_8_69 := v2_2
																	_ = __local_var_8_69
																	// TAST (Let): __local_var_9_70 shape=Other bindingType=Any
																	__local_var_9_70 := (v3_3).V2
																	_ = __local_var_9_70
																	// TAST (Let): __local_var_10_71 shape=Other bindingType=Any
																	__local_var_10_71 := ((v3_3).V3).V2
																	_ = __local_var_10_71
																	__t72 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_65, __local_var_8_69, __local_var_5_66}), __local_var_9_70, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_67, __local_var_10_71, __local_var_7_68})})
																	goto end_branch_72
																} else {

																}
															}
															{
																// TAST (Let): __local_var_4_50 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
																__local_var_4_50 := v1_1
																_ = __local_var_4_50
																// TAST (Let): __local_var_5_51 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
																__local_var_5_51 := v3_3
																_ = __local_var_5_51
																// TAST (Let): __local_var_6_52 shape=Other bindingType=(ADT ["Main","Paint"] [])
																__local_var_6_52 := v_0
																_ = __local_var_6_52
																// TAST (Let): __local_var_7_53 shape=Other bindingType=Int
																__local_var_7_53 := v2_2
																_ = __local_var_7_53
																__t72 = (&Constructor_Main_Branch{1, __local_var_6_52, __local_var_4_50, __local_var_7_53, __local_var_5_51})
															}
														end_branch_72:
															__t83 = __t72
															goto end_branch_83
														} else {

														}
													}
													{
														var __t_tag_73 *Constructor_Main_Branch = (v3_3).V3
														_ = __t_tag_73
														var __t_and_75 bool = false
														if __t_tag_73 != nil {

															var __t_tag_74 uint32 = ((v3_3).V3).V0
															_ = __t_tag_74
															__t_and_75 = (uint32(__t_tag_74) == 2247809753)
														}
														if __t_and_75 {
															// TAST (Let): __local_var_4_76 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
															__local_var_4_76 := v1_1
															_ = __local_var_4_76
															// TAST (Let): __local_var_5_77 shape=Other bindingType=Any
															__local_var_5_77 := (v3_3).V1
															_ = __local_var_5_77
															// TAST (Let): __local_var_6_78 shape=Other bindingType=Any
															__local_var_6_78 := ((v3_3).V3).V1
															_ = __local_var_6_78
															// TAST (Let): __local_var_7_79 shape=Other bindingType=Any
															__local_var_7_79 := ((v3_3).V3).V3
															_ = __local_var_7_79
															// TAST (Let): __local_var_8_80 shape=Other bindingType=Int
															__local_var_8_80 := v2_2
															_ = __local_var_8_80
															// TAST (Let): __local_var_9_81 shape=Other bindingType=Any
															__local_var_9_81 := (v3_3).V2
															_ = __local_var_9_81
															// TAST (Let): __local_var_10_82 shape=Other bindingType=Any
															__local_var_10_82 := ((v3_3).V3).V2
															_ = __local_var_10_82
															__t83 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_76, __local_var_8_80, __local_var_5_77}), __local_var_9_81, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_78, __local_var_10_82, __local_var_7_79})})
															goto end_branch_83
														} else {

														}
													}
													{
														// TAST (Let): __local_var_4_45 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
														__local_var_4_45 := v1_1
														_ = __local_var_4_45
														// TAST (Let): __local_var_5_46 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
														__local_var_5_46 := v3_3
														_ = __local_var_5_46
														// TAST (Let): __local_var_6_47 shape=Other bindingType=(ADT ["Main","Paint"] [])
														__local_var_6_47 := v_0
														_ = __local_var_6_47
														// TAST (Let): __local_var_7_48 shape=Other bindingType=Int
														__local_var_7_48 := v2_2
														_ = __local_var_7_48
														__t83 = (&Constructor_Main_Branch{1, __local_var_6_47, __local_var_4_45, __local_var_7_48, __local_var_5_46})
													}
												end_branch_83:
													__t84 = __t83
													goto end_branch_84
												} else {

												}
											}
											{
												// TAST (Let): __local_var_4_31 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_4_31 := v1_1
												_ = __local_var_4_31
												// TAST (Let): __local_var_5_32 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_5_32 := v3_3
												_ = __local_var_5_32
												// TAST (Let): __local_var_6_33 shape=Other bindingType=(ADT ["Main","Paint"] [])
												__local_var_6_33 := v_0
												_ = __local_var_6_33
												// TAST (Let): __local_var_7_34 shape=Other bindingType=Int
												__local_var_7_34 := v2_2
												_ = __local_var_7_34
												__t84 = (&Constructor_Main_Branch{1, __local_var_6_33, __local_var_4_31, __local_var_7_34, __local_var_5_32})
											}
										end_branch_84:
											__t126 = __t84
											goto end_branch_126
										} else {

										}
									}
									{
										var __t_and_86 bool = false
										if v3_3 != nil {

											var __t_tag_85 uint32 = (v3_3).V0
											_ = __t_tag_85
											__t_and_86 = (uint32(__t_tag_85) == 2247809753)
										}
										if __t_and_86 {
											var __t125 *Constructor_Main_Branch
											{
												var __t_tag_91 *Constructor_Main_Branch = (v3_3).V1
												_ = __t_tag_91
												if __t_tag_91 != nil {
													var __t114 *Constructor_Main_Branch
													{
														var __t_tag_96 uint32 = ((v3_3).V1).V0
														_ = __t_tag_96
														if uint32(__t_tag_96) == 2247809753 {
															// TAST (Let): __local_var_4_97 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
															__local_var_4_97 := v1_1
															_ = __local_var_4_97
															// TAST (Let): __local_var_5_98 shape=Other bindingType=Any
															__local_var_5_98 := ((v3_3).V1).V1
															_ = __local_var_5_98
															// TAST (Let): __local_var_6_99 shape=Other bindingType=Any
															__local_var_6_99 := ((v3_3).V1).V3
															_ = __local_var_6_99
															// TAST (Let): __local_var_7_100 shape=Other bindingType=Any
															__local_var_7_100 := (v3_3).V3
															_ = __local_var_7_100
															// TAST (Let): __local_var_8_101 shape=Other bindingType=Int
															__local_var_8_101 := v2_2
															_ = __local_var_8_101
															// TAST (Let): __local_var_9_102 shape=Other bindingType=Any
															__local_var_9_102 := ((v3_3).V1).V2
															_ = __local_var_9_102
															// TAST (Let): __local_var_10_103 shape=Other bindingType=Any
															__local_var_10_103 := (v3_3).V2
															_ = __local_var_10_103
															__t114 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_97, __local_var_8_101, __local_var_5_98}), __local_var_9_102, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_99, __local_var_10_103, __local_var_7_100})})
															goto end_branch_114
														} else {

														}
													}
													{
														var __t_tag_104 *Constructor_Main_Branch = (v3_3).V3
														_ = __t_tag_104
														var __t_and_106 bool = false
														if __t_tag_104 != nil {

															var __t_tag_105 uint32 = ((v3_3).V3).V0
															_ = __t_tag_105
															__t_and_106 = (uint32(__t_tag_105) == 2247809753)
														}
														if __t_and_106 {
															// TAST (Let): __local_var_4_107 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
															__local_var_4_107 := v1_1
															_ = __local_var_4_107
															// TAST (Let): __local_var_5_108 shape=Other bindingType=Any
															__local_var_5_108 := (v3_3).V1
															_ = __local_var_5_108
															// TAST (Let): __local_var_6_109 shape=Other bindingType=Any
															__local_var_6_109 := ((v3_3).V3).V1
															_ = __local_var_6_109
															// TAST (Let): __local_var_7_110 shape=Other bindingType=Any
															__local_var_7_110 := ((v3_3).V3).V3
															_ = __local_var_7_110
															// TAST (Let): __local_var_8_111 shape=Other bindingType=Int
															__local_var_8_111 := v2_2
															_ = __local_var_8_111
															// TAST (Let): __local_var_9_112 shape=Other bindingType=Any
															__local_var_9_112 := (v3_3).V2
															_ = __local_var_9_112
															// TAST (Let): __local_var_10_113 shape=Other bindingType=Any
															__local_var_10_113 := ((v3_3).V3).V2
															_ = __local_var_10_113
															__t114 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_107, __local_var_8_111, __local_var_5_108}), __local_var_9_112, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_109, __local_var_10_113, __local_var_7_110})})
															goto end_branch_114
														} else {

														}
													}
													{
														// TAST (Let): __local_var_4_92 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
														__local_var_4_92 := v1_1
														_ = __local_var_4_92
														// TAST (Let): __local_var_5_93 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
														__local_var_5_93 := v3_3
														_ = __local_var_5_93
														// TAST (Let): __local_var_6_94 shape=Other bindingType=(ADT ["Main","Paint"] [])
														__local_var_6_94 := v_0
														_ = __local_var_6_94
														// TAST (Let): __local_var_7_95 shape=Other bindingType=Int
														__local_var_7_95 := v2_2
														_ = __local_var_7_95
														__t114 = (&Constructor_Main_Branch{1, __local_var_6_94, __local_var_4_92, __local_var_7_95, __local_var_5_93})
													}
												end_branch_114:
													__t125 = __t114
													goto end_branch_125
												} else {

												}
											}
											{
												var __t_tag_115 *Constructor_Main_Branch = (v3_3).V3
												_ = __t_tag_115
												var __t_and_117 bool = false
												if __t_tag_115 != nil {

													var __t_tag_116 uint32 = ((v3_3).V3).V0
													_ = __t_tag_116
													__t_and_117 = (uint32(__t_tag_116) == 2247809753)
												}
												if __t_and_117 {
													// TAST (Let): __local_var_4_118 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
													__local_var_4_118 := v1_1
													_ = __local_var_4_118
													// TAST (Let): __local_var_5_119 shape=Other bindingType=Any
													__local_var_5_119 := (v3_3).V1
													_ = __local_var_5_119
													// TAST (Let): __local_var_6_120 shape=Other bindingType=Any
													__local_var_6_120 := ((v3_3).V3).V1
													_ = __local_var_6_120
													// TAST (Let): __local_var_7_121 shape=Other bindingType=Any
													__local_var_7_121 := ((v3_3).V3).V3
													_ = __local_var_7_121
													// TAST (Let): __local_var_8_122 shape=Other bindingType=Int
													__local_var_8_122 := v2_2
													_ = __local_var_8_122
													// TAST (Let): __local_var_9_123 shape=Other bindingType=Any
													__local_var_9_123 := (v3_3).V2
													_ = __local_var_9_123
													// TAST (Let): __local_var_10_124 shape=Other bindingType=Any
													__local_var_10_124 := ((v3_3).V3).V2
													_ = __local_var_10_124
													__t125 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_118, __local_var_8_122, __local_var_5_119}), __local_var_9_123, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_120, __local_var_10_124, __local_var_7_121})})
													goto end_branch_125
												} else {

												}
											}
											{
												// TAST (Let): __local_var_4_87 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_4_87 := v1_1
												_ = __local_var_4_87
												// TAST (Let): __local_var_5_88 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_5_88 := v3_3
												_ = __local_var_5_88
												// TAST (Let): __local_var_6_89 shape=Other bindingType=(ADT ["Main","Paint"] [])
												__local_var_6_89 := v_0
												_ = __local_var_6_89
												// TAST (Let): __local_var_7_90 shape=Other bindingType=Int
												__local_var_7_90 := v2_2
												_ = __local_var_7_90
												__t125 = (&Constructor_Main_Branch{1, __local_var_6_89, __local_var_4_87, __local_var_7_90, __local_var_5_88})
											}
										end_branch_125:
											__t126 = __t125
											goto end_branch_126
										} else {

										}
									}
									{
										// TAST (Let): __local_var_4_18 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_4_18 := v1_1
										_ = __local_var_4_18
										// TAST (Let): __local_var_5_19 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_5_19 := v3_3
										_ = __local_var_5_19
										// TAST (Let): __local_var_6_20 shape=Other bindingType=(ADT ["Main","Paint"] [])
										__local_var_6_20 := v_0
										_ = __local_var_6_20
										// TAST (Let): __local_var_7_21 shape=Other bindingType=Int
										__local_var_7_21 := v2_2
										_ = __local_var_7_21
										__t126 = (&Constructor_Main_Branch{1, __local_var_6_20, __local_var_4_18, __local_var_7_21, __local_var_5_19})
									}
								end_branch_126:
									__t223 = __t126
									goto end_branch_223
								} else {

								}
							}
							{
								var __t_tag_127 *Constructor_Main_Branch = (v1_1).V3
								_ = __t_tag_127
								if __t_tag_127 != nil {
									var __t181 *Constructor_Main_Branch
									{
										var __t_tag_132 uint32 = ((v1_1).V3).V0
										_ = __t_tag_132
										if uint32(__t_tag_132) == 2247809753 {
											// TAST (Let): __local_var_4_133 shape=Other bindingType=Any
											__local_var_4_133 := (v1_1).V1
											_ = __local_var_4_133
											// TAST (Let): __local_var_5_134 shape=Other bindingType=Any
											__local_var_5_134 := ((v1_1).V3).V1
											_ = __local_var_5_134
											// TAST (Let): __local_var_6_135 shape=Other bindingType=Any
											__local_var_6_135 := ((v1_1).V3).V3
											_ = __local_var_6_135
											// TAST (Let): __local_var_7_136 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
											__local_var_7_136 := v3_3
											_ = __local_var_7_136
											// TAST (Let): __local_var_8_137 shape=Other bindingType=Any
											__local_var_8_137 := (v1_1).V2
											_ = __local_var_8_137
											// TAST (Let): __local_var_9_138 shape=Other bindingType=Any
											__local_var_9_138 := ((v1_1).V3).V2
											_ = __local_var_9_138
											// TAST (Let): __local_var_10_139 shape=Other bindingType=Int
											__local_var_10_139 := v2_2
											_ = __local_var_10_139
											__t181 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_133, __local_var_8_137, __local_var_5_134}), __local_var_9_138, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_135, __local_var_10_139, __local_var_7_136})})
											goto end_branch_181
										} else {

										}
									}
									{
										var __t_and_141 bool = false
										if v3_3 != nil {

											var __t_tag_140 uint32 = (v3_3).V0
											_ = __t_tag_140
											__t_and_141 = (uint32(__t_tag_140) == 2247809753)
										}
										if __t_and_141 {
											var __t180 *Constructor_Main_Branch
											{
												var __t_tag_146 *Constructor_Main_Branch = (v3_3).V1
												_ = __t_tag_146
												if __t_tag_146 != nil {
													var __t169 *Constructor_Main_Branch
													{
														var __t_tag_151 uint32 = ((v3_3).V1).V0
														_ = __t_tag_151
														if uint32(__t_tag_151) == 2247809753 {
															// TAST (Let): __local_var_4_152 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
															__local_var_4_152 := v1_1
															_ = __local_var_4_152
															// TAST (Let): __local_var_5_153 shape=Other bindingType=Any
															__local_var_5_153 := ((v3_3).V1).V1
															_ = __local_var_5_153
															// TAST (Let): __local_var_6_154 shape=Other bindingType=Any
															__local_var_6_154 := ((v3_3).V1).V3
															_ = __local_var_6_154
															// TAST (Let): __local_var_7_155 shape=Other bindingType=Any
															__local_var_7_155 := (v3_3).V3
															_ = __local_var_7_155
															// TAST (Let): __local_var_8_156 shape=Other bindingType=Int
															__local_var_8_156 := v2_2
															_ = __local_var_8_156
															// TAST (Let): __local_var_9_157 shape=Other bindingType=Any
															__local_var_9_157 := ((v3_3).V1).V2
															_ = __local_var_9_157
															// TAST (Let): __local_var_10_158 shape=Other bindingType=Any
															__local_var_10_158 := (v3_3).V2
															_ = __local_var_10_158
															__t169 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_152, __local_var_8_156, __local_var_5_153}), __local_var_9_157, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_154, __local_var_10_158, __local_var_7_155})})
															goto end_branch_169
														} else {

														}
													}
													{
														var __t_tag_159 *Constructor_Main_Branch = (v3_3).V3
														_ = __t_tag_159
														var __t_and_161 bool = false
														if __t_tag_159 != nil {

															var __t_tag_160 uint32 = ((v3_3).V3).V0
															_ = __t_tag_160
															__t_and_161 = (uint32(__t_tag_160) == 2247809753)
														}
														if __t_and_161 {
															// TAST (Let): __local_var_4_162 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
															__local_var_4_162 := v1_1
															_ = __local_var_4_162
															// TAST (Let): __local_var_5_163 shape=Other bindingType=Any
															__local_var_5_163 := (v3_3).V1
															_ = __local_var_5_163
															// TAST (Let): __local_var_6_164 shape=Other bindingType=Any
															__local_var_6_164 := ((v3_3).V3).V1
															_ = __local_var_6_164
															// TAST (Let): __local_var_7_165 shape=Other bindingType=Any
															__local_var_7_165 := ((v3_3).V3).V3
															_ = __local_var_7_165
															// TAST (Let): __local_var_8_166 shape=Other bindingType=Int
															__local_var_8_166 := v2_2
															_ = __local_var_8_166
															// TAST (Let): __local_var_9_167 shape=Other bindingType=Any
															__local_var_9_167 := (v3_3).V2
															_ = __local_var_9_167
															// TAST (Let): __local_var_10_168 shape=Other bindingType=Any
															__local_var_10_168 := ((v3_3).V3).V2
															_ = __local_var_10_168
															__t169 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_162, __local_var_8_166, __local_var_5_163}), __local_var_9_167, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_164, __local_var_10_168, __local_var_7_165})})
															goto end_branch_169
														} else {

														}
													}
													{
														// TAST (Let): __local_var_4_147 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
														__local_var_4_147 := v1_1
														_ = __local_var_4_147
														// TAST (Let): __local_var_5_148 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
														__local_var_5_148 := v3_3
														_ = __local_var_5_148
														// TAST (Let): __local_var_6_149 shape=Other bindingType=(ADT ["Main","Paint"] [])
														__local_var_6_149 := v_0
														_ = __local_var_6_149
														// TAST (Let): __local_var_7_150 shape=Other bindingType=Int
														__local_var_7_150 := v2_2
														_ = __local_var_7_150
														__t169 = (&Constructor_Main_Branch{1, __local_var_6_149, __local_var_4_147, __local_var_7_150, __local_var_5_148})
													}
												end_branch_169:
													__t180 = __t169
													goto end_branch_180
												} else {

												}
											}
											{
												var __t_tag_170 *Constructor_Main_Branch = (v3_3).V3
												_ = __t_tag_170
												var __t_and_172 bool = false
												if __t_tag_170 != nil {

													var __t_tag_171 uint32 = ((v3_3).V3).V0
													_ = __t_tag_171
													__t_and_172 = (uint32(__t_tag_171) == 2247809753)
												}
												if __t_and_172 {
													// TAST (Let): __local_var_4_173 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
													__local_var_4_173 := v1_1
													_ = __local_var_4_173
													// TAST (Let): __local_var_5_174 shape=Other bindingType=Any
													__local_var_5_174 := (v3_3).V1
													_ = __local_var_5_174
													// TAST (Let): __local_var_6_175 shape=Other bindingType=Any
													__local_var_6_175 := ((v3_3).V3).V1
													_ = __local_var_6_175
													// TAST (Let): __local_var_7_176 shape=Other bindingType=Any
													__local_var_7_176 := ((v3_3).V3).V3
													_ = __local_var_7_176
													// TAST (Let): __local_var_8_177 shape=Other bindingType=Int
													__local_var_8_177 := v2_2
													_ = __local_var_8_177
													// TAST (Let): __local_var_9_178 shape=Other bindingType=Any
													__local_var_9_178 := (v3_3).V2
													_ = __local_var_9_178
													// TAST (Let): __local_var_10_179 shape=Other bindingType=Any
													__local_var_10_179 := ((v3_3).V3).V2
													_ = __local_var_10_179
													__t180 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_173, __local_var_8_177, __local_var_5_174}), __local_var_9_178, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_175, __local_var_10_179, __local_var_7_176})})
													goto end_branch_180
												} else {

												}
											}
											{
												// TAST (Let): __local_var_4_142 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_4_142 := v1_1
												_ = __local_var_4_142
												// TAST (Let): __local_var_5_143 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_5_143 := v3_3
												_ = __local_var_5_143
												// TAST (Let): __local_var_6_144 shape=Other bindingType=(ADT ["Main","Paint"] [])
												__local_var_6_144 := v_0
												_ = __local_var_6_144
												// TAST (Let): __local_var_7_145 shape=Other bindingType=Int
												__local_var_7_145 := v2_2
												_ = __local_var_7_145
												__t180 = (&Constructor_Main_Branch{1, __local_var_6_144, __local_var_4_142, __local_var_7_145, __local_var_5_143})
											}
										end_branch_180:
											__t181 = __t180
											goto end_branch_181
										} else {

										}
									}
									{
										// TAST (Let): __local_var_4_128 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_4_128 := v1_1
										_ = __local_var_4_128
										// TAST (Let): __local_var_5_129 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_5_129 := v3_3
										_ = __local_var_5_129
										// TAST (Let): __local_var_6_130 shape=Other bindingType=(ADT ["Main","Paint"] [])
										__local_var_6_130 := v_0
										_ = __local_var_6_130
										// TAST (Let): __local_var_7_131 shape=Other bindingType=Int
										__local_var_7_131 := v2_2
										_ = __local_var_7_131
										__t181 = (&Constructor_Main_Branch{1, __local_var_6_130, __local_var_4_128, __local_var_7_131, __local_var_5_129})
									}
								end_branch_181:
									__t223 = __t181
									goto end_branch_223
								} else {

								}
							}
							{
								var __t_and_183 bool = false
								if v3_3 != nil {

									var __t_tag_182 uint32 = (v3_3).V0
									_ = __t_tag_182
									__t_and_183 = (uint32(__t_tag_182) == 2247809753)
								}
								if __t_and_183 {
									var __t222 *Constructor_Main_Branch
									{
										var __t_tag_188 *Constructor_Main_Branch = (v3_3).V1
										_ = __t_tag_188
										if __t_tag_188 != nil {
											var __t211 *Constructor_Main_Branch
											{
												var __t_tag_193 uint32 = ((v3_3).V1).V0
												_ = __t_tag_193
												if uint32(__t_tag_193) == 2247809753 {
													// TAST (Let): __local_var_4_194 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
													__local_var_4_194 := v1_1
													_ = __local_var_4_194
													// TAST (Let): __local_var_5_195 shape=Other bindingType=Any
													__local_var_5_195 := ((v3_3).V1).V1
													_ = __local_var_5_195
													// TAST (Let): __local_var_6_196 shape=Other bindingType=Any
													__local_var_6_196 := ((v3_3).V1).V3
													_ = __local_var_6_196
													// TAST (Let): __local_var_7_197 shape=Other bindingType=Any
													__local_var_7_197 := (v3_3).V3
													_ = __local_var_7_197
													// TAST (Let): __local_var_8_198 shape=Other bindingType=Int
													__local_var_8_198 := v2_2
													_ = __local_var_8_198
													// TAST (Let): __local_var_9_199 shape=Other bindingType=Any
													__local_var_9_199 := ((v3_3).V1).V2
													_ = __local_var_9_199
													// TAST (Let): __local_var_10_200 shape=Other bindingType=Any
													__local_var_10_200 := (v3_3).V2
													_ = __local_var_10_200
													__t211 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_194, __local_var_8_198, __local_var_5_195}), __local_var_9_199, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_196, __local_var_10_200, __local_var_7_197})})
													goto end_branch_211
												} else {

												}
											}
											{
												var __t_tag_201 *Constructor_Main_Branch = (v3_3).V3
												_ = __t_tag_201
												var __t_and_203 bool = false
												if __t_tag_201 != nil {

													var __t_tag_202 uint32 = ((v3_3).V3).V0
													_ = __t_tag_202
													__t_and_203 = (uint32(__t_tag_202) == 2247809753)
												}
												if __t_and_203 {
													// TAST (Let): __local_var_4_204 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
													__local_var_4_204 := v1_1
													_ = __local_var_4_204
													// TAST (Let): __local_var_5_205 shape=Other bindingType=Any
													__local_var_5_205 := (v3_3).V1
													_ = __local_var_5_205
													// TAST (Let): __local_var_6_206 shape=Other bindingType=Any
													__local_var_6_206 := ((v3_3).V3).V1
													_ = __local_var_6_206
													// TAST (Let): __local_var_7_207 shape=Other bindingType=Any
													__local_var_7_207 := ((v3_3).V3).V3
													_ = __local_var_7_207
													// TAST (Let): __local_var_8_208 shape=Other bindingType=Int
													__local_var_8_208 := v2_2
													_ = __local_var_8_208
													// TAST (Let): __local_var_9_209 shape=Other bindingType=Any
													__local_var_9_209 := (v3_3).V2
													_ = __local_var_9_209
													// TAST (Let): __local_var_10_210 shape=Other bindingType=Any
													__local_var_10_210 := ((v3_3).V3).V2
													_ = __local_var_10_210
													__t211 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_204, __local_var_8_208, __local_var_5_205}), __local_var_9_209, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_206, __local_var_10_210, __local_var_7_207})})
													goto end_branch_211
												} else {

												}
											}
											{
												// TAST (Let): __local_var_4_189 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_4_189 := v1_1
												_ = __local_var_4_189
												// TAST (Let): __local_var_5_190 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
												__local_var_5_190 := v3_3
												_ = __local_var_5_190
												// TAST (Let): __local_var_6_191 shape=Other bindingType=(ADT ["Main","Paint"] [])
												__local_var_6_191 := v_0
												_ = __local_var_6_191
												// TAST (Let): __local_var_7_192 shape=Other bindingType=Int
												__local_var_7_192 := v2_2
												_ = __local_var_7_192
												__t211 = (&Constructor_Main_Branch{1, __local_var_6_191, __local_var_4_189, __local_var_7_192, __local_var_5_190})
											}
										end_branch_211:
											__t222 = __t211
											goto end_branch_222
										} else {

										}
									}
									{
										var __t_tag_212 *Constructor_Main_Branch = (v3_3).V3
										_ = __t_tag_212
										var __t_and_214 bool = false
										if __t_tag_212 != nil {

											var __t_tag_213 uint32 = ((v3_3).V3).V0
											_ = __t_tag_213
											__t_and_214 = (uint32(__t_tag_213) == 2247809753)
										}
										if __t_and_214 {
											// TAST (Let): __local_var_4_215 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
											__local_var_4_215 := v1_1
											_ = __local_var_4_215
											// TAST (Let): __local_var_5_216 shape=Other bindingType=Any
											__local_var_5_216 := (v3_3).V1
											_ = __local_var_5_216
											// TAST (Let): __local_var_6_217 shape=Other bindingType=Any
											__local_var_6_217 := ((v3_3).V3).V1
											_ = __local_var_6_217
											// TAST (Let): __local_var_7_218 shape=Other bindingType=Any
											__local_var_7_218 := ((v3_3).V3).V3
											_ = __local_var_7_218
											// TAST (Let): __local_var_8_219 shape=Other bindingType=Int
											__local_var_8_219 := v2_2
											_ = __local_var_8_219
											// TAST (Let): __local_var_9_220 shape=Other bindingType=Any
											__local_var_9_220 := (v3_3).V2
											_ = __local_var_9_220
											// TAST (Let): __local_var_10_221 shape=Other bindingType=Any
											__local_var_10_221 := ((v3_3).V3).V2
											_ = __local_var_10_221
											__t222 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_215, __local_var_8_219, __local_var_5_216}), __local_var_9_220, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_217, __local_var_10_221, __local_var_7_218})})
											goto end_branch_222
										} else {

										}
									}
									{
										// TAST (Let): __local_var_4_184 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_4_184 := v1_1
										_ = __local_var_4_184
										// TAST (Let): __local_var_5_185 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_5_185 := v3_3
										_ = __local_var_5_185
										// TAST (Let): __local_var_6_186 shape=Other bindingType=(ADT ["Main","Paint"] [])
										__local_var_6_186 := v_0
										_ = __local_var_6_186
										// TAST (Let): __local_var_7_187 shape=Other bindingType=Int
										__local_var_7_187 := v2_2
										_ = __local_var_7_187
										__t222 = (&Constructor_Main_Branch{1, __local_var_6_186, __local_var_4_184, __local_var_7_187, __local_var_5_185})
									}
								end_branch_222:
									__t223 = __t222
									goto end_branch_223
								} else {

								}
							}
							{
								// TAST (Let): __local_var_4_13 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
								__local_var_4_13 := v1_1
								_ = __local_var_4_13
								// TAST (Let): __local_var_5_14 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
								__local_var_5_14 := v3_3
								_ = __local_var_5_14
								// TAST (Let): __local_var_6_15 shape=Other bindingType=(ADT ["Main","Paint"] [])
								__local_var_6_15 := v_0
								_ = __local_var_6_15
								// TAST (Let): __local_var_7_16 shape=Other bindingType=Int
								__local_var_7_16 := v2_2
								_ = __local_var_7_16
								__t223 = (&Constructor_Main_Branch{1, __local_var_6_15, __local_var_4_13, __local_var_7_16, __local_var_5_14})
							}
						end_branch_223:
							__t265 = __t223
							goto end_branch_265
						} else {

						}
					}
					{
						var __t_and_225 bool = false
						if v3_3 != nil {

							var __t_tag_224 uint32 = (v3_3).V0
							_ = __t_tag_224
							__t_and_225 = (uint32(__t_tag_224) == 2247809753)
						}
						if __t_and_225 {
							var __t264 *Constructor_Main_Branch
							{
								var __t_tag_230 *Constructor_Main_Branch = (v3_3).V1
								_ = __t_tag_230
								if __t_tag_230 != nil {
									var __t253 *Constructor_Main_Branch
									{
										var __t_tag_235 uint32 = ((v3_3).V1).V0
										_ = __t_tag_235
										if uint32(__t_tag_235) == 2247809753 {
											// TAST (Let): __local_var_4_236 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
											__local_var_4_236 := v1_1
											_ = __local_var_4_236
											// TAST (Let): __local_var_5_237 shape=Other bindingType=Any
											__local_var_5_237 := ((v3_3).V1).V1
											_ = __local_var_5_237
											// TAST (Let): __local_var_6_238 shape=Other bindingType=Any
											__local_var_6_238 := ((v3_3).V1).V3
											_ = __local_var_6_238
											// TAST (Let): __local_var_7_239 shape=Other bindingType=Any
											__local_var_7_239 := (v3_3).V3
											_ = __local_var_7_239
											// TAST (Let): __local_var_8_240 shape=Other bindingType=Int
											__local_var_8_240 := v2_2
											_ = __local_var_8_240
											// TAST (Let): __local_var_9_241 shape=Other bindingType=Any
											__local_var_9_241 := ((v3_3).V1).V2
											_ = __local_var_9_241
											// TAST (Let): __local_var_10_242 shape=Other bindingType=Any
											__local_var_10_242 := (v3_3).V2
											_ = __local_var_10_242
											__t253 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_236, __local_var_8_240, __local_var_5_237}), __local_var_9_241, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_238, __local_var_10_242, __local_var_7_239})})
											goto end_branch_253
										} else {

										}
									}
									{
										var __t_tag_243 *Constructor_Main_Branch = (v3_3).V3
										_ = __t_tag_243
										var __t_and_245 bool = false
										if __t_tag_243 != nil {

											var __t_tag_244 uint32 = ((v3_3).V3).V0
											_ = __t_tag_244
											__t_and_245 = (uint32(__t_tag_244) == 2247809753)
										}
										if __t_and_245 {
											// TAST (Let): __local_var_4_246 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
											__local_var_4_246 := v1_1
											_ = __local_var_4_246
											// TAST (Let): __local_var_5_247 shape=Other bindingType=Any
											__local_var_5_247 := (v3_3).V1
											_ = __local_var_5_247
											// TAST (Let): __local_var_6_248 shape=Other bindingType=Any
											__local_var_6_248 := ((v3_3).V3).V1
											_ = __local_var_6_248
											// TAST (Let): __local_var_7_249 shape=Other bindingType=Any
											__local_var_7_249 := ((v3_3).V3).V3
											_ = __local_var_7_249
											// TAST (Let): __local_var_8_250 shape=Other bindingType=Int
											__local_var_8_250 := v2_2
											_ = __local_var_8_250
											// TAST (Let): __local_var_9_251 shape=Other bindingType=Any
											__local_var_9_251 := (v3_3).V2
											_ = __local_var_9_251
											// TAST (Let): __local_var_10_252 shape=Other bindingType=Any
											__local_var_10_252 := ((v3_3).V3).V2
											_ = __local_var_10_252
											__t253 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_246, __local_var_8_250, __local_var_5_247}), __local_var_9_251, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_248, __local_var_10_252, __local_var_7_249})})
											goto end_branch_253
										} else {

										}
									}
									{
										// TAST (Let): __local_var_4_231 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_4_231 := v1_1
										_ = __local_var_4_231
										// TAST (Let): __local_var_5_232 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
										__local_var_5_232 := v3_3
										_ = __local_var_5_232
										// TAST (Let): __local_var_6_233 shape=Other bindingType=(ADT ["Main","Paint"] [])
										__local_var_6_233 := v_0
										_ = __local_var_6_233
										// TAST (Let): __local_var_7_234 shape=Other bindingType=Int
										__local_var_7_234 := v2_2
										_ = __local_var_7_234
										__t253 = (&Constructor_Main_Branch{1, __local_var_6_233, __local_var_4_231, __local_var_7_234, __local_var_5_232})
									}
								end_branch_253:
									__t264 = __t253
									goto end_branch_264
								} else {

								}
							}
							{
								var __t_tag_254 *Constructor_Main_Branch = (v3_3).V3
								_ = __t_tag_254
								var __t_and_256 bool = false
								if __t_tag_254 != nil {

									var __t_tag_255 uint32 = ((v3_3).V3).V0
									_ = __t_tag_255
									__t_and_256 = (uint32(__t_tag_255) == 2247809753)
								}
								if __t_and_256 {
									// TAST (Let): __local_var_4_257 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
									__local_var_4_257 := v1_1
									_ = __local_var_4_257
									// TAST (Let): __local_var_5_258 shape=Other bindingType=Any
									__local_var_5_258 := (v3_3).V1
									_ = __local_var_5_258
									// TAST (Let): __local_var_6_259 shape=Other bindingType=Any
									__local_var_6_259 := ((v3_3).V3).V1
									_ = __local_var_6_259
									// TAST (Let): __local_var_7_260 shape=Other bindingType=Any
									__local_var_7_260 := ((v3_3).V3).V3
									_ = __local_var_7_260
									// TAST (Let): __local_var_8_261 shape=Other bindingType=Int
									__local_var_8_261 := v2_2
									_ = __local_var_8_261
									// TAST (Let): __local_var_9_262 shape=Other bindingType=Any
									__local_var_9_262 := (v3_3).V2
									_ = __local_var_9_262
									// TAST (Let): __local_var_10_263 shape=Other bindingType=Any
									__local_var_10_263 := ((v3_3).V3).V2
									_ = __local_var_10_263
									__t264 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_257, __local_var_8_261, __local_var_5_258}), __local_var_9_262, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_259, __local_var_10_263, __local_var_7_260})})
									goto end_branch_264
								} else {

								}
							}
							{
								// TAST (Let): __local_var_4_226 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
								__local_var_4_226 := v1_1
								_ = __local_var_4_226
								// TAST (Let): __local_var_5_227 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
								__local_var_5_227 := v3_3
								_ = __local_var_5_227
								// TAST (Let): __local_var_6_228 shape=Other bindingType=(ADT ["Main","Paint"] [])
								__local_var_6_228 := v_0
								_ = __local_var_6_228
								// TAST (Let): __local_var_7_229 shape=Other bindingType=Int
								__local_var_7_229 := v2_2
								_ = __local_var_7_229
								__t264 = (&Constructor_Main_Branch{1, __local_var_6_228, __local_var_4_226, __local_var_7_229, __local_var_5_227})
							}
						end_branch_264:
							__t265 = __t264
							goto end_branch_265
						} else {

						}
					}
					{
						// TAST (Let): __local_var_4_8 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
						__local_var_4_8 := v1_1
						_ = __local_var_4_8
						// TAST (Let): __local_var_5_9 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
						__local_var_5_9 := v3_3
						_ = __local_var_5_9
						// TAST (Let): __local_var_6_10 shape=Other bindingType=(ADT ["Main","Paint"] [])
						__local_var_6_10 := v_0
						_ = __local_var_6_10
						// TAST (Let): __local_var_7_11 shape=Other bindingType=Int
						__local_var_7_11 := v2_2
						_ = __local_var_7_11
						__t265 = (&Constructor_Main_Branch{1, __local_var_6_10, __local_var_4_8, __local_var_7_11, __local_var_5_9})
					}
				end_branch_265:
					__t307 = __t265
					goto end_branch_307
				} else {

				}
			}
			{
				var __t_and_267 bool = false
				if v3_3 != nil {

					var __t_tag_266 uint32 = (v3_3).V0
					_ = __t_tag_266
					__t_and_267 = (uint32(__t_tag_266) == 2247809753)
				}
				if __t_and_267 {
					var __t306 *Constructor_Main_Branch
					{
						var __t_tag_272 *Constructor_Main_Branch = (v3_3).V1
						_ = __t_tag_272
						if __t_tag_272 != nil {
							var __t295 *Constructor_Main_Branch
							{
								var __t_tag_277 uint32 = ((v3_3).V1).V0
								_ = __t_tag_277
								if uint32(__t_tag_277) == 2247809753 {
									// TAST (Let): __local_var_4_278 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
									__local_var_4_278 := v1_1
									_ = __local_var_4_278
									// TAST (Let): __local_var_5_279 shape=Other bindingType=Any
									__local_var_5_279 := ((v3_3).V1).V1
									_ = __local_var_5_279
									// TAST (Let): __local_var_6_280 shape=Other bindingType=Any
									__local_var_6_280 := ((v3_3).V1).V3
									_ = __local_var_6_280
									// TAST (Let): __local_var_7_281 shape=Other bindingType=Any
									__local_var_7_281 := (v3_3).V3
									_ = __local_var_7_281
									// TAST (Let): __local_var_8_282 shape=Other bindingType=Int
									__local_var_8_282 := v2_2
									_ = __local_var_8_282
									// TAST (Let): __local_var_9_283 shape=Other bindingType=Any
									__local_var_9_283 := ((v3_3).V1).V2
									_ = __local_var_9_283
									// TAST (Let): __local_var_10_284 shape=Other bindingType=Any
									__local_var_10_284 := (v3_3).V2
									_ = __local_var_10_284
									__t295 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_278, __local_var_8_282, __local_var_5_279}), __local_var_9_283, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_280, __local_var_10_284, __local_var_7_281})})
									goto end_branch_295
								} else {

								}
							}
							{
								var __t_tag_285 *Constructor_Main_Branch = (v3_3).V3
								_ = __t_tag_285
								var __t_and_287 bool = false
								if __t_tag_285 != nil {

									var __t_tag_286 uint32 = ((v3_3).V3).V0
									_ = __t_tag_286
									__t_and_287 = (uint32(__t_tag_286) == 2247809753)
								}
								if __t_and_287 {
									// TAST (Let): __local_var_4_288 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
									__local_var_4_288 := v1_1
									_ = __local_var_4_288
									// TAST (Let): __local_var_5_289 shape=Other bindingType=Any
									__local_var_5_289 := (v3_3).V1
									_ = __local_var_5_289
									// TAST (Let): __local_var_6_290 shape=Other bindingType=Any
									__local_var_6_290 := ((v3_3).V3).V1
									_ = __local_var_6_290
									// TAST (Let): __local_var_7_291 shape=Other bindingType=Any
									__local_var_7_291 := ((v3_3).V3).V3
									_ = __local_var_7_291
									// TAST (Let): __local_var_8_292 shape=Other bindingType=Int
									__local_var_8_292 := v2_2
									_ = __local_var_8_292
									// TAST (Let): __local_var_9_293 shape=Other bindingType=Any
									__local_var_9_293 := (v3_3).V2
									_ = __local_var_9_293
									// TAST (Let): __local_var_10_294 shape=Other bindingType=Any
									__local_var_10_294 := ((v3_3).V3).V2
									_ = __local_var_10_294
									__t295 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_288, __local_var_8_292, __local_var_5_289}), __local_var_9_293, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_290, __local_var_10_294, __local_var_7_291})})
									goto end_branch_295
								} else {

								}
							}
							{
								// TAST (Let): __local_var_4_273 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
								__local_var_4_273 := v1_1
								_ = __local_var_4_273
								// TAST (Let): __local_var_5_274 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
								__local_var_5_274 := v3_3
								_ = __local_var_5_274
								// TAST (Let): __local_var_6_275 shape=Other bindingType=(ADT ["Main","Paint"] [])
								__local_var_6_275 := v_0
								_ = __local_var_6_275
								// TAST (Let): __local_var_7_276 shape=Other bindingType=Int
								__local_var_7_276 := v2_2
								_ = __local_var_7_276
								__t295 = (&Constructor_Main_Branch{1, __local_var_6_275, __local_var_4_273, __local_var_7_276, __local_var_5_274})
							}
						end_branch_295:
							__t306 = __t295
							goto end_branch_306
						} else {

						}
					}
					{
						var __t_tag_296 *Constructor_Main_Branch = (v3_3).V3
						_ = __t_tag_296
						var __t_and_298 bool = false
						if __t_tag_296 != nil {

							var __t_tag_297 uint32 = ((v3_3).V3).V0
							_ = __t_tag_297
							__t_and_298 = (uint32(__t_tag_297) == 2247809753)
						}
						if __t_and_298 {
							// TAST (Let): __local_var_4_299 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
							__local_var_4_299 := v1_1
							_ = __local_var_4_299
							// TAST (Let): __local_var_5_300 shape=Other bindingType=Any
							__local_var_5_300 := (v3_3).V1
							_ = __local_var_5_300
							// TAST (Let): __local_var_6_301 shape=Other bindingType=Any
							__local_var_6_301 := ((v3_3).V3).V1
							_ = __local_var_6_301
							// TAST (Let): __local_var_7_302 shape=Other bindingType=Any
							__local_var_7_302 := ((v3_3).V3).V3
							_ = __local_var_7_302
							// TAST (Let): __local_var_8_303 shape=Other bindingType=Int
							__local_var_8_303 := v2_2
							_ = __local_var_8_303
							// TAST (Let): __local_var_9_304 shape=Other bindingType=Any
							__local_var_9_304 := (v3_3).V2
							_ = __local_var_9_304
							// TAST (Let): __local_var_10_305 shape=Other bindingType=Any
							__local_var_10_305 := ((v3_3).V3).V2
							_ = __local_var_10_305
							__t306 = (&Constructor_Main_Branch{1, 2247809753, (&Constructor_Main_Branch{1, 1685833310, __local_var_4_299, __local_var_8_303, __local_var_5_300}), __local_var_9_304, (&Constructor_Main_Branch{1, 1685833310, __local_var_6_301, __local_var_10_305, __local_var_7_302})})
							goto end_branch_306
						} else {

						}
					}
					{
						// TAST (Let): __local_var_4_268 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
						__local_var_4_268 := v1_1
						_ = __local_var_4_268
						// TAST (Let): __local_var_5_269 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
						__local_var_5_269 := v3_3
						_ = __local_var_5_269
						// TAST (Let): __local_var_6_270 shape=Other bindingType=(ADT ["Main","Paint"] [])
						__local_var_6_270 := v_0
						_ = __local_var_6_270
						// TAST (Let): __local_var_7_271 shape=Other bindingType=Int
						__local_var_7_271 := v2_2
						_ = __local_var_7_271
						__t306 = (&Constructor_Main_Branch{1, __local_var_6_270, __local_var_4_268, __local_var_7_271, __local_var_5_269})
					}
				end_branch_306:
					__t307 = __t306
					goto end_branch_307
				} else {

				}
			}
			{
				// TAST (Let): __local_var_4_4 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
				__local_var_4_4 := v1_1
				_ = __local_var_4_4
				// TAST (Let): __local_var_5_5 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
				__local_var_5_5 := v3_3
				_ = __local_var_5_5
				// TAST (Let): __local_var_6_6 shape=Other bindingType=(ADT ["Main","Paint"] [])
				__local_var_6_6 := v_0
				_ = __local_var_6_6
				// TAST (Let): __local_var_7_7 shape=Other bindingType=Int
				__local_var_7_7 := v2_2
				_ = __local_var_7_7
				__t307 = (&Constructor_Main_Branch{1, __local_var_6_6, __local_var_4_4, __local_var_7_7, __local_var_5_5})
			}
		end_branch_307:
			__t308 = __t307
			goto end_branch_308
		} else {

		}
	}
	{
		// TAST (Let): __local_var_4_0 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
		__local_var_4_0 := v1_1
		_ = __local_var_4_0
		// TAST (Let): __local_var_5_1 shape=Other bindingType=(ADT ["Main","SearchTree"] [])
		__local_var_5_1 := v3_3
		_ = __local_var_5_1
		// TAST (Let): __local_var_6_2 shape=Other bindingType=(ADT ["Main","Paint"] [])
		__local_var_6_2 := v_0
		_ = __local_var_6_2
		// TAST (Let): __local_var_7_3 shape=Other bindingType=Int
		__local_var_7_3 := v2_2
		_ = __local_var_7_3
		__t308 = (&Constructor_Main_Branch{1, __local_var_6_2, __local_var_4_0, __local_var_7_3, __local_var_5_1})
	}
end_branch_308:
	return __t308
}

func Call_Main_leftChild(v_0_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
	var v_0 *Constructor_Main_Branch = v_0_loop
	_ = v_0
	var __t0 *Constructor_Main_Branch
	{
		if v_0 == nil {
			__t0 = (*Constructor_Main_Branch)(nil)
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 != nil {
			__t0 = (v_0).V1
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() *Constructor_Main_Branch { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Main_isRed(v_0_loop *Constructor_Main_Branch) bool {
	var v_0 *Constructor_Main_Branch = v_0_loop
	_ = v_0
	var __t_and_1 bool = false
	if v_0 != nil {

		var __t_tag_0 uint32 = (v_0).V0
		_ = __t_tag_0
		__t_and_1 = (uint32(__t_tag_0) == 2247809753)
	}
	return __t_and_1
}

func Call_Main_descend(v_0_loop int64, v1_1_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
descend:
	for {
		if false {
			continue descend
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Main_Branch = v1_1_loop
		_ = v1_1
		var __t1 *Constructor_Main_Branch
		{
			if v1_1 == nil {
				__t1 = (&Constructor_Main_Branch{1, 2247809753, (*Constructor_Main_Branch)(nil), v_0, (*Constructor_Main_Branch)(nil)})
				goto end_branch_1
			} else {

			}
		}
		{
			if v1_1 != nil {
				var __t0 *Constructor_Main_Branch
				{
					if (v_0) < ((v1_1).V2) {
						__t0 = Call_Main_rebalance((v1_1).V0, Call_Main_descend(v_0, (v1_1).V1), (v1_1).V2, (v1_1).V3)
						goto end_branch_0
					} else {

					}
				}
				{
					if (v_0) > ((v1_1).V2) {
						__t0 = Call_Main_rebalance((v1_1).V0, (v1_1).V1, (v1_1).V2, Call_Main_descend(v_0, (v1_1).V3))
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = (&Constructor_Main_Branch{1, (v1_1).V0, (v1_1).V1, (v1_1).V2, (v1_1).V3})
				}
			end_branch_0:
				__t1 = __t0
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() *Constructor_Main_Branch { panic("Failed pattern match") }()
		}
	end_branch_1:
		return __t1
	}
}

func Call_Main_blacken(v_0_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
	var v_0 *Constructor_Main_Branch = v_0_loop
	_ = v_0
	var __t1 *Constructor_Main_Branch
	{
		if v_0 == nil {
			__t1 = (*Constructor_Main_Branch)(nil)
			goto end_branch_1
		} else {

		}
	}
	{
		if v_0 != nil {
			var __reuse_0 *Constructor_Main_Branch
			if ((v_0) != (nil)) && (((v_0).V0) == (1685833310)) {
				__reuse_0 = v_0
			} else {
				__reuse_0 = (&Constructor_Main_Branch{1, 1685833310, (v_0).V1, (v_0).V2, (v_0).V3})
			}
			__t1 = __reuse_0
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() *Constructor_Main_Branch { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Main_put(value_0_loop int64, tree_1_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
	var value_0 int64 = value_0_loop
	_ = value_0
	var tree_1 *Constructor_Main_Branch = tree_1_loop
	_ = tree_1
	return Call_Main_blacken(Call_Main_descend(value_0, tree_1))
}

func Call_Main_build(v_0_loop int64, v1_1_loop int64, v2_2_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
build:
	for {
		if false {
			continue build
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var v2_2 *Constructor_Main_Branch = v2_2_loop
		_ = v2_2
		var __t0 *Constructor_Main_Branch
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
			v2_2_loop = Call_Main_put((v_0)*(v1_1), v2_2)
			continue build
			__t0 = func() *Constructor_Main_Branch { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_graftAndPut(child_0_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
	var child_0 *Constructor_Main_Branch = child_0_loop
	_ = child_0
	return Call_Main_put(int64(5), (&Constructor_Main_Branch{1, 1685833310, child_0, int64(50), (&Constructor_Main_Branch{1, 1685833310, (*Constructor_Main_Branch)(nil), int64(90), (*Constructor_Main_Branch)(nil)})}))
}

func Call_Main_mixed(v_0_loop int64, v1_1_loop *Constructor_Main_Branch) *Constructor_Main_Branch {
mixed:
	for {
		if false {
			continue mixed
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Main_Branch = v1_1_loop
		_ = v1_1
		var __t0 *Constructor_Main_Branch
		{
			if (v_0) == (int64(0)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = Call_Main_put(gopurs_runtime.IntMod((v_0)*(int64(37)), int64(127)), v1_1)
			continue mixed
			__t0 = func() *Constructor_Main_Branch { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_audit(v_0_loop int64, v1_1_loop int64, v2_2_loop *Constructor_Main_Branch) struct {
	blackHeight int64
	size        int64
	valid       bool
} {
audit:
	for {
		if false {
			continue audit
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var v2_2 *Constructor_Main_Branch = v2_2_loop
		_ = v2_2
		var __t20 struct {
			blackHeight int64
			size        int64
			valid       bool
		}
		{
			if v2_2 == nil {
				__t20 = struct {
					blackHeight int64
					size        int64
					valid       bool
				}{int64(1), int64(0), true}
				goto end_branch_20
			} else {

			}
		}
		{
			if v2_2 != nil {
				// TAST (Let): r_3_0 shape=App(Var) bindingType=(Record (Row [valid: Boolean, blackHeight: Int, size: Int] Empty))
				r_3_0 := Call_Main_audit((v2_2).V2, v1_1, (v2_2).V3)
				_ = r_3_0
				// TAST (Let): l_4_1 shape=App(Var) bindingType=(Record (Row [valid: Boolean, blackHeight: Int, size: Int] Empty))
				l_4_1 := Call_Main_audit(v_0, (v2_2).V2, (v2_2).V1)
				_ = l_4_1
				var __t4 int64
				{
					var __t_tag_2 uint32 = (v2_2).V0
					_ = __t_tag_2
					if uint32(__t_tag_2) == 1685833310 {
						__t4 = (l_4_1.blackHeight) + (int64(1))
						goto end_branch_4
					} else {

					}
				}
				{
					var __t_tag_3 uint32 = (v2_2).V0
					_ = __t_tag_3
					if uint32(__t_tag_3) == 2247809753 {
						__t4 = (l_4_1.blackHeight) + (int64(0))
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() int64 { panic("Failed pattern match") }()
				}
			end_branch_4:
				var __t_and_19 bool = false
				if l_4_1.valid {

					var __t_and_18 bool = false
					if r_3_0.valid {

						var __t_and_17 bool = false
						if (v_0) < ((v2_2).V2) {

							var __t_and_16 bool = false
							if ((v2_2).V2) < (v1_1) {

								var __t_and_15 bool = false
								if (l_4_1.blackHeight) == (r_3_0.blackHeight) {

									var __t_and_6 bool = false
									if v2_2 != nil {

										var __t_tag_5 uint32 = (v2_2).V0
										_ = __t_tag_5
										__t_and_6 = (uint32(__t_tag_5) == 2247809753)
									}
									var __t_or_14 bool = true
									if !((__t_and_6) != (true)) {

										var __t_tag_7 *Constructor_Main_Branch = (v2_2).V1
										_ = __t_tag_7
										var __t_and_9 bool = false
										if __t_tag_7 != nil {

											var __t_tag_8 uint32 = ((v2_2).V1).V0
											_ = __t_tag_8
											__t_and_9 = (uint32(__t_tag_8) == 2247809753)
										}
										var __t_and_13 bool = false
										if (__t_and_9) != (true) {

											var __t_tag_10 *Constructor_Main_Branch = (v2_2).V3
											_ = __t_tag_10
											var __t_and_12 bool = false
											if __t_tag_10 != nil {

												var __t_tag_11 uint32 = ((v2_2).V3).V0
												_ = __t_tag_11
												__t_and_12 = (uint32(__t_tag_11) == 2247809753)
											}
											__t_and_13 = (__t_and_12) != (true)
										}
										__t_or_14 = __t_and_13
									}
									__t_and_15 = __t_or_14
								}
								__t_and_16 = __t_and_15
							}
							__t_and_17 = __t_and_16
						}
						__t_and_18 = __t_and_17
					}
					__t_and_19 = __t_and_18
				}
				__t20 = struct {
					blackHeight int64
					size        int64
					valid       bool
				}{__t4, ((l_4_1.size) + (r_3_0.size)) + (int64(1)), __t_and_19}
				goto end_branch_20
			} else {

			}
		}
		{
			__t20 = func() struct {
				blackHeight int64
				size        int64
				valid       bool
			} {
				panic("Failed pattern match")
			}()
		}
	end_branch_20:
		return __t20
	}
}

func Call_Main_check(expectedSize_0_loop int64, tree_1_loop *Constructor_Main_Branch) gopurs_runtime.Value {
	var expectedSize_0 int64 = expectedSize_0_loop
	_ = expectedSize_0
	var tree_1 *Constructor_Main_Branch = tree_1_loop
	_ = tree_1
	// TAST (Let): result_2_0 shape=App(Var) bindingType=(Record (Row [valid: Boolean, blackHeight: Int, size: Int] Empty))
	result_2_0 := Call_Main_audit(int64(-1000000), int64(1000000), tree_1)
	_ = result_2_0
	return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
		actual   bool
		expected bool
	}{result_2_0.valid, true}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t_and_2 bool = false
		if tree_1 != nil {

			var __t_tag_1 uint32 = (tree_1).V0
			_ = __t_tag_1
			__t_and_2 = (uint32(__t_tag_1) == 2247809753)
		}
		return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
			actual   bool
			expected bool
		}{__t_and_2, false}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{result_2_0.size, expectedSize_0})
		}))
	}))
}

func Call_Main_snapshots(v_0_loop int64, v1_1_loop *Constructor_Main_Branch) gopurs_runtime.Value {
snapshots:
	for {
		if false {
			continue snapshots
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Main_Branch = v1_1_loop
		_ = v1_1
		var __t7 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t7 = Call_Main_check(int64(32), v1_1)
				goto end_branch_7
			} else {

			}
		}
		{
			__t7 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
				__local_var_2_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(v1_1)})
				_ = __local_var_2_0
				__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
				_ = __local_var_3_1
				__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_3_1), gopurs_runtime.Value{})
				_ = __local_var_4_2
				__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str(Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_4_2)))), gopurs_runtime.Value{})
				_ = __local_var_5_3
				__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_5_3), gopurs_runtime.Value{})
				_ = __local_var_6_4
				__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_put(v_0, gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_4_2)))}), gopurs_runtime.Value{})
				_ = __local_var_7_5
				__local_var_8_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_7_5), gopurs_runtime.Value{})
				_ = __local_var_8_6
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check((int64(33))-(v_0), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_8_6)), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_snapshots((v_0)-(int64(1)), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_8_6)), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return Call_Test_Assert_assertEqual_prime___4176622598("", struct {
							actual   string
							expected string
						}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_4_2)), __local_var_6_4.StrVal()})
					}))
				})), gopurs_runtime.Value{})
			})
		}
	end_branch_7:
		return __t7
	}
}

func Call_Main_three(a_0_loop int64, b_1_loop int64, c_2_loop int64) gopurs_runtime.Value {
	var a_0 int64 = a_0_loop
	_ = a_0
	var b_1 int64 = b_1_loop
	_ = b_1
	var c_2 int64 = c_2_loop
	_ = c_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
			orig := struct {
				a int64
				b int64
				c int64
			}{a_0, b_1, c_2}
			_ = orig
			return gopurs_runtime.RecordDict3("a", "b", "c", gopurs_runtime.Int(orig.a), gopurs_runtime.Int(orig.b), gopurs_runtime.Int(orig.c))
		}())
		_ = __local_var_3_0
		__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = __local_var_4_1
		__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
		_ = __local_var_5_2
		__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 2447690122, UnsafePtr: unsafe.Pointer(Call_Main_put(gopurs_runtime.RecordGet(__local_var_5_2, "c").IntVal, Call_Main_put(gopurs_runtime.RecordGet(__local_var_5_2, "b").IntVal, Call_Main___gopurs_owned_put_0(gopurs_runtime.RecordGet(__local_var_5_2, "a").IntVal, (*Constructor_Main_Branch)(nil)))))}), gopurs_runtime.Value{})
		_ = __local_var_6_3
		__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_6_3), gopurs_runtime.Value{})
		_ = __local_var_7_4
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_check(int64(3), gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_7_4)), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Assert_assertEqual_prime___4176622598("", struct {
				actual   string
				expected string
			}{Call_Main_render(gopurs_runtime.CoerceToStruct[Constructor_Main_Branch](__local_var_7_4)), "B2(B1(E)(E))(B3(E)(E))"})
		})), gopurs_runtime.Value{})
	})
}
