// | Prisms are used for selecting cases of a type, most often a sum
// | type. Consider this:
// |
// | ```purescript
// | data Fill -- think of a paint program filling a shape
// |   = NoFill
// |   | Solid Color
// |   | ...
// | ```
// |
// | A prism that focuses on `Solid` fills could be written like this:
// |
// | ```purescript
// | solidFocus :: Prism' Fill Color
// | solidFocus = prism' Solid case _ of
// |   Solid color -> Just color
// |   _ -> Nothing
// | ```
// |
// | ... and used like this:
// |
// | ```purescript
// | preview solidFocus (Solid Color.white) == Just Color.white
// | preview solidFocus NoFill == Nothing
// |
// | is solidFocus (Solid Color.white) == true
// | ```
// |
// | `review` can be used to go from a `Color` to a `Fill`:
// |
// | ```purescript
// | review solidFocus Color.white == Solid Color.white
// | ```
// | 
// | For more information, see `PrismsForSumTypes.purs` in the
// | `examples/src` directory.
// |
// | ---------------
// |
// | A well-behaved `Prism` will follow these laws:
// |
// | **review-preview**: `preview` retrieves what `review` creates. Equationally:
// |   
// | ```purescript
// | review prism >>> preview prism ≡ Just
// | ```
// |
// | An example:
// | 
// | ```purescript
// | Color.white # review solidFocus # preview solidFocus
// |   == Just Color.white
// | ```
// | 
// | **preview-review**: If `preview` retrieves something, `review` can create
// | the original from that something. Equationally:
// | 
// | ```purescript
// | if preview prism s ≡ Just a then review prism a ≡ s
// | ```
// |
// | An example:
// |
// | ```purescript
// | Solid Color.white # preview solidFocus <#> review solidFocus
// |   == Solid Color.white
// | ```
import * as Control_Alternative from "../Control.Alternative/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_HeytingAlgebra from "../Data.HeytingAlgebra/index.js";
import * as Data_Lens_Internal_Market from "../Data.Lens.Internal.Market/index.js";
import * as Data_Lens_Internal_Tagged from "../Data.Lens.Internal.Tagged/index.js";
import * as Data_Lens_Types from "../Data.Lens.Types/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Profunctor from "../Data.Profunctor/index.js";
import * as Data_Profunctor_Choice from "../Data.Profunctor.Choice/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var identity1 = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var guard = /* #__PURE__ */ Control_Alternative.guard(Data_Maybe.alternativeMaybe);
var withPrism = function (l) {
    return function (f) {
        var v = l(new Data_Lens_Internal_Market.Market(identity, Data_Either.Right.create));
        return f(v.value0)(v.value1);
    };
};

// | Create the "whole" corresponding to a specific "part":
// |
// | ```purescript
// | review solidFocus Color.white == Solid Color.white
// | ```
var review = /* #__PURE__ */ Data_Newtype.under()()(Data_Lens_Internal_Tagged.Tagged);

// | Create a `Prism` from a constructor and a matcher function that
// | produces an `Either`:
// | 
// | ```purescript
// | solidFocus :: Prism' Fill Color
// | solidFocus = prism Solid case _ of
// |   Solid color -> Right color
// |   anotherCase -> Left anotherCase
// | ```
// |
// | _Note_: The matcher function returns a result wrapped in `Either t`
// | to allow for type-changing prisms in the case where the input does
// | not match.
var prism = function (to) {
    return function (fro) {
        return function (dictChoice) {
            var Profunctor0 = dictChoice.Profunctor0();
            return function (pab) {
                return Data_Profunctor.dimap(Profunctor0)(fro)(Data_Either.either(identity1)(identity1))(Data_Profunctor_Choice.right(dictChoice)(Data_Profunctor.rmap(Profunctor0)(to)(pab)));
            };
        };
    };
};

// | Create a `Prism` from a constructor and a matcher function that
// | produces a `Maybe`:
// | 
// | ```purescript
// | solidFocus :: Prism' Fill Color
// | solidFocus = prism' Solid case _ of
// |   Solid color -> Just color
// |   _ -> Nothing
// | ```
var prism$prime = function (to) {
    return function (fro) {
        return function (dictChoice) {
            return prism(to)(function (s) {
                return Data_Maybe.maybe(new Data_Either.Left(s))(Data_Either.Right.create)(fro(s));
            })(dictChoice);
        };
    };
};

