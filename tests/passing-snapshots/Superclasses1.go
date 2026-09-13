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
			return gopurs_runtime.Value{Type: 9, IntVal: 999349368, UnsafePtr: unsafe.Pointer(Call_Main_Su_dollar_Dict(func() struct {
				su gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					su gopurs_runtime.Value
				}{}
				clone.su = gopurs_runtime.RecordGet(orig, "su")
				return clone
			}()))}
		})
	})
	return cache_Main_Su_dollar_Dict
}

var cache_Main_Su_dollar_Dict__229017040 gopurs_runtime.Value
var once_Main_Su_dollar_Dict__229017040 sync.Once

func Get_Main_Su_dollar_Dict__229017040() gopurs_runtime.Value {
	once_Main_Su_dollar_Dict__229017040.Do(func() {
		cache_Main_Su_dollar_Dict__229017040 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 999349368, UnsafePtr: unsafe.Pointer(Rebox_Main_673865076_3110518380(Call_Main_Su_dollar_Dict__229017040(func() struct {
				su gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					su gopurs_runtime.Value
				}{}
				clone.su = gopurs_runtime.RecordGet(orig, "su")
				return clone
			}())))}
		})
	})
	return cache_Main_Su_dollar_Dict__229017040
}

var cache_Main_Cl_dollar_Dict gopurs_runtime.Value
var once_Main_Cl_dollar_Dict sync.Once

func Get_Main_Cl_dollar_Dict() gopurs_runtime.Value {
	once_Main_Cl_dollar_Dict.Do(func() {
		cache_Main_Cl_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2792887505, UnsafePtr: unsafe.Pointer(Call_Main_Cl_dollar_Dict(func() struct {
				Su0 gopurs_runtime.Value
				cl  gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Su0 gopurs_runtime.Value
					cl  gopurs_runtime.Value
				}{}
				clone.Su0 = gopurs_runtime.RecordGet(orig, "Su0")
				clone.cl = gopurs_runtime.RecordGet(orig, "cl")
				return clone
			}()))}
		})
	})
	return cache_Main_Cl_dollar_Dict
}

var cache_Main_Cl_dollar_Dict__3543401801 gopurs_runtime.Value
var once_Main_Cl_dollar_Dict__3543401801 sync.Once

func Get_Main_Cl_dollar_Dict__3543401801() gopurs_runtime.Value {
	once_Main_Cl_dollar_Dict__3543401801.Do(func() {
		cache_Main_Cl_dollar_Dict__3543401801 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2792887505, UnsafePtr: unsafe.Pointer(Rebox_Main_203979197_636750501(Call_Main_Cl_dollar_Dict__3543401801(func() struct {
				Su0 gopurs_runtime.Value
				cl  gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Su0 gopurs_runtime.Value
					cl  gopurs_runtime.Value
				}{}
				clone.Su0 = gopurs_runtime.RecordGet(orig, "Su0")
				clone.cl = gopurs_runtime.RecordGet(orig, "cl")
				return clone
			}())))}
		})
	})
	return cache_Main_Cl_dollar_Dict__3543401801
}

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

var cache_Main_su__3096015727 gopurs_runtime.Value
var once_Main_su__3096015727 sync.Once

func Get_Main_su__3096015727() gopurs_runtime.Value {
	once_Main_su__3096015727.Do(func() {
		cache_Main_su__3096015727 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_su__3096015727(__eta_norm_0_0_box.FloatVal()))
		})
	})
	return cache_Main_su__3096015727
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

var cache_Main_cl__2575338094 gopurs_runtime.Value
var once_Main_cl__2575338094 sync.Once

func Get_Main_cl__2575338094() gopurs_runtime.Value {
	once_Main_cl__2575338094.Do(func() {
		cache_Main_cl__2575338094 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_cl__2575338094(__eta_norm_1_0_box.FloatVal(), __eta_norm_0_1_box.FloatVal()))
		})
	})
	return cache_Main_cl__2575338094
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

var cache_Main_test__1387011232 gopurs_runtime.Value
var once_Main_test__1387011232 sync.Once

func Get_Main_test__1387011232() gopurs_runtime.Value {
	once_Main_test__1387011232.Do(func() {
		cache_Main_test__1387011232 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Main_test__1387011232(__eta_norm_0_0_box.FloatVal()))
		})
	})
	return cache_Main_test__1387011232
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Effect_Console_logShow(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))), gopurs_runtime.Float(21.0)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

type Constructor_Main_Su[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

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

func Call_Main_Su_dollar_Dict(x_0_loop struct {
	su gopurs_runtime.Value
}) *Constructor_Main_Su[gopurs_runtime.Value] {
	var x_0 struct {
		su gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Su[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("su", orig.su)
	}())
}

func Call_Main_Su_dollar_Dict__229017040(x_0_loop struct {
	su gopurs_runtime.Value
}) *Constructor_Main_Su[float64] {
Su_dollar_Dict__229017040:
	for {
		if false {
			continue Su_dollar_Dict__229017040
		}
		var x_0 struct {
			su gopurs_runtime.Value
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_Su[float64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict1("su", orig.su)
		}())
	}
}

func Call_Main_Cl_dollar_Dict(x_0_loop struct {
	Su0 gopurs_runtime.Value
	cl  gopurs_runtime.Value
}) *Constructor_Main_Cl[gopurs_runtime.Value] {
	var x_0 struct {
		Su0 gopurs_runtime.Value
		cl  gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Cl[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Su0", "cl", orig.Su0, orig.cl)
	}())
}

func Call_Main_Cl_dollar_Dict__3543401801(x_0_loop struct {
	Su0 gopurs_runtime.Value
	cl  gopurs_runtime.Value
}) *Constructor_Main_Cl[float64] {
Cl_dollar_Dict__3543401801:
	for {
		if false {
			continue Cl_dollar_Dict__3543401801
		}
		var x_0 struct {
			Su0 gopurs_runtime.Value
			cl  gopurs_runtime.Value
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_Cl[float64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Su0", "cl", orig.Su0, orig.cl)
		}())
	}
}

func Call_Main_su(dict_0_loop *Constructor_Main_Su[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Su[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_su__3096015727(__eta_norm_0_0_loop float64) float64 {
su__3096015727:
	for {
		if false {
			continue su__3096015727
		}
		var __eta_norm_0_0 float64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return (__eta_norm_0_0) + (1.0)
	}
}

func Call_Main_cl(dict_0_loop *Constructor_Main_Cl[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Cl[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Main_cl__2575338094(__eta_norm_1_0_loop float64, __eta_norm_0_1_loop float64) float64 {
cl__2575338094:
	for {
		if false {
			continue cl__2575338094
		}
		var __eta_norm_1_0 float64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 float64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (__eta_norm_1_0) + (__eta_norm_0_1)
	}
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

func Call_Main_test__1387011232(__eta_norm_0_0_loop float64) float64 {
test__1387011232:
	for {
		if false {
			continue test__1387011232
		}
		var __eta_norm_0_0 float64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return ((__eta_norm_0_0) + (__eta_norm_0_0)) + (1.0)
	}
}

func Rebox_Main_1386611502_3263178038(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[float64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_203979197_636750501(in *Constructor_Main_Cl[float64]) *Constructor_Main_Cl[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Cl[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	return out
}

func Rebox_Main_3110518380_673865076(in *Constructor_Main_Su[gopurs_runtime.Value]) *Constructor_Main_Su[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Su[float64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_673865076_3110518380(in *Constructor_Main_Su[float64]) *Constructor_Main_Su[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Su[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
