package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

const packageLayout = "gopurs/go-packages-v1"
const packageFilesFormat = "gopurs/go-package-files-v1"

func usesGoPackages(spec Spec) bool {
	for _, setting := range spec.Options {
		if setting.Name == "go-layout" {
			return setting.Value == packageLayout
		}
	}
	return false
}

func packageGoCommand(directory string, args ...string) ([]byte, error) {
	command := exec.Command("go", args...)
	command.Dir = directory
	command.Env = append(os.Environ(), "GOWORK=off")
	output, err := command.Output()
	if failure, ok := err.(*exec.ExitError); ok {
		return nil, fmt.Errorf("go %s: %w\n%s", strings.Join(args, " "), err, failure.Stderr)
	}
	return output, err
}

// Select and fingerprint the Go environment before a whole-build cache hit.
// GOCACHE is an implementation detail, not part of the generated program.
func packageToolchain(spec *Spec) error {
	data, err := packageGoCommand(spec.Output, "env", "-json", "GOOS", "GOARCH", "GOAMD64", "GOARM64", "GOARM", "GOMIPS", "GOMIPS64", "GOPPC64", "GORISCV64", "GOWASM", "GOEXPERIMENT", "CGO_ENABLED", "GOFLAGS", "GOVERSION", "GOROOT", "GOPATH", "GOMODCACHE", "CC", "CXX", "CGO_CFLAGS", "CGO_CPPFLAGS", "CGO_CXXFLAGS", "CGO_LDFLAGS")
	if err != nil {
		return err
	}
	var environment map[string]string
	if err := json.Unmarshal(data, &environment); err != nil {
		return err
	}
	spec.Options = append(spec.Options, Setting{"go-environment", string(data)})
	spec.Compiler = append(spec.Compiler, Input{Name: "go-toolchain", Kind: "file", Path: filepath.Join(environment["GOROOT"], "bin", "go")})
	return nil
}

type packageSymbol struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type packageSource struct {
	Owner        string          `json:"owner"`
	Module       string          `json:"moduleName"`
	Path         string          `json:"path"`
	Source       string          `json:"source"`
	SHA256       string          `json:"sha256"`
	Initializers []packageSymbol `json:"initializers"`
	Records      []any           `json:"records"`
}

type packageExport struct {
	Path   string `json:"path"`
	Source string `json:"source"`
	SHA256 string `json:"sha256"`
}

type packageFileRequest struct {
	Format      string            `json:"format"`
	PackagePath string            `json:"packagePath"`
	GoVersion   string            `json:"goVersion"`
	Arch        string            `json:"arch"`
	Files       []packageSource   `json:"files"`
	Exports     []packageExport   `json:"exports"`
	Entries     []packageSymbol   `json:"entries"`
	Signatures  map[string]string `json:"signatures,omitempty"`
	Output      string            `json:"output"`
}

func sourceImports(name string, data []byte, imports map[string]bool) error {
	file, err := parser.ParseFile(token.NewFileSet(), name, data, parser.ImportsOnly)
	if err != nil {
		return err
	}
	for _, spec := range file.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			return err
		}
		imports[path] = true
	}
	return nil
}

func (resolver *ffiResolver) freezeGoTree(source, destination string, ancestors map[string]bool) error {
	entry, err := resolver.observe("directory", source)
	if err != nil {
		return err
	}
	if entry.fingerprint.State != "directory" {
		return fmt.Errorf("native Go dependency is not a directory: %s", source)
	}
	if ancestors[entry.fingerprint.Resolved] {
		return fmt.Errorf("cyclic native Go dependency directory: %s", source)
	}
	ancestors[entry.fingerprint.Resolved] = true
	defer delete(ancestors, entry.fingerprint.Resolved)
	if err := os.MkdirAll(destination, 0700); err != nil {
		return err
	}
	for _, child := range entry.entries {
		if child.Name() == ".git" {
			continue
		}
		from, to := filepath.Join(source, child.Name()), filepath.Join(destination, child.Name())
		info, err := os.Stat(from)
		if err != nil {
			return err
		}
		if info.IsDir() {
			if err := resolver.freezeGoTree(from, to, ancestors); err != nil {
				return err
			}
			continue
		}
		file, err := resolver.observe("file", from)
		if err != nil {
			return err
		}
		if err := os.WriteFile(to, file.bytes, 0600); err != nil {
			return err
		}
	}
	return nil
}

