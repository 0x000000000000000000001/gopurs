// | This module defines the _exception monad transformer_ `ExceptT`.
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Monad from "../Control.Monad/index.js";
import * as Control_Monad_Cont_Class from "../Control.Monad.Cont.Class/index.js";
import * as Control_Monad_Error_Class from "../Control.Monad.Error.Class/index.js";
import * as Control_Monad_Reader_Class from "../Control.Monad.Reader.Class/index.js";
import * as Control_Monad_Rec_Class from "../Control.Monad.Rec.Class/index.js";
import * as Control_Monad_ST_Class from "../Control.Monad.ST.Class/index.js";
import * as Control_Monad_State_Class from "../Control.Monad.State.Class/index.js";
import * as Control_Monad_Trans_Class from "../Control.Monad.Trans.Class/index.js";
import * as Control_Monad_Writer_Class from "../Control.Monad.Writer.Class/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Effect_Class from "../Effect.Class/index.js";

// | A monad transformer which adds exceptions to other monads, in the same way
// | as `Except`. As before, `e` is the type of exceptions, and `a` is the type
// | of successful results. The new type parameter `m` is the inner monad that
// | computations run in.
var ExceptT = function (x) {
    return x;
};

// | Transform any exceptions thrown by an `ExceptT` computation using the given function.
var withExceptT = function (dictFunctor) {
    return function (f) {
        return function (v) {
            var mapLeft = function (v1) {
                return function (v2) {
                    if (v2 instanceof Data_Either.Right) {
                        return new Data_Either.Right(v2.value0);
                    };
                    if (v2 instanceof Data_Either.Left) {
                        return new Data_Either.Left(v1(v2.value0));
                    };
                    throw new Error("Failed pattern match at Control.Monad.Except.Trans (line 43, column 3 - line 43, column 32): " + [ v1.constructor.name, v2.constructor.name ]);
                };
            };
            return Data_Functor.map(dictFunctor)(mapLeft(f))(v);
        };
    };
};

// | The inverse of `ExceptT`. Run a computation in the `ExceptT` monad.
var runExceptT = function (v) {
    return v;
};
var newtypeExceptT = {
    Coercible0: function () {
        return undefined;
    }
};
var monadTransExceptT = {
    lift: function (dictMonad) {
        var Bind1 = dictMonad.Bind1();
        var pure = Control_Applicative.pure(dictMonad.Applicative0());
        return function (m) {
            return Control_Bind.bind(Bind1)(m)(function (a) {
                return pure(new Data_Either.Right(a));
            });
        };
    }
};
var lift = /* #__PURE__ */ Control_Monad_Trans_Class.lift(monadTransExceptT);

// | Transform the unwrapped computation using the given function.
var mapExceptT = function (f) {
    return function (v) {
        return f(v);
    };
};
var functorExceptT = function (dictFunctor) {
    return {
        map: function (f) {
            return mapExceptT(Data_Functor.map(dictFunctor)(Data_Functor.map(Data_Either.functorEither)(f)));
        }
    };
};

