package Data_Tuple

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var Tuple gopurs_runtime.Value
var once_Tuple sync.Once
func Get_Tuple() gopurs_runtime.Value {
	once_Tuple.Do(func() {
		Tuple = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": value0, "value1": value1})
})
})
	})
	return Tuple
}

var uncurry gopurs_runtime.Value
var once_uncurry sync.Once
func Get_uncurry() gopurs_runtime.Value {
	once_uncurry.Do(func() {
		uncurry = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
	})
	return uncurry
}

var swap gopurs_runtime.Value
var once_swap sync.Once
func Get_swap() gopurs_runtime.Value {
	once_swap.Do(func() {
		swap = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value1": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
})
	})
	return swap
}

var snd gopurs_runtime.Value
var once_snd sync.Once
func Get_snd() gopurs_runtime.Value {
	once_snd.Do(func() {
		snd = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
	})
	return snd
}

var showTuple gopurs_runtime.Value
var once_showTuple sync.Once
func Get_showTuple() gopurs_runtime.Value {
	once_showTuple.Do(func() {
		showTuple = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Tuple ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow1_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Str(")")))))
})})
})
})
	})
	return showTuple
}

var semiringTuple gopurs_runtime.Value
var once_semiringTuple sync.Once
func Get_semiringTuple() gopurs_runtime.Value {
	once_semiringTuple.Do(func() {
		semiringTuple = gopurs_runtime.Func(func(dictSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
one_1_0 := dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["one"]
zero_2_1 := dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["zero"]
return gopurs_runtime.Func(func(dictSemiring1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"add": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["add"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring1_3.PtrVal.(map[string]gopurs_runtime.Value)["add"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "one": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": one_1_0, "value1": dictSemiring1_3.PtrVal.(map[string]gopurs_runtime.Value)["one"]}), "mul": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["mul"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring1_3.PtrVal.(map[string]gopurs_runtime.Value)["mul"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "zero": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": zero_2_1, "value1": dictSemiring1_3.PtrVal.(map[string]gopurs_runtime.Value)["zero"]})})
})
})
	})
	return semiringTuple
}

var semigroupoidTuple gopurs_runtime.Value
var once_semigroupoidTuple sync.Once
func Get_semigroupoidTuple() gopurs_runtime.Value {
	once_semigroupoidTuple.Do(func() {
		semigroupoidTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compose": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
})})
	})
	return semigroupoidTuple
}

var semigroupTuple gopurs_runtime.Value
var once_semigroupTuple sync.Once
func Get_semigroupTuple() gopurs_runtime.Value {
	once_semigroupTuple.Do(func() {
		semigroupTuple = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemigroup1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup1_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
})
})
	})
	return semigroupTuple
}

var ringTuple gopurs_runtime.Value
var once_ringTuple sync.Once
func Get_ringTuple() gopurs_runtime.Value {
	once_ringTuple.Do(func() {
		ringTuple = gopurs_runtime.Func(func(dictRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictRing_0.PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{})
one_2_1 := __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["one"]
zero_3_3 := __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["zero"]
semiringTuple1_3_2 := gopurs_runtime.Func(func(dictSemiring1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"add": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["add"], v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring1_4.PtrVal.(map[string]gopurs_runtime.Value)["add"], v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "one": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": one_2_1, "value1": dictSemiring1_4.PtrVal.(map[string]gopurs_runtime.Value)["one"]}), "mul": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["mul"], v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring1_4.PtrVal.(map[string]gopurs_runtime.Value)["mul"], v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "zero": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": zero_3_3, "value1": dictSemiring1_4.PtrVal.(map[string]gopurs_runtime.Value)["zero"]})})
})
return gopurs_runtime.Func(func(dictRing1_4 gopurs_runtime.Value) gopurs_runtime.Value {
semiringTuple2_5_4 := gopurs_runtime.Apply(semiringTuple1_3_2, gopurs_runtime.Apply(dictRing1_4.PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"sub": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictRing_0.PtrVal.(map[string]gopurs_runtime.Value)["sub"], v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictRing1_4.PtrVal.(map[string]gopurs_runtime.Value)["sub"], v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Semiring0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringTuple2_5_4
})})
})
})
	})
	return ringTuple
}

var monoidTuple gopurs_runtime.Value
var once_monoidTuple sync.Once
func Get_monoidTuple() gopurs_runtime.Value {
	once_monoidTuple.Do(func() {
		monoidTuple = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
__local_var_2_1 := gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictMonoid1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictMonoid1_3.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
semigroupTuple2_5_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": mempty_1_0, "value1": dictMonoid1_3.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupTuple2_5_3
})})
})
})
	})
	return monoidTuple
}

