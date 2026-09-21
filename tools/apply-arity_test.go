package gopurs_runtime_test

import (
	"fmt"
	"reflect"
	"testing"

	rt "gopurs/output/gopurs_runtime"
)

func arityCallback(arity int, callback func([]rt.Value) rt.Value) rt.Value {
	switch arity {
	case 1:
		return rt.Func(func(a rt.Value) rt.Value { return callback([]rt.Value{a}) })
	case 2:
		return rt.Func2(func(a, b rt.Value) rt.Value { return callback([]rt.Value{a, b}) })
	case 3:
		return rt.Func3(func(a, b, c rt.Value) rt.Value { return callback([]rt.Value{a, b, c}) })
	case 4:
		return rt.Func4(func(a, b, c, d rt.Value) rt.Value { return callback([]rt.Value{a, b, c, d}) })
	case 5:
		return rt.Func5(func(a, b, c, d, e rt.Value) rt.Value { return callback([]rt.Value{a, b, c, d, e}) })
	case 6:
		return rt.Func6(func(a, b, c, d, e, f rt.Value) rt.Value { return callback([]rt.Value{a, b, c, d, e, f}) })
	case 7:
		return rt.Func7(func(a, b, c, d, e, f, g rt.Value) rt.Value { return callback([]rt.Value{a, b, c, d, e, f, g}) })
	case 8:
		return rt.Func8(func(a, b, c, d, e, f, g, h rt.Value) rt.Value { return callback([]rt.Value{a, b, c, d, e, f, g, h}) })
	case 9:
		return rt.Func9(func(a, b, c, d, e, f, g, h, i rt.Value) rt.Value {
			return callback([]rt.Value{a, b, c, d, e, f, g, h, i})
		})
	case 10:
		return rt.Func10(func(a, b, c, d, e, f, g, h, i, j rt.Value) rt.Value {
			return callback([]rt.Value{a, b, c, d, e, f, g, h, i, j})
		})
	case 11:
		return rt.Func11(func(a, b, c, d, e, f, g, h, i, j, k rt.Value) rt.Value {
			return callback([]rt.Value{a, b, c, d, e, f, g, h, i, j, k})
		})
	default:
		panic("unsupported arity")
	}
}

// Every argument is observable. The first function consumes its whole arity;
// additional arguments pass through unary closures returned by that function.
func TestApplyArityMatrix(t *testing.T) {
	for arity := 1; arity <= 11; arity++ {
		for supplied := 2; supplied <= 10; supplied++ {
			t.Run(fmt.Sprintf("arity_%d_apply_%d", arity, supplied), func(t *testing.T) {
				total := arity
				if supplied > total {
					total = supplied
				}
				args := make([]rt.Value, total)
				want := make([]int64, total)
				for i := range args {
					args[i] = rt.Int(int64(i + 1))
					want[i] = int64(i + 1)
				}
				var seen []int64
				calls := 0
				var next func([]rt.Value) rt.Value
				next = func(values []rt.Value) rt.Value {
					calls++
					for _, value := range values {
						seen = append(seen, value.IntVal)
					}
					if len(seen) == total {
						return rt.Int(42)
					}
					return rt.Func(func(value rt.Value) rt.Value { return next([]rt.Value{value}) })
				}
				result := applyArguments(arityCallback(arity, next), args[:supplied])
				if supplied < arity {
					if calls != 0 {
						t.Fatalf("partial application invoked function %d times", calls)
					}
					result = applyArguments(result, args[supplied:])
				}
				if !reflect.DeepEqual(seen, want) {
					t.Fatalf("argument order: got %v, want %v", seen, want)
				}
				wantCalls := 1
				if supplied > arity {
					wantCalls += supplied - arity
				}
				if calls != wantCalls {
					t.Fatalf("got %d invocations, want %d", calls, wantCalls)
				}
				if result.Type != rt.TypeInt || result.IntVal != 42 {
					t.Fatalf("got %+v, want Int(42)", result)
				}
			})
		}
	}
}

