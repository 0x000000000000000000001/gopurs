package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_String_Common_eq gopurs_runtime.Value
var once_Data_String_Common_eq sync.Once

func Get_Data_String_Common_eq() gopurs_runtime.Value {
	once_Data_String_Common_eq.Do(func() {
		cache_Data_String_Common_eq = Get_Data_Eq_eq__3478410105()
	})
	return cache_Data_String_Common_eq
}

var cache_Data_String_Common_null gopurs_runtime.Value
var once_Data_String_Common_null sync.Once

func Get_Data_String_Common_null() gopurs_runtime.Value {
	once_Data_String_Common_null.Do(func() {
		cache_Data_String_Common_null = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_String_Common_null(s_0_box.StrVal()))
		})
	})
	return cache_Data_String_Common_null
}

var cache_Data_String_Common_localeCompare gopurs_runtime.Value
var once_Data_String_Common_localeCompare sync.Once

func Get_Data_String_Common_localeCompare() gopurs_runtime.Value {
	once_Data_String_Common_localeCompare.Do(func() {
		cache_Data_String_Common_localeCompare = gopurs_runtime.Apply3(Get_Data_String_Common__localeCompare(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})
	})
	return cache_Data_String_Common_localeCompare
}

func Call_Data_String_Common_null(s_0_loop string) bool {
	var s_0 string = s_0_loop
	_ = s_0
	return (s_0) == ("")
}

func Get_Data_String_Common__localeCompare() gopurs_runtime.Value {
	return _Gopurs_Data_String_Common__LocaleCompare
}

func Get_Data_String_Common_joinWith() gopurs_runtime.Value {
	return _Gopurs_Data_String_Common_JoinWith
}

func Get_Data_String_Common_replace() gopurs_runtime.Value {
	return _Gopurs_Data_String_Common_Replace
}

func Get_Data_String_Common_replaceAll() gopurs_runtime.Value {
	return _Gopurs_Data_String_Common_ReplaceAll
}

func Get_Data_String_Common_split() gopurs_runtime.Value {
	return _Gopurs_Data_String_Common_Split
}

func Get_Data_String_Common_toLower() gopurs_runtime.Value {
	return _Gopurs_Data_String_Common_ToLower
}

func Get_Data_String_Common_toUpper() gopurs_runtime.Value {
	return _Gopurs_Data_String_Common_ToUpper
}

func Get_Data_String_Common_trim() gopurs_runtime.Value {
	return _Gopurs_Data_String_Common_Trim
}
