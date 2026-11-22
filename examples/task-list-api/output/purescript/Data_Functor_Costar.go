package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Costar_composeFlipped gopurs_runtime.Value
var once_Data_Functor_Costar_composeFlipped sync.Once

func Get_Data_Functor_Costar_composeFlipped() gopurs_runtime.Value {
	once_Data_Functor_Costar_composeFlipped.Do(func() {
		cache_Data_Functor_Costar_composeFlipped = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Costar_composeFlipped
}

var cache_Data_Functor_Costar_composeFlipped1 gopurs_runtime.Value
var once_Data_Functor_Costar_composeFlipped1 sync.Once

func Get_Data_Functor_Costar_composeFlipped1() gopurs_runtime.Value {
	once_Data_Functor_Costar_composeFlipped1.Do(func() {
		cache_Data_Functor_Costar_composeFlipped1 = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Costar_composeFlipped1
}

var cache_Data_Functor_Costar_lcmap gopurs_runtime.Value
var once_Data_Functor_Costar_lcmap sync.Once

func Get_Data_Functor_Costar_lcmap() gopurs_runtime.Value {
	once_Data_Functor_Costar_lcmap.Do(func() {
		cache_Data_Functor_Costar_lcmap = gopurs_runtime.Func(func(a2b_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Costar_lcmap(a2b_0_box)
		})
	})
	return cache_Data_Functor_Costar_lcmap
}

var cache_Data_Functor_Costar_compose gopurs_runtime.Value
var once_Data_Functor_Costar_compose sync.Once

func Get_Data_Functor_Costar_compose() gopurs_runtime.Value {
	once_Data_Functor_Costar_compose.Do(func() {
		cache_Data_Functor_Costar_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Costar_compose
}

var cache_Data_Functor_Costar_compose1 gopurs_runtime.Value
var once_Data_Functor_Costar_compose1 sync.Once

func Get_Data_Functor_Costar_compose1() gopurs_runtime.Value {
	once_Data_Functor_Costar_compose1.Do(func() {
		cache_Data_Functor_Costar_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Costar_compose1
}

var cache_Data_Functor_Costar_composeFlipped2 gopurs_runtime.Value
var once_Data_Functor_Costar_composeFlipped2 sync.Once

func Get_Data_Functor_Costar_composeFlipped2() gopurs_runtime.Value {
	once_Data_Functor_Costar_composeFlipped2.Do(func() {
		cache_Data_Functor_Costar_composeFlipped2 = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Costar_composeFlipped2
}

var cache_Data_Functor_Costar_composeFlipped3 gopurs_runtime.Value
var once_Data_Functor_Costar_composeFlipped3 sync.Once

func Get_Data_Functor_Costar_composeFlipped3() gopurs_runtime.Value {
	once_Data_Functor_Costar_composeFlipped3.Do(func() {
		cache_Data_Functor_Costar_composeFlipped3 = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Costar_composeFlipped3
}

var cache_Data_Functor_Costar_Costar gopurs_runtime.Value
var once_Data_Functor_Costar_Costar sync.Once

func Get_Data_Functor_Costar_Costar() gopurs_runtime.Value {
	once_Data_Functor_Costar_Costar.Do(func() {
		cache_Data_Functor_Costar_Costar = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Costar_Costar(x_0_box)
		})
	})
	return cache_Data_Functor_Costar_Costar
}

var cache_Data_Functor_Costar_semigroupoidCostar gopurs_runtime.Value
var once_Data_Functor_Costar_semigroupoidCostar sync.Once

func Get_Data_Functor_Costar_semigroupoidCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_semigroupoidCostar.Do(func() {
		cache_Data_Functor_Costar_semigroupoidCostar = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Costar_semigroupoidCostar(dictExtend_0_box)
		})
	})
	return cache_Data_Functor_Costar_semigroupoidCostar
}

var cache_Data_Functor_Costar_profunctorCostar gopurs_runtime.Value
var once_Data_Functor_Costar_profunctorCostar sync.Once

func Get_Data_Functor_Costar_profunctorCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_profunctorCostar.Do(func() {
		cache_Data_Functor_Costar_profunctorCostar = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Costar_profunctorCostar(dictFunctor_0_box)
		})
	})
	return cache_Data_Functor_Costar_profunctorCostar
}

var cache_Data_Functor_Costar_strongCostar gopurs_runtime.Value
var once_Data_Functor_Costar_strongCostar sync.Once

func Get_Data_Functor_Costar_strongCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_strongCostar.Do(func() {
		cache_Data_Functor_Costar_strongCostar = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Costar_strongCostar(dictComonad_0_box)
		})
	})
	return cache_Data_Functor_Costar_strongCostar
}

var cache_Data_Functor_Costar_newtypeCostar gopurs_runtime.Value
var once_Data_Functor_Costar_newtypeCostar sync.Once

func Get_Data_Functor_Costar_newtypeCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_newtypeCostar.Do(func() {
		cache_Data_Functor_Costar_newtypeCostar = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Functor_Costar_newtypeCostar
}

var cache_Data_Functor_Costar_hoistCostar gopurs_runtime.Value
var once_Data_Functor_Costar_hoistCostar sync.Once

func Get_Data_Functor_Costar_hoistCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_hoistCostar.Do(func() {
		cache_Data_Functor_Costar_hoistCostar = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Costar_hoistCostar(f_0_box, v_1_box)
		})
	})
	return cache_Data_Functor_Costar_hoistCostar
}

