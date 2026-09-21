package gopurs_runtime_test

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"

	rt "gopurs/output/gopurs_runtime"
)

type capturedObject struct {
	seed  int64
	child *int64
	data  [128]int64
}

type closureCase struct {
	value     rt.Value
	arity     int
	prefix    int
	seed      int64
	finalized *atomic.Bool
}

func capturedResult(capture *capturedObject, args ...rt.Value) rt.Value {
	result := capture.seed + *capture.child + capture.data[0] + capture.data[127]
	for index, arg := range args {
		value := rt.ArrayAccess(arg, 0).IntVal
		if rt.ArrayAccess(arg, 1).StrVal() != fmt.Sprintf("argument-%d", value) {
			panic("captured argument corrupted")
		}
		result += int64(index+1) * value
	}
	return rt.Int(result)
}

// Keep the creator frame separate from its caller: only the boxed Value may
// retain the function and its captures after this function returns.
//
//go:noinline
func makeClosure(arity int, seed int64, finalized *atomic.Bool) rt.Value {
	child := seed + 1
	capture := &capturedObject{seed: seed, child: &child}
	capture.data[0], capture.data[127] = seed+2, seed+3
	runtime.SetFinalizer(capture, func(*capturedObject) { finalized.Store(true) })
	switch arity {
	case 1:
		return rt.Func(func(a rt.Value) rt.Value { return capturedResult(capture, a) })
	case 2:
		return rt.Func2(func(a, b rt.Value) rt.Value { return capturedResult(capture, a, b) })
	case 3:
		return rt.Func3(func(a, b, c rt.Value) rt.Value { return capturedResult(capture, a, b, c) })
	case 4:
		return rt.Func4(func(a, b, c, d rt.Value) rt.Value { return capturedResult(capture, a, b, c, d) })
	case 5:
		return rt.Func5(func(a, b, c, d, e rt.Value) rt.Value { return capturedResult(capture, a, b, c, d, e) })
	case 6:
		return rt.Func6(func(a, b, c, d, e, f rt.Value) rt.Value { return capturedResult(capture, a, b, c, d, e, f) })
	case 7:
		return rt.Func7(func(a, b, c, d, e, f, g rt.Value) rt.Value { return capturedResult(capture, a, b, c, d, e, f, g) })
	case 8:
		return rt.Func8(func(a, b, c, d, e, f, g, h rt.Value) rt.Value { return capturedResult(capture, a, b, c, d, e, f, g, h) })
	case 9:
		return rt.Func9(func(a, b, c, d, e, f, g, h, i rt.Value) rt.Value {
			return capturedResult(capture, a, b, c, d, e, f, g, h, i)
		})
	case 10:
		return rt.Func10(func(a, b, c, d, e, f, g, h, i, j rt.Value) rt.Value {
			return capturedResult(capture, a, b, c, d, e, f, g, h, i, j)
		})
	case 11:
		return rt.Func11(func(a, b, c, d, e, f, g, h, i, j, k rt.Value) rt.Value {
			return capturedResult(capture, a, b, c, d, e, f, g, h, i, j, k)
		})
	default:
		panic("unsupported closure arity")
	}
}

func arguments(start, end int) []rt.Value {
	args := make([]rt.Value, 0, end-start)
	for index := start; index < end; index++ {
		value := int64(index + 1)
		args = append(args, rt.Array([]rt.Value{
			rt.Int(value), rt.Str(fmt.Sprintf("argument-%d", value)),
		}))
	}
	return args
}

func applyArguments(value rt.Value, args []rt.Value) rt.Value {
	switch len(args) {
	case 0:
		return value
	case 1:
		return rt.Apply(value, args[0])
	case 2:
		return rt.Apply2(value, args[0], args[1])
	case 3:
		return rt.Apply3(value, args[0], args[1], args[2])
	case 4:
		return rt.Apply4(value, args[0], args[1], args[2], args[3])
	case 5:
		return rt.Apply5(value, args[0], args[1], args[2], args[3], args[4])
	case 6:
		return rt.Apply6(value, args[0], args[1], args[2], args[3], args[4], args[5])
	case 7:
		return rt.Apply7(value, args[0], args[1], args[2], args[3], args[4], args[5], args[6])
	case 8:
		return rt.Apply8(value, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7])
	case 9:
		return rt.Apply9(value, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8])
	case 10:
		return rt.Apply10(value, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9])
	case 11:
		// There is no Apply11 helper.
		return rt.Apply(rt.Apply10(value, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9]), args[10])
	default:
		panic("unsupported argument count")
	}
}

