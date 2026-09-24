package gopurs_runtime

import "testing"

type testJSONEntry struct {
	key   string
	value any
}
type testJSONObject []testJSONEntry

func (o testJSONObject) JSONLength() int               { return len(o) }
func (o testJSONObject) JSONEntry(i int) (string, any) { return o[i].key, o[i].value }
func (o testJSONObject) JSONLookup(key string) (any, bool) {
	for _, entry := range o {
		if entry.key == key {
			return entry.value, true
		}
	}
	return nil, false
}

func TestJSONObjectOwnedMaterialization(t *testing.T) {
	original := testJSONObject{{"null", nil}, {"number", float64(12)}, {"text", "kept"}}
	boxed := Box(original)
	view, ok := ReadJSONObject(boxed)
	if !ok || view.Length() != 3 {
		t.Fatal("not an object")
	}
	if value, present := view.Lookup("null"); !present || value != nil {
		t.Fatal("lost null field")
	}
	if _, present := view.Lookup("missing"); present {
		t.Fatal("invented missing field")
	}
	first, second := UnboxObject(boxed), UnboxObject(boxed)
	first["text"] = "changed"
	delete(first, "null")
	if second["text"] != "kept" {
		t.Fatal("materializations share a map")
	}
	if value, _ := original.JSONLookup("text"); value != "kept" {
		t.Fatal("materialization mutates input")
	}
	if _, present := original.JSONLookup("null"); !present {
		t.Fatal("delete mutates input")
	}
	if RecordGet(boxed, "number").FloatVal() != 12 {
		t.Fatal("record getter lost value")
	}
	updated := RecordSet(boxed, "text", Str("new"))
	if RecordGet(updated, "text").StrVal() != "new" || RecordGet(boxed, "text").StrVal() != "kept" {
		t.Fatal("record update aliasing")
	}
	copy := RecordToMap(boxed)
	copy["number"] = Float(0)
	if RecordGet(boxed, "number").FloatVal() != 12 {
		t.Fatal("boxed map aliases input")
	}
}

func TestJSONObjectViewsAcceptExistingRepresentations(t *testing.T) {
	for _, raw := range []any{map[string]any{"x": float64(4)}, map[string]Value{"x": Float(4)}, RecordDict1("x", Float(4)), testJSONObject{{"x", float64(4)}}} {
		view, ok := ReadJSONObject(raw)
		if !ok || view.Length() != 1 {
			t.Fatalf("not an object: %T", raw)
		}
		value, present := view.Lookup("x")
		if !present || Box(value).FloatVal() != 4 {
			t.Fatalf("wrong lookup: %T", raw)
		}
		count := 0
		view.Each(func(key string, value any) {
			count++
			if key != "x" || Box(value).FloatVal() != 4 {
				t.Fatal("wrong entry")
			}
		})
		if count != 1 {
			t.Fatal("wrong entry count")
		}
	}
	for _, raw := range []any{nil, true, "string", []any{}, Float(1), Array(nil)} {
		if _, ok := ReadJSONObject(raw); ok {
			t.Fatalf("non-object accepted: %T", raw)
		}
	}
}
