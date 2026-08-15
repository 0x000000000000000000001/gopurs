package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(asgård_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f(asgård_0_box)
		})
	})
	return cache_Main_f
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Str("Done").StrVal()))
	})
	return cache_Main_main
}

func Call_Main_f(asgård_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var asgård_0 gopurs_runtime.Value = asgård_0_loop
	_ = asgård_0
	return asgård_0
}
