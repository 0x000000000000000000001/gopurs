package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_retainedIsSymbol gopurs_runtime.Value
var once_Main_retainedIsSymbol sync.Once

func Get_Main_retainedIsSymbol() gopurs_runtime.Value {
	once_Main_retainedIsSymbol.Do(func() {
		cache_Main_retainedIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("retained")
		}))
	})
	return cache_Main_retainedIsSymbol
}

var cache_Main_eqRec gopurs_runtime.Value
var once_Main_eqRec sync.Once

func Get_Main_eqRec() gopurs_runtime.Value {
	once_Main_eqRec.Do(func() {
		cache_Main_eqRec = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1830338253_3790796878(Rebox_Main_3790796878_1830338253(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_retainedIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))))}
	})
	return cache_Main_eqRec
}

var cache_Main_showRecord gopurs_runtime.Value
var once_Main_showRecord sync.Once

func Get_Main_showRecord() gopurs_runtime.Value {
	once_Main_showRecord.Do(func() {
		cache_Main_showRecord = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_920949869_1386611502(Rebox_Main_1386611502_920949869(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(Get_Main_retainedIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}))))))}
	})
	return cache_Main_showRecord
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqArray
}

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main_showArray
}

var cache_Main_makeC gopurs_runtime.Value
var once_Main_makeC sync.Once

func Get_Main_makeC() gopurs_runtime.Value {
	once_Main_makeC.Do(func() {
		cache_Main_makeC = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_makeC(n_0_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordDict([]string{"active", "id", "name", "tail", "values", "z"}, []gopurs_runtime.Value{gopurs_runtime.Bool(orig.active), gopurs_runtime.Int(orig.id), gopurs_runtime.Str(orig.name), gopurs_runtime.Bool(orig.tail), func() gopurs_runtime.Value {
					arr := orig.values
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Str(orig.z)})
			}()
		})
	})
	return cache_Main_makeC
}

var cache_Main_runC gopurs_runtime.Value
var once_Main_runC sync.Once

func Get_Main_runC() gopurs_runtime.Value {
	once_Main_runC.Do(func() {
		cache_Main_runC = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runC(n_0_box.IntVal))
		})
	})
	return cache_Main_runC
}

var cache_Main_makeB gopurs_runtime.Value
var once_Main_makeB sync.Once

func Get_Main_makeB() gopurs_runtime.Value {
	once_Main_makeB.Do(func() {
		cache_Main_makeB = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_makeB(n_0_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordDict5("active", "before", "id", "name", "nested", gopurs_runtime.Bool(orig.active), gopurs_runtime.Float(orig.before), gopurs_runtime.Int(orig.id), gopurs_runtime.Str(orig.name), func() gopurs_runtime.Value {
					orig := orig.nested
					_ = orig
					return gopurs_runtime.RecordDict1("retained", gopurs_runtime.Str(orig.retained))
				}())
			}()
		})
	})
	return cache_Main_makeB
}

var cache_Main_runB gopurs_runtime.Value
var once_Main_runB sync.Once

func Get_Main_runB() gopurs_runtime.Value {
	once_Main_runB.Do(func() {
		cache_Main_runB = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runB(n_0_box.IntVal))
		})
	})
	return cache_Main_runB
}

var cache_Main_makeA gopurs_runtime.Value
var once_Main_makeA sync.Once

func Get_Main_makeA() gopurs_runtime.Value {
	once_Main_makeA.Do(func() {
		cache_Main_makeA = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_makeA(n_0_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordDict4("active", "extra", "id", "name", gopurs_runtime.Bool(orig.active), gopurs_runtime.Int(orig.extra), gopurs_runtime.Int(orig.id), gopurs_runtime.Str(orig.name))
			}()
		})
	})
	return cache_Main_makeA
}

var cache_Main_runA gopurs_runtime.Value
var once_Main_runA sync.Once

func Get_Main_runA() gopurs_runtime.Value {
	once_Main_runA.Do(func() {
		cache_Main_runA = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runA(n_0_box.IntVal))
		})
	})
	return cache_Main_runA
}

var cache_Main_firstSummary gopurs_runtime.Value
var once_Main_firstSummary sync.Once

func Get_Main_firstSummary() gopurs_runtime.Value {
	once_Main_firstSummary.Do(func() {
		cache_Main_firstSummary = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_firstSummary(n_0_box.IntVal)
				_ = orig
				return gopurs_runtime.RecordDict3("active", "id", "name", gopurs_runtime.Bool(orig.active), gopurs_runtime.Int(orig.id), gopurs_runtime.Str(orig.name))
			}()
		})
	})
	return cache_Main_firstSummary
}

