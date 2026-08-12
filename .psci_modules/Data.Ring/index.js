import * as $foreign from "./foreign.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
import * as Data_Symbol from "../Data.Symbol/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Record_Unsafe from "../Record.Unsafe/index.js";
import * as Type_Proxy from "../Type.Proxy/index.js";
var semiringRecord = /* #__PURE__ */ Data_Semiring.semiringRecord();
var subRecord = function (dict) {
    return dict.subRecord;
};
var sub = function (dict) {
    return dict.sub;
};
var ringUnit = {
    sub: function (v) {
        return function (v1) {
            return Data_Unit.unit;
        };
    },
    Semiring0: function () {
        return Data_Semiring.semiringUnit;
    }
};
var ringRecordNil = {
    subRecord: function (v) {
        return function (v1) {
            return function (v2) {
                return {};
            };
        };
    },
    SemiringRecord0: function () {
        return Data_Semiring.semiringRecordNil;
    }
};
var ringRecordCons = function (dictIsSymbol) {
    var semiringRecordCons = Data_Semiring.semiringRecordCons(dictIsSymbol)();
    return function () {
        return function (dictRingRecord) {
            var semiringRecordCons1 = semiringRecordCons(dictRingRecord.SemiringRecord0());
            return function (dictRing) {
                var semiringRecordCons2 = semiringRecordCons1(dictRing.Semiring0());
                return {
                    subRecord: function (v) {
                        return function (ra) {
                            return function (rb) {
                                var tail = subRecord(dictRingRecord)(Type_Proxy["Proxy"].value)(ra)(rb);
                                var key = Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value);
                                var insert = Record_Unsafe.unsafeSet(key);
                                var get = Record_Unsafe.unsafeGet(key);
                                return insert(sub(dictRing)(get(ra))(get(rb)))(tail);
                            };
                        };
                    },
                    SemiringRecord0: function () {
                        return semiringRecordCons2;
                    }
                };
            };
        };
    };
};
var ringRecord = function () {
    return function (dictRingRecord) {
        var semiringRecord1 = semiringRecord(dictRingRecord.SemiringRecord0());
        return {
            sub: subRecord(dictRingRecord)(Type_Proxy["Proxy"].value),
            Semiring0: function () {
                return semiringRecord1;
            }
        };
    };
};
var ringProxy = {
    sub: function (v) {
        return function (v1) {
            return Type_Proxy["Proxy"].value;
        };
    },
    Semiring0: function () {
        return Data_Semiring.semiringProxy;
    }
};
var ringNumber = {
    sub: $foreign.numSub,
    Semiring0: function () {
        return Data_Semiring.semiringNumber;
    }
};
var ringInt = {
    sub: $foreign.intSub,
    Semiring0: function () {
        return Data_Semiring.semiringInt;
    }
};
var ringFn = function (dictRing) {
    var semiringFn = Data_Semiring.semiringFn(dictRing.Semiring0());
    return {
        sub: function (f) {
            return function (g) {
                return function (x) {
                    return sub(dictRing)(f(x))(g(x));
                };
            };
        },
        Semiring0: function () {
            return semiringFn;
        }
    };
};

// | `negate x` can be used as a shorthand for `zero - x`.
var negate = function (dictRing) {
    var Semiring0 = dictRing.Semiring0();
    return function (a) {
        return sub(dictRing)(Data_Semiring.zero(Semiring0))(a);
    };
};
export {
    sub,
    negate,
    subRecord,
    ringInt,
    ringNumber,
    ringUnit,
    ringFn,
    ringProxy,
    ringRecord,
    ringRecordNil,
    ringRecordCons
};
export {
    add,
    mul,
    one,
    zero
} from "../Data.Semiring/index.js";
