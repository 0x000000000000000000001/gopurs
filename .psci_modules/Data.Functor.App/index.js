import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Unsafe_Coerce from "../Unsafe.Coerce/index.js";
var App = function (x) {
    return x;
};
var showApp = function (dictShow) {
    return {
        show: function (v) {
            return "(App " + (Data_Show.show(dictShow)(v) + ")");
        }
    };
};
var semigroupApp = function (dictApply) {
    return function (dictSemigroup) {
        var append = Data_Semigroup.append(dictSemigroup);
        return {
            append: function (v) {
                return function (v1) {
                    return Control_Apply.lift2(dictApply)(append)(v)(v1);
                };
            }
        };
    };
};
var plusApp = function (dictPlus) {
    return dictPlus;
};
var newtypeApp = {
    Coercible0: function () {
        return undefined;
    }
};
var monoidApp = function (dictApplicative) {
    var semigroupApp1 = semigroupApp(dictApplicative.Apply0());
    return function (dictMonoid) {
        var semigroupApp2 = semigroupApp1(dictMonoid.Semigroup0());
        return {
            mempty: Control_Applicative.pure(dictApplicative)(Data_Monoid.mempty(dictMonoid)),
            Semigroup0: function () {
                return semigroupApp2;
            }
        };
    };
};
var monadPlusApp = function (dictMonadPlus) {
    return dictMonadPlus;
};
var monadApp = function (dictMonad) {
    return dictMonad;
};
var lazyApp = function (dictLazy) {
    return dictLazy;
};
var hoistLowerApp = Unsafe_Coerce.unsafeCoerce;
var hoistLiftApp = Unsafe_Coerce.unsafeCoerce;
var hoistApp = function (f) {
    return function (v) {
        return f(v);
    };
};
var functorApp = function (dictFunctor) {
    return dictFunctor;
};
var extendApp = function (dictExtend) {
    return dictExtend;
};
var eqApp = function (dictEq1) {
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
var ordApp = function (dictOrd1) {
    var eqApp1 = eqApp(dictOrd1.Eq10());
    return function (dictOrd) {
        var eqApp2 = eqApp1(dictOrd.Eq0());
        return {
            compare: function (x) {
                return function (y) {
                    return Data_Ord.compare1(dictOrd1)(dictOrd)(x)(y);
                };
            },
            Eq0: function () {
                return eqApp2;
            }
        };
    };
};
var eq1App = function (dictEq1) {
    var eqApp1 = eqApp(dictEq1);
    return {
        eq1: function (dictEq) {
            return Data_Eq.eq(eqApp1(dictEq));
        }
    };
};
var ord1App = function (dictOrd1) {
    var ordApp1 = ordApp(dictOrd1);
    var eq1App1 = eq1App(dictOrd1.Eq10());
    return {
        compare1: function (dictOrd) {
            return Data_Ord.compare(ordApp1(dictOrd));
        },
        Eq10: function () {
            return eq1App1;
        }
    };
};
var comonadApp = function (dictComonad) {
    return dictComonad;
};
var bindApp = function (dictBind) {
    return dictBind;
};
var applyApp = function (dictApply) {
    return dictApply;
};
var applicativeApp = function (dictApplicative) {
    return dictApplicative;
};
var alternativeApp = function (dictAlternative) {
    return dictAlternative;
};
var altApp = function (dictAlt) {
    return dictAlt;
};
export {
    App,
    hoistApp,
    hoistLiftApp,
    hoistLowerApp,
    newtypeApp,
    eqApp,
    eq1App,
    ordApp,
    ord1App,
    showApp,
    semigroupApp,
    monoidApp,
    functorApp,
    applyApp,
    applicativeApp,
    bindApp,
    monadApp,
    altApp,
    plusApp,
    alternativeApp,
    monadPlusApp,
    lazyApp,
    extendApp,
    comonadApp
};
