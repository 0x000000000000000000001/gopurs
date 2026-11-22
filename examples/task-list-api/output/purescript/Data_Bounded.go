package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Bounded_ordRecord gopurs_runtime.Value
var once_Data_Bounded_ordRecord sync.Once

func Get_Data_Bounded_ordRecord() gopurs_runtime.Value {
	once_Data_Bounded_ordRecord.Do(func() {
		cache_Data_Bounded_ordRecord = gopurs_runtime.Apply(Get_Data_Ord_ordRecord(), gopurs_runtime.Value{})
	})
	return cache_Data_Bounded_ordRecord
}

var cache_Data_Bounded_BoundedRecord_dollar_Dict gopurs_runtime.Value
var once_Data_Bounded_BoundedRecord_dollar_Dict sync.Once

func Get_Data_Bounded_BoundedRecord_dollar_Dict() gopurs_runtime.Value {
	once_Data_Bounded_BoundedRecord_dollar_Dict.Do(func() {
		cache_Data_Bounded_BoundedRecord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4260658871, UnsafePtr: unsafe.Pointer(Call_Data_Bounded_BoundedRecord_dollar_Dict(func() struct {
				OrdRecord0   gopurs_runtime.Value
				bottomRecord gopurs_runtime.Value
				topRecord    gopurs_runtime.Value
				/*gopurs:record-v1:3:004f00720064005200650063006f007200640030,0062006f00740074006f006d005200650063006f00720064,0074006f0070005200650063006f00720064*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					OrdRecord0   gopurs_runtime.Value
					bottomRecord gopurs_runtime.Value
					topRecord    gopurs_runtime.Value
					/*gopurs:record-v1:3:004f00720064005200650063006f007200640030,0062006f00740074006f006d005200650063006f00720064,0074006f0070005200650063006f00720064*/
				}{}
				clone.OrdRecord0 = gopurs_runtime.RecordGet(orig, "OrdRecord0")
				clone.bottomRecord = gopurs_runtime.RecordGet(orig, "bottomRecord")
				clone.topRecord = gopurs_runtime.RecordGet(orig, "topRecord")
				return clone
			}()))}
		})
	})
	return cache_Data_Bounded_BoundedRecord_dollar_Dict
}

var cache_Data_Bounded_Bounded_dollar_Dict gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict sync.Once

func Get_Data_Bounded_Bounded_dollar_Dict() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Call_Data_Bounded_Bounded_dollar_Dict(func() struct {
				Ord0   gopurs_runtime.Value
				bottom gopurs_runtime.Value
				top    gopurs_runtime.Value
				/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0   gopurs_runtime.Value
					bottom gopurs_runtime.Value
					top    gopurs_runtime.Value
					/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.bottom = gopurs_runtime.RecordGet(orig, "bottom")
				clone.top = gopurs_runtime.RecordGet(orig, "top")
				return clone
			}()))}
		})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict
}

var cache_Data_Bounded_Bounded_dollar_Dict__771317046 gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict__771317046 sync.Once

func Get_Data_Bounded_Bounded_dollar_Dict__771317046() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict__771317046.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict__771317046 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3247149001_2094947566(Call_Data_Bounded_Bounded_dollar_Dict__771317046(func() struct {
				Ord0   gopurs_runtime.Value
				bottom *Constructor_Data_DateTime_DateTime
				top    *Constructor_Data_DateTime_DateTime
				/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0   gopurs_runtime.Value
					bottom *Constructor_Data_DateTime_DateTime
					top    *Constructor_Data_DateTime_DateTime
					/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.bottom = gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](gopurs_runtime.RecordGet(orig, "bottom"))
				clone.top = gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](gopurs_runtime.RecordGet(orig, "top"))
				return clone
			}())))}
		})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict__771317046
}

var cache_Data_Bounded_Bounded_dollar_Dict__2637305430 gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict__2637305430 sync.Once

func Get_Data_Bounded_Bounded_dollar_Dict__2637305430() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict__2637305430.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict__2637305430 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_832288803_2094947566(Call_Data_Bounded_Bounded_dollar_Dict__2637305430(func() struct {
				Ord0   gopurs_runtime.Value
				bottom uint32
				top    uint32
				/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0   gopurs_runtime.Value
					bottom uint32
					top    uint32
					/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.bottom = uint32(gopurs_runtime.RecordGet(orig, "bottom").IntVal)
				clone.top = uint32(gopurs_runtime.RecordGet(orig, "top").IntVal)
				return clone
			}())))}
		})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict__2637305430
}

var cache_Data_Bounded_Bounded_dollar_Dict__934371414 gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict__934371414 sync.Once

