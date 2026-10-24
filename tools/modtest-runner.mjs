import { basename, join } from "node:path";
import { fileURLToPath } from "node:url";
import { parseOptions, selectModules, UsageError } from "./test-selection.mjs";
import { Interrupted, TestProcesses } from "./test-process.mjs";
import { prepareModule } from "./test-workspace.mjs";
import { resumeTargets, TestCampaign } from "./test-campaign.mjs";
import { compilerBuildScript } from "./test-compiler.mjs";

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
-c rebuilds the selected compiler host once, before starting the module scripts.
-d, --dev compiles library tests with the Dev profile.
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
      const buildScript = options.clean ? compilerBuildScript() : null;
      console.log(`Selected ${modules.length} modules (${options.resume ? "resume" : options.targets.length ? "explicit selection" : "all"}).`);
      if (options.dev) console.log("[gopurs] Build mode: DEV (--dev).");
      campaign = new TestCampaign(root, "modules", modules.map(directory => basename(directory)), options);
      processes = new TestProcesses();
      if (buildScript) await processes.run("build-gopurs", "npm", ["run", buildScript, "--silent"], { cwd: root, env: campaign.env(), log: join(campaign.workspace, "build-gopurs.log") });
      const success = await campaign.run(processes, async (name, directory) => {
        const cwd = prepareModule(root, directory, name, { dev: options.dev });
        await processes.run(name, "./bin/test", [], { cwd, env: campaign.env(directory), log: join(directory, "logs/test.log"), stream: true });
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
