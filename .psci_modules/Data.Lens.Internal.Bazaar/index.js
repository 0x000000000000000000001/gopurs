import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Data_Bitraversable from "../Data.Bitraversable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var Bazaar = function (x) {
    return x;
};
var runBazaar = function (v) {
    return function (dictApplicative) {
        return v(dictApplicative);
    };
};
var profunctorBazaar = {
    dimap: function (f) {
        return function (g) {
            return function (v) {
                return function (dictApplicative) {
                    var Functor0 = (dictApplicative.Apply0()).Functor0();
                    return function (pafb) {
                        return function (s) {
                            return Data_Functor.map(Functor0)(g)(v(dictApplicative)(pafb)(f(s)));
                        };
                    };
                };
            };
        };
    }
};
var strongBazaar = {
    first: function (v) {
        return function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (pafb) {
                return function (v1) {
                    return Data_Functor.map(Functor0)(Data_Function.flip(Data_Tuple.Tuple.create)(v1.value1))(v(dictApplicative)(pafb)(v1.value0));
                };
            };
        };
    },
    second: function (v) {
        return function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (pafb) {
                return function (v1) {
                    return Data_Functor.map(Functor0)(Data_Tuple.Tuple.create(v1.value0))(v(dictApplicative)(pafb)(v1.value1));
                };
            };
        };
    },
    Profunctor0: function () {
        return profunctorBazaar;
    }
};
var choiceBazaar = {
    left: function (v) {
        return function (dictApplicative) {
            var pure = Control_Applicative.pure(dictApplicative);
            return function (pafb) {
                return function (e) {
                    return Data_Bitraversable.bitraverse(Data_Bitraversable.bitraversableEither)(dictApplicative)(v(dictApplicative)(pafb))(pure)(e);
                };
            };
        };
    },
    right: function (v) {
        return function (dictApplicative) {
            return function (pafb) {
                return function (e) {
                    return Data_Traversable.traverse(Data_Traversable.traversableEither)(dictApplicative)(v(dictApplicative)(pafb))(e);
                };
            };
        };
    },
    Profunctor0: function () {
        return profunctorBazaar;
    }
};
var wanderBazaar = {
    wander: function (w) {
        return function (v) {
            return function (dictApplicative) {
                return function (pafb) {
                    return function (s) {
                        return w(dictApplicative)(v(dictApplicative)(pafb))(s);
                    };
                };
            };
        };
    },
    Strong0: function () {
        return strongBazaar;
    },
    Choice1: function () {
        return choiceBazaar;
    }
};
export {
    Bazaar,
    runBazaar,
    profunctorBazaar,
    strongBazaar,
    choiceBazaar,
    wanderBazaar
};
