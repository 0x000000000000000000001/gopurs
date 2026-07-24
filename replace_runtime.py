import re

with open('src/Gopurs/Runtime.purs', 'r') as f:
    content = f.read()

# Replace struct field
content = content.replace("Func   unsafe.Pointer", "UnsafePtr unsafe.Pointer")

# Replace field accesses
content = re.sub(r'\.Func\b', '.UnsafePtr', content)

# Replace field assignments for functions
content = re.sub(r'Func: \*\(\*unsafe\.Pointer\)', 'UnsafePtr: *(*unsafe.Pointer)', content)

# Replace Constructor implementations
content = content.replace(
"""type ConstructorData []Value

func Constructor(tag string, args []Value) Value {
	return Value{Type: TypeConstructor, StrVal: tag, PtrVal: ConstructorData(args)}
}

type ConstructorData0 struct{}
func Constructor0(tag string) Value { return Value{Type: TypeConstructor, StrVal: tag, PtrVal: &ConstructorData0{}} }

type ConstructorData1 struct { V0 Value }
func Constructor1(tag string, v0 Value) Value { return Value{Type: TypeConstructor, StrVal: tag, PtrVal: &ConstructorData1{v0}} }

type ConstructorData2 struct { V0, V1 Value }
func Constructor2(tag string, v0, v1 Value) Value { return Value{Type: TypeConstructor, StrVal: tag, PtrVal: &ConstructorData2{v0, v1}} }

type ConstructorData3 struct { V0, V1, V2 Value }
func Constructor3(tag string, v0, v1, v2 Value) Value { return Value{Type: TypeConstructor, StrVal: tag, PtrVal: &ConstructorData3{v0, v1, v2}} }

type ConstructorData4 struct { V0, V1, V2, V3 Value }
func Constructor4(tag string, v0, v1, v2, v3 Value) Value { return Value{Type: TypeConstructor, StrVal: tag, PtrVal: &ConstructorData4{v0, v1, v2, v3}} }

type ConstructorData5 struct { V0, V1, V2, V3, V4 Value }
func Constructor5(tag string, v0, v1, v2, v3, v4 Value) Value { return Value{Type: TypeConstructor, StrVal: tag, PtrVal: &ConstructorData5{v0, v1, v2, v3, v4}} }""",
"""type ConstructorData []Value

func Constructor(tag string, args []Value) Value {
	ptr := unsafe.Pointer(nil)
	if len(args) > 0 { ptr = unsafe.Pointer(&args[0]) }
	return Value{Type: TypeConstructor, StrVal: tag, UnsafePtr: ptr}
}

type ConstructorData0 struct{}
func Constructor0(tag string) Value { return Value{Type: TypeConstructor, StrVal: tag, UnsafePtr: unsafe.Pointer(&ConstructorData0{})} }

type ConstructorData1 struct { V0 Value }
func Constructor1(tag string, v0 Value) Value { return Value{Type: TypeConstructor, StrVal: tag, UnsafePtr: unsafe.Pointer(&ConstructorData1{v0})} }

type ConstructorData2 struct { V0, V1 Value }
func Constructor2(tag string, v0, v1 Value) Value { return Value{Type: TypeConstructor, StrVal: tag, UnsafePtr: unsafe.Pointer(&ConstructorData2{v0, v1})} }

type ConstructorData3 struct { V0, V1, V2 Value }
func Constructor3(tag string, v0, v1, v2 Value) Value { return Value{Type: TypeConstructor, StrVal: tag, UnsafePtr: unsafe.Pointer(&ConstructorData3{v0, v1, v2})} }

type ConstructorData4 struct { V0, V1, V2, V3 Value }
func Constructor4(tag string, v0, v1, v2, v3 Value) Value { return Value{Type: TypeConstructor, StrVal: tag, UnsafePtr: unsafe.Pointer(&ConstructorData4{v0, v1, v2, v3})} }

type ConstructorData5 struct { V0, V1, V2, V3, V4 Value }
func Constructor5(tag string, v0, v1, v2, v3, v4 Value) Value { return Value{Type: TypeConstructor, StrVal: tag, UnsafePtr: unsafe.Pointer(&ConstructorData5{v0, v1, v2, v3, v4})} }"""
)

with open('src/Gopurs/Runtime.purs', 'w') as f:
    f.write(content)
