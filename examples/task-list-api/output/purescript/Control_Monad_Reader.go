package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Control_Monad_Reader_compose gopurs_runtime.Value
var once_Control_Monad_Reader_compose sync.Once

func Get_Control_Monad_Reader_compose() gopurs_runtime.Value {
	once_Control_Monad_Reader_compose.Do(func() {
		cache_Control_Monad_Reader_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Reader_compose
}

var cache_Control_Monad_Reader_unwrap gopurs_runtime.Value
var once_Control_Monad_Reader_unwrap sync.Once

func Get_Control_Monad_Reader_unwrap() gopurs_runtime.Value {
	once_Control_Monad_Reader_unwrap.Do(func() {
		cache_Control_Monad_Reader_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Control_Monad_Reader_unwrap
}

var cache_Control_Monad_Reader_compose1 gopurs_runtime.Value
var once_Control_Monad_Reader_compose1 sync.Once

func Get_Control_Monad_Reader_compose1() gopurs_runtime.Value {
	once_Control_Monad_Reader_compose1.Do(func() {
		cache_Control_Monad_Reader_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Reader_compose1
}

var cache_Control_Monad_Reader_compose2 gopurs_runtime.Value
var once_Control_Monad_Reader_compose2 sync.Once

func Get_Control_Monad_Reader_compose2() gopurs_runtime.Value {
	once_Control_Monad_Reader_compose2.Do(func() {
		cache_Control_Monad_Reader_compose2 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Reader_compose2
}

var cache_Control_Monad_Reader_unwrap1 gopurs_runtime.Value
var once_Control_Monad_Reader_unwrap1 sync.Once

func Get_Control_Monad_Reader_unwrap1() gopurs_runtime.Value {
	once_Control_Monad_Reader_unwrap1.Do(func() {
		cache_Control_Monad_Reader_unwrap1 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Control_Monad_Reader_unwrap1
}

var cache_Control_Monad_Reader_withReader gopurs_runtime.Value
var once_Control_Monad_Reader_withReader sync.Once

func Get_Control_Monad_Reader_withReader() gopurs_runtime.Value {
	once_Control_Monad_Reader_withReader.Do(func() {
		cache_Control_Monad_Reader_withReader = Get_Control_Monad_Reader_Trans_withReaderT()
	})
	return cache_Control_Monad_Reader_withReader
}

var cache_Control_Monad_Reader_runReader gopurs_runtime.Value
var once_Control_Monad_Reader_runReader sync.Once

func Get_Control_Monad_Reader_runReader() gopurs_runtime.Value {
	once_Control_Monad_Reader_runReader.Do(func() {
		cache_Control_Monad_Reader_runReader = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Reader_runReader(v_0_box)
		})
	})
	return cache_Control_Monad_Reader_runReader
}

var cache_Control_Monad_Reader_mapReader gopurs_runtime.Value
var once_Control_Monad_Reader_mapReader sync.Once

func Get_Control_Monad_Reader_mapReader() gopurs_runtime.Value {
	once_Control_Monad_Reader_mapReader.Do(func() {
		cache_Control_Monad_Reader_mapReader = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Reader_mapReader(f_0_box)
		})
	})
	return cache_Control_Monad_Reader_mapReader
}

func Call_Control_Monad_Reader_runReader(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), v_0)
}

func Call_Control_Monad_Reader_mapReader(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	return gopurs_runtime.Apply(Get_Control_Monad_Reader_Trans_mapReaderT(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Identity_Identity(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})))))
}
