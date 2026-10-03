import {
  accessSync, chmodSync, constants, copyFileSync, mkdirSync, mkdtempSync,
  readFileSync, renameSync, rmSync, symlinkSync, writeFileSync,
} from "node:fs";
import { tmpdir } from "node:os";
import { delimiter, dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { CommandRunner, Interrupted } from "./command-runner.mjs";
import { findTypedCompiler, nativeWorkspaceConfig, verifyTypedOutput } from "./native-workspace.mjs";
import { corePackages, prepareFixture } from "./test-workspace.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const args = process.argv.slice(2);
if (args.includes("--help")) {
  console.log("Usage: npm run build:rust -- [--keep-workspace]\nBuild bin/gopurs-rust: the gopurs Go code generator hosted in Rust.\nPURUST_DIR selects the Rust backend checkout; PURUST_JS=1 uses its JavaScript executable for this bootstrap only.\nGOPURS_PURS selects the typed frontend. GOPURS_NATIVE_TMPDIR selects the workspace parent; GOPURS_RUST_OUTPUT overrides the installed executable.\nGOPURS_KEEP_WORKSPACE=1 also retains successful builds.");
  process.exit(0);
}
if (args.some(arg => arg !== "--keep-workspace")) {
  console.error(`Unknown option: ${args.find(arg => arg !== "--keep-workspace")}`);
  process.exit(1);
}

const commands = new CommandRunner();
const purust = resolve(process.env.PURUST_DIR ?? join(root, "../../purust/purust"));
const destination = resolve(process.env.GOPURS_RUST_OUTPUT ?? join(root, "bin/gopurs-rust"));
const environment = { ...process.env,
  PATH: [join(root, "node_modules/.bin"), dirname(process.execPath), process.env.PATH ?? ""].join(delimiter),
  GOWORK: "off", CARGO_BUILD_JOBS: process.env.CARGO_BUILD_JOBS ?? "8",
  CARGO_INCREMENTAL: process.env.CARGO_INCREMENTAL ?? "0",
};
let workspace, stage = "prepare";
async function run(label, command, args, cwd, env = environment) {
  commands.checkInterrupted();
  stage = label;
  const log = join(workspace, `${label}.log`), started = Date.now();
  console.log(`[${label}] ${command} ${args.join(" ")}\n  log: ${log}`);
  const status = await commands.run(command, args, { cwd, env, log });
  if (commands.signal || status.code !== 0) process.stderr.write(readFileSync(log, "utf8"));
  commands.checkInterrupted();
  if (status.code !== 0) throw new Error(`${label} failed (${status.signal ?? status.code})`);
  console.log(`[${label}] finished in ${((Date.now() - started) / 1000).toFixed(1)} s`);
}

try {
  const compiler = findTypedCompiler(root, process.env.GOPURS_PURS);
  accessSync(join(root, "node_modules/.bin/spago"), constants.X_OK);
  accessSync(join(purust, "bin", environment.PURUST_JS === "1" ? "purust.js" : "purust-native"), constants.R_OK);
  const config = nativeWorkspaceConfig(root, { runtime: "rust", purust });
  workspace = mkdtempSync(join(process.env.GOPURS_NATIVE_TMPDIR ?? tmpdir(), "gopurs-rust-build-"));
  console.log(`Rust-hosted gopurs workspace: ${workspace}\nTAST compiler: ${compiler}`);
  writeFileSync(join(workspace, "spago.yaml"), config);
  symlinkSync(join(root, "src"), join(workspace, "src"), "dir");
  const typedBin = join(workspace, "typed-bin");
  mkdirSync(typedBin);
  symlinkSync(compiler, join(typedBin, "purs"));
  const typedEnvironment = { ...environment, PATH: typedBin + delimiter + environment.PATH };
  await run("embed-go-runtime", process.execPath, [join(root, "tools/embed-runtime.mjs")], root);
  await run("typed-corefn", "spago", ["build"], workspace, typedEnvironment);
  const output = join(workspace, "output");
  stage = "verify-tast";
  const verified = verifyTypedOutput(output, compiler);
  writeFileSync(join(workspace, "verify-tast.json"), JSON.stringify({ compiler, ...verified }, null, 2) + "\n");
  const rust = join(workspace, "rust");
  await run("generate-rust", join(purust, "bin/purust"), ["--source", output, "--out", rust,
    "--main", "Main", "--threaded"], workspace);

  // Link the same Go parser as gopurs-native. No parser subprocess or Node
  // runtime is needed when the resulting compiler generates a project's Go.
  const parser = join(rust, "Purs_Gopurs_FfiSupport", "native-parser");
  mkdirSync(parser);
  await run("native-parser", "go", ["build", "-trimpath", "-tags=carchive", "-buildmode=c-archive",
    "-o", join(parser, "libgopurs_ffi.a"), "."], join(root, "tools/ffi-gen"));
  copyFileSync(join(root, "tools/ffi-gen/rust-build.rs"), join(rust, "Purs_Gopurs_FfiSupport", "build.rs"));
  const target = resolve(environment.CARGO_TARGET_DIR ?? join(workspace, "target"));
  await run("cargo-build", "cargo", ["build", "--release", "--target-dir", target,
    "--config", "profile.release.lto=false", "--config", "profile.release.opt-level=3",
    "--config", "profile.release.debug=false", "--manifest-path", join(rust, "Cargo.toml")], workspace);
  const binary = join(target, "release", process.platform === "win32" ? "purust_output.exe" : "purust_output");

  // A fresh Go-target fixture exercises FFI parsing, generation and application
  // execution before the candidate can replace an installed compiler.
  const smoke = prepareFixture(root, workspace, join(root, "tests/passing/FFIIntegerReturns.purs"), 0, corePackages(root));
  await run("smoke-tast", "spago", ["build", "-q"], smoke.directory, typedEnvironment);
  verifyTypedOutput(join(smoke.directory, "output"), compiler);
  await run("smoke-generate-go", binary, ["--main", "Main"], smoke.directory);
  await run("smoke-go-run", "go", ["run", "./main/main.go"], join(smoke.directory, "output"));
  if (readFileSync(join(workspace, "smoke-go-run.log"), "utf8").trim() !== "Done") {
    throw new Error("The Rust-hosted compiler's Go smoke test did not finish");
  }

  commands.checkInterrupted();
  stage = "publish";
  mkdirSync(dirname(destination), { recursive: true });
  const staging = mkdtempSync(join(dirname(destination), ".gopurs-rust-"));
  try {
    const staged = join(staging, "gopurs-rust");
    copyFileSync(binary, staged);
    chmodSync(staged, 0o755);
    renameSync(staged, destination);
  } finally { rmSync(staging, { recursive: true, force: true }); }
  console.log(`Built ${destination}`);
  if (args.includes("--keep-workspace") || process.env.GOPURS_KEEP_WORKSPACE === "1") console.log(`Workspace retained: ${workspace}`);
  else rmSync(workspace, { recursive: true, force: true, maxRetries: 3 });
} catch (error) {
  console.error(`Rust bootstrap failed during ${stage}: ${error.message}`);
  if (workspace) console.error(`Workspace and logs retained: ${workspace}`);
  process.exitCode = commands.signal ? new Interrupted(commands.signal).exitCode : 1;
} finally { commands.dispose(); }
