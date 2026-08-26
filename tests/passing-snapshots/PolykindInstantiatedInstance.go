package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Proxy gopurs_runtime.Value
var once_Main_Proxy sync.Once

func Get_Main_Proxy() gopurs_runtime.Value {
	once_Main_Proxy.Do(func() {
		cache_Main_Proxy = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_Proxy
}

var cache_Main_F_dollar_Dict gopurs_runtime.Value
var once_Main_F_dollar_Dict sync.Once

func Get_Main_F_dollar_Dict() gopurs_runtime.Value {
	once_Main_F_dollar_Dict.Do(func() {
		cache_Main_F_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_F_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_F_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_fProxy gopurs_runtime.Value
var once_Main_fProxy sync.Once

func Get_Main_fProxy() gopurs_runtime.Value {
	once_Main_fProxy.Do(func() {
		cache_Main_fProxy = gopurs_runtime.Value{Type: 9, IntVal: 1987449688, UnsafePtr: unsafe.Pointer((&Constructor_Main_F[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
			})
		})}))}
	})
	return cache_Main_fProxy
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f(gopurs_runtime.CoerceToStruct[Constructor_Main_F[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_f
}

var cache_Main_f__3489368022 gopurs_runtime.Value
var once_Main_f__3489368022 sync.Once

func Get_Main_f__3489368022() gopurs_runtime.Value {
	once_Main_f__3489368022.Do(func() {
		cache_Main_f__3489368022 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f__3489368022(gopurs_runtime.CoerceToStruct[Constructor_Main_F[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_f__3489368022
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test1
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test2
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test3
}

type Constructor_Main_Proxy[T_a any] struct {
	Rc uint32
}

type Constructor_Main_F[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1987449688] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_F[any])(ptr)
		_ = c
		switch key {
		case "f":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_F: " + key)
		}
	}
}

func Call_Main_F_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_f(dict_0_loop *Constructor_Main_F[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_F[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_f__3489368022(dict_0_loop *Constructor_Main_F[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_F[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
