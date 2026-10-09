package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_shiftState gopurs_runtime.Value
var once_Main_shiftState sync.Once

func Get_Main_shiftState() gopurs_runtime.Value {
	once_Main_shiftState.Do(func() {
		cache_Main_shiftState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_shiftState(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Main_shiftState
}

var cache_Main_localLoop gopurs_runtime.Value
var once_Main_localLoop sync.Once

func Get_Main_localLoop() gopurs_runtime.Value {
	once_Main_localLoop.Do(func() {
		cache_Main_localLoop = gopurs_runtime.Func2(func(count_0_box gopurs_runtime.Value, initial_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_localLoop(count_0_box.IntVal, initial_1_box.IntVal))
		})
	})
	return cache_Main_localLoop
}

var cache_Main_literalStep gopurs_runtime.Value
var once_Main_literalStep sync.Once

func Get_Main_literalStep() gopurs_runtime.Value {
	once_Main_literalStep.Do(func() {
		cache_Main_literalStep = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_literalStep(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Main_literalStep
}

var cache_Main_literalState gopurs_runtime.Value
var once_Main_literalState sync.Once

func Get_Main_literalState() gopurs_runtime.Value {
	once_Main_literalState.Do(func() {
		cache_Main_literalState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_literalState(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Main_literalState
}

var cache_Main_literalCompare gopurs_runtime.Value
var once_Main_literalCompare sync.Once

func Get_Main_literalCompare() gopurs_runtime.Value {
	once_Main_literalCompare.Do(func() {
		cache_Main_literalCompare = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_literalCompare(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Main_literalCompare
}

var cache_Main_divideState gopurs_runtime.Value
var once_Main_divideState sync.Once

func Get_Main_divideState() gopurs_runtime.Value {
	once_Main_divideState.Do(func() {
		cache_Main_divideState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_divideState(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Main_divideState
}

var cache_Main_choose gopurs_runtime.Value
var once_Main_choose sync.Once

func Get_Main_choose() gopurs_runtime.Value {
	once_Main_choose.Do(func() {
		cache_Main_choose = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_choose(v_0_box.IntVal, v1_1_box.IntVal, v2_2_box.IntVal))
		})
	})
	return cache_Main_choose
}

var cache_Main_accumulate gopurs_runtime.Value
var once_Main_accumulate sync.Once

func Get_Main_accumulate() gopurs_runtime.Value {
	once_Main_accumulate.Do(func() {
		cache_Main_accumulate = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_accumulate(v_0_box.IntVal, v1_1_box.IntVal, v2_2_box.IntVal))
		})
	})
	return cache_Main_accumulate
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
				orig := struct {
					count    int64
					high     int64
					low      int64
					negative int64
				}{int64(3), Get_Data_Bounded_topInt().IntVal, Get_Data_Bounded_bottomInt().IntVal, int64(-1)}
				_ = orig
				return gopurs_runtime.RecordDict4("count", "high", "low", "negative", gopurs_runtime.Int(orig.count), gopurs_runtime.Int(orig.high), gopurs_runtime.Int(orig.low), gopurs_runtime.Int(orig.negative))
			}())
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_2_2
			// TAST (Let): unsigned_3_3 shape=Other bindingType=Int
			unsigned_3_3 := gopurs_runtime.Zshr(gopurs_runtime.RecordGet(__local_var_2_2, "negative"), gopurs_runtime.Int(int64(0))).IntVal
			_ = unsigned_3_3
			// TAST (Let): quotient_4_4 shape=Other bindingType=Int
			quotient_4_4 := gopurs_runtime.IntDiv(gopurs_runtime.RecordGet(__local_var_2_2, "low").IntVal, gopurs_runtime.RecordGet(__local_var_2_2, "negative").IntVal)
			_ = quotient_4_4
			__local_var_5_5 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_accumulate(gopurs_runtime.RecordGet(__local_var_2_2, "count").IntVal, gopurs_runtime.RecordGet(__local_var_2_2, "high").IntVal, int64(1)), gopurs_runtime.IntAdd(gopurs_runtime.RecordGet(__local_var_2_2, "low").IntVal, int64(2))}), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_accumulate(gopurs_runtime.RecordGet(__local_var_2_2, "count").IntVal, gopurs_runtime.RecordGet(__local_var_2_2, "low").IntVal, gopurs_runtime.RecordGet(__local_var_2_2, "negative").IntVal), gopurs_runtime.IntSub(gopurs_runtime.RecordGet(__local_var_2_2, "high").IntVal, int64(2))}), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_accumulate(int64(0), unsigned_3_3, int64(1)), unsigned_3_3}), gopurs_runtime.Value{})
			_ = __local_var_7_7
			__local_var_8_8 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_accumulate(int64(0), quotient_4_4, int64(1)), quotient_4_4}), gopurs_runtime.Value{})
			_ = __local_var_8_8
			__local_var_9_9 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_accumulate(gopurs_runtime.RecordGet(__local_var_2_2, "count").IntVal, unsigned_3_3, int64(1)), int64(2)}), gopurs_runtime.Value{})
			_ = __local_var_9_9
			__local_var_10_10 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_accumulate(gopurs_runtime.RecordGet(__local_var_2_2, "count").IntVal, quotient_4_4, int64(1)), gopurs_runtime.IntAdd(gopurs_runtime.RecordGet(__local_var_2_2, "low").IntVal, gopurs_runtime.RecordGet(__local_var_2_2, "count").IntVal)}), gopurs_runtime.Value{})
			_ = __local_var_10_10
			__local_var_11_11 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_accumulate(int64(1), int64(1), unsigned_3_3), int64(0)}), gopurs_runtime.Value{})
			_ = __local_var_11_11
			__local_var_12_12 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_choose(gopurs_runtime.RecordGet(__local_var_2_2, "count").IntVal, gopurs_runtime.RecordGet(__local_var_2_2, "high").IntVal, int64(1)), gopurs_runtime.RecordGet(__local_var_2_2, "low").IntVal}), gopurs_runtime.Value{})
			_ = __local_var_12_12
			__local_var_13_13 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_shiftState(gopurs_runtime.RecordGet(__local_var_2_2, "count").IntVal, gopurs_runtime.RecordGet(__local_var_2_2, "negative").IntVal), unsigned_3_3}), gopurs_runtime.Value{})
			_ = __local_var_13_13
			__local_var_14_14 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_divideState(int64(1), gopurs_runtime.RecordGet(__local_var_2_2, "low").IntVal), quotient_4_4}), gopurs_runtime.Value{})
			_ = __local_var_14_14
			__local_var_15_15 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_divideState(int64(2), gopurs_runtime.RecordGet(__local_var_2_2, "low").IntVal), gopurs_runtime.RecordGet(__local_var_2_2, "low").IntVal}), gopurs_runtime.Value{})
			_ = __local_var_15_15
			__local_var_16_16 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_literalState(gopurs_runtime.RecordGet(__local_var_2_2, "count").IntVal, int64(0)), unsigned_3_3}), gopurs_runtime.Value{})
			_ = __local_var_16_16
			__local_var_17_17 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_literalStep(gopurs_runtime.RecordGet(__local_var_2_2, "count").IntVal, gopurs_runtime.RecordGet(__local_var_2_2, "high").IntVal), gopurs_runtime.IntSub(gopurs_runtime.RecordGet(__local_var_2_2, "high").IntVal, gopurs_runtime.RecordGet(__local_var_2_2, "count").IntVal)}), gopurs_runtime.Value{})
			_ = __local_var_17_17
			__local_var_18_18 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_Main_literalCompare(gopurs_runtime.RecordGet(__local_var_2_2, "count").IntVal, gopurs_runtime.RecordGet(__local_var_2_2, "high").IntVal), gopurs_runtime.IntAdd(gopurs_runtime.RecordGet(__local_var_2_2, "low").IntVal, int64(2))}), gopurs_runtime.Value{})
			_ = __local_var_18_18
			var Call_local_Main_go__467072791_19_20_2 func(int64, int64) int64
			_ = Call_local_Main_go__467072791_19_20_2
			var go__467072791_19_20_2 gopurs_runtime.Value
			_ = go__467072791_19_20_2
			var Call_local_Main_go__go_19_21_3 func(int64, int64) int64
			_ = Call_local_Main_go__go_19_21_3
			var go__go_19_21_3 gopurs_runtime.Value
			_ = go__go_19_21_3
			Call_local_Main_go__467072791_19_20_2 = func(v_20_loop int64, v1_21_loop int64) int64 {
				if ((v_20_loop) == (int64(int32(v_20_loop)))) && ((v1_21_loop) == (int64(int32(v1_21_loop)))) {
					return func() int64 {
						v_20_loop := int32(v_20_loop)
						_ = v_20_loop
						v1_21_loop := int32(v1_21_loop)
						_ = v1_21_loop
					go__467072791_19_20_2:
						for {
							if false {
								continue go__467072791_19_20_2
							}
							v_20 := v_20_loop
							_ = v_20
							v1_21 := v1_21_loop
							_ = v1_21
							var __t22 int64
							{
								if (v_20) == (int32(0)) {
									__t22 = int64(v1_21)
									goto end_branch_22
								} else {

								}
							}
							{
								v_20_loop = (v_20) - (int32(1))
								v1_21_loop = (v1_21) + (v_20)
								continue go__467072791_19_20_2
								__t22 = func() int64 { panic("unreachable") }()
							}
						end_branch_22:
							return __t22
						}
					}()
				} else {

				}
			go__467072791_19_20_2:
				for {
					if false {
						continue go__467072791_19_20_2
					}
					var v_20 int64 = v_20_loop
					_ = v_20
					var v1_21 int64 = v1_21_loop
					_ = v1_21
					var __t22 int64
					{
						if (v_20) == (int64(0)) {
							__t22 = v1_21
							goto end_branch_22
						} else {

						}
					}
					{
						v_20_loop = gopurs_runtime.IntSub(v_20, int64(1))
						v1_21_loop = gopurs_runtime.IntAdd(v1_21, v_20)
						continue go__467072791_19_20_2
						__t22 = func() int64 { panic("unreachable") }()
					}
				end_branch_22:
					return __t22
				}
			}
			go__467072791_19_20_2 = gopurs_runtime.Func(func(v_20_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_21_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_go__467072791_19_20_2(v_20_loop_val.IntVal, v1_21_loop_val.IntVal))
				})
			})
			Call_local_Main_go__go_19_21_3 = func(v_20_loop int64, v1_21_loop int64) int64 {
			go__go_19_21_3:
				for {
					if false {
						continue go__go_19_21_3
					}
					var v_20 int64 = v_20_loop
					_ = v_20
					var v1_21 int64 = v1_21_loop
					_ = v1_21
					var __t23 int64
					{
						if (v_20) == (int64(0)) {
							__t23 = v1_21
							goto end_branch_23
						} else {

						}
					}
					{
						__t23 = Call_local_Main_go__467072791_19_20_2(gopurs_runtime.IntSub(v_20, int64(1)), gopurs_runtime.IntAdd(v1_21, v_20))
					}
				end_branch_23:
					return __t23
				}
			}
			go__go_19_21_3 = gopurs_runtime.Func(func(v_20_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_21_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_go__go_19_21_3(v_20_loop_val.IntVal, v1_21_loop_val.IntVal))
				})
			})
			__local_var_19_19 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_local_Main_go__467072791_19_20_2(gopurs_runtime.RecordGet(__local_var_2_2, "count").IntVal, gopurs_runtime.RecordGet(__local_var_2_2, "high").IntVal), gopurs_runtime.IntAdd(gopurs_runtime.RecordGet(__local_var_2_2, "low").IntVal, int64(5))}), gopurs_runtime.Value{})
			_ = __local_var_19_19
			var Call_local_Main_go__467072791_20_25_4 func(int64, int64) int64
			_ = Call_local_Main_go__467072791_20_25_4
			var go__467072791_20_25_4 gopurs_runtime.Value
			_ = go__467072791_20_25_4
			var Call_local_Main_go__go_20_26_5 func(int64, int64) int64
			_ = Call_local_Main_go__go_20_26_5
			var go__go_20_26_5 gopurs_runtime.Value
			_ = go__go_20_26_5
			Call_local_Main_go__467072791_20_25_4 = func(v_21_loop int64, v1_22_loop int64) int64 {
				if ((v_21_loop) == (int64(int32(v_21_loop)))) && ((v1_22_loop) == (int64(int32(v1_22_loop)))) {
					return func() int64 {
						v_21_loop := int32(v_21_loop)
						_ = v_21_loop
						v1_22_loop := int32(v1_22_loop)
						_ = v1_22_loop
					go__467072791_20_25_4:
						for {
							if false {
								continue go__467072791_20_25_4
							}
							v_21 := v_21_loop
							_ = v_21
							v1_22 := v1_22_loop
							_ = v1_22
							var __t27 int64
							{
								if (v_21) == (int32(0)) {
									__t27 = int64(v1_22)
									goto end_branch_27
								} else {

								}
							}
							{
								v_21_loop = (v_21) - (int32(1))
								v1_22_loop = (v1_22) + (v_21)
								continue go__467072791_20_25_4
								__t27 = func() int64 { panic("unreachable") }()
							}
						end_branch_27:
							return __t27
						}
					}()
				} else {

				}
			go__467072791_20_25_4:
				for {
					if false {
						continue go__467072791_20_25_4
					}
					var v_21 int64 = v_21_loop
					_ = v_21
					var v1_22 int64 = v1_22_loop
					_ = v1_22
					var __t27 int64
					{
						if (v_21) == (int64(0)) {
							__t27 = v1_22
							goto end_branch_27
						} else {

						}
					}
					{
						v_21_loop = gopurs_runtime.IntSub(v_21, int64(1))
						v1_22_loop = gopurs_runtime.IntAdd(v1_22, v_21)
						continue go__467072791_20_25_4
						__t27 = func() int64 { panic("unreachable") }()
					}
				end_branch_27:
					return __t27
				}
			}
			go__467072791_20_25_4 = gopurs_runtime.Func(func(v_21_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_22_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_go__467072791_20_25_4(v_21_loop_val.IntVal, v1_22_loop_val.IntVal))
				})
			})
			Call_local_Main_go__go_20_26_5 = func(v_21_loop int64, v1_22_loop int64) int64 {
			go__go_20_26_5:
				for {
					if false {
						continue go__go_20_26_5
					}
					var v_21 int64 = v_21_loop
					_ = v_21
					var v1_22 int64 = v1_22_loop
					_ = v1_22
					var __t28 int64
					{
						if (v_21) == (int64(0)) {
							__t28 = v1_22
							goto end_branch_28
						} else {

						}
					}
					{
						__t28 = Call_local_Main_go__467072791_20_25_4(gopurs_runtime.IntSub(v_21, int64(1)), gopurs_runtime.IntAdd(v1_22, v_21))
					}
				end_branch_28:
					return __t28
				}
			}
			go__go_20_26_5 = gopurs_runtime.Func(func(v_21_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v1_22_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Int(Call_local_Main_go__go_20_26_5(v_21_loop_val.IntVal, v1_22_loop_val.IntVal))
				})
			})
			__local_var_20_24 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___627669702("", struct {
				actual   int64
				expected int64
			}{Call_local_Main_go__467072791_20_25_4(int64(0), unsigned_3_3), unsigned_3_3}), gopurs_runtime.Value{})
			_ = __local_var_20_24
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_shiftState(v_0_loop int64, v1_1_loop int64) int64 {
	if (v_0_loop) == (int64(int32(v_0_loop))) {
		return func() int64 {
			v_0_loop := int32(v_0_loop)
			_ = v_0_loop
		shiftState:
			for {
				if false {
					continue shiftState
				}
				v_0 := v_0_loop
				_ = v_0
				var v1_1 int64 = v1_1_loop
				_ = v1_1
				var __t0 int64
				{
					if (v_0) == (int32(0)) {
						__t0 = v1_1
						goto end_branch_0
					} else {

					}
				}
				{
					v_0_loop = (v_0) - (int32(1))
					v1_1_loop = gopurs_runtime.Zshr(gopurs_runtime.Int(v1_1), gopurs_runtime.Int(int64(0))).IntVal
					continue shiftState
					__t0 = func() int64 { panic("unreachable") }()
				}
			end_branch_0:
				return __t0
			}
		}()
	} else {

	}
shiftState:
	for {
		if false {
			continue shiftState
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
			v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
			v1_1_loop = gopurs_runtime.Zshr(gopurs_runtime.Int(v1_1), gopurs_runtime.Int(int64(0))).IntVal
			continue shiftState
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_localLoop(count_0_loop int64, initial_1_loop int64) int64 {
	var count_0 int64 = count_0_loop
	_ = count_0
	var initial_1 int64 = initial_1_loop
	_ = initial_1
	var Call_local_Main_go__467072791_2_0_0 func(int64, int64) int64
	_ = Call_local_Main_go__467072791_2_0_0
	var go__467072791_2_0_0 gopurs_runtime.Value
	_ = go__467072791_2_0_0
	var Call_local_Main_go__go_2_1_1 func(int64, int64) int64
	_ = Call_local_Main_go__go_2_1_1
	var go__go_2_1_1 gopurs_runtime.Value
	_ = go__go_2_1_1
	Call_local_Main_go__467072791_2_0_0 = func(v_3_loop int64, v1_4_loop int64) int64 {
		if ((v_3_loop) == (int64(int32(v_3_loop)))) && ((v1_4_loop) == (int64(int32(v1_4_loop)))) {
			return func() int64 {
				v_3_loop := int32(v_3_loop)
				_ = v_3_loop
				v1_4_loop := int32(v1_4_loop)
				_ = v1_4_loop
			go__467072791_2_0_0:
				for {
					if false {
						continue go__467072791_2_0_0
					}
					v_3 := v_3_loop
					_ = v_3
					v1_4 := v1_4_loop
					_ = v1_4
					var __t2 int64
					{
						if (v_3) == (int32(0)) {
							__t2 = int64(v1_4)
							goto end_branch_2
						} else {

						}
					}
					{
						v_3_loop = (v_3) - (int32(1))
						v1_4_loop = (v1_4) + (v_3)
						continue go__467072791_2_0_0
						__t2 = func() int64 { panic("unreachable") }()
					}
				end_branch_2:
					return __t2
				}
			}()
		} else {

		}
	go__467072791_2_0_0:
		for {
			if false {
				continue go__467072791_2_0_0
			}
			var v_3 int64 = v_3_loop
			_ = v_3
			var v1_4 int64 = v1_4_loop
			_ = v1_4
			var __t2 int64
			{
				if (v_3) == (int64(0)) {
					__t2 = v1_4
					goto end_branch_2
				} else {

				}
			}
			{
				v_3_loop = gopurs_runtime.IntSub(v_3, int64(1))
				v1_4_loop = gopurs_runtime.IntAdd(v1_4, v_3)
				continue go__467072791_2_0_0
				__t2 = func() int64 { panic("unreachable") }()
			}
		end_branch_2:
			return __t2
		}
	}
	go__467072791_2_0_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Main_go__467072791_2_0_0(v_3_loop_val.IntVal, v1_4_loop_val.IntVal))
		})
	})
	Call_local_Main_go__go_2_1_1 = func(v_3_loop int64, v1_4_loop int64) int64 {
	go__go_2_1_1:
		for {
			if false {
				continue go__go_2_1_1
			}
			var v_3 int64 = v_3_loop
			_ = v_3
			var v1_4 int64 = v1_4_loop
			_ = v1_4
			var __t3 int64
			{
				if (v_3) == (int64(0)) {
					__t3 = v1_4
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = Call_local_Main_go__467072791_2_0_0(gopurs_runtime.IntSub(v_3, int64(1)), gopurs_runtime.IntAdd(v1_4, v_3))
			}
		end_branch_3:
			return __t3
		}
	}
	go__go_2_1_1 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Main_go__go_2_1_1(v_3_loop_val.IntVal, v1_4_loop_val.IntVal))
		})
	})
	return Call_local_Main_go__467072791_2_0_0(count_0, initial_1)
}

