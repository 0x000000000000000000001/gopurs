package IxMonad

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var pure gopurs_runtime.Value
var once_pure sync.Once
func Get_pure() gopurs_runtime.Value {
	once_pure.Do(func() {
		pure = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"]
})
	})
	return pure
}

var bind gopurs_runtime.Value
var once_bind sync.Once
func Get_bind() gopurs_runtime.Value {
	once_bind.Do(func() {
		bind = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"]
})
	})
	return bind
}


