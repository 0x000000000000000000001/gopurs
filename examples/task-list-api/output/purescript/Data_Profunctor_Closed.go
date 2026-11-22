package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_Closed_Closed_dollar_Dict gopurs_runtime.Value
var once_Data_Profunctor_Closed_Closed_dollar_Dict sync.Once

func Get_Data_Profunctor_Closed_Closed_dollar_Dict() gopurs_runtime.Value {
	once_Data_Profunctor_Closed_Closed_dollar_Dict.Do(func() {
		cache_Data_Profunctor_Closed_Closed_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 768764671, UnsafePtr: unsafe.Pointer(Call_Data_Profunctor_Closed_Closed_dollar_Dict(func() struct {
				Profunctor0 gopurs_runtime.Value
				closed      gopurs_runtime.Value
				/*gopurs:record-v1:2:00500072006f00660075006e00630074006f00720030,0063006c006f007300650064*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Profunctor0 gopurs_runtime.Value
					closed      gopurs_runtime.Value
					/*gopurs:record-v1:2:00500072006f00660075006e00630074006f00720030,0063006c006f007300650064*/
				}{}
				clone.Profunctor0 = gopurs_runtime.RecordGet(orig, "Profunctor0")
				clone.closed = gopurs_runtime.RecordGet(orig, "closed")
				return clone
			}()))}
		})
	})
	return cache_Data_Profunctor_Closed_Closed_dollar_Dict
}

var cache_Data_Profunctor_Closed_closedFunction gopurs_runtime.Value
var once_Data_Profunctor_Closed_closedFunction sync.Once

func Get_Data_Profunctor_Closed_closedFunction() gopurs_runtime.Value {
	once_Data_Profunctor_Closed_closedFunction.Do(func() {
		cache_Data_Profunctor_Closed_closedFunction = gopurs_runtime.Value{Type: 9, IntVal: 768764671, UnsafePtr: unsafe.Pointer((&Constructor_Data_Profunctor_Closed_Closed[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Profunctor[gopurs_runtime.Value]](Get_Data_Profunctor_profunctorFn()))}
		}), gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()).V0}))}
	})
	return cache_Data_Profunctor_Closed_closedFunction
}

var cache_Data_Profunctor_Closed_closed gopurs_runtime.Value
var once_Data_Profunctor_Closed_closed sync.Once

func Get_Data_Profunctor_Closed_closed() gopurs_runtime.Value {
	once_Data_Profunctor_Closed_closed.Do(func() {
		cache_Data_Profunctor_Closed_closed = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Closed_closed(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Closed_Closed[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Profunctor_Closed_closed
}

type Constructor_Data_Profunctor_Closed_Closed[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00500072006f00660075006e00630074006f0072005f0043006c006f007300650064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00500072006f00660075006e00630074006f0072005f0043006c006f007300650064005f0043006c006f007300650064
func init() {
	gopurs_runtime.StructGetters[768764671] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Profunctor_Closed_Closed[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Profunctor0":
			return gopurs_runtime.Box(c.V0)
		case "closed":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Profunctor_Closed_Closed: " + key)
		}
	}
}

func Call_Data_Profunctor_Closed_Closed_dollar_Dict(x_0_loop struct {
	Profunctor0 gopurs_runtime.Value
	closed      gopurs_runtime.Value
	/*gopurs:record-v1:2:00500072006f00660075006e00630074006f00720030,0063006c006f007300650064*/
}) *Constructor_Data_Profunctor_Closed_Closed[gopurs_runtime.Value] {
	var x_0 struct {
		Profunctor0 gopurs_runtime.Value
		closed      gopurs_runtime.Value
		/*gopurs:record-v1:2:00500072006f00660075006e00630074006f00720030,0063006c006f007300650064*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Closed_Closed[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Profunctor0", "closed", orig.Profunctor0, orig.closed)
	}())
}

func Call_Data_Profunctor_Closed_closed(dict_0_loop *Constructor_Data_Profunctor_Closed_Closed[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Profunctor_Closed_Closed[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}
