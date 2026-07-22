package Control_Alt

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var altArray gopurs_runtime.Value
var once_altArray sync.Once
func Get_altArray() gopurs_runtime.Value {
	once_altArray.Do(func() {
		altArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": pkg_Data_Semigroup.Get_concatArray(), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
})})
	})
	return altArray
}

var alt gopurs_runtime.Value
var once_alt sync.Once
func Get_alt() gopurs_runtime.Value {
	once_alt.Do(func() {
		alt = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["alt"]
})
	})
	return alt
}


