package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopurs/c1c"
)

func TestSessionFFIValidationUsesSelectedEntries(t *testing.T) {
	for _, layout := range []string{"gopurs/monolithic-v1", packageLayout} {
		t.Run(layout, func(t *testing.T) {
			_, result, state := driverFixture(t)
			state.Spec.Options = append(state.Spec.Options, Setting{Name: "go-layout", Value: layout})
			var err error
			state.Snapshot, err = snapshot(state.Spec)
			if err != nil {
				t.Fatal(err)
			}
			modules := []SourceModule{{Name: "App.Main", Path: "src/App/Main.purs", HasForeign: true}}
			if err := planSession(result.Stage, &state, modules, &result); err != nil {
				t.Fatal(err)
			}
			put(t, filepath.Join(result.Output, "purescript/App_Main.go"), `package purescript
func Get_App_Main_main() { Missing() }
func Get_App_Main_unused() {}
`)
			put(t, filepath.Join(result.Output, "purescript/App_Main_ffi.go"), `package purescript
import "example.invalid/native-dependency"
`+c1c.MissingFFIMarker+`{"binding":"Missing","module":"App.Main","symbol":"missing","path":"src/App/Main.go"}
var Missing = func() { panic("FFI not implemented: missing") }
`)
			err = validateSessionFFI(result.Stage, state, []string{"App.Main"})
			if err == nil || !strings.Contains(err.Error(), "App.Main") || !strings.Contains(err.Error(), "missing") {
				t.Fatalf("reachable missing FFI accepted: %v", err)
			}
			if strings.Contains(err.Error(), "incremental cache:") {
				t.Fatalf("source error classified as a cache failure: %v", err)
			}
			if err := validateSessionFFI(result.Stage, state, nil); err != nil {
				t.Fatalf("missing FFI blocked a build without an explicit entry: %v", err)
			}
			put(t, filepath.Join(result.Output, "purescript/App_Main.go"), `package purescript
func Get_App_Main_main() {}
func Unreachable() { Missing() }
`)
			if err := validateSessionFFI(result.Stage, state, []string{"App.Main"}); err != nil {
				t.Fatalf("unused FFI or unavailable native import blocked the build: %v", err)
			}
		})
	}
}

func TestSessionFFIValidationKeepsUnneededSourcesLazy(t *testing.T) {
	stage, state, _ := lazyModuleFixture(t)
	if err := validateSessionFFI(stage, state, []string{"Main"}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(stage, "output/purescript/Main.go")); !os.IsNotExist(err) {
		t.Fatalf("validation without missing FFI materialized generated source: %v", err)
	}
	if _, err := os.Stat(filepath.Join(stage, "output/purescript/Main_ffi.go")); err != nil {
		t.Fatalf("cached FFI source was not inspected: %v", err)
	}
}

func TestSessionFFIValidationTraversesRestoredSources(t *testing.T) {
	stage, state, _ := lazyModuleFixture(t)
	state.Modules = append(state.Modules, SourceModule{Name: "Entry", Path: "src/Entry.purs"})
	put(t, filepath.Join(stage, "output/purescript/Entry.go"), `package purescript
func Get_Entry_main() { Use() }
`)
	// The cached Main module calls Foreign. Only the entry and missing stub
	// are fresh, so validation must read the restored declaration to connect them.
	put(t, filepath.Join(stage, "output/purescript/Main_ffi.go"), `package purescript
`+c1c.MissingFFIMarker+`{"binding":"Foreign","module":"Main","symbol":"foreign","path":"src/Main.go"}
var Foreign = func() int { panic("FFI not implemented: foreign") }
`)
	err := validateSessionFFI(stage, state, []string{"Entry"})
	if err == nil || !strings.Contains(err.Error(), "Main") || !strings.Contains(err.Error(), "foreign") {
		t.Fatalf("restored call to missing FFI escaped validation: %v", err)
	}
	if _, err := os.Stat(filepath.Join(stage, "output/purescript/Main.go")); err != nil {
		t.Fatalf("cached generated source was not inspected: %v", err)
	}
}
