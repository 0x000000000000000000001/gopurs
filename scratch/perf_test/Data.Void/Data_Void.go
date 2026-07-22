package Data_Void

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var absurd gopurs_runtime.Value
var once_absurd sync.Once
func Get_absurd() gopurs_runtime.Value {
	once_absurd.Do(func() {
		absurd = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
var spin_1_0 gopurs_runtime.Value
spin_1_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(spin_1_0, v_2)
})
return gopurs_runtime.Apply(spin_1_0, a_0)
})
	})
	return absurd
}


