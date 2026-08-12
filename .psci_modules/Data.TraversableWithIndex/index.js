import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_FoldableWithIndex from "../Data.FoldableWithIndex/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Functor_App from "../Data.Functor.App/index.js";
import * as Data_Functor_Compose from "../Data.Functor.Compose/index.js";
import * as Data_Functor_Coproduct from "../Data.Functor.Coproduct/index.js";
import * as Data_Functor_Product from "../Data.Functor.Product/index.js";
import * as Data_FunctorWithIndex from "../Data.FunctorWithIndex/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Traversable_Accum from "../Data.Traversable.Accum/index.js";
import * as Data_Traversable_Accum_Internal from "../Data.Traversable.Accum.Internal/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
var traverse = /* #__PURE__ */ Data_Traversable.traverse(Data_Traversable.traversableMultiplicative);
var traverse1 = /* #__PURE__ */ Data_Traversable.traverse(Data_Traversable.traversableMaybe);
var traverse2 = /* #__PURE__ */ Data_Traversable.traverse(Data_Traversable.traversableLast);
var traverse3 = /* #__PURE__ */ Data_Traversable.traverse(Data_Traversable.traversableFirst);
var traverse4 = /* #__PURE__ */ Data_Traversable.traverse(Data_Traversable.traversableDual);
var traverse5 = /* #__PURE__ */ Data_Traversable.traverse(Data_Traversable.traversableDisj);
var traverse6 = /* #__PURE__ */ Data_Traversable.traverse(Data_Traversable.traversableConj);
var traverse7 = /* #__PURE__ */ Data_Traversable.traverse(Data_Traversable.traversableAdditive);

// | A default implementation of `traverseWithIndex` using `sequence` and `mapWithIndex`.
var traverseWithIndexDefault = function (dictTraversableWithIndex) {
    var sequence = Data_Traversable.sequence(dictTraversableWithIndex.Traversable2());
    var FunctorWithIndex0 = dictTraversableWithIndex.FunctorWithIndex0();
    return function (dictApplicative) {
        var sequence1 = sequence(dictApplicative);
        return function (f) {
            var $143 = Data_FunctorWithIndex.mapWithIndex(FunctorWithIndex0)(f);
            return function ($144) {
                return sequence1($143($144));
            };
        };
    };
};
var traverseWithIndex = function (dict) {
    return dict.traverseWithIndex;
};

