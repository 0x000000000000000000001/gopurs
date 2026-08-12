import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Parallel_Class from "../Control.Parallel.Class/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);

// | Traverse a collection in parallel, discarding any results.
var parTraverse_ = function (dictParallel) {
    var sequential = Control_Parallel_Class.sequential(dictParallel);
    var parallel = Control_Parallel_Class.parallel(dictParallel);
    return function (dictApplicative) {
        return function (dictFoldable) {
            return function (f) {
                var $33 = Data_Foldable.traverse_(dictApplicative)(dictFoldable)(function ($35) {
                    return parallel(f($35));
                });
                return function ($34) {
                    return sequential($33($34));
                };
            };
        };
    };
};

// | Traverse a collection in parallel.
var parTraverse = function (dictParallel) {
    var sequential = Control_Parallel_Class.sequential(dictParallel);
    var parallel = Control_Parallel_Class.parallel(dictParallel);
    return function (dictApplicative) {
        return function (dictTraversable) {
            return function (f) {
                var $36 = Data_Traversable.traverse(dictTraversable)(dictApplicative)(function ($38) {
                    return parallel(f($38));
                });
                return function ($37) {
                    return sequential($36($37));
                };
            };
        };
    };
};
var parSequence_ = function (dictParallel) {
    return function (dictApplicative) {
        return function (dictFoldable) {
            return parTraverse_(dictParallel)(dictApplicative)(dictFoldable)(identity);
        };
    };
};
var parSequence = function (dictParallel) {
    return function (dictApplicative) {
        return function (dictTraversable) {
            return parTraverse(dictParallel)(dictApplicative)(dictTraversable)(identity);
        };
    };
};

// | Race a collection in parallel while mapping to some effect.
var parOneOfMap = function (dictParallel) {
    var sequential = Control_Parallel_Class.sequential(dictParallel);
    var parallel = Control_Parallel_Class.parallel(dictParallel);
    return function (dictAlternative) {
        var Plus1 = dictAlternative.Plus1();
        return function (dictFoldable) {
            return function (dictFunctor) {
                return function (f) {
                    var $39 = Data_Foldable.oneOfMap(dictFoldable)(Plus1)(function ($41) {
                        return parallel(f($41));
                    });
                    return function ($40) {
                        return sequential($39($40));
                    };
                };
            };
        };
    };
};

// | Race a collection in parallel.
var parOneOf = function (dictParallel) {
    var sequential = Control_Parallel_Class.sequential(dictParallel);
    var parallel = Control_Parallel_Class.parallel(dictParallel);
    return function (dictAlternative) {
        var Plus1 = dictAlternative.Plus1();
        return function (dictFoldable) {
            return function (dictFunctor) {
                var $42 = Data_Foldable.oneOfMap(dictFoldable)(Plus1)(parallel);
                return function ($43) {
                    return sequential($42($43));
                };
            };
        };
    };
};

// | Apply a function to an argument under a type constructor in parallel.
var parApply = function (dictParallel) {
    var Apply1 = dictParallel.Apply1();
    return function (mf) {
        return function (ma) {
            return Control_Parallel_Class.sequential(dictParallel)(Control_Apply.apply(Apply1)(Control_Parallel_Class.parallel(dictParallel)(mf))(Control_Parallel_Class.parallel(dictParallel)(ma)));
        };
    };
};
export {
    parApply,
    parTraverse,
    parTraverse_,
    parSequence,
    parSequence_,
    parOneOf,
    parOneOfMap
};
export {
    ParCont,
    parallel,
    sequential
} from "../Control.Parallel.Class/index.js";
