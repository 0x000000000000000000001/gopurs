import { UsageError } from "./test-selection.mjs";

// Match bin/gopurs so a clean campaign rebuilds the executable it will run.
export function compilerBuildScript(env = process.env) {
  if (env.GOPURS_JS === "1" && env.GOPURS_RUST === "1") {
    throw new UsageError("Select only one compiler host: GOPURS_RUST=1 or GOPURS_JS=1.");
  }
  if (env.GOPURS_RUST === "1") return "build:rust";
  return env.GOPURS_JS === "1" ? "build" : "build:native";
}
