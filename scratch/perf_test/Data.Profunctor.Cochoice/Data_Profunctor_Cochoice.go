package Data_Profunctor_Cochoice

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var unright gopurs_runtime.Value
var once_unright sync.Once
func Get_unright() gopurs_runtime.Value {
	once_unright.Do(func() {
		unright = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["unright"]
})
	})
	return unright
}

var unleft gopurs_runtime.Value
var once_unleft sync.Once
func Get_unleft() gopurs_runtime.Value {
	once_unleft.Do(func() {
		unleft = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["unleft"]
})
	})
	return unleft
}


