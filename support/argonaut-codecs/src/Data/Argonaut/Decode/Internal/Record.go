package Record

import (
	"math"

	"gopurs/output/gopurs_runtime"
)

type recordDecodeField struct {
	reflectName gopurs_runtime.Value
	step        gopurs_runtime.Value
	kind        *fieldKind
}

// A decoder tag says nothing about the symbol callback. Evaluate it at lookup
// and again at insertion, in the same order as the ordinary row decoder.
func (field *recordDecodeField) fieldKey() string {
	return gopurs_runtime.Apply(field.reflectName, gopurs_runtime.Value{}).StrVal()
}

// A plan is immutable after construction and shared by every invocation of
// the dictionary. Custom dictionaries carry no plan and keep their own method.
type recordDecodePlan struct {
	fields []recordDecodeField

	isRight    gopurs_runtime.Value
	rightValue gopurs_runtime.Value
	right      gopurs_runtime.Value

	// Error and Maybe constructors, captured from the caller, so the direct
	// path builds exactly the values the generic path builds.
	left         gopurs_runtime.Value
	atKey        gopurs_runtime.Value
	missingValue gopurs_runtime.Value
	typeMismatch gopurs_runtime.Value
	named        gopurs_runtime.Value
	atIndex      gopurs_runtime.Value
	leftOf       gopurs_runtime.Value
	nothing      gopurs_runtime.Value
	just         gopurs_runtime.Value

	// hasCustom reports whether any field needs the generic step, and therefore
	// whether a boxed object value must be materialised for nested records.
	hasCustom bool
}

// captureErrors reads the support record once per plan construction. The
// values are never inspected while a document is decoded.
func (plan *recordDecodePlan) captureErrors(errors gopurs_runtime.Value) {
	plan.atKey = gopurs_runtime.RecordGet(errors, "atKey")
	plan.missingValue = gopurs_runtime.RecordGet(errors, "missingValue")
	plan.typeMismatch = gopurs_runtime.RecordGet(errors, "typeMismatch")
	plan.named = gopurs_runtime.RecordGet(errors, "named")
	plan.atIndex = gopurs_runtime.RecordGet(errors, "atIndex")
	plan.leftOf = gopurs_runtime.RecordGet(errors, "leftOf")
}

// fieldKind is the direct-path classification of one field decoder. Primitive
// and container decoders carry a typed tag (see typedKind); a record step
// carries its own plan. A container whose element decoder is a custom function
// keeps that function so the elements can still bypass the generic traversals.
type fieldKind struct {
	tag    string
	inner  *fieldKind
	plan   *recordDecodePlan
	custom gopurs_runtime.Value
}

const (
	kindJson    = "Json"
	kindInt     = "Int"
	kindNumber  = "Number"
	kindString  = "String"
	kindBoolean = "Boolean"
	kindMaybe   = "Maybe"
	kindArray   = "Array"
	kindRecord  = "Record"
)

// typedKind tags a decoder whose directly decodable shape is known. The error
// plan is captured once, when the decoder is tagged, so a field access never
// reads the support record again.
type typedKind struct {
	kind   *fieldKind
	errors *recordDecodePlan
	worker argonautSchemaWorker
	// Optional text support lives in Decode.Parser, which need not be linked by
	// callers using only DecodeJson. That module checks its concrete worker type.
	textWorker any
}

var SchemaDecoderABI1 = 1

// Workers are emitted from resolved standard dictionaries, not from result
// types alone. Their schema guards run once at decoder construction. Unknown
// field methods and custom container elements retain their existing callbacks.
type argonautSchemaResult struct {
	value gopurs_runtime.Value
	err   gopurs_runtime.Value
	ok    bool
}
type argonautSchemaWorker func(*recordDecodePlan, *fieldKind, any) argonautSchemaResult

