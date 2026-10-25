import { execFileSync } from "node:child_process";
import { mkdirSync, realpathSync, writeFileSync } from "node:fs";
import { join } from "node:path";

export const devGoFlag = "-gcflags=gopurs/output/...=-N -l";
export const devGoMessage = `[gopurs] Go build mode: DEV (${devGoFlag}).`;

// Library scripts own their Go commands. Intercept only application commands
// inside this library's private copy, leaving the compiler host and Go helpers
// outside it alone. Do not export GOFLAGS to the backend or host rebuild.
export function moduleGoEnvironment(directory, library, env) {
  const go = execFileSync("bash", ["-c", "command -v go"], { env, encoding: "utf8" }).trim();
  const quote = text => "'" + text.replace(/'/g, "'\\''") + "'";
  const bin = join(directory, "go-bin");
  mkdirSync(bin);
  writeFileSync(join(bin, "go"), `#!/usr/bin/env bash
case "$(pwd -P)" in
  ${quote(realpathSync(library))}|${quote(realpathSync(library))}/*)
    case "$1" in
      build|run|test)
        command="$1"
        shift
        echo ${quote(devGoMessage)} >&2
        exec ${quote(go)} "$command" ${quote(devGoFlag)} "$@"
        ;;
    esac
    ;;
esac
exec ${quote(go)} "$@"
`, { mode: 0o755 });
  return { ...env, PATH: bin + ":" + env.PATH };
}
