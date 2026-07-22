package InitializationError

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var alpha gopurs_runtime.Value
var once_alpha sync.Once
func Get_alpha() gopurs_runtime.Value {
	once_alpha.Do(func() {
		alpha = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["alpha"]
})
	})
	return alpha
}

var charlieAlpha gopurs_runtime.Value
var once_charlieAlpha sync.Once
func Get_charlieAlpha() gopurs_runtime.Value {
	once_charlieAlpha.Do(func() {
		charlieAlpha = gopurs_runtime.Func(func(dictCharlie_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictCharlie_0.PtrVal.(map[string]gopurs_runtime.Value)["Bravo0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Alpha0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["alpha"]
})
	})
	return charlieAlpha
}

var charlieArray gopurs_runtime.Value
var once_charlieArray sync.Once
func Get_charlieArray() gopurs_runtime.Value {
	once_charlieArray.Do(func() {
		charlieArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Bravo0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bravoArray()
})})
	})
	return charlieArray
}

var bravoArray gopurs_runtime.Value
var once_bravoArray sync.Once
func Get_bravoArray() gopurs_runtime.Value {
	once_bravoArray.Do(func() {
		bravoArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Alpha0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_alphaArray()
})})
	})
	return bravoArray
}

var alphaArray gopurs_runtime.Value
var once_alphaArray sync.Once
func Get_alphaArray() gopurs_runtime.Value {
	once_alphaArray.Do(func() {
		alphaArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alpha": Get_alphaArray().PtrVal.(map[string]gopurs_runtime.Value)["alpha"]})
	})
	return alphaArray
}