func argonautCompileSchema(original gopurs_runtime.Value, worker argonautSchemaWorker, accepts func(*fieldKind) bool) gopurs_runtime.Value {
	tag, ok := gopurs_runtime.FunctionData[*typedKind](original)
	if !ok || tag == nil || !accepts(tag.kind) {
		return original
	}
	compiled := *tag
	compiled.worker = worker
	compiled.textWorker = nil
	return gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(json gopurs_runtime.Value) gopurs_runtime.Value {
		var input any = json
		if json.Type == gopurs_runtime.TypeAny {
			input = json.PtrVal()
		}
		result := worker(tag.errors, tag.kind, input)
		if !result.ok {
			return gopurs_runtime.Apply(tag.errors.left, result.err)
		}
		return gopurs_runtime.Apply(tag.errors.right, result.value)
	}), &compiled)
}

func argonautSchemaCustom(plan *recordDecodePlan, kind *fieldKind, raw any) argonautSchemaResult {
	value, err, ok := kind.decodeValue(plan, raw)
	return argonautSchemaResult{value, err, ok}
}

func argonautSchemaElement(plan *recordDecodePlan, kind *fieldKind, raw any) argonautSchemaResult {
	value, err, ok := plan.decodeCustom(kind.custom, raw)
	return argonautSchemaResult{value, err, ok}
}

func argonautSchemaField(plan *recordDecodePlan, index int, object gopurs_runtime.Value, key string) argonautSchemaResult {
	result := gopurs_runtime.Apply2(plan.fields[index].step, gopurs_runtime.Str(key), object)
	if !gopurs_runtime.Apply(plan.isRight, result).BoolVal() {
		return argonautSchemaResult{err: gopurs_runtime.Apply(plan.leftOf, result)}
	}
	return argonautSchemaResult{value: gopurs_runtime.Apply(plan.rightValue, result), ok: true}
}

// kindOf classifies a decoder value: a shared record plan, a typed tag, or
// nothing (an arbitrary custom decoder).
func kindOf(decoder gopurs_runtime.Value) *fieldKind {
	if plan, ok := gopurs_runtime.FunctionData[*recordDecodePlan](decoder); ok && plan != nil {
		return &fieldKind{tag: kindRecord, plan: plan}
	}
	if tag, ok := gopurs_runtime.FunctionData[*typedKind](decoder); ok && tag != nil {
		return tag.kind
	}
	return nil
}

// newErrorPlan captures every constructor a direct decode can need.
func newErrorPlan(support gopurs_runtime.Value) *recordDecodePlan {
	plan := &recordDecodePlan{}
	plan.captureErrors(support)
	plan.nothing = gopurs_runtime.RecordGet(support, "nothing")
	plan.just = gopurs_runtime.RecordGet(support, "just")
	plan.isRight = gopurs_runtime.RecordGet(support, "isRight")
	plan.rightValue = gopurs_runtime.RecordGet(support, "rightValue")
	plan.left = gopurs_runtime.RecordGet(support, "left")
	plan.right = gopurs_runtime.RecordGet(support, "right")
	return plan
}

// DecodeFieldFast decodes one field with a tagged decoder and falls back to the
// ordinary decoder call otherwise. Errors are returned without the AtKey
// wrapper; the PureScript caller adds it exactly like the generic path does.
func DecodeFieldFast(decoder, json gopurs_runtime.Value) gopurs_runtime.Value {
	if plan, ok := gopurs_runtime.FunctionData[*recordDecodePlan](decoder); ok && plan != nil {
		return plan.decodeDirect(json, gopurs_runtime.Value{})
	}
	tagData, ok := gopurs_runtime.FunctionData[*typedKind](decoder)
	if !ok || tagData == nil {
		return gopurs_runtime.Apply(decoder, json)
	}
	var value, err gopurs_runtime.Value
	if tagData.worker != nil {
		result := tagData.worker(tagData.errors, tagData.kind, json)
		value, err, ok = result.value, result.err, result.ok
	} else {
		value, err, ok = tagData.kind.decodeValue(tagData.errors, json)
	}
	if !ok {
		return gopurs_runtime.Apply(tagData.errors.left, err)
	}
	return gopurs_runtime.Apply(tagData.errors.right, value)
}

