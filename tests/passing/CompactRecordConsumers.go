import (
	"fmt"
	"gopurs/output/gopurs_runtime"
)

// Value variables preserve the compact representation through FFI wrapping.
// A typed closed-record return would be converted back to a generic record.
var CompactEntry = gopurs_runtime.Func(func(reverse gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		if reverse.BoolVal() {
			return gopurs_runtime.RecordDict2("label", "count", gopurs_runtime.Str("alpha"), gopurs_runtime.Int(5))
		}
		return gopurs_runtime.RecordDict2("count", "label", gopurs_runtime.Int(5), gopurs_runtime.Str("alpha"))
	})
})

var CompactPayload = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	child := gopurs_runtime.RecordDict2("label", "count", gopurs_runtime.Str("nested"), gopurs_runtime.Int(2))
	values := gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Int(-3), gopurs_runtime.Int(0), gopurs_runtime.Int(7)})
	return gopurs_runtime.RecordDict2("values", "child", values, child)
})

var CompactScalars = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.RecordDict2("number", "flag", gopurs_runtime.Float(-1.25), gopurs_runtime.Bool(true))
})

func CompactKeys(record gopurs_runtime.Value) string {
	if record.Type != gopurs_runtime.TypeRecord2 {
		return "not-compact"
	}
	fields := (*gopurs_runtime.RecordData2)(record.UnsafePtr)
	return fields.K0 + "|" + fields.K1
}

func DescribeEntryMap(record map[string]any) string {
	if len(record) != 2 {
		panic("expected exactly count and label")
	}
	count := gopurs_runtime.Unbox[int64](record["count"].(gopurs_runtime.Value))
	label := gopurs_runtime.Unbox[string](record["label"].(gopurs_runtime.Value))
	return fmt.Sprintf("%d:%s", count, label)
}

func DescribePayloadMap(record map[string]any) string {
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

func DescribeScalarsMap(record map[string]any) string {
	if len(record) != 2 {
		panic("expected exactly flag and number")
	}
	flag := gopurs_runtime.Unbox[bool](record["flag"].(gopurs_runtime.Value))
	number := gopurs_runtime.Unbox[float64](record["number"].(gopurs_runtime.Value))
	return fmt.Sprintf("%t:%g", flag, number)
}