func allocationFunction(arity int) rt.Value {
	switch arity {
	case 2:
		return rt.Func2(func(a, b rt.Value) rt.Value { return rt.Int(1*a.IntVal + 2*b.IntVal) })
	case 3:
		return rt.Func3(func(a, b, c rt.Value) rt.Value { return rt.Int(1*a.IntVal + 2*b.IntVal + 3*c.IntVal) })
	case 4:
		return rt.Func4(func(a, b, c, d rt.Value) rt.Value { return rt.Int(1*a.IntVal + 2*b.IntVal + 3*c.IntVal + 4*d.IntVal) })
	case 5:
		return rt.Func5(func(a, b, c, d, e rt.Value) rt.Value {
			return rt.Int(1*a.IntVal + 2*b.IntVal + 3*c.IntVal + 4*d.IntVal + 5*e.IntVal)
		})
	case 6:
		return rt.Func6(func(a, b, c, d, e, f rt.Value) rt.Value {
			return rt.Int(1*a.IntVal + 2*b.IntVal + 3*c.IntVal + 4*d.IntVal + 5*e.IntVal + 6*f.IntVal)
		})
	case 7:
		return rt.Func7(func(a, b, c, d, e, f, g rt.Value) rt.Value {
			return rt.Int(1*a.IntVal + 2*b.IntVal + 3*c.IntVal + 4*d.IntVal + 5*e.IntVal + 6*f.IntVal + 7*g.IntVal)
		})
	case 8:
		return rt.Func8(func(a, b, c, d, e, f, g, h rt.Value) rt.Value {
			return rt.Int(1*a.IntVal + 2*b.IntVal + 3*c.IntVal + 4*d.IntVal + 5*e.IntVal + 6*f.IntVal + 7*g.IntVal + 8*h.IntVal)
		})
	case 9:
		return rt.Func9(func(a, b, c, d, e, f, g, h, i rt.Value) rt.Value {
			return rt.Int(1*a.IntVal + 2*b.IntVal + 3*c.IntVal + 4*d.IntVal + 5*e.IntVal + 6*f.IntVal + 7*g.IntVal + 8*h.IntVal + 9*i.IntVal)
		})
	case 10:
		return rt.Func10(func(a, b, c, d, e, f, g, h, i, j rt.Value) rt.Value {
			return rt.Int(1*a.IntVal + 2*b.IntVal + 3*c.IntVal + 4*d.IntVal + 5*e.IntVal + 6*f.IntVal + 7*g.IntVal + 8*h.IntVal + 9*i.IntVal + 10*j.IntVal)
		})
	default:
		panic("unsupported allocation test arity")
	}
}

var applicationResult rt.Value

func applicationPanic(run func()) (caught any) {
	defer func() { caught = recover() }()
	run()
	return nil
}

func TestApplicationPanics(t *testing.T) {
	for count := 2; count <= 10; count++ {
		for name, value := range map[string]rt.Value{
			"integer": rt.Int(1),
			"string":  rt.Str("not a function"),
			"overapplied": rt.Func(func(arg rt.Value) rt.Value {
				return arg
			}),
		} {
			t.Run(fmt.Sprintf("%s_apply_%d", name, count), func(t *testing.T) {
				args := make([]rt.Value, count)
				for index := range args {
					args[index] = rt.Int(int64(index))
				}
				want := applicationPanic(func() {
					result := value
					for _, arg := range args {
						result = rt.Apply(result, arg)
					}
				})
				got := applicationPanic(func() { applyArguments(value, args) })
				if want == nil || got == nil || fmt.Sprint(got) != fmt.Sprint(want) {
					t.Fatalf("batched panic %v, sequential panic %v", got, want)
				}
			})
		}
	}
}

// Closure construction and argument creation stay outside the measured region.
// Saturated applications must not build intermediate currying closures.
func TestSaturatedApplicationAllocations(t *testing.T) {
	for arity := 2; arity <= 10; arity++ {
		t.Run(fmt.Sprintf("arity_%d", arity), func(t *testing.T) {
			args := make([]rt.Value, arity)
			var want int64
			for i := range args {
				args[i] = rt.Int(int64(i + 1))
				want += int64((i + 1) * (i + 1))
			}
			function := allocationFunction(arity)
			allocs := testing.AllocsPerRun(1000, func() { applicationResult = applyArguments(function, args) })
			if applicationResult.IntVal != want {
				t.Fatalf("got %d, want %d", applicationResult.IntVal, want)
			}
			if allocs != 0 {
				t.Fatalf("saturated Apply%d allocated %g times per call, want 0", arity, allocs)
			}
		})
	}
}
