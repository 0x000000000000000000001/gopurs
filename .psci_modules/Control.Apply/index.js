import * as $foreign from "./foreign.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Type_Proxy from "../Type.Proxy/index.js";
var applyProxy = {
    apply: function (v) {
        return function (v1) {
            return Type_Proxy["Proxy"].value;
        };
    },
    Functor0: function () {
        return Data_Functor.functorProxy;
    }
};
var applyFn = {
    apply: function (f) {
        return function (g) {
            return function (x) {
                return f(x)(g(x));
            };
        };
    },
    Functor0: function () {
        return Data_Functor.functorFn;
    }
};
var applyArray = {
    apply: $foreign.arrayApply,
    Functor0: function () {
        return Data_Functor.functorArray;
    }
};
var apply = function (dict) {
    return dict.apply;
};

// | Combine two effectful actions, keeping only the result of the first.
var applyFirst = function (dictApply) {
    var Functor0 = dictApply.Functor0();
    return function (a) {
        return function (b) {
            return apply(dictApply)(Data_Functor.map(Functor0)(Data_Function["const"])(a))(b);
        };
    };
};

// | Combine two effectful actions, keeping only the result of the second.
var applySecond = function (dictApply) {
    var Functor0 = dictApply.Functor0();
    return function (a) {
        return function (b) {
            return apply(dictApply)(Data_Functor.map(Functor0)(Data_Function["const"](Control_Category.identity(Control_Category.categoryFn)))(a))(b);
        };
    };
};

// | Lift a function of two arguments to a function which accepts and returns
// | values wrapped with the type constructor `f`.
// |
// | ```purescript
// | lift2 add (Just 1) (Just 2) == Just 3
// | lift2 add Nothing (Just 2) == Nothing
// |```
// |
var lift2 = function (dictApply) {
    var Functor0 = dictApply.Functor0();
    return function (f) {
        return function (a) {
            return function (b) {
                return apply(dictApply)(Data_Functor.map(Functor0)(f)(a))(b);
            };
        };
    };
};

// | Lift a function of three arguments to a function which accepts and returns
// | values wrapped with the type constructor `f`.
var lift3 = function (dictApply) {
    var Functor0 = dictApply.Functor0();
    return function (f) {
        return function (a) {
            return function (b) {
                return function (c) {
                    return apply(dictApply)(apply(dictApply)(Data_Functor.map(Functor0)(f)(a))(b))(c);
                };
            };
        };
    };
};

// | Lift a function of four arguments to a function which accepts and returns
// | values wrapped with the type constructor `f`.
var lift4 = function (dictApply) {
    var Functor0 = dictApply.Functor0();
    return function (f) {
        return function (a) {
            return function (b) {
                return function (c) {
                    return function (d) {
                        return apply(dictApply)(apply(dictApply)(apply(dictApply)(Data_Functor.map(Functor0)(f)(a))(b))(c))(d);
                    };
                };
            };
        };
    };
};

// | Lift a function of five arguments to a function which accepts and returns
// | values wrapped with the type constructor `f`.
var lift5 = function (dictApply) {
    var Functor0 = dictApply.Functor0();
    return function (f) {
        return function (a) {
            return function (b) {
                return function (c) {
                    return function (d) {
                        return function (e) {
                            return apply(dictApply)(apply(dictApply)(apply(dictApply)(apply(dictApply)(Data_Functor.map(Functor0)(f)(a))(b))(c))(d))(e);
                        };
                    };
                };
            };
        };
    };
};
export {
    apply,
    applyFirst,
    applySecond,
    lift2,
    lift3,
    lift4,
    lift5,
    applyFn,
    applyArray,
    applyProxy
};
export {
    map,
    void
} from "../Data.Functor/index.js";
