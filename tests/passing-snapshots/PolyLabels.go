package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_fooIsSymbol gopurs_runtime.Value
var once_Main_fooIsSymbol sync.Once

func Get_Main_fooIsSymbol() gopurs_runtime.Value {
	once_Main_fooIsSymbol.Do(func() {
		cache_Main_fooIsSymbol = gopurs_runtime.Value{Type: 9, IntVal: 2134024384, UnsafePtr: unsafe.Pointer((&Constructor_Data_Symbol_IsSymbol{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str("foo")
		})}))}
	})
	return cache_Main_fooIsSymbol
}

var cache_Main_set gopurs_runtime.Value
var once_Main_set sync.Once

func Get_Main_set() gopurs_runtime.Value {
	once_Main_set.Do(func() {
		cache_Main_set = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_set(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_0_box), _dollar___unused_1_box, _dollar___unused_2_box, uint32(l_3_box.IntVal))
		})
	})
	return cache_Main_set
}

var cache_Main_setFoo gopurs_runtime.Value
var once_Main_setFoo sync.Once

func Get_Main_setFoo() gopurs_runtime.Value {
	once_Main_setFoo.Do(func() {
		cache_Main_setFoo = gopurs_runtime.Apply(Get_Main_unsafeSet(), gopurs_runtime.Str("foo"))
	})
	return cache_Main_setFoo
}

var cache_Main_get gopurs_runtime.Value
var once_Main_get sync.Once

func Get_Main_get() gopurs_runtime.Value {
	once_Main_get.Do(func() {
		cache_Main_get = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_0_box), _dollar___unused_1_box, uint32(l_2_box.IntVal))
		})
	})
	return cache_Main_get
}

var cache_Main_getFoo gopurs_runtime.Value
var once_Main_getFoo sync.Once

func Get_Main_getFoo() gopurs_runtime.Value {
	once_Main_getFoo.Do(func() {
		cache_Main_getFoo = gopurs_runtime.Apply(Get_Main_unsafeGet(), gopurs_runtime.Str("foo"))
	})
	return cache_Main_getFoo
}

var cache_Main_lens gopurs_runtime.Value
var once_Main_lens sync.Once

func Get_Main_lens() gopurs_runtime.Value {
	once_Main_lens.Do(func() {
		cache_Main_lens = gopurs_runtime.Func7(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, dictFunctor_3_box gopurs_runtime.Value, l_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, r_6_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_lens(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_0_box), _dollar___unused_1_box, _dollar___unused_2_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_3_box), uint32(l_4_box.IntVal), f_5_box, r_6_box)
		})
	})
	return cache_Main_lens
}

var cache_Main_fooLens gopurs_runtime.Value
var once_Main_fooLens sync.Once

func Get_Main_fooLens() gopurs_runtime.Value {
	once_Main_fooLens.Do(func() {
		cache_Main_fooLens = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fooLens(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, r_2_box)
		})
	})
	return cache_Main_fooLens
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.RecordDict1("foo", gopurs_runtime.Int(1))
			_ = __local_var_0_0
			// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
			__local_var_1_2 := gopurs_runtime.Apply(Get_Main_unsafeSet(), gopurs_runtime.Str("foo"))
			_ = __local_var_1_2
			// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
			__local_var_1_1 := gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(__local_var_1_2, a_2, __local_var_0_0)
			})
			_ = __local_var_1_1
			// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
			__local_var_2_4 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Apply2(Get_Main_unsafeGet(), gopurs_runtime.Str("foo"), __local_var_0_0)).StrVal()))
			_ = __local_var_2_4
			// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
			__local_var_2_3 := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					a_prime__3_5 := gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Value{})
					_ = a_prime__3_5
					return gopurs_runtime.Apply(__local_var_1_1, a_prime__3_5)
				})
			})
			_ = __local_var_2_3
			_dollar___unused_3_6 := gopurs_runtime.Apply(__local_var_2_3, gopurs_runtime.Value{})
			_ = _dollar___unused_3_6
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Main_getFoo(), gopurs_runtime.Apply2(Get_Main_setFoo(), gopurs_runtime.Str("Done"), gopurs_runtime.RecordDict1("foo", gopurs_runtime.Int(1)))).StrVal())), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

