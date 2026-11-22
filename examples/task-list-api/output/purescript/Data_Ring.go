package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Ring_semiringRecord gopurs_runtime.Value
var once_Data_Ring_semiringRecord sync.Once

func Get_Data_Ring_semiringRecord() gopurs_runtime.Value {
	once_Data_Ring_semiringRecord.Do(func() {
		cache_Data_Ring_semiringRecord = gopurs_runtime.Apply(Get_Data_Semiring_semiringRecord(), gopurs_runtime.Value{})
	})
	return cache_Data_Ring_semiringRecord
}

var cache_Data_Ring_RingRecord_dollar_Dict gopurs_runtime.Value
var once_Data_Ring_RingRecord_dollar_Dict sync.Once

func Get_Data_Ring_RingRecord_dollar_Dict() gopurs_runtime.Value {
	once_Data_Ring_RingRecord_dollar_Dict.Do(func() {
		cache_Data_Ring_RingRecord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4246880791, UnsafePtr: unsafe.Pointer(Call_Data_Ring_RingRecord_dollar_Dict(func() struct {
				SemiringRecord0 gopurs_runtime.Value
				subRecord       gopurs_runtime.Value
				/*gopurs:record-v1:2:00530065006d006900720069006e0067005200650063006f007200640030,007300750062005200650063006f00720064*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					SemiringRecord0 gopurs_runtime.Value
					subRecord       gopurs_runtime.Value
					/*gopurs:record-v1:2:00530065006d006900720069006e0067005200650063006f007200640030,007300750062005200650063006f00720064*/
				}{}
				clone.SemiringRecord0 = gopurs_runtime.RecordGet(orig, "SemiringRecord0")
				clone.subRecord = gopurs_runtime.RecordGet(orig, "subRecord")
				return clone
			}()))}
		})
	})
	return cache_Data_Ring_RingRecord_dollar_Dict
}

var cache_Data_Ring_Ring_dollar_Dict gopurs_runtime.Value
var once_Data_Ring_Ring_dollar_Dict sync.Once

func Get_Data_Ring_Ring_dollar_Dict() gopurs_runtime.Value {
	once_Data_Ring_Ring_dollar_Dict.Do(func() {
		cache_Data_Ring_Ring_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Call_Data_Ring_Ring_dollar_Dict(func() struct {
				Semiring0 gopurs_runtime.Value
				sub       gopurs_runtime.Value
				/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semiring0 gopurs_runtime.Value
					sub       gopurs_runtime.Value
					/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
				}{}
				clone.Semiring0 = gopurs_runtime.RecordGet(orig, "Semiring0")
				clone.sub = gopurs_runtime.RecordGet(orig, "sub")
				return clone
			}()))}
		})
	})
	return cache_Data_Ring_Ring_dollar_Dict
}

var cache_Data_Ring_Ring_dollar_Dict__812180516 gopurs_runtime.Value
var once_Data_Ring_Ring_dollar_Dict__812180516 sync.Once

func Get_Data_Ring_Ring_dollar_Dict__812180516() gopurs_runtime.Value {
	once_Data_Ring_Ring_dollar_Dict__812180516.Do(func() {
		cache_Data_Ring_Ring_dollar_Dict__812180516 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_2370240579_3060961102(Call_Data_Ring_Ring_dollar_Dict__812180516(func() struct {
				Semiring0 gopurs_runtime.Value
				sub       gopurs_runtime.Value
				/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semiring0 gopurs_runtime.Value
					sub       gopurs_runtime.Value
					/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
				}{}
				clone.Semiring0 = gopurs_runtime.RecordGet(orig, "Semiring0")
				clone.sub = gopurs_runtime.RecordGet(orig, "sub")
				return clone
			}())))}
		})
	})
	return cache_Data_Ring_Ring_dollar_Dict__812180516
}

var cache_Data_Ring_Ring_dollar_Dict__34472189 gopurs_runtime.Value
var once_Data_Ring_Ring_dollar_Dict__34472189 sync.Once

func Get_Data_Ring_Ring_dollar_Dict__34472189() gopurs_runtime.Value {
	once_Data_Ring_Ring_dollar_Dict__34472189.Do(func() {
		cache_Data_Ring_Ring_dollar_Dict__34472189 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_2534126933_3060961102(Call_Data_Ring_Ring_dollar_Dict__34472189(func() struct {
				Semiring0 gopurs_runtime.Value
				sub       gopurs_runtime.Value
				/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semiring0 gopurs_runtime.Value
					sub       gopurs_runtime.Value
					/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
				}{}
				clone.Semiring0 = gopurs_runtime.RecordGet(orig, "Semiring0")
				clone.sub = gopurs_runtime.RecordGet(orig, "sub")
				return clone
			}())))}
		})
	})
	return cache_Data_Ring_Ring_dollar_Dict__34472189
}