// Archives include machine code that the type checker never reads. Preserve a
// valid ar header and just __.PKGDEF, rather than copying each full object file.
func goExportData(file string) ([]byte, error) {
	input, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer input.Close()
	header := make([]byte, 68)
	if _, err := io.ReadFull(input, header); err != nil {
		return nil, err
	}
	if string(header[:8]) != "!<arch>\n" || strings.TrimSpace(string(header[8:24])) != "__.PKGDEF" || string(header[66:68]) != "`\n" {
		return nil, fmt.Errorf("Go export archive has no leading __.PKGDEF: %s", file)
	}
	size, err := strconv.ParseInt(strings.TrimSpace(string(header[56:66])), 10, 64)
	if err != nil || size < 0 {
		return nil, fmt.Errorf("invalid Go export archive size: %s", file)
	}
	data, err := io.ReadAll(io.LimitReader(input, size+size%2))
	if err != nil || int64(len(data)) != size+size%2 {
		return nil, fmt.Errorf("incomplete Go export archive: %s", file)
	}
	return append(header, data...), nil
}

func sourceSignature(path string, data []byte) (string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, data, parser.SkipObjectResolution)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	for _, declaration := range file.Decls {
		switch value := declaration.(type) {
		case *ast.FuncDecl:
			clone := *value
			clone.Doc = nil
			clone.Body = nil
			if err := printer.Fprint(&out, fset, &clone); err != nil {
				return "", err
			}
		case *ast.GenDecl:
			clone := *value
			clone.Doc = nil
			if err := printer.Fprint(&out, fset, &clone); err != nil {
				return "", err
			}
		}
		out.WriteByte('\n')
	}
	sum := sha256.Sum256(out.Bytes())
	return hex.EncodeToString(sum[:]), nil
}

type packagePlanPackage struct {
	Name         string   `json:"name"`
	Dependencies []string `json:"dependencies"`
	Declarations int      `json:"declarations"`
}

type packagePlanFile struct {
	Path   string `json:"path"`
	Kind   string `json:"kind"`
	Module string `json:"moduleName"`
	SHA256 string `json:"sha256"`
	Bytes  int64  `json:"bytes"`
}

type packagePlanEntry struct {
	Symbol  packageSymbol `json:"symbol"`
	Package string        `json:"package"`
}

type packagePlanResult struct {
	Format   string               `json:"format"`
	PlanHash string               `json:"planHash"`
	Packages []packagePlanPackage `json:"packages"`
	Files    []packagePlanFile    `json:"files"`
	Entries  []packagePlanEntry   `json:"entries"`
}

func readPackagePlan(path string) (*packagePlanResult, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result packagePlanResult
	if err := strictJSON(data, &result); err != nil {
		return nil, err
	}
	if result.Format != packageFilesFormat || !digestPattern.MatchString(result.PlanHash) || len(result.Packages) == 0 {
		return nil, fmt.Errorf("invalid emitted Go package plan")
	}
	for _, file := range result.Files {
		if !outputPath(file.Kind, file.Module, "purescript/"+file.Path) {
			return nil, fmt.Errorf("invalid emitted Go package path: %s", file.Path)
		}
	}
	return &result, nil
}

// Retain the exact verified bytes for the overlay. Parsing and proof consume
// this private snapshot; staged links are independently checked before commit.
type packagePlanSnapshot struct {
	request *packageFileRequest
	result  *packagePlanResult
	sources map[string]string
}

