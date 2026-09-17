package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_eqRec gopurs_runtime.Value
var once_Main_eqRec sync.Once

func Get_Main_eqRec() gopurs_runtime.Value {
	once_Main_eqRec.Do(func() {
		cache_Main_eqRec = gopurs_runtime.Apply(Get_Data_Eq_eqRec(), gopurs_runtime.Value{})
	})
	return cache_Main_eqRec
}

var cache_Main_valueIsSymbol gopurs_runtime.Value
var once_Main_valueIsSymbol sync.Once

func Get_Main_valueIsSymbol() gopurs_runtime.Value {
	once_Main_valueIsSymbol.Do(func() {
		cache_Main_valueIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("value")
		}))
	})
	return cache_Main_valueIsSymbol
}

var cache_Main_keepIsSymbol gopurs_runtime.Value
var once_Main_keepIsSymbol sync.Once

func Get_Main_keepIsSymbol() gopurs_runtime.Value {
	once_Main_keepIsSymbol.Do(func() {
		cache_Main_keepIsSymbol = gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("keep")
		}))
	})
	return cache_Main_keepIsSymbol
}

var cache_Main_eqRowCons gopurs_runtime.Value
var once_Main_eqRowCons sync.Once

func Get_Main_eqRowCons() gopurs_runtime.Value {
	once_Main_eqRowCons.Do(func() {
		cache_Main_eqRowCons = Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_keepIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))})
	})
	return cache_Main_eqRowCons
}

var cache_Main_eqRec1 gopurs_runtime.Value
var once_Main_eqRec1 sync.Once

func Get_Main_eqRec1() gopurs_runtime.Value {
	once_Main_eqRec1.Do(func() {
		cache_Main_eqRec1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2593544496_3790796878(Rebox_Main_3790796878_2593544496(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_keepIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))}))))))}
	})
	return cache_Main_eqRec1
}

var cache_Main_showRecord gopurs_runtime.Value
var once_Main_showRecord sync.Once

func Get_Main_showRecord() gopurs_runtime.Value {
	once_Main_showRecord.Do(func() {
		cache_Main_showRecord = gopurs_runtime.Apply2(Get_Data_Show_showRecord(), gopurs_runtime.Value{}, gopurs_runtime.Value{})
	})
	return cache_Main_showRecord
}

var cache_Main_showRecordFieldsCons gopurs_runtime.Value
var once_Main_showRecordFieldsCons sync.Once

func Get_Main_showRecordFieldsCons() gopurs_runtime.Value {
	once_Main_showRecordFieldsCons.Do(func() {
		cache_Main_showRecordFieldsCons = Call_Data_Show_showRecordFieldsCons(Get_Main_keepIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))})
	})
	return cache_Main_showRecordFieldsCons
}

var cache_Main_showRecord1 gopurs_runtime.Value
var once_Main_showRecord1 sync.Once

func Get_Main_showRecord1() gopurs_runtime.Value {
	once_Main_showRecord1.Do(func() {
		cache_Main_showRecord1 = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2632088208_1386611502(Rebox_Main_1386611502_2632088208(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_keepIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))}))))))}
	})
	return cache_Main_showRecord1
}

var cache_Main_eqList gopurs_runtime.Value
var once_Main_eqList sync.Once

func Get_Main_eqList() gopurs_runtime.Value {
	once_Main_eqList.Do(func() {
		cache_Main_eqList = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1965287036_3790796878(Rebox_Main_3790796878_1965287036(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_List_Types_eqList(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), gopurs_runtime.Value{}, Get_Main_keepIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2737952170_3790796878(Rebox_Main_3790796878_2737952170(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqBoolean()))))})))))))}
	})
	return cache_Main_eqList
}

var cache_Main_showList gopurs_runtime.Value
var once_Main_showList sync.Once

func Get_Main_showList() gopurs_runtime.Value {
	once_Main_showList.Do(func() {
		cache_Main_showList = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_4087401180_1386611502(Rebox_Main_1386611502_4087401180(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_List_Types_showList(Call_Data_Show_showRecord(gopurs_runtime.Value{}, gopurs_runtime.Value{}, Call_Data_Show_showRecordFieldsCons(Get_Main_keepIsSymbol(), Call_Data_Show_showRecordFieldsConsNil(Get_Main_valueIsSymbol(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_2735895690_1386611502(Rebox_Main_1386611502_2735895690(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showBoolean()))))})))))))}
	})
	return cache_Main_showList
}

