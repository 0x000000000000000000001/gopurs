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

var cache_Main_Just__1900459053 gopurs_runtime.Value
var once_Main_Just__1900459053 sync.Once

func Get_Main_Just__1900459053() gopurs_runtime.Value {
	once_Main_Just__1900459053.Do(func() {
		cache_Main_Just__1900459053 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_859947357_4188394610(Call_Main_Just__1900459053(Rebox_Main_4188394610_2858737514(gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](__eta_norm_0_0_box)))))}
		})
	})
	return cache_Main_Just__1900459053
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

var cache_Main_Just__3163753069 gopurs_runtime.Value
var once_Main_Just__3163753069 sync.Once

func Get_Main_Just__3163753069() gopurs_runtime.Value {
	once_Main_Just__3163753069.Do(func() {
		cache_Main_Just__3163753069 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2431535661_4188394610(Call_Main_Just__3163753069(__eta_norm_0_0_box.StrVal())))}
		})
	})
	return cache_Main_Just__3163753069
}

var cache_Main_test8 gopurs_runtime.Value
var once_Main_test8 sync.Once

func Get_Main_test8() gopurs_runtime.Value {
	once_Main_test8.Do(func() {
		cache_Main_test8 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_859947357_4188394610(Call_Main_test8(v_0_box)))}
		})
	})
	return cache_Main_test8
}

var cache_Main_test6 gopurs_runtime.Value
var once_Main_test6 sync.Once

func Get_Main_test6() gopurs_runtime.Value {
	once_Main_test6.Do(func() {
		cache_Main_test6 = gopurs_runtime.Func3(func(dictPartial_0_box gopurs_runtime.Value, mx_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Call_Main_test6(dictPartial_0_box, gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](mx_1_box), v_2_box))}
		})
	})
	return cache_Main_test6
}

var cache_Main_test10 gopurs_runtime.Value
var once_Main_test10 sync.Once

func Get_Main_test10() gopurs_runtime.Value {
	once_Main_test10.Do(func() {
		cache_Main_test10 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2858737514_4188394610(Call_Main_test10(v_0_box)))}
		})
	})
	return cache_Main_test10
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2431535661_4188394610(Call_Main_test1(v_0_box)))}
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

var cache_Main_bindMaybe gopurs_runtime.Value
var once_Main_bindMaybe sync.Once

func Get_Main_bindMaybe() gopurs_runtime.Value {
	once_Main_bindMaybe.Do(func() {
		cache_Main_bindMaybe = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Main_2534018798_2748095225((&Constructor_Control_Bind_Bind[*Constructor_Main_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Main_2381757806_3741347833(Rebox_Main_3741347833_2381757806(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Main_applyMaybe()))))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 *Constructor_Main_Just[gopurs_runtime.Value]
			{
				var __t_tag_0 *Constructor_Main_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](v_0)
				_ = __t_tag_0
				if __t_tag_0 == nil {
					__t2 = (*Constructor_Main_Just[gopurs_runtime.Value])(nil)
					goto end_branch_2
				} else {

				}
			}
			{
				var __t_tag_1 *Constructor_Main_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](v_0)
				_ = __t_tag_1
				if __t_tag_1 != nil {
					__t2 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Main_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0))
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
	return cache_Main_bindMaybe
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

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Func2(func(mx_0_box gopurs_runtime.Value, my_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2858737514_4188394610(Call_Main_test4(Rebox_Main_4188394610_2858737514(gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](mx_0_box)), Rebox_Main_4188394610_2858737514(gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](my_1_box)))))}
		})
	})
	return cache_Main_test4
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.Func3(func(mx_0_box gopurs_runtime.Value, my_1_box gopurs_runtime.Value, mz_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2858737514_4188394610(Call_Main_test5(Rebox_Main_4188394610_2858737514(gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](mx_0_box)), Rebox_Main_4188394610_2858737514(gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](my_1_box)), Rebox_Main_4188394610_2858737514(gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](mz_2_box)))))}
		})
	})
	return cache_Main_test5
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

var cache_Main_monadMaybe gopurs_runtime.Value
var once_Main_monadMaybe sync.Once

func Get_Main_monadMaybe() gopurs_runtime.Value {
	once_Main_monadMaybe.Do(func() {
		cache_Main_monadMaybe = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Main_107968622_2568689657((&Constructor_Control_Monad_Monad[*Constructor_Main_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Main_3923698542_1439734649(Rebox_Main_1439734649_3923698542(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Main_applicativeMaybe()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Main_2534018798_2748095225(Rebox_Main_2748095225_2534018798(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Main_bindMaybe()))))}
		})})))}
	})
	return cache_Main_monadMaybe
}

