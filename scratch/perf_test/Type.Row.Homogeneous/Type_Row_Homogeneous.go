package Type_Row_Homogeneous

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var homogeneousRowListNil gopurs_runtime.Value
var once_homogeneousRowListNil sync.Once
func Get_homogeneousRowListNil() gopurs_runtime.Value {
	once_homogeneousRowListNil.Do(func() {
		homogeneousRowListNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return homogeneousRowListNil
}

var homogeneousRowListCons gopurs_runtime.Value
var once_homogeneousRowListCons sync.Once
func Get_homogeneousRowListCons() gopurs_runtime.Value {
	once_homogeneousRowListCons.Do(func() {
		homogeneousRowListCons = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictTypeEquals_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
	})
	return homogeneousRowListCons
}

var homogeneous gopurs_runtime.Value
var once_homogeneous sync.Once
func Get_homogeneous() gopurs_runtime.Value {
	once_homogeneous.Do(func() {
		homogeneous = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
	})
	return homogeneous
}


