import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Set from "../Data.Set/index.js";

// | `SemigroupMap k v` provides a `Semigroup` instance for `Map k v` whose
// | definition depends on the `Semigroup` instance for the `v` type.
// | You should only use this type when you need `Data.Map` to have
// | a `Semigroup` instance.
// |
// | ```purescript
// | let
// |   s :: forall key value. key -> value -> SemigroupMap key value
// |   s k v = SemigroupMap (singleton k v)
// |
// | (s 1     "foo") <> (s 1     "bar") == (s 1  "foobar")
// | (s 1 (First 1)) <> (s 1 (First 2)) == (s 1 (First 1))
// | (s 1  (Last 1)) <> (s 1  (Last 2)) == (s 1  (Last 2))
// | ```
var SemigroupMap = function (x) {
    return x;
};
var traversableWithIndexSemigroupMap = Data_Map_Internal.traversableWithIndexMap;
var traversableSemigroupMap = Data_Map_Internal.traversableMap;
var showSemigroupMap = function (dictShow) {
    var showMap = Data_Map_Internal.showMap(dictShow);
    return function (dictShow1) {
        return showMap(dictShow1);
    };
};
var semigroupSemigroupMap = function (dictOrd) {
    return function (dictSemigroup) {
        var append = Data_Semigroup.append(dictSemigroup);
        return {
            append: function (v) {
                return function (v1) {
                    return Data_Map_Internal.unionWith(dictOrd)(append)(v)(v1);
                };
            }
        };
    };
};
var plusSemigroupMap = function (dictOrd) {
    return Data_Map_Internal.plusMap(dictOrd);
};
var ordSemigroupMap = function (dictOrd) {
    var ordMap = Data_Map_Internal.ordMap(dictOrd);
    return function (dictOrd1) {
        return ordMap(dictOrd1);
    };
};
var ord1SemigroupMap = function (dictOrd) {
    return Data_Map_Internal.ord1Map(dictOrd);
};
var newtypeSemigroupMap = {
    Coercible0: function () {
        return undefined;
    }
};
var monoidSemigroupMap = function (dictOrd) {
    var semigroupSemigroupMap1 = semigroupSemigroupMap(dictOrd);
    return function (dictSemigroup) {
        var semigroupSemigroupMap2 = semigroupSemigroupMap1(dictSemigroup);
        return {
            mempty: Data_Map_Internal.empty,
            Semigroup0: function () {
                return semigroupSemigroupMap2;
            }
        };
    };
};

// | The set of keys of the given map.
// | See also `Data.Set.fromMap`.
var keys = /* #__PURE__ */ (function () {
    var $36 = Data_Functor["void"](Data_Map_Internal.functorMap);
    return function ($37) {
        return Data_Set.fromMap($36($37));
    };
})();
var functorWithIndexSemigroupMap = Data_Map_Internal.functorWithIndexMap;
var functorSemigroupMap = Data_Map_Internal.functorMap;
var foldableWithIndexSemigroupMap = Data_Map_Internal.foldableWithIndexMap;
var foldableSemigroupMap = Data_Map_Internal.foldableMap;
var eqSemigroupMap = function (dictEq) {
    var eqMap = Data_Map_Internal.eqMap(dictEq);
    return function (dictEq1) {
        return eqMap(dictEq1);
    };
};
var eq1SemigroupMap = function (dictEq) {
    return Data_Map_Internal.eq1Map(dictEq);
};
var bindSemigroupMap = function (dictOrd) {
    return Data_Map_Internal.bindMap(dictOrd);
};
var applySemigroupMap = function (dictOrd) {
    return Data_Map_Internal.applyMap(dictOrd);
};
var altSemigroupMap = function (dictOrd) {
    return Data_Map_Internal.altMap(dictOrd);
};
export {
    keys,
    SemigroupMap,
    eq1SemigroupMap,
    eqSemigroupMap,
    ord1SemigroupMap,
    ordSemigroupMap,
    newtypeSemigroupMap,
    showSemigroupMap,
    semigroupSemigroupMap,
    monoidSemigroupMap,
    altSemigroupMap,
    plusSemigroupMap,
    functorSemigroupMap,
    functorWithIndexSemigroupMap,
    applySemigroupMap,
    bindSemigroupMap,
    foldableSemigroupMap,
    foldableWithIndexSemigroupMap,
    traversableSemigroupMap,
    traversableWithIndexSemigroupMap
};
export {
    alter,
    any,
    anyWithKey,
    catMaybes,
    checkValid,
    delete,
    difference,
    empty,
    filter,
    filterKeys,
    filterWithKey,
    findMax,
    findMin,
    foldSubmap,
    fromFoldable,
    fromFoldableWith,
    fromFoldableWithIndex,
    insert,
    insertWith,
    intersection,
    intersectionWith,
    isEmpty,
    isSubmap,
    lookup,
    lookupGE,
    lookupGT,
    lookupLE,
    lookupLT,
    mapMaybe,
    mapMaybeWithKey,
    member,
    pop,
    showTree,
    singleton,
    size,
    submap,
    toUnfoldable,
    toUnfoldableUnordered,
    union,
    unionWith,
    unions,
    update,
    values
} from "../Data.Map.Internal/index.js";
