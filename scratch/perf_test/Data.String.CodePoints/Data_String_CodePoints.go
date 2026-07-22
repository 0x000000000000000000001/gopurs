package Data_String_CodePoints

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Array "gopurs/output/Data.Array"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Data_String_Unsafe "gopurs/output/Data.String.Unsafe"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
)

var unsurrogate gopurs_runtime.Value
var once_unsurrogate sync.Once
func Get_unsurrogate() gopurs_runtime.Value {
	once_unsurrogate.Do(func() {
		unsurrogate = gopurs_runtime.Func(func(lead_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(trail_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intMul(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), lead_0), gopurs_runtime.Int(55296))), gopurs_runtime.Int(1024))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), trail_1), gopurs_runtime.Int(56320)))), gopurs_runtime.Int(65536))
})
})
	})
	return unsurrogate
}

var showCodePoint gopurs_runtime.Value
var once_showCodePoint sync.Once
func Get_showCodePoint() gopurs_runtime.Value {
	once_showCodePoint.Do(func() {
		showCodePoint = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(CodePoint 0x")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toUpper(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Int.Get_toStringAs(), gopurs_runtime.Int(16)), v_0))), gopurs_runtime.Str(")")))
})})
	})
	return showCodePoint
}

var isTrail gopurs_runtime.Value
var once_isTrail sync.Once
func Get_isTrail() gopurs_runtime.Value {
	once_isTrail.Do(func() {
		isTrail = gopurs_runtime.Func(func(cu_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Int(56320)), cu_0).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], cu_0), gopurs_runtime.Int(57343)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0))
})
	})
	return isTrail
}

var isLead gopurs_runtime.Value
var once_isLead sync.Once
func Get_isLead() gopurs_runtime.Value {
	once_isLead.Do(func() {
		isLead = gopurs_runtime.Func(func(cu_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Int(55296)), cu_0).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], cu_0), gopurs_runtime.Int(56319)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0))
})
	})
	return isLead
}

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
v_1_0 := gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), s_0)
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_1_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v_1_0.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"head": gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0)), s_0)), "tail": gopurs_runtime.Str("")})})
} else {
cu1_2_1 := gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(1)), s_0))
cu0_3_2 := gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0)), s_0))
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(Get_isLead(), cu0_3_2)), gopurs_runtime.Apply(Get_isTrail(), cu1_2_1))).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"head": gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsurrogate(), cu0_3_2), cu1_2_1), "tail": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(2)), s_0)})})
} else {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"head": cu0_3_2, "tail": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(1)), s_0)})})
}
__t4 = __t3
}
}
return __t4
})
	})
	return uncons
}

var unconsButWithTuple gopurs_runtime.Value
var once_unconsButWithTuple sync.Once
func Get_unconsButWithTuple() gopurs_runtime.Value {
	once_unconsButWithTuple.Do(func() {
		unconsButWithTuple = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_uncons(), s_0)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["head"], "value1": __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["tail"]})})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
	})
	return unconsButWithTuple
}

var toCodePointArrayFallback gopurs_runtime.Value
var once_toCodePointArrayFallback sync.Once
func Get_toCodePointArrayFallback() gopurs_runtime.Value {
	once_toCodePointArrayFallback.Do(func() {
		toCodePointArrayFallback = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Unfoldable.Get_unfoldableArray().PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], Get_unconsButWithTuple()), s_0)
})
	})
	return toCodePointArrayFallback
}

var unsafeCodePointAt0Fallback gopurs_runtime.Value
var once_unsafeCodePointAt0Fallback sync.Once
func Get_unsafeCodePointAt0Fallback() gopurs_runtime.Value {
	once_unsafeCodePointAt0Fallback.Do(func() {
		unsafeCodePointAt0Fallback = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
cu0_1_0 := gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0)), s_0))
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(Get_isLead(), cu0_1_0)), gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), s_0)), gopurs_runtime.Int(1)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT"))).IntVal != 0 {
cu1_2_2 := gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(1)), s_0))
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Apply(Get_isTrail(), cu1_2_2)).IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsurrogate(), cu0_1_0), cu1_2_2)
} else {
__t3 = cu0_1_0
}
__t1 = __t3
} else {
__t1 = cu0_1_0
}
return __t1
})
	})
	return unsafeCodePointAt0Fallback
}

