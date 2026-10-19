package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

type SourceModule struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	HasForeign bool   `json:"hasForeign"`
}

type FrozenFFI struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	File   string `json:"file"`
	Found  bool   `json:"found"`
	Digest string `json:"digest"`
}

type SessionRequest struct {
	Operation   string         `json:"operation"`
	Options     []Setting      `json:"options"`
	Force       bool           `json:"force"`
	Stage       string         `json:"stage"`
	Modules     []SourceModule `json:"modules"`
	Entries     []string       `json:"entries"`
	Diagnostics bool           `json:"diagnostics"`
	ModuleCache bool           `json:"moduleCache"`
}

type SessionResponse struct {
	Status         string      `json:"status"`
	Reason         string      `json:"reason"`
	Stage          string      `json:"stage"`
	Input          string      `json:"input"`
	Output         string      `json:"output"`
	Count          int         `json:"count"`
	Context        string      `json:"context"`
	SameTAST       bool        `json:"sameTast"`
	FFI            []FrozenFFI `json:"ffi"`
	PackageRequest string      `json:"packageRequest"`
}

type buildSession struct {
	Spec     Spec           `json:"spec"`
	Snapshot Snapshot       `json:"snapshot"`
	Count    int            `json:"count"`
	Modules  []SourceModule `json:"modules"`
	Planned  bool           `json:"planned"`
}

type recipe struct {
	Schema int     `json:"schema"`
	Inputs []Input `json:"inputs"`
}

// SessionExchange is the driver's producer protocol. Storage v1 remains usable
// independently. No PureScript runtime representation crosses this boundary.
func SessionExchange(request string) string {
	response := SessionResponse{Status: "error", FFI: []FrozenFFI{}}
	var req SessionRequest
	if err := strictJSON([]byte(request), &req); err != nil {
		response.Reason = err.Error()
	} else if err := sessionOperation(req, &response); err != nil {
		response.Status, response.Reason = "error", err.Error()
	}
	data, _ := json.Marshal(response)
	return string(data)
}

func sessionOperation(req SessionRequest, result *SessionResponse) error {
	workspace, err := os.Getwd()
	if err != nil {
		return err
	}
	workspace, err = filepath.EvalSymlinks(workspace)
	if err != nil {
		return err
	}
	output, err := filepath.EvalSymlinks(filepath.Join(workspace, "output"))
	if err != nil {
		return err
	}
	if req.Operation == "begin" {
		compiler, err := CompilerInputs()
		if err != nil {
			return err
		}
		spec := Spec{Workspace: workspace, Output: output, Compiler: compiler, Options: req.Options,
			Inputs: []Input{{Name: "tast", Kind: "tast", Path: output}}}
		if usesGoPackages(spec) {
			if err := (store{root: output}).ensureGoMod(); err != nil {
				return err
			}
			if err := packageToolchain(&spec); err != nil {
				return err
			}
		}
		return beginSession(spec, req.Force, result)
	}
	stage := filepath.Clean(req.Stage)
	if filepath.Dir(stage) != filepath.Join(output, ".gopurs-cache", "stages") || !strings.HasPrefix(filepath.Base(stage), "build-") {
		return fmt.Errorf("invalid build staging directory")
	}
	data, err := os.ReadFile(filepath.Join(stage, "session.json"))
	if err != nil {
		return err
	}
	var state buildSession
	if err := strictJSON(data, &state); err != nil {
		return err
	}
	if state.Spec.Workspace != workspace || state.Spec.Output != output {
		return fmt.Errorf("build workspace changed")
	}
	result.Stage, result.Input, result.Output, result.Count = stage, filepath.Join(stage, "input"), filepath.Join(stage, "output"), state.Count
	switch req.Operation {
	case "plan":
		return planSession(stage, &state, req.Modules, result)
	case "packages":
		return preparePackageSession(stage, &state, req.Entries, result)
	case "publish":
		return publishSession(stage, state, req, result)
	case "abort":
		result.Status = "closed"
		return os.RemoveAll(stage)
	default:
		return fmt.Errorf("unknown driver cache operation")
	}
}

