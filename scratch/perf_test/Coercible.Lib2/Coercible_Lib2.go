package Coercible_Lib2

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var NTLib2 gopurs_runtime.Value
var once_NTLib2 sync.Once
func Get_NTLib2() gopurs_runtime.Value {
	once_NTLib2.Do(func() {
		NTLib2 = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return NTLib2
}