// Fuse a field lookup and a recognised decoder without allocating lookup's
// Just or a temporary Either. Build only the public accessor's final result.
// Untagged decoders use the original accessor, so custom decoding and callback
// behavior (including when a key is missing or null) remain unchanged.
func getFieldDirect(fallback, decoder, object, key gopurs_runtime.Value, optional, nullable bool) gopurs_runtime.Value {
	tag, known := gopurs_runtime.FunctionData[*typedKind](decoder)
	if !known || tag == nil {
		return gopurs_runtime.Apply3(fallback, decoder, object, key)
	}
	dom, ok := gopurs_runtime.ReadJSONObjectValue(object)
	if !ok {
		return gopurs_runtime.Apply3(fallback, decoder, object, key)
	}
	plan := tag.errors
	label := key.StrVal()
	raw, present := dom.Lookup(label)
	if !present || (nullable && domNull(raw)) {
		if optional {
			return gopurs_runtime.Apply(plan.right, plan.mkNothing())
		}
		return gopurs_runtime.Apply(plan.left, plan.mkAtKey(label, plan.missingValue))
	}
	var value, err gopurs_runtime.Value
	if tag.worker != nil {
		result := tag.worker(plan, tag.kind, raw)
		value, err, ok = result.value, result.err, result.ok
	} else {
		value, err, ok = tag.kind.decodeValue(plan, raw)
	}
	if !ok {
		return gopurs_runtime.Apply(plan.left, plan.mkAtKey(label, err))
	}
	if optional {
		value = plan.mkJust(value)
	}
	return gopurs_runtime.Apply(plan.right, value)
}

func GetFieldImpl(fallback, decoder, object, key gopurs_runtime.Value) gopurs_runtime.Value {
	return getFieldDirect(fallback, decoder, object, key, false, false)
}

func GetFieldOptionalImpl(fallback, decoder, object, key gopurs_runtime.Value) gopurs_runtime.Value {
	return getFieldDirect(fallback, decoder, object, key, true, false)
}

func GetFieldOptionalNullableImpl(fallback, decoder, object, key gopurs_runtime.Value) gopurs_runtime.Value {
	return getFieldDirect(fallback, decoder, object, key, true, true)
}

func detectFieldKind(step gopurs_runtime.Value) *fieldKind {
	return kindOf(step)
}

// ---------------------------------------------------------------------------
// Error helpers (only the failure paths allocate)
// ---------------------------------------------------------------------------

func (plan *recordDecodePlan) mkAtKey(key string, err gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply2(plan.atKey, gopurs_runtime.Str(key), err)
}

func (plan *recordDecodePlan) mkTypeMismatch(message string) gopurs_runtime.Value {
	return gopurs_runtime.Apply(plan.typeMismatch, gopurs_runtime.Str(message))
}

func (plan *recordDecodePlan) mkNamed(name string, err gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply2(plan.named, gopurs_runtime.Str(name), err)
}

func (plan *recordDecodePlan) mkAtIndex(index int, err gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply2(plan.atIndex, gopurs_runtime.Int(int64(index)), err)
}

func (plan *recordDecodePlan) mkNothing() gopurs_runtime.Value {
	return plan.nothing
}

func (plan *recordDecodePlan) mkJust(value gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Apply(plan.just, value)
}

// intFromNumber mirrors Data.Int.FromNumberImpl on this runtime.
func intFromNumber(n float64) (int64, bool) {
	if math.IsNaN(n) || math.IsInf(n, 0) || math.Trunc(n) != n {
		return 0, false
	}
	return int64(int(n)), true
}

// ---------------------------------------------------------------------------
// DOM access
//
// The native parser stores raw Go values (float64, string, bool, []any,
// map[string]any, compact JSONObject, nil). Json values built in PureScript (fromObject,
// Object.fromFoldable, record dictionaries) store boxed gopurs values
// instead. Both must decode identically, so every accessor accepts either.
// ---------------------------------------------------------------------------

func domNumber(raw any) (float64, bool) {
	switch value := raw.(type) {
	case float64:
		return value, true
	case gopurs_runtime.Value:
		switch value.Type {
		case gopurs_runtime.TypeInt:
			return float64(value.IntVal), true
		case gopurs_runtime.TypeFloat:
			return value.FloatVal(), true
		}
	}
	return 0, false
}

