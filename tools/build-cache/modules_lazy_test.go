package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"
)

func lazyModuleFixture(t *testing.T) (string, buildSession, *lazyPrior) {
	t.Helper()
	_, result, state, _ := packageModuleFixture(t)
	state.Spec.Options = append(state.Spec.Options, Setting{"module-lazy", "1"})
	stage := t.TempDir()
	if err := os.Mkdir(filepath.Join(stage, "prior"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := freezeModuleCache(stage, state, false, &result); err != nil {
		t.Fatal(err)
	}
	prior, err := getLazyPrior(stage, state)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { forgetLazyPrior(stage) })
	return stage, state, prior
}

func TestLazyModuleMaterializationIsSelectiveAndConcurrent(t *testing.T) {
	stage, state, prior := lazyModuleFixture(t)
	for _, part := range []string{"0.json", "1.json", "2.json", "3.json", "module.go", "ffi.go"} {
		if _, err := os.Stat(filepath.Join(stage, "prior/Main", part)); !os.IsNotExist(err) {
			t.Fatalf("eager materialization of %s: %v", part, err)
		}
	}
	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := materializePrior(stage, state, []string{"Main"}, []string{"1"}); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
	expected, err := prior.store.validateObject(prior.modules["Main"].Parts[1].Object)
	if err != nil {
		t.Fatal(err)
	}
	actual, err := os.ReadFile(filepath.Join(stage, "prior/Main/1.json"))
	if err != nil || !bytes.Equal(actual, expected) {
		t.Fatalf("part bytes differ: %v", err)
	}
	if _, err := os.Stat(filepath.Join(stage, "prior/Main/0.json")); !os.IsNotExist(err) {
		t.Fatalf("materialized unrelated body: %v", err)
	}
	// A fresh helper process reconstructs authority and revalidates private bytes.
	forgetLazyPrior(stage)
	if err := materializePrior(stage, state, []string{"Main"}, []string{"1", "module", "ffi"}); err != nil {
		t.Fatal(err)
	}
	if err := restorePriorSources(stage, state); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"Main.go", "Main_ffi.go"} {
		if _, err := os.Stat(filepath.Join(stage, "output/purescript", name)); err != nil {
			t.Fatal(err)
		}
	}
	if err := materializePrior(stage, state, []string{"../Main"}, []string{"1"}); err == nil {
		t.Fatal("accepted arbitrary owner path")
	}
	if err := materializePrior(stage, state, []string{"Main"}, []string{"../1"}); err == nil {
		t.Fatal("accepted arbitrary part path")
	}
}

