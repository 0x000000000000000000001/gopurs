import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Effect_Console from "../Effect.Console/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript_Backend_Optimizer_Substitute from "../PureScript.Backend.Optimizer.Substitute/index.js";
var main = /* #__PURE__ */ (function () {
    var genericType = new PureScript_Backend_Optimizer_CoreFn.Func([ new PureScript_Backend_Optimizer_CoreFn.ADT("Data.Void.Void", [ "Data.Void.Void" ], [  ]) ], new PureScript_Backend_Optimizer_CoreFn.TypeVar("a"));
    var concrete = new PureScript_Backend_Optimizer_CoreFn.Func([ new PureScript_Backend_Optimizer_CoreFn.ADT("Data.Void.Void", [ "Data.Void.Void" ], [  ]) ], new PureScript_Backend_Optimizer_CoreFn.ADT("Data.Argonaut.Core.Json", [ "Data.Argonaut.Core.Json" ], [  ]));
    var subst = PureScript_Backend_Optimizer_Substitute.unify(genericType)(concrete)(Data_Map_Internal.empty);
    return Effect_Console.log("subst is empty? " + Data_Show.show(Data_Show.showBoolean)(Data_Map_Internal.isEmpty(subst)));
})();
export {
    main
};
