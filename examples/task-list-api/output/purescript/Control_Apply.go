package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Apply_identity gopurs_runtime.Value
var once_Control_Apply_identity sync.Once

func Get_Control_Apply_identity() gopurs_runtime.Value {
	once_Control_Apply_identity.Do(func() {
		cache_Control_Apply_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Control_Apply_identity
}

var cache_Control_Apply_Apply_dollar_Dict gopurs_runtime.Value
var once_Control_Apply_Apply_dollar_Dict sync.Once

func Get_Control_Apply_Apply_dollar_Dict() gopurs_runtime.Value {
	once_Control_Apply_Apply_dollar_Dict.Do(func() {
		cache_Control_Apply_Apply_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Call_Control_Apply_Apply_dollar_Dict(func() struct {
				Functor0 gopurs_runtime.Value
				apply    gopurs_runtime.Value
				/*gopurs:record-v1:2:00460075006e00630074006f00720030,006100700070006c0079*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Functor0 gopurs_runtime.Value
					apply    gopurs_runtime.Value
					/*gopurs:record-v1:2:00460075006e00630074006f00720030,006100700070006c0079*/
				}{}
				clone.Functor0 = gopurs_runtime.RecordGet(orig, "Functor0")
				clone.apply = gopurs_runtime.RecordGet(orig, "apply")
				return clone
			}()))}
		})
	})
	return cache_Control_Apply_Apply_dollar_Dict
}

var cache_Control_Apply_applyProxy gopurs_runtime.Value
var once_Control_Apply_applyProxy sync.Once

func Get_Control_Apply_applyProxy() gopurs_runtime.Value {
	once_Control_Apply_applyProxy.Do(func() {
		cache_Control_Apply_applyProxy = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_1303115796_3741347833((&Constructor_Control_Apply_Apply[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_562069347_2812149806(Rebox_Control_Apply_2812149806_562069347(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorProxy()))))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
		})})))}
	})
	return cache_Control_Apply_applyProxy
}

var cache_Control_Apply_applyFn gopurs_runtime.Value
var once_Control_Apply_applyFn sync.Once

func Get_Control_Apply_applyFn() gopurs_runtime.Value {
	once_Control_Apply_applyFn.Do(func() {
		cache_Control_Apply_applyFn = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Functor_functorFn()
		}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
		})}))}
	})
	return cache_Control_Apply_applyFn
}

var cache_Control_Apply_applyArray gopurs_runtime.Value
var once_Control_Apply_applyArray sync.Once

func Get_Control_Apply_applyArray() gopurs_runtime.Value {
	once_Control_Apply_applyArray.Do(func() {
		cache_Control_Apply_applyArray = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}
		}), Get_Control_Apply_arrayApply()}))}
	})
	return cache_Control_Apply_applyArray
}

var cache_Control_Apply_apply gopurs_runtime.Value
var once_Control_Apply_apply sync.Once

