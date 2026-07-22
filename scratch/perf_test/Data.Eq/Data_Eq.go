package Data_Eq

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var eqVoid gopurs_runtime.Value
var once_eqVoid sync.Once
func Get_eqVoid() gopurs_runtime.Value {
	once_eqVoid.Do(func() {
		eqVoid = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
})})
	})
	return eqVoid
}

var eqUnit gopurs_runtime.Value
var once_eqUnit sync.Once
func Get_eqUnit() gopurs_runtime.Value {
	once_eqUnit.Do(func() {
		eqUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
})})
	})
	return eqUnit
}

var eqString gopurs_runtime.Value
var once_eqString sync.Once
func Get_eqString() gopurs_runtime.Value {
	once_eqString.Do(func() {
		eqString = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": Get_eqStringImpl()})
	})
	return eqString
}

var eqRowNil gopurs_runtime.Value
var once_eqRowNil sync.Once
func Get_eqRowNil() gopurs_runtime.Value {
	once_eqRowNil.Do(func() {
		eqRowNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eqRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
})
})})
	})
	return eqRowNil
}

var eqRecord gopurs_runtime.Value
var once_eqRecord sync.Once
func Get_eqRecord() gopurs_runtime.Value {
	once_eqRecord.Do(func() {
		eqRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["eqRecord"]
})
	})
	return eqRecord
}

var eqRec gopurs_runtime.Value
var once_eqRec sync.Once
func Get_eqRec() gopurs_runtime.Value {
	once_eqRec.Do(func() {
		eqRec = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEqRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Apply(dictEqRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["eqRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))})
})
})
	})
	return eqRec
}

var eqProxy gopurs_runtime.Value
var once_eqProxy sync.Once
func Get_eqProxy() gopurs_runtime.Value {
	once_eqProxy.Do(func() {
		eqProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
})})
	})
	return eqProxy
}

var eqNumber gopurs_runtime.Value
var once_eqNumber sync.Once
func Get_eqNumber() gopurs_runtime.Value {
	once_eqNumber.Do(func() {
		eqNumber = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": Get_eqNumberImpl()})
	})
	return eqNumber
}

var eqInt gopurs_runtime.Value
var once_eqInt sync.Once
func Get_eqInt() gopurs_runtime.Value {
	once_eqInt.Do(func() {
		eqInt = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": Get_eqIntImpl()})
	})
	return eqInt
}

var eqChar gopurs_runtime.Value
var once_eqChar sync.Once
func Get_eqChar() gopurs_runtime.Value {
	once_eqChar.Do(func() {
		eqChar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": Get_eqCharImpl()})
	})
	return eqChar
}

var eqBoolean gopurs_runtime.Value
var once_eqBoolean sync.Once
func Get_eqBoolean() gopurs_runtime.Value {
	once_eqBoolean.Do(func() {
		eqBoolean = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": Get_eqBooleanImpl()})
	})
	return eqBoolean
}

var eq1 gopurs_runtime.Value
var once_eq1 sync.Once
func Get_eq1() gopurs_runtime.Value {
	once_eq1.Do(func() {
		eq1 = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["eq1"]
})
	})
	return eq1
}

var eq gopurs_runtime.Value
var once_eq sync.Once
func Get_eq() gopurs_runtime.Value {
	once_eq.Do(func() {
		eq = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"]
})
	})
	return eq
}

var eqArray gopurs_runtime.Value
var once_eqArray sync.Once
func Get_eqArray() gopurs_runtime.Value {
	once_eqArray.Do(func() {
		eqArray = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Apply(Get_eqArrayImpl(), dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"])})
})
	})
	return eqArray
}

var eq1Array gopurs_runtime.Value
var once_eq1Array sync.Once
func Get_eq1Array() gopurs_runtime.Value {
	once_eq1Array.Do(func() {
		eq1Array = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_eqArrayImpl(), dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"])
})})
	})
	return eq1Array
}

var eqRowCons gopurs_runtime.Value
var once_eqRowCons sync.Once
func Get_eqRowCons() gopurs_runtime.Value {
	once_eqRowCons.Do(func() {
		eqRowCons = gopurs_runtime.Func(func(dictEqRecord_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictIsSymbol_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eqRecord": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
get_7_0 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Apply(dictIsSymbol_2.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})))
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_3.PtrVal.(map[string]gopurs_runtime.Value)["eq"], gopurs_runtime.Apply(get_7_0, ra_5)), gopurs_runtime.Apply(get_7_0, rb_6))), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictEqRecord_0.PtrVal.(map[string]gopurs_runtime.Value)["eqRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), ra_5), rb_6))
})
})
})})
})
})
})
})
	})
	return eqRowCons
}

var notEq gopurs_runtime.Value
var once_notEq sync.Once
func Get_notEq() gopurs_runtime.Value {
	once_notEq.Do(func() {
		notEq = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eqBooleanImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_1), y_2)), gopurs_runtime.Bool(false))
})
})
})
	})
	return notEq
}

var notEq1 gopurs_runtime.Value
var once_notEq1 sync.Once
func Get_notEq1() gopurs_runtime.Value {
	once_notEq1.Do(func() {
		notEq1 = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
eq12_2_0 := gopurs_runtime.Apply(dictEq1_0.PtrVal.(map[string]gopurs_runtime.Value)["eq1"], dictEq_1)
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eqBooleanImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(eq12_2_0, x_3), y_4)), gopurs_runtime.Bool(false))
})
})
})
})
	})
	return notEq1
}

func Get_eqArrayImpl() gopurs_runtime.Value {
	return EqArrayImpl
}

func Get_eqBooleanImpl() gopurs_runtime.Value {
	return EqBooleanImpl
}

func Get_eqCharImpl() gopurs_runtime.Value {
	return EqCharImpl
}

func Get_eqIntImpl() gopurs_runtime.Value {
	return EqIntImpl
}

func Get_eqNumberImpl() gopurs_runtime.Value {
	return EqNumberImpl
}

func Get_eqStringImpl() gopurs_runtime.Value {
	return EqStringImpl
}
