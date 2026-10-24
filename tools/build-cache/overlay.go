package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"gopurs/c1c"
)

const packageOverlayFormat = "gopurs/go-package-overlay-v1"
const packageBundleName = "0_gopurs_package.go"
const sourceMarker = "// gopurs-source-v1: "
const sourceEndMarker = "// gopurs-source-end-v1"

type packageOverlayRecord struct {
	Format          string   `json:"format"`
	PlanHash        string   `json:"planHash"`
	Module          string   `json:"moduleName"`
	SourcePath      string   `json:"sourcePath"`
	Package         string   `json:"package"`
	OriginalPackage string   `json:"originalPackage"`
	Symbols         []string `json:"symbols"`
}

func packageOverlayEnabled() bool {
	return os.Getenv("GOPURS_GO_PACKAGE_OVERLAY") == "1"
}

func readOverlayRecord(cache string) *packageOverlayRecord {
	data, err := os.ReadFile(filepath.Join(cache, "go-packages", "plan-v1", "overlay.json"))
	if err != nil {
		return nil
	}
	var record packageOverlayRecord
	if strictJSON(data, &record) != nil || record.Format != packageOverlayFormat || !digestPattern.MatchString(record.PlanHash) {
		return nil
	}
	return &record
}

type overlaySection struct {
	path   string
	source string
	names  map[string]bool
}

type overlayBundle struct {
	file        packagePlanFile
	packageName string
	source      string
	imports     map[string]string // import path -> effective alias
	aliases     map[string]string // effective alias -> generated package
	sections    []overlaySection
	coalesced   bool
}

type overlayIndex struct {
	bundles        map[string]*overlayBundle
	order          []string
	packageSymbols map[string]map[string]string
	exported       map[string]string
	sections       map[string][]overlaySectionRef
}

type overlaySectionRef struct {
	bundle  *overlayBundle
	section overlaySection
}

func extractSourceSections(source string) ([]overlaySection, error) {
	sections := []overlaySection{}
	position := 0
	for {
		start := strings.Index(source[position:], sourceMarker)
		if start < 0 {
			break
		}
		start += position
		lineEnd := strings.IndexByte(source[start:], '\n')
		if lineEnd < 0 {
			return nil, fmt.Errorf("unterminated Go source marker")
		}
		lineEnd += start
		encoded := strings.TrimSpace(source[start+len(sourceMarker) : lineEnd])
		decoded, err := hex.DecodeString(encoded)
		if err != nil || len(decoded) == 0 {
			return nil, fmt.Errorf("invalid Go source marker")
		}
		name := string(decoded)
		if strings.ContainsAny(name, "/\\\x00") || name == "." || name == ".." {
			return nil, fmt.Errorf("invalid Go source marker path: %s", name)
		}
		bodyStart := lineEnd + 1
		end := strings.Index(source[bodyStart:], sourceEndMarker)
		if end < 0 {
			return nil, fmt.Errorf("missing Go source end marker")
		}
		end += bodyStart
		sections = append(sections, overlaySection{path: name, source: source[bodyStart:end]})
		position = end + len(sourceEndMarker)
	}
	return sections, nil
}

