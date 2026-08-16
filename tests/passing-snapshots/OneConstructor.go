package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_One gopurs_runtime.Value
var once_Main_One sync.Once

func Get_Main_One() gopurs_runtime.Value {
	once_Main_One.Do(func() {
		cache_Main_One = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_One
}

var cache_Main_one_prime_ gopurs_runtime.Value
var once_Main_one_prime_ sync.Once

func Get_Main_one_prime_() gopurs_runtime.Value {
	once_Main_one_prime_.Do(func() {
		cache_Main_one_prime_ = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_one_prime_(v_0_box)
		})
	})
	return cache_Main_one_prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_One struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_one_prime_(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}
