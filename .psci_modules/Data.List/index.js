// | This module defines a type of _strict_ linked lists, and associated helper
// | functions and type class instances.
// |
// | _Note_: Depending on your use-case, you may prefer to use
// | `Data.Sequence` instead, which might give better performance for certain
// | use cases. This module is an improvement over `Data.Array` when working with
// | immutable lists of data in a purely-functional setting, but does not have
// | good random-access performance.
import * as Control_Alt from "../Control.Alt/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Lazy from "../Control.Lazy/index.js";
import * as Control_Monad_Rec_Class from "../Control.Monad.Rec.Class/index.js";
import * as Data_Bifunctor from "../Data.Bifunctor/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_HeytingAlgebra from "../Data.HeytingAlgebra/index.js";
import * as Data_List_Internal from "../Data.List.Internal/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_NonEmpty from "../Data.NonEmpty/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);

//------------------------------------------------------------------------------
// Sublists --------------------------------------------------------------------
//------------------------------------------------------------------------------
// | A newtype used in cases where there is a list to be matched.
var Pattern = function (x) {
    return x;
};

// | Update the element at the specified index, returning a new
// | list or `Nothing` if the index is out-of-bounds.
// |
// | Running time: `O(n)`
var updateAt = function (v) {
    return function (v1) {
        return function (v2) {
            if (v === 0 && v2 instanceof Data_List_Types.Cons) {
                return new Data_Maybe.Just(new Data_List_Types.Cons(v1, v2.value1));
            };
            if (v2 instanceof Data_List_Types.Cons) {
                return Data_Functor.map(Data_Maybe.functorMaybe)(function (v3) {
                    return new Data_List_Types.Cons(v2.value0, v3);
                })(updateAt(v - 1 | 0)(v1)(v2.value1));
            };
            return Data_Maybe.Nothing.value;
        };
    };
};

// | Transforms a list of pairs into a list of first components and a list of
// | second components.
var unzip = /* #__PURE__ */ (function () {
    return Data_Foldable.foldr(Data_List_Types.foldableList)(function (v) {
        return function (v1) {
            return new Data_Tuple.Tuple(new Data_List_Types.Cons(v.value0, v1.value0), new Data_List_Types.Cons(v.value1, v1.value1));
        };
    })(new Data_Tuple.Tuple(Data_List_Types.Nil.value, Data_List_Types.Nil.value));
})();

// | Break a list into its first element, and the remaining elements,
// | or `Nothing` if the list is empty.
// |
// | Running time: `O(1)`
var uncons = function (v) {
    if (v instanceof Data_List_Types.Nil) {
        return Data_Maybe.Nothing.value;
    };
    if (v instanceof Data_List_Types.Cons) {
        return new Data_Maybe.Just({
            head: v.value0,
            tail: v.value1
        });
    };
    throw new Error("Failed pattern match at Data.List (line 259, column 1 - line 259, column 66): " + [ v.constructor.name ]);
};

// | Convert a list into any unfoldable structure.
// |
// | Running time: `O(n)`
var toUnfoldable = function (dictUnfoldable) {
    return Data_Unfoldable.unfoldr(dictUnfoldable)(function (xs) {
        return Data_Functor.map(Data_Maybe.functorMaybe)(function (rec) {
            return new Data_Tuple.Tuple(rec.head, rec.tail);
        })(uncons(xs));
    });
};

// | Get all but the first element of a list, or `Nothing` if the list is empty.
// |
// | Running time: `O(1)`
var tail = function (v) {
    if (v instanceof Data_List_Types.Nil) {
        return Data_Maybe.Nothing.value;
    };
    if (v instanceof Data_List_Types.Cons) {
        return new Data_Maybe.Just(v.value1);
    };
    throw new Error("Failed pattern match at Data.List (line 245, column 1 - line 245, column 43): " + [ v.constructor.name ]);
};

// | If the list starts with the given prefix, return the portion of the
// | list left after removing it, as a Just value. Otherwise, return Nothing.
// | * `stripPrefix (Pattern (1:Nil)) (1:2:Nil) == Just (2:Nil)`
// | * `stripPrefix (Pattern Nil) (1:Nil) == Just (1:Nil)`
// | * `stripPrefix (Pattern (2:Nil)) (1:Nil) == Nothing`
// |
// | Running time: `O(n)` where `n` is the number of elements to strip.
var stripPrefix = function (dictEq) {
    return function (v) {
        return function (s) {
            var go = function (prefix) {
                return function (input) {
                    if (prefix instanceof Data_List_Types.Cons && (input instanceof Data_List_Types.Cons && Data_Eq.eq(dictEq)(prefix.value0)(input.value0))) {
                        return new Data_Maybe.Just(new Control_Monad_Rec_Class.Loop({
                            a: prefix.value1,
                            b: input.value1
                        }));
                    };
                    if (prefix instanceof Data_List_Types.Nil) {
                        return new Data_Maybe.Just(new Control_Monad_Rec_Class.Done(input));
                    };
                    return Data_Maybe.Nothing.value;
                };
            };
            return Control_Monad_Rec_Class.tailRecM2(Control_Monad_Rec_Class.monadRecMaybe)(go)(v)(s);
        };
    };
};

// | Split a list into two parts:
// |
// | 1. the longest initial segment for which all elements satisfy the specified predicate
// | 2. the remaining elements
// |
// | For example,
// |
// | ```purescript
// | span (\n -> n % 2 == 1) (1 : 3 : 2 : 4 : 5 : Nil) == { init: (1 : 3 : Nil), rest: (2 : 4 : 5 : Nil) }
// | ```
// |
// | Running time: `O(n)`
var span = function (v) {
    return function (v1) {
        if (v1 instanceof Data_List_Types.Cons && v(v1.value0)) {
            var v2 = span(v)(v1.value1);
            return {
                init: new Data_List_Types.Cons(v1.value0, v2.init),
                rest: v2.rest
            };
        };
        return {
            init: Data_List_Types.Nil.value,
            rest: v1
        };
    };
};

//------------------------------------------------------------------------------
// Extending lists -------------------------------------------------------------
//------------------------------------------------------------------------------
// | Append an element to the end of a list, creating a new list.
// |
// | Running time: `O(n)`
var snoc = function (xs) {
    return function (x) {
        return Data_Foldable.foldr(Data_List_Types.foldableList)(Data_List_Types.Cons.create)(new Data_List_Types.Cons(x, Data_List_Types.Nil.value))(xs);
    };
};

//------------------------------------------------------------------------------
// List creation ---------------------------------------------------------------
//------------------------------------------------------------------------------
// | Create a list with a single element.
// |
// | Running time: `O(1)`
var singleton = function (a) {
    return new Data_List_Types.Cons(a, Data_List_Types.Nil.value);
};

