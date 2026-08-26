package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_abuseUnused gopurs_runtime.Value
var once_Main_abuseUnused sync.Once

func Get_Main_abuseUnused() gopurs_runtime.Value {
	once_Main_abuseUnused.Do(func() {
		cache_Main_abuseUnused = gopurs_runtime.Func(func(__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_abuseUnused(__unused_0_box)
		})
	})
	return cache_Main_abuseUnused
}

var cache_Main_abuseUnused__2674436160 gopurs_runtime.Value
var once_Main_abuseUnused__2674436160 sync.Once

func Get_Main_abuseUnused__2674436160() gopurs_runtime.Value {
	once_Main_abuseUnused__2674436160.Do(func() {
		cache_Main_abuseUnused__2674436160 = gopurs_runtime.Func(func(__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_abuseUnused__2674436160(__unused_0_box)
		})
	})
	return cache_Main_abuseUnused__2674436160
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_abuseUnused(__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var __unused_0 gopurs_runtime.Value = __unused_0_loop
	_ = __unused_0
	return __unused_0
}

func Call_Main_abuseUnused__2674436160(__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var __unused_0 gopurs_runtime.Value = __unused_0_loop
	_ = __unused_0
	return __unused_0
}
