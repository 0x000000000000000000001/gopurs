package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Argonaut_Decode_Parser_bind gopurs_runtime.Value
var once_Data_Argonaut_Decode_Parser_bind sync.Once

func Get_Data_Argonaut_Decode_Parser_bind() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Parser_bind.Do(func() {
		cache_Data_Argonaut_Decode_Parser_bind = gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind")
	})
	return cache_Data_Argonaut_Decode_Parser_bind
}

var cache_Data_Argonaut_Decode_Parser_parseJson gopurs_runtime.Value
var once_Data_Argonaut_Decode_Parser_parseJson sync.Once

func Get_Data_Argonaut_Decode_Parser_parseJson() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Parser_parseJson.Do(func() {
		cache_Data_Argonaut_Decode_Parser_parseJson = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(__eta2_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}
			{
				if __eta2_0.Type == 9 && __eta2_0.IntVal == 3711209382 {
					__t0 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{Type: 9, IntVal: 2887704423, UnsafePtr: unsafe.Pointer((&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{1, "JSON"}))}, gopurs_runtime.Value{}, false}
					goto end_branch_0
				} else {

				}
			}
			{
				if __eta2_0.Type == 9 && __eta2_0.IntVal == 2465973597 {
					__t0 = struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{gopurs_runtime.Value{}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__eta2_0.UnsafePtr).V0, true}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() struct {
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
		end_branch_0:
			return func() gopurs_runtime.Value {
				_v := __t0
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		}), Get_Data_Argonaut_Parser_jsonParser())
	})
	return cache_Data_Argonaut_Decode_Parser_parseJson
}

var cache_Data_Argonaut_Decode_Parser_decodeJsonStringWith gopurs_runtime.Value
var once_Data_Argonaut_Decode_Parser_decodeJsonStringWith sync.Once

func Get_Data_Argonaut_Decode_Parser_decodeJsonStringWith() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Parser_decodeJsonStringWith.Do(func() {
		cache_Data_Argonaut_Decode_Parser_decodeJsonStringWith = gopurs_runtime.Func(func(decoder_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Parser_decodeJsonStringWith(decoder_0_box)
		})
	})
	return cache_Data_Argonaut_Decode_Parser_decodeJsonStringWith
}

var cache_Data_Argonaut_Decode_Parser_decodeJsonString gopurs_runtime.Value
var once_Data_Argonaut_Decode_Parser_decodeJsonString sync.Once

func Get_Data_Argonaut_Decode_Parser_decodeJsonString() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Parser_decodeJsonString.Do(func() {
		cache_Data_Argonaut_Decode_Parser_decodeJsonString = gopurs_runtime.Func(func(dictDecodeJson_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Parser_decodeJsonString(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](dictDecodeJson_0_box))
		})
	})
	return cache_Data_Argonaut_Decode_Parser_decodeJsonString
}

func Call_Data_Argonaut_Decode_Parser_decodeJsonStringWith(decoder_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var decoder_0 gopurs_runtime.Value = decoder_0_loop
	_ = decoder_0
	return gopurs_runtime.Apply3(Get_Data_Argonaut_Decode_Parser_decodeJsonStringImpl(), gopurs_runtime.Int(Get_Data_Argonaut_Decode_Internal_Record_schemaDecoderABI1().IntVal), gopurs_runtime.Func(func(text_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Parser_parseJson(), gopurs_runtime.NormalizeStringValue(text_1))
		_ = __local_var_2_0
		var __t1 gopurs_runtime.Value
		{
			if __local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3711209382 {
				__t1 = func() gopurs_runtime.Value {
					_v := struct {
						V0 gopurs_runtime.Value
						V1 gopurs_runtime.Value
						V2 bool
					}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
					if _v.V2 {
						return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
					}
					return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
				}()
				goto end_branch_1
			} else {

			}
		}
		{
			if __local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 2465973597 {
				__t1 = gopurs_runtime.Apply(decoder_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0)
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_1:
		return __t1
	}), decoder_0)
}

func Call_Data_Argonaut_Decode_Parser_decodeJsonString(dictDecodeJson_0_loop *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictDecodeJson_0 *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] = dictDecodeJson_0_loop
	_ = dictDecodeJson_0
	return Call_Data_Argonaut_Decode_Parser_decodeJsonStringWith(Call_Data_Argonaut_Decode_Class_decodeJson(dictDecodeJson_0))
}

func Get_Data_Argonaut_Decode_Parser_decodeJsonStringImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Parser_DecodeJsonStringImpl
}

func Get_Data_Argonaut_Decode_Parser_textDecoderABI1() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Parser_TextDecoderABI1
}
