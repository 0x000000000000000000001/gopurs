package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_suspendVary__gopurs_strict_thunk_0 gopurs_runtime.Value
var once_Main_suspendVary__gopurs_strict_thunk_0 sync.Once

func Get_Main_suspendVary__gopurs_strict_thunk_0() gopurs_runtime.Value {
	once_Main_suspendVary__gopurs_strict_thunk_0.Do(func() {
		cache_Main_suspendVary__gopurs_strict_thunk_0 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_suspendVary__gopurs_strict_thunk_0(v_0_box.IntVal, v1_1_box.IntVal, v2_2_box.IntVal))
		})
	})
	return cache_Main_suspendVary__gopurs_strict_thunk_0
}

var cache_Main_suspendVary gopurs_runtime.Value
var once_Main_suspendVary sync.Once

func Get_Main_suspendVary() gopurs_runtime.Value {
	once_Main_suspendVary.Do(func() {
		cache_Main_suspendVary = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_suspendVary(v_0_box.IntVal, v1_1_box.IntVal, v2_2_box)
		})
	})
	return cache_Main_suspendVary
}

var cache_Main_suspendTwice gopurs_runtime.Value
var once_Main_suspendTwice sync.Once

func Get_Main_suspendTwice() gopurs_runtime.Value {
	once_Main_suspendTwice.Do(func() {
		cache_Main_suspendTwice = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_suspendTwice(v_0_box.IntVal, v1_1_box)
		})
	})
	return cache_Main_suspendTwice
}

var cache_Main_suspendOverwrite gopurs_runtime.Value
var once_Main_suspendOverwrite sync.Once

func Get_Main_suspendOverwrite() gopurs_runtime.Value {
	once_Main_suspendOverwrite.Do(func() {
		cache_Main_suspendOverwrite = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_suspendOverwrite(v_0_box.IntVal, v1_1_box)
		})
	})
	return cache_Main_suspendOverwrite
}

var cache_Main_suspendOrder__gopurs_strict_thunk_0 gopurs_runtime.Value
var once_Main_suspendOrder__gopurs_strict_thunk_0 sync.Once

func Get_Main_suspendOrder__gopurs_strict_thunk_0() gopurs_runtime.Value {
	once_Main_suspendOrder__gopurs_strict_thunk_0.Do(func() {
		cache_Main_suspendOrder__gopurs_strict_thunk_0 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_suspendOrder__gopurs_strict_thunk_0(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Main_suspendOrder__gopurs_strict_thunk_0
}

var cache_Main_suspendOrder gopurs_runtime.Value
var once_Main_suspendOrder sync.Once

func Get_Main_suspendOrder() gopurs_runtime.Value {
	once_Main_suspendOrder.Do(func() {
		cache_Main_suspendOrder = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_suspendOrder(v_0_box.IntVal, v1_1_box)
		})
	})
	return cache_Main_suspendOrder
}

var cache_Main_suspendConditionalInt gopurs_runtime.Value
var once_Main_suspendConditionalInt sync.Once

func Get_Main_suspendConditionalInt() gopurs_runtime.Value {
	once_Main_suspendConditionalInt.Do(func() {
		cache_Main_suspendConditionalInt = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_suspendConditionalInt(v_0_box.IntVal, v1_1_box)
		})
	})
	return cache_Main_suspendConditionalInt
}

var cache_Main_suspendConditional gopurs_runtime.Value
var once_Main_suspendConditional sync.Once

func Get_Main_suspendConditional() gopurs_runtime.Value {
	once_Main_suspendConditional.Do(func() {
		cache_Main_suspendConditional = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_suspendConditional(v_0_box.IntVal, (v1_1_box.IntVal) != (0), v2_2_box)
		})
	})
	return cache_Main_suspendConditional
}

var cache_Main_suspendClash_prime___gopurs_strict_thunk_0 gopurs_runtime.Value
var once_Main_suspendClash_prime___gopurs_strict_thunk_0 sync.Once

func Get_Main_suspendClash_prime___gopurs_strict_thunk_0() gopurs_runtime.Value {
	once_Main_suspendClash_prime___gopurs_strict_thunk_0.Do(func() {
		cache_Main_suspendClash_prime___gopurs_strict_thunk_0 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_suspendClash_prime___gopurs_strict_thunk_0(x_0_box.IntVal))
		})
	})
	return cache_Main_suspendClash_prime___gopurs_strict_thunk_0
}

