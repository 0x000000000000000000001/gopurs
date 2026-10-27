package Record

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"testing"

	"gopurs/output/gopurs_runtime"
)

var testRight = gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.RecordDict2("ok", "value", gopurs_runtime.Bool(true), value)
})

var testLeft = gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.RecordDict2("ok", "value", gopurs_runtime.Bool(false), value)
})

var testIsRight = gopurs_runtime.Func(func(result gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.RecordGet(result, "ok")
})

var testRightValue = gopurs_runtime.Func(func(result gopurs_runtime.Value) gopurs_runtime.Value {
	if !gopurs_runtime.RecordGet(result, "ok").BoolVal() {
		panic("extracted an error as a successful field")
	}
	return gopurs_runtime.RecordGet(result, "value")
})

func testField(name string, value int64, tail gopurs_runtime.Value) gopurs_runtime.Value {
	key := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value { return gopurs_runtime.Str(name) })
	step := gopurs_runtime.Func2(func(key, object gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(testRight, gopurs_runtime.Int(value))
	})
	return testCons(key, step, tail, gopurs_runtime.Value{})
}

var testErrAtKey = gopurs_runtime.Func2(func(key, err gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.RecordDict2("atKey", "inner", key, err)
})

var testErrMissing = gopurs_runtime.RecordDict1("missing", gopurs_runtime.Bool(true))

var testErrTypeMismatch = gopurs_runtime.Func(func(message gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.RecordDict2("mismatch", "message", gopurs_runtime.Bool(true), message)
})

var testErrNamed = gopurs_runtime.Func2(func(name, err gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.RecordDict2("named", "inner", name, err)
})

var testErrAtIndex = gopurs_runtime.Func2(func(index, err gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.RecordDict2("atIndex", "inner", index, err)
})

var testNothing = gopurs_runtime.RecordDict1("nothing", gopurs_runtime.Bool(true))

var testJust = gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.RecordDict1("just", value)
})

var testLeftOf = gopurs_runtime.Func(func(either gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.RecordGet(either, "value")
})

var testErrorSupport = gopurs_runtime.RecordDict(
	[]string{"atKey", "missingValue", "typeMismatch", "named", "atIndex", "leftOf",
		"nothing", "just", "isRight", "rightValue", "left", "right"},
	[]gopurs_runtime.Value{testErrAtKey, testErrMissing, testErrTypeMismatch, testErrNamed,
		testErrAtIndex, testLeftOf, testNothing, testJust, testIsRight, testRightValue,
		testLeft, testRight})

func testCons(key, step, tail, fallback gopurs_runtime.Value) gopurs_runtime.Value {
	return RecordConsImpl(key, step, tail, fallback, testIsRight, testRightValue, testRight,
		testLeft, testErrorSupport, testNothing, testJust)
}

func testNil() gopurs_runtime.Value {
	return RecordNilImpl(testLeft, testErrorSupport, testRight)
}

func testDecode(plan gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply2(plan, gopurs_runtime.Value{}, gopurs_runtime.Value{})
}

func TestRecordPlanConstructionAndLifetime(t *testing.T) {
	for _, count := range []int{0, 1, 2, 5, 6, 32} {
		plan := testNil()
		for index := count - 1; index >= 0; index-- {
			plan = testField(fmt.Sprint(index), int64(index), plan)
		}
		runtime.GC()
		partial := gopurs_runtime.Apply(plan, gopurs_runtime.Value{})
		result := gopurs_runtime.Apply(partial, gopurs_runtime.Value{})
		if !gopurs_runtime.RecordGet(result, "ok").BoolVal() {
			t.Fatal("unexpected error")
		}
		record := gopurs_runtime.RecordGet(result, "value")
		for index := 0; index < count; index++ {
			if got := gopurs_runtime.RecordGet(record, fmt.Sprint(index)).IntVal; got != int64(index) {
				t.Fatalf("field %d: got %d", index, got)
			}
		}
		_ = testDecode(plan)
		runtime.GC()
		if got := len(gopurs_runtime.RecordToMap(record)); got != count {
			t.Fatalf("old result changed: got %d fields, want %d", got, count)
		}
	}
}

