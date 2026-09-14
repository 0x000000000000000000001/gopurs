import { copyFileSync, existsSync, mkdirSync, readFileSync } from "node:fs";
import { dirname, join } from "node:path";

export function snapshotFiles(root, fixture) {
  return (fixture.snapshotFfi ? ["", "_ffi"] : [""]).map(suffix => ({
    generated: join(fixture.directory, `output/purescript/Main${suffix}.go`),
    expected: join(root, `tests/passing-snapshots/${fixture.name}${suffix}.go`),
  }));
}

export async function verifySnapshots(files, processes, fixture) {
  for (const { generated, expected } of files) {
    if (!existsSync(expected)) throw new Error(`Missing snapshot: ${expected}. Use --update-snapshots to create it explicitly.`);
    if (!readFileSync(expected).equals(readFileSync(generated))) {
      await processes.run("snapshot diff", "diff", ["-u", expected, generated], { cwd: fixture.directory, log: join(fixture.directory, "logs/snapshot.diff"), display: true });
    }
  }
}

export function updateSnapshots(files) {
  // Called only after the fixture's Go build and execution have succeeded.
  for (const { generated, expected } of files) {
    mkdirSync(dirname(expected), { recursive: true });
    copyFileSync(generated, expected);
    console.log(`   Updated snapshot: ${expected}`);
  }
}
