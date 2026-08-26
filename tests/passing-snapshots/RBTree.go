package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_R gopurs_runtime.Value
var once_Main_R sync.Once

func Get_Main_R() gopurs_runtime.Value {
	once_Main_R.Do(func() {
		cache_Main_R = gopurs_runtime.Value{Type: 9, IntVal: int64(3558538316), UnsafePtr: nil}
	})
	return cache_Main_R
}

var cache_Main_B gopurs_runtime.Value
var once_Main_B sync.Once

func Get_Main_B() gopurs_runtime.Value {
	once_Main_B.Do(func() {
		cache_Main_B = gopurs_runtime.Value{Type: 9, IntVal: int64(4250879068), UnsafePtr: nil}
	})
	return cache_Main_B
}

var cache_Main_E gopurs_runtime.Value
var once_Main_E sync.Once

func Get_Main_E() gopurs_runtime.Value {
	once_Main_E.Do(func() {
		cache_Main_E = gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer((*Constructor_Main_T)(nil))}
	})
	return cache_Main_E
}

var cache_Main_T gopurs_runtime.Value
var once_Main_T sync.Once

func Get_Main_T() gopurs_runtime.Value {
	once_Main_T.Do(func() {
		cache_Main_T = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer((&Constructor_Main_T{1, uint32(value0.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Main_T](value1), value2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_T](value3)}))}
					})
				})
			})
		})
	})
	return cache_Main_T
}

var cache_Main_max gopurs_runtime.Value
var once_Main_max sync.Once

func Get_Main_max() gopurs_runtime.Value {
	once_Main_max.Do(func() {
		cache_Main_max = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_max(x_0_box.IntVal, y_1_box.IntVal))
		})
	})
	return cache_Main_max
}

var cache_Main_describe gopurs_runtime.Value
var once_Main_describe sync.Once

func Get_Main_describe() gopurs_runtime.Value {
	once_Main_describe.Do(func() {
		cache_Main_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Red-Black Tree (100k Worst-Case Insertions):"))
	})
	return cache_Main_describe
}

var cache_Main_depth gopurs_runtime.Value
var once_Main_depth sync.Once

func Get_Main_depth() gopurs_runtime.Value {
	once_Main_depth.Do(func() {
		cache_Main_depth = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_depth(gopurs_runtime.CoerceToStruct[Constructor_Main_T](v_0_box)))
		})
	})
	return cache_Main_depth
}

var cache_Main_balance gopurs_runtime.Value
var once_Main_balance sync.Once

func Get_Main_balance() gopurs_runtime.Value {
	once_Main_balance.Do(func() {
		cache_Main_balance = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer(Call_Main_balance(uint32(v_0_box.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Main_T](v1_1_box), v2_2_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_T](v3_3_box)))}
		})
	})
	return cache_Main_balance
}

var cache_Main_insert gopurs_runtime.Value
var once_Main_insert sync.Once

func Get_Main_insert() gopurs_runtime.Value {
	once_Main_insert.Do(func() {
		cache_Main_insert = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer(Call_Main_insert(x_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_T](s_1_box)))}
		})
	})
	return cache_Main_insert
}

var cache_Main_buildTree gopurs_runtime.Value
var once_Main_buildTree sync.Once

func Get_Main_buildTree() gopurs_runtime.Value {
	once_Main_buildTree.Do(func() {
		cache_Main_buildTree = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer(Call_Main_buildTree(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_T](v1_1_box)))}
		})
	})
	return cache_Main_buildTree
}

var cache_Main_act gopurs_runtime.Value
var once_Main_act sync.Once

