package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			_dollar___unused_0_0 := gopurs_runtime.Apply(Get_Records_main(), gopurs_runtime.Value{})
			_ = _dollar___unused_0_0
			return gopurs_runtime.Apply(Get_Symbols_main(), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}
