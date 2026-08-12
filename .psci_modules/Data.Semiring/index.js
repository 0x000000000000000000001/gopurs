import * as $foreign from "./foreign.js";
import * as Data_Symbol from "../Data.Symbol/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Record_Unsafe from "../Record.Unsafe/index.js";
import * as Type_Proxy from "../Type.Proxy/index.js";
var zeroRecord = function (dict) {
    return dict.zeroRecord;
};
var zero = function (dict) {
    return dict.zero;
};
var semiringUnit = {
    add: function (v) {
        return function (v1) {
            return Data_Unit.unit;
        };
    },
    zero: Data_Unit.unit,
    mul: function (v) {
        return function (v1) {
            return Data_Unit.unit;
        };
    },
    one: Data_Unit.unit
};
var semiringRecordNil = {
    addRecord: function (v) {
        return function (v1) {
            return function (v2) {
                return {};
            };
        };
    },
    mulRecord: function (v) {
        return function (v1) {
            return function (v2) {
                return {};
            };
        };
    },
    oneRecord: function (v) {
        return function (v1) {
            return {};
        };
    },
    zeroRecord: function (v) {
        return function (v1) {
            return {};
        };
    }
};
var semiringProxy = /* #__PURE__ */ (function () {
    return {
        add: function (v) {
            return function (v1) {
                return Type_Proxy["Proxy"].value;
            };
        },
        mul: function (v) {
            return function (v1) {
                return Type_Proxy["Proxy"].value;
            };
        },
        one: Type_Proxy["Proxy"].value,
        zero: Type_Proxy["Proxy"].value
    };
})();
var semiringNumber = {
    add: $foreign.numAdd,
    zero: 0.0,
    mul: $foreign.numMul,
    one: 1.0
};
var semiringInt = {
    add: $foreign.intAdd,
    zero: 0,
    mul: $foreign.intMul,
    one: 1
};
var oneRecord = function (dict) {
    return dict.oneRecord;
};
var one = function (dict) {
    return dict.one;
};
var mulRecord = function (dict) {
    return dict.mulRecord;
};
var mul = function (dict) {
    return dict.mul;
};
var addRecord = function (dict) {
    return dict.addRecord;
};
var semiringRecord = function () {
    return function (dictSemiringRecord) {
        return {
            add: addRecord(dictSemiringRecord)(Type_Proxy["Proxy"].value),
            mul: mulRecord(dictSemiringRecord)(Type_Proxy["Proxy"].value),
            one: oneRecord(dictSemiringRecord)(Type_Proxy["Proxy"].value)(Type_Proxy["Proxy"].value),
            zero: zeroRecord(dictSemiringRecord)(Type_Proxy["Proxy"].value)(Type_Proxy["Proxy"].value)
        };
    };
};
var add = function (dict) {
    return dict.add;
};
var semiringFn = function (dictSemiring) {
    return {
        add: function (f) {
            return function (g) {
                return function (x) {
                    return add(dictSemiring)(f(x))(g(x));
                };
            };
        },
        zero: function (v) {
            return zero(dictSemiring);
        },
        mul: function (f) {
            return function (g) {
                return function (x) {
                    return mul(dictSemiring)(f(x))(g(x));
                };
            };
        },
        one: function (v) {
            return one(dictSemiring);
        }
    };
};
var semiringRecordCons = function (dictIsSymbol) {
    return function () {
        return function (dictSemiringRecord) {
            return function (dictSemiring) {
                var one1 = one(dictSemiring);
                var zero1 = zero(dictSemiring);
                return {
                    addRecord: function (v) {
                        return function (ra) {
                            return function (rb) {
                                var tail = addRecord(dictSemiringRecord)(Type_Proxy["Proxy"].value)(ra)(rb);
                                var key = Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value);
                                var insert = Record_Unsafe.unsafeSet(key);
                                var get = Record_Unsafe.unsafeGet(key);
                                return insert(add(dictSemiring)(get(ra))(get(rb)))(tail);
                            };
                        };
                    },
                    mulRecord: function (v) {
                        return function (ra) {
                            return function (rb) {
                                var tail = mulRecord(dictSemiringRecord)(Type_Proxy["Proxy"].value)(ra)(rb);
                                var key = Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value);
                                var insert = Record_Unsafe.unsafeSet(key);
                                var get = Record_Unsafe.unsafeGet(key);
                                return insert(mul(dictSemiring)(get(ra))(get(rb)))(tail);
                            };
                        };
                    },
                    oneRecord: function (v) {
                        return function (v1) {
                            var tail = oneRecord(dictSemiringRecord)(Type_Proxy["Proxy"].value)(Type_Proxy["Proxy"].value);
                            var key = Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value);
                            var insert = Record_Unsafe.unsafeSet(key);
                            return insert(one1)(tail);
                        };
                    },
                    zeroRecord: function (v) {
                        return function (v1) {
                            var tail = zeroRecord(dictSemiringRecord)(Type_Proxy["Proxy"].value)(Type_Proxy["Proxy"].value);
                            var key = Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value);
                            var insert = Record_Unsafe.unsafeSet(key);
                            return insert(zero1)(tail);
                        };
                    }
                };
            };
        };
    };
};
export {
    add,
    zero,
    mul,
    one,
    addRecord,
    mulRecord,
    oneRecord,
    zeroRecord,
    semiringInt,
    semiringNumber,
    semiringFn,
    semiringUnit,
    semiringProxy,
    semiringRecord,
    semiringRecordNil,
    semiringRecordCons
};
