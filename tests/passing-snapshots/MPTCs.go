package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_NullaryTypeClass_dollar_Dict gopurs_runtime.Value
var once_Main_NullaryTypeClass_dollar_Dict sync.Once

func Get_Main_NullaryTypeClass_dollar_Dict() gopurs_runtime.Value {
	once_Main_NullaryTypeClass_dollar_Dict.Do(func() {
		cache_Main_NullaryTypeClass_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2059088857, UnsafePtr: unsafe.Pointer(Call_Main_NullaryTypeClass_dollar_Dict(func() struct {
				greeting string
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					greeting string
				}{}
				clone.greeting = gopurs_runtime.RecordGet(orig, "greeting").StrVal()
				return clone
			}()))}
		})
	})
	return cache_Main_NullaryTypeClass_dollar_Dict
}

var cache_Main_Coerce_dollar_Dict gopurs_runtime.Value
var once_Main_Coerce_dollar_Dict sync.Once

func Get_Main_Coerce_dollar_Dict() gopurs_runtime.Value {
	once_Main_Coerce_dollar_Dict.Do(func() {
		cache_Main_Coerce_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2049652419, UnsafePtr: unsafe.Pointer(Call_Main_Coerce_dollar_Dict(func() struct {
				coerce gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					coerce gopurs_runtime.Value
				}{}
				clone.coerce = gopurs_runtime.RecordGet(orig, "coerce")
				return clone
			}()))}
		})
	})
	return cache_Main_Coerce_dollar_Dict
}

var cache_Main_nullaryTypeClass gopurs_runtime.Value
var once_Main_nullaryTypeClass sync.Once

func Get_Main_nullaryTypeClass() gopurs_runtime.Value {
	once_Main_nullaryTypeClass.Do(func() {
		cache_Main_nullaryTypeClass = gopurs_runtime.Value{Type: 9, IntVal: 2059088857, UnsafePtr: unsafe.Pointer((&Constructor_Main_NullaryTypeClass{1, "Hello, World!"}))}
	})
	return cache_Main_nullaryTypeClass
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_greeting gopurs_runtime.Value
var once_Main_greeting sync.Once

func Get_Main_greeting() gopurs_runtime.Value {
	once_Main_greeting.Do(func() {
		cache_Main_greeting = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_greeting(dict_0_box)
		})
	})
	return cache_Main_greeting
}

var cache_Main_coerceShow gopurs_runtime.Value
var once_Main_coerceShow sync.Once

func Get_Main_coerceShow() gopurs_runtime.Value {
	once_Main_coerceShow.Do(func() {
		cache_Main_coerceShow = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_coerceShow(dictShow_0_box)
		})
	})
	return cache_Main_coerceShow
}

var cache_Main_coerceRefl gopurs_runtime.Value
var once_Main_coerceRefl sync.Once

func Get_Main_coerceRefl() gopurs_runtime.Value {
	once_Main_coerceRefl.Do(func() {
		cache_Main_coerceRefl = gopurs_runtime.Value{Type: 9, IntVal: 2049652419, UnsafePtr: unsafe.Pointer((&Constructor_Main_Coerce[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return a_0
		})}))}
	})
	return cache_Main_coerceRefl
}

var cache_Main_coerce gopurs_runtime.Value
var once_Main_coerce sync.Once

func Get_Main_coerce() gopurs_runtime.Value {
	once_Main_coerce.Do(func() {
		cache_Main_coerce = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_coerce(gopurs_runtime.CoerceToStruct[Constructor_Main_Coerce[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_coerce
}

type Constructor_Main_NullaryTypeClass struct {
	Rc uint32
	V0 string
}

func init() {
	gopurs_runtime.StructGetters[2059088857] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_NullaryTypeClass)(ptr)
		_ = c
		switch key {
		case "greeting":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_NullaryTypeClass: " + key)
		}
	}
}

type Constructor_Main_Coerce[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2049652419] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Coerce[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "coerce":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Coerce: " + key)
		}
	}
}

func Call_Main_NullaryTypeClass_dollar_Dict(x_0_loop struct {
	greeting string
}) *Constructor_Main_NullaryTypeClass {
	var x_0 struct {
		greeting string
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_NullaryTypeClass](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("greeting", gopurs_runtime.Str(orig.greeting))
	}())
}

func Call_Main_Coerce_dollar_Dict(x_0_loop struct {
	coerce gopurs_runtime.Value
}) *Constructor_Main_Coerce[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		coerce gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Coerce[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("coerce", orig.coerce)
	}())
}

func Call_Main_greeting(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Str(gopurs_runtime.RecordGet(dict_0, "greeting").StrVal())
}

func Call_Main_coerceShow(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
	_ = dictShow_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2049652419, UnsafePtr: unsafe.Pointer(Rebox_Main_1681742382_2651832753((&Constructor_Main_Coerce[gopurs_runtime.Value, string]{1, Call_Data_Show_show(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0))})))}
}

func Call_Main_coerce(dict_0_loop *Constructor_Main_Coerce[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Coerce[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Main_1681742382_2651832753(in *Constructor_Main_Coerce[gopurs_runtime.Value, string]) *Constructor_Main_Coerce[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Coerce[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
