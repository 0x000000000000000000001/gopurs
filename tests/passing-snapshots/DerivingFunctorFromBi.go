package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_bifoldl gopurs_runtime.Value
var once_Main_bifoldl sync.Once

func Get_Main_bifoldl() gopurs_runtime.Value {
	once_Main_bifoldl.Do(func() {
		cache_Main_bifoldl = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bifoldl(f_0_box, g_1_box, z_2_box, Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_3_box)))
		})
	})
	return cache_Main_bifoldl
}

var cache_Main_bifoldl1 gopurs_runtime.Value
var once_Main_bifoldl1 sync.Once

func Get_Main_bifoldl1() gopurs_runtime.Value {
	once_Main_bifoldl1.Do(func() {
		cache_Main_bifoldl1 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bifoldl1(f_0_box, g_1_box, z_2_box, Rebox_Main_138441832_2137532138(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_3_box)))
		})
	})
	return cache_Main_bifoldl1
}

var cache_Main_bifoldr gopurs_runtime.Value
var once_Main_bifoldr sync.Once

func Get_Main_bifoldr() gopurs_runtime.Value {
	once_Main_bifoldr.Do(func() {
		cache_Main_bifoldr = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bifoldr(f_0_box, g_1_box, z_2_box, Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_3_box)))
		})
	})
	return cache_Main_bifoldr
}

var cache_Main_bifoldr1 gopurs_runtime.Value
var once_Main_bifoldr1 sync.Once

func Get_Main_bifoldr1() gopurs_runtime.Value {
	once_Main_bifoldr1.Do(func() {
		cache_Main_bifoldr1 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bifoldr1(f_0_box, g_1_box, z_2_box, Rebox_Main_138441832_2137532138(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_3_box)))
		})
	})
	return cache_Main_bifoldr1
}

var cache_Main_bifoldMap gopurs_runtime.Value
var once_Main_bifoldMap sync.Once

func Get_Main_bifoldMap() gopurs_runtime.Value {
	once_Main_bifoldMap.Do(func() {
		cache_Main_bifoldMap = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_bifoldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
		})
	})
	return cache_Main_bifoldMap
}

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
		cache_Main_eqRowCons = gopurs_runtime.Apply3(Get_Data_Eq_eqRowCons(), Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("a")
		})))
	})
	return cache_Main_eqRowCons
}

var cache_Main_show gopurs_runtime.Value
var once_Main_show sync.Once

func Get_Main_show() gopurs_runtime.Value {
	once_Main_show.Do(func() {
		cache_Main_show = Get_Data_Show_showIntImpl()
	})
	return cache_Main_show
}

var cache_Main_eqArray gopurs_runtime.Value
var once_Main_eqArray sync.Once

func Get_Main_eqArray() gopurs_runtime.Value {
	once_Main_eqArray.Do(func() {
		cache_Main_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_378698611_3790796878(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Eq_eqArray(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqArray
}

var cache_Main_Test1 gopurs_runtime.Value
var once_Main_Test1 sync.Once

func Get_Main_Test1() gopurs_runtime.Value {
	once_Main_Test1.Do(func() {
		cache_Main_Test1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1[gopurs_runtime.Value]{1, Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](value0))}))}
		})
	})
	return cache_Main_Test1
}

var cache_Main_Test1__3723567844 gopurs_runtime.Value
var once_Main_Test1__3723567844 sync.Once

func Get_Main_Test1__3723567844() gopurs_runtime.Value {
	once_Main_Test1__3723567844.Do(func() {
		cache_Main_Test1__3723567844 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test1__3723567844(Rebox_Main_138441832_3363075976(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_0_0_box)))
		})
	})
	return cache_Main_Test1__3723567844
}

var cache_Main_Test1__1738194564 gopurs_runtime.Value
var once_Main_Test1__1738194564 sync.Once

func Get_Main_Test1__1738194564() gopurs_runtime.Value {
	once_Main_Test1__1738194564.Do(func() {
		cache_Main_Test1__1738194564 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test1__1738194564(Rebox_Main_138441832_1830029836(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_0_0_box)))
		})
	})
	return cache_Main_Test1__1738194564
}

var cache_Main_Test2 gopurs_runtime.Value
var once_Main_Test2 sync.Once

func Get_Main_Test2() gopurs_runtime.Value {
	once_Main_Test2.Do(func() {
		cache_Main_Test2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value]{1, Rebox_Main_138441832_3791647502(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](value0))}))}
		})
	})
	return cache_Main_Test2
}

var cache_Main_Test2__2545363874 gopurs_runtime.Value
var once_Main_Test2__2545363874 sync.Once

func Get_Main_Test2__2545363874() gopurs_runtime.Value {
	once_Main_Test2__2545363874.Do(func() {
		cache_Main_Test2__2545363874 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test2__2545363874(Rebox_Main_138441832_1259605934(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_0_0_box)))
		})
	})
	return cache_Main_Test2__2545363874
}

var cache_Main_Test2__877929284 gopurs_runtime.Value
var once_Main_Test2__877929284 sync.Once

func Get_Main_Test2__877929284() gopurs_runtime.Value {
	once_Main_Test2__877929284.Do(func() {
		cache_Main_Test2__877929284 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test2__877929284(Rebox_Main_138441832_839290894(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_0_0_box)))
		})
	})
	return cache_Main_Test2__877929284
}

var cache_Main_Test3 gopurs_runtime.Value
var once_Main_Test3 sync.Once

func Get_Main_Test3() gopurs_runtime.Value {
	once_Main_Test3.Do(func() {
		cache_Main_Test3 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3[gopurs_runtime.Value]{1, func() struct {
				x *Constructor_Data_Tuple_Tuple[struct {
					a gopurs_runtime.Value
				}, int64]
				y *Constructor_Data_Tuple_Tuple[struct {
					a []gopurs_runtime.Value
				}, struct {
					a gopurs_runtime.Value
				}]
			} {
				orig := value0
				_ = orig
				clone := struct {
					x *Constructor_Data_Tuple_Tuple[struct {
						a gopurs_runtime.Value
					}, int64]
					y *Constructor_Data_Tuple_Tuple[struct {
						a []gopurs_runtime.Value
					}, struct {
						a gopurs_runtime.Value
					}]
				}{}
				clone.x = Rebox_Main_138441832_2137532138(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "x")))
				clone.y = Rebox_Main_138441832_372546286(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "y")))
				return clone
			}()}))}
		})
	})
	return cache_Main_Test3
}

var cache_Main_Test3__3463130812 gopurs_runtime.Value
var once_Main_Test3__3463130812 sync.Once

func Get_Main_Test3__3463130812() gopurs_runtime.Value {
	once_Main_Test3__3463130812.Do(func() {
		cache_Main_Test3__3463130812 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test3__3463130812(func() struct {
				x *Constructor_Data_Tuple_Tuple[struct {
					a int64
				}, int64]
				y *Constructor_Data_Tuple_Tuple[struct {
					a []int64
				}, struct {
					a int64
				}]
			} {
				orig := __eta_norm_0_0_box
				_ = orig
				clone := struct {
					x *Constructor_Data_Tuple_Tuple[struct {
						a int64
					}, int64]
					y *Constructor_Data_Tuple_Tuple[struct {
						a []int64
					}, struct {
						a int64
					}]
				}{}
				clone.x = Rebox_Main_138441832_401393041(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "x")))
				clone.y = Rebox_Main_138441832_4123273614(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "y")))
				return clone
			}())
		})
	})
	return cache_Main_Test3__3463130812
}

var cache_Main_Test3__2986798300 gopurs_runtime.Value
var once_Main_Test3__2986798300 sync.Once

func Get_Main_Test3__2986798300() gopurs_runtime.Value {
	once_Main_Test3__2986798300.Do(func() {
		cache_Main_Test3__2986798300 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test3__2986798300(func() struct {
				x *Constructor_Data_Tuple_Tuple[struct {
					a string
				}, int64]
				y *Constructor_Data_Tuple_Tuple[struct {
					a []string
				}, struct {
					a string
				}]
			} {
				orig := __eta_norm_0_0_box
				_ = orig
				clone := struct {
					x *Constructor_Data_Tuple_Tuple[struct {
						a string
					}, int64]
					y *Constructor_Data_Tuple_Tuple[struct {
						a []string
					}, struct {
						a string
					}]
				}{}
				clone.x = Rebox_Main_138441832_1612762549(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "x")))
				clone.y = Rebox_Main_138441832_1971292750(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "y")))
				return clone
			}())
		})
	})
	return cache_Main_Test3__2986798300
}

var cache_Main_functorTest gopurs_runtime.Value
var once_Main_functorTest sync.Once

