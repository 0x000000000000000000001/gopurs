package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Monoid_Generic_GenericMonoid_dollar_Dict gopurs_runtime.Value
var once_Data_Monoid_Generic_GenericMonoid_dollar_Dict sync.Once

func Get_Data_Monoid_Generic_GenericMonoid_dollar_Dict() gopurs_runtime.Value {
	once_Data_Monoid_Generic_GenericMonoid_dollar_Dict.Do(func() {
		cache_Data_Monoid_Generic_GenericMonoid_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2569012965, UnsafePtr: unsafe.Pointer(Call_Data_Monoid_Generic_GenericMonoid_dollar_Dict(func() struct {
				gopurs_field_103_101_110_101_114_105_99_77_101_109_112_116_121_39 gopurs_runtime.Value
				/*gopurs:record-v1:1:00670065006e0065007200690063004d0065006d0070007400790027*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					gopurs_field_103_101_110_101_114_105_99_77_101_109_112_116_121_39 gopurs_runtime.Value
					/*gopurs:record-v1:1:00670065006e0065007200690063004d0065006d0070007400790027*/
				}{}
				clone.gopurs_field_103_101_110_101_114_105_99_77_101_109_112_116_121_39 = gopurs_runtime.RecordGet(orig, "genericMempty'")
				return clone
			}()))}
		})
	})
	return cache_Data_Monoid_Generic_GenericMonoid_dollar_Dict
}

var cache_Data_Monoid_Generic_GenericMonoid_dollar_Dict__121657880 gopurs_runtime.Value
var once_Data_Monoid_Generic_GenericMonoid_dollar_Dict__121657880 sync.Once

func Get_Data_Monoid_Generic_GenericMonoid_dollar_Dict__121657880() gopurs_runtime.Value {
	once_Data_Monoid_Generic_GenericMonoid_dollar_Dict__121657880.Do(func() {
		cache_Data_Monoid_Generic_GenericMonoid_dollar_Dict__121657880 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2569012965, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_Generic_3734302396_4174781905(Call_Data_Monoid_Generic_GenericMonoid_dollar_Dict__121657880(func() struct {
				gopurs_field_103_101_110_101_114_105_99_77_101_109_112_116_121_39 uint32
				/*gopurs:record-v1:1:00670065006e0065007200690063004d0065006d0070007400790027*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					gopurs_field_103_101_110_101_114_105_99_77_101_109_112_116_121_39 uint32
					/*gopurs:record-v1:1:00670065006e0065007200690063004d0065006d0070007400790027*/
				}{}
				clone.gopurs_field_103_101_110_101_114_105_99_77_101_109_112_116_121_39 = uint32(gopurs_runtime.RecordGet(orig, "genericMempty'").IntVal)
				return clone
			}())))}
		})
	})
	return cache_Data_Monoid_Generic_GenericMonoid_dollar_Dict__121657880
}

var cache_Data_Monoid_Generic_genericMonoidNoArguments gopurs_runtime.Value
var once_Data_Monoid_Generic_genericMonoidNoArguments sync.Once

func Get_Data_Monoid_Generic_genericMonoidNoArguments() gopurs_runtime.Value {
	once_Data_Monoid_Generic_genericMonoidNoArguments.Do(func() {
		cache_Data_Monoid_Generic_genericMonoidNoArguments = gopurs_runtime.Value{Type: 9, IntVal: 2569012965, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_Generic_3734302396_4174781905((&Constructor_Data_Monoid_Generic_GenericMonoid[uint32]{1, 1454898258})))}
	})
	return cache_Data_Monoid_Generic_genericMonoidNoArguments
}

var cache_Data_Monoid_Generic_genericMonoidArgument gopurs_runtime.Value
var once_Data_Monoid_Generic_genericMonoidArgument sync.Once

func Get_Data_Monoid_Generic_genericMonoidArgument() gopurs_runtime.Value {
	once_Data_Monoid_Generic_genericMonoidArgument.Do(func() {
		cache_Data_Monoid_Generic_genericMonoidArgument = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Generic_genericMonoidArgument(dictMonoid_0_box)
		})
	})
	return cache_Data_Monoid_Generic_genericMonoidArgument
}

var cache_Data_Monoid_Generic_genericMempty_prime_ gopurs_runtime.Value
var once_Data_Monoid_Generic_genericMempty_prime_ sync.Once

func Get_Data_Monoid_Generic_genericMempty_prime_() gopurs_runtime.Value {
	once_Data_Monoid_Generic_genericMempty_prime_.Do(func() {
		cache_Data_Monoid_Generic_genericMempty_prime_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Generic_genericMempty_prime_(dict_0_box)
		})
	})
	return cache_Data_Monoid_Generic_genericMempty_prime_
}

var cache_Data_Monoid_Generic_genericMonoidConstructor gopurs_runtime.Value
var once_Data_Monoid_Generic_genericMonoidConstructor sync.Once

func Get_Data_Monoid_Generic_genericMonoidConstructor() gopurs_runtime.Value {
	once_Data_Monoid_Generic_genericMonoidConstructor.Do(func() {
		cache_Data_Monoid_Generic_genericMonoidConstructor = gopurs_runtime.Func(func(dictGenericMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Generic_genericMonoidConstructor(dictGenericMonoid_0_box)
		})
	})
	return cache_Data_Monoid_Generic_genericMonoidConstructor
}

var cache_Data_Monoid_Generic_genericMonoidProduct gopurs_runtime.Value
var once_Data_Monoid_Generic_genericMonoidProduct sync.Once

