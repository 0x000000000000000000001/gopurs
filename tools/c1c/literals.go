package c1c

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"sort"
)

type bodyLiteral struct {
	kind       token.Token
	text       string
	start, end int
}

type bodyRange struct {
	name       string
	start, end int
}
type literalSource struct {
	tokens   []bodyLiteral
	bodies   []bodyRange
	literals map[string][]bodyLiteral
}

func scanLiteralSource(source string) (*literalSource, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "literals.go", source, 0)
	if err != nil {
		return nil, err
	}
	result := &literalSource{literals: map[string][]bodyLiteral{}}
	for _, declaration := range file.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || function.Body == nil {
			continue
		}
		if function.Recv != nil || function.Name.Name == "init" {
			return nil, fmt.Errorf("method or init body")
		}
		name := function.Name.Name
		if _, exists := result.literals[name]; exists {
			return nil, fmt.Errorf("duplicate function %s", name)
		}
		result.literals[name] = nil
		result.bodies = append(result.bodies, bodyRange{name, fset.Position(function.Body.Pos()).Offset, fset.Position(function.Body.End()).Offset})
	}
	var scan scanner.Scanner
	scanFile := token.NewFileSet().AddFile("tokens.go", -1, len(source))
	// Record annotations and compiler directives affect the planner even
	// though the Go parser treats them as comments. Include them in the proof.
	scan.Init(scanFile, []byte(source), nil, scanner.ScanComments)
	for {
		position, kind, literal := scan.Scan()
		if kind == token.EOF {
			break
		}
		start := scanFile.Offset(position)
		item := bodyLiteral{kind, literal, start, start + len(literal)}
		result.tokens = append(result.tokens, item)
		if kind.IsLiteral() && kind != token.IDENT {
			if name := result.bodyAt(start); name != "" {
				result.literals[name] = append(result.literals[name], item)
			}
		}
	}
	return result, nil
}

func (source *literalSource) bodyAt(offset int) string {
	i := sort.Search(len(source.bodies), func(i int) bool { return source.bodies[i].end > offset })
	if i < len(source.bodies) && source.bodies[i].start <= offset {
		return source.bodies[i].name
	}
	return ""
}

// LiteralChanges proves that the raw producer changed only literal tokens
// inside function bodies. Published Go may have planner-owned field names and
// qualifications; preserving it and applying these literals avoids recreating
// those type-checked transformations without their original type information.
type LiteralChanges struct{ old, new *literalSource }

func CompareBodyLiterals(previous, current string) (*LiteralChanges, error) {
	old, err := scanLiteralSource(previous)
	if err != nil {
		return nil, err
	}
	new, err := scanLiteralSource(current)
	if err != nil {
		return nil, err
	}
	if len(old.tokens) != len(new.tokens) {
		return nil, fmt.Errorf("source token structure changed")
	}
	for i, before := range old.tokens {
		after := new.tokens[i]
		if before.kind != after.kind {
			return nil, fmt.Errorf("source token kind changed")
		}
		if before.text == after.text {
			continue
		}
		body := old.bodyAt(before.start)
		if !before.kind.IsLiteral() || before.kind == token.IDENT || body == "" || body != new.bodyAt(after.start) {
			return nil, fmt.Errorf("source changed outside body literals")
		}
	}
	return &LiteralChanges{old, new}, nil
}

// Apply checks the complete literal sequence of every published function
// against the verified raw source before applying any replacements.
func (changes *LiteralChanges) Apply(published string) (string, error) {
	section, err := scanLiteralSource(published)
	if err != nil {
		return "", err
	}
	edits := []bodyLiteral{}
	for name, literals := range section.literals {
		old, exists := changes.old.literals[name]
		new := changes.new.literals[name]
		if !exists || len(old) != len(literals) || len(old) != len(new) {
			return "", fmt.Errorf("published literal sequence differs: %s", name)
		}
		for i, actual := range literals {
			if actual.kind != old[i].kind || actual.text != old[i].text {
				return "", fmt.Errorf("published literal differs: %s", name)
			}
			if actual.text != new[i].text {
				actual.text = new[i].text
				edits = append(edits, actual)
			}
		}
	}
	sort.Slice(edits, func(i, j int) bool { return edits[i].start > edits[j].start })
	for _, edit := range edits {
		published = published[:edit.start] + edit.text + published[edit.end:]
	}
	return published, nil
}
