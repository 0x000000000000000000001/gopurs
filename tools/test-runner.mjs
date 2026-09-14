import { existsSync, rmSync } from "node:fs";
import { basename, join } from "node:path";
import { fileURLToPath } from "node:url";
import { parseOptions, selectFixtures, UsageError } from "./test-selection.mjs";
import { Interrupted, TestProcesses } from "./test-process.mjs";
import { corePackages, createWorkspace, prepareFixture } from "./test-workspace.mjs";
import { snapshotFiles, updateSnapshots, verifySnapshots } from "./test-snapshots.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const help = `Usage: ./bin/test [fixtures...] [options]
  --list                 Print the selection without building or writing files
  --all                  Select all non-excluded fixtures (also the default)
  --skip-before NAME     Resume inclusively at a selected fixture
  -c, --clean            Rebuild gopurs; fixture workspaces are always fresh
  --update-snapshots     Create/update snapshots after successful execution
  --keep-workspace       Also keep successful runs for inspection
Legacy skip_before=NAME, --skip-before=NAME and UPDATE_SNAPSHOTS=1 still work.`;

async function runFixture(fixture, options, processes) {
  const { directory } = fixture;
  const phase = (name, command, args, cwd = directory, display = false) => processes.run(name, command, args, { cwd, display, log: join(directory, "logs", name + ".log") });
  await phase("purescript", "spago", ["build", "-q"]);
  await phase("generate-go", join(root, "bin/gopurs"), ["--main", "Main"]);
  const snapshots = snapshotFiles(root, fixture);
  for (const { generated } of snapshots) if (!existsSync(generated)) throw new Error(`Missing generated snapshot source: ${generated}`);
  if (!existsSync(join(directory, "output/main/main.go"))) throw new Error("gopurs did not generate output/main/main.go");
  await phase("format", "gofmt", ["-w", ...snapshots.map(file => file.generated)]);
  if (!options.update) await verifySnapshots(snapshots, processes, fixture);
  const output = join(directory, "output");
  // Main already writes go.mod. Keep its Go version instead of recreating it.
  await phase("go-dependencies", "go", ["mod", "tidy"], output);
  await phase("go-build", "go", ["build", "-o", "gopurs_main", "./main/main.go"], output);
  const result = await phase("execute", join(output, "gopurs_main"), [], output, true);
  if (result.includes("Fail")) throw new Error("Execution output contains Fail.");
  processes.checkInterrupted();
  if (options.update) updateSnapshots(snapshots);
}

async function main() {
  let processes;
  let workspace;
  let options;
  let passed = 0;
  let success = false;
  try {
    options = parseOptions(process.argv.slice(2));
    if (options.help) { console.log(help); return; }
    const { selected, skipped } = selectFixtures(root, options);
    if (options.list) { for (const file of selected) console.log(file); return; }
    console.log(`Selected ${selected.length} fixtures; ${skipped.length} excluded.`);
    for (const file of skipped) console.log(`=> Skipping ${basename(file)} (excluded)`);
    const packages = corePackages(root);
    workspace = createWorkspace();
    console.log(`Workspace: ${workspace}`);
    processes = new TestProcesses();
    if (options.clean) await processes.run("build-gopurs", "npm", ["run", "build", "--silent"], { cwd: root, log: join(workspace, "build-gopurs.log") });
    for (const [index, file] of selected.entries()) {
      processes.checkInterrupted();
      console.log(`=> Testing ${basename(file)}`);
      const fixture = prepareFixture(root, workspace, file, index, packages);
      await runFixture(fixture, options, processes);
      console.log("   [OK]");
      passed++;
    }
    success = true;
    console.log(`Summary: ${passed} passed, 0 failed.`);
  } catch (error) {
    console.error(`[FAILED] ${error.message}`);
    if (!(error instanceof UsageError)) console.error(`Summary: ${passed} passed, 1 failed.`);
    process.exitCode = error instanceof Interrupted ? error.exitCode : error instanceof UsageError ? 2 : 1;
  } finally {
    processes?.dispose();
    if (workspace) {
      if (success && !options.keep) rmSync(workspace, { recursive: true, force: true });
      else console.log(`Kept workspace: ${workspace}`);
    }
  }
}

await main();
