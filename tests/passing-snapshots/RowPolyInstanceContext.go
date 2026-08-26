package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_T_dollar_Dict gopurs_runtime.Value
var once_Main_T_dollar_Dict sync.Once

func Get_Main_T_dollar_Dict() gopurs_runtime.Value {
	once_Main_T_dollar_Dict.Do(func() {
		cache_Main_T_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_T_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_T_dollar_Dict
}

var cache_Main_S gopurs_runtime.Value
var once_Main_S sync.Once

func Get_Main_S() gopurs_runtime.Value {
	once_Main_S.Do(func() {
		cache_Main_S = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_S
}

var cache_Main_state gopurs_runtime.Value
var once_Main_state sync.Once

func Get_Main_state() gopurs_runtime.Value {
	once_Main_state.Do(func() {
		cache_Main_state = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_state(gopurs_runtime.CoerceToStruct[Constructor_Main_T[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_state
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Func(func(dictT_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test2(dictT_0_box)
		})
	})
	return cache_Main_test2
}

var cache_Main_st gopurs_runtime.Value
var once_Main_st sync.Once

func Get_Main_st() gopurs_runtime.Value {
	once_Main_st.Do(func() {
		cache_Main_st = gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer((&Constructor_Main_T[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.RecordDict2("new", "ret", gopurs_runtime.Apply(f_0, s_1), Get_Data_Unit_unit())
			})
		})}))}
	})
	return cache_Main_st
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict2("new", "ret", gopurs_runtime.RecordUpdate1(s_0, "foo", gopurs_runtime.Str((gopurs_runtime.RecordGet(s_0, "foo").StrVal())+("!"))), Get_Data_Unit_unit())
		})
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

type Constructor_Main_S[T_s any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_T[T_s any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[990467018] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_T[any, any])(ptr)
		_ = c
		switch key {
		case "state":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_T: " + key)
		}
	}
}

func Call_Main_T_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_state(dict_0_loop *Constructor_Main_T[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_T[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_test2(dictT_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictT_0 gopurs_runtime.Value = dictT_0_loop
	_ = dictT_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictT_0, "state"), gopurs_runtime.Func(func(o_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.RecordUpdate1(o_1, "foo", gopurs_runtime.Str((gopurs_runtime.RecordGet(o_1, "foo").StrVal())+("!")))
	}))
}