func TestLazyModuleInheritanceUsesFrozenAuthority(t *testing.T) {
	stage, state, prior := lazyModuleFixture(t)
	old := prior.modules["Main"]
	reused := ReusedModule{Name: old.Name, Prepared: old.Prepared, OptimizerEnvironment: Key("new-prefix"), CallerDemands: old.CallerDemands, EmitterEnvironment: old.EmitterEnvironment, Metadata: old.Metadata, FFIDigest: old.FFIDigest}
	req := Request{Modules: []StagedModule{reused.staged()}, inherited: prior.manifest}
	manifest, err := prior.store.stage(req, state.Snapshot, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(manifest.Modules) != 1 {
		t.Fatal("lost inherited module")
	}
	actual := manifest.Modules[0]
	if !reflect.DeepEqual(actual.Parts, old.Parts) || !reflect.DeepEqual(actual.Sources, old.Sources) || !reflect.DeepEqual(actual.Reads, old.Reads) || actual.OptimizerEnvironment != reused.OptimizerEnvironment {
		t.Fatal("inheritance changed artifacts or failed to update key")
	}
	if _, err := os.Stat(filepath.Join(stage, "prior/Main/0.json")); !os.IsNotExist(err) {
		t.Fatal("publication materialized backend body")
	}
	req.inherited = nil
	if _, err := prior.store.stage(req, state.Snapshot, nil); err == nil {
		t.Fatal("external request inherited without frozen authority")
	}
}

func TestLazyModuleCorruptionFailsBeforePublication(t *testing.T) {
	for _, kind := range []string{"part", "source"} {
		t.Run(kind, func(t *testing.T) {
			t.Setenv(packPublication, "0")
			stage, state, prior := lazyModuleFixture(t)
			old := prior.modules["Main"]
			object := old.Parts[1].Object
			part := "1"
			if kind == "source" {
				object = old.Sources[0].Object
				part = "module"
			}
			put(t, filepath.Join(prior.store.cache, "objects", object.SHA256), "corrupt")
			if err := materializePrior(stage, state, []string{"Main"}, []string{part}); err == nil {
				t.Fatal("materialized corrupted object")
			}
			reused := ReusedModule{Name: old.Name, Prepared: old.Prepared, OptimizerEnvironment: old.OptimizerEnvironment, CallerDemands: old.CallerDemands, EmitterEnvironment: old.EmitterEnvironment, Metadata: old.Metadata, FFIDigest: old.FFIDigest}
			req := Request{Modules: []StagedModule{reused.staged()}, inherited: prior.manifest}
			if _, err := prior.store.stage(req, state.Snapshot, nil); err == nil || !strings.Contains(err.Error(), "inherited module") {
				t.Fatalf("inherited corrupt object: %v", err)
			}
		})
	}
}

func TestLazyModuleAbsentForeignIsOptional(t *testing.T) {
	stage, state, prior := lazyModuleFixture(t)
	memo := prior.modules["Main"]
	memo.Sources = memo.Sources[:1]
	prior.modules["Main"] = memo
	delete(prior.outputs["Main"], "ffi")
	if err := materializePrior(stage, state, []string{"Main"}, []string{"ffi"}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(stage, "prior/Main/ffi.go")); !os.IsNotExist(err) {
		t.Fatal("invented absent FFI")
	}
}

func TestLazyModuleReadsRemainCommittedProofs(t *testing.T) {
	t.Setenv(packPublication, "0")
	_, result, state, _ := packageModuleFixture(t)
	s := store{root: state.Spec.Output, cache: filepath.Join(state.Spec.Output, ".gopurs-cache")}
	manifest, why := s.readManifest()
	if manifest == nil {
		t.Fatal(why)
	}
	object, err := s.putObject([]byte("Provider\n"))
	if err != nil {
		t.Fatal(err)
	}
	manifest.Modules[0].Reads = &Artifact{Codec: moduleReadsCodec, Object: object}
	raw, _ := json.Marshal(manifest)
	manifestObject, err := s.putObject(raw)
	if err != nil {
		t.Fatal(err)
	}
	pointerBytes, _ := json.Marshal(pointer{Format: Format, Schema: Schema, Manifest: manifestObject})
	put(t, filepath.Join(s.cache, "current.json"), string(pointerBytes))
	state.Spec.Options = append(state.Spec.Options, Setting{"module-lazy", "1"})
	for _, corrupt := range []bool{false, true} {
		stage := t.TempDir()
		if err := os.Mkdir(filepath.Join(stage, "prior"), 0700); err != nil {
			t.Fatal(err)
		}
		if corrupt {
			put(t, filepath.Join(s.cache, "objects", object.SHA256), "bad")
		}
		if err := freezeModuleCache(stage, state, false, &result); err != nil {
			t.Fatal(err)
		}
		data, err := os.ReadFile(filepath.Join(stage, "prior/index.json"))
		if err != nil {
			t.Fatal(err)
		}
		var index []ModuleMemo
		if err := strictJSON(data, &index); err != nil {
			t.Fatal(err)
		}
		if corrupt {
			if len(index) != 0 {
				t.Fatal("corrupt reads admitted candidate")
			}
		} else {
			data, err := os.ReadFile(filepath.Join(stage, "prior/Main/reads.json"))
			if err != nil || string(data) != "Provider\n" {
				t.Fatalf("lost committed reads: %s %v", data, err)
			}
		}
		forgetLazyPrior(stage)
	}
}

func TestLazyModulePinnedPacksSurviveConcurrentPrune(t *testing.T) {
	t.Setenv(packPublication, "1")
	stage, state, prior := lazyModuleFixture(t)
	public := store{root: state.Spec.Output, cache: filepath.Join(state.Spec.Output, ".gopurs-cache")}
	if err := public.ensurePacks(); err != nil {
		t.Fatal(err)
	}
	old := prior.modules["Main"]
	if len(public.packs.entries) == 0 {
		t.Fatal("fixture did not create packed artifacts")
	}
	if err := public.retirePackObjects(&Manifest{}); err != nil {
		t.Fatal(err)
	}
	if _, err := public.validateObject(old.Parts[0].Object); err == nil {
		t.Fatal("prune did not remove prior object")
	}
	// Also release native file descriptors: the private hard links must support a
	// later independent helper, rather than relying on an unlinked open inode.
	forgetLazyPrior(stage)
	if err := materializePrior(stage, state, []string{"Main"}, []string{"0", "1", "2", "3", "module", "ffi"}); err != nil {
		t.Fatal(err)
	}
	prior, err := getLazyPrior(stage, state)
	if err != nil {
		t.Fatal(err)
	}
	reused := ReusedModule{Name: old.Name, Prepared: old.Prepared, OptimizerEnvironment: old.OptimizerEnvironment, CallerDemands: old.CallerDemands, EmitterEnvironment: old.EmitterEnvironment, Metadata: old.Metadata, FFIDigest: old.FFIDigest}
	req := Request{Modules: []StagedModule{reused.staged()}, inherited: prior.manifest, inheritedStore: &prior.store}
	manifest, err := public.stage(req, state.Snapshot, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(manifest.Modules[0].Parts, old.Parts) || !reflect.DeepEqual(manifest.Modules[0].Sources, old.Sources) {
		t.Fatal("restored references changed")
	}
	for _, part := range artifacts(manifest) {
		if _, err := public.validateObject(part.Object); err != nil {
			t.Fatal(err)
		}
	}
}

func TestNativeMaterializationDoesNotReparseSessionInventory(t *testing.T) {
	_, state, _ := lazyModuleFixture(t)
	stage := filepath.Join(state.Spec.Output, ".gopurs-cache", "stages", "build-fast-materialize")
	if err := os.MkdirAll(filepath.Join(stage, "prior"), 0700); err != nil {
		t.Fatal(err)
	}
	var frozen SessionResponse
	if err := freezeModuleCache(stage, state, false, &frozen); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { forgetLazyPrior(stage) })
	// The already frozen native state is authoritative. A helper without that
	// state must still parse/validate session.json instead of taking this path.
	put(t, filepath.Join(stage, "session.json"), "unparsed inventory sentinel")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(state.Spec.Workspace); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(cwd)
	request := SessionRequest{Operation: "materialize", Stage: stage, Owners: []string{"Main"}, Parts: []string{"1"}}
	var response SessionResponse
	if err := sessionOperation(request, &response); err != nil {
		t.Fatal(err)
	}
	if response.Status != "ready" || response.Count != state.Count {
		t.Fatalf("bad native response: %+v", response)
	}
	forgetLazyPrior(stage)
	if err := sessionOperation(request, &response); err == nil {
		t.Fatal("helper bypassed invalid session file")
	}
}
