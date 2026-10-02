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

var cache_Main_show gopurs_runtime.Value
var once_Main_show sync.Once

func Get_Main_show() gopurs_runtime.Value {
	once_Main_show.Do(func() {
		cache_Main_show = Get_Data_Show_showIntImpl()
	})
	return cache_Main_show
}

var cache_Main_Wrapper gopurs_runtime.Value
var once_Main_Wrapper sync.Once

func Get_Main_Wrapper() gopurs_runtime.Value {
	once_Main_Wrapper.Do(func() {
		cache_Main_Wrapper = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_Wrapper(x_0_box.StrVal()))
		})
	})
	return cache_Main_Wrapper
}

var cache_Main_Pair gopurs_runtime.Value
var once_Main_Pair sync.Once

func Get_Main_Pair() gopurs_runtime.Value {
	once_Main_Pair.Do(func() {
		cache_Main_Pair = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer((&Constructor_Main_Pair[gopurs_runtime.Value]{1, value0, value1}))}
			})
		})
	})
	return cache_Main_Pair
}

var cache_Main_Pair__1394429419 gopurs_runtime.Value
var once_Main_Pair__1394429419 sync.Once

func Get_Main_Pair__1394429419() gopurs_runtime.Value {
	once_Main_Pair__1394429419.Do(func() {
		cache_Main_Pair__1394429419 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer(Rebox_Main_1557125915_791404512(Call_Main_Pair__1394429419(__eta_norm_1_0_box.IntVal, __eta_norm_0_1_box.IntVal)))}
		})
	})
	return cache_Main_Pair__1394429419
}

var cache_Main_Name gopurs_runtime.Value
var once_Main_Name sync.Once

func Get_Main_Name() gopurs_runtime.Value {
	once_Main_Name.Do(func() {
		cache_Main_Name = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_Name(x_0_box.StrVal()))
		})
	})
	return cache_Main_Name
}

var cache_Main_Nil gopurs_runtime.Value
var once_Main_Nil sync.Once

func Get_Main_Nil() gopurs_runtime.Value {
	once_Main_Nil.Do(func() {
		cache_Main_Nil = gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Nil
}

var cache_Main_Cons gopurs_runtime.Value
var once_Main_Cons sync.Once

func Get_Main_Cons() gopurs_runtime.Value {
	once_Main_Cons.Do(func() {
		cache_Main_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](value1)}))}
			})
		})
	})
	return cache_Main_Cons
}

var cache_Main_Cons__41240261 gopurs_runtime.Value
var once_Main_Cons__41240261 sync.Once

func Get_Main_Cons__41240261() gopurs_runtime.Value {
	once_Main_Cons__41240261.Do(func() {
		cache_Main_Cons__41240261 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Rebox_Main_2737593216_176455803(Call_Main_Cons__41240261(__eta_norm_1_0_box.IntVal, Rebox_Main_176455803_2737593216(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](__eta_norm_0_unused_1_box)))))}
		})
	})
	return cache_Main_Cons__41240261
}

var cache_Main_Cons__2678532228 gopurs_runtime.Value
var once_Main_Cons__2678532228 sync.Once

func Get_Main_Cons__2678532228() gopurs_runtime.Value {
	once_Main_Cons__2678532228.Do(func() {
		cache_Main_Cons__2678532228 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Rebox_Main_2737593216_176455803(Call_Main_Cons__2678532228(__eta_norm_1_0_box.IntVal, Rebox_Main_176455803_2737593216(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](__eta_norm_0_1_box)))))}
		})
	})
	return cache_Main_Cons__2678532228
}

var cache_Main_Left2 gopurs_runtime.Value
var once_Main_Left2 sync.Once