func Get_Control_Apply_apply() gopurs_runtime.Value {
	once_Control_Apply_apply.Do(func() {
		cache_Control_Apply_apply = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Apply_apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Apply_apply
}

var cache_Control_Apply_apply__1827729951 gopurs_runtime.Value
var once_Control_Apply_apply__1827729951 sync.Once

func Get_Control_Apply_apply__1827729951() gopurs_runtime.Value {
	once_Control_Apply_apply__1827729951.Do(func() {
		cache_Control_Apply_apply__1827729951 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__1827729951(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box), Rebox_Control_Apply_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Apply_apply__1827729951
}

var cache_Control_Apply_apply__2439745503 gopurs_runtime.Value
var once_Control_Apply_apply__2439745503 sync.Once

func Get_Control_Apply_apply__2439745503() gopurs_runtime.Value {
	once_Control_Apply_apply__2439745503.Do(func() {
		cache_Control_Apply_apply__2439745503 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__2439745503(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box), Rebox_Control_Apply_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Apply_apply__2439745503
}

var cache_Control_Apply_apply__1976881535 gopurs_runtime.Value
var once_Control_Apply_apply__1976881535 sync.Once

func Get_Control_Apply_apply__1976881535() gopurs_runtime.Value {
	once_Control_Apply_apply__1976881535.Do(func() {
		cache_Control_Apply_apply__1976881535 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__1976881535(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box), Rebox_Control_Apply_3094389156_3839235747(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Apply_apply__1976881535
}

var cache_Control_Apply_apply__2672146207 gopurs_runtime.Value
var once_Control_Apply_apply__2672146207 sync.Once

func Get_Control_Apply_apply__2672146207() gopurs_runtime.Value {
	once_Control_Apply_apply__2672146207.Do(func() {
		cache_Control_Apply_apply__2672146207 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__2672146207(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box), Rebox_Control_Apply_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Apply_apply__2672146207
}

var cache_Control_Apply_apply__2106743103 gopurs_runtime.Value
var once_Control_Apply_apply__2106743103 sync.Once

func Get_Control_Apply_apply__2106743103() gopurs_runtime.Value {
	once_Control_Apply_apply__2106743103.Do(func() {
		cache_Control_Apply_apply__2106743103 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__2106743103(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box), Rebox_Control_Apply_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Apply_apply__2106743103
}

var cache_Control_Apply_apply__2598797535 gopurs_runtime.Value
var once_Control_Apply_apply__2598797535 sync.Once

func Get_Control_Apply_apply__2598797535() gopurs_runtime.Value {
	once_Control_Apply_apply__2598797535.Do(func() {
		cache_Control_Apply_apply__2598797535 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__2598797535(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box), Rebox_Control_Apply_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Apply_apply__2598797535
}

var cache_Control_Apply_apply__3976403967 gopurs_runtime.Value
var once_Control_Apply_apply__3976403967 sync.Once

func Get_Control_Apply_apply__3976403967() gopurs_runtime.Value {
	once_Control_Apply_apply__3976403967.Do(func() {
		cache_Control_Apply_apply__3976403967 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__3976403967(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box), Rebox_Control_Apply_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Apply_apply__3976403967
}

var cache_Control_Apply_apply__708417983 gopurs_runtime.Value
var once_Control_Apply_apply__708417983 sync.Once

func Get_Control_Apply_apply__708417983() gopurs_runtime.Value {
	once_Control_Apply_apply__708417983.Do(func() {
		cache_Control_Apply_apply__708417983 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Apply_apply__708417983(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box), Rebox_Control_Apply_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Apply_apply__708417983
}

var cache_Control_Apply_apply__3879048993 gopurs_runtime.Value
var once_Control_Apply_apply__3879048993 sync.Once

func Get_Control_Apply_apply__3879048993() gopurs_runtime.Value {
	once_Control_Apply_apply__3879048993.Do(func() {
		cache_Control_Apply_apply__3879048993 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Apply_apply__3879048993(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Apply_apply__3879048993
}

var cache_Control_Apply_apply__2491776737 gopurs_runtime.Value
var once_Control_Apply_apply__2491776737 sync.Once

func Get_Control_Apply_apply__2491776737() gopurs_runtime.Value {
	once_Control_Apply_apply__2491776737.Do(func() {
		cache_Control_Apply_apply__2491776737 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Apply_apply__2491776737(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Apply_apply__2491776737
}

var cache_Control_Apply_applyFirst gopurs_runtime.Value
var once_Control_Apply_applyFirst sync.Once

func Get_Control_Apply_applyFirst() gopurs_runtime.Value {
	once_Control_Apply_applyFirst.Do(func() {
		cache_Control_Apply_applyFirst = gopurs_runtime.Func3(func(dictApply_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Apply_applyFirst(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box), a_1_box, b_2_box)
		})
	})
	return cache_Control_Apply_applyFirst
}

var cache_Control_Apply_applySecond gopurs_runtime.Value
var once_Control_Apply_applySecond sync.Once

func Get_Control_Apply_applySecond() gopurs_runtime.Value {
	once_Control_Apply_applySecond.Do(func() {
		cache_Control_Apply_applySecond = gopurs_runtime.Func3(func(dictApply_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Apply_applySecond(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box), a_1_box, b_2_box)
		})
	})
	return cache_Control_Apply_applySecond
}

var cache_Control_Apply_applySecond__3440890937 gopurs_runtime.Value
var once_Control_Apply_applySecond__3440890937 sync.Once

func Get_Control_Apply_applySecond__3440890937() gopurs_runtime.Value {
	once_Control_Apply_applySecond__3440890937.Do(func() {
		cache_Control_Apply_applySecond__3440890937 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Apply_applySecond__3440890937(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Apply_applySecond__3440890937
}

var cache_Control_Apply_applySecond__3021785433 gopurs_runtime.Value
var once_Control_Apply_applySecond__3021785433 sync.Once

func Get_Control_Apply_applySecond__3021785433() gopurs_runtime.Value {
	once_Control_Apply_applySecond__3021785433.Do(func() {
		cache_Control_Apply_applySecond__3021785433 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Apply_applySecond__3021785433(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Apply_applySecond__3021785433
}

var cache_Control_Apply_lift2 gopurs_runtime.Value
var once_Control_Apply_lift2 sync.Once

func Get_Control_Apply_lift2() gopurs_runtime.Value {
	once_Control_Apply_lift2.Do(func() {
		cache_Control_Apply_lift2 = gopurs_runtime.Func4(func(dictApply_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Apply_lift2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box), f_1_box, a_2_box, b_3_box)
		})
	})
	return cache_Control_Apply_lift2
}

var cache_Control_Apply_lift3 gopurs_runtime.Value
var once_Control_Apply_lift3 sync.Once

func Get_Control_Apply_lift3() gopurs_runtime.Value {
	once_Control_Apply_lift3.Do(func() {
		cache_Control_Apply_lift3 = gopurs_runtime.Func5(func(dictApply_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Apply_lift3(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box), f_1_box, a_2_box, b_3_box, c_4_box)
		})
	})
	return cache_Control_Apply_lift3
}

var cache_Control_Apply_lift4 gopurs_runtime.Value
var once_Control_Apply_lift4 sync.Once

func Get_Control_Apply_lift4() gopurs_runtime.Value {
	once_Control_Apply_lift4.Do(func() {
		cache_Control_Apply_lift4 = gopurs_runtime.Func6(func(dictApply_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value, d_5_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Apply_lift4(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box), f_1_box, a_2_box, b_3_box, c_4_box, d_5_box)
		})
	})
	return cache_Control_Apply_lift4
}

var cache_Control_Apply_lift5 gopurs_runtime.Value
var once_Control_Apply_lift5 sync.Once

func Get_Control_Apply_lift5() gopurs_runtime.Value {
	once_Control_Apply_lift5.Do(func() {
		cache_Control_Apply_lift5 = gopurs_runtime.Func7(func(dictApply_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value, d_5_box gopurs_runtime.Value, e_6_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Apply_lift5(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0_box), f_1_box, a_2_box, b_3_box, c_4_box, d_5_box, e_6_box)
		})
	})
	return cache_Control_Apply_lift5
}

type Constructor_Control_Apply_Apply[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f004100700070006c0079:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f004100700070006c0079005f004100700070006c0079
func init() {
	gopurs_runtime.StructGetters[3032403085] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Apply_Apply[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Functor0":
			return gopurs_runtime.Box(c.V0)
		case "apply":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Apply_Apply: " + key)
		}
	}
}

func Call_Control_Apply_Apply_dollar_Dict(x_0_loop struct {
	Functor0 gopurs_runtime.Value
	apply    gopurs_runtime.Value
	/*gopurs:record-v1:2:00460075006e00630074006f00720030,006100700070006c0079*/
}) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	var x_0 struct {
		Functor0 gopurs_runtime.Value
		apply    gopurs_runtime.Value
		/*gopurs:record-v1:2:00460075006e00630074006f00720030,006100700070006c0079*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Functor0", "apply", orig.Functor0, orig.apply)
	}())
}

func Call_Control_Apply_apply(dict_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Control_Apply_apply__1827729951(__eta0_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta1_1_loop *Constructor_Data_Maybe_Just[uint32]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
apply__1827729951:
	for {
		if false {
			continue apply__1827729951
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[uint32] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta0_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_622082505_3094389156(__eta1_1))})))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Apply_apply__2439745503(__eta0_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta1_1_loop *Constructor_Data_Maybe_Just[uint32]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
apply__2439745503:
	for {
		if false {
			continue apply__2439745503
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[uint32] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta0_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_622082505_3094389156(__eta1_1))})))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Apply_apply__1976881535(__eta0_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta1_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
apply__1976881535:
	for {
		if false {
			continue apply__1976881535
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_4151030883_3094389156(Rebox_Control_Apply_3094389156_4151030883(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta0_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_3839235747_3094389156(__eta1_1))})))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Apply_apply__2672146207(__eta0_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta1_1_loop *Constructor_Data_Maybe_Just[int64]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
apply__2672146207:
	for {
		if false {
			continue apply__2672146207
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[int64] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_2280409795_3094389156(Rebox_Control_Apply_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta0_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_1170268447_3094389156(__eta1_1))})))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Apply_apply__2106743103(__eta0_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta1_1_loop *Constructor_Data_Maybe_Just[int64]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
apply__2106743103:
	for {
		if false {
			continue apply__2106743103
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[int64] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_3516139266_3094389156(Rebox_Control_Apply_3094389156_3516139266(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta0_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_1170268447_3094389156(__eta1_1))})))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Apply_apply__2598797535(__eta0_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta1_1_loop *Constructor_Data_Maybe_Just[int64]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
apply__2598797535:
	for {
		if false {
			continue apply__2598797535
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[int64] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_3839235747_3094389156(Rebox_Control_Apply_3094389156_3839235747(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta0_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_1170268447_3094389156(__eta1_1))})))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Apply_apply__3976403967(__eta0_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta1_1_loop *Constructor_Data_Maybe_Just[int64]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
apply__3976403967:
	for {
		if false {
			continue apply__3976403967
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[int64] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta0_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_1170268447_3094389156(__eta1_1))})))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Apply_apply__708417983(__eta0_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], __eta1_1_loop *Constructor_Data_Maybe_Just[int64]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
apply__708417983:
	for {
		if false {
			continue apply__708417983
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[int64] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Apply_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta0_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_1170268447_3094389156(__eta1_1))})))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Apply_apply__3879048993(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__3879048993:
	for {
		if false {
			continue apply__3879048993
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Apply_apply__2491776737(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
apply__2491776737:
	for {
		if false {
			continue apply__2491776737
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Apply_applyFirst(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
	_ = dictApply_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}), "map"), Get_Data_Function_go__const(), a_1), b_2)
}

func Call_Control_Apply_applySecond(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
	_ = dictApply_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
	__local_var_3_0 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	_ = __local_var_3_0
	return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return __local_var_3_0
	}), a_1), b_2)
}

func Call_Control_Apply_applySecond__3440890937(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
applySecond__3440890937:
	for {
		if false {
			continue applySecond__3440890937
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()).V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()).V0, gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(Get_Data_Function_const__2389296201(), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})), __eta0_0), __eta1_1)
	}
}

func Call_Control_Apply_applySecond__3021785433(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
applySecond__3021785433:
	for {
		if false {
			continue applySecond__3021785433
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()).V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()).V0, gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(Get_Data_Function_const__3353209993(), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})), __eta0_0), __eta1_1)
	}
}

func Call_Control_Apply_lift2(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
	_ = dictApply_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	var b_3 gopurs_runtime.Value = b_3_loop
	_ = b_3
	return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}), "map"), f_1, a_2), b_3)
}

func Call_Control_Apply_lift3(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value, c_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
	_ = dictApply_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	var b_3 gopurs_runtime.Value = b_3_loop
	_ = b_3
	var c_4 gopurs_runtime.Value = c_4_loop
	_ = c_4
	return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}), "map"), f_1, a_2), b_3), c_4)
}

func Call_Control_Apply_lift4(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value, c_4_loop gopurs_runtime.Value, d_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
	_ = dictApply_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	var b_3 gopurs_runtime.Value = b_3_loop
	_ = b_3
	var c_4 gopurs_runtime.Value = c_4_loop
	_ = c_4
	var d_5 gopurs_runtime.Value = d_5_loop
	_ = d_5
	return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}), "map"), f_1, a_2), b_3), c_4), d_5)
}

func Call_Control_Apply_lift5(dictApply_0_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value, c_4_loop gopurs_runtime.Value, d_5_loop gopurs_runtime.Value, e_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApply_0 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_0_loop
	_ = dictApply_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	var b_3 gopurs_runtime.Value = b_3_loop
	_ = b_3
	var c_4 gopurs_runtime.Value = c_4_loop
	_ = c_4
	var d_5 gopurs_runtime.Value = d_5_loop
	_ = d_5
	var e_6 gopurs_runtime.Value = e_6_loop
	_ = e_6
	return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}), "map"), f_1, a_2), b_3), c_4), d_5), e_6)
}

func Rebox_Control_Apply_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Control_Apply_1303115796_3741347833(in *Constructor_Control_Apply_Apply[uint32]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Apply_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Apply_2812149806_562069347(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[uint32])(unsafe.Pointer(in))
}

func Rebox_Control_Apply_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Control_Apply_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}

func Rebox_Control_Apply_3094389156_3516139266(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]]{}
	out.V0 = Rebox_Control_Apply_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Control_Apply_3094389156_3839235747(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](in.V0)
	return out
}

func Rebox_Control_Apply_3094389156_4151030883(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_DateTime_DateTime] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_DateTime_DateTime]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](in.V0)
	return out
}

func Rebox_Control_Apply_3094389156_622082505(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[uint32] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[uint32]{}
	out.V0 = uint32(in.V0.IntVal)
	return out
}

func Rebox_Control_Apply_3516139266_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Apply_2280409795_3094389156(in.V0))}
	return out
}

func Rebox_Control_Apply_3741347833_3552963512(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Apply_3839235747_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Apply_4151030883_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_DateTime_DateTime]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Apply_562069347_2812149806(in *Constructor_Data_Functor_Functor[uint32]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Apply_622082505_3094389156(in *Constructor_Data_Maybe_Just[uint32]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Get_Control_Apply_arrayApply() gopurs_runtime.Value {
	return _Gopurs_Control_Apply_ArrayApply
}
