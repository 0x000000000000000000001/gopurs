import * as Control_Category from "../Control.Category/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Monoid_Dual from "../Data.Monoid.Dual/index.js";
import * as Data_Monoid_Multiplicative from "../Data.Monoid.Multiplicative/index.js";
import * as Data_Semigroup_Foldable from "../Data.Semigroup.Foldable/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var traverse1 = function (dict) {
    return dict.traverse1;
};
var traversableTuple = {
    traverse1: function (dictApply) {
        var Functor0 = dictApply.Functor0();
        return function (f) {
            return function (v) {
                return Data_Functor.map(Functor0)(Data_Tuple.Tuple.create(v.value0))(f(v.value1));
            };
        };
    },
    sequence1: function (dictApply) {
        var Functor0 = dictApply.Functor0();
        return function (v) {
            return Data_Functor.map(Functor0)(Data_Tuple.Tuple.create(v.value0))(v.value1);
        };
    },
    Foldable10: function () {
        return Data_Semigroup_Foldable.foldableTuple;
    },
    Traversable1: function () {
        return Data_Traversable.traversableTuple;
    }
};
var traversableIdentity = {
    traverse1: function (dictApply) {
        var Functor0 = dictApply.Functor0();
        return function (f) {
            return function (v) {
                return Data_Functor.map(Functor0)(Data_Identity.Identity)(f(v));
            };
        };
    },
    sequence1: function (dictApply) {
        var Functor0 = dictApply.Functor0();
        return function (v) {
            return Data_Functor.map(Functor0)(Data_Identity.Identity)(v);
        };
    },
    Foldable10: function () {
        return Data_Semigroup_Foldable.foldableIdentity;
    },
    Traversable1: function () {
        return Data_Traversable.traversableIdentity;
    }
};

// | A default implementation of `sequence1` using `traverse1`.
var sequence1Default = function (dictTraversable1) {
    return function (dictApply) {
        return traverse1(dictTraversable1)(dictApply)(identity);
    };
};
var traversableDual = {
    traverse1: function (dictApply) {
        var Functor0 = dictApply.Functor0();
        return function (f) {
            return function (v) {
                return Data_Functor.map(Functor0)(Data_Monoid_Dual.Dual)(f(v));
            };
        };
    },
    sequence1: function (dictApply) {
        return sequence1Default(traversableDual)(dictApply);
    },
    Foldable10: function () {
        return Data_Semigroup_Foldable.foldableDual;
    },
    Traversable1: function () {
        return Data_Traversable.traversableDual;
    }
};
var traversableMultiplicative = {
    traverse1: function (dictApply) {
        var Functor0 = dictApply.Functor0();
        return function (f) {
            return function (v) {
                return Data_Functor.map(Functor0)(Data_Monoid_Multiplicative.Multiplicative)(f(v));
            };
        };
    },
    sequence1: function (dictApply) {
        return sequence1Default(traversableMultiplicative)(dictApply);
    },
    Foldable10: function () {
        return Data_Semigroup_Foldable.foldableMultiplicative;
    },
    Traversable1: function () {
        return Data_Traversable.traversableMultiplicative;
    }
};
var sequence1 = function (dict) {
    return dict.sequence1;
};

// | A default implementation of `traverse1` using `sequence1`.
var traverse1Default = function (dictTraversable1) {
    var Functor0 = (dictTraversable1.Traversable1()).Functor0();
    return function (dictApply) {
        return function (f) {
            return function (ta) {
                return sequence1(dictTraversable1)(dictApply)(Data_Functor.map(Functor0)(f)(ta));
            };
        };
    };
};
export {
    sequence1,
    traverse1,
    traverse1Default,
    sequence1Default,
    traversableDual,
    traversableMultiplicative,
    traversableTuple,
    traversableIdentity
};
