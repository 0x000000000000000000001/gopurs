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
	f := newPartitionFixture(t)
	for _, suffix := range []string{"-new", "-newer"} {
		f.stage = t.TempDir()
		source := partitionSource(suffix)
		f.request.Files[0].Source = filepath.Join(f.stage, "Main.go")
		f.request.Files[0].SHA256 = digest([]byte(source))
		put(t, f.request.Files[0].Source, source)
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
		request, _ := json.Marshal(f.request)
		put(t, filepath.Join(f.output, ".gopurs-cache/go-packages/plan-v1/request.json"), string(request))
	}
}

func TestPackageSectionsRejectUnprovedChanges(t *testing.T) {
	for name, edit := range map[string]func(string) string{
		"type": func(s string) string { return s + "\ntype Extra struct{}\n" },
		"init": func(s string) string { return s + "\nfunc init() {}\n" },
		"initialized-variable": func(s string) string {
			return strings.Replace(s, "var privateCache string", "var privateCache = \"initial\"", 1)
		},
		"signature":    func(s string) string { return strings.Replace(s, "func Get_Main()", "func Get_Main(extra int)", 1) },
		"graph":        func(s string) string { return strings.Replace(s, "Helper() + Other()", "Other()", 1) },
		"private-edge": func(s string) string { return strings.Replace(s, "Helper() + Other()", "Helper() + privateCache", 1) },
	} {
		t.Run(name, func(t *testing.T) {
			f := newPartitionFixture(t)
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
	other := testBundle(partitionPackage, "", sectionFor("Main.go", "func Get_Main() string { return \"duplicate\" }\n"))
	replaceFixturePlanFile(t, f, partitionPackage+"/"+packageBundleName, other)
	record, reason, err := preparePackageOverlay(f.stage, &f.state, []string{"Main"}, &f.request)
	if err != nil || record != nil || reason == "" {
		t.Fatalf("accepted ambiguous placement: %+v %s %v", record, reason, err)
	}
}
