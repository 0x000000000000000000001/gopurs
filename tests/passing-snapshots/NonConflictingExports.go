package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_thing gopurs_runtime.Value
var once_Main_thing sync.Once

func Get_Main_thing() gopurs_runtime.Value {
	once_Main_thing.Do(func() {
		cache_Main_thing = gopurs_runtime.Int(2)
	})
	return cache_Main_thing
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}
