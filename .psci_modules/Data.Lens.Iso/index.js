// | This module defines functions for working with isomorphisms.
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Lens_Internal_Exchange from "../Data.Lens.Internal.Exchange/index.js";
import * as Data_Lens_Types from "../Data.Lens.Types/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Profunctor from "../Data.Profunctor/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Safe_Coerce from "../Safe.Coerce/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var identity1 = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var identity2 = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var coerce = /* #__PURE__ */ Safe_Coerce.coerce();
var coerce1 = /* #__PURE__ */ Safe_Coerce.coerce();

// | Extracts the pair of morphisms from an isomorphism.
var withIso = function (l) {
    return function (f) {
        var v = l(new Data_Lens_Internal_Exchange.Exchange(identity, identity1));
        return f(v.value0)(v.value1);
    };
};
var under = function (l) {
    return withIso(l)(function (sa) {
        return function (bt) {
            return function (ts) {
                return function ($25) {
                    return sa(ts(bt($25)));
                };
            };
        };
    });
};

// | Reverses an optic.
var re = function (t) {
    return Data_Newtype.unwrap()(t(identity2));
};

// | Create an `Iso` from a pair of morphisms.
var iso = function (f) {
    return function (g) {
        return function (dictProfunctor) {
            return function (pab) {
                return Data_Profunctor.dimap(dictProfunctor)(f)(g)(pab);
            };
        };
    };
};
var mapping = function (dictFunctor) {
    return function (dictFunctor1) {
        return function (l) {
            return function (dictProfunctor) {
                return withIso(l)(function (sa) {
                    return function (bt) {
                        return iso(Data_Functor.map(dictFunctor)(sa))(Data_Functor.map(dictFunctor1)(bt))(dictProfunctor);
                    };
                });
            };
        };
    };
};

// | If `a1` is obtained from `a` by removing a single value, then
// | `Maybe a1` is isomorphic to `a`.
var non = function (dictEq) {
    return function (def) {
        return function (dictProfunctor) {
            var g = function (a) {
                if (Data_Eq.eq(dictEq)(a)(def)) {
                    return Data_Maybe.Nothing.value;
                };
                if (Data_Boolean.otherwise) {
                    return new Data_Maybe.Just(a);
                };
                throw new Error("Failed pattern match at Data.Lens.Iso (line 60, column 3 - line 62, column 25): " + [ a.constructor.name ]);
            };
            return iso(Data_Maybe.fromMaybe(def))(g)(dictProfunctor);
        };
    };
};
var uncurried = function (dictProfunctor) {
    return iso(Data_Tuple.uncurry)(Data_Tuple.curry)(dictProfunctor);
};
var flipped = function (dictProfunctor) {
    return iso(Data_Function.flip)(Data_Function.flip)(dictProfunctor);
};
var dimapping = function (dictProfunctor) {
    return function (dictProfunctor1) {
        return function (f) {
            return function (g) {
                return function (dictProfunctor2) {
                    return withIso(f)(function (sa) {
                        return function (bt) {
                            return withIso(g)(function (ssaa) {
                                return function (bbtt) {
                                    return iso(Data_Profunctor.dimap(dictProfunctor)(sa)(ssaa))(Data_Profunctor.dimap(dictProfunctor1)(bt)(bbtt))(dictProfunctor2);
                                };
                            });
                        };
                    });
                };
            };
        };
    };
};
var curried = function (dictProfunctor) {
    return iso(Data_Tuple.curry)(Data_Tuple.uncurry)(dictProfunctor);
};

// | An `Iso` between two types that are inferred to have the
// | same representation by the compiler. One scenario that this is
// | particularly useful is the case of nested newtypes:
// |
// |```purescript
// |  newtype UserId = UserId Int
// |  newtype DeletedUserId = DeletedUserId UserId
// |
// |  -- `simple` is used to aid the type inference
// |  deletedUser :: DeletedUserId
// |  deletedUser = review (simple coerced) 42
// |```
var coerced = function () {
    return function () {
        return function (dictProfunctor) {
            return iso(coerce)(coerce1)(dictProfunctor);
        };
    };
};

// | Extracts an `Iso` from `AnIso`.
var cloneIso = function (l) {
    return function (dictProfunctor) {
        return withIso(l)(function (x) {
            return function (y) {
                return function (p) {
                    return iso(x)(y)(dictProfunctor)(p);
                };
            };
        });
    };
};
var auf = function (dictProfunctor) {
    return function (l) {
        return withIso(l)(function (sa) {
            return function (bt) {
                return function (f) {
                    return function (g) {
                        return function (e) {
                            return bt(f(Data_Profunctor.rmap(dictProfunctor)(sa)(g))(e));
                        };
                    };
                };
            };
        });
    };
};
var au = function (l) {
    return withIso(l)(function (sa) {
        return function (bt) {
            return function (f) {
                return function (e) {
                    return sa(f(bt)(e));
                };
            };
        };
    });
};
export {
    iso,
    withIso,
    cloneIso,
    re,
    au,
    auf,
    under,
    non,
    curried,
    uncurried,
    flipped,
    mapping,
    dimapping,
    coerced
};
export {
    Exchange,
    Re
} from "../Data.Lens.Types/index.js";
