package Data_CommutativeRing

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ring "gopurs/output/Data.Ring"
)

var commutativeRingUnit gopurs_runtime.Value
var once_commutativeRingUnit sync.Once
func Get_commutativeRingUnit() gopurs_runtime.Value {
	once_commutativeRingUnit.Do(func() {
		commutativeRingUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Ring0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ring.Get_ringUnit()
})})
	})
	return commutativeRingUnit
}

var commutativeRingRecordNil gopurs_runtime.Value
var once_commutativeRingRecordNil sync.Once
func Get_commutativeRingRecordNil() gopurs_runtime.Value {
	once_commutativeRingRecordNil.Do(func() {
		commutativeRingRecordNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"RingRecord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ring.Get_ringRecordNil()
})})
	})
	return commutativeRingRecordNil
}

var commutativeRingRecordCons gopurs_runtime.Value
var once_commutativeRingRecordCons sync.Once
func Get_commutativeRingRecordCons() gopurs_runtime.Value {
	once_commutativeRingRecordCons.Do(func() {
		commutativeRingRecordCons = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictCommutativeRingRecord_2 gopurs_runtime.Value) gopurs_runtime.Value {
ringRecordCons1_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_ringRecordCons(), dictIsSymbol_0), gopurs_runtime.Value{}), gopurs_runtime.Apply(dictCommutativeRingRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["RingRecord0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictCommutativeRing_4 gopurs_runtime.Value) gopurs_runtime.Value {
ringRecordCons2_5_1 := gopurs_runtime.Apply(ringRecordCons1_3_0, gopurs_runtime.Apply(dictCommutativeRing_4.PtrVal.(map[string]gopurs_runtime.Value)["Ring0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"RingRecord0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return ringRecordCons2_5_1
})})
})
})
})
})
	})
	return commutativeRingRecordCons
}

var commutativeRingRecord gopurs_runtime.Value
var once_commutativeRingRecord sync.Once
func Get_commutativeRingRecord() gopurs_runtime.Value {
	once_commutativeRingRecord.Do(func() {
		commutativeRingRecord = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictCommutativeRingRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictCommutativeRingRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["RingRecord0"], gopurs_runtime.Value{})
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["SemiringRecord0"], gopurs_runtime.Value{})
semiringRecord1_4_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"add": gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["addRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "mul": gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["mulRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "one": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["oneRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "zero": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["zeroRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))})
ringRecord1_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"sub": gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["subRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "Semiring0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringRecord1_4_3
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Ring0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return ringRecord1_4_2
})})
})
})
	})
	return commutativeRingRecord
}

var commutativeRingProxy gopurs_runtime.Value
var once_commutativeRingProxy sync.Once
func Get_commutativeRingProxy() gopurs_runtime.Value {
	once_commutativeRingProxy.Do(func() {
		commutativeRingProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Ring0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ring.Get_ringProxy()
})})
	})
	return commutativeRingProxy
}

var commutativeRingNumber gopurs_runtime.Value
var once_commutativeRingNumber sync.Once
func Get_commutativeRingNumber() gopurs_runtime.Value {
	once_commutativeRingNumber.Do(func() {
		commutativeRingNumber = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Ring0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ring.Get_ringNumber()
})})
	})
	return commutativeRingNumber
}

var commutativeRingInt gopurs_runtime.Value
var once_commutativeRingInt sync.Once
func Get_commutativeRingInt() gopurs_runtime.Value {
	once_commutativeRingInt.Do(func() {
		commutativeRingInt = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Ring0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ring.Get_ringInt()
})})
	})
	return commutativeRingInt
}

var commutativeRingFn gopurs_runtime.Value
var once_commutativeRingFn sync.Once
func Get_commutativeRingFn() gopurs_runtime.Value {
	once_commutativeRingFn.Do(func() {
		commutativeRingFn = gopurs_runtime.Func(func(dictCommutativeRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictCommutativeRing_0.PtrVal.(map[string]gopurs_runtime.Value)["Ring0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{})
zero1_3_3 := __local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["zero"]
one1_4_4 := __local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["one"]
semiringFn_5_5 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"add": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["add"], gopurs_runtime.Apply(f_5, x_7)), gopurs_runtime.Apply(g_6, x_7))
})
})
}), "zero": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return zero1_3_3
}), "mul": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["mul"], gopurs_runtime.Apply(f_5, x_7)), gopurs_runtime.Apply(g_6, x_7))
})
})
}), "one": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return one1_4_4
})})
ringFn_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"sub": gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["sub"], gopurs_runtime.Apply(f_6, x_8)), gopurs_runtime.Apply(g_7, x_8))
})
})
}), "Semiring0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringFn_5_5
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Ring0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ringFn_3_2
})})
})
	})
	return commutativeRingFn
}


