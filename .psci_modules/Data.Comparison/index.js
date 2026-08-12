import * as Data_Function from "../Data.Function/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
var semigroupFn = /* #__PURE__ */ Data_Semigroup.semigroupFn(/* #__PURE__ */ Data_Semigroup.semigroupFn(Data_Ordering.semigroupOrdering));

// | An adaptor allowing `>$<` to map over the inputs of a comparison function.
var Comparison = function (x) {
    return x;
};
var semigroupComparison = {
    append: function (v) {
        return function (v1) {
            return Data_Semigroup.append(semigroupFn)(v)(v1);
        };
    }
};
var newtypeComparison = {
    Coercible0: function () {
        return undefined;
    }
};
var monoidComparison = {
    mempty: function (v) {
        return function (v1) {
            return Data_Ordering.EQ.value;
        };
    },
    Semigroup0: function () {
        return semigroupComparison;
    }
};

// | The default comparison for any values with an `Ord` instance.
var defaultComparison = function (dictOrd) {
    return Data_Ord.compare(dictOrd);
};
var contravariantComparison = {
    cmap: function (f) {
        return function (v) {
            return Data_Function.on(v)(f);
        };
    }
};
export {
    Comparison,
    defaultComparison,
    newtypeComparison,
    contravariantComparison,
    semigroupComparison,
    monoidComparison
};
