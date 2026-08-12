// | This module defines the reader-writer-state monad transformer, `RWST`.
import * as Control_Alt from "../Control.Alt/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad_Error_Class from "../Control.Monad.Error.Class/index.js";
import * as Control_Monad_Rec_Class from "../Control.Monad.Rec.Class/index.js";
import * as Control_Monad_ST_Class from "../Control.Monad.ST.Class/index.js";
import * as Control_Monad_Trans_Class from "../Control.Monad.Trans.Class/index.js";
import * as Control_Plus from "../Control.Plus/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Effect_Class from "../Effect.Class/index.js";
var RWSResult = /* #__PURE__ */ (function () {
    function RWSResult(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    RWSResult.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new RWSResult(value0, value1, value2);
            };
        };
    };
    return RWSResult;
})();

// | The reader-writer-state monad transformer, which combines the operations
// | of `ReaderT`, `WriterT` and `StateT` into a single monad transformer.
var RWST = function (x) {
    return x;
};

// | Change the context type in a `RWST` monad action.
var withRWST = function (f) {
    return function (m) {
        return function (r) {
            return function (s) {
                return Data_Tuple.uncurry(m)(f(r)(s));
            };
        };
    };
};

// | Run a computation in the `RWST` monad.
var runRWST = function (v) {
    return v;
};
var newtypeRWST = {
    Coercible0: function () {
        return undefined;
    }
};
var monadTransRWST = function (dictMonoid) {
    return {
        lift: function (dictMonad) {
            var Bind1 = dictMonad.Bind1();
            var pure = Control_Applicative.pure(dictMonad.Applicative0());
            return function (m) {
                return function (v) {
                    return function (s) {
                        return Control_Bind.bind(Bind1)(m)(function (a) {
                            return pure(new RWSResult(s, a, Data_Monoid.mempty(dictMonoid)));
                        });
                    };
                };
            };
        }
    };
};

// | Change the result and accumulator types in a `RWST` monad action.
var mapRWST = function (f) {
    return function (v) {
        return function (r) {
            return function (s) {
                return f(v(r)(s));
            };
        };
    };
};
var lazyRWST = {
    defer: function (f) {
        return function (r) {
            return function (s) {
                var v = f(Data_Unit.unit);
                return v(r)(s);
            };
        };
    }
};
var functorRWST = function (dictFunctor) {
    return {
        map: function (f) {
            return function (v) {
                return function (r) {
                    return function (s) {
                        return Data_Functor.map(dictFunctor)(function (v1) {
                            return new RWSResult(v1.value0, f(v1.value1), v1.value2);
                        })(v(r)(s));
                    };
                };
            };
        }
    };
};

// | Run a computation in the `RWST` monad, discarding the result.
var execRWST = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var Applicative0 = dictMonad.Applicative0();
    return function (v) {
        return function (r) {
            return function (s) {
                return Control_Bind.bind(Bind1)(v(r)(s))(function (v1) {
                    return Control_Applicative.pure(Applicative0)(new Data_Tuple.Tuple(v1.value0, v1.value2));
                });
            };
        };
    };
};

