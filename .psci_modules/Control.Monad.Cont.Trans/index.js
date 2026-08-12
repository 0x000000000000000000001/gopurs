// | This module defines the CPS monad transformer.
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad_Cont_Class from "../Control.Monad.Cont.Class/index.js";
import * as Control_Monad_Reader_Class from "../Control.Monad.Reader.Class/index.js";
import * as Control_Monad_ST_Class from "../Control.Monad.ST.Class/index.js";
import * as Control_Monad_State_Class from "../Control.Monad.State.Class/index.js";
import * as Control_Monad_Trans_Class from "../Control.Monad.Trans.Class/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Effect_Class from "../Effect.Class/index.js";
var ContT = function (x) {
    return x;
};

// | Modify the continuation in a `ContT` monad action
var withContT = function (f) {
    return function (v) {
        return function (k) {
            return v(f(k));
        };
    };
};

// | Run a computation in the `ContT` monad, by providing a continuation.
var runContT = function (v) {
    return function (k) {
        return v(k);
    };
};
var newtypeContT = {
    Coercible0: function () {
        return undefined;
    }
};
var monadTransContT = {
    lift: function (dictMonad) {
        var Bind1 = dictMonad.Bind1();
        return function (m) {
            return function (k) {
                return Control_Bind.bind(Bind1)(m)(k);
            };
        };
    }
};
var lift = /* #__PURE__ */ Control_Monad_Trans_Class.lift(monadTransContT);

// | Modify the underlying action in a `ContT` monad action.
var mapContT = function (f) {
    return function (v) {
        return function (k) {
            return f(v(k));
        };
    };
};
var functorContT = function (dictFunctor) {
    return {
        map: function (f) {
            return function (v) {
                return function (k) {
                    return v(function (a) {
                        return k(f(a));
                    });
                };
            };
        }
    };
};
var applyContT = function (dictApply) {
    var functorContT1 = functorContT(dictApply.Functor0());
    return {
        apply: function (v) {
            return function (v1) {
                return function (k) {
                    return v(function (g) {
                        return v1(function (a) {
                            return k(g(a));
                        });
                    });
                };
            };
        },
        Functor0: function () {
            return functorContT1;
        }
    };
};
var bindContT = function (dictBind) {
    var applyContT1 = applyContT(dictBind.Apply0());
    return {
        bind: function (v) {
            return function (k) {
                return function (k$prime) {
                    return v(function (a) {
                        var v1 = k(a);
                        return v1(k$prime);
                    });
                };
            };
        },
        Apply0: function () {
            return applyContT1;
        }
    };
};
var semigroupContT = function (dictApply) {
    var applyContT1 = applyContT(dictApply);
    return function (dictSemigroup) {
        return {
            append: Control_Apply.lift2(applyContT1)(Data_Semigroup.append(dictSemigroup))
        };
    };
};
var applicativeContT = function (dictApplicative) {
    var applyContT1 = applyContT(dictApplicative.Apply0());
    return {
        pure: function (a) {
            return function (k) {
                return k(a);
            };
        },
        Apply0: function () {
            return applyContT1;
        }
    };
};
var monadContT = function (dictMonad) {
    var applicativeContT1 = applicativeContT(dictMonad.Applicative0());
    var bindContT1 = bindContT(dictMonad.Bind1());
    return {
        Applicative0: function () {
            return applicativeContT1;
        },
        Bind1: function () {
            return bindContT1;
        }
    };
};
var monadAskContT = function (dictMonadAsk) {
    var monadContT1 = monadContT(dictMonadAsk.Monad0());
    return {
        ask: Control_Monad_Trans_Class.lift(monadTransContT)(dictMonadAsk.Monad0())(Control_Monad_Reader_Class.ask(dictMonadAsk)),
        Monad0: function () {
            return monadContT1;
        }
    };
};
var monadReaderContT = function (dictMonadReader) {
    var MonadAsk0 = dictMonadReader.MonadAsk0();
    var Bind1 = (MonadAsk0.Monad0()).Bind1();
    var ask = Control_Monad_Reader_Class.ask(MonadAsk0);
    var monadAskContT1 = monadAskContT(MonadAsk0);
    return {
        local: function (f) {
            return function (v) {
                return function (k) {
                    return Control_Bind.bind(Bind1)(ask)(function (r) {
                        return Control_Monad_Reader_Class.local(dictMonadReader)(f)(v((function () {
                            var $87 = Control_Monad_Reader_Class.local(dictMonadReader)(Data_Function["const"](r));
                            return function ($88) {
                                return $87(k($88));
                            };
                        })()));
                    });
                };
            };
        },
        MonadAsk0: function () {
            return monadAskContT1;
        }
    };
};
var monadContContT = function (dictMonad) {
    var monadContT1 = monadContT(dictMonad);
    return {
        callCC: function (f) {
            return function (k) {
                var v = f(function (a) {
                    return function (v1) {
                        return k(a);
                    };
                });
                return v(k);
            };
        },
        Monad0: function () {
            return monadContT1;
        }
    };
};
var monadEffectContT = function (dictMonadEffect) {
    var Monad0 = dictMonadEffect.Monad0();
    var monadContT1 = monadContT(Monad0);
    return {
        liftEffect: (function () {
            var $89 = lift(Monad0);
            var $90 = Effect_Class.liftEffect(dictMonadEffect);
            return function ($91) {
                return $89($90($91));
            };
        })(),
        Monad0: function () {
            return monadContT1;
        }
    };
};
var monadStateContT = function (dictMonadState) {
    var Monad0 = dictMonadState.Monad0();
    var monadContT1 = monadContT(Monad0);
    return {
        state: (function () {
            var $92 = lift(Monad0);
            var $93 = Control_Monad_State_Class.state(dictMonadState);
            return function ($94) {
                return $92($93($94));
            };
        })(),
        Monad0: function () {
            return monadContT1;
        }
    };
};
var monadSTContT = function (dictMonadST) {
    var Monad0 = dictMonadST.Monad0();
    var monadContT1 = monadContT(Monad0);
    return {
        liftST: (function () {
            var $95 = lift(Monad0);
            var $96 = Control_Monad_ST_Class.liftST(dictMonadST);
            return function ($97) {
                return $95($96($97));
            };
        })(),
        Monad0: function () {
            return monadContT1;
        }
    };
};
var monoidContT = function (dictApplicative) {
    var applicativeContT1 = applicativeContT(dictApplicative);
    var semigroupContT1 = semigroupContT(dictApplicative.Apply0());
    return function (dictMonoid) {
        var semigroupContT2 = semigroupContT1(dictMonoid.Semigroup0());
        return {
            mempty: Control_Applicative.pure(applicativeContT1)(Data_Monoid.mempty(dictMonoid)),
            Semigroup0: function () {
                return semigroupContT2;
            }
        };
    };
};
export {
    ContT,
    runContT,
    mapContT,
    withContT,
    newtypeContT,
    monadContContT,
    functorContT,
    applyContT,
    applicativeContT,
    bindContT,
    monadContT,
    monadTransContT,
    monadEffectContT,
    monadAskContT,
    monadReaderContT,
    monadStateContT,
    semigroupContT,
    monoidContT,
    monadSTContT
};
export {
    callCC
} from "../Control.Monad.Cont.Class/index.js";
export {
    lift
} from "../Control.Monad.Trans.Class/index.js";
