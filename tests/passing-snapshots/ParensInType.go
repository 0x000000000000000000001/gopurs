package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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

var cache_Main_Foo_dollar_Dict__1254635007 gopurs_runtime.Value
var once_Main_Foo_dollar_Dict__1254635007 sync.Once

func Get_Main_Foo_dollar_Dict__1254635007() gopurs_runtime.Value {
	once_Main_Foo_dollar_Dict__1254635007.Do(func() {
		cache_Main_Foo_dollar_Dict__1254635007 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer(Call_Main_Foo_dollar_Dict__1254635007(func() struct {
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
	return cache_Main_Foo_dollar_Dict__1254635007
}

var cache_Main_fooLogEff gopurs_runtime.Value
var once_Main_fooLogEff sync.Once

func Get_Main_fooLogEff() gopurs_runtime.Value {
	once_Main_fooLogEff.Do(func() {
		cache_Main_fooLogEff = gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer((&Constructor_Main_Foo[gopurs_runtime.Value]{1, Get_Effect_Console_log()}))}
	})
	return cache_Main_fooLogEff
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

var cache_Main_foo__2657262998 gopurs_runtime.Value
var once_Main_foo__2657262998 sync.Once

func Get_Main_foo__2657262998() gopurs_runtime.Value {
	once_Main_foo__2657262998.Do(func() {
		cache_Main_foo__2657262998 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foo__2657262998(__eta_norm_0_0_box.StrVal())
		})
	})
	return cache_Main_foo__2657262998
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
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

func Call_Main_Foo_dollar_Dict__1254635007(x_0_loop struct {
	foo gopurs_runtime.Value
}) *Constructor_Main_Foo[gopurs_runtime.Value] {
Foo_dollar_Dict__1254635007:
	for {
		if false {
			continue Foo_dollar_Dict__1254635007
		}
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
}

func Call_Main_foo(dict_0_loop *Constructor_Main_Foo[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Foo[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_foo__2657262998(__eta_norm_0_0_loop string) gopurs_runtime.Value {
foo__2657262998:
	for {
		if false {
			continue foo__2657262998
		}
		var __eta_norm_0_0 string = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__eta_norm_0_0))
	}
}
