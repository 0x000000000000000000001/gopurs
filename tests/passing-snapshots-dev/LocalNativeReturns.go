package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_valueIsSymbol gopurs_runtime.Value
var once_Main_valueIsSymbol sync.Once

func Get_Main_valueIsSymbol() gopurs_runtime.Value {
	once_Main_valueIsSymbol.Do(func() {
		cache_Main_valueIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("value")
		}))
	})
	return cache_Main_valueIsSymbol
}

var cache_Main_keepIsSymbol gopurs_runtime.Value
var once_Main_keepIsSymbol sync.Once

func Get_Main_keepIsSymbol() gopurs_runtime.Value {
	once_Main_keepIsSymbol.Do(func() {
		cache_Main_keepIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("keep")
		}))
	})
	return cache_Main_keepIsSymbol
}

var cache_Main_keepItems gopurs_runtime.Value
var once_Main_keepItems sync.Once

func Get_Main_keepItems() gopurs_runtime.Value {
	once_Main_keepItems.Do(func() {
		cache_Main_keepItems = func() gopurs_runtime.Value {
			var Call_local_Main_go__go_0_0_0 func(*Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}], *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}]) *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}]
			_ = Call_local_Main_go__go_0_0_0
			var go__go_0_0_0 gopurs_runtime.Value
			_ = go__go_0_0_0
			Call_local_Main_go__go_0_0_0 = func(acc_1_loop *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}], v_2_loop *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}]) *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}] {
			go__go_0_0_0:
				for {
					if false {
						continue go__go_0_0_0
					}
					var acc_1 *Constructor_Data_List_Types_Cons[struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}] = acc_1_loop
					_ = acc_1
					var v_2 *Constructor_Data_List_Types_Cons[struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}] = v_2_loop
					_ = v_2
					var __t2 *Constructor_Data_List_Types_Cons[struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}]
					{
						if v_2 != nil {
							var __t1 *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
								/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
							}]
							{
								if (v_2).V0.keep {
									acc_1_loop = (&Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
										/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
									}]{1, (v_2).V0, acc_1})
									v_2_loop = (v_2).V1
									continue go__go_0_0_0
									__t1 = func() *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
										/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
									}] {
										panic("unreachable")
									}()
									goto end_branch_1
								} else {

								}
							}
							{
								acc_1_loop = acc_1
								v_2_loop = (v_2).V1
								continue go__go_0_0_0
								__t1 = func() *Constructor_Data_List_Types_Cons[struct {
									keep  bool
									value int64
									/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
								}] {
									panic("unreachable")
								}()
							}
						end_branch_1:
							__t2 = __t1
							goto end_branch_2
						} else {

						}
					}
					{
						if v_2 == nil {
							__t2 = Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_reverse(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(acc_1))})))
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = func() *Constructor_Data_List_Types_Cons[struct {
							keep  bool
							value int64
							/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
						}] {
							panic("Failed pattern match")
						}()
					}
				end_branch_2:
					return __t2
				}
			}
			go__go_0_0_0 = gopurs_runtime.Func(func(acc_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(Call_local_Main_go__go_0_0_0(Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](acc_1_loop_val)), Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)))))}
				})
			})
			return gopurs_runtime.Apply(go__go_0_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993((*Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}])(nil)))})
		}()
	})
	return cache_Main_keepItems
}

var cache_Main_checkRecordReturns gopurs_runtime.Value
var once_Main_checkRecordReturns sync.Once

