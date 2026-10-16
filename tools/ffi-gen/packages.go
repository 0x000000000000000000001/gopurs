package main

import (
	"bytes"
	"container/heap"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"path"
	"sort"
	"strconv"
	"strings"
	"unicode/utf16"
)

type PackageRequest struct {
	BindRequest
	Entries []SourceSymbol `json:"entries"`
}

type PackageFile struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

// A PureScript module stem starts with a letter; a leading digit avoids even
// case-insensitive collisions during migration from per-module source files.
const packageBundleName = "0_gopurs_package.go"

type SymbolPlacement struct {
	Symbol  SourceSymbol `json:"symbol"`
	Package string       `json:"package"`
}

type GoPackage struct {
	Name         string   `json:"name"`
	Dependencies []string `json:"dependencies"`
	Declarations int      `json:"declarations"`
}

type PackageResult struct {
	Format   string            `json:"format"`
	Packages []GoPackage       `json:"packages"`
	Files    []PackageFile     `json:"files"`
	Entries  []SymbolPlacement `json:"entries"`
	PlanHash string            `json:"planHash"`
}

type packageDeclaration struct {
	unit     *sourceUnit
	index    int
	identity any
	key      string
	anchor   string
}

type declarationGraph struct {
	sources      *resolvedSources
	declarations []packageDeclaration
	byUnit       map[*sourceUnit][]int
	bySymbol     map[SourceSymbol]int
	edges        []map[int]bool
	parent       []int
	initializers []int
	recordFields map[*types.Var]string
	annotations  map[*sourceUnit][]sourceEdit
}

// Match G01's canonical JSON, including ASCII escaping of UTF-16 pairs.
func identityJSON(value any) string {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(value); err != nil {
		panic(err) // All callers supply strings, integers and lists only.
	}
	var result strings.Builder
	for _, char := range strings.TrimSuffix(buffer.String(), "\n") {
		if char < 128 {
			result.WriteRune(char)
		} else {
			for _, unit := range utf16.Encode([]rune{char}) {
				fmt.Fprintf(&result, "\\u%04x", unit)
			}
		}
	}
	return result.String()
}

func sourceHash(text string) string {
	return sourceBytesHash([]byte(text))
}

func sourceBytesHash(data []byte) string {
	value := sha256.Sum256(data)
	return hex.EncodeToString(value[:])
}

func (graph *declarationGraph) root(index int) int {
	for graph.parent[index] != index {
		graph.parent[index] = graph.parent[graph.parent[index]]
		index = graph.parent[index]
	}
	return index
}

func (graph *declarationGraph) unite(a, b int) {
	a, b = graph.root(a), graph.root(b)
	if a != b {
		if a > b {
			a, b = b, a
		}
		graph.parent[b] = a
	}
}

func (graph *declarationGraph) objectDeclaration(object types.Object) (int, error) {
	unit, index, err := graph.sources.containing(sourceOrigin(object))
	if err != nil {
		return 0, err
	}
	return graph.byUnit[unit][index], nil
}

func (graph *declarationGraph) recordField(object types.Object) string {
	if field, ok := object.(*types.Var); ok && field.IsField() {
		return graph.recordFields[field.Origin()]
	}
	return ""
}

