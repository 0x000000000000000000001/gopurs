package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

func syncDirectory(path string) error {
	if path == os.Getenv("GOPURS_TEST_PAUSE_DIRECTORY") {
		fmt.Fprintln(os.Stdout, "GOPURS_BEFORE_DIRECTORY_BARRIER")
		for {
			time.Sleep(time.Hour)
		}
	}
	return untracedSyncDirectory(path)
}

func TestOutputBarrierProcessDeath(t *testing.T) {
	if request := os.Getenv("GOPURS_TEST_CHILD_REQUEST"); request != "" {
		data, err := os.ReadFile(request)
		if err != nil {
			t.Fatal(err)
		}
		var req Request
		if err := json.Unmarshal(data, &req); err != nil {
			t.Fatal(err)
		}
		_, err = execute(req)
		t.Fatal("child returned before its forced crash", err)
	}
	f := newFixture(t)
	publish(t, f)
	pointerPath := filepath.Join(f.s.cache, "current.json")
	before, _ := os.ReadFile(pointerPath)
	lockPath := filepath.Join(f.s.cache, "write.lock")
	lockBefore, err := os.Stat(lockPath)
	if err != nil {
		t.Fatal(err)
	}
	f.req.Outputs = nil
	for i := 0; i < 32; i++ {
		name := fmt.Sprintf("New%02d", i)
		source := "stage/" + name + ".go"
		put(t, filepath.Join(f.root, source), "package purescript\n// "+name+"\n")
		f.req.Outputs = append(f.req.Outputs, StagedOutput{Path: "purescript/" + name + ".go", Kind: "module", Module: name, Source: source})
	}
	f.req.Outputs = append(f.req.Outputs, StagedOutput{Path: "Nested.Entry/main/main.go", Kind: "entry", Module: "Nested.Entry", Source: "stage/Main.go"})
	f.req.ExpectedKey = call(t, f.req, "snapshot").Snapshot.Key
	f.req.Operation = "publish"
	request := filepath.Join(f.root, "child-request.json")
	data, err := json.Marshal(f.req)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(request, data, 0600); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestOutputBarrierProcessDeath$")
	cmd.Env = append(os.Environ(), "GOPURS_TEST_CHILD_REQUEST="+request,
		"GOPURS_TEST_PAUSE_DIRECTORY="+filepath.Join(f.s.root, "Nested.Entry/main"))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if cmd.ProcessState == nil {
			_ = cmd.Process.Kill()
			_ = cmd.Wait()
		}
	}()
	scanner := bufio.NewScanner(stdout)
	if !scanner.Scan() || scanner.Text() != "GOPURS_BEFORE_DIRECTORY_BARRIER" {
		t.Fatal("child did not reach the directory barrier", scanner.Text(), scanner.Err())
	}
	// Every writer must be joined before the first (deepest) directory barrier.
	dates := map[string]int64{}
	for _, out := range f.req.Outputs {
		path := filepath.Join(f.s.root, out.Path)
		got, err := os.ReadFile(path)
		want, _ := os.ReadFile(filepath.Join(f.root, out.Source))
		if err != nil || !bytes.Equal(got, want) {
			t.Fatal("barrier preceded completed writer", out.Path, err)
		}
		info, _ := os.Stat(path)
		dates[out.Path] = info.ModTime().UnixNano()
	}
	if got, _ := os.ReadFile(pointerPath); !bytes.Equal(got, before) {
		t.Fatal("unbarriered outputs committed")
	}
	owned, err := f.s.readOwnership()
	if err != nil || len(owned) != 34 {
		t.Fatal("union journal not durable before public writes", err, len(owned))
	}
	if _, err := os.Stat(filepath.Join(f.s.root, "purescript/Main.go")); err != nil {
		t.Fatal("obsolete file deleted before barrier", err)
	}
	if err := cmd.Process.Kill(); err != nil {
		t.Fatal(err)
	}
	if err := cmd.Wait(); err == nil {
		t.Fatal("expected real process termination")
	}
	// Model renames/directory entries lost by a power interruption after the
	// actual SIGKILL. All surviving byte-identical outputs must remain untouched.
	if err := os.Remove(filepath.Join(f.s.root, "purescript/New00.go")); err != nil {
		t.Fatal(err)
	}
	if err := os.RemoveAll(filepath.Join(f.s.root, "Nested.Entry")); err != nil {
		t.Fatal(err)
	}
	lookup(t, f, "miss")
	result := publish(t, f)
	for _, out := range result.Manifest.Outputs {
		path := filepath.Join(f.s.root, out.Path)
		if !fileMatches(path, out.Object) {
			t.Fatal("recovery incomplete", out.Path)
		}
		info, _ := os.Stat(path)
		if out.Path != "purescript/New00.go" && out.Path != "Nested.Entry/main/main.go" && info.ModTime().UnixNano() != dates[out.Path] {
			t.Fatal("surviving output rewritten", out.Path)
		}
	}
	if _, err := os.Stat(filepath.Join(f.s.root, "purescript/Main.go")); !os.IsNotExist(err) {
		t.Fatal("obsolete file not removed after recovery")
	}
	lockAfter, err := os.Stat(lockPath)
	if err != nil || !os.SameFile(lockBefore, lockAfter) {
		t.Fatal("lock inode changed", err)
	}
	lookup(t, f, "hit")
	t.Log("SIGKILL before first output-directory barrier recovered 33 outputs; unchanged outputs and lock inode retained")
}
