package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Record_Builder_Builder gopurs_runtime.Value
var once_Record_Builder_Builder sync.Once

func Get_Record_Builder_Builder() gopurs_runtime.Value {
	once_Record_Builder_Builder.Do(func() {
		cache_Record_Builder_Builder = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_Builder_Builder(x_0_box)
		})
	})
	return cache_Record_Builder_Builder
}

var cache_Record_Builder_union gopurs_runtime.Value
var once_Record_Builder_union sync.Once

func Get_Record_Builder_union() gopurs_runtime.Value {
	once_Record_Builder_union.Do(func() {
		cache_Record_Builder_union = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, r1_1_box gopurs_runtime.Value, r2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_Builder_union(_dollar___unused_0_box, r1_1_box, r2_2_box)
		})
	})
	return cache_Record_Builder_union
}

var cache_Record_Builder_semigroupoidBuilder gopurs_runtime.Value
var once_Record_Builder_semigroupoidBuilder sync.Once

func Get_Record_Builder_semigroupoidBuilder() gopurs_runtime.Value {
	once_Record_Builder_semigroupoidBuilder.Do(func() {
		cache_Record_Builder_semigroupoidBuilder = gopurs_runtime.Value{Type: 9, IntVal: 350442445, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))}
	})
	return cache_Record_Builder_semigroupoidBuilder
}

var cache_Record_Builder_rename gopurs_runtime.Value
var once_Record_Builder_rename sync.Once

func Get_Record_Builder_rename() gopurs_runtime.Value {
	once_Record_Builder_rename.Do(func() {
		cache_Record_Builder_rename = gopurs_runtime.Func9(func(dictIsSymbol_0_box gopurs_runtime.Value, dictIsSymbol1_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, _dollar___unused_3_box gopurs_runtime.Value, _dollar___unused_4_box gopurs_runtime.Value, _dollar___unused_5_box gopurs_runtime.Value, l1_6_box gopurs_runtime.Value, l2_7_box gopurs_runtime.Value, r1_8_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_Builder_rename(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol1_1_box), _dollar___unused_2_box, _dollar___unused_3_box, _dollar___unused_4_box, _dollar___unused_5_box, uint32(l1_6_box.IntVal), uint32(l2_7_box.IntVal), r1_8_box)
		})
	})
	return cache_Record_Builder_rename
}

var cache_Record_Builder_nub gopurs_runtime.Value
var once_Record_Builder_nub sync.Once

func Get_Record_Builder_nub() gopurs_runtime.Value {
	once_Record_Builder_nub.Do(func() {
		cache_Record_Builder_nub = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_Builder_nub(_dollar___unused_0_box)
		})
	})
	return cache_Record_Builder_nub
}

var cache_Record_Builder_modify gopurs_runtime.Value
var once_Record_Builder_modify sync.Once

func Get_Record_Builder_modify() gopurs_runtime.Value {
	once_Record_Builder_modify.Do(func() {
		cache_Record_Builder_modify = gopurs_runtime.Func6(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, dictIsSymbol_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value, f_4_box gopurs_runtime.Value, r1_5_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_Builder_modify(_dollar___unused_0_box, _dollar___unused_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_2_box), uint32(l_3_box.IntVal), f_4_box, r1_5_box)
		})
	})
	return cache_Record_Builder_modify
}

var cache_Record_Builder_merge gopurs_runtime.Value
var once_Record_Builder_merge sync.Once

func Get_Record_Builder_merge() gopurs_runtime.Value {
	once_Record_Builder_merge.Do(func() {
		cache_Record_Builder_merge = gopurs_runtime.Func4(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, r1_2_box gopurs_runtime.Value, r2_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_Builder_merge(_dollar___unused_0_box, _dollar___unused_1_box, r1_2_box, r2_3_box)
		})
	})
	return cache_Record_Builder_merge
}

var cache_Record_Builder_insert gopurs_runtime.Value
var once_Record_Builder_insert sync.Once

func Get_Record_Builder_insert() gopurs_runtime.Value {
	once_Record_Builder_insert.Do(func() {
		cache_Record_Builder_insert = gopurs_runtime.Func6(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, dictIsSymbol_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value, a_4_box gopurs_runtime.Value, r1_5_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_Builder_insert(_dollar___unused_0_box, _dollar___unused_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_2_box), uint32(l_3_box.IntVal), a_4_box, r1_5_box)
		})
	})
	return cache_Record_Builder_insert
}

