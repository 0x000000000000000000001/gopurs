package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Semiring_Generic_GenericSemiring_dollar_Dict gopurs_runtime.Value
var once_Data_Semiring_Generic_GenericSemiring_dollar_Dict sync.Once

func Get_Data_Semiring_Generic_GenericSemiring_dollar_Dict() gopurs_runtime.Value {
	once_Data_Semiring_Generic_GenericSemiring_dollar_Dict.Do(func() {
		cache_Data_Semiring_Generic_GenericSemiring_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(Call_Data_Semiring_Generic_GenericSemiring_dollar_Dict(func() struct {
				gopurs_field_103_101_110_101_114_105_99_65_100_100_39     gopurs_runtime.Value
				gopurs_field_103_101_110_101_114_105_99_77_117_108_39     gopurs_runtime.Value
				gopurs_field_103_101_110_101_114_105_99_79_110_101_39     gopurs_runtime.Value
				gopurs_field_103_101_110_101_114_105_99_90_101_114_111_39 gopurs_runtime.Value
				/*gopurs:record-v1:4:00670065006e00650072006900630041006400640027,00670065006e0065007200690063004d0075006c0027,00670065006e0065007200690063004f006e00650027,00670065006e0065007200690063005a00650072006f0027*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					gopurs_field_103_101_110_101_114_105_99_65_100_100_39     gopurs_runtime.Value
					gopurs_field_103_101_110_101_114_105_99_77_117_108_39     gopurs_runtime.Value
					gopurs_field_103_101_110_101_114_105_99_79_110_101_39     gopurs_runtime.Value
					gopurs_field_103_101_110_101_114_105_99_90_101_114_111_39 gopurs_runtime.Value
					/*gopurs:record-v1:4:00670065006e00650072006900630041006400640027,00670065006e0065007200690063004d0075006c0027,00670065006e0065007200690063004f006e00650027,00670065006e0065007200690063005a00650072006f0027*/
				}{}
				clone.gopurs_field_103_101_110_101_114_105_99_65_100_100_39 = gopurs_runtime.RecordGet(orig, "genericAdd'")
				clone.gopurs_field_103_101_110_101_114_105_99_77_117_108_39 = gopurs_runtime.RecordGet(orig, "genericMul'")
				clone.gopurs_field_103_101_110_101_114_105_99_79_110_101_39 = gopurs_runtime.RecordGet(orig, "genericOne'")
				clone.gopurs_field_103_101_110_101_114_105_99_90_101_114_111_39 = gopurs_runtime.RecordGet(orig, "genericZero'")
				return clone
			}()))}
		})
	})
	return cache_Data_Semiring_Generic_GenericSemiring_dollar_Dict
}

var cache_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835 gopurs_runtime.Value
var once_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835 sync.Once

func Get_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835() gopurs_runtime.Value {
	once_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835.Do(func() {
		cache_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(Rebox_Data_Semiring_Generic_2984329148_1889332433(Call_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835(func() struct {
				gopurs_field_103_101_110_101_114_105_99_65_100_100_39     gopurs_runtime.Value
				gopurs_field_103_101_110_101_114_105_99_77_117_108_39     gopurs_runtime.Value
				gopurs_field_103_101_110_101_114_105_99_79_110_101_39     uint32
				gopurs_field_103_101_110_101_114_105_99_90_101_114_111_39 uint32
				/*gopurs:record-v1:4:00670065006e00650072006900630041006400640027,00670065006e0065007200690063004d0075006c0027,00670065006e0065007200690063004f006e00650027,00670065006e0065007200690063005a00650072006f0027*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					gopurs_field_103_101_110_101_114_105_99_65_100_100_39     gopurs_runtime.Value
					gopurs_field_103_101_110_101_114_105_99_77_117_108_39     gopurs_runtime.Value
					gopurs_field_103_101_110_101_114_105_99_79_110_101_39     uint32
					gopurs_field_103_101_110_101_114_105_99_90_101_114_111_39 uint32
					/*gopurs:record-v1:4:00670065006e00650072006900630041006400640027,00670065006e0065007200690063004d0075006c0027,00670065006e0065007200690063004f006e00650027,00670065006e0065007200690063005a00650072006f0027*/
				}{}
				clone.gopurs_field_103_101_110_101_114_105_99_65_100_100_39 = gopurs_runtime.RecordGet(orig, "genericAdd'")
				clone.gopurs_field_103_101_110_101_114_105_99_77_117_108_39 = gopurs_runtime.RecordGet(orig, "genericMul'")
				clone.gopurs_field_103_101_110_101_114_105_99_79_110_101_39 = uint32(gopurs_runtime.RecordGet(orig, "genericOne'").IntVal)
				clone.gopurs_field_103_101_110_101_114_105_99_90_101_114_111_39 = uint32(gopurs_runtime.RecordGet(orig, "genericZero'").IntVal)
				return clone
			}())))}
		})
	})
	return cache_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835
}

