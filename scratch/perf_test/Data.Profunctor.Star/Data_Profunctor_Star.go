package Data_Profunctor_Star

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var Star gopurs_runtime.Value
var once_Star sync.Once
func Get_Star() gopurs_runtime.Value {
	once_Star.Do(func() {
		Star = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Star
}

var semigroupoidStar gopurs_runtime.Value
var once_semigroupoidStar sync.Once
func Get_semigroupoidStar() gopurs_runtime.Value {
	once_semigroupoidStar.Do(func() {
		semigroupoidStar = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compose": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(v1_2, x_3)), v_1)
})
})
})})
})
	})
	return semigroupoidStar
}

var profunctorStar gopurs_runtime.Value
var once_profunctorStar sync.Once
func Get_profunctorStar() gopurs_runtime.Value {
	once_profunctorStar.Do(func() {
		profunctorStar = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"dimap": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], g_2)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(f_1, x_5)))
})
})
})
})})
})
	})
	return profunctorStar
}

var strongStar gopurs_runtime.Value
var once_strongStar sync.Once
func Get_strongStar() gopurs_runtime.Value {
	once_strongStar.Do(func() {
		strongStar = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
profunctorStar1_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"dimap": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], g_2)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(f_1, x_5)))
})
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"first": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v2_5, "value1": __local_var_4_2})
})), gopurs_runtime.Apply(v_2, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
})
}), "second": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(v_2, v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
}), "Profunctor0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorStar1_1_0
})})
})
	})
	return strongStar
}

var newtypeStar gopurs_runtime.Value
var once_newtypeStar sync.Once
func Get_newtypeStar() gopurs_runtime.Value {
	once_newtypeStar.Do(func() {
		newtypeStar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeStar
}

var invariantStar gopurs_runtime.Value
var once_invariantStar sync.Once
func Get_invariantStar() gopurs_runtime.Value {
	once_invariantStar.Do(func() {
		invariantStar = gopurs_runtime.Func(func(dictInvariant_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"imap": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictInvariant_0.PtrVal.(map[string]gopurs_runtime.Value)["imap"], f_1), g_2)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Apply(v_3, x_5))
})
})
})
})})
})
	})
	return invariantStar
}

var hoistStar gopurs_runtime.Value
var once_hoistStar sync.Once
func Get_hoistStar() gopurs_runtime.Value {
	once_hoistStar.Do(func() {
		hoistStar = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
})
})
})
	})
	return hoistStar
}

var functorStar gopurs_runtime.Value
var once_functorStar sync.Once
func Get_functorStar() gopurs_runtime.Value {
	once_functorStar.Do(func() {
		functorStar = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_1)
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Apply(v_2, x_4))
})
})
})})
})
	})
	return functorStar
}

var distributiveStar gopurs_runtime.Value
var once_distributiveStar sync.Once
func Get_distributiveStar() gopurs_runtime.Value {
	once_distributiveStar.Do(func() {
		distributiveStar = gopurs_runtime.Func(func(dictDistributive_0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
__local_var_1_0 := gopurs_runtime.Apply(dictDistributive_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStar1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(v_3, x_5))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"distribute": gopurs_runtime.Func(func(dictFunctor_3 gopurs_runtime.Value) gopurs_runtime.Value {
collect1_4_3 := gopurs_runtime.Apply(dictDistributive_0.PtrVal.(map[string]gopurs_runtime.Value)["collect"], dictFunctor_3)
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(collect1_4_3, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_7, a_6)
})), f_5)
})
})
}), "collect": gopurs_runtime.Func(func(dictFunctor_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_distributiveStar(), dictDistributive_0).PtrVal.(map[string]gopurs_runtime.Value)["distribute"], dictFunctor_3)
__local_var_6_5 := gopurs_runtime.Apply(dictFunctor_3.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4)
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(__local_var_6_5, x_7))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_2_1
})})
}
}()
})
	})
	return distributiveStar
}

var closedStar gopurs_runtime.Value
var once_closedStar sync.Once
func Get_closedStar() gopurs_runtime.Value {
	once_closedStar.Do(func() {
		closedStar = gopurs_runtime.Func(func(dictDistributive_0 gopurs_runtime.Value) gopurs_runtime.Value {
distribute_1_0 := gopurs_runtime.Apply(dictDistributive_0.PtrVal.(map[string]gopurs_runtime.Value)["distribute"], pkg_Data_Functor.Get_functorFn())
__local_var_2_1 := gopurs_runtime.Apply(dictDistributive_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
profunctorStar1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"dimap": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], g_4)
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_3, gopurs_runtime.Apply(v_5, gopurs_runtime.Apply(f_3, x_7)))
})
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"closed": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(distribute_1_0, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Apply(g_5, x_6))
}))
})
}), "Profunctor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorStar1_3_2
})})
})
	})
	return closedStar
}

var choiceStar gopurs_runtime.Value
var once_choiceStar sync.Once
func Get_choiceStar() gopurs_runtime.Value {
	once_choiceStar.Do(func() {
		choiceStar = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
profunctorStar1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"dimap": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], g_3)
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(v_4, gopurs_runtime.Apply(f_2, x_6)))
})
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"left": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Either.Get_Left())
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(v_3, v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
if (gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}))
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t4
})
}), "right": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_5 := gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Either.Get_Right())
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}))
} else {
if (gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t6 = gopurs_runtime.Apply(__local_var_4_5, gopurs_runtime.Apply(v_3, v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
} else {
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t6
})
}), "Profunctor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorStar1_2_1
})})
})
	})
	return choiceStar
}