func packageGraph(sources *resolvedSources) (*declarationGraph, error) {
	graph := &declarationGraph{sources: sources, byUnit: map[*sourceUnit][]int{}, bySymbol: map[SourceSymbol]int{}}
	var err error
	graph.recordFields, graph.annotations, err = bindRecordFields(sources)
	if err != nil {
		return nil, err
	}
	for _, unit := range sources.units {
		if path.Base(unit.input.Path) != unit.input.Path || !strings.HasSuffix(unit.input.Path, ".go") {
			return nil, fmt.Errorf("package source must have a unique Go basename: %s", unit.input.Path)
		}
		for index, bound := range unit.bound.Declarations {
			id := len(graph.declarations)
			graph.byUnit[unit] = append(graph.byUnit[unit], id)
			graph.declarations = append(graph.declarations, packageDeclaration{unit: unit, index: index})
			graph.edges = append(graph.edges, map[int]bool{})
			graph.parent = append(graph.parent, id)
			for _, symbol := range bound.Symbols {
				graph.bySymbol[symbol] = id
			}
		}
	}
	for id := range graph.declarations {
		declaration := &graph.declarations[id]
		unit, index := declaration.unit, declaration.index
		bound := unit.bound.Declarations[index]
		namespace, _, ok := strings.Cut(unit.input.Owner, "/")
		if !ok || (namespace != "generated" && namespace != "ffi") || unit.input.Module == "" {
			return nil, fmt.Errorf("package source requires generated/ffi owner and module: %s", unit.input.Path)
		}
		names := make([]string, 0, len(bound.Symbols))
		for _, symbol := range bound.Symbols {
			names = append(names, symbol.Name)
		}
		sort.Strings(names)
		if bound.Kind != "method" && bound.Kind != "init" {
			declaration.identity = []any{"gopurs/symbol-v1", namespace, unit.input.Module, bound.Kind, names}
		}
		if namespace == "ffi" {
			graph.unite(id, graph.byUnit[unit][0])
		}
		for _, name := range names {
			if bound.Kind == "function" && strings.HasPrefix(name, "Get_") {
				suffix := strings.TrimPrefix(name, "Get_")
				cache, hasCache := graph.bySymbol[SourceSymbol{unit.input.Owner, "cache_" + suffix}]
				once, hasOnce := graph.bySymbol[SourceSymbol{unit.input.Owner, "once_" + suffix}]
				if hasCache && hasOnce {
					graph.unite(id, cache)
					graph.unite(id, once)
					if worker, ok := graph.bySymbol[SourceSymbol{unit.input.Owner, "Call_" + suffix}]; ok {
						graph.unite(id, worker)
					}
				}
			}
		}
		var failure error
		ast.Inspect(unit.decls[index], func(node ast.Node) bool {
			ident, ok := node.(*ast.Ident)
			if !ok || failure != nil {
				return failure == nil
			}
			object := sources.info.Uses[ident]
			if object == nil || object.Pkg() != sources.pkg {
				return true
			}
			object = sourceOrigin(object)
			_, global := sources.owned[object]
			field, isField := object.(*types.Var)
			method, isMethod := object.(*types.Func)
			member := isField && field.IsField() || isMethod && receiverObject(method) != nil
			if !global && !member {
				return true
			}
			var provider int
			provider, failure = graph.objectDeclaration(object)
			if failure != nil {
				return false
			}
			graph.edges[id][provider] = true
			if !object.Exported() && graph.recordField(object) == "" {
				graph.unite(id, provider)
			}
			return true
		})
		if failure != nil {
			return nil, failure
		}
	}
	for id := range graph.declarations {
		declaration := &graph.declarations[id]
		unit, index := declaration.unit, declaration.index
		bound := unit.bound.Declarations[index]
		namespace, _, _ := strings.Cut(unit.input.Owner, "/")
		prefix := []any{"gopurs/symbol-v1", namespace, unit.input.Module}
		switch bound.Kind {
		case "method":
			method := unit.decls[index].(*ast.FuncDecl)
			object := sources.info.Defs[method.Name].(*types.Func)
			receiver, err := graph.objectDeclaration(receiverObject(object))
			if err != nil {
				return nil, err
			}
			graph.unite(id, receiver)
			declaration.identity = append(prefix, "method", graph.declarations[receiver].identity, []string{method.Name.Name})
		case "init":
			graph.initializers = append(graph.initializers, id)
			name := bound.Symbols[0].Name
			if namespace == "ffi" {
				declaration.identity = append(prefix, "init-slot", []string{"init#" + strings.TrimPrefix(name, "init/")})
				declaration.anchor = sourceHash(identityJSON([]string{"gopurs/ffi-init-atom-v1", unit.input.Module}))
			} else if constructor, ok := strings.CutPrefix(name, "register/"); ok {
				target, found := graph.bySymbol[SourceSymbol{unit.input.Owner, constructor}]
				if !found || unit.bound.Declarations[index].Symbols[0].Owner != unit.input.Owner || graph.declarations[target].identity == nil {
					return nil, fmt.Errorf("missing initializer constructor: %+v", bound.Symbols[0])
				}
				declaration.identity = append(prefix, "struct-getter", graph.declarations[target].identity)
			} else {
				declaration.identity = append(prefix, "initializer", name)
			}
		}
		declaration.key = sourceHash(identityJSON(declaration.identity))
		if declaration.anchor == "" {
			declaration.anchor = declaration.key
		}
	}
	variables := []int{}
	for _, initializer := range sources.info.InitOrder {
		previous := -1
		for _, variable := range initializer.Lhs {
			id, err := graph.objectDeclaration(variable)
			if err != nil {
				return nil, err
			}
			if previous >= 0 && previous != id {
				return nil, fmt.Errorf("initializer spans multiple declarations")
			}
			previous = id
		}
		if previous >= 0 {
			variables = append(variables, previous)
		}
	}
	graph.initializers = append(variables, graph.initializers...)
	for i := 1; i < len(graph.initializers); i++ {
		graph.edges[graph.initializers[i]][graph.initializers[i-1]] = true
	}
	return graph, nil
}

