package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_C0_dollar_Dict gopurs_runtime.Value
var once_Main_C0_dollar_Dict sync.Once

func Get_Main_C0_dollar_Dict() gopurs_runtime.Value {
	once_Main_C0_dollar_Dict.Do(func() {
		cache_Main_C0_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_C0_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())), UnsafePtr: nil}
		})
	})
	return cache_Main_C0_dollar_Dict
}

var cache_Main_C0_dollar_Dict__2835222047 gopurs_runtime.Value
var once_Main_C0_dollar_Dict__2835222047 sync.Once

func Get_Main_C0_dollar_Dict__2835222047() gopurs_runtime.Value {
	once_Main_C0_dollar_Dict__2835222047.Do(func() {
		cache_Main_C0_dollar_Dict__2835222047 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_C0_dollar_Dict__2835222047(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())), UnsafePtr: nil}
		})
	})
	return cache_Main_C0_dollar_Dict__2835222047
}

var cache_Main_C1_dollar_Dict gopurs_runtime.Value
var once_Main_C1_dollar_Dict sync.Once

func Get_Main_C1_dollar_Dict() gopurs_runtime.Value {
	once_Main_C1_dollar_Dict.Do(func() {
		cache_Main_C1_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4264042284, UnsafePtr: unsafe.Pointer(Call_Main_C1_dollar_Dict(func() struct {
				C00 gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					C00 gopurs_runtime.Value
				}{}
				clone.C00 = gopurs_runtime.RecordGet(orig, "C00")
				return clone
			}()))}
		})
	})
	return cache_Main_C1_dollar_Dict
}

var cache_Main_C1_dollar_Dict__3035685369 gopurs_runtime.Value
var once_Main_C1_dollar_Dict__3035685369 sync.Once

func Get_Main_C1_dollar_Dict__3035685369() gopurs_runtime.Value {
	once_Main_C1_dollar_Dict__3035685369.Do(func() {
		cache_Main_C1_dollar_Dict__3035685369 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 4264042284, UnsafePtr: unsafe.Pointer(Call_Main_C1_dollar_Dict__3035685369(func() struct {
				C00 gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					C00 gopurs_runtime.Value
				}{}
				clone.C00 = gopurs_runtime.RecordGet(orig, "C00")
				return clone
			}()))}
		})
	})
	return cache_Main_C1_dollar_Dict__3035685369
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_c0 gopurs_runtime.Value
var once_Main_c0 sync.Once

func Get_Main_c0() gopurs_runtime.Value {
	once_Main_c0.Do(func() {
		cache_Main_c0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}().IntVal)), UnsafePtr: nil}
	})
	return cache_Main_c0
}

var cache_Main_c1 gopurs_runtime.Value
var once_Main_c1 sync.Once

func Get_Main_c1() gopurs_runtime.Value {
	once_Main_c1.Do(func() {
		cache_Main_c1 = gopurs_runtime.Value{Type: 9, IntVal: 4264042284, UnsafePtr: unsafe.Pointer((&Constructor_Main_C1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Main_c1
}

type Constructor_Main_C0[T_a any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[1613519245] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_C0[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_C0: " + key)
		}
	}
}

type Constructor_Main_C1[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[4264042284] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_C1[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "C00":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_C1: " + key)
		}
	}
}

func Call_Main_C0_dollar_Dict(x_0_loop struct {
}) uint32 {
	var x_0 struct {
	} = x_0_loop
	_ = x_0
	return uint32(func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}().IntVal)
}

func Call_Main_C0_dollar_Dict__2835222047(x_0_loop struct {
}) uint32 {
C0_dollar_Dict__2835222047:
	for {
		if false {
			continue C0_dollar_Dict__2835222047
		}
		var x_0 struct {
		} = x_0_loop
		_ = x_0
		return uint32(func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict0()
		}().IntVal)
	}
}

func Call_Main_C1_dollar_Dict(x_0_loop struct {
	C00 gopurs_runtime.Value
}) *Constructor_Main_C1[gopurs_runtime.Value] {
	var x_0 struct {
		C00 gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_C1[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("C00", orig.C00)
	}())
}

func Call_Main_C1_dollar_Dict__3035685369(x_0_loop struct {
	C00 gopurs_runtime.Value
}) *Constructor_Main_C1[gopurs_runtime.Value] {
C1_dollar_Dict__3035685369:
	for {
		if false {
			continue C1_dollar_Dict__3035685369
		}
		var x_0 struct {
			C00 gopurs_runtime.Value
		} = x_0_loop
		_ = x_0
		return gopurs_runtime.CoerceToStruct[Constructor_Main_C1[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict1("C00", orig.C00)
		}())
	}
}
