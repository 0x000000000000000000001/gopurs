package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

type store struct{ root, cache string }
type pointer struct {
	Format   string `json:"format"`
	Schema   int    `json:"schema"`
	Manifest Object `json:"manifest"`
}

// The ownership journal is outside the disposable object cache. Before any
// public write it records both old and intended bytes, so an interrupted first
// publication cannot leave an unowned generated file behind.
type ownership struct {
	Format  string        `json:"format"`
	Schema  int           `json:"schema"`
	Output  string        `json:"output"`
	Outputs []OwnedOutput `json:"outputs"`
}
type ownershipEnvelope struct {
	SHA256  string          `json:"sha256"`
	Payload json.RawMessage `json:"payload"`
}

func digest(data []byte) string { sum := sha256.Sum256(data); return hex.EncodeToString(sum[:]) }

func absolute(root, name string) string {
	if filepath.IsAbs(name) {
		return filepath.Clean(name)
	}
	return filepath.Join(root, name)
}

func strictJSON(data []byte, value any) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(value); err != nil {
		return err
	}
	if err := d.Decode(new(any)); err != io.EOF {
		return fmt.Errorf("trailing JSON")
	}
	return nil
}

// Exchange is the only host boundary: UTF-8 JSON in/out. The Go compiler links
// it directly; the JS compiler invokes the same code as a small helper process.
func Exchange(request string) string {
	result := Response{Schema: Schema, Status: "error"}
	var req Request
	if err := strictJSON([]byte(request), &req); err != nil {
		result.Reason = err.Error()
	} else {
		response, err := execute(req)
		if err != nil {
			result.Reason = err.Error()
		} else {
			result = response
		}
	}
	data, _ := json.Marshal(result)
	return string(data)
}

