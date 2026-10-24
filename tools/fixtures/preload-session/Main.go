package Test_PreloadSessionNative

import (
	"encoding/json"
	"fmt"
	cache "gopurs/output/gopurs_build_cache"
	"os"
	"path/filepath"
)

// Test-only private fixtures. No compiler installation or application is used.
var fixtureRoot, previousCwd, previousSnapshot string

func write(path, content string) {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		panic(err)
	}
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		panic(err)
	}
}
func exchange(request cache.Request) cache.Response {
	data, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	var result cache.Response
	if err = json.Unmarshal([]byte(cache.Exchange(string(data))), &result); err != nil {
		panic(err)
	}
	if result.Status == "error" {
		panic(fmt.Errorf("fixture exchange: %s", result.Reason))
	}
	return result
}
func SetupFixture(_ any) any {
	var err error
	previousCwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	previousSnapshot = os.Getenv("GOPURS_COMPILER_SNAPSHOT")
	fixtureRoot, err = os.MkdirTemp("", "gopurs-preload-session-")
	if err != nil {
		panic(err)
	}
	fixtureRoot, err = filepath.EvalSymlinks(fixtureRoot)
	if err != nil {
		panic(err)
	}
	output := filepath.Join(fixtureRoot, "output")
	snapshot := filepath.Join(fixtureRoot, "compiler")
	write(filepath.Join(snapshot, "artifact"), "private fixture compiler bytes")
	write(filepath.Join(snapshot, "compiler.json"), `{"schema":1,"host":"go","resources":[{"name":"compiler","kind":"file","path":"artifact"}]}`)
	if err = os.Setenv("GOPURS_COMPILER_SNAPSHOT", snapshot); err != nil {
		panic(err)
	}
	if err = os.Chdir(fixtureRoot); err != nil {
		panic(err)
	}
	write(filepath.Join(output, "Main/corefn.json"), `{"typeTable":[],"moduleName":["Main"],"value":1}`)
	compiler, err := cache.CompilerInputs()
	if err != nil {
		panic(err)
	}
	options := []cache.Setting{{Name: "main", Value: "Nothing"}, {Name: "ffi", Value: ""}, {Name: "rewriteLimit", Value: "10000"}, {Name: "prepareJobs", Value: "8"}, {Name: "optimizerJobs", Value: "8"}, {Name: "emissionJobs", Value: "8"}, {Name: "pipeline", Value: "true"}, {Name: "go-layout", Value: "gopurs/monolithic-v1"}, {Name: "module-lazy", Value: "1"}, {Name: "input-layout", Value: "packed-v1"}}
	inputs := []cache.Input{{Name: "tast", Kind: "tast", Path: output}}
	request := cache.Request{Schema: 1, Operation: "snapshot", Spec: cache.Spec{Workspace: fixtureRoot, Output: output, Compiler: compiler, Options: options, Inputs: inputs}}
	observed := exchange(request)
	if observed.Status != "snapshot" {
		panic("fixture snapshot missing")
	}
	request.Operation = "publish"
	request.ExpectedKey = observed.Snapshot.Key
	if exchange(request).Status != "published" {
		panic("fixture publication failed")
	}
	recipe, _ := json.Marshal(map[string]any{"schema": 1, "inputs": inputs})
	write(filepath.Join(output, ".gopurs-cache/recipe.json"), string(recipe))
	return nil
}
func EditFixture(_ any) any {
	write(filepath.Join(fixtureRoot, "output/Main/corefn.json"), `{"typeTable":[],"moduleName":["Main"],"value":2}`)
	return nil
}
func ClearRecipe(_ any) any {
	if err := os.Remove(filepath.Join(fixtureRoot, "output/.gopurs-cache/recipe.json")); err != nil {
		panic(err)
	}
	return nil
}
func RemainingStages(_ any) int64 {
	paths, err := filepath.Glob(filepath.Join(fixtureRoot, "output/.gopurs-cache/stages/build-*"))
	if err != nil {
		panic(err)
	}
	return int64(len(paths))
}
func CleanupFixture(_ any) any {
	if err := os.Chdir(previousCwd); err != nil {
		panic(err)
	}
	if previousSnapshot == "" {
		os.Unsetenv("GOPURS_COMPILER_SNAPSHOT")
	} else {
		os.Setenv("GOPURS_COMPILER_SNAPSHOT", previousSnapshot)
	}
	if err := os.RemoveAll(fixtureRoot); err != nil {
		panic(err)
	}
	return nil
}
