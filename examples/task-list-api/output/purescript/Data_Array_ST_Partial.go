package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Array_ST_Partial_poke gopurs_runtime.Value
var once_Data_Array_ST_Partial_poke sync.Once

func Get_Data_Array_ST_Partial_poke() gopurs_runtime.Value {
	once_Data_Array_ST_Partial_poke.Do(func() {
		cache_Data_Array_ST_Partial_poke = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_ST_Partial_poke(_dollar___unused_0_box)
		})
	})
	return cache_Data_Array_ST_Partial_poke
}

var cache_Data_Array_ST_Partial_peek gopurs_runtime.Value
var once_Data_Array_ST_Partial_peek sync.Once

func Get_Data_Array_ST_Partial_peek() gopurs_runtime.Value {
	once_Data_Array_ST_Partial_peek.Do(func() {
		cache_Data_Array_ST_Partial_peek = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Data_Array_ST_Partial_peek(_dollar___unused_0_box)
		})
	})
	return cache_Data_Array_ST_Partial_peek
}

func Call_Data_Array_ST_Partial_poke(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn3(), Get_Data_Array_ST_Partial_pokeImpl())
}

func Call_Data_Array_ST_Partial_peek(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	return gopurs_runtime.Apply(Get_Control_Monad_ST_Uncurried_runSTFn2(), Get_Data_Array_ST_Partial_peekImpl())
}

func Get_Data_Array_ST_Partial_peekImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_Partial_PeekImpl
}

func Get_Data_Array_ST_Partial_pokeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ST_Partial_PokeImpl
}
