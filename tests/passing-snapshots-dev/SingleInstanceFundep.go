package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_singleInstanceFundepRow gopurs_runtime.Value
var once_Main_singleInstanceFundepRow sync.Once

func Get_Main_singleInstanceFundepRow() gopurs_runtime.Value {
	once_Main_singleInstanceFundepRow.Do(func() {
		cache_Main_singleInstanceFundepRow = gopurs_runtime.Value{Type: 9, IntVal: 4179949665, UnsafePtr: unsafe.Pointer((&Constructor_Main_SingleInstanceFundep[gopurs_runtime.Value]{1, 513803634}))}
	})
	return cache_Main_singleInstanceFundepRow
}

var cache_Main_unified gopurs_runtime.Value
var once_Main_unified sync.Once

func Get_Main_unified() gopurs_runtime.Value {
	once_Main_unified.Do(func() {
		cache_Main_unified = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_unified(dict_0_box)
		})
	})
	return cache_Main_unified
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_SingleInstanceFundep[T_r any] struct {
	Rc uint32
	V0 uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00530069006e0067006c00650049006e007300740061006e0063006500460075006e006400650070
func init() {
	gopurs_runtime.StructGetters[4179949665] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_SingleInstanceFundep[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "unified":
			return gopurs_runtime.Value{Type: 9, IntVal: int64(c.V0)}
		default:
			panic("Key not found in dictionary Constructor_Main_SingleInstanceFundep: " + key)
		}
	}
}

func Call_Main_unified(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.RecordGet(dict_0, "unified").IntVal)), UnsafePtr: nil}
}
