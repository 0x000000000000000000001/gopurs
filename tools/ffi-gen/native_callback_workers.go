package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"sort"
	"strconv"
	"strings"
)

// Native callback workers consume the compiler's existing, exact FunctionData
// entry point. This pass never guesses an application's callback contract. It
// clones a worker only when every use of one parameter is an immutable alias,
// a discard, or the complete native/fallback dispatcher emitted by GoFunctions.
// The ordinary worker and its public boxed ABI remain available.
type callbackEdit struct {
	start, end int
	text       string
}

type callbackWorker struct {
	function *ast.FuncDecl
	name     string
	callback *ast.FuncType
	index    int
	edits    []callbackEdit
}

type callbackSource struct {
	source  string
	file    *ast.File
	fset    *token.FileSet
	runtime string
}

func (s *callbackSource) offset(pos token.Pos) int { return s.fset.Position(pos).Offset }
func (s *callbackSource) original(node ast.Node) string {
	return s.source[s.offset(node.Pos()):s.offset(node.End())]
}
func (s *callbackSource) code(node ast.Node) string {
	var out bytes.Buffer
	_ = printer.Fprint(&out, s.fset, node)
	return out.String()
}
func (s *callbackSource) edit(node ast.Node, text string) callbackEdit {
	return callbackEdit{s.offset(node.Pos()), s.offset(node.End()), text}
}
func callbackApplyEdits(source string, edits []callbackEdit) string {
	sort.Slice(edits, func(i, j int) bool { return edits[i].start > edits[j].start })
	for _, edit := range edits {
		source = source[:edit.start] + edit.text + source[edit.end:]
	}
	return source
}
func callbackIdent(expr ast.Expr) *ast.Ident {
	for {
		if p, ok := expr.(*ast.ParenExpr); ok {
			expr = p.X
			continue
		}
		x, _ := expr.(*ast.Ident)
		return x
	}
}
func (s *callbackSource) runtimeSelector(expr ast.Expr, name string) bool {
	sel, ok := expr.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != name {
		return false
	}
	x := callbackIdent(sel.X)
	return x != nil && x.Name == s.runtime && x.Obj == nil
}
func (s *callbackSource) valueType(expr ast.Expr) bool {
	return s.runtimeSelector(expr, "Value")
}
func callbackFields(fields *ast.FieldList) []*ast.Field {
	if fields == nil {
		return nil
	}
	return fields.List
}
func (s *callbackSource) valueResult(fields *ast.FieldList) bool {
	return len(callbackFields(fields)) == 1 && len(fields.List[0].Names) == 0 && s.valueType(fields.List[0].Type)
}
func (s *callbackSource) functionType(ty *ast.FuncType) string {
	copy := *ty
	anonymous := func(fields *ast.FieldList) *ast.FieldList {
		if fields == nil {
			return nil
		}
		result := &ast.FieldList{}
		for _, field := range fields.List {
			count := len(field.Names)
			if count == 0 {
				count = 1
			}
			for i := 0; i < count; i++ {
				result.List = append(result.List, &ast.Field{Type: field.Type})
			}
		}
		return result
	}
	copy.Params, copy.Results = anonymous(ty.Params), anonymous(ty.Results)
	return s.code(&copy) + s.recordAnnotations(ty)
}

// Record labels are producer metadata, not comments that a printer may drop.
// They distinguish equal-looking Go structs and drive their later field ABI.
func (s *callbackSource) recordAnnotations(node ast.Node) string {
	text := s.original(node)
	var markers strings.Builder
	for {
		start := strings.Index(text, "/*gopurs:record-")
		if start < 0 {
			return markers.String()
		}
		text = text[start:]
		end := strings.Index(text, "*/")
		if end < 0 {
			return markers.String() + text
		}
		markers.WriteByte(0)
		markers.WriteString(text[:end+2])
		text = text[end+2:]
	}
}

