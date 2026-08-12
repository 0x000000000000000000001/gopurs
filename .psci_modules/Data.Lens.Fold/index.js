// | This module defines functions for working with folds.
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_HeytingAlgebra from "../Data.HeytingAlgebra/index.js";
import * as Data_Lens_Internal_Forget from "../Data.Lens.Internal.Forget/index.js";
import * as Data_Lens_Types from "../Data.Lens.Types/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Maybe_First from "../Data.Maybe.First/index.js";
import * as Data_Maybe_Last from "../Data.Maybe.Last/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Monoid_Additive from "../Data.Monoid.Additive/index.js";
import * as Data_Monoid_Conj from "../Data.Monoid.Conj/index.js";
import * as Data_Monoid_Disj from "../Data.Monoid.Disj/index.js";
import * as Data_Monoid_Dual from "../Data.Monoid.Dual/index.js";
import * as Data_Monoid_Endo from "../Data.Monoid.Endo/index.js";
import * as Data_Monoid_Multiplicative from "../Data.Monoid.Multiplicative/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Profunctor from "../Data.Profunctor/index.js";
import * as Data_Profunctor_Choice from "../Data.Profunctor.Choice/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
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
var unwrap = /* #__PURE__ */ Data_Newtype.unwrap();
var unwrap1 = /* #__PURE__ */ Data_Newtype.unwrap();
var unwrap2 = /* #__PURE__ */ Data_Newtype.unwrap();
var unwrap3 = /* #__PURE__ */ Data_Newtype.unwrap();
var unwrap4 = /* #__PURE__ */ Data_Newtype.unwrap();
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var fromFoldable = /* #__PURE__ */ Data_Array.fromFoldable(Data_List_Types.foldableList);
var unwrap5 = /* #__PURE__ */ Data_Newtype.unwrap();
var unwrap6 = /* #__PURE__ */ Data_Newtype.unwrap();
var unwrap7 = /* #__PURE__ */ Data_Newtype.unwrap();
var unwrap8 = /* #__PURE__ */ Data_Newtype.unwrap();

// | Builds a `Fold` using an unfold.
var unfolded = function (dictMonoid) {
    var Semigroup0 = dictMonoid.Semigroup0();
    return function (f) {
        return function (p) {
            var $lazy_go = $runtime_lazy("go", "Data.Lens.Fold", function () {
                var $69 = Data_Maybe.maybe(Data_Monoid.mempty(dictMonoid))(function (v) {
                    return Data_Semigroup.append(Semigroup0)(Data_Newtype.unwrap()(p)(v.value0))($lazy_go(232)(v.value1));
                });
                return function ($70) {
                    return $69(f($70));
                };
            });
            var go = $lazy_go(232);
            return go;
        };
    };
};

// | Replicates the elements of a fold.
var replicated = function (dictMonoid) {
    var monoidFn = Data_Monoid.monoidFn(dictMonoid);
    var semigroupFn = Data_Semigroup.semigroupFn(dictMonoid.Semigroup0());
    return function (i) {
        return function (v) {
            var go = function (v1) {
                return function (v2) {
                    if (v1 === 0) {
                        return Data_Monoid.mempty(monoidFn);
                    };
                    return Data_Semigroup.append(semigroupFn)(v2)(go(v1 - 1 | 0)(v2));
                };
            };
            return go(i)(v);
        };
    };
};

// | Fold map over an `IndexedFold`.
var ifoldMapOf = function (p) {
    return function (f) {
        return unwrap(p(Data_Tuple.uncurry(f)));
    };
};

// | Left fold over an `IndexedFold`.
var ifoldlOf = function (p) {
    return function (f) {
        return function (r) {
            var $71 = Data_Function.flip(unwrap1)(r);
            var $72 = ifoldMapOf(p)(function (i) {
                var $74 = Data_Function.flip(f(i));
                return function ($75) {
                    return Data_Monoid_Dual.Dual(Data_Monoid_Endo.Endo($74($75)));
                };
            });
            return function ($73) {
                return $71(unwrap2($72($73)));
            };
        };
    };
};

// | Right fold over an `IndexedFold`.
var ifoldrOf = function (p) {
    return function (f) {
        return function (r) {
            var $76 = Data_Function.flip(unwrap3)(r);
            var $77 = ifoldMapOf(p)(function (i) {
                var $79 = f(i);
                return function ($80) {
                    return Data_Monoid_Endo.Endo($79($80));
                };
            });
            return function ($78) {
                return $76($77($78));
            };
        };
    };
};

