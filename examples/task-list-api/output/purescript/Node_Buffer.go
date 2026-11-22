package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Node_Buffer_show gopurs_runtime.Value
var once_Node_Buffer_show sync.Once

func Get_Node_Buffer_show() gopurs_runtime.Value {
	once_Node_Buffer_show.Do(func() {
		cache_Node_Buffer_show = Get_Data_Show_show__2028723346()
	})
	return cache_Node_Buffer_show
}

var cache_Node_Buffer_go__map gopurs_runtime.Value
var once_Node_Buffer_go__map sync.Once

func Get_Node_Buffer_go__map() gopurs_runtime.Value {
	once_Node_Buffer_go__map.Do(func() {
		cache_Node_Buffer_go__map = gopurs_runtime.Apply(Get_Control_Applicative_liftA1(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()))})
	})
	return cache_Node_Buffer_go__map
}

var cache_Node_Buffer_map__1922202664 gopurs_runtime.Value
var once_Node_Buffer_map__1922202664 sync.Once

func Get_Node_Buffer_map__1922202664() gopurs_runtime.Value {
	once_Node_Buffer_map__1922202664.Do(func() {
		cache_Node_Buffer_map__1922202664 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_map__1922202664(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Node_Buffer_map__1922202664
}

var cache_Node_Buffer_map__488316713 gopurs_runtime.Value
var once_Node_Buffer_map__488316713 sync.Once

func Get_Node_Buffer_map__488316713() gopurs_runtime.Value {
	once_Node_Buffer_map__488316713.Do(func() {
		cache_Node_Buffer_map__488316713 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_map__488316713(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Node_Buffer_map__488316713
}

var cache_Node_Buffer_map__1198218120 gopurs_runtime.Value
var once_Node_Buffer_map__1198218120 sync.Once

func Get_Node_Buffer_map__1198218120() gopurs_runtime.Value {
	once_Node_Buffer_map__1198218120.Do(func() {
		cache_Node_Buffer_map__1198218120 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_map__1198218120(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Node_Buffer_map__1198218120
}

var cache_Node_Buffer_map__3336206575 gopurs_runtime.Value
var once_Node_Buffer_map__3336206575 sync.Once

func Get_Node_Buffer_map__3336206575() gopurs_runtime.Value {
	once_Node_Buffer_map__3336206575.Do(func() {
		cache_Node_Buffer_map__3336206575 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_map__3336206575(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Node_Buffer_map__3336206575
}

var cache_Node_Buffer_map__2503807369 gopurs_runtime.Value
var once_Node_Buffer_map__2503807369 sync.Once

func Get_Node_Buffer_map__2503807369() gopurs_runtime.Value {
	once_Node_Buffer_map__2503807369.Do(func() {
		cache_Node_Buffer_map__2503807369 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_map__2503807369(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Node_Buffer_map__2503807369
}

var cache_Node_Buffer_map__322707977 gopurs_runtime.Value
var once_Node_Buffer_map__322707977 sync.Once

func Get_Node_Buffer_map__322707977() gopurs_runtime.Value {
	once_Node_Buffer_map__322707977.Do(func() {
		cache_Node_Buffer_map__322707977 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_map__322707977(__eta_norm_1_0_box, __eta_norm_0_1_box)
		})
	})
	return cache_Node_Buffer_map__322707977
}

var cache_Node_Buffer_bind gopurs_runtime.Value
var once_Node_Buffer_bind sync.Once

func Get_Node_Buffer_bind() gopurs_runtime.Value {
	once_Node_Buffer_bind.Do(func() {
		cache_Node_Buffer_bind = Get_Control_Bind_bind__3513850932()
	})
	return cache_Node_Buffer_bind
}

var cache_Node_Buffer_writeString gopurs_runtime.Value
var once_Node_Buffer_writeString sync.Once

func Get_Node_Buffer_writeString() gopurs_runtime.Value {
	once_Node_Buffer_writeString.Do(func() {
		cache_Node_Buffer_writeString = gopurs_runtime.Func5(func(enc_0_box gopurs_runtime.Value, offset_1_box gopurs_runtime.Value, len_2_box gopurs_runtime.Value, value_3_box gopurs_runtime.Value, buf_4_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_writeString(uint32(enc_0_box.IntVal), offset_1_box.IntVal, len_2_box.IntVal, value_3_box.StrVal(), buf_4_box)
		})
	})
	return cache_Node_Buffer_writeString
}

var cache_Node_Buffer_write gopurs_runtime.Value
var once_Node_Buffer_write sync.Once

func Get_Node_Buffer_write() gopurs_runtime.Value {
	once_Node_Buffer_write.Do(func() {
		cache_Node_Buffer_write = gopurs_runtime.Func4(func(ty_0_box gopurs_runtime.Value, value_1_box gopurs_runtime.Value, offset_2_box gopurs_runtime.Value, buf_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_write(uint32(ty_0_box.IntVal), value_1_box.FloatVal(), offset_2_box.IntVal, buf_3_box)
		})
	})
	return cache_Node_Buffer_write
}

var cache_Node_Buffer_unsafeThaw gopurs_runtime.Value
var once_Node_Buffer_unsafeThaw sync.Once

func Get_Node_Buffer_unsafeThaw() gopurs_runtime.Value {
	once_Node_Buffer_unsafeThaw.Do(func() {
		cache_Node_Buffer_unsafeThaw = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_pureE(), Get_Unsafe_Coerce_unsafeCoerce())
	})
	return cache_Node_Buffer_unsafeThaw
}

var cache_Node_Buffer_usingToImmutable gopurs_runtime.Value
var once_Node_Buffer_usingToImmutable sync.Once

func Get_Node_Buffer_usingToImmutable() gopurs_runtime.Value {
	once_Node_Buffer_usingToImmutable.Do(func() {
		cache_Node_Buffer_usingToImmutable = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_usingToImmutable(f_0_box, x_1_box)
		})
	})
	return cache_Node_Buffer_usingToImmutable
}

var cache_Node_Buffer_usingToImmutable__2508462662 gopurs_runtime.Value
var once_Node_Buffer_usingToImmutable__2508462662 sync.Once

func Get_Node_Buffer_usingToImmutable__2508462662() gopurs_runtime.Value {
	once_Node_Buffer_usingToImmutable__2508462662.Do(func() {
		cache_Node_Buffer_usingToImmutable__2508462662 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_usingToImmutable__2508462662(f_unused_0_box, x_1_box)
		})
	})
	return cache_Node_Buffer_usingToImmutable__2508462662
}

var cache_Node_Buffer_usingToImmutable__1694059786 gopurs_runtime.Value
var once_Node_Buffer_usingToImmutable__1694059786 sync.Once

func Get_Node_Buffer_usingToImmutable__1694059786() gopurs_runtime.Value {
	once_Node_Buffer_usingToImmutable__1694059786.Do(func() {
		cache_Node_Buffer_usingToImmutable__1694059786 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_usingToImmutable__1694059786(f_0_box, uint32(x_1_box.IntVal))
		})
	})
	return cache_Node_Buffer_usingToImmutable__1694059786
}

var cache_Node_Buffer_usingToImmutable__3986645350 gopurs_runtime.Value
var once_Node_Buffer_usingToImmutable__3986645350 sync.Once

func Get_Node_Buffer_usingToImmutable__3986645350() gopurs_runtime.Value {
	once_Node_Buffer_usingToImmutable__3986645350.Do(func() {
		cache_Node_Buffer_usingToImmutable__3986645350 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_usingToImmutable__3986645350(f_unused_0_box, func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}())
		})
	})
	return cache_Node_Buffer_usingToImmutable__3986645350
}

var cache_Node_Buffer_usingToImmutable__4058095588 gopurs_runtime.Value
var once_Node_Buffer_usingToImmutable__4058095588 sync.Once

func Get_Node_Buffer_usingToImmutable__4058095588() gopurs_runtime.Value {
	once_Node_Buffer_usingToImmutable__4058095588.Do(func() {
		cache_Node_Buffer_usingToImmutable__4058095588 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_usingToImmutable__4058095588(f_unused_0_box, x_1_box.IntVal)
		})
	})
	return cache_Node_Buffer_usingToImmutable__4058095588
}

var cache_Node_Buffer_unsafeFreeze gopurs_runtime.Value
var once_Node_Buffer_unsafeFreeze sync.Once

func Get_Node_Buffer_unsafeFreeze() gopurs_runtime.Value {
	once_Node_Buffer_unsafeFreeze.Do(func() {
		cache_Node_Buffer_unsafeFreeze = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_pureE(), Get_Unsafe_Coerce_unsafeCoerce())
	})
	return cache_Node_Buffer_unsafeFreeze
}

