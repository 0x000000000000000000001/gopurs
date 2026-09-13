package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_X gopurs_runtime.Value
var once_Main_X sync.Once

func Get_Main_X() gopurs_runtime.Value {
	once_Main_X.Do(func() {
		cache_Main_X = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((&Constructor_Main_X{1, value0.IntVal}))}
		})
	})
	return cache_Main_X
}

var cache_Main_Y gopurs_runtime.Value
var once_Main_Y sync.Once

func Get_Main_Y() gopurs_runtime.Value {
	once_Main_Y.Do(func() {
		cache_Main_Y = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1682951303, UnsafePtr: unsafe.Pointer((&Constructor_Main_Y{1, value0}))}
		})
	})
	return cache_Main_Y
}

var cache_Main_Z gopurs_runtime.Value
var once_Main_Z sync.Once

func Get_Main_Z() gopurs_runtime.Value {
	once_Main_Z.Do(func() {
		cache_Main_Z = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				orig := Call_Main_Z(func() struct {
					left  gopurs_runtime.Value
					right gopurs_runtime.Value
				} {
					orig := x_0_box
					_ = orig
					clone := struct {
						left  gopurs_runtime.Value
						right gopurs_runtime.Value
					}{}
					clone.left = gopurs_runtime.RecordGet(orig, "left")
					clone.right = gopurs_runtime.RecordGet(orig, "right")
					return clone
				}())
				_ = orig
				return gopurs_runtime.RecordDict2("left", "right", orig.left, orig.right)
			}()
		})
	})
	return cache_Main_Z
}

var cache_Main_eqX gopurs_runtime.Value
var once_Main_eqX sync.Once

func Get_Main_eqX() gopurs_runtime.Value {
	once_Main_eqX.Do(func() {
		cache_Main_eqX = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t0 bool
			{
				if x_0.Type == 9 && x_0.IntVal == 1409933510 {
					__t0 = (y_1.Type == 9 && y_1.IntVal == 1409933510) && (((*Constructor_Main_X)(x_0.UnsafePtr).V0) == ((*Constructor_Main_X)(y_1.UnsafePtr).V0))
					goto end_branch_0
				} else {

				}
			}
			{
				__t0 = (x_0.Type == 9 && x_0.IntVal == 1682951303) && ((y_1.Type == 9 && y_1.IntVal == 1682951303) && (((*Constructor_Main_Y)(x_0.UnsafePtr).V0.StrVal()) == ((*Constructor_Main_Y)(y_1.UnsafePtr).V0.StrVal())))
			}
		end_branch_0:
			return gopurs_runtime.Bool(__t0)
		})}))}
	})
	return cache_Main_eqX
}

var cache_Main_eqZ gopurs_runtime.Value
var once_Main_eqZ sync.Once

func Get_Main_eqZ() gopurs_runtime.Value {
	once_Main_eqZ.Do(func() {
		cache_Main_eqZ = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Main_2923583716_3790796878((&Constructor_Data_Eq_Eq[struct {
			left  gopurs_runtime.Value
			right gopurs_runtime.Value
		}]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t19 bool
			{
				var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.RecordGet(x_0, "left")
				_ = __t_tag_10
				if __t_tag_10.Type == 9 && __t_tag_10.IntVal == 1409933510 {
					var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.RecordGet(y_1, "left")
					_ = __t_tag_11
					var __t_and_18 bool = false
					if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 1409933510) && (((*Constructor_Main_X)(gopurs_runtime.RecordGet(x_0, "left").UnsafePtr).V0) == ((*Constructor_Main_X)(gopurs_runtime.RecordGet(y_1, "left").UnsafePtr).V0)) {

						var __t17 bool
						{
							var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.RecordGet(x_0, "right")
							_ = __t_tag_15
							if __t_tag_15.Type == 9 && __t_tag_15.IntVal == 1409933510 {
								var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.RecordGet(y_1, "right")
								_ = __t_tag_16
								__t17 = (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 1409933510) && (((*Constructor_Main_X)(gopurs_runtime.RecordGet(x_0, "right").UnsafePtr).V0) == ((*Constructor_Main_X)(gopurs_runtime.RecordGet(y_1, "right").UnsafePtr).V0))
								goto end_branch_17
							} else {

							}
						}
						{
							var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.RecordGet(x_0, "right")
							_ = __t_tag_12
							var __t_and_14 bool = false
							if __t_tag_12.Type == 9 && __t_tag_12.IntVal == 1682951303 {

								var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.RecordGet(y_1, "right")
								_ = __t_tag_13
								__t_and_14 = (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 1682951303) && (((*Constructor_Main_Y)(gopurs_runtime.RecordGet(x_0, "right").UnsafePtr).V0.StrVal()) == ((*Constructor_Main_Y)(gopurs_runtime.RecordGet(y_1, "right").UnsafePtr).V0.StrVal()))
							}
							__t17 = __t_and_14
						}
					end_branch_17:
						__t_and_18 = __t17
					}
					__t19 = __t_and_18
					goto end_branch_19
				} else {

				}
			}
			{
				var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.RecordGet(x_0, "left")
				_ = __t_tag_0
				var __t_and_2 bool = false
				if __t_tag_0.Type == 9 && __t_tag_0.IntVal == 1682951303 {

					var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.RecordGet(y_1, "left")
					_ = __t_tag_1
					__t_and_2 = (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1682951303) && (((*Constructor_Main_Y)(gopurs_runtime.RecordGet(x_0, "left").UnsafePtr).V0.StrVal()) == ((*Constructor_Main_Y)(gopurs_runtime.RecordGet(y_1, "left").UnsafePtr).V0.StrVal()))
				}
				var __t_and_9 bool = false
				if __t_and_2 {

					var __t8 bool
					{
						var __t_tag_6 gopurs_runtime.Value = gopurs_runtime.RecordGet(x_0, "right")
						_ = __t_tag_6
						if __t_tag_6.Type == 9 && __t_tag_6.IntVal == 1409933510 {
							var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.RecordGet(y_1, "right")
							_ = __t_tag_7
							__t8 = (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 1409933510) && (((*Constructor_Main_X)(gopurs_runtime.RecordGet(x_0, "right").UnsafePtr).V0) == ((*Constructor_Main_X)(gopurs_runtime.RecordGet(y_1, "right").UnsafePtr).V0))
							goto end_branch_8
						} else {

						}
					}
					{
						var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.RecordGet(x_0, "right")
						_ = __t_tag_3
						var __t_and_5 bool = false
						if __t_tag_3.Type == 9 && __t_tag_3.IntVal == 1682951303 {

							var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.RecordGet(y_1, "right")
							_ = __t_tag_4
							__t_and_5 = (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1682951303) && (((*Constructor_Main_Y)(gopurs_runtime.RecordGet(x_0, "right").UnsafePtr).V0.StrVal()) == ((*Constructor_Main_Y)(gopurs_runtime.RecordGet(y_1, "right").UnsafePtr).V0.StrVal()))
						}
						__t8 = __t_and_5
					}
				end_branch_8:
					__t_and_9 = __t8
				}
				__t19 = __t_and_9
			}
		end_branch_19:
			return gopurs_runtime.Bool(__t19)
		})})))}
	})
	return cache_Main_eqZ
}

