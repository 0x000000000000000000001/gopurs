package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestSessionFreezesManyModulesWithOrderedHashes(t *testing.T) {
	f := newFixture(t)
	spec := f.req.Spec
	spec.Output = filepath.Join(f.root, "output")
	spec.Inputs = []Input{{Name: "tast", Kind: "tast", Path: spec.Output}}
	original, err := os.ReadFile(filepath.Join(spec.Output, "Main/corefn.json"))
	if err != nil {
		t.Fatal(err)
	}
	sources := map[string]string{"Main": string(original), "Alias": "symlink provider"}
	for i := 31; i >= 0; i-- {
		name := fmt.Sprintf("Module%02d", i)
		sources[name] = strings.Repeat(name, i+1)
		put(t, filepath.Join(spec.Output, name, "corefn.json"), sources[name])
	}
	put(t, filepath.Join(f.root, "provider/corefn.json"), sources["Alias"])
	if err := os.Symlink(filepath.Join(f.root, "provider"), filepath.Join(spec.Output, "Alias")); err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(spec.Output, "WithoutCoreFn/other.txt"), "not a module")
	put(t, filepath.Join(spec.Output, "plain-file"), "not a directory")
	var result SessionResponse
	if err := beginSession(spec, false, &result); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.RemoveAll(result.Stage) })
	if result.Count != len(sources) {
		t.Fatalf("wrong inventory: %d != %d", result.Count, len(sources))
	}
	data, err := os.ReadFile(filepath.Join(result.Stage, "source-hashes.json"))
	if err != nil {
		t.Fatal(err)
	}
	var hashes []frozenSourceHash
	if err := strictJSON(data, &hashes); err != nil {
		t.Fatal(err)
	}
	names := make([]string, 0, len(sources))
	for name := range sources {
		names = append(names, name)
	}
	sort.Strings(names)
	if len(hashes) != len(names) {
		t.Fatalf("wrong hash inventory: %d", len(hashes))
	}
	for i, name := range names {
		if hashes[i] != (frozenSourceHash{Name: name, Hash: digest([]byte(sources[name]))}) {
			t.Fatalf("hash order/content mismatch for %s: %+v", name, hashes[i])
		}
		put(t, filepath.Join(spec.Output, name, "corefn.json"), "edited after freezing")
		frozen, err := os.ReadFile(filepath.Join(result.Stage, "input", name, "corefn.json"))
		if err != nil || string(frozen) != sources[name] {
			t.Fatalf("private snapshot changed for %s: %v", name, err)
		}
	}
}

func TestBaseSnapshotUsesOnlyMatchingFreshObservations(t *testing.T) {
	f := newFixture(t)
	spec := f.req.Spec
	spec.Output = filepath.Join(f.root, "output")
	observed, err := snapshot(spec)
	if err != nil {
		t.Fatal(err)
	}
	spec.Inputs = []Input{{Name: "tast", Kind: "tast", Path: spec.Output}}
	expected, err := snapshot(spec)
	if err != nil {
		t.Fatal(err)
	}
	actual, ok := baseSnapshot(spec, &observed)
	if !ok || !reflect.DeepEqual(actual, expected) {
		t.Fatalf("base snapshot differs from a fresh observation: %+v / %+v", actual, expected)
	}
	for _, inputs := range [][]Input{
		{{Name: "tast", Kind: "tast", Path: "other-output"}},
		{{Name: "tast", Kind: "file", Path: spec.Output}},
		{{Name: "missing", Kind: "file", Path: "missing"}},
		{spec.Inputs[0], spec.Inputs[0]},
	} {
		changed := spec
		changed.Inputs = inputs
		if _, ok := baseSnapshot(changed, &observed); ok {
			t.Fatalf("accepted a changed or unknown recipe: %+v", inputs)
		}
	}
	put(t, filepath.Join(spec.Output, "Main/corefn.json"), "changed after lookup")
	current, err := snapshot(spec)
	if err != nil || current.Key == actual.Key {
		t.Fatalf("post-freeze verification cannot detect a later edit: %+v / %v", current, err)
	}
}

func driverFixture(t *testing.T) (Spec, SessionResponse, buildSession) {
	t.Helper()
	f := newFixture(t)
	spec := f.req.Spec
	spec.Output = filepath.Join(f.root, "output")
	spec.Inputs = []Input{{Name: "tast", Kind: "tast", Path: spec.Output}}
	result := SessionResponse{FFI: []FrozenFFI{}}
	if err := beginSession(spec, false, &result); err != nil {
		t.Fatal(err)
	}
	if result.Status != "miss" || result.Count != 1 {
		t.Fatalf("unexpected begin: %+v", result)
	}
	var state buildSession
	data, err := os.ReadFile(filepath.Join(result.Stage, "session.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := strictJSON(data, &state); err != nil {
		t.Fatal(err)
	}
	return spec, result, state
}

func finishDriver(t *testing.T, result *SessionResponse, state *buildSession, modules []SourceModule, diagnostics bool) {
	t.Helper()
	if err := planSession(result.Stage, state, modules, result); err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(result.Output, "gopurs_runtime/runtime.go"), "package gopurs_runtime\n")
	for _, module := range modules {
		put(t, filepath.Join(result.Output, "purescript", strings.ReplaceAll(module.Name, ".", "_")+".go"), "package purescript\n")
	}
	if err := publishSession(result.Stage, *state, SessionRequest{Diagnostics: diagnostics}, result); err != nil {
		t.Fatal(err)
	}
	if result.Status != "published" {
		t.Fatal(result)
	}
}

