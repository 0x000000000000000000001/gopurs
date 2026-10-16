package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// Model the resolver's public inventory: the module's unqualified source exists
// only in staging, while the committed files are already split/qualified.
func packageModuleFixture(t *testing.T) (Spec, SessionResponse, buildSession, map[string]string) {
	t.Helper()
	spec, result, state := driverFixture(t)
	state.Spec.Options = append(state.Spec.Options, Setting{"go-layout", packageLayout})
	var err error
	state.Snapshot, err = snapshot(state.Spec)
	if err != nil {
		t.Fatal(err)
	}
	if err := planSession(result.Stage, &state, []SourceModule{{Name: "Main", Path: "src/Main.purs", HasForeign: true}}, &result); err != nil {
		t.Fatal(err)
	}
	sources := map[string]string{"module": "package purescript\nfunc Use() int { return Foreign() }\n", "ffi": "package purescript\nfunc Foreign() int { return 42 }\n"}
	put(t, filepath.Join(result.Output, "gopurs_runtime/runtime.go"), "package gopurs_runtime\n")
	put(t, filepath.Join(result.Output, "purescript/Main.go"), sources["module"])
	put(t, filepath.Join(result.Output, "purescript/Main_ffi.go"), sources["ffi"])
	files := []map[string]any{}
	for _, file := range []struct{ path, kind, module, content string }{
		{"p_root/Main.go", "package", "Main", "package p_root\nfunc Use() int {return 42}\n"},
		{"p_root/Main_ffi.go", "ffi-package", "Main", "package p_root\nfunc Foreign() int {return 42}\n"},
		{"p_root/zz_gopurs_order.go", "order", "", "package p_root\n"},
		{"init_all/init.go", "package-init", "", "package init_all\n"},
		{"init_external/init.go", "package-init", "", "package init_external\n"},
	} {
		put(t, filepath.Join(result.Stage, "packages", file.path), file.content)
		files = append(files, map[string]any{"path": file.path, "kind": file.kind, "moduleName": file.module, "sha256": digest([]byte(file.content)), "bytes": len(file.content)})
	}
	plan, _ := json.Marshal(map[string]any{"format": packageFilesFormat, "planHash": Key("plan"), "packages": []any{map[string]any{"name": "p_root", "dependencies": []string{}, "declarations": 2}}, "files": files, "entries": []any{}})
	put(t, filepath.Join(result.Stage, "package-result.json"), string(plan))
	memo, _ := json.Marshal(ModuleMemo{Name: "Main", Prepared: Key("prepared"), OptimizerEnvironment: Key("optimizer"), EmitterEnvironment: Key("emitter"), CallerDemands: Key("callers")})
	put(t, filepath.Join(result.Stage, "modules/Main/memo.json"), string(memo))
	for _, name := range []string{"0.json", "1.json", "2.json", "3.json"} {
		put(t, filepath.Join(result.Stage, "modules/Main", name), `["unit"]`)
	}
	if err := publishSession(result.Stage, state, SessionRequest{ModuleCache: true}, &result); err != nil {
		t.Fatal(err)
	}
	return spec, result, state, sources
}

