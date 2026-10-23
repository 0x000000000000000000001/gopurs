package c1c

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strconv"
	"strings"
)

// Dependency is one generated package imported by a qualified overlay. The
// caller owns the path scheme and may reuse it for publication metadata.
type Dependency struct {
	Package string
	Path    string
	Alias   string
}

// BuildOverlaySource moves the declarations of one generated module into a
// separate package. Symbols owned by other generated packages are qualified
// with fresh imports. Symbols defined by this source (locals, helpers and
// top-level declarations) stay unqualified.
//
// symbols maps an exported top-level name to its generated package. The caller
// must reject private cross-package references before calling this function;
// they cannot be expressed through an import.
func BuildOverlaySource(packageName, source string, symbols map[string]string, packagePath func(string) string) (string, []Dependency, error) {
	if packageName == "" || packagePath == nil {
		return "", nil, fmt.Errorf("overlay requires a package name and path resolver")
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "overlay-input.go", source, parser.ParseComments)
	if err != nil {
		return "", nil, err
	}
	reserved := map[string]bool{}
	ast.Inspect(file, func(node ast.Node) bool {
		if ident, ok := node.(*ast.Ident); ok {
			reserved[ident.Name] = true
		}
		return true
	})
	for _, spec := range file.Imports {
		if spec.Name != nil {
			reserved[spec.Name.Name] = true
		}
	}
	selectors := map[*ast.Ident]bool{}
	ast.Inspect(file, func(node ast.Node) bool {
		if selector, ok := node.(*ast.SelectorExpr); ok {
			selectors[selector.Sel] = true
		}
		return true
	})

	type importSpec struct {
		alias string
		path  string
	}
	existing := []importSpec{}
	importAlias := map[string]string{}
	for _, spec := range file.Imports {
		value, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			return "", nil, err
		}
		alias := ""
		if spec.Name != nil {
			alias = spec.Name.Name
		}
		existing = append(existing, importSpec{alias, value})
		effective := alias
		if effective == "" {
			effective = value[strings.LastIndex(value, "/")+1:]
		}
		importAlias[value] = effective
	}

	dependencies := map[string]Dependency{}
	aliases := map[string]string{}
	aliasFor := func(pkg string) (string, error) {
		if alias, ok := aliases[pkg]; ok {
			return alias, nil
		}
		path := packagePath(pkg)
		if path == "" {
			return "", fmt.Errorf("empty overlay path for package %q", pkg)
		}
		if alias, ok := importAlias[path]; ok {
			aliases[pkg] = alias
			dependencies[pkg] = Dependency{Package: pkg, Path: path, Alias: alias}
			return alias, nil
		}
		digest := sha256.Sum256([]byte(pkg))
		encoded := hex.EncodeToString(digest[:])
		alias := ""
		for width := 6; width <= 64; width++ {
			candidate := "g" + encoded[:width]
			if !reserved[candidate] {
				alias = candidate
				break
			}
		}
		for alias == "" {
			encoded += "_"
			if !reserved["g"+encoded] {
				alias = "g" + encoded
			}
		}
		reserved[alias] = true
		aliases[pkg] = alias
		dependencies[pkg] = Dependency{Package: pkg, Path: path, Alias: alias}
		return alias, nil
	}

	type edit struct {
		start int
		end   int
		text  string
	}
	var edits []edit
	for _, ident := range file.Unresolved {
		if selectors[ident] {
			continue
		}
		pkg, ok := symbols[ident.Name]
		if !ok {
			continue
		}
		alias, err := aliasFor(pkg)
		if err != nil {
			return "", nil, err
		}
		start := fset.Position(ident.Pos()).Offset
		end := fset.Position(ident.End()).Offset
		edits = append(edits, edit{start, end, alias + "." + ident.Name})
	}

	// Rebuild one import block. Generated sources have no comments inside it,
	// and a stable block avoids merging rules that differ between hosts.
	lines := []string{}
	seen := map[string]bool{}
	addImport := func(spec importSpec) {
		if seen[spec.path] {
			return
		}
		seen[spec.path] = true
		if spec.alias == "" {
			lines = append(lines, "\t"+strconv.Quote(spec.path))
		} else {
			lines = append(lines, "\t"+spec.alias+" "+strconv.Quote(spec.path))
		}
	}
	for _, spec := range existing {
		addImport(spec)
	}
	names := make([]string, 0, len(dependencies))
	for pkg := range dependencies {
		names = append(names, pkg)
	}
	sort.Strings(names)
	for _, pkg := range names {
		dependency := dependencies[pkg]
		addImport(importSpec{dependency.Alias, dependency.Path})
	}
	block := "import (\n" + strings.Join(lines, "\n") + "\n)\n"
	var first, last int = -1, -1
	for _, declaration := range file.Decls {
		generic, ok := declaration.(*ast.GenDecl)
		if !ok || generic.Tok != token.IMPORT {
			continue
		}
		start := fset.Position(generic.Pos()).Offset
		end := fset.Position(generic.End()).Offset
		if first < 0 || start < first {
			first = start
		}
		if end > last {
			last = end
		}
	}
	if first >= 0 {
		if last < len(source) && source[last] == '\n' {
			last++
		}
		edits = append(edits, edit{first, last, block})
	} else {
		offset := fset.Position(file.Name.End()).Offset
		edits = append(edits, edit{offset, offset, "\n\n" + block})
	}
	sort.Slice(edits, func(i, j int) bool {
		if edits[i].start != edits[j].start {
			return edits[i].start > edits[j].start
		}
		return edits[i].end > edits[j].end
	})
	out := []byte(source)
	for _, value := range edits {
		if value.start < 0 || value.end < value.start || value.end > len(out) {
			return "", nil, fmt.Errorf("invalid overlay edit")
		}
		out = append(out[:value.start], append([]byte(value.text), out[value.end:]...)...)
	}
	text := string(out)
	text = strings.Replace(text, "package "+file.Name.Name, "package "+packageName, 1)
	if err := CompileCheck(packageName, text); err != nil {
		return "", nil, err
	}
	result := make([]Dependency, 0, len(dependencies))
	for _, pkg := range names {
		result = append(result, dependencies[pkg])
	}
	return text, result, nil
}
