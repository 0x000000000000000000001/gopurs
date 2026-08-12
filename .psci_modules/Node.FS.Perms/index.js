import * as Control_Apply from "../Control.Apply/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Enum from "../Data.Enum/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Int from "../Data.Int/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_String_CodePoints from "../Data.String.CodePoints/index.js";
import * as Data_String_CodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data_String_Common from "../Data.String.Common/index.js";
var fromJust = /* #__PURE__ */ Data_Maybe.fromJust();
var fromJust1 = /* #__PURE__ */ Data_Maybe.fromJust();
var eqMaybe = /* #__PURE__ */ Data_Maybe.eqMaybe(Data_Eq.eqChar);
var eqRec = /* #__PURE__ */ Data_Eq.eqRec();
var eqRowCons = /* #__PURE__ */ Data_Eq.eqRowCons(Data_Eq.eqRowNil)();
var ordArray = /* #__PURE__ */ Data_Ord.ordArray(Data_Ord.ordBoolean);

// | A `Perm` value specifies what is allowed to be done with a particular
// | file by a particular class of user &mdash; that is, whether it is
// | readable, writable, and/or executable. It has a `Semiring` instance, which
// | allows you to combine permissions:
// |
// | - `(+)` adds `Perm` values together. For example, `read + write` means
// |   "readable and writable".
// | - `(*)` masks permissions. It can be thought of as selecting only the
// |    permissions that two `Perm` values have in common. For example:
// |    `(read + write) * (write + execute) == write`.
// |
// | You can think also of a `Perm` value as a subset of the set
// | `{ read, write, execute }`; then, `(+)` and `(*)` represent set union and
// | intersection respectively.
var Perm = function (x) {
    return x;
};

// | A `Perms` value includes all the permissions information about a
// | particular file or directory, by storing a `Perm` value for each of the
// | file owner, the group, and any other users.
var Perms = function (x) {
    return x;
};

// | The "writable" permission.
var write = {
    r: false,
    w: true,
    x: false
};
var semiringPerm = {
    add: function (v) {
        return function (v1) {
            return {
                r: v.r || v1.r,
                w: v.w || v1.w,
                x: v.x || v1.x
            };
        };
    },
    zero: {
        r: false,
        w: false,
        x: false
    },
    mul: function (v) {
        return function (v1) {
            return {
                r: v.r && v1.r,
                w: v.w && v1.w,
                x: v.x && v1.x
            };
        };
    },
    one: {
        r: true,
        w: true,
        x: true
    }
};

// | The "readable" permission.
var read = {
    r: true,
    w: false,
    x: false
};

// | Convert a `Perm` to an octal digit. For example:
// |
// | * `permToInt r == 4`
// | * `permToInt w == 2`
// | * `permToInt (r + w) == 6`
var permToInt = function (v) {
    return ((function () {
        if (v.r) {
            return 4;
        };
        return 0;
    })() + (function () {
        if (v.w) {
            return 2;
        };
        return 0;
    })() | 0) + (function () {
        if (v.x) {
            return 1;
        };
        return 0;
    })() | 0;
};

// | Convert a `Perm` to an octal string, via `permToInt`.
var permToString = /* #__PURE__ */ (function () {
    var $114 = Data_Show.show(Data_Show.showInt);
    return function ($115) {
        return $114(permToInt($115));
    };
})();

// | Convert a `Perms` value to an octal string, in a format similar to that
// | accepted by `chmod`. For example:
// | `permsToString (mkPerms (read + write) read read) == "0644"`
var permsToString = function (v) {
    return "0" + (permToString(v.u) + (permToString(v.g) + permToString(v.o)));
};

// | Convert a `Perms` value to an `Int`, via `permsToString`.
var permsToInt = /* #__PURE__ */ (function () {
    var $116 = Data_Int.fromStringAs(Data_Int.octal);
    return function ($117) {
        return fromJust($116(permsToString($117)));
    };
})();

// | No permissions. This is the identity of the `Semiring` operation `(+)`
// | for `Perm`; that is, it is the same as `zero`.
var none = /* #__PURE__ */ Data_Semiring.zero(semiringPerm);

// | Create a `Perms` value. The arguments represent the owner's, group's, and
// | other users' permission sets, respectively.
var mkPerms = function (u) {
    return function (g) {
        return function (o) {
            return {
                u: u,
                g: g,
                o: o
            };
        };
    };
};

// | Create a `Perm` value. The arguments represent the readable, writable, and
// | executable permissions, in that order.
var mkPerm = function (r) {
    return function (w) {
        return function (x) {
            return {
                r: r,
                w: w,
                x: x
            };
        };
    };
};

// | The "executable" permission.
var execute = {
    r: false,
    w: false,
    x: true
};
var permFromChar = function (c) {
    if (c === "0") {
        return new Data_Maybe.Just(none);
    };
    if (c === "1") {
        return new Data_Maybe.Just(execute);
    };
    if (c === "2") {
        return new Data_Maybe.Just(write);
    };
    if (c === "3") {
        return new Data_Maybe.Just(Data_Semiring.add(semiringPerm)(write)(execute));
    };
    if (c === "4") {
        return new Data_Maybe.Just(read);
    };
    if (c === "5") {
        return new Data_Maybe.Just(Data_Semiring.add(semiringPerm)(read)(execute));
    };
    if (c === "6") {
        return new Data_Maybe.Just(Data_Semiring.add(semiringPerm)(read)(write));
    };
    if (c === "7") {
        return new Data_Maybe.Just(Data_Semiring.add(semiringPerm)(Data_Semiring.add(semiringPerm)(read)(write))(execute));
    };
    return Data_Maybe.Nothing.value;
};

