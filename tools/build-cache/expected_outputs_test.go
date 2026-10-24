package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func expectedOutputWrite(t *testing.T, path string, data []byte) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}
}

func expectedOutputFixture(t *testing.T) (Request, store, string) {
	t.Helper()
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"compiler", "parser", "runtime"} {
		expectedOutputWrite(t, filepath.Join(root, name), []byte(name))
	}
	expectedOutputWrite(t, filepath.Join(root, "output/Main/corefn.json"), []byte(`{"moduleName":["Main"]}`))
	stage := filepath.Join(root, "private-stage")
	plan := packagePlanResult{Format: packageFilesFormat, PlanHash: Key("expected-output-test"), Packages: []packagePlanPackage{{Name: "p_root", Declarations: 1}}, Entries: []packagePlanEntry{}}
	for _, file := range []struct{ path, kind string }{
		{"p_root/0_gopurs_package.go", "package"},
		{"p_root/zz_gopurs_order.go", "order"},
		{"init_all/init.go", "package-init"},
		{"init_external/init.go", "package-init"},
	} {
		path, kind := file.path, file.kind
		content := []byte("package p_root\n// " + path + " stable\n")
		plan.Files = append(plan.Files, packagePlanFile{Path: path, Kind: kind, SHA256: digest(content), Bytes: int64(len(content))})
		expectedOutputWrite(t, filepath.Join(stage, "packages", path), content)
	}
	encoded, err := json.Marshal(plan)
	if err != nil {
		t.Fatal(err)
	}
	expectedOutputWrite(t, filepath.Join(stage, "package-result.json"), encoded)
	outputs, err := packageOutputs(stage, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	deferred, expected, err := packageOutputPlan(stage, nil, nil)
	if err != nil || !reflect.DeepEqual(outputs, deferred) {
		t.Fatalf("metadata output mismatch: %v", err)
	}
	request := Request{Schema: Schema, Operation: "snapshot", Spec: Spec{Workspace: root, Output: filepath.Join(root, "output"),
		Compiler: []Input{{Name: "compiler", Kind: "file", Path: "compiler"}, {Name: "parser", Kind: "file", Path: "parser"}, {Name: "runtime", Kind: "file", Path: "runtime"}},
		Options:  []Setting{}, Inputs: []Input{{Name: "tast", Kind: "tast", Path: filepath.Join(root, "output")}}}, Outputs: outputs, expectedOutputs: expected}
	snap, err := execute(request)
	if err != nil {
		t.Fatal(err)
	}
	request.Operation, request.ExpectedKey = "publish", snap.Snapshot.Key
	committed, err := execute(request)
	if err != nil || committed.Status != "published" {
		t.Fatalf("prime: %+v %v", committed, err)
	}
	return request, store{root: request.Spec.Output, cache: filepath.Join(request.Spec.Output, ".gopurs-cache")}, stage
}

func expectedOutputPublic(t *testing.T, request Request, s store) map[string][]byte {
	t.Helper()
	result := map[string][]byte{}
	paths := []string{filepath.Join(s.cache, "current.json"), filepath.Join(s.root, ".gopurs-ownership.json")}
	for _, out := range request.Outputs {
		paths = append(paths, filepath.Join(s.root, out.Path))
	}
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		result[path] = data
	}
	return result
}

func TestExpectedOutputRejectsChangedBytesAndInventory(t *testing.T) {
	for _, mode := range []string{"same-size", "length", "missing", "symlink", "missing-expectation", "external-expectation", "malformed-expectation", "duplicate-output", "non-package-expectation"} {
		t.Run(mode, func(t *testing.T) {
			request, s, stage := expectedOutputFixture(t)
			before := expectedOutputPublic(t, request, s)
			originalOutputs := append([]StagedOutput(nil), request.Outputs...)
			out := request.Outputs[0]
			original, err := os.ReadFile(out.Source)
			if err != nil {
				t.Fatal(err)
			}
			var invalid []byte
			switch mode {
			case "same-size":
				invalid = append([]byte{}, original...)
				invalid[len(invalid)-2] ^= 1
				expectedOutputWrite(t, out.Source, invalid)
			case "length":
				invalid = append(append([]byte{}, original...), '!')
				expectedOutputWrite(t, out.Source, invalid)
			case "missing":
				if err := os.Remove(out.Source); err != nil {
					t.Fatal(err)
				}
			case "symlink":
				target := filepath.Join(stage, "other.go")
				expectedOutputWrite(t, target, original)
				if err := os.Remove(out.Source); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(target, out.Source); err != nil {
					t.Fatal(err)
				}
			case "missing-expectation":
				delete(request.expectedOutputs, out.Path)
			case "external-expectation":
				request.expectedOutputs["purescript/p_root/external.go"] = request.expectedOutputs[out.Path]
			case "malformed-expectation":
				request.expectedOutputs[out.Path] = Object{SHA256: "broken", Bytes: int64(len(original))}
			case "duplicate-output":
				request.Outputs = append(request.Outputs, out)
			case "non-package-expectation":
				path := "gopurs_runtime/runtime.go"
				request.Outputs = append(request.Outputs, StagedOutput{Path: path, Kind: "runtime", Source: out.Source})
				request.expectedOutputs[path] = request.expectedOutputs[out.Path]
			}
			if mode == "same-size" || mode == "length" || mode == "missing" || mode == "symlink" {
				if _, err := packageOutputs(stage, nil, nil); err == nil {
					t.Fatal("export verification accepted changed staged output")
				}
			}
			_, err = execute(request)
			if err == nil {
				t.Fatal("invalid private output was accepted")
			}
			for path, want := range before {
				got, err := os.ReadFile(path)
				if err != nil || !bytes.Equal(got, want) {
					t.Fatalf("public mutation %s: %v", path, err)
				}
			}
			if invalid != nil {
				if _, err := os.Stat(filepath.Join(s.cache, "objects", digest(invalid))); !os.IsNotExist(err) {
					t.Fatalf("invalid bytes installed as object: %v", err)
				}
			}
			if mode == "symlink" {
				if err := os.Remove(out.Source); err != nil {
					t.Fatal(err)
				}
			}
			expectedOutputWrite(t, out.Source, original)
			_, expected, err := packageOutputPlan(stage, nil, nil)
			if err != nil {
				t.Fatal(err)
			}
			request.expectedOutputs = expected
			request.Outputs = originalOutputs
			result, err := execute(request)
			if err != nil || result.Status != "published" {
				t.Fatalf("retry: %+v %v", result, err)
			}
		})
	}
}

