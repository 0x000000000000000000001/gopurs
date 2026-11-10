import { existsSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

// Resolve the gopurs checkout from the calling module. Tools can live in
// tools/, tools/build/ or tools/release/ without duplicating ../ counts.
export function repoRoot(moduleUrl) {
  let directory = dirname(fileURLToPath(moduleUrl));
  for (;;) {
    if (existsSync(join(directory, "spago.yaml"))) return directory;
    const parent = dirname(directory);
    if (parent === directory) {
      throw new Error(`Cannot find the gopurs checkout from ${moduleUrl}`);
    }
    directory = parent;
  }
}
