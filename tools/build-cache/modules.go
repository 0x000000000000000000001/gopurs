package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
)

const moduleSourceCodec = "gopurs/go-source-gzip-v1"

func (s store) stageModuleSource(root string, source StagedModuleSource, previous []ModuleSource, packer *packWriter) (ModuleSource, error) {
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
	object, err := packer.put(compressed.Bytes())
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
	if err := s.ensurePacks(); err != nil {
		return err
	}
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
		if lazyModules(state.Spec) {
			return freezeLazyModuleCache(stage, state, manifest, s)
		}
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
			if valid && memo.Reads != nil {
				data, err := s.readObject(memo.Reads.Object)
				if err != nil {
					valid = false
				} else {
					files["reads.json"] = data
				}
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

// ReusedModule carries only producer keys. Artifact identities always come from
// the private committed manifest captured at begin, never from this request.
type ReusedModule struct {
	Name                 string `json:"name"`
	Prepared             string `json:"prepared"`
	OptimizerEnvironment string `json:"optimizerEnvironment"`
	CallerDemands        string `json:"callerDemands"`
	EmitterEnvironment   string `json:"emitterEnvironment"`
	Metadata             string `json:"metadata"`
	FFIDigest            string `json:"ffiDigest"`
}

func (m ReusedModule) staged() StagedModule {
	return StagedModule{Inherited: true, Name: m.Name, Prepared: m.Prepared, OptimizerEnvironment: m.OptimizerEnvironment, CallerDemands: m.CallerDemands, EmitterEnvironment: m.EmitterEnvironment, Metadata: m.Metadata, FFIDigest: m.FFIDigest}
}
func lazyModules(spec Spec) bool {
	for _, option := range spec.Options {
		if option.Name == "module-lazy" && option.Value == "1" {
			return true
		}
	}
	return false
}

type lazyPrior struct {
	state        buildSession
	mu           sync.Mutex
	manifest     *Manifest
	modules      map[string]ModuleMemo
	outputs      map[string]map[string]Object
	store        store
	materialized map[string]bool
}

// Native hosts retain one parsed index per build; helper-process hosts read the
// same private manifest. Every materialized file is bound to its expected hash.
var lazyPriors = struct {
	sync.Mutex
	values map[string]*lazyPrior
}{values: map[string]*lazyPrior{}}

func newLazyPrior(manifest *Manifest, s store) *lazyPrior {
	p := &lazyPrior{manifest: manifest, store: s, modules: map[string]ModuleMemo{}, outputs: map[string]map[string]Object{}, materialized: map[string]bool{}}
	for _, memo := range manifest.Modules {
		p.modules[memo.Name] = memo
	}
	for _, output := range manifest.Outputs {
		if output.Kind == "module" || output.Kind == "ffi" {
			if p.outputs[output.Module] == nil {
				p.outputs[output.Module] = map[string]Object{}
			}
			p.outputs[output.Module][output.Kind] = output.Object
		}
	}
	return p
}
func freezeLazyModuleCache(stage string, state buildSession, manifest *Manifest, s store) error {
	pinned, err := pinLazyArtifacts(stage, manifest, s)
	if err != nil {
		return err
	}
	p := newLazyPrior(manifest, pinned)
	p.state = state
	data, err := json.Marshal(manifest)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(stage, "prior", "manifest.json"), data, 0400); err != nil {
		return err
	}
	index := make([]ModuleMemo, len(manifest.Modules))
	valid := make([]bool, len(index))
	if err := boundedEach(len(index), func(i int) error {
		memo := manifest.Modules[i]
		if memo.Reads != nil {
			data, err := p.store.validateObject(memo.Reads.Object)
			if err != nil {
				return nil
			} // absent proof forces ordinary optimization
			dir := filepath.Join(stage, "prior", memo.Name)
			if err := os.MkdirAll(dir, 0700); err != nil {
				return err
			}
			if err := os.WriteFile(filepath.Join(dir, "reads.json"), data, 0400); err != nil {
				return err
			}
		}
		index[i], valid[i] = memo, true
		return nil
	}); err != nil {
		return err
	}
	retained := make([]ModuleMemo, 0, len(index))
	for i, memo := range index {
		if valid[i] {
			retained = append(retained, memo)
		}
	}
	data, err = json.Marshal(retained)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(stage, "prior", "index.json"), data, 0400); err != nil {
		return err
	}
	lazyPriors.Lock()
	lazyPriors.values[stage] = p
	lazyPriors.Unlock()
	return nil
}
func getLazyPrior(stage string, state buildSession) (*lazyPrior, error) {
	lazyPriors.Lock()
	defer lazyPriors.Unlock()
	if p := lazyPriors.values[stage]; p != nil {
		return p, nil
	}
	data, err := os.ReadFile(filepath.Join(stage, "prior", "manifest.json"))
	if err != nil {
		return nil, err
	}
	var manifest Manifest
	if err := strictJSON(data, &manifest); err != nil {
		return nil, err
	}
	if err := validateManifest(manifest); err != nil {
		return nil, err
	}
	if manifest.Snapshot.Compiler != state.Snapshot.Compiler || manifest.Snapshot.Workspace != state.Spec.Workspace || manifest.Snapshot.Output != state.Spec.Output || manifest.Diagnostics != nil {
		return nil, fmt.Errorf("incompatible private module snapshot")
	}
	s := store{root: state.Spec.Output, cache: filepath.Join(stage, "prior", "objects-cache")}
	if err := s.ensurePacks(); err != nil {
		return nil, err
	}
	p := newLazyPrior(&manifest, s)
	p.state = state
	lazyPriors.values[stage] = p
	return p, nil
}
func forgetLazyPrior(stage string) {
	lazyPriors.Lock()
	p := lazyPriors.values[stage]
	delete(lazyPriors.values, stage)
	lazyPriors.Unlock()
	if p == nil {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.store.packs != nil {
		p.store.packs.mu.Lock()
		defer p.store.packs.mu.Unlock()
		for _, file := range p.store.packs.files {
			_ = file.Close()
		}
	}
}
func materializePrior(stage string, state buildSession, owners, parts []string) error {
	p, err := getLazyPrior(stage, state)
	if err != nil {
		return err
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, owner := range owners {
		memo, ok := p.modules[owner]
		if !ok {
			return fmt.Errorf("unknown prior module: %s", owner)
		}
		for _, part := range parts {
			key := owner + "/" + part
			if p.materialized[key] {
				continue
			}
			var object Object
			var source *ModuleSource
			ext := ".json"
			switch part {
			case "0", "1", "2", "3":
				i := int(part[0] - '0')
				if i >= len(memo.Parts) {
					return fmt.Errorf("missing prior part: %s", key)
				}
				object = memo.Parts[i].Object
			case "reads":
				if memo.Reads == nil {
					continue
				}
				object = memo.Reads.Object
			case "module", "ffi":
				ext = ".go"
				for i := range memo.Sources {
					if memo.Sources[i].Kind == part {
						source = &memo.Sources[i]
						object = source.Raw
						break
					}
				}
				if source == nil {
					object = p.outputs[owner][part]
				}
				if object.SHA256 == "" {
					if part == "ffi" {
						continue
					}
					return fmt.Errorf("missing prior source: %s", key)
				}
			default:
				return fmt.Errorf("invalid prior part: %s", part)
			}
			destination := filepath.Join(stage, "prior", owner, part+ext)
			if !fileMatches(destination, object) {
				var data []byte
				if source == nil {
					data, err = p.store.validateObject(object)
				} else {
					data, err = p.store.readModuleSource(*source)
				}
				if err != nil {
					return fmt.Errorf("materializing %s: %w", key, err)
				}
				if err := os.MkdirAll(filepath.Dir(destination), 0700); err != nil {
					return err
				}
				if err := writePrivateFile(destination, data); err != nil {
					return err
				}
			}
			p.materialized[key] = true
		}
	}
	return nil
}

// The current planner consumes unqualified source files. Restore only missing
// files, so fresh producer output always wins. No PBO part is read here.
func restorePriorSources(stage string, state buildSession) error {
	if !lazyModules(state.Spec) {
		return nil
	}
	missing := []string{}
	for _, source := range state.Modules {
		name := strings.ReplaceAll(source.Name, ".", "_")
		if _, err := os.Stat(filepath.Join(stage, "output", "purescript", name+".go")); os.IsNotExist(err) {
			missing = append(missing, source.Name)
		} else if err != nil {
			return err
		}
	}
	if len(missing) == 0 {
		return nil
	}
	if err := materializePrior(stage, state, missing, []string{"module", "ffi"}); err != nil {
		return err
	}
	for _, owner := range missing {
		for _, kind := range []string{"module", "ffi"} {
			from := filepath.Join(stage, "prior", owner, kind+".go")
			name := strings.ReplaceAll(owner, ".", "_")
			if kind == "ffi" {
				name += "_ffi"
			}
			to := filepath.Join(stage, "output", "purescript", name+".go")
			if _, err := os.Stat(from); os.IsNotExist(err) && kind == "ffi" {
				continue
			}
			if err := os.MkdirAll(filepath.Dir(to), 0700); err != nil {
				return err
			}
			if err := os.Link(from, to); err != nil {
				data, err := os.ReadFile(from)
				if err != nil {
					return err
				}
				if err := os.WriteFile(to, data, 0600); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Unlike publication, private staging needs atomic visibility but no durability
// barrier. Concurrent helper processes may materialize the same immutable part.
func writePrivateFile(path string, data []byte) error {
	f, err := os.CreateTemp(filepath.Dir(path), ".materialize-")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	if _, err := f.Write(data); err != nil {
		f.Close()
		return err
	}
	if err := f.Chmod(0400); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(f.Name(), path)
}

// Pin immutable containers rather than unpacking every module. Cache GC may
// compact/remove public packs while this build is active; links keep exactly
// this generation available to both native and independent helper processes.
func pinLazyArtifacts(stage string, manifest *Manifest, source store) (store, error) {
	pinned := store{root: source.root, cache: filepath.Join(stage, "prior", "objects-cache")}
	for _, dir := range []string{"objects", "packs"} {
		if err := os.MkdirAll(filepath.Join(pinned.cache, dir), 0700); err != nil {
			return pinned, err
		}
	}
	index := packIndexFile{Format: packIndexFormat, Schema: packIndexSchema, Objects: map[string]packEntry{}}
	links := map[string]bool{}
	pin := func(relative string) (bool, error) {
		if exists, ok := links[relative]; ok {
			return exists, nil
		}
		from, to := filepath.Join(source.cache, relative), filepath.Join(pinned.cache, relative)
		if err := os.Link(from, to); err != nil {
			if os.IsNotExist(err) {
				links[relative] = false
				return false, nil
			}
			data, readErr := os.ReadFile(from)
			if readErr != nil {
				return false, readErr
			}
			if err := os.WriteFile(to, data, 0400); err != nil {
				return false, err
			}
		}
		links[relative] = true
		return true, nil
	}
	objects := []Object{}
	for _, part := range artifacts(*manifest) {
		objects = append(objects, part.Object)
	}
	for _, output := range manifest.Outputs {
		if output.Kind == "module" || output.Kind == "ffi" {
			objects = append(objects, output.Object)
		}
	}
	for _, object := range objects {
		if entry, ok := source.packs.entries[object.SHA256]; ok {
			exists, err := pin(filepath.Join("packs", entry.Pack+".pack"))
			if err != nil {
				return pinned, err
			}
			if exists {
				index.Objects[object.SHA256] = entry
			}
		} else {
			if _, err := pin(filepath.Join("objects", object.SHA256)); err != nil {
				return pinned, err
			}
		}
	}
	data, err := json.Marshal(index)
	if err != nil {
		return pinned, err
	}
	if err := os.WriteFile(filepath.Join(pinned.cache, "packs", "index.json"), data, 0400); err != nil {
		return pinned, err
	}
	if err := pinned.ensurePacks(); err != nil {
		return pinned, err
	}
	return pinned, nil
}
