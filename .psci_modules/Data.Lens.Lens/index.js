// | This module defines functions for working with lenses.
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Lens_Internal_Indexed from "../Data.Lens.Internal.Indexed/index.js";
import * as Data_Lens_Internal_Shop from "../Data.Lens.Internal.Shop/index.js";
import * as Data_Lens_Types from "../Data.Lens.Types/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Profunctor from "../Data.Profunctor/index.js";
import * as Data_Profunctor_Strong from "../Data.Profunctor.Strong/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var identity1 = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var withLens = function (l) {
    return function (f) {
        var v = l(new Data_Lens_Internal_Shop.Shop(identity, function (v1) {
            return function (b) {
                return b;
            };
        }));
        return f(v.value0)(v.value1);
    };
};
var withIndexedLens = function (l) {
    return function (f) {
        var v = l(new Data_Lens_Internal_Shop.Shop(identity1, function (v1) {
            return function (b) {
                return b;
            };
        }));
        return f(v.value0)(v.value1);
    };
};

// | Converts a lens into the form that `lens'` accepts.
// |
// | Can be useful when defining a lens where the focus appears under multiple
// | constructors of an algebraic data type.  This function would be called for
// | each case of the data type.
// |
// | For example:
// |
// | ```
// | data LensStoreExample = LensStoreA Int | LensStoreB (Tuple Boolean Int)
// |
// | lensStoreExampleInt :: Lens' LensStoreExample Int
// | lensStoreExampleInt = lens' case _ of
// |   LensStoreA i -> map LensStoreA <$> lensStore identity i
// |   LensStoreB i -> map LensStoreB <$> lensStore _2 i
// | ```
var lensStore = function (l) {
    return withLens(l)(Control_Apply.lift2(Control_Apply.applyFn)(Data_Tuple.Tuple.create));
};
var lens$prime = function (to) {
    return function (dictStrong) {
        var Profunctor0 = dictStrong.Profunctor0();
        return function (pab) {
            return Data_Profunctor.dimap(Profunctor0)(to)(function (v) {
                return v.value1(v.value0);
            })(Data_Profunctor_Strong.first(dictStrong)(pab));
        };
    };
};

// | Create a `Lens` from a getter/setter pair.
// |
// | ```purescript
// | > species = lens _.species $ _ {species = _}
// | > view species {species : "bovine"}
// | "bovine"
// |
// | > _2 = lens Tuple.snd $ \(Tuple keep _) new -> Tuple keep new
// | ```
// |
// | Note: `_2` is predefined in `Data.Lens.Tuple`.
var lens = function (get) {
    return function (set) {
        return function (dictStrong) {
            return lens$prime(function (s) {
                return new Data_Tuple.Tuple(get(s), function (b) {
                    return set(s)(b);
                });
            })(dictStrong);
        };
    };
};
var ilens$prime = function (to) {
    return function (dictStrong) {
        var Profunctor0 = dictStrong.Profunctor0();
        return function (pab) {
            return Data_Profunctor.dimap(Profunctor0)(to)(function (v) {
                return v.value1(v.value0);
            })(Data_Profunctor_Strong.first(dictStrong)(Data_Newtype.un()(Data_Lens_Internal_Indexed.Indexed)(pab)));
        };
    };
};

// create an `IndexedLens` from a getter/setter pair.
var ilens = function (get) {
    return function (set) {
        return function (dictStrong) {
            return ilens$prime(function (s) {
                return new Data_Tuple.Tuple(get(s), function (b) {
                    return set(s)(b);
                });
            })(dictStrong);
        };
    };
};
var cloneLens = function (l) {
    return function (dictStrong) {
        return withLens(l)(function (x) {
            return function (y) {
                return function (p) {
                    return lens(x)(y)(dictStrong)(p);
                };
            };
        });
    };
};
var cloneIndexedLens = function (l) {
    return function (dictStrong) {
        return withIndexedLens(l)(function (x) {
            return function (y) {
                return function (p) {
                    return ilens(x)(y)(dictStrong)(p);
                };
            };
        });
    };
};
export {
    lens,
    lens$prime,
    withLens,
    cloneLens,
    ilens,
    ilens$prime,
    withIndexedLens,
    cloneIndexedLens,
    lensStore
};
