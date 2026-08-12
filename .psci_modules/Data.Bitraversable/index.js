import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Bifoldable from "../Data.Bifoldable/index.js";
import * as Data_Bifunctor from "../Data.Bifunctor/index.js";
import * as Data_Const from "../Data.Const/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Functor_Clown from "../Data.Functor.Clown/index.js";
import * as Data_Functor_Flip from "../Data.Functor.Flip/index.js";
import * as Data_Functor_Joker from "../Data.Functor.Joker/index.js";
import * as Data_Functor_Product2 from "../Data.Functor.Product2/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var identity1 = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var bitraverse = function (dict) {
    return dict.bitraverse;
};
var lfor = function (dictBitraversable) {
    return function (dictApplicative) {
        var pure = Control_Applicative.pure(dictApplicative);
        return function (t) {
            return function (f) {
                return bitraverse(dictBitraversable)(dictApplicative)(f)(pure)(t);
            };
        };
    };
};
var ltraverse = function (dictBitraversable) {
    return function (dictApplicative) {
        var pure = Control_Applicative.pure(dictApplicative);
        return function (f) {
            return bitraverse(dictBitraversable)(dictApplicative)(f)(pure);
        };
    };
};
var rfor = function (dictBitraversable) {
    return function (dictApplicative) {
        var pure = Control_Applicative.pure(dictApplicative);
        return function (t) {
            return function (f) {
                return bitraverse(dictBitraversable)(dictApplicative)(pure)(f)(t);
            };
        };
    };
};
var rtraverse = function (dictBitraversable) {
    return function (dictApplicative) {
        return bitraverse(dictBitraversable)(dictApplicative)(Control_Applicative.pure(dictApplicative));
    };
};
var bitraversableTuple = {
    bitraverse: function (dictApplicative) {
        var Apply0 = dictApplicative.Apply0();
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (f) {
            return function (g) {
                return function (v) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Data_Tuple.Tuple.create)(f(v.value0)))(g(v.value1));
                };
            };
        };
    },
    bisequence: function (dictApplicative) {
        var Apply0 = dictApplicative.Apply0();
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (v) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Data_Tuple.Tuple.create)(v.value0))(v.value1);
        };
    },
    Bifunctor0: function () {
        return Data_Bifunctor.bifunctorTuple;
    },
    Bifoldable1: function () {
        return Data_Bifoldable.bifoldableTuple;
    }
};
var bitraversableJoker = function (dictTraversable) {
    var bifunctorJoker = Data_Functor_Joker.bifunctorJoker(dictTraversable.Functor0());
    var bifoldableJoker = Data_Bifoldable.bifoldableJoker(dictTraversable.Foldable1());
    return {
        bitraverse: function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (v) {
                return function (r) {
                    return function (v1) {
                        return Data_Functor.map(Functor0)(Data_Functor_Joker.Joker)(Data_Traversable.traverse(dictTraversable)(dictApplicative)(r)(v1));
                    };
                };
            };
        },
        bisequence: function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (v) {
                return Data_Functor.map(Functor0)(Data_Functor_Joker.Joker)(Data_Traversable.sequence(dictTraversable)(dictApplicative)(v));
            };
        },
        Bifunctor0: function () {
            return bifunctorJoker;
        },
        Bifoldable1: function () {
            return bifoldableJoker;
        }
    };
};
var bitraversableEither = {
    bitraverse: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (v) {
            return function (v1) {
                return function (v2) {
                    if (v2 instanceof Data_Either.Left) {
                        return Data_Functor.map(Functor0)(Data_Either.Left.create)(v(v2.value0));
                    };
                    if (v2 instanceof Data_Either.Right) {
                        return Data_Functor.map(Functor0)(Data_Either.Right.create)(v1(v2.value0));
                    };
                    throw new Error("Failed pattern match at Data.Bitraversable (line 57, column 1 - line 61, column 37): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
                };
            };
        };
    },
    bisequence: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (v) {
            if (v instanceof Data_Either.Left) {
                return Data_Functor.map(Functor0)(Data_Either.Left.create)(v.value0);
            };
            if (v instanceof Data_Either.Right) {
                return Data_Functor.map(Functor0)(Data_Either.Right.create)(v.value0);
            };
            throw new Error("Failed pattern match at Data.Bitraversable (line 57, column 1 - line 61, column 37): " + [ v.constructor.name ]);
        };
    },
    Bifunctor0: function () {
        return Data_Bifunctor.bifunctorEither;
    },
    Bifoldable1: function () {
        return Data_Bifoldable.bifoldableEither;
    }
};
var bitraversableConst = {
    bitraverse: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (f) {
            return function (v) {
                return function (v1) {
                    return Data_Functor.map(Functor0)(Data_Const.Const)(f(v1));
                };
            };
        };
    },
    bisequence: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (v) {
            return Data_Functor.map(Functor0)(Data_Const.Const)(v);
        };
    },
    Bifunctor0: function () {
        return Data_Bifunctor.bifunctorConst;
    },
    Bifoldable1: function () {
        return Data_Bifoldable.bifoldableConst;
    }
};
var bitraversableClown = function (dictTraversable) {
    var bifunctorClown = Data_Functor_Clown.bifunctorClown(dictTraversable.Functor0());
    var bifoldableClown = Data_Bifoldable.bifoldableClown(dictTraversable.Foldable1());
    return {
        bitraverse: function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (l) {
                return function (v) {
                    return function (v1) {
                        return Data_Functor.map(Functor0)(Data_Functor_Clown.Clown)(Data_Traversable.traverse(dictTraversable)(dictApplicative)(l)(v1));
                    };
                };
            };
        },
        bisequence: function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (v) {
                return Data_Functor.map(Functor0)(Data_Functor_Clown.Clown)(Data_Traversable.sequence(dictTraversable)(dictApplicative)(v));
            };
        },
        Bifunctor0: function () {
            return bifunctorClown;
        },
        Bifoldable1: function () {
            return bifoldableClown;
        }
    };
};

