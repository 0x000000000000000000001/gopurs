package Data_Monoid_Alternate

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Alternate gopurs_runtime.Value
var once_Alternate sync.Once
func Get_Alternate() gopurs_runtime.Value {
	once_Alternate.Do(func() {
		Alternate = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Alternate
}

var showAlternate gopurs_runtime.Value
var once_showAlternate sync.Once
func Get_showAlternate() gopurs_runtime.Value {
	once_showAlternate.Do(func() {
		showAlternate = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Alternate ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showAlternate
}

var semigroupAlternate gopurs_runtime.Value
var once_semigroupAlternate sync.Once
func Get_semigroupAlternate() gopurs_runtime.Value {
	once_semigroupAlternate.Do(func() {
		semigroupAlternate = gopurs_runtime.Func(func(dictAlt_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["alt"], v_1), v1_2)
})
})})
})
	})
	return semigroupAlternate
}

var plusAlternate gopurs_runtime.Value
var once_plusAlternate sync.Once
func Get_plusAlternate() gopurs_runtime.Value {
	once_plusAlternate.Do(func() {
		plusAlternate = gopurs_runtime.Func(func(dictPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictPlus_0
})
	})
	return plusAlternate
}

var ordAlternate gopurs_runtime.Value
var once_ordAlternate sync.Once
func Get_ordAlternate() gopurs_runtime.Value {
	once_ordAlternate.Do(func() {
		ordAlternate = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordAlternate
}

var ord1Alternate gopurs_runtime.Value
var once_ord1Alternate sync.Once
func Get_ord1Alternate() gopurs_runtime.Value {
	once_ord1Alternate.Do(func() {
		ord1Alternate = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd1_0
})
	})
	return ord1Alternate
}

var newtypeAlternate gopurs_runtime.Value
var once_newtypeAlternate sync.Once
func Get_newtypeAlternate() gopurs_runtime.Value {
	once_newtypeAlternate.Do(func() {
		newtypeAlternate = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeAlternate
}

var monoidAlternate gopurs_runtime.Value
var once_monoidAlternate sync.Once
func Get_monoidAlternate() gopurs_runtime.Value {
	once_monoidAlternate.Do(func() {
		monoidAlternate = gopurs_runtime.Func(func(dictPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
semigroupAlternate1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["alt"], v_2), v1_3)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": dictPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["empty"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAlternate1_2_1
})})
})
	})
	return monoidAlternate
}

var monadAlternate gopurs_runtime.Value
var once_monadAlternate sync.Once
func Get_monadAlternate() gopurs_runtime.Value {
	once_monadAlternate.Do(func() {
		monadAlternate = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonad_0
})
	})
	return monadAlternate
}

var functorAlternate gopurs_runtime.Value
var once_functorAlternate sync.Once
func Get_functorAlternate() gopurs_runtime.Value {
	once_functorAlternate.Do(func() {
		functorAlternate = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictFunctor_0
})
	})
	return functorAlternate
}

var extendAlternate gopurs_runtime.Value
var once_extendAlternate sync.Once
func Get_extendAlternate() gopurs_runtime.Value {
	once_extendAlternate.Do(func() {
		extendAlternate = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictExtend_0
})
	})
	return extendAlternate
}

var eqAlternate gopurs_runtime.Value
var once_eqAlternate sync.Once
func Get_eqAlternate() gopurs_runtime.Value {
	once_eqAlternate.Do(func() {
		eqAlternate = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqAlternate
}

var eq1Alternate gopurs_runtime.Value
var once_eq1Alternate sync.Once
func Get_eq1Alternate() gopurs_runtime.Value {
	once_eq1Alternate.Do(func() {
		eq1Alternate = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq1_0
})
	})
	return eq1Alternate
}

var comonadAlternate gopurs_runtime.Value
var once_comonadAlternate sync.Once
func Get_comonadAlternate() gopurs_runtime.Value {
	once_comonadAlternate.Do(func() {
		comonadAlternate = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictComonad_0
})
	})
	return comonadAlternate
}

var boundedAlternate gopurs_runtime.Value
var once_boundedAlternate sync.Once
func Get_boundedAlternate() gopurs_runtime.Value {
	once_boundedAlternate.Do(func() {
		boundedAlternate = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBounded_0
})
	})
	return boundedAlternate
}

var bindAlternate gopurs_runtime.Value
var once_bindAlternate sync.Once
func Get_bindAlternate() gopurs_runtime.Value {
	once_bindAlternate.Do(func() {
		bindAlternate = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBind_0
})
	})
	return bindAlternate
}

var applyAlternate gopurs_runtime.Value
var once_applyAlternate sync.Once
func Get_applyAlternate() gopurs_runtime.Value {
	once_applyAlternate.Do(func() {
		applyAlternate = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictApply_0
})
	})
	return applyAlternate
}

var applicativeAlternate gopurs_runtime.Value
var once_applicativeAlternate sync.Once
func Get_applicativeAlternate() gopurs_runtime.Value {
	once_applicativeAlternate.Do(func() {
		applicativeAlternate = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictApplicative_0
})
	})
	return applicativeAlternate
}

var alternativeAlternate gopurs_runtime.Value
var once_alternativeAlternate sync.Once
func Get_alternativeAlternate() gopurs_runtime.Value {
	once_alternativeAlternate.Do(func() {
		alternativeAlternate = gopurs_runtime.Func(func(dictAlternative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictAlternative_0
})
	})
	return alternativeAlternate
}

var altAlternate gopurs_runtime.Value
var once_altAlternate sync.Once
func Get_altAlternate() gopurs_runtime.Value {
	once_altAlternate.Do(func() {
		altAlternate = gopurs_runtime.Func(func(dictAlt_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictAlt_0
})
	})
	return altAlternate
}


