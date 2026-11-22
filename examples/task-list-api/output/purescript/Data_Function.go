package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Function_lessThanOrEq gopurs_runtime.Value
var once_Data_Function_lessThanOrEq sync.Once

func Get_Data_Function_lessThanOrEq() gopurs_runtime.Value {
	once_Data_Function_lessThanOrEq.Do(func() {
		cache_Data_Function_lessThanOrEq = Get_Data_Ord_lessThanOrEq__3348776441()
	})
	return cache_Data_Function_lessThanOrEq
}

var cache_Data_Function_sub gopurs_runtime.Value
var once_Data_Function_sub sync.Once

func Get_Data_Function_sub() gopurs_runtime.Value {
	once_Data_Function_sub.Do(func() {
		cache_Data_Function_sub = Get_Data_Ring_sub__4157804846()
	})
	return cache_Data_Function_sub
}

var cache_Data_Function_on gopurs_runtime.Value
var once_Data_Function_on sync.Once

func Get_Data_Function_on() gopurs_runtime.Value {
	once_Data_Function_on.Do(func() {
		cache_Data_Function_on = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_on(f_0_box, g_1_box, x_2_box, y_3_box)
		})
	})
	return cache_Data_Function_on
}

var cache_Data_Function_flip gopurs_runtime.Value
var once_Data_Function_flip sync.Once

func Get_Data_Function_flip() gopurs_runtime.Value {
	once_Data_Function_flip.Do(func() {
		cache_Data_Function_flip = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_flip(f_0_box, b_1_box, a_2_box)
		})
	})
	return cache_Data_Function_flip
}

var cache_Data_Function_flip__1501157206 gopurs_runtime.Value
var once_Data_Function_flip__1501157206 sync.Once

func Get_Data_Function_flip__1501157206() gopurs_runtime.Value {
	once_Data_Function_flip__1501157206.Do(func() {
		cache_Data_Function_flip__1501157206 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__1501157206(f_unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](a_2_box)))}
		})
	})
	return cache_Data_Function_flip__1501157206
}

var cache_Data_Function_flip__2508043272 gopurs_runtime.Value
var once_Data_Function_flip__2508043272 sync.Once

func Get_Data_Function_flip__2508043272() gopurs_runtime.Value {
	once_Data_Function_flip__2508043272.Do(func() {
		cache_Data_Function_flip__2508043272 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_flip__2508043272(f_unused_0_box, b_1_box, a_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_flip__2508043272
}

var cache_Data_Function_flip__95728040 gopurs_runtime.Value
var once_Data_Function_flip__95728040 sync.Once

func Get_Data_Function_flip__95728040() gopurs_runtime.Value {
	once_Data_Function_flip__95728040.Do(func() {
		cache_Data_Function_flip__95728040 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_flip__95728040(f_unused_0_box, b_1_box, Rebox_Data_Function_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](a_2_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Function_flip__95728040
}

var cache_Data_Function_flip__1515853032 gopurs_runtime.Value
var once_Data_Function_flip__1515853032 sync.Once

func Get_Data_Function_flip__1515853032() gopurs_runtime.Value {
	once_Data_Function_flip__1515853032.Do(func() {
		cache_Data_Function_flip__1515853032 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_flip__1515853032(f_unused_0_box, b_1_box, Rebox_Data_Function_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](a_2_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Function_flip__1515853032
}

var cache_Data_Function_flip__2414803565 gopurs_runtime.Value
var once_Data_Function_flip__2414803565 sync.Once

func Get_Data_Function_flip__2414803565() gopurs_runtime.Value {
	once_Data_Function_flip__2414803565.Do(func() {
		cache_Data_Function_flip__2414803565 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_flip__2414803565(f_unused_0_box, b_1_box, a_2_box)
		})
	})
	return cache_Data_Function_flip__2414803565
}

var cache_Data_Function_flip__377817773 gopurs_runtime.Value
var once_Data_Function_flip__377817773 sync.Once

func Get_Data_Function_flip__377817773() gopurs_runtime.Value {
	once_Data_Function_flip__377817773.Do(func() {
		cache_Data_Function_flip__377817773 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_flip__377817773(f_unused_0_box, b_1_box, a_2_box)
		})
	})
	return cache_Data_Function_flip__377817773
}

var cache_Data_Function_flip__1753236854 gopurs_runtime.Value
var once_Data_Function_flip__1753236854 sync.Once

func Get_Data_Function_flip__1753236854() gopurs_runtime.Value {
	once_Data_Function_flip__1753236854.Do(func() {
		cache_Data_Function_flip__1753236854 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_flip__1753236854(f_unused_0_box, b_1_box, a_2_box)
		})
	})
	return cache_Data_Function_flip__1753236854
}

var cache_Data_Function_flip__379318895 gopurs_runtime.Value
var once_Data_Function_flip__379318895 sync.Once

func Get_Data_Function_flip__379318895() gopurs_runtime.Value {
	once_Data_Function_flip__379318895.Do(func() {
		cache_Data_Function_flip__379318895 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Function_flip__379318895(f_unused_0_box, Rebox_Data_Function_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_1_box)), a_2_box))}
		})
	})
	return cache_Data_Function_flip__379318895
}

var cache_Data_Function_flip__650029772 gopurs_runtime.Value
var once_Data_Function_flip__650029772 sync.Once

func Get_Data_Function_flip__650029772() gopurs_runtime.Value {
	once_Data_Function_flip__650029772.Do(func() {
		cache_Data_Function_flip__650029772 = gopurs_runtime.Func3(func(f_unused_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_flip__650029772(f_unused_0_box, b_1_box.IntVal, a_2_box)
		})
	})
	return cache_Data_Function_flip__650029772
}

var cache_Data_Function_flip__1483642665 gopurs_runtime.Value
var once_Data_Function_flip__1483642665 sync.Once

func Get_Data_Function_flip__1483642665() gopurs_runtime.Value {
	once_Data_Function_flip__1483642665.Do(func() {
		cache_Data_Function_flip__1483642665 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_flip__1483642665(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](b_1_box), a_2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Function_flip__1483642665
}

var cache_Data_Function_go__const gopurs_runtime.Value
var once_Data_Function_go__const sync.Once

func Get_Data_Function_go__const() gopurs_runtime.Value {
	once_Data_Function_go__const.Do(func() {
		cache_Data_Function_go__const = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_go__const(a_0_box, v_1_box)
		})
	})
	return cache_Data_Function_go__const
}

var cache_Data_Function_const__1740749855 gopurs_runtime.Value
var once_Data_Function_const__1740749855 sync.Once

func Get_Data_Function_const__1740749855() gopurs_runtime.Value {
	once_Data_Function_const__1740749855.Do(func() {
		cache_Data_Function_const__1740749855 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_const__1740749855(a_unused_0_box, v_1_box)
		})
	})
	return cache_Data_Function_const__1740749855
}

var cache_Data_Function_const__3605917656 gopurs_runtime.Value
var once_Data_Function_const__3605917656 sync.Once

func Get_Data_Function_const__3605917656() gopurs_runtime.Value {
	once_Data_Function_const__3605917656.Do(func() {
		cache_Data_Function_const__3605917656 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_const__3605917656(a_0_box, Rebox_Data_Function_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box)))
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_const__3605917656
}

var cache_Data_Function_const__1229331832 gopurs_runtime.Value
var once_Data_Function_const__1229331832 sync.Once

func Get_Data_Function_const__1229331832() gopurs_runtime.Value {
	once_Data_Function_const__1229331832.Do(func() {
		cache_Data_Function_const__1229331832 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_const__1229331832(a_0_box, Rebox_Data_Function_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box)))
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_const__1229331832
}

var cache_Data_Function_const__3126182883 gopurs_runtime.Value
var once_Data_Function_const__3126182883 sync.Once

func Get_Data_Function_const__3126182883() gopurs_runtime.Value {
	once_Data_Function_const__3126182883.Do(func() {
		cache_Data_Function_const__3126182883 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_const__3126182883(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](a_0_box), v_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Function_const__3126182883
}

var cache_Data_Function_const__4207166068 gopurs_runtime.Value
var once_Data_Function_const__4207166068 sync.Once

func Get_Data_Function_const__4207166068() gopurs_runtime.Value {
	once_Data_Function_const__4207166068.Do(func() {
		cache_Data_Function_const__4207166068 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_const__4207166068(a_unused_0_box, v_1_box)
		})
	})
	return cache_Data_Function_const__4207166068
}

var cache_Data_Function_const__320519811 gopurs_runtime.Value
var once_Data_Function_const__320519811 sync.Once

func Get_Data_Function_const__320519811() gopurs_runtime.Value {
	once_Data_Function_const__320519811.Do(func() {
		cache_Data_Function_const__320519811 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_const__320519811(a_0_box, v_1_box)
		})
	})
	return cache_Data_Function_const__320519811
}

var cache_Data_Function_const__2941505456 gopurs_runtime.Value
var once_Data_Function_const__2941505456 sync.Once

func Get_Data_Function_const__2941505456() gopurs_runtime.Value {
	once_Data_Function_const__2941505456.Do(func() {
		cache_Data_Function_const__2941505456 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_const__2941505456(a_unused_0_box, v_1_box)
		})
	})
	return cache_Data_Function_const__2941505456
}

var cache_Data_Function_const__2476124256 gopurs_runtime.Value
var once_Data_Function_const__2476124256 sync.Once

func Get_Data_Function_const__2476124256() gopurs_runtime.Value {
	once_Data_Function_const__2476124256.Do(func() {
		cache_Data_Function_const__2476124256 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_const__2476124256(a_unused_0_box, v_1_box)
		})
	})
	return cache_Data_Function_const__2476124256
}

var cache_Data_Function_const__3225710767 gopurs_runtime.Value
var once_Data_Function_const__3225710767 sync.Once

func Get_Data_Function_const__3225710767() gopurs_runtime.Value {
	once_Data_Function_const__3225710767.Do(func() {
		cache_Data_Function_const__3225710767 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Function_const__3225710767((a_0_box.IntVal) != (0), uint32(v_1_box.IntVal)))
		})
	})
	return cache_Data_Function_const__3225710767
}