func domString(raw any) (string, bool) {
	switch value := raw.(type) {
	case string:
		return value, true
	case gopurs_runtime.Value:
		if value.Type == gopurs_runtime.TypeString {
			return value.StrVal(), true
		}
	}
	return "", false
}

func domBool(raw any) (bool, bool) {
	switch value := raw.(type) {
	case bool:
		return value, true
	case gopurs_runtime.Value:
		if value.Type == gopurs_runtime.TypeBool {
			return value.IntVal != 0, true
		}
	}
	return false, false
}

// domNull mirrors Data.Argonaut.Core.isNull, including boxed values.
func domNull(raw any) bool {
	switch value := raw.(type) {
	case nil:
		return true
	case gopurs_runtime.Value:
		if value.Type == 0 {
			return true
		}
		if value.Type == gopurs_runtime.TypeAny {
			if value.UnsafePtr == nil {
				return true
			}
			return *(*any)(value.UnsafePtr) == nil
		}
	}
	return false
}

// domArray is a borrowed, read-only view. Decode allocates only the output
// slice; native arrays need no intermediate recursively boxed input slice.
type domArray struct {
	raw   []any
	boxed []gopurs_runtime.Value
}

func (array domArray) length() int {
	return len(array.raw) + len(array.boxed)
}

func (array domArray) at(index int) any {
	if array.raw != nil {
		return array.raw[index]
	}
	return array.boxed[index]
}

func domArrayValues(raw any) (domArray, bool) {
	switch value := raw.(type) {
	case []any:
		return domArray{raw: value}, true
	case gopurs_runtime.Value:
		if value.Type == gopurs_runtime.TypeArray && value.UnsafePtr != nil {
			return domArray{boxed: *(*[]gopurs_runtime.Value)(value.UnsafePtr)}, true
		}
	}
	return domArray{}, false
}

// domObject borrows a read-only view, including compact parsed objects.
func domObject(raw any) (gopurs_runtime.JSONObjectView, bool) {
	return gopurs_runtime.ReadJSONObject(raw)
}

// ---------------------------------------------------------------------------
// Direct decoding from the DOM
// ---------------------------------------------------------------------------

// decodeValue decodes one present JSON value. Errors are returned without the
// AtKey wrapper, exactly like a step before getField adds it.
func (kind *fieldKind) decodeValue(plan *recordDecodePlan, raw any) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	switch kind.tag {
	case kindJson:
		return gopurs_runtime.Box(raw), gopurs_runtime.Value{}, true
	case kindInt:
		n, ok := domNumber(raw)
		if !ok {
			return gopurs_runtime.Value{}, plan.mkTypeMismatch("Number"), false
		}
		value, ok := intFromNumber(n)
		if !ok {
			return gopurs_runtime.Value{}, plan.mkTypeMismatch("Integer"), false
		}
		return gopurs_runtime.Int(value), gopurs_runtime.Value{}, true
	case kindNumber:
		n, ok := domNumber(raw)
		if !ok {
			return gopurs_runtime.Value{}, plan.mkTypeMismatch("Number"), false
		}
		return gopurs_runtime.Float(n), gopurs_runtime.Value{}, true
	case kindString:
		text, ok := domString(raw)
		if !ok {
			return gopurs_runtime.Value{}, plan.mkTypeMismatch("String"), false
		}
		return gopurs_runtime.Str(text), gopurs_runtime.Value{}, true
	case kindBoolean:
		value, ok := domBool(raw)
		if !ok {
			return gopurs_runtime.Value{}, plan.mkTypeMismatch("Boolean"), false
		}
		return gopurs_runtime.Bool(value), gopurs_runtime.Value{}, true
	case kindMaybe:
		if domNull(raw) {
			return plan.mkNothing(), gopurs_runtime.Value{}, true
		}
		if kind.inner != nil {
			value, err, ok := kind.inner.decodeValue(plan, raw)
			if !ok {
				return gopurs_runtime.Value{}, err, false
			}
			return plan.mkJust(value), gopurs_runtime.Value{}, true
		}
		value, err, ok := plan.decodeCustom(kind.custom, raw)
		if !ok {
			return gopurs_runtime.Value{}, err, false
		}
		return plan.mkJust(value), gopurs_runtime.Value{}, true
	case kindArray:
		items, ok := domArrayValues(raw)
		if !ok {
			return gopurs_runtime.Value{}, plan.mkTypeMismatch("Array"), false
		}
		values := make([]gopurs_runtime.Value, items.length())
		for index := range values {
			item := items.at(index)
			var value, err gopurs_runtime.Value
			var ok bool
			if kind.inner != nil {
				value, err, ok = kind.inner.decodeValue(plan, item)
			} else {
				value, err, ok = plan.decodeCustom(kind.custom, item)
			}
			if !ok {
				return gopurs_runtime.Value{}, plan.mkNamed("Array", plan.mkAtIndex(index, err)), false
			}
			values[index] = value
		}
		return gopurs_runtime.Array(values), gopurs_runtime.Value{}, true
	case kindRecord:
		object := gopurs_runtime.Value{}
		if kind.plan.hasCustom {
			object = gopurs_runtime.Box(raw)
		}
		record, err, ok := kind.plan.decodeAny(object, raw)
		if !ok {
			return gopurs_runtime.Value{}, err, false
		}
		return record, gopurs_runtime.Value{}, true
	}
	return gopurs_runtime.Value{}, plan.mkTypeMismatch("Unknown"), false
}