var cache_Main_callDynamic gopurs_runtime.Value
var once_Main_callDynamic sync.Once

func Get_Main_callDynamic() gopurs_runtime.Value {
	once_Main_callDynamic.Do(func() {
		cache_Main_callDynamic = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, value_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_callDynamic(f_0_box, value_1_box))
		})
	})
	return cache_Main_callDynamic
}

var cache_Main_callDynamic__1326697050 gopurs_runtime.Value
var once_Main_callDynamic__1326697050 sync.Once

func Get_Main_callDynamic__1326697050() gopurs_runtime.Value {
	once_Main_callDynamic__1326697050.Do(func() {
		cache_Main_callDynamic__1326697050 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, value_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_callDynamic__1326697050(f_unused_0_box, func() struct {
				active bool
				before float64
				id     int64
				name   string
				nested struct {
					retained string
				}
			} {
				orig := value_1_box
				_ = orig
				clone := struct {
					active bool
					before float64
					id     int64
					name   string
					nested struct {
						retained string
					}
				}{}
				clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
				clone.before = gopurs_runtime.RecordGet(orig, "before").FloatVal()
				clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
				clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
				clone.nested = func() struct {
					retained string
				} {
					orig := gopurs_runtime.RecordGet(orig, "nested")
					_ = orig
					clone := struct {
						retained string
					}{}
					clone.retained = gopurs_runtime.RecordGet(orig, "retained").StrVal()
					return clone
				}()
				return clone
			}()))
		})
	})
	return cache_Main_callDynamic__1326697050
}

var cache_Main_callDynamic__1974792538 gopurs_runtime.Value
var once_Main_callDynamic__1974792538 sync.Once

