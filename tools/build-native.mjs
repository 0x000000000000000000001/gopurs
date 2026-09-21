import { spawn } from "node:child_process";
import {
  accessSync, chmodSync, closeSync, constants, copyFileSync, existsSync, mkdirSync, mkdtempSync, openSync,
  readFileSync, readdirSync, realpathSync, renameSync, rmSync, statSync, symlinkSync, writeFileSync,
} from "node:fs";
import { tmpdir } from "node:os";
import { delimiter, dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const root = fileURLToPath(new URL("../", import.meta.url));
const args = process.argv.slice(2);
if (args.includes("--help")) {
  console.log("Usage: npm run build:native -- [--keep-workspace]\nBuild an experimental native compiler using the Node backend and local Go library checkouts.\nGOPURS_PURS=/absolute/path/to/typed/purs overrides discovery of the local compiler fork.");
  process.exit(0);
}
if (args.some(arg => arg !== "--keep-workspace")) {
  console.error(`Unknown option: ${args.find(arg => arg !== "--keep-workspace")}`);
  process.exit(1);
}

// Copy the package section verbatim, excluding its optional test stanza. This
// preserves dependency constraints without a second YAML parser dependency.
function compilerPackage() {
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

function nativeWorkspaceConfig() {
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
    // Override the registry name even when the checkout declares a prefixed
    // package name, e.g. gopurs-node-process supplies node-process.
    const packageName = name.slice("gopurs-".length);
    if (packages.has(packageName)) throw new Error(`Duplicate local package: ${packageName}`);
    packages.set(packageName, directory);
  }
  // Native library overrides keep foreign types such as Map opaque in Go.
  return compilerPackage() + "workspace:\n  packageSet:\n    registry: 77.10.1\n  extraPackages:\n" +
    [...packages].map(([name, directory]) => `    ${name}:\n      path: ${JSON.stringify(directory)}\n`).join("");
}

function typedCompiler() {
  let compiler;
  if (process.env.GOPURS_PURS) compiler = resolve(process.env.GOPURS_PURS);
  else {
    const dist = resolve(root, "../../purescript/.stack-work/dist");
    const candidates = [];
    if (existsSync(dist)) {
      for (const platform of readdirSync(dist, { withFileTypes: true })) {
        if (!platform.isDirectory()) continue;
        const platformPath = join(dist, platform.name);
        for (const build of readdirSync(platformPath, { withFileTypes: true })) {
          if (!build.isDirectory()) continue;
          const candidate = join(platformPath, build.name, "build/purs/purs");
          const info = statSync(candidate, { throwIfNoEntry: false });
          if (info?.isFile()) candidates.push({ path: candidate, modified: info.mtimeMs });
        }
      }
    }
    candidates.sort((left, right) => right.modified - left.modified);
    compiler = candidates[0]?.path;
  }
  if (!compiler) throw new Error("No local TAST compiler found; set GOPURS_PURS to the typed fork's purs executable");
  accessSync(compiler, constants.X_OK);
  return realpathSync(compiler);
}

function verifyTypedOutput(output, compiler) {
  let modules = 0;
  let types = 0;
  for (const entry of readdirSync(output, { withFileTypes: true })) {
    if (!entry.isDirectory()) continue;
    const file = join(output, entry.name, "corefn.json");
    if (!existsSync(file)) continue;
    const module = JSON.parse(readFileSync(file, "utf8"));
    if (!["typeTable", "dataDecls", "classDecls"].every(key => Array.isArray(module[key]))) {
      throw new Error(`${file} lacks TAST type metadata. ${compiler} did not produce the required format; select the typed fork with GOPURS_PURS`);
    }
    modules++;
    types += module.typeTable.length;
  }
  if (!modules || !types) throw new Error("The compiler produced no typed CoreFn modules");
  const result = `Verified TAST from ${compiler}: ${modules} modules, ${types} types\n`;
  writeFileSync(join(workspace, "verify-tast.log"), result);
  console.log(result.trimEnd());
}

let active;
let interrupted;
let workspace;
let stage = "prepare";
const handlers = new Map(["SIGINT", "SIGTERM"].map(signal => {
  const handler = () => {
    interrupted = signal;
    if (active?.pid) {
      try { process.kill(-active.pid, signal); }
      catch (error) { if (error.code !== "ESRCH") throw error; }
    }
  };
  process.on(signal, handler);
  return [signal, handler];
}));
const environment = {
  ...process.env,
  PATH: [join(root, "node_modules/.bin"), dirname(process.execPath), process.env.PATH ?? ""].join(delimiter),
  GOWORK: "off",
};

async function run(label, command, commandArgs, cwd, env = environment) {
  if (interrupted) throw new Error(`Interrupted by ${interrupted}`);
  stage = label;
  const log = join(workspace, `${label}.log`);
  const fd = openSync(log, "w");
  const started = Date.now();
  console.log(`[${label}] ${command} ${commandArgs.join(" ")}\n  log: ${log}`);
  let status;
  try {
    active = spawn(command, commandArgs, { cwd, env, detached: true, stdio: ["ignore", fd, fd] });
    status = await new Promise((resolveStatus, reject) => {
      active.once("error", reject);
      active.once("close", (code, signal) => resolveStatus({ code, signal }));
    });
  } finally {
    active = null;
    closeSync(fd);
  }
  if (interrupted || status.code !== 0) {
    process.stderr.write(readFileSync(log, "utf8"));
    throw new Error(interrupted ? `Interrupted by ${interrupted}` : `${label} failed (${status.signal ?? "exit " + status.code})`);
  }
  console.log(`[${label}] finished in ${((Date.now() - started) / 1000).toFixed(1)} s`);
}

try {
  for (const tool of ["purs", "spago"]) {
    if (!existsSync(join(root, "node_modules/.bin", tool))) {
      throw new Error(`Missing local ${tool}; install the compiler's npm dependencies first`);
    }
  }
  const config = nativeWorkspaceConfig();
  const compiler = typedCompiler();
  workspace = mkdtempSync(join(tmpdir(), "gopurs-native-build-"));
  writeFileSync(join(workspace, "spago.yaml"), config);
  symlinkSync(join(root, "src"), join(workspace, "src"), "dir");
  const typedBin = join(workspace, "typed-bin");
  mkdirSync(typedBin);
  symlinkSync(compiler, join(typedBin, "purs"));
  const typedEnvironment = { ...environment, PATH: typedBin + delimiter + environment.PATH };
  console.log(`Native bootstrap workspace: ${workspace}`);
  console.log(`TAST compiler: ${compiler}`);
  await run("node-backend", "npm", ["run", "build"], root);
  await run("typed-corefn", "spago", ["build"], workspace, typedEnvironment);
  stage = "verify-tast";
  const output = join(workspace, "output");
  verifyTypedOutput(output, compiler);
  await run("generate-go", process.execPath, [join(root, "bin/gopurs.js"), "--main", "Main"], workspace);
  await run("native-parser", process.execPath, [join(root, "tools/prepare-native-output.mjs"), output], root);
  const binary = join(workspace, "gopurs-native");
  await run("go-build", "go", ["build", "-trimpath", "-o", binary, "./main/main.go"], output);
  if (interrupted) throw new Error(`Interrupted by ${interrupted}`);
  stage = "publish binary";
  const destination = join(root, "bin/gopurs-native");
  const staging = mkdtempSync(join(root, "bin/.gopurs-native-"));
  try {
    const stagedBinary = join(staging, "gopurs-native");
    copyFileSync(binary, stagedBinary);
    chmodSync(stagedBinary, 0o755);
    renameSync(stagedBinary, destination);
  } finally {
    rmSync(staging, { recursive: true, force: true });
  }
  console.log(`Built ${destination}`);
  if (args.includes("--keep-workspace")) console.log(`Workspace retained: ${workspace}`);
  else rmSync(workspace, { recursive: true, force: true });
} catch (error) {
  console.error(`Native bootstrap failed during ${stage}: ${error.message}`);
  if (workspace) console.error(`Workspace and logs retained: ${workspace}`);
  process.exitCode = interrupted === "SIGINT" ? 130 : interrupted === "SIGTERM" ? 143 : 1;
} finally {
  for (const [signal, handler] of handlers) process.off(signal, handler);
}
