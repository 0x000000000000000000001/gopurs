package Data_Functor_App

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var App gopurs_runtime.Value
var once_App sync.Once
func Get_App() gopurs_runtime.Value {
	once_App.Do(func() {
		App = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return App
}

var showApp gopurs_runtime.Value
var once_showApp sync.Once
func Get_showApp() gopurs_runtime.Value {
	once_showApp.Do(func() {
		showApp = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(App ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showApp
}

var semigroupApp gopurs_runtime.Value
var once_semigroupApp sync.Once
func Get_semigroupApp() gopurs_runtime.Value {
	once_semigroupApp.Do(func() {
		semigroupApp = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
append1_2_0 := dictSemigroup_1.PtrVal.(map[string]gopurs_runtime.Value)["append"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], append1_2_0), v_3)), v1_4)
})
})})
})
})
	})
	return semigroupApp
}

var plusApp gopurs_runtime.Value
var once_plusApp sync.Once
func Get_plusApp() gopurs_runtime.Value {
	once_plusApp.Do(func() {
		plusApp = gopurs_runtime.Func(func(dictPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictPlus_0
})
	})
	return plusApp
}

var newtypeApp gopurs_runtime.Value
var once_newtypeApp sync.Once
func Get_newtypeApp() gopurs_runtime.Value {
	once_newtypeApp.Do(func() {
		newtypeApp = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeApp
}

var monoidApp gopurs_runtime.Value
var once_monoidApp sync.Once
func Get_monoidApp() gopurs_runtime.Value {
	once_monoidApp.Do(func() {
		monoidApp = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
append1_3_1 := gopurs_runtime.Apply(dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"]
semigroupApp2_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], append1_3_1), v_4)), v1_5)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupApp2_4_2
})})
})
})
	})
	return monoidApp
}

var monadPlusApp gopurs_runtime.Value
var once_monadPlusApp sync.Once
func Get_monadPlusApp() gopurs_runtime.Value {
	once_monadPlusApp.Do(func() {
		monadPlusApp = gopurs_runtime.Func(func(dictMonadPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonadPlus_0
})
	})
	return monadPlusApp
}

var monadApp gopurs_runtime.Value
var once_monadApp sync.Once
func Get_monadApp() gopurs_runtime.Value {
	once_monadApp.Do(func() {
		monadApp = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonad_0
})
	})
	return monadApp
}

var lazyApp gopurs_runtime.Value
var once_lazyApp sync.Once
func Get_lazyApp() gopurs_runtime.Value {
	once_lazyApp.Do(func() {
		lazyApp = gopurs_runtime.Func(func(dictLazy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictLazy_0
})
	})
	return lazyApp
}

var hoistLowerApp gopurs_runtime.Value
var once_hoistLowerApp sync.Once
func Get_hoistLowerApp() gopurs_runtime.Value {
	once_hoistLowerApp.Do(func() {
		hoistLowerApp = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return hoistLowerApp
}

var hoistLiftApp gopurs_runtime.Value
var once_hoistLiftApp sync.Once
func Get_hoistLiftApp() gopurs_runtime.Value {
	once_hoistLiftApp.Do(func() {
		hoistLiftApp = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return hoistLiftApp
}

var hoistApp gopurs_runtime.Value
var once_hoistApp sync.Once
func Get_hoistApp() gopurs_runtime.Value {
	once_hoistApp.Do(func() {
		hoistApp = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v_1)
})
})
	})
	return hoistApp
}

var functorApp gopurs_runtime.Value
var once_functorApp sync.Once
func Get_functorApp() gopurs_runtime.Value {
	once_functorApp.Do(func() {
		functorApp = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictFunctor_0
})
	})
	return functorApp
}

var extendApp gopurs_runtime.Value
var once_extendApp sync.Once
func Get_extendApp() gopurs_runtime.Value {
	once_extendApp.Do(func() {
		extendApp = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictExtend_0
})
	})
	return extendApp
}

var eqApp gopurs_runtime.Value
var once_eqApp sync.Once
func Get_eqApp() gopurs_runtime.Value {
	once_eqApp.Do(func() {
		eqApp = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
eq11_2_0 := gopurs_runtime.Apply(dictEq1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_1)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(eq11_2_0, x_3), y_4)
})
})})
})
})
	})
	return eqApp
}

var ordApp gopurs_runtime.Value
var once_ordApp sync.Once
func Get_ordApp() gopurs_runtime.Value {
	once_ordApp.Do(func() {
		ordApp = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_3_1 := gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["compare1"], dictOrd_2)
eq11_4_2 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], gopurs_runtime.Apply(dictOrd_2.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
eqApp2_5_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(eq11_4_2, x_5), y_6)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(compare11_3_1, x_6), y_7)
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqApp2_5_3
})})
})
})
	})
	return ordApp
}

var eq1App gopurs_runtime.Value
var once_eq1App sync.Once
func Get_eq1App() gopurs_runtime.Value {
	once_eq1App.Do(func() {
		eq1App = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictEq1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_1)
})})
})
	})
	return eq1App
}

var ord1App gopurs_runtime.Value
var once_ord1App sync.Once
func Get_ord1App() gopurs_runtime.Value {
	once_ord1App.Do(func() {
		ord1App = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{})
eq1App1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_3)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_5_3 := gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["compare1"], dictOrd_4)
eq11_6_4 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], gopurs_runtime.Apply(dictOrd_4.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
eqApp2_7_5 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(eq11_6_4, x_7), y_8)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(compare11_5_3, x_8), y_9)
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return eqApp2_7_5
})}).PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1App1_3_2
})})
})
	})
	return ord1App
}

var comonadApp gopurs_runtime.Value
var once_comonadApp sync.Once
func Get_comonadApp() gopurs_runtime.Value {
	once_comonadApp.Do(func() {
		comonadApp = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictComonad_0
})
	})
	return comonadApp
}

var bindApp gopurs_runtime.Value
var once_bindApp sync.Once
func Get_bindApp() gopurs_runtime.Value {
	once_bindApp.Do(func() {
		bindApp = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBind_0
})
	})
	return bindApp
}

var applyApp gopurs_runtime.Value
var once_applyApp sync.Once
func Get_applyApp() gopurs_runtime.Value {
	once_applyApp.Do(func() {
		applyApp = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictApply_0
})
	})
	return applyApp
}

var applicativeApp gopurs_runtime.Value
var once_applicativeApp sync.Once
func Get_applicativeApp() gopurs_runtime.Value {
	once_applicativeApp.Do(func() {
		applicativeApp = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictApplicative_0
})
	})
	return applicativeApp
}

var alternativeApp gopurs_runtime.Value
var once_alternativeApp sync.Once
func Get_alternativeApp() gopurs_runtime.Value {
	once_alternativeApp.Do(func() {
		alternativeApp = gopurs_runtime.Func(func(dictAlternative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictAlternative_0
})
	})
	return alternativeApp
}

var altApp gopurs_runtime.Value
var once_altApp sync.Once
func Get_altApp() gopurs_runtime.Value {
	once_altApp.Do(func() {
		altApp = gopurs_runtime.Func(func(dictAlt_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictAlt_0
})
	})
	return altApp
}


