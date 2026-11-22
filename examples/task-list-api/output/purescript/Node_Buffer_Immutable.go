package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Node_Buffer_Immutable_show gopurs_runtime.Value
var once_Node_Buffer_Immutable_show sync.Once

func Get_Node_Buffer_Immutable_show() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_show.Do(func() {
		cache_Node_Buffer_Immutable_show = Get_Data_Show_show__2028723346()
	})
	return cache_Node_Buffer_Immutable_show
}

var cache_Node_Buffer_Immutable_lessThan gopurs_runtime.Value
var once_Node_Buffer_Immutable_lessThan sync.Once

func Get_Node_Buffer_Immutable_lessThan() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_lessThan.Do(func() {
		cache_Node_Buffer_Immutable_lessThan = Get_Data_Ord_lessThan__3348776441()
	})
	return cache_Node_Buffer_Immutable_lessThan
}

var cache_Node_Buffer_Immutable_greaterThan gopurs_runtime.Value
var once_Node_Buffer_Immutable_greaterThan sync.Once

func Get_Node_Buffer_Immutable_greaterThan() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_greaterThan.Do(func() {
		cache_Node_Buffer_Immutable_greaterThan = Get_Data_Ord_greaterThan__3348776441()
	})
	return cache_Node_Buffer_Immutable_greaterThan
}

var cache_Node_Buffer_Immutable_mapFlipped gopurs_runtime.Value
var once_Node_Buffer_Immutable_mapFlipped sync.Once

func Get_Node_Buffer_Immutable_mapFlipped() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_mapFlipped.Do(func() {
		cache_Node_Buffer_Immutable_mapFlipped = Get_Data_Functor_mapFlipped__4096616251()
	})
	return cache_Node_Buffer_Immutable_mapFlipped
}

var cache_Node_Buffer_Immutable_go__append gopurs_runtime.Value
var once_Node_Buffer_Immutable_go__append sync.Once

func Get_Node_Buffer_Immutable_go__append() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_go__append.Do(func() {
		cache_Node_Buffer_Immutable_go__append = Get_Data_Semigroup_append__1694605992()
	})
	return cache_Node_Buffer_Immutable_go__append
}

var cache_Node_Buffer_Immutable_show1 gopurs_runtime.Value
var once_Node_Buffer_Immutable_show1 sync.Once

func Get_Node_Buffer_Immutable_show1() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_show1.Do(func() {
		cache_Node_Buffer_Immutable_show1 = Get_Data_Show_show__1986550628()
	})
	return cache_Node_Buffer_Immutable_show1
}

var cache_Node_Buffer_Immutable_toString_prime_ gopurs_runtime.Value
var once_Node_Buffer_Immutable_toString_prime_ sync.Once

func Get_Node_Buffer_Immutable_toString_prime_() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_toString_prime_.Do(func() {
		cache_Node_Buffer_Immutable_toString_prime_ = gopurs_runtime.Func4(func(enc_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value, end_2_box gopurs_runtime.Value, buf_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Node_Buffer_Immutable_toString_prime_(uint32(enc_0_box.IntVal), start_1_box.IntVal, end_2_box.IntVal, buf_3_box))
		})
	})
	return cache_Node_Buffer_Immutable_toString_prime_
}

var cache_Node_Buffer_Immutable_toString gopurs_runtime.Value
var once_Node_Buffer_Immutable_toString sync.Once

func Get_Node_Buffer_Immutable_toString() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_toString.Do(func() {
		cache_Node_Buffer_Immutable_toString = gopurs_runtime.Func2(func(enc_0_box gopurs_runtime.Value, buf_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Node_Buffer_Immutable_toString(uint32(enc_0_box.IntVal), buf_1_box))
		})
	})
	return cache_Node_Buffer_Immutable_toString
}

var cache_Node_Buffer_Immutable_slice gopurs_runtime.Value
var once_Node_Buffer_Immutable_slice sync.Once

func Get_Node_Buffer_Immutable_slice() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_slice.Do(func() {
		cache_Node_Buffer_Immutable_slice = gopurs_runtime.Func3(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value, buf_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Immutable_slice(start_0_box.IntVal, end_1_box.IntVal, buf_2_box)
		})
	})
	return cache_Node_Buffer_Immutable_slice
}

var cache_Node_Buffer_Immutable_showBuffer gopurs_runtime.Value
var once_Node_Buffer_Immutable_showBuffer sync.Once

