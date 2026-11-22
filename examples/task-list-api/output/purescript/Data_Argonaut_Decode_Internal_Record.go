package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Argonaut_Decode_Internal_Record_rightValue gopurs_runtime.Value
var once_Data_Argonaut_Decode_Internal_Record_rightValue sync.Once

func Get_Data_Argonaut_Decode_Internal_Record_rightValue() gopurs_runtime.Value {
	once_Data_Argonaut_Decode_Internal_Record_rightValue.Do(func() {
		cache_Data_Argonaut_Decode_Internal_Record_rightValue = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Decode_Internal_Record_rightValue(v_0_box)
		})
	})
	return cache_Data_Argonaut_Decode_Internal_Record_rightValue
}

func Call_Data_Argonaut_Decode_Internal_Record_rightValue(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var __t0 gopurs_runtime.Value
	{
		if v_0.Type == 9 && v_0.IntVal == 2465973597 {
			__t0 = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0.Type == 9 && v_0.IntVal == 3711209382 {
			__t0 = gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.StrWithCodeUnits("Record decoder expected a successful field", 42))
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Get_Data_Argonaut_Decode_Internal_Record_borrowObject() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_BorrowObject
}

func Get_Data_Argonaut_Decode_Internal_Record_decodeFieldFast() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_DecodeFieldFast
}

func Get_Data_Argonaut_Decode_Internal_Record_fieldStep() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_FieldStep
}

func Get_Data_Argonaut_Decode_Internal_Record_getFieldImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_GetFieldImpl
}

func Get_Data_Argonaut_Decode_Internal_Record_getFieldOptionalImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_GetFieldOptionalImpl
}

func Get_Data_Argonaut_Decode_Internal_Record_getFieldOptionalNullableImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_GetFieldOptionalNullableImpl
}

func Get_Data_Argonaut_Decode_Internal_Record_recordConsImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_RecordConsImpl
}

func Get_Data_Argonaut_Decode_Internal_Record_recordNilImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_RecordNilImpl
}

func Get_Data_Argonaut_Decode_Internal_Record_schemaDecoderABI1() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_SchemaDecoderABI1
}

func Get_Data_Argonaut_Decode_Internal_Record_typedArray() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_TypedArray
}

func Get_Data_Argonaut_Decode_Internal_Record_typedBoolean() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_TypedBoolean
}

func Get_Data_Argonaut_Decode_Internal_Record_typedFieldId() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_TypedFieldId
}

func Get_Data_Argonaut_Decode_Internal_Record_typedFieldMaybe() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_TypedFieldMaybe
}

func Get_Data_Argonaut_Decode_Internal_Record_typedInt() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_TypedInt
}

func Get_Data_Argonaut_Decode_Internal_Record_typedJson() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_TypedJson
}

func Get_Data_Argonaut_Decode_Internal_Record_typedMaybe() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_TypedMaybe
}

func Get_Data_Argonaut_Decode_Internal_Record_typedNumber() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_TypedNumber
}

func Get_Data_Argonaut_Decode_Internal_Record_typedObject() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_TypedObject
}

func Get_Data_Argonaut_Decode_Internal_Record_typedRecord() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_TypedRecord
}

func Get_Data_Argonaut_Decode_Internal_Record_typedString() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Decode_Internal_Record_TypedString
}
