package SolvingIsSymbol_Lib

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var literalSymbol gopurs_runtime.Value
var once_literalSymbol sync.Once
func Get_literalSymbol() gopurs_runtime.Value {
	once_literalSymbol.Do(func() {
		literalSymbol = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
	})
	return literalSymbol
}

var libReflectSymbol gopurs_runtime.Value
var once_libReflectSymbol sync.Once
func Get_libReflectSymbol() gopurs_runtime.Value {
	once_libReflectSymbol.Do(func() {
		libReflectSymbol = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"]
})
	})
	return libReflectSymbol
}


