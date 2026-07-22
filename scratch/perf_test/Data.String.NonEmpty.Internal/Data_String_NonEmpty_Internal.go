package Data_String_NonEmpty_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var NonEmptyString gopurs_runtime.Value
var once_NonEmptyString sync.Once
func Get_NonEmptyString() gopurs_runtime.Value {
	once_NonEmptyString.Do(func() {
		NonEmptyString = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return NonEmptyString
}

var NonEmptyReplacement gopurs_runtime.Value
var once_NonEmptyReplacement sync.Once
func Get_NonEmptyReplacement() gopurs_runtime.Value {
	once_NonEmptyReplacement.Do(func() {
		NonEmptyReplacement = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return NonEmptyReplacement
}

var toUpper gopurs_runtime.Value
var once_toUpper sync.Once
func Get_toUpper() gopurs_runtime.Value {
	once_toUpper.Do(func() {
		toUpper = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_Common.Get_toUpper(), v_0)
})
	})
	return toUpper
}

var toString gopurs_runtime.Value
var once_toString sync.Once
func Get_toString() gopurs_runtime.Value {
	once_toString.Do(func() {
		toString = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return toString
}

var toLower gopurs_runtime.Value
var once_toLower sync.Once
func Get_toLower() gopurs_runtime.Value {
	once_toLower.Do(func() {
		toLower = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v_0)
})
	})
	return toLower
}

var showNonEmptyString gopurs_runtime.Value
var once_showNonEmptyString sync.Once
func Get_showNonEmptyString() gopurs_runtime.Value {
	once_showNonEmptyString.Do(func() {
		showNonEmptyString = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(NonEmptyString.unsafeFromString ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showStringImpl(), v_0)), gopurs_runtime.Str(")")))
})})
	})
	return showNonEmptyString
}

var showNonEmptyReplacement gopurs_runtime.Value
var once_showNonEmptyReplacement sync.Once
func Get_showNonEmptyReplacement() gopurs_runtime.Value {
	once_showNonEmptyReplacement.Do(func() {
		showNonEmptyReplacement = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(NonEmptyReplacement ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(NonEmptyString.unsafeFromString ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showStringImpl(), v_0)), gopurs_runtime.Str(")")))), gopurs_runtime.Str(")")))
})})
	})
	return showNonEmptyReplacement
}

var semigroupNonEmptyString gopurs_runtime.Value
var once_semigroupNonEmptyString sync.Once
func Get_semigroupNonEmptyString() gopurs_runtime.Value {
	once_semigroupNonEmptyString.Do(func() {
		semigroupNonEmptyString = pkg_Data_Semigroup.Get_semigroupString()
	})
	return semigroupNonEmptyString
}

var semigroupNonEmptyReplacement gopurs_runtime.Value
var once_semigroupNonEmptyReplacement sync.Once
func Get_semigroupNonEmptyReplacement() gopurs_runtime.Value {
	once_semigroupNonEmptyReplacement.Do(func() {
		semigroupNonEmptyReplacement = pkg_Data_Semigroup.Get_semigroupString()
	})
	return semigroupNonEmptyReplacement
}

var replaceAll gopurs_runtime.Value
var once_replaceAll sync.Once
func Get_replaceAll() gopurs_runtime.Value {
	once_replaceAll.Do(func() {
		replaceAll = gopurs_runtime.Func(func(pat_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_Common.Get_replaceAll(), pat_0), v_1), v1_2)
})
})
})
	})
	return replaceAll
}

var replace gopurs_runtime.Value
var once_replace sync.Once
func Get_replace() gopurs_runtime.Value {
	once_replace.Do(func() {
		replace = gopurs_runtime.Func(func(pat_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_Common.Get_replace(), pat_0), v_1), v1_2)
})
})
})
	})
	return replace
}

var prependString gopurs_runtime.Value
var once_prependString sync.Once
func Get_prependString() gopurs_runtime.Value {
	once_prependString.Do(func() {
		prependString = gopurs_runtime.Func(func(s1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), s1_0), v_1)
})
})
	})
	return prependString
}

var ordNonEmptyString gopurs_runtime.Value
var once_ordNonEmptyString sync.Once
func Get_ordNonEmptyString() gopurs_runtime.Value {
	once_ordNonEmptyString.Do(func() {
		ordNonEmptyString = pkg_Data_Ord.Get_ordString()
	})
	return ordNonEmptyString
}

var ordNonEmptyReplacement gopurs_runtime.Value
var once_ordNonEmptyReplacement sync.Once
func Get_ordNonEmptyReplacement() gopurs_runtime.Value {
	once_ordNonEmptyReplacement.Do(func() {
		ordNonEmptyReplacement = pkg_Data_Ord.Get_ordString()
	})
	return ordNonEmptyReplacement
}

var nonEmptyNonEmpty gopurs_runtime.Value
var once_nonEmptyNonEmpty sync.Once
func Get_nonEmptyNonEmpty() gopurs_runtime.Value {
	once_nonEmptyNonEmpty.Do(func() {
		nonEmptyNonEmpty = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"nes": gopurs_runtime.Func(func(p_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], p_1)
})})
})
	})
	return nonEmptyNonEmpty
}

var nes gopurs_runtime.Value
var once_nes sync.Once
func Get_nes() gopurs_runtime.Value {
	once_nes.Do(func() {
		nes = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["nes"]
})
	})
	return nes
}

var makeNonEmptyBad gopurs_runtime.Value
var once_makeNonEmptyBad sync.Once
func Get_makeNonEmptyBad() gopurs_runtime.Value {
	once_makeNonEmptyBad.Do(func() {
		makeNonEmptyBad = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"nes": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
})})
})
	})
	return makeNonEmptyBad
}

