import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";

// | An adaptor allowing `>$<` to map over the inputs of an equivalence
// | relation.
var Equivalence = function (x) {
    return x;
};
var semigroupEquivalence = {
    append: function (v) {
        return function (v1) {
            return function (a) {
                return function (b) {
                    return v(a)(b) && v1(a)(b);
                };
            };
        };
    }
};
var newtypeEquivalence = {
    Coercible0: function () {
        return undefined;
    }
};
var monoidEquivalence = {
    mempty: function (v) {
        return function (v1) {
            return true;
        };
    },
    Semigroup0: function () {
        return semigroupEquivalence;
    }
};

// | The default equivalence relation for any values with an `Eq` instance.
var defaultEquivalence = function (dictEq) {
    return Data_Eq.eq(dictEq);
};
var contravariantEquivalence = {
    cmap: function (f) {
        return function (v) {
            return Data_Function.on(v)(f);
        };
    }
};

// | An equivalence relation for any `Comparison`.
var comparisonEquivalence = function (v) {
    return function (a) {
        return function (b) {
            return Data_Eq.eq(Data_Ordering.eqOrdering)(v(a)(b))(Data_Ordering.EQ.value);
        };
    };
};
export {
    Equivalence,
    defaultEquivalence,
    comparisonEquivalence,
    newtypeEquivalence,
    contravariantEquivalence,
    semigroupEquivalence,
    monoidEquivalence
};
