import assert from "node:assert/strict";
import { readFileSync, readdirSync } from "node:fs";
import { join } from "node:path";

function readSources(directory, prefix = "") {
  const result = {};
  for (const entry of readdirSync(directory, { withFileTypes: true })) {
    const name = prefix + entry.name, path = join(directory, entry.name);
    if (entry.isDirectory()) Object.assign(result, readSources(path, name + "/"));
    else result[name] = readFileSync(path, "utf8");
  }
  return result;
}

export function preparationSourceVariants(directory) {
  const original = readSources(directory);
  let sources = { ...original }, values = [5, 6, 15];
  const states = [];
  const edit = (path, before, after) => {
    assert.equal(sources[path].split(before).length, 2, "Unique preparation mutation: " + path);
    sources[path] = sources[path].replace(before, after);
  };
  const add = (label, observations = {}) => states.push({ label, sources: { ...sources }, stdout: values.join("\n") + "\nDone\n", ...observations });
  add("initial");
  edit("Preparation/Calls.purs", "run value = Cycle.forward 2 value", "run value = Cycle.forward 2 value + if Cycle.forward 1 (value > 0) then 7 else 8");
  values[0] = 12;
  add("cycle-demand-added", { sameTast: ["Preparation.Cycle"], changedPrepared: ["Preparation.Cycle"] });
  edit("Preparation/Calls.purs", "Cycle.forward 1 (value > 0)", 'Cycle.forward 1 (if value > 0 then "yes" else "no") == "yes"');
  add("cycle-demand-replaced", { sameTast: ["Preparation.Cycle"], changedPrepared: ["Preparation.Cycle"], sameCycleCardinalityAs: "cycle-demand-added" });
  sources["Preparation/Calls.purs"] = original["Preparation/Calls.purs"];
  values[0] = 5;
  add("cycle-demand-retracted", { sameTast: ["Preparation.Cycle"], changedPrepared: ["Preparation.Cycle"], restoredPrepared: ["Preparation.Cycle"] });
  edit("Preparation/Provider.purs", "method = { apply: first }", "method = { apply: second }");
  values[1] = 7;
  add("dictionary-provider-changed", { sameTast: ["Preparation.Apply", "Preparation.Calls"], changedPrepared: ["Preparation.Apply"] });
  edit("Main.go", "return value + 10", "return value + 11");
  values[2] = 16;
  add("ffi-only", { sameAllTast: true });
  sources = { ...original }; values = [5, 6, 15];
  add("all-restored", { restoredAll: true });
  return states;
}
