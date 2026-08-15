package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_go_prime_ gopurs_runtime.Value
var once_Main_go_prime_ sync.Once

func Get_Main_go_prime_() gopurs_runtime.Value {
	once_Main_go_prime_.Do(func() {
		cache_Main_go_prime_ = gopurs_runtime.Func(func(dictX_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_go_prime_(gopurs_runtime.CoerceToStruct[Constructor_Lib_X](dictX_0_box))
		})
	})
	return cache_Main_go_prime_
}

func Call_Main_go_prime_(dictX_0_loop *Constructor_Lib_X) gopurs_runtime.Value {
	var dictX_0 *Constructor_Lib_X = dictX_0_loop
	_ = dictX_0
	return gopurs_runtime.Box(dictX_0.V0)
}
