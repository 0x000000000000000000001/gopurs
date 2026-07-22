package Effect_Exception_Unsafe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Exception "gopurs/output/Effect.Exception"
	pkg_Effect_Unsafe "gopurs/output/Effect.Unsafe"
)

var unsafeThrowException gopurs_runtime.Value
var once_unsafeThrowException sync.Once
func Get_unsafeThrowException() gopurs_runtime.Value {
	once_unsafeThrowException.Do(func() {
		unsafeThrowException = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), gopurs_runtime.Apply(pkg_Effect_Exception.Get_throwException(), x_0))
})
	})
	return unsafeThrowException
}

var unsafeThrow gopurs_runtime.Value
var once_unsafeThrow sync.Once
func Get_unsafeThrow() gopurs_runtime.Value {
	once_unsafeThrow.Do(func() {
		unsafeThrow = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), gopurs_runtime.Apply(pkg_Effect_Exception.Get_throwException(), gopurs_runtime.Apply(pkg_Effect_Exception.Get_error(), x_0)))
})
	})
	return unsafeThrow
}


