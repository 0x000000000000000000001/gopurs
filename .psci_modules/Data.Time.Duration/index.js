import * as Control_Category from "../Control.Category/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ring from "../Data.Ring/index.js";
import * as Data_Show from "../Data.Show/index.js";
var negate = /* #__PURE__ */ Data_Ring.negate(Data_Ring.ringNumber);
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);

// | A duration measured in seconds.
var Seconds = function (x) {
    return x;
};

// | A duration measured in minutes.
var Minutes = function (x) {
    return x;
};

// | A duration measured in milliseconds.
var Milliseconds = function (x) {
    return x;
};

// | A duration measured in hours.
var Hours = function (x) {
    return x;
};

// | A duration measured in days, where a day is assumed to be exactly 24 hours.
var Days = function (x) {
    return x;
};
var toDuration = function (dict) {
    return dict.toDuration;
};
var showSeconds = {
    show: function (v) {
        return "(Seconds " + (Data_Show.show(Data_Show.showNumber)(v) + ")");
    }
};
var showMinutes = {
    show: function (v) {
        return "(Minutes " + (Data_Show.show(Data_Show.showNumber)(v) + ")");
    }
};
var showMilliseconds = {
    show: function (v) {
        return "(Milliseconds " + (Data_Show.show(Data_Show.showNumber)(v) + ")");
    }
};
var showHours = {
    show: function (v) {
        return "(Hours " + (Data_Show.show(Data_Show.showNumber)(v) + ")");
    }
};
var showDays = {
    show: function (v) {
        return "(Days " + (Data_Show.show(Data_Show.showNumber)(v) + ")");
    }
};
var semigroupSeconds = {
    append: function (v) {
        return function (v1) {
            return v + v1;
        };
    }
};
var semigroupMinutes = {
    append: function (v) {
        return function (v1) {
            return v + v1;
        };
    }
};
var semigroupMilliseconds = {
    append: function (v) {
        return function (v1) {
            return v + v1;
        };
    }
};
var semigroupHours = {
    append: function (v) {
        return function (v1) {
            return v + v1;
        };
    }
};
var semigroupDays = {
    append: function (v) {
        return function (v1) {
            return v + v1;
        };
    }
};
var ordSeconds = Data_Ord.ordNumber;
var ordMinutes = Data_Ord.ordNumber;
var ordMilliseconds = Data_Ord.ordNumber;
var ordHours = Data_Ord.ordNumber;
var ordDays = Data_Ord.ordNumber;
var newtypeSeconds = {
    Coercible0: function () {
        return undefined;
    }
};
var newtypeMinutes = {
    Coercible0: function () {
        return undefined;
    }
};
var newtypeMilliseconds = {
    Coercible0: function () {
        return undefined;
    }
};
var newtypeHours = {
    Coercible0: function () {
        return undefined;
    }
};
var newtypeDays = {
    Coercible0: function () {
        return undefined;
    }
};
var monoidSeconds = {
    mempty: 0.0,
    Semigroup0: function () {
        return semigroupSeconds;
    }
};
var monoidMinutes = {
    mempty: 0.0,
    Semigroup0: function () {
        return semigroupMinutes;
    }
};
var monoidMilliseconds = {
    mempty: 0.0,
    Semigroup0: function () {
        return semigroupMilliseconds;
    }
};
var monoidHours = {
    mempty: 0.0,
    Semigroup0: function () {
        return semigroupHours;
    }
};
var monoidDays = {
    mempty: 0.0,
    Semigroup0: function () {
        return semigroupDays;
    }
};
var fromDuration = function (dict) {
    return dict.fromDuration;
};

// | Negates a duration, turning a positive duration negative or a negative
// | duration positive.
var negateDuration = function (dictDuration) {
    var $49 = toDuration(dictDuration);
    var $50 = Data_Newtype.over()()(Milliseconds)(negate);
    var $51 = fromDuration(dictDuration);
    return function ($52) {
        return $49($50($51($52)));
    };
};
var eqSeconds = Data_Eq.eqNumber;
var eqMinutes = Data_Eq.eqNumber;
var eqMilliseconds = Data_Eq.eqNumber;
var eqHours = Data_Eq.eqNumber;
var eqDays = Data_Eq.eqNumber;
var durationSeconds = {
    fromDuration: /* #__PURE__ */ Data_Newtype.over()()(Seconds)(function (v) {
        return v * 1000.0;
    }),
    toDuration: /* #__PURE__ */ Data_Newtype.over()()(Milliseconds)(function (v) {
        return v / 1000.0;
    })
};
var durationMinutes = {
    fromDuration: /* #__PURE__ */ Data_Newtype.over()()(Minutes)(function (v) {
        return v * 60000.0;
    }),
    toDuration: /* #__PURE__ */ Data_Newtype.over()()(Milliseconds)(function (v) {
        return v / 60000.0;
    })
};
var durationMilliseconds = {
    fromDuration: identity,
    toDuration: identity
};
var durationHours = {
    fromDuration: /* #__PURE__ */ Data_Newtype.over()()(Hours)(function (v) {
        return v * 3600000.0;
    }),
    toDuration: /* #__PURE__ */ Data_Newtype.over()()(Milliseconds)(function (v) {
        return v / 3600000.0;
    })
};
var durationDays = {
    fromDuration: /* #__PURE__ */ Data_Newtype.over()()(Days)(function (v) {
        return v * 8.64e7;
    }),
    toDuration: /* #__PURE__ */ Data_Newtype.over()()(Milliseconds)(function (v) {
        return v / 8.64e7;
    })
};

// | Converts directly between durations of differing types.
var convertDuration = function (dictDuration) {
    var fromDuration1 = fromDuration(dictDuration);
    return function (dictDuration1) {
        var $53 = toDuration(dictDuration1);
        return function ($54) {
            return $53(fromDuration1($54));
        };
    };
};
export {
    fromDuration,
    toDuration,
    Milliseconds,
    Seconds,
    Minutes,
    Hours,
    Days,
    convertDuration,
    negateDuration,
    newtypeMilliseconds,
    eqMilliseconds,
    ordMilliseconds,
    semigroupMilliseconds,
    monoidMilliseconds,
    showMilliseconds,
    newtypeSeconds,
    eqSeconds,
    ordSeconds,
    semigroupSeconds,
    monoidSeconds,
    showSeconds,
    newtypeMinutes,
    eqMinutes,
    ordMinutes,
    semigroupMinutes,
    monoidMinutes,
    showMinutes,
    newtypeHours,
    eqHours,
    ordHours,
    semigroupHours,
    monoidHours,
    showHours,
    newtypeDays,
    eqDays,
    ordDays,
    semigroupDays,
    monoidDays,
    showDays,
    durationMilliseconds,
    durationSeconds,
    durationMinutes,
    durationHours,
    durationDays
};