func beginSession(spec Spec, force bool, result *SessionResponse) error {
	result.SameTAST = false
	baseInputs := spec.Inputs
	result.Reason = "recipe-absent"
	var saved recipe
	data, err := os.ReadFile(filepath.Join(spec.Output, ".gopurs-cache", "recipe.json"))
	if err == nil && strictJSON(data, &saved) == nil && saved.Schema == Schema {
		spec.Inputs = saved.Inputs
		lookup, lookupErr := execute(Request{Schema: Schema, Operation: "lookup", Spec: spec})
		if lookupErr == nil {
			if lookup.Status == "hit" && !force && lookup.Manifest.Diagnostics == nil {
				result.Status, result.Reason = "hit", ""
				return nil
			}
			result.Reason = lookup.Reason
			if lookup.Manifest != nil && lookup.Manifest.Diagnostics != nil {
				result.Reason = "diagnostics-require-build"
			}
		} else {
			result.Reason = "recipe-unusable"
		}
	} else if err == nil {
		result.Reason = "recipe-corrupt"
	}
	if force {
		result.Reason = "forced"
	}
	spec.Inputs = baseInputs
	initial, err := execute(Request{Schema: Schema, Operation: "snapshot", Spec: spec})
	if err != nil {
		return err
	}
	options, err := sortedSettings(spec.Options)
	if err != nil {
		return err
	}
	result.Context = Key("module-context", initial.Snapshot.Compiler, Key("options", options...))
	parent := filepath.Join(spec.Output, ".gopurs-cache", "stages")
	if err := privateDirectory(parent); err != nil {
		return err
	}
	stage, err := os.MkdirTemp(parent, "build-")
	if err != nil {
		return err
	}
	complete := false
	defer func() {
		if !complete {
			os.RemoveAll(stage)
		}
	}()
	input, output := filepath.Join(stage, "input"), filepath.Join(stage, "output")
	for _, dir := range []string{input, output, filepath.Join(stage, "ffi"), filepath.Join(stage, "modules"), filepath.Join(stage, "prior")} {
		if err := os.Mkdir(dir, 0700); err != nil {
			return err
		}
	}
	// Freeze the actual bytes later decoded by PBO, not just their before/after
	// hashes. An edit-and-revert during optimization cannot poison this build.
	entries, err := os.ReadDir(spec.Output)
	if err != nil {
		return err
	}
	parts := []string{}
	count := 0
	for _, entry := range entries {
		dir := filepath.Join(spec.Output, entry.Name())
		info, err := os.Stat(dir)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return err
		}
		if !info.IsDir() {
			continue
		}
		f, bytes, _, err := readObservation(spec.Workspace, Input{Name: entry.Name(), Kind: "file", Path: filepath.Join(dir, "corefn.json")})
		if err != nil {
			return err
		}
		if f.State == "missing" {
			continue
		}
		parts = append(parts, f.Name, f.Resolved, f.Digest)
		destination := filepath.Join(input, entry.Name())
		if err := os.Mkdir(destination, 0700); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(destination, "corefn.json"), bytes, 0400); err != nil {
			return err
		}
		count++
	}
	if Key("tast", parts...) != initial.Snapshot.Inputs[0].Digest {
		return fmt.Errorf("TAST changed while freezing inputs; retry the build")
	}
	check, err := snapshot(spec)
	if err != nil {
		return err
	}
	if check.Key != initial.Snapshot.Key {
		return fmt.Errorf("inputs changed while freezing TAST; retry the build")
	}
	state := buildSession{Spec: spec, Snapshot: check, Count: count}
	if err := freezeModuleCache(stage, state, force, result); err != nil {
		return err
	}
	if err := saveSession(stage, state); err != nil {
		return err
	}
	result.Status, result.Stage, result.Input, result.Output, result.Count = "miss", stage, input, output, count
	complete = true
	return nil
}

