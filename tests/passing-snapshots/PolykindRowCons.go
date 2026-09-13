package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_Proxy gopurs_runtime.Value
var once_Main_Proxy sync.Once

func Get_Main_Proxy() gopurs_runtime.Value {
	once_Main_Proxy.Do(func() {
		cache_Main_Proxy = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_Proxy
}

var cache_Main_Identity gopurs_runtime.Value
var once_Main_Identity sync.Once

func Get_Main_Identity() gopurs_runtime.Value {
	once_Main_Identity.Do(func() {
		cache_Main_Identity = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Identity
}

var cache_Main_App gopurs_runtime.Value
var once_Main_App sync.Once

func Get_Main_App() gopurs_runtime.Value {
	once_Main_App.Do(func() {
		cache_Main_App = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_App
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_lookup gopurs_runtime.Value
var once_Main_lookup sync.Once

func Get_Main_lookup() gopurs_runtime.Value {
	once_Main_lookup.Do(func() {
		cache_Main_lookup = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup(_dollar___unused_0_box, uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lookup
}

var cache_Main_lookup__4284712939 gopurs_runtime.Value
var once_Main_lookup__4284712939 sync.Once

func Get_Main_lookup__4284712939() gopurs_runtime.Value {
	once_Main_lookup__4284712939.Do(func() {
		cache_Main_lookup__4284712939 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup__4284712939(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lookup__4284712939
}

var cache_Main_lookup__187476285 gopurs_runtime.Value
var once_Main_lookup__187476285 sync.Once

func Get_Main_lookup__187476285() gopurs_runtime.Value {
	once_Main_lookup__187476285.Do(func() {
		cache_Main_lookup__187476285 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup__187476285(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lookup__187476285
}

var cache_Main_lookup__135669723 gopurs_runtime.Value
var once_Main_lookup__135669723 sync.Once

func Get_Main_lookup__135669723() gopurs_runtime.Value {
	once_Main_lookup__135669723.Do(func() {
		cache_Main_lookup__135669723 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup__135669723(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lookup__135669723
}

var cache_Main_lookup__1301770859 gopurs_runtime.Value
var once_Main_lookup__1301770859 sync.Once

func Get_Main_lookup__1301770859() gopurs_runtime.Value {
	once_Main_lookup__1301770859.Do(func() {
		cache_Main_lookup__1301770859 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup__1301770859(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lookup__1301770859
}

var cache_Main_lookup__3186229469 gopurs_runtime.Value
var once_Main_lookup__3186229469 sync.Once

func Get_Main_lookup__3186229469() gopurs_runtime.Value {
	once_Main_lookup__3186229469.Do(func() {
		cache_Main_lookup__3186229469 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup__3186229469(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lookup__3186229469
}

var cache_Main_lookup__1069724763 gopurs_runtime.Value
var once_Main_lookup__1069724763 sync.Once

func Get_Main_lookup__1069724763() gopurs_runtime.Value {
	once_Main_lookup__1069724763.Do(func() {
		cache_Main_lookup__1069724763 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup__1069724763(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lookup__1069724763
}

var cache_Main_lookup__2898168299 gopurs_runtime.Value
var once_Main_lookup__2898168299 sync.Once

func Get_Main_lookup__2898168299() gopurs_runtime.Value {
	once_Main_lookup__2898168299.Do(func() {
		cache_Main_lookup__2898168299 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup__2898168299(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lookup__2898168299
}

var cache_Main_lookup__2248787069 gopurs_runtime.Value
var once_Main_lookup__2248787069 sync.Once

func Get_Main_lookup__2248787069() gopurs_runtime.Value {
	once_Main_lookup__2248787069.Do(func() {
		cache_Main_lookup__2248787069 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup__2248787069(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lookup__2248787069
}

var cache_Main_lookup__4258804571 gopurs_runtime.Value
var once_Main_lookup__4258804571 sync.Once

func Get_Main_lookup__4258804571() gopurs_runtime.Value {
	once_Main_lookup__4258804571.Do(func() {
		cache_Main_lookup__4258804571 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup__4258804571(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_lookup__4258804571
}

var cache_Main_lookup1 gopurs_runtime.Value
var once_Main_lookup1 sync.Once

func Get_Main_lookup1() gopurs_runtime.Value {
	once_Main_lookup1.Do(func() {
		cache_Main_lookup1 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_lookup1
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Main_lookup1().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_test1
}

var cache_Main_lookup2 gopurs_runtime.Value
var once_Main_lookup2 sync.Once

func Get_Main_lookup2() gopurs_runtime.Value {
	once_Main_lookup2.Do(func() {
		cache_Main_lookup2 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_lookup2
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Main_lookup2().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_test2
}

var cache_Main_lookup3 gopurs_runtime.Value
var once_Main_lookup3 sync.Once

func Get_Main_lookup3() gopurs_runtime.Value {
	once_Main_lookup3.Do(func() {
		cache_Main_lookup3 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_lookup3
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Main_lookup3().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_test3
}

var cache_Main_lookup4 gopurs_runtime.Value
var once_Main_lookup4 sync.Once

func Get_Main_lookup4() gopurs_runtime.Value {
	once_Main_lookup4.Do(func() {
		cache_Main_lookup4 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_lookup4
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Main_lookup4().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_test4
}

var cache_Main_lookup5 gopurs_runtime.Value
var once_Main_lookup5 sync.Once

func Get_Main_lookup5() gopurs_runtime.Value {
	once_Main_lookup5.Do(func() {
		cache_Main_lookup5 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_lookup5
}

var cache_Main_test5 gopurs_runtime.Value
var once_Main_test5 sync.Once

func Get_Main_test5() gopurs_runtime.Value {
	once_Main_test5.Do(func() {
		cache_Main_test5 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Main_lookup5().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_test5
}

var cache_Main_lookup6 gopurs_runtime.Value
var once_Main_lookup6 sync.Once

func Get_Main_lookup6() gopurs_runtime.Value {
	once_Main_lookup6.Do(func() {
		cache_Main_lookup6 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_lookup6
}

var cache_Main_test6 gopurs_runtime.Value
var once_Main_test6 sync.Once

func Get_Main_test6() gopurs_runtime.Value {
	once_Main_test6.Do(func() {
		cache_Main_test6 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Main_lookup6().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_test6
}

var cache_Main_lookup7 gopurs_runtime.Value
var once_Main_lookup7 sync.Once

func Get_Main_lookup7() gopurs_runtime.Value {
	once_Main_lookup7.Do(func() {
		cache_Main_lookup7 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_lookup7
}

var cache_Main_test7 gopurs_runtime.Value
var once_Main_test7 sync.Once

func Get_Main_test7() gopurs_runtime.Value {
	once_Main_test7.Do(func() {
		cache_Main_test7 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Main_lookup7().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_test7
}

var cache_Main_lookup8 gopurs_runtime.Value
var once_Main_lookup8 sync.Once

func Get_Main_lookup8() gopurs_runtime.Value {
	once_Main_lookup8.Do(func() {
		cache_Main_lookup8 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_lookup8
}

var cache_Main_test8 gopurs_runtime.Value
var once_Main_test8 sync.Once

func Get_Main_test8() gopurs_runtime.Value {
	once_Main_test8.Do(func() {
		cache_Main_test8 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Main_lookup8().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_test8
}

var cache_Main_lookup9 gopurs_runtime.Value
var once_Main_lookup9 sync.Once

func Get_Main_lookup9() gopurs_runtime.Value {
	once_Main_lookup9.Do(func() {
		cache_Main_lookup9 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_lookup(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_lookup9
}

var cache_Main_test9 gopurs_runtime.Value
var once_Main_test9 sync.Once

func Get_Main_test9() gopurs_runtime.Value {
	once_Main_test9.Do(func() {
		cache_Main_test9 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Main_lookup9().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_test9
}

type Constructor_Main_Proxy[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Identity[T_a any] struct {
	Rc uint32
	V0 T_a
}

type Constructor_Main_App[T_f any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func Call_Main_lookup(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 227768594
}

func Call_Main_lookup__4284712939(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
lookup__4284712939:
	for {
		if false {
			continue lookup__4284712939
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}

func Call_Main_lookup__187476285(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
lookup__187476285:
	for {
		if false {
			continue lookup__187476285
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}

func Call_Main_lookup__135669723(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
lookup__135669723:
	for {
		if false {
			continue lookup__135669723
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}

func Call_Main_lookup__1301770859(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
lookup__1301770859:
	for {
		if false {
			continue lookup__1301770859
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}

func Call_Main_lookup__3186229469(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
lookup__3186229469:
	for {
		if false {
			continue lookup__3186229469
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}

func Call_Main_lookup__1069724763(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
lookup__1069724763:
	for {
		if false {
			continue lookup__1069724763
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}

func Call_Main_lookup__2898168299(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
lookup__2898168299:
	for {
		if false {
			continue lookup__2898168299
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}

func Call_Main_lookup__2248787069(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
lookup__2248787069:
	for {
		if false {
			continue lookup__2248787069
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}

func Call_Main_lookup__4258804571(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
lookup__4258804571:
	for {
		if false {
			continue lookup__4258804571
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}
