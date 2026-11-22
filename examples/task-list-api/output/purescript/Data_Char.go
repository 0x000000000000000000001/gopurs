package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Char_toCharCode gopurs_runtime.Value
var once_Data_Char_toCharCode sync.Once

func Get_Data_Char_toCharCode() gopurs_runtime.Value {
	once_Data_Char_toCharCode.Do(func() {
		cache_Data_Char_toCharCode = Get_Data_Enum_toCharCode()
	})
	return cache_Data_Char_toCharCode
}

var cache_Data_Char_fromCharCode gopurs_runtime.Value
var once_Data_Char_fromCharCode sync.Once

func Get_Data_Char_fromCharCode() gopurs_runtime.Value {
	once_Data_Char_fromCharCode.Do(func() {
		cache_Data_Char_fromCharCode = Get_Data_Enum_charToEnum()
	})
	return cache_Data_Char_fromCharCode
}
