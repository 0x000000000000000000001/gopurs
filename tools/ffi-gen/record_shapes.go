package main

import (
	"encoding/hex"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strconv"
	"strings"
)

// ShareRecordShapes shares immutable literal keys before either output layout
// is planned. Values remain at their original call sites and evaluation order.
// The prefix identifies the source file within its generated Go package.
func ShareRecordShapes(prefix, source string) (string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "records.go", source, parser.ParseComments)
	if err != nil {
		return "", err
	}
	aliases := map[string]bool{}
	for _, spec := range file.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil || path != "gopurs/output/gopurs_runtime" {
			continue
		}
		alias := "gopurs_runtime"
		if spec.Name != nil {
			alias = spec.Name.Name
		}
		if alias != "." && alias != "_" {
			aliases[alias] = true
		}
	}
	if len(aliases) == 0 {
		return source, nil
	}
	// Include local identifiers: a new package variable must not be shadowed at
	// any rewritten call site. Hex encoding makes arbitrary file prefixes unique.
	used := map[string]bool{}
	ast.Inspect(file, func(node ast.Node) bool {
		if ident, ok := node.(*ast.Ident); ok {
			used[ident.Name] = true
		}
		return true
	})
	type edit struct {
		start, end int
		text       string
	}
	var edits []edit
	var declarations strings.Builder
	shapes := map[string]string{}
	serial := 0
	namePrefix := "GopursRecordKeys_" + hex.EncodeToString([]byte(prefix)) + "_"
	ast.Inspect(file, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok || call.Ellipsis.IsValid() {
			return true
		}
		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		receiver, ok := selector.X.(*ast.Ident)
		if !ok || !aliases[receiver.Name] || receiver.Obj != nil {
			return true
		}
		if len(selector.Sel.Name) != len("RecordDict1") || !strings.HasPrefix(selector.Sel.Name, "RecordDict") {
			return true
		}
		size := int(selector.Sel.Name[len("RecordDict")] - '0')
		if size < 1 || size > 5 || len(call.Args) != 2*size {
			return true
		}
		keys := make([]string, size)
		var schema strings.Builder
		for i := range keys {
			literal, ok := call.Args[i].(*ast.BasicLit)
			if !ok || literal.Kind != token.STRING {
				return true
			}
			key, err := strconv.Unquote(literal.Value)
			if err != nil {
				return true
			}
			keys[i] = key
			fmt.Fprintf(&schema, "%d:%s", len(key), key)
		}
		name, found := shapes[schema.String()]
		if !found {
			for {
				name = namePrefix + strconv.Itoa(serial)
				serial++
				if !used[name] {
					break
				}
			}
			used[name] = true
			shapes[schema.String()] = name
			fmt.Fprintf(&declarations, "\nvar %s = %s.RecordKeys%d{", name, receiver.Name, size)
			for i, key := range keys {
				if i != 0 {
					declarations.WriteString(", ")
				}
				fmt.Fprintf(&declarations, "K%d: %s", i, strconv.Quote(key))
			}
			declarations.WriteString("}\n")
		}
		edits = append(edits,
			edit{fset.Position(selector.Sel.Pos()).Offset, fset.Position(selector.Sel.End()).Offset, "RecordShape" + strconv.Itoa(size)},
			edit{fset.Position(call.Args[0].Pos()).Offset, fset.Position(call.Args[size-1].End()).Offset, "&" + name},
		)
		return true
	})
	if len(edits) == 0 {
		return source, nil
	}
	sort.Slice(edits, func(i, j int) bool { return edits[i].start < edits[j].start })
	var result strings.Builder
	result.Grow(len(source) + declarations.Len())
	previous := 0
	for _, edit := range edits {
		if edit.start < previous {
			return "", fmt.Errorf("overlapping record shape edits")
		}
		result.WriteString(source[previous:edit.start])
		result.WriteString(edit.text)
		previous = edit.end
	}
	result.WriteString(source[previous:])
	result.WriteString(declarations.String())
	return result.String(), nil
}
