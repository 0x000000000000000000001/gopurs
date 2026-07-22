package Data_Array_Partial

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Array "gopurs/output/Data.Array"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Function_Uncurried "gopurs/output/Data.Function.Uncurried"
)

var unsafeIndex gopurs_runtime.Value
var once_unsafeIndex sync.Once
func Get_unsafeIndex() gopurs_runtime.Value {
	once_unsafeIndex.Do(func() {
		unsafeIndex = gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), pkg_Data_Array.Get_unsafeIndexImpl())
	})
	return unsafeIndex
}

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array.Get_slice(), gopurs_runtime.Int(1)), gopurs_runtime.Apply(pkg_Data_Array.Get_length(), xs_1)), xs_1)
})
})
	})
	return tail
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeIndex(), xs_1), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(pkg_Data_Array.Get_length(), xs_1)), gopurs_runtime.Int(1)))
})
})
	})
	return last
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array.Get_slice(), gopurs_runtime.Int(0)), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Apply(pkg_Data_Array.Get_length(), xs_1)), gopurs_runtime.Int(1))), xs_1)
})
})
	})
	return init_
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeIndex(), xs_1), gopurs_runtime.Int(0))
})
})
	})
	return head
}


