package Coercible_Lib

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var NTLib3 gopurs_runtime.Value
var once_NTLib3 sync.Once
func Get_NTLib3() gopurs_runtime.Value {
	once_NTLib3.Do(func() {
		NTLib3 = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return NTLib3
}

var NTLib1 gopurs_runtime.Value
var once_NTLib1 sync.Once
func Get_NTLib1() gopurs_runtime.Value {
	once_NTLib1.Do(func() {
		NTLib1 = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return NTLib1
}


