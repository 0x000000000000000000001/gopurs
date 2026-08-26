package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Convert_dollar_Dict gopurs_runtime.Value
var once_Main_Convert_dollar_Dict sync.Once

func Get_Main_Convert_dollar_Dict() gopurs_runtime.Value {
	once_Main_Convert_dollar_Dict.Do(func() {
		cache_Main_Convert_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Convert_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Convert_dollar_Dict
}

var cache_Main_convertSB gopurs_runtime.Value
var once_Main_convertSB sync.Once

func Get_Main_convertSB() gopurs_runtime.Value {
	once_Main_convertSB.Do(func() {
		cache_Main_convertSB = gopurs_runtime.Value{Type: 9, IntVal: 3639075177, UnsafePtr: unsafe.Pointer((&Constructor_Main_Convert[int64, string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 string
			{
				if (v_0.IntVal) == (0) {
					__t0 = "Nope"
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = "Done"
			}
		end_branch_0:
			return gopurs_runtime.Str(__t0)
		})}))}
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

var cache_Main_convert__769515702 gopurs_runtime.Value
var once_Main_convert__769515702 sync.Once

func Get_Main_convert__769515702() gopurs_runtime.Value {
	once_Main_convert__769515702.Do(func() {
		cache_Main_convert__769515702 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_convert__769515702(gopurs_runtime.CoerceToStruct[Constructor_Main_Convert[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_convert__769515702
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

func init() {
	gopurs_runtime.StructGetters[3639075177] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Convert[any, any])(ptr)
		_ = c
		switch key {
		case "convert":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Convert: " + key)
		}
	}
}

func Call_Main_Convert_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_convert(dict_0_loop *Constructor_Main_Convert[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Convert[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_convert__769515702(dict_0_loop *Constructor_Main_Convert[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Convert[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
