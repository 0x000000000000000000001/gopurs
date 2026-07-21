package Gopurs_Runtime

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var runtimeGoCode = gopurs_runtime.Str("package gopurs_runtime

const (
	TypeInt = 1
	TypeString = 2
	TypeRecord = 3
	TypeFunc = 4
	TypeConstructor = 5
)

type Value struct {
	Type   uint8
	IntVal int64
	StrVal string
	PtrVal any
}

func Str(v string) Value {
	return Value{Type: TypeString, StrVal: v}
}

func Record(m map[string]Value) Value {
	return Value{Type: TypeRecord, PtrVal: m}
}

func Cons(tag string, args []Value) Value {
	return Value{Type: TypeConstructor, StrVal: tag, PtrVal: args}
}

// Function with 1 arg (curried)
func Func(f func(Value) Value) Value {
	return Value{Type: TypeFunc, PtrVal: f}
}

// Uncurried application helper
func Apply(f Value, arg Value) Value {
	if f.Type != TypeFunc {
		panic("Attempted to apply a non-function")
	}
	fn := f.PtrVal.(func(Value) Value)
	return fn(arg)
}
")
