package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_Gen_max gopurs_runtime.Value
var once_Data_String_Gen_max sync.Once

func Get_Data_String_Gen_max() gopurs_runtime.Value {
	once_Data_String_Gen_max.Do(func() {
		cache_Data_String_Gen_max = Get_Data_Ord_max__4009478382()
	})
	return cache_Data_String_Gen_max
}

var cache_Data_String_Gen_genString gopurs_runtime.Value
var once_Data_String_Gen_genString sync.Once

func Get_Data_String_Gen_genString() gopurs_runtime.Value {
	once_Data_String_Gen_genString.Do(func() {
		cache_Data_String_Gen_genString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
		})
	})
	return cache_Data_String_Gen_genString
}

var cache_Data_String_Gen_genUnicodeString gopurs_runtime.Value
var once_Data_String_Gen_genUnicodeString sync.Once

func Get_Data_String_Gen_genUnicodeString() gopurs_runtime.Value {
	once_Data_String_Gen_genUnicodeString.Do(func() {
		cache_Data_String_Gen_genUnicodeString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_Gen_genUnicodeString(dictMonadRec_0_box, dictMonadGen_1_box)
		})
	})
	return cache_Data_String_Gen_genUnicodeString
}

var cache_Data_String_Gen_genDigitString gopurs_runtime.Value
var once_Data_String_Gen_genDigitString sync.Once

func Get_Data_String_Gen_genDigitString() gopurs_runtime.Value {
	once_Data_String_Gen_genDigitString.Do(func() {
		cache_Data_String_Gen_genDigitString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_Gen_genDigitString(dictMonadRec_0_box, dictMonadGen_1_box)
		})
	})
	return cache_Data_String_Gen_genDigitString
}

var cache_Data_String_Gen_genAsciiString_prime_ gopurs_runtime.Value
var once_Data_String_Gen_genAsciiString_prime_ sync.Once

func Get_Data_String_Gen_genAsciiString_prime_() gopurs_runtime.Value {
	once_Data_String_Gen_genAsciiString_prime_.Do(func() {
		cache_Data_String_Gen_genAsciiString_prime_ = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_Gen_genAsciiString_prime_(dictMonadRec_0_box, dictMonadGen_1_box)
		})
	})
	return cache_Data_String_Gen_genAsciiString_prime_
}

var cache_Data_String_Gen_genAsciiString gopurs_runtime.Value
var once_Data_String_Gen_genAsciiString sync.Once

func Get_Data_String_Gen_genAsciiString() gopurs_runtime.Value {
	once_Data_String_Gen_genAsciiString.Do(func() {
		cache_Data_String_Gen_genAsciiString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_Gen_genAsciiString(dictMonadRec_0_box, dictMonadGen_1_box)
		})
	})
	return cache_Data_String_Gen_genAsciiString
}

var cache_Data_String_Gen_genAlphaUppercaseString gopurs_runtime.Value
var once_Data_String_Gen_genAlphaUppercaseString sync.Once

func Get_Data_String_Gen_genAlphaUppercaseString() gopurs_runtime.Value {
	once_Data_String_Gen_genAlphaUppercaseString.Do(func() {
		cache_Data_String_Gen_genAlphaUppercaseString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_Gen_genAlphaUppercaseString(dictMonadRec_0_box, dictMonadGen_1_box)
		})
	})
	return cache_Data_String_Gen_genAlphaUppercaseString
}

var cache_Data_String_Gen_genAlphaString gopurs_runtime.Value
var once_Data_String_Gen_genAlphaString sync.Once

func Get_Data_String_Gen_genAlphaString() gopurs_runtime.Value {
	once_Data_String_Gen_genAlphaString.Do(func() {
		cache_Data_String_Gen_genAlphaString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_Gen_genAlphaString(dictMonadRec_0_box, dictMonadGen_1_box)
		})
	})
	return cache_Data_String_Gen_genAlphaString
}

var cache_Data_String_Gen_genAlphaLowercaseString gopurs_runtime.Value
var once_Data_String_Gen_genAlphaLowercaseString sync.Once

func Get_Data_String_Gen_genAlphaLowercaseString() gopurs_runtime.Value {
	once_Data_String_Gen_genAlphaLowercaseString.Do(func() {
		cache_Data_String_Gen_genAlphaLowercaseString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_String_Gen_genAlphaLowercaseString(dictMonadRec_0_box, dictMonadGen_1_box)
		})
	})
	return cache_Data_String_Gen_genAlphaLowercaseString
}

func Call_Data_String_Gen_genString(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
	_ = dictMonadGen_1
	// TAST (Let): Monad0_2_0 shape=App(Other) bindingType=Any
	Monad0_2_0 := gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{})
	_ = Monad0_2_0
	// TAST (Let): unfoldable1_3_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope1) [Char])] (TypeApp (TypeVar m$scope1) [(Array Char)]))
	unfoldable1_3_1 := gopurs_runtime.Apply(Call_Control_Monad_Gen_unfoldable(dictMonadRec_0, dictMonadGen_1), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray()))})
	_ = unfoldable1_3_1
	return gopurs_runtime.Func(func(genChar_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Func(func(size_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(dictMonadGen_1.V3, gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(Call_Data_Ord_max__4009478382(int64(1), size_5.IntVal))), gopurs_runtime.Func(func(newSize_6 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply2(dictMonadGen_1.V4, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
					return newSize_6
				}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_String_CodeUnits_fromCharArray(), gopurs_runtime.Apply(unfoldable1_3_1, genChar_4)))
			}))
		}))
	})
}

func Call_Data_String_Gen_genUnicodeString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
	_ = dictMonadGen_1
	return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), Call_Data_Char_Gen_genUnicodeChar(dictMonadGen_1))
}

func Call_Data_String_Gen_genDigitString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
	_ = dictMonadGen_1
	return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), Call_Data_Char_Gen_genDigitChar(dictMonadGen_1))
}

func Call_Data_String_Gen_genAsciiString_prime_(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
	_ = dictMonadGen_1
	return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), Call_Data_Char_Gen_genAsciiChar_prime_(dictMonadGen_1))
}

func Call_Data_String_Gen_genAsciiString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
	_ = dictMonadGen_1
	return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), Call_Data_Char_Gen_genAsciiChar(dictMonadGen_1))
}

func Call_Data_String_Gen_genAlphaUppercaseString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
	_ = dictMonadGen_1
	return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), Call_Data_Char_Gen_genAlphaUppercase(dictMonadGen_1))
}

func Call_Data_String_Gen_genAlphaString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
	_ = dictMonadGen_1
	return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), Call_Data_Char_Gen_genAlpha(dictMonadGen_1))
}

func Call_Data_String_Gen_genAlphaLowercaseString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
	_ = dictMonadRec_0
	var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
	_ = dictMonadGen_1
	return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), Call_Data_Char_Gen_genAlphaLowercase(dictMonadGen_1))
}
