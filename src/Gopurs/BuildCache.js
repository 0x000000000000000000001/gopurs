import { execFileSync } from "node:child_process";
import { existsSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

export const exchange = request => () => {
  const directory = dirname(fileURLToPath(import.meta.url));
  const runner = [
    join(directory, "../tools/build-cache.mjs"),
    join(directory, "../../tools/build-cache.mjs"),
  ].find(existsSync);
  if (!runner) throw new Error(`Cannot locate the build cache runner from ${directory}`);
  return execFileSync(process.execPath, [runner], {
    input: request, encoding: "utf8", maxBuffer: 64 * 1024 * 1024,
    stdio: ["pipe", "pipe", "pipe"],
  }).trimEnd();
};

export const sessionExchange = request => () => {
  const root = process.env.GOPURS_COMPILER_SNAPSHOT;
  if (!root) throw new Error("Compiler snapshot is required");
  const message = JSON.parse(request);
  if (message.operation === "begin") message.options.push({ name: "js-runtime", value: JSON.stringify(process.versions) });
  return execFileSync(join(root, "tools/build-cache-helper"), ["--session"], {
    input: JSON.stringify(message), encoding: "utf8", maxBuffer: 64 * 1024 * 1024,
    stdio: ["pipe", "pipe", "pipe"],
  }).trimEnd();
};

// The helper process cannot call back into this JS runtime. Keep the ordinary
// sequential path; the caller must accept that preload notification is absent.
export const sessionExchangeOnMiss = request => _onMiss => sessionExchange(request);
