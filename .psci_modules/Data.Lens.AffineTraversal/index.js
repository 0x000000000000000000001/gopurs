// | This module defines functions for working with affine traversals.
// | An `AffineTraversal` is a `Traversal` that applies to at most one element.
// |
// | These arise most frequently as the composition of a `Lens` with a `Prism`.
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Semigroupoid from "../Control.Semigroupoid/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Lens_Internal_Stall from "../Data.Lens.Internal.Stall/index.js";
import * as Data_Lens_Types from "../Data.Lens.Types/index.js";
import * as Data_Profunctor from "../Data.Profunctor/index.js";
import * as Data_Profunctor_Choice from "../Data.Profunctor.Choice/index.js";
import * as Data_Profunctor_Strong from "../Data.Profunctor.Strong/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var withAffineTraversal = function (l) {
    return function (f) {
        var v = l(new Data_Lens_Internal_Stall.Stall(Data_Function["const"](Control_Category.identity(Control_Category.categoryFn)), Data_Either.Right.create));
        return f(v.value0)(v.value1);
    };
};
var affineTraversal$prime = function (to) {
    return function (dictStrong) {
        return function (dictChoice) {
            var Profunctor0 = dictChoice.Profunctor0();
            return function (pab) {
                return Data_Profunctor.dimap(Profunctor0)(to)(function (v) {
                    return Data_Either.either(identity)(v.value0)(v.value1);
                })(Data_Profunctor_Strong.second(dictStrong)(Data_Profunctor_Choice.right(dictChoice)(pab)));
            };
        };
    };
};
var affineTraversal = function (set) {
    return function (pre) {
        return function (dictStrong) {
            return function (dictChoice) {
                return affineTraversal$prime(Data_Profunctor_Strong.fanout(Control_Semigroupoid.semigroupoidFn)(Data_Profunctor_Strong.strongFn)(set)(pre))(dictStrong)(dictChoice);
            };
        };
    };
};
var cloneAffineTraversal = function (l) {
    return function (dictStrong) {
        return function (dictChoice) {
            return withAffineTraversal(l)(function (x) {
                return function (y) {
                    return function (p) {
                        return affineTraversal(x)(y)(dictStrong)(dictChoice)(p);
                    };
                };
            });
        };
    };
};
export {
    affineTraversal,
    affineTraversal$prime,
    withAffineTraversal,
    cloneAffineTraversal
};
