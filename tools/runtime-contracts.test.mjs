import assert from "node:assert/strict";
import { cpSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import { test } from "node:test";
import { runGo, temporaryDirectory } from "./lib/go.mjs";

test("canonical runtime preserves value storage, immutable updates and retained work", async t => {
  const directory = temporaryDirectory(t, "gopurs-runtime-contracts-");
  cpSync(new URL("../runtime/", import.meta.url), join(directory, "gopurs_runtime"), { recursive: true });
  writeFileSync(join(directory, "go.mod"), "module gopurs/output\n\ngo 1.22\n");
  for (const race of [false, true]) {
    await t.test(race ? "race detector" : "optimized build", () => {
      const result = runGo(directory, ["test", "-json", "-count=1", "-timeout=45s",
        ...(race ? ["-race"] : []), "./gopurs_runtime"]);
      const passed = new Set(result.stdout.trim().split("\n").map(line => JSON.parse(line))
        .filter(event => event.Action === "pass").map(event => event.Test));
      for (const name of ["PackedStringsRetainStorage", "ContainersRetainPayloadsAfterCreatorExit",
        "ContainerConstructorsShareCallerStorage", "EventLoopWaitIncludesRetainedWork",
        "JSONObjectOwnedMaterialization", "JSONObjectViewsAcceptExistingRepresentations",
        "RecordSetCompactPromotionAndReplacement", "RecordSetGeneralRepresentationAndMapAliases",
        "RecordSetRetainsClosureAcrossGC", "ConcurrentClosuresSurviveGC"]) {
        assert.ok(passed.has(`Test${name}`), `${name} did not pass`);
      }
    });
  }
});