var cache_Node_Buffer_usingFromImmutable gopurs_runtime.Value
var once_Node_Buffer_usingFromImmutable sync.Once

func Get_Node_Buffer_usingFromImmutable() gopurs_runtime.Value {
	once_Node_Buffer_usingFromImmutable.Do(func() {
		cache_Node_Buffer_usingFromImmutable = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, buf_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_usingFromImmutable(f_0_box, buf_1_box)
		})
	})
	return cache_Node_Buffer_usingFromImmutable
}

var cache_Node_Buffer_usingFromImmutable__3422936011 gopurs_runtime.Value
var once_Node_Buffer_usingFromImmutable__3422936011 sync.Once

func Get_Node_Buffer_usingFromImmutable__3422936011() gopurs_runtime.Value {
	once_Node_Buffer_usingFromImmutable__3422936011.Do(func() {
		cache_Node_Buffer_usingFromImmutable__3422936011 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, buf_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_usingFromImmutable__3422936011(f_unused_0_box, buf_1_box)
		})
	})
	return cache_Node_Buffer_usingFromImmutable__3422936011
}

var cache_Node_Buffer_usingFromImmutable__2525970890 gopurs_runtime.Value
var once_Node_Buffer_usingFromImmutable__2525970890 sync.Once

func Get_Node_Buffer_usingFromImmutable__2525970890() gopurs_runtime.Value {
	once_Node_Buffer_usingFromImmutable__2525970890.Do(func() {
		cache_Node_Buffer_usingFromImmutable__2525970890 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, buf_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_usingFromImmutable__2525970890(f_0_box, buf_1_box)
		})
	})
	return cache_Node_Buffer_usingFromImmutable__2525970890
}

var cache_Node_Buffer_usingFromImmutable__3335896427 gopurs_runtime.Value
var once_Node_Buffer_usingFromImmutable__3335896427 sync.Once

func Get_Node_Buffer_usingFromImmutable__3335896427() gopurs_runtime.Value {
	once_Node_Buffer_usingFromImmutable__3335896427.Do(func() {
		cache_Node_Buffer_usingFromImmutable__3335896427 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, buf_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_usingFromImmutable__3335896427(f_unused_0_box, buf_1_box)
		})
	})
	return cache_Node_Buffer_usingFromImmutable__3335896427
}

var cache_Node_Buffer_usingFromImmutable__2070034188 gopurs_runtime.Value
var once_Node_Buffer_usingFromImmutable__2070034188 sync.Once

func Get_Node_Buffer_usingFromImmutable__2070034188() gopurs_runtime.Value {
	once_Node_Buffer_usingFromImmutable__2070034188.Do(func() {
		cache_Node_Buffer_usingFromImmutable__2070034188 = gopurs_runtime.Func2(func(f_unused_0_box gopurs_runtime.Value, buf_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_usingFromImmutable__2070034188(f_unused_0_box, buf_1_box)
		})
	})
	return cache_Node_Buffer_usingFromImmutable__2070034188
}

var cache_Node_Buffer_usingFromImmutable__2892494410 gopurs_runtime.Value
var once_Node_Buffer_usingFromImmutable__2892494410 sync.Once

func Get_Node_Buffer_usingFromImmutable__2892494410() gopurs_runtime.Value {
	once_Node_Buffer_usingFromImmutable__2892494410.Do(func() {
		cache_Node_Buffer_usingFromImmutable__2892494410 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, buf_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_usingFromImmutable__2892494410(f_0_box, buf_1_box)
		})
	})
	return cache_Node_Buffer_usingFromImmutable__2892494410
}

var cache_Node_Buffer_usingFromImmutable__3537276362 gopurs_runtime.Value
var once_Node_Buffer_usingFromImmutable__3537276362 sync.Once

func Get_Node_Buffer_usingFromImmutable__3537276362() gopurs_runtime.Value {
	once_Node_Buffer_usingFromImmutable__3537276362.Do(func() {
		cache_Node_Buffer_usingFromImmutable__3537276362 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, buf_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_usingFromImmutable__3537276362(f_0_box, buf_1_box)
		})
	})
	return cache_Node_Buffer_usingFromImmutable__3537276362
}

var cache_Node_Buffer_transcode gopurs_runtime.Value
var once_Node_Buffer_transcode sync.Once

func Get_Node_Buffer_transcode() gopurs_runtime.Value {
	once_Node_Buffer_transcode.Do(func() {
		cache_Node_Buffer_transcode = gopurs_runtime.Func3(func(buf_0_box gopurs_runtime.Value, from_1_box gopurs_runtime.Value, to_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_transcode(buf_0_box, uint32(from_1_box.IntVal), uint32(to_2_box.IntVal))
		})
	})
	return cache_Node_Buffer_transcode
}

