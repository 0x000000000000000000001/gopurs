package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Effect_Exception_Unsafe_unsafeThrowException gopurs_runtime.Value
var once_Effect_Exception_Unsafe_unsafeThrowException sync.Once

func Get_Effect_Exception_Unsafe_unsafeThrowException() gopurs_runtime.Value {
	once_Effect_Exception_Unsafe_unsafeThrowException.Do(func() {
		cache_Effect_Exception_Unsafe_unsafeThrowException = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Unsafe_unsafePerformEffect(), Get_Effect_Exception_throwException())
	})
	return cache_Effect_Exception_Unsafe_unsafeThrowException
}

var cache_Effect_Exception_Unsafe_unsafeThrow gopurs_runtime.Value
var once_Effect_Exception_Unsafe_unsafeThrow sync.Once

func Get_Effect_Exception_Unsafe_unsafeThrow() gopurs_runtime.Value {
	once_Effect_Exception_Unsafe_unsafeThrow.Do(func() {
		cache_Effect_Exception_Unsafe_unsafeThrow = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Effect_Exception_Unsafe_unsafeThrowException(), Get_Effect_Exception_error())
	})
	return cache_Effect_Exception_Unsafe_unsafeThrow
}