func Get_Node_Buffer_Immutable_showBuffer() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_showBuffer.Do(func() {
		cache_Node_Buffer_Immutable_showBuffer = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, Get_Node_Buffer_Immutable_showImpl()}))}
	})
	return cache_Node_Buffer_Immutable_showBuffer
}

var cache_Node_Buffer_Immutable_readString gopurs_runtime.Value
var once_Node_Buffer_Immutable_readString sync.Once

func Get_Node_Buffer_Immutable_readString() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_readString.Do(func() {
		cache_Node_Buffer_Immutable_readString = gopurs_runtime.Func4(func(enc_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value, end_2_box gopurs_runtime.Value, buf_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Node_Buffer_Immutable_readString(uint32(enc_0_box.IntVal), start_1_box.IntVal, end_2_box.IntVal, buf_3_box))
		})
	})
	return cache_Node_Buffer_Immutable_readString
}

var cache_Node_Buffer_Immutable_read gopurs_runtime.Value
var once_Node_Buffer_Immutable_read sync.Once

func Get_Node_Buffer_Immutable_read() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_read.Do(func() {
		cache_Node_Buffer_Immutable_read = gopurs_runtime.Func3(func(ty_0_box gopurs_runtime.Value, off_1_box gopurs_runtime.Value, buf_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Float(Call_Node_Buffer_Immutable_read(uint32(ty_0_box.IntVal), off_1_box.IntVal, buf_2_box))
		})
	})
	return cache_Node_Buffer_Immutable_read
}

var cache_Node_Buffer_Immutable_getAtOffset gopurs_runtime.Value
var once_Node_Buffer_Immutable_getAtOffset sync.Once

func Get_Node_Buffer_Immutable_getAtOffset() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_getAtOffset.Do(func() {
		cache_Node_Buffer_Immutable_getAtOffset = gopurs_runtime.Func2(func(off_0_box gopurs_runtime.Value, buf_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Node_Buffer_Immutable_getAtOffset(off_0_box.IntVal, buf_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Node_Buffer_Immutable_getAtOffset
}

var cache_Node_Buffer_Immutable_fromString gopurs_runtime.Value
var once_Node_Buffer_Immutable_fromString sync.Once

func Get_Node_Buffer_Immutable_fromString() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_fromString.Do(func() {
		cache_Node_Buffer_Immutable_fromString = gopurs_runtime.Func2(func(str_0_box gopurs_runtime.Value, enc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Immutable_fromString(str_0_box.StrVal(), uint32(enc_1_box.IntVal))
		})
	})
	return cache_Node_Buffer_Immutable_fromString
}

var cache_Node_Buffer_Immutable_eqBuffer gopurs_runtime.Value
var once_Node_Buffer_Immutable_eqBuffer sync.Once

func Get_Node_Buffer_Immutable_eqBuffer() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_eqBuffer.Do(func() {
		cache_Node_Buffer_Immutable_eqBuffer = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.UncurriedApp2(Get_Node_Buffer_Immutable_eqImpl(), a_0, b_1)
		})}))}
	})
	return cache_Node_Buffer_Immutable_eqBuffer
}

var cache_Node_Buffer_Immutable_ordBuffer gopurs_runtime.Value
var once_Node_Buffer_Immutable_ordBuffer sync.Once

func Get_Node_Buffer_Immutable_ordBuffer() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_ordBuffer.Do(func() {
		cache_Node_Buffer_Immutable_ordBuffer = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Node_Buffer_Immutable_eqBuffer()))}
		}), gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): v_2_0 shape=UncurriedApp(Var) bindingType=Int
			v_2_0 := gopurs_runtime.UncurriedApp2(Get_Node_Buffer_Immutable_compareImpl(), a_0, b_1).IntVal
			_ = v_2_0
			var __t1 gopurs_runtime.Value
			{
				if (v_2_0) < (int64(0)) {
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					goto end_branch_1
				} else {

				}
			}
			{
				if (v_2_0) > (int64(0)) {
					__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_1
				} else {

				}
			}
			{
				__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
			}
		end_branch_1:
			return __t1
		})}))}
	})
	return cache_Node_Buffer_Immutable_ordBuffer
}

var cache_Node_Buffer_Immutable_create gopurs_runtime.Value
var once_Node_Buffer_Immutable_create sync.Once

func Get_Node_Buffer_Immutable_create() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_create.Do(func() {
		cache_Node_Buffer_Immutable_create = Get_Node_Buffer_Immutable_alloc()
	})
	return cache_Node_Buffer_Immutable_create
}

