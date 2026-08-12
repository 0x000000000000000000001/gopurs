import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Effect_Console from "../Effect.Console/index.js";
import * as PureScript_Backend_Optimizer_FreeVars from "../PureScript.Backend.Optimizer.FreeVars/index.js";
var main = function __do() {
    Effect_Console.log(PureScript_Backend_Optimizer_FreeVars.sanitizeName("go"))();
    Effect_Console.log(PureScript_Backend_Optimizer_FreeVars.localId(new Data_Maybe.Just("go"))(1))();
    Effect_Console.log(PureScript_Backend_Optimizer_FreeVars.sanitizeName("go__go"))();
    return Effect_Console.log(PureScript_Backend_Optimizer_FreeVars.localId(new Data_Maybe.Just("go__go"))(1))();
};
export {
    main
};
