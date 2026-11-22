package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Node_Stream_show gopurs_runtime.Value
var once_Node_Stream_show sync.Once

func Get_Node_Stream_show() gopurs_runtime.Value {
	once_Node_Stream_show.Do(func() {
		cache_Node_Stream_show = Get_Data_Show_show__2181193092()
	})
	return cache_Node_Stream_show
}

var cache_Node_Stream_identity gopurs_runtime.Value
var once_Node_Stream_identity sync.Once

func Get_Node_Stream_identity() gopurs_runtime.Value {
	once_Node_Stream_identity.Do(func() {
		cache_Node_Stream_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_identity(x_0_box)
		})
	})
	return cache_Node_Stream_identity
}

var cache_Node_Stream_bind gopurs_runtime.Value
var once_Node_Stream_bind sync.Once

func Get_Node_Stream_bind() gopurs_runtime.Value {
	once_Node_Stream_bind.Do(func() {
		cache_Node_Stream_bind = Get_Control_Bind_bind__845594356()
	})
	return cache_Node_Stream_bind
}

var cache_Node_Stream_pure gopurs_runtime.Value
var once_Node_Stream_pure sync.Once

func Get_Node_Stream_pure() gopurs_runtime.Value {
	once_Node_Stream_pure.Do(func() {
		cache_Node_Stream_pure = Get_Control_Applicative_pure__3252509724()
	})
	return cache_Node_Stream_pure
}

var cache_Node_Stream_compose gopurs_runtime.Value
var once_Node_Stream_compose sync.Once

