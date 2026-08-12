// | This module defines the list monad transformer, `ListT`.
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Monad from "../Control.Monad/index.js";
import * as Control_Monad_Rec_Class from "../Control.Monad.Rec.Class/index.js";
import * as Control_Monad_ST_Class from "../Control.Monad.ST.Class/index.js";
import * as Control_Monad_Trans_Class from "../Control.Monad.Trans.Class/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Lazy from "../Data.Lazy/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Effect_Class from "../Effect.Class/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var identity1 = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);

// | The result of a single step in a `ListT` computation. Either:
// |
// | - Computation has finished (`Done`), or
// | - A result has been returned, along with the next part of the computation (`Yield`).
// |
// | The `Skip` constructor allows us to avoid traversing lists during certain operations.
var Yield = /* #__PURE__ */ (function () {
    function Yield(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Yield.create = function (value0) {
        return function (value1) {
            return new Yield(value0, value1);
        };
    };
    return Yield;
})();

// | The result of a single step in a `ListT` computation. Either:
// |
// | - Computation has finished (`Done`), or
// | - A result has been returned, along with the next part of the computation (`Yield`).
// |
// | The `Skip` constructor allows us to avoid traversing lists during certain operations.
var Skip = /* #__PURE__ */ (function () {
    function Skip(value0) {
        this.value0 = value0;
    };
    Skip.create = function (value0) {
        return new Skip(value0);
    };
    return Skip;
})();

// | The result of a single step in a `ListT` computation. Either:
// |
// | - Computation has finished (`Done`), or
// | - A result has been returned, along with the next part of the computation (`Yield`).
// |
// | The `Skip` constructor allows us to avoid traversing lists during certain operations.
var Done = /* #__PURE__ */ (function () {
    function Done() {

    };
    Done.value = new Done();
    return Done;
})();

// | The list monad transformer.
// |
// | This monad transformer extends the base monad with _non-determinism_.
// | That is, the transformed monad supports the same effects as the base monad
// | but with multiple return values.
var ListT = function (x) {
    return x;
};

// | Defer evaluation of a list.
var wrapLazy = function (dictApplicative) {
    return function (v) {
        return Control_Applicative.pure(dictApplicative)(new Skip(v));
    };
};

// | Lift a computation from the base monad.
var wrapEffect = function (dictFunctor) {
    return function (v) {
        return Data_Functor.map(dictFunctor)(function ($279) {
            return Skip.create(Data_Lazy.defer(Data_Function["const"]($279)));
        })(v);
    };
};

// | Unfold a list using an effectful generator function.
var unfold = function (dictMonad) {
    var Functor0 = ((dictMonad.Bind1()).Apply0()).Functor0();
    return function (f) {
        return function (z) {
            var g = function (v) {
                if (v instanceof Data_Maybe.Just) {
                    return new Yield(v.value0.value1, Data_Lazy.defer(function (v1) {
                        return unfold(dictMonad)(f)(v.value0.value0);
                    }));
                };
                if (v instanceof Data_Maybe.Nothing) {
                    return Done.value;
                };
                throw new Error("Failed pattern match at Control.Monad.List.Trans (line 130, column 3 - line 130, column 60): " + [ v.constructor.name ]);
            };
            return Data_Functor.map(Functor0)(g)(f(z));
        };
    };
};

// | Perform the first step of a computation in the `ListT` monad.
var uncons = function (dictMonad) {
    var pure = Control_Applicative.pure(dictMonad.Applicative0());
    var Applicative0 = dictMonad.Applicative0();
    var Bind1 = dictMonad.Bind1();
    return function (v) {
        var g = function (v1) {
            if (v1 instanceof Yield) {
                return pure(new Data_Maybe.Just(new Data_Tuple.Tuple(v1.value0, Data_Lazy.force(v1.value1))));
            };
            if (v1 instanceof Skip) {
                return uncons(dictMonad)(Data_Lazy.force(v1.value0));
            };
            if (v1 instanceof Done) {
                return Control_Applicative.pure(Applicative0)(Data_Maybe.Nothing.value);
            };
            throw new Error("Failed pattern match at Control.Monad.List.Trans (line 195, column 3 - line 195, column 50): " + [ v1.constructor.name ]);
        };
        return Control_Bind.bind(Bind1)(v)(g);
    };
};

// | Extract all but the first element of a list.
var tail = function (dictMonad) {
    var Functor0 = ((dictMonad.Bind1()).Apply0()).Functor0();
    return function (l) {
        return Data_Functor.map(Functor0)(Data_Functor.map(Data_Maybe.functorMaybe)(Data_Tuple.snd))(uncons(dictMonad)(l));
    };
};

// | Lift a computation on list steps to a computation on whole lists.
var stepMap = function (dictFunctor) {
    return function (f) {
        return function (v) {
            return Data_Functor.map(dictFunctor)(f)(v);
        };
    };
};

// | Take elements from the front of a list while a predicate holds.
var takeWhile = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (f) {
        var g = function (v) {
            if (v instanceof Yield) {
                var $196 = f(v.value0);
                if ($196) {
                    return new Yield(v.value0, Data_Functor.map(Data_Lazy.functorLazy)(takeWhile(dictApplicative)(f))(v.value1));
                };
                return Done.value;
            };
            if (v instanceof Skip) {
                return new Skip(Data_Functor.map(Data_Lazy.functorLazy)(takeWhile(dictApplicative)(f))(v.value0));
            };
            if (v instanceof Done) {
                return Done.value;
            };
            throw new Error("Failed pattern match at Control.Monad.List.Trans (line 154, column 3 - line 154, column 68): " + [ v.constructor.name ]);
        };
        return stepMap(Functor0)(g);
    };
};