func TestRecordPlanEvaluationOrderAndFirstError(t *testing.T) {
	for _, failure := range []string{"", "a", "b"} {
		var events []string
		plan := testNil()
		for _, name := range []string{"b", "a"} {
			name := name
			key := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				events = append(events, "symbol:"+name)
				return gopurs_runtime.Str(name)
			})
			step := gopurs_runtime.Func2(func(key, object gopurs_runtime.Value) gopurs_runtime.Value {
				events = append(events, "decode:"+key.StrVal())
				if name == failure {
					return gopurs_runtime.Apply(testLeft, key)
				}
				return gopurs_runtime.Apply(testRight, gopurs_runtime.Int(1))
			})
			plan = testCons(key, step, plan, gopurs_runtime.Value{})
		}
		if len(events) != 0 {
			t.Fatal("constructing a plan evaluated a callback")
		}
		result := testDecode(plan)
		want := []string{"symbol:a", "decode:a", "symbol:b", "decode:b", "symbol:b", "symbol:a"}
		if failure == "a" {
			want = want[:2]
		} else if failure == "b" {
			want = want[:4]
		}
		if !reflect.DeepEqual(events, want) {
			t.Fatalf("failure=%q events=%v want=%v", failure, events, want)
		}
		if failure != "" && (gopurs_runtime.RecordGet(result, "ok").BoolVal() || gopurs_runtime.RecordGet(result, "value").StrVal() != failure) {
			t.Fatal("first error was not retained")
		}
	}
}

func TestRecordPlanUnknownTailUsesOriginalMethod(t *testing.T) {
	var calls int
	shared := gopurs_runtime.RecordDict1("existing", gopurs_runtime.Int(42))
	fallback := gopurs_runtime.Func2(func(object, proxy gopurs_runtime.Value) gopurs_runtime.Value {
		calls++
		if object.IntVal != 12 || proxy.IntVal != 34 {
			t.Fatal("fallback arguments changed")
		}
		return gopurs_runtime.RecordSet(shared, "new", gopurs_runtime.Int(7))
	})
	unknownTail := gopurs_runtime.Func2(func(_, _ gopurs_runtime.Value) gopurs_runtime.Value { panic("called outside fallback") })
	plan := testCons(gopurs_runtime.Value{}, gopurs_runtime.Value{}, unknownTail, fallback)
	partial := gopurs_runtime.Apply(plan, gopurs_runtime.Int(12))
	result := gopurs_runtime.Apply(partial, gopurs_runtime.Int(34))
	if calls != 1 || len(gopurs_runtime.RecordToMap(shared)) != 1 || len(gopurs_runtime.RecordToMap(result)) != 2 {
		t.Fatal("unknown tail behavior or original record changed")
	}
}

func TestRecordPlanDuplicateLabelsPreserveInsertPrecedence(t *testing.T) {
	plan := testField("same", 1, testField("other", 9, testField("same", 2, testNil())))
	record := gopurs_runtime.RecordGet(testDecode(plan), "value")
	if got := gopurs_runtime.RecordGet(record, "same").IntVal; got != 1 {
		t.Fatalf("got tail value %d instead of head value", got)
	}
	if got := len(gopurs_runtime.RecordToMap(record)); got != 2 {
		t.Fatalf("got %d keys", got)
	}
}

func TestRecordPlanUsesSecondSymbolEvaluationForInsertion(t *testing.T) {
	var reflections int
	key := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		reflections++
		if reflections == 1 {
			return gopurs_runtime.Str("lookup")
		}
		return gopurs_runtime.Str("insert")
	})
	step := gopurs_runtime.Func2(func(key, object gopurs_runtime.Value) gopurs_runtime.Value {
		if key.StrVal() != "lookup" {
			t.Fatal("decoder did not receive the first symbol result")
		}
		return gopurs_runtime.Apply(testRight, gopurs_runtime.Int(7))
	})
	plan := testCons(key, step, testNil(), gopurs_runtime.Value{})
	record := gopurs_runtime.RecordGet(testDecode(plan), "value")
	if reflections != 2 || gopurs_runtime.RecordGet(record, "insert").IntVal != 7 {
		t.Fatal("insertion did not reevaluate the symbol")
	}
}

func TestRecordPlanReentrantAndConcurrent(t *testing.T) {
	inner := testField("nested", 99, testNil())
	key := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value { return gopurs_runtime.Str("outer") })
	step := gopurs_runtime.Func2(func(key, object gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(testRight, gopurs_runtime.RecordGet(testDecode(inner), "value"))
	})
	plan := testCons(key, step, testNil(), gopurs_runtime.Value{})
	var wg sync.WaitGroup
	for worker := 0; worker < 8; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for iteration := 0; iteration < 100; iteration++ {
				record := gopurs_runtime.RecordGet(testDecode(plan), "value")
				innerRecord := gopurs_runtime.RecordGet(record, "outer")
				if got := gopurs_runtime.RecordGet(innerRecord, "nested").IntVal; got != 99 {
					t.Errorf("shared state corrupted: %d", got)
				}
			}
		}()
	}
	wg.Wait()
}

// ---------------------------------------------------------------------------
// Typed direct path
// ---------------------------------------------------------------------------

