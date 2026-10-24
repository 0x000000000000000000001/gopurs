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

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_runA(int64(5)), int64(25)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}())
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_runB(int64(5)), int64(18)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_runC(int64(5)), int64(25)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_runA(int64(-1)), int64(-1)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_runB(int64(-1)), int64(-1)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_runC(int64(-1)), int64(-1)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			// TAST (Let): a2_7_7 shape=LitRecord bindingType=(Record (Row [id: Int, active: Boolean, extra: Int, name: String] Empty))
			a2_7_7 := struct {
				active bool
				extra  int64
				id     int64
				name   string
				/*gopurs:record-v1:4:006100630074006900760065,00650078007400720061,00690064,006e0061006d0065*/
			}{true, int64(42), int64(9), "alpha"}
			_ = a2_7_7
			_dollar___unused_8_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_runA(int64(9)), int64(37)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_8_8
			_dollar___unused_9_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_runB(int64(5)), int64(18)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_9_9
			_dollar___unused_10_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_runC(int64(10)), int64(40)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_10_10
			_dollar___unused_11_11 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{a2_7_7.extra, int64(42)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_11_11
			_dollar___unused_12_12 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1830338253_3790796878(Rebox_Main_3790796878_1830338253(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_0()))), Rebox_Main_920949869_1386611502(Rebox_Main_1386611502_920949869(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_1())))), func() gopurs_runtime.Value {
				orig := struct {
					actual struct {
						retained string
						/*gopurs:record-v1:1:00720065007400610069006e00650064*/
					}
					expected struct {
						retained string
						/*gopurs:record-v1:1:00720065007400610069006e00650064*/
					}
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{struct {
					retained string
					/*gopurs:record-v1:1:00720065007400610069006e00650064*/
				}{"keep"}, struct {
					retained string
					/*gopurs:record-v1:1:00720065007400610069006e00650064*/
				}{"keep"}}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
					orig := orig.actual
					_ = orig
					return gopurs_runtime.RecordDict1("retained", gopurs_runtime.Str(orig.retained))
				}(), func() gopurs_runtime.Value {
					orig := orig.expected
					_ = orig
					return gopurs_runtime.RecordDict1("retained", gopurs_runtime.Str(orig.retained))
				}())
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_12_12
			_dollar___unused_13_13 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main___cached_dict_2()))), Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Main___cached_dict_3())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   []int64
					expected []int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{[]int64{int64(1), int64(2), int64(3)}, []int64{int64(1), int64(2), int64(3)}}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
					arr := orig.actual
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := orig.expected
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}())
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_13_13
			_dollar___unused_14_14 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{a2_7_7.name, "alpha"}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_14_14
			_dollar___unused_15_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_callDynamic(Get_Worker_decode(), func() gopurs_runtime.Value {
					orig := struct {
						active bool
						extra  int64
						id     int64
						name   string
						/*gopurs:record-v1:4:006100630074006900760065,00650078007400720061,00690064,006e0061006d0065*/
					}{true, int64(42), int64(-3), "alpha"}
					_ = orig
					return gopurs_runtime.RecordDict4("active", "extra", "id", "name", gopurs_runtime.Bool(orig.active), gopurs_runtime.Int(orig.extra), gopurs_runtime.Int(orig.id), gopurs_runtime.Str(orig.name))
				}()), "negative"}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_15_15
			_dollar___unused_16_16 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Call_Main_callDynamic(Get_Worker_decode(), func() gopurs_runtime.Value {
					orig := struct {
						active bool
						before float64
						id     int64
						name   string
						nested struct {
							retained string
							/*gopurs:record-v1:1:00720065007400610069006e00650064*/
						}
						/*gopurs:record-v1:5:006100630074006900760065,006200650066006f00720065,00690064,006e0061006d0065,006e00650073007400650064*/
					}{false, 1.25, int64(3), "beta", struct {
						retained string
						/*gopurs:record-v1:1:00720065007400610069006e00650064*/
					}{"keep"}}
					_ = orig
					return gopurs_runtime.RecordDict5("active", "before", "id", "name", "nested", gopurs_runtime.Bool(orig.active), gopurs_runtime.Float(orig.before), gopurs_runtime.Int(orig.id), gopurs_runtime.Str(orig.name), func() gopurs_runtime.Value {
						orig := orig.nested
						_ = orig
						return gopurs_runtime.RecordDict1("retained", gopurs_runtime.Str(orig.retained))
					}())
				}()), "beta"}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_16_16
			saved_17_17 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_go__new(), func() gopurs_runtime.Value {
				_v := Call_Worker_decode(func(record struct {
					active bool
					extra  int64
					id     int64
					name   string
					/*gopurs:record-v1:4:006100630074006900760065,00650078007400720061,00690064,006e0061006d0065*/
				}) struct {
					active bool
					id     int64
					name   string
					/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
				} {
					return struct {
						active bool
						id     int64
						name   string
						/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
					}{record.active, record.id, record.name}
				}(struct {
					active bool
					extra  int64
					id     int64
					name   string
					/*gopurs:record-v1:4:006100630074006900760065,00650078007400720061,00690064,006e0061006d0065*/
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
			_ = saved_17_17
			held_18_18 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), saved_17_17), gopurs_runtime.Value{})
			_ = held_18_18
			var __t20 string
			{
				if held_18_18.Type == 9 && held_18_18.IntVal == 2465973597 {
					__t20 = gopurs_runtime.RecordGet((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(held_18_18.UnsafePtr).V0, "name").StrVal()
					goto end_branch_20
				} else {

				}
			}
			{
				if held_18_18.Type == 9 && held_18_18.IntVal == 3711209382 {
					__t20 = ""
					goto end_branch_20
				} else {

				}
			}
			{
				__t20 = func() string { panic("Failed pattern match") }()
			}
		end_branch_20:
			_dollar___unused_19_19 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{__t20, "alpha"}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_19_19
			// TAST (Let): first_20_21 shape=App(Var) bindingType=(Record (Row [id: Int, name: String, active: Boolean] Empty))
			first_20_21 := Call_Main_firstSummary(int64(7))
			_ = first_20_21
			_dollar___unused_21_22 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{first_20_21.id, int64(8)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_21_22
			_dollar___unused_22_23 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))), Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   string
					expected string
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{first_20_21.name, "alpha"}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(orig.actual), gopurs_runtime.Str(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_22_23
			_dollar___unused_23_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{first_20_21.active, true}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_23_24
			// TAST (Let): missing_24_25 shape=App(Var) bindingType=(Record (Row [id: Int, name: String, active: Boolean] Empty))
			missing_24_25 := Call_Main_firstSummary(int64(-1))
			_ = missing_24_25
			_dollar___unused_25_26 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{missing_24_25.id, int64(-1)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_25_26
			_dollar___unused_26_27 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   bool
					expected bool
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{missing_24_25.active, false}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_26_27
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1830338253_3790796878(Rebox_Main_3790796878_1830338253(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_retainedIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}))))))}
	})
	return cache_Main___cached_dict_0
}