var cache_Main_keepItems gopurs_runtime.Value
var once_Main_keepItems sync.Once

func Get_Main_keepItems() gopurs_runtime.Value {
	once_Main_keepItems.Do(func() {
		cache_Main_keepItems = func() gopurs_runtime.Value {
			var Call_local_Main_go__770030980_0_0_0 func(*Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}], *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}]) *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}]
			_ = Call_local_Main_go__770030980_0_0_0
			var go__770030980_0_0_0 gopurs_runtime.Value
			_ = go__770030980_0_0_0
			var Call_local_Main_go__go_0_1_1 func(*Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}], *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}]) *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}]
			_ = Call_local_Main_go__go_0_1_1
			var go__go_0_1_1 gopurs_runtime.Value
			_ = go__go_0_1_1
			Call_local_Main_go__770030980_0_0_0 = func(acc_1_loop *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}], v_2_loop *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}]) *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}] {
			go__770030980_0_0_0:
				for {
					if false {
						continue go__770030980_0_0_0
					}
					var acc_1 *Constructor_Data_List_Types_Cons[struct {
						keep  bool
						value int64
					}] = acc_1_loop
					_ = acc_1
					var v_2 *Constructor_Data_List_Types_Cons[struct {
						keep  bool
						value int64
					}] = v_2_loop
					_ = v_2
					var __t7 *Constructor_Data_List_Types_Cons[struct {
						keep  bool
						value int64
					}]
					{
						if v_2 != nil {
							var __t2 *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]
							{
								if (v_2).V0.keep {
									acc_1_loop = (&Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}]{1, (v_2).V0, acc_1})
									v_2_loop = (v_2).V1
									continue go__770030980_0_0_0
									__t2 = func() *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}] {
										panic("unreachable")
									}()
									goto end_branch_2
								} else {

								}
							}
							{
								acc_1_loop = acc_1
								v_2_loop = (v_2).V1
								continue go__770030980_0_0_0
								__t2 = func() *Constructor_Data_List_Types_Cons[struct {
									keep  bool
									value int64
								}] {
									panic("unreachable")
								}()
							}
						end_branch_2:
							__t7 = __t2
							goto end_branch_7
						} else {

						}
					}
					{
						if v_2 == nil {
							var Call_local_Main_go__770030980_3_3_2 func(*Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}], *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]) *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]
							_ = Call_local_Main_go__770030980_3_3_2
							var go__770030980_3_3_2 gopurs_runtime.Value
							_ = go__770030980_3_3_2
							var Call_local_Main_go__go_3_4_3 func(*Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}], *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]) *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]
							_ = Call_local_Main_go__go_3_4_3
							var go__go_3_4_3 gopurs_runtime.Value
							_ = go__go_3_4_3
							Call_local_Main_go__770030980_3_3_2 = func(v_4_loop *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}], v1_5_loop *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]) *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}] {
							go__770030980_3_3_2:
								for {
									if false {
										continue go__770030980_3_3_2
									}
									var v_4 *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}] = v_4_loop
									_ = v_4
									var v1_5 *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}] = v1_5_loop
									_ = v1_5
									var __t5 *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}]
									{
										if v1_5 == nil {
											__t5 = v_4
											goto end_branch_5
										} else {

										}
									}
									{
										if v1_5 != nil {
											v_4_loop = (&Constructor_Data_List_Types_Cons[struct {
												keep  bool
												value int64
											}]{1, (v1_5).V0, v_4})
											v1_5_loop = (v1_5).V1
											continue go__770030980_3_3_2
											__t5 = func() *Constructor_Data_List_Types_Cons[struct {
												keep  bool
												value int64
											}] {
												panic("unreachable")
											}()
											goto end_branch_5
										} else {

										}
									}
									{
										__t5 = func() *Constructor_Data_List_Types_Cons[struct {
											keep  bool
											value int64
										}] {
											panic("Failed pattern match")
										}()
									}
								end_branch_5:
									return __t5
								}
							}
							go__770030980_3_3_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(Call_local_Main_go__770030980_3_3_2(Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val)), Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))))}
								})
							})
							Call_local_Main_go__go_3_4_3 = func(v_4_loop *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}], v1_5_loop *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]) *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}] {
							go__go_3_4_3:
								for {
									if false {
										continue go__go_3_4_3
									}
									var v_4 *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}] = v_4_loop
									_ = v_4
									var v1_5 *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}] = v1_5_loop
									_ = v1_5
									var __t6 *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}]
									{
										if v1_5 == nil {
											__t6 = v_4
											goto end_branch_6
										} else {

										}
									}
									{
										if v1_5 != nil {
											__t6 = Call_local_Main_go__770030980_3_3_2((&Constructor_Data_List_Types_Cons[struct {
												keep  bool
												value int64
											}]{1, (v1_5).V0, v_4}), (v1_5).V1)
											goto end_branch_6
										} else {

										}
									}
									{
										__t6 = func() *Constructor_Data_List_Types_Cons[struct {
											keep  bool
											value int64
										}] {
											panic("Failed pattern match")
										}()
									}
								end_branch_6:
									return __t6
								}
							}
							go__go_3_4_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(Call_local_Main_go__go_3_4_3(Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val)), Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))))}
								})
							})
							__t7 = Call_local_Main_go__770030980_3_3_2((*Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}])(nil), acc_1)
							goto end_branch_7
						} else {

						}
					}
					{
						__t7 = func() *Constructor_Data_List_Types_Cons[struct {
							keep  bool
							value int64
						}] {
							panic("Failed pattern match")
						}()
					}
				end_branch_7:
					return __t7
				}
			}
			go__770030980_0_0_0 = gopurs_runtime.Func(func(acc_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(Call_local_Main_go__770030980_0_0_0(Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](acc_1_loop_val)), Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)))))}
				})
			})
			Call_local_Main_go__go_0_1_1 = func(acc_1_loop *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}], v_2_loop *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}]) *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}] {
			go__go_0_1_1:
				for {
					if false {
						continue go__go_0_1_1
					}
					var acc_1 *Constructor_Data_List_Types_Cons[struct {
						keep  bool
						value int64
					}] = acc_1_loop
					_ = acc_1
					var v_2 *Constructor_Data_List_Types_Cons[struct {
						keep  bool
						value int64
					}] = v_2_loop
					_ = v_2
					var __t13 *Constructor_Data_List_Types_Cons[struct {
						keep  bool
						value int64
					}]
					{
						if v_2 != nil {
							var __t8 *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]
							{
								if (v_2).V0.keep {
									__t8 = Call_local_Main_go__770030980_0_0_0((&Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}]{1, (v_2).V0, acc_1}), (v_2).V1)
									goto end_branch_8
								} else {

								}
							}
							{
								__t8 = Call_local_Main_go__770030980_0_0_0(acc_1, (v_2).V1)
							}
						end_branch_8:
							__t13 = __t8
							goto end_branch_13
						} else {

						}
					}
					{
						if v_2 == nil {
							var Call_local_Main_go__770030980_3_9_4 func(*Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}], *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]) *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]
							_ = Call_local_Main_go__770030980_3_9_4
							var go__770030980_3_9_4 gopurs_runtime.Value
							_ = go__770030980_3_9_4
							var Call_local_Main_go__go_3_10_5 func(*Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}], *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]) *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]
							_ = Call_local_Main_go__go_3_10_5
							var go__go_3_10_5 gopurs_runtime.Value
							_ = go__go_3_10_5
							Call_local_Main_go__770030980_3_9_4 = func(v_4_loop *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}], v1_5_loop *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]) *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}] {
							go__770030980_3_9_4:
								for {
									if false {
										continue go__770030980_3_9_4
									}
									var v_4 *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}] = v_4_loop
									_ = v_4
									var v1_5 *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}] = v1_5_loop
									_ = v1_5
									var __t11 *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}]
									{
										if v1_5 == nil {
											__t11 = v_4
											goto end_branch_11
										} else {

										}
									}
									{
										if v1_5 != nil {
											v_4_loop = (&Constructor_Data_List_Types_Cons[struct {
												keep  bool
												value int64
											}]{1, (v1_5).V0, v_4})
											v1_5_loop = (v1_5).V1
											continue go__770030980_3_9_4
											__t11 = func() *Constructor_Data_List_Types_Cons[struct {
												keep  bool
												value int64
											}] {
												panic("unreachable")
											}()
											goto end_branch_11
										} else {

										}
									}
									{
										__t11 = func() *Constructor_Data_List_Types_Cons[struct {
											keep  bool
											value int64
										}] {
											panic("Failed pattern match")
										}()
									}
								end_branch_11:
									return __t11
								}
							}
							go__770030980_3_9_4 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(Call_local_Main_go__770030980_3_9_4(Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val)), Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))))}
								})
							})
							Call_local_Main_go__go_3_10_5 = func(v_4_loop *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}], v1_5_loop *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}]) *Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}] {
							go__go_3_10_5:
								for {
									if false {
										continue go__go_3_10_5
									}
									var v_4 *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}] = v_4_loop
									_ = v_4
									var v1_5 *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}] = v1_5_loop
									_ = v1_5
									var __t12 *Constructor_Data_List_Types_Cons[struct {
										keep  bool
										value int64
									}]
									{
										if v1_5 == nil {
											__t12 = v_4
											goto end_branch_12
										} else {

										}
									}
									{
										if v1_5 != nil {
											__t12 = Call_local_Main_go__770030980_3_9_4((&Constructor_Data_List_Types_Cons[struct {
												keep  bool
												value int64
											}]{1, (v1_5).V0, v_4}), (v1_5).V1)
											goto end_branch_12
										} else {

										}
									}
									{
										__t12 = func() *Constructor_Data_List_Types_Cons[struct {
											keep  bool
											value int64
										}] {
											panic("Failed pattern match")
										}()
									}
								end_branch_12:
									return __t12
								}
							}
							go__go_3_10_5 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(Call_local_Main_go__go_3_10_5(Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val)), Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))))}
								})
							})
							__t13 = Call_local_Main_go__770030980_3_9_4((*Constructor_Data_List_Types_Cons[struct {
								keep  bool
								value int64
							}])(nil), acc_1)
							goto end_branch_13
						} else {

						}
					}
					{
						__t13 = func() *Constructor_Data_List_Types_Cons[struct {
							keep  bool
							value int64
						}] {
							panic("Failed pattern match")
						}()
					}
				end_branch_13:
					return __t13
				}
			}
			go__go_0_1_1 = gopurs_runtime.Func(func(acc_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(Call_local_Main_go__go_0_1_1(Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](acc_1_loop_val)), Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)))))}
				})
			})
			return gopurs_runtime.Apply(go__770030980_0_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993((*Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}])(nil)))})
		}()
	})
	return cache_Main_keepItems
}

