package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

func TestPackedInputsManyConcurrentSlices(t *testing.T) {
	f := newFixture(t)
	spec := f.req.Spec
	spec.Output = filepath.Join(f.root, "packed-output")
	spec.Inputs = []Input{{Name: "tast", Kind: "tast", Path: spec.Output}}
	spec.Options = append(spec.Options, Setting{Name: "input-layout", Value: "packed-v1"})
	sources := map[string]string{}
	for i := 0; i < 257; i++ {
		name := fmt.Sprintf("Module%03d", i)
		// Include empty sources and unequal slices to expose overlap, holes
		// and completion-order mistakes with the real eight-worker freeze.
		text := string(bytes.Repeat([]byte{byte(i), 0, '\n'}, (i*73)%8192))
		sources[name] = text
		put(t, filepath.Join(spec.Output, name, "corefn.json"), text)
	}
	var result SessionResponse
	if err := beginSession(spec, false, &result); err != nil {
		t.Fatal(err)
	}
	var state buildSession
	data, err := os.ReadFile(filepath.Join(result.Stage, "session.json"))
	if err != nil || strictJSON(data, &state) != nil {
		t.Fatal("invalid session", err)
	}
	index, err := readFrozenInputIndex(result.Stage, state)
	if err != nil || len(index.Modules) != len(sources) {
		t.Fatalf("invalid complete slice inventory: %v", err)
	}
	if err := materializeInputs(result.Stage, state, nil, true); err != nil {
		t.Fatal(err)
	}
	for name, expected := range sources {
		assertFrozenInput(t, result, name, expected)
	}
}

func TestPackedInputWriteErrorsJoinWithoutIndex(t *testing.T) {
	stage := t.TempDir()
	writer, err := newInputPackWriter(stage)
	if err != nil {
		t.Fatal(err)
	}
	if err := writer.file.Close(); err != nil {
		t.Fatal(err)
	}
	seen := make([]bool, 33)
	err = boundedEach(len(seen), func(i int) error {
		data := []byte(fmt.Sprintf("failed source %d", i))
		_, err := writer.append(Fingerprint{Name: fmt.Sprintf("Module%03d", i), Digest: digest(data)}, data)
		seen[i] = true
		return err
	})
	if err == nil {
		t.Fatal("closed-file writes were accepted")
	}
	for _, finished := range seen {
		if !finished {
			t.Fatal("write failure returned before all tasks joined")
		}
	}
	if err := writer.finish(stage, nil); err == nil {
		t.Fatal("closed writer published an index")
	}
	if _, err := os.Stat(filepath.Join(stage, "input-index.json")); !os.IsNotExist(err) {
		t.Fatalf("failed writer left a usable index: %v", err)
	}
}

func TestPackedInputAppendFailureCleansStageAndRetries(t *testing.T) {
	f := newFixture(t)
	spec := f.req.Spec
	spec.Output = filepath.Join(f.root, "output")
	spec.Inputs = []Input{{Name: "tast", Kind: "tast", Path: spec.Output}}
	spec.Options = append(spec.Options, Setting{Name: "input-layout", Value: "packed-v1"})
	invalid := filepath.Join(spec.Output, "Invalid..Module")
	put(t, filepath.Join(invalid, "corefn.json"), "rejected during append")
	var result SessionResponse
	err := beginSession(spec, false, &result)
	if err == nil || !strings.Contains(err.Error(), "invalid frozen input module") {
		t.Fatalf("append failure was not exercised: %v", err)
	}
	stages, err := os.ReadDir(filepath.Join(spec.Output, ".gopurs-cache", "stages"))
	if err != nil || len(stages) != 0 {
		t.Fatalf("failed append retained a partial stage: %v, %v", stages, err)
	}
	if err := os.RemoveAll(invalid); err != nil {
		t.Fatal(err)
	}
	if err := beginSession(spec, false, &result); err != nil || result.Count != 1 {
		t.Fatalf("append failure prevented retry: %+v, %v", result, err)
	}
}