func Get_Data_Bounded_Bounded_dollar_Dict__934371414() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict__934371414.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict__934371414 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_832288803_2094947566(Call_Data_Bounded_Bounded_dollar_Dict__934371414(func() struct {
				Ord0   gopurs_runtime.Value
				bottom uint32
				top    uint32
				/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0   gopurs_runtime.Value
					bottom uint32
					top    uint32
					/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.bottom = uint32(gopurs_runtime.RecordGet(orig, "bottom").IntVal)
				clone.top = uint32(gopurs_runtime.RecordGet(orig, "top").IntVal)
				return clone
			}())))}
		})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict__934371414
}

var cache_Data_Bounded_Bounded_dollar_Dict__4087992182 gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict__4087992182 sync.Once

func Get_Data_Bounded_Bounded_dollar_Dict__4087992182() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict__4087992182.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict__4087992182 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3523628265_2094947566(Call_Data_Bounded_Bounded_dollar_Dict__4087992182(func() struct {
				Ord0   gopurs_runtime.Value
				bottom *Constructor_Data_Date_Date
				top    *Constructor_Data_Date_Date
				/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0   gopurs_runtime.Value
					bottom *Constructor_Data_Date_Date
					top    *Constructor_Data_Date_Date
					/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.bottom = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.RecordGet(orig, "bottom"))
				clone.top = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.RecordGet(orig, "top"))
				return clone
			}())))}
		})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict__4087992182
}

var cache_Data_Bounded_Bounded_dollar_Dict__822771094 gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict__822771094 sync.Once

func Get_Data_Bounded_Bounded_dollar_Dict__822771094() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict__822771094.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict__822771094 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_832288803_2094947566(Call_Data_Bounded_Bounded_dollar_Dict__822771094(func() struct {
				Ord0   gopurs_runtime.Value
				bottom uint32
				top    uint32
				/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0   gopurs_runtime.Value
					bottom uint32
					top    uint32
					/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.bottom = uint32(gopurs_runtime.RecordGet(orig, "bottom").IntVal)
				clone.top = uint32(gopurs_runtime.RecordGet(orig, "top").IntVal)
				return clone
			}())))}
		})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict__822771094
}

var cache_Data_Bounded_Bounded_dollar_Dict__1591666678 gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict__1591666678 sync.Once

func Get_Data_Bounded_Bounded_dollar_Dict__1591666678() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict__1591666678.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict__1591666678 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_832288803_2094947566(Call_Data_Bounded_Bounded_dollar_Dict__1591666678(func() struct {
				Ord0   gopurs_runtime.Value
				bottom uint32
				top    uint32
				/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0   gopurs_runtime.Value
					bottom uint32
					top    uint32
					/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.bottom = uint32(gopurs_runtime.RecordGet(orig, "bottom").IntVal)
				clone.top = uint32(gopurs_runtime.RecordGet(orig, "top").IntVal)
				return clone
			}())))}
		})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict__1591666678
}

var cache_Data_Bounded_Bounded_dollar_Dict__419722870 gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict__419722870 sync.Once

func Get_Data_Bounded_Bounded_dollar_Dict__419722870() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict__419722870.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict__419722870 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_4136451977_2094947566(Call_Data_Bounded_Bounded_dollar_Dict__419722870(func() struct {
				Ord0   gopurs_runtime.Value
				bottom *Constructor_Data_Time_Time
				top    *Constructor_Data_Time_Time
				/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0   gopurs_runtime.Value
					bottom *Constructor_Data_Time_Time
					top    *Constructor_Data_Time_Time
					/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.bottom = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](gopurs_runtime.RecordGet(orig, "bottom"))
				clone.top = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](gopurs_runtime.RecordGet(orig, "top"))
				return clone
			}())))}
		})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict__419722870
}

var cache_Data_Bounded_Bounded_dollar_Dict__718008214 gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict__718008214 sync.Once

func Get_Data_Bounded_Bounded_dollar_Dict__718008214() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict__718008214.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict__718008214 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3728870730_2094947566(Call_Data_Bounded_Bounded_dollar_Dict__718008214(func() struct {
				Ord0   gopurs_runtime.Value
				bottom bool
				top    bool
				/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0   gopurs_runtime.Value
					bottom bool
					top    bool
					/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.bottom = (gopurs_runtime.RecordGet(orig, "bottom").IntVal) != (0)
				clone.top = (gopurs_runtime.RecordGet(orig, "top").IntVal) != (0)
				return clone
			}())))}
		})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict__718008214
}

var cache_Data_Bounded_Bounded_dollar_Dict__3297762038 gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict__3297762038 sync.Once

