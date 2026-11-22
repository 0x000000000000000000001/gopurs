package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Applicative_Applicative_dollar_Dict gopurs_runtime.Value
var once_Control_Applicative_Applicative_dollar_Dict sync.Once

func Get_Control_Applicative_Applicative_dollar_Dict() gopurs_runtime.Value {
	once_Control_Applicative_Applicative_dollar_Dict.Do(func() {
		cache_Control_Applicative_Applicative_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Call_Control_Applicative_Applicative_dollar_Dict(func() struct {
				Apply0 gopurs_runtime.Value
				pure   gopurs_runtime.Value
				/*gopurs:record-v1:2:004100700070006c00790030,0070007500720065*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Apply0 gopurs_runtime.Value
					pure   gopurs_runtime.Value
					/*gopurs:record-v1:2:004100700070006c00790030,0070007500720065*/
				}{}
				clone.Apply0 = gopurs_runtime.RecordGet(orig, "Apply0")
				clone.pure = gopurs_runtime.RecordGet(orig, "pure")
				return clone
			}()))}
		})
	})
	return cache_Control_Applicative_Applicative_dollar_Dict
}

var cache_Control_Applicative_pure gopurs_runtime.Value
var once_Control_Applicative_pure sync.Once

func Get_Control_Applicative_pure() gopurs_runtime.Value {
	once_Control_Applicative_pure.Do(func() {
		cache_Control_Applicative_pure = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Applicative_pure
}

var cache_Control_Applicative_pure__2387439490 gopurs_runtime.Value
var once_Control_Applicative_pure__2387439490 sync.Once

func Get_Control_Applicative_pure__2387439490() gopurs_runtime.Value {
	once_Control_Applicative_pure__2387439490.Do(func() {
		cache_Control_Applicative_pure__2387439490 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Applicative_pure__2387439490(__eta0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Control_Applicative_pure__2387439490
}

var cache_Control_Applicative_pure__3632964551 gopurs_runtime.Value
var once_Control_Applicative_pure__3632964551 sync.Once

func Get_Control_Applicative_pure__3632964551() gopurs_runtime.Value {
	once_Control_Applicative_pure__3632964551.Do(func() {
		cache_Control_Applicative_pure__3632964551 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Control_Applicative_pure__3632964551(__eta0_0_box))}
		})
	})
	return cache_Control_Applicative_pure__3632964551
}

var cache_Control_Applicative_pure__1137267138 gopurs_runtime.Value
var once_Control_Applicative_pure__1137267138 sync.Once

func Get_Control_Applicative_pure__1137267138() gopurs_runtime.Value {
	once_Control_Applicative_pure__1137267138.Do(func() {
		cache_Control_Applicative_pure__1137267138 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Applicative_pure__1137267138(uint32(__eta0_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Applicative_pure__1137267138
}

var cache_Control_Applicative_pure__3917743074 gopurs_runtime.Value
var once_Control_Applicative_pure__3917743074 sync.Once

func Get_Control_Applicative_pure__3917743074() gopurs_runtime.Value {
	once_Control_Applicative_pure__3917743074.Do(func() {
		cache_Control_Applicative_pure__3917743074 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Applicative_pure__3917743074(__eta0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Applicative_pure__3917743074
}

var cache_Control_Applicative_pure__1963495204 gopurs_runtime.Value
var once_Control_Applicative_pure__1963495204 sync.Once

func Get_Control_Applicative_pure__1963495204() gopurs_runtime.Value {
	once_Control_Applicative_pure__1963495204.Do(func() {
		cache_Control_Applicative_pure__1963495204 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Applicative_pure__1963495204(__eta0_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Applicative_pure__1963495204
}

var cache_Control_Applicative_pure__3763014981 gopurs_runtime.Value
var once_Control_Applicative_pure__3763014981 sync.Once

func Get_Control_Applicative_pure__3763014981() gopurs_runtime.Value {
	once_Control_Applicative_pure__3763014981.Do(func() {
		cache_Control_Applicative_pure__3763014981 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_pure__3763014981(func() struct {
				completed bool
				id        string
				title     string
				/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
			} {
				orig := __eta0_0_box
				_ = orig
				clone := struct {
					completed bool
					id        string
					title     string
					/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
				}{}
				clone.completed = (gopurs_runtime.RecordGet(orig, "completed").IntVal) != (0)
				clone.id = gopurs_runtime.RecordGet(orig, "id").StrVal()
				clone.title = gopurs_runtime.RecordGet(orig, "title").StrVal()
				return clone
			}())
		})
	})
	return cache_Control_Applicative_pure__3763014981
}

var cache_Control_Applicative_pure__2964015333 gopurs_runtime.Value
var once_Control_Applicative_pure__2964015333 sync.Once

func Get_Control_Applicative_pure__2964015333() gopurs_runtime.Value {
	once_Control_Applicative_pure__2964015333.Do(func() {
		cache_Control_Applicative_pure__2964015333 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_pure__2964015333(__eta0_0_box)
		})
	})
	return cache_Control_Applicative_pure__2964015333
}

var cache_Control_Applicative_pure__738246428 gopurs_runtime.Value
var once_Control_Applicative_pure__738246428 sync.Once

func Get_Control_Applicative_pure__738246428() gopurs_runtime.Value {
	once_Control_Applicative_pure__738246428.Do(func() {
		cache_Control_Applicative_pure__738246428 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_pure__738246428(__eta0_0_box)
		})
	})
	return cache_Control_Applicative_pure__738246428
}

var cache_Control_Applicative_pure__3252509724 gopurs_runtime.Value
var once_Control_Applicative_pure__3252509724 sync.Once

func Get_Control_Applicative_pure__3252509724() gopurs_runtime.Value {
	once_Control_Applicative_pure__3252509724.Do(func() {
		cache_Control_Applicative_pure__3252509724 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_pure__3252509724(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box))
		})
	})
	return cache_Control_Applicative_pure__3252509724
}

var cache_Control_Applicative_pure__1238803548 gopurs_runtime.Value
var once_Control_Applicative_pure__1238803548 sync.Once

func Get_Control_Applicative_pure__1238803548() gopurs_runtime.Value {
	once_Control_Applicative_pure__1238803548.Do(func() {
		cache_Control_Applicative_pure__1238803548 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_pure__1238803548(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box))
		})
	})
	return cache_Control_Applicative_pure__1238803548
}

