// Copy into a retained native bootstrap output/purescript directory.
// Run: go test -race -run '^TestPipelineNative' -count=1 -timeout 30s ./purescript
package purescript

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	rt "gopurs/output/gopurs_runtime"
)

const pipelineNativeTimeout = 3 * time.Second

type pipelineNativeGate struct {
	ch   chan struct{}
	once sync.Once
}

func newPipelineNativeGate() *pipelineNativeGate {
	return &pipelineNativeGate{ch: make(chan struct{})}
}

func (g *pipelineNativeGate) release() { g.once.Do(func() { close(g.ch) }) }

func pipelineNativeEmitter(emit func([]int64) AffFn) rt.Value {
	callback := rt.Func(func(batch rt.Value) rt.Value {
		ids := make([]int64, rt.ArrayLength(batch))
		for i := range ids {
			ids[i] = rt.ArrayAccess(batch, i).IntVal
		}
		return rt.Box(emit(ids))
	})
	effect := rt.Apply2(Get_Gopurs_Emission_createPipelinedEmitter(), rt.Int(2), callback)
	return rt.Apply(effect, rt.Value{})
}

func pipelineNativeEnqueue(emitter rt.Value, id int64) rt.Value {
	return rt.Apply(rt.RecordGet(emitter, "enqueue"), pipelineNativeEntry(id))
}

func pipelineNativeEntry(id int64) rt.Value {
	return rt.RecordDict3("imports", "name", "value",
		Get_Data_Set_empty(), rt.Str(fmt.Sprintf("Pipeline%d", id)), rt.Int(id))
}

func pipelineNativeStart(aff rt.Value) <-chan error {
	done := make(chan error, 1)
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), pipelineNativeTimeout)
		defer cancel()
		_, err := runAffSync(rt.Unbox[AffFn](aff), ctx)
		done <- err
	}()
	return done
}

func pipelineNativeAwait(t *testing.T, label string, done <-chan error) error {
	t.Helper()
	select {
	case err := <-done:
		return err
	case <-time.After(pipelineNativeTimeout):
		t.Fatalf("%s did not return within %s", label, pipelineNativeTimeout)
		return nil
	}
}

func pipelineNativeRun(t *testing.T, label string, aff rt.Value) {
	t.Helper()
	if err := pipelineNativeAwait(t, label, pipelineNativeStart(aff)); err != nil {
		t.Fatalf("%s: %v", label, err)
	}
}

func pipelineNativeBlocked(t *testing.T, label string, done <-chan error) {
	t.Helper()
	select {
	case err := <-done:
		t.Fatalf("%s returned while the worker was suspended: %v", label, err)
	case <-time.After(50 * time.Millisecond):
	}
}

func pipelineNativeSignal(t *testing.T, label string, signal <-chan struct{}) {
	t.Helper()
	select {
	case <-signal:
	case <-time.After(pipelineNativeTimeout):
		t.Fatalf("%s was not observed", label)
	}
}

func pipelineNativeCleanup(t *testing.T, emitter rt.Value, gates ...*pipelineNativeGate) {
	t.Helper()
	t.Cleanup(func() {
		// Release test gates before draining the fiber, including after a failed
		// assertion, so the test cannot strand a deliberately suspended worker.
		for _, gate := range gates {
			gate.release()
		}
		if err := pipelineNativeAwait(t, "cleanup cancel", pipelineNativeStart(rt.RecordGet(emitter, "cancel"))); err != nil {
			t.Errorf("cleanup cancel: %v", err)
		}
	})
}

