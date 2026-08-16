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
			return Call_Main_MonadAsk_dollar_Dict(x_0_box)
		})
	})
	return cache_Main_MonadAsk_dollar_Dict
}

var cache_Main_monadAskFun gopurs_runtime.Value
var once_Main_monadAskFun sync.Once

func Get_Main_monadAskFun() gopurs_runtime.Value {
	once_Main_monadAskFun.Do(func() {
		cache_Main_monadAskFun = gopurs_runtime.Value{Type: 9, IntVal: 1125470254, UnsafePtr: unsafe.Pointer((&Constructor_Main_MonadAsk{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Control_Monad_monadFn()))}
		}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return x_0
		})}))}
	})
	return cache_Main_monadAskFun
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
		cache_Main_test = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Main_test(dictMonadAsk_0_box)
		})
	})
	return cache_Main_test
}

type Constructor_Main_MonadAsk struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[1125470254] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Main_MonadAsk)(ptr)
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

func Call_Main_MonadAsk_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Main_ask(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dict_0 gopurs_runtime.Value = dict_0_loop
	_ = dict_0
	return gopurs_runtime.RecordGet(dict_0, "ask")
}

func Call_Main_test(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
	_ = dictMonadAsk_0
	// TAST (Let): Monad0_1_0 shape=App(Other) expectedFromAst=gopurs_runtime.Value actual=gopurs_runtime.Value bindingType=Any
	Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})
	_ = Monad0_1_0
	// TAST (Let): Applicative0_2_1 shape=App(Other) expectedFromAst=*Constructor_Control_Applicative_Applicative actual=*Constructor_Control_Applicative_Applicative bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
	Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
	_ = Applicative0_2_1
	return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(dictMonadAsk_0, "ask"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Int((x_3.IntVal)+(1)))
	}))
}