var cache_Node_Buffer_toString_prime_ gopurs_runtime.Value
var once_Node_Buffer_toString_prime_ sync.Once

func Get_Node_Buffer_toString_prime_() gopurs_runtime.Value {
	once_Node_Buffer_toString_prime_.Do(func() {
		cache_Node_Buffer_toString_prime_ = gopurs_runtime.Func3(func(enc_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value, end_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_toString_prime_(uint32(enc_0_box.IntVal), start_1_box.IntVal, end_2_box.IntVal)
		})
	})
	return cache_Node_Buffer_toString_prime_
}

var cache_Node_Buffer_toString gopurs_runtime.Value
var once_Node_Buffer_toString sync.Once

func Get_Node_Buffer_toString() gopurs_runtime.Value {
	once_Node_Buffer_toString.Do(func() {
		cache_Node_Buffer_toString = gopurs_runtime.Func(func(enc_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_toString(uint32(enc_0_box.IntVal))
		})
	})
	return cache_Node_Buffer_toString
}

var cache_Node_Buffer_toArrayBuffer gopurs_runtime.Value
var once_Node_Buffer_toArrayBuffer sync.Once

func Get_Node_Buffer_toArrayBuffer() gopurs_runtime.Value {
	once_Node_Buffer_toArrayBuffer.Do(func() {
		cache_Node_Buffer_toArrayBuffer = gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable__3422936011(), Get_Node_Buffer_Immutable_toArrayBuffer())
	})
	return cache_Node_Buffer_toArrayBuffer
}

var cache_Node_Buffer_toArray gopurs_runtime.Value
var once_Node_Buffer_toArray sync.Once

func Get_Node_Buffer_toArray() gopurs_runtime.Value {
	once_Node_Buffer_toArray.Do(func() {
		cache_Node_Buffer_toArray = gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable__3335896427(), Get_Node_Buffer_Immutable_toArray())
	})
	return cache_Node_Buffer_toArray
}

var cache_Node_Buffer_thaw gopurs_runtime.Value
var once_Node_Buffer_thaw sync.Once

func Get_Node_Buffer_thaw() gopurs_runtime.Value {
	once_Node_Buffer_thaw.Do(func() {
		cache_Node_Buffer_thaw = gopurs_runtime.Apply(Get_Effect_Uncurried_runEffectFn1(), Get_Node_Buffer_thawImpl())
	})
	return cache_Node_Buffer_thaw
}

var cache_Node_Buffer_swap64 gopurs_runtime.Value
var once_Node_Buffer_swap64 sync.Once

func Get_Node_Buffer_swap64() gopurs_runtime.Value {
	once_Node_Buffer_swap64.Do(func() {
		cache_Node_Buffer_swap64 = gopurs_runtime.Func(func(b_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_swap64(b_0_box)
		})
	})
	return cache_Node_Buffer_swap64
}

var cache_Node_Buffer_swap32 gopurs_runtime.Value
var once_Node_Buffer_swap32 sync.Once

func Get_Node_Buffer_swap32() gopurs_runtime.Value {
	once_Node_Buffer_swap32.Do(func() {
		cache_Node_Buffer_swap32 = gopurs_runtime.Func(func(b_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_swap32(b_0_box)
		})
	})
	return cache_Node_Buffer_swap32
}

var cache_Node_Buffer_swap16 gopurs_runtime.Value
var once_Node_Buffer_swap16 sync.Once

func Get_Node_Buffer_swap16() gopurs_runtime.Value {
	once_Node_Buffer_swap16.Do(func() {
		cache_Node_Buffer_swap16 = gopurs_runtime.Func(func(b_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_swap16(b_0_box)
		})
	})
	return cache_Node_Buffer_swap16
}

var cache_Node_Buffer_slice gopurs_runtime.Value
var once_Node_Buffer_slice sync.Once

func Get_Node_Buffer_slice() gopurs_runtime.Value {
	once_Node_Buffer_slice.Do(func() {
		cache_Node_Buffer_slice = Get_Node_Buffer_Immutable_slice()
	})
	return cache_Node_Buffer_slice
}

var cache_Node_Buffer_size gopurs_runtime.Value
var once_Node_Buffer_size sync.Once

func Get_Node_Buffer_size() gopurs_runtime.Value {
	once_Node_Buffer_size.Do(func() {
		cache_Node_Buffer_size = gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable__2070034188(), Get_Node_Buffer_Immutable_size())
	})
	return cache_Node_Buffer_size
}

var cache_Node_Buffer_setPoolSize gopurs_runtime.Value
var once_Node_Buffer_setPoolSize sync.Once

func Get_Node_Buffer_setPoolSize() gopurs_runtime.Value {
	once_Node_Buffer_setPoolSize.Do(func() {
		cache_Node_Buffer_setPoolSize = gopurs_runtime.Func(func(sizeInBytes_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_setPoolSize(sizeInBytes_0_box.IntVal)
		})
	})
	return cache_Node_Buffer_setPoolSize
}

var cache_Node_Buffer_setAtOffset gopurs_runtime.Value
var once_Node_Buffer_setAtOffset sync.Once

func Get_Node_Buffer_setAtOffset() gopurs_runtime.Value {
	once_Node_Buffer_setAtOffset.Do(func() {
		cache_Node_Buffer_setAtOffset = gopurs_runtime.Func3(func(val_0_box gopurs_runtime.Value, off_1_box gopurs_runtime.Value, buff_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_setAtOffset(val_0_box.IntVal, off_1_box.IntVal, buff_2_box)
		})
	})
	return cache_Node_Buffer_setAtOffset
}

var cache_Node_Buffer_readString gopurs_runtime.Value
var once_Node_Buffer_readString sync.Once

func Get_Node_Buffer_readString() gopurs_runtime.Value {
	once_Node_Buffer_readString.Do(func() {
		cache_Node_Buffer_readString = gopurs_runtime.Func3(func(enc_0_box gopurs_runtime.Value, o_1_box gopurs_runtime.Value, o_prime__2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_readString(uint32(enc_0_box.IntVal), o_1_box.IntVal, o_prime__2_box.IntVal)
		})
	})
	return cache_Node_Buffer_readString
}

var cache_Node_Buffer_read gopurs_runtime.Value
var once_Node_Buffer_read sync.Once

func Get_Node_Buffer_read() gopurs_runtime.Value {
	once_Node_Buffer_read.Do(func() {
		cache_Node_Buffer_read = gopurs_runtime.Func2(func(t_0_box gopurs_runtime.Value, o_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_read(uint32(t_0_box.IntVal), o_1_box.IntVal)
		})
	})
	return cache_Node_Buffer_read
}

