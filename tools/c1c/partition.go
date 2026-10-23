package c1c

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strconv"
	"strings"
)

// PartitionSource renders only the declarations assigned to packageName using
// the imports already owned by that package. It cannot add imports or move a
// declaration. The caller must verify the resulting signatures and references
// against the published section before reusing the package plan.
func PartitionSource(source, packageName string, placements, symbols, imports map[string]string, packagePath func(string) string) (string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "partition.go", source, parser.ParseComments)
	if err != nil {
		return "", err
	}
	inputImports := map[string]string{}
	for _, spec := range file.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			return "", err
		}
		alias := path[strings.LastIndex(path, "/")+1:]
		if spec.Name != nil {
			alias = spec.Name.Name
		}
		if alias == "." || alias == "_" {
			return "", fmt.Errorf("unsupported source import: %s", alias)
		}
		inputImports[alias] = path
	}
	declarationNames := map[*ast.Ident]bool{}
	selected := []ast.Decl{}
	for _, declaration := range file.Decls {
		names := []*ast.Ident{}
		switch value := declaration.(type) {
		case *ast.FuncDecl:
			if value.Recv != nil || value.Name.Name == "init" {
				return "", fmt.Errorf("method or init declaration")
			}
			names = append(names, value.Name)
		case *ast.GenDecl:
			if value.Tok == token.IMPORT {
				continue
			}
			if value.Tok == token.TYPE {
				return "", fmt.Errorf("type declaration")
			}
			for _, item := range value.Specs {
				spec, ok := item.(*ast.ValueSpec)
				if !ok || value.Tok == token.VAR && len(spec.Values) != 0 {
					return "", fmt.Errorf("initialized package variable")
				}
				names = append(names, spec.Names...)
			}
		}
		owner := ""
		for _, name := range names {
			declarationNames[name] = true
			pkg, ok := placements[name.Name]
			if !ok || pkg == "" || owner != "" && owner != pkg {
				return "", fmt.Errorf("missing or ambiguous placement for %s", name.Name)
			}
			owner = pkg
		}
		if owner == packageName {
			selected = append(selected, declaration)
		}
	}
	selectors := map[*ast.Ident]bool{}
	locals := map[string]bool{}
	ambiguousKeys := map[*ast.Ident]bool{}
	for _, declaration := range selected {
		ast.Inspect(declaration, func(node ast.Node) bool {
			switch value := node.(type) {
			case *ast.SelectorExpr:
				selectors[value.Sel] = true
			case *ast.KeyValueExpr:
				if name, ok := value.Key.(*ast.Ident); ok {
					ambiguousKeys[name] = true
				}
			case *ast.Ident:
				if value.Obj != nil && value.Obj != file.Scope.Lookup(value.Name) {
					locals[value.Name] = true
				}
			}
			return true
		})
	}
	aliasFor := func(path string) (string, error) {
		alias, ok := imports[path]
		if !ok || alias == "." || alias == "_" || locals[alias] {
			return "", fmt.Errorf("missing or shadowed published import: %s", path)
		}
		return alias, nil
	}
	type edit struct {
		start, end int
		text       string
	}
	var out strings.Builder
	for _, declaration := range selected {
		start, end := fset.Position(declaration.Pos()).Offset, fset.Position(declaration.End()).Offset
		edits := []edit{}
		var rewriteErr error
		ast.Inspect(declaration, func(node ast.Node) bool {
			ident, ok := node.(*ast.Ident)
			if !ok || selectors[ident] || declarationNames[ident] || rewriteErr != nil {
				return true
			}
			if ident.Obj != nil && ident.Obj != file.Scope.Lookup(ident.Name) {
				return true
			}
			if ambiguousKeys[ident] {
				if _, exists := symbols[ident.Name]; exists {
					rewriteErr = fmt.Errorf("ambiguous composite literal key: %s", ident.Name)
				}
				return true
			}
			replacement := ""
			if path, ok := inputImports[ident.Name]; ok && ident.Obj == nil {
				replacement, rewriteErr = aliasFor(path)
			} else if pkg, ok := symbols[ident.Name]; ok && pkg != packageName {
				if !ast.IsExported(ident.Name) {
					rewriteErr = fmt.Errorf("private cross-package reference: %s", ident.Name)
					return false
				}
				var alias string
				alias, rewriteErr = aliasFor(packagePath(pkg))
				replacement = alias + "." + ident.Name
			}
			if replacement != "" && replacement != ident.Name {
				edits = append(edits, edit{fset.Position(ident.Pos()).Offset - start, fset.Position(ident.End()).Offset - start, replacement})
			}
			return true
		})
		if rewriteErr != nil {
			return "", rewriteErr
		}
		sort.Slice(edits, func(i, j int) bool { return edits[i].start > edits[j].start })
		body := source[start:end]
		for _, edit := range edits {
			body = body[:edit.start] + edit.text + body[edit.end:]
		}
		out.WriteString("\n\n")
		out.WriteString(body)
		out.WriteByte('\n')
	}
	return out.String(), nil
}