func newClosureCase(arity, prefix int, seed int64) closureCase {
	finalized := new(atomic.Bool)
	value := makeClosure(arity, seed, finalized)
	value = applyArguments(value, arguments(0, prefix))
	return closureCase{value, arity, prefix, seed, finalized}
}

func verifyClosure(c closureCase) error {
	want := 4*c.seed + 6
	for index := 1; index <= c.arity; index++ {
		want += int64(index * index)
	}
	got := applyArguments(c.value, arguments(c.prefix, c.arity))
	finalized := c.finalized.Load()
	runtime.KeepAlive(c.value)
	if finalized {
		return fmt.Errorf("arity %d, prefix %d: capture finalized while closure is reachable", c.arity, c.prefix)
	}
	if got.Type != rt.TypeInt || got.IntVal != want {
		return fmt.Errorf("arity %d, prefix %d: got %+v, want %d", c.arity, c.prefix, got, want)
	}
	return nil
}

// Each live frame holds 4 KiB across recursion, forcing goroutine stack growth.
// Check the padding after returning so the compiler cannot discard it.
//
//go:noinline
func onGrownStack(depth int, create func() []closureCase) []closureCase {
	var padding [4096]byte
	for index := range padding {
		padding[index] = byte(index + depth)
	}
	var values []closureCase
	if depth == 0 {
		values = create()
	} else {
		values = onGrownStack(depth-1, create)
	}
	for index, value := range padding {
		if value != byte(index+depth) {
			panic("creator stack corrupted")
		}
	}
	return values
}

func afterCreatorReturns(create func() []closureCase) []closureCase {
	result := make(chan []closureCase)
	exited := make(chan struct{})
	go func() {
		defer close(exited)
		values := onGrownStack(32, create)
		// The recursive frames have unwound. Repeated collections let Go shrink
		// the creator stack while the boxed closures remain live.
		runtime.GC()
		runtime.GC()
		onGrownStack(32, func() []closureCase { return nil })
		result <- values
	}()
	values := <-result
	<-exited
	// Collect again after the entire creator goroutine has exited.
	runtime.GC()
	runtime.GC()
	return values
}

func TestClosureLifetime(t *testing.T) {
	values := afterCreatorReturns(func() []closureCase {
		var values []closureCase
		for arity := 1; arity <= 11; arity++ {
			values = append(values, newClosureCase(arity, 0, int64(100+arity)))
		}
		return values
	})
	for _, value := range values {
		t.Run(fmt.Sprintf("arity_%d", value.arity), func(t *testing.T) {
			if err := verifyClosure(value); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestPartialApplicationLifetime(t *testing.T) {
	values := afterCreatorReturns(func() []closureCase {
		var values []closureCase
		for arity := 2; arity <= 11; arity++ {
			for prefix := 1; prefix < arity; prefix++ {
				values = append(values, newClosureCase(arity, prefix, int64(100*arity+prefix)))
			}
		}
		return values
	})
	for _, value := range values {
		t.Run(fmt.Sprintf("arity_%d_apply_%d", value.arity, value.prefix), func(t *testing.T) {
			if err := verifyClosure(value); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestConcurrentClosureCreatorsAndCallers(t *testing.T) {
	const workers = 4
	shared := newClosureCase(11, 0, 9000)
	jobs := make(chan closureCase, workers)
	failures := make(chan error, workers)
	stopGC, gcDone := make(chan struct{}), make(chan struct{})
	go func() {
		defer close(gcDone)
		for {
			select {
			case <-stopGC:
				return
			default:
				runtime.GC()
			}
		}
	}()
	var creators, callers sync.WaitGroup
	for worker := 0; worker < workers; worker++ {
		creators.Add(1)
		go func(worker int) {
			defer creators.Done()
			for iteration := 0; iteration < 128; iteration++ {
				arity := 1 + iteration%11
				jobs <- newClosureCase(arity, iteration%arity, int64(worker*128+iteration))
			}
		}(worker)
		callers.Add(1)
		go func() {
			defer callers.Done()
			var failure error
			for value := range jobs {
				if err := verifyClosure(value); err != nil && failure == nil {
					failure = err
				}
				// Every caller also applies the same closure concurrently.
				if err := verifyClosure(shared); err != nil && failure == nil {
					failure = err
				}
			}
			if failure != nil {
				failures <- failure
			}
		}()
	}
	creators.Wait()
	close(jobs)
	callers.Wait()
	close(stopGC)
	<-gcDone
	close(failures)
	for err := range failures {
		t.Error(err)
	}
}