func (s *callbackSource) callText(name string, args []ast.Expr) string {
	parts := make([]string, len(args))
	for i, argument := range args {
		parts[i] = s.original(argument)
	}
	return name + "(" + strings.Join(parts, ", ") + ")"
}

// Matching printed signatures is sound only for compiler-native structural
// types: local aliases and type parameters can have the same spelling while
// denoting different Go types in the worker and the callback's lexical scope.
func callbackClosedType(typ ast.Expr) bool {
	switch ty := typ.(type) {
	case *ast.Ident:
		if ty.Obj != nil {
			return false
		}
		switch ty.Name {
		case "bool", "int", "int32", "int64", "uint32", "uint64", "float64", "string", "byte", "rune", "any":
			return true
		}
	case *ast.SelectorExpr:
		id := callbackIdent(ty.X)
		return id != nil && id.Obj == nil
	case *ast.StarExpr:
		return callbackClosedType(ty.X)
	case *ast.ArrayType:
		return ty.Len == nil && callbackClosedType(ty.Elt)
	case *ast.StructType:
		for _, field := range ty.Fields.List {
			if len(field.Names) != 1 || field.Tag != nil || !callbackClosedType(field.Type) {
				return false
			}
		}
		return true
	case *ast.FuncType:
		if ty.TypeParams != nil {
			return false
		}
		for _, fields := range []*ast.FieldList{ty.Params, ty.Results} {
			for _, field := range callbackFields(fields) {
				if len(field.Names) > 1 || !callbackClosedType(field.Type) {
					return false
				}
			}
		}
		return true
	case *ast.IndexExpr:
		return callbackClosedType(ty.X) && callbackClosedType(ty.Index)
	case *ast.IndexListExpr:
		if !callbackClosedType(ty.X) {
			return false
		}
		for _, argument := range ty.Indices {
			if !callbackClosedType(argument) {
				return false
			}
		}
		return true
	}
	return false
}
func callbackSingleReturn(body *ast.BlockStmt) *ast.ReturnStmt {
	if body == nil || len(body.List) != 1 {
		return nil
	}
	ret, ok := body.List[0].(*ast.ReturnStmt)
	if !ok || len(ret.Results) != 1 {
		return nil
	}
	return ret
}