func Get_Main_functorTest() gopurs_runtime.Value {
	once_Main_functorTest.Do(func() {
		cache_Main_functorTest = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if m_1.Type == 9 && m_1.IntVal == 3720114489 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1[gopurs_runtime.Value]{1, Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Apply(f_0, ((*Constructor_Main_Test1[gopurs_runtime.Value])(m_1.UnsafePtr).V0).V0), gopurs_runtime.Int(((*Constructor_Main_Test1[gopurs_runtime.Value])(m_1.UnsafePtr).V0).V1)}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()))}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_1.Type == 9 && m_1.IntVal == 2375191994 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value]{1, Rebox_Main_138441832_3791647502(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
						_v := struct {
							V0 gopurs_runtime.Value
							V1 gopurs_runtime.Value
						}{gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
							arr_val_arrayMap5 := gopurs_runtime.Array(((*Constructor_Main_Test2[gopurs_runtime.Value])(m_1.UnsafePtr).V0).V0)
							_ = arr_val_arrayMap5
							arr_go_arrayMap5 := (*[]gopurs_runtime.Value)(arr_val_arrayMap5.UnsafePtr)
							_ = arr_go_arrayMap5
							res_go_arrayMap5 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap5))
							_ = res_go_arrayMap5
							for i_arrayMap5, v_arrayMap5 := range *arr_go_arrayMap5 {
								res_go_arrayMap5[i_arrayMap5] = gopurs_runtime.Apply(f_0, v_arrayMap5)
							}
							return gopurs_runtime.Array(res_go_arrayMap5)
						}()).UnsafePtr))), gopurs_runtime.Apply(f_0, ((*Constructor_Main_Test2[gopurs_runtime.Value])(m_1.UnsafePtr).V0).V1)}
						return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
					}()))}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_1.Type == 9 && m_1.IntVal == 227416251 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3[gopurs_runtime.Value]{1, func(record struct {
						x *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]
						y *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]
					}) struct {
						x *Constructor_Data_Tuple_Tuple[struct {
							a gopurs_runtime.Value
						}, int64]
						y *Constructor_Data_Tuple_Tuple[struct {
							a []gopurs_runtime.Value
						}, struct {
							a gopurs_runtime.Value
						}]
					} {
						return struct {
							x *Constructor_Data_Tuple_Tuple[struct {
								a gopurs_runtime.Value
							}, int64]
							y *Constructor_Data_Tuple_Tuple[struct {
								a []gopurs_runtime.Value
							}, struct {
								a gopurs_runtime.Value
							}]
						}{Rebox_Main_138441832_2137532138(record.x), Rebox_Main_138441832_372546286(record.y)}
					}(func() struct {
						x *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]
						y *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]
					} {
						originalRecord := (*Constructor_Main_Test3[gopurs_runtime.Value])(m_1.UnsafePtr).V0
						_ = originalRecord
						_ = originalRecord
						var clone struct {
							x *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]
							y *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]
						}
						clone.x = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
							}{func() gopurs_runtime.Value {
								orig := func() struct {
									a gopurs_runtime.Value
								} {
									clone := ((*Constructor_Main_Test3[gopurs_runtime.Value])(m_1.UnsafePtr).V0.x).V0
									clone.a = gopurs_runtime.Apply(f_0, ((*Constructor_Main_Test3[gopurs_runtime.Value])(m_1.UnsafePtr).V0.x).V0.a)
									return clone
								}()
								_ = orig
								return gopurs_runtime.RecordDict1("a", orig.a)
							}(), gopurs_runtime.Int(((*Constructor_Main_Test3[gopurs_runtime.Value])(m_1.UnsafePtr).V0.x).V1)}
							return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
						}())
						clone.y = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
							_v := struct {
								V0 gopurs_runtime.Value
								V1 gopurs_runtime.Value
							}{func() gopurs_runtime.Value {
								orig := func() struct {
									a []gopurs_runtime.Value
								} {
									clone := ((*Constructor_Main_Test3[gopurs_runtime.Value])(m_1.UnsafePtr).V0.y).V0
									clone.a = (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
										arr_val_arrayMap7 := gopurs_runtime.Array(((*Constructor_Main_Test3[gopurs_runtime.Value])(m_1.UnsafePtr).V0.y).V0.a)
										_ = arr_val_arrayMap7
										arr_go_arrayMap7 := (*[]gopurs_runtime.Value)(arr_val_arrayMap7.UnsafePtr)
										_ = arr_go_arrayMap7
										res_go_arrayMap7 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap7))
										_ = res_go_arrayMap7
										for i_arrayMap7, v_arrayMap7 := range *arr_go_arrayMap7 {
											res_go_arrayMap7[i_arrayMap7] = gopurs_runtime.Apply(f_0, v_arrayMap7)
										}
										return gopurs_runtime.Array(res_go_arrayMap7)
									}()).UnsafePtr))
									return clone
								}()
								_ = orig
								return gopurs_runtime.RecordDict1("a", gopurs_runtime.Array(orig.a))
							}(), func() gopurs_runtime.Value {
								orig := func() struct {
									a gopurs_runtime.Value
								} {
									clone := ((*Constructor_Main_Test3[gopurs_runtime.Value])(m_1.UnsafePtr).V0.y).V1
									clone.a = gopurs_runtime.Apply(f_0, ((*Constructor_Main_Test3[gopurs_runtime.Value])(m_1.UnsafePtr).V0.y).V1.a)
									return clone
								}()
								_ = orig
								return gopurs_runtime.RecordDict1("a", orig.a)
							}()}
							return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
						}())
						return clone
					}())}))}
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_0:
			return __t0
		})}))}
	})
	return cache_Main_functorTest
}

var cache_Main_foldableTest gopurs_runtime.Value
var once_Main_foldableTest sync.Once

