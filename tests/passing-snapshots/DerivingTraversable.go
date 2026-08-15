package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_identity gopurs_runtime.Value
var once_Main_identity sync.Once

func Get_Main_identity() gopurs_runtime.Value {
	once_Main_identity.Do(func() {
		cache_Main_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_identity(x_0_box)
		})
	})
	return cache_Main_identity
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl())}))}
	})
	return cache_Main_eqArray
}

var cache_Main_pure gopurs_runtime.Value
var once_Main_pure sync.Once

func Get_Main_pure() gopurs_runtime.Value {
	once_Main_pure.Do(func() {
		cache_Main_pure = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1)
	})
	return cache_Main_pure
}

var cache_Main_eqRec gopurs_runtime.Value
var once_Main_eqRec sync.Once

func Get_Main_eqRec() gopurs_runtime.Value {
	once_Main_eqRec.Do(func() {
		cache_Main_eqRec = gopurs_runtime.Apply(Get_Data_Eq_eqRec(), gopurs_runtime.Value{})
	})
	return cache_Main_eqRec
}

var cache_Main_eqRowCons gopurs_runtime.Value
var once_Main_eqRowCons sync.Once

func Get_Main_eqRowCons() gopurs_runtime.Value {
	once_Main_eqRowCons.Do(func() {
		cache_Main_eqRowCons = gopurs_runtime.Apply2(Get_Data_Eq_eqRowCons(), gopurs_runtime.Value{Type: 9, IntVal: 1311326743, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_EqRecord](Get_Data_Eq_eqRowNil()))}, gopurs_runtime.Value{})
	})
	return cache_Main_eqRowCons
}

var cache_Main_eqArray1 gopurs_runtime.Value
var once_Main_eqArray1 sync.Once

func Get_Main_eqArray1() gopurs_runtime.Value {
	once_Main_eqArray1.Do(func() {
		cache_Main_eqArray1 = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl()))
	})
	return cache_Main_eqArray1
}

var cache_Main_eqArray2 gopurs_runtime.Value
var once_Main_eqArray2 sync.Once

func Get_Main_eqArray2() gopurs_runtime.Value {
	once_Main_eqArray2.Do(func() {
		cache_Main_eqArray2 = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl()))
	})
	return cache_Main_eqArray2
}

var cache_Main_eqArray3 gopurs_runtime.Value
var once_Main_eqArray3 sync.Once

func Get_Main_eqArray3() gopurs_runtime.Value {
	once_Main_eqArray3.Do(func() {
		cache_Main_eqArray3 = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_1, "zArrayA"), gopurs_runtime.RecordGet(rb_2, "zArrayA")).IntVal) != (0))
					})
				})
			}))
			_ = __local_var_0_0
			// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
			__local_var_1_2 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_2, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_3, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_0_0, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_2, rb_3).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_1_2
			// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
			__local_var_2_3 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_3, "fa"), gopurs_runtime.RecordGet(rb_4, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_2, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_3, rb_4).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_2_3
			// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
			__local_var_3_4 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_4, "fIgnore"), gopurs_runtime.RecordGet(rb_5, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_4, rb_5).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_3_4
			// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
			__local_var_4_5 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_5, "arrayIgnore"), gopurs_runtime.RecordGet(rb_6, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_4, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_4_5
			// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
			__local_var_1_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_5, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_6, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_5, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6).IntVal) != (0)))
				})
			}))
			_ = __local_var_1_1
			return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(rb_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), gopurs_runtime.RecordGet(ra_2, "nested"), gopurs_runtime.RecordGet(rb_3, "nested")).IntVal) != (0))
				})
			})))
		}()
	})
	return cache_Main_eqArray3
}

var cache_Main_M0 gopurs_runtime.Value
var once_Main_M0 sync.Once

func Get_Main_M0() gopurs_runtime.Value {
	once_Main_M0.Do(func() {
		cache_Main_M0 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Main_M0
}

var cache_Main_M1 gopurs_runtime.Value
var once_Main_M1 sync.Once

func Get_Main_M1() gopurs_runtime.Value {
	once_Main_M1.Do(func() {
		cache_Main_M1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, value0, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(value1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()}))}
			})
		})
	})
	return cache_Main_M1
}

var cache_Main_M2 gopurs_runtime.Value
var once_Main_M2 sync.Once

func Get_Main_M2() gopurs_runtime.Value {
	once_Main_M2.Do(func() {
		cache_Main_M2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2{1, value0.IntVal}))}
		})
	})
	return cache_Main_M2
}

var cache_Main_M3 gopurs_runtime.Value
var once_Main_M3 sync.Once

func Get_Main_M3() gopurs_runtime.Value {
	once_Main_M3.Do(func() {
		cache_Main_M3 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, value0}))}
		})
	})
	return cache_Main_M3
}

var cache_Main_M4 gopurs_runtime.Value
var once_Main_M4 sync.Once

func Get_Main_M4() gopurs_runtime.Value {
	once_Main_M4.Do(func() {
		cache_Main_M4 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, value0}))}
		})
	})
	return cache_Main_M4
}

var cache_Main_M5 gopurs_runtime.Value
var once_Main_M5 sync.Once

func Get_Main_M5() gopurs_runtime.Value {
	once_Main_M5.Do(func() {
		cache_Main_M5 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, value0}))}
		})
	})
	return cache_Main_M5
}

var cache_Main_M6 gopurs_runtime.Value
var once_Main_M6 sync.Once

func Get_Main_M6() gopurs_runtime.Value {
	once_Main_M6.Do(func() {
		cache_Main_M6 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(value4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(value5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(value6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(value7 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, value0.IntVal, value1, func() []int64 {
											arr := *(*[]gopurs_runtime.Value)(value2.UnsafePtr)
											unboxed := make([]int64, len(arr))
											for i, v := range arr {
												unboxed[i] = v.IntVal
											}
											return unboxed
										}(), func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(value3.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}(), value4, value5, value6, value7}))}
									})
								})
							})
						})
					})
				})
			})
		})
	})
	return cache_Main_M6
}

var cache_Main_M7 gopurs_runtime.Value
var once_Main_M7 sync.Once

func Get_Main_M7() gopurs_runtime.Value {
	once_Main_M7.Do(func() {
		cache_Main_M7 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, value0}))}
		})
	})
	return cache_Main_M7
}

var cache_Main_functorM gopurs_runtime.Value
var once_Main_functorM sync.Once

func Get_Main_functorM() gopurs_runtime.Value {
	once_Main_functorM.Do(func() {
		cache_Main_functorM = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_functorM(dictFunctor_0_box)
		})
	})
	return cache_Main_functorM
}

var cache_Main_foldableM gopurs_runtime.Value
var once_Main_foldableM sync.Once

func Get_Main_foldableM() gopurs_runtime.Value {
	once_Main_foldableM.Do(func() {
		cache_Main_foldableM = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_foldableM(dictFoldable_0_box)
		})
	})
	return cache_Main_foldableM
}

var cache_Main_traversableM gopurs_runtime.Value
var once_Main_traversableM sync.Once

func Get_Main_traversableM() gopurs_runtime.Value {
	once_Main_traversableM.Do(func() {
		cache_Main_traversableM = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_traversableM(dictTraversable_0_box)
		})
	})
	return cache_Main_traversableM
}

var cache_Main_traversableM1 gopurs_runtime.Value
var once_Main_traversableM1 sync.Once

func Get_Main_traversableM1() gopurs_runtime.Value {
	once_Main_traversableM1.Do(func() {
		cache_Main_traversableM1 = func() gopurs_runtime.Value {
			// TAST (Let): functorM1_0_0 -> *Constructor_Data_Functor_Functor
			functorM1_0_0 := (&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t1 gopurs_runtime.Value
					{
						if m_1.Type == 9 && m_1.IntVal == 3852365315 {
							__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)}
							goto end_branch_1
						} else {

						}
					}
					{
						if m_1.Type == 9 && m_1.IntVal == 769986722 {
							__t1 = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, gopurs_runtime.Apply(f_0, (*Constructor_Main_M1)(m_1.UnsafePtr).V0), func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap6 := gopurs_runtime.Array((*Constructor_Main_M1)(m_1.UnsafePtr).V1)
									_ = arr_val_arrayMap6
									arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
									_ = arr_go_arrayMap6
									res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
									_ = res_go_arrayMap6
									for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
										res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_0, v_arrayMap6)
									}
									return gopurs_runtime.Array(res_go_arrayMap6)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()}))}
							goto end_branch_1
						} else {

						}
					}
					{
						if m_1.Type == 9 && m_1.IntVal == 2727978561 {
							__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2{1, (*Constructor_Main_M2)(m_1.UnsafePtr).V0}))}
							goto end_branch_1
						} else {

						}
					}
					{
						if m_1.Type == 9 && m_1.IntVal == 1830062304 {
							__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, func() gopurs_runtime.Value {
								arr_val_arrayMap6 := (*Constructor_Main_M3)(m_1.UnsafePtr).V0
								_ = arr_val_arrayMap6
								arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
								_ = arr_go_arrayMap6
								res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
								_ = res_go_arrayMap6
								for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
									res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_0, v_arrayMap6)
								}
								return gopurs_runtime.Array(res_go_arrayMap6)
							}()}))}
							goto end_branch_1
						} else {

						}
					}
					{
						if m_1.Type == 9 && m_1.IntVal == 3190619783 {
							__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4)(m_1.UnsafePtr).V0, "a", gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_1.UnsafePtr).V0, "a")), "fa", func() gopurs_runtime.Value {
								arr_val_arrayMap7 := gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_1.UnsafePtr).V0, "fa")
								_ = arr_val_arrayMap7
								arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
								_ = arr_go_arrayMap7
								res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
								_ = res_go_arrayMap7
								for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
									res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_0, v_arrayMap7)
								}
								return gopurs_runtime.Array(res_go_arrayMap7)
							}(), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap7 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_1.UnsafePtr).V0, "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap7
									arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
									_ = arr_go_arrayMap7
									res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
									_ = res_go_arrayMap7
									for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
										res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_0, v_arrayMap7)
									}
									return gopurs_runtime.Array(res_go_arrayMap7)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))}))}
							goto end_branch_1
						} else {

						}
					}
					{
						if m_1.Type == 9 && m_1.IntVal == 108241190 {
							__t1 = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, func() gopurs_runtime.Value {
								origVal := (*Constructor_Main_M5)(m_1.UnsafePtr).V0
								if origVal.Type != gopurs_runtime.TypeRecord1 {
									return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_1.UnsafePtr).V0, "nested"), "a", gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_1.UnsafePtr).V0, "nested"), "a")), "fa", func() gopurs_runtime.Value {
										arr_val_arrayMap8 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_1.UnsafePtr).V0, "nested"), "fa")
										_ = arr_val_arrayMap8
										arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
										_ = arr_go_arrayMap8
										res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
										_ = res_go_arrayMap8
										for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
											res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(f_0, v_arrayMap8)
										}
										return gopurs_runtime.Array(res_go_arrayMap8)
									}(), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
											arr_val_arrayMap8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_1.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}())
											_ = arr_val_arrayMap8
											arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
											_ = arr_go_arrayMap8
											res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
											_ = res_go_arrayMap8
											for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
												res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(f_0, v_arrayMap8)
											}
											return gopurs_runtime.Array(res_go_arrayMap8)
										}().UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()))})
								}
								clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
								clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_1.UnsafePtr).V0, "nested"), "a", gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_1.UnsafePtr).V0, "nested"), "a")), "fa", func() gopurs_runtime.Value {
									arr_val_arrayMap8 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_1.UnsafePtr).V0, "nested"), "fa")
									_ = arr_val_arrayMap8
									arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
									_ = arr_go_arrayMap8
									res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
									_ = res_go_arrayMap8
									for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
										res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(f_0, v_arrayMap8)
									}
									return gopurs_runtime.Array(res_go_arrayMap8)
								}(), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
										arr_val_arrayMap8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_1.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}())
										_ = arr_val_arrayMap8
										arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
										_ = arr_go_arrayMap8
										res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
										_ = res_go_arrayMap8
										for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
											res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(f_0, v_arrayMap8)
										}
										return gopurs_runtime.Array(res_go_arrayMap8)
									}().UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()))
								return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
							}()}))}
							goto end_branch_1
						} else {

						}
					}
					{
						if m_1.Type == 9 && m_1.IntVal == 2066233029 {
							__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, (*Constructor_Main_M6)(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Constructor_Main_M6)(m_1.UnsafePtr).V1), (*Constructor_Main_M6)(m_1.UnsafePtr).V2, func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap6 := gopurs_runtime.Array((*Constructor_Main_M6)(m_1.UnsafePtr).V3)
									_ = arr_val_arrayMap6
									arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
									_ = arr_go_arrayMap6
									res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
									_ = res_go_arrayMap6
									for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
										res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_0, v_arrayMap6)
									}
									return gopurs_runtime.Array(res_go_arrayMap6)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}(), func() gopurs_runtime.Value {
								arr_val_arrayMap6 := (*Constructor_Main_M6)(m_1.UnsafePtr).V4
								_ = arr_val_arrayMap6
								arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
								_ = arr_go_arrayMap6
								res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
								_ = res_go_arrayMap6
								for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
									res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_0, v_arrayMap6)
								}
								return gopurs_runtime.Array(res_go_arrayMap6)
							}(), (*Constructor_Main_M6)(m_1.UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6)(m_1.UnsafePtr).V6, "a", gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_1.UnsafePtr).V6, "a")), "fa", func() gopurs_runtime.Value {
								arr_val_arrayMap7 := gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_1.UnsafePtr).V6, "fa")
								_ = arr_val_arrayMap7
								arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
								_ = arr_go_arrayMap7
								res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
								_ = res_go_arrayMap7
								for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
									res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_0, v_arrayMap7)
								}
								return gopurs_runtime.Array(res_go_arrayMap7)
							}(), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap7 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_1.UnsafePtr).V6, "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap7
									arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
									_ = arr_go_arrayMap7
									res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
									_ = res_go_arrayMap7
									for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
										res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_0, v_arrayMap7)
									}
									return gopurs_runtime.Array(res_go_arrayMap7)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())), func() gopurs_runtime.Value {
								origVal := (*Constructor_Main_M6)(m_1.UnsafePtr).V7
								if origVal.Type != gopurs_runtime.TypeRecord1 {
									return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_1.UnsafePtr).V7, "nested"), "a", gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_1.UnsafePtr).V7, "nested"), "a")), "fa", func() gopurs_runtime.Value {
										arr_val_arrayMap8 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_1.UnsafePtr).V7, "nested"), "fa")
										_ = arr_val_arrayMap8
										arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
										_ = arr_go_arrayMap8
										res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
										_ = res_go_arrayMap8
										for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
											res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(f_0, v_arrayMap8)
										}
										return gopurs_runtime.Array(res_go_arrayMap8)
									}(), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
											arr_val_arrayMap8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_1.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}())
											_ = arr_val_arrayMap8
											arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
											_ = arr_go_arrayMap8
											res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
											_ = res_go_arrayMap8
											for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
												res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(f_0, v_arrayMap8)
											}
											return gopurs_runtime.Array(res_go_arrayMap8)
										}().UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()))})
								}
								clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
								clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_1.UnsafePtr).V7, "nested"), "a", gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_1.UnsafePtr).V7, "nested"), "a")), "fa", func() gopurs_runtime.Value {
									arr_val_arrayMap8 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_1.UnsafePtr).V7, "nested"), "fa")
									_ = arr_val_arrayMap8
									arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
									_ = arr_go_arrayMap8
									res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
									_ = res_go_arrayMap8
									for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
										res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(f_0, v_arrayMap8)
									}
									return gopurs_runtime.Array(res_go_arrayMap8)
								}(), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
										arr_val_arrayMap8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_1.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}())
										_ = arr_val_arrayMap8
										arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
										_ = arr_go_arrayMap8
										res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
										_ = res_go_arrayMap8
										for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
											res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(f_0, v_arrayMap8)
										}
										return gopurs_runtime.Array(res_go_arrayMap8)
									}().UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()))
								return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
							}()}))}
							goto end_branch_1
						} else {

						}
					}
					{
						if m_1.Type == 9 && m_1.IntVal == 1168316772 {
							__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, func() gopurs_runtime.Value {
								arr_val_arrayMap6 := (*Constructor_Main_M7)(m_1.UnsafePtr).V0
								_ = arr_val_arrayMap6
								arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
								_ = arr_go_arrayMap6
								res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
								_ = res_go_arrayMap6
								for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
									res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.RecordUpdate1(v1_2, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_2, "nested"), "a", gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_2, "nested"), "a")), "fa", func() gopurs_runtime.Value {
											arr_val_arrayMap11 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_2, "nested"), "fa")
											_ = arr_val_arrayMap11
											arr_go_arrayMap11 := (*[]gopurs_runtime.Value)(arr_val_arrayMap11.UnsafePtr)
											_ = arr_go_arrayMap11
											res_go_arrayMap11 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap11))
											_ = res_go_arrayMap11
											for i_arrayMap11, v_arrayMap11 := range *arr_go_arrayMap11 {
												res_go_arrayMap11[i_arrayMap11] = gopurs_runtime.Apply(f_0, v_arrayMap11)
											}
											return gopurs_runtime.Array(res_go_arrayMap11)
										}(), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
												arr_val_arrayMap11 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_2, "nested"), "zArrayA").UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}())
												_ = arr_val_arrayMap11
												arr_go_arrayMap11 := (*[]gopurs_runtime.Value)(arr_val_arrayMap11.UnsafePtr)
												_ = arr_go_arrayMap11
												res_go_arrayMap11 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap11))
												_ = res_go_arrayMap11
												for i_arrayMap11, v_arrayMap11 := range *arr_go_arrayMap11 {
													res_go_arrayMap11[i_arrayMap11] = gopurs_runtime.Apply(f_0, v_arrayMap11)
												}
												return gopurs_runtime.Array(res_go_arrayMap11)
											}().UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}())))
									})), v_arrayMap6)
								}
								return gopurs_runtime.Array(res_go_arrayMap6)
							}()}))}
							goto end_branch_1
						} else {

						}
					}
					{
						__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_1:
					return __t1
				})
			})})
			_ = functorM1_0_0
			// TAST (Let): foldableM1_1_2 -> *Constructor_Data_Foldable_Foldable
			foldableM1_1_2 := (&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): mempty_2_3 -> gopurs_runtime.Value
				mempty_2_3 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
				_ = mempty_2_3
				// TAST (Let): Semigroup0_3_4 -> *Constructor_Data_Semigroup_Semigroup
				Semigroup0_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
				_ = Semigroup0_3_4
				return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t22 gopurs_runtime.Value
						{
							if m_5.Type == 9 && m_5.IntVal == 3852365315 {
								__t22 = mempty_2_3
								goto end_branch_22
							} else {

							}
						}
						{
							if m_5.Type == 9 && m_5.IntVal == 769986722 {
								// TAST (Let): Semigroup0_6_5 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_6_5
								__t22 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply(f_4, (*Constructor_Main_M1)(m_5.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_5.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array((*Constructor_Main_M1)(m_5.UnsafePtr).V1)))
								goto end_branch_22
							} else {

							}
						}
						{
							if m_5.Type == 9 && m_5.IntVal == 2727978561 {
								__t22 = mempty_2_3
								goto end_branch_22
							} else {

							}
						}
						{
							if m_5.Type == 9 && m_5.IntVal == 1830062304 {
								// TAST (Let): Semigroup0_6_6 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_6_6
								__t22 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_6.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), (*Constructor_Main_M3)(m_5.UnsafePtr).V0)
								goto end_branch_22
							} else {

							}
						}
						{
							if m_5.Type == 9 && m_5.IntVal == 3190619783 {
								// TAST (Let): Semigroup0_6_7 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_6_7
								// TAST (Let): Semigroup0_6_8 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_6_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_6_8
								__t22 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_5.UnsafePtr).V0, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_7.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_5.UnsafePtr).V0, "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_8.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_5.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()))))
								goto end_branch_22
							} else {

							}
						}
						{
							if m_5.Type == 9 && m_5.IntVal == 108241190 {
								// TAST (Let): Semigroup0_6_9 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_6_9
								// TAST (Let): Semigroup0_6_10 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_6_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_6_10
								__t22 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_5.UnsafePtr).V0, "nested"), "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_9.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_5.UnsafePtr).V0, "nested"), "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_10.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_5.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()))))
								goto end_branch_22
							} else {

							}
						}
						{
							if m_5.Type == 9 && m_5.IntVal == 2066233029 {
								// TAST (Let): Semigroup0_6_11 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_6_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_6_11
								// TAST (Let): Semigroup0_6_12 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_6_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_6_12
								// TAST (Let): Semigroup0_6_13 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_6_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_6_13
								// TAST (Let): Semigroup0_6_14 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_6_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_6_14
								// TAST (Let): Semigroup0_6_15 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_6_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_6_15
								// TAST (Let): Semigroup0_6_16 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_6_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_6_16
								__t22 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply(f_4, (*Constructor_Main_M6)(m_5.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_11.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array((*Constructor_Main_M6)(m_5.UnsafePtr).V3)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_12.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), (*Constructor_Main_M6)(m_5.UnsafePtr).V4), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V6, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_13.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V6, "fa")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_14.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V7, "nested"), "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_15.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V7, "nested"), "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_16.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()))))))))))
								goto end_branch_22
							} else {

							}
						}
						{
							if m_5.Type == 9 && m_5.IntVal == 1168316772 {
								// TAST (Let): Semigroup0_6_17 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_6_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_6_17
								// TAST (Let): Semigroup0_7_19 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_7_19
								// TAST (Let): __local_var_7_18 -> gopurs_runtime.Value
								__local_var_7_18 := gopurs_runtime.Apply2(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
										// TAST (Let): Semigroup0_10_20 -> *Constructor_Data_Semigroup_Semigroup
										Semigroup0_10_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
										_ = Semigroup0_10_20
										// TAST (Let): Semigroup0_10_21 -> *Constructor_Data_Semigroup_Semigroup
										Semigroup0_10_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
										_ = Semigroup0_10_21
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_19.V0), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_8, "nested"), "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_4.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(acc_12 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_20.V0), gopurs_runtime.Apply(f_4, x_11), acc_12)
											})
										}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_8, "nested"), "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(acc_12 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_21.V0), gopurs_runtime.Apply(f_4, x_11), acc_12)
											})
										}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_8, "nested"), "zArrayA").UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}())))), acc_9)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"))
								_ = __local_var_7_18
								__t22 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_17.V0), gopurs_runtime.Apply(__local_var_7_18, x_8), acc_9)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), (*Constructor_Main_M7)(m_5.UnsafePtr).V0)
								goto end_branch_22
							} else {

							}
						}
						{
							__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_22:
						return __t22
					})
				})
			}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t23 gopurs_runtime.Value
						{
							if m_3.Type == 9 && m_3.IntVal == 3852365315 {
								__t23 = z_2
								goto end_branch_23
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 769986722 {
								__t23 = func() gopurs_runtime.Value {
									arr_val_foldlArray7 := gopurs_runtime.Array((*Constructor_Main_M1)(m_3.UnsafePtr).V1)
									_ = arr_val_foldlArray7
									res_go_foldlArray7 := gopurs_runtime.Apply2(f_1, z_2, (*Constructor_Main_M1)(m_3.UnsafePtr).V0)
									_ = res_go_foldlArray7
									arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
									_ = arr_go_foldlArray7
									for _, v_foldlArray7 := range *arr_go_foldlArray7 {
										res_go_foldlArray7 = gopurs_runtime.Apply2(f_1, res_go_foldlArray7, v_foldlArray7)
									}
									return res_go_foldlArray7
								}()
								goto end_branch_23
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 2727978561 {
								__t23 = z_2
								goto end_branch_23
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 1830062304 {
								__t23 = func() gopurs_runtime.Value {
									arr_val_foldlArray7 := (*Constructor_Main_M3)(m_3.UnsafePtr).V0
									_ = arr_val_foldlArray7
									res_go_foldlArray7 := z_2
									_ = res_go_foldlArray7
									arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
									_ = arr_go_foldlArray7
									for _, v_foldlArray7 := range *arr_go_foldlArray7 {
										res_go_foldlArray7 = gopurs_runtime.Apply2(f_1, res_go_foldlArray7, v_foldlArray7)
									}
									return res_go_foldlArray7
								}()
								goto end_branch_23
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 3190619783 {
								__t23 = func() gopurs_runtime.Value {
									arr_val_foldlArray7 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_foldlArray7
									res_go_foldlArray7 := func() gopurs_runtime.Value {
										arr_val_foldlArray8 := gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "fa")
										_ = arr_val_foldlArray8
										res_go_foldlArray8 := gopurs_runtime.Apply2(f_1, z_2, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "a"))
										_ = res_go_foldlArray8
										arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
										_ = arr_go_foldlArray8
										for _, v_foldlArray8 := range *arr_go_foldlArray8 {
											res_go_foldlArray8 = gopurs_runtime.Apply2(f_1, res_go_foldlArray8, v_foldlArray8)
										}
										return res_go_foldlArray8
									}()
									_ = res_go_foldlArray7
									arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
									_ = arr_go_foldlArray7
									for _, v_foldlArray7 := range *arr_go_foldlArray7 {
										res_go_foldlArray7 = gopurs_runtime.Apply2(f_1, res_go_foldlArray7, v_foldlArray7)
									}
									return res_go_foldlArray7
								}()
								goto end_branch_23
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 108241190 {
								__t23 = func() gopurs_runtime.Value {
									arr_val_foldlArray7 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_foldlArray7
									res_go_foldlArray7 := func() gopurs_runtime.Value {
										arr_val_foldlArray8 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "fa")
										_ = arr_val_foldlArray8
										res_go_foldlArray8 := gopurs_runtime.Apply2(f_1, z_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "a"))
										_ = res_go_foldlArray8
										arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
										_ = arr_go_foldlArray8
										for _, v_foldlArray8 := range *arr_go_foldlArray8 {
											res_go_foldlArray8 = gopurs_runtime.Apply2(f_1, res_go_foldlArray8, v_foldlArray8)
										}
										return res_go_foldlArray8
									}()
									_ = res_go_foldlArray7
									arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
									_ = arr_go_foldlArray7
									for _, v_foldlArray7 := range *arr_go_foldlArray7 {
										res_go_foldlArray7 = gopurs_runtime.Apply2(f_1, res_go_foldlArray7, v_foldlArray7)
									}
									return res_go_foldlArray7
								}()
								goto end_branch_23
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 2066233029 {
								__t23 = func() gopurs_runtime.Value {
									arr_val_foldlArray7 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_foldlArray7
									res_go_foldlArray7 := func() gopurs_runtime.Value {
										arr_val_foldlArray8 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "fa")
										_ = arr_val_foldlArray8
										res_go_foldlArray8 := gopurs_runtime.Apply2(f_1, func() gopurs_runtime.Value {
											arr_val_foldlArray10 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "zArrayA").UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}())
											_ = arr_val_foldlArray10
											res_go_foldlArray10 := func() gopurs_runtime.Value {
												arr_val_foldlArray11 := gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "fa")
												_ = arr_val_foldlArray11
												res_go_foldlArray11 := gopurs_runtime.Apply2(f_1, func() gopurs_runtime.Value {
													arr_val_foldlArray13 := (*Constructor_Main_M6)(m_3.UnsafePtr).V4
													_ = arr_val_foldlArray13
													res_go_foldlArray13 := func() gopurs_runtime.Value {
														arr_val_foldlArray14 := gopurs_runtime.Array((*Constructor_Main_M6)(m_3.UnsafePtr).V3)
														_ = arr_val_foldlArray14
														res_go_foldlArray14 := gopurs_runtime.Apply2(f_1, z_2, (*Constructor_Main_M6)(m_3.UnsafePtr).V1)
														_ = res_go_foldlArray14
														arr_go_foldlArray14 := (*[]gopurs_runtime.Value)(arr_val_foldlArray14.UnsafePtr)
														_ = arr_go_foldlArray14
														for _, v_foldlArray14 := range *arr_go_foldlArray14 {
															res_go_foldlArray14 = gopurs_runtime.Apply2(f_1, res_go_foldlArray14, v_foldlArray14)
														}
														return res_go_foldlArray14
													}()
													_ = res_go_foldlArray13
													arr_go_foldlArray13 := (*[]gopurs_runtime.Value)(arr_val_foldlArray13.UnsafePtr)
													_ = arr_go_foldlArray13
													for _, v_foldlArray13 := range *arr_go_foldlArray13 {
														res_go_foldlArray13 = gopurs_runtime.Apply2(f_1, res_go_foldlArray13, v_foldlArray13)
													}
													return res_go_foldlArray13
												}(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "a"))
												_ = res_go_foldlArray11
												arr_go_foldlArray11 := (*[]gopurs_runtime.Value)(arr_val_foldlArray11.UnsafePtr)
												_ = arr_go_foldlArray11
												for _, v_foldlArray11 := range *arr_go_foldlArray11 {
													res_go_foldlArray11 = gopurs_runtime.Apply2(f_1, res_go_foldlArray11, v_foldlArray11)
												}
												return res_go_foldlArray11
											}()
											_ = res_go_foldlArray10
											arr_go_foldlArray10 := (*[]gopurs_runtime.Value)(arr_val_foldlArray10.UnsafePtr)
											_ = arr_go_foldlArray10
											for _, v_foldlArray10 := range *arr_go_foldlArray10 {
												res_go_foldlArray10 = gopurs_runtime.Apply2(f_1, res_go_foldlArray10, v_foldlArray10)
											}
											return res_go_foldlArray10
										}(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "a"))
										_ = res_go_foldlArray8
										arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
										_ = arr_go_foldlArray8
										for _, v_foldlArray8 := range *arr_go_foldlArray8 {
											res_go_foldlArray8 = gopurs_runtime.Apply2(f_1, res_go_foldlArray8, v_foldlArray8)
										}
										return res_go_foldlArray8
									}()
									_ = res_go_foldlArray7
									arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
									_ = arr_go_foldlArray7
									for _, v_foldlArray7 := range *arr_go_foldlArray7 {
										res_go_foldlArray7 = gopurs_runtime.Apply2(f_1, res_go_foldlArray7, v_foldlArray7)
									}
									return res_go_foldlArray7
								}()
								goto end_branch_23
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 1168316772 {
								__t23 = func() gopurs_runtime.Value {
									arr_val_foldlArray7 := (*Constructor_Main_M7)(m_3.UnsafePtr).V0
									_ = arr_val_foldlArray7
									res_go_foldlArray7 := z_2
									_ = res_go_foldlArray7
									arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
									_ = arr_go_foldlArray7
									for _, v_foldlArray7 := range *arr_go_foldlArray7 {
										res_go_foldlArray7 = gopurs_runtime.Apply2(gopurs_runtime.Apply(Get_Data_Foldable_foldlArray(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
												return func() gopurs_runtime.Value {
													arr_val_foldlArray11 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "zArrayA").UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}())
													_ = arr_val_foldlArray11
													res_go_foldlArray11 := func() gopurs_runtime.Value {
														arr_val_foldlArray12 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "fa")
														_ = arr_val_foldlArray12
														res_go_foldlArray12 := gopurs_runtime.Apply2(f_1, v1_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "a"))
														_ = res_go_foldlArray12
														arr_go_foldlArray12 := (*[]gopurs_runtime.Value)(arr_val_foldlArray12.UnsafePtr)
														_ = arr_go_foldlArray12
														for _, v_foldlArray12 := range *arr_go_foldlArray12 {
															res_go_foldlArray12 = gopurs_runtime.Apply2(f_1, res_go_foldlArray12, v_foldlArray12)
														}
														return res_go_foldlArray12
													}()
													_ = res_go_foldlArray11
													arr_go_foldlArray11 := (*[]gopurs_runtime.Value)(arr_val_foldlArray11.UnsafePtr)
													_ = arr_go_foldlArray11
													for _, v_foldlArray11 := range *arr_go_foldlArray11 {
														res_go_foldlArray11 = gopurs_runtime.Apply2(f_1, res_go_foldlArray11, v_foldlArray11)
													}
													return res_go_foldlArray11
												}()
											})
										})), res_go_foldlArray7, v_foldlArray7)
									}
									return res_go_foldlArray7
								}()
								goto end_branch_23
							} else {

							}
						}
						{
							__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_23:
						return __t23
					})
				})
			}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t25 gopurs_runtime.Value
						{
							if m_3.Type == 9 && m_3.IntVal == 3852365315 {
								__t25 = z_2
								goto end_branch_25
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 769986722 {
								__t25 = gopurs_runtime.Apply2(f_1, (*Constructor_Main_M1)(m_3.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array((*Constructor_Main_M1)(m_3.UnsafePtr).V1)))
								goto end_branch_25
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 2727978561 {
								__t25 = z_2
								goto end_branch_25
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 1830062304 {
								__t25 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, (*Constructor_Main_M3)(m_3.UnsafePtr).V0)
								goto end_branch_25
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 3190619783 {
								__t25 = gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "a"), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())), gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "fa")))
								goto end_branch_25
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 108241190 {
								__t25 = gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "a"), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "fa")))
								goto end_branch_25
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 2066233029 {
								__t25 = gopurs_runtime.Apply2(f_1, (*Constructor_Main_M6)(m_3.UnsafePtr).V1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "a"), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "a"), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "fa"))), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())), gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "fa"))), (*Constructor_Main_M6)(m_3.UnsafePtr).V4), gopurs_runtime.Array((*Constructor_Main_M6)(m_3.UnsafePtr).V3)))
								goto end_branch_25
							} else {

							}
						}
						{
							if m_3.Type == 9 && m_3.IntVal == 1168316772 {
								// TAST (Let): __local_var_4_24 -> gopurs_runtime.Value
								__local_var_4_24 := gopurs_runtime.Apply(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "a"), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, v2_5, gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "zArrayA").UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}())), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "fa")))
									})
								}))
								_ = __local_var_4_24
								__t25 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(__local_var_4_24, a_6, b_5)
									})
								}), z_2, (*Constructor_Main_M7)(m_3.UnsafePtr).V0)
								goto end_branch_25
							} else {

							}
						}
						{
							__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_25:
						return __t25
					})
				})
			})})
			_ = foldableM1_1_2
			return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableM1_1_2)}
			}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorM1_0_0)}
			}), gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_Traversable_traversableArray()))}), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_2))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return x_4
					}), v_3)
				})
			}), gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): Apply0_3_26 -> *Constructor_Control_Apply_Apply
				Apply0_3_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
				_ = Apply0_3_26
				// TAST (Let): Functor0_4_27 -> *Constructor_Data_Functor_Functor
				Functor0_4_27 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
				_ = Functor0_4_27
				return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t51 gopurs_runtime.Value
						{
							if m_6.Type == 9 && m_6.IntVal == 3852365315 {
								__t51 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)})
								goto end_branch_51
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 769986722 {
								// TAST (Let): Apply0_7_28 -> gopurs_runtime.Value
								Apply0_7_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_7_28
								__t51 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_27.V0), gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, v2_7, func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v3_8.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()}))}
									})
								}), gopurs_runtime.Apply(f_5, (*Constructor_Main_M1)(m_6.UnsafePtr).V0)), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_7_28, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_28, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, gopurs_runtime.Array((*Constructor_Main_M1)(m_6.UnsafePtr).V1)))
								goto end_branch_51
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 2727978561 {
								__t51 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2{1, (*Constructor_Main_M2)(m_6.UnsafePtr).V0}))})
								goto end_branch_51
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 1830062304 {
								// TAST (Let): Apply0_7_29 -> gopurs_runtime.Value
								Apply0_7_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_7_29
								__t51 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_27.V0), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, v1_7}))}
								}), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_7_29, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_29, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, (*Constructor_Main_M3)(m_6.UnsafePtr).V0))
								goto end_branch_51
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 3190619783 {
								// TAST (Let): __local_var_7_30 -> gopurs_runtime.Value
								__local_var_7_30 := (*Constructor_Main_M4)(m_6.UnsafePtr).V0
								_ = __local_var_7_30
								// TAST (Let): Apply0_8_31 -> gopurs_runtime.Value
								Apply0_8_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_8_31
								// TAST (Let): Apply0_8_32 -> gopurs_runtime.Value
								Apply0_8_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_8_32
								__t51 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_27.V0), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3(__local_var_7_30, "a", v1_8, "fa", v2_9, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_10.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}()))}))}
										})
									})
								}), gopurs_runtime.Apply(f_5, gopurs_runtime.RecordGet(__local_var_7_30, "a"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_8_31, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_31, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, gopurs_runtime.RecordGet(__local_var_7_30, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_8_32, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_32, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_7_30, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
								goto end_branch_51
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 108241190 {
								// TAST (Let): __local_var_7_33 -> gopurs_runtime.Value
								var __local_var_7_33 gopurs_runtime.Value = (*Constructor_Main_M5)(m_6.UnsafePtr).V0
								// TAST (Let): Apply0_8_34 -> gopurs_runtime.Value
								Apply0_8_34 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_8_34
								// TAST (Let): Apply0_8_35 -> gopurs_runtime.Value
								Apply0_8_35 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_8_35
								__t51 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_27.V0), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, gopurs_runtime.RecordUpdate1(__local_var_7_33, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(__local_var_7_33, "nested"), "a", v1_8, "fa", v2_9, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_10.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}())))}))}
										})
									})
								}), gopurs_runtime.Apply(f_5, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_7_33, "nested"), "a"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_8_34, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_34, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_7_33, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_8_35, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_35, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_7_33, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
								goto end_branch_51
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 2066233029 {
								// TAST (Let): __local_var_7_36 -> gopurs_runtime.Value
								var __local_var_7_36 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Main_M6)(m_6.UnsafePtr).V0)
								// TAST (Let): __local_var_8_37 -> gopurs_runtime.Value
								var __local_var_8_37 gopurs_runtime.Value = func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(m_6.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()
								// TAST (Let): __local_var_9_38 -> gopurs_runtime.Value
								__local_var_9_38 := (*Constructor_Main_M6)(m_6.UnsafePtr).V5
								_ = __local_var_9_38
								// TAST (Let): __local_var_10_39 -> gopurs_runtime.Value
								__local_var_10_39 := (*Constructor_Main_M6)(m_6.UnsafePtr).V6
								_ = __local_var_10_39
								// TAST (Let): __local_var_11_40 -> gopurs_runtime.Value
								var __local_var_11_40 gopurs_runtime.Value = (*Constructor_Main_M6)(m_6.UnsafePtr).V7
								// TAST (Let): Apply0_12_41 -> gopurs_runtime.Value
								Apply0_12_41 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_12_41
								// TAST (Let): Apply0_12_42 -> gopurs_runtime.Value
								Apply0_12_42 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_12_42
								// TAST (Let): Apply0_12_43 -> gopurs_runtime.Value
								Apply0_12_43 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_12_43
								// TAST (Let): Apply0_12_44 -> gopurs_runtime.Value
								Apply0_12_44 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_12_44
								// TAST (Let): Apply0_12_45 -> gopurs_runtime.Value
								Apply0_12_45 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_12_45
								// TAST (Let): Apply0_12_46 -> gopurs_runtime.Value
								Apply0_12_46 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_12_46
								__t51 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_27.V0), gopurs_runtime.Func(func(v8_12 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v9_13 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v10_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v11_15 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v12_16 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v13_17 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Func(func(v14_18 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(v15_19 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(v16_20 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, __local_var_7_36.IntVal, v8_12, func() []int64 {
																		arr := *(*[]gopurs_runtime.Value)(__local_var_8_37.UnsafePtr)
																		unboxed := make([]int64, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v.IntVal
																		}
																		return unboxed
																	}(), func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v9_13.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}(), v10_14, __local_var_9_38, gopurs_runtime.RecordUpdate3(__local_var_10_39, "a", v11_15, "fa", v12_16, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v13_17.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}())), gopurs_runtime.RecordUpdate1(__local_var_11_40, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(__local_var_11_40, "nested"), "a", v14_18, "fa", v15_19, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v16_20.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}())))}))}
																})
															})
														})
													})
												})
											})
										})
									})
								}), gopurs_runtime.Apply(f_5, (*Constructor_Main_M6)(m_6.UnsafePtr).V1)), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_12_41, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_12_41, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, gopurs_runtime.Array((*Constructor_Main_M6)(m_6.UnsafePtr).V3))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_12_42, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_12_42, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, (*Constructor_Main_M6)(m_6.UnsafePtr).V4)), gopurs_runtime.Apply(f_5, gopurs_runtime.RecordGet(__local_var_10_39, "a"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_12_43, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_12_43, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, gopurs_runtime.RecordGet(__local_var_10_39, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_12_44, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_12_44, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_10_39, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()))), gopurs_runtime.Apply(f_5, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_11_40, "nested"), "a"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_12_45, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_12_45, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_11_40, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_12_46, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_12_46, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_11_40, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
								goto end_branch_51
							} else {

							}
						}
						{
							if m_6.Type == 9 && m_6.IntVal == 1168316772 {
								// TAST (Let): Apply0_7_47 -> gopurs_runtime.Value
								Apply0_7_47 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_7_47
								// TAST (Let): Apply0_8_48 -> gopurs_runtime.Value
								Apply0_8_48 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_8_48
								__t51 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_27.V0), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, v1_7}))}
								}), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_7_47, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_47, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply5(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_8_48, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_48, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
									// TAST (Let): Apply0_10_49 -> gopurs_runtime.Value
									Apply0_10_49 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
									_ = Apply0_10_49
									// TAST (Let): Apply0_10_50 -> gopurs_runtime.Value
									Apply0_10_50 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{})
									_ = Apply0_10_50
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_26.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_27.V0), gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.RecordUpdate1(v1_9, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_9, "nested"), "a", v2_10, "fa", v3_11, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v4_12.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}())))
											})
										})
									}), gopurs_runtime.Apply(f_5, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_9, "nested"), "a"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_10_49, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_10_49, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_9, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_10_50, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_10_50, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_2, "pure"), Get_Data_Semigroup_concatArray(), f_5, gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_9, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())))
								})), (*Constructor_Main_M7)(m_6.UnsafePtr).V0))
								goto end_branch_51
							} else {

							}
						}
						{
							__t51 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
						}
					end_branch_51:
						return __t51
					})
				})
			})}))}
		}()
	})
	return cache_Main_traversableM1
}

