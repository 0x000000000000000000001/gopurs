package Data_HeytingAlgebra

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var ttRecord gopurs_runtime.Value
var once_ttRecord sync.Once
func Get_ttRecord() gopurs_runtime.Value {
	once_ttRecord.Do(func() {
		ttRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["ttRecord"]
})
	})
	return ttRecord
}

var tt gopurs_runtime.Value
var once_tt sync.Once
func Get_tt() gopurs_runtime.Value {
	once_tt.Do(func() {
		tt = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["tt"]
})
	})
	return tt
}

var notRecord gopurs_runtime.Value
var once_notRecord sync.Once
func Get_notRecord() gopurs_runtime.Value {
	once_notRecord.Do(func() {
		notRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["notRecord"]
})
	})
	return notRecord
}

var not gopurs_runtime.Value
var once_not sync.Once
func Get_not() gopurs_runtime.Value {
	once_not.Do(func() {
		not = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["not"]
})
	})
	return not
}

var impliesRecord gopurs_runtime.Value
var once_impliesRecord sync.Once
func Get_impliesRecord() gopurs_runtime.Value {
	once_impliesRecord.Do(func() {
		impliesRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["impliesRecord"]
})
	})
	return impliesRecord
}

var implies gopurs_runtime.Value
var once_implies sync.Once
func Get_implies() gopurs_runtime.Value {
	once_implies.Do(func() {
		implies = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["implies"]
})
	})
	return implies
}

var heytingAlgebraUnit gopurs_runtime.Value
var once_heytingAlgebraUnit sync.Once
func Get_heytingAlgebraUnit() gopurs_runtime.Value {
	once_heytingAlgebraUnit.Do(func() {
		heytingAlgebraUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ff": pkg_Data_Unit.Get_unit(), "tt": pkg_Data_Unit.Get_unit(), "implies": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}), "conj": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}), "disj": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}), "not": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})})
	})
	return heytingAlgebraUnit
}

var heytingAlgebraRecordNil gopurs_runtime.Value
var once_heytingAlgebraRecordNil sync.Once
func Get_heytingAlgebraRecordNil() gopurs_runtime.Value {
	once_heytingAlgebraRecordNil.Do(func() {
		heytingAlgebraRecordNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"conjRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
}), "disjRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
}), "ffRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
}), "impliesRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
}), "notRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
}), "ttRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})})
	})
	return heytingAlgebraRecordNil
}

var heytingAlgebraProxy gopurs_runtime.Value
var once_heytingAlgebraProxy sync.Once
func Get_heytingAlgebraProxy() gopurs_runtime.Value {
	once_heytingAlgebraProxy.Do(func() {
		heytingAlgebraProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"conj": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
}), "disj": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
}), "implies": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
}), "ff": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}), "not": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
}), "tt": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})})
	})
	return heytingAlgebraProxy
}

var ffRecord gopurs_runtime.Value
var once_ffRecord sync.Once
func Get_ffRecord() gopurs_runtime.Value {
	once_ffRecord.Do(func() {
		ffRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["ffRecord"]
})
	})
	return ffRecord
}

var ff gopurs_runtime.Value
var once_ff sync.Once
func Get_ff() gopurs_runtime.Value {
	once_ff.Do(func() {
		ff = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["ff"]
})
	})
	return ff
}

var disjRecord gopurs_runtime.Value
var once_disjRecord sync.Once
func Get_disjRecord() gopurs_runtime.Value {
	once_disjRecord.Do(func() {
		disjRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["disjRecord"]
})
	})
	return disjRecord
}

var disj gopurs_runtime.Value
var once_disj sync.Once
func Get_disj() gopurs_runtime.Value {
	once_disj.Do(func() {
		disj = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["disj"]
})
	})
	return disj
}

var heytingAlgebraBoolean gopurs_runtime.Value
var once_heytingAlgebraBoolean sync.Once
func Get_heytingAlgebraBoolean() gopurs_runtime.Value {
	once_heytingAlgebraBoolean.Do(func() {
		heytingAlgebraBoolean = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ff": gopurs_runtime.Bool(false), "tt": gopurs_runtime.Bool(true), "implies": gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_heytingAlgebraBoolean().PtrVal.(map[string]gopurs_runtime.Value)["disj"], gopurs_runtime.Apply(Get_heytingAlgebraBoolean().PtrVal.(map[string]gopurs_runtime.Value)["not"], a_0)), b_1)
})
}), "conj": Get_boolConj(), "disj": Get_boolDisj(), "not": Get_boolNot()})
	})
	return heytingAlgebraBoolean
}

var conjRecord gopurs_runtime.Value
var once_conjRecord sync.Once
func Get_conjRecord() gopurs_runtime.Value {
	once_conjRecord.Do(func() {
		conjRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["conjRecord"]
})
	})
	return conjRecord
}

