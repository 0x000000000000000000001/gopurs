package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Node_Buffer_ST_bind gopurs_runtime.Value
var once_Node_Buffer_ST_bind sync.Once

func Get_Node_Buffer_ST_bind() gopurs_runtime.Value {
	once_Node_Buffer_ST_bind.Do(func() {
		cache_Node_Buffer_ST_bind = Get_Control_Monad_ST_Internal_bind_()
	})
	return cache_Node_Buffer_ST_bind
}

var cache_Node_Buffer_ST_writeString gopurs_runtime.Value
var once_Node_Buffer_ST_writeString sync.Once

func Get_Node_Buffer_ST_writeString() gopurs_runtime.Value {
	once_Node_Buffer_ST_writeString.Do(func() {
		cache_Node_Buffer_ST_writeString = Get_Node_Buffer_writeString()
	})
	return cache_Node_Buffer_ST_writeString
}

var cache_Node_Buffer_ST_write gopurs_runtime.Value
var once_Node_Buffer_ST_write sync.Once

func Get_Node_Buffer_ST_write() gopurs_runtime.Value {
	once_Node_Buffer_ST_write.Do(func() {
		cache_Node_Buffer_ST_write = Get_Node_Buffer_write()
	})
	return cache_Node_Buffer_ST_write
}

var cache_Node_Buffer_ST_unsafeThaw gopurs_runtime.Value
var once_Node_Buffer_ST_unsafeThaw sync.Once

func Get_Node_Buffer_ST_unsafeThaw() gopurs_runtime.Value {
	once_Node_Buffer_ST_unsafeThaw.Do(func() {
		cache_Node_Buffer_ST_unsafeThaw = Get_Node_Buffer_unsafeThaw()
	})
	return cache_Node_Buffer_ST_unsafeThaw
}

var cache_Node_Buffer_ST_unsafeFreeze gopurs_runtime.Value
var once_Node_Buffer_ST_unsafeFreeze sync.Once

func Get_Node_Buffer_ST_unsafeFreeze() gopurs_runtime.Value {
	once_Node_Buffer_ST_unsafeFreeze.Do(func() {
		cache_Node_Buffer_ST_unsafeFreeze = Get_Node_Buffer_unsafeFreeze()
	})
	return cache_Node_Buffer_ST_unsafeFreeze
}

var cache_Node_Buffer_ST_transcode gopurs_runtime.Value
var once_Node_Buffer_ST_transcode sync.Once

func Get_Node_Buffer_ST_transcode() gopurs_runtime.Value {
	once_Node_Buffer_ST_transcode.Do(func() {
		cache_Node_Buffer_ST_transcode = Get_Node_Buffer_transcode()
	})
	return cache_Node_Buffer_ST_transcode
}

var cache_Node_Buffer_ST_toString_prime_ gopurs_runtime.Value
var once_Node_Buffer_ST_toString_prime_ sync.Once

func Get_Node_Buffer_ST_toString_prime_() gopurs_runtime.Value {
	once_Node_Buffer_ST_toString_prime_.Do(func() {
		cache_Node_Buffer_ST_toString_prime_ = Get_Node_Buffer_toString_prime_()
	})
	return cache_Node_Buffer_ST_toString_prime_
}

var cache_Node_Buffer_ST_toString gopurs_runtime.Value
var once_Node_Buffer_ST_toString sync.Once

func Get_Node_Buffer_ST_toString() gopurs_runtime.Value {
	once_Node_Buffer_ST_toString.Do(func() {
		cache_Node_Buffer_ST_toString = Get_Node_Buffer_toString()
	})
	return cache_Node_Buffer_ST_toString
}

var cache_Node_Buffer_ST_toArrayBuffer gopurs_runtime.Value
var once_Node_Buffer_ST_toArrayBuffer sync.Once

func Get_Node_Buffer_ST_toArrayBuffer() gopurs_runtime.Value {
	once_Node_Buffer_ST_toArrayBuffer.Do(func() {
		cache_Node_Buffer_ST_toArrayBuffer = gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable__3422936011(), Get_Node_Buffer_Immutable_toArrayBuffer())
	})
	return cache_Node_Buffer_ST_toArrayBuffer
}

var cache_Node_Buffer_ST_toArray gopurs_runtime.Value
var once_Node_Buffer_ST_toArray sync.Once

func Get_Node_Buffer_ST_toArray() gopurs_runtime.Value {
	once_Node_Buffer_ST_toArray.Do(func() {
		cache_Node_Buffer_ST_toArray = gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable__3335896427(), Get_Node_Buffer_Immutable_toArray())
	})
	return cache_Node_Buffer_ST_toArray
}