func Get_Data_Bounded_Bounded_dollar_Dict__3297762038() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict__3297762038.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict__3297762038 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_2420955921_2094947566(Call_Data_Bounded_Bounded_dollar_Dict__3297762038(func() struct {
				Ord0   gopurs_runtime.Value
				bottom string
				top    string
				/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0   gopurs_runtime.Value
					bottom string
					top    string
					/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.bottom = gopurs_runtime.RecordGet(orig, "bottom").StrVal()
				clone.top = gopurs_runtime.RecordGet(orig, "top").StrVal()
				return clone
			}())))}
		})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict__3297762038
}

var cache_Data_Bounded_Bounded_dollar_Dict__3722462422 gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict__3722462422 sync.Once

func Get_Data_Bounded_Bounded_dollar_Dict__3722462422() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict__3722462422.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict__3722462422 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3764732725_2094947566(Call_Data_Bounded_Bounded_dollar_Dict__3722462422(func() struct {
				Ord0   gopurs_runtime.Value
				bottom int64
				top    int64
				/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0   gopurs_runtime.Value
					bottom int64
					top    int64
					/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.bottom = gopurs_runtime.RecordGet(orig, "bottom").IntVal
				clone.top = gopurs_runtime.RecordGet(orig, "top").IntVal
				return clone
			}())))}
		})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict__3722462422
}

var cache_Data_Bounded_Bounded_dollar_Dict__3077815030 gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict__3077815030 sync.Once

func Get_Data_Bounded_Bounded_dollar_Dict__3077815030() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict__3077815030.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict__3077815030 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_4294747894_2094947566(Call_Data_Bounded_Bounded_dollar_Dict__3077815030(func() struct {
				Ord0   gopurs_runtime.Value
				bottom float64
				top    float64
				/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0   gopurs_runtime.Value
					bottom float64
					top    float64
					/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.bottom = gopurs_runtime.RecordGet(orig, "bottom").FloatVal()
				clone.top = gopurs_runtime.RecordGet(orig, "top").FloatVal()
				return clone
			}())))}
		})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict__3077815030
}

var cache_Data_Bounded_Bounded_dollar_Dict__279484406 gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict__279484406 sync.Once

func Get_Data_Bounded_Bounded_dollar_Dict__279484406() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict__279484406.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict__279484406 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Call_Data_Bounded_Bounded_dollar_Dict__279484406(func() struct {
				Ord0   gopurs_runtime.Value
				bottom gopurs_runtime.Value
				top    gopurs_runtime.Value
				/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ord0   gopurs_runtime.Value
					bottom gopurs_runtime.Value
					top    gopurs_runtime.Value
					/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
				}{}
				clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
				clone.bottom = gopurs_runtime.RecordGet(orig, "bottom")
				clone.top = gopurs_runtime.RecordGet(orig, "top")
				return clone
			}()))}
		})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict__279484406
}

var cache_Data_Bounded_topRecord gopurs_runtime.Value
var once_Data_Bounded_topRecord sync.Once

func Get_Data_Bounded_topRecord() gopurs_runtime.Value {
	once_Data_Bounded_topRecord.Do(func() {
		cache_Data_Bounded_topRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bounded_topRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Bounded_topRecord
}

var cache_Data_Bounded_top gopurs_runtime.Value
var once_Data_Bounded_top sync.Once

func Get_Data_Bounded_top() gopurs_runtime.Value {
	once_Data_Bounded_top.Do(func() {
		cache_Data_Bounded_top = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bounded_top(dict_0_box)
		})
	})
	return cache_Data_Bounded_top
}

var cache_Data_Bounded_boundedChar gopurs_runtime.Value
var once_Data_Bounded_boundedChar sync.Once

func Get_Data_Bounded_boundedChar() gopurs_runtime.Value {
	once_Data_Bounded_boundedChar.Do(func() {
		cache_Data_Bounded_boundedChar = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_2420955921_2094947566((&Constructor_Data_Bounded_Bounded[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_2406510097_4177771502(Rebox_Data_Bounded_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordChar()))))}
		}), Get_Data_Bounded_bottomChar().StrVal(), Get_Data_Bounded_topChar().StrVal()})))}
	})
	return cache_Data_Bounded_boundedChar
}

var cache_Data_Bounded_top__2335111300 gopurs_runtime.Value
var once_Data_Bounded_top__2335111300 sync.Once

func Get_Data_Bounded_top__2335111300() gopurs_runtime.Value {
	once_Data_Bounded_top__2335111300.Do(func() {
		cache_Data_Bounded_top__2335111300 = gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar()))
	})
	return cache_Data_Bounded_top__2335111300
}

var cache_Data_Bounded_boundedInt gopurs_runtime.Value
var once_Data_Bounded_boundedInt sync.Once