func Get_Main_Left2() gopurs_runtime.Value {
	once_Main_Left2.Do(func() {
		cache_Main_Left2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3583471031, UnsafePtr: unsafe.Pointer((&Constructor_Main_Left2[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Left2
}

var cache_Main_Left2__1237640998 gopurs_runtime.Value
var once_Main_Left2__1237640998 sync.Once

func Get_Main_Left2__1237640998() gopurs_runtime.Value {
	once_Main_Left2__1237640998.Do(func() {
		cache_Main_Left2__1237640998 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Left2__1237640998(__eta_norm_0_0_box.IntVal)
		})
	})
	return cache_Main_Left2__1237640998
}

var cache_Main_Right2 gopurs_runtime.Value
var once_Main_Right2 sync.Once

func Get_Main_Right2() gopurs_runtime.Value {
	once_Main_Right2.Do(func() {
		cache_Main_Right2 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4276485804, UnsafePtr: unsafe.Pointer((&Constructor_Main_Right2[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Right2
}

var cache_Main_North gopurs_runtime.Value
var once_Main_North sync.Once

func Get_Main_North() gopurs_runtime.Value {
	once_Main_North.Do(func() {
		cache_Main_North = gopurs_runtime.Value{Type: 9, IntVal: int64(663586993), UnsafePtr: nil}
	})
	return cache_Main_North
}

var cache_Main_South gopurs_runtime.Value
var once_Main_South sync.Once

func Get_Main_South() gopurs_runtime.Value {
	once_Main_South.Do(func() {
		cache_Main_South = gopurs_runtime.Value{Type: 9, IntVal: int64(3523815787), UnsafePtr: nil}
	})
	return cache_Main_South
}

var cache_Main_East gopurs_runtime.Value
var once_Main_East sync.Once

func Get_Main_East() gopurs_runtime.Value {
	once_Main_East.Do(func() {
		cache_Main_East = gopurs_runtime.Value{Type: 9, IntVal: int64(2291047933), UnsafePtr: nil}
	})
	return cache_Main_East
}

var cache_Main_West gopurs_runtime.Value
var once_Main_West sync.Once

func Get_Main_West() gopurs_runtime.Value {
	once_Main_West.Do(func() {
		cache_Main_West = gopurs_runtime.Value{Type: 9, IntVal: int64(2882997739), UnsafePtr: nil}
	})
	return cache_Main_West
}

var cache_Main_Red gopurs_runtime.Value
var once_Main_Red sync.Once

func Get_Main_Red() gopurs_runtime.Value {
	once_Main_Red.Do(func() {
		cache_Main_Red = gopurs_runtime.Value{Type: 9, IntVal: int64(1227005933), UnsafePtr: nil}
	})
	return cache_Main_Red
}

var cache_Main_Green gopurs_runtime.Value
var once_Main_Green sync.Once

func Get_Main_Green() gopurs_runtime.Value {
	once_Main_Green.Do(func() {
		cache_Main_Green = gopurs_runtime.Value{Type: 9, IntVal: int64(3772422949), UnsafePtr: nil}
	})
	return cache_Main_Green
}

var cache_Main_Blue gopurs_runtime.Value
var once_Main_Blue sync.Once

func Get_Main_Blue() gopurs_runtime.Value {
	once_Main_Blue.Do(func() {
		cache_Main_Blue = gopurs_runtime.Value{Type: 9, IntVal: int64(3944123360), UnsafePtr: nil}
	})
	return cache_Main_Blue
}

var cache_Main_Empty gopurs_runtime.Value
var once_Main_Empty sync.Once

func Get_Main_Empty() gopurs_runtime.Value {
	once_Main_Empty.Do(func() {
		cache_Main_Empty = gopurs_runtime.Value{Type: 9, IntVal: 1145268909, UnsafePtr: unsafe.Pointer((*Constructor_Main_Full[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Empty
}

var cache_Main_Full gopurs_runtime.Value
var once_Main_Full sync.Once

func Get_Main_Full() gopurs_runtime.Value {
	once_Main_Full.Do(func() {
		cache_Main_Full = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1145268909, UnsafePtr: unsafe.Pointer((&Constructor_Main_Full[gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_Full
}

var cache_Main_Full__3180329226 gopurs_runtime.Value
var once_Main_Full__3180329226 sync.Once

func Get_Main_Full__3180329226() gopurs_runtime.Value {
	once_Main_Full__3180329226.Do(func() {
		cache_Main_Full__3180329226 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1145268909, UnsafePtr: unsafe.Pointer(Rebox_Main_2700107938_1539742553(Call_Main_Full__3180329226(__eta_norm_0_0_box.IntVal)))}
		})
	})
	return cache_Main_Full__3180329226
}

var cache_Main_newtypeWrapper gopurs_runtime.Value
var once_Main_newtypeWrapper sync.Once

func Get_Main_newtypeWrapper() gopurs_runtime.Value {
	once_Main_newtypeWrapper.Do(func() {
		cache_Main_newtypeWrapper = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Main_2199435624_385277032((&Constructor_Data_Newtype_Newtype[string, string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Main_newtypeWrapper
}

var cache_Main_genericDirection gopurs_runtime.Value
var once_Main_genericDirection sync.Once

func Get_Main_genericDirection() gopurs_runtime.Value {
	once_Main_genericDirection.Do(func() {
		cache_Main_genericDirection = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(Rebox_Main_660272253_2818661616((&Constructor_Data_Generic_Rep_Generic[uint32, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 gopurs_runtime.Value
			{
				var __t_tag_0 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_0
				if uint32(__t_tag_0) == 663586993 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(Rebox_Main_497217223_1323331594((&Constructor_Data_Generic_Rep_Inl[uint32, gopurs_runtime.Value]{1, 1454898258})))}
					goto end_branch_4
				} else {

				}
			}
			{
				var __t_tag_1 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_1
				if uint32(__t_tag_1) == 3523815787 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Main_660678937_2687169876((&Constructor_Data_Generic_Rep_Inr[uint32, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(Rebox_Main_497217223_1323331594((&Constructor_Data_Generic_Rep_Inl[uint32, gopurs_runtime.Value]{1, 1454898258})))}})))}
					goto end_branch_4
				} else {

				}
			}
			{
				var __t_tag_2 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_2
				if uint32(__t_tag_2) == 2291047933 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Main_660678937_2687169876((&Constructor_Data_Generic_Rep_Inr[uint32, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Main_660678937_2687169876((&Constructor_Data_Generic_Rep_Inr[uint32, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(Rebox_Main_4191096074_1323331594((&Constructor_Data_Generic_Rep_Inl[uint32, uint32]{1, 1454898258})))}})))}})))}
					goto end_branch_4
				} else {

				}
			}
			{
				var __t_tag_3 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 2882997739 {
					__t4 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Main_660678937_2687169876((&Constructor_Data_Generic_Rep_Inr[uint32, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Main_660678937_2687169876((&Constructor_Data_Generic_Rep_Inr[uint32, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Main_1259967060_2687169876((&Constructor_Data_Generic_Rep_Inr[uint32, uint32]{1, 1454898258})))}})))}})))}
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_4:
			return __t4
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t11 uint32
			{
				if x_0.Type == 9 && x_0.IntVal == 3478632216 {
					__t11 = 663586993
					goto end_branch_11
				} else {

				}
			}
			{
				if x_0.Type == 9 && x_0.IntVal == 492034566 {
					var __t10 uint32
					{
						var __t_tag_5 gopurs_runtime.Value = (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0
						_ = __t_tag_5
						if __t_tag_5.Type == 9 && __t_tag_5.IntVal == 3478632216 {
							__t10 = 3523815787
							goto end_branch_10
						} else {

						}
					}
					{
						var __t_tag_6 gopurs_runtime.Value = (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0
						_ = __t_tag_6
						if __t_tag_6.Type == 9 && __t_tag_6.IntVal == 492034566 {
							var __t9 uint32
							{
								var __t_tag_7 gopurs_runtime.Value = (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0.UnsafePtr).V0
								_ = __t_tag_7
								if __t_tag_7.Type == 9 && __t_tag_7.IntVal == 3478632216 {
									__t9 = 2291047933
									goto end_branch_9
								} else {

								}
							}
							{
								var __t_tag_8 gopurs_runtime.Value = (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0.UnsafePtr).V0
								_ = __t_tag_8
								if __t_tag_8.Type == 9 && __t_tag_8.IntVal == 492034566 {
									__t9 = 2882997739
									goto end_branch_9
								} else {

								}
							}
							{
								__t9 = func() uint32 { panic("Failed pattern match") }()
							}
						end_branch_9:
							__t10 = __t9
							goto end_branch_10
						} else {

						}
					}
					{
						__t10 = func() uint32 { panic("Failed pattern match") }()
					}
				end_branch_10:
					__t11 = __t10
					goto end_branch_11
				} else {

				}
			}
			{
				__t11 = func() uint32 { panic("Failed pattern match") }()
			}
		end_branch_11:
			return gopurs_runtime.Value{Type: 9, IntVal: int64(__t11), UnsafePtr: nil}
		})})))}
	})
	return cache_Main_genericDirection
}

var cache_Main_functorPair gopurs_runtime.Value
var once_Main_functorPair sync.Once

func Get_Main_functorPair() gopurs_runtime.Value {
	once_Main_functorPair.Do(func() {
		cache_Main_functorPair = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Main_451755851_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Main_Pair[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 893478516, UnsafePtr: unsafe.Pointer((&Constructor_Main_Pair[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Main_Pair[gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply(f_0, (*Constructor_Main_Pair[gopurs_runtime.Value])(m_1.UnsafePtr).V1)}))}
		})})))}
	})
	return cache_Main_functorPair
}

