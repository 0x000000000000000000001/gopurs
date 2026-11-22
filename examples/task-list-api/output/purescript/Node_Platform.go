package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Node_Platform_AIX gopurs_runtime.Value
var once_Node_Platform_AIX sync.Once

func Get_Node_Platform_AIX() gopurs_runtime.Value {
	once_Node_Platform_AIX.Do(func() {
		cache_Node_Platform_AIX = gopurs_runtime.Value{Type: 9, IntVal: int64(700987045), UnsafePtr: nil}
	})
	return cache_Node_Platform_AIX
}

var cache_Node_Platform_Darwin gopurs_runtime.Value
var once_Node_Platform_Darwin sync.Once

func Get_Node_Platform_Darwin() gopurs_runtime.Value {
	once_Node_Platform_Darwin.Do(func() {
		cache_Node_Platform_Darwin = gopurs_runtime.Value{Type: 9, IntVal: int64(207948338), UnsafePtr: nil}
	})
	return cache_Node_Platform_Darwin
}

var cache_Node_Platform_FreeBSD gopurs_runtime.Value
var once_Node_Platform_FreeBSD sync.Once

func Get_Node_Platform_FreeBSD() gopurs_runtime.Value {
	once_Node_Platform_FreeBSD.Do(func() {
		cache_Node_Platform_FreeBSD = gopurs_runtime.Value{Type: 9, IntVal: int64(3079948500), UnsafePtr: nil}
	})
	return cache_Node_Platform_FreeBSD
}

var cache_Node_Platform_Linux gopurs_runtime.Value
var once_Node_Platform_Linux sync.Once

func Get_Node_Platform_Linux() gopurs_runtime.Value {
	once_Node_Platform_Linux.Do(func() {
		cache_Node_Platform_Linux = gopurs_runtime.Value{Type: 9, IntVal: int64(4148875795), UnsafePtr: nil}
	})
	return cache_Node_Platform_Linux
}

var cache_Node_Platform_OpenBSD gopurs_runtime.Value
var once_Node_Platform_OpenBSD sync.Once

func Get_Node_Platform_OpenBSD() gopurs_runtime.Value {
	once_Node_Platform_OpenBSD.Do(func() {
		cache_Node_Platform_OpenBSD = gopurs_runtime.Value{Type: 9, IntVal: int64(508021620), UnsafePtr: nil}
	})
	return cache_Node_Platform_OpenBSD
}

var cache_Node_Platform_SunOS gopurs_runtime.Value
var once_Node_Platform_SunOS sync.Once

func Get_Node_Platform_SunOS() gopurs_runtime.Value {
	once_Node_Platform_SunOS.Do(func() {
		cache_Node_Platform_SunOS = gopurs_runtime.Value{Type: 9, IntVal: int64(2031616513), UnsafePtr: nil}
	})
	return cache_Node_Platform_SunOS
}

var cache_Node_Platform_Win32 gopurs_runtime.Value
var once_Node_Platform_Win32 sync.Once

func Get_Node_Platform_Win32() gopurs_runtime.Value {
	once_Node_Platform_Win32.Do(func() {
		cache_Node_Platform_Win32 = gopurs_runtime.Value{Type: 9, IntVal: int64(569341156), UnsafePtr: nil}
	})
	return cache_Node_Platform_Win32
}

var cache_Node_Platform_Android gopurs_runtime.Value
var once_Node_Platform_Android sync.Once

func Get_Node_Platform_Android() gopurs_runtime.Value {
	once_Node_Platform_Android.Do(func() {
		cache_Node_Platform_Android = gopurs_runtime.Value{Type: 9, IntVal: int64(2339342990), UnsafePtr: nil}
	})
	return cache_Node_Platform_Android
}

var cache_Node_Platform_toString gopurs_runtime.Value
var once_Node_Platform_toString sync.Once

func Get_Node_Platform_toString() gopurs_runtime.Value {
	once_Node_Platform_toString.Do(func() {
		cache_Node_Platform_toString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Node_Platform_toString(uint32(v_0_box.IntVal)))
		})
	})
	return cache_Node_Platform_toString
}

