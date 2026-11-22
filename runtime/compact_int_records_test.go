package gopurs_runtime

import (
	"math"
	"reflect"
	"testing"
	"unsafe"
)

var compactIntRecordSink Value

func TestCompactIntRecordsAllocation(t *testing.T) {
	shape := &RecordKeys2{"a", "b"}
	makeRecord := func() { compactIntRecordSink = RecordShape2(shape, Int(1), Int(2)) }
	allocs := testing.AllocsPerRun(100, makeRecord)
	wantSize := unsafe.Sizeof((*RecordKeys2)(nil)) + 2*unsafe.Sizeof(int64(0))
	if unsafe.Sizeof(RecordData2Int{}) != wantSize || allocs != 1 {
		t.Fatalf("size=%d want=%d allocs=%v", unsafe.Sizeof(RecordData2Int{}), wantSize, allocs)
	}
}

func TestCompactIntRecordsLimitsAndAPIs(t *testing.T) {
	shape := &RecordKeys2{"left", "right"}
	pairs := [][2]int64{{math.MinInt64, math.MaxInt64}, {-1, 0}, {math.MinInt32, math.MaxInt32}, {4294967295, 2147483648}}
	for _, pair := range pairs {
		record := RecordShape2(shape, Int(pair[0]), Int(pair[1]))
		if record.Type != TypeRecord2 || record.IntVal != record2IntLayout {
			t.Fatal("wrong representation")
		}
		if RecordGet(record, "left").IntVal != pair[0] || RecordGet(record, "right").IntVal != pair[1] {
			t.Fatal("integer narrowed")
		}
		if RecordAt2(record, 0) != Int(pair[0]) || RecordAt2(record, 1) != Int(pair[1]) {
			t.Fatal("indexed access changed")
		}
		want := map[string]Value{"left": Int(pair[0]), "right": Int(pair[1])}
		if !reflect.DeepEqual(RecordToMap(record), want) || !reflect.DeepEqual(record.AnyVal(), want) {
			t.Fatal("map/AnyVal changed")
		}
		for _, unboxed := range []map[string]any{UnboxObject(record), Unbox[map[string]any](record)} {
			if unboxed["left"] != want["left"] || unboxed["right"] != want["right"] {
				t.Fatal("unboxing changed")
			}
		}
		view, ok := ReadJSONObjectValue(record)
		if !ok || view.Length() != 2 {
			t.Fatal("JSON view rejected record")
		}
		if value, exists := view.Lookup("right"); !exists || value != want["right"] {
			t.Fatal("JSON lookup changed")
		}
		coerced := CoerceToStruct[struct {
			Rc     uint32
			V0, V1 Value
		}](record)
		if coerced.V0 != want["left"] || coerced.V1 != want["right"] {
			t.Fatal("coercion changed")
		}
	}
}

func TestCompactIntRecordsUpdatesAndCoexistence(t *testing.T) {
	shape := &RecordKeys2{"a", "b"}
	initial := RecordShape2(shape, Int(1), Int(2))
	generic := RecordSet(initial, "a", Str("text"))
	compact := RecordSet(generic, "a", Int(math.MaxInt64))
	updatedGeneric := RecordUpdateDict(initial, []string{"b"}, []Value{Str("other")})
	updatedCompact := RecordUpdateDict(updatedGeneric, []string{"b"}, []Value{Int(math.MinInt64)})
	for _, record := range []Value{initial, compact, updatedCompact} {
		if record.IntVal != record2IntLayout {
			t.Fatal("lost compact layout")
		}
	}
	for _, record := range []Value{generic, updatedGeneric} {
		if record.IntVal != 0 {
			t.Fatal("generic layout not restored")
		}
	}
	for _, record := range []Value{initial, generic, compact, updatedGeneric, updatedCompact} {
		keys, _, _ := RecordFields2(record)
		if keys != shape {
			t.Fatal("shape changed")
		}
		if record.Type != TypeRecord2 {
			t.Fatal("public type changed")
		}
	}
	if RecordGet(initial, "a").IntVal != 1 || RecordGet(initial, "b").IntVal != 2 {
		t.Fatal("original mutated")
	}
	if RecordGet(generic, "a").StrVal() != "text" || RecordGet(updatedGeneric, "b").StrVal() != "other" {
		t.Fatal("generic values changed")
	}
	if RecordAt2(generic, 0).StrVal() != "text" || RecordAt2(generic, 1).IntVal != 2 {
		t.Fatal("generic indexed access changed")
	}
	if RecordGet(compact, "a").IntVal != math.MaxInt64 || RecordGet(updatedCompact, "b").IntVal != math.MinInt64 {
		t.Fatal("compact values changed")
	}
	extended := RecordSet(initial, "c", Int(3))
	if extended.Type != TypeRecord3 || len(RecordToMap(extended)) != 3 {
		t.Fatal("extension failed")
	}
	if RecordGet(extended, "a").IntVal != 1 || RecordGet(extended, "b").IntVal != 2 {
		t.Fatal("extension lost fields")
	}
	ignored := RecordUpdateDict(initial, []string{"missing"}, []Value{Int(99)})
	if !reflect.DeepEqual(RecordToMap(ignored), RecordToMap(initial)) {
		t.Fatal("unknown update changed semantics")
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Error("missing key did not panic")
			}
		}()
		RecordGet(initial, "missing")
	}()
	legacy := Value{Type: TypeRecord2, UnsafePtr: unsafe.Pointer(&RecordData2{shape, Int(8), Int(9)})}
	if RecordGet(legacy, "a").IntVal != 8 || RecordGet(legacy, "b").IntVal != 9 {
		t.Fatal("legacy generic layout broken")
	}
}

func TestCompactIntRecordsForeignGraphAndVariant(t *testing.T) {
	compact := RecordShape2(&RecordKeys2{"a", "b"}, Int(math.MinInt64), Int(math.MaxInt64))
	parent := map[string]any{"first": compact, "second": compact}
	parent["self"] = parent
	converted := ForeignValue(Any(parent)).(map[string]any)
	first := converted["first"].(map[string]any)
	second := converted["second"].(map[string]any)
	if first["a"] != int64(math.MinInt64) || first["b"] != int64(math.MaxInt64) {
		t.Fatal("Foreign values changed")
	}
	if reflect.ValueOf(first).UnsafePointer() != reflect.ValueOf(second).UnsafePointer() {
		t.Fatal("Foreign sharing lost")
	}
	if reflect.ValueOf(converted).UnsafePointer() != reflect.ValueOf(converted["self"]).UnsafePointer() {
		t.Fatal("Foreign cycle lost")
	}
	if _, _, ok := ExtractVariant(compact); ok {
		t.Fatal("integer record accepted as variant")
	}
	variant := RecordShape2(&RecordKeys2{"type", "value"}, Str("Example"), compact)
	name, payload, ok := ExtractVariant(variant)
	if !ok || name != "Example" || payload != compact {
		t.Fatal("generic variant broken")
	}
}
