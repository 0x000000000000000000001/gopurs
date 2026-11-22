package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Char_Gen_toEnumWithDefaults gopurs_runtime.Value
var once_Data_Char_Gen_toEnumWithDefaults sync.Once

func Get_Data_Char_Gen_toEnumWithDefaults() gopurs_runtime.Value {
	once_Data_Char_Gen_toEnumWithDefaults.Do(func() {
		cache_Data_Char_Gen_toEnumWithDefaults = Get_Data_Enum_toEnumWithDefaults__3799330569()
	})
	return cache_Data_Char_Gen_toEnumWithDefaults
}

var cache_Data_Char_Gen_bottom gopurs_runtime.Value
var once_Data_Char_Gen_bottom sync.Once

func Get_Data_Char_Gen_bottom() gopurs_runtime.Value {
	once_Data_Char_Gen_bottom.Do(func() {
		cache_Data_Char_Gen_bottom = gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar()))
	})
	return cache_Data_Char_Gen_bottom
}

var cache_Data_Char_Gen_top gopurs_runtime.Value
var once_Data_Char_Gen_top sync.Once

func Get_Data_Char_Gen_top() gopurs_runtime.Value {
	once_Data_Char_Gen_top.Do(func() {
		cache_Data_Char_Gen_top = gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar()))
	})
	return cache_Data_Char_Gen_top
}

var cache_Data_Char_Gen_foldable1NonEmpty gopurs_runtime.Value
var once_Data_Char_Gen_foldable1NonEmpty sync.Once

func Get_Data_Char_Gen_foldable1NonEmpty() gopurs_runtime.Value {
	once_Data_Char_Gen_foldable1NonEmpty.Do(func() {
		cache_Data_Char_Gen_foldable1NonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Char_Gen_4217626592_4151366573(Rebox_Data_Char_Gen_4151366573_4217626592(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))))}
	})
	return cache_Data_Char_Gen_foldable1NonEmpty
}

var cache_Data_Char_Gen_genUnicodeChar gopurs_runtime.Value
var once_Data_Char_Gen_genUnicodeChar sync.Once

func Get_Data_Char_Gen_genUnicodeChar() gopurs_runtime.Value {
	once_Data_Char_Gen_genUnicodeChar.Do(func() {
		cache_Data_Char_Gen_genUnicodeChar = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Char_Gen_genUnicodeChar(dictMonadGen_0_box)
		})
	})
	return cache_Data_Char_Gen_genUnicodeChar
}

var cache_Data_Char_Gen_genDigitChar gopurs_runtime.Value
var once_Data_Char_Gen_genDigitChar sync.Once

func Get_Data_Char_Gen_genDigitChar() gopurs_runtime.Value {
	once_Data_Char_Gen_genDigitChar.Do(func() {
		cache_Data_Char_Gen_genDigitChar = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Char_Gen_genDigitChar(dictMonadGen_0_box)
		})
	})
	return cache_Data_Char_Gen_genDigitChar
}

var cache_Data_Char_Gen_genAsciiChar_prime_ gopurs_runtime.Value
var once_Data_Char_Gen_genAsciiChar_prime_ sync.Once

func Get_Data_Char_Gen_genAsciiChar_prime_() gopurs_runtime.Value {
	once_Data_Char_Gen_genAsciiChar_prime_.Do(func() {
		cache_Data_Char_Gen_genAsciiChar_prime_ = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Char_Gen_genAsciiChar_prime_(dictMonadGen_0_box)
		})
	})
	return cache_Data_Char_Gen_genAsciiChar_prime_
}

var cache_Data_Char_Gen_genAsciiChar gopurs_runtime.Value
var once_Data_Char_Gen_genAsciiChar sync.Once

func Get_Data_Char_Gen_genAsciiChar() gopurs_runtime.Value {
	once_Data_Char_Gen_genAsciiChar.Do(func() {
		cache_Data_Char_Gen_genAsciiChar = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Char_Gen_genAsciiChar(dictMonadGen_0_box)
		})
	})
	return cache_Data_Char_Gen_genAsciiChar
}

