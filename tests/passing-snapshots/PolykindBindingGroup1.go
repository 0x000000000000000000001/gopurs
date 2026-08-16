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
			return gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((&Constructor_Main_X{1, value0}))}
		})
	})
	return cache_Main_X
}

var cache_Main_Z gopurs_runtime.Value
var once_Main_Z sync.Once

func Get_Main_Z() gopurs_runtime.Value {
	once_Main_Z.Do(func() {
		cache_Main_Z = gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((*Constructor_Main_X)(nil))}
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

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((&Constructor_Main_X{1, gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((*Constructor_Main_X)(nil))}}))}
	})
	return cache_Main_test4
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((&Constructor_Main_X{1, gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((*Constructor_Main_X)(nil))}}))}
	})
	return cache_Main_test3
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((&Constructor_Main_X{1, gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((*Constructor_Main_X)(nil))}}))}
	})
	return cache_Main_test2
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((&Constructor_Main_X{1, gopurs_runtime.Value{Type: 9, IntVal: 1409933510, UnsafePtr: unsafe.Pointer((*Constructor_Main_X)(nil))}}))}
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

type Constructor_Main_X struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

type Constructor_Main_Z struct {
	Rc uint32
}

type Constructor_Main_Y struct {
	Rc uint32
	V0 *Constructor_Main_X
}