var cache_Data_Functor_Costar_functorCostar gopurs_runtime.Value
var once_Data_Functor_Costar_functorCostar sync.Once

func Get_Data_Functor_Costar_functorCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_functorCostar.Do(func() {
		cache_Data_Functor_Costar_functorCostar = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, v_1)
		})}))}
	})
	return cache_Data_Functor_Costar_functorCostar
}

var cache_Data_Functor_Costar_invariantCostar gopurs_runtime.Value
var once_Data_Functor_Costar_invariantCostar sync.Once

func Get_Data_Functor_Costar_invariantCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_invariantCostar.Do(func() {
		cache_Data_Functor_Costar_invariantCostar = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Functor_Invariant_imapF(), Get_Data_Functor_Costar_functorCostar())}))}
	})
	return cache_Data_Functor_Costar_invariantCostar
}

var cache_Data_Functor_Costar_distributiveCostar gopurs_runtime.Value
var once_Data_Functor_Costar_distributiveCostar sync.Once

func Get_Data_Functor_Costar_distributiveCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_distributiveCostar.Do(func() {
		cache_Data_Functor_Costar_distributiveCostar = gopurs_runtime.Value{Type: 9, IntVal: 457335066, UnsafePtr: unsafe.Pointer((&Constructor_Data_Distributive_Distributive[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Functor_Costar_functorCostar()
		}), gopurs_runtime.Func2(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Functor_Costar_distributiveCostar(), "distribute"), dictFunctor_0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1))
		}), gopurs_runtime.Func3(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(v_3, a_2)
			}), f_1)
		})}))}
	})
	return cache_Data_Functor_Costar_distributiveCostar
}

var cache_Data_Functor_Costar_closedCostar gopurs_runtime.Value
var once_Data_Functor_Costar_closedCostar sync.Once

func Get_Data_Functor_Costar_closedCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_closedCostar.Do(func() {
		cache_Data_Functor_Costar_closedCostar = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Costar_closedCostar(dictFunctor_0_box)
		})
	})
	return cache_Data_Functor_Costar_closedCostar
}

var cache_Data_Functor_Costar_categoryCostar gopurs_runtime.Value
var once_Data_Functor_Costar_categoryCostar sync.Once

func Get_Data_Functor_Costar_categoryCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_categoryCostar.Do(func() {
		cache_Data_Functor_Costar_categoryCostar = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Costar_categoryCostar(dictComonad_0_box)
		})
	})
	return cache_Data_Functor_Costar_categoryCostar
}

var cache_Data_Functor_Costar_bifunctorCostar gopurs_runtime.Value
var once_Data_Functor_Costar_bifunctorCostar sync.Once

func Get_Data_Functor_Costar_bifunctorCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_bifunctorCostar.Do(func() {
		cache_Data_Functor_Costar_bifunctorCostar = gopurs_runtime.Func(func(dictContravariant_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Costar_bifunctorCostar(dictContravariant_0_box)
		})
	})
	return cache_Data_Functor_Costar_bifunctorCostar
}

var cache_Data_Functor_Costar_applyCostar gopurs_runtime.Value
var once_Data_Functor_Costar_applyCostar sync.Once

func Get_Data_Functor_Costar_applyCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_applyCostar.Do(func() {
		cache_Data_Functor_Costar_applyCostar = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Functor_Costar_functorCostar()
		}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(v_0, a_2, gopurs_runtime.Apply(v1_1, a_2))
		})}))}
	})
	return cache_Data_Functor_Costar_applyCostar
}

var cache_Data_Functor_Costar_bindCostar gopurs_runtime.Value
var once_Data_Functor_Costar_bindCostar sync.Once