// | Sort the elements of a list in increasing order, where elements are
// | compared using the specified ordering.
var sortBy = function (cmp) {
    var merge = function (v) {
        return function (v1) {
            if (v instanceof Data_List_Types.Cons && v1 instanceof Data_List_Types.Cons) {
                if (Data_Eq.eq(Data_Ordering.eqOrdering)(cmp(v.value0)(v1.value0))(Data_Ordering.GT.value)) {
                    return new Data_List_Types.Cons(v1.value0, merge(v)(v1.value1));
                };
                if (Data_Boolean.otherwise) {
                    return new Data_List_Types.Cons(v.value0, merge(v.value1)(v1));
                };
            };
            if (v instanceof Data_List_Types.Nil) {
                return v1;
            };
            if (v1 instanceof Data_List_Types.Nil) {
                return v;
            };
            throw new Error("Failed pattern match at Data.List (line 466, column 3 - line 466, column 38): " + [ v.constructor.name, v1.constructor.name ]);
        };
    };
    var mergePairs = function (v) {
        if (v instanceof Data_List_Types.Cons && v.value1 instanceof Data_List_Types.Cons) {
            return new Data_List_Types.Cons(merge(v.value0)(v.value1.value0), mergePairs(v.value1.value1));
        };
        return v;
    };
    var mergeAll = function ($copy_v) {
        var $tco_done = false;
        var $tco_result;
        function $tco_loop(v) {
            if (v instanceof Data_List_Types.Cons && v.value1 instanceof Data_List_Types.Nil) {
                $tco_done = true;
                return v.value0;
            };
            $copy_v = mergePairs(v);
            return;
        };
        while (!$tco_done) {
            $tco_result = $tco_loop($copy_v);
        };
        return $tco_result;
    };
    var sequences = function (v) {
        if (v instanceof Data_List_Types.Cons && v.value1 instanceof Data_List_Types.Cons) {
            if (Data_Eq.eq(Data_Ordering.eqOrdering)(cmp(v.value0)(v.value1.value0))(Data_Ordering.GT.value)) {
                return descending(v.value1.value0)(singleton(v.value0))(v.value1.value1);
            };
            if (Data_Boolean.otherwise) {
                return ascending(v.value1.value0)(function (v1) {
                    return new Data_List_Types.Cons(v.value0, v1);
                })(v.value1.value1);
            };
        };
        return singleton(v);
    };
    var descending = function ($copy_v) {
        return function ($copy_v1) {
            return function ($copy_v2) {
                var $tco_var_v = $copy_v;
                var $tco_var_v1 = $copy_v1;
                var $tco_done1 = false;
                var $tco_result;
                function $tco_loop(v, v1, v2) {
                    if (v2 instanceof Data_List_Types.Cons && Data_Eq.eq(Data_Ordering.eqOrdering)(cmp(v)(v2.value0))(Data_Ordering.GT.value)) {
                        $tco_var_v = v2.value0;
                        $tco_var_v1 = new Data_List_Types.Cons(v, v1);
                        $copy_v2 = v2.value1;
                        return;
                    };
                    $tco_done1 = true;
                    return new Data_List_Types.Cons(new Data_List_Types.Cons(v, v1), sequences(v2));
                };
                while (!$tco_done1) {
                    $tco_result = $tco_loop($tco_var_v, $tco_var_v1, $copy_v2);
                };
                return $tco_result;
            };
        };
    };
    var ascending = function ($copy_v) {
        return function ($copy_v1) {
            return function ($copy_v2) {
                var $tco_var_v = $copy_v;
                var $tco_var_v1 = $copy_v1;
                var $tco_done2 = false;
                var $tco_result;
                function $tco_loop(v, v1, v2) {
                    if (v2 instanceof Data_List_Types.Cons && Data_Eq.notEq(Data_Ordering.eqOrdering)(cmp(v)(v2.value0))(Data_Ordering.GT.value)) {
                        $tco_var_v = v2.value0;
                        $tco_var_v1 = function (ys) {
                            return v1(new Data_List_Types.Cons(v, ys));
                        };
                        $copy_v2 = v2.value1;
                        return;
                    };
                    $tco_done2 = true;
                    return new Data_List_Types.Cons(v1(singleton(v)), sequences(v2));
                };
                while (!$tco_done2) {
                    $tco_result = $tco_loop($tco_var_v, $tco_var_v1, $copy_v2);
                };
                return $tco_result;
            };
        };
    };
    return function ($401) {
        return mergeAll(sequences($401));
    };
};

//------------------------------------------------------------------------------
// Sorting ---------------------------------------------------------------------
//------------------------------------------------------------------------------
// | Sort the elements of an list in increasing order.
var sort = function (dictOrd) {
    var compare = Data_Ord.compare(dictOrd);
    return function (xs) {
        return sortBy(compare)(xs);
    };
};

// | Returns all final segments of the argument, longest first. For example,
// |
// | ```purescript
// | tails (1 : 2 : 3 : Nil) == ((1 : 2 : 3 : Nil) : (2 : 3 : Nil) : (3 : Nil) : (Nil) : Nil)
// | ```
// | Running time: `O(n)`
var tails = function (v) {
    if (v instanceof Data_List_Types.Nil) {
        return singleton(Data_List_Types.Nil.value);
    };
    if (v instanceof Data_List_Types.Cons) {
        return new Data_List_Types.Cons(v, tails(v.value1));
    };
    throw new Error("Failed pattern match at Data.List (line 641, column 1 - line 641, column 43): " + [ v.constructor.name ]);
};
var showPattern = function (dictShow) {
    var showList = Data_List_Types.showList(dictShow);
    return {
        show: function (v) {
            return "(Pattern " + (Data_Show.show(showList)(v) + ")");
        }
    };
};