func execute(req Request) (Response, error) {
	if req.Schema != Schema {
		return Response{}, fmt.Errorf("request-version")
	}
	if req.Operation != "snapshot" && req.Operation != "lookup" && req.Operation != "publish" {
		return Response{}, fmt.Errorf("unknown operation")
	}
	if req.Spec.Workspace == "" || req.Spec.Output == "" {
		return Response{}, fmt.Errorf("workspace and output are required")
	}
	workspace, err := filepath.Abs(req.Spec.Workspace)
	if err != nil {
		return Response{}, err
	}
	workspace, err = filepath.EvalSymlinks(workspace)
	if err != nil {
		return Response{}, err
	}
	output, err := filepath.EvalSymlinks(absolute(workspace, req.Spec.Output))
	if err != nil {
		return Response{}, err
	}
	stat, err := os.Stat(output)
	if err != nil || !stat.IsDir() {
		return Response{}, fmt.Errorf("output is not a directory: %s", output)
	}
	s := store{root: output, cache: filepath.Join(output, ".gopurs-cache")}
	for _, dir := range []string{s.cache, filepath.Join(s.cache, "objects")} {
		if err := privateDirectory(dir); err != nil {
			return Response{}, err
		}
	}
	lockPath := filepath.Join(s.cache, "write.lock")
	if info, err := os.Lstat(lockPath); err == nil && !info.Mode().IsRegular() {
		return Response{}, fmt.Errorf("cache lock is not a regular file")
	}
	lock, err := os.OpenFile(lockPath, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return Response{}, err
	}
	defer lock.Close()
	// Never remove/replace the lock inode. The kernel releases it after crashes,
	// including SIGKILL, and coordinates JS/Go and containers sharing this file.
	if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_EX); err != nil {
		return Response{}, err
	}
	defer syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)
	req.Spec.Workspace = workspace
	req.Spec.Output = output
	snap, err := snapshot(req.Spec)
	if err != nil {
		return Response{}, err
	}
	response := Response{Schema: Schema, Status: "snapshot", Snapshot: &snap}
	if req.Operation == "snapshot" {
		return response, nil
	}
	old, reason := s.readManifest()
	owned, err := s.readOwnership()
	if err != nil {
		return Response{}, err
	}
	if req.Operation == "lookup" {
		response.Status = "miss"
		response.Reason = reason
		if old == nil {
			return response, nil
		}
		if !sameOwnership(owned, old.Outputs) {
			response.Reason = "publication-incomplete"
			return response, nil
		}
		if old.Snapshot.Workspace != workspace || old.Snapshot.Output != output || old.Snapshot.Key != snap.Key {
			response.Reason = "inputs-changed"
			return response, nil
		}
		for _, out := range old.Outputs {
			if _, err := s.readObject(out.Object); err != nil {
				response.Reason = "object-corrupt"
				return response, nil
			}
			path, err := s.destination(out.Path)
			if err != nil || !fileMatches(path, out.Object) {
				response.Reason = "output-changed"
				return response, nil
			}
		}
		for _, part := range artifacts(*old) {
			data, err := s.readObject(part.Object)
			if err != nil || validateWire(data) != nil {
				response.Reason = "artifact-corrupt"
				return response, nil
			}
		}
		if info, err := os.Lstat(filepath.Join(output, "go.mod")); err != nil || !info.Mode().IsRegular() {
			response.Reason = "go-mod-missing"
			return response, nil
		}
		check, err := snapshot(req.Spec)
		if err != nil {
			return Response{}, err
		}
		if check.Key != snap.Key {
			response.Reason = "inputs-changed"
			return response, nil
		}
		response.Status = "hit"
		response.Reason = ""
		response.Manifest = old
		return response, nil
	}
	if req.ExpectedKey != snap.Key {
		response.Status = "miss"
		response.Reason = "inputs-changed"
		return response, nil
	}
	manifest, err := s.stage(req, snap)
	if err != nil {
		return Response{}, err
	}
	// A valid committed manifest is also ownership evidence. In particular,
	// recover its obsolete outputs if the journal was removed or truncated to an
	// older valid inventory. The journal still supplies any uncommitted outputs.
	if old != nil && old.Snapshot.Output == output {
		owned = mergeOwnership(owned, old.Outputs)
	}
	// Never infer ownership by scanning *.go. An externally edited obsolete file
	// is a conflict; new intended bytes may replace a damaged current output.
	current := map[string]bool{}
	for _, out := range manifest.Outputs {
		current[out.Path] = true
		if _, err := s.destination(out.Path); err != nil {
			return Response{}, err
		}
	}
	obsolete := []string{}
	byPath := map[string][]Object{}
	for _, out := range owned {
		byPath[out.Path] = append(byPath[out.Path], out.Object)
	}
	for relative, versions := range byPath {
		if current[relative] {
			continue
		}
		path, err := s.destination(relative)
		if err != nil {
			return Response{}, err
		}
		if _, err := os.Lstat(path); os.IsNotExist(err) {
			continue
		}
		matches := false
		for _, object := range versions {
			if fileMatches(path, object) {
				matches = true
				break
			}
		}
		if !matches {
			return Response{}, fmt.Errorf("obsolete output modified externally: %s", relative)
		}
		obsolete = append(obsolete, path)
	}
	sort.Strings(obsolete)
	// Staging may be expensive; reject a source edit before touching public files.
	check, err := snapshot(req.Spec)
	if err != nil {
		return Response{}, err
	}
	if check.Key != snap.Key {
		response.Status = "miss"
		response.Reason = "inputs-changed"
		return response, nil
	}
	if err := s.writeOwnership(mergeOwnership(owned, manifest.Outputs)); err != nil {
		return Response{}, err
	}
	for _, out := range manifest.Outputs {
		path, _ := s.destination(out.Path)
		if fileMatches(path, out.Object) {
			continue
		}
		data, err := s.readObject(out.Object)
		if err != nil {
			return Response{}, err
		}
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return Response{}, err
		}
		if err := atomicWrite(path, data, 0644); err != nil {
			return Response{}, err
		}
	}
	for _, path := range obsolete {
		if err := os.Remove(path); err != nil {
			return Response{}, err
		}
		if err := syncDirectory(filepath.Dir(path)); err != nil {
			return Response{}, err
		}
	}
	if err := s.ensureGoMod(); err != nil {
		return Response{}, err
	}
	check, err = snapshot(req.Spec)
	if err != nil {
		return Response{}, err
	}
	if check.Key != snap.Key {
		response.Status = "miss"
		response.Reason = "inputs-changed"
		return response, nil
	}
	// The pointer is the last semantic write. Interrupted publications leave the old
	// coherent manifest (whose output hashes may now miss), or the complete new one.
	data, err := json.Marshal(manifest)
	if err != nil {
		return Response{}, err
	}
	object, err := s.putObject(data)
	if err != nil {
		return Response{}, err
	}
	data, _ = json.Marshal(pointer{Format: Format, Schema: Schema, Manifest: object})
	if err := atomicWrite(filepath.Join(s.cache, "current.json"), data, 0600); err != nil {
		return Response{}, err
	}
	if err := s.writeOwnership(manifest.Outputs); err != nil {
		return Response{}, err
	}
	response.Status = "published"
	response.Manifest = &manifest
	return response, nil
}

