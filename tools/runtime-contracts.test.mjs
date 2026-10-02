import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { cpSync, mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { test } from "node:test";

test("canonical runtime preserves value storage, immutable updates and retained work", async t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-runtime-contracts-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  cpSync(new URL("../runtime/", import.meta.url), join(directory, "gopurs_runtime"), { recursive: true });
  writeFileSync(join(directory, "go.mod"), "module gopurs/output\n\ngo 1.22\n");
  for (const race of [false, true]) {
    await t.test(race ? "race detector" : "optimized build", () => {
      const result = spawnSync("go", ["test", "-json", "-count=1", "-timeout=45s",
        ...(race ? ["-race"] : []), "./gopurs_runtime"], {
        cwd: directory, encoding: "utf8", timeout: 120_000,
        env: { ...process.env, GOWORK: "off" },
      });
      assert.ifError(result.error);
      assert.equal(result.status, 0, result.stdout + result.stderr);
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
