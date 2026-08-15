package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Test gopurs_runtime.Value
var once_Main_Test sync.Once

func Get_Main_Test() gopurs_runtime.Value {
	once_Main_Test.Do(func() {
		cache_Main_Test = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test(x_0_box)
		})
	})
	return cache_Main_Test
}

var cache_Main_First gopurs_runtime.Value
var once_Main_First sync.Once

func Get_Main_First() gopurs_runtime.Value {
	once_Main_First.Do(func() {
		cache_Main_First = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_First(x_0_box)
		})
	})
	return cache_Main_First
}

var cache_Main_newtypeTest gopurs_runtime.Value
var once_Main_newtypeTest sync.Once

func Get_Main_newtypeTest() gopurs_runtime.Value {
	once_Main_newtypeTest.Do(func() {
		cache_Main_newtypeTest = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})}
	})
	return cache_Main_newtypeTest
}

var cache_Main_t gopurs_runtime.Value
var once_Main_t sync.Once

func Get_Main_t() gopurs_runtime.Value {
	once_Main_t.Do(func() {
		cache_Main_t = gopurs_runtime.Str("hello")
	})
	return cache_Main_t
}

var cache_Main_newtypeFirst gopurs_runtime.Value
var once_Main_newtypeFirst sync.Once

func Get_Main_newtypeFirst() gopurs_runtime.Value {
	once_Main_newtypeFirst.Do(func() {
		cache_Main_newtypeFirst = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})}
	})
	return cache_Main_newtypeFirst
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Int(1)
	})
	return cache_Main_f
}

var cache_Main_i gopurs_runtime.Value
var once_Main_i sync.Once

func Get_Main_i() gopurs_runtime.Value {
	once_Main_i.Do(func() {
		cache_Main_i = gopurs_runtime.Int(1)
	})
	return cache_Main_i
}

var cache_Main_a gopurs_runtime.Value
var once_Main_a sync.Once

func Get_Main_a() gopurs_runtime.Value {
	once_Main_a.Do(func() {
		cache_Main_a = gopurs_runtime.Str("hello")
	})
	return cache_Main_a
}

func Call_Main_Test(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_First(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
