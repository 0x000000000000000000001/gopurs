import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Monad from "../Control.Monad/index.js";
import * as Control_Monad_Rec_Class from "../Control.Monad.Rec.Class/index.js";
import * as Data_CatList from "../Data.CatList/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Unsafe_Coerce from "../Unsafe.Coerce/index.js";
var $runtime_lazy = function (name, moduleName, init) {
    var state = 0;
    var val;
    return function (lineNumber) {
        if (state === 2) return val;
        if (state === 1) throw new ReferenceError(name + " was needed before it finished initializing (module " + moduleName + ", line " + lineNumber + ")", moduleName, lineNumber);
        state = 1;
        val = init();
        state = 2;
        return val;
    };
};
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var ExpF = function (x) {
    return x;
};

// | The free monad for a type constructor `f`.
// |
// | Implemented in the spirit of [Reflection without Remorse](http://okmij.org/ftp/Haskell/zseq.pdf),
// | the free monad is represented using a sequential data structure in
// | order to overcome the quadratic complexity of left-associated binds
// | and traversal through the free monad structure.
var Free = /* #__PURE__ */ (function () {
    function Free(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Free.create = function (value0) {
        return function (value1) {
            return new Free(value0, value1);
        };
    };
    return Free;
})();
var Return = /* #__PURE__ */ (function () {
    function Return(value0) {
        this.value0 = value0;
    };
    Return.create = function (value0) {
        return new Return(value0);
    };
    return Return;
})();
var Bind = /* #__PURE__ */ (function () {
    function Bind(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Bind.create = function (value0) {
        return function (value1) {
            return new Bind(value0, value1);
        };
    };
    return Bind;
})();
var toView = function ($copy_v) {
    var $tco_done = false;
    var $tco_result;
    function $tco_loop(v) {
        var runExpF = function (v2) {
            return v2;
        };
        var concatF = function (v2) {
            return function (r) {
                return new Free(v2.value0, Data_Semigroup.append(Data_CatList.semigroupCatList)(v2.value1)(r));
            };
        };
        if (v.value0 instanceof Return) {
            var v2 = Data_CatList.uncons(v.value1);
            if (v2 instanceof Data_Maybe.Nothing) {
                $tco_done = true;
                return new Return(v.value0.value0);
            };
            if (v2 instanceof Data_Maybe.Just) {
                $copy_v = concatF(runExpF(v2.value0.value0)(v.value0.value0))(v2.value0.value1);
                return;
            };
            throw new Error("Failed pattern match at Control.Monad.Free (line 227, column 7 - line 231, column 64): " + [ v2.constructor.name ]);
        };
        if (v.value0 instanceof Bind) {
            $tco_done = true;
            return new Bind(v.value0.value0, function (a) {
                return concatF(v.value0.value1(a))(v.value1);
            });
        };
        throw new Error("Failed pattern match at Control.Monad.Free (line 225, column 3 - line 233, column 56): " + [ v.value0.constructor.name ]);
    };
    while (!$tco_done) {
        $tco_result = $tco_loop($copy_v);
    };
    return $tco_result;
};

// | Run a free monad with a function mapping a functor `f` to a tail-recursive
// | monad `m`. See the `MonadRec` type class for more details.
var runFreeM = function (dictFunctor) {
    return function (dictMonadRec) {
        var Monad0 = dictMonadRec.Monad0();
        var Functor0 = ((Monad0.Bind1()).Apply0()).Functor0();
        var Applicative0 = Monad0.Applicative0();
        return function (k) {
            var go = function (f) {
                var v = toView(f);
                if (v instanceof Return) {
                    return Data_Functor.map(Functor0)(Control_Monad_Rec_Class.Done.create)(Control_Applicative.pure(Applicative0)(v.value0));
                };
                if (v instanceof Bind) {
                    return Data_Functor.map(Functor0)(Control_Monad_Rec_Class.Loop.create)(k(Data_Functor.map(dictFunctor)(v.value1)(v.value0)));
                };
                throw new Error("Failed pattern match at Control.Monad.Free (line 194, column 10 - line 196, column 37): " + [ v.constructor.name ]);
            };
            return Control_Monad_Rec_Class.tailRecM(dictMonadRec)(go);
        };
    };
};

// | Run a free monad with a function that unwraps a single layer of the functor
// | `f` at a time.
var runFree = function (dictFunctor) {
    return function (k) {
        var go = function ($copy_f) {
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(f) {
                var v = toView(f);
                if (v instanceof Return) {
                    $tco_done = true;
                    return v.value0;
                };
                if (v instanceof Bind) {
                    $copy_f = k(Data_Functor.map(dictFunctor)(v.value1)(v.value0));
                    return;
                };
                throw new Error("Failed pattern match at Control.Monad.Free (line 178, column 10 - line 180, column 33): " + [ v.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($copy_f);
            };
            return $tco_result;
        };
        return go;
    };
};

// | Unwraps a single layer of `f`, providing the continuation.
var resume$prime = function (k) {
    return function (j) {
        return function (f) {
            var v = toView(f);
            if (v instanceof Return) {
                return j(v.value0);
            };
            if (v instanceof Bind) {
                return k(v.value0)(v.value1);
            };
            throw new Error("Failed pattern match at Control.Monad.Free (line 213, column 17 - line 215, column 20): " + [ v.constructor.name ]);
        };
    };
};

// | Unwraps a single layer of the functor `f`.
var resume = function (dictFunctor) {
    return resume$prime(function (g) {
        return function (i) {
            return new Data_Either.Left(Data_Functor.map(dictFunctor)(i)(g));
        };
    })(Data_Either.Right.create);
};
var fromView = function (f) {
    return new Free(f, Data_CatList.empty);
};

// | Add a layer.
var wrap = function (f) {
    return fromView(new Bind(f, Unsafe_Coerce.unsafeCoerce));
};

// | Suspend a value given the applicative functor `f` into the free monad.
var suspendF = function (dictApplicative) {
    return function (f) {
        return wrap(Control_Applicative.pure(dictApplicative)(f));
    };
};
var freeMonad = {
    Applicative0: function () {
        return freeApplicative;
    },
    Bind1: function () {
        return freeBind;
    }
};
var freeFunctor = {
    map: function (k) {
        return function (f) {
            return Control_Bind.bindFlipped(freeBind)((function () {
                var $160 = Control_Applicative.pure(freeApplicative);
                return function ($161) {
                    return $160(k($161));
                };
            })())(f);
        };
    }
};
var freeBind = {
    bind: function (v) {
        return function (k) {
            return new Free(v.value0, Data_CatList.snoc(v.value1)(k));
        };
    },
    Apply0: function () {
        return $lazy_freeApply(0);
    }
};
var freeApplicative = {
    pure: function ($162) {
        return fromView(Return.create($162));
    },
    Apply0: function () {
        return $lazy_freeApply(0);
    }
};
var $lazy_freeApply = /* #__PURE__ */ $runtime_lazy("freeApply", "Control.Monad.Free", function () {
    return {
        apply: Control_Monad.ap(freeMonad),
        Functor0: function () {
            return freeFunctor;
        }
    };
});
var freeApply = /* #__PURE__ */ $lazy_freeApply(77);
var pure = /* #__PURE__ */ Control_Applicative.pure(freeApplicative);
var join = /* #__PURE__ */ Control_Bind.join(freeBind);
var pure1 = /* #__PURE__ */ Control_Applicative.pure(freeApplicative);
var semigroupFree = function (dictSemigroup) {
    return {
        append: Control_Apply.lift2(freeApply)(Data_Semigroup.append(dictSemigroup))
    };
};
var freeMonadRec = {
    tailRecM: function (k) {
        return function (a) {
            return Control_Bind.bind(freeBind)(k(a))(function (v) {
                if (v instanceof Control_Monad_Rec_Class.Loop) {
                    return Control_Monad_Rec_Class.tailRecM(freeMonadRec)(k)(v.value0);
                };
                if (v instanceof Control_Monad_Rec_Class.Done) {
                    return Control_Applicative.pure(freeApplicative)(v.value0);
                };
                throw new Error("Failed pattern match at Control.Monad.Free (line 86, column 26 - line 88, column 21): " + [ v.constructor.name ]);
            });
        };
    },
    Monad0: function () {
        return freeMonad;
    }
};

// | Lift an impure value described by the generating type constructor `f` into
// | the free monad.
var liftF = function (f) {
    return fromView(new Bind(f, function ($163) {
        return pure($163);
    }));
};
var freeMonadTrans = {
    lift: function (dictMonad) {
        return liftF;
    }
};
var monoidFree = function (dictMonoid) {
    var semigroupFree1 = semigroupFree(dictMonoid.Semigroup0());
    return {
        mempty: Control_Applicative.pure(freeApplicative)(Data_Monoid.mempty(dictMonoid)),
        Semigroup0: function () {
            return semigroupFree1;
        }
    };
};

// | Like `foldFree`, but for folding into some other Free monad without the
// | overhead that `MonadRec` incurs.
var substFree = function (k) {
    var go = function (f) {
        var v = toView(f);
        if (v instanceof Return) {
            return Control_Applicative.pure(freeApplicative)(v.value0);
        };
        if (v instanceof Bind) {
            return Control_Bind.bind(freeBind)(k(v.value0))(Data_Functor.map(Data_Functor.functorFn)(go)(v.value1));
        };
        throw new Error("Failed pattern match at Control.Monad.Free (line 168, column 10 - line 170, column 33): " + [ v.constructor.name ]);
    };
    return go;
};

// | Use a natural transformation to change the generating type constructor of a
// | free monad.
var hoistFree = function (k) {
    return substFree(function ($164) {
        return liftF(k($164));
    });
};
var foldableFree = function (dictFunctor) {
    var resume1 = resume(dictFunctor);
    return function (dictFoldable) {
        return {
            foldMap: function (dictMonoid) {
                return function (f) {
                    var $lazy_go = $runtime_lazy("go", "Control.Monad.Free", function () {
                        return function ($165) {
                            return (function (v) {
                                if (v instanceof Data_Either.Left) {
                                    return Data_Foldable.foldMap(dictFoldable)(dictMonoid)($lazy_go(94))(v.value0);
                                };
                                if (v instanceof Data_Either.Right) {
                                    return f(v.value0);
                                };
                                throw new Error("Failed pattern match at Control.Monad.Free (line 93, column 21 - line 95, column 21): " + [ v.constructor.name ]);
                            })(resume1($165));
                        };
                    });
                    var go = $lazy_go(93);
                    return go;
                };
            },
            foldl: function (f) {
                var go = function (r) {
                    return function ($166) {
                        return (function (v) {
                            if (v instanceof Data_Either.Left) {
                                return Data_Foldable.foldl(dictFoldable)(go)(r)(v.value0);
                            };
                            if (v instanceof Data_Either.Right) {
                                return f(r)(v.value0);
                            };
                            throw new Error("Failed pattern match at Control.Monad.Free (line 98, column 23 - line 100, column 23): " + [ v.constructor.name ]);
                        })(resume1($166));
                    };
                };
                return go;
            },
            foldr: function (f) {
                var go = function (r) {
                    return function ($167) {
                        return (function (v) {
                            if (v instanceof Data_Either.Left) {
                                return Data_Foldable.foldr(dictFoldable)(Data_Function.flip(go))(r)(v.value0);
                            };
                            if (v instanceof Data_Either.Right) {
                                return f(v.value0)(r);
                            };
                            throw new Error("Failed pattern match at Control.Monad.Free (line 103, column 23 - line 105, column 23): " + [ v.constructor.name ]);
                        })(resume1($167));
                    };
                };
                return go;
            }
        };
    };
};
var traversableFree = function (dictTraversable) {
    var Functor0 = dictTraversable.Functor0();
    var resume1 = resume(Functor0);
    var foldableFree1 = foldableFree(Functor0)(dictTraversable.Foldable1());
    return {
        traverse: function (dictApplicative) {
            var Functor01 = (dictApplicative.Apply0()).Functor0();
            return function (f) {
                var $lazy_go = $runtime_lazy("go", "Control.Monad.Free", function () {
                    return function ($168) {
                        return (function (v) {
                            if (v instanceof Data_Either.Left) {
                                return Data_Functor.map(Functor01)(function ($169) {
                                    return join(liftF($169));
                                })(Data_Traversable.traverse(dictTraversable)(dictApplicative)($lazy_go(111))(v.value0));
                            };
                            if (v instanceof Data_Either.Right) {
                                return Data_Functor.map(Functor01)(pure1)(f(v.value0));
                            };
                            throw new Error("Failed pattern match at Control.Monad.Free (line 110, column 21 - line 112, column 30): " + [ v.constructor.name ]);
                        })(resume1($168));
                    };
                });
                var go = $lazy_go(110);
                return go;
            };
        },
        sequence: function (dictApplicative) {
            return function (tma) {
                return Data_Traversable.traverse(traversableFree(dictTraversable))(dictApplicative)(identity)(tma);
            };
        },
        Functor0: function () {
            return freeFunctor;
        },
        Foldable1: function () {
            return foldableFree1;
        }
    };
};

// | Run a free monad with a natural transformation from the type constructor `f`
// | to the tail-recursive monad `m`. See the `MonadRec` type class for more
// | details.
var foldFree = function (dictMonadRec) {
    var Monad0 = dictMonadRec.Monad0();
    var Functor0 = ((Monad0.Bind1()).Apply0()).Functor0();
    var Applicative0 = Monad0.Applicative0();
    return function (k) {
        var go = function (f) {
            var v = toView(f);
            if (v instanceof Return) {
                return Data_Functor.map(Functor0)(Control_Monad_Rec_Class.Done.create)(Control_Applicative.pure(Applicative0)(v.value0));
            };
            if (v instanceof Bind) {
                return Data_Functor.map(Functor0)(function ($170) {
                    return Control_Monad_Rec_Class.Loop.create(v.value1($170));
                })(k(v.value0));
            };
            throw new Error("Failed pattern match at Control.Monad.Free (line 158, column 10 - line 160, column 37): " + [ v.constructor.name ]);
        };
        return Control_Monad_Rec_Class.tailRecM(dictMonadRec)(go);
    };
};
var eqFree = function (dictFunctor) {
    return function (dictEq1) {
        return function (dictEq) {
            return {
                eq: function (x) {
                    return function (y) {
                        var v = resume(dictFunctor)(y);
                        var v1 = resume(dictFunctor)(x);
                        if (v1 instanceof Data_Either.Left && v instanceof Data_Either.Left) {
                            return Data_Eq.eq1(dictEq1)(eqFree(dictFunctor)(dictEq1)(dictEq))(v1.value0)(v.value0);
                        };
                        if (v1 instanceof Data_Either.Right && v instanceof Data_Either.Right) {
                            return Data_Eq.eq(dictEq)(v1.value0)(v.value0);
                        };
                        return false;
                    };
                }
            };
        };
    };
};
var ordFree = function (dictFunctor) {
    var eqFree1 = eqFree(dictFunctor);
    return function (dictOrd1) {
        var eqFree2 = eqFree1(dictOrd1.Eq10());
        return function (dictOrd) {
            var eqFree3 = eqFree2(dictOrd.Eq0());
            return {
                compare: function (x) {
                    return function (y) {
                        var v = resume(dictFunctor)(y);
                        var v1 = resume(dictFunctor)(x);
                        if (v1 instanceof Data_Either.Left && v instanceof Data_Either.Left) {
                            return Data_Ord.compare1(dictOrd1)(ordFree(dictFunctor)(dictOrd1)(dictOrd))(v1.value0)(v.value0);
                        };
                        if (v1 instanceof Data_Either.Left) {
                            return Data_Ordering.LT.value;
                        };
                        if (v instanceof Data_Either.Left) {
                            return Data_Ordering.GT.value;
                        };
                        if (v1 instanceof Data_Either.Right && v instanceof Data_Either.Right) {
                            return Data_Ord.compare(dictOrd)(v1.value0)(v.value0);
                        };
                        throw new Error("Failed pattern match at Control.Monad.Free (line 56, column 17 - line 60, column 36): " + [ v1.constructor.name, v.constructor.name ]);
                    };
                },
                Eq0: function () {
                    return eqFree3;
                }
            };
        };
    };
};
var eq1Free = function (dictFunctor) {
    var eqFree1 = eqFree(dictFunctor);
    return function (dictEq1) {
        var eqFree2 = eqFree1(dictEq1);
        return {
            eq1: function (dictEq) {
                return Data_Eq.eq(eqFree2(dictEq));
            }
        };
    };
};
var ord1Free = function (dictFunctor) {
    var ordFree1 = ordFree(dictFunctor);
    var eq1Free1 = eq1Free(dictFunctor);
    return function (dictOrd1) {
        var ordFree2 = ordFree1(dictOrd1);
        var eq1Free2 = eq1Free1(dictOrd1.Eq10());
        return {
            compare1: function (dictOrd) {
                return Data_Ord.compare(ordFree2(dictOrd));
            },
            Eq10: function () {
                return eq1Free2;
            }
        };
    };
};
export {
    suspendF,
    wrap,
    liftF,
    hoistFree,
    foldFree,
    substFree,
    runFree,
    runFreeM,
    resume,
    resume$prime,
    eqFree,
    eq1Free,
    ordFree,
    ord1Free,
    freeFunctor,
    freeBind,
    freeApplicative,
    freeApply,
    freeMonad,
    freeMonadTrans,
    freeMonadRec,
    foldableFree,
    traversableFree,
    semigroupFree,
    monoidFree
};
