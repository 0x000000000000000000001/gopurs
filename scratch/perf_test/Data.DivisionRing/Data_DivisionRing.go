package Data_DivisionRing

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Ring "gopurs/output/Data.Ring"
)

var recip gopurs_runtime.Value
var once_recip sync.Once
func Get_recip() gopurs_runtime.Value {
	once_recip.Do(func() {
		recip = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["recip"]
})
	})
	return recip
}

var rightDiv gopurs_runtime.Value
var once_rightDiv sync.Once
func Get_rightDiv() gopurs_runtime.Value {
	once_rightDiv.Do(func() {
		rightDiv = gopurs_runtime.Func(func(dictDivisionRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictDivisionRing_0.PtrVal.(map[string]gopurs_runtime.Value)["Ring0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["mul"], a_1), gopurs_runtime.Apply(dictDivisionRing_0.PtrVal.(map[string]gopurs_runtime.Value)["recip"], b_2))
})
})
})
	})
	return rightDiv
}

var leftDiv gopurs_runtime.Value
var once_leftDiv sync.Once
func Get_leftDiv() gopurs_runtime.Value {
	once_leftDiv.Do(func() {
		leftDiv = gopurs_runtime.Func(func(dictDivisionRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictDivisionRing_0.PtrVal.(map[string]gopurs_runtime.Value)["Ring0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["mul"], gopurs_runtime.Apply(dictDivisionRing_0.PtrVal.(map[string]gopurs_runtime.Value)["recip"], b_2)), a_1)
})
})
})
	})
	return leftDiv
}

var divisionringNumber gopurs_runtime.Value
var once_divisionringNumber sync.Once
func Get_divisionringNumber() gopurs_runtime.Value {
	once_divisionringNumber.Do(func() {
		divisionringNumber = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"recip": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_numDiv(), gopurs_runtime.Float(1.0)), x_0)
}), "Ring0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ring.Get_ringNumber()
})})
	})
	return divisionringNumber
}


