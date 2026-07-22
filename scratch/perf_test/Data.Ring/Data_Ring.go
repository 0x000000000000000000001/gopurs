package Data_Ring

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var subRecord gopurs_runtime.Value
var once_subRecord sync.Once
func Get_subRecord() gopurs_runtime.Value {
	once_subRecord.Do(func() {
		subRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["subRecord"]
})
	})
	return subRecord
}

var sub gopurs_runtime.Value
var once_sub sync.Once
func Get_sub() gopurs_runtime.Value {
	once_sub.Do(func() {
		sub = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["sub"]
})
	})
	return sub
}

var ringUnit gopurs_runtime.Value
var once_ringUnit sync.Once
func Get_ringUnit() gopurs_runtime.Value {
	once_ringUnit.Do(func() {
		ringUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"sub": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}), "Semiring0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semiring.Get_semiringUnit()
})})
	})
	return ringUnit
}

var ringRecordNil gopurs_runtime.Value
var once_ringRecordNil sync.Once
func Get_ringRecordNil() gopurs_runtime.Value {
	once_ringRecordNil.Do(func() {
		ringRecordNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"subRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
}), "SemiringRecord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semiring.Get_semiringRecordNil()
})})
	})
	return ringRecordNil
}

var ringRecordCons gopurs_runtime.Value
var once_ringRecordCons sync.Once
func Get_ringRecordCons() gopurs_runtime.Value {
	once_ringRecordCons.Do(func() {
		ringRecordCons = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictRingRecord_2 gopurs_runtime.Value) gopurs_runtime.Value {
semiringRecordCons1_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_semiringRecordCons(), dictIsSymbol_0), gopurs_runtime.Value{}), gopurs_runtime.Apply(dictRingRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["SemiringRecord0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictRing_4 gopurs_runtime.Value) gopurs_runtime.Value {
semiringRecordCons2_5_1 := gopurs_runtime.Apply(semiringRecordCons1_3_0, gopurs_runtime.Apply(dictRing_4.PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"subRecord": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_2 := gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
get_10_3 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_2)
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), key_9_2), gopurs_runtime.Apply(gopurs_runtime.Apply(dictRing_4.PtrVal.(map[string]gopurs_runtime.Value)["sub"], gopurs_runtime.Apply(get_10_3, ra_7)), gopurs_runtime.Apply(get_10_3, rb_8))), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictRingRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["subRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), ra_7), rb_8))
})
})
}), "SemiringRecord0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringRecordCons2_5_1
})})
})
})
})
})
	})
	return ringRecordCons
}

var ringRecord gopurs_runtime.Value
var once_ringRecord sync.Once
func Get_ringRecord() gopurs_runtime.Value {
	once_ringRecord.Do(func() {
		ringRecord = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictRingRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictRingRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["SemiringRecord0"], gopurs_runtime.Value{})
semiringRecord1_3_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"add": gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["addRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "mul": gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["mulRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "one": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["oneRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "zero": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["zeroRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"sub": gopurs_runtime.Apply(dictRingRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["subRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "Semiring0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringRecord1_3_1
})})
})
})
	})
	return ringRecord
}

var ringProxy gopurs_runtime.Value
var once_ringProxy sync.Once
func Get_ringProxy() gopurs_runtime.Value {
	once_ringProxy.Do(func() {
		ringProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"sub": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
}), "Semiring0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semiring.Get_semiringProxy()
})})
	})
	return ringProxy
}

var ringNumber gopurs_runtime.Value
var once_ringNumber sync.Once
func Get_ringNumber() gopurs_runtime.Value {
	once_ringNumber.Do(func() {
		ringNumber = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"sub": Get_numSub(), "Semiring0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semiring.Get_semiringNumber()
})})
	})
	return ringNumber
}

var ringInt gopurs_runtime.Value
var once_ringInt sync.Once
func Get_ringInt() gopurs_runtime.Value {
	once_ringInt.Do(func() {
		ringInt = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"sub": Get_intSub(), "Semiring0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semiring.Get_semiringInt()
})})
	})
	return ringInt
}

var ringFn gopurs_runtime.Value
var once_ringFn sync.Once
func Get_ringFn() gopurs_runtime.Value {
	once_ringFn.Do(func() {
		ringFn = gopurs_runtime.Func(func(dictRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictRing_0.PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{})
zero1_2_1 := __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["zero"]
one1_3_3 := __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["one"]
semiringFn_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"add": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["add"], gopurs_runtime.Apply(f_4, x_6)), gopurs_runtime.Apply(g_5, x_6))
})
})
}), "zero": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return zero1_2_1
}), "mul": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["mul"], gopurs_runtime.Apply(f_4, x_6)), gopurs_runtime.Apply(g_5, x_6))
})
})
}), "one": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return one1_3_3
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"sub": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictRing_0.PtrVal.(map[string]gopurs_runtime.Value)["sub"], gopurs_runtime.Apply(f_4, x_6)), gopurs_runtime.Apply(g_5, x_6))
})
})
}), "Semiring0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringFn_3_2
})})
})
	})
	return ringFn
}

var negate gopurs_runtime.Value
var once_negate sync.Once
func Get_negate() gopurs_runtime.Value {
	once_negate.Do(func() {
		negate = gopurs_runtime.Func(func(dictRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
zero_1_0 := gopurs_runtime.Apply(dictRing_0.PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["zero"]
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictRing_0.PtrVal.(map[string]gopurs_runtime.Value)["sub"], zero_1_0), a_2)
})
})
	})
	return negate
}

func Get_intSub() gopurs_runtime.Value {
	return IntSub
}

func Get_numSub() gopurs_runtime.Value {
	return NumSub
}
