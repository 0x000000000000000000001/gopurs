import * as $foreign from "./foreign.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Type_Proxy from "../Type.Proxy/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var discard = function (dict) {
    return dict.discard;
};
var bindProxy = {
    bind: function (v) {
        return function (v1) {
            return Type_Proxy["Proxy"].value;
        };
    },
    Apply0: function () {
        return Control_Apply.applyProxy;
    }
};
var bindFn = {
    bind: function (m) {
        return function (f) {
            return function (x) {
                return f(m(x))(x);
            };
        };
    },
    Apply0: function () {
        return Control_Apply.applyFn;
    }
};

// | The `bind`/`>>=` function for `Array` works by applying a function to
// | each element in the array, and flattening the results into a single,
// | new array.
// |
// | Array's `bind`/`>>=` works like a nested for loop. Each `bind` adds
// | another level of nesting in the loop. For example:
// | ```
// | foo :: Array String
// | foo =
// |   ["a", "b"] >>= \eachElementInArray1 ->
// |     ["c", "d"] >>= \eachElementInArray2
// |       pure (eachElementInArray1 <> eachElementInArray2)
// |
// | -- In other words...
// | foo
// | -- ... is the same as...
// | [ ("a" <> "c"), ("a" <> "d"), ("b" <> "c"), ("b" <> "d") ]
// | -- which simplifies to...
// | [ "ac", "ad", "bc", "bd" ]
// | ```
var bindArray = {
    bind: $foreign.arrayBind,
    Apply0: function () {
        return Control_Apply.applyArray;
    }
};
var bind = function (dict) {
    return dict.bind;
};

// | `bindFlipped` is `bind` with its arguments reversed. For example:
// |
// | ```purescript
// | print =<< random
// | ```
var bindFlipped = function (dictBind) {
    return Data_Function.flip(bind(dictBind));
};

// | Backwards Kleisli composition.
var composeKleisliFlipped = function (dictBind) {
    return function (f) {
        return function (g) {
            return function (a) {
                return bindFlipped(dictBind)(f)(g(a));
            };
        };
    };
};

// | Forwards Kleisli composition.
// |
// | For example:
// |
// | ```purescript
// | import Data.Array (head, tail)
// |
// | third = tail >=> tail >=> head
// | ```
var composeKleisli = function (dictBind) {
    return function (f) {
        return function (g) {
            return function (a) {
                return bind(dictBind)(f(a))(g);
            };
        };
    };
};
var discardProxy = {
    discard: function (dictBind) {
        return bind(dictBind);
    }
};
var discardUnit = {
    discard: function (dictBind) {
        return bind(dictBind);
    }
};

// | Execute a monadic action if a condition holds.
// |
// | For example:
// |
// | ```purescript
// | main = ifM ((< 0.5) <$> random)
// |          (trace "Heads")
// |          (trace "Tails")
// | ```
var ifM = function (dictBind) {
    return function (cond) {
        return function (t) {
            return function (f) {
                return bind(dictBind)(cond)(function (cond$prime) {
                    if (cond$prime) {
                        return t;
                    };
                    return f;
                });
            };
        };
    };
};

// | Collapse two applications of a monadic type constructor into one.
var join = function (dictBind) {
    return function (m) {
        return bind(dictBind)(m)(identity);
    };
};
export {
    bind,
    bindFlipped,
    discard,
    join,
    composeKleisli,
    composeKleisliFlipped,
    ifM,
    bindFn,
    bindArray,
    bindProxy,
    discardUnit,
    discardProxy
};
export {
    liftA1,
    pure,
    unless,
    when
} from "../Control.Applicative/index.js";
export {
    apply
} from "../Control.Apply/index.js";
export {
    map,
    void
} from "../Data.Functor/index.js";