var cache_Node_Platform_showPlatform gopurs_runtime.Value
var once_Node_Platform_showPlatform sync.Once

func Get_Node_Platform_showPlatform() gopurs_runtime.Value {
	once_Node_Platform_showPlatform.Do(func() {
		cache_Node_Platform_showPlatform = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Node_Platform_1209612131_1386611502((&Constructor_Data_Show_Show[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t8 gopurs_runtime.Value
			{
				var __t_tag_0 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_0
				if uint32(__t_tag_0) == 700987045 {
					__t8 = gopurs_runtime.StrWithCodeUnits("AIX", 3)
					goto end_branch_8
				} else {

				}
			}
			{
				var __t_tag_1 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_1
				if uint32(__t_tag_1) == 207948338 {
					__t8 = gopurs_runtime.StrWithCodeUnits("Darwin", 6)
					goto end_branch_8
				} else {

				}
			}
			{
				var __t_tag_2 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_2
				if uint32(__t_tag_2) == 3079948500 {
					__t8 = gopurs_runtime.StrWithCodeUnits("FreeBSD", 7)
					goto end_branch_8
				} else {

				}
			}
			{
				var __t_tag_3 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 4148875795 {
					__t8 = gopurs_runtime.StrWithCodeUnits("Linux", 5)
					goto end_branch_8
				} else {

				}
			}
			{
				var __t_tag_4 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_4
				if uint32(__t_tag_4) == 508021620 {
					__t8 = gopurs_runtime.StrWithCodeUnits("OpenBSD", 7)
					goto end_branch_8
				} else {

				}
			}
			{
				var __t_tag_5 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_5
				if uint32(__t_tag_5) == 2031616513 {
					__t8 = gopurs_runtime.StrWithCodeUnits("SunOS", 5)
					goto end_branch_8
				} else {

				}
			}
			{
				var __t_tag_6 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_6
				if uint32(__t_tag_6) == 569341156 {
					__t8 = gopurs_runtime.StrWithCodeUnits("Win32", 5)
					goto end_branch_8
				} else {

				}
			}
			{
				var __t_tag_7 uint32 = uint32(v_0.IntVal)
				_ = __t_tag_7
				if uint32(__t_tag_7) == 2339342990 {
					__t8 = gopurs_runtime.StrWithCodeUnits("Android", 7)
					goto end_branch_8
				} else {

				}
			}
			{
				__t8 = gopurs_runtime.Str(func() string { panic("Failed pattern match") }())
			}
		end_branch_8:
			return __t8
		})})))}
	})
	return cache_Node_Platform_showPlatform
}

var cache_Node_Platform_fromString gopurs_runtime.Value
var once_Node_Platform_fromString sync.Once

