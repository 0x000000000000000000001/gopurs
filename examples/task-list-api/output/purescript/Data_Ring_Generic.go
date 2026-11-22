package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Ring_Generic_GenericRing_dollar_Dict gopurs_runtime.Value
var once_Data_Ring_Generic_GenericRing_dollar_Dict sync.Once

func Get_Data_Ring_Generic_GenericRing_dollar_Dict() gopurs_runtime.Value {
	once_Data_Ring_Generic_GenericRing_dollar_Dict.Do(func() {
		cache_Data_Ring_Generic_GenericRing_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3896698597, UnsafePtr: unsafe.Pointer(Call_Data_Ring_Generic_GenericRing_dollar_Dict(func() struct {
				gopurs_field_103_101_110_101_114_105_99_83_117_98_39 gopurs_runtime.Value
				/*gopurs:record-v1:1:00670065006e00650072006900630053007500620027*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					gopurs_field_103_101_110_101_114_105_99_83_117_98_39 gopurs_runtime.Value
					/*gopurs:record-v1:1:00670065006e00650072006900630053007500620027*/
				}{}
				clone.gopurs_field_103_101_110_101_114_105_99_83_117_98_39 = gopurs_runtime.RecordGet(orig, "genericSub'")
				return clone
			}()))}
		})
	})
	return cache_Data_Ring_Generic_GenericRing_dollar_Dict
}

var cache_Data_Ring_Generic_GenericRing_dollar_Dict__456889733 gopurs_runtime.Value
var once_Data_Ring_Generic_GenericRing_dollar_Dict__456889733 sync.Once

func Get_Data_Ring_Generic_GenericRing_dollar_Dict__456889733() gopurs_runtime.Value {
	once_Data_Ring_Generic_GenericRing_dollar_Dict__456889733.Do(func() {
		cache_Data_Ring_Generic_GenericRing_dollar_Dict__456889733 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3896698597, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_Generic_1088399036_3123454929(Call_Data_Ring_Generic_GenericRing_dollar_Dict__456889733(func() struct {
				gopurs_field_103_101_110_101_114_105_99_83_117_98_39 gopurs_runtime.Value
				/*gopurs:record-v1:1:00670065006e00650072006900630053007500620027*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					gopurs_field_103_101_110_101_114_105_99_83_117_98_39 gopurs_runtime.Value
					/*gopurs:record-v1:1:00670065006e00650072006900630053007500620027*/
				}{}
				clone.gopurs_field_103_101_110_101_114_105_99_83_117_98_39 = gopurs_runtime.RecordGet(orig, "genericSub'")
				return clone
			}())))}
		})
	})
	return cache_Data_Ring_Generic_GenericRing_dollar_Dict__456889733
}

var cache_Data_Ring_Generic_genericSub_prime_ gopurs_runtime.Value
var once_Data_Ring_Generic_genericSub_prime_ sync.Once

func Get_Data_Ring_Generic_genericSub_prime_() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericSub_prime_.Do(func() {
		cache_Data_Ring_Generic_genericSub_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ring_Generic_genericSub_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Ring_Generic_genericSub_prime_
}

var cache_Data_Ring_Generic_genericSub gopurs_runtime.Value
var once_Data_Ring_Generic_genericSub sync.Once

func Get_Data_Ring_Generic_genericSub() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericSub.Do(func() {
		cache_Data_Ring_Generic_genericSub = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ring_Generic_genericSub(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box))
		})
	})
	return cache_Data_Ring_Generic_genericSub
}

var cache_Data_Ring_Generic_genericRingProduct gopurs_runtime.Value
var once_Data_Ring_Generic_genericRingProduct sync.Once

func Get_Data_Ring_Generic_genericRingProduct() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericRingProduct.Do(func() {
		cache_Data_Ring_Generic_genericRingProduct = gopurs_runtime.Func2(func(dictGenericRing_0_box gopurs_runtime.Value, dictGenericRing1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ring_Generic_genericRingProduct(dictGenericRing_0_box, dictGenericRing1_1_box)
		})
	})
	return cache_Data_Ring_Generic_genericRingProduct
}

var cache_Data_Ring_Generic_genericRingNoArguments gopurs_runtime.Value
var once_Data_Ring_Generic_genericRingNoArguments sync.Once

func Get_Data_Ring_Generic_genericRingNoArguments() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericRingNoArguments.Do(func() {
		cache_Data_Ring_Generic_genericRingNoArguments = gopurs_runtime.Value{Type: 9, IntVal: 3896698597, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_Generic_1088399036_3123454929((&Constructor_Data_Ring_Generic_GenericRing[uint32]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
		})})))}
	})
	return cache_Data_Ring_Generic_genericRingNoArguments
}

var cache_Data_Ring_Generic_genericRingConstructor gopurs_runtime.Value
var once_Data_Ring_Generic_genericRingConstructor sync.Once

