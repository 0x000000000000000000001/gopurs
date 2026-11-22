package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Extend_identity gopurs_runtime.Value
var once_Control_Extend_identity sync.Once

func Get_Control_Extend_identity() gopurs_runtime.Value {
	once_Control_Extend_identity.Do(func() {
		cache_Control_Extend_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Extend_identity(x_0_box)
		})
	})
	return cache_Control_Extend_identity
}

var cache_Control_Extend_Extend_dollar_Dict gopurs_runtime.Value
var once_Control_Extend_Extend_dollar_Dict sync.Once

func Get_Control_Extend_Extend_dollar_Dict() gopurs_runtime.Value {
	once_Control_Extend_Extend_dollar_Dict.Do(func() {
		cache_Control_Extend_Extend_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Call_Control_Extend_Extend_dollar_Dict(func() struct {
				Functor0 gopurs_runtime.Value
				extend   gopurs_runtime.Value
				/*gopurs:record-v1:2:00460075006e00630074006f00720030,0065007800740065006e0064*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Functor0 gopurs_runtime.Value
					extend   gopurs_runtime.Value
					/*gopurs:record-v1:2:00460075006e00630074006f00720030,0065007800740065006e0064*/
				}{}
				clone.Functor0 = gopurs_runtime.RecordGet(orig, "Functor0")
				clone.extend = gopurs_runtime.RecordGet(orig, "extend")
				return clone
			}()))}
		})
	})
	return cache_Control_Extend_Extend_dollar_Dict
}

var cache_Control_Extend_extendFn gopurs_runtime.Value
var once_Control_Extend_extendFn sync.Once

func Get_Control_Extend_extendFn() gopurs_runtime.Value {
	once_Control_Extend_extendFn.Do(func() {
		cache_Control_Extend_extendFn = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Extend_extendFn(dictSemigroup_0_box)
		})
	})
	return cache_Control_Extend_extendFn
}

var cache_Control_Extend_extendArray gopurs_runtime.Value
var once_Control_Extend_extendArray sync.Once

func Get_Control_Extend_extendArray() gopurs_runtime.Value {
	once_Control_Extend_extendArray.Do(func() {
		cache_Control_Extend_extendArray = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}
		}), Get_Control_Extend_arrayExtend()}))}
	})
	return cache_Control_Extend_extendArray
}

var cache_Control_Extend_extend gopurs_runtime.Value
var once_Control_Extend_extend sync.Once

func Get_Control_Extend_extend() gopurs_runtime.Value {
	once_Control_Extend_extend.Do(func() {
		cache_Control_Extend_extend = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Extend_extend(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Extend_extend
}

var cache_Control_Extend_extendFlipped gopurs_runtime.Value
var once_Control_Extend_extendFlipped sync.Once

func Get_Control_Extend_extendFlipped() gopurs_runtime.Value {
	once_Control_Extend_extendFlipped.Do(func() {
		cache_Control_Extend_extendFlipped = gopurs_runtime.Func3(func(dictExtend_0_box gopurs_runtime.Value, w_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Extend_extendFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](dictExtend_0_box), w_1_box, f_2_box)
		})
	})
	return cache_Control_Extend_extendFlipped
}

var cache_Control_Extend_duplicate gopurs_runtime.Value
var once_Control_Extend_duplicate sync.Once

func Get_Control_Extend_duplicate() gopurs_runtime.Value {
	once_Control_Extend_duplicate.Do(func() {
		cache_Control_Extend_duplicate = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Extend_duplicate(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](dictExtend_0_box))
		})
	})
	return cache_Control_Extend_duplicate
}

var cache_Control_Extend_composeCoKleisliFlipped gopurs_runtime.Value
var once_Control_Extend_composeCoKleisliFlipped sync.Once

func Get_Control_Extend_composeCoKleisliFlipped() gopurs_runtime.Value {
	once_Control_Extend_composeCoKleisliFlipped.Do(func() {
		cache_Control_Extend_composeCoKleisliFlipped = gopurs_runtime.Func4(func(dictExtend_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, w_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Extend_composeCoKleisliFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](dictExtend_0_box), f_1_box, g_2_box, w_3_box)
		})
	})
	return cache_Control_Extend_composeCoKleisliFlipped
}