func Get_Main_checkRecordReturns() gopurs_runtime.Value {
	once_Main_checkRecordReturns.Do(func() {
		cache_Main_checkRecordReturns = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(40)))
			_ = __local_var_0_0
			seedRef_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = seedRef_1_1
			seed_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), seedRef_1_1), gopurs_runtime.Value{})
			_ = seed_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2593544496_3790796878(Rebox_Main_3790796878_2593544496(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_0()))), Rebox_Main_2632088208_1386611502(Rebox_Main_1386611502_2632088208(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_1())))), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}
					expected struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{struct {
					keep  bool
					value int64
					/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
				}{true, gopurs_runtime.IntAdd(seed_2_2.IntVal, int64(2))}, struct {
					keep  bool
					value int64
					/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
				}{true, int64(42)}}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool(orig.keep), gopurs_runtime.Int(orig.value))
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool(orig.keep), gopurs_runtime.Int(orig.value))
				}())
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			makerRef_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Func(func(offset_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool((offset_4.IntVal) >= (int64(0))), gopurs_runtime.Int(gopurs_runtime.IntAdd(seed_2_2.IntVal, offset_4.IntVal)))
			})), gopurs_runtime.Value{})
			_ = makerRef_4_4
			maker_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), makerRef_4_4), gopurs_runtime.Value{})
			_ = maker_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2593544496_3790796878(Rebox_Main_3790796878_2593544496(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_2()))), Rebox_Main_2632088208_1386611502(Rebox_Main_1386611502_2632088208(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_3())))), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}
					expected struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() struct {
					keep  bool
					value int64
					/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
				} {
					orig := gopurs_runtime.Apply(maker_5_5, gopurs_runtime.Int(int64(5)))
					_ = orig
					clone := struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}{}
					clone.keep = (gopurs_runtime.RecordGet(orig, "keep").IntVal) != (0)
					clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
					return clone
				}(), struct {
					keep  bool
					value int64
					/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
				}{true, int64(45)}}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool(orig.keep), gopurs_runtime.Int(orig.value))
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool(orig.keep), gopurs_runtime.Int(orig.value))
				}())
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			_dollar___unused_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2593544496_3790796878(Rebox_Main_3790796878_2593544496(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_4()))), Rebox_Main_2632088208_1386611502(Rebox_Main_1386611502_2632088208(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_5())))), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}
					expected struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() struct {
					keep  bool
					value int64
					/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
				} {
					orig := gopurs_runtime.Apply(maker_5_5, gopurs_runtime.Int(int64(-1)))
					_ = orig
					clone := struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}{}
					clone.keep = (gopurs_runtime.RecordGet(orig, "keep").IntVal) != (0)
					clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
					return clone
				}(), struct {
					keep  bool
					value int64
					/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
				}{false, int64(39)}}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool(orig.keep), gopurs_runtime.Int(orig.value))
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool(orig.keep), gopurs_runtime.Int(orig.value))
				}())
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			partialRef_8_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Func(func(second_8 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_9_9 shape=Other bindingType=Int
				__local_var_9_9 := gopurs_runtime.IntAdd(int64(2), second_8.IntVal)
				_ = __local_var_9_9
				return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool((__local_var_9_9) >= (int64(0))), gopurs_runtime.Int(gopurs_runtime.IntAdd(seed_2_2.IntVal, __local_var_9_9)))
			})), gopurs_runtime.Value{})
			_ = partialRef_8_8
			partial_9_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), partialRef_8_8), gopurs_runtime.Value{})
			_ = partial_9_10
			_dollar___unused_10_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2593544496_3790796878(Rebox_Main_3790796878_2593544496(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_6()))), Rebox_Main_2632088208_1386611502(Rebox_Main_1386611502_2632088208(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_7())))), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}
					expected struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{func() struct {
					keep  bool
					value int64
					/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
				} {
					orig := gopurs_runtime.Apply(partial_9_10, gopurs_runtime.Int(int64(3)))
					_ = orig
					clone := struct {
						keep  bool
						value int64
						/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
					}{}
					clone.keep = (gopurs_runtime.RecordGet(orig, "keep").IntVal) != (0)
					clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
					return clone
				}(), struct {
					keep  bool
					value int64
					/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
				}{true, int64(45)}}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool(orig.keep), gopurs_runtime.Int(orig.value))
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool(orig.keep), gopurs_runtime.Int(orig.value))
				}())
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_10_11
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("local record returns")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_checkRecordReturns
}

var cache_Main_checkItems gopurs_runtime.Value
var once_Main_checkItems sync.Once