// | Fold a list from the left, accumulating the list of results using the specified function.
var scanl = function (dictMonad) {
    var Functor0 = ((dictMonad.Bind1()).Apply0()).Functor0();
    return function (f) {
        return function (b) {
            return function (l) {
                var g = function (v) {
                    var h = function (v1) {
                        if (v1 instanceof Yield) {
                            var b$prime$prime = f(v.value0)(v1.value0);
                            return new Data_Maybe.Just(new Data_Tuple.Tuple(new Data_Tuple.Tuple(b$prime$prime, Data_Lazy.force(v1.value1)), v.value0));
                        };
                        if (v1 instanceof Skip) {
                            return new Data_Maybe.Just(new Data_Tuple.Tuple(new Data_Tuple.Tuple(v.value0, Data_Lazy.force(v1.value0)), v.value0));
                        };
                        if (v1 instanceof Done) {
                            return Data_Maybe.Nothing.value;
                        };
                        throw new Error("Failed pattern match at Control.Monad.List.Trans (line 248, column 5 - line 248, column 78): " + [ v1.constructor.name ]);
                    };
                    return Data_Functor.map(Functor0)(h)(v.value1);
                };
                return unfold(dictMonad)(g)(new Data_Tuple.Tuple(b, l));
            };
        };
    };
};

// | Prepend an element to a lazily-evaluated list.
var prepend$prime = function (dictApplicative) {
    return function (h) {
        return function (t) {
            return Control_Applicative.pure(dictApplicative)(new Yield(h, t));
        };
    };
};

// | Prepend an element to a list.
var prepend = function (dictApplicative) {
    return function (h) {
        return function (t) {
            return prepend$prime(dictApplicative)(h)(Data_Lazy.defer(Data_Function["const"](t)));
        };
    };
};

// | The empty list.
var nil = function (dictApplicative) {
    return Control_Applicative.pure(dictApplicative)(Done.value);
};

// | Create a list with one element.
var singleton = function (dictApplicative) {
    var nil1 = nil(dictApplicative);
    return function (a) {
        return prepend(dictApplicative)(a)(nil1);
    };
};

// | Take a number of elements from the front of a list.
var take = function (dictApplicative) {
    var nil1 = nil(dictApplicative);
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (v) {
        return function (v1) {
            if (v === 0) {
                return nil1;
            };
            var f = function (v2) {
                if (v2 instanceof Yield) {
                    return new Yield(v2.value0, Data_Functor.map(Data_Lazy.functorLazy)(take(dictApplicative)(v - 1 | 0))(v2.value1));
                };
                if (v2 instanceof Skip) {
                    return new Skip(Data_Functor.map(Data_Lazy.functorLazy)(take(dictApplicative)(v))(v2.value0));
                };
                if (v2 instanceof Done) {
                    return Done.value;
                };
                throw new Error("Failed pattern match at Control.Monad.List.Trans (line 147, column 3 - line 147, column 47): " + [ v2.constructor.name ]);
            };
            return stepMap(Functor0)(f)(v1);
        };
    };
};

