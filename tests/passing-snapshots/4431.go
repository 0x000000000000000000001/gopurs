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
			return gopurs_runtime.Value{Type: 9, IntVal: 2813159464, UnsafePtr: unsafe.Pointer(&Constructor_Main_Get{1, value0})}
		})
	})
	return cache_Main_Get
}

var cache_Main_functorTypedCache gopurs_runtime.Value
var once_Main_functorTypedCache sync.Once

func Get_Main_functorTypedCache() gopurs_runtime.Value {
	once_Main_functorTypedCache.Do(func() {
		cache_Main_functorTypedCache = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_functorTypedCache(dictFunctor_0_box)
		})
	})
	return cache_Main_functorTypedCache
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_Get struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_functorTypedCache(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
	_ = dictFunctor_0
	return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2813159464, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Main_Get](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, m_2)))}
		})
	})})}
}
