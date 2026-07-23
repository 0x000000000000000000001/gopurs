module Gopurs.Runtime where

runtimeGoCode :: String
runtimeGoCode = """
package gopurs_runtime

import (
	"math"
	"unsafe"
)

const (
	TypeFunc = 1
	TypeFunc2 = 2
	TypeFunc3 = 3
	TypeFunc4 = 4
	TypeFunc5 = 5
	TypeInt = 6
	TypeString = 7
	TypeRecord = 8
	TypeConstructor = 9
	TypeFloat = 10
	TypeBool = 11
	TypeArray = 12
	TypeAny = 13
	TypeFunc6 = 14
	TypeFunc7 = 15
	TypeFunc8 = 16
	TypeFunc9 = 17
	TypeFunc10 = 18
)

// We do not add FloatVal or BoolVal fields to keep the struct size minimal.
// Floats are packed into IntVal using math.Float64bits, and Bools are mapped to 1/0 in IntVal.
// Adding more fields would increase the struct size and reduce pass-by-value performance.
type Value struct {
	Type   uint8
	IntVal int64
	StrVal string
	PtrVal any
	Func   unsafe.Pointer
}

func Str(v string) Value {
	return Value{Type: TypeString, StrVal: v}
}

func Int(v int64) Value {
	return Value{Type: TypeInt, IntVal: v}
}

func Float(v float64) Value {
	return Value{Type: TypeFloat, IntVal: int64(math.Float64bits(v))}
}

func Bool(v bool) Value {
	var i int64 = 0
	if v {
		i = 1
	}
	return Value{Type: TypeBool, IntVal: i}
}

func FloatAdd(a, b Value) Value { return Float(math.Float64frombits(uint64(a.IntVal)) + math.Float64frombits(uint64(b.IntVal))) }
func FloatSub(a, b Value) Value { return Float(math.Float64frombits(uint64(a.IntVal)) - math.Float64frombits(uint64(b.IntVal))) }
func FloatMul(a, b Value) Value { return Float(math.Float64frombits(uint64(a.IntVal)) * math.Float64frombits(uint64(b.IntVal))) }
func FloatDiv(a, b Value) Value { return Float(math.Float64frombits(uint64(a.IntVal)) / math.Float64frombits(uint64(b.IntVal))) }
func FloatNeg(a Value) Value { return Float(-math.Float64frombits(uint64(a.IntVal))) }

func FloatEq(a, b Value) Value { return Bool(math.Float64frombits(uint64(a.IntVal)) == math.Float64frombits(uint64(b.IntVal))) }
func FloatNeq(a, b Value) Value { return Bool(math.Float64frombits(uint64(a.IntVal)) != math.Float64frombits(uint64(b.IntVal))) }
func FloatLt(a, b Value) Value { return Bool(math.Float64frombits(uint64(a.IntVal)) < math.Float64frombits(uint64(b.IntVal))) }
func FloatLte(a, b Value) Value { return Bool(math.Float64frombits(uint64(a.IntVal)) <= math.Float64frombits(uint64(b.IntVal))) }
func FloatGt(a, b Value) Value { return Bool(math.Float64frombits(uint64(a.IntVal)) > math.Float64frombits(uint64(b.IntVal))) }
func FloatGte(a, b Value) Value { return Bool(math.Float64frombits(uint64(a.IntVal)) >= math.Float64frombits(uint64(b.IntVal))) }

func Zshr(a Value, b Value) Value {
	return Int(int64(uint32(a.IntVal) >> uint32(b.IntVal)))
}

func Shl(a Value, b Value) Value {
	return Int(int64(int32(a.IntVal) << uint32(b.IntVal)))
}

func Shr(a Value, b Value) Value {
	return Int(int64(int32(a.IntVal) >> uint32(b.IntVal)))
}

func BitAnd(a Value, b Value) Value {
	return Int(int64(int32(a.IntVal) & int32(b.IntVal)))
}

func BitOr(a Value, b Value) Value {
	return Int(int64(int32(a.IntVal) | int32(b.IntVal)))
}

func BitXor(a Value, b Value) Value {
	return Int(int64(int32(a.IntVal) ^ int32(b.IntVal)))
}

func Array(v []Value) Value {
	return Value{Type: TypeArray, PtrVal: v}
}

type RecordData struct {
	Keys []string
	Vals []Value
}

func RecordDict(keys []string, vals []Value) Value {
	return Value{Type: TypeRecord, PtrVal: RecordData{keys, vals}}
}

type RecordData0 struct{}
func RecordDict0() Value { return Value{Type: TypeRecord, PtrVal: &RecordData0{}} }

type RecordData1 struct { K0 string; V0 Value }
func RecordDict1(k0 string, v0 Value) Value { return Value{Type: TypeRecord, PtrVal: &RecordData1{k0, v0}} }

type RecordData2 struct { K0, K1 string; V0, V1 Value }
func RecordDict2(k0, k1 string, v0, v1 Value) Value { return Value{Type: TypeRecord, PtrVal: &RecordData2{k0, k1, v0, v1}} }

type RecordData3 struct { K0, K1, K2 string; V0, V1, V2 Value }
func RecordDict3(k0, k1, k2 string, v0, v1, v2 Value) Value { return Value{Type: TypeRecord, PtrVal: &RecordData3{k0, k1, k2, v0, v1, v2}} }

type RecordData4 struct { K0, K1, K2, K3 string; V0, V1, V2, V3 Value }
func RecordDict4(k0, k1, k2, k3 string, v0, v1, v2, v3 Value) Value { return Value{Type: TypeRecord, PtrVal: &RecordData4{k0, k1, k2, k3, v0, v1, v2, v3}} }

type RecordData5 struct { K0, K1, K2, K3, K4 string; V0, V1, V2, V3, V4 Value }
func RecordDict5(k0, k1, k2, k3, k4 string, v0, v1, v2, v3, v4 Value) Value { return Value{Type: TypeRecord, PtrVal: &RecordData5{k0, k1, k2, k3, k4, v0, v1, v2, v3, v4}} }

func RecordToMap(obj Value) map[string]Value {
	if m, ok := obj.PtrVal.(map[string]Value); ok {
		res := make(map[string]Value, len(m))
		for k, v := range m { res[k] = v }
		return res
	}
	res := make(map[string]Value)
	switch r := obj.PtrVal.(type) {
	case *RecordData0:
	case *RecordData1: res[r.K0] = r.V0
	case *RecordData2: res[r.K0] = r.V0; res[r.K1] = r.V1
	case *RecordData3: res[r.K0] = r.V0; res[r.K1] = r.V1; res[r.K2] = r.V2
	case *RecordData4: res[r.K0] = r.V0; res[r.K1] = r.V1; res[r.K2] = r.V2; res[r.K3] = r.V3
	case *RecordData5: res[r.K0] = r.V0; res[r.K1] = r.V1; res[r.K2] = r.V2; res[r.K3] = r.V3; res[r.K4] = r.V4
	case RecordData:
		for i, k := range r.Keys { res[k] = r.Vals[i] }
	}
	return res
}

func Record(m map[string]Value) Value {
	keys := make([]string, 0, len(m))
	vals := make([]Value, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	return RecordDict(keys, vals)
}

func RecordGet(obj Value, key string) Value {
    if obj.PtrVal == nil {
        _ = obj.PtrVal.(RecordData) // trigger interface conversion panic for backwards compatibility
    }
    if m, ok := obj.PtrVal.(map[string]Value); ok {
        return m[key]
    }
	switch r := obj.PtrVal.(type) {
	case *RecordData1:
		if r.K0 == key { return r.V0 }
	case *RecordData2:
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
	case *RecordData3:
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
		if r.K2 == key { return r.V2 }
	case *RecordData4:
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
		if r.K2 == key { return r.V2 }
		if r.K3 == key { return r.V3 }
	case *RecordData5:
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
		if r.K2 == key { return r.V2 }
		if r.K3 == key { return r.V3 }
		if r.K4 == key { return r.V4 }
	case RecordData:
		for i, k := range r.Keys {
			if k == key { return r.Vals[i] }
		}
	}
	panic("Key not found in record: " + key)
}

func RecordUpdateDict(orig Value, keys []string, vals []Value) Value {
	m := RecordToMap(orig)
	for i, k := range keys { m[k] = vals[i] }
	return Record(m)
}

func RecordUpdate(orig Value, updates map[string]Value) Value {
	m := RecordToMap(orig)
	for k, v := range updates { m[k] = v }
	return Value{Type: TypeRecord, PtrVal: m}
}


func Cons(tag string, args []Value) Value {
	return Value{Type: TypeConstructor, StrVal: tag, PtrVal: args}
}

// Function with 1 arg (curried)
func Func(f func(Value) Value) Value {
	return Value{Type: TypeFunc, Func: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}

// Function constructors
func Func2(f func(Value, Value) Value) Value { return Value{Type: TypeFunc2, Func: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func3(f func(Value, Value, Value) Value) Value { return Value{Type: TypeFunc3, Func: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func4(f func(Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc4, Func: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func5(f func(Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc5, Func: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func6(f func(Value, Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc6, Func: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func7(f func(Value, Value, Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc7, Func: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func8(f func(Value, Value, Value, Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc8, Func: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func9(f func(Value, Value, Value, Value, Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc9, Func: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func10(f func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc10, Func: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }

func FuncAny(f any) Value {
	return Value{Type: TypeFunc, PtrVal: f}
}


func Apply(f Value, arg Value) Value {
	switch f.Type {
	case TypeFunc:
		return (*(*func(Value) Value)(unsafe.Pointer(&f.Func)))(arg)
	case TypeFunc2:
		fn := *(*func(Value, Value) Value)(unsafe.Pointer(&f.Func))
		return Func(func(a Value) Value { return fn(arg, a) })
	case TypeFunc3:
		fn := *(*func(Value, Value, Value) Value)(unsafe.Pointer(&f.Func))
		return Func2(func(a, b Value) Value { return fn(arg, a, b) })
	case TypeFunc4:
		fn := *(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&f.Func))
		return Func3(func(a, b, c Value) Value { return fn(arg, a, b, c) })
	case TypeFunc5:
		fn := *(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.Func))
		return Func4(func(a, b, c, d Value) Value { return fn(arg, a, b, c, d) })
	default:
		panic("Attempted to apply a non-function")
	}
}

func Apply2(f Value, arg1, arg2 Value) Value {
	switch f.Type {
	case TypeFunc2:
		return (*(*func(Value, Value) Value)(unsafe.Pointer(&f.Func)))(arg1, arg2)
	case TypeFunc3:
		fn := *(*func(Value, Value, Value) Value)(unsafe.Pointer(&f.Func))
		return Func(func(a Value) Value { return fn(arg1, arg2, a) })
	case TypeFunc4:
		fn := *(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&f.Func))
		return Func2(func(a, b Value) Value { return fn(arg1, arg2, a, b) })
	case TypeFunc5:
		fn := *(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.Func))
		return Func3(func(a, b, c Value) Value { return fn(arg1, arg2, a, b, c) })
	}
	return Apply(Apply(f, arg1), arg2)
}

func Apply3(f Value, arg1, arg2, arg3 Value) Value {
	switch f.Type {
	case TypeFunc3:
		return (*(*func(Value, Value, Value) Value)(unsafe.Pointer(&f.Func)))(arg1, arg2, arg3)
	case TypeFunc4:
		fn := *(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&f.Func))
		return Func(func(a Value) Value { return fn(arg1, arg2, arg3, a) })
	case TypeFunc5:
		fn := *(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.Func))
		return Func2(func(a, b Value) Value { return fn(arg1, arg2, arg3, a, b) })
	}
	return Apply(Apply2(f, arg1, arg2), arg3)
}

func Apply4(f Value, arg1, arg2, arg3, arg4 Value) Value {
	switch f.Type {
	case TypeFunc4:
		return (*(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&f.Func)))(arg1, arg2, arg3, arg4)
	case TypeFunc5:
		fn := *(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.Func))
		return Func(func(a Value) Value { return fn(arg1, arg2, arg3, arg4, a) })
	}
	return Apply(Apply3(f, arg1, arg2, arg3), arg4)
}

func Apply5(f Value, arg1, arg2, arg3, arg4, arg5 Value) Value {
	if f.Type == TypeFunc5 {
		return (*(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.Func)))(arg1, arg2, arg3, arg4, arg5)
	}
	return Apply(Apply4(f, arg1, arg2, arg3, arg4), arg5)
}


func ArrayAccess(arr Value, index int) Value {
	return arr.PtrVal.([]Value)[index]
}

func Any(v any) Value {
	return Value{Type: TypeAny, PtrVal: v}
}

func UncurriedApp2(fn Value, a, b Value) Value {
	if fn.Type == TypeFunc2 {
		return (*(*func(Value, Value) Value)(unsafe.Pointer(&fn.Func)))(a, b)
	}
	return Apply(Apply(fn, a), b)
}

func UncurriedApp3(fn Value, a, b, c Value) Value {
	if fn.Type == TypeFunc3 {
		return (*(*func(Value, Value, Value) Value)(unsafe.Pointer(&fn.Func)))(a, b, c)
	}
	return Apply(Apply(Apply(fn, a), b), c)
}

func UncurriedApp4(fn Value, a, b, c, d Value) Value {
	if fn.Type == TypeFunc4 {
		return (*(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&fn.Func)))(a, b, c, d)
	}
	return Apply(Apply(Apply(Apply(fn, a), b), c), d)
}

func UncurriedApp5(fn Value, a, b, c, d, e Value) Value {
	if fn.Type == TypeFunc5 {
		return (*(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&fn.Func)))(a, b, c, d, e)
	}
	return Apply(Apply(Apply(Apply(Apply(fn, a), b), c), d), e)
}

func UncurriedApp(fn Value, args ...Value) Value {
	res := fn
	for _, arg := range args {
		res = Apply(res, arg)
	}
	return res
}

func Unbox[T any](v Value) T {
	var t any = *new(T)
	switch t.(type) {
	case int64: return any(v.IntVal).(T)
	case int: return any(int(v.IntVal)).(T)
	case string: return any(v.StrVal).(T)
	case float64: return any(math.Float64frombits(uint64(v.IntVal))).(T)
	case bool: return any(v.IntVal == 1).(T)
	case Value: return any(v).(T)
	default: return v.PtrVal.(T)
	}
}

func Box[T any](val T) Value {
	switch v := any(val).(type) {
	case int64: return Int(v)
	case int: return Int(int64(v))
	case string: return Str(v)
	case float64: return Value{Type: TypeFloat, IntVal: int64(math.Float64bits(v))}
	case bool: 
		if v { return Value{Type: TypeBool, IntVal: 1} }
		return Value{Type: TypeBool, IntVal: 0}
	case func():
		return Func(func(_ Value) Value { v(); return Value{} })
	case func() bool:
		return Func(func(_ Value) Value { if v() { return Value{Type: TypeBool, IntVal: 1} }; return Value{Type: TypeBool, IntVal: 0} })
	case func() int:
		return Func(func(_ Value) Value { return Int(int64(v())) })
	case func() int64:
		return Func(func(_ Value) Value { return Int(v()) })
	case func() string:
		return Func(func(_ Value) Value { return Str(v()) })
	case func() float64:
		return Func(func(_ Value) Value { return Value{Type: TypeFloat, IntVal: int64(math.Float64bits(v()))} })
	case func() Value:
		return Func(func(_ Value) Value { return v() })
	case func(Value) Value:
		return Func(v)
	case Value: return v
	default: return Any(v)
	}
}

func Wrap0[R any](f func() R) Value {
	return Func(func(_ Value) Value {
		return Box(f())
	})
}

func Wrap1[A, R any](f func(A) R) Value {
	return Func(func(a Value) Value {
		return Box(f(Unbox[A](a)))
	})
}

func Wrap2[A, B, R any](f func(A, B) R) Value {
	return Func2(func(a, b Value) Value {
		return Box(f(Unbox[A](a), Unbox[B](b)))
	})
}

func Wrap3[A, B, C, R any](f func(A, B, C) R) Value {
	return Func3(func(a, b, c Value) Value {
		return Box(f(Unbox[A](a), Unbox[B](b), Unbox[C](c)))
	})
}

func Wrap4[A, B, C, D, R any](f func(A, B, C, D) R) Value {
	return Func4(func(a, b, c, d Value) Value {
		return Box(f(Unbox[A](a), Unbox[B](b), Unbox[C](c), Unbox[D](d)))
	})
}

func Wrap5[A, B, C, D, E, R any](f func(A, B, C, D, E) R) Value {
	return Func5(func(a, b, c, d, e Value) Value {
		return Box(f(Unbox[A](a), Unbox[B](b), Unbox[C](c), Unbox[D](d), Unbox[E](e)))
	})
}
"""