// decodeField decodes one record field, including absence and null handling:
// a missing Maybe field is Nothing, everything else reports AtKey MissingValue.
func (kind *fieldKind) decodeField(plan *recordDecodePlan, dom gopurs_runtime.JSONObjectView, key string) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	raw, present := dom.Lookup(key)
	if !present {
		if kind.tag == kindMaybe {
			return plan.mkNothing(), gopurs_runtime.Value{}, true
		}
		return gopurs_runtime.Value{}, plan.mkAtKey(key, plan.missingValue), false
	}
	value, err, ok := kind.decodeValue(plan, raw)
	if !ok {
		return gopurs_runtime.Value{}, plan.mkAtKey(key, err), false
	}
	return value, gopurs_runtime.Value{}, true
}

// decodeCustom applies an untagged element decoder to one DOM value. The
// result keeps the exact Either envelope the generic path produces.
func (plan *recordDecodePlan) decodeCustom(custom gopurs_runtime.Value, raw any) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	result := gopurs_runtime.Apply(custom, gopurs_runtime.Box(raw))
	if !gopurs_runtime.Apply(plan.isRight, result).BoolVal() {
		return gopurs_runtime.Value{}, gopurs_runtime.Apply(plan.leftOf, result), false
	}
	return gopurs_runtime.Apply(plan.rightValue, result), gopurs_runtime.Value{}, true
}

// decodeAny decodes a record from a DOM object. Fields evaluate in plan order
// and the first error wins. Every field reevaluates its symbol during insertion,
// exactly like the generic path and its contract.
func (plan *recordDecodePlan) decodeAny(object gopurs_runtime.Value, raw any) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	dom, ok := domObject(raw)
	if !ok {
		return gopurs_runtime.Value{}, plan.mkTypeMismatch("Object"), false
	}
	count := len(plan.fields)
	if count == 0 {
		return gopurs_runtime.RecordDict0(), gopurs_runtime.Value{}, true
	}
	if count <= 5 {
		var decoded [5]gopurs_runtime.Value
		if err, ok := plan.decodeFields(object, dom, decoded[:count]); !ok {
			return gopurs_runtime.Value{}, err, false
		}
		return plan.recordSmall(decoded[:count]), gopurs_runtime.Value{}, true
	}
	decoded := make([]gopurs_runtime.Value, count)
	if err, ok := plan.decodeFields(object, dom, decoded); !ok {
		return gopurs_runtime.Value{}, err, false
	}
	return plan.record(decoded), gopurs_runtime.Value{}, true
}

