package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_Pattern_go__append gopurs_runtime.Value
var once_Data_String_Pattern_go__append sync.Once

func Get_Data_String_Pattern_go__append() gopurs_runtime.Value {
	once_Data_String_Pattern_go__append.Do(func() {
		cache_Data_String_Pattern_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_String_Pattern_go__append
}

var cache_Data_String_Pattern_show gopurs_runtime.Value
var once_Data_String_Pattern_show sync.Once

func Get_Data_String_Pattern_show() gopurs_runtime.Value {
	once_Data_String_Pattern_show.Do(func() {
		cache_Data_String_Pattern_show = Get_Data_Show_show__3588520802()
	})
	return cache_Data_String_Pattern_show
}

var cache_Data_String_Pattern_eq gopurs_runtime.Value
var once_Data_String_Pattern_eq sync.Once

func Get_Data_String_Pattern_eq() gopurs_runtime.Value {
	once_Data_String_Pattern_eq.Do(func() {
		cache_Data_String_Pattern_eq = Get_Data_Eq_eq__3478410105()
	})
	return cache_Data_String_Pattern_eq
}

var cache_Data_String_Pattern_compare gopurs_runtime.Value
var once_Data_String_Pattern_compare sync.Once

func Get_Data_String_Pattern_compare() gopurs_runtime.Value {
	once_Data_String_Pattern_compare.Do(func() {
		cache_Data_String_Pattern_compare = Get_Data_Ord_compare__3009265731()
	})
	return cache_Data_String_Pattern_compare
}

var cache_Data_String_Pattern_Replacement gopurs_runtime.Value
var once_Data_String_Pattern_Replacement sync.Once

func Get_Data_String_Pattern_Replacement() gopurs_runtime.Value {
	once_Data_String_Pattern_Replacement.Do(func() {
		cache_Data_String_Pattern_Replacement = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_Pattern_Replacement(x_0_box.StrVal()))
		})
	})
	return cache_Data_String_Pattern_Replacement
}

var cache_Data_String_Pattern_Pattern gopurs_runtime.Value
var once_Data_String_Pattern_Pattern sync.Once

func Get_Data_String_Pattern_Pattern() gopurs_runtime.Value {
	once_Data_String_Pattern_Pattern.Do(func() {
		cache_Data_String_Pattern_Pattern = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_Pattern_Pattern(x_0_box.StrVal()))
		})
	})
	return cache_Data_String_Pattern_Pattern
}

var cache_Data_String_Pattern_showReplacement gopurs_runtime.Value
var once_Data_String_Pattern_showReplacement sync.Once

func Get_Data_String_Pattern_showReplacement() gopurs_runtime.Value {
	once_Data_String_Pattern_showReplacement.Do(func() {
		cache_Data_String_Pattern_showReplacement = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_1514099793_1386611502((&Constructor_Data_Show_Show[string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Replacement ", Data_Show_ShowStringImpl(v_0.StrVal())), ")"))
		})})))}
	})
	return cache_Data_String_Pattern_showReplacement
}

var cache_Data_String_Pattern_showPattern gopurs_runtime.Value
var once_Data_String_Pattern_showPattern sync.Once

func Get_Data_String_Pattern_showPattern() gopurs_runtime.Value {
	once_Data_String_Pattern_showPattern.Do(func() {
		cache_Data_String_Pattern_showPattern = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_1514099793_1386611502((&Constructor_Data_Show_Show[string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(Pattern ", Data_Show_ShowStringImpl(v_0.StrVal())), ")"))
		})})))}
	})
	return cache_Data_String_Pattern_showPattern
}

var cache_Data_String_Pattern_newtypeReplacement gopurs_runtime.Value
var once_Data_String_Pattern_newtypeReplacement sync.Once

