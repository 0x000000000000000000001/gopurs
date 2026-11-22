package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Node_Symbol_go__map gopurs_runtime.Value
var once_Node_Symbol_go__map sync.Once

func Get_Node_Symbol_go__map() gopurs_runtime.Value {
	once_Node_Symbol_go__map.Do(func() {
		cache_Node_Symbol_go__map = Get_Data_Functor_map__447590619()
	})
	return cache_Node_Symbol_go__map
}

var cache_Node_Symbol_showSymbol gopurs_runtime.Value
var once_Node_Symbol_showSymbol sync.Once

func Get_Node_Symbol_showSymbol() gopurs_runtime.Value {
	once_Node_Symbol_showSymbol.Do(func() {
		cache_Node_Symbol_showSymbol = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Node_Symbol_showSymbol(s_0_box))
		})
	})
	return cache_Node_Symbol_showSymbol
}

var cache_Node_Symbol_showJsSymbol gopurs_runtime.Value
var once_Node_Symbol_showJsSymbol sync.Once

func Get_Node_Symbol_showJsSymbol() gopurs_runtime.Value {
	once_Node_Symbol_showJsSymbol.Do(func() {
		cache_Node_Symbol_showJsSymbol = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(Get_Node_Symbol_showSymbolImpl(), x_0)
		})}))}
	})
	return cache_Node_Symbol_showJsSymbol
}

var cache_Node_Symbol_keyFor gopurs_runtime.Value
var once_Node_Symbol_keyFor sync.Once

func Get_Node_Symbol_keyFor() gopurs_runtime.Value {
	once_Node_Symbol_keyFor.Do(func() {
		cache_Node_Symbol_keyFor = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Symbol_keyFor(s_0_box)
		})
	})
	return cache_Node_Symbol_keyFor
}

var cache_Node_Symbol_go__for gopurs_runtime.Value
var once_Node_Symbol_go__for sync.Once

func Get_Node_Symbol_go__for() gopurs_runtime.Value {
	once_Node_Symbol_go__for.Do(func() {
		cache_Node_Symbol_go__for = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Node_Symbol_go__for(s_0_box.StrVal())
		})
	})
	return cache_Node_Symbol_go__for
}

func Call_Node_Symbol_showSymbol(s_0_loop gopurs_runtime.Value) string {
	var s_0 gopurs_runtime.Value = s_0_loop
	_ = s_0
	return gopurs_runtime.Apply(Get_Node_Symbol_showSymbolImpl(), s_0).StrVal()
}

func Call_Node_Symbol_keyFor(s_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var s_0 gopurs_runtime.Value = s_0_loop
	_ = s_0
	return Call_Control_Applicative_liftA1(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Effect_applicativeEffect()), Get_Data_Nullable_toMaybe(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Symbol_keyForImpl(), s_0)
	}))
}

func Call_Node_Symbol_go__for(s_0_loop string) gopurs_runtime.Value {
	var s_0 string = s_0_loop
	_ = s_0
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.UncurriedApp(Get_Node_Symbol_forImpl(), gopurs_runtime.Str(s_0))
	})
}

func Get_Node_Symbol_forImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Symbol_ForImpl
}

func Get_Node_Symbol_keyForImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Symbol_KeyForImpl
}

func Get_Node_Symbol_showSymbolImpl() gopurs_runtime.Value {
	return _Gopurs_Node_Symbol_ShowSymbolImpl
}
