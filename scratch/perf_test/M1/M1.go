package M1

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var applyFn gopurs_runtime.Value
var once_applyFn sync.Once
func Get_applyFn() gopurs_runtime.Value {
	once_applyFn.Do(func() {
		applyFn = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, a_1)
})
})
	})
	return applyFn
}