var cache_Data_Function_const__3435219389 gopurs_runtime.Value
var once_Data_Function_const__3435219389 sync.Once

func Get_Data_Function_const__3435219389() gopurs_runtime.Value {
	once_Data_Function_const__3435219389.Do(func() {
		cache_Data_Function_const__3435219389 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Function_const__3435219389((a_0_box.IntVal) != (0), v_1_box))
		})
	})
	return cache_Data_Function_const__3435219389
}

var cache_Data_Function_const__2561705780 gopurs_runtime.Value
var once_Data_Function_const__2561705780 sync.Once

func Get_Data_Function_const__2561705780() gopurs_runtime.Value {
	once_Data_Function_const__2561705780.Do(func() {
		cache_Data_Function_const__2561705780 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Function_const__2561705780((a_0_box.IntVal) != (0), *((*[]gopurs_runtime.Value)(v_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Function_const__2561705780
}

var cache_Data_Function_const__2007806465 gopurs_runtime.Value
var once_Data_Function_const__2007806465 sync.Once

func Get_Data_Function_const__2007806465() gopurs_runtime.Value {
	once_Data_Function_const__2007806465.Do(func() {
		cache_Data_Function_const__2007806465 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Function_const__2007806465((a_0_box.IntVal) != (0), (v_1_box.IntVal) != (0)))
		})
	})
	return cache_Data_Function_const__2007806465
}

var cache_Data_Function_const__318080054 gopurs_runtime.Value
var once_Data_Function_const__318080054 sync.Once

func Get_Data_Function_const__318080054() gopurs_runtime.Value {
	once_Data_Function_const__318080054.Do(func() {
		cache_Data_Function_const__318080054 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Function_const__318080054((a_0_box.IntVal) != (0), v_1_box.IntVal))
		})
	})
	return cache_Data_Function_const__318080054
}

var cache_Data_Function_const__2630302150 gopurs_runtime.Value
var once_Data_Function_const__2630302150 sync.Once

func Get_Data_Function_const__2630302150() gopurs_runtime.Value {
	once_Data_Function_const__2630302150.Do(func() {
		cache_Data_Function_const__2630302150 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Function_const__2630302150((a_0_box.IntVal) != (0), v_1_box.FloatVal()))
		})
	})
	return cache_Data_Function_const__2630302150
}

var cache_Data_Function_const__1410765104 gopurs_runtime.Value
var once_Data_Function_const__1410765104 sync.Once

func Get_Data_Function_const__1410765104() gopurs_runtime.Value {
	once_Data_Function_const__1410765104.Do(func() {
		cache_Data_Function_const__1410765104 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Function_const__1410765104((a_0_box.IntVal) != (0), v_1_box.StrVal()))
		})
	})
	return cache_Data_Function_const__1410765104
}

var cache_Data_Function_const__1399880067 gopurs_runtime.Value
var once_Data_Function_const__1399880067 sync.Once

func Get_Data_Function_const__1399880067() gopurs_runtime.Value {
	once_Data_Function_const__1399880067.Do(func() {
		cache_Data_Function_const__1399880067 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Function_const__1399880067((a_0_box.IntVal) != (0), v_1_box))
		})
	})
	return cache_Data_Function_const__1399880067
}

var cache_Data_Function_const__425793770 gopurs_runtime.Value
var once_Data_Function_const__425793770 sync.Once

func Get_Data_Function_const__425793770() gopurs_runtime.Value {
	once_Data_Function_const__425793770.Do(func() {
		cache_Data_Function_const__425793770 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_const__425793770(a_unused_0_box, v_1_box)
		})
	})
	return cache_Data_Function_const__425793770
}

var cache_Data_Function_const__855710010 gopurs_runtime.Value
var once_Data_Function_const__855710010 sync.Once

func Get_Data_Function_const__855710010() gopurs_runtime.Value {
	once_Data_Function_const__855710010.Do(func() {
		cache_Data_Function_const__855710010 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_const__855710010(a_unused_0_box, v_1_box)
		})
	})
	return cache_Data_Function_const__855710010
}

var cache_Data_Function_const__1353723968 gopurs_runtime.Value
var once_Data_Function_const__1353723968 sync.Once

func Get_Data_Function_const__1353723968() gopurs_runtime.Value {
	once_Data_Function_const__1353723968.Do(func() {
		cache_Data_Function_const__1353723968 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_const__1353723968(a_unused_0_box, v_1_box)
		})
	})
	return cache_Data_Function_const__1353723968
}

var cache_Data_Function_const__244128784 gopurs_runtime.Value
var once_Data_Function_const__244128784 sync.Once

func Get_Data_Function_const__244128784() gopurs_runtime.Value {
	once_Data_Function_const__244128784.Do(func() {
		cache_Data_Function_const__244128784 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_const__244128784(a_unused_0_box, v_1_box)
		})
	})
	return cache_Data_Function_const__244128784
}

var cache_Data_Function_const__2389296201 gopurs_runtime.Value
var once_Data_Function_const__2389296201 sync.Once

func Get_Data_Function_const__2389296201() gopurs_runtime.Value {
	once_Data_Function_const__2389296201.Do(func() {
		cache_Data_Function_const__2389296201 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_const__2389296201(a_unused_0_box, v_1_box)
		})
	})
	return cache_Data_Function_const__2389296201
}

var cache_Data_Function_const__1638250638 gopurs_runtime.Value
var once_Data_Function_const__1638250638 sync.Once

func Get_Data_Function_const__1638250638() gopurs_runtime.Value {
	once_Data_Function_const__1638250638.Do(func() {
		cache_Data_Function_const__1638250638 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_const__1638250638(a_unused_0_box, v_1_box)
		})
	})
	return cache_Data_Function_const__1638250638
}

var cache_Data_Function_const__3353209993 gopurs_runtime.Value
var once_Data_Function_const__3353209993 sync.Once

func Get_Data_Function_const__3353209993() gopurs_runtime.Value {
	once_Data_Function_const__3353209993.Do(func() {
		cache_Data_Function_const__3353209993 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_const__3353209993(a_unused_0_box, v_1_box)
		})
	})
	return cache_Data_Function_const__3353209993
}

var cache_Data_Function_const__3916468150 gopurs_runtime.Value
var once_Data_Function_const__3916468150 sync.Once

func Get_Data_Function_const__3916468150() gopurs_runtime.Value {
	once_Data_Function_const__3916468150.Do(func() {
		cache_Data_Function_const__3916468150 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Function_const__3916468150(a_0_box.IntVal, v_1_box.IntVal))
		})
	})
	return cache_Data_Function_const__3916468150
}

var cache_Data_Function_const__3422218083 gopurs_runtime.Value
var once_Data_Function_const__3422218083 sync.Once

func Get_Data_Function_const__3422218083() gopurs_runtime.Value {
	once_Data_Function_const__3422218083.Do(func() {
		cache_Data_Function_const__3422218083 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Function_const__3422218083(a_0_box.IntVal, v_1_box))
		})
	})
	return cache_Data_Function_const__3422218083
}

var cache_Data_Function_const__657731907 gopurs_runtime.Value
var once_Data_Function_const__657731907 sync.Once

func Get_Data_Function_const__657731907() gopurs_runtime.Value {
	once_Data_Function_const__657731907.Do(func() {
		cache_Data_Function_const__657731907 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_Function_const__657731907(a_0_box.StrVal(), v_1_box))
		})
	})
	return cache_Data_Function_const__657731907
}

var cache_Data_Function_applyN gopurs_runtime.Value
var once_Data_Function_applyN sync.Once

func Get_Data_Function_applyN() gopurs_runtime.Value {
	once_Data_Function_applyN.Do(func() {
		cache_Data_Function_applyN = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_applyN(f_0_box)
		})
	})
	return cache_Data_Function_applyN
}

var cache_Data_Function_applyFlipped gopurs_runtime.Value
var once_Data_Function_applyFlipped sync.Once

func Get_Data_Function_applyFlipped() gopurs_runtime.Value {
	once_Data_Function_applyFlipped.Do(func() {
		cache_Data_Function_applyFlipped = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_applyFlipped(x_0_box, f_1_box)
		})
	})
	return cache_Data_Function_applyFlipped
}

var cache_Data_Function_apply gopurs_runtime.Value
var once_Data_Function_apply sync.Once

