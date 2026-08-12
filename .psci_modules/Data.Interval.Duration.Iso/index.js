import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Plus from "../Control.Plus/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_HeytingAlgebra from "../Data.HeytingAlgebra/index.js";
import * as Data_Interval_Duration from "../Data.Interval.Duration/index.js";
import * as Data_List from "../Data.List/index.js";
import * as Data_List_NonEmpty from "../Data.List.NonEmpty/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Monoid_Additive from "../Data.Monoid.Additive/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Number from "../Data.Number/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var empty = /* #__PURE__ */ Control_Plus.empty(Data_List_Types.plusList);
var foldMap = /* #__PURE__ */ Data_Foldable.foldMap(Data_List_Types.foldableList)(Data_List_Types.monoidList);
var monoidAdditive = /* #__PURE__ */ Data_Monoid_Additive.monoidAdditive(Data_Semiring.semiringNumber);
var heytingAlgebraFunction = /* #__PURE__ */ Data_HeytingAlgebra.heytingAlgebraFunction(Data_HeytingAlgebra.heytingAlgebraBoolean);
var monoidFn = /* #__PURE__ */ Data_Monoid.monoidFn(Data_List_Types.monoidList);
var IsoDuration = function (x) {
    return x;
};
var IsEmpty = /* #__PURE__ */ (function () {
    function IsEmpty() {

    };
    IsEmpty.value = new IsEmpty();
    return IsEmpty;
})();
var InvalidWeekComponentUsage = /* #__PURE__ */ (function () {
    function InvalidWeekComponentUsage() {

    };
    InvalidWeekComponentUsage.value = new InvalidWeekComponentUsage();
    return InvalidWeekComponentUsage;
})();
var ContainsNegativeValue = /* #__PURE__ */ (function () {
    function ContainsNegativeValue(value0) {
        this.value0 = value0;
    };
    ContainsNegativeValue.create = function (value0) {
        return new ContainsNegativeValue(value0);
    };
    return ContainsNegativeValue;
})();
var InvalidFractionalUse = /* #__PURE__ */ (function () {
    function InvalidFractionalUse(value0) {
        this.value0 = value0;
    };
    InvalidFractionalUse.create = function (value0) {
        return new InvalidFractionalUse(value0);
    };
    return InvalidFractionalUse;
})();
var unIsoDuration = function (v) {
    return v;
};
var showIsoDuration = {
    show: function (v) {
        return "(IsoDuration " + (Data_Show.show(Data_Interval_Duration.showDuration)(v) + ")");
    }
};
var showError = {
    show: function (v) {
        if (v instanceof IsEmpty) {
            return "(IsEmpty)";
        };
        if (v instanceof InvalidWeekComponentUsage) {
            return "(InvalidWeekComponentUsage)";
        };
        if (v instanceof ContainsNegativeValue) {
            return "(ContainsNegativeValue " + (Data_Show.show(Data_Interval_Duration.showDurationComponent)(v.value0) + ")");
        };
        if (v instanceof InvalidFractionalUse) {
            return "(InvalidFractionalUse " + (Data_Show.show(Data_Interval_Duration.showDurationComponent)(v.value0) + ")");
        };
        throw new Error("Failed pattern match at Data.Interval.Duration.Iso (line 43, column 1 - line 47, column 76): " + [ v.constructor.name ]);
    }
};
var prettyError = function (v) {
    if (v instanceof IsEmpty) {
        return "Duration is empty (has no components)";
    };
    if (v instanceof InvalidWeekComponentUsage) {
        return "Week component of Duration is used with other components";
    };
    if (v instanceof ContainsNegativeValue) {
        return "Component `" + (Data_Show.show(Data_Interval_Duration.showDurationComponent)(v.value0) + "` contains negative value");
    };
    if (v instanceof InvalidFractionalUse) {
        return "Invalid usage of Fractional value at component `" + (Data_Show.show(Data_Interval_Duration.showDurationComponent)(v.value0) + "`");
    };
    throw new Error("Failed pattern match at Data.Interval.Duration.Iso (line 49, column 1 - line 49, column 31): " + [ v.constructor.name ]);
};
var eqIsoDuration = {
    eq: function (x) {
        return function (y) {
            return Data_Eq.eq(Data_Interval_Duration.eqDuration)(x)(y);
        };
    }
};
var ordIsoDuration = {
    compare: function (x) {
        return function (y) {
            return Data_Ord.compare(Data_Interval_Duration.ordDuration)(x)(y);
        };
    },
    Eq0: function () {
        return eqIsoDuration;
    }
};
var eqError = {
    eq: function (x) {
        return function (y) {
            if (x instanceof IsEmpty && y instanceof IsEmpty) {
                return true;
            };
            if (x instanceof InvalidWeekComponentUsage && y instanceof InvalidWeekComponentUsage) {
                return true;
            };
            if (x instanceof ContainsNegativeValue && y instanceof ContainsNegativeValue) {
                return Data_Eq.eq(Data_Interval_Duration.eqDurationComponent)(x.value0)(y.value0);
            };
            if (x instanceof InvalidFractionalUse && y instanceof InvalidFractionalUse) {
                return Data_Eq.eq(Data_Interval_Duration.eqDurationComponent)(x.value0)(y.value0);
            };
            return false;
        };
    }
};
var ordError = {
    compare: function (x) {
        return function (y) {
            if (x instanceof IsEmpty && y instanceof IsEmpty) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof IsEmpty) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof IsEmpty) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof InvalidWeekComponentUsage && y instanceof InvalidWeekComponentUsage) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof InvalidWeekComponentUsage) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof InvalidWeekComponentUsage) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof ContainsNegativeValue && y instanceof ContainsNegativeValue) {
                return Data_Ord.compare(Data_Interval_Duration.ordDurationComponent)(x.value0)(y.value0);
            };
            if (x instanceof ContainsNegativeValue) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof ContainsNegativeValue) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof InvalidFractionalUse && y instanceof InvalidFractionalUse) {
                return Data_Ord.compare(Data_Interval_Duration.ordDurationComponent)(x.value0)(y.value0);
            };
            throw new Error("Failed pattern match at Data.Interval.Duration.Iso (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqError;
    }
};
var checkWeekUsage = function (v) {
    var $66 = Data_Maybe.isJust(Data_Map_Internal.lookup(Data_Interval_Duration.ordDurationComponent)(Data_Interval_Duration.Week.value)(v.asMap)) && Data_Map_Internal.size(v.asMap) > 1;
    if ($66) {
        return Control_Applicative.pure(Data_List_Types.applicativeList)(InvalidWeekComponentUsage.value);
    };
    return empty;
};
var checkNegativeValues = function (v) {
    return Data_Function.flip(foldMap)(v.asList)(function (v1) {
        var $70 = v1.value1 >= 0.0;
        if ($70) {
            return Control_Plus.empty(Data_List_Types.plusList);
        };
        return Control_Applicative.pure(Data_List_Types.applicativeList)(new ContainsNegativeValue(v1.value0));
    });
};
var checkFractionalUse = function (v) {
    var isFractional = function (a) {
        return Data_Number.floor(a) !== a;
    };
    var checkRest = function (rest) {
        return Data_Newtype.unwrap()(Data_Foldable.foldMap(Data_List_Types.foldableList)(monoidAdditive)(function ($87) {
            return Data_Monoid_Additive.Additive(Data_Number.abs(Data_Tuple.snd($87)));
        })(rest)) > 0.0;
    };
    var v1 = (function (v2) {
        return v2.rest;
    })(Data_List.span((function () {
        var $88 = Data_HeytingAlgebra.not(heytingAlgebraFunction)(isFractional);
        return function ($89) {
            return $88(Data_Tuple.snd($89));
        };
    })())(v.asList));
    if (v1 instanceof Data_List_Types.Cons && checkRest(v1.value1)) {
        return Control_Applicative.pure(Data_List_Types.applicativeList)(new InvalidFractionalUse(v1.value0.value0));
    };
    return empty;
};
var checkEmptiness = function (v) {
    var $82 = Data_List["null"](v.asList);
    if ($82) {
        return Control_Applicative.pure(Data_List_Types.applicativeList)(IsEmpty.value);
    };
    return empty;
};
var checkValidIsoDuration = function (v) {
    var check = Data_Foldable.fold(Data_Foldable.foldableArray)(monoidFn)([ checkWeekUsage, checkEmptiness, checkFractionalUse, checkNegativeValues ]);
    var asList = Data_List.reverse(Data_Map_Internal.toUnfoldable(Data_List_Types.unfoldableList)(v));
    return check({
        asList: asList,
        asMap: v
    });
};
var mkIsoDuration = function (d) {
    var v = Data_List_NonEmpty.fromList(checkValidIsoDuration(d));
    if (v instanceof Data_Maybe.Just) {
        return new Data_Either.Left(v.value0);
    };
    if (v instanceof Data_Maybe.Nothing) {
        return new Data_Either.Right(d);
    };
    throw new Error("Failed pattern match at Data.Interval.Duration.Iso (line 60, column 19 - line 62, column 35): " + [ v.constructor.name ]);
};
export {
    unIsoDuration,
    mkIsoDuration,
    IsEmpty,
    InvalidWeekComponentUsage,
    ContainsNegativeValue,
    InvalidFractionalUse,
    prettyError,
    eqIsoDuration,
    ordIsoDuration,
    showIsoDuration,
    eqError,
    ordError,
    showError
};