var cache_Main_eqM gopurs_runtime.Value
var once_Main_eqM sync.Once

func Get_Main_eqM() gopurs_runtime.Value {
	once_Main_eqM.Do(func() {
		cache_Main_eqM = gopurs_runtime.Func4(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value, dictEq2_2_box gopurs_runtime.Value, dictEq3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqM(dictEq1_0_box, dictEq_1_box, dictEq2_2_box, dictEq3_3_box)
		})
	})
	return cache_Main_eqM
}

var cache_Main_eqArray4 gopurs_runtime.Value
var once_Main_eqArray4 sync.Once

func Get_Main_eqArray4() gopurs_runtime.Value {
	once_Main_eqArray4.Do(func() {
		cache_Main_eqArray4 = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_1, "zArrayA"), gopurs_runtime.RecordGet(rb_2, "zArrayA")).IntVal) != (0))
					})
				})
			}))
			_ = __local_var_0_0
			// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
			__local_var_1_2 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_2, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_3, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_0_0, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_2, rb_3).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_1_2
			// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
			__local_var_2_3 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_3, "fa"), gopurs_runtime.RecordGet(rb_4, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_2, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_3, rb_4).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_2_3
			// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
			__local_var_3_4 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_4, "fIgnore"), gopurs_runtime.RecordGet(rb_5, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_4, rb_5).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_3_4
			// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
			__local_var_4_5 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_5, "arrayIgnore"), gopurs_runtime.RecordGet(rb_6, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_4, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_4_5
			// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
			__local_var_1_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_5, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_6, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_5, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6).IntVal) != (0)))
				})
			}))
			_ = __local_var_1_1
			// TAST (Let): __local_var_2_6 -> gopurs_runtime.Value
			__local_var_2_6 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(rb_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), gopurs_runtime.RecordGet(ra_2, "nested"), gopurs_runtime.RecordGet(rb_3, "nested")).IntVal) != (0))
				})
			})))
			_ = __local_var_2_6
			// TAST (Let): eqArray5_3_7 -> *Constructor_Data_Eq_Eq
			eqArray5_3_7 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
			_ = eqArray5_3_7
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Box((&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t8 bool
					{
						if x_4.Type == 9 && x_4.IntVal == 3852365315 {
							__t8 = (y_5.Type == 9 && y_5.IntVal == 3852365315)
							goto end_branch_8
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 769986722 {
							__t8 = (y_5.Type == 9 && y_5.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_4.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_5.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_3_7.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_4.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_5.UnsafePtr).V1)).IntVal) != (0)))
							goto end_branch_8
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 2727978561 {
							__t8 = (y_5.Type == 9 && y_5.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_4.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_5.UnsafePtr).V0))
							goto end_branch_8
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 1830062304 {
							__t8 = (y_5.Type == 9 && y_5.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_4.UnsafePtr).V0, (*Constructor_Main_M3)(y_5.UnsafePtr).V0).IntVal) != (0))
							goto end_branch_8
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 3190619783 {
							__t8 = (y_5.Type == 9 && y_5.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_4.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_5.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_4.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_5.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_4.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_5.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_4.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_5.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_4.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_5.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_3_7.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_4.UnsafePtr).V0, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_5.UnsafePtr).V0, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())).IntVal) != (0)))
							goto end_branch_8
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 108241190 {
							__t8 = (y_5.Type == 9 && y_5.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_4.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_5.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_4.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_5.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_4.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_5.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_4.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_5.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_4.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_5.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_3_7.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_4.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_5.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())).IntVal) != (0)))
							goto end_branch_8
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 2066233029 {
							__t8 = (y_5.Type == 9 && y_5.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_4.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_5.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_4.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_5.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
								arr := (*Constructor_Main_M6)(x_4.UnsafePtr).V2
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}(), func() gopurs_runtime.Value {
								arr := (*Constructor_Main_M6)(y_5.UnsafePtr).V2
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_3_7.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_4.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_5.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_4.UnsafePtr).V4, (*Constructor_Main_M6)(y_5.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_4.UnsafePtr).V5, (*Constructor_Main_M6)(y_5.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_3_7.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V6, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V6, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_3_7.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())).IntVal) != (0))))
							goto end_branch_8
						} else {

						}
					}
					{
						__t8 = (x_4.Type == 9 && x_4.IntVal == 1168316772) && ((y_5.Type == 9 && y_5.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_2_6, "eq"), (*Constructor_Main_M7)(x_4.UnsafePtr).V0, (*Constructor_Main_M7)(y_5.UnsafePtr).V0).IntVal) != (0)))
					}
				end_branch_8:
					return gopurs_runtime.Bool(__t8)
				})
			})}).V0)))))}
		}()
	})
	return cache_Main_eqArray4
}

var cache_Main_traverseStr gopurs_runtime.Value
var once_Main_traverseStr sync.Once

func Get_Main_traverseStr() gopurs_runtime.Value {
	once_Main_traverseStr.Do(func() {
		cache_Main_traverseStr = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_traverseStr(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box))
		})
	})
	return cache_Main_traverseStr
}

var cache_Main_sequenceStr gopurs_runtime.Value
var once_Main_sequenceStr sync.Once

func Get_Main_sequenceStr() gopurs_runtime.Value {
	once_Main_sequenceStr.Do(func() {
		cache_Main_sequenceStr = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_sequenceStr(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box))
		})
	})
	return cache_Main_sequenceStr
}

var cache_Main_recordValue_prime_ gopurs_runtime.Value
var once_Main_recordValue_prime_ sync.Once

func Get_Main_recordValue_prime_() gopurs_runtime.Value {
	once_Main_recordValue_prime_.Do(func() {
		cache_Main_recordValue_prime_ = gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{func() gopurs_runtime.Value {
			arr := []string{"a"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := []int64{2, 3}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := []int64{4}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := [][]string{[]string{"b"}}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = func() gopurs_runtime.Value {
					arr := v
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()
			}
			return gopurs_runtime.Array(boxed)
		}(), gopurs_runtime.Int(1), func() gopurs_runtime.Value {
			arr := [][]string{[]string{"c"}}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = func() gopurs_runtime.Value {
					arr := v
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()
			}
			return gopurs_runtime.Array(boxed)
		}()})
	})
	return cache_Main_recordValue_prime_
}

var cache_Main_recordValue gopurs_runtime.Value
var once_Main_recordValue sync.Once

func Get_Main_recordValue() gopurs_runtime.Value {
	once_Main_recordValue.Do(func() {
		cache_Main_recordValue = gopurs_runtime.RecordDict([]string{"a", "arrayIgnore", "fIgnore", "fa", "ignore", "zArrayA"}, []gopurs_runtime.Value{gopurs_runtime.Str("a"), func() gopurs_runtime.Value {
			arr := []int64{2, 3}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := []int64{4}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() gopurs_runtime.Value {
			arr := []string{"b"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), gopurs_runtime.Int(1), func() gopurs_runtime.Value {
			arr := []string{"c"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()})
	})
	return cache_Main_recordValue
}

var cache_Main_m7_prime_ gopurs_runtime.Value
var once_Main_m7_prime_ sync.Once

func Get_Main_m7_prime_() gopurs_runtime.Value {
	once_Main_m7_prime_.Do(func() {
		cache_Main_m7_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, func() gopurs_runtime.Value {
			arr := [][]gopurs_runtime.Value{[]gopurs_runtime.Value{gopurs_runtime.RecordDict1("nested", Get_Main_recordValue_prime_())}}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = func() gopurs_runtime.Value {
					arr := v
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = v
					}
					return gopurs_runtime.Array(boxed)
				}()
			}
			return gopurs_runtime.Array(boxed)
		}()}))}
	})
	return cache_Main_m7_prime_
}

var cache_Main_m7 gopurs_runtime.Value
var once_Main_m7 sync.Once

func Get_Main_m7() gopurs_runtime.Value {
	once_Main_m7.Do(func() {
		cache_Main_m7 = gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, func() gopurs_runtime.Value {
			arr := [][]gopurs_runtime.Value{[]gopurs_runtime.Value{gopurs_runtime.RecordDict1("nested", Get_Main_recordValue())}}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = func() gopurs_runtime.Value {
					arr := v
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = v
					}
					return gopurs_runtime.Array(boxed)
				}()
			}
			return gopurs_runtime.Array(boxed)
		}()}))}
	})
	return cache_Main_m7
}

var cache_Main_m6_prime_ gopurs_runtime.Value
var once_Main_m6_prime_ sync.Once

func Get_Main_m6_prime_() gopurs_runtime.Value {
	once_Main_m6_prime_.Do(func() {
		cache_Main_m6_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, 1, func() gopurs_runtime.Value {
			arr := []string{"a"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() []int64 {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
			unboxed := make([]int64, len(arr))
			for i, v := range arr {
				unboxed[i] = v.IntVal
			}
			return unboxed
		}(), func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := [][]string{[]string{"b"}}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						arr := v
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}()
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}(), func() gopurs_runtime.Value {
			arr := [][]string{[]string{"c"}}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = func() gopurs_runtime.Value {
					arr := v
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()
			}
			return gopurs_runtime.Array(boxed)
		}(), gopurs_runtime.Array([]gopurs_runtime.Value{}), Get_Main_recordValue_prime_(), gopurs_runtime.RecordDict1("nested", Get_Main_recordValue_prime_())}))}
	})
	return cache_Main_m6_prime_
}

var cache_Main_m6 gopurs_runtime.Value
var once_Main_m6 sync.Once

func Get_Main_m6() gopurs_runtime.Value {
	once_Main_m6.Do(func() {
		cache_Main_m6 = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, 1, gopurs_runtime.Str("a"), func() []int64 {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
			unboxed := make([]int64, len(arr))
			for i, v := range arr {
				unboxed[i] = v.IntVal
			}
			return unboxed
		}(), func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := []string{"b"}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}(), func() gopurs_runtime.Value {
			arr := []string{"c"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), gopurs_runtime.Array([]gopurs_runtime.Value{}), Get_Main_recordValue(), gopurs_runtime.RecordDict1("nested", Get_Main_recordValue())}))}
	})
	return cache_Main_m6
}

var cache_Main_m5_prime_ gopurs_runtime.Value
var once_Main_m5_prime_ sync.Once

func Get_Main_m5_prime_() gopurs_runtime.Value {
	once_Main_m5_prime_.Do(func() {
		cache_Main_m5_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, gopurs_runtime.RecordDict1("nested", Get_Main_recordValue_prime_())}))}
	})
	return cache_Main_m5_prime_
}

var cache_Main_m5 gopurs_runtime.Value
var once_Main_m5 sync.Once

func Get_Main_m5() gopurs_runtime.Value {
	once_Main_m5.Do(func() {
		cache_Main_m5 = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, gopurs_runtime.RecordDict1("nested", Get_Main_recordValue())}))}
	})
	return cache_Main_m5
}

var cache_Main_m4_prime_ gopurs_runtime.Value
var once_Main_m4_prime_ sync.Once

func Get_Main_m4_prime_() gopurs_runtime.Value {
	once_Main_m4_prime_.Do(func() {
		cache_Main_m4_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, Get_Main_recordValue_prime_()}))}
	})
	return cache_Main_m4_prime_
}

var cache_Main_m4 gopurs_runtime.Value
var once_Main_m4 sync.Once

func Get_Main_m4() gopurs_runtime.Value {
	once_Main_m4.Do(func() {
		cache_Main_m4 = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, Get_Main_recordValue()}))}
	})
	return cache_Main_m4
}

var cache_Main_m3_prime_ gopurs_runtime.Value
var once_Main_m3_prime_ sync.Once

func Get_Main_m3_prime_() gopurs_runtime.Value {
	once_Main_m3_prime_.Do(func() {
		cache_Main_m3_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, func() gopurs_runtime.Value {
			arr := [][]string{[]string{"a"}, []string{"b"}, []string{"c"}}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = func() gopurs_runtime.Value {
					arr := v
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}()
			}
			return gopurs_runtime.Array(boxed)
		}()}))}
	})
	return cache_Main_m3_prime_
}

var cache_Main_m3 gopurs_runtime.Value
var once_Main_m3 sync.Once

func Get_Main_m3() gopurs_runtime.Value {
	once_Main_m3.Do(func() {
		cache_Main_m3 = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, func() gopurs_runtime.Value {
			arr := []string{"a", "b", "c"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()}))}
	})
	return cache_Main_m3
}

var cache_Main_m2_prime_ gopurs_runtime.Value
var once_Main_m2_prime_ sync.Once

func Get_Main_m2_prime_() gopurs_runtime.Value {
	once_Main_m2_prime_.Do(func() {
		cache_Main_m2_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2{1, 0}))}
	})
	return cache_Main_m2_prime_
}

var cache_Main_m2 gopurs_runtime.Value
var once_Main_m2 sync.Once

func Get_Main_m2() gopurs_runtime.Value {
	once_Main_m2.Do(func() {
		cache_Main_m2 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2{1, 0}))}
	})
	return cache_Main_m2
}

var cache_Main_m1_prime_ gopurs_runtime.Value
var once_Main_m1_prime_ sync.Once

func Get_Main_m1_prime_() gopurs_runtime.Value {
	once_Main_m1_prime_.Do(func() {
		cache_Main_m1_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, func() gopurs_runtime.Value {
			arr := []string{"a"}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := [][]string{[]string{"b"}, []string{"c"}}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = func() gopurs_runtime.Value {
						arr := v
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}()
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}()}))}
	})
	return cache_Main_m1_prime_
}

var cache_Main_m1 gopurs_runtime.Value
var once_Main_m1 sync.Once

func Get_Main_m1() gopurs_runtime.Value {
	once_Main_m1.Do(func() {
		cache_Main_m1 = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, gopurs_runtime.Str("a"), func() []gopurs_runtime.Value {
			arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
				arr := []string{"b", "c"}
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Str(v)
				}
				return gopurs_runtime.Array(boxed)
			}().UnsafePtr)
			unboxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				unboxed[i] = v
			}
			return unboxed
		}()}))}
	})
	return cache_Main_m1
}

var cache_Main_m0_prime_ gopurs_runtime.Value
var once_Main_m0_prime_ sync.Once

