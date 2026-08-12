// | This module defines the `MonadError` type class and its instances.
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Effect from "../Effect/index.js";
import * as Effect_Exception from "../Effect.Exception/index.js";
var throwError = function (dict) {
    return dict.throwError;
};
var monadThrowMaybe = /* #__PURE__ */ (function () {
    return {
        throwError: Data_Function["const"](Data_Maybe.Nothing.value),
        Monad0: function () {
            return Data_Maybe.monadMaybe;
        }
    };
})();
var monadThrowEither = /* #__PURE__ */ (function () {
    return {
        throwError: Data_Either.Left.create,
        Monad0: function () {
            return Data_Either.monadEither;
        }
    };
})();
var monadThrowEffect = {
    throwError: Effect_Exception.throwException,
    Monad0: function () {
        return Effect.monadEffect;
    }
};
var monadErrorMaybe = {
    catchError: function (v) {
        return function (v1) {
            if (v instanceof Data_Maybe.Nothing) {
                return v1(Data_Unit.unit);
            };
            if (v instanceof Data_Maybe.Just) {
                return new Data_Maybe.Just(v.value0);
            };
            throw new Error("Failed pattern match at Control.Monad.Error.Class (line 79, column 1 - line 81, column 33): " + [ v.constructor.name, v1.constructor.name ]);
        };
    },
    MonadThrow0: function () {
        return monadThrowMaybe;
    }
};
var monadErrorEither = {
    catchError: function (v) {
        return function (v1) {
            if (v instanceof Data_Either.Left) {
                return v1(v.value0);
            };
            if (v instanceof Data_Either.Right) {
                return new Data_Either.Right(v.value0);
            };
            throw new Error("Failed pattern match at Control.Monad.Error.Class (line 72, column 1 - line 74, column 35): " + [ v.constructor.name, v1.constructor.name ]);
        };
    },
    MonadThrow0: function () {
        return monadThrowEither;
    }
};
var monadErrorEffect = {
    catchError: /* #__PURE__ */ Data_Function.flip(Effect_Exception.catchException),
    MonadThrow0: function () {
        return monadThrowEffect;
    }
};

// | Lift a `Maybe` value to a MonadThrow monad.
var liftMaybe = function (dictMonadThrow) {
    var pure = Control_Applicative.pure((dictMonadThrow.Monad0()).Applicative0());
    return function (error) {
        return Data_Maybe.maybe(throwError(dictMonadThrow)(error))(pure);
    };
};

// | Lift an `Either` value to a MonadThrow monad.
var liftEither = function (dictMonadThrow) {
    return Data_Either.either(throwError(dictMonadThrow))(Control_Applicative.pure((dictMonadThrow.Monad0()).Applicative0()));
};
var catchError = function (dict) {
    return dict.catchError;
};

// | This function allows you to provide a predicate for selecting the
// | exceptions that you're interested in, and handle only those exceptons.
// | If the inner computation throws an exception, and the predicate returns
// | Nothing, then the whole computation will still fail with that exception.
var catchJust = function (dictMonadError) {
    var MonadThrow0 = dictMonadError.MonadThrow0();
    return function (p) {
        return function (act) {
            return function (handler) {
                var handle = function (e) {
                    var v = p(e);
                    if (v instanceof Data_Maybe.Nothing) {
                        return throwError(MonadThrow0)(e);
                    };
                    if (v instanceof Data_Maybe.Just) {
                        return handler(v.value0);
                    };
                    throw new Error("Failed pattern match at Control.Monad.Error.Class (line 57, column 5 - line 59, column 26): " + [ v.constructor.name ]);
                };
                return catchError(dictMonadError)(act)(handle);
            };
        };
    };
};

// | Return `Right` if the given action succeeds, `Left` if it throws.
var $$try = function (dictMonadError) {
    var Monad0 = (dictMonadError.MonadThrow0()).Monad0();
    var Functor0 = ((Monad0.Bind1()).Apply0()).Functor0();
    var pure = Control_Applicative.pure(Monad0.Applicative0());
    return function (a) {
        return catchError(dictMonadError)(Data_Functor.map(Functor0)(Data_Either.Right.create)(a))(function ($43) {
            return pure(Data_Either.Left.create($43));
        });
    };
};

// | Make sure that a resource is cleaned up in the event of an exception. The
// | release action is called regardless of whether the body action throws or
// | returns.
var withResource = function (dictMonadError) {
    var MonadThrow0 = dictMonadError.MonadThrow0();
    var Monad0 = MonadThrow0.Monad0();
    var Bind1 = Monad0.Bind1();
    var try1 = $$try(dictMonadError);
    var throwError1 = throwError(MonadThrow0);
    var pure = Control_Applicative.pure(Monad0.Applicative0());
    return function (acquire) {
        return function (release) {
            return function (kleisli) {
                return Control_Bind.bind(Bind1)(acquire)(function (resource) {
                    return Control_Bind.bind(Bind1)(try1(kleisli(resource)))(function (result) {
                        return Control_Bind.discard(Control_Bind.discardUnit)(Bind1)(release(resource))(function () {
                            return Data_Either.either(throwError1)(pure)(result);
                        });
                    });
                });
            };
        };
    };
};
export {
    catchError,
    throwError,
    catchJust,
    $$try as try,
    withResource,
    liftMaybe,
    liftEither,
    monadThrowEither,
    monadErrorEither,
    monadThrowMaybe,
    monadErrorMaybe,
    monadThrowEffect,
    monadErrorEffect
};
