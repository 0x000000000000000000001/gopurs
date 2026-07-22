package Data_Functor_Coproduct

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var Coproduct gopurs_runtime.Value
var once_Coproduct sync.Once
func Get_Coproduct() gopurs_runtime.Value {
	once_Coproduct.Do(func() {
		Coproduct = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Coproduct
}

var showCoproduct gopurs_runtime.Value
var once_showCoproduct sync.Once
func Get_showCoproduct() gopurs_runtime.Value {
	once_showCoproduct.Do(func() {
		showCoproduct = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(left ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Str(")")))
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(right ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow1_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Str(")")))
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})})
})
})
	})
	return showCoproduct
}

var right gopurs_runtime.Value
var once_right sync.Once
func Get_right() gopurs_runtime.Value {
	once_right.Do(func() {
		right = gopurs_runtime.Func(func(ga_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": ga_0})
})
	})
	return right
}

var newtypeCoproduct gopurs_runtime.Value
var once_newtypeCoproduct sync.Once
func Get_newtypeCoproduct() gopurs_runtime.Value {
	once_newtypeCoproduct.Do(func() {
		newtypeCoproduct = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeCoproduct
}

var left gopurs_runtime.Value
var once_left sync.Once
func Get_left() gopurs_runtime.Value {
	once_left.Do(func() {
		left = gopurs_runtime.Func(func(fa_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": fa_0})
})
	})
	return left
}

var functorCoproduct gopurs_runtime.Value
var once_functorCoproduct sync.Once
func Get_functorCoproduct() gopurs_runtime.Value {
	once_functorCoproduct.Do(func() {
		functorCoproduct = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2)
__local_var_5_1 := gopurs_runtime.Apply(dictFunctor1_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": gopurs_runtime.Apply(__local_var_4_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(__local_var_5_1, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})})
})
})
	})
	return functorCoproduct
}

var eq1Coproduct gopurs_runtime.Value
var once_eq1Coproduct sync.Once
func Get_eq1Coproduct() gopurs_runtime.Value {
	once_eq1Coproduct.Do(func() {
		eq1Coproduct = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq11_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq12_3_0 := gopurs_runtime.Apply(dictEq1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_2)
eq13_4_1 := gopurs_runtime.Apply(dictEq11_1.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_2)
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(eq12_3_0, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)
} else {
__t2 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(eq13_4_1, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t2
})
})
})})
})
})
	})
	return eq1Coproduct
}

var eqCoproduct gopurs_runtime.Value
var once_eqCoproduct sync.Once
func Get_eqCoproduct() gopurs_runtime.Value {
	once_eqCoproduct.Do(func() {
		eqCoproduct = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq11_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
eq12_3_0 := gopurs_runtime.Apply(dictEq1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_2)
eq13_4_1 := gopurs_runtime.Apply(dictEq11_1.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(eq12_3_0, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)
} else {
__t2 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v1_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(eq13_4_1, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t2
})
})})
})
})
})
	})
	return eqCoproduct
}

var ord1Coproduct gopurs_runtime.Value
var once_ord1Coproduct sync.Once
func Get_ord1Coproduct() gopurs_runtime.Value {
	once_ord1Coproduct.Do(func() {
		ord1Coproduct = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictOrd11_2.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{})
eq1Coproduct2_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_4 gopurs_runtime.Value) gopurs_runtime.Value {
eq12_5_3 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_4)
eq13_6_4 := gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_4)
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t5 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(eq12_5_3, v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)
} else {
__t5 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v1_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(eq13_6_4, v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t5
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
compare12_6_6 := gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["compare1"], dictOrd_5)
compare13_7_7 := gopurs_runtime.Apply(dictOrd11_2.PtrVal.(map[string]gopurs_runtime.Value)["compare1"], dictOrd_5)
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
var __t9 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t9 = gopurs_runtime.Apply(gopurs_runtime.Apply(compare12_6_6, v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t9 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
__t8 = __t9
} else {
if (gopurs_runtime.Bool(v1_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(v1_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0)).IntVal != 0 {
__t8 = gopurs_runtime.Apply(gopurs_runtime.Apply(compare13_7_7, v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t8
})
})
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Coproduct2_4_2
})})
})
})
	})
	return ord1Coproduct
}

