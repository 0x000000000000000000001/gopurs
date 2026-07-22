import fs from "fs";
import { decodeModule } from "../../output/PureScript.Backend.Optimizer.CoreFn.Json/index.js";
import { jsonParser } from "../../output/Data.Argonaut.Parser/index.js";
import { convertModule } from "../../output/PureScript.Backend.Optimizer.Convert/index.js";

console.log("Reading corefn.json...");
const jsonStr = fs.readFileSync("output/Main/corefn.json", "utf8");
const astRight = jsonParser(jsonStr);

if (astRight.value0) {
  const coreFnModule = decodeModule(astRight.value0);
  if (coreFnModule.value0) {
    console.time("convertModule");
    const backendModule = convertModule(coreFnModule.value0);
    console.timeEnd("convertModule");
  }
}
