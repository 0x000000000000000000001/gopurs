package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_negate gopurs_runtime.Value
var once_Main_negate sync.Once

func Get_Main_negate() gopurs_runtime.Value {
	once_Main_negate.Do(func() {
		cache_Main_negate = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_negate(a_0_box.IntVal))
		})
	})
	return cache_Main_negate
}

var cache_Main_add gopurs_runtime.Value
var once_Main_add sync.Once

func Get_Main_add() gopurs_runtime.Value {
	once_Main_add.Do(func() {
		cache_Main_add = Get_Data_Semiring_add__4086903470()
	})
	return cache_Main_add
}

var cache_Main_lessThanOrEq gopurs_runtime.Value
var once_Main_lessThanOrEq sync.Once

func Get_Main_lessThanOrEq() gopurs_runtime.Value {
	once_Main_lessThanOrEq.Do(func() {
		cache_Main_lessThanOrEq = Get_Data_Ord_lessThanOrEq__3348776441()
	})
	return cache_Main_lessThanOrEq
}

var cache_Main_eq gopurs_runtime.Value
var once_Main_eq sync.Once

func Get_Main_eq() gopurs_runtime.Value {
	once_Main_eq.Do(func() {
		cache_Main_eq = Get_Data_Eq_eq__3020190905()
	})
	return cache_Main_eq
}

var cache_Main_mod gopurs_runtime.Value
var once_Main_mod sync.Once

func Get_Main_mod() gopurs_runtime.Value {
	once_Main_mod.Do(func() {
		cache_Main_mod = Get_Data_EuclideanRing_mod__478053998()
	})
	return cache_Main_mod
}

var cache_Main_sub gopurs_runtime.Value
var once_Main_sub sync.Once

func Get_Main_sub() gopurs_runtime.Value {
	once_Main_sub.Do(func() {
		cache_Main_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Main_sub
}

var cache_Main_bind gopurs_runtime.Value
var once_Main_bind sync.Once

func Get_Main_bind() gopurs_runtime.Value {
	once_Main_bind.Do(func() {
		cache_Main_bind = Get_Control_Bind_bind__2595690612()
	})
	return cache_Main_bind
}

var cache_Main_discard gopurs_runtime.Value
var once_Main_discard sync.Once

func Get_Main_discard() gopurs_runtime.Value {
	once_Main_discard.Do(func() {
		cache_Main_discard = Get_Control_Bind_discard__143037513()
	})
	return cache_Main_discard
}

var cache_Main_assertEqual gopurs_runtime.Value
var once_Main_assertEqual sync.Once

func Get_Main_assertEqual() gopurs_runtime.Value {
	once_Main_assertEqual.Do(func() {
		cache_Main_assertEqual = Get_Test_Assert_assertEqual__2867988789()
	})
	return cache_Main_assertEqual
}

var cache_Main_Leftward gopurs_runtime.Value
var once_Main_Leftward sync.Once

func Get_Main_Leftward() gopurs_runtime.Value {
	once_Main_Leftward.Do(func() {
		cache_Main_Leftward = gopurs_runtime.Value{Type: 9, IntVal: int64(1170353925), UnsafePtr: nil}
	})
	return cache_Main_Leftward
}

var cache_Main_Rightward gopurs_runtime.Value
var once_Main_Rightward sync.Once

func Get_Main_Rightward() gopurs_runtime.Value {
	once_Main_Rightward.Do(func() {
		cache_Main_Rightward = gopurs_runtime.Value{Type: 9, IntVal: int64(1220419870), UnsafePtr: nil}
	})
	return cache_Main_Rightward
}

var cache_Main_Stationary gopurs_runtime.Value
var once_Main_Stationary sync.Once

func Get_Main_Stationary() gopurs_runtime.Value {
	once_Main_Stationary.Do(func() {
		cache_Main_Stationary = gopurs_runtime.Value{Type: 9, IntVal: int64(736340654), UnsafePtr: nil}
	})
	return cache_Main_Stationary
}

var cache_Main_sumStack gopurs_runtime.Value
var once_Main_sumStack sync.Once

func Get_Main_sumStack() gopurs_runtime.Value {
	once_Main_sumStack.Do(func() {
		cache_Main_sumStack = gopurs_runtime.Func(func(stack_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_sumStack(Rebox_Main_849153993_3967238335(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](stack_0_box))))
		})
	})
	return cache_Main_sumStack
}

