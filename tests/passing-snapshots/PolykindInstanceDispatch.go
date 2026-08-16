package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_ShowP_dollar_Dict gopurs_runtime.Value
var once_Main_ShowP_dollar_Dict sync.Once

func Get_Main_ShowP_dollar_Dict() gopurs_runtime.Value {
	once_Main_ShowP_dollar_Dict.Do(func() {
		cache_Main_ShowP_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ShowP_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_ShowP_dollar_Dict
}

var cache_Main_Proxy gopurs_runtime.Value
var once_Main_Proxy sync.Once

func Get_Main_Proxy() gopurs_runtime.Value {
	once_Main_Proxy.Do(func() {
		cache_Main_Proxy = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_Proxy
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Value{Type: 9, IntVal: 1349270669, UnsafePtr: unsafe.Pointer((&Constructor_Main_ShowP{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("Symbol")
		})}))}
	})
	return cache_Main_test2
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Value{Type: 9, IntVal: 1349270669, UnsafePtr: unsafe.Pointer((&Constructor_Main_ShowP{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("Type")
		})}))}
	})
	return cache_Main_test1
}

var cache_Main_showP gopurs_runtime.Value
var once_Main_showP sync.Once

func Get_Main_showP() gopurs_runtime.Value {
	once_Main_showP.Do(func() {
		cache_Main_showP = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_showP(gopurs_runtime.CoerceToStruct[Constructor_Main_ShowP](dict_0_box))
		})
	})
	return cache_Main_showP
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_showP__1027712357 gopurs_runtime.Value
var once_Main_showP__1027712357 sync.Once

func Get_Main_showP__1027712357() gopurs_runtime.Value {
	once_Main_showP__1027712357.Do(func() {
		cache_Main_showP__1027712357 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_showP__1027712357(gopurs_runtime.CoerceToStruct[Constructor_Main_ShowP](dict_0_box))
		})
	})
	return cache_Main_showP__1027712357
}

var cache_Main_test1__1863477487 gopurs_runtime.Value
var once_Main_test1__1863477487 sync.Once

func Get_Main_test1__1863477487() gopurs_runtime.Value {
	once_Main_test1__1863477487.Do(func() {
		cache_Main_test1__1863477487 = gopurs_runtime.Value{Type: 9, IntVal: 1349270669, UnsafePtr: unsafe.Pointer((&Constructor_Main_ShowP{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("Type")
		})}))}
	})
	return cache_Main_test1__1863477487
}

var cache_Main_test2__61460478 gopurs_runtime.Value
var once_Main_test2__61460478 sync.Once

func Get_Main_test2__61460478() gopurs_runtime.Value {
	once_Main_test2__61460478.Do(func() {
		cache_Main_test2__61460478 = gopurs_runtime.Value{Type: 9, IntVal: 1349270669, UnsafePtr: unsafe.Pointer((&Constructor_Main_ShowP{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("Symbol")
		})}))}
	})
	return cache_Main_test2__61460478
}

type Constructor_Main_Proxy struct {
	Rc uint32
}

type Constructor_Main_ShowP struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1349270669] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ShowP)(ptr)
		_ = c
		switch key {
		case "showP":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_ShowP: " + key)
		}
	}
}

func Call_Main_ShowP_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_showP(dict_0_loop *Constructor_Main_ShowP) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_ShowP = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_showP__1027712357(dict_0_loop *Constructor_Main_ShowP) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_ShowP = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