func Get_Data_Functor_Costar_bindCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_bindCostar.Do(func() {
		cache_Data_Functor_Costar_bindCostar = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Functor_Costar_applyCostar()
		}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(v_0, x_2), x_2)
		})}))}
	})
	return cache_Data_Functor_Costar_bindCostar
}

var cache_Data_Functor_Costar_applicativeCostar gopurs_runtime.Value
var once_Data_Functor_Costar_applicativeCostar sync.Once

func Get_Data_Functor_Costar_applicativeCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_applicativeCostar.Do(func() {
		cache_Data_Functor_Costar_applicativeCostar = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Functor_Costar_applyCostar()
		}), gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return a_0
		})}))}
	})
	return cache_Data_Functor_Costar_applicativeCostar
}

var cache_Data_Functor_Costar_monadCostar gopurs_runtime.Value
var once_Data_Functor_Costar_monadCostar sync.Once

func Get_Data_Functor_Costar_monadCostar() gopurs_runtime.Value {
	once_Data_Functor_Costar_monadCostar.Do(func() {
		cache_Data_Functor_Costar_monadCostar = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Functor_Costar_applicativeCostar()
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Functor_Costar_bindCostar()
		})}))}
	})
	return cache_Data_Functor_Costar_monadCostar
}

func Call_Data_Functor_Costar_lcmap(a2b_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a2b_0 gopurs_runtime.Value = a2b_0_loop
	_ = a2b_0
	return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](Get_Data_Profunctor_profunctorFn()).V0, a2b_0, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}

func Call_Data_Functor_Costar_Costar(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Functor_Costar_semigroupoidCostar(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
	_ = dictExtend_0
	return gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer((&Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, w_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), v1_2, w_3))
	})}))}
}

func Call_Data_Functor_Costar_profunctorCostar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1), gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), v_3, g_2))
	})}))}
}

func Call_Data_Functor_Costar_strongCostar(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
	_ = dictComonad_0
	// TAST (Let): Extend0_1_0 shape=App(Other) bindingType=Any
	Extend0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
	_ = Extend0_1_0
	// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope1)])
	Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Extend0_1_0, "Functor0"), gopurs_runtime.Value{}))
	_ = Functor0_2_1
	// TAST (Let): profunctorCostar1_3_2 shape=App(Var) bindingType=(ADT ["Data","Profunctor","Profunctor"] [(Func [(TypeApp (TypeVar f$scope1) [(TypeVar b)])] (TypeVar a))])
	profunctorCostar1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](Call_Data_Functor_Costar_profunctorCostar(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Extend0_1_0, "Functor0"), gopurs_runtime.Value{})))
	_ = profunctorCostar1_3_2
	return gopurs_runtime.Value{Type: 9, IntVal: 1323482783, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Strong_Strong[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(profunctorCostar1_3_2)}
	}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{gopurs_runtime.Apply(v_4, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_Tuple_fst(), x_5)), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), x_5).UnsafePtr).V1}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
			}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), x_5).UnsafePtr).V0, gopurs_runtime.Apply(v_4, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_Tuple_snd(), x_5))}
			return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
		}()
	})}))}
}

func Call_Data_Functor_Costar_hoistCostar(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), v_1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_2
	})))
}

func Call_Data_Functor_Costar_closedCostar(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	// TAST (Let): profunctorCostar1_1_0 shape=App(Var) bindingType=(ADT ["Data","Profunctor","Profunctor"] [(Func [(TypeApp (TypeVar f$scope122) [(TypeVar b)])] (TypeVar a))])
	profunctorCostar1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](Call_Data_Functor_Costar_profunctorCostar(dictFunctor_0))
	_ = profunctorCostar1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 768764671, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Closed_Closed[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(profunctorCostar1_1_0)}
	}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(v_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(v1_5, x_4)
		}), g_3))
	})}))}
}

func Call_Data_Functor_Costar_categoryCostar(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
	_ = dictComonad_0
	// TAST (Let): semigroupoidCostar1_1_0 shape=App(Var) bindingType=(TypeApp (ADT ["Control","Semigroupoid","Semigroupoid"] []) [(Func [(TypeApp (TypeVar f$scope134) [(TypeVar b)])] (TypeVar a))])
	semigroupoidCostar1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Call_Data_Functor_Costar_semigroupoidCostar(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})))
	_ = semigroupoidCostar1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer((&Constructor_Control_Category_Category[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer(semigroupoidCostar1_1_0)}
	}), Call_Control_Comonad_extract(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](dictComonad_0))}))}
}

func Call_Data_Functor_Costar_bifunctorCostar(dictContravariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
	_ = dictContravariant_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), f_1), gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), v_3, g_2))
	})}))}
}
