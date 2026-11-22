package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_App_go__append gopurs_runtime.Value
var once_Data_Functor_App_go__append sync.Once

func Get_Data_Functor_App_go__append() gopurs_runtime.Value {
	once_Data_Functor_App_go__append.Do(func() {
		cache_Data_Functor_App_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Functor_App_go__append
}

var cache_Data_Functor_App_App gopurs_runtime.Value
var once_Data_Functor_App_App sync.Once

func Get_Data_Functor_App_App() gopurs_runtime.Value {
	once_Data_Functor_App_App.Do(func() {
		cache_Data_Functor_App_App = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_App(x_0_box)
		})
	})
	return cache_Data_Functor_App_App
}

var cache_Data_Functor_App_showApp gopurs_runtime.Value
var once_Data_Functor_App_showApp sync.Once

func Get_Data_Functor_App_showApp() gopurs_runtime.Value {
	once_Data_Functor_App_showApp.Do(func() {
		cache_Data_Functor_App_showApp = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_showApp(dictShow_0_box)
		})
	})
	return cache_Data_Functor_App_showApp
}

var cache_Data_Functor_App_semigroupApp gopurs_runtime.Value
var once_Data_Functor_App_semigroupApp sync.Once

func Get_Data_Functor_App_semigroupApp() gopurs_runtime.Value {
	once_Data_Functor_App_semigroupApp.Do(func() {
		cache_Data_Functor_App_semigroupApp = gopurs_runtime.Func2(func(dictApply_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_semigroupApp(dictApply_0_box, dictSemigroup_1_box)
		})
	})
	return cache_Data_Functor_App_semigroupApp
}

var cache_Data_Functor_App_plusApp gopurs_runtime.Value
var once_Data_Functor_App_plusApp sync.Once

func Get_Data_Functor_App_plusApp() gopurs_runtime.Value {
	once_Data_Functor_App_plusApp.Do(func() {
		cache_Data_Functor_App_plusApp = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_plusApp(dictPlus_0_box)
		})
	})
	return cache_Data_Functor_App_plusApp
}

var cache_Data_Functor_App_newtypeApp gopurs_runtime.Value
var once_Data_Functor_App_newtypeApp sync.Once

func Get_Data_Functor_App_newtypeApp() gopurs_runtime.Value {
	once_Data_Functor_App_newtypeApp.Do(func() {
		cache_Data_Functor_App_newtypeApp = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Functor_App_newtypeApp
}

var cache_Data_Functor_App_monoidApp gopurs_runtime.Value
var once_Data_Functor_App_monoidApp sync.Once

func Get_Data_Functor_App_monoidApp() gopurs_runtime.Value {
	once_Data_Functor_App_monoidApp.Do(func() {
		cache_Data_Functor_App_monoidApp = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_monoidApp(dictApplicative_0_box)
		})
	})
	return cache_Data_Functor_App_monoidApp
}

var cache_Data_Functor_App_monadPlusApp gopurs_runtime.Value
var once_Data_Functor_App_monadPlusApp sync.Once

func Get_Data_Functor_App_monadPlusApp() gopurs_runtime.Value {
	once_Data_Functor_App_monadPlusApp.Do(func() {
		cache_Data_Functor_App_monadPlusApp = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_monadPlusApp(dictMonadPlus_0_box)
		})
	})
	return cache_Data_Functor_App_monadPlusApp
}

var cache_Data_Functor_App_monadApp gopurs_runtime.Value
var once_Data_Functor_App_monadApp sync.Once

func Get_Data_Functor_App_monadApp() gopurs_runtime.Value {
	once_Data_Functor_App_monadApp.Do(func() {
		cache_Data_Functor_App_monadApp = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_monadApp(dictMonad_0_box)
		})
	})
	return cache_Data_Functor_App_monadApp
}

var cache_Data_Functor_App_lazyApp gopurs_runtime.Value
var once_Data_Functor_App_lazyApp sync.Once

func Get_Data_Functor_App_lazyApp() gopurs_runtime.Value {
	once_Data_Functor_App_lazyApp.Do(func() {
		cache_Data_Functor_App_lazyApp = gopurs_runtime.Func(func(dictLazy_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_lazyApp(dictLazy_0_box)
		})
	})
	return cache_Data_Functor_App_lazyApp
}

var cache_Data_Functor_App_hoistLowerApp gopurs_runtime.Value
var once_Data_Functor_App_hoistLowerApp sync.Once

