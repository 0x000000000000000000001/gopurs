package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Partial_crashWith gopurs_runtime.Value
var once_Partial_crashWith sync.Once

func Get_Partial_crashWith() gopurs_runtime.Value {
	once_Partial_crashWith.Do(func() {
		cache_Partial_crashWith = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Partial_crashWith(_dollar___unused_0_box)
		})
	})
	return cache_Partial_crashWith
}

var cache_Partial_crashWith1 gopurs_runtime.Value
var once_Partial_crashWith1 sync.Once

func Get_Partial_crashWith1() gopurs_runtime.Value {
	once_Partial_crashWith1.Do(func() {
		cache_Partial_crashWith1 = Get_Partial__crashWith()
	})
	return cache_Partial_crashWith1
}

var cache_Partial_crash gopurs_runtime.Value
var once_Partial_crash sync.Once

func Get_Partial_crash() gopurs_runtime.Value {
	once_Partial_crash.Do(func() {
		cache_Partial_crash = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Partial_crash(_dollar___unused_0_box)
		})
	})
	return cache_Partial_crash
}

func Call_Partial_crashWith(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Get_Partial__crashWith()
}

func Call_Partial_crash(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.StrWithCodeUnits("Partial.crash: partial function", 31))
}

func Get_Partial__crashWith() gopurs_runtime.Value {
	return _Gopurs_Partial__CrashWith
}
