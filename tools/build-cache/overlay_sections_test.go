package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const partitionPackage = "p_33333333333333333333333333333333"

func partitionSource(suffix string) string {
	return "package purescript\nimport f \"fmt\"\n" +
		"func Get_Main() string { return Helper() + Other() + \"" + suffix + "\" }\n" +
		"var privateCache string\nfunc Other() string { privateCache = f.Sprint(\"other\"); return privateCache }\n"
}

func replaceFixturePlanFile(t *testing.T, f *overlayFixture, path, source string) {
	t.Helper()
	planPath := filepath.Join(f.output, ".gopurs-cache/go-packages/plan-v1/result.json")
	plan, err := readPackagePlan(planPath)
	if err != nil {
		t.Fatal(err)
	}
	for i := range plan.Files {
		if plan.Files[i].Path == path {
			plan.Files[i].SHA256, plan.Files[i].Bytes = digest([]byte(source)), int64(len(source))
		}
	}
	put(t, filepath.Join(f.output, "purescript", path), source)
	encoded, _ := json.Marshal(plan)
	put(t, planPath, string(encoded))
}

func newPartitionFixture(t *testing.T) *overlayFixture {
	t.Helper()
	f := newOverlayFixture(t, "-new")
	source := partitionSource("-new")
	put(t, f.request.Files[0].Source, source)
	f.request.Files[0].SHA256 = digest([]byte(source))
	mainBundle := testBundle(overlayOriginalPackage, "\tgother \"gopurs/output/purescript/"+partitionPackage+"\"\n",
		sectionFor("Main.go", "func Get_Main() string { return Helper() + gother.Other() + \"-old\" }\n"),
		sectionFor("Helper.go", "func Helper() string { return \"helper\" }\n"))
	replaceFixturePlanFile(t, f, overlayOriginalPackage+"/"+packageBundleName, mainBundle)
	other := testBundle(partitionPackage, "\t\"fmt\"\n", sectionFor("Main.go",
		"var privateCache string\nfunc Other() string { privateCache = fmt.Sprint(\"other\"); return privateCache }\n"))
	planPath := filepath.Join(f.output, ".gopurs-cache/go-packages/plan-v1/result.json")
	plan, err := readPackagePlan(planPath)
	if err != nil {
		t.Fatal(err)
	}
	plan.Packages[0].Dependencies = []string{partitionPackage}
	plan.Packages = append(plan.Packages, packagePlanPackage{Name: partitionPackage, Declarations: 2})
	plan.Packages, err = topologicalPackages(plan.Packages)
	if err != nil {
		t.Fatal(err)
	}
	plan.Files = append(plan.Files, packagePlanFile{Path: partitionPackage + "/" + packageBundleName, Kind: "package", SHA256: digest([]byte(other)), Bytes: int64(len(other))})
	order := "package " + partitionPackage + "\n"
	plan.Files = append(plan.Files, packagePlanFile{Path: partitionPackage + "/zz_gopurs_order.go", Kind: "order", SHA256: digest([]byte(order)), Bytes: int64(len(order))})
	encoded, _ := json.Marshal(plan)
	put(t, planPath, string(encoded))
	put(t, filepath.Join(f.output, "purescript", partitionPackage, packageBundleName), other)
	put(t, filepath.Join(f.output, "purescript", partitionPackage, "zz_gopurs_order.go"), order)
	return f
}

func freezePartitionFixture(t *testing.T, f *overlayFixture, previousSuffix string) {
	t.Helper()
	planDir := filepath.Join(f.output, ".gopurs-cache/go-packages/plan-v1")
	request, err := readPackageRequest(filepath.Join(planDir, "request.json"))
	if err != nil {
		t.Fatal(err)
	}
	previous := partitionSource(previousSuffix)
	request.Files[0].SHA256 = digest([]byte(previous))
	encoded, _ := json.Marshal(request)
	put(t, filepath.Join(planDir, "request.json"), string(encoded))
	put(t, filepath.Join(f.stage, "prior/Main/module.go"), previous)
	plan, err := readPackagePlan(filepath.Join(planDir, "result.json"))
	if err != nil {
		t.Fatal(err)
	}
	manifest := &Manifest{}
	for _, file := range plan.Files {
		manifest.Outputs = append(manifest.Outputs, OwnedOutput{
			Path: "purescript/" + file.Path, Kind: file.Kind, Module: file.Module, Object: Object{file.SHA256, file.Bytes},
		})
	}
	f.state.Spec.Options = []Setting{{"module-lazy", "1"}}
	stage := f.stage
	lazyPriors.Lock()
	lazyPriors.values[stage] = newLazyPrior(manifest, store{})
	lazyPriors.Unlock()
	t.Cleanup(func() { forgetLazyPrior(stage) })
	if !committedPackageFiles(stage, f.state, plan) {
		t.Fatal("fixture does not match the committed package files")
	}
}

