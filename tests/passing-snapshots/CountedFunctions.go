package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_identityStep gopurs_runtime.Value
var once_Main_identityStep sync.Once

func Get_Main_identityStep() gopurs_runtime.Value {
	once_Main_identityStep.Do(func() {
		cache_Main_identityStep = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, value_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_identityStep(v_0_box, value_1_box.IntVal))
		})
	})
	return cache_Main_identityStep
}

var cache_Main_repeatStep__gopurs_counted_function_0 gopurs_runtime.Value
var once_Main_repeatStep__gopurs_counted_function_0 sync.Once

func Get_Main_repeatStep__gopurs_counted_function_0() gopurs_runtime.Value {
	once_Main_repeatStep__gopurs_counted_function_0.Do(func() {
		cache_Main_repeatStep__gopurs_counted_function_0 = gopurs_runtime.Func3(func(remaining_0_box gopurs_runtime.Value, callback_1_box gopurs_runtime.Value, result_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_repeatStep__gopurs_counted_function_0(remaining_0_box.IntVal, callback_1_box, result_2_box.IntVal))
		})
	})
	return cache_Main_repeatStep__gopurs_counted_function_0
}

var cache_Main_repeatStep gopurs_runtime.Value
var once_Main_repeatStep sync.Once

func Get_Main_repeatStep() gopurs_runtime.Value {
	once_Main_repeatStep.Do(func() {
		cache_Main_repeatStep = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_repeatStep(v_0_box.IntVal)
		})
	})
	return cache_Main_repeatStep
}

var cache_Main_repeatWithCounter gopurs_runtime.Value
var once_Main_repeatWithCounter sync.Once

func Get_Main_repeatWithCounter() gopurs_runtime.Value {
	once_Main_repeatWithCounter.Do(func() {
		cache_Main_repeatWithCounter = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_repeatWithCounter(v_0_box.IntVal)
		})
	})
	return cache_Main_repeatWithCounter
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(4)))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_2_2
			// TAST (Let): saved__4066693242_3_3 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
			saved__4066693242_3_3 := Call_Main_repeatStep(__local_var_2_2.IntVal)
			_ = saved__4066693242_3_3
			// TAST (Let): applied__3466805691_4_4 shape=App(Other) bindingType=(Func [Int] Int)
			applied__3466805691_4_4 := gopurs_runtime.Apply(saved__4066693242_3_3, gopurs_runtime.Func(func(value_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(((value_4.IntVal) * (int64(3))) - (int64(7)))
			}))
			_ = applied__3466805691_4_4
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{gopurs_runtime.Apply(applied__3466805691_4_4, gopurs_runtime.Int(int64(11))).IntVal, int64(611)}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
					actual   int64
					expected int64
				}{gopurs_runtime.Apply(applied__3466805691_4_4, gopurs_runtime.Int(int64(11))).IntVal, int64(611)}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
						actual   int64
						expected int64
					}{gopurs_runtime.Apply(applied__3466805691_4_4, gopurs_runtime.Int(int64(15))).IntVal, int64(935)}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
							actual   int64
							expected int64
						}{gopurs_runtime.Apply2(saved__4066693242_3_3, gopurs_runtime.Func(func(value_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Int((value_8.IntVal) + (int64(2)))
						}), gopurs_runtime.Int(int64(11))).IntVal, int64(19)}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
								actual   int64
								expected int64
							}{gopurs_runtime.Apply2(Call_Main_repeatStep((__local_var_2_2.IntVal)-(int64(4))), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Int(int64(999))
							}), gopurs_runtime.Int(int64(11))).IntVal, int64(11)}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
									actual   int64
									expected int64
								}{gopurs_runtime.Apply2(Call_Main_repeatStep(__local_var_2_2.IntVal), gopurs_runtime.Func(func(value_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Int(-(value_10.IntVal))
								}), gopurs_runtime.Int(int64(0))).IntVal, int64(0)}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___627669702("", struct {
										actual   int64
										expected int64
									}{gopurs_runtime.Apply2(Call_Main_repeatWithCounter(__local_var_2_2.IntVal), gopurs_runtime.Func(func(value_11 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Int((value_11.IntVal) + (int64(2)))
									}), gopurs_runtime.Int(int64(11))).IntVal, int64(29)}), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
									}))
								}))
							}))
						}))
					}))
				}))
			})), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_identityStep(v_0_loop gopurs_runtime.Value, value_1_loop int64) int64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var value_1 int64 = value_1_loop
	_ = value_1
	return value_1
}