var cache_Main_functorList gopurs_runtime.Value
var once_Main_functorList sync.Once

func Get_Main_functorList() gopurs_runtime.Value {
	once_Main_functorList.Do(func() {
		cache_Main_functorList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Main_2834964816_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Main_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 *Constructor_Main_Cons[gopurs_runtime.Value]
			{
				if m_1.Type == 9 && m_1.IntVal == 322902991 && m_1.UnsafePtr == nil {
					__t0 = (*Constructor_Main_Cons[gopurs_runtime.Value])(nil)
					goto end_branch_0
				} else {

				}
			}
			{
				if m_1.Type == 9 && m_1.IntVal == 322902991 && m_1.UnsafePtr != nil {
					__t0 = (&Constructor_Main_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Main_Cons[gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Main_2812149806_2834964816(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorList())).V0, f_0, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(m_1.UnsafePtr).V1)}))})
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() *Constructor_Main_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_0:
			return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(__t0)}
		})})))}
	})
	return cache_Main_functorList
}

var cache_Main_functorBox gopurs_runtime.Value
var once_Main_functorBox sync.Once

func Get_Main_functorBox() gopurs_runtime.Value {
	once_Main_functorBox.Do(func() {
		cache_Main_functorBox = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Main_2362652146_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Main_Full[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 *Constructor_Main_Full[gopurs_runtime.Value]
			{
				if m_1.Type == 9 && m_1.IntVal == 1145268909 && m_1.UnsafePtr == nil {
					__t0 = (*Constructor_Main_Full[gopurs_runtime.Value])(nil)
					goto end_branch_0
				} else {

				}
			}
			{
				if m_1.Type == 9 && m_1.IntVal == 1145268909 && m_1.UnsafePtr != nil {
					__t0 = (&Constructor_Main_Full[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Main_Full[gopurs_runtime.Value])(m_1.UnsafePtr).V0)})
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = func() *Constructor_Main_Full[gopurs_runtime.Value] { panic("Failed pattern match") }()
			}
		end_branch_0:
			return gopurs_runtime.Value{Type: 9, IntVal: 1145268909, UnsafePtr: unsafe.Pointer(__t0)}
		})})))}
	})
	return cache_Main_functorBox
}

