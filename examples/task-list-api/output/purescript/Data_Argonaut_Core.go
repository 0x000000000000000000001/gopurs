package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Argonaut_Core_eq gopurs_runtime.Value
var once_Data_Argonaut_Core_eq sync.Once

func Get_Data_Argonaut_Core_eq() gopurs_runtime.Value {
	once_Data_Argonaut_Core_eq.Do(func() {
		cache_Data_Argonaut_Core_eq = Get_Data_Eq_eq__2123013075()
	})
	return cache_Data_Argonaut_Core_eq
}

var cache_Data_Argonaut_Core_verbJsonType gopurs_runtime.Value
var once_Data_Argonaut_Core_verbJsonType sync.Once

func Get_Data_Argonaut_Core_verbJsonType() gopurs_runtime.Value {
	once_Data_Argonaut_Core_verbJsonType.Do(func() {
		cache_Data_Argonaut_Core_verbJsonType = gopurs_runtime.Func3(func(def_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_verbJsonType(def_0_box, f_1_box, g_2_box)
		})
	})
	return cache_Data_Argonaut_Core_verbJsonType
}

var cache_Data_Argonaut_Core_caseJsonObject gopurs_runtime.Value
var once_Data_Argonaut_Core_caseJsonObject sync.Once

func Get_Data_Argonaut_Core_caseJsonObject() gopurs_runtime.Value {
	once_Data_Argonaut_Core_caseJsonObject.Do(func() {
		cache_Data_Argonaut_Core_caseJsonObject = gopurs_runtime.Func3(func(d_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, j_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_caseJsonObject(d_0_box, f_1_box, j_2_box)
		})
	})
	return cache_Data_Argonaut_Core_caseJsonObject
}

var cache_Data_Argonaut_Core_verbJsonType__1588320074 gopurs_runtime.Value
var once_Data_Argonaut_Core_verbJsonType__1588320074 sync.Once

func Get_Data_Argonaut_Core_verbJsonType__1588320074() gopurs_runtime.Value {
	once_Data_Argonaut_Core_verbJsonType__1588320074.Do(func() {
		cache_Data_Argonaut_Core_verbJsonType__1588320074 = gopurs_runtime.Func3(func(def_0_box gopurs_runtime.Value, f_unused_1_box gopurs_runtime.Value, g_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_verbJsonType__1588320074(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](def_0_box), f_unused_1_box, g_unused_2_box)
		})
	})
	return cache_Data_Argonaut_Core_verbJsonType__1588320074
}

var cache_Data_Argonaut_Core_caseJsonArray gopurs_runtime.Value
var once_Data_Argonaut_Core_caseJsonArray sync.Once

func Get_Data_Argonaut_Core_caseJsonArray() gopurs_runtime.Value {
	once_Data_Argonaut_Core_caseJsonArray.Do(func() {
		cache_Data_Argonaut_Core_caseJsonArray = gopurs_runtime.Func3(func(d_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, j_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_caseJsonArray(d_0_box, f_1_box, j_2_box)
		})
	})
	return cache_Data_Argonaut_Core_caseJsonArray
}

var cache_Data_Argonaut_Core_verbJsonType__286058406 gopurs_runtime.Value
var once_Data_Argonaut_Core_verbJsonType__286058406 sync.Once

func Get_Data_Argonaut_Core_verbJsonType__286058406() gopurs_runtime.Value {
	once_Data_Argonaut_Core_verbJsonType__286058406.Do(func() {
		cache_Data_Argonaut_Core_verbJsonType__286058406 = gopurs_runtime.Func3(func(def_0_box gopurs_runtime.Value, f_unused_1_box gopurs_runtime.Value, g_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_verbJsonType__286058406(Rebox_Data_Argonaut_Core_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](def_0_box)), f_unused_1_box, g_unused_2_box)
		})
	})
	return cache_Data_Argonaut_Core_verbJsonType__286058406
}

var cache_Data_Argonaut_Core_caseJsonBoolean gopurs_runtime.Value
var once_Data_Argonaut_Core_caseJsonBoolean sync.Once

func Get_Data_Argonaut_Core_caseJsonBoolean() gopurs_runtime.Value {
	once_Data_Argonaut_Core_caseJsonBoolean.Do(func() {
		cache_Data_Argonaut_Core_caseJsonBoolean = gopurs_runtime.Func3(func(d_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, j_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_caseJsonBoolean(d_0_box, f_1_box, j_2_box)
		})
	})
	return cache_Data_Argonaut_Core_caseJsonBoolean
}

var cache_Data_Argonaut_Core_verbJsonType__156013339 gopurs_runtime.Value
var once_Data_Argonaut_Core_verbJsonType__156013339 sync.Once

func Get_Data_Argonaut_Core_verbJsonType__156013339() gopurs_runtime.Value {
	once_Data_Argonaut_Core_verbJsonType__156013339.Do(func() {
		cache_Data_Argonaut_Core_verbJsonType__156013339 = gopurs_runtime.Func3(func(def_0_box gopurs_runtime.Value, f_unused_1_box gopurs_runtime.Value, g_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_verbJsonType__156013339(Rebox_Data_Argonaut_Core_3094389156_1044698560(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](def_0_box)), f_unused_1_box, g_unused_2_box)
		})
	})
	return cache_Data_Argonaut_Core_verbJsonType__156013339
}

var cache_Data_Argonaut_Core_caseJsonNumber gopurs_runtime.Value
var once_Data_Argonaut_Core_caseJsonNumber sync.Once

func Get_Data_Argonaut_Core_caseJsonNumber() gopurs_runtime.Value {
	once_Data_Argonaut_Core_caseJsonNumber.Do(func() {
		cache_Data_Argonaut_Core_caseJsonNumber = gopurs_runtime.Func3(func(d_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, j_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_caseJsonNumber(d_0_box, f_1_box, j_2_box)
		})
	})
	return cache_Data_Argonaut_Core_caseJsonNumber
}

var cache_Data_Argonaut_Core_verbJsonType__425335836 gopurs_runtime.Value
var once_Data_Argonaut_Core_verbJsonType__425335836 sync.Once

func Get_Data_Argonaut_Core_verbJsonType__425335836() gopurs_runtime.Value {
	once_Data_Argonaut_Core_verbJsonType__425335836.Do(func() {
		cache_Data_Argonaut_Core_verbJsonType__425335836 = gopurs_runtime.Func3(func(def_0_box gopurs_runtime.Value, f_unused_1_box gopurs_runtime.Value, g_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_verbJsonType__425335836(Rebox_Data_Argonaut_Core_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](def_0_box)), f_unused_1_box, g_unused_2_box)
		})
	})
	return cache_Data_Argonaut_Core_verbJsonType__425335836
}

