package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const Schema = 1
const Format = "gopurs-build-cache"

// Keys use byte-length framing, not a host's JSON serialization or ADT layout.
// Every component is UTF-8; sorted sets are ordered by UTF-8 bytes.
func Key(domain string, parts ...string) string {
	h := sha256.New()
	for _, part := range append([]string{Format, strconv.Itoa(Schema), domain}, parts...) {
		fmt.Fprintf(h, "%d:", len(part))
		h.Write([]byte(part))
	}
	return hex.EncodeToString(h.Sum(nil))
}

type Setting struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// file includes absence and symlink resolution. directory observes immediate
// entries (for the FFI resolver's search-root inventory). tast observes exactly
// the immediate <module>/corefn.json files, without decoding them.
type Input struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
	Path string `json:"path"`
}

type Spec struct {
	Workspace string    `json:"workspace"`
	Output    string    `json:"output"`
	Compiler  []Input   `json:"compiler"`
	Options   []Setting `json:"options"`
	Inputs    []Input   `json:"inputs"`
}

type Fingerprint struct {
	Name     string `json:"name"`
	Kind     string `json:"kind"`
	Path     string `json:"path"`
	Resolved string `json:"resolved"`
	State    string `json:"state"`
	Digest   string `json:"digest"`
}

type Snapshot struct {
	Workspace string        `json:"workspace"`
	Output    string        `json:"output"`
	Compiler  string        `json:"compiler"`
	Key       string        `json:"key"`
	Inputs    []Fingerprint `json:"inputs"`
}

type Object struct {
	SHA256 string `json:"sha256"`
	Bytes  int64  `json:"bytes"`
}

type OwnedOutput struct {
	Path   string `json:"path"`
	Kind   string `json:"kind"`
	Module string `json:"module"`
	Object Object `json:"object"`
}

type StagedOutput struct {
	Path   string `json:"path"`
	Kind   string `json:"kind"`
	Module string `json:"module"`
	Source string `json:"source"`
}

type Artifact struct {
	Codec  string `json:"codec"`
	Object Object `json:"object"`
}

type StagedArtifact struct {
	Codec  string `json:"codec"`
	Source string `json:"source"`
}

// A module memo is all-or-nothing. These are independent semantic environments,
// not merely an import/signature hash. Their producer must include absent reads,
// inlined bodies, ABI/layouts and transitive caller specialization demands.
type ModuleMemo struct {
	Name                 string     `json:"name"`
	Prepared             string     `json:"prepared"`
	OptimizerEnvironment string     `json:"optimizerEnvironment"`
	EmitterEnvironment   string     `json:"emitterEnvironment"`
	CallerDemands        string     `json:"callerDemands"`
	Parts                []Artifact `json:"parts"`
}

type StagedModule struct {
	Name                 string           `json:"name"`
	Prepared             string           `json:"prepared"`
	OptimizerEnvironment string           `json:"optimizerEnvironment"`
	EmitterEnvironment   string           `json:"emitterEnvironment"`
	CallerDemands        string           `json:"callerDemands"`
	Parts                []StagedArtifact `json:"parts"`
}

type Manifest struct {
	Format   string        `json:"format"`
	Schema   int           `json:"schema"`
	Snapshot Snapshot      `json:"snapshot"`
	Outputs  []OwnedOutput `json:"outputs"`
	// Module order is the producer's canonical dependency order, never map order.
	Modules     []ModuleMemo `json:"modules"`
	Diagnostics *Artifact    `json:"diagnostics"`
	GoMod       string       `json:"goMod"`
	GoSum       string       `json:"goSum"`
}

type Request struct {
	Schema      int             `json:"schema"`
	Operation   string          `json:"operation"`
	Spec        Spec            `json:"spec"`
	ExpectedKey string          `json:"expectedKey"`
	Outputs     []StagedOutput  `json:"outputs"`
	Modules     []StagedModule  `json:"modules"`
	Diagnostics *StagedArtifact `json:"diagnostics"`
}

type Response struct {
	Schema   int       `json:"schema"`
	Status   string    `json:"status"`
	Reason   string    `json:"reason,omitempty"`
	Snapshot *Snapshot `json:"snapshot,omitempty"`
	Manifest *Manifest `json:"manifest,omitempty"`
}

