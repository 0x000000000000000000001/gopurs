package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"gopurs/c1c"
)

// A generated module can span several planner packages. Keep those placements
// and replace its sections in place: private caches, shared callers and the
// initialization order then remain in their original packages. No import or
// declaration edge may change, and no type/init change is accepted here.
func preparePackageSections(stage string, state *buildSession, entries []string, request, previousRequest *packageFileRequest, previous *packagePlanResult, index *overlayIndex, changed packageSource) (*packageOverlayRecord, string, error) {
	return preparePackageSectionsProven(stage, state, entries, request, previousRequest, previous, index, changed, nil)
}

func preparePackageSectionsProven(stage string, state *buildSession, entries []string, request, previousRequest *packageFileRequest, previous *packagePlanResult, index *overlayIndex, changed packageSource, literals *c1c.LiteralChanges) (*packageOverlayRecord, string, error) {
	if _, err := topologicalPackages(previous.Packages); err != nil {
		return nil, "published package graph is not reusable", nil
	}
	source, err := os.ReadFile(changed.Source)
	if err != nil {
		return nil, "", err
	}
	if digest(source) != changed.SHA256 {
		return nil, "generated source changed during preparation", nil
	}
	if strings.Contains(string(source), "//go:") {
		return nil, "module contains a Go directive", nil
	}
	if literals == nil {
		literals = packageLiteralChanges(stage, state, previousRequest, changed, source)
	}

	references := index.sections[changed.Path]
	placements := map[string]string{}
	oldDecls := map[string]*moduleDecls{}
	for _, reference := range references {
		bundle := reference.bundle
		if !bundle.coalesced || oldDecls[bundle.file.Path] != nil {
			return nil, "ambiguous published source partition", nil
		}
		for _, alias := range bundle.imports {
			if alias == "." {
				return nil, "published package uses a dot import", nil
			}
		}
		declarations, err := inspectModuleSource(sectionSource(reference), index.symbolLookup(bundle.packageName), bundle.aliases, externalAliases(bundle.imports))
		if err != nil {
			return nil, "published module partition does not parse", nil
		}
		if declarations.hasType || declarations.hasMethod || declarations.hasInit || declarations.hasVarValue {
			return nil, "module partition has a type, method, init or initialized variable", nil
		}
		if !reflect.DeepEqual(declarations.names, reference.section.names) {
			return nil, "published declaration index is ambiguous", nil
		}
		for name := range declarations.names {
			if _, exists := placements[name]; exists {
				return nil, "duplicate published declaration: " + name, nil
			}
			placements[name] = bundle.packageName
		}
		oldDecls[bundle.file.Path] = declarations
	}
	newDecls, err := inspectModuleSource(string(source), nil, nil, nil)
	if err != nil {
		return nil, "new module source does not parse", nil
	}
	if len(newDecls.names) != len(placements) {
		return nil, "declaration-set-changed", nil
	}
	for name := range newDecls.names {
		if _, exists := placements[name]; !exists {
			return nil, "declaration-set-changed", nil
		}
	}
	replacements := map[string]string{}
	for _, reference := range references {
		bundle := reference.bundle
		fragment := ""
		var err error
		if literals != nil {
			var rewritten string
			rewritten, err = literals.Apply(sectionSource(reference))
			fragment = strings.TrimPrefix(rewritten, "package "+bundle.packageName+"\n")
		} else {
			symbols := index.symbolsFor(bundle.packageName)
			for name, pkg := range placements {
				symbols[name] = pkg
			}
			fragment, err = c1c.PartitionSource(string(source), bundle.packageName, placements, symbols, bundle.imports, func(pkg string) string {
				return request.PackagePath + "/" + pkg
			})
		}
		if err != nil {
			return nil, "module partition cannot reuse published imports: " + err.Error(), nil
		}
		newSection := "package " + bundle.packageName + "\n" + fragment
		declarations, err := inspectModuleSource(newSection, index.symbolLookup(bundle.packageName), bundle.aliases, externalAliases(bundle.imports))
		if err != nil {
			return nil, "rewritten module partition does not parse", nil
		}
		if reason, ok := compareDeclarations(oldDecls[bundle.file.Path], declarations); !ok {
			return nil, reason, nil
		}
		oldSignature, oldErr := sourceSignature(changed.Path, []byte(sectionSource(reference)))
		newSignature, newErr := sourceSignature(changed.Path, []byte(newSection))
		if oldErr != nil || newErr != nil || oldSignature != newSignature {
			return nil, "module partition signature changed", nil
		}
		marker := sourceMarker + hex.EncodeToString([]byte(changed.Path)) + "\n"
		start := strings.Index(bundle.source, marker)
		if start < 0 || strings.Count(bundle.source, marker) != 1 {
			return nil, "ambiguous published source marker", nil
		}
		start += len(marker)
		end := strings.Index(bundle.source[start:], sourceEndMarker)
		if end < 0 || bundle.source[start:start+end] != reference.section.source {
			return nil, "published source section changed", nil
		}
		replacement := bundle.source[:start] + fragment + bundle.source[start+end:]
		if replacement != bundle.source {
			replacements[bundle.file.Path] = replacement
		}
	}
	directory := filepath.Join(stage, "packages")
	if err := os.Mkdir(directory, 0700); err != nil {
		return nil, "", err
	}
	complete := false
	defer func() {
		if !complete {
			os.RemoveAll(directory)
			os.Remove(filepath.Join(stage, "package-result.json"))
		}
	}()
	files := append([]packagePlanFile{}, previous.Files...)
	for i := range files {
		file := &files[i]
		destination := filepath.Join(directory, filepath.FromSlash(file.Path))
		if replacement, ok := replacements[file.Path]; ok {
			file.SHA256, file.Bytes = digest([]byte(replacement)), int64(len(replacement))
			if err := writePackagesFile(destination, replacement); err != nil {
				return nil, "", err
			}
		} else if err := linkOrCopy(filepath.Join(state.Spec.Output, "purescript", filepath.FromSlash(file.Path)), destination); err != nil {
			return nil, "", err
		}
	}
	result := *previous
	result.Files = files
	result.PlanHash = Key("go-package-sections", previous.PlanHash, changed.Path, changed.SHA256)
	encoded, err := json.Marshal(result)
	if err != nil {
		return nil, "", err
	}
	if err := os.WriteFile(filepath.Join(stage, "package-result.json"), encoded, 0400); err != nil {
		return nil, "", err
	}
	if _, err := packageOutputs(stage, state.Modules, entries); err != nil {
		return nil, "section result does not satisfy the package contract: " + err.Error(), nil
	}
	packages := []string{}
	for _, reference := range references {
		packages = append(packages, reference.bundle.packageName)
	}
	sort.Strings(packages)
	complete = true
	// There is no new overlay package. Leave package-overlay.json absent so
	// publication clears any old relocation record; the next edit revalidates
	// the current section partition and its newly published content hashes.
	return &packageOverlayRecord{Format: packageOverlayFormat, PlanHash: result.PlanHash, Module: changed.Module,
		SourcePath: changed.Path, Package: fmt.Sprintf("retained-%d", len(packages)), OriginalPackage: strings.Join(packages, ",")}, "", nil
}