func Get_Node_Platform_fromString() gopurs_runtime.Value {
	once_Node_Platform_fromString.Do(func() {
		cache_Node_Platform_fromString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Node_Platform_fromString(v_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Node_Platform_fromString
}

var cache_Node_Platform_eqPlatform gopurs_runtime.Value
var once_Node_Platform_eqPlatform sync.Once

func Get_Node_Platform_eqPlatform() gopurs_runtime.Value {
	once_Node_Platform_eqPlatform.Do(func() {
		cache_Node_Platform_eqPlatform = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Node_Platform_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t17 bool
			{
				var __t_tag_3 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 700987045 {
					var __t_tag_4 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_4
					__t17 = (uint32(__t_tag_4) == 700987045)
					goto end_branch_17
				} else {

				}
			}
			{
				var __t_tag_5 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_5
				if uint32(__t_tag_5) == 207948338 {
					var __t_tag_6 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_6
					__t17 = (uint32(__t_tag_6) == 207948338)
					goto end_branch_17
				} else {

				}
			}
			{
				var __t_tag_7 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_7
				if uint32(__t_tag_7) == 3079948500 {
					var __t_tag_8 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_8
					__t17 = (uint32(__t_tag_8) == 3079948500)
					goto end_branch_17
				} else {

				}
			}
			{
				var __t_tag_9 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_9
				if uint32(__t_tag_9) == 4148875795 {
					var __t_tag_10 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_10
					__t17 = (uint32(__t_tag_10) == 4148875795)
					goto end_branch_17
				} else {

				}
			}
			{
				var __t_tag_11 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_11
				if uint32(__t_tag_11) == 508021620 {
					var __t_tag_12 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_12
					__t17 = (uint32(__t_tag_12) == 508021620)
					goto end_branch_17
				} else {

				}
			}
			{
				var __t_tag_13 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_13
				if uint32(__t_tag_13) == 2031616513 {
					var __t_tag_14 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_14
					__t17 = (uint32(__t_tag_14) == 2031616513)
					goto end_branch_17
				} else {

				}
			}
			{
				var __t_tag_15 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_15
				if uint32(__t_tag_15) == 569341156 {
					var __t_tag_16 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_16
					__t17 = (uint32(__t_tag_16) == 569341156)
					goto end_branch_17
				} else {

				}
			}
			{
				var __t_tag_0 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_0
				var __t_and_2 bool = false
				if uint32(__t_tag_0) == 2339342990 {

					var __t_tag_1 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_1
					__t_and_2 = (uint32(__t_tag_1) == 2339342990)
				}
				__t17 = __t_and_2
			}
		end_branch_17:
			return gopurs_runtime.Bool(__t17)
		})})))}
	})
	return cache_Node_Platform_eqPlatform
}

var cache_Node_Platform_ordPlatform gopurs_runtime.Value
var once_Node_Platform_ordPlatform sync.Once

