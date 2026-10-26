#!/usr/bin/env node
import { execFileSync } from "node:child_process";
import { createHash } from "node:crypto";
import { chmodSync, copyFileSync, existsSync, mkdirSync, readFileSync, readdirSync, writeFileSync } from "node:fs";
import { arch, platform } from "node:os";
import { join, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const root = fileURLToPath(new URL("../", import.meta.url));
const options = {};
for (let index = 2; index < process.argv.length; index++) {
  const arg = process.argv[index];
  if (arg === "--help") {
    console.log("Usage: node tools/release-prepare.mjs --out DIRECTORY --native FILE [--preview]\nPackages the already built compiler, its lockfile inventory, build metadata and SHA-256 checksums. Final preparation requires a clean tree at the matching annotated version tag.");
    process.exit(0);
  }
  if (arg === "--preview") options.preview = true;
  else if (["--out", "--native"].includes(arg) && process.argv[index + 1]) options[arg.slice(2)] = resolve(process.argv[++index]);
  else throw new Error(`Unknown or incomplete option: ${arg}`);
}
if (!options.out || !options.native) throw new Error("--out and --native are required");
const command = (program, args, cwd = root) => execFileSync(program, args, { cwd, encoding: "utf8", maxBuffer: 16 * 1024 * 1024 });
const json = file => JSON.parse(readFileSync(join(root, file), "utf8"));
const sha256 = file => createHash("sha256").update(readFileSync(file)).digest("hex");
const metadata = json("package.json");
const version = metadata.version;
if (!/^\d+\.\d+\.\d+(?:-rc\.\d+)?$/.test(version)) throw new Error(`Unsupported release version: ${version}`);
if (json("package-lock.json").packages[""].version !== version) throw new Error("package.json and package-lock.json release versions differ");
const target = `${platform()}-${{ x64: "amd64", arm64: "arm64" }[arch()] ?? arch()}`;
if (!["linux-amd64", "linux-arm64", "darwin-arm64"].includes(target)) throw new Error(`Unsupported release target: ${target}`);
const status = command("git", ["status", "--porcelain", "--untracked-files=all"]).trim();
const commit = command("git", ["rev-parse", "HEAD"]).trim();
const tag = `v${version}`;
const changelog = readFileSync(join(root, "CHANGELOG.md"), "utf8");
if (!options.preview) {
  if (status) throw new Error("Release preparation requires a clean source tree; use --preview only for local review");
  if (process.env.GITHUB_REF_TYPE === "tag" && process.env.GITHUB_REF_NAME !== tag) throw new Error(`Workflow tag does not match ${tag}`);
  if (command("git", ["cat-file", "-t", `refs/tags/${tag}`]).trim() !== "tag"
      || command("git", ["rev-parse", `${tag}^{commit}`]).trim() !== commit) {
    throw new Error(`HEAD must be the annotated tag ${tag}`);
  }
  if (!changelog.includes(`## [${version}]`)) throw new Error(`Missing final changelog section for ${version}`);
}
if (existsSync(options.out) && readdirSync(options.out).length) throw new Error(`Output directory is not empty: ${options.out}`);
for (const file of ["bin/gopurs.js", "tools/ffi_gen.wasm", "tools/wasm_exec.js", "bootstrap-sources.json", "flake.lock"]) {
  if (!existsSync(join(root, file))) throw new Error(`Missing release input: ${file}`);
}
const sources = json("bootstrap-sources.json");
if (json("flake.lock").nodes["purs-tast"]?.locked?.rev !== sources.compiler.rev) throw new Error("Nix and bootstrap TAST compiler revisions differ");
const nativeInfo = command("go", ["version", "-m", options.native]).trim().split("\n");
// The first line names the supplied file. Keep only its Go version so build
// metadata does not expose a maintainer's absolute checkout path.
nativeInfo[0] = nativeInfo[0].slice(nativeInfo[0].lastIndexOf(": ") + 2);
const nativeBuildInfo = nativeInfo.join("\n");
const expectedOs = platform();
const expectedArch = arch() === "x64" ? "amd64" : arch();
if (!nativeBuildInfo.includes(`GOOS=${expectedOs}`) || !nativeBuildInfo.includes(`GOARCH=${expectedArch}`)) throw new Error(`Native executable does not match ${target}`);
mkdirSync(options.out, { recursive: true });
const prefix = `gopurs-${version}-${target}`;
const writeJson = (name, value) => writeFileSync(join(options.out, name), JSON.stringify(value, null, 2) + "\n");
const [archive] = JSON.parse(command("npm", ["pack", "--ignore-scripts", "--json", "--pack-destination", options.out]));
const nativeName = prefix;
copyFileSync(options.native, join(options.out, nativeName));
chmodSync(join(options.out, nativeName), 0o755);
const inputFiles = ["package-lock.json", "spago.lock", "flake.lock", "bootstrap-sources.json"];
for (const file of inputFiles) copyFileSync(join(root, file), join(options.out, `${prefix}-${file}`));
const bootstrapName = `${prefix}-bootstrap-checksums.txt`;
const bootstrapFiles = ["bin/gopurs.js", "tools/ffi_gen.wasm", "tools/wasm_exec.js"];
writeFileSync(join(options.out, bootstrapName), bootstrapFiles.map(file => `${sha256(join(root, file))}  ${file}\n`).join(""));

// Lockfiles describe the package build input inventory, including development
// tools. They do not claim every component survives optimizer dead-code removal.
const components = [];
for (const [path, entry] of Object.entries(json("package-lock.json").packages)) {
  if (!path || !entry.version) continue;
  const name = path.split("node_modules/").at(-1);
  const component = { type: "library", "bom-ref": `npm:${path}`, name, version: entry.version, scope: entry.dev ? "optional" : "required", purl: `pkg:npm/${name.replace("@", "%40")}@${entry.version}` };
  if (entry.license) component.licenses = [{ expression: entry.license }];
  if (entry.resolved?.startsWith("https://")) component.externalReferences = [{ type: "distribution", url: entry.resolved }];
  components.push(component);
}
const spago = json("spago.lock");
for (const [name, entry] of Object.entries(spago.packages)) {
  const component = { type: "library", "bom-ref": `purescript:${name}`, name, version: entry.rev ?? entry.version ?? "local", scope: "required", properties: [{ name: "gopurs:ecosystem", value: "purescript" }] };
  if (entry.url) component.externalReferences = [{ type: "vcs", url: entry.url }];
  components.push(component);
}
for (const [name, source] of Object.entries({ "purs-tast": sources.compiler, ...Object.fromEntries(Object.entries(sources.libraries).map(([name, source]) => [`gopurs-${name}`, source])) })) {
  const component = { type: name === "purs-tast" ? "application" : "library", "bom-ref": `bootstrap:${name}`, name, version: source.rev ?? source.sha256 ?? "vendored", scope: "optional" };
  if (source.url) component.externalReferences = [{ type: "vcs", url: source.url }];
  components.push(component);
}
for (const [name, source] of Object.entries(sources.vendored ?? {})) {
  components.push({ type: "library", "bom-ref": `vendored:${name}`, name, version: source.version ?? "vendored", scope: "optional", properties: [{ name: "gopurs:source-path", value: source.path }], ...(source.license ? { licenses: [{ expression: source.license }] } : {}) });
}
const goVersion = command("go", ["version"]).trim();
components.push({ type: "application", "bom-ref": "toolchain:go", name: "go", version: goVersion.split(" ")[2], scope: "required" });
components.push({ type: "application", "bom-ref": "toolchain:node", name: "node", version: process.version.slice(1), scope: "optional" });
const sbomName = `${prefix}-sbom.cdx.json`;
writeJson(sbomName, {
  bomFormat: "CycloneDX", specVersion: "1.6", version: 1,
  metadata: {
    component: { type: "application", "bom-ref": "gopurs", name: "gopurs", version, licenses: [{ expression: metadata.license }] },
    properties: [{ name: "gopurs:inventory-scope", value: "Compiler source/build dependencies, including npm development tools and native bootstrap libraries; not an optimizer reachability report" }, { name: "gopurs:source-commit", value: commit }, { name: "gopurs:target", value: target }],
  },
  components,
  dependencies: [{ ref: "gopurs", dependsOn: components.map(component => component["bom-ref"]) }],
});
const notesName = `${prefix}-release-notes.md`;
const section = changelog.split(/^## /m).find(section => section.startsWith(`[${version}]`))
  ?? (options.preview ? changelog.split(/^## /m).find(section => section.startsWith("[Unreleased]")) : null);
writeFileSync(join(options.out, notesName), `# gopurs ${version}${options.preview ? " — local preview" : ""}\n\n${section ?? "No release notes prepared."}\n`);
const artifacts = [archive.filename, nativeName, sbomName, notesName, bootstrapName, ...inputFiles.map(file => `${prefix}-${file}`)].sort().map(name => ({ name, sha256: sha256(join(options.out, name)) }));
writeJson(`${prefix}-build.json`, {
  schema: 1, version, target, source: { commit, tag: options.preview ? null : tag, dirty: Boolean(status) }, preview: Boolean(options.preview),
  tools: { node: process.version, npm: command("npm", ["--version"]).trim(), go: goVersion },
  nativeBuildInfo,
  inputs: Object.fromEntries(inputFiles.map(file => [file, sha256(join(root, file))])),
  artifacts,
  attestation: "This file is unsigned build metadata. The release workflow signs provenance and SBOM attestations with GitHub OIDC/Sigstore.",
});
const files = readdirSync(options.out).sort();
writeFileSync(join(options.out, `${prefix}-SHA256SUMS`), files.map(file => `${sha256(join(options.out, file))}  ${file}\n`).join(""));
console.log(JSON.stringify({ version, target, preview: Boolean(options.preview), directory: options.out, tarball: archive.filename, native: nativeName, sbom: sbomName }, null, 2));
