import * as Control_Category from "../Control.Category/index.js";
import * as Control_Semigroupoid from "../Control.Semigroupoid/index.js";
import * as Data_Show from "../Data.Show/index.js";
var Endo = function (x) {
    return x;
};
var showEndo = function (dictShow) {
    return {
        show: function (v) {
            return "(Endo " + (Data_Show.show(dictShow)(v) + ")");
        }
    };
};
var semigroupEndo = function (dictSemigroupoid) {
    return {
        append: function (v) {
            return function (v1) {
                return Control_Semigroupoid.compose(dictSemigroupoid)(v)(v1);
            };
        }
    };
};
var ordEndo = function (dictOrd) {
    return dictOrd;
};
var monoidEndo = function (dictCategory) {
    var semigroupEndo1 = semigroupEndo(dictCategory.Semigroupoid0());
    return {
        mempty: Control_Category.identity(dictCategory),
        Semigroup0: function () {
            return semigroupEndo1;
        }
    };
};
var eqEndo = function (dictEq) {
    return dictEq;
};
var boundedEndo = function (dictBounded) {
    return dictBounded;
};
export {
    Endo,
    eqEndo,
    ordEndo,
    boundedEndo,
    showEndo,
    semigroupEndo,
    monoidEndo
};