var cache_Main_foldableList gopurs_runtime.Value
var once_Main_foldableList sync.Once

func Get_Main_foldableList() gopurs_runtime.Value {
	once_Main_foldableList.Do(func() {
		cache_Main_foldableList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Main_464825424_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Main_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): mempty_1_0 shape=App(Var) bindingType=(TypeVar m$scope79)
			mempty_1_0 := Call_Data_Monoid_mempty(dictMonoid_0)
			_ = mempty_1_0
			// TAST (Let): Semigroup0_2_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope79)])
			Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
			_ = Semigroup0_2_1
			return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t2 gopurs_runtime.Value
				{
					if m_4.Type == 9 && m_4.IntVal == 322902991 && m_4.UnsafePtr == nil {
						__t2 = mempty_1_0
						goto end_branch_2
					} else {

					}
				}
				{
					if m_4.Type == 9 && m_4.IntVal == 322902991 && m_4.UnsafePtr != nil {
						__t2 = gopurs_runtime.Apply2(Semigroup0_2_1.V0, gopurs_runtime.Apply(f_3, (*Constructor_Main_Cons[gopurs_runtime.Value])(m_4.UnsafePtr).V0), gopurs_runtime.Apply3(Rebox_Main_1680800814_464825424(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableList())).V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, f_3, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(m_4.UnsafePtr).V1)}))
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
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t3 gopurs_runtime.Value
			{
				if m_2.Type == 9 && m_2.IntVal == 322902991 && m_2.UnsafePtr == nil {
					__t3 = z_1
					goto end_branch_3
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 322902991 && m_2.UnsafePtr != nil {
					__t3 = gopurs_runtime.Apply3(Rebox_Main_1680800814_464825424(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableList())).V1, f_0, gopurs_runtime.Apply2(f_0, z_1, (*Constructor_Main_Cons[gopurs_runtime.Value])(m_2.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(m_2.UnsafePtr).V1)})
					goto end_branch_3
				} else {

				}
			}
			{
				__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_3:
			return __t3
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t4 gopurs_runtime.Value
			{
				if m_2.Type == 9 && m_2.IntVal == 322902991 && m_2.UnsafePtr == nil {
					__t4 = z_1
					goto end_branch_4
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 322902991 && m_2.UnsafePtr != nil {
					__t4 = gopurs_runtime.Apply2(f_0, (*Constructor_Main_Cons[gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Apply3(Rebox_Main_1680800814_464825424(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableList())).V2, f_0, z_1, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(m_2.UnsafePtr).V1)}))
					goto end_branch_4
				} else {

				}
			}
			{
				__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
			}
		end_branch_4:
			return __t4
		})})))}
	})
	return cache_Main_foldableList
}