func Get_Node_Stream_compose() gopurs_runtime.Value {
	once_Node_Stream_compose.Do(func() {
		cache_Node_Stream_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_compose
}

var cache_Node_Stream_pure1 gopurs_runtime.Value
var once_Node_Stream_pure1 sync.Once

func Get_Node_Stream_pure1() gopurs_runtime.Value {
	once_Node_Stream_pure1.Do(func() {
		cache_Node_Stream_pure1 = Get_Effect_pureE()
	})
	return cache_Node_Stream_pure1
}

var cache_Node_Stream_compose1 gopurs_runtime.Value
var once_Node_Stream_compose1 sync.Once

func Get_Node_Stream_compose1() gopurs_runtime.Value {
	once_Node_Stream_compose1.Do(func() {
		cache_Node_Stream_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_compose1
}

var cache_Node_Stream_compose2 gopurs_runtime.Value
var once_Node_Stream_compose2 sync.Once

func Get_Node_Stream_compose2() gopurs_runtime.Value {
	once_Node_Stream_compose2.Do(func() {
		cache_Node_Stream_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_compose2
}

var cache_Node_Stream_compose3 gopurs_runtime.Value
var once_Node_Stream_compose3 sync.Once

func Get_Node_Stream_compose3() gopurs_runtime.Value {
	once_Node_Stream_compose3.Do(func() {
		cache_Node_Stream_compose3 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_compose3
}

var cache_Node_Stream_compose4 gopurs_runtime.Value
var once_Node_Stream_compose4 sync.Once

func Get_Node_Stream_compose4() gopurs_runtime.Value {
	once_Node_Stream_compose4.Do(func() {
		cache_Node_Stream_compose4 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_compose4
}

var cache_Node_Stream_compose5 gopurs_runtime.Value
var once_Node_Stream_compose5 sync.Once

func Get_Node_Stream_compose5() gopurs_runtime.Value {
	once_Node_Stream_compose5.Do(func() {
		cache_Node_Stream_compose5 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_compose5
}

var cache_Node_Stream_compose6 gopurs_runtime.Value
var once_Node_Stream_compose6 sync.Once

func Get_Node_Stream_compose6() gopurs_runtime.Value {
	once_Node_Stream_compose6.Do(func() {
		cache_Node_Stream_compose6 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_compose6
}

var cache_Node_Stream_compose7 gopurs_runtime.Value
var once_Node_Stream_compose7 sync.Once

func Get_Node_Stream_compose7() gopurs_runtime.Value {
	once_Node_Stream_compose7.Do(func() {
		cache_Node_Stream_compose7 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_compose7
}

var cache_Node_Stream_bind1 gopurs_runtime.Value
var once_Node_Stream_bind1 sync.Once

func Get_Node_Stream_bind1() gopurs_runtime.Value {
	once_Node_Stream_bind1.Do(func() {
		cache_Node_Stream_bind1 = Get_Control_Bind_bind__1207231796()
	})
	return cache_Node_Stream_bind1
}

var cache_Node_Stream_pure2 gopurs_runtime.Value
var once_Node_Stream_pure2 sync.Once

func Get_Node_Stream_pure2() gopurs_runtime.Value {
	once_Node_Stream_pure2.Do(func() {
		cache_Node_Stream_pure2 = Get_Control_Applicative_pure__1238803548()
	})
	return cache_Node_Stream_pure2
}

var cache_Node_Stream_pure3 gopurs_runtime.Value
var once_Node_Stream_pure3 sync.Once

func Get_Node_Stream_pure3() gopurs_runtime.Value {
	once_Node_Stream_pure3.Do(func() {
		cache_Node_Stream_pure3 = Get_Effect_pureE()
	})
	return cache_Node_Stream_pure3
}

var cache_Node_Stream_bind2 gopurs_runtime.Value
var once_Node_Stream_bind2 sync.Once

func Get_Node_Stream_bind2() gopurs_runtime.Value {
	once_Node_Stream_bind2.Do(func() {
		cache_Node_Stream_bind2 = Get_Control_Bind_bind__949378708()
	})
	return cache_Node_Stream_bind2
}

var cache_Node_Stream_pure4 gopurs_runtime.Value
var once_Node_Stream_pure4 sync.Once

func Get_Node_Stream_pure4() gopurs_runtime.Value {
	once_Node_Stream_pure4.Do(func() {
		cache_Node_Stream_pure4 = Get_Control_Applicative_pure__1139843612()
	})
	return cache_Node_Stream_pure4
}

var cache_Node_Stream_go__map gopurs_runtime.Value
var once_Node_Stream_go__map sync.Once

func Get_Node_Stream_go__map() gopurs_runtime.Value {
	once_Node_Stream_go__map.Do(func() {
		cache_Node_Stream_go__map = Get_Data_Functor_map__3707757371()
	})
	return cache_Node_Stream_go__map
}

var cache_Node_Stream_compose8 gopurs_runtime.Value
var once_Node_Stream_compose8 sync.Once

func Get_Node_Stream_compose8() gopurs_runtime.Value {
	once_Node_Stream_compose8.Do(func() {
		cache_Node_Stream_compose8 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_compose8
}

var cache_Node_Stream_compose9 gopurs_runtime.Value
var once_Node_Stream_compose9 sync.Once

func Get_Node_Stream_compose9() gopurs_runtime.Value {
	once_Node_Stream_compose9.Do(func() {
		cache_Node_Stream_compose9 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Node_Stream_compose9
}

var cache_Node_Stream_writeableNeedDrain gopurs_runtime.Value
var once_Node_Stream_writeableNeedDrain sync.Once

func Get_Node_Stream_writeableNeedDrain() gopurs_runtime.Value {
	once_Node_Stream_writeableNeedDrain.Do(func() {
		cache_Node_Stream_writeableNeedDrain = gopurs_runtime.Func(func(w_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_writeableNeedDrain(w_0_box)
		})
	})
	return cache_Node_Stream_writeableNeedDrain
}

var cache_Node_Stream_writeableLength gopurs_runtime.Value
var once_Node_Stream_writeableLength sync.Once

func Get_Node_Stream_writeableLength() gopurs_runtime.Value {
	once_Node_Stream_writeableLength.Do(func() {
		cache_Node_Stream_writeableLength = gopurs_runtime.Func(func(w_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_writeableLength(w_0_box)
		})
	})
	return cache_Node_Stream_writeableLength
}

var cache_Node_Stream_writeableHighWaterMark gopurs_runtime.Value
var once_Node_Stream_writeableHighWaterMark sync.Once

func Get_Node_Stream_writeableHighWaterMark() gopurs_runtime.Value {
	once_Node_Stream_writeableHighWaterMark.Do(func() {
		cache_Node_Stream_writeableHighWaterMark = gopurs_runtime.Func(func(w_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_writeableHighWaterMark(w_0_box)
		})
	})
	return cache_Node_Stream_writeableHighWaterMark
}

var cache_Node_Stream_writeableFinished gopurs_runtime.Value
var once_Node_Stream_writeableFinished sync.Once

func Get_Node_Stream_writeableFinished() gopurs_runtime.Value {
	once_Node_Stream_writeableFinished.Do(func() {
		cache_Node_Stream_writeableFinished = gopurs_runtime.Func(func(w_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_writeableFinished(w_0_box)
		})
	})
	return cache_Node_Stream_writeableFinished
}

var cache_Node_Stream_writeableEnded gopurs_runtime.Value
var once_Node_Stream_writeableEnded sync.Once

func Get_Node_Stream_writeableEnded() gopurs_runtime.Value {
	once_Node_Stream_writeableEnded.Do(func() {
		cache_Node_Stream_writeableEnded = gopurs_runtime.Func(func(w_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_writeableEnded(w_0_box)
		})
	})
	return cache_Node_Stream_writeableEnded
}

var cache_Node_Stream_writeableCorked gopurs_runtime.Value
var once_Node_Stream_writeableCorked sync.Once

func Get_Node_Stream_writeableCorked() gopurs_runtime.Value {
	once_Node_Stream_writeableCorked.Do(func() {
		cache_Node_Stream_writeableCorked = gopurs_runtime.Func(func(w_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_writeableCorked(w_0_box)
		})
	})
	return cache_Node_Stream_writeableCorked
}

var cache_Node_Stream_writeable gopurs_runtime.Value
var once_Node_Stream_writeable sync.Once

func Get_Node_Stream_writeable() gopurs_runtime.Value {
	once_Node_Stream_writeable.Do(func() {
		cache_Node_Stream_writeable = gopurs_runtime.Func(func(w_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_writeable(w_0_box)
		})
	})
	return cache_Node_Stream_writeable
}

var cache_Node_Stream_writeString_prime_ gopurs_runtime.Value
var once_Node_Stream_writeString_prime_ sync.Once

func Get_Node_Stream_writeString_prime_() gopurs_runtime.Value {
	once_Node_Stream_writeString_prime_.Do(func() {
		cache_Node_Stream_writeString_prime_ = gopurs_runtime.Func4(func(w_0_box gopurs_runtime.Value, enc_1_box gopurs_runtime.Value, str_2_box gopurs_runtime.Value, cb_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_writeString_prime_(w_0_box, uint32(enc_1_box.IntVal), str_2_box.StrVal(), cb_3_box)
		})
	})
	return cache_Node_Stream_writeString_prime_
}

var cache_Node_Stream_writeString gopurs_runtime.Value
var once_Node_Stream_writeString sync.Once

func Get_Node_Stream_writeString() gopurs_runtime.Value {
	once_Node_Stream_writeString.Do(func() {
		cache_Node_Stream_writeString = gopurs_runtime.Func3(func(w_0_box gopurs_runtime.Value, enc_1_box gopurs_runtime.Value, str_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_writeString(w_0_box, uint32(enc_1_box.IntVal), str_2_box.StrVal())
		})
	})
	return cache_Node_Stream_writeString
}

var cache_Node_Stream_write_prime_ gopurs_runtime.Value
var once_Node_Stream_write_prime_ sync.Once

func Get_Node_Stream_write_prime_() gopurs_runtime.Value {
	once_Node_Stream_write_prime_.Do(func() {
		cache_Node_Stream_write_prime_ = gopurs_runtime.Func3(func(w_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, cb_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_write_prime_(w_0_box, b_1_box, cb_2_box)
		})
	})
	return cache_Node_Stream_write_prime_
}

var cache_Node_Stream_write gopurs_runtime.Value
var once_Node_Stream_write sync.Once

func Get_Node_Stream_write() gopurs_runtime.Value {
	once_Node_Stream_write.Do(func() {
		cache_Node_Stream_write = gopurs_runtime.Func2(func(w_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_write(w_0_box, b_1_box)
		})
	})
	return cache_Node_Stream_write
}

var cache_Node_Stream_unpipeH gopurs_runtime.Value
var once_Node_Stream_unpipeH sync.Once

func Get_Node_Stream_unpipeH() gopurs_runtime.Value {
	once_Node_Stream_unpipeH.Do(func() {
		cache_Node_Stream_unpipeH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "unpipe", Get_Effect_Uncurried_mkEffectFn1()}))}
	})
	return cache_Node_Stream_unpipeH
}

var cache_Node_Stream_unpipeAll gopurs_runtime.Value
var once_Node_Stream_unpipeAll sync.Once

func Get_Node_Stream_unpipeAll() gopurs_runtime.Value {
	once_Node_Stream_unpipeAll.Do(func() {
		cache_Node_Stream_unpipeAll = gopurs_runtime.Func(func(r_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_unpipeAll(r_0_box)
		})
	})
	return cache_Node_Stream_unpipeAll
}

var cache_Node_Stream_unpipe gopurs_runtime.Value
var once_Node_Stream_unpipe sync.Once

func Get_Node_Stream_unpipe() gopurs_runtime.Value {
	once_Node_Stream_unpipe.Do(func() {
		cache_Node_Stream_unpipe = gopurs_runtime.Func2(func(r_0_box gopurs_runtime.Value, w_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_unpipe(r_0_box, w_1_box)
		})
	})
	return cache_Node_Stream_unpipe
}

var cache_Node_Stream_uncork gopurs_runtime.Value
var once_Node_Stream_uncork sync.Once

func Get_Node_Stream_uncork() gopurs_runtime.Value {
	once_Node_Stream_uncork.Do(func() {
		cache_Node_Stream_uncork = gopurs_runtime.Func(func(w_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_uncork(w_0_box)
		})
	})
	return cache_Node_Stream_uncork
}

var cache_Node_Stream_toEventEmitter gopurs_runtime.Value
var once_Node_Stream_toEventEmitter sync.Once

func Get_Node_Stream_toEventEmitter() gopurs_runtime.Value {
	once_Node_Stream_toEventEmitter.Do(func() {
		cache_Node_Stream_toEventEmitter = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Node_Stream_toEventEmitter
}

var cache_Node_Stream_setEncoding gopurs_runtime.Value
var once_Node_Stream_setEncoding sync.Once

func Get_Node_Stream_setEncoding() gopurs_runtime.Value {
	once_Node_Stream_setEncoding.Do(func() {
		cache_Node_Stream_setEncoding = gopurs_runtime.Func2(func(r_0_box gopurs_runtime.Value, enc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_setEncoding(r_0_box, uint32(enc_1_box.IntVal))
		})
	})
	return cache_Node_Stream_setEncoding
}

var cache_Node_Stream_setDefaultEncoding gopurs_runtime.Value
var once_Node_Stream_setDefaultEncoding sync.Once

func Get_Node_Stream_setDefaultEncoding() gopurs_runtime.Value {
	once_Node_Stream_setDefaultEncoding.Do(func() {
		cache_Node_Stream_setDefaultEncoding = gopurs_runtime.Func2(func(r_0_box gopurs_runtime.Value, enc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_setDefaultEncoding(r_0_box, uint32(enc_1_box.IntVal))
		})
	})
	return cache_Node_Stream_setDefaultEncoding
}

var cache_Node_Stream_resumeH gopurs_runtime.Value
var once_Node_Stream_resumeH sync.Once

func Get_Node_Stream_resumeH() gopurs_runtime.Value {
	once_Node_Stream_resumeH.Do(func() {
		cache_Node_Stream_resumeH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "resume", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Node_Stream_resumeH
}

var cache_Node_Stream_resume gopurs_runtime.Value
var once_Node_Stream_resume sync.Once

func Get_Node_Stream_resume() gopurs_runtime.Value {
	once_Node_Stream_resume.Do(func() {
		cache_Node_Stream_resume = gopurs_runtime.Func(func(r_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_resume(r_0_box)
		})
	})
	return cache_Node_Stream_resume
}

var cache_Node_Stream_readableLength gopurs_runtime.Value
var once_Node_Stream_readableLength sync.Once

func Get_Node_Stream_readableLength() gopurs_runtime.Value {
	once_Node_Stream_readableLength.Do(func() {
		cache_Node_Stream_readableLength = gopurs_runtime.Func(func(r_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_readableLength(r_0_box)
		})
	})
	return cache_Node_Stream_readableLength
}

var cache_Node_Stream_readableHighWaterMark gopurs_runtime.Value
var once_Node_Stream_readableHighWaterMark sync.Once

func Get_Node_Stream_readableHighWaterMark() gopurs_runtime.Value {
	once_Node_Stream_readableHighWaterMark.Do(func() {
		cache_Node_Stream_readableHighWaterMark = gopurs_runtime.Func(func(r_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_readableHighWaterMark(r_0_box)
		})
	})
	return cache_Node_Stream_readableHighWaterMark
}

var cache_Node_Stream_readableH gopurs_runtime.Value
var once_Node_Stream_readableH sync.Once

func Get_Node_Stream_readableH() gopurs_runtime.Value {
	once_Node_Stream_readableH.Do(func() {
		cache_Node_Stream_readableH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "readable", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Node_Stream_readableH
}

var cache_Node_Stream_readableFromString gopurs_runtime.Value
var once_Node_Stream_readableFromString sync.Once

func Get_Node_Stream_readableFromString() gopurs_runtime.Value {
	once_Node_Stream_readableFromString.Do(func() {
		cache_Node_Stream_readableFromString = gopurs_runtime.Func2(func(str_0_box gopurs_runtime.Value, enc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_readableFromString(str_0_box.StrVal(), uint32(enc_1_box.IntVal))
		})
	})
	return cache_Node_Stream_readableFromString
}

var cache_Node_Stream_readableFromBuffer gopurs_runtime.Value
var once_Node_Stream_readableFromBuffer sync.Once

func Get_Node_Stream_readableFromBuffer() gopurs_runtime.Value {
	once_Node_Stream_readableFromBuffer.Do(func() {
		cache_Node_Stream_readableFromBuffer = gopurs_runtime.Func(func(buf_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_readableFromBuffer(buf_0_box)
		})
	})
	return cache_Node_Stream_readableFromBuffer
}

var cache_Node_Stream_readableFlowing gopurs_runtime.Value
var once_Node_Stream_readableFlowing sync.Once

func Get_Node_Stream_readableFlowing() gopurs_runtime.Value {
	once_Node_Stream_readableFlowing.Do(func() {
		cache_Node_Stream_readableFlowing = gopurs_runtime.Func(func(r_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_readableFlowing(r_0_box)
		})
	})
	return cache_Node_Stream_readableFlowing
}

var cache_Node_Stream_readableEnded gopurs_runtime.Value
var once_Node_Stream_readableEnded sync.Once

func Get_Node_Stream_readableEnded() gopurs_runtime.Value {
	once_Node_Stream_readableEnded.Do(func() {
		cache_Node_Stream_readableEnded = gopurs_runtime.Func(func(r_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_readableEnded(r_0_box)
		})
	})
	return cache_Node_Stream_readableEnded
}

var cache_Node_Stream_readable gopurs_runtime.Value
var once_Node_Stream_readable sync.Once

func Get_Node_Stream_readable() gopurs_runtime.Value {
	once_Node_Stream_readable.Do(func() {
		cache_Node_Stream_readable = gopurs_runtime.Func(func(r_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_readable(r_0_box)
		})
	})
	return cache_Node_Stream_readable
}

var cache_Node_Stream_readEither_prime_ gopurs_runtime.Value
var once_Node_Stream_readEither_prime_ sync.Once

func Get_Node_Stream_readEither_prime_() gopurs_runtime.Value {
	once_Node_Stream_readEither_prime_.Do(func() {
		cache_Node_Stream_readEither_prime_ = gopurs_runtime.Func2(func(r_0_box gopurs_runtime.Value, size_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_readEither_prime_(r_0_box, size_1_box.IntVal)
		})
	})
	return cache_Node_Stream_readEither_prime_
}

var cache_Node_Stream_readEither gopurs_runtime.Value
var once_Node_Stream_readEither sync.Once

func Get_Node_Stream_readEither() gopurs_runtime.Value {
	once_Node_Stream_readEither.Do(func() {
		cache_Node_Stream_readEither = gopurs_runtime.Func(func(r_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_readEither(r_0_box)
		})
	})
	return cache_Node_Stream_readEither
}

var cache_Node_Stream_read_prime_ gopurs_runtime.Value
var once_Node_Stream_read_prime_ sync.Once

func Get_Node_Stream_read_prime_() gopurs_runtime.Value {
	once_Node_Stream_read_prime_.Do(func() {
		cache_Node_Stream_read_prime_ = gopurs_runtime.Func2(func(r_0_box gopurs_runtime.Value, size_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_read_prime_(r_0_box, size_1_box.IntVal)
		})
	})
	return cache_Node_Stream_read_prime_
}

var cache_Node_Stream_readString_prime_ gopurs_runtime.Value
var once_Node_Stream_readString_prime_ sync.Once

func Get_Node_Stream_readString_prime_() gopurs_runtime.Value {
	once_Node_Stream_readString_prime_.Do(func() {
		cache_Node_Stream_readString_prime_ = gopurs_runtime.Func3(func(r_0_box gopurs_runtime.Value, size_1_box gopurs_runtime.Value, enc_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_readString_prime_(r_0_box, size_1_box.IntVal, uint32(enc_2_box.IntVal))
		})
	})
	return cache_Node_Stream_readString_prime_
}

var cache_Node_Stream_read gopurs_runtime.Value
var once_Node_Stream_read sync.Once

func Get_Node_Stream_read() gopurs_runtime.Value {
	once_Node_Stream_read.Do(func() {
		cache_Node_Stream_read = gopurs_runtime.Func(func(r_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_read(r_0_box)
		})
	})
	return cache_Node_Stream_read
}

var cache_Node_Stream_readString gopurs_runtime.Value
var once_Node_Stream_readString sync.Once

func Get_Node_Stream_readString() gopurs_runtime.Value {
	once_Node_Stream_readString.Do(func() {
		cache_Node_Stream_readString = gopurs_runtime.Func2(func(r_0_box gopurs_runtime.Value, enc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_readString(r_0_box, uint32(enc_1_box.IntVal))
		})
	})
	return cache_Node_Stream_readString
}

var cache_Node_Stream_pipeline gopurs_runtime.Value
var once_Node_Stream_pipeline sync.Once

func Get_Node_Stream_pipeline() gopurs_runtime.Value {
	once_Node_Stream_pipeline.Do(func() {
		cache_Node_Stream_pipeline = gopurs_runtime.Func4(func(src_0_box gopurs_runtime.Value, transforms_1_box gopurs_runtime.Value, dest_2_box gopurs_runtime.Value, cb_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_pipeline(src_0_box, *((*[]gopurs_runtime.Value)(transforms_1_box.UnsafePtr)), dest_2_box, cb_3_box)
		})
	})
	return cache_Node_Stream_pipeline
}

var cache_Node_Stream_pipeH gopurs_runtime.Value
var once_Node_Stream_pipeH sync.Once

func Get_Node_Stream_pipeH() gopurs_runtime.Value {
	once_Node_Stream_pipeH.Do(func() {
		cache_Node_Stream_pipeH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "pipe", Get_Effect_Uncurried_mkEffectFn1()}))}
	})
	return cache_Node_Stream_pipeH
}

var cache_Node_Stream_pipe_prime_ gopurs_runtime.Value
var once_Node_Stream_pipe_prime_ sync.Once

func Get_Node_Stream_pipe_prime_() gopurs_runtime.Value {
	once_Node_Stream_pipe_prime_.Do(func() {
		cache_Node_Stream_pipe_prime_ = gopurs_runtime.Func3(func(r_0_box gopurs_runtime.Value, w_1_box gopurs_runtime.Value, o_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_pipe_prime_(r_0_box, w_1_box, func() struct {
				end bool
				/*gopurs:record-v1:1:0065006e0064*/
			} {
				orig := o_2_box
				_ = orig
				clone := struct {
					end bool
					/*gopurs:record-v1:1:0065006e0064*/
				}{}
				clone.end = (gopurs_runtime.RecordGet(orig, "end").IntVal) != (0)
				return clone
			}())
		})
	})
	return cache_Node_Stream_pipe_prime_
}

var cache_Node_Stream_pipe gopurs_runtime.Value
var once_Node_Stream_pipe sync.Once

func Get_Node_Stream_pipe() gopurs_runtime.Value {
	once_Node_Stream_pipe.Do(func() {
		cache_Node_Stream_pipe = gopurs_runtime.Func2(func(r_0_box gopurs_runtime.Value, w_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_pipe(r_0_box, w_1_box)
		})
	})
	return cache_Node_Stream_pipe
}

var cache_Node_Stream_pauseH gopurs_runtime.Value
var once_Node_Stream_pauseH sync.Once

func Get_Node_Stream_pauseH() gopurs_runtime.Value {
	once_Node_Stream_pauseH.Do(func() {
		cache_Node_Stream_pauseH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "pause", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Node_Stream_pauseH
}

var cache_Node_Stream_pause gopurs_runtime.Value
var once_Node_Stream_pause sync.Once

func Get_Node_Stream_pause() gopurs_runtime.Value {
	once_Node_Stream_pause.Do(func() {
		cache_Node_Stream_pause = gopurs_runtime.Func(func(r_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_pause(r_0_box)
		})
	})
	return cache_Node_Stream_pause
}

var cache_Node_Stream_isPaused gopurs_runtime.Value
var once_Node_Stream_isPaused sync.Once

func Get_Node_Stream_isPaused() gopurs_runtime.Value {
	once_Node_Stream_isPaused.Do(func() {
		cache_Node_Stream_isPaused = gopurs_runtime.Func(func(r_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_isPaused(r_0_box)
		})
	})
	return cache_Node_Stream_isPaused
}

var cache_Node_Stream_finishH gopurs_runtime.Value
var once_Node_Stream_finishH sync.Once

func Get_Node_Stream_finishH() gopurs_runtime.Value {
	once_Node_Stream_finishH.Do(func() {
		cache_Node_Stream_finishH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "finish", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Node_Stream_finishH
}

var cache_Node_Stream_errored gopurs_runtime.Value
var once_Node_Stream_errored sync.Once

func Get_Node_Stream_errored() gopurs_runtime.Value {
	once_Node_Stream_errored.Do(func() {
		cache_Node_Stream_errored = gopurs_runtime.Func(func(rw_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_errored(rw_0_box)
		})
	})
	return cache_Node_Stream_errored
}

var cache_Node_Stream_errorH gopurs_runtime.Value
var once_Node_Stream_errorH sync.Once

func Get_Node_Stream_errorH() gopurs_runtime.Value {
	once_Node_Stream_errorH.Do(func() {
		cache_Node_Stream_errorH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "error", Get_Effect_Uncurried_mkEffectFn1()}))}
	})
	return cache_Node_Stream_errorH
}

var cache_Node_Stream_endH gopurs_runtime.Value
var once_Node_Stream_endH sync.Once

func Get_Node_Stream_endH() gopurs_runtime.Value {
	once_Node_Stream_endH.Do(func() {
		cache_Node_Stream_endH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "end", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Node_Stream_endH
}

var cache_Node_Stream_end_prime_ gopurs_runtime.Value
var once_Node_Stream_end_prime_ sync.Once

func Get_Node_Stream_end_prime_() gopurs_runtime.Value {
	once_Node_Stream_end_prime_.Do(func() {
		cache_Node_Stream_end_prime_ = gopurs_runtime.Func2(func(w_0_box gopurs_runtime.Value, cb_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_end_prime_(w_0_box, cb_1_box)
		})
	})
	return cache_Node_Stream_end_prime_
}

var cache_Node_Stream_end gopurs_runtime.Value
var once_Node_Stream_end sync.Once

func Get_Node_Stream_end() gopurs_runtime.Value {
	once_Node_Stream_end.Do(func() {
		cache_Node_Stream_end = gopurs_runtime.Func(func(w_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_end(w_0_box)
		})
	})
	return cache_Node_Stream_end
}

var cache_Node_Stream_drainH gopurs_runtime.Value
var once_Node_Stream_drainH sync.Once

func Get_Node_Stream_drainH() gopurs_runtime.Value {
	once_Node_Stream_drainH.Do(func() {
		cache_Node_Stream_drainH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "drain", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Node_Stream_drainH
}

var cache_Node_Stream_destroyed gopurs_runtime.Value
var once_Node_Stream_destroyed sync.Once

func Get_Node_Stream_destroyed() gopurs_runtime.Value {
	once_Node_Stream_destroyed.Do(func() {
		cache_Node_Stream_destroyed = gopurs_runtime.Func(func(w_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_destroyed(w_0_box)
		})
	})
	return cache_Node_Stream_destroyed
}

var cache_Node_Stream_destroy_prime_ gopurs_runtime.Value
var once_Node_Stream_destroy_prime_ sync.Once

func Get_Node_Stream_destroy_prime_() gopurs_runtime.Value {
	once_Node_Stream_destroy_prime_.Do(func() {
		cache_Node_Stream_destroy_prime_ = gopurs_runtime.Func2(func(w_0_box gopurs_runtime.Value, e_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_destroy_prime_(w_0_box, e_1_box)
		})
	})
	return cache_Node_Stream_destroy_prime_
}

var cache_Node_Stream_destroy gopurs_runtime.Value
var once_Node_Stream_destroy sync.Once

func Get_Node_Stream_destroy() gopurs_runtime.Value {
	once_Node_Stream_destroy.Do(func() {
		cache_Node_Stream_destroy = gopurs_runtime.Func(func(w_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_destroy(w_0_box)
		})
	})
	return cache_Node_Stream_destroy
}

var cache_Node_Stream_dataHStr gopurs_runtime.Value
var once_Node_Stream_dataHStr sync.Once

func Get_Node_Stream_dataHStr() gopurs_runtime.Value {
	once_Node_Stream_dataHStr.Do(func() {
		cache_Node_Stream_dataHStr = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "data", gopurs_runtime.Func(func(cb_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(chunk_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.UncurriedApp3(Get_Node_Stream_readChunkImpl(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Exception_throwException(), Get_Effect_Exception_error(), gopurs_runtime.StrWithCodeUnits("Got a Buffer, not String. Stream encoding must be set to get a String.", 70)), gopurs_runtime.Value{})
					}), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply(cb_0, __local_var_2), gopurs_runtime.Value{})
					}), chunk_1)
				}), gopurs_runtime.Value{})
			})
		})}))}
	})
	return cache_Node_Stream_dataHStr
}

var cache_Node_Stream_dataHEither gopurs_runtime.Value
var once_Node_Stream_dataHEither sync.Once

func Get_Node_Stream_dataHEither() gopurs_runtime.Value {
	once_Node_Stream_dataHEither.Do(func() {
		cache_Node_Stream_dataHEither = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "data", gopurs_runtime.Func(func(cb_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(chunk_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.UncurriedApp3(Get_Node_Stream_readChunkImpl(), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), cb_0, Get_Data_Either_Left(), __local_var_2), gopurs_runtime.Value{})
					}), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), cb_0, Get_Data_Either_Right(), __local_var_2), gopurs_runtime.Value{})
					}), chunk_1)
				}), gopurs_runtime.Value{})
			})
		})}))}
	})
	return cache_Node_Stream_dataHEither
}