func Get_Data_Bounded_boundedInt() gopurs_runtime.Value {
	once_Data_Bounded_boundedInt.Do(func() {
		cache_Data_Bounded_boundedInt = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3764732725_2094947566((&Constructor_Data_Bounded_Bounded[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3308271157_4177771502(Rebox_Data_Bounded_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}
		}), Get_Data_Bounded_bottomInt().IntVal, Get_Data_Bounded_topInt().IntVal})))}
	})
	return cache_Data_Bounded_boundedInt
}

var cache_Data_Bounded_top__2052445839 gopurs_runtime.Value
var once_Data_Bounded_top__2052445839 sync.Once

func Get_Data_Bounded_top__2052445839() gopurs_runtime.Value {
	once_Data_Bounded_top__2052445839.Do(func() {
		cache_Data_Bounded_top__2052445839 = gopurs_runtime.Int(gopurs_runtime.Int(Get_Data_Bounded_topInt().IntVal).IntVal)
	})
	return cache_Data_Bounded_top__2052445839
}

var cache_Data_Bounded_boundedUnit gopurs_runtime.Value
var once_Data_Bounded_boundedUnit sync.Once

func Get_Data_Bounded_boundedUnit() gopurs_runtime.Value {
	once_Data_Bounded_boundedUnit.Do(func() {
		cache_Data_Bounded_boundedUnit = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordUnit()))}
		}), Get_Data_Unit_unit(), Get_Data_Unit_unit()}))}
	})
	return cache_Data_Bounded_boundedUnit
}

var cache_Data_Bounded_boundedRecordNil gopurs_runtime.Value
var once_Data_Bounded_boundedRecordNil sync.Once

func Get_Data_Bounded_boundedRecordNil() gopurs_runtime.Value {
	once_Data_Bounded_boundedRecordNil.Do(func() {
		cache_Data_Bounded_boundedRecordNil = gopurs_runtime.Value{Type: 9, IntVal: 4260658871, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Ord_ordRecordNil()
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict0()
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.RecordDict0()
		})}))}
	})
	return cache_Data_Bounded_boundedRecordNil
}

var cache_Data_Bounded_boundedProxy gopurs_runtime.Value
var once_Data_Bounded_boundedProxy sync.Once

func Get_Data_Bounded_boundedProxy() gopurs_runtime.Value {
	once_Data_Bounded_boundedProxy.Do(func() {
		cache_Data_Bounded_boundedProxy = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_832288803_2094947566((&Constructor_Data_Bounded_Bounded[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Data_Ord_ordProxy()
		}), uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}.IntVal), uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}.IntVal)})))}
	})
	return cache_Data_Bounded_boundedProxy
}

var cache_Data_Bounded_boundedOrdering gopurs_runtime.Value
var once_Data_Bounded_boundedOrdering sync.Once

func Get_Data_Bounded_boundedOrdering() gopurs_runtime.Value {
	once_Data_Bounded_boundedOrdering.Do(func() {
		cache_Data_Bounded_boundedOrdering = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_832288803_2094947566((&Constructor_Data_Bounded_Bounded[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3730953251_4177771502(Rebox_Data_Bounded_4177771502_3730953251(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordOrdering()))))}
		}), 1527465420, 380165415})))}
	})
	return cache_Data_Bounded_boundedOrdering
}

var cache_Data_Bounded_boundedNumber gopurs_runtime.Value
var once_Data_Bounded_boundedNumber sync.Once

func Get_Data_Bounded_boundedNumber() gopurs_runtime.Value {
	once_Data_Bounded_boundedNumber.Do(func() {
		cache_Data_Bounded_boundedNumber = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_4294747894_2094947566((&Constructor_Data_Bounded_Bounded[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3047586294_4177771502(Rebox_Data_Bounded_4177771502_3047586294(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordNumber()))))}
		}), Get_Data_Bounded_bottomNumber().FloatVal(), Get_Data_Bounded_topNumber().FloatVal()})))}
	})
	return cache_Data_Bounded_boundedNumber
}

var cache_Data_Bounded_boundedBoolean gopurs_runtime.Value
var once_Data_Bounded_boundedBoolean sync.Once

func Get_Data_Bounded_boundedBoolean() gopurs_runtime.Value {
	once_Data_Bounded_boundedBoolean.Do(func() {
		cache_Data_Bounded_boundedBoolean = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3728870730_2094947566((&Constructor_Data_Bounded_Bounded[bool]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_219188042_4177771502(Rebox_Data_Bounded_4177771502_219188042(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordBoolean()))))}
		}), false, true})))}
	})
	return cache_Data_Bounded_boundedBoolean
}

