package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_Proxy gopurs_runtime.Value
var once_Main_Proxy sync.Once

func Get_Main_Proxy() gopurs_runtime.Value {
	once_Main_Proxy.Do(func() {
		cache_Main_Proxy = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_Proxy
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_ctorKind1 gopurs_runtime.Value
var once_Main_ctorKind1 sync.Once

func Get_Main_ctorKind1() gopurs_runtime.Value {
	once_Main_ctorKind1.Do(func() {
		cache_Main_ctorKind1 = func() gopurs_runtime.Value {
			orig := struct {

				/*gopurs:record-v1:0:*/
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_ctorKind1
}

var cache_Main_ctorKind0 gopurs_runtime.Value
var once_Main_ctorKind0 sync.Once

func Get_Main_ctorKind0() gopurs_runtime.Value {
	once_Main_ctorKind0.Do(func() {
		cache_Main_ctorKind0 = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ctorKind0(_dollar___unused_0_box)
		})
	})
	return cache_Main_ctorKind0
}

var cache_Main_ctorKind gopurs_runtime.Value
var once_Main_ctorKind sync.Once

func Get_Main_ctorKind() gopurs_runtime.Value {
	once_Main_ctorKind.Do(func() {
		cache_Main_ctorKind = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_ctorKind(_dollar___unused_0_box, uint32(v_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_ctorKind
}

var cache_Main_testCtor1 gopurs_runtime.Value
var once_Main_testCtor1 sync.Once

func Get_Main_testCtor1() gopurs_runtime.Value {
	once_Main_testCtor1.Do(func() {
		cache_Main_testCtor1 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_ctorKind(gopurs_runtime.Value{}, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_testCtor1
}

var cache_Main_testCtor2 gopurs_runtime.Value
var once_Main_testCtor2 sync.Once

func Get_Main_testCtor2() gopurs_runtime.Value {
	once_Main_testCtor2.Do(func() {
		cache_Main_testCtor2 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_ctorKind(gopurs_runtime.Value{}, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_testCtor2
}

var cache_Main_testCtor3 gopurs_runtime.Value
var once_Main_testCtor3 sync.Once

func Get_Main_testCtor3() gopurs_runtime.Value {
	once_Main_testCtor3.Do(func() {
		cache_Main_testCtor3 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_ctorKind(gopurs_runtime.Value{}, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_testCtor3
}

type Constructor_Main_Proxy[T_a any] struct {
	Rc uint32
}

type Constructor_Main_CtorKind[T_ctor any, T_kind any] struct {
	Rc uint32
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f00430074006f0072004b0069006e0064
func init() {
	gopurs_runtime.StructGetters[133020252] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_CtorKind[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_CtorKind: " + key)
		}
	}
}

func Call_Main_ctorKind0(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return func() gopurs_runtime.Value {
		orig := struct {

			/*gopurs:record-v1:0:*/
		}{}
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Main_ctorKind(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return 227768594
}
