package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Boolean_otherwise gopurs_runtime.Value
var once_Data_Boolean_otherwise sync.Once

func Get_Data_Boolean_otherwise() gopurs_runtime.Value {
	once_Data_Boolean_otherwise.Do(func() {
		cache_Data_Boolean_otherwise = gopurs_runtime.Bool(true)
	})
	return cache_Data_Boolean_otherwise
}
