package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_X gopurs_runtime.Value
var once_Main_X sync.Once

func Get_Main_X() gopurs_runtime.Value {
	once_Main_X.Do(func() {
		cache_Main_X = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((&Constructor_Main_X[gopurs_runtime.Value]{1, value0}))}
		})
	})
	return cache_Main_X
}

var cache_Main_X__2612941499 gopurs_runtime.Value
var once_Main_X__2612941499 sync.Once

func Get_Main_X__2612941499() gopurs_runtime.Value {
	once_Main_X__2612941499.Do(func() {
		cache_Main_X__2612941499 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer(Rebox_Main_2532313673_2799044882(Call_Main_X__2612941499(__eta_norm_0_0_box)))}
		})
	})
	return cache_Main_X__2612941499
}

var cache_Main_X__1955970587 gopurs_runtime.Value
var once_Main_X__1955970587 sync.Once

func Get_Main_X__1955970587() gopurs_runtime.Value {
	once_Main_X__1955970587.Do(func() {
		cache_Main_X__1955970587 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer(Call_Main_X__1955970587(__eta_norm_0_0_box))}
		})
	})
	return cache_Main_X__1955970587
}

var cache_Main_Z gopurs_runtime.Value
var once_Main_Z sync.Once

func Get_Main_Z() gopurs_runtime.Value {
	once_Main_Z.Do(func() {
		cache_Main_Z = gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((*Constructor_Main_X[gopurs_runtime.Value])(nil))}
	})
	return cache_Main_Z
}

var cache_Main_Y gopurs_runtime.Value
var once_Main_Y sync.Once

func Get_Main_Y() gopurs_runtime.Value {
	once_Main_Y.Do(func() {
		cache_Main_Y = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return value0
		})
	})
	return cache_Main_Y
}

var cache_Main_Y__1843750219 gopurs_runtime.Value
var once_Main_Y__1843750219 sync.Once

func Get_Main_Y__1843750219() gopurs_runtime.Value {
	once_Main_Y__1843750219.Do(func() {
		cache_Main_Y__1843750219 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Y__1843750219(Rebox_Main_2799044882_2532313673(gopurs_runtime.CoerceToStruct[Constructor_Main_X[gopurs_runtime.Value]](__eta_norm_0_unused_0_box)))
		})
	})
	return cache_Main_Y__1843750219
}

var cache_Main_Y__2801078043 gopurs_runtime.Value
var once_Main_Y__2801078043 sync.Once

func Get_Main_Y__2801078043() gopurs_runtime.Value {
	once_Main_Y__2801078043.Do(func() {
		cache_Main_Y__2801078043 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Y__2801078043(Rebox_Main_2799044882_2532313673(gopurs_runtime.CoerceToStruct[Constructor_Main_X[gopurs_runtime.Value]](__eta_norm_0_0_box)))
		})
	})
	return cache_Main_Y__2801078043
}

var cache_Main_Y__3768736619 gopurs_runtime.Value
var once_Main_Y__3768736619 sync.Once

func Get_Main_Y__3768736619() gopurs_runtime.Value {
	once_Main_Y__3768736619.Do(func() {
		cache_Main_Y__3768736619 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Y__3768736619(gopurs_runtime.CoerceToStruct[Constructor_Main_X[gopurs_runtime.Value]](__eta_norm_0_unused_0_box))
		})
	})
	return cache_Main_Y__3768736619
}

var cache_Main_Y__2627895611 gopurs_runtime.Value
var once_Main_Y__2627895611 sync.Once

func Get_Main_Y__2627895611() gopurs_runtime.Value {
	once_Main_Y__2627895611.Do(func() {
		cache_Main_Y__2627895611 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Y__2627895611(gopurs_runtime.CoerceToStruct[Constructor_Main_X[gopurs_runtime.Value]](__eta_norm_0_0_box))
		})
	})
	return cache_Main_Y__2627895611
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((&Constructor_Main_X[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((*Constructor_Main_X[gopurs_runtime.Value])(nil))}}))}
	})
	return cache_Main_test4
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer(Rebox_Main_2532313673_2799044882((&Constructor_Main_X[int64]{1, gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer(Rebox_Main_2532313673_2799044882((*Constructor_Main_X[int64])(nil)))}})))}
	})
	return cache_Main_test3
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((&Constructor_Main_X[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((*Constructor_Main_X[gopurs_runtime.Value])(nil))}}))}
	})
	return cache_Main_test2
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((&Constructor_Main_X[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer(Rebox_Main_2532313673_2799044882((*Constructor_Main_X[int64])(nil)))}}))}
	})
	return cache_Main_test1
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

type Constructor_Main_X[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Z[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Y[T_a any] struct {
	Rc uint32
	V0 *Constructor_Main_X[T_a]
}

func Call_Main_X__2612941499(__eta_norm_0_0_loop gopurs_runtime.Value) *Constructor_Main_X[int64] {
X__2612941499:
	for {
		if false {
			continue X__2612941499
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return (&Constructor_Main_X[int64]{1, __eta_norm_0_0})
	}
}

func Call_Main_X__1955970587(__eta_norm_0_0_loop gopurs_runtime.Value) *Constructor_Main_X[gopurs_runtime.Value] {
X__1955970587:
	for {
		if false {
			continue X__1955970587
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return (&Constructor_Main_X[gopurs_runtime.Value]{1, __eta_norm_0_0})
	}
}

func Call_Main_Y__1843750219(__eta_norm_0_unused_0_loop *Constructor_Main_X[int64]) gopurs_runtime.Value {
Y__1843750219:
	for {
		if false {
			continue Y__1843750219
		}
		var __eta_norm_0_unused_0 *Constructor_Main_X[int64] = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer(Rebox_Main_2532313673_2799044882((*Constructor_Main_X[int64])(nil)))}
	}
}

func Call_Main_Y__2801078043(__eta_norm_0_0_loop *Constructor_Main_X[int64]) gopurs_runtime.Value {
Y__2801078043:
	for {
		if false {
			continue Y__2801078043
		}
		var __eta_norm_0_0 *Constructor_Main_X[int64] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer(Rebox_Main_2532313673_2799044882(__eta_norm_0_0))}
	}
}

func Call_Main_Y__3768736619(__eta_norm_0_unused_0_loop *Constructor_Main_X[gopurs_runtime.Value]) gopurs_runtime.Value {
Y__3768736619:
	for {
		if false {
			continue Y__3768736619
		}
		var __eta_norm_0_unused_0 *Constructor_Main_X[gopurs_runtime.Value] = __eta_norm_0_unused_0_loop
		_ = __eta_norm_0_unused_0
		return gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((*Constructor_Main_X[gopurs_runtime.Value])(nil))}
	}
}

func Call_Main_Y__2627895611(__eta_norm_0_0_loop *Constructor_Main_X[gopurs_runtime.Value]) gopurs_runtime.Value {
Y__2627895611:
	for {
		if false {
			continue Y__2627895611
		}
		var __eta_norm_0_0 *Constructor_Main_X[gopurs_runtime.Value] = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer(__eta_norm_0_0)}
	}
}

func Rebox_Main_2532313673_2799044882(in *Constructor_Main_X[int64]) *Constructor_Main_X[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_X[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_2799044882_2532313673(in *Constructor_Main_X[gopurs_runtime.Value]) *Constructor_Main_X[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Main_X[int64]{}
	out.V0 = in.V0
	return out
}
