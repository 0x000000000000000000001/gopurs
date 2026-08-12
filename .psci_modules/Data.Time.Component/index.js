import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Enum from "../Data.Enum/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Show from "../Data.Show/index.js";
var $runtime_lazy = function (name, moduleName, init) {
    var state = 0;
    var val;
    return function (lineNumber) {
        if (state === 2) return val;
        if (state === 1) throw new ReferenceError(name + " was needed before it finished initializing (module " + moduleName + ", line " + lineNumber + ")", moduleName, lineNumber);
        state = 1;
        val = init();
        state = 2;
        return val;
    };
};

// | An second component for a time value.
// |
// | The constructor is private as values for the type are restricted to the
// | range 0 to 59, inclusive. The `toEnum` function can be used to safely
// | acquire an `Second` value from an integer. Correspondingly, a `Second` can
// | be lowered to a plain integer with the `fromEnum` function.
var Second = function (x) {
    return x;
};

// | An minute component for a time value.
// |
// | The constructor is private as values for the type are restricted to the
// | range 0 to 59, inclusive. The `toEnum` function can be used to safely
// | acquire an `Minute` value from an integer. Correspondingly, a `Minute` can
// | be lowered to a plain integer with the `fromEnum` function.
var Minute = function (x) {
    return x;
};

// | An millisecond component for a time value.
// |
// | The constructor is private as values for the type are restricted to the
// | range 0 to 999, inclusive. The `toEnum` function can be used to safely
// | acquire an `Millisecond` value from an integer. Correspondingly, a
// | `Millisecond` can be lowered to a plain integer with the `fromEnum`
// | function.
var Millisecond = function (x) {
    return x;
};

