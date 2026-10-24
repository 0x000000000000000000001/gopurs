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

type compactFixture struct {
	s       store
	objects []Object
	data    map[string][]byte
	packs   map[string][]byte
}

func newCompactFixture(t *testing.T) *compactFixture {
	t.Helper()
	root := t.TempDir()
	f := &compactFixture{
		s: store{root: root, cache: filepath.Join(root, "cache"), packs: &packDB{
			entries: map[string]packEntry{}, files: map[string]*os.File{},
		}},
		data: map[string][]byte{}, packs: map[string][]byte{},
	}
	if err := os.MkdirAll(filepath.Join(f.s.cache, "packs"), 0700); err != nil {
		t.Fatal(err)
	}
	groups := map[int]bool{}
	for i := 0; len(f.objects) < 2; i++ {
		data := []byte(fmt.Sprintf("compaction object %d", i))
		obj := Object{SHA256: digest(data), Bytes: int64(len(data))}
		group := packGroupFor(obj.SHA256)
		if groups[group] {
			continue
		}
		groups[group] = true
		// An obsolete suffix makes each group require compaction.
		packed := append(append([]byte{}, data...), []byte(" retired bytes")...)
		pack := digest(packed)
		if err := os.WriteFile(filepath.Join(f.s.cache, "packs", pack+".pack"), packed, 0600); err != nil {
			t.Fatal(err)
		}
		f.objects = append(f.objects, obj)
		f.data[obj.SHA256] = data
		f.packs[pack] = packed
		f.s.packs.entries[obj.SHA256] = packEntry{Pack: pack, Bytes: obj.Bytes}
	}
	sort.Slice(f.objects, func(i, j int) bool {
		return packGroupFor(f.objects[i].SHA256) < packGroupFor(f.objects[j].SHA256)
	})
	if err := f.s.writePackIndex(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		for _, file := range f.s.packs.files {
			_ = file.Close()
		}
	})
	return f
}

func (f *compactFixture) indexBytes(t *testing.T) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(f.s.cache, "packs", "index.json"))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func (f *compactFixture) assertPriorIndex(t *testing.T, before []byte) {
	t.Helper()
	if !bytes.Equal(before, f.indexBytes(t)) {
		t.Fatal("failed compaction changed the disk index")
	}
	var prior packIndexFile
	if err := json.Unmarshal(before, &prior); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(prior.Objects, f.s.packs.entries) {
		t.Fatal("failed compaction installed partial memory entries")
	}
	for pack := range f.packs {
		if _, err := os.Stat(filepath.Join(f.s.cache, "packs", pack+".pack")); err != nil {
			t.Fatalf("old pack was removed on failure: %v", err)
		}
	}
}

func (f *compactFixture) corrupt(t *testing.T, object Object) string {
	t.Helper()
	pack := f.s.packs.entries[object.SHA256].Pack
	data := append([]byte{}, f.packs[pack]...)
	data[0] ^= 1
	if err := os.WriteFile(filepath.Join(f.s.cache, "packs", pack+".pack"), data, 0600); err != nil {
		t.Fatal(err)
	}
	return pack
}

func (f *compactFixture) assertObjects(t *testing.T) {
	t.Helper()
	var index packIndexFile
	if err := json.Unmarshal(f.indexBytes(t), &index); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(index.Objects, f.s.packs.entries) {
		t.Fatal("committed and memory indices differ")
	}
	for _, obj := range f.objects {
		data, err := f.s.validateObject(obj)
		if err != nil || !bytes.Equal(data, f.data[obj.SHA256]) {
			t.Fatalf("object not recoverable: %v", err)
		}
	}
}

func TestCompactFailureJoinsAndRefillRetries(t *testing.T) {
	t.Setenv(packPublication, "1")
	f := newCompactFixture(t)
	before := f.indexBytes(t)
	bad := f.objects[1]
	pack := f.corrupt(t, bad)
	changed, err := f.s.compactPackGroups()
	if changed || err == nil || !strings.Contains(err.Error(), bad.SHA256) {
		t.Fatalf("corruption was not reported: %v, %v", changed, err)
	}
	f.assertPriorIndex(t, before)
	if f.s.packs.files[pack] != nil {
		t.Fatal("failed handle retained after join")
	}
	// The valid group must finish despite another group's error. Its pack is
	// not indexed, but exists after return: all workers have joined.
	good := f.objects[0]
	if !fileMatches(filepath.Join(f.s.cache, "packs", good.SHA256+".pack"), good) {
		t.Fatal("valid group did not finish before error return")
	}
	writer := newPackWriter(&f.s)
	if _, err := writer.put(f.data[bad.SHA256]); err != nil {
		t.Fatal(err)
	}
	if err := writer.flush(); err != nil {
		t.Fatal(err)
	}
	if changed, err := f.s.compactPackGroups(); err != nil || !changed {
		t.Fatalf("refilled compaction did not retry: %v, %v", changed, err)
	}
	f.assertObjects(t)
}