// decoded is call-local scratch, never retained by the field machinery.
func (plan *recordDecodePlan) decodeFields(object gopurs_runtime.Value, dom gopurs_runtime.JSONObjectView, decoded []gopurs_runtime.Value) (gopurs_runtime.Value, bool) {
	count := len(plan.fields)
	for index := range plan.fields {
		field := &plan.fields[index]
		key := field.fieldKey()
		if field.kind != nil {
			value, err, ok := field.kind.decodeField(plan, dom, key)
			if !ok {
				return err, false
			}
			decoded[count-1-index] = value
			continue
		}
		result := gopurs_runtime.Apply2(field.step, gopurs_runtime.Str(key), object)
		if !gopurs_runtime.Apply(plan.isRight, result).BoolVal() {
			// A generic step returns a complete Either. Extract the error so
			// the caller wraps it exactly once, like a direct kind does.
			return gopurs_runtime.Apply(plan.leftOf, result), false
		}
		decoded[count-1-index] = gopurs_runtime.Apply(plan.rightValue, result)
	}
	return gopurs_runtime.Value{}, true
}

// Small records use stack scratch and one final compact allocation. Symbol
// callbacks still run in reverse order, and duplicate labels retain the first
// insertion slot and the head field's value, just like Record.insert.
func (plan *recordDecodePlan) recordSmall(decoded []gopurs_runtime.Value) gopurs_runtime.Value {
	var keys [5]string
	var values [5]gopurs_runtime.Value
	count := len(plan.fields)
	size := 0
	for index := count - 1; index >= 0; index-- {
		key := plan.fields[index].fieldKey()
		slot := 0
		for slot < size && keys[slot] != key {
			slot++
		}
		if slot == size {
			keys[size] = key
			size++
		}
		values[slot] = decoded[count-1-index]
	}
	switch size {
	case 0:
		return gopurs_runtime.RecordDict0()
	case 1:
		return gopurs_runtime.RecordDict1(keys[0], values[0])
	case 2:
		return gopurs_runtime.RecordDict2(keys[0], keys[1], values[0], values[1])
	case 3:
		return gopurs_runtime.RecordDict3(keys[0], keys[1], keys[2], values[0], values[1], values[2])
	case 4:
		return gopurs_runtime.RecordDict4(keys[0], keys[1], keys[2], keys[3], values[0], values[1], values[2], values[3])
	default:
		return gopurs_runtime.RecordDict5(keys[0], keys[1], keys[2], keys[3], keys[4], values[0], values[1], values[2], values[3], values[4])
	}
}

// record mirrors Record.insert: walk the fields backwards, reuse the first
// position of a duplicated label, and publish one final dictionary record.
func (plan *recordDecodePlan) record(decoded []gopurs_runtime.Value) gopurs_runtime.Value {
	count := len(plan.fields)
	keys := make([]string, 0, count)
	values := decoded[:0]
	for index := count - 1; index >= 0; index-- {
		slot := count - 1 - index
		key := plan.fields[index].fieldKey()
		value := decoded[slot]
		existing := -1
		for i, previous := range keys {
			if previous == key {
				existing = i
				break
			}
		}
		if existing >= 0 {
			values[existing] = value
		} else {
			keys = append(keys, key)
			values = append(values, value)
		}
	}
	return gopurs_runtime.RecordDict(keys, values)
}

// decodeDirect uses the DOM fast path when the input is an object value;
// every other input keeps the generic method.
func (plan *recordDecodePlan) decodeDirect(object, proxy gopurs_runtime.Value) gopurs_runtime.Value {
	if object.Type != gopurs_runtime.TypeAny {
		return plan.decode(object, proxy)
	}
	decoded, err, ok := plan.decodeAny(object, object.PtrVal())
	if !ok {
		return gopurs_runtime.Apply(plan.left, err)
	}
	return gopurs_runtime.Apply(plan.right, decoded)
}

