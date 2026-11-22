package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Control_Comonad_Traced_unwrap gopurs_runtime.Value
var once_Control_Comonad_Traced_unwrap sync.Once

func Get_Control_Comonad_Traced_unwrap() gopurs_runtime.Value {
	once_Control_Comonad_Traced_unwrap.Do(func() {
		cache_Control_Comonad_Traced_unwrap = Call_Safe_Coerce_coerce(gopurs_runtime.Value{})
	})
	return cache_Control_Comonad_Traced_unwrap
}

var cache_Control_Comonad_Traced_traced gopurs_runtime.Value
var once_Control_Comonad_Traced_traced sync.Once

func Get_Control_Comonad_Traced_traced() gopurs_runtime.Value {
	once_Control_Comonad_Traced_traced.Do(func() {
		cache_Control_Comonad_Traced_traced = gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Identity_Identity(), Get_Control_Comonad_Traced_Trans_TracedT())
	})
	return cache_Control_Comonad_Traced_traced
}

var cache_Control_Comonad_Traced_runTraced gopurs_runtime.Value
var once_Control_Comonad_Traced_runTraced sync.Once

func Get_Control_Comonad_Traced_runTraced() gopurs_runtime.Value {
	once_Control_Comonad_Traced_runTraced.Do(func() {
		cache_Control_Comonad_Traced_runTraced = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Comonad_Traced_runTraced(v_0_box)
		})
	})
	return cache_Control_Comonad_Traced_runTraced
}

func Call_Control_Comonad_Traced_runTraced(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), v_0)
}
