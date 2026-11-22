package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Joker_go__append gopurs_runtime.Value
var once_Data_Functor_Joker_go__append sync.Once

func Get_Data_Functor_Joker_go__append() gopurs_runtime.Value {
	once_Data_Functor_Joker_go__append.Do(func() {
		cache_Data_Functor_Joker_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_Functor_Joker_go__append
}

var cache_Data_Functor_Joker_composeFlipped gopurs_runtime.Value
var once_Data_Functor_Joker_composeFlipped sync.Once

func Get_Data_Functor_Joker_composeFlipped() gopurs_runtime.Value {
	once_Data_Functor_Joker_composeFlipped.Do(func() {
		cache_Data_Functor_Joker_composeFlipped = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Joker_composeFlipped
}

var cache_Data_Functor_Joker_un gopurs_runtime.Value
var once_Data_Functor_Joker_un sync.Once

func Get_Data_Functor_Joker_un() gopurs_runtime.Value {
	once_Data_Functor_Joker_un.Do(func() {
		cache_Data_Functor_Joker_un = gopurs_runtime.Apply(Get_Data_Newtype_un(), gopurs_runtime.Value{})
	})
	return cache_Data_Functor_Joker_un
}

var cache_Data_Functor_Joker_compose gopurs_runtime.Value
var once_Data_Functor_Joker_compose sync.Once

func Get_Data_Functor_Joker_compose() gopurs_runtime.Value {
	once_Data_Functor_Joker_compose.Do(func() {
		cache_Data_Functor_Joker_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Data_Functor_Joker_compose
}

var cache_Data_Functor_Joker_Joker gopurs_runtime.Value
var once_Data_Functor_Joker_Joker sync.Once

func Get_Data_Functor_Joker_Joker() gopurs_runtime.Value {
	once_Data_Functor_Joker_Joker.Do(func() {
		cache_Data_Functor_Joker_Joker = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_Joker(x_0_box)
		})
	})
	return cache_Data_Functor_Joker_Joker
}

var cache_Data_Functor_Joker_showJoker gopurs_runtime.Value
var once_Data_Functor_Joker_showJoker sync.Once

func Get_Data_Functor_Joker_showJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_showJoker.Do(func() {
		cache_Data_Functor_Joker_showJoker = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_showJoker(dictShow_0_box)
		})
	})
	return cache_Data_Functor_Joker_showJoker
}

var cache_Data_Functor_Joker_profunctorJoker gopurs_runtime.Value
var once_Data_Functor_Joker_profunctorJoker sync.Once

func Get_Data_Functor_Joker_profunctorJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_profunctorJoker.Do(func() {
		cache_Data_Functor_Joker_profunctorJoker = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_profunctorJoker(dictFunctor_0_box)
		})
	})
	return cache_Data_Functor_Joker_profunctorJoker
}

var cache_Data_Functor_Joker_ordJoker gopurs_runtime.Value
var once_Data_Functor_Joker_ordJoker sync.Once

func Get_Data_Functor_Joker_ordJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_ordJoker.Do(func() {
		cache_Data_Functor_Joker_ordJoker = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_ordJoker(dictOrd_0_box)
		})
	})
	return cache_Data_Functor_Joker_ordJoker
}

var cache_Data_Functor_Joker_newtypeJoker gopurs_runtime.Value
var once_Data_Functor_Joker_newtypeJoker sync.Once

func Get_Data_Functor_Joker_newtypeJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_newtypeJoker.Do(func() {
		cache_Data_Functor_Joker_newtypeJoker = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Data_Functor_Joker_newtypeJoker
}

var cache_Data_Functor_Joker_hoistJoker gopurs_runtime.Value
var once_Data_Functor_Joker_hoistJoker sync.Once

func Get_Data_Functor_Joker_hoistJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_hoistJoker.Do(func() {
		cache_Data_Functor_Joker_hoistJoker = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_hoistJoker(f_0_box, v_1_box)
		})
	})
	return cache_Data_Functor_Joker_hoistJoker
}

var cache_Data_Functor_Joker_functorJoker gopurs_runtime.Value
var once_Data_Functor_Joker_functorJoker sync.Once

