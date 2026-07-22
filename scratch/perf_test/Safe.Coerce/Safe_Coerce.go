package Safe_Coerce

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var coerce gopurs_runtime.Value
var once_coerce sync.Once
func Get_coerce() gopurs_runtime.Value {
	once_coerce.Do(func() {
		coerce = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
})
	})
	return coerce
}


