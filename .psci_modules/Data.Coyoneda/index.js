import * as Control_Alt from "../Control.Alt/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Comonad from "../Control.Comonad/index.js";
import * as Control_Extend from "../Control.Extend/index.js";
import * as Control_Plus from "../Control.Plus/index.js";
import * as Data_Distributive from "../Data.Distributive/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Exists from "../Data.Exists/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Functor_Invariant from "../Data.Functor.Invariant/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Semigroup_Foldable from "../Data.Semigroup.Foldable/index.js";
import * as Data_Semigroup_Traversable from "../Data.Semigroup.Traversable/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";

// | `Coyoneda` is encoded as an existential type using `Data.Exists`.
// |
// | This type constructor encodes the contents of the existential package.
var CoyonedaF = /* #__PURE__ */ (function () {
    function CoyonedaF(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    CoyonedaF.create = function (value0) {
        return function (value1) {
            return new CoyonedaF(value0, value1);
        };
    };
    return CoyonedaF;
})();

// | The `Coyoneda` `Functor`.
// |
// | `Coyoneda f` is a `Functor` for any type constructor `f`. In fact,
// | it is the _free_ `Functor` for `f`, i.e. any natural transformation
// | `nat :: f ~> g`, can be factor through `liftCoyoneda`.  The natural
// | transformation from `Coyoneda f ~> g` is given by `lowerCoyoneda <<<
// | hoistCoyoneda nat`:
// | ```purescript
// | lowerCoyoneda <<< hoistCoyoneda nat <<< liftCoyoneda $ fi
// | = lowerCoyoneda (hoistCoyoneda nat (Coyoneda $ mkExists $ CoyonedaF identity fi))    (by definition of liftCoyoneda)
// | = lowerCoyoneda (coyoneda identity (nat fi))                                         (by definition of hoistCoyoneda)
// | = unCoyoneda map (coyoneda identity (nat fi))                                        (by definition of lowerCoyoneda)
// | = unCoyoneda map (Coyoneda $ mkExists $ CoyonedaF  identity (nat fi))                (by definition of coyoneda)
// | = map identity (nat fi)                                                              (by definition of unCoyoneda)
// | = nat fi                                                                       (since g is a Functor)
// | ```
var Coyoneda = function (x) {
    return x;
};

// | Deconstruct a value of `Coyoneda a` to retrieve the mapping function and
// | original value.
var unCoyoneda = function (f) {
    return function (v) {
        return Data_Exists.runExists(function (v1) {
            return f(v1.value0)(v1.value1);
        })(v);
    };
};

// | Lower a value of type `Coyoneda f a` to the `Functor` `f`.
var lowerCoyoneda = function (dictFunctor) {
    return unCoyoneda(Data_Functor.map(dictFunctor));
};
var foldableCoyoneda = function (dictFoldable) {
    return {
        foldr: function (f) {
            return function (z) {
                return unCoyoneda(function (k) {
                    return Data_Foldable.foldr(dictFoldable)(function ($143) {
                        return f(k($143));
                    })(z);
                });
            };
        },
        foldl: function (f) {
            return function (z) {
                return unCoyoneda(function (k) {
                    return Data_Foldable.foldl(dictFoldable)(function (x) {
                        var $144 = f(x);
                        return function ($145) {
                            return $144(k($145));
                        };
                    })(z);
                });
            };
        },
        foldMap: function (dictMonoid) {
            return function (f) {
                return unCoyoneda(function (k) {
                    return Data_Foldable.foldMap(dictFoldable)(dictMonoid)(function ($146) {
                        return f(k($146));
                    });
                });
            };
        }
    };
};
var foldable1Coyoneda = function (dictFoldable1) {
    var foldableCoyoneda1 = foldableCoyoneda(dictFoldable1.Foldable0());
    return {
        foldMap1: function (dictSemigroup) {
            return function (f) {
                return unCoyoneda(function (k) {
                    return Data_Semigroup_Foldable.foldMap1(dictFoldable1)(dictSemigroup)(function ($147) {
                        return f(k($147));
                    });
                });
            };
        },
        foldr1: Data_Semigroup_Foldable.foldr1Default(foldable1Coyoneda(dictFoldable1)),
        foldl1: Data_Semigroup_Foldable.foldl1Default(foldable1Coyoneda(dictFoldable1)),
        Foldable0: function () {
            return foldableCoyoneda1;
        }
    };
};
var eqCoyoneda = function (dictFunctor) {
    return function (dictEq1) {
        return function (dictEq) {
            return {
                eq: function (x) {
                    return function (y) {
                        return Data_Eq.eq1(dictEq1)(dictEq)(lowerCoyoneda(dictFunctor)(x))(lowerCoyoneda(dictFunctor)(y));
                    };
                }
            };
        };
    };
};
var ordCoyoneda = function (dictFunctor) {
    var eqCoyoneda1 = eqCoyoneda(dictFunctor);
    return function (dictOrd1) {
        var eqCoyoneda2 = eqCoyoneda1(dictOrd1.Eq10());
        return function (dictOrd) {
            var eqCoyoneda3 = eqCoyoneda2(dictOrd.Eq0());
            return {
                compare: function (x) {
                    return function (y) {
                        return Data_Ord.compare1(dictOrd1)(dictOrd)(lowerCoyoneda(dictFunctor)(x))(lowerCoyoneda(dictFunctor)(y));
                    };
                },
                Eq0: function () {
                    return eqCoyoneda3;
                }
            };
        };
    };
};
var eq1Coyoneda = function (dictFunctor) {
    var eqCoyoneda1 = eqCoyoneda(dictFunctor);
    return function (dictEq1) {
        var eqCoyoneda2 = eqCoyoneda1(dictEq1);
        return {
            eq1: function (dictEq) {
                return Data_Eq.eq(eqCoyoneda2(dictEq));
            }
        };
    };
};
var ord1Coyoneda = function (dictFunctor) {
    var ordCoyoneda1 = ordCoyoneda(dictFunctor);
    var eq1Coyoneda1 = eq1Coyoneda(dictFunctor);
    return function (dictOrd1) {
        var ordCoyoneda2 = ordCoyoneda1(dictOrd1);
        var eq1Coyoneda2 = eq1Coyoneda1(dictOrd1.Eq10());
        return {
            compare1: function (dictOrd) {
                return Data_Ord.compare(ordCoyoneda2(dictOrd));
            },
            Eq10: function () {
                return eq1Coyoneda2;
            }
        };
    };
};

// | Construct a value of type `Coyoneda f b` from a mapping function and a
// | value of type `f a`.
var coyoneda = function (k) {
    return function (fi) {
        return Data_Exists.mkExists(new CoyonedaF(k, fi));
    };
};
var functorCoyoneda = {
    map: function (f) {
        return function (v) {
            return Data_Exists.runExists(function (v1) {
                return coyoneda(function ($148) {
                    return f(v1.value0($148));
                })(v1.value1);
            })(v);
        };
    }
};
var invatiantCoyoneda = {
    imap: /* #__PURE__ */ Data_Functor_Invariant.imapF(functorCoyoneda)
};

// | Use a natural transformation to change the generating type constructor of a
// | `Coyoneda`.
var hoistCoyoneda = function (nat) {
    return function (v) {
        return Data_Exists.runExists(function (v1) {
            return coyoneda(v1.value0)(nat(v1.value1));
        })(v);
    };
};

// | Lift a value described by the type constructor `f` to `Coyoneda f`.
// |
// | Note that for any functor `f` `liftCoyoneda` has a right inverse
// | `lowerCoyoneda`:
// | ```purescript
// | liftCoyoneda <<< lowerCoyoneda $ (Coyoneda e)
// | = liftCoyoneda <<< unCoyoneda map $ (Coyoneda e)
// | = liftCoyonead (runExists (\(CoyonedaF k fi) -> map k fi) e)
// | = liftCoyonead (Coyoneda e)
// | = coyoneda identity (Coyoneda e)
// | = Coyoneda e
// | ```
// | Moreover if `f` is a `Functor` then `liftCoyoneda` is an isomorphism of
// | functors with inverse `lowerCoyoneda`:  we already showed that
// | `lowerCoyoneda <<< hoistCoyoneda identity = lowerCoyoneda` is its left inverse
// | whenever `f` is a functor.
var liftCoyoneda = /* #__PURE__ */ coyoneda(/* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn));
var distributiveCoyoneda = function (dictDistributive) {
    var Functor0 = dictDistributive.Functor0();
    var lowerCoyoneda1 = lowerCoyoneda(Functor0);
    var lowerCoyoneda2 = lowerCoyoneda(Functor0);
    return {
        collect: function (dictFunctor) {
            return function (f) {
                var $149 = Data_Distributive.collect(dictDistributive)(dictFunctor)(function ($151) {
                    return lowerCoyoneda1(f($151));
                });
                return function ($150) {
                    return liftCoyoneda($149($150));
                };
            };
        },
        distribute: function (dictFunctor) {
            var $152 = Data_Distributive.collect(dictDistributive)(dictFunctor)(lowerCoyoneda2);
            return function ($153) {
                return liftCoyoneda($152($153));
            };
        },
        Functor0: function () {
            return functorCoyoneda;
        }
    };
};
var extendCoyoneda = function (dictExtend) {
    return {
        extend: function (f) {
            return function (v) {
                return Data_Exists.runExists(function (v1) {
                    return liftCoyoneda(Control_Extend.extend(dictExtend)((function () {
                        var $154 = coyoneda(v1.value0);
                        return function ($155) {
                            return f($154($155));
                        };
                    })())(v1.value1));
                })(v);
            };
        },
        Functor0: function () {
            return functorCoyoneda;
        }
    };
};
var monadTransCoyoneda = {
    lift: function (dictMonad) {
        return liftCoyoneda;
    }
};
var traversableCoyoneda = function (dictTraversable) {
    var foldableCoyoneda1 = foldableCoyoneda(dictTraversable.Foldable1());
    return {
        traverse: function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return function (f) {
                return unCoyoneda(function (k) {
                    var $156 = Data_Functor.map(Functor0)(liftCoyoneda);
                    var $157 = Data_Traversable.traverse(dictTraversable)(dictApplicative)(function ($159) {
                        return f(k($159));
                    });
                    return function ($158) {
                        return $156($157($158));
                    };
                });
            };
        },
        sequence: function (dictApplicative) {
            var Functor0 = (dictApplicative.Apply0()).Functor0();
            return unCoyoneda(function (k) {
                var $160 = Data_Functor.map(Functor0)(liftCoyoneda);
                var $161 = Data_Traversable.traverse(dictTraversable)(dictApplicative)(k);
                return function ($162) {
                    return $160($161($162));
                };
            });
        },
        Functor0: function () {
            return functorCoyoneda;
        },
        Foldable1: function () {
            return foldableCoyoneda1;
        }
    };
};
var traversable1Coyoneda = function (dictTraversable1) {
    var sequence1 = Data_Semigroup_Traversable.sequence1(dictTraversable1);
    var Traversable1 = dictTraversable1.Traversable1();
    var Functor0 = Traversable1.Functor0();
    var foldable1Coyoneda1 = foldable1Coyoneda(dictTraversable1.Foldable10());
    var traversableCoyoneda1 = traversableCoyoneda(Traversable1);
    return {
        traverse1: function (dictApply) {
            var Functor01 = dictApply.Functor0();
            return function (f) {
                return unCoyoneda(function (k) {
                    var $163 = Data_Functor.map(Functor01)(liftCoyoneda);
                    var $164 = Data_Semigroup_Traversable.traverse1(dictTraversable1)(dictApply)(function ($166) {
                        return f(k($166));
                    });
                    return function ($165) {
                        return $163($164($165));
                    };
                });
            };
        },
        sequence1: function (dictApply) {
            var Functor01 = dictApply.Functor0();
            var sequence11 = sequence1(dictApply);
            return unCoyoneda(function (k) {
                var $167 = Data_Functor.map(Functor01)(liftCoyoneda);
                var $168 = Data_Functor.map(Functor0)(k);
                return function ($169) {
                    return $167(sequence11($168($169)));
                };
            });
        },
        Foldable10: function () {
            return foldable1Coyoneda1;
        },
        Traversable1: function () {
            return traversableCoyoneda1;
        }
    };
};

// | As in the monad case: if `w` is a comonad, then it is a functor, thus
// | `liftCoyoneda` is an iso of functors, but moreover it is an iso of
// | comonads, i.e. the following law holds:
// | ```purescript
// | g <<= liftCoyoneda w = liftCoyoneda $ g <<< liftCoyoneda <<= w
// | ```
var comonadCoyoneda = function (dictComonad) {
    var extendCoyoneda1 = extendCoyoneda(dictComonad.Extend0());
    return {
        extract: function (v) {
            return Data_Exists.runExists(function (v1) {
                return v1.value0(Control_Comonad.extract(dictComonad)(v1.value1));
            })(v);
        },
        Extend0: function () {
            return extendCoyoneda1;
        }
    };
};
var applyCoyoneda = function (dictApply) {
    var Functor0 = dictApply.Functor0();
    return {
        apply: function (f) {
            return function (g) {
                return liftCoyoneda(Control_Apply.apply(dictApply)(lowerCoyoneda(Functor0)(f))(lowerCoyoneda(Functor0)(g)));
            };
        },
        Functor0: function () {
            return functorCoyoneda;
        }
    };
};
var bindCoyoneda = function (dictBind) {
    var Apply0 = dictBind.Apply0();
    var lowerCoyoneda1 = lowerCoyoneda(Apply0.Functor0());
    var applyCoyoneda1 = applyCoyoneda(Apply0);
    return {
        bind: function (v) {
            return function (f) {
                return liftCoyoneda(Data_Exists.runExists(function (v1) {
                    return Control_Bind.bindFlipped(dictBind)(function ($170) {
                        return lowerCoyoneda1(f(v1.value0($170)));
                    })(v1.value1);
                })(v));
            };
        },
        Apply0: function () {
            return applyCoyoneda1;
        }
    };
};
var applicativeCoyoneda = function (dictApplicative) {
    var applyCoyoneda1 = applyCoyoneda(dictApplicative.Apply0());
    return {
        pure: (function () {
            var $171 = Control_Applicative.pure(dictApplicative);
            return function ($172) {
                return liftCoyoneda($171($172));
            };
        })(),
        Apply0: function () {
            return applyCoyoneda1;
        }
    };
};

// | When `f` is a Monad then it is a functor as well.  In this case
// | `liftCoyoneda` is not only a functor isomorphism but also a monad
// | isomorphism, i.e. the following law holds
// | ```purescript
// | liftCoyoneda fa >>= liftCoyoneda <<< g = liftCoyoneda $ fa >>= g
// | ```
var monadCoyoneda = function (dictMonad) {
    var applicativeCoyoneda1 = applicativeCoyoneda(dictMonad.Applicative0());
    var bindCoyoneda1 = bindCoyoneda(dictMonad.Bind1());
    return {
        Applicative0: function () {
            return applicativeCoyoneda1;
        },
        Bind1: function () {
            return bindCoyoneda1;
        }
    };
};
var altCoyoneda = function (dictAlt) {
    var Functor0 = dictAlt.Functor0();
    return {
        alt: function (x) {
            return function (y) {
                return liftCoyoneda(Control_Alt.alt(dictAlt)(lowerCoyoneda(Functor0)(x))(lowerCoyoneda(Functor0)(y)));
            };
        },
        Functor0: function () {
            return functorCoyoneda;
        }
    };
};
var plusCoyoneda = function (dictPlus) {
    var altCoyoneda1 = altCoyoneda(dictPlus.Alt0());
    return {
        empty: liftCoyoneda(Control_Plus.empty(dictPlus)),
        Alt0: function () {
            return altCoyoneda1;
        }
    };
};
var alternativeCoyoneda = function (dictAlternative) {
    var applicativeCoyoneda1 = applicativeCoyoneda(dictAlternative.Applicative0());
    var plusCoyoneda1 = plusCoyoneda(dictAlternative.Plus1());
    return {
        Applicative0: function () {
            return applicativeCoyoneda1;
        },
        Plus1: function () {
            return plusCoyoneda1;
        }
    };
};
var monadPlusCoyoneda = function (dictMonadPlus) {
    var monadCoyoneda1 = monadCoyoneda(dictMonadPlus.Monad0());
    var alternativeCoyoneda1 = alternativeCoyoneda(dictMonadPlus.Alternative1());
    return {
        Monad0: function () {
            return monadCoyoneda1;
        },
        Alternative1: function () {
            return alternativeCoyoneda1;
        }
    };
};
export {
    Coyoneda,
    coyoneda,
    unCoyoneda,
    liftCoyoneda,
    lowerCoyoneda,
    hoistCoyoneda,
    eqCoyoneda,
    eq1Coyoneda,
    ordCoyoneda,
    ord1Coyoneda,
    functorCoyoneda,
    invatiantCoyoneda,
    applyCoyoneda,
    applicativeCoyoneda,
    altCoyoneda,
    plusCoyoneda,
    alternativeCoyoneda,
    bindCoyoneda,
    monadCoyoneda,
    monadTransCoyoneda,
    monadPlusCoyoneda,
    extendCoyoneda,
    comonadCoyoneda,
    foldableCoyoneda,
    traversableCoyoneda,
    foldable1Coyoneda,
    traversable1Coyoneda,
    distributiveCoyoneda
};
