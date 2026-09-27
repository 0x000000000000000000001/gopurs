package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqArray
}

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(Rebox_Main_1386611502_1469227923(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))})))))}
	})
	return cache_Main_showArray
}

var cache_Main_checkConsumer gopurs_runtime.Value
var once_Main_checkConsumer sync.Once

func Get_Main_checkConsumer() gopurs_runtime.Value {
	once_Main_checkConsumer.Do(func() {
		cache_Main_checkConsumer = gopurs_runtime.Func(func(consume_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkConsumer(consume_0_box)
		})
	})
	return cache_Main_checkConsumer
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Main_ReturnInt64(int64(0)), int64(0)})
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Main_ReturnInt64(int64(-7)), int64(-7)}), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Main_ReturnInt64(int64(42)), int64(42)}), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Main_ReturnInt64(int64(-2147483648)), int64(-2147483648)}), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Main_ReturnInt64(int64(2147483647)), int64(2147483647)}), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{gopurs_runtime.Apply(Get_Main_returnInt(), gopurs_runtime.Int(int64(0))).IntVal, int64(0)}), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{gopurs_runtime.Apply(Get_Main_returnInt(), gopurs_runtime.Int(int64(-7))).IntVal, int64(-7)}), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{gopurs_runtime.Apply(Get_Main_returnInt(), gopurs_runtime.Int(int64(42))).IntVal, int64(42)}), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{gopurs_runtime.Apply(Get_Main_returnInt(), gopurs_runtime.Int(int64(-2147483648))).IntVal, int64(-2147483648)}), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{gopurs_runtime.Apply(Get_Main_returnInt(), gopurs_runtime.Int(int64(2147483647))).IntVal, int64(2147483647)}), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
				actual   []int64
				expected []int64
			}{Main_ReturnInt64Array([]int64{}), []int64{}}), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
				actual   []int64
				expected []int64
			}{Main_ReturnInt64Array([]int64{int64(0), int64(-7), int64(42), int64(-2147483648), int64(2147483647)}), []int64{int64(0), int64(-7), int64(42), int64(-2147483648), int64(2147483647)}}), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
				actual   []int64
				expected []int64
			}{func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Main_returnIntArray(), func() gopurs_runtime.Value {
					arr := []int64{}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), []int64{}}), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1772858129("", struct {
				actual   []int64
				expected []int64
			}{func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Main_returnIntArray(), func() gopurs_runtime.Value {
					arr := []int64{int64(0), int64(-7), int64(42), int64(-2147483648), int64(2147483647)}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), []int64{int64(0), int64(-7), int64(42), int64(-2147483648), int64(2147483647)}}), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Main_DynamicInt64_nativeWorker(int64(0)).IntVal, int64(0)}), gopurs_runtime.Value{})
			_ = __local_var_15_15
			__local_var_16_16 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Main_DynamicInt64_nativeWorker(int64(-2147483648)).IntVal, int64(-2147483648)}), gopurs_runtime.Value{})
			_ = __local_var_16_16
			__local_var_17_17 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Main_DynamicInt64_nativeWorker(int64(2147483647)).IntVal, int64(2147483647)}), gopurs_runtime.Value{})
			_ = __local_var_17_17
			__local_var_18_18 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{gopurs_runtime.Apply(Get_Main_dynamicInt(), gopurs_runtime.Int(int64(-7))).IntVal, int64(-7)}), gopurs_runtime.Value{})
			_ = __local_var_18_18
			__local_var_19_19 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{gopurs_runtime.Apply(Get_Main_dynamicInt(), gopurs_runtime.Int(int64(-2147483648))).IntVal, int64(-2147483648)}), gopurs_runtime.Value{})
			_ = __local_var_19_19
			__local_var_20_20 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{gopurs_runtime.Apply(Get_Main_dynamicInt(), gopurs_runtime.Int(int64(2147483647))).IntVal, int64(2147483647)}), gopurs_runtime.Value{})
			_ = __local_var_20_20
			__local_var_21_21 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___4176622598("", struct {
				actual   string
				expected string
			}{Main_DynamicString_nativeWorker("fallback").StrVal(), "fallback"}), gopurs_runtime.Value{})
			_ = __local_var_21_21
			__local_var_22_22 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(Main_DynamicBoolean_nativeWorker(false).IntVal) != (0), false}), gopurs_runtime.Value{})
			_ = __local_var_22_22
			__local_var_23_23 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___1138829510("", struct {
				actual   bool
				expected bool
			}{(Main_DynamicBoolean_nativeWorker(true).IntVal) != (0), true}), gopurs_runtime.Value{})
			_ = __local_var_23_23
			__local_var_24_24 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Func(func(value_24 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(((value_24.IntVal) * (value_24.IntVal)) + (int64(1)))
			})), gopurs_runtime.Value{})
			_ = __local_var_24_24
			__local_var_25_25 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_24_24), gopurs_runtime.Value{})
			_ = __local_var_25_25
			__local_var_26_26 := gopurs_runtime.Apply(Call_Main_checkConsumer(__local_var_25_25), gopurs_runtime.Value{})
			_ = __local_var_26_26
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_checkConsumer(consume_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var consume_0 gopurs_runtime.Value = consume_0_loop
	_ = consume_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
		__local_var_1_0 := Call_Test_Assert_assertEqual_prime___627669702("", struct {
			actual   int64
			expected int64
		}{gopurs_runtime.Apply(consume_0, gopurs_runtime.Int(Main_ReturnInt64(int64(-7)))).IntVal, int64(50)})
		_ = __local_var_1_0
		__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = __local_var_2_1
		__local_var_3_2 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
			actual   int64
			expected int64
		}{gopurs_runtime.Apply(consume_0, gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_returnInt(), gopurs_runtime.Int(int64(5))).IntVal)).IntVal, int64(26)}), gopurs_runtime.Value{})
		_ = __local_var_3_2
		__local_var_4_3 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
			actual   int64
			expected int64
		}{gopurs_runtime.Apply(consume_0, gopurs_runtime.Int(Main_DynamicInt64_nativeWorker(int64(-3)).IntVal)).IntVal, int64(10)}), gopurs_runtime.Value{})
		_ = __local_var_4_3
		return gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
			actual   int64
			expected int64
		}{gopurs_runtime.Apply(consume_0, gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_dynamicInt(), gopurs_runtime.Int(int64(0))).IntVal)).IntVal, int64(1)}), gopurs_runtime.Value{})
	})
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1469227923(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[[]int64])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Main_1469227923_1386611502(in *Constructor_Data_Show_Show[[]int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_378698611_3790796878(in *Constructor_Data_Eq_Eq[[]int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_378698611(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]int64])(unsafe.Pointer(in))
}

func Get_Main_dynamicBoolean() gopurs_runtime.Value {
	return _Gopurs_Main_DynamicBoolean
}

func Get_Main_dynamicInt() gopurs_runtime.Value {
	return _Gopurs_Main_DynamicInt
}

func Get_Main_dynamicInt64() gopurs_runtime.Value {
	return _Gopurs_Main_DynamicInt64
}

func Get_Main_dynamicString() gopurs_runtime.Value {
	return _Gopurs_Main_DynamicString
}

func Get_Main_returnInt() gopurs_runtime.Value {
	return _Gopurs_Main_ReturnInt
}

func Get_Main_returnInt64() gopurs_runtime.Value {
	return _Gopurs_Main_ReturnInt64
}

func Get_Main_returnInt64Array() gopurs_runtime.Value {
	return _Gopurs_Main_ReturnInt64Array
}

func Get_Main_returnIntArray() gopurs_runtime.Value {
	return _Gopurs_Main_ReturnIntArray
}
