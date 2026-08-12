import * as Control_Comonad from "../Control.Comonad/index.js";
import * as Control_Extend from "../Control.Extend/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
var IdentityT = function (x) {
    return x;
};
var monadSTIdentityT = function (dictMonadST) {
    return dictMonadST;
};
var traversableIdentityT = function (dictTraversable) {
    return dictTraversable;
};

// | Run a computation in the `IdentityT` monad.
var runIdentityT = function (v) {
    return v;
};
var plusIdentityT = function (dictPlus) {
    return dictPlus;
};
var newtypeIdentityT = {
    Coercible0: function () {
        return undefined;
    }
};
var monadWriterIdentityT = function (dictMonadWriter) {
    return dictMonadWriter;
};
var monadTransIdentityT = {
    lift: function (dictMonad) {
        return IdentityT;
    }
};
var monadThrowIdentityT = function (dictMonadThrow) {
    return dictMonadThrow;
};
var monadTellIdentityT = function (dictMonadTell) {
    return dictMonadTell;
};
var monadStateIdentityT = function (dictMonadState) {
    return dictMonadState;
};
var monadRecIdentityT = function (dictMonadRec) {
    return dictMonadRec;
};
var monadReaderIdentityT = function (dictMonadReader) {
    return dictMonadReader;
};
var monadPlusIdentityT = function (dictMonadPlus) {
    return dictMonadPlus;
};
var monadIdentityT = function (dictMonad) {
    return dictMonad;
};
var monadErrorIdentityT = function (dictMonadError) {
    return dictMonadError;
};
var monadEffectIdentityT = function (dictMonadEffect) {
    return dictMonadEffect;
};
var monadContIdentityT = function (dictMonadCont) {
    return dictMonadCont;
};
var monadAskIdentityT = function (dictMonadAsk) {
    return dictMonadAsk;
};

// | Change the result type of a `IdentityT` monad action.
var mapIdentityT = function (f) {
    return function (v) {
        return f(v);
    };
};
var functorIdentityT = function (dictFunctor) {
    return dictFunctor;
};
var foldableIdentityT = function (dictFoldable) {
    return dictFoldable;
};
var extendIdentityI = function (dictExtend) {
    var functorIdentityT1 = functorIdentityT(dictExtend.Functor0());
    return {
        extend: function (f) {
            return function (v) {
                return Control_Extend.extend(dictExtend)(function ($71) {
                    return f(IdentityT($71));
                })(v);
            };
        },
        Functor0: function () {
            return functorIdentityT1;
        }
    };
};
var eqIdentityT = function (dictEq1) {
    return function (dictEq) {
        return {
            eq: function (x) {
                return function (y) {
                    return Data_Eq.eq1(dictEq1)(dictEq)(x)(y);
                };
            }
        };
    };
};
var ordIdentityT = function (dictOrd1) {
    var eqIdentityT1 = eqIdentityT(dictOrd1.Eq10());
    return function (dictOrd) {
        var eqIdentityT2 = eqIdentityT1(dictOrd.Eq0());
        return {
            compare: function (x) {
                return function (y) {
                    return Data_Ord.compare1(dictOrd1)(dictOrd)(x)(y);
                };
            },
            Eq0: function () {
                return eqIdentityT2;
            }
        };
    };
};
var eq1IdentityT = function (dictEq1) {
    var eqIdentityT1 = eqIdentityT(dictEq1);
    return {
        eq1: function (dictEq) {
            return Data_Eq.eq(eqIdentityT1(dictEq));
        }
    };
};
var ord1IdentityT = function (dictOrd1) {
    var ordIdentityT1 = ordIdentityT(dictOrd1);
    var eq1IdentityT1 = eq1IdentityT(dictOrd1.Eq10());
    return {
        compare1: function (dictOrd) {
            return Data_Ord.compare(ordIdentityT1(dictOrd));
        },
        Eq10: function () {
            return eq1IdentityT1;
        }
    };
};
var comonadIdentityT = function (dictComonad) {
    var extendIdentityI1 = extendIdentityI(dictComonad.Extend0());
    return {
        extract: (function () {
            var $72 = Control_Comonad.extract(dictComonad);
            return function ($73) {
                return $72(runIdentityT($73));
            };
        })(),
        Extend0: function () {
            return extendIdentityI1;
        }
    };
};
var bindIdentityT = function (dictBind) {
    return dictBind;
};
var applyIdentityT = function (dictApply) {
    return dictApply;
};
var applicativeIdentityT = function (dictApplicative) {
    return dictApplicative;
};
var alternativeIdentityT = function (dictAlternative) {
    return dictAlternative;
};
var altIdentityT = function (dictAlt) {
    return dictAlt;
};
export {
    IdentityT,
    runIdentityT,
    mapIdentityT,
    eqIdentityT,
    ordIdentityT,
    eq1IdentityT,
    ord1IdentityT,
    newtypeIdentityT,
    functorIdentityT,
    applyIdentityT,
    applicativeIdentityT,
    altIdentityT,
    plusIdentityT,
    alternativeIdentityT,
    bindIdentityT,
    monadIdentityT,
    monadRecIdentityT,
    monadPlusIdentityT,
    monadTransIdentityT,
    monadEffectIdentityT,
    monadContIdentityT,
    monadThrowIdentityT,
    monadErrorIdentityT,
    monadAskIdentityT,
    monadReaderIdentityT,
    monadStateIdentityT,
    monadTellIdentityT,
    monadWriterIdentityT,
    foldableIdentityT,
    traversableIdentityT,
    monadSTIdentityT,
    extendIdentityI,
    comonadIdentityT
};
