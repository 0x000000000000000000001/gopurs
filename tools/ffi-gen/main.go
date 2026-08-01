package main

import (
	"bytes"
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
	"syscall/js"
)

var fset = token.NewFileSet()

type TypeNode struct {
	Type string      `json:"type"` // "Named", "Func", "Array", "Map", "Unknown"
	Name string      `json:"name,omitempty"`
	Args []*TypeNode `json:"args,omitempty"`
	Ret  *TypeNode   `json:"ret,omitempty"`
	Elem *TypeNode   `json:"elem,omitempty"`
	Key  *TypeNode   `json:"key,omitempty"`
	Val  *TypeNode   `json:"val,omitempty"`
}

type FFIDecl struct {
	Name       string      `json:"name"`
	IsVar      bool        `json:"isVar"`
	TypeParams []string    `json:"typeParams"`
	Args       []*TypeNode `json:"args"`
	Ret        *TypeNode   `json:"ret"`
}

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

func parseFFI(this js.Value, args []js.Value) any {
	content := args[0].String()
	
	ffiMarkerIdx := strings.Index(content, "// --- Auto-generated FFI wrappers ---")
	if ffiMarkerIdx != -1 {
		content = content[:ffiMarkerIdx]
	}
	
	src := content
	if !strings.Contains(content, "package ") {
		src = "package main\n" + content
	}
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return "[]"
	}

	var decls []FFIDecl

	for _, decl := range f.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			funcName := funcDecl.Name.Name
			if len(funcName) == 0 {
				continue
			}
			c := funcName[0]
			if c != '_' && (c < 'A' || c > 'Z') {
				continue
			}
			
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
			
			if typeParamNames == nil { typeParamNames = []string{} }
			if parsedArgs == nil { parsedArgs = []*TypeNode{} }

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
						varName := name.Name
						if len(varName) == 0 {
							continue
						}
						c := varName[0]
						if c != '_' && (c < 'A' || c > 'Z') {
							continue
						}
						decls = append(decls, FFIDecl{
							Name:  varName,
							IsVar: true,
							TypeParams: []string{},
							Args: []*TypeNode{},
						})
					}
				}
			}
		}
	}

	if decls == nil {
		decls = []FFIDecl{}
	}

	jsonBytes, _ := json.Marshal(decls)
	return string(jsonBytes)
}

func main() {
	js.Global().Set("parseFFI", js.FuncOf(parseFFI))
	select {}
}