func Get_Main_foldableTest() gopurs_runtime.Value {
	once_Main_foldableTest.Do(func() {
		cache_Main_foldableTest = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope54)])
			Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
			_ = Semigroup0_1_0
			// TAST (Let): mempty_2_1 shape=App(Var) bindingType=(Func [Int] (TypeVar m$scope37))
			mempty_2_1 := Call_Data_Monoid_mempty(Call_Data_Monoid_monoidFn(dictMonoid_0))
			_ = mempty_2_1
			// TAST (Let): Semigroup0_3_2 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope37)])
			Semigroup0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
			_ = Semigroup0_3_2
			// TAST (Let): Semigroup0_4_3 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope54)])
			Semigroup0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
			_ = Semigroup0_4_3
			return gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, m_6 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t4 gopurs_runtime.Value
				{
					if m_6.Type == 9 && m_6.IntVal == 3720114489 {
						__t4 = gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_5, ((*Constructor_Main_Test1[gopurs_runtime.Value])(m_6.UnsafePtr).V0).V0), gopurs_runtime.Apply(mempty_2_1, gopurs_runtime.Int(((*Constructor_Main_Test1[gopurs_runtime.Value])(m_6.UnsafePtr).V0).V1)))
						goto end_branch_4
					} else {

					}
				}
				{
					if m_6.Type == 9 && m_6.IntVal == 2375191994 {
						__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0)), f_5, gopurs_runtime.Array(((*Constructor_Main_Test2[gopurs_runtime.Value])(m_6.UnsafePtr).V0).V0)), gopurs_runtime.Apply(f_5, ((*Constructor_Main_Test2[gopurs_runtime.Value])(m_6.UnsafePtr).V0).V1))
						goto end_branch_4
					} else {

					}
				}
				{
					if m_6.Type == 9 && m_6.IntVal == 227416251 {
						__t4 = gopurs_runtime.Apply2(Semigroup0_3_2.V0, gopurs_runtime.Apply2(Semigroup0_4_3.V0, gopurs_runtime.Apply(f_5, ((*Constructor_Main_Test3[gopurs_runtime.Value])(m_6.UnsafePtr).V0.x).V0.a), gopurs_runtime.Apply(mempty_2_1, gopurs_runtime.Int(((*Constructor_Main_Test3[gopurs_runtime.Value])(m_6.UnsafePtr).V0.x).V1))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0)), f_5, gopurs_runtime.Array(((*Constructor_Main_Test3[gopurs_runtime.Value])(m_6.UnsafePtr).V0.y).V0.a)), gopurs_runtime.Apply(f_5, ((*Constructor_Main_Test3[gopurs_runtime.Value])(m_6.UnsafePtr).V0.y).V1.a)))
						goto end_branch_4
					} else {

					}
				}
				{
					__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_4:
				return __t4
			})
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t5 gopurs_runtime.Value
			{
				if m_2.Type == 9 && m_2.IntVal == 3720114489 {
					__t5 = Call_Data_Function_go__const(gopurs_runtime.Apply2(f_0, z_1, ((*Constructor_Main_Test1[gopurs_runtime.Value])(m_2.UnsafePtr).V0).V0), gopurs_runtime.Int(((*Constructor_Main_Test1[gopurs_runtime.Value])(m_2.UnsafePtr).V0).V1))
					goto end_branch_5
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 2375191994 {
					__t5 = gopurs_runtime.Apply2(f_0, func() gopurs_runtime.Value {
						arr_val_foldlArray4 := gopurs_runtime.Array(((*Constructor_Main_Test2[gopurs_runtime.Value])(m_2.UnsafePtr).V0).V0)
						_ = arr_val_foldlArray4
						res_go_foldlArray4 := z_1
						_ = res_go_foldlArray4
						arr_go_foldlArray4 := (*[]gopurs_runtime.Value)(arr_val_foldlArray4.UnsafePtr)
						_ = arr_go_foldlArray4
						for _, v_foldlArray4 := range *arr_go_foldlArray4 {
							res_go_foldlArray4 = gopurs_runtime.Apply2(f_0, res_go_foldlArray4, v_foldlArray4)
						}
						return res_go_foldlArray4
					}(), ((*Constructor_Main_Test2[gopurs_runtime.Value])(m_2.UnsafePtr).V0).V1)
					goto end_branch_5
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 227416251 {
					__t5 = gopurs_runtime.Apply2(f_0, func() gopurs_runtime.Value {
						arr_val_foldlArray4 := gopurs_runtime.Array(((*Constructor_Main_Test3[gopurs_runtime.Value])(m_2.UnsafePtr).V0.y).V0.a)
						_ = arr_val_foldlArray4
						res_go_foldlArray4 := Call_Data_Function_go__const(gopurs_runtime.Apply2(f_0, z_1, ((*Constructor_Main_Test3[gopurs_runtime.Value])(m_2.UnsafePtr).V0.x).V0.a), gopurs_runtime.Int(((*Constructor_Main_Test3[gopurs_runtime.Value])(m_2.UnsafePtr).V0.x).V1))
						_ = res_go_foldlArray4
						arr_go_foldlArray4 := (*[]gopurs_runtime.Value)(arr_val_foldlArray4.UnsafePtr)
						_ = arr_go_foldlArray4
						for _, v_foldlArray4 := range *arr_go_foldlArray4 {
							res_go_foldlArray4 = gopurs_runtime.Apply2(f_0, res_go_foldlArray4, v_foldlArray4)
						}
						return res_go_foldlArray4
					}(), ((*Constructor_Main_Test3[gopurs_runtime.Value])(m_2.UnsafePtr).V0.y).V1.a)
					goto end_branch_5
				} else {

				}
			}
			{
				__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_5:
			return __t5
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t6 gopurs_runtime.Value
			{
				if m_2.Type == 9 && m_2.IntVal == 3720114489 {
					__t6 = gopurs_runtime.Apply2(f_0, ((*Constructor_Main_Test1[gopurs_runtime.Value])(m_2.UnsafePtr).V0).V0, gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), z_1))
					goto end_branch_6
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 2375191994 {
					__t6 = gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, gopurs_runtime.Apply2(f_0, ((*Constructor_Main_Test2[gopurs_runtime.Value])(m_2.UnsafePtr).V0).V1, z_1), gopurs_runtime.Array(((*Constructor_Main_Test2[gopurs_runtime.Value])(m_2.UnsafePtr).V0).V0))
					goto end_branch_6
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 227416251 {
					__t6 = gopurs_runtime.Apply2(f_0, ((*Constructor_Main_Test3[gopurs_runtime.Value])(m_2.UnsafePtr).V0.x).V0.a, gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, gopurs_runtime.Apply2(f_0, ((*Constructor_Main_Test3[gopurs_runtime.Value])(m_2.UnsafePtr).V0.y).V1.a, z_1), gopurs_runtime.Array(((*Constructor_Main_Test3[gopurs_runtime.Value])(m_2.UnsafePtr).V0.y).V0.a))))
					goto end_branch_6
				} else {

				}
			}
			{
				__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_6:
			return __t6
		})}))}
	})
	return cache_Main_foldableTest
}

var cache_Main_traversableTest gopurs_runtime.Value
var once_Main_traversableTest sync.Once

func Get_Main_traversableTest() gopurs_runtime.Value {
	once_Main_traversableTest.Do(func() {
		cache_Main_traversableTest = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableTest()))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorTest()))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Main_traversableTest()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_2
			}), v_1)
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope8)])
			Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_1_0
			// TAST (Let): Apply0_2_1 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope8)])
			Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_2_1
			return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t8 gopurs_runtime.Value
				{
					if m_4.Type == 9 && m_4.IntVal == 3720114489 {
						__t8 = gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test1[gopurs_runtime.Value]{1, Rebox_Main_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v1_5))}))}
						}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Tuple_Tuple(), gopurs_runtime.Apply(f_3, ((*Constructor_Main_Test1[gopurs_runtime.Value])(m_4.UnsafePtr).V0).V0)), gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0)), gopurs_runtime.Int(((*Constructor_Main_Test1[gopurs_runtime.Value])(m_4.UnsafePtr).V0).V1))))
						goto end_branch_8
					} else {

					}
				}
				{
					if m_4.Type == 9 && m_4.IntVal == 2375191994 {
						// TAST (Let): Apply0_5_2 shape=App(Other) bindingType=Any
						Apply0_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
						_ = Apply0_5_2
						__t8 = gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test2[gopurs_runtime.Value]{1, Rebox_Main_138441832_3791647502(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v1_5))}))}
						}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Tuple_Tuple(), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Apply0_5_2)), Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_2, "Functor0"), gopurs_runtime.Value{}))), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0)), Get_Data_Semigroup_concatArray(), f_3, gopurs_runtime.Array(((*Constructor_Main_Test2[gopurs_runtime.Value])(m_4.UnsafePtr).V0).V0))), gopurs_runtime.Apply(f_3, ((*Constructor_Main_Test2[gopurs_runtime.Value])(m_4.UnsafePtr).V0).V1)))
						goto end_branch_8
					} else {

					}
				}
				{
					if m_4.Type == 9 && m_4.IntVal == 227416251 {
						// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
						__local_var_5_3 := (*Constructor_Main_Test3[gopurs_runtime.Value])(m_4.UnsafePtr).V0
						_ = __local_var_5_3
						// TAST (Let): __local_var_6_4 shape=Other bindingType=(TypeVar a$scope56)
						__local_var_6_4 := func() gopurs_runtime.Value {
							orig := (__local_var_5_3.x).V0
							_ = orig
							return gopurs_runtime.RecordDict1("a", orig.a)
						}()
						_ = __local_var_6_4
						// TAST (Let): __local_var_6_5 shape=Other bindingType=(TypeVar a$scope56)
						__local_var_6_5 := func() gopurs_runtime.Value {
							orig := (__local_var_5_3.y).V0
							_ = orig
							return gopurs_runtime.RecordDict1("a", gopurs_runtime.Array(orig.a))
						}()
						_ = __local_var_6_5
						// TAST (Let): Apply0_7_6 shape=App(Other) bindingType=Any
						Apply0_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
						_ = Apply0_7_6
						// TAST (Let): __local_var_6_7 shape=Other bindingType=(TypeVar b$scope57)
						__local_var_6_7 := func() gopurs_runtime.Value {
							orig := (__local_var_5_3.y).V1
							_ = orig
							return gopurs_runtime.RecordDict1("a", orig.a)
						}()
						_ = __local_var_6_7
						__t8 = gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func2(func(v1_6 gopurs_runtime.Value, v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer((&Constructor_Main_Test3[gopurs_runtime.Value]{1, func() struct {
								x *Constructor_Data_Tuple_Tuple[struct {
									a gopurs_runtime.Value
								}, int64]
								y *Constructor_Data_Tuple_Tuple[struct {
									a []gopurs_runtime.Value
								}, struct {
									a gopurs_runtime.Value
								}]
							} {
								clone := __local_var_5_3
								clone.x = Rebox_Main_138441832_2137532138(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v1_6))
								clone.y = Rebox_Main_138441832_372546286(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v2_7))
								return clone
							}()}))}
						}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Tuple_Tuple(), gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return func() gopurs_runtime.Value {
								orig := func() struct {
									a gopurs_runtime.Value
								} {
									orig := gopurs_runtime.RecordUpdate1(__local_var_6_4, "a", v2_7)
									_ = orig
									clone := struct {
										a gopurs_runtime.Value
									}{}
									clone.a = gopurs_runtime.RecordGet(orig, "a")
									return clone
								}()
								_ = orig
								return gopurs_runtime.RecordDict1("a", orig.a)
							}()
						}), gopurs_runtime.Apply(f_3, gopurs_runtime.RecordGet(__local_var_6_4, "a")))), gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0)), gopurs_runtime.Int((__local_var_5_3.x).V1)))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Tuple_Tuple(), gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return func() gopurs_runtime.Value {
								orig := func() struct {
									a []gopurs_runtime.Value
								} {
									orig := gopurs_runtime.RecordUpdate1(__local_var_6_5, "a", gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((v2_7).UnsafePtr))))
									_ = orig
									clone := struct {
										a []gopurs_runtime.Value
									}{}
									clone.a = (*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(orig, "a")).UnsafePtr))
									return clone
								}()
								_ = orig
								return gopurs_runtime.RecordDict1("a", gopurs_runtime.Array(orig.a))
							}()
						}), gopurs_runtime.Apply6(Get_Data_Traversable_traverseArrayImpl(), Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Apply0_7_6)), Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_6, "Functor0"), gopurs_runtime.Value{}))), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0)), Get_Data_Semigroup_concatArray(), f_3, gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(__local_var_6_5, "a")).UnsafePtr)))))), gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
							return func() gopurs_runtime.Value {
								orig := func() struct {
									a gopurs_runtime.Value
								} {
									orig := gopurs_runtime.RecordUpdate1(__local_var_6_7, "a", v2_7)
									_ = orig
									clone := struct {
										a gopurs_runtime.Value
									}{}
									clone.a = gopurs_runtime.RecordGet(orig, "a")
									return clone
								}()
								_ = orig
								return gopurs_runtime.RecordDict1("a", orig.a)
							}()
						}), gopurs_runtime.Apply(f_3, gopurs_runtime.RecordGet(__local_var_6_7, "a")))))
						goto end_branch_8
					} else {

					}
				}
				{
					__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
				}
			end_branch_8:
				return __t8
			})
		})}))}
	})
	return cache_Main_traversableTest
}

