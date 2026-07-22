package Data_Profunctor_Closed

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
)

var closedFunction gopurs_runtime.Value
var once_closedFunction sync.Once
func Get_closedFunction() gopurs_runtime.Value {
	once_closedFunction.Do(func() {
		closedFunction = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"closed": pkg_Control_Semigroupoid.Get_semigroupoidFn().PtrVal.(map[string]gopurs_runtime.Value)["compose"], "Profunctor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Profunctor.Get_profunctorFn()
})})
	})
	return closedFunction
}

var closed gopurs_runtime.Value
var once_closed sync.Once
func Get_closed() gopurs_runtime.Value {
	once_closed.Do(func() {
		closed = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["closed"]
})
	})
	return closed
}