// The public plan is a hint. Skipping its global symbol scan additionally
// requires every published file to match the frozen committed generation.
// readPackagePlanFiles has already checked the bytes against these hashes;
// packageOutputs checks the staged bytes again before accepting the overlay.
func committedPackageFiles(stage string, state buildSession, result *packagePlanResult) bool {
	if !lazyModules(state.Spec) {
		return false
	}
	prior, err := getLazyPrior(stage, state)
	if err != nil {
		return false
	}
	outputs := map[string]OwnedOutput{}
	for _, output := range prior.manifest.Outputs {
		switch output.Kind {
		case "package", "ffi-package", "order", "package-init":
			outputs[output.Path] = output
		}
	}
	if len(outputs) != len(result.Files) {
		return false
	}
	for _, file := range result.Files {
		path := "purescript/" + file.Path
		output, ok := outputs[path]
		if !ok || output.Kind != file.Kind || output.Module != file.Module || output.Object != (Object{file.SHA256, file.Bytes}) {
			return false
		}
		delete(outputs, path)
	}
	return len(outputs) == 0
}

func packageLiteralChanges(stage string, state *buildSession, previousRequest *packageFileRequest, changed packageSource, source []byte) *c1c.LiteralChanges {
	priorPath := filepath.Join(stage, "prior", changed.Module, "module.go")
	if _, err := os.Stat(priorPath); os.IsNotExist(err) && lazyModules(state.Spec) {
		// A miss only disables this proof; the ordinary partition path below
		// can still handle modules without planner-owned field rewrites.
		_ = materializePrior(stage, *state, []string{changed.Module}, []string{"module"})
	}
	if raw, err := os.ReadFile(priorPath); err == nil {
		for _, file := range previousRequest.Files {
			if file.Path == changed.Path && file.Module == changed.Module && digest(raw) == file.SHA256 {
				literals, _ := c1c.CompareBodyLiterals(string(raw), string(source))
				return literals
			}
		}
	}
	return nil
}