var cache_Node_Buffer_ST_thaw gopurs_runtime.Value
var once_Node_Buffer_ST_thaw sync.Once

func Get_Node_Buffer_ST_thaw() gopurs_runtime.Value {
	once_Node_Buffer_ST_thaw.Do(func() {
		cache_Node_Buffer_ST_thaw = Get_Node_Buffer_thaw()
	})
	return cache_Node_Buffer_ST_thaw
}

var cache_Node_Buffer_ST_swap64 gopurs_runtime.Value
var once_Node_Buffer_ST_swap64 sync.Once

func Get_Node_Buffer_ST_swap64() gopurs_runtime.Value {
	once_Node_Buffer_ST_swap64.Do(func() {
		cache_Node_Buffer_ST_swap64 = Get_Node_Buffer_swap64()
	})
	return cache_Node_Buffer_ST_swap64
}

var cache_Node_Buffer_ST_swap32 gopurs_runtime.Value
var once_Node_Buffer_ST_swap32 sync.Once

func Get_Node_Buffer_ST_swap32() gopurs_runtime.Value {
	once_Node_Buffer_ST_swap32.Do(func() {
		cache_Node_Buffer_ST_swap32 = Get_Node_Buffer_swap32()
	})
	return cache_Node_Buffer_ST_swap32
}

var cache_Node_Buffer_ST_swap16 gopurs_runtime.Value
var once_Node_Buffer_ST_swap16 sync.Once

func Get_Node_Buffer_ST_swap16() gopurs_runtime.Value {
	once_Node_Buffer_ST_swap16.Do(func() {
		cache_Node_Buffer_ST_swap16 = Get_Node_Buffer_swap16()
	})
	return cache_Node_Buffer_ST_swap16
}

var cache_Node_Buffer_ST_slice gopurs_runtime.Value
var once_Node_Buffer_ST_slice sync.Once

func Get_Node_Buffer_ST_slice() gopurs_runtime.Value {
	once_Node_Buffer_ST_slice.Do(func() {
		cache_Node_Buffer_ST_slice = Get_Node_Buffer_Immutable_slice()
	})
	return cache_Node_Buffer_ST_slice
}

var cache_Node_Buffer_ST_size gopurs_runtime.Value
var once_Node_Buffer_ST_size sync.Once

func Get_Node_Buffer_ST_size() gopurs_runtime.Value {
	once_Node_Buffer_ST_size.Do(func() {
		cache_Node_Buffer_ST_size = gopurs_runtime.Apply(Get_Node_Buffer_usingFromImmutable__2070034188(), Get_Node_Buffer_Immutable_size())
	})
	return cache_Node_Buffer_ST_size
}

var cache_Node_Buffer_ST_setAtOffset gopurs_runtime.Value
var once_Node_Buffer_ST_setAtOffset sync.Once

func Get_Node_Buffer_ST_setAtOffset() gopurs_runtime.Value {
	once_Node_Buffer_ST_setAtOffset.Do(func() {
		cache_Node_Buffer_ST_setAtOffset = Get_Node_Buffer_setAtOffset()
	})
	return cache_Node_Buffer_ST_setAtOffset
}

var cache_Node_Buffer_ST_run gopurs_runtime.Value
var once_Node_Buffer_ST_run sync.Once

func Get_Node_Buffer_ST_run() gopurs_runtime.Value {
	once_Node_Buffer_ST_run.Do(func() {
		cache_Node_Buffer_ST_run = gopurs_runtime.Func(func(st_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Buffer_ST_run(st_0_box)
		})
	})
	return cache_Node_Buffer_ST_run
}

var cache_Node_Buffer_ST_readString gopurs_runtime.Value
var once_Node_Buffer_ST_readString sync.Once

func Get_Node_Buffer_ST_readString() gopurs_runtime.Value {
	once_Node_Buffer_ST_readString.Do(func() {
		cache_Node_Buffer_ST_readString = Get_Node_Buffer_readString()
	})
	return cache_Node_Buffer_ST_readString
}

var cache_Node_Buffer_ST_read gopurs_runtime.Value
var once_Node_Buffer_ST_read sync.Once

func Get_Node_Buffer_ST_read() gopurs_runtime.Value {
	once_Node_Buffer_ST_read.Do(func() {
		cache_Node_Buffer_ST_read = Get_Node_Buffer_read()
	})
	return cache_Node_Buffer_ST_read
}