var cache_Main_suspendClash_prime___gopurs_strict_thunk_1 gopurs_runtime.Value
var once_Main_suspendClash_prime___gopurs_strict_thunk_1 sync.Once

func Get_Main_suspendClash_prime___gopurs_strict_thunk_1() gopurs_runtime.Value {
	once_Main_suspendClash_prime___gopurs_strict_thunk_1.Do(func() {
		cache_Main_suspendClash_prime___gopurs_strict_thunk_1 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_suspendClash_prime___gopurs_strict_thunk_1(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Main_suspendClash_prime___gopurs_strict_thunk_1
}

var cache_Main_suspendClash_prime_ gopurs_runtime.Value
var once_Main_suspendClash_prime_ sync.Once

func Get_Main_suspendClash_prime_() gopurs_runtime.Value {
	once_Main_suspendClash_prime_.Do(func() {
		cache_Main_suspendClash_prime_ = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_suspendClash_prime_(v_0_box.IntVal, v1_1_box)
		})
	})
	return cache_Main_suspendClash_prime_
}

var cache_Main_suspendAdds__gopurs_strict_thunk_0 gopurs_runtime.Value
var once_Main_suspendAdds__gopurs_strict_thunk_0 sync.Once

func Get_Main_suspendAdds__gopurs_strict_thunk_0() gopurs_runtime.Value {
	once_Main_suspendAdds__gopurs_strict_thunk_0.Do(func() {
		cache_Main_suspendAdds__gopurs_strict_thunk_0 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_suspendAdds__gopurs_strict_thunk_0(v_0_box.IntVal, v1_1_box.IntVal, v2_2_box.IntVal))
		})
	})
	return cache_Main_suspendAdds__gopurs_strict_thunk_0
}

var cache_Main_suspendAdds gopurs_runtime.Value
var once_Main_suspendAdds sync.Once

func Get_Main_suspendAdds() gopurs_runtime.Value {
	once_Main_suspendAdds.Do(func() {
		cache_Main_suspendAdds = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_suspendAdds(v_0_box.IntVal, v1_1_box.IntVal, v2_2_box)
		})
	})
	return cache_Main_suspendAdds
}

var cache_Main_scheduleEffects gopurs_runtime.Value
var once_Main_scheduleEffects sync.Once

func Get_Main_scheduleEffects() gopurs_runtime.Value {
	once_Main_scheduleEffects.Do(func() {
		cache_Main_scheduleEffects = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_scheduleEffects(v_0_box.IntVal, v1_1_box)
		})
	})
	return cache_Main_scheduleEffects
}

var cache_Main_runVary gopurs_runtime.Value
var once_Main_runVary sync.Once

func Get_Main_runVary() gopurs_runtime.Value {
	once_Main_runVary.Do(func() {
		cache_Main_runVary = gopurs_runtime.Func3(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value, step_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runVary(depth_0_box.IntVal, seed_1_box.IntVal, step_2_box.IntVal))
		})
	})
	return cache_Main_runVary
}

var cache_Main_runTwice gopurs_runtime.Value
var once_Main_runTwice sync.Once

func Get_Main_runTwice() gopurs_runtime.Value {
	once_Main_runTwice.Do(func() {
		cache_Main_runTwice = gopurs_runtime.Func2(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runTwice(depth_0_box.IntVal, seed_1_box.IntVal))
		})
	})
	return cache_Main_runTwice
}

var cache_Main_runRecursiveBody gopurs_runtime.Value
var once_Main_runRecursiveBody sync.Once

func Get_Main_runRecursiveBody() gopurs_runtime.Value {
	once_Main_runRecursiveBody.Do(func() {
		cache_Main_runRecursiveBody = gopurs_runtime.Func2(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runRecursiveBody(depth_0_box.IntVal, seed_1_box.IntVal))
		})
	})
	return cache_Main_runRecursiveBody
}

var cache_Main_runRecursiveBinding gopurs_runtime.Value
var once_Main_runRecursiveBinding sync.Once

func Get_Main_runRecursiveBinding() gopurs_runtime.Value {
	once_Main_runRecursiveBinding.Do(func() {
		cache_Main_runRecursiveBinding = gopurs_runtime.Func2(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runRecursiveBinding(depth_0_box.IntVal, seed_1_box.IntVal))
		})
	})
	return cache_Main_runRecursiveBinding
}

