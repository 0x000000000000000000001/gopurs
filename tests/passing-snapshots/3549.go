package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_map_prime_ gopurs_runtime.Value
var once_Main_map_prime_ sync.Once

func Get_Main_map_prime_() gopurs_runtime.Value {
	once_Main_map_prime_.Do(func() {
		cache_Main_map_prime_ = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_map_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box))
		})
	})
	return cache_Main_map_prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_identity(x_0_box)
		})
	})
	return cache_Main_identity
}

func Call_Main_map_prime_(dictFunctor_0_loop *Constructor_Data_Functor_Functor) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Box(dictFunctor_0.V0)
}

func Call_Main_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
