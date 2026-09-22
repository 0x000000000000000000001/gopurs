package gopurs_runtime

import (
	"fmt"
	"runtime"
	"testing"
	"unsafe"
)

func TestRecordSetCompactPromotionAndReplacement(t *testing.T) {
	keys := []string{"", "é雪", "a", "b", "c", "d", "e"}
	records := []Value{RecordDict0()}
	for i, key := range keys {
		orig := records[i]
		next := RecordSet(orig, key, Int(int64(i+1)))
		wantType := TypeRecordData
		if i < 5 {
			wantType = TypeRecord1 + i
		}
		if next.Type != wantType {
			t.Fatalf("insert %d: type %d, want %d", i, next.Type, wantType)
		}
		records = append(records, next)
		for j, field := range keys[:i+1] {
			replaced := RecordSet(next, field, Int(99))
			if replaced.Type != next.Type || len(RecordToMap(replaced)) != i+1 {
				t.Fatalf("replace %d/%d changed representation or field count", i, j)
			}
			if RecordGet(replaced, field).IntVal != 99 || RecordGet(next, field).IntVal != int64(j+1) {
				t.Fatalf("replace %d/%d failed or mutated source", i, j)
			}
		}
	}
	for size, record := range records {
		fields := RecordToMap(record)
		if len(fields) != size {
			t.Fatalf("retained record %d now has %d fields", size, len(fields))
		}
		for i, key := range keys[:size] {
			if fields[key].IntVal != int64(i+1) {
				t.Fatalf("retained record %d key %q mutated", size, key)
			}
		}
	}
}

func TestRecordSetGeneralRepresentationAndMapAliases(t *testing.T) {
	valueMap := map[string]Value{"old": Int(1), "other": Int(2)}
	anyMap := map[string]any{"old": int64(1), "other": int64(2)}
	boxedValues := any(valueMap)
	boxedAny := any(anyMap)
	cases := map[string]Value{
		"record data":     RecordDict([]string{"old", "other"}, []Value{Int(1), Int(2)}),
		"map record":      Record(valueMap),
		"boxed value map": {Type: TypeAny, UnsafePtr: unsafe.Pointer(&boxedValues)},
		"boxed any map":   {Type: TypeAny, UnsafePtr: unsafe.Pointer(&boxedAny)},
		"empty data":      RecordDict(nil, nil),
		"nil map":         Record(nil),
	}
	for name, orig := range cases {
		t.Run(name, func(t *testing.T) {
			before := RecordToMap(orig)
			beforeLen := len(before)
			inserted := RecordSet(orig, "new", Int(3))
			replaced := RecordSet(inserted, "new", Int(4))
			if len(RecordToMap(orig)) != beforeLen || len(RecordToMap(inserted)) != beforeLen+1 {
				t.Fatal("insertion changed original or lost fields")
			}
			if RecordGet(inserted, "new").IntVal != 3 || RecordGet(replaced, "new").IntVal != 4 {
				t.Fatal("replacement changed original or failed")
			}
			if beforeLen > 0 {
				updated := RecordSet(orig, "old", Int(5))
				if RecordGet(orig, "old").IntVal != 1 || RecordGet(updated, "old").IntVal != 5 || RecordGet(updated, "other").IntVal != 2 {
					t.Fatal("existing key update lost fields or mutated source")
				}
			}
		})
	}
	if len(valueMap) != 2 || valueMap["old"].IntVal != 1 || len(anyMap) != 2 || anyMap["old"] != int64(1) {
		t.Fatal("source map alias was mutated")
	}
}

//go:noinline
func retainedRecordSetClosure() Value {
	captured := []int64{37}
	closure := Func(func(arg Value) Value { return Int(captured[0] + arg.IntVal) })
	record := RecordSet(RecordDict0(), "closure", closure)
	for i := 0; i < 7; i++ {
		record = RecordSet(record, fmt.Sprintf("field%d", i), Int(int64(i)))
	}
	return RecordSet(record, "field2", Int(99))
}

func TestRecordSetRetainsClosureAcrossGC(t *testing.T) {
	record := retainedRecordSetClosure()
	for i := 0; i < 3; i++ {
		runtime.GC()
	}
	if got := Apply(RecordGet(record, "closure"), Int(5)).IntVal; got != 42 {
		t.Fatalf("closure after promotion and GC: %d", got)
	}
	if RecordGet(record, "field2").IntVal != 99 || RecordGet(record, "field6").IntVal != 6 {
		t.Fatal("fields lost after GC")
	}
}