func packedInputFixture(t *testing.T) (SessionResponse, buildSession, map[string]string) {
	t.Helper()
	f := newFixture(t)
	spec := f.req.Spec
	spec.Output = filepath.Join(f.root, "output")
	spec.Inputs = []Input{{Name: "tast", Kind: "tast", Path: spec.Output}}
	spec.Options = append(spec.Options, Setting{Name: "input-layout", Value: "packed-v1"})
	sources := map[string]string{"Main": "main frozen bytes", "Empty": "", "Alias": "linked provider bytes"}
	for i := 0; i < 17; i++ {
		sources[fmt.Sprintf("Module%02d", i)] = fmt.Sprintf("source %d\x00\n", i)
	}
	for name, text := range sources {
		if name != "Alias" {
			put(t, filepath.Join(spec.Output, name, "corefn.json"), text)
		}
	}
	put(t, filepath.Join(f.root, "provider/corefn.json"), sources["Alias"])
	if err := os.Symlink(filepath.Join(f.root, "provider"), filepath.Join(spec.Output, "Alias")); err != nil {
		t.Fatal(err)
	}
	var result SessionResponse
	if err := beginSession(spec, false, &result); err != nil {
		t.Fatal(err)
	}
	if result.InputLayout != "packed-v1" || result.Count != len(sources) {
		t.Fatalf("unexpected packed session: %+v", result)
	}
	var state buildSession
	data, err := os.ReadFile(filepath.Join(result.Stage, "session.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := strictJSON(data, &state); err != nil {
		t.Fatal(err)
	}
	return result, state, sources
}

func assertFrozenInput(t *testing.T, result SessionResponse, name, expected string) {
	t.Helper()
	path := filepath.Join(result.Input, name, "corefn.json")
	actual, err := os.ReadFile(path)
	if err != nil || string(actual) != expected {
		t.Fatalf("frozen bytes differ for %s: %q, %v", name, actual, err)
	}
	info, err := os.Lstat(path)
	if err != nil || !info.Mode().IsRegular() || info.Mode().Perm() != 0400 {
		t.Fatalf("not a private read-only regular file: %s: %v", name, err)
	}
}

func TestPackedInputsSubsetAllAndConcurrentMaterialization(t *testing.T) {
	result, state, sources := packedInputFixture(t)
	entries, err := os.ReadDir(result.Input)
	if err != nil || len(entries) != 0 {
		t.Fatalf("packed freeze created module directories: %v / %v", entries, err)
	}
	index, err := readFrozenInputIndex(result.Stage, state)
	if err != nil {
		t.Fatal(err)
	}
	var hashes []frozenSourceHash
	data, err := os.ReadFile(filepath.Join(result.Stage, "source-hashes.json"))
	if err != nil || strictJSON(data, &hashes) != nil || len(hashes) != len(index.Modules) {
		t.Fatal("invalid source hash inventory", err)
	}
	for i, entry := range index.Modules {
		if hashes[i] != (frozenSourceHash{Name: entry.Name, Hash: digest([]byte(sources[entry.Name]))}) {
			t.Fatalf("source hashes disagree with original bytes: %+v", hashes[i])
		}
		put(t, filepath.Join(state.Spec.Output, entry.Name, "corefn.json"), "live input edited after freeze")
	}
	if err := materializeInputs(result.Stage, state, []string{"Main", "Empty"}, false); err != nil {
		t.Fatal(err)
	}
	entries, _ = os.ReadDir(result.Input)
	if len(entries) != 2 {
		t.Fatalf("materialized unrequested modules: %v", entries)
	}
	assertFrozenInput(t, result, "Main", sources["Main"])
	assertFrozenInput(t, result, "Empty", "")
	// A regular file damaged after a previous request is restored from the pack.
	if err := writePrivateFile(filepath.Join(result.Input, "Main/corefn.json"), []byte("damaged")); err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	errors := make(chan error, 16)
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			errors <- materializeInputs(result.Stage, state, []string{"Main", fmt.Sprintf("Module%02d", i)}, false)
		}(i)
	}
	wg.Wait()
	close(errors)
	for err := range errors {
		if err != nil {
			t.Fatal(err)
		}
	}
	if err := materializeInputs(result.Stage, state, nil, true); err != nil {
		t.Fatal(err)
	}
	for name, expected := range sources {
		assertFrozenInput(t, result, name, expected)
		entries, _ := os.ReadDir(filepath.Join(result.Input, name))
		if len(entries) != 1 || entries[0].Name() != "corefn.json" {
			t.Fatalf("temporary materialization files remain: %v", entries)
		}
	}
	current, err := snapshot(state.Spec)
	if err != nil || current.Key == state.Snapshot.Key {
		t.Fatalf("live edit escaped publication snapshot check: %v", err)
	}
}

