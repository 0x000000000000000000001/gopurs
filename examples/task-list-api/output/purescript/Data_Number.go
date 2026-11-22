package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Number_tau gopurs_runtime.Value
var once_Data_Number_tau sync.Once

func Get_Data_Number_tau() gopurs_runtime.Value {
	once_Data_Number_tau.Do(func() {
		cache_Data_Number_tau = gopurs_runtime.Float(6.283185307179586)
	})
	return cache_Data_Number_tau
}

var cache_Data_Number_sqrt2 gopurs_runtime.Value
var once_Data_Number_sqrt2 sync.Once

func Get_Data_Number_sqrt2() gopurs_runtime.Value {
	once_Data_Number_sqrt2.Do(func() {
		cache_Data_Number_sqrt2 = gopurs_runtime.Float(1.4142135623730951)
	})
	return cache_Data_Number_sqrt2
}

var cache_Data_Number_sqrt1_2 gopurs_runtime.Value
var once_Data_Number_sqrt1_2 sync.Once

func Get_Data_Number_sqrt1_2() gopurs_runtime.Value {
	once_Data_Number_sqrt1_2.Do(func() {
		cache_Data_Number_sqrt1_2 = gopurs_runtime.Float(0.7071067811865476)
	})
	return cache_Data_Number_sqrt1_2
}

var cache_Data_Number_pi gopurs_runtime.Value
var once_Data_Number_pi sync.Once

func Get_Data_Number_pi() gopurs_runtime.Value {
	once_Data_Number_pi.Do(func() {
		cache_Data_Number_pi = gopurs_runtime.Float(3.141592653589793)
	})
	return cache_Data_Number_pi
}

var cache_Data_Number_log2e gopurs_runtime.Value
var once_Data_Number_log2e sync.Once

func Get_Data_Number_log2e() gopurs_runtime.Value {
	once_Data_Number_log2e.Do(func() {
		cache_Data_Number_log2e = gopurs_runtime.Float(1.4426950408889634)
	})
	return cache_Data_Number_log2e
}

var cache_Data_Number_log10e gopurs_runtime.Value
var once_Data_Number_log10e sync.Once

func Get_Data_Number_log10e() gopurs_runtime.Value {
	once_Data_Number_log10e.Do(func() {
		cache_Data_Number_log10e = gopurs_runtime.Float(0.4342944819032518)
	})
	return cache_Data_Number_log10e
}

var cache_Data_Number_ln2 gopurs_runtime.Value
var once_Data_Number_ln2 sync.Once

func Get_Data_Number_ln2() gopurs_runtime.Value {
	once_Data_Number_ln2.Do(func() {
		cache_Data_Number_ln2 = gopurs_runtime.Float(0.6931471805599453)
	})
	return cache_Data_Number_ln2
}

var cache_Data_Number_ln10 gopurs_runtime.Value
var once_Data_Number_ln10 sync.Once

func Get_Data_Number_ln10() gopurs_runtime.Value {
	once_Data_Number_ln10.Do(func() {
		cache_Data_Number_ln10 = gopurs_runtime.Float(2.302585092994046)
	})
	return cache_Data_Number_ln10
}

var cache_Data_Number_fromString gopurs_runtime.Value
var once_Data_Number_fromString sync.Once

func Get_Data_Number_fromString() gopurs_runtime.Value {
	once_Data_Number_fromString.Do(func() {
		cache_Data_Number_fromString = gopurs_runtime.Func(func(str_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Number_fromString(str_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Number_fromString
}

var cache_Data_Number_e gopurs_runtime.Value
var once_Data_Number_e sync.Once

func Get_Data_Number_e() gopurs_runtime.Value {
	once_Data_Number_e.Do(func() {
		cache_Data_Number_e = gopurs_runtime.Float(2.718281828459045)
	})
	return cache_Data_Number_e
}

func Call_Data_Number_fromString(str_0_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var str_0 string = str_0_loop
	_ = str_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Number_3240988860_3094389156(Rebox_Data_Number_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Number_fromStringImpl(), gopurs_runtime.Str(str_0), Get_Data_Number_isFinite(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		_v := struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{gopurs_runtime.Value{}, false}
		if _v.V1 {
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
		}
		return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
	}()))})))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Rebox_Data_Number_3094389156_3240988860(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[float64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[float64]{}
	out.V0 = in.V0.FloatVal()
	return out
}

func Rebox_Data_Number_3240988860_3094389156(in *Constructor_Data_Maybe_Just[float64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Float(in.V0)
	return out
}

func Get_Data_Number_abs() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Abs
}

func Get_Data_Number_acos() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Acos
}

func Get_Data_Number_asin() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Asin
}

func Get_Data_Number_atan() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Atan
}

func Get_Data_Number_atan2() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Atan2
}

func Get_Data_Number_ceil() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Ceil
}

func Get_Data_Number_cos() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Cos
}

func Get_Data_Number_exp() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Exp
}

func Get_Data_Number_floor() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Floor
}

func Get_Data_Number_fromStringImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Number_FromStringImpl
}

func Get_Data_Number_infinity() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Infinity
}

func Get_Data_Number_isFinite() gopurs_runtime.Value {
	return _Gopurs_Data_Number_IsFinite
}

func Get_Data_Number_isNaN() gopurs_runtime.Value {
	return _Gopurs_Data_Number_IsNaN
}

func Get_Data_Number_log() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Log
}

func Get_Data_Number_max() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Max
}

func Get_Data_Number_min() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Min
}

func Get_Data_Number_nan() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Nan
}

func Get_Data_Number_pow() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Pow
}

func Get_Data_Number_remainder() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Remainder
}

func Get_Data_Number_round() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Round
}

func Get_Data_Number_sign() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Sign
}

func Get_Data_Number_sin() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Sin
}

func Get_Data_Number_sqrt() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Sqrt
}

func Get_Data_Number_tan() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Tan
}

func Get_Data_Number_trunc() gopurs_runtime.Value {
	return _Gopurs_Data_Number_Trunc
}