var cache_Data_Argonaut_Core_caseJsonString gopurs_runtime.Value
var once_Data_Argonaut_Core_caseJsonString sync.Once

func Get_Data_Argonaut_Core_caseJsonString() gopurs_runtime.Value {
	once_Data_Argonaut_Core_caseJsonString.Do(func() {
		cache_Data_Argonaut_Core_caseJsonString = gopurs_runtime.Func3(func(d_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, j_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_caseJsonString(d_0_box, f_1_box, j_2_box)
		})
	})
	return cache_Data_Argonaut_Core_caseJsonString
}

var cache_Data_Argonaut_Core_verbJsonType__1594722986 gopurs_runtime.Value
var once_Data_Argonaut_Core_verbJsonType__1594722986 sync.Once

func Get_Data_Argonaut_Core_verbJsonType__1594722986() gopurs_runtime.Value {
	once_Data_Argonaut_Core_verbJsonType__1594722986.Do(func() {
		cache_Data_Argonaut_Core_verbJsonType__1594722986 = gopurs_runtime.Func3(func(def_0_box gopurs_runtime.Value, f_unused_1_box gopurs_runtime.Value, g_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_verbJsonType__1594722986(Rebox_Data_Argonaut_Core_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](def_0_box)), f_unused_1_box, g_unused_2_box)
		})
	})
	return cache_Data_Argonaut_Core_verbJsonType__1594722986
}

var cache_Data_Argonaut_Core_caseJsonNull gopurs_runtime.Value
var once_Data_Argonaut_Core_caseJsonNull sync.Once

func Get_Data_Argonaut_Core_caseJsonNull() gopurs_runtime.Value {
	once_Data_Argonaut_Core_caseJsonNull.Do(func() {
		cache_Data_Argonaut_Core_caseJsonNull = gopurs_runtime.Func3(func(d_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, j_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_caseJsonNull(d_0_box, f_1_box, j_2_box)
		})
	})
	return cache_Data_Argonaut_Core_caseJsonNull
}

var cache_Data_Argonaut_Core_verbJsonType__601715556 gopurs_runtime.Value
var once_Data_Argonaut_Core_verbJsonType__601715556 sync.Once

func Get_Data_Argonaut_Core_verbJsonType__601715556() gopurs_runtime.Value {
	once_Data_Argonaut_Core_verbJsonType__601715556.Do(func() {
		cache_Data_Argonaut_Core_verbJsonType__601715556 = gopurs_runtime.Func3(func(def_0_box gopurs_runtime.Value, f_unused_1_box gopurs_runtime.Value, g_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_verbJsonType__601715556(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](def_0_box), f_unused_1_box, g_unused_2_box)
		})
	})
	return cache_Data_Argonaut_Core_verbJsonType__601715556
}

var cache_Data_Argonaut_Core_verbJsonType__2562148639 gopurs_runtime.Value
var once_Data_Argonaut_Core_verbJsonType__2562148639 sync.Once

func Get_Data_Argonaut_Core_verbJsonType__2562148639() gopurs_runtime.Value {
	once_Data_Argonaut_Core_verbJsonType__2562148639.Do(func() {
		cache_Data_Argonaut_Core_verbJsonType__2562148639 = gopurs_runtime.Func3(func(def_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_verbJsonType__2562148639((def_0_box.IntVal) != (0), f_1_box, g_unused_2_box)
		})
	})
	return cache_Data_Argonaut_Core_verbJsonType__2562148639
}

var cache_Data_Argonaut_Core_verbJsonType__3871162515 gopurs_runtime.Value
var once_Data_Argonaut_Core_verbJsonType__3871162515 sync.Once

func Get_Data_Argonaut_Core_verbJsonType__3871162515() gopurs_runtime.Value {
	once_Data_Argonaut_Core_verbJsonType__3871162515.Do(func() {
		cache_Data_Argonaut_Core_verbJsonType__3871162515 = gopurs_runtime.Func3(func(def_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_verbJsonType__3871162515((def_0_box.IntVal) != (0), f_1_box, g_unused_2_box)
		})
	})
	return cache_Data_Argonaut_Core_verbJsonType__3871162515
}

var cache_Data_Argonaut_Core_verbJsonType__2126563182 gopurs_runtime.Value
var once_Data_Argonaut_Core_verbJsonType__2126563182 sync.Once

func Get_Data_Argonaut_Core_verbJsonType__2126563182() gopurs_runtime.Value {
	once_Data_Argonaut_Core_verbJsonType__2126563182.Do(func() {
		cache_Data_Argonaut_Core_verbJsonType__2126563182 = gopurs_runtime.Func3(func(def_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_verbJsonType__2126563182((def_0_box.IntVal) != (0), f_1_box, g_unused_2_box)
		})
	})
	return cache_Data_Argonaut_Core_verbJsonType__2126563182
}

var cache_Data_Argonaut_Core_verbJsonType__575025737 gopurs_runtime.Value
var once_Data_Argonaut_Core_verbJsonType__575025737 sync.Once

func Get_Data_Argonaut_Core_verbJsonType__575025737() gopurs_runtime.Value {
	once_Data_Argonaut_Core_verbJsonType__575025737.Do(func() {
		cache_Data_Argonaut_Core_verbJsonType__575025737 = gopurs_runtime.Func3(func(def_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_verbJsonType__575025737((def_0_box.IntVal) != (0), f_1_box, g_unused_2_box)
		})
	})
	return cache_Data_Argonaut_Core_verbJsonType__575025737
}

var cache_Data_Argonaut_Core_verbJsonType__3433149055 gopurs_runtime.Value
var once_Data_Argonaut_Core_verbJsonType__3433149055 sync.Once

func Get_Data_Argonaut_Core_verbJsonType__3433149055() gopurs_runtime.Value {
	once_Data_Argonaut_Core_verbJsonType__3433149055.Do(func() {
		cache_Data_Argonaut_Core_verbJsonType__3433149055 = gopurs_runtime.Func3(func(def_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_verbJsonType__3433149055((def_0_box.IntVal) != (0), f_1_box, g_unused_2_box)
		})
	})
	return cache_Data_Argonaut_Core_verbJsonType__3433149055
}

var cache_Data_Argonaut_Core_verbJsonType__600469169 gopurs_runtime.Value
var once_Data_Argonaut_Core_verbJsonType__600469169 sync.Once

func Get_Data_Argonaut_Core_verbJsonType__600469169() gopurs_runtime.Value {
	once_Data_Argonaut_Core_verbJsonType__600469169.Do(func() {
		cache_Data_Argonaut_Core_verbJsonType__600469169 = gopurs_runtime.Func3(func(def_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_unused_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_verbJsonType__600469169((def_0_box.IntVal) != (0), f_1_box, g_unused_2_box)
		})
	})
	return cache_Data_Argonaut_Core_verbJsonType__600469169
}