var cache_Data_Ring_Ring_dollar_Dict__2257635853 gopurs_runtime.Value
var once_Data_Ring_Ring_dollar_Dict__2257635853 sync.Once

func Get_Data_Ring_Ring_dollar_Dict__2257635853() gopurs_runtime.Value {
	once_Data_Ring_Ring_dollar_Dict__2257635853.Do(func() {
		cache_Data_Ring_Ring_dollar_Dict__2257635853 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_2259598678_3060961102(Call_Data_Ring_Ring_dollar_Dict__2257635853(func() struct {
				Semiring0 gopurs_runtime.Value
				sub       gopurs_runtime.Value
				/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semiring0 gopurs_runtime.Value
					sub       gopurs_runtime.Value
					/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
				}{}
				clone.Semiring0 = gopurs_runtime.RecordGet(orig, "Semiring0")
				clone.sub = gopurs_runtime.RecordGet(orig, "sub")
				return clone
			}())))}
		})
	})
	return cache_Data_Ring_Ring_dollar_Dict__2257635853
}

var cache_Data_Ring_Ring_dollar_Dict__1193243624 gopurs_runtime.Value
var once_Data_Ring_Ring_dollar_Dict__1193243624 sync.Once

func Get_Data_Ring_Ring_dollar_Dict__1193243624() gopurs_runtime.Value {
	once_Data_Ring_Ring_dollar_Dict__1193243624.Do(func() {
		cache_Data_Ring_Ring_dollar_Dict__1193243624 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Call_Data_Ring_Ring_dollar_Dict__1193243624(func() struct {
				Semiring0 gopurs_runtime.Value
				sub       gopurs_runtime.Value
				/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Semiring0 gopurs_runtime.Value
					sub       gopurs_runtime.Value
					/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
				}{}
				clone.Semiring0 = gopurs_runtime.RecordGet(orig, "Semiring0")
				clone.sub = gopurs_runtime.RecordGet(orig, "sub")
				return clone
			}()))}
		})
	})
	return cache_Data_Ring_Ring_dollar_Dict__1193243624
}

var cache_Data_Ring_subRecord gopurs_runtime.Value
var once_Data_Ring_subRecord sync.Once

func Get_Data_Ring_subRecord() gopurs_runtime.Value {
	once_Data_Ring_subRecord.Do(func() {
		cache_Data_Ring_subRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ring_subRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Ring_subRecord
}

var cache_Data_Ring_sub gopurs_runtime.Value
var once_Data_Ring_sub sync.Once

func Get_Data_Ring_sub() gopurs_runtime.Value {
	once_Data_Ring_sub.Do(func() {
		cache_Data_Ring_sub = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ring_sub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Ring_sub
}

var cache_Data_Ring_ringInt gopurs_runtime.Value
var once_Data_Ring_ringInt sync.Once

func Get_Data_Ring_ringInt() gopurs_runtime.Value {
	once_Data_Ring_ringInt.Do(func() {
		cache_Data_Ring_ringInt = gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_2534126933_3060961102((&Constructor_Data_Ring_Ring[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_348932501_2826095630(Rebox_Data_Ring_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))}
		}), Get_Data_Ring_intSub()})))}
	})
	return cache_Data_Ring_ringInt
}

var cache_Data_Ring_sub__4157804846 gopurs_runtime.Value
var once_Data_Ring_sub__4157804846 sync.Once

func Get_Data_Ring_sub__4157804846() gopurs_runtime.Value {
	once_Data_Ring_sub__4157804846.Do(func() {
		cache_Data_Ring_sub__4157804846 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Data_Ring_sub__4157804846(__eta0_0_box.IntVal, __eta1_1_box.IntVal))
		})
	})
	return cache_Data_Ring_sub__4157804846
}

var cache_Data_Ring_ringNumber gopurs_runtime.Value
var once_Data_Ring_ringNumber sync.Once

func Get_Data_Ring_ringNumber() gopurs_runtime.Value {
	once_Data_Ring_ringNumber.Do(func() {
		cache_Data_Ring_ringNumber = gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_2259598678_3060961102((&Constructor_Data_Ring_Ring[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_602713622_2826095630(Rebox_Data_Ring_2826095630_602713622(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringNumber()))))}
		}), Get_Data_Ring_numSub()})))}
	})
	return cache_Data_Ring_ringNumber
}

var cache_Data_Ring_sub__179436318 gopurs_runtime.Value
var once_Data_Ring_sub__179436318 sync.Once

