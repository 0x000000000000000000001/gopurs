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
	TypeRecord0 = 19
	TypeRecord1 = 20
	TypeRecord2 = 21
	TypeRecord3 = 22
	TypeRecord4 = 23
	TypeRecord5 = 24
	TypeRecordData = 25
	TypeFunc11 = 26
)

// We do not add FloatVal or BoolVal fields to keep the struct size minimal.
// Floats are packed into IntVal using math.Float64bits, and Bools are mapped to 1/0 in IntVal.
// Adding more fields would increase the struct size and reduce pass-by-value performance.
type Value struct {
	Type      int
	IntVal    int64
	UnsafePtr unsafe.Pointer

}

// IncRef increments the Rc field if the value is a pointer ADT.
// It assumes that all pointer ADTs have `Rc uint32` as their very first field.
func IncRef(v Value) Value {
	if v.Type == TypeConstructor {
		ptr := (*struct{ Rc uint32 })(v.UnsafePtr)
		if ptr != nil {
			ptr.Rc++
		}
	}
	return v
}

func (v Value) FloatVal() float64 {
	return math.Float64frombits(uint64(v.IntVal))
}

func Str(v string) Value {
	return Value{Type: TypeString, UnsafePtr: unsafe.Pointer(&v)}
}

func (v Value) StrVal() string {
	return *(*string)(v.UnsafePtr)
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
	return Value{Type: TypeArray, UnsafePtr: unsafe.Pointer(&v)}
}

func (v Value) PtrVal() any {
	return *(*any)(v.UnsafePtr)
}

func (v Value) BoolVal() bool {
	return v.IntVal != 0
}

func (v Value) AnyVal() any {
	switch v.Type {
	case TypeInt: return v.IntVal
	case TypeFloat: return math.Float64frombits(uint64(v.IntVal))
	case TypeString: return *(*string)(v.UnsafePtr)
	case TypeBool: return v.IntVal != 0
	case TypeArray: return *(*[]Value)(v.UnsafePtr)
	case TypeRecord: return *(*map[string]Value)(v.UnsafePtr)
	case TypeAny: return v.PtrVal()
	case TypeConstructor: return v
	default:
		if v.Type >= TypeRecord0 && v.Type <= TypeRecordData {
			return RecordToMap(v)
		}
		if (v.Type >= TypeFunc && v.Type <= TypeFunc5) || (v.Type >= TypeFunc6 && v.Type <= TypeFunc11) {
			return v
		}
		if v.UnsafePtr != nil {
			return v
		}
		return nil
	}
}

type RecordData struct {
	Keys []string
	Vals []Value
}

func RecordDict(keys []string, vals []Value) Value {
	r := RecordData{keys, vals}
	return Value{Type: TypeRecordData, UnsafePtr: unsafe.Pointer(&r)}
}

type RecordData0 struct{}
func RecordDict0() Value { return Value{Type: TypeRecord0, UnsafePtr: unsafe.Pointer(&RecordData0{})} }

type RecordData1 struct { K0 string; V0 Value }
func RecordDict1(k0 string, v0 Value) Value { return Value{Type: TypeRecord1, UnsafePtr: unsafe.Pointer(&RecordData1{k0, v0})} }

type RecordData2 struct { K0, K1 string; V0, V1 Value }
func RecordDict2(k0, k1 string, v0, v1 Value) Value { return Value{Type: TypeRecord2, UnsafePtr: unsafe.Pointer(&RecordData2{k0, k1, v0, v1})} }

type RecordData3 struct { K0, K1, K2 string; V0, V1, V2 Value }
func RecordDict3(k0, k1, k2 string, v0, v1, v2 Value) Value { return Value{Type: TypeRecord3, UnsafePtr: unsafe.Pointer(&RecordData3{k0, k1, k2, v0, v1, v2})} }

