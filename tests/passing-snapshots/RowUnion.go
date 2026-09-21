package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Subrow_dollar_Dict gopurs_runtime.Value
var once_Main_Subrow_dollar_Dict sync.Once

func Get_Main_Subrow_dollar_Dict() gopurs_runtime.Value {
	once_Main_Subrow_dollar_Dict.Do(func() {
		cache_Main_Subrow_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Subrow_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_Subrow_dollar_Dict
}

var cache_Main_Proxy gopurs_runtime.Value
var once_Main_Proxy sync.Once

func Get_Main_Proxy() gopurs_runtime.Value {
	once_Main_Proxy.Do(func() {
		cache_Main_Proxy = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_Proxy
}

var cache_Main_subrow gopurs_runtime.Value
var once_Main_subrow sync.Once

func Get_Main_subrow() gopurs_runtime.Value {
	once_Main_subrow.Do(func() {
		cache_Main_subrow = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_subrow(_dollar___unused_0_box)
		})
	})
	return cache_Main_subrow
}

var cache_Main_solve gopurs_runtime.Value
var once_Main_solve sync.Once

func Get_Main_solve() gopurs_runtime.Value {
	once_Main_solve.Do(func() {
		cache_Main_solve = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_solve(_dollar___unused_0_box, uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_solve
}

var cache_Main_solve__1443272647 gopurs_runtime.Value
var once_Main_solve__1443272647 sync.Once

func Get_Main_solve__1443272647() gopurs_runtime.Value {
	once_Main_solve__1443272647.Do(func() {
		cache_Main_solve__1443272647 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_solve__1443272647(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_solve__1443272647
}

var cache_Main_solve__1751852117 gopurs_runtime.Value
var once_Main_solve__1751852117 sync.Once

func Get_Main_solve__1751852117() gopurs_runtime.Value {
	once_Main_solve__1751852117.Do(func() {
		cache_Main_solve__1751852117 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_solve__1751852117(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_solve__1751852117
}

var cache_Main_solveUnionBackwardsCons gopurs_runtime.Value
var once_Main_solveUnionBackwardsCons sync.Once

func Get_Main_solveUnionBackwardsCons() gopurs_runtime.Value {
	once_Main_solveUnionBackwardsCons.Do(func() {
		cache_Main_solveUnionBackwardsCons = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_solve(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_solveUnionBackwardsCons
}

var cache_Main_solveUnionBackwardsDblCons gopurs_runtime.Value
var once_Main_solveUnionBackwardsDblCons sync.Once

func Get_Main_solveUnionBackwardsDblCons() gopurs_runtime.Value {
	once_Main_solveUnionBackwardsDblCons.Do(func() {
		cache_Main_solveUnionBackwardsDblCons = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_solve(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_solveUnionBackwardsDblCons
}

var cache_Main_solveUnionBackwardsNil gopurs_runtime.Value
var once_Main_solveUnionBackwardsNil sync.Once

func Get_Main_solveUnionBackwardsNil() gopurs_runtime.Value {
	once_Main_solveUnionBackwardsNil.Do(func() {
		cache_Main_solveUnionBackwardsNil = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_solve(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_solveUnionBackwardsNil
}

var cache_Main_merge gopurs_runtime.Value
var once_Main_merge sync.Once

func Get_Main_merge() gopurs_runtime.Value {
	once_Main_merge.Do(func() {
		cache_Main_merge = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_merge(_dollar___unused_0_box)
		})
	})
	return cache_Main_merge
}

var cache_Main_merge1 gopurs_runtime.Value
var once_Main_merge1 sync.Once

func Get_Main_merge1() gopurs_runtime.Value {
	once_Main_merge1.Do(func() {
		cache_Main_merge1 = Call_Main_merge(gopurs_runtime.Value{})
	})
	return cache_Main_merge1
}

var cache_Main_mergeWithExtras gopurs_runtime.Value
var once_Main_mergeWithExtras sync.Once

func Get_Main_mergeWithExtras() gopurs_runtime.Value {
	once_Main_mergeWithExtras.Do(func() {
		cache_Main_mergeWithExtras = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_mergeWithExtras(_dollar___unused_0_box)
		})
	})
	return cache_Main_mergeWithExtras
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Apply2(Get_Main_mergeImpl(), func() gopurs_runtime.Value {
			orig := struct {
				x int64
			}{int64(1)}
			_ = orig
			return gopurs_runtime.RecordDict1("x", gopurs_runtime.Int(orig.x))
		}(), func() gopurs_runtime.Value {
			orig := struct {
				y bool
			}{true}
			_ = orig
			return gopurs_runtime.RecordDict1("y", gopurs_runtime.Bool(orig.y))
		}())
	})
	return cache_Main_test1
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Apply2(Get_Main_mergeImpl(), func() gopurs_runtime.Value {
			orig := struct {
				x int64
			}{int64(1)}
			_ = orig
			return gopurs_runtime.RecordDict1("x", gopurs_runtime.Int(orig.x))
		}(), func() gopurs_runtime.Value {
			orig := struct {
				x bool
			}{true}
			_ = orig
			return gopurs_runtime.RecordDict1("x", gopurs_runtime.Bool(orig.x))
		}())
	})
	return cache_Main_test2
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test3(x_0_box)
		})
	})
	return cache_Main_test3
}

var cache_Main_test3_prime_ gopurs_runtime.Value
var once_Main_test3_prime_ sync.Once

func Get_Main_test3_prime_() gopurs_runtime.Value {
	once_Main_test3_prime_.Do(func() {
		cache_Main_test3_prime_ = gopurs_runtime.Func2(func(dictUnion_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test3_prime_(dictUnion_0_box, x_1_box)
		})
	})
	return cache_Main_test3_prime_
}

var cache_Main_withDefaults gopurs_runtime.Value
var once_Main_withDefaults sync.Once

func Get_Main_withDefaults() gopurs_runtime.Value {
	once_Main_withDefaults.Do(func() {
		cache_Main_withDefaults = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, p_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_withDefaults(_dollar___unused_0_box, p_1_box)
		})
	})
	return cache_Main_withDefaults
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Apply2(Get_Main_mergeImpl(), gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2))), func() gopurs_runtime.Value {
			orig := struct {
				y int64
				z int64
			}{int64(1), int64(1)}
			_ = orig
			return gopurs_runtime.RecordDict2("y", "z", gopurs_runtime.Int(orig.y), gopurs_runtime.Int(orig.z))
		}())
	})
	return cache_Main_test4
}

var cache_Main_withDefaultsClosed gopurs_runtime.Value
var once_Main_withDefaultsClosed sync.Once

func Get_Main_withDefaultsClosed() gopurs_runtime.Value {
	once_Main_withDefaultsClosed.Do(func() {
		cache_Main_withDefaultsClosed = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, p_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_withDefaultsClosed(_dollar___unused_0_box, _dollar___unused_1_box, p_2_box)
		})
	})
	return cache_Main_withDefaultsClosed
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(gopurs_runtime.RecordGet(Get_Main_test1(), "x").IntVal)).StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 string
			{
				if (gopurs_runtime.RecordGet(Get_Main_test1(), "y").IntVal) != (0) {
					__t0 = "true"
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = "false"
			}
		end_branch_0:
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t0)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t1 string
				{
					if (gopurs_runtime.RecordGet(Get_Main_test1(), "x").IntVal) == (int64(1)) {
						__t1 = "true"
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = "false"
				}
			end_branch_1:
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t1)), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Call_Main_merge(gopurs_runtime.Value{}), func() gopurs_runtime.Value {
						orig := struct {
							x int64
						}{int64(1)}
						_ = orig
						return gopurs_runtime.RecordDict1("x", gopurs_runtime.Int(orig.x))
					}(), func() gopurs_runtime.Value {
						orig := struct {
							x int64
							y bool
							z float64
						}{int64(0), true, 42.0}
						_ = orig
						return gopurs_runtime.RecordDict3("x", "y", "z", gopurs_runtime.Int(orig.x), gopurs_runtime.Bool(orig.y), gopurs_runtime.Float(orig.z))
					}()), "x")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Main_mergeImpl(), gopurs_runtime.RecordDict1("x", gopurs_runtime.Int(int64(1))), func() gopurs_runtime.Value {
							orig := struct {
								y int64
								z int64
							}{int64(1), int64(1)}
							_ = orig
							return gopurs_runtime.RecordDict2("y", "z", gopurs_runtime.Int(orig.y), gopurs_runtime.Int(orig.z))
						}()), "x")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Main_mergeImpl(), gopurs_runtime.RecordDict1("x", gopurs_runtime.Int(int64(1))), func() gopurs_runtime.Value {
								orig := struct {
									y int64
									z int64
								}{int64(1), int64(1)}
								_ = orig
								return gopurs_runtime.RecordDict2("y", "z", gopurs_runtime.Int(orig.y), gopurs_runtime.Int(orig.z))
							}()), "y")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Main_mergeImpl(), gopurs_runtime.RecordDict1("x", gopurs_runtime.Int(int64(1))), func() gopurs_runtime.Value {
									orig := struct {
										y int64
										z int64
									}{int64(1), int64(1)}
									_ = orig
									return gopurs_runtime.RecordDict2("y", "z", gopurs_runtime.Int(orig.y), gopurs_runtime.Int(orig.z))
								}()), "z")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Main_mergeImpl(), gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2))), func() gopurs_runtime.Value {
										orig := struct {
											y int64
											z int64
										}{int64(1), int64(1)}
										_ = orig
										return gopurs_runtime.RecordDict2("y", "z", gopurs_runtime.Int(orig.y), gopurs_runtime.Int(orig.z))
									}()), "x")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Main_mergeImpl(), gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2))), func() gopurs_runtime.Value {
											orig := struct {
												y int64
												z int64
											}{int64(1), int64(1)}
											_ = orig
											return gopurs_runtime.RecordDict2("y", "z", gopurs_runtime.Int(orig.y), gopurs_runtime.Int(orig.z))
										}()), "y")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Main_mergeImpl(), gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2))), func() gopurs_runtime.Value {
												orig := struct {
													y int64
													z int64
												}{int64(1), int64(1)}
												_ = orig
												return gopurs_runtime.RecordDict2("y", "z", gopurs_runtime.Int(orig.y), gopurs_runtime.Int(orig.z))
											}()), "z")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Main_mergeImpl(), gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2))), func() gopurs_runtime.Value {
													orig := struct {
														y int64
														z int64
													}{int64(1), int64(1)}
													_ = orig
													return gopurs_runtime.RecordDict2("y", "z", gopurs_runtime.Int(orig.y), gopurs_runtime.Int(orig.z))
												}()), "x")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Main_mergeImpl(), gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2))), func() gopurs_runtime.Value {
														orig := struct {
															y int64
															z int64
														}{int64(1), int64(1)}
														_ = orig
														return gopurs_runtime.RecordDict2("y", "z", gopurs_runtime.Int(orig.y), gopurs_runtime.Int(orig.z))
													}()), "y")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Main_mergeImpl(), gopurs_runtime.RecordDict2("x", "y", gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(2))), func() gopurs_runtime.Value {
															orig := struct {
																y int64
																z int64
															}{int64(1), int64(1)}
															_ = orig
															return gopurs_runtime.RecordDict2("y", "z", gopurs_runtime.Int(orig.y), gopurs_runtime.Int(orig.z))
														}()), "z")).StrVal())), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
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
	})
	return cache_Main_main
}

