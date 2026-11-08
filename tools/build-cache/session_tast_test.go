package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const emptyTypedModule = `{"typeTable":[],"dataDecls":[],"classDecls":[]}`

func TestTypedModuleMetadataShapes(t *testing.T) {
	for _, test := range []struct {
		name   string
		source string
		valid  bool
	}{
		{"empty arrays", emptyTypedModule, true},
		{"populated arrays", `{"typeTable":["Int"],"dataDecls":[{}],"classDecls":[{}]}`, true},
		{"whitespace", "{\n\"typeTable\" : [ ], \"dataDecls\" : [], \"classDecls\" : [] }", true},
		{"stock", `{"moduleName":["Main"],"decls":[]}`, false},
		{"missing type table", `{"dataDecls":[],"classDecls":[]}`, false},
		{"null data declarations", `{"typeTable":[],"dataDecls":null,"classDecls":[]}`, false},
		{"class declarations object", `{"typeTable":[],"dataDecls":[],"classDecls":{}}`, false},
		{"wrong key casing", `{"TypeTable":[],"dataDecls":[],"classDecls":[]}`, false},
		{"nested fields", `{"metadata":` + emptyTypedModule + `}`, false},
		{"array root", `[]`, false},
		{"malformed payload", `{"typeTable":[broken],"dataDecls":[],"classDecls":[]}`, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			err := validateTypedModule("Main", "/project/output/Main/corefn.json", []byte(test.source))
			if (err == nil) != test.valid {
				t.Fatalf("valid=%v, error=%v", test.valid, err)
			}
			if err != nil && (!strings.Contains(err.Error(), "Main") || !strings.Contains(err.Error(), "/project/output/Main/corefn.json")) {
				t.Fatalf("diagnostic omits its source: %v", err)
			}
		})
	}
}

func TestSessionValidatesFrozenTypedInputs(t *testing.T) {
	for _, layout := range []string{"", "packed-v1"} {
		for _, scenario := range []string{"empty", "stock", "mixed"} {
			t.Run(layout+"/"+scenario, func(t *testing.T) {
				f := newFixture(t)
				spec := f.req.Spec
				spec.Output = f.s.root
				spec.Inputs = []Input{{Name: "tast", Kind: "tast", Path: spec.Output}}
				spec.Options = append(spec.Options, Setting{Name: "input-format", Value: "tast-v1"}, Setting{Name: "input-layout", Value: layout})
				put(t, filepath.Join(spec.Output, "Main/corefn.json"), emptyTypedModule)
				invalid := filepath.Join(spec.Output, "Main/corefn.json")
				if scenario == "mixed" {
					invalid = filepath.Join(spec.Output, "Dependency/corefn.json")
				}
				if scenario != "empty" {
					put(t, invalid, `{"moduleName":["Stock"],"decls":[]}`)
				}
				var result SessionResponse
				err := beginSession(spec, false, &result)
				if scenario == "empty" {
					if err != nil || result.Status != "miss" || result.Count != 1 {
						t.Fatalf("empty TAST arrays rejected: %+v, %v", result, err)
					}
					return
				}
				if err == nil {
					t.Fatal("stock CoreFn accepted")
				}
				for _, expected := range []string{invalid, "TAST", "typed PureScript fork", "clean frontend output", "go.mod/go.sum", "node_modules/.bin"} {
					if !strings.Contains(err.Error(), expected) {
						t.Fatalf("missing %q in diagnostic: %v", expected, err)
					}
				}
				stages, err := os.ReadDir(filepath.Join(spec.Output, ".gopurs-cache/stages"))
				if err != nil || len(stages) != 0 {
					t.Fatalf("rejected input left a partial stage: %v, %v", stages, err)
				}
			})
		}
	}
}

func TestTypedInputFormatInvalidatesLegacyCache(t *testing.T) {
	spec, result, state := driverFixture(t)
	modules := []SourceModule{{Name: "Main", Path: "src/Main.purs"}}
	finishDriver(t, &result, &state, modules, false)
	var cached SessionResponse
	if err := beginSession(spec, false, &cached); err != nil || cached.Status != "hit" {
		t.Fatalf("legacy fixture was not cached: %+v, %v", cached, err)
	}
	pointerPath := filepath.Join(spec.Output, ".gopurs-cache/current.json")
	pointer, err := os.ReadFile(pointerPath)
	if err != nil {
		t.Fatal(err)
	}
	spec.Options = append(spec.Options, Setting{Name: "input-format", Value: "tast-v1"})
	var rejected SessionResponse
	if err := beginSession(spec, false, &rejected); err == nil || !strings.Contains(err.Error(), "TAST") {
		t.Fatalf("legacy hit bypassed TAST validation: %+v, %v", rejected, err)
	}
	current, err := os.ReadFile(pointerPath)
	if err != nil || string(current) != string(pointer) {
		t.Fatalf("failed validation replaced the committed cache: %v", err)
	}
	put(t, filepath.Join(spec.Output, "Main/corefn.json"), emptyTypedModule)
	if err := beginSession(spec, false, &result); err != nil || result.Status != "miss" {
		t.Fatalf("corrected typed input did not recover: %+v, %v", result, err)
	}
	data, err := os.ReadFile(filepath.Join(result.Stage, "session.json"))
	if err != nil || strictJSON(data, &state) != nil {
		t.Fatalf("invalid recovered session: %v", err)
	}
	finishDriver(t, &result, &state, modules, false)
	if err := beginSession(spec, false, &cached); err != nil || cached.Status != "hit" {
		t.Fatalf("validated typed inputs failed to hit cache: %+v, %v", cached, err)
	}
	pointer, err = os.ReadFile(pointerPath)
	if err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(spec.Output, "Main/corefn.json"), `{"moduleName":["Main"],"decls":[]}`)
	if err := beginSession(spec, false, &rejected); err == nil || !strings.Contains(err.Error(), "TAST") {
		t.Fatalf("validated cache accepted replaced stock input: %+v, %v", rejected, err)
	}
	current, err = os.ReadFile(pointerPath)
	if err != nil || string(current) != string(pointer) {
		t.Fatalf("replaced stock input overwrote the validated cache: %v", err)
	}
}