type RecordData4 struct { K0, K1, K2, K3 string; V0, V1, V2, V3 Value }
func RecordDict4(k0, k1, k2, k3 string, v0, v1, v2, v3 Value) Value { return Value{Type: TypeRecord4, UnsafePtr: unsafe.Pointer(&RecordData4{k0, k1, k2, k3, v0, v1, v2, v3})} }

type RecordData5 struct { K0, K1, K2, K3, K4 string; V0, V1, V2, V3, V4 Value }
func RecordDict5(k0, k1, k2, k3, k4 string, v0, v1, v2, v3, v4 Value) Value { return Value{Type: TypeRecord5, UnsafePtr: unsafe.Pointer(&RecordData5{k0, k1, k2, k3, k4, v0, v1, v2, v3, v4})} }

func RecordToMap(obj Value) map[string]Value {
	if obj.Type == TypeRecord {
		m := *(*map[string]Value)(obj.UnsafePtr)
		res := make(map[string]Value, len(m))
		for k, v := range m { res[k] = v }
		return res
	}
	res := make(map[string]Value)
	switch obj.Type {
	case TypeRecord0:
	case TypeRecord1:
		r := (*RecordData1)(obj.UnsafePtr); res[r.K0] = r.V0
	case TypeRecord2:
		r := (*RecordData2)(obj.UnsafePtr); res[r.K0] = r.V0; res[r.K1] = r.V1
	case TypeRecord3:
		r := (*RecordData3)(obj.UnsafePtr); res[r.K0] = r.V0; res[r.K1] = r.V1; res[r.K2] = r.V2
	case TypeRecord4:
		r := (*RecordData4)(obj.UnsafePtr); res[r.K0] = r.V0; res[r.K1] = r.V1; res[r.K2] = r.V2; res[r.K3] = r.V3
	case TypeRecord5:
		r := (*RecordData5)(obj.UnsafePtr); res[r.K0] = r.V0; res[r.K1] = r.V1; res[r.K2] = r.V2; res[r.K3] = r.V3; res[r.K4] = r.V4
	case TypeRecordData:
		r := (*RecordData)(obj.UnsafePtr)
		for i, k := range r.Keys { res[k] = r.Vals[i] }
	}
	return res
}

func Record(m map[string]Value) Value {
	return Value{Type: TypeRecord, UnsafePtr: unsafe.Pointer(&m)}
}

func RecordGet(obj Value, key string) Value {
	if obj.Type == 0 {
		panic("Attempt to read property '" + key + "' on uninitialized value")
	}
    switch obj.Type {
    case TypeRecord: return (*(*map[string]Value)(obj.UnsafePtr))[key]
	case TypeRecord1:
		r := (*RecordData1)(obj.UnsafePtr)
		if r.K0 == key { return r.V0 }
	case TypeRecord2:
		r := (*RecordData2)(obj.UnsafePtr)
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
	case TypeRecord3:
		r := (*RecordData3)(obj.UnsafePtr)
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
		if r.K2 == key { return r.V2 }
	case TypeRecord4:
		r := (*RecordData4)(obj.UnsafePtr)
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
		if r.K2 == key { return r.V2 }
		if r.K3 == key { return r.V3 }
	case TypeRecord5:
		r := (*RecordData5)(obj.UnsafePtr)
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
		if r.K2 == key { return r.V2 }
		if r.K3 == key { return r.V3 }
		if r.K4 == key { return r.V4 }
	case TypeRecordData:
		r := (*RecordData)(obj.UnsafePtr)
		for i, k := range r.Keys {
			if k == key { return r.Vals[i] }
		}
	}
	panic("Key not found in record: " + key)
}

