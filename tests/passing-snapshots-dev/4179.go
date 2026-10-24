package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_runtimeImport gopurs_runtime.Value
var once_Main_runtimeImport sync.Once

func Get_Main_runtimeImport() gopurs_runtime.Value {
	once_Main_runtimeImport.Do(func() {
		cache_Main_runtimeImport = gopurs_runtime.Apply2(Get_Main_runtimeImportImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_742090555_3094389156(Rebox_Main_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))))}, Get_Data_Maybe_Just())
	})
	return cache_Main_runtimeImport
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

var cache_Main_complicatedIdentity gopurs_runtime.Value
var once_Main_complicatedIdentity sync.Once

func Get_Main_complicatedIdentity() gopurs_runtime.Value {
	once_Main_complicatedIdentity.Do(func() {
		cache_Main_complicatedIdentity = func() gopurs_runtime.Value {
			var h_0_0_0 gopurs_runtime.Value
			_ = h_0_0_0
			var h_0_0_0_cell *gopurs_runtime.Value
			_ = h_0_0_0_cell
			// FALLBACK TCO: isLoop=false len=3
			var g_0_1_1 gopurs_runtime.Value
			_ = g_0_1_1
			var g_0_1_1_cell *gopurs_runtime.Value
			_ = g_0_1_1_cell
			// FALLBACK TCO: isLoop=false len=3
			var f_0_2_2 gopurs_runtime.Value
			_ = f_0_2_2
			var f_0_2_2_cell *gopurs_runtime.Value
			_ = f_0_2_2_cell
			// FALLBACK TCO: isLoop=false len=3
			h_0_0_0 = gopurs_runtime.RecordGet(gopurs_runtime.Apply((*f_0_2_2_cell), gopurs_runtime.Int(int64(10))), "tick")
			h_0_0_0_cell = &h_0_0_0
			g_0_1_1 = gopurs_runtime.Func(func(n_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordGet(gopurs_runtime.Apply((*f_0_2_2_cell), gopurs_runtime.Int(n_1.IntVal)), "tick")
			})
			g_0_1_1_cell = &g_0_1_1
			f_0_2_2 = gopurs_runtime.Func(func(n_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t3 gopurs_runtime.Value
				{
					if (n_1.IntVal) <= (int64(0)) {
						__t3 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_2
						})
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply((*f_0_2_2_cell), gopurs_runtime.Int(gopurs_runtime.IntSub(n_1.IntVal, int64(1)))), "tock"), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
				}
			end_branch_3:
				return gopurs_runtime.RecordDict2("tick", "tock", __t3, gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2((*g_0_1_1_cell), gopurs_runtime.Int(n_1.IntVal), a_2)
				}))
			})
			f_0_2_2_cell = &f_0_2_2
			return h_0_0_0
		}()
	})
	return cache_Main_complicatedIdentity
}

var cache_Main_bravo gopurs_runtime.Value
var once_Main_bravo sync.Once

func Get_Main_bravo() gopurs_runtime.Value {
	once_Main_bravo.Do(func() {
		cache_Main_bravo = gopurs_runtime.Int(Call_Main_force(gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(func() struct {
				backref gopurs_runtime.Value
				x       int64
				/*gopurs:record-v1:2:006200610063006b007200650066,0078*/
			} {
				orig := Get_Main_alpha()
				_ = orig
				clone := struct {
					backref gopurs_runtime.Value
					x       int64
					/*gopurs:record-v1:2:006200610063006b007200650066,0078*/
				}{}
				clone.backref = gopurs_runtime.RecordGet(orig, "backref")
				clone.x = gopurs_runtime.RecordGet(orig, "x").IntVal
				return clone
			}().x)
		})).IntVal)
	})
	return cache_Main_bravo
}

var cache_Main_alpha gopurs_runtime.Value
var once_Main_alpha sync.Once

