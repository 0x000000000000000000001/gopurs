import * as Control_Alt from "../Control.Alt/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Plus from "../Control.Plus/index.js";
import * as Data_Distributive from "../Data.Distributive/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Functor_Invariant from "../Data.Functor.Invariant/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var Star = function (x) {
    return x;
};
var semigroupoidStar = function (dictBind) {
    return {
        compose: function (v) {
            return function (v1) {
                return function (x) {
                    return Control_Bind.bind(dictBind)(v1(x))(v);
                };
            };
        }
    };
};
var profunctorStar = function (dictFunctor) {
    return {
        dimap: function (f) {
            return function (g) {
                return function (v) {
                    var $112 = Data_Functor.map(dictFunctor)(g);
                    return function ($113) {
                        return $112(v(f($113)));
                    };
                };
            };
        }
    };
};
var strongStar = function (dictFunctor) {
    var profunctorStar1 = profunctorStar(dictFunctor);
    return {
        first: function (v) {
            return function (v1) {
                return Data_Functor.map(dictFunctor)(function (v2) {
                    return new Data_Tuple.Tuple(v2, v1.value1);
                })(v(v1.value0));
            };
        },
        second: function (v) {
            return function (v1) {
                return Data_Functor.map(dictFunctor)(Data_Tuple.Tuple.create(v1.value0))(v(v1.value1));
            };
        },
        Profunctor0: function () {
            return profunctorStar1;
        }
    };
};
var newtypeStar = {
    Coercible0: function () {
        return undefined;
    }
};
var invariantStar = function (dictInvariant) {
    return {
        imap: function (f) {
            return function (g) {
                return function (v) {
                    var $114 = Data_Functor_Invariant.imap(dictInvariant)(f)(g);
                    return function ($115) {
                        return $114(v($115));
                    };
                };
            };
        }
    };
};
var hoistStar = function (f) {
    return function (v) {
        return function ($116) {
            return f(v($116));
        };
    };
};
var functorStar = function (dictFunctor) {
    return {
        map: function (f) {
            return function (v) {
                var $117 = Data_Functor.map(dictFunctor)(f);
                return function ($118) {
                    return $117(v($118));
                };
            };
        }
    };
};
var distributiveStar = function (dictDistributive) {
    var functorStar1 = functorStar(dictDistributive.Functor0());
    return {
        distribute: function (dictFunctor) {
            return function (f) {
                return function (a) {
                    return Data_Distributive.collect(dictDistributive)(dictFunctor)(function (v) {
                        return v(a);
                    })(f);
                };
            };
        },
        collect: function (dictFunctor) {
            return function (f) {
                var $119 = Data_Distributive.distribute(distributiveStar(dictDistributive))(dictFunctor);
                var $120 = Data_Functor.map(dictFunctor)(f);
                return function ($121) {
                    return $119($120($121));
                };
            };
        },
        Functor0: function () {
            return functorStar1;
        }
    };
};
var closedStar = function (dictDistributive) {
    var profunctorStar1 = profunctorStar(dictDistributive.Functor0());
    return {
        closed: function (v) {
            return function (g) {
                return Data_Distributive.distribute(dictDistributive)(Data_Functor.functorFn)(function ($122) {
                    return v(g($122));
                });
            };
        },
        Profunctor0: function () {
            return profunctorStar1;
        }
    };
};
var choiceStar = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = Apply0.Functor0();
    var pure = Control_Applicative.pure(dictApplicative);
    var pure1 = Control_Applicative.pure(dictApplicative);
    var profunctorStar1 = profunctorStar(Apply0.Functor0());
    return {
        left: function (v) {
            return Data_Either.either((function () {
                var $123 = Data_Functor.map(Functor0)(Data_Either.Left.create);
                return function ($124) {
                    return $123(v($124));
                };
            })())(function ($125) {
                return pure(Data_Either.Right.create($125));
            });
        },
        right: function (v) {
            return Data_Either.either(function ($126) {
                return pure1(Data_Either.Left.create($126));
            })((function () {
                var $127 = Data_Functor.map(Functor0)(Data_Either.Right.create);
                return function ($128) {
                    return $127(v($128));
                };
            })());
        },
        Profunctor0: function () {
            return profunctorStar1;
        }
    };
};
var categoryStar = function (dictMonad) {
    var semigroupoidStar1 = semigroupoidStar(dictMonad.Bind1());
    return {
        identity: Control_Applicative.pure(dictMonad.Applicative0()),
        Semigroupoid0: function () {
            return semigroupoidStar1;
        }
    };
};
var applyStar = function (dictApply) {
    var functorStar1 = functorStar(dictApply.Functor0());
    return {
        apply: function (v) {
            return function (v1) {
                return function (a) {
                    return Control_Apply.apply(dictApply)(v(a))(v1(a));
                };
            };
        },
        Functor0: function () {
            return functorStar1;
        }
    };
};
var bindStar = function (dictBind) {
    var applyStar1 = applyStar(dictBind.Apply0());
    return {
        bind: function (v) {
            return function (f) {
                return function (x) {
                    return Control_Bind.bind(dictBind)(v(x))(function (a) {
                        var v1 = f(a);
                        return v1(x);
                    });
                };
            };
        },
        Apply0: function () {
            return applyStar1;
        }
    };
};
var applicativeStar = function (dictApplicative) {
    var applyStar1 = applyStar(dictApplicative.Apply0());
    return {
        pure: function (a) {
            return function (v) {
                return Control_Applicative.pure(dictApplicative)(a);
            };
        },
        Apply0: function () {
            return applyStar1;
        }
    };
};
var monadStar = function (dictMonad) {
    var applicativeStar1 = applicativeStar(dictMonad.Applicative0());
    var bindStar1 = bindStar(dictMonad.Bind1());
    return {
        Applicative0: function () {
            return applicativeStar1;
        },
        Bind1: function () {
            return bindStar1;
        }
    };
};
var altStar = function (dictAlt) {
    var functorStar1 = functorStar(dictAlt.Functor0());
    return {
        alt: function (v) {
            return function (v1) {
                return function (a) {
                    return Control_Alt.alt(dictAlt)(v(a))(v1(a));
                };
            };
        },
        Functor0: function () {
            return functorStar1;
        }
    };
};
var plusStar = function (dictPlus) {
    var empty = Control_Plus.empty(dictPlus);
    var altStar1 = altStar(dictPlus.Alt0());
    return {
        empty: function (v) {
            return empty;
        },
        Alt0: function () {
            return altStar1;
        }
    };
};
var alternativeStar = function (dictAlternative) {
    var applicativeStar1 = applicativeStar(dictAlternative.Applicative0());
    var plusStar1 = plusStar(dictAlternative.Plus1());
    return {
        Applicative0: function () {
            return applicativeStar1;
        },
        Plus1: function () {
            return plusStar1;
        }
    };
};
var monadPlusStar = function (dictMonadPlus) {
    var monadStar1 = monadStar(dictMonadPlus.Monad0());
    var alternativeStar1 = alternativeStar(dictMonadPlus.Alternative1());
    return {
        Monad0: function () {
            return monadStar1;
        },
        Alternative1: function () {
            return alternativeStar1;
        }
    };
};
export {
    Star,
    hoistStar,
    newtypeStar,
    semigroupoidStar,
    categoryStar,
    functorStar,
    invariantStar,
    applyStar,
    applicativeStar,
    bindStar,
    monadStar,
    altStar,
    plusStar,
    alternativeStar,
    monadPlusStar,
    distributiveStar,
    profunctorStar,
    strongStar,
    choiceStar,
    closedStar
};
