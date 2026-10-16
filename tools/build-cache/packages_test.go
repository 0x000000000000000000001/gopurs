package main

import (
	"bytes"
	"go/importer"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"
)

func TestPackageLayoutSplitMergeRemovalAndMigration(t *testing.T) {
	f := newFixture(t)
	const a = "p_11111111111111111111111111111111"
	const b = "p_22222222222222222222222222222222"
	previous := map[string]bool{}
	var lock os.FileInfo
	for _, step := range []string{"monolithic", "split", "coalesced", "split", "merge", "coalesced", "removed", "monolithic"} {
		f.req.Outputs = nil
		add := func(kind, module, path string) {
			source := filepath.Join("stage", filepath.FromSlash(path))
			put(t, filepath.Join(f.root, source), "// "+step+" "+kind+"\n")
			f.req.Outputs = append(f.req.Outputs, StagedOutput{Path: path, Kind: kind, Module: module, Source: source})
		}
		add("runtime", "", "gopurs_runtime/runtime.go")
		if step == "monolithic" {
			add("module", "Main", "purescript/Main.go")
			add("ffi", "Main", "purescript/Main_ffi.go")
		} else {
			add("package-init", "", "purescript/init_all/init.go")
			add("package-init", "", "purescript/init_external/init.go")
			if step == "coalesced" {
				for _, provider := range []string{a, b} {
					add("package", "", "purescript/"+provider+"/0_gopurs_package.go")
					add("order", "", "purescript/"+provider+"/zz_gopurs_order.go")
				}
			} else if step != "removed" {
				provider := a
				if step == "merge" {
					provider = b
				}
				add("package", "Main", "purescript/"+provider+"/Main.go")
				add("ffi-package", "Main", "purescript/"+b+"/Main_ffi.go")
				add("order", "", "purescript/"+b+"/zz_gopurs_order.go")
				if provider != b {
					add("order", "", "purescript/"+a+"/zz_gopurs_order.go")
				}
			}
		}
		result := publish(t, f)
		lookup(t, f, "hit")
		current := map[string]bool{}
		for _, file := range result.Manifest.Outputs {
			current[file.Path] = true
		}
		for file := range previous {
			if !current[file] {
				if _, err := os.Stat(filepath.Join(f.s.root, file)); !os.IsNotExist(err) {
					t.Fatalf("%s retained %s: %v", step, file, err)
				}
			}
		}
		previous = current
		currentLock, err := os.Stat(filepath.Join(f.s.cache, "write.lock"))
		if err != nil {
			t.Fatal(err)
		}
		if lock == nil {
			lock = currentLock
			put(t, filepath.Join(f.s.root, "go.mod"), "module gopurs/output\n\ngo 1.22\n// external\n")
			put(t, filepath.Join(f.s.root, "go.sum"), "external module sums\n")
			put(t, filepath.Join(f.s.root, "purescript", a, "external.txt"), "external content\n")
		} else if !os.SameFile(lock, currentLock) {
			t.Fatal("layout migration replaced the lock")
		}
		for file, expected := range map[string]string{"go.mod": "module gopurs/output\n\ngo 1.22\n// external\n", "go.sum": "external module sums\n", "purescript/" + a + "/external.txt": "external content\n"} {
			data, err := os.ReadFile(filepath.Join(f.s.root, file))
			if err != nil || string(data) != expected {
				t.Fatalf("external %s: %s %v", file, data, err)
			}
		}
		replayed := publish(t, f)
		if !reflect.DeepEqual(result.Manifest, replayed.Manifest) {
			t.Fatal("identical layout publication changed the manifest")
		}
	}
}