type Constructor_Main_Nothing[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Just[T_a any] struct {
	Rc uint32
	V0 T_a
}

func Call_Main_Just__1900459053(__eta_norm_0_0_loop *Constructor_Main_Just[float64]) *Constructor_Main_Just[*Constructor_Main_Just[float64]] {
Just__1900459053:
	for {
		if false {
			continue Just__1900459053
		}
		var __eta_norm_0_0 *Constructor_Main_Just[float64] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return (&Constructor_Main_Just[*Constructor_Main_Just[float64]]{1, __eta_norm_0_0})
	}
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

func Call_Main_Just__3163753069(__eta_norm_0_0_loop string) *Constructor_Main_Just[string] {
Just__3163753069:
	for {
		if false {
			continue Just__3163753069
		}
		var __eta_norm_0_0 string = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return (&Constructor_Main_Just[string]{1, __eta_norm_0_0})
	}
}

func Call_Main_test8(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just[*Constructor_Main_Just[float64]] {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return (&Constructor_Main_Just[*Constructor_Main_Just[float64]]{1, (&Constructor_Main_Just[float64]{1, 1.0})})
}

func Call_Main_test6(dictPartial_0_loop gopurs_runtime.Value, mx_1_loop *Constructor_Main_Just[gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) *Constructor_Main_Just[gopurs_runtime.Value] {
	var dictPartial_0 gopurs_runtime.Value = dictPartial_0_loop
	_ = dictPartial_0
	var mx_1 *Constructor_Main_Just[gopurs_runtime.Value] = mx_1_loop
	_ = mx_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	var __t0 gopurs_runtime.Value
	{
		if mx_1 != nil {
			__t0 = (mx_1).V0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return (&Constructor_Main_Just[gopurs_runtime.Value]{1, __t0})
}

func Call_Main_test10(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just[float64] {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var f__512709115_1_0_0 gopurs_runtime.Value
	_ = f__512709115_1_0_0
	var f__512709115_1_0_0_cell *gopurs_runtime.Value
	_ = f__512709115_1_0_0_cell
	// FALLBACK TCO: isLoop=false len=3
	var g_1_1_1 gopurs_runtime.Value
	_ = g_1_1_1
	var g_1_1_1_cell *gopurs_runtime.Value
	_ = g_1_1_1_cell
	// FALLBACK TCO: isLoop=false len=3
	var f_1_2_2 gopurs_runtime.Value
	_ = f_1_2_2
	var f_1_2_2_cell *gopurs_runtime.Value
	_ = f_1_2_2_cell
	// FALLBACK TCO: isLoop=false len=3
	f__512709115_1_0_0 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((gopurs_runtime.Apply((*g_1_1_1_cell), gopurs_runtime.Float(x_2.FloatVal())).FloatVal()) * (3.0))
	})
	f__512709115_1_0_0_cell = &f__512709115_1_0_0
	g_1_1_1 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((gopurs_runtime.Apply((*f__512709115_1_0_0_cell), gopurs_runtime.Float(x_2.FloatVal())).FloatVal()) / (2.0))
	})
	g_1_1_1_cell = &g_1_1_1
	f_1_2_2 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float((gopurs_runtime.Apply((*g_1_1_1_cell), gopurs_runtime.Float(x_2.FloatVal())).FloatVal()) * (3.0))
	})
	f_1_2_2_cell = &f_1_2_2
	return (&Constructor_Main_Just[float64]{1, gopurs_runtime.Apply(f__512709115_1_0_0, gopurs_runtime.Float(10.0)).FloatVal()})
}

func Call_Main_test1(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just[string] {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return (&Constructor_Main_Just[string]{1, "abc"})
}

func Call_Main_test2(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just[float64] {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return (&Constructor_Main_Just[float64]{1, 3.0})
}

func Call_Main_test3(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just[float64] {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return Rebox_Main_4188394610_2858737514((*Constructor_Main_Just[gopurs_runtime.Value])(nil))
}

func Call_Main_test4(mx_0_loop *Constructor_Main_Just[float64], my_1_loop *Constructor_Main_Just[float64]) *Constructor_Main_Just[float64] {
	var mx_0 *Constructor_Main_Just[float64] = mx_0_loop
	_ = mx_0
	var my_1 *Constructor_Main_Just[float64] = my_1_loop
	_ = my_1
	var __t1 *Constructor_Main_Just[gopurs_runtime.Value]
	{
		if mx_0 == nil {
			__t1 = (*Constructor_Main_Just[gopurs_runtime.Value])(nil)
			goto end_branch_1
		} else {

		}
	}
	{
		if mx_0 != nil {
			var __t0 gopurs_runtime.Value
			{
				if my_1 == nil {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just[gopurs_runtime.Value])(nil))}
					goto end_branch_0
				} else {

				}
			}
			{
				if my_1 != nil {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2858737514_4188394610((&Constructor_Main_Just[float64]{1, (((mx_0).V0) + ((my_1).V0)) + (1.0)})))}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(func() *Constructor_Main_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
			}
		end_branch_0:
			__t1 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](__t0)
			goto end_branch_1
		} else {

		}
	}
	{
		__t1 = func() *Constructor_Main_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
	}
end_branch_1:
	return Rebox_Main_4188394610_2858737514(__t1)
}

