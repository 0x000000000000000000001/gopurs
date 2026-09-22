package gopurs_runtime

import (
	"runtime"
	"testing"
	"unsafe"
)

type singleRecordField struct {
	Rc uint32
	V0 Value
}

func TestCoerceCompactRecord(t *testing.T) {
	if got := CoerceToStruct[singleRecordField](RecordDict0()); *got != (singleRecordField{}) {
		t.Fatal("empty record must produce zero fields")
	}
	for _, key := range []string{"", "decodeJson", "é雪"} {
		original := RecordDict1(key, Int(42))
		first := CoerceToStruct[singleRecordField](original)
		second := CoerceToStruct[singleRecordField](original)
		if first.Rc != 0 || first.V0 != Int(42) || first == second {
			t.Fatalf("%q: incorrect value or shared mutable struct", key)
		}
		first.V0 = Int(7)
		if second.V0 != Int(42) || RecordGet(original, key) != Int(42) {
			t.Fatalf("%q: coercion mutated or aliased the record", key)
		}
	}
}

func TestCoerceConstructorAndSortedFallback(t *testing.T) {
	original := &singleRecordField{Rc: 1, V0: Int(42)}
	boxed := Value{Type: TypeConstructor, UnsafePtr: unsafe.Pointer(original)}
	if CoerceToStruct[singleRecordField](boxed) != original {
		t.Fatal("constructor coercion must retain its pointer")
	}
	if CoerceToStruct[singleRecordField](Value{Type: TypeConstructor}) != nil {
		t.Fatal("null constructor pointer changed")
	}
	type twoFields struct {
		Rc     uint32
		V0, V1 Value
	}
	for _, input := range []Value{
		RecordDict2("z", "a", Int(9), Int(1)),
		Record(map[string]Value{"z": Int(9), "a": Int(1)}),
	} {
		got := CoerceToStruct[twoFields](input)
		if got.Rc != 0 || got.V0 != Int(1) || got.V1 != Int(9) {
			t.Fatal("general coercion must retain alphabetical field order")
		}
	}
}

//go:noinline
func coercedClosure() *singleRecordField {
	capture := []int64{37}
	return CoerceToStruct[singleRecordField](RecordDict1("run",
		Func(func(arg Value) Value { return Int(capture[0] + arg.IntVal) })))
}

func TestCoerceCompactRecordRetainsClosure(t *testing.T) {
	record := coercedClosure()
	runtime.GC()
	if Apply(record.V0, Int(5)) != Int(42) {
		t.Fatal("coerced closure lost its captured value")
	}
}
