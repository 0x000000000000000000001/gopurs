package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_Cochoice_Cochoice_dollar_Dict gopurs_runtime.Value
var once_Data_Profunctor_Cochoice_Cochoice_dollar_Dict sync.Once

func Get_Data_Profunctor_Cochoice_Cochoice_dollar_Dict() gopurs_runtime.Value {
	once_Data_Profunctor_Cochoice_Cochoice_dollar_Dict.Do(func() {
		cache_Data_Profunctor_Cochoice_Cochoice_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1827340575, UnsafePtr: unsafe.Pointer(Call_Data_Profunctor_Cochoice_Cochoice_dollar_Dict(func() struct {
				Profunctor0 gopurs_runtime.Value
				unleft      gopurs_runtime.Value
				unright     gopurs_runtime.Value
				/*gopurs:record-v1:3:00500072006f00660075006e00630074006f00720030,0075006e006c006500660074,0075006e00720069006700680074*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Profunctor0 gopurs_runtime.Value
					unleft      gopurs_runtime.Value
					unright     gopurs_runtime.Value
					/*gopurs:record-v1:3:00500072006f00660075006e00630074006f00720030,0075006e006c006500660074,0075006e00720069006700680074*/
				}{}
				clone.Profunctor0 = gopurs_runtime.RecordGet(orig, "Profunctor0")
				clone.unleft = gopurs_runtime.RecordGet(orig, "unleft")
				clone.unright = gopurs_runtime.RecordGet(orig, "unright")
				return clone
			}()))}
		})
	})
	return cache_Data_Profunctor_Cochoice_Cochoice_dollar_Dict
}

var cache_Data_Profunctor_Cochoice_unright gopurs_runtime.Value
var once_Data_Profunctor_Cochoice_unright sync.Once

func Get_Data_Profunctor_Cochoice_unright() gopurs_runtime.Value {
	once_Data_Profunctor_Cochoice_unright.Do(func() {
		cache_Data_Profunctor_Cochoice_unright = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Cochoice_unright(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Cochoice_Cochoice[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Profunctor_Cochoice_unright
}

var cache_Data_Profunctor_Cochoice_unleft gopurs_runtime.Value
var once_Data_Profunctor_Cochoice_unleft sync.Once

func Get_Data_Profunctor_Cochoice_unleft() gopurs_runtime.Value {
	once_Data_Profunctor_Cochoice_unleft.Do(func() {
		cache_Data_Profunctor_Cochoice_unleft = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Profunctor_Cochoice_unleft(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Cochoice_Cochoice[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Profunctor_Cochoice_unleft
}

type Constructor_Data_Profunctor_Cochoice_Cochoice[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00500072006f00660075006e00630074006f0072005f0043006f00630068006f006900630065:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00500072006f00660075006e00630074006f0072005f0043006f00630068006f006900630065005f0043006f00630068006f006900630065
func init() {
	gopurs_runtime.StructGetters[1827340575] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Profunctor_Cochoice_Cochoice[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Profunctor0":
			return gopurs_runtime.Box(c.V0)
		case "unleft":
			return gopurs_runtime.Box(c.V1)
		case "unright":
			return gopurs_runtime.Box(c.V2)
		default:
			panic("Key not found in dictionary Constructor_Data_Profunctor_Cochoice_Cochoice: " + key)
		}
	}
}

func Call_Data_Profunctor_Cochoice_Cochoice_dollar_Dict(x_0_loop struct {
	Profunctor0 gopurs_runtime.Value
	unleft      gopurs_runtime.Value
	unright     gopurs_runtime.Value
	/*gopurs:record-v1:3:00500072006f00660075006e00630074006f00720030,0075006e006c006500660074,0075006e00720069006700680074*/
}) *Constructor_Data_Profunctor_Cochoice_Cochoice[gopurs_runtime.Value] {
	var x_0 struct {
		Profunctor0 gopurs_runtime.Value
		unleft      gopurs_runtime.Value
		unright     gopurs_runtime.Value
		/*gopurs:record-v1:3:00500072006f00660075006e00630074006f00720030,0075006e006c006500660074,0075006e00720069006700680074*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Cochoice_Cochoice[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict3("Profunctor0", "unleft", "unright", orig.Profunctor0, orig.unleft, orig.unright)
	}())
}

func Call_Data_Profunctor_Cochoice_unright(dict_0_loop *Constructor_Data_Profunctor_Cochoice_Cochoice[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Profunctor_Cochoice_Cochoice[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Data_Profunctor_Cochoice_unleft(dict_0_loop *Constructor_Data_Profunctor_Cochoice_Cochoice[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Profunctor_Cochoice_Cochoice[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}
