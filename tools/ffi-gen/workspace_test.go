package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestPackageFilesUseFrozenBytesAndMatchPureTransport(t *testing.T) {
	root := packageEvidence(t)
	fixture := bindingFixture()
	request := PackageFilesRequest{Format: "gopurs/go-package-files-v1", PackagePath: fixture.PackagePath,
		GoVersion: fixture.GoVersion, Arch: fixture.Arch, Output: filepath.Join(root, "emitted")}
	for i, file := range fixture.Files {
		_, module, _ := strings.Cut(file.Owner, "/")
		fixture.Files[i].Module = module
		input := filepath.Join(root, "input", file.Path)
		packagePut(t, root, "input/"+file.Path, file.Content)
		request.Files = append(request.Files, PackageSourceFile{Owner: file.Owner, Module: module, Path: file.Path,
			Source: input, SHA256: sourceHash(file.Content), Initializers: file.Initializers})
	}
	want, err := PackageSources(PackageRequest{BindRequest: fixture})
	if err != nil {
		t.Fatal(err)
	}
	encoded, _ := json.Marshal(request)
	response, err := PackageFiles(string(encoded))
	if err != nil {
		t.Fatal(err)
	}
	var got PackageFilesResult
	if err := json.Unmarshal([]byte(response), &got); err != nil {
		t.Fatal(err)
	}
	packagePut(t, root, "request.json", string(encoded))
	packagePut(t, root, "response.json", response)
	if got.PlanHash != want.PlanHash || !reflect.DeepEqual(got.Packages, want.Packages) || len(got.Files) != len(want.Files) {
		t.Fatal("file transport changed the plan")
	}
	for i, file := range got.Files {
		data, err := os.ReadFile(filepath.Join(request.Output, file.Path))
		if err != nil || string(data) != want.Files[i].Content || file.Path != want.Files[i].Path || sourceHash(string(data)) != file.SHA256 || len(data) != file.Bytes {
			t.Fatalf("file transport changed %s: %v", file.Path, err)
		}
	}
	// Reusing any existing destination must not overwrite its files.
	if _, err := EmitPackageFiles(request); !os.IsExist(err) {
		t.Fatalf("existing directory accepted: %v", err)
	}
	request.Output = filepath.Join(root, "changed-input")
	packagePut(t, root, "input/"+request.Files[0].Path, fixture.Files[0].Content+"\n// edited\n")
	if _, err := EmitPackageFiles(request); err == nil || !strings.Contains(err.Error(), "input changed") {
		t.Fatalf("mutated input accepted: %v", err)
	}
	if _, err := os.Stat(request.Output); !os.IsNotExist(err) {
		t.Fatal("failed resolution wrote its output")
	}
}