var cache_Node_Buffer_ST_getAtOffset gopurs_runtime.Value
var once_Node_Buffer_ST_getAtOffset sync.Once

func Get_Node_Buffer_ST_getAtOffset() gopurs_runtime.Value {
	once_Node_Buffer_ST_getAtOffset.Do(func() {
		cache_Node_Buffer_ST_getAtOffset = Get_Node_Buffer_getAtOffset()
	})
	return cache_Node_Buffer_ST_getAtOffset
}

var cache_Node_Buffer_ST_fromString gopurs_runtime.Value
var once_Node_Buffer_ST_fromString sync.Once

func Get_Node_Buffer_ST_fromString() gopurs_runtime.Value {
	once_Node_Buffer_ST_fromString.Do(func() {
		cache_Node_Buffer_ST_fromString = Get_Node_Buffer_fromString()
	})
	return cache_Node_Buffer_ST_fromString
}

var cache_Node_Buffer_ST_fromArrayBuffer gopurs_runtime.Value
var once_Node_Buffer_ST_fromArrayBuffer sync.Once

func Get_Node_Buffer_ST_fromArrayBuffer() gopurs_runtime.Value {
	once_Node_Buffer_ST_fromArrayBuffer.Do(func() {
		cache_Node_Buffer_ST_fromArrayBuffer = gopurs_runtime.Apply(Get_Node_Buffer_usingToImmutable__2508462662(), Get_Node_Buffer_Immutable_fromArrayBuffer())
	})
	return cache_Node_Buffer_ST_fromArrayBuffer
}

var cache_Node_Buffer_ST_fromArray gopurs_runtime.Value
var once_Node_Buffer_ST_fromArray sync.Once

func Get_Node_Buffer_ST_fromArray() gopurs_runtime.Value {
	once_Node_Buffer_ST_fromArray.Do(func() {
		cache_Node_Buffer_ST_fromArray = gopurs_runtime.Apply(Get_Node_Buffer_usingToImmutable__3986645350(), Get_Node_Buffer_Immutable_fromArray())
	})
	return cache_Node_Buffer_ST_fromArray
}

var cache_Node_Buffer_ST_freeze gopurs_runtime.Value
var once_Node_Buffer_ST_freeze sync.Once

func Get_Node_Buffer_ST_freeze() gopurs_runtime.Value {
	once_Node_Buffer_ST_freeze.Do(func() {
		cache_Node_Buffer_ST_freeze = Get_Node_Buffer_freeze()
	})
	return cache_Node_Buffer_ST_freeze
}

var cache_Node_Buffer_ST_fill gopurs_runtime.Value
var once_Node_Buffer_ST_fill sync.Once

func Get_Node_Buffer_ST_fill() gopurs_runtime.Value {
	once_Node_Buffer_ST_fill.Do(func() {
		cache_Node_Buffer_ST_fill = Get_Node_Buffer_fill()
	})
	return cache_Node_Buffer_ST_fill
}

var cache_Node_Buffer_ST_copy gopurs_runtime.Value
var once_Node_Buffer_ST_copy sync.Once

func Get_Node_Buffer_ST_copy() gopurs_runtime.Value {
	once_Node_Buffer_ST_copy.Do(func() {
		cache_Node_Buffer_ST_copy = Get_Node_Buffer_copy()
	})
	return cache_Node_Buffer_ST_copy
}

var cache_Node_Buffer_ST_concat_prime_ gopurs_runtime.Value
var once_Node_Buffer_ST_concat_prime_ sync.Once

func Get_Node_Buffer_ST_concat_prime_() gopurs_runtime.Value {
	once_Node_Buffer_ST_concat_prime_.Do(func() {
		cache_Node_Buffer_ST_concat_prime_ = Get_Node_Buffer_concat_prime_()
	})
	return cache_Node_Buffer_ST_concat_prime_
}

var cache_Node_Buffer_ST_concat gopurs_runtime.Value
var once_Node_Buffer_ST_concat sync.Once

func Get_Node_Buffer_ST_concat() gopurs_runtime.Value {
	once_Node_Buffer_ST_concat.Do(func() {
		cache_Node_Buffer_ST_concat = Get_Node_Buffer_concat()
	})
	return cache_Node_Buffer_ST_concat
}

var cache_Node_Buffer_ST_compareParts gopurs_runtime.Value
var once_Node_Buffer_ST_compareParts sync.Once

