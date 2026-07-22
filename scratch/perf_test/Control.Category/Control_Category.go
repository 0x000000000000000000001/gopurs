package Control_Category

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
)

var identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		identity = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["identity"]
})
	})
	return identity
}

var categoryFn gopurs_runtime.Value
var once_categoryFn sync.Once
func Get_categoryFn() gopurs_runtime.Value {
	once_categoryFn.Do(func() {
		categoryFn = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"identity": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}), "Semigroupoid0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Semigroupoid.Get_semigroupoidFn()
})})
	})
	return categoryFn
}


