package Data_String_CodeUnits

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_String_Unsafe "gopurs/output/Data.String.Unsafe"
)

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.StrVal == gopurs_runtime.Str("").StrVal)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
} else {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"head": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0)), v_0), "tail": gopurs_runtime.Apply(gopurs_runtime.Apply(Get_drop(), gopurs_runtime.Int(1)), v_0)})})
}
return __t0
})
	})
	return uncons
}

var toChar gopurs_runtime.Value
var once_toChar sync.Once
func Get_toChar() gopurs_runtime.Value {
	once_toChar.Do(func() {
		toChar = gopurs_runtime.Apply(gopurs_runtime.Apply(Get__toChar(), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return toChar
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

var takeRight gopurs_runtime.Value
var once_takeRight sync.Once
func Get_takeRight() gopurs_runtime.Value {
	once_takeRight.Do(func() {
		takeRight = gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_drop(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(Get_length(), s_1)), i_0)), s_1)
})
})
	})
	return takeRight
}

var stripSuffix gopurs_runtime.Value
var once_stripSuffix sync.Once
func Get_stripSuffix() gopurs_runtime.Value {
	once_stripSuffix.Do(func() {
		stripSuffix = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(str_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_splitAt(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(Get_length(), str_1)), gopurs_runtime.Apply(Get_length(), v_0))), str_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqStringImpl(), v1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["after"]), v_0)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["before"]})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
})
	})
	return stripSuffix
}

var stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		stripPrefix = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(str_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_splitAt(), gopurs_runtime.Apply(Get_length(), v_0)), str_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqStringImpl(), v1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["before"]), v_0)).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": v1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["after"]})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
return __t1
})
})
	})
	return stripPrefix
}

var startsWith gopurs_runtime.Value
var once_startsWith sync.Once
func Get_startsWith() gopurs_runtime.Value {
	once_startsWith.Do(func() {
		startsWith = gopurs_runtime.Func(func(pat_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_stripPrefix(), pat_0), x_1)
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

var lastIndexOf_prime gopurs_runtime.Value
var once_lastIndexOf_prime sync.Once
func Get_lastIndexOf_prime() gopurs_runtime.Value {
	once_lastIndexOf_prime.Do(func() {
		lastIndexOf_prime = gopurs_runtime.Apply(gopurs_runtime.Apply(Get__lastIndexOfStartingAt(), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return lastIndexOf_prime
}

var lastIndexOf gopurs_runtime.Value
var once_lastIndexOf sync.Once
func Get_lastIndexOf() gopurs_runtime.Value {
	once_lastIndexOf.Do(func() {
		lastIndexOf = gopurs_runtime.Apply(gopurs_runtime.Apply(Get__lastIndexOf(), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return lastIndexOf
}

var indexOf_prime gopurs_runtime.Value
var once_indexOf_prime sync.Once
func Get_indexOf_prime() gopurs_runtime.Value {
	once_indexOf_prime.Do(func() {
		indexOf_prime = gopurs_runtime.Apply(gopurs_runtime.Apply(Get__indexOfStartingAt(), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return indexOf_prime
}

var indexOf gopurs_runtime.Value
var once_indexOf sync.Once
func Get_indexOf() gopurs_runtime.Value {
	once_indexOf.Do(func() {
		indexOf = gopurs_runtime.Apply(gopurs_runtime.Apply(Get__indexOf(), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return indexOf
}

var endsWith gopurs_runtime.Value
var once_endsWith sync.Once
func Get_endsWith() gopurs_runtime.Value {
	once_endsWith.Do(func() {
		endsWith = gopurs_runtime.Func(func(pat_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_stripSuffix(), pat_0), x_1)
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

var dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		dropWhile = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_drop(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_countPrefix(), p_0), s_1)), s_1)
})
})
	})
	return dropWhile
}

var dropRight gopurs_runtime.Value
var once_dropRight sync.Once
func Get_dropRight() gopurs_runtime.Value {
	once_dropRight.Do(func() {
		dropRight = gopurs_runtime.Func(func(i_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_take(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(Get_length(), s_1)), i_0)), s_1)
})
})
	})
	return dropRight
}

var contains gopurs_runtime.Value
var once_contains sync.Once
func Get_contains() gopurs_runtime.Value {
	once_contains.Do(func() {
		contains = gopurs_runtime.Func(func(pat_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_indexOf(), pat_0)
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(false)
} else {
if (gopurs_runtime.Bool(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Bool(true)
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
return __t2
})
})
	})
	return contains
}

var charAt gopurs_runtime.Value
var once_charAt sync.Once
func Get_charAt() gopurs_runtime.Value {
	once_charAt.Do(func() {
		charAt = gopurs_runtime.Apply(gopurs_runtime.Apply(Get__charAt(), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return charAt
}

func Get__charAt() gopurs_runtime.Value {
	return X_CharAt
}

func Get__indexOf() gopurs_runtime.Value {
	return X_IndexOf
}

func Get__indexOfStartingAt() gopurs_runtime.Value {
	return X_IndexOfStartingAt
}

func Get__lastIndexOf() gopurs_runtime.Value {
	return X_LastIndexOf
}

func Get__lastIndexOfStartingAt() gopurs_runtime.Value {
	return X_LastIndexOfStartingAt
}

func Get__toChar() gopurs_runtime.Value {
	return X_ToChar
}

func Get_countPrefix() gopurs_runtime.Value {
	return CountPrefix
}

func Get_drop() gopurs_runtime.Value {
	return Drop
}

func Get_fromCharArray() gopurs_runtime.Value {
	return FromCharArray
}

func Get_length() gopurs_runtime.Value {
	return Length
}

func Get_singleton() gopurs_runtime.Value {
	return Singleton
}

func Get_slice() gopurs_runtime.Value {
	return Slice
}

func Get_splitAt() gopurs_runtime.Value {
	return SplitAt
}

func Get_take() gopurs_runtime.Value {
	return Take
}

func Get_toCharArray() gopurs_runtime.Value {
	return ToCharArray
}
