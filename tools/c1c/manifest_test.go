package c1c

import "testing"

const source = `package purescript

import (
    r_ "gopurs/output/gopurs_runtime"
    util "gopurs/output/purescript/p_util"
)

type Thing struct {
    Name string
    hidden int
    util.Base
}

func (t *Thing) Get() string { return t.Name }

func Get_Thing() *Thing { return &Thing{} }

func (t *Thing) set(v string) { t.Name = v }

var Once int
`

func TestExtract(t *testing.T) {
	manifest, err := ExtractSource("Generated.Thing", "Thing.go", []byte(source))
	if err != nil {
		t.Fatal(err)
	}
	if err := manifest.Validate(); err != nil {
		t.Fatal(err)
	}
	if len(manifest.Types) != 1 || manifest.Types[0].Name != "Thing" || manifest.Types[0].Kind != "struct" {
		t.Fatalf("types: %+v", manifest.Types)
	}
	if len(manifest.Types[0].Fields) != 3 {
		t.Fatalf("fields: %+v", manifest.Types[0].Fields)
	}
	if len(manifest.Types[0].Methods) != 2 {
		t.Fatalf("methods: %+v", manifest.Types[0].Methods)
	}
	if len(manifest.Functions) != 1 || manifest.Functions[0].Name != "Get_Thing" {
		t.Fatalf("functions: %+v", manifest.Functions)
	}
	if len(manifest.Imports) != 2 || manifest.Imports[0].Path == "" {
		t.Fatalf("imports: %+v", manifest.Imports)
	}
	if len(manifest.Values) != 1 || manifest.Values[0] != "Once" {
		t.Fatalf("values: %+v", manifest.Values)
	}
}

func TestSignatureExcludesBody(t *testing.T) {
	manifest, err := ExtractSource("M", "M.go", []byte("package p\nfunc F(a int) string { return \"\" }\n"))
	if err != nil {
		t.Fatal(err)
	}
	if manifest.Functions[0].Signature != "func(a int) string" {
		t.Fatalf("signature: %q", manifest.Functions[0].Signature)
	}
}
