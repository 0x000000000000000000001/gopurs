package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"
)

type fixture struct {
	root string
	req  Request
	s    store
}

func put(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func newFixture(t *testing.T) fixture {
	t.Helper()
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	for name, content := range map[string]string{
		"compiler": "local compiler bytes", "ffi-parser": "parser bytes", "runtime": "runtime bytes",
		"output/Main/corefn.json": `{"typeTable":[],"moduleName":["Main"]}`,
		"src/Main.go":             "func Main() {}", "stage/Main.go": "package purescript\n// generated A\n",
	} {
		put(t, filepath.Join(root, name), content)
	}
	req := Request{Schema: Schema, Operation: "snapshot", Spec: Spec{
		Workspace: root, Output: "output",
		Compiler: []Input{{Name: "compiler", Kind: "file", Path: "compiler"}, {Name: "parser", Kind: "file", Path: "ffi-parser"}, {Name: "runtime", Kind: "file", Path: "runtime"}},
		Options:  []Setting{{Name: "main", Value: "Main"}, {Name: "rewriteLimit", Value: "10000"}},
		Inputs:   []Input{{Name: "tast", Kind: "tast", Path: "output"}, {Name: "ffi", Kind: "file", Path: "src/Main.go"}, {Name: "preferred-ffi", Kind: "file", Path: "src/Preferred.go"}, {Name: "directives", Kind: "file", Path: "directives.txt"}, {Name: "ffi-roots", Kind: "directory", Path: "src"}},
	}, Outputs: []StagedOutput{{Path: "purescript/Main.go", Kind: "module", Module: "Main", Source: "stage/Main.go"}}}
	f := fixture{root: root, req: req, s: store{root: filepath.Join(root, "output"), cache: filepath.Join(root, "output/.gopurs-cache")}}
	return f
}

func call(t *testing.T, req Request, status string) Response {
	t.Helper()
	data, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var result Response
	if err := json.Unmarshal([]byte(Exchange(string(data))), &result); err != nil {
		t.Fatal(err)
	}
	if result.Status != status {
		t.Fatalf("want %s, got %+v", status, result)
	}
	return result
}

func publish(t *testing.T, f fixture) Response {
	t.Helper()
	req := f.req
	req.Operation = "snapshot"
	req.ExpectedKey = call(t, req, "snapshot").Snapshot.Key
	req.Operation = "publish"
	return call(t, req, "published")
}

func lookup(t *testing.T, f fixture, status string) Response {
	t.Helper()
	req := f.req
	req.Operation = "lookup"
	return call(t, req, status)
}

func TestContentInvalidation(t *testing.T) {
	mutations := map[string]func(*testing.T, fixture){
		"tast-same-mtime": func(t *testing.T, f fixture) {
			path := filepath.Join(f.root, "output/Main/corefn.json")
			info, _ := os.Stat(path)
			put(t, path, `{"typeTable":[],"moduleName":["Maim"]}`)
			if err := os.Chtimes(path, info.ModTime(), info.ModTime()); err != nil {
				t.Fatal(err)
			}
		},
		"tast-add": func(t *testing.T, f fixture) { put(t, filepath.Join(f.root, "output/New/corefn.json"), "{}") },
		"tast-remove": func(t *testing.T, f fixture) {
			if err := os.Remove(filepath.Join(f.root, "output/Main/corefn.json")); err != nil {
				t.Fatal(err)
			}
		},
		"tast-rename": func(t *testing.T, f fixture) {
			if err := os.Rename(filepath.Join(f.root, "output/Main"), filepath.Join(f.root, "output/Other")); err != nil {
				t.Fatal(err)
			}
		},
		"ffi-body": func(t *testing.T, f fixture) {
			put(t, filepath.Join(f.root, "src/Main.go"), "func Main() { panic(1) }")
		},
		"ffi-absence": func(t *testing.T, f fixture) { put(t, filepath.Join(f.root, "src/Preferred.go"), "func Main() {}") },
		"ffi-removal": func(t *testing.T, f fixture) {
			if err := os.Remove(filepath.Join(f.root, "src/Main.go")); err != nil {
				t.Fatal(err)
			}
		},
		"resolution-root": func(t *testing.T, f fixture) {
			if err := os.Mkdir(filepath.Join(f.root, "src/v2"), 0755); err != nil {
				t.Fatal(err)
			}
		},
		"directives":            func(t *testing.T, f fixture) { put(t, filepath.Join(f.root, "directives.txt"), "Main.main never") },
		"compiler-local-change": func(t *testing.T, f fixture) { put(t, filepath.Join(f.root, "compiler"), "dirty build") },
		"parser":                func(t *testing.T, f fixture) { put(t, filepath.Join(f.root, "ffi-parser"), "new parser") },
		"runtime":               func(t *testing.T, f fixture) { put(t, filepath.Join(f.root, "runtime"), "new runtime") },
	}
	for name, change := range mutations {
		t.Run(name, func(t *testing.T) {
			f := newFixture(t)
			publish(t, f)
			lookup(t, f, "hit")
			change(t, f)
			if result := lookup(t, f, "miss"); result.Reason != "inputs-changed" {
				t.Fatal(result)
			}
		})
	}
	t.Run("options", func(t *testing.T) {
		f := newFixture(t)
		publish(t, f)
		f.req.Spec.Options[0].Value = "Other"
		lookup(t, f, "miss")
	})
}

func TestSymlinkAndCanonicalOrdering(t *testing.T) {
	f := newFixture(t)
	put(t, filepath.Join(f.root, "first.go"), "same bytes")
	put(t, filepath.Join(f.root, "second.go"), "same bytes")
	link := filepath.Join(f.root, "linked.go")
	if err := os.Symlink("first.go", link); err != nil {
		t.Fatal(err)
	}
	f.req.Spec.Inputs = append(f.req.Spec.Inputs, Input{Name: "linked", Kind: "file", Path: link})
	publish(t, f)
	for i, j := 0, len(f.req.Spec.Inputs)-1; i < j; i, j = i+1, j-1 {
		f.req.Spec.Inputs[i], f.req.Spec.Inputs[j] = f.req.Spec.Inputs[j], f.req.Spec.Inputs[i]
	}
	f.req.Spec.Options[0], f.req.Spec.Options[1] = f.req.Spec.Options[1], f.req.Spec.Options[0]
	lookup(t, f, "hit")
	if err := os.Remove(link); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("second.go", link); err != nil {
		t.Fatal(err)
	}
	lookup(t, f, "miss")
}

func TestIntegrityAndVersionMisses(t *testing.T) {
	for _, kind := range []string{"absent", "pointer-version", "manifest-version", "manifest-object", "output-object", "output-file", "missing-output", "missing-go-mod", "incomplete-journal"} {
		t.Run(kind, func(t *testing.T) {
			f := newFixture(t)
			if kind == "absent" {
				lookup(t, f, "miss")
				return
			}
			result := publish(t, f)
			manifest := *result.Manifest
			current := filepath.Join(f.s.cache, "current.json")
			switch kind {
			case "pointer-version":
				put(t, current, `{"format":"gopurs-build-cache","schema":999,"manifest":{}}`)
			case "manifest-version":
				manifest.Schema = 999
				data, _ := json.Marshal(manifest)
				obj, err := f.s.putObject(data)
				if err != nil {
					t.Fatal(err)
				}
				data, _ = json.Marshal(pointer{Format: Format, Schema: Schema, Manifest: obj})
				put(t, current, string(data))
			case "manifest-object":
				var p pointer
				data, _ := os.ReadFile(current)
				if err := json.Unmarshal(data, &p); err != nil {
					t.Fatal(err)
				}
				put(t, filepath.Join(f.s.cache, "objects", p.Manifest.SHA256), "corrupt")
			case "output-object":
				put(t, filepath.Join(f.s.cache, "objects", manifest.Outputs[0].Object.SHA256), "corrupt")
			case "output-file":
				put(t, filepath.Join(f.s.root, "purescript/Main.go"), "external edit")
			case "missing-output":
				if err := os.Remove(filepath.Join(f.s.root, "purescript/Main.go")); err != nil {
					t.Fatal(err)
				}
			case "missing-go-mod":
				if err := os.Remove(filepath.Join(f.s.root, "go.mod")); err != nil {
					t.Fatal(err)
				}
			case "incomplete-journal":
				if err := f.s.writeOwnership(nil); err != nil {
					t.Fatal(err)
				}
			}
			lookup(t, f, "miss")
			publish(t, f)
			lookup(t, f, "hit")
		})
	}
}

func TestOwnershipTidyAndWriteElision(t *testing.T) {
	f := newFixture(t)
	put(t, filepath.Join(f.root, "stage/Old.go"), "package purescript\n// obsolete\n")
	f.req.Outputs = append(f.req.Outputs, StagedOutput{Path: "purescript/Old_ffi.go", Kind: "ffi", Module: "Old", Source: "stage/Old.go"})
	publish(t, f)
	owned := filepath.Join(f.s.root, "purescript/Main.go")
	oldTime := time.Unix(10, 0)
	if err := os.Chtimes(owned, oldTime, oldTime); err != nil {
		t.Fatal(err)
	}
	goMod := "module gopurs/output\n\ngo 1.22\n\nrequire example.com/dependency v1.0.0\nreplace example.com/dependency => ../local\n"
	put(t, filepath.Join(f.s.root, "go.mod"), goMod)
	put(t, filepath.Join(f.s.root, "go.sum"), "external sums\n")
	put(t, filepath.Join(f.s.root, "purescript/user.go"), "// user-owned\n")
	lookup(t, f, "hit")
	f.req.Outputs = f.req.Outputs[:1]
	publish(t, f)
	info, err := os.Stat(owned)
	if err != nil || !info.ModTime().Equal(oldTime) {
		t.Fatalf("unchanged output rewritten: %v %v", info, err)
	}
	if _, err := os.Stat(filepath.Join(f.s.root, "purescript/Old_ffi.go")); !os.IsNotExist(err) {
		t.Fatal("obsolete FFI remains")
	}
	for path, want := range map[string]string{"go.mod": goMod, "go.sum": "external sums\n", "purescript/user.go": "// user-owned\n"} {
		got, err := os.ReadFile(filepath.Join(f.s.root, path))
		if err != nil || string(got) != want {
			t.Fatalf("external %s modified: %s %v", path, got, err)
		}
	}
}

func TestFailedAndStalePublication(t *testing.T) {
	for _, kind := range []string{"stale-inputs", "missing-stage", "duplicate-output", "path-escape", "output-symlink", "modified-obsolete"} {
		t.Run(kind, func(t *testing.T) {
			f := newFixture(t)
			publish(t, f)
			before, _ := os.ReadFile(filepath.Join(f.s.cache, "current.json"))
			req := f.req
			req.ExpectedKey = call(t, req, "snapshot").Snapshot.Key
			req.Operation = "publish"
			want := "error"
			switch kind {
			case "stale-inputs":
				put(t, filepath.Join(f.root, "compiler"), "new")
				want = "miss"
			case "missing-stage":
				req.Outputs[0].Source = "no-file"
			case "duplicate-output":
				req.Outputs = append(req.Outputs, req.Outputs[0])
			case "path-escape":
				req.Outputs[0].Path = "../external.go"
			case "output-symlink":
				if err := os.Remove(filepath.Join(f.s.root, "purescript/Main.go")); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(filepath.Join(f.root, "compiler"), filepath.Join(f.s.root, "purescript/Main.go")); err != nil {
					t.Fatal(err)
				}
			case "modified-obsolete":
				put(t, filepath.Join(f.s.root, "purescript/Main.go"), "external edit")
				req.Outputs = nil
			}
			call(t, req, want)
			after, _ := os.ReadFile(filepath.Join(f.s.cache, "current.json"))
			if !bytes.Equal(before, after) {
				t.Fatal("failed build published a manifest")
			}
		})
	}
}

func TestCacheResetRetainsOwnership(t *testing.T) {
	f := newFixture(t)
	publish(t, f)
	if err := os.RemoveAll(filepath.Join(f.s.cache, "objects")); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(filepath.Join(f.s.cache, "current.json")); err != nil {
		t.Fatal(err)
	}
	lookup(t, f, "miss")
	f.req.Outputs = nil
	publish(t, f)
	if _, err := os.Stat(filepath.Join(f.s.root, "purescript/Main.go")); !os.IsNotExist(err) {
		t.Fatal("cache reset lost ownership")
	}
}

func TestMissingJournalRecoversCommittedOwnership(t *testing.T) {
	f := newFixture(t)
	publish(t, f)
	if err := os.Remove(filepath.Join(f.s.root, ".gopurs-ownership.json")); err != nil {
		t.Fatal(err)
	}
	lookup(t, f, "miss")
	f.req.Outputs = nil
	publish(t, f)
	lookup(t, f, "hit")
	if _, err := os.Stat(filepath.Join(f.s.root, "purescript/Main.go")); !os.IsNotExist(err) {
		t.Fatal("valid committed manifest did not restore obsolete-file ownership")
	}
}

func TestWorkspaceIsolation(t *testing.T) {
	first, second := newFixture(t), newFixture(t)
	a := publish(t, first)
	b := publish(t, second)
	if a.Snapshot.Key == b.Snapshot.Key {
		t.Fatal("workspaces share build key")
	}
	if a.Snapshot.Compiler != b.Snapshot.Compiler {
		t.Fatal("identical compiler bytes should share identity")
	}
	data, _ := json.Marshal(a.Manifest)
	obj, err := second.s.putObject(data)
	if err != nil {
		t.Fatal(err)
	}
	data, _ = json.Marshal(pointer{Format: Format, Schema: Schema, Manifest: obj})
	put(t, filepath.Join(second.s.cache, "current.json"), string(data))
	lookup(t, second, "miss")
	lookup(t, first, "hit")
}

func TestWireArtifacts(t *testing.T) {
	valid := `["ctor","Example",[["string","d80000610000dc00"],["number","8000000000000000"],["number","7ff8000000000042"],["int","-9223372036854775808"],["record",[["0061",["boolean",true]]]]]]`
	if err := validateWire([]byte(valid)); err != nil {
		t.Fatal(err)
	}
	for _, invalid := range []string{`null`, `["unknown"]`, `["string","d80"]`, `["string",4]`, `["int",42]`, `["int","01"]`, `["number","NaN"]`, `["int","9223372036854775808"]`, `["ctor","",[]]`, `["record",[["0061",["unit"]],["0061",["unit"]]]]`, `["array",{}]`} {
		if validateWire([]byte(invalid)) == nil {
			t.Fatalf("accepted invalid wire: %s", invalid)
		}
	}
	f := newFixture(t)
	put(t, filepath.Join(f.root, "stage/value.json"), valid)
	mod := StagedModule{Name: "Main", Prepared: Key("prepared"), OptimizerEnvironment: Key("pbo-env"), EmitterEnvironment: Key("emit-env"), CallerDemands: Key("demands")}
	for _, codec := range moduleCodecs {
		mod.Parts = append(mod.Parts, StagedArtifact{Codec: codec, Source: "stage/value.json"})
	}
	f.req.Modules = []StagedModule{mod}
	f.req.Diagnostics = &StagedArtifact{Codec: "gopurs/diagnostics-v1", Source: "stage/value.json"}
	result := publish(t, f)
	lookup(t, f, "hit")
	part := result.Manifest.Modules[0].Parts[0]
	data, err := f.s.readObject(part.Object)
	if err != nil || string(data) != valid {
		t.Fatal("wire bytes changed")
	}
	put(t, filepath.Join(f.s.cache, "objects", part.Object.SHA256), `["unit"]`)
	lookup(t, f, "miss")
	f.req.Modules[0].Parts[0].Codec = "v8/serialize"
	req := f.req
	req.ExpectedKey = call(t, req, "snapshot").Snapshot.Key
	req.Operation = "publish"
	call(t, req, "error")
}

func TestMalformedRequests(t *testing.T) {
	for _, request := range []string{`{}`, `{"schema":99}`, `{"schema":1,"unknown":true}`, `{} {}`, `{`} {
		var response Response
		if err := json.Unmarshal([]byte(Exchange(request)), &response); err != nil {
			t.Fatal(err)
		}
		if response.Status != "error" {
			t.Fatal(response)
		}
	}
	f := newFixture(t)
	f.req.Spec.Inputs = nil
	call(t, f.req, "error")
	f = newFixture(t)
	f.req.Spec.Compiler = nil
	call(t, f.req, "error")
	f = newFixture(t)
	f.req.Spec.Options = append(f.req.Spec.Options, f.req.Spec.Options[0])
	call(t, f.req, "error")
}

// A child provides real process death/OS-lock coverage without production test
// switches. Its marker is sent only after the journal and first file are durable.
func TestCacheProcess(t *testing.T) {
	mode := os.Getenv("CACHE_TEST_CHILD")
	if mode == "" {
		return
	}
	data, err := os.ReadFile(os.Getenv("CACHE_TEST_REQUEST"))
	if err != nil {
		t.Fatal(err)
	}
	if mode == "exchange" {
		fmt.Print(Exchange(string(data)))
		os.Exit(0)
	}
	var req Request
	if err := json.Unmarshal(data, &req); err != nil {
		t.Fatal(err)
	}
	root := filepath.Join(req.Spec.Workspace, "output")
	s := store{root: root, cache: filepath.Join(root, ".gopurs-cache")}
	lock, err := os.OpenFile(filepath.Join(s.cache, "write.lock"), os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer lock.Close()
	if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_EX); err != nil {
		t.Fatal(err)
	}
	owned, err := s.readOwnership()
	if err != nil {
		t.Fatal(err)
	}
	content := []byte("// partial publication\n")
	out := OwnedOutput{Path: "purescript/Partial.go", Kind: "module", Module: "Partial", Object: Object{SHA256: digest(content), Bytes: int64(len(content))}}
	if err := s.writeOwnership(mergeOwnership(owned, []OwnedOutput{out})); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(root, "purescript"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := atomicWrite(filepath.Join(root, out.Path), content, 0644); err != nil {
		t.Fatal(err)
	}
	fmt.Print("ready\n")
	time.Sleep(time.Hour)
}

func child(t *testing.T, f fixture, mode string, req Request) *exec.Cmd {
	t.Helper()
	data, _ := json.Marshal(req)
	file := filepath.Join(t.TempDir(), "request.json")
	put(t, file, string(data))
	cmd := exec.Command(os.Args[0], "-test.run=^TestCacheProcess$")
	cmd.Env = append(os.Environ(), "CACHE_TEST_CHILD="+mode, "CACHE_TEST_REQUEST="+file)
	return cmd
}

func TestInterruptedPublicationAndKernelLock(t *testing.T) {
	for _, firstBuild := range []bool{true, false} {
		t.Run(fmt.Sprint(firstBuild), func(t *testing.T) {
			f := newFixture(t)
			call(t, f.req, "snapshot")
			if !firstBuild {
				publish(t, f)
			}
			cmd := child(t, f, "partial", f.req)
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				t.Fatal(err)
			}
			var stderr bytes.Buffer
			cmd.Stderr = &stderr
			if err := cmd.Start(); err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { _ = cmd.Process.Kill() })
			ready := make([]byte, 6)
			if _, err := io.ReadFull(stdout, ready); err != nil || string(ready) != "ready\n" {
				t.Fatalf("not ready: %q %v %s", ready, err, stderr.String())
			}
			lock, err := os.OpenFile(filepath.Join(f.s.cache, "write.lock"), os.O_RDWR, 0600)
			if err != nil {
				t.Fatal(err)
			}
			defer lock.Close()
			if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != syscall.EWOULDBLOCK {
				t.Fatalf("expected contention, got %v", err)
			}
			other := newFixture(t)
			publish(t, other) // unrelated workspace never waits
			if err := cmd.Process.Kill(); err != nil {
				t.Fatal(err)
			}
			_ = cmd.Wait()
			if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
				t.Fatalf("kernel retained dead owner's lock: %v", err)
			}
			if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_UN); err != nil {
				t.Fatal(err)
			}
			lookup(t, f, "miss")
			publish(t, f)
			lookup(t, f, "hit")
			if _, err := os.Stat(filepath.Join(f.s.root, "purescript/Partial.go")); !os.IsNotExist(err) {
				t.Fatal("interrupted output left behind")
			}
		})
	}
}

func TestConcurrentPublishers(t *testing.T) {
	f := newFixture(t)
	req := f.req
	req.ExpectedKey = call(t, req, "snapshot").Snapshot.Key
	req.Operation = "publish"
	var wg sync.WaitGroup
	for i := 0; i < 6; i++ {
		command := child(t, f, "exchange", req)
		wg.Add(1)
		go func() {
			defer wg.Done()
			out, err := command.CombinedOutput()
			if err != nil || !strings.Contains(string(out), `"status":"published"`) {
				t.Errorf("publisher: %s %v", out, err)
			}
		}()
	}
	wg.Wait()
	lookup(t, f, "hit")
}
