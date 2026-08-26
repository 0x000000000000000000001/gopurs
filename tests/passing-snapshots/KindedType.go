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

var cache_Main_Clazz_dollar_Dict gopurs_runtime.Value
var once_Main_Clazz_dollar_Dict sync.Once

func Get_Main_Clazz_dollar_Dict() gopurs_runtime.Value {
	once_Main_Clazz_dollar_Dict.Do(func() {
		cache_Main_Clazz_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Clazz_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Clazz_dollar_Dict
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.RecordDict1("a", gopurs_runtime.Str("test"))
	})
	return cache_Main_test5
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = func() gopurs_runtime.Value {
			arr := []string{"test"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()
	})
	return cache_Main_test4
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test3
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = func() gopurs_runtime.Value {
			arr := []string{"test"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()
	})
	return cache_Main_test1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_f gopurs_runtime.Value
var once_Main_f sync.Once

func Get_Main_f() gopurs_runtime.Value {
	once_Main_f.Do(func() {
		cache_Main_f = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_f(s_0_box.StrVal()))
		})
	})
	return cache_Main_f
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Str("test")
	})
	return cache_Main_test2
}

var cache_Main_def gopurs_runtime.Value
var once_Main_def sync.Once

func Get_Main_def() gopurs_runtime.Value {
	once_Main_def.Do(func() {
		cache_Main_def = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_def(dict_0_box)
		})
	})
	return cache_Main_def
}

var cache_Main_clazzString gopurs_runtime.Value
var once_Main_clazzString sync.Once

func Get_Main_clazzString() gopurs_runtime.Value {
	once_Main_clazzString.Do(func() {
		cache_Main_clazzString = gopurs_runtime.Value{Type: 9, IntVal: 3678469904, UnsafePtr: unsafe.Pointer((&Constructor_Main_Clazz[string]{1, gopurs_runtime.Str("test")}))}
	})
	return cache_Main_clazzString
}

type Constructor_Main_Proxy[T_f any] struct {
	Rc uint32
}

type Constructor_Main_Clazz[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[3678469904] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Clazz[any])(ptr)
		_ = c
		switch key {
		case "def":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Clazz: " + key)
		}
	}
}

func Call_Main_Clazz_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_f(s_0_loop string) string {
	var s_0 string = s_0_loop
	_ = s_0
	return s_0
}

func Call_Main_def(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "def")
}
