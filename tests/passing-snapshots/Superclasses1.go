package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Su_dollar_Dict gopurs_runtime.Value
var once_Main_Su_dollar_Dict sync.Once

func Get_Main_Su_dollar_Dict() gopurs_runtime.Value {
	once_Main_Su_dollar_Dict.Do(func() {
		cache_Main_Su_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Su_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Su_dollar_Dict
}

var cache_Main_Cl_dollar_Dict gopurs_runtime.Value
var once_Main_Cl_dollar_Dict sync.Once

func Get_Main_Cl_dollar_Dict() gopurs_runtime.Value {
	once_Main_Cl_dollar_Dict.Do(func() {
		cache_Main_Cl_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Cl_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_Cl_dollar_Dict
}

var cache_Main_suNumber gopurs_runtime.Value
var once_Main_suNumber sync.Once

func Get_Main_suNumber() gopurs_runtime.Value {
	once_Main_suNumber.Do(func() {
		cache_Main_suNumber = gopurs_runtime.Value{Type: 9, IntVal: 999349368, UnsafePtr: unsafe.Pointer((&Constructor_Main_Su[float64]{1, gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((n_0.FloatVal()) + (1.0))
		})}))}
	})
	return cache_Main_suNumber
}

var cache_Main_su gopurs_runtime.Value
var once_Main_su sync.Once

func Get_Main_su() gopurs_runtime.Value {
	once_Main_su.Do(func() {
		cache_Main_su = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_su(gopurs_runtime.CoerceToStruct[Constructor_Main_Su[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_su
}

var cache_Main_su__2782801075 gopurs_runtime.Value
var once_Main_su__2782801075 sync.Once

func Get_Main_su__2782801075() gopurs_runtime.Value {
	once_Main_su__2782801075.Do(func() {
		cache_Main_su__2782801075 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_su__2782801075(gopurs_runtime.CoerceToStruct[Constructor_Main_Su[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_su__2782801075
}

var cache_Main_clNumber gopurs_runtime.Value
var once_Main_clNumber sync.Once

func Get_Main_clNumber() gopurs_runtime.Value {
	once_Main_clNumber.Do(func() {
		cache_Main_clNumber = gopurs_runtime.Value{Type: 9, IntVal: 2792887505, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cl[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 999349368, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Su[float64]](Get_Main_suNumber()))}
		}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Float((n_0.FloatVal()) + (m_1.FloatVal()))
			})
		})}))}
	})
	return cache_Main_clNumber
}

var cache_Main_cl gopurs_runtime.Value
var once_Main_cl sync.Once

func Get_Main_cl() gopurs_runtime.Value {
	once_Main_cl.Do(func() {
		cache_Main_cl = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cl(gopurs_runtime.CoerceToStruct[Constructor_Main_Cl[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_cl
}

var cache_Main_cl__826691443 gopurs_runtime.Value
var once_Main_cl__826691443 sync.Once

func Get_Main_cl__826691443() gopurs_runtime.Value {
	once_Main_cl__826691443.Do(func() {
		cache_Main_cl__826691443 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_cl__826691443(gopurs_runtime.CoerceToStruct[Constructor_Main_Cl[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_cl__826691443
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func(func(dictCl_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(gopurs_runtime.CoerceToStruct[Constructor_Main_Cl[gopurs_runtime.Value]](dictCl_0_box))
		})
	})
	return cache_Main_test
}

var cache_Main_test__2037638682 gopurs_runtime.Value
var once_Main_test__2037638682 sync.Once

func Get_Main_test__2037638682() gopurs_runtime.Value {
	once_Main_test__2037638682.Do(func() {
		cache_Main_test__2037638682 = gopurs_runtime.Func(func(dictCl_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test__2037638682(gopurs_runtime.CoerceToStruct[Constructor_Main_Cl[gopurs_runtime.Value]](dictCl_0_box))
		})
	})
	return cache_Main_test__2037638682
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float(gopurs_runtime.Float(21.0).FloatVal())).StrVal()))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Su[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[999349368] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Su[any])(ptr)
		_ = c
		switch key {
		case "su":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Su: " + key)
		}
	}
}

type Constructor_Main_Cl[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2792887505] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Cl[any])(ptr)
		_ = c
		switch key {
		case "Su0":
			return gopurs_runtime.Box(c.V0)
		case "cl":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_Cl: " + key)
		}
	}
}

func Call_Main_Su_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Cl_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_su(dict_0_loop *Constructor_Main_Su[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Su[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_su__2782801075(dict_0_loop *Constructor_Main_Su[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Su[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_cl(dict_0_loop *Constructor_Main_Cl[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Cl[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_cl__826691443(dict_0_loop *Constructor_Main_Cl[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Cl[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V1)
}

func Call_Main_test(dictCl_0_loop *Constructor_Main_Cl[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictCl_0 *Constructor_Main_Cl[gopurs_runtime.Value] = dictCl_0_loop
	_ = dictCl_0
	// TAST (Let): Su0_1_0 shape=App(Other) bindingType=(ADT ["Main","Su"] [(TypeVar a)])
	Su0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Su[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictCl_0.V0), gopurs_runtime.Value{}))
	_ = Su0_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Box(Su0_1_0.V0), gopurs_runtime.Apply2(gopurs_runtime.Box(dictCl_0.V1), a_2, a_2))
	})
}

func Call_Main_test__2037638682(dictCl_0_loop *Constructor_Main_Cl[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictCl_0 *Constructor_Main_Cl[gopurs_runtime.Value] = dictCl_0_loop
	_ = dictCl_0
	// TAST (Let): Su0_1_0 shape=App(Other) bindingType=(ADT ["Main","Su"] [(TypeVar a)])
	Su0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Su[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictCl_0.V0), gopurs_runtime.Value{}))
	_ = Su0_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Box(Su0_1_0.V0), gopurs_runtime.Apply2(gopurs_runtime.Box(dictCl_0.V1), a_2, a_2))
	})
}
