package Data_Field

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var field gopurs_runtime.Value
var once_field sync.Once
func Get_field() gopurs_runtime.Value {
	once_field.Do(func() {
		field = gopurs_runtime.Func(func(dictEuclideanRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictDivisionRing_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"EuclideanRing0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEuclideanRing_0
}), "DivisionRing1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return dictDivisionRing_1
})})
})
})
	})
	return field
}