func Get_Data_Functor_Joker_functorJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_functorJoker.Do(func() {
		cache_Data_Functor_Joker_functorJoker = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_functorJoker(dictFunctor_0_box)
		})
	})
	return cache_Data_Functor_Joker_functorJoker
}

var cache_Data_Functor_Joker_eqJoker gopurs_runtime.Value
var once_Data_Functor_Joker_eqJoker sync.Once

func Get_Data_Functor_Joker_eqJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_eqJoker.Do(func() {
		cache_Data_Functor_Joker_eqJoker = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_eqJoker(dictEq_0_box)
		})
	})
	return cache_Data_Functor_Joker_eqJoker
}

var cache_Data_Functor_Joker_choiceJoker gopurs_runtime.Value
var once_Data_Functor_Joker_choiceJoker sync.Once

func Get_Data_Functor_Joker_choiceJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_choiceJoker.Do(func() {
		cache_Data_Functor_Joker_choiceJoker = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_choiceJoker(dictFunctor_0_box)
		})
	})
	return cache_Data_Functor_Joker_choiceJoker
}

var cache_Data_Functor_Joker_bifunctorJoker gopurs_runtime.Value
var once_Data_Functor_Joker_bifunctorJoker sync.Once

func Get_Data_Functor_Joker_bifunctorJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_bifunctorJoker.Do(func() {
		cache_Data_Functor_Joker_bifunctorJoker = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_bifunctorJoker(dictFunctor_0_box)
		})
	})
	return cache_Data_Functor_Joker_bifunctorJoker
}

var cache_Data_Functor_Joker_biapplyJoker gopurs_runtime.Value
var once_Data_Functor_Joker_biapplyJoker sync.Once

func Get_Data_Functor_Joker_biapplyJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_biapplyJoker.Do(func() {
		cache_Data_Functor_Joker_biapplyJoker = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_biapplyJoker(dictApply_0_box)
		})
	})
	return cache_Data_Functor_Joker_biapplyJoker
}

var cache_Data_Functor_Joker_biapplicativeJoker gopurs_runtime.Value
var once_Data_Functor_Joker_biapplicativeJoker sync.Once

func Get_Data_Functor_Joker_biapplicativeJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_biapplicativeJoker.Do(func() {
		cache_Data_Functor_Joker_biapplicativeJoker = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_biapplicativeJoker(dictApplicative_0_box)
		})
	})
	return cache_Data_Functor_Joker_biapplicativeJoker
}

var cache_Data_Functor_Joker_applyJoker gopurs_runtime.Value
var once_Data_Functor_Joker_applyJoker sync.Once

func Get_Data_Functor_Joker_applyJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_applyJoker.Do(func() {
		cache_Data_Functor_Joker_applyJoker = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_applyJoker(dictApply_0_box)
		})
	})
	return cache_Data_Functor_Joker_applyJoker
}

var cache_Data_Functor_Joker_bindJoker gopurs_runtime.Value
var once_Data_Functor_Joker_bindJoker sync.Once

func Get_Data_Functor_Joker_bindJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_bindJoker.Do(func() {
		cache_Data_Functor_Joker_bindJoker = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_bindJoker(dictBind_0_box)
		})
	})
	return cache_Data_Functor_Joker_bindJoker
}

var cache_Data_Functor_Joker_applicativeJoker gopurs_runtime.Value
var once_Data_Functor_Joker_applicativeJoker sync.Once

func Get_Data_Functor_Joker_applicativeJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_applicativeJoker.Do(func() {
		cache_Data_Functor_Joker_applicativeJoker = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_applicativeJoker(dictApplicative_0_box)
		})
	})
	return cache_Data_Functor_Joker_applicativeJoker
}

var cache_Data_Functor_Joker_monadJoker gopurs_runtime.Value
var once_Data_Functor_Joker_monadJoker sync.Once

func Get_Data_Functor_Joker_monadJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_monadJoker.Do(func() {
		cache_Data_Functor_Joker_monadJoker = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_Joker_monadJoker(dictMonad_0_box)
		})
	})
	return cache_Data_Functor_Joker_monadJoker
}

func Call_Data_Functor_Joker_Joker(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_Functor_Joker_showJoker(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Joker ", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal()), ")"))
	})}))}
}