var cache_Main_traversableList gopurs_runtime.Value
var once_Main_traversableList sync.Once

func Get_Main_traversableList() gopurs_runtime.Value {
	once_Main_traversableList.Do(func() {
		cache_Main_traversableList = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Main_4169608016_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Main_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Main_464825424_1680800814(Rebox_Main_1680800814_464825424(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableList()))))}
		}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Main_2834964816_2812149806(Rebox_Main_2812149806_2834964816(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Main_functorList()))))}
		}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Rebox_Main_3043886126_4169608016(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Main_traversableList())).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return x_2
			}), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](v_1))})
		}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope8)])
			Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
			_ = Apply0_1_0
			// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope8)])
			Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
			_ = Functor0_2_1
			return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t2 gopurs_runtime.Value
				{
					if m_4.Type == 9 && m_4.IntVal == 322902991 && m_4.UnsafePtr == nil {
						__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(nil))}))})
						goto end_branch_2
					} else {

					}
				}
				{
					if m_4.Type == 9 && m_4.IntVal == 322902991 && m_4.UnsafePtr != nil {
						__t2 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func2(func(v2_5 gopurs_runtime.Value, v3_6 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((&Constructor_Main_Cons[gopurs_runtime.Value]{1, v2_5, gopurs_runtime.CoerceToStruct[Constructor_Main_Cons[gopurs_runtime.Value]](v3_6)}))}
						}), gopurs_runtime.Apply(f_3, (*Constructor_Main_Cons[gopurs_runtime.Value])(m_4.UnsafePtr).V0)), gopurs_runtime.Apply3(Rebox_Main_3043886126_4169608016(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Main_traversableList())).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, f_3, gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer((*Constructor_Main_Cons[gopurs_runtime.Value])(m_4.UnsafePtr).V1)}))
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
		})})))}
	})
	return cache_Main_traversableList
}

var cache_Main_eqName gopurs_runtime.Value
var once_Main_eqName sync.Once

func Get_Main_eqName() gopurs_runtime.Value {
	once_Main_eqName.Do(func() {
		cache_Main_eqName = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878((&Constructor_Data_Eq_Eq[string]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((x_0.StrVal()) == (y_1.StrVal()))
		})})))}
	})
	return cache_Main_eqName
}

var cache_Main_ordName gopurs_runtime.Value
var once_Main_ordName sync.Once

func Get_Main_ordName() gopurs_runtime.Value {
	once_Main_ordName.Do(func() {
		cache_Main_ordName = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Main_2406510097_4177771502((&Constructor_Data_Ord_Ord[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1140313009_3790796878(Rebox_Main_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqName()))))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str(x_0.StrVal()), gopurs_runtime.Str(y_1.StrVal())).IntVal)), UnsafePtr: nil}
		})})))}
	})
	return cache_Main_ordName
}

var cache_Main_eqEither2 gopurs_runtime.Value
var once_Main_eqEither2 sync.Once

func Get_Main_eqEither2() gopurs_runtime.Value {
	once_Main_eqEither2.Do(func() {
		cache_Main_eqEither2 = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqEither2(dictEq_0_box)
		})
	})
	return cache_Main_eqEither2
}

var cache_Main_eqEither21 gopurs_runtime.Value
var once_Main_eqEither21 sync.Once

func Get_Main_eqEither21() gopurs_runtime.Value {
	once_Main_eqEither21.Do(func() {
		cache_Main_eqEither21 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Main_eqEither2(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))}
	})
	return cache_Main_eqEither21
}

var cache_Main_eqColor gopurs_runtime.Value
var once_Main_eqColor sync.Once

