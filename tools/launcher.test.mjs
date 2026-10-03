import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { delimiter, join } from "node:path";
import { test } from "node:test";

// Execute the real launcher with instrumented hosts and deterministic memory
// detection. The contracts concern the environment actually passed to exec.
function fixture(t) {
  const root = mkdtempSync(join(tmpdir(), "gopurs-launcher-"));
  t.after(() => rmSync(root, { recursive: true, force: true }));
  const bin = join(root, "bin"), tools = join(root, "tools");
  mkdirSync(bin); mkdirSync(tools);
  copyFileSync(new URL("../bin/gopurs", import.meta.url), join(bin, "gopurs"));
  const keys = ["GOPURS_JOBS", "GOPURS_PREPARE_JOBS", "GOPURS_PBO_JOBS", "GOPURS_EMIT_JOBS", "GOPURS_PIPELINE", "GOGC", "GOMEMLIMIT"];
  for (const [host, path] of [["go", join(bin, "gopurs-native")], ["rust", join(bin, "gopurs-rust")], ["js", join(tools, "node")]]) {
    writeFileSync(path, `#!${process.execPath}\nconsole.log(JSON.stringify({host:${JSON.stringify(host)},args:process.argv.slice(2),env:Object.fromEntries(${JSON.stringify(keys)}.filter(k=>process.env[k]!==undefined).map(k=>[k,process.env[k]]))}));\nprocess.exit(Number(process.env.TEST_EXIT ?? 0));\n`, { mode: 0o755 });
  }
  writeFileSync(join(tools, "sysctl"), '#!/bin/sh\nprintf "%s\\n" "$TEST_MEMORY_BYTES"\n', { mode: 0o755 });
  const base = { ...process.env, PATH: tools + delimiter + process.env.PATH, TEST_MEMORY_BYTES: String(48 * 2 ** 30) };
  for (const key of Object.keys(base)) if (/^(GOPURS_|GOGC$|GOMEMLIMIT$)/.test(key)) delete base[key];
  return { root, run(host, extra = {}, args = ["--main", "Main", "argument with spaces"]) {
    const env = { ...base, GOPURS_RUST: host === "rust" ? "1" : "0", GOPURS_JS: host === "js" ? "1" : "0", ...extra };
    const result = spawnSync("/bin/bash", [join(bin, "gopurs"), ...args], { env, encoding: "utf8", timeout: 10000 });
    assert.ifError(result.error);
    return { ...result, value: result.stdout ? JSON.parse(result.stdout) : null };
  } };
}

test("all hosts receive automatic worker limits; automatic GC settings belong only to Go", t => {
  const f = fixture(t);
  for (const host of ["go", "rust", "js"]) {
    const result = f.run(host);
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.value.host, host);
    assert.deepEqual(result.value.env, { GOPURS_PREPARE_JOBS: "8", GOPURS_PBO_JOBS: "8",
      ...(host === "go" ? { GOGC: "off", GOMEMLIMIT: "10GiB" } : {}) });
    assert.deepEqual(result.value.args, [
      ...(host === "js" ? ["--stack-size=65536", join(f.root, "bin/gopurs.js")] : []),
      "--main", "Main", "argument with spaces",
    ]);
  }
});

test("automatic parallelism has the same 32 GiB boundary for every host", t => {
  const f = fixture(t);
  for (const host of ["go", "rust", "js"]) for (const gib of [16, 31, 32]) {
    const result = f.run(host, { TEST_MEMORY_BYTES: String(gib * 2 ** 30) });
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.value.env.GOPURS_PBO_JOBS, gib >= 32 ? "8" : undefined);
    assert.equal(result.value.env.GOPURS_PREPARE_JOBS, gib >= 32 ? "8" : undefined);
    assert.equal(result.value.env.GOGC, gib >= 32 && host === "go" ? "off" : undefined);
  }
});

test("explicit worker limits and pipeline switches survive every host selection", t => {
  const f = fixture(t);
  const settings = { GOPURS_JOBS: "5", GOPURS_PREPARE_JOBS: "3", GOPURS_PBO_JOBS: "1", GOPURS_EMIT_JOBS: "4", GOPURS_PIPELINE: "0" };
  for (const host of ["go", "rust", "js"]) {
    const result = f.run(host, settings);
    assert.equal(result.status, 0, result.stderr);
    for (const [key, value] of Object.entries(settings)) assert.equal(result.value.env[key], value);
  }
});

test("explicit GC policies are preserved without disabling automatic parallelism", t => {
  const f = fixture(t);
  for (const host of ["go", "rust", "js"]) for (const policy of [{ GOGC: "100" }, { GOMEMLIMIT: "2GiB" }, { GOGC: "50", GOMEMLIMIT: "3GiB" }]) {
    const result = f.run(host, policy);
    assert.equal(result.status, 0, result.stderr);
    assert.deepEqual(result.value.env, { GOPURS_PREPARE_JOBS: "8", GOPURS_PBO_JOBS: "8", ...policy });
  }
});

test("host exit status is forwarded", t => {
  const f = fixture(t);
  for (const host of ["go", "rust", "js"]) assert.equal(f.run(host, { TEST_EXIT: "23" }).status, 23);
});

test("conflicting host selectors and a missing Rust executable fail before dispatch", t => {
  const f = fixture(t);
  const conflict = f.run("rust", { GOPURS_JS: "1" });
  assert.equal(conflict.status, 1); assert.equal(conflict.stdout, "");
  assert.match(conflict.stderr, /select only one compiler host/);
  rmSync(join(f.root, "bin/gopurs-rust"));
  const missing = f.run("rust");
  assert.equal(missing.status, 1); assert.equal(missing.stdout, "");
  assert.match(missing.stderr, /npm run build:rust/);
});
