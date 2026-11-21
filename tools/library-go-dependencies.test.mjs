import assert from "node:assert/strict";
import { execFileSync } from "node:child_process";
import { existsSync, mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import test from "node:test";
import { libraryGoRequirements, prepareLibraryGoDependencies } from "./library-go-dependencies.mjs";

function fixture(t) {
  const root = mkdtempSync(join(tmpdir(), "gopurs-go-dependencies-"));
  t.after(() => rmSync(root, { recursive: true, force: true }));
  const output = join(root, "output");
  const port = (name, require) => {
    const path = join(root, name);
    mkdirSync(path);
    writeFileSync(join(path, "go-dependencies.json"), JSON.stringify({ schema: 1, require }));
    return path;
  };
  const seed = () => {
    mkdirSync(output);
    writeFileSync(join(output, "go.mod"), "// Application settings\nmodule example.test/application\n\ngo 1.22\n\nrequire example.test/kept v1.2.3\nreplace example.test/kept => ../local\nexclude example.test/excluded v1.0.0\n");
    writeFileSync(join(output, "go.sum"), "existing checksums\n");
  };
  const manifests = () => ["go.mod", "go.sum"].map(name => existsSync(join(output, name)) ? readFileSync(join(output, name), "utf8") : null);
  const calls = [];
  // Use the real Go parser/editor offline, but simulate registry resolution
  // and downloads. This verifies merge/preservation without a network oracle.
  const runGo = (args, options) => {
    calls.push(args);
    const execute = args => execFileSync("go", args, {
      ...options, encoding: "utf8", env: { ...options.env, GOPROXY: "off", GOSUMDB: "off" },
    });
    if (args[0] === "get") {
      for (const dependency of args.slice(1)) execute(["mod", "edit", `-require=${dependency}`]);
      return "";
    }
    if (args[0] === "mod" && args[1] === "download") {
      writeFileSync(join(output, "go.sum"), (manifests()[1] ?? "") + "downloaded checksums\n");
      return "";
    }
    return execute(args);
  };
  return { output, root, port, seed, manifests, calls, runGo };
}

test("no declared dependencies performs no writes or Go calls", t => {
  const f = fixture(t);
  assert.deepEqual(prepareLibraryGoDependencies(f.output, [f.root], { runGo: () => assert.fail("unexpected Go invocation") }), []);
  assert.equal(existsSync(f.output), false);
});

test("exact manifests merge while preserving application settings and checksums", t => {
  const f = fixture(t);
  f.seed();
  const a = f.port("a", { "example.test/shared": "v1.0.0" });
  const b = f.port("b", { "example.test/shared": "v1.0.0", "example.test/extra": "v2.0.0+incompatible" });
  assert.deepEqual(prepareLibraryGoDependencies(f.output, [a, b], { runGo: f.runGo }), ["example.test/extra@v2.0.0+incompatible", "example.test/shared@v1.0.0"]);
  const [mod, sum] = f.manifests();
  for (const text of ["// Application settings", "module example.test/application", "go 1.22", "example.test/kept v1.2.3", "replace example.test/kept => ../local", "exclude example.test/excluded v1.0.0"]) assert.ok(mod.includes(text), text);
  assert.equal(sum, "existing checksums\ndownloaded checksums\n");
  const gets = f.calls.filter(args => args[0] === "get");
  assert.deepEqual(gets, [["get", "example.test/extra@v2.0.0+incompatible", "example.test/shared@v1.0.0"]]);
  prepareLibraryGoDependencies(f.output, [a, b], { runGo: f.runGo });
  assert.equal(f.calls.filter(args => args[0] === "get").length, 1, "matching versions do not trigger another resolution");
});

test("conflicting ports and application pins fail before resolution", t => {
  const f = fixture(t);
  f.seed();
  const original = f.manifests();
  const a = f.port("a", { "example.test/kept": "v1.0.0" });
  const b = f.port("b", { "example.test/kept": "v2.0.0" });
  assert.throws(() => prepareLibraryGoDependencies(f.output, [a, b], { runGo: f.runGo }), /Conflicting Go dependency/);
  assert.equal(f.calls.length, 0);
  assert.throws(() => prepareLibraryGoDependencies(f.output, [a], { runGo: f.runGo }), /application pins v1.2.3/);
  assert.deepEqual(f.manifests(), original);
  assert.ok(f.calls.every(args => args[0] !== "get"));
});

test("transitive upgrades roll back both manifests", t => {
  const f = fixture(t);
  f.seed();
  const original = f.manifests();
  const a = f.port("a", { "example.test/added": "v1.0.0" });
  const runGo = (args, options) => {
    const result = f.runGo(args, options);
    if (args[0] === "get") {
      f.runGo(["mod", "edit", "-require=example.test/kept@v1.3.0"], options);
      writeFileSync(join(f.output, "go.sum"), "changed sums\n");
    }
    return result;
  };
  assert.throws(() => prepareLibraryGoDependencies(f.output, [a], { runGo }), /would change application requirement/);
  assert.deepEqual(f.manifests(), original);
});

test("download failure restores existing files and removes newly created files", t => {
  const f = fixture(t);
  const a = f.port("a", { "example.test/added": "v1.0.0" });
  const runGo = (args, options) => {
    if (args[1] === "download") throw new Error("offline module unavailable");
    return f.runGo(args, options);
  };
  assert.throws(() => prepareLibraryGoDependencies(f.output, [a], { runGo }), /offline module unavailable/);
  assert.deepEqual(f.manifests(), [null, null]);
  rmSync(f.output, { recursive: true });
  f.seed();
  const original = f.manifests();
  assert.throws(() => prepareLibraryGoDependencies(f.output, [a], { runGo }), /offline module unavailable/);
  assert.deepEqual(f.manifests(), original);
});

test("unversioned, latest and malformed manifests are rejected", t => {
  const f = fixture(t);
  for (const [index, version] of ["latest", "main", "v1", "v1.2.3 && command"].entries()) {
    const port = f.port(`invalid-${index}`, { "example.test/module": version });
    assert.throws(() => libraryGoRequirements([port]), /exact Go module version/);
  }
  const port = f.port("schema", {});
  writeFileSync(join(port, "go-dependencies.json"), JSON.stringify({ schema: 2, require: {} }));
  assert.throws(() => libraryGoRequirements([port]), /Invalid Go dependency manifest/);
  writeFileSync(join(port, "go-dependencies.json"), "null");
  assert.throws(() => libraryGoRequirements([port]), /Invalid Go dependency manifest/);
  writeFileSync(join(port, "go-dependencies.json"), "{");
  assert.throws(() => libraryGoRequirements([port]), /Invalid Go dependency manifest.*go-dependencies.json/);
  assert.throws(() => libraryGoRequirements([join(f.root, "missing-port")]), /not a directory/);
});