var cache_Main_makeStack gopurs_runtime.Value
var once_Main_makeStack sync.Once

func Get_Main_makeStack() gopurs_runtime.Value {
	once_Main_makeStack.Do(func() {
		cache_Main_makeStack = gopurs_runtime.Func2(func(size_0_box gopurs_runtime.Value, seed_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3967238335_849153993(Call_Main_makeStack(size_0_box.IntVal, seed_1_box.IntVal)))}
		})
	})
	return cache_Main_makeStack
}

var cache_Main_run gopurs_runtime.Value
var once_Main_run sync.Once

func Get_Main_run() gopurs_runtime.Value {
	once_Main_run.Do(func() {
		cache_Main_run = gopurs_runtime.Func2(func(count_0_box gopurs_runtime.Value, size_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_run(count_0_box.IntVal, size_1_box.IntVal))
		})
	})
	return cache_Main_run
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(3)))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply2(Get_Effect_Ref_read(), __local_var_1_1, gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_sumStack(Call_Main_makeStack(int64(3), __local_var_2_2.IntVal)), int64(21)}), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_sumStack(Call_Main_makeStack(int64(0), __local_var_2_2.IntVal)), int64(0)}), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(Call_Test_Assert_assertEqual_prime___2744262273(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt())), "", struct {
				actual   int64
				expected int64
				/*gopurs:record-v1:2:00610063007400750061006c,00650078007000650063007400650064*/
			}{Call_Main_run(__local_var_2_2.IntVal, int64(2)), int64(30)}), gopurs_runtime.Value{})
			_ = __local_var_5_5
			return gopurs_runtime.Apply2(Get_Effect_Console_log(), gopurs_runtime.StrWithCodeUnits("Done", 4), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Leftward struct {
	Rc uint32
}

type Constructor_Main_Rightward struct {
	Rc uint32
}

type Constructor_Main_Stationary struct {
	Rc uint32
}

func Call_Main_negate(a_0_loop int64) int64 {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.IntNegate(a_0)
}

func Call_Main_sumStack(stack_0_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
	offset int64
	row    int64
	/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
}, uint32]]) int64 {
	var stack_0 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]] = stack_0_loop
	_ = stack_0
	var Call_local_Main_go__1241963429_1_0_0 func(int64, *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]]) int64
	_ = Call_local_Main_go__1241963429_1_0_0
	var go__1241963429_1_0_0 gopurs_runtime.Value
	_ = go__1241963429_1_0_0
	var Call_local_Main_go__go_1_1_1 func(int64, *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]]) int64
	_ = Call_local_Main_go__go_1_1_1
	var go__go_1_1_1 gopurs_runtime.Value
	_ = go__go_1_1_1
	Call_local_Main_go__1241963429_1_0_0 = func(total_2_loop int64, v_3_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]]) int64 {
	go__1241963429_1_0_0:
		for {
			if false {
				continue go__1241963429_1_0_0
			}
			var total_2 int64 = total_2_loop
			_ = total_2
			var v_3 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
				offset int64
				row    int64
				/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
			}, uint32]] = v_3_loop
			_ = v_3
			var __t6 int64
			{
				if v_3 == nil {
					__t6 = total_2
					goto end_branch_6
				} else {

				}
			}
			{
				if v_3 != nil {
					var __t5 int64
					{
						var __t_tag_2 uint32 = ((v_3).V0).V1
						_ = __t_tag_2
						if uint32(__t_tag_2) == 1170353925 {
							__t5 = int64(-1)
							goto end_branch_5
						} else {

						}
					}
					{
						var __t_tag_3 uint32 = ((v_3).V0).V1
						_ = __t_tag_3
						if uint32(__t_tag_3) == 1220419870 {
							__t5 = int64(1)
							goto end_branch_5
						} else {

						}
					}
					{
						var __t_tag_4 uint32 = ((v_3).V0).V1
						_ = __t_tag_4
						if uint32(__t_tag_4) == 736340654 {
							__t5 = int64(0)
							goto end_branch_5
						} else {

						}
					}
					{
						__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
					}
				end_branch_5:
					total_2_loop = gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(total_2, ((v_3).V0).V0.offset), ((v_3).V0).V0.row), __t5)
					v_3_loop = (v_3).V1
					continue go__1241963429_1_0_0
					__t6 = func() int64 { panic("unreachable") }()
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = func() int64 { panic("Failed pattern match") }()
			}
		end_branch_6:
			return __t6
		}
	}
	go__1241963429_1_0_0 = gopurs_runtime.Func(func(total_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Main_go__1241963429_1_0_0(total_2_loop_val.IntVal, Rebox_Main_849153993_3967238335(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val))))
		})
	})
	Call_local_Main_go__go_1_1_1 = func(total_2_loop int64, v_3_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]]) int64 {
	go__go_1_1_1:
		for {
			if false {
				continue go__go_1_1_1
			}
			var total_2 int64 = total_2_loop
			_ = total_2
			var v_3 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
				offset int64
				row    int64
				/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
			}, uint32]] = v_3_loop
			_ = v_3
			var __t11 int64
			{
				if v_3 == nil {
					__t11 = total_2
					goto end_branch_11
				} else {

				}
			}
			{
				if v_3 != nil {
					var __t10 int64
					{
						var __t_tag_7 uint32 = ((v_3).V0).V1
						_ = __t_tag_7
						if uint32(__t_tag_7) == 1170353925 {
							__t10 = int64(-1)
							goto end_branch_10
						} else {

						}
					}
					{
						var __t_tag_8 uint32 = ((v_3).V0).V1
						_ = __t_tag_8
						if uint32(__t_tag_8) == 1220419870 {
							__t10 = int64(1)
							goto end_branch_10
						} else {

						}
					}
					{
						var __t_tag_9 uint32 = ((v_3).V0).V1
						_ = __t_tag_9
						if uint32(__t_tag_9) == 736340654 {
							__t10 = int64(0)
							goto end_branch_10
						} else {

						}
					}
					{
						__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
					}
				end_branch_10:
					__t11 = Call_local_Main_go__1241963429_1_0_0(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(gopurs_runtime.IntAdd(total_2, ((v_3).V0).V0.offset), ((v_3).V0).V0.row), __t10), (v_3).V1)
					goto end_branch_11
				} else {

				}
			}
			{
				__t11 = func() int64 { panic("Failed pattern match") }()
			}
		end_branch_11:
			return __t11
		}
	}
	go__go_1_1_1 = gopurs_runtime.Func(func(total_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Main_go__go_1_1_1(total_2_loop_val.IntVal, Rebox_Main_849153993_3967238335(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val))))
		})
	})
	return Call_local_Main_go__1241963429_1_0_0(int64(0), stack_0)
}