var cache_Main_fooLens__3068227129 gopurs_runtime.Value
var once_Main_fooLens__3068227129 sync.Once

func Get_Main_fooLens__3068227129() gopurs_runtime.Value {
	once_Main_fooLens__3068227129.Do(func() {
		cache_Main_fooLens__3068227129 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fooLens__3068227129(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, r_2_box)
		})
	})
	return cache_Main_fooLens__3068227129
}

var cache_Main_get__1925999372 gopurs_runtime.Value
var once_Main_get__1925999372 sync.Once

func Get_Main_get__1925999372() gopurs_runtime.Value {
	once_Main_get__1925999372.Do(func() {
		cache_Main_get__1925999372 = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get__1925999372(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_0_box), _dollar___unused_1_box, uint32(l_2_box.IntVal))
		})
	})
	return cache_Main_get__1925999372
}

var cache_Main_get__1126608997 gopurs_runtime.Value
var once_Main_get__1126608997 sync.Once

func Get_Main_get__1126608997() gopurs_runtime.Value {
	once_Main_get__1126608997.Do(func() {
		cache_Main_get__1126608997 = gopurs_runtime.Func(func(l_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get__1126608997(uint32(l_0_box.IntVal))
		})
	})
	return cache_Main_get__1126608997
}

var cache_Main_get__1573763043 gopurs_runtime.Value
var once_Main_get__1573763043 sync.Once

func Get_Main_get__1573763043() gopurs_runtime.Value {
	once_Main_get__1573763043.Do(func() {
		cache_Main_get__1573763043 = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get__1573763043(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_0_box), _dollar___unused_1_box, uint32(l_2_box.IntVal))
		})
	})
	return cache_Main_get__1573763043
}

var cache_Main_getFoo__81075578 gopurs_runtime.Value
var once_Main_getFoo__81075578 sync.Once

func Get_Main_getFoo__81075578() gopurs_runtime.Value {
	once_Main_getFoo__81075578.Do(func() {
		cache_Main_getFoo__81075578 = gopurs_runtime.Apply(Get_Main_unsafeGet(), gopurs_runtime.Str("foo"))
	})
	return cache_Main_getFoo__81075578
}

var cache_Main_lens__1789924085 gopurs_runtime.Value
var once_Main_lens__1789924085 sync.Once

func Get_Main_lens__1789924085() gopurs_runtime.Value {
	once_Main_lens__1789924085.Do(func() {
		cache_Main_lens__1789924085 = gopurs_runtime.Func4(func(dictFunctor_0_box gopurs_runtime.Value, l_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_lens__1789924085(uint32(dictFunctor_0_box.IntVal), l_1_box, f_2_box, r_3_box)
		})
	})
	return cache_Main_lens__1789924085
}

var cache_Main_lens__409319347 gopurs_runtime.Value
var once_Main_lens__409319347 sync.Once

func Get_Main_lens__409319347() gopurs_runtime.Value {
	once_Main_lens__409319347.Do(func() {
		cache_Main_lens__409319347 = gopurs_runtime.Func7(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, dictFunctor_3_box gopurs_runtime.Value, l_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, r_6_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_lens__409319347(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_0_box), _dollar___unused_1_box, _dollar___unused_2_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_3_box), uint32(l_4_box.IntVal), f_5_box, r_6_box)
		})
	})
	return cache_Main_lens__409319347
}

var cache_Main_set__413201609 gopurs_runtime.Value
var once_Main_set__413201609 sync.Once

func Get_Main_set__413201609() gopurs_runtime.Value {
	once_Main_set__413201609.Do(func() {
		cache_Main_set__413201609 = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_set__413201609(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_0_box), _dollar___unused_1_box, _dollar___unused_2_box, uint32(l_3_box.IntVal))
		})
	})
	return cache_Main_set__413201609
}

var cache_Main_set__3333893858 gopurs_runtime.Value
var once_Main_set__3333893858 sync.Once

