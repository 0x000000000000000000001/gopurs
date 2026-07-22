package Data_String_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Data_Char_Gen "gopurs/output/Data.Char.Gen"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
)

var max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		max = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0), y_1)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = y_1
} else {
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t1 = x_0
} else {
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t1 = x_0
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
})
})
	})
	return max
}

var genString gopurs_runtime.Value
var once_genString sync.Once
func Get_genString() gopurs_runtime.Value {
	once_genString.Do(func() {
		genString = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
unfoldable1_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_Gen.Get_unfoldable(), dictMonadRec_0), dictMonadGen_1), pkg_Data_Unfoldable.Get_unfoldableArray())
return gopurs_runtime.Func(func(genChar_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["sized"], gopurs_runtime.Func(func(size_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["chooseInt"], gopurs_runtime.Int(1)), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_max(), gopurs_runtime.Int(1)), size_5))), gopurs_runtime.Func(func(newSize_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_1.PtrVal.(map[string]gopurs_runtime.Value)["resize"], gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return newSize_6
})), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_String_CodeUnits.Get_fromCharArray()), gopurs_runtime.Apply(unfoldable1_3_1, genChar_4)))
}))
}))
})
})
})
	})
	return genString
}

var genUnicodeString gopurs_runtime.Value
var once_genUnicodeString sync.Once
func Get_genUnicodeString() gopurs_runtime.Value {
	once_genUnicodeString.Do(func() {
		genUnicodeString = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_genString(), dictMonadRec_0), dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genUnicodeChar(), dictMonadGen_1))
})
})
	})
	return genUnicodeString
}

var genDigitString gopurs_runtime.Value
var once_genDigitString sync.Once
func Get_genDigitString() gopurs_runtime.Value {
	once_genDigitString.Do(func() {
		genDigitString = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_genString(), dictMonadRec_0), dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genDigitChar(), dictMonadGen_1))
})
})
	})
	return genDigitString
}

var genAsciiString_prime gopurs_runtime.Value
var once_genAsciiString_prime sync.Once
func Get_genAsciiString_prime() gopurs_runtime.Value {
	once_genAsciiString_prime.Do(func() {
		genAsciiString_prime = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_genString(), dictMonadRec_0), dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAsciiChar_prime(), dictMonadGen_1))
})
})
	})
	return genAsciiString_prime
}

var genAsciiString gopurs_runtime.Value
var once_genAsciiString sync.Once
func Get_genAsciiString() gopurs_runtime.Value {
	once_genAsciiString.Do(func() {
		genAsciiString = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_genString(), dictMonadRec_0), dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAsciiChar(), dictMonadGen_1))
})
})
	})
	return genAsciiString
}

var genAlphaUppercaseString gopurs_runtime.Value
var once_genAlphaUppercaseString sync.Once
func Get_genAlphaUppercaseString() gopurs_runtime.Value {
	once_genAlphaUppercaseString.Do(func() {
		genAlphaUppercaseString = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_genString(), dictMonadRec_0), dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAlphaUppercase(), dictMonadGen_1))
})
})
	})
	return genAlphaUppercaseString
}

var genAlphaString gopurs_runtime.Value
var once_genAlphaString sync.Once
func Get_genAlphaString() gopurs_runtime.Value {
	once_genAlphaString.Do(func() {
		genAlphaString = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_genString(), dictMonadRec_0), dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAlpha(), dictMonadGen_1))
})
})
	})
	return genAlphaString
}

var genAlphaLowercaseString gopurs_runtime.Value
var once_genAlphaLowercaseString sync.Once
func Get_genAlphaLowercaseString() gopurs_runtime.Value {
	once_genAlphaLowercaseString.Do(func() {
		genAlphaLowercaseString = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonadGen_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_genString(), dictMonadRec_0), dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAlphaLowercase(), dictMonadGen_1))
})
})
	})
	return genAlphaLowercaseString
}


