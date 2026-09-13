package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_X gopurs_runtime.Value
var once_Main_X sync.Once

func Get_Main_X() gopurs_runtime.Value {
	once_Main_X.Do(func() {
		cache_Main_X = gopurs_runtime.Value{Type: 9, IntVal: int64(1409933510), UnsafePtr: nil}
	})
	return cache_Main_X
}

var cache_Main_Y gopurs_runtime.Value
var once_Main_Y sync.Once

func Get_Main_Y() gopurs_runtime.Value {
	once_Main_Y.Do(func() {
		cache_Main_Y = gopurs_runtime.Value{Type: 9, IntVal: int64(1682951303), UnsafePtr: nil}
	})
	return cache_Main_Y
}

var cache_Main_Show_dollar_Dict gopurs_runtime.Value
var once_Main_Show_dollar_Dict sync.Once

func Get_Main_Show_dollar_Dict() gopurs_runtime.Value {
	once_Main_Show_dollar_Dict.Do(func() {
		cache_Main_Show_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3143145725, UnsafePtr: unsafe.Pointer(Call_Main_Show_dollar_Dict(func() struct {
				noshow gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					noshow gopurs_runtime.Value
				}{}
				clone.noshow = gopurs_runtime.RecordGet(orig, "noshow")
				return clone
			}()))}
		})
	})
	return cache_Main_Show_dollar_Dict
}

var cache_Main_show gopurs_runtime.Value
var once_Main_show sync.Once

func Get_Main_show() gopurs_runtime.Value {
	once_Main_show.Do(func() {
		cache_Main_show = gopurs_runtime.Float(1.0)
	})
	return cache_Main_show
}

var cache_Main_noshow gopurs_runtime.Value
var once_Main_noshow sync.Once

func Get_Main_noshow() gopurs_runtime.Value {
	once_Main_noshow.Do(func() {
		cache_Main_noshow = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_noshow(gopurs_runtime.CoerceToStruct[Constructor_Main_Show[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_noshow
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(1.0)).StrVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

type Constructor_Main_X struct {
	Rc uint32
}

type Constructor_Main_Y struct {
	Rc uint32
}

type Constructor_Main_Show[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[3143145725] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Show[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "noshow":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Show: " + key)
		}
	}
}

func Call_Main_Show_dollar_Dict(x_0_loop struct {
	noshow gopurs_runtime.Value
}) *Constructor_Main_Show[gopurs_runtime.Value] {
	var x_0 struct {
		noshow gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Show[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("noshow", orig.noshow)
	}())
}

func Call_Main_noshow(dict_0_loop *Constructor_Main_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Show[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}