var unsafeCodePointAt0 gopurs_runtime.Value
var once_unsafeCodePointAt0 sync.Once
func Get_unsafeCodePointAt0() gopurs_runtime.Value {
	once_unsafeCodePointAt0.Do(func() {
		unsafeCodePointAt0 = gopurs_runtime.Apply(Get__unsafeCodePointAt0(), Get_unsafeCodePointAt0Fallback())
	})
	return unsafeCodePointAt0
}

var toCodePointArray gopurs_runtime.Value
var once_toCodePointArray sync.Once
func Get_toCodePointArray() gopurs_runtime.Value {
	once_toCodePointArray.Do(func() {
		toCodePointArray = gopurs_runtime.Apply(gopurs_runtime.Apply(Get__toCodePointArray(), Get_toCodePointArrayFallback()), Get_unsafeCodePointAt0())
	})
	return toCodePointArray
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Array.Get_length(), gopurs_runtime.Apply(Get_toCodePointArray(), x_0))
})
	})
	return length
}

var lastIndexOf gopurs_runtime.Value
var once_lastIndexOf sync.Once
func Get_lastIndexOf() gopurs_runtime.Value {
	once_lastIndexOf.Do(func() {
		lastIndexOf = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_lastIndexOf(), p_0), s_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(pkg_Data_Array.Get_length(), gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_take(), __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), s_1)))})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
})
	})
	return lastIndexOf
}

var indexOf gopurs_runtime.Value
var once_indexOf sync.Once
func Get_indexOf() gopurs_runtime.Value {
	once_indexOf.Do(func() {
		indexOf = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_indexOf(), p_0), s_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(pkg_Data_Array.Get_length(), gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_take(), __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), s_1)))})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
})
	})
	return indexOf
}

var fromCharCode gopurs_runtime.Value
var once_fromCharCode sync.Once
func Get_fromCharCode() gopurs_runtime.Value {
	once_fromCharCode.Do(func() {
		fromCharCode = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Enum.Get_toEnumWithDefaults(), pkg_Data_Enum.Get_boundedEnumChar()), pkg_Data_Bounded.Get_bottomChar()), pkg_Data_Bounded.Get_topChar())
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_singleton(), gopurs_runtime.Apply(__local_var_0_0, x_1))
})
}()
	})
	return fromCharCode
}

var singletonFallback gopurs_runtime.Value
var once_singletonFallback sync.Once
func Get_singletonFallback() gopurs_runtime.Value {
	once_singletonFallback.Do(func() {
		singletonFallback = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_0), gopurs_runtime.Int(65535)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)).IntVal != 0 {
__t0 = gopurs_runtime.Apply(Get_fromCharCode(), v_0)
} else {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(Get_fromCharCode(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_intDiv(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0), gopurs_runtime.Int(65536))), gopurs_runtime.Int(1024))), gopurs_runtime.Int(55296)))), gopurs_runtime.Apply(Get_fromCharCode(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_intMod(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0), gopurs_runtime.Int(65536))), gopurs_runtime.Int(1024))), gopurs_runtime.Int(56320))))
}
return __t0
})
	})
	return singletonFallback
}

var fromCodePointArray gopurs_runtime.Value
var once_fromCodePointArray sync.Once
func Get_fromCodePointArray() gopurs_runtime.Value {
	once_fromCodePointArray.Do(func() {
		fromCodePointArray = gopurs_runtime.Apply(Get__fromCodePointArray(), Get_singletonFallback())
	})
	return fromCodePointArray
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Apply(Get__singleton(), Get_singletonFallback())
	})
	return singleton
}

