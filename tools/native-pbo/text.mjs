#!/usr/bin/env node
// Ready-to-run differential runner for the native CoreFn text decoder.
//
// It reuses the native-PBO mechanics (tools/native-pbo/shared.mjs) and the
// frozen boundary matrix (purust tools/test-native-tast-module.mjs), injects
// the candidate Text.rs through the shared threaded transform, builds one
// fixture crate against the retained Rust workspace and runs:
//
//   native Text.rs  vs  jsonParser >=> decodeModulePS >=> lmap print
//
// on the frozen corpus, on deterministic contract mutations and on the
// boundary matrix. Success values are compared through a stable canonical
// structural dump (Class and ClassShared tolerant); failures through the
// printed Left string byte for byte. Nothing is compared by pointer identity
// and the candidate is never compared to itself.
//
// Standalone final gate for the native CoreFn text decoder. The candidate
// FFI defaults to the live gopurs PBO Text.rs (shared pboSource helper) and
// can be overridden with --text. The reference is a full PureScript oracle:
// the runner copies the generated Json/Usage crates into Oracle packages and
// rewrites their FFI entry points to call their PS bodies.
//
// Usage:
//   node tools/native-pbo/text.mjs --workspace RUST --corpus DIR_OR_JSON \
//     [--expect-cases N] [--expect-corpus N] [--log DIR] [--target-dir DIR] \
//     [--text FILE] [--mutations contracts|none] [--modules all|N] \
//     [--boundaries yes|no] [--plan]
//
// RUST must be a generated rust/ workspace whose purust-argonaut-core includes
// PurustJsonCursor::materialize (the gate verifies this before building).
import assert from "node:assert/strict";
import { appendFileSync, cpSync, existsSync, mkdirSync, mkdtempSync, readFileSync, readdirSync, statSync, writeFileSync } from "node:fs";
import { createHash } from "node:crypto";
import { tmpdir } from "node:os";
import { basename, dirname, join, resolve } from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";

const here = dirname(fileURLToPath(import.meta.url));

function usage() {
  return [
    "Usage: node tests/run-text-differential.mjs --workspace RUST --corpus DIR_OR_JSON [options]",
    "",
    "  --workspace DIR     retained rust/ workspace of the shared-classes candidate",
    "  --corpus PATH       <Module>/corefn.json directory, campaign results.json, or files manifest",
    "  --corpus-name NAME  label recorded in the plan/log",
    "  --text FILE         candidate Text.rs (default live PBO Text.rs via pboSource)",
    "  --log DIR           log directory (default <workspace>/../text-contracts)",
    "  --target-dir DIR    shared cargo target dir (default <workspace>/../target)",
    "  --mutations MODE    contracts (default) or none",
    "  --modules all|N     limit corpus modules (deterministic sorted order)",
    "  --boundaries yes|no include the frozen boundary matrix (default yes)",
    "  --plan              write cases and print the plan without building",
    "  --expect-cases N    fail unless the plan has exactly N cases",
    "  --expect-corpus N   fail unless the corpus has N cases and none declined",
  ].join("\n");
}

function parse(argv) {
  const options = {
    workspace: null, corpus: null, corpusName: null, text: null,
    log: null, targetDir: null, mutations: "contracts", modules: "all", boundaries: true,
    plan: false, help: false, expectCases: null, expectCorpus: null,
  };
  for (let index = 0; index < argv.length; index += 1) {
    const argument = argv[index];
    const value = () => { const next = argv[index + 1]; assert.ok(next, `${argument} needs a value`); index += 1; return next; };
    if (argument === "--help" || argument === "-h") options.help = true;
    else if (argument === "--workspace") options.workspace = resolve(value());
    else if (argument === "--corpus") options.corpus = resolve(value());
    else if (argument === "--corpus-name") options.corpusName = value();
    else if (argument === "--text") options.text = resolve(value());
    else if (argument === "--log") options.log = resolve(value());
    else if (argument === "--target-dir") options.targetDir = resolve(value());
    else if (argument === "--mutations") options.mutations = value();
    else if (argument === "--modules") options.modules = value();
    else if (argument === "--boundaries") options.boundaries = value() !== "no";
    else if (argument === "--plan") options.plan = true;
    else if (argument === "--expect-cases") options.expectCases = Number.parseInt(value(), 10);
    else if (argument === "--expect-corpus") options.expectCorpus = Number.parseInt(value(), 10);
    else throw new Error(`Unknown option: ${argument}\n\n${usage()}`);
  }
  return options;
}

