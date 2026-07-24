const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/Runtime.purs', 'utf8');

// 1. Remove StrVal and PtrVal
code = code.replace(/	StrVal string\n	PtrVal any\n/, '');

// 2. Add TypeRecords
code = code.replace(/	TypeFunc10 = 18\n\)/, `	TypeFunc10 = 18
	TypeRecord0 = 19
	TypeRecord1 = 20
	TypeRecord2 = 21
	TypeRecord3 = 22
	TypeRecord4 = 23
	TypeRecord5 = 24
	TypeRecordData = 25
)`);

// 3. Str
code = code.replace(/func Str\(v string\) Value \{\n\treturn Value\{Type: TypeString, StrVal: v\}\n\}/, 'func Str(v string) Value {\n\treturn Value{Type: TypeString, UnsafePtr: unsafe.Pointer(&v)}\n}');
// StrVal
code = code.replace(/func \(v Value\) StrVal\(\) string \{\n\treturn v\.StrVal\n\}/, 'func (v Value) StrVal() string {\n\treturn *(*string)(v.UnsafePtr)\n}');

// 4. Array
code = code.replace(/func Array\(v \[\]Value\) Value \{\n\treturn Value\{Type: TypeArray, PtrVal: v\}\n\}/, 'func Array(v []Value) Value {\n\treturn Value{Type: TypeArray, UnsafePtr: unsafe.Pointer(&v)}\n}');
// ArrayAccess
code = code.replace(/return arr\.PtrVal\.\(\[\]Value\)\[index\]/, 'return (*(*[]Value)(arr.UnsafePtr))[index]');

// 5. Any
code = code.replace(/func Any\(v any\) Value \{\n\treturn Value\{Type: TypeAny, PtrVal: v\}\n\}/, 'func Any(v any) Value {\n\treturn Value{Type: TypeAny, UnsafePtr: unsafe.Pointer(&v)}\n}');

// 6. FuncAny
code = code.replace(/func FuncAny\(f any\) Value \{\n\treturn Value\{Type: TypeFunc, PtrVal: f\}\n\}/, 'func FuncAny(f any) Value {\n\treturn Value{Type: TypeFunc, UnsafePtr: unsafe.Pointer(&f)}\n}');

// 7. RecordDict
code = code.replace(/func RecordDict\(keys \[\]string, vals \[\]Value\) Value \{\n\treturn Value\{Type: TypeRecord, PtrVal: RecordData\{keys, vals\}\}\n\}/, 'func RecordDict(keys []string, vals []Value) Value {\n\tr := RecordData{keys, vals}\n\treturn Value{Type: TypeRecordData, UnsafePtr: unsafe.Pointer(&r)}\n}');

// 8. RecordDict0-5
code = code.replace(/return Value\{Type: TypeRecord, PtrVal: &RecordData0\{\}\}/, 'return Value{Type: TypeRecord0, UnsafePtr: unsafe.Pointer(&RecordData0{})}');
code = code.replace(/return Value\{Type: TypeRecord, PtrVal: &RecordData1\{k0, v0\}\}/, 'return Value{Type: TypeRecord1, UnsafePtr: unsafe.Pointer(&RecordData1{k0, v0})}');
code = code.replace(/return Value\{Type: TypeRecord, PtrVal: &RecordData2\{k0, k1, v0, v1\}\}/, 'return Value{Type: TypeRecord2, UnsafePtr: unsafe.Pointer(&RecordData2{k0, k1, v0, v1})}');
code = code.replace(/return Value\{Type: TypeRecord, PtrVal: &RecordData3\{k0, k1, k2, v0, v1, v2\}\}/, 'return Value{Type: TypeRecord3, UnsafePtr: unsafe.Pointer(&RecordData3{k0, k1, k2, v0, v1, v2})}');
code = code.replace(/return Value\{Type: TypeRecord, PtrVal: &RecordData4\{k0, k1, k2, k3, v0, v1, v2, v3\}\}/, 'return Value{Type: TypeRecord4, UnsafePtr: unsafe.Pointer(&RecordData4{k0, k1, k2, k3, v0, v1, v2, v3})}');
code = code.replace(/return Value\{Type: TypeRecord, PtrVal: &RecordData5\{k0, k1, k2, k3, k4, v0, v1, v2, v3, v4\}\}/, 'return Value{Type: TypeRecord5, UnsafePtr: unsafe.Pointer(&RecordData5{k0, k1, k2, k3, k4, v0, v1, v2, v3, v4})}');

