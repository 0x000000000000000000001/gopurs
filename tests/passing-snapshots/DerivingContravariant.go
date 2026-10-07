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
			return gopurs_runtime.Array(Call_Main_identity((*(*[]gopurs_runtime.Value)((x_0_box).UnsafePtr))))
		})
	})
	return cache_Main_identity
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqArray
}

var cache_Main_eqArray1 gopurs_runtime.Value
var once_Main_eqArray1 sync.Once

func Get_Main_eqArray1() gopurs_runtime.Value {
	once_Main_eqArray1.Do(func() {
		cache_Main_eqArray1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_131790935_3790796878(Rebox_Main_3790796878_131790935(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))))}
	})
	return cache_Main_eqArray1
}

var cache_Main_Test0 gopurs_runtime.Value
var once_Main_Test0 sync.Once

func Get_Main_Test0() gopurs_runtime.Value {
	once_Main_Test0.Do(func() {
		cache_Main_Test0 = gopurs_runtime.Value{Type: 9, IntVal: 2074462008, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Main_Test0
}

var cache_Main_Test1 gopurs_runtime.Value
var once_Main_Test1 sync.Once

func Get_Main_Test1() gopurs_runtime.Value {
	once_Main_Test1.Do(func() {
		cache_Main_Test1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Test1
}

var cache_Main_Test2 gopurs_runtime.Value
var once_Main_Test2 sync.Once

func Get_Main_Test2() gopurs_runtime.Value {
	once_Main_Test2.Do(func() {
		cache_Main_Test2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Test2
}

var cache_Main_Test3 gopurs_runtime.Value
var once_Main_Test3 sync.Once

func Get_Main_Test3() gopurs_runtime.Value {
	once_Main_Test3.Do(func() {
		cache_Main_Test3 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal, value1}))}
			})
		})
	})
	return cache_Main_Test3
}

var cache_Main_Test4 gopurs_runtime.Value
var once_Main_Test4 sync.Once

func Get_Main_Test4() gopurs_runtime.Value {
	once_Main_Test4.Do(func() {
		cache_Main_Test4 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3712677948, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal, value1}))}
			})
		})
	})
	return cache_Main_Test4
}

var cache_Main_Test5 gopurs_runtime.Value
var once_Main_Test5 sync.Once

func Get_Main_Test5() gopurs_runtime.Value {
	once_Main_Test5.Do(func() {
		cache_Main_Test5 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1063363133, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*(*[]gopurs_runtime.Value)((value0).UnsafePtr)), Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](value1))}))}
			})
		})
	})
	return cache_Main_Test5
}

var cache_Main_Test6 gopurs_runtime.Value
var once_Main_Test6 sync.Once

func Get_Main_Test6() gopurs_runtime.Value {
	once_Main_Test6.Do(func() {
		cache_Main_Test6 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4013407934, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test6[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() struct {
				nested []struct {
					x gopurs_runtime.Value
				}
			} {
				orig := value0
				_ = orig
				clone := struct {
					nested []struct {
						x gopurs_runtime.Value
					}
				}{}
				clone.nested = func() []struct {
					x gopurs_runtime.Value
				} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "nested").UnsafePtr)
					unboxed := make([]struct {
						x gopurs_runtime.Value
					}, len(arr))
					for i, v := range arr {
						unboxed[i] = func() struct {
							x gopurs_runtime.Value
						} {
							orig := v
							_ = orig
							clone := struct {
								x gopurs_runtime.Value
							}{}
							clone.x = gopurs_runtime.RecordGet(orig, "x")
							return clone
						}()
					}
					return unboxed
				}()
				return clone
			}()}))}
		})
	})
	return cache_Main_Test6
}

var cache_Main_contravariantTest gopurs_runtime.Value
var once_Main_contravariantTest sync.Once

func Get_Main_contravariantTest() gopurs_runtime.Value {
	once_Main_contravariantTest.Do(func() {
		cache_Main_contravariantTest = gopurs_runtime.Func(func(dictContravariant_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_contravariantTest(dictContravariant_0_box)
		})
	})
	return cache_Main_contravariantTest
}

var cache_Main_adapt gopurs_runtime.Value
var once_Main_adapt sync.Once

