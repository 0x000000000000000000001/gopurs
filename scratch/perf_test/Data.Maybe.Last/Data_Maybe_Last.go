package Data_Maybe_Last

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Last gopurs_runtime.Value
var once_Last sync.Once
func Get_Last() gopurs_runtime.Value {
	once_Last.Do(func() {
		Last = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Last
}

var showLast gopurs_runtime.Value
var once_showLast sync.Once
func Get_showLast() gopurs_runtime.Value {
	once_showLast.Do(func() {
		showLast = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Last ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Maybe.Get_showMaybe(), dictShow_0).PtrVal.(map[string]gopurs_runtime.Value)["show"], v_1)), gopurs_runtime.Str(")")))
})})
})
	})
	return showLast
}

var semigroupLast gopurs_runtime.Value
var once_semigroupLast sync.Once
func Get_semigroupLast() gopurs_runtime.Value {
	once_semigroupLast.Do(func() {
		semigroupLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = v1_1
} else {
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = v_0
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})})
	})
	return semigroupLast
}

var ordLast gopurs_runtime.Value
var once_ordLast sync.Once
func Get_ordLast() gopurs_runtime.Value {
	once_ordLast.Do(func() {
		ordLast = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return ordLast
}

var ord1Last gopurs_runtime.Value
var once_ord1Last sync.Once
func Get_ord1Last() gopurs_runtime.Value {
	once_ord1Last.Do(func() {
		ord1Last = pkg_Data_Maybe.Get_ord1Maybe()
	})
	return ord1Last
}

var newtypeLast gopurs_runtime.Value
var once_newtypeLast sync.Once
func Get_newtypeLast() gopurs_runtime.Value {
	once_newtypeLast.Do(func() {
		newtypeLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeLast
}

var monoidLast gopurs_runtime.Value
var once_monoidLast sync.Once
func Get_monoidLast() gopurs_runtime.Value {
	once_monoidLast.Do(func() {
		monoidLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupLast()
})})
	})
	return monoidLast
}

var monadLast gopurs_runtime.Value
var once_monadLast sync.Once
func Get_monadLast() gopurs_runtime.Value {
	once_monadLast.Do(func() {
		monadLast = pkg_Data_Maybe.Get_monadMaybe()
	})
	return monadLast
}

var invariantLast gopurs_runtime.Value
var once_invariantLast sync.Once
func Get_invariantLast() gopurs_runtime.Value {
	once_invariantLast.Do(func() {
		invariantLast = pkg_Data_Maybe.Get_invariantMaybe()
	})
	return invariantLast
}

var functorLast gopurs_runtime.Value
var once_functorLast sync.Once
func Get_functorLast() gopurs_runtime.Value {
	once_functorLast.Do(func() {
		functorLast = pkg_Data_Maybe.Get_functorMaybe()
	})
	return functorLast
}

var extendLast gopurs_runtime.Value
var once_extendLast sync.Once
func Get_extendLast() gopurs_runtime.Value {
	once_extendLast.Do(func() {
		extendLast = pkg_Data_Maybe.Get_extendMaybe()
	})
	return extendLast
}

var eqLast gopurs_runtime.Value
var once_eqLast sync.Once
func Get_eqLast() gopurs_runtime.Value {
	once_eqLast.Do(func() {
		eqLast = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return eqLast
}

var eq1Last gopurs_runtime.Value
var once_eq1Last sync.Once
func Get_eq1Last() gopurs_runtime.Value {
	once_eq1Last.Do(func() {
		eq1Last = pkg_Data_Maybe.Get_eq1Maybe()
	})
	return eq1Last
}

var boundedLast gopurs_runtime.Value
var once_boundedLast sync.Once
func Get_boundedLast() gopurs_runtime.Value {
	once_boundedLast.Do(func() {
		boundedLast = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Maybe.Get_boundedMaybe(), dictBounded_0)
})
	})
	return boundedLast
}

var bindLast gopurs_runtime.Value
var once_bindLast sync.Once
func Get_bindLast() gopurs_runtime.Value {
	once_bindLast.Do(func() {
		bindLast = pkg_Data_Maybe.Get_bindMaybe()
	})
	return bindLast
}

var applyLast gopurs_runtime.Value
var once_applyLast sync.Once
func Get_applyLast() gopurs_runtime.Value {
	once_applyLast.Do(func() {
		applyLast = pkg_Data_Maybe.Get_applyMaybe()
	})
	return applyLast
}

var applicativeLast gopurs_runtime.Value
var once_applicativeLast sync.Once
func Get_applicativeLast() gopurs_runtime.Value {
	once_applicativeLast.Do(func() {
		applicativeLast = pkg_Data_Maybe.Get_applicativeMaybe()
	})
	return applicativeLast
}

var altLast gopurs_runtime.Value
var once_altLast sync.Once
func Get_altLast() gopurs_runtime.Value {
	once_altLast.Do(func() {
		altLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": Get_semigroupLast().PtrVal.(map[string]gopurs_runtime.Value)["append"], "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
})})
	})
	return altLast
}

var plusLast gopurs_runtime.Value
var once_plusLast sync.Once
func Get_plusLast() gopurs_runtime.Value {
	once_plusLast.Do(func() {
		plusLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altLast()
})})
	})
	return plusLast
}

var alternativeLast gopurs_runtime.Value
var once_alternativeLast sync.Once
func Get_alternativeLast() gopurs_runtime.Value {
	once_alternativeLast.Do(func() {
		alternativeLast = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applicativeMaybe()
}), "Plus1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusLast()
})})
	})
	return alternativeLast
}