func Call_Main_repeatStep__gopurs_counted_function_0(remaining_0_loop int64, callback_1_loop gopurs_runtime.Value, result_2_loop int64) int64 {
repeatStep__gopurs_counted_function_0:
	for {
		if false {
			continue repeatStep__gopurs_counted_function_0
		}
		var remaining_0 int64 = remaining_0_loop
		_ = remaining_0
		var callback_1 gopurs_runtime.Value = callback_1_loop
		_ = callback_1
		var result_2 int64 = result_2_loop
		_ = result_2
		var __t0 int64
		{
			if (remaining_0) == (int64(0)) {
				__t0 = result_2
				goto end_branch_0
			} else {

			}
		}
		{
			remaining_0_loop = (remaining_0) - (int64(1))
			callback_1_loop = callback_1
			result_2_loop = gopurs_runtime.Apply(callback_1, gopurs_runtime.Int(result_2)).IntVal
			continue repeatStep__gopurs_counted_function_0
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_repeatStep(v_0_loop int64) gopurs_runtime.Value {
repeatStep:
	for {
		if false {
			continue repeatStep
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var __t2 gopurs_runtime.Value
		{
			if (v_0) >= (int64(0)) {
				__t2 = gopurs_runtime.Func2(func(step_2 gopurs_runtime.Value, value_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_Main_repeatStep__gopurs_counted_function_0(v_0, step_2, value_3.IntVal))
				})
				goto end_branch_2
			} else {

			}
		}
		{
			var __t1 gopurs_runtime.Value
			{
				if (v_0) == (int64(0)) {
					__t1 = Get_Main_identityStep()
					goto end_branch_1
				} else {

				}
			}
			{
				// TAST (Let): previous__4066693242_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
				previous__4066693242_1_0 := Call_Main_repeatStep((v_0) - (int64(1)))
				_ = previous__4066693242_1_0
				__t1 = gopurs_runtime.Func2(func(step_2 gopurs_runtime.Value, value_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(gopurs_runtime.Apply(step_2, gopurs_runtime.Int(gopurs_runtime.Apply2(previous__4066693242_1_0, step_2, gopurs_runtime.Int(value_3.IntVal)).IntVal)).IntVal)
				})
			}
		end_branch_1:
			__t2 = __t1
		}
	end_branch_2:
		return __t2
	}
}

func Call_Main_repeatWithCounter(v_0_loop int64) gopurs_runtime.Value {
repeatWithCounter:
	for {
		if false {
			continue repeatWithCounter
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var __t1 gopurs_runtime.Value
		{
			if (v_0) == (int64(0)) {
				__t1 = Get_Main_identityStep()
				goto end_branch_1
			} else {

			}
		}
		{
			// TAST (Let): previous__4066693242_1_0 shape=App(Var) bindingType=(Func [(Func [Int] Int), Int] Int)
			previous__4066693242_1_0 := Call_Main_repeatWithCounter((v_0) - (int64(1)))
			_ = previous__4066693242_1_0
			__t1 = gopurs_runtime.Func2(func(step_2 gopurs_runtime.Value, value_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((gopurs_runtime.Apply(step_2, gopurs_runtime.Int(gopurs_runtime.Apply2(previous__4066693242_1_0, step_2, gopurs_runtime.Int(value_3.IntVal)).IntVal)).IntVal) + (v_0))
			})
		}
	end_branch_1:
		return __t1
	}
}