var cache_Node_Stream_dataH gopurs_runtime.Value
var once_Node_Stream_dataH sync.Once

func Get_Node_Stream_dataH() gopurs_runtime.Value {
	once_Node_Stream_dataH.Do(func() {
		cache_Node_Stream_dataH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "data", gopurs_runtime.Func(func(cb_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(chunk_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.UncurriedApp3(Get_Node_Stream_readChunkImpl(), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply(cb_0, __local_var_2), gopurs_runtime.Value{})
					}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Exception_throwException(), Get_Effect_Exception_error(), gopurs_runtime.StrWithCodeUnits("Got a String, not a Buffer. Stream encoding should not be set", 61)), gopurs_runtime.Value{})
					}), chunk_1)
				}), gopurs_runtime.Value{})
			})
		})}))}
	})
	return cache_Node_Stream_dataH
}

var cache_Node_Stream_cork gopurs_runtime.Value
var once_Node_Stream_cork sync.Once

func Get_Node_Stream_cork() gopurs_runtime.Value {
	once_Node_Stream_cork.Do(func() {
		cache_Node_Stream_cork = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_cork(s_0_box)
		})
	})
	return cache_Node_Stream_cork
}

var cache_Node_Stream_closed gopurs_runtime.Value
var once_Node_Stream_closed sync.Once