func Get_Main_checkItems() gopurs_runtime.Value {
	once_Main_checkItems.Do(func() {
		cache_Main_checkItems = gopurs_runtime.Func3(func(label_0_box gopurs_runtime.Value, expected_1_box gopurs_runtime.Value, original_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkItems(label_0_box.StrVal(), Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](expected_1_box)), Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](original_2_box)))
		})
	})
	return cache_Main_checkItems
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Main_checkItems("guarded list preserves order", (&Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}]{1, struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}{true, int64(7)}, (&Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}]{1, struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}{true, int64(3)}, (*Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}])(nil)})}), (&Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}]{1, struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}{false, int64(0)}, (&Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}]{1, struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}{true, int64(7)}, (&Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}]{1, struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}{false, int64(2)}, (&Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}]{1, struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}{true, int64(3)}, (*Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}])(nil)})})})}))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(Call_Main_checkItems("empty list", (*Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}])(nil), (*Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}])(nil)), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(Call_Main_checkItems("all rejected", (*Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}])(nil), (&Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}]{1, struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}{false, int64(4)}, (&Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}]{1, struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}{false, int64(8)}, (*Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
				/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
			}])(nil)})})), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(Get_Main_checkRecordReturns(), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2593544496_3790796878(Rebox_Main_3790796878_2593544496(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_keepIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))))))}
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2632088208_1386611502(Rebox_Main_1386611502_2632088208(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_keepIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))))))}
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2593544496_3790796878(Rebox_Main_3790796878_2593544496(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_keepIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))))))}
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2632088208_1386611502(Rebox_Main_1386611502_2632088208(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_keepIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))))))}
	})
	return cache_Main___cached_dict_3
}

var cache_Main___cached_dict_4 gopurs_runtime.Value
var once_Main___cached_dict_4 sync.Once

func Get_Main___cached_dict_4() gopurs_runtime.Value {
	once_Main___cached_dict_4.Do(func() {
		cache_Main___cached_dict_4 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2593544496_3790796878(Rebox_Main_3790796878_2593544496(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_keepIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))))))}
	})
	return cache_Main___cached_dict_4
}

var cache_Main___cached_dict_5 gopurs_runtime.Value
var once_Main___cached_dict_5 sync.Once

func Get_Main___cached_dict_5() gopurs_runtime.Value {
	once_Main___cached_dict_5.Do(func() {
		cache_Main___cached_dict_5 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2632088208_1386611502(Rebox_Main_1386611502_2632088208(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_keepIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))))))}
	})
	return cache_Main___cached_dict_5
}

var cache_Main___cached_dict_6 gopurs_runtime.Value
var once_Main___cached_dict_6 sync.Once

func Get_Main___cached_dict_6() gopurs_runtime.Value {
	once_Main___cached_dict_6.Do(func() {
		cache_Main___cached_dict_6 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2593544496_3790796878(Rebox_Main_3790796878_2593544496(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_keepIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))))))}
	})
	return cache_Main___cached_dict_6
}

var cache_Main___cached_dict_7 gopurs_runtime.Value
var once_Main___cached_dict_7 sync.Once

func Get_Main___cached_dict_7() gopurs_runtime.Value {
	once_Main___cached_dict_7.Do(func() {
		cache_Main___cached_dict_7 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2632088208_1386611502(Rebox_Main_1386611502_2632088208(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_keepIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))))))}
	})
	return cache_Main___cached_dict_7
}

var cache_Main___cached_dict_8 gopurs_runtime.Value
var once_Main___cached_dict_8 sync.Once

func Get_Main___cached_dict_8() gopurs_runtime.Value {
	once_Main___cached_dict_8.Do(func() {
		cache_Main___cached_dict_8 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1965287036_3790796878(Rebox_Main_3790796878_1965287036(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_List_Types_eqList(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_keepIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))})))))))}
	})
	return cache_Main___cached_dict_8
}

var cache_Main___cached_dict_9 gopurs_runtime.Value
var once_Main___cached_dict_9 sync.Once

func Get_Main___cached_dict_9() gopurs_runtime.Value {
	once_Main___cached_dict_9.Do(func() {
		cache_Main___cached_dict_9 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_4087401180_1386611502(Rebox_Main_1386611502_4087401180(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_List_Types_showList(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_keepIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))})))))))}
	})
	return cache_Main___cached_dict_9
}

var cache_Main___cached_dict_10 gopurs_runtime.Value
var once_Main___cached_dict_10 sync.Once

func Get_Main___cached_dict_10() gopurs_runtime.Value {
	once_Main___cached_dict_10.Do(func() {
		cache_Main___cached_dict_10 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1965287036_3790796878(Rebox_Main_3790796878_1965287036(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_List_Types_eqList(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_keepIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))})))))))}
	})
	return cache_Main___cached_dict_10
}