func TestPipelineNativeOverlapOrderAndFinish(t *testing.T) {
	first, second := newPipelineNativeGate(), newPipelineNativeGate()
	constructed := make(chan []int64, 4)
	constructedTooSoon := make(chan struct{}, 1)
	firstStarted, secondStarted, firstStopped := make(chan struct{}), make(chan struct{}), make(chan struct{})
	emitter := pipelineNativeEmitter(func(ids []int64) AffFn {
		// Observe callback construction as well as execution: metadata must not
		// be captured for the second batch before the first batch has finished.
		if ids[0] == 3 {
			select {
			case <-firstStopped:
			default:
				constructedTooSoon <- struct{}{}
			}
		}
		constructed <- ids
		return func(ctx context.Context) (any, error) {
			var gate <-chan struct{}
			if ids[0] == 1 {
				close(firstStarted)
				defer close(firstStopped)
				gate = first.ch
			} else {
				close(secondStarted)
				gate = second.ch
			}
			select {
			case <-gate:
				return rt.Value{}, nil
			case <-ctx.Done():
				return nil, context.Cause(ctx)
			}
		}
	})
	pipelineNativeCleanup(t, emitter, first, second)
	pipelineNativeRun(t, "enqueue 1", pipelineNativeEnqueue(emitter, 1))
	pipelineNativeRun(t, "enqueue 2 must return before emission completes", pipelineNativeEnqueue(emitter, 2))
	pipelineNativeSignal(t, "first worker start", firstStarted)
	if got := <-constructed; !reflect.DeepEqual(got, []int64{1, 2}) {
		t.Fatalf("first batch: %v", got)
	}
	pipelineNativeRun(t, "enqueue pending 3", pipelineNativeEnqueue(emitter, 3))
	// Two jobs allow four batches in flight. The second batch is admitted
	// immediately, but must wait for its predecessor before reading metadata.
	pipelineNativeRun(t, "enqueue second batch before first completes", pipelineNativeEnqueue(emitter, 4))
	select {
	case ids := <-constructed:
		t.Fatalf("second callback constructed before first finished: %v", ids)
	default:
	}
	first.release()
	pipelineNativeSignal(t, "second worker start", secondStarted)
	select {
	case <-constructedTooSoon:
		t.Fatal("second callback captured its state before the first callback returned")
	default:
	}
	if got := <-constructed; !reflect.DeepEqual(got, []int64{3, 4}) {
		t.Fatalf("second batch: %v", got)
	}
	finish := pipelineNativeStart(rt.RecordGet(emitter, "finish"))
	pipelineNativeBlocked(t, "finish", finish)
	second.release()
	if err := pipelineNativeAwait(t, "finish after second worker stops", finish); err != nil {
		t.Fatal(err)
	}
}

func TestPipelineNativePropagatesEmissionErrors(t *testing.T) {
	for _, via := range []string{"finish", "backpressure"} {
		t.Run(via, func(t *testing.T) {
			failure := errors.New("pipeline-native-emission-failed")
			fail := newPipelineNativeGate()
			calls := make(chan struct{}, 4)
			emitter := pipelineNativeEmitter(func(_ []int64) AffFn {
				calls <- struct{}{}
				return func(ctx context.Context) (any, error) {
					select {
					case <-fail.ch:
						return nil, failure
					case <-ctx.Done():
						return nil, context.Cause(ctx)
					}
				}
			})
			pipelineNativeCleanup(t, emitter, fail)
			pipelineNativeRun(t, "enqueue 1", pipelineNativeEnqueue(emitter, 1))
			pipelineNativeRun(t, "enqueue 2", pipelineNativeEnqueue(emitter, 2))
			action := rt.RecordGet(emitter, "finish")
			if via == "backpressure" {
				// Fill four batches and leave one pending module. The fifth
				// batch must wait for a slot, then report its predecessor's error.
				for id := int64(3); id <= 9; id++ {
					pipelineNativeRun(t, fmt.Sprintf("enqueue %d", id), pipelineNativeEnqueue(emitter, id))
				}
				action = pipelineNativeEnqueue(emitter, 10)
			}
			done := pipelineNativeStart(action)
			pipelineNativeBlocked(t, via, done)
			fail.release()
			err := pipelineNativeAwait(t, via, done)
			if err == nil || !strings.Contains(err.Error(), failure.Error()) {
				t.Fatalf("%s: got %v, want emission error", via, err)
			}
			if got := len(calls); got != 1 {
				t.Fatalf("callbacks constructed after failure: got %d, want 1", got)
			}
		})
	}
}

