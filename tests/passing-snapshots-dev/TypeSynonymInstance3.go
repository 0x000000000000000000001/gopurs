package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_ltEqD8D256 gopurs_runtime.Value
var once_Main_ltEqD8D256 sync.Once

func Get_Main_ltEqD8D256() gopurs_runtime.Value {
	once_Main_ltEqD8D256.Do(func() {
		cache_Main_ltEqD8D256 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_ltEqD8D256
}

var cache_Main_lte256 gopurs_runtime.Value
var once_Main_lte256 sync.Once

func Get_Main_lte256() gopurs_runtime.Value {
	once_Main_lte256.Do(func() {
		cache_Main_lte256 = gopurs_runtime.Value{Type: 9, IntVal: 907985874, UnsafePtr: unsafe.Pointer((&Constructor_Main_Lte256[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Main_lte256
}

type Constructor_Main_LtEq[T_a any, T_b any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004c007400450071
func init() {
	gopurs_runtime.StructGetters[3632241042] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_LtEq[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_LtEq: " + key)
		}
	}
}

type Constructor_Main_Lte256[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f004c00740065003200350036
func init() {
	gopurs_runtime.StructGetters[907985874] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Lte256[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "LtEq0":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Lte256: " + key)
		}
	}
}