var cache_Node_Buffer_getAtOffset gopurs_runtime.Value
var once_Node_Buffer_getAtOffset sync.Once

func Get_Node_Buffer_getAtOffset() gopurs_runtime.Value {
	once_Node_Buffer_getAtOffset.Do(func() {
		cache_Node_Buffer_getAtOffset = gopurs_runtime.Func(func(o_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_getAtOffset(o_0_box.IntVal)
		})
	})
	return cache_Node_Buffer_getAtOffset
}

var cache_Node_Buffer_fromString gopurs_runtime.Value
var once_Node_Buffer_fromString sync.Once

func Get_Node_Buffer_fromString() gopurs_runtime.Value {
	once_Node_Buffer_fromString.Do(func() {
		cache_Node_Buffer_fromString = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_fromString(s_0_box.StrVal())
		})
	})
	return cache_Node_Buffer_fromString
}

var cache_Node_Buffer_fromArrayBuffer gopurs_runtime.Value
var once_Node_Buffer_fromArrayBuffer sync.Once

func Get_Node_Buffer_fromArrayBuffer() gopurs_runtime.Value {
	once_Node_Buffer_fromArrayBuffer.Do(func() {
		cache_Node_Buffer_fromArrayBuffer = gopurs_runtime.Apply(Get_Node_Buffer_usingToImmutable__2508462662(), Get_Node_Buffer_Immutable_fromArrayBuffer())
	})
	return cache_Node_Buffer_fromArrayBuffer
}

var cache_Node_Buffer_fromArray gopurs_runtime.Value
var once_Node_Buffer_fromArray sync.Once

func Get_Node_Buffer_fromArray() gopurs_runtime.Value {
	once_Node_Buffer_fromArray.Do(func() {
		cache_Node_Buffer_fromArray = gopurs_runtime.Apply(Get_Node_Buffer_usingToImmutable__3986645350(), Get_Node_Buffer_Immutable_fromArray())
	})
	return cache_Node_Buffer_fromArray
}

var cache_Node_Buffer_freeze gopurs_runtime.Value
var once_Node_Buffer_freeze sync.Once

func Get_Node_Buffer_freeze() gopurs_runtime.Value {
	once_Node_Buffer_freeze.Do(func() {
		cache_Node_Buffer_freeze = gopurs_runtime.Apply(Get_Effect_Uncurried_runEffectFn1(), Get_Node_Buffer_freezeImpl())
	})
	return cache_Node_Buffer_freeze
}

var cache_Node_Buffer_fill gopurs_runtime.Value
var once_Node_Buffer_fill sync.Once

func Get_Node_Buffer_fill() gopurs_runtime.Value {
	once_Node_Buffer_fill.Do(func() {
		cache_Node_Buffer_fill = gopurs_runtime.Func4(func(octet_0_box gopurs_runtime.Value, start_1_box gopurs_runtime.Value, end_2_box gopurs_runtime.Value, buf_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_fill(octet_0_box.IntVal, start_1_box.IntVal, end_2_box.IntVal, buf_3_box)
		})
	})
	return cache_Node_Buffer_fill
}

var cache_Node_Buffer_copy gopurs_runtime.Value
var once_Node_Buffer_copy sync.Once

func Get_Node_Buffer_copy() gopurs_runtime.Value {
	once_Node_Buffer_copy.Do(func() {
		cache_Node_Buffer_copy = gopurs_runtime.Func5(func(srcStart_0_box gopurs_runtime.Value, srcEnd_1_box gopurs_runtime.Value, src_2_box gopurs_runtime.Value, targStart_3_box gopurs_runtime.Value, targ_4_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_copy(srcStart_0_box.IntVal, srcEnd_1_box.IntVal, src_2_box, targStart_3_box.IntVal, targ_4_box)
		})
	})
	return cache_Node_Buffer_copy
}

var cache_Node_Buffer_concat_prime_ gopurs_runtime.Value
var once_Node_Buffer_concat_prime_ sync.Once

func Get_Node_Buffer_concat_prime_() gopurs_runtime.Value {
	once_Node_Buffer_concat_prime_.Do(func() {
		cache_Node_Buffer_concat_prime_ = gopurs_runtime.Func3(func(arrs_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_concat_prime_(*((*[]gopurs_runtime.Value)(arrs_0_box.UnsafePtr)), n_1_box.IntVal, v_2_box)
		})
	})
	return cache_Node_Buffer_concat_prime_
}

var cache_Node_Buffer_concat gopurs_runtime.Value
var once_Node_Buffer_concat sync.Once

func Get_Node_Buffer_concat() gopurs_runtime.Value {
	once_Node_Buffer_concat.Do(func() {
		cache_Node_Buffer_concat = gopurs_runtime.Func2(func(arrs_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_concat(*((*[]gopurs_runtime.Value)(arrs_0_box.UnsafePtr)), v_1_box)
		})
	})
	return cache_Node_Buffer_concat
}

var cache_Node_Buffer_compareParts gopurs_runtime.Value
var once_Node_Buffer_compareParts sync.Once

func Get_Node_Buffer_compareParts() gopurs_runtime.Value {
	once_Node_Buffer_compareParts.Do(func() {
		cache_Node_Buffer_compareParts = gopurs_runtime.Func6(func(src_0_box gopurs_runtime.Value, target_1_box gopurs_runtime.Value, targetSrc_2_box gopurs_runtime.Value, targetEnd_3_box gopurs_runtime.Value, srcStart_4_box gopurs_runtime.Value, srcEnd_5_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_compareParts(src_0_box, target_1_box, targetSrc_2_box.IntVal, targetEnd_3_box.IntVal, srcStart_4_box.IntVal, srcEnd_5_box.IntVal)
		})
	})
	return cache_Node_Buffer_compareParts
}

var cache_Node_Buffer_allocUnsafeSlow gopurs_runtime.Value
var once_Node_Buffer_allocUnsafeSlow sync.Once

func Get_Node_Buffer_allocUnsafeSlow() gopurs_runtime.Value {
	once_Node_Buffer_allocUnsafeSlow.Do(func() {
		cache_Node_Buffer_allocUnsafeSlow = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_allocUnsafeSlow(s_0_box.IntVal)
		})
	})
	return cache_Node_Buffer_allocUnsafeSlow
}

var cache_Node_Buffer_allocUnsafe gopurs_runtime.Value
var once_Node_Buffer_allocUnsafe sync.Once

func Get_Node_Buffer_allocUnsafe() gopurs_runtime.Value {
	once_Node_Buffer_allocUnsafe.Do(func() {
		cache_Node_Buffer_allocUnsafe = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_allocUnsafe(s_0_box.IntVal)
		})
	})
	return cache_Node_Buffer_allocUnsafe
}

