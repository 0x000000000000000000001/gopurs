package Record_Unsafe_Union

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Function_Uncurried "gopurs/output/Data.Function.Uncurried"
)

var unsafeUnion gopurs_runtime.Value
var once_unsafeUnion sync.Once
func Get_unsafeUnion() gopurs_runtime.Value {
	once_unsafeUnion.Do(func() {
		unsafeUnion = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_unsafeUnionFn())
	})
	return unsafeUnion
}

func Get_unsafeUnionFn() gopurs_runtime.Value {
	return UnsafeUnionFn
}