func TestPackageModuleCandidatesRestoreUnqualifiedSources(t *testing.T) {
	spec, result, state, sources := packageModuleFixture(t)
	// Public package output is not a reconstruction recipe for the module.
	put(t, filepath.Join(spec.Output, "purescript/p_root/Main.go"), "externally modified")
	next := t.TempDir()
	if err := os.Mkdir(filepath.Join(next, "prior"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := freezeModuleCache(next, state, false, &result); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(next, "prior/index.json"))
	var candidates []ModuleMemo
	if err != nil {
		t.Fatal(err)
	}
	if err := strictJSON(data, &candidates); err != nil {
		t.Fatal(err)
	}
	if len(candidates) != 1 {
		t.Fatalf("package layout lost module candidates: %s", data)
	}
	for kind, want := range sources {
		actual, err := os.ReadFile(filepath.Join(next, "prior/Main", kind+".go"))
		if err != nil || !bytes.Equal(actual, []byte(want)) {
			t.Fatalf("unqualified %s: %s %v", kind, actual, err)
		}
	}
}

func TestPackageModuleSourceIntegrityAndPrivateLifetime(t *testing.T) {
	for _, corruption := range []string{"compressed", "raw-length", "raw-hash", "gzip-trailer", "missing", "version"} {
		t.Run(corruption, func(t *testing.T) {
			spec, result, state, sources := packageModuleFixture(t)
			s := store{root: spec.Output, cache: filepath.Join(spec.Output, ".gopurs-cache")}
			manifest, why := s.readManifest()
			if manifest == nil || len(manifest.Modules[0].Sources) != 2 {
				t.Fatalf("source manifest: %v %s", manifest, why)
			}
			first := manifest.Modules[0].Sources[0]
			freeze := func() string {
				t.Helper()
				directory := t.TempDir()
				if err := os.Mkdir(filepath.Join(directory, "prior"), 0700); err != nil {
					t.Fatal(err)
				}
				if err := freezeModuleCache(directory, state, false, &result); err != nil {
					t.Fatal(err)
				}
				return directory
			}
			frozen := freeze()
			switch corruption {
			case "compressed":
				put(t, filepath.Join(s.cache, "objects", first.Object.SHA256), "corrupt")
			case "missing":
				if err := os.Remove(filepath.Join(s.cache, "objects", first.Object.SHA256)); err != nil {
					t.Fatal(err)
				}
			case "raw-length":
				manifest.Modules[0].Sources[0].Raw.Bytes--
			case "raw-hash":
				manifest.Modules[0].Sources[0].Raw.SHA256 = Key("wrong raw source")
			case "version":
				manifest.Modules[0].Sources[0].Codec = "gopurs/go-source-gzip-v2"
			case "gzip-trailer":
				data, err := s.readObject(first.Object)
				if err != nil {
					t.Fatal(err)
				}
				data[len(data)-1] ^= 1
				object, err := s.putStagedObject(data)
				if err != nil {
					t.Fatal(err)
				}
				manifest.Modules[0].Sources[0].Object = object
			}
			// Rehashing the test manifest exercises schema/decoded-byte checks in
			// addition to the usual content-addressed object integrity check.
			if corruption != "compressed" && corruption != "missing" {
				data, _ := json.Marshal(manifest)
				object, err := s.putStagedObject(data)
				if err != nil {
					t.Fatal(err)
				}
				pointer, _ := json.Marshal(pointer{Format: Format, Schema: Schema, Manifest: object})
				put(t, filepath.Join(s.cache, "current.json"), string(pointer))
			}
			after := freeze()
			index, err := os.ReadFile(filepath.Join(after, "prior/index.json"))
			if err != nil || string(index) != "[]" {
				t.Fatalf("corrupt source remained eligible: %s %v", index, err)
			}
			private, err := os.ReadFile(filepath.Join(frozen, "prior/Main/module.go"))
			if err != nil || string(private) != sources["module"] {
				t.Fatalf("private candidate changed: %s %v", private, err)
			}
			if err := publishSession(result.Stage, state, SessionRequest{ModuleCache: true}, &result); err != nil {
				t.Fatal(err)
			}
			repaired := freeze()
			data, err := os.ReadFile(filepath.Join(repaired, "prior/Main/module.go"))
			if err != nil || string(data) != sources["module"] {
				t.Fatalf("invalid encoded source was not repaired: %s %v", data, err)
			}
		})
	}
}

func TestPackageModuleSourcesAreStableAndPublicationIsAtomic(t *testing.T) {
	spec, result, state, _ := packageModuleFixture(t)
	s := store{root: spec.Output, cache: filepath.Join(spec.Output, ".gopurs-cache")}
	before, _ := s.readManifest()
	lock, err := os.Stat(filepath.Join(s.cache, "write.lock"))
	if err != nil {
		t.Fatal(err)
	}
	if err := publishSession(result.Stage, state, SessionRequest{ModuleCache: true}, &result); err != nil {
		t.Fatal(err)
	}
	after, _ := s.readManifest()
	x, _ := json.Marshal(before)
	y, _ := json.Marshal(after)
	if !bytes.Equal(x, y) {
		t.Fatal("identical source publication changed the manifest")
	}
	pointerBefore, _ := os.ReadFile(filepath.Join(s.cache, "current.json"))
	if err := os.Remove(filepath.Join(result.Output, "purescript/Main_ffi.go")); err != nil {
		t.Fatal(err)
	}
	if err := publishSession(result.Stage, state, SessionRequest{ModuleCache: true}, &result); err == nil {
		t.Fatal("missing raw FFI accepted")
	}
	pointerAfter, _ := os.ReadFile(filepath.Join(s.cache, "current.json"))
	if !bytes.Equal(pointerBefore, pointerAfter) {
		t.Fatal("failed source stage committed a manifest")
	}
	currentLock, err := os.Stat(filepath.Join(s.cache, "write.lock"))
	if err != nil || !os.SameFile(lock, currentLock) {
		t.Fatalf("source publication replaced the lock: %v", err)
	}
}

func TestModuleCandidatesArePrivateVerifiedSnapshots(t *testing.T) {
	spec, result, state := driverFixture(t)
	if err := planSession(result.Stage, &state, []SourceModule{{Name: "Main"}}, &result); err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(result.Output, "gopurs_runtime/runtime.go"), "package gopurs_runtime\n")
	put(t, filepath.Join(result.Output, "purescript/Main.go"), "package purescript\n")
	memo := ModuleMemo{Name: "Main", Prepared: Key("prepared"), OptimizerEnvironment: Key("optimizer"), EmitterEnvironment: Key("emitter"), CallerDemands: Key("callers"), Parts: []Artifact{}}
	data, _ := json.Marshal(memo)
	directory := filepath.Join(result.Stage, "modules/Main")
	put(t, filepath.Join(directory, "memo.json"), string(data))
	for _, name := range []string{"0.json", "1.json", "2.json", "3.json"} {
		put(t, filepath.Join(directory, name), `["unit"]`)
	}
	if err := publishSession(result.Stage, state, SessionRequest{ModuleCache: true}, &result); err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(spec.Output, "Main/corefn.json"), "changed TAST")
	begin := func(force bool) SessionResponse {
		t.Helper()
		var next SessionResponse
		if err := beginSession(spec, force, &next); err != nil {
			t.Fatal(err)
		}
		return next
	}
	index := func(next SessionResponse) []ModuleMemo {
		t.Helper()
		bytes, err := os.ReadFile(filepath.Join(next.Stage, "prior/index.json"))
		if err != nil {
			t.Fatal(err)
		}
		var values []ModuleMemo
		if err := strictJSON(bytes, &values); err != nil {
			t.Fatal(err)
		}
		return values
	}
	next := begin(false)
	candidates := index(next)
	if len(candidates) != 1 || len(candidates[0].Parts) != 4 {
		t.Fatalf("missing candidate: %+v", candidates)
	}
	if len(index(begin(true))) != 0 {
		t.Fatal("forced build reused modules")
	}
	oldContext := next.Context
	spec.Options = append(spec.Options, Setting{Name: "other-option", Value: "changed"})
	if begin(false).Context == oldContext {
		t.Fatal("module context omitted compiler options")
	}
	put(t, filepath.Join(spec.Output, ".gopurs-cache/objects", candidates[0].Parts[0].Object.SHA256), "corrupt")
	if len(index(begin(false))) != 0 {
		t.Fatal("corrupt module remained eligible")
	}
	bytes, err := os.ReadFile(filepath.Join(next.Stage, "prior/Main/0.json"))
	if err != nil || string(bytes) != `["unit"]` {
		t.Fatalf("candidate depends on live cache: %s %v", bytes, err)
	}
}

