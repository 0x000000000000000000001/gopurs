import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad_Gen from "../Control.Monad.Gen/index.js";
import * as Control_Monad_Gen_Class from "../Control.Monad.Gen.Class/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_NonEmpty from "../Data.NonEmpty/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";

// | Creates a generator that outputs `Tuple` values, choosing values from a
// | pair of generators for each slot in the tuple.
var genTuple = function (dictApply) {
    return Control_Apply.lift2(dictApply)(Data_Tuple.Tuple.create);
};

// | Creates a generator that outputs `NonEmpty` values, choosing values from a
// | generator for each of the items.
// |
// | The size of the value will be determined by the current size state
// | for the generator. To generate a value of a particular size, use the
// | `resize` function from the `MonadGen` class first.
var genNonEmpty = function (dictMonadRec) {
    return function (dictMonadGen) {
        var Bind1 = (dictMonadGen.Monad0()).Bind1();
        var Apply0 = Bind1.Apply0();
        var Functor0 = (Bind1.Apply0()).Functor0();
        return function (dictUnfoldable) {
            return function (gen) {
                return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Data_NonEmpty.NonEmpty.create)(gen))(Control_Monad_Gen_Class.resize(dictMonadGen)((function () {
                    var $28 = Data_Ord.max(Data_Ord.ordInt)(0);
                    return function ($29) {
                        return $28((function (v) {
                            return v - 1 | 0;
                        })($29));
                    };
                })())(Control_Monad_Gen.unfoldable(dictMonadRec)(dictMonadGen)(dictUnfoldable)(gen)));
            };
        };
    };
};

// | Creates a generator that outputs `Maybe` values, choosing a value from
// | another generator for the inner value, with an adjustable bias for how
// | often `Just` is returned vs `Nothing`. A bias ≤ 0.0 will always
// | return `Nothing`, a bias ≥ 1.0 will always return `Just`.
var genMaybe$prime = function (dictMonadGen) {
    var Monad0 = dictMonadGen.Monad0();
    var Bind1 = Monad0.Bind1();
    var Functor0 = ((Monad0.Bind1()).Apply0()).Functor0();
    var Applicative0 = Monad0.Applicative0();
    return function (bias) {
        return function (gen) {
            return Control_Bind.bind(Bind1)(Control_Monad_Gen_Class.chooseFloat(dictMonadGen)(0.0)(1.0))(function (n) {
                var $26 = n < bias;
                if ($26) {
                    return Data_Functor.map(Functor0)(Data_Maybe.Just.create)(gen);
                };
                return Control_Applicative.pure(Applicative0)(Data_Maybe.Nothing.value);
            });
        };
    };
};

// | Creates a generator that outputs `Maybe` values, choosing a value from
// | another generator for the inner value. The generator has a 75% chance of
// | returning a `Just` over a `Nothing`.
var genMaybe = function (dictMonadGen) {
    return genMaybe$prime(dictMonadGen)(0.75);
};

// | Creates a generator that outputs `Identity` values, choosing a value from
// | another generator for the inner value.
var genIdentity = function (dictFunctor) {
    return Data_Functor.map(dictFunctor)(Data_Identity.Identity);
};

// | Creates a generator that outputs `Either` values, choosing a value from a
// | `Left` or the `Right` with adjustable bias. As the bias value increases,
// | the chance of returning a `Left` value rises. A bias ≤ 0.0 will always
// | return `Right`, a bias ≥ 1.0 will always return `Left`.
var genEither$prime = function (dictMonadGen) {
    var Monad0 = dictMonadGen.Monad0();
    var Bind1 = Monad0.Bind1();
    var Functor0 = ((Monad0.Bind1()).Apply0()).Functor0();
    return function (bias) {
        return function (genA) {
            return function (genB) {
                return Control_Bind.bind(Bind1)(Control_Monad_Gen_Class.chooseFloat(dictMonadGen)(0.0)(1.0))(function (n) {
                    var $27 = n < bias;
                    if ($27) {
                        return Data_Functor.map(Functor0)(Data_Either.Left.create)(genA);
                    };
                    return Data_Functor.map(Functor0)(Data_Either.Right.create)(genB);
                });
            };
        };
    };
};

// | Creates a generator that outputs `Either` values, choosing a value from a
// | `Left` or the `Right` with even probability.
var genEither = function (dictMonadGen) {
    return genEither$prime(dictMonadGen)(0.5);
};
export {
    genEither,
    genEither$prime,
    genIdentity,
    genMaybe,
    genMaybe$prime,
    genTuple,
    genNonEmpty
};