// | Run a computation in the `RWST` monad, discarding the final state.
var evalRWST = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var Applicative0 = dictMonad.Applicative0();
    return function (v) {
        return function (r) {
            return function (s) {
                return Control_Bind.bind(Bind1)(v(r)(s))(function (v1) {
                    return Control_Applicative.pure(Applicative0)(new Data_Tuple.Tuple(v1.value1, v1.value2));
                });
            };
        };
    };
};
var applyRWST = function (dictBind) {
    var Apply0 = dictBind.Apply0();
    var Functor0 = Apply0.Functor0();
    var functorRWST1 = functorRWST(Apply0.Functor0());
    return function (dictMonoid) {
        var Semigroup0 = dictMonoid.Semigroup0();
        return {
            apply: function (v) {
                return function (v1) {
                    return function (r) {
                        return function (s) {
                            return Control_Bind.bind(dictBind)(v(r)(s))(function (v2) {
                                return Data_Functor.mapFlipped(Functor0)(v1(r)(v2.value0))(function (v3) {
                                    return new RWSResult(v3.value0, v2.value1(v3.value1), Data_Semigroup.append(Semigroup0)(v2.value2)(v3.value2));
                                });
                            });
                        };
                    };
                };
            },
            Functor0: function () {
                return functorRWST1;
            }
        };
    };
};
var bindRWST = function (dictBind) {
    var Functor0 = (dictBind.Apply0()).Functor0();
    var applyRWST1 = applyRWST(dictBind);
    return function (dictMonoid) {
        var Semigroup0 = dictMonoid.Semigroup0();
        var applyRWST2 = applyRWST1(dictMonoid);
        return {
            bind: function (v) {
                return function (f) {
                    return function (r) {
                        return function (s) {
                            return Control_Bind.bind(dictBind)(v(r)(s))(function (v1) {
                                var v2 = f(v1.value1);
                                return Data_Functor.mapFlipped(Functor0)(v2(r)(v1.value0))(function (v3) {
                                    return new RWSResult(v3.value0, v3.value1, Data_Semigroup.append(Semigroup0)(v1.value2)(v3.value2));
                                });
                            });
                        };
                    };
                };
            },
            Apply0: function () {
                return applyRWST2;
            }
        };
    };
};
var semigroupRWST = function (dictBind) {
    var applyRWST1 = applyRWST(dictBind);
    return function (dictMonoid) {
        var applyRWST2 = applyRWST1(dictMonoid);
        return function (dictSemigroup) {
            return {
                append: Control_Apply.lift2(applyRWST2)(Data_Semigroup.append(dictSemigroup))
            };
        };
    };
};
var applicativeRWST = function (dictMonad) {
    var pure = Control_Applicative.pure(dictMonad.Applicative0());
    var applyRWST1 = applyRWST(dictMonad.Bind1());
    return function (dictMonoid) {
        var applyRWST2 = applyRWST1(dictMonoid);
        return {
            pure: function (a) {
                return function (v) {
                    return function (s) {
                        return pure(new RWSResult(s, a, Data_Monoid.mempty(dictMonoid)));
                    };
                };
            },
            Apply0: function () {
                return applyRWST2;
            }
        };
    };
};
var monadRWST = function (dictMonad) {
    var applicativeRWST1 = applicativeRWST(dictMonad);
    var bindRWST1 = bindRWST(dictMonad.Bind1());
    return function (dictMonoid) {
        var applicativeRWST2 = applicativeRWST1(dictMonoid);
        var bindRWST2 = bindRWST1(dictMonoid);
        return {
            Applicative0: function () {
                return applicativeRWST2;
            },
            Bind1: function () {
                return bindRWST2;
            }
        };
    };
};
var monadAskRWST = function (dictMonad) {
    var pure = Control_Applicative.pure(dictMonad.Applicative0());
    var monadRWST1 = monadRWST(dictMonad);
    return function (dictMonoid) {
        var monadRWST2 = monadRWST1(dictMonoid);
        return {
            ask: function (r) {
                return function (s) {
                    return pure(new RWSResult(s, r, Data_Monoid.mempty(dictMonoid)));
                };
            },
            Monad0: function () {
                return monadRWST2;
            }
        };
    };
};
var monadReaderRWST = function (dictMonad) {
    var monadAskRWST1 = monadAskRWST(dictMonad);
    return function (dictMonoid) {
        var monadAskRWST2 = monadAskRWST1(dictMonoid);
        return {
            local: function (f) {
                return function (m) {
                    return function (r) {
                        return function (s) {
                            return m(f(r))(s);
                        };
                    };
                };
            },
            MonadAsk0: function () {
                return monadAskRWST2;
            }
        };
    };
};
var monadEffectRWS = function (dictMonoid) {
    var lift = Control_Monad_Trans_Class.lift(monadTransRWST(dictMonoid));
    return function (dictMonadEffect) {
        var Monad0 = dictMonadEffect.Monad0();
        var monadRWST1 = monadRWST(Monad0)(dictMonoid);
        return {
            liftEffect: (function () {
                var $256 = lift(Monad0);
                var $257 = Effect_Class.liftEffect(dictMonadEffect);
                return function ($258) {
                    return $256($257($258));
                };
            })(),
            Monad0: function () {
                return monadRWST1;
            }
        };
    };
};
var monadRecRWST = function (dictMonadRec) {
    var Monad0 = dictMonadRec.Monad0();
    var Bind1 = Monad0.Bind1();
    var Applicative0 = Monad0.Applicative0();
    var monadRWST1 = monadRWST(Monad0);
    return function (dictMonoid) {
        var Semigroup0 = dictMonoid.Semigroup0();
        var monadRWST2 = monadRWST1(dictMonoid);
        return {
            tailRecM: function (k) {
                return function (a) {
                    var k$prime = function (r) {
                        return function (v) {
                            var v1 = k(v.value1);
                            return Control_Bind.bind(Bind1)(v1(r)(v.value0))(function (v2) {
                                return Control_Applicative.pure(Applicative0)((function () {
                                    if (v2.value1 instanceof Control_Monad_Rec_Class.Loop) {
                                        return new Control_Monad_Rec_Class.Loop(new RWSResult(v2.value0, v2.value1.value0, Data_Semigroup.append(Semigroup0)(v.value2)(v2.value2)));
                                    };
                                    if (v2.value1 instanceof Control_Monad_Rec_Class.Done) {
                                        return new Control_Monad_Rec_Class.Done(new RWSResult(v2.value0, v2.value1.value0, Data_Semigroup.append(Semigroup0)(v.value2)(v2.value2)));
                                    };
                                    throw new Error("Failed pattern match at Control.Monad.RWS.Trans (line 129, column 16 - line 131, column 68): " + [ v2.value1.constructor.name ]);
                                })());
                            });
                        };
                    };
                    return function (r) {
                        return function (s) {
                            return Control_Monad_Rec_Class.tailRecM(dictMonadRec)(k$prime(r))(new RWSResult(s, a, Data_Monoid.mempty(dictMonoid)));
                        };
                    };
                };
            },
            Monad0: function () {
                return monadRWST2;
            }
        };
    };
};
var monadStateRWST = function (dictMonad) {
    var pure = Control_Applicative.pure(dictMonad.Applicative0());
    var monadRWST1 = monadRWST(dictMonad);
    return function (dictMonoid) {
        var monadRWST2 = monadRWST1(dictMonoid);
        return {
            state: function (f) {
                return function (v) {
                    return function (s) {
                        var v1 = f(s);
                        return pure(new RWSResult(v1.value1, v1.value0, Data_Monoid.mempty(dictMonoid)));
                    };
                };
            },
            Monad0: function () {
                return monadRWST2;
            }
        };
    };
};
var monadTellRWST = function (dictMonad) {
    var pure = Control_Applicative.pure(dictMonad.Applicative0());
    var monadRWST1 = monadRWST(dictMonad);
    return function (dictMonoid) {
        var Semigroup0 = dictMonoid.Semigroup0();
        var monadRWST2 = monadRWST1(dictMonoid);
        return {
            tell: function (w) {
                return function (v) {
                    return function (s) {
                        return pure(new RWSResult(s, Data_Unit.unit, w));
                    };
                };
            },
            Semigroup0: function () {
                return Semigroup0;
            },
            Monad1: function () {
                return monadRWST2;
            }
        };
    };
};
var monadWriterRWST = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var Applicative0 = dictMonad.Applicative0();
    var pure = Control_Applicative.pure(Applicative0);
    var pure1 = Control_Applicative.pure(Applicative0);
    var monadTellRWST1 = monadTellRWST(dictMonad);
    return function (dictMonoid) {
        var monadTellRWST2 = monadTellRWST1(dictMonoid);
        return {
            listen: function (m) {
                return function (r) {
                    return function (s) {
                        return Control_Bind.bind(Bind1)(m(r)(s))(function (v) {
                            return pure(new RWSResult(v.value0, new Data_Tuple.Tuple(v.value1, v.value2), v.value2));
                        });
                    };
                };
            },
            pass: function (m) {
                return function (r) {
                    return function (s) {
                        return Control_Bind.bind(Bind1)(m(r)(s))(function (v) {
                            return pure1(new RWSResult(v.value0, v.value1.value0, v.value1.value1(v.value2)));
                        });
                    };
                };
            },
            Monoid0: function () {
                return dictMonoid;
            },
            MonadTell1: function () {
                return monadTellRWST2;
            }
        };
    };
};
var monadThrowRWST = function (dictMonadThrow) {
    var Monad0 = dictMonadThrow.Monad0();
    var monadRWST1 = monadRWST(dictMonadThrow.Monad0());
    return function (dictMonoid) {
        var monadTransRWST1 = monadTransRWST(dictMonoid);
        var monadRWST2 = monadRWST1(dictMonoid);
        return {
            throwError: function (e) {
                return Control_Monad_Trans_Class.lift(monadTransRWST1)(Monad0)(Control_Monad_Error_Class.throwError(dictMonadThrow)(e));
            },
            Monad0: function () {
                return monadRWST2;
            }
        };
    };
};
var monadErrorRWST = function (dictMonadError) {
    var monadThrowRWST1 = monadThrowRWST(dictMonadError.MonadThrow0());
    return function (dictMonoid) {
        var monadThrowRWST2 = monadThrowRWST1(dictMonoid);
        return {
            catchError: function (m) {
                return function (h) {
                    return function (r) {
                        return function (s) {
                            return Control_Monad_Error_Class.catchError(dictMonadError)(m(r)(s))(function (e) {
                                var v = h(e);
                                return v(r)(s);
                            });
                        };
                    };
                };
            },
            MonadThrow0: function () {
                return monadThrowRWST2;
            }
        };
    };
};
var monadSTRWST = function (dictMonoid) {
    var lift = Control_Monad_Trans_Class.lift(monadTransRWST(dictMonoid));
    return function (dictMonadST) {
        var Monad0 = dictMonadST.Monad0();
        var monadRWST1 = monadRWST(Monad0)(dictMonoid);
        return {
            liftST: (function () {
                var $259 = lift(Monad0);
                var $260 = Control_Monad_ST_Class.liftST(dictMonadST);
                return function ($261) {
                    return $259($260($261));
                };
            })(),
            Monad0: function () {
                return monadRWST1;
            }
        };
    };
};
var monoidRWST = function (dictMonad) {
    var applicativeRWST1 = applicativeRWST(dictMonad);
    var semigroupRWST1 = semigroupRWST(dictMonad.Bind1());
    return function (dictMonoid) {
        var applicativeRWST2 = applicativeRWST1(dictMonoid);
        var semigroupRWST2 = semigroupRWST1(dictMonoid);
        return function (dictMonoid1) {
            var semigroupRWST3 = semigroupRWST2(dictMonoid1.Semigroup0());
            return {
                mempty: Control_Applicative.pure(applicativeRWST2)(Data_Monoid.mempty(dictMonoid1)),
                Semigroup0: function () {
                    return semigroupRWST3;
                }
            };
        };
    };
};
var altRWST = function (dictAlt) {
    var functorRWST1 = functorRWST(dictAlt.Functor0());
    return {
        alt: function (v) {
            return function (v1) {
                return function (r) {
                    return function (s) {
                        return Control_Alt.alt(dictAlt)(v(r)(s))(v1(r)(s));
                    };
                };
            };
        },
        Functor0: function () {
            return functorRWST1;
        }
    };
};
var plusRWST = function (dictPlus) {
    var empty = Control_Plus.empty(dictPlus);
    var altRWST1 = altRWST(dictPlus.Alt0());
    return {
        empty: function (v) {
            return function (v1) {
                return empty;
            };
        },
        Alt0: function () {
            return altRWST1;
        }
    };
};
var alternativeRWST = function (dictMonoid) {
    return function (dictAlternative) {
        var plusRWST1 = plusRWST(dictAlternative.Plus1());
        return function (dictMonad) {
            var applicativeRWST1 = applicativeRWST(dictMonad)(dictMonoid);
            return {
                Applicative0: function () {
                    return applicativeRWST1;
                },
                Plus1: function () {
                    return plusRWST1;
                }
            };
        };
    };
};
export {
    RWSResult,
    RWST,
    runRWST,
    evalRWST,
    execRWST,
    mapRWST,
    withRWST,
    newtypeRWST,
    functorRWST,
    applyRWST,
    altRWST,
    alternativeRWST,
    bindRWST,
    applicativeRWST,
    monadRWST,
    monadTransRWST,
    lazyRWST,
    monadEffectRWS,
    monadAskRWST,
    monadReaderRWST,
    monadStateRWST,
    monadTellRWST,
    monadWriterRWST,
    monadThrowRWST,
    monadErrorRWST,
    monadRecRWST,
    plusRWST,
    semigroupRWST,
    monoidRWST,
    monadSTRWST
};
export {
    lift
} from "../Control.Monad.Trans.Class/index.js";
