package Data_Char_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
)

var toEnumWithDefaults gopurs_runtime.Value
var once_toEnumWithDefaults sync.Once
func Get_toEnumWithDefaults() gopurs_runtime.Value {
	once_toEnumWithDefaults.Do(func() {
		toEnumWithDefaults = gopurs_runtime.Apply(pkg_Data_Enum.Get_toEnumWithDefaults(), pkg_Data_Enum.Get_boundedEnumChar())
	})
	return toEnumWithDefaults
}

var foldable1NonEmpty gopurs_runtime.Value
var once_foldable1NonEmpty sync.Once
func Get_foldable1NonEmpty() gopurs_runtime.Value {
	once_foldable1NonEmpty.Do(func() {
		foldable1NonEmpty = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), pkg_Data_Foldable.Get_foldableArray())
	})
	return foldable1NonEmpty
}

var genUnicodeChar gopurs_runtime.Value
var once_genUnicodeChar sync.Once
func Get_genUnicodeChar() gopurs_runtime.Value {
	once_genUnicodeChar.Do(func() {
		genUnicodeChar = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(gopurs_runtime.Apply(Get_toEnumWithDefaults(), pkg_Data_Bounded.Get_bottomChar()), pkg_Data_Bounded.Get_topChar())), gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseInt"], gopurs_runtime.Int(0)), gopurs_runtime.Int(65536)))
})
	})
	return genUnicodeChar
}

var genDigitChar gopurs_runtime.Value
var once_genDigitChar sync.Once
func Get_genDigitChar() gopurs_runtime.Value {
	once_genDigitChar.Do(func() {
		genDigitChar = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(gopurs_runtime.Apply(Get_toEnumWithDefaults(), pkg_Data_Bounded.Get_bottomChar()), pkg_Data_Bounded.Get_topChar())), gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseInt"], gopurs_runtime.Int(48)), gopurs_runtime.Int(57)))
})
	})
	return genDigitChar
}

var genAsciiChar_prime gopurs_runtime.Value
var once_genAsciiChar_prime sync.Once
func Get_genAsciiChar_prime() gopurs_runtime.Value {
	once_genAsciiChar_prime.Do(func() {
		genAsciiChar_prime = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(gopurs_runtime.Apply(Get_toEnumWithDefaults(), pkg_Data_Bounded.Get_bottomChar()), pkg_Data_Bounded.Get_topChar())), gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseInt"], gopurs_runtime.Int(0)), gopurs_runtime.Int(127)))
})
	})
	return genAsciiChar_prime
}

var genAsciiChar gopurs_runtime.Value
var once_genAsciiChar sync.Once
func Get_genAsciiChar() gopurs_runtime.Value {
	once_genAsciiChar.Do(func() {
		genAsciiChar = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(gopurs_runtime.Apply(Get_toEnumWithDefaults(), pkg_Data_Bounded.Get_bottomChar()), pkg_Data_Bounded.Get_topChar())), gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseInt"], gopurs_runtime.Int(32)), gopurs_runtime.Int(127)))
})
	})
	return genAsciiChar
}

var genAlphaUppercase gopurs_runtime.Value
var once_genAlphaUppercase sync.Once
func Get_genAlphaUppercase() gopurs_runtime.Value {
	once_genAlphaUppercase.Do(func() {
		genAlphaUppercase = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(gopurs_runtime.Apply(Get_toEnumWithDefaults(), pkg_Data_Bounded.Get_bottomChar()), pkg_Data_Bounded.Get_topChar())), gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseInt"], gopurs_runtime.Int(65)), gopurs_runtime.Int(90)))
})
	})
	return genAlphaUppercase
}

var genAlphaLowercase gopurs_runtime.Value
var once_genAlphaLowercase sync.Once
func Get_genAlphaLowercase() gopurs_runtime.Value {
	once_genAlphaLowercase.Do(func() {
		genAlphaLowercase = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(gopurs_runtime.Apply(Get_toEnumWithDefaults(), pkg_Data_Bounded.Get_bottomChar()), pkg_Data_Bounded.Get_topChar())), gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseInt"], gopurs_runtime.Int(97)), gopurs_runtime.Int(122)))
})
	})
	return genAlphaLowercase
}

var genAlpha gopurs_runtime.Value
var once_genAlpha sync.Once
func Get_genAlpha() gopurs_runtime.Value {
	once_genAlpha.Do(func() {
		genAlpha = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Control_Monad_Gen.Get_oneOf(), dictMonadGen_0), Get_foldable1NonEmpty()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("NonEmpty"), "value0": gopurs_runtime.Apply(Get_genAlphaLowercase(), dictMonadGen_0), "value1": gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply(Get_genAlphaUppercase(), dictMonadGen_0)})}))
})
	})
	return genAlpha
}