func Get_Node_Platform_ordPlatform() gopurs_runtime.Value {
	once_Node_Platform_ordPlatform.Do(func() {
		cache_Node_Platform_ordPlatform = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Node_Platform_3730953251_4177771502((&Constructor_Data_Ord_Ord[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Node_Platform_3768443459_3790796878(Rebox_Node_Platform_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Node_Platform_eqPlatform()))))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t31 gopurs_runtime.Value
			{
				var __t_tag_0 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_0
				if uint32(__t_tag_0) == 700987045 {
					var __t2 gopurs_runtime.Value
					{
						var __t_tag_1 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_1
						if uint32(__t_tag_1) == 700987045 {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_2:
					__t31 = __t2
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_3 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 700987045 {
					__t31 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_4 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_4
				if uint32(__t_tag_4) == 207948338 {
					var __t6 gopurs_runtime.Value
					{
						var __t_tag_5 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_5
						if uint32(__t_tag_5) == 207948338 {
							__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_6
						} else {

						}
					}
					{
						__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_6:
					__t31 = __t6
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_7 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_7
				if uint32(__t_tag_7) == 207948338 {
					__t31 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_8 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_8
				if uint32(__t_tag_8) == 3079948500 {
					var __t10 gopurs_runtime.Value
					{
						var __t_tag_9 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_9
						if uint32(__t_tag_9) == 3079948500 {
							__t10 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_10
						} else {

						}
					}
					{
						__t10 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_10:
					__t31 = __t10
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_11 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_11
				if uint32(__t_tag_11) == 3079948500 {
					__t31 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_12 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_12
				if uint32(__t_tag_12) == 4148875795 {
					var __t14 gopurs_runtime.Value
					{
						var __t_tag_13 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_13
						if uint32(__t_tag_13) == 4148875795 {
							__t14 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_14
						} else {

						}
					}
					{
						__t14 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_14:
					__t31 = __t14
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_15 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_15
				if uint32(__t_tag_15) == 4148875795 {
					__t31 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_16 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_16
				if uint32(__t_tag_16) == 508021620 {
					var __t18 gopurs_runtime.Value
					{
						var __t_tag_17 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_17
						if uint32(__t_tag_17) == 508021620 {
							__t18 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_18
						} else {

						}
					}
					{
						__t18 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_18:
					__t31 = __t18
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_19 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_19
				if uint32(__t_tag_19) == 508021620 {
					__t31 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_20 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_20
				if uint32(__t_tag_20) == 2031616513 {
					var __t22 gopurs_runtime.Value
					{
						var __t_tag_21 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_21
						if uint32(__t_tag_21) == 2031616513 {
							__t22 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_22
						} else {

						}
					}
					{
						__t22 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_22:
					__t31 = __t22
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_23 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_23
				if uint32(__t_tag_23) == 2031616513 {
					__t31 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_24 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_24
				if uint32(__t_tag_24) == 569341156 {
					var __t26 gopurs_runtime.Value
					{
						var __t_tag_25 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_25
						if uint32(__t_tag_25) == 569341156 {
							__t26 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_26
						} else {

						}
					}
					{
						__t26 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_26:
					__t31 = __t26
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_27 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_27
				if uint32(__t_tag_27) == 569341156 {
					__t31 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_31
				} else {

				}
			}
			{
				var __t_tag_28 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_28
				var __t_and_30 bool = false
				if uint32(__t_tag_28) == 2339342990 {

					var __t_tag_29 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_29
					__t_and_30 = (uint32(__t_tag_29) == 2339342990)
				}
				if __t_and_30 {
					__t31 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
					goto end_branch_31
				} else {

				}
			}
			{
				__t31 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
			}
		end_branch_31:
			return __t31
		})})))}
	})
	return cache_Node_Platform_ordPlatform
}

type Constructor_Node_Platform_AIX struct {
	Rc uint32
}

type Constructor_Node_Platform_Darwin struct {
	Rc uint32
}

type Constructor_Node_Platform_FreeBSD struct {
	Rc uint32
}

type Constructor_Node_Platform_Linux struct {
	Rc uint32
}

type Constructor_Node_Platform_OpenBSD struct {
	Rc uint32
}

type Constructor_Node_Platform_SunOS struct {
	Rc uint32
}

type Constructor_Node_Platform_Win32 struct {
	Rc uint32
}

type Constructor_Node_Platform_Android struct {
	Rc uint32
}

func Call_Node_Platform_toString(v_0_loop uint32) string {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var __t0 string
	{
		if v_0 == 700987045 {
			__t0 = "aix"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 207948338 {
			__t0 = "darwin"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 3079948500 {
			__t0 = "freebsd"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 4148875795 {
			__t0 = "linux"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 508021620 {
			__t0 = "openbsd"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 2031616513 {
			__t0 = "sunos"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 569341156 {
			__t0 = "win32"
			goto end_branch_0
		} else {

		}
	}
	{
		if v_0 == 2339342990 {
			__t0 = "android"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() string { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Node_Platform_fromString(v_0_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var v_0 string = v_0_loop
	_ = v_0
	var __t0 *Constructor_Data_Maybe_Just[uint32]
	{
		if (v_0) == ("aix") {
			__t0 = Rebox_Node_Platform_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(700987045), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (v_0) == ("darwin") {
			__t0 = Rebox_Node_Platform_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(207948338), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (v_0) == ("freebsd") {
			__t0 = Rebox_Node_Platform_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(3079948500), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (v_0) == ("linux") {
			__t0 = Rebox_Node_Platform_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(4148875795), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (v_0) == ("openbsd") {
			__t0 = Rebox_Node_Platform_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(508021620), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (v_0) == ("sunos") {
			__t0 = Rebox_Node_Platform_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(2031616513), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (v_0) == ("win32") {
			__t0 = Rebox_Node_Platform_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(569341156), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (v_0) == ("android") {
			__t0 = Rebox_Node_Platform_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(2339342990), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = Rebox_Node_Platform_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Node_Platform_622082505_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Rebox_Node_Platform_1209612131_1386611502(in *Constructor_Data_Show_Show[uint32]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Node_Platform_3094389156_622082505(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[uint32] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[uint32]{}
	out.V0 = uint32(in.V0.IntVal)
	return out
}

func Rebox_Node_Platform_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Node_Platform_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Node_Platform_3790796878_3768443459(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[uint32])(unsafe.Pointer(in))
}

func Rebox_Node_Platform_622082505_3094389156(in *Constructor_Data_Maybe_Just[uint32]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}
