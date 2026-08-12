import * as $foreign from "./foreign.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad from "../Control.Monad/index.js";
import * as Control_Monad_Error_Class from "../Control.Monad.Error.Class/index.js";
import * as Control_Monad_Rec_Class from "../Control.Monad.Rec.Class/index.js";
import * as Control_Monad_ST_Class from "../Control.Monad.ST.Class/index.js";
import * as Control_Parallel from "../Control.Parallel/index.js";
import * as Control_Parallel_Class from "../Control.Parallel.Class/index.js";
import * as Control_Plus from "../Control.Plus/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Time_Duration from "../Data.Time.Duration/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Effect from "../Effect/index.js";
import * as Effect_Class from "../Effect.Class/index.js";
import * as Effect_Exception from "../Effect.Exception/index.js";
import * as Effect_Unsafe from "../Effect.Unsafe/index.js";
import * as Partial_Unsafe from "../Partial.Unsafe/index.js";
import * as Unsafe_Coerce from "../Unsafe.Coerce/index.js";
var $runtime_lazy = function (name, moduleName, init) {
    var state = 0;
    var val;
    return function (lineNumber) {
        if (state === 2) return val;
        if (state === 1) throw new ReferenceError(name + " was needed before it finished initializing (module " + moduleName + ", line " + lineNumber + ")", moduleName, lineNumber);
        state = 1;
        val = init();
        state = 2;
        return val;
    };
};
var $$void = /* #__PURE__ */ Data_Functor["void"](Effect.functorEffect);
var void1 = /* #__PURE__ */ Data_Functor["void"](Effect.functorEffect);

// | Represents a forked computation by way of `forkAff`. `Fiber`s are
// | memoized, so their results are only computed once.
var Fiber = function (x) {
    return x;
};
var FFIUtil = function (x) {
    return x;
};

// | A cancellation effect for actions run via `makeAff`. If a `Fiber` is
// | killed, and an async action is pending, the canceler will be called to
// | clean it up.
var Canceler = function (x) {
    return x;
};

// | Suspends an `Aff` from within a parent `Aff` context, returning the `Fiber`.
// | A suspended `Aff` is not executed until a consumer observes the result
// | with `joinFiber`.
var suspendAff = /* #__PURE__ */ $foreign["_fork"](false);
var newtypeCanceler = {
    Coercible0: function () {
        return undefined;
    }
};
var functorParAff = {
    map: $foreign["_parAffMap"]
};
var functorAff = {
    map: $foreign["_map"]
};

// | Forks an `Aff` from within a parent `Aff` context, returning the `Fiber`.
var forkAff = /* #__PURE__ */ $foreign["_fork"](true);
var ffiUtil = /* #__PURE__ */ (function () {
    var unsafeFromRight = function (v) {
        if (v instanceof Data_Either.Right) {
            return v.value0;
        };
        if (v instanceof Data_Either.Left) {
            return Partial_Unsafe.unsafeCrashWith("unsafeFromRight: Left");
        };
        throw new Error("Failed pattern match at Effect.Aff (line 412, column 21 - line 414, column 54): " + [ v.constructor.name ]);
    };
    var unsafeFromLeft = function (v) {
        if (v instanceof Data_Either.Left) {
            return v.value0;
        };
        if (v instanceof Data_Either.Right) {
            return Partial_Unsafe.unsafeCrashWith("unsafeFromLeft: Right");
        };
        throw new Error("Failed pattern match at Effect.Aff (line 407, column 20 - line 409, column 55): " + [ v.constructor.name ]);
    };
    var isLeft = function (v) {
        if (v instanceof Data_Either.Left) {
            return true;
        };
        if (v instanceof Data_Either.Right) {
            return false;
        };
        throw new Error("Failed pattern match at Effect.Aff (line 402, column 12 - line 404, column 21): " + [ v.constructor.name ]);
    };
    return {
        isLeft: isLeft,
        fromLeft: unsafeFromLeft,
        fromRight: unsafeFromRight,
        left: Data_Either.Left.create,
        right: Data_Either.Right.create
    };
})();
var makeFiber = function (aff) {
    return $foreign["_makeFiber"](ffiUtil, aff);
};

