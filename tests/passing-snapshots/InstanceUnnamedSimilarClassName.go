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

var cache_Main_ClassName_dollar_Dict gopurs_runtime.Value
var once_Main_ClassName_dollar_Dict sync.Once

func Get_Main_ClassName_dollar_Dict() gopurs_runtime.Value {
	once_Main_ClassName_dollar_Dict.Do(func() {
		cache_Main_ClassName_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1438891703, UnsafePtr: unsafe.Pointer(Call_Main_ClassName_dollar_Dict(func() struct {
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
	return cache_Main_ClassName_dollar_Dict
}

var cache_Main_ClassName_dollar_Dict__3209752335 gopurs_runtime.Value
var once_Main_ClassName_dollar_Dict__3209752335 sync.Once

func Get_Main_ClassName_dollar_Dict__3209752335() gopurs_runtime.Value {
	once_Main_ClassName_dollar_Dict__3209752335.Do(func() {
		cache_Main_ClassName_dollar_Dict__3209752335 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1438891703, UnsafePtr: unsafe.Pointer(Rebox_Main_1049220494_954425347(Call_Main_ClassName_dollar_Dict__3209752335(func() struct {
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
	return cache_Main_ClassName_dollar_Dict__3209752335
}

var cache_Main_classNameFoo gopurs_runtime.Value
var once_Main_classNameFoo sync.Once

func Get_Main_classNameFoo() gopurs_runtime.Value {
	once_Main_classNameFoo.Do(func() {
		cache_Main_classNameFoo = gopurs_runtime.Value{Type: 9, IntVal: 1687123273, UnsafePtr: unsafe.Pointer(Rebox_Main_2145685968_2100274173((&Constructor_ImportedClassName_ClassName[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(int64(0))
		})})))}
	})
	return cache_Main_classNameFoo
}

var cache_Main_classNameFoo1 gopurs_runtime.Value
var once_Main_classNameFoo1 sync.Once

func Get_Main_classNameFoo1() gopurs_runtime.Value {
	once_Main_classNameFoo1.Do(func() {
		cache_Main_classNameFoo1 = gopurs_runtime.Value{Type: 9, IntVal: 1438891703, UnsafePtr: unsafe.Pointer(Rebox_Main_1049220494_954425347((&Constructor_Main_ClassName[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(int64(0))
		})})))}
	})
	return cache_Main_classNameFoo1
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
			return Call_Main_foo(gopurs_runtime.CoerceToStruct[Constructor_Main_ClassName[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_foo
}

type Constructor_Main_Foo struct {
	Rc uint32
}

type Constructor_Main_ClassName[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1438891703] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ClassName[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "foo":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_ClassName: " + key)
		}
	}
}

func Call_Main_ClassName_dollar_Dict(x_0_loop struct {
	foo gopurs_runtime.Value
}) *Constructor_Main_ClassName[gopurs_runtime.Value] {
	var x_0 struct {
		foo gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_ClassName[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("foo", orig.foo)
	}())
}

func Call_Main_ClassName_dollar_Dict__3209752335(x_0_loop struct {
	foo gopurs_runtime.Value
}) *Constructor_Main_ClassName[uint32] {
ClassName_dollar_Dict__3209752335:
	for {
		if false {
			continue ClassName_dollar_Dict__3209752335
		}
		var x_0 struct {
			foo gopurs_runtime.Value
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_ClassName[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict1("foo", orig.foo)
		}())
	}
}

func Call_Main_foo(dict_0_loop *Constructor_Main_ClassName[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_ClassName[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Main_1049220494_954425347(in *Constructor_Main_ClassName[uint32]) *Constructor_Main_ClassName[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_ClassName[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2145685968_2100274173(in *Constructor_ImportedClassName_ClassName[uint32]) *Constructor_ImportedClassName_ClassName[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_ImportedClassName_ClassName[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
