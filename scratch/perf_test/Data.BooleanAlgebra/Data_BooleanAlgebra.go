package Data_BooleanAlgebra

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
)

var booleanAlgebraUnit gopurs_runtime.Value
var once_booleanAlgebraUnit sync.Once
func Get_booleanAlgebraUnit() gopurs_runtime.Value {
	once_booleanAlgebraUnit.Do(func() {
		booleanAlgebraUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"HeytingAlgebra0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_HeytingAlgebra.Get_heytingAlgebraUnit()
})})
	})
	return booleanAlgebraUnit
}

var booleanAlgebraRecordNil gopurs_runtime.Value
var once_booleanAlgebraRecordNil sync.Once
func Get_booleanAlgebraRecordNil() gopurs_runtime.Value {
	once_booleanAlgebraRecordNil.Do(func() {
		booleanAlgebraRecordNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"HeytingAlgebraRecord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_HeytingAlgebra.Get_heytingAlgebraRecordNil()
})})
	})
	return booleanAlgebraRecordNil
}

var booleanAlgebraRecordCons gopurs_runtime.Value
var once_booleanAlgebraRecordCons sync.Once
func Get_booleanAlgebraRecordCons() gopurs_runtime.Value {
	once_booleanAlgebraRecordCons.Do(func() {
		booleanAlgebraRecordCons = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictBooleanAlgebraRecord_2 gopurs_runtime.Value) gopurs_runtime.Value {
heytingAlgebraRecordCons1_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_heytingAlgebraRecordCons(), dictIsSymbol_0), gopurs_runtime.Value{}), gopurs_runtime.Apply(dictBooleanAlgebraRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["HeytingAlgebraRecord0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictBooleanAlgebra_4 gopurs_runtime.Value) gopurs_runtime.Value {
heytingAlgebraRecordCons2_5_1 := gopurs_runtime.Apply(heytingAlgebraRecordCons1_3_0, gopurs_runtime.Apply(dictBooleanAlgebra_4.PtrVal.(map[string]gopurs_runtime.Value)["HeytingAlgebra0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"HeytingAlgebraRecord0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraRecordCons2_5_1
})})
})
})
})
})
	})
	return booleanAlgebraRecordCons
}

var booleanAlgebraRecord gopurs_runtime.Value
var once_booleanAlgebraRecord sync.Once
func Get_booleanAlgebraRecord() gopurs_runtime.Value {
	once_booleanAlgebraRecord.Do(func() {
		booleanAlgebraRecord = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictBooleanAlgebraRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictBooleanAlgebraRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["HeytingAlgebraRecord0"], gopurs_runtime.Value{})
heytingAlgebraRecord1_3_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ff": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["ffRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "tt": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["ttRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "conj": gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["conjRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "disj": gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["disjRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "implies": gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["impliesRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "not": gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["notRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"HeytingAlgebra0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraRecord1_3_1
})})
})
})
	})
	return booleanAlgebraRecord
}

var booleanAlgebraProxy gopurs_runtime.Value
var once_booleanAlgebraProxy sync.Once
func Get_booleanAlgebraProxy() gopurs_runtime.Value {
	once_booleanAlgebraProxy.Do(func() {
		booleanAlgebraProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"HeytingAlgebra0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_HeytingAlgebra.Get_heytingAlgebraProxy()
})})
	})
	return booleanAlgebraProxy
}

var booleanAlgebraFn gopurs_runtime.Value
var once_booleanAlgebraFn sync.Once
func Get_booleanAlgebraFn() gopurs_runtime.Value {
	once_booleanAlgebraFn.Do(func() {
		booleanAlgebraFn = gopurs_runtime.Func(func(dictBooleanAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictBooleanAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["HeytingAlgebra0"], gopurs_runtime.Value{})
ff1_2_1 := __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["ff"]
tt1_3_3 := __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["tt"]
heytingAlgebraFunction_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ff": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ff1_2_1
}), "tt": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return tt1_3_3
}), "implies": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["implies"], gopurs_runtime.Apply(f_4, a_6)), gopurs_runtime.Apply(g_5, a_6))
})
})
}), "conj": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["conj"], gopurs_runtime.Apply(f_4, a_6)), gopurs_runtime.Apply(g_5, a_6))
})
})
}), "disj": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["disj"], gopurs_runtime.Apply(f_4, a_6)), gopurs_runtime.Apply(g_5, a_6))
})
})
}), "not": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["not"], gopurs_runtime.Apply(f_4, a_5))
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"HeytingAlgebra0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraFunction_3_2
})})
})
	})
	return booleanAlgebraFn
}

var booleanAlgebraBoolean gopurs_runtime.Value
var once_booleanAlgebraBoolean sync.Once
func Get_booleanAlgebraBoolean() gopurs_runtime.Value {
	once_booleanAlgebraBoolean.Do(func() {
		booleanAlgebraBoolean = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"HeytingAlgebra0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean()
})})
	})
	return booleanAlgebraBoolean
}