func Call_Data_Functor_Joker_profunctorJoker(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), g_2, v1_3)
	})}))}
}

func Call_Data_Functor_Joker_ordJoker(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))}
}

func Call_Data_Functor_Joker_hoistJoker(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Data_Functor_Joker_functorJoker(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, v_2)
	})}))}
}

func Call_Data_Functor_Joker_eqJoker(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}
}

func Call_Data_Functor_Joker_choiceJoker(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	// TAST (Let): profunctorJoker1_1_0 shape=App(Var) bindingType=(ADT ["Data","Profunctor","Profunctor"] [(TypeApp (TypeVar f$scope72) [(TypeVar b)])])
	profunctorJoker1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](Call_Data_Functor_Joker_profunctorJoker(dictFunctor_0))
	_ = profunctorJoker1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3666633887, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Choice_Choice[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(profunctorJoker1_1_0)}
	}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), Get_Data_Either_Left(), v_2)
	}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), Get_Data_Either_Right(), v_2)
	})}))}
}

func Call_Data_Functor_Joker_bifunctorJoker(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), g_2, v1_3)
	})}))}
}

func Call_Data_Functor_Joker_biapplyJoker(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): bifunctorJoker1_1_0 shape=App(Var) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(TypeApp (TypeVar g$scope127) [(TypeVar b)])])
	bifunctorJoker1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Call_Data_Functor_Joker_bifunctorJoker(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})))
	_ = bifunctorJoker1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3774602829, UnsafePtr: unsafe.Pointer((&Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(bifunctorJoker1_1_0)}
	}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), v_2, v1_3)
	})}))}
}

func Call_Data_Functor_Joker_biapplicativeJoker(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): biapplyJoker1_1_0 shape=App(Var) bindingType=(ADT ["Control","Biapply","Biapply"] [(TypeApp (TypeVar g$scope141) [(TypeVar b)])])
	biapplyJoker1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]](Call_Data_Functor_Joker_biapplyJoker(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})))
	_ = biapplyJoker1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3949191309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Biapplicative_Biapplicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3774602829, UnsafePtr: unsafe.Pointer(biapplyJoker1_1_0)}
	}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), b_3)
	})}))}
}

func Call_Data_Functor_Joker_applyJoker(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
	_ = dictApply_0
	// TAST (Let): functorJoker1_1_0 shape=App(Var) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar f$scope152) [(TypeVar b)])])
	functorJoker1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Data_Functor_Joker_functorJoker(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})))
	_ = functorJoker1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorJoker1_1_0)}
	}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), v_2, v1_3)
	})}))}
}

func Call_Data_Functor_Joker_bindJoker(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
	_ = dictBind_0
	// TAST (Let): applyJoker1_1_0 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar f$scope98) [(TypeVar b)])])
	applyJoker1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Data_Functor_Joker_applyJoker(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})))
	_ = applyJoker1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyJoker1_1_0)}
	}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, amb_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), v_2, gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), amb_3, Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
	})}))}
}

func Call_Data_Functor_Joker_applicativeJoker(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
	_ = dictApplicative_0
	// TAST (Let): applyJoker1_1_0 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar f$scope164) [(TypeVar b)])])
	applyJoker1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Data_Functor_Joker_applyJoker(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})))
	_ = applyJoker1_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyJoker1_1_0)}
	}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Functor_Joker_Joker(), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0)))}))}
}

func Call_Data_Functor_Joker_monadJoker(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	// TAST (Let): applicativeJoker1_1_0 shape=App(Var) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (TypeVar m$scope38) [(TypeVar b)])])
	applicativeJoker1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Data_Functor_Joker_applicativeJoker(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})))
	_ = applicativeJoker1_1_0
	// TAST (Let): bindJoker1_2_1 shape=App(Var) bindingType=(ADT ["Control","Bind","Bind"] [(TypeApp (TypeVar m$scope38) [(TypeVar b)])])
	bindJoker1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Call_Data_Functor_Joker_bindJoker(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})))
	_ = bindJoker1_2_1
	return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeJoker1_1_0)}
	}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindJoker1_2_1)}
	})}))}
}
