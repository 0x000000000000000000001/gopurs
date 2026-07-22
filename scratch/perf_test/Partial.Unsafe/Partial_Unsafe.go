package Partial_Unsafe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Partial "gopurs/output/Partial"
)

var unsafePartial gopurs_runtime.Value
var once_unsafePartial sync.Once
func Get_unsafePartial() gopurs_runtime.Value {
	once_unsafePartial.Do(func() {
		unsafePartial = Get__unsafePartial()
	})
	return unsafePartial
}

var unsafeCrashWith gopurs_runtime.Value
var once_unsafeCrashWith sync.Once
func Get_unsafeCrashWith() gopurs_runtime.Value {
	once_unsafeCrashWith.Do(func() {
		unsafeCrashWith = gopurs_runtime.Func(func(msg_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), msg_0)
}))
})
	})
	return unsafeCrashWith
}

func Get__unsafePartial() gopurs_runtime.Value {
	return X_UnsafePartial
}