var cache_Node_Buffer_Immutable_concat_prime_ gopurs_runtime.Value
var once_Node_Buffer_Immutable_concat_prime_ sync.Once

func Get_Node_Buffer_Immutable_concat_prime_() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_concat_prime_.Do(func() {
		cache_Node_Buffer_Immutable_concat_prime_ = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, i_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Immutable_concat_prime_(*((*[]gopurs_runtime.Value)(a_0_box.UnsafePtr)), i_1_box.IntVal)
		})
	})
	return cache_Node_Buffer_Immutable_concat_prime_
}

var cache_Node_Buffer_Immutable_compareParts gopurs_runtime.Value
var once_Node_Buffer_Immutable_compareParts sync.Once

func Get_Node_Buffer_Immutable_compareParts() gopurs_runtime.Value {
	once_Node_Buffer_Immutable_compareParts.Do(func() {
		cache_Node_Buffer_Immutable_compareParts = gopurs_runtime.Func6(func(src_0_box gopurs_runtime.Value, target_1_box gopurs_runtime.Value, targetStart_2_box gopurs_runtime.Value, targetEnd_3_box gopurs_runtime.Value, sourceStart_4_box gopurs_runtime.Value, sourceEnd_5_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_Immutable_compareParts(src_0_box, target_1_box, targetStart_2_box.IntVal, targetEnd_3_box.IntVal, sourceStart_4_box.IntVal, sourceEnd_5_box.IntVal)
		})
	})
	return cache_Node_Buffer_Immutable_compareParts
}

func Call_Node_Buffer_Immutable_toString_prime_(enc_0_loop uint32, start_1_loop int64, end_2_loop int64, buf_3_loop gopurs_runtime.Value) string {
	var enc_0 uint32 = enc_0_loop
	_ = enc_0
	var start_1 int64 = start_1_loop
	_ = start_1
	var end_2 int64 = end_2_loop
	_ = end_2
	var buf_3 gopurs_runtime.Value = buf_3_loop
	_ = buf_3
	return gopurs_runtime.UncurriedApp4(Get_Node_Buffer_Immutable_toStringSubImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(enc_0), UnsafePtr: nil}, gopurs_runtime.Int(start_1), gopurs_runtime.Int(end_2), buf_3).StrVal()
}

