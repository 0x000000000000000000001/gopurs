import { accessSync, constants, realpathSync } from "node:fs";
import { join, resolve } from "node:path";

// PATH normally contains the TAST frontend in the development/release shell.
// The JavaScript oracle must use the independently installed stock frontend.
export function findOfficialCompiler(root, configured = process.env.GOPURS_OFFICIAL_PURS) {
  const candidate = configured ? resolve(configured) : join(root, "node_modules/purescript/purs.bin");
  try {
    accessSync(candidate, constants.X_OK);
    return realpathSync(candidate);
  } catch (cause) {
    throw new Error(
      `No executable standard PureScript compiler at ${candidate}. Set GOPURS_OFFICIAL_PURS (or --official-purs) to the stock compiler, use the Nix development shell, or install the npm development dependencies with their install scripts. PATH is not used for the JavaScript oracle.`,
      { cause },
    );
  }
}