func Get_Main_callDynamic__1974792538() gopurs_runtime.Value {
	once_Main_callDynamic__1974792538.Do(func() {
		cache_Main_callDynamic__1974792538 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, value_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_callDynamic__1974792538(f_unused_0_box, func() struct {
				active bool
				extra  int64
				id     int64
				name   string
			} {
				orig := value_1_box
				_ = orig
				clone := struct {
					active bool
					extra  int64
					id     int64
					name   string
				}{}
				clone.active = (gopurs_runtime.RecordGet(orig, "active").IntVal) != (0)
				clone.extra = gopurs_runtime.RecordGet(orig, "extra").IntVal
				clone.id = gopurs_runtime.RecordGet(orig, "id").IntVal
				clone.name = gopurs_runtime.RecordGet(orig, "name").StrVal()
				return clone
			}()))
		})
	})
	return cache_Main_callDynamic__1974792538
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
			actual   int64
			expected int64
		}{Call_Main_runA(int64(5)), int64(25)}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_runB(int64(5)), int64(18)}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
					actual   int64
					expected int64
				}{Call_Main_runC(int64(5)), int64(25)}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
						actual   int64
						expected int64
					}{Call_Main_runA(int64(-1)), int64(-1)}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
							actual   int64
							expected int64
						}{Call_Main_runB(int64(-1)), int64(-1)}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
								actual   int64
								expected int64
							}{Call_Main_runC(int64(-1)), int64(-1)}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
								// TAST (Let): a2_6_0 shape=LitRecord bindingType=(Record (Row [id: Int, active: Boolean, extra: Int, name: String] Empty))
								a2_6_0 := struct {
									active bool
									extra  int64
									id     int64
									name   string
								}{true, int64(42), int64(9), "alpha"}
								_ = a2_6_0
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
									actual   int64
									expected int64
								}{Call_Main_runA(int64(9)), int64(37)}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
										actual   int64
										expected int64
									}{Call_Main_runB(int64(5)), int64(18)}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
											actual   int64
											expected int64
										}{Call_Main_runC(int64(10)), int64(40)}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
												actual   int64
												expected int64
											}{a2_6_0.extra, int64(42)}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___82050600("", struct {
													actual struct {
														retained string
													}
													expected struct {
														retained string
													}
												}{struct {
													retained string
												}{"keep"}, struct {
													retained string
												}{"keep"}}), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1772858129("", struct {
														actual   []int64
														expected []int64
													}{[]int64{int64(1), int64(2), int64(3)}, []int64{int64(1), int64(2), int64(3)}}), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
															actual   string
															expected string
														}{a2_6_0.name, "alpha"}), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
																actual   string
																expected string
															}{Call_Main_callDynamic__1974792538(Get_Worker_decode(), struct {
																active bool
																extra  int64
																id     int64
																name   string
															}{true, int64(42), int64(-3), "alpha"}), "negative"}), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
																	actual   string
																	expected string
																}{Call_Main_callDynamic__1326697050(Get_Worker_decode(), struct {
																	active bool
																	before float64
																	id     int64
																	name   string
																	nested struct {
																		retained string
																	}
																}{false, 1.25, int64(3), "beta", struct {
																	retained string
																}{"keep"}}), "beta"}), gopurs_runtime.Func(func(_dollar___unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
																		__local_var_16_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
																			_v := Call_Worker_decode(func(record struct {
																				active bool
																				extra  int64
																				id     int64
																				name   string
																			}) struct {
																				active bool
																				id     int64
																				name   string
																			} {
																				return struct {
																					active bool
																					id     int64
																					name   string
																				}{record.active, record.id, record.name}
																			}(struct {
																				active bool
																				extra  int64
																				id     int64
																				name   string
																			}{true, int64(42), int64(3), "alpha"}))
																			if _v.V2 {
																				return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
																					orig := _v.V1
																					_ = orig
																					return gopurs_runtime.RecordDict3("active", "id", "name", gopurs_runtime.Bool(orig.active), gopurs_runtime.Int(orig.id), gopurs_runtime.Str(orig.name))
																				}()})}
																			}
																			return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
																		}()), gopurs_runtime.Value{})
																		_ = __local_var_16_1
																		__local_var_17_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_16_1), gopurs_runtime.Value{})
																		_ = __local_var_17_2
																		var __t3 string
																		{
																			if __local_var_17_2.Type == 9 && __local_var_17_2.IntVal == 2465973597 {
																				__t3 = gopurs_runtime.RecordGet((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_17_2.UnsafePtr).V0, "name").StrVal()
																				goto end_branch_3
																			} else {

																			}
																		}
																		{
																			if __local_var_17_2.Type == 9 && __local_var_17_2.IntVal == 3711209382 {
																				__t3 = ""
																				goto end_branch_3
																			} else {

																			}
																		}
																		{
																			__t3 = func() string { panic("Failed pattern match") }()
																		}
																	end_branch_3:
																		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
																			actual   string
																			expected string
																		}{__t3, "alpha"}), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
																			// TAST (Let): first_19_4 shape=App(Var) bindingType=(Record (Row [id: Int, name: String, active: Boolean] Empty))
																			first_19_4 := Call_Main_firstSummary(int64(7))
																			_ = first_19_4
																			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
																				actual   int64
																				expected int64
																			}{first_19_4.id, int64(8)}), gopurs_runtime.Func(func(_dollar___unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
																				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___4176622598("", struct {
																					actual   string
																					expected string
																				}{first_19_4.name, "alpha"}), gopurs_runtime.Func(func(_dollar___unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
																					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																						actual   bool
																						expected bool
																					}{first_19_4.active, true}), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
																						// TAST (Let): missing_23_5 shape=App(Var) bindingType=(Record (Row [id: Int, name: String, active: Boolean] Empty))
																						missing_23_5 := Call_Main_firstSummary(int64(-1))
																						_ = missing_23_5
																						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
																							actual   int64
																							expected int64
																						}{missing_23_5.id, int64(-1)}), gopurs_runtime.Func(func(_dollar___unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
																							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
																								actual   bool
																								expected bool
																							}{missing_23_5.active, false}), gopurs_runtime.Func(func(_dollar___unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
																								return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
																							}))
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

func Call_Main_makeC(n_0_loop int64) struct {
	active bool
	id     int64
	name   string
	tail   bool
	values []int64
	z      string
} {
	var n_0 int64 = n_0_loop
	_ = n_0
	return struct {
		active bool
		id     int64
		name   string
		tail   bool
		values []int64
		z      string
	}{true, n_0, "alpha", false, []int64{int64(1), int64(2), int64(3)}, "last"}
}

