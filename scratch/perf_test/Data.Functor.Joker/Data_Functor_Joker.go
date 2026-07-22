package Data_Functor_Joker

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var Joker gopurs_runtime.Value
var once_Joker sync.Once
func Get_Joker() gopurs_runtime.Value {
	once_Joker.Do(func() {
		Joker = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Joker
}

var showJoker gopurs_runtime.Value
var once_showJoker sync.Once
func Get_showJoker() gopurs_runtime.Value {
	once_showJoker.Do(func() {
		showJoker = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Joker ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showJoker
}

var profunctorJoker gopurs_runtime.Value
var once_profunctorJoker sync.Once
func Get_profunctorJoker() gopurs_runtime.Value {
	once_profunctorJoker.Do(func() {
		profunctorJoker = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"dimap": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], g_2), v1_3)
})
})
})})
})
	})
	return profunctorJoker
}

var ordJoker gopurs_runtime.Value
var once_ordJoker sync.Once
func Get_ordJoker() gopurs_runtime.Value {
	once_ordJoker.Do(func() {
		ordJoker = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordJoker
}

var newtypeJoker gopurs_runtime.Value
var once_newtypeJoker sync.Once
func Get_newtypeJoker() gopurs_runtime.Value {
	once_newtypeJoker.Do(func() {
		newtypeJoker = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeJoker
}

var hoistJoker gopurs_runtime.Value
var once_hoistJoker sync.Once
func Get_hoistJoker() gopurs_runtime.Value {
	once_hoistJoker.Do(func() {
		hoistJoker = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v_1)
})
})
	})
	return hoistJoker
}

var functorJoker gopurs_runtime.Value
var once_functorJoker sync.Once
func Get_functorJoker() gopurs_runtime.Value {
	once_functorJoker.Do(func() {
		functorJoker = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_1), v_2)
})
})})
})
	})
	return functorJoker
}

var eqJoker gopurs_runtime.Value
var once_eqJoker sync.Once
func Get_eqJoker() gopurs_runtime.Value {
	once_eqJoker.Do(func() {
		eqJoker = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqJoker
}

var choiceJoker gopurs_runtime.Value
var once_choiceJoker sync.Once
func Get_choiceJoker() gopurs_runtime.Value {
	once_choiceJoker.Do(func() {
		choiceJoker = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
profunctorJoker1_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"dimap": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], g_2), v1_3)
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"left": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Either.Get_Left()), v_2)
}), "right": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Either.Get_Right()), v_2)
}), "Profunctor0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorJoker1_1_0
})})
})
	})
	return choiceJoker
}

var bifunctorJoker gopurs_runtime.Value
var once_bifunctorJoker sync.Once
func Get_bifunctorJoker() gopurs_runtime.Value {
	once_bifunctorJoker.Do(func() {
		bifunctorJoker = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], g_2), v1_3)
})
})
})})
})
	})
	return bifunctorJoker
}

var biapplyJoker gopurs_runtime.Value
var once_biapplyJoker sync.Once
func Get_biapplyJoker() gopurs_runtime.Value {
	once_biapplyJoker.Do(func() {
		biapplyJoker = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
bifunctorJoker1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], g_3), v1_4)
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"biapply": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_3), v1_4)
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoker1_2_1
})})
})
	})
	return biapplyJoker
}

var biapplicativeJoker gopurs_runtime.Value
var once_biapplicativeJoker sync.Once
func Get_biapplicativeJoker() gopurs_runtime.Value {
	once_biapplicativeJoker.Do(func() {
		biapplicativeJoker = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
bifunctorJoker1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], g_4), v1_5)
})
})
})})
biapplyJoker1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"biapply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_4), v1_5)
})
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoker1_3_3
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bipure": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], b_5)
})
}), "Biapply0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyJoker1_3_2
})})
})
	})
	return biapplicativeJoker
}

var applyJoker gopurs_runtime.Value
var once_applyJoker sync.Once
func Get_applyJoker() gopurs_runtime.Value {
	once_applyJoker.Do(func() {
		applyJoker = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorJoker1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2), v_3)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_3), v1_4)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_2_1
})})
})
	})
	return applyJoker
}

var bindJoker gopurs_runtime.Value
var once_bindJoker sync.Once
func Get_bindJoker() gopurs_runtime.Value {
	once_bindJoker.Do(func() {
		bindJoker = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorJoker1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_3), v_4)
})
})})
applyJoker1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_4), v1_5)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_3_3
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(amb_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_4), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(amb_5, x_6))
}))
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyJoker1_3_2
})})
})
	})
	return bindJoker
}

var applicativeJoker gopurs_runtime.Value
var once_applicativeJoker sync.Once
func Get_applicativeJoker() gopurs_runtime.Value {
	once_applicativeJoker.Do(func() {
		applicativeJoker = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorJoker1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_3), v_4)
})
})})
applyJoker1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_4), v1_5)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_3_3
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], x_4)
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyJoker1_3_2
})})
})
	})
	return applicativeJoker
}

var monadJoker gopurs_runtime.Value
var once_monadJoker sync.Once
func Get_monadJoker() gopurs_runtime.Value {
	once_monadJoker.Do(func() {
		monadJoker = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_3_3 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorJoker1_4_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_3.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4), v_5)
})
})})
applyJoker1_5_5 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["apply"], v_5), v1_6)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_4_4
})})
applicativeJoker1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], x_6)
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyJoker1_5_5
})})
bindJoker1_4_6 := gopurs_runtime.Apply(Get_bindJoker(), gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeJoker1_3_2
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bindJoker1_4_6
})})
})
	})
	return monadJoker
}


