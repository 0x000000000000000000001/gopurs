package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"
)

const moduleSourceCodec = "gopurs/go-source-gzip-v1"

func (s store) stageModuleSource(root string, source StagedModuleSource, previous []ModuleSource) (ModuleSource, error) {
	data, err := os.ReadFile(absolute(root, source.Source))
	if err != nil {
		return ModuleSource{}, err
	}
	raw := Object{SHA256: digest(data), Bytes: int64(len(data))}
	for _, prior := range previous {
		if prior.Kind == source.Kind && prior.Codec == moduleSourceCodec && prior.Raw == raw {
			if _, err := s.readModuleSource(prior); err == nil {
				return prior, nil
			}
		}
	}
	var compressed bytes.Buffer
	writer, err := gzip.NewWriterLevel(&compressed, gzip.BestSpeed)
	if err != nil {
		return ModuleSource{}, err
	}
	// No timestamp or filename: all hosts use this producer and byte framing.
	if _, err := writer.Write(data); err != nil {
		return ModuleSource{}, err
	}
	if err := writer.Close(); err != nil {
		return ModuleSource{}, err
	}
	object, err := s.putStagedObject(compressed.Bytes())
	return ModuleSource{Kind: source.Kind, Codec: moduleSourceCodec, Object: object, Raw: raw}, err
}

func (s store) readModuleSource(source ModuleSource) ([]byte, error) {
	data, err := s.readObject(source.Object)
	if err != nil {
		return nil, err
	}
	input := bytes.NewReader(data)
	reader, err := gzip.NewReader(input)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	reader.Multistream(false)
	raw, err := io.ReadAll(io.LimitReader(reader, source.Raw.Bytes+1))
	if err != nil || int64(len(raw)) != source.Raw.Bytes || digest(raw) != source.Raw.SHA256 || input.Len() != 0 {
		return nil, fmt.Errorf("invalid module source object: %s", source.Object.SHA256)
	}
	return raw, nil
}

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
			if len(memo.Sources) != 0 {
				for _, source := range memo.Sources {
					data, err := s.readModuleSource(source)
					if err != nil {
						valid = false
						break
					}
					files[source.Kind+".go"] = data
					body = body || source.Kind == "module"
				}
			} else {
				// Legacy and monolithic manifests already retain these exact bytes
				// as outputs; migration needs no duplicate compressed object.
				for _, output := range outputs[memo.Name] {
					data, err := s.readObject(output.Object)
					if err != nil {
						valid = false
						break
					}
					files[output.Kind+".go"] = data
					body = body || output.Kind == "module"
				}
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
