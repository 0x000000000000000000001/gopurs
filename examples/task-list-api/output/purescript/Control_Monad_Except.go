package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Except_compose gopurs_runtime.Value
var once_Control_Monad_Except_compose sync.Once

func Get_Control_Monad_Except_compose() gopurs_runtime.Value {
	once_Control_Monad_Except_compose.Do(func() {
		cache_Control_Monad_Except_compose = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Except_compose
}

var cache_Control_Monad_Except_compose1 gopurs_runtime.Value
var once_Control_Monad_Except_compose1 sync.Once

func Get_Control_Monad_Except_compose1() gopurs_runtime.Value {
	once_Control_Monad_Except_compose1.Do(func() {
		cache_Control_Monad_Except_compose1 = Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))
	})
	return cache_Control_Monad_Except_compose1
}

var cache_Control_Monad_Except_unwrap gopurs_runtime.Value
var once_Control_Monad_Except_unwrap sync.Once

func Get_Control_Monad_Except_unwrap() gopurs_runtime.Value {
	once_Control_Monad_Except_unwrap.Do(func() {
		cache_Control_Monad_Except_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Control_Monad_Except_unwrap
}

var cache_Control_Monad_Except_withExcept gopurs_runtime.Value
var once_Control_Monad_Except_withExcept sync.Once

func Get_Control_Monad_Except_withExcept() gopurs_runtime.Value {
	once_Control_Monad_Except_withExcept.Do(func() {
		cache_Control_Monad_Except_withExcept = gopurs_runtime.Apply(Get_Control_Monad_Except_Trans_withExceptT(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Identity_functorIdentity()))})
	})
	return cache_Control_Monad_Except_withExcept
}

var cache_Control_Monad_Except_runExcept gopurs_runtime.Value
var once_Control_Monad_Except_runExcept sync.Once

func Get_Control_Monad_Except_runExcept() gopurs_runtime.Value {
	once_Control_Monad_Except_runExcept.Do(func() {
		cache_Control_Monad_Except_runExcept = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), Get_Control_Monad_Except_Trans_runExceptT())
	})
	return cache_Control_Monad_Except_runExcept
}

var cache_Control_Monad_Except_runExcept__1256907226 gopurs_runtime.Value
var once_Control_Monad_Except_runExcept__1256907226 sync.Once

func Get_Control_Monad_Except_runExcept__1256907226() gopurs_runtime.Value {
	once_Control_Monad_Except_runExcept__1256907226.Do(func() {
		cache_Control_Monad_Except_runExcept__1256907226 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Control_Monad_Except_runExcept__1256907226(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: func() gopurs_runtime.Value {
						orig := _v.V1
						_ = orig
						return gopurs_runtime.RecordDict3("completed", "id", "title", gopurs_runtime.Bool(orig.completed), gopurs_runtime.Str(orig.id), gopurs_runtime.Str(orig.title))
					}()})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}()
		})
	})
	return cache_Control_Monad_Except_runExcept__1256907226
}

var cache_Control_Monad_Except_mapExcept gopurs_runtime.Value
var once_Control_Monad_Except_mapExcept sync.Once

func Get_Control_Monad_Except_mapExcept() gopurs_runtime.Value {
	once_Control_Monad_Except_mapExcept.Do(func() {
		cache_Control_Monad_Except_mapExcept = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Control_Monad_Except_mapExcept(f_0_box)
		})
	})
	return cache_Control_Monad_Except_mapExcept
}

func Call_Control_Monad_Except_runExcept__1256907226(__eta_norm_0_0_loop gopurs_runtime.Value) struct {
	V0 gopurs_runtime.Value
	V1 struct {
		completed bool
		id        string
		title     string
		/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
	}
	V2 bool
} {
runExcept__1256907226:
	for {
		if false {
			continue runExcept__1256907226
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return func() struct {
			V0 gopurs_runtime.Value
			V1 struct {
				completed bool
				id        string
				title     string
				/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
			}
			V2 bool
		} {
			_v := gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), Get_Control_Monad_Except_Trans_runExceptT(), __eta_norm_0_0)
			if _v.Type == 9 && _v.IntVal == 2465973597 && _v.UnsafePtr != nil {
				return struct {
					V0 gopurs_runtime.Value
					V1 struct {
						completed bool
						id        string
						title     string
						/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
					}
					V2 bool
				}{V0: gopurs_runtime.Value{}, V1: func() struct {
					completed bool
					id        string
					title     string
					/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
				} {
					orig := (*(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0
					_ = orig
					clone := struct {
						completed bool
						id        string
						title     string
						/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
					}{}
					clone.completed = (gopurs_runtime.RecordGet(orig, "completed").IntVal) != (0)
					clone.id = gopurs_runtime.RecordGet(orig, "id").StrVal()
					clone.title = gopurs_runtime.RecordGet(orig, "title").StrVal()
					return clone
				}(), V2: true}
			}
			return struct {
				V0 gopurs_runtime.Value
				V1 struct {
					completed bool
					id        string
					title     string
					/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
				}
				V2 bool
			}{V0: (*(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: struct {
				completed bool
				id        string
				title     string
				/*gopurs:record-v1:3:0063006f006d0070006c0065007400650064,00690064,007400690074006c0065*/
			}{}, V2: false}
		}()
	}
}

func Call_Control_Monad_Except_mapExcept(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var f_0 gopurs_runtime.Value = f_0_loop
	_ = f_0
	// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
	__local_var_1_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Identity_Identity(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
	_ = __local_var_1_0
	return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(__local_var_1_0, v_2)
	})
}
