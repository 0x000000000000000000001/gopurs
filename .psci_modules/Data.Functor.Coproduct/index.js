import * as Control_Comonad from "../Control.Comonad/index.js";
import * as Control_Extend from "../Control.Extend/index.js";
import * as Data_Bifunctor from "../Data.Bifunctor/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Show from "../Data.Show/index.js";
var Coproduct = function (x) {
    return x;
};
var showCoproduct = function (dictShow) {
    return function (dictShow1) {
        return {
            show: function (v) {
                if (v instanceof Data_Either.Left) {
                    return "(left " + (Data_Show.show(dictShow)(v.value0) + ")");
                };
                if (v instanceof Data_Either.Right) {
                    return "(right " + (Data_Show.show(dictShow1)(v.value0) + ")");
                };
                throw new Error("Failed pattern match at Data.Functor.Coproduct (line 63, column 1 - line 65, column 60): " + [ v.constructor.name ]);
            }
        };
    };
};

// | Right injection
var right = function (ga) {
    return new Data_Either.Right(ga);
};
var newtypeCoproduct = {
    Coercible0: function () {
        return undefined;
    }
};

// | Left injection
var left = function (fa) {
    return new Data_Either.Left(fa);
};
var functorCoproduct = function (dictFunctor) {
    return function (dictFunctor1) {
        return {
            map: function (f) {
                return function (v) {
                    return Data_Bifunctor.bimap(Data_Bifunctor.bifunctorEither)(Data_Functor.map(dictFunctor)(f))(Data_Functor.map(dictFunctor1)(f))(v);
                };
            }
        };
    };
};
var eq1Coproduct = function (dictEq1) {
    return function (dictEq11) {
        return {
            eq1: function (dictEq) {
                return function (v) {
                    return function (v1) {
                        if (v instanceof Data_Either.Left && v1 instanceof Data_Either.Left) {
                            return Data_Eq.eq1(dictEq1)(dictEq)(v.value0)(v1.value0);
                        };
                        if (v instanceof Data_Either.Right && v1 instanceof Data_Either.Right) {
                            return Data_Eq.eq1(dictEq11)(dictEq)(v.value0)(v1.value0);
                        };
                        return false;
                    };
                };
            }
        };
    };
};
var eqCoproduct = function (dictEq1) {
    var eq1Coproduct1 = eq1Coproduct(dictEq1);
    return function (dictEq11) {
        var eq1 = Data_Eq.eq1(eq1Coproduct1(dictEq11));
        return function (dictEq) {
            return {
                eq: eq1(dictEq)
            };
        };
    };
};
var ord1Coproduct = function (dictOrd1) {
    var eq1Coproduct1 = eq1Coproduct(dictOrd1.Eq10());
    return function (dictOrd11) {
        var eq1Coproduct2 = eq1Coproduct1(dictOrd11.Eq10());
        return {
            compare1: function (dictOrd) {
                return function (v) {
                    return function (v1) {
                        if (v instanceof Data_Either.Left && v1 instanceof Data_Either.Left) {
                            return Data_Ord.compare1(dictOrd1)(dictOrd)(v.value0)(v1.value0);
                        };
                        if (v instanceof Data_Either.Left) {
                            return Data_Ordering.LT.value;
                        };
                        if (v1 instanceof Data_Either.Left) {
                            return Data_Ordering.GT.value;
                        };
                        if (v instanceof Data_Either.Right && v1 instanceof Data_Either.Right) {
                            return Data_Ord.compare1(dictOrd11)(dictOrd)(v.value0)(v1.value0);
                        };
                        throw new Error("Failed pattern match at Data.Functor.Coproduct (line 57, column 5 - line 61, column 43): " + [ v.constructor.name, v1.constructor.name ]);
                    };
                };
            },
            Eq10: function () {
                return eq1Coproduct2;
            }
        };
    };
};
var ordCoproduct = function (dictOrd1) {
    var ord1Coproduct1 = ord1Coproduct(dictOrd1);
    var eqCoproduct1 = eqCoproduct(dictOrd1.Eq10());
    return function (dictOrd11) {
        var compare1 = Data_Ord.compare1(ord1Coproduct1(dictOrd11));
        var eqCoproduct2 = eqCoproduct1(dictOrd11.Eq10());
        return function (dictOrd) {
            var eqCoproduct3 = eqCoproduct2(dictOrd.Eq0());
            return {
                compare: compare1(dictOrd),
                Eq0: function () {
                    return eqCoproduct3;
                }
            };
        };
    };
};

// | Eliminate a coproduct by providing eliminators for the left and
// | right components
var coproduct = function (v) {
    return function (v1) {
        return function (v2) {
            if (v2 instanceof Data_Either.Left) {
                return v(v2.value0);
            };
            if (v2 instanceof Data_Either.Right) {
                return v1(v2.value0);
            };
            throw new Error("Failed pattern match at Data.Functor.Coproduct (line 27, column 1 - line 27, column 78): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
        };
    };
};
var extendCoproduct = function (dictExtend) {
    var functorCoproduct1 = functorCoproduct(dictExtend.Functor0());
    return function (dictExtend1) {
        var functorCoproduct2 = functorCoproduct1(dictExtend1.Functor0());
        return {
            extend: function (f) {
                var $89 = coproduct((function () {
                    var $91 = Control_Extend.extend(dictExtend)(function ($93) {
                        return f(Coproduct(Data_Either.Left.create($93)));
                    });
                    return function ($92) {
                        return Data_Either.Left.create($91($92));
                    };
                })())((function () {
                    var $94 = Control_Extend.extend(dictExtend1)(function ($96) {
                        return f(Coproduct(Data_Either.Right.create($96)));
                    });
                    return function ($95) {
                        return Data_Either.Right.create($94($95));
                    };
                })());
                return function ($90) {
                    return Coproduct($89($90));
                };
            },
            Functor0: function () {
                return functorCoproduct2;
            }
        };
    };
};
var comonadCoproduct = function (dictComonad) {
    var extract = Control_Comonad.extract(dictComonad);
    var extendCoproduct1 = extendCoproduct(dictComonad.Extend0());
    return function (dictComonad1) {
        var extendCoproduct2 = extendCoproduct1(dictComonad1.Extend0());
        return {
            extract: coproduct(extract)(Control_Comonad.extract(dictComonad1)),
            Extend0: function () {
                return extendCoproduct2;
            }
        };
    };
};

// | Change the underlying functors in a coproduct
var bihoistCoproduct = function (natF) {
    return function (natG) {
        return function (v) {
            return Data_Bifunctor.bimap(Data_Bifunctor.bifunctorEither)(natF)(natG)(v);
        };
    };
};
export {
    Coproduct,
    left,
    right,
    coproduct,
    bihoistCoproduct,
    newtypeCoproduct,
    eqCoproduct,
    eq1Coproduct,
    ordCoproduct,
    ord1Coproduct,
    showCoproduct,
    functorCoproduct,
    extendCoproduct,
    comonadCoproduct
};