func TestPreparedFingerprintInputProof(t *testing.T) {
	for _, change := range []string{"ffi", "force", "tast", "compiler", "diagnostics"} {
		t.Run(change, func(t *testing.T) {
			spec, result, state := driverFixture(t)
			finishDriver(t, &result, &state, []SourceModule{{Name: "Main", Path: "src/Main.purs", HasForeign: true}}, change == "diagnostics")
			// Every case misses the whole-build cache; only the FFI-only case may
			// prove that preparation still sees the identical complete TAST input.
			put(t, filepath.Join(spec.Workspace, "src/Main.go"), "func ChangedFFI() {}")
			if change == "tast" {
				put(t, filepath.Join(spec.Output, "Main/corefn.json"), "changed caller or provider")
			}
			if change == "compiler" {
				put(t, absolute(spec.Workspace, spec.Compiler[0].Path), "changed compiler")
			}
			var next SessionResponse
			if err := beginSession(spec, change == "force", &next); err != nil {
				t.Fatal(err)
			}
			if next.Status != "miss" || next.SameTAST != (change == "ffi") {
				t.Fatalf("incorrect input proof: %+v", next)
			}
		})
	}
}

func TestBackendCodecUpgradePreservesOwnedOutputs(t *testing.T) {
	f := newFixture(t)
	put(t, filepath.Join(f.root, "stage/body.json"), `["unit"]`)
	mod := StagedModule{Name: "Main", Prepared: Key("prepared"), OptimizerEnvironment: Key("optimizer"), EmitterEnvironment: Key("emitter"), CallerDemands: Key("callers")}
	for i, codec := range moduleCodecs {
		if i == 0 {
			codec = "pbo/backend-v1"
		}
		mod.Parts = append(mod.Parts, StagedArtifact{Codec: codec, Source: "stage/body.json"})
	}
	f.req.Modules = []StagedModule{mod}
	publish(t, f)
	lookup(t, f, "hit")
	lockPath := filepath.Join(f.s.cache, "write.lock")
	before, err := os.Stat(lockPath)
	if err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(f.root, "compiler"), "compiler using backend-v2")
	put(t, filepath.Join(f.root, "stage/Main.go"), "package purescript\n// changed, still compiler-owned\n")
	f.req.Modules[0].Parts[0].Codec = "pbo/backend-v2"
	lookup(t, f, "miss")
	result := publish(t, f)
	if result.Manifest.Modules[0].Parts[0].Codec != "pbo/backend-v2" {
		t.Fatal("backend codec upgrade was not committed")
	}
	lookup(t, f, "hit")
	after, err := os.Stat(lockPath)
	if err != nil || !os.SameFile(before, after) {
		t.Fatalf("codec migration replaced the stable lock: %v", err)
	}
	f.req.Modules[0].Parts[0].Codec = "pbo/backend-v3"
	f.req.Operation, f.req.ExpectedKey = "publish", result.Snapshot.Key
	if _, err := execute(f.req); err == nil {
		t.Fatal("unknown backend version was accepted")
	}
	lookup(t, f, "hit")
}

