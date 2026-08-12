import * as Control_Category from "../Control.Category/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";

// | The opposite of the function category.
var Op = function (x) {
    return x;
};
var semigroupoidOp = {
    compose: function (v) {
        return function (v1) {
            return function ($12) {
                return v1(v($12));
            };
        };
    }
};
var semigroupOp = function (dictSemigroup) {
    return Data_Semigroup.semigroupFn(dictSemigroup);
};
var newtypeOp = {
    Coercible0: function () {
        return undefined;
    }
};
var monoidOp = function (dictMonoid) {
    return Data_Monoid.monoidFn(dictMonoid);
};
var contravariantOp = {
    cmap: function (f) {
        return function (v) {
            return function ($13) {
                return v(f($13));
            };
        };
    }
};
var categoryOp = {
    identity: /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn),
    Semigroupoid0: function () {
        return semigroupoidOp;
    }
};
export {
    Op,
    newtypeOp,
    semigroupOp,
    monoidOp,
    semigroupoidOp,
    categoryOp,
    contravariantOp
};