func Get_Data_Ring_sub__179436318() gopurs_runtime.Value {
	once_Data_Ring_sub__179436318.Do(func() {
		cache_Data_Ring_sub__179436318 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Data_Ring_sub__179436318(__eta0_0_box.FloatVal(), __eta1_1_box.FloatVal()))
		})
	})
	return cache_Data_Ring_sub__179436318
}

var cache_Data_Ring_ringUnit gopurs_runtime.Value
var once_Data_Ring_ringUnit sync.Once

func Get_Data_Ring_ringUnit() gopurs_runtime.Value {
	once_Data_Ring_ringUnit.Do(func() {
		cache_Data_Ring_ringUnit = gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_Ring[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringUnit()))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Unit_unit()
		})}))}
	})
	return cache_Data_Ring_ringUnit
}

var cache_Data_Ring_ringRecordNil gopurs_runtime.Value
var once_Data_Ring_ringRecordNil sync.Once

func Get_Data_Ring_ringRecordNil() gopurs_runtime.Value {
	once_Data_Ring_ringRecordNil.Do(func() {
		cache_Data_Ring_ringRecordNil = gopurs_runtime.Value{Type: 9, IntVal: 4246880791, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Semiring_semiringRecordNil()
		}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict0()
		})}))}
	})
	return cache_Data_Ring_ringRecordNil
}

var cache_Data_Ring_ringRecordCons gopurs_runtime.Value
var once_Data_Ring_ringRecordCons sync.Once

func Get_Data_Ring_ringRecordCons() gopurs_runtime.Value {
	once_Data_Ring_ringRecordCons.Do(func() {
		cache_Data_Ring_ringRecordCons = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, dictRingRecord_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ring_ringRecordCons(dictIsSymbol_0_box, _dollar___unused_1_box, dictRingRecord_2_box)
		})
	})
	return cache_Data_Ring_ringRecordCons
}

var cache_Data_Ring_ringRecord gopurs_runtime.Value
var once_Data_Ring_ringRecord sync.Once

func Get_Data_Ring_ringRecord() gopurs_runtime.Value {
	once_Data_Ring_ringRecord.Do(func() {
		cache_Data_Ring_ringRecord = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, dictRingRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ring_ringRecord(_dollar___unused_0_box, dictRingRecord_1_box)
		})
	})
	return cache_Data_Ring_ringRecord
}

var cache_Data_Ring_ringProxy gopurs_runtime.Value
var once_Data_Ring_ringProxy sync.Once

func Get_Data_Ring_ringProxy() gopurs_runtime.Value {
	once_Data_Ring_ringProxy.Do(func() {
		cache_Data_Ring_ringProxy = gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_2370240579_3060961102((&Constructor_Data_Ring_Ring[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Semiring_semiringProxy()
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
		})})))}
	})
	return cache_Data_Ring_ringProxy
}

var cache_Data_Ring_ringFn gopurs_runtime.Value
var once_Data_Ring_ringFn sync.Once

func Get_Data_Ring_ringFn() gopurs_runtime.Value {
	once_Data_Ring_ringFn.Do(func() {
		cache_Data_Ring_ringFn = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ring_ringFn(dictRing_0_box)
		})
	})
	return cache_Data_Ring_ringFn
}

var cache_Data_Ring_negate gopurs_runtime.Value
var once_Data_Ring_negate sync.Once

func Get_Data_Ring_negate() gopurs_runtime.Value {
	once_Data_Ring_negate.Do(func() {
		cache_Data_Ring_negate = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Ring_negate(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](dictRing_0_box))
		})
	})
	return cache_Data_Ring_negate
}

type Constructor_Data_Ring_RingRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00520069006e0067:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00520069006e0067005f00520069006e0067005200650063006f00720064
func init() {
	gopurs_runtime.StructGetters[4246880791] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "SemiringRecord0":
			return gopurs_runtime.Box(c.V0)
		case "subRecord":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Ring_RingRecord: " + key)
		}
	}
}

type Constructor_Data_Ring_Ring[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f00520069006e0067:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f00520069006e0067005f00520069006e0067
func init() {
	gopurs_runtime.StructGetters[3955491866] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ring_Ring[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Semiring0":
			return gopurs_runtime.Box(c.V0)
		case "sub":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_Ring_Ring: " + key)
		}
	}
}

