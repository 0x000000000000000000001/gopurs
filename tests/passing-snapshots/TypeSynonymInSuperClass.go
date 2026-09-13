package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_MonadAsk_dollar_Dict gopurs_runtime.Value
var once_Main_MonadAsk_dollar_Dict sync.Once

func Get_Main_MonadAsk_dollar_Dict() gopurs_runtime.Value {
	once_Main_MonadAsk_dollar_Dict.Do(func() {
		cache_Main_MonadAsk_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1125470254, UnsafePtr: unsafe.Pointer(Call_Main_MonadAsk_dollar_Dict(func() struct {
				Monad0 gopurs_runtime.Value
				ask    gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Monad0 gopurs_runtime.Value
					ask    gopurs_runtime.Value
				}{}
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				clone.ask = gopurs_runtime.RecordGet(orig, "ask")
				return clone
			}()))}
		})
	})
	return cache_Main_MonadAsk_dollar_Dict
}

var cache_Main_MonadAskEnv_dollar_Dict gopurs_runtime.Value
var once_Main_MonadAskEnv_dollar_Dict sync.Once

func Get_Main_MonadAskEnv_dollar_Dict() gopurs_runtime.Value {
	once_Main_MonadAskEnv_dollar_Dict.Do(func() {
		cache_Main_MonadAskEnv_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 239845587, UnsafePtr: unsafe.Pointer(Call_Main_MonadAskEnv_dollar_Dict(func() struct {
				Monad0    gopurs_runtime.Value
				MonadAsk1 gopurs_runtime.Value
			} {
				orig := x_0_box
				_ = orig
				clone := struct {
					Monad0    gopurs_runtime.Value
					MonadAsk1 gopurs_runtime.Value
				}{}
				clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
				clone.MonadAsk1 = gopurs_runtime.RecordGet(orig, "MonadAsk1")
				return clone
			}()))}
		})
	})
	return cache_Main_MonadAskEnv_dollar_Dict
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_ask gopurs_runtime.Value
var once_Main_ask sync.Once

func Get_Main_ask() gopurs_runtime.Value {
	once_Main_ask.Do(func() {
		cache_Main_ask = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_ask(dict_0_box)
		})
	})
	return cache_Main_ask
}

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once

func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Func(func(dictMonadAskEnv_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(dictMonadAskEnv_0_box)
		})
	})
	return cache_Main_test
}

type Constructor_Main_MonadAsk[T_r any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1125470254] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "ask":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_MonadAsk: " + key)
		}
	}
}

type Constructor_Main_MonadAskEnv[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[239845587] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MonadAskEnv[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0":
			return gopurs_runtime.Box(c.V0)
		case "MonadAsk1":
			return gopurs_runtime.Box(c.V1)
		default:
			panic("Key not found in dictionary Constructor_Main_MonadAskEnv: " + key)
		}
	}
}

func Call_Main_MonadAsk_dollar_Dict(x_0_loop struct {
	Monad0 gopurs_runtime.Value
	ask    gopurs_runtime.Value
}) *Constructor_Main_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] {
	var x_0 struct {
		Monad0 gopurs_runtime.Value
		ask    gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Monad0", "ask", orig.Monad0, orig.ask)
	}())
}

func Call_Main_MonadAskEnv_dollar_Dict(x_0_loop struct {
	Monad0    gopurs_runtime.Value
	MonadAsk1 gopurs_runtime.Value
}) *Constructor_Main_MonadAskEnv[gopurs_runtime.Value] {
	var x_0 struct {
		Monad0    gopurs_runtime.Value
		MonadAsk1 gopurs_runtime.Value
	} = x_0_loop
	_ = x_0
	return gopurs_runtime.CoerceToStruct[Constructor_Main_MonadAskEnv[gopurs_runtime.Value]](func() gopurs_runtime.Value {
		orig := x_0
		_ = orig
		return gopurs_runtime.RecordDict2("Monad0", "MonadAsk1", orig.Monad0, orig.MonadAsk1)
	}())
}

func Call_Main_ask(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "ask")
}

func Call_Main_test(dictMonadAskEnv_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadAskEnv_0 gopurs_runtime.Value = dictMonadAskEnv_0_loop
	_ = dictMonadAskEnv_0
	// TAST (Let): MonadAsk1_1_0 shape=App(Other) bindingType=Any
	MonadAsk1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAskEnv_0, "MonadAsk1"), gopurs_runtime.Value{})
	_ = MonadAsk1_1_0
	// TAST (Let): Monad0_2_1 shape=App(Other) bindingType=Any
	Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadAsk1_1_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_2_1
	// TAST (Let): Applicative0_3_2 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope3)])
	Applicative0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Applicative0"), gopurs_runtime.Value{}))
	_ = Applicative0_3_2
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Bind1"), gopurs_runtime.Value{}), "bind"), Call_Main_ask(MonadAsk1_1_0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(Applicative0_3_2.V1, gopurs_runtime.Bool((gopurs_runtime.RecordGet(v_4, "foo").StrVal()) == ("test")))
	}))
}
