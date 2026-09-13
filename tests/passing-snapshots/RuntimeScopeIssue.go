package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_B_dollar_Dict gopurs_runtime.Value
var once_Main_B_dollar_Dict sync.Once

func Get_Main_B_dollar_Dict() gopurs_runtime.Value {
	once_Main_B_dollar_Dict.Do(func() {
		cache_Main_B_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(Call_Main_B_dollar_Dict(func() struct {
				b gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					b gopurs_runtime.Value
				}{}
				clone.b = gopurs_runtime.RecordGet(orig, "b")
				return clone
			}()))}
		})
	})
	return cache_Main_B_dollar_Dict
}

var cache_Main_B_dollar_Dict__313467447 gopurs_runtime.Value
var once_Main_B_dollar_Dict__313467447 sync.Once

func Get_Main_B_dollar_Dict__313467447() gopurs_runtime.Value {
	once_Main_B_dollar_Dict__313467447.Do(func() {
		cache_Main_B_dollar_Dict__313467447 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(Rebox_Main_2084057936_3159052616(Call_Main_B_dollar_Dict__313467447(func() struct {
				b gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					b gopurs_runtime.Value
				}{}
				clone.b = gopurs_runtime.RecordGet(orig, "b")
				return clone
			}())))}
		})
	})
	return cache_Main_B_dollar_Dict__313467447
}

var cache_Main_A_dollar_Dict gopurs_runtime.Value
var once_Main_A_dollar_Dict sync.Once

func Get_Main_A_dollar_Dict() gopurs_runtime.Value {
	once_Main_A_dollar_Dict.Do(func() {
		cache_Main_A_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4219254943, UnsafePtr: unsafe.Pointer(Call_Main_A_dollar_Dict(func() struct {
				a gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					a gopurs_runtime.Value
				}{}
				clone.a = gopurs_runtime.RecordGet(orig, "a")
				return clone
			}()))}
		})
	})
	return cache_Main_A_dollar_Dict
}

var cache_Main_A_dollar_Dict__3954613399 gopurs_runtime.Value
var once_Main_A_dollar_Dict__3954613399 sync.Once

func Get_Main_A_dollar_Dict__3954613399() gopurs_runtime.Value {
	once_Main_A_dollar_Dict__3954613399.Do(func() {
		cache_Main_A_dollar_Dict__3954613399 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4219254943, UnsafePtr: unsafe.Pointer(Rebox_Main_2868984691_3943979371(Call_Main_A_dollar_Dict__3954613399(func() struct {
				a gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					a gopurs_runtime.Value
				}{}
				clone.a = gopurs_runtime.RecordGet(orig, "a")
				return clone
			}())))}
		})
	})
	return cache_Main_A_dollar_Dict__3954613399
}

var cache_Main_b gopurs_runtime.Value
var once_Main_b sync.Once

