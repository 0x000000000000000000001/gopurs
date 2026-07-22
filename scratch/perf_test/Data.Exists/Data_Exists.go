package Data_Exists

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var runExists gopurs_runtime.Value
var once_runExists sync.Once
func Get_runExists() gopurs_runtime.Value {
	once_runExists.Do(func() {
		runExists = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return runExists
}

var mkExists gopurs_runtime.Value
var once_mkExists sync.Once
func Get_mkExists() gopurs_runtime.Value {
	once_mkExists.Do(func() {
		mkExists = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return mkExists
}