func readPackagePlanFiles(output string, result *packagePlanResult) (map[string]string, error) {
	sources := make(map[string]string, len(result.Files))
	for _, file := range result.Files {
		path := filepath.Join(output, "purescript", filepath.FromSlash(file.Path))
		info, err := os.Lstat(path)
		if !outputPath(file.Kind, file.Module, "purescript/"+file.Path) || err != nil || !info.Mode().IsRegular() || info.Size() != file.Bytes {
			return nil, fmt.Errorf("published Go package file changed: %s", file.Path)
		}
		data, err := os.ReadFile(path)
		if err != nil || int64(len(data)) != file.Bytes || digest(data) != file.SHA256 {
			return nil, fmt.Errorf("published Go package file changed: %s", file.Path)
		}
		sources[file.Path] = string(data)
	}
	return sources, nil
}

func planCacheStatus(output string, request *packageFileRequest) string {
	previous, _ := readPackageRequest(filepath.Join(output, ".gopurs-cache", "go-packages", "plan-v1", "request.json"))
	status, _ := inspectPackagePlan(output, request, previous)
	return status
}

func inspectPackagePlan(output string, request, previous *packageFileRequest) (string, *packagePlanSnapshot) {
	if previous == nil {
		return "miss", nil
	}
	dir := filepath.Join(output, ".gopurs-cache", "go-packages", "plan-v1")
	if previous.Format != request.Format || previous.PackagePath != request.PackagePath || previous.GoVersion != request.GoVersion || previous.Arch != request.Arch {
		return "miss-options", nil
	}
	previousFiles := map[string]string{}
	for _, file := range previous.Files {
		previousFiles[file.Path] = file.Owner + ":" + file.Module + ":" + file.SHA256
	}
	if len(previousFiles) != len(request.Files) {
		return "miss-sources", nil
	}
	sourceChanged := false
	for _, file := range request.Files {
		previous, ok := previousFiles[file.Path]
		if !ok {
			return "miss-sources", nil
		}
		if previous != file.Owner+":"+file.Module+":"+file.SHA256 {
			sourceChanged = true
		}
	}
	previousEntries := map[string]bool{}
	for _, entry := range previous.Entries {
		previousEntries[entry.Owner+"\x00"+entry.Name] = true
	}
	if len(previousEntries) != len(request.Entries) {
		return "miss-entries", nil
	}
	for _, entry := range request.Entries {
		if !previousEntries[entry.Owner+"\x00"+entry.Name] {
			return "miss-entries", nil
		}
	}
	if len(previous.Signatures) != len(request.Signatures) {
		return "miss-signatures", nil
	}
	signatureChanged := false
	for path, signature := range request.Signatures {
		if previous.Signatures[path] != signature {
			signatureChanged = true
		}
	}
	result, err := readPackagePlan(filepath.Join(dir, "result.json"))
	if err != nil {
		return "miss-result", nil
	}
	sources, err := readPackagePlanFiles(output, result)
	if err != nil {
		return "miss-output", nil
	}
	snapshot := &packagePlanSnapshot{request: previous, result: result, sources: sources}
	if signatureChanged {
		return "miss-signatures", nil
	}
	if sourceChanged {
		return "eligible-body", snapshot
	}
	return "validated", snapshot
}

// Missing staged sources belong to modules whose emission was reused. Their
// raw hashes come from the frozen committed manifest, never the public plan
// hint. An overlay needs those references, not thousands of decompressed files.
func packageSources(stage string, state buildSession) ([]packageSource, error) {
	files := []packageSource{}
	for _, module := range state.Modules {
		unit := strings.ReplaceAll(module.Name, ".", "_")
		for _, suffix := range []string{"", "_ffi"} {
			if suffix != "" && !module.HasForeign {
				continue
			}
			name := unit + suffix + ".go"
			file := filepath.Join(stage, "output", "purescript", name)
			data, err := os.ReadFile(file)
			owner := "generated/" + unit
			kind := "module"
			if suffix != "" {
				owner = "ffi/" + unit
				kind = "ffi"
			}
			sha := ""
			if err == nil {
				sha = digest(data)
			} else if os.IsNotExist(err) && lazyModules(state.Spec) {
				object, err := priorSourceObject(stage, state, module.Name, kind)
				if err != nil {
					return nil, err
				}
				sha = object.SHA256
			} else {
				return nil, err
			}
			files = append(files, packageSource{owner, module.Name, name, file, sha, []packageSymbol{}, []any{}})
		}
	}
	return files, nil
}

