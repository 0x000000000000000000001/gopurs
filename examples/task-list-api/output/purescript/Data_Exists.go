package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Exists_runExists gopurs_runtime.Value
var once_Data_Exists_runExists sync.Once

func Get_Data_Exists_runExists() gopurs_runtime.Value {
	once_Data_Exists_runExists.Do(func() {
		cache_Data_Exists_runExists = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Exists_runExists
}

var cache_Data_Exists_mkExists gopurs_runtime.Value
var once_Data_Exists_mkExists sync.Once

func Get_Data_Exists_mkExists() gopurs_runtime.Value {
	once_Data_Exists_mkExists.Do(func() {
		cache_Data_Exists_mkExists = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Exists_mkExists
}