var digestPattern = regexp.MustCompile(`^[0-9a-f]{64}$`)
var hexUnitsPattern = regexp.MustCompile(`^([0-9a-f]{4})*$`)
var floatPattern = regexp.MustCompile(`^[0-9a-f]{16}$`)
var integerPattern = regexp.MustCompile(`^(0|-?[1-9][0-9]*)$`)
var tagPattern = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_./-]*$`)

var moduleCodecs = []string{"pbo/backend-v1", "pbo/implementations-v1", "pbo/directives-v1", "gopurs/workers-v1"}

func validModule(name string) bool {
	if name == "" || strings.ContainsAny(name, "/\\\x00") {
		return false
	}
	for _, part := range strings.Split(name, ".") {
		if part == "" {
			return false
		}
	}
	return true
}

func outputPath(kind, module, path string) bool {
	if filepath.IsAbs(path) || strings.ContainsAny(path, "\\\x00") || filepath.ToSlash(filepath.Clean(path)) != path {
		return false
	}
	goName := strings.ReplaceAll(module, ".", "_")
	switch kind {
	case "module":
		return validModule(module) && path == "purescript/"+goName+".go"
	case "ffi":
		return validModule(module) && path == "purescript/"+goName+"_ffi.go"
	case "runtime":
		return module == "" && path == "gopurs_runtime/runtime.go"
	case "entry":
		return validModule(module) && (path == "main/main.go" || path == module+"/main/main.go")
	}
	return false
}

func validObject(obj Object) bool {
	return digestPattern.MatchString(obj.SHA256) && obj.Bytes >= 0 && obj.Bytes <= 9007199254740991
}

func validateManifest(m Manifest) error {
	if m.Format != Format || m.Schema != Schema {
		return fmt.Errorf("manifest-version")
	}
	if !digestPattern.MatchString(m.Snapshot.Key) || !digestPattern.MatchString(m.Snapshot.Compiler) {
		return fmt.Errorf("manifest-key")
	}
	if m.GoMod != "create-if-absent" || m.GoSum != "external" {
		return fmt.Errorf("module-ownership")
	}
	seen := map[string]bool{}
	previous := ""
	for _, out := range m.Outputs {
		if !outputPath(out.Kind, out.Module, out.Path) || !validObject(out.Object) || seen[out.Path] || out.Path < previous {
			return fmt.Errorf("output-inventory")
		}
		seen[out.Path] = true
		previous = out.Path
	}
	seen = map[string]bool{}
	for _, mod := range m.Modules {
		if !validModule(mod.Name) || seen[mod.Name] {
			return fmt.Errorf("module-inventory")
		}
		seen[mod.Name] = true
		for _, key := range []string{mod.Prepared, mod.OptimizerEnvironment, mod.EmitterEnvironment, mod.CallerDemands} {
			if !digestPattern.MatchString(key) {
				return fmt.Errorf("module-environment")
			}
		}
		if len(mod.Parts) != len(moduleCodecs) {
			return fmt.Errorf("module-parts")
		}
		for i, part := range mod.Parts {
			if part.Codec != moduleCodecs[i] || !validObject(part.Object) {
				return fmt.Errorf("module-codec")
			}
		}
	}
	if m.Diagnostics != nil && (m.Diagnostics.Codec != "gopurs/diagnostics-v1" || !validObject(m.Diagnostics.Object)) {
		return fmt.Errorf("diagnostic-codec")
	}
	return nil
}

// Lossless wire values have explicit, stable tags. Strings are UTF-16 code units
// (including isolated surrogates); Numbers are IEEE-754 bits, preserving -0/NaN.
// Domain codecs must additionally validate their named constructors before use.
func validateWire(data []byte) error {
	var value any
	if err := json.Unmarshal(data, &value); err != nil {
		return fmt.Errorf("wire-json: %w", err)
	}
	var check func(any, int) bool
	check = func(v any, depth int) bool {
		if depth > 4096 {
			return false
		}
		a, ok := v.([]any)
		if !ok || len(a) == 0 {
			return false
		}
		tag, ok := a[0].(string)
		if !ok {
			return false
		}
		str := func(i int) string { s, _ := a[i].(string); return s }
		switch tag {
		case "unit":
			return len(a) == 1
		case "boolean":
			if len(a) != 2 {
				return false
			}
			_, ok := a[1].(bool)
			return ok
		case "int":
			if len(a) != 2 || !integerPattern.MatchString(str(1)) {
				return false
			}
			_, err := strconv.ParseInt(str(1), 10, 64)
			return err == nil
		case "number":
			return len(a) == 2 && floatPattern.MatchString(str(1))
		case "string":
			if len(a) != 2 {
				return false
			}
			s, ok := a[1].(string)
			return ok && hexUnitsPattern.MatchString(s)
		case "array", "ctor":
			i := 1
			if tag == "ctor" {
				if len(a) != 3 || !tagPattern.MatchString(str(1)) {
					return false
				}
				i = 2
			} else if len(a) != 2 {
				return false
			}
			items, ok := a[i].([]any)
			if !ok {
				return false
			}
			for _, item := range items {
				if !check(item, depth+1) {
					return false
				}
			}
			return true
		case "record":
			if len(a) != 2 {
				return false
			}
			items, ok := a[1].([]any)
			if !ok {
				return false
			}
			last := ""
			first := true
			for _, item := range items {
				pair, ok := item.([]any)
				if !ok || len(pair) != 2 {
					return false
				}
				key, ok := pair[0].(string)
				if !ok || !hexUnitsPattern.MatchString(key) || (!first && key <= last) || !check(pair[1], depth+1) {
					return false
				}
				first = false
				last = key
			}
			return true
		}
		return false
	}
	if !check(value, 0) {
		return fmt.Errorf("wire-shape")
	}
	return nil
}

func sortedSettings(settings []Setting) ([]string, error) {
	copy := append([]Setting{}, settings...)
	sort.Slice(copy, func(i, j int) bool { return copy[i].Name < copy[j].Name })
	parts := []string{}
	for i, setting := range copy {
		if setting.Name == "" || (i > 0 && setting.Name == copy[i-1].Name) {
			return nil, fmt.Errorf("duplicate/empty setting")
		}
		parts = append(parts, setting.Name, setting.Value)
	}
	return parts, nil
}
