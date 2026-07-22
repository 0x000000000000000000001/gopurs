package Data_Either

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var Left gopurs_runtime.Value
var once_Left sync.Once
func Get_Left() gopurs_runtime.Value {
	once_Left.Do(func() {
		Left = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": value0})
})
	})
	return Left
}

var Right gopurs_runtime.Value
var once_Right sync.Once
func Get_Right() gopurs_runtime.Value {
	once_Right.Do(func() {
		Right = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": value0})
})
	})
	return Right
}

var showEither gopurs_runtime.Value
var once_showEither sync.Once
func Get_showEither() gopurs_runtime.Value {
	once_showEither.Do(func() {
		showEither = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Left ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Str(")")))
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Right ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow1_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Str(")")))
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})})
})
})
	})
	return showEither
}

var note_prime gopurs_runtime.Value
var once_note_prime sync.Once
func Get_note_prime() gopurs_runtime.Value {
	once_note_prime.Do(func() {
		note_prime = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())})
} else {
if (gopurs_runtime.Bool(v2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": v2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
	})
	return note_prime
}

var note gopurs_runtime.Value
var once_note sync.Once
func Get_note() gopurs_runtime.Value {
	once_note.Do(func() {
		note = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": a_0})
} else {
if (gopurs_runtime.Bool(v2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": v2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
	})
	return note
}

var genericEither gopurs_runtime.Value
var once_genericEither sync.Once
func Get_genericEither() gopurs_runtime.Value {
	once_genericEither.Do(func() {
		genericEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"to": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inl")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Inr")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
}), "from": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inl"), "value0": x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Inr"), "value0": x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})})
	})
	return genericEither
}

var functorEither gopurs_runtime.Value
var once_functorEither sync.Once
func Get_functorEither() gopurs_runtime.Value {
	once_functorEither.Do(func() {
		functorEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": m_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(m_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(f_0, m_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})})
	})
	return functorEither
}

var invariantEither gopurs_runtime.Value
var once_invariantEither sync.Once
func Get_invariantEither() gopurs_runtime.Value {
	once_invariantEither.Do(func() {
		invariantEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"imap": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(m_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": m_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(m_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(f_0, m_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
})})
	})
	return invariantEither
}

var fromRight_prime gopurs_runtime.Value
var once_fromRight_prime sync.Once
func Get_fromRight_prime() gopurs_runtime.Value {
	once_fromRight_prime.Do(func() {
		fromRight_prime = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t0 = gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
}
return __t0
})
})
	})
	return fromRight_prime
}

var fromRight gopurs_runtime.Value
var once_fromRight sync.Once
func Get_fromRight() gopurs_runtime.Value {
	once_fromRight.Do(func() {
		fromRight = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t0 = v_0
}
return __t0
})
})
	})
	return fromRight
}

var fromLeft_prime gopurs_runtime.Value
var once_fromLeft_prime sync.Once
func Get_fromLeft_prime() gopurs_runtime.Value {
	once_fromLeft_prime.Do(func() {
		fromLeft_prime = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t0 = gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
}
return __t0
})
})
	})
	return fromLeft_prime
}

var fromLeft gopurs_runtime.Value
var once_fromLeft sync.Once
func Get_fromLeft() gopurs_runtime.Value {
	once_fromLeft.Do(func() {
		fromLeft = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t0 = v_0
}
return __t0
})
})
	})
	return fromLeft
}

var extendEither gopurs_runtime.Value
var once_extendEither sync.Once
func Get_extendEither() gopurs_runtime.Value {
	once_extendEither.Do(func() {
		extendEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(v_0, v1_1)})
}
return __t0
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
})})
	})
	return extendEither
}

var eqEither gopurs_runtime.Value
var once_eqEither sync.Once
func Get_eqEither() gopurs_runtime.Value {
	once_eqEither.Do(func() {
		eqEither = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(y_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)
} else {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(y_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq1_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t0
})
})})
})
})
	})
	return eqEither
}

var ordEither gopurs_runtime.Value
var once_ordEither sync.Once
func Get_ordEither() gopurs_runtime.Value {
	once_ordEither.Do(func() {
		ordEither = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictOrd1_2.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
eqEither2_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t3 = gopurs_runtime.Bool(gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)
} else {
__t3 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t3
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(y_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
__t4 = __t5
} else {
if (gopurs_runtime.Bool(y_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(y_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0)).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd1_2.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t4
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqEither2_4_2
})})
})
})
	})
	return ordEither
}

var eq1Either gopurs_runtime.Value
var once_eq1Either sync.Once
func Get_eq1Either() gopurs_runtime.Value {
	once_eq1Either.Do(func() {
		eq1Either = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(y_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)
} else {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(y_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq1_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t0
})
})
})})
})
	})
	return eq1Either
}

