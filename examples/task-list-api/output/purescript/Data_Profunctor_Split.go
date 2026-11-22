package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_Split_compose gopurs_runtime.Value
var once_Data_Profunctor_Split_compose sync.Once

func Get_Data_Profunctor_Split_compose() gopurs_runtime.Value {
	once_Data_Profunctor_Split_compose.Do(func() {
		cache_Data_Profunctor_Split_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Profunctor_Split_compose
}

var cache_Data_Profunctor_Split_compose1 gopurs_runtime.Value
var once_Data_Profunctor_Split_compose1 sync.Once

func Get_Data_Profunctor_Split_compose1() gopurs_runtime.Value {
	once_Data_Profunctor_Split_compose1.Do(func() {
		cache_Data_Profunctor_Split_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Profunctor_Split_compose1
}

var cache_Data_Profunctor_Split_identity gopurs_runtime.Value
var once_Data_Profunctor_Split_identity sync.Once

func Get_Data_Profunctor_Split_identity() gopurs_runtime.Value {
	once_Data_Profunctor_Split_identity.Do(func() {
		cache_Data_Profunctor_Split_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Split_identity(x_0_box)
		})
	})
	return cache_Data_Profunctor_Split_identity
}

var cache_Data_Profunctor_Split_compose2 gopurs_runtime.Value
var once_Data_Profunctor_Split_compose2 sync.Once

func Get_Data_Profunctor_Split_compose2() gopurs_runtime.Value {
	once_Data_Profunctor_Split_compose2.Do(func() {
		cache_Data_Profunctor_Split_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Profunctor_Split_compose2
}

var cache_Data_Profunctor_Split_compose3 gopurs_runtime.Value
var once_Data_Profunctor_Split_compose3 sync.Once

func Get_Data_Profunctor_Split_compose3() gopurs_runtime.Value {
	once_Data_Profunctor_Split_compose3.Do(func() {
		cache_Data_Profunctor_Split_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Profunctor_Split_compose3
}

var cache_Data_Profunctor_Split_SplitF gopurs_runtime.Value
var once_Data_Profunctor_Split_SplitF sync.Once

func Get_Data_Profunctor_Split_SplitF() gopurs_runtime.Value {
	once_Data_Profunctor_Split_SplitF.Do(func() {
		cache_Data_Profunctor_Split_SplitF = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, value2}))}
				})
			})
		})
	})
	return cache_Data_Profunctor_Split_SplitF
}

var cache_Data_Profunctor_Split_Split gopurs_runtime.Value
var once_Data_Profunctor_Split_Split sync.Once

func Get_Data_Profunctor_Split_Split() gopurs_runtime.Value {
	once_Data_Profunctor_Split_Split.Do(func() {
		cache_Data_Profunctor_Split_Split = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Split_Split(x_0_box)
		})
	})
	return cache_Data_Profunctor_Split_Split
}

var cache_Data_Profunctor_Split_unSplit gopurs_runtime.Value
var once_Data_Profunctor_Split_unSplit sync.Once

func Get_Data_Profunctor_Split_unSplit() gopurs_runtime.Value {
	once_Data_Profunctor_Split_unSplit.Do(func() {
		cache_Data_Profunctor_Split_unSplit = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Split_unSplit(f_0_box, v_1_box)
		})
	})
	return cache_Data_Profunctor_Split_unSplit
}

var cache_Data_Profunctor_Split_split gopurs_runtime.Value
var once_Data_Profunctor_Split_split sync.Once

func Get_Data_Profunctor_Split_split() gopurs_runtime.Value {
	once_Data_Profunctor_Split_split.Do(func() {
		cache_Data_Profunctor_Split_split = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, fx_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Split_split(f_0_box, g_1_box, fx_2_box)
		})
	})
	return cache_Data_Profunctor_Split_split
}

var cache_Data_Profunctor_Split_profunctorSplit gopurs_runtime.Value
var once_Data_Profunctor_Split_profunctorSplit sync.Once