func Get_Main_eqColor() gopurs_runtime.Value {
	once_Main_eqColor.Do(func() {
		cache_Main_eqColor = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t7 bool
			{
				var __t_tag_3 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 1227005933 {
					var __t_tag_4 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_4
					__t7 = (uint32(__t_tag_4) == 1227005933)
					goto end_branch_7
				} else {

				}
			}
			{
				var __t_tag_5 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_5
				if uint32(__t_tag_5) == 3772422949 {
					var __t_tag_6 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_6
					__t7 = (uint32(__t_tag_6) == 3772422949)
					goto end_branch_7
				} else {

				}
			}
			{
				var __t_tag_0 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_0
				var __t_and_2 bool = false
				if uint32(__t_tag_0) == 3944123360 {

					var __t_tag_1 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_1
					__t_and_2 = (uint32(__t_tag_1) == 3944123360)
				}
				__t7 = __t_and_2
			}
		end_branch_7:
			return gopurs_runtime.Bool(__t7)
		})})))}
	})
	return cache_Main_eqColor
}

var cache_Main_ordColor gopurs_runtime.Value
var once_Main_ordColor sync.Once

func Get_Main_ordColor() gopurs_runtime.Value {
	once_Main_ordColor.Do(func() {
		cache_Main_ordColor = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Main_3730953251_4177771502((&Constructor_Data_Ord_Ord[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_3768443459_3790796878(Rebox_Main_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqColor()))))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t11 uint32
			{
				var __t_tag_0 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_0
				if uint32(__t_tag_0) == 1227005933 {
					var __t2 uint32
					{
						var __t_tag_1 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_1
						if uint32(__t_tag_1) == 1227005933 {
							__t2 = 902936544
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = 1527465420
					}
				end_branch_2:
					__t11 = __t2
					goto end_branch_11
				} else {

				}
			}
			{
				var __t_tag_3 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 1227005933 {
					__t11 = 380165415
					goto end_branch_11
				} else {

				}
			}
			{
				var __t_tag_4 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_4
				if uint32(__t_tag_4) == 3772422949 {
					var __t6 uint32
					{
						var __t_tag_5 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_5
						if uint32(__t_tag_5) == 3772422949 {
							__t6 = 902936544
							goto end_branch_6
						} else {

						}
					}
					{
						__t6 = 1527465420
					}
				end_branch_6:
					__t11 = __t6
					goto end_branch_11
				} else {

				}
			}
			{
				var __t_tag_7 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_7
				if uint32(__t_tag_7) == 3772422949 {
					__t11 = 380165415
					goto end_branch_11
				} else {

				}
			}
			{
				var __t_tag_8 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_8
				var __t_and_10 bool = false
				if uint32(__t_tag_8) == 3944123360 {

					var __t_tag_9 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_9
					__t_and_10 = (uint32(__t_tag_9) == 3944123360)
				}
				if __t_and_10 {
					__t11 = 902936544
					goto end_branch_11
				} else {

				}
			}
			{
				__t11 = func() uint32 { panic("Failed pattern match") }()
			}
		end_branch_11:
			return gopurs_runtime.Value{Type: 9, IntVal: int64(__t11), UnsafePtr: nil}
		})})))}
	})
	return cache_Main_ordColor
}

var cache_Main_eqBox gopurs_runtime.Value
var once_Main_eqBox sync.Once

func Get_Main_eqBox() gopurs_runtime.Value {
	once_Main_eqBox.Do(func() {
		cache_Main_eqBox = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_eqBox(dictEq_0_box)
		})
	})
	return cache_Main_eqBox
}

var cache_Main_eqBox1 gopurs_runtime.Value
var once_Main_eqBox1 sync.Once

func Get_Main_eqBox1() gopurs_runtime.Value {
	once_Main_eqBox1.Do(func() {
		cache_Main_eqBox1 = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2545543625_3790796878(Rebox_Main_3790796878_2545543625(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Main_eqBox(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))})))))}
	})
	return cache_Main_eqBox1
}

var cache_Main_bifunctorEither2 gopurs_runtime.Value
var once_Main_bifunctorEither2 sync.Once

