package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Test1
}

var cache_Main_Test2 gopurs_runtime.Value
var once_Main_Test2 sync.Once

func Get_Main_Test2() gopurs_runtime.Value {
	once_Main_Test2.Do(func() {
		cache_Main_Test2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal, value1}))}
			})
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
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0.IntVal, value1, value2, value3}))}
					})
				})
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
				return gopurs_runtime.Value{Type: 9, IntVal: 3712677948, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*(*[]gopurs_runtime.Value)((value0).UnsafePtr)), Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](value1))}))}
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
			return gopurs_runtime.Value{Type: 9, IntVal: 1063363133, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, func() struct {
				nested []struct {
					x gopurs_runtime.Value
					/*gopurs:record-v1:1:0078*/
				}
				/*gopurs:record-v1:1:006e00650073007400650064*/
			} {
				orig := value0
				_ = orig
				clone := struct {
					nested []struct {
						x gopurs_runtime.Value
						/*gopurs:record-v1:1:0078*/
					}
					/*gopurs:record-v1:1:006e00650073007400650064*/
				}{}
				clone.nested = func() []struct {
					x gopurs_runtime.Value
					/*gopurs:record-v1:1:0078*/
				} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "nested").UnsafePtr)
					unboxed := make([]struct {
						x gopurs_runtime.Value
						/*gopurs:record-v1:1:0078*/
					}, len(arr))
					for i, v := range arr {
						unboxed[i] = func() struct {
							x gopurs_runtime.Value
							/*gopurs:record-v1:1:0078*/
						} {
							orig := v
							_ = orig
							clone := struct {
								x gopurs_runtime.Value
								/*gopurs:record-v1:1:0078*/
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
	return cache_Main_Test5
}

var cache_Main_profunctorTest gopurs_runtime.Value
var once_Main_profunctorTest sync.Once

func Get_Main_profunctorTest() gopurs_runtime.Value {
	once_Main_profunctorTest.Do(func() {
		cache_Main_profunctorTest = gopurs_runtime.Func(func(dictProfunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_profunctorTest(dictProfunctor_0_box)
		})
	})
	return cache_Main_profunctorTest
}

var cache_Main_adapt gopurs_runtime.Value
var once_Main_adapt sync.Once

func Get_Main_adapt() gopurs_runtime.Value {
	once_Main_adapt.Do(func() {
		cache_Main_adapt = gopurs_runtime.Apply2(Call_Data_Profunctor_dimap(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](Get_Main___cached_dict_0())), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordGet(v_0, "value")
		}), Get_Data_Show_showIntImpl())
	})
	return cache_Main_adapt
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref_go__new(), gopurs_runtime.Int(int64(7)))
			_ = __local_var_0_0
			__local_var_1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_2
			n_1_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_2), gopurs_runtime.Value{})
			_ = n_1_1
			var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply(Get_Main_adapt(), gopurs_runtime.Value{Type: 9, IntVal: 2074462008, UnsafePtr: unsafe.Pointer(Rebox_Main_4256224396_3701891308(nil))})
			_ = __t_tag_4
			_dollar___unused_2_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("dimap - empty constructor"), gopurs_runtime.Bool((__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2074462008))), gopurs_runtime.Value{})
			_ = _dollar___unused_2_3
			// TAST (Let): v_3_5 shape=App(Var) bindingType=(ADT ["Main","Test"] [(ADT ["Prim","Function"] []), (Record (Row [value: Int] Empty)), String])
			v_3_5 := gopurs_runtime.Apply(Get_Main_adapt(), gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer(Rebox_Main_2452161453_2724589709((&Constructor_Main_Test1[gopurs_runtime.Value, int64, int64]{1, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((v1_3.IntVal) == (n_1_1.IntVal))
			}), gopurs_runtime.IntAdd(n_1_1.IntVal, int64(1))})))})
			_ = v_3_5
			_dollar___unused_4_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("dimap - predicate and result"), gopurs_runtime.Bool((v_3_5.Type == 9 && v_3_5.IntVal == 3720114489) && (((gopurs_runtime.Apply((*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_3_5.UnsafePtr).V0, func() gopurs_runtime.Value {
				orig := struct {
					value int64
					/*gopurs:record-v1:1:00760061006c00750065*/
				}{n_1_1.IntVal}
				_ = orig
				return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
			}()).IntVal) != (0)) && ((((gopurs_runtime.Apply((*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_3_5.UnsafePtr).V0, func() gopurs_runtime.Value {
				orig := struct {
					value int64
					/*gopurs:record-v1:1:00760061006c00750065*/
				}{int64(0)}
				_ = orig
				return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
			}()).IntVal) != (0)) != (true)) && (((*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_3_5.UnsafePtr).V1.StrVal()) == ("8")))))), gopurs_runtime.Value{})
			_ = _dollar___unused_4_6
			// TAST (Let): v_5_7 shape=App(Var) bindingType=(ADT ["Main","Test"] [(ADT ["Prim","Function"] []), (Record (Row [value: Int] Empty)), String])
			v_5_7 := gopurs_runtime.Apply(Get_Main_adapt(), gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer(Rebox_Main_1838453838_1232183214((&Constructor_Main_Test2[gopurs_runtime.Value, int64, int64]{1, int64(42), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_5
			})})))})
			_ = v_5_7
			_dollar___unused_6_8 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("dimap - quantified argument stays unchanged"), gopurs_runtime.Bool((v_5_7.Type == 9 && v_5_7.IntVal == 2375191994) && ((((*Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_5_7.UnsafePtr).V0) == (int64(42))) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), func() gopurs_runtime.Value {
				arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply((*Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_5_7.UnsafePtr).V1, func() gopurs_runtime.Value {
						arr := []int64{n_1_1.IntVal}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := []int64{n_1_1.IntVal}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), func() gopurs_runtime.Value {
				arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply((*Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_5_7.UnsafePtr).V1, func() gopurs_runtime.Value {
						arr := []string{"x"}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr {
						unboxed[i] = v.StrVal()
					}
					return unboxed
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := []string{"x"}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).IntVal) != (0)))))), gopurs_runtime.Value{})
			_ = _dollar___unused_6_8
			// TAST (Let): v_7_9 shape=App(Var) bindingType=(ADT ["Main","Test"] [(ADT ["Prim","Function"] []), (Record (Row [value: Int] Empty)), String])
			v_7_9 := gopurs_runtime.Apply(Get_Main_adapt(), gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer(Rebox_Main_4085941359_893595471((&Constructor_Main_Test3[gopurs_runtime.Value, int64, int64]{1, int64(42), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(v1_7.IntVal, int64(1)))
			}), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(v1_7.IntVal, int64(2)))
			}), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(v1_7.IntVal, int64(3)))
			})})))})
			_ = v_7_9
			_dollar___unused_8_10 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("dimap - both, left-only and right-only arguments"), gopurs_runtime.Bool((v_7_9.Type == 9 && v_7_9.IntVal == 227416251) && ((((*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_7_9.UnsafePtr).V0) == (int64(42))) && (((gopurs_runtime.Apply((*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_7_9.UnsafePtr).V1, func() gopurs_runtime.Value {
				orig := struct {
					value int64
					/*gopurs:record-v1:1:00760061006c00750065*/
				}{n_1_1.IntVal}
				_ = orig
				return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
			}()).StrVal()) == ("8")) && (((gopurs_runtime.Apply((*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_7_9.UnsafePtr).V2, func() gopurs_runtime.Value {
				orig := struct {
					value int64
					/*gopurs:record-v1:1:00760061006c00750065*/
				}{n_1_1.IntVal}
				_ = orig
				return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
			}()).IntVal) == (int64(9))) && ((gopurs_runtime.Apply((*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_7_9.UnsafePtr).V3, gopurs_runtime.Int(int64(4))).StrVal()) == ("7"))))))), gopurs_runtime.Value{})
			_ = _dollar___unused_8_10
			// TAST (Let): v_9_11 shape=App(Var) bindingType=(ADT ["Main","Test"] [(ADT ["Prim","Function"] []), (Record (Row [value: Int] Empty)), String])
			v_9_11 := gopurs_runtime.Apply(Get_Main_adapt(), gopurs_runtime.Value{Type: 9, IntVal: 3712677948, UnsafePtr: unsafe.Pointer(Rebox_Main_1649529352_1095196264((&Constructor_Main_Test4[gopurs_runtime.Value, int64, int64]{1, (*(*[]gopurs_runtime.Value)((gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.IntAdd(v1_9.IntVal, int64(1)))
			})})).UnsafePtr)), (&Constructor_Data_Tuple_Tuple[int64, int64]{1, n_1_1.IntVal, int64(42)})})))})
			_ = v_9_11
			_dollar___unused_10_12 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("dimap - function and first tuple argument"), gopurs_runtime.Bool((v_9_11.Type == 9 && v_9_11.IntVal == 3712677948) && (((gopurs_runtime.Int(int64(len((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9_11.UnsafePtr).V0))).IntVal) == (int64(1))) && (((gopurs_runtime.Apply(gopurs_runtime.ArrayAccess(gopurs_runtime.Array((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9_11.UnsafePtr).V0), 0), func() gopurs_runtime.Value {
				orig := struct {
					value int64
					/*gopurs:record-v1:1:00760061006c00750065*/
				}{n_1_1.IntVal}
				_ = orig
				return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
			}()).IntVal) == (int64(8))) && (((((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9_11.UnsafePtr).V1).V0.StrVal()) == ("7")) && ((((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9_11.UnsafePtr).V1).V1) == (int64(42)))))))), gopurs_runtime.Value{})
			_ = _dollar___unused_10_12
			// TAST (Let): v_11_13 shape=App(Var) bindingType=(ADT ["Main","Test"] [(ADT ["Prim","Function"] []), (Record (Row [value: Int] Empty)), String])
			v_11_13 := gopurs_runtime.Apply(Get_Main_adapt(), gopurs_runtime.Value{Type: 9, IntVal: 1063363133, UnsafePtr: unsafe.Pointer(Rebox_Main_4140433705_117894665((&Constructor_Main_Test5[gopurs_runtime.Value, int64, int64]{1, struct {
				nested []struct {
					x gopurs_runtime.Value
					/*gopurs:record-v1:1:0078*/
				}
				/*gopurs:record-v1:1:006e00650073007400650064*/
			}{[]struct {
				x gopurs_runtime.Value
				/*gopurs:record-v1:1:0078*/
			}{struct {
				x gopurs_runtime.Value
				/*gopurs:record-v1:1:0078*/
			}{gopurs_runtime.Func(func(r_11 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordDict1("b", gopurs_runtime.Int(gopurs_runtime.IntAdd(gopurs_runtime.RecordGet(r_11, "a").IntVal, int64(1))))
			})}}}})))})
			_ = v_11_13
			_dollar___unused_12_14 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assert_prime_(), gopurs_runtime.Str("dimap - nested records"), gopurs_runtime.Bool((v_11_13.Type == 9 && v_11_13.IntVal == 1063363133) && (((gopurs_runtime.Int(int64(len((*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11_13.UnsafePtr).V0.nested))).IntVal) == (int64(1))) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("b")
			})), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})), "eq"), func() gopurs_runtime.Value {
				orig := func() struct {
					b string
					/*gopurs:record-v1:1:0062*/
				} {
					orig := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.ArrayAccess(func() gopurs_runtime.Value {
						arr := (*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11_13.UnsafePtr).V0.nested
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
								/*gopurs:record-v1:1:00760061006c00750065*/
							}
							/*gopurs:record-v1:1:0061*/
						}{struct {
							value int64
							/*gopurs:record-v1:1:00760061006c00750065*/
						}{n_1_1.IntVal}}
						_ = orig
						return gopurs_runtime.RecordDict1("a", func() gopurs_runtime.Value {
							orig := orig.a
							_ = orig
							return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(orig.value))
						}())
					}())
					_ = orig
					clone := struct {
						b string
						/*gopurs:record-v1:1:0062*/
					}{}
					clone.b = gopurs_runtime.RecordGet(orig, "b").StrVal()
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict1("b", gopurs_runtime.Str(orig.b))
			}(), func() gopurs_runtime.Value {
				orig := struct {
					b string
					/*gopurs:record-v1:1:0062*/
				}{"8"}
				_ = orig
				return gopurs_runtime.RecordDict1("b", gopurs_runtime.Str(orig.b))
			}()).IntVal) != (0))))), gopurs_runtime.Value{})
			_ = _dollar___unused_12_14
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main___cached_dict_0 gopurs_runtime.Value
var once_Main___cached_dict_0 sync.Once