// decode is the original generic method, kept for arbitrary steps and for
// inputs that are not the parser's DOM.
func (plan *recordDecodePlan) decode(object, _ gopurs_runtime.Value) gopurs_runtime.Value {
	count := len(plan.fields)
	if count == 0 {
		return gopurs_runtime.Apply(plan.right, gopurs_runtime.RecordDict0())
	}
	decoded := make([]gopurs_runtime.Value, count)
	for index, field := range plan.fields {
		key := gopurs_runtime.Apply(field.reflectName, gopurs_runtime.Value{})
		result := gopurs_runtime.Apply2(field.step, key, object)
		if !gopurs_runtime.Apply(plan.isRight, result).BoolVal() {
			return result
		}
		decoded[count-1-index] = gopurs_runtime.Apply(plan.rightValue, result)
	}
	return gopurs_runtime.Apply(plan.right, plan.record(decoded))
}

// TagDecoder attaches a typed kind and its captured error plan to a decoder.
// The plan is built once per decoder, never per field access.
func TagDecoder(decoder gopurs_runtime.Value, kind *fieldKind, support gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.WithFunctionData(decoder, &typedKind{kind: kind, errors: newErrorPlan(support)})
}

// scalar/container and the Typed* constructors are the native side of the
// schema tags. They live in this module because the FFI selection keeps only
// the functions declared by their own PureScript module.
func scalar(support, decoder gopurs_runtime.Value, name string) gopurs_runtime.Value {
	return TagDecoder(decoder, &fieldKind{tag: name}, support)
}

func container(support, inner, decoder gopurs_runtime.Value, name string) gopurs_runtime.Value {
	if innerKind := kindOf(inner); innerKind != nil {
		return TagDecoder(decoder, &fieldKind{tag: name, inner: innerKind}, support)
	}
	// A custom element decoder still lets the container itself take the direct
	// path; only the elements go through the custom step.
	return TagDecoder(decoder, &fieldKind{tag: name, custom: inner}, support)
}

func TypedInt(support, decoder gopurs_runtime.Value) gopurs_runtime.Value {
	return scalar(support, decoder, kindInt)
}

func TypedJson(support, decoder gopurs_runtime.Value) gopurs_runtime.Value {
	return scalar(support, decoder, kindJson)
}

func TypedNumber(support, decoder gopurs_runtime.Value) gopurs_runtime.Value {
	return scalar(support, decoder, kindNumber)
}

func TypedString(support, decoder gopurs_runtime.Value) gopurs_runtime.Value {
	return scalar(support, decoder, kindString)
}

func TypedBoolean(support, decoder gopurs_runtime.Value) gopurs_runtime.Value {
	return scalar(support, decoder, kindBoolean)
}

func TypedMaybe(support, inner, decoder gopurs_runtime.Value) gopurs_runtime.Value {
	return container(support, inner, decoder, kindMaybe)
}

func TypedArray(support, inner, decoder gopurs_runtime.Value) gopurs_runtime.Value {
	return container(support, inner, decoder, kindArray)
}

func TypedRecord(support, row, decoder gopurs_runtime.Value) gopurs_runtime.Value {
	plan, known := gopurs_runtime.FunctionData[*recordDecodePlan](row)
	if !known || plan == nil {
		return decoder
	}
	return TagDecoder(decoder, &fieldKind{tag: kindRecord, plan: plan}, support)
}

// A Foreign.Object Json traversal applies only the standard identity decoder.
// Keep a fresh output map and the same boxed entries, but avoid per-entry
// Either/Maybe/record machinery. Arbitrary element decoders retain traversal.
func TypedObject(inner, decoder, right gopurs_runtime.Value) gopurs_runtime.Value {
	kind := kindOf(inner)
	if kind == nil || kind.tag != kindJson {
		return decoder
	}
	return gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(json gopurs_runtime.Value) gopurs_runtime.Value {
		return decodeJsonObjectDirect(decoder, right, json)
	}), &identityObjectDecoder{right: right})
}

type identityObjectDecoder struct{ right gopurs_runtime.Value }

// Only compiler-proven synchronous read-only consumers can call this helper.
// Unknown decoders (including custom element decoders) keep their ordinary
// behavior. Invalid inputs use that same path, preserving exact errors.
func BorrowObject(decoder, json gopurs_runtime.Value) gopurs_runtime.Value {
	if identity, ok := gopurs_runtime.FunctionData[*identityObjectDecoder](decoder); ok && identity != nil {
		if _, object := gopurs_runtime.ReadJSONObjectValue(json); object {
			return gopurs_runtime.Apply(identity.right, json)
		}
	}
	return gopurs_runtime.Apply(decoder, json)
}

