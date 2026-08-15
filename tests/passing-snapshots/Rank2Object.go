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
		cache_Main_Foo = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer((&Constructor_Main_Foo{1, value0}))}
		})
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

var cache_Main_foo gopurs_runtime.Value
var once_Main_foo sync.Once

func Get_Main_foo() gopurs_runtime.Value {
	once_Main_foo.Do(func() {
		cache_Main_foo = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_foo(gopurs_runtime.CoerceToStruct[Constructor_Main_Foo](v_0_box)))
		})
	})
	return cache_Main_foo
}

type Constructor_Main_Foo struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_foo(v_0_loop *Constructor_Main_Foo) float64 {
	var v_0 *Constructor_Main_Foo = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer(v_0)}, "id"), gopurs_runtime.Float(0.0)).FloatVal()
}
