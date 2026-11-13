import { execFileSync } from "node:child_process";
import { accessSync, constants, existsSync, mkdirSync, readFileSync, readdirSync, realpathSync, statSync, writeFileSync } from "node:fs";
import { dirname, join, resolve } from "node:path";
import { readJson } from "../lib/fs.mjs";

// Copy the package section verbatim, excluding its optional test stanza. This
// preserves dependency constraints without a second YAML parser dependency.
function compilerPackage(root) {
  const lines = readFileSync(join(root, "spago.yaml"), "utf8").split(/\r?\n/);
  const start = lines.findIndex((line) => /^package:\s*$/.test(line));
  if (start < 0) throw new Error("spago.yaml has no package section");
  let skipTest = false;
  const result = [];
  for (const line of lines.slice(start)) {
    if (result.length && /^\S/.test(line)) break;
    if (/^  test:/.test(line)) {
      skipTest = true;
      continue;
    }
    if (/^  \S/.test(line)) skipTest = false;
    if (!skipTest) result.push(line);
  }
  if (!result.some((line) => /^  dependencies:/.test(line))) {
    throw new Error("spago.yaml has no package dependencies");
  }
  return result.join("\n").trimEnd() + "\n";
}

function resolveOptimizer(root) {
  const configured = process.env.GOPURS_PBO_DIR;
  if (configured) {
    const directory = resolve(configured);
    if (!existsSync(join(directory, "spago.yaml"))) {
      throw new Error(`GOPURS_PBO_DIR does not contain spago.yaml: ${directory}`);
    }
    return directory;
  }
  try {
    const lock = readJson(join(root, "spago.lock"));
    const entry = lock.packages?.["backend-optimizer"];
    if (entry?.type === "git" && entry.rev) {
      const fetched = join(root, ".spago/p/backend-optimizer", entry.rev);
      if (existsSync(join(fetched, "spago.yaml"))) return fetched;
    }
  } catch {
    // Fall through to the actionable error below.
  }
  throw new Error("Missing backend optimizer; run npm ci or set GOPURS_PBO_DIR to a checkout");
}

// The compiler's own bootstrap must pair its runtime with the library versions
// in its source manifest. Read committed trees without checking out or changing
// the developer's sibling libraries. An explicit library root opts into using
// the caller's development sources instead.
export function prepareNativeGoLibraries(root, workspace, { libraryRoot = process.env.GOPURS_LIBRARY_ROOT } = {}) {
  if (libraryRoot) return resolve(libraryRoot);
  const manifestBytes = readFileSync(join(root, "bootstrap-sources.json"), "utf8");
  const manifest = JSON.parse(manifestBytes);
  if (manifest.schema !== 1 || !manifest.libraries || typeof manifest.libraries !== "object") {
    throw new Error("Invalid bootstrap library manifest");
  }
  // Spago discovers nested spago.yaml files as workspace packages. Keep these
  // dependency snapshots in its excluded cache tree so extraPackages remains
  // their only registration, including support packages without a workspace.
  const destination = join(workspace, ".spago", "bootstrap-libraries");
  mkdirSync(destination, { recursive: true });
  for (const [name, source] of Object.entries({ ...manifest.support, ...manifest.libraries })) {
    if (!/^[a-z0-9]+(?:-[a-z0-9]+)*$/.test(name) || !/^[a-f0-9]{40}$/.test(source?.rev)) {
      throw new Error(`Invalid pinned bootstrap library: ${name}`);
    }
    const checkout = join(dirname(root), `gopurs-${name}`);
    let archive;
    try {
      archive = execFileSync("git", ["-C", checkout, "archive", "--format=tar", source.rev], {
        maxBuffer: 64 * 1024 * 1024, stdio: ["ignore", "pipe", "pipe"],
      });
    } catch (error) {
      throw new Error(
        `Cannot read pinned gopurs-${name} ${source.rev} from ${checkout}; install the manifest sources with ./bin/setup or select GOPURS_LIBRARY_ROOT explicitly`,
        { cause: error },
      );
    }
    const library = join(destination, `gopurs-${name}`);
    mkdirSync(library);
    execFileSync("tar", ["-xf", "-", "-C", library], { input: archive, stdio: ["pipe", "pipe", "pipe"] });
    if (!existsSync(join(library, "spago.yaml"))) throw new Error(`Pinned gopurs-${name} has no spago.yaml`);
  }
  writeFileSync(join(destination, "bootstrap-sources.json"), manifestBytes);
  return destination;
}

