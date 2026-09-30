import { accessSync, constants, existsSync, readFileSync, readdirSync, realpathSync, statSync } from "node:fs";
import { dirname, join, resolve } from "node:path";

// Copy the package section verbatim, excluding its optional test stanza. This
// preserves dependency constraints without a second YAML parser dependency.
function compilerPackage(root) {
  const lines = readFileSync(join(root, "spago.yaml"), "utf8").split(/\r?\n/);
  const start = lines.findIndex(line => /^package:\s*$/.test(line));
  if (start < 0) throw new Error("spago.yaml has no package section");
  let skipTest = false;
  const result = [];
  for (const line of lines.slice(start)) {
    if (result.length && /^\S/.test(line)) break;
    if (/^  test:/.test(line)) { skipTest = true; continue; }
    if (/^  \S/.test(line)) skipTest = false;
    if (!skipTest) result.push(line);
  }
  if (!result.some(line => /^  dependencies:/.test(line))) {
    throw new Error("spago.yaml has no package dependencies");
  }
  return result.join("\n").trimEnd() + "\n";
}

export function nativeWorkspaceConfig(root) {
  const optimizer = resolve(root, "../../purescript-backend-optimizer-gopurs");
  if (!existsSync(join(optimizer, "spago.yaml"))) {
    throw new Error(`Missing local backend optimizer: ${optimizer}`);
  }
  const packages = new Map([["backend-optimizer", optimizer]]);
  for (const name of readdirSync(dirname(root)).sort()) {
    if (!name.startsWith("gopurs-")) continue;
    const directory = join(dirname(root), name);
    if (!statSync(directory, { throwIfNoEntry: false })?.isDirectory()) continue;
    const config = join(directory, "spago.yaml");
    if (!statSync(config, { throwIfNoEntry: false })?.isFile()) continue;
    // The checkout name supplies the registry override, even if its package is
    // prefixed: gopurs-node-process provides the node-process dependency.
    const packageName = name.slice("gopurs-".length);
    if (packages.has(packageName)) throw new Error(`Duplicate local package: ${packageName}`);
    packages.set(packageName, directory);
  }
  // Native library overrides keep foreign types such as Map opaque in Go.
  return compilerPackage(root) + "workspace:\n  packageSet:\n    registry: 77.10.1\n  extraPackages:\n" +
    [...packages].map(([name, directory]) => `    ${name}:\n      path: ${JSON.stringify(directory)}\n`).join("");
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

export function findTypedCompiler(root, configuredCompiler) {
  const compiler = configuredCompiler ? resolve(configuredCompiler)
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
    const metadata = JSON.parse(readFileSync(file, "utf8"));
    if (!["typeTable", "dataDecls", "classDecls"].every(key => Array.isArray(metadata[key]))) {
      throw new Error(`${file} lacks TAST type metadata. ${compiler} did not produce the required format; select the typed fork with GOPURS_PURS`);
    }
    modules++;
    types += metadata.typeTable.length;
  }
  if (!modules || !types) throw new Error("The compiler produced no typed CoreFn modules");
  return { modules, types };
}
