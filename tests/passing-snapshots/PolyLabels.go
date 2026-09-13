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
		cache_Main_fooIsSymbol = gopurs_runtime.Value{Type: 9, IntVal: 2134024384, UnsafePtr: unsafe.Pointer((&Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
			return Call_Main_set(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_0_box), _dollar___unused_1_box, _dollar___unused_2_box, uint32(l_3_box.IntVal))
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

var cache_Main_setFoo__1386493764 gopurs_runtime.Value
var once_Main_setFoo__1386493764 sync.Once

func Get_Main_setFoo__1386493764() gopurs_runtime.Value {
	once_Main_setFoo__1386493764.Do(func() {
		cache_Main_setFoo__1386493764 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_setFoo__1386493764(__eta_norm_1_0_box.StrVal(), __eta_norm_0_1_box)
		})
	})
	return cache_Main_setFoo__1386493764
}

var cache_Main_get gopurs_runtime.Value
var once_Main_get sync.Once

func Get_Main_get() gopurs_runtime.Value {
	once_Main_get.Do(func() {
		cache_Main_get = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_get(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_0_box), _dollar___unused_1_box, uint32(l_2_box.IntVal))
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

var cache_Main_getFoo__1715234564 gopurs_runtime.Value
var once_Main_getFoo__1715234564 sync.Once

func Get_Main_getFoo__1715234564() gopurs_runtime.Value {
	once_Main_getFoo__1715234564.Do(func() {
		cache_Main_getFoo__1715234564 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Main_getFoo__1715234564(__eta_norm_0_0_box))
		})
	})
	return cache_Main_getFoo__1715234564
}

var cache_Main_lens gopurs_runtime.Value
var once_Main_lens sync.Once

func Get_Main_lens() gopurs_runtime.Value {
	once_Main_lens.Do(func() {
		cache_Main_lens = gopurs_runtime.Func7(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, _dollar___unused_2_box gopurs_runtime.Value, dictFunctor_3_box gopurs_runtime.Value, l_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, r_6_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_lens(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dictIsSymbol_0_box), _dollar___unused_1_box, _dollar___unused_2_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_3_box), uint32(l_4_box.IntVal), f_5_box, r_6_box)
		})
	})
	return cache_Main_lens
}

var cache_Main_fooLens gopurs_runtime.Value
var once_Main_fooLens sync.Once

func Get_Main_fooLens() gopurs_runtime.Value {
	once_Main_fooLens.Do(func() {
		cache_Main_fooLens = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fooLens(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box))
		})
	})
	return cache_Main_fooLens
}

var cache_Main_fooLens__3333359234 gopurs_runtime.Value
var once_Main_fooLens__3333359234 sync.Once

func Get_Main_fooLens__3333359234() gopurs_runtime.Value {
	once_Main_fooLens__3333359234.Do(func() {
		cache_Main_fooLens__3333359234 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_fooLens__3333359234(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Main_fooLens__3333359234
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.Apply(Call_Main_lens(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](Get_Main_fooIsSymbol()), gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()), 513803634, gopurs_runtime.Apply(Get_Effect_Console_logShow(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), gopurs_runtime.RecordDict1("foo", gopurs_runtime.Int(int64(1)))), gopurs_runtime.Value{})
			_ = __local_var_0_0
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Main_unsafeGet(), gopurs_runtime.Str("foo"), gopurs_runtime.Apply3(Get_Main_unsafeSet(), gopurs_runtime.Str("foo"), gopurs_runtime.Str("Done"), gopurs_runtime.RecordDict1("foo", gopurs_runtime.Int(int64(1))))).StrVal())), gopurs_runtime.Value{})
		})
	})
	return cache_Main_main
}

func Call_Main_set(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value, l_3_loop uint32) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var l_3 uint32 = l_3_loop
	_ = l_3
	return gopurs_runtime.Apply(Get_Main_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(dictIsSymbol_0.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l_3), UnsafePtr: nil}).StrVal()))
}

func Call_Main_setFoo__1386493764(__eta_norm_1_0_loop string, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
setFoo__1386493764:
	for {
		if false {
			continue setFoo__1386493764
		}
		var __eta_norm_1_0 string = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return gopurs_runtime.Apply3(Get_Main_unsafeSet(), gopurs_runtime.Str("foo"), gopurs_runtime.Str(__eta_norm_1_0), __eta_norm_0_1)
	}
}

func Call_Main_get(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], _dollar___unused_1_loop gopurs_runtime.Value, l_2_loop uint32) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var l_2 uint32 = l_2_loop
	_ = l_2
	return gopurs_runtime.Apply(Get_Main_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(dictIsSymbol_0.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l_2), UnsafePtr: nil}).StrVal()))
}

func Call_Main_getFoo__1715234564(__eta_norm_0_0_loop gopurs_runtime.Value) string {
getFoo__1715234564:
	for {
		if false {
			continue getFoo__1715234564
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Apply2(Get_Main_unsafeGet(), gopurs_runtime.Str("foo"), __eta_norm_0_0).StrVal()
	}
}

func Call_Main_lens(dictIsSymbol_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value], _dollar___unused_1_loop gopurs_runtime.Value, _dollar___unused_2_loop gopurs_runtime.Value, dictFunctor_3_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], l_4_loop uint32, f_5_loop gopurs_runtime.Value, r_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictIsSymbol_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dictIsSymbol_0_loop
	_ = dictIsSymbol_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var _dollar___unused_2 gopurs_runtime.Value = _dollar___unused_2_loop
	_ = _dollar___unused_2
	var dictFunctor_3 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_3_loop
	_ = dictFunctor_3
	var l_4 uint32 = l_4_loop
	_ = l_4
	var f_5 gopurs_runtime.Value = f_5_loop
	_ = f_5
	var r_6 gopurs_runtime.Value = r_6_loop
	_ = r_6
	return gopurs_runtime.Apply2(dictFunctor_3.V0, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(Get_Main_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(dictIsSymbol_0.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l_4), UnsafePtr: nil}).StrVal()), a_7, r_6)
	}), gopurs_runtime.Apply(f_5, gopurs_runtime.Apply2(Get_Main_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(dictIsSymbol_0.V0, gopurs_runtime.Value{Type: 9, IntVal: int64(l_4), UnsafePtr: nil}).StrVal()), r_6)))
}

func Call_Main_fooLens(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Apply5(Get_Main_lens(), gopurs_runtime.Value{Type: 9, IntVal: 2134024384, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](Get_Main_fooIsSymbol()))}, gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(dictFunctor_0)}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
}

func Call_Main_fooLens__3333359234(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
fooLens__3333359234:
	for {
		if false {
			continue fooLens__3333359234
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return Call_Main_lens(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](Get_Main_fooIsSymbol()), gopurs_runtime.Value{}, gopurs_runtime.Value{}, gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()), 513803634, gopurs_runtime.Apply(Get_Effect_Console_logShow(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))))}), __eta_norm_0_1)
	}
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Get_Main_unsafeGet() gopurs_runtime.Value {
	return _Gopurs_Main_UnsafeGet
}

func Get_Main_unsafeSet() gopurs_runtime.Value {
	return _Gopurs_Main_UnsafeSet
}
