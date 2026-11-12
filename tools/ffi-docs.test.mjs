import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { mkdtempSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import test from "node:test";

test("the exact documented Go snippets compile and preserve their observable contracts", t => {
  const documentation = readFileSync(new URL("../docs/ffi-idiomatic.md", import.meta.url), "utf8");
  const snippets = [...documentation.matchAll(/```go\n([\s\S]*?)\n```/g)].map(match => match[1]);
  assert.ok(snippets.length >= 5, "the guide must include its five executable Go examples");
  assert.doesNotMatch(snippets.join("\n"), /gopurs_runtime/);
  const work = mkdtempSync(join(tmpdir(), "gopurs-ffi-docs-"));
  let passed = false;
  t.after(() => {
    if (passed && process.env.GOPURS_TEST_KEEP_WORKSPACE !== "1") rmSync(work, { recursive: true, force: true });
    else console.log(`Documentation examples retained: ${work}`);
  });
  writeFileSync(join(work, "go.mod"), "module documented-ffi\n\ngo 1.23\n");
  // Each block remains its own source file, including any published imports.
  // Only supply package/import clauses abbreviated by the documentation.
  for (const [index, snippet] of snippets.entries()) {
    const missing = ["context", "errors", "os"].filter(
      name => new RegExp("\\b" + name + "\\.").test(snippet) && !snippet.includes('"' + name + '"'),
    );
    const imports = missing.map(name => 'import "' + name + '"\n').join("");
    const source = /^package fixture\n/.test(snippet)
      ? snippet.replace(/^package fixture\n/, "package fixture\n" + imports)
      : "package fixture\n" + imports + snippet;
    writeFileSync(join(work, `example_${index}.go`), source + "\n");
  }
  writeFileSync(join(work, "examples_test.go"), `package fixture
import ("context"; "errors"; "os"; "path/filepath"; "reflect"; "testing")

func TestHello(t *testing.T) {
  if value, err := Hello("Ada"); value != "hello Ada" || err != nil { t.Fatal(value, err) }
  if value, err := Hello("error"); value != "" || err == nil || err.Error() != "expected error" { t.Fatal(value, err) }
}

func TestPending(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background()); cancel()
  if value, err := Pending(ctx); value != "" || err != context.Canceled { t.Fatal(value, err) }
}

func TestOpenAndClose(t *testing.T) {
  path := filepath.Join(t.TempDir(), "resource")
  if file, cleanup, err := Open(path); file != nil || cleanup != nil || !errors.Is(err, os.ErrNotExist) {
    t.Fatalf("failed acquisition: file=%v, cleanup=%v, error=%v", file, cleanup != nil, err)
  }
  if err := os.WriteFile(path, []byte("resource"), 0600); err != nil { t.Fatal(err) }
  file, cleanup, err := Open(path)
  if err != nil || file == nil || cleanup == nil { t.Fatal("successful acquisition did not supply ownership", err) }
  if file.Name() != path { t.Fatal("handle identity lost") }
  bytes := make([]byte, 8)
  if _, err := file.Read(bytes); err != nil || string(bytes) != "resource" { t.Fatal(string(bytes), err) }
  cleanup()
  if _, err := file.Read(bytes); !errors.Is(err, os.ErrClosed) { t.Fatal("cleanup did not close the resource", err) }
  file, cleanup, err = Open(path)
  if err != nil { t.Fatal(err) }
  t.Cleanup(cleanup)
  if err := Close(file); err != nil { t.Fatal(err) }
  if _, err := file.Read(bytes); !errors.Is(err, os.ErrClosed) { t.Fatal("explicit Close did not close the resource", err) }
}

func TestRunCallback(t *testing.T) {
  var values []string
  RunCallback(func(value string) { values = append(values, value) })
  if !reflect.DeepEqual(values, []string{"typed"}) { t.Fatal(values) }
}

func TestChildren(t *testing.T) {
  values := []Child{{Active: false, Label: ""}, {Active: true, Label: "child"}}
  result := Children(values)
  if !reflect.DeepEqual(result, values) || &result[0] != &values[0] { t.Fatal("ordinary Go identity or values changed") }
  if Children(nil) != nil { t.Fatal("nil changed to an empty slice") }
  if result := Children([]Child{}); result == nil || len(result) != 0 { t.Fatal("empty slice changed") }
}
`);
  const result = spawnSync("go", ["test", "-race", "-count=1", "-timeout=20s", "-v", "./..."], {
    cwd: work, env: { ...process.env, GOWORK: "off" }, encoding: "utf8", timeout: 60000,
  });
  writeFileSync(join(work, "test.log"), (result.stdout ?? "") + (result.stderr ?? ""));
  assert.ifError(result.error);
  assert.equal(result.status, 0, result.stdout + result.stderr);
  console.log(`${snippets.length} exact Go examples, five behavior checks under -race`);
  passed = true;
});
