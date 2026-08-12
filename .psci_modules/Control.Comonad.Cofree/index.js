// | The _cofree comonad_ for a `Functor`.
import * as Control_Alt from "../Control.Alt/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Comonad from "../Control.Comonad/index.js";
import * as Control_Monad from "../Control.Monad/index.js";
import * as Control_Monad_Free from "../Control.Monad.Free/index.js";
import * as Control_Monad_Rec_Class from "../Control.Monad.Rec.Class/index.js";
import * as Control_Monad_State from "../Control.Monad.State/index.js";
import * as Control_Monad_State_Class from "../Control.Monad.State.Class/index.js";
import * as Control_Monad_State_Trans from "../Control.Monad.State.Trans/index.js";
import * as Control_Plus from "../Control.Plus/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Lazy from "../Data.Lazy/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var monadStateStateT = /* #__PURE__ */ Control_Monad_State_Trans.monadStateStateT(Data_Identity.monadIdentity);
var monadRecStateT = /* #__PURE__ */ Control_Monad_State_Trans.monadRecStateT(Control_Monad_Rec_Class.monadRecIdentity);

// | The `Cofree` `Comonad` for a functor.
// |
// | A value of type `Cofree f a` consists of an `f`-branching
// | tree, annotated with labels of type `a`.
// |
// | The `Comonad` instance supports _redecoration_, recomputing
// | labels from the local context.
var Cofree = function (x) {
    return x;
};

// | Returns the "subtrees" of a tree.
var tail = function (v) {
    return Data_Tuple.snd(Data_Lazy.force(v));
};

// | Create a value of type `Cofree f a` from a label and a
// | functor-full of "subtrees".
var mkCofree = function (a) {
    return function (t) {
        return Data_Lazy.defer(function (v) {
            return new Data_Tuple.Tuple(a, t);
        });
    };
};
var lazyCofree = {
    defer: function (k) {
        return Data_Lazy.defer(function (v) {
            var v1 = k(Data_Unit.unit);
            return Data_Lazy.force(v1);
        });
    }
};
var hoistCofree = function (dictFunctor) {
    return function (nat) {
        return function (v) {
            return Data_Functor.map(Data_Lazy.functorLazy)(Data_Functor.map(Data_Tuple.functorTuple)((function () {
                var $139 = Data_Functor.map(dictFunctor)(hoistCofree(dictFunctor)(nat));
                return function ($140) {
                    return nat($139($140));
                };
            })()))(v);
        };
    };
};

