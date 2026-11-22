package purescript

import (
	"gopurs/output/gopurs_runtime"
	"reflect"
)

var undefinedForJSON = struct{}{}

// The safe readers check the native kind before these representation changes.
// Reflection preserves support for named primitive and slice types accepted
// by TagOf and IsArray. Array elements remain native Foreign payloads.
func Foreign_UnsafeReadString(value any) string {
	return reflect.ValueOf(value).String()
}

func Foreign_UnsafeReadBoolean(value any) bool {
	return reflect.ValueOf(value).Bool()
}

func Foreign_UnsafeReadArray(value gopurs_runtime.Value) []gopurs_runtime.Value {
	if value.Type == gopurs_runtime.TypeArray {
		return gopurs_runtime.Unbox[[]gopurs_runtime.Value](value)
	}
	// Reading an array inspects only its outer shape. Do not traverse or
	// convert its children: they can include shared/cyclic native containers,
	// undefined, or values retained by the explicit legacy runtime ABI.
	array := reflect.ValueOf(value.PtrVal())
	result := make([]gopurs_runtime.Value, array.Len())
	for i := range result {
		result[i] = gopurs_runtime.Any(array.Index(i).Interface())
	}
	return result
}

func Foreign_UnsafeReadNumber(value any) float64 {
	number := reflect.ValueOf(value)
	switch number.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(number.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(number.Uint())
	default:
		return number.Float()
	}
}

func Foreign_TypeOf(v gopurs_runtime.Value) gopurs_runtime.Value {
	switch v.Type {
	case gopurs_runtime.TypeBool:
		return gopurs_runtime.Str("boolean")
	case gopurs_runtime.TypeInt, gopurs_runtime.TypeFloat:
		return gopurs_runtime.Str("number")
	case gopurs_runtime.TypeString:
		return gopurs_runtime.Str("string")
	case gopurs_runtime.TypeFunc, gopurs_runtime.TypeFunc2, gopurs_runtime.TypeFunc3,
		gopurs_runtime.TypeFunc4, gopurs_runtime.TypeFunc5, gopurs_runtime.TypeFunc6,
		gopurs_runtime.TypeFunc7, gopurs_runtime.TypeFunc8, gopurs_runtime.TypeFunc9,
		gopurs_runtime.TypeFunc10, gopurs_runtime.TypeFunc11, gopurs_runtime.TypeFunctionData:
		return gopurs_runtime.Str("function")
	case 0:
		return gopurs_runtime.Str("undefined")
	default:
		if v.Type == gopurs_runtime.TypeAny && v.UnsafePtr != nil {
			val := *(*any)(v.UnsafePtr)
			if val != nil {
				rt := reflect.TypeOf(val)
				if rt != nil {
					if rt.Name() == "Date" {
						return gopurs_runtime.Str("Date")
					}

					switch rt.Kind() {
					case reflect.String:
						return gopurs_runtime.Str("string")
					case reflect.Bool:
						return gopurs_runtime.Str("boolean")
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
						reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
						reflect.Float32, reflect.Float64:
						return gopurs_runtime.Str("number")
					}
				}
			}
		}
		return gopurs_runtime.Str("object")
	}
}

func Foreign_TagOf(v gopurs_runtime.Value) gopurs_runtime.Value {
	switch v.Type {
	case gopurs_runtime.TypeBool:
		return gopurs_runtime.Str("Boolean")
	case gopurs_runtime.TypeInt, gopurs_runtime.TypeFloat:
		return gopurs_runtime.Str("Number")
	case gopurs_runtime.TypeString:
		return gopurs_runtime.Str("String")
	case gopurs_runtime.TypeArray:
		return gopurs_runtime.Str("Array")
	case gopurs_runtime.TypeFunc, gopurs_runtime.TypeFunc2, gopurs_runtime.TypeFunc3,
		gopurs_runtime.TypeFunc4, gopurs_runtime.TypeFunc5, gopurs_runtime.TypeFunc6,
		gopurs_runtime.TypeFunc7, gopurs_runtime.TypeFunc8, gopurs_runtime.TypeFunc9,
		gopurs_runtime.TypeFunc10, gopurs_runtime.TypeFunc11, gopurs_runtime.TypeFunctionData:
		return gopurs_runtime.Str("Function")
	case 0:
		return gopurs_runtime.Str("Undefined")
	default:
		if v.Type == gopurs_runtime.TypeAny {
			if v.UnsafePtr == nil {
				return gopurs_runtime.Str("Null")
			}
			val := *(*any)(v.UnsafePtr)
			if val == nil {
				return gopurs_runtime.Str("Null")
			}
			rt := reflect.TypeOf(val)
			if rt != nil {
				if rt.Kind() == reflect.Slice {
					return gopurs_runtime.Str("Array")
				}
				if rt.Kind() == reflect.String {
					return gopurs_runtime.Str("String")
				}
				if rt.Kind() == reflect.Bool {
					return gopurs_runtime.Str("Boolean")
				}
				if rt.Name() == "Date" {
					return gopurs_runtime.Str("Date")
				}

				switch rt.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
					reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
					reflect.Float32, reflect.Float64:
					return gopurs_runtime.Str("Number")
				}
			}
		}
		return gopurs_runtime.Str("Object")
	}
}

func Foreign_IsNull(v gopurs_runtime.Value) gopurs_runtime.Value {
	if v.Type == gopurs_runtime.TypeAny {
		if v.UnsafePtr == nil {
			return gopurs_runtime.Bool(true)
		}
		val := *(*any)(v.UnsafePtr)
		if val == nil {
			return gopurs_runtime.Bool(true)
		}
	}
	return gopurs_runtime.Bool(false)
}