func Call_Main_makeStack(size_0_loop int64, seed_1_loop int64) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
	offset int64
	row    int64
	/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
}, uint32]] {
	var size_0 int64 = size_0_loop
	_ = size_0
	var seed_1 int64 = seed_1_loop
	_ = seed_1
	var Call_local_Main_go__1658412951_2_0_2 func(int64, *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]]
	_ = Call_local_Main_go__1658412951_2_0_2
	var go__1658412951_2_0_2 gopurs_runtime.Value
	_ = go__1658412951_2_0_2
	var Call_local_Main_go__go_2_1_3 func(int64, *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]]
	_ = Call_local_Main_go__go_2_1_3
	var go__go_2_1_3 gopurs_runtime.Value
	_ = go__go_2_1_3
	Call_local_Main_go__1658412951_2_0_2 = func(n_3_loop int64, rest_4_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]] {
	go__1658412951_2_0_2:
		for {
			if false {
				continue go__1658412951_2_0_2
			}
			var n_3 int64 = n_3_loop
			_ = n_3
			var rest_4 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
				offset int64
				row    int64
				/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
			}, uint32]] = rest_4_loop
			_ = rest_4
			var __t4 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
				offset int64
				row    int64
				/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
			}, uint32]]
			{
				if (n_3) <= (int64(0)) {
					__t4 = rest_4
					goto end_branch_4
				} else {

				}
			}
			{
				var __t3 uint32
				{
					if (gopurs_runtime.IntMod(n_3, int64(3))) == (int64(0)) {
						__t3 = 736340654
						goto end_branch_3
					} else {

					}
				}
				{
					var __t2 uint32
					{
						if (gopurs_runtime.IntMod(n_3, int64(3))) == (int64(1)) {
							__t2 = 1170353925
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = 1220419870
					}
				end_branch_2:
					__t3 = __t2
				}
			end_branch_3:
				n_3_loop = gopurs_runtime.IntSub(n_3, int64(1))
				rest_4_loop = (&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
					offset int64
					row    int64
					/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
				}, uint32]]{1, (&Constructor_Data_Tuple_Tuple[struct {
					offset int64
					row    int64
					/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
				}, uint32]{1, struct {
					offset int64
					row    int64
					/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
				}{gopurs_runtime.IntAdd(seed_1, n_3), n_3}, __t3}), rest_4})
				continue go__1658412951_2_0_2
				__t4 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
					offset int64
					row    int64
					/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
				}, uint32]] {
					panic("unreachable")
				}()
			}
		end_branch_4:
			return __t4
		}
	}
	go__1658412951_2_0_2 = gopurs_runtime.Func(func(n_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(rest_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3967238335_849153993(Call_local_Main_go__1658412951_2_0_2(n_3_loop_val.IntVal, Rebox_Main_849153993_3967238335(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](rest_4_loop_val)))))}
		})
	})
	Call_local_Main_go__go_2_1_3 = func(n_3_loop int64, rest_4_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]] {
	go__go_2_1_3:
		for {
			if false {
				continue go__go_2_1_3
			}
			var n_3 int64 = n_3_loop
			_ = n_3
			var rest_4 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
				offset int64
				row    int64
				/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
			}, uint32]] = rest_4_loop
			_ = rest_4
			var __t7 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
				offset int64
				row    int64
				/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
			}, uint32]]
			{
				if (n_3) <= (int64(0)) {
					__t7 = rest_4
					goto end_branch_7
				} else {

				}
			}
			{
				var __t6 uint32
				{
					if (gopurs_runtime.IntMod(n_3, int64(3))) == (int64(0)) {
						__t6 = 736340654
						goto end_branch_6
					} else {

					}
				}
				{
					var __t5 uint32
					{
						if (gopurs_runtime.IntMod(n_3, int64(3))) == (int64(1)) {
							__t5 = 1170353925
							goto end_branch_5
						} else {

						}
					}
					{
						__t5 = 1220419870
					}
				end_branch_5:
					__t6 = __t5
				}
			end_branch_6:
				__t7 = Call_local_Main_go__1658412951_2_0_2(gopurs_runtime.IntSub(n_3, int64(1)), (&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
					offset int64
					row    int64
					/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
				}, uint32]]{1, (&Constructor_Data_Tuple_Tuple[struct {
					offset int64
					row    int64
					/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
				}, uint32]{1, struct {
					offset int64
					row    int64
					/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
				}{gopurs_runtime.IntAdd(seed_1, n_3), n_3}, __t6}), rest_4}))
			}
		end_branch_7:
			return __t7
		}
	}
	go__go_2_1_3 = gopurs_runtime.Func(func(n_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(rest_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3967238335_849153993(Call_local_Main_go__go_2_1_3(n_3_loop_val.IntVal, Rebox_Main_849153993_3967238335(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](rest_4_loop_val)))))}
		})
	})
	return Call_local_Main_go__1658412951_2_0_2(size_0, (*Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]])(nil))
}