// | An hour component for a time value.
// |
// | The constructor is private as values for the type are restricted to the
// | range 0 to 23, inclusive. The `toEnum` function can be used to safely
// | acquire an `Hour` value from an integer. Correspondingly, an `Hour` can be
// | lowered to a plain integer with the `fromEnum` function.
var Hour = function (x) {
    return x;
};
var showSecond = {
    show: function (v) {
        return "(Second " + (Data_Show.show(Data_Show.showInt)(v) + ")");
    }
};
var showMinute = {
    show: function (v) {
        return "(Minute " + (Data_Show.show(Data_Show.showInt)(v) + ")");
    }
};
var showMillisecond = {
    show: function (v) {
        return "(Millisecond " + (Data_Show.show(Data_Show.showInt)(v) + ")");
    }
};
var showHour = {
    show: function (v) {
        return "(Hour " + (Data_Show.show(Data_Show.showInt)(v) + ")");
    }
};
var ordSecond = Data_Ord.ordInt;
var ordMinute = Data_Ord.ordInt;
var ordMillisecond = Data_Ord.ordInt;
var ordHour = Data_Ord.ordInt;
var eqSecond = Data_Eq.eqInt;
var eqMinute = Data_Eq.eqInt;
var eqMillisecond = Data_Eq.eqInt;
var eqHour = Data_Eq.eqInt;
var boundedSecond = {
    bottom: 0,
    top: 59,
    Ord0: function () {
        return ordSecond;
    }
};
var boundedMinute = {
    bottom: 0,
    top: 59,
    Ord0: function () {
        return ordMinute;
    }
};
var boundedMillisecond = {
    bottom: 0,
    top: 999,
    Ord0: function () {
        return ordMillisecond;
    }
};
var boundedHour = {
    bottom: 0,
    top: 23,
    Ord0: function () {
        return ordHour;
    }
};
var boundedEnumSecond = {
    cardinality: 60,
    toEnum: function (n) {
        if (n >= 0 && n <= 59) {
            return new Data_Maybe.Just(n);
        };
        if (Data_Boolean.otherwise) {
            return Data_Maybe.Nothing.value;
        };
        throw new Error("Failed pattern match at Data.Time.Component (line 90, column 1 - line 95, column 26): " + [ n.constructor.name ]);
    },
    fromEnum: function (v) {
        return v;
    },
    Bounded0: function () {
        return boundedSecond;
    },
    Enum1: function () {
        return $lazy_enumSecond(0);
    }
};
var $lazy_enumSecond = /* #__PURE__ */ $runtime_lazy("enumSecond", "Data.Time.Component", function () {
    return {
        succ: (function () {
            var $28 = Data_Enum.toEnum(boundedEnumSecond);
            var $29 = Data_Enum.fromEnum(boundedEnumSecond);
            return function ($30) {
                return $28((function (v) {
                    return v + 1 | 0;
                })($29($30)));
            };
        })(),
        pred: (function () {
            var $31 = Data_Enum.toEnum(boundedEnumSecond);
            var $32 = Data_Enum.fromEnum(boundedEnumSecond);
            return function ($33) {
                return $31((function (v) {
                    return v - 1 | 0;
                })($32($33)));
            };
        })(),
        Ord0: function () {
            return ordSecond;
        }
    };
});
var enumSecond = /* #__PURE__ */ $lazy_enumSecond(86);
var boundedEnumMinute = {
    cardinality: 60,
    toEnum: function (n) {
        if (n >= 0 && n <= 59) {
            return new Data_Maybe.Just(n);
        };
        if (Data_Boolean.otherwise) {
            return Data_Maybe.Nothing.value;
        };
        throw new Error("Failed pattern match at Data.Time.Component (line 61, column 1 - line 66, column 26): " + [ n.constructor.name ]);
    },
    fromEnum: function (v) {
        return v;
    },
    Bounded0: function () {
        return boundedMinute;
    },
    Enum1: function () {
        return $lazy_enumMinute(0);
    }
};
var $lazy_enumMinute = /* #__PURE__ */ $runtime_lazy("enumMinute", "Data.Time.Component", function () {
    return {
        succ: (function () {
            var $34 = Data_Enum.toEnum(boundedEnumMinute);
            var $35 = Data_Enum.fromEnum(boundedEnumMinute);
            return function ($36) {
                return $34((function (v) {
                    return v + 1 | 0;
                })($35($36)));
            };
        })(),
        pred: (function () {
            var $37 = Data_Enum.toEnum(boundedEnumMinute);
            var $38 = Data_Enum.fromEnum(boundedEnumMinute);
            return function ($39) {
                return $37((function (v) {
                    return v - 1 | 0;
                })($38($39)));
            };
        })(),
        Ord0: function () {
            return ordMinute;
        }
    };
});
var enumMinute = /* #__PURE__ */ $lazy_enumMinute(57);
var boundedEnumMillisecond = {
    cardinality: 1000,
    toEnum: function (n) {
        if (n >= 0 && n <= 999) {
            return new Data_Maybe.Just(n);
        };
        if (Data_Boolean.otherwise) {
            return Data_Maybe.Nothing.value;
        };
        throw new Error("Failed pattern match at Data.Time.Component (line 120, column 1 - line 125, column 31): " + [ n.constructor.name ]);
    },
    fromEnum: function (v) {
        return v;
    },
    Bounded0: function () {
        return boundedMillisecond;
    },
    Enum1: function () {
        return $lazy_enumMillisecond(0);
    }
};
var $lazy_enumMillisecond = /* #__PURE__ */ $runtime_lazy("enumMillisecond", "Data.Time.Component", function () {
    return {
        succ: (function () {
            var $40 = Data_Enum.toEnum(boundedEnumMillisecond);
            var $41 = Data_Enum.fromEnum(boundedEnumMillisecond);
            return function ($42) {
                return $40((function (v) {
                    return v + 1 | 0;
                })($41($42)));
            };
        })(),
        pred: (function () {
            var $43 = Data_Enum.toEnum(boundedEnumMillisecond);
            var $44 = Data_Enum.fromEnum(boundedEnumMillisecond);
            return function ($45) {
                return $43((function (v) {
                    return v - 1 | 0;
                })($44($45)));
            };
        })(),
        Ord0: function () {
            return ordMillisecond;
        }
    };
});
var enumMillisecond = /* #__PURE__ */ $lazy_enumMillisecond(116);
var boundedEnumHour = {
    cardinality: 24,
    toEnum: function (n) {
        if (n >= 0 && n <= 23) {
            return new Data_Maybe.Just(n);
        };
        if (Data_Boolean.otherwise) {
            return Data_Maybe.Nothing.value;
        };
        throw new Error("Failed pattern match at Data.Time.Component (line 32, column 1 - line 37, column 24): " + [ n.constructor.name ]);
    },
    fromEnum: function (v) {
        return v;
    },
    Bounded0: function () {
        return boundedHour;
    },
    Enum1: function () {
        return $lazy_enumHour(0);
    }
};
var $lazy_enumHour = /* #__PURE__ */ $runtime_lazy("enumHour", "Data.Time.Component", function () {
    return {
        succ: (function () {
            var $46 = Data_Enum.toEnum(boundedEnumHour);
            var $47 = Data_Enum.fromEnum(boundedEnumHour);
            return function ($48) {
                return $46((function (v) {
                    return v + 1 | 0;
                })($47($48)));
            };
        })(),
        pred: (function () {
            var $49 = Data_Enum.toEnum(boundedEnumHour);
            var $50 = Data_Enum.fromEnum(boundedEnumHour);
            return function ($51) {
                return $49((function (v) {
                    return v - 1 | 0;
                })($50($51)));
            };
        })(),
        Ord0: function () {
            return ordHour;
        }
    };
});
var enumHour = /* #__PURE__ */ $lazy_enumHour(28);
export {
    eqHour,
    ordHour,
    boundedHour,
    enumHour,
    boundedEnumHour,
    showHour,
    eqMinute,
    ordMinute,
    boundedMinute,
    enumMinute,
    boundedEnumMinute,
    showMinute,
    eqSecond,
    ordSecond,
    boundedSecond,
    enumSecond,
    boundedEnumSecond,
    showSecond,
    eqMillisecond,
    ordMillisecond,
    boundedMillisecond,
    enumMillisecond,
    boundedEnumMillisecond,
    showMillisecond
};
