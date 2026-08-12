import * as Control_Bind from "../Control.Bind/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_HeytingAlgebra from "../Data.HeytingAlgebra/index.js";
import * as Data_List from "../Data.List/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Map from "../Data.Map/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
import * as Data_Set from "../Data.Set/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";
import * as PureScript_CST_Types from "../PureScript.CST.Types/index.js";
var add = /* #__PURE__ */ Data_Semiring.add(Data_Semiring.semiringInt);
var fromFoldable = /* #__PURE__ */ Data_Map_Internal.fromFoldable(PureScript_CST_Types.ordModuleName);
var fromFoldable1 = /* #__PURE__ */ fromFoldable(Data_Foldable.foldableArray);
var fromFoldable2 = /* #__PURE__ */ fromFoldable(Data_Foldable.foldableArray);
var lookup = /* #__PURE__ */ Data_Map_Internal.lookup(PureScript_CST_Types.ordModuleName);
var toUnfoldable = /* #__PURE__ */ Data_List.toUnfoldable(Data_Unfoldable.unfoldableArray);
var Sorted = /* #__PURE__ */ (function () {
    function Sorted(value0) {
        this.value0 = value0;
    };
    Sorted.create = function (value0) {
        return new Sorted(value0);
    };
    return Sorted;
})();
var CycleDetected = /* #__PURE__ */ (function () {
    function CycleDetected(value0) {
        this.value0 = value0;
    };
    CycleDetected.create = function (value0) {
        return new CycleDetected(value0);
    };
    return CycleDetected;
})();
var topoSort = function (dictOrd) {
    var insert = Data_Set.insert(dictOrd);
    return function (graph) {
        var isRoot = function (v) {
            var $19 = v.value1 === 0;
            if ($19) {
                return new Data_Maybe.Just(v.value0);
            };
            return Data_Maybe.Nothing.value;
        };
        var importCounts = Data_Map_Internal.fromFoldableWith(dictOrd)(Data_Foldable.foldableArray)(add)(Control_Bind.bind(Control_Bind.bindArray)(Data_Map_Internal.toUnfoldable(Data_Unfoldable.unfoldableArray)(graph))(function (v) {
            return Data_Semigroup.append(Data_Semigroup.semigroupArray)([ new Data_Tuple.Tuple(v.value0, 0) ])(Data_Functor.map(Data_Functor.functorArray)(Data_Function.flip(Data_Tuple.Tuple.create)(1))(Data_Set.toUnfoldable(Data_Unfoldable.unfoldableArray)(v.value1)));
        }));
        var startingModules = Data_Map.keys(Data_Map_Internal.filterWithKey(dictOrd)(function (k) {
            return function (v) {
                return Data_Maybe.isJust(isRoot(new Data_Tuple.Tuple(k, v)));
            };
        })(importCounts));
        var depthFirst = function (v) {
            var $26 = Data_Set.member(dictOrd)(v.curr)(v.visited);
            if ($26) {
                return new Data_Maybe.Just(new Data_List_Types.Cons(v.curr, v.path));
            };
            var $27 = Data_Maybe.maybe(true)(Data_Set.isEmpty)(Data_Map_Internal.lookup(dictOrd)(v.curr)(graph));
            if ($27) {
                return Data_Maybe.Nothing.value;
            };
            return Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(dictOrd)(v.curr)(graph))(function (reachable) {
                return Data_Foldable.foldl(Data_Set.foldableSet)(function (b) {
                    return function (a) {
                        var $28 = Data_Maybe.isJust(b);
                        if ($28) {
                            return b;
                        };
                        return depthFirst({
                            path: new Data_List_Types.Cons(v.curr, v.path),
                            visited: Data_Set.insert(dictOrd)(v.curr)(v.visited),
                            curr: a
                        });
                    };
                })(Data_Maybe.Nothing.value)(reachable);
            });
        };
        var decrementImport = function (usages) {
            return function (k) {
                return Data_Map_Internal.insertWith(dictOrd)(add)(k)(-1 | 0)(usages);
            };
        };
        var appendRoots = function (usages) {
            return function (roots) {
                return function (curr) {
                    return Data_Maybe.maybe(roots)(Data_Function.flip(insert)(roots))(Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(dictOrd)(curr)(usages))(function (count) {
                        return isRoot(new Data_Tuple.Tuple(curr, count));
                    }));
                };
            };
        };
        var go = function ($copy_v) {
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(v) {
                var v1 = Data_Set.findMin(v.roots);
                if (v1 instanceof Data_Maybe.Nothing) {
                    var $34 = Data_Foldable.all(Data_Map_Internal.foldableMap)(Data_HeytingAlgebra.heytingAlgebraBoolean)(Data_Eq.eq(Data_Eq.eqInt)(0))(v.usages);
                    if ($34) {
                        $tco_done = true;
                        return new Data_Either.Right({
                            roots: v.roots,
                            sorted: v.sorted,
                            usages: v.usages
                        });
                    };
                    var nonLeaf = Data_Map.keys(Data_Map_Internal.filterWithKey(dictOrd)(function (a) {
                        return function (count) {
                            return count > 0 && !Data_Maybe.maybe(true)(Data_Set.isEmpty)(Data_Map_Internal.lookup(dictOrd)(a)(graph));
                        };
                    })(v.usages));
                    var detectCycles = Data_Foldable.foldl(Data_Set.foldableSet)(function (b) {
                        return function (a) {
                            var $35 = Data_Maybe.isJust(b);
                            if ($35) {
                                return b;
                            };
                            return depthFirst({
                                path: Data_List_Types.Nil.value,
                                visited: Data_Set.empty,
                                curr: a
                            });
                        };
                    })(Data_Maybe.Nothing.value)(nonLeaf);
                    if (detectCycles instanceof Data_Maybe.Just) {
                        $tco_done = true;
                        return new Data_Either.Left(detectCycles.value0);
                    };
                    if (detectCycles instanceof Data_Maybe.Nothing) {
                        $tco_done = true;
                        return new Data_Either.Left(Data_List_Types.Nil.value);
                    };
                    throw new Error("Failed pattern match at PureScript.CST.ModuleGraph (line 79, column 9 - line 81, column 30): " + [ detectCycles.constructor.name ]);
                };
                if (v1 instanceof Data_Maybe.Just) {
                    var reachable = Data_Maybe.fromMaybe(Data_Set.empty)(Data_Map_Internal.lookup(dictOrd)(v1.value0)(graph));
                    var usages$prime = Data_Foldable.foldl(Data_Set.foldableSet)(decrementImport)(v.usages)(reachable);
                    $copy_v = {
                        roots: Data_Foldable.foldl(Data_Set.foldableSet)(appendRoots(usages$prime))(Data_Set["delete"](dictOrd)(v1.value0)(v.roots))(reachable),
                        sorted: new Data_List_Types.Cons(v1.value0, v.sorted),
                        usages: usages$prime
                    };
                    return;
                };
                throw new Error("Failed pattern match at PureScript.CST.ModuleGraph (line 66, column 34 - line 91, column 10): " + [ v1.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($copy_v);
            };
            return $tco_result;
        };
        return Data_Functor.map(Data_Either.functorEither)(function (v) {
            return v.sorted;
        })(go({
            roots: startingModules,
            sorted: Data_List_Types.Nil.value,
            usages: importCounts
        }));
    };
};
var moduleGraph = function (k) {
    var getImportName = function (v) {
        return v.module.name;
    };
    var go = function (v) {
        return new Data_Tuple.Tuple(v.name.name, Data_Set.fromFoldable(Data_Foldable.foldableArray)(PureScript_CST_Types.ordModuleName)(Data_Functor.map(Data_Functor.functorArray)(getImportName)(v.imports)));
    };
    var $55 = Data_Functor.map(Data_Functor.functorArray)(function ($57) {
        return go(k($57));
    });
    return function ($56) {
        return fromFoldable1($55($56));
    };
};
var sortModules = function (k) {
    return function (moduleHeaders) {
        var graph = moduleGraph(k)(moduleHeaders);
        var getModuleName = function (v) {
            return v.name.name;
        };
        var knownModuleHeaders = fromFoldable2(Data_Functor.map(Data_Functor.functorArray)(function (a) {
            return new Data_Tuple.Tuple(getModuleName(k(a)), a);
        })(moduleHeaders));
        var lookupModuleHeaders = (function () {
            var $58 = Data_Array.mapMaybe(Data_Function.flip(lookup)(knownModuleHeaders));
            return function ($59) {
                return $58(toUnfoldable($59));
            };
        })();
        var v = topoSort(PureScript_CST_Types.ordModuleName)(graph);
        if (v instanceof Data_Either.Left) {
            return new CycleDetected(lookupModuleHeaders(v.value0));
        };
        if (v instanceof Data_Either.Right) {
            return new Sorted(lookupModuleHeaders(v.value0));
        };
        throw new Error("Failed pattern match at PureScript.CST.ModuleGraph (line 51, column 3 - line 53, column 56): " + [ v.constructor.name ]);
    };
};
export {
    moduleGraph,
    sortModules,
    Sorted,
    CycleDetected
};
