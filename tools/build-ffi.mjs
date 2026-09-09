import { execFileSync } from "node:child_process";
import {
  copyFileSync,
  mkdtempSync,
  readFileSync,
  renameSync,
  rmSync,
} from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

const toolsDirectory = dirname(fileURLToPath(import.meta.url));
const generatorDirectory = join(toolsDirectory, "ffi-gen");

function buildFfi() {
  const goModPath = join(generatorDirectory, "go.mod");
  const goMod = readFileSync(goModPath, "utf8");
  const requiredVersion = goMod.match(
    /^go[ \t]+(\d+\.\d+\.\d+)[ \t]*(?:\/\/.*)?$/m,
  )?.[1];
  if (!requiredVersion) {
    throw new Error(`${goModPath} must pin a Go version including its patch number.`);
  }

  const goEnvironment = {
    ...process.env,
    GOTOOLCHAIN: "local",
    GOWORK: "off",
  };
  const [actualVersion, goRoot] = execFileSync("go", ["env", "GOVERSION", "GOROOT"], {
    cwd: generatorDirectory,
    env: goEnvironment,
    encoding: "utf8",
    stdio: ["ignore", "pipe", "inherit"],
  }).trim().split(/\r?\n/);

  if (actualVersion !== `go${requiredVersion}`) {
    throw new Error(
      `Go ${requiredVersion} is required by ${goModPath}; found ${actualVersion}. ` +
      `Put Go ${requiredVersion} on PATH and run this command again. ` +
      "Automatic toolchain switching is disabled (GOTOOLCHAIN=local).",
    );
  }
  if (!goRoot) {
    throw new Error("go env GOROOT returned no directory for the pinned toolchain.");
  }

  // Stage both files on the destination filesystem before replacing either one.
  const temporaryDirectory = mkdtempSync(join(toolsDirectory, ".build-ffi-"));
  try {
    const wasmPath = join(temporaryDirectory, "ffi_gen.wasm");
    const runtimePath = join(temporaryDirectory, "wasm_exec.js");
    copyFileSync(join(goRoot, "lib", "wasm", "wasm_exec.js"), runtimePath);
    execFileSync("go", [
      "build",
      "-trimpath",
      "-buildvcs=false",
      "-ldflags=-buildid=",
      "-o", wasmPath,
      ".",
    ], {
      cwd: generatorDirectory,
      env: { ...goEnvironment, GOOS: "js", GOARCH: "wasm", CGO_ENABLED: "0" },
      stdio: "inherit",
    });
    renameSync(wasmPath, join(toolsDirectory, "ffi_gen.wasm"));
    renameSync(runtimePath, join(toolsDirectory, "wasm_exec.js"));
  } finally {
    rmSync(temporaryDirectory, { recursive: true, force: true });
  }

  console.log(`Built tools/ffi_gen.wasm and tools/wasm_exec.js with Go ${requiredVersion}.`);
}

try {
  buildFfi();
} catch (error) {
  console.error(`FFI build failed: ${error.message}`);
  process.exitCode = 1;
}