func TestDriverFreezesConsumedBytes(t *testing.T) {
	spec, result, state := driverFixture(t)
	modules := []SourceModule{{Name: "Main", Path: "src/Main.purs", HasForeign: true}}
	if err := planSession(result.Stage, &state, modules, &result); err != nil {
		t.Fatal(err)
	}
	original, _ := os.ReadFile(filepath.Join(spec.Workspace, "src/Main.go"))
	put(t, filepath.Join(spec.Workspace, "src/Main.go"), "func Edited() {}")
	frozen, err := os.ReadFile(result.FFI[0].File)
	if err != nil || !bytes.Equal(frozen, original) {
		t.Fatalf("FFI did not retain observed bytes: %s %v", frozen, err)
	}
	tastPath := filepath.Join(spec.Output, "Main/corefn.json")
	originalTast, _ := os.ReadFile(tastPath)
	put(t, tastPath, "edited while compiler is running")
	frozenTast, _ := os.ReadFile(filepath.Join(result.Input, "Main/corefn.json"))
	if !bytes.Equal(frozenTast, originalTast) {
		t.Fatal("decoder reads live TAST")
	}
	hashBytes, err := os.ReadFile(filepath.Join(result.Stage, "source-hashes.json"))
	if err != nil {
		t.Fatal(err)
	}
	var hashes []frozenSourceHash
	if err := strictJSON(hashBytes, &hashes); err != nil || len(hashes) != 1 || hashes[0].Name != "Main" || hashes[0].Hash != digest(frozenTast) {
		t.Fatalf("source hashes do not describe the frozen bytes: %s (%v)", hashBytes, err)
	}
	put(t, filepath.Join(result.Output, "gopurs_runtime/runtime.go"), "package gopurs_runtime\n")
	put(t, filepath.Join(result.Output, "purescript/Main.go"), "package purescript\n")
	if err := publishSession(result.Stage, state, SessionRequest{}, &result); err == nil || !strings.Contains(err.Error(), "inputs-changed") {
		t.Fatalf("accepted a stale build: %v", err)
	}
	if _, err := os.Stat(filepath.Join(spec.Output, "purescript/Main.go")); !os.IsNotExist(err) {
		t.Fatal("stale build touched public output")
	}
	// Returning to the original bytes is safe: those are precisely the captured
	// bytes consumed by the decoder and parser, even after an edit-and-revert.
	put(t, tastPath, string(originalTast))
	put(t, filepath.Join(spec.Workspace, "src/Main.go"), string(original))
	if err := publishSession(result.Stage, state, SessionRequest{}, &result); err != nil {
		t.Fatal(err)
	}
}

func TestDriverIdenticalAndOutputRepair(t *testing.T) {
	spec, result, state := driverFixture(t)
	finishDriver(t, &result, &state, []SourceModule{{Name: "Main", Path: "src/Main.purs", HasForeign: true}}, false)
	put(t, filepath.Join(spec.Output, "go.mod"), "module gopurs/output\n\ngo 1.26.0\n// kept by tidy\n")
	put(t, filepath.Join(spec.Output, "go.sum"), "external sum\n")
	for _, step := range []string{"hit", "output", "recipe", "force"} {
		switch step {
		case "output":
			put(t, filepath.Join(spec.Output, "purescript/Main.go"), "damaged")
		case "recipe":
			put(t, filepath.Join(spec.Output, ".gopurs-cache/recipe.json"), "invalid")
		}
		next := SessionResponse{FFI: []FrozenFFI{}}
		if err := beginSession(spec, step == "force", &next); err != nil {
			t.Fatal(err)
		}
		if step == "hit" {
			if next.Status != "hit" || next.Stage != "" {
				t.Fatalf("hit decoded/froze inputs: %+v", next)
			}
		} else if next.Status != "miss" || next.Stage == "" {
			t.Fatal(next)
		}
	}
	data, _ := os.ReadFile(filepath.Join(spec.Output, "go.mod"))
	if !strings.Contains(string(data), "kept by tidy") {
		t.Fatal("go.mod overwritten")
	}
}