var cache_Node_Buffer_alloc gopurs_runtime.Value
var once_Node_Buffer_alloc sync.Once

func Get_Node_Buffer_alloc() gopurs_runtime.Value {
	once_Node_Buffer_alloc.Do(func() {
		cache_Node_Buffer_alloc = gopurs_runtime.Apply(Get_Node_Buffer_usingToImmutable__4058095588(), Get_Node_Buffer_Immutable_alloc())
	})
	return cache_Node_Buffer_alloc
}

var cache_Node_Buffer_create gopurs_runtime.Value
var once_Node_Buffer_create sync.Once

func Get_Node_Buffer_create() gopurs_runtime.Value {
	once_Node_Buffer_create.Do(func() {
		cache_Node_Buffer_create = gopurs_runtime.Apply(Get_Node_Buffer_usingToImmutable__4058095588(), Get_Node_Buffer_Immutable_alloc())
	})
	return cache_Node_Buffer_create
}

var cache_Node_Buffer_mutableBufferEffect gopurs_runtime.Value
var once_Node_Buffer_mutableBufferEffect sync.Once

func Get_Node_Buffer_mutableBufferEffect() gopurs_runtime.Value {
	once_Node_Buffer_mutableBufferEffect.Do(func() {
		cache_Node_Buffer_mutableBufferEffect = gopurs_runtime.Value{Type: 9, IntVal: 1692970589, UnsafePtr: unsafe.Pointer((&Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}
		}), Get_Node_Buffer_concat(), Get_Node_Buffer_concat_prime_(), Get_Node_Buffer_copy(), gopurs_runtime.Apply(Get_Node_Buffer_usingToImmutable__4058095588(), Get_Node_Buffer_Immutable_alloc()), Get_Node_Buffer_fill(), Get_Node_Buffer_freeze(), gopurs_runtime.Apply(Get_Node_Buffer_usingToImmutable__3986645350(), Get_Node_Buffer_Immutable_fromArray()), gopurs_runtime.Apply(Get_Node_Buffer_usingToImmutable__2508462662(), Get_Node_Buffer_Immutable_fromArrayBuffer()), Get_Node_Buffer_fromString(), Get_Node_Buffer_getAtOffset(), Get_Node_Buffer_read(), Get_Node_Buffer_readString(), Get_Node_Buffer_setAtOffset(), gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable__2070034188(), Get_Node_Buffer_Immutable_size()), Get_Node_Buffer_Immutable_slice(), Get_Node_Buffer_thaw(), gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable__3335896427(), Get_Node_Buffer_Immutable_toArray()), gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable__3422936011(), Get_Node_Buffer_Immutable_toArrayBuffer()), Get_Node_Buffer_toString(), Get_Node_Buffer_unsafeFreeze(), Get_Node_Buffer_unsafeThaw(), Get_Node_Buffer_write(), Get_Node_Buffer_writeString()}))}
	})
	return cache_Node_Buffer_mutableBufferEffect
}

func Call_Node_Buffer_map__1922202664(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1922202664:
	for {
		if false {
			continue map__1922202664
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Node_Buffer_Immutable_toArrayBuffer(), __eta_norm_0_1)
	}
}

