// | A module providing a type and operations for the native JavaScript `Date`
// | object.
// |
// | The `JSDate` type and associated functions are provided for interop
// | purposes with JavaScript, but for working with dates in PureScript it is
// | recommended that `DateTime` representation is used instead - `DateTime`
// | offers greater type safety, a more PureScript-friendly interface, and has
// | a `Generic` instance.
import * as $foreign from "./foreign.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Data_Date from "../Data.Date/index.js";
import * as Data_Date_Component from "../Data.Date.Component/index.js";
import * as Data_DateTime from "../Data.DateTime/index.js";
import * as Data_DateTime_Instant from "../Data.DateTime.Instant/index.js";
import * as Data_Enum from "../Data.Enum/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Int from "../Data.Int/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Time from "../Data.Time/index.js";
import * as Data_Time_Component from "../Data.Time.Component/index.js";
import * as Data_Time_Duration from "../Data.Time.Duration/index.js";
import * as Foreign from "../Foreign/index.js";

// | Returns the date as a string using the UTC timezone.
var toUTCString = function (dt) {
    return $foreign.dateMethod("toUTCString", dt);
};

// | Returns the time portion of a date value as a human-readable string.
var toTimeString = function (dt) {
    return $foreign.dateMethod("toTimeString", dt);
};

// | Returns a string representing for a date value.
var toString = function (dt) {
    return $foreign.dateMethod("toString", dt);
};

// | Attempts to construct an `Instant` for a `JSDate`. `Nothing` is returned
// | only when the date value is an invalid date.
var toInstant = /* #__PURE__ */ (function () {
    return Control_Bind.composeKleisliFlipped(Data_Maybe.bindMaybe)(function ($4) {
        return Data_DateTime_Instant.instant(Data_Time_Duration.Milliseconds($4));
    })($foreign.toInstantImpl(Data_Maybe.Just.create)(Data_Maybe.Nothing.value));
})();

// | Converts a date value to an ISO 8601 Extended format date string.
var toISOString = function (dt) {
    return $foreign.dateMethodEff("toISOString", dt);
};

// | Attempts to construct a `DateTime` value for a `JSDate`. `Nothing` is
// | returned only when the date value is an invalid date.
var toDateTime = /* #__PURE__ */ Data_Functor.map(Data_Functor.functorFn)(/* #__PURE__ */ Data_Functor.map(Data_Maybe.functorMaybe)(Data_DateTime_Instant.toDateTime))(toInstant);

// | Returns the date portion of a date value as a human-readable string.
var toDateString = function (dt) {
    return $foreign.dateMethod("toDateString", dt);
};

// | Attempts to construct a `Date` value for a `JSDate`, ignoring any time
// | component of the `JSDate`. `Nothing` is returned only when the date value
// | is an invalid date.
var toDate = /* #__PURE__ */ Data_Functor.map(Data_Functor.functorFn)(/* #__PURE__ */ Data_Functor.map(Data_Maybe.functorMaybe)(Data_DateTime.date))(toDateTime);

// | Attempts to read a `Foreign` value as a `JSDate`.
var readDate = /* #__PURE__ */ Foreign.unsafeReadTagged(Data_Identity.monadIdentity)("Date");

// | Returns the seconds for a date, according to UTC.
var getUTCSeconds = function (dt) {
    return $foreign.dateMethod("getUTCSeconds", dt);
};

// | Returns the month for a date, according to UTC.
var getUTCMonth = function (dt) {
    return $foreign.dateMethod("getUTCMonth", dt);
};

// | Returns the minutes for a date, according to UTC.
var getUTCMinutes = function (dt) {
    return $foreign.dateMethod("getUTCMinutes", dt);
};

// | Returns the milliseconds for a date, according to UTC.
var getUTCMilliseconds = function (dt) {
    return $foreign.dateMethod("getUTCMilliseconds", dt);
};

// | Returns the hours for a date, according to UTC.
var getUTCHours = function (dt) {
    return $foreign.dateMethod("getUTCHours", dt);
};

// | Returns the year for a date, according to UTC.
var getUTCFullYear = function (dt) {
    return $foreign.dateMethod("getUTCFullYear", dt);
};

// | Returns the day of the week for a date, according to UTC.
var getUTCDay = function (dt) {
    return $foreign.dateMethod("getUTCDay", dt);
};

// | Returns the day of the month for a date, according to UTC.
var getUTCDate = function (dt) {
    return $foreign.dateMethod("getUTCDate", dt);
};

