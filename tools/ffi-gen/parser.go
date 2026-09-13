package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/scanner"
	"go/token"
	"sort"
	"strings"
)

var fset = token.NewFileSet()

func parseExprToTypeNode(expr ast.Expr) *TypeNode {
	if expr == nil {
		return nil
	}
	switch t := expr.(type) {
	case *ast.Ident:
		return &TypeNode{Type: "Named", Name: t.Name}
	case *ast.SelectorExpr:
		var buf bytes.Buffer
		printer.Fprint(&buf, fset, t)
		return &TypeNode{Type: "Named", Name: buf.String()}
	case *ast.InterfaceType:
		return &TypeNode{Type: "Named", Name: "any"}
	case *ast.ArrayType:
		return &TypeNode{Type: "Array", Elem: parseExprToTypeNode(t.Elt)}
	case *ast.MapType:
		return &TypeNode{Type: "Map", Key: parseExprToTypeNode(t.Key), Val: parseExprToTypeNode(t.Value)}
	case *ast.FuncType:
		var args []*TypeNode
		if t.Params != nil {
			for _, field := range t.Params.List {
				typNode := parseExprToTypeNode(field.Type)
				if len(field.Names) == 0 {
					args = append(args, typNode)
				} else {
					for range field.Names {
						args = append(args, typNode)
					}
				}
			}
		}
		var ret *TypeNode
		if t.Results != nil && len(t.Results.List) > 0 {
			ret = parseExprToTypeNode(t.Results.List[0].Type)
		}
		if args == nil {
			args = []*TypeNode{} // Ensure empty array, not null
		}
		return &TypeNode{Type: "Func", Args: args, Ret: ret}
	default:
		var buf bytes.Buffer
		printer.Fprint(&buf, fset, expr)
		return &TypeNode{Type: "Unknown", Name: buf.String()}
	}
}

func parseFFI(content string, prefixes ...string) (string, error) {
	ffiMarkerIdx := strings.Index(content, "// --- Auto-generated FFI wrappers ---")
	if ffiMarkerIdx != -1 {
		content = content[:ffiMarkerIdx]
	}

	src := content
	// Preserve source lines and columns while allowing later //line directives.
	const packagePrefix = "package main\n//line :1:1\n"
	addedPackage := !strings.Contains(content, "package ")
	if addedPackage {
		src = packagePrefix + content
	}
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		if parseErrors, ok := err.(scanner.ErrorList); ok && addedPackage {
			for _, parseError := range parseErrors {
				parseError.Pos.Offset -= len(packagePrefix)
			}
		}
		return "", err
	}

	renamedObjects := make(map[*ast.Object]bool)
	prefix := ""
	if len(prefixes) > 0 {
		prefix = prefixes[0]
		for name, object := range f.Scope.Objects {
			if (object.Kind == ast.Fun || object.Kind == ast.Var) && isFFIName(name) {
				renamedObjects[object] = true
			}
		}
	}
	declarationName := func(name *ast.Ident) string {
		if renamedObjects[name.Obj] {
			return prefix + name.Name
		}
		return name.Name
	}

	var decls []FFIDecl

	for _, decl := range f.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			if !isFFIName(funcDecl.Name.Name) {
				continue
			}
			funcName := declarationName(funcDecl.Name)

			var typeParamNames []string
			if funcDecl.Type.TypeParams != nil {
				for _, field := range funcDecl.Type.TypeParams.List {
					for _, name := range field.Names {
						typeParamNames = append(typeParamNames, name.Name)
					}
				}
			}

			var parsedArgs []*TypeNode
			if funcDecl.Type.Params != nil {
				for _, field := range funcDecl.Type.Params.List {
					typNode := parseExprToTypeNode(field.Type)
					if len(field.Names) == 0 {
						parsedArgs = append(parsedArgs, typNode)
					} else {
						for range field.Names {
							parsedArgs = append(parsedArgs, typNode)
						}
					}
				}
			}

			var retNode *TypeNode
			if funcDecl.Type.Results != nil && len(funcDecl.Type.Results.List) > 0 {
				retNode = parseExprToTypeNode(funcDecl.Type.Results.List[0].Type)
			}

			if typeParamNames == nil {
				typeParamNames = []string{}
			}
			if parsedArgs == nil {
				parsedArgs = []*TypeNode{}
			}

			decls = append(decls, FFIDecl{
				Name:       funcName,
				IsVar:      false,
				TypeParams: typeParamNames,
				Args:       parsedArgs,
				Ret:        retNode,
			})

		} else if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.VAR {
			for _, spec := range genDecl.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					for _, name := range valueSpec.Names {
						if !isFFIName(name.Name) {
							continue
						}
						varName := declarationName(name)
						decls = append(decls, FFIDecl{
							Name:       varName,
							IsVar:      true,
							TypeParams: []string{},
							Args:       []*TypeNode{},
						})
					}
				}
			}
		}
	}

	if decls == nil {
		decls = []FFIDecl{}
	}

	var result any = decls
	if len(prefixes) > 0 {
		var offsets []int
		selectorNames := make(map[*ast.Ident]bool)
		ast.Inspect(f, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.SelectorExpr:
				selectorNames[node.Sel] = true
			case *ast.Ident:
				if renamedObjects[node.Obj] && !selectorNames[node] {
					offset := fset.PositionFor(node.Pos(), false).Offset
					if addedPackage {
						offset -= len(packagePrefix)
					}
					offsets = append(offsets, offset)
				}
			}
			return true
		})
		sort.Ints(offsets)
		var renamed strings.Builder
		start := 0
		for _, offset := range offsets {
			renamed.WriteString(content[start:offset])
			renamed.WriteString(prefix)
			start = offset
		}
		renamed.WriteString(content[start:])
		result = struct {
			Decls   []FFIDecl `json:"decls"`
			Content string    `json:"content"`
		}{decls, renamed.String()}
	}
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return "", fmt.Errorf("encode FFI declarations: %w", err)
	}
	return string(jsonBytes), nil
}

func isFFIName(name string) bool {
	return len(name) > 0 && (name[0] == '_' || name[0] >= 'A' && name[0] <= 'Z')
}
