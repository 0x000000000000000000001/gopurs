package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Control_Monad_ST_Global_toEffect gopurs_runtime.Value
var once_Control_Monad_ST_Global_toEffect sync.Once

func Get_Control_Monad_ST_Global_toEffect() gopurs_runtime.Value {
	once_Control_Monad_ST_Global_toEffect.Do(func() {
		cache_Control_Monad_ST_Global_toEffect = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Control_Monad_ST_Global_toEffect
}
