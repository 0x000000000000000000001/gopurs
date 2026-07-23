package main
import (
	"os"
	"strings"
)
func main() {
	b, err := os.ReadFile("src/Gopurs/Runtime.purs")
	if err != nil { panic(err) }
	s := string(b)
	
	s = strings.Replace(s, "TypeFunc = 4", "TypeFunc = 4\n\tTypeFunc2 = 10\n\tTypeFunc3 = 11\n\tTypeFunc4 = 12\n\tTypeFunc5 = 13", 1)
	
	funcDefinitions := `
// Function constructors
func Func2(f func(Value, Value) Value) Value { return Value{Type: TypeFunc2, PtrVal: f} }
func Func3(f func(Value, Value, Value) Value) Value { return Value{Type: TypeFunc3, PtrVal: f} }
func Func4(f func(Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc4, PtrVal: f} }
func Func5(f func(Value, Value, Value, Value, Value) Value) Value { return Value{Type: TypeFunc5, PtrVal: f} }
`
	s = strings.Replace(s, "func FuncAny(f any) Value {", funcDefinitions + "\nfunc FuncAny(f any) Value {", 1)
	
	applyReplacement := `
func Apply(f Value, arg Value) Value {
	switch f.Type {
	case TypeFunc:
		return f.PtrVal.(func(Value) Value)(arg)
	case TypeFunc2:
		fn := f.PtrVal.(func(Value, Value) Value)
		return Func(func(a Value) Value { return fn(arg, a) })
	case TypeFunc3:
		fn := f.PtrVal.(func(Value, Value, Value) Value)
		return Func2(func(a, b Value) Value { return fn(arg, a, b) })
	case TypeFunc4:
		fn := f.PtrVal.(func(Value, Value, Value, Value) Value)
		return Func3(func(a, b, c Value) Value { return fn(arg, a, b, c) })
	case TypeFunc5:
		fn := f.PtrVal.(func(Value, Value, Value, Value, Value) Value)
		return Func4(func(a, b, c, d Value) Value { return fn(arg, a, b, c, d) })
	default:
		panic("Attempted to apply a non-function")
	}
}

func Apply2(f Value, arg1, arg2 Value) Value {
	if f.Type == TypeFunc2 {
		return f.PtrVal.(func(Value, Value) Value)(arg1, arg2)
	}
	return Apply(Apply(f, arg1), arg2)
}

func Apply3(f Value, arg1, arg2, arg3 Value) Value {
	if f.Type == TypeFunc3 {
		return f.PtrVal.(func(Value, Value, Value) Value)(arg1, arg2, arg3)
	}
	return Apply(Apply2(f, arg1, arg2), arg3)
}

func Apply4(f Value, arg1, arg2, arg3, arg4 Value) Value {
	if f.Type == TypeFunc4 {
		return f.PtrVal.(func(Value, Value, Value, Value) Value)(arg1, arg2, arg3, arg4)
	}
	return Apply(Apply3(f, arg1, arg2, arg3), arg4)
}

func Apply5(f Value, arg1, arg2, arg3, arg4, arg5 Value) Value {
	if f.Type == TypeFunc5 {
		return f.PtrVal.(func(Value, Value, Value, Value, Value) Value)(arg1, arg2, arg3, arg4, arg5)
	}
	return Apply(Apply4(f, arg1, arg2, arg3, arg4), arg5)
}
`
	
	// Replace Apply function
	startIdx := strings.Index(s, "// Uncurried application helper\nfunc Apply(f Value, arg Value) Value {")
	endIdx := strings.Index(s[startIdx:], "}\n\nfunc ArrayAccess") + startIdx + 1
	
	s = s[:startIdx] + applyReplacement + s[endIdx:]
	
	os.WriteFile("src/Gopurs/Runtime.purs", []byte(s), 0644)
}
