import { mkdtempSync, readFileSync, readdirSync, realpathSync, rmSync, statSync, unlinkSync } from "node:fs";
import { availableParallelism, tmpdir } from "node:os";
import { join } from "node:path";

export function fixtureWorkers(cpus = availableParallelism()) {
  return Math.max(1, Math.min(6, cpus - 1));
}

export function fixtureBatches(targets) {
  if (!targets.length || new Set(targets).size !== targets.length) throw new Error("Fixture selection must be nonempty and unique");
  const batches = [];
  for (let start = 0; start < targets.length; start += 5) {
    for (const profile of ["release", "dev"]) batches.push({ profile, start, targets: targets.slice(start, start + 5) });
  }
  return batches;
}

// Finish every active process before trimming or propagating a failure. No
// subsequent wave starts after a failed command, report check or interruption.
export async function runFixtureWaves(batches, workers, run, afterWave, beforeWave = () => {}) {
  if (!Number.isInteger(workers) || workers < 1 || workers > 6) throw new Error("Fixture workers must be between 1 and 6");
  for (let index = 0; index < batches.length; index += workers) {
    const wave = batches.slice(index, index + workers);
    await beforeWave(wave);
    const results = await Promise.allSettled(wave.map(batch => Promise.resolve().then(() => run(batch))));
    await afterWave();
    const failed = results.find(result => result.status === "rejected");
    if (failed) throw failed.reason;
  }
}

export function readFixtureReport(log, root, batch) {
  const matches = [...readFileSync(log, "utf8").matchAll(/^Results: (.+)$/gm)];
  if (matches.length !== 1) throw new Error(`Missing or ambiguous fixture report in ${log}`);
  const file = matches[0][1];
  const report = JSON.parse(readFileSync(file, "utf8"));
  if (report?.version !== 1 || report.kind !== "fixtures" || report.root !== realpathSync(root)
      || report.buildProfile !== batch.profile || report.updateSnapshots !== false
      || !Array.isArray(report.skipped) || report.skipped.length || !report.finishedAt
      || !Array.isArray(report.targets) || report.targets.length !== batch.targets.length
      || report.targets.some((row, index) => row?.target !== batch.targets[index] || row.status !== "passed")) {
    throw new Error(`Incomplete or mismatched fixture report: ${file}`);
  }
  return { file, profile: batch.profile, passed: report.targets.length };
}

// Only this newly created cache is ever pruned. In particular, neither the
// caller's GOCACHE nor symlink targets inside this directory are traversed.
export class QualityGoCache {
  constructor({ parent = tmpdir(), limitBytes = 2 * 1024 ** 3 } = {}) {
    if (!Number.isSafeInteger(limitBytes) || limitBytes < 0) throw new Error("Invalid Go cache budget");
    this.path = mkdtempSync(join(parent, "gopurs-quality-gocache-"));
    this.limitBytes = limitBytes;
    this.active = 0;
  }

  environment(env) {
    const result = { ...env, GOCACHE: this.path };
    delete result.GOCACHEPROG;
    return result;
  }

  async withProcess(run) {
    this.active++;
    try { return await run(); }
    finally { this.active--; }
  }

  requireIdle() {
    if (this.active) throw new Error("Cannot clean Go cache while commands are active");
  }

  trim() {
    this.requireIdle();
    const files = [];
    const visit = directory => {
      for (const entry of readdirSync(directory, { withFileTypes: true })) {
        const path = join(directory, entry.name);
        if (entry.isDirectory()) visit(path);
        else if (entry.isFile()) {
          const stat = statSync(path);
          files.push({ path, bytes: stat.size, modified: stat.mtimeMs });
        }
      }
    };
    visit(this.path);
    const beforeBytes = files.reduce((sum, file) => sum + file.bytes, 0);
    let afterBytes = beforeBytes, deletedObjects = 0;
    for (const file of files.sort((a, b) => a.modified - b.modified || a.path.localeCompare(b.path))) {
      if (afterBytes <= this.limitBytes) break;
      unlinkSync(file.path);
      afterBytes -= file.bytes;
      deletedObjects++;
    }
    return { beforeBytes, afterBytes, deletedObjects };
  }

  dispose() {
    this.requireIdle();
    rmSync(this.path, { recursive: true, force: true });
  }
}