func Get_Main___cached_dict_0() gopurs_runtime.Value {
	once_Main___cached_dict_0.Do(func() {
		cache_Main___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](Call_Main_profunctorTest(gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](Get_Data_Profunctor_profunctorFn()))})))}
	})
	return cache_Main___cached_dict_0
}

type Constructor_Main_Test0[T_f any, T_a any, T_b any] struct {
	Rc uint32
}

type Constructor_Main_Test1[T_f any, T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_b
}

type Constructor_Main_Test2[T_f any, T_a any, T_b any] struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
}

type Constructor_Main_Test3[T_f any, T_a any, T_b any] struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}

type Constructor_Main_Test4[T_f any, T_a any, T_b any] struct {
	Rc uint32
	V0 []gopurs_runtime.Value
	V1 *Constructor_Data_Tuple_Tuple[T_b, int64]
}

type Constructor_Main_Test5[T_f any, T_a any, T_b any] struct {
	Rc uint32
	V0 struct {
		nested []struct {
			x gopurs_runtime.Value
			/*gopurs:record-v1:1:0078*/
		}
		/*gopurs:record-v1:1:006e00650073007400650064*/
	}
}

func Call_Main_profunctorTest(dictProfunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
	_ = dictProfunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t4 gopurs_runtime.Value
		{
			if m_3.Type == 9 && m_3.IntVal == 2074462008 {
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2074462008, UnsafePtr: unsafe.Pointer(nil)}
				goto end_branch_4
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 3720114489 {
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), (*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, f_1), gopurs_runtime.Apply(g_2, (*Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)}))}
				goto end_branch_4
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 2375191994 {
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, (*Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1}))}
				goto end_branch_4
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 227416251 {
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0, gopurs_runtime.Apply3(Call_Data_Profunctor_dimap(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0)), f_1, g_2, (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1), gopurs_runtime.Apply(Call_Data_Profunctor_lcmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0), f_1), (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V2), gopurs_runtime.Apply(Call_Data_Profunctor_rmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0), g_2), (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V3)}))}
				goto end_branch_4
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 3712677948 {
				arrayIntrinsic_arg_0 := Call_Data_Profunctor_lcmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](Get_Data_Profunctor_profunctorFn()), f_1)
				_ = arrayIntrinsic_arg_0
				arrayIntrinsic_arg_1 := gopurs_runtime.Array((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0)
				_ = arrayIntrinsic_arg_1
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3712677948, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
					arr_val_arrayMap4 := arrayIntrinsic_arg_1
					_ = arr_val_arrayMap4
					arr_go_arrayMap4 := (*[]gopurs_runtime.Value)(arr_val_arrayMap4.UnsafePtr)
					_ = arr_go_arrayMap4
					res_go_arrayMap4 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap4))
					_ = res_go_arrayMap4
					for i_arrayMap4, v_arrayMap4 := range *arr_go_arrayMap4 {
						res_go_arrayMap4[i_arrayMap4] = gopurs_runtime.Apply(arrayIntrinsic_arg_0, v_arrayMap4)
					}
					return gopurs_runtime.Array(res_go_arrayMap4)
				}()).UnsafePtr)), Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Bifunctor_lmap(Rebox_Main_1495429347_1688994542(Rebox_Main_1688994542_1495429347(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Bifunctor_bifunctorTuple()))), g_2), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3415943795_138441832((*Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1))})))}))}
				goto end_branch_4
			} else {

			}
		}
		{
			if m_3.Type == 9 && m_3.IntVal == 1063363133 {
				arrayIntrinsic_arg_2 := gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.RecordUpdate1(v1_4, "x", gopurs_runtime.Apply3(Call_Data_Profunctor_dimap(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](dictProfunctor_0)), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.RecordUpdate1(v2_5, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(v2_5, "a")))
					}), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.RecordUpdate1(v2_5, "b", gopurs_runtime.Apply(g_2, gopurs_runtime.RecordGet(v2_5, "b")))
					}), gopurs_runtime.RecordGet(v1_4, "x")))
				})
				_ = arrayIntrinsic_arg_2
				arrayIntrinsic_arg_3 := func() gopurs_runtime.Value {
					arr := (*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0.nested
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
				_ = arrayIntrinsic_arg_3
				__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1063363133, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, func() struct {
					nested []struct {
						x gopurs_runtime.Value
						/*gopurs:record-v1:1:0078*/
					}
					/*gopurs:record-v1:1:006e00650073007400650064*/
				} {
					clone := (*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0
					clone.nested = func() []struct {
						x gopurs_runtime.Value
						/*gopurs:record-v1:1:0078*/
					} {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
							arr_val_arrayMap5 := arrayIntrinsic_arg_3
							_ = arr_val_arrayMap5
							arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
							_ = arr_go_arrayMap5
							res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
							_ = res_go_arrayMap5
							for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
								res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(arrayIntrinsic_arg_2, v_arrayMap5)
							}
							return gopurs_runtime.Array(res_go_arrayMap5)
						}()).UnsafePtr))).UnsafePtr)
						unboxed := make([]struct {
							x gopurs_runtime.Value
							/*gopurs:record-v1:1:0078*/
						}, len(arr))
						for i, v := range arr {
							unboxed[i] = func() struct {
								x gopurs_runtime.Value
								/*gopurs:record-v1:1:0078*/
							} {
								orig := v
								_ = orig
								clone := struct {
									x gopurs_runtime.Value
									/*gopurs:record-v1:1:0078*/
								}{}
								clone.x = gopurs_runtime.RecordGet(orig, "x")
								return clone
							}()
						}
						return unboxed
					}()
					return clone
				}()}))}
				goto end_branch_4
			} else {

			}
		}
		{
			__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_4:
		return __t4
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

func Rebox_Main_138441832_3415943795(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]{}
	out.V0 = in.V0
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Main_1495429347_1688994542(in *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1649529352_1095196264(in *Constructor_Main_Test4[gopurs_runtime.Value, int64, int64]) *Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Test4[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = Rebox_Main_3363075976_3415943795(in.V1)
	return out
}

func Rebox_Main_1688994542_1495429347(in *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]) *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Main_1838453838_1232183214(in *Constructor_Main_Test2[gopurs_runtime.Value, int64, int64]) *Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Test2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2452161453_2724589709(in *Constructor_Main_Test1[gopurs_runtime.Value, int64, int64]) *Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Test1[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Main_3363075976_3415943795(in *Constructor_Data_Tuple_Tuple[int64, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Main_3415943795_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
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

func Rebox_Main_4085941359_893595471(in *Constructor_Main_Test3[gopurs_runtime.Value, int64, int64]) *Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Test3[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_4140433705_117894665(in *Constructor_Main_Test5[gopurs_runtime.Value, int64, int64]) *Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Test5[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_4256224396_3701891308(in *Constructor_Main_Test0[gopurs_runtime.Value, int64, int64]) *Constructor_Main_Test0[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Test0[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}
