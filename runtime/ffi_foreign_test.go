package gopurs_runtime

import (
	"math"
	"reflect"
	"testing"
)

func sameForeignMap(left, right map[string]any) bool {
	return reflect.ValueOf(left).UnsafePointer() == reflect.ValueOf(right).UnsafePointer()
}

func TestForeignValueUnboxesLegacyNativeContainersOnDemand(t *testing.T) {
	clean := map[string]any{"large": uint64(math.MaxUint64), "zero": 0, "false": false}
	shared := []any{Int(1), map[string]any{"value": Str("text")}}
	original := map[string]any{"clean": clean, "first": shared, "second": shared, "nil": nil}
	result := ForeignValue(Any(original)).(map[string]any)
	if sameForeignMap(result, original) {
		t.Fatal("mixed container was mutated instead of converted")
	}
	if !sameForeignMap(result["clean"].(map[string]any), clean) {
		t.Fatal("already-native subtree was unnecessarily copied")
	}
	first, second := result["first"].([]any), result["second"].([]any)
	if &first[0] != &second[0] || first[0] != int64(1) || first[1].(map[string]any)["value"] != "text" {
		t.Fatal("converted values or shared-child identity were lost")
	}
	if _, stillBoxed := shared[0].(Value); !stillBoxed {
		t.Fatal("input slice was changed")
	}
	if _, present := result["absent"]; present || result["nil"] != nil {
		t.Fatal("missing and nil values were confused")
	}
	if clean["large"] != uint64(math.MaxUint64) || clean["zero"] != 0 || clean["false"] != false {
		t.Fatal("native numeric types or values were changed")
	}
}

func TestForeignValueKeepsNaturalCyclesAndOpaquePayloads(t *testing.T) {
	cycle := map[string]any{}
	cycle["self"] = cycle
	slice := make([]any, 1)
	slice[0] = slice
	cycle["slice"] = slice
	type opaque struct{ internal Value }
	object := &opaque{internal: Int(17)}
	cycle["object"] = object
	result := ForeignValue(Any(cycle)).(map[string]any)
	if !sameForeignMap(result, cycle) || &result["slice"].([]any)[0] != &slice[0] || result["object"] != object {
		t.Fatal("natural cyclic data or opaque object identity was changed")
	}
}

func TestForeignValueConvertsMixedCyclesWithoutBreakingBackEdges(t *testing.T) {
	root := map[string]any{}
	child := map[string]any{"root": root, "value": Int(42)}
	slice := []any{root, child, nil}
	slice[2] = slice
	root["child"], root["slice"] = child, slice
	result := ForeignValue(Any(root)).(map[string]any)
	convertedChild := result["child"].(map[string]any)
	convertedSlice := result["slice"].([]any)
	if sameForeignMap(result, root) || sameForeignMap(convertedChild, child) || convertedChild["value"] != int64(42) {
		t.Fatal("mixed graph was not converted")
	}
	if !sameForeignMap(convertedChild["root"].(map[string]any), result) || !sameForeignMap(convertedSlice[0].(map[string]any), result) {
		t.Fatal("back-edge still refers to the unconverted graph")
	}
	if !sameForeignMap(convertedSlice[1].(map[string]any), convertedChild) || &convertedSlice[2].([]any)[0] != &convertedSlice[0] {
		t.Fatal("sharing or slice self-cycle was lost")
	}
}

func TestForeignValueConvertsBoxedContainersInsideNativeCycles(t *testing.T) {
	root := map[string]any{}
	boxedArray := Array([]Value{Any(root), Int(3)})
	root["array"] = boxedArray
	root["record"] = RecordDict1("root", Any(root))
	result := ForeignValue(Any(root)).(map[string]any)
	array := result["array"].([]any)
	record := result["record"].(map[string]any)
	if array[1] != int64(3) || !sameForeignMap(array[0].(map[string]any), result) || !sameForeignMap(record["root"].(map[string]any), result) {
		t.Fatal("boxed/native cycle was not converted consistently")
	}
}

