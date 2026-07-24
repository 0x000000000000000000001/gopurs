const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/Runtime.purs', 'utf8');

code = code.replace(/func Str\(v string\) Value \{\n\treturn Value\{Type: TypeString, UnsafePtr: unsafe.Pointer\(&v\)\}\n\}/, 'func Str(v string) Value {\n\treturn Value{Type: TypeString, UnsafePtr: unsafe.Pointer(&v)}\n}\n\nfunc (v Value) StrVal() string {\n\treturn *(*string)(v.UnsafePtr)\n}');

code = code.replace(/func Array\(v \[\]Value\) Value \{\n\treturn Value\{Type: TypeArray, UnsafePtr: unsafe.Pointer\(&v\)\}\n\}/, 'func Array(v []Value) Value {\n\treturn Value{Type: TypeArray, UnsafePtr: unsafe.Pointer(&v)}\n}\n\nfunc (v Value) PtrVal() any {\n\treturn *(*any)(v.UnsafePtr)\n}');

fs.writeFileSync('src/Gopurs/Runtime.purs', code);