var cache_Data_Argonaut_Core_toJsonType gopurs_runtime.Value
var once_Data_Argonaut_Core_toJsonType sync.Once

func Get_Data_Argonaut_Core_toJsonType() gopurs_runtime.Value {
	once_Data_Argonaut_Core_toJsonType.Do(func() {
		cache_Data_Argonaut_Core_toJsonType = gopurs_runtime.Func(func(g_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_toJsonType(g_0_box)
		})
	})
	return cache_Data_Argonaut_Core_toJsonType
}

var cache_Data_Argonaut_Core_toJsonType__1322236059 gopurs_runtime.Value
var once_Data_Argonaut_Core_toJsonType__1322236059 sync.Once

func Get_Data_Argonaut_Core_toJsonType__1322236059() gopurs_runtime.Value {
	once_Data_Argonaut_Core_toJsonType__1322236059.Do(func() {
		cache_Data_Argonaut_Core_toJsonType__1322236059 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_toJsonType__1322236059(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Data_Argonaut_Core_toJsonType__1322236059
}

var cache_Data_Argonaut_Core_toJsonType__3906794974 gopurs_runtime.Value
var once_Data_Argonaut_Core_toJsonType__3906794974 sync.Once

func Get_Data_Argonaut_Core_toJsonType__3906794974() gopurs_runtime.Value {
	once_Data_Argonaut_Core_toJsonType__3906794974.Do(func() {
		cache_Data_Argonaut_Core_toJsonType__3906794974 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_toJsonType__3906794974(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Data_Argonaut_Core_toJsonType__3906794974
}

var cache_Data_Argonaut_Core_toJsonType__1359231606 gopurs_runtime.Value
var once_Data_Argonaut_Core_toJsonType__1359231606 sync.Once

func Get_Data_Argonaut_Core_toJsonType__1359231606() gopurs_runtime.Value {
	once_Data_Argonaut_Core_toJsonType__1359231606.Do(func() {
		cache_Data_Argonaut_Core_toJsonType__1359231606 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_toJsonType__1359231606(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Data_Argonaut_Core_toJsonType__1359231606
}

var cache_Data_Argonaut_Core_toJsonType__2031900118 gopurs_runtime.Value
var once_Data_Argonaut_Core_toJsonType__2031900118 sync.Once

func Get_Data_Argonaut_Core_toJsonType__2031900118() gopurs_runtime.Value {
	once_Data_Argonaut_Core_toJsonType__2031900118.Do(func() {
		cache_Data_Argonaut_Core_toJsonType__2031900118 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_toJsonType__2031900118(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Data_Argonaut_Core_toJsonType__2031900118
}

var cache_Data_Argonaut_Core_toJsonType__917063126 gopurs_runtime.Value
var once_Data_Argonaut_Core_toJsonType__917063126 sync.Once

func Get_Data_Argonaut_Core_toJsonType__917063126() gopurs_runtime.Value {
	once_Data_Argonaut_Core_toJsonType__917063126.Do(func() {
		cache_Data_Argonaut_Core_toJsonType__917063126 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_toJsonType__917063126(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Data_Argonaut_Core_toJsonType__917063126
}

var cache_Data_Argonaut_Core_toJsonType__2083085451 gopurs_runtime.Value
var once_Data_Argonaut_Core_toJsonType__2083085451 sync.Once

func Get_Data_Argonaut_Core_toJsonType__2083085451() gopurs_runtime.Value {
	once_Data_Argonaut_Core_toJsonType__2083085451.Do(func() {
		cache_Data_Argonaut_Core_toJsonType__2083085451 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_toJsonType__2083085451(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Data_Argonaut_Core_toJsonType__2083085451
}

var cache_Data_Argonaut_Core_jsonZero gopurs_runtime.Value
var once_Data_Argonaut_Core_jsonZero sync.Once

func Get_Data_Argonaut_Core_jsonZero() gopurs_runtime.Value {
	once_Data_Argonaut_Core_jsonZero.Do(func() {
		cache_Data_Argonaut_Core_jsonZero = Data_Argonaut_Core_FromNumber_nativeWorker(gopurs_runtime.Float(0.0))
	})
	return cache_Data_Argonaut_Core_jsonZero
}

var cache_Data_Argonaut_Core_jsonTrue gopurs_runtime.Value
var once_Data_Argonaut_Core_jsonTrue sync.Once

func Get_Data_Argonaut_Core_jsonTrue() gopurs_runtime.Value {
	once_Data_Argonaut_Core_jsonTrue.Do(func() {
		cache_Data_Argonaut_Core_jsonTrue = Data_Argonaut_Core_FromBoolean_nativeWorker(gopurs_runtime.Bool(true))
	})
	return cache_Data_Argonaut_Core_jsonTrue
}

var cache_Data_Argonaut_Core_jsonSingletonObject gopurs_runtime.Value
var once_Data_Argonaut_Core_jsonSingletonObject sync.Once

func Get_Data_Argonaut_Core_jsonSingletonObject() gopurs_runtime.Value {
	once_Data_Argonaut_Core_jsonSingletonObject.Do(func() {
		cache_Data_Argonaut_Core_jsonSingletonObject = gopurs_runtime.Func2(func(key_0_box gopurs_runtime.Value, val_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_jsonSingletonObject(key_0_box.StrVal(), val_1_box)
		})
	})
	return cache_Data_Argonaut_Core_jsonSingletonObject
}

var cache_Data_Argonaut_Core_jsonSingletonArray gopurs_runtime.Value
var once_Data_Argonaut_Core_jsonSingletonArray sync.Once

func Get_Data_Argonaut_Core_jsonSingletonArray() gopurs_runtime.Value {
	once_Data_Argonaut_Core_jsonSingletonArray.Do(func() {
		cache_Data_Argonaut_Core_jsonSingletonArray = gopurs_runtime.Func(func(j_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_jsonSingletonArray(j_0_box)
		})
	})
	return cache_Data_Argonaut_Core_jsonSingletonArray
}

var cache_Data_Argonaut_Core_jsonFalse gopurs_runtime.Value
var once_Data_Argonaut_Core_jsonFalse sync.Once

func Get_Data_Argonaut_Core_jsonFalse() gopurs_runtime.Value {
	once_Data_Argonaut_Core_jsonFalse.Do(func() {
		cache_Data_Argonaut_Core_jsonFalse = Data_Argonaut_Core_FromBoolean_nativeWorker(gopurs_runtime.Bool(false))
	})
	return cache_Data_Argonaut_Core_jsonFalse
}

var cache_Data_Argonaut_Core_jsonEmptyString gopurs_runtime.Value
var once_Data_Argonaut_Core_jsonEmptyString sync.Once

func Get_Data_Argonaut_Core_jsonEmptyString() gopurs_runtime.Value {
	once_Data_Argonaut_Core_jsonEmptyString.Do(func() {
		cache_Data_Argonaut_Core_jsonEmptyString = Data_Argonaut_Core_FromString_nativeWorker(gopurs_runtime.StrWithCodeUnits("", 0))
	})
	return cache_Data_Argonaut_Core_jsonEmptyString
}

var cache_Data_Argonaut_Core_jsonEmptyObject gopurs_runtime.Value
var once_Data_Argonaut_Core_jsonEmptyObject sync.Once

func Get_Data_Argonaut_Core_jsonEmptyObject() gopurs_runtime.Value {
	once_Data_Argonaut_Core_jsonEmptyObject.Do(func() {
		cache_Data_Argonaut_Core_jsonEmptyObject = Data_Argonaut_Core_FromObject_nativeWorker(Get_Foreign_Object_empty())
	})
	return cache_Data_Argonaut_Core_jsonEmptyObject
}

var cache_Data_Argonaut_Core_jsonEmptyArray gopurs_runtime.Value
var once_Data_Argonaut_Core_jsonEmptyArray sync.Once

func Get_Data_Argonaut_Core_jsonEmptyArray() gopurs_runtime.Value {
	once_Data_Argonaut_Core_jsonEmptyArray.Do(func() {
		cache_Data_Argonaut_Core_jsonEmptyArray = Data_Argonaut_Core_FromArray_nativeWorker(gopurs_runtime.Array([]gopurs_runtime.Value{}))
	})
	return cache_Data_Argonaut_Core_jsonEmptyArray
}

var cache_Data_Argonaut_Core_isJsonType gopurs_runtime.Value
var once_Data_Argonaut_Core_isJsonType sync.Once

func Get_Data_Argonaut_Core_isJsonType() gopurs_runtime.Value {
	once_Data_Argonaut_Core_isJsonType.Do(func() {
		cache_Data_Argonaut_Core_isJsonType = gopurs_runtime.Func(func(g_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_isJsonType(g_0_box)
		})
	})
	return cache_Data_Argonaut_Core_isJsonType
}

var cache_Data_Argonaut_Core_isJsonType__2816881787 gopurs_runtime.Value
var once_Data_Argonaut_Core_isJsonType__2816881787 sync.Once

func Get_Data_Argonaut_Core_isJsonType__2816881787() gopurs_runtime.Value {
	once_Data_Argonaut_Core_isJsonType__2816881787.Do(func() {
		cache_Data_Argonaut_Core_isJsonType__2816881787 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_isJsonType__2816881787(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Data_Argonaut_Core_isJsonType__2816881787
}

var cache_Data_Argonaut_Core_isJsonType__2120356990 gopurs_runtime.Value
var once_Data_Argonaut_Core_isJsonType__2120356990 sync.Once

func Get_Data_Argonaut_Core_isJsonType__2120356990() gopurs_runtime.Value {
	once_Data_Argonaut_Core_isJsonType__2120356990.Do(func() {
		cache_Data_Argonaut_Core_isJsonType__2120356990 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_isJsonType__2120356990(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Data_Argonaut_Core_isJsonType__2120356990
}

var cache_Data_Argonaut_Core_isJsonType__3607385046 gopurs_runtime.Value
var once_Data_Argonaut_Core_isJsonType__3607385046 sync.Once

func Get_Data_Argonaut_Core_isJsonType__3607385046() gopurs_runtime.Value {
	once_Data_Argonaut_Core_isJsonType__3607385046.Do(func() {
		cache_Data_Argonaut_Core_isJsonType__3607385046 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_isJsonType__3607385046(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Data_Argonaut_Core_isJsonType__3607385046
}

var cache_Data_Argonaut_Core_isJsonType__1394939478 gopurs_runtime.Value
var once_Data_Argonaut_Core_isJsonType__1394939478 sync.Once

func Get_Data_Argonaut_Core_isJsonType__1394939478() gopurs_runtime.Value {
	once_Data_Argonaut_Core_isJsonType__1394939478.Do(func() {
		cache_Data_Argonaut_Core_isJsonType__1394939478 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_isJsonType__1394939478(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Data_Argonaut_Core_isJsonType__1394939478
}

var cache_Data_Argonaut_Core_isJsonType__2126062102 gopurs_runtime.Value
var once_Data_Argonaut_Core_isJsonType__2126062102 sync.Once

func Get_Data_Argonaut_Core_isJsonType__2126062102() gopurs_runtime.Value {
	once_Data_Argonaut_Core_isJsonType__2126062102.Do(func() {
		cache_Data_Argonaut_Core_isJsonType__2126062102 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_isJsonType__2126062102(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Data_Argonaut_Core_isJsonType__2126062102
}

var cache_Data_Argonaut_Core_isJsonType__409139211 gopurs_runtime.Value
var once_Data_Argonaut_Core_isJsonType__409139211 sync.Once

func Get_Data_Argonaut_Core_isJsonType__409139211() gopurs_runtime.Value {
	once_Data_Argonaut_Core_isJsonType__409139211.Do(func() {
		cache_Data_Argonaut_Core_isJsonType__409139211 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_isJsonType__409139211(__eta_norm_0_unused_0_box)
		})
	})
	return cache_Data_Argonaut_Core_isJsonType__409139211
}

var cache_Data_Argonaut_Core_ordJson gopurs_runtime.Value
var once_Data_Argonaut_Core_ordJson sync.Once

func Get_Data_Argonaut_Core_ordJson() gopurs_runtime.Value {
	once_Data_Argonaut_Core_ordJson.Do(func() {
		cache_Data_Argonaut_Core_ordJson = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Core_eqJson()))}
		}), gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.UncurriedApp5(Get_Data_Argonaut_Core__compare(), gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, a_0, b_1)
		})}))}
	})
	return cache_Data_Argonaut_Core_ordJson
}

var cache_Data_Argonaut_Core_eqJson gopurs_runtime.Value
var once_Data_Argonaut_Core_eqJson sync.Once

func Get_Data_Argonaut_Core_eqJson() gopurs_runtime.Value {
	once_Data_Argonaut_Core_eqJson.Do(func() {
		cache_Data_Argonaut_Core_eqJson = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(j1_0 gopurs_runtime.Value, j2_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Argonaut_Core_ordJson()).V1, j1_0, j2_1)
			_ = __t_tag_0
			return gopurs_runtime.Bool((uint32(__t_tag_0.IntVal) == 902936544))
		})}))}
	})
	return cache_Data_Argonaut_Core_eqJson
}

var cache_Data_Argonaut_Core_eqJNull gopurs_runtime.Value
var once_Data_Argonaut_Core_eqJNull sync.Once

func Get_Data_Argonaut_Core_eqJNull() gopurs_runtime.Value {
	once_Data_Argonaut_Core_eqJNull.Do(func() {
		cache_Data_Argonaut_Core_eqJNull = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		})}))}
	})
	return cache_Data_Argonaut_Core_eqJNull
}

var cache_Data_Argonaut_Core_ordJNull gopurs_runtime.Value
var once_Data_Argonaut_Core_ordJNull sync.Once

func Get_Data_Argonaut_Core_ordJNull() gopurs_runtime.Value {
	once_Data_Argonaut_Core_ordJNull.Do(func() {
		cache_Data_Argonaut_Core_ordJNull = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Argonaut_Core_eqJNull()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
		})}))}
	})
	return cache_Data_Argonaut_Core_ordJNull
}

var cache_Data_Argonaut_Core_isString gopurs_runtime.Value
var once_Data_Argonaut_Core_isString sync.Once

func Get_Data_Argonaut_Core_isString() gopurs_runtime.Value {
	once_Data_Argonaut_Core_isString.Do(func() {
		cache_Data_Argonaut_Core_isString = Call_Data_Argonaut_Core_isJsonType__2126062102(Get_Data_Argonaut_Core_caseJsonString())
	})
	return cache_Data_Argonaut_Core_isString
}

var cache_Data_Argonaut_Core_toString gopurs_runtime.Value
var once_Data_Argonaut_Core_toString sync.Once

func Get_Data_Argonaut_Core_toString() gopurs_runtime.Value {
	once_Data_Argonaut_Core_toString.Do(func() {
		cache_Data_Argonaut_Core_toString = Call_Data_Argonaut_Core_toJsonType__917063126(Get_Data_Argonaut_Core_caseJsonString())
	})
	return cache_Data_Argonaut_Core_toString
}

var cache_Data_Argonaut_Core_isObject gopurs_runtime.Value
var once_Data_Argonaut_Core_isObject sync.Once

func Get_Data_Argonaut_Core_isObject() gopurs_runtime.Value {
	once_Data_Argonaut_Core_isObject.Do(func() {
		cache_Data_Argonaut_Core_isObject = Call_Data_Argonaut_Core_isJsonType__2816881787(Get_Data_Argonaut_Core_caseJsonObject())
	})
	return cache_Data_Argonaut_Core_isObject
}

var cache_Data_Argonaut_Core_toObject gopurs_runtime.Value
var once_Data_Argonaut_Core_toObject sync.Once

func Get_Data_Argonaut_Core_toObject() gopurs_runtime.Value {
	once_Data_Argonaut_Core_toObject.Do(func() {
		cache_Data_Argonaut_Core_toObject = Call_Data_Argonaut_Core_toJsonType__1322236059(Get_Data_Argonaut_Core_caseJsonObject())
	})
	return cache_Data_Argonaut_Core_toObject
}

var cache_Data_Argonaut_Core_isNumber gopurs_runtime.Value
var once_Data_Argonaut_Core_isNumber sync.Once

func Get_Data_Argonaut_Core_isNumber() gopurs_runtime.Value {
	once_Data_Argonaut_Core_isNumber.Do(func() {
		cache_Data_Argonaut_Core_isNumber = Call_Data_Argonaut_Core_isJsonType__1394939478(Get_Data_Argonaut_Core_caseJsonNumber())
	})
	return cache_Data_Argonaut_Core_isNumber
}

var cache_Data_Argonaut_Core_toNumber gopurs_runtime.Value
var once_Data_Argonaut_Core_toNumber sync.Once

func Get_Data_Argonaut_Core_toNumber() gopurs_runtime.Value {
	once_Data_Argonaut_Core_toNumber.Do(func() {
		cache_Data_Argonaut_Core_toNumber = Call_Data_Argonaut_Core_toJsonType__2031900118(Get_Data_Argonaut_Core_caseJsonNumber())
	})
	return cache_Data_Argonaut_Core_toNumber
}

var cache_Data_Argonaut_Core_isNull gopurs_runtime.Value
var once_Data_Argonaut_Core_isNull sync.Once

func Get_Data_Argonaut_Core_isNull() gopurs_runtime.Value {
	once_Data_Argonaut_Core_isNull.Do(func() {
		cache_Data_Argonaut_Core_isNull = Call_Data_Argonaut_Core_isJsonType__409139211(Get_Data_Argonaut_Core_caseJsonNull())
	})
	return cache_Data_Argonaut_Core_isNull
}

var cache_Data_Argonaut_Core_toNull gopurs_runtime.Value
var once_Data_Argonaut_Core_toNull sync.Once

func Get_Data_Argonaut_Core_toNull() gopurs_runtime.Value {
	once_Data_Argonaut_Core_toNull.Do(func() {
		cache_Data_Argonaut_Core_toNull = Call_Data_Argonaut_Core_toJsonType__2083085451(Get_Data_Argonaut_Core_caseJsonNull())
	})
	return cache_Data_Argonaut_Core_toNull
}

var cache_Data_Argonaut_Core_isBoolean gopurs_runtime.Value
var once_Data_Argonaut_Core_isBoolean sync.Once

func Get_Data_Argonaut_Core_isBoolean() gopurs_runtime.Value {
	once_Data_Argonaut_Core_isBoolean.Do(func() {
		cache_Data_Argonaut_Core_isBoolean = Call_Data_Argonaut_Core_isJsonType__3607385046(Get_Data_Argonaut_Core_caseJsonBoolean())
	})
	return cache_Data_Argonaut_Core_isBoolean
}

var cache_Data_Argonaut_Core_toBoolean gopurs_runtime.Value
var once_Data_Argonaut_Core_toBoolean sync.Once

func Get_Data_Argonaut_Core_toBoolean() gopurs_runtime.Value {
	once_Data_Argonaut_Core_toBoolean.Do(func() {
		cache_Data_Argonaut_Core_toBoolean = Call_Data_Argonaut_Core_toJsonType__1359231606(Get_Data_Argonaut_Core_caseJsonBoolean())
	})
	return cache_Data_Argonaut_Core_toBoolean
}

var cache_Data_Argonaut_Core_isArray gopurs_runtime.Value
var once_Data_Argonaut_Core_isArray sync.Once

func Get_Data_Argonaut_Core_isArray() gopurs_runtime.Value {
	once_Data_Argonaut_Core_isArray.Do(func() {
		cache_Data_Argonaut_Core_isArray = Call_Data_Argonaut_Core_isJsonType__2120356990(Get_Data_Argonaut_Core_caseJsonArray())
	})
	return cache_Data_Argonaut_Core_isArray
}

var cache_Data_Argonaut_Core_toArray gopurs_runtime.Value
var once_Data_Argonaut_Core_toArray sync.Once

func Get_Data_Argonaut_Core_toArray() gopurs_runtime.Value {
	once_Data_Argonaut_Core_toArray.Do(func() {
		cache_Data_Argonaut_Core_toArray = Call_Data_Argonaut_Core_toJsonType__3906794974(Get_Data_Argonaut_Core_caseJsonArray())
	})
	return cache_Data_Argonaut_Core_toArray
}

var cache_Data_Argonaut_Core_caseJson gopurs_runtime.Value
var once_Data_Argonaut_Core_caseJson sync.Once

func Get_Data_Argonaut_Core_caseJson() gopurs_runtime.Value {
	once_Data_Argonaut_Core_caseJson.Do(func() {
		cache_Data_Argonaut_Core_caseJson = gopurs_runtime.Func7(func(onNull_0_box gopurs_runtime.Value, onBool_1_box gopurs_runtime.Value, onNum_2_box gopurs_runtime.Value, onStr_3_box gopurs_runtime.Value, onArr_4_box gopurs_runtime.Value, onObj_5_box gopurs_runtime.Value, j_6_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Argonaut_Core_caseJson(onNull_0_box, onBool_1_box, onNum_2_box, onStr_3_box, onArr_4_box, onObj_5_box, j_6_box)
		})
	})
	return cache_Data_Argonaut_Core_caseJson
}

func Call_Data_Argonaut_Core_verbJsonType(def_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var def_0 gopurs_runtime.Value = def_0_loop
	_ = def_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var g_2 gopurs_runtime.Value = g_2_loop
	_ = g_2
	return gopurs_runtime.Apply2(g_2, def_0, f_1)
}

func Call_Data_Argonaut_Core_caseJsonObject(d_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, j_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var d_0 gopurs_runtime.Value = d_0_loop
	_ = d_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var j_2 gopurs_runtime.Value = j_2_loop
	_ = j_2
	return Data_Argonaut_Core__CaseJsonObject(d_0, f_1, j_2)
}

func Call_Data_Argonaut_Core_verbJsonType__1588320074(def_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], f_unused_1_loop gopurs_runtime.Value, g_unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
verbJsonType__1588320074:
	for {
		if false {
			continue verbJsonType__1588320074
		}
		var def_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = def_0_loop
		_ = def_0
		var f_unused_1 gopurs_runtime.Value = f_unused_1_loop
		_ = f_unused_1
		var g_unused_2 gopurs_runtime.Value = g_unused_2_loop
		_ = g_unused_2
		return gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonObject(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(def_0)}, Get_Data_Maybe_Just())
	}
}