func Foreign_IsUndefined(v gopurs_runtime.Value) gopurs_runtime.Value {
	if v.Type == 0 {
		return gopurs_runtime.Bool(true)
	}
	return gopurs_runtime.Bool(false)
}

func Foreign_IsArray(v gopurs_runtime.Value) gopurs_runtime.Value {
	if v.Type == gopurs_runtime.TypeArray {
		return gopurs_runtime.Bool(true)
	}
	if v.Type == gopurs_runtime.TypeAny && v.UnsafePtr != nil {
		val := *(*any)(v.UnsafePtr)
		if val != nil && reflect.TypeOf(val).Kind() == reflect.Slice {
			return gopurs_runtime.Bool(true)
		}
	}
	return gopurs_runtime.Bool(false)
}

func Foreign_UnboxForJSON(v interface{}) interface{} {
	return foreignDeepUnbox(v)
}

func foreignDeepUnbox(v interface{}) interface{} {
	if gopurs_runtime.IsForeignUndefined(v) {
		return undefinedForJSON
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
					res[i] = foreignJSONElement(x)
				}
				return res
			}
			return []interface{}{}
		case gopurs_runtime.TypeRecord, gopurs_runtime.TypeRecord0, gopurs_runtime.TypeRecord1, gopurs_runtime.TypeRecord2, gopurs_runtime.TypeRecord3, gopurs_runtime.TypeRecord4, gopurs_runtime.TypeRecord5, gopurs_runtime.TypeRecordData:
			rec := gopurs_runtime.RecordToMap(val)
			res := make(map[string]interface{})
			for k, x := range rec {
				unboxed := foreignDeepUnbox(x)
				if unboxed != undefinedForJSON {
					res[k] = unboxed
				}
			}
			return res
		case gopurs_runtime.TypeAny:
			if val.UnsafePtr == nil {
				return nil
			}
			if *(*any)(val.UnsafePtr) == nil {
				return nil
			}
			return foreignDeepUnbox(*(*any)(val.UnsafePtr))
		case 0:
			return undefinedForJSON
		default:
			return nil
		}
	}
	if valSlice, ok := v.([]gopurs_runtime.Value); ok {
		res := make([]interface{}, len(valSlice))
		for i, x := range valSlice {
			res[i] = foreignJSONElement(x)
		}
		return res
	}
	if mapGopurs, ok := v.(map[string]gopurs_runtime.Value); ok {
		res := make(map[string]interface{})
		for k, x := range mapGopurs {
			unboxed := foreignDeepUnbox(x)
			if unboxed != undefinedForJSON {
				res[k] = unboxed
			}
		}
		return res
	}
	if mapRaw, ok := v.(map[string]interface{}); ok {
		res := make(map[string]interface{})
		for k, x := range mapRaw {
			unboxed := foreignDeepUnbox(x)
			if unboxed != undefinedForJSON {
				res[k] = unboxed
			}
		}
		return res
	}
	if arrRaw, ok := v.([]interface{}); ok {
		res := make([]interface{}, len(arrRaw))
		for i, x := range arrRaw {
			res[i] = foreignJSONElement(x)
		}
		return res
	}
	return v
}

// JSON.stringify retains array positions by encoding undefined as null.
// Object fields are omitted by foreignDeepUnbox instead.
func foreignJSONElement(v interface{}) interface{} {
	value := foreignDeepUnbox(v)
	if value == undefinedForJSON {
		return nil
	}
	return value
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Foreign_IsArray = // TAST: (Func [(ADT ["Foreign","Foreign"] [])] Boolean)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Foreign_IsArray(go_arg0)
	return go_res
})
var _Gopurs_Foreign_IsNull = // TAST: (Func [(ADT ["Foreign","Foreign"] [])] Boolean)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Foreign_IsNull(go_arg0)
	return go_res
})
var _Gopurs_Foreign_IsUndefined = // TAST: (Func [(ADT ["Foreign","Foreign"] [])] Boolean)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Foreign_IsUndefined(go_arg0)
	return go_res
})
var _Gopurs_Foreign_TagOf = // TAST: (Func [(ADT ["Foreign","Foreign"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Foreign_TagOf(go_arg0)
	return go_res
})
var _Gopurs_Foreign_TypeOf = // TAST: (Func [(ADT ["Foreign","Foreign"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Foreign_TypeOf(go_arg0)
	return go_res
})
var _Gopurs_Foreign_UnboxForJSON = // TAST: (Func [(ADT ["Foreign","Foreign"] [])] (ADT ["Foreign","Foreign"] []))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.ForeignValue(arg0)
	go_res := Foreign_UnboxForJSON(go_arg0)
	return gopurs_runtime.Any(go_res)
})
var _Gopurs_Foreign_UnsafeReadArray = // TAST: (Func [(ADT ["Foreign","Foreign"] [])] (Array (ADT ["Foreign","Foreign"] [])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Foreign_UnsafeReadArray(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Foreign_UnsafeReadBoolean = // TAST: (Func [(ADT ["Foreign","Foreign"] [])] Boolean)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.ForeignValue(arg0)
	go_res := Foreign_UnsafeReadBoolean(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Foreign_UnsafeReadNumber = // TAST: (Func [(ADT ["Foreign","Foreign"] [])] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.ForeignValue(arg0)
	go_res := Foreign_UnsafeReadNumber(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Foreign_UnsafeReadString = // TAST: (Func [(ADT ["Foreign","Foreign"] [])] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.ForeignValue(arg0)
	go_res := Foreign_UnsafeReadString(go_arg0)
	return gopurs_runtime.Box(go_res)
})
