module Gopurs.Runtime where

runtimeGoCode :: String
runtimeGoCode = """
package gopurs_runtime

import "math"

const (
	TypeInt = 1
	TypeString = 2
	TypeRecord = 3
	TypeFunc = 4
	TypeFunc2 = 10
	TypeFunc3 = 11
	TypeFunc4 = 12
	TypeFunc5 = 13
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

func Int(v int64) Value {
	return Value{Type: TypeInt, IntVal: v}
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
	return Value{Type: 8, PtrVal: v}
}

type RecordData struct {
	Keys []string
	Vals []Value
}

func RecordDict(keys []string, vals []Value) Value {
	return Value{Type: TypeRecord, PtrVal: RecordData{keys, vals}}
}

type RecordData0 struct{}
func RecordDict0() Value { return Value{Type: TypeRecord, PtrVal: RecordData0{}} }

type RecordData1 struct { K0 string; V0 Value }
func RecordDict1(k0 string, v0 Value) Value { return Value{Type: TypeRecord, PtrVal: RecordData1{k0, v0}} }

type RecordData2 struct { K0, K1 string; V0, V1 Value }
func RecordDict2(k0, k1 string, v0, v1 Value) Value { return Value{Type: TypeRecord, PtrVal: RecordData2{k0, k1, v0, v1}} }

type RecordData3 struct { K0, K1, K2 string; V0, V1, V2 Value }
func RecordDict3(k0, k1, k2 string, v0, v1, v2 Value) Value { return Value{Type: TypeRecord, PtrVal: RecordData3{k0, k1, k2, v0, v1, v2}} }

type RecordData4 struct { K0, K1, K2, K3 string; V0, V1, V2, V3 Value }
func RecordDict4(k0, k1, k2, k3 string, v0, v1, v2, v3 Value) Value { return Value{Type: TypeRecord, PtrVal: RecordData4{k0, k1, k2, k3, v0, v1, v2, v3}} }

type RecordData5 struct { K0, K1, K2, K3, K4 string; V0, V1, V2, V3, V4 Value }
func RecordDict5(k0, k1, k2, k3, k4 string, v0, v1, v2, v3, v4 Value) Value { return Value{Type: TypeRecord, PtrVal: RecordData5{k0, k1, k2, k3, k4, v0, v1, v2, v3, v4}} }

func RecordToMap(obj Value) map[string]Value {
	if m, ok := obj.PtrVal.(map[string]Value); ok {
		res := make(map[string]Value, len(m))
		for k, v := range m { res[k] = v }
		return res
	}
	res := make(map[string]Value)
	switch r := obj.PtrVal.(type) {
	case RecordData0:
	case RecordData1: res[r.K0] = r.V0
	case RecordData2: res[r.K0] = r.V0; res[r.K1] = r.V1
	case RecordData3: res[r.K0] = r.V0; res[r.K1] = r.V1; res[r.K2] = r.V2
	case RecordData4: res[r.K0] = r.V0; res[r.K1] = r.V1; res[r.K2] = r.V2; res[r.K3] = r.V3
	case RecordData5: res[r.K0] = r.V0; res[r.K1] = r.V1; res[r.K2] = r.V2; res[r.K3] = r.V3; res[r.K4] = r.V4
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
	case RecordData1:
		if r.K0 == key { return r.V0 }
	case RecordData2:
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
	case RecordData3:
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
		if r.K2 == key { return r.V2 }
	case RecordData4:
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
		if r.K2 == key { return r.V2 }
		if r.K3 == key { return r.V3 }
	case RecordData5:
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
	return Value{Type: TypeFunc, PtrVal: f}
}


// Function constructors
func Func2(f func(Value, Value) Value) Value { return Value{Type: TypeFunc2, PtrVal: f} }
func Func3(f func(Value, Value, Value) Value) Value { return Value{Type: TypeFunc3, PtrVal: f} }
func Func4(f func(Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc4, PtrVal: f} }
func Func5(f func(Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc5, PtrVal: f} }

func FuncAny(f any) Value {
	return Value{Type: TypeFunc, PtrVal: f}
}


func Apply(f Value, arg Value) Value {
	switch f.Type {
	case TypeFunc:
		return f.PtrVal.(func(Value) Value)(arg)
	case TypeFunc2:
		fn := f.PtrVal.(func(Value, Value) Value)
		return Func(func(a Value) Value { return fn(arg, a) })
	case TypeFunc3:
		fn := f.PtrVal.(func(Value, Value, Value) Value)
		return Func2(func(a, b Value) Value { return fn(arg, a, b) })
	case TypeFunc4:
		fn := f.PtrVal.(func(Value, Value, Value, Value) Value)
		return Func3(func(a, b, c Value) Value { return fn(arg, a, b, c) })
	case TypeFunc5:
		fn := f.PtrVal.(func(Value, Value, Value, Value, Value) Value)
		return Func4(func(a, b, c, d Value) Value { return fn(arg, a, b, c, d) })
	default:
		panic("Attempted to apply a non-function")
	}
}

func Apply2(f Value, arg1, arg2 Value) Value {
	if f.Type == TypeFunc2 {
		return f.PtrVal.(func(Value, Value) Value)(arg1, arg2)
	}
	return Apply(Apply(f, arg1), arg2)
}

func Apply3(f Value, arg1, arg2, arg3 Value) Value {
	if f.Type == TypeFunc3 {
		return f.PtrVal.(func(Value, Value, Value) Value)(arg1, arg2, arg3)
	}
	return Apply(Apply2(f, arg1, arg2), arg3)
}

func Apply4(f Value, arg1, arg2, arg3, arg4 Value) Value {
	if f.Type == TypeFunc4 {
		return f.PtrVal.(func(Value, Value, Value, Value) Value)(arg1, arg2, arg3, arg4)
	}
	return Apply(Apply3(f, arg1, arg2, arg3), arg4)
}

func Apply5(f Value, arg1, arg2, arg3, arg4, arg5 Value) Value {
	if f.Type == TypeFunc5 {
		return f.PtrVal.(func(Value, Value, Value, Value, Value) Value)(arg1, arg2, arg3, arg4, arg5)
	}
	return Apply(Apply4(f, arg1, arg2, arg3, arg4), arg5)
}


func ArrayAccess(arr Value, index int) Value {
	return arr.PtrVal.([]Value)[index]
}

func Any(v any) Value {
	return Value{Type: 9, PtrVal: v}
}

func UncurriedApp2(fn Value, a, b Value) Value {
	if f, ok := fn.PtrVal.(func(Value, Value) Value); ok {
		return f(a, b)
	}
	return Apply(Apply(fn, a), b)
}

func UncurriedApp3(fn Value, a, b, c Value) Value {
	if f, ok := fn.PtrVal.(func(Value, Value, Value) Value); ok {
		return f(a, b, c)
	}
	return Apply(Apply(Apply(fn, a), b), c)
}

func UncurriedApp4(fn Value, a, b, c, d Value) Value {
	if f, ok := fn.PtrVal.(func(Value, Value, Value, Value) Value); ok {
		return f(a, b, c, d)
	}
	return Apply(Apply(Apply(Apply(fn, a), b), c), d)
}

func UncurriedApp5(fn Value, a, b, c, d, e Value) Value {
	if f, ok := fn.PtrVal.(func(Value, Value, Value, Value, Value) Value); ok {
		return f(a, b, c, d, e)
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
"""
