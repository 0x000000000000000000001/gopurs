import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Data_FoldableWithIndex from "../Data.FoldableWithIndex/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_List from "../Data.List/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as PureScript_Backend_Optimizer_Convert from "../PureScript.Backend.Optimizer.Convert/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript_Backend_Optimizer_Semantics from "../PureScript.Backend.Optimizer.Semantics/index.js";
var insert = /* #__PURE__ */ Data_Map_Internal.insert(/* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.ordQualified(PureScript_Backend_Optimizer_CoreFn.ordIdent));
var insert1 = /* #__PURE__ */ Data_Map_Internal.insert(PureScript_Backend_Optimizer_Semantics.ordEvalRef);

// | Builds modules given a _sorted_ list of modules.
// | See `PureScript.Backend.Optimizer.CoreFn.Sort.sortModules`.
var buildModules = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var Applicative0 = dictMonad.Applicative0();
    var $$void = Data_Functor["void"](((dictMonad.Bind1()).Apply0()).Functor0());
    return function (options) {
        return function (coreFnModules) {
            var moduleCount = Data_List.length(coreFnModules);
            var go = function (v) {
                return function (coreFnModule) {
                    var buildEnv = {
                        implementations: v.implementations,
                        moduleCount: moduleCount,
                        moduleIndex: v.moduleIndex
                    };
                    return Control_Bind.bind(Bind1)(options.onPrepareModule(buildEnv)(coreFnModule))(function (v1) {
                        return Control_Bind.bind(Bind1)(options.onSkipModule(buildEnv)(v1))(function (mbCachedMod) {
                            if (mbCachedMod instanceof Data_Maybe.Just) {
                                var newImplementations = Data_FoldableWithIndex.foldrWithIndex(Data_Map_Internal.foldableWithIndexMap)(insert)(v.implementations)(mbCachedMod.value0.implementations);
                                return Control_Applicative.pure(Applicative0)({
                                    directives: Data_FoldableWithIndex.foldrWithIndex(Data_Map_Internal.foldableWithIndexMap)(insert1)(v.directives)(mbCachedMod.value0.directives),
                                    implementations: newImplementations,
                                    moduleIndex: v.moduleIndex + 1 | 0
                                });
                            };
                            if (mbCachedMod instanceof Data_Maybe.Nothing) {
                                var v2 = PureScript_Backend_Optimizer_Convert.toBackendModule(v1)({
                                    analyzeCustom: options.analyzeCustom,
                                    currentModule: v1.name,
                                    currentLevel: 0,
                                    toLevel: Data_Map_Internal.empty,
                                    implementations: v.implementations,
                                    moduleImplementations: Data_Map_Internal.empty,
                                    directives: v.directives,
                                    dataTypes: Data_Map_Internal.empty,
                                    foreignSemantics: options.foreignSemantics,
                                    rewriteLimit: 10000,
                                    traceIdents: options.traceIdents,
                                    optimizationSteps: [  ]
                                });
                                var newImplementations = Data_FoldableWithIndex.foldrWithIndex(Data_Map_Internal.foldableWithIndexMap)(insert)(v.implementations)(v2.value1.implementations);
                                return Control_Bind.discard(Control_Bind.discardUnit)(Bind1)(options.onCodegenModule({
                                    moduleCount: buildEnv.moduleCount,
                                    moduleIndex: buildEnv.moduleIndex,
                                    implementations: newImplementations
                                })(v1)(v2.value1)(v2.value0))(function () {
                                    return Control_Applicative.pure(Applicative0)({
                                        directives: Data_FoldableWithIndex.foldrWithIndex(Data_Map_Internal.foldableWithIndexMap)(insert1)(v.directives)(v2.value1.directives),
                                        implementations: newImplementations,
                                        moduleIndex: v.moduleIndex + 1 | 0
                                    });
                                });
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Builder (line 51, column 5 - line 83, column 12): " + [ mbCachedMod.constructor.name ]);
                        });
                    });
                };
            };
            return $$void(Data_List.foldM(dictMonad)(go)({
                directives: options.directives,
                implementations: Data_Map_Internal.empty,
                moduleIndex: 0
            })(coreFnModules));
        };
    };
};
export {
    buildModules
};
