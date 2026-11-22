package purescript

import (
	"bytes"
	"encoding/json"
	"gopurs/output/gopurs_runtime"
	"sort"
)

func argonautDeepUnbox(v interface{}) interface{} {
	if object, ok := v.(gopurs_runtime.JSONObject); ok {
		result := make(map[string]any, object.JSONLength())
		for i := 0; i < object.JSONLength(); i++ {
			key, value := object.JSONEntry(i)
			result[key] = argonautDeepUnbox(value)
		}
		return result
	}
	if val, ok := v.(gopurs_runtime.Value); ok {
		switch val.Type {
		case gopurs_runtime.TypeInt:
			return val.IntVal
		case gopurs_runtime.TypeFloat:
			return val.FloatVal()
		case gopurs_runtime.TypeString:
			if val.UnsafePtr != nil {
				return gopurs_runtime.StrValue(val)
			}
			return ""
		case gopurs_runtime.TypeBool:
			return val.IntVal != 0
		case gopurs_runtime.TypeArray:
			if val.UnsafePtr != nil {
				arr := *(*[]gopurs_runtime.Value)(val.UnsafePtr)
				res := make([]interface{}, len(arr))
				for i, x := range arr {
					res[i] = argonautDeepUnbox(x)
				}
				return res
			}
			return []interface{}{}
		case gopurs_runtime.TypeRecord, gopurs_runtime.TypeRecord0, gopurs_runtime.TypeRecord1, gopurs_runtime.TypeRecord2, gopurs_runtime.TypeRecord3, gopurs_runtime.TypeRecord4, gopurs_runtime.TypeRecord5:
			rec := gopurs_runtime.RecordToMap(val)
			res := make(map[string]interface{})
			for k, x := range rec {
				res[k] = argonautDeepUnbox(x)
			}
			return res
		case gopurs_runtime.TypeAny:
			if val.UnsafePtr == nil {
				return nil
			}
			if *(*any)(val.UnsafePtr) == nil {
				return nil
			}
			return argonautDeepUnbox(*(*any)(val.UnsafePtr))
		default:
			return nil
		}
	}
	if valSlice, ok := v.([]gopurs_runtime.Value); ok {
		res := make([]interface{}, len(valSlice))
		for i, x := range valSlice {
			res[i] = argonautDeepUnbox(x)
		}
		return res
	}
	if mapRaw, ok := v.(map[string]interface{}); ok {
		res := make(map[string]interface{})
		for k, x := range mapRaw {
			res[k] = argonautDeepUnbox(x)
		}
		return res
	}
	if arrRaw, ok := v.([]interface{}); ok {
		res := make([]interface{}, len(arrRaw))
		for i, x := range arrRaw {
			res[i] = argonautDeepUnbox(x)
		}
		return res
	}
	return v
}

func Data_Argonaut_Core_FromBoolean(b any) any {
	return b
}

func Data_Argonaut_Core_FromNumber(n any) any {
	return n
}

func Data_Argonaut_Core_FromString(s any) any {
	return s
}

func Data_Argonaut_Core_FromArray(a any) any {
	return a
}

func Data_Argonaut_Core_FromObject(o any) any {
	return o
}

var Data_Argonaut_Core_JsonNull any = nil

func Data_Argonaut_Core_Stringify(j any) string {
	b, _ := json.Marshal(argonautDeepUnbox(j))
	return string(b)
}