// scanTopLevelNames is deliberately syntactic. The caller parses only the
// changed module; a whole-program AST walk would dominate the skipped planner.
// Braces in strings and comments are ignored so a column-zero local variable
// inside a function is not mistaken for a package declaration.
func scanTopLevelNames(source string) map[string]bool {
	names := map[string]bool{}
	depth, state, group := 0, 0, ""
	atLineStart := true
	process := func(line string) {
		if group != "" {
			if strings.HasPrefix(line, ")") {
				group = ""
				return
			}
			if line == "" || (line[0] != '\t' && line[0] != ' ') {
				return
			}
			trimmed := strings.TrimSpace(line)
			if trimmed == "" || strings.HasPrefix(trimmed, "//") {
				return
			}
			for _, field := range strings.Split(trimmed, ",") {
				name := field
				if index := strings.IndexAny(name, " \t=[]"); index >= 0 {
					name = name[:index]
				}
				if name != "" && name != "_" {
					names[name] = true
				}
			}
			return
		}
		keyword := ""
		for _, candidate := range []string{"func", "type", "var", "const"} {
			if strings.HasPrefix(line, candidate+" ") {
				keyword = candidate
				break
			}
		}
		if keyword == "" {
			return
		}
		rest := strings.TrimSpace(line[len(keyword):])
		if rest == "(" {
			group = keyword
			return
		}
		if keyword == "func" && strings.HasPrefix(rest, "(") {
			return
		}
		first := rest
		if index := strings.IndexAny(first, " \t(["); index >= 0 {
			first = first[:index]
		}
		for _, field := range strings.Split(first, ",") {
			name := strings.TrimSpace(field)
			if name != "" && name != "_" {
				names[name] = true
			}
		}
	}
	for i := 0; i < len(source); {
		character := source[i]
		if atLineStart && state == 0 && depth == 0 {
			end := strings.IndexByte(source[i:], '\n')
			line := source[i:]
			if end >= 0 {
				line = source[i : i+end]
			}
			process(line)
			atLineStart = false
		}
		switch state {
		case 0:
			switch {
			case character == '/' && i+1 < len(source) && source[i+1] == '/':
				state, i = 1, i+2
				continue
			case character == '/' && i+1 < len(source) && source[i+1] == '*':
				state, i = 2, i+2
				continue
			case character == '"':
				state = 3
			case character == '`':
				state = 4
			case character == '\'':
				state = 5
			case character == '{':
				depth++
			case character == '}':
				if depth > 0 {
					depth--
				}
			case character == '\n':
				atLineStart = true
			}
		case 1:
			if character == '\n' {
				state, atLineStart = 0, true
			}
		case 2:
			if character == '*' && i+1 < len(source) && source[i+1] == '/' {
				state, i = 0, i+2
				continue
			}
		case 3, 5:
			if character == '\\' {
				i++
			} else if (state == 3 && character == '"') || (state == 5 && character == '\'') {
				state = 0
			}
		case 4:
			if character == '`' {
				state = 0
			}
		}
		i++
	}
	return names
}

func parseBundleImports(source []byte, packageName string) (map[string]string, map[string]string, error) {
	file, err := parser.ParseFile(token.NewFileSet(), "bundle.go", source, parser.ImportsOnly)
	if err != nil {
		return nil, nil, err
	}
	if file.Name.Name != packageName {
		return nil, nil, fmt.Errorf("bundle package mismatch: %s", packageName)
	}
	imports := map[string]string{}
	aliases := map[string]string{}
	for _, spec := range file.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			return nil, nil, err
		}
		alias := ""
		if spec.Name != nil {
			alias = spec.Name.Name
		} else {
			alias = path[strings.LastIndex(path, "/")+1:]
		}
		imports[path] = alias
		if generated, ok := strings.CutPrefix(path, "gopurs/output/purescript/"); ok && !strings.ContainsAny(generated, "/\\") {
			aliases[alias] = generated
		}
	}
	return imports, aliases, nil
}

func loadOverlayIndex(result *packagePlanResult, sources map[string]verifiedPackageFile, overlayPackage string) (*overlayIndex, error) {
	return loadOverlayIndexForSource(result, sources, overlayPackage, "")
}

// A partial index is only sufficient for an in-place, proven literal change:
// neither names nor edges can change. Other edits need the complete inventory.
func loadOverlayIndexForSource(result *packagePlanResult, sources map[string]verifiedPackageFile, overlayPackage, onlySource string) (*overlayIndex, error) {
	index := &overlayIndex{
		bundles: map[string]*overlayBundle{}, packageSymbols: map[string]map[string]string{},
		exported: map[string]string{}, sections: map[string][]overlaySectionRef{},
	}
	bundles := make([]*overlayBundle, len(result.Files))
	if err := boundedEach(len(result.Files), func(position int) error {
		file := result.Files[position]
		if file.Kind != "package" && file.Kind != "ffi-package" {
			return nil
		}
		packageName, _, ok := strings.Cut(file.Path, "/")
		if !ok || !goPackagePattern.MatchString(packageName) {
			return fmt.Errorf("invalid published package path: %s", file.Path)
		}
		verified, ok := sources[file.Path]
		if !ok {
			return fmt.Errorf("missing verified package: %s", file.Path)
		}
		if verified.importError != nil {
			return verified.importError
		}
		if verified.imports == nil {
			return fmt.Errorf("missing verified package imports: %s", file.Path)
		}
		text := verified.source
		bundle := &overlayBundle{file: file, packageName: packageName, source: text, imports: verified.imports, aliases: verified.aliases}
		if filepath.Base(file.Path) == packageBundleName {
			sections, err := extractSourceSections(text)
			if err != nil {
				return err
			}
			if len(sections) > 0 {
				bundle.coalesced = true
				bundle.sections = sections
			}
		}
		if len(bundle.sections) == 0 {
			name := filepath.Base(file.Path)
			bundle.sections = []overlaySection{{path: name, source: text}}
		}
		for i := range bundle.sections {
			if onlySource == "" || bundle.sections[i].path == onlySource {
				bundle.sections[i].names = scanTopLevelNames(bundle.sections[i].source)
			}
		}
		bundles[position] = bundle
		return nil
	}); err != nil {
		return nil, err
	}
	counts := map[string]int{}
	total := 0
	for _, bundle := range bundles {
		if bundle != nil {
			for _, section := range bundle.sections {
				counts[bundle.packageName] += len(section.names)
				total += len(section.names)
			}
		}
	}
	index.exported = make(map[string]string, total)
	for name, count := range counts {
		index.packageSymbols[name] = make(map[string]string, count)
	}
	for _, bundle := range bundles {
		if bundle == nil {
			continue
		}
		file, packageName := bundle.file, bundle.packageName
		for i := range bundle.sections {
			index.sections[bundle.sections[i].path] = append(index.sections[bundle.sections[i].path], overlaySectionRef{bundle, bundle.sections[i]})
			for name := range bundle.sections[i].names {
				if name == "init" || name == "_" {
					continue
				}
				if index.packageSymbols[packageName] == nil {
					index.packageSymbols[packageName] = map[string]string{}
				}
				index.packageSymbols[packageName][name] = packageName
				if !ast.IsExported(name) {
					continue
				}
				if previous, ok := index.exported[name]; ok && previous != packageName {
					// A published overlay intentionally duplicates the old
					// declarations. The active overlay owns the exported name;
					// any other exported ambiguity rejects the delta.
					if previous == overlayPackage {
						continue
					}
					if packageName != overlayPackage {
						return nil, fmt.Errorf("ambiguous exported symbol: %s", name)
					}
				}
				index.exported[name] = packageName
			}
		}
		index.bundles[file.Path] = bundle
		index.order = append(index.order, file.Path)
	}
	return index, nil
}

