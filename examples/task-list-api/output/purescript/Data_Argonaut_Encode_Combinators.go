package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Argonaut_Encode_Combinators_extendOptional gopurs_runtime.Value
var once_Data_Argonaut_Encode_Combinators_extendOptional sync.Once

func Get_Data_Argonaut_Encode_Combinators_extendOptional() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_Combinators_extendOptional.Do(func() {
		cache_Data_Argonaut_Encode_Combinators_extendOptional = gopurs_runtime.Func(func(dictEncodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Encode_Combinators_extendOptional(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](dictEncodeJson_0_box))
		})
	})
	return cache_Data_Argonaut_Encode_Combinators_extendOptional
}

var cache_Data_Argonaut_Encode_Combinators_extend gopurs_runtime.Value
var once_Data_Argonaut_Encode_Combinators_extend sync.Once

func Get_Data_Argonaut_Encode_Combinators_extend() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_Combinators_extend.Do(func() {
		cache_Data_Argonaut_Encode_Combinators_extend = gopurs_runtime.Func(func(dictEncodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Encode_Combinators_extend(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](dictEncodeJson_0_box))
		})
	})
	return cache_Data_Argonaut_Encode_Combinators_extend
}

var cache_Data_Argonaut_Encode_Combinators_assocOptional gopurs_runtime.Value
var once_Data_Argonaut_Encode_Combinators_assocOptional sync.Once

func Get_Data_Argonaut_Encode_Combinators_assocOptional() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_Combinators_assocOptional.Do(func() {
		cache_Data_Argonaut_Encode_Combinators_assocOptional = gopurs_runtime.Func(func(dictEncodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Encode_Combinators_assocOptional(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](dictEncodeJson_0_box))
		})
	})
	return cache_Data_Argonaut_Encode_Combinators_assocOptional
}

var cache_Data_Argonaut_Encode_Combinators_assoc gopurs_runtime.Value
var once_Data_Argonaut_Encode_Combinators_assoc sync.Once

func Get_Data_Argonaut_Encode_Combinators_assoc() gopurs_runtime.Value {
	once_Data_Argonaut_Encode_Combinators_assoc.Do(func() {
		cache_Data_Argonaut_Encode_Combinators_assoc = gopurs_runtime.Func(func(dictEncodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Encode_Combinators_assoc(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]](dictEncodeJson_0_box))
		})
	})
	return cache_Data_Argonaut_Encode_Combinators_assoc
}

func Call_Data_Argonaut_Encode_Combinators_extendOptional(dictEncodeJson_0_loop *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEncodeJson_0 *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] = dictEncodeJson_0_loop
	_ = dictEncodeJson_0
	return gopurs_runtime.Apply(Get_Data_Argonaut_Encode_Encoders_extendOptional(), Call_Data_Argonaut_Encode_Class_encodeJson(dictEncodeJson_0))
}

func Call_Data_Argonaut_Encode_Combinators_extend(dictEncodeJson_0_loop *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEncodeJson_0 *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] = dictEncodeJson_0_loop
	_ = dictEncodeJson_0
	return gopurs_runtime.Apply(Get_Data_Argonaut_Encode_Encoders_extend(), Call_Data_Argonaut_Encode_Class_encodeJson(dictEncodeJson_0))
}

func Call_Data_Argonaut_Encode_Combinators_assocOptional(dictEncodeJson_0_loop *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEncodeJson_0 *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] = dictEncodeJson_0_loop
	_ = dictEncodeJson_0
	return gopurs_runtime.Apply(Get_Data_Argonaut_Encode_Encoders_assocOptional(), Call_Data_Argonaut_Encode_Class_encodeJson(dictEncodeJson_0))
}

func Call_Data_Argonaut_Encode_Combinators_assoc(dictEncodeJson_0_loop *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictEncodeJson_0 *Constructor_Data_Argonaut_Encode_Class_EncodeJson[gopurs_runtime.Value] = dictEncodeJson_0_loop
	_ = dictEncodeJson_0
	return gopurs_runtime.Apply(Get_Data_Argonaut_Encode_Encoders_assoc(), Call_Data_Argonaut_Encode_Class_encodeJson(dictEncodeJson_0))
}