func Call_Data_Argonaut_Core_caseJsonArray(d_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, j_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var d_0 gopurs_runtime.Value = d_0_loop
	_ = d_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var j_2 gopurs_runtime.Value = j_2_loop
	_ = j_2
	return Data_Argonaut_Core__CaseJsonArray(d_0, f_1, j_2)
}

func Call_Data_Argonaut_Core_verbJsonType__286058406(def_0_loop *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value], f_unused_1_loop gopurs_runtime.Value, g_unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
verbJsonType__286058406:
	for {
		if false {
			continue verbJsonType__286058406
		}
		var def_0 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value] = def_0_loop
		_ = def_0
		var f_unused_1 gopurs_runtime.Value = f_unused_1_loop
		_ = f_unused_1
		var g_unused_2 gopurs_runtime.Value = g_unused_2_loop
		_ = g_unused_2
		return gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonArray(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Core_2807397954_3094389156(def_0))}, Get_Data_Maybe_Just())
	}
}

func Call_Data_Argonaut_Core_caseJsonBoolean(d_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, j_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var d_0 gopurs_runtime.Value = d_0_loop
	_ = d_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var j_2 gopurs_runtime.Value = j_2_loop
	_ = j_2
	return Data_Argonaut_Core__CaseJsonBoolean(d_0, f_1, j_2)
}