func Get_Data_String_Pattern_newtypeReplacement() gopurs_runtime.Value {
	once_Data_String_Pattern_newtypeReplacement.Do(func() {
		cache_Data_String_Pattern_newtypeReplacement = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_2199435624_385277032((&Constructor_Data_Newtype_Newtype[string, string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_String_Pattern_newtypeReplacement
}

var cache_Data_String_Pattern_newtypePattern gopurs_runtime.Value
var once_Data_String_Pattern_newtypePattern sync.Once

func Get_Data_String_Pattern_newtypePattern() gopurs_runtime.Value {
	once_Data_String_Pattern_newtypePattern.Do(func() {
		cache_Data_String_Pattern_newtypePattern = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_2199435624_385277032((&Constructor_Data_Newtype_Newtype[string, string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_String_Pattern_newtypePattern
}

var cache_Data_String_Pattern_eqReplacement gopurs_runtime.Value
var once_Data_String_Pattern_eqReplacement sync.Once

func Get_Data_String_Pattern_eqReplacement() gopurs_runtime.Value {
	once_Data_String_Pattern_eqReplacement.Do(func() {
		cache_Data_String_Pattern_eqReplacement = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_1140313009_3790796878((&Constructor_Data_Eq_Eq[string]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((x_0.StrVal()) == (y_1.StrVal()))
		})})))}
	})
	return cache_Data_String_Pattern_eqReplacement
}

var cache_Data_String_Pattern_ordReplacement gopurs_runtime.Value
var once_Data_String_Pattern_ordReplacement sync.Once

func Get_Data_String_Pattern_ordReplacement() gopurs_runtime.Value {
	once_Data_String_Pattern_ordReplacement.Do(func() {
		cache_Data_String_Pattern_ordReplacement = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_2406510097_4177771502((&Constructor_Data_Ord_Ord[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_1140313009_3790796878(Rebox_Data_String_Pattern_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_String_Pattern_eqReplacement()))))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Rebox_Data_String_Pattern_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())).V1, gopurs_runtime.NormalizeStringValue(x_0), gopurs_runtime.NormalizeStringValue(y_1))
		})})))}
	})
	return cache_Data_String_Pattern_ordReplacement
}

var cache_Data_String_Pattern_eqPattern gopurs_runtime.Value
var once_Data_String_Pattern_eqPattern sync.Once

func Get_Data_String_Pattern_eqPattern() gopurs_runtime.Value {
	once_Data_String_Pattern_eqPattern.Do(func() {
		cache_Data_String_Pattern_eqPattern = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_1140313009_3790796878((&Constructor_Data_Eq_Eq[string]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((x_0.StrVal()) == (y_1.StrVal()))
		})})))}
	})
	return cache_Data_String_Pattern_eqPattern
}

var cache_Data_String_Pattern_ordPattern gopurs_runtime.Value
var once_Data_String_Pattern_ordPattern sync.Once

func Get_Data_String_Pattern_ordPattern() gopurs_runtime.Value {
	once_Data_String_Pattern_ordPattern.Do(func() {
		cache_Data_String_Pattern_ordPattern = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_2406510097_4177771502((&Constructor_Data_Ord_Ord[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_1140313009_3790796878(Rebox_Data_String_Pattern_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_String_Pattern_eqPattern()))))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Rebox_Data_String_Pattern_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())).V1, gopurs_runtime.NormalizeStringValue(x_0), gopurs_runtime.NormalizeStringValue(y_1))
		})})))}
	})
	return cache_Data_String_Pattern_ordPattern
}

func Call_Data_String_Pattern_Replacement(x_0_loop string) string {
	var x_0 string = x_0_loop
	_ = x_0
	return x_0
}

func Call_Data_String_Pattern_Pattern(x_0_loop string) string {
	var x_0 string = x_0_loop
	_ = x_0
	return x_0
}

func Rebox_Data_String_Pattern_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_Pattern_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_Pattern_2199435624_385277032(in *Constructor_Data_Newtype_Newtype[string, string]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_Pattern_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_Pattern_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[string])(unsafe.Pointer(in))
}

func Rebox_Data_String_Pattern_4177771502_2406510097(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[string])(unsafe.Pointer(in))
}
