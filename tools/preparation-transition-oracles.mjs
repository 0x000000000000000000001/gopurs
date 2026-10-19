import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import { monadIdentity } from "../output/Data.Identity/index.js";
import { Cons, Nil } from "../output/Data.List.Types/index.js";
import { Left } from "../output/Data.Either/index.js";
import * as Map from "../output/Data.Map/index.js";
import { unfoldableArray } from "../output/Data.Unfoldable/index.js";
import * as Aff from "../output/Effect.Aff/index.js";
import { coreFnModulesFromOutput } from "../output/PureScript.Backend.Optimizer.App/index.js";
import { transitiveCollect } from "../output/PureScript.Backend.Optimizer.Monomorphize/index.js";
import { addClassDataDeclarations } from "../output/Gopurs.ClassMetadata/index.js";
import { buildGlobalTypes } from "../output/Gopurs.GlobalTypes/index.js";
import { monomorphizeModulesWith } from "../output/Gopurs.Monomorphization/index.js";

const entries = map => Map.toUnfoldable(unfoldableArray)(map).map(pair => [pair.value0, pair.value1]);
const list = values => values.reduceRight((tail, value) => new Cons(value, tail), Nil.value);
function array(value) {
  const result = [];
  while (value instanceof Cons) { result.push(value.value0); value = value.value1; }
  return result;
}
const hash = value => createHash("sha256").update(JSON.stringify(value)).digest("hex");
const runAff = action => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Left ? reject(result.value0) : resolve(result.value0))(action)());

export async function preparationTransitionOracles(states) {
  const observations = {};
  for (const state of states) {
    const input = await runAff(coreFnModulesFromOutput(state.snapshot)), modules = array(input);
    let observed;
    const collector = ast => direct => {
      const transitive = transitiveCollect(ast)(direct);
      const cycle = Object.fromEntries(entries(transitive).filter(([name]) => name === "Preparation.Cycle.forward" || name === "Preparation.Cycle.backward")
        .map(([name, values]) => [name, entries(values).map(([key]) => key)]));
      assert.equal(Object.keys(cycle).length, 2, "Both sides of the cycle participate");
      const apply = entries(direct).find(([name]) => name === "Preparation.Apply.applyMethod");
      assert(apply, "Static dictionary call must be collected");
      observed = { cycle, directApply: hash(entries(apply[1])) };
      return transitive;
    };
    const prepared = array(monomorphizeModulesWith(monadIdentity)(collector)(buildGlobalTypes(modules))(list(modules.map(addClassDataDeclarations))));
    observed.preparedApply = hash(prepared.find(mod => mod.name === "Preparation.Apply"));
    if (state.sameCycleCardinalityAs) {
      const previous = observations[state.sameCycleCardinalityAs].cycle;
      for (const [name, keys] of Object.entries(observed.cycle)) {
        assert.equal(keys.length, previous[name].length, name + ": replacement preserves cardinality");
        assert.notDeepEqual(keys, previous[name], name + ": replacement changes content");
      }
    }
    if (state.label === "cycle-demand-retracted") assert.deepEqual(observed.cycle, observations.initial.cycle, "The cycle cannot retain support from the removed demand");
    if (state.label === "dictionary-provider-changed") {
      const previous = observations["cycle-demand-retracted"];
      assert.equal(observed.directApply, previous.directApply, "Provider body changes with identical direct request payloads");
      assert.notEqual(observed.preparedApply, previous.preparedApply, "Dictionary body reads affect specialization");
    }
    if (state.sameAllTast) assert.deepEqual(observed, observations[states[states.indexOf(state) - 1].label]);
    if (state.restoredAll) assert.deepEqual(observed, observations.initial);
    observations[state.label] = observed;
  }
  return observations;
}