var cache_Main_eqTest gopurs_runtime.Value
var once_Main_eqTest sync.Once

func Get_Main_eqTest() gopurs_runtime.Value {
	once_Main_eqTest.Do(func() {
		cache_Main_eqTest = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqTest(dictEq_0_box)
		})
	})
	return cache_Main_eqTest
}

var cache_Main_eqTest1 gopurs_runtime.Value
var once_Main_eqTest1 sync.Once

func Get_Main_eqTest1() gopurs_runtime.Value {
	once_Main_eqTest1.Do(func() {
		cache_Main_eqTest1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Main_eqTest(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))})))}
	})
	return cache_Main_eqTest1
}

var cache_Main_check gopurs_runtime.Value
var once_Main_check sync.Once

func Get_Main_check() gopurs_runtime.Value {
	once_Main_check.Do(func() {
		cache_Main_check = gopurs_runtime.Func4(func(name_0_box gopurs_runtime.Value, value_1_box gopurs_runtime.Value, expected_2_box gopurs_runtime.Value, mapped_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_check(name_0_box.StrVal(), value_1_box, func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(expected_2_box.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}(), mapped_3_box)
		})
	})
	return cache_Main_check
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Int(int64(7)))
			_ = __local_var_0_0
			__local_var_1_2 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_2
			__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_1_2), gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_3 := gopurs_runtime.Apply(Call_Main_check("left argument", gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer(Rebox_Main_2574500310_2598952845((&Constructor_Main_Test1[int64]{1, Rebox_Main_138441832_3363075976(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Int(__local_var_1_1.IntVal), gopurs_runtime.Int(int64(42))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))})))}, []int64{int64(7)}, gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer(Rebox_Main_4002987442_2598952845((&Constructor_Main_Test1[string]{1, Rebox_Main_138441832_1830029836(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{gopurs_runtime.Str("7"), gopurs_runtime.Int(int64(42))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))})))}), gopurs_runtime.Value{})
			_ = __local_var_2_3
			__local_var_3_4 := gopurs_runtime.Apply(Call_Main_check("both arguments", gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer(Rebox_Main_3188207925_1106546350((&Constructor_Main_Test2[int64]{1, Rebox_Main_1259605934_1267868309(Rebox_Main_138441832_1259605934(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{func() gopurs_runtime.Value {
					arr := []int64{__local_var_1_1.IntVal, gopurs_runtime.IntAdd(__local_var_1_1.IntVal, int64(1))}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Int(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(gopurs_runtime.IntAdd(__local_var_1_1.IntVal, int64(2)))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}())))})))}, []int64{int64(7), int64(8), int64(9)}, gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer(Rebox_Main_1200426641_1106546350((&Constructor_Main_Test2[string]{1, Rebox_Main_839290894_2922223345(Rebox_Main_138441832_839290894(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 gopurs_runtime.Value
				}{func() gopurs_runtime.Value {
					arr := []string{"7", "8"}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						boxed[i] = gopurs_runtime.Str(v)
					}
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Str("9")}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}())))})))}), gopurs_runtime.Value{})
			_ = __local_var_3_4
			__local_var_4_5 := gopurs_runtime.Apply(Call_Main_check("nested records", gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer(Rebox_Main_697303572_767958607((&Constructor_Main_Test3[int64]{1, func(record struct {
				x *Constructor_Data_Tuple_Tuple[struct {
					a int64
				}, int64]
				y *Constructor_Data_Tuple_Tuple[struct {
					a []int64
				}, struct {
					a int64
				}]
			}) struct {
				x *Constructor_Data_Tuple_Tuple[struct {
					a int64
				}, int64]
				y *Constructor_Data_Tuple_Tuple[struct {
					a []gopurs_runtime.Value
				}, struct {
					a int64
				}]
			} {
				return struct {
					x *Constructor_Data_Tuple_Tuple[struct {
						a int64
					}, int64]
					y *Constructor_Data_Tuple_Tuple[struct {
						a []gopurs_runtime.Value
					}, struct {
						a int64
					}]
				}{record.x, Rebox_Main_4123273614_3746777429(record.y)}
			}(struct {
				x *Constructor_Data_Tuple_Tuple[struct {
					a int64
				}, int64]
				y *Constructor_Data_Tuple_Tuple[struct {
					a []int64
				}, struct {
					a int64
				}]
			}{Rebox_Main_138441832_401393041(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 struct {
						a int64
					}
					V1 gopurs_runtime.Value
				}{struct {
					a int64
				}{__local_var_1_1.IntVal}, gopurs_runtime.Int(int64(42))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict1("a", gopurs_runtime.Int(orig.a))
				}(), V1: _v.V1})}
			}())), Rebox_Main_138441832_4123273614(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 struct {
						a []int64
					}
					V1 struct {
						a int64
					}
				}{struct {
					a []int64
				}{[]int64{gopurs_runtime.IntAdd(__local_var_1_1.IntVal, int64(1)), gopurs_runtime.IntAdd(__local_var_1_1.IntVal, int64(2))}}, struct {
					a int64
				}{gopurs_runtime.IntAdd(__local_var_1_1.IntVal, int64(3))}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict1("a", func() gopurs_runtime.Value {
						arr := orig.a
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}(), V1: func() gopurs_runtime.Value {
					orig := _v.V1
					_ = orig
					return gopurs_runtime.RecordDict1("a", gopurs_runtime.Int(orig.a))
				}()})}
			}()))})})))}, []int64{int64(7), int64(8), int64(9), int64(10)}, gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer(Rebox_Main_2177728240_767958607((&Constructor_Main_Test3[string]{1, func(record struct {
				x *Constructor_Data_Tuple_Tuple[struct {
					a string
				}, int64]
				y *Constructor_Data_Tuple_Tuple[struct {
					a []string
				}, struct {
					a string
				}]
			}) struct {
				x *Constructor_Data_Tuple_Tuple[struct {
					a string
				}, int64]
				y *Constructor_Data_Tuple_Tuple[struct {
					a []gopurs_runtime.Value
				}, struct {
					a string
				}]
			} {
				return struct {
					x *Constructor_Data_Tuple_Tuple[struct {
						a string
					}, int64]
					y *Constructor_Data_Tuple_Tuple[struct {
						a []gopurs_runtime.Value
					}, struct {
						a string
					}]
				}{record.x, Rebox_Main_1971292750_3783055825(record.y)}
			}(struct {
				x *Constructor_Data_Tuple_Tuple[struct {
					a string
				}, int64]
				y *Constructor_Data_Tuple_Tuple[struct {
					a []string
				}, struct {
					a string
				}]
			}{Rebox_Main_138441832_1612762549(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 struct {
						a string
					}
					V1 gopurs_runtime.Value
				}{struct {
					a string
				}{"7"}, gopurs_runtime.Int(int64(42))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict1("a", gopurs_runtime.Str(orig.a))
				}(), V1: _v.V1})}
			}())), Rebox_Main_138441832_1971292750(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 struct {
						a []string
					}
					V1 struct {
						a string
					}
				}{struct {
					a []string
				}{[]string{"8", "9"}}, struct {
					a string
				}{"10"}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
					orig := _v.V0
					_ = orig
					return gopurs_runtime.RecordDict1("a", func() gopurs_runtime.Value {
						arr := orig.a
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Str(v)
						}
						return gopurs_runtime.Array(boxed)
					}())
				}(), V1: func() gopurs_runtime.Value {
					orig := _v.V1
					_ = orig
					return gopurs_runtime.RecordDict1("a", gopurs_runtime.Str(orig.a))
				}()})}
			}()))})})))}), gopurs_runtime.Value{})
			_ = __local_var_4_5
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Test1[T_a any] struct {
	Rc uint32
	V0 *Constructor_Data_Tuple_Tuple[T_a, int64]
}