//------------------------------------------------------------------------------
// Transformations -------------------------------------------------------------
//------------------------------------------------------------------------------
// | Reverse a list.
// |
// | Running time: `O(n)`
var reverse = /* #__PURE__ */ (function () {
    var go = function ($copy_v) {
        return function ($copy_v1) {
            var $tco_var_v = $copy_v;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(v, v1) {
                if (v1 instanceof Data_List_Types.Nil) {
                    $tco_done = true;
                    return v;
                };
                if (v1 instanceof Data_List_Types.Cons) {
                    $tco_var_v = new Data_List_Types.Cons(v1.value0, v);
                    $copy_v1 = v1.value1;
                    return;
                };
                throw new Error("Failed pattern match at Data.List (line 368, column 3 - line 368, column 19): " + [ v.constructor.name, v1.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_v, $copy_v1);
            };
            return $tco_result;
        };
    };
    return go(Data_List_Types.Nil.value);
})();

// | Take the specified number of elements from the front of a list.
// |
// | Running time: `O(n)` where `n` is the number of elements to take.
var take = /* #__PURE__ */ (function () {
    var go = function ($copy_v) {
        return function ($copy_v1) {
            return function ($copy_v2) {
                var $tco_var_v = $copy_v;
                var $tco_var_v1 = $copy_v1;
                var $tco_done = false;
                var $tco_result;
                function $tco_loop(v, v1, v2) {
                    if (v1 < 1) {
                        $tco_done = true;
                        return reverse(v);
                    };
                    if (v2 instanceof Data_List_Types.Nil) {
                        $tco_done = true;
                        return reverse(v);
                    };
                    if (v2 instanceof Data_List_Types.Cons) {
                        $tco_var_v = new Data_List_Types.Cons(v2.value0, v);
                        $tco_var_v1 = v1 - 1 | 0;
                        $copy_v2 = v2.value1;
                        return;
                    };
                    throw new Error("Failed pattern match at Data.List (line 513, column 3 - line 513, column 35): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
                };
                while (!$tco_done) {
                    $tco_result = $tco_loop($tco_var_v, $tco_var_v1, $copy_v2);
                };
                return $tco_result;
            };
        };
    };
    return go(Data_List_Types.Nil.value);
})();

// | Take those elements from the front of a list which match a predicate.
// |
// | Running time (worst case): `O(n)`
var takeWhile = function (p) {
    var go = function ($copy_v) {
        return function ($copy_v1) {
            var $tco_var_v = $copy_v;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(v, v1) {
                if (v1 instanceof Data_List_Types.Cons && p(v1.value0)) {
                    $tco_var_v = new Data_List_Types.Cons(v1.value0, v);
                    $copy_v1 = v1.value1;
                    return;
                };
                $tco_done = true;
                return reverse(v);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_v, $copy_v1);
            };
            return $tco_result;
        };
    };
    return go(Data_List_Types.Nil.value);
};

// | Break a list into its last element, and the preceding elements,
// | or `Nothing` if the list is empty.
// |
// | Running time: `O(n)`
var unsnoc = function (lst) {
    var go = function ($copy_v) {
        return function ($copy_v1) {
            var $tco_var_v = $copy_v;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(v, v1) {
                if (v instanceof Data_List_Types.Nil) {
                    $tco_done = true;
                    return Data_Maybe.Nothing.value;
                };
                if (v instanceof Data_List_Types.Cons && v.value1 instanceof Data_List_Types.Nil) {
                    $tco_done = true;
                    return new Data_Maybe.Just({
                        revInit: v1,
                        last: v.value0
                    });
                };
                if (v instanceof Data_List_Types.Cons) {
                    $tco_var_v = v.value1;
                    $copy_v1 = new Data_List_Types.Cons(v.value0, v1);
                    return;
                };
                throw new Error("Failed pattern match at Data.List (line 270, column 3 - line 270, column 21): " + [ v.constructor.name, v1.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_v, $copy_v1);
            };
            return $tco_result;
        };
    };
    return Data_Functor.map(Data_Maybe.functorMaybe)(function (h) {
        return {
            init: reverse(h.revInit),
            last: h.last
        };
    })(go(lst)(Data_List_Types.Nil.value));
};

//------------------------------------------------------------------------------
// Zipping ---------------------------------------------------------------------
//------------------------------------------------------------------------------
// | Apply a function to pairs of elements at the same positions in two lists,
// | collecting the results in a new list.
// |
// | If one list is longer, elements will be discarded from the longer list.
// |
// | For example
// |
// | ```purescript
// | zipWith (*) (1 : 2 : 3 : Nil) (4 : 5 : 6 : 7 Nil) == 4 : 10 : 18 : Nil
// | ```
// |
// | Running time: `O(min(m, n))`
var zipWith = function (f) {
    return function (xs) {
        return function (ys) {
            var go = function ($copy_v) {
                return function ($copy_v1) {
                    return function ($copy_v2) {
                        var $tco_var_v = $copy_v;
                        var $tco_var_v1 = $copy_v1;
                        var $tco_done = false;
                        var $tco_result;
                        function $tco_loop(v, v1, v2) {
                            if (v instanceof Data_List_Types.Nil) {
                                $tco_done = true;
                                return v2;
                            };
                            if (v1 instanceof Data_List_Types.Nil) {
                                $tco_done = true;
                                return v2;
                            };
                            if (v instanceof Data_List_Types.Cons && v1 instanceof Data_List_Types.Cons) {
                                $tco_var_v = v.value1;
                                $tco_var_v1 = v1.value1;
                                $copy_v2 = new Data_List_Types.Cons(f(v.value0)(v1.value0), v2);
                                return;
                            };
                            throw new Error("Failed pattern match at Data.List (line 779, column 3 - line 779, column 21): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
                        };
                        while (!$tco_done) {
                            $tco_result = $tco_loop($tco_var_v, $tco_var_v1, $copy_v2);
                        };
                        return $tco_result;
                    };
                };
            };
            return reverse(go(xs)(ys)(Data_List_Types.Nil.value));
        };
    };
};

// | Collect pairs of elements at the same positions in two lists.
// |
// | Running time: `O(min(m, n))`
var zip = /* #__PURE__ */ (function () {
    return zipWith(Data_Tuple.Tuple.create);
})();

// | A generalization of `zipWith` which accumulates results in some `Applicative`
// | functor.
var zipWithA = function (dictApplicative) {
    return function (f) {
        return function (xs) {
            return function (ys) {
                return Data_Traversable.sequence(Data_List_Types.traversableList)(dictApplicative)(zipWith(f)(xs)(ys));
            };
        };
    };
};

// | Create a list containing a range of integers, including both endpoints.
var range = function (start) {
    return function (end) {
        if (start === end) {
            return singleton(start);
        };
        if (Data_Boolean.otherwise) {
            var go = function ($copy_s) {
                return function ($copy_e) {
                    return function ($copy_step) {
                        return function ($copy_rest) {
                            var $tco_var_s = $copy_s;
                            var $tco_var_e = $copy_e;
                            var $tco_var_step = $copy_step;
                            var $tco_done = false;
                            var $tco_result;
                            function $tco_loop(s, e, step, rest) {
                                if (s === e) {
                                    $tco_done = true;
                                    return new Data_List_Types.Cons(s, rest);
                                };
                                if (Data_Boolean.otherwise) {
                                    $tco_var_s = s + step | 0;
                                    $tco_var_e = e;
                                    $tco_var_step = step;
                                    $copy_rest = new Data_List_Types.Cons(s, rest);
                                    return;
                                };
                                throw new Error("Failed pattern match at Data.List (line 148, column 3 - line 149, column 65): " + [ s.constructor.name, e.constructor.name, step.constructor.name, rest.constructor.name ]);
                            };
                            while (!$tco_done) {
                                $tco_result = $tco_loop($tco_var_s, $tco_var_e, $tco_var_step, $copy_rest);
                            };
                            return $tco_result;
                        };
                    };
                };
            };
            return go(end)(start)((function () {
                var $282 = start > end;
                if ($282) {
                    return 1;
                };
                return -1 | 0;
            })())(Data_List_Types.Nil.value);
        };
        throw new Error("Failed pattern match at Data.List (line 144, column 1 - line 144, column 32): " + [ start.constructor.name, end.constructor.name ]);
    };
};

// | Returns a lists of elements which do and do not satisfy a predicate.
// |
// | Running time: `O(n)`
var partition = function (p) {
    return function (xs) {
        var select = function (x) {
            return function (v) {
                var $285 = p(x);
                if ($285) {
                    return {
                        no: v.no,
                        yes: new Data_List_Types.Cons(x, v.yes)
                    };
                };
                return {
                    no: new Data_List_Types.Cons(x, v.no),
                    yes: v.yes
                };
            };
        };
        return Data_Foldable.foldr(Data_List_Types.foldableList)(select)({
            no: Data_List_Types.Nil.value,
            yes: Data_List_Types.Nil.value
        })(xs);
    };
};

//------------------------------------------------------------------------------
// List size -------------------------------------------------------------------
//------------------------------------------------------------------------------
// | Test whether a list is empty.
// |
// | Running time: `O(1)`
var $$null = function (v) {
    if (v instanceof Data_List_Types.Nil) {
        return true;
    };
    return false;
};

// | Remove duplicate elements from a list based on the provided comparison function.
// | Keeps the first occurrence of each element in the input list,
// | in the same order they appear in the input list.
// |
// | ```purescript
// | nubBy (compare `on` Array.length) ([1]:[2]:[3,4]:Nil) == [1]:[3,4]:Nil
// | ```
// |
// | Running time: `O(n log n)`
var nubBy = function (p) {
    var go = function ($copy_v) {
        return function ($copy_v1) {
            return function ($copy_v2) {
                var $tco_var_v = $copy_v;
                var $tco_var_v1 = $copy_v1;
                var $tco_done = false;
                var $tco_result;
                function $tco_loop(v, v1, v2) {
                    if (v2 instanceof Data_List_Types.Nil) {
                        $tco_done = true;
                        return v1;
                    };
                    if (v2 instanceof Data_List_Types.Cons) {
                        var v3 = Data_List_Internal.insertAndLookupBy(p)(v2.value0)(v);
                        if (v3.found) {
                            $tco_var_v = v3.result;
                            $tco_var_v1 = v1;
                            $copy_v2 = v2.value1;
                            return;
                        };
                        $tco_var_v = v3.result;
                        $tco_var_v1 = new Data_List_Types.Cons(v2.value0, v1);
                        $copy_v2 = v2.value1;
                        return;
                    };
                    throw new Error("Failed pattern match at Data.List (line 673, column 5 - line 673, column 23): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
                };
                while (!$tco_done) {
                    $tco_result = $tco_loop($tco_var_v, $tco_var_v1, $copy_v2);
                };
                return $tco_result;
            };
        };
    };
    var $402 = go(Data_List_Internal.emptySet)(Data_List_Types.Nil.value);
    return function ($403) {
        return reverse($402($403));
    };
};

//------------------------------------------------------------------------------
// Set-like operations ---------------------------------------------------------
//------------------------------------------------------------------------------
// | Remove duplicate elements from a list.
// | Keeps the first occurrence of each element in the input list,
// | in the same order they appear in the input list.
// |
// | ```purescript
// | nub 1:2:1:3:3:Nil == 1:2:3:Nil
// | ```
// |
// | Running time: `O(n log n)`
var nub = function (dictOrd) {
    return nubBy(Data_Ord.compare(dictOrd));
};
var newtypePattern = {
    Coercible0: function () {
        return undefined;
    }
};

// | Apply a function to each element in a list, keeping only the results which
// | contain a value.
// |
// | Running time: `O(n)`
var mapMaybe = function (f) {
    var go = function ($copy_v) {
        return function ($copy_v1) {
            var $tco_var_v = $copy_v;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(v, v1) {
                if (v1 instanceof Data_List_Types.Nil) {
                    $tco_done = true;
                    return reverse(v);
                };
                if (v1 instanceof Data_List_Types.Cons) {
                    var v2 = f(v1.value0);
                    if (v2 instanceof Data_Maybe.Nothing) {
                        $tco_var_v = v;
                        $copy_v1 = v1.value1;
                        return;
                    };
                    if (v2 instanceof Data_Maybe.Just) {
                        $tco_var_v = new Data_List_Types.Cons(v2.value0, v);
                        $copy_v1 = v1.value1;
                        return;
                    };
                    throw new Error("Failed pattern match at Data.List (line 419, column 5 - line 421, column 32): " + [ v2.constructor.name ]);
                };
                throw new Error("Failed pattern match at Data.List (line 417, column 3 - line 417, column 27): " + [ v.constructor.name, v1.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_v, $copy_v1);
            };
            return $tco_result;
        };
    };
    return go(Data_List_Types.Nil.value);
};

// | A stack-safe version of `many`, at the cost of a `MonadRec` constraint.
var manyRec = function (dictMonadRec) {
    var Bind1 = (dictMonadRec.Monad0()).Bind1();
    return function (dictAlternative) {
        var Plus1 = dictAlternative.Plus1();
        var Alt0 = Plus1.Alt0();
        var Functor0 = (Plus1.Alt0()).Functor0();
        var Applicative0 = dictAlternative.Applicative0();
        var pure = Control_Applicative.pure(dictAlternative.Applicative0());
        return function (p) {
            var go = function (acc) {
                return Control_Bind.bind(Bind1)(Control_Alt.alt(Alt0)(Data_Functor.map(Functor0)(Control_Monad_Rec_Class.Loop.create)(p))(Control_Applicative.pure(Applicative0)(new Control_Monad_Rec_Class.Done(Data_Unit.unit))))(function (aa) {
                    return pure(Data_Bifunctor.bimap(Control_Monad_Rec_Class.bifunctorStep)(function (v) {
                        return new Data_List_Types.Cons(v, acc);
                    })(function (v) {
                        return reverse(acc);
                    })(aa));
                });
            };
            return Control_Monad_Rec_Class.tailRecM(dictMonadRec)(go)(Data_List_Types.Nil.value);
        };
    };
};

// | A stack-safe version of `some`, at the cost of a `MonadRec` constraint.
var someRec = function (dictMonadRec) {
    return function (dictAlternative) {
        var Apply0 = (dictAlternative.Applicative0()).Apply0();
        var Functor0 = ((dictAlternative.Plus1()).Alt0()).Functor0();
        return function (v) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Data_List_Types.Cons.create)(v))(manyRec(dictMonadRec)(dictAlternative)(v));
        };
    };
};

// | Attempt a computation multiple times, requiring at least one success.
// |
// | The `Lazy` constraint is used to generate the result lazily, to ensure
// | termination.
var some = function (dictAlternative) {
    var Apply0 = (dictAlternative.Applicative0()).Apply0();
    var Functor0 = ((dictAlternative.Plus1()).Alt0()).Functor0();
    return function (dictLazy) {
        return function (v) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Data_List_Types.Cons.create)(v))(Control_Lazy.defer(dictLazy)(function (v1) {
                return many(dictAlternative)(dictLazy)(v);
            }));
        };
    };
};

