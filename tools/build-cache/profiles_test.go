package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"
)

func devFixture(f fixture) fixture {
	f.req.Spec.Options = append(append([]Setting{}, f.req.Spec.Options...), Setting{"optimization-profile", "dev-v4"})
	f.s.cache = filepath.Join(f.s.root, ".gopurs-cache", "profiles", "dev")
	return f
}

func TestProfileVersionChangesBuildAndModuleKeys(t *testing.T) {
	f := newFixture(t)
	versions := []string{"dev-v1", "dev-v2", "dev-v3", "dev-v4", "dev-v5"}
	var previous *Snapshot
	var previousContext string
	for _, version := range versions {
		spec := f.req.Spec
		spec.Output = f.s.root
		spec.Options = append(append([]Setting{}, spec.Options...), Setting{"optimization-profile", version})
		if cacheDirectory(spec) != filepath.Join(f.s.root, ".gopurs-cache", "profiles", "dev") {
			t.Fatalf("%s escaped the development namespace", version)
		}
		response, err := execute(Request{Schema: Schema, Operation: "snapshot", Spec: spec})
		if err != nil {
			t.Fatal(err)
		}
		context, err := moduleContext(spec, *response.Snapshot)
		if err != nil {
			t.Fatal(err)
		}
		if previous != nil {
			if previous.Compiler != response.Snapshot.Compiler || previous.Key == response.Snapshot.Key || previousContext == context {
				t.Fatal("policy version must change build and module keys even with the same compiler")
			}
		}
		previous, previousContext = response.Snapshot, context
	}
}

func TestProfilesRecoverOtherCommittedOwnership(t *testing.T) {
	for _, journal := range []string{"missing", "older-valid"} {
		t.Run(journal, func(t *testing.T) {
			release := newFixture(t)
			release.req.Outputs = []StagedOutput{{Path: "purescript/ReleaseOnly.go", Kind: "module", Module: "ReleaseOnly", Source: "stage/Main.go"}}
			dev := devFixture(release)
			dev.req.Outputs = []StagedOutput{{Path: "purescript/DevOnly.go", Kind: "module", Module: "DevOnly", Source: "stage/Main.go"}}
			publish(t, release)
			journalPath := filepath.Join(release.s.root, ".gopurs-ownership.json")
			oldJournal, err := os.ReadFile(journalPath)
			if err != nil {
				t.Fatal(err)
			}
			publish(t, dev)
			if _, err := os.Stat(filepath.Join(release.s.root, "purescript/ReleaseOnly.go")); !os.IsNotExist(err) {
				t.Fatal("development publication retained obsolete release output")
			}
			if journal == "missing" {
				if err := os.Remove(journalPath); err != nil {
					t.Fatal(err)
				}
			} else {
				put(t, journalPath, string(oldJournal))
			}
			publish(t, release)
			lookup(t, release, "hit")
			if _, err := os.Stat(filepath.Join(release.s.root, "purescript/DevOnly.go")); !os.IsNotExist(err) {
				t.Fatal("other profile's obsolete output survived ownership recovery")
			}
		})
	}
}