var cache_Control_Applicative_pure__1139843612 gopurs_runtime.Value
var once_Control_Applicative_pure__1139843612 sync.Once

func Get_Control_Applicative_pure__1139843612() gopurs_runtime.Value {
	once_Control_Applicative_pure__1139843612.Do(func() {
		cache_Control_Applicative_pure__1139843612 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_pure__1139843612(Rebox_Control_Applicative_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box)))
		})
	})
	return cache_Control_Applicative_pure__1139843612
}

var cache_Control_Applicative_pure__1379658716 gopurs_runtime.Value
var once_Control_Applicative_pure__1379658716 sync.Once

func Get_Control_Applicative_pure__1379658716() gopurs_runtime.Value {
	once_Control_Applicative_pure__1379658716.Do(func() {
		cache_Control_Applicative_pure__1379658716 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_pure__1379658716((__eta0_0_box.IntVal) != (0))
		})
	})
	return cache_Control_Applicative_pure__1379658716
}

var cache_Control_Applicative_pure__3156000668 gopurs_runtime.Value
var once_Control_Applicative_pure__3156000668 sync.Once

func Get_Control_Applicative_pure__3156000668() gopurs_runtime.Value {
	once_Control_Applicative_pure__3156000668.Do(func() {
		cache_Control_Applicative_pure__3156000668 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_pure__3156000668(__eta0_0_box)
		})
	})
	return cache_Control_Applicative_pure__3156000668
}

var cache_Control_Applicative_pure__3258610304 gopurs_runtime.Value
var once_Control_Applicative_pure__3258610304 sync.Once