func TestCompactAtomicRepairDoesNotEvict(t *testing.T) {
	f := newCompactFixture(t)
	before := f.indexBytes(t)
	bad := f.objects[1]
	pack := f.corrupt(t, bad)
	if _, err := f.s.validateObject(bad); err == nil {
		t.Fatal("corruption was not observed")
	}
	stale, err := f.s.packs.files[pack].Stat()
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(f.s.cache, "packs", pack+".pack")
	if err := atomicReplaceFile(path, f.packs[pack], 0600); err != nil {
		t.Fatal(err)
	}
	repaired, err := os.Stat(path)
	if err != nil || os.SameFile(stale, repaired) {
		t.Fatalf("test did not replace the cached inode: %v", err)
	}
	// The old descriptor still sees corruption after atomic repair. Returning
	// its error must not evict the now-valid pack or its committed references.
	if changed, err := f.s.compactPackGroups(); changed || err == nil {
		t.Fatalf("stale descriptor was not exercised: %v, %v", changed, err)
	}
	f.assertPriorIndex(t, before)
	if !fileMatches(path, Object{SHA256: pack, Bytes: int64(len(f.packs[pack]))}) {
		t.Fatal("atomic repair was damaged")
	}
	if changed, err := f.s.compactPackGroups(); err != nil || !changed {
		t.Fatalf("retry did not reopen repaired pack: %v, %v", changed, err)
	}
	f.assertObjects(t)
}

func TestCompactIndexFailureDoesNotInstall(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("directory permissions do not prevent root writes")
	}
	f := newCompactFixture(t)
	before := f.indexBytes(t)
	directory := filepath.Join(f.s.cache, "packs")
	// Complete target packs already exist, so the denied write is the index
	// publication, after all group reads have successfully joined.
	for _, obj := range f.objects {
		if err := os.WriteFile(filepath.Join(directory, obj.SHA256+".pack"), f.data[obj.SHA256], 0600); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.Chmod(directory, 0500); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(directory, 0700) })
	changed, err := f.s.compactPackGroups()
	if changed || err == nil || !os.IsPermission(err) {
		t.Fatalf("index publication failure not exercised: %v, %v", changed, err)
	}
	f.assertPriorIndex(t, before)
	if err := os.Chmod(directory, 0700); err != nil {
		t.Fatal(err)
	}
	if changed, err := f.s.compactPackGroups(); err != nil || !changed {
		t.Fatalf("index publication retry failed: %v, %v", changed, err)
	}
	f.assertObjects(t)
}

func TestCompactErrorsFollowGroupOrderAndKeepCloseError(t *testing.T) {
	f := newCompactFixture(t)
	before := f.indexBytes(t)
	for _, obj := range f.objects {
		f.corrupt(t, obj)
	}
	if _, err := f.s.compactPackGroups(); err == nil || !strings.Contains(err.Error(), f.objects[0].SHA256) {
		t.Fatalf("first group error not preserved: %v", err)
	}
	f.assertPriorIndex(t, before)
	obj := f.objects[0]
	pack := f.s.packs.entries[obj.SHA256].Pack
	file, err := f.s.packFile(pack)
	if err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}
	// Inject an already-closed cached descriptor: report the primary read
	// failure and the close error, while clearing the failed cache handle.
	if _, err := f.s.compactPackGroups(); err == nil || !strings.Contains(err.Error(), obj.SHA256) || !strings.Contains(err.Error(), "close compacted pack") {
		t.Fatalf("read/close error lost: %v", err)
	}
	f.assertPriorIndex(t, before)
	if f.s.packs.files[pack] != nil {
		t.Fatal("closed failed handle was retained")
	}
}

func TestCompactReopensCachedTargetAfterRepair(t *testing.T) {
	for _, repairedBefore := range []bool{false, true} {
		t.Run(fmt.Sprintf("repaired-before-%v", repairedBefore), func(t *testing.T) {
			f := newCompactFixture(t)
			obj := f.objects[0]
			path := filepath.Join(f.s.cache, "packs", obj.SHA256+".pack")
			corrupt := append([]byte{}, f.data[obj.SHA256]...)
			corrupt[0] ^= 1
			if err := os.WriteFile(path, corrupt, 0600); err != nil {
				t.Fatal(err)
			}
			// This descriptor can survive from an earlier store operation; the
			// target is not one of the current compaction source packs.
			file, err := f.s.packFile(obj.SHA256)
			if err != nil {
				t.Fatal(err)
			}
			if repairedBefore {
				if err := atomicReplaceFile(path, f.data[obj.SHA256], 0600); err != nil {
					t.Fatal(err)
				}
			}
			if changed, err := f.s.compactPackGroups(); err != nil || !changed {
				t.Fatalf("target repair failed: %v, %v", changed, err)
			}
			if _, err := file.Stat(); err == nil {
				t.Fatal("stale target descriptor remained open")
			}
			f.assertObjects(t)
		})
	}
}