// | A default implementation of `bisequence` using `bitraverse`.
var bisequenceDefault = function (dictBitraversable) {
    return function (dictApplicative) {
        return bitraverse(dictBitraversable)(dictApplicative)(identity)(identity1);
    };
};
var bisequence = function (dict) {
    return dict.bisequence;
};
var bitraversableFlip = function (dictBitraversable) {
    var bifunctorFlip = Data_Functor_Flip.bifunctorFlip(dictBitraversable.Bifunctor0());
    var bifoldableFlip = Data_Bifoldable.bifoldableFlip(dictBitraversable.Bifoldable1());
    return {
        bitraverse: function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (r) {
                return function (l) {
                    return function (v) {
                        return Data_Functor.map(Functor0)(Data_Functor_Flip.Flip)(bitraverse(dictBitraversable)(dictApplicative)(l)(r)(v));
                    };
                };
            };
        },
        bisequence: function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (v) {
                return Data_Functor.map(Functor0)(Data_Functor_Flip.Flip)(bisequence(dictBitraversable)(dictApplicative)(v));
            };
        },
        Bifunctor0: function () {
            return bifunctorFlip;
        },
        Bifoldable1: function () {
            return bifoldableFlip;
        }
    };
};
var bitraversableProduct2 = function (dictBitraversable) {
    var bifunctorProduct2 = Data_Functor_Product2.bifunctorProduct2(dictBitraversable.Bifunctor0());
    var bifoldableProduct2 = Data_Bifoldable.bifoldableProduct2(dictBitraversable.Bifoldable1());
    return function (dictBitraversable1) {
        var bifunctorProduct21 = bifunctorProduct2(dictBitraversable1.Bifunctor0());
        var bifoldableProduct21 = bifoldableProduct2(dictBitraversable1.Bifoldable1());
        return {
            bitraverse: function (dictApplicative) {
                var Apply0 = dictApplicative.Apply0();
                var Functor0 = (dictApplicative.Apply0()).Functor0();
                return function (l) {
                    return function (r) {
                        return function (v) {
                            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Data_Functor_Product2.Product2.create)(bitraverse(dictBitraversable)(dictApplicative)(l)(r)(v.value0)))(bitraverse(dictBitraversable1)(dictApplicative)(l)(r)(v.value1));
                        };
                    };
                };
            },
            bisequence: function (dictApplicative) {
                var Apply0 = dictApplicative.Apply0();
                var Functor0 = (dictApplicative.Apply0()).Functor0();
                return function (v) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Data_Functor_Product2.Product2.create)(bisequence(dictBitraversable)(dictApplicative)(v.value0)))(bisequence(dictBitraversable1)(dictApplicative)(v.value1));
                };
            },
            Bifunctor0: function () {
                return bifunctorProduct21;
            },
            Bifoldable1: function () {
                return bifoldableProduct21;
            }
        };
    };
};

// | A default implementation of `bitraverse` using `bisequence` and `bimap`.
var bitraverseDefault = function (dictBitraversable) {
    var Bifunctor0 = dictBitraversable.Bifunctor0();
    return function (dictApplicative) {
        return function (f) {
            return function (g) {
                return function (t) {
                    return bisequence(dictBitraversable)(dictApplicative)(Data_Bifunctor.bimap(Bifunctor0)(f)(g)(t));
                };
            };
        };
    };
};

// | Traverse a data structure, accumulating effects and results using an `Applicative` functor.
var bifor = function (dictBitraversable) {
    return function (dictApplicative) {
        return function (t) {
            return function (f) {
                return function (g) {
                    return bitraverse(dictBitraversable)(dictApplicative)(f)(g)(t);
                };
            };
        };
    };
};
export {
    bitraverse,
    bisequence,
    bitraverseDefault,
    bisequenceDefault,
    ltraverse,
    rtraverse,
    bifor,
    lfor,
    rfor,
    bitraversableClown,
    bitraversableJoker,
    bitraversableFlip,
    bitraversableProduct2,
    bitraversableEither,
    bitraversableTuple,
    bitraversableConst
};
export {
    biall,
    biany,
    bifold,
    bifoldMap,
    bifoldMapDefaultL,
    bifoldMapDefaultR,
    bifoldl,
    bifoldlDefault,
    bifoldr,
    bifoldrDefault,
    bifor_,
    bisequence_,
    bitraverse_
} from "../Data.Bifoldable/index.js";
