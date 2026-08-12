// | Helper functions for working with mutable arrays using the `ST` effect.
// |
// | This module can be used when performance is important and mutation is a local effect.
import * as $foreign from "./foreign.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad_ST_Internal from "../Control.Monad.ST.Internal/index.js";
import * as Control_Monad_ST_Uncurried from "../Control.Monad.ST.Uncurried/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";

// | Append the values in an immutable array to the front of a mutable array.
// | Returns the new length of the mutable array.
var unshiftAll = /* #__PURE__ */ Control_Monad_ST_Uncurried.runSTFn2($foreign.unshiftAllImpl);

// | Append an element to the front of a mutable array. Returns the new length of
// | the array.
var unshift = function (a) {
    return Control_Monad_ST_Uncurried.runSTFn2($foreign.unshiftAllImpl)([ a ]);
};

// | O(1) Convert an immutable array to a mutable array, without copying. The input
// | array must not be used afterward.
var unsafeThaw = /* #__PURE__ */ Control_Monad_ST_Uncurried.runSTFn1($foreign.unsafeThawImpl);

// | O(1). Convert a mutable array to an immutable array, without copying. The mutable
// | array must not be mutated afterwards.
var unsafeFreeze = /* #__PURE__ */ Control_Monad_ST_Uncurried.runSTFn1($foreign.unsafeFreezeImpl);

// | Create an immutable copy of a mutable array, where each element
// | is labelled with its index in the original array.
var toAssocArray = /* #__PURE__ */ Control_Monad_ST_Uncurried.runSTFn1($foreign.toAssocArrayImpl);
var thaw = /* #__PURE__ */ Control_Monad_ST_Uncurried.runSTFn1($foreign.thawImpl);

// | Perform an effect requiring a mutable array on a copy of an immutable array,
// | safely returning the result as an immutable array.
var withArray = function (f) {
    return function (xs) {
        return function __do() {
            var result = thaw(xs)();
            f(result)();
            return unsafeFreeze(result)();
        };
    };
};

// | Remove and/or insert elements from/into a mutable array at the specified index.
var splice = /* #__PURE__ */ Control_Monad_ST_Uncurried.runSTFn4($foreign.spliceImpl);

// | Sort a mutable array in place using a comparison function. Sorting is
// | stable: the order of elements is preserved if they are equal according to
// | the comparison function.
var sortBy = function (comp) {
    return Control_Monad_ST_Uncurried.runSTFn3($foreign.sortByImpl)(comp)(function (v) {
        if (v instanceof Data_Ordering.GT) {
            return 1;
        };
        if (v instanceof Data_Ordering.EQ) {
            return 0;
        };
        if (v instanceof Data_Ordering.LT) {
            return -1 | 0;
        };
        throw new Error("Failed pattern match at Data.Array.ST (line 129, column 40 - line 132, column 11): " + [ v.constructor.name ]);
    });
};

// | Sort a mutable array in place based on a projection. Sorting is stable: the
// | order of elements is preserved if they are equal according to the projection.
var sortWith = function (dictOrd) {
    return function (f) {
        return sortBy(Data_Ord.comparing(dictOrd)(f));
    };
};

// | Sort a mutable array in place. Sorting is stable: the order of equal
// | elements is preserved.
var sort = function (dictOrd) {
    return sortBy(Data_Ord.compare(dictOrd));
};

// | Remove the first element from an array and return that element.
var shift = /* #__PURE__ */ (function () {
    return Control_Monad_ST_Uncurried.runSTFn3($foreign.shiftImpl)(Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();

// | A safe way to create and work with a mutable array before returning an
// | immutable array for later perusal. This function avoids copying the array
// | before returning it - it uses unsafeFreeze internally, but this wrapper is
// | a safe interface to that function.
var run = function (st) {
    return Control_Bind.bind(Control_Monad_ST_Internal.bindST)(st)(unsafeFreeze)();
};

// | Append the values in an immutable array to the end of a mutable array.
// | Returns the new length of the mutable array.
var pushAll = /* #__PURE__ */ Control_Monad_ST_Uncurried.runSTFn2($foreign.pushAllImpl);

// | Append an element to the end of a mutable array. Returns the new length of
// | the array.
var push = /* #__PURE__ */ Control_Monad_ST_Uncurried.runSTFn2($foreign.pushImpl);

// | Remove the last element from an array and return that element.
var pop = /* #__PURE__ */ (function () {
    return Control_Monad_ST_Uncurried.runSTFn3($foreign.popImpl)(Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();
var poke = /* #__PURE__ */ Control_Monad_ST_Uncurried.runSTFn3($foreign.pokeImpl);

// | Read the value at the specified index in a mutable array.
var peek = /* #__PURE__ */ (function () {
    return Control_Monad_ST_Uncurried.runSTFn4($foreign.peekImpl)(Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();

// | Mutate the element at the specified index using the supplied function.
var modify = function (i) {
    return function (f) {
        return function (xs) {
            return function __do() {
                var entry = peek(i)(xs)();
                if (entry instanceof Data_Maybe.Just) {
                    return poke(i)(f(entry.value0))(xs)();
                };
                if (entry instanceof Data_Maybe.Nothing) {
                    return false;
                };
                throw new Error("Failed pattern match at Data.Array.ST (line 234, column 3 - line 236, column 26): " + [ entry.constructor.name ]);
            };
        };
    };
};

// | Get the number of elements in a mutable array.
var length = /* #__PURE__ */ Control_Monad_ST_Uncurried.runSTFn1($foreign.lengthImpl);

// | Create an immutable copy of a mutable array.
var freeze = /* #__PURE__ */ Control_Monad_ST_Uncurried.runSTFn1($foreign.freezeImpl);

// | Make a mutable copy of a mutable array.
var clone = /* #__PURE__ */ Control_Monad_ST_Uncurried.runSTFn1($foreign.cloneImpl);
export {
    new
} from "./foreign.js";
export {
    run,
    withArray,
    peek,
    poke,
    modify,
    length,
    pop,
    push,
    pushAll,
    shift,
    unshift,
    unshiftAll,
    splice,
    sort,
    sortBy,
    sortWith,
    freeze,
    thaw,
    clone,
    unsafeFreeze,
    unsafeThaw,
    toAssocArray
};
