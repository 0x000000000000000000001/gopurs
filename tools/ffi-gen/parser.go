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
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var fset = token.NewFileSet()

// signatureParser retains source spelling while proving the two special ABI
// identities from lexical declarations. Imported aliases requiring export data
// deliberately remain unresolved here; resolveSources is the authoritative
// go/types checker used when assembling the generated package.
type signatureParser struct {
	imports  map[string]string
	visiting map[*ast.Object]bool
}

func newSignatureParser(file *ast.File) *signatureParser {
	p := &signatureParser{imports: map[string]string{}, visiting: map[*ast.Object]bool{}}
	for _, spec := range file.Imports {
		path, _ := strconv.Unquote(spec.Path.Value)
		name := filepath.Base(path)
		if spec.Name != nil {
			name = spec.Name.Name
		}
		p.imports[name] = path
	}
	return p
}

func sourceText(node ast.Node) string {
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, node)
	return buf.String()
}

func singleResult(results []*TypeNode) *TypeNode {
	if len(results) == 1 {
		return results[0]
	}
	return nil
}

func (p *signatureParser) fields(fields *ast.FieldList) ([]*TypeNode, []string) {
	nodes, names := []*TypeNode{}, []string{}
	if fields != nil {
		for _, field := range fields.List {
			node := p.parse(field.Type)
			if len(field.Names) == 0 {
				nodes, names = append(nodes, node), append(names, "")
			} else {
				for _, name := range field.Names {
					nodes, names = append(nodes, node), append(names, name.Name)
				}
			}
		}
	}
	return nodes, names
}

func (p *signatureParser) parse(expr ast.Expr) *TypeNode {
	if expr == nil {
		return nil
	}
	node := &TypeNode{Args: []*TypeNode{}, ArgNames: []string{}, Results: []*TypeNode{}, ResultNames: []string{}}
	switch t := expr.(type) {
	case *ast.ParenExpr:
		return p.parse(t.X)
	case *ast.Ident:
		node.Type, node.Name = "Named", t.Name
		if t.Obj == nil {
			if t.Name == "error" {
				node.Identity = "builtin.error"
			}
		} else if spec, ok := t.Obj.Decl.(*ast.TypeSpec); ok && !p.visiting[t.Obj] {
			p.visiting[t.Obj] = true
			underlying := p.parse(spec.Type)
			delete(p.visiting, t.Obj)
			if spec.Assign.IsValid() {
				node.Underlying, node.Identity = underlying, underlying.Identity
			} else if hasStructLayout(underlying) {
				node.Underlying = underlying
			}
		}
	case *ast.SelectorExpr:
		node.Type, node.Name = "Named", sourceText(t)
		if qualifier, ok := t.X.(*ast.Ident); ok && qualifier.Obj == nil {
			node.PackagePath = p.imports[qualifier.Name]
			if node.PackagePath == "context" && t.Sel.Name == "Context" {
				node.Identity = "context.Context"
			}
		}
	case *ast.StructType:
		node.Type, node.Name = "Struct", sourceText(t)
		node.Fields = []StructField{}
		for _, field := range t.Fields.List {
			value := p.parse(field.Type)
			tag := ""
			if field.Tag != nil {
				tag, _ = strconv.Unquote(field.Tag.Value)
			}
			names := field.Names
			if len(names) == 0 {
				name := embeddedFieldName(field.Type)
				node.Fields = append(node.Fields, StructField{Name: name, Type: value, Embedded: true, Exported: ast.IsExported(name), Tag: tag})
			} else {
				for _, name := range names {
					node.Fields = append(node.Fields, StructField{Name: name.Name, Type: value, Exported: ast.IsExported(name.Name), Tag: tag})
				}
			}
		}
	case *ast.InterfaceType:
		if t.Methods == nil || len(t.Methods.List) == 0 {
			node.Type, node.Name = "Named", "any"
		} else {
			node.Type, node.Name = "Interface", sourceText(t)
		}
	case *ast.StarExpr:
		node.Type, node.Elem, node.Name = "Pointer", p.parse(t.X), sourceText(t)
	case *ast.ArrayType:
		node.Type, node.Elem = "Array", p.parse(t.Elt)
		if t.Len != nil {
			node.Type, node.Length, node.Name = "FixedArray", sourceText(t.Len), sourceText(t)
		}
	case *ast.Ellipsis:
		node.Type, node.Elem, node.Name = "Variadic", p.parse(t.Elt), sourceText(t)
	case *ast.MapType:
		node.Type, node.Key, node.Val = "Map", p.parse(t.Key), p.parse(t.Value)
	case *ast.FuncType:
		node.Type = "Func"
		node.Args, node.ArgNames = p.fields(t.Params)
		node.Results, node.ResultNames = p.fields(t.Results)
		node.Ret = singleResult(node.Results)
	default:
		node.Type, node.Name = "Unknown", sourceText(expr)
	}
	return node
}

