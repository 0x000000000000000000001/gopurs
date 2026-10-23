package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const overlayOriginalPackage = "p_root"
const overlayDependentPackage = "p_22222222222222222222222222222222"

func sectionFor(path, source string) string {
	return sourceMarker + hex.EncodeToString([]byte(path)) + "\n" + source + sourceEndMarker + "\n"
}

func testBundle(packageName string, imports string, sections ...string) string {
	var out strings.Builder
	out.WriteString("package " + packageName + "\n")
	if imports != "" {
		out.WriteString("\nimport (\n" + imports + ")\n")
	}
	out.WriteString("\n")
	for _, section := range sections {
		out.WriteString(section)
	}
	return out.String()
}

type overlayFixture struct {
	output  string
	stage   string
	state   buildSession
	request packageFileRequest
	record  *packageOverlayRecord
}

func copyDirectory(t *testing.T, source, destination string) {
	t.Helper()
	entries, err := os.ReadDir(source)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(destination, 0700); err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		from := filepath.Join(source, entry.Name())
		to := filepath.Join(destination, entry.Name())
		if entry.IsDir() {
			copyDirectory(t, from, to)
			continue
		}
		data, err := os.ReadFile(from)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(to, data, 0600); err != nil {
			t.Fatal(err)
		}
	}
}

func newOverlayFixture(t *testing.T, newBody string) *overlayFixture {
	t.Helper()
	output := t.TempDir()
	stage := t.TempDir()
	if err := os.MkdirAll(filepath.Join(output, ".gopurs-cache/go-packages/plan-v1"), 0700); err != nil {
		t.Fatal(err)
	}
	oldMain := "func Get_Main() string { return Helper() + \"-old\" }\n"
	helper := "func Helper() string { return \"helper\" }\n"
	newMain := "package purescript\nfunc Get_Main() string { return Helper() + \"" + newBody + "\" }\n"
	put(t, filepath.Join(stage, "output/purescript/Main.go"), newMain)
	put(t, filepath.Join(stage, "output/purescript/Helper.go"), "package purescript\n"+helper)

	mainBundle := testBundle(overlayOriginalPackage, "",
		sectionFor("Main.go", oldMain), sectionFor("Helper.go", helper))
	dependent := testBundle(overlayDependentPackage, "\told \"gopurs/output/purescript/p_root\"\n",
		sectionFor("Dependent.go", "func Run() string { return old.Get_Main() }\n"))
	files := []packagePlanFile{
		{Path: overlayOriginalPackage + "/0_gopurs_package.go", Kind: "package", SHA256: digest([]byte(mainBundle)), Bytes: int64(len(mainBundle))},
		{Path: overlayDependentPackage + "/0_gopurs_package.go", Kind: "package", SHA256: digest([]byte(dependent)), Bytes: int64(len(dependent))},
		{Path: overlayOriginalPackage + "/zz_gopurs_order.go", Kind: "order", SHA256: digest([]byte("package p_root\n")), Bytes: int64(len("package p_root\n"))},
		{Path: overlayDependentPackage + "/zz_gopurs_order.go", Kind: "order", SHA256: digest([]byte("package " + overlayDependentPackage + "\n")), Bytes: int64(len("package " + overlayDependentPackage + "\n"))},
		{Path: "init_all/init.go", Kind: "package-init", SHA256: digest([]byte("package init_all\n")), Bytes: int64(len("package init_all\n"))},
		{Path: "init_external/init.go", Kind: "package-init", SHA256: digest([]byte("package init_external\n")), Bytes: int64(len("package init_external\n"))},
	}
	contents := map[string]string{
		overlayOriginalPackage + "/0_gopurs_package.go":  mainBundle,
		overlayDependentPackage + "/0_gopurs_package.go": dependent,
		overlayOriginalPackage + "/zz_gopurs_order.go":   "package p_root\n",
		overlayDependentPackage + "/zz_gopurs_order.go":  "package " + overlayDependentPackage + "\n",
		"init_all/init.go":      "package init_all\n",
		"init_external/init.go": "package init_external\n",
	}
	for _, file := range files {
		put(t, filepath.Join(output, "purescript", filepath.FromSlash(file.Path)), contents[file.Path])
	}
	previous := packagePlanResult{Format: packageFilesFormat, PlanHash: digest([]byte("base-plan")),
		Packages: []packagePlanPackage{
			{Name: overlayOriginalPackage, Declarations: 2},
			{Name: overlayDependentPackage, Dependencies: []string{overlayOriginalPackage}, Declarations: 1},
		},
		Files: files, Entries: []packagePlanEntry{{Symbol: packageSymbol{"generated/Main", "Get_Main_main"}, Package: "gopurs/output/purescript/" + overlayOriginalPackage}}}
	planData, _ := json.Marshal(previous)
	if err := os.WriteFile(filepath.Join(output, ".gopurs-cache/go-packages/plan-v1/result.json"), planData, 0600); err != nil {
		t.Fatal(err)
	}
	previousRequest := packageFileRequest{Format: packageFilesFormat, PackagePath: "gopurs/output/purescript",
		Files: []packageSource{
			{Owner: "generated/Main", Module: "Main", Path: "Main.go", SHA256: digest([]byte(oldMain))},
			{Owner: "generated/Helper", Module: "Helper", Path: "Helper.go", SHA256: digest([]byte(helper))},
		},
		Entries: []packageSymbol{{Owner: "generated/Main", Name: "Get_Main_main"}}}
	requestData, _ := json.Marshal(previousRequest)
	if err := os.WriteFile(filepath.Join(output, ".gopurs-cache/go-packages/plan-v1/request.json"), requestData, 0600); err != nil {
		t.Fatal(err)
	}
	request := packageFileRequest{Format: packageFilesFormat, PackagePath: "gopurs/output/purescript",
		Files: []packageSource{
			{Owner: "generated/Main", Module: "Main", Path: "Main.go", Source: filepath.Join(stage, "output/purescript/Main.go"), SHA256: digest([]byte(newMain))},
			{Owner: "generated/Helper", Module: "Helper", Path: "Helper.go", SHA256: digest([]byte(helper))},
		},
		Entries: previousRequest.Entries}
	state := buildSession{Spec: Spec{Output: output}, Modules: []SourceModule{{Name: "Main"}, {Name: "Helper"}}}
	return &overlayFixture{output: output, stage: stage, state: state, request: request}
}