func Get_Main_m0_prime_() gopurs_runtime.Value {
	once_Main_m0_prime_.Do(func() {
		cache_Main_m0_prime_ = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Main_m0_prime_
}

var cache_Main_m0 gopurs_runtime.Value
var once_Main_m0 sync.Once

func Get_Main_m0() gopurs_runtime.Value {
	once_Main_m0.Do(func() {
		cache_Main_m0 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Main_m0
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_1 -> gopurs_runtime.Value
			__local_var_0_1 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_1 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_1, "zArrayA"), gopurs_runtime.RecordGet(rb_2, "zArrayA")).IntVal) != (0))
					})
				})
			}))
			_ = __local_var_0_1
			// TAST (Let): __local_var_1_3 -> gopurs_runtime.Value
			__local_var_1_3 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_2, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_3, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_0_1, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_2, rb_3).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_1_3
			// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
			__local_var_2_4 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_3, "fa"), gopurs_runtime.RecordGet(rb_4, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_3, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_3, rb_4).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_2_4
			// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
			__local_var_3_5 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_4, "fIgnore"), gopurs_runtime.RecordGet(rb_5, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_4, rb_5).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_3_5
			// TAST (Let): __local_var_4_6 -> gopurs_runtime.Value
			__local_var_4_6 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_5, "arrayIgnore"), gopurs_runtime.RecordGet(rb_6, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_5, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6).IntVal) != (0)))
					})
				})
			}))
			_ = __local_var_4_6
			// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
			__local_var_1_2 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_5, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_6, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_6, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6).IntVal) != (0)))
				})
			}))
			_ = __local_var_1_2
			// TAST (Let): __local_var_2_7 -> gopurs_runtime.Value
			__local_var_2_7 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(rb_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "eq"), gopurs_runtime.RecordGet(ra_2, "nested"), gopurs_runtime.RecordGet(rb_3, "nested")).IntVal) != (0))
				})
			})))
			_ = __local_var_2_7
			// TAST (Let): eqArray5_3_8 -> *Constructor_Data_Eq_Eq
			eqArray5_3_8 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
			_ = eqArray5_3_8
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m0"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t9 bool
					{
						if x_4.Type == 9 && x_4.IntVal == 3852365315 {
							__t9 = (y_5.Type == 9 && y_5.IntVal == 3852365315)
							goto end_branch_9
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 769986722 {
							__t9 = (y_5.Type == 9 && y_5.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_4.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_5.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_3_8.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_4.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_5.UnsafePtr).V1)).IntVal) != (0)))
							goto end_branch_9
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 2727978561 {
							__t9 = (y_5.Type == 9 && y_5.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_4.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_5.UnsafePtr).V0))
							goto end_branch_9
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 1830062304 {
							__t9 = (y_5.Type == 9 && y_5.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_4.UnsafePtr).V0, (*Constructor_Main_M3)(y_5.UnsafePtr).V0).IntVal) != (0))
							goto end_branch_9
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 3190619783 {
							__t9 = (y_5.Type == 9 && y_5.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_4.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_5.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_4.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_5.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_4.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_5.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_4.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_5.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_4.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_5.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_3_8.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_4.UnsafePtr).V0, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_5.UnsafePtr).V0, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())).IntVal) != (0)))
							goto end_branch_9
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 108241190 {
							__t9 = (y_5.Type == 9 && y_5.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_4.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_5.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_4.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_5.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_4.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_5.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_4.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_5.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_4.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_5.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_3_8.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_4.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_5.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())).IntVal) != (0)))
							goto end_branch_9
						} else {

						}
					}
					{
						if x_4.Type == 9 && x_4.IntVal == 2066233029 {
							__t9 = (y_5.Type == 9 && y_5.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_4.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_5.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_4.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_5.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
								arr := (*Constructor_Main_M6)(x_4.UnsafePtr).V2
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}(), func() gopurs_runtime.Value {
								arr := (*Constructor_Main_M6)(y_5.UnsafePtr).V2
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_3_8.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_4.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_5.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_4.UnsafePtr).V4, (*Constructor_Main_M6)(y_5.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_4.UnsafePtr).V5, (*Constructor_Main_M6)(y_5.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_3_8.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V6, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V6, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								arr := func() []int64 {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
							}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_3_8.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_4.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_5.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())).IntVal) != (0))))
							goto end_branch_9
						} else {

						}
					}
					{
						__t9 = (x_4.Type == 9 && x_4.IntVal == 1168316772) && ((y_5.Type == 9 && y_5.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_2_7, "eq"), (*Constructor_Main_M7)(x_4.UnsafePtr).V0, (*Constructor_Main_M7)(y_5.UnsafePtr).V0).IntVal) != (0)))
					}
				end_branch_9:
					return gopurs_runtime.Bool(__t9)
				})
			}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := []*Constructor_Main_M0{nil}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
					}
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := []*Constructor_Main_M0{nil}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
					}
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}())).IntVal) != (0)))
			_ = __local_var_0_0
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				_dollar___unused_1_10 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
				_ = _dollar___unused_1_10
				// TAST (Let): __local_var_2_12 -> gopurs_runtime.Value
				__local_var_2_12 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_3, "zArrayA"), gopurs_runtime.RecordGet(rb_4, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_2_12
				// TAST (Let): __local_var_3_14 -> gopurs_runtime.Value
				__local_var_3_14 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_5 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_4, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_5, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_12, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_4, rb_5).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_3_14
				// TAST (Let): __local_var_4_15 -> gopurs_runtime.Value
				__local_var_4_15 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_5, "fa"), gopurs_runtime.RecordGet(rb_6, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_14, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_4_15
				// TAST (Let): __local_var_5_16 -> gopurs_runtime.Value
				__local_var_5_16 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_6, "fIgnore"), gopurs_runtime.RecordGet(rb_7, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_15, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_6, rb_7).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_5_16
				// TAST (Let): __local_var_6_17 -> gopurs_runtime.Value
				__local_var_6_17 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_7, "arrayIgnore"), gopurs_runtime.RecordGet(rb_8, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_16, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_6_17
				// TAST (Let): __local_var_3_13 -> gopurs_runtime.Value
				__local_var_3_13 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_7, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_8, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_6_17, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8).IntVal) != (0)))
					})
				}))
				_ = __local_var_3_13
				// TAST (Let): __local_var_4_18 -> gopurs_runtime.Value
				__local_var_4_18 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_13, "eq"), gopurs_runtime.RecordGet(ra_4, "nested"), gopurs_runtime.RecordGet(rb_5, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_4_18
				// TAST (Let): eqArray5_5_19 -> *Constructor_Data_Eq_Eq
				eqArray5_5_19 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_5_19
				var __t29 gopurs_runtime.Value
				{
					var __t_tag_21 gopurs_runtime.Value = Get_Main_m1()
					if __t_tag_21.Type == 9 && __t_tag_21.IntVal == 3852365315 {
						__t29 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M0{nil}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_29
					} else {

					}
				}
				{
					var __t_tag_22 gopurs_runtime.Value = Get_Main_m1()
					if __t_tag_22.Type == 9 && __t_tag_22.IntVal == 769986722 {
						__t29 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Main_M1)(Get_Main_m1().UnsafePtr).V0}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap8
							arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
							_ = arr_go_arrayMap8
							res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
							_ = res_go_arrayMap8
							for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
								res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, v2_3, func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v3_4.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()}))}
									})
								}), v_arrayMap8)
							}
							return gopurs_runtime.Array(res_go_arrayMap8)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array((*Constructor_Main_M1)(Get_Main_m1().UnsafePtr).V1)))
						goto end_branch_29
					} else {

					}
				}
				{
					var __t_tag_23 gopurs_runtime.Value = Get_Main_m1()
					if __t_tag_23.Type == 9 && __t_tag_23.IntVal == 2727978561 {
						__t29 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, (*Constructor_Main_M2)(Get_Main_m1().UnsafePtr).V0})}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_29
					} else {

					}
				}
				{
					var __t_tag_24 gopurs_runtime.Value = Get_Main_m1()
					if __t_tag_24.Type == 9 && __t_tag_24.IntVal == 1830062304 {
						__t29 = func() gopurs_runtime.Value {
							arr_val_arrayMap7 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), (*Constructor_Main_M3)(Get_Main_m1().UnsafePtr).V0)
							_ = arr_val_arrayMap7
							arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
							_ = arr_go_arrayMap7
							res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
							_ = res_go_arrayMap7
							for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
								res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, v1_3}))}
								}), v_arrayMap7)
							}
							return gopurs_runtime.Array(res_go_arrayMap7)
						}()
						goto end_branch_29
					} else {

					}
				}
				{
					var __t_tag_25 gopurs_runtime.Value = Get_Main_m1()
					if __t_tag_25.Type == 9 && __t_tag_25.IntVal == 3190619783 {
						__t29 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap9 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m1().UnsafePtr).V0, "a")}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap9
							arr_go_arrayMap9 := (*[]gopurs_runtime.Value)(arr_val_arrayMap9.UnsafePtr)
							_ = arr_go_arrayMap9
							res_go_arrayMap9 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap9))
							_ = res_go_arrayMap9
							for i_arrayMap9, v_arrayMap9 := range *arr_go_arrayMap9 {
								res_go_arrayMap9[i_arrayMap9] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_5 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4)(Get_Main_m1().UnsafePtr).V0, "a", v1_3, "fa", v2_4, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_5.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}()))}))}
										})
									})
								}), v_arrayMap9)
							}
							return gopurs_runtime.Array(res_go_arrayMap9)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m1().UnsafePtr).V0, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m1().UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_29
					} else {

					}
				}
				{
					var __t_tag_26 gopurs_runtime.Value = Get_Main_m1()
					if __t_tag_26.Type == 9 && __t_tag_26.IntVal == 108241190 {
						__t29 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap9 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m1().UnsafePtr).V0, "nested"), "a")}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap9
							arr_go_arrayMap9 := (*[]gopurs_runtime.Value)(arr_val_arrayMap9.UnsafePtr)
							_ = arr_go_arrayMap9
							res_go_arrayMap9 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap9))
							_ = res_go_arrayMap9
							for i_arrayMap9, v_arrayMap9 := range *arr_go_arrayMap9 {
								res_go_arrayMap9[i_arrayMap9] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_5 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, func() gopurs_runtime.Value {
												origVal := (*Constructor_Main_M5)(Get_Main_m1().UnsafePtr).V0
												if origVal.Type != gopurs_runtime.TypeRecord1 {
													return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m1().UnsafePtr).V0, "nested"), "a", v1_3, "fa", v2_4, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v3_5.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}()))})
												}
												clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
												clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m1().UnsafePtr).V0, "nested"), "a", v1_3, "fa", v2_4, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v3_5.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}()))
												return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
											}()}))}
										})
									})
								}), v_arrayMap9)
							}
							return gopurs_runtime.Array(res_go_arrayMap9)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m1().UnsafePtr).V0, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m1().UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_29
					} else {

					}
				}
				{
					var __t_tag_27 gopurs_runtime.Value = Get_Main_m1()
					if __t_tag_27.Type == 9 && __t_tag_27.IntVal == 2066233029 {
						__t29 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap15 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V1}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap15
							arr_go_arrayMap15 := (*[]gopurs_runtime.Value)(arr_val_arrayMap15.UnsafePtr)
							_ = arr_go_arrayMap15
							res_go_arrayMap15 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap15))
							_ = res_go_arrayMap15
							for i_arrayMap15, v_arrayMap15 := range *arr_go_arrayMap15 {
								res_go_arrayMap15[i_arrayMap15] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v8_3 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v9_4 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v10_5 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v11_6 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v12_7 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v13_8 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Func(func(v14_9 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(v15_10 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(v16_11 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, (*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V0, v8_3, (*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V2, func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v9_4.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}(), v10_5, (*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V6, "a", v11_6, "fa", v12_7, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v13_8.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}())), func() gopurs_runtime.Value {
																		origVal := (*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V7
																		if origVal.Type != gopurs_runtime.TypeRecord1 {
																			return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V7, "nested"), "a", v14_9, "fa", v15_10, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																				arr := *(*[]gopurs_runtime.Value)(v16_11.UnsafePtr)
																				unboxed := make([]gopurs_runtime.Value, len(arr))
																				for i, v := range arr {
																					unboxed[i] = v
																				}
																				return unboxed
																			}()))})
																		}
																		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
																		clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V7, "nested"), "a", v14_9, "fa", v15_10, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																			arr := *(*[]gopurs_runtime.Value)(v16_11.UnsafePtr)
																			unboxed := make([]gopurs_runtime.Value, len(arr))
																			for i, v := range arr {
																				unboxed[i] = v
																			}
																			return unboxed
																		}()))
																		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
																	}()}))}
																})
															})
														})
													})
												})
											})
										})
									})
								}), v_arrayMap15)
							}
							return gopurs_runtime.Array(res_go_arrayMap15)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array((*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V3))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), (*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V4)), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V6, "a")}).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V6, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V7, "nested"), "a")}).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V7, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1().UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_29
					} else {

					}
				}
				{
					var __t_tag_28 gopurs_runtime.Value = Get_Main_m1()
					if __t_tag_28.Type == 9 && __t_tag_28.IntVal == 1168316772 {
						__t29 = func() gopurs_runtime.Value {
							arr_val_arrayMap7 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply5(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
									arr_val_arrayMap13 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "a")}).UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap13
									arr_go_arrayMap13 := (*[]gopurs_runtime.Value)(arr_val_arrayMap13.UnsafePtr)
									_ = arr_go_arrayMap13
									res_go_arrayMap13 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap13))
									_ = res_go_arrayMap13
									for i_arrayMap13, v_arrayMap13 := range *arr_go_arrayMap13 {
										res_go_arrayMap13[i_arrayMap13] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v3_5 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v4_6 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.RecordUpdate1(v1_3, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_3, "nested"), "a", v2_4, "fa", v3_5, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v4_6.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}())))
												})
											})
										}), v_arrayMap13)
									}
									return gopurs_runtime.Array(res_go_arrayMap13)
								}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
							})), (*Constructor_Main_M7)(Get_Main_m1().UnsafePtr).V0)
							_ = arr_val_arrayMap7
							arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
							_ = arr_go_arrayMap7
							res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
							_ = res_go_arrayMap7
							for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
								res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, v1_3}))}
								}), v_arrayMap7)
							}
							return gopurs_runtime.Array(res_go_arrayMap7)
						}()
						goto end_branch_29
					} else {

					}
				}
				{
					__t29 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_29:
				_dollar___unused_2_11 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m1"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_7 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t20 bool
						{
							if x_6.Type == 9 && x_6.IntVal == 3852365315 {
								__t20 = (y_7.Type == 9 && y_7.IntVal == 3852365315)
								goto end_branch_20
							} else {

							}
						}
						{
							if x_6.Type == 9 && x_6.IntVal == 769986722 {
								__t20 = (y_7.Type == 9 && y_7.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_6.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_7.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_5_19.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_6.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_7.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_20
							} else {

							}
						}
						{
							if x_6.Type == 9 && x_6.IntVal == 2727978561 {
								__t20 = (y_7.Type == 9 && y_7.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_6.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_7.UnsafePtr).V0))
								goto end_branch_20
							} else {

							}
						}
						{
							if x_6.Type == 9 && x_6.IntVal == 1830062304 {
								__t20 = (y_7.Type == 9 && y_7.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_6.UnsafePtr).V0, (*Constructor_Main_M3)(y_7.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_20
							} else {

							}
						}
						{
							if x_6.Type == 9 && x_6.IntVal == 3190619783 {
								__t20 = (y_7.Type == 9 && y_7.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_6.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_7.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_6.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_7.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_6.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_7.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_6.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_7.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_6.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_7.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_5_19.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_6.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_7.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_20
							} else {

							}
						}
						{
							if x_6.Type == 9 && x_6.IntVal == 108241190 {
								__t20 = (y_7.Type == 9 && y_7.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_6.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_7.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_6.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_7.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_6.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_7.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_6.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_7.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_6.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_7.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_5_19.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_6.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_7.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_20
							} else {

							}
						}
						{
							if x_6.Type == 9 && x_6.IntVal == 2066233029 {
								__t20 = (y_7.Type == 9 && y_7.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_6.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_7.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_6.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_7.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_6.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_7.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_5_19.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_6.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_7.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_6.UnsafePtr).V4, (*Constructor_Main_M6)(y_7.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_6.UnsafePtr).V5, (*Constructor_Main_M6)(y_7.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_6.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_7.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_6.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_7.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_6.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_7.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_6.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_7.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_6.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_7.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_5_19.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_6.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_7.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_6.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_7.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_6.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_7.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_6.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_7.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_6.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_7.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_6.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_7.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_5_19.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_6.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_7.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_20
							} else {

							}
						}
						{
							__t20 = (x_6.Type == 9 && x_6.IntVal == 1168316772) && ((y_7.Type == 9 && y_7.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_4_18, "eq"), (*Constructor_Main_M7)(x_6.UnsafePtr).V0, (*Constructor_Main_M7)(y_7.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_20:
						return gopurs_runtime.Bool(__t20)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t29.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m1()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_2_11
				// TAST (Let): __local_var_3_31 -> gopurs_runtime.Value
				__local_var_3_31 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_5 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_4, "zArrayA"), gopurs_runtime.RecordGet(rb_5, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_3_31
				// TAST (Let): __local_var_4_33 -> gopurs_runtime.Value
				__local_var_4_33 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_5, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_6, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_31, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_4_33
				// TAST (Let): __local_var_5_34 -> gopurs_runtime.Value
				__local_var_5_34 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_6, "fa"), gopurs_runtime.RecordGet(rb_7, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_33, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_6, rb_7).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_5_34
				// TAST (Let): __local_var_6_35 -> gopurs_runtime.Value
				__local_var_6_35 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_7, "fIgnore"), gopurs_runtime.RecordGet(rb_8, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_34, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_6_35
				// TAST (Let): __local_var_7_36 -> gopurs_runtime.Value
				__local_var_7_36 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_8 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_8, "arrayIgnore"), gopurs_runtime.RecordGet(rb_9, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_6_35, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_8, rb_9).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_7_36
				// TAST (Let): __local_var_4_32 -> gopurs_runtime.Value
				__local_var_4_32 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_8, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_9, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_7_36, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_8, rb_9).IntVal) != (0)))
					})
				}))
				_ = __local_var_4_32
				// TAST (Let): __local_var_5_37 -> gopurs_runtime.Value
				__local_var_5_37 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_32, "eq"), gopurs_runtime.RecordGet(ra_5, "nested"), gopurs_runtime.RecordGet(rb_6, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_5_37
				// TAST (Let): eqArray5_6_38 -> *Constructor_Data_Eq_Eq
				eqArray5_6_38 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_6_38
				_dollar___unused_3_30 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m2"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_8 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t39 bool
						{
							if x_7.Type == 9 && x_7.IntVal == 3852365315 {
								__t39 = (y_8.Type == 9 && y_8.IntVal == 3852365315)
								goto end_branch_39
							} else {

							}
						}
						{
							if x_7.Type == 9 && x_7.IntVal == 769986722 {
								__t39 = (y_8.Type == 9 && y_8.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_7.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_8.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_6_38.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_7.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_8.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_39
							} else {

							}
						}
						{
							if x_7.Type == 9 && x_7.IntVal == 2727978561 {
								__t39 = (y_8.Type == 9 && y_8.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_7.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_8.UnsafePtr).V0))
								goto end_branch_39
							} else {

							}
						}
						{
							if x_7.Type == 9 && x_7.IntVal == 1830062304 {
								__t39 = (y_8.Type == 9 && y_8.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_7.UnsafePtr).V0, (*Constructor_Main_M3)(y_8.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_39
							} else {

							}
						}
						{
							if x_7.Type == 9 && x_7.IntVal == 3190619783 {
								__t39 = (y_8.Type == 9 && y_8.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_7.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_8.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_7.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_8.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_7.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_8.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_7.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_8.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_7.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_8.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_6_38.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_7.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_8.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_39
							} else {

							}
						}
						{
							if x_7.Type == 9 && x_7.IntVal == 108241190 {
								__t39 = (y_8.Type == 9 && y_8.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_7.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_8.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_7.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_8.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_7.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_8.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_7.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_8.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_7.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_8.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_6_38.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_7.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_8.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_39
							} else {

							}
						}
						{
							if x_7.Type == 9 && x_7.IntVal == 2066233029 {
								__t39 = (y_8.Type == 9 && y_8.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_7.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_8.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_7.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_8.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_7.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_8.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_6_38.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_7.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_8.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_7.UnsafePtr).V4, (*Constructor_Main_M6)(y_8.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_7.UnsafePtr).V5, (*Constructor_Main_M6)(y_8.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_7.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_8.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_7.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_8.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_7.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_8.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_7.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_8.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_7.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_8.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_6_38.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_7.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_8.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_7.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_8.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_7.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_8.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_7.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_8.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_7.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_8.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_7.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_8.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_6_38.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_7.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_8.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_39
							} else {

							}
						}
						{
							__t39 = (x_7.Type == 9 && x_7.IntVal == 1168316772) && ((y_8.Type == 9 && y_8.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_5_37, "eq"), (*Constructor_Main_M7)(x_7.UnsafePtr).V0, (*Constructor_Main_M7)(y_8.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_39:
						return gopurs_runtime.Bool(__t39)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, 0})}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, 0})}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_3_30
				// TAST (Let): __local_var_4_41 -> gopurs_runtime.Value
				__local_var_4_41 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_5, "zArrayA"), gopurs_runtime.RecordGet(rb_6, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_4_41
				// TAST (Let): __local_var_5_43 -> gopurs_runtime.Value
				__local_var_5_43 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_6, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_7, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_41, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_6, rb_7).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_5_43
				// TAST (Let): __local_var_6_44 -> gopurs_runtime.Value
				__local_var_6_44 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_7, "fa"), gopurs_runtime.RecordGet(rb_8, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_43, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_6_44
				// TAST (Let): __local_var_7_45 -> gopurs_runtime.Value
				__local_var_7_45 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_8 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_8, "fIgnore"), gopurs_runtime.RecordGet(rb_9, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_6_44, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_8, rb_9).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_7_45
				// TAST (Let): __local_var_8_46 -> gopurs_runtime.Value
				__local_var_8_46 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_9 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_10 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_9, "arrayIgnore"), gopurs_runtime.RecordGet(rb_10, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_7_45, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_9, rb_10).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_8_46
				// TAST (Let): __local_var_5_42 -> gopurs_runtime.Value
				__local_var_5_42 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_9, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_10, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_8_46, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_9, rb_10).IntVal) != (0)))
					})
				}))
				_ = __local_var_5_42
				// TAST (Let): __local_var_6_47 -> gopurs_runtime.Value
				__local_var_6_47 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_42, "eq"), gopurs_runtime.RecordGet(ra_6, "nested"), gopurs_runtime.RecordGet(rb_7, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_6_47
				// TAST (Let): eqArray5_7_48 -> *Constructor_Data_Eq_Eq
				eqArray5_7_48 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_7_48
				var __t58 gopurs_runtime.Value
				{
					var __t_tag_50 gopurs_runtime.Value = Get_Main_m3()
					if __t_tag_50.Type == 9 && __t_tag_50.IntVal == 3852365315 {
						__t58 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M0{nil}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_58
					} else {

					}
				}
				{
					var __t_tag_51 gopurs_runtime.Value = Get_Main_m3()
					if __t_tag_51.Type == 9 && __t_tag_51.IntVal == 769986722 {
						__t58 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap10 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Main_M1)(Get_Main_m3().UnsafePtr).V0}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap10
							arr_go_arrayMap10 := (*[]gopurs_runtime.Value)(arr_val_arrayMap10.UnsafePtr)
							_ = arr_go_arrayMap10
							res_go_arrayMap10 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap10))
							_ = res_go_arrayMap10
							for i_arrayMap10, v_arrayMap10 := range *arr_go_arrayMap10 {
								res_go_arrayMap10[i_arrayMap10] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_6 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, v2_5, func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v3_6.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()}))}
									})
								}), v_arrayMap10)
							}
							return gopurs_runtime.Array(res_go_arrayMap10)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array((*Constructor_Main_M1)(Get_Main_m3().UnsafePtr).V1)))
						goto end_branch_58
					} else {

					}
				}
				{
					var __t_tag_52 gopurs_runtime.Value = Get_Main_m3()
					if __t_tag_52.Type == 9 && __t_tag_52.IntVal == 2727978561 {
						__t58 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, (*Constructor_Main_M2)(Get_Main_m3().UnsafePtr).V0})}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_58
					} else {

					}
				}
				{
					var __t_tag_53 gopurs_runtime.Value = Get_Main_m3()
					if __t_tag_53.Type == 9 && __t_tag_53.IntVal == 1830062304 {
						__t58 = func() gopurs_runtime.Value {
							arr_val_arrayMap9 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), (*Constructor_Main_M3)(Get_Main_m3().UnsafePtr).V0)
							_ = arr_val_arrayMap9
							arr_go_arrayMap9 := (*[]gopurs_runtime.Value)(arr_val_arrayMap9.UnsafePtr)
							_ = arr_go_arrayMap9
							res_go_arrayMap9 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap9))
							_ = res_go_arrayMap9
							for i_arrayMap9, v_arrayMap9 := range *arr_go_arrayMap9 {
								res_go_arrayMap9[i_arrayMap9] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, v1_5}))}
								}), v_arrayMap9)
							}
							return gopurs_runtime.Array(res_go_arrayMap9)
						}()
						goto end_branch_58
					} else {

					}
				}
				{
					var __t_tag_54 gopurs_runtime.Value = Get_Main_m3()
					if __t_tag_54.Type == 9 && __t_tag_54.IntVal == 3190619783 {
						__t58 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap11 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m3().UnsafePtr).V0, "a")}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap11
							arr_go_arrayMap11 := (*[]gopurs_runtime.Value)(arr_val_arrayMap11.UnsafePtr)
							_ = arr_go_arrayMap11
							res_go_arrayMap11 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap11))
							_ = res_go_arrayMap11
							for i_arrayMap11, v_arrayMap11 := range *arr_go_arrayMap11 {
								res_go_arrayMap11[i_arrayMap11] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4)(Get_Main_m3().UnsafePtr).V0, "a", v1_5, "fa", v2_6, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_7.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}()))}))}
										})
									})
								}), v_arrayMap11)
							}
							return gopurs_runtime.Array(res_go_arrayMap11)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m3().UnsafePtr).V0, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m3().UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_58
					} else {

					}
				}
				{
					var __t_tag_55 gopurs_runtime.Value = Get_Main_m3()
					if __t_tag_55.Type == 9 && __t_tag_55.IntVal == 108241190 {
						__t58 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap11 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m3().UnsafePtr).V0, "nested"), "a")}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap11
							arr_go_arrayMap11 := (*[]gopurs_runtime.Value)(arr_val_arrayMap11.UnsafePtr)
							_ = arr_go_arrayMap11
							res_go_arrayMap11 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap11))
							_ = res_go_arrayMap11
							for i_arrayMap11, v_arrayMap11 := range *arr_go_arrayMap11 {
								res_go_arrayMap11[i_arrayMap11] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, func() gopurs_runtime.Value {
												origVal := (*Constructor_Main_M5)(Get_Main_m3().UnsafePtr).V0
												if origVal.Type != gopurs_runtime.TypeRecord1 {
													return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m3().UnsafePtr).V0, "nested"), "a", v1_5, "fa", v2_6, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v3_7.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}()))})
												}
												clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
												clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m3().UnsafePtr).V0, "nested"), "a", v1_5, "fa", v2_6, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v3_7.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}()))
												return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
											}()}))}
										})
									})
								}), v_arrayMap11)
							}
							return gopurs_runtime.Array(res_go_arrayMap11)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m3().UnsafePtr).V0, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m3().UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_58
					} else {

					}
				}
				{
					var __t_tag_56 gopurs_runtime.Value = Get_Main_m3()
					if __t_tag_56.Type == 9 && __t_tag_56.IntVal == 2066233029 {
						__t58 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap17 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V1}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap17
							arr_go_arrayMap17 := (*[]gopurs_runtime.Value)(arr_val_arrayMap17.UnsafePtr)
							_ = arr_go_arrayMap17
							res_go_arrayMap17 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap17))
							_ = res_go_arrayMap17
							for i_arrayMap17, v_arrayMap17 := range *arr_go_arrayMap17 {
								res_go_arrayMap17[i_arrayMap17] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v8_5 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v9_6 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v10_7 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v11_8 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v12_9 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v13_10 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Func(func(v14_11 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(v15_12 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(v16_13 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, (*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V0, v8_5, (*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V2, func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v9_6.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}(), v10_7, (*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V6, "a", v11_8, "fa", v12_9, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v13_10.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}())), func() gopurs_runtime.Value {
																		origVal := (*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V7
																		if origVal.Type != gopurs_runtime.TypeRecord1 {
																			return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V7, "nested"), "a", v14_11, "fa", v15_12, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																				arr := *(*[]gopurs_runtime.Value)(v16_13.UnsafePtr)
																				unboxed := make([]gopurs_runtime.Value, len(arr))
																				for i, v := range arr {
																					unboxed[i] = v
																				}
																				return unboxed
																			}()))})
																		}
																		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
																		clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V7, "nested"), "a", v14_11, "fa", v15_12, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																			arr := *(*[]gopurs_runtime.Value)(v16_13.UnsafePtr)
																			unboxed := make([]gopurs_runtime.Value, len(arr))
																			for i, v := range arr {
																				unboxed[i] = v
																			}
																			return unboxed
																		}()))
																		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
																	}()}))}
																})
															})
														})
													})
												})
											})
										})
									})
								}), v_arrayMap17)
							}
							return gopurs_runtime.Array(res_go_arrayMap17)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array((*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V3))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), (*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V4)), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V6, "a")}).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V6, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V7, "nested"), "a")}).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V7, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3().UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_58
					} else {

					}
				}
				{
					var __t_tag_57 gopurs_runtime.Value = Get_Main_m3()
					if __t_tag_57.Type == 9 && __t_tag_57.IntVal == 1168316772 {
						__t58 = func() gopurs_runtime.Value {
							arr_val_arrayMap9 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply5(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
									arr_val_arrayMap15 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_5, "nested"), "a")}).UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap15
									arr_go_arrayMap15 := (*[]gopurs_runtime.Value)(arr_val_arrayMap15.UnsafePtr)
									_ = arr_go_arrayMap15
									res_go_arrayMap15 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap15))
									_ = res_go_arrayMap15
									for i_arrayMap15, v_arrayMap15 := range *arr_go_arrayMap15 {
										res_go_arrayMap15[i_arrayMap15] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.RecordUpdate1(v1_5, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_5, "nested"), "a", v2_6, "fa", v3_7, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v4_8.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}())))
												})
											})
										}), v_arrayMap15)
									}
									return gopurs_runtime.Array(res_go_arrayMap15)
								}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_5, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_5, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
							})), (*Constructor_Main_M7)(Get_Main_m3().UnsafePtr).V0)
							_ = arr_val_arrayMap9
							arr_go_arrayMap9 := (*[]gopurs_runtime.Value)(arr_val_arrayMap9.UnsafePtr)
							_ = arr_go_arrayMap9
							res_go_arrayMap9 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap9))
							_ = res_go_arrayMap9
							for i_arrayMap9, v_arrayMap9 := range *arr_go_arrayMap9 {
								res_go_arrayMap9[i_arrayMap9] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, v1_5}))}
								}), v_arrayMap9)
							}
							return gopurs_runtime.Array(res_go_arrayMap9)
						}()
						goto end_branch_58
					} else {

					}
				}
				{
					__t58 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_58:
				_dollar___unused_4_40 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m3"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_9 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t49 bool
						{
							if x_8.Type == 9 && x_8.IntVal == 3852365315 {
								__t49 = (y_9.Type == 9 && y_9.IntVal == 3852365315)
								goto end_branch_49
							} else {

							}
						}
						{
							if x_8.Type == 9 && x_8.IntVal == 769986722 {
								__t49 = (y_9.Type == 9 && y_9.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_8.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_9.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_7_48.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_8.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_9.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_49
							} else {

							}
						}
						{
							if x_8.Type == 9 && x_8.IntVal == 2727978561 {
								__t49 = (y_9.Type == 9 && y_9.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_8.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_9.UnsafePtr).V0))
								goto end_branch_49
							} else {

							}
						}
						{
							if x_8.Type == 9 && x_8.IntVal == 1830062304 {
								__t49 = (y_9.Type == 9 && y_9.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_8.UnsafePtr).V0, (*Constructor_Main_M3)(y_9.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_49
							} else {

							}
						}
						{
							if x_8.Type == 9 && x_8.IntVal == 3190619783 {
								__t49 = (y_9.Type == 9 && y_9.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_8.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_9.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_8.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_9.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_8.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_9.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_8.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_9.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_8.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_9.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_7_48.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_8.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_9.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_49
							} else {

							}
						}
						{
							if x_8.Type == 9 && x_8.IntVal == 108241190 {
								__t49 = (y_9.Type == 9 && y_9.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_8.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_9.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_8.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_9.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_8.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_9.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_8.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_9.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_8.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_9.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_7_48.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_8.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_9.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_49
							} else {

							}
						}
						{
							if x_8.Type == 9 && x_8.IntVal == 2066233029 {
								__t49 = (y_9.Type == 9 && y_9.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_8.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_9.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_8.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_9.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_8.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_9.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_7_48.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_8.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_9.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_8.UnsafePtr).V4, (*Constructor_Main_M6)(y_9.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_8.UnsafePtr).V5, (*Constructor_Main_M6)(y_9.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_8.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_9.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_8.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_9.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_8.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_9.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_8.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_9.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_8.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_9.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_7_48.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_8.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_9.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_8.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_9.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_8.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_9.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_8.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_9.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_8.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_9.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_8.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_9.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_7_48.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_8.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_9.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_49
							} else {

							}
						}
						{
							__t49 = (x_8.Type == 9 && x_8.IntVal == 1168316772) && ((y_9.Type == 9 && y_9.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_6_47, "eq"), (*Constructor_Main_M7)(x_8.UnsafePtr).V0, (*Constructor_Main_M7)(y_9.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_49:
						return gopurs_runtime.Bool(__t49)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t58.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m3()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_4_40
				// TAST (Let): __local_var_5_60 -> gopurs_runtime.Value
				__local_var_5_60 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_6, "zArrayA"), gopurs_runtime.RecordGet(rb_7, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_5_60
				// TAST (Let): __local_var_6_62 -> gopurs_runtime.Value
				__local_var_6_62 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_7, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_8, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_60, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_6_62
				// TAST (Let): __local_var_7_63 -> gopurs_runtime.Value
				__local_var_7_63 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_8 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_8, "fa"), gopurs_runtime.RecordGet(rb_9, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_6_62, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_8, rb_9).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_7_63
				// TAST (Let): __local_var_8_64 -> gopurs_runtime.Value
				__local_var_8_64 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_9 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_10 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_9, "fIgnore"), gopurs_runtime.RecordGet(rb_10, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_7_63, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_9, rb_10).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_8_64
				// TAST (Let): __local_var_9_65 -> gopurs_runtime.Value
				__local_var_9_65 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_10, "arrayIgnore"), gopurs_runtime.RecordGet(rb_11, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_8_64, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_10, rb_11).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_9_65
				// TAST (Let): __local_var_6_61 -> gopurs_runtime.Value
				__local_var_6_61 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_11 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_10, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_11, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_9_65, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_10, rb_11).IntVal) != (0)))
					})
				}))
				_ = __local_var_6_61
				// TAST (Let): __local_var_7_66 -> gopurs_runtime.Value
				__local_var_7_66 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_61, "eq"), gopurs_runtime.RecordGet(ra_7, "nested"), gopurs_runtime.RecordGet(rb_8, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_7_66
				// TAST (Let): eqArray5_8_67 -> *Constructor_Data_Eq_Eq
				eqArray5_8_67 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_8_67
				_dollar___unused_5_59 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m4"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_10 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t68 bool
						{
							if x_9.Type == 9 && x_9.IntVal == 3852365315 {
								__t68 = (y_10.Type == 9 && y_10.IntVal == 3852365315)
								goto end_branch_68
							} else {

							}
						}
						{
							if x_9.Type == 9 && x_9.IntVal == 769986722 {
								__t68 = (y_10.Type == 9 && y_10.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_9.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_10.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_8_67.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_9.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_10.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_68
							} else {

							}
						}
						{
							if x_9.Type == 9 && x_9.IntVal == 2727978561 {
								__t68 = (y_10.Type == 9 && y_10.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_9.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_10.UnsafePtr).V0))
								goto end_branch_68
							} else {

							}
						}
						{
							if x_9.Type == 9 && x_9.IntVal == 1830062304 {
								__t68 = (y_10.Type == 9 && y_10.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_9.UnsafePtr).V0, (*Constructor_Main_M3)(y_10.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_68
							} else {

							}
						}
						{
							if x_9.Type == 9 && x_9.IntVal == 3190619783 {
								__t68 = (y_10.Type == 9 && y_10.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_9.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_10.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_9.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_10.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_9.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_10.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_9.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_10.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_9.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_10.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_8_67.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_9.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_10.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_68
							} else {

							}
						}
						{
							if x_9.Type == 9 && x_9.IntVal == 108241190 {
								__t68 = (y_10.Type == 9 && y_10.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_9.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_10.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_9.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_10.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_9.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_10.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_9.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_10.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_9.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_10.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_8_67.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_9.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_10.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_68
							} else {

							}
						}
						{
							if x_9.Type == 9 && x_9.IntVal == 2066233029 {
								__t68 = (y_10.Type == 9 && y_10.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_9.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_10.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_9.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_10.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_9.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_10.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_8_67.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_9.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_10.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_9.UnsafePtr).V4, (*Constructor_Main_M6)(y_10.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_9.UnsafePtr).V5, (*Constructor_Main_M6)(y_10.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_9.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_10.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_9.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_10.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_9.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_10.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_9.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_10.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_9.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_10.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_8_67.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_9.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_10.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_9.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_10.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_9.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_10.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_9.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_10.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_9.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_10.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_9.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_10.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_8_67.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_9.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_10.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_68
							} else {

							}
						}
						{
							__t68 = (x_9.Type == 9 && x_9.IntVal == 1168316772) && ((y_10.Type == 9 && y_10.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_7_66, "eq"), (*Constructor_Main_M7)(x_9.UnsafePtr).V0, (*Constructor_Main_M7)(y_10.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_68:
						return gopurs_runtime.Bool(__t68)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
						arr_val_arrayMap11 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Str("a")}).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						_ = arr_val_arrayMap11
						arr_go_arrayMap11 := (*[]gopurs_runtime.Value)(arr_val_arrayMap11.UnsafePtr)
						_ = arr_go_arrayMap11
						res_go_arrayMap11 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap11))
						_ = res_go_arrayMap11
						for i_arrayMap11, v_arrayMap11 := range *arr_go_arrayMap11 {
							res_go_arrayMap11[i_arrayMap11] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, func() gopurs_runtime.Value {
											origVal := Get_Main_recordValue()
											if origVal.Type != gopurs_runtime.TypeRecordData {
												return gopurs_runtime.RecordUpdateDict(origVal, []string{"a", "fa", "zArrayA"}, []gopurs_runtime.Value{v1_6, v2_7, gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v3_8.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}())})
											}
											r := (*gopurs_runtime.RecordData)(origVal.UnsafePtr)
											newVals := make([]gopurs_runtime.Value, len(r.Vals))
											copy(newVals, r.Vals)
											newVals[0] = v1_6
											newVals[3] = v2_7
											newVals[5] = gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_8.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}())
											newR := gopurs_runtime.RecordData{Keys: r.Keys, Vals: newVals}
											return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecordData, UnsafePtr: unsafe.Pointer(&newR)}
										}()}))}
									})
								})
							}), v_arrayMap11)
						}
						return gopurs_runtime.Array(res_go_arrayMap11)
					}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(Get_Main_recordValue(), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(Get_Main_recordValue(), "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr := []*Constructor_Main_M4{(&Constructor_Main_M4{1, Get_Main_recordValue()})}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_5_59
				// TAST (Let): __local_var_6_70 -> gopurs_runtime.Value
				__local_var_6_70 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_7, "zArrayA"), gopurs_runtime.RecordGet(rb_8, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_6_70
				// TAST (Let): __local_var_7_72 -> gopurs_runtime.Value
				__local_var_7_72 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_8 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_8, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_9, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_6_70, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_8, rb_9).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_7_72
				// TAST (Let): __local_var_8_73 -> gopurs_runtime.Value
				__local_var_8_73 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_9 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_10 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_9, "fa"), gopurs_runtime.RecordGet(rb_10, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_7_72, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_9, rb_10).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_8_73
				// TAST (Let): __local_var_9_74 -> gopurs_runtime.Value
				__local_var_9_74 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_10, "fIgnore"), gopurs_runtime.RecordGet(rb_11, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_8_73, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_10, rb_11).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_9_74
				// TAST (Let): __local_var_10_75 -> gopurs_runtime.Value
				__local_var_10_75 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_11 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_12 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_11, "arrayIgnore"), gopurs_runtime.RecordGet(rb_12, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_9_74, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_11, rb_12).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_10_75
				// TAST (Let): __local_var_7_71 -> gopurs_runtime.Value
				__local_var_7_71 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_12 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_11, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_12, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_10_75, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_11, rb_12).IntVal) != (0)))
					})
				}))
				_ = __local_var_7_71
				// TAST (Let): __local_var_8_76 -> gopurs_runtime.Value
				__local_var_8_76 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_71, "eq"), gopurs_runtime.RecordGet(ra_8, "nested"), gopurs_runtime.RecordGet(rb_9, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_8_76
				// TAST (Let): eqArray5_9_77 -> *Constructor_Data_Eq_Eq
				eqArray5_9_77 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_9_77
				var __t87 gopurs_runtime.Value
				{
					var __t_tag_79 gopurs_runtime.Value = Get_Main_m5()
					if __t_tag_79.Type == 9 && __t_tag_79.IntVal == 3852365315 {
						__t87 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M0{nil}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_87
					} else {

					}
				}
				{
					var __t_tag_80 gopurs_runtime.Value = Get_Main_m5()
					if __t_tag_80.Type == 9 && __t_tag_80.IntVal == 769986722 {
						__t87 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap12 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Main_M1)(Get_Main_m5().UnsafePtr).V0}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap12
							arr_go_arrayMap12 := (*[]gopurs_runtime.Value)(arr_val_arrayMap12.UnsafePtr)
							_ = arr_go_arrayMap12
							res_go_arrayMap12 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap12))
							_ = res_go_arrayMap12
							for i_arrayMap12, v_arrayMap12 := range *arr_go_arrayMap12 {
								res_go_arrayMap12[i_arrayMap12] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, v2_7, func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v3_8.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()}))}
									})
								}), v_arrayMap12)
							}
							return gopurs_runtime.Array(res_go_arrayMap12)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array((*Constructor_Main_M1)(Get_Main_m5().UnsafePtr).V1)))
						goto end_branch_87
					} else {

					}
				}
				{
					var __t_tag_81 gopurs_runtime.Value = Get_Main_m5()
					if __t_tag_81.Type == 9 && __t_tag_81.IntVal == 2727978561 {
						__t87 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, (*Constructor_Main_M2)(Get_Main_m5().UnsafePtr).V0})}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_87
					} else {

					}
				}
				{
					var __t_tag_82 gopurs_runtime.Value = Get_Main_m5()
					if __t_tag_82.Type == 9 && __t_tag_82.IntVal == 1830062304 {
						__t87 = func() gopurs_runtime.Value {
							arr_val_arrayMap11 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), (*Constructor_Main_M3)(Get_Main_m5().UnsafePtr).V0)
							_ = arr_val_arrayMap11
							arr_go_arrayMap11 := (*[]gopurs_runtime.Value)(arr_val_arrayMap11.UnsafePtr)
							_ = arr_go_arrayMap11
							res_go_arrayMap11 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap11))
							_ = res_go_arrayMap11
							for i_arrayMap11, v_arrayMap11 := range *arr_go_arrayMap11 {
								res_go_arrayMap11[i_arrayMap11] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, v1_7}))}
								}), v_arrayMap11)
							}
							return gopurs_runtime.Array(res_go_arrayMap11)
						}()
						goto end_branch_87
					} else {

					}
				}
				{
					var __t_tag_83 gopurs_runtime.Value = Get_Main_m5()
					if __t_tag_83.Type == 9 && __t_tag_83.IntVal == 3190619783 {
						__t87 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap13 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m5().UnsafePtr).V0, "a")}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap13
							arr_go_arrayMap13 := (*[]gopurs_runtime.Value)(arr_val_arrayMap13.UnsafePtr)
							_ = arr_go_arrayMap13
							res_go_arrayMap13 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap13))
							_ = res_go_arrayMap13
							for i_arrayMap13, v_arrayMap13 := range *arr_go_arrayMap13 {
								res_go_arrayMap13[i_arrayMap13] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4)(Get_Main_m5().UnsafePtr).V0, "a", v1_7, "fa", v2_8, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_9.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}()))}))}
										})
									})
								}), v_arrayMap13)
							}
							return gopurs_runtime.Array(res_go_arrayMap13)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m5().UnsafePtr).V0, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m5().UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_87
					} else {

					}
				}
				{
					var __t_tag_84 gopurs_runtime.Value = Get_Main_m5()
					if __t_tag_84.Type == 9 && __t_tag_84.IntVal == 108241190 {
						__t87 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap13 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m5().UnsafePtr).V0, "nested"), "a")}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap13
							arr_go_arrayMap13 := (*[]gopurs_runtime.Value)(arr_val_arrayMap13.UnsafePtr)
							_ = arr_go_arrayMap13
							res_go_arrayMap13 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap13))
							_ = res_go_arrayMap13
							for i_arrayMap13, v_arrayMap13 := range *arr_go_arrayMap13 {
								res_go_arrayMap13[i_arrayMap13] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, func() gopurs_runtime.Value {
												origVal := (*Constructor_Main_M5)(Get_Main_m5().UnsafePtr).V0
												if origVal.Type != gopurs_runtime.TypeRecord1 {
													return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m5().UnsafePtr).V0, "nested"), "a", v1_7, "fa", v2_8, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v3_9.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}()))})
												}
												clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
												clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m5().UnsafePtr).V0, "nested"), "a", v1_7, "fa", v2_8, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v3_9.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}()))
												return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
											}()}))}
										})
									})
								}), v_arrayMap13)
							}
							return gopurs_runtime.Array(res_go_arrayMap13)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m5().UnsafePtr).V0, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m5().UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_87
					} else {

					}
				}
				{
					var __t_tag_85 gopurs_runtime.Value = Get_Main_m5()
					if __t_tag_85.Type == 9 && __t_tag_85.IntVal == 2066233029 {
						__t87 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap19 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V1}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap19
							arr_go_arrayMap19 := (*[]gopurs_runtime.Value)(arr_val_arrayMap19.UnsafePtr)
							_ = arr_go_arrayMap19
							res_go_arrayMap19 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap19))
							_ = res_go_arrayMap19
							for i_arrayMap19, v_arrayMap19 := range *arr_go_arrayMap19 {
								res_go_arrayMap19[i_arrayMap19] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v8_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v9_8 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v10_9 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v11_10 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v12_11 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v13_12 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Func(func(v14_13 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(v15_14 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(v16_15 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, (*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V0, v8_7, (*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V2, func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v9_8.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}(), v10_9, (*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V6, "a", v11_10, "fa", v12_11, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v13_12.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}())), func() gopurs_runtime.Value {
																		origVal := (*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V7
																		if origVal.Type != gopurs_runtime.TypeRecord1 {
																			return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V7, "nested"), "a", v14_13, "fa", v15_14, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																				arr := *(*[]gopurs_runtime.Value)(v16_15.UnsafePtr)
																				unboxed := make([]gopurs_runtime.Value, len(arr))
																				for i, v := range arr {
																					unboxed[i] = v
																				}
																				return unboxed
																			}()))})
																		}
																		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
																		clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V7, "nested"), "a", v14_13, "fa", v15_14, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																			arr := *(*[]gopurs_runtime.Value)(v16_15.UnsafePtr)
																			unboxed := make([]gopurs_runtime.Value, len(arr))
																			for i, v := range arr {
																				unboxed[i] = v
																			}
																			return unboxed
																		}()))
																		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
																	}()}))}
																})
															})
														})
													})
												})
											})
										})
									})
								}), v_arrayMap19)
							}
							return gopurs_runtime.Array(res_go_arrayMap19)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array((*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V3))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), (*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V4)), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V6, "a")}).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V6, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V7, "nested"), "a")}).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V7, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5().UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_87
					} else {

					}
				}
				{
					var __t_tag_86 gopurs_runtime.Value = Get_Main_m5()
					if __t_tag_86.Type == 9 && __t_tag_86.IntVal == 1168316772 {
						__t87 = func() gopurs_runtime.Value {
							arr_val_arrayMap11 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply5(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
									arr_val_arrayMap17 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_7, "nested"), "a")}).UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap17
									arr_go_arrayMap17 := (*[]gopurs_runtime.Value)(arr_val_arrayMap17.UnsafePtr)
									_ = arr_go_arrayMap17
									res_go_arrayMap17 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap17))
									_ = res_go_arrayMap17
									for i_arrayMap17, v_arrayMap17 := range *arr_go_arrayMap17 {
										res_go_arrayMap17[i_arrayMap17] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.RecordUpdate1(v1_7, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_7, "nested"), "a", v2_8, "fa", v3_9, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v4_10.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}())))
												})
											})
										}), v_arrayMap17)
									}
									return gopurs_runtime.Array(res_go_arrayMap17)
								}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_7, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
							})), (*Constructor_Main_M7)(Get_Main_m5().UnsafePtr).V0)
							_ = arr_val_arrayMap11
							arr_go_arrayMap11 := (*[]gopurs_runtime.Value)(arr_val_arrayMap11.UnsafePtr)
							_ = arr_go_arrayMap11
							res_go_arrayMap11 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap11))
							_ = res_go_arrayMap11
							for i_arrayMap11, v_arrayMap11 := range *arr_go_arrayMap11 {
								res_go_arrayMap11[i_arrayMap11] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, v1_7}))}
								}), v_arrayMap11)
							}
							return gopurs_runtime.Array(res_go_arrayMap11)
						}()
						goto end_branch_87
					} else {

					}
				}
				{
					__t87 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_87:
				_dollar___unused_6_69 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m5"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_11 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t78 bool
						{
							if x_10.Type == 9 && x_10.IntVal == 3852365315 {
								__t78 = (y_11.Type == 9 && y_11.IntVal == 3852365315)
								goto end_branch_78
							} else {

							}
						}
						{
							if x_10.Type == 9 && x_10.IntVal == 769986722 {
								__t78 = (y_11.Type == 9 && y_11.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_10.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_11.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_9_77.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_10.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_11.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_78
							} else {

							}
						}
						{
							if x_10.Type == 9 && x_10.IntVal == 2727978561 {
								__t78 = (y_11.Type == 9 && y_11.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_10.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_11.UnsafePtr).V0))
								goto end_branch_78
							} else {

							}
						}
						{
							if x_10.Type == 9 && x_10.IntVal == 1830062304 {
								__t78 = (y_11.Type == 9 && y_11.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_10.UnsafePtr).V0, (*Constructor_Main_M3)(y_11.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_78
							} else {

							}
						}
						{
							if x_10.Type == 9 && x_10.IntVal == 3190619783 {
								__t78 = (y_11.Type == 9 && y_11.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_10.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_11.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_10.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_11.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_10.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_11.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_10.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_11.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_10.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_11.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_9_77.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_10.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_11.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_78
							} else {

							}
						}
						{
							if x_10.Type == 9 && x_10.IntVal == 108241190 {
								__t78 = (y_11.Type == 9 && y_11.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_10.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_11.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_10.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_11.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_10.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_11.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_10.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_11.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_10.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_11.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_9_77.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_10.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_11.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_78
							} else {

							}
						}
						{
							if x_10.Type == 9 && x_10.IntVal == 2066233029 {
								__t78 = (y_11.Type == 9 && y_11.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_10.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_11.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_10.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_11.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_10.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_11.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_9_77.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_10.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_11.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_10.UnsafePtr).V4, (*Constructor_Main_M6)(y_11.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_10.UnsafePtr).V5, (*Constructor_Main_M6)(y_11.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_10.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_11.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_10.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_11.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_10.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_11.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_10.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_11.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_10.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_11.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_9_77.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_10.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_11.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_10.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_11.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_10.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_11.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_10.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_11.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_10.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_11.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_10.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_11.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_9_77.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_10.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_11.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_78
							} else {

							}
						}
						{
							__t78 = (x_10.Type == 9 && x_10.IntVal == 1168316772) && ((y_11.Type == 9 && y_11.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_8_76, "eq"), (*Constructor_Main_M7)(x_10.UnsafePtr).V0, (*Constructor_Main_M7)(y_11.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_78:
						return gopurs_runtime.Bool(__t78)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t87.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m5()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_6_69
				// TAST (Let): __local_var_7_89 -> gopurs_runtime.Value
				__local_var_7_89 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_8 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_8, "zArrayA"), gopurs_runtime.RecordGet(rb_9, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_7_89
				// TAST (Let): __local_var_8_91 -> gopurs_runtime.Value
				__local_var_8_91 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_9 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_10 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_9, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_10, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_7_89, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_9, rb_10).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_8_91
				// TAST (Let): __local_var_9_92 -> gopurs_runtime.Value
				__local_var_9_92 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_10, "fa"), gopurs_runtime.RecordGet(rb_11, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_8_91, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_10, rb_11).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_9_92
				// TAST (Let): __local_var_10_93 -> gopurs_runtime.Value
				__local_var_10_93 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_11 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_12 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_11, "fIgnore"), gopurs_runtime.RecordGet(rb_12, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_9_92, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_11, rb_12).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_10_93
				// TAST (Let): __local_var_11_94 -> gopurs_runtime.Value
				__local_var_11_94 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_12 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_12, "arrayIgnore"), gopurs_runtime.RecordGet(rb_13, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_10_93, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_12, rb_13).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_11_94
				// TAST (Let): __local_var_8_90 -> gopurs_runtime.Value
				__local_var_8_90 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_12 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_13 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_12, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_13, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_11_94, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_12, rb_13).IntVal) != (0)))
					})
				}))
				_ = __local_var_8_90
				// TAST (Let): __local_var_9_95 -> gopurs_runtime.Value
				__local_var_9_95 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_90, "eq"), gopurs_runtime.RecordGet(ra_9, "nested"), gopurs_runtime.RecordGet(rb_10, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_9_95
				// TAST (Let): eqArray5_10_96 -> *Constructor_Data_Eq_Eq
				eqArray5_10_96 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_10_96
				var __t106 gopurs_runtime.Value
				{
					var __t_tag_98 gopurs_runtime.Value = Get_Main_m6()
					if __t_tag_98.Type == 9 && __t_tag_98.IntVal == 3852365315 {
						__t106 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M0{nil}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_106
					} else {

					}
				}
				{
					var __t_tag_99 gopurs_runtime.Value = Get_Main_m6()
					if __t_tag_99.Type == 9 && __t_tag_99.IntVal == 769986722 {
						__t106 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap13 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Main_M1)(Get_Main_m6().UnsafePtr).V0}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap13
							arr_go_arrayMap13 := (*[]gopurs_runtime.Value)(arr_val_arrayMap13.UnsafePtr)
							_ = arr_go_arrayMap13
							res_go_arrayMap13 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap13))
							_ = res_go_arrayMap13
							for i_arrayMap13, v_arrayMap13 := range *arr_go_arrayMap13 {
								res_go_arrayMap13[i_arrayMap13] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, v2_8, func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v3_9.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()}))}
									})
								}), v_arrayMap13)
							}
							return gopurs_runtime.Array(res_go_arrayMap13)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array((*Constructor_Main_M1)(Get_Main_m6().UnsafePtr).V1)))
						goto end_branch_106
					} else {

					}
				}
				{
					var __t_tag_100 gopurs_runtime.Value = Get_Main_m6()
					if __t_tag_100.Type == 9 && __t_tag_100.IntVal == 2727978561 {
						__t106 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, (*Constructor_Main_M2)(Get_Main_m6().UnsafePtr).V0})}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_106
					} else {

					}
				}
				{
					var __t_tag_101 gopurs_runtime.Value = Get_Main_m6()
					if __t_tag_101.Type == 9 && __t_tag_101.IntVal == 1830062304 {
						__t106 = func() gopurs_runtime.Value {
							arr_val_arrayMap12 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), (*Constructor_Main_M3)(Get_Main_m6().UnsafePtr).V0)
							_ = arr_val_arrayMap12
							arr_go_arrayMap12 := (*[]gopurs_runtime.Value)(arr_val_arrayMap12.UnsafePtr)
							_ = arr_go_arrayMap12
							res_go_arrayMap12 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap12))
							_ = res_go_arrayMap12
							for i_arrayMap12, v_arrayMap12 := range *arr_go_arrayMap12 {
								res_go_arrayMap12[i_arrayMap12] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, v1_8}))}
								}), v_arrayMap12)
							}
							return gopurs_runtime.Array(res_go_arrayMap12)
						}()
						goto end_branch_106
					} else {

					}
				}
				{
					var __t_tag_102 gopurs_runtime.Value = Get_Main_m6()
					if __t_tag_102.Type == 9 && __t_tag_102.IntVal == 3190619783 {
						__t106 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap14 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m6().UnsafePtr).V0, "a")}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap14
							arr_go_arrayMap14 := (*[]gopurs_runtime.Value)(arr_val_arrayMap14.UnsafePtr)
							_ = arr_go_arrayMap14
							res_go_arrayMap14 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap14))
							_ = res_go_arrayMap14
							for i_arrayMap14, v_arrayMap14 := range *arr_go_arrayMap14 {
								res_go_arrayMap14[i_arrayMap14] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4)(Get_Main_m6().UnsafePtr).V0, "a", v1_8, "fa", v2_9, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_10.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}()))}))}
										})
									})
								}), v_arrayMap14)
							}
							return gopurs_runtime.Array(res_go_arrayMap14)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m6().UnsafePtr).V0, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m6().UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_106
					} else {

					}
				}
				{
					var __t_tag_103 gopurs_runtime.Value = Get_Main_m6()
					if __t_tag_103.Type == 9 && __t_tag_103.IntVal == 108241190 {
						__t106 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap14 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m6().UnsafePtr).V0, "nested"), "a")}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap14
							arr_go_arrayMap14 := (*[]gopurs_runtime.Value)(arr_val_arrayMap14.UnsafePtr)
							_ = arr_go_arrayMap14
							res_go_arrayMap14 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap14))
							_ = res_go_arrayMap14
							for i_arrayMap14, v_arrayMap14 := range *arr_go_arrayMap14 {
								res_go_arrayMap14[i_arrayMap14] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, func() gopurs_runtime.Value {
												origVal := (*Constructor_Main_M5)(Get_Main_m6().UnsafePtr).V0
												if origVal.Type != gopurs_runtime.TypeRecord1 {
													return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m6().UnsafePtr).V0, "nested"), "a", v1_8, "fa", v2_9, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v3_10.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}()))})
												}
												clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
												clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m6().UnsafePtr).V0, "nested"), "a", v1_8, "fa", v2_9, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v3_10.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}()))
												return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
											}()}))}
										})
									})
								}), v_arrayMap14)
							}
							return gopurs_runtime.Array(res_go_arrayMap14)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m6().UnsafePtr).V0, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m6().UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_106
					} else {

					}
				}
				{
					var __t_tag_104 gopurs_runtime.Value = Get_Main_m6()
					if __t_tag_104.Type == 9 && __t_tag_104.IntVal == 2066233029 {
						__t106 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap20 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V1}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap20
							arr_go_arrayMap20 := (*[]gopurs_runtime.Value)(arr_val_arrayMap20.UnsafePtr)
							_ = arr_go_arrayMap20
							res_go_arrayMap20 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap20))
							_ = res_go_arrayMap20
							for i_arrayMap20, v_arrayMap20 := range *arr_go_arrayMap20 {
								res_go_arrayMap20[i_arrayMap20] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v8_8 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v9_9 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v10_10 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v11_11 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v12_12 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v13_13 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Func(func(v14_14 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(v15_15 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(v16_16 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, (*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V0, v8_8, (*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V2, func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v9_9.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}(), v10_10, (*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V6, "a", v11_11, "fa", v12_12, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v13_13.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}())), func() gopurs_runtime.Value {
																		origVal := (*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V7
																		if origVal.Type != gopurs_runtime.TypeRecord1 {
																			return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V7, "nested"), "a", v14_14, "fa", v15_15, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																				arr := *(*[]gopurs_runtime.Value)(v16_16.UnsafePtr)
																				unboxed := make([]gopurs_runtime.Value, len(arr))
																				for i, v := range arr {
																					unboxed[i] = v
																				}
																				return unboxed
																			}()))})
																		}
																		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
																		clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V7, "nested"), "a", v14_14, "fa", v15_15, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																			arr := *(*[]gopurs_runtime.Value)(v16_16.UnsafePtr)
																			unboxed := make([]gopurs_runtime.Value, len(arr))
																			for i, v := range arr {
																				unboxed[i] = v
																			}
																			return unboxed
																		}()))
																		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
																	}()}))}
																})
															})
														})
													})
												})
											})
										})
									})
								}), v_arrayMap20)
							}
							return gopurs_runtime.Array(res_go_arrayMap20)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array((*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V3))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), (*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V4)), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V6, "a")}).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V6, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V7, "nested"), "a")}).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V7, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6().UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_106
					} else {

					}
				}
				{
					var __t_tag_105 gopurs_runtime.Value = Get_Main_m6()
					if __t_tag_105.Type == 9 && __t_tag_105.IntVal == 1168316772 {
						__t106 = func() gopurs_runtime.Value {
							arr_val_arrayMap12 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply5(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
									arr_val_arrayMap18 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "a")}).UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap18
									arr_go_arrayMap18 := (*[]gopurs_runtime.Value)(arr_val_arrayMap18.UnsafePtr)
									_ = arr_go_arrayMap18
									res_go_arrayMap18 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap18))
									_ = res_go_arrayMap18
									for i_arrayMap18, v_arrayMap18 := range *arr_go_arrayMap18 {
										res_go_arrayMap18[i_arrayMap18] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.RecordUpdate1(v1_8, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_8, "nested"), "a", v2_9, "fa", v3_10, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v4_11.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}())))
												})
											})
										}), v_arrayMap18)
									}
									return gopurs_runtime.Array(res_go_arrayMap18)
								}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
							})), (*Constructor_Main_M7)(Get_Main_m6().UnsafePtr).V0)
							_ = arr_val_arrayMap12
							arr_go_arrayMap12 := (*[]gopurs_runtime.Value)(arr_val_arrayMap12.UnsafePtr)
							_ = arr_go_arrayMap12
							res_go_arrayMap12 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap12))
							_ = res_go_arrayMap12
							for i_arrayMap12, v_arrayMap12 := range *arr_go_arrayMap12 {
								res_go_arrayMap12[i_arrayMap12] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, v1_8}))}
								}), v_arrayMap12)
							}
							return gopurs_runtime.Array(res_go_arrayMap12)
						}()
						goto end_branch_106
					} else {

					}
				}
				{
					__t106 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_106:
				_dollar___unused_7_88 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m6"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_12 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t97 bool
						{
							if x_11.Type == 9 && x_11.IntVal == 3852365315 {
								__t97 = (y_12.Type == 9 && y_12.IntVal == 3852365315)
								goto end_branch_97
							} else {

							}
						}
						{
							if x_11.Type == 9 && x_11.IntVal == 769986722 {
								__t97 = (y_12.Type == 9 && y_12.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_11.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_12.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_10_96.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_11.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_12.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_97
							} else {

							}
						}
						{
							if x_11.Type == 9 && x_11.IntVal == 2727978561 {
								__t97 = (y_12.Type == 9 && y_12.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_11.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_12.UnsafePtr).V0))
								goto end_branch_97
							} else {

							}
						}
						{
							if x_11.Type == 9 && x_11.IntVal == 1830062304 {
								__t97 = (y_12.Type == 9 && y_12.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_11.UnsafePtr).V0, (*Constructor_Main_M3)(y_12.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_97
							} else {

							}
						}
						{
							if x_11.Type == 9 && x_11.IntVal == 3190619783 {
								__t97 = (y_12.Type == 9 && y_12.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_11.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_12.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_11.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_12.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_11.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_12.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_11.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_12.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_11.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_12.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_10_96.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_11.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_12.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_97
							} else {

							}
						}
						{
							if x_11.Type == 9 && x_11.IntVal == 108241190 {
								__t97 = (y_12.Type == 9 && y_12.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_11.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_12.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_11.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_12.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_11.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_12.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_11.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_12.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_11.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_12.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_10_96.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_11.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_12.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_97
							} else {

							}
						}
						{
							if x_11.Type == 9 && x_11.IntVal == 2066233029 {
								__t97 = (y_12.Type == 9 && y_12.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_11.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_12.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_11.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_12.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_11.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_12.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_10_96.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_11.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_12.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_11.UnsafePtr).V4, (*Constructor_Main_M6)(y_12.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_11.UnsafePtr).V5, (*Constructor_Main_M6)(y_12.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_11.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_12.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_11.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_12.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_11.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_12.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_11.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_12.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_11.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_12.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_10_96.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_11.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_12.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_11.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_12.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_11.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_12.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_11.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_12.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_11.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_12.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_11.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_12.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_10_96.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_11.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_12.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_97
							} else {

							}
						}
						{
							__t97 = (x_11.Type == 9 && x_11.IntVal == 1168316772) && ((y_12.Type == 9 && y_12.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_9_95, "eq"), (*Constructor_Main_M7)(x_11.UnsafePtr).V0, (*Constructor_Main_M7)(y_12.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_97:
						return gopurs_runtime.Bool(__t97)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t106.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m6()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_7_88
				// TAST (Let): __local_var_8_108 -> gopurs_runtime.Value
				__local_var_8_108 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_9 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_10 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_9, "zArrayA"), gopurs_runtime.RecordGet(rb_10, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_8_108
				// TAST (Let): __local_var_9_110 -> gopurs_runtime.Value
				__local_var_9_110 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_10, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_11, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_8_108, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_10, rb_11).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_9_110
				// TAST (Let): __local_var_10_111 -> gopurs_runtime.Value
				__local_var_10_111 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_11 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_12 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_11, "fa"), gopurs_runtime.RecordGet(rb_12, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_9_110, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_11, rb_12).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_10_111
				// TAST (Let): __local_var_11_112 -> gopurs_runtime.Value
				__local_var_11_112 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_12 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_12, "fIgnore"), gopurs_runtime.RecordGet(rb_13, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_10_111, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_12, rb_13).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_11_112
				// TAST (Let): __local_var_12_113 -> gopurs_runtime.Value
				__local_var_12_113 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_13 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_14 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_13, "arrayIgnore"), gopurs_runtime.RecordGet(rb_14, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_11_112, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_13, rb_14).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_12_113
				// TAST (Let): __local_var_9_109 -> gopurs_runtime.Value
				__local_var_9_109 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_13 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_14 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_13, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_14, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_12_113, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_13, rb_14).IntVal) != (0)))
					})
				}))
				_ = __local_var_9_109
				// TAST (Let): __local_var_10_114 -> gopurs_runtime.Value
				__local_var_10_114 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_11 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_109, "eq"), gopurs_runtime.RecordGet(ra_10, "nested"), gopurs_runtime.RecordGet(rb_11, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_10_114
				// TAST (Let): eqArray5_11_115 -> *Constructor_Data_Eq_Eq
				eqArray5_11_115 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_11_115
				var __t125 gopurs_runtime.Value
				{
					var __t_tag_117 gopurs_runtime.Value = Get_Main_m7()
					if __t_tag_117.Type == 9 && __t_tag_117.IntVal == 3852365315 {
						__t125 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M0{nil}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_125
					} else {

					}
				}
				{
					var __t_tag_118 gopurs_runtime.Value = Get_Main_m7()
					if __t_tag_118.Type == 9 && __t_tag_118.IntVal == 769986722 {
						__t125 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap14 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Main_M1)(Get_Main_m7().UnsafePtr).V0}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap14
							arr_go_arrayMap14 := (*[]gopurs_runtime.Value)(arr_val_arrayMap14.UnsafePtr)
							_ = arr_go_arrayMap14
							res_go_arrayMap14 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap14))
							_ = res_go_arrayMap14
							for i_arrayMap14, v_arrayMap14 := range *arr_go_arrayMap14 {
								res_go_arrayMap14[i_arrayMap14] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, v2_9, func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v3_10.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()}))}
									})
								}), v_arrayMap14)
							}
							return gopurs_runtime.Array(res_go_arrayMap14)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array((*Constructor_Main_M1)(Get_Main_m7().UnsafePtr).V1)))
						goto end_branch_125
					} else {

					}
				}
				{
					var __t_tag_119 gopurs_runtime.Value = Get_Main_m7()
					if __t_tag_119.Type == 9 && __t_tag_119.IntVal == 2727978561 {
						__t125 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, (*Constructor_Main_M2)(Get_Main_m7().UnsafePtr).V0})}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_125
					} else {

					}
				}
				{
					var __t_tag_120 gopurs_runtime.Value = Get_Main_m7()
					if __t_tag_120.Type == 9 && __t_tag_120.IntVal == 1830062304 {
						__t125 = func() gopurs_runtime.Value {
							arr_val_arrayMap13 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), (*Constructor_Main_M3)(Get_Main_m7().UnsafePtr).V0)
							_ = arr_val_arrayMap13
							arr_go_arrayMap13 := (*[]gopurs_runtime.Value)(arr_val_arrayMap13.UnsafePtr)
							_ = arr_go_arrayMap13
							res_go_arrayMap13 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap13))
							_ = res_go_arrayMap13
							for i_arrayMap13, v_arrayMap13 := range *arr_go_arrayMap13 {
								res_go_arrayMap13[i_arrayMap13] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, v1_9}))}
								}), v_arrayMap13)
							}
							return gopurs_runtime.Array(res_go_arrayMap13)
						}()
						goto end_branch_125
					} else {

					}
				}
				{
					var __t_tag_121 gopurs_runtime.Value = Get_Main_m7()
					if __t_tag_121.Type == 9 && __t_tag_121.IntVal == 3190619783 {
						__t125 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap15 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m7().UnsafePtr).V0, "a")}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap15
							arr_go_arrayMap15 := (*[]gopurs_runtime.Value)(arr_val_arrayMap15.UnsafePtr)
							_ = arr_go_arrayMap15
							res_go_arrayMap15 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap15))
							_ = res_go_arrayMap15
							for i_arrayMap15, v_arrayMap15 := range *arr_go_arrayMap15 {
								res_go_arrayMap15[i_arrayMap15] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4)(Get_Main_m7().UnsafePtr).V0, "a", v1_9, "fa", v2_10, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_11.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}()))}))}
										})
									})
								}), v_arrayMap15)
							}
							return gopurs_runtime.Array(res_go_arrayMap15)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m7().UnsafePtr).V0, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m7().UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_125
					} else {

					}
				}
				{
					var __t_tag_122 gopurs_runtime.Value = Get_Main_m7()
					if __t_tag_122.Type == 9 && __t_tag_122.IntVal == 108241190 {
						__t125 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap15 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m7().UnsafePtr).V0, "nested"), "a")}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap15
							arr_go_arrayMap15 := (*[]gopurs_runtime.Value)(arr_val_arrayMap15.UnsafePtr)
							_ = arr_go_arrayMap15
							res_go_arrayMap15 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap15))
							_ = res_go_arrayMap15
							for i_arrayMap15, v_arrayMap15 := range *arr_go_arrayMap15 {
								res_go_arrayMap15[i_arrayMap15] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, func() gopurs_runtime.Value {
												origVal := (*Constructor_Main_M5)(Get_Main_m7().UnsafePtr).V0
												if origVal.Type != gopurs_runtime.TypeRecord1 {
													return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m7().UnsafePtr).V0, "nested"), "a", v1_9, "fa", v2_10, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v3_11.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}()))})
												}
												clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
												clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m7().UnsafePtr).V0, "nested"), "a", v1_9, "fa", v2_10, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v3_11.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}()))
												return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
											}()}))}
										})
									})
								}), v_arrayMap15)
							}
							return gopurs_runtime.Array(res_go_arrayMap15)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m7().UnsafePtr).V0, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m7().UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_125
					} else {

					}
				}
				{
					var __t_tag_123 gopurs_runtime.Value = Get_Main_m7()
					if __t_tag_123.Type == 9 && __t_tag_123.IntVal == 2066233029 {
						__t125 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap21 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V1}).UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap21
							arr_go_arrayMap21 := (*[]gopurs_runtime.Value)(arr_val_arrayMap21.UnsafePtr)
							_ = arr_go_arrayMap21
							res_go_arrayMap21 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap21))
							_ = res_go_arrayMap21
							for i_arrayMap21, v_arrayMap21 := range *arr_go_arrayMap21 {
								res_go_arrayMap21[i_arrayMap21] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v8_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v9_10 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v10_11 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v11_12 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v12_13 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v13_14 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Func(func(v14_15 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(v15_16 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(v16_17 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, (*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V0, v8_9, (*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V2, func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v9_10.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}(), v10_11, (*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V6, "a", v11_12, "fa", v12_13, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v13_14.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}())), func() gopurs_runtime.Value {
																		origVal := (*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V7
																		if origVal.Type != gopurs_runtime.TypeRecord1 {
																			return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V7, "nested"), "a", v14_15, "fa", v15_16, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																				arr := *(*[]gopurs_runtime.Value)(v16_17.UnsafePtr)
																				unboxed := make([]gopurs_runtime.Value, len(arr))
																				for i, v := range arr {
																					unboxed[i] = v
																				}
																				return unboxed
																			}()))})
																		}
																		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
																		clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V7, "nested"), "a", v14_15, "fa", v15_16, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																			arr := *(*[]gopurs_runtime.Value)(v16_17.UnsafePtr)
																			unboxed := make([]gopurs_runtime.Value, len(arr))
																			for i, v := range arr {
																				unboxed[i] = v
																			}
																			return unboxed
																		}()))
																		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
																	}()}))}
																})
															})
														})
													})
												})
											})
										})
									})
								}), v_arrayMap21)
							}
							return gopurs_runtime.Array(res_go_arrayMap21)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array((*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V3))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), (*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V4)), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V6, "a")}).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V6, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V7, "nested"), "a")}).UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V7, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7().UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_125
					} else {

					}
				}
				{
					var __t_tag_124 gopurs_runtime.Value = Get_Main_m7()
					if __t_tag_124.Type == 9 && __t_tag_124.IntVal == 1168316772 {
						__t125 = func() gopurs_runtime.Value {
							arr_val_arrayMap13 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply5(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
									arr_val_arrayMap19 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_9, "nested"), "a")}).UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap19
									arr_go_arrayMap19 := (*[]gopurs_runtime.Value)(arr_val_arrayMap19.UnsafePtr)
									_ = arr_go_arrayMap19
									res_go_arrayMap19 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap19))
									_ = res_go_arrayMap19
									for i_arrayMap19, v_arrayMap19 := range *arr_go_arrayMap19 {
										res_go_arrayMap19[i_arrayMap19] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v4_12 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.RecordUpdate1(v1_9, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_9, "nested"), "a", v2_10, "fa", v3_11, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v4_12.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}())))
												})
											})
										}), v_arrayMap19)
									}
									return gopurs_runtime.Array(res_go_arrayMap19)
								}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_9, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_9, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
							})), (*Constructor_Main_M7)(Get_Main_m7().UnsafePtr).V0)
							_ = arr_val_arrayMap13
							arr_go_arrayMap13 := (*[]gopurs_runtime.Value)(arr_val_arrayMap13.UnsafePtr)
							_ = arr_go_arrayMap13
							res_go_arrayMap13 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap13))
							_ = res_go_arrayMap13
							for i_arrayMap13, v_arrayMap13 := range *arr_go_arrayMap13 {
								res_go_arrayMap13[i_arrayMap13] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, v1_9}))}
								}), v_arrayMap13)
							}
							return gopurs_runtime.Array(res_go_arrayMap13)
						}()
						goto end_branch_125
					} else {

					}
				}
				{
					__t125 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_125:
				_dollar___unused_8_107 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("traverse - m7"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_13 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t116 bool
						{
							if x_12.Type == 9 && x_12.IntVal == 3852365315 {
								__t116 = (y_13.Type == 9 && y_13.IntVal == 3852365315)
								goto end_branch_116
							} else {

							}
						}
						{
							if x_12.Type == 9 && x_12.IntVal == 769986722 {
								__t116 = (y_13.Type == 9 && y_13.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_12.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_13.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_11_115.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_12.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_13.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_116
							} else {

							}
						}
						{
							if x_12.Type == 9 && x_12.IntVal == 2727978561 {
								__t116 = (y_13.Type == 9 && y_13.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_12.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_13.UnsafePtr).V0))
								goto end_branch_116
							} else {

							}
						}
						{
							if x_12.Type == 9 && x_12.IntVal == 1830062304 {
								__t116 = (y_13.Type == 9 && y_13.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_12.UnsafePtr).V0, (*Constructor_Main_M3)(y_13.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_116
							} else {

							}
						}
						{
							if x_12.Type == 9 && x_12.IntVal == 3190619783 {
								__t116 = (y_13.Type == 9 && y_13.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_12.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_13.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_12.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_13.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_12.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_13.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_12.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_13.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_12.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_13.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_11_115.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_12.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_13.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_116
							} else {

							}
						}
						{
							if x_12.Type == 9 && x_12.IntVal == 108241190 {
								__t116 = (y_13.Type == 9 && y_13.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_12.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_13.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_12.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_13.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_12.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_13.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_12.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_13.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_12.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_13.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_11_115.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_12.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_13.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_116
							} else {

							}
						}
						{
							if x_12.Type == 9 && x_12.IntVal == 2066233029 {
								__t116 = (y_13.Type == 9 && y_13.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_12.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_13.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_12.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_13.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_12.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_13.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_11_115.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_12.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_13.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_12.UnsafePtr).V4, (*Constructor_Main_M6)(y_13.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_12.UnsafePtr).V5, (*Constructor_Main_M6)(y_13.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_12.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_13.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_12.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_13.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_12.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_13.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_12.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_13.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_12.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_13.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_11_115.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_12.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_13.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_12.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_13.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_12.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_13.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_12.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_13.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_12.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_13.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_12.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_13.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_11_115.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_12.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_13.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_116
							} else {

							}
						}
						{
							__t116 = (x_12.Type == 9 && x_12.IntVal == 1168316772) && ((y_13.Type == 9 && y_13.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_10_114, "eq"), (*Constructor_Main_M7)(x_12.UnsafePtr).V0, (*Constructor_Main_M7)(y_13.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_116:
						return gopurs_runtime.Bool(__t116)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t125.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m7()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_8_107
				// TAST (Let): __local_var_9_127 -> gopurs_runtime.Value
				__local_var_9_127 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_10 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_10, "zArrayA"), gopurs_runtime.RecordGet(rb_11, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_9_127
				// TAST (Let): __local_var_10_129 -> gopurs_runtime.Value
				__local_var_10_129 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_11 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_12 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_11, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_12, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_9_127, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_11, rb_12).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_10_129
				// TAST (Let): __local_var_11_130 -> gopurs_runtime.Value
				__local_var_11_130 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_12 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_12, "fa"), gopurs_runtime.RecordGet(rb_13, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_10_129, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_12, rb_13).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_11_130
				// TAST (Let): __local_var_12_131 -> gopurs_runtime.Value
				__local_var_12_131 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_13 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_14 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_13, "fIgnore"), gopurs_runtime.RecordGet(rb_14, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_11_130, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_13, rb_14).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_12_131
				// TAST (Let): __local_var_13_132 -> gopurs_runtime.Value
				__local_var_13_132 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_14 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_14, "arrayIgnore"), gopurs_runtime.RecordGet(rb_15, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_12_131, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_14, rb_15).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_13_132
				// TAST (Let): __local_var_10_128 -> gopurs_runtime.Value
				__local_var_10_128 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_14 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_15 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_14, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_15, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_13_132, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_14, rb_15).IntVal) != (0)))
					})
				}))
				_ = __local_var_10_128
				// TAST (Let): __local_var_11_133 -> gopurs_runtime.Value
				__local_var_11_133 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_12 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_128, "eq"), gopurs_runtime.RecordGet(ra_11, "nested"), gopurs_runtime.RecordGet(rb_12, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_11_133
				// TAST (Let): eqArray5_12_134 -> *Constructor_Data_Eq_Eq
				eqArray5_12_134 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_12_134
				_dollar___unused_9_126 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m0"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_14 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t135 bool
						{
							if x_13.Type == 9 && x_13.IntVal == 3852365315 {
								__t135 = (y_14.Type == 9 && y_14.IntVal == 3852365315)
								goto end_branch_135
							} else {

							}
						}
						{
							if x_13.Type == 9 && x_13.IntVal == 769986722 {
								__t135 = (y_14.Type == 9 && y_14.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_13.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_14.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_12_134.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_13.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_14.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_135
							} else {

							}
						}
						{
							if x_13.Type == 9 && x_13.IntVal == 2727978561 {
								__t135 = (y_14.Type == 9 && y_14.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_13.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_14.UnsafePtr).V0))
								goto end_branch_135
							} else {

							}
						}
						{
							if x_13.Type == 9 && x_13.IntVal == 1830062304 {
								__t135 = (y_14.Type == 9 && y_14.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_13.UnsafePtr).V0, (*Constructor_Main_M3)(y_14.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_135
							} else {

							}
						}
						{
							if x_13.Type == 9 && x_13.IntVal == 3190619783 {
								__t135 = (y_14.Type == 9 && y_14.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_13.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_14.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_13.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_14.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_13.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_14.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_13.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_14.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_13.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_14.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_12_134.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_13.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_14.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_135
							} else {

							}
						}
						{
							if x_13.Type == 9 && x_13.IntVal == 108241190 {
								__t135 = (y_14.Type == 9 && y_14.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_13.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_14.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_13.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_14.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_13.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_14.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_13.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_14.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_13.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_14.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_12_134.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_13.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_14.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_135
							} else {

							}
						}
						{
							if x_13.Type == 9 && x_13.IntVal == 2066233029 {
								__t135 = (y_14.Type == 9 && y_14.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_13.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_14.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_13.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_14.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_13.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_14.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_12_134.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_13.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_14.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_13.UnsafePtr).V4, (*Constructor_Main_M6)(y_14.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_13.UnsafePtr).V5, (*Constructor_Main_M6)(y_14.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_13.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_14.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_13.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_14.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_13.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_14.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_13.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_14.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_13.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_14.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_12_134.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_13.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_14.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_13.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_14.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_13.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_14.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_13.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_14.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_13.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_14.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_13.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_14.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_12_134.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_13.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_14.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_135
							} else {

							}
						}
						{
							__t135 = (x_13.Type == 9 && x_13.IntVal == 1168316772) && ((y_14.Type == 9 && y_14.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_11_133, "eq"), (*Constructor_Main_M7)(x_13.UnsafePtr).V0, (*Constructor_Main_M7)(y_14.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_135:
						return gopurs_runtime.Bool(__t135)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
							arr := []*Constructor_Main_M0{nil}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
							}
							return gopurs_runtime.Array(boxed)
						}().UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr := []*Constructor_Main_M0{nil}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_9_126
				// TAST (Let): __local_var_10_137 -> gopurs_runtime.Value
				__local_var_10_137 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_11 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_12 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_11, "zArrayA"), gopurs_runtime.RecordGet(rb_12, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_10_137
				// TAST (Let): __local_var_11_139 -> gopurs_runtime.Value
				__local_var_11_139 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_12 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_12, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_13, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_10_137, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_12, rb_13).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_11_139
				// TAST (Let): __local_var_12_140 -> gopurs_runtime.Value
				__local_var_12_140 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_13 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_14 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_13, "fa"), gopurs_runtime.RecordGet(rb_14, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_11_139, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_13, rb_14).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_12_140
				// TAST (Let): __local_var_13_141 -> gopurs_runtime.Value
				__local_var_13_141 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_14 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_14, "fIgnore"), gopurs_runtime.RecordGet(rb_15, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_12_140, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_14, rb_15).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_13_141
				// TAST (Let): __local_var_14_142 -> gopurs_runtime.Value
				__local_var_14_142 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_15 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_15, "arrayIgnore"), gopurs_runtime.RecordGet(rb_16, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_13_141, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_15, rb_16).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_14_142
				// TAST (Let): __local_var_11_138 -> gopurs_runtime.Value
				__local_var_11_138 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_15 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_16 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_15, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_16, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_14_142, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_15, rb_16).IntVal) != (0)))
					})
				}))
				_ = __local_var_11_138
				// TAST (Let): __local_var_12_143 -> gopurs_runtime.Value
				__local_var_12_143 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_12 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_13 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_138, "eq"), gopurs_runtime.RecordGet(ra_12, "nested"), gopurs_runtime.RecordGet(rb_13, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_12_143
				// TAST (Let): eqArray5_13_144 -> *Constructor_Data_Eq_Eq
				eqArray5_13_144 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_13_144
				var __t154 gopurs_runtime.Value
				{
					var __t_tag_146 gopurs_runtime.Value = Get_Main_m1_prime_()
					if __t_tag_146.Type == 9 && __t_tag_146.IntVal == 3852365315 {
						__t154 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M0{nil}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_154
					} else {

					}
				}
				{
					var __t_tag_147 gopurs_runtime.Value = Get_Main_m1_prime_()
					if __t_tag_147.Type == 9 && __t_tag_147.IntVal == 769986722 {
						__t154 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap16 := (*Constructor_Main_M1)(Get_Main_m1_prime_().UnsafePtr).V0
							_ = arr_val_arrayMap16
							arr_go_arrayMap16 := (*[]gopurs_runtime.Value)(arr_val_arrayMap16.UnsafePtr)
							_ = arr_go_arrayMap16
							res_go_arrayMap16 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap16))
							_ = res_go_arrayMap16
							for i_arrayMap16, v_arrayMap16 := range *arr_go_arrayMap16 {
								res_go_arrayMap16[i_arrayMap16] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, v2_11, func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v3_12.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()}))}
									})
								}), v_arrayMap16)
							}
							return gopurs_runtime.Array(res_go_arrayMap16)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_11
						}), gopurs_runtime.Array((*Constructor_Main_M1)(Get_Main_m1_prime_().UnsafePtr).V1)))
						goto end_branch_154
					} else {

					}
				}
				{
					var __t_tag_148 gopurs_runtime.Value = Get_Main_m1_prime_()
					if __t_tag_148.Type == 9 && __t_tag_148.IntVal == 2727978561 {
						__t154 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, (*Constructor_Main_M2)(Get_Main_m1_prime_().UnsafePtr).V0})}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_154
					} else {

					}
				}
				{
					var __t_tag_149 gopurs_runtime.Value = Get_Main_m1_prime_()
					if __t_tag_149.Type == 9 && __t_tag_149.IntVal == 1830062304 {
						__t154 = func() gopurs_runtime.Value {
							arr_val_arrayMap15 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return x_11
							}), (*Constructor_Main_M3)(Get_Main_m1_prime_().UnsafePtr).V0)
							_ = arr_val_arrayMap15
							arr_go_arrayMap15 := (*[]gopurs_runtime.Value)(arr_val_arrayMap15.UnsafePtr)
							_ = arr_go_arrayMap15
							res_go_arrayMap15 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap15))
							_ = res_go_arrayMap15
							for i_arrayMap15, v_arrayMap15 := range *arr_go_arrayMap15 {
								res_go_arrayMap15[i_arrayMap15] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, v1_11}))}
								}), v_arrayMap15)
							}
							return gopurs_runtime.Array(res_go_arrayMap15)
						}()
						goto end_branch_154
					} else {

					}
				}
				{
					var __t_tag_150 gopurs_runtime.Value = Get_Main_m1_prime_()
					if __t_tag_150.Type == 9 && __t_tag_150.IntVal == 3190619783 {
						__t154 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap17 := gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m1_prime_().UnsafePtr).V0, "a")
							_ = arr_val_arrayMap17
							arr_go_arrayMap17 := (*[]gopurs_runtime.Value)(arr_val_arrayMap17.UnsafePtr)
							_ = arr_go_arrayMap17
							res_go_arrayMap17 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap17))
							_ = res_go_arrayMap17
							for i_arrayMap17, v_arrayMap17 := range *arr_go_arrayMap17 {
								res_go_arrayMap17[i_arrayMap17] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4)(Get_Main_m1_prime_().UnsafePtr).V0, "a", v1_11, "fa", v2_12, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_13.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}()))}))}
										})
									})
								}), v_arrayMap17)
							}
							return gopurs_runtime.Array(res_go_arrayMap17)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_11
						}), gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m1_prime_().UnsafePtr).V0, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_11
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m1_prime_().UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_154
					} else {

					}
				}
				{
					var __t_tag_151 gopurs_runtime.Value = Get_Main_m1_prime_()
					if __t_tag_151.Type == 9 && __t_tag_151.IntVal == 108241190 {
						__t154 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap17 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m1_prime_().UnsafePtr).V0, "nested"), "a")
							_ = arr_val_arrayMap17
							arr_go_arrayMap17 := (*[]gopurs_runtime.Value)(arr_val_arrayMap17.UnsafePtr)
							_ = arr_go_arrayMap17
							res_go_arrayMap17 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap17))
							_ = res_go_arrayMap17
							for i_arrayMap17, v_arrayMap17 := range *arr_go_arrayMap17 {
								res_go_arrayMap17[i_arrayMap17] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, func() gopurs_runtime.Value {
												origVal := (*Constructor_Main_M5)(Get_Main_m1_prime_().UnsafePtr).V0
												if origVal.Type != gopurs_runtime.TypeRecord1 {
													return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m1_prime_().UnsafePtr).V0, "nested"), "a", v1_11, "fa", v2_12, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v3_13.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}()))})
												}
												clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
												clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m1_prime_().UnsafePtr).V0, "nested"), "a", v1_11, "fa", v2_12, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v3_13.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}()))
												return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
											}()}))}
										})
									})
								}), v_arrayMap17)
							}
							return gopurs_runtime.Array(res_go_arrayMap17)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_11
						}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m1_prime_().UnsafePtr).V0, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_11
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m1_prime_().UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_154
					} else {

					}
				}
				{
					var __t_tag_152 gopurs_runtime.Value = Get_Main_m1_prime_()
					if __t_tag_152.Type == 9 && __t_tag_152.IntVal == 2066233029 {
						__t154 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap23 := (*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V1
							_ = arr_val_arrayMap23
							arr_go_arrayMap23 := (*[]gopurs_runtime.Value)(arr_val_arrayMap23.UnsafePtr)
							_ = arr_go_arrayMap23
							res_go_arrayMap23 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap23))
							_ = res_go_arrayMap23
							for i_arrayMap23, v_arrayMap23 := range *arr_go_arrayMap23 {
								res_go_arrayMap23[i_arrayMap23] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v8_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v9_12 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v10_13 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v11_14 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v12_15 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v13_16 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Func(func(v14_17 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(v15_18 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(v16_19 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, (*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V0, v8_11, (*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V2, func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v9_12.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}(), v10_13, (*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V6, "a", v11_14, "fa", v12_15, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v13_16.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}())), func() gopurs_runtime.Value {
																		origVal := (*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V7
																		if origVal.Type != gopurs_runtime.TypeRecord1 {
																			return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V7, "nested"), "a", v14_17, "fa", v15_18, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																				arr := *(*[]gopurs_runtime.Value)(v16_19.UnsafePtr)
																				unboxed := make([]gopurs_runtime.Value, len(arr))
																				for i, v := range arr {
																					unboxed[i] = v
																				}
																				return unboxed
																			}()))})
																		}
																		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
																		clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V7, "nested"), "a", v14_17, "fa", v15_18, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																			arr := *(*[]gopurs_runtime.Value)(v16_19.UnsafePtr)
																			unboxed := make([]gopurs_runtime.Value, len(arr))
																			for i, v := range arr {
																				unboxed[i] = v
																			}
																			return unboxed
																		}()))
																		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
																	}()}))}
																})
															})
														})
													})
												})
											})
										})
									})
								}), v_arrayMap23)
							}
							return gopurs_runtime.Array(res_go_arrayMap23)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_11
						}), gopurs_runtime.Array((*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V3))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_11
						}), (*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V4)), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V6, "a")), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_11
						}), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V6, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_11
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V7, "nested"), "a")), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_11
						}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V7, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_11
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m1_prime_().UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_154
					} else {

					}
				}
				{
					var __t_tag_153 gopurs_runtime.Value = Get_Main_m1_prime_()
					if __t_tag_153.Type == 9 && __t_tag_153.IntVal == 1168316772 {
						__t154 = func() gopurs_runtime.Value {
							arr_val_arrayMap15 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply5(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
									arr_val_arrayMap21 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_11, "nested"), "a")
									_ = arr_val_arrayMap21
									arr_go_arrayMap21 := (*[]gopurs_runtime.Value)(arr_val_arrayMap21.UnsafePtr)
									_ = arr_go_arrayMap21
									res_go_arrayMap21 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap21))
									_ = res_go_arrayMap21
									for i_arrayMap21, v_arrayMap21 := range *arr_go_arrayMap21 {
										res_go_arrayMap21[i_arrayMap21] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v4_14 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.RecordUpdate1(v1_11, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_11, "nested"), "a", v2_12, "fa", v3_13, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v4_14.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}())))
												})
											})
										}), v_arrayMap21)
									}
									return gopurs_runtime.Array(res_go_arrayMap21)
								}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
									return x_12
								}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_11, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
									return x_12
								}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_11, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
							})), (*Constructor_Main_M7)(Get_Main_m1_prime_().UnsafePtr).V0)
							_ = arr_val_arrayMap15
							arr_go_arrayMap15 := (*[]gopurs_runtime.Value)(arr_val_arrayMap15.UnsafePtr)
							_ = arr_go_arrayMap15
							res_go_arrayMap15 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap15))
							_ = res_go_arrayMap15
							for i_arrayMap15, v_arrayMap15 := range *arr_go_arrayMap15 {
								res_go_arrayMap15[i_arrayMap15] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, v1_11}))}
								}), v_arrayMap15)
							}
							return gopurs_runtime.Array(res_go_arrayMap15)
						}()
						goto end_branch_154
					} else {

					}
				}
				{
					__t154 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_154:
				_dollar___unused_10_136 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m1"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_15 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t145 bool
						{
							if x_14.Type == 9 && x_14.IntVal == 3852365315 {
								__t145 = (y_15.Type == 9 && y_15.IntVal == 3852365315)
								goto end_branch_145
							} else {

							}
						}
						{
							if x_14.Type == 9 && x_14.IntVal == 769986722 {
								__t145 = (y_15.Type == 9 && y_15.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_14.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_15.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_13_144.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_14.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_15.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_145
							} else {

							}
						}
						{
							if x_14.Type == 9 && x_14.IntVal == 2727978561 {
								__t145 = (y_15.Type == 9 && y_15.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_14.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_15.UnsafePtr).V0))
								goto end_branch_145
							} else {

							}
						}
						{
							if x_14.Type == 9 && x_14.IntVal == 1830062304 {
								__t145 = (y_15.Type == 9 && y_15.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_14.UnsafePtr).V0, (*Constructor_Main_M3)(y_15.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_145
							} else {

							}
						}
						{
							if x_14.Type == 9 && x_14.IntVal == 3190619783 {
								__t145 = (y_15.Type == 9 && y_15.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_14.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_15.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_14.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_15.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_14.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_15.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_14.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_15.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_14.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_15.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_13_144.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_14.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_15.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_145
							} else {

							}
						}
						{
							if x_14.Type == 9 && x_14.IntVal == 108241190 {
								__t145 = (y_15.Type == 9 && y_15.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_14.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_15.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_14.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_15.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_14.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_15.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_14.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_15.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_14.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_15.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_13_144.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_14.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_15.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_145
							} else {

							}
						}
						{
							if x_14.Type == 9 && x_14.IntVal == 2066233029 {
								__t145 = (y_15.Type == 9 && y_15.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_14.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_15.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_14.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_15.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_14.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_15.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_13_144.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_14.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_15.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_14.UnsafePtr).V4, (*Constructor_Main_M6)(y_15.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_14.UnsafePtr).V5, (*Constructor_Main_M6)(y_15.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_14.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_15.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_14.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_15.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_14.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_15.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_14.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_15.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_14.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_15.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_13_144.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_14.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_15.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_14.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_15.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_14.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_15.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_14.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_15.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_14.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_15.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_14.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_15.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_13_144.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_14.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_15.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_145
							} else {

							}
						}
						{
							__t145 = (x_14.Type == 9 && x_14.IntVal == 1168316772) && ((y_15.Type == 9 && y_15.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_12_143, "eq"), (*Constructor_Main_M7)(x_14.UnsafePtr).V0, (*Constructor_Main_M7)(y_15.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_145:
						return gopurs_runtime.Bool(__t145)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t154.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m1()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_10_136
				// TAST (Let): __local_var_11_156 -> gopurs_runtime.Value
				__local_var_11_156 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_12 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_12, "zArrayA"), gopurs_runtime.RecordGet(rb_13, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_11_156
				// TAST (Let): __local_var_12_158 -> gopurs_runtime.Value
				__local_var_12_158 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_13 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_14 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_13, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_14, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_11_156, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_13, rb_14).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_12_158
				// TAST (Let): __local_var_13_159 -> gopurs_runtime.Value
				__local_var_13_159 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_14 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_14, "fa"), gopurs_runtime.RecordGet(rb_15, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_12_158, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_14, rb_15).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_13_159
				// TAST (Let): __local_var_14_160 -> gopurs_runtime.Value
				__local_var_14_160 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_15 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_15, "fIgnore"), gopurs_runtime.RecordGet(rb_16, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_13_159, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_15, rb_16).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_14_160
				// TAST (Let): __local_var_15_161 -> gopurs_runtime.Value
				__local_var_15_161 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_16 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_16, "arrayIgnore"), gopurs_runtime.RecordGet(rb_17, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_14_160, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_16, rb_17).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_15_161
				// TAST (Let): __local_var_12_157 -> gopurs_runtime.Value
				__local_var_12_157 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_16 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_17 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_16, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_17, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_15_161, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_16, rb_17).IntVal) != (0)))
					})
				}))
				_ = __local_var_12_157
				// TAST (Let): __local_var_13_162 -> gopurs_runtime.Value
				__local_var_13_162 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_13 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_14 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_12_157, "eq"), gopurs_runtime.RecordGet(ra_13, "nested"), gopurs_runtime.RecordGet(rb_14, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_13_162
				// TAST (Let): eqArray5_14_163 -> *Constructor_Data_Eq_Eq
				eqArray5_14_163 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_14_163
				_dollar___unused_11_155 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m2"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_16 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t164 bool
						{
							if x_15.Type == 9 && x_15.IntVal == 3852365315 {
								__t164 = (y_16.Type == 9 && y_16.IntVal == 3852365315)
								goto end_branch_164
							} else {

							}
						}
						{
							if x_15.Type == 9 && x_15.IntVal == 769986722 {
								__t164 = (y_16.Type == 9 && y_16.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_15.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_16.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_14_163.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_15.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_16.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_164
							} else {

							}
						}
						{
							if x_15.Type == 9 && x_15.IntVal == 2727978561 {
								__t164 = (y_16.Type == 9 && y_16.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_15.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_16.UnsafePtr).V0))
								goto end_branch_164
							} else {

							}
						}
						{
							if x_15.Type == 9 && x_15.IntVal == 1830062304 {
								__t164 = (y_16.Type == 9 && y_16.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_15.UnsafePtr).V0, (*Constructor_Main_M3)(y_16.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_164
							} else {

							}
						}
						{
							if x_15.Type == 9 && x_15.IntVal == 3190619783 {
								__t164 = (y_16.Type == 9 && y_16.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_15.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_16.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_15.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_16.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_15.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_16.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_15.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_16.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_15.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_16.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_14_163.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_15.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_16.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_164
							} else {

							}
						}
						{
							if x_15.Type == 9 && x_15.IntVal == 108241190 {
								__t164 = (y_16.Type == 9 && y_16.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_15.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_16.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_15.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_16.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_15.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_16.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_15.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_16.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_15.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_16.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_14_163.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_15.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_16.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_164
							} else {

							}
						}
						{
							if x_15.Type == 9 && x_15.IntVal == 2066233029 {
								__t164 = (y_16.Type == 9 && y_16.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_15.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_16.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_15.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_16.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_15.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_16.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_14_163.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_15.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_16.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_15.UnsafePtr).V4, (*Constructor_Main_M6)(y_16.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_15.UnsafePtr).V5, (*Constructor_Main_M6)(y_16.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_15.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_16.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_15.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_16.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_15.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_16.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_15.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_16.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_15.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_16.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_14_163.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_15.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_16.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_15.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_16.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_15.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_16.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_15.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_16.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_15.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_16.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_15.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_16.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_14_163.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_15.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_16.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_164
							} else {

							}
						}
						{
							__t164 = (x_15.Type == 9 && x_15.IntVal == 1168316772) && ((y_16.Type == 9 && y_16.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_13_162, "eq"), (*Constructor_Main_M7)(x_15.UnsafePtr).V0, (*Constructor_Main_M7)(y_16.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_164:
						return gopurs_runtime.Bool(__t164)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
							arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, 0})}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
							}
							return gopurs_runtime.Array(boxed)
						}().UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, 0})}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_11_155
				// TAST (Let): __local_var_12_166 -> gopurs_runtime.Value
				__local_var_12_166 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_13 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_14 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_13, "zArrayA"), gopurs_runtime.RecordGet(rb_14, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_12_166
				// TAST (Let): __local_var_13_168 -> gopurs_runtime.Value
				__local_var_13_168 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_14 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_14, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_15, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_12_166, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_14, rb_15).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_13_168
				// TAST (Let): __local_var_14_169 -> gopurs_runtime.Value
				__local_var_14_169 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_15 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_15, "fa"), gopurs_runtime.RecordGet(rb_16, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_13_168, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_15, rb_16).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_14_169
				// TAST (Let): __local_var_15_170 -> gopurs_runtime.Value
				__local_var_15_170 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_16 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_16, "fIgnore"), gopurs_runtime.RecordGet(rb_17, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_14_169, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_16, rb_17).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_15_170
				// TAST (Let): __local_var_16_171 -> gopurs_runtime.Value
				__local_var_16_171 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_17 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_18 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_17, "arrayIgnore"), gopurs_runtime.RecordGet(rb_18, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_15_170, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_17, rb_18).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_16_171
				// TAST (Let): __local_var_13_167 -> gopurs_runtime.Value
				__local_var_13_167 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_17 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_18 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_17, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_18, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_16_171, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_17, rb_18).IntVal) != (0)))
					})
				}))
				_ = __local_var_13_167
				// TAST (Let): __local_var_14_172 -> gopurs_runtime.Value
				__local_var_14_172 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_14 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_15 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_167, "eq"), gopurs_runtime.RecordGet(ra_14, "nested"), gopurs_runtime.RecordGet(rb_15, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_14_172
				// TAST (Let): eqArray5_15_173 -> *Constructor_Data_Eq_Eq
				eqArray5_15_173 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_15_173
				var __t183 gopurs_runtime.Value
				{
					var __t_tag_175 gopurs_runtime.Value = Get_Main_m3_prime_()
					if __t_tag_175.Type == 9 && __t_tag_175.IntVal == 3852365315 {
						__t183 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M0{nil}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_183
					} else {

					}
				}
				{
					var __t_tag_176 gopurs_runtime.Value = Get_Main_m3_prime_()
					if __t_tag_176.Type == 9 && __t_tag_176.IntVal == 769986722 {
						__t183 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap18 := (*Constructor_Main_M1)(Get_Main_m3_prime_().UnsafePtr).V0
							_ = arr_val_arrayMap18
							arr_go_arrayMap18 := (*[]gopurs_runtime.Value)(arr_val_arrayMap18.UnsafePtr)
							_ = arr_go_arrayMap18
							res_go_arrayMap18 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap18))
							_ = res_go_arrayMap18
							for i_arrayMap18, v_arrayMap18 := range *arr_go_arrayMap18 {
								res_go_arrayMap18[i_arrayMap18] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, v2_13, func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v3_14.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()}))}
									})
								}), v_arrayMap18)
							}
							return gopurs_runtime.Array(res_go_arrayMap18)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_13
						}), gopurs_runtime.Array((*Constructor_Main_M1)(Get_Main_m3_prime_().UnsafePtr).V1)))
						goto end_branch_183
					} else {

					}
				}
				{
					var __t_tag_177 gopurs_runtime.Value = Get_Main_m3_prime_()
					if __t_tag_177.Type == 9 && __t_tag_177.IntVal == 2727978561 {
						__t183 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, (*Constructor_Main_M2)(Get_Main_m3_prime_().UnsafePtr).V0})}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_183
					} else {

					}
				}
				{
					var __t_tag_178 gopurs_runtime.Value = Get_Main_m3_prime_()
					if __t_tag_178.Type == 9 && __t_tag_178.IntVal == 1830062304 {
						__t183 = func() gopurs_runtime.Value {
							arr_val_arrayMap17 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
								return x_13
							}), (*Constructor_Main_M3)(Get_Main_m3_prime_().UnsafePtr).V0)
							_ = arr_val_arrayMap17
							arr_go_arrayMap17 := (*[]gopurs_runtime.Value)(arr_val_arrayMap17.UnsafePtr)
							_ = arr_go_arrayMap17
							res_go_arrayMap17 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap17))
							_ = res_go_arrayMap17
							for i_arrayMap17, v_arrayMap17 := range *arr_go_arrayMap17 {
								res_go_arrayMap17[i_arrayMap17] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, v1_13}))}
								}), v_arrayMap17)
							}
							return gopurs_runtime.Array(res_go_arrayMap17)
						}()
						goto end_branch_183
					} else {

					}
				}
				{
					var __t_tag_179 gopurs_runtime.Value = Get_Main_m3_prime_()
					if __t_tag_179.Type == 9 && __t_tag_179.IntVal == 3190619783 {
						__t183 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap19 := gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m3_prime_().UnsafePtr).V0, "a")
							_ = arr_val_arrayMap19
							arr_go_arrayMap19 := (*[]gopurs_runtime.Value)(arr_val_arrayMap19.UnsafePtr)
							_ = arr_go_arrayMap19
							res_go_arrayMap19 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap19))
							_ = res_go_arrayMap19
							for i_arrayMap19, v_arrayMap19 := range *arr_go_arrayMap19 {
								res_go_arrayMap19[i_arrayMap19] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4)(Get_Main_m3_prime_().UnsafePtr).V0, "a", v1_13, "fa", v2_14, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_15.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}()))}))}
										})
									})
								}), v_arrayMap19)
							}
							return gopurs_runtime.Array(res_go_arrayMap19)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_13
						}), gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m3_prime_().UnsafePtr).V0, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_13
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m3_prime_().UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_183
					} else {

					}
				}
				{
					var __t_tag_180 gopurs_runtime.Value = Get_Main_m3_prime_()
					if __t_tag_180.Type == 9 && __t_tag_180.IntVal == 108241190 {
						__t183 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap19 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m3_prime_().UnsafePtr).V0, "nested"), "a")
							_ = arr_val_arrayMap19
							arr_go_arrayMap19 := (*[]gopurs_runtime.Value)(arr_val_arrayMap19.UnsafePtr)
							_ = arr_go_arrayMap19
							res_go_arrayMap19 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap19))
							_ = res_go_arrayMap19
							for i_arrayMap19, v_arrayMap19 := range *arr_go_arrayMap19 {
								res_go_arrayMap19[i_arrayMap19] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, func() gopurs_runtime.Value {
												origVal := (*Constructor_Main_M5)(Get_Main_m3_prime_().UnsafePtr).V0
												if origVal.Type != gopurs_runtime.TypeRecord1 {
													return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m3_prime_().UnsafePtr).V0, "nested"), "a", v1_13, "fa", v2_14, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v3_15.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}()))})
												}
												clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
												clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m3_prime_().UnsafePtr).V0, "nested"), "a", v1_13, "fa", v2_14, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v3_15.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}()))
												return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
											}()}))}
										})
									})
								}), v_arrayMap19)
							}
							return gopurs_runtime.Array(res_go_arrayMap19)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_13
						}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m3_prime_().UnsafePtr).V0, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_13
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m3_prime_().UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_183
					} else {

					}
				}
				{
					var __t_tag_181 gopurs_runtime.Value = Get_Main_m3_prime_()
					if __t_tag_181.Type == 9 && __t_tag_181.IntVal == 2066233029 {
						__t183 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap25 := (*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V1
							_ = arr_val_arrayMap25
							arr_go_arrayMap25 := (*[]gopurs_runtime.Value)(arr_val_arrayMap25.UnsafePtr)
							_ = arr_go_arrayMap25
							res_go_arrayMap25 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap25))
							_ = res_go_arrayMap25
							for i_arrayMap25, v_arrayMap25 := range *arr_go_arrayMap25 {
								res_go_arrayMap25[i_arrayMap25] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v8_13 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v9_14 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v10_15 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v11_16 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v12_17 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v13_18 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Func(func(v14_19 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(v15_20 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(v16_21 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, (*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V0, v8_13, (*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V2, func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v9_14.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}(), v10_15, (*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V6, "a", v11_16, "fa", v12_17, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v13_18.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}())), func() gopurs_runtime.Value {
																		origVal := (*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V7
																		if origVal.Type != gopurs_runtime.TypeRecord1 {
																			return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V7, "nested"), "a", v14_19, "fa", v15_20, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																				arr := *(*[]gopurs_runtime.Value)(v16_21.UnsafePtr)
																				unboxed := make([]gopurs_runtime.Value, len(arr))
																				for i, v := range arr {
																					unboxed[i] = v
																				}
																				return unboxed
																			}()))})
																		}
																		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
																		clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V7, "nested"), "a", v14_19, "fa", v15_20, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																			arr := *(*[]gopurs_runtime.Value)(v16_21.UnsafePtr)
																			unboxed := make([]gopurs_runtime.Value, len(arr))
																			for i, v := range arr {
																				unboxed[i] = v
																			}
																			return unboxed
																		}()))
																		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
																	}()}))}
																})
															})
														})
													})
												})
											})
										})
									})
								}), v_arrayMap25)
							}
							return gopurs_runtime.Array(res_go_arrayMap25)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_13
						}), gopurs_runtime.Array((*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V3))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_13
						}), (*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V4)), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V6, "a")), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_13
						}), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V6, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_13
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V7, "nested"), "a")), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_13
						}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V7, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_13
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m3_prime_().UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_183
					} else {

					}
				}
				{
					var __t_tag_182 gopurs_runtime.Value = Get_Main_m3_prime_()
					if __t_tag_182.Type == 9 && __t_tag_182.IntVal == 1168316772 {
						__t183 = func() gopurs_runtime.Value {
							arr_val_arrayMap17 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply5(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
									arr_val_arrayMap23 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_13, "nested"), "a")
									_ = arr_val_arrayMap23
									arr_go_arrayMap23 := (*[]gopurs_runtime.Value)(arr_val_arrayMap23.UnsafePtr)
									_ = arr_go_arrayMap23
									res_go_arrayMap23 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap23))
									_ = res_go_arrayMap23
									for i_arrayMap23, v_arrayMap23 := range *arr_go_arrayMap23 {
										res_go_arrayMap23[i_arrayMap23] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v4_16 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.RecordUpdate1(v1_13, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_13, "nested"), "a", v2_14, "fa", v3_15, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v4_16.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}())))
												})
											})
										}), v_arrayMap23)
									}
									return gopurs_runtime.Array(res_go_arrayMap23)
								}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
									return x_14
								}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_13, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
									return x_14
								}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_13, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
							})), (*Constructor_Main_M7)(Get_Main_m3_prime_().UnsafePtr).V0)
							_ = arr_val_arrayMap17
							arr_go_arrayMap17 := (*[]gopurs_runtime.Value)(arr_val_arrayMap17.UnsafePtr)
							_ = arr_go_arrayMap17
							res_go_arrayMap17 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap17))
							_ = res_go_arrayMap17
							for i_arrayMap17, v_arrayMap17 := range *arr_go_arrayMap17 {
								res_go_arrayMap17[i_arrayMap17] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, v1_13}))}
								}), v_arrayMap17)
							}
							return gopurs_runtime.Array(res_go_arrayMap17)
						}()
						goto end_branch_183
					} else {

					}
				}
				{
					__t183 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_183:
				_dollar___unused_12_165 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m3"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_17 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t174 bool
						{
							if x_16.Type == 9 && x_16.IntVal == 3852365315 {
								__t174 = (y_17.Type == 9 && y_17.IntVal == 3852365315)
								goto end_branch_174
							} else {

							}
						}
						{
							if x_16.Type == 9 && x_16.IntVal == 769986722 {
								__t174 = (y_17.Type == 9 && y_17.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_16.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_17.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_15_173.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_16.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_17.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_174
							} else {

							}
						}
						{
							if x_16.Type == 9 && x_16.IntVal == 2727978561 {
								__t174 = (y_17.Type == 9 && y_17.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_16.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_17.UnsafePtr).V0))
								goto end_branch_174
							} else {

							}
						}
						{
							if x_16.Type == 9 && x_16.IntVal == 1830062304 {
								__t174 = (y_17.Type == 9 && y_17.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_16.UnsafePtr).V0, (*Constructor_Main_M3)(y_17.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_174
							} else {

							}
						}
						{
							if x_16.Type == 9 && x_16.IntVal == 3190619783 {
								__t174 = (y_17.Type == 9 && y_17.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_16.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_17.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_16.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_17.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_16.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_17.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_16.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_17.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_16.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_17.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_15_173.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_16.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_17.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_174
							} else {

							}
						}
						{
							if x_16.Type == 9 && x_16.IntVal == 108241190 {
								__t174 = (y_17.Type == 9 && y_17.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_16.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_17.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_16.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_17.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_16.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_17.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_16.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_17.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_16.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_17.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_15_173.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_16.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_17.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_174
							} else {

							}
						}
						{
							if x_16.Type == 9 && x_16.IntVal == 2066233029 {
								__t174 = (y_17.Type == 9 && y_17.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_16.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_17.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_16.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_17.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_16.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_17.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_15_173.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_16.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_17.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_16.UnsafePtr).V4, (*Constructor_Main_M6)(y_17.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_16.UnsafePtr).V5, (*Constructor_Main_M6)(y_17.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_16.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_17.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_16.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_17.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_16.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_17.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_16.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_17.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_16.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_17.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_15_173.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_16.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_17.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_16.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_17.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_16.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_17.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_16.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_17.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_16.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_17.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_16.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_17.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_15_173.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_16.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_17.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_174
							} else {

							}
						}
						{
							__t174 = (x_16.Type == 9 && x_16.IntVal == 1168316772) && ((y_17.Type == 9 && y_17.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_14_172, "eq"), (*Constructor_Main_M7)(x_16.UnsafePtr).V0, (*Constructor_Main_M7)(y_17.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_174:
						return gopurs_runtime.Bool(__t174)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t183.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m3()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_12_165
				// TAST (Let): __local_var_13_185 -> gopurs_runtime.Value
				__local_var_13_185 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_14 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_14, "zArrayA"), gopurs_runtime.RecordGet(rb_15, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_13_185
				// TAST (Let): __local_var_14_187 -> gopurs_runtime.Value
				__local_var_14_187 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_15 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_15, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_16, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_13_185, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_15, rb_16).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_14_187
				// TAST (Let): __local_var_15_188 -> gopurs_runtime.Value
				__local_var_15_188 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_16 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_16, "fa"), gopurs_runtime.RecordGet(rb_17, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_14_187, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_16, rb_17).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_15_188
				// TAST (Let): __local_var_16_189 -> gopurs_runtime.Value
				__local_var_16_189 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_17 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_18 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_17, "fIgnore"), gopurs_runtime.RecordGet(rb_18, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_15_188, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_17, rb_18).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_16_189
				// TAST (Let): __local_var_17_190 -> gopurs_runtime.Value
				__local_var_17_190 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_18 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_19 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_18, "arrayIgnore"), gopurs_runtime.RecordGet(rb_19, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_16_189, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_18, rb_19).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_17_190
				// TAST (Let): __local_var_14_186 -> gopurs_runtime.Value
				__local_var_14_186 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_18 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_19 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_18, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_19, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_17_190, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_18, rb_19).IntVal) != (0)))
					})
				}))
				_ = __local_var_14_186
				// TAST (Let): __local_var_15_191 -> gopurs_runtime.Value
				__local_var_15_191 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_15 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_16 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_186, "eq"), gopurs_runtime.RecordGet(ra_15, "nested"), gopurs_runtime.RecordGet(rb_16, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_15_191
				// TAST (Let): eqArray5_16_192 -> *Constructor_Data_Eq_Eq
				eqArray5_16_192 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_16_192
				_dollar___unused_13_184 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m4"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_18 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t193 bool
						{
							if x_17.Type == 9 && x_17.IntVal == 3852365315 {
								__t193 = (y_18.Type == 9 && y_18.IntVal == 3852365315)
								goto end_branch_193
							} else {

							}
						}
						{
							if x_17.Type == 9 && x_17.IntVal == 769986722 {
								__t193 = (y_18.Type == 9 && y_18.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_17.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_18.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_16_192.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_17.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_18.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_193
							} else {

							}
						}
						{
							if x_17.Type == 9 && x_17.IntVal == 2727978561 {
								__t193 = (y_18.Type == 9 && y_18.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_17.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_18.UnsafePtr).V0))
								goto end_branch_193
							} else {

							}
						}
						{
							if x_17.Type == 9 && x_17.IntVal == 1830062304 {
								__t193 = (y_18.Type == 9 && y_18.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_17.UnsafePtr).V0, (*Constructor_Main_M3)(y_18.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_193
							} else {

							}
						}
						{
							if x_17.Type == 9 && x_17.IntVal == 3190619783 {
								__t193 = (y_18.Type == 9 && y_18.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_17.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_18.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_17.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_18.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_17.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_18.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_17.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_18.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_17.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_18.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_16_192.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_17.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_18.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_193
							} else {

							}
						}
						{
							if x_17.Type == 9 && x_17.IntVal == 108241190 {
								__t193 = (y_18.Type == 9 && y_18.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_17.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_18.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_17.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_18.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_17.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_18.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_17.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_18.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_17.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_18.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_16_192.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_17.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_18.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_193
							} else {

							}
						}
						{
							if x_17.Type == 9 && x_17.IntVal == 2066233029 {
								__t193 = (y_18.Type == 9 && y_18.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_17.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_18.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_17.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_18.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_17.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_18.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_16_192.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_17.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_18.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_17.UnsafePtr).V4, (*Constructor_Main_M6)(y_18.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_17.UnsafePtr).V5, (*Constructor_Main_M6)(y_18.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_17.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_18.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_17.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_18.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_17.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_18.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_17.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_18.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_17.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_18.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_16_192.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_17.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_18.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_17.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_18.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_17.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_18.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_17.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_18.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_17.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_18.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_17.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_18.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_16_192.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_17.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_18.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_193
							} else {

							}
						}
						{
							__t193 = (x_17.Type == 9 && x_17.IntVal == 1168316772) && ((y_18.Type == 9 && y_18.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_15_191, "eq"), (*Constructor_Main_M7)(x_17.UnsafePtr).V0, (*Constructor_Main_M7)(y_18.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_193:
						return gopurs_runtime.Bool(__t193)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
						arr_val_arrayMap19 := gopurs_runtime.RecordGet(Get_Main_recordValue_prime_(), "a")
						_ = arr_val_arrayMap19
						arr_go_arrayMap19 := (*[]gopurs_runtime.Value)(arr_val_arrayMap19.UnsafePtr)
						_ = arr_go_arrayMap19
						res_go_arrayMap19 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap19))
						_ = res_go_arrayMap19
						for i_arrayMap19, v_arrayMap19 := range *arr_go_arrayMap19 {
							res_go_arrayMap19[i_arrayMap19] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, func() gopurs_runtime.Value {
											origVal := Get_Main_recordValue_prime_()
											if origVal.Type != gopurs_runtime.TypeRecordData {
												return gopurs_runtime.RecordUpdateDict(origVal, []string{"a", "fa", "zArrayA"}, []gopurs_runtime.Value{v1_14, v2_15, gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v3_16.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}())})
											}
											r := (*gopurs_runtime.RecordData)(origVal.UnsafePtr)
											newVals := make([]gopurs_runtime.Value, len(r.Vals))
											copy(newVals, r.Vals)
											newVals[0] = v1_14
											newVals[3] = v2_15
											newVals[5] = gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_16.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}())
											newR := gopurs_runtime.RecordData{Keys: r.Keys, Vals: newVals}
											return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecordData, UnsafePtr: unsafe.Pointer(&newR)}
										}()}))}
									})
								})
							}), v_arrayMap19)
						}
						return gopurs_runtime.Array(res_go_arrayMap19)
					}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
						return x_14
					}), gopurs_runtime.RecordGet(Get_Main_recordValue_prime_(), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
						return x_14
					}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(Get_Main_recordValue_prime_(), "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
						arr := []*Constructor_Main_M4{(&Constructor_Main_M4{1, Get_Main_recordValue()})}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer(v)}
						}
						return gopurs_runtime.Array(boxed)
					}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_13_184
				// TAST (Let): __local_var_14_195 -> gopurs_runtime.Value
				__local_var_14_195 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_15 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_15, "zArrayA"), gopurs_runtime.RecordGet(rb_16, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_14_195
				// TAST (Let): __local_var_15_197 -> gopurs_runtime.Value
				__local_var_15_197 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_16 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_16, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_17, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_14_195, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_16, rb_17).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_15_197
				// TAST (Let): __local_var_16_198 -> gopurs_runtime.Value
				__local_var_16_198 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_17 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_18 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_17, "fa"), gopurs_runtime.RecordGet(rb_18, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_15_197, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_17, rb_18).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_16_198
				// TAST (Let): __local_var_17_199 -> gopurs_runtime.Value
				__local_var_17_199 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_18 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_19 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_18, "fIgnore"), gopurs_runtime.RecordGet(rb_19, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_16_198, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_18, rb_19).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_17_199
				// TAST (Let): __local_var_18_200 -> gopurs_runtime.Value
				__local_var_18_200 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_19 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_20 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_19, "arrayIgnore"), gopurs_runtime.RecordGet(rb_20, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_17_199, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_19, rb_20).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_18_200
				// TAST (Let): __local_var_15_196 -> gopurs_runtime.Value
				__local_var_15_196 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_19 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_20 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_19, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_20, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_18_200, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_19, rb_20).IntVal) != (0)))
					})
				}))
				_ = __local_var_15_196
				// TAST (Let): __local_var_16_201 -> gopurs_runtime.Value
				__local_var_16_201 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_16 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_17 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_196, "eq"), gopurs_runtime.RecordGet(ra_16, "nested"), gopurs_runtime.RecordGet(rb_17, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_16_201
				// TAST (Let): eqArray5_17_202 -> *Constructor_Data_Eq_Eq
				eqArray5_17_202 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_17_202
				var __t212 gopurs_runtime.Value
				{
					var __t_tag_204 gopurs_runtime.Value = Get_Main_m5_prime_()
					if __t_tag_204.Type == 9 && __t_tag_204.IntVal == 3852365315 {
						__t212 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M0{nil}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_212
					} else {

					}
				}
				{
					var __t_tag_205 gopurs_runtime.Value = Get_Main_m5_prime_()
					if __t_tag_205.Type == 9 && __t_tag_205.IntVal == 769986722 {
						__t212 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap20 := (*Constructor_Main_M1)(Get_Main_m5_prime_().UnsafePtr).V0
							_ = arr_val_arrayMap20
							arr_go_arrayMap20 := (*[]gopurs_runtime.Value)(arr_val_arrayMap20.UnsafePtr)
							_ = arr_go_arrayMap20
							res_go_arrayMap20 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap20))
							_ = res_go_arrayMap20
							for i_arrayMap20, v_arrayMap20 := range *arr_go_arrayMap20 {
								res_go_arrayMap20[i_arrayMap20] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, v2_15, func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v3_16.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()}))}
									})
								}), v_arrayMap20)
							}
							return gopurs_runtime.Array(res_go_arrayMap20)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_15
						}), gopurs_runtime.Array((*Constructor_Main_M1)(Get_Main_m5_prime_().UnsafePtr).V1)))
						goto end_branch_212
					} else {

					}
				}
				{
					var __t_tag_206 gopurs_runtime.Value = Get_Main_m5_prime_()
					if __t_tag_206.Type == 9 && __t_tag_206.IntVal == 2727978561 {
						__t212 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, (*Constructor_Main_M2)(Get_Main_m5_prime_().UnsafePtr).V0})}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_212
					} else {

					}
				}
				{
					var __t_tag_207 gopurs_runtime.Value = Get_Main_m5_prime_()
					if __t_tag_207.Type == 9 && __t_tag_207.IntVal == 1830062304 {
						__t212 = func() gopurs_runtime.Value {
							arr_val_arrayMap19 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
								return x_15
							}), (*Constructor_Main_M3)(Get_Main_m5_prime_().UnsafePtr).V0)
							_ = arr_val_arrayMap19
							arr_go_arrayMap19 := (*[]gopurs_runtime.Value)(arr_val_arrayMap19.UnsafePtr)
							_ = arr_go_arrayMap19
							res_go_arrayMap19 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap19))
							_ = res_go_arrayMap19
							for i_arrayMap19, v_arrayMap19 := range *arr_go_arrayMap19 {
								res_go_arrayMap19[i_arrayMap19] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, v1_15}))}
								}), v_arrayMap19)
							}
							return gopurs_runtime.Array(res_go_arrayMap19)
						}()
						goto end_branch_212
					} else {

					}
				}
				{
					var __t_tag_208 gopurs_runtime.Value = Get_Main_m5_prime_()
					if __t_tag_208.Type == 9 && __t_tag_208.IntVal == 3190619783 {
						__t212 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap21 := gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m5_prime_().UnsafePtr).V0, "a")
							_ = arr_val_arrayMap21
							arr_go_arrayMap21 := (*[]gopurs_runtime.Value)(arr_val_arrayMap21.UnsafePtr)
							_ = arr_go_arrayMap21
							res_go_arrayMap21 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap21))
							_ = res_go_arrayMap21
							for i_arrayMap21, v_arrayMap21 := range *arr_go_arrayMap21 {
								res_go_arrayMap21[i_arrayMap21] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4)(Get_Main_m5_prime_().UnsafePtr).V0, "a", v1_15, "fa", v2_16, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_17.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}()))}))}
										})
									})
								}), v_arrayMap21)
							}
							return gopurs_runtime.Array(res_go_arrayMap21)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_15
						}), gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m5_prime_().UnsafePtr).V0, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_15
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m5_prime_().UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_212
					} else {

					}
				}
				{
					var __t_tag_209 gopurs_runtime.Value = Get_Main_m5_prime_()
					if __t_tag_209.Type == 9 && __t_tag_209.IntVal == 108241190 {
						__t212 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap21 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m5_prime_().UnsafePtr).V0, "nested"), "a")
							_ = arr_val_arrayMap21
							arr_go_arrayMap21 := (*[]gopurs_runtime.Value)(arr_val_arrayMap21.UnsafePtr)
							_ = arr_go_arrayMap21
							res_go_arrayMap21 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap21))
							_ = res_go_arrayMap21
							for i_arrayMap21, v_arrayMap21 := range *arr_go_arrayMap21 {
								res_go_arrayMap21[i_arrayMap21] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, func() gopurs_runtime.Value {
												origVal := (*Constructor_Main_M5)(Get_Main_m5_prime_().UnsafePtr).V0
												if origVal.Type != gopurs_runtime.TypeRecord1 {
													return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m5_prime_().UnsafePtr).V0, "nested"), "a", v1_15, "fa", v2_16, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v3_17.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}()))})
												}
												clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
												clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m5_prime_().UnsafePtr).V0, "nested"), "a", v1_15, "fa", v2_16, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v3_17.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}()))
												return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
											}()}))}
										})
									})
								}), v_arrayMap21)
							}
							return gopurs_runtime.Array(res_go_arrayMap21)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_15
						}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m5_prime_().UnsafePtr).V0, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_15
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m5_prime_().UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_212
					} else {

					}
				}
				{
					var __t_tag_210 gopurs_runtime.Value = Get_Main_m5_prime_()
					if __t_tag_210.Type == 9 && __t_tag_210.IntVal == 2066233029 {
						__t212 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap27 := (*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V1
							_ = arr_val_arrayMap27
							arr_go_arrayMap27 := (*[]gopurs_runtime.Value)(arr_val_arrayMap27.UnsafePtr)
							_ = arr_go_arrayMap27
							res_go_arrayMap27 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap27))
							_ = res_go_arrayMap27
							for i_arrayMap27, v_arrayMap27 := range *arr_go_arrayMap27 {
								res_go_arrayMap27[i_arrayMap27] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v8_15 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v9_16 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v10_17 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v11_18 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v12_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v13_20 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Func(func(v14_21 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(v15_22 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(v16_23 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, (*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V0, v8_15, (*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V2, func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v9_16.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}(), v10_17, (*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V6, "a", v11_18, "fa", v12_19, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v13_20.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}())), func() gopurs_runtime.Value {
																		origVal := (*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V7
																		if origVal.Type != gopurs_runtime.TypeRecord1 {
																			return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V7, "nested"), "a", v14_21, "fa", v15_22, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																				arr := *(*[]gopurs_runtime.Value)(v16_23.UnsafePtr)
																				unboxed := make([]gopurs_runtime.Value, len(arr))
																				for i, v := range arr {
																					unboxed[i] = v
																				}
																				return unboxed
																			}()))})
																		}
																		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
																		clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V7, "nested"), "a", v14_21, "fa", v15_22, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																			arr := *(*[]gopurs_runtime.Value)(v16_23.UnsafePtr)
																			unboxed := make([]gopurs_runtime.Value, len(arr))
																			for i, v := range arr {
																				unboxed[i] = v
																			}
																			return unboxed
																		}()))
																		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
																	}()}))}
																})
															})
														})
													})
												})
											})
										})
									})
								}), v_arrayMap27)
							}
							return gopurs_runtime.Array(res_go_arrayMap27)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_15
						}), gopurs_runtime.Array((*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V3))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_15
						}), (*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V4)), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V6, "a")), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_15
						}), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V6, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_15
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V7, "nested"), "a")), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_15
						}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V7, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_15
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m5_prime_().UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_212
					} else {

					}
				}
				{
					var __t_tag_211 gopurs_runtime.Value = Get_Main_m5_prime_()
					if __t_tag_211.Type == 9 && __t_tag_211.IntVal == 1168316772 {
						__t212 = func() gopurs_runtime.Value {
							arr_val_arrayMap19 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply5(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
									arr_val_arrayMap25 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_15, "nested"), "a")
									_ = arr_val_arrayMap25
									arr_go_arrayMap25 := (*[]gopurs_runtime.Value)(arr_val_arrayMap25.UnsafePtr)
									_ = arr_go_arrayMap25
									res_go_arrayMap25 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap25))
									_ = res_go_arrayMap25
									for i_arrayMap25, v_arrayMap25 := range *arr_go_arrayMap25 {
										res_go_arrayMap25[i_arrayMap25] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v4_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.RecordUpdate1(v1_15, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_15, "nested"), "a", v2_16, "fa", v3_17, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v4_18.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}())))
												})
											})
										}), v_arrayMap25)
									}
									return gopurs_runtime.Array(res_go_arrayMap25)
								}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
									return x_16
								}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_15, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
									return x_16
								}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_15, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
							})), (*Constructor_Main_M7)(Get_Main_m5_prime_().UnsafePtr).V0)
							_ = arr_val_arrayMap19
							arr_go_arrayMap19 := (*[]gopurs_runtime.Value)(arr_val_arrayMap19.UnsafePtr)
							_ = arr_go_arrayMap19
							res_go_arrayMap19 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap19))
							_ = res_go_arrayMap19
							for i_arrayMap19, v_arrayMap19 := range *arr_go_arrayMap19 {
								res_go_arrayMap19[i_arrayMap19] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, v1_15}))}
								}), v_arrayMap19)
							}
							return gopurs_runtime.Array(res_go_arrayMap19)
						}()
						goto end_branch_212
					} else {

					}
				}
				{
					__t212 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_212:
				_dollar___unused_14_194 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m5"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_19 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t203 bool
						{
							if x_18.Type == 9 && x_18.IntVal == 3852365315 {
								__t203 = (y_19.Type == 9 && y_19.IntVal == 3852365315)
								goto end_branch_203
							} else {

							}
						}
						{
							if x_18.Type == 9 && x_18.IntVal == 769986722 {
								__t203 = (y_19.Type == 9 && y_19.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_18.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_19.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_17_202.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_18.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_19.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_203
							} else {

							}
						}
						{
							if x_18.Type == 9 && x_18.IntVal == 2727978561 {
								__t203 = (y_19.Type == 9 && y_19.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_18.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_19.UnsafePtr).V0))
								goto end_branch_203
							} else {

							}
						}
						{
							if x_18.Type == 9 && x_18.IntVal == 1830062304 {
								__t203 = (y_19.Type == 9 && y_19.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_18.UnsafePtr).V0, (*Constructor_Main_M3)(y_19.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_203
							} else {

							}
						}
						{
							if x_18.Type == 9 && x_18.IntVal == 3190619783 {
								__t203 = (y_19.Type == 9 && y_19.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_18.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_19.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_18.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_19.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_18.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_19.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_18.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_19.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_18.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_19.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_17_202.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_18.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_19.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_203
							} else {

							}
						}
						{
							if x_18.Type == 9 && x_18.IntVal == 108241190 {
								__t203 = (y_19.Type == 9 && y_19.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_18.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_19.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_18.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_19.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_18.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_19.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_18.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_19.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_18.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_19.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_17_202.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_18.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_19.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_203
							} else {

							}
						}
						{
							if x_18.Type == 9 && x_18.IntVal == 2066233029 {
								__t203 = (y_19.Type == 9 && y_19.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_18.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_19.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_18.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_19.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_18.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_19.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_17_202.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_18.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_19.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_18.UnsafePtr).V4, (*Constructor_Main_M6)(y_19.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_18.UnsafePtr).V5, (*Constructor_Main_M6)(y_19.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_18.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_19.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_18.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_19.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_18.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_19.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_18.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_19.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_18.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_19.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_17_202.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_18.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_19.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_18.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_19.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_18.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_19.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_18.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_19.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_18.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_19.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_18.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_19.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_17_202.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_18.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_19.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_203
							} else {

							}
						}
						{
							__t203 = (x_18.Type == 9 && x_18.IntVal == 1168316772) && ((y_19.Type == 9 && y_19.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_16_201, "eq"), (*Constructor_Main_M7)(x_18.UnsafePtr).V0, (*Constructor_Main_M7)(y_19.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_203:
						return gopurs_runtime.Bool(__t203)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t212.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m5()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_14_194
				// TAST (Let): __local_var_15_214 -> gopurs_runtime.Value
				__local_var_15_214 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_16 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_16, "zArrayA"), gopurs_runtime.RecordGet(rb_17, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_15_214
				// TAST (Let): __local_var_16_216 -> gopurs_runtime.Value
				__local_var_16_216 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_17 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_18 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_17, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_18, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_15_214, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_17, rb_18).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_16_216
				// TAST (Let): __local_var_17_217 -> gopurs_runtime.Value
				__local_var_17_217 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_18 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_19 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_18, "fa"), gopurs_runtime.RecordGet(rb_19, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_16_216, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_18, rb_19).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_17_217
				// TAST (Let): __local_var_18_218 -> gopurs_runtime.Value
				__local_var_18_218 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_19 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_20 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_19, "fIgnore"), gopurs_runtime.RecordGet(rb_20, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_17_217, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_19, rb_20).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_18_218
				// TAST (Let): __local_var_19_219 -> gopurs_runtime.Value
				__local_var_19_219 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_20 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_21 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_20, "arrayIgnore"), gopurs_runtime.RecordGet(rb_21, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_18_218, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_20, rb_21).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_19_219
				// TAST (Let): __local_var_16_215 -> gopurs_runtime.Value
				__local_var_16_215 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_20 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_21 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_20, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_21, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_19_219, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_20, rb_21).IntVal) != (0)))
					})
				}))
				_ = __local_var_16_215
				// TAST (Let): __local_var_17_220 -> gopurs_runtime.Value
				__local_var_17_220 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_17 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_18 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_215, "eq"), gopurs_runtime.RecordGet(ra_17, "nested"), gopurs_runtime.RecordGet(rb_18, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_17_220
				// TAST (Let): eqArray5_18_221 -> *Constructor_Data_Eq_Eq
				eqArray5_18_221 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_18_221
				var __t231 gopurs_runtime.Value
				{
					var __t_tag_223 gopurs_runtime.Value = Get_Main_m6_prime_()
					if __t_tag_223.Type == 9 && __t_tag_223.IntVal == 3852365315 {
						__t231 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M0{nil}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_231
					} else {

					}
				}
				{
					var __t_tag_224 gopurs_runtime.Value = Get_Main_m6_prime_()
					if __t_tag_224.Type == 9 && __t_tag_224.IntVal == 769986722 {
						__t231 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap21 := (*Constructor_Main_M1)(Get_Main_m6_prime_().UnsafePtr).V0
							_ = arr_val_arrayMap21
							arr_go_arrayMap21 := (*[]gopurs_runtime.Value)(arr_val_arrayMap21.UnsafePtr)
							_ = arr_go_arrayMap21
							res_go_arrayMap21 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap21))
							_ = res_go_arrayMap21
							for i_arrayMap21, v_arrayMap21 := range *arr_go_arrayMap21 {
								res_go_arrayMap21[i_arrayMap21] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, v2_16, func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v3_17.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()}))}
									})
								}), v_arrayMap21)
							}
							return gopurs_runtime.Array(res_go_arrayMap21)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_16
						}), gopurs_runtime.Array((*Constructor_Main_M1)(Get_Main_m6_prime_().UnsafePtr).V1)))
						goto end_branch_231
					} else {

					}
				}
				{
					var __t_tag_225 gopurs_runtime.Value = Get_Main_m6_prime_()
					if __t_tag_225.Type == 9 && __t_tag_225.IntVal == 2727978561 {
						__t231 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, (*Constructor_Main_M2)(Get_Main_m6_prime_().UnsafePtr).V0})}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_231
					} else {

					}
				}
				{
					var __t_tag_226 gopurs_runtime.Value = Get_Main_m6_prime_()
					if __t_tag_226.Type == 9 && __t_tag_226.IntVal == 1830062304 {
						__t231 = func() gopurs_runtime.Value {
							arr_val_arrayMap20 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
								return x_16
							}), (*Constructor_Main_M3)(Get_Main_m6_prime_().UnsafePtr).V0)
							_ = arr_val_arrayMap20
							arr_go_arrayMap20 := (*[]gopurs_runtime.Value)(arr_val_arrayMap20.UnsafePtr)
							_ = arr_go_arrayMap20
							res_go_arrayMap20 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap20))
							_ = res_go_arrayMap20
							for i_arrayMap20, v_arrayMap20 := range *arr_go_arrayMap20 {
								res_go_arrayMap20[i_arrayMap20] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, v1_16}))}
								}), v_arrayMap20)
							}
							return gopurs_runtime.Array(res_go_arrayMap20)
						}()
						goto end_branch_231
					} else {

					}
				}
				{
					var __t_tag_227 gopurs_runtime.Value = Get_Main_m6_prime_()
					if __t_tag_227.Type == 9 && __t_tag_227.IntVal == 3190619783 {
						__t231 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap22 := gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m6_prime_().UnsafePtr).V0, "a")
							_ = arr_val_arrayMap22
							arr_go_arrayMap22 := (*[]gopurs_runtime.Value)(arr_val_arrayMap22.UnsafePtr)
							_ = arr_go_arrayMap22
							res_go_arrayMap22 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap22))
							_ = res_go_arrayMap22
							for i_arrayMap22, v_arrayMap22 := range *arr_go_arrayMap22 {
								res_go_arrayMap22[i_arrayMap22] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_18 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4)(Get_Main_m6_prime_().UnsafePtr).V0, "a", v1_16, "fa", v2_17, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_18.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}()))}))}
										})
									})
								}), v_arrayMap22)
							}
							return gopurs_runtime.Array(res_go_arrayMap22)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_16
						}), gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m6_prime_().UnsafePtr).V0, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_16
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m6_prime_().UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_231
					} else {

					}
				}
				{
					var __t_tag_228 gopurs_runtime.Value = Get_Main_m6_prime_()
					if __t_tag_228.Type == 9 && __t_tag_228.IntVal == 108241190 {
						__t231 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap22 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m6_prime_().UnsafePtr).V0, "nested"), "a")
							_ = arr_val_arrayMap22
							arr_go_arrayMap22 := (*[]gopurs_runtime.Value)(arr_val_arrayMap22.UnsafePtr)
							_ = arr_go_arrayMap22
							res_go_arrayMap22 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap22))
							_ = res_go_arrayMap22
							for i_arrayMap22, v_arrayMap22 := range *arr_go_arrayMap22 {
								res_go_arrayMap22[i_arrayMap22] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_18 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, func() gopurs_runtime.Value {
												origVal := (*Constructor_Main_M5)(Get_Main_m6_prime_().UnsafePtr).V0
												if origVal.Type != gopurs_runtime.TypeRecord1 {
													return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m6_prime_().UnsafePtr).V0, "nested"), "a", v1_16, "fa", v2_17, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v3_18.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}()))})
												}
												clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
												clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m6_prime_().UnsafePtr).V0, "nested"), "a", v1_16, "fa", v2_17, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v3_18.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}()))
												return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
											}()}))}
										})
									})
								}), v_arrayMap22)
							}
							return gopurs_runtime.Array(res_go_arrayMap22)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_16
						}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m6_prime_().UnsafePtr).V0, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_16
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m6_prime_().UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_231
					} else {

					}
				}
				{
					var __t_tag_229 gopurs_runtime.Value = Get_Main_m6_prime_()
					if __t_tag_229.Type == 9 && __t_tag_229.IntVal == 2066233029 {
						__t231 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap28 := (*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V1
							_ = arr_val_arrayMap28
							arr_go_arrayMap28 := (*[]gopurs_runtime.Value)(arr_val_arrayMap28.UnsafePtr)
							_ = arr_go_arrayMap28
							res_go_arrayMap28 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap28))
							_ = res_go_arrayMap28
							for i_arrayMap28, v_arrayMap28 := range *arr_go_arrayMap28 {
								res_go_arrayMap28[i_arrayMap28] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v8_16 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v9_17 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v10_18 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v11_19 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v12_20 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v13_21 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Func(func(v14_22 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(v15_23 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(v16_24 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, (*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V0, v8_16, (*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V2, func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v9_17.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}(), v10_18, (*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V6, "a", v11_19, "fa", v12_20, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v13_21.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}())), func() gopurs_runtime.Value {
																		origVal := (*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V7
																		if origVal.Type != gopurs_runtime.TypeRecord1 {
																			return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V7, "nested"), "a", v14_22, "fa", v15_23, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																				arr := *(*[]gopurs_runtime.Value)(v16_24.UnsafePtr)
																				unboxed := make([]gopurs_runtime.Value, len(arr))
																				for i, v := range arr {
																					unboxed[i] = v
																				}
																				return unboxed
																			}()))})
																		}
																		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
																		clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V7, "nested"), "a", v14_22, "fa", v15_23, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																			arr := *(*[]gopurs_runtime.Value)(v16_24.UnsafePtr)
																			unboxed := make([]gopurs_runtime.Value, len(arr))
																			for i, v := range arr {
																				unboxed[i] = v
																			}
																			return unboxed
																		}()))
																		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
																	}()}))}
																})
															})
														})
													})
												})
											})
										})
									})
								}), v_arrayMap28)
							}
							return gopurs_runtime.Array(res_go_arrayMap28)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_16
						}), gopurs_runtime.Array((*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V3))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_16
						}), (*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V4)), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V6, "a")), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_16
						}), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V6, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_16
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V7, "nested"), "a")), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_16
						}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V7, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_16
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m6_prime_().UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_231
					} else {

					}
				}
				{
					var __t_tag_230 gopurs_runtime.Value = Get_Main_m6_prime_()
					if __t_tag_230.Type == 9 && __t_tag_230.IntVal == 1168316772 {
						__t231 = func() gopurs_runtime.Value {
							arr_val_arrayMap20 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply5(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
									arr_val_arrayMap26 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_16, "nested"), "a")
									_ = arr_val_arrayMap26
									arr_go_arrayMap26 := (*[]gopurs_runtime.Value)(arr_val_arrayMap26.UnsafePtr)
									_ = arr_go_arrayMap26
									res_go_arrayMap26 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap26))
									_ = res_go_arrayMap26
									for i_arrayMap26, v_arrayMap26 := range *arr_go_arrayMap26 {
										res_go_arrayMap26[i_arrayMap26] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v3_18 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v4_19 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.RecordUpdate1(v1_16, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_16, "nested"), "a", v2_17, "fa", v3_18, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v4_19.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}())))
												})
											})
										}), v_arrayMap26)
									}
									return gopurs_runtime.Array(res_go_arrayMap26)
								}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
									return x_17
								}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_16, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
									return x_17
								}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_16, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
							})), (*Constructor_Main_M7)(Get_Main_m6_prime_().UnsafePtr).V0)
							_ = arr_val_arrayMap20
							arr_go_arrayMap20 := (*[]gopurs_runtime.Value)(arr_val_arrayMap20.UnsafePtr)
							_ = arr_go_arrayMap20
							res_go_arrayMap20 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap20))
							_ = res_go_arrayMap20
							for i_arrayMap20, v_arrayMap20 := range *arr_go_arrayMap20 {
								res_go_arrayMap20[i_arrayMap20] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, v1_16}))}
								}), v_arrayMap20)
							}
							return gopurs_runtime.Array(res_go_arrayMap20)
						}()
						goto end_branch_231
					} else {

					}
				}
				{
					__t231 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_231:
				_dollar___unused_15_213 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m6"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_20 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t222 bool
						{
							if x_19.Type == 9 && x_19.IntVal == 3852365315 {
								__t222 = (y_20.Type == 9 && y_20.IntVal == 3852365315)
								goto end_branch_222
							} else {

							}
						}
						{
							if x_19.Type == 9 && x_19.IntVal == 769986722 {
								__t222 = (y_20.Type == 9 && y_20.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_19.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_20.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_18_221.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_19.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_20.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_222
							} else {

							}
						}
						{
							if x_19.Type == 9 && x_19.IntVal == 2727978561 {
								__t222 = (y_20.Type == 9 && y_20.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_19.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_20.UnsafePtr).V0))
								goto end_branch_222
							} else {

							}
						}
						{
							if x_19.Type == 9 && x_19.IntVal == 1830062304 {
								__t222 = (y_20.Type == 9 && y_20.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_19.UnsafePtr).V0, (*Constructor_Main_M3)(y_20.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_222
							} else {

							}
						}
						{
							if x_19.Type == 9 && x_19.IntVal == 3190619783 {
								__t222 = (y_20.Type == 9 && y_20.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_19.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_20.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_19.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_20.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_19.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_20.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_19.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_20.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_19.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_20.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_18_221.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_19.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_20.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_222
							} else {

							}
						}
						{
							if x_19.Type == 9 && x_19.IntVal == 108241190 {
								__t222 = (y_20.Type == 9 && y_20.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_19.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_20.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_19.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_20.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_19.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_20.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_19.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_20.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_19.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_20.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_18_221.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_19.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_20.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_222
							} else {

							}
						}
						{
							if x_19.Type == 9 && x_19.IntVal == 2066233029 {
								__t222 = (y_20.Type == 9 && y_20.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_19.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_20.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_19.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_20.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_19.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_20.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_18_221.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_19.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_20.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_19.UnsafePtr).V4, (*Constructor_Main_M6)(y_20.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_19.UnsafePtr).V5, (*Constructor_Main_M6)(y_20.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_19.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_20.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_19.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_20.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_19.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_20.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_19.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_20.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_19.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_20.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_18_221.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_19.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_20.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_19.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_20.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_19.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_20.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_19.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_20.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_19.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_20.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_19.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_20.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_18_221.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_19.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_20.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_222
							} else {

							}
						}
						{
							__t222 = (x_19.Type == 9 && x_19.IntVal == 1168316772) && ((y_20.Type == 9 && y_20.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_17_220, "eq"), (*Constructor_Main_M7)(x_19.UnsafePtr).V0, (*Constructor_Main_M7)(y_20.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_222:
						return gopurs_runtime.Bool(__t222)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t231.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m6()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_15_213
				// TAST (Let): __local_var_16_233 -> gopurs_runtime.Value
				__local_var_16_233 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_17 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_18 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_17, "zArrayA"), gopurs_runtime.RecordGet(rb_18, "zArrayA")).IntVal) != (0))
						})
					})
				}))
				_ = __local_var_16_233
				// TAST (Let): __local_var_17_235 -> gopurs_runtime.Value
				__local_var_17_235 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_18 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_19 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_18, "ignore").IntVal) == (gopurs_runtime.RecordGet(rb_19, "ignore").IntVal)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_16_233, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_18, rb_19).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_17_235
				// TAST (Let): __local_var_18_236 -> gopurs_runtime.Value
				__local_var_18_236 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_19 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_20 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray1(), "eq"), gopurs_runtime.RecordGet(ra_19, "fa"), gopurs_runtime.RecordGet(rb_20, "fa")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_17_235, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_19, rb_20).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_18_236
				// TAST (Let): __local_var_19_237 -> gopurs_runtime.Value
				__local_var_19_237 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_20 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_21 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_20, "fIgnore"), gopurs_runtime.RecordGet(rb_21, "fIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_18_236, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_20, rb_21).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_19_237
				// TAST (Let): __local_var_20_238 -> gopurs_runtime.Value
				__local_var_20_238 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(ra_21 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(rb_22 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Main_eqArray2(), "eq"), gopurs_runtime.RecordGet(ra_21, "arrayIgnore"), gopurs_runtime.RecordGet(rb_22, "arrayIgnore")).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_19_237, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_21, rb_22).IntVal) != (0)))
						})
					})
				}))
				_ = __local_var_20_238
				// TAST (Let): __local_var_17_234 -> gopurs_runtime.Value
				__local_var_17_234 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_21 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_22 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool(((gopurs_runtime.RecordGet(ra_21, "a").StrVal()) == (gopurs_runtime.RecordGet(rb_22, "a").StrVal())) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_20_238, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_21, rb_22).IntVal) != (0)))
					})
				}))
				_ = __local_var_17_234
				// TAST (Let): __local_var_18_239 -> gopurs_runtime.Value
				__local_var_18_239 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(ra_18 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(rb_19 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_17_234, "eq"), gopurs_runtime.RecordGet(ra_18, "nested"), gopurs_runtime.RecordGet(rb_19, "nested")).IntVal) != (0))
					})
				})))
				_ = __local_var_18_239
				// TAST (Let): eqArray5_19_240 -> *Constructor_Data_Eq_Eq
				eqArray5_19_240 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl())})
				_ = eqArray5_19_240
				var __t250 gopurs_runtime.Value
				{
					var __t_tag_242 gopurs_runtime.Value = Get_Main_m7_prime_()
					if __t_tag_242.Type == 9 && __t_tag_242.IntVal == 3852365315 {
						__t250 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M0{nil}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_250
					} else {

					}
				}
				{
					var __t_tag_243 gopurs_runtime.Value = Get_Main_m7_prime_()
					if __t_tag_243.Type == 9 && __t_tag_243.IntVal == 769986722 {
						__t250 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap22 := (*Constructor_Main_M1)(Get_Main_m7_prime_().UnsafePtr).V0
							_ = arr_val_arrayMap22
							arr_go_arrayMap22 := (*[]gopurs_runtime.Value)(arr_val_arrayMap22.UnsafePtr)
							_ = arr_go_arrayMap22
							res_go_arrayMap22 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap22))
							_ = res_go_arrayMap22
							for i_arrayMap22, v_arrayMap22 := range *arr_go_arrayMap22 {
								res_go_arrayMap22[i_arrayMap22] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_18 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, v2_17, func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v3_18.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()}))}
									})
								}), v_arrayMap22)
							}
							return gopurs_runtime.Array(res_go_arrayMap22)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_17
						}), gopurs_runtime.Array((*Constructor_Main_M1)(Get_Main_m7_prime_().UnsafePtr).V1)))
						goto end_branch_250
					} else {

					}
				}
				{
					var __t_tag_244 gopurs_runtime.Value = Get_Main_m7_prime_()
					if __t_tag_244.Type == 9 && __t_tag_244.IntVal == 2727978561 {
						__t250 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr := []*Constructor_Main_M2{(&Constructor_Main_M2{1, (*Constructor_Main_M2)(Get_Main_m7_prime_().UnsafePtr).V0})}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer(v)}
								}
								return gopurs_runtime.Array(boxed)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())
						goto end_branch_250
					} else {

					}
				}
				{
					var __t_tag_245 gopurs_runtime.Value = Get_Main_m7_prime_()
					if __t_tag_245.Type == 9 && __t_tag_245.IntVal == 1830062304 {
						__t250 = func() gopurs_runtime.Value {
							arr_val_arrayMap21 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
								return x_17
							}), (*Constructor_Main_M3)(Get_Main_m7_prime_().UnsafePtr).V0)
							_ = arr_val_arrayMap21
							arr_go_arrayMap21 := (*[]gopurs_runtime.Value)(arr_val_arrayMap21.UnsafePtr)
							_ = arr_go_arrayMap21
							res_go_arrayMap21 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap21))
							_ = res_go_arrayMap21
							for i_arrayMap21, v_arrayMap21 := range *arr_go_arrayMap21 {
								res_go_arrayMap21[i_arrayMap21] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, v1_17}))}
								}), v_arrayMap21)
							}
							return gopurs_runtime.Array(res_go_arrayMap21)
						}()
						goto end_branch_250
					} else {

					}
				}
				{
					var __t_tag_246 gopurs_runtime.Value = Get_Main_m7_prime_()
					if __t_tag_246.Type == 9 && __t_tag_246.IntVal == 3190619783 {
						__t250 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap23 := gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m7_prime_().UnsafePtr).V0, "a")
							_ = arr_val_arrayMap23
							arr_go_arrayMap23 := (*[]gopurs_runtime.Value)(arr_val_arrayMap23.UnsafePtr)
							_ = arr_go_arrayMap23
							res_go_arrayMap23 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap23))
							_ = res_go_arrayMap23
							for i_arrayMap23, v_arrayMap23 := range *arr_go_arrayMap23 {
								res_go_arrayMap23[i_arrayMap23] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_19 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4)(Get_Main_m7_prime_().UnsafePtr).V0, "a", v1_17, "fa", v2_18, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v3_19.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}()))}))}
										})
									})
								}), v_arrayMap23)
							}
							return gopurs_runtime.Array(res_go_arrayMap23)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_17
						}), gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m7_prime_().UnsafePtr).V0, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_17
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(Get_Main_m7_prime_().UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_250
					} else {

					}
				}
				{
					var __t_tag_247 gopurs_runtime.Value = Get_Main_m7_prime_()
					if __t_tag_247.Type == 9 && __t_tag_247.IntVal == 108241190 {
						__t250 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap23 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m7_prime_().UnsafePtr).V0, "nested"), "a")
							_ = arr_val_arrayMap23
							arr_go_arrayMap23 := (*[]gopurs_runtime.Value)(arr_val_arrayMap23.UnsafePtr)
							_ = arr_go_arrayMap23
							res_go_arrayMap23 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap23))
							_ = res_go_arrayMap23
							for i_arrayMap23, v_arrayMap23 := range *arr_go_arrayMap23 {
								res_go_arrayMap23[i_arrayMap23] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v3_19 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, func() gopurs_runtime.Value {
												origVal := (*Constructor_Main_M5)(Get_Main_m7_prime_().UnsafePtr).V0
												if origVal.Type != gopurs_runtime.TypeRecord1 {
													return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m7_prime_().UnsafePtr).V0, "nested"), "a", v1_17, "fa", v2_18, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v3_19.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}()))})
												}
												clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
												clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m7_prime_().UnsafePtr).V0, "nested"), "a", v1_17, "fa", v2_18, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
													arr := *(*[]gopurs_runtime.Value)(v3_19.UnsafePtr)
													unboxed := make([]gopurs_runtime.Value, len(arr))
													for i, v := range arr {
														unboxed[i] = v
													}
													return unboxed
												}()))
												return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
											}()}))}
										})
									})
								}), v_arrayMap23)
							}
							return gopurs_runtime.Array(res_go_arrayMap23)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_17
						}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m7_prime_().UnsafePtr).V0, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_17
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(Get_Main_m7_prime_().UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_250
					} else {

					}
				}
				{
					var __t_tag_248 gopurs_runtime.Value = Get_Main_m7_prime_()
					if __t_tag_248.Type == 9 && __t_tag_248.IntVal == 2066233029 {
						__t250 = gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
							arr_val_arrayMap29 := (*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V1
							_ = arr_val_arrayMap29
							arr_go_arrayMap29 := (*[]gopurs_runtime.Value)(arr_val_arrayMap29.UnsafePtr)
							_ = arr_go_arrayMap29
							res_go_arrayMap29 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap29))
							_ = res_go_arrayMap29
							for i_arrayMap29, v_arrayMap29 := range *arr_go_arrayMap29 {
								res_go_arrayMap29[i_arrayMap29] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v8_17 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v9_18 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v10_19 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v11_20 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v12_21 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v13_22 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Func(func(v14_23 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(v15_24 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Func(func(v16_25 gopurs_runtime.Value) gopurs_runtime.Value {
																	return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, (*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V0, v8_17, (*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V2, func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v9_18.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}(), v10_19, (*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V6, "a", v11_20, "fa", v12_21, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																		arr := *(*[]gopurs_runtime.Value)(v13_22.UnsafePtr)
																		unboxed := make([]gopurs_runtime.Value, len(arr))
																		for i, v := range arr {
																			unboxed[i] = v
																		}
																		return unboxed
																	}())), func() gopurs_runtime.Value {
																		origVal := (*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V7
																		if origVal.Type != gopurs_runtime.TypeRecord1 {
																			return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V7, "nested"), "a", v14_23, "fa", v15_24, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																				arr := *(*[]gopurs_runtime.Value)(v16_25.UnsafePtr)
																				unboxed := make([]gopurs_runtime.Value, len(arr))
																				for i, v := range arr {
																					unboxed[i] = v
																				}
																				return unboxed
																			}()))})
																		}
																		clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
																		clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V7, "nested"), "a", v14_23, "fa", v15_24, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																			arr := *(*[]gopurs_runtime.Value)(v16_25.UnsafePtr)
																			unboxed := make([]gopurs_runtime.Value, len(arr))
																			for i, v := range arr {
																				unboxed[i] = v
																			}
																			return unboxed
																		}()))
																		return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
																	}()}))}
																})
															})
														})
													})
												})
											})
										})
									})
								}), v_arrayMap29)
							}
							return gopurs_runtime.Array(res_go_arrayMap29)
						}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_17
						}), gopurs_runtime.Array((*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V3))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_17
						}), (*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V4)), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V6, "a")), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_17
						}), gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V6, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_17
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V7, "nested"), "a")), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_17
						}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V7, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_17 gopurs_runtime.Value) gopurs_runtime.Value {
							return x_17
						}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(Get_Main_m7_prime_().UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
						goto end_branch_250
					} else {

					}
				}
				{
					var __t_tag_249 gopurs_runtime.Value = Get_Main_m7_prime_()
					if __t_tag_249.Type == 9 && __t_tag_249.IntVal == 1168316772 {
						__t250 = func() gopurs_runtime.Value {
							arr_val_arrayMap21 := gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply5(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), gopurs_runtime.Apply2(Get_Control_Apply_arrayApply(), func() gopurs_runtime.Value {
									arr_val_arrayMap27 := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_17, "nested"), "a")
									_ = arr_val_arrayMap27
									arr_go_arrayMap27 := (*[]gopurs_runtime.Value)(arr_val_arrayMap27.UnsafePtr)
									_ = arr_go_arrayMap27
									res_go_arrayMap27 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap27))
									_ = res_go_arrayMap27
									for i_arrayMap27, v_arrayMap27 := range *arr_go_arrayMap27 {
										res_go_arrayMap27[i_arrayMap27] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v3_19 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v4_20 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.RecordUpdate1(v1_17, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_17, "nested"), "a", v2_18, "fa", v3_19, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
														arr := *(*[]gopurs_runtime.Value)(v4_20.UnsafePtr)
														unboxed := make([]gopurs_runtime.Value, len(arr))
														for i, v := range arr {
															unboxed[i] = v
														}
														return unboxed
													}())))
												})
											})
										}), v_arrayMap27)
									}
									return gopurs_runtime.Array(res_go_arrayMap27)
								}(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
									return x_18
								}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_17, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Get_Control_Apply_arrayApply(), Get_Data_Functor_arrayMap(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1), Get_Data_Semigroup_concatArray(), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
									return x_18
								}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_17, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
							})), (*Constructor_Main_M7)(Get_Main_m7_prime_().UnsafePtr).V0)
							_ = arr_val_arrayMap21
							arr_go_arrayMap21 := (*[]gopurs_runtime.Value)(arr_val_arrayMap21.UnsafePtr)
							_ = arr_go_arrayMap21
							res_go_arrayMap21 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap21))
							_ = res_go_arrayMap21
							for i_arrayMap21, v_arrayMap21 := range *arr_go_arrayMap21 {
								res_go_arrayMap21[i_arrayMap21] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, v1_17}))}
								}), v_arrayMap21)
							}
							return gopurs_runtime.Array(res_go_arrayMap21)
						}()
						goto end_branch_250
					} else {

					}
				}
				{
					__t250 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_250:
				_dollar___unused_16_232 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str("sequence - m7"), gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.Func(func(x_20 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(y_21 gopurs_runtime.Value) gopurs_runtime.Value {
						var __t241 bool
						{
							if x_20.Type == 9 && x_20.IntVal == 3852365315 {
								__t241 = (y_21.Type == 9 && y_21.IntVal == 3852365315)
								goto end_branch_241
							} else {

							}
						}
						{
							if x_20.Type == 9 && x_20.IntVal == 769986722 {
								__t241 = (y_21.Type == 9 && y_21.IntVal == 769986722) && ((((*Constructor_Main_M1)(x_20.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_M1)(y_21.UnsafePtr).V0.StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_19_240.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_20.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_21.UnsafePtr).V1)).IntVal) != (0)))
								goto end_branch_241
							} else {

							}
						}
						{
							if x_20.Type == 9 && x_20.IntVal == 2727978561 {
								__t241 = (y_21.Type == 9 && y_21.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_20.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_21.UnsafePtr).V0))
								goto end_branch_241
							} else {

							}
						}
						{
							if x_20.Type == 9 && x_20.IntVal == 1830062304 {
								__t241 = (y_21.Type == 9 && y_21.IntVal == 1830062304) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M3)(x_20.UnsafePtr).V0, (*Constructor_Main_M3)(y_21.UnsafePtr).V0).IntVal) != (0))
								goto end_branch_241
							} else {

							}
						}
						{
							if x_20.Type == 9 && x_20.IntVal == 3190619783 {
								__t241 = (y_21.Type == 9 && y_21.IntVal == 3190619783) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_20.UnsafePtr).V0, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_21.UnsafePtr).V0, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_20.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_21.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_20.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_21.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_20.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_21.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_20.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_21.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_19_240.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_20.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_21.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_241
							} else {

							}
						}
						{
							if x_20.Type == 9 && x_20.IntVal == 108241190 {
								__t241 = (y_21.Type == 9 && y_21.IntVal == 108241190) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_20.UnsafePtr).V0, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_21.UnsafePtr).V0, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_20.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_21.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_20.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_21.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_20.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_21.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_20.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_21.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_19_240.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_20.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_21.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))
								goto end_branch_241
							} else {

							}
						}
						{
							if x_20.Type == 9 && x_20.IntVal == 2066233029 {
								__t241 = (y_21.Type == 9 && y_21.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_20.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_21.UnsafePtr).V0)) && (((*Constructor_Main_M6)(x_20.UnsafePtr).V1.StrVal()) == ((*Constructor_Main_M6)(y_21.UnsafePtr).V1.StrVal()))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(x_20.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}(), func() gopurs_runtime.Value {
									arr := (*Constructor_Main_M6)(y_21.UnsafePtr).V2
									boxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										boxed[i] = gopurs_runtime.Int(v)
									}
									return gopurs_runtime.Array(boxed)
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_19_240.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_20.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_21.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), (*Constructor_Main_M6)(x_20.UnsafePtr).V4, (*Constructor_Main_M6)(y_21.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), (*Constructor_Main_M6)(x_20.UnsafePtr).V5, (*Constructor_Main_M6)(y_21.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_20.UnsafePtr).V6, "a").StrVal()) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_21.UnsafePtr).V6, "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_20.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_21.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_20.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_21.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_20.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_21.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_20.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_21.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_19_240.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_20.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_21.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0)))) && (((((((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_20.UnsafePtr).V7, "nested"), "a").StrVal()) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_21.UnsafePtr).V7, "nested"), "a").StrVal())) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_20.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
									arr := func() []int64 {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_21.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
								}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqIntImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_20.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_21.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), Get_Data_Eq_eqStringImpl(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_20.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_21.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_20.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_21.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_19_240.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_20.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_21.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())).IntVal) != (0))))
								goto end_branch_241
							} else {

							}
						}
						{
							__t241 = (x_20.Type == 9 && x_20.IntVal == 1168316772) && ((y_21.Type == 9 && y_21.IntVal == 1168316772) && ((gopurs_runtime.Apply3(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(__local_var_18_239, "eq"), (*Constructor_Main_M7)(x_20.UnsafePtr).V0, (*Constructor_Main_M7)(y_21.UnsafePtr).V0).IntVal) != (0)))
						}
					end_branch_241:
						return gopurs_runtime.Bool(__t241)
					})
				}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t250.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Main_m7()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())).IntVal) != (0))), gopurs_runtime.Value{})
				_ = _dollar___unused_16_232
				return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
			})
		}()
	})
	return cache_Main_main
}