// Iterative Kosaraju: recursive Go/PureScript stacks are not a corpus limit.
func stronglyConnected(edges [][]int) [][]int {
	reverse := make([][]int, len(edges))
	for caller, providers := range edges {
		for _, provider := range providers {
			reverse[provider] = append(reverse[provider], caller)
		}
	}
	type frame struct{ node, next int }
	seen, order := make([]bool, len(edges)), []int{}
	for root := range edges {
		if seen[root] {
			continue
		}
		seen[root] = true
		stack := []frame{{node: root}}
		for len(stack) > 0 {
			current := &stack[len(stack)-1]
			if current.next == len(edges[current.node]) {
				order = append(order, current.node)
				stack = stack[:len(stack)-1]
				continue
			}
			child := edges[current.node][current.next]
			current.next++
			if !seen[child] {
				seen[child] = true
				stack = append(stack, frame{node: child})
			}
		}
	}
	seen = make([]bool, len(edges))
	components := [][]int{}
	for i := len(order) - 1; i >= 0; i-- {
		root := order[i]
		if seen[root] {
			continue
		}
		seen[root] = true
		stack, component := []int{root}, []int{}
		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			component = append(component, node)
			for _, caller := range reverse[node] {
				if !seen[caller] {
					seen[caller] = true
					stack = append(stack, caller)
				}
			}
		}
		components = append(components, component)
	}
	return components
}

type anchorQueue struct {
	keys  []string
	nodes []int
}

func (q anchorQueue) Len() int           { return len(q.nodes) }
func (q anchorQueue) Less(i, j int) bool { return q.keys[q.nodes[i]] < q.keys[q.nodes[j]] }
func (q anchorQueue) Swap(i, j int)      { q.nodes[i], q.nodes[j] = q.nodes[j], q.nodes[i] }
func (q *anchorQueue) Push(x any)        { q.nodes = append(q.nodes, x.(int)) }
func (q *anchorQueue) Pop() any {
	x := q.nodes[len(q.nodes)-1]
	q.nodes = q.nodes[:len(q.nodes)-1]
	return x
}

type packagePlan struct {
	packages  []GoPackage
	placement []int
	hash      string
}

func planPackages(graph *declarationGraph, density uint64) (*packagePlan, error) {
	if density == 0 {
		return nil, fmt.Errorf("zero package anchor density")
	}
	n := len(graph.declarations)
	edges := make([][]int, n)
	// Bidirectional ownership edges let the same SCC implementation absorb unions.
	for caller, providers := range graph.edges {
		root := graph.root(caller)
		if root != caller {
			edges[caller] = append(edges[caller], root)
			edges[root] = append(edges[root], caller)
		}
		for provider := range providers {
			edges[caller] = append(edges[caller], provider)
		}
	}
	components := stronglyConnected(edges)
	componentOf, keys := make([]int, n), make([]string, len(components))
	seenKeys := map[string]bool{}
	for id, component := range components {
		anchors := map[string]bool{}
		for _, node := range component {
			componentOf[node] = id
			anchors[graph.declarations[node].anchor] = true
		}
		keys[id] = sourceHash(strings.Join(sortedKeys(anchors), "\n") + "\n")
		if seenKeys[keys[id]] {
			return nil, fmt.Errorf("component identity collision: %s", keys[id])
		}
		seenKeys[keys[id]] = true
	}
	dag := make([]map[int]bool, len(components))
	for i := range dag {
		dag[i] = map[int]bool{}
	}
	for caller, providers := range edges {
		for _, provider := range providers {
			a, b := componentOf[caller], componentOf[provider]
			if a != b {
				dag[a][b] = true
			}
		}
	}
	callers, pending := make([][]int, len(dag)), make([]int, len(dag))
	ready := &anchorQueue{keys: keys}
	for caller, providers := range dag {
		pending[caller] = len(providers)
		if len(providers) == 0 {
			heap.Push(ready, caller)
		}
		for provider := range providers {
			callers[provider] = append(callers[provider], caller)
		}
	}
	plan := &packagePlan{placement: make([]int, n)}
	packageNames := map[string]bool{}
	visited := 0
	for ready.Len() > 0 {
		id := heap.Pop(ready).(int)
		visited++
		anchor, _ := hex.DecodeString(keys[id][:16])
		boundary := binary.BigEndian.Uint64(anchor)%density == 0
		if len(plan.packages) == 0 || boundary {
			name := "p_root"
			if boundary {
				name = "p_" + keys[id][:32]
			}
			if packageNames[name] {
				return nil, fmt.Errorf("package name collision: %s", name)
			}
			packageNames[name] = true
			plan.packages = append(plan.packages, GoPackage{Name: name, Dependencies: []string{}})
		}
		p := len(plan.packages) - 1
		for _, node := range components[id] {
			plan.placement[node] = p
			plan.packages[p].Declarations++
		}
		for _, caller := range callers[id] {
			pending[caller]--
			if pending[caller] == 0 {
				heap.Push(ready, caller)
			}
		}
	}
	if visited != len(components) {
		return nil, fmt.Errorf("cyclic condensed package graph")
	}
	dependencies := make([]map[string]bool, len(plan.packages))
	for i := range dependencies {
		dependencies[i] = map[string]bool{}
	}
	assignments := make([][2]string, n)
	for id, declaration := range graph.declarations {
		p := plan.placement[id]
		assignments[id] = [2]string{declaration.key, plan.packages[p].Name}
		for provider := range graph.edges[id] {
			q := plan.placement[provider]
			if q != p {
				if q > p {
					return nil, fmt.Errorf("non-topological package edge: %d -> %d", p, q)
				}
				dependencies[p][plan.packages[q].Name] = true
			}
		}
	}
	for p := range plan.packages {
		plan.packages[p].Dependencies = sortedKeys(dependencies[p])
	}
	sort.Slice(assignments, func(i, j int) bool { return assignments[i][0] < assignments[j][0] })
	plan.hash = sourceHash(identityJSON(assignments))
	return plan, nil
}