// | Returns the label for a tree.
var head = function (v) {
    return Data_Tuple.fst(Data_Lazy.force(v));
};
var functorCofree = function (dictFunctor) {
    return {
        map: function (f) {
            var loop = function (v) {
                return Data_Functor.map(Data_Lazy.functorLazy)(function (v1) {
                    return new Data_Tuple.Tuple(f(v1.value0), Data_Functor.map(dictFunctor)(loop)(v1.value1));
                })(v);
            };
            return loop;
        }
    };
};
var functorWithIndexCofree = function (dictFunctor) {
    var functorCofree1 = functorCofree(dictFunctor);
    return {
        mapWithIndex: function (f) {
            var loop = function (n) {
                return function (v) {
                    return Data_Functor.map(Data_Lazy.functorLazy)(function (v1) {
                        return new Data_Tuple.Tuple(f(n)(v1.value0), Data_Functor.map(dictFunctor)(loop(n + 1 | 0))(v1.value1));
                    })(v);
                };
            };
            return loop(0);
        },
        Functor0: function () {
            return functorCofree1;
        }
    };
};
var foldableCofree = function (dictFoldable) {
    return {
        foldr: function (f) {
            var go = function (fa) {
                return function (b) {
                    return f(head(fa))(Data_Foldable.foldr(dictFoldable)(go)(b)(tail(fa)));
                };
            };
            return Data_Function.flip(go);
        },
        foldl: function (f) {
            var go = function (b) {
                return function (fa) {
                    return Data_Foldable.foldl(dictFoldable)(go)(f(b)(head(fa)))(tail(fa));
                };
            };
            return go;
        },
        foldMap: function (dictMonoid) {
            var Semigroup0 = dictMonoid.Semigroup0();
            return function (f) {
                var go = function (fa) {
                    return Data_Semigroup.append(Semigroup0)(f(head(fa)))(Data_Foldable.foldMap(dictFoldable)(dictMonoid)(go)(tail(fa)));
                };
                return go;
            };
        }
    };
};
var foldableWithIndexCofree = function (dictFoldable) {
    var foldableCofree1 = foldableCofree(dictFoldable);
    return {
        foldrWithIndex: function (f) {
            var go = function (ix) {
                return function (b) {
                    return function (fa) {
                        return f(ix)(head(fa))(Data_Foldable.foldr(dictFoldable)(Data_Function.flip(go(ix + 1 | 0)))(b)(tail(fa)));
                    };
                };
            };
            return go(0);
        },
        foldlWithIndex: function (f) {
            var go = function (ix) {
                return function (b) {
                    return function (fa) {
                        return Data_Foldable.foldl(dictFoldable)(go(ix + 1 | 0))(f(ix)(b)(head(fa)))(tail(fa));
                    };
                };
            };
            return go(0);
        },
        foldMapWithIndex: function (dictMonoid) {
            var Semigroup0 = dictMonoid.Semigroup0();
            return function (f) {
                var go = function (ix) {
                    return function (fa) {
                        return Data_Semigroup.append(Semigroup0)(f(ix)(head(fa)))(Data_Foldable.foldMap(dictFoldable)(dictMonoid)(go(ix + 1 | 0))(tail(fa)));
                    };
                };
                return go(0);
            };
        },
        Foldable0: function () {
            return foldableCofree1;
        }
    };
};
var traversableCofree = function (dictTraversable) {
    var functorCofree1 = functorCofree(dictTraversable.Functor0());
    var foldableCofree1 = foldableCofree(dictTraversable.Foldable1());
    return {
        sequence: function (dictApplicative) {
            return Data_Traversable.traverse(traversableCofree(dictTraversable))(dictApplicative)(identity);
        },
        traverse: function (dictApplicative) {
            var Apply0 = dictApplicative.Apply0();
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (f) {
                var loop = function (ta) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(mkCofree)(f(head(ta))))(Data_Traversable.traverse(dictTraversable)(dictApplicative)(loop)(tail(ta)));
                };
                return loop;
            };
        },
        Functor0: function () {
            return functorCofree1;
        },
        Foldable1: function () {
            return foldableCofree1;
        }
    };
};
var traversableWithIndexCofree = function (dictTraversable) {
    var functorWithIndexCofree1 = functorWithIndexCofree(dictTraversable.Functor0());
    var foldableWithIndexCofree1 = foldableWithIndexCofree(dictTraversable.Foldable1());
    var traversableCofree1 = traversableCofree(dictTraversable);
    return {
        traverseWithIndex: function (dictApplicative) {
            var Apply0 = dictApplicative.Apply0();
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (f) {
                var loop = function (ix) {
                    return function (ta) {
                        return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(mkCofree)(f(ix)(head(ta))))(Data_Traversable.traverse(dictTraversable)(dictApplicative)(loop(ix + 1 | 0))(tail(ta)));
                    };
                };
                return loop(0);
            };
        },
        FunctorWithIndex0: function () {
            return functorWithIndexCofree1;
        },
        FoldableWithIndex1: function () {
            return foldableWithIndexCofree1;
        },
        Traversable2: function () {
            return traversableCofree1;
        }
    };
};
var extendCofree = function (dictFunctor) {
    var functorCofree1 = functorCofree(dictFunctor);
    return {
        extend: function (f) {
            var loop = function (v) {
                return Data_Functor.map(Data_Lazy.functorLazy)(function (v1) {
                    return new Data_Tuple.Tuple(f(v), Data_Functor.map(dictFunctor)(loop)(v1.value1));
                })(v);
            };
            return loop;
        },
        Functor0: function () {
            return functorCofree1;
        }
    };
};
var eqCofree = function (dictEq1) {
    return function (dictEq) {
        return {
            eq: function (x) {
                return function (y) {
                    return Data_Eq.eq(dictEq)(head(x))(head(y)) && Data_Eq.eq1(dictEq1)(eqCofree(dictEq1)(dictEq))(tail(x))(tail(y));
                };
            }
        };
    };
};
var ordCofree = function (dictOrd1) {
    var eqCofree1 = eqCofree(dictOrd1.Eq10());
    return function (dictOrd) {
        var eqCofree2 = eqCofree1(dictOrd.Eq0());
        return {
            compare: function (x) {
                return function (y) {
                    var v = Data_Ord.compare(dictOrd)(head(x))(head(y));
                    if (v instanceof Data_Ordering.EQ) {
                        return Data_Ord.compare1(dictOrd1)(ordCofree(dictOrd1)(dictOrd))(tail(x))(tail(y));
                    };
                    return v;
                };
            },
            Eq0: function () {
                return eqCofree2;
            }
        };
    };
};
var eq1Cofree = function (dictEq1) {
    var eqCofree1 = eqCofree(dictEq1);
    return {
        eq1: function (dictEq) {
            return Data_Eq.eq(eqCofree1(dictEq));
        }
    };
};
var ord1Cofree = function (dictOrd1) {
    var ordCofree1 = ordCofree(dictOrd1);
    var eq1Cofree1 = eq1Cofree(dictOrd1.Eq10());
    return {
        compare1: function (dictOrd) {
            return Data_Ord.compare(ordCofree1(dictOrd));
        },
        Eq10: function () {
            return eq1Cofree1;
        }
    };
};