// | A default implementation of `traverse` in terms of `traverseWithIndex`
var traverseDefault = function (dictTraversableWithIndex) {
    return function (dictApplicative) {
        return function (f) {
            return traverseWithIndex(dictTraversableWithIndex)(dictApplicative)(Data_Function["const"](f));
        };
    };
};
var traversableWithIndexTuple = {
    traverseWithIndex: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (f) {
            return function (v) {
                return Data_Functor.map(Functor0)(Data_Tuple.Tuple.create(v.value0))(f(Data_Unit.unit)(v.value1));
            };
        };
    },
    FunctorWithIndex0: function () {
        return Data_FunctorWithIndex.functorWithIndexTuple;
    },
    FoldableWithIndex1: function () {
        return Data_FoldableWithIndex.foldableWithIndexTuple;
    },
    Traversable2: function () {
        return Data_Traversable.traversableTuple;
    }
};
var traversableWithIndexProduct = function (dictTraversableWithIndex) {
    var functorWithIndexProduct = Data_FunctorWithIndex.functorWithIndexProduct(dictTraversableWithIndex.FunctorWithIndex0());
    var foldableWithIndexProduct = Data_FoldableWithIndex.foldableWithIndexProduct(dictTraversableWithIndex.FoldableWithIndex1());
    var traversableProduct = Data_Traversable.traversableProduct(dictTraversableWithIndex.Traversable2());
    return function (dictTraversableWithIndex1) {
        var functorWithIndexProduct1 = functorWithIndexProduct(dictTraversableWithIndex1.FunctorWithIndex0());
        var foldableWithIndexProduct1 = foldableWithIndexProduct(dictTraversableWithIndex1.FoldableWithIndex1());
        var traversableProduct1 = traversableProduct(dictTraversableWithIndex1.Traversable2());
        return {
            traverseWithIndex: function (dictApplicative) {
                var Apply0 = dictApplicative.Apply0();
                return function (f) {
                    return function (v) {
                        return Control_Apply.lift2(Apply0)(Data_Functor_Product.product)(traverseWithIndex(dictTraversableWithIndex)(dictApplicative)(function ($145) {
                            return f(Data_Either.Left.create($145));
                        })(v.value0))(traverseWithIndex(dictTraversableWithIndex1)(dictApplicative)(function ($146) {
                            return f(Data_Either.Right.create($146));
                        })(v.value1));
                    };
                };
            },
            FunctorWithIndex0: function () {
                return functorWithIndexProduct1;
            },
            FoldableWithIndex1: function () {
                return foldableWithIndexProduct1;
            },
            Traversable2: function () {
                return traversableProduct1;
            }
        };
    };
};
var traversableWithIndexMultiplicative = {
    traverseWithIndex: function (dictApplicative) {
        var traverse8 = traverse(dictApplicative);
        return function (f) {
            return traverse8(f(Data_Unit.unit));
        };
    },
    FunctorWithIndex0: function () {
        return Data_FunctorWithIndex.functorWithIndexMultiplicative;
    },
    FoldableWithIndex1: function () {
        return Data_FoldableWithIndex.foldableWithIndexMultiplicative;
    },
    Traversable2: function () {
        return Data_Traversable.traversableMultiplicative;
    }
};
var traversableWithIndexMaybe = {
    traverseWithIndex: function (dictApplicative) {
        var traverse8 = traverse1(dictApplicative);
        return function (f) {
            return traverse8(f(Data_Unit.unit));
        };
    },
    FunctorWithIndex0: function () {
        return Data_FunctorWithIndex.functorWithIndexMaybe;
    },
    FoldableWithIndex1: function () {
        return Data_FoldableWithIndex.foldableWithIndexMaybe;
    },
    Traversable2: function () {
        return Data_Traversable.traversableMaybe;
    }
};
var traversableWithIndexLast = {
    traverseWithIndex: function (dictApplicative) {
        var traverse8 = traverse2(dictApplicative);
        return function (f) {
            return traverse8(f(Data_Unit.unit));
        };
    },
    FunctorWithIndex0: function () {
        return Data_FunctorWithIndex.functorWithIndexLast;
    },
    FoldableWithIndex1: function () {
        return Data_FoldableWithIndex.foldableWithIndexLast;
    },
    Traversable2: function () {
        return Data_Traversable.traversableLast;
    }
};
var traversableWithIndexIdentity = {
    traverseWithIndex: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (f) {
            return function (v) {
                return Data_Functor.map(Functor0)(Data_Identity.Identity)(f(Data_Unit.unit)(v));
            };
        };
    },
    FunctorWithIndex0: function () {
        return Data_FunctorWithIndex.functorWithIndexIdentity;
    },
    FoldableWithIndex1: function () {
        return Data_FoldableWithIndex.foldableWithIndexIdentity;
    },
    Traversable2: function () {
        return Data_Traversable.traversableIdentity;
    }
};
var traversableWithIndexFirst = {
    traverseWithIndex: function (dictApplicative) {
        var traverse8 = traverse3(dictApplicative);
        return function (f) {
            return traverse8(f(Data_Unit.unit));
        };
    },
    FunctorWithIndex0: function () {
        return Data_FunctorWithIndex.functorWithIndexFirst;
    },
    FoldableWithIndex1: function () {
        return Data_FoldableWithIndex.foldableWithIndexFirst;
    },
    Traversable2: function () {
        return Data_Traversable.traversableFirst;
    }
};
var traversableWithIndexEither = {
    traverseWithIndex: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (v) {
            return function (v1) {
                if (v1 instanceof Data_Either.Left) {
                    return Control_Applicative.pure(dictApplicative)(new Data_Either.Left(v1.value0));
                };
                if (v1 instanceof Data_Either.Right) {
                    return Data_Functor.map(Functor0)(Data_Either.Right.create)(v(Data_Unit.unit)(v1.value0));
                };
                throw new Error("Failed pattern match at Data.TraversableWithIndex (line 95, column 1 - line 97, column 53): " + [ v.constructor.name, v1.constructor.name ]);
            };
        };
    },
    FunctorWithIndex0: function () {
        return Data_FunctorWithIndex.functorWithIndexEither;
    },
    FoldableWithIndex1: function () {
        return Data_FoldableWithIndex.foldableWithIndexEither;
    },
    Traversable2: function () {
        return Data_Traversable.traversableEither;
    }
};
var traversableWithIndexDual = {
    traverseWithIndex: function (dictApplicative) {
        var traverse8 = traverse4(dictApplicative);
        return function (f) {
            return traverse8(f(Data_Unit.unit));
        };
    },
    FunctorWithIndex0: function () {
        return Data_FunctorWithIndex.functorWithIndexDual;
    },
    FoldableWithIndex1: function () {
        return Data_FoldableWithIndex.foldableWithIndexDual;
    },
    Traversable2: function () {
        return Data_Traversable.traversableDual;
    }
};
var traversableWithIndexDisj = {
    traverseWithIndex: function (dictApplicative) {
        var traverse8 = traverse5(dictApplicative);
        return function (f) {
            return traverse8(f(Data_Unit.unit));
        };
    },
    FunctorWithIndex0: function () {
        return Data_FunctorWithIndex.functorWithIndexDisj;
    },
    FoldableWithIndex1: function () {
        return Data_FoldableWithIndex.foldableWithIndexDisj;
    },
    Traversable2: function () {
        return Data_Traversable.traversableDisj;
    }
};
var traversableWithIndexCoproduct = function (dictTraversableWithIndex) {
    var functorWithIndexCoproduct = Data_FunctorWithIndex.functorWithIndexCoproduct(dictTraversableWithIndex.FunctorWithIndex0());
    var foldableWithIndexCoproduct = Data_FoldableWithIndex.foldableWithIndexCoproduct(dictTraversableWithIndex.FoldableWithIndex1());
    var traversableCoproduct = Data_Traversable.traversableCoproduct(dictTraversableWithIndex.Traversable2());
    return function (dictTraversableWithIndex1) {
        var functorWithIndexCoproduct1 = functorWithIndexCoproduct(dictTraversableWithIndex1.FunctorWithIndex0());
        var foldableWithIndexCoproduct1 = foldableWithIndexCoproduct(dictTraversableWithIndex1.FoldableWithIndex1());
        var traversableCoproduct1 = traversableCoproduct(dictTraversableWithIndex1.Traversable2());
        return {
            traverseWithIndex: function (dictApplicative) {
                var Functor0 = (dictApplicative.Apply0()).Functor0();
                return function (f) {
                    return Data_Functor_Coproduct.coproduct((function () {
                        var $147 = Data_Functor.map(Functor0)(function ($150) {
                            return Data_Functor_Coproduct.Coproduct(Data_Either.Left.create($150));
                        });
                        var $148 = traverseWithIndex(dictTraversableWithIndex)(dictApplicative)(function ($151) {
                            return f(Data_Either.Left.create($151));
                        });
                        return function ($149) {
                            return $147($148($149));
                        };
                    })())((function () {
                        var $152 = Data_Functor.map(Functor0)(function ($155) {
                            return Data_Functor_Coproduct.Coproduct(Data_Either.Right.create($155));
                        });
                        var $153 = traverseWithIndex(dictTraversableWithIndex1)(dictApplicative)(function ($156) {
                            return f(Data_Either.Right.create($156));
                        });
                        return function ($154) {
                            return $152($153($154));
                        };
                    })());
                };
            },
            FunctorWithIndex0: function () {
                return functorWithIndexCoproduct1;
            },
            FoldableWithIndex1: function () {
                return foldableWithIndexCoproduct1;
            },
            Traversable2: function () {
                return traversableCoproduct1;
            }
        };
    };
};
var traversableWithIndexConst = {
    traverseWithIndex: function (dictApplicative) {
        return function (v) {
            return function (v1) {
                return Control_Applicative.pure(dictApplicative)(v1);
            };
        };
    },
    FunctorWithIndex0: function () {
        return Data_FunctorWithIndex.functorWithIndexConst;
    },
    FoldableWithIndex1: function () {
        return Data_FoldableWithIndex.foldableWithIndexConst;
    },
    Traversable2: function () {
        return Data_Traversable.traversableConst;
    }
};
var traversableWithIndexConj = {
    traverseWithIndex: function (dictApplicative) {
        var traverse8 = traverse6(dictApplicative);
        return function (f) {
            return traverse8(f(Data_Unit.unit));
        };
    },
    FunctorWithIndex0: function () {
        return Data_FunctorWithIndex.functorWithIndexConj;
    },
    FoldableWithIndex1: function () {
        return Data_FoldableWithIndex.foldableWithIndexConj;
    },
    Traversable2: function () {
        return Data_Traversable.traversableConj;
    }
};
var traversableWithIndexCompose = function (dictTraversableWithIndex) {
    var functorWithIndexCompose = Data_FunctorWithIndex.functorWithIndexCompose(dictTraversableWithIndex.FunctorWithIndex0());
    var foldableWithIndexCompose = Data_FoldableWithIndex.foldableWithIndexCompose(dictTraversableWithIndex.FoldableWithIndex1());
    var traversableCompose = Data_Traversable.traversableCompose(dictTraversableWithIndex.Traversable2());
    return function (dictTraversableWithIndex1) {
        var traverseWithIndex1 = traverseWithIndex(dictTraversableWithIndex1);
        var functorWithIndexCompose1 = functorWithIndexCompose(dictTraversableWithIndex1.FunctorWithIndex0());
        var foldableWithIndexCompose1 = foldableWithIndexCompose(dictTraversableWithIndex1.FoldableWithIndex1());
        var traversableCompose1 = traversableCompose(dictTraversableWithIndex1.Traversable2());
        return {
            traverseWithIndex: function (dictApplicative) {
                var Functor0 = (dictApplicative.Apply0()).Functor0();
                var traverseWithIndex2 = traverseWithIndex1(dictApplicative);
                return function (f) {
                    return function (v) {
                        return Data_Functor.map(Functor0)(Data_Functor_Compose.Compose)(traverseWithIndex(dictTraversableWithIndex)(dictApplicative)((function () {
                            var $157 = Data_Tuple.curry(f);
                            return function ($158) {
                                return traverseWithIndex2($157($158));
                            };
                        })())(v));
                    };
                };
            },
            FunctorWithIndex0: function () {
                return functorWithIndexCompose1;
            },
            FoldableWithIndex1: function () {
                return foldableWithIndexCompose1;
            },
            Traversable2: function () {
                return traversableCompose1;
            }
        };
    };
};
var traversableWithIndexArray = {
    traverseWithIndex: function (dictApplicative) {
        return traverseWithIndexDefault(traversableWithIndexArray)(dictApplicative);
    },
    FunctorWithIndex0: function () {
        return Data_FunctorWithIndex.functorWithIndexArray;
    },
    FoldableWithIndex1: function () {
        return Data_FoldableWithIndex.foldableWithIndexArray;
    },
    Traversable2: function () {
        return Data_Traversable.traversableArray;
    }
};
var traversableWithIndexApp = function (dictTraversableWithIndex) {
    var functorWithIndexApp = Data_FunctorWithIndex.functorWithIndexApp(dictTraversableWithIndex.FunctorWithIndex0());
    var foldableWithIndexApp = Data_FoldableWithIndex.foldableWithIndexApp(dictTraversableWithIndex.FoldableWithIndex1());
    var traversableApp = Data_Traversable.traversableApp(dictTraversableWithIndex.Traversable2());
    return {
        traverseWithIndex: function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (f) {
                return function (v) {
                    return Data_Functor.map(Functor0)(Data_Functor_App.App)(traverseWithIndex(dictTraversableWithIndex)(dictApplicative)(f)(v));
                };
            };
        },
        FunctorWithIndex0: function () {
            return functorWithIndexApp;
        },
        FoldableWithIndex1: function () {
            return foldableWithIndexApp;
        },
        Traversable2: function () {
            return traversableApp;
        }
    };
};
var traversableWithIndexAdditive = {
    traverseWithIndex: function (dictApplicative) {
        var traverse8 = traverse7(dictApplicative);
        return function (f) {
            return traverse8(f(Data_Unit.unit));
        };
    },
    FunctorWithIndex0: function () {
        return Data_FunctorWithIndex.functorWithIndexAdditive;
    },
    FoldableWithIndex1: function () {
        return Data_FoldableWithIndex.foldableWithIndexAdditive;
    },
    Traversable2: function () {
        return Data_Traversable.traversableAdditive;
    }
};

