package main

import (
	"encoding/json"
	"strings"
	"testing"
	"unicode/utf16"
)

func bindingFixture() BindRequest {
	return BindRequest{PackagePath: "example/purescript", GoVersion: "go1.22", Arch: "arm64", Files: []SourceInput{
		{Owner: "ffi/Provider", Path: "Provider.go", Content: `package purescript
type Box[T any] struct { Key T }
func (renamed *Box[T]) Read() T { return renamed.Key }
type IntMap map[int]int
const Key = 3
func Produce() int { return 21 }
type Readable[T any] interface { Read() T }
`},
		{Owner: "generated/Consumer", Path: "Consumer.go", Content: `package purescript
// 😀 Produce Key: never a source reference.
func Use() int {
  value := Box[int]{Key: Produce()}
  entries := IntMap{Key: 10}
  callback := func(Produce func() int) int { return Produce() }
  var boxed any = value
  { Key := 99; _ = Box[int]{Key: Key} }
  return boxed.(Box[int]).Key + entries[Key] + callback(Produce) - 10
}
func ReadBox(value Readable[int]) int { return value.Read() }
`},
	}}
}

func TestBindSourceScopes(t *testing.T) {
	request := bindingFixture()
	result, err := BindSources(request)
	if err != nil {
		t.Fatal(err)
	}
	consumer := result.Files[0]
	if consumer.Path != "Consumer.go" || consumer.Text != request.Files[1].Content {
		t.Fatal("source text or canonical order changed")
	}
	counts := map[string]int{}
	text := utf16.Encode([]rune(consumer.Text))
	for _, ref := range consumer.References {
		if string(utf16.Decode(text[ref.Start:ref.End])) != ref.Name {
			t.Fatalf("nonportable source position: %+v", ref)
		}
		counts[ref.Role+":"+ref.Name]++
		if ref.Role == "reference" && ref.Symbol.Owner != "ffi/Provider" {
			t.Fatalf("wrong owner: %+v", ref)
		}
	}
	for role, count := range map[string]int{"reference:Produce": 2, "reference:Key": 2, "field:Key": 3, "reference:Box": 3, "reference:IntMap": 1} {
		if counts[role] != count {
			t.Errorf("%s: got %d, want %d", role, counts[role], count)
		}
	}
	method := result.Files[1].Declarations[1].Symbols[0]
	if method != (SourceSymbol{"ffi/Provider", "Box.Read"}) {
		t.Fatalf("method identity depends on receiver spelling: %+v", method)
	}
	interfaceMethod := false
	for _, reference := range consumer.References {
		if reference.Role == "method" && reference.Name == "Read" {
			interfaceMethod = reference.Symbol == (SourceSymbol{"ffi/Provider", "Readable"})
		}
	}
	if !interfaceMethod {
		t.Fatal("interface method has no owning interface")
	}
	encoded, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	transport, err := Bind(string(encoded))
	if err != nil {
		t.Fatal(err)
	}
	expected, _ := json.Marshal(result)
	if transport != string(expected) {
		t.Fatal("JSON transport changed source bindings")
	}
}

func TestBindSourceRejectsUnknownsAndMissingInitIdentity(t *testing.T) {
	for _, source := range []string{
		"package purescript\nfunc F() { Missing() }",
		"package purescript\nfunc init() {}",
		"package purescript\nimport \"missing/export\"\nvar _ = export.Name",
	} {
		request := bindingFixture()
		request.Files = []SourceInput{{Owner: "generated/Probe", Path: "probe.go", Content: source}}
		if _, err := BindSources(request); err == nil {
			t.Fatalf("accepted an unresolved source: %s", source)
		}
	}
	request := bindingFixture()
	request.Files = []SourceInput{{Owner: "generated/Probe", Path: "probe.go", Content: "package purescript\nfunc init() {}",
		Initializers: []SourceSymbol{{"generated/Probe", "register/Constructor_Probe_Box"}}}}
	result, err := BindSources(request)
	if err != nil || result.Files[0].Declarations[0].Symbols[0].Name != "register/Constructor_Probe_Box" {
		t.Fatalf("lost explicit registration identity: %+v, %v", result, err)
	}
	for _, reference := range result.Files[0].References {
		if reference.Name == "init" {
			t.Fatalf("init is not a qualifiable package name: %+v", reference)
		}
	}
	request.Files[0].Content = strings.ReplaceAll(request.Files[0].Content, "func init", "//line foreign.go:400\nfunc init")
	if _, err := BindSources(request); err != nil {
		t.Fatal(err)
	}
}

func TestEmbeddedFieldRetainsTypeReference(t *testing.T) {
	request := bindingFixture()
	request.Files[1].Content = `package purescript
type Wrapped struct { Box[int] }
func ReadWrapped(value Wrapped) int { return value.Box.Read() }
`
	result, err := BindSources(request)
	if err != nil {
		t.Fatal(err)
	}
	foundType, foundField := false, false
	for _, reference := range result.Files[0].References {
		if reference.Name != "Box" {
			continue
		}
		if reference.Role == "reference" && reference.Symbol == (SourceSymbol{"ffi/Provider", "Box"}) {
			foundType = true
		}
		if reference.Role == "field" && reference.Symbol == (SourceSymbol{"generated/Consumer", "Wrapped"}) {
			foundField = true
		}
	}
	if !foundType || !foundField {
		t.Fatalf("embedded type and field must retain distinct targets: %+v", result.Files[0].References)
	}
}