// | Lazily creates a value of type `Cofree f a` from a label and a
// | functor-full of "subtrees".
var deferCofree = function ($141) {
    return Cofree(Data_Lazy.defer($141));
};
var semigroupCofree = function (dictApply) {
    var Functor0 = dictApply.Functor0();
    return function (dictSemigroup) {
        return {
            append: function (x) {
                return function (y) {
                    return deferCofree(function (v) {
                        return new Data_Tuple.Tuple(Data_Semigroup.append(dictSemigroup)(head(x))(head(y)), Control_Apply.apply(dictApply)(Data_Functor.map(Functor0)(Data_Semigroup.append(semigroupCofree(dictApply)(dictSemigroup)))(tail(x)))(tail(y)));
                    });
                };
            }
        };
    };
};
var monoidCofree = function (dictApplicative) {
    var semigroupCofree1 = semigroupCofree(dictApplicative.Apply0());
    return function (dictMonoid) {
        var semigroupCofree2 = semigroupCofree1(dictMonoid.Semigroup0());
        return {
            mempty: deferCofree(function (v) {
                return new Data_Tuple.Tuple(Data_Monoid.mempty(dictMonoid), Control_Applicative.pure(dictApplicative)(Data_Monoid.mempty(monoidCofree(dictApplicative)(dictMonoid))));
            }),
            Semigroup0: function () {
                return semigroupCofree2;
            }
        };
    };
};
var comonadCofree = function (dictFunctor) {
    var extendCofree1 = extendCofree(dictFunctor);
    return {
        extract: head,
        Extend0: function () {
            return extendCofree1;
        }
    };
};

