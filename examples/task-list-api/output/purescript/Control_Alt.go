package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Alt_Alt_dollar_Dict gopurs_runtime.Value
var once_Control_Alt_Alt_dollar_Dict sync.Once

func Get_Control_Alt_Alt_dollar_Dict() gopurs_runtime.Value {
	once_Control_Alt_Alt_dollar_Dict.Do(func() {
		cache_Control_Alt_Alt_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Call_Control_Alt_Alt_dollar_Dict(func() struct {
				Functor0 gopurs_runtime.Value
				alt      gopurs_runtime.Value
				/*gopurs:record-v1:2:00460075006e00630074006f00720030,0061006c0074*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Functor0 gopurs_runtime.Value
					alt      gopurs_runtime.Value
					/*gopurs:record-v1:2:00460075006e00630074006f00720030,0061006c0074*/
				}{}
				clone.Functor0 = gopurs_runtime.RecordGet(orig, "Functor0")
				clone.alt = gopurs_runtime.RecordGet(orig, "alt")
				return clone
			}()))}
		})
	})
	return cache_Control_Alt_Alt_dollar_Dict
}

var cache_Control_Alt_altArray gopurs_runtime.Value
var once_Control_Alt_altArray sync.Once

func Get_Control_Alt_altArray() gopurs_runtime.Value {
	once_Control_Alt_altArray.Do(func() {
		cache_Control_Alt_altArray = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}
		}), Get_Data_Semigroup_concatArray()}))}
	})
	return cache_Control_Alt_altArray
}

var cache_Control_Alt_alt gopurs_runtime.Value
var once_Control_Alt_alt sync.Once

func Get_Control_Alt_alt() gopurs_runtime.Value {
	once_Control_Alt_alt.Do(func() {
		cache_Control_Alt_alt = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Alt_alt(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Alt_alt
}

type Constructor_Control_Alt_Alt[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f0041006c0074:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f0041006c0074005f0041006c0074
func init() {
	gopurs_runtime.StructGetters[4060500237] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Alt_Alt[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Functor0":
			return gopurs_runtime.Box(c.V0)
		case "alt":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Alt_Alt: " + key)
		}
	}
}

func Call_Control_Alt_Alt_dollar_Dict(x_0_loop struct {
	Functor0 gopurs_runtime.Value
	alt      gopurs_runtime.Value
	/*gopurs:record-v1:2:00460075006e00630074006f00720030,0061006c0074*/
}) *Constructor_Control_Alt_Alt[gopurs_runtime.Value] {
	var x_0 struct {
		Functor0 gopurs_runtime.Value
		alt      gopurs_runtime.Value
		/*gopurs:record-v1:2:00460075006e00630074006f00720030,0061006c0074*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Functor0", "alt", orig.Functor0, orig.alt)
	}())
}

func Call_Control_Alt_alt(dict_0_loop *Constructor_Control_Alt_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Alt_Alt[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}