func Call_Node_Buffer_map__488316713(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__488316713:
	for {
		if false {
			continue map__488316713
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Node_Buffer_map__1198218120(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1198218120:
	for {
		if false {
			continue map__1198218120
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Node_Buffer_Immutable_toArray(), __eta_norm_0_1)
	}
}

func Call_Node_Buffer_map__3336206575(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__3336206575:
	for {
		if false {
			continue map__3336206575
		}
		var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
		_ = __eta_norm_1_unused_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Node_Buffer_Immutable_size(), __eta_norm_0_1)
	}
}

func Call_Node_Buffer_map__2503807369(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2503807369:
	for {
		if false {
			continue map__2503807369
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Node_Buffer_map__322707977(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__322707977:
	for {
		if false {
			continue map__322707977
		}
		var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
		_ = __eta_norm_1_0
		var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
		_ = __eta_norm_0_1
		return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), __eta_norm_1_0, __eta_norm_0_1)
	}
}

func Call_Node_Buffer_writeString(enc_0_loop uint32, offset_1_loop int64, len_2_loop int64, value_3_loop string, buf_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var enc_0 uint32 = enc_0_loop
	_ = enc_0
	var offset_1 int64 = offset_1_loop
	_ = offset_1
	var len_2 int64 = len_2_loop
	_ = len_2
	var value_3 string = value_3_loop
	_ = value_3
	var buf_4 gopurs_runtime.Value = buf_4_loop
	_ = buf_4
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 string
		{
			if enc_0 == 2598402200 {
				__t1 = "ascii"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_0 == 346813718 {
				__t1 = "utf8"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_0 == 2352678688 {
				__t1 = "utf16le"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_0 == 4287456318 {
				__t1 = "ucs2"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_0 == 2405785054 {
				__t1 = "base64"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_0 == 2476550165 {
				__t1 = "base64url"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_0 == 3576003686 {
				__t1 = "latin1"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_0 == 349879974 {
				__t1 = "binary"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_0 == 106823740 {
				__t1 = "hex"
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
		}
	end_branch_1:
		// TAST (Let): __local_var_5_0 shape=Branch(LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, def=Other) bindingType=Any
		__local_var_5_0 := __t1
		_ = __local_var_5_0
		return gopurs_runtime.UncurriedApp5(Get_Node_Buffer_writeStringInternal(), gopurs_runtime.Str(__local_var_5_0), gopurs_runtime.Int(offset_1), gopurs_runtime.Int(len_2), gopurs_runtime.Str(value_3), buf_4)
	})
}

func Call_Node_Buffer_write(ty_0_loop uint32, value_1_loop float64, offset_2_loop int64, buf_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var ty_0 uint32 = ty_0_loop
	_ = ty_0
	var value_1 float64 = value_1_loop
	_ = value_1
	var offset_2 int64 = offset_2_loop
	_ = offset_2
	var buf_3 gopurs_runtime.Value = buf_3_loop
	_ = buf_3
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 string
		{
			if ty_0 == 3314799936 {
				__t1 = "UInt8"
				goto end_branch_1
			} else {

			}
		}
		{
			if ty_0 == 852326966 {
				__t1 = "UInt16LE"
				goto end_branch_1
			} else {

			}
		}
		{
			if ty_0 == 305778936 {
				__t1 = "UInt16BE"
				goto end_branch_1
			} else {

			}
		}
		{
			if ty_0 == 3460170864 {
				__t1 = "UInt32LE"
				goto end_branch_1
			} else {

			}
		}
		{
			if ty_0 == 1992902334 {
				__t1 = "UInt32BE"
				goto end_branch_1
			} else {

			}
		}
		{
			if ty_0 == 1095082613 {
				__t1 = "Int8"
				goto end_branch_1
			} else {

			}
		}
		{
			if ty_0 == 408746499 {
				__t1 = "Int16LE"
				goto end_branch_1
			} else {

			}
		}
		{
			if ty_0 == 3066827405 {
				__t1 = "Int16BE"
				goto end_branch_1
			} else {

			}
		}
		{
			if ty_0 == 199750917 {
				__t1 = "Int32LE"
				goto end_branch_1
			} else {

			}
		}
		{
			if ty_0 == 3275822987 {
				__t1 = "Int32BE"
				goto end_branch_1
			} else {

			}
		}
		{
			if ty_0 == 1492841799 {
				__t1 = "FloatLE"
				goto end_branch_1
			} else {

			}
		}
		{
			if ty_0 == 4210130121 {
				__t1 = "FloatBE"
				goto end_branch_1
			} else {

			}
		}
		{
			if ty_0 == 496370018 {
				__t1 = "DoubleLE"
				goto end_branch_1
			} else {

			}
		}
		{
			if ty_0 == 894396972 {
				__t1 = "DoubleBE"
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
		}
	end_branch_1:
		// TAST (Let): __local_var_4_0 shape=Branch(LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, def=Other) bindingType=Any
		__local_var_4_0 := __t1
		_ = __local_var_4_0
		return gopurs_runtime.UncurriedApp4(Get_Node_Buffer_writeInternal(), gopurs_runtime.Str(__local_var_4_0), gopurs_runtime.Float(value_1), gopurs_runtime.Int(offset_2), buf_3)
	})
}

func Call_Node_Buffer_usingToImmutable(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var x_1 gopurs_runtime.Value = x_1_loop
	_ = x_1
	return gopurs_runtime.Apply(Get_Node_Buffer_unsafeThaw(), gopurs_runtime.Apply(f_0, x_1))
}

func Call_Node_Buffer_usingToImmutable__2508462662(f_unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
usingToImmutable__2508462662:
	for {
		if false {
			continue usingToImmutable__2508462662
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 gopurs_runtime.Value = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_Buffer_unsafeThaw(), gopurs_runtime.Apply(Get_Node_Buffer_Immutable_fromArrayBuffer(), x_1))
	}
}

func Call_Node_Buffer_usingToImmutable__1694059786(f_0_loop gopurs_runtime.Value, x_1_loop uint32) gopurs_runtime.Value {
usingToImmutable__1694059786:
	for {
		if false {
			continue usingToImmutable__1694059786
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var x_1 uint32 = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_Buffer_unsafeThaw(), gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: int64(x_1), UnsafePtr: nil}))
	}
}

func Call_Node_Buffer_usingToImmutable__3986645350(f_unused_0_loop gopurs_runtime.Value, x_1_loop []int64) gopurs_runtime.Value {
usingToImmutable__3986645350:
	for {
		if false {
			continue usingToImmutable__3986645350
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 []int64 = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_Buffer_unsafeThaw(), gopurs_runtime.Apply(Get_Node_Buffer_Immutable_fromArray(), func() gopurs_runtime.Value {
			arr := x_1
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()))
	}
}

func Call_Node_Buffer_usingToImmutable__4058095588(f_unused_0_loop gopurs_runtime.Value, x_1_loop int64) gopurs_runtime.Value {
usingToImmutable__4058095588:
	for {
		if false {
			continue usingToImmutable__4058095588
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var x_1 int64 = x_1_loop
		_ = x_1
		return gopurs_runtime.Apply(Get_Node_Buffer_unsafeThaw(), gopurs_runtime.Apply(Get_Node_Buffer_Immutable_alloc(), gopurs_runtime.Int(x_1)))
	}
}

func Call_Node_Buffer_usingFromImmutable(f_0_loop gopurs_runtime.Value, buf_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var buf_1 gopurs_runtime.Value = buf_1_loop
	_ = buf_1
	return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), f_0, gopurs_runtime.Apply(Get_Node_Buffer_unsafeFreeze(), buf_1))
}

func Call_Node_Buffer_usingFromImmutable__3422936011(f_unused_0_loop gopurs_runtime.Value, buf_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
usingFromImmutable__3422936011:
	for {
		if false {
			continue usingFromImmutable__3422936011
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var buf_1 gopurs_runtime.Value = buf_1_loop
		_ = buf_1
		return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Node_Buffer_Immutable_toArrayBuffer(), gopurs_runtime.Apply(Get_Node_Buffer_unsafeFreeze(), buf_1))
	}
}

func Call_Node_Buffer_usingFromImmutable__2525970890(f_0_loop gopurs_runtime.Value, buf_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
usingFromImmutable__2525970890:
	for {
		if false {
			continue usingFromImmutable__2525970890
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var buf_1 gopurs_runtime.Value = buf_1_loop
		_ = buf_1
		return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), f_0, gopurs_runtime.Apply(Get_Node_Buffer_unsafeFreeze(), buf_1))
	}
}

func Call_Node_Buffer_usingFromImmutable__3335896427(f_unused_0_loop gopurs_runtime.Value, buf_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
usingFromImmutable__3335896427:
	for {
		if false {
			continue usingFromImmutable__3335896427
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var buf_1 gopurs_runtime.Value = buf_1_loop
		_ = buf_1
		return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Node_Buffer_Immutable_toArray(), gopurs_runtime.Apply(Get_Node_Buffer_unsafeFreeze(), buf_1))
	}
}

func Call_Node_Buffer_usingFromImmutable__2070034188(f_unused_0_loop gopurs_runtime.Value, buf_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
usingFromImmutable__2070034188:
	for {
		if false {
			continue usingFromImmutable__2070034188
		}
		var f_unused_0 gopurs_runtime.Value = f_unused_0_loop
		_ = f_unused_0
		var buf_1 gopurs_runtime.Value = buf_1_loop
		_ = buf_1
		return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Node_Buffer_Immutable_size(), gopurs_runtime.Apply(Get_Node_Buffer_unsafeFreeze(), buf_1))
	}
}

func Call_Node_Buffer_usingFromImmutable__2892494410(f_0_loop gopurs_runtime.Value, buf_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
usingFromImmutable__2892494410:
	for {
		if false {
			continue usingFromImmutable__2892494410
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var buf_1 gopurs_runtime.Value = buf_1_loop
		_ = buf_1
		return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), f_0, gopurs_runtime.Apply(Get_Node_Buffer_unsafeFreeze(), buf_1))
	}
}

func Call_Node_Buffer_usingFromImmutable__3537276362(f_0_loop gopurs_runtime.Value, buf_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
usingFromImmutable__3537276362:
	for {
		if false {
			continue usingFromImmutable__3537276362
		}
		var f_0 gopurs_runtime.Value = f_0_loop
		_ = f_0
		var buf_1 gopurs_runtime.Value = buf_1_loop
		_ = buf_1
		return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), f_0, gopurs_runtime.Apply(Get_Node_Buffer_unsafeFreeze(), buf_1))
	}
}

func Call_Node_Buffer_transcode(buf_0_loop gopurs_runtime.Value, from_1_loop uint32, to_2_loop uint32) gopurs_runtime.Value {
	var buf_0 gopurs_runtime.Value = buf_0_loop
	_ = buf_0
	var from_1 uint32 = from_1_loop
	_ = from_1
	var to_2 uint32 = to_2_loop
	_ = to_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 string
		{
			if from_1 == 2598402200 {
				__t1 = "ascii"
				goto end_branch_1
			} else {

			}
		}
		{
			if from_1 == 346813718 {
				__t1 = "utf8"
				goto end_branch_1
			} else {

			}
		}
		{
			if from_1 == 2352678688 {
				__t1 = "utf16le"
				goto end_branch_1
			} else {

			}
		}
		{
			if from_1 == 4287456318 {
				__t1 = "ucs2"
				goto end_branch_1
			} else {

			}
		}
		{
			if from_1 == 2405785054 {
				__t1 = "base64"
				goto end_branch_1
			} else {

			}
		}
		{
			if from_1 == 2476550165 {
				__t1 = "base64url"
				goto end_branch_1
			} else {

			}
		}
		{
			if from_1 == 3576003686 {
				__t1 = "latin1"
				goto end_branch_1
			} else {

			}
		}
		{
			if from_1 == 349879974 {
				__t1 = "binary"
				goto end_branch_1
			} else {

			}
		}
		{
			if from_1 == 106823740 {
				__t1 = "hex"
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
		}
	end_branch_1:
		// TAST (Let): __local_var_3_0 shape=Branch(LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, def=Other) bindingType=Any
		__local_var_3_0 := __t1
		_ = __local_var_3_0
		var __t3 string
		{
			if to_2 == 2598402200 {
				__t3 = "ascii"
				goto end_branch_3
			} else {

			}
		}
		{
			if to_2 == 346813718 {
				__t3 = "utf8"
				goto end_branch_3
			} else {

			}
		}
		{
			if to_2 == 2352678688 {
				__t3 = "utf16le"
				goto end_branch_3
			} else {

			}
		}
		{
			if to_2 == 4287456318 {
				__t3 = "ucs2"
				goto end_branch_3
			} else {

			}
		}
		{
			if to_2 == 2405785054 {
				__t3 = "base64"
				goto end_branch_3
			} else {

			}
		}
		{
			if to_2 == 2476550165 {
				__t3 = "base64url"
				goto end_branch_3
			} else {

			}
		}
		{
			if to_2 == 3576003686 {
				__t3 = "latin1"
				goto end_branch_3
			} else {

			}
		}
		{
			if to_2 == 349879974 {
				__t3 = "binary"
				goto end_branch_3
			} else {

			}
		}
		{
			if to_2 == 106823740 {
				__t3 = "hex"
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
		}
	end_branch_3:
		// TAST (Let): __local_var_4_2 shape=Branch(LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, def=Other) bindingType=Any
		__local_var_4_2 := __t3
		_ = __local_var_4_2
		return gopurs_runtime.UncurriedApp3(Get_Node_Buffer_transcodeImpl(), buf_0, gopurs_runtime.Str(__local_var_3_0), gopurs_runtime.Str(__local_var_4_2))
	})
}

func Call_Node_Buffer_toString_prime_(enc_0_loop uint32, start_1_loop int64, end_2_loop int64) gopurs_runtime.Value {
	var enc_0 uint32 = enc_0_loop
	_ = enc_0
	var start_1 int64 = start_1_loop
	_ = start_1
	var end_2 int64 = end_2_loop
	_ = end_2
	return gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable(), gopurs_runtime.Apply3(Get_Node_Buffer_Immutable_toString_prime_(), gopurs_runtime.Value{Type: 9, IntVal: int64(enc_0), UnsafePtr: nil}, gopurs_runtime.Int(start_1), gopurs_runtime.Int(end_2)))
}

func Call_Node_Buffer_toString(enc_0_loop uint32) gopurs_runtime.Value {
	var enc_0 uint32 = enc_0_loop
	_ = enc_0
	return gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable(), gopurs_runtime.Apply(Get_Node_Buffer_Immutable_toString(), gopurs_runtime.Value{Type: 9, IntVal: int64(enc_0), UnsafePtr: nil}))
}