func Get_Main_act() gopurs_runtime.Value {
	once_Main_act.Do(func() {
		cache_Main_act = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Main_depth(Call_Main_buildTree(100000, (*Constructor_Main_T)(nil))))).StrVal()))
	})
	return cache_Main_act
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			_dollar___unused_0_0 := gopurs_runtime.Apply(Get_Main_describe(), gopurs_runtime.Value{})
			_ = _dollar___unused_0_0
			return gopurs_runtime.Apply(Get_Main_act(), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_R struct {
	Rc uint32
}

type Constructor_Main_B struct {
	Rc uint32
}

type Constructor_Main_E struct {
	Rc uint32
}

type Constructor_Main_T struct {
	Rc uint32
	V0 uint32
	V1 *Constructor_Main_T
	V2 int64
	V3 *Constructor_Main_T
}

func Call_Main_max(x_0_loop int64, y_1_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	var y_1 int64 = y_1_loop
	_ = y_1
	var __t0 int64
	{
		if (x_0) > (y_1) {
			__t0 = x_0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = y_1
	}
end_branch_0:
	return __t0
}

func Call_Main_depth(v_0_loop *Constructor_Main_T) int64 {
depth:
	for {
		if false {
			continue depth
		}
		var v_0 *Constructor_Main_T = v_0_loop
		_ = v_0
		var __t3 int64
		{
			if v_0 == nil {
				__t3 = 0
				goto end_branch_3
			} else {

			}
		}
		{
			if v_0 != nil {
				// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
				__local_var_1_0 := Call_Main_depth((v_0).V1)
				_ = __local_var_1_0
				// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=Int
				__local_var_2_1 := Call_Main_depth((v_0).V3)
				_ = __local_var_2_1
				var __t2 int64
				{
					if (__local_var_1_0) > (__local_var_2_1) {
						__t2 = __local_var_1_0
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = __local_var_2_1
				}
			end_branch_2:
				__t3 = (1) + (__t2)
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
		}
	end_branch_3:
		return __t3
	}
}

func Call_Main_balance(v_0_loop uint32, v1_1_loop *Constructor_Main_T, v2_2_loop int64, v3_3_loop *Constructor_Main_T) *Constructor_Main_T {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var v1_1 *Constructor_Main_T = v1_1_loop
	_ = v1_1
	var v2_2 int64 = v2_2_loop
	_ = v2_2
	var v3_3 *Constructor_Main_T = v3_3_loop
	_ = v3_3
	var __t85 *Constructor_Main_T
	{
		if v_0 == 4250879068 {
			var __t84 *Constructor_Main_T
			{
				if v1_1 != nil {
					var __t71 *Constructor_Main_T
					{
						var __t_tag_0 uint32 = (v1_1).V0
						if uint32(__t_tag_0) == 3558538316 {
							var __t58 *Constructor_Main_T
							{
								var __t_tag_1 *Constructor_Main_T = (v1_1).V1
								if __t_tag_1 != nil {
									var __t30 *Constructor_Main_T
									{
										var __t_tag_2 uint32 = ((v1_1).V1).V0
										if uint32(__t_tag_2) == 3558538316 {
											__t30 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, ((v1_1).V1).V1, ((v1_1).V1).V2, ((v1_1).V1).V3}), (v1_1).V2, (&Constructor_Main_T{1, 4250879068, (v1_1).V3, v2_2, v3_3})})
											goto end_branch_30
										} else {

										}
									}
									{
										var __t_tag_3 *Constructor_Main_T = (v1_1).V3
										if __t_tag_3 != nil {
											var __t17 *Constructor_Main_T
											{
												var __t_tag_4 uint32 = ((v1_1).V3).V0
												if uint32(__t_tag_4) == 3558538316 {
													__t17 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, (v1_1).V1, (v1_1).V2, ((v1_1).V3).V1}), ((v1_1).V3).V2, (&Constructor_Main_T{1, 4250879068, ((v1_1).V3).V3, v2_2, v3_3})})
													goto end_branch_17
												} else {

												}
											}
											{
												var __t_and_6 bool = false
												if v3_3 != nil {

													var __t_tag_5 uint32 = (v3_3).V0
													__t_and_6 = (uint32(__t_tag_5) == 3558538316)
												}
												if __t_and_6 {
													var __t16 *Constructor_Main_T
													{
														var __t_tag_7 *Constructor_Main_T = (v3_3).V1
														if __t_tag_7 != nil {
															var __t12 *Constructor_Main_T
															{
																var __t_tag_8 uint32 = ((v3_3).V1).V0
																if uint32(__t_tag_8) == 3558538316 {
																	__t12 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, ((v3_3).V1).V1}), ((v3_3).V1).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V1).V3, (v3_3).V2, (v3_3).V3})})
																	goto end_branch_12
																} else {

																}
															}
															{
																var __t_tag_9 *Constructor_Main_T = (v3_3).V3
																var __t_and_11 bool = false
																if __t_tag_9 != nil {

																	var __t_tag_10 uint32 = ((v3_3).V3).V0
																	__t_and_11 = (uint32(__t_tag_10) == 3558538316)
																}
																if __t_and_11 {
																	__t12 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
																	goto end_branch_12
																} else {

																}
															}
															{
																__t12 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
															}
														end_branch_12:
															__t16 = __t12
															goto end_branch_16
														} else {

														}
													}
													{
														var __t_tag_13 *Constructor_Main_T = (v3_3).V3
														var __t_and_15 bool = false
														if __t_tag_13 != nil {

															var __t_tag_14 uint32 = ((v3_3).V3).V0
															__t_and_15 = (uint32(__t_tag_14) == 3558538316)
														}
														if __t_and_15 {
															__t16 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
															goto end_branch_16
														} else {

														}
													}
													{
														__t16 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
													}
												end_branch_16:
													__t17 = __t16
													goto end_branch_17
												} else {

												}
											}
											{
												__t17 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
											}
										end_branch_17:
											__t30 = __t17
											goto end_branch_30
										} else {

										}
									}
									{
										var __t_and_19 bool = false
										if v3_3 != nil {

											var __t_tag_18 uint32 = (v3_3).V0
											__t_and_19 = (uint32(__t_tag_18) == 3558538316)
										}
										if __t_and_19 {
											var __t29 *Constructor_Main_T
											{
												var __t_tag_20 *Constructor_Main_T = (v3_3).V1
												if __t_tag_20 != nil {
													var __t25 *Constructor_Main_T
													{
														var __t_tag_21 uint32 = ((v3_3).V1).V0
														if uint32(__t_tag_21) == 3558538316 {
															__t25 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, ((v3_3).V1).V1}), ((v3_3).V1).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V1).V3, (v3_3).V2, (v3_3).V3})})
															goto end_branch_25
														} else {

														}
													}
													{
														var __t_tag_22 *Constructor_Main_T = (v3_3).V3
														var __t_and_24 bool = false
														if __t_tag_22 != nil {

															var __t_tag_23 uint32 = ((v3_3).V3).V0
															__t_and_24 = (uint32(__t_tag_23) == 3558538316)
														}
														if __t_and_24 {
															__t25 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
															goto end_branch_25
														} else {

														}
													}
													{
														__t25 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
													}
												end_branch_25:
													__t29 = __t25
													goto end_branch_29
												} else {

												}
											}
											{
												var __t_tag_26 *Constructor_Main_T = (v3_3).V3
												var __t_and_28 bool = false
												if __t_tag_26 != nil {

													var __t_tag_27 uint32 = ((v3_3).V3).V0
													__t_and_28 = (uint32(__t_tag_27) == 3558538316)
												}
												if __t_and_28 {
													__t29 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
													goto end_branch_29
												} else {

												}
											}
											{
												__t29 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
											}
										end_branch_29:
											__t30 = __t29
											goto end_branch_30
										} else {

										}
									}
									{
										__t30 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
									}
								end_branch_30:
									__t58 = __t30
									goto end_branch_58
								} else {

								}
							}
							{
								var __t_tag_31 *Constructor_Main_T = (v1_1).V3
								if __t_tag_31 != nil {
									var __t45 *Constructor_Main_T
									{
										var __t_tag_32 uint32 = ((v1_1).V3).V0
										if uint32(__t_tag_32) == 3558538316 {
											__t45 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, (v1_1).V1, (v1_1).V2, ((v1_1).V3).V1}), ((v1_1).V3).V2, (&Constructor_Main_T{1, 4250879068, ((v1_1).V3).V3, v2_2, v3_3})})
											goto end_branch_45
										} else {

										}
									}
									{
										var __t_and_34 bool = false
										if v3_3 != nil {

											var __t_tag_33 uint32 = (v3_3).V0
											__t_and_34 = (uint32(__t_tag_33) == 3558538316)
										}
										if __t_and_34 {
											var __t44 *Constructor_Main_T
											{
												var __t_tag_35 *Constructor_Main_T = (v3_3).V1
												if __t_tag_35 != nil {
													var __t40 *Constructor_Main_T
													{
														var __t_tag_36 uint32 = ((v3_3).V1).V0
														if uint32(__t_tag_36) == 3558538316 {
															__t40 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, ((v3_3).V1).V1}), ((v3_3).V1).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V1).V3, (v3_3).V2, (v3_3).V3})})
															goto end_branch_40
														} else {

														}
													}
													{
														var __t_tag_37 *Constructor_Main_T = (v3_3).V3
														var __t_and_39 bool = false
														if __t_tag_37 != nil {

															var __t_tag_38 uint32 = ((v3_3).V3).V0
															__t_and_39 = (uint32(__t_tag_38) == 3558538316)
														}
														if __t_and_39 {
															__t40 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
															goto end_branch_40
														} else {

														}
													}
													{
														__t40 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
													}
												end_branch_40:
													__t44 = __t40
													goto end_branch_44
												} else {

												}
											}
											{
												var __t_tag_41 *Constructor_Main_T = (v3_3).V3
												var __t_and_43 bool = false
												if __t_tag_41 != nil {

													var __t_tag_42 uint32 = ((v3_3).V3).V0
													__t_and_43 = (uint32(__t_tag_42) == 3558538316)
												}
												if __t_and_43 {
													__t44 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
													goto end_branch_44
												} else {

												}
											}
											{
												__t44 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
											}
										end_branch_44:
											__t45 = __t44
											goto end_branch_45
										} else {

										}
									}
									{
										__t45 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
									}
								end_branch_45:
									__t58 = __t45
									goto end_branch_58
								} else {

								}
							}
							{
								var __t_and_47 bool = false
								if v3_3 != nil {

									var __t_tag_46 uint32 = (v3_3).V0
									__t_and_47 = (uint32(__t_tag_46) == 3558538316)
								}
								if __t_and_47 {
									var __t57 *Constructor_Main_T
									{
										var __t_tag_48 *Constructor_Main_T = (v3_3).V1
										if __t_tag_48 != nil {
											var __t53 *Constructor_Main_T
											{
												var __t_tag_49 uint32 = ((v3_3).V1).V0
												if uint32(__t_tag_49) == 3558538316 {
													__t53 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, ((v3_3).V1).V1}), ((v3_3).V1).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V1).V3, (v3_3).V2, (v3_3).V3})})
													goto end_branch_53
												} else {

												}
											}
											{
												var __t_tag_50 *Constructor_Main_T = (v3_3).V3
												var __t_and_52 bool = false
												if __t_tag_50 != nil {

													var __t_tag_51 uint32 = ((v3_3).V3).V0
													__t_and_52 = (uint32(__t_tag_51) == 3558538316)
												}
												if __t_and_52 {
													__t53 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
													goto end_branch_53
												} else {

												}
											}
											{
												__t53 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
											}
										end_branch_53:
											__t57 = __t53
											goto end_branch_57
										} else {

										}
									}
									{
										var __t_tag_54 *Constructor_Main_T = (v3_3).V3
										var __t_and_56 bool = false
										if __t_tag_54 != nil {

											var __t_tag_55 uint32 = ((v3_3).V3).V0
											__t_and_56 = (uint32(__t_tag_55) == 3558538316)
										}
										if __t_and_56 {
											__t57 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
											goto end_branch_57
										} else {

										}
									}
									{
										__t57 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
									}
								end_branch_57:
									__t58 = __t57
									goto end_branch_58
								} else {

								}
							}
							{
								__t58 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
							}
						end_branch_58:
							__t71 = __t58
							goto end_branch_71
						} else {

						}
					}
					{
						var __t_and_60 bool = false
						if v3_3 != nil {

							var __t_tag_59 uint32 = (v3_3).V0
							__t_and_60 = (uint32(__t_tag_59) == 3558538316)
						}
						if __t_and_60 {
							var __t70 *Constructor_Main_T
							{
								var __t_tag_61 *Constructor_Main_T = (v3_3).V1
								if __t_tag_61 != nil {
									var __t66 *Constructor_Main_T
									{
										var __t_tag_62 uint32 = ((v3_3).V1).V0
										if uint32(__t_tag_62) == 3558538316 {
											__t66 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, ((v3_3).V1).V1}), ((v3_3).V1).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V1).V3, (v3_3).V2, (v3_3).V3})})
											goto end_branch_66
										} else {

										}
									}
									{
										var __t_tag_63 *Constructor_Main_T = (v3_3).V3
										var __t_and_65 bool = false
										if __t_tag_63 != nil {

											var __t_tag_64 uint32 = ((v3_3).V3).V0
											__t_and_65 = (uint32(__t_tag_64) == 3558538316)
										}
										if __t_and_65 {
											__t66 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
											goto end_branch_66
										} else {

										}
									}
									{
										__t66 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
									}
								end_branch_66:
									__t70 = __t66
									goto end_branch_70
								} else {

								}
							}
							{
								var __t_tag_67 *Constructor_Main_T = (v3_3).V3
								var __t_and_69 bool = false
								if __t_tag_67 != nil {

									var __t_tag_68 uint32 = ((v3_3).V3).V0
									__t_and_69 = (uint32(__t_tag_68) == 3558538316)
								}
								if __t_and_69 {
									__t70 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
									goto end_branch_70
								} else {

								}
							}
							{
								__t70 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
							}
						end_branch_70:
							__t71 = __t70
							goto end_branch_71
						} else {

						}
					}
					{
						__t71 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
					}
				end_branch_71:
					__t84 = __t71
					goto end_branch_84
				} else {

				}
			}
			{
				var __t_and_73 bool = false
				if v3_3 != nil {

					var __t_tag_72 uint32 = (v3_3).V0
					__t_and_73 = (uint32(__t_tag_72) == 3558538316)
				}
				if __t_and_73 {
					var __t83 *Constructor_Main_T
					{
						var __t_tag_74 *Constructor_Main_T = (v3_3).V1
						if __t_tag_74 != nil {
							var __t79 *Constructor_Main_T
							{
								var __t_tag_75 uint32 = ((v3_3).V1).V0
								if uint32(__t_tag_75) == 3558538316 {
									__t79 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, ((v3_3).V1).V1}), ((v3_3).V1).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V1).V3, (v3_3).V2, (v3_3).V3})})
									goto end_branch_79
								} else {

								}
							}
							{
								var __t_tag_76 *Constructor_Main_T = (v3_3).V3
								var __t_and_78 bool = false
								if __t_tag_76 != nil {

									var __t_tag_77 uint32 = ((v3_3).V3).V0
									__t_and_78 = (uint32(__t_tag_77) == 3558538316)
								}
								if __t_and_78 {
									__t79 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
									goto end_branch_79
								} else {

								}
							}
							{
								__t79 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
							}
						end_branch_79:
							__t83 = __t79
							goto end_branch_83
						} else {

						}
					}
					{
						var __t_tag_80 *Constructor_Main_T = (v3_3).V3
						var __t_and_82 bool = false
						if __t_tag_80 != nil {

							var __t_tag_81 uint32 = ((v3_3).V3).V0
							__t_and_82 = (uint32(__t_tag_81) == 3558538316)
						}
						if __t_and_82 {
							__t83 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, v1_1, v2_2, (v3_3).V1}), (v3_3).V2, (&Constructor_Main_T{1, 4250879068, ((v3_3).V3).V1, ((v3_3).V3).V2, ((v3_3).V3).V3})})
							goto end_branch_83
						} else {

						}
					}
					{
						__t83 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
					}
				end_branch_83:
					__t84 = __t83
					goto end_branch_84
				} else {

				}
			}
			{
				__t84 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
			}
		end_branch_84:
			__t85 = __t84
			goto end_branch_85
		} else {

		}
	}
	{
		__t85 = (&Constructor_Main_T{1, v_0, v1_1, v2_2, v3_3})
	}
