import * as Data_Array from "../Data.Array/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Set from "../Data.Set/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";

// | Computes the set of reachable modules starting from a list of entry modules.
var moduleReachability = function (entryMods) {
    return function (modulesMap) {
        var getModuleDeps = function (modName) {
            var v = Data_Map_Internal.lookup(PureScript_Backend_Optimizer_CoreFn.ordModuleName)(modName)(modulesMap);
            if (v instanceof Data_Maybe.Just) {
                var impls = Data_Map_Internal.toUnfoldable(Data_Unfoldable.unfoldableArray)(v.value0.implementations);
                
                // Accumulate module names directly into a Set to avoid massive intermediate arrays
var allMods = Data_Foldable.foldl(Data_Foldable.foldableArray)(function (acc) {
                    return function (v1) {
                        return Data_Foldable.foldl(Data_Set.foldableSet)(function (acc2) {
                            return function (v2) {
                                if (v2.value0 instanceof Data_Maybe.Just) {
                                    return Data_Set.insert(PureScript_Backend_Optimizer_CoreFn.ordModuleName)(v2.value0.value0)(acc2);
                                };
                                if (v2.value0 instanceof Data_Maybe.Nothing) {
                                    return acc2;
                                };
                                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Reachability (line 50, column 27 - line 52, column 44): " + [ v2.value0.constructor.name ]);
                            };
                        })(acc)(v1.value1.value0.deps);
                    };
                })(Data_Set.empty)(impls);
                return Data_Set.union(PureScript_Backend_Optimizer_CoreFn.ordModuleName)(v.value0.imports)(allMods);
            };
            if (v instanceof Data_Maybe.Nothing) {
                return Data_Set.empty;
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Reachability (line 42, column 7 - line 60, column 20): " + [ v.constructor.name ]);
        };
        var go = function ($copy_seen) {
            return function ($copy_unvisited) {
                var $tco_var_seen = $copy_seen;
                var $tco_done = false;
                var $tco_result;
                function $tco_loop(seen, unvisited) {
                    var $15 = Data_Array["null"](unvisited);
                    if ($15) {
                        $tco_done = true;
                        return seen;
                    };
                    var newSeen = Data_Set.union(PureScript_Backend_Optimizer_CoreFn.ordModuleName)(seen)(Data_Set.fromFoldable(Data_Foldable.foldableArray)(PureScript_Backend_Optimizer_CoreFn.ordModuleName)(unvisited));
                    
                    // Use a Set to accumulate new dependencies to avoid huge arrays
var newDepsSet = Data_Foldable.foldl(Data_Foldable.foldableArray)(function (acc) {
                        return function (modName) {
                            return Data_Set.union(PureScript_Backend_Optimizer_CoreFn.ordModuleName)(acc)(getModuleDeps(modName));
                        };
                    })(Data_Set.empty)(unvisited);
                    var nextUnvisitedSet = Data_Set.difference(PureScript_Backend_Optimizer_CoreFn.ordModuleName)(newDepsSet)(newSeen);
                    $tco_var_seen = newSeen;
                    $copy_unvisited = Data_Array.fromFoldable(Data_Set.foldableSet)(nextUnvisitedSet);
                    return;
                };
                while (!$tco_done) {
                    $tco_result = $tco_loop($tco_var_seen, $copy_unvisited);
                };
                return $tco_result;
            };
        };
        return go(Data_Set.empty)(entryMods);
    };
};
export {
    moduleReachability
};