func testKey(name string) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value { return gopurs_runtime.Str(name) })
}

// testGenericStep fails loudly if the generic step is used, so the direct-path
// tests prove the DOM fast path ran.
func testGenericStep(tag string) gopurs_runtime.Value {
	step := gopurs_runtime.Func2(func(key, object gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(testLeft, gopurs_runtime.Str("generic:"+tag))
	})
	return TagDecoder(step, &fieldKind{tag: tag}, testErrorSupport)
}

func testTypedPlan() gopurs_runtime.Value {
	arrayString := TagDecoder(testGenericStep(kindArray),
		&fieldKind{tag: kindArray, inner: &fieldKind{tag: kindString}}, testErrorSupport)
	plan := testNil()
	plan = testCons(testKey("tags"), arrayString, plan, gopurs_runtime.Value{})
	plan = testCons(testKey("label"), testGenericStep(kindString), plan, gopurs_runtime.Value{})
	plan = testCons(testKey("count"), testGenericStep(kindInt), plan, gopurs_runtime.Value{})
	return plan
}

func TestRecordPlanTypedDirectValues(t *testing.T) {
	plan := testTypedPlan()
	dom := gopurs_runtime.Any(map[string]any{
		"count": float64(7),
		"label": "ok",
		"tags":  []any{"a", "b"},
	})
	result := gopurs_runtime.Apply(gopurs_runtime.Apply(plan, dom), gopurs_runtime.Value{})
	if !gopurs_runtime.RecordGet(result, "ok").BoolVal() {
		t.Fatalf("direct path failed: %s", gopurs_runtime.RecordGet(result, "value").StrVal())
	}
	record := gopurs_runtime.RecordGet(result, "value")
	if got := gopurs_runtime.RecordGet(record, "count").IntVal; got != 7 {
		t.Fatalf("count: got %d", got)
	}
	if got := gopurs_runtime.RecordGet(record, "label").StrVal(); got != "ok" {
		t.Fatalf("label: got %q", got)
	}
	tags := gopurs_runtime.RecordGet(record, "tags")
	if gopurs_runtime.ArrayLength(tags) != 2 || gopurs_runtime.ArrayAccess(tags, 1).StrVal() != "b" {
		t.Fatal("tags were not decoded from the DOM")
	}
}

func TestTypedSymbolsAreReevaluatedForInsertionAndEveryCall(t *testing.T) {
	var calls int
	key := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		calls++
		return gopurs_runtime.Str(fmt.Sprintf("key%d", calls))
	})
	plan := testCons(key, testGenericStep(kindInt), testNil(), gopurs_runtime.Value{})
	if calls != 0 {
		t.Fatal("construction evaluated a symbol")
	}
	var retained []gopurs_runtime.Value
	for pass := 0; pass < 2; pass++ {
		lookup := fmt.Sprintf("key%d", pass*2+1)
		insert := fmt.Sprintf("key%d", pass*2+2)
		input := gopurs_runtime.Any(map[string]any{lookup: float64(pass + 7)})
		result := gopurs_runtime.Apply2(plan, input, gopurs_runtime.Value{})
		if !gopurs_runtime.RecordGet(result, "ok").BoolVal() {
			t.Fatal("lookup used a cached symbol")
		}
		record := gopurs_runtime.RecordGet(result, "value")
		if calls != (pass+1)*2 || gopurs_runtime.RecordGet(record, insert).IntVal != int64(pass+7) {
			t.Fatal("insertion did not use the second symbol result")
		}
		retained = append(retained, record)
	}
	runtime.GC()
	if got := gopurs_runtime.RecordToMap(retained[0]); len(got) != 1 || got["key2"].IntVal != 7 {
		t.Fatal("later decoding changed a retained record")
	}
}

func TestTypedSymbolsStopAtFirstFailure(t *testing.T) {
	var events []string
	key := func(name string) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			events = append(events, name)
			return gopurs_runtime.Str(name)
		})
	}
	plan := testCons(key("a"), testGenericStep(kindInt),
		testCons(key("b"), testGenericStep(kindInt), testNil(), gopurs_runtime.Value{}), gopurs_runtime.Value{})
	for _, fail := range []bool{true, false} {
		events = nil
		input := map[string]any{"a": float64(1), "b": float64(2)}
		want := []string{"a", "b", "b", "a"}
		if fail {
			input["a"], input["b"] = "bad a", "bad b"
			want = []string{"a"}
		}
		result := gopurs_runtime.Apply2(plan, gopurs_runtime.Any(input), gopurs_runtime.Value{})
		if !reflect.DeepEqual(events, want) {
			t.Fatalf("events = %v, want %v", events, want)
		}
		if fail && gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(result, "value"), "atKey").StrVal() != "a" {
			t.Fatal("did not retain the first error")
		}
	}
}