var cache_Data_Semiring_Generic_genericZero_prime_ gopurs_runtime.Value
var once_Data_Semiring_Generic_genericZero_prime_ sync.Once

func Get_Data_Semiring_Generic_genericZero_prime_() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericZero_prime_.Do(func() {
		cache_Data_Semiring_Generic_genericZero_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semiring_Generic_genericZero_prime_(dict_0_box)
		})
	})
	return cache_Data_Semiring_Generic_genericZero_prime_
}

var cache_Data_Semiring_Generic_genericZero gopurs_runtime.Value
var once_Data_Semiring_Generic_genericZero sync.Once

func Get_Data_Semiring_Generic_genericZero() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericZero.Do(func() {
		cache_Data_Semiring_Generic_genericZero = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semiring_Generic_genericZero(dictGeneric_0_box, dictGenericSemiring_1_box)
		})
	})
	return cache_Data_Semiring_Generic_genericZero
}

var cache_Data_Semiring_Generic_genericSemiringNoArguments gopurs_runtime.Value
var once_Data_Semiring_Generic_genericSemiringNoArguments sync.Once

func Get_Data_Semiring_Generic_genericSemiringNoArguments() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericSemiringNoArguments.Do(func() {
		cache_Data_Semiring_Generic_genericSemiringNoArguments = gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(Rebox_Data_Semiring_Generic_2984329148_1889332433((&Constructor_Data_Semiring_Generic_GenericSemiring[uint32]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
		}), 1454898258, 1454898258})))}
	})
	return cache_Data_Semiring_Generic_genericSemiringNoArguments
}

var cache_Data_Semiring_Generic_genericSemiringArgument gopurs_runtime.Value
var once_Data_Semiring_Generic_genericSemiringArgument sync.Once

func Get_Data_Semiring_Generic_genericSemiringArgument() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericSemiringArgument.Do(func() {
		cache_Data_Semiring_Generic_genericSemiringArgument = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semiring_Generic_genericSemiringArgument(dictSemiring_0_box)
		})
	})
	return cache_Data_Semiring_Generic_genericSemiringArgument
}

var cache_Data_Semiring_Generic_genericOne_prime_ gopurs_runtime.Value
var once_Data_Semiring_Generic_genericOne_prime_ sync.Once

func Get_Data_Semiring_Generic_genericOne_prime_() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericOne_prime_.Do(func() {
		cache_Data_Semiring_Generic_genericOne_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semiring_Generic_genericOne_prime_(dict_0_box)
		})
	})
	return cache_Data_Semiring_Generic_genericOne_prime_
}

var cache_Data_Semiring_Generic_genericOne gopurs_runtime.Value
var once_Data_Semiring_Generic_genericOne sync.Once

func Get_Data_Semiring_Generic_genericOne() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericOne.Do(func() {
		cache_Data_Semiring_Generic_genericOne = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semiring_Generic_genericOne(dictGeneric_0_box, dictGenericSemiring_1_box)
		})
	})
	return cache_Data_Semiring_Generic_genericOne
}

var cache_Data_Semiring_Generic_genericMul_prime_ gopurs_runtime.Value
var once_Data_Semiring_Generic_genericMul_prime_ sync.Once

