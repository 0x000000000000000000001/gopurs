package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

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
