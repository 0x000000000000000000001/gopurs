package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), Get_Data_Show_showIntImpl())})}
	})
	return cache_Main_showArray
}

var cache_Main_add gopurs_runtime.Value
var once_Main_add sync.Once

func Get_Main_add() gopurs_runtime.Value {
	once_Main_add.Do(func() {
		cache_Main_add = Get_Data_Semiring_numAdd()
	})
	return cache_Main_add
}

var cache_Main_Nothing gopurs_runtime.Value
var once_Main_Nothing sync.Once

func Get_Main_Nothing() gopurs_runtime.Value {
	once_Main_Nothing.Do(func() {
		cache_Main_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just)(nil))}
	})
	return cache_Main_Nothing
}

var cache_Main_Just gopurs_runtime.Value
var once_Main_Just sync.Once

func Get_Main_Just() gopurs_runtime.Value {
	once_Main_Just.Do(func() {
		cache_Main_Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(&Constructor_Main_Just{1, value0})}
		})
	})
	return cache_Main_Just
}

var cache_Main_test8 gopurs_runtime.Value
var once_Main_test8 sync.Once

func Get_Main_test8() gopurs_runtime.Value {
	once_Main_test8.Do(func() {
		cache_Main_test8 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, dictApplicative1_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test8(dictApplicative_0_box, dictApplicative1_1_box, v_2_box)
		})
	})
	return cache_Main_test8
}

var cache_Main_test6 gopurs_runtime.Value
var once_Main_test6 sync.Once

func Get_Main_test6() gopurs_runtime.Value {
	once_Main_test6.Do(func() {
		cache_Main_test6 = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, dictPartial_1_box gopurs_runtime.Value, mx_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test6(dictApplicative_0_box, dictPartial_1_box, mx_2_box, v_3_box)
		})
	})
	return cache_Main_test6
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test5(dictApply_0_box)
		})
	})
	return cache_Main_test5
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test4(dictApply_0_box)
		})
	})
	return cache_Main_test4
}

var cache_Main_test11 gopurs_runtime.Value
var once_Main_test11 sync.Once

func Get_Main_test11() gopurs_runtime.Value {
	once_Main_test11.Do(func() {
		cache_Main_test11 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test11(dictApply_0_box)
		})
	})
	return cache_Main_test11
}

var cache_Main_test10 gopurs_runtime.Value
var once_Main_test10 sync.Once

func Get_Main_test10() gopurs_runtime.Value {
	once_Main_test10.Do(func() {
		cache_Main_test10 = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test10(dictApplicative_0_box, v_1_box)
		})
	})
	return cache_Main_test10
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test1(dictApplicative_0_box, v_1_box)
		})
	})
	return cache_Main_test1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str("X"))
			_ = __local_var_0_0
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				r_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
				_ = r_1_1
				f_prime__2_5 := gopurs_runtime.Apply(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						a_prime__2_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Str("D"), r_1_1), gopurs_runtime.Value{})
						_ = a_prime__2_6
						return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v3_5 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Str((((v1_3.StrVal()) + (v2_4.StrVal())) + ("n")) + (v3_5.StrVal()))
								})
							})
						})
					})
				}), gopurs_runtime.Value{})
				_ = f_prime__2_5
				a_prime__3_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), r_1_1), gopurs_runtime.Value{})
				_ = a_prime__3_7
				f_prime__2_4 := gopurs_runtime.Apply(f_prime__2_5, a_prime__3_7)
				_ = f_prime__2_4
				f_prime__2_3 := gopurs_runtime.Apply(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(f_prime__2_4, gopurs_runtime.Str("o"))
					})
				}), gopurs_runtime.Value{})
				_ = f_prime__2_3
				__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(f_prime__2_3, gopurs_runtime.Str("e"))
					})
				}), gopurs_runtime.Value{})
				_ = __local_var_2_2
				return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), __local_var_2_2), gopurs_runtime.Value{})
			})
		}()
	})
	return cache_Main_main
}

var cache_Main_functorMaybe gopurs_runtime.Value
var once_Main_functorMaybe sync.Once

func Get_Main_functorMaybe() gopurs_runtime.Value {
	once_Main_functorMaybe.Do(func() {
		cache_Main_functorMaybe = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 *Constructor_Main_Just
				{
					if v1_1.Type == 9 && v1_1.IntVal == 3271839782 && v1_1.UnsafePtr == nil {
						__t0 = (*Constructor_Main_Just)(nil)
						goto end_branch_0
					} else {

					}
				}
				{
					if v1_1.Type == 9 && v1_1.IntVal == 3271839782 && v1_1.UnsafePtr != nil {
						__t0 = &Constructor_Main_Just{1, gopurs_runtime.Apply(v_0, (*Constructor_Main_Just)(v1_1.UnsafePtr).V0)}
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
				}
			end_branch_0:
				return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t0)}
			})
		})})}
	})
	return cache_Main_functorMaybe
}

var cache_Main_applyMaybe gopurs_runtime.Value
var once_Main_applyMaybe sync.Once