func Get_Control_Applicative_pure__3258610304() gopurs_runtime.Value {
	once_Control_Applicative_pure__3258610304.Do(func() {
		cache_Control_Applicative_pure__3258610304 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_pure__3258610304(func() struct {
				isSuspended gopurs_runtime.Value
				join        gopurs_runtime.Value
				kill        gopurs_runtime.Value
				onComplete  gopurs_runtime.Value
				run         gopurs_runtime.Value
				/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
			} {
				orig := __eta_norm_0_0_box
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
	return cache_Control_Applicative_pure__3258610304
}

var cache_Control_Applicative_pure__2088919292 gopurs_runtime.Value
var once_Control_Applicative_pure__2088919292 sync.Once

func Get_Control_Applicative_pure__2088919292() gopurs_runtime.Value {
	once_Control_Applicative_pure__2088919292.Do(func() {
		cache_Control_Applicative_pure__2088919292 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_pure__2088919292(__eta0_0_box)
		})
	})
	return cache_Control_Applicative_pure__2088919292
}

var cache_Control_Applicative_pure__3641479109 gopurs_runtime.Value
var once_Control_Applicative_pure__3641479109 sync.Once

func Get_Control_Applicative_pure__3641479109() gopurs_runtime.Value {
	once_Control_Applicative_pure__3641479109.Do(func() {
		cache_Control_Applicative_pure__3641479109 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_pure__3641479109(__eta0_0_box)
		})
	})
	return cache_Control_Applicative_pure__3641479109
}

var cache_Control_Applicative_unless gopurs_runtime.Value
var once_Control_Applicative_unless sync.Once

func Get_Control_Applicative_unless() gopurs_runtime.Value {
	once_Control_Applicative_unless.Do(func() {
		cache_Control_Applicative_unless = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_unless(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
		})
	})
	return cache_Control_Applicative_unless
}

var cache_Control_Applicative_when gopurs_runtime.Value
var once_Control_Applicative_when sync.Once

func Get_Control_Applicative_when() gopurs_runtime.Value {
	once_Control_Applicative_when.Do(func() {
		cache_Control_Applicative_when = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_when(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
		})
	})
	return cache_Control_Applicative_when
}

var cache_Control_Applicative_liftA1 gopurs_runtime.Value
var once_Control_Applicative_liftA1 sync.Once

func Get_Control_Applicative_liftA1() gopurs_runtime.Value {
	once_Control_Applicative_liftA1.Do(func() {
		cache_Control_Applicative_liftA1 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), f_1_box, a_2_box)
		})
	})
	return cache_Control_Applicative_liftA1
}

var cache_Control_Applicative_applicativeProxy gopurs_runtime.Value
var once_Control_Applicative_applicativeProxy sync.Once

func Get_Control_Applicative_applicativeProxy() gopurs_runtime.Value {
	once_Control_Applicative_applicativeProxy.Do(func() {
		cache_Control_Applicative_applicativeProxy = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Control_Applicative_3924528404_1439734649((&Constructor_Control_Applicative_Applicative[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Control_Applicative_1303115796_3741347833(Rebox_Control_Applicative_3741347833_1303115796(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Control_Apply_applyProxy()))))}
		}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
		})})))}
	})
	return cache_Control_Applicative_applicativeProxy
}

var cache_Control_Applicative_applicativeFn gopurs_runtime.Value
var once_Control_Applicative_applicativeFn sync.Once

func Get_Control_Applicative_applicativeFn() gopurs_runtime.Value {
	once_Control_Applicative_applicativeFn.Do(func() {
		cache_Control_Applicative_applicativeFn = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Apply_applyFn()
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Control_Applicative_applicativeFn
}

var cache_Control_Applicative_applicativeArray gopurs_runtime.Value
var once_Control_Applicative_applicativeArray sync.Once

func Get_Control_Applicative_applicativeArray() gopurs_runtime.Value {
	once_Control_Applicative_applicativeArray.Do(func() {
		cache_Control_Applicative_applicativeArray = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Control_Apply_applyArray()))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array([]gopurs_runtime.Value{x_0})
		})}))}
	})
	return cache_Control_Applicative_applicativeArray
}

