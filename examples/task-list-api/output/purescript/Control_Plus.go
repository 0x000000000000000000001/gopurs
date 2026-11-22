package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Plus_Plus_dollar_Dict gopurs_runtime.Value
var once_Control_Plus_Plus_dollar_Dict sync.Once

func Get_Control_Plus_Plus_dollar_Dict() gopurs_runtime.Value {
	once_Control_Plus_Plus_dollar_Dict.Do(func() {
		cache_Control_Plus_Plus_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Call_Control_Plus_Plus_dollar_Dict(func() struct {
				Alt0  gopurs_runtime.Value
				empty gopurs_runtime.Value
				/*gopurs:record-v1:2:0041006c00740030,0065006d007000740079*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Alt0  gopurs_runtime.Value
					empty gopurs_runtime.Value
					/*gopurs:record-v1:2:0041006c00740030,0065006d007000740079*/
				}{}
				clone.Alt0 = gopurs_runtime.RecordGet(orig, "Alt0")
				clone.empty = gopurs_runtime.RecordGet(orig, "empty")
				return clone
			}()))}
		})
	})
	return cache_Control_Plus_Plus_dollar_Dict
}

var cache_Control_Plus_plusArray gopurs_runtime.Value
var once_Control_Plus_plusArray sync.Once

func Get_Control_Plus_plusArray() gopurs_runtime.Value {
	once_Control_Plus_plusArray.Do(func() {
		cache_Control_Plus_plusArray = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Get_Control_Alt_altArray()))}
		}), gopurs_runtime.Array([]gopurs_runtime.Value{})}))}
	})
	return cache_Control_Plus_plusArray
}

var cache_Control_Plus_empty gopurs_runtime.Value
var once_Control_Plus_empty sync.Once

func Get_Control_Plus_empty() gopurs_runtime.Value {
	once_Control_Plus_empty.Do(func() {
		cache_Control_Plus_empty = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Plus_empty(dict_0_box)
		})
	})
	return cache_Control_Plus_empty
}

var cache_Control_Plus_empty__3264985657 gopurs_runtime.Value
var once_Control_Plus_empty__3264985657 sync.Once

func Get_Control_Plus_empty__3264985657() gopurs_runtime.Value {
	once_Control_Plus_empty__3264985657.Do(func() {
		cache_Control_Plus_empty__3264985657 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Rebox_Control_Plus_3706288089_2173899317(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_List_Types_plusList())).V1))}
	})
	return cache_Control_Plus_empty__3264985657
}

type Constructor_Control_Plus_Plus[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f0050006c00750073:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f0050006c00750073005f0050006c00750073
func init() {
	gopurs_runtime.StructGetters[3709470893] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Plus_Plus[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Alt0":
			return gopurs_runtime.Box(c.V0)
		case "empty":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Plus_Plus: " + key)
		}
	}
}

func Call_Control_Plus_Plus_dollar_Dict(x_0_loop struct {
	Alt0  gopurs_runtime.Value
	empty gopurs_runtime.Value
	/*gopurs:record-v1:2:0041006c00740030,0065006d007000740079*/
}) *Constructor_Control_Plus_Plus[gopurs_runtime.Value] {
	var x_0 struct {
		Alt0  gopurs_runtime.Value
		empty gopurs_runtime.Value
		/*gopurs:record-v1:2:0041006c00740030,0065006d007000740079*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Alt0", "empty", orig.Alt0, orig.empty)
	}())
}

func Call_Control_Plus_empty(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "empty")
}

func Rebox_Control_Plus_3706288089_2173899317(in *Constructor_Control_Plus_Plus[gopurs_runtime.Value]) *Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}
