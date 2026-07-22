package Def

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var what gopurs_runtime.Value
var once_what sync.Once
func Get_what() gopurs_runtime.Value {
	once_what.Do(func() {
		what = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
})
})
	})
	return what
}