func TestPipelineNativeManagedFailureJoinsProducerWorkers(t *testing.T) {
	for _, mode := range []struct {
		name      string
		jobs      int64
		pipelined bool
	}{
		{"sequential", 1, true},
		{"batched", 3, false},
		{"pipeline", 3, true},
	} {
		for _, stage := range []string{"producer", "emission"} {
			t.Run(mode.name+"/"+stage, func(t *testing.T) {
				failure := errors.New("managed-" + stage + "-failed")
				started, stopped := make(chan struct{}), make(chan struct{})
				// Model an optimizer worker publishing to an occupied result slot.
				results := make(chan int, 1)
				results <- 1
				t.Cleanup(func() {
					// Release a stranded worker too if supervision is accidentally removed.
					select {
					case <-results:
					default:
					}
				})
				worker := AffFn(func(ctx context.Context) (any, error) {
					close(started)
					defer close(stopped)
					select {
					case results <- 2:
						return rt.Value{}, nil
					case <-ctx.Done():
						return nil, context.Cause(ctx)
					}
				})
				emit := rt.Func(func(rt.Value) rt.Value {
					return rt.Box(AffFn(func(context.Context) (any, error) { return nil, failure }))
				})
				produce := rt.Func(func(enqueue rt.Value) rt.Value {
					return rt.Box(AffFn(func(ctx context.Context) (any, error) {
						fork := rt.Apply(Get_Effect_Aff_forkAff(), rt.Box(worker))
						if _, err := runAffSync(rt.Unbox[AffFn](fork), ctx); err != nil {
							return nil, err
						}
						select {
						case <-started:
						case <-ctx.Done():
							return nil, context.Cause(ctx)
						}
						if stage == "producer" {
							return nil, failure
						}
						return runAffSync(rt.Unbox[AffFn](rt.Apply(enqueue, pipelineNativeEntry(1))), ctx)
					}))
				})
				options := rt.RecordDict2("jobs", "pipelined", rt.Int(mode.jobs), rt.Bool(mode.pipelined))
				build := rt.Apply3(Get_Gopurs_Emission_withEmitter(), options, emit, produce)
				err := pipelineNativeAwait(t, "managed build", pipelineNativeStart(build))
				if err == nil || !strings.Contains(err.Error(), failure.Error()) {
					t.Fatalf("got %v, want original build error", err)
				}
				select {
				case <-stopped:
				default:
					t.Fatal("build returned before its producer worker stopped")
				}
				if got := <-results; got != 1 {
					t.Fatalf("result slot was consumed: got %d, want 1", got)
				}
			})
		}
	}
}

func TestPipelineNativeCancelDrainsDetachedWrite(t *testing.T) {
	allowWrite := newPipelineNativeGate()
	started, writeFinished := make(chan struct{}), make(chan struct{})
	workerCancelled := make(chan struct{}, 1)
	callbacks := make(chan []int64, 4)
	emitter := pipelineNativeEmitter(func(ids []int64) AffFn {
		callbacks <- ids
		return func(ctx context.Context) (any, error) {
			callback := make(chan struct{})
			// Model Node.FS.Async.writeFileImpl: the write runs independently of
			// the Aff context and reports completion through a callback.
			go func() {
				close(started)
				<-allowWrite.ch
				close(writeFinished)
				close(callback)
			}()
			select {
			case <-callback:
				return rt.Value{}, nil
			case <-ctx.Done():
				// A nonCanceler cannot stop or join the detached write. Killing
				// this Aff would let the fiber end before the write callback.
				workerCancelled <- struct{}{}
				return nil, context.Cause(ctx)
			}
		}
	})
	pipelineNativeCleanup(t, emitter, allowWrite)
	pipelineNativeRun(t, "enqueue 1", pipelineNativeEnqueue(emitter, 1))
	pipelineNativeRun(t, "enqueue 2", pipelineNativeEnqueue(emitter, 2))
	pipelineNativeSignal(t, "detached write start", started)
	pipelineNativeRun(t, "enqueue pending 3", pipelineNativeEnqueue(emitter, 3))
	cancel := pipelineNativeStart(rt.RecordGet(emitter, "cancel"))
	pipelineNativeBlocked(t, "cancel while detached write has not called back", cancel)
	allowWrite.release()
	if err := pipelineNativeAwait(t, "cancel after write callback", cancel); err != nil {
		t.Fatal(err)
	}
	select {
	case <-writeFinished:
	default:
		t.Fatal("cancel returned before the detached write completed")
	}
	select {
	case <-workerCancelled:
		t.Fatal("cancel interrupted the worker instead of draining its write")
	default:
	}
	err := pipelineNativeAwait(t, "enqueue after cancellation", pipelineNativeStart(pipelineNativeEnqueue(emitter, 4)))
	if err == nil || !strings.Contains(err.Error(), "Go emission cancelled") {
		t.Fatalf("enqueue after cancellation: got %v", err)
	}
	pipelineNativeRun(t, "finish after cancellation", rt.RecordGet(emitter, "finish"))
	if got := len(callbacks); got != 1 {
		t.Fatalf("cancel emitted pending work: got %d callbacks, want 1", got)
	}
}
