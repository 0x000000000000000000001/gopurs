package purescript

import (
	"encoding/json"
	"reflect"
	"testing"
	"unsafe"

	rt "gopurs/output/gopurs_runtime"
)

func TestTextOwnedPublicJSONAndForeignObject(t *testing.T) {
	tag, ok := rt.FunctionData[*typedKind](Get_Main_decodeJSONFields())
	if !ok || tag == nil || tag.textWorker == nil {
		t.Fatal("owned fixture did not install the text worker")
	}
	decode := func(decoder rt.Value, text string) rt.Value {
		result := rt.Apply(decoder, rt.Str(text))
		var out rt.Value
		rt.Apply3(Get_Data_Either_either(), rt.Func(func(err rt.Value) rt.Value { t.Fatal("decode failed", err); return rt.Value{} }), rt.Func(func(value rt.Value) rt.Value { out = value; return rt.Value{} }), result)
		return out
	}
	text := `{"json":{"inner":["owned-é🙂"]},"name":"kept","obj":{"keep":"owned-object"}}`
	bytes := []byte(text)
	borrowed := unsafe.String(unsafe.SliceData(bytes), len(bytes))
	first := decode(Get_Main_decodeOwnedText(), borrowed)
	second := decode(Get_Main_decodeOwnedText(), text)
	firstFast := decode(Get_Main_decodeJSONFieldsText(), borrowed)
	secondFast := decode(Get_Main_decodeJSONFieldsText(), text)
	for i := range bytes {
		bytes[i] = 'x'
	}
	jsonText := func(value rt.Value) string {
		return rt.Apply(Get_Data_Argonaut_Core_stringify(), rt.RecordGet(value, "json")).StrVal()
	}
	var saved, expected any
	if err := json.Unmarshal([]byte(jsonText(firstFast)), &saved); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(jsonText(secondFast)), &expected); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(saved, expected) {
		t.Fatal("public Json retained the input buffer")
	}
	firstObject := rt.RecordGet(first, "obj")
	secondObject := rt.RecordGet(second, "obj")
	if firstObject.Type != rt.TypeAny || secondObject.Type != rt.TypeAny {
		t.Fatal("public object native representation changed")
	}
	left, ok := firstObject.PtrVal().(map[string]any)
	if !ok {
		t.Fatal("public Foreign.Object must have its owned map")
	}
	right, ok := secondObject.PtrVal().(map[string]any)
	if !ok {
		t.Fatal("second public object")
	}
	objectText := func(value rt.Value) string {
		return rt.Apply(Get_Data_Argonaut_Core_stringify(), rt.Apply(Get_Data_Argonaut_Core_fromObject(), value)).StrVal()
	}
	if objectText(firstObject) != objectText(secondObject) {
		t.Fatal("public object retained input")
	}
	left["new"] = float64(9)
	delete(left, "keep")
	if _, present := right["new"]; present {
		t.Fatal("public objects share map storage")
	}
	if _, present := right["keep"]; !present {
		t.Fatal("deletion escaped to another decode")
	}
	if jsonText(first) != jsonText(second) || jsonText(firstFast) != jsonText(secondFast) {
		t.Fatal("mutating the public object changed its Json sibling")
	}
}