var ordCoproduct gopurs_runtime.Value
var once_ordCoproduct sync.Once
func Get_ordCoproduct() gopurs_runtime.Value {
	once_ordCoproduct.Do(func() {
		ordCoproduct = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
ord1Coproduct1_1_0 := gopurs_runtime.Apply(Get_ord1Coproduct(), dictOrd1_0)
__local_var_2_1 := gopurs_runtime.Apply(dictOrd1_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictOrd11_3.PtrVal.(map[string]gopurs_runtime.Value)["Eq10"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(dictOrd_5.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
eq12_7_5 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], __local_var_6_3)
eq13_8_6 := gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], __local_var_6_3)
eqCoproduct3_7_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t7 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(eq12_7_5, v_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0)
} else {
__t7 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(v1_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(eq13_8_6, v_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
return __t7
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Apply(gopurs_runtime.Apply(ord1Coproduct1_1_0, dictOrd11_3).PtrVal.(map[string]gopurs_runtime.Value)["compare1"], dictOrd_5), "Eq0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return eqCoproduct3_7_4
})})
})
})
})
	})
	return ordCoproduct
}

var coproduct gopurs_runtime.Value
var once_coproduct sync.Once
func Get_coproduct() gopurs_runtime.Value {
	once_coproduct.Do(func() {
		coproduct = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return coproduct
}

var extendCoproduct gopurs_runtime.Value
var once_extendCoproduct sync.Once
func Get_extendCoproduct() gopurs_runtime.Value {
	once_extendCoproduct.Do(func() {
		extendCoproduct = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictExtend_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictExtend1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictExtend1_2.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorCoproduct2_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4)
__local_var_7_4 := gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_4)
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": gopurs_runtime.Apply(__local_var_6_3, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
if (gopurs_runtime.Bool(v_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(__local_var_7_4, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t5
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_6 := gopurs_runtime.Apply(dictExtend_0.PtrVal.(map[string]gopurs_runtime.Value)["extend"], gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_6}))
}))
__local_var_7_7 := gopurs_runtime.Apply(dictExtend1_2.PtrVal.(map[string]gopurs_runtime.Value)["extend"], gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": x_7}))
}))
return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": gopurs_runtime.Apply(__local_var_6_6, v2_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
if (gopurs_runtime.Bool(v2_8.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(__local_var_7_7, v2_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t8
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCoproduct2_4_2
})})
})
})
	})
	return extendCoproduct
}

var comonadCoproduct gopurs_runtime.Value
var once_comonadCoproduct sync.Once
func Get_comonadCoproduct() gopurs_runtime.Value {
	once_comonadCoproduct.Do(func() {
		comonadCoproduct = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
extendCoproduct1_1_0 := gopurs_runtime.Apply(Get_extendCoproduct(), gopurs_runtime.Apply(dictComonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictComonad1_2 gopurs_runtime.Value) gopurs_runtime.Value {
extendCoproduct2_3_1 := gopurs_runtime.Apply(extendCoproduct1_1_0, gopurs_runtime.Apply(dictComonad1_2.PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extract": gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(dictComonad_0.PtrVal.(map[string]gopurs_runtime.Value)["extract"], v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
if (gopurs_runtime.Bool(v2_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(dictComonad1_2.PtrVal.(map[string]gopurs_runtime.Value)["extract"], v2_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
}), "Extend0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return extendCoproduct2_3_1
})})
})
})
	})
	return comonadCoproduct
}

var bihoistCoproduct gopurs_runtime.Value
var once_bihoistCoproduct sync.Once
func Get_bihoistCoproduct() gopurs_runtime.Value {
	once_bihoistCoproduct.Do(func() {
		bihoistCoproduct = gopurs_runtime.Func(func(natF_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(natG_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": gopurs_runtime.Apply(natF_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Right"), "value0": gopurs_runtime.Apply(natG_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t0
})
})
})
	})
	return bihoistCoproduct
}


