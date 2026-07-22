package Data_Boolean

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var otherwise gopurs_runtime.Value
var once_otherwise sync.Once
func Get_otherwise() gopurs_runtime.Value {
	once_otherwise.Do(func() {
		otherwise = gopurs_runtime.Bool(true)
	})
	return otherwise
}


