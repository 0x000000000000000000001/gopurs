import * as $foreign from "./foreign.js";
import * as Data_Bifunctor from "../Data.Bifunctor/index.js";
import * as Data_Const from "../Data.Const/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Functor_App from "../Data.Functor.App/index.js";
import * as Data_Functor_Compose from "../Data.Functor.Compose/index.js";
import * as Data_Functor_Coproduct from "../Data.Functor.Coproduct/index.js";
import * as Data_Functor_Product from "../Data.Functor.Product/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Maybe_First from "../Data.Maybe.First/index.js";
import * as Data_Maybe_Last from "../Data.Maybe.Last/index.js";
import * as Data_Monoid_Additive from "../Data.Monoid.Additive/index.js";
import * as Data_Monoid_Conj from "../Data.Monoid.Conj/index.js";
import * as Data_Monoid_Disj from "../Data.Monoid.Disj/index.js";
import * as Data_Monoid_Dual from "../Data.Monoid.Dual/index.js";
import * as Data_Monoid_Multiplicative from "../Data.Monoid.Multiplicative/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
var map = /* #__PURE__ */ Data_Functor.map(Data_Tuple.functorTuple);
var map1 = /* #__PURE__ */ Data_Functor.map(Data_Monoid_Multiplicative.functorMultiplicative);
var map2 = /* #__PURE__ */ Data_Functor.map(Data_Maybe.functorMaybe);
var map3 = /* #__PURE__ */ Data_Functor.map(Data_Maybe_Last.functorLast);
var map4 = /* #__PURE__ */ Data_Functor.map(Data_Maybe_First.functorFirst);
var map5 = /* #__PURE__ */ Data_Functor.map(Data_Either.functorEither);
var map6 = /* #__PURE__ */ Data_Functor.map(Data_Monoid_Dual.functorDual);
var map7 = /* #__PURE__ */ Data_Functor.map(Data_Monoid_Disj.functorDisj);
var map8 = /* #__PURE__ */ Data_Functor.map(Data_Monoid_Conj.functorConj);
var map9 = /* #__PURE__ */ Data_Functor.map(Data_Monoid_Additive.functorAdditive);
var mapWithIndex = function (dict) {
    return dict.mapWithIndex;
};

