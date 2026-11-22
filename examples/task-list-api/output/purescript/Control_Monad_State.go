package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Control_Monad_State_compose gopurs_runtime.Value
var once_Control_Monad_State_compose sync.Once

func Get_Control_Monad_State_compose() gopurs_runtime.Value {
	once_Control_Monad_State_compose.Do(func() {
		cache_Control_Monad_State_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_State_compose
}

var cache_Control_Monad_State_unwrap gopurs_runtime.Value
var once_Control_Monad_State_unwrap sync.Once

func Get_Control_Monad_State_unwrap() gopurs_runtime.Value {
	once_Control_Monad_State_unwrap.Do(func() {
		cache_Control_Monad_State_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Control_Monad_State_unwrap
}

var cache_Control_Monad_State_compose1 gopurs_runtime.Value
var once_Control_Monad_State_compose1 sync.Once

func Get_Control_Monad_State_compose1() gopurs_runtime.Value {
	once_Control_Monad_State_compose1.Do(func() {
		cache_Control_Monad_State_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_State_compose1
}

var cache_Control_Monad_State_compose2 gopurs_runtime.Value
var once_Control_Monad_State_compose2 sync.Once

func Get_Control_Monad_State_compose2() gopurs_runtime.Value {
	once_Control_Monad_State_compose2.Do(func() {
		cache_Control_Monad_State_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_State_compose2
}

var cache_Control_Monad_State_unwrap1 gopurs_runtime.Value
var once_Control_Monad_State_unwrap1 sync.Once

func Get_Control_Monad_State_unwrap1() gopurs_runtime.Value {
	once_Control_Monad_State_unwrap1.Do(func() {
		cache_Control_Monad_State_unwrap1 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Control_Monad_State_unwrap1
}

var cache_Control_Monad_State_withState gopurs_runtime.Value
var once_Control_Monad_State_withState sync.Once

func Get_Control_Monad_State_withState() gopurs_runtime.Value {
	once_Control_Monad_State_withState.Do(func() {
		cache_Control_Monad_State_withState = Get_Control_Monad_State_Trans_withStateT()
	})
	return cache_Control_Monad_State_withState
}

var cache_Control_Monad_State_runState gopurs_runtime.Value
var once_Control_Monad_State_runState sync.Once

func Get_Control_Monad_State_runState() gopurs_runtime.Value {
	once_Control_Monad_State_runState.Do(func() {
		cache_Control_Monad_State_runState = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_runState(v_0_box)
		})
	})
	return cache_Control_Monad_State_runState
}

var cache_Control_Monad_State_mapState gopurs_runtime.Value
var once_Control_Monad_State_mapState sync.Once

func Get_Control_Monad_State_mapState() gopurs_runtime.Value {
	once_Control_Monad_State_mapState.Do(func() {
		cache_Control_Monad_State_mapState = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_mapState(f_0_box)
		})
	})
	return cache_Control_Monad_State_mapState
}

var cache_Control_Monad_State_execState gopurs_runtime.Value
var once_Control_Monad_State_execState sync.Once

func Get_Control_Monad_State_execState() gopurs_runtime.Value {
	once_Control_Monad_State_execState.Do(func() {
		cache_Control_Monad_State_execState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_execState(v_0_box, s_1_box)
		})
	})
	return cache_Control_Monad_State_execState
}

var cache_Control_Monad_State_evalState gopurs_runtime.Value
var once_Control_Monad_State_evalState sync.Once

func Get_Control_Monad_State_evalState() gopurs_runtime.Value {
	once_Control_Monad_State_evalState.Do(func() {
		cache_Control_Monad_State_evalState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_State_evalState(v_0_box, s_1_box)
		})
	})
	return cache_Control_Monad_State_evalState
}

func Call_Control_Monad_State_runState(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), v_0)
}

func Call_Control_Monad_State_mapState(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply(Get_Control_Monad_State_Trans_mapStateT(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Identity_Identity(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})))))
}

func Call_Control_Monad_State_execState(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var s_1 gopurs_runtime.Value = s_1_loop
	_ = s_1
	return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(v_0, s_1).UnsafePtr).V1
}

func Call_Control_Monad_State_evalState(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var s_1 gopurs_runtime.Value = s_1_loop
	_ = s_1
	return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(v_0, s_1).UnsafePtr).V0
}