func Get_Main_set__3333893858() gopurs_runtime.Value {
	once_Main_set__3333893858.Do(func() {
		cache_Main_set__3333893858 = gopurs_runtime.Func(func(l_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_set__3333893858(uint32(l_0_box.IntVal))
		})
	})
	return cache_Main_set__3333893858
}

var cache_Main_set__2051929444 gopurs_runtime.Value
var once_Main_set__2051929444 sync.Once

func Get_Main_set__2051929444() gopurs_runtime.Value {
	once_Main_set__2051929444.Do(func() {
		cache_Main_set__2051929444 = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_set__2051929444(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol](dictIsSymbol_0_box), _dollar___unused_1_box, _dollar___unused_2_box, uint32(l_3_box.IntVal))
		})
	})
	return cache_Main_set__2051929444
}

var cache_Main_setFoo__3907007033 gopurs_runtime.Value
var once_Main_setFoo__3907007033 sync.Once

func Get_Main_setFoo__3907007033() gopurs_runtime.Value {
	once_Main_setFoo__3907007033.Do(func() {
		cache_Main_setFoo__3907007033 = gopurs_runtime.Apply(Get_Main_unsafeSet(), gopurs_runtime.Str("foo"))
	})
	return cache_Main_setFoo__3907007033
}

func Call_Main_set(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol, _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value, l_3_loop uint32) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var l_3 uint32 = l_3_loop
	_ = l_3
	return gopurs_runtime.Apply(Get_Main_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictIsSymbol_0.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(l_3), UnsafePtr: nil}).StrVal()))
}

func Call_Main_get(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol, _dollar___unused_1_loop gopurs_runtime.Value, l_2_loop uint32) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var l_2 uint32 = l_2_loop
	_ = l_2
	return gopurs_runtime.Apply(Get_Main_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictIsSymbol_0.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(l_2), UnsafePtr: nil}).StrVal()))
}

func Call_Main_lens(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol, _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value, dictFunctor_3_loop *Constructor_Data_Functor_Functor, l_4_loop uint32, f_5_loop gopurs_runtime.Value, r_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var dictFunctor_3 *Constructor_Data_Functor_Functor = dictFunctor_3_loop
	_ = dictFunctor_3
	var l_4 uint32 = l_4_loop
	_ = l_4
	var f_5 gopurs_runtime.Value = f_5_loop
	_ = f_5
	var r_6 gopurs_runtime.Value = r_6_loop
	_ = r_6
	// TAST (Let): __local_var_7_0 -> gopurs_runtime.Value
	__local_var_7_0 := gopurs_runtime.Apply(Get_Main_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictIsSymbol_0.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(l_4), UnsafePtr: nil}).StrVal()))
	_ = __local_var_7_0
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_3.V0), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(__local_var_7_0, a_8, r_6)
	}), gopurs_runtime.Apply(f_5, gopurs_runtime.Apply2(Get_Main_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictIsSymbol_0.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(l_4), UnsafePtr: nil}).StrVal()), r_6)))
}

func Call_Main_fooLens(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
	_ = dictFunctor_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var r_2 gopurs_runtime.Value = r_2_loop
	_ = r_2
	// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
	__local_var_3_0 := gopurs_runtime.Apply(Get_Main_unsafeSet(), gopurs_runtime.Str("foo"))
	_ = __local_var_3_0
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(__local_var_3_0, a_4, r_2)
	}), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(Get_Main_unsafeGet(), gopurs_runtime.Str("foo"), r_2)))
}

func Call_Main_fooLens__3068227129(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
	_ = dictFunctor_0
	var f_1 gopurs_runtime.Value = f_1_loop
	_ = f_1
	var r_2 gopurs_runtime.Value = r_2_loop
	_ = r_2
	// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
	__local_var_3_0 := gopurs_runtime.Apply(Get_Main_unsafeSet(), gopurs_runtime.Str("foo"))
	_ = __local_var_3_0
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(__local_var_3_0, a_4, r_2)
	}), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(Get_Main_unsafeGet(), gopurs_runtime.Str("foo"), r_2)))
}