type moduleDecls struct {
	names       map[string]bool
	references  map[string]map[string]bool
	imports     map[string]bool
	hasType     bool
	hasInit     bool
	hasMethod   bool
	hasVarValue bool
}

// inspectModuleSource resolves cross-package references with the same symbol
// table the planner owns. References are keyed by declaration so a body edit
// that keeps the declaration set but changes an edge is rejected.
func inspectModuleSource(source string, symbol func(string) (string, bool), aliases map[string]string, external map[string]string) (*moduleDecls, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "module.go", source, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	result := &moduleDecls{names: map[string]bool{}, references: map[string]map[string]bool{}, imports: map[string]bool{}}
	selectors := map[*ast.Ident]bool{}
	ast.Inspect(file, func(node ast.Node) bool {
		if selector, ok := node.(*ast.SelectorExpr); ok {
			selectors[selector.Sel] = true
		}
		return true
	})
	occurrences := map[string]int{}
	add := func(name string, node ast.Node) {
		key := name
		if occurrences[name] > 0 {
			key = name + "#" + strconv.Itoa(occurrences[name])
		}
		occurrences[name]++
		references := map[string]bool{}
		ast.Inspect(node, func(value ast.Node) bool {
			switch typed := value.(type) {
			case *ast.SelectorExpr:
				base, ok := typed.X.(*ast.Ident)
				if ok {
					if pkg, found := aliases[base.Name]; found {
						references[pkg+"\x00"+typed.Sel.Name] = true
					} else if path, found := external[base.Name]; found {
						result.imports[path] = true
					}
				}
			case *ast.Ident:
				if selectors[typed] || typed.Obj != nil || symbol == nil {
					return true
				}
				if pkg, found := symbol(typed.Name); found {
					references[pkg+"\x00"+typed.Name] = true
				}
			}
			return true
		})
		result.references[key] = references
	}
	for _, declaration := range file.Decls {
		switch value := declaration.(type) {
		case *ast.FuncDecl:
			if value.Recv != nil && len(value.Recv.List) > 0 {
				result.hasMethod = true
			}
			name := value.Name.Name
			if name == "init" {
				result.hasInit = true
			}
			result.names[name] = true
			add(name, value)
		case *ast.GenDecl:
			if value.Tok == token.TYPE {
				result.hasType = true
			}
			for _, item := range value.Specs {
				switch spec := item.(type) {
				case *ast.TypeSpec:
					result.names[spec.Name.Name] = true
					add(spec.Name.Name, spec)
				case *ast.ValueSpec:
					if value.Tok == token.VAR && len(spec.Values) > 0 {
						result.hasVarValue = true
					}
					for _, name := range spec.Names {
						if name.Name == "_" {
							continue
						}
						result.names[name.Name] = true
						add(name.Name, spec)
					}
				}
			}
		}
	}
	return result, nil
}

