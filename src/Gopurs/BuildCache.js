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