// | Zip the elements of two lists, combining elements at the same position from each list.
var zipWith$prime = function (dictMonad) {
    var Applicative0 = dictMonad.Applicative0();
    var Functor0 = ((dictMonad.Bind1()).Apply0()).Functor0();
    var prepend$prime1 = prepend$prime(dictMonad.Applicative0());
    var Bind1 = dictMonad.Bind1();
    return function (f) {
        var g = function (v) {
            return function (v1) {
                if (v1 instanceof Data_Maybe.Nothing) {
                    return Control_Applicative.pure(Applicative0)(nil(Applicative0));
                };
                if (v instanceof Data_Maybe.Nothing) {
                    return Control_Applicative.pure(Applicative0)(nil(Applicative0));
                };
                if (v instanceof Data_Maybe.Just && v1 instanceof Data_Maybe.Just) {
                    return Data_Functor.map(Functor0)(Data_Function.flip(prepend$prime1)(Data_Lazy.defer(function (v2) {
                        return zipWith$prime(dictMonad)(f)(v.value0.value1)(v1.value0.value1);
                    })))(f(v.value0.value0)(v1.value0.value0));
                };
                throw new Error("Failed pattern match at Control.Monad.List.Trans (line 260, column 3 - line 260, column 25): " + [ v.constructor.name, v1.constructor.name ]);
            };
        };
        var loop = function (fa) {
            return function (fb) {
                return wrapEffect(Functor0)(Control_Bind.bind(Bind1)(uncons(dictMonad)(fa))(function (ua) {
                    return Control_Bind.bind(Bind1)(uncons(dictMonad)(fb))(function (ub) {
                        return g(ua)(ub);
                    });
                }));
            };
        };
        return loop;
    };
};

// | Zip the elements of two lists, combining elements at the same position from each list.
var zipWith = function (dictMonad) {
    var pure = Control_Applicative.pure(dictMonad.Applicative0());
    return function (f) {
        var g = function (a) {
            return function (b) {
                return pure(f(a)(b));
            };
        };
        return zipWith$prime(dictMonad)(g);
    };
};
var newtypeListT = {
    Coercible0: function () {
        return undefined;
    }
};

// | Apply a function to the elements of a list, keeping only those return values which contain a result.
var mapMaybe = function (dictFunctor) {
    return function (f) {
        var g = function (v) {
            if (v instanceof Yield) {
                return Data_Maybe.fromMaybe(Skip.create)(Data_Functor.map(Data_Maybe.functorMaybe)(Yield.create)(f(v.value0)))(Data_Functor.map(Data_Lazy.functorLazy)(mapMaybe(dictFunctor)(f))(v.value1));
            };
            if (v instanceof Skip) {
                return new Skip(Data_Functor.map(Data_Lazy.functorLazy)(mapMaybe(dictFunctor)(f))(v.value0));
            };
            if (v instanceof Done) {
                return Done.value;
            };
            throw new Error("Failed pattern match at Control.Monad.List.Trans (line 183, column 3 - line 183, column 72): " + [ v.constructor.name ]);
        };
        return stepMap(dictFunctor)(g);
    };
};

// | Generate an infinite list by iterating a function.
var iterate = function (dictMonad) {
    var pure = Control_Applicative.pure(dictMonad.Applicative0());
    return function (f) {
        return function (a) {
            var g = function (x) {
                return pure(new Data_Maybe.Just(new Data_Tuple.Tuple(f(x), x)));
            };
            return unfold(dictMonad)(g)(a);
        };
    };
};

// | Generate an infinite list by repeating a value.
var repeat = function (dictMonad) {
    return iterate(dictMonad)(identity);
};

// | Extract the first element of a list.
var head = function (dictMonad) {
    var Functor0 = ((dictMonad.Bind1()).Apply0()).Functor0();
    return function (l) {
        return Data_Functor.map(Functor0)(Data_Functor.map(Data_Maybe.functorMaybe)(Data_Tuple.fst))(uncons(dictMonad)(l));
    };
};
var functorListT = function (dictFunctor) {
    return {
        map: function (f) {
            var g = function (v) {
                if (v instanceof Yield) {
                    return new Yield(f(v.value0), Data_Functor.map(Data_Lazy.functorLazy)(Data_Functor.map(functorListT(dictFunctor))(f))(v.value1));
                };
                if (v instanceof Skip) {
                    return new Skip(Data_Functor.map(Data_Lazy.functorLazy)(Data_Functor.map(functorListT(dictFunctor))(f))(v.value0));
                };
                if (v instanceof Done) {
                    return Done.value;
                };
                throw new Error("Failed pattern match at Control.Monad.List.Trans (line 280, column 5 - line 280, column 48): " + [ v.constructor.name ]);
            };
            return stepMap(dictFunctor)(g);
        }
    };
};