func sortedKeys[V any](values map[string]V) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// The production density is fixed by G01, not chosen from the current body sizes.
func PackageSources(request PackageRequest) (*PackageResult, error) {
	return packageSources(request, 8192)
}

func packageSources(request PackageRequest, density uint64) (*PackageResult, error) {
	sources, err := resolveSources(request.BindRequest)
	if err != nil {
		return nil, err
	}
	graph, err := packageGraph(sources)
	if err != nil {
		return nil, err
	}
	plan, err := planPackages(graph, density)
	if err != nil {
		return nil, err
	}
	return renderPackages(request, graph, plan)
}

func Packages(content string) (string, error) {
	var request PackageRequest
	if err := json.Unmarshal([]byte(content), &request); err != nil {
		return "", fmt.Errorf("decode Go package request: %w", err)
	}
	result, err := PackageSources(request)
	if err != nil {
		return "", err
	}
	encoded, err := json.Marshal(result)
	return string(encoded), err
}

// An alias depends on its path and the complete reserved-name set. All hash
// collisions extend together; iteration order cannot select a shorter winner.
func packageAliases(paths []string, reserved map[string]bool) (map[string]string, error) {
	digests := map[string]string{}
	seen := map[string]string{}
	for _, p := range paths {
		digest := sourceHash(p)
		if previous, ok := seen[digest]; ok && previous != p {
			return nil, fmt.Errorf("import hash collision: %s and %s", previous, p)
		}
		digests[p], seen[digest] = digest, p
	}
	result := map[string]string{}
	for _, p := range paths {
		digest, width := digests[p], 6
		for {
			alias := "g" + digest[:width]
			collision := reserved[alias]
			for other, hash := range digests {
				if other != p && hash[:width] == digest[:width] {
					collision = true
				}
			}
			if !collision {
				result[p] = alias
				break
			}
			if width < 64 {
				width++
			} else {
				for reserved[alias] {
					alias += "_"
				}
				result[p] = alias
				break
			}
		}
	}
	return result, nil
}

type sourceEdit struct {
	start, end int
	text       string
	qualifier  string
}

func writeSourceEdits(buffer *strings.Builder, text string, start, end int, edits []sourceEdit, aliases map[string]string) error {
	sort.Slice(edits, func(i, j int) bool { return edits[i].start < edits[j].start })
	for _, edit := range edits {
		if edit.start < start || edit.end <= edit.start || edit.end > end {
			return fmt.Errorf("overlapping or out-of-range Go package edit")
		}
		buffer.WriteString(text[start:edit.start])
		if edit.qualifier != "" {
			buffer.WriteString(aliases[edit.qualifier])
		}
		buffer.WriteString(edit.text)
		start = edit.end
	}
	buffer.WriteString(text[start:end])
	return nil
}

