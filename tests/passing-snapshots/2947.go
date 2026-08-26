package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Foo gopurs_runtime.Value
var once_Main_Foo sync.Once

func Get_Main_Foo() gopurs_runtime.Value {
	once_Main_Foo.Do(func() {
		cache_Main_Foo = gopurs_runtime.Value{Type: 9, IntVal: int64(2763139640), UnsafePtr: nil}
	})
	return cache_Main_Foo
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_eqFoo gopurs_runtime.Value
var once_Main_eqFoo sync.Once

func Get_Main_eqFoo() gopurs_runtime.Value {
	once_Main_eqFoo.Do(func() {
		cache_Main_eqFoo = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			})
		})}))}
	})
	return cache_Main_eqFoo
}

type Constructor_Main_Foo struct {
	Rc uint32
}
