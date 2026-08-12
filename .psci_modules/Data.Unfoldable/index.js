// | This module provides a type class for _unfoldable functors_, i.e.
// | functors which support an `unfoldr` operation.
// |
// | This allows us to unify various operations on arrays, lists,
// | sequences, etc.
import * as $foreign from "./foreign.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unfoldable1 from "../Data.Unfoldable1/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
var unfoldr = function (dict) {
    return dict.unfoldr;
};
var unfoldableMaybe = {
    unfoldr: function (f) {
        return function (b) {
            return Data_Functor.map(Data_Maybe.functorMaybe)(Data_Tuple.fst)(f(b));
        };
    },
    Unfoldable10: function () {
        return Data_Unfoldable1.unfoldable1Maybe;
    }
};
var unfoldableArray = {
    unfoldr: /* #__PURE__ */ $foreign.unfoldrArrayImpl(Data_Maybe.isNothing)(/* #__PURE__ */ Data_Maybe.fromJust())(Data_Tuple.fst)(Data_Tuple.snd),
    Unfoldable10: function () {
        return Data_Unfoldable1.unfoldable1Array;
    }
};

// | Replicate a value some natural number of times.
// | For example:
// |
// | ``` purescript
// | replicate 2 "foo" == (["foo", "foo"] :: Array String)
// | ```
var replicate = function (dictUnfoldable) {
    return function (n) {
        return function (v) {
            var step = function (i) {
                var $7 = i <= 0;
                if ($7) {
                    return Data_Maybe.Nothing.value;
                };
                return new Data_Maybe.Just(new Data_Tuple.Tuple(v, i - 1 | 0));
            };
            return unfoldr(dictUnfoldable)(step)(n);
        };
    };
};

// | Perform an Applicative action `n` times, and accumulate all the results.
// |
// | ``` purescript
// | > replicateA 5 (randomInt 1 10) :: Effect (Array Int)
// | [1,3,2,7,5]
// | ```
var replicateA = function (dictApplicative) {
    return function (dictUnfoldable) {
        return function (dictTraversable) {
            return function (n) {
                return function (m) {
                    return Data_Traversable.sequence(dictTraversable)(dictApplicative)(replicate(dictUnfoldable)(n)(m));
                };
            };
        };
    };
};

// | The container with no elements - unfolded with zero iterations.
// | For example:
// |
// | ``` purescript
// | none == ([] :: Array Unit)
// | ```
var none = function (dictUnfoldable) {
    return unfoldr(dictUnfoldable)(Data_Function["const"](Data_Maybe.Nothing.value))(Data_Unit.unit);
};

// | Convert a Maybe to any Unfoldable, such as lists or arrays.
// |
// | ``` purescript
// | fromMaybe (Nothing :: Maybe Int) == []
// | fromMaybe (Just 1) == [1]
// | ```
var fromMaybe = function (dictUnfoldable) {
    return unfoldr(dictUnfoldable)(function (b) {
        return Data_Functor.map(Data_Maybe.functorMaybe)(Data_Function.flip(Data_Tuple.Tuple.create)(Data_Maybe.Nothing.value))(b);
    });
};
export {
    unfoldr,
    replicate,
    replicateA,
    none,
    fromMaybe,
    unfoldableArray,
    unfoldableMaybe
};
export {
    iterateN,
    range,
    replicate1,
    replicate1A,
    singleton,
    unfoldr1
} from "../Data.Unfoldable1/index.js";
