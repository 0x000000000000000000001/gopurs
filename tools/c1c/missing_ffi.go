package c1c

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strconv"
	"strings"
)

const MissingFFIMarker = "//gopurs:missing-ffi-v1:"

type missingFFI struct {
	Binding string `json:"binding"`
	Module  string `json:"module"`
	Symbol  string `json:"symbol"`
	Path    string `json:"path"`
}

type ffiSource struct {
	file      *ast.File
	imports   map[string]string
	selectors map[*ast.Ident]bool
}

type ffiDeclaration struct {
	source     *ffiSource
	node       ast.Node
	evaluation []ast.Node
	function   bool
}

type ffiReference struct {
	name string
	// A value reached from an entry or passed as an argument can carry a
	// callback. Its complete body is needed. Eager initialization only visits
	// executed expressions, leaving stored closures alone.
	value bool
}

type ffiReachability struct {
	declarations map[string]*ffiDeclaration
	missing      map[string]missingFFI
	roots        []ffiReference
}

// CheckMissingFFI follows references in the emitted, monolithic Go source.
// Intrinsics have already been lowered, and unreferenced library bindings do
// not become roots merely because their modules were compiled. Native Go
// dynamic dispatch remains guarded by the generated runtime stubs.
func CheckMissingFFI(files map[string]string, entries []string) error {
	if len(entries) == 0 {
		return nil
	}
	hasMissing := false
	paths := make([]string, 0, len(files))
	for path, text := range files {
		paths = append(paths, path)
		hasMissing = hasMissing || strings.Contains(text, MissingFFIMarker)
	}
	if !hasMissing {
		return nil
	}
	sort.Strings(paths)
	graph := ffiReachability{
		declarations: map[string]*ffiDeclaration{},
		missing:      map[string]missingFFI{},
	}
	for _, path := range paths {
		if err := graph.addSource(path, files[path]); err != nil {
			return err
		}
	}
	for _, entry := range entries {
		if graph.declarations[entry] == nil {
			return fmt.Errorf("missing Go entry for FFI validation: %s", entry)
		}
		graph.roots = append(graph.roots, ffiReference{entry, true})
	}
	visited := map[ffiReference]bool{}
	required := map[string]missingFFI{}
	for len(graph.roots) != 0 {
		last := len(graph.roots) - 1
		ref := graph.roots[last]
		graph.roots = graph.roots[:last]
		if visited[ref] {
			continue
		}
		visited[ref] = true
		if missing, ok := graph.missing[ref.name]; ok && ref.value {
			required[ref.name] = missing
		}
		decl := graph.declarations[ref.name]
		if decl == nil {
			continue
		}
		if ref.value {
			graph.followValue(decl.source, decl.node)
		} else {
			for _, expression := range decl.evaluation {
				graph.followEvaluation(decl.source, expression)
			}
		}
	}
	diagnostics := make([]string, 0, len(required))
	for _, missing := range required {
		diagnostics = append(diagnostics, fmt.Sprintf(
			"FFI error in module %s (%s): missing Go implementation for %s. Add the Go binding in that file or the --ffi directory, or select the Go library port in workspace.extraPackages.",
			missing.Module, missing.Path, missing.Symbol))
	}
	if len(diagnostics) == 0 {
		return nil
	}
	sort.Strings(diagnostics)
	return fmt.Errorf("%s", strings.Join(diagnostics, "\n"))
}

func (graph *ffiReachability) addSource(path, text string) error {
	file, err := parser.ParseFile(token.NewFileSet(), path, text, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("Go FFI reachability: %w", err)
	}
	source := &ffiSource{file, map[string]string{}, map[*ast.Ident]bool{}}
	for _, group := range file.Comments {
		for _, comment := range group.List {
			payload, found := strings.CutPrefix(comment.Text, MissingFFIMarker)
			if !found {
				continue
			}
			var missing missingFFI
			if err := json.Unmarshal([]byte(payload), &missing); err != nil || missing.Binding == "" || missing.Module == "" || missing.Symbol == "" || missing.Path == "" {
				return fmt.Errorf("invalid missing Go FFI annotation in %s", path)
			}
			graph.missing[missing.Binding] = missing
		}
	}
	for _, spec := range file.Imports {
		path, _ := strconv.Unquote(spec.Path.Value)
		name := path[strings.LastIndex(path, "/")+1:]
		if spec.Name != nil {
			name = spec.Name.Name
		}
		source.imports[name] = path
	}
	ast.Inspect(file, func(node ast.Node) bool {
		switch value := node.(type) {
		case *ast.SelectorExpr:
			source.selectors[value.Sel] = true
		case *ast.KeyValueExpr:
			if key, ok := value.Key.(*ast.Ident); ok && key.Obj == nil {
				source.selectors[key] = true
			}
		}
		return true
	})
	for index, declaration := range file.Decls {
		switch value := declaration.(type) {
		case *ast.FuncDecl:
			if value.Recv != nil {
				continue
			}
			name := value.Name.Name
			decl := &ffiDeclaration{source: source, node: value, function: true}
			if value.Body != nil {
				decl.evaluation = []ast.Node{value.Body}
			}
			if name == "init" {
				name = fmt.Sprintf("init/%s/%d", path, index)
				graph.roots = append(graph.roots, ffiReference{name, false})
			}
			graph.declarations[name] = decl
		case *ast.GenDecl:
			for specIndex, spec := range value.Specs {
				switch item := spec.(type) {
				case *ast.ValueSpec:
					decl := &ffiDeclaration{source: source, node: item}
					for _, expression := range item.Values {
						decl.evaluation = append(decl.evaluation, expression)
					}
					for nameIndex, name := range item.Names {
						binding := name.Name
						if binding == "_" {
							binding = fmt.Sprintf("blank/%s/%d/%d/%d", path, index, specIndex, nameIndex)
						}
						graph.declarations[binding] = decl
						if value.Tok == token.VAR && len(item.Values) != 0 {
							graph.roots = append(graph.roots, ffiReference{binding, false})
						}
					}
				case *ast.TypeSpec:
					graph.declarations[item.Name.Name] = &ffiDeclaration{source: source, node: item}
				}
			}
		}
	}
	return nil
}