func TestPackageExportsFreezeNativeAndReplacementSources(t *testing.T) {
	t.Setenv("GOCACHE", t.TempDir())
	t.Setenv("GOFLAGS", "")
	spec, response, state := driverFixture(t)
	put(t, filepath.Join(spec.Output, "go.mod"), "module gopurs/output\n\ngo 1.22\nrequire example/offset v0.0.0\nreplace example/offset => ../offset\n")
	put(t, filepath.Join(spec.Workspace, "offset/go.mod"), "module example/offset\n\ngo 1.22\n")
	put(t, filepath.Join(spec.Workspace, "offset/offset.go"), "package offset\nfunc Value() int64 {return 0}\n")
	put(t, filepath.Join(spec.Workspace, "offset/unused/unused.go"), "package unused\nimport _ \"gopurs/output/not-a-dependency\"\n")
	put(t, filepath.Join(spec.Output, "native/native.go"), "package native\nimport \"gopurs/output/transitive\"\nfunc Offset() int64 {return transitive.Value()}\n")
	put(t, filepath.Join(spec.Output, "transitive/value.go"), "package transitive\nimport \"example/offset\"\nfunc Value() int64 {return offset.Value()}\n")
	put(t, filepath.Join(spec.Output, "native/native_test.go"), "package native\nimport _ \"gopurs/output/not-a-test-dependency\"\n")
	state.Spec.Options = append(state.Spec.Options, Setting{"go-layout", packageLayout})
	if err := packageToolchain(&state.Spec); err != nil {
		t.Fatal(err)
	}
	var err error
	state.Snapshot, err = snapshot(state.Spec)
	if err != nil {
		t.Fatal(err)
	}
	if err := planSession(response.Stage, &state, []SourceModule{{Name: "Main", Path: "src/Main.purs"}}, &response); err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(response.Output, "gopurs_runtime/runtime.go"), "package gopurs_runtime\n")
	put(t, filepath.Join(response.Output, "purescript/Main.go"), "package purescript\nimport \"gopurs/output/native\"\nfunc Value() int64 {return native.Offset()}\n")
	if err := preparePackageSession(response.Stage, &state, nil, &response); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(response.PackageRequest)
	if err != nil {
		t.Fatal(err)
	}
	var request packageFileRequest
	if err := strictJSON(data, &request); err != nil {
		t.Fatal(err)
	}
	exports := map[string][]byte{}
	for _, export := range request.Exports {
		data, err := os.ReadFile(export.Source)
		if err != nil || digest(data) != export.SHA256 {
			t.Fatalf("unfrozen export: %s %v", export.Path, err)
		}
		exports[export.Path] = data
	}
	read := importer.ForCompiler(token.NewFileSet(), "gc", func(path string) (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(exports[path])), nil
	})
	pkg, err := read.Import("gopurs/output/native")
	if err != nil || pkg.Scope().Lookup("Offset").Type().String() != "func() int64" {
		t.Fatalf("native exports changed: %v", err)
	}
	if strings.Contains(string(data), "func Value()") || len(request.Files) != 1 {
		t.Fatal("Go bodies leaked into the compact host request")
	}
	// A subsequent executable build uses this same GOCACHE. Export-only
	// archives must not masquerade as ordinary compiled packages during linking.
	workspace := filepath.Join(response.Stage, "go-exports")
	put(t, filepath.Join(workspace, "main.go"), "package main\nimport \"gopurs/output/native\"\nfunc main() {if native.Offset() != 0 {panic(\"native result changed\")}}\n")
	if _, err := packageGoCommand(workspace, "run", "-p=1", "./main.go"); err != nil {
		t.Fatalf("normal build after export acquisition: %v", err)
	}
	previous := state.Snapshot.Key
	put(t, filepath.Join(spec.Workspace, "offset/offset.go"), "package offset\nfunc Value() int {return 0}\n")
	next, err := snapshot(state.Spec)
	if err != nil || next.Key == previous {
		t.Fatalf("local replacement edit not observed: %v", err)
	}
	frozen := filepath.Join(response.Stage, "go-replacements", digest([]byte("example/offset")), "offset.go")
	content, err := os.ReadFile(frozen)
	if err != nil || !strings.Contains(string(content), "Value() int64") {
		t.Fatalf("external edit changed consumed bytes: %v", err)
	}
}