var cache_Main_checkRecordReturns gopurs_runtime.Value
var once_Main_checkRecordReturns sync.Once

func Get_Main_checkRecordReturns() gopurs_runtime.Value {
	once_Main_checkRecordReturns.Do(func() {
		cache_Main_checkRecordReturns = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(40)))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_1), gopurs_runtime.Value{})
			_ = __local_var_2_2
			return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___3363913000("", struct {
				actual struct {
					keep  bool
					value int64
				}
				expected struct {
					keep  bool
					value int64
				}
			}{struct {
				keep  bool
				value int64
			}{true, (__local_var_2_2.IntVal) + (int64(2))}, struct {
				keep  bool
				value int64
			}{true, int64(42)}}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Func(func(offset_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return func() gopurs_runtime.Value {
							orig := struct {
								keep  bool
								value int64
							}{(offset_4.IntVal) >= (int64(0)), (__local_var_2_2.IntVal) + (offset_4.IntVal)}
							_ = orig
							return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool(orig.keep), gopurs_runtime.Int(orig.value))
						}()
					})), gopurs_runtime.Value{})
					_ = __local_var_4_3
					__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_3), gopurs_runtime.Value{})
					_ = __local_var_5_4
					return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___3363913000("", struct {
						actual struct {
							keep  bool
							value int64
						}
						expected struct {
							keep  bool
							value int64
						}
					}{func() struct {
						keep  bool
						value int64
					} {
						orig := gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Int(int64(5)))
						_ = orig
						clone := struct {
							keep  bool
							value int64
						}{}
						clone.keep = (gopurs_runtime.RecordGet(orig, "keep").IntVal) != (0)
						clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
						return clone
					}(), struct {
						keep  bool
						value int64
					}{true, int64(45)}}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___3363913000("", struct {
							actual struct {
								keep  bool
								value int64
							}
							expected struct {
								keep  bool
								value int64
							}
						}{func() struct {
							keep  bool
							value int64
						} {
							orig := gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Int(int64(-1)))
							_ = orig
							clone := struct {
								keep  bool
								value int64
							}{}
							clone.keep = (gopurs_runtime.RecordGet(orig, "keep").IntVal) != (0)
							clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
							return clone
						}(), struct {
							keep  bool
							value int64
						}{false, int64(39)}}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								__local_var_8_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Func(func(second_8 gopurs_runtime.Value) gopurs_runtime.Value {
									// TAST (Let): __local_var_9_6 shape=Other bindingType=Int
									__local_var_9_6 := (int64(2)) + (second_8.IntVal)
									_ = __local_var_9_6
									return func() gopurs_runtime.Value {
										orig := struct {
											keep  bool
											value int64
										}{(__local_var_9_6) >= (int64(0)), (__local_var_2_2.IntVal) + (__local_var_9_6)}
										_ = orig
										return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool(orig.keep), gopurs_runtime.Int(orig.value))
									}()
								})), gopurs_runtime.Value{})
								_ = __local_var_8_5
								__local_var_9_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_8_5), gopurs_runtime.Value{})
								_ = __local_var_9_7
								return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___3363913000("", struct {
									actual struct {
										keep  bool
										value int64
									}
									expected struct {
										keep  bool
										value int64
									}
								}{func() struct {
									keep  bool
									value int64
								} {
									orig := gopurs_runtime.Apply(__local_var_9_7, gopurs_runtime.Int(int64(3)))
									_ = orig
									clone := struct {
										keep  bool
										value int64
									}{}
									clone.keep = (gopurs_runtime.RecordGet(orig, "keep").IntVal) != (0)
									clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
									return clone
								}(), struct {
									keep  bool
									value int64
								}{true, int64(45)}}), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("local record returns"))
								})), gopurs_runtime.Value{})
							})
						}))
					})), gopurs_runtime.Value{})
				})
			})), gopurs_runtime.Value{})
		})
	})
	return cache_Main_checkRecordReturns
}

