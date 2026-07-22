package Data_Functor_Costar

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var Costar gopurs_runtime.Value
var once_Costar sync.Once
func Get_Costar() gopurs_runtime.Value {
	once_Costar.Do(func() {
		Costar = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Costar
}

var semigroupoidCostar gopurs_runtime.Value
var once_semigroupoidCostar sync.Once
func Get_semigroupoidCostar() gopurs_runtime.Value {
	once_semigroupoidCostar.Do(func() {
		semigroupoidCostar = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compose": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(gopurs_runtime.Apply(dictExtend_0.PtrVal.(map[string]gopurs_runtime.Value)["extend"], v1_2), w_3))
})
})
})})
})
	})
	return semigroupoidCostar
}

var profunctorCostar gopurs_runtime.Value
var once_profunctorCostar sync.Once
func Get_profunctorCostar() gopurs_runtime.Value {
	once_profunctorCostar.Do(func() {
		profunctorCostar = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"dimap": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_1)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(__local_var_4_0, x_5)))
})
})
})
})})
})
	})
	return profunctorCostar
}

var strongCostar gopurs_runtime.Value
var once_strongCostar sync.Once
func Get_strongCostar() gopurs_runtime.Value {
	once_strongCostar.Do(func() {
		strongCostar = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
profunctorCostar1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"dimap": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2)
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_3, gopurs_runtime.Apply(v_4, gopurs_runtime.Apply(__local_var_5_2, x_6)))
})
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"first": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Tuple.Get_fst()), x_4)), "value1": gopurs_runtime.Apply(dictComonad_0.PtrVal.(map[string]gopurs_runtime.Value)["extract"], x_4).PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
}), "second": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(dictComonad_0.PtrVal.(map[string]gopurs_runtime.Value)["extract"], x_4).PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Tuple.Get_snd()), x_4))})
})
}), "Profunctor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorCostar1_2_1
})})
})
	})
	return strongCostar
}

var newtypeCostar gopurs_runtime.Value
var once_newtypeCostar sync.Once
func Get_newtypeCostar() gopurs_runtime.Value {
	once_newtypeCostar.Do(func() {
		newtypeCostar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeCostar
}

var hoistCostar gopurs_runtime.Value
var once_hoistCostar sync.Once
func Get_hoistCostar() gopurs_runtime.Value {
	once_hoistCostar.Do(func() {
		hoistCostar = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
})
	})
	return hoistCostar
}

var functorCostar gopurs_runtime.Value
var once_functorCostar sync.Once
func Get_functorCostar() gopurs_runtime.Value {
	once_functorCostar.Do(func() {
		functorCostar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
})
})
})})
	})
	return functorCostar
}

var invariantCostar gopurs_runtime.Value
var once_invariantCostar sync.Once
func Get_invariantCostar() gopurs_runtime.Value {
	once_invariantCostar.Do(func() {
		invariantCostar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"imap": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_2, x_3))
})
})
})
})})
	})
	return invariantCostar
}

var distributiveCostar gopurs_runtime.Value
var once_distributiveCostar sync.Once
func Get_distributiveCostar() gopurs_runtime.Value {
	once_distributiveCostar.Do(func() {
		distributiveCostar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"distribute": gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, a_2)
})), f_1)
})
})
}), "collect": gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(Get_distributiveCostar().PtrVal.(map[string]gopurs_runtime.Value)["distribute"], dictFunctor_0)
__local_var_3_1 := gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_1)
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorCostar()
})})
	})
	return distributiveCostar
}

var closedCostar gopurs_runtime.Value
var once_closedCostar sync.Once
func Get_closedCostar() gopurs_runtime.Value {
	once_closedCostar.Do(func() {
		closedCostar = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
profunctorCostar1_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"dimap": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_1)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(__local_var_4_1, x_5)))
})
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"closed": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_5, x_4)
})), g_3))
})
})
}), "Profunctor0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorCostar1_1_0
})})
})
	})
	return closedCostar
}

var categoryCostar gopurs_runtime.Value
var once_categoryCostar sync.Once
func Get_categoryCostar() gopurs_runtime.Value {
	once_categoryCostar.Do(func() {
		categoryCostar = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictComonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{})
semigroupoidCostar1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compose": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(w_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["extend"], v1_3), w_4))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"identity": dictComonad_0.PtrVal.(map[string]gopurs_runtime.Value)["extract"], "Semigroupoid0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupoidCostar1_2_1
})})
})
	})
	return categoryCostar
}

var bifunctorCostar gopurs_runtime.Value
var once_bifunctorCostar sync.Once
func Get_bifunctorCostar() gopurs_runtime.Value {
	once_bifunctorCostar.Do(func() {
		bifunctorCostar = gopurs_runtime.Func(func(dictContravariant_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(dictContravariant_0.PtrVal.(map[string]gopurs_runtime.Value)["cmap"], f_1)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(__local_var_4_0, x_5)))
})
})
})
})})
})
	})
	return bifunctorCostar
}

var applyCostar gopurs_runtime.Value
var once_applyCostar sync.Once
func Get_applyCostar() gopurs_runtime.Value {
	once_applyCostar.Do(func() {
		applyCostar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, a_2), gopurs_runtime.Apply(v1_1, a_2))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorCostar()
})})
	})
	return applyCostar
}

var bindCostar gopurs_runtime.Value
var once_bindCostar sync.Once
func Get_bindCostar() gopurs_runtime.Value {
	once_bindCostar.Do(func() {
		bindCostar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(v_0, x_2)), x_2)
})
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyCostar()
})})
	})
	return bindCostar
}

var applicativeCostar gopurs_runtime.Value
var once_applicativeCostar sync.Once
func Get_applicativeCostar() gopurs_runtime.Value {
	once_applicativeCostar.Do(func() {
		applicativeCostar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyCostar()
})})
	})
	return applicativeCostar
}

var monadCostar gopurs_runtime.Value
var once_monadCostar sync.Once
func Get_monadCostar() gopurs_runtime.Value {
	once_monadCostar.Do(func() {
		monadCostar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeCostar()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindCostar()
})})
	})
	return monadCostar
}


