package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Example_HTTP_compose gopurs_runtime.Value
var once_Example_HTTP_compose sync.Once

func Get_Example_HTTP_compose() gopurs_runtime.Value {
	once_Example_HTTP_compose.Do(func() {
		cache_Example_HTTP_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Example_HTTP_compose
}

var cache_Example_HTTP_compose1 gopurs_runtime.Value
var once_Example_HTTP_compose1 sync.Once

func Get_Example_HTTP_compose1() gopurs_runtime.Value {
	once_Example_HTTP_compose1.Do(func() {
		cache_Example_HTTP_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Example_HTTP_compose1
}

var cache_Example_HTTP_serve gopurs_runtime.Value
var once_Example_HTTP_serve sync.Once

func Get_Example_HTTP_serve() gopurs_runtime.Value {
	once_Example_HTTP_serve.Do(func() {
		cache_Example_HTTP_serve = gopurs_runtime.Func2(func(address_0_box gopurs_runtime.Value, handler_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Example_HTTP_serve(address_0_box.StrVal(), handler_1_box)
		})
	})
	return cache_Example_HTTP_serve
}

func Call_Example_HTTP_serve(address_0_loop string, handler_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var address_0 string = address_0_loop
	_ = address_0
	var handler_1 gopurs_runtime.Value = handler_1_loop
	_ = handler_1
	return Call_Effect_Aff_bracket__3139827911(Call_Effect_Aff_Compat_fromEffectFnAff__2688986814(gopurs_runtime.Apply2(Get_Example_HTTP_startImpl(), gopurs_runtime.Str(address_0), gopurs_runtime.Func2(func(request_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply2(handler_1, request_2, gopurs_runtime.Apply(Get_Effect_Uncurried_runEffectFn1(), __local_var_3)), gopurs_runtime.Value{})
	}))), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff_Compat_fromEffectFnAff(), Get_Example_HTTP_stopImpl()), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Aff_Compat_fromEffectFnAff(), Get_Example_HTTP_waitImpl()))
}

func Get_Example_HTTP_startImpl() gopurs_runtime.Value {
	return _Gopurs_Example_HTTP_StartImpl
}

func Get_Example_HTTP_stopImpl() gopurs_runtime.Value {
	return _Gopurs_Example_HTTP_StopImpl
}

func Get_Example_HTTP_waitImpl() gopurs_runtime.Value {
	return _Gopurs_Example_HTTP_WaitImpl
}