const options = parse(process.argv.slice(2));
if (options.help) { console.log(usage()); process.exit(0); }
assert.ok(options.workspace, `--workspace is required\n\n${usage()}`);
assert.ok(["contracts", "none"].includes(options.mutations), "unknown --mutations mode");

// Shared native-PBO plumbing: workspace creation, crate manifests, logging,
// cargo invocation and the threaded FFI transform all come from the campaign
// harness rather than a copy.
const sharedPath = join(here, "shared.mjs");
assert.ok(existsSync(sharedPath), `Missing native-pbo shared.mjs: ${sharedPath}`);
const shared = await import(pathToFileURL(sharedPath).href);
if (!options.corpus) options.corpus = shared.defaultCorpus(options.workspace);
assert.ok(options.corpus, `--corpus is required (or a sibling output directory)`);
const fixturePath = join(here, "fixtures/test-native-text.rs");

// ---------------------------------------------------------------------------
// Candidate and workspace checks

const ffiSource = options.text
  ? readFileSync(resolve(options.text), "utf8")
  : shared.pboSource("PureScript/Backend/Optimizer/CoreFn/Json/Text.rs");
assert.ok(ffiSource.includes("PureScript_Backend_Optimizer_CoreFn_Json_Text_parseModuleTextImpl"),
  `Not a parseModuleTextImpl candidate: ${options.text ?? "live PBO Text.rs"}`);
assert.ok(existsSync(options.workspace), `Missing workspace: ${options.workspace}`);
const argonaut = join(options.workspace, "Purs_Data_Argonaut_Core/src/lib.rs");
assert.ok(existsSync(argonaut), `Workspace lacks Purs_Data_Argonaut_Core: ${options.workspace}`);

// ---------------------------------------------------------------------------
// Corpus resolution

function resolveCorpus(path, modulesLimit) {
  const entries = [];
  const stat = statSync(path, { throwIfNoEntry: false });
  assert.ok(stat, `Missing corpus: ${path}`);
  if (stat.isDirectory()) {
    for (const name of readdirSync(path).sort()) {
      const file = join(path, name, "corefn.json");
      if (existsSync(file)) entries.push({ name, path: file });
    }
  } else {
    const parsed = JSON.parse(readFileSync(path, "utf8"));
    const files = Array.isArray(parsed) ? parsed
      : parsed?.tast?.files ?? parsed?.files ?? [];
    assert.ok(Array.isArray(files) && files.length > 0, `No module list in ${path}`);
    for (const entry of files) {
      const modulePath = entry.path ?? entry.file;
      assert.ok(modulePath, `Corpus entry without path: ${JSON.stringify(entry)}`);
      entries.push({ name: entry.module ?? basename(dirname(modulePath)), path: modulePath });
    }
  }
  assert.ok(entries.length > 0, `No corefn.json modules in ${path}`);
  entries.sort((left, right) => left.name.localeCompare(right.name));
  if (modulesLimit !== "all") {
    const limit = Number.parseInt(modulesLimit, 10);
    assert.ok(Number.isInteger(limit) && limit > 0, `Invalid --modules ${modulesLimit}`);
    entries.length = Math.min(entries.length, limit);
  }
  return entries;
}

const modules = options.corpus ? resolveCorpus(options.corpus, options.modules) : [];
const corpusName = options.corpusName ?? (options.corpus ? basename(options.corpus) : "none");


