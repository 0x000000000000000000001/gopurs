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
			return Call_Main_B_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_B_dollar_Dict
}

var cache_Main_A_dollar_Dict gopurs_runtime.Value
var once_Main_A_dollar_Dict sync.Once

func Get_Main_A_dollar_Dict() gopurs_runtime.Value {
	once_Main_A_dollar_Dict.Do(func() {
		cache_Main_A_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_A_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_A_dollar_Dict
}

var cache_Main_b gopurs_runtime.Value
var once_Main_b sync.Once

func Get_Main_b() gopurs_runtime.Value {
	once_Main_b.Do(func() {
		cache_Main_b = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_b(gopurs_runtime.CoerceToStruct[Constructor_Main_B](dict_0_box))
		})
	})
	return cache_Main_b
}

var cache_Main_a gopurs_runtime.Value
var once_Main_a sync.Once

func Get_Main_a() gopurs_runtime.Value {
	once_Main_a.Do(func() {
		cache_Main_a = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_a(gopurs_runtime.CoerceToStruct[Constructor_Main_A](dict_0_box))
		})
	})
	return cache_Main_a
}

var cache_Main_bNumber gopurs_runtime.Value
var once_Main_bNumber sync.Once

func Get_Main_bNumber() gopurs_runtime.Value {
	once_Main_bNumber.Do(func() {
		cache_Main_bNumber = gopurs_runtime.Value{Type: 9, IntVal: 4250879068, UnsafePtr: unsafe.Pointer((&Constructor_Main_B{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 bool
			{
				if (v_0.FloatVal()) == (0.0) {
					__t0 = false
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = (gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Main_A](Get_Main_aNumber()).V0), gopurs_runtime.Float((v_0.FloatVal())-(1.0))).IntVal) != (0)
			}
		end_branch_0:
			return gopurs_runtime.Bool(__t0)
		})}))}
	})
	return cache_Main_bNumber
}

var cache_Main_aNumber gopurs_runtime.Value
var once_Main_aNumber sync.Once

func Get_Main_aNumber() gopurs_runtime.Value {
	once_Main_aNumber.Do(func() {
		cache_Main_aNumber = gopurs_runtime.Value{Type: 9, IntVal: 4219254943, UnsafePtr: unsafe.Pointer((&Constructor_Main_A{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t2 bool
			{
				if (v_0.FloatVal()) == (0.0) {
					__t2 = true
					goto end_branch_2
				} else {

				}
			}
			{
				// TAST (Let): __local_var_1_0 -> float64
				__local_var_1_0 := (v_0.FloatVal()) - (1.0)
				_ = __local_var_1_0
				var __t1 bool
				{
					if (__local_var_1_0) == (0.0) {
						__t1 = false
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = (gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Main_A](Get_Main_aNumber()).V0), gopurs_runtime.Float((__local_var_1_0)-(1.0))).IntVal) != (0)
				}
			end_branch_1:
				__t2 = __t1
			}
		end_branch_2:
			return gopurs_runtime.Bool(__t2)
		})}))}
	})
	return cache_Main_aNumber
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 string
			{
				if (gopurs_runtime.Bool((gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Main_A](Get_Main_aNumber()).V0), gopurs_runtime.Float(8.0)).IntVal) != (0)).IntVal) != (0) {
					__t1 = "true"
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = "false"
			}
		end_branch_1:
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t1))
			_ = __local_var_0_0
			_dollar___unused_1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_a__1664477318 gopurs_runtime.Value
var once_Main_a__1664477318 sync.Once

func Get_Main_a__1664477318() gopurs_runtime.Value {
	once_Main_a__1664477318.Do(func() {
		cache_Main_a__1664477318 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_a__1664477318(gopurs_runtime.CoerceToStruct[Constructor_Main_A](dict_0_box))
		})
	})
	return cache_Main_a__1664477318
}

var cache_Main_b__9469573 gopurs_runtime.Value
var once_Main_b__9469573 sync.Once

func Get_Main_b__9469573() gopurs_runtime.Value {
	once_Main_b__9469573.Do(func() {
		cache_Main_b__9469573 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_b__9469573(gopurs_runtime.CoerceToStruct[Constructor_Main_B](dict_0_box))
		})
	})
	return cache_Main_b__9469573
}

type Constructor_Main_B struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4250879068] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_B)(ptr)
		_ = c
		switch key {
		case "b":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_B: " + key)
		}
	}
}

type Constructor_Main_A struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4219254943] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_A)(ptr)
		_ = c
		switch key {
		case "a":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_A: " + key)
		}
	}
}

func Call_Main_B_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_A_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_b(dict_0_loop *Constructor_Main_B) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_B = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_a(dict_0_loop *Constructor_Main_A) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_A = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_a__1664477318(dict_0_loop *Constructor_Main_A) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_A = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_b__9469573(dict_0_loop *Constructor_Main_B) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_B = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}
