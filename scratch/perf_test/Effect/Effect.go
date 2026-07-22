package Effect

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var monadEffect gopurs_runtime.Value
var once_monadEffect sync.Once
func Get_monadEffect() gopurs_runtime.Value {
	once_monadEffect.Do(func() {
		monadEffect = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeEffect()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindEffect()
})})
	})
	return monadEffect
}

var bindEffect gopurs_runtime.Value
var once_bindEffect sync.Once
func Get_bindEffect() gopurs_runtime.Value {
	once_bindEffect.Do(func() {
		bindEffect = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": Get_bindE(), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
})})
	})
	return bindEffect
}

var applyEffect gopurs_runtime.Value
var once_applyEffect sync.Once
func Get_applyEffect() gopurs_runtime.Value {
	once_applyEffect.Do(func() {
		applyEffect = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bindE(), f_0), gopurs_runtime.Func(func(f_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bindE(), a_1), gopurs_runtime.Func(func(a_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeEffect().PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Apply(f_prime_2, a_prime_3))
}))
}))
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEffect()
})})
	})
	return applyEffect
}

var applicativeEffect gopurs_runtime.Value
var once_applicativeEffect sync.Once
func Get_applicativeEffect() gopurs_runtime.Value {
	once_applicativeEffect.Do(func() {
		applicativeEffect = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": Get_pureE(), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
})})
	})
	return applicativeEffect
}

var functorEffect gopurs_runtime.Value
var once_functorEffect sync.Once
func Get_functorEffect() gopurs_runtime.Value {
	once_functorEffect.Do(func() {
		functorEffect = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_applyEffect().PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(Get_pureE(), f_0)), a_1)
})
})})
	})
	return functorEffect
}

var lift2 gopurs_runtime.Value
var once_lift2 sync.Once
func Get_lift2() gopurs_runtime.Value {
	once_lift2.Do(func() {
		lift2 = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_applyEffect().PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Get_applyEffect().PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(Get_pureE(), f_0)), a_1)), b_2)
})
})
})
	})
	return lift2
}

var semigroupEffect gopurs_runtime.Value
var once_semigroupEffect sync.Once
func Get_semigroupEffect() gopurs_runtime.Value {
	once_semigroupEffect.Do(func() {
		semigroupEffect = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Apply(Get_lift2(), dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"])})
})
	})
	return semigroupEffect
}

var monoidEffect gopurs_runtime.Value
var once_monoidEffect sync.Once
func Get_monoidEffect() gopurs_runtime.Value {
	once_monoidEffect.Do(func() {
		monoidEffect = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Apply(Get_pureE(), dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Apply(Get_lift2(), gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"])})
})})
})
	})
	return monoidEffect
}

func Get_bindE() gopurs_runtime.Value {
	return BindE
}

func Get_forE() gopurs_runtime.Value {
	return ForE
}

func Get_foreachE() gopurs_runtime.Value {
	return ForeachE
}

func Get_pureE() gopurs_runtime.Value {
	return PureE
}

func Get_untilE() gopurs_runtime.Value {
	return UntilE
}

func Get_whileE() gopurs_runtime.Value {
	return WhileE
}