func Call_Main_run(count_0_loop int64, size_1_loop int64) int64 {
	var count_0 int64 = count_0_loop
	_ = count_0
	var size_1 int64 = size_1_loop
	_ = size_1
	var Call_local_Main_go__467072791_2_0_4 func(int64, int64) int64
	_ = Call_local_Main_go__467072791_2_0_4
	var go__467072791_2_0_4 gopurs_runtime.Value
	_ = go__467072791_2_0_4
	var Call_local_Main_go__go_2_1_5 func(int64, int64) int64
	_ = Call_local_Main_go__go_2_1_5
	var go__go_2_1_5 gopurs_runtime.Value
	_ = go__go_2_1_5
	Call_local_Main_go__467072791_2_0_4 = func(n_3_loop int64, total_4_loop int64) int64 {
	go__467072791_2_0_4:
		for {
			if false {
				continue go__467072791_2_0_4
			}
			var n_3 int64 = n_3_loop
			_ = n_3
			var total_4 int64 = total_4_loop
			_ = total_4
			var __t2 int64
			{
				if (n_3) <= (int64(0)) {
					__t2 = total_4
					goto end_branch_2
				} else {

				}
			}
			{
				n_3_loop = gopurs_runtime.IntSub(n_3, int64(1))
				total_4_loop = gopurs_runtime.IntAdd(total_4, Call_Main_sumStack(Call_Main_makeStack(size_1, gopurs_runtime.IntMod(n_3, int64(11)))))
				continue go__467072791_2_0_4
				__t2 = func() int64 { panic("unreachable") }()
			}
		end_branch_2:
			return __t2
		}
	}
	go__467072791_2_0_4 = gopurs_runtime.Func(func(n_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(total_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Main_go__467072791_2_0_4(n_3_loop_val.IntVal, total_4_loop_val.IntVal))
		})
	})
	Call_local_Main_go__go_2_1_5 = func(n_3_loop int64, total_4_loop int64) int64 {
	go__go_2_1_5:
		for {
			if false {
				continue go__go_2_1_5
			}
			var n_3 int64 = n_3_loop
			_ = n_3
			var total_4 int64 = total_4_loop
			_ = total_4
			var __t3 int64
			{
				if (n_3) <= (int64(0)) {
					__t3 = total_4
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = Call_local_Main_go__467072791_2_0_4(gopurs_runtime.IntSub(n_3, int64(1)), gopurs_runtime.IntAdd(total_4, Call_Main_sumStack(Call_Main_makeStack(size_1, gopurs_runtime.IntMod(n_3, int64(11))))))
			}
		end_branch_3:
			return __t3
		}
	}
	go__go_2_1_5 = gopurs_runtime.Func(func(n_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(total_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_local_Main_go__go_2_1_5(n_3_loop_val.IntVal, total_4_loop_val.IntVal))
		})
	})
	return Call_local_Main_go__467072791_2_0_4(count_0, int64(0))
}

