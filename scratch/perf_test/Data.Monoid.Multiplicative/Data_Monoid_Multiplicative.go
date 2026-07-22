package Data_Monoid_Multiplicative

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Multiplicative gopurs_runtime.Value
var once_Multiplicative sync.Once
func Get_Multiplicative() gopurs_runtime.Value {
	once_Multiplicative.Do(func() {
		Multiplicative = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Multiplicative
}

var showMultiplicative gopurs_runtime.Value
var once_showMultiplicative sync.Once
func Get_showMultiplicative() gopurs_runtime.Value {
	once_showMultiplicative.Do(func() {
		showMultiplicative = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Multiplicative ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showMultiplicative
}

var semigroupMultiplicative gopurs_runtime.Value
var once_semigroupMultiplicative sync.Once
func Get_semigroupMultiplicative() gopurs_runtime.Value {
	once_semigroupMultiplicative.Do(func() {
		semigroupMultiplicative = gopurs_runtime.Func(func(dictSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["mul"], v_1), v1_2)
})
})})
})
	})
	return semigroupMultiplicative
}

var ordMultiplicative gopurs_runtime.Value
var once_ordMultiplicative sync.Once
func Get_ordMultiplicative() gopurs_runtime.Value {
	once_ordMultiplicative.Do(func() {
		ordMultiplicative = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordMultiplicative
}

var monoidMultiplicative gopurs_runtime.Value
var once_monoidMultiplicative sync.Once
func Get_monoidMultiplicative() gopurs_runtime.Value {
	once_monoidMultiplicative.Do(func() {
		monoidMultiplicative = gopurs_runtime.Func(func(dictSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupMultiplicative1_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["mul"], v_1), v1_2)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["one"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMultiplicative1_1_0
})})
})
	})
	return monoidMultiplicative
}

var functorMultiplicative gopurs_runtime.Value
var once_functorMultiplicative sync.Once
func Get_functorMultiplicative() gopurs_runtime.Value {
	once_functorMultiplicative.Do(func() {
		functorMultiplicative = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})
	})
	return functorMultiplicative
}

var eqMultiplicative gopurs_runtime.Value
var once_eqMultiplicative sync.Once
func Get_eqMultiplicative() gopurs_runtime.Value {
	once_eqMultiplicative.Do(func() {
		eqMultiplicative = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqMultiplicative
}

var eq1Multiplicative gopurs_runtime.Value
var once_eq1Multiplicative sync.Once
func Get_eq1Multiplicative() gopurs_runtime.Value {
	once_eq1Multiplicative.Do(func() {
		eq1Multiplicative = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
	})
	return eq1Multiplicative
}

var ord1Multiplicative gopurs_runtime.Value
var once_ord1Multiplicative sync.Once
func Get_ord1Multiplicative() gopurs_runtime.Value {
	once_ord1Multiplicative.Do(func() {
		ord1Multiplicative = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Multiplicative()
})})
	})
	return ord1Multiplicative
}

var boundedMultiplicative gopurs_runtime.Value
var once_boundedMultiplicative sync.Once
func Get_boundedMultiplicative() gopurs_runtime.Value {
	once_boundedMultiplicative.Do(func() {
		boundedMultiplicative = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBounded_0
})
	})
	return boundedMultiplicative
}

var applyMultiplicative gopurs_runtime.Value
var once_applyMultiplicative sync.Once
func Get_applyMultiplicative() gopurs_runtime.Value {
	once_applyMultiplicative.Do(func() {
		applyMultiplicative = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMultiplicative()
})})
	})
	return applyMultiplicative
}

var bindMultiplicative gopurs_runtime.Value
var once_bindMultiplicative sync.Once
func Get_bindMultiplicative() gopurs_runtime.Value {
	once_bindMultiplicative.Do(func() {
		bindMultiplicative = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMultiplicative()
})})
	})
	return bindMultiplicative
}

var applicativeMultiplicative gopurs_runtime.Value
var once_applicativeMultiplicative sync.Once
func Get_applicativeMultiplicative() gopurs_runtime.Value {
	once_applicativeMultiplicative.Do(func() {
		applicativeMultiplicative = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": Get_Multiplicative(), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMultiplicative()
})})
	})
	return applicativeMultiplicative
}

var monadMultiplicative gopurs_runtime.Value
var once_monadMultiplicative sync.Once
func Get_monadMultiplicative() gopurs_runtime.Value {
	once_monadMultiplicative.Do(func() {
		monadMultiplicative = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeMultiplicative()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindMultiplicative()
})})
	})
	return monadMultiplicative
}


