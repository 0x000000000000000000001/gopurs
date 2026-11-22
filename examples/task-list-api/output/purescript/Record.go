package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Record_conj gopurs_runtime.Value
var once_Record_conj sync.Once

func Get_Record_conj() gopurs_runtime.Value {
	once_Record_conj.Do(func() {
		cache_Record_conj = Get_Data_HeytingAlgebra_conj__2421783737()
	})
	return cache_Record_conj
}

var cache_Record_EqualFields_dollar_Dict gopurs_runtime.Value
var once_Record_EqualFields_dollar_Dict sync.Once

func Get_Record_EqualFields_dollar_Dict() gopurs_runtime.Value {
	once_Record_EqualFields_dollar_Dict.Do(func() {
		cache_Record_EqualFields_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2299966053, UnsafePtr: unsafe.Pointer(Call_Record_EqualFields_dollar_Dict(func() struct {
				equalFields gopurs_runtime.Value
				/*gopurs:record-v1:1:0065007100750061006c004600690065006c00640073*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					equalFields gopurs_runtime.Value
					/*gopurs:record-v1:1:0065007100750061006c004600690065006c00640073*/
				}{}
				clone.equalFields = gopurs_runtime.RecordGet(orig, "equalFields")
				return clone
			}()))}
		})
	})
	return cache_Record_EqualFields_dollar_Dict
}

var cache_Record_union gopurs_runtime.Value
var once_Record_union sync.Once

func Get_Record_union() gopurs_runtime.Value {
	once_Record_union.Do(func() {
		cache_Record_union = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_union(_dollar___unused_0_box, l_1_box, r_2_box)
		})
	})
	return cache_Record_union
}

var cache_Record_set gopurs_runtime.Value
var once_Record_set sync.Once

func Get_Record_set() gopurs_runtime.Value {
	once_Record_set.Do(func() {
		cache_Record_set = gopurs_runtime.Func6(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value, b_4_box gopurs_runtime.Value, r_5_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_set(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_0_box), _dollar___unused_1_box, _dollar___unused_2_box, uint32(l_3_box.IntVal), b_4_box, r_5_box)
		})
	})
	return cache_Record_set
}

var cache_Record_nub gopurs_runtime.Value
var once_Record_nub sync.Once

func Get_Record_nub() gopurs_runtime.Value {
	once_Record_nub.Do(func() {
		cache_Record_nub = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_nub(_dollar___unused_0_box)
		})
	})
	return cache_Record_nub
}

var cache_Record_merge gopurs_runtime.Value
var once_Record_merge sync.Once

func Get_Record_merge() gopurs_runtime.Value {
	once_Record_merge.Do(func() {
		cache_Record_merge = gopurs_runtime.Func4(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_merge(_dollar___unused_0_box, _dollar___unused_1_box, l_2_box, r_3_box)
		})
	})
	return cache_Record_merge
}

var cache_Record_insert gopurs_runtime.Value
var once_Record_insert sync.Once

func Get_Record_insert() gopurs_runtime.Value {
	once_Record_insert.Do(func() {
		cache_Record_insert = gopurs_runtime.Func6(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value, a_4_box gopurs_runtime.Value, r_5_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_0_box), _dollar___unused_1_box, _dollar___unused_2_box, uint32(l_3_box.IntVal), a_4_box, r_5_box)
		})
	})
	return cache_Record_insert
}

var cache_Record_get gopurs_runtime.Value
var once_Record_get sync.Once

func Get_Record_get() gopurs_runtime.Value {
	once_Record_get.Do(func() {
		cache_Record_get = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_get(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_0_box), _dollar___unused_1_box, uint32(l_2_box.IntVal), r_3_box)
		})
	})
	return cache_Record_get
}

var cache_Record_modify gopurs_runtime.Value
var once_Record_modify sync.Once

func Get_Record_modify() gopurs_runtime.Value {
	once_Record_modify.Do(func() {
		cache_Record_modify = gopurs_runtime.Func6(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value, f_4_box gopurs_runtime.Value, r_5_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_modify(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_0_box), _dollar___unused_1_box, _dollar___unused_2_box, uint32(l_3_box.IntVal), f_4_box, r_5_box)
		})
	})
	return cache_Record_modify
}

var cache_Record_equalFieldsNil gopurs_runtime.Value
var once_Record_equalFieldsNil sync.Once

func Get_Record_equalFieldsNil() gopurs_runtime.Value {
	once_Record_equalFieldsNil.Do(func() {
		cache_Record_equalFieldsNil = gopurs_runtime.Value{Type: 9, IntVal: 2299966053, UnsafePtr: unsafe.Pointer((&Constructor_Record_EqualFields[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})}))}
	})
	return cache_Record_equalFieldsNil
}

var cache_Record_equalFields gopurs_runtime.Value
var once_Record_equalFields sync.Once

