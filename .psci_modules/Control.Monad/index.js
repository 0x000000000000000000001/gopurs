import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Data_Functor from "../Data.Functor/index.js";

// | Perform a monadic action when a condition is true, where the conditional
// | value is also in a monadic context.
var whenM = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var Applicative0 = dictMonad.Applicative0();
    return function (mb) {
        return function (m) {
            return Control_Bind.bind(Bind1)(mb)(function (b) {
                return Control_Applicative.when(Applicative0)(b)(m);
            });
        };
    };
};

// | Perform a monadic action unless a condition is true, where the conditional
// | value is also in a monadic context.
var unlessM = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var Applicative0 = dictMonad.Applicative0();
    return function (mb) {
        return function (m) {
            return Control_Bind.bind(Bind1)(mb)(function (b) {
                return Control_Applicative.unless(Applicative0)(b)(m);
            });
        };
    };
};
var monadProxy = {
    Applicative0: function () {
        return Control_Applicative.applicativeProxy;
    },
    Bind1: function () {
        return Control_Bind.bindProxy;
    }
};
var monadFn = {
    Applicative0: function () {
        return Control_Applicative.applicativeFn;
    },
    Bind1: function () {
        return Control_Bind.bindFn;
    }
};
var monadArray = {
    Applicative0: function () {
        return Control_Applicative.applicativeArray;
    },
    Bind1: function () {
        return Control_Bind.bindArray;
    }
};

// | `liftM1` provides a default implementation of `(<$>)` for any
// | [`Monad`](#monad), without using `(<$>)` as provided by the
// | [`Functor`](#functor)-[`Monad`](#monad) superclass relationship.
// |
// | `liftM1` can therefore be used to write [`Functor`](#functor) instances
// | as follows:
// |
// | ```purescript
// | instance functorF :: Functor F where
// |   map = liftM1
// | ```
var liftM1 = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var Applicative0 = dictMonad.Applicative0();
    return function (f) {
        return function (a) {
            return Control_Bind.bind(Bind1)(a)(function (a$prime) {
                return Control_Applicative.pure(Applicative0)(f(a$prime));
            });
        };
    };
};

// | `ap` provides a default implementation of `(<*>)` for any `Monad`, without
// | using `(<*>)` as provided by the `Apply`-`Monad` superclass relationship.
// |
// | `ap` can therefore be used to write `Apply` instances as follows:
// |
// | ```purescript
// | instance applyF :: Apply F where
// |   apply = ap
// | ```
// Note: Only a `Bind` constraint is needed, but this can
// produce loops when used with other default implementations
// (i.e. `liftA1`).
// See https://github.com/purescript/purescript-prelude/issues/232
var ap = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var Applicative0 = dictMonad.Applicative0();
    return function (f) {
        return function (a) {
            return Control_Bind.bind(Bind1)(f)(function (f$prime) {
                return Control_Bind.bind(Bind1)(a)(function (a$prime) {
                    return Control_Applicative.pure(Applicative0)(f$prime(a$prime));
                });
            });
        };
    };
};
export {
    liftM1,
    whenM,
    unlessM,
    ap,
    monadFn,
    monadArray,
    monadProxy
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
    bind,
    ifM,
    join
} from "../Control.Bind/index.js";
export {
    map,
    void
} from "../Data.Functor/index.js";
