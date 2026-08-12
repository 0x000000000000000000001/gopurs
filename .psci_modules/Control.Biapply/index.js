import * as Control_Category from "../Control.Category/index.js";
import * as Data_Bifunctor from "../Data.Bifunctor/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var biapplyTuple = {
    biapply: function (v) {
        return function (v1) {
            return new Data_Tuple.Tuple(v.value0(v1.value0), v.value1(v1.value1));
        };
    },
    Bifunctor0: function () {
        return Data_Bifunctor.bifunctorTuple;
    }
};
var biapply = function (dict) {
    return dict.biapply;
};

// | Keep the results of the second computation.
var biapplyFirst = function (dictBiapply) {
    var Bifunctor0 = dictBiapply.Bifunctor0();
    return function (a) {
        return function (b) {
            return biapply(dictBiapply)(Data_Bifunctor.bimap(Bifunctor0)(Data_Function["const"](Control_Category.identity(Control_Category.categoryFn)))(Data_Function["const"](Control_Category.identity(Control_Category.categoryFn)))(a))(b);
        };
    };
};

// | Keep the results of the first computation.
var biapplySecond = function (dictBiapply) {
    var Bifunctor0 = dictBiapply.Bifunctor0();
    return function (a) {
        return function (b) {
            return biapply(dictBiapply)(Data_Bifunctor.bimap(Bifunctor0)(Data_Function["const"])(Data_Function["const"])(a))(b);
        };
    };
};

// | Lift a function of two arguments.
var bilift2 = function (dictBiapply) {
    var Bifunctor0 = dictBiapply.Bifunctor0();
    return function (f) {
        return function (g) {
            return function (a) {
                return function (b) {
                    return biapply(dictBiapply)(Data_Bifunctor.bimap(Bifunctor0)(f)(g)(a))(b);
                };
            };
        };
    };
};

// | Lift a function of three arguments.
var bilift3 = function (dictBiapply) {
    var Bifunctor0 = dictBiapply.Bifunctor0();
    return function (f) {
        return function (g) {
            return function (a) {
                return function (b) {
                    return function (c) {
                        return biapply(dictBiapply)(biapply(dictBiapply)(Data_Bifunctor.bimap(Bifunctor0)(f)(g)(a))(b))(c);
                    };
                };
            };
        };
    };
};
export {
    biapply,
    biapplyFirst,
    biapplySecond,
    bilift2,
    bilift3,
    biapplyTuple
};