// | Attempt a computation multiple times, returning as many successful results
// | as possible (possibly zero).
// |
// | The `Lazy` constraint is used to generate the result lazily, to ensure
// | termination.
var many = function (dictAlternative) {
    var Alt0 = (dictAlternative.Plus1()).Alt0();
    var Applicative0 = dictAlternative.Applicative0();
    return function (dictLazy) {
        return function (v) {
            return Control_Alt.alt(Alt0)(some(dictAlternative)(dictLazy)(v))(Control_Applicative.pure(Applicative0)(Data_List_Types.Nil.value));
        };
    };
};

// | Get the length of a list
// |
// | Running time: `O(n)`
var length = /* #__PURE__ */ Data_Foldable.foldl(Data_List_Types.foldableList)(function (acc) {
    return function (v) {
        return acc + 1 | 0;
    };
})(0);

// | Get the last element in a list, or `Nothing` if the list is empty.
// |
// | Running time: `O(n)`.
var last = function ($copy_v) {
    var $tco_done = false;
    var $tco_result;
    function $tco_loop(v) {
        if (v instanceof Data_List_Types.Cons && v.value1 instanceof Data_List_Types.Nil) {
            $tco_done = true;
            return new Data_Maybe.Just(v.value0);
        };
        if (v instanceof Data_List_Types.Cons) {
            $copy_v = v.value1;
            return;
        };
        $tco_done = true;
        return Data_Maybe.Nothing.value;
    };
    while (!$tco_done) {
        $tco_result = $tco_loop($copy_v);
    };
    return $tco_result;
};

