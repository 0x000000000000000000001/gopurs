package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_bar_prime_ gopurs_runtime.Value
var once_Main_bar_prime_ sync.Once

func Get_Main_bar_prime_() gopurs_runtime.Value {
	once_Main_bar_prime_.Do(func() {
		cache_Main_bar_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 1341194597, UnsafePtr: unsafe.Pointer(&Constructor_Eg_Bar_prime_{1, 4, 5})}
	})
	return cache_Main_bar_prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}
