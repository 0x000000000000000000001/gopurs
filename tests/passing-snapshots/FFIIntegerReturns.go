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
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878((&Constructor_Data_Eq_Eq[[]int64]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl())})))}
	})
	return cache_Main_eqArray
}

var cache_Main_showArray gopurs_runtime.Value
var once_Main_showArray sync.Once

func Get_Main_showArray() gopurs_runtime.Value {
	once_Main_showArray.Do(func() {
		cache_Main_showArray = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502((&Constructor_Data_Show_Show[[]int64]{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), Get_Data_Show_showIntImpl())})))}
	})
	return cache_Main_showArray
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_returnInt64(), gopurs_runtime.Int(int64(0))).IntVal), gopurs_runtime.Int(int64(0))))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_returnInt64(), gopurs_runtime.Int(int64(-7))).IntVal), gopurs_runtime.Int(int64(-7)))), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_returnInt64(), gopurs_runtime.Int(int64(42))).IntVal), gopurs_runtime.Int(int64(42)))), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_returnInt64(), gopurs_runtime.Int(int64(-2147483648))).IntVal), gopurs_runtime.Int(int64(-2147483648)))), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_returnInt64(), gopurs_runtime.Int(int64(2147483647))).IntVal), gopurs_runtime.Int(int64(2147483647)))), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_returnInt(), gopurs_runtime.Int(int64(0))).IntVal), gopurs_runtime.Int(int64(0)))), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_returnInt(), gopurs_runtime.Int(int64(-7))).IntVal), gopurs_runtime.Int(int64(-7)))), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_returnInt(), gopurs_runtime.Int(int64(42))).IntVal), gopurs_runtime.Int(int64(42)))), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_returnInt(), gopurs_runtime.Int(int64(-2147483648))).IntVal), gopurs_runtime.Int(int64(-2147483648)))), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_returnInt(), gopurs_runtime.Int(int64(2147483647))).IntVal), gopurs_runtime.Int(int64(2147483647)))), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]int64]](Get_Main_showArray())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
				arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Main_returnInt64Array(), func() gopurs_runtime.Value {
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
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := []int64{}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}())), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]int64]](Get_Main_showArray())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
				arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Main_returnInt64Array(), func() gopurs_runtime.Value {
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
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := []int64{int64(0), int64(-7), int64(42), int64(-2147483648), int64(2147483647)}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}())), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]int64]](Get_Main_showArray())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
				arr := func() []int64 {
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
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := []int64{}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}())), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]int64]](Get_Main_eqArray())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1469227923_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[[]int64]](Get_Main_showArray())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", func() gopurs_runtime.Value {
				arr := func() []int64 {
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
				}()
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}(), func() gopurs_runtime.Value {
				arr := []int64{int64(0), int64(-7), int64(42), int64(-2147483648), int64(2147483647)}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}())), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_dynamicInt64(), gopurs_runtime.Int(int64(0))).IntVal), gopurs_runtime.Int(int64(0)))), gopurs_runtime.Value{})
			_ = __local_var_15_15
			__local_var_16_16 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_dynamicInt64(), gopurs_runtime.Int(int64(-2147483648))).IntVal), gopurs_runtime.Int(int64(-2147483648)))), gopurs_runtime.Value{})
			_ = __local_var_16_16
			__local_var_17_17 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_dynamicInt64(), gopurs_runtime.Int(int64(2147483647))).IntVal), gopurs_runtime.Int(int64(2147483647)))), gopurs_runtime.Value{})
			_ = __local_var_17_17
			__local_var_18_18 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_dynamicInt(), gopurs_runtime.Int(int64(-7))).IntVal), gopurs_runtime.Int(int64(-7)))), gopurs_runtime.Value{})
			_ = __local_var_18_18
			__local_var_19_19 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_dynamicInt(), gopurs_runtime.Int(int64(-2147483648))).IntVal), gopurs_runtime.Int(int64(-2147483648)))), gopurs_runtime.Value{})
			_ = __local_var_19_19
			__local_var_20_20 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Get_Main_dynamicInt(), gopurs_runtime.Int(int64(2147483647))).IntVal), gopurs_runtime.Int(int64(2147483647)))), gopurs_runtime.Value{})
			_ = __local_var_20_20
			__local_var_21_21 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_Eq_eqString())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1514099793_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[string]](Get_Data_Show_showString())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Str(gopurs_runtime.Apply(Get_Main_dynamicString(), gopurs_runtime.Str("fallback")).StrVal()), gopurs_runtime.Str("fallback"))), gopurs_runtime.Value{})
			_ = __local_var_21_21
			__local_var_22_22 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[bool]](Get_Data_Eq_eqBoolean())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[bool]](Get_Data_Show_showBoolean())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Main_dynamicBoolean(), gopurs_runtime.Bool(false)).IntVal) != (0)), gopurs_runtime.Bool(false))), gopurs_runtime.Value{})
			_ = __local_var_22_22
			__local_var_23_23 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[bool]](Get_Data_Eq_eqBoolean())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[bool]](Get_Data_Show_showBoolean())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Main_dynamicBoolean(), gopurs_runtime.Bool(true)).IntVal) != (0)), gopurs_runtime.Bool(true))), gopurs_runtime.Value{})
			_ = __local_var_23_23
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1469227923_1386611502(in *Constructor_Data_Show_Show[[]int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_378698611_3790796878(in *Constructor_Data_Eq_Eq[[]int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
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