type Constructor_Main_Test2[T_a any] struct {
	Rc uint32
	V0 *Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, T_a]
}

type Constructor_Main_Test3[T_a any] struct {
	Rc uint32
	V0 struct {
		x *Constructor_Data_Tuple_Tuple[struct {
			a T_a
		}, int64]
		y *Constructor_Data_Tuple_Tuple[struct {
			a []gopurs_runtime.Value
		}, struct {
			a T_a
		}]
	}
}

func Call_Main_bifoldl(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var z_2 gopurs_runtime.Value = z_2_loop
	_ = z_2
	var v_3 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] = v_3_loop
	_ = v_3
	return gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, z_2, (v_3).V0), gopurs_runtime.Int((v_3).V1))
}

func Call_Main_bifoldl1(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_Tuple_Tuple[struct {
	a gopurs_runtime.Value
}, int64]) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var z_2 gopurs_runtime.Value = z_2_loop
	_ = z_2
	var v_3 *Constructor_Data_Tuple_Tuple[struct {
		a gopurs_runtime.Value
	}, int64] = v_3_loop
	_ = v_3
	return gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, z_2, func() gopurs_runtime.Value {
		orig := (v_3).V0
		_ = orig
		return gopurs_runtime.RecordDict1("a", orig.a)
	}()), gopurs_runtime.Int((v_3).V1))
}

func Call_Main_bifoldr(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var z_2 gopurs_runtime.Value = z_2_loop
	_ = z_2
	var v_3 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] = v_3_loop
	_ = v_3
	return gopurs_runtime.Apply2(f_0, (v_3).V0, gopurs_runtime.Apply2(g_1, gopurs_runtime.Int((v_3).V1), z_2))
}

func Call_Main_bifoldr1(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_Tuple_Tuple[struct {
	a gopurs_runtime.Value
}, int64]) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var z_2 gopurs_runtime.Value = z_2_loop
	_ = z_2
	var v_3 *Constructor_Data_Tuple_Tuple[struct {
		a gopurs_runtime.Value
	}, int64] = v_3_loop
	_ = v_3
	return gopurs_runtime.Apply2(f_0, func() gopurs_runtime.Value {
		orig := (v_3).V0
		_ = orig
		return gopurs_runtime.RecordDict1("a", orig.a)
	}(), gopurs_runtime.Apply2(g_1, gopurs_runtime.Int((v_3).V1), z_2))
}

func Call_Main_bifoldMap(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
	_ = dictMonoid_0
	// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope54)])
	Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_0.V0, gopurs_runtime.Value{}))
	_ = Semigroup0_1_0
	return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply(g_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
	})
}

func Call_Main_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Test1__3723567844(__eta_norm_0_0_loop *Constructor_Data_Tuple_Tuple[int64, int64]) gopurs_runtime.Value {
Test1__3723567844:
	for {
		if false {
			continue Test1__3723567844
		}
		var __eta_norm_0_0 *Constructor_Data_Tuple_Tuple[int64, int64] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer(Rebox_Main_2574500310_2598952845((&Constructor_Main_Test1[int64]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_Test1__1738194564(__eta_norm_0_0_loop *Constructor_Data_Tuple_Tuple[string, int64]) gopurs_runtime.Value {
Test1__1738194564:
	for {
		if false {
			continue Test1__1738194564
		}
		var __eta_norm_0_0 *Constructor_Data_Tuple_Tuple[string, int64] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 3720114489, UnsafePtr: unsafe.Pointer(Rebox_Main_4002987442_2598952845((&Constructor_Main_Test1[string]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_Test2__2545363874(__eta_norm_0_0_loop *Constructor_Data_Tuple_Tuple[[]int64, int64]) gopurs_runtime.Value {
Test2__2545363874:
	for {
		if false {
			continue Test2__2545363874
		}
		var __eta_norm_0_0 *Constructor_Data_Tuple_Tuple[[]int64, int64] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer(Rebox_Main_3188207925_1106546350((&Constructor_Main_Test2[int64]{1, Rebox_Main_1259605934_1267868309(__eta_norm_0_0)})))}
	}
}

func Call_Main_Test2__877929284(__eta_norm_0_0_loop *Constructor_Data_Tuple_Tuple[[]string, string]) gopurs_runtime.Value {
Test2__877929284:
	for {
		if false {
			continue Test2__877929284
		}
		var __eta_norm_0_0 *Constructor_Data_Tuple_Tuple[[]string, string] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 2375191994, UnsafePtr: unsafe.Pointer(Rebox_Main_1200426641_1106546350((&Constructor_Main_Test2[string]{1, Rebox_Main_839290894_2922223345(__eta_norm_0_0)})))}
	}
}

func Call_Main_Test3__3463130812(__eta_norm_0_0_loop struct {
	x *Constructor_Data_Tuple_Tuple[struct {
		a int64
	}, int64]
	y *Constructor_Data_Tuple_Tuple[struct {
		a []int64
	}, struct {
		a int64
	}]
}) gopurs_runtime.Value {
Test3__3463130812:
	for {
		if false {
			continue Test3__3463130812
		}
		var __eta_norm_0_0 struct {
			x *Constructor_Data_Tuple_Tuple[struct {
				a int64
			}, int64]
			y *Constructor_Data_Tuple_Tuple[struct {
				a []int64
			}, struct {
				a int64
			}]
		} = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer(Rebox_Main_697303572_767958607((&Constructor_Main_Test3[int64]{1, func(record struct {
			x *Constructor_Data_Tuple_Tuple[struct {
				a int64
			}, int64]
			y *Constructor_Data_Tuple_Tuple[struct {
				a []int64
			}, struct {
				a int64
			}]
		}) struct {
			x *Constructor_Data_Tuple_Tuple[struct {
				a int64
			}, int64]
			y *Constructor_Data_Tuple_Tuple[struct {
				a []gopurs_runtime.Value
			}, struct {
				a int64
			}]
		} {
			return struct {
				x *Constructor_Data_Tuple_Tuple[struct {
					a int64
				}, int64]
				y *Constructor_Data_Tuple_Tuple[struct {
					a []gopurs_runtime.Value
				}, struct {
					a int64
				}]
			}{record.x, Rebox_Main_4123273614_3746777429(record.y)}
		}(__eta_norm_0_0)})))}
	}
}

func Call_Main_Test3__2986798300(__eta_norm_0_0_loop struct {
	x *Constructor_Data_Tuple_Tuple[struct {
		a string
	}, int64]
	y *Constructor_Data_Tuple_Tuple[struct {
		a []string
	}, struct {
		a string
	}]
}) gopurs_runtime.Value {
Test3__2986798300:
	for {
		if false {
			continue Test3__2986798300
		}
		var __eta_norm_0_0 struct {
			x *Constructor_Data_Tuple_Tuple[struct {
				a string
			}, int64]
			y *Constructor_Data_Tuple_Tuple[struct {
				a []string
			}, struct {
				a string
			}]
		} = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 227416251, UnsafePtr: unsafe.Pointer(Rebox_Main_2177728240_767958607((&Constructor_Main_Test3[string]{1, func(record struct {
			x *Constructor_Data_Tuple_Tuple[struct {
				a string
			}, int64]
			y *Constructor_Data_Tuple_Tuple[struct {
				a []string
			}, struct {
				a string
			}]
		}) struct {
			x *Constructor_Data_Tuple_Tuple[struct {
				a string
			}, int64]
			y *Constructor_Data_Tuple_Tuple[struct {
				a []gopurs_runtime.Value
			}, struct {
				a string
			}]
		} {
			return struct {
				x *Constructor_Data_Tuple_Tuple[struct {
					a string
				}, int64]
				y *Constructor_Data_Tuple_Tuple[struct {
					a []gopurs_runtime.Value
				}, struct {
					a string
				}]
			}{record.x, Rebox_Main_1971292750_3783055825(record.y)}
		}(__eta_norm_0_0)})))}
	}
}

func Call_Main_eqTest(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	// TAST (Let): eqTuple_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope39), Int])])
	eqTuple_1_0 := Rebox_Main_3790796878_2257251064(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Tuple_eqTuple(dictEq_0, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))
	_ = eqTuple_1_0
	// TAST (Let): eqArray1_2_1 shape=App(Var) bindingType=Any
	eqArray1_2_1 := Call_Data_Eq_eqArray(dictEq_0)
	_ = eqArray1_2_1
	// TAST (Let): eqTuple1_3_2 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Tuple","Tuple"] [(Array (TypeVar a$scope39)), (TypeVar a$scope39)])])
	eqTuple1_3_2 := Rebox_Main_3790796878_2737345957(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Tuple_eqTuple(eqArray1_2_1, dictEq_0)))
	_ = eqTuple1_3_2
	// TAST (Let): eqRec1_4_3 shape=App(Var) bindingType=Any
	eqRec1_4_3 := Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str("a")
	})), dictEq_0))
	_ = eqRec1_4_3
	// TAST (Let): eqTuple2_5_4 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Tuple","Tuple"] [(Record (Row [a: (TypeVar a$scope39)] Empty)), Int])])
	eqTuple2_5_4 := Rebox_Main_3790796878_3148176769(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Tuple_eqTuple(eqRec1_4_3, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))
	_ = eqTuple2_5_4
	// TAST (Let): eqTuple3_6_5 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Tuple","Tuple"] [(Record (Row [a: (Array (TypeVar a$scope39))] Empty)), (Record (Row [a: (TypeVar a$scope39)] Empty))])])
	eqTuple3_6_5 := Rebox_Main_3790796878_1848348165(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Tuple_eqTuple(Call_Data_Eq_eqRec(gopurs_runtime.Value{}, Call_Data_Eq_eqRowCons(Get_Data_Eq_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str("a")
	})), eqArray1_2_1)), eqRec1_4_3)))
	_ = eqTuple3_6_5
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_7 gopurs_runtime.Value, y_8 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t6 bool
		{
			if x_7.Type == 9 && x_7.IntVal == 3720114489 {
				__t6 = (y_8.Type == 9 && y_8.IntVal == 3720114489) && ((gopurs_runtime.Apply2(eqTuple_1_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3415943795_138441832((*Constructor_Main_Test1[gopurs_runtime.Value])(x_7.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3415943795_138441832((*Constructor_Main_Test1[gopurs_runtime.Value])(y_8.UnsafePtr).V0))}).IntVal) != (0))
				goto end_branch_6
			} else {

			}
		}
		{
			if x_7.Type == 9 && x_7.IntVal == 2375191994 {
				__t6 = (y_8.Type == 9 && y_8.IntVal == 2375191994) && ((gopurs_runtime.Apply2(eqTuple1_3_2.V0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3791647502_138441832((*Constructor_Main_Test2[gopurs_runtime.Value])(x_7.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_3791647502_138441832((*Constructor_Main_Test2[gopurs_runtime.Value])(y_8.UnsafePtr).V0))}).IntVal) != (0))
				goto end_branch_6
			} else {

			}
		}
		{
			__t6 = (x_7.Type == 9 && x_7.IntVal == 227416251) && ((y_8.Type == 9 && y_8.IntVal == 227416251) && (((gopurs_runtime.Apply2(eqTuple2_5_4.V0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_2137532138_138441832((*Constructor_Main_Test3[gopurs_runtime.Value])(x_7.UnsafePtr).V0.x))}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_2137532138_138441832((*Constructor_Main_Test3[gopurs_runtime.Value])(y_8.UnsafePtr).V0.x))}).IntVal) != (0)) && ((gopurs_runtime.Apply2(eqTuple3_6_5.V0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_372546286_138441832((*Constructor_Main_Test3[gopurs_runtime.Value])(x_7.UnsafePtr).V0.y))}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Main_372546286_138441832((*Constructor_Main_Test3[gopurs_runtime.Value])(y_8.UnsafePtr).V0.y))}).IntVal) != (0))))
		}
	end_branch_6:
		return gopurs_runtime.Bool(__t6)
	})}))}
}