// | Lift a computation from the base functor.
var fromEffect = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (fa) {
        return Data_Functor.map(Functor0)(Data_Function.flip(Yield.create)(Data_Lazy.defer(function (v) {
            return nil(dictApplicative);
        })))(fa);
    };
};
var monadTransListT = {
    lift: function (dictMonad) {
        return fromEffect(dictMonad.Applicative0());
    }
};
var lift = /* #__PURE__ */ Control_Monad_Trans_Class.lift(monadTransListT);

// | Fold a list from the left, accumulating the result (effectfully) using the specified function.
// | Uses tail call optimization.
var foldlRec$prime = function (dictMonadRec) {
    var Monad0 = dictMonadRec.Monad0();
    var Applicative0 = Monad0.Applicative0();
    var Bind1 = Monad0.Bind1();
    var Monad01 = dictMonadRec.Monad0();
    return function (f) {
        var loop = function (b) {
            return function (l) {
                var g = function (v) {
                    if (v instanceof Data_Maybe.Nothing) {
                        return Control_Applicative.pure(Applicative0)(new Control_Monad_Rec_Class.Done(b));
                    };
                    if (v instanceof Data_Maybe.Just) {
                        return Control_Bind.bind(Bind1)(f(b)(v.value0.value0))(function (b$prime) {
                            return Control_Applicative.pure(Applicative0)(new Control_Monad_Rec_Class.Loop({
                                a: b$prime,
                                b: v.value0.value1
                            }));
                        });
                    };
                    throw new Error("Failed pattern match at Control.Monad.List.Trans (line 221, column 5 - line 221, column 45): " + [ v.constructor.name ]);
                };
                return Control_Bind.bind(Bind1)(uncons(Monad01)(l))(g);
            };
        };
        return Control_Monad_Rec_Class.tailRecM2(dictMonadRec)(loop);
    };
};

// | Drain a `ListT`, running it to completion and discarding all values.
// | Stack safe: Uses tail call optimization.
var runListTRec = function (dictMonadRec) {
    var Applicative0 = (dictMonadRec.Monad0()).Applicative0();
    return foldlRec$prime(dictMonadRec)(function (v) {
        return function (v1) {
            return Control_Applicative.pure(Applicative0)(Data_Unit.unit);
        };
    })(Data_Unit.unit);
};

// | Fold a list from the left, accumulating the result using the specified function.
// | Uses tail call optimization.
var foldlRec = function (dictMonadRec) {
    var Monad0 = dictMonadRec.Monad0();
    var Applicative0 = Monad0.Applicative0();
    var Bind1 = Monad0.Bind1();
    var Monad01 = dictMonadRec.Monad0();
    return function (f) {
        var loop = function (b) {
            return function (l) {
                var g = function (v) {
                    if (v instanceof Data_Maybe.Nothing) {
                        return Control_Applicative.pure(Applicative0)(new Control_Monad_Rec_Class.Done(b));
                    };
                    if (v instanceof Data_Maybe.Just) {
                        return Control_Applicative.pure(Applicative0)(new Control_Monad_Rec_Class.Loop({
                            a: f(b)(v.value0.value0),
                            b: v.value0.value1
                        }));
                    };
                    throw new Error("Failed pattern match at Control.Monad.List.Trans (line 239, column 7 - line 239, column 47): " + [ v.constructor.name ]);
                };
                return Control_Bind.bind(Bind1)(uncons(Monad01)(l))(g);
            };
        };
        return Control_Monad_Rec_Class.tailRecM2(dictMonadRec)(loop);
    };
};

