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
			return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer((&Constructor_Main_Foo[gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Foo
}

var cache_Main_Bar gopurs_runtime.Value
var once_Main_Bar sync.Once

func Get_Main_Bar() gopurs_runtime.Value {
	once_Main_Bar.Do(func() {
		cache_Main_Bar = gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer((*Constructor_Main_Foo[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Bar
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
		cache_Main_foo = gopurs_runtime.Func2(func(dictPartial_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foo(dictPartial_0_box, v_1_box)
		})
	})
	return cache_Main_foo
}

type Constructor_Main_Foo[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Bar[T_a any] struct {
	Rc uint32
}

func Call_Main_foo(dictPartial_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictPartial_0 gopurs_runtime.Value = dictPartial_0_loop
	_ = dictPartial_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	var __t0 *Constructor_Main_Foo[gopurs_runtime.Value]
	{
		if (v_1.Type == 9 && v_1.IntVal == 2763139640 && v_1.UnsafePtr != nil) && ((gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength((*Constructor_Main_Foo[gopurs_runtime.Value])(v_1.UnsafePtr).V0))).IntVal) == (0)) {
			__t0 = (*Constructor_Main_Foo[gopurs_runtime.Value])(nil)
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_Foo[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
	}
end_branch_0:
	return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer(__t0)}
}