func hasStructLayout(node *TypeNode) bool {
	for node != nil {
		if node.Type == "Struct" {
			return true
		}
		node = node.Underlying
	}
	return false
}

func embeddedFieldName(expr ast.Expr) string {
	switch value := expr.(type) {
	case *ast.Ident:
		return value.Name
	case *ast.SelectorExpr:
		return value.Sel.Name
	case *ast.StarExpr:
		return embeddedFieldName(value.X)
	case *ast.IndexExpr:
		return embeddedFieldName(value.X)
	case *ast.IndexListExpr:
		return embeddedFieldName(value.X)
	}
	return ""
}

func parseFFI(content string, prefixes ...string) (string, error) {
	ffiMarkerIdx := strings.Index(content, "// --- Auto-generated FFI wrappers ---")
	if ffiMarkerIdx != -1 {
		content = content[:ffiMarkerIdx]
	}

	src := content
	// Preserve source lines and columns while allowing later //line directives.
	const packagePrefix = "package main\n//line :1:1\n"
	var sourceScanner scanner.Scanner
	sourceFile := token.NewFileSet().AddFile("", -1, len(content))
	sourceScanner.Init(sourceFile, []byte(content), nil, 0)
	_, firstToken, _ := sourceScanner.Scan()
	addedPackage := firstToken != token.PACKAGE
	packageSemicolon := -1
	if !addedPackage {
		sourceScanner.Scan() // package name; ParseFile below validates it.
		position, kind, literal := sourceScanner.Scan()
		if kind == token.SEMICOLON && literal == ";" {
			packageSemicolon = sourceFile.Offset(position)
		}
	}
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

	types := newSignatureParser(f)
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

			parsedArgs, argNames := types.fields(funcDecl.Type.Params)
			results, resultNames := types.fields(funcDecl.Type.Results)
			if typeParamNames == nil {
				typeParamNames = []string{}
			}
			position := fset.Position(funcDecl.Name.Pos())
			var receiver *TypeNode
			if funcDecl.Recv != nil && len(funcDecl.Recv.List) > 0 {
				receiver = types.parse(funcDecl.Recv.List[0].Type)
			}
			decls = append(decls, FFIDecl{
				SchemaVersion: ffiSchemaVersion,
				Name:          funcName, IsVar: false, TypeParams: typeParamNames,
				Args: parsedArgs, ArgNames: argNames,
				Ret: singleResult(results), Results: results, ResultNames: resultNames,
				Receiver: receiver, Signature: sourceText(funcDecl.Type),
				Line: position.Line, Column: position.Column,
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
							SchemaVersion: ffiSchemaVersion,
							Name:          varName,
							IsVar:         true,
							TypeParams:    []string{},
							Args:          []*TypeNode{},
							ArgNames:      []string{}, Results: []*TypeNode{}, ResultNames: []string{},
							ValueType: types.parse(valueSpec.Type),
							Line:      fset.Position(name.Pos()).Line, Column: fset.Position(name.Pos()).Column,
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
		type sourceEdit struct {
			start, end int
			text       string
		}
		var edits []sourceEdit
		if !addedPackage {
			// Main supplies the destination package. Remove only its source
			// tokens, retaining comments, whitespace and literal contents.
			packageStart := fset.PositionFor(f.Package, false).Offset
			nameStart := fset.PositionFor(f.Name.Pos(), false).Offset
			nameEnd := fset.PositionFor(f.Name.End(), false).Offset
			edits = append(edits, sourceEdit{packageStart, packageStart + len("package"), ""}, sourceEdit{nameStart, nameEnd, ""})
			if packageSemicolon >= 0 {
				edits = append(edits, sourceEdit{packageSemicolon, packageSemicolon + 1, ""})
			}
		}
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
					edits = append(edits, sourceEdit{offset, offset, prefix})
				}
			}
			return true
		})
		sort.Slice(edits, func(i, j int) bool { return edits[i].start < edits[j].start })
		var renamed strings.Builder
		start := 0
		for _, edit := range edits {
			renamed.WriteString(content[start:edit.start])
			renamed.WriteString(edit.text)
			start = edit.end
		}
		renamed.WriteString(content[start:])
		result = struct {
			SchemaVersion int       `json:"schemaVersion"`
			Decls         []FFIDecl `json:"decls"`
			Content       string    `json:"content"`
		}{ffiSchemaVersion, decls, renamed.String()}
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