func Get_Node_Stream_closed() gopurs_runtime.Value {
	once_Node_Stream_closed.Do(func() {
		cache_Node_Stream_closed = gopurs_runtime.Func(func(w_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_closed(w_0_box)
		})
	})
	return cache_Node_Stream_closed
}

var cache_Node_Stream_closeH gopurs_runtime.Value
var once_Node_Stream_closeH sync.Once

func Get_Node_Stream_closeH() gopurs_runtime.Value {
	once_Node_Stream_closeH.Do(func() {
		cache_Node_Stream_closeH = gopurs_runtime.Value{Type: 9, IntVal: 4152640982, UnsafePtr: unsafe.Pointer((&Constructor_Node_EventEmitter_EventHandle[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, "close", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Node_Stream_closeH
}

var cache_Node_Stream_allowHalfOpen gopurs_runtime.Value
var once_Node_Stream_allowHalfOpen sync.Once

func Get_Node_Stream_allowHalfOpen() gopurs_runtime.Value {
	once_Node_Stream_allowHalfOpen.Do(func() {
		cache_Node_Stream_allowHalfOpen = gopurs_runtime.Func(func(d_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Stream_allowHalfOpen(d_0_box)
		})
	})
	return cache_Node_Stream_allowHalfOpen
}

func Call_Node_Stream_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Node_Stream_writeableNeedDrain(w_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_writeableNeedDrainImpl(), w_0)
	})
}

func Call_Node_Stream_writeableLength(w_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_writeableLengthImpl(), w_0)
	})
}

