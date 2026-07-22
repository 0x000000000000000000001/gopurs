package Eg

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Bar_prime gopurs_runtime.Value
var once_Bar_prime sync.Once
func Get_Bar_prime() gopurs_runtime.Value {
	once_Bar_prime.Do(func() {
		Bar_prime = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Bar'"), "value0": value0, "value1": value1})
})
})
	})
	return Bar_prime
}