var heytingAlgebraTuple gopurs_runtime.Value
var once_heytingAlgebraTuple sync.Once
func Get_heytingAlgebraTuple() gopurs_runtime.Value {
	once_heytingAlgebraTuple.Do(func() {
		heytingAlgebraTuple = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
tt_1_0 := dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["tt"]
ff_2_1 := dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["ff"]
return gopurs_runtime.Func(func(dictHeytingAlgebra1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tt": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": tt_1_0, "value1": dictHeytingAlgebra1_3.PtrVal.(map[string]gopurs_runtime.Value)["tt"]}), "ff": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": ff_2_1, "value1": dictHeytingAlgebra1_3.PtrVal.(map[string]gopurs_runtime.Value)["ff"]}), "implies": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["implies"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra1_3.PtrVal.(map[string]gopurs_runtime.Value)["implies"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "conj": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["conj"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra1_3.PtrVal.(map[string]gopurs_runtime.Value)["conj"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "disj": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["disj"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra1_3.PtrVal.(map[string]gopurs_runtime.Value)["disj"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "not": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["not"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(dictHeytingAlgebra1_3.PtrVal.(map[string]gopurs_runtime.Value)["not"], v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})})
})
})
	})
	return heytingAlgebraTuple
}

var genericTuple gopurs_runtime.Value
var once_genericTuple sync.Once
func Get_genericTuple() gopurs_runtime.Value {
	once_genericTuple.Do(func() {
		genericTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"to": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
}), "from": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Product"), "value0": x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})})
	})
	return genericTuple
}

var functorTuple gopurs_runtime.Value
var once_functorTuple sync.Once
func Get_functorTuple() gopurs_runtime.Value {
	once_functorTuple.Do(func() {
		functorTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": m_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(f_0, m_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
	})
	return functorTuple
}

var invariantTuple gopurs_runtime.Value
var once_invariantTuple sync.Once
func Get_invariantTuple() gopurs_runtime.Value {
	once_invariantTuple.Do(func() {
		invariantTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"imap": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": m_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(f_0, m_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
})})
	})
	return invariantTuple
}

var fst gopurs_runtime.Value
var once_fst sync.Once
func Get_fst() gopurs_runtime.Value {
	once_fst.Do(func() {
		fst = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
})
	})
	return fst
}

var lazyTuple gopurs_runtime.Value
var once_lazyTuple sync.Once
func Get_lazyTuple() gopurs_runtime.Value {
	once_lazyTuple.Do(func() {
		lazyTuple = gopurs_runtime.Func(func(dictLazy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictLazy1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"defer": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(dictLazy_0.PtrVal.(map[string]gopurs_runtime.Value)["defer"], gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()).PtrVal.(map[string]gopurs_runtime.Value)["value0"]
})), "value1": gopurs_runtime.Apply(dictLazy1_1.PtrVal.(map[string]gopurs_runtime.Value)["defer"], gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()).PtrVal.(map[string]gopurs_runtime.Value)["value1"]
}))})
})})
})
})
	})
	return lazyTuple
}

var extendTuple gopurs_runtime.Value
var once_extendTuple sync.Once
func Get_extendTuple() gopurs_runtime.Value {
	once_extendTuple.Do(func() {
		extendTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(f_0, v_1)})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorTuple()
})})
	})
	return extendTuple
}

var eqTuple gopurs_runtime.Value
var once_eqTuple sync.Once
func Get_eqTuple() gopurs_runtime.Value {
	once_eqTuple.Do(func() {
		eqTuple = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq1_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})})
})
})
	})
	return eqTuple
}

var ordTuple gopurs_runtime.Value
var once_ordTuple sync.Once
func Get_ordTuple() gopurs_runtime.Value {
	once_ordTuple.Do(func() {
		ordTuple = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
eqTuple1_1_0 := gopurs_runtime.Apply(Get_eqTuple(), gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
eqTuple2_3_1 := gopurs_runtime.Apply(eqTuple1_1_0, gopurs_runtime.Apply(dictOrd1_2.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
} else {
if (gopurs_runtime.Bool(v_6_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd1_2.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
}
return __t3
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqTuple2_3_1
})})
})
})
	})
	return ordTuple
}

var eq1Tuple gopurs_runtime.Value
var once_eq1Tuple sync.Once
func Get_eq1Tuple() gopurs_runtime.Value {
	once_eq1Tuple.Do(func() {
		eq1Tuple = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eqTuple(), dictEq_0), dictEq1_1).PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
})
	})
	return eq1Tuple
}

var ord1Tuple gopurs_runtime.Value
var once_ord1Tuple sync.Once
func Get_ord1Tuple() gopurs_runtime.Value {
	once_ord1Tuple.Do(func() {
		ord1Tuple = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
ordTuple1_1_0 := gopurs_runtime.Apply(Get_ordTuple(), dictOrd_0)
__local_var_2_1 := gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
eq1Tuple1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eqTuple(), __local_var_2_1), dictEq1_3).PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(ordTuple1_1_0, dictOrd1_4).PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Tuple1_3_2
})})
})
	})
	return ord1Tuple
}

var curry gopurs_runtime.Value
var once_curry sync.Once
func Get_curry() gopurs_runtime.Value {
	once_curry.Do(func() {
		curry = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_1, "value1": b_2}))
})
})
})
	})
	return curry
}

