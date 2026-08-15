package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_LBox gopurs_runtime.Value
var once_Main_LBox sync.Once

func Get_Main_LBox() gopurs_runtime.Value {
	once_Main_LBox.Do(func() {
		cache_Main_LBox = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_LBox(x_0_box)
		})
	})
	return cache_Main_LBox
}

var cache_Main_unLBox gopurs_runtime.Value
var once_Main_unLBox sync.Once

func Get_Main_unLBox() gopurs_runtime.Value {
	once_Main_unLBox.Do(func() {
		cache_Main_unLBox = gopurs_runtime.Func2(func(g_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_unLBox(g_0_box, v_1_box)
		})
	})
	return cache_Main_unLBox
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_lboxIdentity gopurs_runtime.Value
var once_Main_lboxIdentity sync.Once

func Get_Main_lboxIdentity() gopurs_runtime.Value {
	once_Main_lboxIdentity.Do(func() {
		cache_Main_lboxIdentity = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_lboxIdentity(v_0_box)
		})
	})
	return cache_Main_lboxIdentity
}

var cache_Main_get gopurs_runtime.Value
var once_Main_get sync.Once

func Get_Main_get() gopurs_runtime.Value {
	once_Main_get.Do(func() {
		cache_Main_get = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_0_box), _dollar___unused_1_box, uint32(l_2_box.IntVal), r_3_box)
		})
	})
	return cache_Main_get
}

var cache_Main_read gopurs_runtime.Value
var once_Main_read sync.Once

func Get_Main_read() gopurs_runtime.Value {
	once_Main_read.Do(func() {
		cache_Main_read = gopurs_runtime.Func2(func(rec_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_read(rec_0_box, v_1_box)
		})
	})
	return cache_Main_read
}

var cache_Main_get__1126608997 gopurs_runtime.Value
var once_Main_get__1126608997 sync.Once

func Get_Main_get__1126608997() gopurs_runtime.Value {
	once_Main_get__1126608997.Do(func() {
		cache_Main_get__1126608997 = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get__1126608997(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_0_box), uint32(l_1_box.IntVal), r_2_box)
		})
	})
	return cache_Main_get__1126608997
}

var cache_Main_get__1573763043 gopurs_runtime.Value
var once_Main_get__1573763043 sync.Once

func Get_Main_get__1573763043() gopurs_runtime.Value {
	once_Main_get__1573763043.Do(func() {
		cache_Main_get__1573763043 = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get__1573763043(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_0_box), _dollar___unused_1_box, uint32(l_2_box.IntVal), r_3_box)
		})
	})
	return cache_Main_get__1573763043
}

var cache_Main_unLBox__2682485592 gopurs_runtime.Value
var once_Main_unLBox__2682485592 sync.Once

func Get_Main_unLBox__2682485592() gopurs_runtime.Value {
	once_Main_unLBox__2682485592.Do(func() {
		cache_Main_unLBox__2682485592 = gopurs_runtime.Func2(func(g_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_unLBox__2682485592(g_0_box, v_1_box)
		})
	})
	return cache_Main_unLBox__2682485592
}

var cache_Main_unLBox__3788829592 gopurs_runtime.Value
var once_Main_unLBox__3788829592 sync.Once

func Get_Main_unLBox__3788829592() gopurs_runtime.Value {
	once_Main_unLBox__3788829592.Do(func() {
		cache_Main_unLBox__3788829592 = gopurs_runtime.Func2(func(g_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_unLBox__3788829592(g_0_box, v_1_box)
		})
	})
	return cache_Main_unLBox__3788829592
}

func Call_Main_LBox(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_unLBox(g_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var g_0 gopurs_runtime.Value = g_0_loop
	_ = g_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	// TAST (Let): g2_2_0 -> gopurs_runtime.Value
	g2_2_0 := gopurs_runtime.Apply(g_0, gopurs_runtime.Value{})
	_ = g2_2_0
	return gopurs_runtime.Apply(v_1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(dictIsSymbol_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(g2_2_0, dictIsSymbol_4)
		})
	}))
}

func Call_Main_lboxIdentity(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(v_0, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(dictIsSymbol_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(lbl_3 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply3(f_4, gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 2134024384, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_2))}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(lbl_3.IntVal)), UnsafePtr: nil})
				})
			})
		})
	}))
}

func Call_Main_get(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol, _dollar___unused_1_loop gopurs_runtime.Value, l_2_loop uint32, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var l_2 uint32 = l_2_loop
	_ = l_2
	var r_3 gopurs_runtime.Value = r_3_loop
	_ = r_3
	return gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictIsSymbol_0.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(l_2), UnsafePtr: nil}).StrVal()), r_3)
}

func Call_Main_read(rec_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var rec_0 gopurs_runtime.Value = rec_0_loop
	_ = rec_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(v_1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(dictIsSymbol_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(lbl_4 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_3, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(lbl_4.IntVal)), UnsafePtr: nil}).StrVal()), rec_0)
			})
		})
	}))
}

func Call_Main_get__1126608997(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol, l_1_loop uint32, r_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var l_1 uint32 = l_1_loop
	_ = l_1
	var r_2 gopurs_runtime.Value = r_2_loop
	_ = r_2
	return gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictIsSymbol_0.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(l_1), UnsafePtr: nil}).StrVal()), r_2)
}

func Call_Main_get__1573763043(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol, _dollar___unused_1_loop gopurs_runtime.Value, l_2_loop uint32, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var l_2 uint32 = l_2_loop
	_ = l_2
	var r_3 gopurs_runtime.Value = r_3_loop
	_ = r_3
	return gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictIsSymbol_0.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(l_2), UnsafePtr: nil}).StrVal()), r_3)
}

func Call_Main_unLBox__2682485592(g_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var g_0 gopurs_runtime.Value = g_0_loop
	_ = g_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	// TAST (Let): g2_2_0 -> gopurs_runtime.Value
	g2_2_0 := gopurs_runtime.Apply(g_0, gopurs_runtime.Value{})
	_ = g2_2_0
	return gopurs_runtime.Apply(v_1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(dictIsSymbol_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(g2_2_0, dictIsSymbol_4)
		})
	}))
}

func Call_Main_unLBox__3788829592(g_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var g_0 gopurs_runtime.Value = g_0_loop
	_ = g_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	// TAST (Let): g2_2_0 -> gopurs_runtime.Value
	g2_2_0 := gopurs_runtime.Apply(g_0, gopurs_runtime.Value{})
	_ = g2_2_0
	return gopurs_runtime.Apply(v_1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(dictIsSymbol_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(g2_2_0, dictIsSymbol_4)
		})
	}))
}