var cache_Main_runOverwriteTotal gopurs_runtime.Value
var once_Main_runOverwriteTotal sync.Once

func Get_Main_runOverwriteTotal() gopurs_runtime.Value {
	once_Main_runOverwriteTotal.Do(func() {
		cache_Main_runOverwriteTotal = gopurs_runtime.Func2(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runOverwriteTotal(depth_0_box.IntVal, seed_1_box.IntVal))
		})
	})
	return cache_Main_runOverwriteTotal
}

var cache_Main_runOverwrite gopurs_runtime.Value
var once_Main_runOverwrite sync.Once

func Get_Main_runOverwrite() gopurs_runtime.Value {
	once_Main_runOverwrite.Do(func() {
		cache_Main_runOverwrite = gopurs_runtime.Func(func(depth_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runOverwrite(depth_0_box.IntVal))
		})
	})
	return cache_Main_runOverwrite
}

var cache_Main_runOrder gopurs_runtime.Value
var once_Main_runOrder sync.Once

func Get_Main_runOrder() gopurs_runtime.Value {
	once_Main_runOrder.Do(func() {
		cache_Main_runOrder = gopurs_runtime.Func2(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runOrder(depth_0_box.IntVal, seed_1_box.IntVal))
		})
	})
	return cache_Main_runOrder
}

var cache_Main_runConditionalInt gopurs_runtime.Value
var once_Main_runConditionalInt sync.Once

func Get_Main_runConditionalInt() gopurs_runtime.Value {
	once_Main_runConditionalInt.Do(func() {
		cache_Main_runConditionalInt = gopurs_runtime.Func2(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runConditionalInt(depth_0_box.IntVal, seed_1_box.IntVal))
		})
	})
	return cache_Main_runConditionalInt
}

var cache_Main_runConditional gopurs_runtime.Value
var once_Main_runConditional sync.Once

func Get_Main_runConditional() gopurs_runtime.Value {
	once_Main_runConditional.Do(func() {
		cache_Main_runConditional = gopurs_runtime.Func2(func(depth_0_box gopurs_runtime.Value, demand_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runConditional(depth_0_box.IntVal, (demand_1_box.IntVal) != (0)))
		})
	})
	return cache_Main_runConditional
}

var cache_Main_runClash gopurs_runtime.Value
var once_Main_runClash sync.Once

func Get_Main_runClash() gopurs_runtime.Value {
	once_Main_runClash.Do(func() {
		cache_Main_runClash = gopurs_runtime.Func2(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runClash(depth_0_box.IntVal, seed_1_box.IntVal))
		})
	})
	return cache_Main_runClash
}

var cache_Main_runAdds gopurs_runtime.Value
var once_Main_runAdds sync.Once

func Get_Main_runAdds() gopurs_runtime.Value {
	once_Main_runAdds.Do(func() {
		cache_Main_runAdds = gopurs_runtime.Func3(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value, step_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_runAdds(depth_0_box.IntVal, seed_1_box.IntVal, step_2_box.IntVal))
		})
	})
	return cache_Main_runAdds
}

var cache_Main_keepSuspended gopurs_runtime.Value
var once_Main_keepSuspended sync.Once

func Get_Main_keepSuspended() gopurs_runtime.Value {
	once_Main_keepSuspended.Do(func() {
		cache_Main_keepSuspended = gopurs_runtime.Func2(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_keepSuspended(depth_0_box.IntVal, seed_1_box.IntVal)
		})
	})
	return cache_Main_keepSuspended
}

var cache_Main_checkAdds gopurs_runtime.Value
var once_Main_checkAdds sync.Once

