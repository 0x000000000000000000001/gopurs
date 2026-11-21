package gopurs_runtime

import (
	"runtime"
	"strings"
	"testing"
	"unicode/utf8"
	"unsafe"
)

func referenceCodeUnits(s string) int {
	units := 0
	for len(s) > 0 {
		if len(s) >= 3 && s[0] == 0xed && s[1] >= 0xa0 && s[1] <= 0xbf && s[2]&0xc0 == 0x80 {
			units++
			s = s[3:]
			continue
		}
		r, n := utf8.DecodeRuneInString(s)
		units++
		if r > 0xffff {
			units++
		}
		s = s[n:]
	}
	return units
}
func TestOptionalStringMetadata(t *testing.T) {
	if unsafe.Sizeof(Value{}) != 24 {
		t.Fatal("Value grew")
	}
	inputs := []string{"", "a", "abc", "a😀é", "\xed\xa0\x80", "\xed\xbf\xbf", "\xff\xc0\xaf", "\xe0\x9f\xbf", "\xf0\x8f\xbf\xbf", "\xf4\x90\x80\x80", strings.Repeat("a😀é", 1000)}
	for _, s := range inputs {
		units := referenceCodeUnits(s)
		plain := Str(s)
		if plain.IntVal != int64(len(s)) {
			t.Fatal("ordinary Str gained metadata")
		}
		cached := StrWithCodeUnits(s, units)
		if got, known := CachedStringCodeUnits(cached); !known || got != units {
			t.Fatal("metadata inaccessible")
		}
		if len(s) > 0 {
			if _, known := CachedStringCodeUnits(plain); known {
				t.Fatal("ordinary nonempty string gained metadata")
			}
		}
		for _, value := range []Value{plain, cached, NormalizeStringValue(cached), Box(any(cached)), AnyToValue(any(cached))} {
			if value.StrVal() != s || StringCodeUnitsLength(value) != units {
				t.Fatalf("roundtrip %x", s)
			}
		}
		if cached != NormalizeStringValue(cached) {
			t.Fatal("normalization lost metadata")
		}
		if cached != Box(any(cached)) {
			t.Fatal("polymorphic FFI lost metadata")
		}
		if plain != NormalizeStringValue(plain) {
			t.Fatal("normalization changed plain representation")
		}
	}
	empty := StrWithCodeUnits("", 0)
	if empty.IntVal != 0 || empty.UnsafePtr != nil || StringCodeUnitsLength(empty) != 0 {
		t.Fatal("empty layout")
	}
	for _, value := range []Value{Value{}, Int(42), Bool(true)} {
		if _, known := CachedStringCodeUnits(value); known {
			t.Fatal("cached length on non-string")
		}
		if NormalizeStringValue(value) != Str(value.StrVal()) {
			t.Fatal("non-string normalization changed semantics")
		}
	}
}
func TestStringMetadataBounds(t *testing.T) {
	max := int(stringCodeUnitsMask)
	for _, pair := range [][2]int{{0, 0}, {1, 1}, {max, 0}, {max, max}} {
		packed, ok := packStringCodeUnits(pair[0], pair[1])
		if !ok || packed >= 0 || int((-1-packed)>>31) != pair[0] || int((-1-packed)&stringCodeUnitsMask) != pair[1] {
			t.Fatal(pair, packed, ok)
		}
	}
	for _, pair := range [][2]int{{-1, 0}, {1, -1}, {1, 2}, {max + 1, 0}, {max + 1, max + 1}} {
		if _, ok := packStringCodeUnits(pair[0], pair[1]); ok {
			t.Fatal("must fall back", pair)
		}
	}
	for _, units := range []int{-1, 4, max + 1} {
		plain := Str("abc")
		value := StrWithCodeUnits("abc", units)
		if value != plain || StringCodeUnitsLength(value) != 3 {
			t.Fatal("fallback", units)
		}
	}
}

var stringMetadataSink Value

func TestStringBoxingRemainsAllocationFree(t *testing.T) {
	text := strings.Repeat("a😀é", 1000)
	units := referenceCodeUnits(text)
	for name, box := range map[string]func() Value{
		"ordinary":     func() Value { return Str(text) },
		"known length": func() Value { return StrWithCodeUnits(text, units) },
	} {
		if allocations := testing.AllocsPerRun(100, func() { stringMetadataSink = box() }); allocations != 0 {
			t.Fatalf("%s boxing allocated: %v", name, allocations)
		}
	}
}

//go:noinline
func retainedStringWithMetadata() Value {
	text := strings.Repeat("padding", 1024) + "é😀\xed\xa0\x80\x00"
	return StrWithCodeUnits(text[len("padding")*1024:], 5)
}

func TestStringMetadataRetainsBytes(t *testing.T) {
	value := retainedStringWithMetadata()
	runtime.GC()
	if value.StrVal() != "é😀\xed\xa0\x80\x00" || StringCodeUnitsLength(value) != 5 {
		t.Fatal("metadata string lost its backing storage")
	}
}