var cache_Main_sequenceStr__1451196854 gopurs_runtime.Value
var once_Main_sequenceStr__1451196854 sync.Once

func Get_Main_sequenceStr__1451196854() gopurs_runtime.Value {
	once_Main_sequenceStr__1451196854.Do(func() {
		cache_Main_sequenceStr__1451196854 = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_sequenceStr__1451196854(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box))
		})
	})
	return cache_Main_sequenceStr__1451196854
}

var cache_Main_traverseStr__2257532016 gopurs_runtime.Value
var once_Main_traverseStr__2257532016 sync.Once

func Get_Main_traverseStr__2257532016() gopurs_runtime.Value {
	once_Main_traverseStr__2257532016.Do(func() {
		cache_Main_traverseStr__2257532016 = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_traverseStr__2257532016(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box))
		})
	})
	return cache_Main_traverseStr__2257532016
}

type Constructor_Main_M0 struct {
	Rc uint32
}

type Constructor_Main_M1 struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 []gopurs_runtime.Value
}

type Constructor_Main_M2 struct {
	Rc uint32
	V0 int64
}

type Constructor_Main_M3 struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_M4 struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_M5 struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_M6 struct {
	Rc uint32
	V0 int64
	V1 gopurs_runtime.Value
	V2 []int64
	V3 []gopurs_runtime.Value
	V4 gopurs_runtime.Value
	V5 gopurs_runtime.Value
	V6 gopurs_runtime.Value
	V7 gopurs_runtime.Value
}