// | Collects the foci of an `IndexedFold` into a list.
var itoListOf = function (p) {
    return ifoldrOf(p)(function (i) {
        return function (x) {
            return function (xs) {
                return new Data_List_Types.Cons(new Data_Tuple.Tuple(i, x), xs);
            };
        };
    })(Data_List_Types.Nil.value);
};

// | Traverse the foci of an `IndexedFold`, discarding the results.
var itraverseOf_ = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (p) {
        return function (f) {
            return ifoldrOf(p)(function (i) {
                return function (a) {
                    return function (fu) {
                        return Control_Apply.applySecond(Apply0)(Data_Functor["void"](Functor0)(f(i)(a)))(fu);
                    };
                };
            })(Control_Applicative.pure(dictApplicative)(Data_Unit.unit));
        };
    };
};

// | Flipped version of `itraverseOf_`.
var iforOf_ = function (dictApplicative) {
    var $81 = itraverseOf_(dictApplicative);
    return function ($82) {
        return Data_Function.flip($81($82));
    };
};

// | Find the first focus of an `IndexedFold` that satisfies a predicate, if
// | there is any.
var ifindOf = function (p) {
    return function (f) {
        return ifoldrOf(p)(function (i) {
            return function (a) {
                return Data_Maybe.maybe((function () {
                    var $63 = f(i)(a);
                    if ($63) {
                        return new Data_Maybe.Just(a);
                    };
                    return Data_Maybe.Nothing.value;
                })())(Data_Maybe.Just.create);
            };
        })(Data_Maybe.Nothing.value);
    };
};

// | Whether any focus of an `IndexedFold` satisfies a predicate.
var ianyOf = function (dictHeytingAlgebra) {
    return function (p) {
        return function (f) {
            var $83 = ifoldMapOf(p)(function (i) {
                var $85 = f(i);
                return function ($86) {
                    return Data_Monoid_Disj.Disj($85($86));
                };
            });
            return function ($84) {
                return unwrap4($83($84));
            };
        };
    };
};

// | Whether all foci of an `IndexedFold` satisfy a predicate.
var iallOf = function (dictHeytingAlgebra) {
    return function (p) {
        return function (f) {
            var $87 = ifoldMapOf(p)(function (i) {
                var $89 = f(i);
                return function ($90) {
                    return Data_Monoid_Conj.Conj($89($90));
                };
            });
            return function ($88) {
                return unwrap4($87($88));
            };
        };
    };
};

// | Folds over a `Foldable` container.
var folded = function (dictMonoid) {
    return function (dictFoldable) {
        return function (v) {
            return Data_Foldable.foldMap(dictFoldable)(dictMonoid)(v);
        };
    };
};

// | Maps and then folds all foci of a `Fold`.
var foldMapOf = /* #__PURE__ */ Data_Newtype.under()()(Data_Lens_Internal_Forget.Forget);

// | Folds all foci of a `Fold` to one. Note that this is the same as `view`.
var foldOf = function (p) {
    return foldMapOf(p)(identity);
};

// | Left fold over a `Fold`.
var foldlOf = function (p) {
    return function (f) {
        return function (r) {
            var $91 = Data_Function.flip(unwrap1)(r);
            var $92 = foldMapOf(p)((function () {
                var $94 = Data_Function.flip(f);
                return function ($95) {
                    return Data_Monoid_Dual.Dual(Data_Monoid_Endo.Endo($94($95)));
                };
            })());
            return function ($93) {
                return $91(unwrap2($92($93)));
            };
        };
    };
};

// | Right fold over a `Fold`.
var foldrOf = function (p) {
    return function (f) {
        return function (r) {
            var $96 = Data_Function.flip(unwrap3)(r);
            var $97 = foldMapOf(p)(function ($99) {
                return Data_Monoid_Endo.Endo(f($99));
            });
            return function ($98) {
                return $96($97($98));
            };
        };
    };
};