var cache_Data_Bounded_bottomRecord gopurs_runtime.Value
var once_Data_Bounded_bottomRecord sync.Once

func Get_Data_Bounded_bottomRecord() gopurs_runtime.Value {
	once_Data_Bounded_bottomRecord.Do(func() {
		cache_Data_Bounded_bottomRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bounded_bottomRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_Bounded_bottomRecord
}

var cache_Data_Bounded_boundedRecord gopurs_runtime.Value
var once_Data_Bounded_boundedRecord sync.Once

func Get_Data_Bounded_boundedRecord() gopurs_runtime.Value {
	once_Data_Bounded_boundedRecord.Do(func() {
		cache_Data_Bounded_boundedRecord = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, dictBoundedRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bounded_boundedRecord(_dollar___unused_0_box, dictBoundedRecord_1_box)
		})
	})
	return cache_Data_Bounded_boundedRecord
}

var cache_Data_Bounded_bottom gopurs_runtime.Value
var once_Data_Bounded_bottom sync.Once

func Get_Data_Bounded_bottom() gopurs_runtime.Value {
	once_Data_Bounded_bottom.Do(func() {
		cache_Data_Bounded_bottom = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bounded_bottom(dict_0_box)
		})
	})
	return cache_Data_Bounded_bottom
}

var cache_Data_Bounded_bottom__2335111300 gopurs_runtime.Value
var once_Data_Bounded_bottom__2335111300 sync.Once

func Get_Data_Bounded_bottom__2335111300() gopurs_runtime.Value {
	once_Data_Bounded_bottom__2335111300.Do(func() {
		cache_Data_Bounded_bottom__2335111300 = gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar()))
	})
	return cache_Data_Bounded_bottom__2335111300
}

var cache_Data_Bounded_boundedRecordCons gopurs_runtime.Value
var once_Data_Bounded_boundedRecordCons sync.Once

func Get_Data_Bounded_boundedRecordCons() gopurs_runtime.Value {
	once_Data_Bounded_boundedRecordCons.Do(func() {
		cache_Data_Bounded_boundedRecordCons = gopurs_runtime.Func2(func(dictIsSymbol_0_box gopurs_runtime.Value, dictBounded_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Bounded_boundedRecordCons(dictIsSymbol_0_box, dictBounded_1_box)
		})
	})
	return cache_Data_Bounded_boundedRecordCons
}

type Constructor_Data_Bounded_BoundedRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f0042006f0075006e006400650064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f0042006f0075006e006400650064005f0042006f0075006e006400650064005200650063006f00720064
func init() {
	gopurs_runtime.StructGetters[4260658871] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "OrdRecord0":
			return gopurs_runtime.Box(c.V0)
		case "bottomRecord":
			return gopurs_runtime.Box(c.V1)
		case "topRecord":
			return gopurs_runtime.Box(c.V2)
		default:
			panic("Key not found in dictionary Constructor_Data_Bounded_BoundedRecord: " + key)
		}
	}
}

type Constructor_Data_Bounded_Bounded[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_a
	V2 T_a
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f0042006f0075006e006400650064:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f0042006f0075006e006400650064005f0042006f0075006e006400650064
func init() {
	gopurs_runtime.StructGetters[3510799738] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Bounded_Bounded[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Ord0":
			return gopurs_runtime.Box(c.V0)
		case "bottom":
			return gopurs_runtime.Box(c.V1)
		case "top":
			return gopurs_runtime.Box(c.V2)
		default:
			panic("Key not found in dictionary Constructor_Data_Bounded_Bounded: " + key)
		}
	}
}

