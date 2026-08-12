// | Functions for working with PureScripts builtin `Number` type.
import * as $foreign from "./foreign.js";
import * as Data_Maybe from "../Data.Maybe/index.js";

// | The ratio of the circumference of a circle to its radius.
// | ```purs
// | > 2.0 * pi == tau
// | true
// |
// | > tau
// | 6.283185307179586
// | ```
var tau = 6.283185307179586;

// | The square root of two.
// | ```purs
// | > sqrt 2.0 == sqrt2
// | true
// |
// | > sqrt2
// | 1.4142135623730951
// | ```
var sqrt2 = 1.4142135623730951;

// | The square root of one half.
// | ```purs
// | > sqrt 0.5 == sqrt1_2
// | true
// |
// | > sqrt1_2
// | 0.7071067811865476
// | ```
var sqrt1_2 = 0.7071067811865476;

// | The ratio of the circumference of a circle to its diameter.
// | ```purs
// | > pi
// | 3.141592653589793
// | ```
var pi = 3.141592653589793;

// | The base 2 logarithm of `e`.
// | ```purs
// | > 1.0 / ln2 == log2e
// | true
// |
// | > log2e
// | 1.4426950408889634
// | ```
var log2e = 1.4426950408889634;

// | Base 10 logarithm of `e`.
// | ```purs
// | > 1.0 / ln10 - log10e
// | -5.551115123125783e-17
// |
// | > log10e
// | 0.4342944819032518
// | ```
var log10e = 0.4342944819032518;

// | The natural logarithm of 2.
// | ```purs
// | > log 2.0 == ln2
// | true
// |
// | > ln2
// | 0.6931471805599453
// | ```
var ln2 = 0.6931471805599453;

// | The natural logarithm of 10.
// | ```purs
// | > log 10.0 == ln10
// | true
// |
// | > ln10
// | 2.302585092994046
// | ```
var ln10 = 2.302585092994046;

// | Attempt to parse a `Number` using JavaScripts `parseFloat`. Returns
// | `Nothing` if the parse fails or if the result is not a finite number.
// |
// | Example:
// | ```purs
// | > fromString "123"
// | (Just 123.0)
// |
// | > fromString "12.34"
// | (Just 12.34)
// |
// | > fromString "1e4"
// | (Just 10000.0)
// |
// | > fromString "1.2e4"
// | (Just 12000.0)
// |
// | > fromString "bad"
// | Nothing
// | ```
// |
// | Note that `parseFloat` allows for trailing non-digit characters and
// | whitespace as a prefix:
// | ```
// | > fromString "  1.2 ??"
// | (Just 1.2)
// | ```
var fromString = function (str) {
    return $foreign.fromStringImpl(str, $foreign["isFinite"], Data_Maybe.Just.create, Data_Maybe.Nothing.value);
};

// | The base of the natural logarithm, also known as Euler's number or *e*.
// | ```purs
// | > log e
// | 1.0
// |
// | > exp 1.0 == e
// | true
// |
// | > e
// | 2.718281828459045
// | ```
var e = 2.718281828459045;
export {
    nan,
    isNaN,
    infinity,
    isFinite,
    abs,
    acos,
    asin,
    atan,
    atan2,
    ceil,
    cos,
    exp,
    floor,
    log,
    max,
    min,
    pow,
    remainder,
    round,
    sign,
    sin,
    sqrt,
    tan,
    trunc
} from "./foreign.js";
export {
    fromString,
    e,
    ln2,
    ln10,
    log10e,
    log2e,
    pi,
    sqrt1_2,
    sqrt2,
    tau
};