var takeFallback gopurs_runtime.Value
var once_takeFallback sync.Once
func Get_takeFallback() gopurs_runtime.Value {
	once_takeFallback.Do(func() {
		takeFallback = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_0), gopurs_runtime.Int(1)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t2 = gopurs_runtime.Str("")
} else {
v2_2_0 := gopurs_runtime.Apply(Get_uncons(), v1_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v2_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(Get_singleton(), v2_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["head"])), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_takeFallback(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0), gopurs_runtime.Int(1))), v2_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["tail"]))
} else {
__t1 = v1_1
}
__t2 = __t1
}
return __t2
}
}()
})
})
	})
	return takeFallback
}

var take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		take = gopurs_runtime.Apply(Get__take(), Get_takeFallback())
	})
	return take
}

var lastIndexOf_prime gopurs_runtime.Value
var once_lastIndexOf_prime sync.Once
func Get_lastIndexOf_prime() gopurs_runtime.Value {
	once_lastIndexOf_prime.Do(func() {
		lastIndexOf_prime = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_lastIndexOf_prime(), p_0), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_take(), i_1), s_2))), s_2)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(pkg_Data_Array.Get_length(), gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_take(), __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), s_2)))})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
})
})
	})
	return lastIndexOf_prime
}

var splitAt gopurs_runtime.Value
var once_splitAt sync.Once
func Get_splitAt() gopurs_runtime.Value {
	once_splitAt.Do(func() {
		splitAt = gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
before_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_take(), i_0), s_1)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"before": before_2_0, "after": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), before_2_0)), s_1)})
})
})
	})
	return splitAt
}

var eqCodePoint gopurs_runtime.Value
var once_eqCodePoint sync.Once
func Get_eqCodePoint() gopurs_runtime.Value {
	once_eqCodePoint.Do(func() {
		eqCodePoint = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), x_0), y_1)
})
})})
	})
	return eqCodePoint
}

var ordCodePoint gopurs_runtime.Value
var once_ordCodePoint sync.Once
func Get_ordCodePoint() gopurs_runtime.Value {
	once_ordCodePoint.Do(func() {
		ordCodePoint = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0), y_1)
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqCodePoint()
})})
	})
	return ordCodePoint
}

var drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		drop = gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_take(), n_0), s_1))), s_1)
})
})
	})
	return drop
}

var indexOf_prime gopurs_runtime.Value
var once_indexOf_prime sync.Once
func Get_indexOf_prime() gopurs_runtime.Value {
	once_indexOf_prime.Do(func() {
		indexOf_prime = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_take(), i_1), s_2))), s_2)
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_indexOf(), p_0), s_prime_3_0)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_4_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), i_1), gopurs_runtime.Apply(pkg_Data_Array.Get_length(), gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_take(), __local_var_4_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), s_prime_3_0))))})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t2
})
})
})
	})
	return indexOf_prime
}

var countTail gopurs_runtime.Value
var once_countTail sync.Once
func Get_countTail() gopurs_runtime.Value {
	once_countTail.Do(func() {
		countTail = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(accum_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
v_3_0 := gopurs_runtime.Apply(Get_uncons(), s_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply(p_0, v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["head"]).IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_countTail(), p_0), v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["tail"]), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), accum_2), gopurs_runtime.Int(1)))
} else {
__t1 = accum_2
}
return __t1
}
}()
})
})
})
	})
	return countTail
}

var countFallback gopurs_runtime.Value
var once_countFallback sync.Once
func Get_countFallback() gopurs_runtime.Value {
	once_countFallback.Do(func() {
		countFallback = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_countTail(), p_0), s_1), gopurs_runtime.Int(0))
})
})
	})
	return countFallback
}

var countPrefix gopurs_runtime.Value
var once_countPrefix sync.Once
func Get_countPrefix() gopurs_runtime.Value {
	once_countPrefix.Do(func() {
		countPrefix = gopurs_runtime.Apply(gopurs_runtime.Apply(Get__countPrefix(), Get_countFallback()), Get_unsafeCodePointAt0())
	})
	return countPrefix
}

var dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		dropWhile = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_take(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_countPrefix(), p_0), s_1)), s_1))), s_1)
})
})
	})
	return dropWhile
}

var takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		takeWhile = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_take(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_countPrefix(), p_0), s_1)), s_1)
})
})
	})
	return takeWhile
}

var codePointFromChar gopurs_runtime.Value
var once_codePointFromChar sync.Once
func Get_codePointFromChar() gopurs_runtime.Value {
	once_codePointFromChar.Do(func() {
		codePointFromChar = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), x_0)
})
	})
	return codePointFromChar
}

var codePointAtFallback gopurs_runtime.Value
var once_codePointAtFallback sync.Once
func Get_codePointAtFallback() gopurs_runtime.Value {
	once_codePointAtFallback.Do(func() {
		codePointAtFallback = gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
for {
v_2_0 := gopurs_runtime.Apply(Get_uncons(), s_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), n_0), gopurs_runtime.Int(0))).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["head"]})
} else {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_codePointAtFallback(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), n_0), gopurs_runtime.Int(1))), v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["tail"])
}
__t1 = __t2
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
}
}()
})
})
	})
	return codePointAtFallback
}

var codePointAt gopurs_runtime.Value
var once_codePointAt sync.Once
func Get_codePointAt() gopurs_runtime.Value {
	once_codePointAt.Do(func() {
		codePointAt = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], v_0), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.StrVal == gopurs_runtime.Str("").StrVal)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(Get_unsafeCodePointAt0(), v1_1)})
}
__t0 = __t1
} else {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get__codePointAt(), Get_codePointAtFallback()), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})), Get_unsafeCodePointAt0()), v_0), v1_1)
}
}
return __t0
})
})
	})
	return codePointAt
}

var boundedCodePoint gopurs_runtime.Value
var once_boundedCodePoint sync.Once
func Get_boundedCodePoint() gopurs_runtime.Value {
	once_boundedCodePoint.Do(func() {
		boundedCodePoint = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bottom": gopurs_runtime.Int(0), "top": gopurs_runtime.Int(1114111), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordCodePoint()
})})
	})
	return boundedCodePoint
}

var boundedEnumCodePoint gopurs_runtime.Value
var once_boundedEnumCodePoint sync.Once
func Get_boundedEnumCodePoint() gopurs_runtime.Value {
	once_boundedEnumCodePoint.Do(func() {
		boundedEnumCodePoint = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"cardinality": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1114111)), gopurs_runtime.Int(1)), "fromEnum": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), "toEnum": gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], n_0), gopurs_runtime.Int(1114111)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0))).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": n_0})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t0
}), "Bounded0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedCodePoint()
}), "Enum1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumCodePoint()
})})
	})
	return boundedEnumCodePoint
}

var enumCodePoint gopurs_runtime.Value
var once_enumCodePoint sync.Once
func Get_enumCodePoint() gopurs_runtime.Value {
	once_enumCodePoint.Do(func() {
		enumCodePoint = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"succ": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Enum.Get_defaultSucc(), Get_boundedEnumCodePoint().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"]), Get_boundedEnumCodePoint().PtrVal.(map[string]gopurs_runtime.Value)["fromEnum"]), "pred": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Enum.Get_defaultPred(), Get_boundedEnumCodePoint().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"]), Get_boundedEnumCodePoint().PtrVal.(map[string]gopurs_runtime.Value)["fromEnum"]), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordCodePoint()
})})
	})
	return enumCodePoint
}

func Get__codePointAt() gopurs_runtime.Value {
	return X_CodePointAt
}

func Get__countPrefix() gopurs_runtime.Value {
	return X_CountPrefix
}

func Get__fromCodePointArray() gopurs_runtime.Value {
	return X_FromCodePointArray
}

func Get__singleton() gopurs_runtime.Value {
	return X_Singleton
}

func Get__take() gopurs_runtime.Value {
	return X_Take
}

func Get__toCodePointArray() gopurs_runtime.Value {
	return X_ToCodePointArray
}

func Get__unsafeCodePointAt0() gopurs_runtime.Value {
	return X_UnsafeCodePointAt0
}