func TestRecordPlanTypedDirectErrors(t *testing.T) {
	plan := testTypedPlan()
	wrongType := gopurs_runtime.Any(map[string]any{
		"count": "seven",
		"label": "ok",
		"tags":  []any{},
	})
	result := gopurs_runtime.Apply(gopurs_runtime.Apply(plan, wrongType), gopurs_runtime.Value{})
	if gopurs_runtime.RecordGet(result, "ok").BoolVal() {
		t.Fatal("wrong type was accepted")
	}
	err := gopurs_runtime.RecordGet(result, "value")
	if got := gopurs_runtime.RecordGet(err, "atKey").StrVal(); got != "count" {
		t.Fatalf("atKey: got %q", got)
	}
	inner := gopurs_runtime.RecordGet(err, "inner")
	if !gopurs_runtime.RecordGet(inner, "mismatch").BoolVal() {
		t.Fatal("expected a type mismatch")
	}
	if got := gopurs_runtime.RecordGet(inner, "message").StrVal(); got != "Number" {
		t.Fatalf("message: got %q", got)
	}

	missing := gopurs_runtime.Any(map[string]any{"label": "ok", "tags": []any{}})
	result = gopurs_runtime.Apply(gopurs_runtime.Apply(plan, missing), gopurs_runtime.Value{})
	if gopurs_runtime.RecordGet(result, "ok").BoolVal() {
		t.Fatal("missing field was accepted")
	}
	err = gopurs_runtime.RecordGet(result, "value")
	if got := gopurs_runtime.RecordGet(err, "atKey").StrVal(); got != "count" {
		t.Fatalf("missing atKey: got %q", got)
	}
	if !gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(err, "inner"), "missing").BoolVal() {
		t.Fatal("expected MissingValue")
	}

	badElement := gopurs_runtime.Any(map[string]any{
		"count": float64(1),
		"label": "ok",
		"tags":  []any{"a", float64(2)},
	})
	result = gopurs_runtime.Apply(gopurs_runtime.Apply(plan, badElement), gopurs_runtime.Value{})
	err = gopurs_runtime.RecordGet(result, "value")
	if got := gopurs_runtime.RecordGet(err, "atKey").StrVal(); got != "tags" {
		t.Fatalf("tags atKey: got %q", got)
	}
	named := gopurs_runtime.RecordGet(err, "inner")
	if got := gopurs_runtime.RecordGet(named, "named").StrVal(); got != "Array" {
		t.Fatalf("named: got %q", got)
	}
	atIndex := gopurs_runtime.RecordGet(named, "inner")
	if got := gopurs_runtime.RecordGet(atIndex, "atIndex").IntVal; got != 1 {
		t.Fatalf("atIndex: got %d", got)
	}
}

func TestRecordPlanTypedDirectBoxedValues(t *testing.T) {
	plan := testTypedPlan()
	// Json built with fromObject: every entry is already a boxed Value.
	dom := gopurs_runtime.Any(map[string]any{
		"count": gopurs_runtime.Int(7),
		"label": gopurs_runtime.Str("ok"),
		"tags":  gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Str("a"), gopurs_runtime.Str("b")}),
	})
	result := gopurs_runtime.Apply(gopurs_runtime.Apply(plan, dom), gopurs_runtime.Value{})
	if !gopurs_runtime.RecordGet(result, "ok").BoolVal() {
		t.Fatalf("boxed values failed: %s", gopurs_runtime.RecordGet(result, "value").StrVal())
	}
	record := gopurs_runtime.RecordGet(result, "value")
	if got := gopurs_runtime.RecordGet(record, "count").IntVal; got != 7 {
		t.Fatalf("boxed count: got %d", got)
	}
	if got := gopurs_runtime.RecordGet(record, "label").StrVal(); got != "ok" {
		t.Fatalf("boxed label: got %q", got)
	}
	tags := gopurs_runtime.RecordGet(record, "tags")
	if gopurs_runtime.ArrayLength(tags) != 2 || gopurs_runtime.ArrayAccess(tags, 1).StrVal() != "b" {
		t.Fatal("boxed tags were not decoded")
	}
}

