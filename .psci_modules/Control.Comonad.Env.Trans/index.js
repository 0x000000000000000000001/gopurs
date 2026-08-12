// | This module defines the environment comonad transformer, `EnvT`.
import * as Control_Comonad from "../Control.Comonad/index.js";
import * as Control_Extend from "../Control.Extend/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_FoldableWithIndex from "../Data.FoldableWithIndex/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_FunctorWithIndex from "../Data.FunctorWithIndex/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_TraversableWithIndex from "../Data.TraversableWithIndex/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var EnvT = function (x) {
    return x;
};

// | Change the environment type in an `EnvT` context.
var withEnvT = function (f) {
    return function (v) {
        return new Data_Tuple.Tuple(f(v.value0), v.value1);
    };
};

// | Unwrap a value in the `EnvT` comonad.
var runEnvT = function (v) {
    return v;
};
var newtypeEnvT = {
    Coercible0: function () {
        return undefined;
    }
};

// | Change the underlying comonad and data type in an `EnvT` context.
var mapEnvT = function (f) {
    return function (v) {
        return new Data_Tuple.Tuple(v.value0, f(v.value1));
    };
};
var functorEnvT = function (dictFunctor) {
    return {
        map: function (f) {
            return function (v) {
                return new Data_Tuple.Tuple(v.value0, Data_Functor.map(dictFunctor)(f)(v.value1));
            };
        }
    };
};
var functorWithIndexEnvT = function (dictFunctorWithIndex) {
    var functorEnvT1 = functorEnvT(dictFunctorWithIndex.Functor0());
    return {
        mapWithIndex: function (f) {
            return function (v) {
                return new Data_Tuple.Tuple(v.value0, Data_FunctorWithIndex.mapWithIndex(dictFunctorWithIndex)(f)(v.value1));
            };
        },
        Functor0: function () {
            return functorEnvT1;
        }
    };
};
var foldableEnvT = function (dictFoldable) {
    return {
        foldl: function (fn) {
            return function (a) {
                return function (v) {
                    return Data_Foldable.foldl(dictFoldable)(fn)(a)(v.value1);
                };
            };
        },
        foldr: function (fn) {
            return function (a) {
                return function (v) {
                    return Data_Foldable.foldr(dictFoldable)(fn)(a)(v.value1);
                };
            };
        },
        foldMap: function (dictMonoid) {
            return function (fn) {
                return function (v) {
                    return Data_Foldable.foldMap(dictFoldable)(dictMonoid)(fn)(v.value1);
                };
            };
        }
    };
};
var foldableWithIndexEnvT = function (dictFoldableWithIndex) {
    var foldableEnvT1 = foldableEnvT(dictFoldableWithIndex.Foldable0());
    return {
        foldlWithIndex: function (f) {
            return function (a) {
                return function (v) {
                    return Data_FoldableWithIndex.foldlWithIndex(dictFoldableWithIndex)(f)(a)(v.value1);
                };
            };
        },
        foldrWithIndex: function (f) {
            return function (a) {
                return function (v) {
                    return Data_FoldableWithIndex.foldrWithIndex(dictFoldableWithIndex)(f)(a)(v.value1);
                };
            };
        },
        foldMapWithIndex: function (dictMonoid) {
            return function (f) {
                return function (v) {
                    return Data_FoldableWithIndex.foldMapWithIndex(dictFoldableWithIndex)(dictMonoid)(f)(v.value1);
                };
            };
        },
        Foldable0: function () {
            return foldableEnvT1;
        }
    };
};
var traversableEnvT = function (dictTraversable) {
    var functorEnvT1 = functorEnvT(dictTraversable.Functor0());
    var foldableEnvT1 = foldableEnvT(dictTraversable.Foldable1());
    return {
        sequence: function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (v) {
                return Data_Functor.map(Functor0)(Data_Functor.map(Data_Functor.functorFn)(EnvT)(Data_Tuple.Tuple.create(v.value0)))(Data_Traversable.sequence(dictTraversable)(dictApplicative)(v.value1));
            };
        },
        traverse: function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (f) {
                return function (v) {
                    return Data_Functor.map(Functor0)(Data_Functor.map(Data_Functor.functorFn)(EnvT)(Data_Tuple.Tuple.create(v.value0)))(Data_Traversable.traverse(dictTraversable)(dictApplicative)(f)(v.value1));
                };
            };
        },
        Functor0: function () {
            return functorEnvT1;
        },
        Foldable1: function () {
            return foldableEnvT1;
        }
    };
};
var traversableWithIndexEnvT = function (dictTraversableWithIndex) {
    var functorWithIndexEnvT1 = functorWithIndexEnvT(dictTraversableWithIndex.FunctorWithIndex0());
    var foldableWithIndexEnvT1 = foldableWithIndexEnvT(dictTraversableWithIndex.FoldableWithIndex1());
    var traversableEnvT1 = traversableEnvT(dictTraversableWithIndex.Traversable2());
    return {
        traverseWithIndex: function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (f) {
                return function (v) {
                    return Data_Functor.map(Functor0)(Data_Functor.map(Data_Functor.functorFn)(EnvT)(Data_Tuple.Tuple.create(v.value0)))(Data_TraversableWithIndex.traverseWithIndex(dictTraversableWithIndex)(dictApplicative)(f)(v.value1));
                };
            };
        },
        FunctorWithIndex0: function () {
            return functorWithIndexEnvT1;
        },
        FoldableWithIndex1: function () {
            return foldableWithIndexEnvT1;
        },
        Traversable2: function () {
            return traversableEnvT1;
        }
    };
};
var extendEnvT = function (dictExtend) {
    var Functor0 = dictExtend.Functor0();
    var functorEnvT1 = functorEnvT(dictExtend.Functor0());
    return {
        extend: function (f) {
            return function (v) {
                return new Data_Tuple.Tuple(v.value0, Data_Functor.map(Functor0)(f)(Control_Extend.extend(dictExtend)((function () {
                    var $122 = Data_Tuple.Tuple.create(v.value0);
                    return function ($123) {
                        return EnvT($122($123));
                    };
                })())(v.value1)));
            };
        },
        Functor0: function () {
            return functorEnvT1;
        }
    };
};
var comonadTransEnvT = {
    lower: function (dictComonad) {
        return function (v) {
            return v.value1;
        };
    }
};
var comonadEnvT = function (dictComonad) {
    var extendEnvT1 = extendEnvT(dictComonad.Extend0());
    return {
        extract: function (v) {
            return Control_Comonad.extract(dictComonad)(v.value1);
        },
        Extend0: function () {
            return extendEnvT1;
        }
    };
};
export {
    EnvT,
    runEnvT,
    withEnvT,
    mapEnvT,
    newtypeEnvT,
    functorEnvT,
    extendEnvT,
    comonadEnvT,
    comonadTransEnvT,
    foldableEnvT,
    traversableEnvT,
    functorWithIndexEnvT,
    foldableWithIndexEnvT,
    traversableWithIndexEnvT
};
