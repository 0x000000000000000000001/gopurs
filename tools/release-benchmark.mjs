#!/usr/bin/env node
import assert from "node:assert/strict";
import { execFileSync } from "node:child_process";
import { createHash } from "node:crypto";
import { chmodSync, constants, copyFileSync, cpSync, existsSync, mkdirSync, openSync, closeSync, readFileSync, readdirSync, readlinkSync, renameSync, rmSync, statfsSync, statSync, symlinkSync, writeFileSync } from "node:fs";
import { cpus, loadavg, platform, release, totalmem } from "node:os";
import { basename, delimiter, dirname, isAbsolute, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { CommandRunner } from "./command-runner.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const script = fileURLToPath(import.meta.url);
const [operation, ...args] = process.argv.slice(2);
const options = {};
for (let index = 0; index < args.length; index++) {
  const arg = args[index];
  if (["--pilot", "--retry-failed", "--keep-going"].includes(arg)) options[arg.slice(2)] = true;
  else if (["--out", "--archive", "--altbak", "--compiler", "--cases", "--min-free-gib"].includes(arg) && args[index + 1]) options[arg.slice(2)] = args[++index];
  else throw new Error(`Unknown or incomplete option: ${arg}`);
}
if (!["prepare", "run", "status"].includes(operation) || !options.out) {
  console.log("Usage:\n  node tools/release-benchmark.mjs prepare --out DIR --archive FROZEN_20261025 --altbak ALTBAK_DIR [--compiler DIR] [--cases NAME,...] [--pilot] [--min-free-gib 2]\n  node tools/release-benchmark.mjs run --out DIR [--cases NAME,...] [--retry-failed] [--keep-going]\n  node tools/release-benchmark.mjs status --out DIR\nNo frontend, compiler bootstrap, final Go build or application execution is timed. Prepare does not launch a compiler.");
  process.exit(operation === "--help" ? 0 : 2);
}
const directory = resolve(options.out);
const statePath = join(directory, "campaign.json");
const hash = value => createHash("sha256").update(value).digest("hex");
const fileHash = path => hash(readFileSync(path));
const readJson = path => JSON.parse(readFileSync(path, "utf8"));
function writeJson(path, value) {
  mkdirSync(dirname(path), { recursive: true });
  writeFileSync(path + ".next", JSON.stringify(value, null, 2) + "\n");
  renameSync(path + ".next", path);
}
function files(directory, prefix = "") {
  return readdirSync(directory, { withFileTypes: true }).sort((a, b) => a.name.localeCompare(b.name, "en")).flatMap(entry => {
    const name = prefix + entry.name;
    return entry.isDirectory() ? files(join(directory, entry.name), name + "/") : entry.isFile() ? [name] : [];
  });
}
function manifest(directory, selected = files(directory)) {
  return selected.sort().map(path => ({ path, bytes: statSync(join(directory, path)).size, sha256: fileHash(join(directory, path)) }));
}
function contained(directory, path) {
  assert(!isAbsolute(path), `Absolute manifest path: ${path}`);
  const result = resolve(directory, path);
  assert(result.startsWith(resolve(directory) + "/"), `Escaping manifest path: ${path}`);
  return result;
}
function verify(directory, expected) {
  for (const entry of expected) assert.equal(fileHash(contained(directory, entry.path)), entry.sha256, `Changed input: ${entry.path}`);
}
function copyManifest(from, to, entries) {
  for (const entry of entries) {
    const source = contained(from, entry.path), target = contained(to, entry.path);
    assert.equal(fileHash(source), entry.sha256, `Frozen input changed: ${source}`);
    mkdirSync(dirname(target), { recursive: true });
    copyFileSync(source, target, constants.COPYFILE_FICLONE);
  }
}
const capture = (command, args, cwd = root) => execFileSync(command, args, { cwd, encoding: "utf8", stdio: ["ignore", "pipe", "pipe"], maxBuffer: 4 * 1024 * 1024 }).trim();
const variants = ["js", "go", "dev"];
const orders = [["js", "go", "dev"], ["go", "dev", "js"], ["dev", "js", "go"], ["js", "go", "dev"]];
const select = rows => {
  if (!options.cases) return rows;
  const names = options.cases.split(",");
  assert(names.length && new Set(names).size === names.length, "Duplicate or empty case selection");
  for (const name of names) assert(rows.some(row => row.name === name), `Unknown corpus: ${name}`);
  return rows.filter(row => names.includes(row.name));
};
function space(minimum) {
  const disk = statfsSync(directory);
  assert(disk.bavail * disk.bsize >= minimum * 2 ** 30, `Less than ${minimum} GiB available; no measurement started`);
}
function baseline(altbak) {
  const readme = join(altbak, "README.md");
  const report = join(altbak, "docs/benchmark-results/2026-10-25-gopurs-compilation-refresh.md");
  const results = join(altbak, "docs/benchmark-results/2026-10-25-gopurs-compilation-refresh.json");
  const data = readJson(results);
  const text = readFileSync(readme, "utf8").split("#### ... to Go")[1]?.split("#### ... to Rust")[0];
  assert(text, "Cannot find the official Go compilation table in altbak README");
  for (const [name, row] of Object.entries(data.results)) {
    const line = text.split("\n").find(line => line.startsWith(`[${name}](`));
    assert(line, `README baseline row is missing: ${name}`);
    const cells = line.split("|").slice(1, 4).map(cell => {
      const match = cell.match(/~\s*([\d.]+)\s*(ms|s)/);
      assert(match, `Invalid README baseline cell: ${cell}`);
      return Number(match[1]) * (match[2] === "s" ? 1000 : 1);
    });
    for (const [index, variant] of variants.entries()) assert(Math.abs(cells[index] - row[variant].median_ms) < 0.51, `README/JSON baseline mismatch: ${name}/${variant}`);
  }
  return { files: [readme, report, results].map(path => ({ path, sha256: fileHash(path) })), machine: data.machine, protocol: data.protocol, rows: data.results, totals_ms: data.totals_ms };
}
function environment(state, variant) {
  const env = { ...process.env };
  for (const name of Object.keys(env)) if (/^(GOPURS_|PURUST_|PBO_|DEV_AUDIT_|NODE_|RUST|CARGO|GOGC$|GOMEMLIMIT$|GOMAXPROCS$|GODEBUG$|GOFLAGS$|GOCACHE$|GOMODCACHE$|GOWORK$|GOTOOLCHAIN$|GOPROXY$|TMPDIR$|TMP$|TEMP$|PPROF$)/.test(name)) delete env[name];
  Object.assign(env, { GOWORK: "off", GOPROXY: "off", GOTOOLCHAIN: "local", GOCACHE: join(directory, "go-cache"), GOMODCACHE: join(directory, "go-mod-cache"), TMPDIR: join(directory, "tmp"),
    GOPURS_JS: variant === "js" ? "1" : "0", GOPURS_RUST: "0", GOPURS_DAEMON: "0", GOPURS_RESIDENT: "0",
    GOPURS_GO_PACKAGES: "1", GOPURS_GO_PACKAGE_OVERLAY: "1", GOPURS_BUILD_MANIFEST: "1",
    GOPURS_JOBS: "8", GOPURS_PREPARE_JOBS: "8", GOPURS_PBO_JOBS: "8", GOPURS_EMIT_JOBS: "8", GOPURS_PIPELINE: "1" });
  env.PATH = [dirname(process.execPath), dirname(state.tools.go_path), process.env.PATH].join(delimiter);
  if (variant !== "js") Object.assign(env, { GOGC: "off", GOMEMLIMIT: "16GiB" });
  return env;
}
function generated(output) {
  return manifest(output, files(output).filter(path => !path.split("/").some(part => part.startsWith(".gopurs-cache")) && (path.endsWith(".go") || ["go.mod", "go.sum"].includes(basename(path)))));
}
function summary(state) {
  const rows = state.cases.map(row => ({ name: row.name, modules: row.modules, status: row.status, attempts: row.attempts.length, completed_invocations: row.attempts.at(-1)?.runs.filter(run => run.status === "passed").length ?? 0, expected_invocations: 24, latest: row.attempts.at(-1)?.summary ?? null, baseline_cold_ms: Object.fromEntries(variants.map(variant => [variant, state.baseline.rows[row.name][variant].median_ms])) }));
  const all = rows.length === 51 && rows.every(row => row.status === "passed");
  const totals = all ? Object.fromEntries(["cold", "warm"].map(cache => [cache, Object.fromEntries(variants.map(variant => [variant, rows.reduce((sum, row) => sum + row.latest[variant][cache].median_ms, 0)]))])) : null;
  return { purpose: state.purpose, completed: rows.filter(row => row.status === "passed").length, selected: rows.length, total_official_corpora: 51, rows, totals_ms: totals, baseline_totals_ms: state.baseline.totals_ms, interpretation: "Cold values may be compared with the recorded 25 October protocol after checking machine and input hashes. Warm values have no historical baseline. A pilot is not a performance result." };
}

if (operation === "prepare") {
  assert(options.archive && options.altbak, "prepare requires --archive and --altbak");
  assert(!existsSync(directory) || readdirSync(directory).length === 0, "Output directory must be empty");
  const archive = resolve(options.archive), compiler = resolve(options.compiler ?? root);
  assert(!directory.startsWith(archive + "/") && directory !== archive, "Campaign output must be outside the frozen archive");
  const reference = baseline(resolve(options.altbak));
  const definitions = select(readdirSync(join(archive, "cases")).filter(name => existsSync(join(archive, "cases", name, "definition.json"))).sort().map(name => ({ name, definition: join(archive, "cases", name, "definition.json") })));
  const selected = definitions.map(({ name, definition }) => {
    const data = readJson(definition);
    assert.equal(data.name, name); assert(reference.rows[name], `No official baseline for ${name}`);
    verify(data.input, data.input_manifest); verify(dirname(data.input), data.sibling_inputs ?? []);
    const tast = data.input_manifest.filter(file => file.path.endsWith("/corefn.json"));
    let types = 0;
    for (const file of tast) {
      const module = readJson(join(data.input, file.path));
      for (const key of ["typeTable", "dataDecls", "classDecls"]) assert(Array.isArray(module[key]), `${name}/${file.path}: missing TAST ${key}`);
      assert(!isAbsolute(module.modulePath), `Unrelocatable TAST module path: ${module.modulePath}`);
      const source = resolve(data.input, module.modulePath);
      assert(source.startsWith(dirname(data.input) + "/"), `TAST source escapes its frozen family: ${module.modulePath}`);
      assert(existsSync(source), `Frozen TAST source is missing: ${module.modulePath}`);
      types += module.typeTable.length;
    }
    return { name, definition, definition_sha256: fileHash(definition), input_sha256: hash(JSON.stringify(data.input_manifest)), sibling_sha256: hash(JSON.stringify(data.sibling_inputs ?? [])), tast_sha256: hash(JSON.stringify(tast)), modules: tast.length, types, input_bytes: data.input_manifest.reduce((sum, file) => sum + file.bytes, 0), frontend: { sha256: data.frontend?.sha256 ?? data.frontend?.compiler_sha256, version: data.frontend?.version ?? null }, status: "pending", attempts: [] };
  });
  mkdirSync(directory, { recursive: true });
  const minimum = Number(options["min-free-gib"] ?? 2);
  assert(Number.isFinite(minimum) && minimum > 0, "--min-free-gib must be positive"); space(minimum);
  const packageJson = readJson(join(compiler, "package.json"));
  const distribution = [...new Set(["package.json", "LICENSE", "bin/gopurs-native", ...packageJson.files].flatMap(path => statSync(join(compiler, path)).isDirectory() ? files(join(compiler, path)).map(child => path + "/" + child) : [path]))];
  const compilerFiles = manifest(compiler, distribution);
  copyManifest(compiler, join(directory, "compiler"), compilerFiles);
  for (const file of compilerFiles) chmodSync(join(directory, "compiler", file.path), file.path.startsWith("bin/") ? 0o555 : 0o444);
  const sources = ["src", "runtime"].flatMap(name => manifest(join(compiler, name)).map(file => ({ ...file, path: `${name}/${file.path}` })));
  sources.push(...manifest(compiler, ["package-lock.json", "spago.lock", "bootstrap-sources.json", "flake.lock"]));
  writeJson(join(directory, "source-manifest.json"), sources);
  const moduleCache = join(archive, "go-mod-cache");
  cpSync(moduleCache, join(directory, "go-mod-cache"), { recursive: true, mode: constants.COPYFILE_FICLONE });
  const moduleFiles = manifest(moduleCache);
  writeJson(join(directory, "module-cache-manifest.json"), moduleFiles);
  for (const name of ["tmp", "go-cache", "cases"]) mkdirSync(join(directory, name));
  let git = {};
  try { git = { commit: capture("git", ["rev-parse", "HEAD"], compiler), status: capture("git", ["status", "--short"], compiler) }; } catch {}
  const goPath = process.env.PATH.split(delimiter).map(path => join(path, "go")).find(path => existsSync(path) && statSync(path).isFile());
  assert(goPath, "Go executable is missing from PATH");
  const state = { schema: 1, purpose: options.pilot ? "pilot-functional-validation" : "release-backend-benchmark", created: new Date().toISOString(), archive, compiler_origin: compiler, version: packageJson.version,
    machine: { cpu: cpus()[0]?.model, logical_cpus: cpus().length, memory: totalmem(), platform: platform(), os_release: release() },
    tools: { node: process.version, node_sha256: fileHash(process.execPath), go: capture(goPath, ["version"]), go_path: goPath, go_sha256: fileHash(goPath) }, source: { ...git, manifest_sha256: hash(JSON.stringify(sources)), runtime_sha256: hash(JSON.stringify(sources.filter(file => file.path.startsWith("runtime/")))) },
    runner: { sha256: fileHash(script), process_helper_sha256: fileHash(join(root, "tools/command-runner.mjs")) }, compiler_files: compilerFiles,
    module_cache: { original: moduleCache, manifest_sha256: hash(JSON.stringify(moduleFiles)), files: moduleFiles.length }, baseline: reference,
    protocol: { metric: "backend total", warmups: 1, samples: 3, orders, cache_order: ["cold", "warm"], min_free_gib: minimum, go_cache_budget_mib: 1024, cold: "fresh output and backend caches per pair; warm filesystem/Go export cache", warm: "fresh process immediately after the same variant's cold generation; exact input/output cache hit required", excludes: ["frontend", "compiler bootstrap", "final go build", "application execution", "process startup/exit"] }, cases: selected };
  writeJson(statePath, state);
  console.log(`Prepared ${selected.length} frozen corpora; no compiler was launched. ${statePath}`);
} else if (operation === "status") {
  console.log(JSON.stringify(summary(readJson(statePath)), null, 2));
} else {
  const state = readJson(statePath);
  assert.equal(state.schema, 1);
  assert.equal(fileHash(script), state.runner.sha256, "Benchmark runner changed; prepare a new campaign");
  assert.equal(fileHash(join(root, "tools/command-runner.mjs")), state.runner.process_helper_sha256, "Process helper changed");
  assert.equal(fileHash(process.execPath), state.tools.node_sha256, "Node executable changed");
  assert.equal(fileHash(state.tools.go_path), state.tools.go_sha256, "Go executable changed");
  assert.equal(capture(state.tools.go_path, ["version"]), state.tools.go, "Go toolchain changed");
  const moduleFiles = readJson(join(directory, "module-cache-manifest.json"));
  assert.equal(hash(JSON.stringify(moduleFiles)), state.module_cache.manifest_sha256, "Go module cache manifest changed");
  verify(join(directory, "go-mod-cache"), moduleFiles);
  for (const file of state.baseline.files) assert.equal(fileHash(file.path), file.sha256, `Baseline changed: ${file.path}`);
  const lock = join(directory, "run.lock");
  if (existsSync(lock)) {
    const prior = readJson(lock);
    try { process.kill(prior.pid, 0); throw new Error(`Campaign process ${prior.pid} is still running`); }
    catch (error) { if (error.code !== "ESRCH") throw error; rmSync(lock); }
  }
  const lockFd = openSync(lock, "wx"); writeFileSync(lockFd, JSON.stringify({ pid: process.pid, started: new Date().toISOString() })); closeSync(lockFd);
  const runner = new CommandRunner();
  const save = () => writeJson(statePath, state);
  try {
    for (const row of select(state.cases)) {
      if (row.status === "passed") { console.log(`${row.name}: completed case retained`); continue; }
      assert(row.status === "pending" || options["retry-failed"], `${row.name} is ${row.status}; use --retry-failed to preserve and exclude its old attempt, then restart the complete case`);
      runner.checkInterrupted();
      const definition = readJson(row.definition);
      assert.equal(fileHash(row.definition), row.definition_sha256, "Frozen corpus definition changed");
      verify(definition.input, definition.input_manifest); verify(dirname(definition.input), definition.sibling_inputs ?? []);
      verify(join(directory, "compiler"), state.compiler_files);
      const active = join(directory, "active"); rmSync(active, { recursive: true, force: true }); space(state.protocol.min_free_gib); mkdirSync(active);
      const input = join(active, basename(definition.input));
      copyManifest(definition.input, input, definition.input_manifest.filter(file => !file.path.startsWith("output/")));
      copyManifest(dirname(definition.input), active, definition.sibling_inputs ?? []);
      // Some source snapshots use a local spago.yaml → spago.go.yaml link.
      // Recreate only links whose targets remain inside this private family.
      for (const entry of readdirSync(definition.input, { withFileTypes: true }).filter(entry => entry.isSymbolicLink())) {
        const target = readlinkSync(join(definition.input, entry.name));
        assert(!isAbsolute(target) && resolve(input, target).startsWith(active + "/"), `External snapshot symlink: ${entry.name}`);
        symlinkSync(target, join(input, entry.name));
      }
      const attempt = { number: row.attempts.length + 1, status: "running", started: new Date().toISOString(), runs: [], oracles: {} };
      const evidence = join(directory, "cases", row.name, `attempt-${attempt.number}`); mkdirSync(evidence, { recursive: true });
      row.attempts.push(attempt); row.status = "running"; save();
      try {
        for (let round = 0; round < orders.length; round++) for (const variant of orders[round]) {
          space(state.protocol.min_free_gib); verify(join(directory, "compiler"), state.compiler_files);
          const output = join(input, "output"); rmSync(output, { recursive: true, force: true });
          copyManifest(definition.input, input, definition.input_manifest.filter(file => file.path.startsWith("output/")));
          for (const name of [".cache", ".purmeta", ".gopurs-cache", ".gopurs-daemon"]) rmSync(join(input, name), { recursive: true, force: true });
          assert(!existsSync(join(output, ".gopurs-cache")));
          for (const cache of ["cold", "warm"]) {
            runner.checkInterrupted(); space(state.protocol.min_free_gib);
            const label = `${round ? `sample-${round}` : "warmup"}-${variant}-${cache}`;
            const env = environment(state, variant), command = join(directory, "compiler/bin/gopurs");
            const args = [...definition.invocation, ...(variant === "dev" ? ["--dev"] : [])];
            const record = { variant, cache, round, measured: round > 0, status: "running", started: new Date().toISOString(), load_before: loadavg(), command: [command, ...args], environment: Object.fromEntries(Object.entries(env).filter(([name]) => /^(GOPURS|GOGC|GOMEMLIMIT|GOMAXPROCS|GOWORK|GOCACHE|GOMODCACHE|GOPROXY|GOTOOLCHAIN)/.test(name))), log: join(evidence, `${label}.log`) };
            attempt.runs.push(record); save(); console.log(`${row.name} ${label}: started`);
            const started = performance.now(); let timedOut = false; let forceTimer;
            const timer = setTimeout(() => {
              timedOut = true;
              if (runner.active?.pid) {
                const pid = runner.active.pid;
                try { process.kill(-pid, "SIGTERM"); } catch (error) { if (error.code !== "ESRCH") throw error; }
                forceTimer = setTimeout(() => { try { process.kill(-pid, "SIGKILL"); } catch (error) { if (error.code !== "ESRCH") throw error; } }, 5000);
              }
            }, definition.timeout_ms ?? 600000);
            let result;
            try { result = await runner.run(command, args, { cwd: input, env, log: record.log }); }
            finally { clearTimeout(timer); clearTimeout(forceTimer); record.wall_ms = performance.now() - started; record.finished = new Date().toISOString(); record.load_after = loadavg(); record.timed_out = timedOut; save(); }
            Object.assign(record, { exit: result.code, signal: result.signal }); save(); runner.checkInterrupted();
            assert(!timedOut, `${label}: timed out`); assert.equal(result.code, 0, `${label}: compiler failed; see ${record.log}`);
            const raw = readFileSync(record.log, "utf8");
            const totals = [...raw.matchAll(/^\[gopurs\] backend total: (\d+) ms$/gm)];
            assert.equal(totals.length, 1, `${label}: missing/duplicate backend timer`);
            assert(!raw.includes("[gopurs] full fallback:"), `${label}: full fallback is not a comparable sample`);
            record.cache_hit = raw.includes("[gopurs] cache hit: inputs and outputs verified; 0 modules optimized/emitted");
            assert.equal(record.cache_hit, cache === "warm", `${label}: unexpected backend cache state`);
            if (cache === "cold") assert(raw.includes("[gopurs] cache miss:"), `${label}: missing cold-cache evidence`);
            record.backend_ms = Number(totals[0][1]);
            record.phases_ms = Object.fromEntries([...raw.matchAll(/^\[gopurs\] (.+): (\d+) ms$/gm)].map(([, name, ms]) => [name, Number(ms)]));
            const actual = generated(output); assert(actual.length > 0);
            record.generated_sha256 = hash(JSON.stringify(actual)); record.generated_files = actual.length;
            const profile = variant === "dev" ? "dev" : "release";
            const oracle = attempt.oracles[profile];
            if (!oracle) {
              assert(variant === (profile === "release" ? "js" : "dev"), "Unexpected oracle variant");
              attempt.oracles[profile] = { sha256: record.generated_sha256, file: join(evidence, `${profile}-oracle.json`), established_by: label };
              writeJson(attempt.oracles[profile].file, actual);
            } else if (record.generated_sha256 !== oracle.sha256) {
              writeJson(join(evidence, `${label}-actual.json`), actual);
              throw new Error(`${label}: generated files differ from the ${profile} oracle`);
            }
            verify(input, definition.input_manifest); verify(active, definition.sibling_inputs ?? []);
            record.checked_input = true; record.status = "passed"; writeJson(join(evidence, `${label}.json`), record); save();
            console.log(`${row.name} ${label}: ${record.backend_ms} ms; ${record.generated_files} file hashes verified`);
          }
        }
        attempt.summary = Object.fromEntries(variants.map(variant => [variant, Object.fromEntries(["cold", "warm"].map(cache => {
          const samples = attempt.runs.filter(run => run.variant === variant && run.cache === cache && run.measured && run.status === "passed").map(run => run.backend_ms);
          assert.equal(samples.length, 3);
          return [cache, { samples_ms: samples, median_ms: samples.toSorted((a, b) => a - b)[1] }];
        }))]));
        assert.equal(attempt.runs.length, 24); attempt.status = "passed"; row.status = "passed";
        rmSync(active, { recursive: true, force: true });
        rmSync(join(directory, "tmp"), { recursive: true }); mkdirSync(join(directory, "tmp"));
        const cacheBytes = files(join(directory, "go-cache")).reduce((sum, file) => sum + statSync(join(directory, "go-cache", file)).size, 0);
        attempt.go_cache_bytes_after = cacheBytes;
        if (cacheBytes > state.protocol.go_cache_budget_mib * 2 ** 20) { rmSync(join(directory, "go-cache"), { recursive: true }); mkdirSync(join(directory, "go-cache")); attempt.go_cache_evicted_between_cases = true; }
      } catch (error) {
        attempt.status = runner.signal ? "interrupted" : "failed"; row.status = attempt.status; attempt.error = error.stack;
        delete attempt.summary;
        const current = attempt.runs.at(-1); if (current?.status === "running") { current.status = attempt.status; current.error = error.message; }
        console.error(`${row.name}: ${error.message}`); process.exitCode = runner.signal === "SIGINT" ? 130 : runner.signal ? 143 : 1;
      } finally {
        attempt.finished = new Date().toISOString();
        try {
          verify(join(directory, "compiler"), state.compiler_files);
          verify(join(directory, "go-mod-cache"), moduleFiles);
          verify(definition.input, definition.input_manifest); verify(dirname(definition.input), definition.sibling_inputs ?? []);
        } catch (error) {
          attempt.status = "failed"; row.status = "failed"; attempt.error = error.stack; delete attempt.summary; process.exitCode = 1;
          console.error(`${row.name}: final input verification failed: ${error.message}`);
        }
        save(); writeJson(join(directory, "summary.json"), summary(state));
      }
      if (row.status !== "passed" && (!options["keep-going"] || runner.signal)) break;
    }
  } finally { runner.dispose(); rmSync(lock, { force: true }); }
  console.log(JSON.stringify({ completed: state.cases.filter(row => row.status === "passed").length, selected: state.cases.length, purpose: state.purpose }));
}