// | The maximum of all foci of a `Fold`, if there is any.
var maximumOf = function (dictOrd) {
    return function (p) {
        var max = function (a) {
            return function (b) {
                var $65 = Data_Ord.greaterThan(dictOrd)(a)(b);
                if ($65) {
                    return a;
                };
                return b;
            };
        };
        return foldrOf(p)(function (a) {
            var $100 = Data_Maybe.maybe(a)(max(a));
            return function ($101) {
                return Data_Maybe.Just.create($100($101));
            };
        })(Data_Maybe.Nothing.value);
    };
};

// | The minimum of all foci of a `Fold`, if there is any.
var minimumOf = function (dictOrd) {
    return function (p) {
        var min = function (a) {
            return function (b) {
                var $66 = Data_Ord.lessThan(dictOrd)(a)(b);
                if ($66) {
                    return a;
                };
                return b;
            };
        };
        return foldrOf(p)(function (a) {
            var $102 = Data_Maybe.maybe(a)(min(a));
            return function ($103) {
                return Data_Maybe.Just.create($102($103));
            };
        })(Data_Maybe.Nothing.value);
    };
};

// | Collects the foci of a `Fold` into a list.
var toListOf = function (p) {
    return foldrOf(p)(Data_List_Types.Cons.create)(Data_List_Types.Nil.value);
};

// | Collects the foci of a `Fold` into an array.
var toArrayOf = function (p) {
    var $104 = toListOf(p);
    return function ($105) {
        return fromFoldable($104($105));
    };
};

// | Synonym for `toArrayOf`, reversed.
var toArrayOfOn = function (s) {
    return function (p) {
        return toArrayOf(p)(s);
    };
};

// | Synonym for `toListOf`, reversed.
var toListOfOn = function (s) {
    return function (p) {
        return toListOf(p)(s);
    };
};

// | Traverse the foci of a `Fold`, discarding the results.
var traverseOf_ = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (p) {
        return function (f) {
            return foldrOf(p)(function (a) {
                return function (fu) {
                    return Control_Apply.applySecond(Apply0)(Data_Functor["void"](Functor0)(f(a)))(fu);
                };
            })(Control_Applicative.pure(dictApplicative)(Data_Unit.unit));
        };
    };
};

// | Determines whether a `Fold` has at least one focus.
var has = function (dictHeytingAlgebra) {
    return function (p) {
        var $106 = foldMapOf(p)(Data_Function["const"](Data_HeytingAlgebra.tt(dictHeytingAlgebra)));
        return function ($107) {
            return unwrap4($106($107));
        };
    };
};

// | Determines whether a `Fold` does not have a focus.
var hasn$primet = function (dictHeytingAlgebra) {
    return function (p) {
        var $108 = foldMapOf(p)(Data_Function["const"](Data_HeytingAlgebra.ff(dictHeytingAlgebra)));
        return function ($109) {
            return unwrap4($108($109));
        };
    };
};

// | The last focus of a `Fold`, if there is any.
var lastOf = function (p) {
    var $110 = foldMapOf(p)(function ($112) {
        return Data_Maybe_Last.Last(Data_Maybe.Just.create($112));
    });
    return function ($111) {
        return unwrap5($110($111));
    };
};

// | The number of foci of a `Fold`.
var lengthOf = function (p) {
    var $113 = foldMapOf(p)(Data_Function["const"](1));
    return function ($114) {
        return unwrap6($113($114));
    };
};

// | Previews the first value of a fold, if there is any.
var preview = function (p) {
    var $115 = foldMapOf(p)(function ($117) {
        return Data_Maybe_First.First(Data_Maybe.Just.create($117));
    });
    return function ($116) {
        return unwrap5($115($116));
    };
};

// | Synonym for `preview`, flipped.
var previewOn = function (s) {
    return function (p) {
        return preview(p)(s);
    };
};

// | The product of all foci of a `Fold`.
var productOf = function (dictSemiring) {
    return function (p) {
        var $118 = foldMapOf(p)(Data_Monoid_Multiplicative.Multiplicative);
        return function ($119) {
            return unwrap7($118($119));
        };
    };
};

// | Sequence the foci of a `Fold`, pulling out an `Applicative`, and ignore
// | the result. If you need the result, see `sequenceOf` for `Traversal`s.
var sequenceOf_ = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    return function (p) {
        var $120 = Data_Function.flip(unwrap8)(Control_Applicative.pure(dictApplicative)(Data_Unit.unit));
        var $121 = foldMapOf(p)(function (f) {
            return function (v) {
                return Control_Apply.applySecond(Apply0)(f)(v);
            };
        });
        return function ($122) {
            return $120($121($122));
        };
    };
};

