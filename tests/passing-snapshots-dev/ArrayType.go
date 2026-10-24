package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_pointedArray gopurs_runtime.Value
var once_Main_pointedArray sync.Once

func Get_Main_pointedArray() gopurs_runtime.Value {
	once_Main_pointedArray.Do(func() {
		cache_Main_pointedArray = gopurs_runtime.Value{Type: 9, IntVal: 4236620371, UnsafePtr: unsafe.Pointer((&Constructor_Main_Pointed[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array([]gopurs_runtime.Value{a_0})
		})}))}
	})
	return cache_Main_pointedArray
}

var cache_Main_point gopurs_runtime.Value
var once_Main_point sync.Once

func Get_Main_point() gopurs_runtime.Value {
	once_Main_point.Do(func() {
		cache_Main_point = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_point(gopurs_runtime.CoerceToStruct[Constructor_Main_Pointed[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_point
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Pointed[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0050006f0069006e007400650064
func init() {
	gopurs_runtime.StructGetters[4236620371] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Pointed[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "point":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Pointed: " + key)
		}
	}
}

func Call_Main_point(dict_0_loop *Constructor_Main_Pointed[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Pointed[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}
