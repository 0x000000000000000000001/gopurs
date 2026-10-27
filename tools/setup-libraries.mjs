#!/usr/bin/env node
import { execFileSync } from "node:child_process";
import { existsSync, mkdirSync, mkdtempSync, readFileSync, renameSync, rmSync, statSync } from "node:fs";
import { join, resolve } from "node:path";

const [root, destination, lockedFlag, ...names] = process.argv.slice(2);
if (!root || !destination || !["0", "1"].includes(lockedFlag) || !names.length) {
  console.error("Use bin/setup [--core | --all] [--locked].");
  process.exit(2);
}

function git(args, cwd, capture = true) {
  return execFileSync("git", args, { cwd, encoding: "utf8", stdio: capture ? ["ignore", "pipe", "pipe"] : "inherit" })?.trim();
}

try {
  const manifest = JSON.parse(readFileSync(join(root, "bootstrap-sources.json"), "utf8"));
  if (manifest.schema !== 1) throw new Error("Unsupported bootstrap-sources.json schema");
  const libraries = resolve(destination);
  const selected = names.map(name => {
    const source = manifest.libraries[name];
    if (!/^[a-z0-9]+(?:-[a-z0-9]+)*$/.test(name) || !source || !/^[a-f0-9]{40}$/.test(source.rev)) {
      throw new Error(`Missing pinned source for ${name}`);
    }
    return { name, source, target: join(libraries, `gopurs-${name}`) };
  });

  // Validate every existing destination before changing anything. A developer's
  // different branch or uncommitted work must never be reset by setup.
  for (const { name, source, target } of selected) {
    if (!existsSync(target)) continue;
    if (!statSync(target).isDirectory() || !existsSync(join(target, "spago.yaml"))) {
      throw new Error(`Existing destination is not a configured library: ${target}`);
    }
    if (lockedFlag === "1") {
      const actual = git(["rev-parse", "HEAD"], target);
      if (actual !== source.rev) throw new Error(`gopurs-${name} is at ${actual}; release requires ${source.rev}. Use an empty GOPURS_LIBRARY_ROOT.`);
      if (git(["status", "--porcelain"], target)) throw new Error(`gopurs-${name} has local changes; use an empty GOPURS_LIBRARY_ROOT for a locked installation.`);
    }
  }

  mkdirSync(libraries, { recursive: true });
  for (const { name, source, target } of selected) {
    if (existsSync(target)) {
      console.log(`[${lockedFlag === "1" ? "VERIFIED" : "SKIP"}] gopurs-${name}`);
      continue;
    }
    const temporary = mkdtempSync(join(libraries, `.gopurs-${name}-`));
    try {
      const checkout = join(temporary, "checkout");
      console.log(`[CLONE] gopurs-${name} ${source.rev}`);
      git(["clone", "--no-checkout", "--filter=blob:none", source.url, checkout], root, false);
      git(["checkout", "--detach", source.rev], checkout, false);
      if (!existsSync(join(checkout, "spago.yaml"))) throw new Error(`Pinned gopurs-${name} lacks spago.yaml`);
      renameSync(checkout, target);
    } finally {
      rmSync(temporary, { recursive: true, force: true });
    }
  }
  console.log(`Installed ${selected.length} library checkouts from bootstrap-sources.json.`);
} catch (error) {
  console.error(`[setup] ${error.message}`);
  process.exitCode = 1;
}