type Constructor_Main_M7 struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_functorM(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if m_2.Type == 9 && m_2.IntVal == 3852365315 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 769986722 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, gopurs_runtime.Apply(f_1, (*Constructor_Main_M1)(m_2.UnsafePtr).V0), func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
							arr_val_arrayMap5 := gopurs_runtime.Array((*Constructor_Main_M1)(m_2.UnsafePtr).V1)
							_ = arr_val_arrayMap5
							arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
							_ = arr_go_arrayMap5
							res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
							_ = res_go_arrayMap5
							for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
								res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(f_1, v_arrayMap5)
							}
							return gopurs_runtime.Array(res_go_arrayMap5)
						}().UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 2727978561 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2{1, (*Constructor_Main_M2)(m_2.UnsafePtr).V0}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 1830062304 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*Constructor_Main_M3)(m_2.UnsafePtr).V0)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 3190619783 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4)(m_2.UnsafePtr).V0, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_2.UnsafePtr).V0, "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_2.UnsafePtr).V0, "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
							arr_val_arrayMap6 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_2.UnsafePtr).V0, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap6
							arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
							_ = arr_go_arrayMap6
							res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
							_ = res_go_arrayMap6
							for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
								res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_1, v_arrayMap6)
							}
							return gopurs_runtime.Array(res_go_arrayMap6)
						}().UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()))}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 108241190 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, func() gopurs_runtime.Value {
						origVal := (*Constructor_Main_M5)(m_2.UnsafePtr).V0
						if origVal.Type != gopurs_runtime.TypeRecord1 {
							return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_2.UnsafePtr).V0, "nested"), "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_2.UnsafePtr).V0, "nested"), "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_2.UnsafePtr).V0, "nested"), "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap7 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_2.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap7
									arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
									_ = arr_go_arrayMap7
									res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
									_ = res_go_arrayMap7
									for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
										res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_1, v_arrayMap7)
									}
									return gopurs_runtime.Array(res_go_arrayMap7)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))})
						}
						clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
						clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_2.UnsafePtr).V0, "nested"), "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_2.UnsafePtr).V0, "nested"), "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_2.UnsafePtr).V0, "nested"), "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap7 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_2.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_arrayMap7
								arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
								_ = arr_go_arrayMap7
								res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
								_ = res_go_arrayMap7
								for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
									res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_1, v_arrayMap7)
								}
								return gopurs_runtime.Array(res_go_arrayMap7)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))
						return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
					}()}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 2066233029 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, (*Constructor_Main_M6)(m_2.UnsafePtr).V0, gopurs_runtime.Apply(f_1, (*Constructor_Main_M6)(m_2.UnsafePtr).V1), (*Constructor_Main_M6)(m_2.UnsafePtr).V2, func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
							arr_val_arrayMap5 := gopurs_runtime.Array((*Constructor_Main_M6)(m_2.UnsafePtr).V3)
							_ = arr_val_arrayMap5
							arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
							_ = arr_go_arrayMap5
							res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
							_ = res_go_arrayMap5
							for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
								res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(f_1, v_arrayMap5)
							}
							return gopurs_runtime.Array(res_go_arrayMap5)
						}().UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*Constructor_Main_M6)(m_2.UnsafePtr).V4), (*Constructor_Main_M6)(m_2.UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6)(m_2.UnsafePtr).V6, "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_2.UnsafePtr).V6, "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_2.UnsafePtr).V6, "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
							arr_val_arrayMap6 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_2.UnsafePtr).V6, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_arrayMap6
							arr_go_arrayMap6 := (*[]gopurs_runtime.Value)(arr_val_arrayMap6.UnsafePtr)
							_ = arr_go_arrayMap6
							res_go_arrayMap6 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap6))
							_ = res_go_arrayMap6
							for i_arrayMap6, v_arrayMap6 := range *arr_go_arrayMap6 {
								res_go_arrayMap6[i_arrayMap6] = gopurs_runtime.Apply(f_1, v_arrayMap6)
							}
							return gopurs_runtime.Array(res_go_arrayMap6)
						}().UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())), func() gopurs_runtime.Value {
						origVal := (*Constructor_Main_M6)(m_2.UnsafePtr).V7
						if origVal.Type != gopurs_runtime.TypeRecord1 {
							return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_2.UnsafePtr).V7, "nested"), "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_2.UnsafePtr).V7, "nested"), "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_2.UnsafePtr).V7, "nested"), "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap7 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_2.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap7
									arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
									_ = arr_go_arrayMap7
									res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
									_ = res_go_arrayMap7
									for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
										res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_1, v_arrayMap7)
									}
									return gopurs_runtime.Array(res_go_arrayMap7)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))})
						}
						clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
						clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_2.UnsafePtr).V7, "nested"), "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_2.UnsafePtr).V7, "nested"), "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_2.UnsafePtr).V7, "nested"), "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap7 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_2.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_arrayMap7
								arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
								_ = arr_go_arrayMap7
								res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
								_ = res_go_arrayMap7
								for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
									res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_1, v_arrayMap7)
								}
								return gopurs_runtime.Array(res_go_arrayMap7)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))
						return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
					}()}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 1168316772 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.RecordUpdate1(v1_3, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_3, "nested"), "a", gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap10 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_arrayMap10
								arr_go_arrayMap10 := (*[]gopurs_runtime.Value)(arr_val_arrayMap10.UnsafePtr)
								_ = arr_go_arrayMap10
								res_go_arrayMap10 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap10))
								_ = res_go_arrayMap10
								for i_arrayMap10, v_arrayMap10 := range *arr_go_arrayMap10 {
									res_go_arrayMap10[i_arrayMap10] = gopurs_runtime.Apply(f_1, v_arrayMap10)
								}
								return gopurs_runtime.Array(res_go_arrayMap10)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())))
					})), (*Constructor_Main_M7)(m_2.UnsafePtr).V0)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		})
	})}))}
}

