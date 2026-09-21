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

var cache_Main_Learn_dollar_Dict gopurs_runtime.Value
var once_Main_Learn_dollar_Dict sync.Once

func Get_Main_Learn_dollar_Dict() gopurs_runtime.Value {
	once_Main_Learn_dollar_Dict.Do(func() {
		cache_Main_Learn_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Learn_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_Learn_dollar_Dict
}

var cache_Main_IsString_dollar_Dict gopurs_runtime.Value
var once_Main_IsString_dollar_Dict sync.Once

func Get_Main_IsString_dollar_Dict() gopurs_runtime.Value {
	once_Main_IsString_dollar_Dict.Do(func() {
		cache_Main_IsString_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_IsString_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_IsString_dollar_Dict
}

var cache_Main_IsString_dollar_Dict__2463103355 gopurs_runtime.Value
var once_Main_IsString_dollar_Dict__2463103355 sync.Once

func Get_Main_IsString_dollar_Dict__2463103355() gopurs_runtime.Value {
	once_Main_IsString_dollar_Dict__2463103355.Do(func() {
		cache_Main_IsString_dollar_Dict__2463103355 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_IsString_dollar_Dict__2463103355(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_IsString_dollar_Dict__2463103355
}

var cache_Main_IsEq_dollar_Dict gopurs_runtime.Value
var once_Main_IsEq_dollar_Dict sync.Once

func Get_Main_IsEq_dollar_Dict() gopurs_runtime.Value {
	once_Main_IsEq_dollar_Dict.Do(func() {
		cache_Main_IsEq_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_IsEq_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_IsEq_dollar_Dict
}

var cache_Main_Arg_dollar_Dict gopurs_runtime.Value
var once_Main_Arg_dollar_Dict sync.Once

func Get_Main_Arg_dollar_Dict() gopurs_runtime.Value {
	once_Main_Arg_dollar_Dict.Do(func() {
		cache_Main_Arg_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Arg_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_Arg_dollar_Dict
}

var cache_Main_reflIsEq gopurs_runtime.Value
var once_Main_reflIsEq sync.Once

func Get_Main_reflIsEq() gopurs_runtime.Value {
	once_Main_reflIsEq.Do(func() {
		cache_Main_reflIsEq = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_reflIsEq
}

var cache_Main_reflArg gopurs_runtime.Value
var once_Main_reflArg sync.Once

func Get_Main_reflArg() gopurs_runtime.Value {
	once_Main_reflArg.Do(func() {
		cache_Main_reflArg = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_reflArg
}

var cache_Main_notIsEq gopurs_runtime.Value
var once_Main_notIsEq sync.Once

func Get_Main_notIsEq() gopurs_runtime.Value {
	once_Main_notIsEq.Do(func() {
		cache_Main_notIsEq = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_notIsEq
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_learnIsString gopurs_runtime.Value
var once_Main_learnIsString sync.Once

func Get_Main_learnIsString() gopurs_runtime.Value {
	once_Main_learnIsString.Do(func() {
		cache_Main_learnIsString = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_learnIsString(_dollar___unused_0_box, _dollar___unused_1_box, uint32(v_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_learnIsString
}

var cache_Main_learnIsString__2172666286 gopurs_runtime.Value
var once_Main_learnIsString__2172666286 sync.Once

func Get_Main_learnIsString__2172666286() gopurs_runtime.Value {
	once_Main_learnIsString__2172666286.Do(func() {
		cache_Main_learnIsString__2172666286 = gopurs_runtime.Func(func(v_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_learnIsString__2172666286(uint32(v_unused_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_learnIsString__2172666286
}

var cache_Main_learnIsString__2315041832 gopurs_runtime.Value
var once_Main_learnIsString__2315041832 sync.Once

func Get_Main_learnIsString__2315041832() gopurs_runtime.Value {
	once_Main_learnIsString__2315041832.Do(func() {
		cache_Main_learnIsString__2315041832 = gopurs_runtime.Func(func(v_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_learnIsString__2315041832(uint32(v_unused_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_learnIsString__2315041832
}

var cache_Main_learnInst gopurs_runtime.Value
var once_Main_learnInst sync.Once

func Get_Main_learnInst() gopurs_runtime.Value {
	once_Main_learnInst.Do(func() {
		cache_Main_learnInst = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_learnInst
}

var cache_Main_isStringString gopurs_runtime.Value
var once_Main_isStringString sync.Once

func Get_Main_isStringString() gopurs_runtime.Value {
	once_Main_isStringString.Do(func() {
		cache_Main_isStringString = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_isStringString
}

var cache_Main_isStringElse gopurs_runtime.Value
var once_Main_isStringElse sync.Once

func Get_Main_isStringElse() gopurs_runtime.Value {
	once_Main_isStringElse.Do(func() {
		cache_Main_isStringElse = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_isStringElse
}

var cache_Main_isStringEg1 gopurs_runtime.Value
var once_Main_isStringEg1 sync.Once

func Get_Main_isStringEg1() gopurs_runtime.Value {
	once_Main_isStringEg1.Do(func() {
		cache_Main_isStringEg1 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_learnIsString(gopurs_runtime.Value{}, gopurs_runtime.Value{}, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_isStringEg1
}

var cache_Main_isStringEg0 gopurs_runtime.Value
var once_Main_isStringEg0 sync.Once

func Get_Main_isStringEg0() gopurs_runtime.Value {
	once_Main_isStringEg0.Do(func() {
		cache_Main_isStringEg0 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_learnIsString(gopurs_runtime.Value{}, gopurs_runtime.Value{}, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_isStringEg0
}

var cache_Main_isEq gopurs_runtime.Value
var once_Main_isEq sync.Once

func Get_Main_isEq() gopurs_runtime.Value {
	once_Main_isEq.Do(func() {
		cache_Main_isEq = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_isEq(_dollar___unused_0_box, uint32(v_1_box.IntVal), uint32(v1_2_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_isEq
}

var cache_Main_isEq__1533064487 gopurs_runtime.Value
var once_Main_isEq__1533064487 sync.Once

func Get_Main_isEq__1533064487() gopurs_runtime.Value {
	once_Main_isEq__1533064487.Do(func() {
		cache_Main_isEq__1533064487 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_isEq__1533064487(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_isEq__1533064487
}

var cache_Main_isEq__4057056784 gopurs_runtime.Value
var once_Main_isEq__4057056784 sync.Once

func Get_Main_isEq__4057056784() gopurs_runtime.Value {
	once_Main_isEq__4057056784.Do(func() {
		cache_Main_isEq__4057056784 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_isEq__4057056784(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_isEq__4057056784
}

var cache_Main_isEq__3632306064 gopurs_runtime.Value
var once_Main_isEq__3632306064 sync.Once

func Get_Main_isEq__3632306064() gopurs_runtime.Value {
	once_Main_isEq__3632306064.Do(func() {
		cache_Main_isEq__3632306064 = gopurs_runtime.Func2(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_isEq__3632306064(uint32(v_unused_0_box.IntVal), uint32(v1_unused_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_isEq__3632306064
}

var cache_Main_isEqEg0 gopurs_runtime.Value
var once_Main_isEqEg0 sync.Once

func Get_Main_isEqEg0() gopurs_runtime.Value {
	once_Main_isEqEg0.Do(func() {
		cache_Main_isEqEg0 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_isEq(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_isEqEg0
}

var cache_Main_isEqEg1 gopurs_runtime.Value
var once_Main_isEqEg1 sync.Once

func Get_Main_isEqEg1() gopurs_runtime.Value {
	once_Main_isEqEg1.Do(func() {
		cache_Main_isEqEg1 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_isEq(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_isEqEg1
}

var cache_Main_isEqEg2 gopurs_runtime.Value
var once_Main_isEqEg2 sync.Once

func Get_Main_isEqEg2() gopurs_runtime.Value {
	once_Main_isEqEg2.Do(func() {
		cache_Main_isEqEg2 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_isEq(gopurs_runtime.Value{}, 227768594, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_isEqEg2
}

var cache_Main_arg gopurs_runtime.Value
var once_Main_arg sync.Once

func Get_Main_arg() gopurs_runtime.Value {
	once_Main_arg.Do(func() {
		cache_Main_arg = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_arg(_dollar___unused_0_box, uint32(v_1_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_arg
}

var cache_Main_arg__2298718983 gopurs_runtime.Value
var once_Main_arg__2298718983 sync.Once

func Get_Main_arg__2298718983() gopurs_runtime.Value {
	once_Main_arg__2298718983.Do(func() {
		cache_Main_arg__2298718983 = gopurs_runtime.Func(func(v_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_arg__2298718983(uint32(v_unused_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_arg__2298718983
}

var cache_Main_arg__4084523869 gopurs_runtime.Value
var once_Main_arg__4084523869 sync.Once

func Get_Main_arg__4084523869() gopurs_runtime.Value {
	once_Main_arg__4084523869.Do(func() {
		cache_Main_arg__4084523869 = gopurs_runtime.Func(func(v_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_arg__4084523869(uint32(v_unused_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_arg__4084523869
}

var cache_Main_arg__2274491521 gopurs_runtime.Value
var once_Main_arg__2274491521 sync.Once

func Get_Main_arg__2274491521() gopurs_runtime.Value {
	once_Main_arg__2274491521.Do(func() {
		cache_Main_arg__2274491521 = gopurs_runtime.Func(func(v_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_arg__2274491521(uint32(v_unused_0_box.IntVal))), UnsafePtr: nil}
		})
	})
	return cache_Main_arg__2274491521
}

var cache_Main_argEg0 gopurs_runtime.Value
var once_Main_argEg0 sync.Once

func Get_Main_argEg0() gopurs_runtime.Value {
	once_Main_argEg0.Do(func() {
		cache_Main_argEg0 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_arg(gopurs_runtime.Value{}, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_argEg0
}

var cache_Main_appArg gopurs_runtime.Value
var once_Main_appArg sync.Once

func Get_Main_appArg() gopurs_runtime.Value {
	once_Main_appArg.Do(func() {
		cache_Main_appArg = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_appArg(_dollar___unused_0_box)
		})
	})
	return cache_Main_appArg
}

var cache_Main_argEg1 gopurs_runtime.Value
var once_Main_argEg1 sync.Once

func Get_Main_argEg1() gopurs_runtime.Value {
	once_Main_argEg1.Do(func() {
		cache_Main_argEg1 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_arg(gopurs_runtime.Value{}, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_argEg1
}

var cache_Main_argEg2 gopurs_runtime.Value
var once_Main_argEg2 sync.Once

func Get_Main_argEg2() gopurs_runtime.Value {
	once_Main_argEg2.Do(func() {
		cache_Main_argEg2 = gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Main_arg(gopurs_runtime.Value{}, 227768594)), UnsafePtr: nil}
	})
	return cache_Main_argEg2
}

type Constructor_Main_Proxy[T_p any] struct {
	Rc uint32
}

type Constructor_Main_Learn[T_a any, T_b any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[2338118058] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Learn[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Learn: " + key)
		}
	}
}

type Constructor_Main_IsString[T_t any, T_o any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[2058630097] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_IsString[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_IsString: " + key)
		}
	}
}

type Constructor_Main_IsEq[T_l any, T_r any, T_o any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3793749168] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_IsEq[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_IsEq: " + key)
		}
	}
}

type Constructor_Main_Arg[T_i any, T_o any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3676657418] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Arg[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Arg: " + key)
		}
	}
}

func Call_Main_Learn_dollar_Dict(x_0_loop struct {
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

func Call_Main_IsString_dollar_Dict(x_0_loop struct {
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

func Call_Main_IsString_dollar_Dict__2463103355(x_0_loop struct {
}) gopurs_runtime.Value {
IsString_dollar_Dict__2463103355:
	for {
		if false {
			continue IsString_dollar_Dict__2463103355
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

func Call_Main_IsEq_dollar_Dict(x_0_loop struct {
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

func Call_Main_Arg_dollar_Dict(x_0_loop struct {
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

func Call_Main_learnIsString(_dollar___unused_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, v_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
	_ = _dollar___unused_1
	var v_2 uint32 = v_2_loop
	_ = v_2
	return 227768594
}

func Call_Main_learnIsString__2172666286(v_unused_0_loop uint32) uint32 {
learnIsString__2172666286:
	for {
		if false {
			continue learnIsString__2172666286
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		return 227768594
	}
}

func Call_Main_learnIsString__2315041832(v_unused_0_loop uint32) uint32 {
learnIsString__2315041832:
	for {
		if false {
			continue learnIsString__2315041832
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		return 227768594
	}
}

func Call_Main_isEq(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32, v1_2_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	var v1_2 uint32 = v1_2_loop
	_ = v1_2
	return 227768594
}

func Call_Main_isEq__1533064487(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
isEq__1533064487:
	for {
		if false {
			continue isEq__1533064487
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}

func Call_Main_isEq__4057056784(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
isEq__4057056784:
	for {
		if false {
			continue isEq__4057056784
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}

func Call_Main_isEq__3632306064(v_unused_0_loop uint32, v1_unused_1_loop uint32) uint32 {
isEq__3632306064:
	for {
		if false {
			continue isEq__3632306064
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		var v1_unused_1 uint32 = v1_unused_1_loop
		_ = v1_unused_1
		return 227768594
	}
}

func Call_Main_arg(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop uint32) uint32 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var v_1 uint32 = v_1_loop
	_ = v_1
	return 227768594
}

func Call_Main_arg__2298718983(v_unused_0_loop uint32) uint32 {
arg__2298718983:
	for {
		if false {
			continue arg__2298718983
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		return 227768594
	}
}

func Call_Main_arg__4084523869(v_unused_0_loop uint32) uint32 {
arg__4084523869:
	for {
		if false {
			continue arg__4084523869
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		return 227768594
	}
}

func Call_Main_arg__2274491521(v_unused_0_loop uint32) uint32 {
arg__2274491521:
	for {
		if false {
			continue arg__2274491521
		}
		var v_unused_0 uint32 = v_unused_0_loop
		_ = v_unused_0
		return 227768594
	}
}

func Call_Main_appArg(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return func() gopurs_runtime.Value {
		orig := struct {
		}{}
		_ = orig
		return gopurs_runtime.RecordDict0()
	}()
}