// | Explore a value in the cofree comonad by using an expression in a
// | corresponding free monad.
// |
// | The free monad should be built from a functor which pairs with the
// | functor underlying the cofree comonad.
var explore = function (dictFunctor) {
    return function (dictFunctor1) {
        var comonadCofree1 = comonadCofree(dictFunctor1);
        return function (pair) {
            return function (m) {
                return function (w) {
                    var step = function (ff) {
                        return Control_Monad_State_Class.state(monadStateStateT)(function (cof) {
                            return pair(Data_Functor.map(dictFunctor)(Data_Tuple.Tuple.create)(ff))(tail(cof));
                        });
                    };
                    var v = Control_Monad_State.runState(Control_Monad_Free.runFreeM(dictFunctor)(monadRecStateT)(step)(m))(w);
                    return v.value0(Control_Comonad.extract(comonadCofree1)(v.value1));
                };
            };
        };
    };
};
var exploreM = function (dictFunctor) {
    return function (dictFunctor1) {
        var comonadCofree1 = comonadCofree(dictFunctor1);
        return function (dictMonadRec) {
            var Functor0 = (((dictMonadRec.Monad0()).Bind1()).Apply0()).Functor0();
            var monadRecStateT1 = Control_Monad_State_Trans.monadRecStateT(dictMonadRec);
            return function (pair) {
                return function (m) {
                    return function (w) {
                        var step = function (ff) {
                            return function (cof) {
                                return pair(Data_Functor.map(dictFunctor)(Data_Tuple.Tuple.create)(ff))(tail(cof));
                            };
                        };
                        var $$eval = function (v) {
                            return v.value0(Control_Comonad.extract(comonadCofree1)(v.value1));
                        };
                        return Data_Functor.map(Functor0)($$eval)(Control_Monad_State_Trans.runStateT(Control_Monad_Free.runFreeM(dictFunctor)(monadRecStateT1)(step)(m))(w));
                    };
                };
            };
        };
    };
};

// | Recursively unfolds a `Cofree` structure given a seed.
var buildCofree = function (dictFunctor) {
    return function (k) {
        return function (s) {
            return Data_Lazy.defer(function (v) {
                return Data_Functor.map(Data_Tuple.functorTuple)(Data_Functor.map(dictFunctor)(buildCofree(dictFunctor)(k)))(k(s));
            });
        };
    };
};
var monadCofree = function (dictAlternative) {
    return {
        Applicative0: function () {
            return applicativeCofree(dictAlternative);
        },
        Bind1: function () {
            return bindCofree(dictAlternative);
        }
    };
};
var bindCofree = function (dictAlternative) {
    var Plus1 = dictAlternative.Plus1();
    var Alt0 = Plus1.Alt0();
    var Functor0 = (Plus1.Alt0()).Functor0();
    return {
        bind: function (fa) {
            return function (f) {
                var loop = function (fa$prime) {
                    var fh = f(head(fa$prime));
                    return mkCofree(head(fh))(Control_Alt.alt(Alt0)(tail(fh))(Data_Functor.map(Functor0)(loop)(tail(fa$prime))));
                };
                return loop(fa);
            };
        },
        Apply0: function () {
            return applyCofree(dictAlternative);
        }
    };
};
var applyCofree = function (dictAlternative) {
    var functorCofree1 = functorCofree(((dictAlternative.Plus1()).Alt0()).Functor0());
    return {
        apply: Control_Monad.ap(monadCofree(dictAlternative)),
        Functor0: function () {
            return functorCofree1;
        }
    };
};
var applicativeCofree = function (dictAlternative) {
    var empty = Control_Plus.empty(dictAlternative.Plus1());
    return {
        pure: function (a) {
            return mkCofree(a)(empty);
        },
        Apply0: function () {
            return applyCofree(dictAlternative);
        }
    };
};
export {
    deferCofree,
    mkCofree,
    head,
    tail,
    hoistCofree,
    buildCofree,
    explore,
    exploreM,
    semigroupCofree,
    monoidCofree,
    eqCofree,
    eq1Cofree,
    ordCofree,
    ord1Cofree,
    functorCofree,
    functorWithIndexCofree,
    foldableCofree,
    foldableWithIndexCofree,
    traversableCofree,
    traversableWithIndexCofree,
    extendCofree,
    comonadCofree,
    applyCofree,
    applicativeCofree,
    bindCofree,
    monadCofree,
    lazyCofree
};
