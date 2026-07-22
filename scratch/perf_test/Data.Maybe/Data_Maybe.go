package Data_Maybe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var Nothing gopurs_runtime.Value
var once_Nothing sync.Once
func Get_Nothing() gopurs_runtime.Value {
	once_Nothing.Do(func() {
		Nothing = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
	})
	return Nothing
}

var Just gopurs_runtime.Value
var once_Just sync.Once
func Get_Just() gopurs_runtime.Value {
	once_Just.Do(func() {
		Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": value0})
})
	})
	return Just
}

var showMaybe gopurs_runtime.Value
var once_showMaybe sync.Once
func Get_showMaybe() gopurs_runtime.Value {
	once_showMaybe.Do(func() {
		showMaybe = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Just ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Str(")")))
} else {
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Str("Nothing")
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})})
})
	})
	return showMaybe
}

var semigroupMaybe gopurs_runtime.Value
var once_semigroupMaybe sync.Once
func Get_semigroupMaybe() gopurs_runtime.Value {
	once_semigroupMaybe.Do(func() {
		semigroupMaybe = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = v1_2
} else {
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = v_1
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t0
})
})})
})
	})
	return semigroupMaybe
}

var optional gopurs_runtime.Value
var once_optional sync.Once
func Get_optional() gopurs_runtime.Value {
	once_optional.Do(func() {
		optional = gopurs_runtime.Func(func(dictAlt_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_Just()), a_2)), gopurs_runtime.Apply(dictApplicative_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})))
})
})
})
	})
	return optional
}

var monoidMaybe gopurs_runtime.Value
var once_monoidMaybe sync.Once
func Get_monoidMaybe() gopurs_runtime.Value {
	once_monoidMaybe.Do(func() {
		monoidMaybe = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupMaybe1_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = v1_2
} else {
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = v_1
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMaybe1_1_0
})})
})
	})
	return monoidMaybe
}

var maybe_prime gopurs_runtime.Value
var once_maybe_prime sync.Once
func Get_maybe_prime() gopurs_runtime.Value {
	once_maybe_prime.Do(func() {
		maybe_prime = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(v1_1, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
})
	})
	return maybe_prime
}

var maybe gopurs_runtime.Value
var once_maybe sync.Once
func Get_maybe() gopurs_runtime.Value {
	once_maybe.Do(func() {
		maybe = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = v_0
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(v1_1, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
})
	})
	return maybe
}

var isNothing gopurs_runtime.Value
var once_isNothing sync.Once
func Get_isNothing() gopurs_runtime.Value {
	once_isNothing.Do(func() {
		isNothing = gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(true)
} else {
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(false)
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
	})
	return isNothing
}

var isJust gopurs_runtime.Value
var once_isJust sync.Once
func Get_isJust() gopurs_runtime.Value {
	once_isJust.Do(func() {
		isJust = gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(false)
} else {
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(true)
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
	})
	return isJust
}

var genericMaybe gopurs_runtime.Value
var once_genericMaybe sync.Once
func Get_genericMaybe() gopurs_runtime.Value {
	once_genericMaybe.Do(func() {
		genericMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"to": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inl")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inr")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
}), "from": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inl"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NoArguments")})})
} else {
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inr"), "value0": x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})})
	})
	return genericMaybe
}

var functorMaybe gopurs_runtime.Value
var once_functorMaybe sync.Once
func Get_functorMaybe() gopurs_runtime.Value {
	once_functorMaybe.Do(func() {
		functorMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(v_0, v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t0
})
})})
	})
	return functorMaybe
}

var invariantMaybe gopurs_runtime.Value
var once_invariantMaybe sync.Once
func Get_invariantMaybe() gopurs_runtime.Value {
	once_invariantMaybe.Do(func() {
		invariantMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"imap": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(f_0, v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t0
})
})
})})
	})
	return invariantMaybe
}

var fromMaybe_prime gopurs_runtime.Value
var once_fromMaybe_prime sync.Once
func Get_fromMaybe_prime() gopurs_runtime.Value {
	once_fromMaybe_prime.Do(func() {
		fromMaybe_prime = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_maybe_prime(), a_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
	})
	return fromMaybe_prime
}

var fromMaybe gopurs_runtime.Value
var once_fromMaybe sync.Once
func Get_fromMaybe() gopurs_runtime.Value {
	once_fromMaybe.Do(func() {
		fromMaybe = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = a_0
} else {
if (gopurs_runtime.Bool(v2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = v2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
	})
	return fromMaybe
}

var fromJust gopurs_runtime.Value
var once_fromJust sync.Once
func Get_fromJust() gopurs_runtime.Value {
	once_fromJust.Do(func() {
		fromJust = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
return __t0
})
})
	})
	return fromJust
}

