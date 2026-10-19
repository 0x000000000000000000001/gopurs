package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/parser"
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
	Format      string          `json:"format"`
	PackagePath string          `json:"packagePath"`
	GoVersion   string          `json:"goVersion"`
	Arch        string          `json:"arch"`
	Files       []packageSource `json:"files"`
	Exports     []packageExport `json:"exports"`
	Entries     []packageSymbol `json:"entries"`
	Output      string          `json:"output"`
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

func preparePackageSession(stage string, state *buildSession, entries []string, response *SessionResponse) error {
	if !state.Planned || !usesGoPackages(state.Spec) {
		return fmt.Errorf("Go package inputs require a planned package-layout build")
	}
	request := packageFileRequest{Format: packageFilesFormat, PackagePath: "gopurs/output/purescript",
		Files: []packageSource{}, Exports: []packageExport{}, Entries: []packageSymbol{}, Output: filepath.Join(stage, "packages")}
	imports := map[string]bool{}
	for _, module := range state.Modules {
		unit := strings.ReplaceAll(module.Name, ".", "_")
		for _, suffix := range []string{"", "_ffi"} {
			if suffix != "" && !module.HasForeign {
				continue
			}
			name := unit + suffix + ".go"
			file := filepath.Join(stage, "output", "purescript", name)
			data, err := os.ReadFile(file)
			if err != nil {
				return err
			}
			owner := "generated/" + unit
			if suffix != "" {
				owner = "ffi/" + unit
			}
			request.Files = append(request.Files, packageSource{owner, module.Name, name, file, digest(data), []packageSymbol{}, []any{}})
			if err := sourceImports(file, data, imports); err != nil {
				return err
			}
		}
	}
	for _, module := range entries {
		unit := strings.ReplaceAll(module, ".", "_")
		request.Entries = append(request.Entries, packageSymbol{"generated/" + unit, "Get_" + unit + "_main"})
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
	var environment map[string]string
	for _, option := range state.Spec.Options {
		if option.Name == "go-environment" {
			if err := json.Unmarshal([]byte(option.Value), &environment); err != nil {
				return err
			}
		}
	}
	request.Arch = environment["GOARCH"]
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
	data, err = json.Marshal(request)
	if err != nil {
		return err
	}
	response.PackageRequest = filepath.Join(stage, "package-request.json")
	response.Status = "packages-ready"
	return os.WriteFile(response.PackageRequest, data, 0400)
}

func packageOutputs(stage string, modules []SourceModule, entries []string) ([]StagedOutput, error) {
	var result struct {
		Format   string `json:"format"`
		PlanHash string `json:"planHash"`
		Packages []struct {
			Name         string   `json:"name"`
			Dependencies []string `json:"dependencies"`
			Declarations int      `json:"declarations"`
		} `json:"packages"`
		Files []struct {
			Path   string `json:"path"`
			Kind   string `json:"kind"`
			Module string `json:"moduleName"`
			SHA256 string `json:"sha256"`
			Bytes  int64  `json:"bytes"`
		} `json:"files"`
		Entries []struct {
			Symbol  packageSymbol `json:"symbol"`
			Package string        `json:"package"`
		} `json:"entries"`
	}
	data, err := os.ReadFile(filepath.Join(stage, "package-result.json"))
	if err != nil {
		return nil, err
	}
	if err := strictJSON(data, &result); err != nil {
		return nil, err
	}
	if result.Format != packageFilesFormat || !digestPattern.MatchString(result.PlanHash) || len(result.Packages) == 0 {
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
