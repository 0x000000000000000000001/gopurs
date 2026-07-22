package Records

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Test_Assert "gopurs/output/Test.Assert"
)

var LoneSurrogateKeys gopurs_runtime.Value
var once_LoneSurrogateKeys sync.Once
func Get_LoneSurrogateKeys() gopurs_runtime.Value {
	once_LoneSurrogateKeys.Do(func() {
		LoneSurrogateKeys = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return LoneSurrogateKeys
}

var AstralKeys gopurs_runtime.Value
var once_AstralKeys sync.Once
func Get_AstralKeys() gopurs_runtime.Value {
	once_AstralKeys.Do(func() {
		AstralKeys = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return AstralKeys
}

var testLoneSurrogateKeys gopurs_runtime.Value
var once_testLoneSurrogateKeys sync.Once
func Get_testLoneSurrogateKeys() gopurs_runtime.Value {
	once_testLoneSurrogateKeys.Do(func() {
		testLoneSurrogateKeys = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Test_Assert.Get_assertImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("lone surrogate keys: ")), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Int(5)))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), gopurs_runtime.Int(5)), gopurs_runtime.Int(5)))
	})
	return testLoneSurrogateKeys
}

var testAstralKeys gopurs_runtime.Value
var once_testAstralKeys sync.Once
func Get_testAstralKeys() gopurs_runtime.Value {
	once_testAstralKeys.Do(func() {
		testAstralKeys = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Test_Assert.Get_assertImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("astral keys: ")), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Int(5)))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), gopurs_runtime.Int(5)), gopurs_runtime.Int(5)))
	})
	return testAstralKeys
}

var main gopurs_runtime.Value
var once_main sync.Once
func Get_main() gopurs_runtime.Value {
	once_main.Do(func() {
		main = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), Get_testLoneSurrogateKeys()), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_bindE(), Get_testAstralKeys()), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Done"))
}))
}))
	})
	return main
}


