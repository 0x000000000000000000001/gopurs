package IxApplicative

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

var map_ gopurs_runtime.Value
var once_map_ sync.Once
func Get_map_() gopurs_runtime.Value {
	once_map_.Do(func() {
		map_ = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["map"]
})
	})
	return map_
}

var apply gopurs_runtime.Value
var once_apply sync.Once
func Get_apply() gopurs_runtime.Value {
	once_apply.Do(func() {
		apply = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"]
})
	})
	return apply
}