func TestPackageExportsPreserveExplicitCompilerFlags(t *testing.T) {
	t.Setenv("GOFLAGS", "-gcflags=all=-gopurs-invalid-compiler-flag")
	spec, response, state := driverFixture(t)
	put(t, filepath.Join(spec.Output, "go.mod"), "module gopurs/output\n\ngo 1.22\n")
	state.Spec.Options = append(state.Spec.Options, Setting{"go-layout", packageLayout})
	if err := packageToolchain(&state.Spec); err != nil {
		t.Fatal(err)
	}
	var err error
	state.Snapshot, err = snapshot(state.Spec)
	if err != nil {
		t.Fatal(err)
	}
	if err := planSession(response.Stage, &state, []SourceModule{{Name: "Main", Path: "src/Main.purs"}}, &response); err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(response.Output, "gopurs_runtime/runtime.go"), "package gopurs_runtime\nfunc Value() int64 {return 0}\n")
	put(t, filepath.Join(response.Output, "purescript/Main.go"), "package purescript\nimport \"gopurs/output/gopurs_runtime\"\nfunc Value() int64 {return gopurs_runtime.Value()}\n")
	err = preparePackageSession(response.Stage, &state, nil, &response)
	if err == nil || !strings.Contains(err.Error(), "gopurs-invalid-compiler-flag") {
		t.Fatalf("explicit compiler flags were discarded: %v", err)
	}
}

func TestPackagePublicationRetiresUnreachableObjects(t *testing.T) {
	spec, result, state, sources := packageModuleFixture(t)
	s := store{root: spec.Output, cache: filepath.Join(spec.Output, ".gopurs-cache")}
	before, why := s.readManifest()
	if before == nil {
		t.Fatal(why)
	}
	lock, err := os.Stat(filepath.Join(s.cache, "write.lock"))
	if err != nil {
		t.Fatal(err)
	}
	orphan, err := s.putObject([]byte("obsolete generated object"))
	if err != nil {
		t.Fatal(err)
	}
	// Another build has already frozen its inputs before this publication.
	frozen := t.TempDir()
	if err := os.Mkdir(filepath.Join(frozen, "prior"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := freezeModuleCache(frozen, state, false, &result); err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(s.cache, "objects", "external-note.txt"), "retained note")
	put(t, filepath.Join(result.Output, "purescript/Main.go"), sources["module"]+"// next generation\n")
	if err := publishSession(result.Stage, state, SessionRequest{ModuleCache: true}, &result); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(s.cache, "objects", orphan.SHA256)); !os.IsNotExist(err) {
		t.Fatalf("obsolete object survived successful package publication: %v", err)
	}
	for _, source := range before.Modules[0].Sources {
		if source.Kind == "module" {
			if _, err := os.Stat(filepath.Join(s.cache, "objects", source.Object.SHA256)); !os.IsNotExist(err) {
				t.Fatalf("old source object retained: %v", err)
			}
		}
	}
	for kind, want := range sources {
		data, err := os.ReadFile(filepath.Join(frozen, "prior/Main", kind+".go"))
		if err != nil || string(data) != want {
			t.Fatalf("active build's private %s changed: %s %v", kind, data, err)
		}
	}
	request := Request{Schema: Schema, Operation: "lookup", Spec: state.Spec}
	call(t, request, "hit")
	after, why := s.readManifest()
	if after == nil {
		t.Fatal(why)
	}
	for _, source := range after.Modules[0].Sources {
		if _, err := s.readModuleSource(source); err != nil {
			t.Fatalf("live source object removed: %v", err)
		}
	}
	currentLock, err := os.Stat(filepath.Join(s.cache, "write.lock"))
	if err != nil || !os.SameFile(lock, currentLock) {
		t.Fatalf("lock inode changed: %v", err)
	}
	if data, err := os.ReadFile(filepath.Join(s.cache, "objects", "external-note.txt")); err != nil || string(data) != "retained note" {
		t.Fatalf("unrelated file removed: %s %v", data, err)
	}
	// The previously frozen build may still publish: its private source bytes
	// recreate an object that maintenance just removed from the shared cache.
	frozenSource, err := os.ReadFile(filepath.Join(frozen, "prior/Main/module.go"))
	if err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(result.Output, "purescript/Main.go"), string(frozenSource))
	if err := publishSession(result.Stage, state, SessionRequest{ModuleCache: true, Diagnostics: true}, &result); err != nil {
		t.Fatal(err)
	}
	after, why = s.readManifest()
	if after == nil || after.Diagnostics == nil {
		t.Fatalf("diagnostics lost: %s", why)
	}
	for _, artifact := range artifacts(*after) {
		if _, err := s.readObject(artifact.Object); err != nil {
			t.Fatalf("live business/diagnostic object removed: %v", err)
		}
	}
	for _, source := range after.Modules[0].Sources {
		data, err := s.readModuleSource(source)
		if err != nil || string(data) != sources[source.Kind] {
			t.Fatalf("frozen source did not republish: %s %v", data, err)
		}
	}
}