func (index *overlayIndex) activeSection(sourcePath string, result *packagePlanResult, overlay *packageOverlayRecord, entries []string) (overlaySectionRef, string, error) {
	if overlay != nil && overlay.SourcePath == sourcePath && overlay.PlanHash == result.PlanHash {
		// The overlay bundle is a complete package (package clause and imports
		// included), not a marked fragment. Its single section is the module.
		if bundle := index.bundles[overlay.Package+"/"+packageBundleName]; bundle != nil && len(bundle.sections) == 1 {
			return overlaySectionRef{bundle, bundle.sections[0]}, overlay.Package, nil
		}
		return overlaySectionRef{}, "", fmt.Errorf("overlay package is missing from the published plan")
	}
	candidates := index.sections[sourcePath]
	if len(candidates) == 0 {
		return overlaySectionRef{}, "", fmt.Errorf("changed source is not published as a generated section")
	}
	if len(candidates) == 1 {
		return candidates[0], candidates[0].bundle.packageName, nil
	}
	// A previous overlay can leave the original section next to its replacement.
	// Without a matching record, only an entry target disambiguates them.
	unit := strings.TrimSuffix(sourcePath, ".go")
	target := ""
	for _, entry := range entries {
		if strings.ReplaceAll(entry, ".", "_") == unit {
			for _, candidate := range candidates {
				for _, placement := range result.Entries {
					if placement.Symbol.Owner == "generated/"+unit && placement.Package == "gopurs/output/purescript/"+candidate.bundle.packageName {
						target = candidate.bundle.packageName
					}
				}
			}
		}
	}
	if target != "" {
		for _, candidate := range candidates {
			if candidate.bundle.packageName == target {
				return candidate, target, nil
			}
		}
	}
	return overlaySectionRef{}, "", fmt.Errorf("ambiguous published source section: %s", sourcePath)
}

func sectionSource(ref overlaySectionRef) string {
	if ref.bundle.coalesced {
		return "package " + ref.bundle.packageName + "\n" + ref.section.source
	}
	return ref.section.source
}

func compareDeclarations(old, new *moduleDecls) (string, bool) {
	if !reflect.DeepEqual(old.names, new.names) {
		return "declaration-set-changed", false
	}
	if !reflect.DeepEqual(old.references, new.references) {
		return "declaration-references-changed", false
	}
	if !reflect.DeepEqual(old.imports, new.imports) {
		return "module-imports-changed", false
	}
	return "", true
}

func externalAliases(imports map[string]string) map[string]string {
	out := map[string]string{}
	for path, alias := range imports {
		if strings.HasPrefix(path, "gopurs/output/purescript/") || alias == "_" || alias == "." {
			continue
		}
		out[alias] = path
	}
	return out
}

// Lookup keeps local precedence without copying every exported declaration
// for each section of the changed module. The index is immutable after load.
func (index *overlayIndex) symbolLookup(pkg string) func(string) (string, bool) {
	local := index.packageSymbols[pkg]
	return func(name string) (string, bool) {
		if owner, found := local[name]; found {
			return owner, true
		}
		owner, found := index.exported[name]
		return owner, found
	}
}

func (index *overlayIndex) symbolsFor(pkg string) map[string]string {
	symbols := map[string]string{}
	for name, owner := range index.packageSymbols[pkg] {
		symbols[name] = owner
	}
	for name, owner := range index.exported {
		if _, local := symbols[name]; !local {
			symbols[name] = owner
		}
	}
	return symbols
}

func exportedNames(names map[string]bool) []string {
	out := []string{}
	for name := range names {
		if ast.IsExported(name) {
			out = append(out, name)
		}
	}
	sort.Strings(out)
	return out
}

func identifierByte(value byte) bool {
	return value == '_' || value >= '0' && value <= '9' || value >= 'A' && value <= 'Z' || value >= 'a' && value <= 'z'
}

func anyIdentifier(source string, wanted map[string]bool) bool {
	for i := 0; i < len(source); {
		if !identifierByte(source[i]) || source[i] >= '0' && source[i] <= '9' {
			i++
			continue
		}
		start := i
		for i < len(source) && identifierByte(source[i]) {
			i++
		}
		if wanted[source[start:i]] {
			return true
		}
	}
	return false
}

func freeRewriteAlias(source string) string {
	alias := "c1c_overlay"
	for strings.Contains(source, alias) {
		alias += "_"
	}
	return alias
}