// | Forks an `Aff` from an `Effect` context, returning the `Fiber`.
var launchAff = function (aff) {
    return function __do() {
        var fiber = makeFiber(aff)();
        fiber.run();
        return fiber;
    };
};

// | Forks an `Aff` from an `Effect` context, discarding the `Fiber`.
var launchAff_ = function ($55) {
    return $$void(launchAff($55));
};

// | Suspends an `Aff` from an `Effect` context, returning the `Fiber`.
var launchSuspendedAff = makeFiber;

// | Pauses the running fiber.
var delay = function (v) {
    return $foreign["_delay"](Data_Either.Right.create, v);
};

// | Guarantees resource acquisition and cleanup. The first effect may acquire
// | some resource, while the second will dispose of it. The third effect makes
// | use of the resource. Disposal is always run last, regardless. Neither
// | acquisition nor disposal may be cancelled and are guaranteed to run until
// | they complete.
var bracket = function (acquire) {
    return function (completed) {
        return $foreign.generalBracket(acquire)({
            killed: Data_Function["const"](completed),
            failed: Data_Function["const"](completed),
            completed: Data_Function["const"](completed)
        });
    };
};

// | Runs effects in parallel, combining their results.
var applyParAff = {
    apply: $foreign["_parAffApply"],
    Functor0: function () {
        return functorParAff;
    }
};
var semigroupParAff = function (dictSemigroup) {
    return {
        append: Control_Apply.lift2(applyParAff)(Data_Semigroup.append(dictSemigroup))
    };
};
var monadAff = {
    Applicative0: function () {
        return applicativeAff;
    },
    Bind1: function () {
        return bindAff;
    }
};
var bindAff = {
    bind: $foreign["_bind"],
    Apply0: function () {
        return $lazy_applyAff(0);
    }
};
var applicativeAff = {
    pure: $foreign["_pure"],
    Apply0: function () {
        return $lazy_applyAff(0);
    }
};
var $lazy_applyAff = /* #__PURE__ */ $runtime_lazy("applyAff", "Effect.Aff", function () {
    return {
        apply: Control_Monad.ap(monadAff),
        Functor0: function () {
            return functorAff;
        }
    };
});
var applyAff = /* #__PURE__ */ $lazy_applyAff(73);
var pure = /* #__PURE__ */ Control_Applicative.pure(applicativeAff);

// | Attaches a custom `Canceler` to an action. If the computation is canceled,
// | then the custom `Canceler` will be run afterwards.
var cancelWith = function (aff) {
    return function (v) {
        return $foreign.generalBracket(Control_Applicative.pure(applicativeAff)(Data_Unit.unit))({
            killed: function (e) {
                return function (v1) {
                    return v(e);
                };
            },
            failed: Data_Function["const"](Control_Applicative.pure(applicativeAff)),
            completed: Data_Function["const"](Control_Applicative.pure(applicativeAff))
        })(Data_Function["const"](aff));
    };
};

// | Runs the first effect after the second, regardless of whether it completed
// | successfully or the fiber was cancelled.
var $$finally = function (fin) {
    return function (a) {
        return bracket(Control_Applicative.pure(applicativeAff)(Data_Unit.unit))(Data_Function["const"](fin))(Data_Function["const"](a));
    };
};

