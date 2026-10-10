import { createHash } from "node:crypto";
import { execFileSync } from "node:child_process";
import { existsSync, mkdirSync, mkdtempSync, readFileSync, renameSync, rmSync } from "node:fs";
import { dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const script = fileURLToPath(import.meta.url);
const tools = dirname(script);
const sources = join(tools, "build-cache");

// One implementation of locking, hashes, validation and publication for both
// hosts. Compilation is amortized by source content, never source timestamps.
export function resolveBuildCacheHelper() {
  if (!["darwin", "linux"].includes(process.platform)) {
    throw new Error(`Build cache locking is unsupported on ${process.platform}`);
  }
  const hash = createHash("sha256");
  for (const name of ["go.mod", "contract.go", "store.go", "main.go"]) {
    const content = readFileSync(join(sources, name));
    hash.update(`${name}:${content.length}:`).update(content);
  }
  const directory = join(tools, "../node_modules/.cache/gopurs-build-cache");
  const binary = join(directory, `${process.platform}-${process.arch}-${hash.digest("hex")}`);
  if (!existsSync(binary)) {
    mkdirSync(directory, { recursive: true });
    const staging = mkdtempSync(join(directory, ".build-"));
    try {
      execFileSync("go", ["build", "-trimpath", "-buildvcs=false", "-o", join(staging, "helper"), "."], {
        cwd: sources,
        env: {
          ...process.env, GOWORK: "off", GOFLAGS: "", CGO_ENABLED: "0",
          GOOS: process.platform, GOARCH: { x64: "amd64", arm64: "arm64" }[process.arch] ?? process.arch,
        },
        stdio: ["ignore", "pipe", "pipe"],
      });
      renameSync(join(staging, "helper"), binary);
    } finally {
      rmSync(staging, { recursive: true, force: true });
    }
  }
  return binary;
}

export function exchange(request) {
  return execFileSync(resolveBuildCacheHelper(), [], {
    input: request, encoding: "utf8", maxBuffer: 64 * 1024 * 1024,
    stdio: ["pipe", "pipe", "pipe"],
  }).trimEnd();
}

if (process.argv[1] && resolve(process.argv[1]) === script) {
  process.stdout.write(exchange(readFileSync(0, "utf8")) + "\n");
}
