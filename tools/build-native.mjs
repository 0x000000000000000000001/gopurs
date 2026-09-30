import {
  chmodSync, copyFileSync, existsSync, mkdirSync, mkdtempSync, readFileSync,
  renameSync, rmSync, symlinkSync, writeFileSync,
} from "node:fs";
import { tmpdir } from "node:os";
import { delimiter, dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import { CommandRunner, Interrupted } from "./command-runner.mjs";
import { findTypedCompiler, nativeWorkspaceConfig, verifyTypedOutput } from "./native-workspace.mjs";

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

// Stage on the destination filesystem so the final rename replaces the binary
// atomically. A failed build or copy must leave the previous compiler usable.
function publishNativeBinary(binary, destination) {
  const staging = mkdtempSync(join(dirname(destination), ".gopurs-native-"));
  try {
    const stagedBinary = join(staging, "gopurs-native");
    copyFileSync(binary, stagedBinary);
    chmodSync(stagedBinary, 0o755);
    renameSync(stagedBinary, destination);
  } finally {
    rmSync(staging, { recursive: true, force: true });
  }
}

async function buildNative({ keepWorkspace }) {
  const commands = new CommandRunner();
  const environment = {
    ...process.env,
    PATH: [join(root, "node_modules/.bin"), dirname(process.execPath), process.env.PATH ?? ""].join(delimiter),
    GOWORK: "off",
  };
  let workspace;
  let stage = "prepare";

  async function runStage(label, command, commandArgs, { cwd, env = environment }) {
    commands.checkInterrupted();
    stage = label;
    const log = join(workspace, `${label}.log`);
    const started = Date.now();
    console.log(`[${label}] ${command} ${commandArgs.join(" ")}\n  log: ${log}`);
    const status = await commands.run(command, commandArgs, { cwd, env, log });
    if (commands.signal || status.code !== 0) process.stderr.write(readFileSync(log, "utf8"));
    commands.checkInterrupted();
    if (status.code !== 0) throw new Error(`${label} failed (${status.signal ?? "exit " + status.code})`);
    console.log(`[${label}] finished in ${((Date.now() - started) / 1000).toFixed(1)} s`);
  }

  try {
    for (const tool of ["purs", "spago"]) {
      if (!existsSync(join(root, "node_modules/.bin", tool))) {
        throw new Error(`Missing local ${tool}; install the compiler's npm dependencies first`);
      }
    }
    const config = nativeWorkspaceConfig(root);
    const compiler = findTypedCompiler(root, process.env.GOPURS_PURS);
    // Record the workspace before populating it, so setup failures retain its path.
    workspace = mkdtempSync(join(tmpdir(), "gopurs-native-build-"));
    writeFileSync(join(workspace, "spago.yaml"), config);
    symlinkSync(join(root, "src"), join(workspace, "src"), "dir");
    const typedBin = join(workspace, "typed-bin");
    mkdirSync(typedBin);
    symlinkSync(compiler, join(typedBin, "purs"));
    const typedEnvironment = { ...environment, PATH: typedBin + delimiter + environment.PATH };
    const output = join(workspace, "output");
    console.log(`Native bootstrap workspace: ${workspace}`);
    console.log(`TAST compiler: ${compiler}`);

    await runStage("node-backend", "npm", ["run", "build"], { cwd: root });
    await runStage("typed-corefn", "spago", ["build"], { cwd: workspace, env: typedEnvironment });
    stage = "verify-tast";
    const { modules, types } = verifyTypedOutput(output, compiler);
    const verification = `Verified TAST from ${compiler}: ${modules} modules, ${types} types\n`;
    writeFileSync(join(workspace, "verify-tast.log"), verification);
    console.log(verification.trimEnd());

    await runStage("generate-go", process.execPath, [join(root, "bin/gopurs.js"), "--main", "Main"], { cwd: workspace });
    await runStage("native-parser", process.execPath, [join(root, "tools/prepare-native-output.mjs"), output], { cwd: root });
    const binary = join(workspace, "gopurs-native");
    await runStage("go-build", "go", ["build", "-trimpath", "-o", binary, "./main/main.go"], { cwd: output });

    commands.checkInterrupted();
    stage = "publish binary";
    const destination = join(root, "bin/gopurs-native");
    publishNativeBinary(binary, destination);
    console.log(`Built ${destination}`);
    if (keepWorkspace) console.log(`Workspace retained: ${workspace}`);
    else rmSync(workspace, { recursive: true, force: true });
  } catch (error) {
    console.error(`Native bootstrap failed during ${stage}: ${error.message}`);
    if (workspace) console.error(`Workspace and logs retained: ${workspace}`);
    const interruption = commands.signal ? new Interrupted(commands.signal) : null;
    process.exitCode = interruption?.exitCode ?? 1;
  } finally {
    commands.dispose();
  }
}

await buildNative({ keepWorkspace: args.includes("--keep-workspace") });
