import * as Control_Comonad from "../Control.Comonad/index.js";
import * as Control_Extend from "../Control.Extend/index.js";
import * as Data_Distributive from "../Data.Distributive/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Functor_Contravariant from "../Data.Functor.Contravariant/index.js";
import * as Data_Functor_Invariant from "../Data.Functor.Invariant/index.js";
import * as Data_Profunctor from "../Data.Profunctor/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var Costar = function (x) {
    return x;
};
var semigroupoidCostar = function (dictExtend) {
    return {
        compose: function (v) {
            return function (v1) {
                return Control_Extend.composeCoKleisliFlipped(dictExtend)(v)(v1);
            };
        }
    };
};
var profunctorCostar = function (dictFunctor) {
    return {
        dimap: function (f) {
            return function (g) {
                return function (v) {
                    var $55 = Data_Functor.map(dictFunctor)(f);
                    return function ($56) {
                        return g(v($55($56)));
                    };
                };
            };
        }
    };
};
var strongCostar = function (dictComonad) {
    var Extend0 = dictComonad.Extend0();
    var Functor0 = Extend0.Functor0();
    var profunctorCostar1 = profunctorCostar(Extend0.Functor0());
    return {
        first: function (v) {
            return function (x) {
                return new Data_Tuple.Tuple(v(Data_Functor.map(Functor0)(Data_Tuple.fst)(x)), Data_Tuple.snd(Control_Comonad.extract(dictComonad)(x)));
            };
        },
        second: function (v) {
            return function (x) {
                return new Data_Tuple.Tuple(Data_Tuple.fst(Control_Comonad.extract(dictComonad)(x)), v(Data_Functor.map(Functor0)(Data_Tuple.snd)(x)));
            };
        },
        Profunctor0: function () {
            return profunctorCostar1;
        }
    };
};
var newtypeCostar = {
    Coercible0: function () {
        return undefined;
    }
};
var hoistCostar = function (f) {
    return function (v) {
        return Data_Profunctor.lcmap(Data_Profunctor.profunctorFn)(f)(v);
    };
};
var functorCostar = {
    map: function (f) {
        return function (v) {
            return function ($57) {
                return f(v($57));
            };
        };
    }
};
var invariantCostar = {
    imap: /* #__PURE__ */ Data_Functor_Invariant.imapF(functorCostar)
};
var distributiveCostar = {
    distribute: function (dictFunctor) {
        return function (f) {
            return function (a) {
                return Data_Functor.map(dictFunctor)(function (v) {
                    return v(a);
                })(f);
            };
        };
    },
    collect: function (dictFunctor) {
        return function (f) {
            var $58 = Data_Distributive.distribute(distributiveCostar)(dictFunctor);
            var $59 = Data_Functor.map(dictFunctor)(f);
            return function ($60) {
                return $58($59($60));
            };
        };
    },
    Functor0: function () {
        return functorCostar;
    }
};
var closedCostar = function (dictFunctor) {
    var profunctorCostar1 = profunctorCostar(dictFunctor);
    return {
        closed: function (v) {
            return function (g) {
                return function (x) {
                    return v(Data_Functor.map(dictFunctor)(function (v1) {
                        return v1(x);
                    })(g));
                };
            };
        },
        Profunctor0: function () {
            return profunctorCostar1;
        }
    };
};
var categoryCostar = function (dictComonad) {
    var semigroupoidCostar1 = semigroupoidCostar(dictComonad.Extend0());
    return {
        identity: Control_Comonad.extract(dictComonad),
        Semigroupoid0: function () {
            return semigroupoidCostar1;
        }
    };
};
var bifunctorCostar = function (dictContravariant) {
    return {
        bimap: function (f) {
            return function (g) {
                return function (v) {
                    var $61 = Data_Functor_Contravariant.cmap(dictContravariant)(f);
                    return function ($62) {
                        return g(v($61($62)));
                    };
                };
            };
        }
    };
};
var applyCostar = {
    apply: function (v) {
        return function (v1) {
            return function (a) {
                return v(a)(v1(a));
            };
        };
    },
    Functor0: function () {
        return functorCostar;
    }
};
var bindCostar = {
    bind: function (v) {
        return function (f) {
            return function (x) {
                var v1 = f(v(x));
                return v1(x);
            };
        };
    },
    Apply0: function () {
        return applyCostar;
    }
};
var applicativeCostar = {
    pure: function (a) {
        return function (v) {
            return a;
        };
    },
    Apply0: function () {
        return applyCostar;
    }
};
var monadCostar = {
    Applicative0: function () {
        return applicativeCostar;
    },
    Bind1: function () {
        return bindCostar;
    }
};
export {
    Costar,
    hoistCostar,
    newtypeCostar,
    semigroupoidCostar,
    categoryCostar,
    functorCostar,
    invariantCostar,
    applyCostar,
    applicativeCostar,
    bindCostar,
    monadCostar,
    distributiveCostar,
    bifunctorCostar,
    profunctorCostar,
    strongCostar,
    closedCostar
};