func Get_Data_Function_apply() gopurs_runtime.Value {
	once_Data_Function_apply.Do(func() {
		cache_Data_Function_apply = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply(f_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply
}

var cache_Data_Function_apply__3946155321 gopurs_runtime.Value
var once_Data_Function_apply__3946155321 sync.Once

func Get_Data_Function_apply__3946155321() gopurs_runtime.Value {
	once_Data_Function_apply__3946155321.Do(func() {
		cache_Data_Function_apply__3946155321 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__3946155321(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__3946155321
}

var cache_Data_Function_apply__2302590153 gopurs_runtime.Value
var once_Data_Function_apply__2302590153 sync.Once

func Get_Data_Function_apply__2302590153() gopurs_runtime.Value {
	once_Data_Function_apply__2302590153.Do(func() {
		cache_Data_Function_apply__2302590153 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__2302590153(f_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__2302590153
}

var cache_Data_Function_apply__1928284458 gopurs_runtime.Value
var once_Data_Function_apply__1928284458 sync.Once

func Get_Data_Function_apply__1928284458() gopurs_runtime.Value {
	once_Data_Function_apply__1928284458.Do(func() {
		cache_Data_Function_apply__1928284458 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__1928284458(f_unused_0_box, x_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_apply__1928284458
}

var cache_Data_Function_apply__3425554218 gopurs_runtime.Value
var once_Data_Function_apply__3425554218 sync.Once

func Get_Data_Function_apply__3425554218() gopurs_runtime.Value {
	once_Data_Function_apply__3425554218.Do(func() {
		cache_Data_Function_apply__3425554218 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__3425554218(f_unused_0_box, x_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_apply__3425554218
}

var cache_Data_Function_apply__1520674922 gopurs_runtime.Value
var once_Data_Function_apply__1520674922 sync.Once

func Get_Data_Function_apply__1520674922() gopurs_runtime.Value {
	once_Data_Function_apply__1520674922.Do(func() {
		cache_Data_Function_apply__1520674922 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__1520674922(f_unused_0_box, x_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_apply__1520674922
}

var cache_Data_Function_apply__3344524906 gopurs_runtime.Value
var once_Data_Function_apply__3344524906 sync.Once

func Get_Data_Function_apply__3344524906() gopurs_runtime.Value {
	once_Data_Function_apply__3344524906.Do(func() {
		cache_Data_Function_apply__3344524906 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__3344524906(f_unused_0_box, x_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_apply__3344524906
}

var cache_Data_Function_apply__2793916074 gopurs_runtime.Value
var once_Data_Function_apply__2793916074 sync.Once

func Get_Data_Function_apply__2793916074() gopurs_runtime.Value {
	once_Data_Function_apply__2793916074.Do(func() {
		cache_Data_Function_apply__2793916074 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__2793916074(f_unused_0_box, x_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_apply__2793916074
}

var cache_Data_Function_apply__1173000582 gopurs_runtime.Value
var once_Data_Function_apply__1173000582 sync.Once

func Get_Data_Function_apply__1173000582() gopurs_runtime.Value {
	once_Data_Function_apply__1173000582.Do(func() {
		cache_Data_Function_apply__1173000582 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Function_apply__1173000582(f_unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](x_1_box)))}
		})
	})
	return cache_Data_Function_apply__1173000582
}

var cache_Data_Function_apply__428839251 gopurs_runtime.Value
var once_Data_Function_apply__428839251 sync.Once

func Get_Data_Function_apply__428839251() gopurs_runtime.Value {
	once_Data_Function_apply__428839251.Do(func() {
		cache_Data_Function_apply__428839251 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__428839251(f_unused_0_box, x_unused_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_apply__428839251
}

var cache_Data_Function_apply__569550801 gopurs_runtime.Value
var once_Data_Function_apply__569550801 sync.Once

func Get_Data_Function_apply__569550801() gopurs_runtime.Value {
	once_Data_Function_apply__569550801.Do(func() {
		cache_Data_Function_apply__569550801 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__569550801(f_unused_0_box, x_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_apply__569550801
}

var cache_Data_Function_apply__3998497865 gopurs_runtime.Value
var once_Data_Function_apply__3998497865 sync.Once

func Get_Data_Function_apply__3998497865() gopurs_runtime.Value {
	once_Data_Function_apply__3998497865.Do(func() {
		cache_Data_Function_apply__3998497865 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__3998497865(f_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__3998497865
}

var cache_Data_Function_apply__543526857 gopurs_runtime.Value
var once_Data_Function_apply__543526857 sync.Once

func Get_Data_Function_apply__543526857() gopurs_runtime.Value {
	once_Data_Function_apply__543526857.Do(func() {
		cache_Data_Function_apply__543526857 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__543526857(f_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__543526857
}

var cache_Data_Function_apply__2561392351 gopurs_runtime.Value
var once_Data_Function_apply__2561392351 sync.Once

func Get_Data_Function_apply__2561392351() gopurs_runtime.Value {
	once_Data_Function_apply__2561392351.Do(func() {
		cache_Data_Function_apply__2561392351 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__2561392351(f_unused_0_box, Rebox_Data_Function_849153993_3232828411(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](x_1_box)))
		})
	})
	return cache_Data_Function_apply__2561392351
}

var cache_Data_Function_apply__2998264905 gopurs_runtime.Value
var once_Data_Function_apply__2998264905 sync.Once

func Get_Data_Function_apply__2998264905() gopurs_runtime.Value {
	once_Data_Function_apply__2998264905.Do(func() {
		cache_Data_Function_apply__2998264905 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__2998264905(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1_box))
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_apply__2998264905
}

var cache_Data_Function_apply__26935913 gopurs_runtime.Value
var once_Data_Function_apply__26935913 sync.Once

func Get_Data_Function_apply__26935913() gopurs_runtime.Value {
	once_Data_Function_apply__26935913.Do(func() {
		cache_Data_Function_apply__26935913 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_Function_apply__26935913(f_0_box, Rebox_Data_Function_3094389156_3839235747(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1_box))))}
		})
	})
	return cache_Data_Function_apply__26935913
}

var cache_Data_Function_apply__3632476649 gopurs_runtime.Value
var once_Data_Function_apply__3632476649 sync.Once

func Get_Data_Function_apply__3632476649() gopurs_runtime.Value {
	once_Data_Function_apply__3632476649.Do(func() {
		cache_Data_Function_apply__3632476649 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__3632476649(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1_box))
		})
	})
	return cache_Data_Function_apply__3632476649
}

var cache_Data_Function_apply__972864507 gopurs_runtime.Value
var once_Data_Function_apply__972864507 sync.Once

func Get_Data_Function_apply__972864507() gopurs_runtime.Value {
	once_Data_Function_apply__972864507.Do(func() {
		cache_Data_Function_apply__972864507 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__972864507(f_unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1_box))
		})
	})
	return cache_Data_Function_apply__972864507
}

var cache_Data_Function_apply__2647861350 gopurs_runtime.Value
var once_Data_Function_apply__2647861350 sync.Once

func Get_Data_Function_apply__2647861350() gopurs_runtime.Value {
	once_Data_Function_apply__2647861350.Do(func() {
		cache_Data_Function_apply__2647861350 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Function_apply__2647861350(f_unused_0_box, Rebox_Data_Function_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1_box))))
		})
	})
	return cache_Data_Function_apply__2647861350
}

var cache_Data_Function_apply__2639368963 gopurs_runtime.Value
var once_Data_Function_apply__2639368963 sync.Once

func Get_Data_Function_apply__2639368963() gopurs_runtime.Value {
	once_Data_Function_apply__2639368963.Do(func() {
		cache_Data_Function_apply__2639368963 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Function_apply__2639368963(f_unused_0_box, Rebox_Data_Function_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1_box))))
		})
	})
	return cache_Data_Function_apply__2639368963
}

var cache_Data_Function_apply__2730021546 gopurs_runtime.Value
var once_Data_Function_apply__2730021546 sync.Once

func Get_Data_Function_apply__2730021546() gopurs_runtime.Value {
	once_Data_Function_apply__2730021546.Do(func() {
		cache_Data_Function_apply__2730021546 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__2730021546(f_unused_0_box, Rebox_Data_Function_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](x_1_box)))
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_apply__2730021546
}

var cache_Data_Function_apply__3653967818 gopurs_runtime.Value
var once_Data_Function_apply__3653967818 sync.Once

func Get_Data_Function_apply__3653967818() gopurs_runtime.Value {
	once_Data_Function_apply__3653967818.Do(func() {
		cache_Data_Function_apply__3653967818 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__3653967818(f_unused_0_box, Rebox_Data_Function_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](x_1_box)))
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_apply__3653967818
}

var cache_Data_Function_apply__1272484046 gopurs_runtime.Value
var once_Data_Function_apply__1272484046 sync.Once

func Get_Data_Function_apply__1272484046() gopurs_runtime.Value {
	once_Data_Function_apply__1272484046.Do(func() {
		cache_Data_Function_apply__1272484046 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__1272484046(f_unused_0_box, x_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Function_apply__1272484046
}

var cache_Data_Function_apply__1141988375 gopurs_runtime.Value
var once_Data_Function_apply__1141988375 sync.Once

func Get_Data_Function_apply__1141988375() gopurs_runtime.Value {
	once_Data_Function_apply__1141988375.Do(func() {
		cache_Data_Function_apply__1141988375 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__1141988375(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__1141988375
}

var cache_Data_Function_apply__1383316684 gopurs_runtime.Value
var once_Data_Function_apply__1383316684 sync.Once

func Get_Data_Function_apply__1383316684() gopurs_runtime.Value {
	once_Data_Function_apply__1383316684.Do(func() {
		cache_Data_Function_apply__1383316684 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__1383316684(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__1383316684
}

var cache_Data_Function_apply__1171255079 gopurs_runtime.Value
var once_Data_Function_apply__1171255079 sync.Once

func Get_Data_Function_apply__1171255079() gopurs_runtime.Value {
	once_Data_Function_apply__1171255079.Do(func() {
		cache_Data_Function_apply__1171255079 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__1171255079(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__1171255079
}

var cache_Data_Function_apply__2435095346 gopurs_runtime.Value
var once_Data_Function_apply__2435095346 sync.Once

func Get_Data_Function_apply__2435095346() gopurs_runtime.Value {
	once_Data_Function_apply__2435095346.Do(func() {
		cache_Data_Function_apply__2435095346 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__2435095346(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__2435095346
}

var cache_Data_Function_apply__734028734 gopurs_runtime.Value
var once_Data_Function_apply__734028734 sync.Once

func Get_Data_Function_apply__734028734() gopurs_runtime.Value {
	once_Data_Function_apply__734028734.Do(func() {
		cache_Data_Function_apply__734028734 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__734028734(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__734028734
}

var cache_Data_Function_apply__1422622261 gopurs_runtime.Value
var once_Data_Function_apply__1422622261 sync.Once

func Get_Data_Function_apply__1422622261() gopurs_runtime.Value {
	once_Data_Function_apply__1422622261.Do(func() {
		cache_Data_Function_apply__1422622261 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__1422622261(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__1422622261
}

var cache_Data_Function_apply__2813923095 gopurs_runtime.Value
var once_Data_Function_apply__2813923095 sync.Once

func Get_Data_Function_apply__2813923095() gopurs_runtime.Value {
	once_Data_Function_apply__2813923095.Do(func() {
		cache_Data_Function_apply__2813923095 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__2813923095(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__2813923095
}

var cache_Data_Function_apply__1413608268 gopurs_runtime.Value
var once_Data_Function_apply__1413608268 sync.Once

func Get_Data_Function_apply__1413608268() gopurs_runtime.Value {
	once_Data_Function_apply__1413608268.Do(func() {
		cache_Data_Function_apply__1413608268 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__1413608268(f_0_box, x_unused_1_box)
		})
	})
	return cache_Data_Function_apply__1413608268
}

var cache_Data_Function_apply__1659262775 gopurs_runtime.Value
var once_Data_Function_apply__1659262775 sync.Once

func Get_Data_Function_apply__1659262775() gopurs_runtime.Value {
	once_Data_Function_apply__1659262775.Do(func() {
		cache_Data_Function_apply__1659262775 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__1659262775(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__1659262775
}

var cache_Data_Function_apply__24380833 gopurs_runtime.Value
var once_Data_Function_apply__24380833 sync.Once

func Get_Data_Function_apply__24380833() gopurs_runtime.Value {
	once_Data_Function_apply__24380833.Do(func() {
		cache_Data_Function_apply__24380833 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__24380833(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__24380833
}

var cache_Data_Function_apply__341522372 gopurs_runtime.Value
var once_Data_Function_apply__341522372 sync.Once

func Get_Data_Function_apply__341522372() gopurs_runtime.Value {
	once_Data_Function_apply__341522372.Do(func() {
		cache_Data_Function_apply__341522372 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__341522372(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__341522372
}

var cache_Data_Function_apply__556989731 gopurs_runtime.Value
var once_Data_Function_apply__556989731 sync.Once

func Get_Data_Function_apply__556989731() gopurs_runtime.Value {
	once_Data_Function_apply__556989731.Do(func() {
		cache_Data_Function_apply__556989731 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__556989731(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__556989731
}

var cache_Data_Function_apply__483590224 gopurs_runtime.Value
var once_Data_Function_apply__483590224 sync.Once

func Get_Data_Function_apply__483590224() gopurs_runtime.Value {
	once_Data_Function_apply__483590224.Do(func() {
		cache_Data_Function_apply__483590224 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__483590224(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__483590224
}

var cache_Data_Function_apply__2279232018 gopurs_runtime.Value
var once_Data_Function_apply__2279232018 sync.Once

func Get_Data_Function_apply__2279232018() gopurs_runtime.Value {
	once_Data_Function_apply__2279232018.Do(func() {
		cache_Data_Function_apply__2279232018 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__2279232018(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__2279232018
}

var cache_Data_Function_apply__1755036686 gopurs_runtime.Value
var once_Data_Function_apply__1755036686 sync.Once

func Get_Data_Function_apply__1755036686() gopurs_runtime.Value {
	once_Data_Function_apply__1755036686.Do(func() {
		cache_Data_Function_apply__1755036686 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__1755036686(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__1755036686
}

var cache_Data_Function_apply__1851282762 gopurs_runtime.Value
var once_Data_Function_apply__1851282762 sync.Once

func Get_Data_Function_apply__1851282762() gopurs_runtime.Value {
	once_Data_Function_apply__1851282762.Do(func() {
		cache_Data_Function_apply__1851282762 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__1851282762(f_unused_0_box, x_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict2("buffers", "readagain", gopurs_runtime.Array(orig.buffers), gopurs_runtime.Bool(orig.readagain))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_apply__1851282762
}

var cache_Data_Function_apply__2798404781 gopurs_runtime.Value
var once_Data_Function_apply__2798404781 sync.Once

func Get_Data_Function_apply__2798404781() gopurs_runtime.Value {
	once_Data_Function_apply__2798404781.Do(func() {
		cache_Data_Function_apply__2798404781 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__2798404781(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__2798404781
}

var cache_Data_Function_apply__1937891721 gopurs_runtime.Value
var once_Data_Function_apply__1937891721 sync.Once

func Get_Data_Function_apply__1937891721() gopurs_runtime.Value {
	once_Data_Function_apply__1937891721.Do(func() {
		cache_Data_Function_apply__1937891721 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__1937891721(f_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__1937891721
}

var cache_Data_Function_apply__2514240489 gopurs_runtime.Value
var once_Data_Function_apply__2514240489 sync.Once

func Get_Data_Function_apply__2514240489() gopurs_runtime.Value {
	once_Data_Function_apply__2514240489.Do(func() {
		cache_Data_Function_apply__2514240489 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__2514240489(f_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__2514240489
}

var cache_Data_Function_apply__1942606537 gopurs_runtime.Value
var once_Data_Function_apply__1942606537 sync.Once

func Get_Data_Function_apply__1942606537() gopurs_runtime.Value {
	once_Data_Function_apply__1942606537.Do(func() {
		cache_Data_Function_apply__1942606537 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__1942606537(f_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__1942606537
}

var cache_Data_Function_apply__2422281551 gopurs_runtime.Value
var once_Data_Function_apply__2422281551 sync.Once

func Get_Data_Function_apply__2422281551() gopurs_runtime.Value {
	once_Data_Function_apply__2422281551.Do(func() {
		cache_Data_Function_apply__2422281551 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_3123684004_1293498952(Call_Data_Function_apply__2422281551(f_unused_0_box, x_1_box)))}
		})
	})
	return cache_Data_Function_apply__2422281551
}

var cache_Data_Function_apply__650006357 gopurs_runtime.Value
var once_Data_Function_apply__650006357 sync.Once

func Get_Data_Function_apply__650006357() gopurs_runtime.Value {
	once_Data_Function_apply__650006357.Do(func() {
		cache_Data_Function_apply__650006357 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__650006357(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__650006357
}

var cache_Data_Function_apply__1319276646 gopurs_runtime.Value
var once_Data_Function_apply__1319276646 sync.Once

func Get_Data_Function_apply__1319276646() gopurs_runtime.Value {
	once_Data_Function_apply__1319276646.Do(func() {
		cache_Data_Function_apply__1319276646 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__1319276646(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__1319276646
}

var cache_Data_Function_apply__2293662334 gopurs_runtime.Value
var once_Data_Function_apply__2293662334 sync.Once

func Get_Data_Function_apply__2293662334() gopurs_runtime.Value {
	once_Data_Function_apply__2293662334.Do(func() {
		cache_Data_Function_apply__2293662334 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Function_apply__2293662334(f_unused_0_box, *((*[]gopurs_runtime.Value)(x_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Function_apply__2293662334
}

var cache_Data_Function_apply__862467209 gopurs_runtime.Value
var once_Data_Function_apply__862467209 sync.Once

func Get_Data_Function_apply__862467209() gopurs_runtime.Value {
	once_Data_Function_apply__862467209.Do(func() {
		cache_Data_Function_apply__862467209 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Data_Function_apply__862467209(f_0_box, *((*[]gopurs_runtime.Value)(x_1_box.UnsafePtr))))
		})
	})
	return cache_Data_Function_apply__862467209
}

var cache_Data_Function_apply__1378432126 gopurs_runtime.Value
var once_Data_Function_apply__1378432126 sync.Once

func Get_Data_Function_apply__1378432126() gopurs_runtime.Value {
	once_Data_Function_apply__1378432126.Do(func() {
		cache_Data_Function_apply__1378432126 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				arr := Call_Data_Function_apply__1378432126(f_unused_0_box, func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr {
						unboxed[i] = v.IntVal
					}
					return unboxed
				}())
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()
		})
	})
	return cache_Data_Function_apply__1378432126
}

var cache_Data_Function_apply__2114047305 gopurs_runtime.Value
var once_Data_Function_apply__2114047305 sync.Once

func Get_Data_Function_apply__2114047305() gopurs_runtime.Value {
	once_Data_Function_apply__2114047305.Do(func() {
		cache_Data_Function_apply__2114047305 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool(Call_Data_Function_apply__2114047305(f_0_box, (x_1_box.IntVal) != (0)))
		})
	})
	return cache_Data_Function_apply__2114047305
}

var cache_Data_Function_apply__676573833 gopurs_runtime.Value
var once_Data_Function_apply__676573833 sync.Once

func Get_Data_Function_apply__676573833() gopurs_runtime.Value {
	once_Data_Function_apply__676573833.Do(func() {
		cache_Data_Function_apply__676573833 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__676573833(f_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__676573833
}

var cache_Data_Function_apply__973594793 gopurs_runtime.Value
var once_Data_Function_apply__973594793 sync.Once

func Get_Data_Function_apply__973594793() gopurs_runtime.Value {
	once_Data_Function_apply__973594793.Do(func() {
		cache_Data_Function_apply__973594793 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__973594793(f_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__973594793
}

var cache_Data_Function_apply__3402160552 gopurs_runtime.Value
var once_Data_Function_apply__3402160552 sync.Once

func Get_Data_Function_apply__3402160552() gopurs_runtime.Value {
	once_Data_Function_apply__3402160552.Do(func() {
		cache_Data_Function_apply__3402160552 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__3402160552(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__3402160552
}

var cache_Data_Function_apply__1639137429 gopurs_runtime.Value
var once_Data_Function_apply__1639137429 sync.Once

func Get_Data_Function_apply__1639137429() gopurs_runtime.Value {
	once_Data_Function_apply__1639137429.Do(func() {
		cache_Data_Function_apply__1639137429 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__1639137429(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__1639137429
}

var cache_Data_Function_apply__3034864212 gopurs_runtime.Value
var once_Data_Function_apply__3034864212 sync.Once

func Get_Data_Function_apply__3034864212() gopurs_runtime.Value {
	once_Data_Function_apply__3034864212.Do(func() {
		cache_Data_Function_apply__3034864212 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__3034864212(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__3034864212
}

var cache_Data_Function_apply__2984437716 gopurs_runtime.Value
var once_Data_Function_apply__2984437716 sync.Once

func Get_Data_Function_apply__2984437716() gopurs_runtime.Value {
	once_Data_Function_apply__2984437716.Do(func() {
		cache_Data_Function_apply__2984437716 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__2984437716(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__2984437716
}

var cache_Data_Function_apply__1415232852 gopurs_runtime.Value
var once_Data_Function_apply__1415232852 sync.Once

func Get_Data_Function_apply__1415232852() gopurs_runtime.Value {
	once_Data_Function_apply__1415232852.Do(func() {
		cache_Data_Function_apply__1415232852 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__1415232852(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__1415232852
}

var cache_Data_Function_apply__3334826457 gopurs_runtime.Value
var once_Data_Function_apply__3334826457 sync.Once

func Get_Data_Function_apply__3334826457() gopurs_runtime.Value {
	once_Data_Function_apply__3334826457.Do(func() {
		cache_Data_Function_apply__3334826457 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__3334826457(f_unused_0_box, x_1_box)
		})
	})
	return cache_Data_Function_apply__3334826457
}

var cache_Data_Function_apply__3228670985 gopurs_runtime.Value
var once_Data_Function_apply__3228670985 sync.Once

func Get_Data_Function_apply__3228670985() gopurs_runtime.Value {
	once_Data_Function_apply__3228670985.Do(func() {
		cache_Data_Function_apply__3228670985 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer(Call_Data_Function_apply__3228670985(f_0_box, x_1_box))}
		})
	})
	return cache_Data_Function_apply__3228670985
}

var cache_Data_Function_apply__355458268 gopurs_runtime.Value
var once_Data_Function_apply__355458268 sync.Once

func Get_Data_Function_apply__355458268() gopurs_runtime.Value {
	once_Data_Function_apply__355458268.Do(func() {
		cache_Data_Function_apply__355458268 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Function_apply__355458268(f_unused_0_box, x_1_box.FloatVal()))
		})
	})
	return cache_Data_Function_apply__355458268
}

var cache_Data_Function_apply__2245704861 gopurs_runtime.Value
var once_Data_Function_apply__2245704861 sync.Once

func Get_Data_Function_apply__2245704861() gopurs_runtime.Value {
	once_Data_Function_apply__2245704861.Do(func() {
		cache_Data_Function_apply__2245704861 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Function_apply__2245704861(f_unused_0_box, x_1_box.FloatVal()))
		})
	})
	return cache_Data_Function_apply__2245704861
}

var cache_Data_Function_apply__2146671458 gopurs_runtime.Value
var once_Data_Function_apply__2146671458 sync.Once

func Get_Data_Function_apply__2146671458() gopurs_runtime.Value {
	once_Data_Function_apply__2146671458.Do(func() {
		cache_Data_Function_apply__2146671458 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__2146671458(f_unused_0_box, func() struct {
				isSuspended gopurs_runtime.Value
				join        gopurs_runtime.Value
				kill        gopurs_runtime.Value
				onComplete  gopurs_runtime.Value
				run         gopurs_runtime.Value
				/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
			} {
				orig := x_1_box
				_ = orig
				clone := struct {
					isSuspended gopurs_runtime.Value
					join        gopurs_runtime.Value
					kill        gopurs_runtime.Value
					onComplete  gopurs_runtime.Value
					run         gopurs_runtime.Value
					/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
				}{}
				clone.isSuspended = gopurs_runtime.RecordGet(orig, "isSuspended")
				clone.join = gopurs_runtime.RecordGet(orig, "join")
				clone.kill = gopurs_runtime.RecordGet(orig, "kill")
				clone.onComplete = gopurs_runtime.RecordGet(orig, "onComplete")
				clone.run = gopurs_runtime.RecordGet(orig, "run")
				return clone
			}())
		})
	})
	return cache_Data_Function_apply__2146671458
}

var cache_Data_Function_apply__190415331 gopurs_runtime.Value
var once_Data_Function_apply__190415331 sync.Once

func Get_Data_Function_apply__190415331() gopurs_runtime.Value {
	once_Data_Function_apply__190415331.Do(func() {
		cache_Data_Function_apply__190415331 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__190415331(f_unused_0_box, func() struct {
				body   string
				status int64
				/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
			} {
				orig := x_1_box
				_ = orig
				clone := struct {
					body   string
					status int64
					/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
				}{}
				clone.body = gopurs_runtime.RecordGet(orig, "body").StrVal()
				clone.status = gopurs_runtime.RecordGet(orig, "status").IntVal
				return clone
			}())
		})
	})
	return cache_Data_Function_apply__190415331
}

var cache_Data_Function_apply__2201874889 gopurs_runtime.Value
var once_Data_Function_apply__2201874889 sync.Once

func Get_Data_Function_apply__2201874889() gopurs_runtime.Value {
	once_Data_Function_apply__2201874889.Do(func() {
		cache_Data_Function_apply__2201874889 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__2201874889(f_0_box, func() struct {
				body   string
				status int64
				/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
			} {
				orig := x_1_box
				_ = orig
				clone := struct {
					body   string
					status int64
					/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
				}{}
				clone.body = gopurs_runtime.RecordGet(orig, "body").StrVal()
				clone.status = gopurs_runtime.RecordGet(orig, "status").IntVal
				return clone
			}())
		})
	})
	return cache_Data_Function_apply__2201874889
}

var cache_Data_Function_apply__2515510025 gopurs_runtime.Value
var once_Data_Function_apply__2515510025 sync.Once

func Get_Data_Function_apply__2515510025() gopurs_runtime.Value {
	once_Data_Function_apply__2515510025.Do(func() {
		cache_Data_Function_apply__2515510025 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Function_apply__2515510025(f_0_box, x_1_box.StrVal())
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Function_apply__2515510025
}

var cache_Data_Function_apply__1333176700 gopurs_runtime.Value
var once_Data_Function_apply__1333176700 sync.Once

func Get_Data_Function_apply__1333176700() gopurs_runtime.Value {
	once_Data_Function_apply__1333176700.Do(func() {
		cache_Data_Function_apply__1333176700 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Function_apply__1333176700(f_unused_0_box, x_1_box.StrVal())), UnsafePtr: nil}
		})
	})
	return cache_Data_Function_apply__1333176700
}

var cache_Data_Function_apply__826080446 gopurs_runtime.Value
var once_Data_Function_apply__826080446 sync.Once

func Get_Data_Function_apply__826080446() gopurs_runtime.Value {
	once_Data_Function_apply__826080446.Do(func() {
		cache_Data_Function_apply__826080446 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__826080446(f_unused_0_box, x_1_box.StrVal())
		})
	})
	return cache_Data_Function_apply__826080446
}

var cache_Data_Function_apply__596732553 gopurs_runtime.Value
var once_Data_Function_apply__596732553 sync.Once

func Get_Data_Function_apply__596732553() gopurs_runtime.Value {
	once_Data_Function_apply__596732553.Do(func() {
		cache_Data_Function_apply__596732553 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Function_apply__596732553(f_0_box, x_1_box.StrVal())
		})
	})
	return cache_Data_Function_apply__596732553
}

func Call_Data_Function_on(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var g_1 gopurs_runtime.Value = g_1_loop
	_ = g_1
	var x_2 gopurs_runtime.Value = x_2_loop
	_ = x_2
	var y_3 gopurs_runtime.Value = y_3_loop
	_ = y_3
	return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3))
}

func Call_Data_Function_flip(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var b_1 gopurs_runtime.Value = b_1_loop
	_ = b_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_Data_Function_flip__1501157206(f_unused_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Time_Time, a_2_loop *Constructor_Data_Date_Date) *Constructor_Data_DateTime_DateTime {
flip__1501157206:
	for {
		if false {
			continue flip__1501157206
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var b_1 *Constructor_Data_Time_Time = b_1_loop
		_ = b_1
		var a_2 *Constructor_Data_Date_Date = a_2_loop
		_ = a_2
		return gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](gopurs_runtime.Apply2(Get_Data_DateTime_DateTime(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(a_2)}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(b_1)}))
	}
}

func Call_Data_Function_flip__2508043272(f_unused_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
flip__2508043272:
	for {
		if false {
			continue flip__2508043272
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var b_1 gopurs_runtime.Value = b_1_loop
		_ = b_1
		var a_2 gopurs_runtime.Value = a_2_loop
		_ = a_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply3(Get_Control_Bind_bind(), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Either_bindEither()))}, a_2, b_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_flip__95728040(f_unused_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
flip__95728040:
	for {
		if false {
			continue flip__95728040
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var b_1 gopurs_runtime.Value = b_1_loop
		_ = b_1
		var a_2 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = a_2_loop
		_ = a_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_2280409795_3094389156(Rebox_Data_Function_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Control_Bind_bind(), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_3887487416_2748095225(Rebox_Data_Function_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe()))))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_2280409795_3094389156(a_2))}, b_1)))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Function_flip__1515853032(f_unused_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop *Constructor_Data_Maybe_Just[string]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
flip__1515853032:
	for {
		if false {
			continue flip__1515853032
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var b_1 gopurs_runtime.Value = b_1_loop
		_ = b_1
		var a_2 *Constructor_Data_Maybe_Just[string] = a_2_loop
		_ = a_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_742090555_3094389156(Rebox_Data_Function_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Control_Bind_bind(), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_3887487416_2748095225(Rebox_Data_Function_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe()))))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_742090555_3094389156(a_2))}, b_1)))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Function_flip__2414803565(f_unused_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
flip__2414803565:
	for {
		if false {
			continue flip__2414803565
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var b_1 gopurs_runtime.Value = b_1_loop
		_ = b_1
		var a_2 gopurs_runtime.Value = a_2_loop
		_ = a_2
		return gopurs_runtime.Apply3(Get_Control_Bind_bind(), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()))}, a_2, b_1)
	}
}

func Call_Data_Function_flip__377817773(f_unused_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
flip__377817773:
	for {
		if false {
			continue flip__377817773
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var b_1 gopurs_runtime.Value = b_1_loop
		_ = b_1
		var a_2 gopurs_runtime.Value = a_2_loop
		_ = a_2
		return gopurs_runtime.Apply3(Get_Control_Bind_bind(), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()))}, a_2, b_1)
	}
}

func Call_Data_Function_flip__1753236854(f_unused_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
flip__1753236854:
	for {
		if false {
			continue flip__1753236854
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var b_1 gopurs_runtime.Value = b_1_loop
		_ = b_1
		var a_2 gopurs_runtime.Value = a_2_loop
		_ = a_2
		return gopurs_runtime.Apply3(Get_Control_Bind_bind(), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()))}, a_2, b_1)
	}
}

func Call_Data_Function_flip__379318895(f_unused_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]], a_2_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
flip__379318895:
	for {
		if false {
			continue flip__379318895
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var b_1 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = b_1_loop
		_ = b_1
		var a_2 gopurs_runtime.Value = a_2_loop
		_ = a_2
		return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_Interval_Duration_Iso_foldMap(), a_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_2442833393_849153993(b_1))}))
	}
}

func Call_Data_Function_flip__650029772(f_unused_0_loop gopurs_runtime.Value, b_1_loop int64, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
flip__650029772:
	for {
		if false {
			continue flip__650029772
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var b_1 int64 = b_1_loop
		_ = b_1
		var a_2 gopurs_runtime.Value = a_2_loop
		_ = a_2
		return gopurs_runtime.Apply2(Get_Node_Stream_Aff_tailRecM(), a_2, gopurs_runtime.Int(b_1))
	}
}

func Call_Data_Function_flip__1483642665(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Date_Date, a_2_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
flip__1483642665:
	for {
		if false {
			continue flip__1483642665
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var b_1 *Constructor_Data_Date_Date = b_1_loop
		_ = b_1
		var a_2 int64 = a_2_loop
		_ = a_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_2280409795_3094389156(Rebox_Data_Function_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(a_2), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(b_1)})))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Function_go__const(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var a_0 gopurs_runtime.Value = a_0_loop
	_ = a_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return a_0
}

func Call_Data_Function_const__1740749855(a_unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__1740749855:
	for {
		if false {
			continue const__1740749855
		}
		var a_unused_0 gopurs_runtime.Value = a_unused_0_loop
		_ = a_unused_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return Get_Data_Argonaut_Core_jsonNull()
	}
}

func Call_Data_Function_const__3605917656(a_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
const__3605917656:
	for {
		if false {
			continue const__3605917656
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
		_ = v_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := a_0; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_const__1229331832(a_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
const__1229331832:
	for {
		if false {
			continue const__1229331832
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
		_ = v_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := a_0; if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_const__3126182883(a_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
const__3126182883:
	for {
		if false {
			continue const__3126182883
		}
		var a_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = a_0_loop
		_ = a_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(a_0)}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Function_const__4207166068(a_unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__4207166068:
	for {
		if false {
			continue const__4207166068
		}
		var a_unused_0 gopurs_runtime.Value = a_unused_0_loop
		_ = a_unused_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return gopurs_runtime.Apply(Get_Effect_Aff_pure6(), Get_Data_Unit_unit())
	}
}

func Call_Data_Function_const__320519811(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__320519811:
	for {
		if false {
			continue const__320519811
		}
		var a_0 gopurs_runtime.Value = a_0_loop
		_ = a_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return a_0
	}
}

func Call_Data_Function_const__2941505456(a_unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__2941505456:
	for {
		if false {
			continue const__2941505456
		}
		var a_unused_0 gopurs_runtime.Value = a_unused_0_loop
		_ = a_unused_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return gopurs_runtime.Apply(Get_Effect_Aff_pure2(), Get_Data_Unit_unit())
	}
}

func Call_Data_Function_const__2476124256(a_unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__2476124256:
	for {
		if false {
			continue const__2476124256
		}
		var a_unused_0 gopurs_runtime.Value = a_unused_0_loop
		_ = a_unused_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return gopurs_runtime.Apply(Get_Effect_Aff_pure2(), Get_Data_Unit_unit())
	}
}

func Call_Data_Function_const__3225710767(a_0_loop bool, v_1_loop uint32) bool {
const__3225710767:
	for {
		if false {
			continue const__3225710767
		}
		var a_0 bool = a_0_loop
		_ = a_0
		var v_1 uint32 = v_1_loop
		_ = v_1
		return a_0
	}
}

func Call_Data_Function_const__3435219389(a_0_loop bool, v_1_loop gopurs_runtime.Value) bool {
const__3435219389:
	for {
		if false {
			continue const__3435219389
		}
		var a_0 bool = a_0_loop
		_ = a_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return a_0
	}
}

func Call_Data_Function_const__2561705780(a_0_loop bool, v_1_loop []gopurs_runtime.Value) bool {
const__2561705780:
	for {
		if false {
			continue const__2561705780
		}
		var a_0 bool = a_0_loop
		_ = a_0
		var v_1 []gopurs_runtime.Value = v_1_loop
		_ = v_1
		return a_0
	}
}

func Call_Data_Function_const__2007806465(a_0_loop bool, v_1_loop bool) bool {
const__2007806465:
	for {
		if false {
			continue const__2007806465
		}
		var a_0 bool = a_0_loop
		_ = a_0
		var v_1 bool = v_1_loop
		_ = v_1
		return a_0
	}
}

func Call_Data_Function_const__318080054(a_0_loop bool, v_1_loop int64) bool {
const__318080054:
	for {
		if false {
			continue const__318080054
		}
		var a_0 bool = a_0_loop
		_ = a_0
		var v_1 int64 = v_1_loop
		_ = v_1
		return a_0
	}
}

func Call_Data_Function_const__2630302150(a_0_loop bool, v_1_loop float64) bool {
const__2630302150:
	for {
		if false {
			continue const__2630302150
		}
		var a_0 bool = a_0_loop
		_ = a_0
		var v_1 float64 = v_1_loop
		_ = v_1
		return a_0
	}
}

func Call_Data_Function_const__1410765104(a_0_loop bool, v_1_loop string) bool {
const__1410765104:
	for {
		if false {
			continue const__1410765104
		}
		var a_0 bool = a_0_loop
		_ = a_0
		var v_1 string = v_1_loop
		_ = v_1
		return a_0
	}
}

func Call_Data_Function_const__1399880067(a_0_loop bool, v_1_loop gopurs_runtime.Value) bool {
const__1399880067:
	for {
		if false {
			continue const__1399880067
		}
		var a_0 bool = a_0_loop
		_ = a_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return a_0
	}
}

func Call_Data_Function_const__425793770(a_unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__425793770:
	for {
		if false {
			continue const__425793770
		}
		var a_unused_0 gopurs_runtime.Value = a_unused_0_loop
		_ = a_unused_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return Get_Database_SQLite3_go__close()
	}
}

func Call_Data_Function_const__855710010(a_unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__855710010:
	for {
		if false {
			continue const__855710010
		}
		var a_unused_0 gopurs_runtime.Value = a_unused_0_loop
		_ = a_unused_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return Get_Database_SQLite3_go__close()
	}
}

func Call_Data_Function_const__1353723968(a_unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__1353723968:
	for {
		if false {
			continue const__1353723968
		}
		var a_unused_0 gopurs_runtime.Value = a_unused_0_loop
		_ = a_unused_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return gopurs_runtime.Apply2(Get_Example_HTTP_compose(), Get_Effect_Aff_Compat_fromEffectFnAff(), Get_Example_HTTP_stopImpl())
	}
}

func Call_Data_Function_const__244128784(a_unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__244128784:
	for {
		if false {
			continue const__244128784
		}
		var a_unused_0 gopurs_runtime.Value = a_unused_0_loop
		_ = a_unused_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return gopurs_runtime.Apply2(Get_Example_HTTP_compose(), Get_Effect_Aff_Compat_fromEffectFnAff(), Get_Example_HTTP_stopImpl())
	}
}

func Call_Data_Function_const__2389296201(a_unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__2389296201:
	for {
		if false {
			continue const__2389296201
		}
		var a_unused_0 gopurs_runtime.Value = a_unused_0_loop
		_ = a_unused_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return Get_Control_Apply_identity()
	}
}

func Call_Data_Function_const__1638250638(a_unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__1638250638:
	for {
		if false {
			continue const__1638250638
		}
		var a_unused_0 gopurs_runtime.Value = a_unused_0_loop
		_ = a_unused_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return Get_Effect_Aff_pure6()
	}
}

func Call_Data_Function_const__3353209993(a_unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
const__3353209993:
	for {
		if false {
			continue const__3353209993
		}
		var a_unused_0 gopurs_runtime.Value = a_unused_0_loop
		_ = a_unused_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return Get_Control_Apply_identity()
	}
}

func Call_Data_Function_const__3916468150(a_0_loop int64, v_1_loop int64) int64 {
const__3916468150:
	for {
		if false {
			continue const__3916468150
		}
		var a_0 int64 = a_0_loop
		_ = a_0
		var v_1 int64 = v_1_loop
		_ = v_1
		return a_0
	}
}

func Call_Data_Function_const__3422218083(a_0_loop int64, v_1_loop gopurs_runtime.Value) int64 {
const__3422218083:
	for {
		if false {
			continue const__3422218083
		}
		var a_0 int64 = a_0_loop
		_ = a_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return a_0
	}
}

func Call_Data_Function_const__657731907(a_0_loop string, v_1_loop gopurs_runtime.Value) string {
const__657731907:
	for {
		if false {
			continue const__657731907
		}
		var a_0 string = a_0_loop
		_ = a_0
		var v_1 gopurs_runtime.Value = v_1_loop
		_ = v_1
		return a_0
	}
}

func Call_Data_Function_applyN(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var Call_local_Data_Function_go__go_1_0_0 func(int64, gopurs_runtime.Value) gopurs_runtime.Value
	_ = Call_local_Data_Function_go__go_1_0_0
	var go__go_1_0_0 gopurs_runtime.Value
	_ = go__go_1_0_0
	Call_local_Data_Function_go__go_1_0_0 = func(n_2_loop int64, acc_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	go__go_1_0_0:
		for {
			if false {
				continue go__go_1_0_0
			}
			var n_2 int64 = n_2_loop
			_ = n_2
			var acc_3 gopurs_runtime.Value = acc_3_loop
			_ = acc_3
			var __t1 gopurs_runtime.Value
			{
				if (n_2) <= (int64(0)) {
					__t1 = acc_3
					goto end_branch_1
				} else {

				}
			}
			{
				n_2_loop = gopurs_runtime.IntSub(n_2, int64(1))
				acc_3_loop = gopurs_runtime.Apply(f_0, acc_3)
				continue go__go_1_0_0
				__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
			}
		end_branch_1:
			return __t1
		}
	}
	go__go_1_0_0 = gopurs_runtime.Func(func(n_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(acc_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_local_Data_Function_go__go_1_0_0(n_2_loop_val.IntVal, acc_3_loop_val)
		})
	})
	return go__go_1_0_0
}

func Call_Data_Function_applyFlipped(x_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	return gopurs_runtime.Apply(f_1, x_0)
}

func Call_Data_Function_apply(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(f_0, x_1)
}

func Call_Data_Function_apply__3946155321(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__3946155321:
	for {
		if false {
			continue apply__3946155321
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Error_UnexpectedValue(), x_1)
	}
}

func Call_Data_Function_apply__2302590153(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2302590153:
	for {
		if false {
			continue apply__2302590153
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(f_0, x_1)
	}
}

func Call_Data_Function_apply__1928284458(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
apply__1928284458:
	for {
		if false {
			continue apply__1928284458
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply(Get_Data_Either_Left(), x_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_apply__3425554218(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
apply__3425554218:
	for {
		if false {
			continue apply__3425554218
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply(Get_Data_Either_Left(), x_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_apply__1520674922(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
apply__1520674922:
	for {
		if false {
			continue apply__1520674922
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply(Get_Data_Either_Left(), x_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_apply__3344524906(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
apply__3344524906:
	for {
		if false {
			continue apply__3344524906
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply(Get_Data_Either_Left__3194200049(), x_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_apply__2793916074(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
apply__2793916074:
	for {
		if false {
			continue apply__2793916074
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply(Get_Data_Either_Left(), x_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_apply__1173000582(f_unused_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Date_Date) *Constructor_Data_Date_Date {
apply__1173000582:
	for {
		if false {
			continue apply__1173000582
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 *Constructor_Data_Date_Date = x_1_loop
		_ = x_1
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(Get_Partial_Unsafe_unsafePartial__3011796551(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(x_1)}))
	}
}

func Call_Data_Function_apply__428839251(f_unused_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
apply__428839251:
	for {
		if false {
			continue apply__428839251
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
		_ = x_unused_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply(Get_Data_Function_go__const(), gopurs_runtime.Apply(Get_Data_Either_Right__3870859976(), Get_Data_Unit_unit())); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_apply__569550801(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
apply__569550801:
	for {
		if false {
			continue apply__569550801
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply(Get_Data_Function_go__const(), x_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_apply__3998497865(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__3998497865:
	for {
		if false {
			continue apply__3998497865
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(f_0, x_1)
	}
}

func Call_Data_Function_apply__543526857(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__543526857:
	for {
		if false {
			continue apply__543526857
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(f_0, x_1)
	}
}

func Call_Data_Function_apply__2561392351(f_unused_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]]) gopurs_runtime.Value {
apply__2561392351:
	for {
		if false {
			continue apply__2561392351
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]] = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Data_Argonaut_Encode_Encoders_fromFoldable(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_3232828411_849153993(x_1))})
	}
}

func Call_Data_Function_apply__2998264905(f_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
apply__2998264905:
	for {
		if false {
			continue apply__2998264905
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = x_1_loop
		_ = x_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(x_1)}); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_apply__26935913(f_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]) *Constructor_Data_Time_Time {
apply__26935913:
	for {
		if false {
			continue apply__26935913
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] = x_1_loop
		_ = x_1
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_3839235747_3094389156(x_1))}))
	}
}

func Call_Data_Function_apply__3632476649(f_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
apply__3632476649:
	for {
		if false {
			continue apply__3632476649
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(x_1)})
	}
}

func Call_Data_Function_apply__972864507(f_unused_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
apply__972864507:
	for {
		if false {
			continue apply__972864507
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_Stream_pure3(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(x_1)})
	}
}

func Call_Data_Function_apply__2647861350(f_unused_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Maybe_Just[int64]) bool {
apply__2647861350:
	for {
		if false {
			continue apply__2647861350
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 *Constructor_Data_Maybe_Just[int64] = x_1_loop
		_ = x_1
		return (gopurs_runtime.Apply(Get_Data_Maybe_isJust__1937207052(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_1170268447_3094389156(x_1))}).IntVal) != (0)
	}
}

func Call_Data_Function_apply__2639368963(f_unused_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Maybe_Just[int64]) bool {
apply__2639368963:
	for {
		if false {
			continue apply__2639368963
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 *Constructor_Data_Maybe_Just[int64] = x_1_loop
		_ = x_1
		return (gopurs_runtime.Apply(Get_Data_Maybe_isNothing__1937207052(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_1170268447_3094389156(x_1))}).IntVal) != (0)
	}
}

func Call_Data_Function_apply__2730021546(f_unused_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
apply__2730021546:
	for {
		if false {
			continue apply__2730021546
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = x_1_loop
		_ = x_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply(Get_Data_Either_Left(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_3123684004_1293498952(x_1))}); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_apply__3653967818(f_unused_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
apply__3653967818:
	for {
		if false {
			continue apply__3653967818
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = x_1_loop
		_ = x_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply(Get_Data_Either_Left(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_3123684004_1293498952(x_1))}); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_apply__1272484046(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
apply__1272484046:
	for {
		if false {
			continue apply__1272484046
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_1170268447_3094389156(Rebox_Data_Function_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Nullable_toMaybe__1451648826(), x_1)))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: gopurs_runtime.Value{}, V1: false} }()
	}
}

func Call_Data_Function_apply__1141988375(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1141988375:
	for {
		if false {
			continue apply__1141988375
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Main_void(), x_1)
	}
}

func Call_Data_Function_apply__1383316684(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1383316684:
	for {
		if false {
			continue apply__1383316684
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Effect_Aff_launchAff__4085038308(), x_1)
	}
}

func Call_Data_Function_apply__1171255079(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1171255079:
	for {
		if false {
			continue apply__1171255079
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Effect_Aff_launchSuspendedAff__4085038308(), x_1)
	}
}

func Call_Data_Function_apply__2435095346(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2435095346:
	for {
		if false {
			continue apply__2435095346
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Effect_Aff_Compat_fromEffectFnAff__2688986814(), x_1)
	}
}

func Call_Data_Function_apply__734028734(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__734028734:
	for {
		if false {
			continue apply__734028734
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply2(Get_Node_Symbol_go__map(), Get_Data_Nullable_toMaybe(), x_1)
	}
}

func Call_Data_Function_apply__1422622261(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1422622261:
	for {
		if false {
			continue apply__1422622261
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Effect_Aff_void1(), x_1)
	}
}

func Call_Data_Function_apply__2813923095(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2813923095:
	for {
		if false {
			continue apply__2813923095
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_Stream_Aff_join(), x_1)
	}
}

func Call_Data_Function_apply__1413608268(f_0_loop gopurs_runtime.Value, x_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1413608268:
	for {
		if false {
			continue apply__1413608268
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_unused_1 gopurs_runtime.Value = x_unused_1_loop
		_ = x_unused_1
		return gopurs_runtime.Apply(f_0, Get_Node_Process_unsafeGetEnv())
	}
}

func Call_Data_Function_apply__1659262775(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1659262775:
	for {
		if false {
			continue apply__1659262775
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply2(Get_Node_Stream_Aff_go__map(), Get_Node_Stream_Aff_pure5(), x_1)
	}
}

func Call_Data_Function_apply__24380833(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__24380833:
	for {
		if false {
			continue apply__24380833
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_Stream_Aff_void(), x_1)
	}
}

func Call_Data_Function_apply__341522372(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__341522372:
	for {
		if false {
			continue apply__341522372
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Effect_Aff_void(), x_1)
	}
}

func Call_Data_Function_apply__556989731(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__556989731:
	for {
		if false {
			continue apply__556989731
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Main_liftEffect(), x_1)
	}
}

func Call_Data_Function_apply__483590224(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__483590224:
	for {
		if false {
			continue apply__483590224
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Effect_Aff_liftEffect(), x_1)
	}
}

func Call_Data_Function_apply__2279232018(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2279232018:
	for {
		if false {
			continue apply__2279232018
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Main_liftEffect1(), x_1)
	}
}

func Call_Data_Function_apply__1755036686(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1755036686:
	for {
		if false {
			continue apply__1755036686
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_EventEmitter_pure(), x_1)
	}
}

func Call_Data_Function_apply__1851282762(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 struct {
		buffers   []gopurs_runtime.Value
		readagain bool
		/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
	}
	V2 bool
} {
apply__1851282762:
	for {
		if false {
			continue apply__1851282762
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 struct {
				buffers   []gopurs_runtime.Value
				readagain bool
				/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
			}
			V2 bool
		} {
			_v := gopurs_runtime.Apply(Get_Data_Either_Left__2132796443(), x_1)
			if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
				return struct {
					V0 gopurs_runtime.Value
					V1 struct {
						buffers   []gopurs_runtime.Value
						readagain bool
						/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
					}
					V2 bool
				}{V0: gopurs_runtime.Value{}, V1: func() struct {
					buffers   []gopurs_runtime.Value
					readagain bool
					/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
				} {
					orig := (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0
					_ = orig
					clone := struct {
						buffers   []gopurs_runtime.Value
						readagain bool
						/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
					}{}
					clone.buffers = *((*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "buffers").UnsafePtr))
					clone.readagain = (gopurs_runtime.RecordGet(orig, "readagain").IntVal) != (0)
					return clone
				}(), V2: true}
			}
			return struct {
				V0 gopurs_runtime.Value
				V1 struct {
					buffers   []gopurs_runtime.Value
					readagain bool
					/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
				}
				V2 bool
			}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: struct {
				buffers   []gopurs_runtime.Value
				readagain bool
				/*gopurs:record-v1:2:0062007500660066006500720073,00720065006100640061006700610069006e*/
			}{}, V2: false}
		}()
	}
}

func Call_Data_Function_apply__2798404781(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2798404781:
	for {
		if false {
			continue apply__2798404781
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Main_throwError(), x_1)
	}
}

func Call_Data_Function_apply__1937891721(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1937891721:
	for {
		if false {
			continue apply__1937891721
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(f_0, x_1)
	}
}

func Call_Data_Function_apply__2514240489(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2514240489:
	for {
		if false {
			continue apply__2514240489
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(f_0, x_1)
	}
}

func Call_Data_Function_apply__1942606537(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1942606537:
	for {
		if false {
			continue apply__1942606537
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(f_0, x_1)
	}
}

func Call_Data_Function_apply__2422281551(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
apply__2422281551:
	for {
		if false {
			continue apply__2422281551
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return Rebox_Data_Function_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_NonEmpty_singleton__1307584590(), x_1)))
	}
}

func Call_Data_Function_apply__650006357(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__650006357:
	for {
		if false {
			continue apply__650006357
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Data_Argonaut_Core_fromObject(), x_1)
	}
}

func Call_Data_Function_apply__1319276646(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1319276646:
	for {
		if false {
			continue apply__1319276646
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_Buffer_unsafeThaw(), x_1)
	}
}

func Call_Data_Function_apply__2293662334(f_unused_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
apply__2293662334:
	for {
		if false {
			continue apply__2293662334
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 []gopurs_runtime.Value = x_1_loop
		_ = x_1
		return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray__4190607066(), gopurs_runtime.Array(x_1)).UnsafePtr))
	}
}

func Call_Data_Function_apply__862467209(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
apply__862467209:
	for {
		if false {
			continue apply__862467209
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 []gopurs_runtime.Value = x_1_loop
		_ = x_1
		return *((*[]gopurs_runtime.Value)(gopurs_runtime.Apply(f_0, gopurs_runtime.Array(x_1)).UnsafePtr))
	}
}

func Call_Data_Function_apply__1378432126(f_unused_0_loop gopurs_runtime.Value, x_1_loop []int64) []int64 {
apply__1378432126:
	for {
		if false {
			continue apply__1378432126
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 []int64 = x_1_loop
		_ = x_1
		return func() []int64 {
			arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unsafeFromArray__4020493786(), func() gopurs_runtime.Value {
				arr := x_1
				boxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					boxed[i] = gopurs_runtime.Int(v)
				}
				return gopurs_runtime.Array(boxed)
			}()).UnsafePtr)
			unboxed := make([]int64, len(arr))
			for i, v := range arr {
				unboxed[i] = v.IntVal
			}
			return unboxed
		}()
	}
}

func Call_Data_Function_apply__2114047305(f_0_loop gopurs_runtime.Value, x_1_loop bool) bool {
apply__2114047305:
	for {
		if false {
			continue apply__2114047305
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 bool = x_1_loop
		_ = x_1
		return (gopurs_runtime.Apply(f_0, gopurs_runtime.Bool(x_1)).IntVal) != (0)
	}
}

func Call_Data_Function_apply__676573833(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__676573833:
	for {
		if false {
			continue apply__676573833
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(f_0, x_1)
	}
}

func Call_Data_Function_apply__973594793(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__973594793:
	for {
		if false {
			continue apply__973594793
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(f_0, x_1)
	}
}

func Call_Data_Function_apply__3402160552(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__3402160552:
	for {
		if false {
			continue apply__3402160552
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Effect_Aff_Compat_pure(), x_1)
	}
}

func Call_Data_Function_apply__1639137429(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1639137429:
	for {
		if false {
			continue apply__1639137429
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_Stream_Aff_pure2(), x_1)
	}
}

func Call_Data_Function_apply__3034864212(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__3034864212:
	for {
		if false {
			continue apply__3034864212
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable__2525970890(), x_1)
	}
}

func Call_Data_Function_apply__2984437716(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2984437716:
	for {
		if false {
			continue apply__2984437716
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable__2892494410(), x_1)
	}
}

func Call_Data_Function_apply__1415232852(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__1415232852:
	for {
		if false {
			continue apply__1415232852
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable__3537276362(), x_1)
	}
}

func Call_Data_Function_apply__3334826457(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__3334826457:
	for {
		if false {
			continue apply__3334826457
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_Buffer_usingToImmutable__1694059786(), x_1)
	}
}

func Call_Data_Function_apply__3228670985(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) *Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
apply__3228670985:
	for {
		if false {
			continue apply__3228670985
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.CoerceToStruct[Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, x_1))
	}
}

func Call_Data_Function_apply__355458268(f_unused_0_loop gopurs_runtime.Value, x_1_loop float64) float64 {
apply__355458268:
	for {
		if false {
			continue apply__355458268
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 float64 = x_1_loop
		_ = x_1
		return x_1
	}
}

func Call_Data_Function_apply__2245704861(f_unused_0_loop gopurs_runtime.Value, x_1_loop float64) float64 {
apply__2245704861:
	for {
		if false {
			continue apply__2245704861
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 float64 = x_1_loop
		_ = x_1
		return x_1
	}
}

func Call_Data_Function_apply__2146671458(f_unused_0_loop gopurs_runtime.Value, x_1_loop struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
}) gopurs_runtime.Value {
apply__2146671458:
	for {
		if false {
			continue apply__2146671458
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 struct {
			isSuspended gopurs_runtime.Value
			join        gopurs_runtime.Value
			kill        gopurs_runtime.Value
			onComplete  gopurs_runtime.Value
			run         gopurs_runtime.Value
			/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
		} = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Effect_Aff_pure__621669172(), func() gopurs_runtime.Value {
			orig := x_1
			_ = orig
			return gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", orig.isSuspended, orig.join, orig.kill, orig.onComplete, orig.run)
		}())
	}
}

func Call_Data_Function_apply__190415331(f_unused_0_loop gopurs_runtime.Value, x_1_loop struct {
	body   string
	status int64
	/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
}) gopurs_runtime.Value {
apply__190415331:
	for {
		if false {
			continue apply__190415331
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 struct {
			body   string
			status int64
			/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
		} = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Main_pure2(), func() gopurs_runtime.Value {
			orig := x_1
			_ = orig
			return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
		}())
	}
}

func Call_Data_Function_apply__2201874889(f_0_loop gopurs_runtime.Value, x_1_loop struct {
	body   string
	status int64
	/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
}) gopurs_runtime.Value {
apply__2201874889:
	for {
		if false {
			continue apply__2201874889
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 struct {
			body   string
			status int64
			/*gopurs:record-v1:2:0062006f00640079,007300740061007400750073*/
		} = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(f_0, func() gopurs_runtime.Value {
			orig := x_1
			_ = orig
			return gopurs_runtime.RecordDict2("body", "status", gopurs_runtime.Str(orig.body), gopurs_runtime.Int(orig.status))
		}())
	}
}

func Call_Data_Function_apply__2515510025(f_0_loop gopurs_runtime.Value, x_1_loop string) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
apply__2515510025:
	for {
		if false {
			continue apply__2515510025
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 string = x_1_loop
		_ = x_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply(f_0, gopurs_runtime.Str(x_1)); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
			return struct {
				V0 gopurs_runtime.Value
				V1 gopurs_runtime.Value
				V2 bool
			}{V0: gopurs_runtime.Value{}, V1: (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V2: true}
		}; return struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: gopurs_runtime.Value{}, V2: false} }()
	}
}

func Call_Data_Function_apply__1333176700(f_unused_0_loop gopurs_runtime.Value, x_1_loop string) uint32 {
apply__1333176700:
	for {
		if false {
			continue apply__1333176700
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 string = x_1_loop
		_ = x_1
		return uint32(gopurs_runtime.Apply(Get_Partial_Unsafe_unsafeCrashWith(), gopurs_runtime.Str(x_1)).IntVal)
	}
}

func Call_Data_Function_apply__826080446(f_unused_0_loop gopurs_runtime.Value, x_1_loop string) gopurs_runtime.Value {
apply__826080446:
	for {
		if false {
			continue apply__826080446
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 string = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Effect_Exception_error(), gopurs_runtime.Str(x_1))
	}
}

func Call_Data_Function_apply__596732553(f_0_loop gopurs_runtime.Value, x_1_loop string) gopurs_runtime.Value {
apply__596732553:
	for {
		if false {
			continue apply__596732553
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 string = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(x_1))
	}
}

func Rebox_Data_Function_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Function_1293498952_3123684004(in *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Function_138441832_3103076375(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]{}
	out.V0 = in.V0.StrVal()
	out.V1 = in.V1
	return out
}

func Rebox_Data_Function_138441832_3132786365(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[uint32, float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[uint32, float64]{}
	out.V0 = uint32(in.V0.IntVal)
	out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_Function_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Function_2442833393_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_3132786365_138441832(in.V0))}
	out.V1 = Rebox_Data_Function_2442833393_849153993(in.V1)
	return out
}

func Rebox_Data_Function_2748095225_3887487416(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Function_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Function_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}

func Rebox_Data_Function_3094389156_3839235747(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](in.V0)
	return out
}

func Rebox_Data_Function_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_Function_3103076375_138441832(in *Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	out.V1 = in.V1
	return out
}

func Rebox_Data_Function_3123684004_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Function_3132786365_138441832(in *Constructor_Data_Tuple_Tuple[uint32, float64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Function_3232828411_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Function_3103076375_138441832(in.V0))}
	out.V1 = Rebox_Data_Function_3232828411_849153993(in.V1)
	return out
}

func Rebox_Data_Function_3839235747_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Function_3887487416_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Function_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}

func Rebox_Data_Function_849153993_2442833393(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{}
	out.V0 = Rebox_Data_Function_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	out.V1 = Rebox_Data_Function_849153993_2442833393(in.V1)
	return out
}

func Rebox_Data_Function_849153993_3232828411(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[string, gopurs_runtime.Value]]{}
	out.V0 = Rebox_Data_Function_138441832_3103076375(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	out.V1 = Rebox_Data_Function_849153993_3232828411(in.V1)
	return out
}
