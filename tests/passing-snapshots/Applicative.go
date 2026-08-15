package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Nothing gopurs_runtime.Value
var once_Main_Nothing sync.Once

func Get_Main_Nothing() gopurs_runtime.Value {
	once_Main_Nothing.Do(func() {
		cache_Main_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer((*Constructor_Main_Just)(nil))}
	})
	return cache_Main_Nothing
}

var cache_Main_Just gopurs_runtime.Value
var once_Main_Just sync.Once

func Get_Main_Just() gopurs_runtime.Value {
	once_Main_Just.Do(func() {
		cache_Main_Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(&Constructor_Main_Just{1, value0})}
		})
	})
	return cache_Main_Just
}

var cache_Main_Applicative_dollar_Dict gopurs_runtime.Value
var once_Main_Applicative_dollar_Dict sync.Once

func Get_Main_Applicative_dollar_Dict() gopurs_runtime.Value {
	once_Main_Applicative_dollar_Dict.Do(func() {
		cache_Main_Applicative_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Applicative_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Applicative_dollar_Dict
}

var cache_Main_pure gopurs_runtime.Value
var once_Main_pure sync.Once

func Get_Main_pure() gopurs_runtime.Value {
	once_Main_pure.Do(func() {
		cache_Main_pure = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_pure(gopurs_runtime.CoerceToStruct[Constructor_Main_Applicative](dict_0_box))
		})
	})
	return cache_Main_pure
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_apply gopurs_runtime.Value
var once_Main_apply sync.Once

func Get_Main_apply() gopurs_runtime.Value {
	once_Main_apply.Do(func() {
		cache_Main_apply = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_apply(gopurs_runtime.CoerceToStruct[Constructor_Main_Applicative](dict_0_box))
		})
	})
	return cache_Main_apply
}

var cache_Main_applicativeMaybe gopurs_runtime.Value
var once_Main_applicativeMaybe sync.Once

func Get_Main_applicativeMaybe() gopurs_runtime.Value {
	once_Main_applicativeMaybe.Do(func() {
		cache_Main_applicativeMaybe = gopurs_runtime.Value{Type: 9, IntVal: 4228518006, UnsafePtr: unsafe.Pointer(&Constructor_Main_Applicative{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 *Constructor_Main_Just
				{
					if (v_0.Type == 9 && v_0.IntVal == 3271839782 && v_0.UnsafePtr != nil) && (v1_1.Type == 9 && v1_1.IntVal == 3271839782 && v1_1.UnsafePtr != nil) {
						__t0 = &Constructor_Main_Just{1, gopurs_runtime.Apply((*Constructor_Main_Just)(v_0.UnsafePtr).V0, (*Constructor_Main_Just)(v1_1.UnsafePtr).V0)}
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = (*Constructor_Main_Just)(nil)
				}
			end_branch_0:
				return gopurs_runtime.Value{Type: 9, IntVal: 3271839782, UnsafePtr: unsafe.Pointer(__t0)}
			})
		}), Get_Main_Just()})}
	})
	return cache_Main_applicativeMaybe
}

type Constructor_Main_Nothing struct {
	Rc uint32
}

type Constructor_Main_Just struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Applicative struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4228518006] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Applicative)(ptr)
		_ = c
		switch key {
		case "apply":
			return gopurs_runtime.Box(c.V0)
		case "pure":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Applicative: " + key)
		}
	}
}

func Call_Main_Applicative_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_pure(dict_0_loop *Constructor_Main_Applicative) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Applicative = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_apply(dict_0_loop *Constructor_Main_Applicative) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Applicative = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
