import * as $foreign from "./foreign.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Bounded from "../Data.Bounded/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Number from "../Data.Number/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
var top = /* #__PURE__ */ Data_Bounded.top(Data_Bounded.boundedInt);
var bottom = /* #__PURE__ */ Data_Bounded.bottom(Data_Bounded.boundedInt);

// | The number of unique digits (including zero) used to represent integers in
// | a specific base.
var Radix = function (x) {
    return x;
};

// | A type for describing whether an integer is even or odd.
// |
// | The `Ord` instance considers `Even` to be less than `Odd`.
// |
// | The `Semiring` instance allows you to ask about the parity of the results
// | of arithmetical operations, given only the parities of the inputs. For
// | example, the sum of an odd number and an even number is odd, so
// | `Odd + Even == Odd`. This also works for multiplication, eg. the product
// | of two odd numbers is odd, and therefore `Odd * Odd == Odd`.
// |
// | More generally, we have that
// |
// | ```purescript
// | parity x + parity y == parity (x + y)
// | parity x * parity y == parity (x * y)
// | ```
// |
// | for any integers `x`, `y`. (A mathematician would say that `parity` is a
// | *ring homomorphism*.)
// |
// | After defining addition and multiplication on `Parity` in this way, the
// | `Semiring` laws now force us to choose `zero = Even` and `one = Odd`.
// | This `Semiring` instance actually turns out to be a `Field`.
var Even = /* #__PURE__ */ (function () {
    function Even() {

    };
    Even.value = new Even();
    return Even;
})();

// | A type for describing whether an integer is even or odd.
// |
// | The `Ord` instance considers `Even` to be less than `Odd`.
// |
// | The `Semiring` instance allows you to ask about the parity of the results
// | of arithmetical operations, given only the parities of the inputs. For
// | example, the sum of an odd number and an even number is odd, so
// | `Odd + Even == Odd`. This also works for multiplication, eg. the product
// | of two odd numbers is odd, and therefore `Odd * Odd == Odd`.
// |
// | More generally, we have that
// |
// | ```purescript
// | parity x + parity y == parity (x + y)
// | parity x * parity y == parity (x * y)
// | ```
// |
// | for any integers `x`, `y`. (A mathematician would say that `parity` is a
// | *ring homomorphism*.)
// |
// | After defining addition and multiplication on `Parity` in this way, the
// | `Semiring` laws now force us to choose `zero = Even` and `one = Odd`.
// | This `Semiring` instance actually turns out to be a `Field`.
var Odd = /* #__PURE__ */ (function () {
    function Odd() {

    };
    Odd.value = new Odd();
    return Odd;
})();
var showParity = {
    show: function (v) {
        if (v instanceof Even) {
            return "Even";
        };
        if (v instanceof Odd) {
            return "Odd";
        };
        throw new Error("Failed pattern match at Data.Int (line 117, column 1 - line 119, column 19): " + [ v.constructor.name ]);
    }
};

// | Create a `Radix` from a number between 2 and 36.
var radix = function (n) {
    if (n >= 2 && n <= 36) {
        return new Data_Maybe.Just(n);
    };
    if (Data_Boolean.otherwise) {
        return Data_Maybe.Nothing.value;
    };
    throw new Error("Failed pattern match at Data.Int (line 198, column 1 - line 198, column 28): " + [ n.constructor.name ]);
};

// | The negation of `even`.
// |
// | ``` purescript
// | odd 0 == false
// | odd 1 == true
// | ```
var odd = function (x) {
    return (x & 1) !== 0;
};

// | The base-8 system.
var octal = 8;

// | The base-16 system.
var hexadecimal = 16;

