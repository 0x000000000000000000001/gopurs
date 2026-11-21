package gopurs_runtime

import "testing"

func TestSTReferencesPreserveAliasesAndReadSnapshots(t *testing.T) {
	initial := StrWithCodeUnits("😀", 2)
	legacyValue := any(initial)
	for _, ref := range []Value{NewSTRef(initial), Any(&legacyValue)} {
		alias := ref
		snapshot := ReadSTRef(alias)
		if got := WriteSTRef(ref, Int(42)); got.IntVal != 42 {
			t.Fatal("write result")
		}
		if ReadSTRef(alias).IntVal != 42 || snapshot != initial {
			t.Fatal("reference alias or immutable read snapshot lost")
		}
		WriteSTRef(alias, initial)
		if ReadSTRef(ref) != initial {
			t.Fatal("string metadata lost")
		}
	}
}

func TestSTValueCellWritesDoNotAllocate(t *testing.T) {
	ref := NewSTRef(Int(0))
	allocations := testing.AllocsPerRun(10, func() {
		for i := int64(0); i < 10000; i++ {
			WriteSTRef(ref, Int(i))
		}
	})
	if allocations != 0 || ReadSTRef(ref).IntVal != 9999 {
		t.Fatalf("writes allocate %v or lose state", allocations)
	}
}

// These legacy operations preserve the published gopurs-st FFI bodies from
// revision b915b18e254f57994f88827a4172f6022ff47166. In particular, the type
// assertions must remain *interface{}, not adapt to the new representation.
func legacySTNew(value Value) Value {
	cell := any(value)
	return Box(&cell)
}

func legacySTRead(ref Value) Value {
	return AnyToValue(*ref.PtrVal().(*interface{}))
}

func legacySTWrite(ref, value Value) Value {
	*ref.PtrVal().(*interface{}) = value
	return value
}

func legacySTModify(f, ref Value) Value {
	ptr := ref.PtrVal().(*interface{})
	updated := Apply(f, AnyToValue(*ptr))
	*ptr = RecordGet(updated, "state")
	return RecordGet(updated, "value")
}

func currentSTModify(f, ref Value) Value {
	updated := Apply(f, ReadSTRef(ref))
	WriteSTRef(ref, RecordGet(updated, "state"))
	return RecordGet(updated, "value")
}

func TestMixedSTReferencePrimitives(t *testing.T) {
	constructors := map[string]func(Value) Value{
		"compiler": NewSTRef,
		"legacy":   legacySTNew,
		"ffi":      func(value Value) Value { return NewSTRef(value) },
	}
	readers := map[string]func(Value) Value{
		"compiler": ReadSTRef,
		"legacy":   legacySTRead,
		"ffi":      func(ref Value) Value { return ReadSTRef(ref) },
	}
	writers := map[string]func(Value, Value) Value{
		"compiler": WriteSTRef,
		"legacy":   legacySTWrite,
		"ffi":      func(ref, value Value) Value { return WriteSTRef(ref, value) },
	}
	modifiers := map[string]func(Value, Value) Value{
		"legacy": legacySTModify,
		"ffi":    currentSTModify,
	}
	initial := StrWithCodeUnits("😀", 2)
	for cn, create := range constructors {
		for rn, read := range readers {
			for wn, write := range writers {
				for mn, modify := range modifiers {
					t.Run(cn+"/"+rn+"/"+wn+"/"+mn, func(t *testing.T) {
						ref := create(initial)
						alias := ref
						snapshot := read(alias)
						if snapshot != initial {
							t.Fatal("initial value or string metadata changed")
						}
						if got := write(ref, Int(42)); got.IntVal != 42 {
							t.Fatal("write result")
						}
						if read(alias).IntVal != 42 || snapshot != initial {
							t.Fatal("alias or immutable read snapshot changed")
						}
						seen := 0
						callback := Func(func(value Value) Value {
							seen++
							if value.IntVal != 42 {
								t.Fatal("modify callback input")
							}
							return RecordDict2("state", "value", Int(43), initial)
						})
						if got := modify(callback, alias); got != initial || seen != 1 || read(ref).IntVal != 43 {
							t.Fatal("modify callback count, result or state")
						}
						WriteSTRef(ref, Int(44))
						legacySTWrite(alias, Int(45))
						WriteSTRef(alias, initial)
						if read(ref) != initial || snapshot != initial {
							t.Fatal("mixed writes or metadata")
						}
					})
				}
			}
		}
	}
}

func TestSTReferencesCanStoreReferences(t *testing.T) {
	inner := NewSTRef(Int(1))
	outer := NewSTRef(inner)
	for _, read := range []func(Value) Value{ReadSTRef, legacySTRead} {
		stored := read(outer)
		if stored != inner {
			t.Fatal("a reference stored as a value was unwrapped")
		}
		WriteSTRef(stored, Int(2))
		if ReadSTRef(inner).IntVal != 2 {
			t.Fatal("stored reference lost its identity")
		}
	}
}
