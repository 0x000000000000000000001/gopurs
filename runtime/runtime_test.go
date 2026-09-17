package gopurs_runtime

import (
	"runtime"
	"sync"
	"testing"
)

func TestConcurrentClosuresSurviveGC(t *testing.T) {
	var workers sync.WaitGroup
	for worker := int64(0); worker < 8; worker++ {
		workers.Add(1)
		go func(seed int64) {
			defer workers.Done()
			for iteration := int64(0); iteration < 128; iteration++ {
				captured := seed*1000 + iteration
				unary := Func(func(arg Value) Value { return Int(captured + arg.IntVal) })
				binary := Func2(func(a, b Value) Value { return Int(captured + a.IntVal + b.IntVal) })
				if iteration%32 == 0 {
					runtime.GC()
				}
				if got := Apply(unary, Int(7)).IntVal; got != captured+7 {
					t.Errorf("unary closure: got %d, want %d", got, captured+7)
				}
				// Partial application must retain both the captured value and argument.
				partial := Apply(binary, Int(3))
				if got := Apply(partial, Int(4)).IntVal; got != captured+7 {
					t.Errorf("partial closure: got %d, want %d", got, captured+7)
				}
			}
		}(worker)
	}
	workers.Wait()
}