func renderPackages(request PackageRequest, graph *declarationGraph, plan *packagePlan) (*PackageResult, error) {
	// File-level build/line directives retain their original file boundaries.
	// Ordinary generated/FFI sources share one file per package: this does not
	// change the dependency partition, variable initialization or init order.
	coalesced := true
	for _, unit := range graph.sources.units {
		prefix := unit.input.Content[:unit.token.Offset(unit.file.Package)]
		if strings.Contains(prefix, "//go:") || strings.Contains(prefix, "// +build") || strings.Contains(prefix, "//line ") {
			coalesced = false
		}
		for _, group := range unit.file.Comments {
			for _, comment := range group.List {
				if strings.HasPrefix(comment.Text, "// gopurs-source-") || strings.HasPrefix(comment.Text, "//line ") || strings.HasPrefix(comment.Text, "/*line ") {
					coalesced = false
				}
			}
		}
	}
	type sourceFragment struct {
		unit *sourceUnit
		ids  []int
	}
	byPackage := make([][]sourceFragment, len(plan.packages))
	imports := map[string]bool{}
	for _, unit := range graph.sources.units {
		byPlacement := map[int][]int{}
		for _, id := range graph.byUnit[unit] {
			p := plan.placement[id]
			byPlacement[p] = append(byPlacement[p], id)
		}
		for p, ids := range byPlacement {
			byPackage[p] = append(byPackage[p], sourceFragment{unit, ids})
		}
		for _, spec := range unit.file.Imports {
			target, _ := strconv.Unquote(spec.Path.Value)
			if target == "C" {
				return nil, fmt.Errorf("cgo must remain in a native package: %s", unit.input.Path)
			}
			imports[target] = true
		}
	}
	result := &PackageResult{Format: "gopurs/go-packages-v1", Packages: plan.packages, Files: []PackageFile{}, Entries: []SymbolPlacement{}, PlanHash: plan.hash}
	sources := graph.sources
	type fileGroup struct {
		placement int
		fragments []sourceFragment
		name      string
	}
	groups := []fileGroup{}
	packageReserved := make([]map[string]bool, len(plan.packages))
	for p, fragments := range byPackage {
		sort.Slice(fragments, func(i, j int) bool { return fragments[i].unit.input.Path < fragments[j].unit.input.Path })
		packageReserved[p] = map[string]bool{}
		for _, fragment := range fragments {
			for _, id := range fragment.ids {
				for _, symbol := range fragment.unit.bound.Declarations[graph.declarations[id].index].Symbols {
					packageReserved[p][symbol.Name] = true
				}
			}
		}
		if coalesced {
			groups = append(groups, fileGroup{p, fragments, packageBundleName})
		} else {
			for _, fragment := range fragments {
				groups = append(groups, fileGroup{p, []sourceFragment{fragment}, fragment.unit.input.Path})
			}
		}
	}
	// Typed edits are collected while discovering imports and reserved names.
	// Their import aliases become known only after that complete first pass.
	editsByDeclaration := make([][]sourceEdit, len(graph.declarations))
	for _, group := range groups {
		p, fragments := group.placement, group.fragments
		reserved, needed, blank := map[string]bool{}, map[string]bool{}, map[string]bool{}
		for name := range packageReserved[p] {
			reserved[name] = true
		}
		preferred := map[string]map[string]bool{}
		selectors := map[*ast.Ident]bool{}
		var failure error
		for _, fragment := range fragments {
			unit := fragment.unit
			dotted := map[string]bool{}
			for _, spec := range unit.file.Imports {
				target, _ := strconv.Unquote(spec.Path.Value)
				if spec.Name != nil {
					if spec.Name.Name == "." {
						dotted[target] = true
					}
					if spec.Name.Name == "_" {
						blank[target] = true
					}
				}
			}
			for _, id := range fragment.ids {
				ast.Inspect(unit.decls[graph.declarations[id].index], func(node ast.Node) bool {
					if selector, ok := node.(*ast.SelectorExpr); ok && len(dotted) > 0 {
						selectors[selector.Sel] = true
					}
					ident, ok := node.(*ast.Ident)
					if !ok || failure != nil {
						return failure == nil
					}
					object := sources.info.Uses[ident]
					fieldObject := object
					if fieldObject == nil {
						fieldObject = sources.info.Defs[ident]
					}
					fieldName := graph.recordField(fieldObject)
					if fieldName != "" {
						editsByDeclaration[id] = append(editsByDeclaration[id], sourceEdit{
							start: unit.token.Offset(ident.Pos()), end: unit.token.Offset(ident.End()), text: fieldName,
						})
					}
					object = sourceOrigin(object)
					if imported, ok := object.(*types.PkgName); ok {
						target := imported.Imported().Path()
						needed[target] = true
						if preferred[target] == nil {
							preferred[target] = map[string]bool{}
						}
						preferred[target][imported.Name()] = true
						if fieldName == "" {
							editsByDeclaration[id] = append(editsByDeclaration[id], sourceEdit{
								start: unit.token.Offset(ident.Pos()), end: unit.token.Offset(ident.End()), qualifier: target,
							})
						}
						return true
					}
					reserved[ident.Name] = true
					if object == nil {
						return true
					}
					target := ""
					if !selectors[ident] && object.Pkg() != nil && dotted[object.Pkg().Path()] && object.Parent() == object.Pkg().Scope() {
						target = object.Pkg().Path()
						needed[target] = true
					} else if symbol, ok := sources.owned[object]; ok {
						if method, ok := object.(*types.Func); ok && receiverObject(method) != nil {
							return true
						}
						provider := plan.placement[graph.bySymbol[symbol]]
						if provider != p {
							if !object.Exported() {
								failure = fmt.Errorf("private symbol crosses packages: %+v", symbol)
								return false
							}
							target = request.PackagePath + "/" + plan.packages[provider].Name
							needed[target] = true
						}
					}
					if fieldName == "" && target != "" {
						editsByDeclaration[id] = append(editsByDeclaration[id], sourceEdit{
							start: unit.token.Offset(ident.Pos()), end: unit.token.Offset(ident.End()),
							text: "." + ident.Name, qualifier: target,
						})
					}
					return true
				})
			}
		}
		if failure != nil {
			return nil, failure
		}
		aliases, counts := map[string]string{}, map[string]int{}
		for _, names := range preferred {
			for name := range names {
				counts[name]++
			}
		}
		const runtimePath = "gopurs/output/gopurs_runtime"
		if needed[runtimePath] {
			alias := "r"
			for reserved[alias] {
				alias += "_"
			}
			aliases[runtimePath], reserved[alias] = alias, true
		}
		// Keep an existing external alias only when every use agrees, and no
		// declaration/local name or other import can shadow it after coalescing.
		for target, names := range preferred {
			if target == runtimePath || len(names) != 1 {
				continue
			}
			for name := range names {
				if !reserved[name] && counts[name] == 1 {
					aliases[target], reserved[name] = name, true
				}
			}
		}
		paths := []string{}
		for target := range needed {
			if aliases[target] == "" {
				paths = append(paths, target)
			}
		}
		generated, err := packageAliases(paths, reserved)
		if err != nil {
			return nil, err
		}
		for target, alias := range generated {
			aliases[target] = alias
		}
		for target := range blank {
			if !needed[target] {
				aliases[target] = "_"
			}
		}
		var content strings.Builder
		// Most bytes are retained source text. Reserve their size once, then
		// splice edits directly into the final file instead of copying every
		// declaration through another growing buffer.
		sourceBytes := 0
		for _, fragment := range fragments {
			for _, id := range fragment.ids {
				decl := fragment.unit.bound.Declarations[graph.declarations[id].index]
				sourceBytes += decl.End - decl.Start + 1
			}
		}
		content.Grow(sourceBytes)
		if !coalesced {
			unit := fragments[0].unit
			content.WriteString(unit.input.Content[:unit.token.Offset(unit.file.Package)])
		}
		fmt.Fprintf(&content, "package %s\n", plan.packages[p].Name)
		writePackageImports(&content, aliases)
		for _, fragment := range fragments {
			unit := fragment.unit
			if coalesced {
				fmt.Fprintf(&content, "\n// gopurs-source-v1: %x\n", []byte(unit.input.Path))
				content.WriteString(unit.input.Content[:unit.token.Offset(unit.file.Package)])
			}
			for _, id := range fragment.ids {
				index := graph.declarations[id].index
				decl := unit.decls[index]
				start := unit.token.Offset(unit.file.Name.End())
				if index > 0 {
					start = unit.token.Offset(unit.decls[index-1].End())
				} else {
					for _, spec := range unit.file.Decls {
						if imported, ok := spec.(*ast.GenDecl); ok && imported.Tok == token.IMPORT {
							start = unit.token.Offset(imported.End())
						}
					}
				}
				edits := editsByDeclaration[id]
				editsByDeclaration[id] = nil
				annotations := graph.annotations[unit]
				first := sort.Search(len(annotations), func(i int) bool { return annotations[i].start >= start })
				for _, annotation := range annotations[first:] {
					if annotation.end > unit.token.Offset(decl.End()) {
						break
					}
					edits = append(edits, annotation)
				}
				if err := writeSourceEdits(&content, unit.input.Content, start, unit.token.Offset(decl.End()), edits, aliases); err != nil {
					return nil, err
				}
				content.WriteByte('\n')
			}
			if coalesced {
				content.WriteString("// gopurs-source-end-v1\n")
			}
		}
		result.Files = append(result.Files, PackageFile{plan.packages[p].Name + "/" + group.name, content.String()})
	}
	return completePackageRendering(request, graph, plan, result, imports)
}

