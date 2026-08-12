import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor_Coproduct from "../Data.Functor.Coproduct/index.js";
import * as Data_FunctorWithIndex from "../Data.FunctorWithIndex/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Monoid_Conj from "../Data.Monoid.Conj/index.js";
import * as Data_Monoid_Disj from "../Data.Monoid.Disj/index.js";
import * as Data_Monoid_Dual from "../Data.Monoid.Dual/index.js";
import * as Data_Monoid_Endo from "../Data.Monoid.Endo/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
var foldr = /* #__PURE__ */ Data_Foldable.foldr(Data_Foldable.foldableMultiplicative);
var foldl = /* #__PURE__ */ Data_Foldable.foldl(Data_Foldable.foldableMultiplicative);
var foldMap = /* #__PURE__ */ Data_Foldable.foldMap(Data_Foldable.foldableMultiplicative);
var foldr1 = /* #__PURE__ */ Data_Foldable.foldr(Data_Foldable.foldableMaybe);
var foldl1 = /* #__PURE__ */ Data_Foldable.foldl(Data_Foldable.foldableMaybe);
var foldMap1 = /* #__PURE__ */ Data_Foldable.foldMap(Data_Foldable.foldableMaybe);
var foldr2 = /* #__PURE__ */ Data_Foldable.foldr(Data_Foldable.foldableLast);
var foldl2 = /* #__PURE__ */ Data_Foldable.foldl(Data_Foldable.foldableLast);
var foldMap2 = /* #__PURE__ */ Data_Foldable.foldMap(Data_Foldable.foldableLast);
var foldr3 = /* #__PURE__ */ Data_Foldable.foldr(Data_Foldable.foldableFirst);
var foldl3 = /* #__PURE__ */ Data_Foldable.foldl(Data_Foldable.foldableFirst);
var foldMap3 = /* #__PURE__ */ Data_Foldable.foldMap(Data_Foldable.foldableFirst);
var foldr4 = /* #__PURE__ */ Data_Foldable.foldr(Data_Foldable.foldableDual);
var foldl4 = /* #__PURE__ */ Data_Foldable.foldl(Data_Foldable.foldableDual);
var foldMap4 = /* #__PURE__ */ Data_Foldable.foldMap(Data_Foldable.foldableDual);
var foldr5 = /* #__PURE__ */ Data_Foldable.foldr(Data_Foldable.foldableDisj);
var foldl5 = /* #__PURE__ */ Data_Foldable.foldl(Data_Foldable.foldableDisj);
var foldMap5 = /* #__PURE__ */ Data_Foldable.foldMap(Data_Foldable.foldableDisj);
var foldr6 = /* #__PURE__ */ Data_Foldable.foldr(Data_Foldable.foldableConj);
var foldl6 = /* #__PURE__ */ Data_Foldable.foldl(Data_Foldable.foldableConj);
var foldMap6 = /* #__PURE__ */ Data_Foldable.foldMap(Data_Foldable.foldableConj);
var foldr7 = /* #__PURE__ */ Data_Foldable.foldr(Data_Foldable.foldableAdditive);
var foldl7 = /* #__PURE__ */ Data_Foldable.foldl(Data_Foldable.foldableAdditive);
var foldMap7 = /* #__PURE__ */ Data_Foldable.foldMap(Data_Foldable.foldableAdditive);
var monoidDual = /* #__PURE__ */ Data_Monoid_Dual.monoidDual(/* #__PURE__ */ Data_Monoid_Endo.monoidEndo(Control_Category.categoryFn));
var monoidEndo = /* #__PURE__ */ Data_Monoid_Endo.monoidEndo(Control_Category.categoryFn);
var monoidEndo1 = /* #__PURE__ */ Data_Monoid_Endo.monoidEndo(Control_Category.categoryFn);
var unwrap = /* #__PURE__ */ Data_Newtype.unwrap();
var foldrWithIndex = function (dict) {
    return dict.foldrWithIndex;
};

