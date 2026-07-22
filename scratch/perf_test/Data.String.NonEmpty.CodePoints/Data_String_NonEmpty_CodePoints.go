package Data_String_NonEmpty_CodePoints

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_String_CodePoints "gopurs/output/Data.String.CodePoints"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
)

var fromJust gopurs_runtime.Value
var once_fromJust sync.Once
func Get_fromJust() gopurs_runtime.Value {
	once_fromJust.Do(func() {
		fromJust = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
return __t0
})
	})
	return fromJust
}

var snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		snoc = gopurs_runtime.Func(func(c_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), s_1), gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_singleton(), c_0))
})
})
	})
	return snoc
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_singleton(), x_0)
})
	})
	return singleton
}

var takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		takeWhile = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_take(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_countPrefix(), f_0), x_1)), x_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.StrVal == gopurs_runtime.Str("").StrVal)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_2_0})
}
return __t1
})
})
	})
	return takeWhile
}

var lastIndexOf_prime gopurs_runtime.Value
var once_lastIndexOf_prime sync.Once
func Get_lastIndexOf_prime() gopurs_runtime.Value {
	once_lastIndexOf_prime.Do(func() {
		lastIndexOf_prime = gopurs_runtime.Func(func(pat_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_lastIndexOf_prime(), pat_0), x_1), v_2)
})
})
})
	})
	return lastIndexOf_prime
}

var lastIndexOf gopurs_runtime.Value
var once_lastIndexOf sync.Once
func Get_lastIndexOf() gopurs_runtime.Value {
	once_lastIndexOf.Do(func() {
		lastIndexOf = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_lastIndexOf(), x_0), v_1)
})
})
	})
	return lastIndexOf
}

var indexOf_prime gopurs_runtime.Value
var once_indexOf_prime sync.Once
func Get_indexOf_prime() gopurs_runtime.Value {
	once_indexOf_prime.Do(func() {
		indexOf_prime = gopurs_runtime.Func(func(pat_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_indexOf_prime(), pat_0), x_1), v_2)
})
})
})
	})
	return indexOf_prime
}

var indexOf gopurs_runtime.Value
var once_indexOf sync.Once
func Get_indexOf() gopurs_runtime.Value {
	once_indexOf.Do(func() {
		indexOf = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_indexOf(), x_0), v_1)
})
})
	})
	return indexOf
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Array.Get_length(), gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_toCodePointArray(), x_0))
})
	})
	return length
}

var splitAt gopurs_runtime.Value
var once_splitAt sync.Once
func Get_splitAt() gopurs_runtime.Value {
	once_splitAt.Do(func() {
		splitAt = gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(nes_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_splitAt(), i_0), nes_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["before"].StrVal == gopurs_runtime.Str("").StrVal)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["before"]})
}
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["after"].StrVal == gopurs_runtime.Str("").StrVal)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["after"]})
}
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"before": __t1, "after": __t2})
})
})
	})
	return splitAt
}

var take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		take = gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(nes_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], i_0), gopurs_runtime.Int(1)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_take(), i_0), nes_1)})
}
return __t0
})
})
	})
	return take
}

var toCodePointArray gopurs_runtime.Value
var once_toCodePointArray sync.Once
func Get_toCodePointArray() gopurs_runtime.Value {
	once_toCodePointArray.Do(func() {
		toCodePointArray = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_toCodePointArray(), x_0)
})
	})
	return toCodePointArray
}

var toNonEmptyCodePointArray gopurs_runtime.Value
var once_toNonEmptyCodePointArray sync.Once
func Get_toNonEmptyCodePointArray() gopurs_runtime.Value {
	once_toNonEmptyCodePointArray.Do(func() {
		toNonEmptyCodePointArray = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust()
}))
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_toCodePointArray(), x_1)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(pkg_Data_Array.Get_length(), __local_var_2_1)), gopurs_runtime.Int(0)).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_2_1})
} else {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return gopurs_runtime.Apply(__local_var_0_0, __t2)
})
}()
	})
	return toNonEmptyCodePointArray
}

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(nes_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_take(), gopurs_runtime.Int(1)), nes_0))), nes_0)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_1_0.StrVal == gopurs_runtime.Str("").StrVal)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_1_0})
}
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"head": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust()
})), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_codePointAt(), gopurs_runtime.Int(0)), nes_0)), "tail": __t1})
})
	})
	return uncons
}

var fromFoldable1 gopurs_runtime.Value
var once_fromFoldable1 sync.Once
func Get_fromFoldable1() gopurs_runtime.Value {
	once_fromFoldable1.Do(func() {
		fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"], pkg_Data_Semigroup.Get_semigroupString()), Get_singleton())
})
	})
	return fromFoldable1
}

var fromCodePointArray gopurs_runtime.Value
var once_fromCodePointArray sync.Once
func Get_fromCodePointArray() gopurs_runtime.Value {
	once_fromCodePointArray.Do(func() {
		fromCodePointArray = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Int(len(v_0.PtrVal.([]gopurs_runtime.Value))).IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_fromCodePointArray(), v_0)})
}
return __t0
})
	})
	return fromCodePointArray
}

var fromNonEmptyCodePointArray gopurs_runtime.Value
var once_fromNonEmptyCodePointArray sync.Once
func Get_fromNonEmptyCodePointArray() gopurs_runtime.Value {
	once_fromNonEmptyCodePointArray.Do(func() {
		fromNonEmptyCodePointArray = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust()
}))
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Int(len(x_1.PtrVal.([]gopurs_runtime.Value))).IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_fromCodePointArray(), x_1)})
}
return gopurs_runtime.Apply(__local_var_0_0, __t1)
})
}()
	})
	return fromNonEmptyCodePointArray
}

var dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		dropWhile = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_dropWhile(), f_0), x_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_2_0.StrVal == gopurs_runtime.Str("").StrVal)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_2_0})
}
return __t1
})
})
	})
	return dropWhile
}

var drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		drop = gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(nes_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], i_0), gopurs_runtime.Apply(pkg_Data_Array.Get_length(), gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_toCodePointArray(), nes_1))).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_take(), i_0), nes_1))), nes_1)})
}
return __t0
})
})
	})
	return drop
}

var countPrefix gopurs_runtime.Value
var once_countPrefix sync.Once
func Get_countPrefix() gopurs_runtime.Value {
	once_countPrefix.Do(func() {
		countPrefix = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_countPrefix(), x_0)
})
	})
	return countPrefix
}

var cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cons = gopurs_runtime.Func(func(c_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_singleton(), c_0)), s_1)
})
})
	})
	return cons
}

var codePointAt gopurs_runtime.Value
var once_codePointAt sync.Once
func Get_codePointAt() gopurs_runtime.Value {
	once_codePointAt.Do(func() {
		codePointAt = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_codePointAt(), x_0), v_1)
})
})
	})
	return codePointAt
}


