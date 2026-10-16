package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"sort"
	"strconv"
	"strings"
	"unicode/utf16"
)

// Source ownership is an input from the producer, not a guess based on a name
// prefix. Export data is supplied as bytes so WASM, Go and the Rust C transport
// have the same inputs; the resolver never invokes Go or reads the filesystem.
type SourceSymbol struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type SourceInput struct {
	Owner        string         `json:"owner"`
	Module       string         `json:"moduleName,omitempty"`
	Path         string         `json:"path"`
	Content      string         `json:"content"`
	Initializers []SourceSymbol `json:"initializers"`
	Records      []SourceRecord `json:"records,omitempty"`
}

// Type is the producer's exact Go type spelling. Labels are UTF-16 hex, not
// JSON strings whose isolated surrogates encoding/json would replace.
type SourceRecord struct {
	Type   string   `json:"type"`
	Labels []string `json:"labels"`
}

type BindRequest struct {
	PackagePath string            `json:"packagePath"`
	GoVersion   string            `json:"goVersion"`
	Arch        string            `json:"arch"`
	Files       []SourceInput     `json:"files"`
	Exports     map[string][]byte `json:"exports"`
}

type SourceReference struct {
	Start  int          `json:"start"` // UTF-16 code units, half-open.
	End    int          `json:"end"`
	Role   string       `json:"role"`
	Name   string       `json:"name"`
	Symbol SourceSymbol `json:"symbol"`
}

type SourceDeclaration struct {
	Start   int            `json:"start"`
	End     int            `json:"end"`
	Kind    string         `json:"kind"`
	Symbols []SourceSymbol `json:"symbols"`
}

type SourceImport struct {
	Path  string `json:"path"`
	Alias string `json:"alias"`
}

type BoundSource struct {
	Owner        string              `json:"owner"`
	Path         string              `json:"path"`
	Text         string              `json:"text"`
	Imports      []SourceImport      `json:"imports"`
	Identifiers  []string            `json:"identifiers"`
	References   []SourceReference   `json:"references"`
	Declarations []SourceDeclaration `json:"declarations"`
	Uses         map[string]int      `json:"uses"`
}

type BindResult struct {
	Format string        `json:"format"`
	Files  []BoundSource `json:"files"`
}

type sourceUnit struct {
	input SourceInput
	file  *ast.File
	token *token.File
	decls []ast.Decl
	bound BoundSource
}

func receiverObject(object *types.Func) *types.TypeName {
	receiver := object.Type().(*types.Signature).Recv()
	if receiver == nil {
		return nil
	}
	value := types.Unalias(receiver.Type())
	if pointer, ok := value.(*types.Pointer); ok {
		value = types.Unalias(pointer.Elem())
	}
	if named, ok := value.(*types.Named); ok {
		return named.Origin().Obj()
	}
	return nil
}

// Instantiation creates new field/method objects, but ownership belongs to the
// original declaration. In particular receiver variable names are irrelevant.
func sourceOrigin(object types.Object) types.Object {
	switch value := object.(type) {
	case *types.Var:
		return value.Origin()
	case *types.Func:
		return value.Origin()
	}
	return object
}

// Resolution is shared by the diagnostic binder and the package emitter.
// The latter keeps the typed graph here rather than transporting every use.
type resolvedSources struct {
	set   *token.FileSet
	units []*sourceUnit
	info  *types.Info
	pkg   *types.Package
	owned map[types.Object]SourceSymbol
}

