import { execFileSync } from "node:child_process";
import { cpSync, existsSync, mkdirSync, mkdtempSync, readFileSync, readdirSync, realpathSync, statSync, symlinkSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { basename, dirname, join, resolve } from "node:path";
import { isLibraryCheckout } from "./test-selection.mjs";

export function libraryRoot(root) {
  return resolve(process.env.GOPURS_LIBRARY_ROOT || dirname(root));
}

export function libraryPath(root, name) {
  const vendored = join(root, "support", name);
  return existsSync(join(vendored, "spago.yaml")) ? vendored : join(libraryRoot(root), "gopurs-" + name);
}

export function corePackages(root) {
  // bin/pkg remains the shared list used by bin/setup and the test runner.
  const names = execFileSync("bash", ["-c", 'source "$1"; printf "%s\\n" "${CORE_PACKAGES[@]}"', "pkg", join(root, "bin/pkg")], { encoding: "utf8" }).trim().split(/\s+/);
  for (const name of names) {
    if (!statSync(libraryPath(root, name), { throwIfNoEntry: false })?.isDirectory()) {
      throw new Error(`Missing library checkout: gopurs-${name}. Run ./bin/setup first.`);
    }
  }
  return names;
}

export function createWorkspace() {
  return mkdtempSync(join(tmpdir(), "gopurs-tests-"));
}

export function prepareFixture(root, workspace, file, index, packages) {
  const content = readFileSync(file, "utf8");
  const declared = [...content.matchAll(/^-- @dependencies: (.*)$/gm)].map(match => match[1]).join(" ").trim();
  const dependencies = declared ? declared.split(/\s+/) : packages;
  if (!dependencies.every(name => /^[a-z0-9]+(?:-[a-z0-9]+)*$/.test(name))) throw new Error(`Invalid @dependencies in ${file}`);
  // Explicit fixture dependencies can include native libraries outside the
  // bootstrap's core set. Use their local FFI instead of the registry's JS-only
  // package, while retaining the registry for dependencies without a checkout.
  const overrides = [...new Set([...packages, ...dependencies.filter(name =>
    existsSync(join(libraryPath(root, name), "spago.yaml")))])];

  const name = basename(file, ".purs");
  const directory = join(workspace, `${index}-${name}`);
  const source = join(directory, "src");
  mkdirSync(source, { recursive: true });
  mkdirSync(join(directory, "logs"), { recursive: true });
  // A fresh source/output directory replaces timestamp edits and module-name
  // cleanup. No files from tests/runner are read, modified, or removed.
  cpSync(file, join(source, "Main.purs"));
  for (const extension of ["go", "js"]) {
    const companion = file.slice(0, -5) + "." + extension;
    if (existsSync(companion)) cpSync(companion, join(source, "Main." + extension));
  }
  const companionDirectory = file.slice(0, -5);
  if (statSync(companionDirectory, { throwIfNoEntry: false })?.isDirectory()) cpSync(companionDirectory, source, { recursive: true });

  const config = "package:\n  name: runner\n  dependencies:\n" + dependencies.map(name => `    - ${name}\n`).join("") +
    "workspace:\n  packageSet:\n    registry: 77.10.1\n  extraPackages:\n" + overrides.map(name =>
      `    ${name}:\n      path: ${JSON.stringify(libraryPath(root, name))}\n`).join("");
  writeFileSync(join(directory, "spago.yaml"), config);
  return { name, directory, snapshotFfi: /^-- @snapshot-ffi\r?$/m.test(content) };
}

// Library runners assume sibling-relative paths and can clean their siblings.
// Give each target its own complete source family; never symlink a library back
// to a live checkout. Relative config symlinks stay relative inside the copy.
export function prepareModule(root, directory, name, { dev = false } = {}) {
  const packages = join(directory, "packages");
  mkdirSync(packages);
  const generated = new Set([".git", ".spago", ".cache", ".purmeta", "node_modules", "output", ".stack-work", ".pulp-cache"]);
  const parent = libraryRoot(root);
  for (const sibling of readdirSync(parent)) {
    const source = join(parent, sibling);
    if (!isLibraryCheckout(source)) continue;
    cpSync(realpathSync(source), join(packages, sibling), {
      recursive: true, verbatimSymlinks: true,
      filter: path => !generated.has(basename(path)),
    });
  }
  const support = join(root, "support");
  if (existsSync(support)) for (const name of readdirSync(support)) {
    const source = join(support, name);
    if (existsSync(join(source, "spago.yaml"))) cpSync(source, join(packages, `gopurs-${name}`), {
      recursive: true, verbatimSymlinks: true,
    });
  }
  // The compiler is consumed read-only; -c is handled once by the outer runner.
  // Library scripts call ../gopurs/bin/gopurs themselves. A private launcher
  // supplies the selected profile without changing each library's arguments.
  const compiler = join(packages, "gopurs");
  if (!dev) symlinkSync(root, compiler, "dir");
  else {
    mkdirSync(join(compiler, "bin"), { recursive: true });
    for (const entry of readdirSync(root)) {
      if (entry !== "bin") symlinkSync(join(root, entry), join(compiler, entry));
    }
    for (const entry of readdirSync(join(root, "bin"))) {
      if (entry !== "gopurs") symlinkSync(join(root, "bin", entry), join(compiler, "bin", entry));
    }
    const launcher = join(root, "bin/gopurs").replace(/'/g, "'\\''");
    writeFileSync(join(compiler, "bin/gopurs"), `#!/usr/bin/env bash\necho '[gopurs] Build mode: DEV (--dev).' >&2\nexec '${launcher}' --dev "$@"\n`, { mode: 0o755 });
  }
  return join(packages, name);
}