func (graph *ffiReachability) reference(source *ffiSource, node ast.Node) (string, *ffiDeclaration) {
	ident, ok := node.(*ast.Ident)
	if !ok || source.selectors[ident] || source.imports[ident.Name] != "" {
		return "", nil
	}
	// Parser object resolution distinguishes local bindings from package
	// bindings in this file. Unresolved identifiers may belong to another file.
	if ident.Obj != nil && ident.Obj != source.file.Scope.Lookup(ident.Name) {
		return "", nil
	}
	return ident.Name, graph.declarations[ident.Name]
}

func (graph *ffiReachability) followValue(source *ffiSource, node ast.Node) {
	ast.Inspect(node, func(node ast.Node) bool {
		if name, decl := graph.reference(source, node); decl != nil {
			graph.roots = append(graph.roots, ffiReference{name, true})
		}
		return true
	})
}

func (graph *ffiReachability) followEvaluation(source *ffiSource, node ast.Node) {
	ast.Inspect(node, func(node ast.Node) bool {
		switch value := node.(type) {
		case *ast.FuncLit:
			return false
		case *ast.CallExpr:
			if storesRuntimeValue(source, value.Fun) {
				for _, argument := range value.Args {
					graph.followEvaluation(source, argument)
				}
				return false
			}
			callee := calledExpression(value.Fun)
			if function, ok := callee.(*ast.FuncLit); ok {
				graph.followEvaluation(source, function.Body)
			} else if name, decl := graph.reference(source, callee); decl != nil {
				graph.roots = append(graph.roots, ffiReference{name, !decl.function})
			} else {
				graph.followEvaluation(source, value.Fun)
			}
			for _, argument := range value.Args {
				graph.followValue(source, argument)
			}
			return false
		}
		if name, decl := graph.reference(source, node); decl != nil && !decl.function {
			graph.roots = append(graph.roots, ffiReference{name, false})
		}
		return true
	})
}

func calledExpression(expression ast.Expr) ast.Expr {
	for {
		switch value := expression.(type) {
		case *ast.ParenExpr:
			expression = value.X
		case *ast.IndexExpr:
			expression = value.X
		case *ast.IndexListExpr:
			expression = value.X
		default:
			return expression
		}
	}
}

func storesRuntimeValue(source *ffiSource, expression ast.Expr) bool {
	selector, ok := calledExpression(expression).(*ast.SelectorExpr)
	if !ok {
		return false
	}
	base, ok := selector.X.(*ast.Ident)
	if !ok || base.Obj != nil || source.imports[base.Name] != "gopurs/output/gopurs_runtime" {
		return false
	}
	// These runtime constructors store or wrap their values without invoking
	// callbacks. Their arguments still execute before storage. Keep this list
	// tied to the implementations in runtime/runtime.go, including Func11.
	switch selector.Sel.Name {
	case "Func", "Func2", "Func3", "Func4", "Func5", "Func6", "Func7",
		"Func8", "Func9", "Func10", "Func11", "FuncAny", "WithFunctionData",
		"Wrap0", "Wrap1", "Wrap2", "Wrap3", "Wrap4", "Wrap5",
		"Any", "AnyToValue", "Box", "Array", "Record", "RecordDict",
		"RecordDict0", "RecordDict1", "RecordDict2", "RecordDict3", "RecordDict4", "RecordDict5",
		"Constructor", "Constructor0", "Constructor1", "Constructor2", "Constructor3", "Constructor4", "Constructor5":
		return true
	default:
		return false
	}
}