// ---------------------------------------------------------------------------
// Full PureScript oracle. The runner copies the generated Json and Usage
// crates into the fixture workspace under Oracle package names and rewrites
// their FFI entry points to call their PureScript bodies, so the reference
// composition shares no native decoder with the candidate. The candidate
// still links the original crates (cold decoders and validate).

const ORACLE_PACKAGES = {
  json: { source: "Purs_PureScript_Backend_Optimizer_CoreFn_Json", oracle: "Purs_PureScript_Backend_Optimizer_CoreFn_JsonOracle" },
  usage: { source: "Purs_PureScript_Backend_Optimizer_CoreFn_Usage", oracle: "Purs_PureScript_Backend_Optimizer_CoreFn_UsageOracle" },
};

const NATIVE_CUTS = {
  "Purs_PureScript_Backend_Optimizer_CoreFn_JsonOracle": {
    start: "mod purust_type_table {",
    end: "pub fn PureScript_Backend_Optimizer_CoreFn_Json_decodeTypeTableImpl(",
  },
  "Purs_PureScript_Backend_Optimizer_CoreFn_UsageOracle": {
    start: "use std::collections::HashSet;",
    end: "pub fn PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModuleImpl(",
  },
};

const PS_BODIES = {
  "PureScript_Backend_Optimizer_CoreFn_Json_decodeAnnWithUsageImpl": "{ fallback(module, table, path, input) }",
  "PureScript_Backend_Optimizer_CoreFn_Json_decodeArrayImpl": "{ fallback(decoder, input) }",
  "PureScript_Backend_Optimizer_CoreFn_Json_decodeModuleImpl": "{ fallback(input) }",
  "PureScript_Backend_Optimizer_CoreFn_Json_decodeTypeTableImpl": "{ Purs_PureScript_Backend_Optimizer_CoreFn_TypeTable::PureScript_Backend_Optimizer_CoreFn_TypeTable_decodeTypeTablePS(input) }",
  "PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModuleImpl": "{ fallback(module) }",
};

const sha256 = text => createHash("sha256").update(text).digest("hex");

function findFunctionBody(source, name) {
  const anchor = `pub fn ${name}(`;
  assert.equal(source.split(anchor).length - 1, 1, `oracle: expected one ${name}`);
  let at = source.indexOf(anchor) + anchor.length;
  for (let depth = 1; at < source.length; at += 1) {
    if (source[at] === "(") depth += 1;
    else if (source[at] === ")" && (depth -= 1) === 0) { at = source.indexOf("{", at); break; }
  }
  assert.ok(at > 0, `oracle: no body for ${name}`);
  const start = at;
  for (let depth = 0; at < source.length; at += 1) {
    if (source.startsWith("//", at)) { at = source.indexOf("\n", at); continue; }
    if (source.startsWith("/*", at)) {
      let nest = 1; at += 2;
      while (at < source.length && nest) {
        if (source.startsWith("/*", at)) { nest += 1; at += 2; }
        else if (source.startsWith("*/", at)) { nest -= 1; at += 2; }
        else at += 1;
      }
      at -= 1; continue;
    }
    const c = source[at];
    if (c === '"') { for (at += 1; at < source.length && source[at] !== '"'; at += 1) if (source[at] === "\\") at += 1; continue; }
    if (c === "'") {
      if (source[at + 1] === "\\") { at = source.indexOf("'", at + 2); continue; }
      const close = source.indexOf("'", at + 1);
      if (close > at && close <= at + 4) at = close;
      continue;
    }
    if (c === "{") depth += 1;
    else if (c === "}" && (depth -= 1) === 0) return { start, end: at };
  }
  throw new Error(`oracle: unbalanced body for ${name}`);
}

function rewriteFunction(source, name, body) {
  const { start, end } = findFunctionBody(source, name);
  return source.slice(0, start) + body + source.slice(end + 1);
}

