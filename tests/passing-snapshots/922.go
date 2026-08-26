package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_I gopurs_runtime.Value
var once_Main_I sync.Once

func Get_Main_I() gopurs_runtime.Value {
	once_Main_I.Do(func() {
		cache_Main_I = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_I
}

var cache_Main_Default_dollar_Dict gopurs_runtime.Value
var once_Main_Default_dollar_Dict sync.Once

func Get_Main_Default_dollar_Dict() gopurs_runtime.Value {
	once_Main_Default_dollar_Dict.Do(func() {
		cache_Main_Default_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Default_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Default_dollar_Dict
}

var cache_Main_defaultString gopurs_runtime.Value
var once_Main_defaultString sync.Once

func Get_Main_defaultString() gopurs_runtime.Value {
	once_Main_defaultString.Do(func() {
		cache_Main_defaultString = gopurs_runtime.Value{Type: 9, IntVal: 1853528597, UnsafePtr: unsafe.Pointer((&Constructor_Main_Default[string]{1, gopurs_runtime.Str("Done")}))}
	})
	return cache_Main_defaultString
}

var cache_Main_def gopurs_runtime.Value
var once_Main_def sync.Once

func Get_Main_def() gopurs_runtime.Value {
	once_Main_def.Do(func() {
		cache_Main_def = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_def(dict_0_box)
		})
	})
	return cache_Main_def
}

var cache_Main_def__193435443 gopurs_runtime.Value
var once_Main_def__193435443 sync.Once

func Get_Main_def__193435443() gopurs_runtime.Value {
	once_Main_def__193435443.Do(func() {
		cache_Main_def__193435443 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_def__193435443(dict_0_box)
		})
	})
	return cache_Main_def__193435443
}

var cache_Main_def__3240737238 gopurs_runtime.Value
var once_Main_def__3240737238 sync.Once

func Get_Main_def__3240737238() gopurs_runtime.Value {
	once_Main_def__3240737238.Do(func() {
		cache_Main_def__3240737238 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_def__3240737238(dict_0_box)
		})
	})
	return cache_Main_def__3240737238
}

var cache_Main_defaultI gopurs_runtime.Value
var once_Main_defaultI sync.Once

func Get_Main_defaultI() gopurs_runtime.Value {
	once_Main_defaultI.Do(func() {
		cache_Main_defaultI = gopurs_runtime.Func(func(dictDefault_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_defaultI(dictDefault_0_box)
		})
	})
	return cache_Main_defaultI
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_I[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Default[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1853528597] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Default[any])(ptr)
		_ = c
		switch key {
		case "def":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Default: " + key)
		}
	}
}

func Call_Main_Default_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_def(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "def")
}

func Call_Main_def__193435443(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "def")
}

func Call_Main_def__3240737238(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "def")
}

func Call_Main_defaultI(dictDefault_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictDefault_0 gopurs_runtime.Value = dictDefault_0_loop
	_ = dictDefault_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1853528597, UnsafePtr: unsafe.Pointer((&Constructor_Main_Default[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictDefault_0, "def")}))}
}