func runPartitionProgram(t *testing.T, packages, entry string) string {
	t.Helper()
	workspace := t.TempDir()
	put(t, filepath.Join(workspace, "go.mod"), "module gopurs/output\n\ngo 1.22\n")
	copyDirectory(t, packages, filepath.Join(workspace, "purescript"))
	put(t, filepath.Join(workspace, "cmd/main.go"), "package main\nimport (\"fmt\"; p \"gopurs/output/purescript/"+entry+"\")\nfunc main() { fmt.Print(p.Get_Main()) }\n")
	command := exec.Command("go", "run", "./cmd")
	command.Dir = workspace
	command.Env = append(os.Environ(), "GOWORK=off")
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("execute partition: %v\n%s", err, output)
	}
	return string(output)
}

func TestPackageSectionsExecuteAgainstFullAndRepeat(t *testing.T) {
	for _, committed := range []bool{false, true} {
		name := "partition"
		if committed {
			name = "committed-literals"
		}
		t.Run(name, func(t *testing.T) {
			f := newPartitionFixture(t)
			previousSuffix := "-old"
			for _, suffix := range []string{"-new", "-newer"} {
				f.stage = t.TempDir()
				source := partitionSource(suffix)
				f.request.Files[0].Source = filepath.Join(f.stage, "Main.go")
				f.request.Files[0].SHA256 = digest([]byte(source))
				put(t, f.request.Files[0].Source, source)
				if committed {
					freezePartitionFixture(t, f, previousSuffix)
				}
				record, reason, err := preparePackageOverlay(f.stage, &f.state, []string{"Main"}, &f.request)
				if err != nil || record == nil {
					t.Fatalf("partition refused: %s %v", reason, err)
				}
				if _, err := os.Stat(filepath.Join(f.stage, "package-overlay.json")); !os.IsNotExist(err) {
					t.Fatalf("in-place edit unexpectedly wrote relocation record: %v", err)
				}
				actual := runPartitionProgram(t, filepath.Join(f.stage, "packages"), overlayOriginalPackage)
				full := t.TempDir()
				put(t, filepath.Join(full, "purescript/full/full.go"), strings.Replace(source, "package purescript", "package full", 1)+"\nfunc Helper() string { return \"helper\" }\n")
				oracle := runPartitionProgram(t, filepath.Join(full, "purescript"), "full")
				if actual != oracle || actual != "helperother"+suffix {
					t.Fatalf("partition %q != full %q", actual, oracle)
				}
				copyDirectory(t, filepath.Join(f.stage, "packages"), filepath.Join(f.output, "purescript"))
				result, err := os.ReadFile(filepath.Join(f.stage, "package-result.json"))
				if err != nil {
					t.Fatal(err)
				}
				put(t, filepath.Join(f.output, ".gopurs-cache/go-packages/plan-v1/result.json"), string(result))
				previousSuffix = suffix
				request, _ := json.Marshal(f.request)
				put(t, filepath.Join(f.output, ".gopurs-cache/go-packages/plan-v1/request.json"), string(request))
			}
		})
	}
}

func TestPackageSectionsRejectUnprovedChanges(t *testing.T) {
	for name, edit := range map[string]func(string) string{
		"name":   func(s string) string { return strings.ReplaceAll(s, "Other", "Another") },
		"import": func(s string) string { return strings.Replace(s, "\"fmt\"", "\"strings\"", 1) },
		"type":   func(s string) string { return s + "\ntype Extra struct{}\n" },
		"init":   func(s string) string { return s + "\nfunc init() {}\n" },
		"initialized-variable": func(s string) string {
			return strings.Replace(s, "var privateCache string", "var privateCache = \"initial\"", 1)
		},
		"signature":    func(s string) string { return strings.Replace(s, "func Get_Main()", "func Get_Main(extra int)", 1) },
		"graph":        func(s string) string { return strings.Replace(s, "Helper() + Other()", "Other()", 1) },
		"private-edge": func(s string) string { return strings.Replace(s, "Helper() + Other()", "Helper() + privateCache", 1) },
	} {
		t.Run(name, func(t *testing.T) {
			f := newPartitionFixture(t)
			freezePartitionFixture(t, f, "-old")
			source := edit(partitionSource("-new"))
			put(t, f.request.Files[0].Source, source)
			f.request.Files[0].SHA256 = digest([]byte(source))
			record, reason, err := preparePackageOverlay(f.stage, &f.state, []string{"Main"}, &f.request)
			if err != nil || record != nil || reason == "" {
				t.Fatalf("accepted unproved edit: %+v %s %v", record, reason, err)
			}
		})
	}
}

