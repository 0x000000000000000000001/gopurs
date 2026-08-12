// | This module exports the `NonEmptyArray` constructor.
// |
// | It is **NOT** intended for public use and is **NOT** versioned.
// |
// | Its content may change **in any way**, **at any time** and
// | **without notice**.
import * as $foreign from "./foreign.js";
import * as Control_Alt from "../Control.Alt/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad from "../Control.Monad/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_FoldableWithIndex from "../Data.FoldableWithIndex/index.js";
import * as Data_Function_Uncurried from "../Data.Function.Uncurried/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_FunctorWithIndex from "../Data.FunctorWithIndex/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Semigroup_Foldable from "../Data.Semigroup.Foldable/index.js";
import * as Data_Semigroup_Traversable from "../Data.Semigroup.Traversable/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_TraversableWithIndex from "../Data.TraversableWithIndex/index.js";
import * as Data_Unfoldable1 from "../Data.Unfoldable1/index.js";

// | An array that is known not to be empty.
// |
// | You can use the constructor to create a `NonEmptyArray` that isn't
// | non-empty, breaking the guarantee behind this newtype. It is
// | provided as an escape hatch mainly for the `Data.Array.NonEmpty`
// | and `Data.Array` modules. Use this at your own risk when you know
// | what you are doing.
var NonEmptyArray = function (x) {
    return x;
};
var unfoldable1NonEmptyArray = Data_Unfoldable1.unfoldable1Array;
var traversableWithIndexNonEmptyArray = Data_TraversableWithIndex.traversableWithIndexArray;
var traversableNonEmptyArray = Data_Traversable.traversableArray;
var showNonEmptyArray = function (dictShow) {
    var showArray = Data_Show.showArray(dictShow);
    return {
        show: function (v) {
            return "(NonEmptyArray " + (Data_Show.show(showArray)(v) + ")");
        }
    };
};
var semigroupNonEmptyArray = Data_Semigroup.semigroupArray;
var ordNonEmptyArray = function (dictOrd) {
    return Data_Ord.ordArray(dictOrd);
};
var ord1NonEmptyArray = Data_Ord.ord1Array;
var monadNonEmptyArray = Control_Monad.monadArray;
var functorWithIndexNonEmptyArray = Data_FunctorWithIndex.functorWithIndexArray;
var functorNonEmptyArray = Data_Functor.functorArray;
var foldableWithIndexNonEmptyArray = Data_FoldableWithIndex.foldableWithIndexArray;
var foldableNonEmptyArray = Data_Foldable.foldableArray;
var foldable1NonEmptyArray = {
    foldMap1: function (dictSemigroup) {
        return Data_Semigroup_Foldable.foldMap1DefaultL(foldable1NonEmptyArray)(functorNonEmptyArray)(dictSemigroup);
    },
    foldr1: /* #__PURE__ */ Data_Function_Uncurried.runFn2($foreign.foldr1Impl),
    foldl1: /* #__PURE__ */ Data_Function_Uncurried.runFn2($foreign.foldl1Impl),
    Foldable0: function () {
        return foldableNonEmptyArray;
    }
};
var traversable1NonEmptyArray = {
    traverse1: function (dictApply) {
        var apply = Control_Apply.apply(dictApply);
        var map = Data_Functor.map(dictApply.Functor0());
        return function (f) {
            return $foreign.traverse1Impl(apply, map, f);
        };
    },
    sequence1: function (dictApply) {
        return Data_Semigroup_Traversable.sequence1Default(traversable1NonEmptyArray)(dictApply);
    },
    Foldable10: function () {
        return foldable1NonEmptyArray;
    },
    Traversable1: function () {
        return traversableNonEmptyArray;
    }
};
var eqNonEmptyArray = function (dictEq) {
    return Data_Eq.eqArray(dictEq);
};
var eq1NonEmptyArray = Data_Eq.eq1Array;
var bindNonEmptyArray = Control_Bind.bindArray;
var applyNonEmptyArray = Control_Apply.applyArray;
var applicativeNonEmptyArray = Control_Applicative.applicativeArray;
var altNonEmptyArray = Control_Alt.altArray;
export {
    NonEmptyArray,
    showNonEmptyArray,
    eqNonEmptyArray,
    eq1NonEmptyArray,
    ordNonEmptyArray,
    ord1NonEmptyArray,
    semigroupNonEmptyArray,
    functorNonEmptyArray,
    functorWithIndexNonEmptyArray,
    foldableNonEmptyArray,
    foldableWithIndexNonEmptyArray,
    foldable1NonEmptyArray,
    unfoldable1NonEmptyArray,
    traversableNonEmptyArray,
    traversableWithIndexNonEmptyArray,
    traversable1NonEmptyArray,
    applyNonEmptyArray,
    applicativeNonEmptyArray,
    bindNonEmptyArray,
    monadNonEmptyArray,
    altNonEmptyArray
};