var cache_Main_ordX gopurs_runtime.Value
var once_Main_ordX sync.Once

func Get_Main_ordX() gopurs_runtime.Value {
	once_Main_ordX.Do(func() {
		cache_Main_ordX = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqX()))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t1 uint32
			{
				if x_0.Type == 9 && x_0.IntVal == 1409933510 {
					var __t0 uint32
					{
						if y_1.Type == 9 && y_1.IntVal == 1409933510 {
							__t0 = uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Main_X)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Main_X)(y_1.UnsafePtr).V0)).IntVal)
							goto end_branch_0
						} else {

						}
					}
					{
						__t0 = 1527465420
					}
				end_branch_0:
					__t1 = __t0
					goto end_branch_1
				} else {

				}
			}
			{
				if y_1.Type == 9 && y_1.IntVal == 1409933510 {
					__t1 = 380165415
					goto end_branch_1
				} else {

				}
			}
			{
				if (x_0.Type == 9 && x_0.IntVal == 1682951303) && (y_1.Type == 9 && y_1.IntVal == 1682951303) {
					__t1 = uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, (*Constructor_Main_Y)(x_0.UnsafePtr).V0, (*Constructor_Main_Y)(y_1.UnsafePtr).V0).IntVal)
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = func() uint32 { panic("Failed pattern match") }()
			}
		end_branch_1:
			return gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
		})}))}
	})
	return cache_Main_ordX
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
								var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str("Bar"), gopurs_runtime.Str("Baz"))
								_ = __t_tag_0
								return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_0.IntVal) == 1527465420))), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
									}))
								}))
							}))
						}))
					}))
				}))
			}))
		}))
	})
	return cache_Main_main
}

var cache_Main_eqV gopurs_runtime.Value
var once_Main_eqV sync.Once

func Get_Main_eqV() gopurs_runtime.Value {
	once_Main_eqV.Do(func() {
		cache_Main_eqV = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(false)
		})}))}
	})
	return cache_Main_eqV
}

var cache_Main_ordV gopurs_runtime.Value
var once_Main_ordV sync.Once

func Get_Main_ordV() gopurs_runtime.Value {
	once_Main_ordV.Do(func() {
		cache_Main_ordV = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Main_eqV()))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
		})}))}
	})
	return cache_Main_ordV
}

type Constructor_Main_X struct {
	Rc uint32
	V0 int64
}

type Constructor_Main_Y struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_Z(x_0_loop struct {
	left  gopurs_runtime.Value
	right gopurs_runtime.Value
}) struct {
	left  gopurs_runtime.Value
	right gopurs_runtime.Value
} {
	var x_0 struct {
		left  gopurs_runtime.Value
		right gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return x_0
}

func Rebox_Main_2923583716_3790796878(in *Constructor_Data_Eq_Eq[struct {
	left  gopurs_runtime.Value
	right gopurs_runtime.Value
}]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