func Call_Node_Stream_writeableHighWaterMark(w_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_writeableHighWaterMarkImpl(), w_0)
	})
}

func Call_Node_Stream_writeableFinished(w_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_writeableFinishedImpl(), w_0)
	})
}

func Call_Node_Stream_writeableEnded(w_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_writeableEndedImpl(), w_0)
	})
}

func Call_Node_Stream_writeableCorked(w_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_writeableCorkedImpl(), w_0)
	})
}

func Call_Node_Stream_writeable(w_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_writeableImpl(), w_0)
	})
}

func Call_Node_Stream_writeString_prime_(w_0_loop gopurs_runtime.Value, enc_1_loop uint32, str_2_loop string, cb_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	var enc_1 uint32 = enc_1_loop
	_ = enc_1
	var str_2 string = str_2_loop
	_ = str_2
	var cb_3 gopurs_runtime.Value = cb_3_loop
	_ = cb_3
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 string
		{
			if enc_1 == 2598402200 {
				__t1 = "ascii"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 346813718 {
				__t1 = "utf8"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2352678688 {
				__t1 = "utf16le"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 4287456318 {
				__t1 = "ucs2"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2405785054 {
				__t1 = "base64"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2476550165 {
				__t1 = "base64url"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 3576003686 {
				__t1 = "latin1"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 349879974 {
				__t1 = "binary"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 106823740 {
				__t1 = "hex"
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
		}
	end_branch_1:
		// TAST (Let): __local_var_4_0 shape=Branch(LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, def=Other) bindingType=Any
		__local_var_4_0 := __t1
		_ = __local_var_4_0
		return gopurs_runtime.UncurriedApp4(Get_Node_Stream_writeStringCbImpl(), w_0, gopurs_runtime.Str(str_2), gopurs_runtime.Str(__local_var_4_0), gopurs_runtime.Func(func(err_5 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(cb_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), err_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}), gopurs_runtime.Value{})
		}))
	})
}

func Call_Node_Stream_writeString(w_0_loop gopurs_runtime.Value, enc_1_loop uint32, str_2_loop string) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	var enc_1 uint32 = enc_1_loop
	_ = enc_1
	var str_2 string = str_2_loop
	_ = str_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 string
		{
			if enc_1 == 2598402200 {
				__t1 = "ascii"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 346813718 {
				__t1 = "utf8"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2352678688 {
				__t1 = "utf16le"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 4287456318 {
				__t1 = "ucs2"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2405785054 {
				__t1 = "base64"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2476550165 {
				__t1 = "base64url"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 3576003686 {
				__t1 = "latin1"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 349879974 {
				__t1 = "binary"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 106823740 {
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
		return gopurs_runtime.UncurriedApp3(Get_Node_Stream_writeStringImpl(), w_0, gopurs_runtime.Str(str_2), gopurs_runtime.Str(__local_var_3_0))
	})
}

func Call_Node_Stream_write_prime_(w_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, cb_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	var b_1 gopurs_runtime.Value = b_1_loop
	_ = b_1
	var cb_2 gopurs_runtime.Value = cb_2_loop
	_ = cb_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp3(Get_Node_Stream_writeCbImpl(), w_0, b_1, gopurs_runtime.Func(func(err_3 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(cb_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), err_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}), gopurs_runtime.Value{})
		}))
	})
}

func Call_Node_Stream_write(w_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	var b_1 gopurs_runtime.Value = b_1_loop
	_ = b_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp2(Get_Node_Stream_writeImpl(), w_0, b_1)
	})
}

func Call_Node_Stream_unpipeAll(r_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_unpipeAllImpl(), r_0)
	})
}

func Call_Node_Stream_unpipe(r_0_loop gopurs_runtime.Value, w_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var w_1 gopurs_runtime.Value = w_1_loop
	_ = w_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp2(Get_Node_Stream_unpipeImpl(), r_0, w_1)
	})
}