var heytingAlgebraRecord gopurs_runtime.Value
var once_heytingAlgebraRecord sync.Once
func Get_heytingAlgebraRecord() gopurs_runtime.Value {
	once_heytingAlgebraRecord.Do(func() {
		heytingAlgebraRecord = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictHeytingAlgebraRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ff": gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebraRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["ffRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "tt": gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebraRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["ttRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "conj": gopurs_runtime.Apply(dictHeytingAlgebraRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["conjRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "disj": gopurs_runtime.Apply(dictHeytingAlgebraRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["disjRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "implies": gopurs_runtime.Apply(dictHeytingAlgebraRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["impliesRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "not": gopurs_runtime.Apply(dictHeytingAlgebraRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["notRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))})
})
})
	})
	return heytingAlgebraRecord
}

var conj gopurs_runtime.Value
var once_conj sync.Once
func Get_conj() gopurs_runtime.Value {
	once_conj.Do(func() {
		conj = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["conj"]
})
	})
	return conj
}

var heytingAlgebraFunction gopurs_runtime.Value
var once_heytingAlgebraFunction sync.Once
func Get_heytingAlgebraFunction() gopurs_runtime.Value {
	once_heytingAlgebraFunction.Do(func() {
		heytingAlgebraFunction = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
ff1_1_0 := dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["ff"]
tt1_2_1 := dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["tt"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ff": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return ff1_1_0
}), "tt": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return tt1_2_1
}), "implies": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["implies"], gopurs_runtime.Apply(f_3, a_5)), gopurs_runtime.Apply(g_4, a_5))
})
})
}), "conj": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["conj"], gopurs_runtime.Apply(f_3, a_5)), gopurs_runtime.Apply(g_4, a_5))
})
})
}), "disj": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["disj"], gopurs_runtime.Apply(f_3, a_5)), gopurs_runtime.Apply(g_4, a_5))
})
})
}), "not": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictHeytingAlgebra_0.PtrVal.(map[string]gopurs_runtime.Value)["not"], gopurs_runtime.Apply(f_3, a_4))
})
})})
})
	})
	return heytingAlgebraFunction
}

var heytingAlgebraRecordCons gopurs_runtime.Value
var once_heytingAlgebraRecordCons sync.Once
func Get_heytingAlgebraRecordCons() gopurs_runtime.Value {
	once_heytingAlgebraRecordCons.Do(func() {
		heytingAlgebraRecordCons = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictHeytingAlgebraRecord_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictHeytingAlgebra_3 gopurs_runtime.Value) gopurs_runtime.Value {
ff1_4_0 := dictHeytingAlgebra_3.PtrVal.(map[string]gopurs_runtime.Value)["ff"]
tt1_5_1 := dictHeytingAlgebra_3.PtrVal.(map[string]gopurs_runtime.Value)["tt"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"conjRecord": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_2 := gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
get_10_3 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_2)
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), key_9_2), gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_3.PtrVal.(map[string]gopurs_runtime.Value)["conj"], gopurs_runtime.Apply(get_10_3, ra_7)), gopurs_runtime.Apply(get_10_3, rb_8))), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebraRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["conjRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), ra_7), rb_8))
})
})
}), "disjRecord": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_4 := gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
get_10_5 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_4)
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), key_9_4), gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_3.PtrVal.(map[string]gopurs_runtime.Value)["disj"], gopurs_runtime.Apply(get_10_5, ra_7)), gopurs_runtime.Apply(get_10_5, rb_8))), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebraRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["disjRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), ra_7), rb_8))
})
})
}), "impliesRecord": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_6 := gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
get_10_7 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_6)
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), key_9_6), gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebra_3.PtrVal.(map[string]gopurs_runtime.Value)["implies"], gopurs_runtime.Apply(get_10_7, ra_7)), gopurs_runtime.Apply(get_10_7, rb_8))), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebraRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["impliesRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), ra_7), rb_8))
})
})
}), "ffRecord": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(row_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))), ff1_4_0), gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebraRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["ffRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), row_7))
})
}), "notRecord": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(row_7 gopurs_runtime.Value) gopurs_runtime.Value {
key_8_8 := gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), key_8_8), gopurs_runtime.Apply(dictHeytingAlgebra_3.PtrVal.(map[string]gopurs_runtime.Value)["not"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_8_8), row_7))), gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebraRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["notRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), row_7))
})
}), "ttRecord": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(row_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))), tt1_5_1), gopurs_runtime.Apply(gopurs_runtime.Apply(dictHeytingAlgebraRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["ttRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), row_7))
})
})})
})
})
})
})
	})
	return heytingAlgebraRecordCons
}

func Get_boolConj() gopurs_runtime.Value {
	return BoolConj
}

func Get_boolDisj() gopurs_runtime.Value {
	return BoolDisj
}

func Get_boolNot() gopurs_runtime.Value {
	return BoolNot
}