func TestDriverResolutionPriorityAndAbsentRoots(t *testing.T) {
	for _, change := range []string{"companion", "new-version", "skipped-symlink-directory", "missing-body"} {
		t.Run(change, func(t *testing.T) {
			spec, result, state := driverFixture(t)
			os.Remove(filepath.Join(spec.Workspace, "src/Main.go"))
			put(t, filepath.Join(spec.Workspace, ".spago/lib/v2/src/Main.go"), "func Second() {}")
			// Directory inspection must cover a skipped symlink target's type.
			put(t, filepath.Join(spec.Workspace, "target"), "file")
			if err := os.Symlink(filepath.Join(spec.Workspace, "target"), filepath.Join(spec.Workspace, ".spago/earlier")); err != nil {
				t.Fatal(err)
			}
			if change == "missing-body" {
				os.Remove(filepath.Join(spec.Workspace, ".spago/lib/v2/src/Main.go"))
			}
			finishDriver(t, &result, &state, []SourceModule{{Name: "Main", Path: "src/Main.purs", HasForeign: true}}, false)
			if change != "missing-body" && !strings.HasSuffix(result.FFI[0].Path, ".spago/lib/v2/src/Main.go") {
				t.Fatal(result.FFI)
			}
			switch change {
			case "companion":
				put(t, filepath.Join(spec.Workspace, "src/Main.go"), "func Preferred() {}")
			case "new-version":
				put(t, filepath.Join(spec.Workspace, ".spago/lib/v1/src/Main.go"), "func Earlier() {}")
			case "skipped-symlink-directory":
				os.Remove(filepath.Join(spec.Workspace, "target"))
				put(t, filepath.Join(spec.Workspace, "target/src/Main.go"), "func EarlierRoot() {}")
			case "missing-body":
				put(t, filepath.Join(spec.Workspace, ".spago/lib/v2/src/Main.go"), "func Appeared() {}")
			}
			next := SessionResponse{FFI: []FrozenFFI{}}
			if err := beginSession(spec, false, &next); err != nil {
				t.Fatal(err)
			}
			if next.Status != "miss" {
				t.Fatalf("resolution change reused stale output: %+v", next)
			}
		})
	}
}

func TestDriverDiagnosticsAndLegacyOwnership(t *testing.T) {
	for _, scenario := range []string{"diagnostics", "matching-legacy", "unknown-legacy", "changed-legacy"} {
		t.Run(scenario, func(t *testing.T) {
			spec, result, state := driverFixture(t)
			if scenario == "matching-legacy" {
				put(t, filepath.Join(spec.Output, "purescript/Main.go"), "package purescript\n")
			}
			if scenario == "unknown-legacy" {
				put(t, filepath.Join(spec.Output, "purescript/Old.go"), "package purescript\n")
			}
			if scenario == "changed-legacy" {
				put(t, filepath.Join(spec.Output, "purescript/Main.go"), "user modified bytes")
			}
			if err := planSession(result.Stage, &state, []SourceModule{{Name: "Main"}}, &result); err != nil {
				t.Fatal(err)
			}
			put(t, filepath.Join(result.Output, "gopurs_runtime/runtime.go"), "package gopurs_runtime\n")
			put(t, filepath.Join(result.Output, "purescript/Main.go"), "package purescript\n")
			err := publishSession(result.Stage, state, SessionRequest{Diagnostics: scenario == "diagnostics"}, &result)
			if strings.HasPrefix(scenario, "unknown") || strings.HasPrefix(scenario, "changed") {
				if err == nil || !strings.Contains(err.Error(), "verified ownership") {
					t.Fatal(err)
				}
				if _, err := os.Stat(filepath.Join(spec.Output, ".gopurs-ownership.json")); !os.IsNotExist(err) {
					t.Fatal("legacy files adopted without evidence")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			next := SessionResponse{FFI: []FrozenFFI{}}
			if err := beginSession(spec, false, &next); err != nil {
				t.Fatal(err)
			}
			if scenario == "diagnostics" && (next.Status != "miss" || next.Reason != "diagnostics-require-build") {
				t.Fatal(next)
			}
			if scenario == "matching-legacy" && next.Status != "hit" {
				t.Fatal(next)
			}
		})
	}
}

func TestDriverCompilerIdentityIncludesResources(t *testing.T) {
	spec, result, state := driverFixture(t)
	finishDriver(t, &result, &state, []SourceModule{{Name: "Main"}}, false)
	put(t, filepath.Join(spec.Workspace, "ffi-parser"), "new parser executed")
	next := SessionResponse{FFI: []FrozenFFI{}}
	if err := beginSession(spec, false, &next); err != nil {
		t.Fatal(err)
	}
	if next.Status != "miss" {
		t.Fatal("compiler resources not part of identity")
	}
	encoded, _ := json.Marshal(state.Spec)
	if !bytes.Contains(encoded, []byte("ffi-parser")) {
		t.Fatal("missing resource proof")
	}
}
