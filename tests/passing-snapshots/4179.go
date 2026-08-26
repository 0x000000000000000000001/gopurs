package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_identity(x_0_box)
		})
	})
	return cache_Main_identity
}

var cache_Main_eqMaybe gopurs_runtime.Value
var once_Main_eqMaybe sync.Once

func Get_Main_eqMaybe() gopurs_runtime.Value {
	once_Main_eqMaybe.Do(func() {
		cache_Main_eqMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[string]]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 bool
				{
					if x_0.Type == 9 && x_0.IntVal == 930809136 && x_0.UnsafePtr == nil {
						__t0 = (y_1.Type == 9 && y_1.IntVal == 930809136 && y_1.UnsafePtr == nil)
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = (x_0.Type == 9 && x_0.IntVal == 930809136 && x_0.UnsafePtr != nil) && ((y_1.Type == 9 && y_1.IntVal == 930809136 && y_1.UnsafePtr != nil) && (((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_0.UnsafePtr).V0.StrVal()) == ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_1.UnsafePtr).V0.StrVal())))
				}
			end_branch_0:
				return gopurs_runtime.Bool(__t0)
			})
		})}))}
	})
	return cache_Main_eqMaybe
}

var cache_Main_runtimeImport gopurs_runtime.Value
var once_Main_runtimeImport sync.Once

func Get_Main_runtimeImport() gopurs_runtime.Value {
	once_Main_runtimeImport.Do(func() {
		cache_Main_runtimeImport = gopurs_runtime.Apply2(Get_Main_runtimeImportImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[string])(nil))}, Get_Data_Maybe_Just())
	})
	return cache_Main_runtimeImport
}

var cache_Main_runtimeImport__1627706496 gopurs_runtime.Value
var once_Main_runtimeImport__1627706496 sync.Once

func Get_Main_runtimeImport__1627706496() gopurs_runtime.Value {
	once_Main_runtimeImport__1627706496.Do(func() {
		cache_Main_runtimeImport__1627706496 = gopurs_runtime.Apply2(Get_Main_runtimeImportImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[string])(nil))}, Get_Data_Maybe_Just())
	})
	return cache_Main_runtimeImport__1627706496
}

var cache_Main_force gopurs_runtime.Value
var once_Main_force sync.Once

func Get_Main_force() gopurs_runtime.Value {
	once_Main_force.Do(func() {
		cache_Main_force = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_force(f_0_box)
		})
	})
	return cache_Main_force
}

var cache_Main_force__270261239 gopurs_runtime.Value
var once_Main_force__270261239 sync.Once

func Get_Main_force__270261239() gopurs_runtime.Value {
	once_Main_force__270261239.Do(func() {
		cache_Main_force__270261239 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_force__270261239(f_0_box)
		})
	})
	return cache_Main_force__270261239
}

var cache_Main_complicatedIdentity gopurs_runtime.Value
var once_Main_complicatedIdentity sync.Once

func Get_Main_complicatedIdentity() gopurs_runtime.Value {
	once_Main_complicatedIdentity.Do(func() {
		cache_Main_complicatedIdentity = func() gopurs_runtime.Value {
			var h_0_0_0 gopurs_runtime.Value
			_ = h_0_0_0
			var g_0_1_1 gopurs_runtime.Value
			_ = g_0_1_1
			var f_0_2_2 gopurs_runtime.Value
			_ = f_0_2_2
			h_0_0_0 = gopurs_runtime.RecordGet(gopurs_runtime.Apply(f_0_2_2, gopurs_runtime.Int(10)), "tick")
			g_0_1_1 = gopurs_runtime.Func(func(n_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordGet(gopurs_runtime.Apply(f_0_2_2, gopurs_runtime.Int(n_1.IntVal)), "tick")
			})
			f_0_2_2 = gopurs_runtime.Func(func(n_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 gopurs_runtime.Value
				{
					if (n_1.IntVal) <= (0) {
						__t3 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_2
						})
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(f_0_2_2, gopurs_runtime.Int((n_1.IntVal)-(1))), "tock"), Get_Control_Category_identity__193435443())
				}
			end_branch_3:
				return func() gopurs_runtime.Value {
					orig := func() *struct {
						tick gopurs_runtime.Value
						tock gopurs_runtime.Value
					} {
						orig := gopurs_runtime.RecordDict2("tick", "tock", __t3, gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(g_0_1_1, gopurs_runtime.Int(n_1.IntVal), a_2)
						}))
						_ = orig
						clone := struct {
							tick gopurs_runtime.Value
							tock gopurs_runtime.Value
						}{}
						clone.tick = gopurs_runtime.RecordGet(orig, "tick")
						clone.tock = gopurs_runtime.RecordGet(orig, "tock")
						return &clone
					}()
					_ = orig
					return gopurs_runtime.RecordDict([]string{"tick", "tock"}, []gopurs_runtime.Value{orig.tick, orig.tock})
				}()
			})
			return h_0_0_0
		}()
	})
	return cache_Main_complicatedIdentity
}

