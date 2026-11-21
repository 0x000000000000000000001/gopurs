#!/usr/bin/env node
import { execFileSync } from "node:child_process";
import { existsSync, mkdirSync, readFileSync, rmSync, statSync, writeFileSync } from "node:fs";
import { join, resolve } from "node:path";
import { fileURLToPath } from "node:url";

// A port declares exact Go module versions independently of its Spago packages.
// Missing manifests mean no external Go dependencies, not an implicit lookup.
export function libraryGoRequirements(ports) {
  const requirements = new Map();
  for (const port of ports) {
    if (!statSync(resolve(port), { throwIfNoEntry: false })?.isDirectory()) {
      throw new Error(`Go dependency port is not a directory: ${resolve(port)}`);
    }
    const path = join(resolve(port), "go-dependencies.json");
    if (!existsSync(path)) continue;
    let manifest;
    try {
      manifest = JSON.parse(readFileSync(path, "utf8"));
    } catch (error) {
      throw new Error(`Invalid Go dependency manifest ${path}: ${error.message}`);
    }
    if (
      manifest?.schema !== 1 ||
      !manifest.require ||
      typeof manifest.require !== "object" ||
      Array.isArray(manifest.require) ||
      Object.keys(manifest).some((key) => key !== "schema" && key !== "require")
    ) {
      throw new Error(`Invalid Go dependency manifest: ${path}`);
    }
    for (const [name, version] of Object.entries(manifest.require)) {
      if (
        !/^[A-Za-z0-9][A-Za-z0-9._~/-]*$/.test(name) ||
        typeof version !== "string" ||
        !/^v\d+\.\d+\.\d+(?:-[0-9A-Za-z.-]+)?(?:\+incompatible)?$/.test(version)
      ) {
        throw new Error(`Expected an exact Go module version in ${path}: ${name}@${version}`);
      }
      const previous = requirements.get(name);
      if (previous && previous.version !== version) {
        throw new Error(
          `Conflicting Go dependency ${name}: ${previous.version} (${previous.path}) and ${version} (${path})`,
        );
      }
      requirements.set(name, { version, path });
    }
  }
  return new Map([...requirements].sort(([a], [b]) => a.localeCompare(b)));
}

function executeGo(args, options) {
  return execFileSync("go", args, { ...options, encoding: "utf8", timeout: 120000, maxBuffer: 8 * 1024 * 1024 });
}

export function prepareLibraryGoDependencies(output, ports, { runGo = executeGo } = {}) {
  const requirements = libraryGoRequirements(ports);
  if (requirements.size === 0) return [];
  output = resolve(output);
  mkdirSync(output, { recursive: true });
  const original = new Map(
    ["go.mod", "go.sum"].map((name) => {
      const path = join(output, name);
      return [path, existsSync(path) ? readFileSync(path) : null];
    }),
  );
  const options = { cwd: output, env: { ...process.env, GOWORK: "off", GOTOOLCHAIN: "local", GOFLAGS: "" } };
  const run = (args) => runGo(args, options);
  const inspect = () => JSON.parse(run(["mod", "edit", "-json"]));
  try {
    if (original.get(join(output, "go.mod")) === null) {
      writeFileSync(join(output, "go.mod"), "module gopurs/output\n\ngo 1.22\n");
    }
    const before = inspect();
    const existing = new Map((before.Require ?? []).map((item) => [item.Path, item.Version]));
    for (const [name, { version, path }] of requirements) {
      if (existing.has(name) && existing.get(name) !== version) {
        throw new Error(
          `Go dependency ${name}: application pins ${existing.get(name)}, but ${path} requires ${version}. Reconcile the versions explicitly.`,
        );
      }
    }
    const missing = [...requirements].filter(([name]) => !existing.has(name));
    // Exact go get arguments install the module's declared transitive graph.
    // tidy would discard these requirements before generated sources exist.
    if (missing.length) run(["get", ...missing.map(([name, { version }]) => `${name}@${version}`)]);
    const after = inspect();
    const selected = new Map((after.Require ?? []).map((item) => [item.Path, item.Version]));
    for (const [name, version] of existing) {
      if (selected.get(name) !== version)
        throw new Error(
          `Go dependency preparation would change application requirement ${name}@${version}; reconcile its version explicitly.`,
        );
    }
    for (const [name, { version }] of requirements) {
      if (selected.get(name) !== version)
        throw new Error(
          `Go selected ${name}@${selected.get(name)} instead of the declared ${version}. Reconcile the port manifests.`,
        );
    }
    if (
      JSON.stringify(before.Module) !== JSON.stringify(after.Module) ||
      before.Go !== after.Go ||
      JSON.stringify(before.Toolchain) !== JSON.stringify(after.Toolchain)
    ) {
      throw new Error(
        "Go dependency preparation would change the application module or toolchain directive; update it explicitly first.",
      );
    }
    run(["mod", "download", "all"]);
    return [...requirements].map(([name, { version }]) => `${name}@${version}`);
  } catch (error) {
    // A failed download or version conflict must not leave half-updated user
    // manifests. The shared Go download cache may still contain fetched data.
    for (const [path, content] of original) {
      if (content === null) rmSync(path, { force: true });
      else if (!existsSync(path) || !readFileSync(path).equals(content)) writeFileSync(path, content);
    }
    throw error;
  }
}

if (process.argv[1] && resolve(process.argv[1]) === fileURLToPath(import.meta.url)) {
  try {
    const [output, ...ports] = process.argv.slice(2);
    if (!output || !ports.length) throw new Error("Usage: node tools/library-go-dependencies.mjs OUTPUT PORT...");
    const dependencies = prepareLibraryGoDependencies(output, ports);
    if (dependencies.length) console.log(`[gopurs] prepared Go dependencies: ${dependencies.join(", ")}`);
  } catch (error) {
    console.error(`[gopurs] ${error.message}`);
    process.exitCode = 1;
  }
}
