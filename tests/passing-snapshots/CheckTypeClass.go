package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Bar gopurs_runtime.Value
var once_Main_Bar sync.Once

func Get_Main_Bar() gopurs_runtime.Value {
	once_Main_Bar.Do(func() {
		cache_Main_Bar = gopurs_runtime.Value{Type: 9, IntVal: int64(2512729583), UnsafePtr: nil}
	})
	return cache_Main_Bar
}

var cache_Main_Foo_dollar_Dict gopurs_runtime.Value
var once_Main_Foo_dollar_Dict sync.Once

func Get_Main_Foo_dollar_Dict() gopurs_runtime.Value {
	once_Main_Foo_dollar_Dict.Do(func() {
		cache_Main_Foo_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer(Call_Main_Foo_dollar_Dict(func() struct {
				foo gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					foo gopurs_runtime.Value
				}{}
				clone.foo = gopurs_runtime.RecordGet(orig, "foo")
				return clone
			}()))}
		})
	})
	return cache_Main_Foo_dollar_Dict
}

var cache_Main_mkBar gopurs_runtime.Value
var once_Main_mkBar sync.Once

func Get_Main_mkBar() gopurs_runtime.Value {
	once_Main_mkBar.Do(func() {
		cache_Main_mkBar = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_mkBar(v_0_box)), UnsafePtr: nil}
		})
	})
	return cache_Main_mkBar
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
		cache_Main_foo = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foo(gopurs_runtime.CoerceToStruct[Constructor_Main_Foo[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_foo
}

var cache_Main_foo_ gopurs_runtime.Value
var once_Main_foo_ sync.Once

func Get_Main_foo_() gopurs_runtime.Value {
	once_Main_foo_.Do(func() {
		cache_Main_foo_ = gopurs_runtime.Func2(func(dictFoo_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foo_(gopurs_runtime.CoerceToStruct[Constructor_Main_Foo[gopurs_runtime.Value]](dictFoo_0_box), x_1_box)
		})
	})
	return cache_Main_foo_
}

type Constructor_Main_Bar[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Foo[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2763139640] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Foo[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "foo":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Foo: " + key)
		}
	}
}

func Call_Main_Foo_dollar_Dict(x_0_loop struct {
	foo gopurs_runtime.Value
}) *Constructor_Main_Foo[gopurs_runtime.Value] {
	var x_0 struct {
		foo gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Foo[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("foo", orig.foo)
	}())
}

func Call_Main_mkBar(v_0_loop gopurs_runtime.Value) uint32 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return 2512729583
}

func Call_Main_foo(dict_0_loop *Constructor_Main_Foo[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Foo[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_foo_(dictFoo_0_loop *Constructor_Main_Foo[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoo_0 *Constructor_Main_Foo[gopurs_runtime.Value] = dictFoo_0_loop
	_ = dictFoo_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(Call_Main_foo(dictFoo_0), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Func(func(dictFoo1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return Get_Main_mkBar()
	}), gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer(dictFoo_0)}, x_1).IntVal)), UnsafePtr: nil})
}