var cache_Main___cached_dict_11 gopurs_runtime.Value
var once_Main___cached_dict_11 sync.Once

func Get_Main___cached_dict_11() gopurs_runtime.Value {
	once_Main___cached_dict_11.Do(func() {
		cache_Main___cached_dict_11 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_4087401180_1386611502(Rebox_Main_1386611502_4087401180(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_List_Types_showList(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_keepIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))})))))))}
	})
	return cache_Main___cached_dict_11
}

func Call_Main_checkItems(label_0_loop string, expected_1_loop *Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
	/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
}], original_2_loop *Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
	/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
}]) gopurs_runtime.Value {
	var label_0 string = label_0_loop
	_ = label_0
	var expected_1 *Constructor_Data_List_Types_Cons[struct {
		keep  bool
		value int64
		/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
	}] = expected_1_loop
	_ = expected_1
	var original_2 *Constructor_Data_List_Types_Cons[struct {
		keep  bool
		value int64
		/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
	}] = original_2_loop
	_ = original_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(original_2))})
		_ = __local_var_3_0
		inputRef_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = inputRef_4_1
		input_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), inputRef_4_1), gopurs_runtime.Value{})
		_ = input_5_2
		_dollar___unused_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1965287036_3790796878(Rebox_Main_3790796878_1965287036(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_8()))), Rebox_Main_4087401180_1386611502(Rebox_Main_1386611502_4087401180(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_9())))), func() gopurs_runtime.Value {
			orig := struct {
				actual *Constructor_Data_List_Types_Cons[struct {
					keep  bool
					value int64
					/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
				}]
				expected *Constructor_Data_List_Types_Cons[struct {
					keep  bool
					value int64
					/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
				}]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Main_keepItems(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](input_5_2))))}))), expected_1}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(orig.actual))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(orig.expected))})
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_6_3
		preserved_7_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), inputRef_4_1), gopurs_runtime.Value{})
		_ = preserved_7_4
		_dollar___unused_8_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1965287036_3790796878(Rebox_Main_3790796878_1965287036(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_10()))), Rebox_Main_4087401180_1386611502(Rebox_Main_1386611502_4087401180(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_11())))), func() gopurs_runtime.Value {
			orig := struct {
				actual *Constructor_Data_List_Types_Cons[struct {
					keep  bool
					value int64
					/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
				}]
				expected *Constructor_Data_List_Types_Cons[struct {
					keep  bool
					value int64
					/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
				}]
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](preserved_7_4)), original_2}
			_ = orig
			return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(orig.actual))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(orig.expected))})
		}()), gopurs_runtime.Value{})
		_ = _dollar___unused_8_5
		return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(label_0)), gopurs_runtime.Value{})
	})
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

func Rebox_Main_1386611502_2632088208(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	keep  bool
	value int64
	/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[struct {
		keep  bool
		value int64
		/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_4087401180(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
	/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
}]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[struct {
		keep  bool
		value int64
		/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
	}]])(unsafe.Pointer(in))
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1965287036_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
	/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
}]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2593544496_3790796878(in *Constructor_Data_Eq_Eq[struct {
	keep  bool
	value int64
	/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2632088208_1386611502(in *Constructor_Data_Show_Show[struct {
	keep  bool
	value int64
	/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1965287036(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
	/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
}]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[struct {
		keep  bool
		value int64
		/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
	}]])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_2593544496(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	keep  bool
	value int64
	/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[struct {
		keep  bool
		value int64
		/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[bool])(unsafe.Pointer(in))
}

func Rebox_Main_4087401180_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
	/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
}]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_849153993_917490743(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
	/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[struct {
		keep  bool
		value int64
		/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
	}]{}
	out.V0 = func() struct {
		keep  bool
		value int64
		/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			keep  bool
			value int64
			/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
		}{}
		clone.keep = (gopurs_runtime.RecordGet(orig, "keep").IntVal) != (0)
		clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
		return clone
	}()
	out.V1 = Rebox_Main_849153993_917490743(in.V1)
	return out
}

func Rebox_Main_917490743_849153993(in *Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
	/*gopurs:record-v1:2:006b006500650070,00760061006c00750065*/
}]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool(orig.keep), gopurs_runtime.Int(orig.value))
	}()
	out.V1 = Rebox_Main_917490743_849153993(in.V1)
	return out
}
