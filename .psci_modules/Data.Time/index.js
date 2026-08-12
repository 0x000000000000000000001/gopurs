import * as Control_Apply from "../Control.Apply/index.js";
import * as Data_Bounded from "../Data.Bounded/index.js";
import * as Data_Enum from "../Data.Enum/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Int from "../Data.Int/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Number from "../Data.Number/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Time_Component from "../Data.Time.Component/index.js";
import * as Data_Time_Duration from "../Data.Time.Duration/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var Time = /* #__PURE__ */ (function () {
    function Time(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    Time.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new Time(value0, value1, value2, value3);
                };
            };
        };
    };
    return Time;
})();
var showTime = {
    show: function (v) {
        return "(Time " + (Data_Show.show(Data_Time_Component.showHour)(v.value0) + (" " + (Data_Show.show(Data_Time_Component.showMinute)(v.value1) + (" " + (Data_Show.show(Data_Time_Component.showSecond)(v.value2) + (" " + (Data_Show.show(Data_Time_Component.showMillisecond)(v.value3) + ")")))))));
    }
};

// | Alters the second component of a time value.
var setSecond = function (s) {
    return function (v) {
        return new Time(v.value0, v.value1, s, v.value3);
    };
};

// | Alters the minute component of a time value.
var setMinute = function (m) {
    return function (v) {
        return new Time(v.value0, m, v.value2, v.value3);
    };
};

// | Alters the millisecond component of a time value.
var setMillisecond = function (ms) {
    return function (v) {
        return new Time(v.value0, v.value1, v.value2, ms);
    };
};

// | Alters the hour component of a time value.
var setHour = function (h) {
    return function (v) {
        return new Time(h, v.value1, v.value2, v.value3);
    };
};

// | The second component of a time value.
var second = function (v) {
    return v.value2;
};

// | The minute component of a time value.
var minute = function (v) {
    return v.value1;
};

// | The millisecond component of a time value.
var millisecond = function (v) {
    return v.value3;
};
var millisToTime = function (v) {
    var hours = Data_Number.floor(v / 3600000.0);
    var minutes = Data_Number.floor((v - hours * 3600000.0) / 60000.0);
    var seconds = Data_Number.floor((v - (hours * 3600000.0 + minutes * 60000.0)) / 1000.0);
    var milliseconds = v - (hours * 3600000.0 + minutes * 60000.0 + seconds * 1000.0);
    return Data_Maybe.fromJust()(Control_Apply.apply(Data_Maybe.applyMaybe)(Control_Apply.apply(Data_Maybe.applyMaybe)(Control_Apply.apply(Data_Maybe.applyMaybe)(Data_Functor.map(Data_Maybe.functorMaybe)(Time.create)(Data_Enum.toEnum(Data_Time_Component.boundedEnumHour)(Data_Int.floor(hours))))(Data_Enum.toEnum(Data_Time_Component.boundedEnumMinute)(Data_Int.floor(minutes))))(Data_Enum.toEnum(Data_Time_Component.boundedEnumSecond)(Data_Int.floor(seconds))))(Data_Enum.toEnum(Data_Time_Component.boundedEnumMillisecond)(Data_Int.floor(milliseconds))));
};

// | The hour component of a time value.
var hour = function (v) {
    return v.value0;
};
var timeToMillis = function (t) {
    return 3600000.0 * Data_Int.toNumber(Data_Enum.fromEnum(Data_Time_Component.boundedEnumHour)(hour(t))) + 60000.0 * Data_Int.toNumber(Data_Enum.fromEnum(Data_Time_Component.boundedEnumMinute)(minute(t))) + 1000.0 * Data_Int.toNumber(Data_Enum.fromEnum(Data_Time_Component.boundedEnumSecond)(second(t))) + Data_Int.toNumber(Data_Enum.fromEnum(Data_Time_Component.boundedEnumMillisecond)(millisecond(t)));
};
var eqTime = {
    eq: function (x) {
        return function (y) {
            return Data_Eq.eq(Data_Time_Component.eqHour)(x.value0)(y.value0) && Data_Eq.eq(Data_Time_Component.eqMinute)(x.value1)(y.value1) && Data_Eq.eq(Data_Time_Component.eqSecond)(x.value2)(y.value2) && Data_Eq.eq(Data_Time_Component.eqMillisecond)(x.value3)(y.value3);
        };
    }
};
var ordTime = {
    compare: function (x) {
        return function (y) {
            var v = Data_Ord.compare(Data_Time_Component.ordHour)(x.value0)(y.value0);
            if (v instanceof Data_Ordering.LT) {
                return Data_Ordering.LT.value;
            };
            if (v instanceof Data_Ordering.GT) {
                return Data_Ordering.GT.value;
            };
            var v1 = Data_Ord.compare(Data_Time_Component.ordMinute)(x.value1)(y.value1);
            if (v1 instanceof Data_Ordering.LT) {
                return Data_Ordering.LT.value;
            };
            if (v1 instanceof Data_Ordering.GT) {
                return Data_Ordering.GT.value;
            };
            var v2 = Data_Ord.compare(Data_Time_Component.ordSecond)(x.value2)(y.value2);
            if (v2 instanceof Data_Ordering.LT) {
                return Data_Ordering.LT.value;
            };
            if (v2 instanceof Data_Ordering.GT) {
                return Data_Ordering.GT.value;
            };
            return Data_Ord.compare(Data_Time_Component.ordMillisecond)(x.value3)(y.value3);
        };
    },
    Eq0: function () {
        return eqTime;
    }
};