func TestProfileLookupDetectsUnjournaledOtherOutputs(t *testing.T) {
	for _, externalEdit := range []bool{false, true} {
		t.Run(map[bool]string{false: "restore", true: "external-conflict"}[externalEdit], func(t *testing.T) {
			release := newFixture(t)
			dev := devFixture(release)
			dev.req.Outputs = append(append([]StagedOutput{}, release.req.Outputs...), StagedOutput{
				Path: "purescript/DevOnly.go", Kind: "module", Module: "DevOnly", Source: "stage/Main.go",
			})
			publish(t, release)
			journalPath := filepath.Join(release.s.root, ".gopurs-ownership.json")
			oldJournal, err := os.ReadFile(journalPath)
			if err != nil {
				t.Fatal(err)
			}
			publish(t, dev)
			put(t, journalPath, string(oldJournal))
			extraPath := filepath.Join(release.s.root, "purescript/DevOnly.go")
			if externalEdit {
				put(t, extraPath, "external changes must be preserved")
			}
			if got := lookup(t, release, "miss"); got.Reason != "profile-outputs-changed" {
				t.Fatalf("stale journal authorized an incomplete release tree: %+v", got)
			}
			if externalEdit {
				req := release.req
				req.ExpectedKey = call(t, req, "snapshot").Snapshot.Key
				req.Operation = "publish"
				before, _ := os.ReadFile(filepath.Join(release.s.cache, "current.json"))
				_, err := execute(req)
				if err == nil || !strings.Contains(err.Error(), "obsolete output modified externally") {
					t.Fatalf("ambiguous obsolete bytes were not rejected: %v", err)
				}
				after, _ := os.ReadFile(filepath.Join(release.s.cache, "current.json"))
				journalAfter, _ := os.ReadFile(journalPath)
				extra, _ := os.ReadFile(extraPath)
				if !bytes.Equal(before, after) || !bytes.Equal(oldJournal, journalAfter) || string(extra) != "external changes must be preserved" {
					t.Fatal("conflicting publication modified ownership or external bytes")
				}
				return
			}
			put(t, filepath.Join(release.s.root, "purescript/User.go"), "external asset")
			publish(t, release)
			for i := 0; i < 3; i++ {
				lookup(t, release, "hit")
			}
			if _, err := os.Stat(extraPath); !os.IsNotExist(err) {
				t.Fatal("development-only output remains after recovery")
			}
			user, err := os.ReadFile(filepath.Join(release.s.root, "purescript/User.go"))
			if err != nil || string(user) != "external asset" {
				t.Fatal("recovery removed an unowned external file")
			}
		})
	}
}

func TestProfilesRecoverOtherBytesWhenRemovingSharedOutput(t *testing.T) {
	release := newFixture(t)
	dev := devFixture(release)
	dev.req.Outputs = append([]StagedOutput{}, release.req.Outputs...)
	dev.req.Outputs[0].Source = "stage/Dev.go"
	put(t, filepath.Join(release.root, "stage/Dev.go"), "package purescript\n// other profile bytes\n")
	publish(t, release)
	publish(t, dev)
	if err := os.Remove(filepath.Join(release.s.root, ".gopurs-ownership.json")); err != nil {
		t.Fatal(err)
	}
	release.req.Outputs = nil
	publish(t, release)
	lookup(t, release, "hit")
	if _, err := os.Stat(filepath.Join(release.s.root, "purescript/Main.go")); !os.IsNotExist(err) {
		t.Fatal("shared obsolete output was not recovered from the other profile's digest")
	}
}

