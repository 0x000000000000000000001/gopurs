package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

// Module candidates are an immutable private copy of one committed manifest.
// They are hints until the typed producer checks the prepared/optimizer/emitter
// keys and decodes the business schemas. No public output is used as authority.
func freezeModuleCache(stage string, state buildSession, force bool, result *SessionResponse) error {
	index := []ModuleMemo{}
	s := store{root: state.Spec.Output, cache: filepath.Join(state.Spec.Output, ".gopurs-cache")}
	lock, err := os.OpenFile(filepath.Join(s.cache, "write.lock"), os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer lock.Close()
	if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_EX); err != nil {
		return err
	}
	defer syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)
	manifest, _ := s.readManifest()
	if !force && manifest != nil && manifest.Diagnostics == nil &&
		manifest.Snapshot.Compiler == state.Snapshot.Compiler &&
		manifest.Snapshot.Workspace == state.Spec.Workspace && manifest.Snapshot.Output == state.Spec.Output {
		// The preparation producer is deterministic for a complete TAST inventory
		// and compiler. The typed driver must additionally match its options and
		// initial directives before reusing any prepared fingerprint from this index.
		result.SameTAST = sameTASTInputs(manifest.Snapshot, state.Snapshot)
		outputs := map[string][]OwnedOutput{}
		for _, output := range manifest.Outputs {
			if output.Kind == "module" || output.Kind == "ffi" {
				outputs[output.Module] = append(outputs[output.Module], output)
			}
		}
		for _, memo := range manifest.Modules {
			files := map[string][]byte{}
			valid := true
			for i, part := range memo.Parts {
				data, err := s.readObject(part.Object)
				// The committed producer validated these bytes; the typed driver
				// will additionally validate the business schema before reuse.
				if err != nil {
					valid = false
					break
				}
				files[fmt.Sprintf("%d.json", i)] = data
			}
			body := false
			for _, output := range outputs[memo.Name] {
				data, err := s.readObject(output.Object)
				if err != nil {
					valid = false
					break
				}
				files[output.Kind+".go"] = data
				body = body || output.Kind == "module"
			}
			if !valid || !body {
				continue
			}
			directory := filepath.Join(stage, "prior", memo.Name)
			if err := os.Mkdir(directory, 0700); err != nil {
				return err
			}
			for name, bytes := range files {
				if err := os.WriteFile(filepath.Join(directory, name), bytes, 0400); err != nil {
					return err
				}
			}
			index = append(index, memo)
		}
	}
	data, err := json.Marshal(index)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(stage, "prior", "index.json"), data, 0400)
}

func sameTASTInputs(a, b Snapshot) bool {
	find := func(snapshot Snapshot) (Fingerprint, bool) {
		var found Fingerprint
		count := 0
		for _, input := range snapshot.Inputs {
			if input.Kind == "tast" {
				found = input
				count++
			}
		}
		return found, count == 1 && found.Name == "tast" && found.State == "directory"
	}
	x, okA := find(a)
	y, okB := find(b)
	return okA && okB && x == y
}
