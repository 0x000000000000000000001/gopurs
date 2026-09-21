package purescript

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	rt "gopurs/output/gopurs_runtime"
)

const preparationNativeTimeout = 3 * time.Second

type preparationNativeResult struct {
	value any
	err   error
}

func preparationNativeAff(jobs int64, thunks []rt.Value) rt.Value {
	return rt.Apply2(Get_Gopurs_Preparation_runPreparationJobs(), rt.Int(jobs), rt.Array(thunks))
}

func preparationNativeStart(aff rt.Value) <-chan preparationNativeResult {
	done := make(chan preparationNativeResult, 1)
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		value, err := runAffSync(rt.Unbox[AffFn](aff), ctx)
		done <- preparationNativeResult{value, err}
	}()
	return done
}

func preparationNativeValues(t *testing.T, done <-chan preparationNativeResult) []int64 {
	t.Helper()
	select {
	case result := <-done:
		if result.err != nil {
			t.Fatal(result.err)
		}
		value := rt.Box(result.value)
		values := make([]int64, rt.ArrayLength(value))
		for i := range values {
			values[i] = rt.ArrayAccess(value, i).IntVal
		}
		return values
	case <-time.After(preparationNativeTimeout):
		t.Fatal("preparation did not finish")
		return nil
	}
}

func preparationNativeReceive(t *testing.T, label string, events <-chan int) int {
	t.Helper()
	select {
	case index := <-events:
		return index
	case <-time.After(preparationNativeTimeout):
		t.Fatalf("%s was not observed; pure computations may not run concurrently", label)
		return -1
	}
}

func TestPreparationNativeDefersAndReruns(t *testing.T) {
	var calls [5]atomic.Int32
	thunks := make([]rt.Value, len(calls))
	for i := range thunks {
		index := i
		thunks[i] = rt.Func(func(rt.Value) rt.Value {
			calls[index].Add(1)
			return rt.Int(int64(index))
		})
	}
	aff := preparationNativeAff(3, thunks)
	for i := range calls {
		if got := calls[i].Load(); got != 0 {
			t.Fatalf("thunk %d ran while constructing Aff: %d calls", i, got)
		}
	}
	for run := int32(1); run <= 2; run++ {
		got := preparationNativeValues(t, preparationNativeStart(aff))
		if !reflect.DeepEqual(got, []int64{0, 1, 2, 3, 4}) {
			t.Fatalf("run %d results: %v", run, got)
		}
		for i := range calls {
			if got := calls[i].Load(); got != run {
				t.Fatalf("run %d: thunk %d has %d calls", run, i, got)
			}
		}
	}
	if got := preparationNativeValues(t, preparationNativeStart(preparationNativeAff(3, nil))); len(got) != 0 {
		t.Fatalf("empty input returned %v", got)
	}
}

func preparationNativeExercise(t *testing.T, jobs int64, width, count int) {
	t.Helper()
	started, finished := make(chan int, count), make(chan int, count)
	gates := make([]chan struct{}, count)
	releases := make([]sync.Once, count)
	for i := range gates {
		gates[i] = make(chan struct{})
	}
	release := func(index int) { releases[index].Do(func() { close(gates[index]) }) }
	t.Cleanup(func() {
		// A failed assertion must not leave a pure thunk suspended forever.
		for i := range gates {
			release(i)
		}
	})
	var active, peak atomic.Int32
	thunks := make([]rt.Value, count)
	for i := range thunks {
		index := i
		thunks[i] = rt.Func(func(rt.Value) rt.Value {
			current := active.Add(1)
			for previous := peak.Load(); current > previous; previous = peak.Load() {
				if peak.CompareAndSwap(previous, current) {
					break
				}
			}
			defer func() {
				active.Add(-1)
				finished <- index
			}()
			started <- index
			<-gates[index]
			return rt.Int(int64(index*10 + 7))
		})
	}
	// Construction is deliberately included in this goroutine: a missing
	// defer then fails the overlap assertion rather than hanging the test.
	done := make(chan preparationNativeResult, 1)
	go func() {
		done <- <-preparationNativeStart(preparationNativeAff(jobs, thunks))
	}()
	seen := make(map[int]bool)
	receiveStart := func() int {
		index := preparationNativeReceive(t, "computation start", started)
		if index < 0 || index >= count || seen[index] {
			t.Fatalf("unexpected or repeated computation %d", index)
		}
		seen[index] = true
		return index
	}
	complete := func(index int) {
		release(index)
		if got := preparationNativeReceive(t, "computation completion", finished); got != index {
			t.Fatalf("completion: got %d, want %d", got, index)
		}
	}
	// Observe whichever computations the scheduler actually starts. Neither
	// their indices nor the worker partition sizes are part of this contract.
	initial := make([]int, width)
	for i := range initial {
		initial[i] = receiveStart()
	}
	if got := active.Load(); got != int32(width) {
		t.Fatalf("active pure computations: got %d, want %d", got, width)
	}
	select {
	case index := <-started:
		t.Fatalf("computation %d started beyond jobs=%d while all workers were blocked", index, jobs)
	case <-done:
		t.Fatal("preparation completed while its computations were blocked")
	case <-time.After(25 * time.Millisecond):
	}
	// Force the initially overlapping computations to finish out of input
	// order, then release every later computation as it arrives.
	sort.Sort(sort.Reverse(sort.IntSlice(initial)))
	for _, index := range initial {
		complete(index)
	}
	for completed := width; completed < count; completed++ {
		complete(receiveStart())
	}
	got := preparationNativeValues(t, done)
	want := make([]int64, count)
	for i := range want {
		want[i] = int64(i*10 + 7)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("results: got %v, want %v", got, want)
	}
	if got := peak.Load(); got != int32(width) {
		t.Fatalf("peak active computations: got %d, want %d", got, width)
	}
	if got := active.Load(); got != 0 {
		t.Fatalf("%d computations still active after completion", got)
	}
}

func TestPreparationNativeParallelBoundAndOrder(t *testing.T) {
	t.Run("jobs=3", func(t *testing.T) {
		preparationNativeExercise(t, 3, 3, 107)
	})
	t.Run("jobs-clamped-to-8", func(t *testing.T) {
		preparationNativeExercise(t, 999, 8, 145)
	})
}

func TestPreparationNativeSequential(t *testing.T) {
	for _, jobs := range []int64{-2, 0, 1} {
		t.Run(fmt.Sprintf("jobs=%d", jobs), func(t *testing.T) {
			preparationNativeExercise(t, jobs, 1, 4)
		})
	}
}