func Get_Data_Ring_Generic_genericRingConstructor() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericRingConstructor.Do(func() {
		cache_Data_Ring_Generic_genericRingConstructor = gopurs_runtime.Func(func(dictGenericRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ring_Generic_genericRingConstructor(dictGenericRing_0_box)
		})
	})
	return cache_Data_Ring_Generic_genericRingConstructor
}

var cache_Data_Ring_Generic_genericRingArgument gopurs_runtime.Value
var once_Data_Ring_Generic_genericRingArgument sync.Once

func Get_Data_Ring_Generic_genericRingArgument() gopurs_runtime.Value {
	once_Data_Ring_Generic_genericRingArgument.Do(func() {
		cache_Data_Ring_Generic_genericRingArgument = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ring_Generic_genericRingArgument(dictRing_0_box)
		})
	})
	return cache_Data_Ring_Generic_genericRingArgument
}

type Constructor_Data_Ring_Generic_GenericRing[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00520069006e0067005f00470065006e0065007200690063:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00520069006e0067005f00470065006e0065007200690063005f00470065006e006500720069006300520069006e0067
func init() {
	gopurs_runtime.StructGetters[3896698597] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "genericSub'":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Data_Ring_Generic_GenericRing: " + key)
		}
	}
}

func Call_Data_Ring_Generic_GenericRing_dollar_Dict(x_0_loop struct {
	gopurs_field_103_101_110_101_114_105_99_83_117_98_39 gopurs_runtime.Value
	/*gopurs:record-v1:1:00670065006e00650072006900630053007500620027*/
}) *Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value] {
	var x_0 struct {
		gopurs_field_103_101_110_101_114_105_99_83_117_98_39 gopurs_runtime.Value
		/*gopurs:record-v1:1:00670065006e00650072006900630053007500620027*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("genericSub'", orig.gopurs_field_103_101_110_101_114_105_99_83_117_98_39)
	}())
}

func Call_Data_Ring_Generic_GenericRing_dollar_Dict__456889733(x_0_loop struct {
	gopurs_field_103_101_110_101_114_105_99_83_117_98_39 gopurs_runtime.Value
	/*gopurs:record-v1:1:00670065006e00650072006900630053007500620027*/
}) *Constructor_Data_Ring_Generic_GenericRing[uint32] {
GenericRing_dollar_Dict__456889733:
	for {
		if false {
			continue GenericRing_dollar_Dict__456889733
		}
		var x_0 struct {
			gopurs_field_103_101_110_101_114_105_99_83_117_98_39 gopurs_runtime.Value
			/*gopurs:record-v1:1:00670065006e00650072006900630053007500620027*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Generic_GenericRing[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict1("genericSub'", orig.gopurs_field_103_101_110_101_114_105_99_83_117_98_39)
		}())
	}
}

func Call_Data_Ring_Generic_genericSub_prime_(dict_0_loop *Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Data_Ring_Generic_genericSub(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
	_ = dictGeneric_0
	// TAST (Let): to_1_0 shape=App(Var) bindingType=(Func [Any] (TypeVar a$scope3))
	to_1_0 := Call_Data_Generic_Rep_to(dictGeneric_0)
	_ = to_1_0
	return gopurs_runtime.Func3(func(dictGenericRing_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(to_1_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_2, "genericSub'"), gopurs_runtime.Apply(dictGeneric_0.V0, x_3), gopurs_runtime.Apply(dictGeneric_0.V0, y_4)))
	})
}

func Call_Data_Ring_Generic_genericRingProduct(dictGenericRing_0_loop gopurs_runtime.Value, dictGenericRing1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictGenericRing_0 gopurs_runtime.Value = dictGenericRing_0_loop
	_ = dictGenericRing_0
	var dictGenericRing1_1 gopurs_runtime.Value = dictGenericRing1_1_loop
	_ = dictGenericRing1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 3896698597, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_Generic_2426474238_3123454929((&Constructor_Data_Ring_Generic_GenericRing[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing1_1, "genericSub'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
	})})))}
}

func Call_Data_Ring_Generic_genericRingConstructor(dictGenericRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictGenericRing_0 gopurs_runtime.Value = dictGenericRing_0_loop
	_ = dictGenericRing_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3896698597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), v_1, v1_2)
	})}))}
}

func Call_Data_Ring_Generic_genericRingArgument(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
	_ = dictRing_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3896698597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), v_1, v1_2)
	})}))}
}

func Rebox_Data_Ring_Generic_1088399036_3123454929(in *Constructor_Data_Ring_Generic_GenericRing[uint32]) *Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ring_Generic_2426474238_3123454929(in *Constructor_Data_Ring_Generic_GenericRing[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ring_Generic_GenericRing[gopurs_runtime.Value])(unsafe.Pointer(in))
}
