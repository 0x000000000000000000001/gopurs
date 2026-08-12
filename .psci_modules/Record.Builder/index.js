import * as $foreign from "./foreign.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Semigroupoid from "../Control.Semigroupoid/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Symbol from "../Data.Symbol/index.js";
import * as Record_Unsafe_Union from "../Record.Unsafe.Union/index.js";
import * as Unsafe_Coerce from "../Unsafe.Coerce/index.js";

// | A `Builder` can be used to `build` a record by incrementally adding
// | fields in-place, instead of using `insert` and repeatedly generating new
// | immutable records which need to be garbage collected.
// |
// | The mutations accumulated in a `Builder` are safe because intermediate states can't be
// | observed. These mutations, then, are performed all-at-once in the `build` function.
// |
// | The `Category` instance for `Builder` can be used to compose builders.
// |
// | For example:
// |
// | ```purescript
// | build (insert x 42 >>> insert y "testing") {} :: { x :: Int, y :: String }
// | ```
var Builder = function (x) {
    return x;
};

// | Build by merging existing fields from another record, taking precedence
// | in the case of overlaps. Unlike `merge`, this does not remove duplicate
// | labels from the resulting record type. This can result in better inference
// | for some pipelines, deferring the need for a `Nub` constraint.
// |
// | For example:
// |
// | ```purescript
// | build (union { x: 1, y: "y" }) { y: 2, z: true }
// |  :: { x :: Int, y :: String, y :: Int, z :: Boolean }
// | ```
var union = function () {
    return function (r1) {
        return function (r2) {
            return Record_Unsafe_Union.unsafeUnionFn(r1, r2);
        };
    };
};
var semigroupoidBuilder = Control_Semigroupoid.semigroupoidFn;

// | Build by renaming an existing field.
var rename = function (dictIsSymbol) {
    return function (dictIsSymbol1) {
        return function () {
            return function () {
                return function () {
                    return function () {
                        return function (l1) {
                            return function (l2) {
                                return function (r1) {
                                    return $foreign.unsafeRename(Data_Symbol.reflectSymbol(dictIsSymbol)(l1))(Data_Symbol.reflectSymbol(dictIsSymbol1)(l2))(r1);
                                };
                            };
                        };
                    };
                };
            };
        };
    };
};

// | A coercion which removes duplicate labels from a record's type.
var nub = function () {
    return Unsafe_Coerce.unsafeCoerce;
};

// | Build by modifying an existing field.
var modify = function () {
    return function () {
        return function (dictIsSymbol) {
            return function (l) {
                return function (f) {
                    return function (r1) {
                        return $foreign.unsafeModify(Data_Symbol.reflectSymbol(dictIsSymbol)(l))(f)(r1);
                    };
                };
            };
        };
    };
};

// | Build by merging existing fields from another record, taking precedence
// | in the case of overlaps.
// |
// | For example:
// |
// | ```purescript
// | build (merge { x: 1, y: "y" }) { y: 2, z: true }
// |  :: { x :: Int, y :: String, z :: Boolean }
// | ```
var merge = function () {
    return function () {
        return function (r1) {
            return function (r2) {
                return Record_Unsafe_Union.unsafeUnionFn(r1, r2);
            };
        };
    };
};

// | Build by inserting a new field.
var insert = function () {
    return function () {
        return function (dictIsSymbol) {
            return function (l) {
                return function (a) {
                    return function (r1) {
                        return $foreign.unsafeInsert(Data_Symbol.reflectSymbol(dictIsSymbol)(l))(a)(r1);
                    };
                };
            };
        };
    };
};

// | Build by merging some disjoint set of fields from another record.
var disjointUnion = function () {
    return function () {
        return function (r1) {
            return function (r2) {
                return Record_Unsafe_Union.unsafeUnionFn(r1, r2);
            };
        };
    };
};

// | Build by deleting an existing field.
var $$delete = function (dictIsSymbol) {
    return function () {
        return function () {
            return function (l) {
                return function (r2) {
                    return $foreign.unsafeDelete(Data_Symbol.reflectSymbol(dictIsSymbol)(l))(r2);
                };
            };
        };
    };
};
var categoryBuilder = Control_Category.categoryFn;

// | Build a record, starting from some other record.
var build = function (v) {
    return function (r1) {
        return v($foreign.copyRecord(r1));
    };
};

// | Build a record from scratch.
var buildFromScratch = /* #__PURE__ */ Data_Function.flip(build)({});

// | Flip a function of one argument returning a builder.
var flip = function (f) {
    return function (b) {
        return function (a) {
            return build(f(a))(b);
        };
    };
};
export {
    build,
    buildFromScratch,
    flip,
    insert,
    modify,
    $$delete as delete,
    rename,
    merge,
    union,
    disjointUnion,
    nub,
    semigroupoidBuilder,
    categoryBuilder
};