// | Runs an effect such that it cannot be killed.
var invincible = function (a) {
    return bracket(a)(Data_Function["const"](Control_Applicative.pure(applicativeAff)(Data_Unit.unit)))(pure);
};
var lazyAff = {
    defer: function (f) {
        return Control_Bind.bind(bindAff)(Control_Applicative.pure(applicativeAff)(Data_Unit.unit))(f);
    }
};
var parallelAff = {
    parallel: Unsafe_Coerce.unsafeCoerce,
    sequential: $foreign["_sequential"],
    Apply0: function () {
        return applyAff;
    },
    Apply1: function () {
        return applyParAff;
    }
};
var applicativeParAff = {
    pure: /* #__PURE__ */ (function () {
        var $56 = Control_Parallel_Class.parallel(parallelAff);
        var $57 = Control_Applicative.pure(applicativeAff);
        return function ($58) {
            return $56($57($58));
        };
    })(),
    Apply0: function () {
        return applyParAff;
    }
};
var monoidParAff = function (dictMonoid) {
    var semigroupParAff1 = semigroupParAff(dictMonoid.Semigroup0());
    return {
        mempty: Control_Applicative.pure(applicativeParAff)(Data_Monoid.mempty(dictMonoid)),
        Semigroup0: function () {
            return semigroupParAff1;
        }
    };
};
var semigroupCanceler = {
    append: function (v) {
        return function (v1) {
            return function (err) {
                return Control_Parallel.parSequence_(parallelAff)(applicativeParAff)(Data_Foldable.foldableArray)([ v(err), v1(err) ]);
            };
        };
    }
};
var semigroupAff = function (dictSemigroup) {
    return {
        append: Control_Apply.lift2(applyAff)(Data_Semigroup.append(dictSemigroup))
    };
};
var monadEffectAff = {
    liftEffect: $foreign["_liftEffect"],
    Monad0: function () {
        return monadAff;
    }
};
var liftEffect = /* #__PURE__ */ Effect_Class.liftEffect(monadEffectAff);

// | A canceler from an Effect action.
var effectCanceler = function ($59) {
    return Canceler(Data_Function["const"](liftEffect($59)));
};

// | Blocks until the fiber completes, yielding the result. If the fiber
// | throws an exception, it is rethrown in the current fiber.
var joinFiber = function (v) {
    return $foreign.makeAff(function (k) {
        return Data_Functor.map(Effect.functorEffect)(effectCanceler)(v.join(k));
    });
};
var functorFiber = {
    map: function (f) {
        return function (t) {
            return Effect_Unsafe.unsafePerformEffect(makeFiber(Data_Functor.map(functorAff)(f)(joinFiber(t))));
        };
    }
};
var applyFiber = {
    apply: function (t1) {
        return function (t2) {
            return Effect_Unsafe.unsafePerformEffect(makeFiber(Control_Apply.apply(applyAff)(joinFiber(t1))(joinFiber(t2))));
        };
    },
    Functor0: function () {
        return functorFiber;
    }
};
var applicativeFiber = {
    pure: function (a) {
        return Effect_Unsafe.unsafePerformEffect(makeFiber(Control_Applicative.pure(applicativeAff)(a)));
    },
    Apply0: function () {
        return applyFiber;
    }
};

// | Invokes pending cancelers in a fiber and runs cleanup effects. Blocks
// | until the fiber has fully exited.
var killFiber = function (e) {
    return function (v) {
        return Control_Bind.bind(bindAff)(Effect_Class.liftEffect(monadEffectAff)(v.isSuspended))(function (suspended) {
            if (suspended) {
                return liftEffect(void1(v.kill(e, Data_Function["const"](Control_Applicative.pure(Effect.applicativeEffect)(Data_Unit.unit)))));
            };
            return $foreign.makeAff(function (k) {
                return Data_Functor.map(Effect.functorEffect)(effectCanceler)(v.kill(e, k));
            });
        });
    };
};

// | A canceler from a Fiber.
var fiberCanceler = /* #__PURE__ */ (function () {
    var $60 = Data_Function.flip(killFiber);
    return function ($61) {
        return Canceler($60($61));
    };
})();

