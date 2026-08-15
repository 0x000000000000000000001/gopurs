package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_n gopurs_runtime.Value
var once_Main_n sync.Once

func Get_Main_n() gopurs_runtime.Value {
	once_Main_n.Do(func() {
		cache_Main_n = gopurs_runtime.Int(gopurs_runtime.Int(2147483648).IntVal)
	})
	return cache_Main_n
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}
