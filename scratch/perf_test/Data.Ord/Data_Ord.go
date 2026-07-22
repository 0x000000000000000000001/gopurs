package Data_Ord

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var ordVoid gopurs_runtime.Value
var once_ordVoid sync.Once
func Get_ordVoid() gopurs_runtime.Value {
	once_ordVoid.Do(func() {
		ordVoid = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqVoid()
})})
	})
	return ordVoid
}

var ordUnit gopurs_runtime.Value
var once_ordUnit sync.Once
func Get_ordUnit() gopurs_runtime.Value {
	once_ordUnit.Do(func() {
		ordUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqUnit()
})})
	})
	return ordUnit
}

var ordString gopurs_runtime.Value
var once_ordString sync.Once
func Get_ordString() gopurs_runtime.Value {
	once_ordString.Do(func() {
		ordString = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_ordStringImpl(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqString()
})})
	})
	return ordString
}

var ordRecordNil gopurs_runtime.Value
var once_ordRecordNil sync.Once
func Get_ordRecordNil() gopurs_runtime.Value {
	once_ordRecordNil.Do(func() {
		ordRecordNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compareRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
})
})
}), "EqRecord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqRowNil()
})})
	})
	return ordRecordNil
}

var ordProxy gopurs_runtime.Value
var once_ordProxy sync.Once
func Get_ordProxy() gopurs_runtime.Value {
	once_ordProxy.Do(func() {
		ordProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqProxy()
})})
	})
	return ordProxy
}

var ordOrdering gopurs_runtime.Value
var once_ordOrdering sync.Once
func Get_ordOrdering() gopurs_runtime.Value {
	once_ordOrdering.Do(func() {
		ordOrdering = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
var __t3 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
} else {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
__t0 = __t3
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
} else {
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
} else {
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
__t0 = __t2
} else {
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
} else {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
}
__t0 = __t1
} else {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t0
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ordering.Get_eqOrdering()
})})
	})
	return ordOrdering
}

var ordNumber gopurs_runtime.Value
var once_ordNumber sync.Once
func Get_ordNumber() gopurs_runtime.Value {
	once_ordNumber.Do(func() {
		ordNumber = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_ordNumberImpl(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqNumber()
})})
	})
	return ordNumber
}

var ordInt gopurs_runtime.Value
var once_ordInt sync.Once
func Get_ordInt() gopurs_runtime.Value {
	once_ordInt.Do(func() {
		ordInt = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_ordIntImpl(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqInt()
})})
	})
	return ordInt
}

var ordChar gopurs_runtime.Value
var once_ordChar sync.Once
func Get_ordChar() gopurs_runtime.Value {
	once_ordChar.Do(func() {
		ordChar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_ordCharImpl(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqChar()
})})
	})
	return ordChar
}

var ordBoolean gopurs_runtime.Value
var once_ordBoolean sync.Once
func Get_ordBoolean() gopurs_runtime.Value {
	once_ordBoolean.Do(func() {
		ordBoolean = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_ordBooleanImpl(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eqBoolean()
})})
	})
	return ordBoolean
}

var compareRecord gopurs_runtime.Value
var once_compareRecord sync.Once
func Get_compareRecord() gopurs_runtime.Value {
	once_compareRecord.Do(func() {
		compareRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["compareRecord"]
})
	})
	return compareRecord
}

var ordRecord gopurs_runtime.Value
var once_ordRecord sync.Once
func Get_ordRecord() gopurs_runtime.Value {
	once_ordRecord.Do(func() {
		ordRecord = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictOrdRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
eqRec1_2_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrdRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["EqRecord0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["eqRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Apply(dictOrdRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["compareRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), "Eq0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRec1_2_0
})})
})
})
	})
	return ordRecord
}

var compare1 gopurs_runtime.Value
var once_compare1 sync.Once
func Get_compare1() gopurs_runtime.Value {
	once_compare1.Do(func() {
		compare1 = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["compare1"]
})
	})
	return compare1
}

