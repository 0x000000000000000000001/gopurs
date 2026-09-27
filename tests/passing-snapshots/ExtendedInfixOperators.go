package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_null gopurs_runtime.Value
var once_Main_null sync.Once

func Get_Main_null() gopurs_runtime.Value {
	once_Main_null.Do(func() {
		cache_Main_null = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Main_null((*(*[]gopurs_runtime.Value)((v_0_box).UnsafePtr))))
		})
	})
	return cache_Main_null
}

var cache_Main_comparing gopurs_runtime.Value
var once_Main_comparing sync.Once

func Get_Main_comparing() gopurs_runtime.Value {
	once_Main_comparing.Do(func() {
		cache_Main_comparing = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_comparing(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
		})
	})
	return cache_Main_comparing
}

var cache_Main_comparing__1991096015 gopurs_runtime.Value
var once_Main_comparing__1991096015 sync.Once

func Get_Main_comparing__1991096015() gopurs_runtime.Value {
	once_Main_comparing__1991096015.Do(func() {
		cache_Main_comparing__1991096015 = gopurs_runtime.Func3(func(__eta_norm_2_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_comparing__1991096015(__eta_norm_2_unused_0_box, func() []float64 {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_1_1_box.UnsafePtr)
				unboxed := make([]float64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.FloatVal()
				}
				return unboxed
			}(), func() []float64 {
				arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_2_box.UnsafePtr)
				unboxed := make([]float64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.FloatVal()
				}
				return unboxed
			}())), UnsafePtr: nil}
		})
	})
	return cache_Main_comparing__1991096015
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Data_Ord_OrdBooleanImpl_nativeWorker(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, (gopurs_runtime.Bool(false).IntVal) != (0), (gopurs_runtime.Bool(false).IntVal) != (0)).IntVal)), UnsafePtr: nil}
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 string
			{
				var __t_tag_1 uint32 = uint32(Get_Main_test().IntVal)
				_ = __t_tag_1
				if uint32(__t_tag_1) == 1527465420 {
					__t4 = "LT"
					goto end_branch_4
				} else {

				}
			}
			{
				var __t_tag_2 uint32 = uint32(Get_Main_test().IntVal)
				_ = __t_tag_2
				if uint32(__t_tag_2) == 380165415 {
					__t4 = "GT"
					goto end_branch_4
				} else {

				}
			}
			{
				var __t_tag_3 uint32 = uint32(Get_Main_test().IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 902936544 {
					__t4 = "EQ"
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = func() string { panic("Failed pattern match") }()
			}
		end_branch_4:
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__t4))
			_ = __local_var_0_0
			_dollar___unused_1_5 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_5
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_null(v_0_loop []gopurs_runtime.Value) bool {
	var v_0 []gopurs_runtime.Value = v_0_loop
	_ = v_0
	return (gopurs_runtime.Int(int64(len(v_0))).IntVal) == (int64(0))
}

func Call_Main_comparing(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
	_ = dictOrd_0
	// TAST (Let): compare_1_0 shape=App(Var) bindingType=(Func [(TypeVar b$scope3), (TypeVar b$scope3)] (ADT ["Data","Ordering","Ordering"] []))
	compare_1_0 := Call_Data_Ord_compare(dictOrd_0)
	_ = compare_1_0
	return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(compare_1_0, gopurs_runtime.Apply(f_2, x_3), gopurs_runtime.Apply(f_2, y_4))
	})
}

func Call_Main_comparing__1991096015(__eta_norm_2_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop []float64, __eta_norm_0_2_loop []float64) uint32 {
comparing__1991096015:
	for {
		if false {
			continue comparing__1991096015
		}
		var __eta_norm_2_unused_0 gopurs_runtime.Value = __eta_norm_2_unused_0_loop
		_ = __eta_norm_2_unused_0
		var __eta_norm_1_1 []float64 = __eta_norm_1_1_loop
		_ = __eta_norm_1_1
		var __eta_norm_0_2 []float64 = __eta_norm_0_2_loop
		_ = __eta_norm_0_2
		return uint32(Data_Ord_OrdBooleanImpl_nativeWorker(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, (gopurs_runtime.Bool((gopurs_runtime.Int(int64(len((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := __eta_norm_1_1
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Float(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).UnsafePtr))))).IntVal) == (int64(0))).IntVal) != (0), (gopurs_runtime.Bool((gopurs_runtime.Int(int64(len((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := __eta_norm_0_2
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Float(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).UnsafePtr))))).IntVal) == (int64(0))).IntVal) != (0)).IntVal)
	}
}
