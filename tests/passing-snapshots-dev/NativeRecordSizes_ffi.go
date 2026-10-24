package purescript

import (
	"fmt"
	"gopurs/output/gopurs_runtime"
	"sort"
	"strings"
)

// An opaque PureScript field must remain a Go Value inside the native struct.
var Main_OpaqueValue = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Str("opaque")
})

func Main_NativeShape(record gopurs_runtime.Value) string {
	var keys []string
	shape := ""
	switch record.Type {
	case gopurs_runtime.TypeRecord0:
		shape = "compact0"
	case gopurs_runtime.TypeRecord1:
		r := (*gopurs_runtime.RecordData1)(record.UnsafePtr)
		shape, keys = "compact1", []string{r.K0}
	case gopurs_runtime.TypeRecord2:
		r := (*gopurs_runtime.RecordData2)(record.UnsafePtr)
		shape, keys = "compact2", []string{r.K0, r.K1}
	case gopurs_runtime.TypeRecord3:
		r := (*gopurs_runtime.RecordData3)(record.UnsafePtr)
		shape, keys = "compact3", []string{r.K0, r.K1, r.K2}
	case gopurs_runtime.TypeRecord4:
		r := (*gopurs_runtime.RecordData4)(record.UnsafePtr)
		shape, keys = "compact4", []string{r.K0, r.K1, r.K2, r.K3}
	case gopurs_runtime.TypeRecord5:
		r := (*gopurs_runtime.RecordData5)(record.UnsafePtr)
		shape, keys = "compact5", []string{r.K0, r.K1, r.K2, r.K3, r.K4}
	case gopurs_runtime.TypeRecordData:
		r := (*gopurs_runtime.RecordData)(record.UnsafePtr)
		shape, keys = fmt.Sprintf("generic%d", len(r.Keys)), r.Keys
	default:
		panic(fmt.Sprintf("unexpected record representation: %d", record.Type))
	}
	// Preserve physical order here; the map renderer below sorts independently.
	return shape + ":" + strings.Join(keys, "|")
}

func nativeSizesRender(value gopurs_runtime.Value) string {
	switch value.Type {
	case gopurs_runtime.TypeInt:
		return fmt.Sprintf("i:%d", gopurs_runtime.Unbox[int64](value))
	case gopurs_runtime.TypeString:
		return "s:" + gopurs_runtime.Unbox[string](value)
	case gopurs_runtime.TypeBool:
		return fmt.Sprintf("b:%t", gopurs_runtime.Unbox[bool](value))
	case gopurs_runtime.TypeFloat:
		return fmt.Sprintf("n:%g", gopurs_runtime.Unbox[float64](value))
	case gopurs_runtime.TypeArray:
		values := gopurs_runtime.Unbox[[]any](value)
		parts := make([]string, len(values))
		for i, item := range values {
			parts[i] = nativeSizesRender(item.(gopurs_runtime.Value))
		}
		return "[" + strings.Join(parts, ",") + "]"
	default:
		if value.Type == gopurs_runtime.TypeRecord ||
			(value.Type >= gopurs_runtime.TypeRecord0 && value.Type <= gopurs_runtime.TypeRecordData) {
			return nativeSizesRenderRecord(value)
		}
		panic(fmt.Sprintf("unexpected field representation: %d", value.Type))
	}
}

func nativeSizesRenderRecord(record gopurs_runtime.Value) string {
	fields := gopurs_runtime.RecordToMap(record)
	keys := make([]string, 0, len(fields))
	for key := range fields {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	parts := make([]string, len(keys))
	for i, key := range keys {
		parts[i] = key + "=" + nativeSizesRender(gopurs_runtime.RecordGet(record, key))
	}
	return "{" + strings.Join(parts, ",") + "}"
}

// Exercise the generated record -> map[string]any bridge. All fields must
// still be boxed Values; no conversion of their tags is performed here.
func Main_NativeMap(record map[string]any) string {
	fields := make(map[string]gopurs_runtime.Value, len(record))
	for key, value := range record {
		fields[key] = value.(gopurs_runtime.Value)
	}
	return nativeSizesRenderRecord(gopurs_runtime.Record(fields))
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Main_NativeMap = // TAST: (ForAll [a] (Func [(TypeVar a)] String))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.UnboxObject(arg0)
	go_res := Main_NativeMap(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Main_NativeShape = // TAST: (ForAll [a] (Func [(TypeVar a)] String))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Main_NativeShape(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Main_OpaqueValue = // TAST: (ADT ["Effect","Effect"] [(ADT ["Main","Opaque"] [])])
gopurs_runtime.Box(Main_OpaqueValue)
