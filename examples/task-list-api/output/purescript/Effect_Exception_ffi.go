package purescript

import "gopurs/output/gopurs_runtime"

import (
	"errors"
	"fmt"
)

func Effect_Exception_Error(msg string) error       { return errors.New(msg) }
func Effect_Exception_Message(e error) string       { return e.Error() }
func Effect_Exception_Name(e error) string          { return "Error" }
func Effect_Exception_ShowErrorImpl(e error) string { return e.Error() }
func Effect_Exception_StackImpl(just func(string) interface{}, nothing interface{}, e error) interface{} {
	return nothing
}

func Effect_Exception_ThrowException(e error, _ interface{}) interface{} {
	panic(e)
}

func Effect_Exception_CatchException(c func(error) func(interface{}) interface{}, t func(interface{}) interface{}, _ interface{}) (res interface{}) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				res = c(err)(nil)
			} else {
				res = c(fmt.Errorf("%v", r))(nil)
			}
		}
	}()
	return t(nil)
}

func Effect_Exception_ErrorWithCause(msg string, cause error) error { return errors.New(msg) }
func Effect_Exception_ErrorWithName(name string, msg string) error  { return errors.New(msg) }

// --- Auto-generated FFI wrappers ---
var _Gopurs_Effect_Exception_CatchException = // TAST: (ForAll [a] (Func [(Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Effect"] [(TypeVar a)])), (ADT ["Effect","Effect"] [(TypeVar a)])] (ADT ["Effect","Effect"] [(TypeVar a)])))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 error) func(any) any {
		inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
	}
	go_arg1 := func(p0_0 any) any {
		return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
	}
	go_arg2 := arg2
	go_res := Effect_Exception_CatchException(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Exception_Error = // TAST: (Func [String] (ADT ["Effect","Exception","Error"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Effect_Exception_Error(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Exception_ErrorWithCause = // TAST: (Func [String, (ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Exception","Error"] []))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[error](arg1)
	go_res := Effect_Exception_ErrorWithCause(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Exception_ErrorWithName = // TAST: (Func [String, String] (ADT ["Effect","Exception","Error"] []))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Effect_Exception_ErrorWithName(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Exception_Message = // TAST: (Func [(ADT ["Effect","Exception","Error"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[error](arg0)
	go_res := Effect_Exception_Message(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Exception_Name = // TAST: (Func [(ADT ["Effect","Exception","Error"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[error](arg0)
	go_res := Effect_Exception_Name(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Exception_ShowErrorImpl = // TAST: (Func [(ADT ["Effect","Exception","Error"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[error](arg0)
	go_res := Effect_Exception_ShowErrorImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Exception_StackImpl = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Effect","Exception","Error"] [])] (ADT ["Data","Maybe","Maybe"] [String]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[error](arg2)
	go_res := Effect_Exception_StackImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Exception_ThrowException = // TAST: (ForAll [a] (Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Effect"] [(TypeVar a)])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[error](arg0)
	go_arg1 := arg1
	go_res := Effect_Exception_ThrowException(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})

// --- Boxed-result FFI workers ---
func Effect_Exception_StackImpl_nativeWorker(arg0 gopurs_runtime.Value, arg1 any, arg2 error) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) any {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Effect_Exception_StackImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
}