var compare gopurs_runtime.Value
var once_compare sync.Once
func Get_compare() gopurs_runtime.Value {
	once_compare.Do(func() {
		compare = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
})
	})
	return compare
}

var comparing gopurs_runtime.Value
var once_comparing sync.Once
func Get_comparing() gopurs_runtime.Value {
	once_comparing.Do(func() {
		comparing = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(f_1, x_2)), gopurs_runtime.Apply(f_1, y_3))
})
})
})
})
	})
	return comparing
}

var greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		greaterThan = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], a1_1), a2_2).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")
})
})
})
	})
	return greaterThan
}

var greaterThanOrEq gopurs_runtime.Value
var once_greaterThanOrEq sync.Once
func Get_greaterThanOrEq() gopurs_runtime.Value {
	once_greaterThanOrEq.Do(func() {
		greaterThanOrEq = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], a1_1), a2_2).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)
})
})
})
	})
	return greaterThanOrEq
}

var lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		lessThan = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], a1_1), a2_2).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")
})
})
})
	})
	return lessThan
}

var signum gopurs_runtime.Value
var once_signum sync.Once
func Get_signum() gopurs_runtime.Value {
	once_signum.Do(func() {
		signum = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictRing_1 gopurs_runtime.Value) gopurs_runtime.Value {
Semiring0_2_0 := gopurs_runtime.Apply(dictRing_1.PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{})
zero_3_1 := Semiring0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["zero"]
zero_4_2 := gopurs_runtime.Apply(dictRing_1.PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["zero"]
one_5_3 := Semiring0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["one"]
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_6), zero_3_1).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictRing_1.PtrVal.(map[string]gopurs_runtime.Value)["sub"], zero_4_2), one_5_3)
} else {
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_6), zero_3_1).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t4 = one_5_3
} else {
__t4 = x_6
}
}
return __t4
})
})
})
	})
	return signum
}

var lessThanOrEq gopurs_runtime.Value
var once_lessThanOrEq sync.Once
func Get_lessThanOrEq() gopurs_runtime.Value {
	once_lessThanOrEq.Do(func() {
		lessThanOrEq = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], a1_1), a2_2).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)
})
})
})
	})
	return lessThanOrEq
}

var max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		max = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_1), y_2)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = y_2
} else {
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t1 = x_1
} else {
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t1 = x_1
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
})
})
})
	})
	return max
}

var min gopurs_runtime.Value
var once_min sync.Once
func Get_min() gopurs_runtime.Value {
	once_min.Do(func() {
		min = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_1), y_2)
var __t1 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = x_1
} else {
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t1 = x_1
} else {
if (gopurs_runtime.Bool(v_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t1 = y_2
} else {
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t1
})
})
})
	})
	return min
}

var ordArray gopurs_runtime.Value
var once_ordArray sync.Once
func Get_ordArray() gopurs_runtime.Value {
	once_ordArray.Do(func() {
		ordArray = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
eqArray_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["eq"])})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Int(0)), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_ordArrayImpl(), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4), y_5)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t2 = gopurs_runtime.Int(0)
} else {
if (gopurs_runtime.Bool(v_6_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t2 = gopurs_runtime.Int(1)
} else {
if (gopurs_runtime.Bool(v_6_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), gopurs_runtime.Int(0)), gopurs_runtime.Int(1))
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t2
})
})), xs_2), ys_3))
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqArray_1_0
})})
})
	})
	return ordArray
}

var ord1Array gopurs_runtime.Value
var once_ord1Array sync.Once
func Get_ord1Array() gopurs_runtime.Value {
	once_ord1Array.Do(func() {
		ord1Array = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_ordArray(), dictOrd_0).PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eq1Array()
})})
	})
	return ord1Array
}

