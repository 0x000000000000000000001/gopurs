import assert from "node:assert/strict";
import { existsSync, mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { dirname, join } from "node:path";
import test from "node:test";
import {
  cleanGenerated, emittedManifest, encodedRustFlags, freezeTrainingOutput,
  rustFlagsEnvironment, walk,
} from "./pgo.mjs";

function temporary(t, name) {
  const directory = mkdtempSync(join(tmpdir(), `gopurs-pgo-${name}-`));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  return directory;
}

function put(path, contents) {
  mkdirSync(dirname(path), { recursive: true });
  writeFileSync(path, contents);
}

function typedCoreFn(modulePath) {
  return JSON.stringify({ moduleName: modulePath, modulePath, typeTable: ["Int"], dataDecls: [], classDecls: [] });
}

function relativeFiles(directory, base) {
  return walk(directory).map(path => path.slice(base.length + 1).split("\\").join("/"));
}

test("training freeze keeps Main, excludes Test.Main, and rewrites modulePath to frozen sources", t => {
  const project = temporary(t, "freeze");
  put(join(project, "src/Main.purs"), "module Main where\n");
  put(join(project, "src/Main.go"), "package main\n");
  put(join(project, "src/Test/Main.purs"), "module Test.Main where\n");
  put(join(project, "src/Test/Main.js"), "export const run = () => {};\n");
  put(join(project, "output/Main/corefn.json"), typedCoreFn("src/Main.purs"));
  put(join(project, "output/Test.Main/corefn.json"), typedCoreFn("src/Test/Main.purs"));
  const training = join(project, "training");
  const frozen = freezeTrainingOutput({ output: join(project, "output"), projectRoot: project, trainingDirectory: training });

  assert.deepEqual(frozen.modules.map(record => record.module), ["Main"]);
  assert.deepEqual(frozen.modules[0].sources, ["sources/Main/Main.purs", "sources/Main/Main.go"]);
  assert.deepEqual(frozen.excluded, ["Test.Main"]);
  assert.equal(JSON.parse(readFileSync(join(training, "output/Main/corefn.json"), "utf8")).modulePath,
    "sources/Main/Main.purs");
  assert.equal(existsSync(join(training, "output/Test.Main/corefn.json")), false);
  assert.equal(existsSync(join(training, "sources/Test.Main/Main.purs")), false);
  assert.ok(frozen.manifest.some(file => file.path === "output/Main/corefn.json"));
  assert.ok(frozen.manifest.some(file => file.path === "sources/Main/Main.go"));
  assert.ok(frozen.manifest.every(file => file.sha256.length === 64));
});

test("training freeze requires Main and does not accept only excluded modules", t => {
  const project = temporary(t, "missing-main");
  put(join(project, "src/Test/Main.purs"), "module Test.Main where\n");
  put(join(project, "output/Test.Main/corefn.json"), typedCoreFn("src/Test/Main.purs"));
  const training = join(project, "training");
  assert.throws(() => freezeTrainingOutput({
    output: join(project, "output"), projectRoot: project, trainingDirectory: training,
  }), /missing required module Main/);
  assert.equal(existsSync(join(training, "output/Test.Main/corefn.json")), false);
});

test("cleanup removes generated Go files and preserves frozen CoreFn, Main and sources", t => {
  const training = temporary(t, "clean");
  put(join(training, "output/Main/corefn.json"), typedCoreFn("sources/Main/Main.purs"));
  put(join(training, "sources/Main/Main.purs"), "module Main where\n");
  put(join(training, "sources/Main/Main.go"), "package main\n");
  // The generated entry point is output/main/main.go; on case-insensitive APFS
  // that directory aliases output/Main, so writing it must not remove CoreFn.
  put(join(training, "output/main/main.go"), "package main\n");
  put(join(training, "output/gopurs_runtime/runtime.go"), "package gopurs_runtime\n");
  put(join(training, "output/purescript/Main_ffi.go"), "package purescript\n");
  put(join(training, "output/go.mod"), "module gopurs/output\n");
  put(join(training, "output/go.sum"), "gopurs/output v0.0.0\n");
  put(join(training, ".purmeta/Main.json"), "{}");
  put(join(training, ".cache/Main.json"), "{}");
  const corefn = readFileSync(join(training, "output/Main/corefn.json"));

  const removed = cleanGenerated(training);
  assert.ok(removed >= 5, `expected generated files to be removed, got ${removed}`);
  assert.deepEqual(readFileSync(join(training, "output/Main/corefn.json")), corefn);
  assert.deepEqual(relativeFiles(join(training, "output"), training), ["output/Main/corefn.json"]);
  assert.equal(existsSync(join(training, "sources/Main/Main.go")), true);
  assert.equal(existsSync(join(training, ".purmeta")), false);
  assert.equal(existsSync(join(training, ".cache")), false);
});

test("profile flags keep spaces in paths inside one encoded argument", () => {
  const flags = ["-C", "profile-use=/tmp/My Profile/training.profdata"];
  const encoded = encodedRustFlags(flags);
  assert.equal(encoded, "-C\u001fprofile-use=/tmp/My Profile/training.profdata");
  assert.equal(encoded.split("\u001f")[1], "profile-use=/tmp/My Profile/training.profdata");

  const base = { PATH: "/bin" };
  const environment = rustFlagsEnvironment(base, flags);
  assert.equal(environment.CARGO_ENCODED_RUSTFLAGS, encoded);
  assert.deepEqual(base, { PATH: "/bin" });
  assert.equal(rustFlagsEnvironment(base, []).CARGO_ENCODED_RUSTFLAGS, undefined);
  assert.throws(() => encodedRustFlags(["-C", "bad\u001fflag"]), /separator/);
  assert.throws(() => encodedRustFlags([""]), /non-empty/);
});

test("emitted manifest compares generated Go and go.mod only", t => {
  const output = temporary(t, "emitted");
  put(join(output, "Main/corefn.json"), "{}");
  put(join(output, "main/main.go"), "package main\n");
  put(join(output, "go.mod"), "module gopurs/output\n");
  put(join(output, "go.sum"), "sum\n");
  const files = emittedManifest(output);
  assert.equal(files.length, 2);
  assert.ok(files.every(file => file.path.endsWith(".go") || file.path.endsWith("go.mod")));
  assert.ok(!files.some(file => file.path.endsWith("go.sum") || file.path.endsWith("corefn.json")));
});