// | Calculates the difference between two times, returning the result as a
// | duration.
var diff = function (dictDuration) {
    return function (t1) {
        return function (t2) {
            return Data_Time_Duration.toDuration(dictDuration)(Data_Semigroup.append(Data_Time_Duration.semigroupMilliseconds)(timeToMillis(t1))(Data_Time_Duration.negateDuration(Data_Time_Duration.durationMilliseconds)(timeToMillis(t2))));
        };
    };
};
var boundedTime = /* #__PURE__ */ (function () {
    return {
        bottom: new Time(Data_Bounded.bottom(Data_Time_Component.boundedHour), Data_Bounded.bottom(Data_Time_Component.boundedMinute), Data_Bounded.bottom(Data_Time_Component.boundedSecond), Data_Bounded.bottom(Data_Time_Component.boundedMillisecond)),
        top: new Time(Data_Bounded.top(Data_Time_Component.boundedHour), Data_Bounded.top(Data_Time_Component.boundedMinute), Data_Bounded.top(Data_Time_Component.boundedSecond), Data_Bounded.top(Data_Time_Component.boundedMillisecond)),
        Ord0: function () {
            return ordTime;
        }
    };
})();
var maxTime = /* #__PURE__ */ timeToMillis(/* #__PURE__ */ Data_Bounded.top(boundedTime));
var minTime = /* #__PURE__ */ timeToMillis(/* #__PURE__ */ Data_Bounded.bottom(boundedTime));

// | Adjusts a time value with a duration offset. The result includes a
// | remainder value of the whole number of days involved in the adjustment,
// | for example, if a time of 23:00:00:00 has a duration of +2 hours added to
// | it, the result will be 1 day, and 01:00:00:00. Correspondingly, if the
// | duration is negative, a negative number of days may also be returned as
// | the remainder.
var adjust = function (dictDuration) {
    return function (d) {
        return function (t) {
            var tLength = timeToMillis(t);
            var d$prime = Data_Time_Duration.fromDuration(dictDuration)(d);
            var wholeDays = Data_Number.floor(Data_Newtype.unwrap()(d$prime) / 8.64e7);
            var msAdjust = Data_Semigroup.append(Data_Time_Duration.semigroupMilliseconds)(d$prime)(Data_Time_Duration.negateDuration(Data_Time_Duration.durationMilliseconds)(Data_Time_Duration.fromDuration(Data_Time_Duration.durationDays)(wholeDays)));
            var msAdjusted = Data_Semigroup.append(Data_Time_Duration.semigroupMilliseconds)(tLength)(msAdjust);
            var wrap = (function () {
                var $118 = Data_Ord.greaterThan(Data_Time_Duration.ordMilliseconds)(msAdjusted)(maxTime);
                if ($118) {
                    return 1.0;
                };
                var $119 = Data_Ord.lessThan(Data_Time_Duration.ordMilliseconds)(msAdjusted)(minTime);
                if ($119) {
                    return -1.0;
                };
                return 0.0;
            })();
            return new Data_Tuple.Tuple(Data_Semigroup.append(Data_Time_Duration.semigroupDays)(wholeDays)(wrap), millisToTime(Data_Semigroup.append(Data_Time_Duration.semigroupMilliseconds)(msAdjusted)(8.64e7 * -wrap)));
        };
    };
};
export {
    Time,
    hour,
    setHour,
    minute,
    setMinute,
    second,
    setSecond,
    millisecond,
    setMillisecond,
    adjust,
    diff,
    eqTime,
    ordTime,
    boundedTime,
    showTime
};
