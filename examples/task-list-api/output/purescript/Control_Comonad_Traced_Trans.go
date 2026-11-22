package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Traced_Trans_TracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_TracedT sync.Once

func Get_Control_Comonad_Traced_Trans_TracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_TracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_TracedT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Trans_TracedT(x_0_box)
		})
	})
	return cache_Control_Comonad_Traced_Trans_TracedT
}

var cache_Control_Comonad_Traced_Trans_runTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_runTracedT sync.Once

func Get_Control_Comonad_Traced_Trans_runTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_runTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_runTracedT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Trans_runTracedT(v_0_box)
		})
	})
	return cache_Control_Comonad_Traced_Trans_runTracedT
}

var cache_Control_Comonad_Traced_Trans_newtypeTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_newtypeTracedT sync.Once

func Get_Control_Comonad_Traced_Trans_newtypeTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_newtypeTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_newtypeTracedT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Control_Comonad_Traced_Trans_newtypeTracedT
}

var cache_Control_Comonad_Traced_Trans_functorTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_functorTracedT sync.Once

func Get_Control_Comonad_Traced_Trans_functorTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_functorTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_functorTracedT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Trans_functorTracedT(dictFunctor_0_box)
		})
	})
	return cache_Control_Comonad_Traced_Trans_functorTracedT
}

var cache_Control_Comonad_Traced_Trans_extendTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_extendTracedT sync.Once

func Get_Control_Comonad_Traced_Trans_extendTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_extendTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_extendTracedT = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Trans_extendTracedT(dictExtend_0_box)
		})
	})
	return cache_Control_Comonad_Traced_Trans_extendTracedT
}

var cache_Control_Comonad_Traced_Trans_comonadTransTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_comonadTransTracedT sync.Once

func Get_Control_Comonad_Traced_Trans_comonadTransTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_comonadTransTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_comonadTransTracedT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Trans_comonadTransTracedT(dictMonoid_0_box)
		})
	})
	return cache_Control_Comonad_Traced_Trans_comonadTransTracedT
}

var cache_Control_Comonad_Traced_Trans_comonadTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_comonadTracedT sync.Once

func Get_Control_Comonad_Traced_Trans_comonadTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_comonadTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_comonadTracedT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_Trans_comonadTracedT(dictComonad_0_box)
		})
	})
	return cache_Control_Comonad_Traced_Trans_comonadTracedT
}

func Call_Control_Comonad_Traced_Trans_TracedT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Comonad_Traced_Trans_runTracedT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Control_Comonad_Traced_Trans_functorTracedT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func2(func(g_3 gopurs_runtime.Value, t_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(g_3, t_4))
		}), v_2)
	})}))}
}

func Call_Control_Comonad_Traced_Trans_extendTracedT(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
	_ = dictExtend_0
	// TAST (Let): functorTracedT1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar w$scope29) [(Func [(TypeVar t$scope30)] (TypeVar a))])])
	functorTracedT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Comonad_Traced_Trans_functorTracedT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{})))
	_ = functorTracedT1_1_0
	return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorTracedT1_1_0)}
		}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func2(func(w_prime__5 gopurs_runtime.Value, t_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(f_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(h_7 gopurs_runtime.Value, t_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(h_7, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), t_6, t_prime__8))
				}), w_prime__5))
			}), v_4)
		})}))}
	})
}

func Call_Control_Comonad_Traced_Trans_comonadTransTracedT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): mempty_1_0 shape=Other bindingType=(TypeVar t$scope43)
	mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
	_ = mempty_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3399197123, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Trans_Class_ComonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(dictComonad_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_2, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(f_4, mempty_1_0)
		}), v_3)
	})}))}
}

func Call_Control_Comonad_Traced_Trans_comonadTracedT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
	_ = dictComonad_0
	// TAST (Let): extendTracedT1_1_0 shape=App(Var) bindingType=Any
	extendTracedT1_1_0 := Call_Control_Comonad_Traced_Trans_extendTracedT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{}))
	_ = extendTracedT1_1_0
	return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): mempty_3_1 shape=Other bindingType=(TypeVar t$scope52)
		mempty_3_1 := gopurs_runtime.RecordGet(dictMonoid_2, "mempty")
		_ = mempty_3_1
		// TAST (Let): extendTracedT2_4_2 shape=App(Other) bindingType=(ADT ["Control","Extend","Extend"] [(TypeApp (TypeVar w$scope51) [(Func [(TypeVar t$scope52)] (TypeVar a))])])
		extendTracedT2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](gopurs_runtime.Apply(extendTracedT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{})))
		_ = extendTracedT2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(extendTracedT2_4_2)}
		}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), v_5, mempty_3_1)
		})}))}
	})
}
