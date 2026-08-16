package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_id_prime_ gopurs_runtime.Value
var once_Main_id_prime_ sync.Once

func Get_Main_id_prime_() gopurs_runtime.Value {
	once_Main_id_prime_.Do(func() {
		cache_Main_id_prime_ = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_id_prime_(x_0_box)
		})
	})
	return cache_Main_id_prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_id_prime_(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
