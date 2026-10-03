import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { cpSync, existsSync, readFileSync, readdirSync, rmSync, writeFileSync } from "node:fs";
import { delimiter, join } from "node:path";
import { fileURLToPath } from "node:url";
import { test } from "node:test";
import { corePackages, createWorkspace, prepareFixture } from "./test-workspace.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const launcher = join(root, "bin/gopurs");
const candidate = process.env.GOPURS_TEST_RUST_BINARY;
const enabled = process.env.GOPURS_TEST_RUST === "1";
const environment = { ...process.env,
  PATH: join(root, "node_modules/.bin") + delimiter + process.env.PATH,
  GOWORK: "off", GOPURS_PBO_JOBS: "4", GOPURS_PREPARE_JOBS: "4", GOPURS_EMIT_JOBS: "4",
};
delete environment.GOPURS_ALLOC_PROFILE;

function run(command, args, cwd, env, log) {
  const result = spawnSync(command, args, { cwd, env, encoding: "utf8", timeout: 120000, maxBuffer: 16 * 1024 * 1024 });
  writeFileSync(log + ".stdout", result.stdout ?? "");
  writeFileSync(log + ".stderr", result.stderr ?? "");
  assert.ifError(result.error);
  assert.equal(result.status, 0, `${command}: ${result.stdout}\n${result.stderr}`);
  return result;
}

function generatedFiles(directory, prefix = "") {
  return readdirSync(join(directory, prefix), { withFileTypes: true }).sort((a, b) => a.name.localeCompare(b.name)).flatMap(entry => {
    const path = join(prefix, entry.name);
    return entry.isDirectory() ? generatedFiles(directory, path)
      : path.endsWith(".go") || path === "go.mod" ? [[path, readFileSync(join(directory, path))]] : [];
  });
}

test("Rust-hosted gopurs preserves Go sources and execution on native regression fixtures", {
  skip: enabled ? false : "run npm run test:rust after building the Rust compiler",
}, async t => {
  assert.ok(existsSync(candidate ?? join(root, "bin/gopurs-rust")), "run npm run build:rust first");
  const workspace = createWorkspace(), packages = corePackages(root);
  let passed = false, completed = 0;
  t.after(() => {
    if (passed && process.env.GOPURS_TEST_KEEP_WORKSPACE !== "1") rmSync(workspace, { recursive: true, force: true });
    else console.log(`Rust host test workspace retained: ${workspace}`);
  });
  const fixtures = ["CompilerHostStrings", "FFIIntegerReturns", "NativeRecordWorkers", "OwnedTrees", "JsonRecordPlan"];
  for (const [index, name] of fixtures.entries()) {
    await t.test(name, () => {
      const fixture = prepareFixture(root, workspace, join(root, "tests/passing", name + ".purs"), index, packages);
      const directory = fixture.directory, output = join(directory, "output"), logs = join(directory, "logs");
      run("spago", ["build", "-q"], directory, environment, join(logs, "frontend"));
      const inputs = join(directory, "frozen-output");
      cpSync(output, inputs, { recursive: true });
      let reference, expected;
      for (const host of ["js", "go", "rust"]) {
        rmSync(output, { recursive: true, force: true });
        cpSync(inputs, output, { recursive: true });
        for (const cache of [".purmeta", ".cache"]) rmSync(join(directory, cache), { recursive: true, force: true });
        const env = { ...environment, GOPURS_JS: host === "js" ? "1" : "0", GOPURS_RUST: host === "rust" ? "1" : "0" };
        run(host === "rust" && candidate ? candidate : launcher, ["--main", "Main"], directory, env, join(logs, host + "-compile"));
        const files = generatedFiles(output);
        assert(files.length > 0);
        if (reference) assert.deepEqual(files, reference, `${name}: ${host} generated different Go`);
        else reference = files;
        // Execute every variant's generated application, not just the oracle.
        run("go", ["build", "-o", "test-app", "./main/main.go"], output, env, join(logs, host + "-build"));
        const result = run(join(output, "test-app"), [], output, env, join(logs, host + "-execute"));
        assert.equal(result.stderr, "");
        if (expected !== undefined) assert.equal(result.stdout, expected);
        else expected = result.stdout;
        assert.doesNotMatch(result.stdout, /Fail/);
      }
      completed++;
    });
  }
  passed = completed === fixtures.length;
});
