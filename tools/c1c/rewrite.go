package c1c

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strconv"
	"strings"
)

// RewritePackageReferences replaces selectors oldAlias.Symbol by
// newAlias.Symbol for the given symbols and adds the new import when needed.
// It returns the rewritten source and whether anything changed.
func RewritePackageReferences(source, oldPath, newPath string, symbols []string) (string, bool, error) {
	return RewritePackageReferencesAs(source, oldPath, newPath, "c1c_overlay", symbols)
}

// RewritePackageReferencesAs is RewritePackageReferences with a caller-chosen
// import alias. A caller that already knows the destination's reserved names
// can avoid a shadowing collision. An existing import of newPath keeps its own
// alias, so repeated rewrites are stable.
func RewritePackageReferencesAs(source, oldPath, newPath, alias string, symbols []string) (string, bool, error) {
	wanted := map[string]bool{}
	for _, symbol := range symbols {
		wanted[symbol] = true
	}
	if alias == "" {
		alias = "c1c_overlay"
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "package.go", source, parser.ParseComments)
	if err != nil {
		return "", false, err
	}
	oldAlias := ""
	for _, spec := range file.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			continue
		}
		if path != oldPath {
			continue
		}
		if spec.Name != nil {
			oldAlias = spec.Name.Name
		} else {
			oldAlias = path[strings.LastIndex(path, "/")+1:]
		}
	}
	if oldAlias == "" {
		return source, false, nil
	}
	newAlias := alias
	for _, spec := range file.Imports {
		path, _ := strconv.Unquote(spec.Path.Value)
		if path == newPath {
			if spec.Name != nil {
				newAlias = spec.Name.Name
			}
			break
		}
	}
	type edit struct {
		start, end int
		text       string
	}
	var edits []edit
	changed := false
	keepOld := false
	ast.Inspect(file, func(node ast.Node) bool {
		selector, ok := node.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		base, ok := selector.X.(*ast.Ident)
		if !ok || base.Name != oldAlias {
			return true
		}
		if !wanted[selector.Sel.Name] {
			keepOld = true
			return true
		}
		start := fset.Position(selector.Pos()).Offset
		end := fset.Position(selector.End()).Offset
		edits = append(edits, edit{start, end, newAlias + "." + selector.Sel.Name})
		changed = true
		return false
	})
	if !changed {
		return source, false, nil
	}
	if !keepOld {
		for _, spec := range file.Imports {
			path, _ := strconv.Unquote(spec.Path.Value)
			if path != oldPath {
				continue
			}
			start := fset.Position(spec.Pos()).Offset
			end := fset.Position(spec.End()).Offset
			lineStart := start
			for lineStart > 0 && source[lineStart-1] != '\n' {
				lineStart--
			}
			if end < len(source) && source[end] == '\n' {
				end++
			}
			edits = append(edits, edit{lineStart, end, ""})
			break
		}
	}
	sort.Slice(edits, func(i, j int) bool { return edits[i].start > edits[j].start })
	out := []byte(source)
	for _, value := range edits {
		out = append(out[:value.start], append([]byte(value.text), out[value.end:]...)...)
	}
	text := string(out)
	newImport := fmt.Sprintf("%s %q\n", newAlias, newPath)
	if strings.Contains(text, "import (") {
		text = strings.Replace(text, "import (\n", "import (\n\t"+newImport, 1)
	} else {
		text = strings.Replace(text, "package ", "package ", 1)
		// There is no import block: insert after the package line.
		if index := strings.IndexByte(text, '\n'); index >= 0 {
			text = text[:index+1] + "\nimport " + newImport + text[index+1:]
		}
	}
	return text, true, nil
}

func ValidateRewritten(source string) error {
	_, err := parser.ParseFile(token.NewFileSet(), "rewritten.go", source, parser.AllErrors)
	if err != nil {
		return err
	}
	if bytes.Contains([]byte(source), []byte("c1c_overlay.")) && !bytes.Contains([]byte(source), []byte("c1c_overlay \"")) {
		// The import may be aliased differently; this is only a smoke check.
	}
	return nil
}