func Get_Node_Buffer_ST_compareParts() gopurs_runtime.Value {
	once_Node_Buffer_ST_compareParts.Do(func() {
		cache_Node_Buffer_ST_compareParts = Get_Node_Buffer_compareParts()
	})
	return cache_Node_Buffer_ST_compareParts
}

var cache_Node_Buffer_ST_allocUnsafeSlow gopurs_runtime.Value
var once_Node_Buffer_ST_allocUnsafeSlow sync.Once

func Get_Node_Buffer_ST_allocUnsafeSlow() gopurs_runtime.Value {
	once_Node_Buffer_ST_allocUnsafeSlow.Do(func() {
		cache_Node_Buffer_ST_allocUnsafeSlow = Get_Node_Buffer_allocUnsafeSlow()
	})
	return cache_Node_Buffer_ST_allocUnsafeSlow
}

var cache_Node_Buffer_ST_allocUnsafe gopurs_runtime.Value
var once_Node_Buffer_ST_allocUnsafe sync.Once

func Get_Node_Buffer_ST_allocUnsafe() gopurs_runtime.Value {
	once_Node_Buffer_ST_allocUnsafe.Do(func() {
		cache_Node_Buffer_ST_allocUnsafe = Get_Node_Buffer_allocUnsafe()
	})
	return cache_Node_Buffer_ST_allocUnsafe
}

var cache_Node_Buffer_ST_alloc gopurs_runtime.Value
var once_Node_Buffer_ST_alloc sync.Once

func Get_Node_Buffer_ST_alloc() gopurs_runtime.Value {
	once_Node_Buffer_ST_alloc.Do(func() {
		cache_Node_Buffer_ST_alloc = gopurs_runtime.Apply(Get_Node_Buffer_usingToImmutable__4058095588(), Get_Node_Buffer_Immutable_alloc())
	})
	return cache_Node_Buffer_ST_alloc
}

var cache_Node_Buffer_ST_create gopurs_runtime.Value
var once_Node_Buffer_ST_create sync.Once

func Get_Node_Buffer_ST_create() gopurs_runtime.Value {
	once_Node_Buffer_ST_create.Do(func() {
		cache_Node_Buffer_ST_create = Get_Node_Buffer_ST_alloc()
	})
	return cache_Node_Buffer_ST_create
}

var cache_Node_Buffer_ST_mutableBufferST gopurs_runtime.Value
var once_Node_Buffer_ST_mutableBufferST sync.Once

func Get_Node_Buffer_ST_mutableBufferST() gopurs_runtime.Value {
	once_Node_Buffer_ST_mutableBufferST.Do(func() {
		cache_Node_Buffer_ST_mutableBufferST = gopurs_runtime.Value{Type: 9, IntVal: 1692970589, UnsafePtr: unsafe.Pointer((&Constructor_Node_Buffer_Class_MutableBuffer[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Control_Monad_ST_Internal_monadST()
		}), Get_Node_Buffer_ST_concat(), Get_Node_Buffer_ST_concat_prime_(), Get_Node_Buffer_ST_copy(), Get_Node_Buffer_ST_create(), Get_Node_Buffer_ST_fill(), Get_Node_Buffer_ST_freeze(), Get_Node_Buffer_ST_fromArray(), Get_Node_Buffer_ST_fromArrayBuffer(), Get_Node_Buffer_ST_fromString(), Get_Node_Buffer_ST_getAtOffset(), Get_Node_Buffer_ST_read(), Get_Node_Buffer_ST_readString(), Get_Node_Buffer_ST_setAtOffset(), Get_Node_Buffer_ST_size(), Get_Node_Buffer_ST_slice(), Get_Node_Buffer_ST_thaw(), Get_Node_Buffer_ST_toArray(), Get_Node_Buffer_ST_toArrayBuffer(), Get_Node_Buffer_ST_toString(), Get_Node_Buffer_ST_unsafeFreeze(), Get_Node_Buffer_ST_unsafeThaw(), Get_Node_Buffer_ST_write(), Get_Node_Buffer_ST_writeString()}))}
	})
	return cache_Node_Buffer_ST_mutableBufferST
}

func Call_Node_Buffer_ST_run(st_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var st_0 gopurs_runtime.Value = st_0_loop
	_ = st_0
	return gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		__local_var_1_0 := gopurs_runtime.Apply(st_0, gopurs_runtime.Value{})
		_ = __local_var_1_0
		return gopurs_runtime.Apply2(Get_Node_Buffer_ST_unsafeFreeze(), __local_var_1_0, gopurs_runtime.Value{})
	}))
}