var cache_Main_bravo gopurs_runtime.Value
var once_Main_bravo sync.Once

func Get_Main_bravo() gopurs_runtime.Value {
	once_Main_bravo.Do(func() {
		cache_Main_bravo = gopurs_runtime.Int(gopurs_runtime.Int(func() *struct {
			backref gopurs_runtime.Value
			x       int64
		} {
			orig := Get_Main_alpha()
			_ = orig
			clone := struct {
				backref gopurs_runtime.Value
				x       int64
			}{}
			clone.backref = gopurs_runtime.RecordGet(orig, "backref")
			clone.x = gopurs_runtime.RecordGet(orig, "x").IntVal
			return &clone
		}().x).IntVal)
	})
	return cache_Main_bravo
}

var cache_Main_alpha gopurs_runtime.Value
var once_Main_alpha sync.Once

func Get_Main_alpha() gopurs_runtime.Value {
	once_Main_alpha.Do(func() {
		cache_Main_alpha = func() gopurs_runtime.Value {
			orig := func() *struct {
				backref gopurs_runtime.Value
				x       int64
			} {
				orig := gopurs_runtime.RecordDict2("backref", "x", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(gopurs_runtime.Int(func() *struct {
						backref gopurs_runtime.Value
						x       int64
					} {
						orig := Get_Main_alpha()
						_ = orig
						clone := struct {
							backref gopurs_runtime.Value
							x       int64
						}{}
						clone.backref = gopurs_runtime.RecordGet(orig, "backref")
						clone.x = gopurs_runtime.RecordGet(orig, "x").IntVal
						return &clone
					}().x).IntVal)
				}), gopurs_runtime.Int(1))
				_ = orig
				clone := struct {
					backref gopurs_runtime.Value
					x       int64
				}{}
				clone.backref = gopurs_runtime.RecordGet(orig, "backref")
				clone.x = gopurs_runtime.RecordGet(orig, "x").IntVal
				return &clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict([]string{"backref", "x"}, []gopurs_runtime.Value{orig.backref, gopurs_runtime.Int(orig.x)})
		}()
	})
	return cache_Main_alpha
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_CustomAssert_assertThrows(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				var selfOwn_1_1_3 *struct {
					a int64
					b int64
				}
				_ = selfOwn_1_1_3
				selfOwn_1_1_3 = func() *struct {
					a int64
					b int64
				} {
					orig := gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Int(selfOwn_1_1_3.a).IntVal))
					_ = orig
					clone := struct {
						a int64
						b int64
					}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a").IntVal
					clone.b = gopurs_runtime.RecordGet(orig, "b").IntVal
					return &clone
				}()
				return func() gopurs_runtime.Value {
					orig := selfOwn_1_1_3
					_ = orig
					return gopurs_runtime.RecordDict([]string{"a", "b"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.a), gopurs_runtime.Int(orig.b)})
				}()
			}))
			_ = __local_var_0_0
			err_1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = err_1_2
			_dollar___unused_2_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(((gopurs_runtime.Apply2(Get_Data_String_CodeUnits_contains(), gopurs_runtime.Str("interface conversion"), gopurs_runtime.Str(err_1_2.StrVal())).IntVal) != (0)) || (((gopurs_runtime.Apply2(Get_Data_String_CodeUnits_contains(), gopurs_runtime.Str("Attempt to read property"), gopurs_runtime.Str(err_1_2.StrVal())).IntVal) != (0)) || ((gopurs_runtime.Apply2(Get_Data_String_CodeUnits_contains(), gopurs_runtime.Str("nil pointer"), gopurs_runtime.Str(err_1_2.StrVal())).IntVal) != (0))))), gopurs_runtime.Value{})
			_ = _dollar___unused_2_3
			err2_3_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_CustomAssert_assertThrows(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var j_4_5_4 gopurs_runtime.Value
				_ = j_4_5_4
				var h_4_6_5 gopurs_runtime.Value
				_ = h_4_6_5
				var g_4_7_6 gopurs_runtime.Value
				_ = g_4_7_6
				var f_4_8_7 *struct {
					left  gopurs_runtime.Value
					right gopurs_runtime.Value
				}
				_ = f_4_8_7
				j_4_5_4 = gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(z_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return func() gopurs_runtime.Value {
								orig := func() *struct {
									left  gopurs_runtime.Value
									right gopurs_runtime.Value
								} {
									orig := gopurs_runtime.RecordDict2("left", "right", gopurs_runtime.Apply2(x_5, y_6, z_7), f_4_8_7.left)
									_ = orig
									clone := struct {
										left  gopurs_runtime.Value
										right gopurs_runtime.Value
									}{}
									clone.left = gopurs_runtime.RecordGet(orig, "left")
									clone.right = gopurs_runtime.RecordGet(orig, "right")
									return &clone
								}()
								_ = orig
								return gopurs_runtime.RecordDict([]string{"left", "right"}, []gopurs_runtime.Value{orig.left, orig.right})
							}()
						})
					})
				})
				h_4_6_5 = gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(j_4_5_4, x_5, x_5)
				})
				g_4_7_6 = gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(j_4_5_4, x_5, x_5, x_5), "right")
				})
				f_4_8_7 = func() *struct {
					left  gopurs_runtime.Value
					right gopurs_runtime.Value
				} {
					orig := gopurs_runtime.RecordDict2("left", "right", gopurs_runtime.Apply(g_4_7_6, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return x_5
					})), gopurs_runtime.Apply(h_4_6_5, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return x_5
					})))
					_ = orig
					clone := struct {
						left  gopurs_runtime.Value
						right gopurs_runtime.Value
					}{}
					clone.left = gopurs_runtime.RecordGet(orig, "left")
					clone.right = gopurs_runtime.RecordGet(orig, "right")
					return &clone
				}()
				return func() gopurs_runtime.Value {
					orig := f_4_8_7
					_ = orig
					return gopurs_runtime.RecordDict([]string{"left", "right"}, []gopurs_runtime.Value{orig.left, orig.right})
				}()
			})), gopurs_runtime.Value{})
			_ = err2_3_4
			_dollar___unused_4_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(((gopurs_runtime.Apply2(Get_Data_String_CodeUnits_contains(), gopurs_runtime.Str("interface conversion"), gopurs_runtime.Str(err2_3_4.StrVal())).IntVal) != (0)) || (((gopurs_runtime.Apply2(Get_Data_String_CodeUnits_contains(), gopurs_runtime.Str("Attempt to read property"), gopurs_runtime.Str(err2_3_4.StrVal())).IntVal) != (0)) || ((gopurs_runtime.Apply2(Get_Data_String_CodeUnits_contains(), gopurs_runtime.Str("nil pointer"), gopurs_runtime.Str(err2_3_4.StrVal())).IntVal) != (0))))), gopurs_runtime.Value{})
			_ = _dollar___unused_4_9
			// TAST (Let): __local_var_5_10 shape=LitRecord bindingType=(Record (Row [actual: Int, expected: Int] Any))
			__local_var_5_10 := func() *struct {
				actual   int64
				expected int64
			} {
				orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal), gopurs_runtime.Int(1))
				_ = orig
				clone := struct {
					actual   int64
					expected int64
				}{}
				clone.actual = gopurs_runtime.RecordGet(orig, "actual").IntVal
				clone.expected = gopurs_runtime.RecordGet(orig, "expected").IntVal
				return &clone
			}()
			_ = __local_var_5_10
			_dollar___unused_6_12 := Get_Data_Unit_unit()
			_ = _dollar___unused_6_12
			_dollar___unused_6_11 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(((("Expected: ")+(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(__local_var_5_10.expected)).StrVal()))+("\x0aActual:   "))+(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(__local_var_5_10.actual)).StrVal())), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_6_11
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Main_runtimeImport(), gopurs_runtime.Str("InitializationError"), gopurs_runtime.Func(func(err3_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					var __t_tag_14 *Constructor_Data_Maybe_Just[string] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](err3_7)
					// TAST (Let): __local_var_8_13 shape=LitRecord bindingType=(Record (Row [actual: Boolean, expected: Boolean] Any))
					__local_var_8_13 := func() *struct {
						actual   bool
						expected bool
					} {
						orig := gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool((__t_tag_14 == nil) != (true)), gopurs_runtime.Bool(true))
						_ = orig
						clone := struct {
							actual   bool
							expected bool
						}{}
						clone.actual = (gopurs_runtime.RecordGet(orig, "actual").IntVal) != (0)
						clone.expected = (gopurs_runtime.RecordGet(orig, "expected").IntVal) != (0)
						return &clone
					}()
					_ = __local_var_8_13
					// TAST (Let): result_9_15 shape=Other bindingType=Boolean
					result_9_15 := __local_var_8_13.actual
					_ = result_9_15
					var __t17 string
					{
						if (gopurs_runtime.Bool(__local_var_8_13.actual).IntVal) != (0) {
							__t17 = ("Expected: true\x0aActual:   ") + ("true")
							goto end_branch_17
						} else {

						}
					}
					{
						__t17 = ("Expected: true\x0aActual:   ") + ("false")
					}
				end_branch_17:
					// TAST (Let): message_10_16 shape=Branch(Other, def=Other) bindingType=String
					message_10_16 := __t17
					_ = message_10_16
					// TAST (Let): __local_var_11_18 shape=Let(Let(EffectBind(App(Var)))) bindingType=(ADT ["Effect","Effect"] [Unit])
					__local_var_11_18 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						// TAST (Let): __local_var_11_19 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
						__local_var_11_19 := gopurs_runtime.Apply(Get_Effect_Console_error(), gopurs_runtime.Str(message_10_16))
						_ = __local_var_11_19
						var __t21 gopurs_runtime.Value
						{
							if (result_9_15) != (true) {
								__t21 = __local_var_11_19
								goto end_branch_21
							} else {

							}
						}
						{
							if result_9_15 {
								__t21 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									return Get_Data_Unit_unit()
								})
								goto end_branch_21
							} else {

							}
						}
						{
							__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_21:
						// TAST (Let): __local_var_12_20 shape=Branch(Other, EffectPure, def=Other) bindingType=(TypeApp (TypeVar m) [Unit])
						__local_var_12_20 := __t21
						_ = __local_var_12_20
						_dollar___unused_13_22 := gopurs_runtime.Apply(__local_var_12_20, gopurs_runtime.Value{})
						_ = _dollar___unused_13_22
						return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(message_10_16), gopurs_runtime.Bool(result_9_15)), gopurs_runtime.Value{})
					})
					_ = __local_var_11_18
					_dollar___unused_12_23 := gopurs_runtime.Apply(__local_var_11_18, gopurs_runtime.Value{})
					_ = _dollar___unused_12_23
					return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
				})
			})), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_force(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
}

func Call_Main_force__270261239(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
}

func Get_Main_runtimeImportImpl() gopurs_runtime.Value {
	return _Gopurs_Main_RuntimeImportImpl
}