// | `nearly` is a variant of `only`. Like `only`, `nearly` produces
// | a prism that matches
// | a single value. Unlike `only`, it uses a predicate you supply
// | instead of depending on `class Eq`: 
// | 
// | ```purescript
// | solidWhiteFocus :: Prism' Fill Unit
// | solidWhiteFocus = nearly (Solid Color.white) predicate
// |   where
// |     predicate candidate =
// |       color.toHexString == Color.white.toHexString
// | ```
var nearly = function (x) {
    return function (f) {
        return function (dictChoice) {
            return prism$prime(Data_Function["const"](x))(function ($26) {
                return guard(f($26));
            })(dictChoice);
        };
    };
};

// | `only` focuses not just on a case, but a specific value of that case.
// | 
// | ```purescript
// | solidWhiteFocus :: Prism' Fill Unit
// | solidWhiteFocus = only $ Solid Color.white
// |
// | is      solidWhiteFocus (Solid Color.white) == true
// | preview solidWhiteFocus (Solid Color.white) == Just unit
// | review  solidWhiteFocus unit                == Solid Color.white
// | ```
// |
// | *Note*: `only` depends on `Eq`. Strange definitions of `(==)`
// | (for example, that it counts any `Fill` as being equal to `Solid Color.white`)
// | will create a prism that violates the preview-review law. 
var only = function (dictEq) {
    return function (x) {
        return function (dictChoice) {
            return nearly(x)(function (v) {
                return Data_Eq.eq(dictEq)(v)(x);
            })(dictChoice);
        };
    };
};
var matching = function (l) {
    return withPrism(l)(function (v) {
        return function (f) {
            return f;
        };
    });
};

//| Ask if `preview prism` would produce a `Just`.
var is = function (dictHeytingAlgebra) {
    return function (l) {
        var $27 = Data_Either.either(Data_Function["const"](Data_HeytingAlgebra.ff(dictHeytingAlgebra)))(Data_Function["const"](Data_HeytingAlgebra.tt(dictHeytingAlgebra)));
        var $28 = matching(l);
        return function ($29) {
            return $27($28($29));
        };
    };
};

//| Ask if `preview prism` would produce a `Nothing`.
var isn$primet = function (dictHeytingAlgebra) {
    var not = Data_HeytingAlgebra.not(dictHeytingAlgebra);
    return function (l) {
        var $30 = is(dictHeytingAlgebra)(l);
        return function ($31) {
            return not($30($31));
        };
    };
};
var clonePrism = function (l) {
    return function (dictChoice) {
        return withPrism(l)(function (x) {
            return function (y) {
                return function (p) {
                    return prism(x)(y)(dictChoice)(p);
                };
            };
        });
    };
};

// Ported from Haskell: https://hackage.haskell.org/package/lens-4.16/docs/src/Control-Lens-Prism.html#below
// | `lift` a `Prism` through a `Traversable` functor, giving a `Prism` that matches 
// | only if all the elements of the container match the `Prism`.
// |
// | ``` purescript
// | >>> [Left 1, Right "foo", Left 4, Right "woot"]^..below _Right
// | []
// | ```
// | 
// | ``` purescript
// | >>> [Right "hail hydra!", Right "foo", Right "blah", Right "woot"]^..below _Right
// | [["hail hydra!","foo","blah","woot"]]
// | ```
var below = function (dictTraversable) {
    var Functor0 = dictTraversable.Functor0();
    return function (k) {
        return function (dictChoice) {
            return withPrism(k)(function (bt) {
                return function (seta) {
                    return prism(Data_Functor.map(Functor0)(bt))(function (s) {
                        var v = Data_Traversable.traverse(dictTraversable)(Data_Either.applicativeEither)(seta)(s);
                        if (v instanceof Data_Either.Left) {
                            return new Data_Either.Left(s);
                        };
                        if (v instanceof Data_Either.Right) {
                            return new Data_Either.Right(v.value0);
                        };
                        throw new Error("Failed pattern match at Data.Lens.Prism (line 199, column 7 - line 201, column 27): " + [ v.constructor.name ]);
                    })(dictChoice);
                };
            });
        };
    };
};
export {
    prism$prime,
    prism,
    only,
    nearly,
    review,
    is,
    isn$primet,
    matching,
    clonePrism,
    withPrism,
    below
};
