package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_UpdateDto gopurs_runtime.Value
var once_Main_UpdateDto sync.Once

func Get_Main_UpdateDto() gopurs_runtime.Value {
	once_Main_UpdateDto.Do(func() {
		cache_Main_UpdateDto = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_UpdateDto(x_0_box)
		})
	})
	return cache_Main_UpdateDto
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_eqUpdateDto gopurs_runtime.Value
var once_Main_eqUpdateDto sync.Once

func Get_Main_eqUpdateDto() gopurs_runtime.Value {
	once_Main_eqUpdateDto.Do(func() {
		cache_Main_eqUpdateDto = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Bool((gopurs_runtime.RecordGet(x_0, "bio").StrVal()) == (gopurs_runtime.RecordGet(y_1, "bio").StrVal()))
			})
		})})}
	})
	return cache_Main_eqUpdateDto
}

func Call_Main_UpdateDto(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}