func resolveSources(request BindRequest) (*resolvedSources, error) {
	if request.PackagePath == "" || request.GoVersion == "" || len(request.Files) == 0 {
		return nil, fmt.Errorf("Go source binding requires a package, Go version and source files")
	}
	sizes := types.SizesFor("gc", request.Arch)
	if sizes == nil {
		return nil, fmt.Errorf("unsupported Go target architecture: %s", request.Arch)
	}
	set := token.NewFileSet()
	inputs := append([]SourceInput(nil), request.Files...)
	sort.Slice(inputs, func(i, j int) bool { return inputs[i].Path < inputs[j].Path })
	units := make([]*sourceUnit, 0, len(inputs))
	files := make([]*ast.File, 0, len(inputs))
	byToken := map[*token.File]*sourceUnit{}
	for i, input := range inputs {
		if input.Owner == "" || input.Path == "" || i > 0 && input.Path == inputs[i-1].Path {
			return nil, fmt.Errorf("missing owner or duplicate Go source path: %q", input.Path)
		}
		file, err := parser.ParseFile(set, input.Path, input.Content, parser.ParseComments|parser.SkipObjectResolution)
		if err != nil {
			return nil, err
		}
		unit := &sourceUnit{input: input, file: file, token: set.File(file.Pos()), bound: BoundSource{
			Owner: input.Owner, Path: input.Path, Text: input.Content,
			Imports: []SourceImport{}, Identifiers: []string{}, References: []SourceReference{},
			Declarations: []SourceDeclaration{}, Uses: map[string]int{},
		}}
		for _, declaration := range file.Decls {
			if decl, ok := declaration.(*ast.GenDecl); ok && decl.Tok == token.IMPORT {
				continue
			}
			unit.decls = append(unit.decls, declaration)
		}
		byToken[unit.token] = unit
		units = append(units, unit)
		files = append(files, file)
	}
	info := &types.Info{Defs: map[*ast.Ident]types.Object{}, Uses: map[*ast.Ident]types.Object{}, Implicits: map[ast.Node]types.Object{}}
	config := types.Config{GoVersion: request.GoVersion, Sizes: sizes,
		Importer: importer.ForCompiler(set, "gc", func(path string) (io.ReadCloser, error) {
			data, ok := request.Exports[path]
			if !ok {
				return nil, fmt.Errorf("missing Go export data: %s", path)
			}
			return io.NopCloser(bytes.NewReader(data)), nil
		})}
	pkg, err := config.Check(request.PackagePath, set, files, info)
	if err != nil {
		return nil, fmt.Errorf("resolve Go source: %w", err)
	}
	owned := map[types.Object]SourceSymbol{}
	methodDeclarations := map[types.Object]bool{}
	for _, unit := range units {
		for _, declaration := range unit.decls {
			if method, ok := declaration.(*ast.FuncDecl); ok && method.Recv != nil {
				methodDeclarations[info.Defs[method.Name]] = true
			}
		}
	}
	for _, object := range info.Defs {
		if object == nil || object.Pkg() != pkg {
			continue
		}
		name := object.Name()
		if function, ok := object.(*types.Func); ok && name == "init" && receiverObject(function) == nil {
			continue // init has an explicit initialization identity, never a referable name.
		}
		if function, ok := object.(*types.Func); ok && receiverObject(function) != nil {
			if !methodDeclarations[object] {
				continue // Interface methods belong to the enclosing type declaration.
			}
			name = receiverObject(function).Name() + "." + name
		} else if object.Parent() != pkg.Scope() || name == "_" {
			continue
		}
		unit := byToken[set.File(object.Pos())]
		if unit == nil {
			return nil, fmt.Errorf("unowned Go definition: %s", name)
		}
		owned[object] = SourceSymbol{unit.input.Owner, name}
	}
	identities := map[SourceSymbol]string{}
	for _, unit := range units {
		initIndex := 0
		for _, declaration := range unit.decls {
			bound := SourceDeclaration{Start: unit.token.Offset(declaration.Pos()), End: unit.token.Offset(declaration.End()), Symbols: []SourceSymbol{}}
			switch decl := declaration.(type) {
			case *ast.FuncDecl:
				bound.Kind = "function"
				if decl.Recv != nil {
					bound.Kind = "method"
				}
				if decl.Recv == nil && decl.Name.Name == "init" {
					bound.Kind = "init"
					annotation, err := initializerAnnotation(decl)
					if err != nil {
						return nil, fmt.Errorf("%s: %w", unit.input.Path, err)
					}
					if initIndex < len(unit.input.Initializers) {
						if annotation != nil && *annotation != unit.input.Initializers[initIndex] {
							return nil, fmt.Errorf("conflicting initializer identity in %s", unit.input.Path)
						}
						bound.Symbols = append(bound.Symbols, unit.input.Initializers[initIndex])
					} else if annotation != nil {
						bound.Symbols = append(bound.Symbols, *annotation)
					} else if strings.HasPrefix(unit.input.Owner, "ffi/") {
						bound.Symbols = append(bound.Symbols, SourceSymbol{unit.input.Owner, fmt.Sprintf("init/%d", initIndex)})
					} else {
						return nil, fmt.Errorf("missing generated initializer identity in %s", unit.input.Path)
					}
					initIndex++
				} else {
					bound.Symbols = append(bound.Symbols, owned[info.Defs[decl.Name]])
				}
			case *ast.GenDecl:
				bound.Kind = decl.Tok.String()
				for _, spec := range decl.Specs {
					var names []*ast.Ident
					switch spec := spec.(type) {
					case *ast.ValueSpec:
						names = spec.Names
					case *ast.TypeSpec:
						names = []*ast.Ident{spec.Name}
					}
					for _, name := range names {
						if name.Name != "_" {
							bound.Symbols = append(bound.Symbols, owned[info.Defs[name]])
						}
					}
				}
				if len(bound.Symbols) == 0 {
					// Blank initializers have no referable symbol. Their source-order
					// token is only an initialization constraint, not a named identity.
					bound.Symbols = append(bound.Symbols, SourceSymbol{unit.input.Owner, fmt.Sprintf("blank/%d", len(unit.bound.Declarations))})
				}
			}
			for _, symbol := range bound.Symbols {
				if symbol.Owner != unit.input.Owner || symbol.Name == "" {
					return nil, fmt.Errorf("invalid Go declaration identity in %s: %+v", unit.input.Path, symbol)
				}
				if previous, exists := identities[symbol]; exists {
					return nil, fmt.Errorf("duplicate Go declaration identity in %s and %s: %+v", previous, unit.input.Path, symbol)
				}
				identities[symbol] = unit.input.Path
			}
			unit.bound.Declarations = append(unit.bound.Declarations, bound)
		}
		if len(unit.input.Initializers) > initIndex {
			return nil, fmt.Errorf("unused initializer identities in %s", unit.input.Path)
		}
	}
	return &resolvedSources{set, units, info, pkg, owned}, nil
}