func Get_Main_alpha() gopurs_runtime.Value {
	once_Main_alpha.Do(func() {
		cache_Main_alpha = func() gopurs_runtime.Value {
			orig := struct {
				backref gopurs_runtime.Value
				x       int64
				/*gopurs:record-v1:2:006200610063006b007200650066,0078*/
			}{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(Get_Main_bravo().IntVal)
			}), int64(1)}
			_ = orig
			return gopurs_runtime.RecordDict2("backref", "x", orig.backref, gopurs_runtime.Int(orig.x))
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
				var selfOwn_1_1_3 struct {
					a int64
					b int64
					/*gopurs:record-v1:2:0061,0062*/
				}
				_ = selfOwn_1_1_3
				var selfOwn_1_1_3_cell *struct {
					a int64
					b int64
					/*gopurs:record-v1:2:0061,0062*/
				}
				_ = selfOwn_1_1_3_cell
				// FALLBACK TCO: isLoop=false len=1
				selfOwn_1_1_3 = struct {
					a int64
					b int64
					/*gopurs:record-v1:2:0061,0062*/
				}{int64(1), Call_Main_force(gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int((*selfOwn_1_1_3_cell).a)
				})).IntVal}
				selfOwn_1_1_3_cell = &selfOwn_1_1_3
				return func() gopurs_runtime.Value {
					orig := selfOwn_1_1_3
					_ = orig
					return gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Int(orig.a), gopurs_runtime.Int(orig.b))
				}()
			}))
			_ = __local_var_0_0
			err_1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = err_1_2
			_dollar___unused_2_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(((gopurs_runtime.Apply(Call_Data_String_CodeUnits_contains("interface conversion"), gopurs_runtime.Str(err_1_2.StrVal())).IntVal) != (0)) || (((gopurs_runtime.Apply(Call_Data_String_CodeUnits_contains("Attempt to read property"), gopurs_runtime.Str(err_1_2.StrVal())).IntVal) != (0)) || ((gopurs_runtime.Apply(Call_Data_String_CodeUnits_contains("nil pointer"), gopurs_runtime.Str(err_1_2.StrVal())).IntVal) != (0))))), gopurs_runtime.Value{})
			_ = _dollar___unused_2_3
			err2_3_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_CustomAssert_assertThrows(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var j_4_5_4 gopurs_runtime.Value
				_ = j_4_5_4
				var j_4_5_4_cell *gopurs_runtime.Value
				_ = j_4_5_4_cell
				// FALLBACK TCO: isLoop=false len=4
				var h_4_6_5 gopurs_runtime.Value
				_ = h_4_6_5
				var h_4_6_5_cell *gopurs_runtime.Value
				_ = h_4_6_5_cell
				// FALLBACK TCO: isLoop=false len=4
				var g_4_7_6 gopurs_runtime.Value
				_ = g_4_7_6
				var g_4_7_6_cell *gopurs_runtime.Value
				_ = g_4_7_6_cell
				// FALLBACK TCO: isLoop=false len=4
				var f_4_8_7 struct {
					left  gopurs_runtime.Value
					right gopurs_runtime.Value
					/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
				}
				_ = f_4_8_7
				var f_4_8_7_cell *struct {
					left  gopurs_runtime.Value
					right gopurs_runtime.Value
					/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
				}
				_ = f_4_8_7_cell
				// FALLBACK TCO: isLoop=false len=4
				j_4_5_4 = gopurs_runtime.Func3(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value, z_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.RecordDict2("left", "right", gopurs_runtime.Apply2(x_5, y_6, z_7), (*f_4_8_7_cell).left)
				})
				j_4_5_4_cell = &j_4_5_4
				h_4_6_5 = gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2((*j_4_5_4_cell), x_5, x_5)
				})
				h_4_6_5_cell = &h_4_6_5
				g_4_7_6 = gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.RecordGet(gopurs_runtime.Apply3((*j_4_5_4_cell), x_5, x_5, x_5), "right")
				})
				g_4_7_6_cell = &g_4_7_6
				f_4_8_7 = struct {
					left  gopurs_runtime.Value
					right gopurs_runtime.Value
					/*gopurs:record-v1:2:006c006500660074,00720069006700680074*/
				}{gopurs_runtime.Apply((*g_4_7_6_cell), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_5
				})), gopurs_runtime.Apply((*h_4_6_5_cell), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_5
				}))}
				f_4_8_7_cell = &f_4_8_7
				return func() gopurs_runtime.Value {
					orig := f_4_8_7
					_ = orig
					return gopurs_runtime.RecordDict2("left", "right", orig.left, orig.right)
				}()
			})), gopurs_runtime.Value{})
			_ = err2_3_4
			_dollar___unused_4_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(((gopurs_runtime.Apply(Call_Data_String_CodeUnits_contains("interface conversion"), gopurs_runtime.Str(err2_3_4.StrVal())).IntVal) != (0)) || (((gopurs_runtime.Apply(Call_Data_String_CodeUnits_contains("Attempt to read property"), gopurs_runtime.Str(err2_3_4.StrVal())).IntVal) != (0)) || ((gopurs_runtime.Apply(Call_Data_String_CodeUnits_contains("nil pointer"), gopurs_runtime.Str(err2_3_4.StrVal())).IntVal) != (0))))), gopurs_runtime.Value{})
			_ = _dollar___unused_4_9
			_dollar___unused_5_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))), Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())))), func() gopurs_runtime.Value {
				orig := struct {
					actual   int64
					expected int64
					/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
				}{Get_Main_bravo().IntVal, int64(1)}
				_ = orig
				return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(orig.actual), gopurs_runtime.Int(orig.expected))
			}()), gopurs_runtime.Value{})
			_ = _dollar___unused_5_10
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Main_runtimeImport(), gopurs_runtime.Str("InitializationError"), gopurs_runtime.Func(func(err3_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_7_11 shape=App(Var) bindingType=Any
					__local_var_7_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))), Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean())))), func() gopurs_runtime.Value {
						orig := struct {
							actual   bool
							expected bool
							/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
						}{((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Maybe_eqMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_742090555_3094389156(Rebox_Main_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](err3_6))))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_742090555_3094389156(Rebox_Main_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 bool
							}{gopurs_runtime.Value{}, false}
							if _v.V1 {
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
							}
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
						}()))))}).IntVal) != (0)) != (true), true}
						_ = orig
						return gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool(orig.actual), gopurs_runtime.Bool(orig.expected))
					}())
					_ = __local_var_7_11
					_dollar___unused_8_12 := gopurs_runtime.Apply(__local_var_7_11, gopurs_runtime.Value{})
					_ = _dollar___unused_8_12
					return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
				})
			})), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_force(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
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

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
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

func Rebox_Main_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
	return out
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

func Rebox_Main_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[bool])(unsafe.Pointer(in))
}

func Rebox_Main_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}

func Get_Main_runtimeImportImpl() gopurs_runtime.Value {
	return _Gopurs_Main_RuntimeImportImpl
}