func Rebox_Main_138441832_3505406419(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[struct {
	offset int64
	row    int64
	/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
}, uint32] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]{}
	out.V0 = func() struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			offset int64
			row    int64
			/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
		}{}
		clone.offset = gopurs_runtime.RecordGet(orig, "offset").IntVal
		clone.row = gopurs_runtime.RecordGet(orig, "row").IntVal
		return clone
	}()
	out.V1 = uint32(in.V1.IntVal)
	return out
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[int64])(unsafe.Pointer(in))
}

func Rebox_Main_3505406419_138441832(in *Constructor_Data_Tuple_Tuple[struct {
	offset int64
	row    int64
	/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
}, uint32]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordShape2(&GopursRecordKeys_4d61696e_0, gopurs_runtime.Int(orig.offset), gopurs_runtime.Int(orig.row))
	}()
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
	return out
}

func Rebox_Main_3967238335_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
	offset int64
	row    int64
	/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
}, uint32]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3505406419_138441832(in.V0))}
	out.V1 = Rebox_Main_3967238335_849153993(in.V1)
	return out
}

func Rebox_Main_849153993_3967238335(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
	offset int64
	row    int64
	/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
}, uint32]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[struct {
		offset int64
		row    int64
		/*gopurs:record-v1:2:006f00660066007300650074,0072006f0077*/
	}, uint32]]{}
	out.V0 = Rebox_Main_138441832_3505406419(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	out.V1 = Rebox_Main_849153993_3967238335(in.V1)
	return out
}

var GopursRecordKeys_4d61696e_0 = gopurs_runtime.RecordKeys2{K0: "offset", K1: "row"}