func privateDirectory(path string) error {
	if err := os.Mkdir(path, 0700); err != nil && !os.IsExist(err) {
		return err
	}
	info, err := os.Lstat(path)
	if err != nil || !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
		return fmt.Errorf("cache directory is not a real directory: %s", path)
	}
	return nil
}

func snapshot(spec Spec) (Snapshot, error) {
	compiler, err := fingerprints(spec.Workspace, spec.Compiler)
	if err != nil {
		return Snapshot{}, err
	}
	parts := []string{}
	hasCompiler := false
	for _, f := range compiler {
		if f.Kind != "file" || f.State != "file" {
			return Snapshot{}, fmt.Errorf("compiler resource absent/not a file: %s", f.Path)
		}
		if f.Name == "compiler" {
			hasCompiler = true
		}
		parts = append(parts, f.Name, f.Digest)
	}
	if !hasCompiler {
		return Snapshot{}, fmt.Errorf("compiler artifact is required")
	}
	identity := Key("compiler", parts...)
	inputs, err := fingerprints(spec.Workspace, spec.Inputs)
	if err != nil {
		return Snapshot{}, err
	}
	options, err := sortedSettings(spec.Options)
	if err != nil {
		return Snapshot{}, err
	}
	parts = []string{spec.Workspace, spec.Output, identity, Key("options", options...)}
	hasTast := false
	for _, f := range inputs {
		if f.Kind == "tast" && f.Resolved == spec.Output {
			hasTast = true
		}
		parts = append(parts, f.Name, f.Kind, f.Path, f.Resolved, f.State, f.Digest)
	}
	if !hasTast {
		return Snapshot{}, fmt.Errorf("output TAST inventory is required")
	}
	return Snapshot{Workspace: spec.Workspace, Output: spec.Output, Compiler: identity, Key: Key("build", parts...), Inputs: inputs}, nil
}

func fingerprints(root string, inputs []Input) ([]Fingerprint, error) {
	items := append([]Input{}, inputs...)
	sort.Slice(items, func(i, j int) bool { return items[i].Name < items[j].Name })
	result := []Fingerprint{}
	for i, input := range items {
		if input.Name == "" || (i > 0 && items[i-1].Name == input.Name) {
			return nil, fmt.Errorf("duplicate/empty input name")
		}
		f, err := fingerprint(root, input)
		if err != nil {
			return nil, err
		}
		result = append(result, f)
	}
	return result, nil
}

func fingerprint(root string, input Input) (Fingerprint, error) {
	f := Fingerprint{Name: input.Name, Kind: input.Kind, Path: absolute(root, input.Path), State: "missing"}
	if input.Kind != "file" && input.Kind != "directory" && input.Kind != "tast" {
		return f, fmt.Errorf("unknown input kind")
	}
	info, err := os.Stat(f.Path)
	if os.IsNotExist(err) {
		return f, nil
	}
	if err != nil {
		return f, err
	}
	f.Resolved, err = filepath.EvalSymlinks(f.Path)
	if err != nil {
		return f, err
	}
	if input.Kind == "file" {
		if !info.Mode().IsRegular() {
			return f, fmt.Errorf("input is not a regular file: %s", f.Path)
		}
		data, err := os.ReadFile(f.Path)
		if err != nil {
			return f, err
		}
		f.State = "file"
		f.Digest = digest(data)
		return f, nil
	}
	if !info.IsDir() {
		return f, fmt.Errorf("input is not a directory: %s", f.Path)
	}
	entries, err := os.ReadDir(f.Path)
	if err != nil {
		return f, err
	}
	parts := []string{}
	for _, entry := range entries {
		path := filepath.Join(f.Path, entry.Name())
		if input.Kind == "tast" {
			info, err := os.Stat(path)
			if os.IsNotExist(err) {
				continue
			}
			if err != nil {
				return f, err
			}
			if !info.IsDir() {
				continue
			}
			child, err := fingerprint(root, Input{Name: entry.Name(), Kind: "file", Path: filepath.Join(path, "corefn.json")})
			if err != nil {
				return f, err
			}
			if child.State == "missing" {
				continue
			}
			parts = append(parts, child.Name, child.Resolved, child.Digest)
		} else {
			info, err := os.Lstat(path)
			if err != nil {
				return f, err
			}
			link := ""
			if info.Mode()&os.ModeSymlink != 0 {
				link, err = os.Readlink(path)
				if err != nil {
					return f, err
				}
			}
			parts = append(parts, entry.Name(), strconv.FormatUint(uint64(info.Mode().Type()), 10), link)
		}
	}
	f.State = "directory"
	f.Digest = Key(input.Kind, parts...)
	return f, nil
}

