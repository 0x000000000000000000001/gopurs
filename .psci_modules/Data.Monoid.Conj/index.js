import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_HeytingAlgebra from "../Data.HeytingAlgebra/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Show from "../Data.Show/index.js";

// | Monoid and semigroup for conjunction.
// |
// | ``` purescript
// | Conj x <> Conj y == Conj (x && y)
// | (mempty :: Conj _) == Conj tt
// | ```
var Conj = function (x) {
    return x;
};
var showConj = function (dictShow) {
    return {
        show: function (v) {
            return "(Conj " + (Data_Show.show(dictShow)(v) + ")");
        }
    };
};
var semiringConj = function (dictHeytingAlgebra) {
    return {
        zero: Data_HeytingAlgebra.tt(dictHeytingAlgebra),
        one: Data_HeytingAlgebra.ff(dictHeytingAlgebra),
        add: function (v) {
            return function (v1) {
                return Data_HeytingAlgebra.conj(dictHeytingAlgebra)(v)(v1);
            };
        },
        mul: function (v) {
            return function (v1) {
                return Data_HeytingAlgebra.disj(dictHeytingAlgebra)(v)(v1);
            };
        }
    };
};
var semigroupConj = function (dictHeytingAlgebra) {
    return {
        append: function (v) {
            return function (v1) {
                return Data_HeytingAlgebra.conj(dictHeytingAlgebra)(v)(v1);
            };
        }
    };
};
var ordConj = function (dictOrd) {
    return dictOrd;
};
var monoidConj = function (dictHeytingAlgebra) {
    var semigroupConj1 = semigroupConj(dictHeytingAlgebra);
    return {
        mempty: Data_HeytingAlgebra.tt(dictHeytingAlgebra),
        Semigroup0: function () {
            return semigroupConj1;
        }
    };
};
var functorConj = {
    map: function (f) {
        return function (m) {
            return f(m);
        };
    }
};
var eqConj = function (dictEq) {
    return dictEq;
};
var eq1Conj = {
    eq1: function (dictEq) {
        return Data_Eq.eq(eqConj(dictEq));
    }
};
var ord1Conj = {
    compare1: function (dictOrd) {
        return Data_Ord.compare(ordConj(dictOrd));
    },
    Eq10: function () {
        return eq1Conj;
    }
};
var boundedConj = function (dictBounded) {
    return dictBounded;
};
var applyConj = {
    apply: function (v) {
        return function (v1) {
            return v(v1);
        };
    },
    Functor0: function () {
        return functorConj;
    }
};
var bindConj = {
    bind: function (v) {
        return function (f) {
            return f(v);
        };
    },
    Apply0: function () {
        return applyConj;
    }
};
var applicativeConj = {
    pure: Conj,
    Apply0: function () {
        return applyConj;
    }
};
var monadConj = {
    Applicative0: function () {
        return applicativeConj;
    },
    Bind1: function () {
        return bindConj;
    }
};
export {
    Conj,
    eqConj,
    eq1Conj,
    ordConj,
    ord1Conj,
    boundedConj,
    showConj,
    functorConj,
    applyConj,
    applicativeConj,
    bindConj,
    monadConj,
    semigroupConj,
    monoidConj,
    semiringConj
};
