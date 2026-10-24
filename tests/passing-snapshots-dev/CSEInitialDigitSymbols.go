package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_twoThirty gopurs_runtime.Value
var once_Main_twoThirty sync.Once

func Get_Main_twoThirty() gopurs_runtime.Value {
	once_Main_twoThirty.Do(func() {
		cache_Main_twoThirty = gopurs_runtime.Str("2:30")
	})
	return cache_Main_twoThirty
}

var cache_Main_two gopurs_runtime.Value
var once_Main_two sync.Once

func Get_Main_two() gopurs_runtime.Value {
	once_Main_two.Do(func() {
		cache_Main_two = gopurs_runtime.Str("2")
	})
	return cache_Main_two
}

var cache_Main_reflectSymbol_prime_ gopurs_runtime.Value
var once_Main_reflectSymbol_prime_ sync.Once

func Get_Main_reflectSymbol_prime_() gopurs_runtime.Value {
	once_Main_reflectSymbol_prime_.Do(func() {
		cache_Main_reflectSymbol_prime_ = gopurs_runtime.Func(func(dictIsSymbol_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_reflectSymbol_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_0_box))
		})
	})
	return cache_Main_reflectSymbol_prime_
}

var cache_Main_two2 gopurs_runtime.Value
var once_Main_two2 sync.Once

func Get_Main_two2() gopurs_runtime.Value {
	once_Main_two2.Do(func() {
		cache_Main_two2 = gopurs_runtime.Str("2")
	})
	return cache_Main_two2
}

var cache_Main_twoThirty2 gopurs_runtime.Value
var once_Main_twoThirty2 sync.Once

func Get_Main_twoThirty2() gopurs_runtime.Value {
	once_Main_twoThirty2.Do(func() {
		cache_Main_twoThirty2 = gopurs_runtime.Str("2:30")
	})
	return cache_Main_twoThirty2
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_reflectSymbol_prime_(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	return Call_Data_Symbol_reflectSymbol(dictIsSymbol_0)
}