var ordRecordCons gopurs_runtime.Value
var once_ordRecordCons sync.Once
func Get_ordRecordCons() gopurs_runtime.Value {
	once_ordRecordCons.Do(func() {
		ordRecordCons = gopurs_runtime.Func(func(dictOrdRecord_0 gopurs_runtime.Value) gopurs_runtime.Value {
eqRowCons_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply(dictOrdRecord_0.PtrVal.(map[string]gopurs_runtime.Value)["EqRecord0"], gopurs_runtime.Value{})), gopurs_runtime.Value{})
return gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictIsSymbol_3 gopurs_runtime.Value) gopurs_runtime.Value {
eqRowCons1_4_1 := gopurs_runtime.Apply(eqRowCons_1_0, dictIsSymbol_3)
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
eqRowCons2_6_2 := gopurs_runtime.Apply(eqRowCons1_4_1, gopurs_runtime.Apply(dictOrd_5.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compareRecord": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
key_10_3 := gopurs_runtime.Apply(dictIsSymbol_3.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
left_11_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_5.PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_10_3), ra_8)), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_10_3), rb_9))
var __t5 gopurs_runtime.Value
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqBooleanImpl(), gopurs_runtime.Bool(left_11_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")), gopurs_runtime.Bool(false))).IntVal != 0 {
__t5 = left_11_4
} else {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrdRecord_0.PtrVal.(map[string]gopurs_runtime.Value)["compareRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), ra_8), rb_9)
}
return __t5
})
})
}), "EqRecord0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRowCons2_6_2
})})
})
})
})
})
	})
	return ordRecordCons
}

var clamp gopurs_runtime.Value
var once_clamp sync.Once
func Get_clamp() gopurs_runtime.Value {
	once_clamp.Do(func() {
		clamp = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(low_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(hi_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], low_1), x_3)
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_4_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t2 = x_3
} else {
if (gopurs_runtime.Bool(v_4_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t2 = low_1
} else {
if (gopurs_runtime.Bool(v_4_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t2 = low_1
} else {
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
__local_var_5_1 := __t2
v_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], hi_2), __local_var_5_1)
var __t4 gopurs_runtime.Value
if (gopurs_runtime.Bool(v_6_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t4 = hi_2
} else {
if (gopurs_runtime.Bool(v_6_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
__t4 = hi_2
} else {
if (gopurs_runtime.Bool(v_6_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t4 = __local_var_5_1
} else {
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
}
}
return __t4
})
})
})
})
	})
	return clamp
}

var between gopurs_runtime.Value
var once_between sync.Once
func Get_between() gopurs_runtime.Value {
	once_between.Do(func() {
		between = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(low_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(hi_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_3), low_1).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(false)
} else {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_3), hi_2).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT").IntVal == 0)
}
return __t0
})
})
})
})
	})
	return between
}

var abs gopurs_runtime.Value
var once_abs sync.Once
func Get_abs() gopurs_runtime.Value {
	once_abs.Do(func() {
		abs = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictRing_1 gopurs_runtime.Value) gopurs_runtime.Value {
zero_2_0 := gopurs_runtime.Apply(dictRing_1.PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["zero"]
zero_3_1 := gopurs_runtime.Apply(dictRing_1.PtrVal.(map[string]gopurs_runtime.Value)["Semiring0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["zero"]
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4), zero_2_0).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT").IntVal == 0)).IntVal != 0 {
__t2 = x_4
} else {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictRing_1.PtrVal.(map[string]gopurs_runtime.Value)["sub"], zero_3_1), x_4)
}
return __t2
})
})
})
	})
	return abs
}

func Get_ordArrayImpl() gopurs_runtime.Value {
	return OrdArrayImpl
}

func Get_ordBooleanImpl() gopurs_runtime.Value {
	return OrdBooleanImpl
}

func Get_ordCharImpl() gopurs_runtime.Value {
	return OrdCharImpl
}

func Get_ordIntImpl() gopurs_runtime.Value {
	return OrdIntImpl
}

func Get_ordNumberImpl() gopurs_runtime.Value {
	return OrdNumberImpl
}

func Get_ordStringImpl() gopurs_runtime.Value {
	return OrdStringImpl
}