func TestWorkerCodecUpgradePreservesOwnedOutputs(t *testing.T) {
	f := newFixture(t)
	put(t, filepath.Join(f.root, "stage/workers.json"), `["unit"]`)
	mod := StagedModule{Name: "Main", Prepared: Key("prepared"), OptimizerEnvironment: Key("optimizer"), EmitterEnvironment: Key("emitter"), CallerDemands: Key("callers")}
	for i, codec := range moduleCodecs {
		if i == 3 {
			codec = "gopurs/workers-v1"
		}
		mod.Parts = append(mod.Parts, StagedArtifact{Codec: codec, Source: "stage/workers.json"})
	}
	f.req.Modules = []StagedModule{mod}
	publish(t, f)
	lookup(t, f, "hit")
	before, err := os.Stat(filepath.Join(f.s.cache, "write.lock"))
	if err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(f.root, "compiler"), "compiler using owned workers-v2")
	// This body differs from the old output: migration must use the readable
	// ownership journal, rather than treating it as an unowned legacy file.
	put(t, filepath.Join(f.root, "stage/Main.go"), "package purescript\n// owned worker\n")
	f.req.Modules[0].Parts[3].Codec = "gopurs/workers-v2"
	lookup(t, f, "miss")
	result := publish(t, f)
	if result.Manifest.Modules[0].Parts[3].Codec != "gopurs/workers-v2" {
		t.Fatal("worker codec upgrade was not committed")
	}
	lookup(t, f, "hit")
	after, err := os.Stat(filepath.Join(f.s.cache, "write.lock"))
	if err != nil || !os.SameFile(before, after) {
		t.Fatalf("worker migration replaced the stable lock: %v", err)
	}
	f.req.Modules[0].Parts[3].Codec = "gopurs/workers-v3"
	f.req.Operation, f.req.ExpectedKey = "publish", result.Snapshot.Key
	if _, err := execute(f.req); err == nil {
		t.Fatal("unknown worker version was accepted")
	}
	lookup(t, f, "hit")
}
