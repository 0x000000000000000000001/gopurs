package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Compose_go__append gopurs_runtime.Value
var once_Data_Functor_Compose_go__append sync.Once

func Get_Data_Functor_Compose_go__append() gopurs_runtime.Value {
	once_Data_Functor_Compose_go__append.Do(func() {
		cache_Data_Functor_Compose_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Functor_Compose_go__append
}

var cache_Data_Functor_Compose_compose gopurs_runtime.Value
var once_Data_Functor_Compose_compose sync.Once

func Get_Data_Functor_Compose_compose() gopurs_runtime.Value {
	once_Data_Functor_Compose_compose.Do(func() {
		cache_Data_Functor_Compose_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Compose_compose
}

var cache_Data_Functor_Compose_compose1 gopurs_runtime.Value
var once_Data_Functor_Compose_compose1 sync.Once

func Get_Data_Functor_Compose_compose1() gopurs_runtime.Value {
	once_Data_Functor_Compose_compose1.Do(func() {
		cache_Data_Functor_Compose_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Compose_compose1
}

var cache_Data_Functor_Compose_Compose gopurs_runtime.Value
var once_Data_Functor_Compose_Compose sync.Once

func Get_Data_Functor_Compose_Compose() gopurs_runtime.Value {
	once_Data_Functor_Compose_Compose.Do(func() {
		cache_Data_Functor_Compose_Compose = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Compose_Compose(x_0_box)
		})
	})
	return cache_Data_Functor_Compose_Compose
}

var cache_Data_Functor_Compose_showCompose gopurs_runtime.Value
var once_Data_Functor_Compose_showCompose sync.Once

func Get_Data_Functor_Compose_showCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_showCompose.Do(func() {
		cache_Data_Functor_Compose_showCompose = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Compose_showCompose(dictShow_0_box)
		})
	})
	return cache_Data_Functor_Compose_showCompose
}

var cache_Data_Functor_Compose_newtypeCompose gopurs_runtime.Value
var once_Data_Functor_Compose_newtypeCompose sync.Once

func Get_Data_Functor_Compose_newtypeCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_newtypeCompose.Do(func() {
		cache_Data_Functor_Compose_newtypeCompose = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Functor_Compose_newtypeCompose
}

var cache_Data_Functor_Compose_functorCompose gopurs_runtime.Value
var once_Data_Functor_Compose_functorCompose sync.Once

func Get_Data_Functor_Compose_functorCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_functorCompose.Do(func() {
		cache_Data_Functor_Compose_functorCompose = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Compose_functorCompose(dictFunctor_0_box, dictFunctor1_1_box)
		})
	})
	return cache_Data_Functor_Compose_functorCompose
}

var cache_Data_Functor_Compose_eqCompose gopurs_runtime.Value
var once_Data_Functor_Compose_eqCompose sync.Once

func Get_Data_Functor_Compose_eqCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_eqCompose.Do(func() {
		cache_Data_Functor_Compose_eqCompose = gopurs_runtime.Func3(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value, dictEq_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Compose_eqCompose(dictEq1_0_box, dictEq11_1_box, dictEq_2_box)
		})
	})
	return cache_Data_Functor_Compose_eqCompose
}

var cache_Data_Functor_Compose_ordCompose gopurs_runtime.Value
var once_Data_Functor_Compose_ordCompose sync.Once

func Get_Data_Functor_Compose_ordCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_ordCompose.Do(func() {
		cache_Data_Functor_Compose_ordCompose = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Compose_ordCompose(dictOrd1_0_box)
		})
	})
	return cache_Data_Functor_Compose_ordCompose
}

var cache_Data_Functor_Compose_eq1Compose gopurs_runtime.Value
var once_Data_Functor_Compose_eq1Compose sync.Once

func Get_Data_Functor_Compose_eq1Compose() gopurs_runtime.Value {
	once_Data_Functor_Compose_eq1Compose.Do(func() {
		cache_Data_Functor_Compose_eq1Compose = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Compose_eq1Compose(dictEq1_0_box, dictEq11_1_box)
		})
	})
	return cache_Data_Functor_Compose_eq1Compose
}

