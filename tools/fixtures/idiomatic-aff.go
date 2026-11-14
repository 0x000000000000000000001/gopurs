package fixture

import (
	"context"
	"errors"
	"sync/atomic"
)

var calls atomic.Int64
var started atomic.Bool
var stopped atomic.Bool

func Hello(name string) (string, error) {
	calls.Add(1)
	if name == "error" {
		return "", errors.New("expected error")
	}
	return "hello " + name, nil
}

func Notify() error                     { return nil }
func NoError() string                   { return "plain result" }
func Fire()                             {}
func ReadCalls() int64                  { return calls.Load() }
func Label() string                     { return "effect without Unit" }
func RunCallback(callback func(string)) { callback("typed") }

func Pending(ctx context.Context) (string, error) {
	started.Store(true)
	<-ctx.Done()
	stopped.Store(true)
	return "", ctx.Err()
}

func IsStarted() bool { return started.Load() }
func IsStopped() bool { return stopped.Load() }