func Call_Main_test5(mx_0_loop *Constructor_Main_Just[float64], my_1_loop *Constructor_Main_Just[float64], mz_2_loop *Constructor_Main_Just[float64]) *Constructor_Main_Just[float64] {
	var mx_0 *Constructor_Main_Just[float64] = mx_0_loop
	_ = mx_0
	var my_1 *Constructor_Main_Just[float64] = my_1_loop
	_ = my_1
	var mz_2 *Constructor_Main_Just[float64] = mz_2_loop
	_ = mz_2
	var __t3 *Constructor_Main_Just[gopurs_runtime.Value]
	{
		if mx_0 == nil {
			__t3 = (*Constructor_Main_Just[gopurs_runtime.Value])(nil)
			goto end_branch_3
		} else {

		}
	}
	{
		if mx_0 != nil {
			var __t2 *Constructor_Main_Just[gopurs_runtime.Value]
			{
				if my_1 == nil {
					__t2 = (*Constructor_Main_Just[gopurs_runtime.Value])(nil)
					goto end_branch_2
				} else {

				}
			}
			{
				if my_1 != nil {
					// TAST (Let): sum_3_0 shape=Other bindingType=Number
					sum_3_0 := ((mx_0).V0) + ((my_1).V0)
					_ = sum_3_0
					var __t1 gopurs_runtime.Value
					{
						if mz_2 == nil {
							__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just[gopurs_runtime.Value])(nil))}
							goto end_branch_1
						} else {

						}
					}
					{
						if mz_2 != nil {
							__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2858737514_4188394610((&Constructor_Main_Just[float64]{1, (((mz_2).V0) + (sum_3_0)) + (1.0)})))}
							goto end_branch_1
						} else {

						}
					}
					{
						__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(func() *Constructor_Main_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
					}
				end_branch_1:
					__t2 = gopurs_runtime.CoerceToStruct[Constructor_Main_Just[gopurs_runtime.Value]](__t1)
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() *Constructor_Main_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_2:
			__t3 = __t2
			goto end_branch_3
		} else {

		}
	}
	{
		__t3 = func() *Constructor_Main_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
	}
end_branch_3:
	return Rebox_Main_4188394610_2858737514(__t3)
}

func Call_Main_test9(v_0_loop gopurs_runtime.Value) *Constructor_Main_Just[float64] {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return Rebox_Main_4188394610_2858737514((&Constructor_Main_Just[gopurs_runtime.Value]{1, gopurs_runtime.Float(gopurs_runtime.Apply2(Call_Data_Semiring_add(Rebox_Main_602713622_2826095630(Rebox_Main_2826095630_602713622(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringNumber())))), gopurs_runtime.Float(1.0), gopurs_runtime.Float(2.0)).FloatVal())}))
}

func Rebox_Main_107968622_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Main_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Monad_Monad[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_1439734649_3923698542(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Main_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Applicative_Applicative[*Constructor_Main_Just[gopurs_runtime.Value]]{}
	out.V0 = in.V0
	out.V1 = in.V1
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

func Rebox_Main_2431535661_4188394610(in *Constructor_Main_Just[string]) *Constructor_Main_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}

func Rebox_Main_2534018798_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Main_Just[gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_2748095225_2534018798(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Main_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Control_Bind_Bind[*Constructor_Main_Just[gopurs_runtime.Value]]{}
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

func Rebox_Main_859947357_4188394610(in *Constructor_Main_Just[*Constructor_Main_Just[float64]]) *Constructor_Main_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(Rebox_Main_2858737514_4188394610(in.V0))}
	return out
}
