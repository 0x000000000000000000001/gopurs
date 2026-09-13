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
		cache_Main_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_identity
}

var cache_Main_identity1 gopurs_runtime.Value
var once_Main_identity1 sync.Once

func Get_Main_identity1() gopurs_runtime.Value {
	once_Main_identity1.Do(func() {
		cache_Main_identity1 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_identity1
}

var cache_Main_identity2 gopurs_runtime.Value
var once_Main_identity2 sync.Once

func Get_Main_identity2() gopurs_runtime.Value {
	once_Main_identity2.Do(func() {
		cache_Main_identity2 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Main_identity2
}

var cache_Main_eqMaybe gopurs_runtime.Value
var once_Main_eqMaybe sync.Once

func Get_Main_eqMaybe() gopurs_runtime.Value {
	once_Main_eqMaybe.Do(func() {
		cache_Main_eqMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1144047024_3790796878(Rebox_Main_3790796878_1144047024(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Maybe_eqMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))}
	})
	return cache_Main_eqMaybe
}

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

var cache_Main_runtimeImport__2128991624 gopurs_runtime.Value
var once_Main_runtimeImport__2128991624 sync.Once

func Get_Main_runtimeImport__2128991624() gopurs_runtime.Value {
	once_Main_runtimeImport__2128991624.Do(func() {
		cache_Main_runtimeImport__2128991624 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runtimeImport__2128991624(__eta_norm_1_0_box.StrVal(), __eta_norm_0_1_box)
		})
	})
	return cache_Main_runtimeImport__2128991624
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

var cache_Main_force__3600449186 gopurs_runtime.Value
var once_Main_force__3600449186 sync.Once

func Get_Main_force__3600449186() gopurs_runtime.Value {
	once_Main_force__3600449186.Do(func() {
		cache_Main_force__3600449186 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_force__3600449186(f_0_box))
		})
	})
	return cache_Main_force__3600449186
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
						__t3 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
						goto end_branch_3
					} else {

					}
				}
				{
					__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply((*f_0_2_2_cell), gopurs_runtime.Int((n_1.IntVal)-(int64(1)))), "tock"), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
				}
			end_branch_3:
				return func() gopurs_runtime.Value {
					orig := struct {
						tick gopurs_runtime.Value
						tock gopurs_runtime.Value
					}{__t3, gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2((*g_0_1_1_cell), gopurs_runtime.Int(n_1.IntVal), a_2)
					})}
					_ = orig
					return gopurs_runtime.RecordDict2("tick", "tock", orig.tick, orig.tock)
				}()
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
		cache_Main_bravo = gopurs_runtime.Int(gopurs_runtime.Int(func() struct {
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
			return clone
		}().x).IntVal)
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
			__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_CustomAssert_assertThrowsImpl(), Get_Data_Unit_unit(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				var selfOwn_1_1_3 struct {
					a int64
					b int64
				}
				_ = selfOwn_1_1_3
				var selfOwn_1_1_3_cell *struct {
					a int64
					b int64
				}
				_ = selfOwn_1_1_3_cell
				// FALLBACK TCO: isLoop=false len=1
				selfOwn_1_1_3 = struct {
					a int64
					b int64
				}{int64(1), gopurs_runtime.Int((*selfOwn_1_1_3_cell).a).IntVal}
				selfOwn_1_1_3_cell = &selfOwn_1_1_3
				return func() gopurs_runtime.Value {
					orig := selfOwn_1_1_3
					_ = orig
					return gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Int(orig.a), gopurs_runtime.Int(orig.b))
				}()
			})), gopurs_runtime.Value{})
			_ = __local_var_0_0
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(((gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.Str("interface conversion")), gopurs_runtime.Str(__local_var_0_0.StrVal())).IntVal) != (0)) || (((gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.Str("Attempt to read property")), gopurs_runtime.Str(__local_var_0_0.StrVal())).IntVal) != (0)) || ((gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.Str("nil pointer")), gopurs_runtime.Str(__local_var_0_0.StrVal())).IntVal) != (0))))), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_2_2 shape=App(Var) bindingType=Any
					__local_var_2_2 := gopurs_runtime.Apply(Get_CustomAssert_assertThrows(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
						var j_3_3_4 gopurs_runtime.Value
						_ = j_3_3_4
						var j_3_3_4_cell *gopurs_runtime.Value
						_ = j_3_3_4_cell
						// FALLBACK TCO: isLoop=false len=4
						var h_3_4_5 gopurs_runtime.Value
						_ = h_3_4_5
						var h_3_4_5_cell *gopurs_runtime.Value
						_ = h_3_4_5_cell
						// FALLBACK TCO: isLoop=false len=4
						var g_3_5_6 gopurs_runtime.Value
						_ = g_3_5_6
						var g_3_5_6_cell *gopurs_runtime.Value
						_ = g_3_5_6_cell
						// FALLBACK TCO: isLoop=false len=4
						var f_3_6_7 struct {
							left  gopurs_runtime.Value
							right gopurs_runtime.Value
						}
						_ = f_3_6_7
						var f_3_6_7_cell *struct {
							left  gopurs_runtime.Value
							right gopurs_runtime.Value
						}
						_ = f_3_6_7_cell
						// FALLBACK TCO: isLoop=false len=4
						j_3_3_4 = gopurs_runtime.Func3(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value, z_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return func() gopurs_runtime.Value {
								orig := struct {
									left  gopurs_runtime.Value
									right gopurs_runtime.Value
								}{gopurs_runtime.Apply2(x_4, y_5, z_6), (*f_3_6_7_cell).left}
								_ = orig
								return gopurs_runtime.RecordDict2("left", "right", orig.left, orig.right)
							}()
						})
						j_3_3_4_cell = &j_3_3_4
						h_3_4_5 = gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2((*j_3_3_4_cell), x_4, x_4)
						})
						h_3_4_5_cell = &h_3_4_5
						g_3_5_6 = gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.RecordGet(gopurs_runtime.Apply3((*j_3_3_4_cell), x_4, x_4, x_4), "right")
						})
						g_3_5_6_cell = &g_3_5_6
						f_3_6_7 = struct {
							left  gopurs_runtime.Value
							right gopurs_runtime.Value
						}{gopurs_runtime.Apply((*g_3_5_6_cell), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})), gopurs_runtime.Apply((*h_3_4_5_cell), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))}
						f_3_6_7_cell = &f_3_6_7
						return func() gopurs_runtime.Value {
							orig := f_3_6_7
							_ = orig
							return gopurs_runtime.RecordDict2("left", "right", orig.left, orig.right)
						}()
					}))
					_ = __local_var_2_2
					__local_var_3_7 := gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Value{})
					_ = __local_var_3_7
					return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(((gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.Str("interface conversion")), gopurs_runtime.Str(__local_var_3_7.StrVal())).IntVal) != (0)) || (((gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.Str("Attempt to read property")), gopurs_runtime.Str(__local_var_3_7.StrVal())).IntVal) != (0)) || ((gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isJust(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.Str("nil pointer")), gopurs_runtime.Str(__local_var_3_7.StrVal())).IntVal) != (0))))), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
							actual   int64
							expected int64
						}{gopurs_runtime.Int(int64(1)).IntVal, int64(1)}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply4(Get_Main_runtimeImportImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_742090555_3094389156(Rebox_Main_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
								_v := struct {
									V0 gopurs_runtime.Value
									V1 bool
								}{gopurs_runtime.Value{}, false}
								if _v.V1 {
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
								}
								return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
							}()))))}, Get_Data_Maybe_Just(), gopurs_runtime.Str("InitializationError"), gopurs_runtime.Func(func(err3_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___1138829510("", struct {
									actual   bool
									expected bool
								}{((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Maybe_eqMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), err3_6, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_742090555_3094389156(Rebox_Main_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
									_v := struct {
										V0 gopurs_runtime.Value
										V1 bool
									}{gopurs_runtime.Value{}, false}
									if _v.V1 {
										return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
									}
									return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
								}()))))}).IntVal) != (0)) != (true), true}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
								}))
							}))
						}))
					})), gopurs_runtime.Value{})
				})
			})), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_runtimeImport__2128991624(__eta_norm_1_0_loop string, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
runtimeImport__2128991624:
	for {
		if false {
			continue runtimeImport__2128991624
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply4(Get_Main_runtimeImportImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Main_742090555_3094389156(Rebox_Main_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))))}, Get_Data_Maybe_Just(), gopurs_runtime.Str(__eta_norm_1_0), __eta_norm_0_1)
	}
}

func Call_Main_force(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
}

func Call_Main_force__3600449186(f_0_loop gopurs_runtime.Value) int64 {
force__3600449186:
	for {
		if false {
			continue force__3600449186
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		return gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()).IntVal
	}
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1144047024_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[string]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
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

func Rebox_Main_3790796878_1144047024(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[string]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[string]]{}
	out.V0 = in.V0
	return out
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
