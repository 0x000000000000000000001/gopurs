import * as $foreign from "./foreign.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var extendFn = function (dictSemigroup) {
    return {
        extend: function (f) {
            return function (g) {
                return function (w) {
                    return f(function (w$prime) {
                        return g(Data_Semigroup.append(dictSemigroup)(w)(w$prime));
                    });
                };
            };
        },
        Functor0: function () {
            return Data_Functor.functorFn;
        }
    };
};
var extendArray = {
    extend: $foreign.arrayExtend,
    Functor0: function () {
        return Data_Functor.functorArray;
    }
};
var extend = function (dict) {
    return dict.extend;
};

// | A version of `extend` with its arguments flipped.
var extendFlipped = function (dictExtend) {
    return function (w) {
        return function (f) {
            return extend(dictExtend)(f)(w);
        };
    };
};

// | Duplicate a comonadic context.
// |
// | `duplicate` is dual to `Control.Bind.join`.
var duplicate = function (dictExtend) {
    return extend(dictExtend)(identity);
};

// | Backwards co-Kleisli composition.
var composeCoKleisliFlipped = function (dictExtend) {
    return function (f) {
        return function (g) {
            return function (w) {
                return f(extend(dictExtend)(g)(w));
            };
        };
    };
};

// | Forwards co-Kleisli composition.
var composeCoKleisli = function (dictExtend) {
    return function (f) {
        return function (g) {
            return function (w) {
                return g(extend(dictExtend)(f)(w));
            };
        };
    };
};
export {
    extend,
    extendFlipped,
    composeCoKleisli,
    composeCoKleisliFlipped,
    duplicate,
    extendFn,
    extendArray
};
export {
    map,
    void
} from "../Data.Functor/index.js";
