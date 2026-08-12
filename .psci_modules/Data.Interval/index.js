import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Extend from "../Control.Extend/index.js";
import * as Data_Bifoldable from "../Data.Bifoldable/index.js";
import * as Data_Bifunctor from "../Data.Bifunctor/index.js";
import * as Data_Bitraversable from "../Data.Bitraversable/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Interval_Duration from "../Data.Interval.Duration/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
var showMaybe = /* #__PURE__ */ Data_Maybe.showMaybe(Data_Show.showInt);
var eqMaybe = /* #__PURE__ */ Data_Maybe.eqMaybe(Data_Eq.eqInt);
var ordMaybe = /* #__PURE__ */ Data_Maybe.ordMaybe(Data_Ord.ordInt);
var StartEnd = /* #__PURE__ */ (function () {
    function StartEnd(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    StartEnd.create = function (value0) {
        return function (value1) {
            return new StartEnd(value0, value1);
        };
    };
    return StartEnd;
})();
var DurationEnd = /* #__PURE__ */ (function () {
    function DurationEnd(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    DurationEnd.create = function (value0) {
        return function (value1) {
            return new DurationEnd(value0, value1);
        };
    };
    return DurationEnd;
})();
var StartDuration = /* #__PURE__ */ (function () {
    function StartDuration(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    StartDuration.create = function (value0) {
        return function (value1) {
            return new StartDuration(value0, value1);
        };
    };
    return StartDuration;
})();
var DurationOnly = /* #__PURE__ */ (function () {
    function DurationOnly(value0) {
        this.value0 = value0;
    };
    DurationOnly.create = function (value0) {
        return new DurationOnly(value0);
    };
    return DurationOnly;
})();
var RecurringInterval = /* #__PURE__ */ (function () {
    function RecurringInterval(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    RecurringInterval.create = function (value0) {
        return function (value1) {
            return new RecurringInterval(value0, value1);
        };
    };
    return RecurringInterval;
})();
var showInterval = function (dictShow) {
    return function (dictShow1) {
        return {
            show: function (v) {
                if (v instanceof StartEnd) {
                    return "(StartEnd " + (Data_Show.show(dictShow1)(v.value0) + (" " + (Data_Show.show(dictShow1)(v.value1) + ")")));
                };
                if (v instanceof DurationEnd) {
                    return "(DurationEnd " + (Data_Show.show(dictShow)(v.value0) + (" " + (Data_Show.show(dictShow1)(v.value1) + ")")));
                };
                if (v instanceof StartDuration) {
                    return "(StartDuration " + (Data_Show.show(dictShow1)(v.value0) + (" " + (Data_Show.show(dictShow)(v.value1) + ")")));
                };
                if (v instanceof DurationOnly) {
                    return "(DurationOnly " + (Data_Show.show(dictShow)(v.value0) + ")");
                };
                throw new Error("Failed pattern match at Data.Interval (line 66, column 1 - line 70, column 60): " + [ v.constructor.name ]);
            }
        };
    };
};
var showRecurringInterval = function (dictShow) {
    var showInterval1 = showInterval(dictShow);
    return function (dictShow1) {
        var showInterval2 = showInterval1(dictShow1);
        return {
            show: function (v) {
                return "(RecurringInterval " + (Data_Show.show(showMaybe)(v.value0) + (" " + (Data_Show.show(showInterval2)(v.value1) + ")")));
            }
        };
    };
};
var over = function (dictFunctor) {
    return function (f) {
        return function (v) {
            return Data_Functor.map(dictFunctor)(RecurringInterval.create(v.value0))(f(v.value1));
        };
    };
};
var interval = function (v) {
    return v.value1;
};
var foldableInterval = {
    foldl: function (v) {
        return function (v1) {
            return function (v2) {
                if (v2 instanceof StartEnd) {
                    return v(v(v1)(v2.value0))(v2.value1);
                };
                if (v2 instanceof DurationEnd) {
                    return v(v1)(v2.value1);
                };
                if (v2 instanceof StartDuration) {
                    return v(v1)(v2.value0);
                };
                return v1;
            };
        };
    },
    foldr: function (x) {
        return Data_Foldable.foldrDefault(foldableInterval)(x);
    },
    foldMap: function (dictMonoid) {
        return Data_Foldable.foldMapDefaultL(foldableInterval)(dictMonoid);
    }
};
var foldableRecurringInterval = {
    foldl: function (f) {
        return function (i) {
            var $275 = Data_Foldable.foldl(foldableInterval)(f)(i);
            return function ($276) {
                return $275(interval($276));
            };
        };
    },
    foldr: function (f) {
        return function (i) {
            var $277 = Data_Foldable.foldr(foldableInterval)(f)(i);
            return function ($278) {
                return $277(interval($278));
            };
        };
    },
    foldMap: function (dictMonoid) {
        return Data_Foldable.foldMapDefaultL(foldableRecurringInterval)(dictMonoid);
    }
};
var eqInterval = function (dictEq) {
    return function (dictEq1) {
        return {
            eq: function (x) {
                return function (y) {
                    if (x instanceof StartEnd && y instanceof StartEnd) {
                        return Data_Eq.eq(dictEq1)(x.value0)(y.value0) && Data_Eq.eq(dictEq1)(x.value1)(y.value1);
                    };
                    if (x instanceof DurationEnd && y instanceof DurationEnd) {
                        return Data_Eq.eq(dictEq)(x.value0)(y.value0) && Data_Eq.eq(dictEq1)(x.value1)(y.value1);
                    };
                    if (x instanceof StartDuration && y instanceof StartDuration) {
                        return Data_Eq.eq(dictEq1)(x.value0)(y.value0) && Data_Eq.eq(dictEq)(x.value1)(y.value1);
                    };
                    if (x instanceof DurationOnly && y instanceof DurationOnly) {
                        return Data_Eq.eq(dictEq)(x.value0)(y.value0);
                    };
                    return false;
                };
            }
        };
    };
};
var eqRecurringInterval = function (dictEq) {
    var eqInterval1 = eqInterval(dictEq);
    return function (dictEq1) {
        var eqInterval2 = eqInterval1(dictEq1);
        return {
            eq: function (x) {
                return function (y) {
                    return Data_Eq.eq(eqMaybe)(x.value0)(y.value0) && Data_Eq.eq(eqInterval2)(x.value1)(y.value1);
                };
            }
        };
    };
};
var ordInterval = function (dictOrd) {
    var eqInterval1 = eqInterval(dictOrd.Eq0());
    return function (dictOrd1) {
        var eqInterval2 = eqInterval1(dictOrd1.Eq0());
        return {
            compare: function (x) {
                return function (y) {
                    if (x instanceof StartEnd && y instanceof StartEnd) {
                        var v = Data_Ord.compare(dictOrd1)(x.value0)(y.value0);
                        if (v instanceof Data_Ordering.LT) {
                            return Data_Ordering.LT.value;
                        };
                        if (v instanceof Data_Ordering.GT) {
                            return Data_Ordering.GT.value;
                        };
                        return Data_Ord.compare(dictOrd1)(x.value1)(y.value1);
                    };
                    if (x instanceof StartEnd) {
                        return Data_Ordering.LT.value;
                    };
                    if (y instanceof StartEnd) {
                        return Data_Ordering.GT.value;
                    };
                    if (x instanceof DurationEnd && y instanceof DurationEnd) {
                        var v = Data_Ord.compare(dictOrd)(x.value0)(y.value0);
                        if (v instanceof Data_Ordering.LT) {
                            return Data_Ordering.LT.value;
                        };
                        if (v instanceof Data_Ordering.GT) {
                            return Data_Ordering.GT.value;
                        };
                        return Data_Ord.compare(dictOrd1)(x.value1)(y.value1);
                    };
                    if (x instanceof DurationEnd) {
                        return Data_Ordering.LT.value;
                    };
                    if (y instanceof DurationEnd) {
                        return Data_Ordering.GT.value;
                    };
                    if (x instanceof StartDuration && y instanceof StartDuration) {
                        var v = Data_Ord.compare(dictOrd1)(x.value0)(y.value0);
                        if (v instanceof Data_Ordering.LT) {
                            return Data_Ordering.LT.value;
                        };
                        if (v instanceof Data_Ordering.GT) {
                            return Data_Ordering.GT.value;
                        };
                        return Data_Ord.compare(dictOrd)(x.value1)(y.value1);
                    };
                    if (x instanceof StartDuration) {
                        return Data_Ordering.LT.value;
                    };
                    if (y instanceof StartDuration) {
                        return Data_Ordering.GT.value;
                    };
                    if (x instanceof DurationOnly && y instanceof DurationOnly) {
                        return Data_Ord.compare(dictOrd)(x.value0)(y.value0);
                    };
                    throw new Error("Failed pattern match at Data.Interval (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
                };
            },
            Eq0: function () {
                return eqInterval2;
            }
        };
    };
};
var ordRecurringInterval = function (dictOrd) {
    var ordInterval1 = ordInterval(dictOrd);
    var eqRecurringInterval1 = eqRecurringInterval(dictOrd.Eq0());
    return function (dictOrd1) {
        var ordInterval2 = ordInterval1(dictOrd1);
        var eqRecurringInterval2 = eqRecurringInterval1(dictOrd1.Eq0());
        return {
            compare: function (x) {
                return function (y) {
                    var v = Data_Ord.compare(ordMaybe)(x.value0)(y.value0);
                    if (v instanceof Data_Ordering.LT) {
                        return Data_Ordering.LT.value;
                    };
                    if (v instanceof Data_Ordering.GT) {
                        return Data_Ordering.GT.value;
                    };
                    return Data_Ord.compare(ordInterval2)(x.value1)(y.value1);
                };
            },
            Eq0: function () {
                return eqRecurringInterval2;
            }
        };
    };
};
var bifunctorInterval = {
    bimap: function (v) {
        return function (v1) {
            return function (v2) {
                if (v2 instanceof StartEnd) {
                    return new StartEnd(v1(v2.value0), v1(v2.value1));
                };
                if (v2 instanceof DurationEnd) {
                    return new DurationEnd(v(v2.value0), v1(v2.value1));
                };
                if (v2 instanceof StartDuration) {
                    return new StartDuration(v1(v2.value0), v(v2.value1));
                };
                if (v2 instanceof DurationOnly) {
                    return new DurationOnly(v(v2.value0));
                };
                throw new Error("Failed pattern match at Data.Interval (line 75, column 1 - line 79, column 50): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
            };
        };
    }
};
var bifunctorRecurringInterval = {
    bimap: function (f) {
        return function (g) {
            return function (v) {
                return new RecurringInterval(v.value0, Data_Bifunctor.bimap(bifunctorInterval)(f)(g)(v.value1));
            };
        };
    }
};
var functorInterval = {
    map: /* #__PURE__ */ Data_Bifunctor.bimap(bifunctorInterval)(/* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn))
};
var extendInterval = {
    extend: function (v) {
        return function (v1) {
            if (v1 instanceof StartEnd) {
                return new StartEnd(v(v1), v(v1));
            };
            if (v1 instanceof DurationEnd) {
                return new DurationEnd(v1.value0, v(v1));
            };
            if (v1 instanceof StartDuration) {
                return new StartDuration(v(v1), v1.value1);
            };
            if (v1 instanceof DurationOnly) {
                return new DurationOnly(v1.value0);
            };
            throw new Error("Failed pattern match at Data.Interval (line 111, column 1 - line 115, column 45): " + [ v.constructor.name, v1.constructor.name ]);
        };
    },
    Functor0: function () {
        return functorInterval;
    }
};
var functorRecurringInterval = {
    map: function (f) {
        return function (v) {
            return new RecurringInterval(v.value0, Data_Functor.map(functorInterval)(f)(v.value1));
        };
    }
};
var extendRecurringInterval = {
    extend: function (f) {
        return function (v) {
            return new RecurringInterval(v.value0, Control_Extend.extend(extendInterval)(Data_Function["const"](f(v)))(v.value1));
        };
    },
    Functor0: function () {
        return functorRecurringInterval;
    }
};
var traversableInterval = {
    traverse: function (dictApplicative) {
        var Apply0 = dictApplicative.Apply0();
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (v) {
            return function (v1) {
                if (v1 instanceof StartEnd) {
                    return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(StartEnd.create)(v(v1.value0)))(v(v1.value1));
                };
                if (v1 instanceof DurationEnd) {
                    return Data_Functor.mapFlipped(Functor0)(v(v1.value1))(DurationEnd.create(v1.value0));
                };
                if (v1 instanceof StartDuration) {
                    return Data_Functor.mapFlipped(Functor0)(v(v1.value0))(function (v2) {
                        return new StartDuration(v2, v1.value1);
                    });
                };
                if (v1 instanceof DurationOnly) {
                    return Control_Applicative.pure(dictApplicative)(new DurationOnly(v1.value0));
                };
                throw new Error("Failed pattern match at Data.Interval (line 97, column 1 - line 102, column 29): " + [ v.constructor.name, v1.constructor.name ]);
            };
        };
    },
    sequence: function (dictApplicative) {
        return Data_Traversable.sequenceDefault(traversableInterval)(dictApplicative);
    },
    Functor0: function () {
        return functorInterval;
    },
    Foldable1: function () {
        return foldableInterval;
    }
};
var traversableRecurringInterval = {
    traverse: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (f) {
            return function (i) {
                return over(Functor0)(Data_Traversable.traverse(traversableInterval)(dictApplicative)(f))(i);
            };
        };
    },
    sequence: function (dictApplicative) {
        return Data_Traversable.sequenceDefault(traversableRecurringInterval)(dictApplicative);
    },
    Functor0: function () {
        return functorRecurringInterval;
    },
    Foldable1: function () {
        return foldableRecurringInterval;
    }
};
var bifoldableInterval = {
    bifoldl: function (v) {
        return function (v1) {
            return function (v2) {
                return function (v3) {
                    if (v3 instanceof StartEnd) {
                        return v1(v1(v2)(v3.value0))(v3.value1);
                    };
                    if (v3 instanceof DurationEnd) {
                        return v1(v(v2)(v3.value0))(v3.value1);
                    };
                    if (v3 instanceof StartDuration) {
                        return v1(v(v2)(v3.value1))(v3.value0);
                    };
                    if (v3 instanceof DurationOnly) {
                        return v(v2)(v3.value0);
                    };
                    throw new Error("Failed pattern match at Data.Interval (line 89, column 1 - line 95, column 32): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name, v3.constructor.name ]);
                };
            };
        };
    },
    bifoldr: function (x) {
        return Data_Bifoldable.bifoldrDefault(bifoldableInterval)(x);
    },
    bifoldMap: function (dictMonoid) {
        return Data_Bifoldable.bifoldMapDefaultL(bifoldableInterval)(dictMonoid);
    }
};
var bifoldableRecurringInterval = {
    bifoldl: function (f) {
        return function (g) {
            return function (i) {
                var $279 = Data_Bifoldable.bifoldl(bifoldableInterval)(f)(g)(i);
                return function ($280) {
                    return $279(interval($280));
                };
            };
        };
    },
    bifoldr: function (f) {
        return function (g) {
            return function (i) {
                var $281 = Data_Bifoldable.bifoldr(bifoldableInterval)(f)(g)(i);
                return function ($282) {
                    return $281(interval($282));
                };
            };
        };
    },
    bifoldMap: function (dictMonoid) {
        return Data_Bifoldable.bifoldMapDefaultL(bifoldableRecurringInterval)(dictMonoid);
    }
};
var bitraversableInterval = {
    bitraverse: function (dictApplicative) {
        var Apply0 = dictApplicative.Apply0();
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (v) {
            return function (v1) {
                return function (v2) {
                    if (v2 instanceof StartEnd) {
                        return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(StartEnd.create)(v1(v2.value0)))(v1(v2.value1));
                    };
                    if (v2 instanceof DurationEnd) {
                        return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(DurationEnd.create)(v(v2.value0)))(v1(v2.value1));
                    };
                    if (v2 instanceof StartDuration) {
                        return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(StartDuration.create)(v1(v2.value0)))(v(v2.value1));
                    };
                    if (v2 instanceof DurationOnly) {
                        return Data_Functor.map(Functor0)(DurationOnly.create)(v(v2.value0));
                    };
                    throw new Error("Failed pattern match at Data.Interval (line 104, column 1 - line 109, column 33): " + [ v.constructor.name, v1.constructor.name, v2.constructor.name ]);
                };
            };
        };
    },
    bisequence: function (dictApplicative) {
        return Data_Bitraversable.bisequenceDefault(bitraversableInterval)(dictApplicative);
    },
    Bifunctor0: function () {
        return bifunctorInterval;
    },
    Bifoldable1: function () {
        return bifoldableInterval;
    }
};
var bitraversableRecurringInterval = {
    bitraverse: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (l) {
            return function (r) {
                return function (i) {
                    return over(Functor0)(Data_Bitraversable.bitraverse(bitraversableInterval)(dictApplicative)(l)(r))(i);
                };
            };
        };
    },
    bisequence: function (dictApplicative) {
        return Data_Bitraversable.bisequenceDefault(bitraversableRecurringInterval)(dictApplicative);
    },
    Bifunctor0: function () {
        return bifunctorRecurringInterval;
    },
    Bifoldable1: function () {
        return bifoldableRecurringInterval;
    }
};
export {
    StartEnd,
    DurationEnd,
    StartDuration,
    DurationOnly,
    RecurringInterval,
    eqRecurringInterval,
    ordRecurringInterval,
    showRecurringInterval,
    functorRecurringInterval,
    bifunctorRecurringInterval,
    foldableRecurringInterval,
    bifoldableRecurringInterval,
    traversableRecurringInterval,
    bitraversableRecurringInterval,
    extendRecurringInterval,
    eqInterval,
    ordInterval,
    showInterval,
    functorInterval,
    bifunctorInterval,
    foldableInterval,
    bifoldableInterval,
    traversableInterval,
    bitraversableInterval,
    extendInterval
};
export {
    Duration,
    Day,
    Hour,
    Minute,
    Month,
    Second,
    Week,
    Year,
    day,
    hour,
    millisecond,
    minute,
    month,
    second,
    week,
    year
} from "../Data.Interval.Duration/index.js";
