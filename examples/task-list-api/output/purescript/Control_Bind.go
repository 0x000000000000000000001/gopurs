package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Bind_identity gopurs_runtime.Value
var once_Control_Bind_identity sync.Once

func Get_Control_Bind_identity() gopurs_runtime.Value {
	once_Control_Bind_identity.Do(func() {
		cache_Control_Bind_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_identity(x_0_box)
		})
	})
	return cache_Control_Bind_identity
}

var cache_Control_Bind_Bind_dollar_Dict gopurs_runtime.Value
var once_Control_Bind_Bind_dollar_Dict sync.Once

func Get_Control_Bind_Bind_dollar_Dict() gopurs_runtime.Value {
	once_Control_Bind_Bind_dollar_Dict.Do(func() {
		cache_Control_Bind_Bind_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Call_Control_Bind_Bind_dollar_Dict(func() struct {
				Apply0 gopurs_runtime.Value
				bind   gopurs_runtime.Value
				/*gopurs:record-v1:2:004100700070006c00790030,00620069006e0064*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Apply0 gopurs_runtime.Value
					bind   gopurs_runtime.Value
					/*gopurs:record-v1:2:004100700070006c00790030,00620069006e0064*/
				}{}
				clone.Apply0 = gopurs_runtime.RecordGet(orig, "Apply0")
				clone.bind = gopurs_runtime.RecordGet(orig, "bind")
				return clone
			}()))}
		})
	})
	return cache_Control_Bind_Bind_dollar_Dict
}

var cache_Control_Bind_Discard_dollar_Dict gopurs_runtime.Value
var once_Control_Bind_Discard_dollar_Dict sync.Once

func Get_Control_Bind_Discard_dollar_Dict() gopurs_runtime.Value {
	once_Control_Bind_Discard_dollar_Dict.Do(func() {
		cache_Control_Bind_Discard_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2260728934, UnsafePtr: unsafe.Pointer(Call_Control_Bind_Discard_dollar_Dict(func() struct {
				discard gopurs_runtime.Value
				/*gopurs:record-v1:1:0064006900730063006100720064*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					discard gopurs_runtime.Value
					/*gopurs:record-v1:1:0064006900730063006100720064*/
				}{}
				clone.discard = gopurs_runtime.RecordGet(orig, "discard")
				return clone
			}()))}
		})
	})
	return cache_Control_Bind_Discard_dollar_Dict
}

var cache_Control_Bind_discard gopurs_runtime.Value
var once_Control_Bind_discard sync.Once

func Get_Control_Bind_discard() gopurs_runtime.Value {
	once_Control_Bind_discard.Do(func() {
		cache_Control_Bind_discard = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_discard(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Discard[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Bind_discard
}

var cache_Control_Bind_bind gopurs_runtime.Value
var once_Control_Bind_bind sync.Once

func Get_Control_Bind_bind() gopurs_runtime.Value {
	once_Control_Bind_bind.Do(func() {
		cache_Control_Bind_bind = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Bind_bind
}

var cache_Control_Bind_discardUnit gopurs_runtime.Value
var once_Control_Bind_discardUnit sync.Once

func Get_Control_Bind_discardUnit() gopurs_runtime.Value {
	once_Control_Bind_discardUnit.Do(func() {
		cache_Control_Bind_discardUnit = gopurs_runtime.Value{Type: 9, IntVal: 2260728934, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Discard[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0))
		})}))}
	})
	return cache_Control_Bind_discardUnit
}

var cache_Control_Bind_discard__4273756336 gopurs_runtime.Value
var once_Control_Bind_discard__4273756336 sync.Once

func Get_Control_Bind_discard__4273756336() gopurs_runtime.Value {
	once_Control_Bind_discard__4273756336.Do(func() {
		cache_Control_Bind_discard__4273756336 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_discard__4273756336(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_discard__4273756336
}

var cache_Control_Bind_discard__932405161 gopurs_runtime.Value
var once_Control_Bind_discard__932405161 sync.Once

func Get_Control_Bind_discard__932405161() gopurs_runtime.Value {
	once_Control_Bind_discard__932405161.Do(func() {
		cache_Control_Bind_discard__932405161 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_discard__932405161(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_discard__932405161
}

var cache_Control_Bind_discard__1714184905 gopurs_runtime.Value
var once_Control_Bind_discard__1714184905 sync.Once

func Get_Control_Bind_discard__1714184905() gopurs_runtime.Value {
	once_Control_Bind_discard__1714184905.Do(func() {
		cache_Control_Bind_discard__1714184905 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_discard__1714184905(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_discard__1714184905
}

var cache_Control_Bind_discard__2928924777 gopurs_runtime.Value
var once_Control_Bind_discard__2928924777 sync.Once

func Get_Control_Bind_discard__2928924777() gopurs_runtime.Value {
	once_Control_Bind_discard__2928924777.Do(func() {
		cache_Control_Bind_discard__2928924777 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_discard__2928924777(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_discard__2928924777
}

var cache_Control_Bind_discard__1192935977 gopurs_runtime.Value
var once_Control_Bind_discard__1192935977 sync.Once

func Get_Control_Bind_discard__1192935977() gopurs_runtime.Value {
	once_Control_Bind_discard__1192935977.Do(func() {
		cache_Control_Bind_discard__1192935977 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_discard__1192935977(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_discard__1192935977
}

var cache_Control_Bind_discard__1759547399 gopurs_runtime.Value
var once_Control_Bind_discard__1759547399 sync.Once

func Get_Control_Bind_discard__1759547399() gopurs_runtime.Value {
	once_Control_Bind_discard__1759547399.Do(func() {
		cache_Control_Bind_discard__1759547399 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_discard__1759547399(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Bind_discard__1759547399
}

var cache_Control_Bind_discard__143037513 gopurs_runtime.Value
var once_Control_Bind_discard__143037513 sync.Once

func Get_Control_Bind_discard__143037513() gopurs_runtime.Value {
	once_Control_Bind_discard__143037513.Do(func() {
		cache_Control_Bind_discard__143037513 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_discard__143037513(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_discard__143037513
}

var cache_Control_Bind_bindProxy gopurs_runtime.Value
var once_Control_Bind_bindProxy sync.Once

func Get_Control_Bind_bindProxy() gopurs_runtime.Value {
	once_Control_Bind_bindProxy.Do(func() {
		cache_Control_Bind_bindProxy = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_4014717844_2748095225((&Constructor_Control_Bind_Bind[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_1303115796_3741347833(Rebox_Control_Bind_3741347833_1303115796(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Control_Apply_applyProxy()))))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
		})})))}
	})
	return cache_Control_Bind_bindProxy
}

var cache_Control_Bind_bindFn gopurs_runtime.Value
var once_Control_Bind_bindFn sync.Once

func Get_Control_Bind_bindFn() gopurs_runtime.Value {
	once_Control_Bind_bindFn.Do(func() {
		cache_Control_Bind_bindFn = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Apply_applyFn()
		}), gopurs_runtime.Func3(func(m_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(m_0, x_2), x_2)
		})}))}
	})
	return cache_Control_Bind_bindFn
}

var cache_Control_Bind_bindArray gopurs_runtime.Value
var once_Control_Bind_bindArray sync.Once

func Get_Control_Bind_bindArray() gopurs_runtime.Value {
	once_Control_Bind_bindArray.Do(func() {
		cache_Control_Bind_bindArray = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Control_Apply_applyArray()))}
		}), Get_Control_Bind_arrayBind()}))}
	})
	return cache_Control_Bind_bindArray
}

