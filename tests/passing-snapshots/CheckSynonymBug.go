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

var cache_Main_length gopurs_runtime.Value
var once_Main_length sync.Once

func Get_Main_length() gopurs_runtime.Value {
	once_Main_length.Do(func() {
		cache_Main_length = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_length(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}()))
		})
	})
	return cache_Main_length
}

var cache_Main_length__4151727363 gopurs_runtime.Value
var once_Main_length__4151727363 sync.Once

func Get_Main_length__4151727363() gopurs_runtime.Value {
	once_Main_length__4151727363.Do(func() {
		cache_Main_length__4151727363 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_length__4151727363(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}()))
		})
	})
	return cache_Main_length__4151727363
}

var cache_Main_foo gopurs_runtime.Value
var once_Main_foo sync.Once

func Get_Main_foo() gopurs_runtime.Value {
	once_Main_foo.Do(func() {
		cache_Main_foo = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_foo(v_0_box))
		})
	})
	return cache_Main_foo
}

func Call_Main_length(v_0_loop []gopurs_runtime.Value) int64 {
	var v_0 []gopurs_runtime.Value = v_0_loop
	_ = v_0
	return 0
}

func Call_Main_length__4151727363(v_0_loop []gopurs_runtime.Value) int64 {
	var v_0 []gopurs_runtime.Value = v_0_loop
	_ = v_0
	return 0
}

func Call_Main_foo(v_0_loop gopurs_runtime.Value) int64 {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return 0
}