var cache_Data_Char_Gen_genAlphaUppercase gopurs_runtime.Value
var once_Data_Char_Gen_genAlphaUppercase sync.Once

func Get_Data_Char_Gen_genAlphaUppercase() gopurs_runtime.Value {
	once_Data_Char_Gen_genAlphaUppercase.Do(func() {
		cache_Data_Char_Gen_genAlphaUppercase = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Char_Gen_genAlphaUppercase(dictMonadGen_0_box)
		})
	})
	return cache_Data_Char_Gen_genAlphaUppercase
}

var cache_Data_Char_Gen_genAlphaLowercase gopurs_runtime.Value
var once_Data_Char_Gen_genAlphaLowercase sync.Once

func Get_Data_Char_Gen_genAlphaLowercase() gopurs_runtime.Value {
	once_Data_Char_Gen_genAlphaLowercase.Do(func() {
		cache_Data_Char_Gen_genAlphaLowercase = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Char_Gen_genAlphaLowercase(dictMonadGen_0_box)
		})
	})
	return cache_Data_Char_Gen_genAlphaLowercase
}

var cache_Data_Char_Gen_genAlpha gopurs_runtime.Value
var once_Data_Char_Gen_genAlpha sync.Once

func Get_Data_Char_Gen_genAlpha() gopurs_runtime.Value {
	once_Data_Char_Gen_genAlpha.Do(func() {
		cache_Data_Char_Gen_genAlpha = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Char_Gen_genAlpha(dictMonadGen_0_box)
		})
	})
	return cache_Data_Char_Gen_genAlpha
}

var cache_Data_Char_Gen___cached_dict_0 gopurs_runtime.Value
var once_Data_Char_Gen___cached_dict_0 sync.Once

func Get_Data_Char_Gen___cached_dict_0() gopurs_runtime.Value {
	once_Data_Char_Gen___cached_dict_0.Do(func() {
		cache_Data_Char_Gen___cached_dict_0 = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Char_Gen_4217626592_4151366573(Rebox_Data_Char_Gen_4151366573_4217626592(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))))}
	})
	return cache_Data_Char_Gen___cached_dict_0
}

func Call_Data_Char_Gen_genUnicodeChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3799330569(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar()))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(65536))))
}

func Call_Data_Char_Gen_genDigitChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3799330569(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar()))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(48)), gopurs_runtime.Int(int64(57))))
}

func Call_Data_Char_Gen_genAsciiChar_prime_(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3799330569(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar()))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(127))))
}

func Call_Data_Char_Gen_genAsciiChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3799330569(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar()))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(32)), gopurs_runtime.Int(int64(127))))
}

func Call_Data_Char_Gen_genAlphaUppercase(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3799330569(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar()))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(65)), gopurs_runtime.Int(int64(90))))
}

func Call_Data_Char_Gen_genAlphaLowercase(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3799330569(), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_bottomChar())), gopurs_runtime.NormalizeStringValue(gopurs_runtime.NormalizeStringValue(Get_Data_Bounded_topChar()))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(97)), gopurs_runtime.Int(int64(122))))
}

func Call_Data_Char_Gen_genAlpha(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
	_ = dictMonadGen_0
	return gopurs_runtime.Apply(Call_Control_Monad_Gen_oneOf(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0), Rebox_Data_Char_Gen_4217626592_4151366573(Rebox_Data_Char_Gen_4151366573_4217626592(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Char_Gen___cached_dict_0())))), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_Data_Char_Gen_genAlphaLowercase(gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0))}), gopurs_runtime.Array([]gopurs_runtime.Value{Call_Data_Char_Gen_genAlphaUppercase(dictMonadGen_0)})}))})
}

func Rebox_Data_Char_Gen_4151366573_4217626592(in *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]])(unsafe.Pointer(in))
}

func Rebox_Data_Char_Gen_4217626592_4151366573(in *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value])(unsafe.Pointer(in))
}
