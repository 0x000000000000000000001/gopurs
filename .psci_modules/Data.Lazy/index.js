import * as $foreign from "./foreign.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Data_Bounded from "../Data.Bounded/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_EuclideanRing from "../Data.EuclideanRing/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Functor_Invariant from "../Data.Functor.Invariant/index.js";
import * as Data_HeytingAlgebra from "../Data.HeytingAlgebra/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ring from "../Data.Ring/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
var showLazy = function (dictShow) {
    return {
        show: function (x) {
            return "(defer \\_ -> " + (Data_Show.show(dictShow)($foreign.force(x)) + ")");
        }
    };
};
var semiringLazy = function (dictSemiring) {
    return {
        add: function (a) {
            return function (b) {
                return $foreign.defer(function (v) {
                    return Data_Semiring.add(dictSemiring)($foreign.force(a))($foreign.force(b));
                });
            };
        },
        zero: $foreign.defer(function (v) {
            return Data_Semiring.zero(dictSemiring);
        }),
        mul: function (a) {
            return function (b) {
                return $foreign.defer(function (v) {
                    return Data_Semiring.mul(dictSemiring)($foreign.force(a))($foreign.force(b));
                });
            };
        },
        one: $foreign.defer(function (v) {
            return Data_Semiring.one(dictSemiring);
        })
    };
};
var semigroupLazy = function (dictSemigroup) {
    return {
        append: function (a) {
            return function (b) {
                return $foreign.defer(function (v) {
                    return Data_Semigroup.append(dictSemigroup)($foreign.force(a))($foreign.force(b));
                });
            };
        }
    };
};
var ringLazy = function (dictRing) {
    var semiringLazy1 = semiringLazy(dictRing.Semiring0());
    return {
        sub: function (a) {
            return function (b) {
                return $foreign.defer(function (v) {
                    return Data_Ring.sub(dictRing)($foreign.force(a))($foreign.force(b));
                });
            };
        },
        Semiring0: function () {
            return semiringLazy1;
        }
    };
};
var monoidLazy = function (dictMonoid) {
    var semigroupLazy1 = semigroupLazy(dictMonoid.Semigroup0());
    return {
        mempty: $foreign.defer(function (v) {
            return Data_Monoid.mempty(dictMonoid);
        }),
        Semigroup0: function () {
            return semigroupLazy1;
        }
    };
};
var lazyLazy = {
    defer: function (f) {
        return $foreign.defer(function (v) {
            return $foreign.force(f(Data_Unit.unit));
        });
    }
};
var functorLazy = {
    map: function (f) {
        return function (l) {
            return $foreign.defer(function (v) {
                return f($foreign.force(l));
            });
        };
    }
};
var map = /* #__PURE__ */ Data_Functor.map(functorLazy);
var functorWithIndexLazy = {
    mapWithIndex: function (f) {
        return map(f(Data_Unit.unit));
    },
    Functor0: function () {
        return functorLazy;
    }
};
var invariantLazy = {
    imap: /* #__PURE__ */ Data_Functor_Invariant.imapF(functorLazy)
};
var foldableLazy = {
    foldr: function (f) {
        return function (z) {
            return function (l) {
                return f($foreign.force(l))(z);
            };
        };
    },
    foldl: function (f) {
        return function (z) {
            return function (l) {
                return f(z)($foreign.force(l));
            };
        };
    },
    foldMap: function (dictMonoid) {
        return function (f) {
            return function (l) {
                return f($foreign.force(l));
            };
        };
    }
};
var foldr = /* #__PURE__ */ Data_Foldable.foldr(foldableLazy);
var foldl = /* #__PURE__ */ Data_Foldable.foldl(foldableLazy);
var foldMap = /* #__PURE__ */ Data_Foldable.foldMap(foldableLazy);
var foldableWithIndexLazy = {
    foldrWithIndex: function (f) {
        return foldr(f(Data_Unit.unit));
    },
    foldlWithIndex: function (f) {
        return foldl(f(Data_Unit.unit));
    },
    foldMapWithIndex: function (dictMonoid) {
        var foldMap1 = foldMap(dictMonoid);
        return function (f) {
            return foldMap1(f(Data_Unit.unit));
        };
    },
    Foldable0: function () {
        return foldableLazy;
    }
};
var traversableLazy = {
    traverse: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (f) {
            return function (l) {
                return Data_Functor.map(Functor0)(function ($80) {
                    return $foreign.defer(Data_Function["const"]($80));
                })(f($foreign.force(l)));
            };
        };
    },
    sequence: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (l) {
            return Data_Functor.map(Functor0)(function ($81) {
                return $foreign.defer(Data_Function["const"]($81));
            })($foreign.force(l));
        };
    },
    Functor0: function () {
        return functorLazy;
    },
    Foldable1: function () {
        return foldableLazy;
    }
};
var traverse = /* #__PURE__ */ Data_Traversable.traverse(traversableLazy);
var traversableWithIndexLazy = {
    traverseWithIndex: function (dictApplicative) {
        var traverse1 = traverse(dictApplicative);
        return function (f) {
            return traverse1(f(Data_Unit.unit));
        };
    },
    FunctorWithIndex0: function () {
        return functorWithIndexLazy;
    },
    FoldableWithIndex1: function () {
        return foldableWithIndexLazy;
    },
    Traversable2: function () {
        return traversableLazy;
    }
};
var foldable1Lazy = {
    foldMap1: function (dictSemigroup) {
        return function (f) {
            return function (l) {
                return f($foreign.force(l));
            };
        };
    },
    foldr1: function (v) {
        return function (l) {
            return $foreign.force(l);
        };
    },
    foldl1: function (v) {
        return function (l) {
            return $foreign.force(l);
        };
    },
    Foldable0: function () {
        return foldableLazy;
    }
};
var traversable1Lazy = {
    traverse1: function (dictApply) {
        var Functor0 = dictApply.Functor0();
        return function (f) {
            return function (l) {
                return Data_Functor.map(Functor0)(function ($82) {
                    return $foreign.defer(Data_Function["const"]($82));
                })(f($foreign.force(l)));
            };
        };
    },
    sequence1: function (dictApply) {
        var Functor0 = dictApply.Functor0();
        return function (l) {
            return Data_Functor.map(Functor0)(function ($83) {
                return $foreign.defer(Data_Function["const"]($83));
            })($foreign.force(l));
        };
    },
    Foldable10: function () {
        return foldable1Lazy;
    },
    Traversable1: function () {
        return traversableLazy;
    }
};
var extendLazy = {
    extend: function (f) {
        return function (x) {
            return $foreign.defer(function (v) {
                return f(x);
            });
        };
    },
    Functor0: function () {
        return functorLazy;
    }
};
var eqLazy = function (dictEq) {
    return {
        eq: function (x) {
            return function (y) {
                return Data_Eq.eq(dictEq)($foreign.force(x))($foreign.force(y));
            };
        }
    };
};
var ordLazy = function (dictOrd) {
    var eqLazy1 = eqLazy(dictOrd.Eq0());
    return {
        compare: function (x) {
            return function (y) {
                return Data_Ord.compare(dictOrd)($foreign.force(x))($foreign.force(y));
            };
        },
        Eq0: function () {
            return eqLazy1;
        }
    };
};
var eq1Lazy = {
    eq1: function (dictEq) {
        return Data_Eq.eq(eqLazy(dictEq));
    }
};
var ord1Lazy = {
    compare1: function (dictOrd) {
        return Data_Ord.compare(ordLazy(dictOrd));
    },
    Eq10: function () {
        return eq1Lazy;
    }
};
var comonadLazy = {
    extract: $foreign.force,
    Extend0: function () {
        return extendLazy;
    }
};
var commutativeRingLazy = function (dictCommutativeRing) {
    var ringLazy1 = ringLazy(dictCommutativeRing.Ring0());
    return {
        Ring0: function () {
            return ringLazy1;
        }
    };
};
var euclideanRingLazy = function (dictEuclideanRing) {
    var commutativeRingLazy1 = commutativeRingLazy(dictEuclideanRing.CommutativeRing0());
    return {
        degree: (function () {
            var $84 = Data_EuclideanRing.degree(dictEuclideanRing);
            return function ($85) {
                return $84($foreign.force($85));
            };
        })(),
        div: function (a) {
            return function (b) {
                return $foreign.defer(function (v) {
                    return Data_EuclideanRing.div(dictEuclideanRing)($foreign.force(a))($foreign.force(b));
                });
            };
        },
        mod: function (a) {
            return function (b) {
                return $foreign.defer(function (v) {
                    return Data_EuclideanRing.mod(dictEuclideanRing)($foreign.force(a))($foreign.force(b));
                });
            };
        },
        CommutativeRing0: function () {
            return commutativeRingLazy1;
        }
    };
};
var boundedLazy = function (dictBounded) {
    var ordLazy1 = ordLazy(dictBounded.Ord0());
    return {
        top: $foreign.defer(function (v) {
            return Data_Bounded.top(dictBounded);
        }),
        bottom: $foreign.defer(function (v) {
            return Data_Bounded.bottom(dictBounded);
        }),
        Ord0: function () {
            return ordLazy1;
        }
    };
};
var applyLazy = {
    apply: function (f) {
        return function (x) {
            return $foreign.defer(function (v) {
                return $foreign.force(f)($foreign.force(x));
            });
        };
    },
    Functor0: function () {
        return functorLazy;
    }
};
var bindLazy = {
    bind: function (l) {
        return function (f) {
            return $foreign.defer(function (v) {
                return $foreign.force(f($foreign.force(l)));
            });
        };
    },
    Apply0: function () {
        return applyLazy;
    }
};
var heytingAlgebraLazy = function (dictHeytingAlgebra) {
    var implies = Data_HeytingAlgebra.implies(dictHeytingAlgebra);
    var conj = Data_HeytingAlgebra.conj(dictHeytingAlgebra);
    var disj = Data_HeytingAlgebra.disj(dictHeytingAlgebra);
    var not = Data_HeytingAlgebra.not(dictHeytingAlgebra);
    return {
        ff: $foreign.defer(function (v) {
            return Data_HeytingAlgebra.ff(dictHeytingAlgebra);
        }),
        tt: $foreign.defer(function (v) {
            return Data_HeytingAlgebra.tt(dictHeytingAlgebra);
        }),
        implies: function (a) {
            return function (b) {
                return Control_Apply.apply(applyLazy)(Data_Functor.map(functorLazy)(implies)(a))(b);
            };
        },
        conj: function (a) {
            return function (b) {
                return Control_Apply.apply(applyLazy)(Data_Functor.map(functorLazy)(conj)(a))(b);
            };
        },
        disj: function (a) {
            return function (b) {
                return Control_Apply.apply(applyLazy)(Data_Functor.map(functorLazy)(disj)(a))(b);
            };
        },
        not: function (a) {
            return Data_Functor.map(functorLazy)(not)(a);
        }
    };
};
var booleanAlgebraLazy = function (dictBooleanAlgebra) {
    var heytingAlgebraLazy1 = heytingAlgebraLazy(dictBooleanAlgebra.HeytingAlgebra0());
    return {
        HeytingAlgebra0: function () {
            return heytingAlgebraLazy1;
        }
    };
};
var applicativeLazy = {
    pure: function (a) {
        return $foreign.defer(function (v) {
            return a;
        });
    },
    Apply0: function () {
        return applyLazy;
    }
};
var monadLazy = {
    Applicative0: function () {
        return applicativeLazy;
    },
    Bind1: function () {
        return bindLazy;
    }
};
export {
    defer,
    force
} from "./foreign.js";
export {
    semiringLazy,
    ringLazy,
    commutativeRingLazy,
    euclideanRingLazy,
    eqLazy,
    eq1Lazy,
    ordLazy,
    ord1Lazy,
    boundedLazy,
    semigroupLazy,
    monoidLazy,
    heytingAlgebraLazy,
    booleanAlgebraLazy,
    functorLazy,
    functorWithIndexLazy,
    foldableLazy,
    foldableWithIndexLazy,
    foldable1Lazy,
    traversableLazy,
    traversableWithIndexLazy,
    traversable1Lazy,
    invariantLazy,
    applyLazy,
    applicativeLazy,
    bindLazy,
    monadLazy,
    extendLazy,
    comonadLazy,
    showLazy,
    lazyLazy
};