func overlayPackageName(module string) string {
	sum := sha256.Sum256([]byte("gopurs/go-package-overlay/" + module))
	return "p_" + hex.EncodeToString(sum[:])[:32]
}

func readPackageRequest(path string) (*packageFileRequest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var request packageFileRequest
	if err := strictJSON(data, &request); err != nil {
		return nil, err
	}
	return &request, nil
}

func changedPackageSource(previous, current *packageFileRequest) (packageSource, string) {
	oldFiles := map[string]packageSource{}
	for _, file := range previous.Files {
		oldFiles[file.Path] = file
	}
	changed := []packageSource{}
	for _, file := range current.Files {
		old, ok := oldFiles[file.Path]
		if !ok || old.Owner != file.Owner || old.Module != file.Module || old.SHA256 != file.SHA256 {
			changed = append(changed, file)
		}
	}
	if len(changed) != 1 {
		return packageSource{}, "expected exactly one changed Go source"
	}
	return changed[0], ""
}

func topologicalPackages(packages []packagePlanPackage) ([]packagePlanPackage, error) {
	byName := map[string]packagePlanPackage{}
	for _, pkg := range packages {
		if _, exists := byName[pkg.Name]; exists {
			return nil, fmt.Errorf("duplicate package: %s", pkg.Name)
		}
		byName[pkg.Name] = pkg
	}
	state := map[string]int{}
	order := []string{}
	var visit func(string) error
	visit = func(name string) error {
		switch state[name] {
		case 1:
			return fmt.Errorf("cyclic package graph at %s", name)
		case 2:
			return nil
		}
		pkg, ok := byName[name]
		if !ok {
			return fmt.Errorf("missing package dependency: %s", name)
		}
		state[name] = 1
		dependencies := append([]string{}, pkg.Dependencies...)
		sort.Strings(dependencies)
		for _, dependency := range dependencies {
			if err := visit(dependency); err != nil {
				return err
			}
		}
		state[name] = 2
		order = append(order, name)
		return nil
	}
	names := make([]string, 0, len(byName))
	for name := range byName {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		if err := visit(name); err != nil {
			return nil, err
		}
	}
	result := make([]packagePlanPackage, 0, len(order))
	for _, name := range order {
		result = append(result, byName[name])
	}
	return result, nil
}

func writePackagesFile(destination, content string) error {
	if err := os.MkdirAll(filepath.Dir(destination), 0700); err != nil {
		return err
	}
	return os.WriteFile(destination, []byte(content), 0600)
}

// Reuse an unchanged published blob without copying it. Publication reads the
// private link, so the public file is never modified; packageOutputs verifies
// the bytes before commit.
func linkOrCopy(source, destination string) error {
	if err := os.MkdirAll(filepath.Dir(destination), 0700); err != nil {
		return err
	}
	if err := os.Link(source, destination); err == nil {
		return nil
	} else if os.IsExist(err) {
		return err
	}
	data, err := os.ReadFile(source)
	if err != nil {
		return err
	}
	return os.WriteFile(destination, data, 0600)
}

// preparePackageOverlay builds a private package result from the last complete
// publication. It never mutates public output; any missing proof returns an
// empty reason and the caller runs the full planner.
func preparePackageOverlay(stage string, state *buildSession, entries []string, request *packageFileRequest) (*packageOverlayRecord, string, error) {
	return preparePackageOverlaySnapshot(stage, state, entries, request, nil)
}