var cache_Control_Bind_bind__968027882 gopurs_runtime.Value
var once_Control_Bind_bind__968027882 sync.Once

func Get_Control_Bind_bind__968027882() gopurs_runtime.Value {
	once_Control_Bind_bind__968027882.Do(func() {
		cache_Control_Bind_bind__968027882 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bind__968027882(__eta0_0_box, __eta1_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Control_Bind_bind__968027882
}

var cache_Control_Bind_bind__1394626596 gopurs_runtime.Value
var once_Control_Bind_bind__1394626596 sync.Once

func Get_Control_Bind_bind__1394626596() gopurs_runtime.Value {
	once_Control_Bind_bind__1394626596.Do(func() {
		cache_Control_Bind_bind__1394626596 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bind__1394626596(__eta_norm_1_0_box, __eta_norm_0_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Control_Bind_bind__1394626596
}

var cache_Control_Bind_bind__3973282666 gopurs_runtime.Value
var once_Control_Bind_bind__3973282666 sync.Once

func Get_Control_Bind_bind__3973282666() gopurs_runtime.Value {
	once_Control_Bind_bind__3973282666.Do(func() {
		cache_Control_Bind_bind__3973282666 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bind__3973282666(Rebox_Control_Bind_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box)), __eta1_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Bind_bind__3973282666
}

var cache_Control_Bind_bind__2253334596 gopurs_runtime.Value
var once_Control_Bind_bind__2253334596 sync.Once

func Get_Control_Bind_bind__2253334596() gopurs_runtime.Value {
	once_Control_Bind_bind__2253334596.Do(func() {
		cache_Control_Bind_bind__2253334596 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bind__2253334596(Rebox_Control_Bind_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0_box)), __eta_norm_0_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Bind_bind__2253334596
}

var cache_Control_Bind_bind__1743947658 gopurs_runtime.Value
var once_Control_Bind_bind__1743947658 sync.Once

func Get_Control_Bind_bind__1743947658() gopurs_runtime.Value {
	once_Control_Bind_bind__1743947658.Do(func() {
		cache_Control_Bind_bind__1743947658 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bind__1743947658(Rebox_Control_Bind_3094389156_3516139266(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box)), __eta1_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Bind_bind__1743947658
}

var cache_Control_Bind_bind__2848886154 gopurs_runtime.Value
var once_Control_Bind_bind__2848886154 sync.Once

func Get_Control_Bind_bind__2848886154() gopurs_runtime.Value {
	once_Control_Bind_bind__2848886154.Do(func() {
		cache_Control_Bind_bind__2848886154 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bind__2848886154(Rebox_Control_Bind_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box)), __eta1_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Bind_bind__2848886154
}

var cache_Control_Bind_bind__3209581674 gopurs_runtime.Value
var once_Control_Bind_bind__3209581674 sync.Once

func Get_Control_Bind_bind__3209581674() gopurs_runtime.Value {
	once_Control_Bind_bind__3209581674.Do(func() {
		cache_Control_Bind_bind__3209581674 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bind__3209581674(Rebox_Control_Bind_3094389156_3302008615(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box)), __eta1_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Bind_bind__3209581674
}

var cache_Control_Bind_bind__3488709124 gopurs_runtime.Value
var once_Control_Bind_bind__3488709124 sync.Once

func Get_Control_Bind_bind__3488709124() gopurs_runtime.Value {
	once_Control_Bind_bind__3488709124.Do(func() {
		cache_Control_Bind_bind__3488709124 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bind__3488709124(Rebox_Control_Bind_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0_box)), __eta_norm_0_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Bind_bind__3488709124
}

var cache_Control_Bind_bind__2984238339 gopurs_runtime.Value
var once_Control_Bind_bind__2984238339 sync.Once

func Get_Control_Bind_bind__2984238339() gopurs_runtime.Value {
	once_Control_Bind_bind__2984238339.Do(func() {
		cache_Control_Bind_bind__2984238339 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__2984238339(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Bind_bind__2984238339
}

var cache_Control_Bind_bind__3987706115 gopurs_runtime.Value
var once_Control_Bind_bind__3987706115 sync.Once

func Get_Control_Bind_bind__3987706115() gopurs_runtime.Value {
	once_Control_Bind_bind__3987706115.Do(func() {
		cache_Control_Bind_bind__3987706115 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__3987706115(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Bind_bind__3987706115
}

var cache_Control_Bind_bind__4176309933 gopurs_runtime.Value
var once_Control_Bind_bind__4176309933 sync.Once

func Get_Control_Bind_bind__4176309933() gopurs_runtime.Value {
	once_Control_Bind_bind__4176309933.Do(func() {
		cache_Control_Bind_bind__4176309933 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__4176309933(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__4176309933
}

var cache_Control_Bind_bind__156704557 gopurs_runtime.Value
var once_Control_Bind_bind__156704557 sync.Once

func Get_Control_Bind_bind__156704557() gopurs_runtime.Value {
	once_Control_Bind_bind__156704557.Do(func() {
		cache_Control_Bind_bind__156704557 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__156704557(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__156704557
}

var cache_Control_Bind_bind__1318103565 gopurs_runtime.Value
var once_Control_Bind_bind__1318103565 sync.Once

func Get_Control_Bind_bind__1318103565() gopurs_runtime.Value {
	once_Control_Bind_bind__1318103565.Do(func() {
		cache_Control_Bind_bind__1318103565 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__1318103565(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__1318103565
}

var cache_Control_Bind_bind__2624874445 gopurs_runtime.Value
var once_Control_Bind_bind__2624874445 sync.Once

func Get_Control_Bind_bind__2624874445() gopurs_runtime.Value {
	once_Control_Bind_bind__2624874445.Do(func() {
		cache_Control_Bind_bind__2624874445 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__2624874445(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__2624874445
}

var cache_Control_Bind_bind__4145807597 gopurs_runtime.Value
var once_Control_Bind_bind__4145807597 sync.Once

func Get_Control_Bind_bind__4145807597() gopurs_runtime.Value {
	once_Control_Bind_bind__4145807597.Do(func() {
		cache_Control_Bind_bind__4145807597 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__4145807597(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__4145807597
}

var cache_Control_Bind_bind__949378708 gopurs_runtime.Value
var once_Control_Bind_bind__949378708 sync.Once

func Get_Control_Bind_bind__949378708() gopurs_runtime.Value {
	once_Control_Bind_bind__949378708.Do(func() {
		cache_Control_Bind_bind__949378708 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__949378708(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__949378708
}

var cache_Control_Bind_bind__2676662740 gopurs_runtime.Value
var once_Control_Bind_bind__2676662740 sync.Once

func Get_Control_Bind_bind__2676662740() gopurs_runtime.Value {
	once_Control_Bind_bind__2676662740.Do(func() {
		cache_Control_Bind_bind__2676662740 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__2676662740(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__2676662740
}

var cache_Control_Bind_bind__845594356 gopurs_runtime.Value
var once_Control_Bind_bind__845594356 sync.Once

func Get_Control_Bind_bind__845594356() gopurs_runtime.Value {
	once_Control_Bind_bind__845594356.Do(func() {
		cache_Control_Bind_bind__845594356 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__845594356(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__845594356
}

var cache_Control_Bind_bind__1207231796 gopurs_runtime.Value
var once_Control_Bind_bind__1207231796 sync.Once

func Get_Control_Bind_bind__1207231796() gopurs_runtime.Value {
	once_Control_Bind_bind__1207231796.Do(func() {
		cache_Control_Bind_bind__1207231796 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__1207231796(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__1207231796
}

var cache_Control_Bind_bind__3135657434 gopurs_runtime.Value
var once_Control_Bind_bind__3135657434 sync.Once

func Get_Control_Bind_bind__3135657434() gopurs_runtime.Value {
	once_Control_Bind_bind__3135657434.Do(func() {
		cache_Control_Bind_bind__3135657434 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__3135657434(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Bind_bind__3135657434
}

var cache_Control_Bind_bind__4149781044 gopurs_runtime.Value
var once_Control_Bind_bind__4149781044 sync.Once

func Get_Control_Bind_bind__4149781044() gopurs_runtime.Value {
	once_Control_Bind_bind__4149781044.Do(func() {
		cache_Control_Bind_bind__4149781044 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__4149781044(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__4149781044
}

var cache_Control_Bind_bind__1370440500 gopurs_runtime.Value
var once_Control_Bind_bind__1370440500 sync.Once

func Get_Control_Bind_bind__1370440500() gopurs_runtime.Value {
	once_Control_Bind_bind__1370440500.Do(func() {
		cache_Control_Bind_bind__1370440500 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__1370440500(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__1370440500
}

var cache_Control_Bind_bind__1455934228 gopurs_runtime.Value
var once_Control_Bind_bind__1455934228 sync.Once

func Get_Control_Bind_bind__1455934228() gopurs_runtime.Value {
	once_Control_Bind_bind__1455934228.Do(func() {
		cache_Control_Bind_bind__1455934228 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__1455934228(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__1455934228
}

var cache_Control_Bind_bind__1941644084 gopurs_runtime.Value
var once_Control_Bind_bind__1941644084 sync.Once

func Get_Control_Bind_bind__1941644084() gopurs_runtime.Value {
	once_Control_Bind_bind__1941644084.Do(func() {
		cache_Control_Bind_bind__1941644084 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__1941644084(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__1941644084
}

var cache_Control_Bind_bind__1903947764 gopurs_runtime.Value
var once_Control_Bind_bind__1903947764 sync.Once

func Get_Control_Bind_bind__1903947764() gopurs_runtime.Value {
	once_Control_Bind_bind__1903947764.Do(func() {
		cache_Control_Bind_bind__1903947764 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__1903947764(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__1903947764
}

var cache_Control_Bind_bind__406760756 gopurs_runtime.Value
var once_Control_Bind_bind__406760756 sync.Once

func Get_Control_Bind_bind__406760756() gopurs_runtime.Value {
	once_Control_Bind_bind__406760756.Do(func() {
		cache_Control_Bind_bind__406760756 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__406760756(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__406760756
}

var cache_Control_Bind_bind__3500963866 gopurs_runtime.Value
var once_Control_Bind_bind__3500963866 sync.Once

func Get_Control_Bind_bind__3500963866() gopurs_runtime.Value {
	once_Control_Bind_bind__3500963866.Do(func() {
		cache_Control_Bind_bind__3500963866 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__3500963866(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Bind_bind__3500963866
}

var cache_Control_Bind_bind__3513850932 gopurs_runtime.Value
var once_Control_Bind_bind__3513850932 sync.Once

func Get_Control_Bind_bind__3513850932() gopurs_runtime.Value {
	once_Control_Bind_bind__3513850932.Do(func() {
		cache_Control_Bind_bind__3513850932 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__3513850932(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__3513850932
}

var cache_Control_Bind_bind__2639965556 gopurs_runtime.Value
var once_Control_Bind_bind__2639965556 sync.Once

func Get_Control_Bind_bind__2639965556() gopurs_runtime.Value {
	once_Control_Bind_bind__2639965556.Do(func() {
		cache_Control_Bind_bind__2639965556 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__2639965556(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__2639965556
}

var cache_Control_Bind_bind__931913300 gopurs_runtime.Value
var once_Control_Bind_bind__931913300 sync.Once

func Get_Control_Bind_bind__931913300() gopurs_runtime.Value {
	once_Control_Bind_bind__931913300.Do(func() {
		cache_Control_Bind_bind__931913300 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__931913300(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__931913300
}

var cache_Control_Bind_bind__4216591092 gopurs_runtime.Value
var once_Control_Bind_bind__4216591092 sync.Once

func Get_Control_Bind_bind__4216591092() gopurs_runtime.Value {
	once_Control_Bind_bind__4216591092.Do(func() {
		cache_Control_Bind_bind__4216591092 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__4216591092(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__4216591092
}

var cache_Control_Bind_bind__3398628340 gopurs_runtime.Value
var once_Control_Bind_bind__3398628340 sync.Once

func Get_Control_Bind_bind__3398628340() gopurs_runtime.Value {
	once_Control_Bind_bind__3398628340.Do(func() {
		cache_Control_Bind_bind__3398628340 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__3398628340(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__3398628340
}

var cache_Control_Bind_bind__4020646868 gopurs_runtime.Value
var once_Control_Bind_bind__4020646868 sync.Once

func Get_Control_Bind_bind__4020646868() gopurs_runtime.Value {
	once_Control_Bind_bind__4020646868.Do(func() {
		cache_Control_Bind_bind__4020646868 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__4020646868(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__4020646868
}

var cache_Control_Bind_bind__967963220 gopurs_runtime.Value
var once_Control_Bind_bind__967963220 sync.Once

func Get_Control_Bind_bind__967963220() gopurs_runtime.Value {
	once_Control_Bind_bind__967963220.Do(func() {
		cache_Control_Bind_bind__967963220 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__967963220(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__967963220
}

var cache_Control_Bind_bind__2595690612 gopurs_runtime.Value
var once_Control_Bind_bind__2595690612 sync.Once

func Get_Control_Bind_bind__2595690612() gopurs_runtime.Value {
	once_Control_Bind_bind__2595690612.Do(func() {
		cache_Control_Bind_bind__2595690612 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__2595690612(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__2595690612
}

var cache_Control_Bind_bind__3918566810 gopurs_runtime.Value
var once_Control_Bind_bind__3918566810 sync.Once

func Get_Control_Bind_bind__3918566810() gopurs_runtime.Value {
	once_Control_Bind_bind__3918566810.Do(func() {
		cache_Control_Bind_bind__3918566810 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__3918566810(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Bind_bind__3918566810
}

var cache_Control_Bind_bind__4262485428 gopurs_runtime.Value
var once_Control_Bind_bind__4262485428 sync.Once

func Get_Control_Bind_bind__4262485428() gopurs_runtime.Value {
	once_Control_Bind_bind__4262485428.Do(func() {
		cache_Control_Bind_bind__4262485428 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__4262485428(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bind__4262485428
}

var cache_Control_Bind_bind__1864922452 gopurs_runtime.Value
var once_Control_Bind_bind__1864922452 sync.Once

func Get_Control_Bind_bind__1864922452() gopurs_runtime.Value {
	once_Control_Bind_bind__1864922452.Do(func() {
		cache_Control_Bind_bind__1864922452 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind__1864922452(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Control_Bind_bind__1864922452
}

var cache_Control_Bind_bindFlipped gopurs_runtime.Value
var once_Control_Bind_bindFlipped sync.Once

func Get_Control_Bind_bindFlipped() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped.Do(func() {
		cache_Control_Bind_bindFlipped = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bindFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0_box))
		})
	})
	return cache_Control_Bind_bindFlipped
}

var cache_Control_Bind_bindFlipped__2935978538 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__2935978538 sync.Once

func Get_Control_Bind_bindFlipped__2935978538() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__2935978538.Do(func() {
		cache_Control_Bind_bindFlipped__2935978538 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bindFlipped__2935978538(__eta0_0_box, __eta1_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Control_Bind_bindFlipped__2935978538
}

var cache_Control_Bind_bindFlipped__79681002 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__79681002 sync.Once

func Get_Control_Bind_bindFlipped__79681002() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__79681002.Do(func() {
		cache_Control_Bind_bindFlipped__79681002 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bindFlipped__79681002(__eta0_0_box, __eta1_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Control_Bind_bindFlipped__79681002
}

var cache_Control_Bind_bindFlipped__1041122314 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__1041122314 sync.Once

func Get_Control_Bind_bindFlipped__1041122314() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__1041122314.Do(func() {
		cache_Control_Bind_bindFlipped__1041122314 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bindFlipped__1041122314(__eta0_0_box, __eta1_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Control_Bind_bindFlipped__1041122314
}

var cache_Control_Bind_bindFlipped__3340339946 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__3340339946 sync.Once

func Get_Control_Bind_bindFlipped__3340339946() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__3340339946.Do(func() {
		cache_Control_Bind_bindFlipped__3340339946 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bindFlipped__3340339946(__eta0_0_box, Rebox_Control_Bind_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Bind_bindFlipped__3340339946
}

var cache_Control_Bind_bindFlipped__1302780650 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__1302780650 sync.Once

func Get_Control_Bind_bindFlipped__1302780650() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__1302780650.Do(func() {
		cache_Control_Bind_bindFlipped__1302780650 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_bindFlipped__1302780650(__eta0_0_box, Rebox_Control_Bind_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Bind_bindFlipped__1302780650
}

var cache_Control_Bind_bindFlipped__2373983747 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__2373983747 sync.Once

func Get_Control_Bind_bindFlipped__2373983747() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__2373983747.Do(func() {
		cache_Control_Bind_bindFlipped__2373983747 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bindFlipped__2373983747(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Bind_bindFlipped__2373983747
}

var cache_Control_Bind_bindFlipped__900387395 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__900387395 sync.Once

func Get_Control_Bind_bindFlipped__900387395() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__900387395.Do(func() {
		cache_Control_Bind_bindFlipped__900387395 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bindFlipped__900387395(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Control_Bind_bindFlipped__900387395
}

var cache_Control_Bind_bindFlipped__857303188 gopurs_runtime.Value
var once_Control_Bind_bindFlipped__857303188 sync.Once

func Get_Control_Bind_bindFlipped__857303188() gopurs_runtime.Value {
	once_Control_Bind_bindFlipped__857303188.Do(func() {
		cache_Control_Bind_bindFlipped__857303188 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bindFlipped__857303188(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Control_Bind_bindFlipped__857303188
}

var cache_Control_Bind_composeKleisliFlipped gopurs_runtime.Value
var once_Control_Bind_composeKleisliFlipped sync.Once

func Get_Control_Bind_composeKleisliFlipped() gopurs_runtime.Value {
	once_Control_Bind_composeKleisliFlipped.Do(func() {
		cache_Control_Bind_composeKleisliFlipped = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_composeKleisliFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0_box))
		})
	})
	return cache_Control_Bind_composeKleisliFlipped
}

var cache_Control_Bind_composeKleisliFlipped__1244140651 gopurs_runtime.Value
var once_Control_Bind_composeKleisliFlipped__1244140651 sync.Once

func Get_Control_Bind_composeKleisliFlipped__1244140651() gopurs_runtime.Value {
	once_Control_Bind_composeKleisliFlipped__1244140651.Do(func() {
		cache_Control_Bind_composeKleisliFlipped__1244140651 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_composeKleisliFlipped__1244140651(__eta0_0_box, __eta1_1_box, __eta2_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Control_Bind_composeKleisliFlipped__1244140651
}

var cache_Control_Bind_composeKleisliFlipped__2774842251 gopurs_runtime.Value
var once_Control_Bind_composeKleisliFlipped__2774842251 sync.Once

func Get_Control_Bind_composeKleisliFlipped__2774842251() gopurs_runtime.Value {
	once_Control_Bind_composeKleisliFlipped__2774842251.Do(func() {
		cache_Control_Bind_composeKleisliFlipped__2774842251 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_composeKleisliFlipped__2774842251(__eta0_0_box, __eta1_1_box, __eta2_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Control_Bind_composeKleisliFlipped__2774842251
}

var cache_Control_Bind_composeKleisliFlipped__647481931 gopurs_runtime.Value
var once_Control_Bind_composeKleisliFlipped__647481931 sync.Once

func Get_Control_Bind_composeKleisliFlipped__647481931() gopurs_runtime.Value {
	once_Control_Bind_composeKleisliFlipped__647481931.Do(func() {
		cache_Control_Bind_composeKleisliFlipped__647481931 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_composeKleisliFlipped__647481931(__eta0_0_box, __eta1_1_box, __eta2_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Control_Bind_composeKleisliFlipped__647481931
}

var cache_Control_Bind_composeKleisliFlipped__2887028321 gopurs_runtime.Value
var once_Control_Bind_composeKleisliFlipped__2887028321 sync.Once

func Get_Control_Bind_composeKleisliFlipped__2887028321() gopurs_runtime.Value {
	once_Control_Bind_composeKleisliFlipped__2887028321.Do(func() {
		cache_Control_Bind_composeKleisliFlipped__2887028321 = gopurs_runtime.Func3(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_composeKleisliFlipped__2887028321(__eta_norm_1_0_box, __eta_norm_0_unused_1_box, a_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Control_Bind_composeKleisliFlipped__2887028321
}

var cache_Control_Bind_composeKleisliFlipped__354839169 gopurs_runtime.Value
var once_Control_Bind_composeKleisliFlipped__354839169 sync.Once

func Get_Control_Bind_composeKleisliFlipped__354839169() gopurs_runtime.Value {
	once_Control_Bind_composeKleisliFlipped__354839169.Do(func() {
		cache_Control_Bind_composeKleisliFlipped__354839169 = gopurs_runtime.Func3(func(__eta_norm_2_0_box gopurs_runtime.Value, __eta_norm_1_unused_1_box gopurs_runtime.Value, __eta_norm_0_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_composeKleisliFlipped__354839169(__eta_norm_2_0_box, __eta_norm_1_unused_1_box, __eta_norm_0_2_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Control_Bind_composeKleisliFlipped__354839169
}

var cache_Control_Bind_composeKleisliFlipped__904307531 gopurs_runtime.Value
var once_Control_Bind_composeKleisliFlipped__904307531 sync.Once

func Get_Control_Bind_composeKleisliFlipped__904307531() gopurs_runtime.Value {
	once_Control_Bind_composeKleisliFlipped__904307531.Do(func() {
		cache_Control_Bind_composeKleisliFlipped__904307531 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_composeKleisliFlipped__904307531(__eta0_0_box, __eta1_1_box, __eta2_2_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Bind_composeKleisliFlipped__904307531
}

var cache_Control_Bind_composeKleisli gopurs_runtime.Value
var once_Control_Bind_composeKleisli sync.Once

func Get_Control_Bind_composeKleisli() gopurs_runtime.Value {
	once_Control_Bind_composeKleisli.Do(func() {
		cache_Control_Bind_composeKleisli = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_composeKleisli(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0_box), f_1_box, g_2_box, a_3_box)
		})
	})
	return cache_Control_Bind_composeKleisli
}

var cache_Control_Bind_composeKleisli__2131995638 gopurs_runtime.Value
var once_Control_Bind_composeKleisli__2131995638 sync.Once

func Get_Control_Bind_composeKleisli__2131995638() gopurs_runtime.Value {
	once_Control_Bind_composeKleisli__2131995638.Do(func() {
		cache_Control_Bind_composeKleisli__2131995638 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_composeKleisli__2131995638(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Bind_composeKleisli__2131995638
}

var cache_Control_Bind_composeKleisli__1522139914 gopurs_runtime.Value
var once_Control_Bind_composeKleisli__1522139914 sync.Once

func Get_Control_Bind_composeKleisli__1522139914() gopurs_runtime.Value {
	once_Control_Bind_composeKleisli__1522139914.Do(func() {
		cache_Control_Bind_composeKleisli__1522139914 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_composeKleisli__1522139914(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
		})
	})
	return cache_Control_Bind_composeKleisli__1522139914
}

var cache_Control_Bind_discardProxy gopurs_runtime.Value
var once_Control_Bind_discardProxy sync.Once

func Get_Control_Bind_discardProxy() gopurs_runtime.Value {
	once_Control_Bind_discardProxy.Do(func() {
		cache_Control_Bind_discardProxy = gopurs_runtime.Value{Type: 9, IntVal: 2260728934, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_3559868223_177467250((&Constructor_Control_Bind_Discard[uint32]{1, gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0))
		})})))}
	})
	return cache_Control_Bind_discardProxy
}

var cache_Control_Bind_ifM gopurs_runtime.Value
var once_Control_Bind_ifM sync.Once

func Get_Control_Bind_ifM() gopurs_runtime.Value {
	once_Control_Bind_ifM.Do(func() {
		cache_Control_Bind_ifM = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, cond_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_ifM(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0_box), cond_1_box, t_2_box, f_3_box)
		})
	})
	return cache_Control_Bind_ifM
}

var cache_Control_Bind_join gopurs_runtime.Value
var once_Control_Bind_join sync.Once

func Get_Control_Bind_join() gopurs_runtime.Value {
	once_Control_Bind_join.Do(func() {
		cache_Control_Bind_join = gopurs_runtime.Func2(func(dictBind_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Bind_join(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](dictBind_0_box), m_1_box)
		})
	})
	return cache_Control_Bind_join
}

var cache_Control_Bind_join__3351431019 gopurs_runtime.Value
var once_Control_Bind_join__3351431019 sync.Once

func Get_Control_Bind_join__3351431019() gopurs_runtime.Value {
	once_Control_Bind_join__3351431019.Do(func() {
		cache_Control_Bind_join__3351431019 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Bind_join__3351431019(Rebox_Control_Bind_3094389156_3516139266(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta0_0_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Control_Bind_join__3351431019
}

type Constructor_Control_Bind_Bind[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f00420069006e0064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f00420069006e0064005f00420069006e0064
func init() {
	gopurs_runtime.StructGetters[4032919565] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Bind_Bind[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Apply0":
			return gopurs_runtime.Box(c.V0)
		case "bind":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Control_Bind_Bind: " + key)
		}
	}
}

type Constructor_Control_Bind_Discard[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0043006f006e00740072006f006c005f00420069006e0064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0043006f006e00740072006f006c005f00420069006e0064005f0044006900730063006100720064
func init() {
	gopurs_runtime.StructGetters[2260728934] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Bind_Discard[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "discard":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Control_Bind_Discard: " + key)
		}
	}
}

func Call_Control_Bind_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Control_Bind_Bind_dollar_Dict(x_0_loop struct {
	Apply0 gopurs_runtime.Value
	bind   gopurs_runtime.Value
	/*gopurs:record-v1:2:004100700070006c00790030,00620069006e0064*/
}) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	var x_0 struct {
		Apply0 gopurs_runtime.Value
		bind   gopurs_runtime.Value
		/*gopurs:record-v1:2:004100700070006c00790030,00620069006e0064*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Apply0", "bind", orig.Apply0, orig.bind)
	}())
}

func Call_Control_Bind_Discard_dollar_Dict(x_0_loop struct {
	discard gopurs_runtime.Value
	/*gopurs:record-v1:1:0064006900730063006100720064*/
}) *Constructor_Control_Bind_Discard[gopurs_runtime.Value] {
	var x_0 struct {
		discard gopurs_runtime.Value
		/*gopurs:record-v1:1:0064006900730063006100720064*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Discard[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("discard", orig.discard)
	}())
}

func Call_Control_Bind_discard(dict_0_loop *Constructor_Control_Bind_Discard[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Bind_Discard[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Control_Bind_bind(dict_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Control_Bind_discard__4273756336(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__4273756336:
	for {
		if false {
			continue discard__4273756336
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_discard__932405161(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__932405161:
	for {
		if false {
			continue discard__932405161
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_discard__1714184905(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1714184905:
	for {
		if false {
			continue discard__1714184905
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_discard__2928924777(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__2928924777:
	for {
		if false {
			continue discard__2928924777
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_discard__1192935977(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1192935977:
	for {
		if false {
			continue discard__1192935977
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_discard__1759547399(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__1759547399:
	for {
		if false {
			continue discard__1759547399
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Control_Bind_discard__143037513(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
discard__143037513:
	for {
		if false {
			continue discard__143037513
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__968027882(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
bind__968027882:
	for {
		if false {
			continue bind__968027882
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind"), __eta0_0, __eta1_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_bind__1394626596(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
bind__1394626596:
	for {
		if false {
			continue bind__1394626596
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind"), __eta_norm_1_0, __eta_norm_0_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_bind__3973282666(__eta0_0_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date], __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
bind__3973282666:
	for {
		if false {
			continue bind__3973282666
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156(Rebox_Control_Bind_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156(__eta0_0))}, __eta1_1)))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_bind__2253334596(__eta_norm_1_0_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date], __eta_norm_0_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
bind__2253334596:
	for {
		if false {
			continue bind__2253334596
		}
		var __eta_norm_1_0 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156(__eta_norm_1_0))}, __eta_norm_0_1); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_bind__1743947658(__eta0_0_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]], __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
bind__1743947658:
	for {
		if false {
			continue bind__1743947658
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156(Rebox_Control_Bind_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_3516139266_3094389156(__eta0_0))}, __eta1_1)))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_bind__2848886154(__eta0_0_loop *Constructor_Data_Maybe_Just[int64], __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
bind__2848886154:
	for {
		if false {
			continue bind__2848886154
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[int64] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156(Rebox_Control_Bind_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_1170268447_3094389156(__eta0_0))}, __eta1_1)))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_bind__3209581674(__eta0_0_loop *Constructor_Data_Maybe_Just[struct {
	day         int64
	hour        int64
	millisecond int64
	minute      int64
	month       int64
	second      int64
	year        int64
	/*gopurs:record-v1:7:006400610079,0068006f00750072,006d0069006c006c0069007300650063006f006e0064,006d0069006e007500740065,006d006f006e00740068,007300650063006f006e0064,0079006500610072*/
}], __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
bind__3209581674:
	for {
		if false {
			continue bind__3209581674
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[struct {
			day         int64
			hour        int64
			millisecond int64
			minute      int64
			month       int64
			second      int64
			year        int64
			/*gopurs:record-v1:7:006400610079,0068006f00750072,006d0069006c006c0069007300650063006f006e0064,006d0069006e007500740065,006d006f006e00740068,007300650063006f006e0064,0079006500610072*/
		}] = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_4151030883_3094389156(Rebox_Control_Bind_3094389156_4151030883(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_3302008615_3094389156(__eta0_0))}, __eta1_1)))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_bind__3488709124(__eta_norm_1_0_loop *Constructor_Data_Maybe_Just[string], __eta_norm_0_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
bind__3488709124:
	for {
		if false {
			continue bind__3488709124
		}
		var __eta_norm_1_0 *Constructor_Data_Maybe_Just[string] = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_742090555_3094389156(__eta_norm_1_0))}, __eta_norm_0_1); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_bind__2984238339(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2984238339:
	for {
		if false {
			continue bind__2984238339
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()).V1, __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Control_Bind_bind__3987706115(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3987706115:
	for {
		if false {
			continue bind__3987706115
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()).V1, __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Control_Bind_bind__4176309933(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__4176309933:
	for {
		if false {
			continue bind__4176309933
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__156704557(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__156704557:
	for {
		if false {
			continue bind__156704557
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__1318103565(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1318103565:
	for {
		if false {
			continue bind__1318103565
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__2624874445(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2624874445:
	for {
		if false {
			continue bind__2624874445
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__4145807597(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__4145807597:
	for {
		if false {
			continue bind__4145807597
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__949378708(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__949378708:
	for {
		if false {
			continue bind__949378708
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__2676662740(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2676662740:
	for {
		if false {
			continue bind__2676662740
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__845594356(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__845594356:
	for {
		if false {
			continue bind__845594356
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__1207231796(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1207231796:
	for {
		if false {
			continue bind__1207231796
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__3135657434(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3135657434:
	for {
		if false {
			continue bind__3135657434
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Control_Bind_bind__4149781044(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__4149781044:
	for {
		if false {
			continue bind__4149781044
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__1370440500(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1370440500:
	for {
		if false {
			continue bind__1370440500
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__1455934228(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1455934228:
	for {
		if false {
			continue bind__1455934228
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__1941644084(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1941644084:
	for {
		if false {
			continue bind__1941644084
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__1903947764(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__1903947764:
	for {
		if false {
			continue bind__1903947764
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__406760756(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__406760756:
	for {
		if false {
			continue bind__406760756
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__3500963866(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3500963866:
	for {
		if false {
			continue bind__3500963866
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Control_Bind_bind__3513850932(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3513850932:
	for {
		if false {
			continue bind__3513850932
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__2639965556(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2639965556:
	for {
		if false {
			continue bind__2639965556
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__931913300(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__931913300:
	for {
		if false {
			continue bind__931913300
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__4216591092(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__4216591092:
	for {
		if false {
			continue bind__4216591092
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__3398628340(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3398628340:
	for {
		if false {
			continue bind__3398628340
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__4020646868(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__4020646868:
	for {
		if false {
			continue bind__4020646868
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__967963220(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__967963220:
	for {
		if false {
			continue bind__967963220
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__2595690612(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__2595690612:
	for {
		if false {
			continue bind__2595690612
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__3918566810(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__3918566810:
	for {
		if false {
			continue bind__3918566810
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Control_Bind_bind__4262485428(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bind__4262485428:
	for {
		if false {
			continue bind__4262485428
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta0_0, __eta1_1)
	}
}

func Call_Control_Bind_bind__1864922452(dict_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
bind__1864922452:
	for {
		if false {
			continue bind__1864922452
		}
		var dict_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dict_0_loop
		_ = dict_0
		return dict_0.V1
	}
}

func Call_Control_Bind_bindFlipped(dictBind_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictBind_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dictBind_0_loop
	_ = dictBind_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Control_Bind_bind(dictBind_0)
	_ = __local_var_1_0
	return gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(__local_var_1_0, a_3, b_2)
	})
}

func Call_Control_Bind_bindFlipped__2935978538(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
bindFlipped__2935978538:
	for {
		if false {
			continue bindFlipped__2935978538
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind"), __eta1_1, __eta0_0); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_bindFlipped__79681002(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
bindFlipped__79681002:
	for {
		if false {
			continue bindFlipped__79681002
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind"), __eta1_1, __eta0_0); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_bindFlipped__1041122314(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
bindFlipped__1041122314:
	for {
		if false {
			continue bindFlipped__1041122314
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind"), __eta1_1, __eta0_0); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_bindFlipped__3340339946(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
bindFlipped__3340339946:
	for {
		if false {
			continue bindFlipped__3340339946
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156(Rebox_Control_Bind_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156(__eta1_1))}, __eta0_0)))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_bindFlipped__1302780650(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop *Constructor_Data_Maybe_Just[string]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
bindFlipped__1302780650:
	for {
		if false {
			continue bindFlipped__1302780650
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[string] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_742090555_3094389156(Rebox_Control_Bind_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_742090555_3094389156(__eta1_1))}, __eta0_0)))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_bindFlipped__2373983747(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bindFlipped__2373983747:
	for {
		if false {
			continue bindFlipped__2373983747
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()).V1, __eta_norm_0_1, __eta_norm_1_0)
	}
}

func Call_Control_Bind_bindFlipped__900387395(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bindFlipped__900387395:
	for {
		if false {
			continue bindFlipped__900387395
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff()).V1, __eta_norm_0_1, __eta_norm_1_0)
	}
}

func Call_Control_Bind_bindFlipped__857303188(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bindFlipped__857303188:
	for {
		if false {
			continue bindFlipped__857303188
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect()).V1, __eta1_1, __eta0_0)
	}
}

func Call_Control_Bind_composeKleisliFlipped(dictBind_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictBind_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dictBind_0_loop
	_ = dictBind_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := Call_Control_Bind_bind(dictBind_0)
	_ = __local_var_1_0
	return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(__local_var_1_0, gopurs_runtime.Apply(g_3, a_4), f_2)
	})
}

func Call_Control_Bind_composeKleisliFlipped__1244140651(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
composeKleisliFlipped__1244140651:
	for {
		if false {
			continue composeKleisliFlipped__1244140651
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
		_ = __eta2_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind"), gopurs_runtime.Apply(__eta1_1, __eta2_2), __eta0_0); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_composeKleisliFlipped__2774842251(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
composeKleisliFlipped__2774842251:
	for {
		if false {
			continue composeKleisliFlipped__2774842251
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
		_ = __eta2_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind"), gopurs_runtime.Apply(__eta1_1, __eta2_2), __eta0_0); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_composeKleisliFlipped__647481931(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
composeKleisliFlipped__647481931:
	for {
		if false {
			continue composeKleisliFlipped__647481931
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
		_ = __eta2_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind"), gopurs_runtime.Apply(__eta1_1, __eta2_2), __eta0_0); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_composeKleisliFlipped__2887028321(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
composeKleisliFlipped__2887028321:
	for {
		if false {
			continue composeKleisliFlipped__2887028321
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		var a_2 gopurs_runtime.Value = a_2_loop
		_ = a_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind"), gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Decoders_decodeNumber(), a_2), __eta_norm_1_0); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_composeKleisliFlipped__354839169(__eta_norm_2_0_loop gopurs_runtime.Value, __eta_norm_1_unused_1_loop gopurs_runtime.Value, __eta_norm_0_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
composeKleisliFlipped__354839169:
	for {
		if false {
			continue composeKleisliFlipped__354839169
		}
		var __eta_norm_2_0 gopurs_runtime.Value = __eta_norm_2_0_loop
		_ = __eta_norm_2_0
		var __eta_norm_1_unused_1 gopurs_runtime.Value = __eta_norm_1_unused_1_loop
		_ = __eta_norm_1_unused_1
		var __eta_norm_0_2 gopurs_runtime.Value = __eta_norm_0_2_loop
		_ = __eta_norm_0_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind"), gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Decoders_decodeJArray(), __eta_norm_0_2), __eta_norm_2_0); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_composeKleisliFlipped__904307531(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
composeKleisliFlipped__904307531:
	for {
		if false {
			continue composeKleisliFlipped__904307531
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var __eta2_2 string = __eta2_2_loop
		_ = __eta2_2
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_742090555_3094389156(Rebox_Control_Bind_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_742090555_3094389156(Rebox_Control_Bind_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__eta1_1, gopurs_runtime.Str(__eta2_2))))))}, __eta0_0)))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Control_Bind_composeKleisli(dictBind_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBind_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dictBind_0_loop
	_ = dictBind_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var g_2 gopurs_runtime.Value = g_2_loop
	_ = g_2
	var a_3 gopurs_runtime.Value = a_3_loop
	_ = a_3
	return gopurs_runtime.Apply2(dictBind_0.V1, gopurs_runtime.Apply(f_1, a_3), g_2)
}

func Call_Control_Bind_composeKleisli__2131995638(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeKleisli__2131995638:
	for {
		if false {
			continue composeKleisli__2131995638
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (Record (Row [completed: Boolean] Empty))]))
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Class_decodeJson(), gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2648973148_1277236813(Rebox_Control_Bind_1277236813_2648973148(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_Argonaut_Decode_Class_decodeRecord(), gopurs_runtime.Apply5(Get_Data_Argonaut_Decode_Class_gDecodeJsonCons(), gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Class_decodeFieldId(), gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_3459698345_1277236813(Rebox_Control_Bind_1277236813_3459698345(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonBoolean()))))}), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_completedIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{}), gopurs_runtime.Value{})))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind"), gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Parser_parseJson(), gopurs_runtime.NormalizeStringValue(a_3)), __local_var_2_0)
		})
	}
}

func Call_Control_Bind_composeKleisli__1522139914(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
composeKleisli__1522139914:
	for {
		if false {
			continue composeKleisli__1522139914
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
		_ = __eta_norm_0_unused_1
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Either","Either"] [(ADT ["Data","Argonaut","Decode","Error","JsonDecodeError"] []), (Record (Row [title: String] Empty))]))
		__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Class_decodeJson(), gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2400514730_1277236813(Rebox_Control_Bind_1277236813_2400514730(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_Argonaut_Decode_Class_decodeRecord(), gopurs_runtime.Apply5(Get_Data_Argonaut_Decode_Class_gDecodeJsonCons(), gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Class_decodeFieldId(), gopurs_runtime.Value{Type: 9, IntVal: 1358001017, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_48201874_1277236813(Rebox_Control_Bind_1277236813_48201874(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_decodeJsonString()))))}), gopurs_runtime.Value{Type: 9, IntVal: 2727814238, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Argonaut_Decode_Class_GDecodeJson[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Argonaut_Decode_Class_gDecodeJsonNil()))}, Get_Main_titleIsSymbol(), gopurs_runtime.Value{}, gopurs_runtime.Value{}), gopurs_runtime.Value{})))))})
		_ = __local_var_2_0
		return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_bindEither(), "bind"), gopurs_runtime.Apply(Get_Data_Argonaut_Decode_Parser_parseJson(), gopurs_runtime.NormalizeStringValue(a_3)), __local_var_2_0)
		})
	}
}

func Call_Control_Bind_ifM(dictBind_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value], cond_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBind_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dictBind_0_loop
	_ = dictBind_0
	var cond_1 gopurs_runtime.Value = cond_1_loop
	_ = cond_1
	var t_2 gopurs_runtime.Value = t_2_loop
	_ = t_2
	var f_3 gopurs_runtime.Value = f_3_loop
	_ = f_3
	return gopurs_runtime.Apply2(dictBind_0.V1, cond_1, gopurs_runtime.Func(func(cond_prime__4 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if (cond_prime__4.IntVal) != (0) {
				__t0 = t_2
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = f_3
		}
	end_branch_0:
		return __t0
	}))
}

func Call_Control_Bind_join(dictBind_0_loop *Constructor_Control_Bind_Bind[gopurs_runtime.Value], m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictBind_0 *Constructor_Control_Bind_Bind[gopurs_runtime.Value] = dictBind_0_loop
	_ = dictBind_0
	var m_1 gopurs_runtime.Value = m_1_loop
	_ = m_1
	return gopurs_runtime.Apply2(dictBind_0.V1, m_1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_2
	}))
}

func Call_Control_Bind_join__3351431019(__eta0_0_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
join__3351431019:
	for {
		if false {
			continue join__3351431019
		}
		var __eta0_0 *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]] = __eta0_0_loop
		_ = __eta0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156(Rebox_Control_Bind_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Control_Bind_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_3516139266_3094389156(__eta0_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_1
		}))))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Rebox_Control_Bind_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Control_Bind_1277236813_2400514730(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	title string
	/*gopurs:record-v1:1:007400690074006c0065*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
		title string
		/*gopurs:record-v1:1:007400690074006c0065*/
	}])(unsafe.Pointer(in))
}

func Rebox_Control_Bind_1277236813_2648973148(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	completed bool
	/*gopurs:record-v1:1:0063006f006d0070006c0065007400650064*/
}] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
		completed bool
		/*gopurs:record-v1:1:0063006f006d0070006c0065007400650064*/
	}])(unsafe.Pointer(in))
}

func Rebox_Control_Bind_1277236813_3459698345(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[bool])(unsafe.Pointer(in))
}

func Rebox_Control_Bind_1277236813_48201874(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[string])(unsafe.Pointer(in))
}

func Rebox_Control_Bind_1303115796_3741347833(in *Constructor_Control_Apply_Apply[uint32]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Bind_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Bind_2400514730_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	title string
	/*gopurs:record-v1:1:007400690074006c0065*/
}]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Bind_2648973148_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[struct {
	completed bool
	/*gopurs:record-v1:1:0063006f006d0070006c0065007400650064*/
}]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Bind_2748095225_3887487416(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Control_Bind_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Control_Bind_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}

func Rebox_Control_Bind_3094389156_3302008615(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
	day         int64
	hour        int64
	millisecond int64
	minute      int64
	month       int64
	second      int64
	year        int64
	/*gopurs:record-v1:7:006400610079,0068006f00750072,006d0069006c006c0069007300650063006f006e0064,006d0069006e007500740065,006d006f006e00740068,007300650063006f006e0064,0079006500610072*/
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[struct {
		day         int64
		hour        int64
		millisecond int64
		minute      int64
		month       int64
		second      int64
		year        int64
		/*gopurs:record-v1:7:006400610079,0068006f00750072,006d0069006c006c0069007300650063006f006e0064,006d0069006e007500740065,006d006f006e00740068,007300650063006f006e0064,0079006500610072*/
	}]{}
	out.V0 = func() struct {
		day         int64
		hour        int64
		millisecond int64
		minute      int64
		month       int64
		second      int64
		year        int64
		/*gopurs:record-v1:7:006400610079,0068006f00750072,006d0069006c006c0069007300650063006f006e0064,006d0069006e007500740065,006d006f006e00740068,007300650063006f006e0064,0079006500610072*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			day         int64
			hour        int64
			millisecond int64
			minute      int64
			month       int64
			second      int64
			year        int64
			/*gopurs:record-v1:7:006400610079,0068006f00750072,006d0069006c006c0069007300650063006f006e0064,006d0069006e007500740065,006d006f006e00740068,007300650063006f006e0064,0079006500610072*/
		}{}
		clone.day = gopurs_runtime.RecordGet(orig, "day").IntVal
		clone.hour = gopurs_runtime.RecordGet(orig, "hour").IntVal
		clone.millisecond = gopurs_runtime.RecordGet(orig, "millisecond").IntVal
		clone.minute = gopurs_runtime.RecordGet(orig, "minute").IntVal
		clone.month = gopurs_runtime.RecordGet(orig, "month").IntVal
		clone.second = gopurs_runtime.RecordGet(orig, "second").IntVal
		clone.year = gopurs_runtime.RecordGet(orig, "year").IntVal
		return clone
	}()
	return out
}

func Rebox_Control_Bind_3094389156_3516139266(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]]{}
	out.V0 = Rebox_Control_Bind_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Control_Bind_3094389156_4151030883(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_DateTime_DateTime] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_DateTime_DateTime]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](in.V0)
	return out
}

func Rebox_Control_Bind_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[string]{}
	out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Control_Bind_3302008615_3094389156(in *Constructor_Data_Maybe_Just[struct {
	day         int64
	hour        int64
	millisecond int64
	minute      int64
	month       int64
	second      int64
	year        int64
	/*gopurs:record-v1:7:006400610079,0068006f00750072,006d0069006c006c0069007300650063006f006e0064,006d0069006e007500740065,006d006f006e00740068,007300650063006f006e0064,0079006500610072*/
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict([]string{"day", "hour", "millisecond", "minute", "month", "second", "year"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.day), gopurs_runtime.Int(orig.hour), gopurs_runtime.Int(orig.millisecond), gopurs_runtime.Int(orig.minute), gopurs_runtime.Int(orig.month), gopurs_runtime.Int(orig.second), gopurs_runtime.Int(orig.year)})
	}()
	return out
}

func Rebox_Control_Bind_3459698345_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[bool]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Bind_3516139266_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Control_Bind_2280409795_3094389156(in.V0))}
	return out
}

func Rebox_Control_Bind_3559868223_177467250(in *Constructor_Control_Bind_Discard[uint32]) *Constructor_Control_Bind_Discard[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Discard[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Bind_3741347833_1303115796(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Apply_Apply[uint32])(unsafe.Pointer(in))
}

func Rebox_Control_Bind_4014717844_2748095225(in *Constructor_Control_Bind_Bind[uint32]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Control_Bind_Bind[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Bind_4151030883_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_DateTime_DateTime]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Bind_48201874_1277236813(in *Constructor_Data_Argonaut_Decode_Class_DecodeJson[string]) *Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Argonaut_Decode_Class_DecodeJson[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Control_Bind_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Str(in.V0)
	return out
}

func Get_Control_Bind_arrayBind() gopurs_runtime.Value {
	return _Gopurs_Control_Bind_ArrayBind
}
