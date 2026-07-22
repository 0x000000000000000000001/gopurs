package Data_Show

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Void "gopurs/output/Data.Void"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var showVoid gopurs_runtime.Value
var once_showVoid sync.Once
func Get_showVoid() gopurs_runtime.Value {
	once_showVoid.Do(func() {
		showVoid = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": pkg_Data_Void.Get_absurd()})
	})
	return showVoid
}

var showUnit gopurs_runtime.Value
var once_showUnit sync.Once
func Get_showUnit() gopurs_runtime.Value {
	once_showUnit.Do(func() {
		showUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("unit")
})})
	})
	return showUnit
}

var showString gopurs_runtime.Value
var once_showString sync.Once
func Get_showString() gopurs_runtime.Value {
	once_showString.Do(func() {
		showString = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": Get_showStringImpl()})
	})
	return showString
}

var showRecordFieldsNil gopurs_runtime.Value
var once_showRecordFieldsNil sync.Once
func Get_showRecordFieldsNil() gopurs_runtime.Value {
	once_showRecordFieldsNil.Do(func() {
		showRecordFieldsNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"showRecordFields": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
})
})})
	})
	return showRecordFieldsNil
}

var showRecordFields gopurs_runtime.Value
var once_showRecordFields sync.Once
func Get_showRecordFields() gopurs_runtime.Value {
	once_showRecordFields.Do(func() {
		showRecordFields = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["showRecordFields"]
})
	})
	return showRecordFields
}

var showRecord gopurs_runtime.Value
var once_showRecord sync.Once
func Get_showRecord() gopurs_runtime.Value {
	once_showRecord.Do(func() {
		showRecord = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShowRecordFields_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(record_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("{")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictShowRecordFields_2.PtrVal.(map[string]gopurs_runtime.Value)["showRecordFields"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), record_3)), gopurs_runtime.Str("}")))
})})
})
})
})
	})
	return showRecord
}

var showProxy gopurs_runtime.Value
var once_showProxy sync.Once
func Get_showProxy() gopurs_runtime.Value {
	once_showProxy.Do(func() {
		showProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("Proxy")
})})
	})
	return showProxy
}

var showNumber gopurs_runtime.Value
var once_showNumber sync.Once
func Get_showNumber() gopurs_runtime.Value {
	once_showNumber.Do(func() {
		showNumber = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": Get_showNumberImpl()})
	})
	return showNumber
}

var showInt gopurs_runtime.Value
var once_showInt sync.Once
func Get_showInt() gopurs_runtime.Value {
	once_showInt.Do(func() {
		showInt = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": Get_showIntImpl()})
	})
	return showInt
}

var showChar gopurs_runtime.Value
var once_showChar sync.Once
func Get_showChar() gopurs_runtime.Value {
	once_showChar.Do(func() {
		showChar = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": Get_showCharImpl()})
	})
	return showChar
}

var showBoolean gopurs_runtime.Value
var once_showBoolean sync.Once
func Get_showBoolean() gopurs_runtime.Value {
	once_showBoolean.Do(func() {
		showBoolean = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
if (v_0).IntVal != 0 {
__t0 = gopurs_runtime.Str("true")
} else {
__t0 = gopurs_runtime.Str("false")
}
return __t0
})})
	})
	return showBoolean
}

var show gopurs_runtime.Value
var once_show sync.Once
func Get_show() gopurs_runtime.Value {
	once_show.Do(func() {
		show = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["show"]
})
	})
	return show
}

var showArray gopurs_runtime.Value
var once_showArray sync.Once
func Get_showArray() gopurs_runtime.Value {
	once_showArray.Do(func() {
		showArray = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Apply(Get_showArrayImpl(), dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"])})
})
	})
	return showArray
}

var showRecordFieldsCons gopurs_runtime.Value
var once_showRecordFieldsCons sync.Once
func Get_showRecordFieldsCons() gopurs_runtime.Value {
	once_showRecordFieldsCons.Do(func() {
		showRecordFieldsCons = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShowRecordFields_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"showRecordFields": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(record_4 gopurs_runtime.Value) gopurs_runtime.Value {
key_5_0 := gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), key_5_0), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(": ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_2.PtrVal.(map[string]gopurs_runtime.Value)["show"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_5_0), record_4))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(",")), gopurs_runtime.Apply(gopurs_runtime.Apply(dictShowRecordFields_1.PtrVal.(map[string]gopurs_runtime.Value)["showRecordFields"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), record_4))))))
})
})})
})
})
})
	})
	return showRecordFieldsCons
}

var showRecordFieldsConsNil gopurs_runtime.Value
var once_showRecordFieldsConsNil sync.Once
func Get_showRecordFieldsConsNil() gopurs_runtime.Value {
	once_showRecordFieldsConsNil.Do(func() {
		showRecordFieldsConsNil = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"showRecordFields": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(record_3 gopurs_runtime.Value) gopurs_runtime.Value {
key_4_0 := gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), key_4_0), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(": ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_4_0), record_3))), gopurs_runtime.Str(" ")))))
})
})})
})
})
	})
	return showRecordFieldsConsNil
}

func Get_showArrayImpl() gopurs_runtime.Value {
	return ShowArrayImpl
}

func Get_showCharImpl() gopurs_runtime.Value {
	return ShowCharImpl
}

func Get_showIntImpl() gopurs_runtime.Value {
	return ShowIntImpl
}

func Get_showNumberImpl() gopurs_runtime.Value {
	return ShowNumberImpl
}

func Get_showStringImpl() gopurs_runtime.Value {
	return ShowStringImpl
}