var cache_Control_Extend_composeCoKleisli gopurs_runtime.Value
var once_Control_Extend_composeCoKleisli sync.Once

func Get_Control_Extend_composeCoKleisli() gopurs_runtime.Value {
	once_Control_Extend_composeCoKleisli.Do(func() {
		cache_Control_Extend_composeCoKleisli = gopurs_runtime.Func4(func(dictExtend_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, w_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Extend_composeCoKleisli(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](dictExtend_0_box), f_1_box, g_2_box, w_3_box)
		})
	})
	return cache_Control_Extend_composeCoKleisli
}

type Constructor_Control_Extend_Extend[T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f0045007800740065006e0064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f0045007800740065006e0064005f0045007800740065006e0064
func init() {
	gopurs_runtime.StructGetters[3028639021] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Extend_Extend[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Functor0":
			return gopurs_runtime.Box(c.V0)
		case "extend":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Extend_Extend: " + key)
		}
	}
}

func Call_Control_Extend_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Extend_Extend_dollar_Dict(x_0_loop struct {
	Functor0 gopurs_runtime.Value
	extend   gopurs_runtime.Value
	/*gopurs:record-v1:2:00460075006e00630074006f00720030,0065007800740065006e0064*/
}) *Constructor_Control_Extend_Extend[gopurs_runtime.Value] {
	var x_0 struct {
		Functor0 gopurs_runtime.Value
		extend   gopurs_runtime.Value
		/*gopurs:record-v1:2:00460075006e00630074006f00720030,0065007800740065006e0064*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Functor0", "extend", orig.Functor0, orig.extend)
	}())
}

func Call_Control_Extend_extendFn(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
	_ = dictSemigroup_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Get_Data_Functor_functorFn()
	}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, w_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(f_1, gopurs_runtime.Func(func(w_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), w_3, w_prime__4))
		}))
	})}))}
}

func Call_Control_Extend_extend(dict_0_loop *Constructor_Control_Extend_Extend[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Extend_Extend[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Control_Extend_extendFlipped(dictExtend_0_loop *Constructor_Control_Extend_Extend[gopurs_runtime.Value], w_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictExtend_0 *Constructor_Control_Extend_Extend[gopurs_runtime.Value] = dictExtend_0_loop
	_ = dictExtend_0
	var w_1 gopurs_runtime.Value = w_1_loop
	_ = w_1
	var f_2 gopurs_runtime.Value = f_2_loop
	_ = f_2
	return gopurs_runtime.Apply2(dictExtend_0.V1, f_2, w_1)
}

func Call_Control_Extend_duplicate(dictExtend_0_loop *Constructor_Control_Extend_Extend[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictExtend_0 *Constructor_Control_Extend_Extend[gopurs_runtime.Value] = dictExtend_0_loop
	_ = dictExtend_0
	return gopurs_runtime.Apply(dictExtend_0.V1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}

func Call_Control_Extend_composeCoKleisliFlipped(dictExtend_0_loop *Constructor_Control_Extend_Extend[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, w_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictExtend_0 *Constructor_Control_Extend_Extend[gopurs_runtime.Value] = dictExtend_0_loop
	_ = dictExtend_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var g_2 gopurs_runtime.Value = g_2_loop
	_ = g_2
	var w_3 gopurs_runtime.Value = w_3_loop
	_ = w_3
	return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(dictExtend_0.V1, g_2, w_3))
}

func Call_Control_Extend_composeCoKleisli(dictExtend_0_loop *Constructor_Control_Extend_Extend[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, w_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictExtend_0 *Constructor_Control_Extend_Extend[gopurs_runtime.Value] = dictExtend_0_loop
	_ = dictExtend_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var g_2 gopurs_runtime.Value = g_2_loop
	_ = g_2
	var w_3 gopurs_runtime.Value = w_3_loop
	_ = w_3
	return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply2(dictExtend_0.V1, f_1, w_3))
}

func Get_Control_Extend_arrayExtend() gopurs_runtime.Value {
	return _Gopurs_Control_Extend_ArrayExtend
}