// | Like `fromString`, but the integer can be specified in a different base.
// |
// | Example:
// | ``` purs
// | fromStringAs binary      "100" == Just 4
// | fromStringAs hexadecimal "ff"  == Just 255
// | ```
var fromStringAs = /* #__PURE__ */ (function () {
    return $foreign.fromStringAsImpl(Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();

// | Reads an `Int` from a `String` value. The number must parse as an integer
// | and fall within the valid range of values for the `Int` type, otherwise
// | `Nothing` is returned.
var fromString = /* #__PURE__ */ fromStringAs(10);

// | Creates an `Int` from a `Number` value. The number must already be an
// | integer and fall within the valid range of values for the `Int` type
// | otherwise `Nothing` is returned.
var fromNumber = /* #__PURE__ */ (function () {
    return $foreign.fromNumberImpl(Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();

// | Convert an integral `Number` to an `Int`, by clamping to the `Int` range.
// | This function will return 0 if the input is `NaN` or an `Infinity`.
var unsafeClamp = function (x) {
    if (!Data_Number["isFinite"](x)) {
        return 0;
    };
    if (x >= $foreign.toNumber(top)) {
        return top;
    };
    if (x <= $foreign.toNumber(bottom)) {
        return bottom;
    };
    if (Data_Boolean.otherwise) {
        return Data_Maybe.fromMaybe(0)(fromNumber(x));
    };
    throw new Error("Failed pattern match at Data.Int (line 72, column 1 - line 72, column 29): " + [ x.constructor.name ]);
};

// | Convert a `Number` to an `Int`, by taking the nearest integer to the
// | argument. Values outside the `Int` range are clamped, `NaN` and `Infinity`
// | values return 0.
var round = function ($27) {
    return unsafeClamp(Data_Number.round($27));
};

// | Convert a `Number` to an `Int`, by dropping the decimal.
// | Values outside the `Int` range are clamped, `NaN` and `Infinity`
// | values return 0.
var trunc = function ($28) {
    return unsafeClamp(Data_Number.trunc($28));
};

// | Convert a `Number` to an `Int`, by taking the closest integer equal to or
// | less than the argument. Values outside the `Int` range are clamped, `NaN`
// | and `Infinity` values return 0.
var floor = function ($29) {
    return unsafeClamp(Data_Number.floor($29));
};

// | Returns whether an `Int` is an even number.
// |
// | ``` purescript
// | even 0 == true
// | even 1 == false
// | ```
var even = function (x) {
    return (x & 1) === 0;
};

// | Returns whether an `Int` is `Even` or `Odd`.
// |
// | ``` purescript
// | parity 0 == Even
// | parity 1 == Odd
// | ```
var parity = function (n) {
    var $18 = even(n);
    if ($18) {
        return Even.value;
    };
    return Odd.value;
};
var eqParity = {
    eq: function (x) {
        return function (y) {
            if (x instanceof Even && y instanceof Even) {
                return true;
            };
            if (x instanceof Odd && y instanceof Odd) {
                return true;
            };
            return false;
        };
    }
};
var ordParity = {
    compare: function (x) {
        return function (y) {
            if (x instanceof Even && y instanceof Even) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Even) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Even) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Odd && y instanceof Odd) {
                return Data_Ordering.EQ.value;
            };
            throw new Error("Failed pattern match at Data.Int (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqParity;
    }
};
var semiringParity = /* #__PURE__ */ (function () {
    return {
        zero: Even.value,
        add: function (x) {
            return function (y) {
                var $23 = Data_Eq.eq(eqParity)(x)(y);
                if ($23) {
                    return Even.value;
                };
                return Odd.value;
            };
        },
        one: Odd.value,
        mul: function (v) {
            return function (v1) {
                if (v instanceof Odd && v1 instanceof Odd) {
                    return Odd.value;
                };
                return Even.value;
            };
        }
    };
})();
var ringParity = {
    sub: /* #__PURE__ */ Data_Semiring.add(semiringParity),
    Semiring0: function () {
        return semiringParity;
    }
};
var divisionRingParity = {
    recip: /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn),
    Ring0: function () {
        return ringParity;
    }
};

// | The base-10 system.
var decimal = 10;
var commutativeRingParity = {
    Ring0: function () {
        return ringParity;
    }
};
var euclideanRingParity = {
    degree: function (v) {
        if (v instanceof Even) {
            return 0;
        };
        if (v instanceof Odd) {
            return 1;
        };
        throw new Error("Failed pattern match at Data.Int (line 137, column 1 - line 141, column 17): " + [ v.constructor.name ]);
    },
    div: function (x) {
        return function (v) {
            return x;
        };
    },
    mod: function (v) {
        return function (v1) {
            return Even.value;
        };
    },
    CommutativeRing0: function () {
        return commutativeRingParity;
    }
};

// | Convert a `Number` to an `Int`, by taking the closest integer equal to or
// | greater than the argument. Values outside the `Int` range are clamped,
// | `NaN` and `Infinity` values return 0.
var ceil = function ($30) {
    return unsafeClamp(Data_Number.ceil($30));
};
var boundedParity = /* #__PURE__ */ (function () {
    return {
        bottom: Even.value,
        top: Odd.value,
        Ord0: function () {
            return ordParity;
        }
    };
})();

// | The base-2 system.
var binary = 2;

// | The base-36 system.
var base36 = 36;
export {
    toNumber,
    toStringAs,
    quot,
    rem,
    pow
} from "./foreign.js";
export {
    fromNumber,
    ceil,
    floor,
    trunc,
    round,
    fromString,
    radix,
    binary,
    octal,
    decimal,
    hexadecimal,
    base36,
    fromStringAs,
    Even,
    Odd,
    parity,
    even,
    odd,
    eqParity,
    ordParity,
    showParity,
    boundedParity,
    semiringParity,
    ringParity,
    commutativeRingParity,
    euclideanRingParity,
    divisionRingParity
};