func Get_Main_bifunctorEither2() gopurs_runtime.Value {
	once_Main_bifunctorEither2.Do(func() {
		cache_Main_bifunctorEither2 = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 gopurs_runtime.Value
			{
				if m_2.Type == 9 && m_2.IntVal == 3583471031 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3583471031, UnsafePtr: unsafe.Pointer((&Constructor_Main_Left2[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Main_Left2[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0)}))}
					goto end_branch_0
				} else {

				}
			}
			{
				if m_2.Type == 9 && m_2.IntVal == 4276485804 {
					__t0 = gopurs_runtime.Value{Type: 9, IntVal: 4276485804, UnsafePtr: unsafe.Pointer((&Constructor_Main_Right2[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(g_1, (*Constructor_Main_Right2[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0)}))}
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
	return cache_Main_bifunctorEither2
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=(ADT ["Effect","Effect"] [Unit])
			__local_var_0_0 := gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = __local_var_2_2
			__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = __local_var_3_3
			__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((gopurs_runtime.Str((Data_Show_ShowIntImpl(gopurs_runtime.Int(int64(1)).IntVal))+(gopurs_runtime.Apply3(Rebox_Main_1680800814_464825424(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Main_foldableList())).V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Main_1950344881_1201789390(Rebox_Main_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}, Get_Data_Show_showIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 322902991, UnsafePtr: unsafe.Pointer(Rebox_Main_2737593216_176455803((&Constructor_Main_Cons[int64]{1, int64(2), (*Constructor_Main_Cons[int64])(nil)})))}).StrVal())).StrVal()) == ("12"))), gopurs_runtime.Value{})
			_ = __local_var_4_4
			__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqEither2(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 3583471031, UnsafePtr: unsafe.Pointer((&Constructor_Main_Left2[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(int64(4))}))}, gopurs_runtime.Value{Type: 9, IntVal: 3583471031, UnsafePtr: unsafe.Pointer((&Constructor_Main_Left2[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(int64(4))}))}).IntVal) != (0))), gopurs_runtime.Value{})
			_ = __local_var_5_5
			__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Main_eqBox(gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_1053099733_3790796878(Rebox_Main_3790796878_1053099733(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqInt()))))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 1145268909, UnsafePtr: unsafe.Pointer((&Constructor_Main_Full[gopurs_runtime.Value]{1, gopurs_runtime.Int(int64(2))}))}, gopurs_runtime.Value{Type: 9, IntVal: 1145268909, UnsafePtr: unsafe.Pointer((&Constructor_Main_Full[gopurs_runtime.Value]{1, gopurs_runtime.Int(int64(2))}))}).IntVal) != (0))), gopurs_runtime.Value{})
			_ = __local_var_6_6
			__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = __local_var_7_7
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

type Constructor_Main_Pair[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 T_a
}

type Constructor_Main_Nil[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Main_Cons[T_a]
}

type Constructor_Main_Left2[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
}

type Constructor_Main_Right2[T_a any, T_b any] struct {
	Rc uint32
	V0 T_b
}

type Constructor_Main_North struct {
	Rc uint32
}

type Constructor_Main_South struct {
	Rc uint32
}

type Constructor_Main_East struct {
	Rc uint32
}

type Constructor_Main_West struct {
	Rc uint32
}

type Constructor_Main_Red struct {
	Rc uint32
}

type Constructor_Main_Green struct {
	Rc uint32
}

type Constructor_Main_Blue struct {
	Rc uint32
}

