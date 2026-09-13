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

var cache_Main_Foo_dollar_Dict__1164380134 gopurs_runtime.Value
var once_Main_Foo_dollar_Dict__1164380134 sync.Once

func Get_Main_Foo_dollar_Dict__1164380134() gopurs_runtime.Value {
	once_Main_Foo_dollar_Dict__1164380134.Do(func() {
		cache_Main_Foo_dollar_Dict__1164380134 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer(Rebox_Main_752630355_3584142444(Call_Main_Foo_dollar_Dict__1164380134(func() struct {
				foo gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					foo gopurs_runtime.Value
				}{}
				clone.foo = gopurs_runtime.RecordGet(orig, "foo")
				return clone
			}())))}
		})
	})
	return cache_Main_Foo_dollar_Dict__1164380134
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_fooString gopurs_runtime.Value
var once_Main_fooString sync.Once

func Get_Main_fooString() gopurs_runtime.Value {
	once_Main_fooString.Do(func() {
		cache_Main_fooString = gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer(Rebox_Main_752630355_3584142444((&Constructor_Main_Foo[string]{1, gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(s_0.StrVal())
		})})))}
	})
	return cache_Main_fooString
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

func Call_Main_Foo_dollar_Dict__1164380134(x_0_loop struct {
	foo gopurs_runtime.Value
}) *Constructor_Main_Foo[string] {
Foo_dollar_Dict__1164380134:
	for {
		if false {
			continue Foo_dollar_Dict__1164380134
		}
		var x_0 struct {
			foo gopurs_runtime.Value
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_Foo[string]](func() gopurs_runtime.Value {
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

func Rebox_Main_752630355_3584142444(in *Constructor_Main_Foo[string]) *Constructor_Main_Foo[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Foo[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
