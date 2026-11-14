package gopurs_runtime

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type ffiAffResult struct {
	value Value
	err   error
}

func ffiAffCallbacks() (Value, Value, <-chan ffiAffResult) {
	results := make(chan ffiAffResult, 2)
	onError := Func(func(value Value) Value {
		results <- ffiAffResult{err: Unbox[error](value)}
		return Value{}
	})
	onSuccess := Func(func(value Value) Value {
		results <- ffiAffResult{value: value}
		return Value{}
	})
	return onError, onSuccess, results
}

func ffiAffReceive[T any](t *testing.T, ch <-chan T) T {
	t.Helper()
	select {
	case value := <-ch:
		return value
	case <-time.After(5 * time.Second):
		t.Fatal("FFI operation did not finish")
		var zero T
		return zero
	}
}

func TestEffectFnAffDeferredRepeatedSuccessAndError(t *testing.T) {
	var calls atomic.Int32
	failure := errors.New("ordinary error")
	action := EffectFnAff(func(ctx context.Context) (Value, error) {
		if calls.Add(1) == 2 {
			return Value{}, failure
		}
		return Str("ordinary success"), nil
	})
	if calls.Load() != 0 {
		t.Fatal("constructing an action ran the operation")
	}
	for execution := 1; execution <= 3; execution++ {
		onError, onSuccess, results := ffiAffCallbacks()
		canceler := Apply2(action, onError, onSuccess)
		if canceler.Type != TypeFunc3 {
			t.Fatalf("EffectFnAff canceler has wrong arity: %d", canceler.Type)
		}
		result := ffiAffReceive(t, results)
		if execution == 2 {
			if result.err != failure {
				t.Fatalf("error identity was lost: %v", result.err)
			}
		} else if result.err != nil || Unbox[string](result.value) != "ordinary success" {
			t.Fatalf("unexpected success: %+v", result)
		}
	}
	if calls.Load() != 3 {
		t.Fatalf("executions were memoized: %d", calls.Load())
	}
}

func TestEffectFnAffInstallationReturnsBeforeBlockingOperation(t *testing.T) {
	started := make(chan context.Context, 1)
	finished := make(chan struct{})
	action := EffectFnAff(func(ctx context.Context) (Value, error) {
		defer close(finished)
		started <- ctx
		<-ctx.Done()
		return Value{}, ctx.Err()
	})
	onError, onSuccess, results := ffiAffCallbacks()
	installed := make(chan Value, 1)
	go func() { installed <- Apply2(action, onError, onSuccess) }()
	canceler := ffiAffReceive(t, installed)
	ctx := ffiAffReceive(t, started)
	reason := errors.New("requested cancellation")
	var acknowledgments atomic.Int32
	onCanceled := Func(func(_ Value) Value { acknowledgments.Add(1); return Value{} })
	for i := 0; i < 2; i++ {
		Apply3(canceler, Box(reason), onError, onCanceled)
	}
	ffiAffReceive(t, finished)
	if context.Cause(ctx) != reason || acknowledgments.Load() != 2 {
		t.Fatal("cancellation was not propagated or acknowledged")
	}
	select {
	case result := <-results:
		t.Fatalf("canceled operation delivered a result: %+v", result)
	default:
	}
}

func TestEffectFnAffCancellationBeforeWorkerStarts(t *testing.T) {
	onError, onSuccess, results := ffiAffCallbacks()
	operation := newFFIAffOperation(onError, onSuccess)
	operation.stop(Box(errors.New("cancel before scheduling")))
	operation.run(func(context.Context) (Value, error) {
		t.Fatal("canceled worker started an operation")
		return Value{}, nil
	})
	select {
	case <-results:
		t.Fatal("pre-canceled operation delivered a callback")
	default:
	}
}

func TestEffectFnAffLateResultAndCancellationRace(t *testing.T) {
	for _, fail := range []bool{false, true} {
		for iteration := 0; iteration < 250; iteration++ {
			var callbacks atomic.Int32
			callback := Func(func(_ Value) Value { callbacks.Add(1); return Value{} })
			operation := newFFIAffOperation(callback, callback)
			gate := make(chan struct{})
			var workers sync.WaitGroup
			workers.Add(2)
			go func() {
				defer workers.Done()
				operation.run(func(context.Context) (Value, error) {
					<-gate
					if fail {
						return Value{}, errors.New("failure")
					}
					return Str("success"), nil
				})
			}()
			go func() {
				defer workers.Done()
				<-gate
				operation.stop(Box(errors.New("cancel")))
			}()
			close(gate)
			workers.Wait()
			want := int32(0)
			if operation.state.Load() == 1 {
				want = 1
			}
			if callbacks.Load() != want {
				t.Fatalf("state=%d callbacks=%d", operation.state.Load(), callbacks.Load())
			}
			if operation.ctx.Err() == nil {
				t.Fatal("completed context was not released")
			}
		}
	}
}

