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

var cache_Main_Id__3670367442 gopurs_runtime.Value
var once_Main_Id__3670367442 sync.Once

func Get_Main_Id__3670367442() gopurs_runtime.Value {
	once_Main_Id__3670367442.Do(func() {
		cache_Main_Id__3670367442 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Id__3670367442(__eta_norm_0_0_box.FloatVal())
		})
	})
	return cache_Main_Id__3670367442
}

var cache_Main_E_dollar_Dict gopurs_runtime.Value
var once_Main_E_dollar_Dict sync.Once

func Get_Main_E_dollar_Dict() gopurs_runtime.Value {
	once_Main_E_dollar_Dict.Do(func() {
		cache_Main_E_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1955825563, UnsafePtr: unsafe.Pointer(Call_Main_E_dollar_Dict(func() struct {
				add gopurs_runtime.Value
				num gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					add gopurs_runtime.Value
					num gopurs_runtime.Value
				}{}
				clone.add = gopurs_runtime.RecordGet(orig, "add")
				clone.num = gopurs_runtime.RecordGet(orig, "num")
				return clone
			}()))}
		})
	})
	return cache_Main_E_dollar_Dict
}

var cache_Main_E_dollar_Dict__3879561812 gopurs_runtime.Value
var once_Main_E_dollar_Dict__3879561812 sync.Once

func Get_Main_E_dollar_Dict__3879561812() gopurs_runtime.Value {
	once_Main_E_dollar_Dict__3879561812.Do(func() {
		cache_Main_E_dollar_Dict__3879561812 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1955825563, UnsafePtr: unsafe.Pointer(Call_Main_E_dollar_Dict__3879561812(func() struct {
				add gopurs_runtime.Value
				num gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					add gopurs_runtime.Value
					num gopurs_runtime.Value
				}{}
				clone.add = gopurs_runtime.RecordGet(orig, "add")
				clone.num = gopurs_runtime.RecordGet(orig, "num")
				return clone
			}()))}
		})
	})
	return cache_Main_E_dollar_Dict__3879561812
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

var cache_Main_exprId gopurs_runtime.Value
var once_Main_exprId sync.Once

func Get_Main_exprId() gopurs_runtime.Value {
	once_Main_exprId.Do(func() {
		cache_Main_exprId = gopurs_runtime.Value{Type: 9, IntVal: 1955825563, UnsafePtr: unsafe.Pointer((&Constructor_Main_E[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
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
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Effect_Console_logShow(Rebox_Main_3263178038_1386611502(Rebox_Main_1386611502_3263178038(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showNumber()))), gopurs_runtime.Float(Call_Main_runId(Call_Main_three(gopurs_runtime.Value{Type: 9, IntVal: 1955825563, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_E[gopurs_runtime.Value]](Get_Main_exprId()))})).FloatVal())), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

type Constructor_Main_Id[T_a any] struct {
	Rc uint32
	V0 T_a
}

type Constructor_Main_E[T_e any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1955825563] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_E[gopurs_runtime.Value])(ptr)
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

func Call_Main_Id__3670367442(__eta_norm_0_0_loop float64) gopurs_runtime.Value {
Id__3670367442:
	for {
		if false {
			continue Id__3670367442
		}
		var __eta_norm_0_0 float64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Float(__eta_norm_0_0)
	}
}

func Call_Main_E_dollar_Dict(x_0_loop struct {
	add gopurs_runtime.Value
	num gopurs_runtime.Value
}) *Constructor_Main_E[gopurs_runtime.Value] {
	var x_0 struct {
		add gopurs_runtime.Value
		num gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_E[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("add", "num", orig.add, orig.num)
	}())
}

func Call_Main_E_dollar_Dict__3879561812(x_0_loop struct {
	add gopurs_runtime.Value
	num gopurs_runtime.Value
}) *Constructor_Main_E[gopurs_runtime.Value] {
E_dollar_Dict__3879561812:
	for {
		if false {
			continue E_dollar_Dict__3879561812
		}
		var x_0 struct {
			add gopurs_runtime.Value
			num gopurs_runtime.Value
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_E[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("add", "num", orig.add, orig.num)
		}())
	}
}

func Call_Main_runId(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return v_0
}

func Call_Main_num(dict_0_loop *Constructor_Main_E[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_E[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Main_add(dict_0_loop *Constructor_Main_E[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_E[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_three(dictE_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictE_0 gopurs_runtime.Value = dictE_0_loop
	_ = dictE_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictE_0, "add"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictE_0, "num"), gopurs_runtime.Float(1.0)), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictE_0, "num"), gopurs_runtime.Float(2.0)))
}

func Rebox_Main_1386611502_3263178038(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[float64]{}
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