func TestPackageMaintenancePreservesIncompletePublications(t *testing.T) {
	for _, mode := range []string{"pending-outputs", "absent-journal", "corrupt-journal", "corrupt-pointer", "corrupt-manifest"} {
		t.Run(mode, func(t *testing.T) {
			spec, _, _, _ := packageModuleFixture(t)
			s := store{root: spec.Output, cache: filepath.Join(spec.Output, ".gopurs-cache")}
			orphan, err := s.putObject([]byte("pending transaction object"))
			if err != nil {
				t.Fatal(err)
			}
			manifest, why := s.readManifest()
			if manifest == nil {
				t.Fatal(why)
			}
			switch mode {
			case "pending-outputs":
				pending := append(append([]OwnedOutput{}, manifest.Outputs...), OwnedOutput{"purescript/p_root/Pending.go", "package", "", orphan})
				if err := s.writeOwnership(pending); err != nil {
					t.Fatal(err)
				}
			case "absent-journal":
				if err := os.Remove(filepath.Join(s.root, ".gopurs-ownership.json")); err != nil {
					t.Fatal(err)
				}
			case "corrupt-journal":
				put(t, filepath.Join(s.root, ".gopurs-ownership.json"), "corrupt")
			case "corrupt-pointer":
				put(t, filepath.Join(s.cache, "current.json"), "corrupt")
			case "corrupt-manifest":
				data, err := os.ReadFile(filepath.Join(s.cache, "current.json"))
				var current pointer
				if err != nil || strictJSON(data, &current) != nil {
					t.Fatal("invalid fixture pointer")
				}
				put(t, filepath.Join(s.cache, "objects", current.Manifest.SHA256), "corrupt")
			}
			_ = s.retirePackageObjects()
			if _, err := s.readObject(orphan); err != nil {
				t.Fatalf("incomplete publication's object removed: %v", err)
			}
		})
	}
}

func TestPackageMaintenanceConcurrentPublishersAndReaders(t *testing.T) {
	spec, result, state, sources := packageModuleFixture(t)
	s := store{root: spec.Output, cache: filepath.Join(spec.Output, ".gopurs-cache")}
	other := t.TempDir()
	if err := filepath.Walk(result.Stage, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(result.Stage, path)
		if err != nil {
			return err
		}
		target := filepath.Join(other, relative)
		if info.IsDir() {
			return os.MkdirAll(target, 0700)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0600)
	}); err != nil {
		t.Fatal(err)
	}
	alternate := sources["module"] + "// concurrent version\n"
	put(t, filepath.Join(other, "output/purescript/Main.go"), alternate)
	readers := t.TempDir()
	var workers sync.WaitGroup
	for _, stage := range []string{result.Stage, other} {
		stage := stage
		workers.Add(1)
		go func() {
			defer workers.Done()
			for i := 0; i < 4; i++ {
				var response SessionResponse
				if err := publishSession(stage, state, SessionRequest{ModuleCache: true}, &response); err != nil {
					t.Errorf("concurrent publication: %v", err)
					return
				}
			}
		}()
	}
	workers.Add(2)
	go func() {
		defer workers.Done()
		for i := 0; i < 8; i++ {
			if err := s.retirePackageObjects(); err != nil {
				t.Errorf("concurrent maintenance: %v", err)
				return
			}
		}
	}()
	go func() {
		defer workers.Done()
		for i := 0; i < 8; i++ {
			directory, err := os.MkdirTemp(readers, "frozen-")
			if err == nil {
				err = os.Mkdir(filepath.Join(directory, "prior"), 0700)
			}
			if err != nil {
				t.Errorf("reader staging: %v", err)
				return
			}
			var response SessionResponse
			if err := freezeModuleCache(directory, state, false, &response); err != nil {
				t.Errorf("concurrent freeze: %v", err)
				return
			}
			data, err := os.ReadFile(filepath.Join(directory, "prior/Main/module.go"))
			if err != nil || string(data) != sources["module"] && string(data) != alternate {
				t.Errorf("concurrent reader lost its source: %s %v", data, err)
				return
			}
		}
	}()
	workers.Wait()
	call(t, Request{Schema: Schema, Operation: "lookup", Spec: state.Spec}, "hit")
}