type Constructor_Main_Empty[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Full[T_a any] struct {
	Rc uint32
	V0 T_a
}

func Call_Main_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Wrapper(x_0_loop string) string {
	var x_0 string = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Pair__1394429419(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop int64) *Constructor_Main_Pair[int64] {
Pair__1394429419:
	for {
		if false {
			continue Pair__1394429419
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 int64 = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Main_Pair[int64]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Main_Name(x_0_loop string) string {
	var x_0 string = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_Cons__41240261(__eta_norm_1_0_loop int64, __eta_norm_0_unused_1_loop *Constructor_Main_Cons[int64]) *Constructor_Main_Cons[int64] {
Cons__41240261:
	for {
		if false {
			continue Cons__41240261
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 *Constructor_Main_Cons[int64] = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		return (&Constructor_Main_Cons[int64]{1, __eta_norm_1_0, (*Constructor_Main_Cons[int64])(nil)})
	}
}

func Call_Main_Cons__2678532228(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop *Constructor_Main_Cons[int64]) *Constructor_Main_Cons[int64] {
Cons__2678532228:
	for {
		if false {
			continue Cons__2678532228
		}
		var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 *Constructor_Main_Cons[int64] = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return (&Constructor_Main_Cons[int64]{1, __eta_norm_1_0, __eta_norm_0_1})
	}
}

func Call_Main_Left2__1237640998(__eta_norm_0_0_loop int64) gopurs_runtime.Value {
Left2__1237640998:
	for {
		if false {
			continue Left2__1237640998
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 3583471031, UnsafePtr: unsafe.Pointer(Rebox_Main_4137439973_3184105157((&Constructor_Main_Left2[int64, int64]{1, __eta_norm_0_0})))}
	}
}

func Call_Main_Full__3180329226(__eta_norm_0_0_loop int64) *Constructor_Main_Full[int64] {
Full__3180329226:
	for {
		if false {
			continue Full__3180329226
		}
		var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return (&Constructor_Main_Full[int64]{1, __eta_norm_0_0})
	}
}

func Call_Main_eqEither2(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 bool
		{
			if x_1.Type == 9 && x_1.IntVal == 3583471031 {
				__t0 = (y_2.Type == 9 && y_2.IntVal == 3583471031) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Main_Left2[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Main_Left2[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0))
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = (x_1.Type == 9 && x_1.IntVal == 4276485804) && ((y_2.Type == 9 && y_2.IntVal == 4276485804) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Main_Right2[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Main_Right2[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0)))
		}
	end_branch_0:
		return gopurs_runtime.Bool(__t0)
	})}))}
}

func Call_Main_eqBox(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
	_ = dictEq_0
	return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_966678290_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Main_Full[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 bool
		{
			if x_1.Type == 9 && x_1.IntVal == 1145268909 && x_1.UnsafePtr == nil {
				__t0 = (y_2.Type == 9 && y_2.IntVal == 1145268909 && y_2.UnsafePtr == nil)
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = (x_1.Type == 9 && x_1.IntVal == 1145268909 && x_1.UnsafePtr != nil) && ((y_2.Type == 9 && y_2.IntVal == 1145268909 && y_2.UnsafePtr != nil) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Main_Full[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Main_Full[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0)))
		}
	end_branch_0:
		return gopurs_runtime.Bool(__t0)
	})})))}
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

func Rebox_Main_1201789390_1950344881(in *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) *Constructor_Data_Monoid_Monoid[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[string]{}
	out.V0 = in.V0
	out.V1 = in.V1.StrVal()
	return out
}

func Rebox_Main_1259967060_2687169876(in *Constructor_Data_Generic_Rep_Inr[uint32, uint32]) *Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Main_1557125915_791404512(in *Constructor_Main_Pair[int64]) *Constructor_Main_Pair[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Pair[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Main_1680800814_464825424(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Main_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[*Constructor_Main_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Main_176455803_2737593216(in *Constructor_Main_Cons[gopurs_runtime.Value]) *Constructor_Main_Cons[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Cons[int64]{}
	out.V0 = in.V0.IntVal
	out.V1 = Rebox_Main_176455803_2737593216(in.V1)
	return out
}

func Rebox_Main_1950344881_1201789390(in *Constructor_Data_Monoid_Monoid[string]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Str(in.V1)
	return out
}

func Rebox_Main_2199435624_385277032(in *Constructor_Data_Newtype_Newtype[string, string]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2362652146_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Main_Full[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2545543625_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Main_Full[int64]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_2700107938_1539742553(in *Constructor_Main_Full[int64]) *Constructor_Main_Full[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Full[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Main_2737593216_176455803(in *Constructor_Main_Cons[int64]) *Constructor_Main_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = Rebox_Main_2737593216_176455803(in.V1)
	return out
}

func Rebox_Main_2812149806_2834964816(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Main_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Main_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Main_2834964816_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Main_Cons[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3043886126_4169608016(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Main_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[*Constructor_Main_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Main_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
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

func Rebox_Main_3790796878_2545543625(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Main_Full[int64]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[*Constructor_Main_Full[int64]])(unsafe.Pointer(in))
}

func Rebox_Main_3790796878_3768443459(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[uint32])(unsafe.Pointer(in))
}

func Rebox_Main_4137439973_3184105157(in *Constructor_Main_Left2[int64, int64]) *Constructor_Main_Left2[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_Left2[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Main_4169608016_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Main_Cons[gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Traversable_Traversable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_4191096074_1323331594(in *Constructor_Data_Generic_Rep_Inl[uint32, uint32]) *Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Main_451755851_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Main_Pair[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_464825424_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Main_Cons[gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_497217223_1323331594(in *Constructor_Data_Generic_Rep_Inl[uint32, gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Main_660272253_2818661616(in *Constructor_Data_Generic_Rep_Generic[uint32, gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_660678937_2687169876(in *Constructor_Data_Generic_Rep_Inr[uint32, gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Main_966678290_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Main_Full[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}
