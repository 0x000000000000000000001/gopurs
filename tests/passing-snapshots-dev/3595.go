package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_showString gopurs_runtime.Value
var once_Main_showString sync.Once

func Get_Main_showString() gopurs_runtime.Value {
	once_Main_showString.Do(func() {
		cache_Main_showString = gopurs_runtime.Value{Type: 9, IntVal: 3143145725, UnsafePtr: unsafe.Pointer(Rebox_Main_2177390102_590014985((&Constructor_Main_Show[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(Rebox_Main_1386611502_1514099793(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showString()))))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})})))}
	})
	return cache_Main_showString
}

var cache_Main_id gopurs_runtime.Value
var once_Main_id sync.Once

func Get_Main_id() gopurs_runtime.Value {
	once_Main_id.Do(func() {
		cache_Main_id = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_id(gopurs_runtime.CoerceToStruct[Constructor_Main_Show[gopurs_runtime.Value]](dict_0_box))
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

type Constructor_Main_Show[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00530068006f0077
func init() {
	gopurs_runtime.StructGetters[3143145725] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Show[gopurs_runtime.Value])(ptr)
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

func Call_Main_id(dict_0_loop *Constructor_Main_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Show[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Rebox_Main_1386611502_1514099793(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[string])(unsafe.Pointer(in))
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2177390102_590014985(in *Constructor_Main_Show[string]) *Constructor_Main_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}