func Call_Data_Argonaut_Core_verbJsonType__156013339(def_0_loop *Constructor_Data_Maybe_Just[bool], f_unused_1_loop gopurs_runtime.Value, g_unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
verbJsonType__156013339:
	for {
		if false {
			continue verbJsonType__156013339
		}
		var def_0 *Constructor_Data_Maybe_Just[bool] = def_0_loop
		_ = def_0
		var f_unused_1 gopurs_runtime.Value = f_unused_1_loop
		_ = f_unused_1
		var g_unused_2 gopurs_runtime.Value = g_unused_2_loop
		_ = g_unused_2
		return gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonBoolean(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Core_1044698560_3094389156(def_0))}, Get_Data_Maybe_Just())
	}
}

func Call_Data_Argonaut_Core_caseJsonNumber(d_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, j_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var d_0 gopurs_runtime.Value = d_0_loop
	_ = d_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var j_2 gopurs_runtime.Value = j_2_loop
	_ = j_2
	return Data_Argonaut_Core__CaseJsonNumber(d_0, f_1, j_2)
}

func Call_Data_Argonaut_Core_verbJsonType__425335836(def_0_loop *Constructor_Data_Maybe_Just[float64], f_unused_1_loop gopurs_runtime.Value, g_unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
verbJsonType__425335836:
	for {
		if false {
			continue verbJsonType__425335836
		}
		var def_0 *Constructor_Data_Maybe_Just[float64] = def_0_loop
		_ = def_0
		var f_unused_1 gopurs_runtime.Value = f_unused_1_loop
		_ = f_unused_1
		var g_unused_2 gopurs_runtime.Value = g_unused_2_loop
		_ = g_unused_2
		return gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonNumber(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Core_3240988860_3094389156(def_0))}, Get_Data_Maybe_Just())
	}
}

