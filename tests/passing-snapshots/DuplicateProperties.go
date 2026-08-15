package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_subtractX gopurs_runtime.Value
var once_Main_subtractX sync.Once

func Get_Main_subtractX() gopurs_runtime.Value {
	once_Main_subtractX.Do(func() {
		cache_Main_subtractX = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_subtractX(uint32(v_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_subtractX
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_hasX gopurs_runtime.Value
var once_Main_hasX sync.Once

func Get_Main_hasX() gopurs_runtime.Value {
	once_Main_hasX.Do(func() {
		cache_Main_hasX = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_hasX
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_test1
}

var cache_Main_extractX gopurs_runtime.Value
var once_Main_extractX sync.Once

func Get_Main_extractX() gopurs_runtime.Value {
	once_Main_extractX.Do(func() {
		cache_Main_extractX = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_extractX(uint32(v_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_extractX
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_test2(uint32(x_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_test2
}

var cache_Main_extractX__651711577 gopurs_runtime.Value
var once_Main_extractX__651711577 sync.Once

func Get_Main_extractX__651711577() gopurs_runtime.Value {
	once_Main_extractX__651711577.Do(func() {
		cache_Main_extractX__651711577 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_extractX__651711577(uint32(v_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_extractX__651711577
}

var cache_Main_extractX__3239128985 gopurs_runtime.Value
var once_Main_extractX__3239128985 sync.Once

func Get_Main_extractX__3239128985() gopurs_runtime.Value {
	once_Main_extractX__3239128985.Do(func() {
		cache_Main_extractX__3239128985 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_extractX__3239128985(uint32(v_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_extractX__3239128985
}

var cache_Main_hasX__3736288358 gopurs_runtime.Value
var once_Main_hasX__3736288358 sync.Once

func Get_Main_hasX__3736288358() gopurs_runtime.Value {
	once_Main_hasX__3736288358.Do(func() {
		cache_Main_hasX__3736288358 = gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
	})
	return cache_Main_hasX__3736288358
}

var cache_Main_subtractX__3539893288 gopurs_runtime.Value
var once_Main_subtractX__3539893288 sync.Once

func Get_Main_subtractX__3539893288() gopurs_runtime.Value {
	once_Main_subtractX__3539893288.Do(func() {
		cache_Main_subtractX__3539893288 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_subtractX__3539893288(uint32(v_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_subtractX__3539893288
}

var cache_Main_subtractX__4267503976 gopurs_runtime.Value
var once_Main_subtractX__4267503976 sync.Once

func Get_Main_subtractX__4267503976() gopurs_runtime.Value {
	once_Main_subtractX__4267503976.Do(func() {
		cache_Main_subtractX__4267503976 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_subtractX__4267503976(uint32(v_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_subtractX__4267503976
}

func Call_Main_subtractX(v_0_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	return 513803634
}

func Call_Main_extractX(v_0_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	return 513803634
}

func Call_Main_test2(x_0_loop uint32) uint32 {
	var x_0 uint32 = x_0_loop
	_ = x_0
	return 513803634
}

func Call_Main_extractX__651711577(v_0_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	return 513803634
}

func Call_Main_extractX__3239128985(v_0_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	return 513803634
}

func Call_Main_subtractX__3539893288(v_0_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	return 513803634
}

func Call_Main_subtractX__4267503976(v_0_loop uint32) uint32 {
	var v_0 uint32 = v_0_loop
	_ = v_0
	return 513803634
}