var cache_Data_Functor_Compose_ord1Compose gopurs_runtime.Value
var once_Data_Functor_Compose_ord1Compose sync.Once

func Get_Data_Functor_Compose_ord1Compose() gopurs_runtime.Value {
	once_Data_Functor_Compose_ord1Compose.Do(func() {
		cache_Data_Functor_Compose_ord1Compose = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Compose_ord1Compose(dictOrd1_0_box)
		})
	})
	return cache_Data_Functor_Compose_ord1Compose
}

var cache_Data_Functor_Compose_bihoistCompose gopurs_runtime.Value
var once_Data_Functor_Compose_bihoistCompose sync.Once

func Get_Data_Functor_Compose_bihoistCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_bihoistCompose.Do(func() {
		cache_Data_Functor_Compose_bihoistCompose = gopurs_runtime.Func4(func(dictFunctor_0_box gopurs_runtime.Value, natF_1_box gopurs_runtime.Value, natG_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Compose_bihoistCompose(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), natF_1_box, natG_2_box, v_3_box)
		})
	})
	return cache_Data_Functor_Compose_bihoistCompose
}

var cache_Data_Functor_Compose_applyCompose gopurs_runtime.Value
var once_Data_Functor_Compose_applyCompose sync.Once

func Get_Data_Functor_Compose_applyCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_applyCompose.Do(func() {
		cache_Data_Functor_Compose_applyCompose = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Compose_applyCompose(dictApply_0_box)
		})
	})
	return cache_Data_Functor_Compose_applyCompose
}

var cache_Data_Functor_Compose_applicativeCompose gopurs_runtime.Value
var once_Data_Functor_Compose_applicativeCompose sync.Once

func Get_Data_Functor_Compose_applicativeCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_applicativeCompose.Do(func() {
		cache_Data_Functor_Compose_applicativeCompose = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Compose_applicativeCompose(dictApplicative_0_box)
		})
	})
	return cache_Data_Functor_Compose_applicativeCompose
}

var cache_Data_Functor_Compose_altCompose gopurs_runtime.Value
var once_Data_Functor_Compose_altCompose sync.Once

func Get_Data_Functor_Compose_altCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_altCompose.Do(func() {
		cache_Data_Functor_Compose_altCompose = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Compose_altCompose(dictAlt_0_box)
		})
	})
	return cache_Data_Functor_Compose_altCompose
}

var cache_Data_Functor_Compose_plusCompose gopurs_runtime.Value
var once_Data_Functor_Compose_plusCompose sync.Once

func Get_Data_Functor_Compose_plusCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_plusCompose.Do(func() {
		cache_Data_Functor_Compose_plusCompose = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Compose_plusCompose(dictPlus_0_box)
		})
	})
	return cache_Data_Functor_Compose_plusCompose
}

var cache_Data_Functor_Compose_alternativeCompose gopurs_runtime.Value
var once_Data_Functor_Compose_alternativeCompose sync.Once

func Get_Data_Functor_Compose_alternativeCompose() gopurs_runtime.Value {
	once_Data_Functor_Compose_alternativeCompose.Do(func() {
		cache_Data_Functor_Compose_alternativeCompose = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Compose_alternativeCompose(dictAlternative_0_box)
		})
	})
	return cache_Data_Functor_Compose_alternativeCompose
}

func Call_Data_Functor_Compose_Compose(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Functor_Compose_showCompose(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Compose ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Functor_Compose_functorCompose(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
	_ = dictFunctor1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2), v_3)
	})}))}
}

func Call_Data_Functor_Compose_eqCompose(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value, dictEq_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
	_ = dictEq11_1
	var dictEq_2 gopurs_runtime.Value = dictEq_2_loop
	_ = dictEq_2
	// TAST (Let): eq11_3_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope61) [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar g$scope62), (TypeVar a$scope63)])]), (TypeApp (TypeVar f$scope61) [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar g$scope62), (TypeVar a$scope63)])])] Boolean)
	eq11_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Functor_App_eqApp(dictEq11_1, dictEq_2)))})
	_ = eq11_3_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(eq11_3_0, v_4, v1_5)
	})}))}
}

