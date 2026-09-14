import { execFileSync } from "node:child_process";
import { cpSync, existsSync, mkdirSync, mkdtempSync, readFileSync, statSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { basename, dirname, join } from "node:path";

export function corePackages(root) {
  // bin/pkg remains the shared list used by bin/setup and the test runner.
  const names = execFileSync("bash", ["-c", 'source "$1"; printf "%s\\n" "${CORE_PACKAGES[@]}"', "pkg", join(root, "bin/pkg")], { encoding: "utf8" }).trim().split(/\s+/);
  for (const name of names) {
    if (!statSync(join(dirname(root), "gopurs-" + name), { throwIfNoEntry: false })?.isDirectory()) {
      throw new Error(`Missing sibling checkout: gopurs-${name}. Run ./bin/setup first.`);
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

  const name = basename(file, ".purs");
  const directory = join(workspace, `${index}-${name}`);
  const source = join(directory, "src");
  mkdirSync(source, { recursive: true });
  mkdirSync(join(directory, "logs"));
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
    "workspace:\n  packageSet:\n    registry: 77.10.1\n  extraPackages:\n" + packages.map(name =>
      `    ${name}:\n      path: ${JSON.stringify(join(dirname(root), "gopurs-" + name))}\n`).join("");
  writeFileSync(join(directory, "spago.yaml"), config);
  return { name, directory, snapshotFfi: /^-- @snapshot-ffi\r?$/m.test(content) };
}