func Call_Data_Argonaut_Core_caseJsonString(d_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, j_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var d_0 gopurs_runtime.Value = d_0_loop
	_ = d_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var j_2 gopurs_runtime.Value = j_2_loop
	_ = j_2
	return Data_Argonaut_Core__CaseJsonString(d_0, f_1, j_2)
}

func Call_Data_Argonaut_Core_verbJsonType__1594722986(def_0_loop *Constructor_Data_Maybe_Just[string], f_unused_1_loop gopurs_runtime.Value, g_unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
verbJsonType__1594722986:
	for {
		if false {
			continue verbJsonType__1594722986
		}
		var def_0 *Constructor_Data_Maybe_Just[string] = def_0_loop
		_ = def_0
		var f_unused_1 gopurs_runtime.Value = f_unused_1_loop
		_ = f_unused_1
		var g_unused_2 gopurs_runtime.Value = g_unused_2_loop
		_ = g_unused_2
		return gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonString(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Argonaut_Core_742090555_3094389156(def_0))}, Get_Data_Maybe_Just())
	}
}

func Call_Data_Argonaut_Core_caseJsonNull(d_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, j_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var d_0 gopurs_runtime.Value = d_0_loop
	_ = d_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var j_2 gopurs_runtime.Value = j_2_loop
	_ = j_2
	return Data_Argonaut_Core__CaseJsonNull(d_0, f_1, j_2)
}

