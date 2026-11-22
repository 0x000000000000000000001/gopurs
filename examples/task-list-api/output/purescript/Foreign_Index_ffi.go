package purescript

import (
	"gopurs/output/gopurs_runtime"
	"reflect"
	"strconv"
)

func nativeArrayIndex(key string, length int) (int, bool) {
	index, err := strconv.ParseInt(key, 10, 64)
	valid := err == nil && strconv.FormatInt(index, 10) == key && index >= 0 && index < int64(length)
	return int(index), valid
}

func Foreign_Index_UnsafeReadPropImpl(f gopurs_runtime.Value, s gopurs_runtime.Value, key gopurs_runtime.Value, value gopurs_runtime.Value) gopurs_runtime.Value {
	if value.Type == gopurs_runtime.TypeAny && value.UnsafePtr == nil {
		return f
	}

	var propVal gopurs_runtime.Value
	found := false

	var kStr string
	isStr := false
	if key.Type == gopurs_runtime.TypeString {
		kStr = gopurs_runtime.StrValue(key)
		isStr = true
	} else if key.Type == gopurs_runtime.TypeInt {
		kStr = strconv.FormatInt(key.IntVal, 10)
		isStr = true
	}

	if value.Type == gopurs_runtime.TypeRecord && isStr {
		m := gopurs_runtime.RecordToMap(value)
		propVal, found = m[kStr]
	} else if value.Type == gopurs_runtime.TypeArray {
		arr := *(*[]gopurs_runtime.Value)(value.UnsafePtr)
		var idx int
		if key.Type == gopurs_runtime.TypeInt {
			idx = int(key.IntVal)
		} else if isStr {
			idx, _ = strconv.Atoi(kStr)
		}
		if idx >= 0 && idx < len(arr) {
			propVal = arr[idx]
			found = true
		}
	} else if value.Type == gopurs_runtime.TypeAny && value.UnsafePtr != nil && isStr {
		native := *(*any)(value.UnsafePtr)
		if m, ok := native.(map[string]interface{}); ok {
			if v, ok := m[kStr]; ok {
				propVal = gopurs_runtime.Any(v)
				found = true
			}
		} else if m, ok := native.(map[string]gopurs_runtime.Value); ok {
			if v, ok := m[kStr]; ok {
				propVal = v
				found = true
			}
		} else if array := reflect.ValueOf(native); array.Kind() == reflect.Slice {
			// String keys must spell the index exactly: "01", "+1" and
			// non-numeric properties are not array indices in JavaScript.
			if index, valid := nativeArrayIndex(kStr, array.Len()); valid {
				// Preserve native child identity and explicit legacy Values.
				// Reading one property must not convert its descendants.
				propVal = gopurs_runtime.Any(array.Index(index).Interface())
				found = true
			}
		}
	}

	if !found {
		propVal = gopurs_runtime.Value{Type: 0}
	}

	return gopurs_runtime.Apply(s, propVal)
}

func unsafeHasOwnPropertyImpl(prop gopurs_runtime.Value, value gopurs_runtime.Value) gopurs_runtime.Value {
	var kStr string
	isStr := false
	if prop.Type == gopurs_runtime.TypeString {
		kStr = gopurs_runtime.StrValue(prop)
		isStr = true
	} else if prop.Type == gopurs_runtime.TypeInt {
		kStr = strconv.FormatInt(prop.IntVal, 10)
		isStr = true
	}

	if value.Type == gopurs_runtime.TypeRecord && isStr {
		m := gopurs_runtime.RecordToMap(value)
		_, ok := m[kStr]
		return gopurs_runtime.Bool(ok)
	} else if value.Type == gopurs_runtime.TypeAny && value.UnsafePtr != nil && isStr {
		native := *(*any)(value.UnsafePtr)
		if m, ok := native.(map[string]interface{}); ok {
			_, ok := m[kStr]
			return gopurs_runtime.Bool(ok)
		} else if m, ok := native.(map[string]gopurs_runtime.Value); ok {
			_, ok := m[kStr]
			return gopurs_runtime.Bool(ok)
		} else if array := reflect.ValueOf(native); array.Kind() == reflect.Slice {
			_, found := nativeArrayIndex(kStr, array.Len())
			return gopurs_runtime.Bool(found)
		}
	}
	return gopurs_runtime.Bool(false)
}

func Foreign_Index_UnsafeHasOwnProperty(prop gopurs_runtime.Value, value gopurs_runtime.Value) gopurs_runtime.Value {
	return unsafeHasOwnPropertyImpl(prop, value)
}

func Foreign_Index_UnsafeHasProperty(prop gopurs_runtime.Value, value gopurs_runtime.Value) gopurs_runtime.Value {
	return unsafeHasOwnPropertyImpl(prop, value)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Foreign_Index_UnsafeHasOwnProperty = // TAST: (ForAll [k] (ADT ["Data","Function","Uncurried","Fn2"] [(TypeVar k), (ADT ["Foreign","Foreign"] []), Boolean]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Foreign_Index_UnsafeHasOwnProperty(go_arg0, go_arg1)
	return go_res
})
var _Gopurs_Foreign_Index_UnsafeHasProperty = // TAST: (ForAll [k] (ADT ["Data","Function","Uncurried","Fn2"] [(TypeVar k), (ADT ["Foreign","Foreign"] []), Boolean]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Foreign_Index_UnsafeHasProperty(go_arg0, go_arg1)
	return go_res
})
var _Gopurs_Foreign_Index_UnsafeReadPropImpl = // TAST: (ForAll [r, k] (ADT ["Data","Function","Uncurried","Fn4"] [(TypeVar r), (Func [(ADT ["Foreign","Foreign"] [])] (TypeVar r)), (TypeVar k), (ADT ["Foreign","Foreign"] []), (TypeVar r)]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := arg3
	go_res := Foreign_Index_UnsafeReadPropImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return go_res
})