// Return the native call and exact callback type of a complete dispatcher.
// No statements, initializers, extra branches, or alternate result shapes may
// disappear together with its fallback.
func (s *callbackSource) dispatcher(call *ast.CallExpr, aliases map[*ast.Object]bool) (*ast.CallExpr, *ast.FuncType) {
	fn, ok := call.Fun.(*ast.FuncLit)
	if !ok || len(call.Args) != 1 || len(callbackFields(fn.Type.Params)) != 2 || !s.valueResult(fn.Type.Results) || len(fn.Body.List) != 2 {
		return nil, nil
	}
	parameters := fn.Type.Params.List
	if len(parameters[0].Names) != 1 || len(parameters[1].Names) != 1 {
		return nil, nil
	}
	nativeType, ok := parameters[0].Type.(*ast.FuncType)
	boolean := callbackIdent(parameters[1].Type)
	if !ok || boolean == nil || boolean.Name != "bool" || boolean.Obj != nil || !s.valueResult(nativeType.Results) || !callbackClosedType(nativeType) || len(nativeType.Params.List) < 1 || len(nativeType.Params.List) > 5 {
		return nil, nil
	}
	data, ok := call.Args[0].(*ast.CallExpr)
	if !ok || len(data.Args) != 1 {
		return nil, nil
	}
	typed, ok := data.Fun.(*ast.IndexExpr)
	if !ok || !s.runtimeSelector(typed.X, "FunctionData") || s.code(typed.Index) != s.code(nativeType) || s.recordAnnotations(typed.Index) != s.recordAnnotations(nativeType) {
		return nil, nil
	}
	callback := callbackIdent(data.Args[0])
	if callback == nil || !aliases[callback.Obj] {
		return nil, nil
	}
	branch, ok := fn.Body.List[0].(*ast.IfStmt)
	if !ok || branch.Init != nil {
		return nil, nil
	}
	condition := callbackIdent(branch.Cond)
	if condition == nil || condition.Obj != parameters[1].Names[0].Obj {
		return nil, nil
	}
	if branch.Else != nil {
		other, ok := branch.Else.(*ast.BlockStmt)
		if !ok || len(other.List) != 0 {
			return nil, nil
		}
	}
	returned := callbackSingleReturn(branch.Body)
	if returned == nil {
		return nil, nil
	}
	nativeCall, ok := returned.Results[0].(*ast.CallExpr)
	if !ok {
		return nil, nil
	}
	callee := callbackIdent(nativeCall.Fun)
	if callee == nil || callee.Obj != parameters[0].Names[0].Obj || len(nativeCall.Args) != len(callbackFields(nativeType.Params)) {
		return nil, nil
	}
	fallback, ok := fn.Body.List[1].(*ast.ReturnStmt)
	if !ok || len(fallback.Results) != 1 {
		return nil, nil
	}
	boxed, ok := fallback.Results[0].(*ast.CallExpr)
	if !ok || len(boxed.Args) != len(nativeCall.Args)+1 {
		return nil, nil
	}
	apply := "Apply"
	if len(nativeCall.Args) != 1 {
		apply += strconv.Itoa(len(nativeCall.Args))
	}
	boxedFn := callbackIdent(boxed.Args[0])
	if !s.runtimeSelector(boxed.Fun, apply) || boxedFn == nil || boxedFn.Obj != callback.Obj {
		return nil, nil
	}
	// Every use of this callback inside the discarded dispatcher is one of
	// the two proven operands. A conversion must not additionally observe it.
	valid := true
	ast.Inspect(call, func(node ast.Node) bool {
		if id, ok := node.(*ast.Ident); ok && aliases[id.Obj] && id != callback && id != boxedFn {
			valid = false
		}
		return true
	})
	if !valid {
		return nil, nil
	}
	replacement := *nativeCall
	replacement.Fun = callback
	return &replacement, nativeType
}

func (s *callbackSource) worker(fn *ast.FuncDecl, index int, name string) *callbackWorker {
	if fn.Recv != nil || fn.Type.TypeParams != nil || !s.valueResult(fn.Type.Results) || index >= len(callbackFields(fn.Type.Params)) {
		return nil
	}
	parameter := fn.Type.Params.List[index]
	if len(parameter.Names) != 1 || !s.valueType(parameter.Type) || parameter.Names[0].Obj == nil {
		return nil
	}
	aliases := map[*ast.Object]bool{parameter.Names[0].Obj: true}
	allowed := map[*ast.Ident]bool{parameter.Names[0]: true}
	typedAliases := []*ast.ValueSpec{}
	// Alias closure is monotone; the use audit below rejects all other writes,
	// shared declarations, escaping uses, and assignments after initialization.
	for changed := true; changed; {
		changed = false
		ast.Inspect(fn.Body, func(node ast.Node) bool {
			var lhs, rhs *ast.Ident
			var spec *ast.ValueSpec
			switch n := node.(type) {
			case *ast.AssignStmt:
				if n.Tok == token.DEFINE && len(n.Lhs) == 1 && len(n.Rhs) == 1 {
					lhs, rhs = callbackIdent(n.Lhs[0]), callbackIdent(n.Rhs[0])
				}
			case *ast.ValueSpec:
				if len(n.Names) == 1 && len(n.Values) == 1 && (n.Type == nil || s.valueType(n.Type)) {
					lhs, rhs, spec = n.Names[0], callbackIdent(n.Values[0]), n
				}
			}
			if lhs != nil && rhs != nil && lhs.Obj != nil && aliases[rhs.Obj] && !aliases[lhs.Obj] {
				aliases[lhs.Obj], allowed[lhs], allowed[rhs] = true, true, true
				if spec != nil && spec.Type != nil {
					typedAliases = append(typedAliases, spec)
				}
				changed = true
			}
			return true
		})
	}
	plan := &callbackWorker{function: fn, name: name, index: index}
	valid := true
	ast.Inspect(fn.Body, func(node ast.Node) bool {
		if call, ok := node.(*ast.CallExpr); ok {
			if native, ty := s.dispatcher(call, aliases); native != nil {
				if plan.callback != nil && s.functionType(plan.callback) != s.functionType(ty) {
					valid = false
				}
				plan.callback = ty
				plan.edits = append(plan.edits, s.edit(call, s.callText(s.original(native.Fun), native.Args)))
				return false
			}
		}
		if assignment, ok := node.(*ast.AssignStmt); ok && assignment.Tok == token.ASSIGN && len(assignment.Lhs) == 1 && len(assignment.Rhs) == 1 {
			lhs, rhs := callbackIdent(assignment.Lhs[0]), callbackIdent(assignment.Rhs[0])
			if lhs != nil && lhs.Name == "_" && rhs != nil && aliases[rhs.Obj] {
				allowed[rhs] = true
			}
		}
		if id, ok := node.(*ast.Ident); ok && aliases[id.Obj] && !allowed[id] {
			valid = false
		}
		return true
	})
	if !valid || plan.callback == nil {
		return nil
	}
	plan.edits = append(plan.edits, s.edit(parameter.Type, s.original(plan.callback)), s.edit(fn.Name, name))
	for _, alias := range typedAliases {
		plan.edits = append(plan.edits, s.edit(alias.Type, s.original(plan.callback)))
	}
	return plan
}

