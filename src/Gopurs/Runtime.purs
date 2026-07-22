module Gopurs.Runtime where

runtimeGoCode :: String
runtimeGoCode = """package gopurs_runtime

import "math"

const (
	TypeInt = 1
	TypeString = 2
	TypeRecord = 3
	TypeFunc = 4
	TypeConstructor = 5
)

// We do not add FloatVal or BoolVal fields to keep the struct size minimal.
// Floats are packed into IntVal using math.Float64bits, and Bools are mapped to 1/0 in IntVal.
// Adding more fields would increase the struct size and reduce pass-by-value performance.
type Value struct {
	Type   uint8
	IntVal int64
	StrVal string
	PtrVal any
}

func Str(v string) Value {
	return Value{Type: TypeString, StrVal: v}
}

func Int(v int) Value {
	return Value{Type: TypeInt, IntVal: int64(v)}
}

func Float(v float64) Value {
	return Value{Type: 7, IntVal: int64(math.Float64bits(v))}
}

func Bool(v bool) Value {
	var i int64 = 0
	if v {
		i = 1
	}
	return Value{Type: 6, IntVal: i}
}

func Array(v []Value) Value {
	return Value{Type: 8, PtrVal: v}
}

func Record(m map[string]Value) Value {
	return Value{Type: TypeRecord, PtrVal: m}
}

func RecordUpdate(orig Value, updates map[string]Value) Value {
	origMap := orig.PtrVal.(map[string]Value)
	newMap := make(map[string]Value, len(origMap)+len(updates))
	for k, v := range origMap {
		newMap[k] = v
	}
	for k, v := range updates {
		newMap[k] = v
	}
	return Record(newMap)
}

func Cons(tag string, args []Value) Value {
	return Value{Type: TypeConstructor, StrVal: tag, PtrVal: args}
}

// Function with 1 arg (curried)
func Func(f func(Value) Value) Value {
	return Value{Type: TypeFunc, PtrVal: f}
}

func FuncAny(f any) Value {
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

func ArrayAccess(arr Value, index int) Value {
	return arr.PtrVal.([]Value)[index]
}

func Any(v any) Value {
	return Value{Type: 9, PtrVal: v}
}
"""