func Call_Main_literalStep(v_0_loop int64, v1_1_loop int64) int64 {
	if ((v_0_loop) == (int64(int32(v_0_loop)))) && ((v1_1_loop) == (int64(int32(v1_1_loop)))) {
		return func() int64 {
			v_0_loop := int32(v_0_loop)
			_ = v_0_loop
			v1_1_loop := int32(v1_1_loop)
			_ = v1_1_loop
		literalStep:
			for {
				if false {
					continue literalStep
				}
				v_0 := v_0_loop
				_ = v_0
				v1_1 := v1_1_loop
				_ = v1_1
				var __t0 int64
				{
					if (v_0) == (int32(0)) {
						__t0 = int64(v1_1)
						goto end_branch_0
					} else {

					}
				}
				{
					v_0_loop = (v_0) - (int32(1))
					v1_1_loop = (v1_1) + (int32(gopurs_runtime.IntAdd(4294967295, int64(0))))
					continue literalStep
					__t0 = func() int64 { panic("unreachable") }()
				}
			end_branch_0:
				return __t0
			}
		}()
	} else {

	}
literalStep:
	for {
		if false {
			continue literalStep
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
			v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
			v1_1_loop = gopurs_runtime.IntAdd(v1_1, int64(4294967295))
			continue literalStep
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_literalState(v_0_loop int64, v1_1_loop int64) int64 {
	if (v_0_loop) == (int64(int32(v_0_loop))) {
		return func() int64 {
			v_0_loop := int32(v_0_loop)
			_ = v_0_loop
		literalState:
			for {
				if false {
					continue literalState
				}
				v_0 := v_0_loop
				_ = v_0
				var v1_1 int64 = v1_1_loop
				_ = v1_1
				var __t0 int64
				{
					if (v_0) == (int32(0)) {
						__t0 = v1_1
						goto end_branch_0
					} else {

					}
				}
				{
					v_0_loop = (v_0) - (int32(1))
					v1_1_loop = int64(4294967295)
					continue literalState
					__t0 = func() int64 { panic("unreachable") }()
				}
			end_branch_0:
				return __t0
			}
		}()
	} else {

	}
literalState:
	for {
		if false {
			continue literalState
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
			v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
			v1_1_loop = int64(4294967295)
			continue literalState
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_literalCompare(v_0_loop int64, v1_1_loop int64) int64 {
	if ((v_0_loop) == (int64(int32(v_0_loop)))) && ((v1_1_loop) == (int64(int32(v1_1_loop)))) {
		return func() int64 {
			v_0_loop := int32(v_0_loop)
			_ = v_0_loop
			v1_1_loop := int32(v1_1_loop)
			_ = v1_1_loop
		literalCompare:
			for {
				if false {
					continue literalCompare
				}
				v_0 := v_0_loop
				_ = v_0
				v1_1 := v1_1_loop
				_ = v1_1
				var __t1 int64
				{
					if (v_0) == (int32(0)) {
						__t1 = int64(v1_1)
						goto end_branch_1
					} else {

					}
				}
				{
					var __t0 int64
					{
						if (int64(v1_1)) < (int64(4294967295)) {
							v_0_loop = (v_0) - (int32(1))
							v1_1_loop = (v1_1) + (int32(1))
							continue literalCompare
							__t0 = func() int64 { panic("unreachable") }()
							goto end_branch_0
						} else {

						}
					}
					{
						v_0_loop = (v_0) - (int32(1))
						v1_1_loop = (v1_1) - (int32(1))
						continue literalCompare
						__t0 = func() int64 { panic("unreachable") }()
					}
				end_branch_0:
					__t1 = __t0
				}
			end_branch_1:
				return __t1
			}
		}()
	} else {

	}
literalCompare:
	for {
		if false {
			continue literalCompare
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var __t1 int64
		{
			if (v_0) == (int64(0)) {
				__t1 = v1_1
				goto end_branch_1
			} else {

			}
		}
		{
			var __t0 int64
			{
				if (v1_1) < (int64(4294967295)) {
					v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
					v1_1_loop = gopurs_runtime.IntAdd(v1_1, int64(1))
					continue literalCompare
					__t0 = func() int64 { panic("unreachable") }()
					goto end_branch_0
				} else {

				}
			}
			{
				v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
				v1_1_loop = gopurs_runtime.IntSub(v1_1, int64(1))
				continue literalCompare
				__t0 = func() int64 { panic("unreachable") }()
			}
		end_branch_0:
			__t1 = __t0
		}
	end_branch_1:
		return __t1
	}
}

func Call_Main_divideState(v_0_loop int64, v1_1_loop int64) int64 {
	if (v_0_loop) == (int64(int32(v_0_loop))) {
		return func() int64 {
			v_0_loop := int32(v_0_loop)
			_ = v_0_loop
		divideState:
			for {
				if false {
					continue divideState
				}
				v_0 := v_0_loop
				_ = v_0
				var v1_1 int64 = v1_1_loop
				_ = v1_1
				var __t0 int64
				{
					if (v_0) == (int32(0)) {
						__t0 = v1_1
						goto end_branch_0
					} else {

					}
				}
				{
					v_0_loop = (v_0) - (int32(1))
					v1_1_loop = gopurs_runtime.IntDiv(v1_1, int64(-1))
					continue divideState
					__t0 = func() int64 { panic("unreachable") }()
				}
			end_branch_0:
				return __t0
			}
		}()
	} else {

	}
divideState:
	for {
		if false {
			continue divideState
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
			v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
			v1_1_loop = gopurs_runtime.IntDiv(v1_1, int64(-1))
			continue divideState
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Call_Main_choose(v_0_loop int64, v1_1_loop int64, v2_2_loop int64) int64 {
	if ((v_0_loop) == (int64(int32(v_0_loop)))) && ((v1_1_loop) == (int64(int32(v1_1_loop)))) {
		return func() int64 {
			v_0_loop := int32(v_0_loop)
			_ = v_0_loop
			v1_1_loop := int32(v1_1_loop)
			_ = v1_1_loop
		choose:
			for {
				if false {
					continue choose
				}
				v_0 := v_0_loop
				_ = v_0
				v1_1 := v1_1_loop
				_ = v1_1
				var v2_2 int64 = v2_2_loop
				_ = v2_2
				var __t1 int64
				{
					if (v_0) == (int32(0)) {
						__t1 = int64(v1_1)
						goto end_branch_1
					} else {

					}
				}
				{
					var __t0 int64
					{
						if (v1_1) < (int32(0)) {
							v_0_loop = (v_0) - (int32(1))
							v1_1_loop = (v1_1) - (int32(gopurs_runtime.IntAdd(v2_2, int64(0))))
							v2_2_loop = v2_2
							continue choose
							__t0 = func() int64 { panic("unreachable") }()
							goto end_branch_0
						} else {

						}
					}
					{
						v_0_loop = (v_0) - (int32(1))
						v1_1_loop = (v1_1) + (int32(gopurs_runtime.IntAdd(v2_2, int64(0))))
						v2_2_loop = v2_2
						continue choose
						__t0 = func() int64 { panic("unreachable") }()
					}
				end_branch_0:
					__t1 = __t0
				}
			end_branch_1:
				return __t1
			}
		}()
	} else {

	}
choose:
	for {
		if false {
			continue choose
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var v2_2 int64 = v2_2_loop
		_ = v2_2
		var __t1 int64
		{
			if (v_0) == (int64(0)) {
				__t1 = v1_1
				goto end_branch_1
			} else {

			}
		}
		{
			var __t0 int64
			{
				if (v1_1) < (int64(0)) {
					v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
					v1_1_loop = gopurs_runtime.IntSub(v1_1, v2_2)
					v2_2_loop = v2_2
					continue choose
					__t0 = func() int64 { panic("unreachable") }()
					goto end_branch_0
				} else {

				}
			}
			{
				v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
				v1_1_loop = gopurs_runtime.IntAdd(v1_1, v2_2)
				v2_2_loop = v2_2
				continue choose
				__t0 = func() int64 { panic("unreachable") }()
			}
		end_branch_0:
			__t1 = __t0
		}
	end_branch_1:
		return __t1
	}
}

func Call_Main_accumulate(v_0_loop int64, v1_1_loop int64, v2_2_loop int64) int64 {
	if ((v_0_loop) == (int64(int32(v_0_loop)))) && ((v1_1_loop) == (int64(int32(v1_1_loop)))) {
		return func() int64 {
			v_0_loop := int32(v_0_loop)
			_ = v_0_loop
			v1_1_loop := int32(v1_1_loop)
			_ = v1_1_loop
		accumulate:
			for {
				if false {
					continue accumulate
				}
				v_0 := v_0_loop
				_ = v_0
				v1_1 := v1_1_loop
				_ = v1_1
				var v2_2 int64 = v2_2_loop
				_ = v2_2
				var __t0 int64
				{
					if (v_0) == (int32(0)) {
						__t0 = int64(v1_1)
						goto end_branch_0
					} else {

					}
				}
				{
					v_0_loop = (v_0) - (int32(1))
					v1_1_loop = (v1_1) + (int32(gopurs_runtime.IntAdd(v2_2, int64(0))))
					v2_2_loop = v2_2
					continue accumulate
					__t0 = func() int64 { panic("unreachable") }()
				}
			end_branch_0:
				return __t0
			}
		}()
	} else {

	}
accumulate:
	for {
		if false {
			continue accumulate
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
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = gopurs_runtime.IntSub(v_0, int64(1))
			v1_1_loop = gopurs_runtime.IntAdd(v1_1, v2_2)
			v2_2_loop = v2_2
			continue accumulate
			__t0 = func() int64 { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}
