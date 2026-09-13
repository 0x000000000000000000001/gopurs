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
		cache_Main_showArray = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main_showArray
}

var cache_Main_add gopurs_runtime.Value
var once_Main_add sync.Once

func Get_Main_add() gopurs_runtime.Value {
	once_Main_add.Do(func() {
		cache_Main_add = Call_Data_Semiring_add(Rebox_Main_602713622_2826095630(Rebox_Main_2826095630_602713622(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringNumber()))))
	})
	return cache_Main_add
}

var cache_Main_Nothing gopurs_runtime.Value
var once_Main_Nothing sync.Once

func Get_Main_Nothing() gopurs_runtime.Value {
	once_Main_Nothing.Do(func() {
		cache_Main_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Nothing
}

var cache_Main_Just gopurs_runtime.Value
var once_Main_Just sync.Once

func Get_Main_Just() gopurs_runtime.Value {
	once_Main_Just.Do(func() {
		cache_Main_Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just[gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Just
}

var cache_Main_Just__309914989 gopurs_runtime.Value
var once_Main_Just__309914989 sync.Once

func Get_Main_Just__309914989() gopurs_runtime.Value {
	once_Main_Just__309914989.Do(func() {
		cache_Main_Just__309914989 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2858737514_4188394610(Call_Main_Just__309914989(__eta_norm_0_0_box.FloatVal())))}
		})
	})
	return cache_Main_Just__309914989
}

var cache_Main_test8 gopurs_runtime.Value
var once_Main_test8 sync.Once

func Get_Main_test8() gopurs_runtime.Value {
	once_Main_test8.Do(func() {
		cache_Main_test8 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, dictApplicative1_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test8(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative1_1_box), v_2_box)
		})
	})
	return cache_Main_test8
}

var cache_Main_test6 gopurs_runtime.Value
var once_Main_test6 sync.Once

func Get_Main_test6() gopurs_runtime.Value {
	once_Main_test6.Do(func() {
		cache_Main_test6 = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, dictPartial_1_box gopurs_runtime.Value, mx_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test6(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), dictPartial_1_box, gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](mx_2_box), v_3_box)
		})
	})
	return cache_Main_test6
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test5(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box))
		})
	})
	return cache_Main_test5
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test4(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box))
		})
	})
	return cache_Main_test4
}

var cache_Main_test11 gopurs_runtime.Value
var once_Main_test11 sync.Once

func Get_Main_test11() gopurs_runtime.Value {
	once_Main_test11.Do(func() {
		cache_Main_test11 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test11(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box))
		})
	})
	return cache_Main_test11
}

var cache_Main_test10 gopurs_runtime.Value
var once_Main_test10 sync.Once

func Get_Main_test10() gopurs_runtime.Value {
	once_Main_test10.Do(func() {
		cache_Main_test10 = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test10(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), v_1_box)
		})
	})
	return cache_Main_test10
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), v_1_box)
		})
	})
	return cache_Main_test1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Str("X"))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()).V1, gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()).V1, gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()).V1, gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func4(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value, v3_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str((((v1_3.StrVal()) + (v2_4.StrVal())) + ("n")) + (v3_5.StrVal()))
			}), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Str("D"), __local_var_1_1)), gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1)), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("o")
			})), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Str("e")
			})), Get_Effect_Console_log()), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_functorMaybe gopurs_runtime.Value
var once_Main_functorMaybe sync.Once

func Get_Main_functorMaybe() gopurs_runtime.Value {
	once_Main_functorMaybe.Do(func() {
		cache_Main_functorMaybe = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Main_4035997145_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Main_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 *Constructor_Main_Just[gopurs_runtime.Value]
			{
				var __t_tag_0 *Constructor_Main_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](v1_1)
				_ = __t_tag_0
				if __t_tag_0 == nil {
					__t2 = (*Constructor_Main_Just[gopurs_runtime.Value])(nil)
					goto end_branch_2
				} else {

				}
			}
			{
				var __t_tag_1 *Constructor_Main_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](v1_1)
				_ = __t_tag_1
				if __t_tag_1 != nil {
					__t2 = (&Constructor_Main_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Main_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() *Constructor_Main_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_2:
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t2)}
		})})))}
	})
	return cache_Main_functorMaybe
}

var cache_Main_applyMaybe gopurs_runtime.Value
var once_Main_applyMaybe sync.Once