type observation struct {
	input       Input
	fingerprint Fingerprint
	bytes       []byte
	entries     []os.DirEntry
}
type ffiResolver struct {
	workspace string
	observed  map[string]observation
	roots     []string
}

func (r *ffiResolver) observe(kind, path string) (observation, error) {
	path = absolute(r.workspace, path)
	name := kind + "/" + path
	if saved, ok := r.observed[name]; ok {
		return saved, nil
	}
	input := Input{Name: name, Kind: kind, Path: path}
	f, bytes, entries, err := readObservation(r.workspace, input)
	if err != nil {
		return observation{}, err
	}
	value := observation{input, f, bytes, entries}
	r.observed[name] = value
	return value, nil
}

// This resolver is used by both compiler hosts. Its trace includes unsuccessful
// candidates, skipped paths and every directory used to choose search roots.
func (r *ffiResolver) searchRoots(ffiDir string) ([]string, error) {
	if r.roots != nil {
		return r.roots, nil
	}
	roots := []string{}
	for _, name := range []string{".spago", "spago.d"} {
		dir := filepath.Join(r.workspace, name)
		parent, err := r.observe("directory", dir)
		if err != nil {
			return nil, err
		}
		for _, entry := range parent.entries {
			pkgDir := filepath.Join(dir, entry.Name())
			pkg, err := r.observe("directory", pkgDir)
			if err != nil {
				return nil, err
			}
			if pkg.fingerprint.State != "directory" {
				continue
			}
			versions := []string{}
			for _, child := range pkg.entries {
				if !strings.HasPrefix(child.Name(), "v") {
					continue
				}
				path := filepath.Join(pkgDir, child.Name())
				version, err := r.observe("directory", path)
				if err != nil {
					return nil, err
				}
				if version.fingerprint.State == "directory" {
					versions = append(versions, path)
				}
			}
			if len(versions) == 0 {
				roots = append(roots, pkgDir)
			} else {
				roots = append(roots, versions...)
			}
		}
	}
	if ffiDir != "" {
		roots = append(roots, filepath.Join(r.workspace, ffiDir))
	}
	roots = append(roots, r.workspace)
	for _, dir := range roots {
		if _, err := r.observe("directory", dir); err != nil {
			return nil, err
		}
	}
	r.roots = roots
	return roots, nil
}

func (r *ffiResolver) resolve(module SourceModule, ffiDir string) (observation, bool, error) {
	if module.Path != "" {
		candidate := module.Path
		if strings.HasSuffix(candidate, ".purs") {
			candidate = strings.TrimSuffix(candidate, ".purs") + ".go"
		}
		found, err := r.observe("file", candidate)
		if err != nil {
			return observation{}, false, err
		}
		if found.fingerprint.State == "file" {
			return found, true, nil
		}
	}
	roots, err := r.searchRoots(ffiDir)
	if err != nil {
		return observation{}, false, err
	}
	for _, root := range roots {
		for _, candidate := range []string{
			filepath.Join(root, "src", filepath.FromSlash(strings.ReplaceAll(module.Name, ".", "/"))) + ".go",
			filepath.Join(root, "src", module.Name) + ".go", filepath.Join(root, module.Name) + ".go",
		} {
			found, err := r.observe("file", candidate)
			if err != nil {
				return observation{}, false, err
			}
			if found.fingerprint.State == "file" {
				return found, true, nil
			}
		}
	}
	return observation{}, false, nil
}

