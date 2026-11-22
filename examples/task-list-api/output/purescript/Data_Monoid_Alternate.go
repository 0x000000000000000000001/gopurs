package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Monoid_Alternate_go__append gopurs_runtime.Value
var once_Data_Monoid_Alternate_go__append sync.Once

func Get_Data_Monoid_Alternate_go__append() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_go__append.Do(func() {
		cache_Data_Monoid_Alternate_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Monoid_Alternate_go__append
}

var cache_Data_Monoid_Alternate_Alternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_Alternate sync.Once

func Get_Data_Monoid_Alternate_Alternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_Alternate.Do(func() {
		cache_Data_Monoid_Alternate_Alternate = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_Alternate(x_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_Alternate
}

var cache_Data_Monoid_Alternate_showAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_showAlternate sync.Once

func Get_Data_Monoid_Alternate_showAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_showAlternate.Do(func() {
		cache_Data_Monoid_Alternate_showAlternate = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_showAlternate(dictShow_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_showAlternate
}

var cache_Data_Monoid_Alternate_semigroupAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_semigroupAlternate sync.Once

func Get_Data_Monoid_Alternate_semigroupAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_semigroupAlternate.Do(func() {
		cache_Data_Monoid_Alternate_semigroupAlternate = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_semigroupAlternate(dictAlt_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_semigroupAlternate
}

var cache_Data_Monoid_Alternate_plusAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_plusAlternate sync.Once

func Get_Data_Monoid_Alternate_plusAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_plusAlternate.Do(func() {
		cache_Data_Monoid_Alternate_plusAlternate = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_plusAlternate(dictPlus_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_plusAlternate
}

var cache_Data_Monoid_Alternate_ordAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_ordAlternate sync.Once

func Get_Data_Monoid_Alternate_ordAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_ordAlternate.Do(func() {
		cache_Data_Monoid_Alternate_ordAlternate = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_ordAlternate(dictOrd_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_ordAlternate
}

var cache_Data_Monoid_Alternate_ord1Alternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_ord1Alternate sync.Once

func Get_Data_Monoid_Alternate_ord1Alternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_ord1Alternate.Do(func() {
		cache_Data_Monoid_Alternate_ord1Alternate = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_ord1Alternate(dictOrd1_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_ord1Alternate
}

var cache_Data_Monoid_Alternate_newtypeAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_newtypeAlternate sync.Once

func Get_Data_Monoid_Alternate_newtypeAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_newtypeAlternate.Do(func() {
		cache_Data_Monoid_Alternate_newtypeAlternate = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Monoid_Alternate_newtypeAlternate
}

var cache_Data_Monoid_Alternate_monoidAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_monoidAlternate sync.Once

func Get_Data_Monoid_Alternate_monoidAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_monoidAlternate.Do(func() {
		cache_Data_Monoid_Alternate_monoidAlternate = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_monoidAlternate(dictPlus_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_monoidAlternate
}

var cache_Data_Monoid_Alternate_monadAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_monadAlternate sync.Once

func Get_Data_Monoid_Alternate_monadAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_monadAlternate.Do(func() {
		cache_Data_Monoid_Alternate_monadAlternate = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_monadAlternate(dictMonad_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_monadAlternate
}

var cache_Data_Monoid_Alternate_functorAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_functorAlternate sync.Once

func Get_Data_Monoid_Alternate_functorAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_functorAlternate.Do(func() {
		cache_Data_Monoid_Alternate_functorAlternate = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_functorAlternate(dictFunctor_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_functorAlternate
}

var cache_Data_Monoid_Alternate_extendAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_extendAlternate sync.Once

func Get_Data_Monoid_Alternate_extendAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_extendAlternate.Do(func() {
		cache_Data_Monoid_Alternate_extendAlternate = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_extendAlternate(dictExtend_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_extendAlternate
}

var cache_Data_Monoid_Alternate_eqAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_eqAlternate sync.Once

func Get_Data_Monoid_Alternate_eqAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_eqAlternate.Do(func() {
		cache_Data_Monoid_Alternate_eqAlternate = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_eqAlternate(dictEq_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_eqAlternate
}

var cache_Data_Monoid_Alternate_eq1Alternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_eq1Alternate sync.Once

func Get_Data_Monoid_Alternate_eq1Alternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_eq1Alternate.Do(func() {
		cache_Data_Monoid_Alternate_eq1Alternate = gopurs_runtime.Func(func(dictEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_eq1Alternate(dictEq1_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_eq1Alternate
}

var cache_Data_Monoid_Alternate_comonadAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_comonadAlternate sync.Once

func Get_Data_Monoid_Alternate_comonadAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_comonadAlternate.Do(func() {
		cache_Data_Monoid_Alternate_comonadAlternate = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_comonadAlternate(dictComonad_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_comonadAlternate
}

var cache_Data_Monoid_Alternate_boundedAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_boundedAlternate sync.Once

func Get_Data_Monoid_Alternate_boundedAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_boundedAlternate.Do(func() {
		cache_Data_Monoid_Alternate_boundedAlternate = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_boundedAlternate(dictBounded_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_boundedAlternate
}

var cache_Data_Monoid_Alternate_bindAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_bindAlternate sync.Once

func Get_Data_Monoid_Alternate_bindAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_bindAlternate.Do(func() {
		cache_Data_Monoid_Alternate_bindAlternate = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_bindAlternate(dictBind_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_bindAlternate
}

var cache_Data_Monoid_Alternate_applyAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_applyAlternate sync.Once

func Get_Data_Monoid_Alternate_applyAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_applyAlternate.Do(func() {
		cache_Data_Monoid_Alternate_applyAlternate = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_applyAlternate(dictApply_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_applyAlternate
}

var cache_Data_Monoid_Alternate_applicativeAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_applicativeAlternate sync.Once

func Get_Data_Monoid_Alternate_applicativeAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_applicativeAlternate.Do(func() {
		cache_Data_Monoid_Alternate_applicativeAlternate = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_applicativeAlternate(dictApplicative_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_applicativeAlternate
}

var cache_Data_Monoid_Alternate_alternativeAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_alternativeAlternate sync.Once

func Get_Data_Monoid_Alternate_alternativeAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_alternativeAlternate.Do(func() {
		cache_Data_Monoid_Alternate_alternativeAlternate = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_alternativeAlternate(dictAlternative_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_alternativeAlternate
}

var cache_Data_Monoid_Alternate_altAlternate gopurs_runtime.Value
var once_Data_Monoid_Alternate_altAlternate sync.Once

func Get_Data_Monoid_Alternate_altAlternate() gopurs_runtime.Value {
	once_Data_Monoid_Alternate_altAlternate.Do(func() {
		cache_Data_Monoid_Alternate_altAlternate = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Alternate_altAlternate(dictAlt_0_box)
		})
	})
	return cache_Data_Monoid_Alternate_altAlternate
}

func Call_Data_Monoid_Alternate_Alternate(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Monoid_Alternate_showAlternate(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Alternate ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Monoid_Alternate_semigroupAlternate(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
	_ = dictAlt_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), v_1, v1_2)
	})}))}
}

func Call_Data_Monoid_Alternate_plusAlternate(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
	_ = dictPlus_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](dictPlus_0))}
}

func Call_Data_Monoid_Alternate_ordAlternate(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))}
}

func Call_Data_Monoid_Alternate_ord1Alternate(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
	_ = dictOrd1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](dictOrd1_0))}
}

func Call_Data_Monoid_Alternate_monoidAlternate(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
	_ = dictPlus_0
	// TAST (Let): semigroupAlternate1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f$scope30), (TypeVar a$scope31)])])
	semigroupAlternate1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Monoid_Alternate_semigroupAlternate(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})))
	_ = semigroupAlternate1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAlternate1_1_0)}
	}), Call_Control_Plus_empty(dictPlus_0)}))}
}

func Call_Data_Monoid_Alternate_monadAlternate(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0))}
}

func Call_Data_Monoid_Alternate_functorAlternate(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0))}
}

func Call_Data_Monoid_Alternate_extendAlternate(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
	_ = dictExtend_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](dictExtend_0))}
}

func Call_Data_Monoid_Alternate_eqAlternate(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}
}

func Call_Data_Monoid_Alternate_eq1Alternate(dictEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0))}
}

func Call_Data_Monoid_Alternate_comonadAlternate(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
	_ = dictComonad_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](dictComonad_0))}
}

func Call_Data_Monoid_Alternate_boundedAlternate(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
	_ = dictBounded_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](dictBounded_0))}
}

func Call_Data_Monoid_Alternate_bindAlternate(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
	_ = dictBind_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0))}
}

func Call_Data_Monoid_Alternate_applyAlternate(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))}
}

func Call_Data_Monoid_Alternate_applicativeAlternate(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}
}

func Call_Data_Monoid_Alternate_alternativeAlternate(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
	_ = dictAlternative_0
	return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0))}
}

func Call_Data_Monoid_Alternate_altAlternate(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
	_ = dictAlt_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](dictAlt_0))}
}