func Call_Main_check(name_0_loop string, value_1_loop gopurs_runtime.Value, expected_2_loop []int64, mapped_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var name_0 string = name_0_loop
	_ = name_0
	var value_1 gopurs_runtime.Value = value_1_loop
	_ = value_1
	var expected_2 []int64 = expected_2_loop
	_ = expected_2
	var mapped_3 gopurs_runtime.Value = mapped_3_loop
	_ = mapped_3
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
		__local_var_4_0 := gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(gopurs_runtime.ConcatString(name_0, " - map")), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqTest(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorTest()).V0, Get_Data_Show_showIntImpl(), value_1), mapped_3).IntVal) != (0)))
		_ = __local_var_4_0
		__local_var_5_1 := gopurs_runtime.Apply(__local_var_4_0, gopurs_runtime.Value{})
		_ = __local_var_5_1
		__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(gopurs_runtime.ConcatString(name_0, " - foldl")), gopurs_runtime.Bool((gopurs_runtime.Apply2(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqArray())).V0, func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableTest()).V1, gopurs_runtime.Func2(func(xs_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), xs_6, func() gopurs_runtime.Value {
								arr := []int64{x_7.IntVal}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}())).UnsafePtr))).UnsafePtr)
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
					}()
				}), gopurs_runtime.Array([]gopurs_runtime.Value{}), value_1).UnsafePtr)
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
			arr := expected_2
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).IntVal) != (0))), gopurs_runtime.Value{})
		_ = __local_var_6_2
		__local_var_7_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(gopurs_runtime.ConcatString(name_0, " - foldr")), gopurs_runtime.Bool((gopurs_runtime.Apply2(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqArray())).V0, func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableTest()).V2, gopurs_runtime.Func2(func(x_7 gopurs_runtime.Value, xs_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						arr := func() []int64 {
							arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), func() gopurs_runtime.Value {
								arr := []int64{x_7.IntVal}
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = gopurs_runtime.Int(v)
								}
								return gopurs_runtime.Array(boxed)
							}(), xs_8)).UnsafePtr))).UnsafePtr)
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
					}()
				}), gopurs_runtime.Array([]gopurs_runtime.Value{}), value_1).UnsafePtr)
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
			arr := expected_2
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).IntVal) != (0))), gopurs_runtime.Value{})
		_ = __local_var_7_3
		__local_var_8_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(gopurs_runtime.ConcatString(name_0, " - foldMap")), gopurs_runtime.Bool((gopurs_runtime.Apply2(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqArray())).V0, func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableTest()).V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_267785971_1201789390(Rebox_Main_1201789390_267785971(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidArray()))))}, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
					return func() gopurs_runtime.Value {
						arr := []int64{x_8.IntVal}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = gopurs_runtime.Int(v)
						}
						return gopurs_runtime.Array(boxed)
					}()
				}), value_1).UnsafePtr)
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
			arr := expected_2
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).IntVal) != (0))), gopurs_runtime.Value{})
		_ = __local_var_8_4
		__local_var_9_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref__new(), func() gopurs_runtime.Value {
			arr := []int64{}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()), gopurs_runtime.Value{})
		_ = __local_var_9_5
		var Call_local_Main_visit_10_6 func(gopurs_runtime.Value) gopurs_runtime.Value
		_ = Call_local_Main_visit_10_6
		var visit_10_6 gopurs_runtime.Value
		_ = visit_10_6
		Call_local_Main_visit_10_6 = func(x_10_loop gopurs_runtime.Value) gopurs_runtime.Value {
			var x_10 gopurs_runtime.Value = x_10_loop
			_ = x_10
			return Call_Control_Apply_applySecond__1408114518(gopurs_runtime.Apply2(Get_Effect_Ref_modify_(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
				return func() gopurs_runtime.Value {
					arr := func() []int64 {
						arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array((*(*[]gopurs_runtime.Value)((gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), v_11, func() gopurs_runtime.Value {
							arr := []int64{x_10.IntVal}
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}())).UnsafePtr))).UnsafePtr)
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
				}()
			}), __local_var_9_5), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				// TAST (Let): __local_var_11_7 shape=App(Var) bindingType=Any
				__local_var_11_7 := Data_Show_ShowIntImpl(x_10.IntVal)
				_ = __local_var_11_7
				return gopurs_runtime.Str(__local_var_11_7)
			}))
		}
		visit_10_6 = gopurs_runtime.Func(func(x_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Main_visit_10_6(x_10_loop_val)
		})
		__local_var_11_8 := gopurs_runtime.Apply(gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Main_traversableTest()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()))}, visit_10_6, value_1), gopurs_runtime.Value{})
		_ = __local_var_11_8
		__local_var_12_9 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(gopurs_runtime.ConcatString(name_0, " - traverse")), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqTest(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), __local_var_11_8, mapped_3).IntVal) != (0))), gopurs_runtime.Value{})
		_ = __local_var_12_9
		__local_var_13_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_9_5), gopurs_runtime.Value{})
		_ = __local_var_13_10
		__local_var_14_11 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(gopurs_runtime.ConcatString(name_0, " - traverse order")), gopurs_runtime.Bool((gopurs_runtime.Apply2(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqArray())).V0, __local_var_13_10, func() gopurs_runtime.Value {
			arr := expected_2
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).IntVal) != (0))), gopurs_runtime.Value{})
		_ = __local_var_14_11
		__local_var_15_12 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Effect_Ref_write(), func() gopurs_runtime.Value {
			arr := []int64{}
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}(), __local_var_9_5), gopurs_runtime.Value{})
		_ = __local_var_15_12
		__local_var_16_13 := gopurs_runtime.Apply(gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Main_traversableTest()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()))}, gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_16
		}), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorTest()).V0, visit_10_6, value_1)), gopurs_runtime.Value{})
		_ = __local_var_16_13
		__local_var_17_14 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(gopurs_runtime.ConcatString(name_0, " - sequence")), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqTest(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}), "eq"), __local_var_16_13, mapped_3).IntVal) != (0))), gopurs_runtime.Value{})
		_ = __local_var_17_14
		__local_var_18_15 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Ref_read(), __local_var_9_5), gopurs_runtime.Value{})
		_ = __local_var_18_15
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Test_Assert_assertImpl(), gopurs_runtime.Str(gopurs_runtime.ConcatString(name_0, " - sequence order")), gopurs_runtime.Bool((gopurs_runtime.Apply2(Rebox_Main_3790796878_378698611(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqArray())).V0, __local_var_18_15, func() gopurs_runtime.Value {
			arr := expected_2
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).IntVal) != (0))), gopurs_runtime.Value{})
	})
}

