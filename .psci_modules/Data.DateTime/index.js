import * as $foreign from "./foreign.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Data_Bounded from "../Data.Bounded/index.js";
import * as Data_Date from "../Data.Date/index.js";
import * as Data_Date_Component from "../Data.Date.Component/index.js";
import * as Data_Enum from "../Data.Enum/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Time from "../Data.Time/index.js";
import * as Data_Time_Component from "../Data.Time.Component/index.js";
import * as Data_Time_Duration from "../Data.Time.Duration/index.js";

// | A date/time value in the Gregorian calendar/UTC time zone.
var DateTime = /* #__PURE__ */ (function () {
    function DateTime(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    DateTime.create = function (value0) {
        return function (value1) {
            return new DateTime(value0, value1);
        };
    };
    return DateTime;
})();
var toRecord = function (v) {
    return {
        year: Data_Enum.fromEnum(Data_Date_Component.boundedEnumYear)(Data_Date.year(v.value0)),
        month: Data_Enum.fromEnum(Data_Date_Component.boundedEnumMonth)(Data_Date.month(v.value0)),
        day: Data_Enum.fromEnum(Data_Date_Component.boundedEnumDay)(Data_Date.day(v.value0)),
        hour: Data_Enum.fromEnum(Data_Time_Component.boundedEnumHour)(Data_Time.hour(v.value1)),
        minute: Data_Enum.fromEnum(Data_Time_Component.boundedEnumMinute)(Data_Time.minute(v.value1)),
        second: Data_Enum.fromEnum(Data_Time_Component.boundedEnumSecond)(Data_Time.second(v.value1)),
        millisecond: Data_Enum.fromEnum(Data_Time_Component.boundedEnumMillisecond)(Data_Time.millisecond(v.value1))
    };
};
var time = function (v) {
    return v.value1;
};
var showDateTime = {
    show: function (v) {
        return "(DateTime " + (Data_Show.show(Data_Date.showDate)(v.value0) + (" " + (Data_Show.show(Data_Time.showTime)(v.value1) + ")")));
    }
};
var modifyTimeF = function (dictFunctor) {
    return function (f) {
        return function (v) {
            return Data_Functor.map(dictFunctor)(DateTime.create(v.value0))(f(v.value1));
        };
    };
};
var modifyTime = function (f) {
    return function (v) {
        return new DateTime(v.value0, f(v.value1));
    };
};
var modifyDateF = function (dictFunctor) {
    return function (f) {
        return function (v) {
            return Data_Functor.map(dictFunctor)(Data_Function.flip(DateTime.create)(v.value1))(f(v.value0));
        };
    };
};
var modifyDate = function (f) {
    return function (v) {
        return new DateTime(f(v.value0), v.value1);
    };
};
var eqDateTime = {
    eq: function (x) {
        return function (y) {
            return Data_Eq.eq(Data_Date.eqDate)(x.value0)(y.value0) && Data_Eq.eq(Data_Time.eqTime)(x.value1)(y.value1);
        };
    }
};
var ordDateTime = {
    compare: function (x) {
        return function (y) {
            var v = Data_Ord.compare(Data_Date.ordDate)(x.value0)(y.value0);
            if (v instanceof Data_Ordering.LT) {
                return Data_Ordering.LT.value;
            };
            if (v instanceof Data_Ordering.GT) {
                return Data_Ordering.GT.value;
            };
            return Data_Ord.compare(Data_Time.ordTime)(x.value1)(y.value1);
        };
    },
    Eq0: function () {
        return eqDateTime;
    }
};

// | Calculates the difference between two date/time values, returning the
// | result as a duration.
var diff = function (dictDuration) {
    var toDuration = Data_Time_Duration.toDuration(dictDuration);
    return function (dt1) {
        return function (dt2) {
            return toDuration($foreign.calcDiff(toRecord(dt1), toRecord(dt2)));
        };
    };
};
var date = function (v) {
    return v.value0;
};
var boundedDateTime = /* #__PURE__ */ (function () {
    return {
        bottom: new DateTime(Data_Bounded.bottom(Data_Date.boundedDate), Data_Bounded.bottom(Data_Time.boundedTime)),
        top: new DateTime(Data_Bounded.top(Data_Date.boundedDate), Data_Bounded.top(Data_Time.boundedTime)),
        Ord0: function () {
            return ordDateTime;
        }
    };
})();

// | Adjusts a date/time value with a duration offset. `Nothing` is returned
// | if the resulting date would be outside of the range of valid dates.
var adjust = function (dictDuration) {
    return function (d) {
        return function (dt) {
            return Control_Bind.bind(Data_Maybe.bindMaybe)($foreign.adjustImpl(Data_Maybe.Just.create)(Data_Maybe.Nothing.value)(Data_Time_Duration.fromDuration(dictDuration)(d))(toRecord(dt)))(function (rec) {
                return Control_Apply.apply(Data_Maybe.applyMaybe)(Data_Functor.map(Data_Maybe.functorMaybe)(DateTime.create)(Control_Bind.join(Data_Maybe.bindMaybe)(Control_Apply.apply(Data_Maybe.applyMaybe)(Control_Apply.apply(Data_Maybe.applyMaybe)(Data_Functor.map(Data_Maybe.functorMaybe)(Data_Date.exactDate)(Data_Enum.toEnum(Data_Date_Component.boundedEnumYear)(rec.year)))(Data_Enum.toEnum(Data_Date_Component.boundedEnumMonth)(rec.month)))(Data_Enum.toEnum(Data_Date_Component.boundedEnumDay)(rec.day)))))(Control_Apply.apply(Data_Maybe.applyMaybe)(Control_Apply.apply(Data_Maybe.applyMaybe)(Control_Apply.apply(Data_Maybe.applyMaybe)(Data_Functor.map(Data_Maybe.functorMaybe)(Data_Time.Time.create)(Data_Enum.toEnum(Data_Time_Component.boundedEnumHour)(rec.hour)))(Data_Enum.toEnum(Data_Time_Component.boundedEnumMinute)(rec.minute)))(Data_Enum.toEnum(Data_Time_Component.boundedEnumSecond)(rec.second)))(Data_Enum.toEnum(Data_Time_Component.boundedEnumMillisecond)(rec.millisecond)));
            });
        };
    };
};
export {
    DateTime,
    date,
    modifyDate,
    modifyDateF,
    time,
    modifyTime,
    modifyTimeF,
    adjust,
    diff,
    eqDateTime,
    ordDateTime,
    boundedDateTime,
    showDateTime
};
export {
    April,
    August,
    December,
    February,
    January,
    July,
    June,
    March,
    May,
    November,
    October,
    September,
    Friday,
    Monday,
    Saturday,
    Sunday,
    Thursday,
    Tuesday,
    Wednesday,
    canonicalDate,
    day,
    exactDate,
    month,
    weekday,
    year
} from "../Data.Date/index.js";
export {
    Time,
    hour,
    millisecond,
    minute,
    second,
    setHour,
    setMillisecond,
    setMinute,
    setSecond
} from "../Data.Time/index.js";
