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

var cache_Main_baz gopurs_runtime.Value
var once_Main_baz sync.Once

func Get_Main_baz() gopurs_runtime.Value {
	once_Main_baz.Do(func() {
		cache_Main_baz = gopurs_runtime.Int(3)
	})
	return cache_Main_baz
}
