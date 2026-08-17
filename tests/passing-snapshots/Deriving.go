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
			return Call_Main_Z(x_0_box)
		})
	})
	return cache_Main_Z
}

var cache_Main_eqX gopurs_runtime.Value
var once_Main_eqX sync.Once

func Get_Main_eqX() gopurs_runtime.Value {
	once_Main_eqX.Do(func() {
		cache_Main_eqX = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
			})
		})}))}
	})
	return cache_Main_eqX
}

var cache_Main_eqZ gopurs_runtime.Value
var once_Main_eqZ sync.Once

func Get_Main_eqZ() gopurs_runtime.Value {
	once_Main_eqZ.Do(func() {
		cache_Main_eqZ = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				var __t5 bool
				{
					var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.RecordGet(x_0, "left")
					if __t_tag_3.Type == 9 && __t_tag_3.IntVal == 1409933510 {
						var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.RecordGet(y_1, "left")
						__t5 = (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1409933510) && (((*Constructor_Main_X)(gopurs_runtime.RecordGet(x_0, "left").UnsafePtr).V0) == ((*Constructor_Main_X)(gopurs_runtime.RecordGet(y_1, "left").UnsafePtr).V0))
						goto end_branch_5
					} else {

					}
				}
				{
					var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.RecordGet(x_0, "left")
					var __t_and_2 bool = false
					if __t_tag_0.Type == 9 && __t_tag_0.IntVal == 1682951303 {

						var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.RecordGet(y_1, "left")
						__t_and_2 = (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1682951303) && (((*Constructor_Main_Y)(gopurs_runtime.RecordGet(x_0, "left").UnsafePtr).V0.StrVal()) == ((*Constructor_Main_Y)(gopurs_runtime.RecordGet(y_1, "left").UnsafePtr).V0.StrVal()))
					}
					__t5 = __t_and_2
				}
			end_branch_5:
				var __t_and_12 bool = false
				if __t5 {

					var __t11 bool
					{
						var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.RecordGet(x_0, "right")
						if __t_tag_9.Type == 9 && __t_tag_9.IntVal == 1409933510 {
							var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.RecordGet(y_1, "right")
							__t11 = (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 1409933510) && (((*Constructor_Main_X)(gopurs_runtime.RecordGet(x_0, "right").UnsafePtr).V0) == ((*Constructor_Main_X)(gopurs_runtime.RecordGet(y_1, "right").UnsafePtr).V0))
							goto end_branch_11
						} else {

						}
					}
					{
						var __t_tag_6 gopurs_runtime.Value = gopurs_runtime.RecordGet(x_0, "right")
						var __t_and_8 bool = false
						if __t_tag_6.Type == 9 && __t_tag_6.IntVal == 1682951303 {

							var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.RecordGet(y_1, "right")
							__t_and_8 = (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 1682951303) && (((*Constructor_Main_Y)(gopurs_runtime.RecordGet(x_0, "right").UnsafePtr).V0.StrVal()) == ((*Constructor_Main_Y)(gopurs_runtime.RecordGet(y_1, "right").UnsafePtr).V0.StrVal()))
						}
						__t11 = __t_and_8
					}
				end_branch_11:
					__t_and_12 = __t11
				}
				return gopurs_runtime.Bool(__t_and_12)
			})
		})}))}
	})
	return cache_Main_eqZ
}

var cache_Main_ordX gopurs_runtime.Value
var once_Main_ordX sync.Once

func Get_Main_ordX() gopurs_runtime.Value {
	once_Main_ordX.Do(func() {
		cache_Main_ordX = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqX()))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
						__t1 = uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str((*Constructor_Main_Y)(x_0.UnsafePtr).V0.StrVal()), gopurs_runtime.Str((*Constructor_Main_Y)(y_1.UnsafePtr).V0.StrVal())).IntVal)
						goto end_branch_1
					} else {

					}
				}
				{
					__t1 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
				}
			end_branch_1:
				return gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
			})
		})}))}
	})
	return cache_Main_ordX
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true))
			_ = __local_var_0_0
			_dollar___unused_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = _dollar___unused_1_1
			_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_2_2
			_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_3_3
			_dollar___unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_4_4
			_dollar___unused_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_5_5
			_dollar___unused_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_6_6
			var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str("Bar"), gopurs_runtime.Str("Baz"))
			_dollar___unused_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool((uint32(__t_tag_8.IntVal) == 1527465420))), gopurs_runtime.Value{})
			_ = _dollar___unused_7_7
			_dollar___unused_8_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_Assert_assert(), gopurs_runtime.Bool(true)), gopurs_runtime.Value{})
			_ = _dollar___unused_8_9
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done")), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_eqV gopurs_runtime.Value
var once_Main_eqV sync.Once

func Get_Main_eqV() gopurs_runtime.Value {
	once_Main_eqV.Do(func() {
		cache_Main_eqV = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(false)
			})
		})}))}
	})
	return cache_Main_eqV
}

var cache_Main_ordV gopurs_runtime.Value
var once_Main_ordV sync.Once

func Get_Main_ordV() gopurs_runtime.Value {
	once_Main_ordV.Do(func() {
		cache_Main_ordV = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Main_eqV()))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
			})
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

func Call_Main_Z(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