func Call_Data_Argonaut_Core_verbJsonType__601715556(def_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], f_unused_1_loop gopurs_runtime.Value, g_unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
verbJsonType__601715556:
	for {
		if false {
			continue verbJsonType__601715556
		}
		var def_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = def_0_loop
		_ = def_0
		var f_unused_1 gopurs_runtime.Value = f_unused_1_loop
		_ = f_unused_1
		var g_unused_2 gopurs_runtime.Value = g_unused_2_loop
		_ = g_unused_2
		return gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonNull(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(def_0)}, Get_Data_Maybe_Just())
	}
}

func Call_Data_Argonaut_Core_verbJsonType__2562148639(def_0_loop bool, f_1_loop gopurs_runtime.Value, g_unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
verbJsonType__2562148639:
	for {
		if false {
			continue verbJsonType__2562148639
		}
		var def_0 bool = def_0_loop
		_ = def_0
		var f_1 gopurs_runtime.Value = f_1_loop
		_ = f_1
		var g_unused_2 gopurs_runtime.Value = g_unused_2_loop
		_ = g_unused_2
		return gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonObject(), gopurs_runtime.Bool(def_0), f_1)
	}
}

func Call_Data_Argonaut_Core_verbJsonType__3871162515(def_0_loop bool, f_1_loop gopurs_runtime.Value, g_unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
verbJsonType__3871162515:
	for {
		if false {
			continue verbJsonType__3871162515
		}
		var def_0 bool = def_0_loop
		_ = def_0
		var f_1 gopurs_runtime.Value = f_1_loop
		_ = f_1
		var g_unused_2 gopurs_runtime.Value = g_unused_2_loop
		_ = g_unused_2
		return gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonArray(), gopurs_runtime.Bool(def_0), f_1)
	}
}

func Call_Data_Argonaut_Core_verbJsonType__2126563182(def_0_loop bool, f_1_loop gopurs_runtime.Value, g_unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
verbJsonType__2126563182:
	for {
		if false {
			continue verbJsonType__2126563182
		}
		var def_0 bool = def_0_loop
		_ = def_0
		var f_1 gopurs_runtime.Value = f_1_loop
		_ = f_1
		var g_unused_2 gopurs_runtime.Value = g_unused_2_loop
		_ = g_unused_2
		return gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonBoolean(), gopurs_runtime.Bool(def_0), f_1)
	}
}

func Call_Data_Argonaut_Core_verbJsonType__575025737(def_0_loop bool, f_1_loop gopurs_runtime.Value, g_unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
verbJsonType__575025737:
	for {
		if false {
			continue verbJsonType__575025737
		}
		var def_0 bool = def_0_loop
		_ = def_0
		var f_1 gopurs_runtime.Value = f_1_loop
		_ = f_1
		var g_unused_2 gopurs_runtime.Value = g_unused_2_loop
		_ = g_unused_2
		return gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonNumber(), gopurs_runtime.Bool(def_0), f_1)
	}
}

func Call_Data_Argonaut_Core_verbJsonType__3433149055(def_0_loop bool, f_1_loop gopurs_runtime.Value, g_unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
verbJsonType__3433149055:
	for {
		if false {
			continue verbJsonType__3433149055
		}
		var def_0 bool = def_0_loop
		_ = def_0
		var f_1 gopurs_runtime.Value = f_1_loop
		_ = f_1
		var g_unused_2 gopurs_runtime.Value = g_unused_2_loop
		_ = g_unused_2
		return gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonString(), gopurs_runtime.Bool(def_0), f_1)
	}
}

func Call_Data_Argonaut_Core_verbJsonType__600469169(def_0_loop bool, f_1_loop gopurs_runtime.Value, g_unused_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
verbJsonType__600469169:
	for {
		if false {
			continue verbJsonType__600469169
		}
		var def_0 bool = def_0_loop
		_ = def_0
		var f_1 gopurs_runtime.Value = f_1_loop
		_ = f_1
		var g_unused_2 gopurs_runtime.Value = g_unused_2_loop
		_ = g_unused_2
		return gopurs_runtime.Apply2(Get_Data_Argonaut_Core_caseJsonNull(), gopurs_runtime.Bool(def_0), f_1)
	}
}

func Call_Data_Argonaut_Core_toJsonType(g_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var g_0 gopurs_runtime.Value = g_0_loop
	_ = g_0
	return gopurs_runtime.Apply2(g_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, Get_Data_Maybe_Just())
}

func Call_Data_Argonaut_Core_toJsonType__1322236059(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
toJsonType__1322236059:
	for {
		if false {
			continue toJsonType__1322236059
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return Call_Data_Argonaut_Core_verbJsonType__1588320074(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}), Get_Data_Maybe_Just(), Get_Data_Argonaut_Core_caseJsonObject())
	}
}

func Call_Data_Argonaut_Core_toJsonType__3906794974(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
toJsonType__3906794974:
	for {
		if false {
			continue toJsonType__3906794974
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return Call_Data_Argonaut_Core_verbJsonType__286058406(Rebox_Data_Argonaut_Core_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})), Get_Data_Maybe_Just(), Get_Data_Argonaut_Core_caseJsonArray())
	}
}

func Call_Data_Argonaut_Core_toJsonType__1359231606(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
toJsonType__1359231606:
	for {
		if false {
			continue toJsonType__1359231606
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return Call_Data_Argonaut_Core_verbJsonType__156013339(Rebox_Data_Argonaut_Core_3094389156_1044698560(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})), Get_Data_Maybe_Just(), Get_Data_Argonaut_Core_caseJsonBoolean())
	}
}

func Call_Data_Argonaut_Core_toJsonType__2031900118(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
toJsonType__2031900118:
	for {
		if false {
			continue toJsonType__2031900118
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return Call_Data_Argonaut_Core_verbJsonType__425335836(Rebox_Data_Argonaut_Core_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})), Get_Data_Maybe_Just(), Get_Data_Argonaut_Core_caseJsonNumber())
	}
}

func Call_Data_Argonaut_Core_toJsonType__917063126(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
toJsonType__917063126:
	for {
		if false {
			continue toJsonType__917063126
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return Call_Data_Argonaut_Core_verbJsonType__1594722986(Rebox_Data_Argonaut_Core_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})), Get_Data_Maybe_Just(), Get_Data_Argonaut_Core_caseJsonString())
	}
}