func TestPackageSectionsRejectAmbiguousPlacement(t *testing.T) {
	f := newPartitionFixture(t)
	freezePartitionFixture(t, f, "-old")
	other := testBundle(partitionPackage, "", sectionFor("Main.go", "func Get_Main() string { return \"duplicate\" }\n"))
	replaceFixturePlanFile(t, f, partitionPackage+"/"+packageBundleName, other)
	record, reason, err := preparePackageOverlay(f.stage, &f.state, []string{"Main"}, &f.request)
	if err != nil || record != nil || reason == "" {
		t.Fatalf("accepted ambiguous placement: %+v %s %v", record, reason, err)
	}
}

func TestPackageLiteralSectionsMatchCompleteIndex(t *testing.T) {
	f := newPartitionFixture(t)
	freezePartitionFixture(t, f, "-old")
	var expected string
	for _, complete := range []bool{false, true} {
		if complete {
			f.stage = t.TempDir()
			f.state.Spec.Options = nil // An eager client retains the complete index.
			put(t, filepath.Join(f.stage, "prior/Main/module.go"), partitionSource("-old"))
		}
		record, reason, err := preparePackageOverlay(f.stage, &f.state, []string{"Main"}, &f.request)
		if err != nil || record == nil {
			t.Fatalf("complete=%t: %s %v", complete, reason, err)
		}
		result, err := os.ReadFile(filepath.Join(f.stage, "package-result.json"))
		if err != nil {
			t.Fatal(err)
		}
		if complete && string(result) != expected {
			t.Fatal("literal path changed package bytes or plan versus complete index")
		}
		expected = string(result)
	}
}

func TestPackageLiteralSectionsRejectAlteredPlanHints(t *testing.T) {
	for _, mode := range []string{"name", "import", "outside-section"} {
		t.Run(mode, func(t *testing.T) {
			f := newPartitionFixture(t)
			freezePartitionFixture(t, f, "-old")
			path := overlayDependentPackage + "/" + packageBundleName
			before, err := os.ReadFile(filepath.Join(f.output, "purescript", path))
			if err != nil {
				t.Fatal(err)
			}
			source := string(before)
			switch mode {
			case "name":
				// A duplicate exported name outside the changed module must
				// still reach the ordinary global-index rejection.
				source = strings.Replace(source, "func Run", "func Helper", 1)
			case "import":
				// Imports of a package are checked even when it contributes
				// no section to the changed producer.
				source = strings.Replace(source, "old \"", "old old \"", 1)
			case "outside-section":
				source += sectionFor("../outside.go", "func Outside() {}\n")
			}
			// Recomputing the public hint hash must not grant the partial
			// index permission to trust a different committed generation.
			replaceFixturePlanFile(t, f, path, source)
			plan, err := readPackagePlan(filepath.Join(f.output, ".gopurs-cache/go-packages/plan-v1/result.json"))
			if err != nil {
				t.Fatal(err)
			}
			if committedPackageFiles(f.stage, f.state, plan) {
				t.Fatal("altered hint still matches frozen manifest")
			}
			record, reason, err := preparePackageOverlay(f.stage, &f.state, []string{"Main"}, &f.request)
			if err != nil || record != nil || reason != "published bundles are not reusable" {
				t.Fatalf("lost complete-index rejection: %+v %q %v", record, reason, err)
			}
		})
	}
}

