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

var cache_Main_P gopurs_runtime.Value
var once_Main_P sync.Once

func Get_Main_P() gopurs_runtime.Value {
	once_Main_P.Do(func() {
		cache_Main_P = gopurs_runtime.Value{Type: 9, IntVal: int64(3253896398), UnsafePtr: nil}
	})
	return cache_Main_P
}

var cache_Main_Test_dollar_Dict gopurs_runtime.Value
var once_Main_Test_dollar_Dict sync.Once

func Get_Main_Test_dollar_Dict() gopurs_runtime.Value {
	once_Main_Test_dollar_Dict.Do(func() {
		cache_Main_Test_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test_dollar_Dict(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_Test_dollar_Dict
}

var cache_Main_Test_dollar_Dict__3889889207 gopurs_runtime.Value
var once_Main_Test_dollar_Dict__3889889207 sync.Once

func Get_Main_Test_dollar_Dict__3889889207() gopurs_runtime.Value {
	once_Main_Test_dollar_Dict__3889889207.Do(func() {
		cache_Main_Test_dollar_Dict__3889889207 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test_dollar_Dict__3889889207(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_Test_dollar_Dict__3889889207
}

var cache_Main_Test_dollar_Dict__1879757193 gopurs_runtime.Value
var once_Main_Test_dollar_Dict__1879757193 sync.Once

func Get_Main_Test_dollar_Dict__1879757193() gopurs_runtime.Value {
	once_Main_Test_dollar_Dict__1879757193.Do(func() {
		cache_Main_Test_dollar_Dict__1879757193 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_Test_dollar_Dict__1879757193(func() struct {
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
				}{}

				return clone
			}())
		})
	})
	return cache_Main_Test_dollar_Dict__1879757193
}

var cache_Main_testClass2 gopurs_runtime.Value
var once_Main_testClass2 sync.Once

func Get_Main_testClass2() gopurs_runtime.Value {
	once_Main_testClass2.Do(func() {
		cache_Main_testClass2 = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_testClass2
}

var cache_Main_testClass1 gopurs_runtime.Value
var once_Main_testClass1 sync.Once

func Get_Main_testClass1() gopurs_runtime.Value {
	once_Main_testClass1.Do(func() {
		cache_Main_testClass1 = func() gopurs_runtime.Value {
			orig := struct {
			}{}
			_ = orig
			return gopurs_runtime.RecordDict0()
		}()
	})
	return cache_Main_testClass1
}

var cache_Main_test4 gopurs_runtime.Value
var once_Main_test4 sync.Once

func Get_Main_test4() gopurs_runtime.Value {
	once_Main_test4.Do(func() {
		cache_Main_test4 = gopurs_runtime.Value{Type: 9, IntVal: int64(3253896398), UnsafePtr: nil}
	})
	return cache_Main_test4
}

var cache_Main_test3 gopurs_runtime.Value
var once_Main_test3 sync.Once

func Get_Main_test3() gopurs_runtime.Value {
	once_Main_test3.Do(func() {
		cache_Main_test3 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test3
}

var cache_Main_test2 gopurs_runtime.Value
var once_Main_test2 sync.Once

func Get_Main_test2() gopurs_runtime.Value {
	once_Main_test2.Do(func() {
		cache_Main_test2 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
	})
	return cache_Main_test2
}

var cache_Main_test1 gopurs_runtime.Value
var once_Main_test1 sync.Once

func Get_Main_test1() gopurs_runtime.Value {
	once_Main_test1.Do(func() {
		cache_Main_test1 = gopurs_runtime.Value{Type: 9, IntVal: int64(227768594), UnsafePtr: nil}
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

type Constructor_Main_Proxy[T_a any] struct {
	Rc uint32
}

type Constructor_Main_P[T_a any] struct {
	Rc uint32
}

type Constructor_Main_Test[T_a any] struct {
	Rc uint32
}

func init() {
	gopurs_runtime.StructGetters[3423238824] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_Test[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {

		default:
			panic("Key not found in dictionary Constructor_Main_Test: " + key)
		}
	}
}

func Call_Main_Test_dollar_Dict(x_0_loop struct {
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

func Call_Main_Test_dollar_Dict__3889889207(x_0_loop struct {
}) gopurs_runtime.Value {
Test_dollar_Dict__3889889207:
	for {
		if false {
			continue Test_dollar_Dict__3889889207
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

func Call_Main_Test_dollar_Dict__1879757193(x_0_loop struct {
}) gopurs_runtime.Value {
Test_dollar_Dict__1879757193:
	for {
		if false {
			continue Test_dollar_Dict__1879757193
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
