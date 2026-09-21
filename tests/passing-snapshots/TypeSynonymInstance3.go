package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_LtEq_dollar_Dict gopurs_runtime.Value
var once_Main_LtEq_dollar_Dict sync.Once

func Get_Main_LtEq_dollar_Dict() gopurs_runtime.Value {
	once_Main_LtEq_dollar_Dict.Do(func() {
		cache_Main_LtEq_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_LtEq_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_LtEq_dollar_Dict
}

var cache_Main_LtEq_dollar_Dict__3764444104 gopurs_runtime.Value
var once_Main_LtEq_dollar_Dict__3764444104 sync.Once

func Get_Main_LtEq_dollar_Dict__3764444104() gopurs_runtime.Value {
	once_Main_LtEq_dollar_Dict__3764444104.Do(func() {
		cache_Main_LtEq_dollar_Dict__3764444104 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_LtEq_dollar_Dict__3764444104(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_LtEq_dollar_Dict__3764444104
}

var cache_Main_Lte256_dollar_Dict gopurs_runtime.Value
var once_Main_Lte256_dollar_Dict sync.Once

func Get_Main_Lte256_dollar_Dict() gopurs_runtime.Value {
	once_Main_Lte256_dollar_Dict.Do(func() {
		cache_Main_Lte256_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 907985874, UnsafePtr: unsafe.Pointer(Call_Main_Lte256_dollar_Dict(func() struct {
				LtEq0 gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					LtEq0 gopurs_runtime.Value
				}{}
				clone.LtEq0 = gopurs_runtime.RecordGet(orig, "LtEq0")
				return clone
			}()))}
		})
	})
	return cache_Main_Lte256_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_ltEqD8D256 gopurs_runtime.Value
var once_Main_ltEqD8D256 sync.Once

func Get_Main_ltEqD8D256() gopurs_runtime.Value {
	once_Main_ltEqD8D256.Do(func() {
		cache_Main_ltEqD8D256 = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_ltEqD8D256
}

var cache_Main_lte256 gopurs_runtime.Value
var once_Main_lte256 sync.Once

func Get_Main_lte256() gopurs_runtime.Value {
	once_Main_lte256.Do(func() {
		cache_Main_lte256 = gopurs_runtime.Value{Type: 9, IntVal: 907985874, UnsafePtr: unsafe.Pointer((&Constructor_Main_Lte256[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})}))}
	})
	return cache_Main_lte256
}

type Constructor_Main_LtEq[T_a any, T_b any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3632241042] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_LtEq[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_LtEq: " + key)
		}
	}
}

type Constructor_Main_Lte256[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[907985874] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Lte256[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "LtEq0":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Main_Lte256: " + key)
		}
	}
}

func Call_Main_LtEq_dollar_Dict(x_0_loop struct {
}) gopurs_runtime.Value {
	var x_0 struct {
	} = x_0_loop
	_ = x_0
	return func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}

func Call_Main_LtEq_dollar_Dict__3764444104(x_0_loop struct {
}) gopurs_runtime.Value {
LtEq_dollar_Dict__3764444104:
	for {
		if false {
			continue LtEq_dollar_Dict__3764444104
		}
		var x_0 struct {
		} = x_0_loop
		_ = x_0
		return func() gopurs_runtime.Value {
			orig := x_0
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	}
}

func Call_Main_Lte256_dollar_Dict(x_0_loop struct {
	LtEq0 gopurs_runtime.Value
}) *Constructor_Main_Lte256[gopurs_runtime.Value] {
	var x_0 struct {
		LtEq0 gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_Lte256[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict1("LtEq0", orig.LtEq0)
	}())
}