func TestDecodeFieldFastTaggedAndCustom(t *testing.T) {
	tagged := testGenericStep(kindInt)

	// A tagged decoder decodes directly: the generic step would have failed.
	result := DecodeFieldFast(tagged, gopurs_runtime.Int(5))
	if !gopurs_runtime.RecordGet(result, "ok").BoolVal() {
		t.Fatal("tagged field did not take the direct path")
	}
	if got := gopurs_runtime.RecordGet(result, "value").IntVal; got != 5 {
		t.Fatalf("tagged value: got %d", got)
	}

	// Errors are reported without AtKey; the PureScript caller adds it.
	result = DecodeFieldFast(tagged, gopurs_runtime.Str("x"))
	if gopurs_runtime.RecordGet(result, "ok").BoolVal() {
		t.Fatal("wrong type was accepted")
	}
	inner := gopurs_runtime.RecordGet(result, "value")
	if !gopurs_runtime.RecordGet(inner, "mismatch").BoolVal() {
		t.Fatal("expected a type mismatch")
	}
	if got := gopurs_runtime.RecordGet(inner, "message").StrVal(); got != "Number" {
		t.Fatalf("message: got %q", got)
	}

	// Untagged decoders keep their own call.
	custom := gopurs_runtime.Func(func(json gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(testRight, gopurs_runtime.Str("custom"))
	})
	result = DecodeFieldFast(custom, gopurs_runtime.Int(1))
	if got := gopurs_runtime.RecordGet(result, "value").StrVal(); got != "custom" {
		t.Fatalf("custom decoder: got %q", got)
	}

	// Null and Maybe: a tagged Maybe reports Nothing without calling anything.
	maybe := TagDecoder(testGenericStep(kindMaybe), &fieldKind{tag: kindMaybe, inner: &fieldKind{tag: kindInt}}, testErrorSupport)
	result = DecodeFieldFast(maybe, gopurs_runtime.Any(nil))
	if got := gopurs_runtime.RecordGet(result, "value"); got.Type != testNothing.Type || got.IntVal != testNothing.IntVal {
		t.Fatal("tagged Maybe did not produce Nothing")
	}
	result = DecodeFieldFast(maybe, gopurs_runtime.Int(3))
	just := gopurs_runtime.RecordGet(result, "value")
	if just.Type != gopurs_runtime.TypeRecord1 {
		t.Fatalf("tagged Maybe did not produce Just: type %d", just.Type)
	}
	if got := gopurs_runtime.RecordGet(just, "just").IntVal; got != 3 {
		t.Fatalf("Just payload: got %d", got)
	}
}

func TestRecordPlanTypedDirectFallsBackForNonDOM(t *testing.T) {
	plan := testTypedPlan()
	recordInput := gopurs_runtime.RecordDict1("count", gopurs_runtime.Int(3))
	result := gopurs_runtime.Apply(gopurs_runtime.Apply(plan, recordInput), gopurs_runtime.Value{})
	if gopurs_runtime.RecordGet(result, "ok").BoolVal() {
		t.Fatal("generic path was not used")
	}
	if got := gopurs_runtime.RecordGet(result, "value").StrVal(); got != "generic:Int" {
		t.Fatalf("generic step result: %q", got)
	}
}

func TestFieldAccessorFallbackPreservesArgumentsAndResult(t *testing.T) {
	for _, accessor := range []func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value{
		GetFieldImpl, GetFieldOptionalImpl, GetFieldOptionalNullableImpl,
	} {
		decoder := gopurs_runtime.Func(func(gopurs_runtime.Value) gopurs_runtime.Value {
			t.Fatal("untagged decoder was called outside its original accessor")
			return gopurs_runtime.Value{}
		})
		for _, object := range []gopurs_runtime.Value{gopurs_runtime.Any(map[string]any{"field": nil}), gopurs_runtime.Any(map[string]any{}), gopurs_runtime.Int(0)} {
			key := gopurs_runtime.Str("field")
			marker := gopurs_runtime.RecordDict0()
			calls := 0
			fallback := gopurs_runtime.Func3(func(d, o, k gopurs_runtime.Value) gopurs_runtime.Value {
				calls++
				if d != decoder || o != object || k != key {
					t.Fatal("fallback arguments changed")
				}
				return marker
			})
			if result := accessor(fallback, decoder, object, key); result != marker || calls != 1 {
				t.Fatal("fallback result or call count changed")
			}
		}
	}
}