// | Fold a list from the left, accumulating the result (effectfully) using the specified function.
var foldl$prime = function (dictMonad) {
    var Applicative0 = dictMonad.Applicative0();
    var Bind1 = dictMonad.Bind1();
    return function (f) {
        var loop = function (b) {
            return function (l) {
                var g = function (v) {
                    if (v instanceof Data_Maybe.Nothing) {
                        return Control_Applicative.pure(Applicative0)(b);
                    };
                    if (v instanceof Data_Maybe.Just) {
                        return Control_Bind.bind(Bind1)(f(b)(v.value0.value0))(Data_Function.flip(loop)(v.value0.value1));
                    };
                    throw new Error("Failed pattern match at Control.Monad.List.Trans (line 212, column 5 - line 212, column 35): " + [ v.constructor.name ]);
                };
                return Control_Bind.bind(Bind1)(uncons(dictMonad)(l))(g);
            };
        };
        return loop;
    };
};

// | Drain a `ListT`, running it to completion and discarding all values.
var runListT = function (dictMonad) {
    var Applicative0 = dictMonad.Applicative0();
    return foldl$prime(dictMonad)(function (v) {
        return function (v1) {
            return Control_Applicative.pure(Applicative0)(Data_Unit.unit);
        };
    })(Data_Unit.unit);
};

// | Fold a list from the left, accumulating the result using the specified function.
var foldl = function (dictMonad) {
    var Applicative0 = dictMonad.Applicative0();
    var Bind1 = dictMonad.Bind1();
    return function (f) {
        var loop = function (b) {
            return function (l) {
                var g = function (v) {
                    if (v instanceof Data_Maybe.Nothing) {
                        return Control_Applicative.pure(Applicative0)(b);
                    };
                    if (v instanceof Data_Maybe.Just) {
                        return loop(f(b)(v.value0.value0))(v.value0.value1);
                    };
                    throw new Error("Failed pattern match at Control.Monad.List.Trans (line 229, column 5 - line 229, column 35): " + [ v.constructor.name ]);
                };
                return Control_Bind.bind(Bind1)(uncons(dictMonad)(l))(g);
            };
        };
        return loop;
    };
};

// | Remove elements from a list for which a predicate fails to hold.
var filter = function (dictFunctor) {
    return function (f) {
        var g = function (v) {
            if (v instanceof Yield) {
                var s$prime = Data_Functor.map(Data_Lazy.functorLazy)(filter(dictFunctor)(f))(v.value1);
                var $246 = f(v.value0);
                if ($246) {
                    return new Yield(v.value0, s$prime);
                };
                return new Skip(s$prime);
            };
            if (v instanceof Skip) {
                var s$prime = Data_Functor.map(Data_Lazy.functorLazy)(filter(dictFunctor)(f))(v.value0);
                return new Skip(s$prime);
            };
            if (v instanceof Done) {
                return Done.value;
            };
            throw new Error("Failed pattern match at Control.Monad.List.Trans (line 176, column 3 - line 176, column 80): " + [ v.constructor.name ]);
        };
        return stepMap(dictFunctor)(g);
    };
};

// | Drop elements from the front of a list while a predicate holds.
var dropWhile = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (f) {
        var g = function (v) {
            if (v instanceof Yield) {
                var $251 = f(v.value0);
                if ($251) {
                    return new Skip(Data_Functor.map(Data_Lazy.functorLazy)(dropWhile(dictApplicative)(f))(v.value1));
                };
                return new Yield(v.value0, v.value1);
            };
            if (v instanceof Skip) {
                return new Skip(Data_Functor.map(Data_Lazy.functorLazy)(dropWhile(dictApplicative)(f))(v.value0));
            };
            if (v instanceof Done) {
                return Done.value;
            };
            throw new Error("Failed pattern match at Control.Monad.List.Trans (line 169, column 3 - line 169, column 70): " + [ v.constructor.name ]);
        };
        return stepMap(Functor0)(g);
    };
};

// | Drop a number of elements from the front of a list.
var drop = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (v) {
        return function (v1) {
            if (v === 0) {
                return v1;
            };
            var f = function (v2) {
                if (v2 instanceof Yield) {
                    return new Skip(Data_Functor.map(Data_Lazy.functorLazy)(drop(dictApplicative)(v - 1 | 0))(v2.value1));
                };
                if (v2 instanceof Skip) {
                    return new Skip(Data_Functor.map(Data_Lazy.functorLazy)(drop(dictApplicative)(v))(v2.value0));
                };
                if (v2 instanceof Done) {
                    return Done.value;
                };
                throw new Error("Failed pattern match at Control.Monad.List.Trans (line 162, column 3 - line 162, column 44): " + [ v2.constructor.name ]);
            };
            return stepMap(Functor0)(f)(v1);
        };
    };
};