// Only a missing signature or the full planner requires the original bytes.
// Materialization verifies both the compressed object and its raw digest. Keep
// the ordinary source path in requests so the external planner stays unchanged.
func readPackageSource(stage string, state buildSession, file packageSource) ([]byte, error) {
	data, err := os.ReadFile(file.Source)
	if os.IsNotExist(err) && lazyModules(state.Spec) {
		kind := "module"
		if strings.HasPrefix(file.Owner, "ffi/") {
			kind = "ffi"
		}
		if err := materializePrior(stage, state, []string{file.Module}, []string{kind}); err != nil {
			return nil, err
		}
		if err := linkOrCopy(filepath.Join(stage, "prior", file.Module, kind+".go"), file.Source); err != nil {
			return nil, err
		}
		data, err = os.ReadFile(file.Source)
	}
	if err != nil {
		return nil, err
	}
	if digest(data) != file.SHA256 {
		return nil, fmt.Errorf("Go package source changed during preparation: %s", file.Path)
	}
	return data, nil
}

func preparePackageSession(stage string, state *buildSession, entries []string, response *SessionResponse) error {
	if !state.Planned || !usesGoPackages(state.Spec) {
		return fmt.Errorf("Go package inputs require a planned package-layout build")
	}
	files, err := packageSources(stage, *state)
	if err != nil {
		return err
	}
	request := packageFileRequest{Format: packageFilesFormat, PackagePath: "gopurs/output/purescript",
		Files: files, Exports: []packageExport{}, Entries: []packageSymbol{}, Output: filepath.Join(stage, "packages")}
	request.Signatures = map[string]string{}
	type cachedSignature struct{ owner, module, sha, signature string }
	previousRequest, _ := readPackageRequest(filepath.Join(state.Spec.Output, ".gopurs-cache", "go-packages", "plan-v1", "request.json"))
	previousSignatures := map[string]cachedSignature{}
	if previousRequest != nil {
		for _, file := range previousRequest.Files {
			if signature, ok := previousRequest.Signatures[file.Path]; ok {
				previousSignatures[file.Path] = cachedSignature{file.Owner, file.Module, file.SHA256, signature}
			}
		}
	}
	for _, file := range request.Files {
		if previous, ok := previousSignatures[file.Path]; ok &&
			previous.owner == file.Owner && previous.module == file.Module && previous.sha == file.SHA256 {
			request.Signatures[file.Path] = previous.signature
			continue
		}
		data, err := readPackageSource(stage, *state, file)
		if err != nil {
			return err
		}
		signature, err := sourceSignature(file.Path, data)
		if err != nil {
			return err
		}
		request.Signatures[file.Path] = signature
	}
	for _, module := range entries {
		unit := strings.ReplaceAll(module, ".", "_")
		request.Entries = append(request.Entries, packageSymbol{"generated/" + unit, "Get_" + unit + "_main"})
	}
	var environment map[string]string
	for _, option := range state.Spec.Options {
		if option.Name == "go-environment" {
			if err := json.Unmarshal([]byte(option.Value), &environment); err != nil {
				return err
			}
		}
	}
	request.Arch = environment["GOARCH"]
	versionData, err := packageGoCommand(state.Spec.Output, "mod", "edit", "-json")
	if err != nil {
		return err
	}
	var version struct{ Go string }
	if err := json.Unmarshal(versionData, &version); err != nil {
		return err
	}
	request.GoVersion = "go" + version.Go
	status, plan := inspectPackagePlan(state.Spec.Output, &request, previousRequest)
	fmt.Fprintf(os.Stderr, "[gopurs] Go plan cache: %s\n", status)
	encoded, err := json.Marshal(request)
	if err != nil {
		return err
	}
	response.PackageRequest = filepath.Join(stage, "package-request.json")
	if err := os.WriteFile(response.PackageRequest, encoded, 0600); err != nil {
		return err
	}
	// The overlay only consumes generated sources and the published plan. Try
	// it before freezing Go exports so a proven body edit does not pay for the
	// full dependency type-check.
	if status == "eligible-body" && packageOverlayEnabled() {
		record, reason, err := preparePackageOverlaySnapshot(stage, state, entries, &request, plan)
		if err != nil {
			return err
		}
		if record != nil {
			response.PackageResult = filepath.Join(stage, "package-result.json")
			response.Status = "packages-overlay"
			fmt.Fprintf(os.Stderr, "[gopurs] Go overlay: module=%s package=%s source=%s\n", record.Module, record.Package, record.OriginalPackage)
			return nil
		}
		fmt.Fprintf(os.Stderr, "[gopurs] Go overlay fallback: %s\n", reason)
	}
	// Only the full planner consumes the native import closure. The overlay
	// proof is complete before this point and does not parse 2 704 headers.
	imports := map[string]bool{}
	for _, file := range request.Files {
		data, err := readPackageSource(stage, *state, file)
		if err != nil {
			return err
		}
		if err := sourceImports(file.Source, data, imports); err != nil {
			return err
		}
	}
	workspace := filepath.Join(stage, "go-exports")
	if err := os.Mkdir(workspace, 0700); err != nil {
		return err
	}
	resolver := ffiResolver{workspace: state.Spec.Workspace, observed: map[string]observation{}}
	for _, name := range []string{"go.mod", "go.sum"} {
		file, err := resolver.observe("file", filepath.Join(state.Spec.Output, name))
		if err != nil {
			return err
		}
		data := file.bytes
		if file.fingerprint.State == "missing" {
			if name == "go.sum" {
				continue
			}
			data = []byte("module gopurs/output\n\ngo 1.22\n")
		}
		if err := os.WriteFile(filepath.Join(workspace, name), data, 0600); err != nil {
			return err
		}
	}
	data, err := packageGoCommand(workspace, "mod", "edit", "-json")
	if err != nil {
		return err
	}
	var module struct {
		Module  struct{ Path string }
		Go      string
		Replace []struct {
			Old, New struct{ Path, Version string }
		}
	}
	if err := json.Unmarshal(data, &module); err != nil {
		return err
	}
	if module.Module.Path != "gopurs/output" {
		return fmt.Errorf("generated Go module must be gopurs/output, found %s", module.Module.Path)
	}
	request.GoVersion = "go" + module.Go
	for _, replacement := range module.Replace {
		if replacement.New.Version != "" {
			continue
		}
		original := absolute(state.Spec.Output, replacement.New.Path)
		frozen := filepath.Join(stage, "go-replacements", digest([]byte(replacement.Old.Path)))
		if err := resolver.freezeGoTree(original, frozen, map[string]bool{}); err != nil {
			return err
		}
		old := replacement.Old.Path
		if replacement.Old.Version != "" {
			old += "@" + replacement.Old.Version
		}
		if _, err := packageGoCommand(workspace, "mod", "edit", "-replace", old+"="+frozen); err != nil {
			return err
		}
	}
	// Runtime is a producer output, not an external input. Native companion
	// packages are frozen from the public module before collecting their exports.
	data, err = os.ReadFile(filepath.Join(stage, "output", "gopurs_runtime", "runtime.go"))
	if err != nil {
		return err
	}
	if err := os.Mkdir(filepath.Join(workspace, "gopurs_runtime"), 0700); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(workspace, "gopurs_runtime", "runtime.go"), data, 0600); err != nil {
		return err
	}
	copied := map[string]bool{"gopurs_runtime": true}
	paths := make([]string, 0, len(imports))
	for path := range imports {
		if path != "unsafe" {
			paths = append(paths, path)
		}
	}
	sort.Strings(paths)
	for len(paths) > 0 {
		// Go selects files by target/build tags and excludes tests. Resolve
		// missing local roots until the complete native import closure exists.
		args := append([]string{"list", "-deps", "-e", "-json", "-mod=readonly"}, paths...)
		data, err := packageGoCommand(workspace, args...)
		if err != nil {
			return err
		}
		decoder := json.NewDecoder(bytes.NewReader(data))
		pending := map[string]bool{}
		for {
			var pkg struct{ ImportPath string }
			if err := decoder.Decode(&pkg); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			if local, ok := strings.CutPrefix(pkg.ImportPath, "gopurs/output/"); ok {
				name, _, _ := strings.Cut(local, "/")
				if name == "purescript" || name == "main" || name == "" || strings.ContainsAny(name, ".\\") {
					return fmt.Errorf("invalid native Go package dependency: %s", pkg.ImportPath)
				}
				if !copied[name] {
					pending[name] = true
				}
			}
		}
		if len(pending) == 0 {
			break
		}
		names := make([]string, 0, len(pending))
		for name := range pending {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			if err := resolver.freezeGoTree(filepath.Join(state.Spec.Output, name), filepath.Join(workspace, name), map[string]bool{}); err != nil {
				return err
			}
			copied[name] = true
		}
	}
	if len(paths) > 0 {
		args := []string{"list", "-deps", "-export", "-json", "-mod=readonly"}
		if environment["GOFLAGS"] == "" {
			// The binder consumes types, not inline bodies or machine code.
			// Split linker objects out of these archives and discard them; Go's
			// action keys include the flags, so ordinary builds have distinct
			// cache entries. Local source paths are irrelevant to binding, too.
			// Explicit caller flags retain their existing policy.
			// Independent type-only dependency builds use a small process cap;
			// the driver's selected Go runtime limits are inherited by each one.
			args = append(args, "-p=8", "-trimpath", "-gcflags=all=-N -l -linkobj="+os.DevNull)
		} else {
			args = append(args, "-p=1")
		}
		args = append(args, paths...)
		data, err := packageGoCommand(workspace, args...)
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(stage, "go-exports.json"), data, 0600); err != nil {
			return err
		}
		decoder := json.NewDecoder(bytes.NewReader(data))
		for {
			var pkg struct{ ImportPath, Export string }
			if err := decoder.Decode(&pkg); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			if pkg.Export == "" {
				continue
			}
			data, err := goExportData(pkg.Export)
			if err != nil {
				return err
			}
			sha := digest(data)
			file := filepath.Join(workspace, sha+".a")
			if err := os.WriteFile(file, data, 0400); err != nil {
				return err
			}
			request.Exports = append(request.Exports, packageExport{pkg.ImportPath, file, sha})
		}
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
		return fmt.Errorf("inputs changed while freezing Go package dependencies; retry the build")
	}
	state.Snapshot = current
	if err := saveSession(stage, *state); err != nil {
		return err
	}
	encoded, err = json.Marshal(request)
	if err != nil {
		return err
	}
	if err := os.WriteFile(response.PackageRequest, encoded, 0400); err != nil {
		return err
	}
	response.Status = "packages-ready"
	return nil
}