func decodeJsonObjectDirect(decoder, right, json gopurs_runtime.Value) gopurs_runtime.Value {
	object, ok := gopurs_runtime.ReadJSONObjectValue(json)
	if !ok {
		return gopurs_runtime.Apply(decoder, json)
	}
	owned := make(map[string]any, object.Length())
	object.Each(func(key string, raw any) {
		owned[key] = gopurs_runtime.Box(raw)
	})
	return gopurs_runtime.Apply(right, gopurs_runtime.Any(owned))
}

// TypedFieldId tags a Maybe-free field method with the element decoder's kind.
func TypedFieldId(inner, impl gopurs_runtime.Value) gopurs_runtime.Value {
	if plan, ok := gopurs_runtime.FunctionData[*recordDecodePlan](inner); ok && plan != nil {
		return gopurs_runtime.WithFunctionData(impl, plan)
	}
	if tagData, ok := gopurs_runtime.FunctionData[*typedKind](inner); ok && tagData != nil {
		return gopurs_runtime.WithFunctionData(impl, tagData)
	}
	return impl
}

// TypedFieldMaybe tags a Maybe field method so a missing key becomes Nothing.
func TypedFieldMaybe(inner, impl gopurs_runtime.Value) gopurs_runtime.Value {
	var errors *recordDecodePlan
	var innerKind *fieldKind
	if plan, ok := gopurs_runtime.FunctionData[*recordDecodePlan](inner); ok && plan != nil {
		errors = plan
		innerKind = &fieldKind{tag: kindRecord, plan: plan}
	} else if tagData, ok := gopurs_runtime.FunctionData[*typedKind](inner); ok && tagData != nil {
		errors = tagData.errors
		innerKind = tagData.kind
	}
	if errors == nil {
		return impl
	}
	return gopurs_runtime.WithFunctionData(impl, &typedKind{
		kind:   &fieldKind{tag: kindMaybe, inner: innerKind},
		errors: errors,
	})
}

// FieldStep carries a field method's tag onto the record step closure built in
// gDecodeJsonCons, which is what a record plan actually receives.
func FieldStep(field, step gopurs_runtime.Value) gopurs_runtime.Value {
	if plan, ok := gopurs_runtime.FunctionData[*recordDecodePlan](field); ok && plan != nil {
		return gopurs_runtime.WithFunctionData(step, plan)
	}
	if tagData, ok := gopurs_runtime.FunctionData[*typedKind](field); ok && tagData != nil {
		return gopurs_runtime.WithFunctionData(step, tagData)
	}
	return step
}

func RecordNilImpl(left, errors, right gopurs_runtime.Value) gopurs_runtime.Value {
	plan := &recordDecodePlan{left: left, right: right}
	plan.captureErrors(errors)
	return gopurs_runtime.WithFunctionData(gopurs_runtime.Func2(plan.decodeDirect), plan)
}

func RecordConsImpl(reflectName, step, tail, fallback, isRight, rightValue, right, left, errors, nothing, just gopurs_runtime.Value) gopurs_runtime.Value {
	tailPlan, known := gopurs_runtime.FunctionData[*recordDecodePlan](tail)
	if !known {
		return fallback
	}
	fields := make([]recordDecodeField, 1+len(tailPlan.fields))
	kind := detectFieldKind(step)
	field := recordDecodeField{reflectName: reflectName, step: step, kind: kind}
	fields[0] = field
	copy(fields[1:], tailPlan.fields)
	plan := &recordDecodePlan{
		fields:     fields,
		isRight:    isRight,
		rightValue: rightValue,
		right:      right,
		left:       left,
		nothing:    nothing,
		just:       just,
		hasCustom:  tailPlan.hasCustom || kind == nil,
	}
	plan.captureErrors(errors)
	return gopurs_runtime.WithFunctionData(gopurs_runtime.Func2(plan.decodeDirect), plan)
}
