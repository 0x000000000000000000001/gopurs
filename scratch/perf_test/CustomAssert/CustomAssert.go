package CustomAssert

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var assertThrows gopurs_runtime.Value
var once_assertThrows sync.Once
func Get_assertThrows() gopurs_runtime.Value {
	once_assertThrows.Do(func() {
		assertThrows = gopurs_runtime.Apply(Get_assertThrowsImpl(), pkg_Data_Unit.Get_unit())
	})
	return assertThrows
}

func Get_assertThrowsImpl() gopurs_runtime.Value {
	return AssertThrowsImpl
}
