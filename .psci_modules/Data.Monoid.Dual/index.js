import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Show from "../Data.Show/index.js";

// | The dual of a monoid.
// |
// | ``` purescript
// | Dual x <> Dual y == Dual (y <> x)
// | (mempty :: Dual _) == Dual mempty
// | ```
var Dual = function (x) {
    return x;
};
var showDual = function (dictShow) {
    return {
        show: function (v) {
            return "(Dual " + (Data_Show.show(dictShow)(v) + ")");
        }
    };
};
var semigroupDual = function (dictSemigroup) {
    return {
        append: function (v) {
            return function (v1) {
                return Data_Semigroup.append(dictSemigroup)(v1)(v);
            };
        }
    };
};
var ordDual = function (dictOrd) {
    return dictOrd;
};
var monoidDual = function (dictMonoid) {
    var semigroupDual1 = semigroupDual(dictMonoid.Semigroup0());
    return {
        mempty: Data_Monoid.mempty(dictMonoid),
        Semigroup0: function () {
            return semigroupDual1;
        }
    };
};
var functorDual = {
    map: function (f) {
        return function (m) {
            return f(m);
        };
    }
};
var eqDual = function (dictEq) {
    return dictEq;
};
var eq1Dual = {
    eq1: function (dictEq) {
        return Data_Eq.eq(eqDual(dictEq));
    }
};
var ord1Dual = {
    compare1: function (dictOrd) {
        return Data_Ord.compare(ordDual(dictOrd));
    },
    Eq10: function () {
        return eq1Dual;
    }
};
var boundedDual = function (dictBounded) {
    return dictBounded;
};
var applyDual = {
    apply: function (v) {
        return function (v1) {
            return v(v1);
        };
    },
    Functor0: function () {
        return functorDual;
    }
};
var bindDual = {
    bind: function (v) {
        return function (f) {
            return f(v);
        };
    },
    Apply0: function () {
        return applyDual;
    }
};
var applicativeDual = {
    pure: Dual,
    Apply0: function () {
        return applyDual;
    }
};
var monadDual = {
    Applicative0: function () {
        return applicativeDual;
    },
    Bind1: function () {
        return bindDual;
    }
};
export {
    Dual,
    eqDual,
    eq1Dual,
    ordDual,
    ord1Dual,
    boundedDual,
    showDual,
    functorDual,
    applyDual,
    applicativeDual,
    bindDual,
    monadDual,
    semigroupDual,
    monoidDual
};