function rewriteCargo(source, definition, overrides, workspace) {
  const rewrites = [];
  const lines = source.split("\n").map(line => {
    if (line === `name = "${definition.source}"`) { rewrites.push(`package ${definition.source} -> ${definition.oracle}`); return `name = "${definition.oracle}"`; }
    const match = line.match(/^([A-Za-z0-9_]+) = \{ path = "([^"]+)"(,\s*(.*))?\s*\}$/);
    if (!match) return line;
    const [, dependency, relativePath, , tail] = match;
    const override = overrides[dependency];
    const absolute = override ? override.path : join(workspace, basename(relativePath));
    rewrites.push(`${dependency}: ${relativePath} -> ${absolute}`);
    return override
      ? `${dependency} = { package = "${override.package}", path = ${JSON.stringify(absolute)}${tail ? `, ${tail}` : ""} }`
      : `${dependency} = { path = ${JSON.stringify(absolute)}${tail ? `, ${tail}` : ""} }`;
  });
  return { source: lines.join("\n"), rewrites };
}

function prepareOracles(fixture, workspace, logDir, ctx) {
  const { json, usage } = ORACLE_PACKAGES;
  const usageOverride = { [usage.source]: { package: usage.oracle, path: join(fixture, "oracle", usage.oracle) } };
  const definitions = [
    { ...json, overrides: usageOverride, functions: Object.keys(PS_BODIES).filter(name => name.includes("CoreFn_Json")) },
    { ...usage, overrides: {}, functions: Object.keys(PS_BODIES).filter(name => name.includes("CoreFn_Usage")) },
  ];
  const report = [];
  for (const [index, definition] of definitions.entries()) {
    const destination = join(fixture, "oracle", definition.oracle);
    mkdirSync(destination, { recursive: true });
    cpSync(join(workspace, definition.source), destination, { recursive: true });
    const cargoPath = join(destination, "Cargo.toml");
    const cargoBefore = readFileSync(cargoPath, "utf8");
    const cargo = rewriteCargo(cargoBefore, definition, definition.overrides, workspace);
    writeFileSync(cargoPath, cargo.source);
    const libPath = join(destination, "src/lib.rs");
    const before = readFileSync(libPath, "utf8");
    let after = before;
    for (const name of definition.functions) after = rewriteFunction(after, name, PS_BODIES[name]);
    // Drop the dead native helper modules that preceded the entry points; the
    // rewritten entries call only their PS bodies.
    const cut = NATIVE_CUTS[definition.oracle];
    assert.equal(after.split(cut.start).length - 1, 1, `${definition.oracle}: native start anchor is not unique`);
    const from = after.indexOf(cut.start), to = after.indexOf(cut.end);
    assert.ok(from > 0 && to > from, `${definition.oracle}: native suffix anchors out of order`);
    after = after.slice(0, from) + after.slice(to);
    writeFileSync(libPath, after);
    for (const name of definition.functions) {
      const { start, end } = findFunctionBody(after, name);
      assert.equal(after.slice(start, end + 1), PS_BODIES[name], `${definition.oracle}: ${name} body is not PS-only`);
    }
    if (index === 0) {
      for (const call of ["purust_type_table::decode(", "purust_ann::decode(", "purust_module::decode("])
        assert.equal(after.split(call).length - 1, 0, `${definition.oracle}: native call ${call} remains`);
      for (const [wrapper, fallback] of [
        ["PureScript_Backend_Optimizer_CoreFn_Json_decodeModule", "PureScript_Backend_Optimizer_CoreFn_Json_decodeModulePS"],
        ["PureScript_Backend_Optimizer_CoreFn_Json_decodeAnnWithUsage", "PureScript_Backend_Optimizer_CoreFn_Json_decodeAnnWithUsagePS"],
        ["PureScript_Backend_Optimizer_CoreFn_Json_decodeArray", "PureScript_Backend_Optimizer_CoreFn_Json_decodeArrayPS"],
      ]) {
        const { start, end } = findFunctionBody(after, wrapper);
        const body = after.slice(start, end + 1);
        assert.ok(body.includes(fallback), `${definition.oracle}: ${wrapper} does not pass ${fallback}`);
        assert.ok(!body.includes(`Static(${wrapper})`), `${definition.oracle}: ${wrapper} would loop through itself`);
      }
    } else {
      assert.equal(after.split("PurustUsageChecker {").length - 1, 0, `${definition.oracle}: native validator body remains`);
    }
    const retained = join(logDir, "oracle", definition.oracle);
    mkdirSync(retained, { recursive: true });
    writeFileSync(join(retained, "Cargo.toml"), cargo.source);
    writeFileSync(join(retained, "lib.rs"), after);
    report.push({
      source: definition.source,
      oracle: definition.oracle,
      original_lib_sha256: sha256(before),
      rewritten_lib_sha256: sha256(after),
      rewritten_cargo_sha256: sha256(cargo.source),
      function_overrides: definition.functions.map(name => ({ function: name, body: PS_BODIES[name] })),
      cargo_path_rewrites: cargo.rewrites,
      retained,
    });
    ctx.note(`PS-only oracle ${definition.oracle}: ${definition.functions.length} entry points rewritten, lib sha ${sha256(after).slice(0, 16)}`);
  }
  writeFileSync(join(logDir, "oracle-report.json"), JSON.stringify(report, null, 1) + "\n");
  ctx.note(`oracle report ${join(logDir, "oracle-report.json")}`);
  return report;
}

