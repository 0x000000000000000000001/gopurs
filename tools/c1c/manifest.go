// Package c1c extracts the stable ABI surface of generated Go modules.
package c1c

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"sort"
)

type Import struct {
	Alias string `json:"alias,omitempty"`
	Path  string `json:"path"`
}

type Field struct {
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Exported bool   `json:"exported"`
	Embedded bool   `json:"embedded,omitempty"`
}

type Method struct {
	Name      string `json:"name"`
	Signature string `json:"signature"`
	Exported  bool   `json:"exported"`
}

type Type struct {
	Name     string   `json:"name"`
	Kind     string   `json:"kind"`
	Exported bool     `json:"exported"`
	Fields   []Field  `json:"fields,omitempty"`
	Methods  []Method `json:"methods,omitempty"`
}

type Function struct {
	Name      string `json:"name"`
	Receiver  string `json:"receiver,omitempty"`
	Signature string `json:"signature"`
	Exported  bool   `json:"exported"`
	Init      bool   `json:"init,omitempty"`
}

type Manifest struct {
	Module    string     `json:"module"`
	Package   string     `json:"package"`
	Imports   []Import   `json:"imports,omitempty"`
	Types     []Type     `json:"types,omitempty"`
	Functions []Function `json:"functions,omitempty"`
	Values    []string   `json:"values,omitempty"`
}

func Extract(module, path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ExtractSource(module, path, data)
}

func ExtractSource(module, path string, data []byte) (*Manifest, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, data, parser.SkipObjectResolution)
	if err != nil {
		return nil, err
	}
	manifest := &Manifest{Module: module, Package: file.Name.Name}
	types := map[string]*Type{}
	methods := map[string][]Method{}
	var values []string
	for _, spec := range file.Imports {
		value := ""
		if spec.Path != nil {
			value = spec.Path.Value
		}
		value = trimQuotes(value)
		alias := ""
		if spec.Name != nil {
			alias = spec.Name.Name
		}
		manifest.Imports = append(manifest.Imports, Import{Alias: alias, Path: value})
	}
	for _, declaration := range file.Decls {
		switch value := declaration.(type) {
		case *ast.FuncDecl:
			signature, err := nodeString(fset, value.Type)
			if err != nil {
				return nil, err
			}
			receiver := receiverName(value)
			function := Function{Name: value.Name.Name, Receiver: receiver, Signature: signature, Exported: ast.IsExported(value.Name.Name), Init: value.Name.Name == "init"}
			if receiver != "" {
				methods[receiver] = append(methods[receiver], Method{Name: function.Name, Signature: signature, Exported: function.Exported})
			} else {
				manifest.Functions = append(manifest.Functions, function)
			}
		case *ast.GenDecl:
			for _, item := range value.Specs {
				switch spec := item.(type) {
				case *ast.TypeSpec:
					kind := "alias"
					if spec.Assign == 0 {
						kind = "defined"
					}
					fields := []Field{}
					if structure, ok := spec.Type.(*ast.StructType); ok {
						kind = "struct"
						for _, field := range structure.Fields.List {
							fieldType, err := nodeString(fset, field.Type)
							if err != nil {
								return nil, err
							}
							if len(field.Names) == 0 {
								fields = append(fields, Field{Type: fieldType, Embedded: true})
								continue
							}
							for _, name := range field.Names {
								fields = append(fields, Field{Name: name.Name, Type: fieldType, Exported: ast.IsExported(name.Name)})
							}
						}
					}
					types[spec.Name.Name] = &Type{Name: spec.Name.Name, Kind: kind, Exported: ast.IsExported(spec.Name.Name), Fields: fields}
				case *ast.ValueSpec:
					for _, name := range spec.Names {
						values = append(values, name.Name)
					}
				}
			}
		}
	}
	for name, value := range types {
		value.Methods = methods[name]
		manifest.Types = append(manifest.Types, *value)
	}
	sort.Slice(manifest.Imports, func(i, j int) bool { return manifest.Imports[i].Path < manifest.Imports[j].Path })
	sort.Slice(manifest.Types, func(i, j int) bool { return manifest.Types[i].Name < manifest.Types[j].Name })
	sort.Slice(manifest.Functions, func(i, j int) bool { return manifest.Functions[i].Name < manifest.Functions[j].Name })
	sort.Strings(values)
	manifest.Values = values
	return manifest, nil
}

func JSON(manifest *Manifest) ([]byte, error) { return json.MarshalIndent(manifest, "", "  ") }

func trimQuotes(value string) string {
	if len(value) >= 2 && value[0] == '"' && value[len(value)-1] == '"' {
		return value[1 : len(value)-1]
	}
	return value
}

func receiverName(declaration *ast.FuncDecl) string {
	if declaration.Recv == nil || len(declaration.Recv.List) == 0 {
		return ""
	}
	return typeName(declaration.Recv.List[0].Type)
}

func typeName(expression ast.Expr) string {
	switch value := expression.(type) {
	case *ast.StarExpr:
		return typeName(value.X)
	case *ast.Ident:
		return value.Name
	case *ast.IndexExpr:
		return typeName(value.X)
	case *ast.IndexListExpr:
		return typeName(value.X)
	default:
		return ""
	}
}

func nodeString(fset *token.FileSet, node ast.Node) (string, error) {
	var output bytes.Buffer
	if err := printer.Fprint(&output, fset, node); err != nil {
		return "", err
	}
	return output.String(), nil
}

func (m *Manifest) Validate() error {
	if m.Module == "" {
		return fmt.Errorf("empty module")
	}
	if m.Package == "" {
		return fmt.Errorf("empty package")
	}
	for _, typ := range m.Types {
		for _, method := range typ.Methods {
			if method.Name == "" || method.Signature == "" {
				return fmt.Errorf("invalid method on %s", typ.Name)
			}
		}
	}
	return nil
}