var ord1Either gopurs_runtime.Value
var once_ord1Either sync.Once
func Get_ord1Either() gopurs_runtime.Value {
	once_ord1Either.Do(func() {
		ord1Either = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
ordEither1_1_0 := gopurs_runtime.Apply(Get_ordEither(), dictOrd_0)
__local_var_2_1 := gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
eq1Either1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t3 = gopurs_runtime.Bool(gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)
} else {
__t3 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq1_3.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t3
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(ordEither1_1_0, dictOrd1_4).PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Either1_3_2
})})
})
	})
	return ord1Either
}

var either gopurs_runtime.Value
var once_either sync.Once
func Get_either() gopurs_runtime.Value {
	once_either.Do(func() {
		either = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(v_0, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
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
	return either
}

var hush gopurs_runtime.Value
var once_hush sync.Once
func Get_hush() gopurs_runtime.Value {
	once_hush.Do(func() {
		hush = gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
	})
	return hush
}

var isLeft gopurs_runtime.Value
var once_isLeft sync.Once
func Get_isLeft() gopurs_runtime.Value {
	once_isLeft.Do(func() {
		isLeft = gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(true)
} else {
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(false)
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
	})
	return isLeft
}

var isRight gopurs_runtime.Value
var once_isRight sync.Once
func Get_isRight() gopurs_runtime.Value {
	once_isRight.Do(func() {
		isRight = gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(false)
} else {
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(true)
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
	})
	return isRight
}

var choose gopurs_runtime.Value
var once_choose sync.Once
func Get_choose() gopurs_runtime.Value {
	once_choose.Do(func() {
		choose = gopurs_runtime.Func(func(dictAlt_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_Left()), a_2)), gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_Right()), b_3))
})
})
})
	})
	return choose
}

var boundedEither gopurs_runtime.Value
var once_boundedEither sync.Once
func Get_boundedEither() gopurs_runtime.Value {
	once_boundedEither.Do(func() {
		boundedEither = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
bottom_1_0 := dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["bottom"]
ordEither1_2_1 := gopurs_runtime.Apply(Get_ordEither(), gopurs_runtime.Apply(dictBounded_0.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictBounded1_3 gopurs_runtime.Value) gopurs_runtime.Value {
ordEither2_4_2 := gopurs_runtime.Apply(ordEither1_2_1, gopurs_runtime.Apply(dictBounded1_3.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"top": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": dictBounded1_3.PtrVal.(map[string]gopurs_runtime.Value)["top"]}), "bottom": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": bottom_1_0}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return ordEither2_4_2
})})
})
})
	})
	return boundedEither
}

var blush gopurs_runtime.Value
var once_blush sync.Once
func Get_blush() gopurs_runtime.Value {
	once_blush.Do(func() {
		blush = gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
	})
	return blush
}

var applyEither gopurs_runtime.Value
var once_applyEither sync.Once
func Get_applyEither() gopurs_runtime.Value {
	once_applyEither.Do(func() {
		applyEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t0 = __t1
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
})})
	})
	return applyEither
}

var bindEither gopurs_runtime.Value
var once_bindEither sync.Once
func Get_bindEither() gopurs_runtime.Value {
	once_bindEither.Do(func() {
		bindEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__local_var_1_2 := v2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__t0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": __local_var_1_2})
})
} else {
if (gopurs_runtime.Bool(v2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__local_var_1_1 := v2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
__t0 = gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, __local_var_1_1)
})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEither()
})})
	})
	return bindEither
}

var semigroupEither gopurs_runtime.Value
var once_semigroupEither sync.Once
func Get_semigroupEither() gopurs_runtime.Value {
	once_semigroupEither.Do(func() {
		semigroupEither = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(x_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(y_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": y_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
} else {
if (gopurs_runtime.Bool(y_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], x_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
__t0 = __t1
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})})
})
	})
	return semigroupEither
}

var applicativeEither gopurs_runtime.Value
var once_applicativeEither sync.Once
func Get_applicativeEither() gopurs_runtime.Value {
	once_applicativeEither.Do(func() {
		applicativeEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": Get_Right(), "Apply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEither()
})})
	})
	return applicativeEither
}

var monadEither gopurs_runtime.Value
var once_monadEither sync.Once
func Get_monadEither() gopurs_runtime.Value {
	once_monadEither.Do(func() {
		monadEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeEither()
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindEither()
})})
	})
	return monadEither
}

var altEither gopurs_runtime.Value
var once_altEither sync.Once
func Get_altEither() gopurs_runtime.Value {
	once_altEither.Do(func() {
		altEither = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = v1_1
} else {
__t0 = v_0
}
return __t0
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
})})
	})
	return altEither
}


