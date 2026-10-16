package gopurs_runtime

import (
	"reflect"
	"testing"
)

func TestRecordFieldReflectionUsesLabelsWithoutChangingOrder(t *testing.T) {
	record := struct {
		P_type     string
		P_value    int64
		PX_d800    bool
		P_go__type string
	}{"tag", 42, true, "other"}
	value := reflect.ValueOf(record)
	labels := []string{"type", "value", "\xed\xa0\x80", "go__type"}
	for index, label := range labels {
		field := value.FieldByName(RecordFieldName(label))
		if !field.IsValid() || !field.CanInterface() || field.Interface() != value.Field(index).Interface() {
			t.Fatalf("reflection/layout mismatch for %q", label)
		}
		decoded, ok := RecordFieldLabel(value.Type().Field(index).Name)
		if !ok || decoded != label {
			t.Fatalf("field lost its label: %q -> %q", label, decoded)
		}
	}
}

func TestRecordFieldNamesPreserveUTF16Identity(t *testing.T) {
	vectors := map[string]string{
		"": "PX_", "_": "P__", "a1": "P_a1", "1": "PX_0031", "type": "P_type",
		"go__type": "P_go__type", "a-b": "PX_0061002d0062", "a_minus_b": "P_a_minus_b",
		"💡": "PX_d83ddca1", "\xed\xa0\x80": "PX_d800", "\xed\xb0\x80": "PX_dc00",
		"\x00": "PX_0000", "é": "PX_00e9", "P_type": "P_P_type",
	}
	for label, name := range vectors {
		if got := RecordFieldName(label); got != name {
			t.Fatalf("%q: %s != %s", label, got, name)
		}
		if got, ok := RecordFieldLabel(name); !ok || got != label {
			t.Fatalf("%s: %q != %q (%t)", name, got, label, ok)
		}
	}
	for _, name := range []string{"value", "V0", "Rc", "P_", "P_a-b", "PX_0061", "PX_D800", "PX_d80", "PX_zzzz"} {
		if _, ok := RecordFieldLabel(name); ok {
			t.Fatalf("noncanonical record field accepted: %q", name)
		}
	}
}