var cache_Main_checkItems gopurs_runtime.Value
var once_Main_checkItems sync.Once

func Get_Main_checkItems() gopurs_runtime.Value {
	once_Main_checkItems.Do(func() {
		cache_Main_checkItems = gopurs_runtime.Func3(func(label_0_box gopurs_runtime.Value, expected_1_box gopurs_runtime.Value, original_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_checkItems(label_0_box.StrVal(), Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](expected_1_box)), Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](original_2_box)))
		})
	})
	return cache_Main_checkItems
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkItems("guarded list preserves order", (&Constructor_Data_List_Types_Cons[struct {
			keep  bool
			value int64
		}]{1, struct {
			keep  bool
			value int64
		}{true, int64(7)}, (&Constructor_Data_List_Types_Cons[struct {
			keep  bool
			value int64
		}]{1, struct {
			keep  bool
			value int64
		}{true, int64(3)}, (*Constructor_Data_List_Types_Cons[struct {
			keep  bool
			value int64
		}])(nil)})}), (&Constructor_Data_List_Types_Cons[struct {
			keep  bool
			value int64
		}]{1, struct {
			keep  bool
			value int64
		}{false, int64(0)}, (&Constructor_Data_List_Types_Cons[struct {
			keep  bool
			value int64
		}]{1, struct {
			keep  bool
			value int64
		}{true, int64(7)}, (&Constructor_Data_List_Types_Cons[struct {
			keep  bool
			value int64
		}]{1, struct {
			keep  bool
			value int64
		}{false, int64(2)}, (&Constructor_Data_List_Types_Cons[struct {
			keep  bool
			value int64
		}]{1, struct {
			keep  bool
			value int64
		}{true, int64(3)}, (*Constructor_Data_List_Types_Cons[struct {
			keep  bool
			value int64
		}])(nil)})})})})), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkItems("empty list", (*Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}])(nil), (*Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}])(nil)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Main_checkItems("all rejected", (*Constructor_Data_List_Types_Cons[struct {
					keep  bool
					value int64
				}])(nil), (&Constructor_Data_List_Types_Cons[struct {
					keep  bool
					value int64
				}]{1, struct {
					keep  bool
					value int64
				}{false, int64(4)}, (&Constructor_Data_List_Types_Cons[struct {
					keep  bool
					value int64
				}]{1, struct {
					keep  bool
					value int64
				}{false, int64(8)}, (*Constructor_Data_List_Types_Cons[struct {
					keep  bool
					value int64
				}])(nil)})})), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_checkRecordReturns(), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
					}))
				}))
			}))
		}))
	})
	return cache_Main_main
}

