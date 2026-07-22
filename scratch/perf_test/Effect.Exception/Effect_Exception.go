package Effect_Exception

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Effect "gopurs/output/Effect"
)

var try gopurs_runtime.Value
var once_try sync.Once
func Get_try() gopurs_runtime.Value {
	once_try.Do(func() {
		try = gopurs_runtime.Func(func(action_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_catchException(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect.Get_pureE(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Left"), "value0": x_1}))
})), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect.Get_applyEffect().PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(pkg_Effect.Get_pureE(), pkg_Data_Either.Get_Right())), action_0))
})
	})
	return try
}

var throw gopurs_runtime.Value
var once_throw sync.Once
func Get_throw() gopurs_runtime.Value {
	once_throw.Do(func() {
		throw = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_throwException(), gopurs_runtime.Apply(Get_error(), x_0))
})
	})
	return throw
}

var stack gopurs_runtime.Value
var once_stack sync.Once
func Get_stack() gopurs_runtime.Value {
	once_stack.Do(func() {
		stack = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_stackImpl(), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")}))
	})
	return stack
}

var showError gopurs_runtime.Value
var once_showError sync.Once
func Get_showError() gopurs_runtime.Value {
	once_showError.Do(func() {
		showError = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": Get_showErrorImpl()})
	})
	return showError
}

func Get_catchException() gopurs_runtime.Value {
	return CatchException
}

func Get_error() gopurs_runtime.Value {
	return Error
}

func Get_errorWithCause() gopurs_runtime.Value {
	return ErrorWithCause
}

func Get_errorWithName() gopurs_runtime.Value {
	return ErrorWithName
}

func Get_message() gopurs_runtime.Value {
	return Message
}

func Get_name() gopurs_runtime.Value {
	return Name
}

func Get_showErrorImpl() gopurs_runtime.Value {
	return ShowErrorImpl
}

func Get_stackImpl() gopurs_runtime.Value {
	return StackImpl
}

func Get_throwException() gopurs_runtime.Value {
	return ThrowException
}
