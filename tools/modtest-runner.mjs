import { basename, join } from "node:path";
import { fileURLToPath } from "node:url";
import { parseOptions, selectModules, UsageError } from "./test-selection.mjs";
import { Interrupted, TestProcesses } from "./test-process.mjs";
import { prepareModule } from "./test-workspace.mjs";
import { resumeTargets, TestCampaign } from "./test-campaign.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
async function main() {
  let processes;
  let campaign;
  let failure;
  try {
    const options = parseOptions(process.argv.slice(2), { modules: true });
    if (options.help) {
      console.log(`Usage: ./bin/modtest [modules...] [options]
With no module names, run all sibling gopurs-* repositories with an executable bin/test.
Resume is inclusive; names accept either prelude or gopurs-prelude.
--all, --list, --skip-before NAME retain their selection behavior.
--keep-going attempts all selected libraries after individual failures.
--resume-failed FILE selects unsuccessful targets from a JSON campaign report.
--keep-workspace also keeps successful library copies and generated files.
-c rebuilds the compiler once, before starting the selected module scripts.
Each library runs in a fresh copy of the sibling sources, with private temporaries.`);
      return;
    }
    const retry = resumeTargets(root, options, "modules");
    if (retry?.length === 0) {
      if (!options.list) console.log("No unsuccessful targets to resume.");
      return;
    }
    if (retry) options.targets = retry;
    const modules = selectModules(root, options);
    for (const directory of modules) console.log(basename(directory));
    if (!options.list) {
      console.log(`Selected ${modules.length} modules (${options.resume ? "resume" : options.targets.length ? "explicit selection" : "all"}).`);
      campaign = new TestCampaign(root, "modules", modules.map(directory => basename(directory)), options);
      processes = new TestProcesses();
      if (options.clean) await processes.run("build-gopurs", "npm", ["run", "build", "--silent"], { cwd: root, env: campaign.env(), log: join(campaign.workspace, "build-gopurs.log") });
      const success = await campaign.run(processes, async (name, directory) => {
        const cwd = prepareModule(root, directory, name);
        await processes.run(name, "./bin/test", [], { cwd, env: campaign.env(directory), log: join(directory, "logs/test.log"), display: true });
      });
      if (!success) process.exitCode = 1;
    }
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