func Call_Node_Buffer_Immutable_toString(enc_0_loop uint32, buf_1_loop gopurs_runtime.Value) string {
	var enc_0 uint32 = enc_0_loop
	_ = enc_0
	var buf_1 gopurs_runtime.Value = buf_1_loop
	_ = buf_1
	var __t0 string
	{
		if enc_0 == 2598402200 {
			__t0 = "ascii"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 346813718 {
			__t0 = "utf8"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 2352678688 {
			__t0 = "utf16le"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 4287456318 {
			__t0 = "ucs2"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 2405785054 {
			__t0 = "base64"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 2476550165 {
			__t0 = "base64url"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 3576003686 {
			__t0 = "latin1"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 349879974 {
			__t0 = "binary"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 106823740 {
			__t0 = "hex"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() string { panic("Failed pattern match") }()
	}
end_branch_0:
	return gopurs_runtime.UncurriedApp2(Get_Node_Buffer_Immutable_toStringImpl(), gopurs_runtime.Str(__t0), buf_1).StrVal()
}

func Call_Node_Buffer_Immutable_slice(start_0_loop int64, end_1_loop int64, buf_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var start_0 int64 = start_0_loop
	_ = start_0
	var end_1 int64 = end_1_loop
	_ = end_1
	var buf_2 gopurs_runtime.Value = buf_2_loop
	_ = buf_2
	return gopurs_runtime.UncurriedApp3(Get_Node_Buffer_Immutable_sliceImpl(), gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1), buf_2)
}

func Call_Node_Buffer_Immutable_readString(enc_0_loop uint32, start_1_loop int64, end_2_loop int64, buf_3_loop gopurs_runtime.Value) string {
	var enc_0 uint32 = enc_0_loop
	_ = enc_0
	var start_1 int64 = start_1_loop
	_ = start_1
	var end_2 int64 = end_2_loop
	_ = end_2
	var buf_3 gopurs_runtime.Value = buf_3_loop
	_ = buf_3
	var __t0 string
	{
		if enc_0 == 2598402200 {
			__t0 = "ascii"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 346813718 {
			__t0 = "utf8"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 2352678688 {
			__t0 = "utf16le"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 4287456318 {
			__t0 = "ucs2"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 2405785054 {
			__t0 = "base64"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 2476550165 {
			__t0 = "base64url"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 3576003686 {
			__t0 = "latin1"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 349879974 {
			__t0 = "binary"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_0 == 106823740 {
			__t0 = "hex"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() string { panic("Failed pattern match") }()
	}
end_branch_0:
	return gopurs_runtime.UncurriedApp4(Get_Node_Buffer_Immutable_readStringImpl(), gopurs_runtime.Str(__t0), gopurs_runtime.Int(start_1), gopurs_runtime.Int(end_2), buf_3).StrVal()
}

func Call_Node_Buffer_Immutable_read(ty_0_loop uint32, off_1_loop int64, buf_2_loop gopurs_runtime.Value) float64 {
	var ty_0 uint32 = ty_0_loop
	_ = ty_0
	var off_1 int64 = off_1_loop
	_ = off_1
	var buf_2 gopurs_runtime.Value = buf_2_loop
	_ = buf_2
	var __t0 string
	{
		if ty_0 == 3314799936 {
			__t0 = "UInt8"
			goto end_branch_0
		} else {

		}
	}
	{
		if ty_0 == 852326966 {
			__t0 = "UInt16LE"
			goto end_branch_0
		} else {

		}
	}
	{
		if ty_0 == 305778936 {
			__t0 = "UInt16BE"
			goto end_branch_0
		} else {

		}
	}
	{
		if ty_0 == 3460170864 {
			__t0 = "UInt32LE"
			goto end_branch_0
		} else {

		}
	}
	{
		if ty_0 == 1992902334 {
			__t0 = "UInt32BE"
			goto end_branch_0
		} else {

		}
	}
	{
		if ty_0 == 1095082613 {
			__t0 = "Int8"
			goto end_branch_0
		} else {

		}
	}
	{
		if ty_0 == 408746499 {
			__t0 = "Int16LE"
			goto end_branch_0
		} else {

		}
	}
	{
		if ty_0 == 3066827405 {
			__t0 = "Int16BE"
			goto end_branch_0
		} else {

		}
	}
	{
		if ty_0 == 199750917 {
			__t0 = "Int32LE"
			goto end_branch_0
		} else {

		}
	}
	{
		if ty_0 == 3275822987 {
			__t0 = "Int32BE"
			goto end_branch_0
		} else {

		}
	}
	{
		if ty_0 == 1492841799 {
			__t0 = "FloatLE"
			goto end_branch_0
		} else {

		}
	}
	{
		if ty_0 == 4210130121 {
			__t0 = "FloatBE"
			goto end_branch_0
		} else {

		}
	}
	{
		if ty_0 == 496370018 {
			__t0 = "DoubleLE"
			goto end_branch_0
		} else {

		}
	}
	{
		if ty_0 == 894396972 {
			__t0 = "DoubleBE"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() string { panic("Failed pattern match") }()
	}
end_branch_0:
	return gopurs_runtime.UncurriedApp3(Get_Node_Buffer_Immutable_readImpl(), gopurs_runtime.Str(__t0), gopurs_runtime.Int(off_1), buf_2).FloatVal()
}

func Call_Node_Buffer_Immutable_getAtOffset(off_0_loop int64, buf_1_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var off_0 int64 = off_0_loop
	_ = off_0
	var buf_1 gopurs_runtime.Value = buf_1_loop
	_ = buf_1
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Node_Buffer_Immutable_1170268447_3094389156(Rebox_Node_Buffer_Immutable_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), gopurs_runtime.UncurriedApp2(Get_Node_Buffer_Immutable_getAtOffsetImpl(), gopurs_runtime.Int(off_0), buf_1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Node_Buffer_Immutable_1170268447_3094389156(Rebox_Node_Buffer_Immutable_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))}, Get_Data_Maybe_Just())))))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Call_Node_Buffer_Immutable_fromString(str_0_loop string, enc_1_loop uint32) gopurs_runtime.Value {
	var str_0 string = str_0_loop
	_ = str_0
	var enc_1 uint32 = enc_1_loop
	_ = enc_1
	var __t0 string
	{
		if enc_1 == 2598402200 {
			__t0 = "ascii"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 346813718 {
			__t0 = "utf8"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 2352678688 {
			__t0 = "utf16le"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 4287456318 {
			__t0 = "ucs2"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 2405785054 {
			__t0 = "base64"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 2476550165 {
			__t0 = "base64url"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 3576003686 {
			__t0 = "latin1"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 349879974 {
			__t0 = "binary"
			goto end_branch_0
		} else {

		}
	}
	{
		if enc_1 == 106823740 {
			__t0 = "hex"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() string { panic("Failed pattern match") }()
	}
end_branch_0:
	return gopurs_runtime.UncurriedApp2(Get_Node_Buffer_Immutable_fromStringImpl(), gopurs_runtime.Str(str_0), gopurs_runtime.Str(__t0))
}

func Call_Node_Buffer_Immutable_concat_prime_(a_0_loop []gopurs_runtime.Value, i_1_loop int64) gopurs_runtime.Value {
	var a_0 []gopurs_runtime.Value = a_0_loop
	_ = a_0
	var i_1 int64 = i_1_loop
	_ = i_1
	return gopurs_runtime.UncurriedApp2(Get_Node_Buffer_Immutable_concatToLength(), gopurs_runtime.Array(a_0), gopurs_runtime.Int(i_1))
}

func Call_Node_Buffer_Immutable_compareParts(src_0_loop gopurs_runtime.Value, target_1_loop gopurs_runtime.Value, targetStart_2_loop int64, targetEnd_3_loop int64, sourceStart_4_loop int64, sourceEnd_5_loop int64) gopurs_runtime.Value {
	var src_0 gopurs_runtime.Value = src_0_loop
	_ = src_0
	var target_1 gopurs_runtime.Value = target_1_loop
	_ = target_1
	var targetStart_2 int64 = targetStart_2_loop
	_ = targetStart_2
	var targetEnd_3 int64 = targetEnd_3_loop
	_ = targetEnd_3
	var sourceStart_4 int64 = sourceStart_4_loop
	_ = sourceStart_4
	var sourceEnd_5 int64 = sourceEnd_5_loop
	_ = sourceEnd_5
	return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t0 gopurs_runtime.Value
		{
			if (v_6.IntVal) == (int64(-1)) {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
				goto end_branch_0
			} else {

			}
		}
		{
			if (v_6.IntVal) == (int64(0)) {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
				goto end_branch_0
			} else {

			}
		}
		{
			if (v_6.IntVal) == (int64(1)) {
				__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = Call_Partial_Unsafe_unsafeCrashWith(gopurs_runtime.ConcatString("Impossible: Invalid value: ", Data_Show_ShowIntImpl(v_6.IntVal)))
		}
	end_branch_0:
		return __t0
	}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Buffer_Immutable_comparePartsImpl(), src_0, target_1, gopurs_runtime.Int(targetStart_2), gopurs_runtime.Int(targetEnd_3), gopurs_runtime.Int(sourceStart_4), gopurs_runtime.Int(sourceEnd_5))
	}))
}

func Rebox_Node_Buffer_Immutable_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Node_Buffer_Immutable_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[int64]{}
	out.V0 = in.V0.IntVal
	return out
}

func Get_Node_Buffer_Immutable_alloc() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_Alloc
}

func Get_Node_Buffer_Immutable_compareImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_CompareImpl
}

func Get_Node_Buffer_Immutable_comparePartsImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_ComparePartsImpl
}

func Get_Node_Buffer_Immutable_concat() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_Concat
}

func Get_Node_Buffer_Immutable_concatToLength() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_ConcatToLength
}

func Get_Node_Buffer_Immutable_eqImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_EqImpl
}

func Get_Node_Buffer_Immutable_fromArray() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_FromArray
}

func Get_Node_Buffer_Immutable_fromArrayBuffer() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_FromArrayBuffer
}

func Get_Node_Buffer_Immutable_fromStringImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_FromStringImpl
}

func Get_Node_Buffer_Immutable_getAtOffsetImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_GetAtOffsetImpl
}

func Get_Node_Buffer_Immutable_readImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_ReadImpl
}

func Get_Node_Buffer_Immutable_readStringImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_ReadStringImpl
}

func Get_Node_Buffer_Immutable_showImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_ShowImpl
}

func Get_Node_Buffer_Immutable_size() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_Size
}

func Get_Node_Buffer_Immutable_sliceImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_SliceImpl
}

func Get_Node_Buffer_Immutable_toArray() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_ToArray
}

func Get_Node_Buffer_Immutable_toArrayBuffer() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_ToArrayBuffer
}

func Get_Node_Buffer_Immutable_toStringImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_ToStringImpl
}

func Get_Node_Buffer_Immutable_toStringSubImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Immutable_ToStringSubImpl
}