func Call_Node_Buffer_swap64(b_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Buffer_swap64Impl(), b_0)
	})
}

func Call_Node_Buffer_swap32(b_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Buffer_swap32Impl(), b_0)
	})
}

func Call_Node_Buffer_swap16(b_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var b_0 gopurs_runtime.Value = b_0_loop
	_ = b_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Buffer_swap16Impl(), b_0)
	})
}

func Call_Node_Buffer_setPoolSize(sizeInBytes_0_loop int64) gopurs_runtime.Value {
	var sizeInBytes_0 int64 = sizeInBytes_0_loop
	_ = sizeInBytes_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Buffer_setPoolSizeImpl(), gopurs_runtime.Int(sizeInBytes_0))
	})
}

func Call_Node_Buffer_setAtOffset(val_0_loop int64, off_1_loop int64, buff_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var val_0 int64 = val_0_loop
	_ = val_0
	var off_1 int64 = off_1_loop
	_ = off_1
	var buff_2 gopurs_runtime.Value = buff_2_loop
	_ = buff_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp3(Get_Node_Buffer_setAtOffsetImpl(), gopurs_runtime.Int(val_0), gopurs_runtime.Int(off_1), buff_2)
	})
}

func Call_Node_Buffer_readString(enc_0_loop uint32, o_1_loop int64, o_prime__2_loop int64) gopurs_runtime.Value {
	var enc_0 uint32 = enc_0_loop
	_ = enc_0
	var o_1 int64 = o_1_loop
	_ = o_1
	var o_prime__2 int64 = o_prime__2_loop
	_ = o_prime__2
	return gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable(), gopurs_runtime.Apply3(Get_Node_Buffer_Immutable_readString(), gopurs_runtime.Value{Type: 9, IntVal: int64(enc_0), UnsafePtr: nil}, gopurs_runtime.Int(o_1), gopurs_runtime.Int(o_prime__2)))
}