func TestEffectFnAffNonInterruptibleWorkOwnsLateResources(t *testing.T) {
	for _, fail := range []bool{false, true} {
		onError, onSuccess, results := ffiAffCallbacks()
		operation := newFFIAffOperation(onError, onSuccess)
		started, finish, done := make(chan struct{}), make(chan struct{}), make(chan struct{})
		var releases atomic.Int32
		go func() {
			defer close(done)
			operation.run(func(ctx context.Context) (Value, error) {
				close(started)
				<-finish // Models an SDK that cannot stop its blocking call.
				if ctx.Err() != nil {
					releases.Add(1) // Explicit ownership; no inferred Close method.
				}
				if fail {
					return Value{}, errors.New("late error")
				}
				return Str("late result"), nil
			})
		}()
		ffiAffReceive(t, started)
		operation.stop(Box(errors.New("cancel")))
		select {
		case <-done:
			t.Fatal("cancellation claimed to stop non-interruptible work")
		default:
		}
		close(finish)
		ffiAffReceive(t, done)
		if releases.Load() != 1 {
			t.Fatal("resource was not released exactly once")
		}
		select {
		case <-results:
			t.Fatal("late result escaped cancellation")
		default:
		}
	}
}

type ffiNilError struct{}

func (*ffiNilError) Error() string { return "typed nil error" }

func TestEffectFnAffTypedNilPreservesGoSemantics(t *testing.T) {
	onError, onSuccess, results := ffiAffCallbacks()
	var failure *ffiNilError
	operation := newFFIAffOperation(onError, onSuccess)
	operation.run(func(context.Context) (Value, error) { return Value{}, failure })
	result := ffiAffReceive(t, results)
	if result.err == nil || result.err != error(failure) {
		t.Fatal("typed nil error was incorrectly treated as success")
	}

	onError, onSuccess, results = ffiAffCallbacks()
	operation = newFFIAffOperation(onError, onSuccess)
	operation.run(func(context.Context) (Value, error) { return Box(failure), nil })
	result = ffiAffReceive(t, results)
	if result.err != nil || Unbox[*ffiNilError](result.value) != nil {
		t.Fatal("typed nil success did not round-trip")
	}
}

func TestEffectFnAffNoRemainingWorkers(t *testing.T) {
	before := runtime.NumGoroutine()
	const count = 200
	var workers sync.WaitGroup
	workers.Add(count)
	callback := Func(func(_ Value) Value { return Value{} })
	for i := 0; i < count; i++ {
		operation := newFFIAffOperation(callback, callback)
		go func() {
			defer workers.Done()
			operation.run(func(ctx context.Context) (Value, error) {
				<-ctx.Done()
				return Value{}, ctx.Err()
			})
		}()
		operation.stop(Value{})
	}
	workers.Wait()
	// WaitGroup completion happens immediately before goroutine exit.
	deadline := time.Now().Add(5 * time.Second)
	for runtime.NumGoroutine() > before && time.Now().Before(deadline) {
		runtime.Gosched()
	}
	if after := runtime.NumGoroutine(); after > before {
		t.Fatalf("worker goroutines remain: before=%d after=%d", before, after)
	}
}

func TestEffectFnAffPanicIsNotBusinessError(t *testing.T) {
	if os.Getenv("GOPURS_FFI_PANIC_CHILD") == "1" {
		callback := Func(func(_ Value) Value { os.Exit(7); return Value{} })
		Apply2(EffectFnAff(func(context.Context) (Value, error) {
			panic("ffi panic must remain a panic")
		}), callback, callback)
		select {}
	}
	command := exec.Command(os.Args[0], "-test.run=^TestEffectFnAffPanicIsNotBusinessError$")
	command.Env = append(os.Environ(), "GOPURS_FFI_PANIC_CHILD=1")
	output, err := command.CombinedOutput()
	if err == nil || !strings.Contains(string(output), "panic: ffi panic must remain a panic") {
		t.Fatalf("panic was hidden: error=%v output=%s", err, output)
	}
}