func Get_Data_Monoid_Generic_genericMonoidProduct() gopurs_runtime.Value {
	once_Data_Monoid_Generic_genericMonoidProduct.Do(func() {
		cache_Data_Monoid_Generic_genericMonoidProduct = gopurs_runtime.Func(func(dictGenericMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Generic_genericMonoidProduct(dictGenericMonoid_0_box)
		})
	})
	return cache_Data_Monoid_Generic_genericMonoidProduct
}

var cache_Data_Monoid_Generic_genericMempty gopurs_runtime.Value
var once_Data_Monoid_Generic_genericMempty sync.Once

func Get_Data_Monoid_Generic_genericMempty() gopurs_runtime.Value {
	once_Data_Monoid_Generic_genericMempty.Do(func() {
		cache_Data_Monoid_Generic_genericMempty = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Monoid_Generic_genericMempty(dictGeneric_0_box, dictGenericMonoid_1_box)
		})
	})
	return cache_Data_Monoid_Generic_genericMempty
}

type Constructor_Data_Monoid_Generic_GenericMonoid[T_a any] struct {
	Rc uint32
	V0 T_a
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f004d006f006e006f00690064005f00470065006e0065007200690063:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f004d006f006e006f00690064005f00470065006e0065007200690063005f00470065006e0065007200690063004d006f006e006f00690064
func init() {
	gopurs_runtime.StructGetters[2569012965] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Monoid_Generic_GenericMonoid[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "genericMempty'":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Data_Monoid_Generic_GenericMonoid: " + key)
		}
	}
}

func Call_Data_Monoid_Generic_GenericMonoid_dollar_Dict(x_0_loop struct {
	gopurs_field_103_101_110_101_114_105_99_77_101_109_112_116_121_39 gopurs_runtime.Value
	/*gopurs:record-v1:1:00670065006e0065007200690063004d0065006d0070007400790027*/
}) *Constructor_Data_Monoid_Generic_GenericMonoid[gopurs_runtime.Value] {
	var x_0 struct {
		gopurs_field_103_101_110_101_114_105_99_77_101_109_112_116_121_39 gopurs_runtime.Value
		/*gopurs:record-v1:1:00670065006e0065007200690063004d0065006d0070007400790027*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Generic_GenericMonoid[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("genericMempty'", orig.gopurs_field_103_101_110_101_114_105_99_77_101_109_112_116_121_39)
	}())
}

func Call_Data_Monoid_Generic_GenericMonoid_dollar_Dict__121657880(x_0_loop struct {
	gopurs_field_103_101_110_101_114_105_99_77_101_109_112_116_121_39 uint32
	/*gopurs:record-v1:1:00670065006e0065007200690063004d0065006d0070007400790027*/
}) *Constructor_Data_Monoid_Generic_GenericMonoid[uint32] {
GenericMonoid_dollar_Dict__121657880:
	for {
		if false {
			continue GenericMonoid_dollar_Dict__121657880
		}
		var x_0 struct {
			gopurs_field_103_101_110_101_114_105_99_77_101_109_112_116_121_39 uint32
			/*gopurs:record-v1:1:00670065006e0065007200690063004d0065006d0070007400790027*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Generic_GenericMonoid[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict1("genericMempty'", gopurs_runtime.Value{Type: 9, IntVal: int64(orig.gopurs_field_103_101_110_101_114_105_99_77_101_109_112_116_121_39), UnsafePtr: nil})
		}())
	}
}

func Call_Data_Monoid_Generic_genericMonoidArgument(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
	_ = dictMonoid_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2569012965, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Generic_GenericMonoid[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}))}
}

func Call_Data_Monoid_Generic_genericMempty_prime_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "genericMempty'")
}

func Call_Data_Monoid_Generic_genericMonoidConstructor(dictGenericMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictGenericMonoid_0 gopurs_runtime.Value = dictGenericMonoid_0_loop
	_ = dictGenericMonoid_0
	return gopurs_runtime.Value{Type: 9, IntVal: 2569012965, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Generic_GenericMonoid[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericMonoid_0, "genericMempty'")}))}
}

func Call_Data_Monoid_Generic_genericMonoidProduct(dictGenericMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictGenericMonoid_0 gopurs_runtime.Value = dictGenericMonoid_0_loop
	_ = dictGenericMonoid_0
	// TAST (Let): genericMempty_prime_1_1_0 shape=Other bindingType=(TypeVar a$scope15)
	genericMempty_prime_1_1_0 := gopurs_runtime.RecordGet(dictGenericMonoid_0, "genericMempty'")
	_ = genericMempty_prime_1_1_0
	return gopurs_runtime.Func(func(dictGenericMonoid1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 2569012965, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_Generic_2471368446_4174781905((&Constructor_Data_Monoid_Generic_GenericMonoid[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]{1, (&Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, genericMempty_prime_1_1_0, gopurs_runtime.RecordGet(dictGenericMonoid1_2, "genericMempty'")})})))}
	})
}

func Call_Data_Monoid_Generic_genericMempty(dictGeneric_0_loop gopurs_runtime.Value, dictGenericMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
	_ = dictGeneric_0
	var dictGenericMonoid_1 gopurs_runtime.Value = dictGenericMonoid_1_loop
	_ = dictGenericMonoid_1
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericMonoid_1, "genericMempty'"))
}

func Rebox_Data_Monoid_Generic_2471368446_4174781905(in *Constructor_Data_Monoid_Generic_GenericMonoid[*Constructor_Data_Generic_Rep_Product[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Monoid_Generic_GenericMonoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Generic_GenericMonoid[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Monoid_Generic_3734302396_4174781905(in *Constructor_Data_Monoid_Generic_GenericMonoid[uint32]) *Constructor_Data_Monoid_Generic_GenericMonoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Generic_GenericMonoid[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}
