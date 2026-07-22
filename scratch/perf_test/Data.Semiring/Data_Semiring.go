package Data_Semiring

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var zeroRecord gopurs_runtime.Value
var once_zeroRecord sync.Once
func Get_zeroRecord() gopurs_runtime.Value {
	once_zeroRecord.Do(func() {
		zeroRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["zeroRecord"]
})
	})
	return zeroRecord
}

var zero gopurs_runtime.Value
var once_zero sync.Once
func Get_zero() gopurs_runtime.Value {
	once_zero.Do(func() {
		zero = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["zero"]
})
	})
	return zero
}

var semiringUnit gopurs_runtime.Value
var once_semiringUnit sync.Once
func Get_semiringUnit() gopurs_runtime.Value {
	once_semiringUnit.Do(func() {
		semiringUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"add": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}), "zero": pkg_Data_Unit.Get_unit(), "mul": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}), "one": pkg_Data_Unit.Get_unit()})
	})
	return semiringUnit
}

var semiringRecordNil gopurs_runtime.Value
var once_semiringRecordNil sync.Once
func Get_semiringRecordNil() gopurs_runtime.Value {
	once_semiringRecordNil.Do(func() {
		semiringRecordNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"addRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
}), "mulRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
}), "oneRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
}), "zeroRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})})
	})
	return semiringRecordNil
}

var semiringProxy gopurs_runtime.Value
var once_semiringProxy sync.Once
func Get_semiringProxy() gopurs_runtime.Value {
	once_semiringProxy.Do(func() {
		semiringProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"add": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
}), "mul": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
}), "one": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}), "zero": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})})
	})
	return semiringProxy
}

var semiringNumber gopurs_runtime.Value
var once_semiringNumber sync.Once
func Get_semiringNumber() gopurs_runtime.Value {
	once_semiringNumber.Do(func() {
		semiringNumber = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"add": Get_numAdd(), "zero": gopurs_runtime.Float(0.0), "mul": Get_numMul(), "one": gopurs_runtime.Float(1.0)})
	})
	return semiringNumber
}

var semiringInt gopurs_runtime.Value
var once_semiringInt sync.Once
func Get_semiringInt() gopurs_runtime.Value {
	once_semiringInt.Do(func() {
		semiringInt = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"add": Get_intAdd(), "zero": gopurs_runtime.Int(0), "mul": Get_intMul(), "one": gopurs_runtime.Int(1)})
	})
	return semiringInt
}

var oneRecord gopurs_runtime.Value
var once_oneRecord sync.Once
func Get_oneRecord() gopurs_runtime.Value {
	once_oneRecord.Do(func() {
		oneRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["oneRecord"]
})
	})
	return oneRecord
}

var one gopurs_runtime.Value
var once_one sync.Once
func Get_one() gopurs_runtime.Value {
	once_one.Do(func() {
		one = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["one"]
})
	})
	return one
}

var mulRecord gopurs_runtime.Value
var once_mulRecord sync.Once
func Get_mulRecord() gopurs_runtime.Value {
	once_mulRecord.Do(func() {
		mulRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["mulRecord"]
})
	})
	return mulRecord
}

var mul gopurs_runtime.Value
var once_mul sync.Once
func Get_mul() gopurs_runtime.Value {
	once_mul.Do(func() {
		mul = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["mul"]
})
	})
	return mul
}

var addRecord gopurs_runtime.Value
var once_addRecord sync.Once
func Get_addRecord() gopurs_runtime.Value {
	once_addRecord.Do(func() {
		addRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["addRecord"]
})
	})
	return addRecord
}

var semiringRecord gopurs_runtime.Value
var once_semiringRecord sync.Once
func Get_semiringRecord() gopurs_runtime.Value {
	once_semiringRecord.Do(func() {
		semiringRecord = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemiringRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"add": gopurs_runtime.Apply(dictSemiringRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["addRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "mul": gopurs_runtime.Apply(dictSemiringRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["mulRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "one": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiringRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["oneRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "zero": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiringRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["zeroRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))})
})
})
	})
	return semiringRecord
}

var add gopurs_runtime.Value
var once_add sync.Once
func Get_add() gopurs_runtime.Value {
	once_add.Do(func() {
		add = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["add"]
})
	})
	return add
}

var semiringFn gopurs_runtime.Value
var once_semiringFn sync.Once
func Get_semiringFn() gopurs_runtime.Value {
	once_semiringFn.Do(func() {
		semiringFn = gopurs_runtime.Func(func(dictSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
zero1_1_0 := dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["zero"]
one1_2_1 := dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["one"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"add": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["add"], gopurs_runtime.Apply(f_3, x_5)), gopurs_runtime.Apply(g_4, x_5))
})
})
}), "zero": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return zero1_1_0
}), "mul": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_0.PtrVal.(map[string]gopurs_runtime.Value)["mul"], gopurs_runtime.Apply(f_3, x_5)), gopurs_runtime.Apply(g_4, x_5))
})
})
}), "one": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return one1_2_1
})})
})
	})
	return semiringFn
}

var semiringRecordCons gopurs_runtime.Value
var once_semiringRecordCons sync.Once
func Get_semiringRecordCons() gopurs_runtime.Value {
	once_semiringRecordCons.Do(func() {
		semiringRecordCons = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemiringRecord_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemiring_3 gopurs_runtime.Value) gopurs_runtime.Value {
one1_4_0 := dictSemiring_3.PtrVal.(map[string]gopurs_runtime.Value)["one"]
zero1_5_1 := dictSemiring_3.PtrVal.(map[string]gopurs_runtime.Value)["zero"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"addRecord": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_2 := gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
get_10_3 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_2)
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), key_9_2), gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_3.PtrVal.(map[string]gopurs_runtime.Value)["add"], gopurs_runtime.Apply(get_10_3, ra_7)), gopurs_runtime.Apply(get_10_3, rb_8))), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiringRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["addRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), ra_7), rb_8))
})
})
}), "mulRecord": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_4 := gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
get_10_5 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_4)
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), key_9_4), gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiring_3.PtrVal.(map[string]gopurs_runtime.Value)["mul"], gopurs_runtime.Apply(get_10_5, ra_7)), gopurs_runtime.Apply(get_10_5, rb_8))), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiringRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["mulRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), ra_7), rb_8))
})
})
}), "oneRecord": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))), one1_4_0), gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiringRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["oneRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})))
})
}), "zeroRecord": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))), zero1_5_1), gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemiringRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["zeroRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})))
})
})})
})
})
})
})
	})
	return semiringRecordCons
}

func Get_intAdd() gopurs_runtime.Value {
	return IntAdd
}

func Get_intMul() gopurs_runtime.Value {
	return IntMul
}

func Get_numAdd() gopurs_runtime.Value {
	return NumAdd
}

func Get_numMul() gopurs_runtime.Value {
	return NumMul
}