func Get_Record_equalFields() gopurs_runtime.Value {
	once_Record_equalFields.Do(func() {
		cache_Record_equalFields = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_equalFields(gopurs_runtime.CoerceToStruct[Constructor_Record_EqualFields[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Record_equalFields
}

var cache_Record_equalFieldsCons gopurs_runtime.Value
var once_Record_equalFieldsCons sync.Once

func Get_Record_equalFieldsCons() gopurs_runtime.Value {
	once_Record_equalFieldsCons.Do(func() {
		cache_Record_equalFieldsCons = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, dictEqualFields_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_equalFieldsCons(dictIsSymbol_0_box, dictEq_1_box, _dollar___unused_2_box, dictEqualFields_3_box)
		})
	})
	return cache_Record_equalFieldsCons
}

var cache_Record_equal gopurs_runtime.Value
var once_Record_equal sync.Once

func Get_Record_equal() gopurs_runtime.Value {
	once_Record_equal.Do(func() {
		cache_Record_equal = gopurs_runtime.Func4(func(_dollar___unused_0_box gopurs_runtime.Value, dictEqualFields_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Record_equal(_dollar___unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Record_EqualFields[gopurs_runtime.Value, gopurs_runtime.Value]](dictEqualFields_1_box), a_2_box, b_3_box))
		})
	})
	return cache_Record_equal
}

var cache_Record_disjointUnion gopurs_runtime.Value
var once_Record_disjointUnion sync.Once

func Get_Record_disjointUnion() gopurs_runtime.Value {
	once_Record_disjointUnion.Do(func() {
		cache_Record_disjointUnion = gopurs_runtime.Func4(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_disjointUnion(_dollar___unused_0_box, _dollar___unused_1_box, l_2_box, r_3_box)
		})
	})
	return cache_Record_disjointUnion
}

var cache_Record_go__delete gopurs_runtime.Value
var once_Record_go__delete sync.Once

func Get_Record_go__delete() gopurs_runtime.Value {
	once_Record_go__delete.Do(func() {
		cache_Record_go__delete = gopurs_runtime.Func5(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value, r_4_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_0_box), _dollar___unused_1_box, _dollar___unused_2_box, uint32(l_3_box.IntVal), r_4_box)
		})
	})
	return cache_Record_go__delete
}

var cache_Record_rename gopurs_runtime.Value
var once_Record_rename sync.Once

func Get_Record_rename() gopurs_runtime.Value {
	once_Record_rename.Do(func() {
		cache_Record_rename = gopurs_runtime.Func9(func(dictIsSymbol_0_box gopurs_runtime.Value, dictIsSymbol1_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, _dollar___unused_3_box gopurs_runtime.Value, _dollar___unused_4_box gopurs_runtime.Value, _dollar___unused_5_box gopurs_runtime.Value, prev_6_box gopurs_runtime.Value, next_7_box gopurs_runtime.Value, record_8_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_rename(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol1_1_box), _dollar___unused_2_box, _dollar___unused_3_box, _dollar___unused_4_box, _dollar___unused_5_box, uint32(prev_6_box.IntVal), uint32(next_7_box.IntVal), record_8_box)
		})
	})
	return cache_Record_rename
}

type Constructor_Record_EqualFields[T_rs any, T_row any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f005200650063006f00720064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f005200650063006f00720064005f0045007100750061006c004600690065006c00640073
func init() {
	gopurs_runtime.StructGetters[2299966053] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Record_EqualFields[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "equalFields":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Record_EqualFields: " + key)
		}
	}
}

