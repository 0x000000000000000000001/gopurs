import { basename } from "node:path";
import { fileURLToPath } from "node:url";
import { parseOptions, selectModules, UsageError } from "./test-selection.mjs";
import { Interrupted, TestProcesses } from "./test-process.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
let processes;
try {
  const options = parseOptions(process.argv.slice(2), { modules: true });
  if (options.help) {
    console.log(`Usage: ./bin/modtest [modules...] [--all] [--skip-before NAME] [--list] [-c]
With no module names, run all sibling gopurs-* repositories with an executable bin/test.
Resume is inclusive; names accept either prelude or gopurs-prelude.
-c rebuilds the compiler once, before starting the selected module scripts.
Each sibling script still controls its own build, caches, and cleanup.`);
  } else {
    const modules = selectModules(root, options);
    for (const directory of modules) console.log(basename(directory));
    if (!options.list) {
      console.log(`Selected ${modules.length} modules (${options.resume ? "resume" : options.targets.length ? "explicit selection" : "all"}).`);
      processes = new TestProcesses();
      if (options.clean) await processes.run("build-gopurs", "npm", ["run", "build", "--silent"], { cwd: root });
      for (const directory of modules) await processes.run(basename(directory), "./bin/test", [], { cwd: directory });
      console.log(`Summary: ${modules.length} modules passed.`);
    }
  }
} catch (error) {
  console.error(`[FAILED] ${error.message}`);
  process.exitCode = error instanceof Interrupted ? error.exitCode : error instanceof UsageError ? 2 : 1;
} finally {
  processes?.dispose();
}
