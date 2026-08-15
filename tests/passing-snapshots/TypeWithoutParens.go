package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_idY gopurs_runtime.Value
var once_Main_idY sync.Once

func Get_Main_idY() gopurs_runtime.Value {
	once_Main_idY.Do(func() {
		cache_Main_idY = gopurs_runtime.Func(func(y_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_idY(uint32(y_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_idY
}

var cache_Main_idX gopurs_runtime.Value
var once_Main_idX sync.Once

func Get_Main_idX() gopurs_runtime.Value {
	once_Main_idX.Do(func() {
		cache_Main_idX = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_idX(uint32(x_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_idX
}

func Call_Main_idY(y_0_loop uint32) uint32 {
	var y_0 uint32 = y_0_loop
	_ = y_0
	return y_0
}

func Call_Main_idX(x_0_loop uint32) uint32 {
	var x_0 uint32 = x_0_loop
	_ = x_0
	return x_0
}
