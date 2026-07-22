import fs from "fs";
import { decodeModule } from "../../output/PureScript.Backend.Optimizer.CoreFn.Json/index.js";
import { jsonParser } from "../../output/Data.Argonaut.Parser/index.js";
import { toBackendModule } from "../../output/PureScript.Backend.Optimizer.Convert/index.js";
import * as Map from "../../output/Data.Map.Internal/index.js";
import * as Set from "../../output/Data.Set/index.js";
import { List } from "../../output/Data.List.Types/index.js";

console.log("Reading corefn.json...");
const jsonStr = fs.readFileSync("output/Main/corefn.json", "utf8");
console.log("Parsing...");
const astRight = jsonParser(jsonStr);

if (astRight.value0) {
  console.log("Decoding...");
  const coreFnModule = decodeModule(astRight.value0);
  if (coreFnModule.value0) {
    console.log("toBackendModule...");
    console.time("toBackendModule");
    const result = toBackendModule(coreFnModule.value0)({
      analyzeCustom: function(a) { return function(b) { return { value0: null, value1: null }; }; },
      currentModule: coreFnModule.value0.value1.name,
      currentLevel: 0,
      toLevel: Map.empty,
      implementations: Map.empty,
      moduleImplementations: Map.empty,
      directives: Map.empty,
      dataTypes: Map.empty,
      foreignSemantics: Map.empty,
      rewriteLimit: 10000,
      traceIdents: Set.empty,
      optimizationSteps: []
    });
    console.timeEnd("toBackendModule");
  }
}