type Constructor_Control_Applicative_Applicative[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f004100700070006c0069006300610074006900760065:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f004100700070006c0069006300610074006900760065005f004100700070006c0069006300610074006900760065
func init() {
	gopurs_runtime.StructGetters[1459134221] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Applicative_Applicative[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Apply0":
			return gopurs_runtime.Box(c.V0)
		case "pure":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Applicative_Applicative: " + key)
		}
	}
}

func Call_Control_Applicative_Applicative_dollar_Dict(x_0_loop struct {
	Apply0 gopurs_runtime.Value
	pure   gopurs_runtime.Value
	/*gopurs:record-v1:2:004100700070006c00790030,0070007500720065*/
}) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	var x_0 struct {
		Apply0 gopurs_runtime.Value
		pure   gopurs_runtime.Value
		/*gopurs:record-v1:2:004100700070006c00790030,0070007500720065*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Apply0", "pure", orig.Apply0, orig.pure)
	}())
}

func Call_Control_Applicative_pure(dict_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Control_Applicative_pure__2387439490(__eta0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
pure__2387439490:
	for {
		if false {
			continue pure__2387439490
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_applicativeEither(), "pure"), __eta0_0); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Control_Applicative_pure__3632964551(__eta0_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
pure__3632964551:
	for {
		if false {
			continue pure__3632964551
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Control_Applicative_1439734649_2870828117(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Types_applicativeList())).V1, __eta0_0))
	}
}

func Call_Control_Applicative_pure__1137267138(__eta0_0_loop uint32) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
pure__1137267138:
	for {
		if false {
			continue pure__1137267138
		}
		var __eta0_0 uint32 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Applicative_622082505_3094389156(Rebox_Control_Applicative_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Control_Applicative_1439734649_649684152(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Maybe_applicativeMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: int64(__eta0_0), UnsafePtr: nil})))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Applicative_pure__3917743074(__eta0_0_loop int64) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
pure__3917743074:
	for {
		if false {
			continue pure__3917743074
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Applicative_1170268447_3094389156(Rebox_Control_Applicative_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Control_Applicative_1439734649_649684152(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Maybe_applicativeMaybe())).V1, gopurs_runtime.Int(__eta0_0))))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Applicative_pure__1963495204(__eta0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
pure__1963495204:
	for {
		if false {
			continue pure__1963495204
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Rebox_Control_Applicative_2307501113_673668088(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Data_Maybe_alternativeMaybe())).V0, gopurs_runtime.Value{}), "pure"), __eta0_0)))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Applicative_pure__3763014981(__eta0_0_loop struct {
	completed bool
	id        string
	title     string
	/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
}) gopurs_runtime.Value {
pure__3763014981:
	for {
		if false {
			continue pure__3763014981
		}
		var __eta0_0 struct {
			completed bool
			id        string
			title     string
			/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
		} = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeAff()).V1, func() gopurs_runtime.Value {
			orig := __eta0_0
			_ = orig
			return gopurs_runtime.RecordDict3("completed", "id", "title", gopurs_runtime.Bool(orig.completed), gopurs_runtime.Str(orig.id), gopurs_runtime.Str(orig.title))
		}())
	}
}

func Call_Control_Applicative_pure__2964015333(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
pure__2964015333:
	for {
		if false {
			continue pure__2964015333
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_Aff_applicativeAff()).V1, __eta0_0)
	}
}

func Call_Control_Applicative_pure__738246428(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
pure__738246428:
	for {
		if false {
			continue pure__738246428
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, __eta0_0)
	}
}

func Call_Control_Applicative_pure__3252509724(__eta0_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
pure__3252509724:
	for {
		if false {
			continue pure__3252509724
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta0_0)})
	}
}

func Call_Control_Applicative_pure__1238803548(__eta0_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
pure__1238803548:
	for {
		if false {
			continue pure__1238803548
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__eta0_0)})
	}
}

func Call_Control_Applicative_pure__1139843612(__eta0_0_loop *Constructor_Data_Maybe_Just[string]) gopurs_runtime.Value {
pure__1139843612:
	for {
		if false {
			continue pure__1139843612
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[string] = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Applicative_742090555_3094389156(__eta0_0))})
	}
}

func Call_Control_Applicative_pure__1379658716(__eta0_0_loop bool) gopurs_runtime.Value {
pure__1379658716:
	for {
		if false {
			continue pure__1379658716
		}
		var __eta0_0 bool = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, gopurs_runtime.Bool(__eta0_0))
	}
}

func Call_Control_Applicative_pure__3156000668(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
pure__3156000668:
	for {
		if false {
			continue pure__3156000668
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, __eta0_0)
	}
}

func Call_Control_Applicative_pure__3258610304(__eta_norm_0_0_loop struct {
	isSuspended gopurs_runtime.Value
	join        gopurs_runtime.Value
	kill        gopurs_runtime.Value
	onComplete  gopurs_runtime.Value
	run         gopurs_runtime.Value
	/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
}) gopurs_runtime.Value {
pure__3258610304:
	for {
		if false {
			continue pure__3258610304
		}
		var __eta_norm_0_0 struct {
			isSuspended gopurs_runtime.Value
			join        gopurs_runtime.Value
			kill        gopurs_runtime.Value
			onComplete  gopurs_runtime.Value
			run         gopurs_runtime.Value
			/*gopurs:record-v1:5:0069007300530075007300700065006e006400650064,006a006f0069006e,006b0069006c006c,006f006e0043006f006d0070006c006500740065,00720075006e*/
		} = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, func() gopurs_runtime.Value {
			orig := __eta_norm_0_0
			_ = orig
			return gopurs_runtime.RecordDict5("isSuspended", "join", "kill", "onComplete", "run", orig.isSuspended, orig.join, orig.kill, orig.onComplete, orig.run)
		}())
	}
}