func Rebox_Main_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_1200426641_1106546350(in *Constructor_Main_Test2[string]) *Constructor_Main_Test2[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Test2[gopurs_runtime.Value]{}
	out.V0 = Rebox_Main_2922223345_3791647502(in.V0)
	return out
}

func Rebox_Main_1201789390_267785971(in *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) *Constructor_Data_Monoid_Monoid[[]int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[[]int64]{}
	out.V0 = in.V0
	out.V1 = func() []int64 {
		arr := *(*[]gopurs_runtime.Value)(in.V1.UnsafePtr)
		unboxed := make([]int64, len(arr))
		for i, v := range arr {
			unboxed[i] = v.IntVal
		}
		return unboxed
	}()
	return out
}

func Rebox_Main_1259605934_1267868309(in *Constructor_Data_Tuple_Tuple[[]int64, int64]) *Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, int64]{}
	out.V0 = (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
		arr := in.V0
		boxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			boxed[i] = gopurs_runtime.Int(v)
		}
		return gopurs_runtime.Array(boxed)
	}()).UnsafePtr))
	out.V1 = in.V1
	return out
}

func Rebox_Main_1267868309_3791647502(in *Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, int64]) *Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Main_138441832_1259605934(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[[]int64, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[[]int64, int64]{}
	out.V0 = func() []int64 {
		arr := *(*[]gopurs_runtime.Value)(in.V0.UnsafePtr)
		unboxed := make([]int64, len(arr))
		for i, v := range arr {
			unboxed[i] = v.IntVal
		}
		return unboxed
	}()
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Main_138441832_1612762549(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[struct {
	a string
}, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[struct {
		a string
	}, int64]{}
	out.V0 = func() struct {
		a string
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			a string
		}{}
		clone.a = gopurs_runtime.RecordGet(orig, "a").StrVal()
		return clone
	}()
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Main_138441832_1830029836(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[string, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[string, int64]{}
	out.V0 = in.V0.StrVal()
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Main_138441832_1971292750(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[struct {
	a []string
}, struct {
	a string
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[struct {
		a []string
	}, struct {
		a string
	}]{}
	out.V0 = func() struct {
		a []string
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			a []string
		}{}
		clone.a = func() []string {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "a").UnsafePtr)
			unboxed := make([]string, len(arr))
			for i, v := range arr {
				unboxed[i] = v.StrVal()
			}
			return unboxed
		}()
		return clone
	}()
	out.V1 = func() struct {
		a string
	} {
		orig := in.V1
		_ = orig
		clone := struct {
			a string
		}{}
		clone.a = gopurs_runtime.RecordGet(orig, "a").StrVal()
		return clone
	}()
	return out
}

func Rebox_Main_138441832_2137532138(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[struct {
	a gopurs_runtime.Value
}, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[struct {
		a gopurs_runtime.Value
	}, int64]{}
	out.V0 = func() struct {
		a gopurs_runtime.Value
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			a gopurs_runtime.Value
		}{}
		clone.a = gopurs_runtime.RecordGet(orig, "a")
		return clone
	}()
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Main_138441832_3363075976(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[int64, int64]{}
	out.V0 = in.V0.IntVal
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Main_138441832_3415943795(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]{}
	out.V0 = in.V0
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Main_138441832_372546286(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[struct {
	a []gopurs_runtime.Value
}, struct {
	a gopurs_runtime.Value
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[struct {
		a []gopurs_runtime.Value
	}, struct {
		a gopurs_runtime.Value
	}]{}
	out.V0 = func() struct {
		a []gopurs_runtime.Value
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			a []gopurs_runtime.Value
		}{}
		clone.a = (*(*[]gopurs_runtime.Value)((gopurs_runtime.RecordGet(orig, "a")).UnsafePtr))
		return clone
	}()
	out.V1 = func() struct {
		a gopurs_runtime.Value
	} {
		orig := in.V1
		_ = orig
		clone := struct {
			a gopurs_runtime.Value
		}{}
		clone.a = gopurs_runtime.RecordGet(orig, "a")
		return clone
	}()
	return out
}

func Rebox_Main_138441832_3791647502(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = (*(*[]gopurs_runtime.Value)((in.V0).UnsafePtr))
	out.V1 = in.V1
	return out
}

func Rebox_Main_138441832_401393041(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[struct {
	a int64
}, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[struct {
		a int64
	}, int64]{}
	out.V0 = func() struct {
		a int64
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			a int64
		}{}
		clone.a = gopurs_runtime.RecordGet(orig, "a").IntVal
		return clone
	}()
	out.V1 = in.V1.IntVal
	return out
}

func Rebox_Main_138441832_4123273614(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[struct {
	a []int64
}, struct {
	a int64
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[struct {
		a []int64
	}, struct {
		a int64
	}]{}
	out.V0 = func() struct {
		a []int64
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			a []int64
		}{}
		clone.a = func() []int64 {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "a").UnsafePtr)
			unboxed := make([]int64, len(arr))
			for i, v := range arr {
				unboxed[i] = v.IntVal
			}
			return unboxed
		}()
		return clone
	}()
	out.V1 = func() struct {
		a int64
	} {
		orig := in.V1
		_ = orig
		clone := struct {
			a int64
		}{}
		clone.a = gopurs_runtime.RecordGet(orig, "a").IntVal
		return clone
	}()
	return out
}

func Rebox_Main_138441832_839290894(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[[]string, string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[[]string, string]{}
	out.V0 = func() []string {
		arr := *(*[]gopurs_runtime.Value)(in.V0.UnsafePtr)
		unboxed := make([]string, len(arr))
		for i, v := range arr {
			unboxed[i] = v.StrVal()
		}
		return unboxed
	}()
	out.V1 = in.V1.StrVal()
	return out
}

func Rebox_Main_1612762549_2137532138(in *Constructor_Data_Tuple_Tuple[struct {
	a string
}, int64]) *Constructor_Data_Tuple_Tuple[struct {
	a gopurs_runtime.Value
}, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[struct {
		a gopurs_runtime.Value
	}, int64]{}
	out.V0 = func(record struct {
		a string
	}) struct {
		a gopurs_runtime.Value
	} {
		return struct {
			a gopurs_runtime.Value
		}{gopurs_runtime.Str(record.a)}
	}(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Main_1830029836_3415943795(in *Constructor_Data_Tuple_Tuple[string, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Main_1971292750_3783055825(in *Constructor_Data_Tuple_Tuple[struct {
	a []string
}, struct {
	a string
}]) *Constructor_Data_Tuple_Tuple[struct {
	a []gopurs_runtime.Value
}, struct {
	a string
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[struct {
		a []gopurs_runtime.Value
	}, struct {
		a string
	}]{}
	out.V0 = func(record struct {
		a []string
	}) struct {
		a []gopurs_runtime.Value
	} {
		return struct {
			a []gopurs_runtime.Value
		}{(*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := record.a
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).UnsafePtr))}
	}(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Main_2137532138_138441832(in *Constructor_Data_Tuple_Tuple[struct {
	a gopurs_runtime.Value
}, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict1("a", orig.a)
	}()
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Main_2177728240_767958607(in *Constructor_Main_Test3[string]) *Constructor_Main_Test3[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Test3[gopurs_runtime.Value]{}
	out.V0 = func(record struct {
		x *Constructor_Data_Tuple_Tuple[struct {
			a string
		}, int64]
		y *Constructor_Data_Tuple_Tuple[struct {
			a []gopurs_runtime.Value
		}, struct {
			a string
		}]
	}) struct {
		x *Constructor_Data_Tuple_Tuple[struct {
			a gopurs_runtime.Value
		}, int64]
		y *Constructor_Data_Tuple_Tuple[struct {
			a []gopurs_runtime.Value
		}, struct {
			a gopurs_runtime.Value
		}]
	} {
		return struct {
			x *Constructor_Data_Tuple_Tuple[struct {
				a gopurs_runtime.Value
			}, int64]
			y *Constructor_Data_Tuple_Tuple[struct {
				a []gopurs_runtime.Value
			}, struct {
				a gopurs_runtime.Value
			}]
		}{Rebox_Main_1612762549_2137532138(record.x), Rebox_Main_3783055825_372546286(record.y)}
	}(in.V0)
	return out
}

func Rebox_Main_2574500310_2598952845(in *Constructor_Main_Test1[int64]) *Constructor_Main_Test1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Test1[gopurs_runtime.Value]{}
	out.V0 = Rebox_Main_3363075976_3415943795(in.V0)
	return out
}

func Rebox_Main_267785971_1201789390(in *Constructor_Data_Monoid_Monoid[[]int64]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = func() gopurs_runtime.Value {
		arr := in.V1
		boxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			boxed[i] = gopurs_runtime.Int(v)
		}
		return gopurs_runtime.Array(boxed)
	}()
	return out
}

func Rebox_Main_2922223345_3791647502(in *Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, string]) *Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Str(in.V1)
	return out
}

func Rebox_Main_3188207925_1106546350(in *Constructor_Main_Test2[int64]) *Constructor_Main_Test2[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Test2[gopurs_runtime.Value]{}
	out.V0 = Rebox_Main_1267868309_3791647502(in.V0)
	return out
}

func Rebox_Main_3363075976_3415943795(in *Constructor_Data_Tuple_Tuple[int64, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Main_3415943795_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Main_372546286_138441832(in *Constructor_Data_Tuple_Tuple[struct {
	a []gopurs_runtime.Value
}, struct {
	a gopurs_runtime.Value
}]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict1("a", gopurs_runtime.Array(orig.a))
	}()
	out.V1 = func() gopurs_runtime.Value {
		orig := in.V1
		_ = orig
		return gopurs_runtime.RecordDict1("a", orig.a)
	}()
	return out
}

func Rebox_Main_3746777429_372546286(in *Constructor_Data_Tuple_Tuple[struct {
	a []gopurs_runtime.Value
}, struct {
	a int64
}]) *Constructor_Data_Tuple_Tuple[struct {
	a []gopurs_runtime.Value
}, struct {
	a gopurs_runtime.Value
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[struct {
		a []gopurs_runtime.Value
	}, struct {
		a gopurs_runtime.Value
	}]{}
	out.V0 = in.V0
	out.V1 = func(record struct {
		a int64
	}) struct {
		a gopurs_runtime.Value
	} {
		return struct {
			a gopurs_runtime.Value
		}{gopurs_runtime.Int(record.a)}
	}(in.V1)
	return out
}

func Rebox_Main_3783055825_372546286(in *Constructor_Data_Tuple_Tuple[struct {
	a []gopurs_runtime.Value
}, struct {
	a string
}]) *Constructor_Data_Tuple_Tuple[struct {
	a []gopurs_runtime.Value
}, struct {
	a gopurs_runtime.Value
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[struct {
		a []gopurs_runtime.Value
	}, struct {
		a gopurs_runtime.Value
	}]{}
	out.V0 = in.V0
	out.V1 = func(record struct {
		a string
	}) struct {
		a gopurs_runtime.Value
	} {
		return struct {
			a gopurs_runtime.Value
		}{gopurs_runtime.Str(record.a)}
	}(in.V1)
	return out
}

func Rebox_Main_378698611_3790796878(in *Constructor_Data_Eq_Eq[[]int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1053099733(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[int64])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[string])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_1848348165(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[struct {
	a []gopurs_runtime.Value
}, struct {
	a gopurs_runtime.Value
}]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[struct {
		a []gopurs_runtime.Value
	}, struct {
		a gopurs_runtime.Value
	}]])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_2257251064(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_2737345957(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_3148176769(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[struct {
	a gopurs_runtime.Value
}, int64]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[*Constructor_Data_Tuple_Tuple[struct {
		a gopurs_runtime.Value
	}, int64]])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_378698611(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[[]int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[[]int64])(unsafe.Pointer(in))
}

func Rebox_Main_3791647502_138441832(in *Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Array(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Main_4002987442_2598952845(in *Constructor_Main_Test1[string]) *Constructor_Main_Test1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Test1[gopurs_runtime.Value]{}
	out.V0 = Rebox_Main_1830029836_3415943795(in.V0)
	return out
}

func Rebox_Main_401393041_2137532138(in *Constructor_Data_Tuple_Tuple[struct {
	a int64
}, int64]) *Constructor_Data_Tuple_Tuple[struct {
	a gopurs_runtime.Value
}, int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[struct {
		a gopurs_runtime.Value
	}, int64]{}
	out.V0 = func(record struct {
		a int64
	}) struct {
		a gopurs_runtime.Value
	} {
		return struct {
			a gopurs_runtime.Value
		}{gopurs_runtime.Int(record.a)}
	}(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Main_4123273614_3746777429(in *Constructor_Data_Tuple_Tuple[struct {
	a []int64
}, struct {
	a int64
}]) *Constructor_Data_Tuple_Tuple[struct {
	a []gopurs_runtime.Value
}, struct {
	a int64
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[struct {
		a []gopurs_runtime.Value
	}, struct {
		a int64
	}]{}
	out.V0 = func(record struct {
		a []int64
	}) struct {
		a []gopurs_runtime.Value
	} {
		return struct {
			a []gopurs_runtime.Value
		}{(*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
			arr := record.a
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()).UnsafePtr))}
	}(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Main_697303572_767958607(in *Constructor_Main_Test3[int64]) *Constructor_Main_Test3[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Test3[gopurs_runtime.Value]{}
	out.V0 = func(record struct {
		x *Constructor_Data_Tuple_Tuple[struct {
			a int64
		}, int64]
		y *Constructor_Data_Tuple_Tuple[struct {
			a []gopurs_runtime.Value
		}, struct {
			a int64
		}]
	}) struct {
		x *Constructor_Data_Tuple_Tuple[struct {
			a gopurs_runtime.Value
		}, int64]
		y *Constructor_Data_Tuple_Tuple[struct {
			a []gopurs_runtime.Value
		}, struct {
			a gopurs_runtime.Value
		}]
	} {
		return struct {
			x *Constructor_Data_Tuple_Tuple[struct {
				a gopurs_runtime.Value
			}, int64]
			y *Constructor_Data_Tuple_Tuple[struct {
				a []gopurs_runtime.Value
			}, struct {
				a gopurs_runtime.Value
			}]
		}{Rebox_Main_401393041_2137532138(record.x), Rebox_Main_3746777429_372546286(record.y)}
	}(in.V0)
	return out
}

func Rebox_Main_839290894_2922223345(in *Constructor_Data_Tuple_Tuple[[]string, string]) *Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, string]{}
	out.V0 = (*(*[]gopurs_runtime.Value)((func() gopurs_runtime.Value {
		arr := in.V0
		boxed := make([]gopurs_runtime.Value, len(arr))
		for i, v := range arr {
			boxed[i] = gopurs_runtime.Str(v)
		}
		return gopurs_runtime.Array(boxed)
	}()).UnsafePtr))
	out.V1 = in.V1
	return out
}
