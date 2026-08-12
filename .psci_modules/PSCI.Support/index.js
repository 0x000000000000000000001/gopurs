// | This module provides support for the
// | PureScript interactive mode, PSCI.
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Effect from "../Effect/index.js";
import * as Effect_Console from "../Effect.Console/index.js";
var evalShow = function (dictShow) {
    return {
        "eval": Effect_Console.logShow(dictShow)
    };
};
var evalEffectUnit = {
    "eval": /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn)
};
var $$eval = function (dict) {
    return dict["eval"];
};
var evalEffect = function (dictEval) {
    var eval1 = $$eval(dictEval);
    return {
        "eval": function (x) {
            return Control_Bind.bind(Effect.bindEffect)(x)(eval1);
        }
    };
};
export {
    $$eval as eval,
    evalEffectUnit,
    evalEffect,
    evalShow
};
