import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad from "../Control.Monad/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Effect from "../Effect/index.js";
import * as Effect_Ref from "../Effect.Ref/index.js";

// | The result of a computation: either `Loop` containing the updated
// | accumulator, or `Done` containing the final result of the computation.
var Loop = /* #__PURE__ */ (function () {
    function Loop(value0) {
        this.value0 = value0;
    };
    Loop.create = function (value0) {
        return new Loop(value0);
    };
    return Loop;
})();

// | The result of a computation: either `Loop` containing the updated
// | accumulator, or `Done` containing the final result of the computation.
var Done = /* #__PURE__ */ (function () {
    function Done(value0) {
        this.value0 = value0;
    };
    Done.create = function (value0) {
        return new Done(value0);
    };
    return Done;
})();
var tailRecM = function (dict) {
    return dict.tailRecM;
};

// | Create a tail-recursive function of two arguments which uses constant stack space.
// |
// | The `loop2` helper function provides a curried alternative to the `Loop`
// | constructor for this function.
var tailRecM2 = function (dictMonadRec) {
    return function (f) {
        return function (a) {
            return function (b) {
                return tailRecM(dictMonadRec)(function (o) {
                    return f(o.a)(o.b);
                })({
                    a: a,
                    b: b
                });
            };
        };
    };
};

// | Create a tail-recursive function of three arguments which uses constant stack space.
// |
// | The `loop3` helper function provides a curried alternative to the `Loop`
// | constructor for this function.
var tailRecM3 = function (dictMonadRec) {
    return function (f) {
        return function (a) {
            return function (b) {
                return function (c) {
                    return tailRecM(dictMonadRec)(function (o) {
                        return f(o.a)(o.b)(o.c);
                    })({
                        a: a,
                        b: b,
                        c: c
                    });
                };
            };
        };
    };
};

// | Supplied computation will be executed repeatedly until it evaluates
// | to `Just value` and then that `value` will be returned.
var untilJust = function (dictMonadRec) {
    var Functor0 = (((dictMonadRec.Monad0()).Bind1()).Apply0()).Functor0();
    return function (m) {
        return tailRecM(dictMonadRec)(function (v) {
            return Data_Functor.mapFlipped(Functor0)(m)(function (v1) {
                if (v1 instanceof Data_Maybe.Nothing) {
                    return new Loop(Data_Unit.unit);
                };
                if (v1 instanceof Data_Maybe.Just) {
                    return new Done(v1.value0);
                };
                throw new Error("Failed pattern match at Control.Monad.Rec.Class (line 179, column 43 - line 181, column 19): " + [ v1.constructor.name ]);
            });
        })(Data_Unit.unit);
    };
};

// | While supplied computation evaluates to `Just _`, it will be
// | executed repeatedly and results will be combined using monoid instance.
var whileJust = function (dictMonoid) {
    var Semigroup0 = dictMonoid.Semigroup0();
    return function (dictMonadRec) {
        var Functor0 = (((dictMonadRec.Monad0()).Bind1()).Apply0()).Functor0();
        return function (m) {
            return tailRecM(dictMonadRec)(function (v) {
                return Data_Functor.mapFlipped(Functor0)(m)(function (v1) {
                    if (v1 instanceof Data_Maybe.Nothing) {
                        return new Done(v);
                    };
                    if (v1 instanceof Data_Maybe.Just) {
                        return new Loop(Data_Semigroup.append(Semigroup0)(v)(v1.value0));
                    };
                    throw new Error("Failed pattern match at Control.Monad.Rec.Class (line 172, column 45 - line 174, column 26): " + [ v1.constructor.name ]);
                });
            })(Data_Monoid.mempty(dictMonoid));
        };
    };
};

// | Create a pure tail-recursive function of one argument
// |
// | For example:
// |
// | ```purescript
// | pow :: Int -> Int -> Int
// | pow n p = tailRec go { accum: 1, power: p }
// |   where
// |   go :: _ -> Step _ Int
// |   go { accum: acc, power: 0 } = Done acc
// |   go { accum: acc, power: p } = Loop { accum: acc * n, power: p - 1 }
// | ```
var tailRec = function (f) {
    var go = function ($copy_v) {
        var $tco_done = false;
        var $tco_result;
        function $tco_loop(v) {
            if (v instanceof Loop) {
                $copy_v = f(v.value0);
                return;
            };
            if (v instanceof Done) {
                $tco_done = true;
                return v.value0;
            };
            throw new Error("Failed pattern match at Control.Monad.Rec.Class (line 103, column 3 - line 103, column 25): " + [ v.constructor.name ]);
        };
        while (!$tco_done) {
            $tco_result = $tco_loop($copy_v);
        };
        return $tco_result;
    };
    return function ($68) {
        return go(f($68));
    };
};

// | Create a pure tail-recursive function of two arguments
// |
// | The `loop2` helper function provides a curried alternative to the `Loop`
// | constructor for this function.
var tailRec2 = function (f) {
    return function (a) {
        return function (b) {
            return tailRec(function (o) {
                return f(o.a)(o.b);
            })({
                a: a,
                b: b
            });
        };
    };
};