// | Returns the time-zone offset for a date, according to the current machine's
// | date/time locale.
var getTimezoneOffset = function (dt) {
    return $foreign.dateMethodEff("getTimezoneOffset", dt);
};

// | Returns the date as a number of milliseconds since 1970-01-01 00:00:00 UTC.
var getTime = function (dt) {
    return $foreign.dateMethod("getTime", dt);
};
var showJSDate = {
    show: function (a) {
        return "(fromTime " + (Data_Show.show(Data_Show.showNumber)(getTime(a)) + ")");
    }
};

// | Returns the seconds for a date, according to the current machine's
// | date/time locale.
var getSeconds = function (dt) {
    return $foreign.dateMethodEff("getSeconds", dt);
};

// | Returns the month for a date, according to the current machine's
// | date/time locale.
var getMonth = function (dt) {
    return $foreign.dateMethodEff("getMonth", dt);
};

// | Returns the minutes for a date, according to the current machine's
// | date/time locale.
var getMinutes = function (dt) {
    return $foreign.dateMethodEff("getMinutes", dt);
};

// | Returns the milliseconds for a date, according to the current machine's
// | date/time locale.
var getMilliseconds = function (dt) {
    return $foreign.dateMethodEff("getMilliseconds", dt);
};

// | Returns the hour for a date, according to the current machine's date/time
// | locale.
var getHours = function (dt) {
    return $foreign.dateMethodEff("getHours", dt);
};

// | Returns the year for a date, according to the current machine's date/time
// | locale.
var getFullYear = function (dt) {
    return $foreign.dateMethodEff("getFullYear", dt);
};

// | Returns the day of the week for a date, according to the current
// | machine's date/time locale.
var getDay = function (dt) {
    return $foreign.dateMethodEff("getDay", dt);
};

// | Returns the day of the month for a date, according to the current
// | machine's date/time locale.
var getDate = function (dt) {
    return $foreign.dateMethodEff("getDate", dt);
};

// | Converts a `DateTime` value into a native JavaScript date object. The
// | resulting date value is guaranteed to be valid.
var fromDateTime = function (v) {
    return $foreign.jsdate({
        year: Data_Int.toNumber(Data_Enum.fromEnum(Data_Date_Component.boundedEnumYear)(Data_Date.year(v.value0))),
        month: Data_Int.toNumber(Data_Enum.fromEnum(Data_Date_Component.boundedEnumMonth)(Data_Date.month(v.value0)) - 1 | 0),
        day: Data_Int.toNumber(Data_Enum.fromEnum(Data_Date_Component.boundedEnumDay)(Data_Date.day(v.value0))),
        hour: Data_Int.toNumber(Data_Enum.fromEnum(Data_Time_Component.boundedEnumHour)(Data_Time.hour(v.value1))),
        minute: Data_Int.toNumber(Data_Enum.fromEnum(Data_Time_Component.boundedEnumMinute)(Data_Time.minute(v.value1))),
        second: Data_Int.toNumber(Data_Enum.fromEnum(Data_Time_Component.boundedEnumSecond)(Data_Time.second(v.value1))),
        millisecond: Data_Int.toNumber(Data_Enum.fromEnum(Data_Time_Component.boundedEnumMillisecond)(Data_Time.millisecond(v.value1)))
    });
};
var eqJSDate = {
    eq: function (a) {
        return function (b) {
            return getTime(a) === getTime(b);
        };
    }
};
var ordJSDate = {
    compare: function (a) {
        return function (b) {
            return Data_Ord.compare(Data_Ord.ordNumber)(getTime(a))(getTime(b));
        };
    },
    Eq0: function () {
        return eqJSDate;
    }
};
export {
    isValid,
    fromInstant,
    jsdate,
    jsdateLocal,
    now,
    parse,
    fromTime
} from "./foreign.js";
export {
    readDate,
    fromDateTime,
    toDateTime,
    toDate,
    toInstant,
    getTime,
    getUTCDate,
    getUTCDay,
    getUTCFullYear,
    getUTCHours,
    getUTCMilliseconds,
    getUTCMinutes,
    getUTCMonth,
    getUTCSeconds,
    getDate,
    getDay,
    getFullYear,
    getHours,
    getMilliseconds,
    getMinutes,
    getMonth,
    getSeconds,
    getTimezoneOffset,
    toDateString,
    toISOString,
    toString,
    toTimeString,
    toUTCString,
    eqJSDate,
    ordJSDate,
    showJSDate
};