func Call_Node_Stream_uncork(w_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_uncorkImpl(), w_0)
	})
}

func Call_Node_Stream_setEncoding(r_0_loop gopurs_runtime.Value, enc_1_loop uint32) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var enc_1 uint32 = enc_1_loop
	_ = enc_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 string
		{
			if enc_1 == 2598402200 {
				__t1 = "ASCII"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 346813718 {
				__t1 = "UTF8"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2352678688 {
				__t1 = "UTF16LE"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 4287456318 {
				__t1 = "UCS2"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2405785054 {
				__t1 = "Base64"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2476550165 {
				__t1 = "Base64Url"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 3576003686 {
				__t1 = "Latin1"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 349879974 {
				__t1 = "Binary"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 106823740 {
				__t1 = "Hex"
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
		}
	end_branch_1:
		// TAST (Let): __local_var_2_0 shape=Branch(LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, def=Other) bindingType=Any
		__local_var_2_0 := __t1
		_ = __local_var_2_0
		return gopurs_runtime.UncurriedApp2(Get_Node_Stream_setEncodingImpl(), r_0, gopurs_runtime.Str(__local_var_2_0))
	})
}

func Call_Node_Stream_setDefaultEncoding(r_0_loop gopurs_runtime.Value, enc_1_loop uint32) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var enc_1 uint32 = enc_1_loop
	_ = enc_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 string
		{
			if enc_1 == 2598402200 {
				__t1 = "ASCII"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 346813718 {
				__t1 = "UTF8"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2352678688 {
				__t1 = "UTF16LE"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 4287456318 {
				__t1 = "UCS2"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2405785054 {
				__t1 = "Base64"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2476550165 {
				__t1 = "Base64Url"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 3576003686 {
				__t1 = "Latin1"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 349879974 {
				__t1 = "Binary"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 106823740 {
				__t1 = "Hex"
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
		}
	end_branch_1:
		// TAST (Let): __local_var_2_0 shape=Branch(LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, def=Other) bindingType=Any
		__local_var_2_0 := __t1
		_ = __local_var_2_0
		return gopurs_runtime.UncurriedApp2(Get_Node_Stream_setDefaultEncodingImpl(), r_0, gopurs_runtime.Str(__local_var_2_0))
	})
}

func Call_Node_Stream_resume(r_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_resumeImpl(), r_0)
	})
}

func Call_Node_Stream_readableLength(r_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_readableLengthImpl(), r_0)
	})
}

func Call_Node_Stream_readableHighWaterMark(r_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_readableHighWaterMarkImpl(), r_0)
	})
}

func Call_Node_Stream_readableFromString(str_0_loop string, enc_1_loop uint32) gopurs_runtime.Value {
	var str_0 string = str_0_loop
	_ = str_0
	var enc_1 uint32 = enc_1_loop
	_ = enc_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		var __t1 string
		{
			if enc_1 == 2598402200 {
				__t1 = "ascii"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 346813718 {
				__t1 = "utf8"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2352678688 {
				__t1 = "utf16le"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 4287456318 {
				__t1 = "ucs2"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2405785054 {
				__t1 = "base64"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 2476550165 {
				__t1 = "base64url"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 3576003686 {
				__t1 = "latin1"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 349879974 {
				__t1 = "binary"
				goto end_branch_1
			} else {

			}
		}
		{
			if enc_1 == 106823740 {
				__t1 = "hex"
				goto end_branch_1
			} else {

			}
		}
		{
			__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
		}
	end_branch_1:
		// TAST (Let): __local_var_2_0 shape=Branch(LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, LitString, def=Other) bindingType=Any
		__local_var_2_0 := __t1
		_ = __local_var_2_0
		return gopurs_runtime.UncurriedApp2(Get_Node_Stream_readableFromStrImpl(), gopurs_runtime.Str(str_0), gopurs_runtime.Str(__local_var_2_0))
	})
}

func Call_Node_Stream_readableFromBuffer(buf_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var buf_0 gopurs_runtime.Value = buf_0_loop
	_ = buf_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_readableFromBufImpl(), buf_0)
	})
}

func Call_Node_Stream_readableFlowing(r_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_readableFlowingImpl(), r_0)
	})
}

func Call_Node_Stream_readableEnded(r_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_readableEndedImpl(), r_0)
	})
}