var cache_Main___cached_dict_1 gopurs_runtime.Value
var once_Main___cached_dict_1 sync.Once

func Get_Main___cached_dict_1() gopurs_runtime.Value {
	once_Main___cached_dict_1.Do(func() {
		cache_Main___cached_dict_1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_920949869_1386611502(Rebox_Main_1386611502_920949869(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsConsNil(Get_Main_retainedIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}))))))}
	})
	return cache_Main___cached_dict_1
}

var cache_Main___cached_dict_2 gopurs_runtime.Value
var once_Main___cached_dict_2 sync.Once

func Get_Main___cached_dict_2() gopurs_runtime.Value {
	once_Main___cached_dict_2.Do(func() {
		cache_Main___cached_dict_2 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main___cached_dict_2
}

var cache_Main___cached_dict_3 gopurs_runtime.Value
var once_Main___cached_dict_3 sync.Once

func Get_Main___cached_dict_3() gopurs_runtime.Value {
	once_Main___cached_dict_3.Do(func() {
		cache_Main___cached_dict_3 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main___cached_dict_3
}

func Call_Main_makeC(n_0_loop int64) struct {
	active bool
	id     int64
	name   string
	tail   bool
	values []int64
	z      string
	/*gopurs:record-v1:6:006100630074006900760065,00690064,006e0061006d0065,007400610069006c,00760061006c007500650073,007a*/
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
		/*gopurs:record-v1:6:006100630074006900760065,00690064,006e0061006d0065,007400610069006c,00760061006c007500650073,007a*/
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
		/*gopurs:record-v1:6:006100630074006900760065,00690064,006e0061006d0065,007400610069006c,00760061006c007500650073,007a*/
	}) struct {
		active bool
		id     int64
		name   string
		/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
	} {
		return struct {
			active bool
			id     int64
			name   string
			/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
		}{record.active, record.id, record.name}
	}(struct {
		active bool
		id     int64
		name   string
		tail   bool
		values []int64
		z      string
		/*gopurs:record-v1:6:006100630074006900760065,00690064,006e0061006d0065,007400610069006c,00760061006c007500650073,007a*/
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
			__t2 = gopurs_runtime.IntAdd(gopurs_runtime.IntMul(v_1_0.V1.id, int64(3)), __t1)
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
		/*gopurs:record-v1:1:00720065007400610069006e00650064*/
	}
	/*gopurs:record-v1:5:006100630074006900760065,006200650066006f00720065,00690064,006e0061006d0065,006e00650073007400650064*/
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
			/*gopurs:record-v1:1:00720065007400610069006e00650064*/
		}
		/*gopurs:record-v1:5:006100630074006900760065,006200650066006f00720065,00690064,006e0061006d0065,006e00650073007400650064*/
	}{false, 1.25, n_0, "beta", struct {
		retained string
		/*gopurs:record-v1:1:00720065007400610069006e00650064*/
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
			/*gopurs:record-v1:1:00720065007400610069006e00650064*/
		}
		/*gopurs:record-v1:5:006100630074006900760065,006200650066006f00720065,00690064,006e0061006d0065,006e00650073007400650064*/
	}) struct {
		active bool
		id     int64
		name   string
		/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
	} {
		return struct {
			active bool
			id     int64
			name   string
			/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
		}{record.active, record.id, record.name}
	}(struct {
		active bool
		before float64
		id     int64
		name   string
		nested struct {
			retained string
			/*gopurs:record-v1:1:00720065007400610069006e00650064*/
		}
		/*gopurs:record-v1:5:006100630074006900760065,006200650066006f00720065,00690064,006e0061006d0065,006e00650073007400650064*/
	}{false, 1.25, n_0, "beta", struct {
		retained string
		/*gopurs:record-v1:1:00720065007400610069006e00650064*/
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
			__t2 = gopurs_runtime.IntAdd(gopurs_runtime.IntMul(v_1_0.V1.id, int64(3)), __t1)
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
	/*gopurs:record-v1:4:006100630074006900760065,00650078007400720061,00690064,006e0061006d0065*/
} {
	var n_0 int64 = n_0_loop
	_ = n_0
	return struct {
		active bool
		extra  int64
		id     int64
		name   string
		/*gopurs:record-v1:4:006100630074006900760065,00650078007400720061,00690064,006e0061006d0065*/
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
		/*gopurs:record-v1:4:006100630074006900760065,00650078007400720061,00690064,006e0061006d0065*/
	}) struct {
		active bool
		id     int64
		name   string
		/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
	} {
		return struct {
			active bool
			id     int64
			name   string
			/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
		}{record.active, record.id, record.name}
	}(struct {
		active bool
		extra  int64
		id     int64
		name   string
		/*gopurs:record-v1:4:006100630074006900760065,00650078007400720061,00690064,006e0061006d0065*/
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
			__t2 = gopurs_runtime.IntAdd(gopurs_runtime.IntMul(v_1_0.V1.id, int64(3)), __t1)
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
	/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
} {
	var n_0 int64 = n_0_loop
	_ = n_0
	// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [String, (Record (Row [id: Int, name: String, active: Boolean] Empty))])
	v_1_0 := Call_Worker_decode(func(record struct {
		active bool
		extra  int64
		id     int64
		name   string
		/*gopurs:record-v1:4:006100630074006900760065,00650078007400720061,00690064,006e0061006d0065*/
	}) struct {
		active bool
		id     int64
		name   string
		/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
	} {
		return struct {
			active bool
			id     int64
			name   string
			/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
		}{record.active, record.id, record.name}
	}(struct {
		active bool
		extra  int64
		id     int64
		name   string
		/*gopurs:record-v1:4:006100630074006900760065,00650078007400720061,00690064,006e0061006d0065*/
	}{true, int64(42), n_0, "alpha"}))
	_ = v_1_0
	var __t1 struct {
		active bool
		id     int64
		name   string
		/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
	}
	{
		if !v_1_0.V2 {
			__t1 = struct {
				active bool
				id     int64
				name   string
				/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
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
			/*gopurs:record-v1:3:006100630074006900760065,00690064,006e0061006d0065*/
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

func Rebox_Main_1386611502_1469227923(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[[]int64])(unsafe.Pointer(in))
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

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[bool])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_920949869(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	retained string
	/*gopurs:record-v1:1:00720065007400610069006e00650064*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[struct {
		retained string
		/*gopurs:record-v1:1:00720065007400610069006e00650064*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_1469227923_1386611502(in *Constructor_Data_Show_Show[[]int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
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

func Rebox_Main_1830338253_3790796878(in *Constructor_Data_Eq_Eq[struct {
	retained string
	/*gopurs:record-v1:1:00720065007400610069006e00650064*/
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
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

func Rebox_Main_378698611_3790796878(in *Constructor_Data_Eq_Eq[[]int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
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

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[string])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1830338253(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	retained string
	/*gopurs:record-v1:1:00720065007400610069006e00650064*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[struct {
		retained string
		/*gopurs:record-v1:1:00720065007400610069006e00650064*/
	}])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[bool])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_378698611(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]int64])(unsafe.Pointer(in))
}

func Rebox_Main_920949869_1386611502(in *Constructor_Data_Show_Show[struct {
	retained string
	/*gopurs:record-v1:1:00720065007400610069006e00650064*/
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}
