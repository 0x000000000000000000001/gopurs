package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Cons_dollar_Dict gopurs_runtime.Value
var once_Main_Cons_dollar_Dict sync.Once

func Get_Main_Cons_dollar_Dict() gopurs_runtime.Value {
	once_Main_Cons_dollar_Dict.Do(func() {
		cache_Main_Cons_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Call_Main_Cons_dollar_Dict(func() struct {
				cons gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					cons gopurs_runtime.Value
				}{}
				clone.cons = gopurs_runtime.RecordGet(orig, "cons")
				return clone
			}()))}
		})
	})
	return cache_Main_Cons_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_cons gopurs_runtime.Value
var once_Main_cons sync.Once

func Get_Main_cons() gopurs_runtime.Value {
	once_Main_cons.Do(func() {
		cache_Main_cons = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cons(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_cons
}

type Constructor_Main_Cons[T_x any, T_xs any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[322902991] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Cons[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "cons":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Cons: " + key)
		}
	}
}

func Call_Main_Cons_dollar_Dict(x_0_loop struct {
	cons gopurs_runtime.Value
}) *Constructor_Main_Cons[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		cons gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("cons", orig.cons)
	}())
}

func Call_Main_cons(dict_0_loop *Constructor_Main_Cons[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Cons[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}