func TestPackageOverlayMaterializesOnlyMissingSignatures(t *testing.T) {
	for _, missingSignature := range []bool{false, true} {
		t.Run(fmt.Sprintf("missing-signature-%t", missingSignature), func(t *testing.T) {
			t.Setenv(packPublication, "0")
			t.Setenv("GOPURS_GO_PACKAGE_OVERLAY", "1")
			f := newOverlayFixture(t, "-new")
			f.state.Planned = true
			var err error
			f.state.Spec.Output, err = filepath.EvalSymlinks(f.output)
			if err != nil {
				t.Fatal(err)
			}
			f.state.Spec.Workspace = f.state.Spec.Output
			f.state.Spec.Options = []Setting{{"go-layout", packageLayout}, {"module-lazy", "1"}}
			f.state.Spec.Compiler = []Input{{Name: "compiler", Kind: "file", Path: filepath.Join(f.output, "compiler")}}
			f.state.Spec.Inputs = []Input{{Name: "tast", Kind: "tast", Path: f.state.Spec.Output}}
			put(t, f.state.Spec.Compiler[0].Path, "test compiler")
			put(t, filepath.Join(f.output, "go.mod"), "module gopurs/output\n\ngo 1.22\n")
			put(t, filepath.Join(f.stage, "output/gopurs_runtime/runtime.go"), "package gopurs_runtime\n")
			f.state.Snapshot, err = snapshot(f.state.Spec)
			if err != nil {
				t.Fatal(err)
			}
			s := store{root: f.output, cache: filepath.Join(f.output, ".gopurs-cache")}
			if err := os.MkdirAll(filepath.Join(s.cache, "objects"), 0700); err != nil {
				t.Fatal(err)
			}
			helperPath := filepath.Join(f.stage, "output/purescript/Helper.go")
			helper, err := os.ReadFile(helperPath)
			if err != nil {
				t.Fatal(err)
			}
			source, err := s.stageModuleSource("", StagedModuleSource{"module", helperPath}, nil, newPackWriter(&s))
			if err != nil {
				t.Fatal(err)
			}
			manifest := &Manifest{Modules: []ModuleMemo{{Name: "Helper", Sources: []ModuleSource{source}}}}
			lazyPriors.Lock()
			lazyPriors.values[f.stage] = newLazyPrior(manifest, s)
			lazyPriors.Unlock()
			t.Cleanup(func() { forgetLazyPrior(f.stage) })
			if err := os.Remove(helperPath); err != nil {
				t.Fatal(err)
			}
			previous := f.request
			previous.GoVersion = "go1.22"
			previous.Files = append([]packageSource{}, previous.Files...)
			previous.Files[0].SHA256 = digest([]byte("package purescript\nfunc Get_Main() string { return Helper() + \"-old\" }\n"))
			previous.Files[1].SHA256 = source.Raw.SHA256
			previous.Signatures = map[string]string{}
			for _, file := range previous.Files {
				data := helper
				if file.Module == "Main" {
					data, err = os.ReadFile(file.Source)
					if err != nil {
						t.Fatal(err)
					}
				}
				previous.Signatures[file.Path], err = sourceSignature(file.Path, data)
				if err != nil {
					t.Fatal(err)
				}
			}
			if missingSignature {
				delete(previous.Signatures, "Helper.go")
			}
			encoded, _ := json.Marshal(previous)
			put(t, filepath.Join(f.output, ".gopurs-cache/go-packages/plan-v1/request.json"), string(encoded))
			var response SessionResponse
			err = preparePackageSession(f.stage, &f.state, []string{"Main"}, &response)
			if missingSignature {
				if err != nil || response.Status != "packages-ready" {
					t.Fatalf("missing signature did not fall back to the planner: %s %v", response.Status, err)
				}
			} else if err != nil || response.Status != "packages-overlay" {
				t.Fatalf("lazy overlay: %s %v", response.Status, err)
			}
			_, statErr := os.Stat(helperPath)
			if missingSignature && statErr != nil || !missingSignature && !os.IsNotExist(statErr) {
				t.Fatalf("unexpected helper materialization: %v", statErr)
			}
		})
	}
}

