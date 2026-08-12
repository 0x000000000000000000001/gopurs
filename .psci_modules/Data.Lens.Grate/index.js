// | This module defines functions for working with grates.
// |
// | See <http://r6research.livejournal.com/28050.html>.
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Distributive from "../Data.Distributive/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Lens_Types from "../Data.Lens.Types/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Profunctor from "../Data.Profunctor/index.js";
import * as Data_Profunctor_Closed from "../Data.Profunctor.Closed/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var identity1 = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var zipWithOf = function (g) {
    return function (f) {
        return Data_Newtype.unwrap()(g(f));
    };
};
var zipFWithOf = function (g) {
    return function (f) {
        return Data_Newtype.unwrap()(g(f));
    };
};
var withGrate = function (g) {
    return Data_Newtype.unwrap()(g(function (f) {
        return f(identity);
    }));
};
var grate = function (f) {
    return function (dictClosed) {
        var Profunctor0 = dictClosed.Profunctor0();
        return function (pab) {
            return Data_Profunctor.dimap(Profunctor0)(Data_Function.applyFlipped)(f)(Data_Profunctor_Closed.closed(dictClosed)(pab));
        };
    };
};
var cotraversed = function (dictDistributive) {
    return function (dictClosed) {
        return grate(function (f) {
            return Data_Distributive.cotraverse(dictDistributive)(Data_Functor.functorFn)(f)(identity1);
        })(dictClosed);
    };
};
var collectOf = function (dictFunctor) {
    return function (g) {
        return function (f) {
            var $8 = zipFWithOf(g)(identity1);
            var $9 = Data_Functor.map(dictFunctor)(f);
            return function ($10) {
                return $8($9($10));
            };
        };
    };
};
var cloneGrate = function (g) {
    return function (dictClosed) {
        return grate(withGrate(g))(dictClosed);
    };
};
export {
    grate,
    withGrate,
    cloneGrate,
    cotraversed,
    zipWithOf,
    zipFWithOf,
    collectOf
};