func Get_Main_applyMaybe() gopurs_runtime.Value {
	once_Main_applyMaybe.Do(func() {
		cache_Main_applyMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Main_2381757806_3741347833((&Constructor_Control_Apply_Apply[*Constructor_Main_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Main_4035997145_2812149806(Rebox_Main_2812149806_4035997145(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorMaybe()))))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 *Constructor_Main_Just[gopurs_runtime.Value]
			{
				var __t_tag_0 *Constructor_Main_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](v_0)
				_ = __t_tag_0
				var __t_and_2 bool = false
				if __t_tag_0 != nil {

					var __t_tag_1 *Constructor_Main_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](v1_1)
					_ = __t_tag_1
					__t_and_2 = (__t_tag_1 != nil)
				}
				if __t_and_2 {
					__t3 = (&Constructor_Main_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply((*Constructor_Main_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Main_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = (*Constructor_Main_Just[gopurs_runtime.Value])(nil)
			}
		end_branch_3:
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t3)}
		})})))}
	})
	return cache_Main_applyMaybe
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2858737514_4188394610(Call_Main_test2(v_0_box)))}
		})
	})
	return cache_Main_test2
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2858737514_4188394610(Call_Main_test3(v_0_box)))}
		})
	})
	return cache_Main_test3
}

var cache_Main_test9 gopurs_runtime.Value
var once_Main_test9 sync.Once

func Get_Main_test9() gopurs_runtime.Value {
	once_Main_test9.Do(func() {
		cache_Main_test9 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2858737514_4188394610(Call_Main_test9(v_0_box)))}
		})
	})
	return cache_Main_test9
}

var cache_Main_applicativeMaybe gopurs_runtime.Value
var once_Main_applicativeMaybe sync.Once

func Get_Main_applicativeMaybe() gopurs_runtime.Value {
	once_Main_applicativeMaybe.Do(func() {
		cache_Main_applicativeMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Main_3923698542_1439734649((&Constructor_Control_Applicative_Applicative[*Constructor_Main_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Main_2381757806_3741347833(Rebox_Main_3741347833_2381757806(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyMaybe()))))}
		}), Get_Main_Just()})))}
	})
	return cache_Main_applicativeMaybe
}

type Constructor_Main_Nothing[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Just[T_a any] struct {
	Rc uint32
	V0 T_a
}

func Call_Main_Just__309914989(__eta_norm_0_0_loop float64) *Constructor_Main_Just[float64] {
Just__309914989:
	for {
		if false {
			continue Just__309914989
		}
		var __eta_norm_0_0 float64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return (&Constructor_Main_Just[float64]{1, __eta_norm_0_0})
	}
}

func Call_Main_test8(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictApplicative1_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	var dictApplicative1_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative1_1_loop
	_ = dictApplicative1_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Apply(dictApplicative1_1.V1, gopurs_runtime.Float(1.0)))
}

func Call_Main_test6(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictPartial_1_loop gopurs_runtime.Value, mx_2_loop *Constructor_Main_Just[gopurs_runtime.Value], v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	var dictPartial_1 gopurs_runtime.Value = dictPartial_1_loop
	_ = dictPartial_1
	var mx_2 *Constructor_Main_Just[gopurs_runtime.Value] = mx_2_loop
	_ = mx_2
	var v_3 gopurs_runtime.Value = v_3_loop
	_ = v_3
	var __t0 gopurs_runtime.Value
	{
		if mx_2 != nil {
			__t0 = (mx_2).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return gopurs_runtime.Apply(dictApplicative_0.V1, __t0)
}

func Call_Main_test5(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [Any])
	Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
	_ = Functor0_1_0
	return gopurs_runtime.Func3(func(mx_2 gopurs_runtime.Value, my_3 gopurs_runtime.Value, mz_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): sum_7_1 shape=Other bindingType=Number
			sum_7_1 := (v_5.FloatVal()) + (v1_6.FloatVal())
			_ = sum_7_1
			return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Float(((v2_8.FloatVal()) + (sum_7_1)) + (1.0))
			})
		}), mx_2), my_3), mz_4)
	})
}

func Call_Main_test4(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [Any])
	Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
	_ = Functor0_1_0
	return gopurs_runtime.Func2(func(mx_2 gopurs_runtime.Value, my_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(((v_4.FloatVal()) + (v1_5.FloatVal())) + (1.0))
		}), mx_2), my_3)
	})
}