// | Fold a data structure from the right with access to the indices, keeping
// | all intermediate results instead of only the final result.
// |
// | Unlike `scanrWithIndex`, `imapAccumRWithIndex` allows the type of accumulator to differ
// | from the element type of the final data structure.
var mapAccumRWithIndex = function (dictTraversableWithIndex) {
    return function (f) {
        return function (s0) {
            return function (xs) {
                return Data_Traversable_Accum_Internal.stateR(traverseWithIndex(dictTraversableWithIndex)(Data_Traversable_Accum_Internal.applicativeStateR)(function (i) {
                    return function (a) {
                        return function (s) {
                            return f(i)(s)(a);
                        };
                    };
                })(xs))(s0);
            };
        };
    };
};

// | Fold a data structure from the right with access to the indices, keeping
// | all intermediate results instead of only the final result. Note that the
// | initial value does not appear in the result (unlike Haskell's `Prelude.scanr`).
// |
// | ```purescript
// | scanrWithIndex (\i x y -> i + x + y) 0 [1, 2, 3] = [9, 8, 5]
// | ```
var scanrWithIndex = function (dictTraversableWithIndex) {
    return function (f) {
        return function (b0) {
            return function (xs) {
                return (mapAccumRWithIndex(dictTraversableWithIndex)(function (i) {
                    return function (b) {
                        return function (a) {
                            var b$prime = f(i)(a)(b);
                            return {
                                accum: b$prime,
                                value: b$prime
                            };
                        };
                    };
                })(b0)(xs)).value;
            };
        };
    };
};