func Call_Data_Argonaut_Core_toJsonType__2083085451(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
toJsonType__2083085451:
	for {
		if false {
			continue toJsonType__2083085451
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return Call_Data_Argonaut_Core_verbJsonType__601715556(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}), Get_Data_Maybe_Just(), Get_Data_Argonaut_Core_caseJsonNull())
	}
}

func Call_Data_Argonaut_Core_jsonSingletonObject(key_0_loop string, val_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var key_0 string = key_0_loop
	_ = key_0
	var val_1 gopurs_runtime.Value = val_1_loop
	_ = val_1
	return Data_Argonaut_Core_FromObject_nativeWorker(Call_Foreign_Object_singleton__2505508685(key_0, val_1))
}

func Call_Data_Argonaut_Core_jsonSingletonArray(j_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var j_0 gopurs_runtime.Value = j_0_loop
	_ = j_0
	return Data_Argonaut_Core_FromArray_nativeWorker(gopurs_runtime.Array([]gopurs_runtime.Value{j_0}))
}

func Call_Data_Argonaut_Core_isJsonType(g_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var g_0 gopurs_runtime.Value = g_0_loop
	_ = g_0
	return gopurs_runtime.Apply2(g_0, gopurs_runtime.Bool(false), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Bool(true)
	}))
}

func Call_Data_Argonaut_Core_isJsonType__2816881787(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
isJsonType__2816881787:
	for {
		if false {
			continue isJsonType__2816881787
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return Call_Data_Argonaut_Core_verbJsonType__2562148639(false, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		}), Get_Data_Argonaut_Core_caseJsonObject())
	}
}

func Call_Data_Argonaut_Core_isJsonType__2120356990(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
isJsonType__2120356990:
	for {
		if false {
			continue isJsonType__2120356990
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return Call_Data_Argonaut_Core_verbJsonType__3871162515(false, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		}), Get_Data_Argonaut_Core_caseJsonArray())
	}
}

func Call_Data_Argonaut_Core_isJsonType__3607385046(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
isJsonType__3607385046:
	for {
		if false {
			continue isJsonType__3607385046
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return Call_Data_Argonaut_Core_verbJsonType__2126563182(false, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		}), Get_Data_Argonaut_Core_caseJsonBoolean())
	}
}

func Call_Data_Argonaut_Core_isJsonType__1394939478(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
isJsonType__1394939478:
	for {
		if false {
			continue isJsonType__1394939478
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return Call_Data_Argonaut_Core_verbJsonType__575025737(false, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		}), Get_Data_Argonaut_Core_caseJsonNumber())
	}
}

func Call_Data_Argonaut_Core_isJsonType__2126062102(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
isJsonType__2126062102:
	for {
		if false {
			continue isJsonType__2126062102
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return Call_Data_Argonaut_Core_verbJsonType__3433149055(false, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		}), Get_Data_Argonaut_Core_caseJsonString())
	}
}

func Call_Data_Argonaut_Core_isJsonType__409139211(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
isJsonType__409139211:
	for {
		if false {
			continue isJsonType__409139211
		}
		var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return Call_Data_Argonaut_Core_verbJsonType__600469169(false, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(true)
		}), Get_Data_Argonaut_Core_caseJsonNull())
	}
}

func Call_Data_Argonaut_Core_caseJson(onNull_0_loop gopurs_runtime.Value, onBool_1_loop gopurs_runtime.Value, onNum_2_loop gopurs_runtime.Value, onStr_3_loop gopurs_runtime.Value, onArr_4_loop gopurs_runtime.Value, onObj_5_loop gopurs_runtime.Value, j_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var onNull_0 gopurs_runtime.Value = onNull_0_loop
	_ = onNull_0
	var onBool_1 gopurs_runtime.Value = onBool_1_loop
	_ = onBool_1
	var onNum_2 gopurs_runtime.Value = onNum_2_loop
	_ = onNum_2
	var onStr_3 gopurs_runtime.Value = onStr_3_loop
	_ = onStr_3
	var onArr_4 gopurs_runtime.Value = onArr_4_loop
	_ = onArr_4
	var onObj_5 gopurs_runtime.Value = onObj_5_loop
	_ = onObj_5
	var j_6 gopurs_runtime.Value = j_6_loop
	_ = j_6
	return Data_Argonaut_Core_CaseJsonImpl(onNull_0, onBool_1, onNum_2, onStr_3, onArr_4, onObj_5, j_6)
}

func Rebox_Data_Argonaut_Core_1044698560_3094389156(in *Constructor_Data_Maybe_Just[bool]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Bool(in.V0)
	return out
}

func Rebox_Data_Argonaut_Core_2807397954_3094389156(in *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Array(in.V0)
	return out
}

func Rebox_Data_Argonaut_Core_3094389156_1044698560(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[bool]{}
	out.V0 = (in.V0.IntVal) != (0)
	return out
}

func Rebox_Data_Argonaut_Core_3094389156_2807397954(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]{}
	out.V0 = *((*[]gopurs_runtime.Value)(in.V0.UnsafePtr))
	return out
}

func Rebox_Data_Argonaut_Core_3094389156_3240988860(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[float64]{}
	out.V0 = in.V0.FloatVal()
	return out
}

func Rebox_Data_Argonaut_Core_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_Argonaut_Core_3240988860_3094389156(in *Constructor_Data_Maybe_Just[float64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Float(in.V0)
	return out
}

func Rebox_Data_Argonaut_Core_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}

func Get_Data_Argonaut_Core__caseJsonArray() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core__CaseJsonArray
}

func Get_Data_Argonaut_Core__caseJsonBoolean() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core__CaseJsonBoolean
}

func Get_Data_Argonaut_Core__caseJsonNull() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core__CaseJsonNull
}

func Get_Data_Argonaut_Core__caseJsonNumber() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core__CaseJsonNumber
}

func Get_Data_Argonaut_Core__caseJsonObject() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core__CaseJsonObject
}

func Get_Data_Argonaut_Core__caseJsonString() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core__CaseJsonString
}

func Get_Data_Argonaut_Core__compare() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core__Compare
}

func Get_Data_Argonaut_Core_caseJsonImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core_CaseJsonImpl
}

func Get_Data_Argonaut_Core_fromArray() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core_FromArray
}

func Get_Data_Argonaut_Core_fromBoolean() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core_FromBoolean
}

func Get_Data_Argonaut_Core_fromNumber() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core_FromNumber
}

func Get_Data_Argonaut_Core_fromObject() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core_FromObject
}

func Get_Data_Argonaut_Core_fromString() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core_FromString
}

func Get_Data_Argonaut_Core_jsonNull() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core_JsonNull
}

func Get_Data_Argonaut_Core_stringify() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core_Stringify
}

func Get_Data_Argonaut_Core_stringifyWithIndent() gopurs_runtime.Value {
	return _Gopurs_Data_Argonaut_Core_StringifyWithIndent
}
