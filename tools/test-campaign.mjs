import { mkdirSync, readFileSync, readdirSync, realpathSync, renameSync, rmSync, writeFileSync } from "node:fs";
import { basename, join, resolve } from "node:path";
import { Interrupted } from "./test-process.mjs";
import { UsageError } from "./test-selection.mjs";
import { createWorkspace } from "./test-workspace.mjs";

// Resume selects work for a new run. Previous successes remain evidence of the
// previous run, not cached successes claimed against possibly changed sources.
export function resumeTargets(root, options, kind) {
  if (!options.resumeFailed) return null;
  let report;
  try { report = JSON.parse(readFileSync(resolve(options.resumeFailed), "utf8")); }
  catch (error) { throw new UsageError(`Cannot read campaign report: ${error.message}`); }
  const statuses = new Set(["pending", "running", "passed", "failed", "interrupted"]);
  if (report?.version !== 1 || report.kind !== kind || report.root !== realpathSync(root)
      || !Array.isArray(report.targets) || !report.targets.every(row =>
        row && typeof row.target === "string" && row.target.length && statuses.has(row.status))
      || new Set(report.targets.map(row => row.target)).size !== report.targets.length) {
    throw new UsageError("Invalid campaign report, kind, or repository root.");
  }
  return report.targets.filter(row => row.status !== "passed").map(row => row.target);
}

export class TestCampaign {
  constructor(root, kind, targets, options, skipped = []) {
    this.options = options;
    this.workspace = createWorkspace();
    this.file = join(this.workspace, "results.json");
    this.report = {
      version: 1, kind, root: realpathSync(root), startedAt: new Date().toISOString(),
      resumedFrom: options.resumeFailed ? resolve(options.resumeFailed) : null,
      updateSnapshots: options.update, skipped,
      targets: targets.map((target, index) => ({
        target, status: "pending",
        directory: join(this.workspace, `${index}-${basename(target, kind === "fixtures" ? ".purs" : undefined)}`),
      })),
    };
    mkdirSync(join(this.workspace, "tmp"));
    this.save();
    console.log(`Workspace: ${this.workspace}`);
    console.log(`Results: ${this.file}`);
  }

  env(directory = this.workspace) {
    const temporary = join(directory, "tmp");
    return { ...process.env, TMPDIR: temporary, TMP: temporary, TEMP: temporary };
  }

  save() {
    // A failed write must not truncate the last complete checkpoint.
    const temporary = this.file + ".tmp";
    writeFileSync(temporary, JSON.stringify(this.report, null, 2) + "\n");
    renameSync(temporary, this.file);
  }

  async run(processes, action) {
    for (const [index, row] of this.report.targets.entries()) {
      processes.checkInterrupted();
      row.status = "running";
      row.startedAt = new Date().toISOString();
      this.save();
      let interruption;
      try {
        mkdirSync(join(row.directory, "logs"), { recursive: true });
        mkdirSync(join(row.directory, "tmp"));
        await action(row.target, row.directory, index);
        processes.checkInterrupted();
        row.status = "passed";
      } catch (error) {
        row.status = error instanceof Interrupted ? "interrupted" : "failed";
        row.error = error.message;
        console.error(`[FAILED] ${row.target}: ${error.message}`);
        if (error instanceof Interrupted) interruption = error;
      }
      row.finishedAt = new Date().toISOString();
      this.save();
      if (row.status === "passed" && !this.options.keep) {
        // Keep logs/report even on success, but bound generated data on disk.
        for (const name of readdirSync(row.directory)) {
          if (name !== "logs") rmSync(join(row.directory, name), { recursive: true, force: true });
        }
      }
      if (interruption) throw interruption;
      if (row.status !== "passed" && !this.options.keepGoing) break;
    }
    return this.report.targets.every(row => row.status === "passed");
  }

  finish(error) {
    if (error) this.report.error = error.message;
    this.report.finishedAt = new Date().toISOString();
    this.save();
    const count = status => this.report.targets.filter(row => row.status === status).length;
    const passed = count("passed");
    const failed = count("failed") + count("interrupted");
    const pending = count("pending") + count("running");
    console.log(`Summary: ${passed} passed, ${failed} failed, ${pending} pending.`);
    if (error || failed || pending || this.options.keep) console.log(`Kept workspace: ${this.workspace}`);
    else {
      rmSync(join(this.workspace, "tmp"), { recursive: true, force: true });
      console.log(`Kept reports and logs: ${this.workspace}`);
    }
  }
}