func TestPackageOverlayExecutesBodyEdit(t *testing.T) {
	f := newOverlayFixture(t, "-new")
	t.Setenv("GOPURS_GO_PACKAGE_OVERLAY", "1")
	record, reason, err := preparePackageOverlay(f.stage, &f.state, []string{"Main"}, &f.request)
	if err != nil {
		t.Fatal(err)
	}
	if record == nil {
		t.Fatalf("overlay was rejected: %s", reason)
	}
	if record.OriginalPackage != overlayOriginalPackage || record.Package == overlayOriginalPackage {
		t.Fatalf("bad overlay record: %+v", record)
	}
	result, err := readPackagePlan(filepath.Join(f.stage, "package-result.json"))
	if err != nil {
		t.Fatal(err)
	}
	entries := map[string]string{}
	for _, entry := range result.Entries {
		entries[entry.Symbol.Name] = entry.Package
	}
	if entries["Get_Main_main"] != "gopurs/output/purescript/"+record.Package {
		t.Fatalf("entry not redirected: %+v", result.Entries)
	}
	packageDir := filepath.Join(f.stage, "packages")
	workspace := t.TempDir()
	put(t, filepath.Join(workspace, "go.mod"), "module gopurs/output\n\ngo 1.22\n")
	entriesDir, err := os.ReadDir(packageDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entriesDir {
		if !entry.IsDir() {
			continue
		}
		copyDirectory(t, filepath.Join(packageDir, entry.Name()), filepath.Join(workspace, "purescript", entry.Name()))
	}
	put(t, filepath.Join(workspace, "cmd/main.go"), `package main
import (
	"fmt"
	d "gopurs/output/purescript/`+overlayDependentPackage+`"
)
func main() { fmt.Print(d.Run()) }
`)
	command := exec.Command("go", "run", "./cmd")
	command.Dir = workspace
	command.Env = append(os.Environ(), "GOWORK=off")
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("go run overlay: %v: %s", err, output)
	}
	if string(output) != "helper-new" {
		t.Fatalf("overlay output %q", output)
	}
	// The oracle is the same edit rendered in the original package.
	oracle := t.TempDir()
	source := "package p_root\nfunc Get_Main() string { return Helper() + \"-new\" }\nfunc Helper() string { return \"helper\" }\n"
	put(t, filepath.Join(oracle, "go.mod"), "module gopurs/output\n\ngo 1.22\n")
	put(t, filepath.Join(oracle, "purescript/p_root/0_gopurs_package.go"), source)
	put(t, filepath.Join(oracle, "cmd/main.go"), `package main
import (
	"fmt"
	"gopurs/output/purescript/p_root"
)
func main() { fmt.Print(p_root.Get_Main()) }
`)
	command = exec.Command("go", "run", "./cmd")
	command.Dir = oracle
	command.Env = append(os.Environ(), "GOWORK=off")
	full, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("go run full: %v: %s", err, full)
	}
	if string(output) != string(full) {
		t.Fatalf("overlay output %q differs from full %q", output, full)
	}
}

