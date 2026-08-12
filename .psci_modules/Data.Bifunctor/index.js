import * as Control_Category from "../Control.Category/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var identity1 = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var bimap = function (dict) {
    return dict.bimap;
};

// | The bivoid function is used to ignore the types wrapped by a Bifunctor.
var bivoid = function (dictBifunctor) {
    return bimap(dictBifunctor)(Data_Function["const"](Data_Unit.unit))(Data_Function["const"](Data_Unit.unit));
};

// | Map a function over the first type argument of a `Bifunctor`.
var lmap = function (dictBifunctor) {
    return function (f) {
        return bimap(dictBifunctor)(f)(identity);
    };
};

// | Map a function over the second type arguments of a `Bifunctor`.
var rmap = function (dictBifunctor) {
    return bimap(dictBifunctor)(identity1);
};
var bifunctorTuple = {
    bimap: function (f) {
        return function (g) {
            return function (v) {
                return new Data_Tuple.Tuple(f(v.value0), g(v.value1));
            };
        };
    }
};
var bifunctorEither = {
    bimap: function (v) {
        return function (v1) {
            return function (v2) {
                if (v2 instanceof Data_Either.Left) {
                    return new Data_Either.Left(v(v2.value0));
                };
                if (v2 instanceof Data_Either.Right) {
                    return new Data_Either.Right(v1(v2.value0));
                };
                throw new Error("Failed pattern match at Data.Bifunctor (line 38, column 1 - line 40, column 36): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
            };
        };
    }
};
var bifunctorConst = {
    bimap: function (f) {
        return function (v) {
            return function (v1) {
                return f(v1);
            };
        };
    }
};
export {
    bimap,
    lmap,
    rmap,
    bivoid,
    bifunctorEither,
    bifunctorTuple,
    bifunctorConst
};
