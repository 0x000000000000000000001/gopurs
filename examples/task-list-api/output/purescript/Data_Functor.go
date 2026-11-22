package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Functor_dollar_Dict gopurs_runtime.Value
var once_Data_Functor_Functor_dollar_Dict sync.Once

func Get_Data_Functor_Functor_dollar_Dict() gopurs_runtime.Value {
	once_Data_Functor_Functor_dollar_Dict.Do(func() {
		cache_Data_Functor_Functor_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Functor_dollar_Dict(func() struct {
				go__map gopurs_runtime.Value
				/*gopurs:record-v1:1:006d00610070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					go__map gopurs_runtime.Value
					/*gopurs:record-v1:1:006d00610070*/
				}{}
				clone.go__map = gopurs_runtime.RecordGet(orig, "map")
				return clone
			}()))}
		})
	})
	return cache_Data_Functor_Functor_dollar_Dict
}

var cache_Data_Functor_go__map gopurs_runtime.Value
var once_Data_Functor_go__map sync.Once

func Get_Data_Functor_go__map() gopurs_runtime.Value {
	once_Data_Functor_go__map.Do(func() {
		cache_Data_Functor_go__map = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Functor_go__map
}

var cache_Data_Functor_map__3659863366 gopurs_runtime.Value
var once_Data_Functor_map__3659863366 sync.Once

func Get_Data_Functor_map__3659863366() gopurs_runtime.Value {
	once_Data_Functor_map__3659863366.Do(func() {
		cache_Data_Functor_map__3659863366 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__3659863366(__eta0_0_box, __eta1_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Functor_map__3659863366
}

var cache_Data_Functor_map__1891958886 gopurs_runtime.Value
var once_Data_Functor_map__1891958886 sync.Once

func Get_Data_Functor_map__1891958886() gopurs_runtime.Value {
	once_Data_Functor_map__1891958886.Do(func() {
		cache_Data_Functor_map__1891958886 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__1891958886(__eta0_0_box, __eta1_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Functor_map__1891958886
}

var cache_Data_Functor_map__1962508006 gopurs_runtime.Value
var once_Data_Functor_map__1962508006 sync.Once

func Get_Data_Functor_map__1962508006() gopurs_runtime.Value {
	once_Data_Functor_map__1962508006.Do(func() {
		cache_Data_Functor_map__1962508006 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__1962508006(__eta0_0_box, __eta1_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Functor_map__1962508006
}

var cache_Data_Functor_map__3477210790 gopurs_runtime.Value
var once_Data_Functor_map__3477210790 sync.Once

func Get_Data_Functor_map__3477210790() gopurs_runtime.Value {
	once_Data_Functor_map__3477210790.Do(func() {
		cache_Data_Functor_map__3477210790 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__3477210790(__eta0_0_box, __eta1_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Functor_map__3477210790
}

var cache_Data_Functor_map__1881518470 gopurs_runtime.Value
var once_Data_Functor_map__1881518470 sync.Once

func Get_Data_Functor_map__1881518470() gopurs_runtime.Value {
	once_Data_Functor_map__1881518470.Do(func() {
		cache_Data_Functor_map__1881518470 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__1881518470(__eta0_0_box, __eta1_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Functor_map__1881518470
}

var cache_Data_Functor_map__2013991078 gopurs_runtime.Value
var once_Data_Functor_map__2013991078 sync.Once

func Get_Data_Functor_map__2013991078() gopurs_runtime.Value {
	once_Data_Functor_map__2013991078.Do(func() {
		cache_Data_Functor_map__2013991078 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__2013991078(__eta0_0_box, __eta1_1_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Data_Functor_map__2013991078
}

var cache_Data_Functor_map__1601533770 gopurs_runtime.Value
var once_Data_Functor_map__1601533770 sync.Once

func Get_Data_Functor_map__1601533770() gopurs_runtime.Value {
	once_Data_Functor_map__1601533770.Do(func() {
		cache_Data_Functor_map__1601533770 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__1601533770(__eta0_0_box, Rebox_Data_Functor_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Functor_map__1601533770
}

var cache_Data_Functor_map__1787310276 gopurs_runtime.Value
var once_Data_Functor_map__1787310276 sync.Once

func Get_Data_Functor_map__1787310276() gopurs_runtime.Value {
	once_Data_Functor_map__1787310276.Do(func() {
		cache_Data_Functor_map__1787310276 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__1787310276(__eta_norm_1_0_box, Rebox_Data_Functor_3094389156_2304530460(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Functor_map__1787310276
}

var cache_Data_Functor_map__1167843562 gopurs_runtime.Value
var once_Data_Functor_map__1167843562 sync.Once

func Get_Data_Functor_map__1167843562() gopurs_runtime.Value {
	once_Data_Functor_map__1167843562.Do(func() {
		cache_Data_Functor_map__1167843562 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__1167843562(__eta0_0_box, Rebox_Data_Functor_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Functor_map__1167843562
}

var cache_Data_Functor_map__3216300106 gopurs_runtime.Value
var once_Data_Functor_map__3216300106 sync.Once

func Get_Data_Functor_map__3216300106() gopurs_runtime.Value {
	once_Data_Functor_map__3216300106.Do(func() {
		cache_Data_Functor_map__3216300106 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__3216300106(__eta0_0_box, Rebox_Data_Functor_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Functor_map__3216300106
}

var cache_Data_Functor_map__1221272394 gopurs_runtime.Value
var once_Data_Functor_map__1221272394 sync.Once

func Get_Data_Functor_map__1221272394() gopurs_runtime.Value {
	once_Data_Functor_map__1221272394.Do(func() {
		cache_Data_Functor_map__1221272394 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__1221272394(__eta0_0_box, Rebox_Data_Functor_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Functor_map__1221272394
}

var cache_Data_Functor_map__1798682762 gopurs_runtime.Value
var once_Data_Functor_map__1798682762 sync.Once

func Get_Data_Functor_map__1798682762() gopurs_runtime.Value {
	once_Data_Functor_map__1798682762.Do(func() {
		cache_Data_Functor_map__1798682762 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__1798682762(__eta0_0_box, Rebox_Data_Functor_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Functor_map__1798682762
}

var cache_Data_Functor_map__3627393546 gopurs_runtime.Value
var once_Data_Functor_map__3627393546 sync.Once

func Get_Data_Functor_map__3627393546() gopurs_runtime.Value {
	once_Data_Functor_map__3627393546.Do(func() {
		cache_Data_Functor_map__3627393546 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__3627393546(__eta0_0_box, Rebox_Data_Functor_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Functor_map__3627393546
}

var cache_Data_Functor_map__1382837546 gopurs_runtime.Value
var once_Data_Functor_map__1382837546 sync.Once

func Get_Data_Functor_map__1382837546() gopurs_runtime.Value {
	once_Data_Functor_map__1382837546.Do(func() {
		cache_Data_Functor_map__1382837546 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__1382837546(__eta0_0_box, Rebox_Data_Functor_3094389156_1387998409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta1_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Functor_map__1382837546
}

var cache_Data_Functor_map__1317168042 gopurs_runtime.Value
var once_Data_Functor_map__1317168042 sync.Once

func Get_Data_Functor_map__1317168042() gopurs_runtime.Value {
	once_Data_Functor_map__1317168042.Do(func() {
		cache_Data_Functor_map__1317168042 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__1317168042(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__1317168042
}

var cache_Data_Functor_map__2345995242 gopurs_runtime.Value
var once_Data_Functor_map__2345995242 sync.Once

func Get_Data_Functor_map__2345995242() gopurs_runtime.Value {
	once_Data_Functor_map__2345995242.Do(func() {
		cache_Data_Functor_map__2345995242 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__2345995242(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__2345995242
}

var cache_Data_Functor_map__2883899760 gopurs_runtime.Value
var once_Data_Functor_map__2883899760 sync.Once

func Get_Data_Functor_map__2883899760() gopurs_runtime.Value {
	once_Data_Functor_map__2883899760.Do(func() {
		cache_Data_Functor_map__2883899760 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__2883899760(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__2883899760
}

var cache_Data_Functor_map__3631854384 gopurs_runtime.Value
var once_Data_Functor_map__3631854384 sync.Once

func Get_Data_Functor_map__3631854384() gopurs_runtime.Value {
	once_Data_Functor_map__3631854384.Do(func() {
		cache_Data_Functor_map__3631854384 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__3631854384(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__3631854384
}

var cache_Data_Functor_map__3654485947 gopurs_runtime.Value
var once_Data_Functor_map__3654485947 sync.Once

func Get_Data_Functor_map__3654485947() gopurs_runtime.Value {
	once_Data_Functor_map__3654485947.Do(func() {
		cache_Data_Functor_map__3654485947 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__3654485947(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__3654485947
}

var cache_Data_Functor_map__4150859291 gopurs_runtime.Value
var once_Data_Functor_map__4150859291 sync.Once

func Get_Data_Functor_map__4150859291() gopurs_runtime.Value {
	once_Data_Functor_map__4150859291.Do(func() {
		cache_Data_Functor_map__4150859291 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__4150859291(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__4150859291
}

var cache_Data_Functor_map__447590619 gopurs_runtime.Value
var once_Data_Functor_map__447590619 sync.Once

func Get_Data_Functor_map__447590619() gopurs_runtime.Value {
	once_Data_Functor_map__447590619.Do(func() {
		cache_Data_Functor_map__447590619 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__447590619(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__447590619
}

var cache_Data_Functor_map__2042768475 gopurs_runtime.Value
var once_Data_Functor_map__2042768475 sync.Once

func Get_Data_Functor_map__2042768475() gopurs_runtime.Value {
	once_Data_Functor_map__2042768475.Do(func() {
		cache_Data_Functor_map__2042768475 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__2042768475(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__2042768475
}

var cache_Data_Functor_map__1260122395 gopurs_runtime.Value
var once_Data_Functor_map__1260122395 sync.Once

func Get_Data_Functor_map__1260122395() gopurs_runtime.Value {
	once_Data_Functor_map__1260122395.Do(func() {
		cache_Data_Functor_map__1260122395 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__1260122395(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__1260122395
}

var cache_Data_Functor_map__2482553499 gopurs_runtime.Value
var once_Data_Functor_map__2482553499 sync.Once

func Get_Data_Functor_map__2482553499() gopurs_runtime.Value {
	once_Data_Functor_map__2482553499.Do(func() {
		cache_Data_Functor_map__2482553499 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__2482553499(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__2482553499
}

var cache_Data_Functor_map__612405524 gopurs_runtime.Value
var once_Data_Functor_map__612405524 sync.Once

func Get_Data_Functor_map__612405524() gopurs_runtime.Value {
	once_Data_Functor_map__612405524.Do(func() {
		cache_Data_Functor_map__612405524 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__612405524(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Data_Functor_map__612405524
}

var cache_Data_Functor_map__3407405013 gopurs_runtime.Value
var once_Data_Functor_map__3407405013 sync.Once

func Get_Data_Functor_map__3407405013() gopurs_runtime.Value {
	once_Data_Functor_map__3407405013.Do(func() {
		cache_Data_Functor_map__3407405013 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__3407405013(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Data_Functor_map__3407405013
}

var cache_Data_Functor_map__477560180 gopurs_runtime.Value
var once_Data_Functor_map__477560180 sync.Once

func Get_Data_Functor_map__477560180() gopurs_runtime.Value {
	once_Data_Functor_map__477560180.Do(func() {
		cache_Data_Functor_map__477560180 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__477560180(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Data_Functor_map__477560180
}

var cache_Data_Functor_map__1813422419 gopurs_runtime.Value
var once_Data_Functor_map__1813422419 sync.Once

func Get_Data_Functor_map__1813422419() gopurs_runtime.Value {
	once_Data_Functor_map__1813422419.Do(func() {
		cache_Data_Functor_map__1813422419 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__1813422419(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Data_Functor_map__1813422419
}

var cache_Data_Functor_map__3041220533 gopurs_runtime.Value
var once_Data_Functor_map__3041220533 sync.Once

func Get_Data_Functor_map__3041220533() gopurs_runtime.Value {
	once_Data_Functor_map__3041220533.Do(func() {
		cache_Data_Functor_map__3041220533 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__3041220533(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Data_Functor_map__3041220533
}

var cache_Data_Functor_map__855557813 gopurs_runtime.Value
var once_Data_Functor_map__855557813 sync.Once

func Get_Data_Functor_map__855557813() gopurs_runtime.Value {
	once_Data_Functor_map__855557813.Do(func() {
		cache_Data_Functor_map__855557813 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__855557813(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Data_Functor_map__855557813
}

var cache_Data_Functor_map__1860138811 gopurs_runtime.Value
var once_Data_Functor_map__1860138811 sync.Once

func Get_Data_Functor_map__1860138811() gopurs_runtime.Value {
	once_Data_Functor_map__1860138811.Do(func() {
		cache_Data_Functor_map__1860138811 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__1860138811(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__1860138811
}

var cache_Data_Functor_map__3707757371 gopurs_runtime.Value
var once_Data_Functor_map__3707757371 sync.Once

func Get_Data_Functor_map__3707757371() gopurs_runtime.Value {
	once_Data_Functor_map__3707757371.Do(func() {
		cache_Data_Functor_map__3707757371 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__3707757371(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__3707757371
}

var cache_Data_Functor_map__2772546662 gopurs_runtime.Value
var once_Data_Functor_map__2772546662 sync.Once

func Get_Data_Functor_map__2772546662() gopurs_runtime.Value {
	once_Data_Functor_map__2772546662.Do(func() {
		cache_Data_Functor_map__2772546662 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__2772546662(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__2772546662
}

var cache_Data_Functor_map__1137213094 gopurs_runtime.Value
var once_Data_Functor_map__1137213094 sync.Once

func Get_Data_Functor_map__1137213094() gopurs_runtime.Value {
	once_Data_Functor_map__1137213094.Do(func() {
		cache_Data_Functor_map__1137213094 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__1137213094(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__1137213094
}

var cache_Data_Functor_functorArray gopurs_runtime.Value
var once_Data_Functor_functorArray sync.Once

func Get_Data_Functor_functorArray() gopurs_runtime.Value {
	once_Data_Functor_functorArray.Do(func() {
		cache_Data_Functor_functorArray = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, Get_Data_Functor_arrayMap()}))}
	})
	return cache_Data_Functor_functorArray
}

var cache_Data_Functor_map__4008632948 gopurs_runtime.Value
var once_Data_Functor_map__4008632948 sync.Once

func Get_Data_Functor_map__4008632948() gopurs_runtime.Value {
	once_Data_Functor_map__4008632948.Do(func() {
		cache_Data_Functor_map__4008632948 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__4008632948(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__4008632948
}

var cache_Data_Functor_map__1207317140 gopurs_runtime.Value
var once_Data_Functor_map__1207317140 sync.Once

func Get_Data_Functor_map__1207317140() gopurs_runtime.Value {
	once_Data_Functor_map__1207317140.Do(func() {
		cache_Data_Functor_map__1207317140 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__1207317140(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_map__1207317140
}

var cache_Data_Functor_functorFn gopurs_runtime.Value
var once_Data_Functor_functorFn sync.Once

func Get_Data_Functor_functorFn() gopurs_runtime.Value {
	once_Data_Functor_functorFn.Do(func() {
		cache_Data_Functor_functorFn = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()).V0}))}
	})
	return cache_Data_Functor_functorFn
}

var cache_Data_Functor_map__4275211926 gopurs_runtime.Value
var once_Data_Functor_map__4275211926 sync.Once

func Get_Data_Functor_map__4275211926() gopurs_runtime.Value {
	once_Data_Functor_map__4275211926.Do(func() {
		cache_Data_Functor_map__4275211926 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_map__4275211926(__eta0_0_box, __eta1_1_box, x_2_box)
		})
	})
	return cache_Data_Functor_map__4275211926
}

var cache_Data_Functor_mapFlipped gopurs_runtime.Value
var once_Data_Functor_mapFlipped sync.Once

func Get_Data_Functor_mapFlipped() gopurs_runtime.Value {
	once_Data_Functor_mapFlipped.Do(func() {
		cache_Data_Functor_mapFlipped = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_mapFlipped(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
		})
	})
	return cache_Data_Functor_mapFlipped
}

var cache_Data_Functor_mapFlipped__4096616251 gopurs_runtime.Value
var once_Data_Functor_mapFlipped__4096616251 sync.Once

func Get_Data_Functor_mapFlipped__4096616251() gopurs_runtime.Value {
	once_Data_Functor_mapFlipped__4096616251.Do(func() {
		cache_Data_Functor_mapFlipped__4096616251 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_mapFlipped__4096616251(__eta0_0_box, __eta1_1_box)
		})
	})
	return cache_Data_Functor_mapFlipped__4096616251
}

var cache_Data_Functor_void gopurs_runtime.Value
var once_Data_Functor_void sync.Once

func Get_Data_Functor_void() gopurs_runtime.Value {
	once_Data_Functor_void.Do(func() {
		cache_Data_Functor_void = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box))
		})
	})
	return cache_Data_Functor_void
}

var cache_Data_Functor_voidLeft gopurs_runtime.Value
var once_Data_Functor_voidLeft sync.Once

func Get_Data_Functor_voidLeft() gopurs_runtime.Value {
	once_Data_Functor_voidLeft.Do(func() {
		cache_Data_Functor_voidLeft = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_voidLeft(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, x_2_box)
		})
	})
	return cache_Data_Functor_voidLeft
}

var cache_Data_Functor_voidLeft__3176963436 gopurs_runtime.Value
var once_Data_Functor_voidLeft__3176963436 sync.Once

func Get_Data_Functor_voidLeft__3176963436() gopurs_runtime.Value {
	once_Data_Functor_voidLeft__3176963436.Do(func() {
		cache_Data_Functor_voidLeft__3176963436 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_voidLeft__3176963436(__eta0_0_box, __eta1_1_box.StrVal())
		})
	})
	return cache_Data_Functor_voidLeft__3176963436
}

var cache_Data_Functor_voidRight gopurs_runtime.Value
var once_Data_Functor_voidRight sync.Once

func Get_Data_Functor_voidRight() gopurs_runtime.Value {
	once_Data_Functor_voidRight.Do(func() {
		cache_Data_Functor_voidRight = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_voidRight(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), x_1_box)
		})
	})
	return cache_Data_Functor_voidRight
}

var cache_Data_Functor_functorProxy gopurs_runtime.Value
var once_Data_Functor_functorProxy sync.Once

func Get_Data_Functor_functorProxy() gopurs_runtime.Value {
	once_Data_Functor_functorProxy.Do(func() {
		cache_Data_Functor_functorProxy = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_562069347_2812149806((&Constructor_Data_Functor_Functor[uint32]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
		})})))}
	})
	return cache_Data_Functor_functorProxy
}

var cache_Data_Functor_flap gopurs_runtime.Value
var once_Data_Functor_flap sync.Once

func Get_Data_Functor_flap() gopurs_runtime.Value {
	once_Data_Functor_flap.Do(func() {
		cache_Data_Functor_flap = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, ff_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Functor_flap(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), ff_1_box, x_2_box)
		})
	})
	return cache_Data_Functor_flap
}

type Constructor_Data_Functor_Functor[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00460075006e00630074006f0072:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00460075006e00630074006f0072005f00460075006e00630074006f0072
func init() {
	gopurs_runtime.StructGetters[929368378] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "map":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Data_Functor_Functor: " + key)
		}
	}
}

func Call_Data_Functor_Functor_dollar_Dict(x_0_loop struct {
	go__map gopurs_runtime.Value
	/*gopurs:record-v1:1:006d00610070*/
}) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	var x_0 struct {
		go__map gopurs_runtime.Value
		/*gopurs:record-v1:1:006d00610070*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("map", orig.go__map)
	}())
}

func Call_Data_Functor_go__map(dict_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Data_Functor_map__3659863366(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
map__3659863366:
	for {
		if false {
			continue map__3659863366
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), __eta0_0, __eta1_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__1891958886(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
map__1891958886:
	for {
		if false {
			continue map__1891958886
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), __eta0_0, __eta1_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__1962508006(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
map__1962508006:
	for {
		if false {
			continue map__1962508006
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), __eta0_0, __eta1_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__3477210790(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
map__3477210790:
	for {
		if false {
			continue map__3477210790
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), __eta0_0, __eta1_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__1881518470(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
map__1881518470:
	for {
		if false {
			continue map__1881518470
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), __eta0_0, __eta1_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__2013991078(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 bool
} {
map__2013991078:
	for {
		if false {
			continue map__2013991078
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 gopurs_runtime.Value
			V2 bool
		} { _v := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), __eta0_0, __eta1_1); if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__1601533770(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
map__1601533770:
	for {
		if false {
			continue map__1601533770
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, __eta0_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_2280409795_3094389156(__eta1_1))})))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__1787310276(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[uint32, float64]]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
map__1787310276:
	for {
		if false {
			continue map__1787310276
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[uint32, float64]] = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, __eta_norm_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_2304530460_3094389156(__eta_norm_0_1))}); if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__1167843562(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop *Constructor_Data_Maybe_Just[int64]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
map__1167843562:
	for {
		if false {
			continue map__1167843562
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[int64] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_2280409795_3094389156(Rebox_Data_Functor_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, __eta0_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_1170268447_3094389156(__eta1_1))})))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__3216300106(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop *Constructor_Data_Maybe_Just[int64]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
map__3216300106:
	for {
		if false {
			continue map__3216300106
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[int64] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, __eta0_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_1170268447_3094389156(__eta1_1))})))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__1221272394(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop *Constructor_Data_Maybe_Just[int64]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
map__1221272394:
	for {
		if false {
			continue map__1221272394
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[int64] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, __eta0_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_1170268447_3094389156(__eta1_1))})))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__1798682762(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop *Constructor_Data_Maybe_Just[int64]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
map__1798682762:
	for {
		if false {
			continue map__1798682762
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[int64] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, __eta0_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_1170268447_3094389156(__eta1_1))})))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__3627393546(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop *Constructor_Data_Maybe_Just[int64]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
map__3627393546:
	for {
		if false {
			continue map__3627393546
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[int64] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_1170268447_3094389156(Rebox_Data_Functor_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, __eta0_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_1170268447_3094389156(__eta1_1))})))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__1382837546(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop *Constructor_Data_Maybe_Just[struct {
	head int64
	tail string
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
}]) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
map__1382837546:
	for {
		if false {
			continue map__1382837546
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 *Constructor_Data_Maybe_Just[struct {
			head int64
			tail string
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}] = __eta1_1_loop
		_ = __eta1_1
		return func() struct {
			V0 gopurs_runtime.Value
			V1 bool
		} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_639395693_3094389156(Rebox_Data_Functor_3094389156_639395693(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, __eta0_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_1387998409_3094389156(__eta1_1))})))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Call_Data_Functor_map__1317168042(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1317168042:
	for {
		if false {
			continue map__1317168042
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_Aff_functorAff()).V0, __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__2345995242(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2345995242:
	for {
		if false {
			continue map__2345995242
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_Aff_functorAff()).V0, __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__2883899760(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2883899760:
	for {
		if false {
			continue map__2883899760
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_monadErrorAff()).V0, gopurs_runtime.Value{}), "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__3631854384(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__3631854384:
	for {
		if false {
			continue map__3631854384
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Effect_Aff_monadErrorAff()).V0, gopurs_runtime.Value{}), "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__3654485947(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__3654485947:
	for {
		if false {
			continue map__3654485947
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__4150859291(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__4150859291:
	for {
		if false {
			continue map__4150859291
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__447590619(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__447590619:
	for {
		if false {
			continue map__447590619
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__2042768475(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2042768475:
	for {
		if false {
			continue map__2042768475
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__1260122395(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1260122395:
	for {
		if false {
			continue map__1260122395
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__2482553499(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2482553499:
	for {
		if false {
			continue map__2482553499
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__612405524(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__612405524:
	for {
		if false {
			continue map__612405524
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, Get_Node_Buffer_Immutable_toArrayBuffer(), __eta_norm_0_1)
	}
}

func Call_Data_Functor_map__3407405013(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__3407405013:
	for {
		if false {
			continue map__3407405013
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Data_Functor_map__477560180(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__477560180:
	for {
		if false {
			continue map__477560180
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, Get_Node_Buffer_Immutable_toArray(), __eta_norm_0_1)
	}
}

func Call_Data_Functor_map__1813422419(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1813422419:
	for {
		if false {
			continue map__1813422419
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, Get_Node_Buffer_Immutable_size(), __eta_norm_0_1)
	}
}

func Call_Data_Functor_map__3041220533(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__3041220533:
	for {
		if false {
			continue map__3041220533
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Data_Functor_map__855557813(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__855557813:
	for {
		if false {
			continue map__855557813
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Data_Functor_map__1860138811(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1860138811:
	for {
		if false {
			continue map__1860138811
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__3707757371(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__3707757371:
	for {
		if false {
			continue map__3707757371
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__2772546662(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2772546662:
	for {
		if false {
			continue map__2772546662
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()).V0, gopurs_runtime.Value{}), "map"), __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__1137213094(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1137213094:
	for {
		if false {
			continue map__1137213094
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Effect_applyEffect()).V0, gopurs_runtime.Value{}), "map"), __eta0_0, __eta1_1)
	}
}

func Call_Data_Functor_map__4008632948(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__4008632948:
	for {
		if false {
			continue map__4008632948
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		arrayIntrinsic_arg_0 := __eta0_0
		_ = arrayIntrinsic_arg_0
		_ = arrayIntrinsic_arg_0
		arrayIntrinsic_arg_1 := __eta1_1
		_ = arrayIntrinsic_arg_1
		_ = arrayIntrinsic_arg_1
		return gopurs_runtime.Array(func() []gopurs_runtime.Value {
			arr_val_arrayMap0 := arrayIntrinsic_arg_1
			_ = arr_val_arrayMap0
			_ = arr_val_arrayMap0
			arrayIntrinsic_callback_2 := arrayIntrinsic_arg_0
			_ = arrayIntrinsic_callback_2
			arr_go_arrayMap0 := (*[]gopurs_runtime.Value)(arr_val_arrayMap0.UnsafePtr)
			_ = arr_go_arrayMap0
			res_go_arrayMap0 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap0))
			_ = res_go_arrayMap0
			for i_arrayMap0, v_arrayMap0 := range *arr_go_arrayMap0 {
				res_go_arrayMap0[i_arrayMap0] = gopurs_runtime.Apply(arrayIntrinsic_callback_2, v_arrayMap0)
			}
			return res_go_arrayMap0
		}())
	}
}

func Call_Data_Functor_map__1207317140(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1207317140:
	for {
		if false {
			continue map__1207317140
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		arrayIntrinsic_arg_0 := __eta0_0
		_ = arrayIntrinsic_arg_0
		_ = arrayIntrinsic_arg_0
		arrayIntrinsic_arg_1 := __eta1_1
		_ = arrayIntrinsic_arg_1
		_ = arrayIntrinsic_arg_1
		return gopurs_runtime.Array(func() []gopurs_runtime.Value {
			arr_val_arrayMap0 := arrayIntrinsic_arg_1
			_ = arr_val_arrayMap0
			_ = arr_val_arrayMap0
			arrayIntrinsic_callback_2 := arrayIntrinsic_arg_0
			_ = arrayIntrinsic_callback_2
			arr_go_arrayMap0 := (*[]gopurs_runtime.Value)(arr_val_arrayMap0.UnsafePtr)
			_ = arr_go_arrayMap0
			res_go_arrayMap0 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap0))
			_ = res_go_arrayMap0
			for i_arrayMap0, v_arrayMap0 := range *arr_go_arrayMap0 {
				res_go_arrayMap0[i_arrayMap0] = gopurs_runtime.Apply(arrayIntrinsic_callback_2, v_arrayMap0)
			}
			return res_go_arrayMap0
		}())
	}
}

func Call_Data_Functor_map__4275211926(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__4275211926:
	for {
		if false {
			continue map__4275211926
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		var x_2 gopurs_runtime.Value = x_2_loop
		_ = x_2
		return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
	}
}

func Call_Data_Functor_mapFlipped(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var fa_1 gopurs_runtime.Value = fa_1_loop
	_ = fa_1
	var f_2 gopurs_runtime.Value = f_2_loop
	_ = f_2
	return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_Data_Functor_mapFlipped__4096616251(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapFlipped__4096616251:
	for {
		if false {
			continue mapFlipped__4096616251
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta1_1, __eta0_0)
	}
}

func Call_Data_Functor_void(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Apply(dictFunctor_0.V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return Get_Data_Unit_unit()
	}))
}

func Call_Data_Functor_voidLeft(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var x_2 gopurs_runtime.Value = x_2_loop
	_ = x_2
	return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_2
	}), f_1)
}

func Call_Data_Functor_voidLeft__3176963436(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop string) gopurs_runtime.Value {
voidLeft__3176963436:
	for {
		if false {
			continue voidLeft__3176963436
		}
		var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 string = __eta1_1_loop
		_ = __eta1_1
		arrayIntrinsic_arg_0 := func(v_2 gopurs_runtime.Value) string {
			return __eta1_1
		}
		_ = arrayIntrinsic_arg_0
		_ = arrayIntrinsic_arg_0
		arrayIntrinsic_arg_1 := __eta0_0
		_ = arrayIntrinsic_arg_1
		_ = arrayIntrinsic_arg_1
		return gopurs_runtime.Array(*((*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
			arr := func() []string {
				arr_val_arrayMap0 := arrayIntrinsic_arg_1
				_ = arr_val_arrayMap0
				_ = arr_val_arrayMap0
				arr_go_arrayMap0 := (*[]gopurs_runtime.Value)(arr_val_arrayMap0.UnsafePtr)
				_ = arr_go_arrayMap0
				res_go_arrayMap0 := make([]string, len(*arr_go_arrayMap0))
				_ = res_go_arrayMap0
				for i_arrayMap0, v_arrayMap0 := range *arr_go_arrayMap0 {
					res_go_arrayMap0[i_arrayMap0] = arrayIntrinsic_arg_0(v_arrayMap0)
				}
				return res_go_arrayMap0
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Str(v)
			}
			return gopurs_runtime.Array(boxed)
		}().UnsafePtr)))
	}
}

func Call_Data_Functor_voidRight(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(dictFunctor_0.V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return x_1
	}))
}

func Call_Data_Functor_flap(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], ff_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	var ff_1 gopurs_runtime.Value = ff_1_loop
	_ = ff_1
	var x_2 gopurs_runtime.Value = x_2_loop
	_ = x_2
	return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(f_3, x_2)
	}), ff_1)
}

func Rebox_Data_Functor_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Functor_138441832_2169642284(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, string] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[int64, string]{}
	out.V0 = in.V0.IntVal
	out.V1 = in.V1.StrVal()
	return out
}

func Rebox_Data_Functor_138441832_3132786365(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[uint32, float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[uint32, float64]{}
	out.V0 = uint32(in.V0.IntVal)
	out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_Functor_1387998409_3094389156(in *Constructor_Data_Maybe_Just[struct {
	head int64
	tail string
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = func() gopurs_runtime.Value {
		orig := in.V0
		_ = orig
		return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Str(orig.tail))
	}()
	return out
}

func Rebox_Data_Functor_2169642284_138441832(in *Constructor_Data_Tuple_Tuple[int64, string]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	out.V1 = gopurs_runtime.Str(in.V1)
	return out
}

func Rebox_Data_Functor_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Functor_2304530460_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_3132786365_138441832(in.V0))}
	return out
}

func Rebox_Data_Functor_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Functor_3094389156_1387998409(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct {
	head int64
	tail string
	/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
}] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[struct {
		head int64
		tail string
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	}]{}
	out.V0 = func() struct {
		head int64
		tail string
		/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
	} {
		orig := in.V0
		_ = orig
		clone := struct {
			head int64
			tail string
			/*gopurs:record-v1:2:0068006500610064,007400610069006c*/
		}{}
		clone.head = gopurs_runtime.RecordGet(orig, "head").IntVal
		clone.tail = gopurs_runtime.RecordGet(orig, "tail").StrVal()
		return clone
	}()
	return out
}

func Rebox_Data_Functor_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
	out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}

func Rebox_Data_Functor_3094389156_2304530460(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[uint32, float64]]{}
	out.V0 = Rebox_Data_Functor_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_Functor_3094389156_639395693(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, string]] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, string]]{}
	out.V0 = Rebox_Data_Functor_138441832_2169642284(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_Functor_3132786365_138441832(in *Constructor_Data_Tuple_Tuple[uint32, float64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Functor_562069347_2812149806(in *Constructor_Data_Functor_Functor[uint32]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Functor_639395693_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, string]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_2169642284_138441832(in.V0))}
	return out
}

func Get_Data_Functor_arrayMap() gopurs_runtime.Value {
	return _Gopurs_Data_Functor_ArrayMap
}
