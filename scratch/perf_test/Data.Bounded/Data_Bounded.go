package Data_Bounded

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var topRecord gopurs_runtime.Value
var once_topRecord sync.Once
func Get_topRecord() gopurs_runtime.Value {
	once_topRecord.Do(func() {
		topRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["topRecord"]
})
	})
	return topRecord
}

var top gopurs_runtime.Value
var once_top sync.Once
func Get_top() gopurs_runtime.Value {
	once_top.Do(func() {
		top = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["top"]
})
	})
	return top
}

var boundedUnit gopurs_runtime.Value
var once_boundedUnit sync.Once
func Get_boundedUnit() gopurs_runtime.Value {
	once_boundedUnit.Do(func() {
		boundedUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"top": pkg_Data_Unit.Get_unit(), "bottom": pkg_Data_Unit.Get_unit(), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordUnit()
})})
	})
	return boundedUnit
}

var boundedRecordNil gopurs_runtime.Value
var once_boundedRecordNil sync.Once
func Get_boundedRecordNil() gopurs_runtime.Value {
	once_boundedRecordNil.Do(func() {
		boundedRecordNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"topRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
}), "bottomRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
}), "OrdRecord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordRecordNil()
})})
	})
	return boundedRecordNil
}

var boundedProxy gopurs_runtime.Value
var once_boundedProxy sync.Once
func Get_boundedProxy() gopurs_runtime.Value {
	once_boundedProxy.Do(func() {
		boundedProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bottom": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}), "top": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordProxy()
})})
	})
	return boundedProxy
}

var boundedOrdering gopurs_runtime.Value
var once_boundedOrdering sync.Once
func Get_boundedOrdering() gopurs_runtime.Value {
	once_boundedOrdering.Do(func() {
		boundedOrdering = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"top": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")}), "bottom": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordOrdering()
})})
	})
	return boundedOrdering
}

var boundedNumber gopurs_runtime.Value
var once_boundedNumber sync.Once
func Get_boundedNumber() gopurs_runtime.Value {
	once_boundedNumber.Do(func() {
		boundedNumber = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"top": Get_topNumber(), "bottom": Get_bottomNumber(), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordNumber()
})})
	})
	return boundedNumber
}

var boundedInt gopurs_runtime.Value
var once_boundedInt sync.Once
func Get_boundedInt() gopurs_runtime.Value {
	once_boundedInt.Do(func() {
		boundedInt = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"top": Get_topInt(), "bottom": Get_bottomInt(), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
})})
	})
	return boundedInt
}

var boundedChar gopurs_runtime.Value
var once_boundedChar sync.Once
func Get_boundedChar() gopurs_runtime.Value {
	once_boundedChar.Do(func() {
		boundedChar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"top": Get_topChar(), "bottom": Get_bottomChar(), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordChar()
})})
	})
	return boundedChar
}

var boundedBoolean gopurs_runtime.Value
var once_boundedBoolean sync.Once
func Get_boundedBoolean() gopurs_runtime.Value {
	once_boundedBoolean.Do(func() {
		boundedBoolean = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"top": gopurs_runtime.Bool(true), "bottom": gopurs_runtime.Bool(false), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordBoolean()
})})
	})
	return boundedBoolean
}

var bottomRecord gopurs_runtime.Value
var once_bottomRecord sync.Once
func Get_bottomRecord() gopurs_runtime.Value {
	once_bottomRecord.Do(func() {
		bottomRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["bottomRecord"]
})
	})
	return bottomRecord
}

var boundedRecord gopurs_runtime.Value
var once_boundedRecord sync.Once
func Get_boundedRecord() gopurs_runtime.Value {
	once_boundedRecord.Do(func() {
		boundedRecord = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictBoundedRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictBoundedRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["OrdRecord0"], gopurs_runtime.Value{})
eqRec1_3_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["EqRecord0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["eqRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))})
ordRecord1_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["compareRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "Eq0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRec1_3_1
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"top": gopurs_runtime.Apply(gopurs_runtime.Apply(dictBoundedRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["topRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "bottom": gopurs_runtime.Apply(gopurs_runtime.Apply(dictBoundedRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["bottomRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "Ord0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return ordRecord1_4_2
})})
})
})
	})
	return boundedRecord
}

var bottom gopurs_runtime.Value
var once_bottom sync.Once
func Get_bottom() gopurs_runtime.Value {
	once_bottom.Do(func() {
		bottom = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["bottom"]
})
	})
	return bottom
}

var boundedRecordCons gopurs_runtime.Value
var once_boundedRecordCons sync.Once
func Get_boundedRecordCons() gopurs_runtime.Value {
	once_boundedRecordCons.Do(func() {
		boundedRecordCons = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictBounded_1 gopurs_runtime.Value) gopurs_runtime.Value {
top1_2_0 := dictBounded_1.PtrVal.(map[string]gopurs_runtime.Value)["top"]
bottom1_3_1 := dictBounded_1.PtrVal.(map[string]gopurs_runtime.Value)["bottom"]
Ord0_4_2 := gopurs_runtime.Apply(dictBounded_1.PtrVal.(map[string]gopurs_runtime.Value)["Ord0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictBoundedRecord_7 gopurs_runtime.Value) gopurs_runtime.Value {
ordRecordCons_8_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordRecordCons(), gopurs_runtime.Apply(dictBoundedRecord_7.PtrVal.(map[string]gopurs_runtime.Value)["OrdRecord0"], gopurs_runtime.Value{})), gopurs_runtime.Value{}), dictIsSymbol_0), Ord0_4_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"topRecord": gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rowProxy_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))), top1_2_0), gopurs_runtime.Apply(gopurs_runtime.Apply(dictBoundedRecord_7.PtrVal.(map[string]gopurs_runtime.Value)["topRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), rowProxy_10))
})
}), "bottomRecord": gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rowProxy_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))), bottom1_3_1), gopurs_runtime.Apply(gopurs_runtime.Apply(dictBoundedRecord_7.PtrVal.(map[string]gopurs_runtime.Value)["bottomRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), rowProxy_10))
})
}), "OrdRecord0": gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return ordRecordCons_8_3
})})
})
})
})
})
})
	})
	return boundedRecordCons
}

func Get_bottomChar() gopurs_runtime.Value {
	return BottomChar
}

func Get_bottomInt() gopurs_runtime.Value {
	return BottomInt
}

func Get_bottomNumber() gopurs_runtime.Value {
	return BottomNumber
}

func Get_topChar() gopurs_runtime.Value {
	return TopChar
}

func Get_topInt() gopurs_runtime.Value {
	return TopInt
}

func Get_topNumber() gopurs_runtime.Value {
	return TopNumber
}