func Get_Data_Profunctor_Split_profunctorSplit() gopurs_runtime.Value {
	once_Data_Profunctor_Split_profunctorSplit.Do(func() {
		cache_Data_Profunctor_Split_profunctorSplit = gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.Func2(func(h_3 gopurs_runtime.Value, i_4 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_5_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope34)] (TypeVar x$scope38))
				__local_var_5_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), h_3, f_0)
				_ = __local_var_5_0
				// TAST (Let): __local_var_6_1 shape=App(Var) bindingType=(Func [(TypeVar x$scope38)] (TypeVar d$scope37))
				__local_var_6_1 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), g_1, i_4)
				_ = __local_var_6_1
				return gopurs_runtime.Func(func(fx_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_5_0, __local_var_6_1, fx_7}))}
				})
			}), (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
		})}))}
	})
	return cache_Data_Profunctor_Split_profunctorSplit
}

var cache_Data_Profunctor_Split_lowerSplit gopurs_runtime.Value
var once_Data_Profunctor_Split_lowerSplit sync.Once

func Get_Data_Profunctor_Split_lowerSplit() gopurs_runtime.Value {
	once_Data_Profunctor_Split_lowerSplit.Do(func() {
		cache_Data_Profunctor_Split_lowerSplit = gopurs_runtime.Func(func(dictInvariant_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Split_lowerSplit(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]](dictInvariant_0_box))
		})
	})
	return cache_Data_Profunctor_Split_lowerSplit
}

var cache_Data_Profunctor_Split_liftSplit gopurs_runtime.Value
var once_Data_Profunctor_Split_liftSplit sync.Once

func Get_Data_Profunctor_Split_liftSplit() gopurs_runtime.Value {
	once_Data_Profunctor_Split_liftSplit.Do(func() {
		cache_Data_Profunctor_Split_liftSplit = gopurs_runtime.Func(func(fx_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Split_liftSplit(fx_0_box)
		})
	})
	return cache_Data_Profunctor_Split_liftSplit
}

var cache_Data_Profunctor_Split_hoistSplit gopurs_runtime.Value
var once_Data_Profunctor_Split_hoistSplit sync.Once

func Get_Data_Profunctor_Split_hoistSplit() gopurs_runtime.Value {
	once_Data_Profunctor_Split_hoistSplit.Do(func() {
		cache_Data_Profunctor_Split_hoistSplit = gopurs_runtime.Func2(func(nat_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Split_hoistSplit(nat_0_box, v_1_box)
		})
	})
	return cache_Data_Profunctor_Split_hoistSplit
}

var cache_Data_Profunctor_Split_functorSplit gopurs_runtime.Value
var once_Data_Profunctor_Split_functorSplit sync.Once

func Get_Data_Profunctor_Split_functorSplit() gopurs_runtime.Value {
	once_Data_Profunctor_Split_functorSplit.Do(func() {
		cache_Data_Profunctor_Split_functorSplit = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.Func3(func(g_2 gopurs_runtime.Value, h_3 gopurs_runtime.Value, fx_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, g_2, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, h_3), fx_4}))}
			}), (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)
		})}))}
	})
	return cache_Data_Profunctor_Split_functorSplit
}

type Constructor_Data_Profunctor_Split_SplitF[T_f any, T_a any, T_b any, T_x any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}

func Call_Data_Profunctor_Split_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Profunctor_Split_Split(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Profunctor_Split_unSplit(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply3(f_0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)
}

func Call_Data_Profunctor_Split_split(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, fx_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var fx_2 gopurs_runtime.Value = fx_2_loop
	_ = fx_2
	return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, f_0, g_1, fx_2}))}
}

func Call_Data_Profunctor_Split_lowerSplit(dictInvariant_0_loop *Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictInvariant_0 *Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value] = dictInvariant_0_loop
	_ = dictInvariant_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Data_Functor_Invariant_imap(dictInvariant_0)
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(__local_var_1_0, a_4, b_3)
		}), (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
	})
}

func Call_Data_Profunctor_Split_liftSplit(fx_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var fx_0 gopurs_runtime.Value = fx_0_loop
	_ = fx_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}), fx_0}))}
}

func Call_Data_Profunctor_Split_hoistSplit(nat_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var nat_0 gopurs_runtime.Value = nat_0_loop
	_ = nat_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply3(gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(fx_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, f_2, g_3, fx_4}))}
		}), nat_0)
	}), (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)
}