func Call_Record_EqualFields_dollar_Dict(x_0_loop struct {
	equalFields gopurs_runtime.Value
	/*gopurs:record-v1:1:0065007100750061006c004600690065006c00640073*/
}) *Constructor_Record_EqualFields[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		equalFields gopurs_runtime.Value
		/*gopurs:record-v1:1:0065007100750061006c004600690065006c00640073*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Record_EqualFields[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("equalFields", orig.equalFields)
	}())
}

func Call_Record_union(_dollar___unused_0_loop gopurs_runtime.Value, l_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var l_1 gopurs_runtime.Value = l_1_loop
	_ = l_1
	var r_2 gopurs_runtime.Value = r_2_loop
	_ = r_2
	return gopurs_runtime.UncurriedApp2(Get_Record_Unsafe_Union_unsafeUnionFn(), l_1, r_2)
}

func Call_Record_set(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value, l_3_loop uint32, b_4_loop gopurs_runtime.Value, r_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var l_3 uint32 = l_3_loop
	_ = l_3
	var b_4 gopurs_runtime.Value = b_4_loop
	_ = b_4
	var r_5 gopurs_runtime.Value = r_5_loop
	_ = r_5
	return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictIsSymbol_0.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l_3), UnsafePtr: nil})), b_4, r_5)
}

func Call_Record_nub(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Get_Unsafe_Coerce_unsafeCoerce()
}

func Call_Record_merge(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var l_2 gopurs_runtime.Value = l_2_loop
	_ = l_2
	var r_3 gopurs_runtime.Value = r_3_loop
	_ = r_3
	return gopurs_runtime.UncurriedApp2(Get_Record_Unsafe_Union_unsafeUnionFn(), l_2, r_3)
}

func Call_Record_insert(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value, l_3_loop uint32, a_4_loop gopurs_runtime.Value, r_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var l_3 uint32 = l_3_loop
	_ = l_3
	var a_4 gopurs_runtime.Value = a_4_loop
	_ = a_4
	var r_5 gopurs_runtime.Value = r_5_loop
	_ = r_5
	return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictIsSymbol_0.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l_3), UnsafePtr: nil})), a_4, r_5)
}

func Call_Record_get(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], _dollar___unused_1_loop gopurs_runtime.Value, l_2_loop uint32, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var l_2 uint32 = l_2_loop
	_ = l_2
	var r_3 gopurs_runtime.Value = r_3_loop
	_ = r_3
	return gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictIsSymbol_0.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l_2), UnsafePtr: nil})), r_3)
}

func Call_Record_modify(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value, l_3_loop uint32, f_4_loop gopurs_runtime.Value, r_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var l_3 uint32 = l_3_loop
	_ = l_3
	var f_4 gopurs_runtime.Value = f_4_loop
	_ = f_4
	var r_5 gopurs_runtime.Value = r_5_loop
	_ = r_5
	return Call_Record_set(dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Value{}, l_3, gopurs_runtime.Apply(f_4, gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictIsSymbol_0.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l_3), UnsafePtr: nil})), r_5)), r_5)
}

func Call_Record_equalFields(dict_0_loop *Constructor_Record_EqualFields[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Record_EqualFields[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Record_equalFieldsCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value, dictEqualFields_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
	_ = dictEq_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var dictEqualFields_3 gopurs_runtime.Value = dictEqualFields_3_loop
	_ = dictEqualFields_3
	return gopurs_runtime.Value{Type: 9, IntVal: 2299966053, UnsafePtr: unsafe.Pointer((&Constructor_Record_EqualFields[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})), a_5), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})), b_6)).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEqualFields_3, "equalFields"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, a_5, b_6).IntVal) != (0)))
	})}))}
}

func Call_Record_equal(_dollar___unused_0_loop gopurs_runtime.Value, dictEqualFields_1_loop *Constructor_Record_EqualFields[gopurs_runtime.Value, gopurs_runtime.Value], a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) bool {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var dictEqualFields_1 *Constructor_Record_EqualFields[gopurs_runtime.Value, gopurs_runtime.Value] = dictEqualFields_1_loop
	_ = dictEqualFields_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	var b_3 gopurs_runtime.Value = b_3_loop
	_ = b_3
	return (gopurs_runtime.Apply3(dictEqualFields_1.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, a_2, b_3).IntVal) != (0)
}

func Call_Record_disjointUnion(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var l_2 gopurs_runtime.Value = l_2_loop
	_ = l_2
	var r_3 gopurs_runtime.Value = r_3_loop
	_ = r_3
	return gopurs_runtime.UncurriedApp2(Get_Record_Unsafe_Union_unsafeUnionFn(), l_2, r_3)
}

func Call_Record_go__delete(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value, l_3_loop uint32, r_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var l_3 uint32 = l_3_loop
	_ = l_3
	var r_4 gopurs_runtime.Value = r_4_loop
	_ = r_4
	return gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeDelete(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictIsSymbol_0.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l_3), UnsafePtr: nil})), r_4)
}

func Call_Record_rename(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], dictIsSymbol1_1_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], _dollar___unused_2_loop gopurs_runtime.Value, _dollar___unused_3_loop gopurs_runtime.Value, _dollar___unused_4_loop gopurs_runtime.Value, _dollar___unused_5_loop gopurs_runtime.Value, prev_6_loop uint32, next_7_loop uint32, record_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var dictIsSymbol1_1 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol1_1_loop
	_ = dictIsSymbol1_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var _dollar___unused_3 gopurs_runtime.Value = _dollar___unused_3_loop
	_ = _dollar___unused_3
	var _dollar___unused_4 gopurs_runtime.Value = _dollar___unused_4_loop
	_ = _dollar___unused_4
	var _dollar___unused_5 gopurs_runtime.Value = _dollar___unused_5_loop
	_ = _dollar___unused_5
	var prev_6 uint32 = prev_6_loop
	_ = prev_6
	var next_7 uint32 = next_7_loop
	_ = next_7
	var record_8 gopurs_runtime.Value = record_8_loop
	_ = record_8
	return Call_Record_insert(dictIsSymbol1_1, gopurs_runtime.Value{}, gopurs_runtime.Value{}, next_7, gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictIsSymbol_0.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(prev_6), UnsafePtr: nil})), record_8), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeDelete(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictIsSymbol_0.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(prev_6), UnsafePtr: nil})), record_8))
}