func planSession(stage string, state *buildSession, modules []SourceModule, result *SessionResponse) error {
	if state.Planned {
		return fmt.Errorf("build inputs already planned")
	}
	if len(modules) != state.Count {
		return fmt.Errorf("TAST decoding did not produce every inventoried module")
	}
	resolver := ffiResolver{workspace: state.Spec.Workspace, observed: map[string]observation{}}
	ffiDir := ""
	for _, option := range state.Spec.Options {
		if option.Name == "ffi" {
			ffiDir = option.Value
		}
	}
	seen := map[string]bool{}
	for i, module := range modules {
		if !validModule(module.Name) || seen[module.Name] {
			return fmt.Errorf("invalid/duplicate module: %s", module.Name)
		}
		seen[module.Name] = true
		if !module.HasForeign {
			continue
		}
		found, exists, err := resolver.resolve(module, ffiDir)
		if err != nil {
			return err
		}
		ffi := FrozenFFI{Name: module.Name, Found: exists, Digest: Key("ffi-absent", module.Name)}
		if exists {
			ffi.Digest = Key("ffi-module", found.input.Path, found.fingerprint.Resolved, found.fingerprint.Digest)
			ffi.Path, ffi.File = found.input.Path, filepath.Join(stage, "ffi", fmt.Sprintf("%d.go", i))
			if err := os.WriteFile(ffi.File, found.bytes, 0400); err != nil {
				return err
			}
		}
		result.FFI = append(result.FFI, ffi)
	}
	expected := append([]Fingerprint{}, state.Snapshot.Inputs...)
	for _, value := range resolver.observed {
		state.Spec.Inputs = append(state.Spec.Inputs, value.input)
		expected = append(expected, value.fingerprint)
	}
	sort.Slice(expected, func(i, j int) bool { return expected[i].Name < expected[j].Name })
	current, err := snapshot(state.Spec)
	if err != nil {
		return err
	}
	if current.Compiler != state.Snapshot.Compiler || !reflect.DeepEqual(current.Inputs, expected) {
		return fmt.Errorf("inputs changed while freezing FFI; retry the build")
	}
	state.Snapshot, state.Modules, state.Planned = current, modules, true
	if err := saveSession(stage, *state); err != nil {
		return err
	}
	result.Status = "planned"
	return nil
}

func saveSession(stage string, state buildSession) error {
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return atomicWrite(filepath.Join(stage, "session.json"), data, 0600)
}

