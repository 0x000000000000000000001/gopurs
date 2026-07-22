package Data_Monoid_Dual

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Dual gopurs_runtime.Value
var once_Dual sync.Once
func Get_Dual() gopurs_runtime.Value {
	once_Dual.Do(func() {
		Dual = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Dual
}

var showDual gopurs_runtime.Value
var once_showDual sync.Once
func Get_showDual() gopurs_runtime.Value {
	once_showDual.Do(func() {
		showDual = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Dual ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showDual
}

var semigroupDual gopurs_runtime.Value
var once_semigroupDual sync.Once
func Get_semigroupDual() gopurs_runtime.Value {
	once_semigroupDual.Do(func() {
		semigroupDual = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v1_2), v_1)
})
})})
})
	})
	return semigroupDual
}

var ordDual gopurs_runtime.Value
var once_ordDual sync.Once
func Get_ordDual() gopurs_runtime.Value {
	once_ordDual.Do(func() {
		ordDual = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordDual
}

var monoidDual gopurs_runtime.Value
var once_monoidDual sync.Once
func Get_monoidDual() gopurs_runtime.Value {
	once_monoidDual.Do(func() {
		monoidDual = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
semigroupDual1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v1_3), v_2)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDual1_2_1
})})
})
	})
	return monoidDual
}

var functorDual gopurs_runtime.Value
var once_functorDual sync.Once
func Get_functorDual() gopurs_runtime.Value {
	once_functorDual.Do(func() {
		functorDual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})
	})
	return functorDual
}

var eqDual gopurs_runtime.Value
var once_eqDual sync.Once
func Get_eqDual() gopurs_runtime.Value {
	once_eqDual.Do(func() {
		eqDual = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqDual
}

var eq1Dual gopurs_runtime.Value
var once_eq1Dual sync.Once
func Get_eq1Dual() gopurs_runtime.Value {
	once_eq1Dual.Do(func() {
		eq1Dual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
	})
	return eq1Dual
}

var ord1Dual gopurs_runtime.Value
var once_ord1Dual sync.Once
func Get_ord1Dual() gopurs_runtime.Value {
	once_ord1Dual.Do(func() {
		ord1Dual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Dual()
})})
	})
	return ord1Dual
}

var boundedDual gopurs_runtime.Value
var once_boundedDual sync.Once
func Get_boundedDual() gopurs_runtime.Value {
	once_boundedDual.Do(func() {
		boundedDual = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBounded_0
})
	})
	return boundedDual
}

var applyDual gopurs_runtime.Value
var once_applyDual sync.Once
func Get_applyDual() gopurs_runtime.Value {
	once_applyDual.Do(func() {
		applyDual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorDual()
})})
	})
	return applyDual
}

var bindDual gopurs_runtime.Value
var once_bindDual sync.Once
func Get_bindDual() gopurs_runtime.Value {
	once_bindDual.Do(func() {
		bindDual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDual()
})})
	})
	return bindDual
}

var applicativeDual gopurs_runtime.Value
var once_applicativeDual sync.Once
func Get_applicativeDual() gopurs_runtime.Value {
	once_applicativeDual.Do(func() {
		applicativeDual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": Get_Dual(), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDual()
})})
	})
	return applicativeDual
}

var monadDual gopurs_runtime.Value
var once_monadDual sync.Once
func Get_monadDual() gopurs_runtime.Value {
	once_monadDual.Do(func() {
		monadDual = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeDual()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindDual()
})})
	})
	return monadDual
}