func Call_Data_Functor_Compose_ordCompose(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
	_ = dictOrd1_0
	// TAST (Let): eqCompose1_1_0 shape=App(Var) bindingType=Any
	eqCompose1_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Compose_eqCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))
	_ = eqCompose1_1_0
	return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): ordApp_3_1 shape=App(Var) bindingType=Any
		ordApp_3_1 := Call_Data_Functor_App_ordApp(dictOrd11_2)
		_ = ordApp_3_1
		// TAST (Let): eqCompose2_4_2 shape=App(Other) bindingType=Any
		eqCompose2_4_2 := gopurs_runtime.Apply(eqCompose1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{}))
		_ = eqCompose2_4_2
		return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): compare11_6_3 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f$scope22) [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar g$scope23), (TypeVar a$scope24)])]), (TypeApp (TypeVar f$scope22) [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar g$scope23), (TypeVar a$scope24)])])] (ADT ["Data","Ordering","Ordering"] []))
			compare11_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordApp_3_1, dictOrd_5)))})
			_ = compare11_6_3
			// TAST (Let): eqCompose3_7_4 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope22), (TypeVar g$scope23), (TypeVar a$scope24)])])
			eqCompose3_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqCompose2_4_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})))
			_ = eqCompose3_7_4
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqCompose3_7_4)}
			}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(compare11_6_3, v_8, v1_9)
			})}))}
		})
	})
}

func Call_Data_Functor_Compose_eq1Compose(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
	_ = dictEq11_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Functor_Compose_eqCompose(dictEq1_0, dictEq11_1, dictEq_2)))
	})}))}
}

func Call_Data_Functor_Compose_ord1Compose(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
	_ = dictOrd1_0
	// TAST (Let): ordCompose1_1_0 shape=App(Var) bindingType=Any
	ordCompose1_1_0 := Call_Data_Functor_Compose_ordCompose(dictOrd1_0)
	_ = ordCompose1_1_0
	// TAST (Let): eq1Compose1_2_1 shape=App(Var) bindingType=Any
	eq1Compose1_2_1 := gopurs_runtime.Apply(Get_Data_Functor_Compose_eq1Compose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))
	_ = eq1Compose1_2_1
	return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): ordCompose2_4_2 shape=App(Other) bindingType=Any
		ordCompose2_4_2 := gopurs_runtime.Apply(ordCompose1_1_0, dictOrd11_3)
		_ = ordCompose2_4_2
		// TAST (Let): eq1Compose2_5_3 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq1"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope30), (TypeVar g$scope31)])])
		eq1Compose2_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](gopurs_runtime.Apply(eq1Compose1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{})))
		_ = eq1Compose2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1Compose2_5_3)}
		}), gopurs_runtime.Func(func(dictOrd_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordCompose2_4_2, dictOrd_6)))
		})}))}
	})
}

func Call_Data_Functor_Compose_bihoistCompose(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], natF_1_loop gopurs_runtime.Value, natG_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var natF_1 gopurs_runtime.Value = natF_1_loop
	_ = natF_1
	var natG_2 gopurs_runtime.Value = natG_2_loop
	_ = natG_2
	var v_3 gopurs_runtime.Value = v_3_loop
	_ = v_3
	return gopurs_runtime.Apply(natF_1, gopurs_runtime.Apply2(dictFunctor_0.V0, natG_2, v_3))
}

func Call_Data_Functor_Compose_applyCompose(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): functorCompose1_1_0 shape=App(Var) bindingType=Any
	functorCompose1_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Compose_functorCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
	_ = functorCompose1_1_0
	return gopurs_runtime.Func(func(dictApply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): apply1_3_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar g$scope94) [(Func [(TypeVar a$scope98)] (TypeVar b$scope99))]), (TypeApp (TypeVar g$scope94) [(TypeVar a$scope98)])] (TypeApp (TypeVar g$scope94) [(TypeVar b$scope99)]))
		apply1_3_1 := Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply1_2))
		_ = apply1_3_1
		// TAST (Let): functorCompose2_4_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope93), (TypeVar g$scope94)])])
		functorCompose2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCompose1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply1_2, "Functor0"), gopurs_runtime.Value{})))
		_ = functorCompose2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose2_4_2)}
		}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), apply1_3_1, v_5), v1_6)
		})}))}
	})
}

