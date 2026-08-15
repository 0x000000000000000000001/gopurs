package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just{1, value0}))}
		})
	})
	return cache_Main_Just
}

var cache_Main_test8 gopurs_runtime.Value
var once_Main_test8 sync.Once

func Get_Main_test8() gopurs_runtime.Value {
	once_Main_test8.Do(func() {
		cache_Main_test8 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Call_Main_test8(v_0_box))}
		})
	})
	return cache_Main_test8
}

var cache_Main_test6 gopurs_runtime.Value
var once_Main_test6 sync.Once

func Get_Main_test6() gopurs_runtime.Value {
	once_Main_test6.Do(func() {
		cache_Main_test6 = gopurs_runtime.Func3(func(dictPartial_0_box gopurs_runtime.Value, mx_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test6(dictPartial_0_box, mx_1_box, v_2_box)
		})
	})
	return cache_Main_test6
}

var cache_Main_test10 gopurs_runtime.Value
var once_Main_test10 sync.Once

func Get_Main_test10() gopurs_runtime.Value {
	once_Main_test10.Do(func() {
		cache_Main_test10 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Call_Main_test10(v_0_box))}
		})
	})
	return cache_Main_test10
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Call_Main_test1(v_0_box))}
		})
	})
	return cache_Main_test1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_functorMaybe gopurs_runtime.Value
var once_Main_functorMaybe sync.Once

func Get_Main_functorMaybe() gopurs_runtime.Value {
	once_Main_functorMaybe.Do(func() {
		cache_Main_functorMaybe = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
						__t0 = (&Constructor_Main_Just{1, gopurs_runtime.Apply(v_0, (*Constructor_Main_Just)(v1_1.UnsafePtr).V0)})
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
		})}))}
	})
	return cache_Main_functorMaybe
}

var cache_Main_applyMaybe gopurs_runtime.Value
var once_Main_applyMaybe sync.Once

func Get_Main_applyMaybe() gopurs_runtime.Value {
	once_Main_applyMaybe.Do(func() {
		cache_Main_applyMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Main_functorMaybe()))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 *Constructor_Main_Just
				{
					if (v_0.Type == 9 && v_0.IntVal == 3271839782 && v_0.UnsafePtr != nil) && (v1_1.Type == 9 && v1_1.IntVal == 3271839782 && v1_1.UnsafePtr != nil) {
						__t0 = (&Constructor_Main_Just{1, gopurs_runtime.Apply((*Constructor_Main_Just)(v_0.UnsafePtr).V0, (*Constructor_Main_Just)(v1_1.UnsafePtr).V0)})
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
		})}))}
	})
	return cache_Main_applyMaybe
}

var cache_Main_bindMaybe gopurs_runtime.Value
var once_Main_bindMaybe sync.Once

func Get_Main_bindMaybe() gopurs_runtime.Value {
	once_Main_bindMaybe.Do(func() {
		cache_Main_bindMaybe = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Main_applyMaybe()))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 *Constructor_Main_Just
				{
					if v_0.Type == 9 && v_0.IntVal == 3271839782 && v_0.UnsafePtr == nil {
						__t0 = (*Constructor_Main_Just)(nil)
						goto end_branch_0
					} else {

					}
				}
				{
					if v_0.Type == 9 && v_0.IntVal == 3271839782 && v_0.UnsafePtr != nil {
						__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Apply(v1_1, (*Constructor_Main_Just)(v_0.UnsafePtr).V0))
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
		})}))}
	})
	return cache_Main_bindMaybe
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

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Func2(func(mx_0_box gopurs_runtime.Value, my_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Call_Main_test4(gopurs_runtime.CoerceToStruct[Constructor_Main_Just](mx_0_box), gopurs_runtime.CoerceToStruct[Constructor_Main_Just](my_1_box)))}
		})
	})
	return cache_Main_test4
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.Func3(func(mx_0_box gopurs_runtime.Value, my_1_box gopurs_runtime.Value, mz_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Call_Main_test5(gopurs_runtime.CoerceToStruct[Constructor_Main_Just](mx_0_box), gopurs_runtime.CoerceToStruct[Constructor_Main_Just](my_1_box), gopurs_runtime.CoerceToStruct[Constructor_Main_Just](mz_2_box)))}
		})
	})
	return cache_Main_test5
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
		cache_Main_applicativeMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Main_applyMaybe()))}
		}), Get_Main_Just()}))}
	})
	return cache_Main_applicativeMaybe
}

var cache_Main_monadMaybe gopurs_runtime.Value
var once_Main_monadMaybe sync.Once

func Get_Main_monadMaybe() gopurs_runtime.Value {
	once_Main_monadMaybe.Do(func() {
		cache_Main_monadMaybe = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Main_applicativeMaybe()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Main_bindMaybe()))}
		})}))}
	})
	return cache_Main_monadMaybe
}

