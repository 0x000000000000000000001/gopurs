package Data_Semigroup

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Void "gopurs/output/Data.Void"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var semigroupVoid gopurs_runtime.Value
var once_semigroupVoid sync.Once
func Get_semigroupVoid() gopurs_runtime.Value {
	once_semigroupVoid.Do(func() {
		semigroupVoid = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Void.Get_absurd()
})})
	})
	return semigroupVoid
}

var semigroupUnit gopurs_runtime.Value
var once_semigroupUnit sync.Once
func Get_semigroupUnit() gopurs_runtime.Value {
	once_semigroupUnit.Do(func() {
		semigroupUnit = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
})})
	})
	return semigroupUnit
}

var semigroupString gopurs_runtime.Value
var once_semigroupString sync.Once
func Get_semigroupString() gopurs_runtime.Value {
	once_semigroupString.Do(func() {
		semigroupString = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": Get_concatString()})
	})
	return semigroupString
}

var semigroupRecordNil gopurs_runtime.Value
var once_semigroupRecordNil sync.Once
func Get_semigroupRecordNil() gopurs_runtime.Value {
	once_semigroupRecordNil.Do(func() {
		semigroupRecordNil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"appendRecord": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
})
})
})})
	})
	return semigroupRecordNil
}

var semigroupProxy gopurs_runtime.Value
var once_semigroupProxy sync.Once
func Get_semigroupProxy() gopurs_runtime.Value {
	once_semigroupProxy.Do(func() {
		semigroupProxy = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})
})
})})
	})
	return semigroupProxy
}

var semigroupArray gopurs_runtime.Value
var once_semigroupArray sync.Once
func Get_semigroupArray() gopurs_runtime.Value {
	once_semigroupArray.Do(func() {
		semigroupArray = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": Get_concatArray()})
	})
	return semigroupArray
}

var appendRecord gopurs_runtime.Value
var once_appendRecord sync.Once
func Get_appendRecord() gopurs_runtime.Value {
	once_appendRecord.Do(func() {
		appendRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["appendRecord"]
})
	})
	return appendRecord
}

var semigroupRecord gopurs_runtime.Value
var once_semigroupRecord sync.Once
func Get_semigroupRecord() gopurs_runtime.Value {
	once_semigroupRecord.Do(func() {
		semigroupRecord = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemigroupRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Apply(dictSemigroupRecord_1.PtrVal.(map[string]gopurs_runtime.Value)["appendRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))})
})
})
	})
	return semigroupRecord
}

var append_ gopurs_runtime.Value
var once_append_ sync.Once
func Get_append_() gopurs_runtime.Value {
	once_append_.Do(func() {
		append_ = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["append"]
})
	})
	return append_
}

var semigroupFn gopurs_runtime.Value
var once_semigroupFn sync.Once
func Get_semigroupFn() gopurs_runtime.Value {
	once_semigroupFn.Do(func() {
		semigroupFn = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(f_1, x_3)), gopurs_runtime.Apply(g_2, x_3))
})
})
})})
})
	})
	return semigroupFn
}

var semigroupRecordCons gopurs_runtime.Value
var once_semigroupRecordCons sync.Once
func Get_semigroupRecordCons() gopurs_runtime.Value {
	once_semigroupRecordCons.Do(func() {
		semigroupRecordCons = gopurs_runtime.Func(func(dictIsSymbol_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemigroupRecord_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"appendRecord": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
key_7_0 := gopurs_runtime.Apply(dictIsSymbol_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectSymbol"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
get_8_1 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_7_0)
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeSet(), key_7_0), gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_3.PtrVal.(map[string]gopurs_runtime.Value)["append"], gopurs_runtime.Apply(get_8_1, ra_5)), gopurs_runtime.Apply(get_8_1, rb_6))), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroupRecord_2.PtrVal.(map[string]gopurs_runtime.Value)["appendRecord"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")})), ra_5), rb_6))
})
})
})})
})
})
})
})
	})
	return semigroupRecordCons
}

func Get_concatArray() gopurs_runtime.Value {
	return ConcatArray
}

func Get_concatString() gopurs_runtime.Value {
	return ConcatString
}
