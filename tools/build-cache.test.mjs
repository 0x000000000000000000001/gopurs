import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { execFileSync } from "node:child_process";
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import { test } from "node:test";
import { exchange as jsExchange } from "../src/Gopurs/BuildCache.js";
import { resolveBuildCacheHelper } from "./build-cache.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const env = { ...process.env, GOWORK: "off" };
const run = (command, args, cwd, input) => execFileSync(command, args, {
  cwd, input, encoding: "utf8", env, timeout: 120_000,
  maxBuffer: 16 * 1024 * 1024, stdio: ["pipe", "pipe", "pipe"],
});

test("cache integrity, invalidation, publication and process death under the Go race detector", () => {
  const result = run("go", ["test", "-race", "-count=1", "-timeout=90s", "-v", "./..."], join(root, "tools/build-cache"));
  for (const name of ["ContentInvalidation", "IntegrityAndVersionMisses", "WorkspaceIsolation",
    "OwnershipTidyAndWriteElision", "WireArtifacts", "ConcurrentPublishers", "InterruptedPublicationAndKernelLock"]) {
    assert.ok(result.includes(`--- PASS: Test${name} `), result);
  }
});

test("JS and directly linked Go bridges exchange a content-verified persistent manifest", t => {
  const workspace = mkdtempSync(join(tmpdir(), "gopurs-build-cache-hosts-"));
  t.after(() => rmSync(workspace, { recursive: true, force: true }));
  const put = (path, content) => {
    mkdirSync(dirname(join(workspace, path)), { recursive: true });
    writeFileSync(join(workspace, path), content);
  };
  // Compile the actual FFI and the exact same copied package used by bootstrap.
  put("native/go.mod", "module gopurs/output\n\ngo 1.22\n");
  run(process.execPath, [join(root, "tools/prepare-native-output.mjs"), join(workspace, "native")], root);
  put("native/bridge/bridge.go", readFileSync(join(root, "src/Gopurs/BuildCache.go"), "utf8")
    .replace("package Gopurs_BuildCache", "package bridge"));
  put("native/main.go", `package main
import ("fmt"; "io"; "os"; "gopurs/output/bridge")
func main() { b, err := io.ReadAll(os.Stdin); if err != nil { panic(err) }; fmt.Print(bridge.Exchange(string(b), nil)) }
`);
  const native = join(workspace, "linked-cache");
  run("go", ["build", "-o", native, "."], join(workspace, "native"));
  const nativeExchange = request => run(native, [], workspace, request);
  put("compiler", "identical artifact for this protocol comparison");
  put("output/Main/corefn.json", '{"moduleName":["Main"],"typeTable":[]}');
  put("stage/Main.go", "package purescript\n// revision one\n");
  const request = {
    schema: 1, operation: "snapshot",
    spec: {
      workspace, output: "output",
      compiler: [{ name: "compiler", kind: "file", path: "compiler" }],
      options: [{ name: "main", value: "Main" }],
      inputs: [{ name: "tast", kind: "tast", path: "output" }],
    },
    outputs: [{ path: "purescript/Main.go", kind: "module", module: "Main", source: "stage/Main.go" }],
  };
  const exchange = req => jsExchange(JSON.stringify(req))();
  const js = exchange(request);
  assert.equal(js, nativeExchange(JSON.stringify(request)), "wire responses must be byte-identical");
  request.operation = "publish";
  request.expectedKey = JSON.parse(js).snapshot.key;
  const published = JSON.parse(exchange(request));
  assert.equal(published.status, "published");
  request.operation = "lookup";
  const found = JSON.parse(nativeExchange(JSON.stringify(request)));
  assert.equal(found.status, "hit");
  assert.deepEqual(found.manifest, published.manifest);
  assert.equal(exchange(request), nativeExchange(JSON.stringify(request)));
  const file = readFileSync(join(workspace, "output/purescript/Main.go"));
  assert.equal(createHash("sha256").update(file).digest("hex"), found.manifest.outputs[0].object.sha256);
  put("stage/Main.go", "package purescript\n// revision two\n");
  request.operation = "publish";
  assert.equal(JSON.parse(nativeExchange(JSON.stringify(request))).status, "published");
  request.operation = "lookup";
  assert.equal(JSON.parse(exchange(request)).status, "hit");
  assert.match(readFileSync(join(workspace, "output/purescript/Main.go"), "utf8"), /revision two/);
  assert.equal(resolveBuildCacheHelper(), resolveBuildCacheHelper(), "helper is reused by source content");
});
