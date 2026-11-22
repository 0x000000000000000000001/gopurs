package gopurs_runtime

import (
	"reflect"
	"strings"
	"testing"
)

func TestRecordShapesLookupMapAndUnbox(t *testing.T) {
	v0, v1, v2, v3, v4 := Int(10), Int(20), Int(30), Int(40), Int(50)
	cases := []struct {
		name            string
		keys            []string
		values          []Value
		static, dynamic Value
	}{
		{"one", []string{"z"}, []Value{v0}, RecordShape1(&RecordKeys1{"z"}, v0), RecordDict1("z", v0)},
		{"two", []string{"b", "a"}, []Value{v0, v1}, RecordShape2(&RecordKeys2{"b", "a"}, v0, v1), RecordDict2("b", "a", v0, v1)},
		{"three", []string{"c", "a", "b"}, []Value{v0, v1, v2}, RecordShape3(&RecordKeys3{"c", "a", "b"}, v0, v1, v2), RecordDict3("c", "a", "b", v0, v1, v2)},
		{"four", []string{"d", "b", "a", "c"}, []Value{v0, v1, v2, v3}, RecordShape4(&RecordKeys4{"d", "b", "a", "c"}, v0, v1, v2, v3), RecordDict4("d", "b", "a", "c", v0, v1, v2, v3)},
		{"five", []string{"e", "d", "b", "a", "c"}, []Value{v0, v1, v2, v3, v4}, RecordShape5(&RecordKeys5{"e", "d", "b", "a", "c"}, v0, v1, v2, v3, v4), RecordDict5("e", "d", "b", "a", "c", v0, v1, v2, v3, v4)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for _, record := range []Value{tc.static, tc.dynamic} {
				m := RecordToMap(record)
				u := UnboxObject(record)
				if len(m) != len(tc.keys) || len(u) != len(tc.keys) {
					t.Fatal("wrong map size")
				}
				for i, key := range tc.keys {
					if got := RecordGet(record, key); got != tc.values[i] {
						t.Fatalf("lookup %q", key)
					}
					if got := m[key]; got != tc.values[i] {
						t.Fatalf("map %q", key)
					}
					if got := u[key]; got != tc.values[i] {
						t.Fatalf("unbox %q: %#v", key, got)
					}
				}
				func() {
					defer func() {
						if recover() == nil {
							t.Error("missing key did not panic")
						}
					}()
					RecordGet(record, "missing")
				}()
			}
			if !reflect.DeepEqual(RecordToMap(tc.static), RecordToMap(tc.dynamic)) {
				t.Fatal("static/dynamic mismatch")
			}
		})
	}
}

func TestRecordShapesSharingAndUpdates(t *testing.T) {
	shape := RecordKeys2{K0: "b", K1: "a"}
	one := RecordShape2(&shape, Int(1), Int(2))
	two := RecordShape2(&shape, Int(3), Int(4))
	set := RecordSet(one, "a", Int(20))
	updated := RecordUpdateDict(one, []string{"b"}, []Value{Int(10)})
	for _, record := range []Value{one, two, set, updated} {
		keys, _, _ := RecordFields2(record)
		if keys != &shape {
			t.Fatal("shape was not shared")
		}
	}
	if one.UnsafePtr == set.UnsafePtr || one.UnsafePtr == updated.UnsafePtr {
		t.Fatal("update reused record storage")
	}
	for _, tc := range []struct {
		record Value
		b, a   int64
	}{{one, 1, 2}, {two, 3, 4}, {set, 1, 20}, {updated, 10, 2}} {
		if RecordGet(tc.record, "b").IntVal != tc.b || RecordGet(tc.record, "a").IntVal != tc.a {
			t.Fatal("shared record mutated")
		}
	}
	if shape != (RecordKeys2{K0: "b", K1: "a"}) {
		t.Fatal("shape mutated")
	}
	extended := RecordSet(one, "c", Int(30))
	if extended.Type != TypeRecord3 || RecordGet(extended, "c").IntVal != 30 {
		t.Fatal("extension failed")
	}
	if len(RecordToMap(one)) != 2 {
		t.Fatal("extension mutated source")
	}
	coerced := CoerceToStruct[struct {
		Rc     uint32
		V0, V1 Value
	}](one)
	if coerced.V0.IntVal != 2 || coerced.V1.IntVal != 1 {
		t.Fatal("coerce/sorted field order changed")
	}
}

func TestRecordShapesDynamicFallbackAndOrder(t *testing.T) {
	key := strings.Join([]string{"run", "time"}, "-")
	dynamic := RecordDict2(key, "fixed", Str("value"), Int(9))
	if RecordGet(dynamic, key).StrVal() != "value" || RecordGet(dynamic, "fixed").IntVal != 9 {
		t.Fatal("dynamic key lost")
	}
	shape := RecordKeys2{K0: "same", K1: "same"}
	duplicate := RecordShape2(&shape, Int(1), Int(2))
	if RecordGet(duplicate, "same").IntVal != 1 || RecordToMap(duplicate)["same"].IntVal != 2 {
		t.Fatal("key order semantics changed")
	}
	five := RecordShape5(&RecordKeys5{"a", "b", "c", "d", "e"}, Int(1), Int(2), Int(3), Int(4), Int(5))
	six := RecordSet(five, "f", Int(6))
	if len(RecordToMap(six)) != 6 || RecordGet(six, "f").IntVal != 6 || len(RecordToMap(five)) != 5 {
		t.Fatal("map extension fallback failed")
	}
}

var recordShapeAllocationSink Value

func TestRecordShapesAllocationCounts(t *testing.T) {
	shape := RecordKeys2{"a", "b"}
	static := testing.AllocsPerRun(100, func() { recordShapeAllocationSink = RecordShape2(&shape, Int(1), Int(2)) })
	dynamic := testing.AllocsPerRun(100, func() { recordShapeAllocationSink = RecordDict2("a", "b", Int(1), Int(2)) })
	if static != 1 || dynamic != 2 {
		t.Fatalf("allocation counts static=%v dynamic=%v", static, dynamic)
	}
	t.Logf("allocation counts static=%v dynamic=%v", static, dynamic)
}