func Call_Main_foldableM(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
	_ = dictFoldable_0
	return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): mempty_2_0 -> gopurs_runtime.Value
		mempty_2_0 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
		_ = mempty_2_0
		// TAST (Let): Semigroup0_3_1 -> *Constructor_Data_Semigroup_Semigroup
		Semigroup0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
		_ = Semigroup0_3_1
		return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t9 gopurs_runtime.Value
				{
					if m_5.Type == 9 && m_5.IntVal == 3852365315 {
						__t9 = mempty_2_0
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 769986722 {
						// TAST (Let): Semigroup0_6_2 -> *Constructor_Data_Semigroup_Semigroup
						Semigroup0_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
						_ = Semigroup0_6_2
						__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, (*Constructor_Main_M1)(m_5.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_2.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
							})
						}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array((*Constructor_Main_M1)(m_5.UnsafePtr).V1)))
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 2727978561 {
						__t9 = mempty_2_0
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 1830062304 {
						__t9 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, f_4, (*Constructor_Main_M3)(m_5.UnsafePtr).V0)
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 3190619783 {
						// TAST (Let): Semigroup0_6_3 -> *Constructor_Data_Semigroup_Semigroup
						Semigroup0_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
						_ = Semigroup0_6_3
						__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_5.UnsafePtr).V0, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_5.UnsafePtr).V0, "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_3.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
							})
						}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_5.UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))))
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 108241190 {
						// TAST (Let): Semigroup0_6_4 -> *Constructor_Data_Semigroup_Semigroup
						Semigroup0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
						_ = Semigroup0_6_4
						__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_5.UnsafePtr).V0, "nested"), "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_5.UnsafePtr).V0, "nested"), "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_4.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
							})
						}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_5.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))))
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 2066233029 {
						// TAST (Let): Semigroup0_6_5 -> *Constructor_Data_Semigroup_Semigroup
						Semigroup0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
						_ = Semigroup0_6_5
						// TAST (Let): Semigroup0_6_6 -> *Constructor_Data_Semigroup_Semigroup
						Semigroup0_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
						_ = Semigroup0_6_6
						// TAST (Let): Semigroup0_6_7 -> *Constructor_Data_Semigroup_Semigroup
						Semigroup0_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
						_ = Semigroup0_6_7
						__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, (*Constructor_Main_M6)(m_5.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_5.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
							})
						}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array((*Constructor_Main_M6)(m_5.UnsafePtr).V3)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, f_4, (*Constructor_Main_M6)(m_5.UnsafePtr).V4), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V6, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V6, "fa")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_6.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
							})
						}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V7, "nested"), "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V7, "nested"), "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(acc_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_7.V0), gopurs_runtime.Apply(f_4, x_7), acc_8)
							})
						}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))))))))))
						goto end_branch_9
					} else {

					}
				}
				{
					if m_5.Type == 9 && m_5.IntVal == 1168316772 {
						__t9 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
							// TAST (Let): Semigroup0_7_8 -> *Constructor_Data_Semigroup_Semigroup
							Semigroup0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_7_8
							return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_1.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, f_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_8.V0), gopurs_runtime.Apply(f_4, x_8), acc_9)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))))
						})), (*Constructor_Main_M7)(m_5.UnsafePtr).V0)
						goto end_branch_9
					} else {

					}
				}
				{
					__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_9:
				return __t9
			})
		})
	}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t10 gopurs_runtime.Value
				{
					if m_3.Type == 9 && m_3.IntVal == 3852365315 {
						__t10 = z_2
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 769986722 {
						__t10 = func() gopurs_runtime.Value {
							arr_val_foldlArray5 := gopurs_runtime.Array((*Constructor_Main_M1)(m_3.UnsafePtr).V1)
							_ = arr_val_foldlArray5
							res_go_foldlArray5 := gopurs_runtime.Apply2(f_1, z_2, (*Constructor_Main_M1)(m_3.UnsafePtr).V0)
							_ = res_go_foldlArray5
							arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
							_ = arr_go_foldlArray5
							for _, v_foldlArray5 := range *arr_go_foldlArray5 {
								res_go_foldlArray5 = gopurs_runtime.Apply2(f_1, res_go_foldlArray5, v_foldlArray5)
							}
							return res_go_foldlArray5
						}()
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 2727978561 {
						__t10 = z_2
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 1830062304 {
						__t10 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, z_2, (*Constructor_Main_M3)(m_3.UnsafePtr).V0)
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 3190619783 {
						__t10 = func() gopurs_runtime.Value {
							arr_val_foldlArray5 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_foldlArray5
							res_go_foldlArray5 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, z_2, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "fa"))
							_ = res_go_foldlArray5
							arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
							_ = arr_go_foldlArray5
							for _, v_foldlArray5 := range *arr_go_foldlArray5 {
								res_go_foldlArray5 = gopurs_runtime.Apply2(f_1, res_go_foldlArray5, v_foldlArray5)
							}
							return res_go_foldlArray5
						}()
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 108241190 {
						__t10 = func() gopurs_runtime.Value {
							arr_val_foldlArray5 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_foldlArray5
							res_go_foldlArray5 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, z_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "a")), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "fa"))
							_ = res_go_foldlArray5
							arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
							_ = arr_go_foldlArray5
							for _, v_foldlArray5 := range *arr_go_foldlArray5 {
								res_go_foldlArray5 = gopurs_runtime.Apply2(f_1, res_go_foldlArray5, v_foldlArray5)
							}
							return res_go_foldlArray5
						}()
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 2066233029 {
						__t10 = func() gopurs_runtime.Value {
							arr_val_foldlArray5 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())
							_ = arr_val_foldlArray5
							res_go_foldlArray5 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, func() gopurs_runtime.Value {
								arr_val_foldlArray8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_foldlArray8
								res_go_foldlArray8 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, func() gopurs_runtime.Value {
									arr_val_foldlArray12 := gopurs_runtime.Array((*Constructor_Main_M6)(m_3.UnsafePtr).V3)
									_ = arr_val_foldlArray12
									res_go_foldlArray12 := gopurs_runtime.Apply2(f_1, z_2, (*Constructor_Main_M6)(m_3.UnsafePtr).V1)
									_ = res_go_foldlArray12
									arr_go_foldlArray12 := (*[]gopurs_runtime.Value)(arr_val_foldlArray12.UnsafePtr)
									_ = arr_go_foldlArray12
									for _, v_foldlArray12 := range *arr_go_foldlArray12 {
										res_go_foldlArray12 = gopurs_runtime.Apply2(f_1, res_go_foldlArray12, v_foldlArray12)
									}
									return res_go_foldlArray12
								}(), (*Constructor_Main_M6)(m_3.UnsafePtr).V4), gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "fa"))
								_ = res_go_foldlArray8
								arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
								_ = arr_go_foldlArray8
								for _, v_foldlArray8 := range *arr_go_foldlArray8 {
									res_go_foldlArray8 = gopurs_runtime.Apply2(f_1, res_go_foldlArray8, v_foldlArray8)
								}
								return res_go_foldlArray8
							}(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "a")), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "fa"))
							_ = res_go_foldlArray5
							arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
							_ = arr_go_foldlArray5
							for _, v_foldlArray5 := range *arr_go_foldlArray5 {
								res_go_foldlArray5 = gopurs_runtime.Apply2(f_1, res_go_foldlArray5, v_foldlArray5)
							}
							return res_go_foldlArray5
						}()
						goto end_branch_10
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 1168316772 {
						__t10 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return func() gopurs_runtime.Value {
									arr_val_foldlArray9 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_foldlArray9
									res_go_foldlArray9 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, v1_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "a")), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_5, "nested"), "fa"))
									_ = res_go_foldlArray9
									arr_go_foldlArray9 := (*[]gopurs_runtime.Value)(arr_val_foldlArray9.UnsafePtr)
									_ = arr_go_foldlArray9
									for _, v_foldlArray9 := range *arr_go_foldlArray9 {
										res_go_foldlArray9 = gopurs_runtime.Apply2(f_1, res_go_foldlArray9, v_foldlArray9)
									}
									return res_go_foldlArray9
								}()
							})
						})), z_2, (*Constructor_Main_M7)(m_3.UnsafePtr).V0)
						goto end_branch_10
					} else {

					}
				}
				{
					__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_10:
				return __t10
			})
		})
	}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t12 gopurs_runtime.Value
				{
					if m_3.Type == 9 && m_3.IntVal == 3852365315 {
						__t12 = z_2
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 769986722 {
						__t12 = gopurs_runtime.Apply2(f_1, (*Constructor_Main_M1)(m_3.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array((*Constructor_Main_M1)(m_3.UnsafePtr).V1)))
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 2727978561 {
						__t12 = z_2
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 1830062304 {
						__t12 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, z_2, (*Constructor_Main_M3)(m_3.UnsafePtr).V0)
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 3190619783 {
						__t12 = gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "fa")))
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 108241190 {
						__t12 = gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "fa")))
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 2066233029 {
						__t12 = gopurs_runtime.Apply2(f_1, (*Constructor_Main_M6)(m_3.UnsafePtr).V1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, z_2, gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "fa"))), gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "zArrayA").UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "fa"))), (*Constructor_Main_M6)(m_3.UnsafePtr).V4), gopurs_runtime.Array((*Constructor_Main_M6)(m_3.UnsafePtr).V3)))
						goto end_branch_12
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 1168316772 {
						// TAST (Let): __local_var_4_11 -> gopurs_runtime.Value
						__local_var_4_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, v2_5, gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "fa")))
							})
						}))
						_ = __local_var_4_11
						__t12 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply2(__local_var_4_11, a_6, b_5)
							})
						}), z_2, (*Constructor_Main_M7)(m_3.UnsafePtr).V0)
						goto end_branch_12
					} else {

					}
				}
				{
					__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_12:
				return __t12
			})
		})
	})}))}
}