var categoryStar gopurs_runtime.Value
var once_categoryStar sync.Once
func Get_categoryStar() gopurs_runtime.Value {
	once_categoryStar.Do(func() {
		categoryStar = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
semigroupoidStar1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compose": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(v1_3, x_4)), v_2)
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"identity": gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], "Semigroupoid0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupoidStar1_2_1
})})
})
	})
	return categoryStar
}

var applyStar gopurs_runtime.Value
var once_applyStar sync.Once
func Get_applyStar() gopurs_runtime.Value {
	once_applyStar.Do(func() {
		applyStar = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStar1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(v_3, x_5))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_3, a_5)), gopurs_runtime.Apply(v1_4, a_5))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_2_1
})})
})
	})
	return applyStar
}

var bindStar gopurs_runtime.Value
var once_bindStar sync.Once
func Get_bindStar() gopurs_runtime.Value {
	once_bindStar.Do(func() {
		bindStar = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStar1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_3)
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
})
})})
applyStar1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_4, a_6)), gopurs_runtime.Apply(v1_5, a_6))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_3_3
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(v_4, x_6)), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_5, a_7), x_6)
}))
})
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_3_2
})})
})
	})
	return bindStar
}

var applicativeStar gopurs_runtime.Value
var once_applicativeStar sync.Once
func Get_applicativeStar() gopurs_runtime.Value {
	once_applicativeStar.Do(func() {
		applicativeStar = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStar1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_3)
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
})
})})
applyStar1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_4, a_6)), gopurs_runtime.Apply(v1_5, a_6))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_3_3
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], a_4)
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_3_2
})})
})
	})
	return applicativeStar
}

var monadStar gopurs_runtime.Value
var once_monadStar sync.Once
func Get_monadStar() gopurs_runtime.Value {
	once_monadStar.Do(func() {
		monadStar = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_3_3 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStar1_4_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_5 := gopurs_runtime.Apply(__local_var_3_3.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4)
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5, gopurs_runtime.Apply(v_5, x_7))
})
})
})})
applyStar1_5_6 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_5, a_7)), gopurs_runtime.Apply(v1_6, a_7))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_4_4
})})
applicativeStar1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], a_6)
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_5_6
})})
bindStar1_4_7 := gopurs_runtime.Apply(Get_bindStar(), gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeStar1_3_2
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bindStar1_4_7
})})
})
	})
	return monadStar
}

var altStar gopurs_runtime.Value
var once_altStar sync.Once
func Get_altStar() gopurs_runtime.Value {
	once_altStar.Do(func() {
		altStar = gopurs_runtime.Func(func(dictAlt_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStar1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(v_3, x_5))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(v_3, a_5)), gopurs_runtime.Apply(v1_4, a_5))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_2_1
})})
})
	})
	return altStar
}

var plusStar gopurs_runtime.Value
var once_plusStar sync.Once
func Get_plusStar() gopurs_runtime.Value {
	once_plusStar.Do(func() {
		plusStar = gopurs_runtime.Func(func(dictPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
empty_1_0 := dictPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
__local_var_2_1 := gopurs_runtime.Apply(dictPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
__local_var_3_2 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStar1_4_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_5 := gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4)
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5, gopurs_runtime.Apply(v_5, x_7))
})
})
})})
altStar1_4_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(v_5, a_7)), gopurs_runtime.Apply(v1_6, a_7))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_4_4
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_1_0
}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return altStar1_4_3
})})
})
	})
	return plusStar
}

var alternativeStar gopurs_runtime.Value
var once_alternativeStar sync.Once
func Get_alternativeStar() gopurs_runtime.Value {
	once_alternativeStar.Do(func() {
		alternativeStar = gopurs_runtime.Func(func(dictAlternative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_3_3 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStar1_4_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_5 := gopurs_runtime.Apply(__local_var_3_3.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4)
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5, gopurs_runtime.Apply(v_5, x_7))
})
})
})})
applyStar1_5_6 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(v_5, a_7)), gopurs_runtime.Apply(v1_6, a_7))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_4_4
})})
applicativeStar1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], a_6)
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_5_6
})})
__local_var_4_7 := gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{})
empty_5_8 := __local_var_4_7.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
__local_var_6_10 := gopurs_runtime.Apply(__local_var_4_7.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
__local_var_7_11 := gopurs_runtime.Apply(__local_var_6_10.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStar1_8_12 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_10_13 := gopurs_runtime.Apply(__local_var_7_11.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_8)
return gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_13, gopurs_runtime.Apply(v_9, x_11))
})
})
})})
altStar1_9_14 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_6_10.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(v_9, a_11)), gopurs_runtime.Apply(v1_10, a_11))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_8_12
})})
plusStar1_6_9 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_5_8
}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return altStar1_9_14
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeStar1_3_2
}), "Plus1": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return plusStar1_6_9
})})
})
	})
	return alternativeStar
}

var monadPlusStar gopurs_runtime.Value
var once_monadPlusStar sync.Once
func Get_monadPlusStar() gopurs_runtime.Value {
	once_monadPlusStar.Do(func() {
		monadPlusStar = gopurs_runtime.Func(func(dictMonadPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadStar1_1_0 := gopurs_runtime.Apply(Get_monadStar(), gopurs_runtime.Apply(dictMonadPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}))
alternativeStar1_2_1 := gopurs_runtime.Apply(Get_alternativeStar(), gopurs_runtime.Apply(dictMonadPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["Alternative1"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStar1_1_0
}), "Alternative1": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeStar1_2_1
})})
})
	})
	return monadPlusStar
}


