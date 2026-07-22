package Symbols

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Test_Assert "gopurs/output/Test.Assert"
)

var lowS gopurs_runtime.Value
var once_lowS sync.Once
func Get_lowS() gopurs_runtime.Value {
	once_lowS.Do(func() {
		lowS = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
	})
	return lowS
}

var notReplacing gopurs_runtime.Value
var once_notReplacing sync.Once
func Get_notReplacing() gopurs_runtime.Value {
	once_notReplacing.Do(func() {
		notReplacing = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqBooleanImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqStringImpl(), gopurs_runtime.Str("�")), gopurs_runtime.Str("�"))), gopurs_runtime.Bool(false))
	})
	return notReplacing
}

var highS gopurs_runtime.Value
var once_highS sync.Once
func Get_highS() gopurs_runtime.Value {
	once_highS.Do(func() {
		highS = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
	})
	return highS
}

var loneSurrogates gopurs_runtime.Value
var once_loneSurrogates sync.Once
func Get_loneSurrogates() gopurs_runtime.Value {
	once_loneSurrogates.Do(func() {
		loneSurrogates = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqStringImpl(), gopurs_runtime.Str("𝌆")), gopurs_runtime.Str("𝌆"))
	})
	return loneSurrogates
}

var outOfOrderSurrogates gopurs_runtime.Value
var once_outOfOrderSurrogates sync.Once
func Get_outOfOrderSurrogates() gopurs_runtime.Value {
	once_outOfOrderSurrogates.Do(func() {
		outOfOrderSurrogates = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqStringImpl(), gopurs_runtime.Str("��")), gopurs_runtime.Str("��"))
	})
	return outOfOrderSurrogates
}

var main gopurs_runtime.Value
var once_main sync.Once
func Get_main() gopurs_runtime.Value {
	once_main.Do(func() {
		main = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Test_Assert.Get_assertImpl(), gopurs_runtime.Str("lone surrogates may be combined into a surrogate pair")), Get_loneSurrogates())), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Test_Assert.Get_assertImpl(), gopurs_runtime.Str("lone surrogates may be combined out of order to remain lone surrogates")), Get_outOfOrderSurrogates())), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Test_Assert.Get_assertImpl(), gopurs_runtime.Str("lone surrogates are not replaced with the Unicode replacement character U+FFFD")), Get_notReplacing())), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Done"))
}))
}))
}))
	})
	return main
}


