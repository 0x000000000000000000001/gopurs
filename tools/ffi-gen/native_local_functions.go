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

// OptimizeGeneratedSource runs only on compiler-generated modules, before
// source binding and package planning. User FFI keeps its representation and
// passes through ShareRecordShapes separately. Keeping the analysis here also
// covers the opaque fragments emitted by the PureScript GoAst printer.
func OptimizeGeneratedSource(prefix, source string) (string, error) {
	// Callback specialization runs first: its native workers expose additional
	// local functions whose Value wrappers are no longer observed.
	output, err := OptimizeNativeCallbacks(source)
	if err != nil {
		return "", err
	}
	output, err = optimizeNativeLocalFunctions(output)
	if err != nil {
		return "", err
	}
	return ShareRecordShapes(prefix, output)
}

type localFunctionCandidate struct {
	object *ast.Object
	assign *ast.AssignStmt
	block  *ast.BlockStmt
	call   *ast.CallExpr
	value  *ast.FuncLit
	arity  int
}

type localFunctionEdit struct {
	start, end int
	text       string
}

// A wrapper can become an ordinary Go function only when all observations of
// its lexical binding, and every alias, are exact saturated runtime applies.
// Calls from a captured closure are allowed; Go then owns the closure lifetime.
// Returning, storing, passing, inspecting, partially applying or reassigning
// the function rejects the whole alias group. No ownership claim is inferred
// from PureScript usage annotations, and runtime forceEscape is unchanged.
func optimizeNativeLocalFunctions(source string) (string, error) {
	set := token.NewFileSet()
	file, err := parser.ParseFile(set, "generated.go", source, parser.ParseComments)
	if err != nil {
		return "", err
	}
	runtimeAliases := map[string]bool{}
	for _, spec := range file.Imports {
		path, _ := strconv.Unquote(spec.Path.Value)
		if path != "gopurs/output/gopurs_runtime" {
			continue
		}
		alias := "gopurs_runtime"
		if spec.Name != nil {
			alias = spec.Name.Name
		}
		if alias != "." && alias != "_" {
			runtimeAliases[alias] = true
		}
	}
	if len(runtimeAliases) == 0 {
		return source, nil
	}
	runtimeCall := func(expr ast.Expr, stem string) (*ast.CallExpr, int, bool) {
		call, ok := expr.(*ast.CallExpr)
		if !ok || call.Ellipsis.IsValid() {
			return nil, 0, false
		}
		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return nil, 0, false
		}
		receiver, ok := selector.X.(*ast.Ident)
		if !ok || !runtimeAliases[receiver.Name] || receiver.Obj != nil {
			return nil, 0, false
		}
		name := selector.Sel.Name
		if name == stem {
			return call, 1, true
		}
		if !strings.HasPrefix(name, stem) {
			return nil, 0, false
		}
		arity, err := strconv.Atoi(strings.TrimPrefix(name, stem))
		return call, arity, err == nil && arity >= 2 && arity <= 11 && name == stem+strconv.Itoa(arity)
	}

	parents := map[ast.Node]ast.Node{}
	uses := map[*ast.Object][]*ast.Ident{}
	assignments := map[*ast.Object]int{}
	var stack []ast.Node
	ast.Inspect(file, func(node ast.Node) bool {
		if node == nil {
			stack = stack[:len(stack)-1]
			return true
		}
		if len(stack) > 0 {
			parents[node] = stack[len(stack)-1]
		}
		stack = append(stack, node)
		if ident, ok := node.(*ast.Ident); ok && ident.Obj != nil {
			uses[ident.Obj] = append(uses[ident.Obj], ident)
		}
		if assign, ok := node.(*ast.AssignStmt); ok {
			for _, left := range assign.Lhs {
				if ident, ok := left.(*ast.Ident); ok && ident.Obj != nil {
					assignments[ident.Obj]++
				}
			}
		}
		return true
	})
	functionOf := func(node ast.Node) ast.Node {
		for node != nil {
			switch node.(type) {
			case *ast.FuncDecl, *ast.FuncLit:
				return node
			}
			node = parents[node]
		}
		return nil
	}
	labels := map[ast.Node]map[string]*ast.LabeledStmt{}
	jumps := map[ast.Node][]*ast.BranchStmt{}
	ast.Inspect(file, func(node ast.Node) bool {
		switch value := node.(type) {
		case *ast.LabeledStmt:
			function := functionOf(value)
			if labels[function] == nil {
				labels[function] = map[string]*ast.LabeledStmt{}
			}
			labels[function][value.Label.Name] = value
		case *ast.BranchStmt:
			if value.Tok == token.GOTO {
				function := functionOf(value)
				jumps[function] = append(jumps[function], value)
			}
		}
		return true
	})

	var candidates []localFunctionCandidate
	for _, declaration := range file.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || function.Body == nil {
			continue
		}
		ast.Inspect(function.Body, func(node ast.Node) bool {
			assign, ok := node.(*ast.AssignStmt)
			if !ok || len(assign.Lhs) != 1 || len(assign.Rhs) != 1 {
				return true
			}
			ident, ok := assign.Lhs[0].(*ast.Ident)
			// Package bindings may be observed in other source units. Only local
			// declarations have a complete set of uses in this syntax tree.
			if !ok || ident.Obj == nil || ident.Name == "_" || ident.Obj.Pos() < function.Body.Pos() || ident.Obj.Pos() >= function.Body.End() {
				return true
			}
			// Parameters, named results and range bindings cannot change type.
			// A declaration may precede a loop containing the assignment; each
			// observation must still be dominated by that assignment below.
			switch value := ident.Obj.Decl.(type) {
			case *ast.AssignStmt:
				if value != assign || value.Tok != token.DEFINE {
					return true
				}
			case *ast.ValueSpec:
			default:
				return true
			}
			block, ok := parents[assign].(*ast.BlockStmt)
			if !ok {
				return true
			}
			// Generated pattern matches use forward jumps to join labels. A
			// jump leaving this block is harmless because observations must stay
			// inside it. Reject any entry after the assignment that can bypass
			// initialization; labels belong to their own Go function scope.
			scope := functionOf(assign)
			for _, jump := range jumps[scope] {
				target := labels[scope][jump.Label.Name]
				if target != nil && target.Pos() >= assign.End() && target.Pos() < block.End() &&
					(jump.Pos() < assign.End() || jump.Pos() >= block.End()) {
					return true
				}
			}
			call, arity, ok := runtimeCall(assign.Rhs[0], "Func")
			if !ok || len(call.Args) != 1 {
				return true
			}
			value, ok := call.Args[0].(*ast.FuncLit)
			if ok {
				candidates = append(candidates, localFunctionCandidate{ident.Obj, assign, block, call, value, arity})
			}
			return true
		})
	}

	accepted := map[*ast.Object]localFunctionCandidate{}
	calls := map[*ast.CallExpr]bool{}
	typeAvailableAt := func(functionType *ast.FuncType, declaration *ast.ValueSpec) bool {
		available := true
		ast.Inspect(functionType, func(node ast.Node) bool {
			if !available {
				return false
			}
			ident, ok := node.(*ast.Ident)
			if !ok || ident.Obj == nil || ident.Obj.Kind != ast.Typ {
				return available
			}
			// Qualified and package types remain visible. A local type alias
			// used by the lambda may not be in scope at the earlier slot
			// declaration where its native function type would be emitted.
			var scope *ast.BlockStmt
			for owner, _ := ident.Obj.Decl.(ast.Node); owner != nil; owner = parents[owner] {
				if block, ok := owner.(*ast.BlockStmt); ok {
					scope = block
					break
				}
			}
			if scope != nil {
				available = false
				if ident.Obj.Pos() <= declaration.Pos() {
					for owner := ast.Node(declaration); owner != nil; owner = parents[owner] {
						if owner == scope {
							available = true
							break
						}
					}
				}
			}
			return available
		})
		return available
	}
	dominated := func(ident *ast.Ident, candidate localFunctionCandidate) bool {
		// Self-references inside the wrapper body execute after construction.
		if ident.Pos() >= candidate.value.Body.Pos() && ident.Pos() < candidate.value.Body.End() {
			return true
		}
		// The use must belong to a later statement in the assignment's block.
		// In particular, a conditional or loop assignment does not dominate a
		// use outside its block, even when its source position is later.
		for node := ast.Node(ident); node != nil; node = parents[node] {
			if parents[node] == candidate.block {
				return node.Pos() >= candidate.assign.End()
			}
		}
		return false
	}
	for _, candidate := range candidates {
		group := map[*ast.Object]bool{candidate.object: true}
		queue := []*ast.Object{candidate.object}
		localCalls := map[*ast.CallExpr]bool{}
		valid := true
		for len(queue) > 0 && valid {
			object := queue[0]
			queue = queue[1:]
			if assignments[object] != 1 {
				valid = false
				break
			}
			if spec, ok := object.Decl.(*ast.ValueSpec); ok && (len(spec.Names) != 1 || len(spec.Values) != 0 || spec.Type == nil || !typeAvailableAt(candidate.value.Type, spec)) {
				valid = false
				break
			}
			for _, ident := range uses[object] {
				switch parent := parents[ident].(type) {
				case *ast.ValueSpec:
					valid = len(parent.Names) == 1 && parent.Names[0] == ident
				case *ast.AssignStmt:
					if len(parent.Lhs) != 1 || len(parent.Rhs) != 1 {
						valid = false
						break
					}
					if parent.Lhs[0] == ident {
						break
					}
					alias, ok := parent.Lhs[0].(*ast.Ident)
					if parent.Rhs[0] != ident || !ok {
						valid = false
						break
					}
					if alias.Name == "_" {
						break
					}
					if !dominated(ident, candidate) {
						valid = false
						break
					}
					if parent.Tok != token.DEFINE || alias.Obj == nil {
						valid = false
						break
					}
					if !group[alias.Obj] {
						group[alias.Obj] = true
						queue = append(queue, alias.Obj)
					}
				case *ast.CallExpr:
					call, arity, ok := runtimeCall(parent, "Apply")
					valid = ok && arity == candidate.arity && len(call.Args) == arity+1 && call.Args[0] == ident
					valid = valid && dominated(ident, candidate)
					if valid {
						localCalls[call] = true
					}
				default:
					valid = false
				}
				if !valid {
					break
				}
			}
		}
		if valid {
			for object := range group {
				accepted[object] = candidate
			}
			for call := range localCalls {
				calls[call] = true
			}
		}
	}

	// Make edits from resolved AST positions, preserving unrelated source bytes.
	// Wrapper delimiters and Apply prefixes do not overlap nested call edits.
	var edits []localFunctionEdit
	edit := func(start, end token.Pos, text string) {
		edits = append(edits, localFunctionEdit{set.Position(start).Offset, set.Position(end).Offset, text})
	}
	for _, candidate := range candidates {
		if _, ok := accepted[candidate.object]; ok {
			edit(candidate.call.Pos(), candidate.value.Pos(), "")
			edit(candidate.value.End(), candidate.call.End(), "")
		}
	}
	for object, candidate := range accepted {
		if spec, ok := object.Decl.(*ast.ValueSpec); ok {
			var out bytes.Buffer
			if err := printer.Fprint(&out, set, candidate.value.Type); err != nil {
				return "", err
			}
			edit(spec.Type.Pos(), spec.Type.End(), out.String())
		}
	}
	for call := range calls {
		edit(call.Fun.Pos(), call.Fun.End(), call.Args[0].(*ast.Ident).Name)
		edit(call.Lparen+1, call.Args[1].Pos(), "")
	}
	sort.Slice(edits, func(i, j int) bool { return edits[i].start < edits[j].start })
	var out strings.Builder
	previous := 0
	for _, edit := range edits {
		if edit.start < previous || edit.end < edit.start {
			return "", fmt.Errorf("overlapping native local function transformations")
		}
		out.WriteString(source[previous:edit.start])
		out.WriteString(edit.text)
		previous = edit.end
	}
	if len(edits) == 0 {
		return source, nil
	}
	out.WriteString(source[previous:])
	return out.String(), nil
}