export function nativeWorkspaceConfig(root, { runtime = "go", purust = resolve(root, "../../purust/purust"), libraryRoot } = {}) {
  if (!["go", "rust"].includes(runtime)) throw new Error(`Unknown compiler runtime: ${runtime}`);
  const optimizer = resolveOptimizer(root);
  const packages = new Map([["backend-optimizer", optimizer]]);
  const family = runtime === "rust" ? "purust-" : "gopurs-";
  const libraries = runtime === "rust" ? dirname(purust) : resolve(libraryRoot || process.env.GOPURS_LIBRARY_ROOT || dirname(root));
  for (const name of readdirSync(libraries).sort()) {
    if (!name.startsWith(family)) continue;
    const directory = join(libraries, name);
    if (!statSync(directory, { throwIfNoEntry: false })?.isDirectory()) continue;
    const config = join(directory, "spago.yaml");
    if (!statSync(config, { throwIfNoEntry: false })?.isFile()) continue;
    // The checkout name supplies the registry override, even if its package is
    // prefixed: gopurs-node-process provides the node-process dependency.
    const packageName = name.slice(family.length);
    if (packages.has(packageName)) throw new Error(`Duplicate local package: ${packageName}`);
    packages.set(packageName, directory);
  }
  // A Go bootstrap needs the pinned Argonaut codec support checkout. The
  // sibling scan above supplies it; fail early rather than silently falling
  // back to the registry package's JavaScript-only FFI.
  if (runtime === "go" && !packages.has("argonaut-codecs")) {
    throw new Error(
      "Missing gopurs-argonaut-codecs; run ./bin/setup to install the pinned Argonaut codec support checkout",
    );
  }
  // Select the host runtime's library representations, retaining gopurs's PBO
  // sources and Go code generator for either native executable.
  return (
    compilerPackage(root) +
    "workspace:\n  packageSet:\n    registry: 77.10.1\n  extraPackages:\n" +
    [...packages].map(([name, directory]) => `    ${name}:\n      path: ${JSON.stringify(directory)}\n`).join("")
  );
}

function localCompilerCandidates(root) {
  const dist = resolve(root, "../../purescript/.stack-work/dist");
  const candidates = [];
  if (!existsSync(dist)) return candidates;
  for (const platform of readdirSync(dist, { withFileTypes: true })) {
    if (!platform.isDirectory()) continue;
    const platformPath = join(dist, platform.name);
    for (const build of readdirSync(platformPath, { withFileTypes: true })) {
      if (!build.isDirectory()) continue;
      const path = join(platformPath, build.name, "build/purs/purs");
      const info = statSync(path, { throwIfNoEntry: false });
      if (info?.isFile()) candidates.push({ path, modified: info.mtimeMs });
    }
  }
  return candidates;
}

export function findTypedCompiler(root, configuredCompiler = process.env.GOPURS_PURS) {
  const compiler = configuredCompiler
    ? resolve(configuredCompiler)
    : localCompilerCandidates(root).sort((left, right) => right.modified - left.modified)[0]?.path;
  if (!compiler) throw new Error("No local TAST compiler found; set GOPURS_PURS to the typed fork's purs executable");
  accessSync(compiler, constants.X_OK);
  return realpathSync(compiler);
}

// Return a summary for the caller to log. Validation does not depend on where
// the workspace or its logs live, and must finish before Go generation begins.
export function verifyTypedOutput(output, compiler) {
  let modules = 0;
  let types = 0;
  for (const entry of readdirSync(output, { withFileTypes: true })) {
    if (!entry.isDirectory()) continue;
    const file = join(output, entry.name, "corefn.json");
    if (!existsSync(file)) continue;
    const metadata = readJson(file);
    if (!["typeTable", "dataDecls", "classDecls"].every((key) => Array.isArray(metadata[key]))) {
      throw new Error(
        `${file} lacks TAST type metadata. ${compiler} did not produce the required format; select the typed fork with GOPURS_PURS`,
      );
    }
    modules++;
    types += metadata.typeTable.length;
  }
  if (!modules || !types) throw new Error("The compiler produced no typed CoreFn modules");
  return { modules, types };
}
