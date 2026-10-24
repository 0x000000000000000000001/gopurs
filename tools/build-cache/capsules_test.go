package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func capsuleBytes(t *testing.T, context, reader, implementation string) []byte {
	t.Helper()
	absent := false
	value := bindingCapsule{Codec: bindingCapsuleCodec, Context: context, Reader: reader,
		Anchors:  []capsuleAnchor{{reader, digest([]byte(implementation))}},
		Coverage: []capsuleCoverage{{reader, "missing", &absent}}, Payload: `["array",[]]`}
	value.PayloadSHA256 = digest([]byte(value.Payload))
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func TestCapsuleEnvelopeRejectsChangedProofs(t *testing.T) {
	context := Key("context")
	data := capsuleBytes(t, context, "Main", `["unit"]`)
	lookup := func(name string) string {
		if name == "Main" {
			return digest([]byte(`["unit"]`))
		}
		return ""
	}
	if !validateCapsule(data, context, "Main", lookup) {
		t.Fatal("valid negative coverage refused")
	}
	for name, change := range map[string]func(*bindingCapsule){
		"context":          func(c *bindingCapsule) { c.Context = Key("other") },
		"reader":           func(c *bindingCapsule) { c.Reader = "Other" },
		"codec":            func(c *bindingCapsule) { c.Codec += "-old" },
		"changed-anchor":   func(c *bindingCapsule) { c.Anchors[0].SHA256 = Key("other") },
		"missing-owner":    func(c *bindingCapsule) { c.Anchors[0].Owner = "Other" },
		"duplicate-anchor": func(c *bindingCapsule) { c.Anchors = append(c.Anchors, c.Anchors[0]) },
		"duplicate-query":  func(c *bindingCapsule) { c.Coverage = append(c.Coverage, c.Coverage[0]) },
		"unanchored-query": func(c *bindingCapsule) { c.Coverage[0].Owner = "Other" },
		"unused-anchor":    func(c *bindingCapsule) { c.Coverage = []capsuleCoverage{} },
		"missing-presence": func(c *bindingCapsule) { c.Coverage[0].Present = nil },
		"payload-hash":     func(c *bindingCapsule) { c.PayloadSHA256 = Key("other") },
		"payload-shape":    func(c *bindingCapsule) { c.Payload = `{}`; c.PayloadSHA256 = digest([]byte(c.Payload)) },
	} {
		t.Run(name, func(t *testing.T) {
			var value bindingCapsule
			if err := json.Unmarshal(data, &value); err != nil {
				t.Fatal(err)
			}
			change(&value)
			bad, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			if validateCapsule(bad, context, "Main", lookup) {
				t.Fatal("accepted changed capsule proof")
			}
		})
	}
}

func beginCapsuleSession(t *testing.T, spec Spec, force bool) (SessionResponse, buildSession) {
	t.Helper()
	var result SessionResponse
	if err := beginSession(spec, force, &result); err != nil {
		t.Fatal(err)
	}
	if result.Status != "miss" || result.Stage == "" {
		t.Fatalf("expected producer session: %+v", result)
	}
	t.Cleanup(func() { forgetLazyPrior(result.Stage); _ = os.RemoveAll(result.Stage) })
	data, err := os.ReadFile(filepath.Join(result.Stage, "session.json"))
	if err != nil {
		t.Fatal(err)
	}
	var state buildSession
	if err := strictJSON(data, &state); err != nil {
		t.Fatal(err)
	}
	return result, state
}

func finishCapsuleSession(t *testing.T, result SessionResponse, state buildSession, implementation string, capsule []byte) *Manifest {
	t.Helper()
	if err := planSession(result.Stage, &state, []SourceModule{{Name: "Main", Path: "src/Main.purs"}}, &result); err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(result.Output, "gopurs_runtime/runtime.go"), "package gopurs_runtime\n")
	put(t, filepath.Join(result.Output, "purescript/Main.go"), "package purescript\n")
	memo := ModuleMemo{Name: "Main", Prepared: Key("prepared"), OptimizerEnvironment: Key("optimizer"), EmitterEnvironment: Key("emitter"), CallerDemands: Key("callers")}
	data, _ := json.Marshal(memo)
	dir := filepath.Join(result.Stage, "modules/Main")
	put(t, filepath.Join(dir, "memo.json"), string(data))
	for _, name := range []string{"0.json", "1.json", "2.json", "3.json"} {
		put(t, filepath.Join(dir, name), `["unit"]`)
	}
	put(t, filepath.Join(dir, "1.json"), implementation)
	if capsule != nil {
		put(t, filepath.Join(dir, "capsule.json"), string(capsule))
	}
	if err := publishSession(result.Stage, state, SessionRequest{ModuleCache: true}, &result); err != nil {
		t.Fatal(err)
	}
	s := store{root: state.Spec.Output, cache: filepath.Join(state.Spec.Output, ".gopurs-cache")}
	manifest, reason := s.readManifest()
	if manifest == nil {
		t.Fatal(reason)
	}
	return manifest
}

