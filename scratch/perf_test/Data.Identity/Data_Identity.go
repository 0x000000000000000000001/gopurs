package Data_Identity

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Identity gopurs_runtime.Value
var once_Identity sync.Once
func Get_Identity() gopurs_runtime.Value {
	once_Identity.Do(func() {
		Identity = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Identity
}

var showIdentity gopurs_runtime.Value
var once_showIdentity sync.Once
func Get_showIdentity() gopurs_runtime.Value {
	once_showIdentity.Do(func() {
		showIdentity = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Identity ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showIdentity
}

var semiringIdentity gopurs_runtime.Value
var once_semiringIdentity sync.Once
func Get_semiringIdentity() gopurs_runtime.Value {
	once_semiringIdentity.Do(func() {
		semiringIdentity = gopurs_runtime.Func(func(dictSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictSemiring_0
})
	})
	return semiringIdentity
}

var semigroupIdentity gopurs_runtime.Value
var once_semigroupIdentity sync.Once
func Get_semigroupIdentity() gopurs_runtime.Value {
	once_semigroupIdentity.Do(func() {
		semigroupIdentity = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictSemigroup_0
})
	})
	return semigroupIdentity
}

var ringIdentity gopurs_runtime.Value
var once_ringIdentity sync.Once
func Get_ringIdentity() gopurs_runtime.Value {
	once_ringIdentity.Do(func() {
		ringIdentity = gopurs_runtime.Func(func(dictRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictRing_0
})
	})
	return ringIdentity
}

var ordIdentity gopurs_runtime.Value
var once_ordIdentity sync.Once
func Get_ordIdentity() gopurs_runtime.Value {
	once_ordIdentity.Do(func() {
		ordIdentity = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordIdentity
}

var newtypeIdentity gopurs_runtime.Value
var once_newtypeIdentity sync.Once
func Get_newtypeIdentity() gopurs_runtime.Value {
	once_newtypeIdentity.Do(func() {
		newtypeIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeIdentity
}

var monoidIdentity gopurs_runtime.Value
var once_monoidIdentity sync.Once
func Get_monoidIdentity() gopurs_runtime.Value {
	once_monoidIdentity.Do(func() {
		monoidIdentity = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonoid_0
})
	})
	return monoidIdentity
}

var lazyIdentity gopurs_runtime.Value
var once_lazyIdentity sync.Once
func Get_lazyIdentity() gopurs_runtime.Value {
	once_lazyIdentity.Do(func() {
		lazyIdentity = gopurs_runtime.Func(func(dictLazy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictLazy_0
})
	})
	return lazyIdentity
}

var heytingAlgebraIdentity gopurs_runtime.Value
var once_heytingAlgebraIdentity sync.Once
func Get_heytingAlgebraIdentity() gopurs_runtime.Value {
	once_heytingAlgebraIdentity.Do(func() {
		heytingAlgebraIdentity = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictHeytingAlgebra_0
})
	})
	return heytingAlgebraIdentity
}

var functorIdentity gopurs_runtime.Value
var once_functorIdentity sync.Once
func Get_functorIdentity() gopurs_runtime.Value {
	once_functorIdentity.Do(func() {
		functorIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
})})
	})
	return functorIdentity
}

var invariantIdentity gopurs_runtime.Value
var once_invariantIdentity sync.Once
func Get_invariantIdentity() gopurs_runtime.Value {
	once_invariantIdentity.Do(func() {
		invariantIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"imap": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_2)
})
})
})})
	})
	return invariantIdentity
}

var extendIdentity gopurs_runtime.Value
var once_extendIdentity sync.Once
func Get_extendIdentity() gopurs_runtime.Value {
	once_extendIdentity.Do(func() {
		extendIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorIdentity()
})})
	})
	return extendIdentity
}

var euclideanRingIdentity gopurs_runtime.Value
var once_euclideanRingIdentity sync.Once
func Get_euclideanRingIdentity() gopurs_runtime.Value {
	once_euclideanRingIdentity.Do(func() {
		euclideanRingIdentity = gopurs_runtime.Func(func(dictEuclideanRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEuclideanRing_0
})
	})
	return euclideanRingIdentity
}

var eqIdentity gopurs_runtime.Value
var once_eqIdentity sync.Once
func Get_eqIdentity() gopurs_runtime.Value {
	once_eqIdentity.Do(func() {
		eqIdentity = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqIdentity
}

var eq1Identity gopurs_runtime.Value
var once_eq1Identity sync.Once
func Get_eq1Identity() gopurs_runtime.Value {
	once_eq1Identity.Do(func() {
		eq1Identity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})})
	})
	return eq1Identity
}

var ord1Identity gopurs_runtime.Value
var once_ord1Identity sync.Once
func Get_ord1Identity() gopurs_runtime.Value {
	once_ord1Identity.Do(func() {
		ord1Identity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Identity()
})})
	})
	return ord1Identity
}

var comonadIdentity gopurs_runtime.Value
var once_comonadIdentity sync.Once
func Get_comonadIdentity() gopurs_runtime.Value {
	once_comonadIdentity.Do(func() {
		comonadIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extract": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), "Extend0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendIdentity()
})})
	})
	return comonadIdentity
}

var commutativeRingIdentity gopurs_runtime.Value
var once_commutativeRingIdentity sync.Once
func Get_commutativeRingIdentity() gopurs_runtime.Value {
	once_commutativeRingIdentity.Do(func() {
		commutativeRingIdentity = gopurs_runtime.Func(func(dictCommutativeRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictCommutativeRing_0
})
	})
	return commutativeRingIdentity
}

var boundedIdentity gopurs_runtime.Value
var once_boundedIdentity sync.Once
func Get_boundedIdentity() gopurs_runtime.Value {
	once_boundedIdentity.Do(func() {
		boundedIdentity = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBounded_0
})
	})
	return boundedIdentity
}

var booleanAlgebraIdentity gopurs_runtime.Value
var once_booleanAlgebraIdentity sync.Once
func Get_booleanAlgebraIdentity() gopurs_runtime.Value {
	once_booleanAlgebraIdentity.Do(func() {
		booleanAlgebraIdentity = gopurs_runtime.Func(func(dictBooleanAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBooleanAlgebra_0
})
	})
	return booleanAlgebraIdentity
}

var applyIdentity gopurs_runtime.Value
var once_applyIdentity sync.Once
func Get_applyIdentity() gopurs_runtime.Value {
	once_applyIdentity.Do(func() {
		applyIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorIdentity()
})})
	})
	return applyIdentity
}

var bindIdentity gopurs_runtime.Value
var once_bindIdentity sync.Once
func Get_bindIdentity() gopurs_runtime.Value {
	once_bindIdentity.Do(func() {
		bindIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyIdentity()
})})
	})
	return bindIdentity
}

var applicativeIdentity gopurs_runtime.Value
var once_applicativeIdentity sync.Once
func Get_applicativeIdentity() gopurs_runtime.Value {
	once_applicativeIdentity.Do(func() {
		applicativeIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": Get_Identity(), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyIdentity()
})})
	})
	return applicativeIdentity
}

var monadIdentity gopurs_runtime.Value
var once_monadIdentity sync.Once
func Get_monadIdentity() gopurs_runtime.Value {
	once_monadIdentity.Do(func() {
		monadIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeIdentity()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindIdentity()
})})
	})
	return monadIdentity
}

var altIdentity gopurs_runtime.Value
var once_altIdentity sync.Once
func Get_altIdentity() gopurs_runtime.Value {
	once_altIdentity.Do(func() {
		altIdentity = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorIdentity()
})})
	})
	return altIdentity
}