func Get_Data_Functor_App_hoistLowerApp() gopurs_runtime.Value {
	once_Data_Functor_App_hoistLowerApp.Do(func() {
		cache_Data_Functor_App_hoistLowerApp = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Functor_App_hoistLowerApp
}

var cache_Data_Functor_App_hoistLiftApp gopurs_runtime.Value
var once_Data_Functor_App_hoistLiftApp sync.Once

func Get_Data_Functor_App_hoistLiftApp() gopurs_runtime.Value {
	once_Data_Functor_App_hoistLiftApp.Do(func() {
		cache_Data_Functor_App_hoistLiftApp = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Functor_App_hoistLiftApp
}

var cache_Data_Functor_App_hoistApp gopurs_runtime.Value
var once_Data_Functor_App_hoistApp sync.Once

func Get_Data_Functor_App_hoistApp() gopurs_runtime.Value {
	once_Data_Functor_App_hoistApp.Do(func() {
		cache_Data_Functor_App_hoistApp = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_hoistApp(f_0_box, v_1_box)
		})
	})
	return cache_Data_Functor_App_hoistApp
}

var cache_Data_Functor_App_functorApp gopurs_runtime.Value
var once_Data_Functor_App_functorApp sync.Once

func Get_Data_Functor_App_functorApp() gopurs_runtime.Value {
	once_Data_Functor_App_functorApp.Do(func() {
		cache_Data_Functor_App_functorApp = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_functorApp(dictFunctor_0_box)
		})
	})
	return cache_Data_Functor_App_functorApp
}

var cache_Data_Functor_App_extendApp gopurs_runtime.Value
var once_Data_Functor_App_extendApp sync.Once

func Get_Data_Functor_App_extendApp() gopurs_runtime.Value {
	once_Data_Functor_App_extendApp.Do(func() {
		cache_Data_Functor_App_extendApp = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_extendApp(dictExtend_0_box)
		})
	})
	return cache_Data_Functor_App_extendApp
}

var cache_Data_Functor_App_eqApp gopurs_runtime.Value
var once_Data_Functor_App_eqApp sync.Once

func Get_Data_Functor_App_eqApp() gopurs_runtime.Value {
	once_Data_Functor_App_eqApp.Do(func() {
		cache_Data_Functor_App_eqApp = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_eqApp(dictEq1_0_box, dictEq_1_box)
		})
	})
	return cache_Data_Functor_App_eqApp
}

var cache_Data_Functor_App_ordApp gopurs_runtime.Value
var once_Data_Functor_App_ordApp sync.Once

func Get_Data_Functor_App_ordApp() gopurs_runtime.Value {
	once_Data_Functor_App_ordApp.Do(func() {
		cache_Data_Functor_App_ordApp = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_ordApp(dictOrd1_0_box)
		})
	})
	return cache_Data_Functor_App_ordApp
}

var cache_Data_Functor_App_eq1App gopurs_runtime.Value
var once_Data_Functor_App_eq1App sync.Once

func Get_Data_Functor_App_eq1App() gopurs_runtime.Value {
	once_Data_Functor_App_eq1App.Do(func() {
		cache_Data_Functor_App_eq1App = gopurs_runtime.Func(func(dictEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_eq1App(dictEq1_0_box)
		})
	})
	return cache_Data_Functor_App_eq1App
}

var cache_Data_Functor_App_ord1App gopurs_runtime.Value
var once_Data_Functor_App_ord1App sync.Once

func Get_Data_Functor_App_ord1App() gopurs_runtime.Value {
	once_Data_Functor_App_ord1App.Do(func() {
		cache_Data_Functor_App_ord1App = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_ord1App(dictOrd1_0_box)
		})
	})
	return cache_Data_Functor_App_ord1App
}

var cache_Data_Functor_App_comonadApp gopurs_runtime.Value
var once_Data_Functor_App_comonadApp sync.Once

func Get_Data_Functor_App_comonadApp() gopurs_runtime.Value {
	once_Data_Functor_App_comonadApp.Do(func() {
		cache_Data_Functor_App_comonadApp = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_comonadApp(dictComonad_0_box)
		})
	})
	return cache_Data_Functor_App_comonadApp
}

var cache_Data_Functor_App_bindApp gopurs_runtime.Value
var once_Data_Functor_App_bindApp sync.Once

func Get_Data_Functor_App_bindApp() gopurs_runtime.Value {
	once_Data_Functor_App_bindApp.Do(func() {
		cache_Data_Functor_App_bindApp = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_bindApp(dictBind_0_box)
		})
	})
	return cache_Data_Functor_App_bindApp
}

var cache_Data_Functor_App_applyApp gopurs_runtime.Value
var once_Data_Functor_App_applyApp sync.Once

func Get_Data_Functor_App_applyApp() gopurs_runtime.Value {
	once_Data_Functor_App_applyApp.Do(func() {
		cache_Data_Functor_App_applyApp = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_applyApp(dictApply_0_box)
		})
	})
	return cache_Data_Functor_App_applyApp
}

var cache_Data_Functor_App_applicativeApp gopurs_runtime.Value
var once_Data_Functor_App_applicativeApp sync.Once

func Get_Data_Functor_App_applicativeApp() gopurs_runtime.Value {
	once_Data_Functor_App_applicativeApp.Do(func() {
		cache_Data_Functor_App_applicativeApp = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_applicativeApp(dictApplicative_0_box)
		})
	})
	return cache_Data_Functor_App_applicativeApp
}

var cache_Data_Functor_App_alternativeApp gopurs_runtime.Value
var once_Data_Functor_App_alternativeApp sync.Once