var cache_Record_Builder_disjointUnion gopurs_runtime.Value
var once_Record_Builder_disjointUnion sync.Once

func Get_Record_Builder_disjointUnion() gopurs_runtime.Value {
	once_Record_Builder_disjointUnion.Do(func() {
		cache_Record_Builder_disjointUnion = gopurs_runtime.Func4(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, r1_2_box gopurs_runtime.Value, r2_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_Builder_disjointUnion(_dollar___unused_0_box, _dollar___unused_1_box, r1_2_box, r2_3_box)
		})
	})
	return cache_Record_Builder_disjointUnion
}

var cache_Record_Builder_go__delete gopurs_runtime.Value
var once_Record_Builder_go__delete sync.Once

func Get_Record_Builder_go__delete() gopurs_runtime.Value {
	once_Record_Builder_go__delete.Do(func() {
		cache_Record_Builder_go__delete = gopurs_runtime.Func5(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value, r2_4_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_Builder_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_0_box), _dollar___unused_1_box, _dollar___unused_2_box, uint32(l_3_box.IntVal), r2_4_box)
		})
	})
	return cache_Record_Builder_go__delete
}

var cache_Record_Builder_categoryBuilder gopurs_runtime.Value
var once_Record_Builder_categoryBuilder sync.Once

func Get_Record_Builder_categoryBuilder() gopurs_runtime.Value {
	once_Record_Builder_categoryBuilder.Do(func() {
		cache_Record_Builder_categoryBuilder = gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}
	})
	return cache_Record_Builder_categoryBuilder
}

var cache_Record_Builder_build gopurs_runtime.Value
var once_Record_Builder_build sync.Once

func Get_Record_Builder_build() gopurs_runtime.Value {
	once_Record_Builder_build.Do(func() {
		cache_Record_Builder_build = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, r1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_Builder_build(v_0_box, r1_1_box)
		})
	})
	return cache_Record_Builder_build
}

var cache_Record_Builder_buildFromScratch gopurs_runtime.Value
var once_Record_Builder_buildFromScratch sync.Once

func Get_Record_Builder_buildFromScratch() gopurs_runtime.Value {
	once_Record_Builder_buildFromScratch.Do(func() {
		cache_Record_Builder_buildFromScratch = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_Builder_buildFromScratch(a_0_box)
		})
	})
	return cache_Record_Builder_buildFromScratch
}

var cache_Record_Builder_flip gopurs_runtime.Value
var once_Record_Builder_flip sync.Once

func Get_Record_Builder_flip() gopurs_runtime.Value {
	once_Record_Builder_flip.Do(func() {
		cache_Record_Builder_flip = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Record_Builder_flip(f_0_box, b_1_box, a_2_box)
		})
	})
	return cache_Record_Builder_flip
}

func Call_Record_Builder_Builder(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Record_Builder_union(_dollar___unused_0_loop gopurs_runtime.Value, r1_1_loop gopurs_runtime.Value, r2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var r1_1 gopurs_runtime.Value = r1_1_loop
	_ = r1_1
	var r2_2 gopurs_runtime.Value = r2_2_loop
	_ = r2_2
	return gopurs_runtime.UncurriedApp2(Get_Record_Unsafe_Union_unsafeUnionFn(), r1_1, r2_2)
}

func Call_Record_Builder_rename(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], dictIsSymbol1_1_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], _dollar___unused_2_loop gopurs_runtime.Value, _dollar___unused_3_loop gopurs_runtime.Value, _dollar___unused_4_loop gopurs_runtime.Value, _dollar___unused_5_loop gopurs_runtime.Value, l1_6_loop uint32, l2_7_loop uint32, r1_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
	var l1_6 uint32 = l1_6_loop
	_ = l1_6
	var l2_7 uint32 = l2_7_loop
	_ = l2_7
	var r1_8 gopurs_runtime.Value = r1_8_loop
	_ = r1_8
	return gopurs_runtime.Apply3(Get_Record_Builder_unsafeRename(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictIsSymbol_0.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l1_6), UnsafePtr: nil})), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictIsSymbol1_1.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l2_7), UnsafePtr: nil})), r1_8)
}

