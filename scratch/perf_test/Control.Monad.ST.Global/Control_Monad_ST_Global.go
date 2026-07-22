package Control_Monad_ST_Global

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var toEffect gopurs_runtime.Value
var once_toEffect sync.Once
func Get_toEffect() gopurs_runtime.Value {
	once_toEffect.Do(func() {
		toEffect = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return toEffect
}