func Call_Data_Functor_Compose_applicativeCompose(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): pure_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar g$scope108) [(TypeVar a$scope111)])] (TypeApp (TypeVar f$scope107) [(TypeApp (TypeVar g$scope108) [(TypeVar a$scope111)])]))
	pure_1_0 := Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))
	_ = pure_1_0
	// TAST (Let): applyCompose1_2_1 shape=App(Var) bindingType=Any
	applyCompose1_2_1 := Call_Data_Functor_Compose_applyCompose(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
	_ = applyCompose1_2_1
	return gopurs_runtime.Func(func(dictApplicative1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): applyCompose2_4_2 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope107), (TypeVar g$scope108)])])
		applyCompose2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyCompose1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative1_3, "Apply0"), gopurs_runtime.Value{})))
		_ = applyCompose2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyCompose2_4_2)}
		}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Functor_Compose_Compose(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), pure_1_0, Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative1_3))))}))}
	})
}

func Call_Data_Functor_Compose_altCompose(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
	_ = dictAlt_0
	// TAST (Let): functorCompose1_1_0 shape=App(Var) bindingType=Any
	functorCompose1_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Compose_functorCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{}))
	_ = functorCompose1_1_0
	return gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): functorCompose2_3_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope126), (TypeVar g$scope127)])])
		functorCompose2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCompose1_1_0, dictFunctor_2))
		_ = functorCompose2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose2_3_1)}
		}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), v_4, v1_5)
		})}))}
	})
}

func Call_Data_Functor_Compose_plusCompose(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
	_ = dictPlus_0
	// TAST (Let): empty_1_0 shape=App(Var) bindingType=(TypeApp (TypeVar f$scope9) [(TypeApp (TypeVar g$scope10) [(TypeVar a$scope13)])])
	empty_1_0 := Call_Control_Plus_empty(dictPlus_0)
	_ = empty_1_0
	// TAST (Let): altCompose1_2_1 shape=App(Var) bindingType=Any
	altCompose1_2_1 := Call_Data_Functor_Compose_altCompose(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{}))
	_ = altCompose1_2_1
	return gopurs_runtime.Func(func(dictFunctor_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): altCompose2_4_2 shape=App(Other) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope9), (TypeVar g$scope10)])])
		altCompose2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(altCompose1_2_1, dictFunctor_3))
		_ = altCompose2_4_2
		return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altCompose2_4_2)}
		}), empty_1_0}))}
	})
}

func Call_Data_Functor_Compose_alternativeCompose(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
	_ = dictAlternative_0
	// TAST (Let): applicativeCompose1_1_0 shape=App(Var) bindingType=Any
	applicativeCompose1_1_0 := Call_Data_Functor_Compose_applicativeCompose(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}))
	_ = applicativeCompose1_1_0
	// TAST (Let): plusCompose1_2_1 shape=App(Var) bindingType=Any
	plusCompose1_2_1 := Call_Data_Functor_Compose_plusCompose(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}))
	_ = plusCompose1_2_1
	return gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): applicativeCompose2_4_2 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope121), (TypeVar g$scope122)])])
		applicativeCompose2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeCompose1_1_0, dictApplicative_3))
		_ = applicativeCompose2_4_2
		// TAST (Let): plusCompose2_5_3 shape=App(Other) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope121), (TypeVar g$scope122)])])
		plusCompose2_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](gopurs_runtime.Apply(plusCompose1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})))
		_ = plusCompose2_5_3
		return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeCompose2_4_2)}
		}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusCompose2_5_3)}
		})}))}
	})
}
