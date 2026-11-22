package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Array_Partial_unsafeIndex gopurs_runtime.Value
var once_Data_Array_Partial_unsafeIndex sync.Once

func Get_Data_Array_Partial_unsafeIndex() gopurs_runtime.Value {
	once_Data_Array_Partial_unsafeIndex.Do(func() {
		cache_Data_Array_Partial_unsafeIndex = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_Partial_unsafeIndex(*((*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)), __local_var_1_box.IntVal)
		})
	})
	return cache_Data_Array_Partial_unsafeIndex
}

var cache_Data_Array_Partial_sub gopurs_runtime.Value
var once_Data_Array_Partial_sub sync.Once

func Get_Data_Array_Partial_sub() gopurs_runtime.Value {
	once_Data_Array_Partial_sub.Do(func() {
		cache_Data_Array_Partial_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_Array_Partial_sub
}

var cache_Data_Array_Partial_unsafeIndex1 gopurs_runtime.Value
var once_Data_Array_Partial_unsafeIndex1 sync.Once

func Get_Data_Array_Partial_unsafeIndex1() gopurs_runtime.Value {
	once_Data_Array_Partial_unsafeIndex1.Do(func() {
		cache_Data_Array_Partial_unsafeIndex1 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_Partial_unsafeIndex1(*((*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)), __local_var_1_box.IntVal)
		})
	})
	return cache_Data_Array_Partial_unsafeIndex1
}

var cache_Data_Array_Partial_tail gopurs_runtime.Value
var once_Data_Array_Partial_tail sync.Once

func Get_Data_Array_Partial_tail() gopurs_runtime.Value {
	once_Data_Array_Partial_tail.Do(func() {
		cache_Data_Array_Partial_tail = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_Partial_tail(_dollar___unused_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_Partial_tail
}

var cache_Data_Array_Partial_last gopurs_runtime.Value
var once_Data_Array_Partial_last sync.Once

func Get_Data_Array_Partial_last() gopurs_runtime.Value {
	once_Data_Array_Partial_last.Do(func() {
		cache_Data_Array_Partial_last = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_Partial_last(_dollar___unused_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)))
		})
	})
	return cache_Data_Array_Partial_last
}

var cache_Data_Array_Partial_go__init gopurs_runtime.Value
var once_Data_Array_Partial_go__init sync.Once

func Get_Data_Array_Partial_go__init() gopurs_runtime.Value {
	once_Data_Array_Partial_go__init.Do(func() {
		cache_Data_Array_Partial_go__init = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Array_Partial_go__init(_dollar___unused_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Array_Partial_go__init
}

var cache_Data_Array_Partial_head gopurs_runtime.Value
var once_Data_Array_Partial_head sync.Once

func Get_Data_Array_Partial_head() gopurs_runtime.Value {
	once_Data_Array_Partial_head.Do(func() {
		cache_Data_Array_Partial_head = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_Partial_head(_dollar___unused_0_box, *((*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)))
		})
	})
	return cache_Data_Array_Partial_head
}

func Call_Data_Array_Partial_unsafeIndex(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) gopurs_runtime.Value {
	var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 int64 = __local_var_1_loop
	_ = __local_var_1
	arrayUnsafe_value_0 := __local_var_0
	_ = arrayUnsafe_value_0
	arrayUnsafe_source_0 := arrayUnsafe_value_0
	_ = arrayUnsafe_source_0
	arrayUnsafe_index_1 := __local_var_1
	_ = arrayUnsafe_index_1
	return (arrayUnsafe_source_0)[arrayUnsafe_index_1]
}

func Call_Data_Array_Partial_unsafeIndex1(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) gopurs_runtime.Value {
	var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
	_ = __local_var_0
	var __local_var_1 int64 = __local_var_1_loop
	_ = __local_var_1
	arrayUnsafe_value_0 := __local_var_0
	_ = arrayUnsafe_value_0
	arrayUnsafe_source_0 := arrayUnsafe_value_0
	_ = arrayUnsafe_source_0
	arrayUnsafe_index_1 := __local_var_1
	_ = arrayUnsafe_index_1
	return (arrayUnsafe_source_0)[arrayUnsafe_index_1]
}

func Call_Data_Array_Partial_tail(_dollar___unused_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(int64(len(xs_1))), gopurs_runtime.Array(xs_1)).UnsafePtr))
}

func Call_Data_Array_Partial_last(_dollar___unused_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	arrayUnsafe_value_0 := xs_1
	_ = arrayUnsafe_value_0
	arrayUnsafe_source_0 := arrayUnsafe_value_0
	_ = arrayUnsafe_source_0
	arrayUnsafe_index_1 := gopurs_runtime.IntSub(gopurs_runtime.Int(int64(len(xs_1))).IntVal, int64(1))
	_ = arrayUnsafe_index_1
	return (arrayUnsafe_source_0)[arrayUnsafe_index_1]
}

func Call_Data_Array_Partial_go__init(_dollar___unused_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	return *((*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(gopurs_runtime.IntSub(gopurs_runtime.Int(int64(len(xs_1))).IntVal, int64(1))), gopurs_runtime.Array(xs_1)).UnsafePtr))
}

func Call_Data_Array_Partial_head(_dollar___unused_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var xs_1 []gopurs_runtime.Value = xs_1_loop
	_ = xs_1
	arrayUnsafe_value_0 := xs_1
	_ = arrayUnsafe_value_0
	arrayUnsafe_source_0 := arrayUnsafe_value_0
	_ = arrayUnsafe_source_0
	arrayUnsafe_index_1 := int64(0)
	_ = arrayUnsafe_index_1
	return (arrayUnsafe_source_0)[arrayUnsafe_index_1]
}
