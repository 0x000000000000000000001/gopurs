// | This module defines a type of sets as height-balanced (AVL) binary trees.
// | Efficient set operations are implemented in terms of
// | <https://www.cs.cmu.edu/~guyb/papers/BFS16.pdf>
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_List from "../Data.List/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Safe_Coerce from "../Safe.Coerce/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);

// | `Set a` represents a set of values of type `a`
var $$Set = function (x) {
    return x;
};

// | Form the union of two sets
// |
// | Running time: `O(n + m)`
var union = function (dictOrd) {
    return Safe_Coerce.coerce()(Data_Map_Internal.union(dictOrd));
};

// | Insert a value into a set if it is not already present, if it is present, delete it.
var toggle = function (dictOrd) {
    return function (a) {
        return function (v) {
            return Data_Map_Internal.alter(dictOrd)(Data_Maybe.maybe(new Data_Maybe.Just(Data_Unit.unit))(function (v1) {
                return Data_Maybe.Nothing.value;
            }))(a)(v);
        };
    };
};

// | A set is a map with no value attached to each key.
var toMap = function (v) {
    return v;
};
var toList = function (v) {
    return Data_Map_Internal.keys(v);
};

// | Convert a set to an unfoldable structure.
var toUnfoldable = function (dictUnfoldable) {
    var $70 = Data_List.toUnfoldable(dictUnfoldable);
    return function ($71) {
        return $70(toList($71));
    };
};

// | Find the size of a set
var size = /* #__PURE__ */ Safe_Coerce.coerce()(Data_Map_Internal.size);

// | Create a set with one element
var singleton = function (a) {
    return Data_Map_Internal.singleton(a)(Data_Unit.unit);
};
var showSet = function (dictShow) {
    var showArray = Data_Show.showArray(dictShow);
    return {
        show: function (s) {
            return "(fromFoldable " + (Data_Show.show(showArray)(toUnfoldable(Data_Unfoldable.unfoldableArray)(s)) + ")");
        }
    };
};
var semigroupSet = function (dictOrd) {
    return {
        append: union(dictOrd)
    };
};

// | Test if a value is a member of a set
var member = function (dictOrd) {
    return Safe_Coerce.coerce()(Data_Map_Internal.member(dictOrd));
};

// | Test if a set is empty
var isEmpty = /* #__PURE__ */ Safe_Coerce.coerce()(Data_Map_Internal.isEmpty);

// | The set of elements which are in both the first and second set
var intersection = function (dictOrd) {
    return Safe_Coerce.coerce()(Data_Map_Internal.intersection(dictOrd));
};

// | Insert a value into a set
var insert = function (dictOrd) {
    return function (a) {
        return function (v) {
            return Data_Map_Internal.insert(dictOrd)(a)(Data_Unit.unit)(v);
        };
    };
};

// | A map with no value attached to each key is a set.
// | See also `Data.Map.keys`.
var fromMap = $$Set;
var foldableSet = {
    foldMap: function (dictMonoid) {
        return function (f) {
            var $72 = Data_Foldable.foldMap(Data_List_Types.foldableList)(dictMonoid)(f);
            return function ($73) {
                return $72(toList($73));
            };
        };
    },
    foldl: function (f) {
        return function (x) {
            var $74 = Data_Foldable.foldl(Data_List_Types.foldableList)(f)(x);
            return function ($75) {
                return $74(toList($75));
            };
        };
    },
    foldr: function (f) {
        return function (x) {
            var $76 = Data_Foldable.foldr(Data_List_Types.foldableList)(f)(x);
            return function ($77) {
                return $76(toList($77));
            };
        };
    }
};
var findMin = function (v) {
    return Data_Functor.map(Data_Maybe.functorMaybe)(function (v1) {
        return v1.key;
    })(Data_Map_Internal.findMin(v));
};
var findMax = function (v) {
    return Data_Functor.map(Data_Maybe.functorMaybe)(function (v1) {
        return v1.key;
    })(Data_Map_Internal.findMax(v));
};