var extendMaybe gopurs_runtime.Value
var once_extendMaybe sync.Once
func Get_extendMaybe() gopurs_runtime.Value {
	once_extendMaybe.Do(func() {
		extendMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(v_0, v1_1)})
}
return __t0
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMaybe()
})})
	})
	return extendMaybe
}

var eqMaybe gopurs_runtime.Value
var once_eqMaybe sync.Once
func Get_eqMaybe() gopurs_runtime.Value {
	once_eqMaybe.Do(func() {
		eqMaybe = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")
} else {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(y_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t0
})
})})
})
	})
	return eqMaybe
}

var ordMaybe gopurs_runtime.Value
var once_ordMaybe sync.Once
func Get_ordMaybe() gopurs_runtime.Value {
	once_ordMaybe.Do(func() {
		ordMaybe = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
eqMaybe1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(y_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")
} else {
__t2 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(y_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t2
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(y_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
} else {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
__t3 = __t4
} else {
if (gopurs_runtime.Bool(y_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0)).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t3
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_2_1
})})
})
	})
	return ordMaybe
}

var eq1Maybe gopurs_runtime.Value
var once_eq1Maybe sync.Once
func Get_eq1Maybe() gopurs_runtime.Value {
	once_eq1Maybe.Do(func() {
		eq1Maybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")
} else {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(y_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t0
})
})
})})
	})
	return eq1Maybe
}

var ord1Maybe gopurs_runtime.Value
var once_ord1Maybe sync.Once
func Get_ord1Maybe() gopurs_runtime.Value {
	once_ord1Maybe.Do(func() {
		ord1Maybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(y_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
__t0 = __t1
} else {
if (gopurs_runtime.Bool(y_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t0
})
})
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Maybe()
})})
	})
	return ord1Maybe
}

var boundedMaybe gopurs_runtime.Value
var once_boundedMaybe sync.Once
func Get_boundedMaybe() gopurs_runtime.Value {
	once_boundedMaybe.Do(func() {
		boundedMaybe = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
eqMaybe1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t4 = gopurs_runtime.Bool(y_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")
} else {
__t4 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(y_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t4
})
})})
ordMaybe1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
var __t6 gopurs_runtime.Value
if (gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
} else {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
__t5 = __t6
} else {
if (gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0)).IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t5
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_3_3
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"top": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["top"]}), "bottom": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ordMaybe1_3_2
})})
})
	})
	return boundedMaybe
}

var applyMaybe gopurs_runtime.Value
var once_applyMaybe sync.Once
func Get_applyMaybe() gopurs_runtime.Value {
	once_applyMaybe.Do(func() {
		applyMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
__t0 = __t1
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMaybe()
})})
	})
	return applyMaybe
}

var bindMaybe gopurs_runtime.Value
var once_bindMaybe sync.Once
func Get_bindMaybe() gopurs_runtime.Value {
	once_bindMaybe.Do(func() {
		bindMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(v1_1, v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMaybe()
})})
	})
	return bindMaybe
}

var semiringMaybe gopurs_runtime.Value
var once_semiringMaybe sync.Once
func Get_semiringMaybe() gopurs_runtime.Value {
	once_semiringMaybe.Do(func() {
		semiringMaybe = gopurs_runtime.Func(func(dictSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"zero": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}), "one": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["one"]}), "add": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = v1_2
} else {
if (gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = v_1
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(v1_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["add"], v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t0
})
}), "mul": gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["mul"], x_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
})})
})
	})
	return semiringMaybe
}

var applicativeMaybe gopurs_runtime.Value
var once_applicativeMaybe sync.Once
func Get_applicativeMaybe() gopurs_runtime.Value {
	once_applicativeMaybe.Do(func() {
		applicativeMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": Get_Just(), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMaybe()
})})
	})
	return applicativeMaybe
}

var monadMaybe gopurs_runtime.Value
var once_monadMaybe sync.Once
func Get_monadMaybe() gopurs_runtime.Value {
	once_monadMaybe.Do(func() {
		monadMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeMaybe()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindMaybe()
})})
	})
	return monadMaybe
}

var altMaybe gopurs_runtime.Value
var once_altMaybe sync.Once
func Get_altMaybe() gopurs_runtime.Value {
	once_altMaybe.Do(func() {
		altMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = v1_1
} else {
__t0 = v_0
}
return __t0
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMaybe()
})})
	})
	return altMaybe
}

var plusMaybe gopurs_runtime.Value
var once_plusMaybe sync.Once
func Get_plusMaybe() gopurs_runtime.Value {
	once_plusMaybe.Do(func() {
		plusMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altMaybe()
})})
	})
	return plusMaybe
}

var alternativeMaybe gopurs_runtime.Value
var once_alternativeMaybe sync.Once
func Get_alternativeMaybe() gopurs_runtime.Value {
	once_alternativeMaybe.Do(func() {
		alternativeMaybe = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeMaybe()
}), "Plus1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusMaybe()
})})
	})
	return alternativeMaybe
}