// | Construct a computation in the `ExceptT` transformer from an `Either` value.
var except = function (dictApplicative) {
    var $169 = Control_Applicative.pure(dictApplicative);
    return function ($170) {
        return ExceptT($169($170));
    };
};
var monadExceptT = function (dictMonad) {
    return {
        Applicative0: function () {
            return applicativeExceptT(dictMonad);
        },
        Bind1: function () {
            return bindExceptT(dictMonad);
        }
    };
};
var bindExceptT = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var pure = Control_Applicative.pure(dictMonad.Applicative0());
    return {
        bind: function (v) {
            return function (k) {
                return Control_Bind.bind(Bind1)(v)(Data_Either.either(function ($171) {
                    return pure(Data_Either.Left.create($171));
                })(function (a) {
                    var v1 = k(a);
                    return v1;
                }));
            };
        },
        Apply0: function () {
            return applyExceptT(dictMonad);
        }
    };
};
var applyExceptT = function (dictMonad) {
    var functorExceptT1 = functorExceptT(((dictMonad.Bind1()).Apply0()).Functor0());
    return {
        apply: Control_Monad.ap(monadExceptT(dictMonad)),
        Functor0: function () {
            return functorExceptT1;
        }
    };
};
var applicativeExceptT = function (dictMonad) {
    return {
        pure: (function () {
            var $172 = Control_Applicative.pure(dictMonad.Applicative0());
            return function ($173) {
                return ExceptT($172(Data_Either.Right.create($173)));
            };
        })(),
        Apply0: function () {
            return applyExceptT(dictMonad);
        }
    };
};
var semigroupExceptT = function (dictMonad) {
    var applyExceptT1 = applyExceptT(dictMonad);
    return function (dictSemigroup) {
        return {
            append: Control_Apply.lift2(applyExceptT1)(Data_Semigroup.append(dictSemigroup))
        };
    };
};
var monadAskExceptT = function (dictMonadAsk) {
    var monadExceptT1 = monadExceptT(dictMonadAsk.Monad0());
    return {
        ask: Control_Monad_Trans_Class.lift(monadTransExceptT)(dictMonadAsk.Monad0())(Control_Monad_Reader_Class.ask(dictMonadAsk)),
        Monad0: function () {
            return monadExceptT1;
        }
    };
};
var monadReaderExceptT = function (dictMonadReader) {
    var monadAskExceptT1 = monadAskExceptT(dictMonadReader.MonadAsk0());
    return {
        local: function (f) {
            return mapExceptT(Control_Monad_Reader_Class.local(dictMonadReader)(f));
        },
        MonadAsk0: function () {
            return monadAskExceptT1;
        }
    };
};
var monadContExceptT = function (dictMonadCont) {
    var monadExceptT1 = monadExceptT(dictMonadCont.Monad0());
    return {
        callCC: function (f) {
            return Control_Monad_Cont_Class.callCC(dictMonadCont)(function (c) {
                var v = f(function (a) {
                    return c(new Data_Either.Right(a));
                });
                return v;
            });
        },
        Monad0: function () {
            return monadExceptT1;
        }
    };
};
var monadEffectExceptT = function (dictMonadEffect) {
    var Monad0 = dictMonadEffect.Monad0();
    var monadExceptT1 = monadExceptT(Monad0);
    return {
        liftEffect: (function () {
            var $174 = lift(Monad0);
            var $175 = Effect_Class.liftEffect(dictMonadEffect);
            return function ($176) {
                return $174($175($176));
            };
        })(),
        Monad0: function () {
            return monadExceptT1;
        }
    };
};
var monadRecExceptT = function (dictMonadRec) {
    var Monad0 = dictMonadRec.Monad0();
    var Bind1 = Monad0.Bind1();
    var Applicative0 = Monad0.Applicative0();
    var monadExceptT1 = monadExceptT(Monad0);
    return {
        tailRecM: function (f) {
            var $177 = Control_Monad_Rec_Class.tailRecM(dictMonadRec)(function (a) {
                var v = f(a);
                return Control_Bind.bind(Bind1)(v)(function (m$prime) {
                    return Control_Applicative.pure(Applicative0)((function () {
                        if (m$prime instanceof Data_Either.Left) {
                            return new Control_Monad_Rec_Class.Done(new Data_Either.Left(m$prime.value0));
                        };
                        if (m$prime instanceof Data_Either.Right && m$prime.value0 instanceof Control_Monad_Rec_Class.Loop) {
                            return new Control_Monad_Rec_Class.Loop(m$prime.value0.value0);
                        };
                        if (m$prime instanceof Data_Either.Right && m$prime.value0 instanceof Control_Monad_Rec_Class.Done) {
                            return new Control_Monad_Rec_Class.Done(new Data_Either.Right(m$prime.value0.value0));
                        };
                        throw new Error("Failed pattern match at Control.Monad.Except.Trans (line 75, column 14 - line 78, column 43): " + [ m$prime.constructor.name ]);
                    })());
                });
            });
            return function ($178) {
                return ExceptT($177($178));
            };
        },
        Monad0: function () {
            return monadExceptT1;
        }
    };
};
var monadStateExceptT = function (dictMonadState) {
    var Monad0 = dictMonadState.Monad0();
    var monadExceptT1 = monadExceptT(dictMonadState.Monad0());
    return {
        state: function (f) {
            return Control_Monad_Trans_Class.lift(monadTransExceptT)(Monad0)(Control_Monad_State_Class.state(dictMonadState)(f));
        },
        Monad0: function () {
            return monadExceptT1;
        }
    };
};
var monadTellExceptT = function (dictMonadTell) {
    var Monad1 = dictMonadTell.Monad1();
    var Semigroup0 = dictMonadTell.Semigroup0();
    var monadExceptT1 = monadExceptT(Monad1);
    return {
        tell: (function () {
            var $179 = lift(Monad1);
            var $180 = Control_Monad_Writer_Class.tell(dictMonadTell);
            return function ($181) {
                return $179($180($181));
            };
        })(),
        Semigroup0: function () {
            return Semigroup0;
        },
        Monad1: function () {
            return monadExceptT1;
        }
    };
};
var monadWriterExceptT = function (dictMonadWriter) {
    var MonadTell1 = dictMonadWriter.MonadTell1();
    var Monad1 = MonadTell1.Monad1();
    var Bind1 = Monad1.Bind1();
    var pure = Control_Applicative.pure(Monad1.Applicative0());
    var Applicative0 = Monad1.Applicative0();
    var Monoid0 = dictMonadWriter.Monoid0();
    var monadTellExceptT1 = monadTellExceptT(MonadTell1);
    return {
        listen: mapExceptT(function (m) {
            return Control_Bind.bind(Bind1)(Control_Monad_Writer_Class.listen(dictMonadWriter)(m))(function (v) {
                return pure(Data_Functor.map(Data_Either.functorEither)(function (r) {
                    return new Data_Tuple.Tuple(r, v.value1);
                })(v.value0));
            });
        }),
        pass: mapExceptT(function (m) {
            return Control_Monad_Writer_Class.pass(dictMonadWriter)(Control_Bind.bind(Bind1)(m)(function (a) {
                return Control_Applicative.pure(Applicative0)((function () {
                    if (a instanceof Data_Either.Left) {
                        return new Data_Tuple.Tuple(new Data_Either.Left(a.value0), Control_Category.identity(Control_Category.categoryFn));
                    };
                    if (a instanceof Data_Either.Right) {
                        return new Data_Tuple.Tuple(new Data_Either.Right(a.value0.value0), a.value0.value1);
                    };
                    throw new Error("Failed pattern match at Control.Monad.Except.Trans (line 135, column 10 - line 137, column 45): " + [ a.constructor.name ]);
                })());
            }));
        }),
        Monoid0: function () {
            return Monoid0;
        },
        MonadTell1: function () {
            return monadTellExceptT1;
        }
    };
};
var monadThrowExceptT = function (dictMonad) {
    var monadExceptT1 = monadExceptT(dictMonad);
    return {
        throwError: (function () {
            var $182 = Control_Applicative.pure(dictMonad.Applicative0());
            return function ($183) {
                return ExceptT($182(Data_Either.Left.create($183)));
            };
        })(),
        Monad0: function () {
            return monadExceptT1;
        }
    };
};
var monadErrorExceptT = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var pure = Control_Applicative.pure(dictMonad.Applicative0());
    var monadThrowExceptT1 = monadThrowExceptT(dictMonad);
    return {
        catchError: function (v) {
            return function (k) {
                return Control_Bind.bind(Bind1)(v)(Data_Either.either(function (a) {
                    var v1 = k(a);
                    return v1;
                })(function ($184) {
                    return pure(Data_Either.Right.create($184));
                }));
            };
        },
        MonadThrow0: function () {
            return monadThrowExceptT1;
        }
    };
};
var monadSTExceptT = function (dictMonadST) {
    var Monad0 = dictMonadST.Monad0();
    var monadExceptT1 = monadExceptT(Monad0);
    return {
        liftST: (function () {
            var $185 = lift(Monad0);
            var $186 = Control_Monad_ST_Class.liftST(dictMonadST);
            return function ($187) {
                return $185($186($187));
            };
        })(),
        Monad0: function () {
            return monadExceptT1;
        }
    };
};
var monoidExceptT = function (dictMonad) {
    var applicativeExceptT1 = applicativeExceptT(dictMonad);
    var semigroupExceptT1 = semigroupExceptT(dictMonad);
    return function (dictMonoid) {
        var semigroupExceptT2 = semigroupExceptT1(dictMonoid.Semigroup0());
        return {
            mempty: Control_Applicative.pure(applicativeExceptT1)(Data_Monoid.mempty(dictMonoid)),
            Semigroup0: function () {
                return semigroupExceptT2;
            }
        };
    };
};
var altExceptT = function (dictSemigroup) {
    return function (dictMonad) {
        var Bind1 = dictMonad.Bind1();
        var Applicative0 = dictMonad.Applicative0();
        var functorExceptT1 = functorExceptT(((dictMonad.Bind1()).Apply0()).Functor0());
        return {
            alt: function (v) {
                return function (v1) {
                    return Control_Bind.bind(Bind1)(v)(function (rm) {
                        if (rm instanceof Data_Either.Right) {
                            return Control_Applicative.pure(Applicative0)(new Data_Either.Right(rm.value0));
                        };
                        if (rm instanceof Data_Either.Left) {
                            return Control_Bind.bind(Bind1)(v1)(function (rn) {
                                if (rn instanceof Data_Either.Right) {
                                    return Control_Applicative.pure(Applicative0)(new Data_Either.Right(rn.value0));
                                };
                                if (rn instanceof Data_Either.Left) {
                                    return Control_Applicative.pure(Applicative0)(new Data_Either.Left(Data_Semigroup.append(dictSemigroup)(rm.value0)(rn.value0)));
                                };
                                throw new Error("Failed pattern match at Control.Monad.Except.Trans (line 87, column 9 - line 89, column 49): " + [ rn.constructor.name ]);
                            });
                        };
                        throw new Error("Failed pattern match at Control.Monad.Except.Trans (line 83, column 5 - line 89, column 49): " + [ rm.constructor.name ]);
                    });
                };
            },
            Functor0: function () {
                return functorExceptT1;
            }
        };
    };
};
var plusExceptT = function (dictMonoid) {
    var mempty = Data_Monoid.mempty(dictMonoid);
    var altExceptT1 = altExceptT(dictMonoid.Semigroup0());
    return function (dictMonad) {
        var altExceptT2 = altExceptT1(dictMonad);
        return {
            empty: Control_Monad_Error_Class.throwError(monadThrowExceptT(dictMonad))(mempty),
            Alt0: function () {
                return altExceptT2;
            }
        };
    };
};
var alternativeExceptT = function (dictMonoid) {
    var plusExceptT1 = plusExceptT(dictMonoid);
    return function (dictMonad) {
        var applicativeExceptT1 = applicativeExceptT(dictMonad);
        var plusExceptT2 = plusExceptT1(dictMonad);
        return {
            Applicative0: function () {
                return applicativeExceptT1;
            },
            Plus1: function () {
                return plusExceptT2;
            }
        };
    };
};
var monadPlusExceptT = function (dictMonoid) {
    var alternativeExceptT1 = alternativeExceptT(dictMonoid);
    return function (dictMonad) {
        var monadExceptT1 = monadExceptT(dictMonad);
        var alternativeExceptT2 = alternativeExceptT1(dictMonad);
        return {
            Monad0: function () {
                return monadExceptT1;
            },
            Alternative1: function () {
                return alternativeExceptT2;
            }
        };
    };
};
export {
    ExceptT,
    runExceptT,
    withExceptT,
    mapExceptT,
    except,
    newtypeExceptT,
    functorExceptT,
    applyExceptT,
    applicativeExceptT,
    bindExceptT,
    monadExceptT,
    monadRecExceptT,
    altExceptT,
    plusExceptT,
    alternativeExceptT,
    monadPlusExceptT,
    monadTransExceptT,
    monadEffectExceptT,
    monadContExceptT,
    monadThrowExceptT,
    monadErrorExceptT,
    monadAskExceptT,
    monadReaderExceptT,
    monadStateExceptT,
    monadTellExceptT,
    monadWriterExceptT,
    semigroupExceptT,
    monoidExceptT,
    monadSTExceptT
};
export {
    catchError,
    throwError
} from "../Control.Monad.Error.Class/index.js";
export {
    lift
} from "../Control.Monad.Trans.Class/index.js";