func packageOutputs(stage string, modules []SourceModule, entries []string) ([]StagedOutput, error) {
	result, err := readPackagePlan(filepath.Join(stage, "package-result.json"))
	if err != nil {
		return nil, err
	}
	if len(result.Packages) == 0 {
		return nil, fmt.Errorf("invalid emitted Go package plan")
	}
	packages := map[string]bool{}
	for _, pkg := range result.Packages {
		if !goPackagePattern.MatchString(pkg.Name) || packages[pkg.Name] || pkg.Declarations < 1 {
			return nil, fmt.Errorf("invalid emitted Go package: %s", pkg.Name)
		}
		for _, dependency := range pkg.Dependencies {
			if !packages[dependency] {
				return nil, fmt.Errorf("non-topological emitted Go packages")
			}
		}
		packages[pkg.Name] = true
	}
	if len(result.Entries) != len(entries) {
		return nil, fmt.Errorf("emitted Go entry inventory changed")
	}
	for i, entry := range result.Entries {
		unit := strings.ReplaceAll(entries[i], ".", "_")
		pkg, ok := strings.CutPrefix(entry.Package, "gopurs/output/purescript/")
		if !ok || !packages[pkg] || entry.Symbol != (packageSymbol{"generated/" + unit, "Get_" + unit + "_main"}) {
			return nil, fmt.Errorf("emitted Go entry changed: %s", entries[i])
		}
	}
	owners := map[string]bool{}
	for _, module := range modules {
		owners[module.Name] = true
	}
	files := []StagedOutput{}
	seen := map[string]bool{}
	for _, file := range result.Files {
		public := "purescript/" + file.Path
		pkg, _, _ := strings.Cut(file.Path, "/")
		if seen[public] || !outputPath(file.Kind, file.Module, public) ||
			(file.Kind != "package-init" && !packages[pkg]) || (file.Module != "" && !owners[file.Module]) {
			return nil, fmt.Errorf("unowned emitted Go file: %s", file.Path)
		}
		seen[public] = true
		source := filepath.Join(stage, "packages", filepath.FromSlash(file.Path))
		if !fileMatches(source, Object{file.SHA256, file.Bytes}) {
			return nil, fmt.Errorf("emitted Go file changed before publication: %s", file.Path)
		}
		files = append(files, StagedOutput{public, file.Kind, file.Module, source})
	}
	if !seen["purescript/init_all/init.go"] || !seen["purescript/init_external/init.go"] {
		return nil, fmt.Errorf("missing Go initialization packages")
	}
	for pkg := range packages {
		if !seen["purescript/"+pkg+"/zz_gopurs_order.go"] {
			return nil, fmt.Errorf("missing Go package order file: %s", pkg)
		}
	}
	return files, nil
}