func initializerAnnotation(declaration *ast.FuncDecl) (*SourceSymbol, error) {
	if declaration.Doc == nil {
		return nil, nil
	}
	var symbol *SourceSymbol
	for _, comment := range declaration.Doc.List {
		if !strings.HasPrefix(comment.Text, "//gopurs:init-") {
			continue
		}
		encoded, valid := strings.CutPrefix(comment.Text, "//gopurs:init-v1:")
		owner, name, ok := strings.Cut(encoded, ":")
		if !valid || !ok || symbol != nil {
			return nil, fmt.Errorf("invalid initializer annotation: %s", comment.Text)
		}
		decode := func(text string) (string, error) {
			if text == "" || len(text)%4 != 0 {
				return "", fmt.Errorf("invalid initializer UTF-16 identity")
			}
			units := []uint16{}
			for i := 0; i < len(text); i += 4 {
				unit, err := strconv.ParseUint(text[i:i+4], 16, 16)
				if err != nil {
					return "", err
				}
				units = append(units, uint16(unit))
			}
			return string(utf16.Decode(units)), nil
		}
		decodedOwner, err := decode(owner)
		if err != nil {
			return nil, err
		}
		decodedName, err := decode(name)
		if err != nil {
			return nil, err
		}
		symbol = &SourceSymbol{decodedOwner, decodedName}
	}
	return symbol, nil
}

func (sources *resolvedSources) containing(object types.Object) (*sourceUnit, int, error) {
	file := sources.set.File(object.Pos())
	if file == nil {
		return nil, 0, fmt.Errorf("unowned Go member: %s", object.Name())
	}
	i := sort.Search(len(sources.units), func(i int) bool { return sources.units[i].token.Base() >= file.Base() })
	if i >= len(sources.units) || sources.units[i].token != file {
		return nil, 0, fmt.Errorf("unowned Go member: %s", object.Name())
	}
	unit := sources.units[i]
	i = sort.Search(len(unit.decls), func(i int) bool { return unit.decls[i].End() >= object.Pos() })
	if i >= len(unit.decls) || unit.decls[i].Pos() > object.Pos() {
		return nil, 0, fmt.Errorf("Go member outside a declaration: %s", object.Name())
	}
	return unit, i, nil
}