func completePackageRendering(request PackageRequest, graph *declarationGraph, plan *packagePlan, result *PackageResult, imports map[string]bool) (*PackageResult, error) {
	// All external packages initialized before the former monolithic package.
	// Retain that union even for an entry referencing only a small generated leaf.
	externalImports := map[string]string{}
	for p := range imports {
		externalImports[p] = "_"
	}
	var external strings.Builder
	external.WriteString("package init_external\n")
	writePackageImports(&external, externalImports)
	result.Files = append(result.Files, PackageFile{"init_external/init.go", external.String()})
	effects := map[int]bool{}
	for _, id := range graph.initializers {
		effects[plan.placement[id]] = true
	}
	for p, pkg := range plan.packages {
		dependencies := map[string]string{}
		for _, name := range pkg.Dependencies {
			dependencies[request.PackagePath+"/"+name] = "_"
		}
		if effects[p] {
			dependencies[request.PackagePath+"/init_external"] = "_"
		}
		var content strings.Builder
		fmt.Fprintf(&content, "package %s\n", pkg.Name)
		writePackageImports(&content, dependencies)
		result.Files = append(result.Files, PackageFile{pkg.Name + "/zz_gopurs_order.go", content.String()})
	}
	retained := map[string]string{request.PackagePath + "/init_external": "_"}
	if n := len(graph.initializers); n > 0 {
		retained[request.PackagePath+"/"+plan.packages[plan.placement[graph.initializers[n-1]]].Name] = "_"
	}
	var initialization strings.Builder
	initialization.WriteString("package init_all\n")
	writePackageImports(&initialization, retained)
	result.Files = append(result.Files, PackageFile{"init_all/init.go", initialization.String()})
	for _, entry := range request.Entries {
		id, ok := graph.bySymbol[entry]
		if !ok || !ast.IsExported(entry.Name) {
			return nil, fmt.Errorf("missing or private Go entry: %+v", entry)
		}
		result.Entries = append(result.Entries, SymbolPlacement{entry, request.PackagePath + "/" + plan.packages[plan.placement[id]].Name})
	}
	sort.Slice(result.Files, func(i, j int) bool { return result.Files[i].Path < result.Files[j].Path })
	for i := 1; i < len(result.Files); i++ {
		if result.Files[i-1].Path == result.Files[i].Path {
			return nil, fmt.Errorf("generated package file collision: %s", result.Files[i].Path)
		}
	}
	return result, nil
}