func Get_Main_applyMaybe() gopurs_runtime.Value {
	once_Main_applyMaybe.Do(func() {
		cache_Main_applyMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Main_functorMaybe()))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 *Constructor_Main_Just
				{
					if (v_0.Type == 9 && v_0.IntVal == 3271839782 && v_0.UnsafePtr != nil) && (v1_1.Type == 9 && v1_1.IntVal == 3271839782 && v1_1.UnsafePtr != nil) {
						__t0 = &Constructor_Main_Just{1, gopurs_runtime.Apply((*Constructor_Main_Just)(v_0.UnsafePtr).V0, (*Constructor_Main_Just)(v1_1.UnsafePtr).V0)}
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = (*Constructor_Main_Just)(nil)
				}
			end_branch_0:
				return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t0)}
			})
		})})}
	})
	return cache_Main_applyMaybe
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Call_Main_test2(v_0_box))}
		})
	})
	return cache_Main_test2
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Call_Main_test3(v_0_box))}
		})
	})
	return cache_Main_test3
}

var cache_Main_test9 gopurs_runtime.Value
var once_Main_test9 sync.Once

func Get_Main_test9() gopurs_runtime.Value {
	once_Main_test9.Do(func() {
		cache_Main_test9 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Call_Main_test9(v_0_box))}
		})
	})
	return cache_Main_test9
}

var cache_Main_applicativeMaybe gopurs_runtime.Value
var once_Main_applicativeMaybe sync.Once

func Get_Main_applicativeMaybe() gopurs_runtime.Value {
	once_Main_applicativeMaybe.Do(func() {
		cache_Main_applicativeMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Main_applyMaybe()))}
		}), Get_Main_Just()})}
	})
	return cache_Main_applicativeMaybe
}

type Constructor_Main_Nothing struct {
	Rc uint32
}

type Constructor_Main_Just struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_test8(dictApplicative_0_loop gopurs_runtime.Value, dictApplicative1_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	var dictApplicative1_1 gopurs_runtime.Value = dictApplicative1_1_loop
	_ = dictApplicative1_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_1, "pure"), gopurs_runtime.Float(1.0)))
}

func Call_Main_test6(dictApplicative_0_loop gopurs_runtime.Value, dictPartial_1_loop gopurs_runtime.Value, mx_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	var dictPartial_1 gopurs_runtime.Value = dictPartial_1_loop
	_ = dictPartial_1
	var mx_2 gopurs_runtime.Value = mx_2_loop
	_ = mx_2
	var v_3 gopurs_runtime.Value = v_3_loop
	_ = v_3
	var __t1 gopurs_runtime.Value
	{
		var __t_tag_0 *Constructor_Main_Just = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](mx_2)
		if __t_tag_0 != nil {
			__t1 = (*Constructor_Main_Just)(mx_2.UnsafePtr).V0
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), __t1)
}

func Call_Main_test5(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
	Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
	_ = Functor0_1_0
	return gopurs_runtime.Func(func(mx_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(my_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(mz_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
						// TAST (Let): sum_7_1 -> float64
						sum_7_1 := (v_5.FloatVal()) + (v1_6.FloatVal())
						_ = sum_7_1
						return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Float(((v2_8.FloatVal()) + (sum_7_1)) + (1.0))
						})
					})
				}), mx_2), my_3), mz_4)
			})
		})
	})
}

func Call_Main_test4(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
	Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
	_ = Functor0_1_0
	return gopurs_runtime.Func(func(mx_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(my_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Float(((v_4.FloatVal()) + (v1_5.FloatVal())) + (1.0))
				})
			}), mx_2), my_3)
		})
	})
}

func Call_Main_test11(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
	Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
	_ = Functor0_1_0
	return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(v3_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Str(((gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(v1_4.IntVal)).StrVal()) + (v2_5.StrVal())) + (gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Main_showArray()).V0), func() gopurs_runtime.Value {
							arr := func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(v3_6.UnsafePtr)
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
						}()).StrVal()))
					})
				})
			}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Int(1))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Str("A"))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), func() gopurs_runtime.Value {
				arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
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
			}()))
		})
	})
}

func Call_Main_test10(dictApplicative_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var g_2_0_0 gopurs_runtime.Value
	_ = g_2_0_0
	var f_2_1_1 gopurs_runtime.Value
	_ = f_2_1_1
	g_2_0_0 = gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((gopurs_runtime.Apply(f_2_1_1, gopurs_runtime.Float(x_3.FloatVal())).FloatVal()) / (2.0))
	})
	f_2_1_1 = gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((gopurs_runtime.Apply(g_2_0_0, gopurs_runtime.Float(x_3.FloatVal())).FloatVal()) * (3.0))
	})
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Float(gopurs_runtime.Apply(f_2_1_1, gopurs_runtime.Float(10.0)).FloatVal()))
}

func Call_Main_test1(dictApplicative_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Str("abc"))
}

func Call_Main_test2(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(&Constructor_Main_Just{1, gopurs_runtime.Float(3.0)})})
}

func Call_Main_test3(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just)(nil))})
}

func Call_Main_test9(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(&Constructor_Main_Just{1, gopurs_runtime.Float(3.0)})})
}
