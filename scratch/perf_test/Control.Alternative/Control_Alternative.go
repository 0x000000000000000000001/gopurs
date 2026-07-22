package Control_Alternative

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Plus "gopurs/output/Control.Plus"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var guard gopurs_runtime.Value
var once_guard sync.Once
func Get_guard() gopurs_runtime.Value {
	once_guard.Do(func() {
		guard = gopurs_runtime.Func(func(dictAlternative_0 gopurs_runtime.Value) gopurs_runtime.Value {
empty_1_0 := gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["empty"]
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
if (v_2).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlternative_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], pkg_Data_Unit.Get_unit())
} else {
__t1 = empty_1_0
}
return __t1
})
})
	})
	return guard
}

var alternativeArray gopurs_runtime.Value
var once_alternativeArray sync.Once
func Get_alternativeArray() gopurs_runtime.Value {
	once_alternativeArray.Do(func() {
		alternativeArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray()
}), "Plus1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Plus.Get_plusArray()
})})
	})
	return alternativeArray
}


