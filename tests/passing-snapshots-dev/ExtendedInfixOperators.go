package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(Call_Main_comparing(Rebox_Main_219188042_4177771502(Rebox_Main_4177771502_219188042(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordBoolean())))), Get_Main_null(), func() gopurs_runtime.Value {
			arr := []float64{1.0, 2.0, 3.0}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Float(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := []float64{4.0, 5.0, 6.0}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Float(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).IntVal)), UnsafePtr: nil}
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := Call_Effect_Console_logShow(Rebox_Main_1209612131_1386611502(Rebox_Main_1386611502_1209612131(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Ordering_showOrdering()))), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Main_test().IntVal)), UnsafePtr: nil})
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
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

func Rebox_Main_1209612131_1386611502(in *Constructor_Data_Show_Show[uint32]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1386611502_1209612131(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[uint32])(unsafe.Pointer(in))
}

func Rebox_Main_219188042_4177771502(in *Constructor_Data_Ord_Ord[bool]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_4177771502_219188042(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[bool])(unsafe.Pointer(in))
}
