import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_HeytingAlgebra from "../Data.HeytingAlgebra/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Show from "../Data.Show/index.js";

// | Monoid and semigroup for disjunction.
// |
// | ``` purescript
// | Disj x <> Disj y == Disj (x || y)
// | (mempty :: Disj _) == Disj bottom
// | ```
var Disj = function (x) {
    return x;
};
var showDisj = function (dictShow) {
    return {
        show: function (v) {
            return "(Disj " + (Data_Show.show(dictShow)(v) + ")");
        }
    };
};
var semiringDisj = function (dictHeytingAlgebra) {
    return {
        zero: Data_HeytingAlgebra.ff(dictHeytingAlgebra),
        one: Data_HeytingAlgebra.tt(dictHeytingAlgebra),
        add: function (v) {
            return function (v1) {
                return Data_HeytingAlgebra.disj(dictHeytingAlgebra)(v)(v1);
            };
        },
        mul: function (v) {
            return function (v1) {
                return Data_HeytingAlgebra.conj(dictHeytingAlgebra)(v)(v1);
            };
        }
    };
};
var semigroupDisj = function (dictHeytingAlgebra) {
    return {
        append: function (v) {
            return function (v1) {
                return Data_HeytingAlgebra.disj(dictHeytingAlgebra)(v)(v1);
            };
        }
    };
};
var ordDisj = function (dictOrd) {
    return dictOrd;
};
var monoidDisj = function (dictHeytingAlgebra) {
    var semigroupDisj1 = semigroupDisj(dictHeytingAlgebra);
    return {
        mempty: Data_HeytingAlgebra.ff(dictHeytingAlgebra),
        Semigroup0: function () {
            return semigroupDisj1;
        }
    };
};
var functorDisj = {
    map: function (f) {
        return function (m) {
            return f(m);
        };
    }
};
var eqDisj = function (dictEq) {
    return dictEq;
};
var eq1Disj = {
    eq1: function (dictEq) {
        return Data_Eq.eq(eqDisj(dictEq));
    }
};
var ord1Disj = {
    compare1: function (dictOrd) {
        return Data_Ord.compare(ordDisj(dictOrd));
    },
    Eq10: function () {
        return eq1Disj;
    }
};
var boundedDisj = function (dictBounded) {
    return dictBounded;
};
var applyDisj = {
    apply: function (v) {
        return function (v1) {
            return v(v1);
        };
    },
    Functor0: function () {
        return functorDisj;
    }
};
var bindDisj = {
    bind: function (v) {
        return function (f) {
            return f(v);
        };
    },
    Apply0: function () {
        return applyDisj;
    }
};
var applicativeDisj = {
    pure: Disj,
    Apply0: function () {
        return applyDisj;
    }
};
var monadDisj = {
    Applicative0: function () {
        return applicativeDisj;
    },
    Bind1: function () {
        return bindDisj;
    }
};
export {
    Disj,
    eqDisj,
    eq1Disj,
    ordDisj,
    ord1Disj,
    boundedDisj,
    showDisj,
    functorDisj,
    applyDisj,
    applicativeDisj,
    bindDisj,
    monadDisj,
    semigroupDisj,
    monoidDisj,
    semiringDisj
};