func Call_Main_test11(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [Any])
	Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
	_ = Functor0_1_0
	return gopurs_runtime.Func2(func(dictApplicative_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func3(func(v1_4 gopurs_runtime.Value, v2_5 gopurs_runtime.Value, v3_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(((gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(v1_4.IntVal)).StrVal()) + (v2_5.StrVal())) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), "show"), func() gopurs_runtime.Value {
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
		}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Int(int64(1)))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Str("A"))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), func() gopurs_runtime.Value {
			arr := []int64{}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()))
	})
}

func Call_Main_test10(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var f__512709115_2_0_0 gopurs_runtime.Value
	_ = f__512709115_2_0_0
	var f__512709115_2_0_0_cell *gopurs_runtime.Value
	_ = f__512709115_2_0_0_cell
	// FALLBACK TCO: isLoop=false len=3
	var g_2_1_1 gopurs_runtime.Value
	_ = g_2_1_1
	var g_2_1_1_cell *gopurs_runtime.Value
	_ = g_2_1_1_cell
	// FALLBACK TCO: isLoop=false len=3
	var f_2_2_2 gopurs_runtime.Value
	_ = f_2_2_2
	var f_2_2_2_cell *gopurs_runtime.Value
	_ = f_2_2_2_cell
	// FALLBACK TCO: isLoop=false len=3
	f__512709115_2_0_0 = gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((gopurs_runtime.Apply((*g_2_1_1_cell), gopurs_runtime.Float(x_3.FloatVal())).FloatVal()) * (3.0))
	})
	f__512709115_2_0_0_cell = &f__512709115_2_0_0
	g_2_1_1 = gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((gopurs_runtime.Apply((*f__512709115_2_0_0_cell), gopurs_runtime.Float(x_3.FloatVal())).FloatVal()) / (2.0))
	})
	g_2_1_1_cell = &g_2_1_1
	f_2_2_2 = gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((gopurs_runtime.Apply((*g_2_1_1_cell), gopurs_runtime.Float(x_3.FloatVal())).FloatVal()) * (3.0))
	})
	f_2_2_2_cell = &f_2_2_2
	return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Float(gopurs_runtime.Apply(f__512709115_2_0_0, gopurs_runtime.Float(10.0)).FloatVal()))
}

func Call_Main_test1(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(dictApplicative_0.V1, gopurs_runtime.Str("abc"))
}

func Call_Main_test2(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just[float64] {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return Rebox_Main_4188394610_2858737514((&Constructor_Main_Just[gopurs_runtime.Value]{1, gopurs_runtime.Float(3.0)}))
}

func Call_Main_test3(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just[float64] {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return Rebox_Main_4188394610_2858737514((*Constructor_Main_Just[gopurs_runtime.Value])(nil))
}

func Call_Main_test9(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just[float64] {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return Rebox_Main_4188394610_2858737514((&Constructor_Main_Just[gopurs_runtime.Value]{1, gopurs_runtime.Float(gopurs_runtime.Apply2(Call_Data_Semiring_add(Rebox_Main_602713622_2826095630(Rebox_Main_2826095630_602713622(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringNumber())))), gopurs_runtime.Float(1.0), gopurs_runtime.Float(2.0)).FloatVal())}))
}

func Rebox_Main_1386611502_1469227923(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[[]int64]{}
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

func Rebox_Main_1469227923_1386611502(in *Constructor_Data_Show_Show[[]int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
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

func Rebox_Main_2381757806_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Main_Just[gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Apply_Apply[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_2812149806_4035997145(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Main_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Functor_Functor[*Constructor_Main_Just[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2826095630_602713622(in *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) *Constructor_Data_Semiring_Semiring[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[float64]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2.FloatVal()
	out.V3 = in.V3.FloatVal()
	return out
}

func Rebox_Main_2858737514_4188394610(in *Constructor_Main_Just[float64]) *Constructor_Main_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Float(in.V0)
	return out
}

func Rebox_Main_3741347833_2381757806(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Main_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Apply_Apply[*Constructor_Main_Just[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_3923698542_1439734649(in *Constructor_Control_Applicative_Applicative[*Constructor_Main_Just[gopurs_runtime.Value]]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_4035997145_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Main_Just[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_4188394610_2858737514(in *Constructor_Main_Just[gopurs_runtime.Value]) *Constructor_Main_Just[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Just[float64]{}
	out.V0 = in.V0.FloatVal()
	return out
}

func Rebox_Main_602713622_2826095630(in *Constructor_Data_Semiring_Semiring[float64]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Float(in.V2)
	out.V3 = gopurs_runtime.Float(in.V3)
	return out
}