func Data_Argonaut_Core_StringifyWithIndent(i int, j any) string {
	spaces := ""
	if i > 0 {
		if i > 10 {
			i = 10
		}
		for k := 0; k < i; k++ {
			spaces += " "
		}
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", spaces)
	enc.SetEscapeHTML(false)
	enc.Encode(argonautDeepUnbox(j))
	return buf.String()
}

func isArray(a any) bool {
	_, ok := argonautDeepUnbox(a).([]any)
	return ok
}

func Data_Argonaut_Core__Compare(EQ any, GT any, LT any, a any, b any) any {
	return argonautCompare(EQ, GT, LT, a, b)
}

func argonautCompare(EQ any, GT any, LT any, a any, b any) any {
	a = argonautDeepUnbox(a)
	b = argonautDeepUnbox(b)
	if a == nil {
		if b == nil {
			return EQ
		}
		return LT
	}
	switch va := a.(type) {
	case bool:
		if vb, ok := b.(bool); ok {
			if va == vb {
				return EQ
			} else if !va {
				return LT
			}
			return GT
		}
		if b == nil {
			return GT
		}
		return LT
	case float64:
		if vb, ok := b.(float64); ok {
			if va == vb {
				return EQ
			} else if va < vb {
				return LT
			}
			return GT
		}
		if b == nil || isBool(b) {
			return GT
		}
		return LT
	case string:
		if vb, ok := b.(string); ok {
			if va == vb {
				return EQ
			} else if va < vb {
				return LT
			}
			return GT
		}
		if b == nil || isBool(b) || isFloat64(b) {
			return GT
		}
		return LT
	case []any:
		if vb, ok := b.([]any); ok {
			minLen := len(va)
			if len(vb) < minLen {
				minLen = len(vb)
			}
			for i := 0; i < minLen; i++ {
				ca := argonautCompare(EQ, GT, LT, va[i], vb[i])
				if ca != EQ {
					return ca
				}
			}
			if len(va) == len(vb) {
				return EQ
			} else if len(va) < len(vb) {
				return LT
			}
			return GT
		}
		if b == nil || isBool(b) || isFloat64(b) || isString(b) {
			return GT
		}
		return LT
	case map[string]any:
		if vb, ok := b.(map[string]any); ok {
			if len(va) < len(vb) {
				return LT
			} else if len(va) > len(vb) {
				return GT
			}
			akeys := make([]string, 0, len(va))
			bkeys := make([]string, 0, len(vb))
			for k := range va {
				akeys = append(akeys, k)
			}
			for k := range vb {
				bkeys = append(bkeys, k)
			}
			keysMap := make(map[string]bool)
			for _, k := range akeys {
				keysMap[k] = true
			}
			for _, k := range bkeys {
				keysMap[k] = true
			}
			keys := make([]string, 0, len(keysMap))
			for k := range keysMap {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				valA, okA := va[k]
				valB, okB := vb[k]
				if !okA {
					return LT
				} else if !okB {
					return GT
				}
				ck := argonautCompare(EQ, GT, LT, valA, valB)
				if ck != EQ {
					return ck
				}
			}
			return EQ
		}
		return GT
	default:
		panic("unknown JSON type in argonautCompare")
	}
}

func isBool(v any) bool {
	_, ok := argonautDeepUnbox(v).(bool)
	return ok
}

func isFloat64(v any) bool {
	_, ok := argonautDeepUnbox(v).(float64)
	return ok
}

func isString(v any) bool {
	_, ok := argonautDeepUnbox(v).(string)
	return ok
}

func Data_Argonaut_Core_CaseJsonImpl(onNull, onBool, onNum, onStr, onArr, onObj, j gopurs_runtime.Value) gopurs_runtime.Value {
	switch j.Type {
	case gopurs_runtime.TypeBool:
		return gopurs_runtime.Apply(onBool, j)
	case gopurs_runtime.TypeInt, gopurs_runtime.TypeFloat:
		return gopurs_runtime.Apply(onNum, j)
	case gopurs_runtime.TypeString:
		return gopurs_runtime.Apply(onStr, j)
	case gopurs_runtime.TypeArray:
		return gopurs_runtime.Apply(onArr, j)
	case 0:
		return gopurs_runtime.Apply(onNull, gopurs_runtime.Value{})
	case gopurs_runtime.TypeRecord, gopurs_runtime.TypeRecord0, gopurs_runtime.TypeRecord1, gopurs_runtime.TypeRecord2, gopurs_runtime.TypeRecord3, gopurs_runtime.TypeRecord4, gopurs_runtime.TypeRecord5:
		return gopurs_runtime.Apply(onObj, j)
	case gopurs_runtime.TypeAny:
		if j.UnsafePtr == nil {
			return gopurs_runtime.Apply(onNull, gopurs_runtime.Value{})
		}
		if *(*any)(j.UnsafePtr) == nil {
			return gopurs_runtime.Apply(onNull, gopurs_runtime.Value{})
		}
		return gopurs_runtime.Apply(onObj, j)
	default:
		return gopurs_runtime.Apply(onObj, j)
	}
}

// Single-type case analysis. These replace the generic caseJson path that
// builds six constant callbacks per call. Dispatch matches CaseJsonImpl.
func argoJsonIsNull(j gopurs_runtime.Value) bool {
	if j.Type == 0 {
		return true
	}
	if j.Type == gopurs_runtime.TypeAny {
		if j.UnsafePtr == nil {
			return true
		}
		return *(*any)(j.UnsafePtr) == nil
	}
	return false
}

func argoJsonIsScalarOrNull(j gopurs_runtime.Value) bool {
	switch j.Type {
	case gopurs_runtime.TypeBool, gopurs_runtime.TypeInt, gopurs_runtime.TypeFloat, gopurs_runtime.TypeString, gopurs_runtime.TypeArray:
		return true
	case 0:
		return true
	case gopurs_runtime.TypeAny:
		return argoJsonIsNull(j)
	default:
		// Records, record dictionaries and every other constructor fall through
		// to the object branch, exactly like CaseJsonImpl.
		return false
	}
}

func Data_Argonaut_Core__CaseJsonNull(d, f, j gopurs_runtime.Value) gopurs_runtime.Value {
	if argoJsonIsNull(j) {
		return gopurs_runtime.Apply(f, gopurs_runtime.Value{})
	}
	return d
}

func Data_Argonaut_Core__CaseJsonBoolean(d, f, j gopurs_runtime.Value) gopurs_runtime.Value {
	if j.Type == gopurs_runtime.TypeBool {
		return gopurs_runtime.Apply(f, j)
	}
	return d
}

func Data_Argonaut_Core__CaseJsonNumber(d, f, j gopurs_runtime.Value) gopurs_runtime.Value {
	if j.Type == gopurs_runtime.TypeInt || j.Type == gopurs_runtime.TypeFloat {
		return gopurs_runtime.Apply(f, j)
	}
	return d
}

func Data_Argonaut_Core__CaseJsonString(d, f, j gopurs_runtime.Value) gopurs_runtime.Value {
	if j.Type == gopurs_runtime.TypeString {
		return gopurs_runtime.Apply(f, j)
	}
	return d
}

func Data_Argonaut_Core__CaseJsonArray(d, f, j gopurs_runtime.Value) gopurs_runtime.Value {
	if j.Type == gopurs_runtime.TypeArray {
		return gopurs_runtime.Apply(f, j)
	}
	return d
}

func Data_Argonaut_Core__CaseJsonObject(d, f, j gopurs_runtime.Value) gopurs_runtime.Value {
	if argoJsonIsScalarOrNull(j) {
		return d
	}
	return gopurs_runtime.Apply(f, j)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_Argonaut_Core__CaseJsonArray = // TAST: (Func [(ADT ["Data","Argonaut","Core","Json"] []), (Func [(Array (ADT ["Data","Argonaut","Core","Json"] []))] (ADT ["Data","Argonaut","Core","Json"] [])), (ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Argonaut","Core","Json"] []))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Data_Argonaut_Core__CaseJsonArray(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Data_Argonaut_Core__CaseJsonBoolean = // TAST: (Func [(ADT ["Data","Argonaut","Core","Json"] []), (Func [Boolean] (ADT ["Data","Argonaut","Core","Json"] [])), (ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Argonaut","Core","Json"] []))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Data_Argonaut_Core__CaseJsonBoolean(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Data_Argonaut_Core__CaseJsonNull = // TAST: (Func [(ADT ["Data","Argonaut","Core","Json"] []), (Func [Unit] (ADT ["Data","Argonaut","Core","Json"] [])), (ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Argonaut","Core","Json"] []))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Data_Argonaut_Core__CaseJsonNull(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Data_Argonaut_Core__CaseJsonNumber = // TAST: (Func [(ADT ["Data","Argonaut","Core","Json"] []), (Func [Number] (ADT ["Data","Argonaut","Core","Json"] [])), (ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Argonaut","Core","Json"] []))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Data_Argonaut_Core__CaseJsonNumber(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Data_Argonaut_Core__CaseJsonObject = // TAST: (Func [(ADT ["Data","Argonaut","Core","Json"] []), (Func [(ADT ["Foreign","Object","Object"] [(ADT ["Data","Argonaut","Core","Json"] [])])] (ADT ["Data","Argonaut","Core","Json"] [])), (ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Argonaut","Core","Json"] []))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Data_Argonaut_Core__CaseJsonObject(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Data_Argonaut_Core__CaseJsonString = // TAST: (Func [(ADT ["Data","Argonaut","Core","Json"] []), (Func [String] (ADT ["Data","Argonaut","Core","Json"] [])), (ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Argonaut","Core","Json"] []))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Data_Argonaut_Core__CaseJsonString(go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Data_Argonaut_Core__Compare = // TAST: (ADT ["Data","Function","Uncurried","Fn5"] [(ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Argonaut","Core","Json"] []), (ADT ["Data","Argonaut","Core","Json"] []), (ADT ["Data","Ordering","Ordering"] [])])
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := Data_Argonaut_Core__Compare(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Argonaut_Core_CaseJsonImpl = // TAST: (Func [(Func [Unit] (ADT ["Data","Argonaut","Core","Json"] [])), (Func [Boolean] (ADT ["Data","Argonaut","Core","Json"] [])), (Func [Number] (ADT ["Data","Argonaut","Core","Json"] [])), (Func [String] (ADT ["Data","Argonaut","Core","Json"] [])), (Func [(Array (ADT ["Data","Argonaut","Core","Json"] []))] (ADT ["Data","Argonaut","Core","Json"] [])), (Func [(ADT ["Foreign","Object","Object"] [(ADT ["Data","Argonaut","Core","Json"] [])])] (ADT ["Data","Argonaut","Core","Json"] [])), (ADT ["Data","Argonaut","Core","Json"] [])] (ADT ["Data","Argonaut","Core","Json"] []))
gopurs_runtime.Func7(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value, arg6 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_arg5 := arg5
	go_arg6 := arg6
	go_res := Data_Argonaut_Core_CaseJsonImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5, go_arg6)
	return go_res
})
var _Gopurs_Data_Argonaut_Core_FromArray = // TAST: (Func [(Array (ADT ["Data","Argonaut","Core","Json"] []))] (ADT ["Data","Argonaut","Core","Json"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Argonaut_Core_FromArray(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Argonaut_Core_FromBoolean = // TAST: (Func [Boolean] (ADT ["Data","Argonaut","Core","Json"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Argonaut_Core_FromBoolean(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Argonaut_Core_FromNumber = // TAST: (Func [Number] (ADT ["Data","Argonaut","Core","Json"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Argonaut_Core_FromNumber(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Argonaut_Core_FromObject = // TAST: (Func [(ADT ["Foreign","Object","Object"] [(ADT ["Data","Argonaut","Core","Json"] [])])] (ADT ["Data","Argonaut","Core","Json"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Argonaut_Core_FromObject(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Argonaut_Core_FromString = // TAST: (Func [String] (ADT ["Data","Argonaut","Core","Json"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Argonaut_Core_FromString(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Argonaut_Core_JsonNull = // TAST: (ADT ["Data","Argonaut","Core","Json"] [])
gopurs_runtime.Box(Data_Argonaut_Core_JsonNull)
var _Gopurs_Data_Argonaut_Core_Stringify = // TAST: (Func [(ADT ["Data","Argonaut","Core","Json"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Argonaut_Core_Stringify(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_Argonaut_Core_StringifyWithIndent = // TAST: (Func [Int, (ADT ["Data","Argonaut","Core","Json"] [])] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := arg1
	go_res := Data_Argonaut_Core_StringifyWithIndent(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})

// --- Boxed-result FFI workers ---
func Data_Argonaut_Core__Compare_nativeWorker(arg0 any, arg1 any, arg2 any, arg3 any, arg4 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := arg4
	go_res := Data_Argonaut_Core__Compare(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
}

func Data_Argonaut_Core_FromArray_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Argonaut_Core_FromArray(go_arg0)
	return gopurs_runtime.Box(go_res)
}

func Data_Argonaut_Core_FromBoolean_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Argonaut_Core_FromBoolean(go_arg0)
	return gopurs_runtime.Box(go_res)
}

func Data_Argonaut_Core_FromNumber_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Argonaut_Core_FromNumber(go_arg0)
	return gopurs_runtime.Box(go_res)
}

func Data_Argonaut_Core_FromObject_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Argonaut_Core_FromObject(go_arg0)
	return gopurs_runtime.Box(go_res)
}

func Data_Argonaut_Core_FromString_nativeWorker(arg0 any) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_Argonaut_Core_FromString(go_arg0)
	return gopurs_runtime.Box(go_res)
}