func TestForeignValuePreservesTypedNilAndEmptyContainers(t *testing.T) {
	var nilSlice []any
	var nilMap map[string]any
	var nilPointer *int
	values := map[string]any{"nilSlice": nilSlice, "nilMap": nilMap, "nilPointer": nilPointer, "emptySlice": []any{}, "emptyMap": map[string]any{}, "boxed": Int(0)}
	result := ForeignValue(Any(values)).(map[string]any)
	for _, key := range []string{"nilSlice", "nilMap", "nilPointer"} {
		if reflect.TypeOf(result[key]) != reflect.TypeOf(values[key]) || !reflect.ValueOf(result[key]).IsNil() {
			t.Fatalf("typed nil %s was changed", key)
		}
	}
	for _, key := range []string{"emptySlice", "emptyMap"} {
		if reflect.ValueOf(result[key]).IsNil() || reflect.ValueOf(result[key]).Len() != 0 {
			t.Fatalf("empty %s became nil", key)
		}
	}
}

func TestForeignValuePreservesUndefinedAcrossNativeRoundTrip(t *testing.T) {
	undefined := ForeignValue(Value{})
	if undefined == nil || !IsForeignUndefined(undefined) {
		t.Fatal("undefined became null")
	}
	if _, boxed := undefined.(Value); boxed {
		t.Fatal("undefined leaked a boxed runtime Value")
	}
	if Any(undefined).Type != 0 || Box(undefined).Type != 0 {
		t.Fatal("native undefined was not restored when passed back to PureScript")
	}
	if null := ForeignValue(Any(nil)); null != nil || Any(null).Type != TypeAny {
		t.Fatal("null no longer has its distinct representation")
	}
	for _, other := range []any{nil, Value{}, map[string]any{}, []any{}, struct{}{}} {
		if IsForeignUndefined(other) {
			t.Fatalf("%T was mistaken for the native undefined marker", other)
		}
	}
}

func TestForeignValuePreservesNestedUndefinedWithoutChangingTheGraph(t *testing.T) {
	trace := map[string]any{"cause": Value{}, "null": Any(nil), "run": Str("test")}
	meta := map[string]any{"trace": trace}
	input := map[string]any{"first": meta, "second": meta}
	input["self"] = input
	result := ForeignValue(Any(input)).(map[string]any)
	first := result["first"].(map[string]any)
	second := result["second"].(map[string]any)
	converted := first["trace"].(map[string]any)
	undefined, present := converted["cause"]
	if !present || undefined == nil || Any(undefined).Type != 0 {
		t.Fatal("the undefined property was removed or became null")
	}
	if _, boxed := undefined.(Value); boxed {
		t.Fatal("nested undefined leaked a boxed runtime Value")
	}
	if converted["null"] != nil || converted["run"] != "test" {
		t.Fatal("native data was changed")
	}
	if _, present := converted["missing"]; present {
		t.Fatal("an absent property was created")
	}
	if !sameForeignMap(first, second) || !sameForeignMap(result["self"].(map[string]any), result) {
		t.Fatal("shared children or cycle identity were changed")
	}
	if _, boxed := trace["cause"].(Value); !boxed {
		t.Fatal("the source object was mutated")
	}
	if again := ForeignValue(Any(result)).(map[string]any); !sameForeignMap(result, again) {
		t.Fatal("an already-native graph was unnecessarily copied")
	}
}

func TestForeignValuePreservesEventStoreUndefinedParameter(t *testing.T) {
	params := Array([]Value{Array([]Value{Record(map[string]Value{
		"trace": Record(map[string]Value{"run": Str("test"), "cause": {}}),
	})})})
	result := ForeignValue(params).([]any)
	trace := result[0].([]any)[0].(map[string]any)["trace"].(map[string]any)
	if cause, present := trace["cause"]; !present || cause == nil || Any(cause).Type != 0 {
		t.Fatal("the EventStore optional trace field did not cross intact")
	}
}

func TestForeignValueStillRejectsFunctionsAndConstructors(t *testing.T) {
	values := map[string]Value{
		"function":    Func(func(value Value) Value { return value }),
		"constructor": Constructor1("Payload", Int(42)),
	}
	for name, value := range values {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("a non-data value silently crossed the native Foreign boundary")
				}
			}()
			ForeignValue(Any(map[string]any{"nested": value}))
		})
	}
}
