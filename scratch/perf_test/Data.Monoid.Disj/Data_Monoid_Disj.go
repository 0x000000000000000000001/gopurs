package Data_Monoid_Disj

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Disj gopurs_runtime.Value
var once_Disj sync.Once
func Get_Disj() gopurs_runtime.Value {
	once_Disj.Do(func() {
		Disj = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Disj
}

var showDisj gopurs_runtime.Value
var once_showDisj sync.Once
func Get_showDisj() gopurs_runtime.Value {
	once_showDisj.Do(func() {
		showDisj = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Disj ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showDisj
}

var semiringDisj gopurs_runtime.Value
var once_semiringDisj sync.Once
func Get_semiringDisj() gopurs_runtime.Value {
	once_semiringDisj.Do(func() {
		semiringDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"zero": dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["ff"], "one": dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["tt"], "add": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["disj"], v_1), v1_2)
})
}), "mul": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["conj"], v_1), v1_2)
})
})})
})
	})
	return semiringDisj
}

var semigroupDisj gopurs_runtime.Value
var once_semigroupDisj sync.Once
func Get_semigroupDisj() gopurs_runtime.Value {
	once_semigroupDisj.Do(func() {
		semigroupDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["disj"], v_1), v1_2)
})
})})
})
	})
	return semigroupDisj
}

var ordDisj gopurs_runtime.Value
var once_ordDisj sync.Once
func Get_ordDisj() gopurs_runtime.Value {
	once_ordDisj.Do(func() {
		ordDisj = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordDisj
}

var monoidDisj gopurs_runtime.Value
var once_monoidDisj sync.Once
func Get_monoidDisj() gopurs_runtime.Value {
	once_monoidDisj.Do(func() {
		monoidDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupDisj1_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["disj"], v_1), v1_2)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["ff"], "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_1_0
})})
})
	})
	return monoidDisj
}

var functorDisj gopurs_runtime.Value
var once_functorDisj sync.Once
func Get_functorDisj() gopurs_runtime.Value {
	once_functorDisj.Do(func() {
		functorDisj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})
	})
	return functorDisj
}

var eqDisj gopurs_runtime.Value
var once_eqDisj sync.Once
func Get_eqDisj() gopurs_runtime.Value {
	once_eqDisj.Do(func() {
		eqDisj = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqDisj
}

var eq1Disj gopurs_runtime.Value
var once_eq1Disj sync.Once
func Get_eq1Disj() gopurs_runtime.Value {
	once_eq1Disj.Do(func() {
		eq1Disj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
	})
	return eq1Disj
}

var ord1Disj gopurs_runtime.Value
var once_ord1Disj sync.Once
func Get_ord1Disj() gopurs_runtime.Value {
	once_ord1Disj.Do(func() {
		ord1Disj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Disj()
})})
	})
	return ord1Disj
}

var boundedDisj gopurs_runtime.Value
var once_boundedDisj sync.Once
func Get_boundedDisj() gopurs_runtime.Value {
	once_boundedDisj.Do(func() {
		boundedDisj = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBounded_0
})
	})
	return boundedDisj
}

var applyDisj gopurs_runtime.Value
var once_applyDisj sync.Once
func Get_applyDisj() gopurs_runtime.Value {
	once_applyDisj.Do(func() {
		applyDisj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorDisj()
})})
	})
	return applyDisj
}

var bindDisj gopurs_runtime.Value
var once_bindDisj sync.Once
func Get_bindDisj() gopurs_runtime.Value {
	once_bindDisj.Do(func() {
		bindDisj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDisj()
})})
	})
	return bindDisj
}

var applicativeDisj gopurs_runtime.Value
var once_applicativeDisj sync.Once
func Get_applicativeDisj() gopurs_runtime.Value {
	once_applicativeDisj.Do(func() {
		applicativeDisj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": Get_Disj(), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDisj()
})})
	})
	return applicativeDisj
}

var monadDisj gopurs_runtime.Value
var once_monadDisj sync.Once
func Get_monadDisj() gopurs_runtime.Value {
	once_monadDisj.Do(func() {
		monadDisj = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeDisj()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindDisj()
})})
	})
	return monadDisj
}