func BindSources(request BindRequest) (*BindResult, error) {
	sources, err := resolveSources(request)
	if err != nil {
		return nil, err
	}
	units, info, pkg, owned := sources.units, sources.info, sources.pkg, sources.owned
	// Fields and methods are edges to their declaration, never free names.
	containing := func(object types.Object) (SourceSymbol, error) {
		unit, index, err := sources.containing(object)
		if err != nil {
			return SourceSymbol{}, err
		}
		return unit.bound.Declarations[index].Symbols[0], nil
	}
	for _, unit := range units {
		identifiers := map[string]bool{}
		for _, spec := range unit.file.Imports {
			path, err := strconv.Unquote(spec.Path.Value)
			if err != nil {
				return nil, err
			}
			name := ""
			if spec.Name != nil {
				name = spec.Name.Name
			} else if imported, ok := info.Implicits[spec].(*types.PkgName); ok {
				name = imported.Name()
			}
			unit.bound.Imports = append(unit.bound.Imports, SourceImport{path, name})
		}
		var bindingError error
		ast.Inspect(unit.file, func(node ast.Node) bool {
			identifier, ok := node.(*ast.Ident)
			if !ok || bindingError != nil {
				return bindingError == nil
			}
			identifiers[identifier.Name] = true
			object, definition := info.Defs[identifier]
			if used := info.Uses[identifier]; used != nil {
				// An embedded field defines its implicit field and uses its type
				// at the same token. Qualify the type here; selectors of the field
				// still resolve to the enclosing declaration through Defs/Origin.
				object, definition = used, false
			}
			if object == nil {
				// Checked package clauses, blanks, init names and type-switch
				// implicit bindings have no single referable object in types.Info.
				unit.bound.Uses["syntax"]++
				return true
			}
			object = sourceOrigin(object)
			if function, ok := object.(*types.Func); ok && definition && function.Name() == "init" && receiverObject(function) == nil {
				unit.bound.Uses["syntax"]++
				return true
			}
			role, symbol := "local", SourceSymbol{}
			if global, ok := owned[object]; ok {
				role, symbol = "reference", global
				if definition {
					role = "definition"
				}
				if function, ok := object.(*types.Func); ok && receiverObject(function) != nil {
					role = "method"
				}
			} else if imported, ok := object.(*types.PkgName); ok {
				role, symbol = "import", SourceSymbol{"import/" + imported.Imported().Path(), imported.Name()}
			} else if method, ok := object.(*types.Func); ok && method.Type().(*types.Signature).Recv() != nil {
				role = "method"
				if object.Pkg() == pkg {
					symbol, bindingError = containing(object)
				} else if object.Pkg() != nil {
					symbol = SourceSymbol{"import/" + object.Pkg().Path(), object.Name()}
				} else {
					// The predeclared error interface is owned by the Go universe.
					role = "predeclared"
				}
			} else if field, ok := object.(*types.Var); ok && field.IsField() {
				role = "field"
				if object.Pkg() == pkg {
					symbol, bindingError = containing(object)
				} else if object.Pkg() != nil {
					symbol = SourceSymbol{"import/" + object.Pkg().Path(), object.Name()}
				} else {
					bindingError = fmt.Errorf("field without owner: %s", object.Name())
				}
			} else if object.Pkg() == nil {
				role = "predeclared"
			} else if object.Pkg() != pkg {
				role, symbol = "external", SourceSymbol{"import/" + object.Pkg().Path(), object.Name()}
			}
			unit.bound.Uses[role]++
			if role != "local" && role != "predeclared" {
				unit.bound.References = append(unit.bound.References, SourceReference{
					Start: unit.token.Offset(identifier.Pos()), End: unit.token.Offset(identifier.End()),
					Role: role, Name: identifier.Name, Symbol: symbol,
				})
			}
			return true
		})
		if bindingError != nil {
			return nil, bindingError
		}
		for name := range identifiers {
			unit.bound.Identifiers = append(unit.bound.Identifiers, name)
		}
		sort.Strings(unit.bound.Identifiers)
		// Convert all boundaries in one ordered pass, independent of //line
		// directives. Go uses UTF-8 byte positions; PureScript strings use UTF-16.
		positions := []*int{}
		for i := range unit.bound.References {
			ref := &unit.bound.References[i]
			positions = append(positions, &ref.Start, &ref.End)
		}
		for i := range unit.bound.Declarations {
			decl := &unit.bound.Declarations[i]
			positions = append(positions, &decl.Start, &decl.End)
		}
		sort.Slice(positions, func(i, j int) bool { return *positions[i] < *positions[j] })
		previous, units := 0, 0
		for _, position := range positions {
			current := *position
			for _, char := range unit.input.Content[previous:current] {
				units++
				if char > 0xffff {
					units++
				}
			}
			*position, previous = units, current
		}
		sort.Slice(unit.bound.References, func(i, j int) bool { return unit.bound.References[i].Start < unit.bound.References[j].Start })
	}
	result := &BindResult{Format: "gopurs/go-source-v1", Files: []BoundSource{}}
	for _, unit := range units {
		result.Files = append(result.Files, unit.bound)
	}
	return result, nil
}

func Bind(content string) (string, error) {
	var request BindRequest
	if err := json.Unmarshal([]byte(content), &request); err != nil {
		return "", fmt.Errorf("decode Go source binding request: %w", err)
	}
	result, err := BindSources(request)
	if err != nil {
		return "", err
	}
	encoded, err := json.Marshal(result)
	return string(encoded), err
}