// | Fold a data structure from the left with access to the indices, keeping
// | all intermediate results instead of only the final result.
// |
// | Unlike `scanlWithIndex`, `mapAccumLWithIndex` allows the type of accumulator to differ
// | from the element type of the final data structure.
var mapAccumLWithIndex = function (dictTraversableWithIndex) {
    return function (f) {
        return function (s0) {
            return function (xs) {
                return Data_Traversable_Accum_Internal.stateL(traverseWithIndex(dictTraversableWithIndex)(Data_Traversable_Accum_Internal.applicativeStateL)(function (i) {
                    return function (a) {
                        return function (s) {
                            return f(i)(s)(a);
                        };
                    };
                })(xs))(s0);
            };
        };
    };
};

// | Fold a data structure from the left with access to the indices, keeping
// | all intermediate results instead of only the final result. Note that the
// | initial value does not appear in the result (unlike Haskell's
// | `Prelude.scanl`).
// |
// | ```purescript
// | scanlWithIndex (\i y x -> i + y + x) 0 [1, 2, 3] = [1, 4, 9]
// | ```
var scanlWithIndex = function (dictTraversableWithIndex) {
    return function (f) {
        return function (b0) {
            return function (xs) {
                return (mapAccumLWithIndex(dictTraversableWithIndex)(function (i) {
                    return function (b) {
                        return function (a) {
                            var b$prime = f(i)(b)(a);
                            return {
                                accum: b$prime,
                                value: b$prime
                            };
                        };
                    };
                })(b0)(xs)).value;
            };
        };
    };
};