// | Attempt to parse a `Perms` value from a `String` containing an octal
// | integer. For example,
// | `permsFromString "0644" == Just (mkPerms (read + write) read read)`.
var permsFromString = /* #__PURE__ */ (function () {
    var zeroChar = fromJust1(Data_Enum.toEnum(Data_Enum.boundedEnumChar)(48));
    var dropPrefix = function (x) {
        return function (xs) {
            if (Data_Eq.eq(eqMaybe)(Data_String_CodeUnits.charAt(0)(xs))(new Data_Maybe.Just(x))) {
                return Data_String_CodePoints.drop(1)(xs);
            };
            if (Data_Boolean.otherwise) {
                return xs;
            };
            throw new Error("Failed pattern match at Node.FS.Perms (line 127, column 3 - line 129, column 21): " + [ x.constructor.name, xs.constructor.name ]);
        };
    };
    var _perms = function (v) {
        if (v.length === 3) {
            return Control_Apply.apply(Data_Maybe.applyMaybe)(Control_Apply.apply(Data_Maybe.applyMaybe)(Data_Functor.map(Data_Maybe.functorMaybe)(mkPerms)(permFromChar(v[0])))(permFromChar(v[1])))(permFromChar(v[2]));
        };
        return Data_Maybe.Nothing.value;
    };
    var $118 = dropPrefix(zeroChar);
    return function ($119) {
        return _perms(Data_String_CodeUnits.toCharArray($118($119)));
    };
})();
var eqPerm = /* #__PURE__ */ eqRec(/* #__PURE__ */ Data_Eq.eqRowCons(/* #__PURE__ */ Data_Eq.eqRowCons(/* #__PURE__ */ eqRowCons({
    reflectSymbol: function () {
        return "x";
    }
})(Data_Eq.eqBoolean))()({
    reflectSymbol: function () {
        return "w";
    }
})(Data_Eq.eqBoolean))()({
    reflectSymbol: function () {
        return "r";
    }
})(Data_Eq.eqBoolean));
var eqPerms = /* #__PURE__ */ eqRec(/* #__PURE__ */ Data_Eq.eqRowCons(/* #__PURE__ */ Data_Eq.eqRowCons(/* #__PURE__ */ eqRowCons({
    reflectSymbol: function () {
        return "u";
    }
})(eqPerm))()({
    reflectSymbol: function () {
        return "o";
    }
})(eqPerm))()({
    reflectSymbol: function () {
        return "g";
    }
})(eqPerm));
var ordPerm = {
    compare: function (v) {
        return function (v1) {
            return Data_Ord.compare(ordArray)([ v.r, v.w, v.x ])([ v1.r, v1.w, v1.x ]);
        };
    },
    Eq0: function () {
        return eqPerm;
    }
};
var ordArray1 = /* #__PURE__ */ Data_Ord.ordArray(ordPerm);
var ordPerms = {
    compare: function (v) {
        return function (v1) {
            return Data_Ord.compare(ordArray1)([ v.u, v.g, v.o ])([ v1.u, v1.g, v1.o ]);
        };
    },
    Eq0: function () {
        return eqPerms;
    }
};

// | All permissions: readable, writable, and executable. This is the identity
// | of the `Semiring` operation `(*)` for `Perm`; that is, it is the same as
// | `one`.
var all = /* #__PURE__ */ Data_Semiring.one(semiringPerm);
var permsAll = /* #__PURE__ */ mkPerms(all)(all)(all);
var permsReadWrite = /* #__PURE__ */ mkPerms(all)(all)(none);
var showPerm = {
    show: function (v) {
        if (Data_Eq.eq(eqPerm)(v)(none)) {
            return "none";
        };
        if (Data_Eq.eq(eqPerm)(v)(all)) {
            return "all";
        };
        var ps = Data_Semigroup.append(Data_Semigroup.semigroupArray)((function () {
            if (v.r) {
                return [ "read" ];
            };
            return [  ];
        })())(Data_Semigroup.append(Data_Semigroup.semigroupArray)((function () {
            if (v.w) {
                return [ "write" ];
            };
            return [  ];
        })())((function () {
            if (v.x) {
                return [ "execute" ];
            };
            return [  ];
        })()));
        return Data_String_Common.joinWith(" + ")(ps);
    }
};
var showPerms = {
    show: function (v) {
        var f = function (perm) {
            var str = Data_Show.show(showPerm)(perm);
            var $110 = Data_Maybe.isNothing(Data_String_CodePoints.indexOf(" ")(str));
            if ($110) {
                return str;
            };
            return "(" + (str + ")");
        };
        return "mkPerms " + Data_String_Common.joinWith(" ")(Data_Functor.map(Data_Functor.functorArray)(f)([ v.u, v.g, v.o ]));
    }
};
export {
    mkPerm,
    none,
    read,
    write,
    execute,
    all,
    mkPerms,
    permsAll,
    permsReadWrite,
    permsFromString,
    permsToString,
    permsToInt,
    eqPerm,
    ordPerm,
    showPerm,
    semiringPerm,
    eqPerms,
    ordPerms,
    showPerms
};