// | Attach an element to the front of a list.
var cons = function (dictApplicative) {
    var pure = Control_Applicative.pure(dictApplicative);
    return function (lh) {
        return function (t) {
            return pure(new Yield(Data_Lazy.force(lh), t));
        };
    };
};
var unfoldable1ListT = function (dictMonad) {
    var Applicative0 = dictMonad.Applicative0();
    return {
        unfoldr1: function (f) {
            return function (b) {
                var go = function (v) {
                    if (v.value1 instanceof Data_Maybe.Nothing) {
                        return singleton(Applicative0)(v.value0);
                    };
                    if (v.value1 instanceof Data_Maybe.Just) {
                        return cons(Applicative0)(Control_Applicative.pure(Data_Lazy.applicativeLazy)(v.value0))(Data_Lazy.defer(function (v1) {
                            return go(f(v.value1.value0));
                        }));
                    };
                    throw new Error("Failed pattern match at Control.Monad.List.Trans (line 294, column 12 - line 296, column 67): " + [ v.constructor.name ]);
                };
                return go(f(b));
            };
        }
    };
};
var unfoldableListT = function (dictMonad) {
    var Applicative0 = dictMonad.Applicative0();
    var unfoldable1ListT1 = unfoldable1ListT(dictMonad);
    return {
        unfoldr: function (f) {
            return function (b) {
                var go = function (v) {
                    if (v instanceof Data_Maybe.Nothing) {
                        return nil(Applicative0);
                    };
                    if (v instanceof Data_Maybe.Just) {
                        return cons(Applicative0)(Control_Applicative.pure(Data_Lazy.applicativeLazy)(v.value0.value0))(Data_Lazy.defer(function (v1) {
                            return go(f(v.value0.value1));
                        }));
                    };
                    throw new Error("Failed pattern match at Control.Monad.List.Trans (line 287, column 12 - line 289, column 67): " + [ v.constructor.name ]);
                };
                return go(f(b));
            };
        },
        Unfoldable10: function () {
            return unfoldable1ListT1;
        }
    };
};
var semigroupListT = function (dictApplicative) {
    return {
        append: concat(dictApplicative)
    };
};

// | Append one list to another.
var concat = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (x) {
        return function (y) {
            var f = function (v) {
                if (v instanceof Yield) {
                    return new Yield(v.value0, Data_Functor.map(Data_Lazy.functorLazy)(function (v1) {
                        return Data_Semigroup.append(semigroupListT(dictApplicative))(v1)(y);
                    })(v.value1));
                };
                if (v instanceof Skip) {
                    return new Skip(Data_Functor.map(Data_Lazy.functorLazy)(function (v1) {
                        return Data_Semigroup.append(semigroupListT(dictApplicative))(v1)(y);
                    })(v.value0));
                };
                if (v instanceof Done) {
                    return new Skip(Data_Lazy.defer(Data_Function["const"](y)));
                };
                throw new Error("Failed pattern match at Control.Monad.List.Trans (line 106, column 3 - line 106, column 43): " + [ v.constructor.name ]);
            };
            return stepMap(Functor0)(f)(x);
        };
    };
};
var monoidListT = function (dictApplicative) {
    var semigroupListT1 = semigroupListT(dictApplicative);
    return {
        mempty: nil(dictApplicative),
        Semigroup0: function () {
            return semigroupListT1;
        }
    };
};