type Constructor_Main_Nothing struct {
	Rc uint32
}

type Constructor_Main_Just struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_test8(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just{1, gopurs_runtime.Float(1.0)}))}}))})
}

func Call_Main_test6(dictPartial_0_loop gopurs_runtime.Value, mx_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictPartial_0 gopurs_runtime.Value = dictPartial_0_loop
	_ = dictPartial_0
	var mx_1 gopurs_runtime.Value = mx_1_loop
	_ = mx_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	var __t1 gopurs_runtime.Value
	{
		var __t_tag_0 *Constructor_Main_Just = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](mx_1)
		if __t_tag_0 != nil {
			__t1 = (*Constructor_Main_Just)(mx_1.UnsafePtr).V0
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_1:
	return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just{1, __t1}))}
}

func Call_Main_test10(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var g_1_0_0 gopurs_runtime.Value
	_ = g_1_0_0
	var f_1_1_1 gopurs_runtime.Value
	_ = f_1_1_1
	g_1_0_0 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((gopurs_runtime.Apply(f_1_1_1, gopurs_runtime.Float(x_2.FloatVal())).FloatVal()) / (2.0))
	})
	f_1_1_1 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((gopurs_runtime.Apply(g_1_0_0, gopurs_runtime.Float(x_2.FloatVal())).FloatVal()) * (3.0))
	})
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just{1, gopurs_runtime.Apply(f_1_1_1, gopurs_runtime.Float(10.0))}))})
}

func Call_Main_test1(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just{1, gopurs_runtime.Str("abc")}))})
}

func Call_Main_test2(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just{1, gopurs_runtime.Float(3.0)}))})
}

func Call_Main_test3(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just)(nil))})
}

func Call_Main_test4(mx_0_loop *Constructor_Main_Just, my_1_loop *Constructor_Main_Just) *Constructor_Main_Just {
	var mx_0 *Constructor_Main_Just = mx_0_loop
	_ = mx_0
	var my_1 *Constructor_Main_Just = my_1_loop
	_ = my_1
	var __t1 *Constructor_Main_Just
	{
		if mx_0 == nil {
			__t1 = (*Constructor_Main_Just)(nil)
			goto end_branch_1
		} else {

		}
	}
	{
		if mx_0 != nil {
			var __t0 *Constructor_Main_Just
			{
				if my_1 == nil {
					__t0 = (*Constructor_Main_Just)(nil)
					goto end_branch_0
				} else {

				}
			}
			{
				if my_1 != nil {
					__t0 = (&Constructor_Main_Just{1, gopurs_runtime.Float((((mx_0).V0.FloatVal()) + ((my_1).V0.FloatVal())) + (1.0))})
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
			}
		end_branch_0:
			__t1 = __t0
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
	}
end_branch_1:
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Main_test5(mx_0_loop *Constructor_Main_Just, my_1_loop *Constructor_Main_Just, mz_2_loop *Constructor_Main_Just) *Constructor_Main_Just {
	var mx_0 *Constructor_Main_Just = mx_0_loop
	_ = mx_0
	var my_1 *Constructor_Main_Just = my_1_loop
	_ = my_1
	var mz_2 *Constructor_Main_Just = mz_2_loop
	_ = mz_2
	var __t3 *Constructor_Main_Just
	{
		if mx_0 == nil {
			__t3 = (*Constructor_Main_Just)(nil)
			goto end_branch_3
		} else {

		}
	}
	{
		if mx_0 != nil {
			var __t2 *Constructor_Main_Just
			{
				if my_1 == nil {
					__t2 = (*Constructor_Main_Just)(nil)
					goto end_branch_2
				} else {

				}
			}
			{
				if my_1 != nil {
					// TAST (Let): sum_3_0 -> float64
					sum_3_0 := ((mx_0).V0.FloatVal()) + ((my_1).V0.FloatVal())
					_ = sum_3_0
					var __t1 *Constructor_Main_Just
					{
						if mz_2 == nil {
							__t1 = (*Constructor_Main_Just)(nil)
							goto end_branch_1
						} else {

						}
					}
					{
						if mz_2 != nil {
							__t1 = (&Constructor_Main_Just{1, gopurs_runtime.Float((((mz_2).V0.FloatVal()) + (sum_3_0)) + (1.0))})
							goto end_branch_1
						} else {

						}
					}
					{
						__t1 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
					}
				end_branch_1:
					__t2 = __t1
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
			}
		end_branch_2:
			__t3 = __t2
			goto end_branch_3
		} else {

		}
	}
	{
		__t3 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
	}
end_branch_3:
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t3)})
}

func Call_Main_test9(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Just](gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((&Constructor_Main_Just{1, gopurs_runtime.Float(3.0)}))})
}