func Get_Data_Functor_App_alternativeApp() gopurs_runtime.Value {
	once_Data_Functor_App_alternativeApp.Do(func() {
		cache_Data_Functor_App_alternativeApp = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_alternativeApp(dictAlternative_0_box)
		})
	})
	return cache_Data_Functor_App_alternativeApp
}

var cache_Data_Functor_App_altApp gopurs_runtime.Value
var once_Data_Functor_App_altApp sync.Once

func Get_Data_Functor_App_altApp() gopurs_runtime.Value {
	once_Data_Functor_App_altApp.Do(func() {
		cache_Data_Functor_App_altApp = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_App_altApp(dictAlt_0_box)
		})
	})
	return cache_Data_Functor_App_altApp
}

func Call_Data_Functor_App_App(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Functor_App_showApp(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(App ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Functor_App_semigroupApp(dictApply_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
	_ = dictSemigroup_1
	// TAST (Let): append1_2_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope8), (TypeVar a$scope8)] (TypeVar a$scope8))
	append1_2_0 := Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_1))
	_ = append1_2_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), append1_2_0, v_3), v1_4)
	})}))}
}

func Call_Data_Functor_App_plusApp(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
	_ = dictPlus_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](dictPlus_0))}
}

func Call_Data_Functor_App_monoidApp(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): semigroupApp1_1_0 shape=App(Var) bindingType=Any
	semigroupApp1_1_0 := gopurs_runtime.Apply(Get_Data_Functor_App_semigroupApp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
	_ = semigroupApp1_1_0
	return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): semigroupApp2_3_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f$scope34), (TypeVar a$scope35)])])
		semigroupApp2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(semigroupApp1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{})))
		_ = semigroupApp2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupApp2_3_1)}
		}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordGet(dictMonoid_2, "mempty"))}))}
	})
}

func Call_Data_Functor_App_monadPlusApp(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
	_ = dictMonadPlus_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]](dictMonadPlus_0))}
}

func Call_Data_Functor_App_monadApp(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0))}
}

func Call_Data_Functor_App_lazyApp(dictLazy_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictLazy_0 gopurs_runtime.Value = dictLazy_0_loop
	_ = dictLazy_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_0))}
}

func Call_Data_Functor_App_hoistApp(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Data_Functor_App_functorApp(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0))}
}

func Call_Data_Functor_App_extendApp(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
	_ = dictExtend_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](dictExtend_0))}
}

func Call_Data_Functor_App_eqApp(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
	_ = dictEq_1
	// TAST (Let): eq11_2_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope80) [(TypeVar a$scope81)]), (TypeApp (TypeVar f$scope80) [(TypeVar a$scope81)])] Boolean)
	eq11_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1))})
	_ = eq11_2_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(eq11_2_0, x_3, y_4)
	})}))}
}

func Call_Data_Functor_App_ordApp(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
	_ = dictOrd1_0
	// TAST (Let): eqApp1_1_0 shape=App(Var) bindingType=Any
	eqApp1_1_0 := gopurs_runtime.Apply(Get_Data_Functor_App_eqApp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))
	_ = eqApp1_1_0
	return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): compare11_3_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope17) [(TypeVar a$scope18)]), (TypeApp (TypeVar f$scope17) [(TypeVar a$scope18)])] (ADT ["Data","Ordering","Ordering"] []))
		compare11_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_2))})
		_ = compare11_3_1
		// TAST (Let): eqApp2_4_2 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f$scope17), (TypeVar a$scope18)])])
		eqApp2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqApp1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_2, "Eq0"), gopurs_runtime.Value{})))
		_ = eqApp2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqApp2_4_2)}
		}), gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(compare11_3_1, x_5, y_6)
		})}))}
	})
}

func Call_Data_Functor_App_eq1App(dictEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Functor_App_eqApp(dictEq1_0, dictEq_1)))
	})}))}
}

func Call_Data_Functor_App_ord1App(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
	_ = dictOrd1_0
	// TAST (Let): ordApp1_1_0 shape=App(Var) bindingType=Any
	ordApp1_1_0 := Call_Data_Functor_App_ordApp(dictOrd1_0)
	_ = ordApp1_1_0
	// TAST (Let): eq1App1_2_1 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq1"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f$scope21)])])
	eq1App1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Call_Data_Functor_App_eq1App(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})))
	_ = eq1App1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1App1_2_1)}
	}), gopurs_runtime.Func(func(dictOrd_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordApp1_1_0, dictOrd_3)))
	})}))}
}

func Call_Data_Functor_App_comonadApp(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
	_ = dictComonad_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](dictComonad_0))}
}

func Call_Data_Functor_App_bindApp(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
	_ = dictBind_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0))}
}

func Call_Data_Functor_App_applyApp(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))}
}

func Call_Data_Functor_App_applicativeApp(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}
}

func Call_Data_Functor_App_alternativeApp(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
	_ = dictAlternative_0
	return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0))}
}

func Call_Data_Functor_App_altApp(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
	_ = dictAlt_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](dictAlt_0))}
}
