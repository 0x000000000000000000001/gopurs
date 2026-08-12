import * as $foreign from "./foreign.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Semigroup_Traversable from "../Data.Semigroup.Traversable/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var unfoldr1 = function (dict) {
    return dict.unfoldr1;
};
var unfoldable1Maybe = {
    unfoldr1: function (f) {
        return function (b) {
            return new Data_Maybe.Just(Data_Tuple.fst(f(b)));
        };
    }
};
var unfoldable1Array = {
    unfoldr1: /* #__PURE__ */ $foreign.unfoldr1ArrayImpl(Data_Maybe.isNothing)(/* #__PURE__ */ Data_Maybe.fromJust())(Data_Tuple.fst)(Data_Tuple.snd)
};

// | Replicate a value `n` times. At least one value will be produced, so values
// | `n` less than 1 will be treated as 1.
// |
// | ``` purescript
// | replicate1 2 "foo" == (NEL.cons "foo" (NEL.singleton "foo") :: NEL.NonEmptyList String)
// | replicate1 0 "foo" == (NEL.singleton "foo" :: NEL.NonEmptyList String)
// | ```
var replicate1 = function (dictUnfoldable1) {
    return function (n) {
        return function (v) {
            var step = function (i) {
                if (i <= 0) {
                    return new Data_Tuple.Tuple(v, Data_Maybe.Nothing.value);
                };
                if (Data_Boolean.otherwise) {
                    return new Data_Tuple.Tuple(v, new Data_Maybe.Just(i - 1 | 0));
                };
                throw new Error("Failed pattern match at Data.Unfoldable1 (line 68, column 5 - line 68, column 39): " + [ i.constructor.name ]);
            };
            return unfoldr1(dictUnfoldable1)(step)(n - 1 | 0);
        };
    };
};

// | Perform an `Apply` action `n` times (at least once, so values `n` less
// | than 1 will be treated as 1), and accumulate the results.
// |
// | ``` purescript
// | > replicate1A 2 (randomInt 1 10) :: Effect (NEL.NonEmptyList Int)
// | (NonEmptyList (NonEmpty 8 (2 : Nil)))
// | > replicate1A 0 (randomInt 1 10) :: Effect (NEL.NonEmptyList Int)
// | (NonEmptyList (NonEmpty 4 Nil))
// | ```
var replicate1A = function (dictApply) {
    return function (dictUnfoldable1) {
        return function (dictTraversable1) {
            return function (n) {
                return function (m) {
                    return Data_Semigroup_Traversable.sequence1(dictTraversable1)(dictApply)(replicate1(dictUnfoldable1)(n)(m));
                };
            };
        };
    };
};

// | Contain a single value. For example:
// |
// | ``` purescript
// | singleton "foo" == (NEL.singleton "foo" :: NEL.NonEmptyList String)
// | ```
var singleton = function (dictUnfoldable1) {
    return replicate1(dictUnfoldable1)(1);
};

// | Create an `Unfoldable1` containing a range of values, including both
// | endpoints.
// |
// | ``` purescript
// | range 0 0 == (NEL.singleton 0 :: NEL.NonEmptyList Int)
// | range 1 2 == (NEL.cons 1 (NEL.singleton 2) :: NEL.NonEmptyList Int)
// | range 2 0 == (NEL.cons 2 (NEL.cons 1 (NEL.singleton 0)) :: NEL.NonEmptyList Int)
// | ```
var range = function (dictUnfoldable1) {
    return function (start) {
        return function (end) {
            var go = function (delta) {
                return function (i) {
                    var i$prime = i + delta | 0;
                    return new Data_Tuple.Tuple(i, (function () {
                        var $10 = i === end;
                        if ($10) {
                            return Data_Maybe.Nothing.value;
                        };
                        return new Data_Maybe.Just(i$prime);
                    })());
                };
            };
            var delta = (function () {
                var $11 = end >= start;
                if ($11) {
                    return 1;
                };
                return -1 | 0;
            })();
            return unfoldr1(dictUnfoldable1)(go(delta))(start);
        };
    };
};

// | Create an `Unfoldable1` by repeated application of a function to a seed value.
// | For example:
// |
// | ``` purescript
// | (iterateN 5 (_ + 1) 0 :: Array Int) == [0, 1, 2, 3, 4]
// | (iterateN 5 (_ + 1) 0 :: NonEmptyArray Int) == NonEmptyArray [0, 1, 2, 3, 4]
// |
// | (iterateN 0 (_ + 1) 0 :: Array Int) == [0]
// | (iterateN 0 (_ + 1) 0 :: NonEmptyArray Int) == NonEmptyArray [0]
// | ```
var iterateN = function (dictUnfoldable1) {
    return function (n) {
        return function (f) {
            return function (s) {
                var go = function (v) {
                    return new Data_Tuple.Tuple(v.value0, (function () {
                        var $13 = v.value1 > 0;
                        if ($13) {
                            return new Data_Maybe.Just(new Data_Tuple.Tuple(f(v.value0), v.value1 - 1 | 0));
                        };
                        return Data_Maybe.Nothing.value;
                    })());
                };
                return unfoldr1(dictUnfoldable1)(go)(new Data_Tuple.Tuple(s, n - 1 | 0));
            };
        };
    };
};
export {
    unfoldr1,
    replicate1,
    replicate1A,
    singleton,
    range,
    iterateN,
    unfoldable1Array,
    unfoldable1Maybe
};
