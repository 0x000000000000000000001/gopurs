import * as Control_Category from "../Control.Category/index.js";
import * as Control_Semigroupoid from "../Control.Semigroupoid/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Profunctor from "../Data.Profunctor/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var right = function (dict) {
    return dict.right;
};
var left = function (dict) {
    return dict.left;
};

// | Compose a value acting on a sum from two values, each acting on one of
// | the components of the sum.
// |
// | Specializing `(+++)` to function application would look like this:
// | ```
// | (+++) :: forall a b c d. (a -> b) -> (c -> d) -> (Either a c) -> (Either b d)
// | ```
// | We take two functions, `f` and `g`, and we transform them into a single function which
// | takes an `Either`and maps `f` over the left side and `g` over the right side.  Just like
// | `bi-map` would do for the `bi-functor` instance of `Either`.
var splitChoice = function (dictSemigroupoid) {
    return function (dictChoice) {
        return function (l) {
            return function (r) {
                return Control_Semigroupoid.composeFlipped(dictSemigroupoid)(left(dictChoice)(l))(right(dictChoice)(r));
            };
        };
    };
};

// | Compose a value which eliminates a sum from two values, each eliminating
// | one side of the sum.
// |
// | This combinator is useful when assembling values from smaller components,
// | because it provides a way to support two different types of input.
// |
// | Specializing `(|||)` to function application would look like this:
// | ```
// | (|||) :: forall a b c d. (a -> c) -> (b -> c) -> Either a b -> c
// | ```
// | We take two functions, `f` and `g`, which both return the same type `c` and we transform them into a
// | single function which takes an `Either` value with the parameter type of `f` on the left side and
// | the parameter type of `g` on the right side. The function then runs either `f` or `g`, depending on
// | whether the `Either` value is a `Left` or a `Right`.
// | This allows us to bundle two different computations which both have the same result type into one
// | function which will run the approriate computation based on the parameter supplied in the `Either` value.
var fanin = function (dictSemigroupoid) {
    return function (dictChoice) {
        var Profunctor0 = dictChoice.Profunctor0();
        return function (l) {
            return function (r) {
                return Data_Profunctor.rmap(Profunctor0)(Data_Either.either(identity)(identity))(splitChoice(dictSemigroupoid)(dictChoice)(l)(r));
            };
        };
    };
};
var choiceFn = {
    left: function (v) {
        return function (v1) {
            if (v1 instanceof Data_Either.Left) {
                return new Data_Either.Left(v(v1.value0));
            };
            if (v1 instanceof Data_Either.Right) {
                return new Data_Either.Right(v1.value0);
            };
            throw new Error("Failed pattern match at Data.Profunctor.Choice (line 32, column 1 - line 35, column 16): " + [ v.constructor.name, v1.constructor.name ]);
        };
    },
    right: /* #__PURE__ */ Data_Functor.map(Data_Either.functorEither),
    Profunctor0: function () {
        return Data_Profunctor.profunctorFn;
    }
};
export {
    left,
    right,
    splitChoice,
    fanin,
    choiceFn
};