// This is the field ABI, not a capitalization heuristic. In particular, labels
// "type", "go__type", "a-b" and "a_minus_b" have distinct exported spellings.
func exportedRecordField(encoded string) (string, error) {
	if len(encoded)%4 != 0 {
		return "", fmt.Errorf("invalid UTF-16 record label: %q", encoded)
	}
	units := make([]uint16, 0, len(encoded)/4)
	ascii := encoded != ""
	for i := 0; i < len(encoded); i += 4 {
		value, err := strconv.ParseUint(encoded[i:i+4], 16, 16)
		if err != nil || strings.ToLower(encoded[i:i+4]) != encoded[i:i+4] {
			return "", fmt.Errorf("invalid UTF-16 record label: %q", encoded)
		}
		units = append(units, uint16(value))
		letter := value == '_' || value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z'
		ascii = ascii && (letter || i > 0 && value >= '0' && value <= '9')
	}
	if ascii {
		return "P_" + string(utf16.Decode(units)), nil
	}
	return "PX_" + encoded, nil
}

func bindRecordFields(sources *resolvedSources) (map[*types.Var]string, map[*sourceUnit][]sourceEdit, error) {
	// A tuple, ADT payload and record can have identical anonymous Go shapes.
	// Only explicit producer metadata owns a field rename, including in FFI.
	fields := map[*types.Var]string{}
	annotations := map[*sourceUnit][]sourceEdit{}
	for _, unit := range sources.units {
		declared := map[string]SourceRecord{}
		for _, layout := range unit.input.Records {
			if previous, exists := declared[layout.Type]; exists && identityJSON(previous.Labels) != identityJSON(layout.Labels) {
				return nil, nil, fmt.Errorf("ambiguous record labels in %s: %s", unit.input.Path, layout.Type)
			}
			declared[layout.Type] = layout
		}
		used := map[string]bool{}
		comments := []*ast.Comment{}
		markers := map[*ast.Comment]bool{}
		for _, group := range unit.file.Comments {
			for _, comment := range group.List {
				comments = append(comments, comment)
				if strings.HasPrefix(comment.Text, "/*gopurs:record-") {
					markers[comment] = false
				}
				if strings.HasPrefix(comment.Text, "//gopurs:init-v1:") {
					annotations[unit] = append(annotations[unit], sourceEdit{start: unit.token.Offset(comment.Pos()), end: unit.token.Offset(comment.End())})
				}
			}
		}
		var failure error
		ast.Inspect(unit.file, func(node ast.Node) bool {
			structure, ok := node.(*ast.StructType)
			if !ok || failure != nil {
				return failure == nil
			}
			spelling := unit.input.Content[unit.token.Offset(structure.Pos()):unit.token.Offset(structure.End())]
			layout, ok := declared[spelling]
			last := sort.Search(len(comments), func(i int) bool { return comments[i].End() >= structure.Fields.Closing }) - 1
			if last >= 0 {
				comment := comments[last]
				if _, marker := markers[comment]; marker && comment.Pos() > structure.Fields.Opening && strings.TrimSpace(unit.input.Content[unit.token.Offset(comment.End()):unit.token.Offset(structure.Fields.Closing)]) == "" {
					labels, err := recordAnnotation(comment.Text)
					if err != nil {
						failure = fmt.Errorf("%s: %w", unit.input.Path, err)
						return false
					}
					if ok && identityJSON(labels) != identityJSON(layout.Labels) {
						failure = fmt.Errorf("contradictory record annotation in %s", unit.input.Path)
						return false
					}
					layout, ok = SourceRecord{spelling, labels}, true
					markers[comment] = true
					annotations[unit] = append(annotations[unit], sourceEdit{start: unit.token.Offset(comment.Pos()), end: unit.token.Offset(comment.End())})
				}
			}
			if !ok {
				return true
			}
			used[spelling] = true
			objects := []*types.Var{}
			for _, field := range structure.Fields.List {
				if len(field.Names) == 0 {
					failure = fmt.Errorf("record metadata names an embedded Go field in %s", unit.input.Path)
					return false
				}
				for _, name := range field.Names {
					objects = append(objects, sources.info.Defs[name].(*types.Var).Origin())
				}
			}
			if len(objects) != len(layout.Labels) {
				failure = fmt.Errorf("record label count mismatch in %s: %s", unit.input.Path, spelling)
				return false
			}
			names := []string{}
			for i, label := range layout.Labels {
				name, err := exportedRecordField(label)
				if err != nil {
					failure = err
					return false
				}
				for _, previous := range names {
					if previous == name {
						failure = fmt.Errorf("duplicate record label in %s: %s", unit.input.Path, label)
						return false
					}
				}
				names = append(names, name)
				if previous, ok := fields[objects[i]]; ok && previous != name {
					failure = fmt.Errorf("ambiguous record field identity in %s: %s", unit.input.Path, objects[i].Name())
					return false
				}
				fields[objects[i]] = name
			}
			return true
		})
		if failure != nil {
			return nil, nil, failure
		}
		for spelling := range declared {
			if !used[spelling] {
				return nil, nil, fmt.Errorf("unused record metadata in %s: %s", unit.input.Path, spelling)
			}
		}
		for marker, used := range markers {
			if !used {
				return nil, nil, fmt.Errorf("record annotation outside its struct in %s at %d", unit.input.Path, unit.token.Offset(marker.Pos()))
			}
		}
		sort.Slice(annotations[unit], func(i, j int) bool { return annotations[unit][i].start < annotations[unit][j].start })
	}
	return fields, annotations, nil
}

func recordAnnotation(comment string) ([]string, error) {
	const prefix = "/*gopurs:record-v1:"
	if !strings.HasPrefix(comment, prefix) || !strings.HasSuffix(comment, "*/") {
		return nil, fmt.Errorf("unsupported record annotation: %s", comment)
	}
	count, encoded, ok := strings.Cut(strings.TrimSuffix(strings.TrimPrefix(comment, prefix), "*/"), ":")
	n, err := strconv.Atoi(count)
	if !ok || err != nil || n < 0 || strconv.Itoa(n) != count {
		return nil, fmt.Errorf("invalid record annotation count: %s", comment)
	}
	labels := strings.Split(encoded, ",")
	if n == 0 && encoded == "" {
		labels = []string{}
	}
	if len(labels) != n {
		return nil, fmt.Errorf("invalid record annotation labels: %s", comment)
	}
	return labels, nil
}

func writePackageImports(buffer *strings.Builder, imports map[string]string) {
	if len(imports) == 0 {
		return
	}
	buffer.WriteString("\nimport (\n")
	for _, p := range sortedKeys(imports) {
		fmt.Fprintf(buffer, "\t%s %s\n", imports[p], strconv.Quote(p))
	}
	buffer.WriteString(")\n")
}