// | Creates a new supervision context for some `Aff`, guaranteeing fiber
// | cleanup when the parent completes. Any pending fibers forked within
// | the context will be killed and have their cancelers run.
var supervise = function (aff) {
    var killError = Effect_Exception.error("[Aff] Child fiber outlived parent");
    var killAll = function (err) {
        return function (sup) {
            return $foreign.makeAff(function (k) {
                return $foreign["_killAll"](err, sup.supervisor, k(Control_Applicative.pure(Data_Either.applicativeEither)(Data_Unit.unit)));
            });
        };
    };
    var acquire = function __do() {
        var sup = $foreign["_makeSupervisedFiber"](ffiUtil, aff)();
        sup.fiber.run();
        return sup;
    };
    return $foreign.generalBracket(Effect_Class.liftEffect(monadEffectAff)(acquire))({
        killed: function (err) {
            return function (sup) {
                return Control_Parallel.parSequence_(parallelAff)(applicativeParAff)(Data_Foldable.foldableArray)([ killFiber(err)(sup.fiber), killAll(err)(sup) ]);
            };
        },
        failed: Data_Function["const"](killAll(killError)),
        completed: Data_Function["const"](killAll(killError))
    })(function ($62) {
        return joinFiber((function (v) {
            return v.fiber;
        })($62));
    });
};
var monadSTAff = {
    liftST: /* #__PURE__ */ (function () {
        var $63 = Effect_Class.liftEffect(monadEffectAff);
        var $64 = Control_Monad_ST_Class.liftST(Control_Monad_ST_Class.monadSTEffect);
        return function ($65) {
            return $63($64($65));
        };
    })(),
    Monad0: function () {
        return monadAff;
    }
};
var monadThrowAff = {
    throwError: $foreign["_throwError"],
    Monad0: function () {
        return monadAff;
    }
};
var monadErrorAff = {
    catchError: $foreign["_catchError"],
    MonadThrow0: function () {
        return monadThrowAff;
    }
};

// | A monomorphic version of `try`. Catches thrown errors and lifts them
// | into an `Either`.
var attempt = /* #__PURE__ */ Control_Monad_Error_Class["try"](monadErrorAff);

// | Forks an `Aff` from an `Effect` context and also takes a callback to run when
// | it completes. Returns the pending `Fiber`.
var runAff = function (k) {
    return function (aff) {
        return launchAff(Control_Bind.bindFlipped(bindAff)(function ($66) {
            return liftEffect(k($66));
        })(Control_Monad_Error_Class["try"](monadErrorAff)(aff)));
    };
};

// | Forks an `Aff` from an `Effect` context and also takes a callback to run when
// | it completes, discarding the `Fiber`.
var runAff_ = function (k) {
    return function (aff) {
        return $$void(runAff(k)(aff));
    };
};

// | Suspends an `Aff` from an `Effect` context and also takes a callback to run
// | when it completes. Returns the suspended `Fiber`.
var runSuspendedAff = function (k) {
    return function (aff) {
        return launchSuspendedAff(Control_Bind.bindFlipped(bindAff)(function ($67) {
            return liftEffect(k($67));
        })(Control_Monad_Error_Class["try"](monadErrorAff)(aff)));
    };
};

// | This instance is provided for compatibility. `Aff` is always stack-safe
// | within a given fiber. This instance will just result in unnecessary
// | bind overhead.
var monadRecAff = {
    tailRecM: function (k) {
        var go = function (a) {
            return Control_Bind.bind(bindAff)(k(a))(function (res) {
                if (res instanceof Control_Monad_Rec_Class.Done) {
                    return Control_Applicative.pure(applicativeAff)(res.value0);
                };
                if (res instanceof Control_Monad_Rec_Class.Loop) {
                    return go(res.value0);
                };
                throw new Error("Failed pattern match at Effect.Aff (line 104, column 7 - line 106, column 23): " + [ res.constructor.name ]);
            });
        };
        return go;
    },
    Monad0: function () {
        return monadAff;
    }
};
var monoidAff = function (dictMonoid) {
    var semigroupAff1 = semigroupAff(dictMonoid.Semigroup0());
    return {
        mempty: Control_Applicative.pure(applicativeAff)(Data_Monoid.mempty(dictMonoid)),
        Semigroup0: function () {
            return semigroupAff1;
        }
    };
};

