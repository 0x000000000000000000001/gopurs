package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_id_prime_ gopurs_runtime.Value
var once_Main_id_prime_ sync.Once

func Get_Main_id_prime_() gopurs_runtime.Value {
	once_Main_id_prime_.Do(func() {
		cache_Main_id_prime_ = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		}), gopurs_runtime.Func(func(y_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return y_0
		}))
	})
	return cache_Main_id_prime_
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Main_id_prime_(), gopurs_runtime.Str("Done")).StrVal()))
	})
	return cache_Main_main
}
