module Gopurs.Runtime where

runtimeGoCode :: String
runtimeGoCode = """package gopurs_runtime

const (
	TypeInt = 1
	TypeString = 2
	TypeRecord = 3
	TypeFunc = 4
	TypeConstructor = 5
)

type Value struct {
	Type   uint8
	IntVal   int
	FloatVal float64
	BoolVal  bool
	StrVal   string
	ArrayVal []Value
	PtrVal   any
}

func Str(v string) Value {
	return Value{Type: TypeString, StrVal: v}
}

func Int(v int) Value {
	return Value{Type: TypeInt, IntVal: v}
}

func Float(v float64) Value {
	return Value{Type: 7, FloatVal: v}
}

func Bool(v bool) Value {
	return Value{Type: 6, BoolVal: v}
}

func Array(v []Value) Value {
	return Value{Type: 8, ArrayVal: v}
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
"""
