package Control_Biapplicative

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Biapply "gopurs/output/Control.Biapply"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var bipure gopurs_runtime.Value
var once_bipure sync.Once
func Get_bipure() gopurs_runtime.Value {
	once_bipure.Do(func() {
		bipure = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["bipure"]
})
	})
	return bipure
}

var biapplicativeTuple gopurs_runtime.Value
var once_biapplicativeTuple sync.Once
func Get_biapplicativeTuple() gopurs_runtime.Value {
	once_biapplicativeTuple.Do(func() {
		biapplicativeTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bipure": pkg_Data_Tuple.Get_Tuple(), "Biapply0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Biapply.Get_biapplyTuple()
})})
	})
	return biapplicativeTuple
}