// | Insert an element into a sorted list, using the specified function to
// | determine the ordering of elements.
// |
// | Running time: `O(n)`
var insertBy = function (v) {
    return function (v1) {
        return function (v2) {
            if (v2 instanceof Data_List_Types.Nil) {
                return singleton(v1);
            };
            if (v2 instanceof Data_List_Types.Cons) {
                var v3 = v(v1)(v2.value0);
                if (v3 instanceof Data_Ordering.GT) {
                    return new Data_List_Types.Cons(v2.value0, insertBy(v)(v1)(v2.value1));
                };
                return new Data_List_Types.Cons(v1, v2);
            };
            throw new Error("Failed pattern match at Data.List (line 216, column 1 - line 216, column 68): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
        };
    };
};

// | Insert an element into a list at the specified index, returning a new
// | list or `Nothing` if the index is out-of-bounds.
// |
// | Running time: `O(n)`
var insertAt = function (v) {
    return function (v1) {
        return function (v2) {
            if (v === 0) {
                return new Data_Maybe.Just(new Data_List_Types.Cons(v1, v2));
            };
            if (v2 instanceof Data_List_Types.Cons) {
                return Data_Functor.map(Data_Maybe.functorMaybe)(function (v3) {
                    return new Data_List_Types.Cons(v2.value0, v3);
                })(insertAt(v - 1 | 0)(v1)(v2.value1));
            };
            return Data_Maybe.Nothing.value;
        };
    };
};

// | Insert an element into a sorted list.
// |
// | Running time: `O(n)`
var insert = function (dictOrd) {
    return insertBy(Data_Ord.compare(dictOrd));
};

// | Get all but the last element of a list, or `Nothing` if the list is empty.
// |
// | Running time: `O(n)`
var init = function (lst) {
    return Data_Functor.map(Data_Maybe.functorMaybe)(function (v) {
        return v.init;
    })(unsnoc(lst));
};

//------------------------------------------------------------------------------
// Indexed operations ----------------------------------------------------------
//------------------------------------------------------------------------------
// | Get the element at the specified index, or `Nothing` if the index is out-of-bounds.
// |
// | Running time: `O(n)` where `n` is the required index.
var index = function ($copy_v) {
    return function ($copy_v1) {
        var $tco_var_v = $copy_v;
        var $tco_done = false;
        var $tco_result;
        function $tco_loop(v, v1) {
            if (v instanceof Data_List_Types.Nil) {
                $tco_done = true;
                return Data_Maybe.Nothing.value;
            };
            if (v instanceof Data_List_Types.Cons && v1 === 0) {
                $tco_done = true;
                return new Data_Maybe.Just(v.value0);
            };
            if (v instanceof Data_List_Types.Cons) {
                $tco_var_v = v.value1;
                $copy_v1 = v1 - 1 | 0;
                return;
            };
            throw new Error("Failed pattern match at Data.List (line 281, column 1 - line 281, column 44): " + [ v.constructor.name, v1.constructor.name ]);
        };
        while (!$tco_done) {
            $tco_result = $tco_loop($tco_var_v, $copy_v1);
        };
        return $tco_result;
    };
};

//------------------------------------------------------------------------------
// Non-indexed reads -----------------------------------------------------------
//------------------------------------------------------------------------------
// | Get the first element in a list, or `Nothing` if the list is empty.
// |
// | Running time: `O(1)`.
var head = function (v) {
    if (v instanceof Data_List_Types.Nil) {
        return Data_Maybe.Nothing.value;
    };
    if (v instanceof Data_List_Types.Cons) {
        return new Data_Maybe.Just(v.value0);
    };
    throw new Error("Failed pattern match at Data.List (line 230, column 1 - line 230, column 22): " + [ v.constructor.name ]);
};

//------------------------------------------------------------------------------
// Transpose -------------------------------------------------------------------
//------------------------------------------------------------------------------
// | The 'transpose' function transposes the rows and columns of its argument.
// | For example,
// |
// |     transpose ((1:2:3:Nil) : (4:5:6:Nil) : Nil) ==
// |       ((1:4:Nil) : (2:5:Nil) : (3:6:Nil) : Nil)
// |
// | If some of the rows are shorter than the following rows, their elements are skipped:
// |
// |     transpose ((10:11:Nil) : (20:Nil) : Nil : (30:31:32:Nil) : Nil) ==
// |       ((10:20:30:Nil) : (11:31:Nil) : (32:Nil) : Nil)
var transpose = function (v) {
    if (v instanceof Data_List_Types.Nil) {
        return Data_List_Types.Nil.value;
    };
    if (v instanceof Data_List_Types.Cons && v.value0 instanceof Data_List_Types.Nil) {
        return transpose(v.value1);
    };
    if (v instanceof Data_List_Types.Cons && v.value0 instanceof Data_List_Types.Cons) {
        return new Data_List_Types.Cons(new Data_List_Types.Cons(v.value0.value0, mapMaybe(head)(v.value1)), transpose(new Data_List_Types.Cons(v.value0.value1, mapMaybe(tail)(v.value1))));
    };
    throw new Error("Failed pattern match at Data.List (line 813, column 1 - line 813, column 54): " + [ v.constructor.name ]);
};

// | Group equal, consecutive elements of a list into lists, using the specified
// | equivalence relation to determine equality.
// |
// | For example,
// |
// | ```purescript
// | groupBy (\a b -> odd a && odd b) (1 : 3 : 2 : 4 : 3 : 3 : Nil) ==
// |   (NonEmptyList (NonEmpty 1 (3 : Nil))) : (NonEmptyList (NonEmpty 2 Nil)) : (NonEmptyList (NonEmpty 4 Nil)) : (NonEmptyList (NonEmpty 3 (3 : Nil))) : Nil
// | ```
// |
// | Running time: `O(n)`
var groupBy = function (v) {
    return function (v1) {
        if (v1 instanceof Data_List_Types.Nil) {
            return Data_List_Types.Nil.value;
        };
        if (v1 instanceof Data_List_Types.Cons) {
            var v2 = span(v(v1.value0))(v1.value1);
            return new Data_List_Types.Cons(new Data_NonEmpty.NonEmpty(v1.value0, v2.init), groupBy(v)(v2.rest));
        };
        throw new Error("Failed pattern match at Data.List (line 609, column 1 - line 609, column 80): " + [ v.constructor.name, v1.constructor.name ]);
    };
};

// | Sort, then group equal elements of a list into lists, using the provided comparison function.
// |
// | ```purescript
// | groupAllBy (compare `on` (_ `div` 10)) (32 : 31 : 21 : 22 : 11 : 33 : Nil) ==
// |   NonEmptyList (11 :| Nil) : NonEmptyList (21 :| 22 : Nil) : NonEmptyList (32 :| 31 : 33) : Nil
// | ```
// |
// | Running time: `O(n log n)`
var groupAllBy = function (p) {
    var $404 = groupBy(function (x) {
        return function (y) {
            return Data_Eq.eq(Data_Ordering.eqOrdering)(p(x)(y))(Data_Ordering.EQ.value);
        };
    });
    var $405 = sortBy(p);
    return function ($406) {
        return $404($405($406));
    };
};

// | Group equal, consecutive elements of a list into lists.
// |
// | For example,
// |
// | ```purescript
// | group (1 : 1 : 2 : 2 : 1 : Nil) ==
// |   (NonEmptyList (NonEmpty 1 (1 : Nil))) : (NonEmptyList (NonEmpty 2 (2 : Nil))) : (NonEmptyList (NonEmpty 1 Nil)) : Nil
// | ```
// |
// | Running time: `O(n)`
var group = function (dictEq) {
    return groupBy(Data_Eq.eq(dictEq));
};

// | Group equal elements of a list into lists.
// |
// | For example,
// |
// | ```purescript
// | groupAll (1 : 1 : 2 : 2 : 1 : Nil) ==
// |   (NonEmptyList (NonEmpty 1 (1 : 1 : Nil))) : (NonEmptyList (NonEmpty 2 (2 : Nil))) : Nil
// | ```
var groupAll = function (dictOrd) {
    var $407 = group(dictOrd.Eq0());
    var $408 = sort(dictOrd);
    return function ($409) {
        return $407($408($409));
    };
};

// | Construct a list from a foldable structure.
// |
// | Running time: `O(n)`
var fromFoldable = function (dictFoldable) {
    return Data_Foldable.foldr(dictFoldable)(Data_List_Types.Cons.create)(Data_List_Types.Nil.value);
};

//------------------------------------------------------------------------------
// Folding ---------------------------------------------------------------------
//------------------------------------------------------------------------------
// | Perform a fold using a monadic step function.
var foldM = function (dictMonad) {
    var Applicative0 = dictMonad.Applicative0();
    var Bind1 = dictMonad.Bind1();
    return function (v) {
        return function (v1) {
            return function (v2) {
                if (v2 instanceof Data_List_Types.Nil) {
                    return Control_Applicative.pure(Applicative0)(v1);
                };
                if (v2 instanceof Data_List_Types.Cons) {
                    return Control_Bind.bind(Bind1)(v(v1)(v2.value0))(function (b$prime) {
                        return foldM(dictMonad)(v)(b$prime)(v2.value1);
                    });
                };
                throw new Error("Failed pattern match at Data.List (line 824, column 1 - line 824, column 72): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
            };
        };
    };
};

// | Find the first index for which a predicate holds.
var findIndex = function (fn) {
    var go = function ($copy_v) {
        return function ($copy_v1) {
            var $tco_var_v = $copy_v;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(v, v1) {
                if (v1 instanceof Data_List_Types.Cons) {
                    if (fn(v1.value0)) {
                        $tco_done = true;
                        return new Data_Maybe.Just(v);
                    };
                    if (Data_Boolean.otherwise) {
                        $tco_var_v = v + 1 | 0;
                        $copy_v1 = v1.value1;
                        return;
                    };
                };
                if (v1 instanceof Data_List_Types.Nil) {
                    $tco_done = true;
                    return Data_Maybe.Nothing.value;
                };
                throw new Error("Failed pattern match at Data.List (line 301, column 3 - line 301, column 35): " + [ v.constructor.name, v1.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_v, $copy_v1);
            };
            return $tco_result;
        };
    };
    return go(0);
};

// | Find the last index for which a predicate holds.
var findLastIndex = function (fn) {
    return function (xs) {
        return Data_Functor.map(Data_Maybe.functorMaybe)(function (v) {
            return (length(xs) - 1 | 0) - v | 0;
        })(findIndex(fn)(reverse(xs)));
    };
};

// | Filter where the predicate returns a monadic `Boolean`.
// |
// | For example:
// |
// | ```purescript
// | powerSet :: forall a. [a] -> [[a]]
// | powerSet = filterM (const [true, false])
// | ```
var filterM = function (dictMonad) {
    var Applicative0 = dictMonad.Applicative0();
    var Bind1 = dictMonad.Bind1();
    return function (v) {
        return function (v1) {
            if (v1 instanceof Data_List_Types.Nil) {
                return Control_Applicative.pure(Applicative0)(Data_List_Types.Nil.value);
            };
            if (v1 instanceof Data_List_Types.Cons) {
                return Control_Bind.bind(Bind1)(v(v1.value0))(function (b) {
                    return Control_Bind.bind(Bind1)(filterM(dictMonad)(v)(v1.value1))(function (xs$prime) {
                        return Control_Applicative.pure(Applicative0)((function () {
                            if (b) {
                                return new Data_List_Types.Cons(v1.value0, xs$prime);
                            };
                            return xs$prime;
                        })());
                    });
                });
            };
            throw new Error("Failed pattern match at Data.List (line 403, column 1 - line 403, column 75): " + [ v.constructor.name, v1.constructor.name ]);
        };
    };
};

// | Filter a list, keeping the elements which satisfy a predicate function.
// |
// | Running time: `O(n)`
var filter = function (p) {
    var go = function ($copy_v) {
        return function ($copy_v1) {
            var $tco_var_v = $copy_v;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(v, v1) {
                if (v1 instanceof Data_List_Types.Nil) {
                    $tco_done = true;
                    return reverse(v);
                };
                if (v1 instanceof Data_List_Types.Cons) {
                    if (p(v1.value0)) {
                        $tco_var_v = new Data_List_Types.Cons(v1.value0, v);
                        $copy_v1 = v1.value1;
                        return;
                    };
                    if (Data_Boolean.otherwise) {
                        $tco_var_v = v;
                        $copy_v1 = v1.value1;
                        return;
                    };
                };
                throw new Error("Failed pattern match at Data.List (line 390, column 3 - line 390, column 27): " + [ v.constructor.name, v1.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_v, $copy_v1);
            };
            return $tco_result;
        };
    };
    return go(Data_List_Types.Nil.value);
};

// | Calculate the intersection of two lists, using the specified
// | function to determine equality of elements.
// |
// | Running time: `O(n^2)`
var intersectBy = function (v) {
    return function (v1) {
        return function (v2) {
            if (v1 instanceof Data_List_Types.Nil) {
                return Data_List_Types.Nil.value;
            };
            if (v2 instanceof Data_List_Types.Nil) {
                return Data_List_Types.Nil.value;
            };
            return filter(function (x) {
                return Data_Foldable.any(Data_List_Types.foldableList)(Data_HeytingAlgebra.heytingAlgebraBoolean)(v(x))(v2);
            })(v1);
        };
    };
};

// | Calculate the intersection of two lists.
// |
// | Running time: `O(n^2)`
var intersect = function (dictEq) {
    return intersectBy(Data_Eq.eq(dictEq));
};

// | Remove duplicate elements from a list, using the provided equivalence function.
// | Keeps the first occurrence of each element in the input list,
// | in the same order they appear in the input list.
// | This less efficient version of `nubBy` only requires an equivalence
// | function, rather than an ordering function.
// |
// | ```purescript
// | mod3eq = eq `on` \n -> mod n 3
// | nubByEq mod3eq 1:3:4:5:6:Nil == 1:3:5:Nil
// | ```
// |
// | Running time: `O(n^2)`
var nubByEq = function (v) {
    return function (v1) {
        if (v1 instanceof Data_List_Types.Nil) {
            return Data_List_Types.Nil.value;
        };
        if (v1 instanceof Data_List_Types.Cons) {
            return new Data_List_Types.Cons(v1.value0, nubByEq(v)(filter(function (y) {
                return !v(v1.value0)(y);
            })(v1.value1)));
        };
        throw new Error("Failed pattern match at Data.List (line 705, column 1 - line 705, column 61): " + [ v.constructor.name, v1.constructor.name ]);
    };
};

// | Remove duplicate elements from a list.
// | Keeps the first occurrence of each element in the input list,
// | in the same order they appear in the input list.
// | This less efficient version of `nub` only requires an `Eq` instance.
// |
// | ```purescript
// | nubEq 1:2:1:3:3:Nil == 1:2:3:Nil
// | ```
// |
// | Running time: `O(n^2)`
var nubEq = function (dictEq) {
    return nubByEq(Data_Eq.eq(dictEq));
};
var eqPattern = function (dictEq) {
    var eqList = Data_List_Types.eqList(dictEq);
    return {
        eq: function (x) {
            return function (y) {
                return Data_Eq.eq(eqList)(x)(y);
            };
        }
    };
};
var ordPattern = function (dictOrd) {
    var ordList = Data_List_Types.ordList(dictOrd);
    var eqPattern1 = eqPattern(dictOrd.Eq0());
    return {
        compare: function (x) {
            return function (y) {
                return Data_Ord.compare(ordList)(x)(y);
            };
        },
        Eq0: function () {
            return eqPattern1;
        }
    };
};

// | Find the index of the last element equal to the specified element.
var elemLastIndex = function (dictEq) {
    return function (x) {
        return findLastIndex(function (v) {
            return Data_Eq.eq(dictEq)(v)(x);
        });
    };
};

// | Find the index of the first element equal to the specified element.
var elemIndex = function (dictEq) {
    return function (x) {
        return findIndex(function (v) {
            return Data_Eq.eq(dictEq)(v)(x);
        });
    };
};

// | Drop those elements from the front of a list which match a predicate.
// |
// | Running time (worst case): `O(n)`
var dropWhile = function (p) {
    var go = function ($copy_v) {
        var $tco_done = false;
        var $tco_result;
        function $tco_loop(v) {
            if (v instanceof Data_List_Types.Cons && p(v.value0)) {
                $copy_v = v.value1;
                return;
            };
            $tco_done = true;
            return v;
        };
        while (!$tco_done) {
            $tco_result = $tco_loop($copy_v);
        };
        return $tco_result;
    };
    return go;
};

// | Drop the specified number of elements from the end of a list.
// |
// | Running time: `O(2n - m)` where `n` is the number of elements in list
// | and `m` is number of elements to drop.
var dropEnd = function (n) {
    return function (xs) {
        return take(length(xs) - n | 0)(xs);
    };
};

// | Drop the specified number of elements from the front of a list.
// |
// | Running time: `O(n)` where `n` is the number of elements to drop.
var drop = function ($copy_v) {
    return function ($copy_v1) {
        var $tco_var_v = $copy_v;
        var $tco_done = false;
        var $tco_result;
        function $tco_loop(v, v1) {
            if (v < 1) {
                $tco_done = true;
                return v1;
            };
            if (v1 instanceof Data_List_Types.Nil) {
                $tco_done = true;
                return Data_List_Types.Nil.value;
            };
            if (v1 instanceof Data_List_Types.Cons) {
                $tco_var_v = v - 1 | 0;
                $copy_v1 = v1.value1;
                return;
            };
            throw new Error("Failed pattern match at Data.List (line 536, column 1 - line 536, column 42): " + [ v.constructor.name, v1.constructor.name ]);
        };
        while (!$tco_done) {
            $tco_result = $tco_loop($tco_var_v, $copy_v1);
        };
        return $tco_result;
    };
};

// | Extract a sublist by a start and end index.
var slice = function (start) {
    return function (end) {
        return function (xs) {
            return take(end - start | 0)(drop(start)(xs));
        };
    };
};

// | Take the specified number of elements from the end of a list.
// |
// | Running time: `O(2n - m)` where `n` is the number of elements in list
// | and `m` is number of elements to take.
var takeEnd = function (n) {
    return function (xs) {
        return drop(length(xs) - n | 0)(xs);
    };
};

// | Delete the first occurrence of an element from a list, using the specified
// | function to determine equality of elements.
// |
// | Running time: `O(n)`
var deleteBy = function (v) {
    return function (v1) {
        return function (v2) {
            if (v2 instanceof Data_List_Types.Nil) {
                return Data_List_Types.Nil.value;
            };
            if (v2 instanceof Data_List_Types.Cons && v(v1)(v2.value0)) {
                return v2.value1;
            };
            if (v2 instanceof Data_List_Types.Cons) {
                return new Data_List_Types.Cons(v2.value0, deleteBy(v)(v1)(v2.value1));
            };
            throw new Error("Failed pattern match at Data.List (line 732, column 1 - line 732, column 67): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
        };
    };
};

// | Calculate the union of two lists, using the specified
// | function to determine equality of elements.
// |
// | Running time: `O(n^2)`
var unionBy = function (eq) {
    return function (xs) {
        return function (ys) {
            return Data_Semigroup.append(Data_List_Types.semigroupList)(xs)(Data_Foldable.foldl(Data_List_Types.foldableList)(Data_Function.flip(deleteBy(eq)))(nubByEq(eq)(ys))(xs));
        };
    };
};

// | Calculate the union of two lists.
// |
// | Running time: `O(n^2)`
var union = function (dictEq) {
    return unionBy(Data_Eq.eq(dictEq));
};

// | Delete an element from a list at the specified index, returning a new
// | list or `Nothing` if the index is out-of-bounds.
// |
// | Running time: `O(n)`
var deleteAt = function (v) {
    return function (v1) {
        if (v === 0 && v1 instanceof Data_List_Types.Cons) {
            return new Data_Maybe.Just(v1.value1);
        };
        if (v1 instanceof Data_List_Types.Cons) {
            return Data_Functor.map(Data_Maybe.functorMaybe)(function (v2) {
                return new Data_List_Types.Cons(v1.value0, v2);
            })(deleteAt(v - 1 | 0)(v1.value1));
        };
        return Data_Maybe.Nothing.value;
    };
};

// | Delete the first occurrence of an element from a list.
// |
// | Running time: `O(n)`
var $$delete = function (dictEq) {
    return deleteBy(Data_Eq.eq(dictEq));
};

// | Delete the first occurrence of each element in the second list from the first list.
// |
// | Running time: `O(n^2)`
var difference = function (dictEq) {
    return Data_Foldable.foldl(Data_List_Types.foldableList)(Data_Function.flip($$delete(dictEq)));
};

// | Apply a function to each element in a list, and flatten the results
// | into a single, new list.
// |
// | Running time: `O(n)`, where `n` is the total number of elements.
var concatMap = /* #__PURE__ */ Data_Function.flip(/* #__PURE__ */ Control_Bind.bind(Data_List_Types.bindList));

// | Flatten a list of lists.
// |
// | Running time: `O(n)`, where `n` is the total number of elements.
var concat = function (v) {
    return Control_Bind.bind(Data_List_Types.bindList)(v)(identity);
};

// | Filter a list of optional values, keeping only the elements which contain
// | a value.
var catMaybes = /* #__PURE__ */ mapMaybe(/* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn));

// | Update or delete the element at the specified index by applying a
// | function to the current value, returning a new list or `Nothing` if the
// | index is out-of-bounds.
// |
// | Running time: `O(n)`
var alterAt = function (v) {
    return function (v1) {
        return function (v2) {
            if (v === 0 && v2 instanceof Data_List_Types.Cons) {
                return new Data_Maybe.Just((function () {
                    var v3 = v1(v2.value0);
                    if (v3 instanceof Data_Maybe.Nothing) {
                        return v2.value1;
                    };
                    if (v3 instanceof Data_Maybe.Just) {
                        return new Data_List_Types.Cons(v3.value0, v2.value1);
                    };
                    throw new Error("Failed pattern match at Data.List (line 352, column 3 - line 354, column 23): " + [ v3.constructor.name ]);
                })());
            };
            if (v2 instanceof Data_List_Types.Cons) {
                return Data_Functor.map(Data_Maybe.functorMaybe)(function (v3) {
                    return new Data_List_Types.Cons(v2.value0, v3);
                })(alterAt(v - 1 | 0)(v1)(v2.value1));
            };
            return Data_Maybe.Nothing.value;
        };
    };
};

// | Update the element at the specified index by applying a function to
// | the current value, returning a new list or `Nothing` if the index is
// | out-of-bounds.
// |
// | Running time: `O(n)`
var modifyAt = function (n) {
    return function (f) {
        return alterAt(n)(function ($410) {
            return Data_Maybe.Just.create(f($410));
        });
    };
};
export {
    toUnfoldable,
    fromFoldable,
    singleton,
    range,
    some,
    someRec,
    many,
    manyRec,
    $$null as null,
    length,
    snoc,
    insert,
    insertBy,
    head,
    last,
    tail,
    init,
    uncons,
    unsnoc,
    index,
    elemIndex,
    elemLastIndex,
    findIndex,
    findLastIndex,
    insertAt,
    deleteAt,
    updateAt,
    modifyAt,
    alterAt,
    reverse,
    concat,
    concatMap,
    filter,
    filterM,
    mapMaybe,
    catMaybes,
    sort,
    sortBy,
    Pattern,
    stripPrefix,
    slice,
    take,
    takeEnd,
    takeWhile,
    drop,
    dropEnd,
    dropWhile,
    span,
    group,
    groupAll,
    groupBy,
    groupAllBy,
    partition,
    nub,
    nubBy,
    nubEq,
    nubByEq,
    union,
    unionBy,
    $$delete as delete,
    deleteBy,
    difference,
    intersect,
    intersectBy,
    zipWith,
    zipWithA,
    zip,
    unzip,
    transpose,
    foldM,
    eqPattern,
    ordPattern,
    newtypePattern,
    showPattern
};
export {
    all,
    any,
    elem,
    find,
    findMap,
    fold,
    foldMap,
    foldl,
    foldr,
    intercalate,
    notElem
} from "../Data.Foldable/index.js";
export {
    Cons,
    Nil
} from "../Data.List.Types/index.js";
export {
    scanl,
    scanr
} from "../Data.Traversable/index.js";