// | The sum of all foci of a `Fold`.
var sumOf = function (dictSemiring) {
    return function (p) {
        var $123 = foldMapOf(p)(Data_Monoid_Additive.Additive);
        return function ($124) {
            return unwrap7($123($124));
        };
    };
};

// | The first focus of a `Fold`, if there is any. Synonym for `preview`.
var firstOf = function (p) {
    var $125 = foldMapOf(p)(function ($127) {
        return Data_Maybe_First.First(Data_Maybe.Just.create($127));
    });
    return function ($126) {
        return unwrap5($125($126));
    };
};

// | Find the first focus of a `Fold` that satisfies a predicate, if there is any.
var findOf = function (p) {
    return function (f) {
        return foldrOf(p)(function (a) {
            return Data_Maybe.maybe((function () {
                var $67 = f(a);
                if ($67) {
                    return new Data_Maybe.Just(a);
                };
                return Data_Maybe.Nothing.value;
            })())(Data_Maybe.Just.create);
        })(Data_Maybe.Nothing.value);
    };
};

// | Filters on a predicate.
var filtered = function (dictChoice) {
    var right = Data_Profunctor_Choice.right(dictChoice);
    var Profunctor0 = dictChoice.Profunctor0();
    return function (f) {
        var $128 = Data_Profunctor.dimap(Profunctor0)(function (x) {
            var $68 = f(x);
            if ($68) {
                return new Data_Either.Right(x);
            };
            return new Data_Either.Left(x);
        })(Data_Either.either(identity)(identity));
        return function ($129) {
            return $128(right($129));
        };
    };
};

// | Whether any focus of a `Fold` satisfies a predicate.
var anyOf = function (dictHeytingAlgebra) {
    return function (p) {
        return function (f) {
            var $130 = foldMapOf(p)(function ($132) {
                return Data_Monoid_Disj.Disj(f($132));
            });
            return function ($131) {
                return unwrap4($130($131));
            };
        };
    };
};

// | Whether a `Fold` contains a given element.
var elemOf = function (dictEq) {
    return function (p) {
        return function (a) {
            return anyOf(Data_HeytingAlgebra.heytingAlgebraBoolean)(p)(function (v) {
                return Data_Eq.eq(dictEq)(v)(a);
            });
        };
    };
};

// | The disjunction of all foci of a `Fold`.
var orOf = function (dictHeytingAlgebra) {
    return function (p) {
        return anyOf(dictHeytingAlgebra)(p)(identity);
    };
};

// | Whether all foci of a `Fold` satisfy a predicate.
var allOf = function (dictHeytingAlgebra) {
    return function (p) {
        return function (f) {
            var $133 = foldMapOf(p)(function ($135) {
                return Data_Monoid_Conj.Conj(f($135));
            });
            return function ($134) {
                return unwrap4($133($134));
            };
        };
    };
};

// | The conjunction of all foci of a `Fold`.
var andOf = function (dictHeytingAlgebra) {
    return function (p) {
        return allOf(dictHeytingAlgebra)(p)(identity);
    };
};

// | Whether a `Fold` not contains a given element.
var notElemOf = function (dictEq) {
    return function (p) {
        return function (a) {
            return allOf(Data_HeytingAlgebra.heytingAlgebraBoolean)(p)(function (v) {
                return Data_Eq.notEq(dictEq)(v)(a);
            });
        };
    };
};
export {
    previewOn,
    toListOfOn,
    preview,
    foldOf,
    foldMapOf,
    foldrOf,
    foldlOf,
    toListOf,
    firstOf,
    lastOf,
    maximumOf,
    minimumOf,
    allOf,
    anyOf,
    andOf,
    orOf,
    elemOf,
    notElemOf,
    sumOf,
    productOf,
    lengthOf,
    findOf,
    sequenceOf_,
    traverseOf_,
    has,
    hasn$primet,
    replicated,
    filtered,
    folded,
    unfolded,
    toArrayOf,
    toArrayOfOn,
    ifoldMapOf,
    ifoldrOf,
    ifoldlOf,
    iallOf,
    ianyOf,
    ifindOf,
    itoListOf,
    itraverseOf_,
    iforOf_
};
