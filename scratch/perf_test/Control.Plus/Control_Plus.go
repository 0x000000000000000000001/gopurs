package Control_Plus

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Alt "gopurs/output/Control.Alt"
)

var plusArray gopurs_runtime.Value
var once_plusArray sync.Once
func Get_plusArray() gopurs_runtime.Value {
	once_plusArray.Do(func() {
		plusArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Array([]gopurs_runtime.Value{}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Alt.Get_altArray()
})})
	})
	return plusArray
}

var empty gopurs_runtime.Value
var once_empty sync.Once
func Get_empty() gopurs_runtime.Value {
	once_empty.Do(func() {
		empty = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
})
	})
	return empty
}