func Call_Control_Applicative_pure__2088919292(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
pure__2088919292:
	for {
		if false {
			continue pure__2088919292
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()).V1, __eta0_0)
	}
}

func Call_Control_Applicative_pure__3641479109(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
pure__3641479109:
	for {
		if false {
			continue pure__3641479109
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Control_Alternative_alternativeArray()).V0, gopurs_runtime.Value{}), "pure"), __eta0_0)
	}
}

func Call_Control_Applicative_unless(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	var v_1 bool = v_1_loop
	_ = v_1
	var v1_2 gopurs_runtime.Value = v1_2_loop
	_ = v1_2
	var __t0 gopurs_runtime.Value
	{
		if (v_1) != (true) {
			__t0 = v1_2
			goto end_branch_0
		} else {

		}
	}
	{
		if v_1 {
			__t0 = gopurs_runtime.Apply(dictApplicative_0.V1, Get_Data_Unit_unit())
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Control_Applicative_when(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	var v_1 bool = v_1_loop
	_ = v_1
	var v1_2 gopurs_runtime.Value = v1_2_loop
	_ = v1_2
	var __t0 gopurs_runtime.Value
	{
		if v_1 {
			__t0 = v1_2
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = gopurs_runtime.Apply(dictApplicative_0.V1, Get_Data_Unit_unit())
	}
end_branch_0:
	return __t0
}

func Call_Control_Applicative_liftA1(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
	_ = dictApplicative_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var a_2 gopurs_runtime.Value = a_2_loop
	_ = a_2
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(dictApplicative_0.V1, f_1), a_2)
}

func Rebox_Control_Applicative_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Control_Applicative_1303115796_3741347833(in *Constructor_Control_Apply_Apply[uint32]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Applicative_1439734649_2870828117(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Applicative_Applicative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Applicative_1439734649_649684152(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Applicative_2307501113_673668088(in *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) *Constructor_Control_Alternative_Alternative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Alternative_Alternative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Applicative_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Control_Applicative_3094389156_622082505(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[uint32] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[uint32]{}
	out.V0 = uint32(in.V0.IntVal)
	return out
}

func Rebox_Control_Applicative_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Control_Applicative_3741347833_1303115796(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[uint32])(unsafe.Pointer(in))
}

func Rebox_Control_Applicative_3924528404_1439734649(in *Constructor_Control_Applicative_Applicative[uint32]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Applicative_Applicative[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Applicative_622082505_3094389156(in *Constructor_Data_Maybe_Just[uint32]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Control_Applicative_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}