type Constructor_Main_Proxy[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Subrow[T_r any, T_s any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3309093968] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Subrow[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Subrow: " + key)
		}
	}
}

func Call_Main_Subrow_dollar_Dict(x_0_loop struct {
}) gopurs_runtime.Value {
	var x_0 struct {
	} = x_0_loop
	_ = x_0
	return func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Main_subrow(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return func() gopurs_runtime.Value {
		orig := struct {
		}{}
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Main_solve(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 227768594
}

func Call_Main_solve__1443272647(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
solve__1443272647:
	for {
		if false {
			continue solve__1443272647
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}

func Call_Main_solve__1751852117(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
solve__1751852117:
	for {
		if false {
			continue solve__1751852117
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}

func Call_Main_merge(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Get_Main_mergeImpl()
}

func Call_Main_mergeWithExtras(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Call_Main_merge(gopurs_runtime.Value{})
}

func Call_Main_test3(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return gopurs_runtime.Apply2(Get_Main_mergeImpl(), func() gopurs_runtime.Value {
		orig := struct {
			x int64
		}{int64(1)}
		_ = orig
		return gopurs_runtime.RecordDict1("x", gopurs_runtime.Int(orig.x))
	}(), x_0)
}

func Call_Main_test3_prime_(dictUnion_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictUnion_0 gopurs_runtime.Value = dictUnion_0_loop
	_ = dictUnion_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply2(Get_Main_mergeImpl(), x_1, func() gopurs_runtime.Value {
		orig := struct {
			x int64
		}{int64(1)}
		_ = orig
		return gopurs_runtime.RecordDict1("x", gopurs_runtime.Int(orig.x))
	}())
}

func Call_Main_withDefaults(_dollar___unused_0_loop gopurs_runtime.Value, p_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var p_1 gopurs_runtime.Value = p_1_loop
	_ = p_1
	return gopurs_runtime.Apply2(Get_Main_mergeImpl(), p_1, func() gopurs_runtime.Value {
		orig := struct {
			y int64
			z int64
		}{int64(1), int64(1)}
		_ = orig
		return gopurs_runtime.RecordDict2("y", "z", gopurs_runtime.Int(orig.y), gopurs_runtime.Int(orig.z))
	}())
}

func Call_Main_withDefaultsClosed(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, p_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var p_2 gopurs_runtime.Value = p_2_loop
	_ = p_2
	return gopurs_runtime.Apply2(Get_Main_mergeImpl(), p_2, func() gopurs_runtime.Value {
		orig := struct {
			y int64
			z int64
		}{int64(1), int64(1)}
		_ = orig
		return gopurs_runtime.RecordDict2("y", "z", gopurs_runtime.Int(orig.y), gopurs_runtime.Int(orig.z))
	}())
}

func Get_Main_mergeImpl() gopurs_runtime.Value {
	return _Gopurs_Main_MergeImpl
}