func TestCapsuleSurvivesForcedSessionAndRejectsChangedGeneration(t *testing.T) {
	for _, changed := range []string{"same", "implementation", "context"} {
		t.Run(changed, func(t *testing.T) {
			f := newFixture(t)
			spec := f.req.Spec
			spec.Output = filepath.Join(f.root, "output")
			spec.Inputs = []Input{{Name: "tast", Kind: "tast", Path: spec.Output}}
			spec.Options = append(spec.Options, Setting{"module-lazy", "1"})
			prime, state := beginCapsuleSession(t, spec, false)
			capsule := capsuleBytes(t, prime.Context, "Main", `["unit"]`)
			before := finishCapsuleSession(t, prime, state, `["unit"]`, capsule)
			if before.Modules[0].Capsule == nil {
				t.Fatal("fresh capsule not committed")
			}
			if changed == "context" {
				spec.Options = append(spec.Options, Setting{"changed", "1"})
			}
			full, fullState := beginCapsuleSession(t, spec, true)
			index, err := os.ReadFile(filepath.Join(full.Stage, "prior/index.json"))
			if err != nil || string(index) != "[]" {
				t.Fatalf("forced full exposed candidates: %s %v", index, err)
			}
			implementation := `["unit"]`
			if changed == "implementation" {
				implementation = `["int","9"]`
			}
			after := finishCapsuleSession(t, full, fullState, implementation, nil)
			if changed != "same" {
				if after.Modules[0].Capsule != nil {
					t.Fatal("capsule survived changed generation")
				}
				return
			}
			if !reflect.DeepEqual(before.Modules[0].Capsule, after.Modules[0].Capsule) {
				t.Fatal("forced full lost exact prior capsule")
			}
			// A later cache miss obtains the retained reference from a fresh private
			// snapshot. A forged sidecar cannot replace its committed object.
			put(t, filepath.Join(spec.Output, "Main/corefn.json"), `{"changed":true}`)
			next, nextState := beginCapsuleSession(t, spec, false)
			put(t, filepath.Join(next.Stage, "prior/Main/capsule.json"), "forged sidecar")
			var read SessionResponse
			if err := readSessionCapsule(next.Stage, nextState, SessionRequest{Owners: []string{"Main"}}, &read); err != nil {
				t.Fatal(err)
			}
			if read.Capsule == nil || *read.Capsule != string(capsule) {
				t.Fatal("private committed capsule changed")
			}
			prior, err := getLazyPrior(next.Stage, nextState)
			if err != nil {
				t.Fatal(err)
			}
			object := after.Modules[0].Capsule.Object
			entry, packed := prior.store.packs.entries[object.SHA256]
			public := filepath.Join(spec.Output, ".gopurs-cache", "objects", object.SHA256)
			if packed {
				public = filepath.Join(spec.Output, ".gopurs-cache", "packs", entry.Pack+".pack")
			}
			if err := os.Remove(public); err != nil {
				t.Fatal(err)
			}
			forgetLazyPrior(next.Stage)
			read = SessionResponse{}
			cwd, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			if err := os.Chdir(spec.Workspace); err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { _ = os.Chdir(cwd) })
			request, _ := json.Marshal(SessionRequest{Operation: "read-capsule", Stage: next.Stage, Owners: []string{"Main"}})
			if err := json.Unmarshal([]byte(SessionExchange(string(request))), &read); err != nil || read.Status != "ready" || read.Capsule == nil || *read.Capsule != string(capsule) {
				t.Fatalf("helper reconstruction: %+v %v", read, err)
			}
			prior, err = getLazyPrior(next.Stage, nextState)
			if err != nil {
				t.Fatal(err)
			}
			entriesBefore := len(prior.store.packs.entries)
			if packed {
				file, err := os.OpenFile(filepath.Join(prior.store.cache, "packs", entry.Pack+".pack"), os.O_WRONLY, 0600)
				if err != nil {
					t.Fatal(err)
				}
				_, err = file.WriteAt([]byte("!"), entry.Offset)
				_ = file.Close()
				if err != nil {
					t.Fatal(err)
				}
				// A loose duplicate never overrides an authoritative indexed pack.
				put(t, filepath.Join(prior.store.cache, "objects", object.SHA256), string(capsule))
			} else if err := os.Remove(filepath.Join(prior.store.cache, "objects", object.SHA256)); err != nil {
				t.Fatal(err)
			}
			read = SessionResponse{}
			if err := readSessionCapsule(next.Stage, nextState, SessionRequest{Owners: []string{"Main"}}, &read); err != nil || read.Capsule != nil {
				t.Fatalf("optional corruption must be a local miss: %+v %v", read, err)
			}
			if len(prior.store.packs.entries) != entriesBefore {
				t.Fatal("optional read invalidated a pack")
			}
		})
	}
}

