package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_suNumber gopurs_runtime.Value
var once_Main_suNumber sync.Once

func Get_Main_suNumber() gopurs_runtime.Value {
	once_Main_suNumber.Do(func() {
		cache_Main_suNumber = gopurs_runtime.Value{Type: 9, IntVal: 999349368, UnsafePtr: unsafe.Pointer(Rebox_Main_673865076_3110518380((&Constructor_Main_Su[float64]{1, gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((n_0.FloatVal()) + (1.0))
		})})))}
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

var cache_Main_clNumber gopurs_runtime.Value
var once_Main_clNumber sync.Once

func Get_Main_clNumber() gopurs_runtime.Value {
	once_Main_clNumber.Do(func() {
		cache_Main_clNumber = gopurs_runtime.Value{Type: 9, IntVal: 2792887505, UnsafePtr: unsafe.Pointer(Rebox_Main_203979197_636750501((&Constructor_Main_Cl[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 999349368, UnsafePtr: unsafe.Pointer(Rebox_Main_673865076_3110518380(Rebox_Main_3110518380_673865076(gopurs_runtime.CoerceToStruct[Constructor_Main_Su[gopurs_runtime.Value]](Get_Main_suNumber()))))}
		}), gopurs_runtime.Func2(func(n_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((n_0.FloatVal()) + (m_1.FloatVal()))
		})})))}
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

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Effect_Console_logShow(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))), gopurs_runtime.Float(21.0))
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

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00530075
func init() {
	gopurs_runtime.StructGetters[999349368] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Su[gopurs_runtime.Value])(ptr)
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

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0043006c
func init() {
	gopurs_runtime.StructGetters[2792887505] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Cl[gopurs_runtime.Value])(ptr)
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

func Call_Main_su(dict_0_loop *Constructor_Main_Su[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Su[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_cl(dict_0_loop *Constructor_Main_Cl[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Cl[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Main_test(dictCl_0_loop *Constructor_Main_Cl[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictCl_0 *Constructor_Main_Cl[gopurs_runtime.Value] = dictCl_0_loop
	_ = dictCl_0
	// TAST (Let): Su0_1_0 shape=App(Other) bindingType=(ADT ["Main","Su"] [(TypeVar a$scope5)])
	Su0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Main_Su[gopurs_runtime.Value]](gopurs_runtime.Apply(dictCl_0.V0, gopurs_runtime.Value{}))
	_ = Su0_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Su0_1_0.V0, gopurs_runtime.Apply2(dictCl_0.V1, a_2, a_2))
	})
}

func Rebox_Main_1386611502_3263178038(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[float64])(unsafe.Pointer(in))
}

func Rebox_Main_203979197_636750501(in *Constructor_Main_Cl[float64]) *Constructor_Main_Cl[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Cl[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3110518380_673865076(in *Constructor_Main_Su[gopurs_runtime.Value]) *Constructor_Main_Su[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Su[float64])(unsafe.Pointer(in))
}

func Rebox_Main_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_673865076_3110518380(in *Constructor_Main_Su[float64]) *Constructor_Main_Su[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Main_Su[gopurs_runtime.Value])(unsafe.Pointer(in))
}