func (s store) destination(relative string) (string, error) {
	path := s.root
	for _, part := range strings.Split(relative, "/") {
		if part == "" || part == "." || part == ".." {
			return "", fmt.Errorf("invalid output path")
		}
		path = filepath.Join(path, part)
		info, err := os.Lstat(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return "", err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return "", fmt.Errorf("output symlink: %s", path)
		}
		if path != filepath.Join(s.root, filepath.FromSlash(relative)) && !info.IsDir() {
			return "", fmt.Errorf("output parent is not a directory: %s", path)
		}
		if path == filepath.Join(s.root, filepath.FromSlash(relative)) && !info.Mode().IsRegular() {
			return "", fmt.Errorf("output is not a regular file: %s", path)
		}
	}
	return path, nil
}

func fileMatches(path string, obj Object) bool {
	info, err := os.Lstat(path)
	if err != nil || !info.Mode().IsRegular() || info.Size() != obj.Bytes {
		return false
	}
	data, err := os.ReadFile(path)
	return err == nil && digest(data) == obj.SHA256
}

func (s store) putObject(data []byte) (Object, error) {
	obj := Object{SHA256: digest(data), Bytes: int64(len(data))}
	path := filepath.Join(s.cache, "objects", obj.SHA256)
	if fileMatches(path, obj) {
		return obj, nil
	}
	return obj, atomicWrite(path, data, 0600)
}

func (s store) readObject(obj Object) ([]byte, error) {
	if !validObject(obj) {
		return nil, fmt.Errorf("invalid object reference")
	}
	path := filepath.Join(s.cache, "objects", obj.SHA256)
	info, err := os.Lstat(path)
	if err != nil || !info.Mode().IsRegular() || info.Size() != obj.Bytes {
		return nil, fmt.Errorf("missing/corrupt object: %s", obj.SHA256)
	}
	data, err := os.ReadFile(path)
	if err != nil || int64(len(data)) != obj.Bytes || digest(data) != obj.SHA256 {
		return nil, fmt.Errorf("missing/corrupt object: %s", obj.SHA256)
	}
	return data, nil
}

