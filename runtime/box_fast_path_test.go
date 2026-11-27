package gopurs_runtime

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"testing"
)

func TestBoxValueIdentityPreservesPayloadAndMetadata(t *testing.T) {
	callback := Func(func(value Value) Value { return value })
	withData := WithFunctionData(callback, "callback metadata")
	values := []Value{
		{}, Int(1 << 40), Float(math.Float64frombits(0x7ff8000000000042)),
		StrWithCodeUnits("a😀", 3), Array([]Value{Int(1234)}),
		RecordDict2("a", "b", Int(1), Int(2)), callback, withData,
	}
	for _, value := range values {
		if got := Box(value); got != value {
			t.Fatalf("typed Value identity: got %+v, want %+v", got, value)
		}
		if got := Box[any](value); got != value {
			t.Fatalf("interface Value identity: got %+v, want %+v", got, value)
		}
	}
	if metadata, ok := FunctionData[string](Box[any](withData)); !ok || metadata != "callback metadata" {
		t.Fatal("Box discarded callback metadata")
	}
}

func TestBoxScalarAndForeignRepresentations(t *testing.T) {
	for _, sample := range []struct {
		input    any
		expected Value
	}{
		{int32(-2147483648), Int(-2147483648)}, {int64(1 << 40), Int(1 << 40)},
		{int(1234), Int(1234)}, {false, Bool(false)}, {true, Bool(true)},
		{math.Float64frombits(0x8000000000000000), Float(math.Float64frombits(0x8000000000000000))},
		{math.Float64frombits(0x7ff8000000000042), Float(math.Float64frombits(0x7ff8000000000042))},
	} {
		if got := Box(sample.input); got != sample.expected {
			t.Fatalf("%T: got %+v, want %+v", sample.input, got, sample.expected)
		}
	}
	if got := Box(int64(1234)); got != Int(1234) {
		t.Fatal("native int64")
	}
	if got := Box(int(1234)); got != Int(1234) {
		t.Fatal("native int")
	}
	if got := Box(int32(1234)); got != Int(1234) {
		t.Fatal("native int32")
	}
	for _, text := range []string{"", "a😀", "\xed\xa0\x80", "\xff"} {
		if got := Box(text); got.Type != TypeString || got.StrVal() != text {
			t.Fatalf("string %q changed", text)
		}
		if got := Box[any](text); got.Type != TypeString || got.StrVal() != text {
			t.Fatalf("interface string %q changed", text)
		}
	}
	type namedInt int64
	type namedValue Value
	var pointer *int
	var mapping map[string]int
	var channel chan int
	for _, input := range []any{nil, pointer, mapping, channel, namedInt(1234), namedValue(Int(1234)), struct{ N int }{1234}} {
		got := Box(input)
		if got.Type != TypeAny || !reflect.DeepEqual(got.PtrVal(), input) {
			t.Fatalf("foreign %T representation changed", input)
		}
	}
	if got := Box(pointer); got.Type != TypeAny || got.PtrVal().(*int) != nil {
		t.Fatal("typed nil pointer changed")
	}
	if got := Box(namedInt(1234)); got.Type != TypeAny || got.PtrVal().(namedInt) != 1234 {
		t.Fatal("named integer became primitive")
	}
}

func TestBoxArraysKeepRecursiveConversionAndNativeAliasing(t *testing.T) {
	values := []Value{Int(1234), Str("hello")}
	boxed := Box(values)
	values[0] = Int(5678)
	if ArrayAccess(boxed, 0) != Int(5678) {
		t.Fatal("native Value slice was copied")
	}
	nested := Box([]any{int64(1234), []any{true, "x"}, values})
	if ArrayLength(nested) != 3 || ArrayAccess(nested, 0) != Int(1234) || ArrayAccess(ArrayAccess(nested, 1), 0) != Bool(true) {
		t.Fatal("recursive array conversion changed")
	}
	values[1] = Str("changed")
	if ArrayAccess(ArrayAccess(nested, 2), 1).StrVal() != "changed" {
		t.Fatal("nested native slice alias changed")
	}
	for _, input := range []any{[]Value(nil), []any(nil)} {
		if got := Box(input); got.Type != TypeArray || ArrayLength(got) != 0 {
			t.Fatalf("typed nil slice %T changed", input)
		}
	}
}

func TestBoxFunctionAdaptersRemainDeferredAcrossGC(t *testing.T) {
	calls := 0
	mark := func() { calls++ }
	for _, sample := range []struct {
		name     string
		input    any
		args     []Value
		expected Value
	}{
		{"unit", func() { mark() }, nil, Value{}},
		{"bool", func() bool { mark(); return true }, nil, Bool(true)},
		{"int", func() int { mark(); return 1234 }, nil, Int(1234)},
		{"int64", func() int64 { mark(); return 1 << 40 }, nil, Int(1 << 40)},
		{"string", func() string { mark(); return "x" }, nil, Str("x")},
		{"float", func() float64 { mark(); return -1.25 }, nil, Float(-1.25)},
		{"value", func() Value { mark(); return Int(1234) }, nil, Int(1234)},
		{"unaryValue", func(v Value) Value { mark(); return v }, []Value{Int(1234)}, Int(1234)},
		{"unaryAny", func(v any) any { mark(); return v }, []Value{Int(1234)}, Int(1234)},
		{"curried2", func(a any) func(any) any {
			mark()
			return func(b any) any { mark(); return IntAdd(a.(Value).IntVal, b.(Value).IntVal) }
		}, []Value{Int(1200), Int(34)}, Int(1234)},
		{"curried3", func(a any) func(any) func(any) any {
			mark()
			return func(b any) func(any) any {
				mark()
				return func(c any) any { mark(); return IntAdd(IntAdd(a.(Value).IntVal, b.(Value).IntVal), c.(Value).IntVal) }
			}
		}, []Value{Int(1200), Int(30), Int(4)}, Int(1234)},
	} {
		t.Run(sample.name, func(t *testing.T) {
			before := calls
			function := Box(sample.input)
			if calls != before {
				t.Fatal("boxing invoked callback")
			}
			runtime.GC()
			args := sample.args
			if len(args) == 0 {
				args = []Value{{}}
			}
			result := function
			for _, arg := range args {
				result = Apply(result, arg)
			}
			if result != sample.expected {
				t.Fatalf("got %+v want %+v", result, sample.expected)
			}
			if calls-before != len(args) {
				t.Fatalf("callback calls got %d want %d", calls-before, len(args))
			}
		})
	}
}

func TestBoxTypedNilFunctionsStillFailOnlyWhenApplied(t *testing.T) {
	for _, input := range []any{(func())(nil), (func() bool)(nil), (func() int)(nil), (func() int64)(nil),
		(func() string)(nil), (func() float64)(nil), (func() Value)(nil), (func(Value) Value)(nil),
		(func(any) any)(nil), (func(any) func(any) any)(nil), (func(any) func(any) func(any) any)(nil)} {
		t.Run(fmt.Sprintf("%T", input), func(t *testing.T) {
			boxed := Box(input)
			if boxed.Type != TypeFunc {
				t.Fatal("nil function no longer boxed as callable")
			}
			defer func() {
				if recover() == nil {
					t.Fatal("applying nil callback did not panic")
				}
			}()
			Apply(boxed, Value{})
		})
	}
}