func Call_Node_Stream_readable(r_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_readableImpl(), r_0)
	})
}

func Call_Node_Stream_readEither_prime_(r_0_loop gopurs_runtime.Value, size_1_loop int64) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var size_1 int64 = size_1_loop
	_ = size_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_2_0 := gopurs_runtime.UncurriedApp2(Get_Node_Stream_readSizeImpl(), r_0, gopurs_runtime.Int(size_1))
		_ = __local_var_2_0
		// TAST (Let): v_3_1 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Node","Stream","Chunk"] [])])
		v_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), __local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just()))
		_ = v_3_1
		var __t2 gopurs_runtime.Value
		{
			if v_3_1 == nil {
				__t2 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				})
				goto end_branch_2
			} else {

			}
		}
		{
			if v_3_1 != nil {
				__t2 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.UncurriedApp3(Get_Node_Stream_readChunkImpl(), gopurs_runtime.Func(func(__local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_pureE(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), Get_Data_Either_Right()), __local_var_4), gopurs_runtime.Value{})
					}), gopurs_runtime.Func(func(__local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_pureE(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), Get_Data_Either_Left()), __local_var_4), gopurs_runtime.Value{})
					}), (v_3_1).V0)
				})
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_2:
		return gopurs_runtime.Apply(__t2, gopurs_runtime.Value{})
	})
}

func Call_Node_Stream_readEither(r_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_1_0 := gopurs_runtime.UncurriedApp(Get_Node_Stream_readImpl(), r_0)
		_ = __local_var_1_0
		// TAST (Let): v_2_1 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Node","Stream","Chunk"] [])])
		v_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), __local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just()))
		_ = v_2_1
		var __t2 gopurs_runtime.Value
		{
			if v_2_1 == nil {
				__t2 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				})
				goto end_branch_2
			} else {

			}
		}
		{
			if v_2_1 != nil {
				__t2 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.UncurriedApp3(Get_Node_Stream_readChunkImpl(), gopurs_runtime.Func(func(__local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_pureE(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), Get_Data_Either_Right()), __local_var_3), gopurs_runtime.Value{})
					}), gopurs_runtime.Func(func(__local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_pureE(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), Get_Data_Either_Left()), __local_var_3), gopurs_runtime.Value{})
					}), (v_2_1).V0)
				})
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_2:
		return gopurs_runtime.Apply(__t2, gopurs_runtime.Value{})
	})
}

func Call_Node_Stream_read_prime_(r_0_loop gopurs_runtime.Value, size_1_loop int64) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var size_1 int64 = size_1_loop
	_ = size_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_2_0 := gopurs_runtime.UncurriedApp2(Get_Node_Stream_readSizeImpl(), r_0, gopurs_runtime.Int(size_1))
		_ = __local_var_2_0
		// TAST (Let): v_3_1 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Node","Stream","Chunk"] [])])
		v_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), __local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just()))
		_ = v_3_1
		var __t2 gopurs_runtime.Value
		{
			if v_3_1 == nil {
				__t2 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				})
				goto end_branch_2
			} else {

			}
		}
		{
			if v_3_1 != nil {
				__t2 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.UncurriedApp3(Get_Node_Stream_readChunkImpl(), gopurs_runtime.Func(func(buf_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, buf_4}))}
						}), gopurs_runtime.Value{})
					}), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Exception_throwException(), Get_Effect_Exception_error(), gopurs_runtime.StrWithCodeUnits("Stream encoding should not be set", 33)), gopurs_runtime.Value{})
					}), (v_3_1).V0)
				})
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_2:
		return gopurs_runtime.Apply(__t2, gopurs_runtime.Value{})
	})
}

func Call_Node_Stream_readString_prime_(r_0_loop gopurs_runtime.Value, size_1_loop int64, enc_2_loop uint32) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var size_1 int64 = size_1_loop
	_ = size_1
	var enc_2 uint32 = enc_2_loop
	_ = enc_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
		__local_var_3_0 := gopurs_runtime.Apply2(Get_Node_Stream_read_prime_(), r_0, gopurs_runtime.Int(size_1))
		_ = __local_var_3_0
		__local_var_4_1 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
		_ = __local_var_4_1
		var __t2 gopurs_runtime.Value
		{
			if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 930809136 && __local_var_4_1.UnsafePtr == nil {
				__t2 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				})
				goto end_branch_2
			} else {

			}
		}
		{
			if __local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 930809136 && __local_var_4_1.UnsafePtr != nil {
				__t2 = Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Data_Maybe_Just(), Call_Node_Buffer_usingFromImmutable(gopurs_runtime.Apply(Get_Node_Buffer_Immutable_toString(), gopurs_runtime.Value{Type: 9, IntVal: int64(enc_2), UnsafePtr: nil}), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0))
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_2:
		return gopurs_runtime.Apply(__t2, gopurs_runtime.Value{})
	})
}

func Call_Node_Stream_read(r_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_1_0 := gopurs_runtime.UncurriedApp(Get_Node_Stream_readImpl(), r_0)
		_ = __local_var_1_0
		// TAST (Let): v_2_1 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Node","Stream","Chunk"] [])])
		v_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), __local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just()))
		_ = v_2_1
		var __t2 gopurs_runtime.Value
		{
			if v_2_1 == nil {
				__t2 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				})
				goto end_branch_2
			} else {

			}
		}
		{
			if v_2_1 != nil {
				__t2 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.UncurriedApp3(Get_Node_Stream_readChunkImpl(), gopurs_runtime.Func(func(buf_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, buf_3}))}
						}), gopurs_runtime.Value{})
					}), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Exception_throwException(), Get_Effect_Exception_error(), gopurs_runtime.StrWithCodeUnits("Stream encoding should not be set", 33)), gopurs_runtime.Value{})
					}), (v_2_1).V0)
				})
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_2:
		return gopurs_runtime.Apply(__t2, gopurs_runtime.Value{})
	})
}

func Call_Node_Stream_readString(r_0_loop gopurs_runtime.Value, enc_1_loop uint32) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var enc_1 uint32 = enc_1_loop
	_ = enc_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
		__local_var_2_0 := Call_Node_Stream_read(r_0)
		_ = __local_var_2_0
		__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
		_ = __local_var_3_1
		var __t2 gopurs_runtime.Value
		{
			if __local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 930809136 && __local_var_3_1.UnsafePtr == nil {
				__t2 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
				})
				goto end_branch_2
			} else {

			}
		}
		{
			if __local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 930809136 && __local_var_3_1.UnsafePtr != nil {
				__t2 = Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Data_Maybe_Just(), Call_Node_Buffer_usingFromImmutable(gopurs_runtime.Apply(Get_Node_Buffer_Immutable_toString(), gopurs_runtime.Value{Type: 9, IntVal: int64(enc_1), UnsafePtr: nil}), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_1.UnsafePtr).V0))
				goto end_branch_2
			} else {

			}
		}
		{
			__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
		}
	end_branch_2:
		return gopurs_runtime.Apply(__t2, gopurs_runtime.Value{})
	})
}

