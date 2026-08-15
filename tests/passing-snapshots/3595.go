package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Show_dollar_Dict gopurs_runtime.Value
var once_Main_Show_dollar_Dict sync.Once

func Get_Main_Show_dollar_Dict() gopurs_runtime.Value {
	once_Main_Show_dollar_Dict.Do(func() {
		cache_Main_Show_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Show_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Show_dollar_Dict
}

var cache_Main_showString gopurs_runtime.Value
var once_Main_showString sync.Once

func Get_Main_showString() gopurs_runtime.Value {
	once_Main_showString.Do(func() {
		cache_Main_showString = gopurs_runtime.Value{Type: 9, IntVal: 3143145725, UnsafePtr: unsafe.Pointer(&Constructor_Main_Show{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Data_Show_showString()))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(x_0.StrVal())
		})})}
	})
	return cache_Main_showString
}

var cache_Main_id gopurs_runtime.Value
var once_Main_id sync.Once

func Get_Main_id() gopurs_runtime.Value {
	once_Main_id.Do(func() {
		cache_Main_id = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_id(gopurs_runtime.CoerceToStruct[Constructor_Main_Show](dict_0_box))
		})
	})
	return cache_Main_id
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_id__7826678 gopurs_runtime.Value
var once_Main_id__7826678 sync.Once

func Get_Main_id__7826678() gopurs_runtime.Value {
	once_Main_id__7826678.Do(func() {
		cache_Main_id__7826678 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_id__7826678(gopurs_runtime.CoerceToStruct[Constructor_Main_Show](dict_0_box))
		})
	})
	return cache_Main_id__7826678
}

type Constructor_Main_Show struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[3143145725] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Show)(ptr)
		_ = c
		switch key {
		case "Show0":
			return gopurs_runtime.Box(c.V0)
		case "id":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Show: " + key)
		}
	}
}

func Call_Main_Show_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_id(dict_0_loop *Constructor_Main_Show) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Show = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_id__7826678(dict_0_loop *Constructor_Main_Show) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Show = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}