func preparePackageOverlaySnapshot(stage string, state *buildSession, entries []string, request *packageFileRequest, snapshot *packagePlanSnapshot) (*packageOverlayRecord, string, error) {
	if snapshot == nil {
		planDir := filepath.Join(cacheDirectory(state.Spec), "go-packages", "plan-v1")
		previous, err := readPackageRequest(filepath.Join(planDir, "request.json"))
		if err != nil {
			return nil, "missing previous request", nil
		}
		result, err := readPackagePlan(filepath.Join(planDir, "result.json"))
		if err != nil {
			return nil, "missing previous result", nil
		}
		sources, err := readPackagePlanFiles(state.Spec.Output, result)
		if err != nil {
			return nil, "published plan is stale", nil
		}
		snapshot = &packagePlanSnapshot{request: previous, result: result, sources: sources}
	}
	previous, previousResult := snapshot.request, snapshot.result
	changed, reason := changedPackageSource(previous, request)
	if reason != "" {
		return nil, reason, nil
	}
	if !strings.HasPrefix(changed.Owner, "generated/") || changed.Module == "" {
		return nil, "changed source is not a generated module", nil
	}
	module := changed.Module
	overlay := readOverlayRecord(cacheDirectory(state.Spec))
	if overlay != nil {
		if overlay.Module != module || overlay.SourcePath != changed.Path || overlay.PlanHash != previousResult.PlanHash {
			return nil, "previous overlay does not cover this edit", nil
		}
	}
	overlayPackage := overlayPackageName(module)
	if overlay != nil {
		overlayPackage = overlay.Package
	}
	if overlay == nil && committedPackageFiles(stage, *state, previousResult) {
		source, readErr := os.ReadFile(changed.Source)
		if readErr == nil && digest(source) == changed.SHA256 {
			literals := packageLiteralChanges(stage, state, previous, changed, source)
			if literals != nil {
				marker := sourceMarker + hex.EncodeToString([]byte(changed.Path)) + "\n"
				count := 0
				for _, file := range snapshot.sources {
					count += strings.Count(file.source, marker)
				}
				if count > 1 {
					index, indexErr := loadOverlayIndexForSource(previousResult, snapshot.sources, overlayPackage, changed.Path)
					if indexErr != nil {
						return nil, "published bundles are not reusable", nil
					}
					if len(index.sections[changed.Path]) > 1 {
						return preparePackageSectionsProven(stage, state, entries, request, previous, previousResult, index, changed, literals)
					}
				}
			}
		}
	}
	index, err := loadOverlayIndex(previousResult, snapshot.sources, overlayPackage)
	if err != nil {
		return nil, "published bundles are not reusable", nil
	}
	if overlay == nil && len(index.sections[changed.Path]) > 1 {
		return preparePackageSections(stage, state, entries, request, previous, previousResult, index, changed)
	}
	reference, originalPackage, err := index.activeSection(changed.Path, previousResult, overlay, entries)
	if err != nil {
		return nil, err.Error(), nil
	}
	oldSource := sectionSource(reference)
	oldDecls, err := inspectModuleSource(oldSource, index.symbolLookup(reference.bundle.packageName), reference.bundle.aliases, externalAliases(reference.bundle.imports))
	if err != nil {
		return nil, "previous module source is not reusable", nil
	}
	for _, alias := range reference.bundle.imports {
		if alias == "." {
			return nil, "published package uses a dot import", nil
		}
	}
	newSourceBytes, err := os.ReadFile(changed.Source)
	if err != nil {
		return nil, "", err
	}
	newSource := string(newSourceBytes)
	if digest(newSourceBytes) != changed.SHA256 {
		return nil, "generated source changed during preparation", nil
	}
	newImports, newAliases, err := parseBundleImports(newSourceBytes, "purescript")
	if err != nil {
		return nil, "", err
	}
	for _, alias := range newImports {
		if alias == "." || alias == "_" {
			return nil, "new module source uses a dot or blank import", nil
		}
	}
	newDecls, err := inspectModuleSource(newSource, index.symbolLookup(reference.bundle.packageName), newAliases, externalAliases(newImports))
	if err != nil {
		return nil, "new module source is not parseable", nil
	}
	if reason, ok := compareDeclarations(oldDecls, newDecls); !ok {
		return nil, reason, nil
	}
	for _, declaration := range []*moduleDecls{oldDecls, newDecls} {
		switch {
		case declaration.hasType:
			return nil, "module declares a type", nil
		case declaration.hasMethod:
			return nil, "module declares a method", nil
		case declaration.hasInit:
			return nil, "module declares init", nil
		case declaration.hasVarValue:
			return nil, "module has an initialized package variable", nil
		}
	}
	for declaration, references := range oldDecls.references {
		for reference := range references {
			pkg, symbol, _ := strings.Cut(reference, "\x00")
			if pkg == originalPackage && !ast.IsExported(symbol) && !oldDecls.names[symbol] {
				return nil, "module references a private symbol of its package: " + declaration, nil
			}
			if pkg != originalPackage && !ast.IsExported(symbol) {
				return nil, "module references a private symbol of another package: " + symbol, nil
			}
		}
	}
	// A package that keeps the old bundle cannot call the overlay (that would
	// be an import cycle). Any intra-package use of the edited module is a
	// fallback, not a silently stale call site. A textual search over the
	// other sections is conservative and avoids re-parsing a 28 MB bundle.
	if overlay == nil {
		var otherSections strings.Builder
		for _, candidate := range reference.bundle.sections {
			if candidate.path == changed.Path {
				continue
			}
			otherSections.WriteString(candidate.source)
			otherSections.WriteByte('\n')
		}
		others := otherSections.String()
		wanted := map[string]bool{}
		for name := range oldDecls.names {
			if name != "init" && name != "_" {
				wanted[name] = true
			}
		}
		if anyIdentifier(others, wanted) {
			return nil, "another module in the same package calls the edited module", nil
		}
	}
	if overlay == nil && overlayPackage == originalPackage {
		return nil, "overlay package collides with its source package", nil
	}
	overlaySource, dependencies, err := c1c.BuildOverlaySource(overlayPackage, newSource, index.exported, func(pkg string) string {
		return request.PackagePath + "/" + pkg
	})
	if err != nil {
		return nil, "overlay source generation failed", nil
	}
	if err := c1c.CompileCheck(overlayPackage, overlaySource); err != nil {
		return nil, "overlay source does not parse", nil
	}
	dependencyNames := []string{}
	for _, dependency := range dependencies {
		if dependency.Package == overlayPackage {
			return nil, "overlay imports itself", nil
		}
		dependencyNames = append(dependencyNames, dependency.Package)
	}
	sort.Strings(dependencyNames)

	symbols := exportedNames(oldDecls.names)
	if len(symbols) == 0 {
		return nil, "edited module has no exported symbols to redirect", nil
	}

	packages := append([]packagePlanPackage{}, previousResult.Packages...)
	packageByName := map[string]*packagePlanPackage{}
	for i := range packages {
		packageByName[packages[i].Name] = &packages[i]
	}
	existingOverlay := packageByName[overlayPackage]
	if overlay == nil && existingOverlay != nil {
		return nil, "overlay package already exists without a record", nil
	}
	newBundlePath := overlayPackage + "/" + packageBundleName
	newOrderPath := overlayPackage + "/zz_gopurs_order.go"
	replacements := map[string]string{}
	rewritten := map[string]bool{}
	for _, file := range previousResult.Files {
		if overlay != nil {
			break
		}
		bundle := index.bundles[file.Path]
		if bundle == nil {
			continue
		}
		if bundle.packageName == overlayPackage {
			continue
		}
		oldPath := request.PackagePath + "/" + originalPackage
		alias, ok := bundle.imports[oldPath]
		if !ok {
			continue
		}
		if alias == "." {
			return nil, "dependent package uses a dot import", nil
		}
		if alias == "_" {
			continue
		}
		needs := false
		for _, symbol := range symbols {
			if strings.Contains(bundle.source, alias+"."+symbol) {
				needs = true
				break
			}
		}
		if !needs {
			continue
		}
		rewrittenSource, changedSource, err := c1c.RewritePackageReferencesAs(bundle.source, oldPath, request.PackagePath+"/"+overlayPackage,
			freeRewriteAlias(bundle.source), symbols)
		if err != nil {
			return nil, "dependent package rewrite failed", nil
		}
		if !changedSource {
			return nil, "dependent package references the edited module without a rewrite", nil
		}
		for _, symbol := range symbols {
			if strings.Contains(rewrittenSource, alias+"."+symbol) {
				return nil, "dependent package kept a stale reference", nil
			}
		}
		replacements[file.Path] = rewrittenSource
		rewritten[bundle.packageName] = true
	}
	if overlay == nil {
		declarations := 0
		for name := range newDecls.names {
			if name != "init" && name != "_" {
				declarations++
			}
		}
		if declarations == 0 {
			return nil, "overlay module has no declarations", nil
		}
		dependencySet := map[string]bool{}
		for _, name := range dependencyNames {
			if _, ok := packageByName[name]; !ok {
				return nil, "overlay dependency is not in the published plan", nil
			}
			dependencySet[name] = true
		}
		dependencies := make([]string, 0, len(dependencySet))
		for name := range dependencySet {
			dependencies = append(dependencies, name)
		}
		sort.Strings(dependencies)
		packages = append(packages, packagePlanPackage{Name: overlayPackage, Dependencies: dependencies, Declarations: declarations})
		packageByName[overlayPackage] = &packages[len(packages)-1]
	} else {
		if existingOverlay == nil {
			return nil, "overlay package disappeared from the plan", nil
		}
		if !reflect.DeepEqual(existingOverlay.Dependencies, dependencyNames) {
			return nil, "overlay package dependencies changed", nil
		}
	}
	// Appending the overlay may reallocate the package slice; refresh pointers
	// before rewriting dependent dependency sets.
	packageByName = map[string]*packagePlanPackage{}
	for i := range packages {
		packageByName[packages[i].Name] = &packages[i]
	}
	for name := range rewritten {
		pkg := packageByName[name]
		if pkg == nil {
			return nil, "rewritten package is missing from the plan", nil
		}
		found := false
		for _, dependency := range pkg.Dependencies {
			if dependency == overlayPackage {
				found = true
			}
		}
		if !found {
			pkg.Dependencies = append(pkg.Dependencies, overlayPackage)
			sort.Strings(pkg.Dependencies)
		}
	}
	ordered, err := topologicalPackages(packages)
	if err != nil {
		return nil, "overlay package graph is not reusable", nil
	}
	packages = ordered

	// Stage every published package file, replacing only proven objects.
	directory := filepath.Join(stage, "packages")
	if err := os.Mkdir(directory, 0700); err != nil {
		return nil, "", err
	}
	complete := false
	defer func() {
		if !complete {
			os.RemoveAll(directory)
		}
	}()
	files := append([]packagePlanFile{}, previousResult.Files...)
	seenBundle := false
	for i := range files {
		file := &files[i]
		if file.Path == newBundlePath {
			seenBundle = true
			continue
		}
		destination := filepath.Join(directory, filepath.FromSlash(file.Path))
		if replacement, ok := replacements[file.Path]; ok {
			data := []byte(replacement)
			file.SHA256, file.Bytes = digest(data), int64(len(data))
			if err := writePackagesFile(destination, replacement); err != nil {
				return nil, "", err
			}
			continue
		}
		if err := linkOrCopy(filepath.Join(state.Spec.Output, "purescript", filepath.FromSlash(file.Path)), destination); err != nil {
			return nil, "", err
		}
	}
	overlayBytes := []byte(overlaySource)
	if seenBundle {
		for i := range files {
			if files[i].Path == newBundlePath {
				files[i].SHA256, files[i].Bytes = digest(overlayBytes), int64(len(overlayBytes))
			}
		}
	} else {
		files = append(files, packagePlanFile{Path: newBundlePath, Kind: "package", SHA256: digest(overlayBytes), Bytes: int64(len(overlayBytes))})
	}
	if err := writePackagesFile(filepath.Join(directory, filepath.FromSlash(newBundlePath)), overlaySource); err != nil {
		return nil, "", err
	}
	orderPath := filepath.Join(directory, filepath.FromSlash(newOrderPath))
	if _, err := os.Stat(orderPath); os.IsNotExist(err) {
		var order strings.Builder
		fmt.Fprintf(&order, "package %s\n", overlayPackage)
		if len(dependencyNames) > 0 {
			order.WriteString("\nimport (\n")
			for _, name := range dependencyNames {
				fmt.Fprintf(&order, "\t_ %q\n", request.PackagePath+"/"+name)
			}
			order.WriteString(")\n")
		}
		content := order.String()
		if err := writePackagesFile(orderPath, content); err != nil {
			return nil, "", err
		}
		files = append(files, packagePlanFile{Path: newOrderPath, Kind: "order", SHA256: digest([]byte(content)), Bytes: int64(len(content))})
	}
	sort.Slice(files, func(i, j int) bool { return files[i].Path < files[j].Path })

	placements := append([]packagePlanEntry{}, previousResult.Entries...)
	unit := strings.TrimSuffix(changed.Path, ".go")
	for i := range placements {
		if placements[i].Symbol.Owner == "generated/"+unit && placements[i].Symbol.Name == "Get_"+unit+"_main" {
			placements[i].Package = request.PackagePath + "/" + overlayPackage
		}
	}
	hashParts := []string{previousResult.PlanHash, module, changed.Path, digest(overlayBytes)}
	hashParts = append(hashParts, symbols...)
	hashParts = append(hashParts, dependencyNames...)
	result := packagePlanResult{Format: packageFilesFormat, PlanHash: Key("go-package-overlay", hashParts...),
		Packages: packages, Files: files, Entries: placements}
	encoded, err := json.Marshal(result)
	if err != nil {
		return nil, "", err
	}
	if err := os.WriteFile(filepath.Join(stage, "package-result.json"), encoded, 0400); err != nil {
		return nil, "", err
	}
	if _, err := packageOutputs(stage, state.Modules, entries); err != nil {
		os.Remove(filepath.Join(stage, "package-result.json"))
		return nil, "overlay result does not satisfy the package contract", nil
	}
	record := &packageOverlayRecord{Format: packageOverlayFormat, PlanHash: result.PlanHash, Module: module,
		SourcePath: changed.Path, Package: overlayPackage, OriginalPackage: originalPackage, Symbols: symbols}
	recordData, err := json.Marshal(record)
	if err != nil {
		return nil, "", err
	}
	if err := os.WriteFile(filepath.Join(stage, "package-overlay.json"), recordData, 0400); err != nil {
		return nil, "", err
	}
	complete = true
	return record, "", nil
}
