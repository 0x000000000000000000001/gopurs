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

var cache_Main_Foo_dollar_Dict__883817645 gopurs_runtime.Value
var once_Main_Foo_dollar_Dict__883817645 sync.Once

func Get_Main_Foo_dollar_Dict__883817645() gopurs_runtime.Value {
	once_Main_Foo_dollar_Dict__883817645.Do(func() {
		cache_Main_Foo_dollar_Dict__883817645 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer(Rebox_Main_1157739124_3584142444(Call_Main_Foo_dollar_Dict__883817645(func() struct {
				foo float64
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					foo float64
				}{}
				clone.foo = gopurs_runtime.RecordGet(orig, "foo").FloatVal()
				return clone
			}())))}
		})
	})
	return cache_Main_Foo_dollar_Dict__883817645
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_fooNumber gopurs_runtime.Value
var once_Main_fooNumber sync.Once

func Get_Main_fooNumber() gopurs_runtime.Value {
	once_Main_fooNumber.Do(func() {
		cache_Main_fooNumber = gopurs_runtime.Value{Type: 9, IntVal: 2763139640, UnsafePtr: unsafe.Pointer(Rebox_Main_1157739124_3584142444((&Constructor_Main_Foo[float64]{1, 0.0})))}
	})
	return cache_Main_fooNumber
}

var cache_Main_foo gopurs_runtime.Value
var once_Main_foo sync.Once

func Get_Main_foo() gopurs_runtime.Value {
	once_Main_foo.Do(func() {
		cache_Main_foo = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foo(dict_0_box)
		})
	})
	return cache_Main_foo
}

type Constructor_Main_Foo[T_a any] struct {
	Rc uint32
	V0 T_a
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

func Call_Main_Foo_dollar_Dict__883817645(x_0_loop struct {
	foo float64
}) *Constructor_Main_Foo[float64] {
Foo_dollar_Dict__883817645:
	for {
		if false {
			continue Foo_dollar_Dict__883817645
		}
		var x_0 struct {
			foo float64
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_Foo[float64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict1("foo", gopurs_runtime.Float(orig.foo))
		}())
	}
}

func Call_Main_foo(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "foo")
}

func Rebox_Main_1157739124_3584142444(in *Constructor_Main_Foo[float64]) *Constructor_Main_Foo[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Foo[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Float(in.V0)
	return out
}
