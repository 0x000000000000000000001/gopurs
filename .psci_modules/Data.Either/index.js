import * as Control_Alt from "../Control.Alt/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Data_Bounded from "../Data.Bounded/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Functor_Invariant from "../Data.Functor.Invariant/index.js";
import * as Data_Generic_Rep from "../Data.Generic.Rep/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Unit from "../Data.Unit/index.js";

// | The `Either` type is used to represent a choice between two types of value.
// |
// | A common use case for `Either` is error handling, where `Left` is used to
// | carry an error value and `Right` is used to carry a success value.
var Left = /* #__PURE__ */ (function () {
    function Left(value0) {
        this.value0 = value0;
    };
    Left.create = function (value0) {
        return new Left(value0);
    };
    return Left;
})();

// | The `Either` type is used to represent a choice between two types of value.
// |
// | A common use case for `Either` is error handling, where `Left` is used to
// | carry an error value and `Right` is used to carry a success value.
var Right = /* #__PURE__ */ (function () {
    function Right(value0) {
        this.value0 = value0;
    };
    Right.create = function (value0) {
        return new Right(value0);
    };
    return Right;
})();

// | The `Show` instance allows `Either` values to be rendered as a string with
// | `show` whenever there is an `Show` instance for both type the `Either` can
// | contain.
var showEither = function (dictShow) {
    return function (dictShow1) {
        return {
            show: function (v) {
                if (v instanceof Left) {
                    return "(Left " + (Data_Show.show(dictShow)(v.value0) + ")");
                };
                if (v instanceof Right) {
                    return "(Right " + (Data_Show.show(dictShow1)(v.value0) + ")");
                };
                throw new Error("Failed pattern match at Data.Either (line 173, column 1 - line 175, column 46): " + [ v.constructor.name ]);
            }
        };
    };
};

// | Similar to `note`, but for use in cases where the default value may be
// | expensive to compute.
// |
// | ```purescript
// | note' (\_ -> "default") Nothing = Left "default"
// | note' (\_ -> "default") (Just 1) = Right 1
// | ```
var note$prime = function (f) {
    return Data_Maybe["maybe$prime"](function ($124) {
        return Left.create(f($124));
    })(Right.create);
};

