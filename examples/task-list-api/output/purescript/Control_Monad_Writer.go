package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Writer_compose gopurs_runtime.Value
var once_Control_Monad_Writer_compose sync.Once

func Get_Control_Monad_Writer_compose() gopurs_runtime.Value {
	once_Control_Monad_Writer_compose.Do(func() {
		cache_Control_Monad_Writer_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Writer_compose
}

var cache_Control_Monad_Writer_compose1 gopurs_runtime.Value
var once_Control_Monad_Writer_compose1 sync.Once

func Get_Control_Monad_Writer_compose1() gopurs_runtime.Value {
	once_Control_Monad_Writer_compose1.Do(func() {
		cache_Control_Monad_Writer_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Writer_compose1
}

var cache_Control_Monad_Writer_unwrap gopurs_runtime.Value
var once_Control_Monad_Writer_unwrap sync.Once

func Get_Control_Monad_Writer_unwrap() gopurs_runtime.Value {
	once_Control_Monad_Writer_unwrap.Do(func() {
		cache_Control_Monad_Writer_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Control_Monad_Writer_unwrap
}

var cache_Control_Monad_Writer_writer gopurs_runtime.Value
var once_Control_Monad_Writer_writer sync.Once

func Get_Control_Monad_Writer_writer() gopurs_runtime.Value {
	once_Control_Monad_Writer_writer.Do(func() {
		cache_Control_Monad_Writer_writer = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Control_Monad_Writer_Trans_WriterT(), Get_Data_Identity_Identity())
	})
	return cache_Control_Monad_Writer_writer
}

var cache_Control_Monad_Writer_runWriter gopurs_runtime.Value
var once_Control_Monad_Writer_runWriter sync.Once

func Get_Control_Monad_Writer_runWriter() gopurs_runtime.Value {
	once_Control_Monad_Writer_runWriter.Do(func() {
		cache_Control_Monad_Writer_runWriter = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), Get_Control_Monad_Writer_Trans_runWriterT())
	})
	return cache_Control_Monad_Writer_runWriter
}

var cache_Control_Monad_Writer_mapWriter gopurs_runtime.Value
var once_Control_Monad_Writer_mapWriter sync.Once

func Get_Control_Monad_Writer_mapWriter() gopurs_runtime.Value {
	once_Control_Monad_Writer_mapWriter.Do(func() {
		cache_Control_Monad_Writer_mapWriter = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_mapWriter(f_0_box)
		})
	})
	return cache_Control_Monad_Writer_mapWriter
}

var cache_Control_Monad_Writer_execWriter gopurs_runtime.Value
var once_Control_Monad_Writer_execWriter sync.Once

func Get_Control_Monad_Writer_execWriter() gopurs_runtime.Value {
	once_Control_Monad_Writer_execWriter.Do(func() {
		cache_Control_Monad_Writer_execWriter = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Writer_execWriter(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](m_0_box))
		})
	})
	return cache_Control_Monad_Writer_execWriter
}

func Call_Control_Monad_Writer_mapWriter(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Identity_Identity(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(__local_var_1_0, v_2)
	})
}

func Call_Control_Monad_Writer_execWriter(m_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
	var m_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = m_0_loop
	_ = m_0
	return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(Get_Control_Monad_Writer_runWriter(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(m_0)}).UnsafePtr).V1
}