func Call_Data_Ring_RingRecord_dollar_Dict(x_0_loop struct {
	SemiringRecord0 gopurs_runtime.Value
	subRecord       gopurs_runtime.Value
	/*gopurs:record-v1:2:00530065006d006900720069006e0067005200650063006f007200640030,007300750062005200650063006f00720064*/
}) *Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		SemiringRecord0 gopurs_runtime.Value
		subRecord       gopurs_runtime.Value
		/*gopurs:record-v1:2:00530065006d006900720069006e0067005200650063006f007200640030,007300750062005200650063006f00720064*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("SemiringRecord0", "subRecord", orig.SemiringRecord0, orig.subRecord)
	}())
}

func Call_Data_Ring_Ring_dollar_Dict(x_0_loop struct {
	Semiring0 gopurs_runtime.Value
	sub       gopurs_runtime.Value
	/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
}) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
	var x_0 struct {
		Semiring0 gopurs_runtime.Value
		sub       gopurs_runtime.Value
		/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Semiring0", "sub", orig.Semiring0, orig.sub)
	}())
}

func Call_Data_Ring_Ring_dollar_Dict__812180516(x_0_loop struct {
	Semiring0 gopurs_runtime.Value
	sub       gopurs_runtime.Value
	/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
}) *Constructor_Data_Ring_Ring[uint32] {
Ring_dollar_Dict__812180516:
	for {
		if false {
			continue Ring_dollar_Dict__812180516
		}
		var x_0 struct {
			Semiring0 gopurs_runtime.Value
			sub       gopurs_runtime.Value
			/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Semiring0", "sub", orig.Semiring0, orig.sub)
		}())
	}
}

func Call_Data_Ring_Ring_dollar_Dict__34472189(x_0_loop struct {
	Semiring0 gopurs_runtime.Value
	sub       gopurs_runtime.Value
	/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
}) *Constructor_Data_Ring_Ring[int64] {
Ring_dollar_Dict__34472189:
	for {
		if false {
			continue Ring_dollar_Dict__34472189
		}
		var x_0 struct {
			Semiring0 gopurs_runtime.Value
			sub       gopurs_runtime.Value
			/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[int64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Semiring0", "sub", orig.Semiring0, orig.sub)
		}())
	}
}

func Call_Data_Ring_Ring_dollar_Dict__2257635853(x_0_loop struct {
	Semiring0 gopurs_runtime.Value
	sub       gopurs_runtime.Value
	/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
}) *Constructor_Data_Ring_Ring[float64] {
Ring_dollar_Dict__2257635853:
	for {
		if false {
			continue Ring_dollar_Dict__2257635853
		}
		var x_0 struct {
			Semiring0 gopurs_runtime.Value
			sub       gopurs_runtime.Value
			/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[float64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Semiring0", "sub", orig.Semiring0, orig.sub)
		}())
	}
}

func Call_Data_Ring_Ring_dollar_Dict__1193243624(x_0_loop struct {
	Semiring0 gopurs_runtime.Value
	sub       gopurs_runtime.Value
	/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
}) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
Ring_dollar_Dict__1193243624:
	for {
		if false {
			continue Ring_dollar_Dict__1193243624
		}
		var x_0 struct {
			Semiring0 gopurs_runtime.Value
			sub       gopurs_runtime.Value
			/*gopurs:record-v1:2:00530065006d006900720069006e00670030,007300750062*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Semiring0", "sub", orig.Semiring0, orig.sub)
		}())
	}
}

func Call_Data_Ring_subRecord(dict_0_loop *Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Ring_sub(dict_0_loop *Constructor_Data_Ring_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Ring_Ring[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Ring_sub__4157804846(__eta0_0_loop int64, __eta1_1_loop int64) int64 {
sub__4157804846:
	for {
		if false {
			continue sub__4157804846
		}
		var __eta0_0 int64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 int64 = __eta1_1_loop
		_ = __eta1_1
		return gopurs_runtime.IntSub(__eta0_0, __eta1_1)
	}
}

func Call_Data_Ring_sub__179436318(__eta0_0_loop float64, __eta1_1_loop float64) float64 {
sub__179436318:
	for {
		if false {
			continue sub__179436318
		}
		var __eta0_0 float64 = __eta0_0_loop
		_ = __eta0_0
		var __eta1_1 float64 = __eta1_1_loop
		_ = __eta1_1
		return (__eta0_0) - (__eta1_1)
	}
}

func Call_Data_Ring_ringRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, dictRingRecord_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var dictRingRecord_2 gopurs_runtime.Value = dictRingRecord_2_loop
	_ = dictRingRecord_2
	// TAST (Let): semiringRecordCons1_3_0 shape=App(Var) bindingType=Any
	semiringRecordCons1_3_0 := gopurs_runtime.Apply3(Get_Data_Semiring_semiringRecordCons(), dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_2, "SemiringRecord0"), gopurs_runtime.Value{}))
	_ = semiringRecordCons1_3_0
	return gopurs_runtime.Func(func(dictRing_4 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): semiringRecordCons2_5_1 shape=App(Other) bindingType=(TypeApp (ADT ["Data","Semiring","SemiringRecord"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key$scope17), (TypeVar focus$scope18), (TypeVar rowlistTail$scope21)]), (TypeVar row$scope22), (TypeVar subrow$scope20)])
		semiringRecordCons2_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(semiringRecordCons1_3_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_4, "Semiring0"), gopurs_runtime.Value{})))
		_ = semiringRecordCons2_5_1
		return gopurs_runtime.Value{Type: 9, IntVal: 4246880791, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3914418263, UnsafePtr: unsafe.Pointer(semiringRecordCons2_5_1)}
		}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, ra_7 gopurs_runtime.Value, rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): key_9_2 shape=App(Other) bindingType=String
			key_9_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()
			_ = key_9_2
			// TAST (Let): get__3905987523_10_3 shape=App(Var) bindingType=(Func [(Record (Row [] (TypeVar row$scope22)))] (TypeVar focus$scope18))
			get__3905987523_10_3 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_9_2))
			_ = get__3905987523_10_3
			return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_9_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_4, "sub"), gopurs_runtime.Apply(get__3905987523_10_3, ra_7), gopurs_runtime.Apply(get__3905987523_10_3, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictRingRecord_2, "subRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8))
		})}))}
	})
}

func Call_Data_Ring_ringRecord(_dollar___unused_0_loop gopurs_runtime.Value, dictRingRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var dictRingRecord_1 gopurs_runtime.Value = dictRingRecord_1_loop
	_ = dictRingRecord_1
	// TAST (Let): semiringRecord1_2_0 shape=App(Var) bindingType=(ADT ["Data","Semiring","Semiring"] [(Record (Row [] (TypeVar row$scope38)))])
	semiringRecord1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Call_Data_Semiring_semiringRecord(gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_1, "SemiringRecord0"), gopurs_runtime.Value{})))
	_ = semiringRecord1_2_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_Ring[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(semiringRecord1_2_0)}
	}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_1, "subRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_Ring_ringFn(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
	_ = dictRing_0
	// TAST (Let): semiringFn_1_0 shape=App(Var) bindingType=(ADT ["Data","Semiring","Semiring"] [(Func [(TypeVar a$scope53)] (TypeVar b$scope52))])
	semiringFn_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Call_Data_Semiring_semiringFn(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{})))
	_ = semiringFn_1_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_Ring[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(semiringFn_1_0)}
	}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
	})}))}
}

func Call_Data_Ring_negate(dictRing_0_loop *Constructor_Data_Ring_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictRing_0 *Constructor_Data_Ring_Ring[gopurs_runtime.Value] = dictRing_0_loop
	_ = dictRing_0
	// TAST (Let): zero_1_0 shape=Other bindingType=(TypeVar a$scope56)
	zero_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictRing_0.V0, gopurs_runtime.Value{}), "zero")
	_ = zero_1_0
	return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(dictRing_0.V1, zero_1_0, a_2)
	})
}

func Rebox_Data_Ring_2259598678_3060961102(in *Constructor_Data_Ring_Ring[float64]) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ring_Ring[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ring_2370240579_3060961102(in *Constructor_Data_Ring_Ring[uint32]) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ring_Ring[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ring_2534126933_3060961102(in *Constructor_Data_Ring_Ring[int64]) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ring_Ring[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Ring_2826095630_348932501(in *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) *Constructor_Data_Semiring_Semiring[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[int64]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2.IntVal
	out.V3 = in.V3.IntVal
	return out
}

func Rebox_Data_Ring_2826095630_602713622(in *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) *Constructor_Data_Semiring_Semiring[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[float64]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = in.V2.FloatVal()
	out.V3 = in.V3.FloatVal()
	return out
}

func Rebox_Data_Ring_348932501_2826095630(in *Constructor_Data_Semiring_Semiring[int64]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Int(in.V2)
	out.V3 = gopurs_runtime.Int(in.V3)
	return out
}

func Rebox_Data_Ring_602713622_2826095630(in *Constructor_Data_Semiring_Semiring[float64]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = in.V1
	out.V2 = gopurs_runtime.Float(in.V2)
	out.V3 = gopurs_runtime.Float(in.V3)
	return out
}

func Get_Data_Ring_intSub() gopurs_runtime.Value {
	return _Gopurs_Data_Ring_IntSub
}

func Get_Data_Ring_numSub() gopurs_runtime.Value {
	return _Gopurs_Data_Ring_NumSub
}