// | Takes a default and a `Maybe` value, if the value is a `Just`, turn it into
// | a `Right`, if the value is a `Nothing` use the provided default as a `Left`
// |
// | ```purescript
// | note "default" Nothing = Left "default"
// | note "default" (Just 1) = Right 1
// | ```
var note = function (a) {
    return Data_Maybe.maybe(new Left(a))(Right.create);
};
var genericEither = {
    to: function (x) {
        if (x instanceof Data_Generic_Rep.Inl) {
            return new Left(x.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr) {
            return new Right(x.value0);
        };
        throw new Error("Failed pattern match at Data.Either (line 33, column 1 - line 33, column 56): " + [ x.constructor.name ]);
    },
    from: function (x) {
        if (x instanceof Left) {
            return new Data_Generic_Rep.Inl(x.value0);
        };
        if (x instanceof Right) {
            return new Data_Generic_Rep.Inr(x.value0);
        };
        throw new Error("Failed pattern match at Data.Either (line 33, column 1 - line 33, column 56): " + [ x.constructor.name ]);
    }
};

// | The `Functor` instance allows functions to transform the contents of a
// | `Right` with the `<$>` operator:
// |
// | ``` purescript
// | f <$> Right x == Right (f x)
// | ```
// |
// | `Left` values are untouched:
// |
// | ``` purescript
// | f <$> Left y == Left y
// | ```
var functorEither = {
    map: function (f) {
        return function (m) {
            if (m instanceof Left) {
                return new Left(m.value0);
            };
            if (m instanceof Right) {
                return new Right(f(m.value0));
            };
            throw new Error("Failed pattern match at Data.Either (line 0, column 0 - line 0, column 0): " + [ m.constructor.name ]);
        };
    }
};
var invariantEither = {
    imap: /* #__PURE__ */ Data_Functor_Invariant.imapF(functorEither)
};

// | Similar to `fromRight` but for use in cases where the default value may be
// | expensive to compute. As PureScript is not lazy, the standard `fromRight`
// | has to evaluate the default value before returning the result,
// | whereas here the value is only computed when the `Either` is known
// | to be `Left`.
var fromRight$prime = function (v) {
    return function (v1) {
        if (v1 instanceof Right) {
            return v1.value0;
        };
        return v(Data_Unit.unit);
    };
};

// | A function that extracts the value from the `Right` data constructor.
// | The first argument is a default value, which will be returned in the
// | case where a `Left` is passed to `fromRight`.
var fromRight = function (v) {
    return function (v1) {
        if (v1 instanceof Right) {
            return v1.value0;
        };
        return v;
    };
};

// | Similar to `fromLeft` but for use in cases where the default value may be
// | expensive to compute. As PureScript is not lazy, the standard `fromLeft`
// | has to evaluate the default value before returning the result,
// | whereas here the value is only computed when the `Either` is known
// | to be `Right`.
var fromLeft$prime = function (v) {
    return function (v1) {
        if (v1 instanceof Left) {
            return v1.value0;
        };
        return v(Data_Unit.unit);
    };
};

// | A function that extracts the value from the `Left` data constructor.
// | The first argument is a default value, which will be returned in the
// | case where a `Right` is passed to `fromLeft`.
var fromLeft = function (v) {
    return function (v1) {
        if (v1 instanceof Left) {
            return v1.value0;
        };
        return v;
    };
};

// | The `Extend` instance allows sequencing of `Either` values and functions
// | that accept an `Either` and return a non-`Either` result using the
// | `<<=` operator.
// |
// | ``` purescript
// | f <<= Left x = Left x
// | f <<= Right x = Right (f (Right x))
// | ```
var extendEither = {
    extend: function (v) {
        return function (v1) {
            if (v1 instanceof Left) {
                return new Left(v1.value0);
            };
            return new Right(v(v1));
        };
    },
    Functor0: function () {
        return functorEither;
    }
};

// | The `Eq` instance allows `Either` values to be checked for equality with
// | `==` and inequality with `/=` whenever there is an `Eq` instance for both
// | types the `Either` can contain.
var eqEither = function (dictEq) {
    return function (dictEq1) {
        return {
            eq: function (x) {
                return function (y) {
                    if (x instanceof Left && y instanceof Left) {
                        return Data_Eq.eq(dictEq)(x.value0)(y.value0);
                    };
                    if (x instanceof Right && y instanceof Right) {
                        return Data_Eq.eq(dictEq1)(x.value0)(y.value0);
                    };
                    return false;
                };
            }
        };
    };
};

// | The `Ord` instance allows `Either` values to be compared with
// | `compare`, `>`, `>=`, `<` and `<=` whenever there is an `Ord` instance for
// | both types the `Either` can contain.
// |
// | Any `Left` value is considered to be less than a `Right` value.
var ordEither = function (dictOrd) {
    var eqEither1 = eqEither(dictOrd.Eq0());
    return function (dictOrd1) {
        var eqEither2 = eqEither1(dictOrd1.Eq0());
        return {
            compare: function (x) {
                return function (y) {
                    if (x instanceof Left && y instanceof Left) {
                        return Data_Ord.compare(dictOrd)(x.value0)(y.value0);
                    };
                    if (x instanceof Left) {
                        return Data_Ordering.LT.value;
                    };
                    if (y instanceof Left) {
                        return Data_Ordering.GT.value;
                    };
                    if (x instanceof Right && y instanceof Right) {
                        return Data_Ord.compare(dictOrd1)(x.value0)(y.value0);
                    };
                    throw new Error("Failed pattern match at Data.Either (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
                };
            },
            Eq0: function () {
                return eqEither2;
            }
        };
    };
};
var eq1Either = function (dictEq) {
    var eqEither1 = eqEither(dictEq);
    return {
        eq1: function (dictEq1) {
            return Data_Eq.eq(eqEither1(dictEq1));
        }
    };
};
var ord1Either = function (dictOrd) {
    var ordEither1 = ordEither(dictOrd);
    var eq1Either1 = eq1Either(dictOrd.Eq0());
    return {
        compare1: function (dictOrd1) {
            return Data_Ord.compare(ordEither1(dictOrd1));
        },
        Eq10: function () {
            return eq1Either1;
        }
    };
};

// | Takes two functions and an `Either` value, if the value is a `Left` the
// | inner value is applied to the first function, if the value is a `Right`
// | the inner value is applied to the second function.
// |
// | ``` purescript
// | either f g (Left x) == f x
// | either f g (Right y) == g y
// | ```
var either = function (v) {
    return function (v1) {
        return function (v2) {
            if (v2 instanceof Left) {
                return v(v2.value0);
            };
            if (v2 instanceof Right) {
                return v1(v2.value0);
            };
            throw new Error("Failed pattern match at Data.Either (line 208, column 1 - line 208, column 64): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
        };
    };
};

// | Turns an `Either` into a `Maybe`, by throwing potential `Left` values away and converting
// | them into `Nothing`. `Right` values get turned into `Just`s.
// |
// | ```purescript
// | hush (Left "ParseError") = Nothing
// | hush (Right 42) = Just 42
// | ```
var hush = /* #__PURE__ */ (function () {
    return either(Data_Function["const"](Data_Maybe.Nothing.value))(Data_Maybe.Just.create);
})();

// | Returns `true` when the `Either` value was constructed with `Left`.
var isLeft = /* #__PURE__ */ either(/* #__PURE__ */ Data_Function["const"](true))(/* #__PURE__ */ Data_Function["const"](false));

// | Returns `true` when the `Either` value was constructed with `Right`.
var isRight = /* #__PURE__ */ either(/* #__PURE__ */ Data_Function["const"](false))(/* #__PURE__ */ Data_Function["const"](true));

// | Combine two alternatives.
var choose = function (dictAlt) {
    var Functor0 = dictAlt.Functor0();
    return function (a) {
        return function (b) {
            return Control_Alt.alt(dictAlt)(Data_Functor.map(Functor0)(Left.create)(a))(Data_Functor.map(Functor0)(Right.create)(b));
        };
    };
};
var boundedEither = function (dictBounded) {
    var ordEither1 = ordEither(dictBounded.Ord0());
    return function (dictBounded1) {
        var ordEither2 = ordEither1(dictBounded1.Ord0());
        return {
            top: new Right(Data_Bounded.top(dictBounded1)),
            bottom: new Left(Data_Bounded.bottom(dictBounded)),
            Ord0: function () {
                return ordEither2;
            }
        };
    };
};

// | Turns an `Either` into a `Maybe`, by throwing potential `Right` values away and converting
// | them into `Nothing`. `Left` values get turned into `Just`s.
// |
// | ```purescript
// | blush (Left "ParseError") = Just "Parse Error"
// | blush (Right 42) = Nothing
// | ```
var blush = /* #__PURE__ */ (function () {
    return either(Data_Maybe.Just.create)(Data_Function["const"](Data_Maybe.Nothing.value));
})();

// | The `Apply` instance allows functions contained within a `Right` to
// | transform a value contained within a `Right` using the `(<*>)` operator:
// |
// | ``` purescript
// | Right f <*> Right x == Right (f x)
// | ```
// |
// | `Left` values are left untouched:
// |
// | ``` purescript
// | Left f <*> Right x == Left f
// | Right f <*> Left y == Left y
// | ```
// |
// | Combining `Functor`'s `<$>` with `Apply`'s `<*>` can be used to transform a
// | pure function to take `Either`-typed arguments so `f :: a -> b -> c`
// | becomes `f :: Either l a -> Either l b -> Either l c`:
// |
// | ``` purescript
// | f <$> Right x <*> Right y == Right (f x y)
// | ```
// |
// | The `Left`-preserving behaviour of both operators means the result of
// | an expression like the above but where any one of the values is `Left`
// | means the whole result becomes `Left` also, taking the first `Left` value
// | found:
// |
// | ``` purescript
// | f <$> Left x <*> Right y == Left x
// | f <$> Right x <*> Left y == Left y
// | f <$> Left x <*> Left y == Left x
// | ```
var applyEither = {
    apply: function (v) {
        return function (v1) {
            if (v instanceof Left) {
                return new Left(v.value0);
            };
            if (v instanceof Right) {
                return Data_Functor.map(functorEither)(v.value0)(v1);
            };
            throw new Error("Failed pattern match at Data.Either (line 70, column 1 - line 72, column 30): " + [ v.constructor.name, v1.constructor.name ]);
        };
    },
    Functor0: function () {
        return functorEither;
    }
};

// | The `Bind` instance allows sequencing of `Either` values and functions that
// | return an `Either` by using the `>>=` operator:
// |
// | ``` purescript
// | Left x >>= f = Left x
// | Right x >>= f = f x
// | ```
// |
// | `Either`'s "do notation" can be understood to work like this:
// | ``` purescript
// | x :: forall e a. Either e a
// | x = --
// |
// | y :: forall e b. Either e b
// | y = --
// |
// | foo :: forall e a. (a -> b -> c) -> Either e c
// | foo f = do
// |   x' <- x
// |   y' <- y
// |   pure (f x' y')
// | ```
// |
// | ...which is equivalent to...
// |
// | ``` purescript
// | x >>= (\x' -> y >>= (\y' -> pure (f x' y')))
// | ```
// |
// | ...and is the same as writing...
// |
// | ```
// | foo :: forall e a. (a -> b -> c) -> Either e c
// | foo f = case x of
// |   Left e ->
// |     Left e
// |   Right x -> case y of
// |     Left e ->
// |       Left e
// |     Right y ->
// |       Right (f x y)
// | ```
var bindEither = {
    bind: /* #__PURE__ */ either(function (e) {
        return function (v) {
            return new Left(e);
        };
    })(function (a) {
        return function (f) {
            return f(a);
        };
    }),
    Apply0: function () {
        return applyEither;
    }
};
var semigroupEither = function (dictSemigroup) {
    var append = Data_Semigroup.append(dictSemigroup);
    return {
        append: function (x) {
            return function (y) {
                return Control_Apply.apply(applyEither)(Data_Functor.map(functorEither)(append)(x))(y);
            };
        }
    };
};

// | The `Applicative` instance enables lifting of values into `Either` with the
// | `pure` function:
// |
// | ``` purescript
// | pure x :: Either _ _ == Right x
// | ```
// |
// | Combining `Functor`'s `<$>` with `Apply`'s `<*>` and `Applicative`'s
// | `pure` can be used to pass a mixture of `Either` and non-`Either` typed
// | values to a function that does not usually expect them, by using `pure`
// | for any value that is not already `Either` typed:
// |
// | ``` purescript
// | f <$> Right x <*> pure y == Right (f x y)
// | ```
// |
// | Even though `pure = Right` it is recommended to use `pure` in situations
// | like this as it allows the choice of `Applicative` to be changed later
// | without having to go through and replace `Right` with a new constructor.
var applicativeEither = /* #__PURE__ */ (function () {
    return {
        pure: Right.create,
        Apply0: function () {
            return applyEither;
        }
    };
})();

// | The `Monad` instance guarantees that there are both `Applicative` and
// | `Bind` instances for `Either`.
var monadEither = {
    Applicative0: function () {
        return applicativeEither;
    },
    Bind1: function () {
        return bindEither;
    }
};

// | The `Alt` instance allows for a choice to be made between two `Either`
// | values with the `<|>` operator, where the first `Right` encountered
// | is taken.
// |
// | ``` purescript
// | Right x <|> Right y == Right x
// | Left x <|> Right y == Right y
// | Left x <|> Left y == Left y
// | ```
var altEither = {
    alt: function (v) {
        return function (v1) {
            if (v instanceof Left) {
                return v1;
            };
            return v;
        };
    },
    Functor0: function () {
        return functorEither;
    }
};
export {
    Left,
    Right,
    either,
    choose,
    isLeft,
    isRight,
    fromLeft,
    fromLeft$prime,
    fromRight,
    fromRight$prime,
    note,
    note$prime,
    hush,
    blush,
    functorEither,
    genericEither,
    invariantEither,
    applyEither,
    applicativeEither,
    altEither,
    bindEither,
    monadEither,
    extendEither,
    showEither,
    eqEither,
    eq1Either,
    ordEither,
    ord1Either,
    boundedEither,
    semigroupEither
};
