package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_RWS_pure gopurs_runtime.Value
var once_Control_Monad_RWS_pure sync.Once

func Get_Control_Monad_RWS_pure() gopurs_runtime.Value {
	once_Control_Monad_RWS_pure.Do(func() {
		cache_Control_Monad_RWS_pure = Get_Data_Identity_Identity()
	})
	return cache_Control_Monad_RWS_pure
}

var cache_Control_Monad_RWS_composeFlipped gopurs_runtime.Value
var once_Control_Monad_RWS_composeFlipped sync.Once

func Get_Control_Monad_RWS_composeFlipped() gopurs_runtime.Value {
	once_Control_Monad_RWS_composeFlipped.Do(func() {
		cache_Control_Monad_RWS_composeFlipped = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_RWS_composeFlipped
}

var cache_Control_Monad_RWS_unwrap gopurs_runtime.Value
var once_Control_Monad_RWS_unwrap sync.Once

func Get_Control_Monad_RWS_unwrap() gopurs_runtime.Value {
	once_Control_Monad_RWS_unwrap.Do(func() {
		cache_Control_Monad_RWS_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Control_Monad_RWS_unwrap
}

var cache_Control_Monad_RWS_composeFlipped1 gopurs_runtime.Value
var once_Control_Monad_RWS_composeFlipped1 sync.Once

func Get_Control_Monad_RWS_composeFlipped1() gopurs_runtime.Value {
	once_Control_Monad_RWS_composeFlipped1.Do(func() {
		cache_Control_Monad_RWS_composeFlipped1 = Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_RWS_composeFlipped1
}

var cache_Control_Monad_RWS_unwrap1 gopurs_runtime.Value
var once_Control_Monad_RWS_unwrap1 sync.Once

func Get_Control_Monad_RWS_unwrap1() gopurs_runtime.Value {
	once_Control_Monad_RWS_unwrap1.Do(func() {
		cache_Control_Monad_RWS_unwrap1 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Control_Monad_RWS_unwrap1
}

var cache_Control_Monad_RWS_execRWST gopurs_runtime.Value
var once_Control_Monad_RWS_execRWST sync.Once

func Get_Control_Monad_RWS_execRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_execRWST.Do(func() {
		cache_Control_Monad_RWS_execRWST = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Monad_RWS_execRWST(v_0_box, r_1_box, s_2_box)
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Control_Monad_RWS_execRWST
}

var cache_Control_Monad_RWS_unwrap2 gopurs_runtime.Value
var once_Control_Monad_RWS_unwrap2 sync.Once

func Get_Control_Monad_RWS_unwrap2() gopurs_runtime.Value {
	once_Control_Monad_RWS_unwrap2.Do(func() {
		cache_Control_Monad_RWS_unwrap2 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Control_Monad_RWS_unwrap2
}

var cache_Control_Monad_RWS_evalRWST gopurs_runtime.Value
var once_Control_Monad_RWS_evalRWST sync.Once

func Get_Control_Monad_RWS_evalRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_evalRWST.Do(func() {
		cache_Control_Monad_RWS_evalRWST = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Monad_RWS_evalRWST(v_0_box, r_1_box, s_2_box)
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Control_Monad_RWS_evalRWST
}

var cache_Control_Monad_RWS_withRWS gopurs_runtime.Value
var once_Control_Monad_RWS_withRWS sync.Once

func Get_Control_Monad_RWS_withRWS() gopurs_runtime.Value {
	once_Control_Monad_RWS_withRWS.Do(func() {
		cache_Control_Monad_RWS_withRWS = Get_Control_Monad_RWS_Trans_withRWST()
	})
	return cache_Control_Monad_RWS_withRWS
}

var cache_Control_Monad_RWS_rws gopurs_runtime.Value
var once_Control_Monad_RWS_rws sync.Once

func Get_Control_Monad_RWS_rws() gopurs_runtime.Value {
	once_Control_Monad_RWS_rws.Do(func() {
		cache_Control_Monad_RWS_rws = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Call_Control_Monad_RWS_rws(f_0_box, r_1_box, s_2_box))}
		})
	})
	return cache_Control_Monad_RWS_rws
}

var cache_Control_Monad_RWS_runRWS gopurs_runtime.Value
var once_Control_Monad_RWS_runRWS sync.Once

func Get_Control_Monad_RWS_runRWS() gopurs_runtime.Value {
	once_Control_Monad_RWS_runRWS.Do(func() {
		cache_Control_Monad_RWS_runRWS = gopurs_runtime.Func3(func(m_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Call_Control_Monad_RWS_runRWS(m_0_box, r_1_box, s_2_box))}
		})
	})
	return cache_Control_Monad_RWS_runRWS
}

var cache_Control_Monad_RWS_mapRWS gopurs_runtime.Value
var once_Control_Monad_RWS_mapRWS sync.Once

func Get_Control_Monad_RWS_mapRWS() gopurs_runtime.Value {
	once_Control_Monad_RWS_mapRWS.Do(func() {
		cache_Control_Monad_RWS_mapRWS = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_RWS_mapRWS(f_0_box)
		})
	})
	return cache_Control_Monad_RWS_mapRWS
}

var cache_Control_Monad_RWS_execRWS gopurs_runtime.Value
var once_Control_Monad_RWS_execRWS sync.Once

func Get_Control_Monad_RWS_execRWS() gopurs_runtime.Value {
	once_Control_Monad_RWS_execRWS.Do(func() {
		cache_Control_Monad_RWS_execRWS = gopurs_runtime.Func3(func(m_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Monad_RWS_execRWS(m_0_box, r_1_box, s_2_box)
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Control_Monad_RWS_execRWS
}

var cache_Control_Monad_RWS_evalRWS gopurs_runtime.Value
var once_Control_Monad_RWS_evalRWS sync.Once

func Get_Control_Monad_RWS_evalRWS() gopurs_runtime.Value {
	once_Control_Monad_RWS_evalRWS.Do(func() {
		cache_Control_Monad_RWS_evalRWS = gopurs_runtime.Func3(func(m_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Monad_RWS_evalRWS(m_0_box, r_1_box, s_2_box)
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
		})
	})
	return cache_Control_Monad_RWS_evalRWS
}

func Call_Control_Monad_RWS_execRWST(v_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var r_1 gopurs_runtime.Value = r_1_loop
	_ = r_1
	var s_2 gopurs_runtime.Value = s_2_loop
	_ = s_2
	// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeVar a$scope62)
	__local_var_3_0 := gopurs_runtime.Apply2(v_0, r_1, s_2)
	_ = __local_var_3_0
	return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{gopurs_runtime.Value{}, gopurs_runtime.Value{}}
}

func Call_Control_Monad_RWS_evalRWST(v_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	var r_1 gopurs_runtime.Value = r_1_loop
	_ = r_1
	var s_2 gopurs_runtime.Value = s_2_loop
	_ = s_2
	// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeVar a$scope62)
	__local_var_3_0 := gopurs_runtime.Apply2(v_0, r_1, s_2)
	_ = __local_var_3_0
	return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{gopurs_runtime.Value{}, gopurs_runtime.Value{}}
}

func Call_Control_Monad_RWS_rws(f_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	var r_1 gopurs_runtime.Value = r_1_loop
	_ = r_1
	var s_2 gopurs_runtime.Value = s_2_loop
	_ = s_2
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, r_1, s_2))
}

func Call_Control_Monad_RWS_runRWS(m_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	var m_0 gopurs_runtime.Value = m_0_loop
	_ = m_0
	var r_1 gopurs_runtime.Value = r_1_loop
	_ = r_1
	var s_2 gopurs_runtime.Value = s_2_loop
	_ = s_2
	return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(m_0, r_1, s_2))
}

func Call_Control_Monad_RWS_mapRWS(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply2(Call_Control_Semigroupoid_composeFlipped(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, Get_Data_Identity_Identity()))
	_ = __local_var_1_0
	return gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value, s_4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply2(v_2, r_3, s_4))
	})
}

func Call_Control_Monad_RWS_execRWS(m_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var m_0 gopurs_runtime.Value = m_0_loop
	_ = m_0
	var r_1 gopurs_runtime.Value = r_1_loop
	_ = r_1
	var s_2 gopurs_runtime.Value = s_2_loop
	_ = s_2
	// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeVar a$scope62)
	__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(m_0, r_1, s_2))
	_ = __local_var_3_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_3_0).V0, (__local_var_3_0).V2}))})))}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
}

func Call_Control_Monad_RWS_evalRWS(m_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
} {
	var m_0 gopurs_runtime.Value = m_0_loop
	_ = m_0
	var r_1 gopurs_runtime.Value = r_1_loop
	_ = r_1
	var s_2 gopurs_runtime.Value = s_2_loop
	_ = s_2
	// TAST (Let): __local_var_3_0 shape=App(Other) bindingType=(TypeVar a$scope62)
	__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(m_0, r_1, s_2))
	_ = __local_var_3_0
	return func() struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_3_0).V1, (__local_var_3_0).V2}))})))}; return struct {
		V0 gopurs_runtime.Value
		V1 gopurs_runtime.Value
	}{V0: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: (*(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V1} }()
}
