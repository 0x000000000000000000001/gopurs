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
		cache_Main_X = gopurs_runtime.Value{Type: 9, IntVal: int64(1409933510), UnsafePtr: nil}
	})
	return cache_Main_X
}

var cache_Main_x gopurs_runtime.Value
var once_Main_x sync.Once

func Get_Main_x() gopurs_runtime.Value {
	once_Main_x.Do(func() {
		cache_Main_x = gopurs_runtime.Value{Type: 9, IntVal: int64(1409933510), UnsafePtr: nil}
	})
	return cache_Main_x
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(dictMonad_0_box)
		})
	})
	return cache_Main_test
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_cA gopurs_runtime.Value
var once_Main_cA sync.Once

func Get_Main_cA() gopurs_runtime.Value {
	once_Main_cA.Do(func() {
		cache_Main_cA = gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer((&Constructor_Main_C[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x1_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
			return x1_0
		})}))}
	})
	return cache_Main_cA
}

var cache_Main_c gopurs_runtime.Value
var once_Main_c sync.Once

func Get_Main_c() gopurs_runtime.Value {
	once_Main_c.Do(func() {
		cache_Main_c = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c(gopurs_runtime.CoerceToStruct[Constructor_Main_C[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Main_c
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test2(dictMonad_0_box)
		})
	})
	return cache_Main_test2
}

type Constructor_Main_X[T_a any] struct {
	Rc uint32
}

type Constructor_Main_C[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

//gopurs:init-v1:00670065006e006500720061007400650064002f004d00610069006e:00720065006700690073007400650072002f0043006f006e007300740072007500630074006f0072005f004d00610069006e005f0043
func init() {
	gopurs_runtime.StructGetters[2167983901] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_C[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "c":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_C: " + key)
		}
	}
}

func Call_Main_test(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))), func() gopurs_runtime.Value {
		orig := struct {
			x uint32
			/*gopurs:record-v1:1:0078*/
		}{uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1409933510), UnsafePtr: nil}.IntVal)}
		_ = orig
		return gopurs_runtime.RecordDict1("x", gopurs_runtime.Value{Type: 9, IntVal: int64(orig.x), UnsafePtr: nil})
	}())
}

func Call_Main_c(dict_0_loop *Constructor_Main_C[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_C[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return dict_0.V0
}

func Call_Main_test2(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Apply(Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))), func() gopurs_runtime.Value {
		orig := struct {
			ccc gopurs_runtime.Value
			/*gopurs:record-v1:1:006300630063*/
		}{gopurs_runtime.Func2(func(x1_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return x1_1
		})}
		_ = orig
		return gopurs_runtime.RecordDict1("ccc", orig.ccc)
	}())
}