// 9. RecordToMap
code = code.replace(/func RecordToMap\(obj Value\) map\[string\]Value \{[\s\S]*?\n\}\n/, `func RecordToMap(obj Value) map[string]Value {
	if obj.Type == TypeRecord {
		m := *(*map[string]Value)(obj.UnsafePtr)
		res := make(map[string]Value, len(m))
		for k, v := range m { res[k] = v }
		return res
	}
	res := make(map[string]Value)
	switch obj.Type {
	case TypeRecord0:
	case TypeRecord1:
		r := (*RecordData1)(obj.UnsafePtr); res[r.K0] = r.V0
	case TypeRecord2:
		r := (*RecordData2)(obj.UnsafePtr); res[r.K0] = r.V0; res[r.K1] = r.V1
	case TypeRecord3:
		r := (*RecordData3)(obj.UnsafePtr); res[r.K0] = r.V0; res[r.K1] = r.V1; res[r.K2] = r.V2
	case TypeRecord4:
		r := (*RecordData4)(obj.UnsafePtr); res[r.K0] = r.V0; res[r.K1] = r.V1; res[r.K2] = r.V2; res[r.K3] = r.V3
	case TypeRecord5:
		r := (*RecordData5)(obj.UnsafePtr); res[r.K0] = r.V0; res[r.K1] = r.V1; res[r.K2] = r.V2; res[r.K3] = r.V3; res[r.K4] = r.V4
	case TypeRecordData:
		r := (*RecordData)(obj.UnsafePtr)
		for i, k := range r.Keys { res[k] = r.Vals[i] }
	}
	return res
}
`);

// 10. RecordGet
code = code.replace(/func RecordGet\(obj Value, key string\) Value \{[\s\S]*?\n\}\n/, `func RecordGet(obj Value, key string) Value {
    switch obj.Type {
    case TypeRecord: return (*(*map[string]Value)(obj.UnsafePtr))[key]
	case TypeRecord1:
		r := (*RecordData1)(obj.UnsafePtr)
		if r.K0 == key { return r.V0 }
	case TypeRecord2:
		r := (*RecordData2)(obj.UnsafePtr)
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
	case TypeRecord3:
		r := (*RecordData3)(obj.UnsafePtr)
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
		if r.K2 == key { return r.V2 }
	case TypeRecord4:
		r := (*RecordData4)(obj.UnsafePtr)
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
		if r.K2 == key { return r.V2 }
		if r.K3 == key { return r.V3 }
	case TypeRecord5:
		r := (*RecordData5)(obj.UnsafePtr)
		if r.K0 == key { return r.V0 }
		if r.K1 == key { return r.V1 }
		if r.K2 == key { return r.V2 }
		if r.K3 == key { return r.V3 }
		if r.K4 == key { return r.V4 }
	case TypeRecordData:
		r := (*RecordData)(obj.UnsafePtr)
		for i, k := range r.Keys {
			if k == key { return r.Vals[i] }
		}
	}
	panic("Key not found in record: " + key)
}
`);

// 11. Record (from map)
code = code.replace(/func Record\(m map\[string\]Value\) Value \{\n\tkeys := make\(\[\]string, 0, len\(m\)\)\n\tvals := make\(\[\]Value, 0, len\(m\)\)\n\tfor k, v := range m \{\n\t\tkeys = append\(keys, k\)\n\t\tvals = append\(vals, v\)\n\t\}\n\treturn RecordDict\(keys, vals\)\n\}/, 'func Record(m map[string]Value) Value {\n\treturn Value{Type: TypeRecord, UnsafePtr: unsafe.Pointer(&m)}\n}');

// 12. RecordUpdate
code = code.replace(/return Value\{Type: TypeRecord, PtrVal: m\}/, 'return Record(m)');

// 13. Unbox (replaces all Unbox cases reading from StrVal or PtrVal)
code = code.replace(/case string: return any\(v\.StrVal\)\.\(T\)/, 'case string: return any(*(*string)(v.UnsafePtr)).(T)');
code = code.replace(/default: return v\.PtrVal\.\(T\)/, 'default: return *(*T)(v.UnsafePtr)');

fs.writeFileSync('src/Gopurs/Runtime.purs', code);