func Get_Main_checkAdds() gopurs_runtime.Value {
	once_Main_checkAdds.Do(func() {
		cache_Main_checkAdds = gopurs_runtime.Func3(func(depth_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value, step_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkAdds(depth_0_box.IntVal, seed_1_box.IntVal, step_2_box.IntVal)
		})
	})
	return cache_Main_checkAdds
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(1000)))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(7))), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(-3))), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_2_2), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_3_3), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(Call_Main_checkAdds(int64(0), __local_var_5_5.IntVal, __local_var_6_6.IntVal), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Main_checkAdds(int64(1), __local_var_5_5.IntVal, __local_var_6_6.IntVal), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Main_checkAdds(int64(2), __local_var_5_5.IntVal, __local_var_6_6.IntVal), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Call_Main_checkAdds(int64(7), __local_var_5_5.IntVal, __local_var_6_6.IntVal), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(Call_Main_checkAdds(__local_var_4_4.IntVal, __local_var_5_5.IntVal, __local_var_6_6.IntVal), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Main_checkAdds(int64(7), __local_var_6_6.IntVal, __local_var_5_5.IntVal), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_suspendOrder__gopurs_strict_thunk_0(int64(0), __local_var_5_5.IntVal)), gopurs_runtime.Int(__local_var_5_5.IntVal))), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_suspendOrder__gopurs_strict_thunk_0(int64(3), __local_var_5_5.IntVal)), gopurs_runtime.Int(int64(73)))), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_suspendVary__gopurs_strict_thunk_0(int64(3), int64(2), __local_var_5_5.IntVal)), gopurs_runtime.Int(int64(30)))), gopurs_runtime.Value{})
			_ = __local_var_15_15
			__local_var_16_16 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Main_suspendClash_prime___gopurs_strict_thunk_1(), gopurs_runtime.Int(int64(3)), gopurs_runtime.Int(__local_var_5_5.IntVal)).IntVal), gopurs_runtime.Int(int64(10)))), gopurs_runtime.Value{})
			_ = __local_var_16_16
			__local_var_17_17 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int((__local_var_5_5.IntVal)+(int64(100))), gopurs_runtime.Int(int64(107)))), gopurs_runtime.Value{})
			_ = __local_var_17_17
			// TAST (Let): escaped__3314126478_18_18 shape=App(Var) bindingType=(Func [Unit] Int)
			escaped__3314126478_18_18 := Call_Main_suspendAdds(int64(7), int64(2), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(__local_var_5_5.IntVal)
			}))
			_ = escaped__3314126478_18_18
			__local_var_19_19 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int((gopurs_runtime.Apply(escaped__3314126478_18_18, Get_Data_Unit_unit()).IntVal)+(gopurs_runtime.Apply(escaped__3314126478_18_18, Get_Data_Unit_unit()).IntVal)), gopurs_runtime.Int(int64(42)))), gopurs_runtime.Value{})
			_ = __local_var_19_19
			__local_var_20_20 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Call_Main_suspendConditional(int64(3), false, gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("Unused seed was forced")).IntVal)
			})), Get_Data_Unit_unit()).IntVal), gopurs_runtime.Int(int64(1)))), gopurs_runtime.Value{})
			_ = __local_var_20_20
			__local_var_21_21 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_runOverwrite(int64(3))), gopurs_runtime.Int(int64(1)))), gopurs_runtime.Value{})
			_ = __local_var_21_21
			__local_var_22_22 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Call_Main_suspendConditionalInt(int64(0), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(__local_var_5_5.IntVal)
			})), Get_Data_Unit_unit()).IntVal), gopurs_runtime.Int(__local_var_5_5.IntVal))), gopurs_runtime.Value{})
			_ = __local_var_22_22
			__local_var_23_23 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Call_Main_suspendConditionalInt(int64(1), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(__local_var_5_5.IntVal)
			})), Get_Data_Unit_unit()).IntVal), gopurs_runtime.Int(int64(8)))), gopurs_runtime.Value{})
			_ = __local_var_23_23
			__local_var_24_24 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Call_Main_suspendConditionalInt(int64(3), gopurs_runtime.Func(func(v_24 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(__local_var_5_5.IntVal)
			})), Get_Data_Unit_unit()).IntVal), gopurs_runtime.Int(int64(3)))), gopurs_runtime.Value{})
			_ = __local_var_24_24
			__local_var_25_25 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(__local_var_5_5.IntVal), gopurs_runtime.Int(__local_var_5_5.IntVal))), gopurs_runtime.Value{})
			_ = __local_var_25_25
			__local_var_26_26 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Call_Main_suspendOverwrite(int64(1), gopurs_runtime.Func(func(v2_26 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(int64(2))
			})), Get_Data_Unit_unit()).IntVal), gopurs_runtime.Int(int64(1)))), gopurs_runtime.Value{})
			_ = __local_var_26_26
			__local_var_27_27 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Call_Main_suspendTwice(int64(0), gopurs_runtime.Func(func(v_27 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(__local_var_5_5.IntVal)
			})), Get_Data_Unit_unit()).IntVal), gopurs_runtime.Int(__local_var_5_5.IntVal))), gopurs_runtime.Value{})
			_ = __local_var_27_27
			__local_var_28_28 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(gopurs_runtime.Apply(Call_Main_suspendTwice(int64(3), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(__local_var_5_5.IntVal)
			})), Get_Data_Unit_unit()).IntVal), gopurs_runtime.Int(int64(56)))), gopurs_runtime.Value{})
			_ = __local_var_28_28
			__local_var_29_29 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_runRecursiveBinding(int64(0), __local_var_5_5.IntVal)), gopurs_runtime.Int(__local_var_5_5.IntVal))), gopurs_runtime.Value{})
			_ = __local_var_29_29
			__local_var_30_30 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_runRecursiveBinding(int64(3), __local_var_5_5.IntVal)), gopurs_runtime.Int(int64(16)))), gopurs_runtime.Value{})
			_ = __local_var_30_30
			__local_var_31_31 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_runRecursiveBinding(__local_var_4_4.IntVal, __local_var_5_5.IntVal)), gopurs_runtime.Int((__local_var_5_5.IntVal)+(int64(3000))))), gopurs_runtime.Value{})
			_ = __local_var_31_31
			__local_var_32_32 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_runRecursiveBody(int64(0), __local_var_5_5.IntVal)), gopurs_runtime.Int((int64(2))*(__local_var_5_5.IntVal)))), gopurs_runtime.Value{})
			_ = __local_var_32_32
			__local_var_33_33 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_runRecursiveBody(int64(3), __local_var_5_5.IntVal)), gopurs_runtime.Int(int64(23)))), gopurs_runtime.Value{})
			_ = __local_var_33_33
			__local_var_34_34 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_runRecursiveBody(__local_var_4_4.IntVal, __local_var_5_5.IntVal)), gopurs_runtime.Int(((int64(2))*(__local_var_5_5.IntVal))+(int64(3000))))), gopurs_runtime.Value{})
			_ = __local_var_34_34
			__local_var_35_35 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(0))), gopurs_runtime.Value{})
			_ = __local_var_35_35
			// TAST (Let): pending__1854579232_36_36 shape=App(Var) bindingType=(Func [Unit] (ADT ["Effect","Effect"] [Int]))
			pending__1854579232_36_36 := Call_Main_scheduleEffects(int64(3), gopurs_runtime.Func(func(v_36 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_37_37 shape=App(Var) bindingType=Any
					__local_var_37_37 := gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(count_37 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Int((count_37.IntVal) + (int64(1)))
					}), __local_var_35_35)
					_ = __local_var_37_37
					_dollar___unused_38_38 := gopurs_runtime.Apply(__local_var_37_37, gopurs_runtime.Value{})
					_ = _dollar___unused_38_38
					return __local_var_5_5
				})
			}))
			_ = pending__1854579232_36_36
			__local_var_37_39 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_35_35), gopurs_runtime.Value{})
			_ = __local_var_37_39
			__local_var_38_40 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(__local_var_37_39.IntVal), gopurs_runtime.Int(int64(0)))), gopurs_runtime.Value{})
			_ = __local_var_38_40
			__local_var_39_41 := gopurs_runtime.Apply(gopurs_runtime.Apply(pending__1854579232_36_36, Get_Data_Unit_unit()), gopurs_runtime.Value{})
			_ = __local_var_39_41
			__local_var_40_42 := gopurs_runtime.Apply(gopurs_runtime.Apply(pending__1854579232_36_36, Get_Data_Unit_unit()), gopurs_runtime.Value{})
			_ = __local_var_40_42
			__local_var_41_43 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_35_35), gopurs_runtime.Value{})
			_ = __local_var_41_43
			__local_var_42_44 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(__local_var_39_41.IntVal), gopurs_runtime.Int(int64(10)))), gopurs_runtime.Value{})
			_ = __local_var_42_44
			__local_var_43_45 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(__local_var_40_42.IntVal), gopurs_runtime.Int(int64(10)))), gopurs_runtime.Value{})
			_ = __local_var_43_45
			__local_var_44_46 := gopurs_runtime.Apply(gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(__local_var_41_43.IntVal), gopurs_runtime.Int(int64(2)))), gopurs_runtime.Value{})
			_ = __local_var_44_46
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_suspendVary__gopurs_strict_thunk_0(v_0_loop int64, v1_1_loop int64, v2_2_loop int64) int64 {
suspendVary__gopurs_strict_thunk_0:
	for {
		if false {
			continue suspendVary__gopurs_strict_thunk_0
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var v2_2 int64 = v2_2_loop
		_ = v2_2
		var __t0 int64
		{
			if (v_0) == (int64(0)) {
				__t0 = v2_2
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = (v1_1) + (v_0)
			v2_2_loop = (v2_2) + ((v1_1) * (v_0))
			continue suspendVary__gopurs_strict_thunk_0
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_suspendVary(v_0_loop int64, v1_1_loop int64, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
suspendVary:
	for {
		if false {
			continue suspendVary
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var v2_2 gopurs_runtime.Value = v2_2_loop
		_ = v2_2
		var __t0 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t0 = v2_2
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = (v1_1) + (v_0)
			v2_2_loop = gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((gopurs_runtime.Apply(v2_2, Get_Data_Unit_unit()).IntVal) + ((v1_1) * (v_0)))
			})
			continue suspendVary
			__t0 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_suspendTwice(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
suspendTwice:
	for {
		if false {
			continue suspendTwice
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var __t0 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit()).IntVal) + (gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit()).IntVal))
			})
			continue suspendTwice
			__t0 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_suspendOverwrite(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
suspendOverwrite:
	for {
		if false {
			continue suspendOverwrite
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var __t0 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(v_0)
			})
			continue suspendOverwrite
			__t0 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_suspendOrder__gopurs_strict_thunk_0(v_0_loop int64, v1_1_loop int64) int64 {
suspendOrder__gopurs_strict_thunk_0:
	for {
		if false {
			continue suspendOrder__gopurs_strict_thunk_0
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var __t0 int64
		{
			if (v_0) == (int64(0)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = ((int64(2)) * (v1_1)) + (v_0)
			continue suspendOrder__gopurs_strict_thunk_0
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_suspendOrder(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
suspendOrder:
	for {
		if false {
			continue suspendOrder
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var __t0 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(((int64(2)) * (gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit()).IntVal)) + (v_0))
			})
			continue suspendOrder
			__t0 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_suspendConditionalInt(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
suspendConditionalInt:
	for {
		if false {
			continue suspendConditionalInt
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var __t1 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t1 = v1_1
				goto end_branch_1
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 int64
				{
					if (v_0) == (int64(2)) {
						__t0 = v_0
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = (gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit()).IntVal) + (int64(1))
				}
			end_branch_0:
				return gopurs_runtime.Int(__t0)
			})
			continue suspendConditionalInt
			__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_1:
		return __t1
	}
}

func Call_Main_suspendConditional(v_0_loop int64, v1_1_loop bool, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
suspendConditional:
	for {
		if false {
			continue suspendConditional
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 bool = v1_1_loop
		_ = v1_1
		var v2_2 gopurs_runtime.Value = v2_2_loop
		_ = v2_2
		var __t1 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t1 = v2_2
				goto end_branch_1
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = v1_1
			v2_2_loop = gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t0 int64
				{
					if v1_1 {
						__t0 = (gopurs_runtime.Apply(v2_2, Get_Data_Unit_unit()).IntVal) + (int64(1))
						goto end_branch_0
					} else {

					}
				}
				{
					__t0 = v_0
				}
			end_branch_0:
				return gopurs_runtime.Int(__t0)
			})
			continue suspendConditional
			__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_1:
		return __t1
	}
}

func Call_Main_suspendClash_prime___gopurs_strict_thunk_0(x_0_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	return (x_0) + (int64(100))
}

func Call_Main_suspendClash_prime___gopurs_strict_thunk_1(v_0_loop int64, v1_1_loop int64) int64 {
suspendClash_prime___gopurs_strict_thunk_1:
	for {
		if false {
			continue suspendClash_prime___gopurs_strict_thunk_1
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var __t0 int64
		{
			if (v_0) == (int64(0)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = (v1_1) + (int64(1))
			continue suspendClash_prime___gopurs_strict_thunk_1
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_suspendClash_prime_(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
suspendClash_prime_:
	for {
		if false {
			continue suspendClash_prime_
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var __t0 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit()).IntVal) + (int64(1)))
			})
			continue suspendClash_prime_
			__t0 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_suspendAdds__gopurs_strict_thunk_0(v_0_loop int64, v1_1_loop int64, v2_2_loop int64) int64 {
suspendAdds__gopurs_strict_thunk_0:
	for {
		if false {
			continue suspendAdds__gopurs_strict_thunk_0
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var v2_2 int64 = v2_2_loop
		_ = v2_2
		var __t0 int64
		{
			if (v_0) == (int64(0)) {
				__t0 = v2_2
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = v1_1
			v2_2_loop = (v2_2) + (v1_1)
			continue suspendAdds__gopurs_strict_thunk_0
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_suspendAdds(v_0_loop int64, v1_1_loop int64, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
suspendAdds:
	for {
		if false {
			continue suspendAdds
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var v2_2 gopurs_runtime.Value = v2_2_loop
		_ = v2_2
		var __t0 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t0 = v2_2
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = v1_1
			v2_2_loop = gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((gopurs_runtime.Apply(v2_2, Get_Data_Unit_unit()).IntVal) + (v1_1))
			})
			continue suspendAdds
			__t0 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_scheduleEffects(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
scheduleEffects:
	for {
		if false {
			continue scheduleEffects
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var __t2 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t2 = v1_1
				goto end_branch_2
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(ADT ["Effect","Effect"] [Int])
					__local_var_3_0 := gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit())
					_ = __local_var_3_0
					__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
					_ = __local_var_4_1
					return gopurs_runtime.Int((__local_var_4_1.IntVal) + (int64(1)))
				})
			})
			continue scheduleEffects
			__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
		}
	end_branch_2:
		return __t2
	}
}

func Call_Main_runVary(depth_0_loop int64, seed_1_loop int64, step_2_loop int64) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	var step_2 int64 = step_2_loop
	_ = step_2
	return Call_Main_suspendVary__gopurs_strict_thunk_0(depth_0, step_2, seed_1)
}

func Call_Main_runTwice(depth_0_loop int64, seed_1_loop int64) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	return gopurs_runtime.Apply(Call_Main_suspendTwice(depth_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int(seed_1)
	})), Get_Data_Unit_unit()).IntVal
}

func Call_Main_runRecursiveBody(depth_0_loop int64, seed_1_loop int64) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	var visit__3466805691_2_0_0 gopurs_runtime.Value
	_ = visit__3466805691_2_0_0
	// FALLBACK TCO: isLoop=false len=1
	visit__3466805691_2_0_0 = gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 int64
		{
			if (n_3.IntVal) == (int64(0)) {
				__t1 = seed_1
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = (gopurs_runtime.Apply(visit__3466805691_2_0_0, gopurs_runtime.Int((n_3.IntVal)-(int64(1)))).IntVal) + (int64(1))
		}
	end_branch_1:
		return gopurs_runtime.Int(__t1)
	})
	var visit_3_2_1 gopurs_runtime.Value
	_ = visit_3_2_1
	// FALLBACK TCO: isLoop=false len=1
	visit_3_2_1 = gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t3 int64
		{
			if (n_4.IntVal) == (int64(0)) {
				__t3 = seed_1
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = (gopurs_runtime.Apply(visit__3466805691_2_0_0, gopurs_runtime.Int((n_4.IntVal)-(int64(1)))).IntVal) + (int64(1))
		}
	end_branch_3:
		return gopurs_runtime.Int(__t3)
	})
	return (gopurs_runtime.Apply(visit__3466805691_2_0_0, gopurs_runtime.Int(depth_0)).IntVal) + (gopurs_runtime.Apply(Call_Main_suspendAdds(depth_0, int64(2), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int(seed_1)
	})), Get_Data_Unit_unit()).IntVal)
}

func Call_Main_runRecursiveBinding(depth_0_loop int64, seed_1_loop int64) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	var visit__3466805691_2_0_2 gopurs_runtime.Value
	_ = visit__3466805691_2_0_2
	// FALLBACK TCO: isLoop=false len=1
	visit__3466805691_2_0_2 = gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 int64
		{
			if (n_3.IntVal) == (int64(0)) {
				__t1 = gopurs_runtime.Apply(Call_Main_suspendAdds(depth_0, int64(2), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(seed_1)
				})), Get_Data_Unit_unit()).IntVal
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = (gopurs_runtime.Apply(visit__3466805691_2_0_2, gopurs_runtime.Int((n_3.IntVal)-(int64(1)))).IntVal) + (int64(1))
		}
	end_branch_1:
		return gopurs_runtime.Int(__t1)
	})
	var visit_3_2_3 gopurs_runtime.Value
	_ = visit_3_2_3
	// FALLBACK TCO: isLoop=false len=1
	visit_3_2_3 = gopurs_runtime.Func(func(n_4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t3 int64
		{
			if (n_4.IntVal) == (int64(0)) {
				__t3 = gopurs_runtime.Apply(Call_Main_suspendAdds(depth_0, int64(2), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(seed_1)
				})), Get_Data_Unit_unit()).IntVal
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = (gopurs_runtime.Apply(visit__3466805691_2_0_2, gopurs_runtime.Int((n_4.IntVal)-(int64(1)))).IntVal) + (int64(1))
		}
	end_branch_3:
		return gopurs_runtime.Int(__t3)
	})
	return gopurs_runtime.Apply(visit__3466805691_2_0_2, gopurs_runtime.Int(depth_0)).IntVal
}

