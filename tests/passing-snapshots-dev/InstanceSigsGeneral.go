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

var cache_Main_eqNumber gopurs_runtime.Value
var once_Main_eqNumber sync.Once

func Get_Main_eqNumber() gopurs_runtime.Value {
	once_Main_eqNumber.Do(func() {
		cache_Main_eqNumber = gopurs_runtime.Value{Type: 9, IntVal: 61300330, UnsafePtr: unsafe.Pointer(Rebox_Main_3136802470_2946378686((&Constructor_Main_Eq[float64]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})})))}
	})
	return cache_Main_eqNumber
}

var cache_Main_eq gopurs_runtime.Value
var once_Main_eq sync.Once

func Get_Main_eq() gopurs_runtime.Value {
	once_Main_eq.Do(func() {
		cache_Main_eq = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eq(gopurs_runtime.CoerceToStruct[Constructor_Main_Eq[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_eq
}

type Constructor_Main_Eq[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00450071
func init() {
	gopurs_runtime.StructGetters[61300330] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Eq[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "eq":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Eq: " + key)
		}
	}
}

func Call_Main_eq(dict_0_loop *Constructor_Main_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Eq[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Main_3136802470_2946378686(in *Constructor_Main_Eq[float64]) *Constructor_Main_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}