func Call_Node_Stream_pipeline(src_0_loop gopurs_runtime.Value, transforms_1_loop []gopurs_runtime.Value, dest_2_loop gopurs_runtime.Value, cb_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var src_0 gopurs_runtime.Value = src_0_loop
	_ = src_0
	var transforms_1 []gopurs_runtime.Value = transforms_1_loop
	_ = transforms_1
	var dest_2 gopurs_runtime.Value = dest_2_loop
	_ = dest_2
	var cb_3 gopurs_runtime.Value = cb_3_loop
	_ = cb_3
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp4(Get_Node_Stream_pipelineImpl(), src_0, gopurs_runtime.Array(transforms_1), dest_2, gopurs_runtime.Func(func(err_4 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(cb_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), err_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}), gopurs_runtime.Value{})
		}))
	})
}

func Call_Node_Stream_pipe_prime_(r_0_loop gopurs_runtime.Value, w_1_loop gopurs_runtime.Value, o_2_loop struct {
	end bool
	/*gopurs:record-v1:1:0065006e0064*/
}) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var w_1 gopurs_runtime.Value = w_1_loop
	_ = w_1
	var o_2 struct {
		end bool
		/*gopurs:record-v1:1:0065006e0064*/
	} = o_2_loop
	_ = o_2
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp3(Get_Node_Stream_pipeCbImpl(), r_0, w_1, func() gopurs_runtime.Value {
			orig := o_2
			_ = orig
			return gopurs_runtime.RecordDict1("end", gopurs_runtime.Bool(orig.end))
		}())
	})
}

func Call_Node_Stream_pipe(r_0_loop gopurs_runtime.Value, w_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	var w_1 gopurs_runtime.Value = w_1_loop
	_ = w_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp2(Get_Node_Stream_pipeImpl(), r_0, w_1)
	})
}

func Call_Node_Stream_pause(r_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_pauseImpl(), r_0)
	})
}

func Call_Node_Stream_isPaused(r_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var r_0 gopurs_runtime.Value = r_0_loop
	_ = r_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_isPausedImpl(), r_0)
	})
}

func Call_Node_Stream_errored(rw_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var rw_0 gopurs_runtime.Value = rw_0_loop
	_ = rw_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_erroredImpl(), rw_0)
	})
}

func Call_Node_Stream_end_prime_(w_0_loop gopurs_runtime.Value, cb_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	var cb_1 gopurs_runtime.Value = cb_1_loop
	_ = cb_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp2(Get_Node_Stream_endCbImpl(), w_0, gopurs_runtime.Func(func(err_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(cb_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Nullable_nullable(), err_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, Get_Data_Maybe_Just())))}), gopurs_runtime.Value{})
		}))
	})
}

func Call_Node_Stream_end(w_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_endImpl(), w_0)
	})
}

func Call_Node_Stream_destroyed(w_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_destroyedImpl(), w_0)
	})
}

func Call_Node_Stream_destroy_prime_(w_0_loop gopurs_runtime.Value, e_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	var e_1 gopurs_runtime.Value = e_1_loop
	_ = e_1
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp2(Get_Node_Stream_destroyErrorImpl(), w_0, e_1)
	})
}

func Call_Node_Stream_destroy(w_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_destroyImpl(), w_0)
	})
}

func Call_Node_Stream_cork(s_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var s_0 gopurs_runtime.Value = s_0_loop
	_ = s_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_corkImpl(), s_0)
	})
}

func Call_Node_Stream_closed(w_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var w_0 gopurs_runtime.Value = w_0_loop
	_ = w_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_closedImpl(), w_0)
	})
}

func Call_Node_Stream_allowHalfOpen(d_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var d_0 gopurs_runtime.Value = d_0_loop
	_ = d_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Stream_allowHalfOpenImpl(), d_0)
	})
}

func Get_Node_Stream_allowHalfOpenImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_AllowHalfOpenImpl
}

func Get_Node_Stream_closedImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_ClosedImpl
}

func Get_Node_Stream_corkImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_CorkImpl
}

func Get_Node_Stream_destroyErrorImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_DestroyErrorImpl
}

func Get_Node_Stream_destroyImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_DestroyImpl
}

func Get_Node_Stream_destroyedImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_DestroyedImpl
}

func Get_Node_Stream_endCbImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_EndCbImpl
}

func Get_Node_Stream_endImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_EndImpl
}

func Get_Node_Stream_erroredImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_ErroredImpl
}

func Get_Node_Stream_isPausedImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_IsPausedImpl
}

func Get_Node_Stream_newPassThrough() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_NewPassThrough
}

func Get_Node_Stream_pauseImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_PauseImpl
}

func Get_Node_Stream_pipeCbImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_PipeCbImpl
}

func Get_Node_Stream_pipeImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_PipeImpl
}

func Get_Node_Stream_pipelineImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_PipelineImpl
}

func Get_Node_Stream_readChunkImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_ReadChunkImpl
}

func Get_Node_Stream_readImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_ReadImpl
}

func Get_Node_Stream_readSizeImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_ReadSizeImpl
}

func Get_Node_Stream_readableEndedImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_ReadableEndedImpl
}

func Get_Node_Stream_readableFlowingImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_ReadableFlowingImpl
}

func Get_Node_Stream_readableFromBufImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_ReadableFromBufImpl
}

func Get_Node_Stream_readableFromStrImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_ReadableFromStrImpl
}

func Get_Node_Stream_readableHighWaterMarkImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_ReadableHighWaterMarkImpl
}

func Get_Node_Stream_readableImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_ReadableImpl
}

func Get_Node_Stream_readableLengthImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_ReadableLengthImpl
}

func Get_Node_Stream_resumeImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_ResumeImpl
}

func Get_Node_Stream_setDefaultEncodingImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_SetDefaultEncodingImpl
}

func Get_Node_Stream_setEncodingImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_SetEncodingImpl
}

func Get_Node_Stream_uncorkImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_UncorkImpl
}

func Get_Node_Stream_unpipeAllImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_UnpipeAllImpl
}

func Get_Node_Stream_unpipeImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_UnpipeImpl
}

func Get_Node_Stream_writeCbImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_WriteCbImpl
}

func Get_Node_Stream_writeImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_WriteImpl
}

func Get_Node_Stream_writeStringCbImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_WriteStringCbImpl
}

func Get_Node_Stream_writeStringImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_WriteStringImpl
}

func Get_Node_Stream_writeableCorkedImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_WriteableCorkedImpl
}

func Get_Node_Stream_writeableEndedImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_WriteableEndedImpl
}

func Get_Node_Stream_writeableFinishedImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_WriteableFinishedImpl
}

func Get_Node_Stream_writeableHighWaterMarkImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_WriteableHighWaterMarkImpl
}

func Get_Node_Stream_writeableImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_WriteableImpl
}

func Get_Node_Stream_writeableLengthImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_WriteableLengthImpl
}

func Get_Node_Stream_writeableNeedDrainImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Stream_WriteableNeedDrainImpl
}
