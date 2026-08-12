import * as Control_Category from "../Control.Category/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var identity1 = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var wrap = /* #__PURE__ */ Data_Newtype.wrap();
var unwrap = /* #__PURE__ */ Data_Newtype.unwrap();
var profunctorFn = {
    dimap: function (a2b) {
        return function (c2d) {
            return function (b2c) {
                return function ($13) {
                    return c2d(b2c(a2b($13)));
                };
            };
        };
    }
};
var dimap = function (dict) {
    return dict.dimap;
};

// | Map a function over the (contravariant) first type argument only.
var lcmap = function (dictProfunctor) {
    return function (a2b) {
        return dimap(dictProfunctor)(a2b)(identity);
    };
};

// | Map a function over the (covariant) second type argument only.
var rmap = function (dictProfunctor) {
    return function (b2c) {
        return dimap(dictProfunctor)(identity1)(b2c);
    };
};
var unwrapIso = function (dictProfunctor) {
    return function () {
        return dimap(dictProfunctor)(wrap)(unwrap);
    };
};
var wrapIso = function (dictProfunctor) {
    return function () {
        return function (v) {
            return dimap(dictProfunctor)(unwrap)(wrap);
        };
    };
};

// | Lift a pure function into any `Profunctor` which is also a `Category`.
var arr = function (dictCategory) {
    var identity2 = Control_Category.identity(dictCategory);
    return function (dictProfunctor) {
        return function (f) {
            return rmap(dictProfunctor)(f)(identity2);
        };
    };
};
export {
    dimap,
    lcmap,
    rmap,
    arr,
    unwrapIso,
    wrapIso,
    profunctorFn
};