func TestExpectedOutputValidatesConsumedBytesAndKeepsPublicJSON(t *testing.T) {
	request, s, _ := expectedOutputFixture(t)
	encoded, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	legacy := request
	legacy.expectedOutputs = nil
	oldJSON, err := json.Marshal(legacy)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(encoded, oldJSON) || bytes.Contains(encoded, []byte("expectedOutputs")) {
		t.Fatal("private field changed public JSON")
	}
	var decoded Request
	if err := strictJSON(encoded, &decoded); err != nil || decoded.expectedOutputs != nil {
		t.Fatalf("public request decode: %v", err)
	}
	modified := append([]byte{}, encoded[:len(encoded)-1]...)
	modified = append(modified, []byte(`,"expectedOutputs":{}}`)...)
	if err := strictJSON(modified, &decoded); err == nil {
		t.Fatal("external expectation accepted")
	}
	out := request.Outputs[0]
	data, err := os.ReadFile(out.Source)
	if err != nil {
		t.Fatal(err)
	}
	expected := request.expectedOutputs[out.Path]
	changed := append(append([]byte{}, data...), []byte("// later mutation\n")...)
	expectedOutputWrite(t, out.Source, changed)
	// The exact consumed slice, not a fresh reread or remembered path hash,
	// determines the object that may be installed.
	got, err := s.putStagedObjectExpected(data, &expected)
	if err != nil || got != expected {
		t.Fatalf("consumed bytes rejected: %+v %v", got, err)
	}
	if _, err := s.putStagedObjectExpected(changed, &expected); err == nil {
		t.Fatal("wrong consumed bytes accepted")
	}
	if _, err := os.Stat(filepath.Join(s.cache, "objects", digest(changed))); !os.IsNotExist(err) {
		t.Fatal("wrong slice installed")
	}
	// Historical nil/API path continues to accept the caller's generated bytes.
	result, err := execute(legacy)
	if err != nil || result.Status != "published" {
		t.Fatalf("nil expectation path: %+v %v", result, err)
	}
	gotBytes, err := os.ReadFile(filepath.Join(s.root, out.Path))
	if err != nil || !bytes.Equal(gotBytes, changed) {
		t.Fatalf("legacy source bytes changed: %v", err)
	}
}

func TestExpectedOutputJoinsSiblingsBeforeFailure(t *testing.T) {
	request, s, stage := expectedOutputFixture(t)
	before := expectedOutputPublic(t, request, s)
	var installed []Object
	for i := 0; i < 25; i++ {
		name := strings.Repeat("x", i+1) + ".go"
		path := "purescript/p_root/" + name
		source := filepath.Join(stage, name)
		data := []byte("package p_root\n// " + name + "\n")
		expectedOutputWrite(t, source, data)
		object := Object{SHA256: digest(data), Bytes: int64(len(data))}
		request.Outputs = append(request.Outputs, StagedOutput{Path: path, Kind: "package", Module: strings.TrimSuffix(name, ".go"), Source: source})
		request.expectedOutputs[path] = object
		installed = append(installed, object)
	}
	bad := request.Outputs[len(request.Outputs)-1]
	expectedOutputWrite(t, bad.Source, []byte("corrupt"))
	if _, err := execute(request); err == nil {
		t.Fatal("bad final sibling accepted")
	}
	for _, object := range installed[:len(installed)-1] {
		if !fileMatches(filepath.Join(s.cache, "objects", object.SHA256), object) {
			t.Fatal("staging returned before a valid sibling completed")
		}
	}
	for path, want := range before {
		got, err := os.ReadFile(path)
		if err != nil || !bytes.Equal(got, want) {
			t.Fatalf("public mutation %s: %v", path, err)
		}
	}
}
