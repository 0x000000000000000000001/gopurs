package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Test2_dollar_Dict gopurs_runtime.Value
var once_Main_Test2_dollar_Dict sync.Once

func Get_Main_Test2_dollar_Dict() gopurs_runtime.Value {
	once_Main_Test2_dollar_Dict.Do(func() {
		cache_Main_Test2_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test2_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Test2_dollar_Dict
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Main_test2
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = func() gopurs_runtime.Value {
			orig := func() *struct {
				attr gopurs_runtime.Value
			} {
				orig := gopurs_runtime.RecordDict1("attr", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(0)
				}))
				_ = orig
				clone := struct {
					attr gopurs_runtime.Value
				}{}
				clone.attr = gopurs_runtime.RecordGet(orig, "attr")
				return &clone
			}()
			_ = orig
			return gopurs_runtime.RecordDict([]string{"attr"}, []gopurs_runtime.Value{orig.attr})
		}()
	})
	return cache_Main_test1
}

var cache_Main_test0 gopurs_runtime.Value
var once_Main_test0 sync.Once

func Get_Main_test0() gopurs_runtime.Value {
	once_Main_test0.Do(func() {
		cache_Main_test0 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_test0(v_0_box))
		})
	})
	return cache_Main_test0
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
		cache_Main_f = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_f(gopurs_runtime.CoerceToStruct[Constructor_Main_Test2](dict_0_box))
		})
	})
	return cache_Main_f
}

type Constructor_Main_Test2 struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2375191994] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Test2)(ptr)
		_ = c
		switch key {
		case "f":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Test2: " + key)
		}
	}
}

func Call_Main_Test2_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_test0(v_0_loop gopurs_runtime.Value) int64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return 0
}

func Call_Main_f(dict_0_loop *Constructor_Main_Test2) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Test2 = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
