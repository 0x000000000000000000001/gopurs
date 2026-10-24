package purescript

import (
	"fmt"
	"gopurs/output/gopurs_runtime"
)

// Value variables preserve the compact representation through FFI wrapping.
// A typed closed-record return would be converted back to a generic record.
var Main_CompactEntry = gopurs_runtime.Func(func(reverse gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		if reverse.BoolVal() {
			return gopurs_runtime.RecordDict2("label", "count", gopurs_runtime.Str("alpha"), gopurs_runtime.Int(5))
		}
		return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(5), gopurs_runtime.Str("alpha"))
	})
})

var Main_CompactPayload = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	child := gopurs_runtime.RecordDict2("label", "count", gopurs_runtime.Str("nested"), gopurs_runtime.Int(2))
	values := gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Int(-3), gopurs_runtime.Int(0), gopurs_runtime.Int(7)})
	return gopurs_runtime.RecordDict2("values", "child", values, child)
})

var Main_CompactScalars = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.RecordDict2("number", "flag", gopurs_runtime.Float(-1.25), gopurs_runtime.Bool(true))
})

func Main_CompactKeys(record gopurs_runtime.Value) string {
	if record.Type != gopurs_runtime.TypeRecord2 {
		return "not-compact"
	}
	fields := (*gopurs_runtime.RecordData2)(record.UnsafePtr)
	return fields.K0 + "|" + fields.K1
}

func Main_DescribeEntryMap(record map[string]any) string {
	if len(record) != 2 {
		panic("expected exactly count and label")
	}
	count := gopurs_runtime.Unbox[int64](record["count"].(gopurs_runtime.Value))
	label := gopurs_runtime.Unbox[string](record["label"].(gopurs_runtime.Value))
	return fmt.Sprintf("%d:%s", count, label)
}

func Main_DescribePayloadMap(record map[string]any) string {
	if len(record) != 2 {
		panic("expected exactly child and values")
	}
	child := record["child"].(gopurs_runtime.Value)
	count := gopurs_runtime.Unbox[int64](gopurs_runtime.RecordGet(child, "count"))
	label := gopurs_runtime.Unbox[string](gopurs_runtime.RecordGet(child, "label"))
	values := gopurs_runtime.Unbox[[]any](record["values"].(gopurs_runtime.Value))
	ints := make([]int64, len(values))
	for i, value := range values {
		ints[i] = gopurs_runtime.Unbox[int64](value.(gopurs_runtime.Value))
	}
	return fmt.Sprintf("%d:%s:%v", count, label, ints)
}

func Main_DescribeScalarsMap(record map[string]any) string {
	if len(record) != 2 {
		panic("expected exactly flag and number")
	}
	flag := gopurs_runtime.Unbox[bool](record["flag"].(gopurs_runtime.Value))
	number := gopurs_runtime.Unbox[float64](record["number"].(gopurs_runtime.Value))
	return fmt.Sprintf("%t:%g", flag, number)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Main_CompactEntry = // TAST: (Func [Boolean] (ADT ["Effect","Effect"] [Any]))
gopurs_runtime.Box(Main_CompactEntry)
var _Gopurs_Main_CompactKeys = // TAST: (ForAll [a] (Func [(TypeVar a)] String))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Main_CompactKeys(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Main_CompactPayload = // TAST: (ADT ["Effect","Effect"] [Any])
gopurs_runtime.Box(Main_CompactPayload)
var _Gopurs_Main_CompactScalars = // TAST: (ADT ["Effect","Effect"] [Any])
gopurs_runtime.Box(Main_CompactScalars)
var _Gopurs_Main_DescribeEntryMap = // TAST: (Func [Any] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.UnboxObject(arg0)
	go_res := Main_DescribeEntryMap(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Main_DescribePayloadMap = // TAST: (Func [Any] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.UnboxObject(arg0)
	go_res := Main_DescribePayloadMap(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Main_DescribeScalarsMap = // TAST: (Func [Any] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.UnboxObject(arg0)
	go_res := Main_DescribeScalarsMap(go_arg0)
	return gopurs_runtime.Box(go_res)
})