func Call_Main_traversableM(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
traversableM:
	for {
		if false {
			continue traversableM
		}
		var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
		_ = dictTraversable_0
		// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
		__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
		_ = __local_var_1_1
		// TAST (Let): functorM1_1_0 -> *Constructor_Data_Functor_Functor
		functorM1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t2 gopurs_runtime.Value
				{
					if m_3.Type == 9 && m_3.IntVal == 3852365315 {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)}
						goto end_branch_2
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 769986722 {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, gopurs_runtime.Apply(f_2, (*Constructor_Main_M1)(m_3.UnsafePtr).V0), func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap7 := gopurs_runtime.Array((*Constructor_Main_M1)(m_3.UnsafePtr).V1)
								_ = arr_val_arrayMap7
								arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
								_ = arr_go_arrayMap7
								res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
								_ = res_go_arrayMap7
								for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
									res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_2, v_arrayMap7)
								}
								return gopurs_runtime.Array(res_go_arrayMap7)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()}))}
						goto end_branch_2
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 2727978561 {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2{1, (*Constructor_Main_M2)(m_3.UnsafePtr).V0}))}
						goto end_branch_2
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 1830062304 {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, (*Constructor_Main_M3)(m_3.UnsafePtr).V0)}))}
						goto end_branch_2
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 3190619783 {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "a", gopurs_runtime.Apply(f_2, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_3.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_arrayMap8
								arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
								_ = arr_go_arrayMap8
								res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
								_ = res_go_arrayMap8
								for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
									res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(f_2, v_arrayMap8)
								}
								return gopurs_runtime.Array(res_go_arrayMap8)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}()))}))}
						goto end_branch_2
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 108241190 {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, func() gopurs_runtime.Value {
							origVal := (*Constructor_Main_M5)(m_3.UnsafePtr).V0
							if origVal.Type != gopurs_runtime.TypeRecord1 {
								return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "a", gopurs_runtime.Apply(f_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
										arr_val_arrayMap9 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}())
										_ = arr_val_arrayMap9
										arr_go_arrayMap9 := (*[]gopurs_runtime.Value)(arr_val_arrayMap9.UnsafePtr)
										_ = arr_go_arrayMap9
										res_go_arrayMap9 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap9))
										_ = res_go_arrayMap9
										for i_arrayMap9, v_arrayMap9 := range *arr_go_arrayMap9 {
											res_go_arrayMap9[i_arrayMap9] = gopurs_runtime.Apply(f_2, v_arrayMap9)
										}
										return gopurs_runtime.Array(res_go_arrayMap9)
									}().UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()))})
							}
							clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
							clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "a", gopurs_runtime.Apply(f_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap9 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_3.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap9
									arr_go_arrayMap9 := (*[]gopurs_runtime.Value)(arr_val_arrayMap9.UnsafePtr)
									_ = arr_go_arrayMap9
									res_go_arrayMap9 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap9))
									_ = res_go_arrayMap9
									for i_arrayMap9, v_arrayMap9 := range *arr_go_arrayMap9 {
										res_go_arrayMap9[i_arrayMap9] = gopurs_runtime.Apply(f_2, v_arrayMap9)
									}
									return gopurs_runtime.Array(res_go_arrayMap9)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))
							return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
						}()}))}
						goto end_branch_2
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 2066233029 {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, (*Constructor_Main_M6)(m_3.UnsafePtr).V0, gopurs_runtime.Apply(f_2, (*Constructor_Main_M6)(m_3.UnsafePtr).V1), (*Constructor_Main_M6)(m_3.UnsafePtr).V2, func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap7 := gopurs_runtime.Array((*Constructor_Main_M6)(m_3.UnsafePtr).V3)
								_ = arr_val_arrayMap7
								arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
								_ = arr_go_arrayMap7
								res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
								_ = res_go_arrayMap7
								for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
									res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_2, v_arrayMap7)
								}
								return gopurs_runtime.Array(res_go_arrayMap7)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, (*Constructor_Main_M6)(m_3.UnsafePtr).V4), (*Constructor_Main_M6)(m_3.UnsafePtr).V5, gopurs_runtime.RecordUpdate3((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "a", gopurs_runtime.Apply(f_2, gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
							arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
								arr_val_arrayMap8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V6, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_arrayMap8
								arr_go_arrayMap8 := (*[]gopurs_runtime.Value)(arr_val_arrayMap8.UnsafePtr)
								_ = arr_go_arrayMap8
								res_go_arrayMap8 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap8))
								_ = res_go_arrayMap8
								for i_arrayMap8, v_arrayMap8 := range *arr_go_arrayMap8 {
									res_go_arrayMap8[i_arrayMap8] = gopurs_runtime.Apply(f_2, v_arrayMap8)
								}
								return gopurs_runtime.Array(res_go_arrayMap8)
							}().UnsafePtr)
							unboxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								unboxed[i] = v
							}
							return unboxed
						}())), func() gopurs_runtime.Value {
							origVal := (*Constructor_Main_M6)(m_3.UnsafePtr).V7
							if origVal.Type != gopurs_runtime.TypeRecord1 {
								return gopurs_runtime.RecordUpdateDict(origVal, []string{"nested"}, []gopurs_runtime.Value{gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "a", gopurs_runtime.Apply(f_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
										arr_val_arrayMap9 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}())
										_ = arr_val_arrayMap9
										arr_go_arrayMap9 := (*[]gopurs_runtime.Value)(arr_val_arrayMap9.UnsafePtr)
										_ = arr_go_arrayMap9
										res_go_arrayMap9 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap9))
										_ = res_go_arrayMap9
										for i_arrayMap9, v_arrayMap9 := range *arr_go_arrayMap9 {
											res_go_arrayMap9[i_arrayMap9] = gopurs_runtime.Apply(f_2, v_arrayMap9)
										}
										return gopurs_runtime.Array(res_go_arrayMap9)
									}().UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()))})
							}
							clone := *((*gopurs_runtime.RecordData1)(origVal.UnsafePtr))
							clone.V0 = gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "a", gopurs_runtime.Apply(f_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap9 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_3.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap9
									arr_go_arrayMap9 := (*[]gopurs_runtime.Value)(arr_val_arrayMap9.UnsafePtr)
									_ = arr_go_arrayMap9
									res_go_arrayMap9 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap9))
									_ = res_go_arrayMap9
									for i_arrayMap9, v_arrayMap9 := range *arr_go_arrayMap9 {
										res_go_arrayMap9[i_arrayMap9] = gopurs_runtime.Apply(f_2, v_arrayMap9)
									}
									return gopurs_runtime.Array(res_go_arrayMap9)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))
							return gopurs_runtime.Value{Type: gopurs_runtime.TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
						}()}))}
						goto end_branch_2
					} else {

					}
				}
				{
					if m_3.Type == 9 && m_3.IntVal == 1168316772 {
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.RecordUpdate1(v1_4, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_4, "nested"), "a", gopurs_runtime.Apply(f_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "a")), "fa", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "fa")), "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
									arr_val_arrayMap12 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_4, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_arrayMap12
									arr_go_arrayMap12 := (*[]gopurs_runtime.Value)(arr_val_arrayMap12.UnsafePtr)
									_ = arr_go_arrayMap12
									res_go_arrayMap12 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap12))
									_ = res_go_arrayMap12
									for i_arrayMap12, v_arrayMap12 := range *arr_go_arrayMap12 {
										res_go_arrayMap12[i_arrayMap12] = gopurs_runtime.Apply(f_2, v_arrayMap12)
									}
									return gopurs_runtime.Array(res_go_arrayMap12)
								}().UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())))
						})), (*Constructor_Main_M7)(m_3.UnsafePtr).V0)}))}
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_2:
				return __t2
			})
		})))
		_ = functorM1_1_0
		// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
		__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
		_ = __local_var_2_4
		// TAST (Let): foldableM1_2_3 -> *Constructor_Data_Foldable_Foldable
		foldableM1_2_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): mempty_4_5 -> gopurs_runtime.Value
			mempty_4_5 := gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
			_ = mempty_4_5
			// TAST (Let): Semigroup0_5_6 -> *Constructor_Data_Semigroup_Semigroup
			Semigroup0_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
			_ = Semigroup0_5_6
			return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t14 gopurs_runtime.Value
					{
						if m_7.Type == 9 && m_7.IntVal == 3852365315 {
							__t14 = mempty_4_5
							goto end_branch_14
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 769986722 {
							// TAST (Let): Semigroup0_8_7 -> *Constructor_Data_Semigroup_Semigroup
							Semigroup0_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_8_7
							__t14 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply(f_6, (*Constructor_Main_M1)(m_7.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_7.V0), gopurs_runtime.Apply(f_6, x_9), acc_10)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), gopurs_runtime.Array((*Constructor_Main_M1)(m_7.UnsafePtr).V1)))
							goto end_branch_14
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 2727978561 {
							__t14 = mempty_4_5
							goto end_branch_14
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 1830062304 {
							__t14 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, f_6, (*Constructor_Main_M3)(m_7.UnsafePtr).V0)
							goto end_branch_14
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 3190619783 {
							// TAST (Let): Semigroup0_8_8 -> *Constructor_Data_Semigroup_Semigroup
							Semigroup0_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_8_8
							__t14 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_7.UnsafePtr).V0, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, f_6, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_7.UnsafePtr).V0, "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_8.V0), gopurs_runtime.Apply(f_6, x_9), acc_10)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_7.UnsafePtr).V0, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))))
							goto end_branch_14
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 108241190 {
							// TAST (Let): Semigroup0_8_9 -> *Constructor_Data_Semigroup_Semigroup
							Semigroup0_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_8_9
							__t14 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_7.UnsafePtr).V0, "nested"), "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_7.UnsafePtr).V0, "nested"), "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_9.V0), gopurs_runtime.Apply(f_6, x_9), acc_10)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_7.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))))
							goto end_branch_14
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 2066233029 {
							// TAST (Let): Semigroup0_8_10 -> *Constructor_Data_Semigroup_Semigroup
							Semigroup0_8_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_8_10
							// TAST (Let): Semigroup0_8_11 -> *Constructor_Data_Semigroup_Semigroup
							Semigroup0_8_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_8_11
							// TAST (Let): Semigroup0_8_12 -> *Constructor_Data_Semigroup_Semigroup
							Semigroup0_8_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
							_ = Semigroup0_8_12
							__t14 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply(f_6, (*Constructor_Main_M6)(m_7.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_10.V0), gopurs_runtime.Apply(f_6, x_9), acc_10)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), gopurs_runtime.Array((*Constructor_Main_M6)(m_7.UnsafePtr).V3)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, f_6, (*Constructor_Main_M6)(m_7.UnsafePtr).V4), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_7.UnsafePtr).V6, "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, f_6, gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_7.UnsafePtr).V6, "fa")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_11.V0), gopurs_runtime.Apply(f_6, x_9), acc_10)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_7.UnsafePtr).V6, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_7.UnsafePtr).V7, "nested"), "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_7.UnsafePtr).V7, "nested"), "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(acc_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_12.V0), gopurs_runtime.Apply(f_6, x_9), acc_10)
								})
							}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_7.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))))))))))
							goto end_branch_14
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 1168316772 {
							__t14 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_4, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
								// TAST (Let): Semigroup0_9_13 -> *Constructor_Data_Semigroup_Semigroup
								Semigroup0_9_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
								_ = Semigroup0_9_13
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "a")), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_6.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "fa")), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(acc_11 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_13.V0), gopurs_runtime.Apply(f_6, x_10), acc_11)
									})
								}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()))))
							})), (*Constructor_Main_M7)(m_7.UnsafePtr).V0)
							goto end_branch_14
						} else {

						}
					}
					{
						__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_14:
					return __t14
				})
			})
		}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(z_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t15 gopurs_runtime.Value
					{
						if m_5.Type == 9 && m_5.IntVal == 3852365315 {
							__t15 = z_4
							goto end_branch_15
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 769986722 {
							__t15 = func() gopurs_runtime.Value {
								arr_val_foldlArray8 := gopurs_runtime.Array((*Constructor_Main_M1)(m_5.UnsafePtr).V1)
								_ = arr_val_foldlArray8
								res_go_foldlArray8 := gopurs_runtime.Apply2(f_3, z_4, (*Constructor_Main_M1)(m_5.UnsafePtr).V0)
								_ = res_go_foldlArray8
								arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
								_ = arr_go_foldlArray8
								for _, v_foldlArray8 := range *arr_go_foldlArray8 {
									res_go_foldlArray8 = gopurs_runtime.Apply2(f_3, res_go_foldlArray8, v_foldlArray8)
								}
								return res_go_foldlArray8
							}()
							goto end_branch_15
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 2727978561 {
							__t15 = z_4
							goto end_branch_15
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 1830062304 {
							__t15 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldl"), f_3, z_4, (*Constructor_Main_M3)(m_5.UnsafePtr).V0)
							goto end_branch_15
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 3190619783 {
							__t15 = func() gopurs_runtime.Value {
								arr_val_foldlArray8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_5.UnsafePtr).V0, "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_foldlArray8
								res_go_foldlArray8 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldl"), f_3, gopurs_runtime.Apply2(f_3, z_4, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_5.UnsafePtr).V0, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_5.UnsafePtr).V0, "fa"))
								_ = res_go_foldlArray8
								arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
								_ = arr_go_foldlArray8
								for _, v_foldlArray8 := range *arr_go_foldlArray8 {
									res_go_foldlArray8 = gopurs_runtime.Apply2(f_3, res_go_foldlArray8, v_foldlArray8)
								}
								return res_go_foldlArray8
							}()
							goto end_branch_15
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 108241190 {
							__t15 = func() gopurs_runtime.Value {
								arr_val_foldlArray8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_5.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_foldlArray8
								res_go_foldlArray8 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldl"), f_3, gopurs_runtime.Apply2(f_3, z_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_5.UnsafePtr).V0, "nested"), "a")), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_5.UnsafePtr).V0, "nested"), "fa"))
								_ = res_go_foldlArray8
								arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
								_ = arr_go_foldlArray8
								for _, v_foldlArray8 := range *arr_go_foldlArray8 {
									res_go_foldlArray8 = gopurs_runtime.Apply2(f_3, res_go_foldlArray8, v_foldlArray8)
								}
								return res_go_foldlArray8
							}()
							goto end_branch_15
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 2066233029 {
							__t15 = func() gopurs_runtime.Value {
								arr_val_foldlArray8 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())
								_ = arr_val_foldlArray8
								res_go_foldlArray8 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldl"), f_3, gopurs_runtime.Apply2(f_3, func() gopurs_runtime.Value {
									arr_val_foldlArray11 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V6, "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())
									_ = arr_val_foldlArray11
									res_go_foldlArray11 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldl"), f_3, gopurs_runtime.Apply2(f_3, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldl"), f_3, func() gopurs_runtime.Value {
										arr_val_foldlArray15 := gopurs_runtime.Array((*Constructor_Main_M6)(m_5.UnsafePtr).V3)
										_ = arr_val_foldlArray15
										res_go_foldlArray15 := gopurs_runtime.Apply2(f_3, z_4, (*Constructor_Main_M6)(m_5.UnsafePtr).V1)
										_ = res_go_foldlArray15
										arr_go_foldlArray15 := (*[]gopurs_runtime.Value)(arr_val_foldlArray15.UnsafePtr)
										_ = arr_go_foldlArray15
										for _, v_foldlArray15 := range *arr_go_foldlArray15 {
											res_go_foldlArray15 = gopurs_runtime.Apply2(f_3, res_go_foldlArray15, v_foldlArray15)
										}
										return res_go_foldlArray15
									}(), (*Constructor_Main_M6)(m_5.UnsafePtr).V4), gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V6, "a")), gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V6, "fa"))
									_ = res_go_foldlArray11
									arr_go_foldlArray11 := (*[]gopurs_runtime.Value)(arr_val_foldlArray11.UnsafePtr)
									_ = arr_go_foldlArray11
									for _, v_foldlArray11 := range *arr_go_foldlArray11 {
										res_go_foldlArray11 = gopurs_runtime.Apply2(f_3, res_go_foldlArray11, v_foldlArray11)
									}
									return res_go_foldlArray11
								}(), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V7, "nested"), "a")), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V7, "nested"), "fa"))
								_ = res_go_foldlArray8
								arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
								_ = arr_go_foldlArray8
								for _, v_foldlArray8 := range *arr_go_foldlArray8 {
									res_go_foldlArray8 = gopurs_runtime.Apply2(f_3, res_go_foldlArray8, v_foldlArray8)
								}
								return res_go_foldlArray8
							}()
							goto end_branch_15
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 1168316772 {
							__t15 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "foldl"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return func() gopurs_runtime.Value {
										arr_val_foldlArray12 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_7, "nested"), "zArrayA").UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}())
										_ = arr_val_foldlArray12
										res_go_foldlArray12 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldl"), f_3, gopurs_runtime.Apply2(f_3, v1_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_7, "nested"), "a")), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v2_7, "nested"), "fa"))
										_ = res_go_foldlArray12
										arr_go_foldlArray12 := (*[]gopurs_runtime.Value)(arr_val_foldlArray12.UnsafePtr)
										_ = arr_go_foldlArray12
										for _, v_foldlArray12 := range *arr_go_foldlArray12 {
											res_go_foldlArray12 = gopurs_runtime.Apply2(f_3, res_go_foldlArray12, v_foldlArray12)
										}
										return res_go_foldlArray12
									}()
								})
							})), z_4, (*Constructor_Main_M7)(m_5.UnsafePtr).V0)
							goto end_branch_15
						} else {

						}
					}
					{
						__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_15:
					return __t15
				})
			})
		}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(z_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t17 gopurs_runtime.Value
					{
						if m_5.Type == 9 && m_5.IntVal == 3852365315 {
							__t17 = z_4
							goto end_branch_17
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 769986722 {
							__t17 = gopurs_runtime.Apply2(f_3, (*Constructor_Main_M1)(m_5.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_3, z_4, gopurs_runtime.Array((*Constructor_Main_M1)(m_5.UnsafePtr).V1)))
							goto end_branch_17
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 2727978561 {
							__t17 = z_4
							goto end_branch_17
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 1830062304 {
							__t17 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldr"), f_3, z_4, (*Constructor_Main_M3)(m_5.UnsafePtr).V0)
							goto end_branch_17
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 3190619783 {
							__t17 = gopurs_runtime.Apply2(f_3, gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_5.UnsafePtr).V0, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldr"), f_3, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_3, z_4, gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_5.UnsafePtr).V0, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())), gopurs_runtime.RecordGet((*Constructor_Main_M4)(m_5.UnsafePtr).V0, "fa")))
							goto end_branch_17
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 108241190 {
							__t17 = gopurs_runtime.Apply2(f_3, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_5.UnsafePtr).V0, "nested"), "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldr"), f_3, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_3, z_4, gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_5.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(m_5.UnsafePtr).V0, "nested"), "fa")))
							goto end_branch_17
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 2066233029 {
							__t17 = gopurs_runtime.Apply2(f_3, (*Constructor_Main_M6)(m_5.UnsafePtr).V1, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_3, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldr"), f_3, gopurs_runtime.Apply2(f_3, gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V6, "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldr"), f_3, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_3, gopurs_runtime.Apply2(f_3, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V7, "nested"), "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldr"), f_3, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_3, z_4, gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V7, "nested"), "fa"))), gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V6, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())), gopurs_runtime.RecordGet((*Constructor_Main_M6)(m_5.UnsafePtr).V6, "fa"))), (*Constructor_Main_M6)(m_5.UnsafePtr).V4), gopurs_runtime.Array((*Constructor_Main_M6)(m_5.UnsafePtr).V3)))
							goto end_branch_17
						} else {

						}
					}
					{
						if m_5.Type == 9 && m_5.IntVal == 1168316772 {
							// TAST (Let): __local_var_6_16 -> gopurs_runtime.Value
							__local_var_6_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "foldr"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(f_3, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "a"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldr"), f_3, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_3, v2_7, gopurs_runtime.Array(func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "zArrayA").UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}())), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_6, "nested"), "fa")))
								})
							}))
							_ = __local_var_6_16
							__t17 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "foldr"), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(__local_var_6_16, a_8, b_7)
								})
							}), z_4, (*Constructor_Main_M7)(m_5.UnsafePtr).V0)
							goto end_branch_17
						} else {

						}
					}
					{
						__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_17:
					return __t17
				})
			})
		})))
		_ = foldableM1_2_3
		return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableM1_2_3)}
		}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorM1_1_0)}
		}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Main_traversableM(dictTraversable_0), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
					return x_5
				}), v_4)
			})
		}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Apply0_4_18 -> *Constructor_Control_Apply_Apply
			Apply0_4_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_4_18
			// TAST (Let): Functor0_5_19 -> *Constructor_Data_Functor_Functor
			Functor0_5_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_5_19
			return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
					var __t34 gopurs_runtime.Value
					{
						if m_7.Type == 9 && m_7.IntVal == 3852365315 {
							__t34 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3852365315, UnsafePtr: unsafe.Pointer(nil)})
							goto end_branch_34
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 769986722 {
							// TAST (Let): Apply0_8_20 -> gopurs_runtime.Value
							Apply0_8_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
							_ = Apply0_8_20
							__t34 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_19.V0), gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 769986722, UnsafePtr: unsafe.Pointer((&Constructor_Main_M1{1, v2_8, func() []gopurs_runtime.Value {
										arr := *(*[]gopurs_runtime.Value)(v3_9.UnsafePtr)
										unboxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											unboxed[i] = v
										}
										return unboxed
									}()}))}
								})
							}), gopurs_runtime.Apply(f_6, (*Constructor_Main_M1)(m_7.UnsafePtr).V0)), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_8_20, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_20, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_3, "pure"), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array((*Constructor_Main_M1)(m_7.UnsafePtr).V1)))
							goto end_branch_34
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 2727978561 {
							__t34 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2727978561, UnsafePtr: unsafe.Pointer((&Constructor_Main_M2{1, (*Constructor_Main_M2)(m_7.UnsafePtr).V0}))})
							goto end_branch_34
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 1830062304 {
							__t34 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_19.V0), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Value{Type: 9, IntVal: 1830062304, UnsafePtr: unsafe.Pointer((&Constructor_Main_M3{1, v1_8}))}
							}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, f_6, (*Constructor_Main_M3)(m_7.UnsafePtr).V0))
							goto end_branch_34
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 3190619783 {
							// TAST (Let): __local_var_8_21 -> gopurs_runtime.Value
							__local_var_8_21 := (*Constructor_Main_M4)(m_7.UnsafePtr).V0
							_ = __local_var_8_21
							// TAST (Let): Apply0_9_22 -> gopurs_runtime.Value
							Apply0_9_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
							_ = Apply0_9_22
							__t34 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_19.V0), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 3190619783, UnsafePtr: unsafe.Pointer((&Constructor_Main_M4{1, gopurs_runtime.RecordUpdate3(__local_var_8_21, "a", v1_9, "fa", v2_10, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v3_11.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}()))}))}
									})
								})
							}), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(__local_var_8_21, "a"))), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(__local_var_8_21, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_9_22, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_22, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_3, "pure"), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_8_21, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())))
							goto end_branch_34
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 108241190 {
							// TAST (Let): __local_var_8_23 -> gopurs_runtime.Value
							var __local_var_8_23 gopurs_runtime.Value = (*Constructor_Main_M5)(m_7.UnsafePtr).V0
							// TAST (Let): Apply0_9_24 -> gopurs_runtime.Value
							Apply0_9_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
							_ = Apply0_9_24
							__t34 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_19.V0), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Value{Type: 9, IntVal: 108241190, UnsafePtr: unsafe.Pointer((&Constructor_Main_M5{1, gopurs_runtime.RecordUpdate1(__local_var_8_23, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(__local_var_8_23, "nested"), "a", v1_9, "fa", v2_10, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
											arr := *(*[]gopurs_runtime.Value)(v3_11.UnsafePtr)
											unboxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												unboxed[i] = v
											}
											return unboxed
										}())))}))}
									})
								})
							}), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_8_23, "nested"), "a"))), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_8_23, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_9_24, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_24, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_3, "pure"), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_8_23, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())))
							goto end_branch_34
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 2066233029 {
							// TAST (Let): __local_var_8_25 -> gopurs_runtime.Value
							var __local_var_8_25 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Main_M6)(m_7.UnsafePtr).V0)
							// TAST (Let): __local_var_9_26 -> gopurs_runtime.Value
							var __local_var_9_26 gopurs_runtime.Value = func() gopurs_runtime.Value {
								arr := (*Constructor_Main_M6)(m_7.UnsafePtr).V2
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}()
							// TAST (Let): __local_var_10_27 -> gopurs_runtime.Value
							__local_var_10_27 := (*Constructor_Main_M6)(m_7.UnsafePtr).V5
							_ = __local_var_10_27
							// TAST (Let): __local_var_11_28 -> gopurs_runtime.Value
							__local_var_11_28 := (*Constructor_Main_M6)(m_7.UnsafePtr).V6
							_ = __local_var_11_28
							// TAST (Let): __local_var_12_29 -> gopurs_runtime.Value
							var __local_var_12_29 gopurs_runtime.Value = (*Constructor_Main_M6)(m_7.UnsafePtr).V7
							// TAST (Let): Apply0_13_30 -> gopurs_runtime.Value
							Apply0_13_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
							_ = Apply0_13_30
							// TAST (Let): Apply0_13_31 -> gopurs_runtime.Value
							Apply0_13_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
							_ = Apply0_13_31
							// TAST (Let): Apply0_13_32 -> gopurs_runtime.Value
							Apply0_13_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
							_ = Apply0_13_32
							__t34 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_19.V0), gopurs_runtime.Func(func(v8_13 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v9_14 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v10_15 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v11_16 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(v12_17 gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Func(func(v13_18 gopurs_runtime.Value) gopurs_runtime.Value {
													return gopurs_runtime.Func(func(v14_19 gopurs_runtime.Value) gopurs_runtime.Value {
														return gopurs_runtime.Func(func(v15_20 gopurs_runtime.Value) gopurs_runtime.Value {
															return gopurs_runtime.Func(func(v16_21 gopurs_runtime.Value) gopurs_runtime.Value {
																return gopurs_runtime.Value{Type: 9, IntVal: 2066233029, UnsafePtr: unsafe.Pointer((&Constructor_Main_M6{1, __local_var_8_25.IntVal, v8_13, func() []int64 {
																	arr := *(*[]gopurs_runtime.Value)(__local_var_9_26.UnsafePtr)
																	unboxed := make([]int64, len(arr))
																	for i, v := range arr {
																		unboxed[i] = v.IntVal
																	}
																	return unboxed
																}(), func() []gopurs_runtime.Value {
																	arr := *(*[]gopurs_runtime.Value)(v9_14.UnsafePtr)
																	unboxed := make([]gopurs_runtime.Value, len(arr))
																	for i, v := range arr {
																		unboxed[i] = v
																	}
																	return unboxed
																}(), v10_15, __local_var_10_27, gopurs_runtime.RecordUpdate3(__local_var_11_28, "a", v11_16, "fa", v12_17, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																	arr := *(*[]gopurs_runtime.Value)(v13_18.UnsafePtr)
																	unboxed := make([]gopurs_runtime.Value, len(arr))
																	for i, v := range arr {
																		unboxed[i] = v
																	}
																	return unboxed
																}())), gopurs_runtime.RecordUpdate1(__local_var_12_29, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(__local_var_12_29, "nested"), "a", v14_19, "fa", v15_20, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
																	arr := *(*[]gopurs_runtime.Value)(v16_21.UnsafePtr)
																	unboxed := make([]gopurs_runtime.Value, len(arr))
																	for i, v := range arr {
																		unboxed[i] = v
																	}
																	return unboxed
																}())))}))}
															})
														})
													})
												})
											})
										})
									})
								})
							}), gopurs_runtime.Apply(f_6, (*Constructor_Main_M6)(m_7.UnsafePtr).V1)), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_13_30, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_13_30, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_3, "pure"), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array((*Constructor_Main_M6)(m_7.UnsafePtr).V3))), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, f_6, (*Constructor_Main_M6)(m_7.UnsafePtr).V4)), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(__local_var_11_28, "a"))), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(__local_var_11_28, "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_13_31, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_13_31, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_3, "pure"), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_11_28, "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}()))), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_12_29, "nested"), "a"))), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_12_29, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_13_32, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_13_32, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_3, "pure"), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array(func() []gopurs_runtime.Value {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_12_29, "nested"), "zArrayA").UnsafePtr)
								unboxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									unboxed[i] = v
								}
								return unboxed
							}())))
							goto end_branch_34
						} else {

						}
					}
					{
						if m_7.Type == 9 && m_7.IntVal == 1168316772 {
							__t34 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_19.V0), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Value{Type: 9, IntVal: 1168316772, UnsafePtr: unsafe.Pointer((&Constructor_Main_M7{1, v1_8}))}
							}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
								// TAST (Let): Apply0_9_33 -> gopurs_runtime.Value
								Apply0_9_33 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{})
								_ = Apply0_9_33
								return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_18.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_19.V0), gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(v4_11 gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.RecordUpdate1(v1_8, "nested", gopurs_runtime.RecordUpdate3(gopurs_runtime.RecordGet(v1_8, "nested"), "a", v2_9, "fa", v3_10, "zArrayA", gopurs_runtime.Array(func() []gopurs_runtime.Value {
												arr := *(*[]gopurs_runtime.Value)(v4_11.UnsafePtr)
												unboxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													unboxed[i] = v
												}
												return unboxed
											}())))
										})
									})
								}), gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "a"))), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, f_6, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "fa"))), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_9_33, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_33, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_3, "pure"), Get_Data_Semigroup_concatArray(), f_6, gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_8, "nested"), "zArrayA").UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}())))
							})), (*Constructor_Main_M7)(m_7.UnsafePtr).V0))
							goto end_branch_34
						} else {

						}
					}
					{
						__t34 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
					}
				end_branch_34:
					return __t34
				})
			})
		})}))}
	}
}