func Get_Main_adapt() gopurs_runtime.Value {
	once_Main_adapt.Do(func() {
		cache_Main_adapt = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Main_contravariantTest(gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant[gopurs_runtime.Value]](Get_Data_Predicate_contravariantPredicate()))}), "cmap"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(gopurs_runtime.RecordGet(v_0, "value").IntVal)
		}))
	})
	return cache_Main_adapt
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(7)))
			_ = __local_var_0_0
			__local_var_1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_2
			__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_2), gopurs_runtime.Value{})
			_ = __local_var_1_1
			var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply(Get_Main_adapt(), gopurs_runtime.Value{Type: 9, IntVal: 2074462008, UnsafePtr: unsafe.Pointer(Rebox_Main_1204441265_2762907754(nil))})
			_ = __t_tag_4
			__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("cmap - empty constructor"), gopurs_runtime.Bool((__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2074462008))), gopurs_runtime.Value{})
			_ = __local_var_2_3
			// TAST (Let): v_3_5 shape=App(Var) bindingType=(ADT ["Main","Test"] [(Func [(TypeVar a)] Boolean), (Record (Row [value: Int] Empty))])
			v_3_5 := gopurs_runtime.Apply(Get_Main_adapt(), gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer(Rebox_Main_3008504208_1610699019((&Constructor_Main_Test1[gopurs_runtime.Value, int64]{1, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((v1_3.IntVal) == (__local_var_1_1.IntVal))
			})})))})
			_ = v_3_5
			__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("cmap - predicate"), gopurs_runtime.Bool((v_3_5.Type == 9 && v_3_5.IntVal == 3720114489) && (((gopurs_runtime.Apply((*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_5.UnsafePtr).V0, func() gopurs_runtime.Value {
				orig := struct {
					value int64
				}{__local_var_1_1.IntVal}
				_ = orig
				return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
			}()).IntVal) != (0)) && (((gopurs_runtime.Apply((*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_5.UnsafePtr).V0, func() gopurs_runtime.Value {
				orig := struct {
					value int64
				}{int64(0)}
				_ = orig
				return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
			}()).IntVal) != (0)) != (true))))), gopurs_runtime.Value{})
			_ = __local_var_4_6
			// TAST (Let): v_5_7 shape=App(Var) bindingType=(ADT ["Main","Test"] [(Func [(TypeVar a)] Boolean), (Record (Row [value: Int] Empty))])
			v_5_7 := gopurs_runtime.Apply(Get_Main_adapt(), gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer(Rebox_Main_2838221171_1161568424((&Constructor_Main_Test2[gopurs_runtime.Value, int64]{1, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((gopurs_runtime.Apply(v1_5, gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((v2_6.IntVal) == (__local_var_1_1.IntVal))
				})).IntVal) != (0))
			})})))})
			_ = v_5_7
			__local_var_6_8 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("cmap - triple contravariance"), gopurs_runtime.Bool((v_5_7.Type == 9 && v_5_7.IntVal == 2375191994) && ((gopurs_runtime.Apply((*Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_7.UnsafePtr).V0, gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((gopurs_runtime.Apply(v1_6, func() gopurs_runtime.Value {
					orig := struct {
						value int64
					}{__local_var_1_1.IntVal}
					_ = orig
					return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
				}()).IntVal) != (0))
			})).IntVal) != (0)))), gopurs_runtime.Value{})
			_ = __local_var_6_8
			// TAST (Let): v_7_9 shape=App(Var) bindingType=(ADT ["Main","Test"] [(Func [(TypeVar a)] Boolean), (Record (Row [value: Int] Empty))])
			v_7_9 := gopurs_runtime.Apply(Get_Main_adapt(), gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer(Rebox_Main_590733650_648073545((&Constructor_Main_Test3[gopurs_runtime.Value, int64]{1, int64(42), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_7
			})})))})
			_ = v_7_9
			__local_var_8_10 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("cmap - quantified argument stays unchanged"), gopurs_runtime.Bool((v_7_9.Type == 9 && v_7_9.IntVal == 227416251) && ((((*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value])(v_7_9.UnsafePtr).V0) == (int64(42))) && (((gopurs_runtime.Apply2(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqArray())).V0, gopurs_runtime.Apply((*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value])(v_7_9.UnsafePtr).V1, func() gopurs_runtime.Value {
				arr := []int64{__local_var_1_1.IntVal}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()), func() gopurs_runtime.Value {
				arr := []int64{__local_var_1_1.IntVal}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).IntVal) != (0)) && ((gopurs_runtime.Apply2(Rebox_Main_3790796878_131790935(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqArray1())).V0, gopurs_runtime.Apply((*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value])(v_7_9.UnsafePtr).V1, func() gopurs_runtime.Value {
				arr := []string{"x"}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()), func() gopurs_runtime.Value {
				arr := []string{"x"}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).IntVal) != (0)))))), gopurs_runtime.Value{})
			_ = __local_var_8_10
			// TAST (Let): v_9_11 shape=App(Var) bindingType=(ADT ["Main","Test"] [(Func [(TypeVar a)] Boolean), (Record (Row [value: Int] Empty))])
			v_9_11 := gopurs_runtime.Apply(Get_Main_adapt(), gopurs_runtime.Value{Type: 9, IntVal: 3712677948, UnsafePtr: unsafe.Pointer(Rebox_Main_3811136309_1074635502((&Constructor_Main_Test4[gopurs_runtime.Value, int64]{1, int64(42), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((v1_9.IntVal) == (__local_var_1_1.IntVal))
			})})))})
			_ = v_9_11
			__local_var_10_12 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("cmap - constrained constructor"), gopurs_runtime.Bool((v_9_11.Type == 9 && v_9_11.IntVal == 3712677948) && ((((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value])(v_9_11.UnsafePtr).V0) == (int64(42))) && (((gopurs_runtime.Apply((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value])(v_9_11.UnsafePtr).V1, func() gopurs_runtime.Value {
				orig := struct {
					value int64
				}{__local_var_1_1.IntVal}
				_ = orig
				return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
			}()).IntVal) != (0)) && (((gopurs_runtime.Apply((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value])(v_9_11.UnsafePtr).V1, func() gopurs_runtime.Value {
				orig := struct {
					value int64
				}{int64(0)}
				_ = orig
				return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
			}()).IntVal) != (0)) != (true)))))), gopurs_runtime.Value{})
			_ = __local_var_10_12
			// TAST (Let): v_11_13 shape=App(Var) bindingType=(ADT ["Main","Test"] [(Func [(TypeVar a)] Boolean), (Record (Row [value: Int] Empty))])
			v_11_13 := gopurs_runtime.Apply(Get_Main_adapt(), gopurs_runtime.Value{Type: 9, IntVal: 1063363133, UnsafePtr: unsafe.Pointer(Rebox_Main_1320231956_4217394063((&Constructor_Main_Test5[gopurs_runtime.Value, int64]{1, (*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(v1_11.IntVal, int64(2)))
			})})).UnsafePtr)), Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((v1_11.IntVal) == (__local_var_1_1.IntVal))
				}), gopurs_runtime.Int(int64(42))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))})))})
			_ = v_11_13
			__local_var_12_14 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("cmap - function and first tuple argument"), gopurs_runtime.Bool((v_11_13.Type == 9 && v_11_13.IntVal == 1063363133) && (((gopurs_runtime.Int(int64(len((*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value])(v_11_13.UnsafePtr).V0))).IntVal) == (int64(1))) && (((gopurs_runtime.Apply(gopurs_runtime.ArrayAccess(gopurs_runtime.Array((*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value])(v_11_13.UnsafePtr).V0), 0), func() gopurs_runtime.Value {
				orig := struct {
					value int64
				}{__local_var_1_1.IntVal}
				_ = orig
				return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
			}()).IntVal) == (int64(9))) && (((gopurs_runtime.Apply(((*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value])(v_11_13.UnsafePtr).V1).V0, func() gopurs_runtime.Value {
				orig := struct {
					value int64
				}{__local_var_1_1.IntVal}
				_ = orig
				return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
			}()).IntVal) != (0)) && ((((gopurs_runtime.Apply(((*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value])(v_11_13.UnsafePtr).V1).V0, func() gopurs_runtime.Value {
				orig := struct {
					value int64
				}{int64(0)}
				_ = orig
				return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
			}()).IntVal) != (0)) != (true)) && ((((*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value])(v_11_13.UnsafePtr).V1).V1) == (int64(42))))))))), gopurs_runtime.Value{})
			_ = __local_var_12_14
			// TAST (Let): v_13_15 shape=App(Var) bindingType=(ADT ["Main","Test"] [(Func [(TypeVar a)] Boolean), (Record (Row [value: Int] Empty))])
			v_13_15 := gopurs_runtime.Apply(Get_Main_adapt(), gopurs_runtime.Value{Type: 9, IntVal: 4013407934, UnsafePtr: unsafe.Pointer(Rebox_Main_1149948919_3768263468((&Constructor_Main_Test6[gopurs_runtime.Value, int64]{1, func() struct {
				nested []struct {
					x gopurs_runtime.Value
				}
			} {
				orig := gopurs_runtime.RecordDict1("nested", func() gopurs_runtime.Value {
					arr := []struct {
						x gopurs_runtime.Value
					}{struct {
						x gopurs_runtime.Value
					}{gopurs_runtime.Func(func(r_13 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.RecordGet(r_13, "a").IntVal) == (__local_var_1_1.IntVal))
					})}}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = func() gopurs_runtime.Value {
							orig := v
							_ = orig
							return gopurs_runtime.RecordDict1("x", orig.x)
						}()
					}
					return gopurs_runtime.Array(boxed)
				}())
				_ = orig
				clone := struct {
					nested []struct {
						x gopurs_runtime.Value
					}
				}{}
				clone.nested = func() []struct {
					x gopurs_runtime.Value
				} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "nested").UnsafePtr)
					unboxed := make([]struct {
						x gopurs_runtime.Value
					}, len(arr))
					for i, v := range arr {
						unboxed[i] = func() struct {
							x gopurs_runtime.Value
						} {
							orig := v
							_ = orig
							clone := struct {
								x gopurs_runtime.Value
							}{}
							clone.x = gopurs_runtime.RecordGet(orig, "x")
							return clone
						}()
					}
					return unboxed
				}()
				return clone
			}()})))})
			_ = v_13_15
			__local_var_14_16 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("cmap - nested records"), gopurs_runtime.Bool((v_13_15.Type == 9 && v_13_15.IntVal == 4013407934) && (((gopurs_runtime.Int(int64(len((*Constructor_Main_Test6[gopurs_runtime.Value, gopurs_runtime.Value])(v_13_15.UnsafePtr).V0.nested))).IntVal) == (int64(1))) && (((gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
				arr := (*Constructor_Main_Test6[gopurs_runtime.Value, gopurs_runtime.Value])(v_13_15.UnsafePtr).V0.nested
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						orig := v
						_ = orig
						return gopurs_runtime.RecordDict1("x", orig.x)
					}()
				}
				return gopurs_runtime.Array(boxed)
			}(), 0), "x"), func() gopurs_runtime.Value {
				orig := struct {
					a struct {
						value int64
					}
				}{struct {
					value int64
				}{__local_var_1_1.IntVal}}
				_ = orig
				return gopurs_runtime.RecordDict1("a", func() gopurs_runtime.Value {
					orig := orig.a
					_ = orig
					return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
				}())
			}()).IntVal) != (0)) && (((gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
				arr := (*Constructor_Main_Test6[gopurs_runtime.Value, gopurs_runtime.Value])(v_13_15.UnsafePtr).V0.nested
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						orig := v
						_ = orig
						return gopurs_runtime.RecordDict1("x", orig.x)
					}()
				}
				return gopurs_runtime.Array(boxed)
			}(), 0), "x"), func() gopurs_runtime.Value {
				orig := struct {
					a struct {
						value int64
					}
				}{struct {
					value int64
				}{int64(0)}}
				_ = orig
				return gopurs_runtime.RecordDict1("a", func() gopurs_runtime.Value {
					orig := orig.a
					_ = orig
					return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
				}())
			}()).IntVal) != (0)) != (true)))))), gopurs_runtime.Value{})
			_ = __local_var_14_16
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Test0[T_f any, T_a any] struct {
	Rc uint32
}