// | Traverse a data structure with access to the index, performing some
// | effects encoded by an `Applicative` functor at each value, ignoring the
// | final result.
// |
// | For example:
// |
// | ```purescript
// | > traverseWithIndex_ (curry logShow) ["a", "b", "c"]
// | (Tuple 0 "a")
// | (Tuple 1 "b")
// | (Tuple 2 "c")
// | ```
var traverseWithIndex_ = function (dictApplicative) {
    var applySecond = Control_Apply.applySecond(dictApplicative.Apply0());
    return function (dictFoldableWithIndex) {
        return function (f) {
            return foldrWithIndex(dictFoldableWithIndex)(function (i) {
                var $235 = f(i);
                return function ($236) {
                    return applySecond($235($236));
                };
            })(Control_Applicative.pure(dictApplicative)(Data_Unit.unit));
        };
    };
};

// | A version of `traverseWithIndex_` with its arguments flipped.
// |
// | This can be useful when running an action written using do notation
// | for every element in a data structure:
// |
// | For example:
// |
// | ```purescript
// | forWithIndex_ ["a", "b", "c"] \i x -> do
// |   logShow i
// |   log x
// | ```
var forWithIndex_ = function (dictApplicative) {
    var traverseWithIndex_1 = traverseWithIndex_(dictApplicative);
    return function (dictFoldableWithIndex) {
        return Data_Function.flip(traverseWithIndex_1(dictFoldableWithIndex));
    };
};

// | A default implementation of `foldr` using `foldrWithIndex`
var foldrDefault = function (dictFoldableWithIndex) {
    return function (f) {
        return foldrWithIndex(dictFoldableWithIndex)(Data_Function["const"](f));
    };
};
var foldlWithIndex = function (dict) {
    return dict.foldlWithIndex;
};

