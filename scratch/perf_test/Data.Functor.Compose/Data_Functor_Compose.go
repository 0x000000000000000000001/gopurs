package Data_Functor_Compose

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var Compose gopurs_runtime.Value
var once_Compose sync.Once
func Get_Compose() gopurs_runtime.Value {
	once_Compose.Do(func() {
		Compose = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Compose
}

var showCompose gopurs_runtime.Value
var once_showCompose sync.Once
func Get_showCompose() gopurs_runtime.Value {
	once_showCompose.Do(func() {
		showCompose = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Compose ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showCompose
}

var newtypeCompose gopurs_runtime.Value
var once_newtypeCompose sync.Once
func Get_newtypeCompose() gopurs_runtime.Value {
	once_newtypeCompose.Do(func() {
		newtypeCompose = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeCompose
}

var functorCompose gopurs_runtime.Value
var once_functorCompose sync.Once
func Get_functorCompose() gopurs_runtime.Value {
	once_functorCompose.Do(func() {
		functorCompose = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(dictFunctor1_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2)), v_3)
})
})})
})
})
	})
	return functorCompose
}

var eqCompose gopurs_runtime.Value
var once_eqCompose sync.Once
func Get_eqCompose() gopurs_runtime.Value {
	once_eqCompose.Do(func() {
		eqCompose = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq11_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq11_3_1 := gopurs_runtime.Apply(dictEq11_1.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_2)
eq11_3_0 := gopurs_runtime.Apply(dictEq1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(eq11_3_1, x_4), y_5)
})
})}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(eq11_3_0, gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), v_4)), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), v1_5))
})
})})
})
})
})
	})
	return eqCompose
}

var ordCompose gopurs_runtime.Value
var once_ordCompose sync.Once
func Get_ordCompose() gopurs_runtime.Value {
	once_ordCompose.Do(func() {
		ordCompose = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
eqCompose1_1_0 := gopurs_runtime.Apply(Get_eqCompose(), gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictOrd11_2.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{})
eqCompose2_4_2 := gopurs_runtime.Apply(eqCompose1_1_0, gopurs_runtime.Apply(dictOrd11_2.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_6_4 := gopurs_runtime.Apply(dictOrd11_2.PtrVal.(map[string]gopurs_runtime.Value)["compare1"], dictOrd_5)
eq11_7_5 := gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], gopurs_runtime.Apply(dictOrd_5.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
eqApp2_8_6 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(eq11_7_5, x_8), y_9)
})
})})
compare11_6_3 := gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["compare1"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(compare11_6_4, x_9), y_10)
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return eqApp2_8_6
})}))
eqCompose3_7_7 := gopurs_runtime.Apply(eqCompose2_4_2, gopurs_runtime.Apply(dictOrd_5.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(compare11_6_3, gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), v_8)), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), v1_9))
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return eqCompose3_7_7
})})
})
})
})
	})
	return ordCompose
}

var eq1Compose gopurs_runtime.Value
var once_eq1Compose sync.Once
func Get_eq1Compose() gopurs_runtime.Value {
	once_eq1Compose.Do(func() {
		eq1Compose = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq11_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eqCompose(), dictEq1_0), dictEq11_1), dictEq_2).PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
})
})
	})
	return eq1Compose
}

var ord1Compose gopurs_runtime.Value
var once_ord1Compose sync.Once
func Get_ord1Compose() gopurs_runtime.Value {
	once_ord1Compose.Do(func() {
		ord1Compose = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
ordCompose1_1_0 := gopurs_runtime.Apply(Get_ordCompose(), dictOrd1_0)
eq1Compose1_2_1 := gopurs_runtime.Apply(Get_eq1Compose(), gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
ordCompose2_4_2 := gopurs_runtime.Apply(ordCompose1_1_0, dictOrd11_3)
eq1Compose2_5_3 := gopurs_runtime.Apply(eq1Compose1_2_1, gopurs_runtime.Apply(dictOrd11_3.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(ordCompose2_4_2, dictOrd_6).PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Compose2_5_3
})})
})
})
	})
	return ord1Compose
}