func publishSession(stage string, state buildSession, req SessionRequest, result *SessionResponse) error {
	if !state.Planned {
		return fmt.Errorf("build inputs were not frozen")
	}
	outputs := []StagedOutput{}
	seen := map[string]bool{}
	add := func(kind, module, path string, optional bool) error {
		if seen[path] {
			return fmt.Errorf("generated path collision: %s", path)
		}
		seen[path] = true
		source := filepath.Join(stage, "output", filepath.FromSlash(path))
		if _, err := os.Stat(source); optional && os.IsNotExist(err) {
			return nil
		}
		outputs = append(outputs, StagedOutput{Path: path, Kind: kind, Module: module, Source: source})
		return nil
	}
	if err := add("runtime", "", "gopurs_runtime/runtime.go", false); err != nil {
		return err
	}
	if usesGoPackages(state.Spec) {
		files, err := packageOutputs(stage, state.Modules, req.Entries)
		if err != nil {
			return err
		}
		outputs = append(outputs, files...)
	} else {
		for _, module := range state.Modules {
			name := strings.ReplaceAll(module.Name, ".", "_")
			if err := add("module", module.Name, "purescript/"+name+".go", false); err != nil {
				return err
			}
			if module.HasForeign {
				if err := add("ffi", module.Name, "purescript/"+name+"_ffi.go", true); err != nil {
					return err
				}
			}
		}
	}
	for _, module := range req.Entries {
		if err := add("entry", module, module+"/main/main.go", false); err != nil {
			return err
		}
	}
	if len(req.Entries) > 0 {
		if err := add("entry", req.Entries[len(req.Entries)-1], "main/main.go", false); err != nil {
			return err
		}
	}
	publication := Request{Schema: Schema, Operation: "publish", Spec: state.Spec, ExpectedKey: state.Snapshot.Key, Outputs: outputs, checkLegacy: true}
	if req.ModuleCache {
		for _, source := range state.Modules {
			var memo ModuleMemo
			directory := filepath.Join(stage, "modules", source.Name)
			data, err := os.ReadFile(filepath.Join(directory, "memo.json"))
			if err != nil {
				return err
			}
			if err := strictJSON(data, &memo); err != nil {
				return err
			}
			if memo.Name != source.Name {
				return fmt.Errorf("module memo name mismatch")
			}
			mod := StagedModule{Name: memo.Name, Prepared: memo.Prepared, OptimizerEnvironment: memo.OptimizerEnvironment,
				EmitterEnvironment: memo.EmitterEnvironment, CallerDemands: memo.CallerDemands}
			for i, codec := range moduleCodecs {
				mod.Parts = append(mod.Parts, StagedArtifact{Codec: codec, Source: filepath.Join(directory, fmt.Sprintf("%d.json", i))})
			}
			if usesGoPackages(state.Spec) {
				name := strings.ReplaceAll(source.Name, ".", "_")
				mod.Sources = append(mod.Sources, StagedModuleSource{"module", filepath.Join(stage, "output", "purescript", name+".go")})
				if source.HasForeign {
					mod.Sources = append(mod.Sources, StagedModuleSource{"ffi", filepath.Join(stage, "output", "purescript", name+"_ffi.go")})
				}
			}
			publication.Modules = append(publication.Modules, mod)
		}
	}
	if req.Diagnostics {
		file := filepath.Join(stage, "diagnostics.json")
		if err := os.WriteFile(file, []byte(`["unit"]`), 0600); err != nil {
			return err
		}
		publication.Diagnostics = &StagedArtifact{Codec: "gopurs/diagnostics-v1", Source: file}
	}
	response, err := execute(publication)
	if err != nil {
		return err
	}
	if response.Status != "published" {
		return fmt.Errorf("%s during publication; retry the build", response.Reason)
	}
	// This is a lookup hint, not authority. A stale/corrupt recipe cannot match
	// the manifest's content key; concurrent writers can at worst cause a miss.
	data, _ := json.Marshal(recipe{Schema: Schema, Inputs: state.Spec.Inputs})
	if err := atomicWrite(filepath.Join(state.Spec.Output, ".gopurs-cache", "recipe.json"), data, 0600); err != nil {
		return err
	}
	if usesGoPackages(state.Spec) || packPublicationEnabled() {
		// Cache maintenance is best-effort after the semantic commit. A failed
		// removal leaves disposable objects; it cannot undo a successful build.
		s := store{root: state.Spec.Output, cache: filepath.Join(state.Spec.Output, ".gopurs-cache")}
		_ = s.retirePackageObjects()
	}
	result.Status = "published"
	return nil
}

// Initial adoption proves each legacy file against the fresh generation. Files
// without this evidence require a fresh generated-output namespace, not guesses
// about ownership based on a .go suffix. Text/assets outside it remain external.
func (s store) checkLegacyOutputs(manifest Manifest) error {
	intended := map[string]Object{}
	for _, out := range manifest.Outputs {
		intended[out.Path] = out.Object
	}
	entries, err := os.ReadDir(s.root)
	if err != nil {
		return err
	}
	dirs := []string{"purescript", "gopurs_runtime", "main"}
	for _, entry := range entries {
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") && entry.Name() != "purescript" && entry.Name() != "gopurs_runtime" && entry.Name() != "main" {
			dirs = append(dirs, filepath.Join(entry.Name(), "main"))
		}
	}
	for _, dir := range dirs {
		files, err := os.ReadDir(filepath.Join(s.root, dir))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return err
		}
		for _, file := range files {
			if !strings.HasSuffix(file.Name(), ".go") {
				continue
			}
			rel := filepath.ToSlash(filepath.Join(dir, file.Name()))
			obj, ok := intended[rel]
			if !ok || !fileMatches(filepath.Join(s.root, rel), obj) {
				return fmt.Errorf("legacy output has no verified ownership: %s; use an output directory containing only frontend artifacts for the first cached build", rel)
			}
		}
	}
	return nil
}