// ---------------------------------------------------------------------------
// Deterministic contract mutations

function firstStringRange(text) {
  for (let index = 0; index < text.length; index += 1) {
    if (text[index] !== '"') continue;
    let end = index + 1;
    while (end < text.length) {
      if (text[end] === "\\") { end += 2; continue; }
      if (text[end] === '"') break;
      end += 1;
    }
    if (end - index - 1 >= 2) return [index + 1, end];
    index = end;
  }
  return null;
}

function insertIntoFirstString(text, escape) {
  const range = firstStringRange(text);
  if (!range) return null;
  const [start, end] = range;
  let at = start + Math.floor((end - start) / 2);
  while (at > start && text[at - 1] === "\\") at -= 1;
  return text.slice(0, at) + escape + text.slice(at);
}

const MUTATIONS = [
  ["truncate-60", text => text.slice(0, Math.max(1, Math.floor(text.length * 3 / 5)))],
  ["truncate-90", text => text.slice(0, Math.max(1, Math.floor(text.length * 9 / 10)))],
  ["delete-byte", text => {
    const at = Math.min(text.length - 1, Math.max(1, Math.floor(text.length * 2 / 3)));
    return text.slice(0, at) + text.slice(at + 1);
  }],
  ["trailing-garbage", text => `${text}garbage`],
  ["leading-bom", text => `\ufeff${text}`],
  ["bad-escape", text => insertIntoFirstString(text, "\\x")],
  ["lone-surrogate", text => insertIntoFirstString(text, "\\ud800")],
  ["surrogate-pair", text => insertIntoFirstString(text, "\\ud83d\\ude00")],
  ["duplicate-modulePath", text => {
    const match = text.match(/"modulePath"\s*:\s*"(?:[^"\\]|\\.)*"/);
    return match ? `${text.slice(0, match.index + match[0].length)},"modulePath":"Mutated/Path.purs"${text.slice(match.index + match[0].length)}` : null;
  }],
  ["duplicate-type-tag", text => {
    const match = text.match(/("type"\s*:\s*")([A-Za-z]+)(")/);
    return match ? `${text.slice(0, match.index + match[0].length)},"type":"Mutated"${text.slice(match.index + match[0].length)}` : null;
  }],
  ["unknown-modulePath", text => text.replace('"modulePath":', '"modulePathRenamed":')],
  ["rename-decls", text => text.replace('"decls"', '"declsRenamed"')],
  ["rename-annotation", text => text.replace('"annotation"', '"annotationRenamed"')],
  ["null-modulePath", text => {
    const match = text.match(/"modulePath"\s*:\s*"(?:[^"\\]|\\.)*"/);
    return match ? `${text.slice(0, match.index)}"modulePath":null${text.slice(match.index + match[0].length)}` : null;
  }],
  ["cycle-typetable", text => text.includes('"typeTable":[')
    ? text.replace('"typeTable":[', '"typeTable":[{"type":"Record","row":0},') : null],
  ["deep-comments-130", text => text.includes('"comments":[]')
    ? text.replace('"comments":[]', `"comments":${"[".repeat(130)}${"]".repeat(130)}`) : null],
];

