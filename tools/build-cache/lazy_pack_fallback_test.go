package main

import (
	"os"
	"path/filepath"
	"testing"
)

// Exercise the storage steps used by the driver's automatic full fallback.
// The producer's fresh bytes can be identical after an unrelated TAST edit;
// their presence in the old pack index must not bypass integrity checks.
func TestLazyPackFailureFullRebuildRepairsIndexedObject(t *testing.T) {
	for _, testCase := range []struct {
		name  string
		loose bool
	}{{"packed-only", false}, {"valid-loose-copy", true}} {
		t.Run(testCase.name, func(t *testing.T) {
			t.Setenv(packPublication, "1")
			f := packedFixture(t)
			f.req.Spec.Options = append(f.req.Spec.Options, Setting{"module-lazy", "1"})
			first := publish(t, f)
			part := first.Manifest.Modules[0].Parts[1].Object
			index := readPackIndex(t, f)
			entry, ok := index.Objects[part.SHA256]
			if !ok {
				t.Fatal("fixture did not pack the implementation")
			}

			// Corrupt only this object's bytes, preserving the pack's size and all
			// neighboring objects, in particular any compressed module sources.
			pack, err := os.OpenFile(filepath.Join(f.s.cache, "packs", entry.Pack+".pack"), os.O_RDWR, 0)
			if err != nil {
				t.Fatal(err)
			}
			byteAtOffset := []byte{0}
			if _, err := pack.ReadAt(byteAtOffset, entry.Offset); err != nil {
				_ = pack.Close()
				t.Fatal(err)
			}
			byteAtOffset[0] ^= 1
			if _, err := pack.WriteAt(byteAtOffset, entry.Offset); err != nil {
				_ = pack.Close()
				t.Fatal(err)
			}
			if err := pack.Close(); err != nil {
				t.Fatal(err)
			}
			if testCase.loose {
				// A prior non-packed publication can leave this valid duplicate.
				put(t, filepath.Join(f.s.cache, "objects", part.SHA256), `["unit"]`)
			}
			put(t, filepath.Join(f.s.root, "Main/corefn.json"), `{"typeTable":[],"moduleName":["Main"],"edited":true}`)
			lookupRequest := f.req
			lookupRequest.Operation = "lookup"
			miss := call(t, lookupRequest, "miss")
			if miss.Reason != "inputs-changed" {
				t.Fatalf("test must take the pre-validation lookup exit: %s", miss.Reason)
			}

			state := buildSession{Spec: f.req.Spec, Snapshot: *miss.Snapshot, Count: 1}
			state.Spec.Output = f.s.root
			newStage := func(force bool) string {
				t.Helper()
				stage := t.TempDir()
				if err := os.Mkdir(filepath.Join(stage, "prior"), 0700); err != nil {
					t.Fatal(err)
				}
				if err := freezeModuleCache(stage, state, force, &SessionResponse{}); err != nil {
					t.Fatal(err)
				}
				t.Cleanup(func() { forgetLazyPrior(stage) })
				return stage
			}
			lazy := newStage(false)
			if err := materializePrior(lazy, state, []string{"Main"}, []string{"1"}); err == nil {
				t.Fatal("lazy materialization accepted the corrupted implementation")
			}
			if _, ok := readPackIndex(t, f).Objects[part.SHA256]; !ok {
				t.Fatal("fixture unexpectedly repaired the public pack before fallback")
			}
			forced := newStage(true)
			data, err := os.ReadFile(filepath.Join(forced, "prior/index.json"))
			if err != nil || string(data) != "[]" {
				t.Fatalf("forced build retained module candidates: %q %v", data, err)
			}

			// Re-publish the producer's correct bytes through the actual pack writer,
			// as a full reconstruction does. Validate before lookup: lookup itself
			// would invalidate the broken index and mask the failed repair.
			rebuilt := publish(t, f)
			s := store{root: f.s.root, cache: f.s.cache}
			if err := s.ensurePacks(); err != nil {
				t.Fatal(err)
			}
			for _, artifact := range artifacts(*rebuilt.Manifest) {
				if _, err := s.validateObject(artifact.Object); err != nil {
					t.Fatalf("full fallback published a corrupt %s artifact: %v", artifact.Codec, err)
				}
			}
			lookup(t, f, "hit")
		})
	}
}

func TestPackRepairReopensReplacedPackInSameStore(t *testing.T) {
	t.Setenv(packPublication, "1")
	f := packedFixture(t)
	// One object makes the repaired pack retain the same content address, so
	// atomic replacement must invalidate the open handle to the corrupt inode.
	f.req.Modules[0].Sources = nil
	first := publish(t, f)
	object := first.Manifest.Modules[0].Parts[0].Object
	s := store{root: f.s.root, cache: f.s.cache}
	if err := s.ensurePacks(); err != nil {
		t.Fatal(err)
	}
	good, err := s.validateObject(object)
	if err != nil {
		t.Fatal(err)
	}
	entry := s.packs.entries[object.SHA256]
	if entry.Pack != object.SHA256 {
		t.Fatal("fixture must contain exactly one packed object")
	}
	corrupt := append([]byte{}, good...)
	corrupt[0] ^= 1
	if err := os.WriteFile(filepath.Join(s.cache, "packs", entry.Pack+".pack"), corrupt, 0600); err != nil {
		t.Fatal(err)
	}
	writer := newPackWriter(&s)
	if _, err := writer.put(good); err != nil {
		t.Fatal(err)
	}
	if err := writer.flush(); err != nil {
		t.Fatal(err)
	}
	if _, err := s.validateObject(object); err != nil {
		t.Fatalf("same store read the old pack inode after repair: %v", err)
	}
}
