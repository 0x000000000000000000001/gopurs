package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_DivisionRing_div gopurs_runtime.Value
var once_Data_DivisionRing_div sync.Once

func Get_Data_DivisionRing_div() gopurs_runtime.Value {
	once_Data_DivisionRing_div.Do(func() {
		cache_Data_DivisionRing_div = Get_Data_EuclideanRing_div__266314078()
	})
	return cache_Data_DivisionRing_div
}

var cache_Data_DivisionRing_DivisionRing_dollar_Dict gopurs_runtime.Value
var once_Data_DivisionRing_DivisionRing_dollar_Dict sync.Once

func Get_Data_DivisionRing_DivisionRing_dollar_Dict() gopurs_runtime.Value {
	once_Data_DivisionRing_DivisionRing_dollar_Dict.Do(func() {
		cache_Data_DivisionRing_DivisionRing_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2548491258, UnsafePtr: unsafe.Pointer(Call_Data_DivisionRing_DivisionRing_dollar_Dict(func() struct {
				Ring0 gopurs_runtime.Value
				recip gopurs_runtime.Value
				/*gopurs:record-v1:2:00520069006e00670030,00720065006300690070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ring0 gopurs_runtime.Value
					recip gopurs_runtime.Value
					/*gopurs:record-v1:2:00520069006e00670030,00720065006300690070*/
				}{}
				clone.Ring0 = gopurs_runtime.RecordGet(orig, "Ring0")
				clone.recip = gopurs_runtime.RecordGet(orig, "recip")
				return clone
			}()))}
		})
	})
	return cache_Data_DivisionRing_DivisionRing_dollar_Dict
}

var cache_Data_DivisionRing_DivisionRing_dollar_Dict__2494093930 gopurs_runtime.Value
var once_Data_DivisionRing_DivisionRing_dollar_Dict__2494093930 sync.Once

func Get_Data_DivisionRing_DivisionRing_dollar_Dict__2494093930() gopurs_runtime.Value {
	once_Data_DivisionRing_DivisionRing_dollar_Dict__2494093930.Do(func() {
		cache_Data_DivisionRing_DivisionRing_dollar_Dict__2494093930 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2548491258, UnsafePtr: unsafe.Pointer(Rebox_Data_DivisionRing_2384868323_2381053422(Call_Data_DivisionRing_DivisionRing_dollar_Dict__2494093930(func() struct {
				Ring0 gopurs_runtime.Value
				recip gopurs_runtime.Value
				/*gopurs:record-v1:2:00520069006e00670030,00720065006300690070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ring0 gopurs_runtime.Value
					recip gopurs_runtime.Value
					/*gopurs:record-v1:2:00520069006e00670030,00720065006300690070*/
				}{}
				clone.Ring0 = gopurs_runtime.RecordGet(orig, "Ring0")
				clone.recip = gopurs_runtime.RecordGet(orig, "recip")
				return clone
			}())))}
		})
	})
	return cache_Data_DivisionRing_DivisionRing_dollar_Dict__2494093930
}

var cache_Data_DivisionRing_DivisionRing_dollar_Dict__1513685994 gopurs_runtime.Value
var once_Data_DivisionRing_DivisionRing_dollar_Dict__1513685994 sync.Once

func Get_Data_DivisionRing_DivisionRing_dollar_Dict__1513685994() gopurs_runtime.Value {
	once_Data_DivisionRing_DivisionRing_dollar_Dict__1513685994.Do(func() {
		cache_Data_DivisionRing_DivisionRing_dollar_Dict__1513685994 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2548491258, UnsafePtr: unsafe.Pointer(Rebox_Data_DivisionRing_1264867830_2381053422(Call_Data_DivisionRing_DivisionRing_dollar_Dict__1513685994(func() struct {
				Ring0 gopurs_runtime.Value
				recip gopurs_runtime.Value
				/*gopurs:record-v1:2:00520069006e00670030,00720065006300690070*/
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Ring0 gopurs_runtime.Value
					recip gopurs_runtime.Value
					/*gopurs:record-v1:2:00520069006e00670030,00720065006300690070*/
				}{}
				clone.Ring0 = gopurs_runtime.RecordGet(orig, "Ring0")
				clone.recip = gopurs_runtime.RecordGet(orig, "recip")
				return clone
			}())))}
		})
	})
	return cache_Data_DivisionRing_DivisionRing_dollar_Dict__1513685994
}

var cache_Data_DivisionRing_recip gopurs_runtime.Value
var once_Data_DivisionRing_recip sync.Once

