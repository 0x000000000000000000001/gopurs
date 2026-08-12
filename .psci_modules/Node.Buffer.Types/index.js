
// | Enumeration of the numeric types that can be written to a buffer.
var UInt8 = /* #__PURE__ */ (function () {
    function UInt8() {

    };
    UInt8.value = new UInt8();
    return UInt8;
})();

// | Enumeration of the numeric types that can be written to a buffer.
var UInt16LE = /* #__PURE__ */ (function () {
    function UInt16LE() {

    };
    UInt16LE.value = new UInt16LE();
    return UInt16LE;
})();

// | Enumeration of the numeric types that can be written to a buffer.
var UInt16BE = /* #__PURE__ */ (function () {
    function UInt16BE() {

    };
    UInt16BE.value = new UInt16BE();
    return UInt16BE;
})();

// | Enumeration of the numeric types that can be written to a buffer.
var UInt32LE = /* #__PURE__ */ (function () {
    function UInt32LE() {

    };
    UInt32LE.value = new UInt32LE();
    return UInt32LE;
})();

// | Enumeration of the numeric types that can be written to a buffer.
var UInt32BE = /* #__PURE__ */ (function () {
    function UInt32BE() {

    };
    UInt32BE.value = new UInt32BE();
    return UInt32BE;
})();

// | Enumeration of the numeric types that can be written to a buffer.
var Int8 = /* #__PURE__ */ (function () {
    function Int8() {

    };
    Int8.value = new Int8();
    return Int8;
})();

// | Enumeration of the numeric types that can be written to a buffer.
var Int16LE = /* #__PURE__ */ (function () {
    function Int16LE() {

    };
    Int16LE.value = new Int16LE();
    return Int16LE;
})();

// | Enumeration of the numeric types that can be written to a buffer.
var Int16BE = /* #__PURE__ */ (function () {
    function Int16BE() {

    };
    Int16BE.value = new Int16BE();
    return Int16BE;
})();

// | Enumeration of the numeric types that can be written to a buffer.
var Int32LE = /* #__PURE__ */ (function () {
    function Int32LE() {

    };
    Int32LE.value = new Int32LE();
    return Int32LE;
})();

// | Enumeration of the numeric types that can be written to a buffer.
var Int32BE = /* #__PURE__ */ (function () {
    function Int32BE() {

    };
    Int32BE.value = new Int32BE();
    return Int32BE;
})();

// | Enumeration of the numeric types that can be written to a buffer.
var FloatLE = /* #__PURE__ */ (function () {
    function FloatLE() {

    };
    FloatLE.value = new FloatLE();
    return FloatLE;
})();

// | Enumeration of the numeric types that can be written to a buffer.
var FloatBE = /* #__PURE__ */ (function () {
    function FloatBE() {

    };
    FloatBE.value = new FloatBE();
    return FloatBE;
})();

// | Enumeration of the numeric types that can be written to a buffer.
var DoubleLE = /* #__PURE__ */ (function () {
    function DoubleLE() {

    };
    DoubleLE.value = new DoubleLE();
    return DoubleLE;
})();

// | Enumeration of the numeric types that can be written to a buffer.
var DoubleBE = /* #__PURE__ */ (function () {
    function DoubleBE() {

    };
    DoubleBE.value = new DoubleBE();
    return DoubleBE;
})();
var showBufferValueType = {
    show: function (v) {
        if (v instanceof UInt8) {
            return "UInt8";
        };
        if (v instanceof UInt16LE) {
            return "UInt16LE";
        };
        if (v instanceof UInt16BE) {
            return "UInt16BE";
        };
        if (v instanceof UInt32LE) {
            return "UInt32LE";
        };
        if (v instanceof UInt32BE) {
            return "UInt32BE";
        };
        if (v instanceof Int8) {
            return "Int8";
        };
        if (v instanceof Int16LE) {
            return "Int16LE";
        };
        if (v instanceof Int16BE) {
            return "Int16BE";
        };
        if (v instanceof Int32LE) {
            return "Int32LE";
        };
        if (v instanceof Int32BE) {
            return "Int32BE";
        };
        if (v instanceof FloatLE) {
            return "FloatLE";
        };
        if (v instanceof FloatBE) {
            return "FloatBE";
        };
        if (v instanceof DoubleLE) {
            return "DoubleLE";
        };
        if (v instanceof DoubleBE) {
            return "DoubleBE";
        };
        throw new Error("Failed pattern match at Node.Buffer.Types (line 33, column 1 - line 47, column 29): " + [ v.constructor.name ]);
    }
};
export {
    UInt8,
    UInt16LE,
    UInt16BE,
    UInt32LE,
    UInt32BE,
    Int8,
    Int16LE,
    Int16BE,
    Int32LE,
    Int32BE,
    FloatLE,
    FloatBE,
    DoubleLE,
    DoubleBE,
    showBufferValueType
};
