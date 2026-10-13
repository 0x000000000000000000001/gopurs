package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"
)

// This recorder exists only in an instrumented private copy. It wraps the real
// syncDirectory; production has no test hooks or global mutable callbacks.
var barrierProbe struct {
	sync.Mutex
	root  string
	calls []string
	fail  string
}

func syncDirectory(path string) error {
	barrierProbe.Lock()
	root, fail := barrierProbe.root, barrierProbe.fail
	if root != "" && (path == root || strings.HasPrefix(path, root+string(os.PathSeparator))) {
		name, _ := filepath.Rel(root, path)
		barrierProbe.calls = append(barrierProbe.calls, filepath.ToSlash(name))
	}
	barrierProbe.Unlock()
	if fail != "" && path == fail {
		return fmt.Errorf("injected output directory barrier failure")
	}
	return untracedSyncDirectory(path)
}

func recordBarriers(root, fail string) {
	barrierProbe.Lock()
	defer barrierProbe.Unlock()
	barrierProbe.root, barrierProbe.fail, barrierProbe.calls = root, fail, nil
}

func recordedBarriers() []string {
	barrierProbe.Lock()
	defer barrierProbe.Unlock()
	return append([]string{}, barrierProbe.calls...)
}

func multiOutputFixture(t *testing.T) fixture {
	f := newFixture(t)
	for i := 0; i < 32; i++ {
		name := fmt.Sprintf("Mod%02d", i)
		source := "stage/" + name + ".go"
		put(t, filepath.Join(f.root, source), "package purescript\n// "+name+"\n")
		f.req.Outputs = append(f.req.Outputs, StagedOutput{Path: "purescript/" + name + ".go", Kind: "module", Module: name, Source: source})
	}
	f.req.Outputs = append(f.req.Outputs, StagedOutput{Path: "Nested.Entry/main/main.go", Kind: "entry", Module: "Nested.Entry", Source: "stage/Main.go"})
	return f
}

func TestOutputDirectoryBarriers(t *testing.T) {
	f := multiOutputFixture(t)
	snap := call(t, f.req, "snapshot").Snapshot
	m, err := f.s.stage(f.req, *snap, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := f.s.writeOwnership(m.Outputs); err != nil {
		t.Fatal(err)
	}
	for _, mode := range []string{"new", "identical-recovery"} {
		recordBarriers(f.s.root, "")
		err := f.s.writeOutputs(m.Outputs)
		calls := recordedBarriers()
		recordBarriers("", "")
		if err != nil {
			t.Fatal(err)
		}
		counts := map[string]int{}
		positions := map[string]int{}
		for i, path := range calls {
			counts[path]++
			positions[path] = i
		}
		want := map[string]int{"purescript": 1, "Nested.Entry/main": 1, "Nested.Entry": 1, ".": 1}
		t.Logf("%s: %d output files; directory syncs=%v", mode, len(m.Outputs), counts)
		if !reflect.DeepEqual(counts, want) {
			t.Errorf("%s: want one barrier per directory including recovery ancestors: got %v", mode, counts)
		}
		if positions["Nested.Entry/main"] >= positions["Nested.Entry"] || positions["Nested.Entry"] >= positions["."] {
			t.Errorf("new directory entries must be durable bottom-up: %v", calls)
		}
		for _, output := range m.Outputs {
			if !fileMatches(filepath.Join(f.s.root, output.Path), output.Object) {
				t.Fatal(output.Path)
			}
		}
	}
}

func TestOutputBarrierFailureRecovery(t *testing.T) {
	f := newFixture(t)
	publish(t, f)
	pointerPath := filepath.Join(f.s.cache, "current.json")
	before, _ := os.ReadFile(pointerPath)
	lockBefore, err := os.Stat(filepath.Join(f.s.cache, "write.lock"))
	if err != nil {
		t.Fatal(err)
	}
	// An obsolete file must survive a failure before all output directories have
	// reached their barriers. The union journal must still cover both versions.
	f.req.Outputs = nil
	for i := 0; i < 16; i++ {
		name := fmt.Sprintf("New%02d", i)
		source := "stage/" + name + ".go"
		put(t, filepath.Join(f.root, source), "package purescript\n// "+name+"\n")
		f.req.Outputs = append(f.req.Outputs, StagedOutput{Path: "purescript/" + name + ".go", Kind: "module", Module: name, Source: source})
	}
	f.req.Operation = "snapshot"
	f.req.ExpectedKey = call(t, f.req, "snapshot").Snapshot.Key
	f.req.Operation = "publish"
	recordBarriers(f.s.root, filepath.Join(f.s.root, "purescript"))
	_, err = execute(f.req)
	calls := recordedBarriers()
	recordBarriers("", "")
	if err == nil || !strings.Contains(err.Error(), "injected output directory barrier failure") {
		t.Fatal("did not report barrier failure", err)
	}
	if got, _ := os.ReadFile(pointerPath); !bytes.Equal(got, before) {
		t.Fatal("failed barrier committed a manifest")
	}
	if _, err := os.Stat(filepath.Join(f.s.root, "purescript/Main.go")); err != nil {
		t.Fatal("deleted obsolete output before successful barrier", err)
	}
	owned, err := f.s.readOwnership()
	if err != nil || len(owned) != 17 {
		t.Fatal("incomplete ownership union", len(owned), err)
	}
	dates := map[string]int64{}
	for _, item := range f.req.Outputs {
		path := filepath.Join(f.s.root, item.Path)
		info, err := os.Stat(path)
		if err != nil {
			t.Fatal("returned before every writer finished", err)
		}
		dates[item.Path] = info.ModTime().UnixNano()
	}
	// Model one uncommitted rename being lost after a crash; all other output
	// bytes survived but still require a directory barrier on the retry.
	if err := os.Remove(filepath.Join(f.s.root, "purescript/New00.go")); err != nil {
		t.Fatal(err)
	}
	recordBarriers(f.s.root, "")
	response := publish(t, f)
	retry := recordedBarriers()
	recordBarriers("", "")
	for _, item := range response.Manifest.Outputs {
		path := filepath.Join(f.s.root, item.Path)
		if !fileMatches(path, item.Object) {
			t.Fatal("incomplete recovery", item.Path)
		}
		info, _ := os.Stat(path)
		if item.Path != "purescript/New00.go" && info.ModTime().UnixNano() != dates[item.Path] {
			t.Fatal("recovery rewrote identical file", item.Path)
		}
	}
	if _, err := os.Stat(filepath.Join(f.s.root, "purescript/Main.go")); !os.IsNotExist(err) {
		t.Fatal("obsolete output survived successful retry")
	}
	lockAfter, err := os.Stat(filepath.Join(f.s.cache, "write.lock"))
	if err != nil || !os.SameFile(lockBefore, lockAfter) {
		t.Fatal("lock inode replaced", err)
	}
	lookup(t, f, "hit")
	t.Logf("failure barriers=%v; retry barriers=%v", calls, retry)
}