func Call_Main_checkItems(label_0_loop string, expected_1_loop *Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
}], original_2_loop *Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
}]) gopurs_runtime.Value {
	var label_0 string = label_0_loop
	_ = label_0
	var expected_1 *Constructor_Data_List_Types_Cons[struct {
		keep  bool
		value int64
	}] = expected_1_loop
	_ = expected_1
	var original_2 *Constructor_Data_List_Types_Cons[struct {
		keep  bool
		value int64
	}] = original_2_loop
	_ = original_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(original_2))})
		_ = __local_var_3_0
		__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = __local_var_4_1
		__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1), gopurs_runtime.Value{})
		_ = __local_var_5_2
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___282923697("", struct {
			actual *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}]
			expected *Constructor_Data_List_Types_Cons[struct {
				keep  bool
				value int64
			}]
		}{Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Main_keepItems(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_917490743_849153993(Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_5_2))))}))), expected_1}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_7_3 shape=App(Var) bindingType=Any
				__local_var_7_3 := gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_4_1)
				_ = __local_var_7_3
				__local_var_8_4 := gopurs_runtime.Apply(__local_var_7_3, gopurs_runtime.Value{})
				_ = __local_var_8_4
				return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Call_Test_Assert_assertEqual_prime___282923697("", struct {
					actual *Constructor_Data_List_Types_Cons[struct {
						keep  bool
						value int64
					}]
					expected *Constructor_Data_List_Types_Cons[struct {
						keep  bool
						value int64
					}]
				}{Rebox_Main_849153993_917490743(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_8_4)), original_2}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(label_0))
				})), gopurs_runtime.Value{})
			})
		})), gopurs_runtime.Value{})
	})
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_2632088208(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[struct {
	keep  bool
	value int64
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[struct {
		keep  bool
		value int64
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_2735895690(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[bool]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1386611502_4087401180(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
}]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[struct {
		keep  bool
		value int64
	}]]{}
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

func Rebox_Main_1965287036_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
}]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2593544496_3790796878(in *Constructor_Data_Eq_Eq[struct {
	keep  bool
	value int64
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2632088208_1386611502(in *Constructor_Data_Show_Show[struct {
	keep  bool
	value int64
}]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2735895690_1386611502(in *Constructor_Data_Show_Show[bool]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_1965287036(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
}]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[struct {
		keep  bool
		value int64
	}]]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_2593544496(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[struct {
	keep  bool
	value int64
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[struct {
		keep  bool
		value int64
	}]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_3790796878_2737952170(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[bool] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[bool]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_4087401180_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
}]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_849153993_917490743(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[struct {
		keep  bool
		value int64
	}]{}
	out.V0 = func() struct {
		keep  bool
		value int64
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			keep  bool
			value int64
		}{}
		clone.keep = (gopurs_runtime.RecordGet(orig, "keep").IntVal) != (0)
		clone.value = gopurs_runtime.RecordGet(orig, "value").IntVal
		return clone
	}()
	out.V1 = Rebox_Main_849153993_917490743(in.V1)
	return out
}

func Rebox_Main_917490743_849153993(in *Constructor_Data_List_Types_Cons[struct {
	keep  bool
	value int64
}]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict2("keep", "value", gopurs_runtime.Bool(orig.keep), gopurs_runtime.Int(orig.value))
	}()
	out.V1 = Rebox_Main_917490743_849153993(in.V1)
	return out
}
