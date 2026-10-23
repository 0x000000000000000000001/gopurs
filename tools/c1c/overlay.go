package c1c

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"sort"
	"strings"
)

// OverlaySource moves one generated module into its own package. Unqualified
// references to exported declarations of the old bundle are qualified with
// oldpkg, which is imported from oldPackagePath.
func OverlaySource(packageName, oldPackagePath, oldSource, newSource string) (string, error) {
	oldManifest, err := ExtractSource("old", "old.go", []byte(oldSource))
	if err != nil {
		return "", err
	}
	oldExports := map[string]bool{}
	for _, typ := range oldManifest.Types {
		if typ.Exported {
			oldExports[typ.Name] = true
		}
	}
	for _, function := range oldManifest.Functions {
		if function.Exported {
			oldExports[function.Name] = true
		}
	}
	for _, value := range oldManifest.Values {
		if ast.IsExported(value) {
			oldExports[value] = true
		}
	}
	newManifest, err := ExtractSource("new", "new.go", []byte(newSource))
	if err != nil {
		return "", err
	}
	local := map[string]bool{}
	for _, typ := range newManifest.Types {
		local[typ.Name] = true
	}
	for _, function := range newManifest.Functions {
		local[function.Name] = true
	}
	for _, value := range newManifest.Values {
		local[value] = true
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "new.go", newSource, parser.ParseComments)
	if err != nil {
		return "", err
	}
	var edits []struct {
		start int
		end   int
		text  string
	}
	src := []byte(newSource)
	for _, declaration := range file.Decls {
		ast.Inspect(declaration, func(node ast.Node) bool {
			switch value := node.(type) {
			case *ast.SelectorExpr:
				// Do not rewrite the selection half of an existing selector.
				ast.Inspect(value.X, func(inner ast.Node) bool {
					return true
				})
				return false
			case *ast.Ident:
				if local[value.Name] || !oldExports[value.Name] {
					return true
				}
				start := fset.Position(value.Pos()).Offset
				end := fset.Position(value.End()).Offset
				if start < 0 || end > len(src) || start >= end {
					return true
				}
				edits = append(edits, struct {
					start int
					end   int
					text  string
				}{start, end, "oldpkg." + value.Name})
			}
			return true
		})
	}
	sort.Slice(edits, func(i, j int) bool { return edits[i].start > edits[j].start })
	out := append([]byte(nil), src...)
	for _, edit := range edits {
		out = append(out[:edit.start], append([]byte(edit.text), out[edit.end:]...)...)
	}
	text := string(out)
	text = strings.Replace(text, "package purescript", "package "+packageName, 1)
	importLine := fmt.Sprintf("oldpkg %q\n", oldPackagePath)
	if strings.Contains(text, "import (") {
		text = strings.Replace(text, "import (\n", "import (\n\t"+importLine, 1)
	} else {
		text = strings.Replace(text, "package "+packageName+"\n", "package "+packageName+"\n\nimport "+importLine, 1)
	}
	return text, nil
}

// ExportedSymbols returns the top-level exported names in a generated source.
func ExportedSymbols(source string) ([]string, error) {
	manifest, err := ExtractSource("source", "source.go", []byte(source))
	if err != nil {
		return nil, err
	}
	set := map[string]bool{}
	for _, typ := range manifest.Types {
		if typ.Exported {
			set[typ.Name] = true
		}
	}
	for _, function := range manifest.Functions {
		if function.Exported {
			set[function.Name] = true
		}
	}
	for _, value := range manifest.Values {
		if ast.IsExported(value) {
			set[value] = true
		}
	}
	out := make([]string, 0, len(set))
	for name := range set {
		out = append(out, name)
	}
	sort.Strings(out)
	return out, nil
}

// CompileCheck parses the generated overlay and rejects malformed output.
func CompileCheck(packageName, source string) error {
	var scan scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("overlay.go", -1, len(source))
	_ = scan
	_ = file
	_, err := parser.ParseFile(fset, "overlay.go", source, parser.AllErrors)
	if err != nil {
		return err
	}
	if !bytes.Contains([]byte(source), []byte("package "+packageName)) {
		return fmt.Errorf("missing overlay package name %q", packageName)
	}
	return nil
}
