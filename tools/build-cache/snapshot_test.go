package main

import (
	"errors"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"
)

func TestSnapshotPreservesErrorPriority(t *testing.T) {
	f := newFixture(t)
	base := f.req.Spec
	base.Output = f.s.root
	// Every compiler rejection must outrank this independent input error.
	base.Inputs = []Input{{Name: "invalid", Kind: "unknown", Path: "compiler"}}
	cases := []struct {
		name     string
		compiler []Input
		want     string
	}{
		{"compiler-input-error", []Input{{Name: "", Kind: "file", Path: "compiler"}}, "duplicate/empty input name"},
		{"missing-compiler", []Input{{Name: "compiler", Kind: "file", Path: "missing"}}, "compiler resource absent/not a file: " + filepath.Join(f.root, "missing")},
		{"non-file-compiler", []Input{{Name: "compiler", Kind: "directory", Path: "src"}}, "compiler resource absent/not a file: " + filepath.Join(f.root, "src")},
		{"compiler-required", []Input{{Name: "resource", Kind: "file", Path: "compiler"}}, "compiler artifact is required"},
		{"no-compiler-resources", nil, "compiler artifact is required"},
		{"input-error", base.Compiler, "unknown input kind"},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			spec := base
			spec.Compiler = test.compiler
			_, err := snapshot(spec)
			if err == nil || err.Error() != test.want {
				t.Fatalf("error priority changed: got %v, want %q", err, test.want)
			}
		})
	}
}

func TestSnapshotJoinsIndependentObservationsOnFailure(t *testing.T) {
	for _, blocked := range []string{"compiler", "inputs"} {
		t.Run(blocked, func(t *testing.T) {
			spec := Spec{Compiler: []Input{{Name: "compiler"}}, Inputs: []Input{{Name: "inputs"}}}
			started := make(chan struct{}, 2)
			unblockedFinished := make(chan struct{})
			release := make(chan struct{})
			defer close(release)
			compilerFailure := errors.New("compiler observation failed")
			inputFailure := errors.New("input observation failed")
			var entered, finished atomic.Int32
			observe := func(_ string, items []Input) ([]Fingerprint, error) {
				entered.Add(1)
				started <- struct{}{}
				if items[0].Name == blocked {
					<-release
				} else {
					close(unblockedFinished)
				}
				finished.Add(1)
				if items[0].Name == "compiler" {
					return nil, compilerFailure
				}
				return nil, inputFailure
			}
			result := make(chan error, 1)
			go func() {
				_, err := snapshotWithFingerprints(spec, observe)
				result <- err
			}()
			wait := func(signal <-chan struct{}) {
				t.Helper()
				select {
				case <-signal:
				case <-time.After(time.Second):
					t.Fatal("independent observations did not both start")
				}
			}
			wait(started)
			wait(started)
			wait(unblockedFinished)
			select {
			case err := <-result:
				t.Fatalf("returned while %s observation was still running: %v", blocked, err)
			case <-time.After(10 * time.Millisecond):
			}
			// Sending releases exactly this observation; the deferred close also
			// releases it if an earlier assertion aborts the test.
			release <- struct{}{}
			if err := <-result; err != compilerFailure {
				t.Fatalf("compiler error identity/priority changed: %v", err)
			}
			if entered.Load() != 2 || finished.Load() != 2 {
				t.Fatalf("observations not joined exactly once: %d entered, %d finished", entered.Load(), finished.Load())
			}
		})
	}
}
