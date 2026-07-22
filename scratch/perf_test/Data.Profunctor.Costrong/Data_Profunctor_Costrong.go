package Data_Profunctor_Costrong

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var unsecond gopurs_runtime.Value
var once_unsecond sync.Once
func Get_unsecond() gopurs_runtime.Value {
	once_unsecond.Do(func() {
		unsecond = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["unsecond"]
})
	})
	return unsecond
}

var unfirst gopurs_runtime.Value
var once_unfirst sync.Once
func Get_unfirst() gopurs_runtime.Value {
	once_unfirst.Do(func() {
		unfirst = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["unfirst"]
})
	})
	return unfirst
}