// | Create a pure tail-recursive function of three arguments
// |
// | The `loop3` helper function provides a curried alternative to the `Loop`
// | constructor for this function.
var tailRec3 = function (f) {
    return function (a) {
        return function (b) {
            return function (c) {
                return tailRec(function (o) {
                    return f(o.a)(o.b)(o.c);
                })({
                    a: a,
                    b: b,
                    c: c
                });
            };
        };
    };
};
var monadRecMaybe = {
    tailRecM: function (f) {
        return function (a0) {
            var g = function (v) {
                if (v instanceof Data_Maybe.Nothing) {
                    return new Done(Data_Maybe.Nothing.value);
                };
                if (v instanceof Data_Maybe.Just && v.value0 instanceof Loop) {
                    return new Loop(f(v.value0.value0));
                };
                if (v instanceof Data_Maybe.Just && v.value0 instanceof Done) {
                    return new Done(new Data_Maybe.Just(v.value0.value0));
                };
                throw new Error("Failed pattern match at Control.Monad.Rec.Class (line 153, column 7 - line 153, column 31): " + [ v.constructor.name ]);
            };
            return tailRec(g)(f(a0));
        };
    },
    Monad0: function () {
        return Data_Maybe.monadMaybe;
    }
};
var monadRecIdentity = {
    tailRecM: function (f) {
        var runIdentity = function (v) {
            return v;
        };
        var $69 = tailRec(function ($71) {
            return runIdentity(f($71));
        });
        return function ($70) {
            return Data_Identity.Identity($69($70));
        };
    },
    Monad0: function () {
        return Data_Identity.monadIdentity;
    }
};
var monadRecFunction = {
    tailRecM: function (f) {
        return function (a0) {
            return function (e) {
                return tailRec(function (a) {
                    return f(a)(e);
                })(a0);
            };
        };
    },
    Monad0: function () {
        return Control_Monad.monadFn;
    }
};
var monadRecEither = {
    tailRecM: function (f) {
        return function (a0) {
            var g = function (v) {
                if (v instanceof Data_Either.Left) {
                    return new Done(new Data_Either.Left(v.value0));
                };
                if (v instanceof Data_Either.Right && v.value0 instanceof Loop) {
                    return new Loop(f(v.value0.value0));
                };
                if (v instanceof Data_Either.Right && v.value0 instanceof Done) {
                    return new Done(new Data_Either.Right(v.value0.value0));
                };
                throw new Error("Failed pattern match at Control.Monad.Rec.Class (line 145, column 7 - line 145, column 33): " + [ v.constructor.name ]);
            };
            return tailRec(g)(f(a0));
        };
    },
    Monad0: function () {
        return Data_Either.monadEither;
    }
};
var monadRecEffect = {
    tailRecM: function (f) {
        return function (a) {
            var fromDone = function (v) {
                if (v instanceof Done) {
                    return v.value0;
                };
                throw new Error("Failed pattern match at Control.Monad.Rec.Class (line 137, column 30 - line 137, column 44): " + [ v.constructor.name ]);
            };
            return function __do() {
                var r = Control_Bind.bindFlipped(Effect.bindEffect)(Effect_Ref["new"])(f(a))();
                (function () {
                    while (!(function __do() {
                        var v = Effect_Ref.read(r)();
                        if (v instanceof Loop) {
                            var e = f(v.value0)();
                            Effect_Ref.write(e)(r)();
                            return false;
                        };
                        if (v instanceof Done) {
                            return true;
                        };
                        throw new Error("Failed pattern match at Control.Monad.Rec.Class (line 128, column 22 - line 133, column 28): " + [ v.constructor.name ]);
                    })()) {

                    };
                    return {};
                })();
                return Data_Functor.map(Effect.functorEffect)(fromDone)(Effect_Ref.read(r))();
            };
        };
    },
    Monad0: function () {
        return Effect.monadEffect;
    }
};

// | A curried version of the `Loop` constructor, provided as a convenience for
// | use with `tailRec3` and `tailRecM3`.
var loop3 = function (a) {
    return function (b) {
        return function (c) {
            return new Loop({
                a: a,
                b: b,
                c: c
            });
        };
    };
};

// | A curried version of the `Loop` constructor, provided as a convenience for
// | use with `tailRec2` and `tailRecM2`.
var loop2 = function (a) {
    return function (b) {
        return new Loop({
            a: a,
            b: b
        });
    };
};
var functorStep = {
    map: function (f) {
        return function (m) {
            if (m instanceof Loop) {
                return new Loop(m.value0);
            };
            if (m instanceof Done) {
                return new Done(f(m.value0));
            };
            throw new Error("Failed pattern match at Control.Monad.Rec.Class (line 0, column 0 - line 0, column 0): " + [ m.constructor.name ]);
        };
    }
};

// | `forever` runs an action indefinitely, using the `MonadRec` instance to
// | ensure constant stack usage.
// |
// | For example:
// |
// | ```purescript
// | main = forever $ trace "Hello, World!"
// | ```
var forever = function (dictMonadRec) {
    var Functor0 = (((dictMonadRec.Monad0()).Bind1()).Apply0()).Functor0();
    return function (ma) {
        return tailRecM(dictMonadRec)(function (u) {
            return Data_Functor.voidRight(Functor0)(new Loop(u))(ma);
        })(Data_Unit.unit);
    };
};
var bifunctorStep = {
    bimap: function (v) {
        return function (v1) {
            return function (v2) {
                if (v2 instanceof Loop) {
                    return new Loop(v(v2.value0));
                };
                if (v2 instanceof Done) {
                    return new Done(v1(v2.value0));
                };
                throw new Error("Failed pattern match at Control.Monad.Rec.Class (line 33, column 1 - line 35, column 34): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
            };
        };
    }
};
export {
    Loop,
    Done,
    tailRec,
    tailRec2,
    tailRec3,
    tailRecM,
    tailRecM2,
    tailRecM3,
    forever,
    whileJust,
    untilJust,
    loop2,
    loop3,
    functorStep,
    bifunctorStep,
    monadRecIdentity,
    monadRecEffect,
    monadRecFunction,
    monadRecEither,
    monadRecMaybe
};
