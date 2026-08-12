import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Lazy from "../Control.Lazy/index.js";
import * as Control_Monad from "../Control.Monad/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_FoldableWithIndex from "../Data.FoldableWithIndex/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_FunctorWithIndex from "../Data.FunctorWithIndex/index.js";
import * as Data_Lazy from "../Data.Lazy/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_NonEmpty from "../Data.NonEmpty/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_TraversableWithIndex from "../Data.TraversableWithIndex/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unfoldable1 from "../Data.Unfoldable1/index.js";
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
var unwrap = /* #__PURE__ */ Data_Newtype.unwrap();

// | A lazy linked list.
var List = function (x) {
    return x;
};

// | A list is either empty (represented by the `Nil` constructor) or non-empty, in
// | which case it consists of a head element, and another list (represented by the
// | `Cons` constructor).
var Nil = /* #__PURE__ */ (function () {
    function Nil() {

    };
    Nil.value = new Nil();
    return Nil;
})();

// | A list is either empty (represented by the `Nil` constructor) or non-empty, in
// | which case it consists of a head element, and another list (represented by the
// | `Cons` constructor).
var Cons = /* #__PURE__ */ (function () {
    function Cons(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Cons.create = function (value0) {
        return function (value1) {
            return new Cons(value0, value1);
        };
    };
    return Cons;
})();
var NonEmptyList = function (x) {
    return x;
};

// | The empty list.
// |
// | Running time: `O(1)`
var nil = /* #__PURE__ */ Data_Lazy.defer(function (v) {
    return Nil.value;
});
var newtypeNonEmptyList = {
    Coercible0: function () {
        return undefined;
    }
};
var newtypeList = {
    Coercible0: function () {
        return undefined;
    }
};

// | Unwrap a lazy linked list
var step = /* #__PURE__ */ (function () {
    var $272 = Data_Newtype.unwrap();
    return function ($273) {
        return Data_Lazy.force($272($273));
    };
})();
var semigroupList = {
    append: function (xs) {
        return function (ys) {
            var go = function (v) {
                if (v instanceof Nil) {
                    return step(ys);
                };
                if (v instanceof Cons) {
                    return new Cons(v.value0, Data_Semigroup.append(semigroupList)(v.value1)(ys));
                };
                throw new Error("Failed pattern match at Data.List.Lazy.Types (line 103, column 5 - line 103, column 21): " + [ v.constructor.name ]);
            };
            return Data_Functor.map(Data_Lazy.functorLazy)(go)(Data_Newtype.unwrap()(xs));
        };
    }
};
var monoidList = {
    mempty: nil,
    Semigroup0: function () {
        return semigroupList;
    }
};
var lazyList = {
    defer: function (f) {
        return Data_Lazy.defer(function ($274) {
            return step(f($274));
        });
    }
};
var functorList = {
    map: function (f) {
        return function (xs) {
            var go = function (v) {
                if (v instanceof Nil) {
                    return Nil.value;
                };
                if (v instanceof Cons) {
                    return new Cons(f(v.value0), Data_Functor.map(functorList)(f)(v.value1));
                };
                throw new Error("Failed pattern match at Data.List.Lazy.Types (line 112, column 5 - line 112, column 17): " + [ v.constructor.name ]);
            };
            return Data_Functor.map(Data_Lazy.functorLazy)(go)(Data_Newtype.unwrap()(xs));
        };
    }
};
var functorNonEmpty = /* #__PURE__ */ Data_NonEmpty.functorNonEmpty(functorList);
var functorNonEmptyList = {
    map: function (f) {
        return function (v) {
            return Data_Functor.map(Data_Lazy.functorLazy)(Data_Functor.map(functorNonEmpty)(f))(v);
        };
    }
};
var eq1List = {
    eq1: function (dictEq) {
        return function (xs) {
            return function (ys) {
                var go = function ($copy_v) {
                    return function ($copy_v1) {
                        var $tco_var_v = $copy_v;
                        var $tco_done = false;
                        var $tco_result;
                        function $tco_loop(v, v1) {
                            if (v instanceof Nil && v1 instanceof Nil) {
                                $tco_done = true;
                                return true;
                            };
                            if (v instanceof Cons && (v1 instanceof Cons && Data_Eq.eq(dictEq)(v.value0)(v1.value0))) {
                                $tco_var_v = step(v.value1);
                                $copy_v1 = step(v1.value1);
                                return;
                            };
                            $tco_done = true;
                            return false;
                        };
                        while (!$tco_done) {
                            $tco_result = $tco_loop($tco_var_v, $copy_v1);
                        };
                        return $tco_result;
                    };
                };
                return go(step(xs))(step(ys));
            };
        };
    }
};
var eqNonEmpty = /* #__PURE__ */ Data_NonEmpty.eqNonEmpty(eq1List);
var eq1 = /* #__PURE__ */ Data_Eq.eq1(eq1List);
var eq1NonEmptyList = {
    eq1: function (dictEq) {
        var eqNonEmpty1 = eqNonEmpty(dictEq);
        return function (v) {
            return function (v1) {
                return Data_Eq.eq1(Data_Lazy.eq1Lazy)(eqNonEmpty1)(v)(v1);
            };
        };
    }
};
var eqList = function (dictEq) {
    return {
        eq: eq1(dictEq)
    };
};
var eqNonEmptyList = function (dictEq) {
    return Data_Lazy.eqLazy(eqNonEmpty(dictEq));
};
var ord1List = {
    compare1: function (dictOrd) {
        return function (xs) {
            return function (ys) {
                var go = function ($copy_v) {
                    return function ($copy_v1) {
                        var $tco_var_v = $copy_v;
                        var $tco_done = false;
                        var $tco_result;
                        function $tco_loop(v, v1) {
                            if (v instanceof Nil && v1 instanceof Nil) {
                                $tco_done = true;
                                return Data_Ordering.EQ.value;
                            };
                            if (v instanceof Nil) {
                                $tco_done = true;
                                return Data_Ordering.LT.value;
                            };
                            if (v1 instanceof Nil) {
                                $tco_done = true;
                                return Data_Ordering.GT.value;
                            };
                            if (v instanceof Cons && v1 instanceof Cons) {
                                var v2 = Data_Ord.compare(dictOrd)(v.value0)(v1.value0);
                                if (v2 instanceof Data_Ordering.EQ) {
                                    $tco_var_v = step(v.value1);
                                    $copy_v1 = step(v1.value1);
                                    return;
                                };
                                $tco_done = true;
                                return v2;
                            };
                            throw new Error("Failed pattern match at Data.List.Lazy.Types (line 89, column 5 - line 89, column 20): " + [ v.constructor.name, v1.constructor.name ]);
                        };
                        while (!$tco_done) {
                            $tco_result = $tco_loop($tco_var_v, $copy_v1);
                        };
                        return $tco_result;
                    };
                };
                return go(step(xs))(step(ys));
            };
        };
    },
    Eq10: function () {
        return eq1List;
    }
};
var ordNonEmpty = /* #__PURE__ */ Data_NonEmpty.ordNonEmpty(ord1List);
var compare1 = /* #__PURE__ */ Data_Ord.compare1(ord1List);
var ord1NonEmptyList = {
    compare1: function (dictOrd) {
        var ordNonEmpty1 = ordNonEmpty(dictOrd);
        return function (v) {
            return function (v1) {
                return Data_Ord.compare1(Data_Lazy.ord1Lazy)(ordNonEmpty1)(v)(v1);
            };
        };
    },
    Eq10: function () {
        return eq1NonEmptyList;
    }
};
var ordList = function (dictOrd) {
    var eqList1 = eqList(dictOrd.Eq0());
    return {
        compare: compare1(dictOrd),
        Eq0: function () {
            return eqList1;
        }
    };
};
var ordNonEmptyList = function (dictOrd) {
    return Data_Lazy.ordLazy(ordNonEmpty(dictOrd));
};

// | Attach an element to the front of a lazy list.
// |
// | Running time: `O(1)`
var cons = function (x) {
    return function (xs) {
        return Data_Lazy.defer(function (v) {
            return new Cons(x, xs);
        });
    };
};
var foldableList = {
    foldr: function (op) {
        return function (z) {
            return function (xs) {
                var rev = Data_Foldable.foldl(foldableList)(Data_Function.flip(cons))(nil);
                return Data_Foldable.foldl(foldableList)(Data_Function.flip(op))(z)(rev(xs));
            };
        };
    },
    foldl: function (op) {
        
        // `go` is needed to ensure the function is tail-call optimized
var go = function ($copy_b) {
            return function ($copy_xs) {
                var $tco_var_b = $copy_b;
                var $tco_done = false;
                var $tco_result;
                function $tco_loop(b, xs) {
                    var v = step(xs);
                    if (v instanceof Nil) {
                        $tco_done = true;
                        return b;
                    };
                    if (v instanceof Cons) {
                        $tco_var_b = op(b)(v.value0);
                        $copy_xs = v.value1;
                        return;
                    };
                    throw new Error("Failed pattern match at Data.List.Lazy.Types (line 127, column 7 - line 129, column 40): " + [ v.constructor.name ]);
                };
                while (!$tco_done) {
                    $tco_result = $tco_loop($tco_var_b, $copy_xs);
                };
                return $tco_result;
            };
        };
        return go;
    },
    foldMap: function (dictMonoid) {
        var Semigroup0 = dictMonoid.Semigroup0();
        var mempty = Data_Monoid.mempty(dictMonoid);
        return function (f) {
            return Data_Foldable.foldl(foldableList)(function (b) {
                return function (a) {
                    return Data_Semigroup.append(Semigroup0)(b)(f(a));
                };
            })(mempty);
        };
    }
};
var foldableNonEmpty = /* #__PURE__ */ Data_NonEmpty.foldableNonEmpty(foldableList);
var extendList = {
    extend: function (f) {
        return function (l) {
            var go = function (a) {
                return function (v) {
                    var acc$prime = cons(a)(v.acc);
                    return {
                        val: cons(f(acc$prime))(v.val),
                        acc: acc$prime
                    };
                };
            };
            var v = step(l);
            if (v instanceof Nil) {
                return nil;
            };
            if (v instanceof Cons) {
                return cons(f(l))((Data_Foldable.foldr(foldableList)(go)({
                    val: nil,
                    acc: nil
                })(v.value1)).val);
            };
            throw new Error("Failed pattern match at Data.List.Lazy.Types (line 197, column 5 - line 200, column 55): " + [ v.constructor.name ]);
        };
    },
    Functor0: function () {
        return functorList;
    }
};
var extendNonEmptyList = {
    extend: function (f) {
        return function (v) {
            var go = function (a) {
                return function (v1) {
                    return {
                        val: cons(f(Data_Lazy.defer(function (v2) {
                            return new Data_NonEmpty.NonEmpty(a, v1.acc);
                        })))(v1.val),
                        acc: cons(a)(v1.acc)
                    };
                };
            };
            var v1 = Data_Lazy.force(v);
            return Data_Lazy.defer(function (v2) {
                return new Data_NonEmpty.NonEmpty(f(v), (Data_Foldable.foldr(foldableList)(go)({
                    val: nil,
                    acc: nil
                })(v1.value1)).val);
            });
        };
    },
    Functor0: function () {
        return functorNonEmptyList;
    }
};
var foldableNonEmptyList = {
    foldr: function (f) {
        return function (b) {
            return function (v) {
                return Data_Foldable.foldr(foldableNonEmpty)(f)(b)(Data_Lazy.force(v));
            };
        };
    },
    foldl: function (f) {
        return function (b) {
            return function (v) {
                return Data_Foldable.foldl(foldableNonEmpty)(f)(b)(Data_Lazy.force(v));
            };
        };
    },
    foldMap: function (dictMonoid) {
        return function (f) {
            return function (v) {
                return Data_Foldable.foldMap(foldableNonEmpty)(dictMonoid)(f)(Data_Lazy.force(v));
            };
        };
    }
};
var showList = function (dictShow) {
    return {
        show: function (xs) {
            return "(fromFoldable [" + ((function () {
                var v = step(xs);
                if (v instanceof Nil) {
                    return "";
                };
                if (v instanceof Cons) {
                    return Data_Show.show(dictShow)(v.value0) + Data_Foldable.foldl(foldableList)(function (shown) {
                        return function (x$prime) {
                            return shown + ("," + Data_Show.show(dictShow)(x$prime));
                        };
                    })("")(v.value1);
                };
                throw new Error("Failed pattern match at Data.List.Lazy.Types (line 66, column 13 - line 69, column 78): " + [ v.constructor.name ]);
            })() + "])");
        }
    };
};
var showNonEmptyList = function (dictShow) {
    var showLazy = Data_Lazy.showLazy(Data_NonEmpty.showNonEmpty(dictShow)(showList(dictShow)));
    return {
        show: function (v) {
            return "(NonEmptyList " + (Data_Show.show(showLazy)(v) + ")");
        }
    };
};
var showStep = function (dictShow) {
    var showList1 = showList(dictShow);
    return {
        show: function (v) {
            if (v instanceof Nil) {
                return "Nil";
            };
            if (v instanceof Cons) {
                return "(" + (Data_Show.show(dictShow)(v.value0) + (" : " + (Data_Show.show(showList1)(v.value1) + ")")));
            };
            throw new Error("Failed pattern match at Data.List.Lazy.Types (line 36, column 1 - line 38, column 62): " + [ v.constructor.name ]);
        }
    };
};
var foldableWithIndexList = {
    foldrWithIndex: function (f) {
        return function (b) {
            return function (xs) {
                var v = (function () {
                    
                    // As we create our reversed list, we count elements.
var rev = Data_Foldable.foldl(foldableList)(function (v1) {
                        return function (a) {
                            return new Data_Tuple.Tuple(v1.value0 + 1 | 0, cons(a)(v1.value1));
                        };
                    });
                    return rev(new Data_Tuple.Tuple(0, nil))(xs);
                })();
                return Data_Tuple.snd(Data_Foldable.foldl(foldableList)(function (v1) {
                    return function (a) {
                        return new Data_Tuple.Tuple(v1.value0 - 1 | 0, f(v1.value0 - 1 | 0)(a)(v1.value1));
                    };
                })(new Data_Tuple.Tuple(v.value0, b))(v.value1));
            };
        };
    },
    foldlWithIndex: function (f) {
        return function (acc) {
            var $275 = Data_Foldable.foldl(foldableList)(function (v) {
                return function (a) {
                    return new Data_Tuple.Tuple(v.value0 + 1 | 0, f(v.value0)(v.value1)(a));
                };
            })(new Data_Tuple.Tuple(0, acc));
            return function ($276) {
                return Data_Tuple.snd($275($276));
            };
        };
    },
    foldMapWithIndex: function (dictMonoid) {
        var Semigroup0 = dictMonoid.Semigroup0();
        var mempty = Data_Monoid.mempty(dictMonoid);
        return function (f) {
            return Data_FoldableWithIndex.foldlWithIndex(foldableWithIndexList)(function (i) {
                return function (acc) {
                    var $277 = Data_Semigroup.append(Semigroup0)(acc);
                    var $278 = f(i);
                    return function ($279) {
                        return $277($278($279));
                    };
                };
            })(mempty);
        };
    },
    Foldable0: function () {
        return foldableList;
    }
};
var foldableWithIndexNonEmpty = /* #__PURE__ */ Data_NonEmpty.foldableWithIndexNonEmpty(foldableWithIndexList);
var foldableWithIndexNonEmptyList = {
    foldMapWithIndex: function (dictMonoid) {
        return function (f) {
            return function (v) {
                return Data_FoldableWithIndex.foldMapWithIndex(foldableWithIndexNonEmpty)(dictMonoid)((function () {
                    var $280 = Data_Maybe.maybe(0)(Data_Semiring.add(Data_Semiring.semiringInt)(1));
                    return function ($281) {
                        return f($280($281));
                    };
                })())(Data_Lazy.force(v));
            };
        };
    },
    foldlWithIndex: function (f) {
        return function (b) {
            return function (v) {
                return Data_FoldableWithIndex.foldlWithIndex(foldableWithIndexNonEmpty)((function () {
                    var $282 = Data_Maybe.maybe(0)(Data_Semiring.add(Data_Semiring.semiringInt)(1));
                    return function ($283) {
                        return f($282($283));
                    };
                })())(b)(Data_Lazy.force(v));
            };
        };
    },
    foldrWithIndex: function (f) {
        return function (b) {
            return function (v) {
                return Data_FoldableWithIndex.foldrWithIndex(foldableWithIndexNonEmpty)((function () {
                    var $284 = Data_Maybe.maybe(0)(Data_Semiring.add(Data_Semiring.semiringInt)(1));
                    return function ($285) {
                        return f($284($285));
                    };
                })())(b)(Data_Lazy.force(v));
            };
        };
    },
    Foldable0: function () {
        return foldableNonEmptyList;
    }
};
var functorWithIndexList = {
    mapWithIndex: function (f) {
        return Data_FoldableWithIndex.foldrWithIndex(foldableWithIndexList)(function (i) {
            return function (x) {
                return function (acc) {
                    return cons(f(i)(x))(acc);
                };
            };
        })(nil);
    },
    Functor0: function () {
        return functorList;
    }
};
var functorWithIndex = /* #__PURE__ */ Data_NonEmpty.functorWithIndex(functorWithIndexList);
var functorWithIndexNonEmptyList = {
    mapWithIndex: function (f) {
        return function (v) {
            return Data_Lazy.defer(function (v1) {
                return Data_FunctorWithIndex.mapWithIndex(functorWithIndex)((function () {
                    var $286 = Data_Maybe.maybe(0)(Data_Semiring.add(Data_Semiring.semiringInt)(1));
                    return function ($287) {
                        return f($286($287));
                    };
                })())(Data_Lazy.force(v));
            });
        };
    },
    Functor0: function () {
        return functorNonEmptyList;
    }
};
var toList = function (v) {
    return Control_Lazy.defer(lazyList)(function (v1) {
        var v2 = Data_Lazy.force(v);
        return cons(v2.value0)(v2.value1);
    });
};
var semigroupNonEmptyList = {
    append: function (v) {
        return function (as$prime) {
            var v1 = Data_Lazy.force(v);
            return Data_Lazy.defer(function (v2) {
                return new Data_NonEmpty.NonEmpty(v1.value0, Data_Semigroup.append(semigroupList)(v1.value1)(toList(as$prime)));
            });
        };
    }
};
var traversableList = {
    traverse: function (dictApplicative) {
        var Apply0 = dictApplicative.Apply0();
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (f) {
            return Data_Foldable.foldr(foldableList)(function (a) {
                return function (b) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(cons)(f(a)))(b);
                };
            })(Control_Applicative.pure(dictApplicative)(nil));
        };
    },
    sequence: function (dictApplicative) {
        return Data_Traversable.traverse(traversableList)(dictApplicative)(identity);
    },
    Functor0: function () {
        return functorList;
    },
    Foldable1: function () {
        return foldableList;
    }
};
var traversableNonEmpty = /* #__PURE__ */ Data_NonEmpty.traversableNonEmpty(traversableList);
var traversableNonEmptyList = {
    traverse: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (f) {
            return function (v) {
                return Data_Functor.map(Functor0)(function (xxs) {
                    return Data_Lazy.defer(function (v1) {
                        return xxs;
                    });
                })(Data_Traversable.traverse(traversableNonEmpty)(dictApplicative)(f)(Data_Lazy.force(v)));
            };
        };
    },
    sequence: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (v) {
            return Data_Functor.map(Functor0)(function (xxs) {
                return Data_Lazy.defer(function (v1) {
                    return xxs;
                });
            })(Data_Traversable.sequence(traversableNonEmpty)(dictApplicative)(Data_Lazy.force(v)));
        };
    },
    Functor0: function () {
        return functorNonEmptyList;
    },
    Foldable1: function () {
        return foldableNonEmptyList;
    }
};
var traversableWithIndexList = {
    traverseWithIndex: function (dictApplicative) {
        var Apply0 = dictApplicative.Apply0();
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (f) {
            return Data_FoldableWithIndex.foldrWithIndex(foldableWithIndexList)(function (i) {
                return function (a) {
                    return function (b) {
                        return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(cons)(f(i)(a)))(b);
                    };
                };
            })(Control_Applicative.pure(dictApplicative)(nil));
        };
    },
    FunctorWithIndex0: function () {
        return functorWithIndexList;
    },
    FoldableWithIndex1: function () {
        return foldableWithIndexList;
    },
    Traversable2: function () {
        return traversableList;
    }
};
var traversableWithIndexNonEmpty = /* #__PURE__ */ Data_NonEmpty.traversableWithIndexNonEmpty(traversableWithIndexList);
var traversableWithIndexNonEmptyList = {
    traverseWithIndex: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (f) {
            return function (v) {
                return Data_Functor.map(Functor0)(function (xxs) {
                    return Data_Lazy.defer(function (v1) {
                        return xxs;
                    });
                })(Data_TraversableWithIndex.traverseWithIndex(traversableWithIndexNonEmpty)(dictApplicative)((function () {
                    var $288 = Data_Maybe.maybe(0)(Data_Semiring.add(Data_Semiring.semiringInt)(1));
                    return function ($289) {
                        return f($288($289));
                    };
                })())(Data_Lazy.force(v)));
            };
        };
    },
    FunctorWithIndex0: function () {
        return functorWithIndexNonEmptyList;
    },
    FoldableWithIndex1: function () {
        return foldableWithIndexNonEmptyList;
    },
    Traversable2: function () {
        return traversableNonEmptyList;
    }
};
var unfoldable1List = {
    unfoldr1: /* #__PURE__ */ (function () {
        var go = function (f) {
            return function (b) {
                return Control_Lazy.defer(lazyList)(function (v) {
                    var v1 = f(b);
                    if (v1.value1 instanceof Data_Maybe.Just) {
                        return cons(v1.value0)(go(f)(v1.value1.value0));
                    };
                    if (v1.value1 instanceof Data_Maybe.Nothing) {
                        return cons(v1.value0)(nil);
                    };
                    throw new Error("Failed pattern match at Data.List.Lazy.Types (line 151, column 28 - line 153, column 33): " + [ v1.constructor.name ]);
                });
            };
        };
        return go;
    })()
};
var unfoldableList = {
    unfoldr: /* #__PURE__ */ (function () {
        var go = function (f) {
            return function (b) {
                return Control_Lazy.defer(lazyList)(function (v) {
                    var v1 = f(b);
                    if (v1 instanceof Data_Maybe.Nothing) {
                        return nil;
                    };
                    if (v1 instanceof Data_Maybe.Just) {
                        return cons(v1.value0.value0)(go(f)(v1.value0.value1));
                    };
                    throw new Error("Failed pattern match at Data.List.Lazy.Types (line 157, column 28 - line 159, column 39): " + [ v1.constructor.name ]);
                });
            };
        };
        return go;
    })(),
    Unfoldable10: function () {
        return unfoldable1List;
    }
};
var unfoldable1NonEmpty = /* #__PURE__ */ Data_NonEmpty.unfoldable1NonEmpty(unfoldableList);
var unfoldable1NonEmptyList = {
    unfoldr1: function (f) {
        return function (b) {
            return Data_Lazy.defer(function (v) {
                return Data_Unfoldable1.unfoldr1(unfoldable1NonEmpty)(f)(b);
            });
        };
    }
};
var comonadNonEmptyList = {
    extract: function (v) {
        return Data_NonEmpty.head(Data_Lazy.force(v));
    },
    Extend0: function () {
        return extendNonEmptyList;
    }
};
var monadList = {
    Applicative0: function () {
        return applicativeList;
    },
    Bind1: function () {
        return bindList;
    }
};
var bindList = {
    bind: function (xs) {
        return function (f) {
            var go = function (v) {
                if (v instanceof Nil) {
                    return Nil.value;
                };
                if (v instanceof Cons) {
                    return step(Data_Semigroup.append(semigroupList)(f(v.value0))(Control_Bind.bind(bindList)(v.value1)(f)));
                };
                throw new Error("Failed pattern match at Data.List.Lazy.Types (line 180, column 5 - line 180, column 17): " + [ v.constructor.name ]);
            };
            return Data_Functor.map(Data_Lazy.functorLazy)(go)(Data_Newtype.unwrap()(xs));
        };
    },
    Apply0: function () {
        return $lazy_applyList(0);
    }
};
var applicativeList = {
    pure: function (a) {
        return cons(a)(nil);
    },
    Apply0: function () {
        return $lazy_applyList(0);
    }
};
var $lazy_applyList = /* #__PURE__ */ $runtime_lazy("applyList", "Data.List.Lazy.Types", function () {
    return {
        apply: Control_Monad.ap(monadList),
        Functor0: function () {
            return functorList;
        }
    };
});
var applyList = /* #__PURE__ */ $lazy_applyList(171);
var applyNonEmptyList = {
    apply: function (v) {
        return function (v1) {
            var v2 = Data_Lazy.force(v1);
            var v3 = Data_Lazy.force(v);
            return Data_Lazy.defer(function (v4) {
                return new Data_NonEmpty.NonEmpty(v3.value0(v2.value0), Data_Semigroup.append(semigroupList)(Control_Apply.apply(applyList)(v3.value1)(cons(v2.value0)(nil)))(Control_Apply.apply(applyList)(cons(v3.value0)(v3.value1))(v2.value1)));
            });
        };
    },
    Functor0: function () {
        return functorNonEmptyList;
    }
};
var bindNonEmptyList = {
    bind: function (v) {
        return function (f) {
            var v1 = Data_Lazy.force(v);
            var v2 = Data_Lazy.force(unwrap(f(v1.value0)));
            return Data_Lazy.defer(function (v3) {
                return new Data_NonEmpty.NonEmpty(v2.value0, Data_Semigroup.append(semigroupList)(v2.value1)(Control_Bind.bind(bindList)(v1.value1)(function ($290) {
                    return toList(f($290));
                })));
            });
        };
    },
    Apply0: function () {
        return applyNonEmptyList;
    }
};
var altNonEmptyList = {
    alt: /* #__PURE__ */ Data_Semigroup.append(semigroupNonEmptyList),
    Functor0: function () {
        return functorNonEmptyList;
    }
};
var altList = {
    alt: /* #__PURE__ */ Data_Semigroup.append(semigroupList),
    Functor0: function () {
        return functorList;
    }
};
var plusList = {
    empty: nil,
    Alt0: function () {
        return altList;
    }
};
var alternativeList = {
    Applicative0: function () {
        return applicativeList;
    },
    Plus1: function () {
        return plusList;
    }
};
var monadPlusList = {
    Monad0: function () {
        return monadList;
    },
    Alternative1: function () {
        return alternativeList;
    }
};
var applicativeNonEmptyList = {
    pure: function (a) {
        return Data_Lazy.defer(function (v) {
            return Data_NonEmpty.singleton(plusList)(a);
        });
    },
    Apply0: function () {
        return applyNonEmptyList;
    }
};
var monadNonEmptyList = {
    Applicative0: function () {
        return applicativeNonEmptyList;
    },
    Bind1: function () {
        return bindNonEmptyList;
    }
};
export {
    List,
    Nil,
    Cons,
    step,
    nil,
    cons,
    NonEmptyList,
    toList,
    showStep,
    newtypeList,
    showList,
    eqList,
    eq1List,
    ordList,
    ord1List,
    lazyList,
    semigroupList,
    monoidList,
    functorList,
    functorWithIndexList,
    foldableList,
    foldableWithIndexList,
    unfoldable1List,
    unfoldableList,
    traversableList,
    traversableWithIndexList,
    applyList,
    applicativeList,
    bindList,
    monadList,
    altList,
    plusList,
    alternativeList,
    monadPlusList,
    extendList,
    newtypeNonEmptyList,
    eqNonEmptyList,
    ordNonEmptyList,
    eq1NonEmptyList,
    ord1NonEmptyList,
    showNonEmptyList,
    functorNonEmptyList,
    applyNonEmptyList,
    applicativeNonEmptyList,
    bindNonEmptyList,
    monadNonEmptyList,
    altNonEmptyList,
    extendNonEmptyList,
    comonadNonEmptyList,
    semigroupNonEmptyList,
    foldableNonEmptyList,
    traversableNonEmptyList,
    unfoldable1NonEmptyList,
    functorWithIndexNonEmptyList,
    foldableWithIndexNonEmptyList,
    traversableWithIndexNonEmptyList
};
