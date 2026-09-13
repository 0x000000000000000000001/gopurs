package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Eq_dollar_Dict gopurs_runtime.Value
var once_Main_Eq_dollar_Dict sync.Once

func Get_Main_Eq_dollar_Dict() gopurs_runtime.Value {
	once_Main_Eq_dollar_Dict.Do(func() {
		cache_Main_Eq_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 61300330, UnsafePtr: unsafe.Pointer(Call_Main_Eq_dollar_Dict(func() struct {
				eq gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					eq gopurs_runtime.Value
				}{}
				clone.eq = gopurs_runtime.RecordGet(orig, "eq")
				return clone
			}()))}
		})
	})
	return cache_Main_Eq_dollar_Dict
}

var cache_Main_Eq_dollar_Dict__3060985643 gopurs_runtime.Value
var once_Main_Eq_dollar_Dict__3060985643 sync.Once

func Get_Main_Eq_dollar_Dict__3060985643() gopurs_runtime.Value {
	once_Main_Eq_dollar_Dict__3060985643.Do(func() {
		cache_Main_Eq_dollar_Dict__3060985643 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 61300330, UnsafePtr: unsafe.Pointer(Rebox_Main_3136802470_2946378686(Call_Main_Eq_dollar_Dict__3060985643(func() struct {
				eq gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					eq gopurs_runtime.Value
				}{}
				clone.eq = gopurs_runtime.RecordGet(orig, "eq")
				return clone
			}())))}
		})
	})
	return cache_Main_Eq_dollar_Dict__3060985643
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_eqNumber gopurs_runtime.Value
var once_Main_eqNumber sync.Once

func Get_Main_eqNumber() gopurs_runtime.Value {
	once_Main_eqNumber.Do(func() {
		cache_Main_eqNumber = gopurs_runtime.Value{Type: 9, IntVal: 61300330, UnsafePtr: unsafe.Pointer(Rebox_Main_3136802470_2946378686((&Constructor_Main_Eq[float64]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})})))}
	})
	return cache_Main_eqNumber
}

var cache_Main_eq gopurs_runtime.Value
var once_Main_eq sync.Once

func Get_Main_eq() gopurs_runtime.Value {
	once_Main_eq.Do(func() {
		cache_Main_eq = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eq(gopurs_runtime.CoerceToStruct[Constructor_Main_Eq[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_eq
}

type Constructor_Main_Eq[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[61300330] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Eq[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "eq":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Eq: " + key)
		}
	}
}

func Call_Main_Eq_dollar_Dict(x_0_loop struct {
	eq gopurs_runtime.Value
}) *Constructor_Main_Eq[gopurs_runtime.Value] {
	var x_0 struct {
		eq gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Eq[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("eq", orig.eq)
	}())
}

func Call_Main_Eq_dollar_Dict__3060985643(x_0_loop struct {
	eq gopurs_runtime.Value
}) *Constructor_Main_Eq[float64] {
Eq_dollar_Dict__3060985643:
	for {
		if false {
			continue Eq_dollar_Dict__3060985643
		}
		var x_0 struct {
			eq gopurs_runtime.Value
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_Eq[float64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict1("eq", orig.eq)
		}())
	}
}

func Call_Main_eq(dict_0_loop *Constructor_Main_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_Eq[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Rebox_Main_3136802470_2946378686(in *Constructor_Main_Eq[float64]) *Constructor_Main_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