// The getter must consist exactly of a sync.Once cache of a boxed wrapper
// around this worker. All arguments and its result are Value, so bypassing the
// wrapper introduces neither conversion nor eager evaluation of its body.
func (s *callbackSource) getter(fn *ast.FuncDecl, workers map[*ast.Object]*callbackWorker) *callbackWorker {
	if fn.Recv != nil || len(callbackFields(fn.Type.Params)) != 0 || !s.valueResult(fn.Type.Results) || len(fn.Body.List) != 2 {
		return nil
	}
	last, ok := fn.Body.List[1].(*ast.ReturnStmt)
	if !ok || len(last.Results) != 1 {
		return nil
	}
	cache := callbackIdent(last.Results[0])
	first, ok := fn.Body.List[0].(*ast.ExprStmt)
	if cache == nil || cache.Obj == nil || !ok {
		return nil
	}
	once, ok := first.X.(*ast.CallExpr)
	if !ok || len(once.Args) != 1 {
		return nil
	}
	do, ok := once.Fun.(*ast.SelectorExpr)
	if !ok || do.Sel.Name != "Do" {
		return nil
	}
	onceID := callbackIdent(do.X)
	if onceID == nil || onceID.Obj == nil {
		return nil
	}
	onceDecl, ok := onceID.Obj.Decl.(*ast.ValueSpec)
	if !ok || len(onceDecl.Names) != 1 || len(onceDecl.Values) != 0 {
		return nil
	}
	onceType, ok := onceDecl.Type.(*ast.SelectorExpr)
	if !ok || onceType.Sel.Name != "Once" {
		return nil
	}
	packageID := callbackIdent(onceType.X)
	if packageID == nil || packageID.Obj != nil || packageID.Name != "sync" {
		return nil
	}
	syncImport := false
	for _, spec := range s.file.Imports {
		path, _ := strconv.Unquote(spec.Path.Value)
		if path == "sync" && (spec.Name == nil || spec.Name.Name == "sync") {
			syncImport = true
		}
	}
	if !syncImport {
		return nil
	}
	init, ok := once.Args[0].(*ast.FuncLit)
	if !ok || len(callbackFields(init.Type.Params)) != 0 || init.Type.Results != nil || len(init.Body.List) != 1 {
		return nil
	}
	assignment, ok := init.Body.List[0].(*ast.AssignStmt)
	if !ok || assignment.Tok != token.ASSIGN || len(assignment.Lhs) != 1 || len(assignment.Rhs) != 1 {
		return nil
	}
	lhs := callbackIdent(assignment.Lhs[0])
	boxed, ok := assignment.Rhs[0].(*ast.CallExpr)
	if lhs == nil || lhs.Obj != cache.Obj || !ok || len(boxed.Args) != 1 {
		return nil
	}
	wrapper, ok := boxed.Args[0].(*ast.FuncLit)
	if !ok || !s.valueResult(wrapper.Type.Results) {
		return nil
	}
	returned := callbackSingleReturn(wrapper.Body)
	if returned == nil {
		return nil
	}
	call, ok := returned.Results[0].(*ast.CallExpr)
	if !ok {
		return nil
	}
	target := callbackIdent(call.Fun)
	if target == nil || workers[target.Obj] == nil {
		return nil
	}
	worker := workers[target.Obj]
	parameters := callbackFields(wrapper.Type.Params)
	if len(parameters) != len(callbackFields(worker.function.Type.Params)) || len(call.Args) != len(parameters) {
		return nil
	}
	constructor := "Func"
	if len(parameters) != 1 {
		constructor += strconv.Itoa(len(parameters))
	}
	if !s.runtimeSelector(boxed.Fun, constructor) {
		return nil
	}
	for i, field := range parameters {
		arg := callbackIdent(call.Args[i])
		if len(field.Names) != 1 || !s.valueType(field.Type) || arg == nil || arg.Obj != field.Names[0].Obj || !s.valueType(worker.function.Type.Params.List[i].Type) {
			return nil
		}
	}
	return worker
}

