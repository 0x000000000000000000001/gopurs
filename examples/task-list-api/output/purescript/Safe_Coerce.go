package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Safe_Coerce_coerce gopurs_runtime.Value
var once_Safe_Coerce_coerce sync.Once

func Get_Safe_Coerce_coerce() gopurs_runtime.Value {
	once_Safe_Coerce_coerce.Do(func() {
		cache_Safe_Coerce_coerce = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Safe_Coerce_coerce(_dollar___unused_0_box)
		})
	})
	return cache_Safe_Coerce_coerce
}

func Call_Safe_Coerce_coerce(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Get_Unsafe_Coerce_unsafeCoerce()
}