func RecordUpdateDict(orig Value, keys []string, vals []Value) Value {
	switch orig.Type {
	case TypeRecord1:
		clone := *(*RecordData1)(orig.UnsafePtr)
		for i, k := range keys {
			if clone.K0 == k { clone.V0 = vals[i] }
		}
		return Value{Type: TypeRecord1, UnsafePtr: unsafe.Pointer(&clone)}
	case TypeRecord2:
		clone := *(*RecordData2)(orig.UnsafePtr)
		for i, k := range keys {
			if clone.K0 == k { clone.V0 = vals[i] } else if clone.K1 == k { clone.V1 = vals[i] }
		}
		return Value{Type: TypeRecord2, UnsafePtr: unsafe.Pointer(&clone)}
	case TypeRecord3:
		clone := *(*RecordData3)(orig.UnsafePtr)
		for i, k := range keys {
			if clone.K0 == k { clone.V0 = vals[i] } else if clone.K1 == k { clone.V1 = vals[i] } else if clone.K2 == k { clone.V2 = vals[i] }
		}
		return Value{Type: TypeRecord3, UnsafePtr: unsafe.Pointer(&clone)}
	case TypeRecord4:
		clone := *(*RecordData4)(orig.UnsafePtr)
		for i, k := range keys {
			if clone.K0 == k { clone.V0 = vals[i] } else if clone.K1 == k { clone.V1 = vals[i] } else if clone.K2 == k { clone.V2 = vals[i] } else if clone.K3 == k { clone.V3 = vals[i] }
		}
		return Value{Type: TypeRecord4, UnsafePtr: unsafe.Pointer(&clone)}
	case TypeRecord5:
		clone := *(*RecordData5)(orig.UnsafePtr)
		for i, k := range keys {
			if clone.K0 == k { clone.V0 = vals[i] } else if clone.K1 == k { clone.V1 = vals[i] } else if clone.K2 == k { clone.V2 = vals[i] } else if clone.K3 == k { clone.V3 = vals[i] } else if clone.K4 == k { clone.V4 = vals[i] }
		}
		return Value{Type: TypeRecord5, UnsafePtr: unsafe.Pointer(&clone)}
	case TypeRecordData:
		r := (*RecordData)(orig.UnsafePtr)
		newVals := make([]Value, len(r.Vals))
		copy(newVals, r.Vals)
		for i, k := range keys {
			for j, rk := range r.Keys {
				if rk == k { newVals[j] = vals[i]; break }
			}
		}
		newR := RecordData{Keys: r.Keys, Vals: newVals}
		return Value{Type: TypeRecordData, UnsafePtr: unsafe.Pointer(&newR)}
	}
	m := RecordToMap(orig)
	for i, k := range keys { m[k] = vals[i] }
	return Record(m)
}

func RecordUpdate1(orig Value, k1 string, v1 Value) Value {
	return RecordUpdateDict(orig, []string{k1}, []Value{v1})
}

func RecordUpdate2(orig Value, k1 string, v1 Value, k2 string, v2 Value) Value {
	return RecordUpdateDict(orig, []string{k1, k2}, []Value{v1, v2})
}

func RecordUpdate3(orig Value, k1 string, v1 Value, k2 string, v2 Value, k3 string, v3 Value) Value {
	return RecordUpdateDict(orig, []string{k1, k2, k3}, []Value{v1, v2, v3})
}