func TestPackedInputsRejectCorruption(t *testing.T) {
	for _, kind := range []string{"pack-byte", "pack-truncated", "index-hash", "index-resolved", "index-offset", "index-count", "index-duplicate", "index-name"} {
		t.Run(kind, func(t *testing.T) {
			result, state, _ := packedInputFixture(t)
			index, err := readFrozenInputIndex(result.Stage, state)
			if err != nil {
				t.Fatal(err)
			}
			switch kind {
			case "pack-byte", "pack-truncated":
				path := filepath.Join(result.Stage, "input.pack")
				data, err := os.ReadFile(path)
				if err != nil {
					t.Fatal(err)
				}
				if kind == "pack-byte" {
					data[0] ^= 0xff
				} else {
					data = data[:len(data)-1]
				}
				if err := writePrivateFile(path, data); err != nil {
					t.Fatal(err)
				}
			default:
				switch kind {
				case "index-hash":
					index.Modules[0].Object.SHA256 = digest([]byte("different"))
				case "index-resolved":
					index.Modules[0].Resolved += "-changed"
				case "index-offset":
					index.Modules[0].Offset++
				case "index-count":
					index.Modules = index.Modules[1:]
				case "index-duplicate":
					index.Modules[1] = index.Modules[0]
				case "index-name":
					index.Modules[0].Name = "../escape"
				}
				data, _ := json.Marshal(index)
				if err := writePrivateFile(filepath.Join(result.Stage, "input-index.json"), data); err != nil {
					t.Fatal(err)
				}
			}
			if err := materializeInputs(result.Stage, state, nil, true); err == nil {
				t.Fatal("accepted corrupt frozen input", kind)
			}
		})
	}
}

func TestPackedInputsRejectPrivateSymlinks(t *testing.T) {
	for _, path := range []string{"input.pack", "input-index.json", "input", "input/Main", "input/Main/corefn.json"} {
		t.Run(path, func(t *testing.T) {
			result, state, _ := packedInputFixture(t)
			external := t.TempDir()
			put(t, filepath.Join(external, "sentinel"), "untouched")
			target := filepath.Join(result.Stage, filepath.FromSlash(path))
			if err := os.MkdirAll(filepath.Dir(target), 0700); err != nil {
				t.Fatal(err)
			}
			os.RemoveAll(target)
			if err := os.Symlink(external, target); err != nil {
				t.Fatal(err)
			}
			if err := materializeInputs(result.Stage, state, []string{"Main"}, false); err == nil {
				t.Fatal("accepted a private symlink")
			}
			entries, _ := os.ReadDir(external)
			data, _ := os.ReadFile(filepath.Join(external, "sentinel"))
			if len(entries) != 1 || string(data) != "untouched" {
				t.Fatal("modified symlink target")
			}
		})
	}
}

func TestPackedInputsValidateWholeRequestBeforeWriting(t *testing.T) {
	result, state, _ := packedInputFixture(t)
	for _, owners := range [][]string{{"Main", "Unknown"}, {"Main", "Main"}, {"Main", "../escape"}} {
		if err := materializeInputs(result.Stage, state, owners, false); err == nil {
			t.Fatalf("accepted invalid request: %v", owners)
		}
	}
	if err := materializeInputs(result.Stage, state, []string{"Main"}, true); err == nil {
		t.Fatal("accepted all with named owners")
	}
	entries, _ := os.ReadDir(result.Input)
	if len(entries) != 0 {
		t.Fatal("invalid request partially materialized")
	}
}

func TestPackedInputsProtocolAndEagerCompatibility(t *testing.T) {
	result, state, sources := packedInputFixture(t)
	previous, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(state.Spec.Workspace); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(previous) })
	request, _ := json.Marshal(SessionRequest{Operation: "materialize-input", Stage: result.Stage, All: true})
	var response SessionResponse
	if err := strictJSON([]byte(SessionExchange(string(request))), &response); err != nil || response.Status != "ready" || response.InputLayout != "packed-v1" {
		t.Fatalf("materialize protocol failed: %+v %v", response, err)
	}
	for name, expected := range sources {
		assertFrozenInput(t, result, name, expected)
	}
	if err := materializeInputs("absent-stage", buildSession{}, nil, true); err != nil {
		t.Fatal("eager session is not a no-op", err)
	}
	spec, eager, _ := driverFixture(t)
	if eager.InputLayout != "" {
		t.Fatal("old client no longer uses eager inputs")
	}
	original, _ := os.ReadFile(filepath.Join(spec.Output, "Main/corefn.json"))
	frozen, _ := os.ReadFile(filepath.Join(eager.Input, "Main/corefn.json"))
	if !bytes.Equal(original, frozen) || inputLayout(spec) != "" {
		t.Fatal("eager client snapshot changed")
	}
}
