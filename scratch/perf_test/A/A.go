package A

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var foo gopurs_runtime.Value
var once_foo sync.Once
func Get_foo() gopurs_runtime.Value {
	once_foo.Do(func() {
		foo = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return foo
}


