package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Id gopurs_runtime.Value
var once_Main_Id sync.Once

func Get_Main_Id() gopurs_runtime.Value {
	once_Main_Id.Do(func() {
		cache_Main_Id = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Id
}

var cache_Main_E_dollar_Dict gopurs_runtime.Value
var once_Main_E_dollar_Dict sync.Once

func Get_Main_E_dollar_Dict() gopurs_runtime.Value {
	once_Main_E_dollar_Dict.Do(func() {
		cache_Main_E_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_E_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_E_dollar_Dict
}

var cache_Main_runId gopurs_runtime.Value
var once_Main_runId sync.Once

func Get_Main_runId() gopurs_runtime.Value {
	once_Main_runId.Do(func() {
		cache_Main_runId = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_runId(v_0_box)
		})
	})
	return cache_Main_runId
}

var cache_Main_num gopurs_runtime.Value
var once_Main_num sync.Once

func Get_Main_num() gopurs_runtime.Value {
	once_Main_num.Do(func() {
		cache_Main_num = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_num(gopurs_runtime.CoerceToStruct[Constructor_Main_E[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_num
}

var cache_Main_num__3369776827 gopurs_runtime.Value
var once_Main_num__3369776827 sync.Once

func Get_Main_num__3369776827() gopurs_runtime.Value {
	once_Main_num__3369776827.Do(func() {
		cache_Main_num__3369776827 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_num__3369776827(gopurs_runtime.CoerceToStruct[Constructor_Main_E[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_num__3369776827
}

var cache_Main_exprId gopurs_runtime.Value
var once_Main_exprId sync.Once

func Get_Main_exprId() gopurs_runtime.Value {
	once_Main_exprId.Do(func() {
		cache_Main_exprId = gopurs_runtime.Value{Type: 9, IntVal: 1955825563, UnsafePtr: unsafe.Pointer((&Constructor_Main_E[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
			})
		}), Get_Main_Id()}))}
	})
	return cache_Main_exprId
}

var cache_Main_add gopurs_runtime.Value
var once_Main_add sync.Once

func Get_Main_add() gopurs_runtime.Value {
	once_Main_add.Do(func() {
		cache_Main_add = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_add(gopurs_runtime.CoerceToStruct[Constructor_Main_E[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_add
}

var cache_Main_add__1253445479 gopurs_runtime.Value
var once_Main_add__1253445479 sync.Once

func Get_Main_add__1253445479() gopurs_runtime.Value {
	once_Main_add__1253445479.Do(func() {
		cache_Main_add__1253445479 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_add__1253445479(gopurs_runtime.CoerceToStruct[Constructor_Main_E[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_add__1253445479
}

var cache_Main_three gopurs_runtime.Value
var once_Main_three sync.Once

func Get_Main_three() gopurs_runtime.Value {
	once_Main_three.Do(func() {
		cache_Main_three = gopurs_runtime.Func(func(dictE_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_three(dictE_0_box)
		})
	})
	return cache_Main_three
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(gopurs_runtime.Float(3.0).FloatVal())).StrVal()))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Id[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_E[T_e any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1955825563] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_E[any])(ptr)
		_ = c
		switch key {
		case "add":
			return gopurs_runtime.Box(c.V0)
		case "num":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_E: " + key)
		}
	}
}

func Call_Main_E_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_runId(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Main_num(dict_0_loop *Constructor_Main_E[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_E[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_num__3369776827(dict_0_loop *Constructor_Main_E[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_E[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_add(dict_0_loop *Constructor_Main_E[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_E[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_add__1253445479(dict_0_loop *Constructor_Main_E[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_E[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_three(dictE_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictE_0 gopurs_runtime.Value = dictE_0_loop
	_ = dictE_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictE_0, "add"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictE_0, "num"), gopurs_runtime.Float(1.0)), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictE_0, "num"), gopurs_runtime.Float(2.0)))
}