function contractMutations(text) {
  const results = [];
  for (const [kind, apply] of MUTATIONS) {
    const mutated = apply(text);
    if (typeof mutated === "string") results.push([kind, mutated]);
  }
  return results;
}

// ---------------------------------------------------------------------------
// Boundaries (frozen matrix from the Purust harness)

async function boundaryCases() {
  if (!options.boundaries) return [];
  const module = join(shared.purustDir, "tools/test-native-tast-module.mjs");
  assert.ok(existsSync(module), `Missing boundary matrix: ${module}`);
  const { buildBoundaries } = await import(pathToFileURL(module).href);
  return buildBoundaries().map(([mode, name, value]) => [`boundary-${mode}`, name, JSON.stringify(value)]);
}

// ---------------------------------------------------------------------------
// Case materialization

const cases = [];
for (const module of modules) cases.push(["corpus", module.name, readFileSync(module.path, "utf8")]);
if (options.mutations === "contracts") {
  for (const module of modules) {
    const text = readFileSync(module.path, "utf8");
    for (const [kind, mutated] of contractMutations(text)) cases.push([`mutation-${kind}`, module.name, mutated]);
  }
}
for (const entry of await boundaryCases()) cases.push(entry);

const logDir = options.log ? resolve(options.log) : join(dirname(options.workspace), "text-contracts");
const planDirectory = mkdtempSync(join(tmpdir(), "gopurs-text-diff-plan-"));
mkdirSync(join(planDirectory, "cases"), { recursive: true });
const manifest = [];
cases.forEach(([kind, name, text], index) => {
  const file = `case-${String(index).padStart(5, "0")}.txt`;
  writeFileSync(join(planDirectory, "cases", file), text, "utf8");
  manifest.push(`${kind}\t${name}\t${file}`);
});
writeFileSync(join(planDirectory, "cases", "manifest.tsv"), `${manifest.join("\n")}\n`);

const kindCounts = {};
for (const [kind] of cases) kindCounts[kind] = (kindCounts[kind] ?? 0) + 1;
if (options.expectCases !== null) {
  assert.ok(Number.isInteger(options.expectCases) && options.expectCases > 0, "--expect-cases needs a positive integer");
  assert.equal(cases.length, options.expectCases, `expected ${options.expectCases} cases, planned ${cases.length}`);
}
console.log(`corpus ${corpusName}: ${modules.length} modules, ${cases.length} cases`);
for (const [kind, count] of Object.entries(kindCounts).sort()) console.log(`  ${kind}: ${count}`);
console.log(`cases: ${join(planDirectory, "cases")}`);

if (options.plan) {
  console.log("plan only: no cargo build was run");
  process.exit(0);
}

// The fixture build links the candidate against the retained workspace, so the
// cursor materialize extension must be present there.
if (!readFileSync(argonaut, "utf8").includes("pub fn materialize")) {
  throw new Error(
    `${argonaut} has no PurustJsonCursor::materialize. Rebuild the workspace with a ` +
    "purust-argonaut-core package that includes the cursor materialize extension, then rerun this gate.");
}

// ---------------------------------------------------------------------------
// Fixture build and run

const threadedRust = await shared.loadThreadedRust();
assert.ok(existsSync(fixturePath), `Missing fixture: ${fixturePath}`);
const fixture = readFileSync(fixturePath, "utf8");
assert.ok(/^\s*\/\/ NATIVE_TEXT_FFI$/m.test(fixture), "fixture lacks the NATIVE_TEXT_FFI marker");

