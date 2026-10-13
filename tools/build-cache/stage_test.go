package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

func TestConcurrentStagingPreservesModuleOrderAndSharedObjects(t *testing.T) {
	f := newFixture(t)
	// Multiple workers address the same object. Module order deliberately differs
	// from output path order: it belongs to the producer's canonical schedule.
	const wire = `["unit"]`
	put(t, filepath.Join(f.root, "stage/shared.json"), wire)
	for i := 31; i >= 0; i-- {
		name := fmt.Sprintf("Mod%02d", i)
		mod := StagedModule{Name: name, Prepared: Key("prepared", name), OptimizerEnvironment: Key("pbo", name), EmitterEnvironment: Key("emitter", name), CallerDemands: Key("callers", name)}
		for _, codec := range moduleCodecs {
			mod.Parts = append(mod.Parts, StagedArtifact{Codec: codec, Source: "stage/shared.json"})
		}
		f.req.Modules = append(f.req.Modules, mod)
		f.req.Outputs = append(f.req.Outputs, StagedOutput{Path: "purescript/" + name + ".go", Kind: "module", Module: name, Source: "stage/Main.go"})
	}
	result := publish(t, f)
	for i, mod := range result.Manifest.Modules {
		if mod.Name != f.req.Modules[i].Name {
			t.Fatalf("completion reordered module %d: %s", i, mod.Name)
		}
		for _, part := range mod.Parts {
			data, err := f.s.readObject(part.Object)
			if err != nil || string(data) != wire {
				t.Fatalf("shared object not committed intact: %s %v", data, err)
			}
		}
	}
	lookup(t, f, "hit")
	// A late validation failure may leave private immutable objects, but must not
	// commit its changed outputs or ownership, even if other workers finish first.
	current := filepath.Join(f.s.cache, "current.json")
	journal := filepath.Join(f.s.root, ".gopurs-ownership.json")
	before, _ := os.ReadFile(current)
	owned, _ := os.ReadFile(journal)
	public, _ := os.ReadFile(filepath.Join(f.s.root, "purescript/Main.go"))
	put(t, filepath.Join(f.root, "stage/Main.go"), "package purescript\n// new output\n")
	put(t, filepath.Join(f.root, "stage/invalid.json"), `["invalid"]`)
	f.req.Modules[len(f.req.Modules)-1].Parts[3].Source = "stage/invalid.json"
	f.req.Operation = "publish"
	f.req.ExpectedKey = result.Snapshot.Key
	_, err := execute(f.req)
	if err == nil {
		t.Fatal("invalid part was accepted")
	}
	for path, want := range map[string][]byte{current: before, journal: owned, filepath.Join(f.s.root, "purescript/Main.go"): public} {
		got, err := os.ReadFile(path)
		if err != nil || !bytes.Equal(got, want) {
			t.Fatalf("failed staging changed %s: %v", path, err)
		}
	}
	entries, err := os.ReadDir(filepath.Join(f.s.cache, "objects"))
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if !digestPattern.MatchString(entry.Name()) {
			t.Fatalf("staging returned before its temporary files were cleaned: %s", entry.Name())
		}
	}
	lookup(t, f, "hit")
}

func TestFailedParallelOutputPublicationRecoversFromJournal(t *testing.T) {
	f := newFixture(t)
	committed := publish(t, f)
	pointerPath := filepath.Join(f.s.cache, "current.json")
	before, err := os.ReadFile(pointerPath)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 32; i++ {
		name := fmt.Sprintf("Mod%02d", i)
		source := "stage/" + name + ".go"
		put(t, filepath.Join(f.root, source), "package purescript\n// "+name+"\n")
		f.req.Outputs = append(f.req.Outputs, StagedOutput{Path: "purescript/" + name + ".go", Kind: "module", Module: name, Source: source})
	}
	lock, err := os.OpenFile(filepath.Join(f.s.cache, "write.lock"), os.O_RDWR, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer lock.Close()
	if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_EX); err != nil {
		t.Fatal(err)
	}
	staged, err := f.s.stage(f.req, *committed.Snapshot, committed.Manifest)
	if err != nil {
		t.Fatal(err)
	}
	if err := f.s.writeOwnership(mergeOwnership(committed.Manifest.Outputs, staged.Outputs)); err != nil {
		t.Fatal(err)
	}
	// Fail an early new output after journalling. Other workers and subsequent
	// items must all finish before the failure returns to the lock owner.
	var broken OwnedOutput
	for _, output := range staged.Outputs {
		if output.Path == "purescript/Mod00.go" {
			broken = output
		}
	}
	if broken.Path == "" {
		t.Fatal("missing failure target")
	}
	put(t, filepath.Join(f.s.cache, "objects", broken.Object.SHA256), "corrupt")
	if err := f.s.writeOutputs(staged.Outputs); err == nil {
		t.Fatal("corrupt object was published")
	}
	for _, output := range staged.Outputs {
		if output == broken {
			continue
		}
		if !fileMatches(filepath.Join(f.s.root, output.Path), output.Object) {
			t.Fatalf("publication returned before joining output %s", output.Path)
		}
	}
	if _, err := os.Lstat(filepath.Join(f.s.root, broken.Path)); !os.IsNotExist(err) {
		t.Fatal("failed output became public", err)
	}
	after, err := os.ReadFile(pointerPath)
	if err != nil || !bytes.Equal(before, after) {
		t.Fatal("failed publication committed a manifest", err)
	}
	if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_UN); err != nil {
		t.Fatal(err)
	}
	lookup(t, f, "miss")
	// The normal path repairs the object, accepts journal-owned partial outputs,
	// and commits the complete manifest. A second publication elides every write.
	repaired := publish(t, f)
	for _, output := range repaired.Manifest.Outputs {
		if !fileMatches(filepath.Join(f.s.root, output.Path), output.Object) {
			t.Fatal("incomplete recovery", output.Path)
		}
	}
	lookup(t, f, "hit")
}
