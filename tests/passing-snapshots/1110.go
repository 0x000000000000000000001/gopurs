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

var cache_Main_C_dollar_Dict gopurs_runtime.Value
var once_Main_C_dollar_Dict sync.Once

func Get_Main_C_dollar_Dict() gopurs_runtime.Value {
	once_Main_C_dollar_Dict.Do(func() {
		cache_Main_C_dollar_Dict = gopurs_runtime.Func(func(x1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_C_dollar_Dict(x1_0_box)
		})
	})
	return cache_Main_C_dollar_Dict
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
		cache_Main_cA = gopurs_runtime.Value{Type: 9, IntVal: 2167983901, UnsafePtr: unsafe.Pointer(&Constructor_Main_C{1, gopurs_runtime.Func(func(x1_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x1_0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}())
			})
		})})}
	})
	return cache_Main_cA
}

var cache_Main_c gopurs_runtime.Value
var once_Main_c sync.Once

func Get_Main_c() gopurs_runtime.Value {
	once_Main_c.Do(func() {
		cache_Main_c = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_c(gopurs_runtime.CoerceToStruct[Constructor_Main_C](dict_0_box))
		})
	})
	return cache_Main_c
}

var cache_Main_c1 gopurs_runtime.Value
var once_Main_c1 sync.Once

func Get_Main_c1() gopurs_runtime.Value {
	once_Main_c1.Do(func() {
		cache_Main_c1 = gopurs_runtime.Func2(func(x1_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(Call_Main_c1(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(x1_0_box.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}(), v_1_box))
		})
	})
	return cache_Main_c1
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

var cache_Main_x__2502356749 gopurs_runtime.Value
var once_Main_x__2502356749 sync.Once

func Get_Main_x__2502356749() gopurs_runtime.Value {
	once_Main_x__2502356749.Do(func() {
		cache_Main_x__2502356749 = gopurs_runtime.Value{Type: 9, IntVal: int64(1409933510), UnsafePtr: nil}
	})
	return cache_Main_x__2502356749
}

type Constructor_Main_X struct {
	Rc uint32
}

type Constructor_Main_C struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[2167983901] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_C)(ptr)
		_ = c
		switch key {
		case "c":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_C: " + key)
		}
	}
}

func Call_Main_C_dollar_Dict(x1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x1_0 gopurs_runtime.Value = x1_0_loop
	_ = x1_0
	return x1_0
}

func Call_Main_test(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.RecordDict1("x", gopurs_runtime.Value{Type: 9, IntVal: int64(1409933510), UnsafePtr: nil}))
}

func Call_Main_c(dict_0_loop *Constructor_Main_C) gopurs_runtime.Value {
	var dict_0 *Constructor_Main_C = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Main_c1(x1_0_loop []gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
	var x1_0 []gopurs_runtime.Value = x1_0_loop
	_ = x1_0
	var v_1 gopurs_runtime.Value = v_1_loop
	_ = v_1
	return x1_0
}

func Call_Main_test2(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
	_ = dictMonad_0
	return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.RecordDict1("ccc", gopurs_runtime.Func(func(x1_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Array(func() []gopurs_runtime.Value {
				arr := *(*[]gopurs_runtime.Value)(x1_1.UnsafePtr)
				unboxed := make([]gopurs_runtime.Value, len(arr))
				for i, v := range arr {
					unboxed[i] = v
				}
				return unboxed
			}())
		})
	})))
}
