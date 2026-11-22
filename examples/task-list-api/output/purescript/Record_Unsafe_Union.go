package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Record_Unsafe_Union_unsafeUnionFn gopurs_runtime.Value
var once_Record_Unsafe_Union_unsafeUnionFn sync.Once

func Get_Record_Unsafe_Union_unsafeUnionFn() gopurs_runtime.Value {
	once_Record_Unsafe_Union_unsafeUnionFn.Do(func() {
		cache_Record_Unsafe_Union_unsafeUnionFn = Get_Record_Unsafe_Union__unsafeUnionFn()
	})
	return cache_Record_Unsafe_Union_unsafeUnionFn
}

var cache_Record_Unsafe_Union_unsafeUnion gopurs_runtime.Value
var once_Record_Unsafe_Union_unsafeUnion sync.Once

func Get_Record_Unsafe_Union_unsafeUnion() gopurs_runtime.Value {
	once_Record_Unsafe_Union_unsafeUnion.Do(func() {
		cache_Record_Unsafe_Union_unsafeUnion = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_Unsafe_Union_unsafeUnion(__local_var_0_box, __local_var_1_box)
		})
	})
	return cache_Record_Unsafe_Union_unsafeUnion
}

func Call_Record_Unsafe_Union_unsafeUnion(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
	_ = __local_var_1
	return gopurs_runtime.UncurriedApp2(Get_Record_Unsafe_Union_unsafeUnionFn(), __local_var_0, __local_var_1)
}

func Get_Record_Unsafe_Union__unsafeUnionFn() gopurs_runtime.Value {
	return _Gopurs_Record_Unsafe_Union__UnsafeUnionFn
}
