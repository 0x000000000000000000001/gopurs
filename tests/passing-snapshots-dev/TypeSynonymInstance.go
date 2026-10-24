package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_convertSB gopurs_runtime.Value
var once_Main_convertSB sync.Once

func Get_Main_convertSB() gopurs_runtime.Value {
	once_Main_convertSB.Do(func() {
		cache_Main_convertSB = gopurs_runtime.Value{Type: 9, IntVal: 3639075177, UnsafePtr: unsafe.Pointer(Rebox_Main_1758696223_1198875995((&Constructor_Main_Convert[int64, string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if (v_0.IntVal) == (int64(0)) {
					__t0 = gopurs_runtime.Str("Nope")
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = gopurs_runtime.Str("Done")
			}
		end_branch_0:
			return __t0
		})})))}
	})
	return cache_Main_convertSB
}

var cache_Main_convert gopurs_runtime.Value
var once_Main_convert sync.Once

func Get_Main_convert() gopurs_runtime.Value {
	once_Main_convert.Do(func() {
		cache_Main_convert = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_convert(gopurs_runtime.CoerceToStruct[Constructor_Main_Convert[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_convert
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Convert[T_a any, T_b any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0043006f006e0076006500720074
func init() {
	gopurs_runtime.StructGetters[3639075177] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Convert[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "convert":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Convert: " + key)
		}
	}
}

func Call_Main_convert(dict_0_loop *Constructor_Main_Convert[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Convert[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Main_1758696223_1198875995(in *Constructor_Main_Convert[int64, string]) *Constructor_Main_Convert[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Convert[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}