func TestRecordPlanTypedDirectCustomElements(t *testing.T) {
	custom := gopurs_runtime.Func(func(json gopurs_runtime.Value) gopurs_runtime.Value {
		if json.Type == gopurs_runtime.TypeString {
			return gopurs_runtime.Apply(testRight, json)
		}
		return gopurs_runtime.Apply(testLeft, gopurs_runtime.Str("bad element"))
	})
	arrayCustom := TagDecoder(testGenericStep(kindArray),
		&fieldKind{tag: kindArray, custom: custom}, testErrorSupport)
	plan := testCons(testKey("items"), arrayCustom, testNil(), gopurs_runtime.Value{})

	good := gopurs_runtime.Any(map[string]any{"items": []any{"a", "b"}})
	result := gopurs_runtime.Apply(gopurs_runtime.Apply(plan, good), gopurs_runtime.Value{})
	if !gopurs_runtime.RecordGet(result, "ok").BoolVal() {
		t.Fatal("custom elements were rejected")
	}
	record := gopurs_runtime.RecordGet(result, "value")
	items := gopurs_runtime.RecordGet(record, "items")
	if gopurs_runtime.ArrayLength(items) != 2 || gopurs_runtime.ArrayAccess(items, 0).StrVal() != "a" {
		t.Fatal("custom elements were not decoded")
	}

	bad := gopurs_runtime.Any(map[string]any{"items": []any{"a", float64(2)}})
	result = gopurs_runtime.Apply(gopurs_runtime.Apply(plan, bad), gopurs_runtime.Value{})
	err := gopurs_runtime.RecordGet(result, "value")
	if got := gopurs_runtime.RecordGet(err, "atKey").StrVal(); got != "items" {
		t.Fatalf("custom atKey: got %q", got)
	}
	named := gopurs_runtime.RecordGet(err, "inner")
	if got := gopurs_runtime.RecordGet(named, "named").StrVal(); got != "Array" {
		t.Fatalf("custom named: got %q", got)
	}
	atIndex := gopurs_runtime.RecordGet(named, "inner")
	if got := gopurs_runtime.RecordGet(atIndex, "atIndex").IntVal; got != 1 {
		t.Fatalf("custom atIndex: got %d", got)
	}
	if got := gopurs_runtime.RecordGet(atIndex, "inner").StrVal(); got != "bad element" {
		t.Fatalf("custom inner error: %q", got)
	}
}

func testRecordDecoder(row gopurs_runtime.Value) gopurs_runtime.Value {
	return TypedRecord(testErrorSupport, row, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		panic("tagged record wrapper was called by the direct path")
	}))
}

func TestNestedPlansRawBoxedAndCustomFields(t *testing.T) {
	customCalls := 0
	step := gopurs_runtime.Func2(func(key, object gopurs_runtime.Value) gopurs_runtime.Value {
		customCalls++
		fields := gopurs_runtime.UnboxObject(object)
		return gopurs_runtime.Apply(testRight, gopurs_runtime.Box(fields[key.StrVal()]))
	})
	inner := testRecordDecoder(testCons(testKey("x"), testGenericStep(kindInt),
		testCons(testKey("custom"), step, testNil(), gopurs_runtime.Value{}), gopurs_runtime.Value{}))
	field := FieldStep(TypedFieldId(inner, step), step)
	outer := testCons(testKey("nested"), field, testNil(), gopurs_runtime.Value{})
	for _, nested := range []any{
		map[string]any{"x": float64(7), "custom": "raw"},
		gopurs_runtime.Any(map[string]any{"x": gopurs_runtime.Int(7), "custom": gopurs_runtime.Str("raw")}),
		gopurs_runtime.RecordDict2("x", "custom", gopurs_runtime.Int(7), gopurs_runtime.Str("raw")),
	} {
		result := gopurs_runtime.Apply2(outer, gopurs_runtime.Any(map[string]any{"nested": nested}), gopurs_runtime.Value{})
		if !gopurs_runtime.RecordGet(result, "ok").BoolVal() {
			t.Fatal("nested direct record failed")
		}
		record := gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(result, "value"), "nested")
		if gopurs_runtime.RecordGet(record, "x").IntVal != 7 || gopurs_runtime.RecordGet(record, "custom").StrVal() != "raw" {
			t.Fatal("nested custom field received the wrong object representation")
		}
	}
	if customCalls != 3 {
		t.Fatalf("custom calls: %d", customCalls)
	}
}