func Get_Data_Semiring_Generic_genericMul_prime_() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericMul_prime_.Do(func() {
		cache_Data_Semiring_Generic_genericMul_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semiring_Generic_genericMul_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Semiring_Generic_genericMul_prime_
}

var cache_Data_Semiring_Generic_genericMul gopurs_runtime.Value
var once_Data_Semiring_Generic_genericMul sync.Once

func Get_Data_Semiring_Generic_genericMul() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericMul.Do(func() {
		cache_Data_Semiring_Generic_genericMul = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semiring_Generic_genericMul(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box))
		})
	})
	return cache_Data_Semiring_Generic_genericMul
}

var cache_Data_Semiring_Generic_genericAdd_prime_ gopurs_runtime.Value
var once_Data_Semiring_Generic_genericAdd_prime_ sync.Once

func Get_Data_Semiring_Generic_genericAdd_prime_() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericAdd_prime_.Do(func() {
		cache_Data_Semiring_Generic_genericAdd_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semiring_Generic_genericAdd_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Semiring_Generic_genericAdd_prime_
}

var cache_Data_Semiring_Generic_genericSemiringConstructor gopurs_runtime.Value
var once_Data_Semiring_Generic_genericSemiringConstructor sync.Once

func Get_Data_Semiring_Generic_genericSemiringConstructor() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericSemiringConstructor.Do(func() {
		cache_Data_Semiring_Generic_genericSemiringConstructor = gopurs_runtime.Func(func(dictGenericSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semiring_Generic_genericSemiringConstructor(dictGenericSemiring_0_box)
		})
	})
	return cache_Data_Semiring_Generic_genericSemiringConstructor
}

var cache_Data_Semiring_Generic_genericSemiringProduct gopurs_runtime.Value
var once_Data_Semiring_Generic_genericSemiringProduct sync.Once

func Get_Data_Semiring_Generic_genericSemiringProduct() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericSemiringProduct.Do(func() {
		cache_Data_Semiring_Generic_genericSemiringProduct = gopurs_runtime.Func(func(dictGenericSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semiring_Generic_genericSemiringProduct(dictGenericSemiring_0_box)
		})
	})
	return cache_Data_Semiring_Generic_genericSemiringProduct
}

var cache_Data_Semiring_Generic_genericAdd gopurs_runtime.Value
var once_Data_Semiring_Generic_genericAdd sync.Once

func Get_Data_Semiring_Generic_genericAdd() gopurs_runtime.Value {
	once_Data_Semiring_Generic_genericAdd.Do(func() {
		cache_Data_Semiring_Generic_genericAdd = gopurs_runtime.Func(func(dictGeneric_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Semiring_Generic_genericAdd(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box))
		})
	})
	return cache_Data_Semiring_Generic_genericAdd
}

type Constructor_Data_Semiring_Generic_GenericSemiring[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 T_a
	V3 T_a
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00530065006d006900720069006e0067005f00470065006e0065007200690063:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00530065006d006900720069006e0067005f00470065006e0065007200690063005f00470065006e006500720069006300530065006d006900720069006e0067
func init() {
	gopurs_runtime.StructGetters[3093875941] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "genericAdd'":
			return gopurs_runtime.Box(c.V0)
		case "genericMul'":
			return gopurs_runtime.Box(c.V1)
		case "genericOne'":
			return gopurs_runtime.Box(c.V2)
		case "genericZero'":
			return gopurs_runtime.Box(c.V3)
		default:
			panic("Key not found in dictionary Constructor_Data_Semiring_Generic_GenericSemiring: " + key)
		}
	}
}