// Grouped sources can replace tens of megabytes for a one-line edit. Once a
// package publication commits, retain only the current manifest's objects.
// Readers freeze their module inputs under this same lock into private stages;
// publishers put every object before committing, including reused inputs.
func (s store) retirePackageObjects() error {
	lockPath := filepath.Join(s.cache, "write.lock")
	if info, err := os.Lstat(lockPath); err != nil || !info.Mode().IsRegular() {
		return fmt.Errorf("cache lock is not a regular file")
	}
	lock, err := os.OpenFile(lockPath, os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer lock.Close()
	if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_EX); err != nil {
		return err
	}
	defer syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)
	if err := s.ensurePacks(); err != nil {
		return err
	}
	// Another publisher may have run since this session committed. Read its
	// current roots under the lock, and leave incomplete transactions alone.
	manifest, _ := s.readManifest()
	if manifest == nil || manifest.Snapshot.Output != s.root {
		return nil
	}
	owned, err := s.readOwnership()
	if err != nil {
		return err
	}
	if !sameOwnership(owned, manifest.Outputs) {
		return nil
	}
	data, err := os.ReadFile(filepath.Join(s.cache, "current.json"))
	if err != nil {
		return err
	}
	var current pointer
	if err := strictJSON(data, &current); err != nil {
		return err
	}
	retained := map[string]bool{current.Manifest.SHA256: true}
	for _, output := range manifest.Outputs {
		retained[output.Object.SHA256] = true
	}
	for _, artifact := range artifacts(*manifest) {
		retained[artifact.Object.SHA256] = true
	}
	for _, module := range manifest.Modules {
		for _, source := range module.Sources {
			retained[source.Object.SHA256] = true
		}
	}
	directory := filepath.Join(s.cache, "objects")
	info, err := os.Lstat(directory)
	if err != nil || !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
		return fmt.Errorf("cache directory is not a real directory: %s", directory)
	}
	entries, err := os.ReadDir(directory)
	if err != nil {
		return err
	}
	removed := false
	for _, entry := range entries {
		if entry.Type().IsRegular() && digestPattern.MatchString(entry.Name()) && !retained[entry.Name()] {
			if err := os.Remove(filepath.Join(directory, entry.Name())); err != nil {
				return err
			}
			removed = true
		}
	}
	if removed {
		if err := syncDirectory(directory); err != nil {
			return err
		}
	}
	return s.retirePackObjects(manifest)
}
