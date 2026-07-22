package Data_Reflectable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var reifiableString gopurs_runtime.Value
var once_reifiableString sync.Once
func Get_reifiableString() gopurs_runtime.Value {
	once_reifiableString.Do(func() {
		reifiableString = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return reifiableString
}

var reifiableOrdering gopurs_runtime.Value
var once_reifiableOrdering sync.Once
func Get_reifiableOrdering() gopurs_runtime.Value {
	once_reifiableOrdering.Do(func() {
		reifiableOrdering = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return reifiableOrdering
}

var reifiableInt gopurs_runtime.Value
var once_reifiableInt sync.Once
func Get_reifiableInt() gopurs_runtime.Value {
	once_reifiableInt.Do(func() {
		reifiableInt = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return reifiableInt
}

var reifiableBoolean gopurs_runtime.Value
var once_reifiableBoolean sync.Once
func Get_reifiableBoolean() gopurs_runtime.Value {
	once_reifiableBoolean.Do(func() {
		reifiableBoolean = gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
	})
	return reifiableBoolean
}

var reifyType gopurs_runtime.Value
var once_reifyType sync.Once
func Get_reifyType() gopurs_runtime.Value {
	once_reifyType.Do(func() {
		reifyType = gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_unsafeCoerce(), gopurs_runtime.Func(func(dictReflectable_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, dictReflectable_3)
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"reflectType": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return s_1
})})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Proxy")}))
})
})
})
	})
	return reifyType
}

var reflectType gopurs_runtime.Value
var once_reflectType sync.Once
func Get_reflectType() gopurs_runtime.Value {
	once_reflectType.Do(func() {
		reflectType = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["reflectType"]
})
	})
	return reflectType
}

func Get_unsafeCoerce() gopurs_runtime.Value {
	return UnsafeCoerce
}