func Call_Record_Builder_nub(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return Get_Unsafe_Coerce_unsafeCoerce()
}

func Call_Record_Builder_modify(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, dictIsSymbol_2_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], l_3_loop uint32, f_4_loop gopurs_runtime.Value, r1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var dictIsSymbol_2 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol_2_loop
	_ = dictIsSymbol_2
	var l_3 uint32 = l_3_loop
	_ = l_3
	var f_4 gopurs_runtime.Value = f_4_loop
	_ = f_4
	var r1_5 gopurs_runtime.Value = r1_5_loop
	_ = r1_5
	return gopurs_runtime.Apply3(Get_Record_Builder_unsafeModify(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictIsSymbol_2.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l_3), UnsafePtr: nil})), f_4, r1_5)
}

func Call_Record_Builder_merge(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, r1_2_loop gopurs_runtime.Value, r2_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var r1_2 gopurs_runtime.Value = r1_2_loop
	_ = r1_2
	var r2_3 gopurs_runtime.Value = r2_3_loop
	_ = r2_3
	return gopurs_runtime.UncurriedApp2(Get_Record_Unsafe_Union_unsafeUnionFn(), r1_2, r2_3)
}

func Call_Record_Builder_insert(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, dictIsSymbol_2_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], l_3_loop uint32, a_4_loop gopurs_runtime.Value, r1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var dictIsSymbol_2 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol_2_loop
	_ = dictIsSymbol_2
	var l_3 uint32 = l_3_loop
	_ = l_3
	var a_4 gopurs_runtime.Value = a_4_loop
	_ = a_4
	var r1_5 gopurs_runtime.Value = r1_5_loop
	_ = r1_5
	return gopurs_runtime.Apply3(Get_Record_Builder_unsafeInsert(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictIsSymbol_2.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l_3), UnsafePtr: nil})), a_4, r1_5)
}

func Call_Record_Builder_disjointUnion(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, r1_2_loop gopurs_runtime.Value, r2_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var r1_2 gopurs_runtime.Value = r1_2_loop
	_ = r1_2
	var r2_3 gopurs_runtime.Value = r2_3_loop
	_ = r2_3
	return gopurs_runtime.UncurriedApp2(Get_Record_Unsafe_Union_unsafeUnionFn(), r1_2, r2_3)
}

func Call_Record_Builder_go__delete(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value, l_3_loop uint32, r2_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var l_3 uint32 = l_3_loop
	_ = l_3
	var r2_4 gopurs_runtime.Value = r2_4_loop
	_ = r2_4
	return gopurs_runtime.Apply2(Get_Record_Builder_unsafeDelete(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(dictIsSymbol_0.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l_3), UnsafePtr: nil})), r2_4)
}

func Call_Record_Builder_build(v_0_loop gopurs_runtime.Value, r1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var r1_1 gopurs_runtime.Value = r1_1_loop
	_ = r1_1
	return gopurs_runtime.Apply(v_0, gopurs_runtime.Apply(Get_Record_Builder_copyRecord(), r1_1))
}

func Call_Record_Builder_buildFromScratch(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	return Call_Record_Builder_build(a_0, func() gopurs_runtime.Value {
		orig := struct {

			/*gopurs:record-v1:0:*/
		}{}
		_ = orig
		return gopurs_runtime.RecordDict0()
	}())
}

func Call_Record_Builder_flip(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var b_1 gopurs_runtime.Value = b_1_loop
	_ = b_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	return gopurs_runtime.Apply2(f_0, a_2, gopurs_runtime.Apply(Get_Record_Builder_copyRecord(), b_1))
}

func Get_Record_Builder_copyRecord() gopurs_runtime.Value {
	return _Gopurs_Record_Builder_CopyRecord
}

func Get_Record_Builder_unsafeDelete() gopurs_runtime.Value {
	return _Gopurs_Record_Builder_UnsafeDelete
}

func Get_Record_Builder_unsafeInsert() gopurs_runtime.Value {
	return _Gopurs_Record_Builder_UnsafeInsert
}

func Get_Record_Builder_unsafeModify() gopurs_runtime.Value {
	return _Gopurs_Record_Builder_UnsafeModify
}

func Get_Record_Builder_unsafeRename() gopurs_runtime.Value {
	return _Gopurs_Record_Builder_UnsafeRename
}
