package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_fn gopurs_runtime.Value
var once_Main_fn sync.Once

func Get_Main_fn() gopurs_runtime.Value {
	once_Main_fn.Do(func() {
		cache_Main_fn = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_fn(a_0_box.StrVal()))
		})
	})
	return cache_Main_fn
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

func Call_Main_fn(a_0_loop string) string {
	var a_0 string = a_0_loop
	_ = a_0
	return a_0
}
