package Data_Char

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Enum "gopurs/output/Data.Enum"
)

var toCharCode gopurs_runtime.Value
var once_toCharCode sync.Once
func Get_toCharCode() gopurs_runtime.Value {
	once_toCharCode.Do(func() {
		toCharCode = pkg_Data_Enum.Get_toCharCode()
	})
	return toCharCode
}

var fromCharCode gopurs_runtime.Value
var once_fromCharCode sync.Once
func Get_fromCharCode() gopurs_runtime.Value {
	once_fromCharCode.Do(func() {
		fromCharCode = pkg_Data_Enum.Get_charToEnum()
	})
	return fromCharCode
}