func TestInvalidCapsuleReferenceDoesNotPoisonCandidateIndex(t *testing.T) {
	_, _, state, _ := packageModuleFixture(t)
	state.Spec.Options = append(state.Spec.Options, Setting{"module-lazy", "1"})
	s := store{root: state.Spec.Output, cache: filepath.Join(state.Spec.Output, ".gopurs-cache")}
	if err := s.ensurePacks(); err != nil {
		t.Fatal(err)
	}
	manifest, reason := s.readManifest()
	if manifest == nil {
		t.Fatal(reason)
	}
	manifest.Modules[0].Capsule = &Artifact{Codec: bindingCapsuleCodec, Object: Object{SHA256: "../../untrusted", Bytes: -1}}
	if err := validateManifest(*manifest); err != nil {
		t.Fatalf("optional invalid reference poisoned manifest: %v", err)
	}
	stage := t.TempDir()
	if err := os.Mkdir(filepath.Join(stage, "prior"), 0700); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { forgetLazyPrior(stage) })
	if err := freezeLazyModuleCache(stage, state, manifest, s); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(stage, "prior/index.json"))
	if err != nil {
		t.Fatal(err)
	}
	var memos []ModuleMemo
	if err := strictJSON(data, &memos); err != nil || len(memos) != 1 || memos[0].Capsule != nil {
		t.Fatalf("invalid reference entered typed index: %s %v", data, err)
	}
}

func TestCapsuleCorruptionIsLocalAndRetainedByGC(t *testing.T) {
	for _, packed := range []string{"0", "1"} {
		t.Run("packed-"+packed, func(t *testing.T) {
			t.Setenv(packPublication, packed)
			f := packedFixture(t)
			snap := call(t, f.req, "snapshot").Snapshot
			context, err := moduleContext(f.req.Spec, *snap)
			if err != nil {
				t.Fatal(err)
			}
			capsule := capsuleBytes(t, context, "Main", `["unit"]`)
			path := filepath.Join(f.root, "stage/capsule.json")
			put(t, path, string(capsule))
			f.req.Modules[0].Capsule = &StagedArtifact{Codec: bindingCapsuleCodec, Source: path}
			result := publish(t, f)
			part := result.Manifest.Modules[0].Capsule
			if part == nil {
				t.Fatal("capsule not committed")
			}
			s := store{root: f.s.root, cache: f.s.cache}
			if err := s.ensurePacks(); err != nil {
				t.Fatal(err)
			}
			if err := s.retirePackObjects(result.Manifest); err != nil {
				t.Fatal(err)
			}
			if err := s.retirePackageObjects(); err != nil {
				t.Fatal(err)
			}
			if actual, err := s.validateObject(part.Object); err != nil || string(actual) != string(capsule) {
				t.Fatalf("GC lost capsule: %v", err)
			}
			// The optional object can fail without invalidating mandatory parts or
			// discarding other entries in its indexed pack.
			corrupt := func() {
				s := store{root: f.s.root, cache: f.s.cache}
				if err := s.ensurePacks(); err != nil {
					t.Fatal(err)
				}
				if entry, exists := s.packs.entries[part.Object.SHA256]; exists {
					file, err := os.OpenFile(filepath.Join(s.cache, "packs", entry.Pack+".pack"), os.O_WRONLY, 0600)
					if err != nil {
						t.Fatal(err)
					}
					_, err = file.WriteAt([]byte("!"), entry.Offset)
					_ = file.Close()
					if err != nil {
						t.Fatal(err)
					}
				} else if err := os.Remove(filepath.Join(s.cache, "objects", part.Object.SHA256)); err != nil {
					t.Fatal(err)
				}
			}
			corrupt()
			lookup(t, f, "hit")
			repaired := publish(t, f)
			if !reflect.DeepEqual(part, repaired.Manifest.Modules[0].Capsule) {
				t.Fatal("fresh capsule was not recommitted after corruption")
			}
			if data, err := readPackedOrLoose(t, f, part.Object); err != nil || string(data) != string(capsule) {
				t.Fatalf("fresh capsule did not repair authoritative object: %v", err)
			}
			corrupt()
			f.req.Modules[0].Capsule = nil
			after := publish(t, f)
			if after.Manifest.Modules[0].Capsule != nil {
				t.Fatal("corrupt optional object carried forward")
			}
			for _, required := range artifacts(*after.Manifest) {
				if _, err := readPackedOrLoose(t, f, required.Object); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func readPackedOrLoose(t *testing.T, f fixture, object Object) ([]byte, error) {
	t.Helper()
	s := store{root: f.s.root, cache: f.s.cache}
	if err := s.ensurePacks(); err != nil {
		return nil, err
	}
	return s.validateObject(object)
}
