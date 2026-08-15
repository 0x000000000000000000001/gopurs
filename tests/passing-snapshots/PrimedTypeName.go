package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_TP gopurs_runtime.Value
var once_Main_TP sync.Once

func Get_Main_TP() gopurs_runtime.Value {
	once_Main_TP.Do(func() {
		cache_Main_TP = gopurs_runtime.Value{Type: 9, IntVal: int64(3421219130), UnsafePtr: nil}
	})
	return cache_Main_TP
}

var cache_Main_T gopurs_runtime.Value
var once_Main_T sync.Once

func Get_Main_T() gopurs_runtime.Value {
	once_Main_T.Do(func() {
		cache_Main_T = gopurs_runtime.Value{Type: 9, IntVal: int64(990467018), UnsafePtr: nil}
	})
	return cache_Main_T
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_eqT gopurs_runtime.Value
var once_Main_eqT sync.Once

func Get_Main_eqT() gopurs_runtime.Value {
	once_Main_eqT.Do(func() {
		cache_Main_eqT = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool(true)
			})
		})}))}
	})
	return cache_Main_eqT
}

type Constructor_Main_TP struct {
	Rc uint32
}

type Constructor_Main_T struct {
	Rc uint32
}