var localeCompare gopurs_runtime.Value
var once_localeCompare sync.Once
func Get_localeCompare() gopurs_runtime.Value {
	once_localeCompare.Do(func() {
		localeCompare = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_Common.Get_localeCompare(), v_0), v1_1)
})
})
	})
	return localeCompare
}

var liftS gopurs_runtime.Value
var once_liftS sync.Once
func Get_liftS() gopurs_runtime.Value {
	once_liftS.Do(func() {
		liftS = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v_1)
})
})
	})
	return liftS
}

var startsWith gopurs_runtime.Value
var once_startsWith sync.Once
func Get_startsWith() gopurs_runtime.Value {
	once_startsWith.Do(func() {
		startsWith = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_stripPrefix(), x_0), v_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
} else {
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(true)
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
	})
	return startsWith
}

var joinWith1 gopurs_runtime.Value
var once_joinWith1 sync.Once
func Get_joinWith1() gopurs_runtime.Value {
	once_joinWith1.Do(func() {
		joinWith1 = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(sep_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (v_4.PtrVal.(map[string]gopurs_runtime.Value)["init"]).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(false), "acc": v1_5})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(false), "acc": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), v_4.PtrVal.(map[string]gopurs_runtime.Value)["acc"]), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), sep_2), v1_5))})
}
return __t1
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(true), "acc": gopurs_runtime.Str("")})), xs_3).PtrVal.(map[string]gopurs_runtime.Value)["acc"]
})
})
})
	})
	return joinWith1
}

var joinWith gopurs_runtime.Value
var once_joinWith sync.Once
func Get_joinWith() gopurs_runtime.Value {
	once_joinWith.Do(func() {
		joinWith = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(splice_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (v_3.PtrVal.(map[string]gopurs_runtime.Value)["init"]).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(false), "acc": v1_4})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(false), "acc": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), v_3.PtrVal.(map[string]gopurs_runtime.Value)["acc"]), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), splice_1), v1_4))})
}
return __t0
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"init": gopurs_runtime.Bool(true), "acc": gopurs_runtime.Str("")})), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), x_2)).PtrVal.(map[string]gopurs_runtime.Value)["acc"]
})
})
})
	})
	return joinWith
}

var join1With gopurs_runtime.Value
var once_join1With sync.Once
func Get_join1With() gopurs_runtime.Value {
	once_join1With.Do(func() {
		join1With = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_joinWith(), gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{}))
})
	})
	return join1With
}

var fromString gopurs_runtime.Value
var once_fromString sync.Once
func Get_fromString() gopurs_runtime.Value {
	once_fromString.Do(func() {
		fromString = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.StrVal == gopurs_runtime.Str("").StrVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_0})
}
return __t0
})
	})
	return fromString
}

var stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		stripPrefix = gopurs_runtime.Func(func(pat_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_stripPrefix(), pat_0), a_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].StrVal == gopurs_runtime.Str("").StrVal)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
}
__t1 = __t2
} else {
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
	})
	return stripPrefix
}

var stripSuffix gopurs_runtime.Value
var once_stripSuffix sync.Once
func Get_stripSuffix() gopurs_runtime.Value {
	once_stripSuffix.Do(func() {
		stripSuffix = gopurs_runtime.Func(func(pat_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_stripSuffix(), pat_0), a_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].StrVal == gopurs_runtime.Str("").StrVal)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
}
__t1 = __t2
} else {
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
	})
	return stripSuffix
}

var trim gopurs_runtime.Value
var once_trim sync.Once
func Get_trim() gopurs_runtime.Value {
	once_trim.Do(func() {
		trim = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_Common.Get_trim(), v_0)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_1_0.StrVal == gopurs_runtime.Str("").StrVal)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_1_0})
}
return __t1
})
	})
	return trim
}

var unsafeFromString gopurs_runtime.Value
var once_unsafeFromString sync.Once
func Get_unsafeFromString() gopurs_runtime.Value {
	once_unsafeFromString.Do(func() {
		unsafeFromString = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(x_1.StrVal == gopurs_runtime.Str("").StrVal)).IntVal != 0 {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
} else {
__t0 = x_1
}
return __t0
})
})
	})
	return unsafeFromString
}

var eqNonEmptyString gopurs_runtime.Value
var once_eqNonEmptyString sync.Once
func Get_eqNonEmptyString() gopurs_runtime.Value {
	once_eqNonEmptyString.Do(func() {
		eqNonEmptyString = pkg_Data_Eq.Get_eqString()
	})
	return eqNonEmptyString
}

var eqNonEmptyReplacement gopurs_runtime.Value
var once_eqNonEmptyReplacement sync.Once
func Get_eqNonEmptyReplacement() gopurs_runtime.Value {
	once_eqNonEmptyReplacement.Do(func() {
		eqNonEmptyReplacement = pkg_Data_Eq.Get_eqString()
	})
	return eqNonEmptyReplacement
}

var endsWith gopurs_runtime.Value
var once_endsWith sync.Once
func Get_endsWith() gopurs_runtime.Value {
	once_endsWith.Do(func() {
		endsWith = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_stripSuffix(), x_0), v_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
} else {
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(true)
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t1
})
})
	})
	return endsWith
}

var contains gopurs_runtime.Value
var once_contains sync.Once
func Get_contains() gopurs_runtime.Value {
	once_contains.Do(func() {
		contains = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_contains(), x_0)
})
	})
	return contains
}

var appendString gopurs_runtime.Value
var once_appendString sync.Once
func Get_appendString() gopurs_runtime.Value {
	once_appendString.Do(func() {
		appendString = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), v_0), s2_1)
})
})
	})
	return appendString
}


