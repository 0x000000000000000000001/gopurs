import * as Control_Semigroupoid from "../Control.Semigroupoid/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Profunctor from "../Data.Profunctor/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var strongFn = {
    first: function (a2b) {
        return function (v) {
            return new Data_Tuple.Tuple(a2b(v.value0), v.value1);
        };
    },
    second: /* #__PURE__ */ Data_Functor.map(Data_Tuple.functorTuple),
    Profunctor0: function () {
        return Data_Profunctor.profunctorFn;
    }
};
var second = function (dict) {
    return dict.second;
};
var first = function (dict) {
    return dict.first;
};

// | Compose a value acting on a `Tuple` from two values, each acting on one of
// | the components of the `Tuple`.
// |
// | Specializing `(***)` to function application would look like this:
// | ```
// | (***) :: forall a b c d. (a -> b) -> (c -> d) -> (Tuple a c) -> (Tuple b d)
// | ```
// | We take two functions, `f` and `g`, and we transform them into a single function which
// | takes a `Tuple` and maps `f` over the first element and `g` over the second.  Just like `bi-map`
// | would do for the `bi-functor` instance of `Tuple`.
var splitStrong = function (dictSemigroupoid) {
    return function (dictStrong) {
        return function (l) {
            return function (r) {
                return Control_Semigroupoid.composeFlipped(dictSemigroupoid)(first(dictStrong)(l))(second(dictStrong)(r));
            };
        };
    };
};

// | Compose a value which introduces a `Tuple` from two values, each introducing
// | one side of the `Tuple`.
// |
// | This combinator is useful when assembling values from smaller components,
// | because it provides a way to support two different types of output.
// |
// | Specializing `(&&&)` to function application would look like this:
// | ```
// | (&&&) :: forall a b c. (a -> b) -> (a -> c) -> (a -> (Tuple b c))
// | ```
// | We take two functions, `f` and `g`, with the same parameter type and we transform them into a
// | single function which takes one parameter and returns a `Tuple` of the results of running
// | `f` and `g` on the parameter, respectively.  This allows us to run two parallel computations
// | on the same input and return both results in a `Tuple`.
var fanout = function (dictSemigroupoid) {
    return function (dictStrong) {
        var Profunctor0 = dictStrong.Profunctor0();
        return function (l) {
            return function (r) {
                return Data_Profunctor.lcmap(Profunctor0)(function (a) {
                    return new Data_Tuple.Tuple(a, a);
                })(splitStrong(dictSemigroupoid)(dictStrong)(l)(r));
            };
        };
    };
};
export {
    first,
    second,
    splitStrong,
    fanout,
    strongFn
};