// | Remove elements from a list which do not contain a value.
var catMaybes = function (dictFunctor) {
    return mapMaybe(dictFunctor)(identity1);
};
var monadListT = function (dictMonad) {
    return {
        Applicative0: function () {
            return applicativeListT(dictMonad);
        },
        Bind1: function () {
            return bindListT(dictMonad);
        }
    };
};
var bindListT = function (dictMonad) {
    var semigroupListT1 = semigroupListT(dictMonad.Applicative0());
    var Functor0 = ((dictMonad.Bind1()).Apply0()).Functor0();
    return {
        bind: function (fa) {
            return function (f) {
                var g = function (v) {
                    if (v instanceof Yield) {
                        var h = function (s$prime) {
                            return Data_Semigroup.append(semigroupListT1)(f(v.value0))(Control_Bind.bind(bindListT(dictMonad))(s$prime)(f));
                        };
                        return new Skip(Data_Functor.map(Data_Lazy.functorLazy)(h)(v.value1));
                    };
                    if (v instanceof Skip) {
                        return new Skip(Data_Functor.map(Data_Lazy.functorLazy)(function (v1) {
                            return Control_Bind.bind(bindListT(dictMonad))(v1)(f);
                        })(v.value0));
                    };
                    if (v instanceof Done) {
                        return Done.value;
                    };
                    throw new Error("Failed pattern match at Control.Monad.List.Trans (line 306, column 5 - line 308, column 31): " + [ v.constructor.name ]);
                };
                return stepMap(Functor0)(g)(fa);
            };
        },
        Apply0: function () {
            return applyListT(dictMonad);
        }
    };
};
var applyListT = function (dictMonad) {
    var functorListT1 = functorListT(((dictMonad.Bind1()).Apply0()).Functor0());
    return {
        apply: Control_Monad.ap(monadListT(dictMonad)),
        Functor0: function () {
            return functorListT1;
        }
    };
};
var applicativeListT = function (dictMonad) {
    return {
        pure: singleton(dictMonad.Applicative0()),
        Apply0: function () {
            return applyListT(dictMonad);
        }
    };
};
var monadEffectListT = function (dictMonadEffect) {
    var Monad0 = dictMonadEffect.Monad0();
    var monadListT1 = monadListT(Monad0);
    return {
        liftEffect: (function () {
            var $280 = lift(Monad0);
            var $281 = Effect_Class.liftEffect(dictMonadEffect);
            return function ($282) {
                return $280($281($282));
            };
        })(),
        Monad0: function () {
            return monadListT1;
        }
    };
};
var monadSTListT = function (dictMonadST) {
    var Monad0 = dictMonadST.Monad0();
    var monadListT1 = monadListT(Monad0);
    return {
        liftST: (function () {
            var $283 = lift(Monad0);
            var $284 = Control_Monad_ST_Class.liftST(dictMonadST);
            return function ($285) {
                return $283($284($285));
            };
        })(),
        Monad0: function () {
            return monadListT1;
        }
    };
};
var altListT = function (dictApplicative) {
    var functorListT1 = functorListT((dictApplicative.Apply0()).Functor0());
    return {
        alt: concat(dictApplicative),
        Functor0: function () {
            return functorListT1;
        }
    };
};
var plusListT = function (dictMonad) {
    var Applicative0 = dictMonad.Applicative0();
    var altListT1 = altListT(Applicative0);
    return {
        empty: nil(Applicative0),
        Alt0: function () {
            return altListT1;
        }
    };
};
var alternativeListT = function (dictMonad) {
    var applicativeListT1 = applicativeListT(dictMonad);
    var plusListT1 = plusListT(dictMonad);
    return {
        Applicative0: function () {
            return applicativeListT1;
        },
        Plus1: function () {
            return plusListT1;
        }
    };
};
var monadPlusListT = function (dictMonad) {
    var monadListT1 = monadListT(dictMonad);
    var alternativeListT1 = alternativeListT(dictMonad);
    return {
        Monad0: function () {
            return monadListT1;
        },
        Alternative1: function () {
            return alternativeListT1;
        }
    };
};
export {
    ListT,
    Yield,
    Skip,
    Done,
    catMaybes,
    cons,
    drop,
    dropWhile,
    filter,
    foldl,
    foldlRec,
    foldl$prime,
    foldlRec$prime,
    fromEffect,
    head,
    iterate,
    mapMaybe,
    nil,
    prepend,
    prepend$prime,
    repeat,
    runListT,
    runListTRec,
    scanl,
    singleton,
    tail,
    take,
    takeWhile,
    uncons,
    unfold,
    wrapEffect,
    wrapLazy,
    zipWith,
    zipWith$prime,
    newtypeListT,
    semigroupListT,
    monoidListT,
    functorListT,
    unfoldableListT,
    unfoldable1ListT,
    applyListT,
    applicativeListT,
    bindListT,
    monadListT,
    monadTransListT,
    altListT,
    plusListT,
    alternativeListT,
    monadPlusListT,
    monadEffectListT,
    monadSTListT
};
export {
    lift
} from "../Control.Monad.Trans.Class/index.js";
