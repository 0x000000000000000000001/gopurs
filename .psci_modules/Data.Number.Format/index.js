// | A module for formatting numbers as strings.
// |
// | Usage:
// | ``` purs
// | > let x = 1234.56789
// |
// | > toStringWith (precision 6) x
// | "1234.57"
// |
// | > toStringWith (fixed 3) x
// | "1234.568"
// |
// | > toStringWith (exponential 2) x
// | "1.23e+3"
// | ```
// |
// | The main method of this module is the `toStringWith` function that accepts
// | a `Format` argument which can be constructed through one of the smart
// | constructors `precision`, `fixed` and `exponential`. Internally, the
// | number will be formatted with JavaScripts `toPrecision`, `toFixed` or
// | `toExponential`.
import * as $foreign from "./foreign.js";
import * as Data_Ord from "../Data.Ord/index.js";

// | The `Format` data type specifies how a number will be formatted.
var Precision = /* #__PURE__ */ (function () {
    function Precision(value0) {
        this.value0 = value0;
    };
    Precision.create = function (value0) {
        return new Precision(value0);
    };
    return Precision;
})();

// | The `Format` data type specifies how a number will be formatted.
var Fixed = /* #__PURE__ */ (function () {
    function Fixed(value0) {
        this.value0 = value0;
    };
    Fixed.create = function (value0) {
        return new Fixed(value0);
    };
    return Fixed;
})();

// | The `Format` data type specifies how a number will be formatted.
var Exponential = /* #__PURE__ */ (function () {
    function Exponential(value0) {
        this.value0 = value0;
    };
    Exponential.create = function (value0) {
        return new Exponential(value0);
    };
    return Exponential;
})();

// | Convert a number to a string with a given format.
var toStringWith = function (v) {
    if (v instanceof Precision) {
        return $foreign.toPrecisionNative(v.value0);
    };
    if (v instanceof Fixed) {
        return $foreign.toFixedNative(v.value0);
    };
    if (v instanceof Exponential) {
        return $foreign.toExponentialNative(v.value0);
    };
    throw new Error("Failed pattern match at Data.Number.Format (line 59, column 1 - line 59, column 43): " + [ v.constructor.name ]);
};

// | Create a `toPrecision`-based format from an integer. Values smaller than
// | `1` and larger than `21` will be clamped.
var precision = /* #__PURE__ */ (function () {
    var $5 = Data_Ord.clamp(Data_Ord.ordInt)(1)(21);
    return function ($6) {
        return Precision.create($5($6));
    };
})();

// | Create a `toFixed`-based format from an integer. Values smaller than `0`
// | and larger than `20` will be clamped.
var fixed = /* #__PURE__ */ (function () {
    var $7 = Data_Ord.clamp(Data_Ord.ordInt)(0)(20);
    return function ($8) {
        return Fixed.create($7($8));
    };
})();

// | Create a `toExponential`-based format from an integer. Values smaller than
// | `0` and larger than `20` will be clamped.
var exponential = /* #__PURE__ */ (function () {
    var $9 = Data_Ord.clamp(Data_Ord.ordInt)(0)(20);
    return function ($10) {
        return Exponential.create($9($10));
    };
})();
export {
    toString
} from "./foreign.js";
export {
    precision,
    fixed,
    exponential,
    toStringWith
};
