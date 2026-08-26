package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Get gopurs_runtime.Value
var once_Main_Get sync.Once

func Get_Main_Get() gopurs_runtime.Value {
	once_Main_Get.Do(func() {
		cache_Main_Get = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Get
}

var cache_Main_functorTypedCacheConst gopurs_runtime.Value
var once_Main_functorTypedCacheConst sync.Once

func Get_Main_functorTypedCacheConst() gopurs_runtime.Value {
	once_Main_functorTypedCacheConst.Do(func() {
		cache_Main_functorTypedCacheConst = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[*Constructor_Main_Get[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Value{Type: 9, IntVal: 2813159464, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Get[gopurs_runtime.Value, gopurs_runtime.Value]](m_1))}
			})
		})}))}
	})
	return cache_Main_functorTypedCacheConst
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Get[T_key any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}