func RecordUpdate(orig Value, updates map[string]Value) Value {
	keys := make([]string, 0, len(updates))
	vals := make([]Value, 0, len(updates))
	for k, v := range updates {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	return RecordUpdateDict(orig, keys, vals)
}


type ConstructorData []Value

func Constructor(tag string, args []Value) Value {
	ptr := unsafe.Pointer(nil)
	if len(args) > 0 { ptr = unsafe.Pointer(&args[0]) }
	return Value{Type: TypeConstructor, UnsafePtr: ptr}
}

type ConstructorData0 struct{}
func Constructor0(tag string) Value { return Value{Type: TypeConstructor, UnsafePtr: unsafe.Pointer(&ConstructorData0{})} }

type ConstructorData1 struct { V0 Value }
func Constructor1(tag string, v0 Value) Value { return Value{Type: TypeConstructor, UnsafePtr: unsafe.Pointer(&ConstructorData1{v0})} }

type ConstructorData2 struct { V0, V1 Value }
func Constructor2(tag string, v0, v1 Value) Value { return Value{Type: TypeConstructor, UnsafePtr: unsafe.Pointer(&ConstructorData2{v0, v1})} }

type ConstructorData3 struct { V0, V1, V2 Value }
func Constructor3(tag string, v0, v1, v2 Value) Value { return Value{Type: TypeConstructor, UnsafePtr: unsafe.Pointer(&ConstructorData3{v0, v1, v2})} }

type ConstructorData4 struct { V0, V1, V2, V3 Value }
func Constructor4(tag string, v0, v1, v2, v3 Value) Value { return Value{Type: TypeConstructor, UnsafePtr: unsafe.Pointer(&ConstructorData4{v0, v1, v2, v3})} }

type ConstructorData5 struct { V0, V1, V2, V3, V4 Value }
func Constructor5(tag string, v0, v1, v2, v3, v4 Value) Value { return Value{Type: TypeConstructor, UnsafePtr: unsafe.Pointer(&ConstructorData5{v0, v1, v2, v3, v4})} }

// Function with 1 arg (curried)
func Func(f func(Value) Value) Value {
	return Value{Type: TypeFunc, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}

// Function constructors
func Func2(f func(Value, Value) Value) Value { return Value{Type: TypeFunc2, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func3(f func(Value, Value, Value) Value) Value { return Value{Type: TypeFunc3, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func4(f func(Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc4, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func5(f func(Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc5, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func6(f func(Value, Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc6, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func7(f func(Value, Value, Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc7, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func8(f func(Value, Value, Value, Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc8, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func9(f func(Value, Value, Value, Value, Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc9, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func10(f func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc10, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }
func Func11(f func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc11, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))} }

func FuncAny(f any) Value {
	ptr := new(any)
	*ptr = f
	return Value{Type: TypeFunc, UnsafePtr: unsafe.Pointer(ptr)}
}


func Apply(f Value, arg Value) Value {
	switch f.Type {
	case TypeFunc:
		return (*(*func(Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg)
	case TypeFunc2:
		fn := *(*func(Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg, a) })
	case TypeFunc3:
		fn := *(*func(Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func2(func(a, b Value) Value { return fn(arg, a, b) })
	case TypeFunc4:
		fn := *(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func3(func(a, b, c Value) Value { return fn(arg, a, b, c) })
	case TypeFunc5:
		fn := *(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func4(func(a, b, c, d Value) Value { return fn(arg, a, b, c, d) })
	case TypeFunc6:
		fn := *(*func(Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func5(func(a, b, c, d, e Value) Value { return fn(arg, a, b, c, d, e) })
	case TypeFunc7:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func6(func(a, b, c, d, e, f Value) Value { return fn(arg, a, b, c, d, e, f) })
	case TypeFunc8:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func7(func(a, b, c, d, e, f, g Value) Value { return fn(arg, a, b, c, d, e, f, g) })
	case TypeFunc9:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func8(func(a, b, c, d, e, f, g, h Value) Value { return fn(arg, a, b, c, d, e, f, g, h) })
	case TypeFunc10:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func9(func(a, b, c, d, e, f, g, h, i Value) Value { return fn(arg, a, b, c, d, e, f, g, h, i) })
	case TypeFunc11:
		fn := *(*func(Value, Value, Value, Value, Value, Value, Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func10(func(a, b, c, d, e, f, g, h, i, j Value) Value { return fn(arg, a, b, c, d, e, f, g, h, i, j) })
	default:
		panic("Attempted to apply a non-function")
	}
}

func Apply2(f Value, arg1, arg2 Value) Value {
	switch f.Type {
	case TypeFunc2:
		return (*(*func(Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg1, arg2)
	case TypeFunc3:
		fn := *(*func(Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg1, arg2, a) })
	case TypeFunc4:
		fn := *(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func2(func(a, b Value) Value { return fn(arg1, arg2, a, b) })
	case TypeFunc5:
		fn := *(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func3(func(a, b, c Value) Value { return fn(arg1, arg2, a, b, c) })
	}
	return Apply(Apply(f, arg1), arg2)
}

func Apply3(f Value, arg1, arg2, arg3 Value) Value {
	switch f.Type {
	case TypeFunc3:
		return (*(*func(Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg1, arg2, arg3)
	case TypeFunc4:
		fn := *(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg1, arg2, arg3, a) })
	case TypeFunc5:
		fn := *(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func2(func(a, b Value) Value { return fn(arg1, arg2, arg3, a, b) })
	}
	return Apply(Apply2(f, arg1, arg2), arg3)
}

func Apply4(f Value, arg1, arg2, arg3, arg4 Value) Value {
	switch f.Type {
	case TypeFunc4:
		return (*(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg1, arg2, arg3, arg4)
	case TypeFunc5:
		fn := *(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
		return Func(func(a Value) Value { return fn(arg1, arg2, arg3, arg4, a) })
	}
	return Apply(Apply3(f, arg1, arg2, arg3), arg4)
}

func Apply5(f Value, arg1, arg2, arg3, arg4, arg5 Value) Value {
	if f.Type == TypeFunc5 {
		return (*(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr)))(arg1, arg2, arg3, arg4, arg5)
	}
	return Apply(Apply4(f, arg1, arg2, arg3, arg4), arg5)
}


func ArrayAccess(arr Value, index int) Value {
	return (*(*[]Value)(arr.UnsafePtr))[index]
}

func ArrayLength(arr Value) int {
	return len(*(*[]Value)(arr.UnsafePtr))
}

func Any(v any) Value {
	if val, ok := v.(Value); ok {
		return val
	}
	return Value{Type: TypeAny, UnsafePtr: unsafe.Pointer(&v)}
}

func UncurriedApp2(fn Value, a, b Value) Value {
	if fn.Type == TypeFunc2 {
		return (*(*func(Value, Value) Value)(unsafe.Pointer(&fn.UnsafePtr)))(a, b)
	}
	return Apply(Apply(fn, a), b)
}

func UncurriedApp3(fn Value, a, b, c Value) Value {
	if fn.Type == TypeFunc3 {
		return (*(*func(Value, Value, Value) Value)(unsafe.Pointer(&fn.UnsafePtr)))(a, b, c)
	}
	return Apply(Apply(Apply(fn, a), b), c)
}

func UncurriedApp4(fn Value, a, b, c, d Value) Value {
	if fn.Type == TypeFunc4 {
		return (*(*func(Value, Value, Value, Value) Value)(unsafe.Pointer(&fn.UnsafePtr)))(a, b, c, d)
	}
	return Apply(Apply(Apply(Apply(fn, a), b), c), d)
}

func UncurriedApp5(fn Value, a, b, c, d, e Value) Value {
	if fn.Type == TypeFunc5 {
		return (*(*func(Value, Value, Value, Value, Value) Value)(unsafe.Pointer(&fn.UnsafePtr)))(a, b, c, d, e)
	}
	return Apply(Apply(Apply(Apply(Apply(fn, a), b), c), d), e)
}

func UncurriedApp6(fn Value, a, b, c, d, e, f Value) Value {
	return Apply6(fn, a, b, c, d, e, f)
}

func UncurriedApp7(fn Value, a, b, c, d, e, f, g Value) Value {
	return Apply7(fn, a, b, c, d, e, f, g)
}

func UncurriedApp8(fn Value, a, b, c, d, e, f, g, h Value) Value {
	return Apply8(fn, a, b, c, d, e, f, g, h)
}

func UncurriedApp9(fn Value, a, b, c, d, e, f, g, h, i Value) Value {
	return Apply9(fn, a, b, c, d, e, f, g, h, i)
}

func UncurriedApp10(fn Value, a, b, c, d, e, f, g, h, i, j Value) Value {
	return Apply10(fn, a, b, c, d, e, f, g, h, i, j)
}

func UncurriedApp(fn Value, args ...Value) Value {
	res := fn
	for _, arg := range args {
		res = Apply(res, arg)
	}
	return res
}

func Unbox[T any](v any) T {
	if res, ok := v.(T); ok {
		return res
	}
	
	val, isValue := v.(Value)
	if !isValue {
		var t any = *new(T)
		// If T is any, we can just return v
		switch t.(type) {
		case any: return any(v).(T)
		}
		panic("Cannot unbox non-Value into T")
	}
	
	var t any = *new(T)
	switch t.(type) {
	case int64: return any(val.IntVal).(T)
	case int: return any(int(val.IntVal)).(T)
	case string: return any(*(*string)(val.UnsafePtr)).(T)
	case float64: return any(math.Float64frombits(uint64(val.IntVal))).(T)
	case bool: return any(val.IntVal == 1).(T)
	case Value: return any(val).(T)
	case map[string]any:
		m := RecordToMap(val)
		res := make(map[string]any, len(m))
		for k, v2 := range m { res[k] = v2 }
		return any(res).(T)
	case []any:
		arr := *(*[]Value)(val.UnsafePtr)
		res := make([]any, len(arr))
		for i, v2 := range arr { res[i] = v2 }
		return any(res).(T)
	case func(any) any:
		res := func(arg any) any { return Apply(val, Box(arg)) }
		return any(res).(T)
	default:
		if val.Type == TypeAny {
			return (*(*any)(val.UnsafePtr)).(T)
		}
		return *(*T)(val.UnsafePtr)
	}
}

func Box[T any](val T) Value {
	if v, ok := any(val).(Value); ok {
		return v
	}
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
	case func(any) func(any) func(any) any:
		return Func(func(arg Value) Value { return Box(v(arg)) })
	case func(any) func(any) any:
		return Func(func(arg Value) Value { return Box(v(arg)) })
	case func(any) any:
		return Func(func(arg Value) Value { return Box(v(arg)) })
	case []any:
		arr := make([]Value, len(v))
		for i, val := range v { arr[i] = Box(val) }
		return Array(arr)
	case []Value:
		return Array(v)
	case map[string]any:
		m := make(map[string]Value)
		for k, val := range v { m[k] = Box(val) }
		return Record(m)
	case map[string]Value:
		return Record(v)
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

func Apply6(fn Value, arg1 Value, arg2 Value, arg3 Value, arg4 Value, arg5 Value, arg6 Value) Value {
	return Apply(Apply(Apply(Apply(Apply(Apply(fn, arg1), arg2), arg3), arg4), arg5), arg6)
}

func Apply7(fn Value, arg1 Value, arg2 Value, arg3 Value, arg4 Value, arg5 Value, arg6 Value, arg7 Value) Value {
	return Apply(Apply(Apply(Apply(Apply(Apply(Apply(fn, arg1), arg2), arg3), arg4), arg5), arg6), arg7)
}

func Apply8(fn Value, arg1 Value, arg2 Value, arg3 Value, arg4 Value, arg5 Value, arg6 Value, arg7 Value, arg8 Value) Value {
	return Apply(Apply(Apply(Apply(Apply(Apply(Apply(Apply(fn, arg1), arg2), arg3), arg4), arg5), arg6), arg7), arg8)
}

func Apply9(fn Value, arg1 Value, arg2 Value, arg3 Value, arg4 Value, arg5 Value, arg6 Value, arg7 Value, arg8 Value, arg9 Value) Value {
	return Apply(Apply(Apply(Apply(Apply(Apply(Apply(Apply(Apply(fn, arg1), arg2), arg3), arg4), arg5), arg6), arg7), arg8), arg9)
}

func Apply10(fn Value, arg1 Value, arg2 Value, arg3 Value, arg4 Value, arg5 Value, arg6 Value, arg7 Value, arg8 Value, arg9 Value, arg10 Value) Value {
	return Apply(Apply(Apply(Apply(Apply(Apply(Apply(Apply(Apply(Apply(fn, arg1), arg2), arg3), arg4), arg5), arg6), arg7), arg8), arg9), arg10)
}



"""