func mergeOwnership(left, right []OwnedOutput) []OwnedOutput {
	items := map[string]OwnedOutput{}
	for _, list := range [][]OwnedOutput{left, right} {
		for _, out := range list {
			items[out.Path+"\x00"+out.Object.SHA256] = out
		}
	}
	keys := make([]string, 0, len(items))
	for key := range items {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	result := make([]OwnedOutput, 0, len(keys))
	for _, key := range keys {
		result = append(result, items[key])
	}
	return result
}

func sameOwnership(left, right []OwnedOutput) bool {
	l, r := mergeOwnership(nil, left), mergeOwnership(nil, right)
	if len(l) != len(r) {
		return false
	}
	for i := range l {
		if l[i] != r[i] {
			return false
		}
	}
	return true
}

func (s store) readOwnership() ([]OwnedOutput, error) {
	path := filepath.Join(s.root, ".gopurs-ownership.json")
	info, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil || !info.Mode().IsRegular() {
		return nil, fmt.Errorf("ownership journal is not a regular file")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var envelope ownershipEnvelope
	var journal ownership
	if strictJSON(data, &envelope) != nil || digest(envelope.Payload) != envelope.SHA256 || strictJSON(envelope.Payload, &journal) != nil || journal.Format != Format || journal.Schema != Schema || journal.Output != s.root {
		return nil, fmt.Errorf("ownership journal is incompatible/corrupt")
	}
	for _, out := range journal.Outputs {
		if !outputPath(out.Kind, out.Module, out.Path) || !validObject(out.Object) {
			return nil, fmt.Errorf("ownership journal has an invalid output")
		}
	}
	return journal.Outputs, nil
}

func (s store) writeOwnership(outputs []OwnedOutput) error {
	payload, err := json.Marshal(ownership{Format: Format, Schema: Schema, Output: s.root, Outputs: mergeOwnership(nil, outputs)})
	if err != nil {
		return err
	}
	data, err := json.Marshal(ownershipEnvelope{SHA256: digest(payload), Payload: payload})
	if err != nil {
		return err
	}
	path := filepath.Join(s.root, ".gopurs-ownership.json")
	if previous, err := os.ReadFile(path); err == nil && bytes.Equal(previous, data) {
		return nil
	}
	return atomicWrite(path, data, 0600)
}

func (s store) readManifest() (*Manifest, string) {
	data, err := os.ReadFile(filepath.Join(s.cache, "current.json"))
	if os.IsNotExist(err) {
		return nil, "cache-absent"
	}
	if err != nil {
		return nil, "manifest-unreadable"
	}
	var p pointer
	if strictJSON(data, &p) != nil || p.Format != Format || p.Schema != Schema {
		return nil, "manifest-version-or-corrupt"
	}
	data, err = s.readObject(p.Manifest)
	if err != nil {
		return nil, "manifest-object-corrupt"
	}
	var m Manifest
	if strictJSON(data, &m) != nil || validateManifest(m) != nil {
		return nil, "manifest-invalid"
	}
	return &m, ""
}

func (s store) stage(req Request, snap Snapshot) (Manifest, error) {
	m := Manifest{Format: Format, Schema: Schema, Snapshot: snap, Outputs: []OwnedOutput{}, Modules: []ModuleMemo{}, GoMod: "create-if-absent", GoSum: "external"}
	for _, out := range req.Outputs {
		if !outputPath(out.Kind, out.Module, out.Path) {
			return m, fmt.Errorf("unowned output path: %s", out.Path)
		}
		data, err := os.ReadFile(absolute(req.Spec.Workspace, out.Source))
		if err != nil {
			return m, err
		}
		obj, err := s.putObject(data)
		if err != nil {
			return m, err
		}
		m.Outputs = append(m.Outputs, OwnedOutput{Path: out.Path, Kind: out.Kind, Module: out.Module, Object: obj})
	}
	sort.Slice(m.Outputs, func(i, j int) bool { return m.Outputs[i].Path < m.Outputs[j].Path })
	for _, mod := range req.Modules {
		memo := ModuleMemo{Name: mod.Name, Prepared: mod.Prepared, OptimizerEnvironment: mod.OptimizerEnvironment, EmitterEnvironment: mod.EmitterEnvironment, CallerDemands: mod.CallerDemands, Parts: []Artifact{}}
		for _, part := range mod.Parts {
			a, err := s.stageArtifact(req.Spec.Workspace, part)
			if err != nil {
				return m, err
			}
			memo.Parts = append(memo.Parts, a)
		}
		m.Modules = append(m.Modules, memo)
	}
	if req.Diagnostics != nil {
		a, err := s.stageArtifact(req.Spec.Workspace, *req.Diagnostics)
		if err != nil {
			return m, err
		}
		m.Diagnostics = &a
	}
	return m, validateManifest(m)
}

func (s store) stageArtifact(root string, part StagedArtifact) (Artifact, error) {
	data, err := os.ReadFile(absolute(root, part.Source))
	if err != nil {
		return Artifact{}, err
	}
	if err := validateWire(data); err != nil {
		return Artifact{}, err
	}
	obj, err := s.putObject(data)
	return Artifact{Codec: part.Codec, Object: obj}, err
}

func artifacts(m Manifest) []Artifact {
	parts := []Artifact{}
	for _, mod := range m.Modules {
		parts = append(parts, mod.Parts...)
	}
	if m.Diagnostics != nil {
		parts = append(parts, *m.Diagnostics)
	}
	return parts
}

func atomicWrite(path string, data []byte, mode os.FileMode) error {
	f, err := os.CreateTemp(filepath.Dir(path), ".gopurs-txn-")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	if _, err := f.Write(data); err != nil {
		f.Close()
		return err
	}
	if err := f.Chmod(mode); err != nil {
		f.Close()
		return err
	}
	if err := f.Sync(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	if err := os.Rename(f.Name(), path); err != nil {
		return err
	}
	return syncDirectory(filepath.Dir(path))
}

func syncDirectory(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return f.Sync()
}

func (s store) ensureGoMod() error {
	path := filepath.Join(s.root, "go.mod")
	if info, err := os.Lstat(path); err == nil {
		if !info.Mode().IsRegular() {
			return fmt.Errorf("go.mod is not a regular file")
		}
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}
	// Exclusive link publishes a complete initial file without replacing a file
	// concurrently created by Go. Dependencies, replace/toolchain and go.sum are
	// not compiler-owned and are never overwritten or garbage-collected.
	f, err := os.CreateTemp(s.root, ".gopurs-mod-")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	if _, err := f.WriteString("module gopurs/output\n\ngo 1.22\n"); err != nil {
		f.Close()
		return err
	}
	if err := f.Chmod(0644); err != nil {
		f.Close()
		return err
	}
	if err := f.Sync(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	if err := os.Link(f.Name(), path); err != nil && !os.IsExist(err) {
		return err
	}
	return syncDirectory(s.root)
}