func Call_Main_runOverwriteTotal(depth_0_loop int64, seed_1_loop int64) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	var __t0 int64
	{
		if (depth_0) == (int64(0)) {
			__t0 = seed_1
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.Apply(Call_Main_suspendOverwrite((depth_0)-(int64(1)), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(depth_0)
		})), Get_Data_Unit_unit()).IntVal
	}
end_branch_0:
	return __t0
}

func Call_Main_runOverwrite(depth_0_loop int64) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var __t0 int64
	{
		if (depth_0) == (int64(0)) {
			__t0 = gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("Discarded seed was forced")).IntVal
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.Apply(Call_Main_suspendOverwrite((depth_0)-(int64(1)), gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(depth_0)
		})), Get_Data_Unit_unit()).IntVal
	}
end_branch_0:
	return __t0
}

func Call_Main_runOrder(depth_0_loop int64, seed_1_loop int64) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	return Call_Main_suspendOrder__gopurs_strict_thunk_0(depth_0, seed_1)
}

func Call_Main_runConditionalInt(depth_0_loop int64, seed_1_loop int64) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	return gopurs_runtime.Apply(Call_Main_suspendConditionalInt(depth_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int(seed_1)
	})), Get_Data_Unit_unit()).IntVal
}

func Call_Main_runConditional(depth_0_loop int64, demand_1_loop bool) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var demand_1 bool = demand_1_loop
	_ = demand_1
	return gopurs_runtime.Apply(Call_Main_suspendConditional(depth_0, demand_1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int(gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("Unused seed was forced")).IntVal)
	})), Get_Data_Unit_unit()).IntVal
}