end_branch_85:
	return __t85
}

func Call_Main_insert(x_0_loop int64, s_1_loop *Constructor_Main_T) *Constructor_Main_T {
	var x_0 int64 = x_0_loop
	_ = x_0
	var s_1 *Constructor_Main_T = s_1_loop
	_ = s_1
	var ins_2_0_0 gopurs_runtime.Value
	_ = ins_2_0_0
	ins_2_0_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t3 *Constructor_Main_T
		{
			if v_3.Type == 9 && v_3.IntVal == 990467018 && v_3.UnsafePtr == nil {
				__t3 = (&Constructor_Main_T{1, 3558538316, (*Constructor_Main_T)(nil), x_0, (*Constructor_Main_T)(nil)})
				goto end_branch_3
			} else {

			}
		}
		{
			if v_3.Type == 9 && v_3.IntVal == 990467018 && v_3.UnsafePtr != nil {
				var __t2 *Constructor_Main_T
				{
					if (x_0) < ((*Constructor_Main_T)(v_3.UnsafePtr).V2) {
						__t2 = Call_Main_balance((*Constructor_Main_T)(v_3.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Main_T](gopurs_runtime.Apply(ins_2_0_0, gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer((*Constructor_Main_T)(v_3.UnsafePtr).V1)})), (*Constructor_Main_T)(v_3.UnsafePtr).V2, (*Constructor_Main_T)(v_3.UnsafePtr).V3)
						goto end_branch_2
					} else {

					}
				}
				{
					var __t1 *Constructor_Main_T
					{
						if (x_0) > ((*Constructor_Main_T)(v_3.UnsafePtr).V2) {
							__t1 = Call_Main_balance((*Constructor_Main_T)(v_3.UnsafePtr).V0, (*Constructor_Main_T)(v_3.UnsafePtr).V1, (*Constructor_Main_T)(v_3.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Main_T](gopurs_runtime.Apply(ins_2_0_0, gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer((*Constructor_Main_T)(v_3.UnsafePtr).V3)})))
							goto end_branch_1
						} else {

						}
					}
					{
						__t1 = (&Constructor_Main_T{1, (*Constructor_Main_T)(v_3.UnsafePtr).V0, (*Constructor_Main_T)(v_3.UnsafePtr).V1, (*Constructor_Main_T)(v_3.UnsafePtr).V2, (*Constructor_Main_T)(v_3.UnsafePtr).V3})
					}
				end_branch_1:
					__t2 = __t1
				}
			end_branch_2:
				__t3 = __t2
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = gopurs_runtime.CoerceToStruct[Constructor_Main_T](func() gopurs_runtime.Value { panic("Failed pattern match") }())
		}
	end_branch_3:
		return gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer(__t3)}
	})
	// TAST (Let): __local_var_3_4 shape=App(Other) bindingType=(ADT ["Main","Tree"] [])
	__local_var_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Main_T](gopurs_runtime.Apply(ins_2_0_0, gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer(s_1)}))
	_ = __local_var_3_4
	var __t5 *Constructor_Main_T
	{
		if __local_var_3_4 != nil {
			__t5 = (&Constructor_Main_T{1, 4250879068, (__local_var_3_4).V1, (__local_var_3_4).V2, (__local_var_3_4).V3})
			goto end_branch_5
		} else {

		}
	}
	{
		if __local_var_3_4 == nil {
			__t5 = (*Constructor_Main_T)(nil)
			goto end_branch_5
		} else {

		}
	}
	{
		__t5 = gopurs_runtime.CoerceToStruct[Constructor_Main_T](func() gopurs_runtime.Value { panic("Failed pattern match") }())
	}
end_branch_5:
	return __t5
}

func Call_Main_buildTree(v_0_loop int64, v1_1_loop *Constructor_Main_T) *Constructor_Main_T {
buildTree:
	for {
		if false {
			continue buildTree
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Main_T = v1_1_loop
		_ = v1_1
		var __t0 *Constructor_Main_T
		{
			if (v_0) == (0) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (1)
			v1_1_loop = Call_Main_insert(v_0, v1_1)
			continue buildTree
			__t0 = gopurs_runtime.CoerceToStruct[Constructor_Main_T](gopurs_runtime.Value{})
		}
	end_branch_0:
		return __t0
	}
}
