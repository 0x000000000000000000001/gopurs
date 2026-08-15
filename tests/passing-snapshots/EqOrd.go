package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Pair gopurs_runtime.Value
var once_Main_Pair sync.Once

func Get_Main_Pair() gopurs_runtime.Value {
	once_Main_Pair.Do(func() {
		cache_Main_Pair = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer((&Constructor_Main_Pair{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Pair
}

var cache_Main_eqPair gopurs_runtime.Value
var once_Main_eqPair sync.Once

func Get_Main_eqPair() gopurs_runtime.Value {
	once_Main_eqPair.Do(func() {
		cache_Main_eqPair = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqPair(dictEq_0_box, dictEq1_1_box)
		})
	})
	return cache_Main_eqPair
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("true"))
			_ = __local_var_0_0
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
				_ = _dollar___unused_1_1
				return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
			})
		}()
	})
	return cache_Main_main
}

var cache_Main_ordPair gopurs_runtime.Value
var once_Main_ordPair sync.Once

func Get_Main_ordPair() gopurs_runtime.Value {
	once_Main_ordPair.Do(func() {
		cache_Main_ordPair = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ordPair(dictOrd_0_box)
		})
	})
	return cache_Main_ordPair
}

type Constructor_Main_Pair struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func Call_Main_eqPair(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
	_ = dictEq1_1
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Main_Pair)(v_2.UnsafePtr).V0, (*Constructor_Main_Pair)(v1_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Main_Pair)(v_2.UnsafePtr).V1, (*Constructor_Main_Pair)(v1_3.UnsafePtr).V1).IntVal) != (0)))
		})
	})}))}
}

func Call_Main_ordPair(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
	__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
		__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})
		_ = __local_var_3_2
		// TAST (Let): eqPair2_3_1 -> *Constructor_Data_Eq_Eq
		eqPair2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Constructor_Main_Pair)(v_4.UnsafePtr).V0, (*Constructor_Main_Pair)(v1_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "eq"), (*Constructor_Main_Pair)(v_4.UnsafePtr).V1, (*Constructor_Main_Pair)(v1_5.UnsafePtr).V1).IntVal) != (0)))
			})
		})))
		_ = eqPair2_3_1
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqPair2_3_1)}
		}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): v2_6_3 -> gopurs_runtime.Value
				v2_6_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Main_Pair)(v_4.UnsafePtr).V0, (*Constructor_Main_Pair)(v1_5.UnsafePtr).V0)
				_ = v2_6_3
				var __t4 uint32
				{
					if uint32(v2_6_3.IntVal) == 902936544 {
						__t4 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Main_Pair)(v_4.UnsafePtr).V1, (*Constructor_Main_Pair)(v1_5.UnsafePtr).V1).IntVal)
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = uint32(v2_6_3.IntVal)
				}
			end_branch_4:
				return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
			})
		})}))}
	})
}