func TestPackageSnapshotRetainsProofAndRechecksStagedLinks(t *testing.T) {
	for _, revert := range []bool{false, true} {
		t.Run(map[bool]string{false: "changed-after-snapshot", true: "edited-and-reverted"}[revert], func(t *testing.T) {
			f := newPartitionFixture(t)
			freezePartitionFixture(t, f, "-old")
			planDir := filepath.Join(f.output, ".gopurs-cache/go-packages/plan-v1")
			previous, err := readPackageRequest(filepath.Join(planDir, "request.json"))
			if err != nil {
				t.Fatal(err)
			}
			result, err := readPackagePlan(filepath.Join(planDir, "result.json"))
			if err != nil {
				t.Fatal(err)
			}
			sources, err := readPackagePlanFiles(f.output, result)
			if err != nil {
				t.Fatal(err)
			}
			snapshot := &packagePlanSnapshot{request: previous, result: result, sources: sources}
			path := filepath.Join(f.output, "purescript", overlayDependentPackage, packageBundleName)
			before, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			info, err := os.Stat(path)
			if err != nil {
				t.Fatal(err)
			}
			changed := strings.Replace(string(before), "Get_Main", "Get_Fake", 1)
			if changed == string(before) {
				t.Fatal("missing dependent call")
			}
			put(t, path, changed)
			if err := os.Chtimes(path, info.ModTime(), info.ModTime()); err != nil {
				t.Fatal(err)
			}
			if revert {
				put(t, path, string(before))
			}
			record, reason, err := preparePackageOverlaySnapshot(f.stage, &f.state, []string{"Main"}, &f.request, snapshot)
			if err != nil {
				t.Fatal(err)
			}
			if revert {
				if record == nil || reason != "" {
					t.Fatalf("reverted bytes refused: %s", reason)
				}
				if _, err := packageOutputs(f.stage, f.state.Modules, []string{"Main"}); err != nil {
					t.Fatal(err)
				}
			} else {
				if record != nil || !strings.Contains(reason, "package contract") {
					t.Fatalf("accepted changed linked bytes: %+v %s", record, reason)
				}
				if _, err := os.Stat(filepath.Join(f.stage, "package-result.json")); !os.IsNotExist(err) {
					t.Fatalf("failed snapshot left a result: %v", err)
				}
			}
		})
	}
}

func TestPackageSnapshotRejectsCorruptionAndSymlinks(t *testing.T) {
	for _, mode := range []string{"bundle", "order", "symlink", "missing", "unreadable"} {
		t.Run(mode, func(t *testing.T) {
			if mode == "unreadable" && os.Geteuid() == 0 {
				t.Skip("root can read files without permission")
			}
			f := newPartitionFixture(t)
			freezePartitionFixture(t, f, "-old")
			plan, err := readPackagePlan(filepath.Join(f.output, ".gopurs-cache/go-packages/plan-v1/result.json"))
			if err != nil {
				t.Fatal(err)
			}
			name := packageBundleName
			if mode == "order" {
				name = "zz_gopurs_order.go"
			}
			path := filepath.Join(f.output, "purescript", overlayOriginalPackage, name)
			if mode == "symlink" {
				data, err := os.ReadFile(path)
				if err != nil {
					t.Fatal(err)
				}
				target := filepath.Join(t.TempDir(), "same.go")
				put(t, target, string(data))
				if err := os.Remove(path); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(target, path); err != nil {
					t.Fatal(err)
				}
			} else if mode == "missing" {
				if err := os.Remove(path); err != nil {
					t.Fatal(err)
				}
			} else if mode == "unreadable" {
				if err := os.Chmod(path, 0000); err != nil {
					t.Fatal(err)
				}
				t.Cleanup(func() { _ = os.Chmod(path, 0600) })
			} else {
				put(t, path, "corrupt")
			}
			if _, err := readPackagePlanFiles(f.output, plan); err == nil {
				t.Fatal("accepted invalid published bytes")
			}
		})
	}
}

func TestPackageSnapshotPreservesReadFailureOrderAndImportRefusal(t *testing.T) {
	output := t.TempDir()
	plan := &packagePlanResult{}
	for _, name := range []string{"Invalid", "MissingFirst", "MissingSecond"} {
		source := "package " + overlayOriginalPackage + "\n"
		if name == "Invalid" {
			source += "import (\n"
		}
		path := overlayOriginalPackage + "/" + name + ".go"
		plan.Files = append(plan.Files, packagePlanFile{Path: path, Kind: "package", Module: name, SHA256: digest([]byte(source)), Bytes: int64(len(source))})
		if name == "Invalid" {
			put(t, filepath.Join(output, "purescript", path), source)
		}
	}
	// Import parsing must not supersede the existing byte-validation failures.
	// Parallel reads must also report the first failure in plan order.
	for _, missing := range plan.Files[1:] {
		if _, err := readPackagePlanFiles(output, plan); err == nil || err.Error() != "published Go package file changed: "+missing.Path {
			t.Fatalf("wrong first read failure for %s: %v", missing.Path, err)
		}
		put(t, filepath.Join(output, "purescript", missing.Path), "package "+overlayOriginalPackage+"\n")
	}
	snapshot, err := readPackagePlanFiles(output, plan)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := loadOverlayIndex(plan, snapshot, overlayOriginalPackage); err == nil {
		t.Fatal("accepted malformed imports from verified bytes")
	}
}