func TestPackageOverlayRepeatedBodyEdit(t *testing.T) {
	f := newOverlayFixture(t, "-new")
	record, reason, err := preparePackageOverlay(f.stage, &f.state, []string{"Main"}, &f.request)
	if err != nil || record == nil {
		t.Fatalf("first overlay: %+v %s %v", record, reason, err)
	}
	// Publish the overlay into the fixture output, mirroring the driver's
	// commit order: package files, result, then the overlay record.
	entries, err := os.ReadDir(filepath.Join(f.stage, "packages"))
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		copyDirectory(t, filepath.Join(f.stage, "packages", entry.Name()), filepath.Join(f.output, "purescript", entry.Name()))
	}
	for _, name := range []string{"package-result.json", "package-overlay.json"} {
		data, err := os.ReadFile(filepath.Join(f.stage, name))
		if err != nil {
			t.Fatal(err)
		}
		target := filepath.Join(f.output, ".gopurs-cache/go-packages/plan-v1/result.json")
		if name == "package-overlay.json" {
			target = filepath.Join(f.output, ".gopurs-cache/go-packages/plan-v1/overlay.json")
		}
		if err := os.WriteFile(target, data, 0600); err != nil {
			t.Fatal(err)
		}
	}
	previousRequest, _ := json.Marshal(f.request)
	if err := os.WriteFile(filepath.Join(f.output, ".gopurs-cache/go-packages/plan-v1/request.json"), previousRequest, 0600); err != nil {
		t.Fatal(err)
	}
	stage := t.TempDir()
	newer := "package purescript\nfunc Get_Main() string { return Helper() + \"-newer\" }\n"
	put(t, filepath.Join(stage, "output/purescript/Main.go"), newer)
	request := f.request
	request.Files = append([]packageSource{}, f.request.Files...)
	request.Files[0].Source = filepath.Join(stage, "output/purescript/Main.go")
	request.Files[0].SHA256 = digest([]byte(newer))
	second, reason, err := preparePackageOverlay(stage, &f.state, []string{"Main"}, &request)
	if err != nil || second == nil {
		t.Fatalf("second overlay: %+v %s %v", second, reason, err)
	}
	if second.Package != record.Package {
		t.Fatalf("overlay package changed: %s -> %s", record.Package, second.Package)
	}
	bundle, err := os.ReadFile(filepath.Join(stage, "packages", filepath.FromSlash(second.Package+"/0_gopurs_package.go")))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(bundle), "-newer") {
		t.Fatalf("second overlay kept old body: %s", bundle)
	}
}

func TestPackageOverlayRejectsNewDeclarationReference(t *testing.T) {
	f := newOverlayFixture(t, "-new")
	// The reference set is a proof of the planner's declaration edges. A new
	// call changes it even when the module signature still matches.
	put(t, filepath.Join(f.stage, "output/purescript/Main.go"),
		"package purescript\nfunc Get_Main() string { return Helper() + Run() }\n")
	data, err := os.ReadFile(filepath.Join(f.stage, "output/purescript/Main.go"))
	if err != nil {
		t.Fatal(err)
	}
	f.request.Files[0].SHA256 = digest(data)
	record, reason, err := preparePackageOverlay(f.stage, &f.state, []string{"Main"}, &f.request)
	if err != nil {
		t.Fatal(err)
	}
	if record != nil || reason == "" {
		t.Fatalf("unproved reference change was accepted: %+v", record)
	}
}

func TestPackageOverlayRejectsTypeDeclaration(t *testing.T) {
	f := newOverlayFixture(t, "-new")
	put(t, filepath.Join(f.stage, "output/purescript/Main.go"),
		"package purescript\ntype Extra struct{ A int }\nfunc Get_Main() string { return Helper() + \"-new\" }\n")
	data, err := os.ReadFile(filepath.Join(f.stage, "output/purescript/Main.go"))
	if err != nil {
		t.Fatal(err)
	}
	f.request.Files[0].SHA256 = digest(data)
	record, reason, err := preparePackageOverlay(f.stage, &f.state, []string{"Main"}, &f.request)
	if err != nil {
		t.Fatal(err)
	}
	if record != nil || reason == "" {
		t.Fatalf("type declaration was accepted: %+v", record)
	}
}