func Call_Main_runC(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [String, (Record (Row [id: Int, name: String, active: Boolean] Empty))])
	v_1_0 := Call_Worker_decode(func(record struct {
		active bool
		id     int64
		name   string
		tail   bool
		values []int64
		z      string
	}) struct {
		active bool
		id     int64
		name   string
	} {
		return struct {
			active bool
			id     int64
			name   string
		}{record.active, record.id, record.name}
	}(struct {
		active bool
		id     int64
		name   string
		tail   bool
		values []int64
		z      string
	}{true, n_0, "alpha", false, []int64{int64(1), int64(2), int64(3)}, "last"}))
	_ = v_1_0
	var __t2 int64
	{
		if !v_1_0.V2 {
			__t2 = int64(-1)
			goto end_branch_2
		} else {

		}
	}
	{
		if v_1_0.V2 {
			var __t1 int64
			{
				if v_1_0.V1.active {
					__t1 = int64(7)
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = int64(0)
			}
		end_branch_1:
			__t2 = ((v_1_0.V1.id) * (int64(3))) + (__t1)
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = func() int64 { panic("Failed pattern match") }()
	}
end_branch_2:
	return __t2
}

func Call_Main_makeB(n_0_loop int64) struct {
	active bool
	before float64
	id     int64
	name   string
	nested struct {
		retained string
	}
} {
	var n_0 int64 = n_0_loop
	_ = n_0
	return struct {
		active bool
		before float64
		id     int64
		name   string
		nested struct {
			retained string
		}
	}{false, 1.25, n_0, "beta", struct {
		retained string
	}{"keep"}}
}

func Call_Main_runB(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [String, (Record (Row [id: Int, name: String, active: Boolean] Empty))])
	v_1_0 := Call_Worker_decode(func(record struct {
		active bool
		before float64
		id     int64
		name   string
		nested struct {
			retained string
		}
	}) struct {
		active bool
		id     int64
		name   string
	} {
		return struct {
			active bool
			id     int64
			name   string
		}{record.active, record.id, record.name}
	}(struct {
		active bool
		before float64
		id     int64
		name   string
		nested struct {
			retained string
		}
	}{false, 1.25, n_0, "beta", struct {
		retained string
	}{"keep"}}))
	_ = v_1_0
	var __t2 int64
	{
		if !v_1_0.V2 {
			__t2 = int64(-1)
			goto end_branch_2
		} else {

		}
	}
	{
		if v_1_0.V2 {
			var __t1 int64
			{
				if v_1_0.V1.active {
					__t1 = int64(7)
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = int64(0)
			}
		end_branch_1:
			__t2 = ((v_1_0.V1.id) * (int64(3))) + (__t1)
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = func() int64 { panic("Failed pattern match") }()
	}
end_branch_2:
	return __t2
}

func Call_Main_makeA(n_0_loop int64) struct {
	active bool
	extra  int64
	id     int64
	name   string
} {
	var n_0 int64 = n_0_loop
	_ = n_0
	return struct {
		active bool
		extra  int64
		id     int64
		name   string
	}{true, int64(42), n_0, "alpha"}
}

func Call_Main_runA(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [String, (Record (Row [id: Int, name: String, active: Boolean] Empty))])
	v_1_0 := Call_Worker_decode(func(record struct {
		active bool
		extra  int64
		id     int64
		name   string
	}) struct {
		active bool
		id     int64
		name   string
	} {
		return struct {
			active bool
			id     int64
			name   string
		}{record.active, record.id, record.name}
	}(struct {
		active bool
		extra  int64
		id     int64
		name   string
	}{true, int64(42), n_0, "alpha"}))
	_ = v_1_0
	var __t2 int64
	{
		if !v_1_0.V2 {
			__t2 = int64(-1)
			goto end_branch_2
		} else {

		}
	}
	{
		if v_1_0.V2 {
			var __t1 int64
			{
				if v_1_0.V1.active {
					__t1 = int64(7)
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = int64(0)
			}
		end_branch_1:
			__t2 = ((v_1_0.V1.id) * (int64(3))) + (__t1)
			goto end_branch_2
		} else {

		}
	}
	{
		__t2 = func() int64 { panic("Failed pattern match") }()
	}
end_branch_2:
	return __t2
}

func Call_Main_firstSummary(n_0_loop int64) struct {
	active bool
	id     int64
	name   string
} {
	var n_0 int64 = n_0_loop
	_ = n_0
	// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [String, (Record (Row [id: Int, name: String, active: Boolean] Empty))])
	v_1_0 := Call_Worker_decode(func(record struct {
		active bool
		extra  int64
		id     int64
		name   string
	}) struct {
		active bool
		id     int64
		name   string
	} {
		return struct {
			active bool
			id     int64
			name   string
		}{record.active, record.id, record.name}
	}(struct {
		active bool
		extra  int64
		id     int64
		name   string
	}{true, int64(42), n_0, "alpha"}))
	_ = v_1_0
	var __t1 struct {
		active bool
		id     int64
		name   string
	}
	{
		if !v_1_0.V2 {
			__t1 = struct {
				active bool
				id     int64
				name   string
			}{false, int64(-1), ""}
			goto end_branch_1
		} else {

		}
	}
	{
		if v_1_0.V2 {
			__t1 = v_1_0.V1
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() struct {
			active bool
			id     int64
			name   string
		} {
			panic("Failed pattern match")
		}()
	}
end_branch_1:
	return __t1
}

func Call_Main_callDynamic(f_0_loop gopurs_runtime.Value, value_1_loop gopurs_runtime.Value) string {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var value_1 gopurs_runtime.Value = value_1_loop
	_ = value_1
	// TAST (Let): v_2_0 shape=App(Other) bindingType=(ADT ["Data","Either","Either"] [String, (Record (Row [id: Int, name: String, active: Boolean] Empty))])
	v_2_0 := gopurs_runtime.Apply(f_0, value_1)
	_ = v_2_0
	var __t1 string
	{
		if v_2_0.Type == 9 && v_2_0.IntVal == 3711209382 {
			__t1 = (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2_0.UnsafePtr).V0.StrVal()
			goto end_branch_1
		} else {

		}
	}
	{
		if v_2_0.Type == 9 && v_2_0.IntVal == 2465973597 {
			__t1 = gopurs_runtime.RecordGet((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2_0.UnsafePtr).V0, "name").StrVal()
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() string { panic("Failed pattern match") }()
	}
end_branch_1:
	return __t1
}

func Call_Main_callDynamic__1326697050(f_unused_0_loop gopurs_runtime.Value, value_1_loop struct {
	active bool
	before float64
	id     int64
	name   string
	nested struct {
		retained string
	}
}) string {
callDynamic__1326697050:
	for {
		if false {
			continue callDynamic__1326697050
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var value_1 struct {
			active bool
			before float64
			id     int64
			name   string
			nested struct {
				retained string
			}
		} = value_1_loop
		_ = value_1
		// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [String, (Record (Row [id: Int, name: String, active: Boolean] Empty))])
		v_2_0 := Call_Worker_decode(func(record struct {
			active bool
			before float64
			id     int64
			name   string
			nested struct {
				retained string
			}
		}) struct {
			active bool
			id     int64
			name   string
		} {
			return struct {
				active bool
				id     int64
				name   string
			}{record.active, record.id, record.name}
		}(value_1))
		_ = v_2_0
		var __t1 string
		{
			if !v_2_0.V2 {
				__t1 = v_2_0.V0.StrVal()
				goto end_branch_1
			} else {

			}
		}
		{
			if v_2_0.V2 {
				__t1 = v_2_0.V1.name
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() string { panic("Failed pattern match") }()
		}
	end_branch_1:
		return __t1
	}
}

func Call_Main_callDynamic__1974792538(f_unused_0_loop gopurs_runtime.Value, value_1_loop struct {
	active bool
	extra  int64
	id     int64
	name   string
}) string {
callDynamic__1974792538:
	for {
		if false {
			continue callDynamic__1974792538
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var value_1 struct {
			active bool
			extra  int64
			id     int64
			name   string
		} = value_1_loop
		_ = value_1
		// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [String, (Record (Row [id: Int, name: String, active: Boolean] Empty))])
		v_2_0 := Call_Worker_decode(func(record struct {
			active bool
			extra  int64
			id     int64
			name   string
		}) struct {
			active bool
			id     int64
			name   string
		} {
			return struct {
				active bool
				id     int64
				name   string
			}{record.active, record.id, record.name}
		}(value_1))
		_ = v_2_0
		var __t1 string
		{
			if !v_2_0.V2 {
				__t1 = v_2_0.V0.StrVal()
				goto end_branch_1
			} else {

			}
		}
		{
			if v_2_0.V2 {
				__t1 = v_2_0.V1.name
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() string { panic("Failed pattern match") }()
		}
	end_branch_1:
		return __t1
	}
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

func Rebox_Main_1386611502_1469227923(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[[]int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[string]{}
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

func Rebox_Main_1386611502_920949869(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	retained string
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[struct {
		retained string
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1469227923_1386611502(in *Constructor_Data_Show_Show[[]int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
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

func Rebox_Main_1830338253_3790796878(in *Constructor_Data_Eq_Eq[struct {
	retained string
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_378698611_3790796878(in *Constructor_Data_Eq_Eq[[]int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
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

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[string]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_1830338253(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	retained string
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[struct {
		retained string
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_378698611(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[[]int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_920949869_1386611502(in *Constructor_Data_Show_Show[struct {
	retained string
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
