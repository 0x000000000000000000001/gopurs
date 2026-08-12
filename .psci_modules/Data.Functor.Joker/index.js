import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Show from "../Data.Show/index.js";
var Joker = function (x) {
    return x;
};
var showJoker = function (dictShow) {
    return {
        show: function (v) {
            return "(Joker " + (Data_Show.show(dictShow)(v) + ")");
        }
    };
};
var profunctorJoker = function (dictFunctor) {
    return {
        dimap: function (v) {
            return function (g) {
                return function (v1) {
                    return Data_Functor.map(dictFunctor)(g)(v1);
                };
            };
        }
    };
};
var ordJoker = function (dictOrd) {
    return dictOrd;
};
var newtypeJoker = {
    Coercible0: function () {
        return undefined;
    }
};
var hoistJoker = function (f) {
    return function (v) {
        return f(v);
    };
};
var functorJoker = function (dictFunctor) {
    return {
        map: function (f) {
            return function (v) {
                return Data_Functor.map(dictFunctor)(f)(v);
            };
        }
    };
};
var eqJoker = function (dictEq) {
    return dictEq;
};
var choiceJoker = function (dictFunctor) {
    var profunctorJoker1 = profunctorJoker(dictFunctor);
    return {
        left: function (v) {
            return Data_Functor.map(dictFunctor)(Data_Either.Left.create)(v);
        },
        right: function (v) {
            return Data_Functor.map(dictFunctor)(Data_Either.Right.create)(v);
        },
        Profunctor0: function () {
            return profunctorJoker1;
        }
    };
};
var bifunctorJoker = function (dictFunctor) {
    return {
        bimap: function (v) {
            return function (g) {
                return function (v1) {
                    return Data_Functor.map(dictFunctor)(g)(v1);
                };
            };
        }
    };
};
var biapplyJoker = function (dictApply) {
    var bifunctorJoker1 = bifunctorJoker(dictApply.Functor0());
    return {
        biapply: function (v) {
            return function (v1) {
                return Control_Apply.apply(dictApply)(v)(v1);
            };
        },
        Bifunctor0: function () {
            return bifunctorJoker1;
        }
    };
};
var biapplicativeJoker = function (dictApplicative) {
    var biapplyJoker1 = biapplyJoker(dictApplicative.Apply0());
    return {
        bipure: function (v) {
            return function (b) {
                return Control_Applicative.pure(dictApplicative)(b);
            };
        },
        Biapply0: function () {
            return biapplyJoker1;
        }
    };
};
var applyJoker = function (dictApply) {
    var functorJoker1 = functorJoker(dictApply.Functor0());
    return {
        apply: function (v) {
            return function (v1) {
                return Control_Apply.apply(dictApply)(v)(v1);
            };
        },
        Functor0: function () {
            return functorJoker1;
        }
    };
};
var bindJoker = function (dictBind) {
    var applyJoker1 = applyJoker(dictBind.Apply0());
    return {
        bind: function (v) {
            return function (amb) {
                return Control_Bind.bind(dictBind)(v)((function () {
                    var $63 = Data_Newtype.un()(Joker);
                    return function ($64) {
                        return $63(amb($64));
                    };
                })());
            };
        },
        Apply0: function () {
            return applyJoker1;
        }
    };
};
var applicativeJoker = function (dictApplicative) {
    var applyJoker1 = applyJoker(dictApplicative.Apply0());
    return {
        pure: (function () {
            var $65 = Control_Applicative.pure(dictApplicative);
            return function ($66) {
                return Joker($65($66));
            };
        })(),
        Apply0: function () {
            return applyJoker1;
        }
    };
};
var monadJoker = function (dictMonad) {
    var applicativeJoker1 = applicativeJoker(dictMonad.Applicative0());
    var bindJoker1 = bindJoker(dictMonad.Bind1());
    return {
        Applicative0: function () {
            return applicativeJoker1;
        },
        Bind1: function () {
            return bindJoker1;
        }
    };
};
export {
    Joker,
    hoistJoker,
    newtypeJoker,
    eqJoker,
    ordJoker,
    showJoker,
    functorJoker,
    applyJoker,
    applicativeJoker,
    bindJoker,
    monadJoker,
    bifunctorJoker,
    biapplyJoker,
    biapplicativeJoker,
    profunctorJoker,
    choiceJoker
};