// | A version of `traverseWithIndex` with its arguments flipped.
// |
// |
// | This can be useful when running an action written using do notation
// | for every element in a data structure:
// |
// | For example:
// |
// | ```purescript
// | for [1, 2, 3] \i x -> do
// |   logShow i
// |   pure (x * x)
// | ```
var forWithIndex = function (dictApplicative) {
    return function (dictTraversableWithIndex) {
        return Data_Function.flip(traverseWithIndex(dictTraversableWithIndex)(dictApplicative));
    };
};
export {
    traverseWithIndex,
    traverseWithIndexDefault,
    forWithIndex,
    scanlWithIndex,
    mapAccumLWithIndex,
    scanrWithIndex,
    mapAccumRWithIndex,
    traverseDefault,
    traversableWithIndexArray,
    traversableWithIndexMaybe,
    traversableWithIndexFirst,
    traversableWithIndexLast,
    traversableWithIndexAdditive,
    traversableWithIndexDual,
    traversableWithIndexConj,
    traversableWithIndexDisj,
    traversableWithIndexMultiplicative,
    traversableWithIndexEither,
    traversableWithIndexTuple,
    traversableWithIndexIdentity,
    traversableWithIndexConst,
    traversableWithIndexProduct,
    traversableWithIndexCoproduct,
    traversableWithIndexCompose,
    traversableWithIndexApp
};