func Call_Main_eqM(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value, dictEq2_2_loop gopurs_runtime.Value, dictEq3_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
	_ = dictEq1_0
	var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
	_ = dictEq_1
	var dictEq2_2 gopurs_runtime.Value = dictEq2_2_loop
	_ = dictEq2_2
	var dictEq3_3 gopurs_runtime.Value = dictEq3_3_loop
	_ = dictEq3_3
	// TAST (Let): eqArray5_4_0 -> *Constructor_Data_Eq_Eq
	eqArray5_4_0 := (&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq3_3, "eq"))})
	_ = eqArray5_4_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 bool
			{
				if x_5.Type == 9 && x_5.IntVal == 3852365315 {
					__t1 = (y_6.Type == 9 && y_6.IntVal == 3852365315)
					goto end_branch_1
				} else {

				}
			}
			{
				if x_5.Type == 9 && x_5.IntVal == 769986722 {
					__t1 = (y_6.Type == 9 && y_6.IntVal == 769986722) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), (*Constructor_Main_M1)(x_5.UnsafePtr).V0, (*Constructor_Main_M1)(y_6.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_4_0.V0), gopurs_runtime.Array((*Constructor_Main_M1)(x_5.UnsafePtr).V1), gopurs_runtime.Array((*Constructor_Main_M1)(y_6.UnsafePtr).V1)).IntVal) != (0)))
					goto end_branch_1
				} else {

				}
			}
			{
				if x_5.Type == 9 && x_5.IntVal == 2727978561 {
					__t1 = (y_6.Type == 9 && y_6.IntVal == 2727978561) && (((*Constructor_Main_M2)(x_5.UnsafePtr).V0) == ((*Constructor_Main_M2)(y_6.UnsafePtr).V0))
					goto end_branch_1
				} else {

				}
			}
			{
				if x_5.Type == 9 && x_5.IntVal == 1830062304 {
					__t1 = (y_6.Type == 9 && y_6.IntVal == 1830062304) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq3_3))}, (*Constructor_Main_M3)(x_5.UnsafePtr).V0, (*Constructor_Main_M3)(y_6.UnsafePtr).V0).IntVal) != (0))
					goto end_branch_1
				} else {

				}
			}
			{
				if x_5.Type == 9 && x_5.IntVal == 3190619783 {
					__t1 = (y_6.Type == 9 && y_6.IntVal == 3190619783) && (((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_5.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_6.UnsafePtr).V0, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_5.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_6.UnsafePtr).V0, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqInt()))}, gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_5.UnsafePtr).V0, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_6.UnsafePtr).V0, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq3_3))}, gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_5.UnsafePtr).V0, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_6.UnsafePtr).V0, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_5.UnsafePtr).V0, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_6.UnsafePtr).V0, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_4_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(x_5.UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M4)(y_6.UnsafePtr).V0, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))
					goto end_branch_1
				} else {

				}
			}
			{
				if x_5.Type == 9 && x_5.IntVal == 108241190 {
					__t1 = (y_6.Type == 9 && y_6.IntVal == 108241190) && (((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_5.UnsafePtr).V0, "nested"), "a"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_6.UnsafePtr).V0, "nested"), "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_5.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_6.UnsafePtr).V0, "nested"), "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqInt()))}, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_5.UnsafePtr).V0, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_6.UnsafePtr).V0, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq3_3))}, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_5.UnsafePtr).V0, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_6.UnsafePtr).V0, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_5.UnsafePtr).V0, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_6.UnsafePtr).V0, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_4_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(x_5.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M5)(y_6.UnsafePtr).V0, "nested"), "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))
					goto end_branch_1
				} else {

				}
			}
			{
				if x_5.Type == 9 && x_5.IntVal == 2066233029 {
					__t1 = (y_6.Type == 9 && y_6.IntVal == 2066233029) && ((((((((((*Constructor_Main_M6)(x_5.UnsafePtr).V0) == ((*Constructor_Main_M6)(y_6.UnsafePtr).V0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), (*Constructor_Main_M6)(x_5.UnsafePtr).V1, (*Constructor_Main_M6)(y_6.UnsafePtr).V1).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M6)(x_5.UnsafePtr).V2
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}(), func() gopurs_runtime.Value {
						arr := (*Constructor_Main_M6)(y_6.UnsafePtr).V2
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_4_0.V0), gopurs_runtime.Array((*Constructor_Main_M6)(x_5.UnsafePtr).V3), gopurs_runtime.Array((*Constructor_Main_M6)(y_6.UnsafePtr).V3)).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq3_3))}, (*Constructor_Main_M6)(x_5.UnsafePtr).V4, (*Constructor_Main_M6)(y_6.UnsafePtr).V4).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqInt()))}, (*Constructor_Main_M6)(x_5.UnsafePtr).V5, (*Constructor_Main_M6)(y_6.UnsafePtr).V5).IntVal) != (0))) && (((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_5.UnsafePtr).V6, "a"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_6.UnsafePtr).V6, "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_5.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_6.UnsafePtr).V6, "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqInt()))}, gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_5.UnsafePtr).V6, "fIgnore"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_6.UnsafePtr).V6, "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq3_3))}, gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_5.UnsafePtr).V6, "fa"), gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_6.UnsafePtr).V6, "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_5.UnsafePtr).V6, "ignore").IntVal) == (gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_6.UnsafePtr).V6, "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_4_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_5.UnsafePtr).V6, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_6.UnsafePtr).V6, "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0)))) && (((((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq3_3, "eq"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_5.UnsafePtr).V7, "nested"), "a"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_6.UnsafePtr).V7, "nested"), "a")).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqArray()).V0), func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_5.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_6.UnsafePtr).V7, "nested"), "arrayIgnore").UnsafePtr)
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
					}()).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqInt()))}, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_5.UnsafePtr).V7, "nested"), "fIgnore"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_6.UnsafePtr).V7, "nested"), "fIgnore")).IntVal) != (0))) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq3_3))}, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_5.UnsafePtr).V7, "nested"), "fa"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_6.UnsafePtr).V7, "nested"), "fa")).IntVal) != (0))) && ((gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_5.UnsafePtr).V7, "nested"), "ignore").IntVal) == (gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_6.UnsafePtr).V7, "nested"), "ignore").IntVal))) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqArray5_4_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(x_5.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet((*Constructor_Main_M6)(y_6.UnsafePtr).V7, "nested"), "zArrayA").UnsafePtr)
						unboxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							unboxed[i] = v
						}
						return unboxed
					}())).IntVal) != (0))))
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = (x_5.Type == 9 && x_5.IntVal == 1168316772) && ((y_6.Type == 9 && y_6.IntVal == 1168316772) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq2_2))}, (*Constructor_Main_M7)(x_5.UnsafePtr).V0, (*Constructor_Main_M7)(y_6.UnsafePtr).V0).IntVal) != (0)))
			}
		end_branch_1:
			return gopurs_runtime.Bool(__t1)
		})
	})}))}
}

func Call_Main_traverseStr(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
	_ = dictTraversable_0
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1))
}

func Call_Main_sequenceStr(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
	_ = dictTraversable_0
	return gopurs_runtime.Apply(gopurs_runtime.Box(dictTraversable_0.V2), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()))})
}

func Call_Main_sequenceStr__1451196854(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
	_ = dictTraversable_0
	return gopurs_runtime.Apply(gopurs_runtime.Box(dictTraversable_0.V2), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()))})
}

func Call_Main_traverseStr__2257532016(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
	var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
	_ = dictTraversable_0
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()).V1))
}