func Get_Data_DivisionRing_recip() gopurs_runtime.Value {
	once_Data_DivisionRing_recip.Do(func() {
		cache_Data_DivisionRing_recip = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_DivisionRing_recip(gopurs_runtime.CoerceToStruct[Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Data_DivisionRing_recip
}

var cache_Data_DivisionRing_rightDiv gopurs_runtime.Value
var once_Data_DivisionRing_rightDiv sync.Once

func Get_Data_DivisionRing_rightDiv() gopurs_runtime.Value {
	once_Data_DivisionRing_rightDiv.Do(func() {
		cache_Data_DivisionRing_rightDiv = gopurs_runtime.Func3(func(dictDivisionRing_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_DivisionRing_rightDiv(gopurs_runtime.CoerceToStruct[Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value]](dictDivisionRing_0_box), a_1_box, b_2_box)
		})
	})
	return cache_Data_DivisionRing_rightDiv
}

var cache_Data_DivisionRing_leftDiv gopurs_runtime.Value
var once_Data_DivisionRing_leftDiv sync.Once

func Get_Data_DivisionRing_leftDiv() gopurs_runtime.Value {
	once_Data_DivisionRing_leftDiv.Do(func() {
		cache_Data_DivisionRing_leftDiv = gopurs_runtime.Func3(func(dictDivisionRing_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_DivisionRing_leftDiv(gopurs_runtime.CoerceToStruct[Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value]](dictDivisionRing_0_box), a_1_box, b_2_box)
		})
	})
	return cache_Data_DivisionRing_leftDiv
}

var cache_Data_DivisionRing_divisionringNumber gopurs_runtime.Value
var once_Data_DivisionRing_divisionringNumber sync.Once

func Get_Data_DivisionRing_divisionringNumber() gopurs_runtime.Value {
	once_Data_DivisionRing_divisionringNumber.Do(func() {
		cache_Data_DivisionRing_divisionringNumber = gopurs_runtime.Value{Type: 9, IntVal: 2548491258, UnsafePtr: unsafe.Pointer(Rebox_Data_DivisionRing_1264867830_2381053422((&Constructor_Data_DivisionRing_DivisionRing[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_DivisionRing_2259598678_3060961102(Rebox_Data_DivisionRing_3060961102_2259598678(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](Get_Data_Ring_ringNumber()))))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float((1.0) / (x_0.FloatVal()))
		})})))}
	})
	return cache_Data_DivisionRing_divisionringNumber
}

type Constructor_Data_DivisionRing_DivisionRing[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f0044006100740061005f004400690076006900730069006f006e00520069006e0067:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f0044006100740061005f004400690076006900730069006f006e00520069006e0067005f004400690076006900730069006f006e00520069006e0067
func init() {
	gopurs_runtime.StructGetters[2548491258] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Ring0":
			return gopurs_runtime.Box(c.V0)
		case "recip":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Data_DivisionRing_DivisionRing: " + key)
		}
	}
}

func Call_Data_DivisionRing_DivisionRing_dollar_Dict(x_0_loop struct {
	Ring0 gopurs_runtime.Value
	recip gopurs_runtime.Value
	/*gopurs:record-v1:2:00520069006e00670030,00720065006300690070*/
}) *Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value] {
	var x_0 struct {
		Ring0 gopurs_runtime.Value
		recip gopurs_runtime.Value
		/*gopurs:record-v1:2:00520069006e00670030,00720065006300690070*/
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Ring0", "recip", orig.Ring0, orig.recip)
	}())
}

func Call_Data_DivisionRing_DivisionRing_dollar_Dict__2494093930(x_0_loop struct {
	Ring0 gopurs_runtime.Value
	recip gopurs_runtime.Value
	/*gopurs:record-v1:2:00520069006e00670030,00720065006300690070*/
}) *Constructor_Data_DivisionRing_DivisionRing[uint32] {
DivisionRing_dollar_Dict__2494093930:
	for {
		if false {
			continue DivisionRing_dollar_Dict__2494093930
		}
		var x_0 struct {
			Ring0 gopurs_runtime.Value
			recip gopurs_runtime.Value
			/*gopurs:record-v1:2:00520069006e00670030,00720065006300690070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_DivisionRing_DivisionRing[uint32]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Ring0", "recip", orig.Ring0, orig.recip)
		}())
	}
}

func Call_Data_DivisionRing_DivisionRing_dollar_Dict__1513685994(x_0_loop struct {
	Ring0 gopurs_runtime.Value
	recip gopurs_runtime.Value
	/*gopurs:record-v1:2:00520069006e00670030,00720065006300690070*/
}) *Constructor_Data_DivisionRing_DivisionRing[float64] {
DivisionRing_dollar_Dict__1513685994:
	for {
		if false {
			continue DivisionRing_dollar_Dict__1513685994
		}
		var x_0 struct {
			Ring0 gopurs_runtime.Value
			recip gopurs_runtime.Value
			/*gopurs:record-v1:2:00520069006e00670030,00720065006300690070*/
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Data_DivisionRing_DivisionRing[float64]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict2("Ring0", "recip", orig.Ring0, orig.recip)
		}())
	}
}

func Call_Data_DivisionRing_recip(dict_0_loop *Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V1
}

func Call_Data_DivisionRing_rightDiv(dictDivisionRing_0_loop *Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictDivisionRing_0 *Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value] = dictDivisionRing_0_loop
	_ = dictDivisionRing_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictDivisionRing_0.V0, gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "mul"), a_1, gopurs_runtime.Apply(dictDivisionRing_0.V1, b_2))
}

func Call_Data_DivisionRing_leftDiv(dictDivisionRing_0_loop *Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictDivisionRing_0 *Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value] = dictDivisionRing_0_loop
	_ = dictDivisionRing_0
	var a_1 gopurs_runtime.Value = a_1_loop
	_ = a_1
	var b_2 gopurs_runtime.Value = b_2_loop
	_ = b_2
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictDivisionRing_0.V0, gopurs_runtime.Value{}), "Semiring0"), gopurs_runtime.Value{}), "mul"), gopurs_runtime.Apply(dictDivisionRing_0.V1, b_2), a_1)
}

func Rebox_Data_DivisionRing_1264867830_2381053422(in *Constructor_Data_DivisionRing_DivisionRing[float64]) *Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_DivisionRing_2259598678_3060961102(in *Constructor_Data_Ring_Ring[float64]) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ring_Ring[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_DivisionRing_2384868323_2381053422(in *Constructor_Data_DivisionRing_DivisionRing[uint32]) *Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_DivisionRing_DivisionRing[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_DivisionRing_3060961102_2259598678(in *Constructor_Data_Ring_Ring[gopurs_runtime.Value]) *Constructor_Data_Ring_Ring[float64] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ring_Ring[float64])(unsafe.Pointer(in))
}
