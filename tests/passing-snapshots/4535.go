package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_singleArgument gopurs_runtime.Value
var once_Main_singleArgument sync.Once

func Get_Main_singleArgument() gopurs_runtime.Value {
	once_Main_singleArgument.Do(func() {
		cache_Main_singleArgument = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_singleArgument(v_0_box)
		})
	})
	return cache_Main_singleArgument
}

var cache_Main_singleApplication gopurs_runtime.Value
var once_Main_singleApplication sync.Once

func Get_Main_singleApplication() gopurs_runtime.Value {
	once_Main_singleApplication.Do(func() {
		cache_Main_singleApplication = Get_Main_singleArgument()
	})
	return cache_Main_singleApplication
}

var cache_Main_otherNestingWorks gopurs_runtime.Value
var once_Main_otherNestingWorks sync.Once

func Get_Main_otherNestingWorks() gopurs_runtime.Value {
	once_Main_otherNestingWorks.Do(func() {
		cache_Main_otherNestingWorks = func() gopurs_runtime.Value {
			arr := []*Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, float64]]{(&Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, float64]]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), gopurs_runtime.Float(0.0)}))}}), (&Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, float64]]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(1), gopurs_runtime.Float(1.0)}))}})}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v)}
			}
			return gopurs_runtime.Array(boxed)
		}()
	})
	return cache_Main_otherNestingWorks
}

var cache_Main_operatorAsArgument gopurs_runtime.Value
var once_Main_operatorAsArgument sync.Once

func Get_Main_operatorAsArgument() gopurs_runtime.Value {
	once_Main_operatorAsArgument.Do(func() {
		cache_Main_operatorAsArgument = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_operatorAsArgument
}

var cache_Main_multiArgument gopurs_runtime.Value
var once_Main_multiArgument sync.Once

func Get_Main_multiArgument() gopurs_runtime.Value {
	once_Main_multiArgument.Do(func() {
		cache_Main_multiArgument = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_multiArgument(v_0_box, v1_1_box)
		})
	})
	return cache_Main_multiArgument
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_inSynonym gopurs_runtime.Value
var once_Main_inSynonym sync.Once

func Get_Main_inSynonym() gopurs_runtime.Value {
	once_Main_inSynonym.Do(func() {
		cache_Main_inSynonym = Get_Main_singleArgument()
	})
	return cache_Main_inSynonym
}

var cache_Main_appNestingWorks gopurs_runtime.Value
var once_Main_appNestingWorks sync.Once

func Get_Main_appNestingWorks() gopurs_runtime.Value {
	once_Main_appNestingWorks.Do(func() {
		cache_Main_appNestingWorks = Get_Main_multiArgument()
	})
	return cache_Main_appNestingWorks
}

func Call_Main_singleArgument(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return Get_Data_Unit_unit()
}

func Call_Main_multiArgument(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var v1_1 gopurs_runtime.Value = v1_1_loop
	_ = v1_1
	return Get_Data_Unit_unit()
}
