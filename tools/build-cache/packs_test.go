package main

import (
	"os"
	"path/filepath"
	"testing"
)

func packedFixture(t *testing.T) fixture {
	t.Helper()
	f := newFixture(t)
	put(t, filepath.Join(f.root, "stage/unit.json"), `["unit"]`)
	put(t, filepath.Join(f.root, "stage/module.go"), "package purescript\n// module\n")
	mod := StagedModule{Name: "Main", Prepared: Key("prepared"), OptimizerEnvironment: Key("optimizer"), EmitterEnvironment: Key("emitter"), CallerDemands: Key("callers")}
	for _, codec := range moduleCodecs {
		mod.Parts = append(mod.Parts, StagedArtifact{Codec: codec, Source: "stage/unit.json"})
	}
	mod.Sources = []StagedModuleSource{{Kind: "module", Source: "stage/module.go"}}
	f.req.Modules = []StagedModule{mod}
	return f
}

func readPackedObject(t *testing.T, f fixture, obj Object) []byte {
	t.Helper()
	s := store{root: f.s.root, cache: f.s.cache}
	if err := s.ensurePacks(); err != nil {
		t.Fatal(err)
	}
	data, err := s.readObject(obj)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func readPackIndex(t *testing.T, f fixture) packIndexFile {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(f.s.cache, "packs/index.json"))
	if err != nil {
		t.Fatal(err)
	}
	var index packIndexFile
	if err := strictJSON(data, &index); err != nil {
		t.Fatal(err)
	}
	return index
}

func TestPackedPublicationRoundTripAndCorruption(t *testing.T) {
	t.Setenv(packPublication, "1")
	f := packedFixture(t)
	result := publish(t, f)

	index := readPackIndex(t, f)
	unique := map[string]bool{}
	for _, artifact := range artifacts(*result.Manifest) {
		unique[artifact.Object.SHA256] = true
	}
	if len(index.Objects) != len(unique) {
		t.Fatalf("packed objects: %d, unique artifacts: %d", len(index.Objects), len(unique))
	}
	for _, artifact := range artifacts(*result.Manifest) {
		if _, ok := index.Objects[artifact.Object.SHA256]; !ok {
			t.Fatalf("artifact not packed: %s", artifact.Object.SHA256)
		}
		data := readPackedObject(t, f, artifact.Object)
		// Module sources are compressed; only the PBO parts have this payload.
		if artifact.Codec != moduleSourceCodec && string(data) != `["unit"]` {
			t.Fatalf("packed artifact %s: %q", artifact.Codec, data)
		}
	}
	lookup(t, f, "hit")

	first := result.Manifest.Modules[0].Parts[0]
	entry := index.Objects[first.Object.SHA256]
	packPath := filepath.Join(f.s.cache, "packs", entry.Pack+".pack")
	if err := os.WriteFile(packPath, []byte("corrupt"), 0600); err != nil {
		t.Fatal(err)
	}
	lookup(t, f, "miss")

	// The failed read invalidates the pack in the persisted index; a rebuild
	// must create a fresh pack rather than silently reusing the corrupt one.
	index = readPackIndex(t, f)
	if _, ok := index.Objects[first.Object.SHA256]; ok {
		t.Fatal("corrupt pack entry survived invalidation")
	}
	rebuilt := publish(t, f)
	for _, artifact := range artifacts(*rebuilt.Manifest) {
		readPackedObject(t, f, artifact.Object)
	}
	lookup(t, f, "hit")
}

func TestPackedModuleSourcesFreeze(t *testing.T) {
	t.Setenv(packPublication, "1")
	_, result, state, sources := packageModuleFixture(t)
	next := t.TempDir()
	if err := os.Mkdir(filepath.Join(next, "prior"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := freezeModuleCache(next, state, false, &result); err != nil {
		t.Fatal(err)
	}
	for kind, want := range sources {
		data, err := os.ReadFile(filepath.Join(next, "prior/Main", kind+".go"))
		if err != nil || string(data) != want {
			t.Fatalf("frozen %s: %q %v", kind, data, err)
		}
	}
}

func TestPackPublicationDisabledKeepsIndividualObjects(t *testing.T) {
	t.Setenv(packPublication, "0")
	f := packedFixture(t)
	result := publish(t, f)
	if _, err := os.Stat(filepath.Join(f.s.cache, "packs/index.json")); !os.IsNotExist(err) {
		t.Fatalf("pack index written while disabled: %v", err)
	}
	for _, artifact := range artifacts(*result.Manifest) {
		if _, err := os.Stat(filepath.Join(f.s.cache, "objects", artifact.Object.SHA256)); err != nil {
			t.Fatalf("individual artifact missing: %v", err)
		}
	}
	lookup(t, f, "hit")
}