// | A default implementation of `foldl` using `foldlWithIndex`
var foldlDefault = function (dictFoldableWithIndex) {
    return function (f) {
        return foldlWithIndex(dictFoldableWithIndex)(Data_Function["const"](f));
    };
};
var foldableWithIndexTuple = {
    foldrWithIndex: function (f) {
        return function (z) {
            return function (v) {
                return f(Data_Unit.unit)(v.value1)(z);
            };
        };
    },
    foldlWithIndex: function (f) {
        return function (z) {
            return function (v) {
                return f(Data_Unit.unit)(z)(v.value1);
            };
        };
    },
    foldMapWithIndex: function (dictMonoid) {
        return function (f) {
            return function (v) {
                return f(Data_Unit.unit)(v.value1);
            };
        };
    },
    Foldable0: function () {
        return Data_Foldable.foldableTuple;
    }
};
var foldableWithIndexMultiplicative = {
    foldrWithIndex: function (f) {
        return foldr(f(Data_Unit.unit));
    },
    foldlWithIndex: function (f) {
        return foldl(f(Data_Unit.unit));
    },
    foldMapWithIndex: function (dictMonoid) {
        var foldMap8 = foldMap(dictMonoid);
        return function (f) {
            return foldMap8(f(Data_Unit.unit));
        };
    },
    Foldable0: function () {
        return Data_Foldable.foldableMultiplicative;
    }
};
var foldableWithIndexMaybe = {
    foldrWithIndex: function (f) {
        return foldr1(f(Data_Unit.unit));
    },
    foldlWithIndex: function (f) {
        return foldl1(f(Data_Unit.unit));
    },
    foldMapWithIndex: function (dictMonoid) {
        var foldMap8 = foldMap1(dictMonoid);
        return function (f) {
            return foldMap8(f(Data_Unit.unit));
        };
    },
    Foldable0: function () {
        return Data_Foldable.foldableMaybe;
    }
};
var foldableWithIndexLast = {
    foldrWithIndex: function (f) {
        return foldr2(f(Data_Unit.unit));
    },
    foldlWithIndex: function (f) {
        return foldl2(f(Data_Unit.unit));
    },
    foldMapWithIndex: function (dictMonoid) {
        var foldMap8 = foldMap2(dictMonoid);
        return function (f) {
            return foldMap8(f(Data_Unit.unit));
        };
    },
    Foldable0: function () {
        return Data_Foldable.foldableLast;
    }
};
var foldableWithIndexIdentity = {
    foldrWithIndex: function (f) {
        return function (z) {
            return function (v) {
                return f(Data_Unit.unit)(v)(z);
            };
        };
    },
    foldlWithIndex: function (f) {
        return function (z) {
            return function (v) {
                return f(Data_Unit.unit)(z)(v);
            };
        };
    },
    foldMapWithIndex: function (dictMonoid) {
        return function (f) {
            return function (v) {
                return f(Data_Unit.unit)(v);
            };
        };
    },
    Foldable0: function () {
        return Data_Foldable.foldableIdentity;
    }
};
var foldableWithIndexFirst = {
    foldrWithIndex: function (f) {
        return foldr3(f(Data_Unit.unit));
    },
    foldlWithIndex: function (f) {
        return foldl3(f(Data_Unit.unit));
    },
    foldMapWithIndex: function (dictMonoid) {
        var foldMap8 = foldMap3(dictMonoid);
        return function (f) {
            return foldMap8(f(Data_Unit.unit));
        };
    },
    Foldable0: function () {
        return Data_Foldable.foldableFirst;
    }
};
var foldableWithIndexEither = {
    foldrWithIndex: function (v) {
        return function (v1) {
            return function (v2) {
                if (v2 instanceof Data_Either.Left) {
                    return v1;
                };
                if (v2 instanceof Data_Either.Right) {
                    return v(Data_Unit.unit)(v2.value0)(v1);
                };
                throw new Error("Failed pattern match at Data.FoldableWithIndex (line 164, column 1 - line 170, column 42): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
            };
        };
    },
    foldlWithIndex: function (v) {
        return function (v1) {
            return function (v2) {
                if (v2 instanceof Data_Either.Left) {
                    return v1;
                };
                if (v2 instanceof Data_Either.Right) {
                    return v(Data_Unit.unit)(v1)(v2.value0);
                };
                throw new Error("Failed pattern match at Data.FoldableWithIndex (line 164, column 1 - line 170, column 42): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
            };
        };
    },
    foldMapWithIndex: function (dictMonoid) {
        var mempty = Data_Monoid.mempty(dictMonoid);
        return function (v) {
            return function (v1) {
                if (v1 instanceof Data_Either.Left) {
                    return mempty;
                };
                if (v1 instanceof Data_Either.Right) {
                    return v(Data_Unit.unit)(v1.value0);
                };
                throw new Error("Failed pattern match at Data.FoldableWithIndex (line 164, column 1 - line 170, column 42): " + [ v.constructor.name, v1.constructor.name ]);
            };
        };
    },
    Foldable0: function () {
        return Data_Foldable.foldableEither;
    }
};
var foldableWithIndexDual = {
    foldrWithIndex: function (f) {
        return foldr4(f(Data_Unit.unit));
    },
    foldlWithIndex: function (f) {
        return foldl4(f(Data_Unit.unit));
    },
    foldMapWithIndex: function (dictMonoid) {
        var foldMap8 = foldMap4(dictMonoid);
        return function (f) {
            return foldMap8(f(Data_Unit.unit));
        };
    },
    Foldable0: function () {
        return Data_Foldable.foldableDual;
    }
};
var foldableWithIndexDisj = {
    foldrWithIndex: function (f) {
        return foldr5(f(Data_Unit.unit));
    },
    foldlWithIndex: function (f) {
        return foldl5(f(Data_Unit.unit));
    },
    foldMapWithIndex: function (dictMonoid) {
        var foldMap8 = foldMap5(dictMonoid);
        return function (f) {
            return foldMap8(f(Data_Unit.unit));
        };
    },
    Foldable0: function () {
        return Data_Foldable.foldableDisj;
    }
};
var foldableWithIndexConst = {
    foldrWithIndex: function (v) {
        return function (z) {
            return function (v1) {
                return z;
            };
        };
    },
    foldlWithIndex: function (v) {
        return function (z) {
            return function (v1) {
                return z;
            };
        };
    },
    foldMapWithIndex: function (dictMonoid) {
        var mempty = Data_Monoid.mempty(dictMonoid);
        return function (v) {
            return function (v1) {
                return mempty;
            };
        };
    },
    Foldable0: function () {
        return Data_Foldable.foldableConst;
    }
};
var foldableWithIndexConj = {
    foldrWithIndex: function (f) {
        return foldr6(f(Data_Unit.unit));
    },
    foldlWithIndex: function (f) {
        return foldl6(f(Data_Unit.unit));
    },
    foldMapWithIndex: function (dictMonoid) {
        var foldMap8 = foldMap6(dictMonoid);
        return function (f) {
            return foldMap8(f(Data_Unit.unit));
        };
    },
    Foldable0: function () {
        return Data_Foldable.foldableConj;
    }
};
var foldableWithIndexAdditive = {
    foldrWithIndex: function (f) {
        return foldr7(f(Data_Unit.unit));
    },
    foldlWithIndex: function (f) {
        return foldl7(f(Data_Unit.unit));
    },
    foldMapWithIndex: function (dictMonoid) {
        var foldMap8 = foldMap7(dictMonoid);
        return function (f) {
            return foldMap8(f(Data_Unit.unit));
        };
    },
    Foldable0: function () {
        return Data_Foldable.foldableAdditive;
    }
};

// | Similar to 'foldlWithIndex', but the result is encapsulated in a monad.
// |
// | Note: this function is not generally stack-safe, e.g., for monads which
// | build up thunks a la `Eff`.
var foldWithIndexM = function (dictFoldableWithIndex) {
    return function (dictMonad) {
        var Bind1 = dictMonad.Bind1();
        var Applicative0 = dictMonad.Applicative0();
        return function (f) {
            return function (a0) {
                return foldlWithIndex(dictFoldableWithIndex)(function (i) {
                    return function (ma) {
                        return function (b) {
                            return Control_Bind.bind(Bind1)(ma)(Data_Function.flip(f(i))(b));
                        };
                    };
                })(Control_Applicative.pure(Applicative0)(a0));
            };
        };
    };
};

// | A default implementation of `foldMapWithIndex` using `foldrWithIndex`.
// |
// | Note: when defining a `FoldableWithIndex` instance, this function is
// | unsafe to use in combination with `foldrWithIndexDefault`.
var foldMapWithIndexDefaultR = function (dictFoldableWithIndex) {
    return function (dictMonoid) {
        var Semigroup0 = dictMonoid.Semigroup0();
        var mempty = Data_Monoid.mempty(dictMonoid);
        return function (f) {
            return foldrWithIndex(dictFoldableWithIndex)(function (i) {
                return function (x) {
                    return function (acc) {
                        return Data_Semigroup.append(Semigroup0)(f(i)(x))(acc);
                    };
                };
            })(mempty);
        };
    };
};
var foldableWithIndexArray = {
    foldrWithIndex: function (f) {
        return function (z) {
            var $237 = Data_Foldable.foldr(Data_Foldable.foldableArray)(function (v) {
                return function (y) {
                    return f(v.value0)(v.value1)(y);
                };
            })(z);
            var $238 = Data_FunctorWithIndex.mapWithIndex(Data_FunctorWithIndex.functorWithIndexArray)(Data_Tuple.Tuple.create);
            return function ($239) {
                return $237($238($239));
            };
        };
    },
    foldlWithIndex: function (f) {
        return function (z) {
            var $240 = Data_Foldable.foldl(Data_Foldable.foldableArray)(function (y) {
                return function (v) {
                    return f(v.value0)(y)(v.value1);
                };
            })(z);
            var $241 = Data_FunctorWithIndex.mapWithIndex(Data_FunctorWithIndex.functorWithIndexArray)(Data_Tuple.Tuple.create);
            return function ($242) {
                return $240($241($242));
            };
        };
    },
    foldMapWithIndex: function (dictMonoid) {
        return foldMapWithIndexDefaultR(foldableWithIndexArray)(dictMonoid);
    },
    Foldable0: function () {
        return Data_Foldable.foldableArray;
    }
};

// | A default implementation of `foldMapWithIndex` using `foldlWithIndex`.
// |
// | Note: when defining a `FoldableWithIndex` instance, this function is
// | unsafe to use in combination with `foldlWithIndexDefault`.
var foldMapWithIndexDefaultL = function (dictFoldableWithIndex) {
    return function (dictMonoid) {
        var Semigroup0 = dictMonoid.Semigroup0();
        var mempty = Data_Monoid.mempty(dictMonoid);
        return function (f) {
            return foldlWithIndex(dictFoldableWithIndex)(function (i) {
                return function (acc) {
                    return function (x) {
                        return Data_Semigroup.append(Semigroup0)(acc)(f(i)(x));
                    };
                };
            })(mempty);
        };
    };
};
var foldMapWithIndex = function (dict) {
    return dict.foldMapWithIndex;
};
var foldableWithIndexApp = function (dictFoldableWithIndex) {
    var foldableApp = Data_Foldable.foldableApp(dictFoldableWithIndex.Foldable0());
    return {
        foldrWithIndex: function (f) {
            return function (z) {
                return function (v) {
                    return foldrWithIndex(dictFoldableWithIndex)(f)(z)(v);
                };
            };
        },
        foldlWithIndex: function (f) {
            return function (z) {
                return function (v) {
                    return foldlWithIndex(dictFoldableWithIndex)(f)(z)(v);
                };
            };
        },
        foldMapWithIndex: function (dictMonoid) {
            return function (f) {
                return function (v) {
                    return foldMapWithIndex(dictFoldableWithIndex)(dictMonoid)(f)(v);
                };
            };
        },
        Foldable0: function () {
            return foldableApp;
        }
    };
};
var foldableWithIndexCompose = function (dictFoldableWithIndex) {
    var foldableCompose = Data_Foldable.foldableCompose(dictFoldableWithIndex.Foldable0());
    return function (dictFoldableWithIndex1) {
        var foldlWithIndex1 = foldlWithIndex(dictFoldableWithIndex1);
        var foldMapWithIndex1 = foldMapWithIndex(dictFoldableWithIndex1);
        var foldableCompose1 = foldableCompose(dictFoldableWithIndex1.Foldable0());
        return {
            foldrWithIndex: function (f) {
                return function (i) {
                    return function (v) {
                        return foldrWithIndex(dictFoldableWithIndex)(function (a) {
                            return Data_Function.flip(foldrWithIndex(dictFoldableWithIndex1)(Data_Tuple.curry(f)(a)));
                        })(i)(v);
                    };
                };
            },
            foldlWithIndex: function (f) {
                return function (i) {
                    return function (v) {
                        return foldlWithIndex(dictFoldableWithIndex)((function () {
                            var $243 = Data_Tuple.curry(f);
                            return function ($244) {
                                return foldlWithIndex1($243($244));
                            };
                        })())(i)(v);
                    };
                };
            },
            foldMapWithIndex: function (dictMonoid) {
                var foldMapWithIndex2 = foldMapWithIndex1(dictMonoid);
                return function (f) {
                    return function (v) {
                        return foldMapWithIndex(dictFoldableWithIndex)(dictMonoid)((function () {
                            var $245 = Data_Tuple.curry(f);
                            return function ($246) {
                                return foldMapWithIndex2($245($246));
                            };
                        })())(v);
                    };
                };
            },
            Foldable0: function () {
                return foldableCompose1;
            }
        };
    };
};
var foldableWithIndexCoproduct = function (dictFoldableWithIndex) {
    var foldableCoproduct = Data_Foldable.foldableCoproduct(dictFoldableWithIndex.Foldable0());
    return function (dictFoldableWithIndex1) {
        var foldableCoproduct1 = foldableCoproduct(dictFoldableWithIndex1.Foldable0());
        return {
            foldrWithIndex: function (f) {
                return function (z) {
                    return Data_Functor_Coproduct.coproduct(foldrWithIndex(dictFoldableWithIndex)(function ($247) {
                        return f(Data_Either.Left.create($247));
                    })(z))(foldrWithIndex(dictFoldableWithIndex1)(function ($248) {
                        return f(Data_Either.Right.create($248));
                    })(z));
                };
            },
            foldlWithIndex: function (f) {
                return function (z) {
                    return Data_Functor_Coproduct.coproduct(foldlWithIndex(dictFoldableWithIndex)(function ($249) {
                        return f(Data_Either.Left.create($249));
                    })(z))(foldlWithIndex(dictFoldableWithIndex1)(function ($250) {
                        return f(Data_Either.Right.create($250));
                    })(z));
                };
            },
            foldMapWithIndex: function (dictMonoid) {
                return function (f) {
                    return Data_Functor_Coproduct.coproduct(foldMapWithIndex(dictFoldableWithIndex)(dictMonoid)(function ($251) {
                        return f(Data_Either.Left.create($251));
                    }))(foldMapWithIndex(dictFoldableWithIndex1)(dictMonoid)(function ($252) {
                        return f(Data_Either.Right.create($252));
                    }));
                };
            },
            Foldable0: function () {
                return foldableCoproduct1;
            }
        };
    };
};
var foldableWithIndexProduct = function (dictFoldableWithIndex) {
    var foldableProduct = Data_Foldable.foldableProduct(dictFoldableWithIndex.Foldable0());
    return function (dictFoldableWithIndex1) {
        var foldableProduct1 = foldableProduct(dictFoldableWithIndex1.Foldable0());
        return {
            foldrWithIndex: function (f) {
                return function (z) {
                    return function (v) {
                        return foldrWithIndex(dictFoldableWithIndex)(function ($253) {
                            return f(Data_Either.Left.create($253));
                        })(foldrWithIndex(dictFoldableWithIndex1)(function ($254) {
                            return f(Data_Either.Right.create($254));
                        })(z)(v.value1))(v.value0);
                    };
                };
            },
            foldlWithIndex: function (f) {
                return function (z) {
                    return function (v) {
                        return foldlWithIndex(dictFoldableWithIndex1)(function ($255) {
                            return f(Data_Either.Right.create($255));
                        })(foldlWithIndex(dictFoldableWithIndex)(function ($256) {
                            return f(Data_Either.Left.create($256));
                        })(z)(v.value0))(v.value1);
                    };
                };
            },
            foldMapWithIndex: function (dictMonoid) {
                var Semigroup0 = dictMonoid.Semigroup0();
                return function (f) {
                    return function (v) {
                        return Data_Semigroup.append(Semigroup0)(foldMapWithIndex(dictFoldableWithIndex)(dictMonoid)(function ($257) {
                            return f(Data_Either.Left.create($257));
                        })(v.value0))(foldMapWithIndex(dictFoldableWithIndex1)(dictMonoid)(function ($258) {
                            return f(Data_Either.Right.create($258));
                        })(v.value1));
                    };
                };
            },
            Foldable0: function () {
                return foldableProduct1;
            }
        };
    };
};

// | A default implementation of `foldlWithIndex` using `foldMapWithIndex`.
// |
// | Note: when defining a `FoldableWithIndex` instance, this function is
// | unsafe to use in combination with `foldMapWithIndexDefaultL`.
var foldlWithIndexDefault = function (dictFoldableWithIndex) {
    return function (c) {
        return function (u) {
            return function (xs) {
                return Data_Newtype.unwrap()(Data_Newtype.unwrap()(foldMapWithIndex(dictFoldableWithIndex)(monoidDual)(function (i) {
                    var $259 = Data_Function.flip(c(i));
                    return function ($260) {
                        return Data_Monoid_Dual.Dual(Data_Monoid_Endo.Endo($259($260)));
                    };
                })(xs)))(u);
            };
        };
    };
};

// | A default implementation of `foldrWithIndex` using `foldMapWithIndex`.
// |
// | Note: when defining a `FoldableWithIndex` instance, this function is
// | unsafe to use in combination with `foldMapWithIndexDefaultR`.
var foldrWithIndexDefault = function (dictFoldableWithIndex) {
    return function (c) {
        return function (u) {
            return function (xs) {
                return Data_Newtype.unwrap()(foldMapWithIndex(dictFoldableWithIndex)(monoidEndo)(function (i) {
                    var $261 = c(i);
                    return function ($262) {
                        return Data_Monoid_Endo.Endo($261($262));
                    };
                })(xs))(u);
            };
        };
    };
};

// | `foldMapWithIndex` but with each element surrounded by some fixed value.
// |
// | For example:
// |
// | ```purescript
// | > surroundMapWithIndex "*" (\i x -> show i <> x) []
// | = "*"
// |
// | > surroundMapWithIndex "*" (\i x -> show i <> x) ["a"]
// | = "*0a*"
// |
// | > surroundMapWithIndex "*" (\i x -> show i <> x) ["a", "b"]
// | = "*0a*1b*"
// |
// | > surroundMapWithIndex "*" (\i x -> show i <> x) ["a", "b", "c"]
// | = "*0a*1b*2c*"
// | ```
var surroundMapWithIndex = function (dictFoldableWithIndex) {
    return function (dictSemigroup) {
        return function (d) {
            return function (t) {
                return function (f) {
                    var joined = function (i) {
                        return function (a) {
                            return function (m) {
                                return Data_Semigroup.append(dictSemigroup)(d)(Data_Semigroup.append(dictSemigroup)(t(i)(a))(m));
                            };
                        };
                    };
                    return Data_Newtype.unwrap()(foldMapWithIndex(dictFoldableWithIndex)(monoidEndo1)(joined)(f))(d);
                };
            };
        };
    };
};

// | A default implementation of `foldMap` using `foldMapWithIndex`
var foldMapDefault = function (dictFoldableWithIndex) {
    return function (dictMonoid) {
        return function (f) {
            return foldMapWithIndex(dictFoldableWithIndex)(dictMonoid)(Data_Function["const"](f));
        };
    };
};

// | Try to find an element in a data structure which satisfies a predicate
// | with access to the index.
var findWithIndex = function (dictFoldableWithIndex) {
    return function (p) {
        var go = function (v) {
            return function (v1) {
                return function (v2) {
                    if (v1 instanceof Data_Maybe.Nothing && p(v)(v2)) {
                        return new Data_Maybe.Just({
                            index: v,
                            value: v2
                        });
                    };
                    return v1;
                };
            };
        };
        return foldlWithIndex(dictFoldableWithIndex)(go)(Data_Maybe.Nothing.value);
    };
};

// | Try to find an element in a data structure which satisfies a predicate mapping
// | with access to the index.
var findMapWithIndex = function (dictFoldableWithIndex) {
    return function (f) {
        var go = function (v) {
            return function (v1) {
                return function (v2) {
                    if (v1 instanceof Data_Maybe.Nothing) {
                        return f(v)(v2);
                    };
                    return v1;
                };
            };
        };
        return foldlWithIndex(dictFoldableWithIndex)(go)(Data_Maybe.Nothing.value);
    };
};

// | `anyWithIndex f` is the same as `or <<< mapWithIndex f`; map a function over the
// | structure, and then get the disjunction of the results.
var anyWithIndex = function (dictFoldableWithIndex) {
    return function (dictHeytingAlgebra) {
        var monoidDisj = Data_Monoid_Disj.monoidDisj(dictHeytingAlgebra);
        return function (t) {
            var $263 = foldMapWithIndex(dictFoldableWithIndex)(monoidDisj)(function (i) {
                var $265 = t(i);
                return function ($266) {
                    return Data_Monoid_Disj.Disj($265($266));
                };
            });
            return function ($264) {
                return unwrap($263($264));
            };
        };
    };
};

// | `allWithIndex f` is the same as `and <<< mapWithIndex f`; map a function over the
// | structure, and then get the conjunction of the results.
var allWithIndex = function (dictFoldableWithIndex) {
    return function (dictHeytingAlgebra) {
        var monoidConj = Data_Monoid_Conj.monoidConj(dictHeytingAlgebra);
        return function (t) {
            var $267 = foldMapWithIndex(dictFoldableWithIndex)(monoidConj)(function (i) {
                var $269 = t(i);
                return function ($270) {
                    return Data_Monoid_Conj.Conj($269($270));
                };
            });
            return function ($268) {
                return unwrap($267($268));
            };
        };
    };
};
export {
    foldrWithIndex,
    foldlWithIndex,
    foldMapWithIndex,
    foldrWithIndexDefault,
    foldlWithIndexDefault,
    foldMapWithIndexDefaultR,
    foldMapWithIndexDefaultL,
    foldWithIndexM,
    traverseWithIndex_,
    forWithIndex_,
    surroundMapWithIndex,
    allWithIndex,
    anyWithIndex,
    findWithIndex,
    findMapWithIndex,
    foldrDefault,
    foldlDefault,
    foldMapDefault,
    foldableWithIndexArray,
    foldableWithIndexMaybe,
    foldableWithIndexFirst,
    foldableWithIndexLast,
    foldableWithIndexAdditive,
    foldableWithIndexDual,
    foldableWithIndexDisj,
    foldableWithIndexConj,
    foldableWithIndexMultiplicative,
    foldableWithIndexEither,
    foldableWithIndexTuple,
    foldableWithIndexIdentity,
    foldableWithIndexConst,
    foldableWithIndexProduct,
    foldableWithIndexCoproduct,
    foldableWithIndexCompose,
    foldableWithIndexApp
};
