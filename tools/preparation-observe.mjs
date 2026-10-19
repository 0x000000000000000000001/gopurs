import { spawnSync } from "node:child_process";
import { fileURLToPath } from "node:url";

export function observe(executable, args, options, trace) {
  return spawnSync("python3", ["-B", fileURLToPath(new URL("./preparation-observe.py", import.meta.url)), trace, "--", executable, ...args], options);
}