var comonadTuple gopurs_runtime.Value
var once_comonadTuple sync.Once
func Get_comonadTuple() gopurs_runtime.Value {
	once_comonadTuple.Do(func() {
		comonadTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extract": Get_snd(), "Extend0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendTuple()
})})
	})
	return comonadTuple
}

var commutativeRingTuple gopurs_runtime.Value
var once_commutativeRingTuple sync.Once
func Get_commutativeRingTuple() gopurs_runtime.Value {
	once_commutativeRingTuple.Do(func() {
		commutativeRingTuple = gopurs_runtime.Func(func(dictCommutativeRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
ringTuple1_1_0 := gopurs_runtime.Apply(Get_ringTuple(), gopurs_runtime.Apply(dictCommutativeRing_0.PtrVal.(map[string]gopurs_runtime.Value)["Ring0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictCommutativeRing1_2 gopurs_runtime.Value) gopurs_runtime.Value {
ringTuple2_3_1 := gopurs_runtime.Apply(ringTuple1_1_0, gopurs_runtime.Apply(dictCommutativeRing1_2.PtrVal.(map[string]gopurs_runtime.Value)["Ring0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Ring0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ringTuple2_3_1
})})
})
})
	})
	return commutativeRingTuple
}

var boundedTuple gopurs_runtime.Value
var once_boundedTuple sync.Once
func Get_boundedTuple() gopurs_runtime.Value {
	once_boundedTuple.Do(func() {
		boundedTuple = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
top_1_0 := dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["top"]
bottom_2_1 := dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["bottom"]
ordTuple1_3_2 := gopurs_runtime.Apply(Get_ordTuple(), gopurs_runtime.Apply(dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictBounded1_4 gopurs_runtime.Value) gopurs_runtime.Value {
ordTuple2_5_3 := gopurs_runtime.Apply(ordTuple1_3_2, gopurs_runtime.Apply(dictBounded1_4.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"top": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": top_1_0, "value1": dictBounded1_4.PtrVal.(map[string]gopurs_runtime.Value)["top"]}), "bottom": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": bottom_2_1, "value1": dictBounded1_4.PtrVal.(map[string]gopurs_runtime.Value)["bottom"]}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return ordTuple2_5_3
})})
})
})
	})
	return boundedTuple
}

var booleanAlgebraTuple gopurs_runtime.Value
var once_booleanAlgebraTuple sync.Once
func Get_booleanAlgebraTuple() gopurs_runtime.Value {
	once_booleanAlgebraTuple.Do(func() {
		booleanAlgebraTuple = gopurs_runtime.Func(func(dictBooleanAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
heytingAlgebraTuple1_1_0 := gopurs_runtime.Apply(Get_heytingAlgebraTuple(), gopurs_runtime.Apply(dictBooleanAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["HeytingAlgebra0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictBooleanAlgebra1_2 gopurs_runtime.Value) gopurs_runtime.Value {
heytingAlgebraTuple2_3_1 := gopurs_runtime.Apply(heytingAlgebraTuple1_1_0, gopurs_runtime.Apply(dictBooleanAlgebra1_2.PtrVal.(map[string]gopurs_runtime.Value)["HeytingAlgebra0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"HeytingAlgebra0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraTuple2_3_1
})})
})
})
	})
	return booleanAlgebraTuple
}

var applyTuple gopurs_runtime.Value
var once_applyTuple sync.Once
func Get_applyTuple() gopurs_runtime.Value {
	once_applyTuple.Do(func() {
		applyTuple = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"], v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorTuple()
})})
})
	})
	return applyTuple
}

var bindTuple gopurs_runtime.Value
var once_bindTuple sync.Once
func Get_bindTuple() gopurs_runtime.Value {
	once_bindTuple.Do(func() {
		bindTuple = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
applyTuple1_1_0 := gopurs_runtime.Apply(Get_applyTuple(), dictSemigroup_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
v1_4_1 := gopurs_runtime.Apply(f_3, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v1_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyTuple1_1_0
})})
})
	})
	return bindTuple
}

var applicativeTuple gopurs_runtime.Value
var once_applicativeTuple sync.Once
func Get_applicativeTuple() gopurs_runtime.Value {
	once_applicativeTuple.Do(func() {
		applicativeTuple = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
applyTuple1_1_0 := gopurs_runtime.Apply(Get_applyTuple(), gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Apply(Get_Tuple(), dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]), "Apply0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyTuple1_1_0
})})
})
	})
	return applicativeTuple
}

var monadTuple gopurs_runtime.Value
var once_monadTuple sync.Once
func Get_monadTuple() gopurs_runtime.Value {
	once_monadTuple.Do(func() {
		monadTuple = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeTuple1_1_0 := gopurs_runtime.Apply(Get_applicativeTuple(), dictMonoid_0)
bindTuple1_2_1 := gopurs_runtime.Apply(Get_bindTuple(), gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeTuple1_1_0
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bindTuple1_2_1
})})
})
	})
	return monadTuple
}