func Call_Main_runClash(depth_0_loop int64, seed_1_loop int64) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	return gopurs_runtime.Apply2(Get_Main_suspendClash_prime___gopurs_strict_thunk_1(), gopurs_runtime.Int(depth_0), gopurs_runtime.Int(seed_1)).IntVal
}

func Call_Main_runAdds(depth_0_loop int64, seed_1_loop int64, step_2_loop int64) int64 {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	var step_2 int64 = step_2_loop
	_ = step_2
	return Call_Main_suspendAdds__gopurs_strict_thunk_0(depth_0, step_2, seed_1)
}

func Call_Main_keepSuspended(depth_0_loop int64, seed_1_loop int64) gopurs_runtime.Value {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	return Call_Main_suspendAdds(depth_0, int64(2), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int(seed_1)
	}))
}

func Call_Main_checkAdds(depth_0_loop int64, seed_1_loop int64, step_2_loop int64) gopurs_runtime.Value {
	var depth_0 int64 = depth_0_loop
	_ = depth_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	var step_2 int64 = step_2_loop
	_ = step_2
	return gopurs_runtime.Apply4(Get_Test_Assert_assertEqual_prime_(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[int64]](Get_Data_Show_showInt())))}, gopurs_runtime.Str(""), gopurs_runtime.RecordDict2("actual", "expected", gopurs_runtime.Int(Call_Main_suspendAdds__gopurs_strict_thunk_0(depth_0, step_2, seed_1)), gopurs_runtime.Int((seed_1)+((depth_0)*(step_2)))))
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
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