func Call_Data_Semiring_Generic_GenericSemiring_dollar_Dict(x_0_loop struct {
	gopurs_field_103_101_110_101_114_105_99_65_100_100_39     gopurs_runtime.Value
	gopurs_field_103_101_110_101_114_105_99_77_117_108_39     gopurs_runtime.Value
	gopurs_field_103_101_110_101_114_105_99_79_110_101_39     gopurs_runtime.Value
	gopurs_field_103_101_110_101_114_105_99_90_101_114_111_39 gopurs_runtime.Value
	/*gopurs:record-v1:4:00670065006e00650072006900630041006400640027,00670065006e0065007200690063004d0075006c0027,00670065006e0065007200690063004f006e00650027,00670065006e0065007200690063005a00650072006f0027*/
}) *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value] {
	var x_0 struct {
		gopurs_field_103_101_110_101_114_105_99_65_100_100_39     gopurs_runtime.Value
		gopurs_field_103_101_110_101_114_105_99_77_117_108_39     gopurs_runtime.Value
		gopurs_field_103_101_110_101_114_105_99_79_110_101_39     gopurs_runtime.Value
		gopurs_field_103_101_110_101_114_105_99_90_101_114_111_39 gopurs_runtime.Value
		/*gopurs:record-v1:4:00670065006e00650072006900630041006400640027,00670065006e0065007200690063004d0075006c0027,00670065006e0065007200690063004f006e00650027,00670065006e0065007200690063005a00650072006f0027*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict4("genericAdd'", "genericMul'", "genericOne'", "genericZero'", orig.gopurs_field_103_101_110_101_114_105_99_65_100_100_39, orig.gopurs_field_103_101_110_101_114_105_99_77_117_108_39, orig.gopurs_field_103_101_110_101_114_105_99_79_110_101_39, orig.gopurs_field_103_101_110_101_114_105_99_90_101_114_111_39)
	}())
}

func Call_Data_Semiring_Generic_GenericSemiring_dollar_Dict__1911646835(x_0_loop struct {
	gopurs_field_103_101_110_101_114_105_99_65_100_100_39     gopurs_runtime.Value
	gopurs_field_103_101_110_101_114_105_99_77_117_108_39     gopurs_runtime.Value
	gopurs_field_103_101_110_101_114_105_99_79_110_101_39     uint32
	gopurs_field_103_101_110_101_114_105_99_90_101_114_111_39 uint32
	/*gopurs:record-v1:4:00670065006e00650072006900630041006400640027,00670065006e0065007200690063004d0075006c0027,00670065006e0065007200690063004f006e00650027,00670065006e0065007200690063005a00650072006f0027*/
}) *Constructor_Data_Semiring_Generic_GenericSemiring[uint32] {
GenericSemiring_dollar_Dict__1911646835:
	for {
		if false {
			continue GenericSemiring_dollar_Dict__1911646835
		}
		var x_0 struct {
			gopurs_field_103_101_110_101_114_105_99_65_100_100_39     gopurs_runtime.Value
			gopurs_field_103_101_110_101_114_105_99_77_117_108_39     gopurs_runtime.Value
			gopurs_field_103_101_110_101_114_105_99_79_110_101_39     uint32
			gopurs_field_103_101_110_101_114_105_99_90_101_114_111_39 uint32
			/*gopurs:record-v1:4:00670065006e00650072006900630041006400640027,00670065006e0065007200690063004d0075006c0027,00670065006e0065007200690063004f006e00650027,00670065006e0065007200690063005a00650072006f0027*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Generic_GenericSemiring[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict4("genericAdd'", "genericMul'", "genericOne'", "genericZero'", orig.gopurs_field_103_101_110_101_114_105_99_65_100_100_39, orig.gopurs_field_103_101_110_101_114_105_99_77_117_108_39, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.gopurs_field_103_101_110_101_114_105_99_79_110_101_39), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.gopurs_field_103_101_110_101_114_105_99_90_101_114_111_39), UnsafePtr: nil})
		}())
	}
}

func Call_Data_Semiring_Generic_genericZero_prime_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "genericZero'")
}

func Call_Data_Semiring_Generic_genericZero(dictGeneric_0_loop gopurs_runtime.Value, dictGenericSemiring_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
	_ = dictGeneric_0
	var dictGenericSemiring_1 gopurs_runtime.Value = dictGenericSemiring_1_loop
	_ = dictGenericSemiring_1
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericSemiring_1, "genericZero'"))
}

func Call_Data_Semiring_Generic_genericSemiringArgument(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
	_ = dictSemiring_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), v_1, v1_2)
	}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), v_1, v1_2)
	}), gopurs_runtime.RecordGet(dictSemiring_0, "one"), gopurs_runtime.RecordGet(dictSemiring_0, "zero")}))}
}

func Call_Data_Semiring_Generic_genericOne_prime_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "genericOne'")
}

func Call_Data_Semiring_Generic_genericOne(dictGeneric_0_loop gopurs_runtime.Value, dictGenericSemiring_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
	_ = dictGeneric_0
	var dictGenericSemiring_1 gopurs_runtime.Value = dictGenericSemiring_1_loop
	_ = dictGenericSemiring_1
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericSemiring_1, "genericOne'"))
}

func Call_Data_Semiring_Generic_genericMul_prime_(dict_0_loop *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Semiring_Generic_genericMul(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
	_ = dictGeneric_0
	// TAST (Let): to_1_0 shape=App(Var) bindingType=(Func [Any] (TypeVar a$scope21))
	to_1_0 := Call_Data_Generic_Rep_to(dictGeneric_0)
	_ = to_1_0
	return gopurs_runtime.Func3(func(dictGenericSemiring_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(to_1_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_2, "genericMul'"), gopurs_runtime.Apply(dictGeneric_0.V0, x_3), gopurs_runtime.Apply(dictGeneric_0.V0, y_4)))
	})
}

func Call_Data_Semiring_Generic_genericAdd_prime_(dict_0_loop *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Data_Semiring_Generic_genericSemiringConstructor(dictGenericSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictGenericSemiring_0 gopurs_runtime.Value = dictGenericSemiring_0_loop
	_ = dictGenericSemiring_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericAdd'"), v_1, v1_2)
	}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericMul'"), v_1, v1_2)
	}), gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericOne'"), gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericZero'")}))}
}

func Call_Data_Semiring_Generic_genericSemiringProduct(dictGenericSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictGenericSemiring_0 gopurs_runtime.Value = dictGenericSemiring_0_loop
	_ = dictGenericSemiring_0
	// TAST (Let): genericZero_prime_1_1_0 shape=Other bindingType=(TypeVar a$scope39)
	genericZero_prime_1_1_0 := gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericZero'")
	_ = genericZero_prime_1_1_0
	// TAST (Let): genericOne_prime_1_2_1 shape=Other bindingType=(TypeVar a$scope39)
	genericOne_prime_1_2_1 := gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericOne'")
	_ = genericOne_prime_1_2_1
	return gopurs_runtime.Func(func(dictGenericSemiring1_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 3093875941, UnsafePtr: unsafe.Pointer(Rebox_Data_Semiring_Generic_180684798_1889332433((&Constructor_Data_Semiring_Generic_GenericSemiring[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericAdd'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericAdd'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}))}
		}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0, "genericMul'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericMul'"), (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}))}
		}), (&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, genericOne_prime_1_2_1, gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericOne'")}), (&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, genericZero_prime_1_1_0, gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericZero'")})})))}
	})
}

func Call_Data_Semiring_Generic_genericAdd(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
	_ = dictGeneric_0
	// TAST (Let): to_1_0 shape=App(Var) bindingType=(Func [Any] (TypeVar a$scope52))
	to_1_0 := Call_Data_Generic_Rep_to(dictGeneric_0)
	_ = to_1_0
	return gopurs_runtime.Func3(func(dictGenericSemiring_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(to_1_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_2, "genericAdd'"), gopurs_runtime.Apply(dictGeneric_0.V0, x_3), gopurs_runtime.Apply(dictGeneric_0.V0, y_4)))
	})
}

func Rebox_Data_Semiring_Generic_180684798_1889332433(in *Constructor_Data_Semiring_Generic_GenericSemiring[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(in.V2)}
	out.V3 = gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(in.V3)}
	return out
}

func Rebox_Data_Semiring_Generic_2984329148_1889332433(in *Constructor_Data_Semiring_Generic_GenericSemiring[uint32]) *Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Generic_GenericSemiring[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
	out.V3 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V3), UnsafePtr: nil}
	return out
}