func TestNestedArrayPlansRetainResults(t *testing.T) {
	leaf := testRecordDecoder(testCons(testKey("x"), testGenericStep(kindInt), testNil(), gopurs_runtime.Value{}))
	fallback := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value { panic("array fallback") })
	array := TypedArray(testErrorSupport, TypedArray(testErrorSupport, leaf, fallback), fallback)
	var retained []gopurs_runtime.Value
	for pass := 0; pass < 3; pass++ {
		raw := []any{[]any{map[string]any{"x": float64(pass + 1)}}, []any{}}
		kind := kindOf(array)
		result, _, ok := kind.decodeValue(newErrorPlan(testErrorSupport), raw)
		if !ok {
			t.Fatal("raw nested array failed")
		}
		retained = append(retained, result)
		// The input view is borrowed only during decode; decoded arrays and
		// records must own their output storage.
		raw[0].([]any)[0].(map[string]any)["x"] = float64(999)
	}
	runtime.GC()
	for i, result := range retained {
		row := gopurs_runtime.ArrayAccess(gopurs_runtime.ArrayAccess(result, 0), 0)
		if gopurs_runtime.RecordGet(row, "x").IntVal != int64(i+1) || gopurs_runtime.ArrayLength(gopurs_runtime.ArrayAccess(result, 1)) != 0 {
			t.Fatal("retained nested output was changed")
		}
	}
	boxed := gopurs_runtime.Array([]gopurs_runtime.Value{
		gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Any(map[string]any{"x": gopurs_runtime.Int(9)})}),
	})
	result := DecodeFieldFast(array, boxed)
	if !gopurs_runtime.RecordGet(result, "ok").BoolVal() {
		t.Fatal("boxed nested array failed")
	}
}

func TestRecordTagUnknownTailAndEmptyObject(t *testing.T) {
	calls := 0
	custom := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
		calls++
		return gopurs_runtime.Apply(testRight, value)
	})
	decoder := TypedRecord(testErrorSupport, custom, custom)
	if kindOf(decoder) != nil {
		t.Fatal("an arbitrary row method acquired a plan")
	}
	DecodeFieldFast(decoder, gopurs_runtime.Int(5))
	if calls != 1 {
		t.Fatal("unknown row method was bypassed")
	}
	empty := testRecordDecoder(testNil())
	for _, input := range []gopurs_runtime.Value{
		gopurs_runtime.Any(map[string]any{}), gopurs_runtime.RecordDict0(),
	} {
		result := DecodeFieldFast(empty, input)
		if !gopurs_runtime.RecordGet(result, "ok").BoolVal() || len(gopurs_runtime.RecordToMap(gopurs_runtime.RecordGet(result, "value"))) != 0 {
			t.Fatal("empty object failed")
		}
	}
	for _, input := range []gopurs_runtime.Value{gopurs_runtime.Any(nil), gopurs_runtime.Int(1), gopurs_runtime.Array(nil)} {
		result := DecodeFieldFast(empty, input)
		if gopurs_runtime.RecordGet(result, "ok").BoolVal() || gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(result, "value"), "message").StrVal() != "Object" {
			t.Fatal("empty record decoder accepted a non-object")
		}
	}
}

func TestDirectRecordSizesAndDuplicateInsertion(t *testing.T) {
	for count := 0; count <= 7; count++ {
		plan := testNil()
		input := map[string]any{}
		for i := count - 1; i >= 0; i-- {
			name := fmt.Sprint(i)
			input[name] = float64(i)
			plan = testCons(testKey(name), testGenericStep(kindInt), plan, gopurs_runtime.Value{})
		}
		result := gopurs_runtime.Apply2(plan, gopurs_runtime.Any(input), gopurs_runtime.Value{})
		record := gopurs_runtime.RecordToMap(gopurs_runtime.RecordGet(result, "value"))
		if len(record) != count {
			t.Fatalf("size %d: %d fields", count, len(record))
		}
		for i := 0; i < count; i++ {
			if record[fmt.Sprint(i)].IntVal != int64(i) {
				t.Fatal("wrong field value")
			}
		}
	}
	// Distinct lookup labels collide only during insertion. All callbacks
	// still run and the earlier field must overwrite the later field.
	key := func(lookup string) gopurs_runtime.Value {
		calls := 0
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			calls++
			if calls == 1 {
				return gopurs_runtime.Str(lookup)
			}
			return gopurs_runtime.Str("same")
		})
	}
	plan := testCons(key("a"), testGenericStep(kindInt),
		testCons(key("b"), testGenericStep(kindInt), testNil(), gopurs_runtime.Value{}), gopurs_runtime.Value{})
	result := gopurs_runtime.Apply2(plan, gopurs_runtime.Any(map[string]any{"a": float64(1), "b": float64(2)}), gopurs_runtime.Value{})
	record := gopurs_runtime.RecordToMap(gopurs_runtime.RecordGet(result, "value"))
	if len(record) != 1 || record["same"].IntVal != 1 {
		t.Fatal("duplicate insertion precedence changed")
	}
}

