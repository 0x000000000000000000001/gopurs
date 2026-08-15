package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_yield gopurs_runtime.Value
var once_Main_yield sync.Once

func Get_Main_yield() gopurs_runtime.Value {
	once_Main_yield.Do(func() {
		cache_Main_yield = gopurs_runtime.Int(0)
	})
	return cache_Main_yield
}

var cache_Main_this gopurs_runtime.Value
var once_Main_this sync.Once

func Get_Main_this() gopurs_runtime.Value {
	once_Main_this.Do(func() {
		cache_Main_this = gopurs_runtime.Func(func(catch_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_this(catch_0_box)
		})
	})
	return cache_Main_this
}

var cache_Main_public gopurs_runtime.Value
var once_Main_public sync.Once

func Get_Main_public() gopurs_runtime.Value {
	once_Main_public.Do(func() {
		cache_Main_public = gopurs_runtime.Func(func(go__return_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_public(go__return_0_box)
		})
	})
	return cache_Main_public
}

var cache_Main_member gopurs_runtime.Value
var once_Main_member sync.Once

func Get_Main_member() gopurs_runtime.Value {
	once_Main_member.Do(func() {
		cache_Main_member = gopurs_runtime.Int(1)
	})
	return cache_Main_member
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_this(catch_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var catch_0 gopurs_runtime.Value = catch_0_loop
	_ = catch_0
	return catch_0
}

func Call_Main_public(go__return_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var go__return_0 gopurs_runtime.Value = go__return_0_loop
	_ = go__return_0
	return go__return_0
}
