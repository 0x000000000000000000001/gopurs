package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_CaseInsensitive_go__append gopurs_runtime.Value
var once_Data_String_CaseInsensitive_go__append sync.Once

func Get_Data_String_CaseInsensitive_go__append() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_go__append.Do(func() {
		cache_Data_String_CaseInsensitive_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Data_String_CaseInsensitive_go__append
}

var cache_Data_String_CaseInsensitive_show gopurs_runtime.Value
var once_Data_String_CaseInsensitive_show sync.Once

func Get_Data_String_CaseInsensitive_show() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_show.Do(func() {
		cache_Data_String_CaseInsensitive_show = Get_Data_Show_show__3588520802()
	})
	return cache_Data_String_CaseInsensitive_show
}

var cache_Data_String_CaseInsensitive_eq gopurs_runtime.Value
var once_Data_String_CaseInsensitive_eq sync.Once

func Get_Data_String_CaseInsensitive_eq() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_eq.Do(func() {
		cache_Data_String_CaseInsensitive_eq = Get_Data_Eq_eq__3478410105()
	})
	return cache_Data_String_CaseInsensitive_eq
}

var cache_Data_String_CaseInsensitive_compare gopurs_runtime.Value
var once_Data_String_CaseInsensitive_compare sync.Once

func Get_Data_String_CaseInsensitive_compare() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_compare.Do(func() {
		cache_Data_String_CaseInsensitive_compare = Get_Data_Ord_compare__3009265731()
	})
	return cache_Data_String_CaseInsensitive_compare
}

var cache_Data_String_CaseInsensitive_CaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_CaseInsensitiveString sync.Once

func Get_Data_String_CaseInsensitive_CaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_CaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_CaseInsensitiveString = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_String_CaseInsensitive_CaseInsensitiveString(x_0_box.StrVal()))
		})
	})
	return cache_Data_String_CaseInsensitive_CaseInsensitiveString
}

var cache_Data_String_CaseInsensitive_showCaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_showCaseInsensitiveString sync.Once

func Get_Data_String_CaseInsensitive_showCaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_showCaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_showCaseInsensitiveString = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CaseInsensitive_1514099793_1386611502((&Constructor_Data_Show_Show[string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(gopurs_runtime.ConcatString(gopurs_runtime.ConcatString("(CaseInsensitiveString ", Data_Show_ShowStringImpl(v_0.StrVal())), ")"))
		})})))}
	})
	return cache_Data_String_CaseInsensitive_showCaseInsensitiveString
}

var cache_Data_String_CaseInsensitive_newtypeCaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_newtypeCaseInsensitiveString sync.Once

func Get_Data_String_CaseInsensitive_newtypeCaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_newtypeCaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_newtypeCaseInsensitiveString = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CaseInsensitive_2199435624_385277032((&Constructor_Data_Newtype_Newtype[string, string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})})))}
	})
	return cache_Data_String_CaseInsensitive_newtypeCaseInsensitiveString
}

var cache_Data_String_CaseInsensitive_eqCaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_eqCaseInsensitiveString sync.Once

func Get_Data_String_CaseInsensitive_eqCaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_eqCaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_eqCaseInsensitiveString = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CaseInsensitive_1140313009_3790796878((&Constructor_Data_Eq_Eq[string]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Bool((Data_String_Common_ToLower(v_0.StrVal())) == (Data_String_Common_ToLower(v1_1.StrVal())))
		})})))}
	})
	return cache_Data_String_CaseInsensitive_eqCaseInsensitiveString
}

var cache_Data_String_CaseInsensitive_ordCaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_ordCaseInsensitiveString sync.Once

func Get_Data_String_CaseInsensitive_ordCaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_ordCaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_ordCaseInsensitiveString = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CaseInsensitive_2406510097_4177771502((&Constructor_Data_Ord_Ord[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CaseInsensitive_1140313009_3790796878(Rebox_Data_String_CaseInsensitive_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_String_CaseInsensitive_eqCaseInsensitiveString()))))}
		}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply2(Rebox_Data_String_CaseInsensitive_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString())).V1, gopurs_runtime.Str(Data_String_Common_ToLower(v_0.StrVal())), gopurs_runtime.Str(Data_String_Common_ToLower(v1_1.StrVal())))
		})})))}
	})
	return cache_Data_String_CaseInsensitive_ordCaseInsensitiveString
}

func Call_Data_String_CaseInsensitive_CaseInsensitiveString(x_0_loop string) string {
	var x_0 string = x_0_loop
	_ = x_0
	return x_0
}

func Rebox_Data_String_CaseInsensitive_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_CaseInsensitive_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_CaseInsensitive_2199435624_385277032(in *Constructor_Data_Newtype_Newtype[string, string]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_CaseInsensitive_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_String_CaseInsensitive_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[string])(unsafe.Pointer(in))
}

func Rebox_Data_String_CaseInsensitive_4177771502_2406510097(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[string] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[string])(unsafe.Pointer(in))
}