func TestIdentityObjectShortcutOwnsContainer(t *testing.T) {
	identity := TypedJson(testErrorSupport, testRight)
	fallbackCalls := 0
	fallback := gopurs_runtime.Func(func(json gopurs_runtime.Value) gopurs_runtime.Value {
		fallbackCalls++
		return gopurs_runtime.Apply(testLeft, gopurs_runtime.Str("fallback"))
	})
	decoder := TypedObject(identity, fallback, testRight)
	for _, boxed := range []bool{false, true} {
		input := map[string]any{"n": float64(7), "array": []any{"kept", nil}}
		if boxed {
			input["n"] = gopurs_runtime.Int(7)
			input["array"] = gopurs_runtime.Box(input["array"])
		}
		result := gopurs_runtime.Apply(decoder, gopurs_runtime.Any(input))
		if !gopurs_runtime.RecordGet(result, "ok").BoolVal() {
			t.Fatal("identity object did not use the shortcut")
		}
		output := gopurs_runtime.UnboxObject(gopurs_runtime.RecordGet(result, "value"))
		input["n"] = float64(99)
		output["new"] = true
		number, ok := domNumber(output["n"])
		if _, shared := input["new"]; shared || !ok || number != 7 {
			t.Fatal("input and output object storage alias")
		}
		array := gopurs_runtime.Box(output["array"])
		if gopurs_runtime.ArrayAccess(array, 0).StrVal() != "kept" || !domNull(gopurs_runtime.ArrayAccess(array, 1)) {
			t.Fatal("identity object changed nested JSON")
		}
	}
	for _, input := range []gopurs_runtime.Value{gopurs_runtime.Any(nil), gopurs_runtime.Int(1), gopurs_runtime.Array(nil)} {
		gopurs_runtime.Apply(decoder, input)
	}
	// Both arbitrary and non-identity tagged element decoders retain traversal.
	for _, element := range []gopurs_runtime.Value{testRight, testGenericStep(kindInt)} {
		gopurs_runtime.Apply(TypedObject(element, fallback, testRight), gopurs_runtime.Any(map[string]any{}))
	}
	if fallbackCalls != 5 {
		t.Fatalf("fallback calls: %d", fallbackCalls)
	}
}

func TestBorrowObjectIdentityAndFallback(t *testing.T) {
	identity := TypedJson(testErrorSupport, testRight)
	calls := 0
	fallback := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value { calls++; return value })
	decoder := TypedObject(identity, fallback, testRight)
	input := gopurs_runtime.Any(map[string]any{"x": float64(7)})
	result := BorrowObject(decoder, input)
	if gopurs_runtime.RecordGet(result, "value").UnsafePtr != input.UnsafePtr {
		t.Fatal("identity borrowed result lost source")
	}
	if calls != 0 {
		t.Fatal("identity fallback used")
	}
	bad := gopurs_runtime.Int(7)
	if BorrowObject(decoder, bad) != bad || calls != 1 {
		t.Fatal("invalid input fallback changed")
	}
	if BorrowObject(fallback, input) != input || calls != 2 {
		t.Fatal("custom fallback changed")
	}
}

func TestCompiledSchemaGuardOwnershipAndCalls(t *testing.T) {
	original := TypedInt(testErrorSupport, testRight)
	calls := 0
	worker := func(plan *recordDecodePlan, kind *fieldKind, raw any) argonautSchemaResult {
		calls++
		value, err, ok := kind.decodeValue(plan, raw)
		return argonautSchemaResult{value, err, ok}
	}
	accepts := func(kind *fieldKind) bool { return kind != nil && kind.tag == kindInt }
	compiled := argonautCompileSchema(original, worker, accepts)
	for _, input := range []gopurs_runtime.Value{gopurs_runtime.Int(7), gopurs_runtime.Float(2.5), gopurs_runtime.Str("bad")} {
		got := gopurs_runtime.Apply(compiled, input)
		expected := DecodeFieldFast(original, input)
		if gopurs_runtime.RecordGet(got, "ok").BoolVal() != gopurs_runtime.RecordGet(expected, "ok").BoolVal() {
			t.Fatal("success status changed")
		}
		if gopurs_runtime.RecordGet(got, "ok").BoolVal() {
			if gopurs_runtime.RecordGet(got, "value").IntVal != 7 {
				t.Fatal("decoded value changed")
			}
		} else if gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(got, "value"), "message").StrVal() != gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(expected, "value"), "message").StrVal() {
			t.Fatal("error changed")
		}
	}
	if calls != 3 {
		t.Fatalf("calls: %d", calls)
	}
	if tag, _ := gopurs_runtime.FunctionData[*typedKind](original); tag.worker != nil {
		t.Fatal("mutated original decoder metadata")
	}
	unknown := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value { return value })
	if argonautCompileSchema(unknown, worker, accepts) != unknown {
		t.Fatal("custom fallback replaced")
	}
	wrong := TypedString(testErrorSupport, testRight)
	if argonautCompileSchema(wrong, worker, accepts) != wrong {
		t.Fatal("mismatching tag accepted")
	}
}