func Get_Main_b() gopurs_runtime.Value {
	once_Main_b.Do(func() {
		cache_Main_b = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_b(gopurs_runtime.CoerceToStruct[Constructor_Main_B[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_b
}

var cache_Main_b__4024554 gopurs_runtime.Value
var once_Main_b__4024554 sync.Once

func Get_Main_b__4024554() gopurs_runtime.Value {
	once_Main_b__4024554.Do(func() {
		cache_Main_b__4024554 = gopurs_runtime.Func(func(__eta_norm_0_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((gopurs_runtime.Apply(Rebox_Main_3159052616_2084057936(gopurs_runtime.CoerceToStruct[Constructor_Main_B[gopurs_runtime.Value]](Get_Main_bNumber())).V0, __eta_norm_0_0).IntVal) != (0))
		})
	})
	return cache_Main_b__4024554
}

var cache_Main_a__61794345 gopurs_runtime.Value
var once_Main_a__61794345 sync.Once

func Get_Main_a__61794345() gopurs_runtime.Value {
	once_Main_a__61794345.Do(func() {
		cache_Main_a__61794345 = gopurs_runtime.Func(func(__eta_norm_0_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((gopurs_runtime.Apply(Rebox_Main_3943979371_2868984691(gopurs_runtime.CoerceToStruct[Constructor_Main_A[gopurs_runtime.Value]](Get_Main_aNumber())).V0, __eta_norm_0_0).IntVal) != (0))
		})
	})
	return cache_Main_a__61794345
}

var cache_Main_bNumber gopurs_runtime.Value
var once_Main_bNumber sync.Once

func Get_Main_bNumber() gopurs_runtime.Value {
	once_Main_bNumber.Do(func() {
		cache_Main_bNumber = gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer(Rebox_Main_2084057936_3159052616((&Constructor_Main_B[float64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 bool
			{
				if (v_0.FloatVal()) == (0.0) {
					__t0 = false
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = (gopurs_runtime.Apply(Get_Main_a__61794345(), gopurs_runtime.Float((v_0.FloatVal())-(1.0))).IntVal) != (0)
			}
		end_branch_0:
			return gopurs_runtime.Bool(__t0)
		})})))}
	})
	return cache_Main_bNumber
}

var cache_Main_aNumber gopurs_runtime.Value
var once_Main_aNumber sync.Once

func Get_Main_aNumber() gopurs_runtime.Value {
	once_Main_aNumber.Do(func() {
		cache_Main_aNumber = gopurs_runtime.Value{Type: 9, IntVal: 4219254943, UnsafePtr: unsafe.Pointer(Rebox_Main_2868984691_3943979371((&Constructor_Main_A[float64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 bool
			{
				if (v_0.FloatVal()) == (0.0) {
					__t0 = true
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = (gopurs_runtime.Apply(Get_Main_b__4024554(), gopurs_runtime.Float((v_0.FloatVal())-(1.0))).IntVal) != (0)
			}
		end_branch_0:
			return gopurs_runtime.Bool(__t0)
		})})))}
	})
	return cache_Main_aNumber
}

var cache_Main_a gopurs_runtime.Value
var once_Main_a sync.Once

func Get_Main_a() gopurs_runtime.Value {
	once_Main_a.Do(func() {
		cache_Main_a = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_a(gopurs_runtime.CoerceToStruct[Constructor_Main_A[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_a
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Effect_Console_logShow(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))), gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Main_a__61794345(), gopurs_runtime.Float(8.0)).IntVal) != (0))), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
		}))
	})
	return cache_Main_main
}

type Constructor_Main_B[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4250879068] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_B[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "b":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_B: " + key)
		}
	}
}

type Constructor_Main_A[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4219254943] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_A[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "a":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_A: " + key)
		}
	}
}

func Call_Main_B_dollar_Dict(x_0_loop struct {
	b gopurs_runtime.Value
}) *Constructor_Main_B[gopurs_runtime.Value] {
	var x_0 struct {
		b gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_B[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("b", orig.b)
	}())
}

func Call_Main_B_dollar_Dict__313467447(x_0_loop struct {
	b gopurs_runtime.Value
}) *Constructor_Main_B[float64] {
B_dollar_Dict__313467447:
	for {
		if false {
			continue B_dollar_Dict__313467447
		}
		var x_0 struct {
			b gopurs_runtime.Value
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_B[float64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict1("b", orig.b)
		}())
	}
}

func Call_Main_A_dollar_Dict(x_0_loop struct {
	a gopurs_runtime.Value
}) *Constructor_Main_A[gopurs_runtime.Value] {
	var x_0 struct {
		a gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_A[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("a", orig.a)
	}())
}

func Call_Main_A_dollar_Dict__3954613399(x_0_loop struct {
	a gopurs_runtime.Value
}) *Constructor_Main_A[float64] {
A_dollar_Dict__3954613399:
	for {
		if false {
			continue A_dollar_Dict__3954613399
		}
		var x_0 struct {
			a gopurs_runtime.Value
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_A[float64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict1("a", orig.a)
		}())
	}
}

func Call_Main_b(dict_0_loop *Constructor_Main_B[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_B[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_a(dict_0_loop *Constructor_Main_A[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_A[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[bool]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2084057936_3159052616(in *Constructor_Main_B[float64]) *Constructor_Main_B[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_B[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2868984691_3943979371(in *Constructor_Main_A[float64]) *Constructor_Main_A[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_A[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3159052616_2084057936(in *Constructor_Main_B[gopurs_runtime.Value]) *Constructor_Main_B[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_B[float64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3943979371_2868984691(in *Constructor_Main_A[gopurs_runtime.Value]) *Constructor_Main_A[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_A[float64]{}
	out.V0 = in.V0
	return out
}