// | Filter out those values of a set for which a predicate on the value fails
// | to hold.
var filter = function (dictOrd) {
    return Safe_Coerce.coerce()(Data_Map_Internal.filterKeys(dictOrd));
};
var eqSet = function (dictEq) {
    var eqMap = Data_Map_Internal.eqMap(dictEq)(Data_Eq.eqUnit);
    return {
        eq: function (v) {
            return function (v1) {
                return Data_Eq.eq(eqMap)(v)(v1);
            };
        }
    };
};
var ordSet = function (dictOrd) {
    var ordList = Data_List_Types.ordList(dictOrd);
    var eqSet1 = eqSet(dictOrd.Eq0());
    return {
        compare: function (s1) {
            return function (s2) {
                return Data_Ord.compare(ordList)(toList(s1))(toList(s2));
            };
        },
        Eq0: function () {
            return eqSet1;
        }
    };
};
var eq1Set = {
    eq1: function (dictEq) {
        return Data_Eq.eq(eqSet(dictEq));
    }
};
var ord1Set = {
    compare1: function (dictOrd) {
        return Data_Ord.compare(ordSet(dictOrd));
    },
    Eq10: function () {
        return eq1Set;
    }
};

// | An empty set
var empty = Data_Map_Internal.empty;

// | Create a set from a foldable structure.
var fromFoldable = function (dictFoldable) {
    return function (dictOrd) {
        return Data_Foldable.foldl(dictFoldable)(function (m) {
            return function (a) {
                return insert(dictOrd)(a)(m);
            };
        })(empty);
    };
};

// | Maps over the values in a set.
// |
// | This operation is not structure-preserving for sets, so is not a valid
// | `Functor`. An example case: mapping `const x` over a set with `n > 0`
// | elements will result in a set with one element.
var map = function (dictOrd) {
    return function (f) {
        return Data_Foldable.foldl(foldableSet)(function (m) {
            return function (a) {
                return insert(dictOrd)(f(a))(m);
            };
        })(empty);
    };
};

// | Applies a function to each value in a set, discarding entries where the
// | function returns `Nothing`.
var mapMaybe = function (dictOrd) {
    return function (f) {
        return Data_Foldable.foldr(foldableSet)(function (a) {
            return function (acc) {
                return Data_Maybe.maybe(acc)(function (b) {
                    return insert(dictOrd)(b)(acc);
                })(f(a));
            };
        })(empty);
    };
};
var monoidSet = function (dictOrd) {
    var semigroupSet1 = semigroupSet(dictOrd);
    return {
        mempty: empty,
        Semigroup0: function () {
            return semigroupSet1;
        }
    };
};

// | Form the union of a collection of sets
var unions = function (dictFoldable) {
    return function (dictOrd) {
        return Data_Foldable.foldl(dictFoldable)(union(dictOrd))(empty);
    };
};

// | Form the set difference
var difference = function (dictOrd) {
    return Safe_Coerce.coerce()(Data_Map_Internal.difference(dictOrd));
};

// | True if and only if every element in the first set
// | is an element of the second set
var subset = function (dictOrd) {
    return function (s1) {
        return function (s2) {
            return isEmpty(difference(dictOrd)(s1)(s2));
        };
    };
};

// | True if and only if the first set is a subset of the second set
// | and the sets are not equal
var properSubset = function (dictOrd) {
    return function (s1) {
        return function (s2) {
            return size(s1) !== size(s2) && subset(dictOrd)(s1)(s2);
        };
    };
};

// | Delete a value from a set
var $$delete = function (dictOrd) {
    return Safe_Coerce.coerce()(Data_Map_Internal["delete"](dictOrd));
};

// | Check whether the underlying tree satisfies the height, size, and ordering invariants.
// |
// | This function is provided for internal use.
var checkValid = function (dictOrd) {
    return Safe_Coerce.coerce()(Data_Map_Internal.checkValid(dictOrd));
};

// | Filter a set of optional values, discarding values that contain `Nothing`
var catMaybes = function (dictOrd) {
    return mapMaybe(dictOrd)(identity);
};
export {
    fromFoldable,
    toUnfoldable,
    empty,
    isEmpty,
    singleton,
    map,
    checkValid,
    insert,
    member,
    $$delete as delete,
    toggle,
    size,
    findMin,
    findMax,
    union,
    unions,
    difference,
    subset,
    properSubset,
    intersection,
    filter,
    mapMaybe,
    catMaybes,
    toMap,
    fromMap,
    eqSet,
    eq1Set,
    showSet,
    ordSet,
    ord1Set,
    monoidSet,
    semigroupSet,
    foldableSet
};
