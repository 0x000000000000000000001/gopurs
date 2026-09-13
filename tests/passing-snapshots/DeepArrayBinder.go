package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Cons gopurs_runtime.Value
var once_Main_Cons sync.Once

func Get_Main_Cons() gopurs_runtime.Value {
	once_Main_Cons.Do(func() {
		cache_Main_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](value1)}))}
			})
		})
	})
	return cache_Main_Cons
}

var cache_Main_Nil gopurs_runtime.Value
var once_Main_Nil sync.Once

func Get_Main_Nil() gopurs_runtime.Value {
	once_Main_Nil.Do(func() {
		cache_Main_Nil = gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Nil
}

var cache_Main_Cons__2372756309 gopurs_runtime.Value
var once_Main_Cons__2372756309 sync.Once

func Get_Main_Cons__2372756309() gopurs_runtime.Value {
	once_Main_Cons__2372756309.Do(func() {
		cache_Main_Cons__2372756309 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Rebox_Main_332595043_176455803(Call_Main_Cons__2372756309(__eta_norm_1_0_box.FloatVal(), Rebox_Main_176455803_332595043(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](__eta_norm_0_unused_1_box)))))}
		})
	})
	return cache_Main_Cons__2372756309
}

var cache_Main_Cons__3816333300 gopurs_runtime.Value
var once_Main_Cons__3816333300 sync.Once

func Get_Main_Cons__3816333300() gopurs_runtime.Value {
	once_Main_Cons__3816333300.Do(func() {
		cache_Main_Cons__3816333300 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Rebox_Main_332595043_176455803(Call_Main_Cons__3816333300(__eta_norm_1_0_box.FloatVal(), Rebox_Main_176455803_332595043(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](__eta_norm_0_1_box)))))}
		})
	})
	return cache_Main_Cons__3816333300
}

var cache_Main_match2 gopurs_runtime.Value
var once_Main_match2 sync.Once

func Get_Main_match2() gopurs_runtime.Value {
	once_Main_match2.Do(func() {
		cache_Main_match2 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_match2(Rebox_Main_176455803_332595043(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](v_0_box))))
		})
	})
	return cache_Main_match2
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("Incorrect result!"), gopurs_runtime.Bool((Call_Main_match2((&Constructor_Main_Cons[float64]{1, 1.0, (&Constructor_Main_Cons[float64]{1, 2.0, (&Constructor_Main_Cons[float64]{1, 3.0, (&Constructor_Main_Cons[float64]{1, 4.0, (&Constructor_Main_Cons[float64]{1, 5.0, (&Constructor_Main_Cons[float64]{1, 6.0, (&Constructor_Main_Cons[float64]{1, 7.0, (&Constructor_Main_Cons[float64]{1, 8.0, (&Constructor_Main_Cons[float64]{1, 9.0, (*Constructor_Main_Cons[float64])(nil)})})})})})})})})}))) == (100.0))), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

type Constructor_Main_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Main_Cons[T_a]
}

type Constructor_Main_Nil[T_a any] struct {
	Rc uint32
}

func Call_Main_Cons__2372756309(__eta_norm_1_0_loop float64, __eta_norm_0_unused_1_loop *Constructor_Main_Cons[float64]) *Constructor_Main_Cons[float64] {
Cons__2372756309:
	for {
		if false {
			continue Cons__2372756309
		}
		var __eta_norm_1_0 float64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 *Constructor_Main_Cons[float64] = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return (&Constructor_Main_Cons[float64]{1, __eta_norm_1_0, (*Constructor_Main_Cons[float64])(nil)})
	}
}

func Call_Main_Cons__3816333300(__eta_norm_1_0_loop float64, __eta_norm_0_1_loop *Constructor_Main_Cons[float64]) *Constructor_Main_Cons[float64] {
Cons__3816333300:
	for {
		if false {
			continue Cons__3816333300
		}
		var __eta_norm_1_0 float64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 *Constructor_Main_Cons[float64] = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Main_Cons[float64]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Main_match2(v_0_loop *Constructor_Main_Cons[float64]) float64 {
match2:
	for {
		if false {
			continue match2
		}
		var v_0 *Constructor_Main_Cons[float64] = v_0_loop
		_ = v_0
		var __t2 float64
		{
			var __t_and_1 bool = false
			if v_0 != nil {

				var __t_tag_0 *Constructor_Main_Cons[float64] = (v_0).V1
				_ = __t_tag_0
				__t_and_1 = (__t_tag_0 != nil)
			}
			if __t_and_1 {
				__t2 = (((v_0).V0) * (((v_0).V1).V0)) + (Call_Main_match2(((v_0).V1).V1))
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = 0.0
		}
	end_branch_2:
		return __t2
	}
}

func Rebox_Main_176455803_332595043(in *Constructor_Main_Cons[gopurs_runtime.Value]) *Constructor_Main_Cons[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Cons[float64]{}
	out.V0 = in.V0.FloatVal()
	out.V1 = Rebox_Main_176455803_332595043(in.V1)
	return out
}

func Rebox_Main_332595043_176455803(in *Constructor_Main_Cons[float64]) *Constructor_Main_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Float(in.V0)
	out.V1 = Rebox_Main_332595043_176455803(in.V1)
	return out
}
