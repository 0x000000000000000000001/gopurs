import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { existsSync, mkdirSync, mkdtempSync, realpathSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { delimiter, dirname, join } from "node:path";
import test from "node:test";
import { findOfficialCompiler } from "./release/official-compiler.mjs";

function fixture(t) {
  const root = mkdtempSync(join(tmpdir(), "gopurs-official-compiler-"));
  t.after(() => rmSync(root, { recursive: true, force: true }));
  const typed = join(root, "typed/purs");
  const invoked = join(root, "typed-invoked");
  const official = join(root, "node_modules/purescript/purs.bin");
  for (const file of [typed, official]) mkdirSync(dirname(file), { recursive: true });
  writeFileSync(typed, `#!/bin/sh\ntouch '${invoked}'\necho typed\n`, { mode: 0o755 });
  const env = { ...process.env, PATH: dirname(typed) + delimiter + process.env.PATH };
  delete env.GOPURS_OFFICIAL_PURS;
  const run = () => spawnSync(process.execPath, ["--input-type=module", "-e", `
    import { execFileSync } from "node:child_process";
    import { findOfficialCompiler } from ${JSON.stringify(new URL("./release/official-compiler.mjs", import.meta.url).href)};
    const compiler = findOfficialCompiler(${JSON.stringify(root)});
    process.stdout.write(execFileSync(compiler, ["--version"], { encoding: "utf8" }));
  `], { env, encoding: "utf8", timeout: 10000 });
  return { root, official, invoked, env, run };
}

test("the JavaScript oracle selects the installed stock frontend even with a typed purs on PATH", (t) => {
  const f = fixture(t);
  writeFileSync(f.official, "#!/bin/sh\necho stock\n", { mode: 0o755 });
  const result = f.run();
  assert.equal(result.status, 0, result.stderr);
  assert.equal(result.stdout.trim(), "stock");
  assert.equal(existsSync(f.invoked), false);
});

test("a missing stock frontend fails instead of falling back to the typed purs on PATH", (t) => {
  const f = fixture(t);
  const result = f.run();
  assert.notEqual(result.status, 0);
  assert.match(result.stderr, /Set GOPURS_OFFICIAL_PURS/);
  assert.match(result.stderr, /PATH is not used/);
  assert.equal(existsSync(f.invoked), false);
});

test("the stock frontend supports explicit and environment paths without requiring npm install scripts", (t) => {
  const f = fixture(t);
  const configured = join(f.root, "nix-purs");
  writeFileSync(configured, "#!/bin/sh\necho configured stock\n", { mode: 0o755 });
  f.env.GOPURS_OFFICIAL_PURS = configured;
  const result = f.run();
  assert.equal(result.status, 0, result.stderr);
  assert.equal(result.stdout.trim(), "configured stock");
  assert.equal(findOfficialCompiler(f.root, configured), realpathSync(configured));
  assert.throws(() => findOfficialCompiler(f.root, join(f.root, "missing-purs")), /No executable standard PureScript compiler/);
  assert.equal(existsSync(f.invoked), false);
});