func Call_Data_Bounded_BoundedRecord_dollar_Dict(x_0_loop struct {
	OrdRecord0   gopurs_runtime.Value
	bottomRecord gopurs_runtime.Value
	topRecord    gopurs_runtime.Value
	/*gopurs:record-v1:3:004f00720064005200650063006f007200640030,0062006f00740074006f006d005200650063006f00720064,0074006f0070005200650063006f00720064*/
}) *Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		OrdRecord0   gopurs_runtime.Value
		bottomRecord gopurs_runtime.Value
		topRecord    gopurs_runtime.Value
		/*gopurs:record-v1:3:004f00720064005200650063006f007200640030,0062006f00740074006f006d005200650063006f00720064,0074006f0070005200650063006f00720064*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict3("OrdRecord0", "bottomRecord", "topRecord", orig.OrdRecord0, orig.bottomRecord, orig.topRecord)
	}())
}

func Call_Data_Bounded_Bounded_dollar_Dict(x_0_loop struct {
	Ord0   gopurs_runtime.Value
	bottom gopurs_runtime.Value
	top    gopurs_runtime.Value
	/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
}) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	var x_0 struct {
		Ord0   gopurs_runtime.Value
		bottom gopurs_runtime.Value
		top    gopurs_runtime.Value
		/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", orig.Ord0, orig.bottom, orig.top)
	}())
}

func Call_Data_Bounded_Bounded_dollar_Dict__771317046(x_0_loop struct {
	Ord0   gopurs_runtime.Value
	bottom *Constructor_Data_DateTime_DateTime
	top    *Constructor_Data_DateTime_DateTime
	/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
}) *Constructor_Data_Bounded_Bounded[*Constructor_Data_DateTime_DateTime] {
Bounded_dollar_Dict__771317046:
	for {
		if false {
			continue Bounded_dollar_Dict__771317046
		}
		var x_0 struct {
			Ord0   gopurs_runtime.Value
			bottom *Constructor_Data_DateTime_DateTime
			top    *Constructor_Data_DateTime_DateTime
			/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[*Constructor_Data_DateTime_DateTime]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", orig.Ord0, gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(orig.bottom)}, gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(orig.top)})
		}())
	}
}

func Call_Data_Bounded_Bounded_dollar_Dict__2637305430(x_0_loop struct {
	Ord0   gopurs_runtime.Value
	bottom uint32
	top    uint32
	/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
}) *Constructor_Data_Bounded_Bounded[uint32] {
Bounded_dollar_Dict__2637305430:
	for {
		if false {
			continue Bounded_dollar_Dict__2637305430
		}
		var x_0 struct {
			Ord0   gopurs_runtime.Value
			bottom uint32
			top    uint32
			/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", orig.Ord0, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.bottom), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.top), UnsafePtr: nil})
		}())
	}
}

func Call_Data_Bounded_Bounded_dollar_Dict__934371414(x_0_loop struct {
	Ord0   gopurs_runtime.Value
	bottom uint32
	top    uint32
	/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
}) *Constructor_Data_Bounded_Bounded[uint32] {
Bounded_dollar_Dict__934371414:
	for {
		if false {
			continue Bounded_dollar_Dict__934371414
		}
		var x_0 struct {
			Ord0   gopurs_runtime.Value
			bottom uint32
			top    uint32
			/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", orig.Ord0, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.bottom), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.top), UnsafePtr: nil})
		}())
	}
}

func Call_Data_Bounded_Bounded_dollar_Dict__4087992182(x_0_loop struct {
	Ord0   gopurs_runtime.Value
	bottom *Constructor_Data_Date_Date
	top    *Constructor_Data_Date_Date
	/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
}) *Constructor_Data_Bounded_Bounded[*Constructor_Data_Date_Date] {
Bounded_dollar_Dict__4087992182:
	for {
		if false {
			continue Bounded_dollar_Dict__4087992182
		}
		var x_0 struct {
			Ord0   gopurs_runtime.Value
			bottom *Constructor_Data_Date_Date
			top    *Constructor_Data_Date_Date
			/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[*Constructor_Data_Date_Date]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", orig.Ord0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(orig.bottom)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(orig.top)})
		}())
	}
}

func Call_Data_Bounded_Bounded_dollar_Dict__822771094(x_0_loop struct {
	Ord0   gopurs_runtime.Value
	bottom uint32
	top    uint32
	/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
}) *Constructor_Data_Bounded_Bounded[uint32] {
Bounded_dollar_Dict__822771094:
	for {
		if false {
			continue Bounded_dollar_Dict__822771094
		}
		var x_0 struct {
			Ord0   gopurs_runtime.Value
			bottom uint32
			top    uint32
			/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", orig.Ord0, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.bottom), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.top), UnsafePtr: nil})
		}())
	}
}

func Call_Data_Bounded_Bounded_dollar_Dict__1591666678(x_0_loop struct {
	Ord0   gopurs_runtime.Value
	bottom uint32
	top    uint32
	/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
}) *Constructor_Data_Bounded_Bounded[uint32] {
Bounded_dollar_Dict__1591666678:
	for {
		if false {
			continue Bounded_dollar_Dict__1591666678
		}
		var x_0 struct {
			Ord0   gopurs_runtime.Value
			bottom uint32
			top    uint32
			/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", orig.Ord0, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.bottom), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(orig.top), UnsafePtr: nil})
		}())
	}
}

func Call_Data_Bounded_Bounded_dollar_Dict__419722870(x_0_loop struct {
	Ord0   gopurs_runtime.Value
	bottom *Constructor_Data_Time_Time
	top    *Constructor_Data_Time_Time
	/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
}) *Constructor_Data_Bounded_Bounded[*Constructor_Data_Time_Time] {
Bounded_dollar_Dict__419722870:
	for {
		if false {
			continue Bounded_dollar_Dict__419722870
		}
		var x_0 struct {
			Ord0   gopurs_runtime.Value
			bottom *Constructor_Data_Time_Time
			top    *Constructor_Data_Time_Time
			/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[*Constructor_Data_Time_Time]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", orig.Ord0, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(orig.bottom)}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(orig.top)})
		}())
	}
}

func Call_Data_Bounded_Bounded_dollar_Dict__718008214(x_0_loop struct {
	Ord0   gopurs_runtime.Value
	bottom bool
	top    bool
	/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
}) *Constructor_Data_Bounded_Bounded[bool] {
Bounded_dollar_Dict__718008214:
	for {
		if false {
			continue Bounded_dollar_Dict__718008214
		}
		var x_0 struct {
			Ord0   gopurs_runtime.Value
			bottom bool
			top    bool
			/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[bool]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", orig.Ord0, gopurs_runtime.Bool(orig.bottom), gopurs_runtime.Bool(orig.top))
		}())
	}
}

func Call_Data_Bounded_Bounded_dollar_Dict__3297762038(x_0_loop struct {
	Ord0   gopurs_runtime.Value
	bottom string
	top    string
	/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
}) *Constructor_Data_Bounded_Bounded[string] {
Bounded_dollar_Dict__3297762038:
	for {
		if false {
			continue Bounded_dollar_Dict__3297762038
		}
		var x_0 struct {
			Ord0   gopurs_runtime.Value
			bottom string
			top    string
			/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[string]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", orig.Ord0, gopurs_runtime.Str(orig.bottom), gopurs_runtime.Str(orig.top))
		}())
	}
}

func Call_Data_Bounded_Bounded_dollar_Dict__3722462422(x_0_loop struct {
	Ord0   gopurs_runtime.Value
	bottom int64
	top    int64
	/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
}) *Constructor_Data_Bounded_Bounded[int64] {
Bounded_dollar_Dict__3722462422:
	for {
		if false {
			continue Bounded_dollar_Dict__3722462422
		}
		var x_0 struct {
			Ord0   gopurs_runtime.Value
			bottom int64
			top    int64
			/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[int64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", orig.Ord0, gopurs_runtime.Int(orig.bottom), gopurs_runtime.Int(orig.top))
		}())
	}
}

func Call_Data_Bounded_Bounded_dollar_Dict__3077815030(x_0_loop struct {
	Ord0   gopurs_runtime.Value
	bottom float64
	top    float64
	/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
}) *Constructor_Data_Bounded_Bounded[float64] {
Bounded_dollar_Dict__3077815030:
	for {
		if false {
			continue Bounded_dollar_Dict__3077815030
		}
		var x_0 struct {
			Ord0   gopurs_runtime.Value
			bottom float64
			top    float64
			/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[float64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", orig.Ord0, gopurs_runtime.Float(orig.bottom), gopurs_runtime.Float(orig.top))
		}())
	}
}

func Call_Data_Bounded_Bounded_dollar_Dict__279484406(x_0_loop struct {
	Ord0   gopurs_runtime.Value
	bottom gopurs_runtime.Value
	top    gopurs_runtime.Value
	/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
}) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
Bounded_dollar_Dict__279484406:
	for {
		if false {
			continue Bounded_dollar_Dict__279484406
		}
		var x_0 struct {
			Ord0   gopurs_runtime.Value
			bottom gopurs_runtime.Value
			top    gopurs_runtime.Value
			/*gopurs:record-v1:3:004f007200640030,0062006f00740074006f006d,0074006f0070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", orig.Ord0, orig.bottom, orig.top)
		}())
	}
}

func Call_Data_Bounded_topRecord(dict_0_loop *Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V2
}

func Call_Data_Bounded_top(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "top")
}

func Call_Data_Bounded_bottomRecord(dict_0_loop *Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_Bounded_boundedRecord(_dollar___unused_0_loop gopurs_runtime.Value, dictBoundedRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var dictBoundedRecord_1 gopurs_runtime.Value = dictBoundedRecord_1_loop
	_ = dictBoundedRecord_1
	// TAST (Let): ordRecord1_2_0 shape=App(Var) bindingType=(ADT ["Data","Ord","Ord"] [(Record (Row [] (TypeVar row$scope28)))])
	ordRecord1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Ord_ordRecord(gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedRecord_1, "OrdRecord0"), gopurs_runtime.Value{})))
	_ = ordRecord1_2_0
	return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordRecord1_2_0)}
	}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_1, "bottomRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_1, "topRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_Bounded_bottom(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "bottom")
}

func Call_Data_Bounded_boundedRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictBounded_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var dictBounded_1 gopurs_runtime.Value = dictBounded_1_loop
	_ = dictBounded_1
	// TAST (Let): top1_2_0 shape=App(Var) bindingType=(TypeVar focus$scope48)
	top1_2_0 := Call_Data_Bounded_top(dictBounded_1)
	_ = top1_2_0
	// TAST (Let): bottom1_3_1 shape=App(Var) bindingType=(TypeVar focus$scope48)
	bottom1_3_1 := Call_Data_Bounded_bottom(dictBounded_1)
	_ = bottom1_3_1
	// TAST (Let): Ord0_4_2 shape=App(Other) bindingType=Any
	Ord0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_1, "Ord0"), gopurs_runtime.Value{})
	_ = Ord0_4_2
	return gopurs_runtime.Func3(func(_dollar___unused_5 gopurs_runtime.Value, _dollar___unused_6 gopurs_runtime.Value, dictBoundedRecord_7 gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): ordRecordCons_8_3 shape=App(Var) bindingType=(TypeApp (ADT ["Data","Ord","OrdRecord"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key$scope47), (TypeVar focus$scope48), (TypeVar rowlistTail$scope53)]), (TypeVar row$scope50)])
		ordRecordCons_8_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Data_Ord_ordRecordCons(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedRecord_7, "OrdRecord0"), gopurs_runtime.Value{})), gopurs_runtime.Value{}, dictIsSymbol_0, Ord0_4_2))
		_ = ordRecordCons_8_3
		return gopurs_runtime.Value{Type: 9, IntVal: 4260658871, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4162894775, UnsafePtr: unsafe.Pointer(ordRecordCons_8_3)}
		}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, rowProxy_10 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})), bottom1_3_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_7, "bottomRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(rowProxy_10.IntVal)), UnsafePtr: nil}))
		}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, rowProxy_10 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})), top1_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_7, "topRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(rowProxy_10.IntVal)), UnsafePtr: nil}))
		})}))}
	})
}

func Rebox_Data_Bounded_219188042_4177771502(in *Constructor_Data_Ord_Ord[bool]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Bounded_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Bounded_2420955921_2094947566(in *Constructor_Data_Bounded_Bounded[string]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Str(in.V1)
	out.V2 = gopurs_runtime.Str(in.V2)
	return out
}

func Rebox_Data_Bounded_3047586294_4177771502(in *Constructor_Data_Ord_Ord[float64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Bounded_3247149001_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_DateTime_DateTime]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(in.V1)}
	out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_Bounded_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Bounded_3523628265_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_Date_Date]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V1)}
	out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_Bounded_3728870730_2094947566(in *Constructor_Data_Bounded_Bounded[bool]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Bool(in.V1)
	out.V2 = gopurs_runtime.Bool(in.V2)
	return out
}

func Rebox_Data_Bounded_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Bounded_3764732725_2094947566(in *Constructor_Data_Bounded_Bounded[int64]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Int(in.V1)
	out.V2 = gopurs_runtime.Int(in.V2)
	return out
}

func Rebox_Data_Bounded_4136451977_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_Time_Time]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(in.V1)}
	out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_Bounded_4177771502_219188042(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[bool] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[bool])(unsafe.Pointer(in))
}

func Rebox_Data_Bounded_4177771502_2406510097(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[string])(unsafe.Pointer(in))
}

func Rebox_Data_Bounded_4177771502_3047586294(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[float64])(unsafe.Pointer(in))
}

func Rebox_Data_Bounded_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[int64])(unsafe.Pointer(in))
}

func Rebox_Data_Bounded_4177771502_3730953251(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Bounded_4294747894_2094947566(in *Constructor_Data_Bounded_Bounded[float64]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Float(in.V1)
	out.V2 = gopurs_runtime.Float(in.V2)
	return out
}

func Rebox_Data_Bounded_832288803_2094947566(in *Constructor_Data_Bounded_Bounded[uint32]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
	out.V0 = in.V0
	out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
	out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
	return out
}

func Get_Data_Bounded_bottomChar() gopurs_runtime.Value {
	return _Gopurs_Data_Bounded_BottomChar
}

func Get_Data_Bounded_bottomInt() gopurs_runtime.Value {
	return _Gopurs_Data_Bounded_BottomInt
}

func Get_Data_Bounded_bottomNumber() gopurs_runtime.Value {
	return _Gopurs_Data_Bounded_BottomNumber
}

func Get_Data_Bounded_topChar() gopurs_runtime.Value {
	return _Gopurs_Data_Bounded_TopChar
}

func Get_Data_Bounded_topInt() gopurs_runtime.Value {
	return _Gopurs_Data_Bounded_TopInt
}

func Get_Data_Bounded_topNumber() gopurs_runtime.Value {
	return _Gopurs_Data_Bounded_TopNumber
}
