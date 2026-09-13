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
			return gopurs_runtime.Value{Type: 9, IntVal: 1853528597, UnsafePtr: unsafe.Pointer(Call_Main_Default_dollar_Dict(func() struct {
				def gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					def gopurs_runtime.Value
				}{}
				clone.def = gopurs_runtime.RecordGet(orig, "def")
				return clone
			}()))}
		})
	})
	return cache_Main_Default_dollar_Dict
}

var cache_Main_Default_dollar_Dict__629339585 gopurs_runtime.Value
var once_Main_Default_dollar_Dict__629339585 sync.Once

func Get_Main_Default_dollar_Dict__629339585() gopurs_runtime.Value {
	once_Main_Default_dollar_Dict__629339585.Do(func() {
		cache_Main_Default_dollar_Dict__629339585 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1853528597, UnsafePtr: unsafe.Pointer(Rebox_Main_998170654_401051297(Call_Main_Default_dollar_Dict__629339585(func() struct {
				def string
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					def string
				}{}
				clone.def = gopurs_runtime.RecordGet(orig, "def").StrVal()
				return clone
			}())))}
		})
	})
	return cache_Main_Default_dollar_Dict__629339585
}

var cache_Main_defaultString gopurs_runtime.Value
var once_Main_defaultString sync.Once

func Get_Main_defaultString() gopurs_runtime.Value {
	once_Main_defaultString.Do(func() {
		cache_Main_defaultString = gopurs_runtime.Value{Type: 9, IntVal: 1853528597, UnsafePtr: unsafe.Pointer(Rebox_Main_998170654_401051297((&Constructor_Main_Default[string]{1, "Done"})))}
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

var cache_Main_def__1122847620 gopurs_runtime.Value
var once_Main_def__1122847620 sync.Once

func Get_Main_def__1122847620() gopurs_runtime.Value {
	once_Main_def__1122847620.Do(func() {
		cache_Main_def__1122847620 = gopurs_runtime.RecordGet(Call_Main_defaultI(gopurs_runtime.Value{Type: 9, IntVal: 1853528597, UnsafePtr: unsafe.Pointer(Rebox_Main_998170654_401051297(Rebox_Main_401051297_998170654(gopurs_runtime.CoerceToStruct[Constructor_Main_Default[gopurs_runtime.Value]](Get_Main_defaultString()))))}), "def")
	})
	return cache_Main_def__1122847620
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(Get_Main_def__1122847620().StrVal()))
	})
	return cache_Main_main
}

type Constructor_Main_I[T_a any] struct {
	Rc uint32
	V0 T_a
}

type Constructor_Main_Default[T_a any] struct {
	Rc uint32
	V0 T_a
}

func init() {
	gopurs_runtime.StructGetters[1853528597] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Default[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "def":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Default: " + key)
		}
	}
}

func Call_Main_Default_dollar_Dict(x_0_loop struct {
	def gopurs_runtime.Value
}) *Constructor_Main_Default[gopurs_runtime.Value] {
	var x_0 struct {
		def gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Default[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("def", orig.def)
	}())
}

func Call_Main_Default_dollar_Dict__629339585(x_0_loop struct {
	def string
}) *Constructor_Main_Default[string] {
Default_dollar_Dict__629339585:
	for {
		if false {
			continue Default_dollar_Dict__629339585
		}
		var x_0 struct {
			def string
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_Default[string]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict1("def", gopurs_runtime.Str(orig.def))
		}())
	}
}

func Call_Main_def(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "def")
}

func Call_Main_defaultI(dictDefault_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictDefault_0 gopurs_runtime.Value = dictDefault_0_loop
	_ = dictDefault_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1853528597, UnsafePtr: unsafe.Pointer((&Constructor_Main_Default[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictDefault_0, "def")}))}
}

func Rebox_Main_401051297_998170654(in *Constructor_Main_Default[gopurs_runtime.Value]) *Constructor_Main_Default[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Default[string]{}
	out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Main_998170654_401051297(in *Constructor_Main_Default[string]) *Constructor_Main_Default[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Default[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}
