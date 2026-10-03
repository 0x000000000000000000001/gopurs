// After npm run build:native: npm run test:cli
import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { cpSync, existsSync, mkdirSync, readFileSync, readdirSync, rmSync, writeFileSync } from "node:fs";
import { delimiter, join } from "node:path";
import { test } from "node:test";
import { fileURLToPath } from "node:url";
import { corePackages, createWorkspace, prepareFixture } from "./test-workspace.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const launcher = join(root, "bin/gopurs");
const testRust = process.env.GOPURS_TEST_RUST === "1";
const rustCandidate = process.env.GOPURS_TEST_RUST_BINARY;
const environment = {
  ...process.env,
  PATH: join(root, "node_modules/.bin") + delimiter + process.env.PATH,
  GOWORK: "off",
};
// An inherited profile path could turn an otherwise successful build into an
// unrelated failure. Each case below explicitly selects compiler concurrency.
delete environment.GOPURS_ALLOC_PROFILE;

function run(command, args, cwd, env = environment) {
  const result = spawnSync(command, args, {
    cwd, env, encoding: "utf8", timeout: 30_000, maxBuffer: 4 * 1024 * 1024,
  });
  assert.ifError(result.error);
  assert.equal(result.signal, null, result.stdout + result.stderr);
  return result;
}

function generatedFiles(output, directory = output, files = {}) {
  for (const entry of readdirSync(directory, { withFileTypes: true })) {
    const path = join(directory, entry.name);
    if (entry.isDirectory()) generatedFiles(output, path, files);
    else if (entry.name.endsWith(".go") || entry.name === "go.mod") {
      files[path.slice(output.length + 1)] = readFileSync(path);
    }
  }
  return files;
}

test("CLI reports failures and preserves successful output across backends", async t => {
  assert.ok(existsSync(join(root, "bin/gopurs-native")), "run npm run build:native first");
  if (testRust) assert.ok(existsSync(rustCandidate ?? join(root, "bin/gopurs-rust")), "run npm run build:rust first");
  const workspace = createWorkspace();
  t.after(() => rmSync(workspace, { recursive: true, force: true }));
  const fixture = prepareFixture(root, workspace,
    join(root, "tests/passing/FFIIntegerReturns.purs"), 0, corePackages(root));
  const output = join(fixture.directory, "output");
  const ffiPath = join(fixture.directory, "src/Main.go");
  const originalFfi = readFileSync(ffiPath, "utf8");
  const compiled = run("spago", ["build", "-q"], fixture.directory);
  assert.equal(compiled.status, 0, compiled.stdout + compiled.stderr);
  assert.ok(Array.isArray(JSON.parse(readFileSync(join(output, "Main/corefn.json"), "utf8")).typeTable),
    "the fixture requires the TAST-capable purs fork");
  const inputs = join(workspace, "tast-output");
  cpSync(output, inputs, { recursive: true });

  function reset() {
    rmSync(output, { recursive: true, force: true });
    cpSync(inputs, output, { recursive: true });
    writeFileSync(ffiPath, originalFfi);
  }

  let reference;
  for (const mode of [
    { name: "JS", js: "1", jobs: "1", pipeline: "0" },
    { name: "native sequential", js: "0", jobs: "1", pipeline: "0" },
    { name: "native parallel", js: "0", jobs: "8", pipeline: "1" },
    ...(testRust ? [
      { name: "Rust sequential", js: "0", rust: "1", jobs: "1", pipeline: "0" },
      { name: "Rust parallel", js: "0", rust: "1", jobs: "8", pipeline: "1" },
    ] : []),
  ]) {
    const env = {
      ...environment,
      GOPURS_JS: mode.js,
      GOPURS_RUST: mode.rust ?? "0",
      GOPURS_JOBS: mode.jobs,
      GOPURS_PBO_JOBS: mode.jobs,
      GOPURS_PREPARE_JOBS: mode.jobs,
      GOPURS_EMIT_JOBS: mode.jobs,
      GOPURS_PIPELINE: mode.pipeline,
    };
    const compile = () => run(mode.rust && rustCandidate ? rustCandidate : launcher, ["--main", "Main"], fixture.directory, env);

    await t.test(`${mode.name}: success exits zero with byte-identical Go`, () => {
      reset();
      const result = compile();
      assert.equal(result.status, 0, result.stdout + result.stderr);
      assert.ok(result.stderr.includes(`[gopurs] workers: prepare=${mode.jobs}, pbo=${mode.jobs}, emit=${mode.jobs}, pipeline=${mode.pipeline === "1"}`), result.stderr);
      assert.doesNotMatch(result.stderr, /\(failed\)|\[gopurs\] error:/);
      assert.ok(existsSync(join(output, "main/main.go")));
      assert.ok(existsSync(join(output, "purescript/Main_ffi.go")));
      const files = generatedFiles(output);
      if (reference) assert.deepEqual(files, reference);
      else {
        reference = files;
        const execute = run("go", ["run", "./main/main.go"], output);
        assert.equal(execute.status, 0, execute.stdout + execute.stderr);
        assert.match(execute.stdout, /Done/);
      }
    });

    for (const scenario of [
      {
        name: "TAST loading",
        prepare: () => rmSync(output, { recursive: true }),
        diagnostic: /output/,
        phase: "load TAST + sort",
      },
      {
        name: "runtime write",
        prepare: () => mkdirSync(join(output, "gopurs_runtime/runtime.go"), { recursive: true }),
        diagnostic: /gopurs_runtime\/runtime\.go/,
        phase: "runtime",
      },
      {
        name: "module write with optimizer workers",
        prepare: () => mkdirSync(join(output, "purescript/Control_Bind.go"), { recursive: true }),
        diagnostic: /Control_Bind\.go/,
        phase: "optimize + emit",
      },
      {
        name: "invalid FFI",
        prepare: () => writeFileSync(ffiPath, "func Broken("),
        diagnostic: /FFI error in module Main \(.*Main\.go\)/,
        phase: "optimize + emit",
      },
      {
        name: "entry point write",
        prepare: () => mkdirSync(join(output, "main/main.go"), { recursive: true }),
        diagnostic: /main\/main\.go/,
        phase: "entry points",
      },
    ]) {
      await t.test(`${mode.name}: ${scenario.name} exits one with the original diagnostic`, () => {
        reset();
        scenario.prepare();
        const result = compile();
        assert.equal(result.status, 1, result.stdout + result.stderr);
        assert.equal(result.stdout, "", "diagnostics belong on stderr");
        const errors = result.stderr.split("\n").filter(line => line.startsWith("[gopurs] error: "));
        assert.equal(errors.length, 1, result.stderr);
        assert.match(errors[0], scenario.diagnostic);
        assert.ok(result.stderr.includes(`[gopurs] ${scenario.phase}:`), result.stderr);
        assert.ok(result.stderr.includes("(failed)"), result.stderr);
        assert.doesNotMatch(result.stderr, /deadlock|panic:/);
      });
    }
  }
});