const dependencies = [
  "purust_core", "perceus_ptr",
  "Purs_Data_Argonaut_Core", "Purs_Data_Argonaut_Decode_Error",
  "Purs_Data_Either", "Purs_Data_Maybe", "Purs_Data_Tuple", "Purs_Data_Unfoldable",
  "Purs_Foreign_Object", "Purs_Data_Map_Internal", "Purs_Data_Ord", "Purs_Data_Foldable",
  "Purs_PureScript_Backend_Optimizer_CoreFn",
  "Purs_PureScript_Backend_Optimizer_CoreFn_Json",
  "Purs_PureScript_Backend_Optimizer_CoreFn_Usage",
];

if (options.targetDir) process.env.PBO_NATIVE_TARGET_DIR = options.targetDir;
await shared.runCase("text", options.workspace, async ctx => {
  const directory = shared.createWorkspace("gopurs-text-differential");
  ctx.state.workspace = directory;
  mkdirSync(logDir, { recursive: true });
  ctx.note(`fixture workspace ${directory}`);
  shared.writeCrate(directory, "gopurs_text_differential", dependencies, options.workspace);
  prepareOracles(directory, options.workspace, logDir, ctx);
  appendFileSync(join(directory, "Cargo.toml"),
    `Purs_PureScript_Backend_Optimizer_CoreFn_JsonOracle = { path = "oracle/Purs_PureScript_Backend_Optimizer_CoreFn_JsonOracle" }\n`);
  const casesDir = join(directory, "cases");
  mkdirSync(casesDir, { recursive: true });
  for (const line of manifest) {
    const [, , file] = line.split("\t");
    writeFileSync(join(casesDir, file), readFileSync(join(planDirectory, "cases", file)));
  }
  writeFileSync(join(casesDir, "manifest.tsv"), `${manifest.join("\n")}\n`);
  writeFileSync(join(directory, "src/main.rs"),
    fixture.replace(/^\s*\/\/ NATIVE_TEXT_FFI$/m, () => threadedRust(ffiSource)));
  ctx.note(`injected candidate ${options.text ?? "live PBO Text.rs"} through threadedRust`);
  ctx.note(`${cases.length} cases (${Object.entries(kindCounts).map(([kind, count]) => `${kind}=${count}`).join(", ")})`);
  const executable = ctx.cargoBuild(directory, "gopurs_text_differential");
  const report = join(directory, "report.json");
  ctx.spawn("run", executable, [casesDir, report]);
  mkdirSync(logDir, { recursive: true });
  const saved = join(logDir, "report.json");
  writeFileSync(saved, readFileSync(report));
  const summary = JSON.parse(readFileSync(saved, "utf8"));
  assert.deepEqual(summary.failures, []);
  const groups = Object.values(summary.kinds);
  assert.equal(groups.reduce((total, group) => total + group.total, 0), cases.length);
  for (const group of groups) assert.equal(group.ok, group.total);
  assert(groups.some(group => group.class_carriers > 0), "legacy carriers were not exercised");
  assert(groups.some(group => group.shared_carriers > 0), "shared carriers were not exercised");
  const corpus = summary.kinds?.corpus;
  assert.ok(corpus, "report lacks corpus stats");
  if (options.expectCorpus !== null) {
    assert.ok(Number.isInteger(options.expectCorpus) && options.expectCorpus > 0, "--expect-corpus needs a positive integer");
    assert.equal(corpus.total, options.expectCorpus, `corpus total ${corpus.total} != ${options.expectCorpus}`);
  }
  assert.equal(corpus.declined, 0, "corpus cases must not decline the native path");
  assert.equal(corpus.ok, corpus.total, "corpus cases must all match");
  ctx.note(`corpus ${corpus.total} cases, 0 declined, ${corpus.ok} equal`);
  ctx.note(`report ${saved}`);
  ctx.note(`plan cases ${join(planDirectory, "cases")}`);
});