func TestProfilesKeepIndependentGenerationsAndRestoreSharedOutput(t *testing.T) {
	release := newFixture(t)
	dev := devFixture(release)
	dev.req.Outputs = append([]StagedOutput{}, release.req.Outputs...)
	dev.req.Outputs[0].Source = "stage/Dev.go"
	put(t, filepath.Join(release.root, "stage/Dev.go"), "package purescript\n// development generation\n")
	releaseManifest := publish(t, release).Manifest
	releasePointer, err := os.ReadFile(filepath.Join(release.s.cache, "current.json"))
	if err != nil {
		t.Fatal(err)
	}
	lookup(t, release, "hit")
	devManifest := publish(t, dev).Manifest
	if releaseManifest.Snapshot.Key == devManifest.Snapshot.Key {
		t.Fatal("optimization profile missing from the snapshot")
	}
	lookup(t, dev, "hit")
	lookup(t, release, "miss")
	current, err := os.ReadFile(filepath.Join(release.s.cache, "current.json"))
	if err != nil || !bytes.Equal(releasePointer, current) {
		t.Fatal("development publication replaced the release generation", err)
	}
	for _, target := range []fixture{release, dev, release} {
		publish(t, target)
		lookup(t, target, "hit")
		want, _ := os.ReadFile(filepath.Join(target.root, target.req.Outputs[0].Source))
		got, err := os.ReadFile(filepath.Join(target.s.root, target.req.Outputs[0].Path))
		if err != nil || !bytes.Equal(got, want) {
			t.Fatal("profile switch did not restore the selected public bytes", err)
		}
	}
	// Cache damage in the inactive profile must neither invalidate nor poison
	// the active release build. The damaged profile must never become a hit.
	put(t, filepath.Join(dev.s.cache, "objects", devManifest.Outputs[0].Object.SHA256), "corrupt")
	if _, err := dev.s.readObject(devManifest.Outputs[0].Object); err == nil {
		t.Fatal("corrupted development object remains readable")
	}
	lookup(t, release, "hit")
	lookup(t, dev, "miss")
	publish(t, dev)
	lookup(t, dev, "hit")
	put(t, filepath.Join(dev.s.cache, "objects", devManifest.Outputs[0].Object.SHA256), "corrupt")
	if got := lookup(t, dev, "miss"); got.Reason != "object-corrupt" {
		t.Fatalf("active development profile failed to detect corruption: %+v", got)
	}
	if _, err := release.s.readObject(releaseManifest.Outputs[0].Object); err != nil {
		t.Fatal("development corruption damaged release object", err)
	}
}

func TestProfilesUseTheSamePublicLock(t *testing.T) {
	release := newFixture(t)
	call(t, release.req, "snapshot")
	dev := devFixture(release)
	lock, err := os.OpenFile(publicLockPath(release.s.root), os.O_RDWR, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer lock.Close()
	if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_EX); err != nil {
		t.Fatal(err)
	}
	defer syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)
	finished := make(chan error, 1)
	go func() { _, err := execute(dev.req); finished <- err }()
	select {
	case err := <-finished:
		t.Fatalf("development bypassed the shared lock: %v", err)
	case <-time.After(50 * time.Millisecond):
	}
	if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_UN); err != nil {
		t.Fatal(err)
	}
	select {
	case err := <-finished:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("development remained blocked after release unlocked")
	}
	if _, err := os.Stat(filepath.Join(dev.s.cache, "write.lock")); !os.IsNotExist(err) {
		t.Fatal("profile created a separate publication lock")
	}
}

func TestProfileSessionKeepsReleaseModuleCandidates(t *testing.T) {
	_, release, releaseState, sources := packageModuleFixture(t)
	devState := releaseState
	devState.Spec.Options = append(append([]Setting{}, releaseState.Spec.Options...), Setting{"optimization-profile", "dev-v4"})
	response, err := execute(Request{Schema: Schema, Operation: "snapshot", Spec: devState.Spec})
	if err != nil {
		t.Fatal(err)
	}
	devState.Snapshot = *response.Snapshot
	if err := publishSession(release.Stage, devState, SessionRequest{ModuleCache: true}, &SessionResponse{}); err != nil {
		t.Fatal(err)
	}
	for _, state := range []buildSession{releaseState, devState} {
		stage := t.TempDir()
		if err := os.Mkdir(filepath.Join(stage, "prior"), 0700); err != nil {
			t.Fatal(err)
		}
		if err := freezeModuleCache(stage, state, false, &SessionResponse{}); err != nil {
			t.Fatal(err)
		}
		data, err := os.ReadFile(filepath.Join(stage, "prior/index.json"))
		var candidates []ModuleMemo
		if err != nil || json.Unmarshal(data, &candidates) != nil || len(candidates) != 1 {
			t.Fatalf("profile lost its module candidates: %s / %v", data, err)
		}
		for kind, want := range sources {
			got, err := os.ReadFile(filepath.Join(stage, "prior/Main", kind+".go"))
			if err != nil || string(got) != want {
				t.Fatalf("profile lost reusable %s source: %v", kind, err)
			}
		}
	}
}