// | A default implementation of Functor's `map` in terms of `mapWithIndex`
var mapDefault = function (dictFunctorWithIndex) {
    return function (f) {
        return mapWithIndex(dictFunctorWithIndex)(Data_Function["const"](f));
    };
};
var functorWithIndexTuple = {
    mapWithIndex: function (f) {
        return map(f(Data_Unit.unit));
    },
    Functor0: function () {
        return Data_Tuple.functorTuple;
    }
};
var functorWithIndexProduct = function (dictFunctorWithIndex) {
    var functorProduct = Data_Functor_Product.functorProduct(dictFunctorWithIndex.Functor0());
    return function (dictFunctorWithIndex1) {
        var functorProduct1 = functorProduct(dictFunctorWithIndex1.Functor0());
        return {
            mapWithIndex: function (f) {
                return function (v) {
                    return Data_Bifunctor.bimap(Data_Bifunctor.bifunctorTuple)(mapWithIndex(dictFunctorWithIndex)(function ($53) {
                        return f(Data_Either.Left.create($53));
                    }))(mapWithIndex(dictFunctorWithIndex1)(function ($54) {
                        return f(Data_Either.Right.create($54));
                    }))(v);
                };
            },
            Functor0: function () {
                return functorProduct1;
            }
        };
    };
};
var functorWithIndexMultiplicative = {
    mapWithIndex: function (f) {
        return map1(f(Data_Unit.unit));
    },
    Functor0: function () {
        return Data_Monoid_Multiplicative.functorMultiplicative;
    }
};
var functorWithIndexMaybe = {
    mapWithIndex: function (f) {
        return map2(f(Data_Unit.unit));
    },
    Functor0: function () {
        return Data_Maybe.functorMaybe;
    }
};
var functorWithIndexLast = {
    mapWithIndex: function (f) {
        return map3(f(Data_Unit.unit));
    },
    Functor0: function () {
        return Data_Maybe_Last.functorLast;
    }
};
var functorWithIndexIdentity = {
    mapWithIndex: function (f) {
        return function (v) {
            return f(Data_Unit.unit)(v);
        };
    },
    Functor0: function () {
        return Data_Identity.functorIdentity;
    }
};
var functorWithIndexFirst = {
    mapWithIndex: function (f) {
        return map4(f(Data_Unit.unit));
    },
    Functor0: function () {
        return Data_Maybe_First.functorFirst;
    }
};
var functorWithIndexEither = {
    mapWithIndex: function (f) {
        return map5(f(Data_Unit.unit));
    },
    Functor0: function () {
        return Data_Either.functorEither;
    }
};
var functorWithIndexDual = {
    mapWithIndex: function (f) {
        return map6(f(Data_Unit.unit));
    },
    Functor0: function () {
        return Data_Monoid_Dual.functorDual;
    }
};
var functorWithIndexDisj = {
    mapWithIndex: function (f) {
        return map7(f(Data_Unit.unit));
    },
    Functor0: function () {
        return Data_Monoid_Disj.functorDisj;
    }
};
var functorWithIndexCoproduct = function (dictFunctorWithIndex) {
    var functorCoproduct = Data_Functor_Coproduct.functorCoproduct(dictFunctorWithIndex.Functor0());
    return function (dictFunctorWithIndex1) {
        var functorCoproduct1 = functorCoproduct(dictFunctorWithIndex1.Functor0());
        return {
            mapWithIndex: function (f) {
                return function (v) {
                    return Data_Bifunctor.bimap(Data_Bifunctor.bifunctorEither)(mapWithIndex(dictFunctorWithIndex)(function ($55) {
                        return f(Data_Either.Left.create($55));
                    }))(mapWithIndex(dictFunctorWithIndex1)(function ($56) {
                        return f(Data_Either.Right.create($56));
                    }))(v);
                };
            },
            Functor0: function () {
                return functorCoproduct1;
            }
        };
    };
};
var functorWithIndexConst = {
    mapWithIndex: function (v) {
        return function (v1) {
            return v1;
        };
    },
    Functor0: function () {
        return Data_Const.functorConst;
    }
};
var functorWithIndexConj = {
    mapWithIndex: function (f) {
        return map8(f(Data_Unit.unit));
    },
    Functor0: function () {
        return Data_Monoid_Conj.functorConj;
    }
};
var functorWithIndexCompose = function (dictFunctorWithIndex) {
    var functorCompose = Data_Functor_Compose.functorCompose(dictFunctorWithIndex.Functor0());
    return function (dictFunctorWithIndex1) {
        var mapWithIndex1 = mapWithIndex(dictFunctorWithIndex1);
        var functorCompose1 = functorCompose(dictFunctorWithIndex1.Functor0());
        return {
            mapWithIndex: function (f) {
                return function (v) {
                    return mapWithIndex(dictFunctorWithIndex)((function () {
                        var $57 = Data_Tuple.curry(f);
                        return function ($58) {
                            return mapWithIndex1($57($58));
                        };
                    })())(v);
                };
            },
            Functor0: function () {
                return functorCompose1;
            }
        };
    };
};
var functorWithIndexArray = {
    mapWithIndex: $foreign.mapWithIndexArray,
    Functor0: function () {
        return Data_Functor.functorArray;
    }
};
var functorWithIndexApp = function (dictFunctorWithIndex) {
    var functorApp = Data_Functor_App.functorApp(dictFunctorWithIndex.Functor0());
    return {
        mapWithIndex: function (f) {
            return function (v) {
                return mapWithIndex(dictFunctorWithIndex)(f)(v);
            };
        },
        Functor0: function () {
            return functorApp;
        }
    };
};
var functorWithIndexAdditive = {
    mapWithIndex: function (f) {
        return map9(f(Data_Unit.unit));
    },
    Functor0: function () {
        return Data_Monoid_Additive.functorAdditive;
    }
};
export {
    mapWithIndex,
    mapDefault,
    functorWithIndexArray,
    functorWithIndexMaybe,
    functorWithIndexFirst,
    functorWithIndexLast,
    functorWithIndexAdditive,
    functorWithIndexDual,
    functorWithIndexConj,
    functorWithIndexDisj,
    functorWithIndexMultiplicative,
    functorWithIndexEither,
    functorWithIndexTuple,
    functorWithIndexIdentity,
    functorWithIndexConst,
    functorWithIndexProduct,
    functorWithIndexCoproduct,
    functorWithIndexCompose,
    functorWithIndexApp
};
