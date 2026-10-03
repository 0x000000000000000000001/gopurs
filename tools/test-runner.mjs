import { existsSync } from "node:fs";
import { basename, join } from "node:path";
import { fileURLToPath } from "node:url";
import { parseOptions, selectFixtures, UsageError } from "./test-selection.mjs";
import { Interrupted, TestProcesses } from "./test-process.mjs";
import { corePackages, prepareFixture } from "./test-workspace.mjs";
import { snapshotFiles, updateSnapshots, verifySnapshots } from "./test-snapshots.mjs";
import { resumeTargets, TestCampaign } from "./test-campaign.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const help = `Usage: ./bin/test [fixtures...] [options]
  --list                 Print the selection without building or writing files
  --all                  Select all non-excluded fixtures (also the default)
  --skip-before NAME     Resume inclusively at a selected fixture
  --resume-failed FILE   Select failed/interrupted/pending targets from a JSON report
  --keep-going           Attempt every selected fixture after individual failures
  -c, --clean            Rebuild gopurs; fixture workspaces are always fresh
  --update-snapshots     Create/update snapshots after successful execution
  --keep-workspace       Also keep successful runs for inspection
Legacy skip_before=NAME, --skip-before=NAME and UPDATE_SNAPSHOTS=1 still work.`;

async function runFixture(fixture, options, processes, env) {
  const { directory } = fixture;
  const phase = (name, command, args, cwd = directory, display = false) => processes.run(name, command, args, { cwd, env, display, log: join(directory, "logs", name + ".log") });
  await phase("purescript", "spago", ["build", "-q"]);
  await phase("generate-go", join(root, "bin/gopurs"), ["--main", "Main"]);
  const snapshots = snapshotFiles(root, fixture);
  for (const { generated } of snapshots) if (!existsSync(generated)) throw new Error(`Missing generated snapshot source: ${generated}`);
  if (!existsSync(join(directory, "output/main/main.go"))) throw new Error("gopurs did not generate output/main/main.go");
  await phase("format", "gofmt", ["-w", ...snapshots.map(file => file.generated)]);
  if (!options.update) await verifySnapshots(snapshots, processes, fixture, env);
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
  let campaign;
  let failure;
  try {
    const options = parseOptions(process.argv.slice(2));
    if (options.help) { console.log(help); return; }
    const retry = resumeTargets(root, options, "fixtures");
    if (retry?.length === 0) {
      if (!options.list) console.log("No unsuccessful targets to resume.");
      return;
    }
    if (retry) options.targets = retry;
    const { selected, skipped } = selectFixtures(root, options);
    if (options.list) { for (const file of selected) console.log(file); return; }
    // Fixture compilations are tiny: use the sequential PBO builder so the
    // harness stays deterministic and independent from machine sizing rules.
    // Sequential and parallel output is byte-identical (validated on b8x and
    // on a fixture sample run with GOPURS_PBO_JOBS=8).
    process.env.GOPURS_PBO_JOBS ??= "1";
    process.env.GOPURS_PREPARE_JOBS ??= "1";
    console.log(`Selected ${selected.length} fixtures; ${skipped.length} excluded.`);
    for (const file of skipped) console.log(`=> Skipping ${basename(file)} (excluded)`);
    const packages = corePackages(root);
    campaign = new TestCampaign(root, "fixtures", selected, options, skipped);
    processes = new TestProcesses();
    if (options.clean) await processes.run("build-gopurs", "npm", ["run", "build", "--silent"], { cwd: root, env: campaign.env(), log: join(campaign.workspace, "build-gopurs.log") });
    const success = await campaign.run(processes, async (file, directory, index) => {
      console.log(`=> Testing ${basename(file)}`);
      const fixture = prepareFixture(root, campaign.workspace, file, index, packages);
      await runFixture(fixture, options, processes, campaign.env(directory));
      console.log("   [OK]");
    });
    if (!success) process.exitCode = 1;
  } catch (error) {
    failure = error;
    console.error(`[FAILED] ${error.message}`);
    process.exitCode = error instanceof Interrupted ? error.exitCode : error instanceof UsageError ? 2 : 1;
  } finally {
    processes?.dispose();
    campaign?.finish(failure);
  }
}

await main();
