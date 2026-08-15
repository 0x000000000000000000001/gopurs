package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_z gopurs_runtime.Value
var once_Main_z sync.Once

func Get_Main_z() gopurs_runtime.Value {
	once_Main_z.Do(func() {
		cache_Main_z = gopurs_runtime.Int(0)
	})
	return cache_Main_z
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}