// | A canceler which does not cancel anything.
var nonCanceler = /* #__PURE__ */ Data_Function["const"](/* #__PURE__ */ Control_Applicative.pure(applicativeAff)(Data_Unit.unit));

// | A no-op `Canceler` can be constructed with `mempty`.
var monoidCanceler = {
    mempty: nonCanceler,
    Semigroup0: function () {
        return semigroupCanceler;
    }
};

// | An async computation which does not resolve.
var never = /* #__PURE__ */ $foreign.makeAff(function (v) {
    return Control_Applicative.pure(Effect.applicativeEffect)(Data_Monoid.mempty(monoidCanceler));
});

// | Ignores any errors.
var apathize = /* #__PURE__ */ (function () {
    var $68 = Data_Functor.map(functorAff)(Data_Function["const"](Data_Unit.unit));
    return function ($69) {
        return $68(attempt($69));
    };
})();

// | Races effects in parallel. Returns the first successful result or the
// | first error if all fail with an exception. Losing branches will be
// | cancelled.
var altParAff = {
    alt: $foreign["_parAffAlt"],
    Functor0: function () {
        return functorParAff;
    }
};
var altAff = {
    alt: function (a1) {
        return function (a2) {
            return Control_Monad_Error_Class.catchError(monadErrorAff)(a1)(Data_Function["const"](a2));
        };
    },
    Functor0: function () {
        return functorAff;
    }
};
var plusAff = {
    empty: /* #__PURE__ */ Control_Monad_Error_Class.throwError(monadThrowAff)(/* #__PURE__ */ Effect_Exception.error("Always fails")),
    Alt0: function () {
        return altAff;
    }
};
var plusParAff = {
    empty: /* #__PURE__ */ Control_Parallel_Class.parallel(parallelAff)(/* #__PURE__ */ Control_Plus.empty(plusAff)),
    Alt0: function () {
        return altParAff;
    }
};
var alternativeParAff = {
    Applicative0: function () {
        return applicativeParAff;
    },
    Plus1: function () {
        return plusParAff;
    }
};
export {
    makeAff,
    generalBracket
} from "./foreign.js";
export {
    Canceler,
    launchAff,
    launchAff_,
    launchSuspendedAff,
    runAff,
    runAff_,
    runSuspendedAff,
    forkAff,
    suspendAff,
    supervise,
    attempt,
    apathize,
    delay,
    never,
    $$finally as finally,
    invincible,
    killFiber,
    joinFiber,
    cancelWith,
    bracket,
    nonCanceler,
    effectCanceler,
    fiberCanceler,
    functorAff,
    applyAff,
    applicativeAff,
    bindAff,
    monadAff,
    semigroupAff,
    monoidAff,
    altAff,
    plusAff,
    monadRecAff,
    monadThrowAff,
    monadErrorAff,
    monadEffectAff,
    lazyAff,
    monadSTAff,
    functorParAff,
    applyParAff,
    applicativeParAff,
    semigroupParAff,
    monoidParAff,
    altParAff,
    plusParAff,
    alternativeParAff,
    parallelAff,
    functorFiber,
    applyFiber,
    applicativeFiber,
    newtypeCanceler,
    semigroupCanceler,
    monoidCanceler
};
export {
    catchError,
    throwError,
    try
} from "../Control.Monad.Error.Class/index.js";
export {
    parallel,
    sequential
} from "../Control.Parallel.Class/index.js";
export {
    Milliseconds
} from "../Data.Time.Duration/index.js";
export {
    error,
    message
} from "../Effect.Exception/index.js";