var bihoistCompose gopurs_runtime.Value
var once_bihoistCompose sync.Once
func Get_bihoistCompose() gopurs_runtime.Value {
	once_bihoistCompose.Do(func() {
		bihoistCompose = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(natF_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(natG_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(natF_1, gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], natG_2), v_3))
})
})
})
})
	})
	return bihoistCompose
}

var applyCompose gopurs_runtime.Value
var once_applyCompose sync.Once
func Get_applyCompose() gopurs_runtime.Value {
	once_applyCompose.Do(func() {
		applyCompose = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictApply1_2 gopurs_runtime.Value) gopurs_runtime.Value {
apply1_3_1 := dictApply1_2.PtrVal.(map[string]gopurs_runtime.Value)["apply"]
__local_var_4_2 := gopurs_runtime.Apply(dictApply1_2.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorCompose2_5_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_5)), v_6)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], apply1_3_1), v_6)), v1_7)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_5_3
})})
})
})
	})
	return applyCompose
}

var applicativeCompose gopurs_runtime.Value
var once_applicativeCompose sync.Once
func Get_applicativeCompose() gopurs_runtime.Value {
	once_applicativeCompose.Do(func() {
		applicativeCompose = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
Functor0_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictApplicative1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictApplicative1_3.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
apply1_5_3 := __local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["apply"]
__local_var_6_5 := gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorCompose2_7_6 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(__local_var_6_5.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_7)), v_8)
})
})})
applyCompose2_6_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], apply1_5_3), v_8)), v1_9)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_7_6
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Apply(dictApplicative1_3.PtrVal.(map[string]gopurs_runtime.Value)["pure"], x_7))
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose2_6_4
})})
})
})
	})
	return applicativeCompose
}

var altCompose gopurs_runtime.Value
var once_altCompose sync.Once
func Get_altCompose() gopurs_runtime.Value {
	once_altCompose.Do(func() {
		altCompose = gopurs_runtime.Func(func(dictAlt_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
functorCompose2_3_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(dictFunctor_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_3)), v_4)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["alt"], v_4), v1_5)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_3_1
})})
})
})
	})
	return altCompose
}

var plusCompose gopurs_runtime.Value
var once_plusCompose sync.Once
func Get_plusCompose() gopurs_runtime.Value {
	once_plusCompose.Do(func() {
		plusCompose = gopurs_runtime.Func(func(dictPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
empty_1_0 := dictPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
__local_var_2_1 := gopurs_runtime.Apply(dictPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
__local_var_3_2 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictFunctor_4 gopurs_runtime.Value) gopurs_runtime.Value {
functorCompose2_5_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(dictFunctor_4.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_5)), v_6)
})
})})
altCompose2_6_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["alt"], v_6), v1_7)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_5_3
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": empty_1_0, "Alt0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return altCompose2_6_4
})})
})
})
	})
	return plusCompose
}

var alternativeCompose gopurs_runtime.Value
var once_alternativeCompose sync.Once
func Get_alternativeCompose() gopurs_runtime.Value {
	once_alternativeCompose.Do(func() {
		alternativeCompose = gopurs_runtime.Func(func(dictAlternative_0 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeCompose1_1_0 := gopurs_runtime.Apply(Get_applicativeCompose(), gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}))
__local_var_2_1 := gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{})
empty_3_2 := __local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
__local_var_4_4 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
__local_var_5_5 := gopurs_runtime.Apply(__local_var_4_4.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
plusCompose1_4_3 := gopurs_runtime.Func(func(dictFunctor_6 gopurs_runtime.Value) gopurs_runtime.Value {
functorCompose2_7_6 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_5.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(dictFunctor_6.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_7)), v_8)
})
})})
altCompose2_8_7 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_4.PtrVal.(map[string]gopurs_runtime.Value)["alt"], v_8), v1_9)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_7_6
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": empty_3_2, "Alt0": gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return altCompose2_8_7
})})
})
return gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeCompose2_6_8 := gopurs_runtime.Apply(applicativeCompose1_1_0, dictApplicative_5)
plusCompose2_7_9 := gopurs_runtime.Apply(plusCompose1_4_3, gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_5.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeCompose2_6_8
}), "Plus1": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return plusCompose2_7_9
})})
})
})
	})
	return alternativeCompose
}


