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
			return gopurs_runtime.Value{Type: 9, IntVal: 1349270669, UnsafePtr: unsafe.Pointer(Call_Main_ShowP_dollar_Dict(func() struct {
				showP gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					showP gopurs_runtime.Value
				}{}
				clone.showP = gopurs_runtime.RecordGet(orig, "showP")
				return clone
			}()))}
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
		cache_Main_test2 = gopurs_runtime.Value{Type: 9, IntVal: 1349270669, UnsafePtr: unsafe.Pointer(Rebox_Main_1815506164_4226897273((&Constructor_Main_ShowP[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("Symbol")
		})})))}
	})
	return cache_Main_test2
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Value{Type: 9, IntVal: 1349270669, UnsafePtr: unsafe.Pointer(Rebox_Main_1815506164_4226897273((&Constructor_Main_ShowP[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("Type")
		})})))}
	})
	return cache_Main_test1
}

var cache_Main_showP gopurs_runtime.Value
var once_Main_showP sync.Once

func Get_Main_showP() gopurs_runtime.Value {
	once_Main_showP.Do(func() {
		cache_Main_showP = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_showP(gopurs_runtime.CoerceToStruct[Constructor_Main_ShowP[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_showP
}

var cache_Main_showP__4115038596 gopurs_runtime.Value
var once_Main_showP__4115038596 sync.Once

func Get_Main_showP__4115038596() gopurs_runtime.Value {
	once_Main_showP__4115038596.Do(func() {
		cache_Main_showP__4115038596 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_showP__4115038596(uint32(__eta_norm_0_unused_0_box.IntVal)))
		})
	})
	return cache_Main_showP__4115038596
}

var cache_Main_showP__2960655734 gopurs_runtime.Value
var once_Main_showP__2960655734 sync.Once

func Get_Main_showP__2960655734() gopurs_runtime.Value {
	once_Main_showP__2960655734.Do(func() {
		cache_Main_showP__2960655734 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_showP__2960655734(uint32(__eta_norm_0_unused_0_box.IntVal)))
		})
	})
	return cache_Main_showP__2960655734
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
			}))
		}))
	})
	return cache_Main_main
}

type Constructor_Main_Proxy[T_a any] struct {
	Rc uint32
}

type Constructor_Main_ShowP[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1349270669] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_ShowP[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "showP":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_ShowP: " + key)
		}
	}
}

func Call_Main_ShowP_dollar_Dict(x_0_loop struct {
	showP gopurs_runtime.Value
}) *Constructor_Main_ShowP[gopurs_runtime.Value] {
	var x_0 struct {
		showP gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_ShowP[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("showP", orig.showP)
	}())
}

func Call_Main_showP(dict_0_loop *Constructor_Main_ShowP[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_ShowP[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_showP__4115038596(__eta_norm_0_unused_0_loop uint32) string {
showP__4115038596:
	for {
		if false {
			continue showP__4115038596
		}
		var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return "Type"
	}
}

func Call_Main_showP__2960655734(__eta_norm_0_unused_0_loop uint32) string {
showP__2960655734:
	for {
		if false {
			continue showP__2960655734
		}
		var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return "Symbol"
	}
}

func Rebox_Main_1815506164_4226897273(in *Constructor_Main_ShowP[uint32]) *Constructor_Main_ShowP[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_ShowP[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
