package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Argonaut_Decode_Combinators_go__map gopurs_runtime.Value
var once_Data_Argonaut_Decode_Combinators_go__map sync.Once

func Get_Data_Argonaut_Decode_Combinators_go__map() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Combinators_go__map.Do(func() {
		cache_Data_Argonaut_Decode_Combinators_go__map = gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map")
	})
	return cache_Data_Argonaut_Decode_Combinators_go__map
}

var cache_Data_Argonaut_Decode_Combinators_getFieldOptional_prime_ gopurs_runtime.Value
var once_Data_Argonaut_Decode_Combinators_getFieldOptional_prime_ sync.Once

func Get_Data_Argonaut_Decode_Combinators_getFieldOptional_prime_() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Combinators_getFieldOptional_prime_.Do(func() {
		cache_Data_Argonaut_Decode_Combinators_getFieldOptional_prime_ = gopurs_runtime.Func(func(dictDecodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Combinators_getFieldOptional_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](dictDecodeJson_0_box))
		})
	})
	return cache_Data_Argonaut_Decode_Combinators_getFieldOptional_prime_
}

var cache_Data_Argonaut_Decode_Combinators_getFieldOptional gopurs_runtime.Value
var once_Data_Argonaut_Decode_Combinators_getFieldOptional sync.Once

func Get_Data_Argonaut_Decode_Combinators_getFieldOptional() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Combinators_getFieldOptional.Do(func() {
		cache_Data_Argonaut_Decode_Combinators_getFieldOptional = gopurs_runtime.Func(func(dictDecodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Combinators_getFieldOptional(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](dictDecodeJson_0_box))
		})
	})
	return cache_Data_Argonaut_Decode_Combinators_getFieldOptional
}

var cache_Data_Argonaut_Decode_Combinators_getField gopurs_runtime.Value
var once_Data_Argonaut_Decode_Combinators_getField sync.Once

func Get_Data_Argonaut_Decode_Combinators_getField() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Combinators_getField.Do(func() {
		cache_Data_Argonaut_Decode_Combinators_getField = gopurs_runtime.Func(func(dictDecodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Combinators_getField(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](dictDecodeJson_0_box))
		})
	})
	return cache_Data_Argonaut_Decode_Combinators_getField
}

var cache_Data_Argonaut_Decode_Combinators_defaultField gopurs_runtime.Value
var once_Data_Argonaut_Decode_Combinators_defaultField sync.Once

func Get_Data_Argonaut_Decode_Combinators_defaultField() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Combinators_defaultField.Do(func() {
		cache_Data_Argonaut_Decode_Combinators_defaultField = gopurs_runtime.Func2(func(parser_0_box gopurs_runtime.Value, go__default_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Argonaut_Decode_Combinators_defaultField(parser_0_box, go__default_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Argonaut_Decode_Combinators_defaultField
}

func Call_Data_Argonaut_Decode_Combinators_getFieldOptional_prime_(dictDecodeJson_0_loop *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictDecodeJson_0 *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] = dictDecodeJson_0_loop
	_ = dictDecodeJson_0
	return gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Decoders_getFieldOptional_prime_(), Call_Data_Argonaut_Decode_Class_decodeJson(dictDecodeJson_0))
}

func Call_Data_Argonaut_Decode_Combinators_getFieldOptional(dictDecodeJson_0_loop *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictDecodeJson_0 *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] = dictDecodeJson_0_loop
	_ = dictDecodeJson_0
	return gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Decoders_getFieldOptional(), Call_Data_Argonaut_Decode_Class_decodeJson(dictDecodeJson_0))
}

func Call_Data_Argonaut_Decode_Combinators_getField(dictDecodeJson_0_loop *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictDecodeJson_0 *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] = dictDecodeJson_0_loop
	_ = dictDecodeJson_0
	return gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Decoders_getField(), Call_Data_Argonaut_Decode_Class_decodeJson(dictDecodeJson_0))
}

func Call_Data_Argonaut_Decode_Combinators_defaultField(parser_0_loop gopurs_runtime.Value, go__default_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
	var parser_0 gopurs_runtime.Value = parser_0_loop
	_ = parser_0
	var go__default_1 gopurs_runtime.Value = go__default_1_loop
	_ = go__default_1
	var __t3 struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
		V2 bool
	}
	{
		if parser_0.Type == 9 && parser_0.IntVal == 3711209382 {
			__t3 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(parser_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
			goto end_branch_3
		} else {

		}
	}
	{
		if parser_0.Type == 9 && parser_0.IntVal == 2465973597 {
			var __t2 gopurs_runtime.Value
			{
				var __t_tag_0 gopurs_runtime.Value = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(parser_0.UnsafePtr).V0
				_ = __t_tag_0
				if __t_tag_0.Type == 9 && __t_tag_0.IntVal == 930809136 && __t_tag_0.UnsafePtr == nil {
					__t2 = go__default_1
					goto end_branch_2
				} else {

				}
			}
			{
				var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(parser_0.UnsafePtr).V0
				_ = __t_tag_1
				if __t_tag_1.Type == 9 && __t_tag_1.IntVal == 930809136 && __t_tag_1.UnsafePtr != nil {
					__t2 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(parser_0.UnsafePtr).V0.UnsafePtr).V0
					goto end_branch_2
				} else {

				}
			}
			{
				__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_2:
			__t3 = struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{gopurs_runtime.Value{}, __t2, true}
			goto end_branch_3
		} else {

		}
	}
	{
		__t3 = func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := func() gopurs_runtime.Value { panic("Failed pattern match") }(); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
end_branch_3:
	return __t3
}