type Constructor_Main_Test1[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Test2[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Test3[T_f any, T_a any] struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
}

type Constructor_Main_Test4[T_f any, T_a any] struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
}

type Constructor_Main_Test5[T_f any, T_a any] struct {
	Rc uint32
	V0 []gopurs_runtime.Value
	V1 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]
}

type Constructor_Main_Test6[T_f any, T_a any] struct {
	Rc uint32
	V0 struct {
		nested []struct {
			x gopurs_runtime.Value
		}
	}
}

func Call_Main_identity(x_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var x_0 []gopurs_runtime.Value = x_0_loop
	_ = x_0
	return (*(*[]gopurs_runtime.Value)((gopurs_runtime.Array(x_0)).UnsafePtr))
}

func Call_Main_contravariantTest(dictContravariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
	_ = dictContravariant_0
	return gopurs_runtime.Value{Type: 9, IntVal: 85171506, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Contravariant_Contravariant[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if m_2.Type == 9 && m_2.IntVal == 2074462008 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2074462008, UnsafePtr: unsafe.Pointer(nil)}
				goto end_branch_0
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 3720114489 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), (*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, f_1)}))}
				goto end_branch_0
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 2375191994 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), (*Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), v_3, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), v_4, f_1)
					}))
				}))}))}
				goto end_branch_0
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 227416251 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1}))}
				goto end_branch_0
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 3712677948 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3712677948, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), f_1, (*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)}))}
				goto end_branch_0
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 1063363133 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1063363133, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap4 := gopurs_runtime.Array((*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0)
					_ = arr_val_arrayMap4
					arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
					_ = arr_go_arrayMap4
					res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
					_ = res_go_arrayMap4
					for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
						res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(gopurs_runtime.Func(func(b2c_3 gopurs_runtime.Value) gopurs_runtime.Value {
							return Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), f_1, Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()), b2c_3, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
								return x_4
							})))
						}), v_arrayMap4)
					}
					return gopurs_runtime.Array(res_go_arrayMap4)
				}()).UnsafePtr)), Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
					}{gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), ((*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1).V0, f_1), gopurs_runtime.Int(((*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1).V1)}
					return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
				}()))}))}
				goto end_branch_0
			} else {

			}
		}
		{
			if m_2.Type == 9 && m_2.IntVal == 4013407934 {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 4013407934, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test6[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() struct {
					nested []struct {
						x gopurs_runtime.Value
					}
				} {
					clone := (*Constructor_Main_Test6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0
					clone.nested = func() []struct {
						x gopurs_runtime.Value
					} {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
							arr_val_arrayMap5 := func() gopurs_runtime.Value {
								arr := (*Constructor_Main_Test6[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0.nested
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = func() gopurs_runtime.Value {
										orig := v
										_ = orig
										return gopurs_runtime.RecordDict1("x", orig.x)
									}()
								}
								return gopurs_runtime.Array(boxed)
							}()
							_ = arr_val_arrayMap5
							arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
							_ = arr_go_arrayMap5
							res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
							_ = res_go_arrayMap5
							for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
								res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
									return func() gopurs_runtime.Value {
										orig := func() struct {
											x gopurs_runtime.Value
										} {
											orig := gopurs_runtime.RecordUpdate1(v1_3, "x", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
												return func() gopurs_runtime.Value {
													orig := func() struct {
														a gopurs_runtime.Value
													} {
														orig := gopurs_runtime.RecordUpdate1(v2_4, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(v2_4, "a")))
														_ = orig
														clone := struct {
															a gopurs_runtime.Value
														}{}
														clone.a = gopurs_runtime.RecordGet(orig, "a")
														return clone
													}()
													_ = orig
													return gopurs_runtime.RecordDict1("a", orig.a)
												}()
											}), gopurs_runtime.RecordGet(v1_3, "x")))
											_ = orig
											clone := struct {
												x gopurs_runtime.Value
											}{}
											clone.x = gopurs_runtime.RecordGet(orig, "x")
											return clone
										}()
										_ = orig
										return gopurs_runtime.RecordDict1("x", orig.x)
									}()
								}), v_arrayMap5)
							}
							return gopurs_runtime.Array(res_go_arrayMap5)
						}()).UnsafePtr))).UnsafePtr)
						unboxed := make([]struct {
							x gopurs_runtime.Value
						}, len(arr))
						for i, v := range arr {
							unboxed[i] = func() struct {
								x gopurs_runtime.Value
							} {
								orig := v
								_ = orig
								clone := struct {
									x gopurs_runtime.Value
								}{}
								clone.x = gopurs_runtime.RecordGet(orig, "x")
								return clone
							}()
						}
						return unboxed
					}()
					return clone
				}()}))}
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_0:
		return __t0
	})}))}
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

func Rebox_Main_1149948919_3768263468(in *Constructor_Main_Test6[gopurs_runtime.Value, int64]) *Constructor_Main_Test6[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Test6[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1204441265_2762907754(in *Constructor_Main_Test0[gopurs_runtime.Value, int64]) *Constructor_Main_Test0[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Test0[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_131790935_3790796878(in *Constructor_Data_Eq_Eq[[]string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1320231956_4217394063(in *Constructor_Main_Test5[gopurs_runtime.Value, int64]) *Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_138441832_3415943795(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]{}
	out.V0 = in.V0
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Main_2838221171_1161568424(in *Constructor_Main_Test2[gopurs_runtime.Value, int64]) *Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3008504208_1610699019(in *Constructor_Main_Test1[gopurs_runtime.Value, int64]) *Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
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

func Rebox_Main_3790796878_131790935(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]string])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_378698611(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]int64])(unsafe.Pointer(in))
}

func Rebox_Main_3811136309_1074635502(in *Constructor_Main_Test4[gopurs_runtime.Value, int64]) *Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_590733650_648073545(in *Constructor_Main_Test3[gopurs_runtime.Value, int64]) *Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}