func Call_Node_Buffer_read(t_0_loop uint32, o_1_loop int64) gopurs_runtime.Value {
	var t_0 uint32 = t_0_loop
	_ = t_0
	var o_1 int64 = o_1_loop
	_ = o_1
	return gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable(), gopurs_runtime.Apply2(Get_Node_Buffer_Immutable_read(), gopurs_runtime.Value{Type: 9, IntVal: int64(t_0), UnsafePtr: nil}, gopurs_runtime.Int(o_1)))
}

func Call_Node_Buffer_getAtOffset(o_0_loop int64) gopurs_runtime.Value {
	var o_0 int64 = o_0_loop
	_ = o_0
	return gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable(), gopurs_runtime.Apply(Get_Node_Buffer_Immutable_getAtOffset(), gopurs_runtime.Int(o_0)))
}

func Call_Node_Buffer_fromString(s_0_loop string) gopurs_runtime.Value {
	var s_0 string = s_0_loop
	_ = s_0
	return gopurs_runtime.Apply(Get_Node_Buffer_usingToImmutable(), gopurs_runtime.Apply(Get_Node_Buffer_Immutable_fromString(), gopurs_runtime.Str(s_0)))
}

func Call_Node_Buffer_fill(octet_0_loop int64, start_1_loop int64, end_2_loop int64, buf_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var octet_0 int64 = octet_0_loop
	_ = octet_0
	var start_1 int64 = start_1_loop
	_ = start_1
	var end_2 int64 = end_2_loop
	_ = end_2
	var buf_3 gopurs_runtime.Value = buf_3_loop
	_ = buf_3
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp4(Get_Node_Buffer_fillImpl(), gopurs_runtime.Int(octet_0), gopurs_runtime.Int(start_1), gopurs_runtime.Int(end_2), buf_3)
	})
}

func Call_Node_Buffer_copy(srcStart_0_loop int64, srcEnd_1_loop int64, src_2_loop gopurs_runtime.Value, targStart_3_loop int64, targ_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var srcStart_0 int64 = srcStart_0_loop
	_ = srcStart_0
	var srcEnd_1 int64 = srcEnd_1_loop
	_ = srcEnd_1
	var src_2 gopurs_runtime.Value = src_2_loop
	_ = src_2
	var targStart_3 int64 = targStart_3_loop
	_ = targStart_3
	var targ_4 gopurs_runtime.Value = targ_4_loop
	_ = targ_4
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp5(Get_Node_Buffer_copyImpl(), gopurs_runtime.Int(srcStart_0), gopurs_runtime.Int(srcEnd_1), src_2, gopurs_runtime.Int(targStart_3), targ_4)
	})
}

func Call_Node_Buffer_concat_prime_(arrs_0_loop []gopurs_runtime.Value, n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var arrs_0 []gopurs_runtime.Value = arrs_0_loop
	_ = arrs_0
	var n_1 int64 = n_1_loop
	_ = n_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.UncurriedApp2(Get_Node_Buffer_Immutable_concatToLength(), gopurs_runtime.Array(arrs_0), gopurs_runtime.Int(n_1))
}

func Call_Node_Buffer_concat(arrs_0_loop []gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var arrs_0 []gopurs_runtime.Value = arrs_0_loop
	_ = arrs_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return gopurs_runtime.Apply(Get_Node_Buffer_Immutable_concat(), gopurs_runtime.Array(arrs_0))
}

func Call_Node_Buffer_compareParts(src_0_loop gopurs_runtime.Value, target_1_loop gopurs_runtime.Value, targetSrc_2_loop int64, targetEnd_3_loop int64, srcStart_4_loop int64, srcEnd_5_loop int64) gopurs_runtime.Value {
	var src_0 gopurs_runtime.Value = src_0_loop
	_ = src_0
	var target_1 gopurs_runtime.Value = target_1_loop
	_ = target_1
	var targetSrc_2 int64 = targetSrc_2_loop
	_ = targetSrc_2
	var targetEnd_3 int64 = targetEnd_3_loop
	_ = targetEnd_3
	var srcStart_4 int64 = srcStart_4_loop
	_ = srcStart_4
	var srcEnd_5 int64 = srcEnd_5_loop
	_ = srcEnd_5
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_6_0 shape=App(Var) bindingType=Any
		__local_var_6_0 := gopurs_runtime.Apply(Get_Node_Buffer_unsafeFreeze(), src_0)
		_ = __local_var_6_0
		__local_var_7_1 := gopurs_runtime.Apply(__local_var_6_0, gopurs_runtime.Value{})
		_ = __local_var_7_1
		__local_var_8_2 := gopurs_runtime.Apply2(Get_Node_Buffer_unsafeFreeze(), target_1, gopurs_runtime.Value{})
		_ = __local_var_8_2
		return gopurs_runtime.Apply(Call_Node_Buffer_Immutable_compareParts(__local_var_7_1, __local_var_8_2, targetSrc_2, targetEnd_3, srcStart_4, srcEnd_5), gopurs_runtime.Value{})
	})
}

func Call_Node_Buffer_allocUnsafeSlow(s_0_loop int64) gopurs_runtime.Value {
	var s_0 int64 = s_0_loop
	_ = s_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Buffer_allocUnsafeSlowImpl(), gopurs_runtime.Int(s_0))
	})
}

func Call_Node_Buffer_allocUnsafe(s_0_loop int64) gopurs_runtime.Value {
	var s_0 int64 = s_0_loop
	_ = s_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Buffer_allocUnsafeImpl(), gopurs_runtime.Int(s_0))
	})
}

func Get_Node_Buffer_allocUnsafeImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_AllocUnsafeImpl
}

func Get_Node_Buffer_allocUnsafeSlowImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_AllocUnsafeSlowImpl
}

func Get_Node_Buffer_copyImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_CopyImpl
}

func Get_Node_Buffer_fillImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_FillImpl
}

func Get_Node_Buffer_freezeImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_FreezeImpl
}

func Get_Node_Buffer_poolSize() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_PoolSize
}

func Get_Node_Buffer_setAtOffsetImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_SetAtOffsetImpl
}

func Get_Node_Buffer_setPoolSizeImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_SetPoolSizeImpl
}

func Get_Node_Buffer_swap16Impl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Swap16Impl
}

func Get_Node_Buffer_swap32Impl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Swap32Impl
}

func Get_Node_Buffer_swap64Impl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_Swap64Impl
}

func Get_Node_Buffer_thawImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_ThawImpl
}

func Get_Node_Buffer_transcodeImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_TranscodeImpl
}

func Get_Node_Buffer_writeInternal() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_WriteInternal
}

func Get_Node_Buffer_writeStringInternal() gopurs_runtime.Value {
	return _Gopurs_Node_Buffer_WriteStringInternal
}