func Call_Main_get__1925999372(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol, _dollar___unused_1_loop gopurs_runtime.Value, l_2_loop uint32) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var l_2 uint32 = l_2_loop
	_ = l_2
	return gopurs_runtime.Apply(Get_Main_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictIsSymbol_0.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(l_2), UnsafePtr: nil}).StrVal()))
}

func Call_Main_get__1126608997(l_0_loop uint32) gopurs_runtime.Value {
	var l_0 uint32 = l_0_loop
	_ = l_0
	return gopurs_runtime.Apply(Get_Main_unsafeGet(), gopurs_runtime.Str("foo"))
}

func Call_Main_get__1573763043(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol, _dollar___unused_1_loop gopurs_runtime.Value, l_2_loop uint32) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var l_2 uint32 = l_2_loop
	_ = l_2
	return gopurs_runtime.Apply(Get_Main_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictIsSymbol_0.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(l_2), UnsafePtr: nil}).StrVal()))
}

func Call_Main_lens__1789924085(dictFunctor_0_loop uint32, l_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 uint32 = dictFunctor_0_loop
	_ = dictFunctor_0
	var l_1 gopurs_runtime.Value = l_1_loop
	_ = l_1
	var f_2 gopurs_runtime.Value = f_2_loop
	_ = f_2
	var r_3 gopurs_runtime.Value = r_3_loop
	_ = r_3
	// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
	__local_var_4_0 := gopurs_runtime.Apply(Get_Main_unsafeSet(), gopurs_runtime.Str("foo"))
	_ = __local_var_4_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 9, IntVal: int64(dictFunctor_0), UnsafePtr: nil}, "map"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(__local_var_4_0, a_5, r_3)
	}), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply2(Get_Main_unsafeGet(), gopurs_runtime.Str("foo"), r_3)))
}

func Call_Main_lens__409319347(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol, _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value, dictFunctor_3_loop *Constructor_Data_Functor_Functor, l_4_loop uint32, f_5_loop gopurs_runtime.Value, r_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var dictFunctor_3 *Constructor_Data_Functor_Functor = dictFunctor_3_loop
	_ = dictFunctor_3
	var l_4 uint32 = l_4_loop
	_ = l_4
	var f_5 gopurs_runtime.Value = f_5_loop
	_ = f_5
	var r_6 gopurs_runtime.Value = r_6_loop
	_ = r_6
	// TAST (Let): __local_var_7_0 -> gopurs_runtime.Value
	__local_var_7_0 := gopurs_runtime.Apply(Get_Main_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictIsSymbol_0.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(l_4), UnsafePtr: nil}).StrVal()))
	_ = __local_var_7_0
	return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_3.V0), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(__local_var_7_0, a_8, r_6)
	}), gopurs_runtime.Apply(f_5, gopurs_runtime.Apply2(Get_Main_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictIsSymbol_0.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(l_4), UnsafePtr: nil}).StrVal()), r_6)))
}

func Call_Main_set__413201609(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol, _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value, l_3_loop uint32) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var l_3 uint32 = l_3_loop
	_ = l_3
	return gopurs_runtime.Apply(Get_Main_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictIsSymbol_0.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(l_3), UnsafePtr: nil}).StrVal()))
}

func Call_Main_set__3333893858(l_0_loop uint32) gopurs_runtime.Value {
	var l_0 uint32 = l_0_loop
	_ = l_0
	return gopurs_runtime.Apply(Get_Main_unsafeSet(), gopurs_runtime.Str("foo"))
}

func Call_Main_set__2051929444(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol, _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value, l_3_loop uint32) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var l_3 uint32 = l_3_loop
	_ = l_3
	return gopurs_runtime.Apply(Get_Main_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.Box(dictIsSymbol_0.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(l_3), UnsafePtr: nil}).StrVal()))
}

func Get_Main_unsafeGet() gopurs_runtime.Value {
	return _Gopurs_Main_UnsafeGet
}

func Get_Main_unsafeSet() gopurs_runtime.Value {
	return _Gopurs_Main_UnsafeSet
}