func TestEffectFnAffAcquireTransfersOwnershipOnlyOnSuccess(t *testing.T) {
	for _, fail := range []bool{false, true} {
		var releases atomic.Int32
		failure := errors.New("acquisition failed")
		onError, onSuccess, results := ffiAffCallbacks()
		operation := newFFIAffOperation(onError, onSuccess)
		operation.runAcquire(func(context.Context) (Value, func(), error) {
			release := func() { releases.Add(1) }
			if fail {
				return Value{}, release, failure
			}
			return Str("resource"), release, nil
		})
		result := ffiAffReceive(t, results)
		if fail {
			if result.err != failure || releases.Load() != 1 {
				t.Fatal("failed acquisition did not release before reporting its error")
			}
		} else if result.err != nil || Unbox[string](result.value) != "resource" || releases.Load() != 0 {
			t.Fatal("successful acquisition did not transfer ownership")
		}
		operation.stop(Value{})
		wantReleases := int32(0)
		if fail {
			wantReleases = 1
		}
		if releases.Load() != wantReleases {
			t.Fatal("cancel after completion released a transferred resource")
		}
	}
}

func TestEffectFnAffAcquireReleasesLateSuccessAndFailureOnce(t *testing.T) {
	for _, fail := range []bool{false, true} {
		var releases atomic.Int32
		started, finish, released := make(chan struct{}), make(chan struct{}), make(chan struct{})
		onError, onSuccess, results := ffiAffCallbacks()
		action := EffectFnAffAcquire(func(context.Context) (Value, func(), error) {
			close(started)
			<-finish
			release := func() { releases.Add(1); close(released) }
			if fail {
				return Value{}, release, errors.New("late failure")
			}
			return Str("late acquisition"), release, nil
		})
		canceler := Apply2(action, onError, onSuccess)
		ffiAffReceive(t, started)
		var canceled atomic.Int32
		acknowledge := Func(func(Value) Value { canceled.Add(1); return Value{} })
		Apply3(canceler, Box(errors.New("cancel acquisition")), onError, acknowledge)
		if canceled.Load() != 1 || releases.Load() != 0 {
			t.Fatal("cancellation waited for or prematurely released a pending acquisition")
		}
		close(finish)
		ffiAffReceive(t, released)
		Apply3(canceler, Value{}, onError, acknowledge)
		if releases.Load() != 1 {
			t.Fatal("late acquisition released more than once")
		}
		select {
		case <-results:
			t.Fatal("late acquired result reached a stopped action")
		default:
		}
	}
}

func TestEffectFnAffAcquireReleaseRaceHasOneOwner(t *testing.T) {
	for iteration := 0; iteration < 250; iteration++ {
		var transfers, releases atomic.Int32
		callback := Func(func(Value) Value { transfers.Add(1); return Value{} })
		operation := newFFIAffOperation(callback, callback)
		started, finish := make(chan struct{}), make(chan struct{})
		var workers sync.WaitGroup
		workers.Add(2)
		go func() {
			defer workers.Done()
			operation.runAcquire(func(context.Context) (Value, func(), error) {
				close(started)
				<-finish
				return Str("resource"), func() { releases.Add(1) }, nil
			})
		}()
		ffiAffReceive(t, started)
		go func() { defer workers.Done(); <-finish; operation.stop(Value{}) }()
		close(finish)
		workers.Wait()
		if transfers.Load()+releases.Load() != 1 {
			t.Fatalf("acquisition ownership lost/duplicated: transfers=%d releases=%d", transfers.Load(), releases.Load())
		}
	}
}

func TestEffectFnAffCallbackCanInvokeItsOwnCanceler(t *testing.T) {
	start := make(chan struct{})
	returned, acknowledged := make(chan struct{}), make(chan struct{})
	var canceler Value
	onCanceled := Func(func(Value) Value {
		select {
		case <-returned:
		default:
			t.Error("cancellation was acknowledged before the winning callback returned")
		}
		close(acknowledged)
		return Value{}
	})
	onSuccess := Func(func(Value) Value {
		Apply3(canceler, Value{}, Value{}, onCanceled)
		close(returned)
		return Value{}
	})
	action := EffectFnAff(func(context.Context) (Value, error) {
		<-start
		return Str("success"), nil
	})
	canceler = Apply2(action, Value{}, onSuccess)
	close(start)
	ffiAffReceive(t, returned)
	ffiAffReceive(t, acknowledged)
}