func (s *callbackSource) nativeArgument(expr ast.Expr, worker *callbackWorker) *ast.FuncLit {
	call, ok := expr.(*ast.CallExpr)
	if !ok || !s.runtimeSelector(call.Fun, "WithFunctionData") || len(call.Args) != 2 {
		return nil
	}
	boxed, ok := call.Args[0].(*ast.CallExpr)
	if !ok || len(boxed.Args) != 1 {
		return nil
	}
	old, ok := boxed.Args[0].(*ast.FuncLit)
	if !ok || !s.valueResult(old.Type.Results) {
		return nil
	}
	constructor := "Func"
	if len(callbackFields(old.Type.Params)) != 1 {
		constructor += strconv.Itoa(len(callbackFields(old.Type.Params)))
	}
	if !s.runtimeSelector(boxed.Fun, constructor) {
		return nil
	}
	native, ok := call.Args[1].(*ast.FuncLit)
	if !ok || !callbackClosedType(native.Type) || s.functionType(native.Type) != s.functionType(worker.callback) {
		return nil
	}
	return native
}

// OptimizeNativeCallbacks keeps unchanged source bytes intact. It only emits
// clones actually selected by a known literal callback. The callback literal
// may move inside a partial application's returned closure: creating a Go
// closure evaluates no user code, and its lexical captures remain the same.
// Other supplied arguments are currently rejected for partial calls, avoiding
// any change to their evaluation order or the time at which values are read.
func OptimizeNativeCallbacks(source string) (string, error) {
	set := token.NewFileSet()
	file, err := parser.ParseFile(set, "generated.go", source, 0)
	if err != nil {
		return "", err
	}
	s := callbackSource{source: source, file: file, fset: set}
	for _, imported := range file.Imports {
		path, _ := strconv.Unquote(imported.Path.Value)
		if path == "gopurs/output/gopurs_runtime" {
			s.runtime = "gopurs_runtime"
			if imported.Name != nil {
				s.runtime = imported.Name.Name
			}
		}
	}
	if s.runtime == "" || s.runtime == "." || s.runtime == "_" {
		return source, nil
	}
	workers := map[*ast.Object]*callbackWorker{}
	names := map[string]bool{}
	ast.Inspect(file, func(node ast.Node) bool {
		if id, ok := node.(*ast.Ident); ok {
			names[id.Name] = true
		}
		return true
	})
	for _, declaration := range file.Decls {
		fn, ok := declaration.(*ast.FuncDecl)
		if !ok || fn.Name.Obj == nil {
			continue
		}
		name := "GopursNativeCallback_" + fn.Name.Name
		for n := 1; names[name]; n++ {
			name = "GopursNativeCallback_" + fn.Name.Name + "_" + strconv.Itoa(n)
		}
		// One independent leading callback keeps worker proliferation bounded.
		if plan := s.worker(fn, 0, name); plan != nil {
			workers[fn.Name.Obj], names[name] = plan, true
		}
	}
	getters := map[*ast.Object]*callbackWorker{}
	for _, declaration := range file.Decls {
		if fn, ok := declaration.(*ast.FuncDecl); ok {
			if worker := s.getter(fn, workers); worker != nil {
				getters[fn.Name.Obj] = worker
			}
		}
	}
	edits := []callbackEdit{}
	selected := map[*callbackWorker]bool{}
	ast.Inspect(file, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}
		if id := callbackIdent(call.Fun); id != nil {
			if worker := workers[id.Obj]; worker != nil && len(call.Args) == len(worker.function.Type.Params.List) {
				if native := s.nativeArgument(call.Args[0], worker); native != nil {
					args := append([]ast.Expr(nil), call.Args...)
					args[0] = native
					edits = append(edits, s.edit(call, s.callText(worker.name, args)))
					selected[worker] = true
					return false
				}
			}
		}
		if !s.runtimeSelector(call.Fun, "Apply") || len(call.Args) != 2 {
			return true
		}
		getter, ok := call.Args[0].(*ast.CallExpr)
		if !ok || len(getter.Args) != 0 {
			return true
		}
		id := callbackIdent(getter.Fun)
		if id == nil || getters[id.Obj] == nil {
			return true
		}
		worker := getters[id.Obj]
		native := s.nativeArgument(call.Args[1], worker)
		remaining := len(worker.function.Type.Params.List) - 1
		if native == nil || remaining < 1 || remaining > 5 {
			return true
		}
		parameters := []string{}
		args := []string{s.original(native)}
		for i := 0; i < remaining; i++ {
			name := fmt.Sprintf("__gopurs_callback_arg_%d", i)
			for names[name] {
				name += "_"
			}
			parameters = append(parameters, name+" "+s.original(worker.function.Type.Params.List[i+1].Type))
			args = append(args, name)
		}
		constructor := "Func"
		if remaining != 1 {
			constructor += strconv.Itoa(remaining)
		}
		replacement := s.runtime + "." + constructor + "(func(" + strings.Join(parameters, ", ") + ") " +
			s.original(worker.function.Type.Results.List[0].Type) + " { return " + worker.name + "(" + strings.Join(args, ", ") + ") })"
		edits = append(edits, s.edit(call, replacement))
		selected[worker] = true
		return false
	})
	if len(selected) == 0 {
		return source, nil
	}
	result := callbackApplyEdits(source, edits)
	// Source order makes output stable independently of Go map iteration.
	for _, declaration := range file.Decls {
		fn, ok := declaration.(*ast.FuncDecl)
		if !ok {
			continue
		}
		worker := workers[fn.Name.Obj]
		if !selected[worker] {
			continue
		}
		start, end := s.offset(fn.Pos()), s.offset(fn.End())
		localEdits := make([]callbackEdit, len(worker.edits))
		for i, edit := range worker.edits {
			localEdits[i] = callbackEdit{edit.start - start, edit.end - start, edit.text}
		}
		result += "\n\n" + callbackApplyEdits(source[start:end], localEdits) + "\n"
	}
	return result, nil
}
