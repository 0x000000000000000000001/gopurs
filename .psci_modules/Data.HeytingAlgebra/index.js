import * as $foreign from "./foreign.js";
import * as Data_Symbol from "../Data.Symbol/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Record_Unsafe from "../Record.Unsafe/index.js";
import * as Type_Proxy from "../Type.Proxy/index.js";
var ttRecord = function (dict) {
    return dict.ttRecord;
};
var tt = function (dict) {
    return dict.tt;
};
var notRecord = function (dict) {
    return dict.notRecord;
};
var not = function (dict) {
    return dict.not;
};
var impliesRecord = function (dict) {
    return dict.impliesRecord;
};
var implies = function (dict) {
    return dict.implies;
};
var heytingAlgebraUnit = {
    ff: Data_Unit.unit,
    tt: Data_Unit.unit,
    implies: function (v) {
        return function (v1) {
            return Data_Unit.unit;
        };
    },
    conj: function (v) {
        return function (v1) {
            return Data_Unit.unit;
        };
    },
    disj: function (v) {
        return function (v1) {
            return Data_Unit.unit;
        };
    },
    not: function (v) {
        return Data_Unit.unit;
    }
};
var heytingAlgebraRecordNil = {
    conjRecord: function (v) {
        return function (v1) {
            return function (v2) {
                return {};
            };
        };
    },
    disjRecord: function (v) {
        return function (v1) {
            return function (v2) {
                return {};
            };
        };
    },
    ffRecord: function (v) {
        return function (v1) {
            return {};
        };
    },
    impliesRecord: function (v) {
        return function (v1) {
            return function (v2) {
                return {};
            };
        };
    },
    notRecord: function (v) {
        return function (v1) {
            return {};
        };
    },
    ttRecord: function (v) {
        return function (v1) {
            return {};
        };
    }
};
var heytingAlgebraProxy = /* #__PURE__ */ (function () {
    return {
        conj: function (v) {
            return function (v1) {
                return Type_Proxy["Proxy"].value;
            };
        },
        disj: function (v) {
            return function (v1) {
                return Type_Proxy["Proxy"].value;
            };
        },
        implies: function (v) {
            return function (v1) {
                return Type_Proxy["Proxy"].value;
            };
        },
        ff: Type_Proxy["Proxy"].value,
        not: function (v) {
            return Type_Proxy["Proxy"].value;
        },
        tt: Type_Proxy["Proxy"].value
    };
})();
var ffRecord = function (dict) {
    return dict.ffRecord;
};
var ff = function (dict) {
    return dict.ff;
};
var disjRecord = function (dict) {
    return dict.disjRecord;
};
var disj = function (dict) {
    return dict.disj;
};
var heytingAlgebraBoolean = {
    ff: false,
    tt: true,
    implies: function (a) {
        return function (b) {
            return disj(heytingAlgebraBoolean)(not(heytingAlgebraBoolean)(a))(b);
        };
    },
    conj: $foreign.boolConj,
    disj: $foreign.boolDisj,
    not: $foreign.boolNot
};
var conjRecord = function (dict) {
    return dict.conjRecord;
};
var heytingAlgebraRecord = function () {
    return function (dictHeytingAlgebraRecord) {
        return {
            ff: ffRecord(dictHeytingAlgebraRecord)(Type_Proxy["Proxy"].value)(Type_Proxy["Proxy"].value),
            tt: ttRecord(dictHeytingAlgebraRecord)(Type_Proxy["Proxy"].value)(Type_Proxy["Proxy"].value),
            conj: conjRecord(dictHeytingAlgebraRecord)(Type_Proxy["Proxy"].value),
            disj: disjRecord(dictHeytingAlgebraRecord)(Type_Proxy["Proxy"].value),
            implies: impliesRecord(dictHeytingAlgebraRecord)(Type_Proxy["Proxy"].value),
            not: notRecord(dictHeytingAlgebraRecord)(Type_Proxy["Proxy"].value)
        };
    };
};
var conj = function (dict) {
    return dict.conj;
};
var heytingAlgebraFunction = function (dictHeytingAlgebra) {
    return {
        ff: function (v) {
            return ff(dictHeytingAlgebra);
        },
        tt: function (v) {
            return tt(dictHeytingAlgebra);
        },
        implies: function (f) {
            return function (g) {
                return function (a) {
                    return implies(dictHeytingAlgebra)(f(a))(g(a));
                };
            };
        },
        conj: function (f) {
            return function (g) {
                return function (a) {
                    return conj(dictHeytingAlgebra)(f(a))(g(a));
                };
            };
        },
        disj: function (f) {
            return function (g) {
                return function (a) {
                    return disj(dictHeytingAlgebra)(f(a))(g(a));
                };
            };
        },
        not: function (f) {
            return function (a) {
                return not(dictHeytingAlgebra)(f(a));
            };
        }
    };
};
var heytingAlgebraRecordCons = function (dictIsSymbol) {
    return function () {
        return function (dictHeytingAlgebraRecord) {
            return function (dictHeytingAlgebra) {
                var ff1 = ff(dictHeytingAlgebra);
                var tt1 = tt(dictHeytingAlgebra);
                return {
                    conjRecord: function (v) {
                        return function (ra) {
                            return function (rb) {
                                var tail = conjRecord(dictHeytingAlgebraRecord)(Type_Proxy["Proxy"].value)(ra)(rb);
                                var key = Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value);
                                var insert = Record_Unsafe.unsafeSet(key);
                                var get = Record_Unsafe.unsafeGet(key);
                                return insert(conj(dictHeytingAlgebra)(get(ra))(get(rb)))(tail);
                            };
                        };
                    },
                    disjRecord: function (v) {
                        return function (ra) {
                            return function (rb) {
                                var tail = disjRecord(dictHeytingAlgebraRecord)(Type_Proxy["Proxy"].value)(ra)(rb);
                                var key = Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value);
                                var insert = Record_Unsafe.unsafeSet(key);
                                var get = Record_Unsafe.unsafeGet(key);
                                return insert(disj(dictHeytingAlgebra)(get(ra))(get(rb)))(tail);
                            };
                        };
                    },
                    impliesRecord: function (v) {
                        return function (ra) {
                            return function (rb) {
                                var tail = impliesRecord(dictHeytingAlgebraRecord)(Type_Proxy["Proxy"].value)(ra)(rb);
                                var key = Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value);
                                var insert = Record_Unsafe.unsafeSet(key);
                                var get = Record_Unsafe.unsafeGet(key);
                                return insert(implies(dictHeytingAlgebra)(get(ra))(get(rb)))(tail);
                            };
                        };
                    },
                    ffRecord: function (v) {
                        return function (row) {
                            var tail = ffRecord(dictHeytingAlgebraRecord)(Type_Proxy["Proxy"].value)(row);
                            var key = Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value);
                            var insert = Record_Unsafe.unsafeSet(key);
                            return insert(ff1)(tail);
                        };
                    },
                    notRecord: function (v) {
                        return function (row) {
                            var tail = notRecord(dictHeytingAlgebraRecord)(Type_Proxy["Proxy"].value)(row);
                            var key = Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value);
                            var insert = Record_Unsafe.unsafeSet(key);
                            var get = Record_Unsafe.unsafeGet(key);
                            return insert(not(dictHeytingAlgebra)(get(row)))(tail);
                        };
                    },
                    ttRecord: function (v) {
                        return function (row) {
                            var tail = ttRecord(dictHeytingAlgebraRecord)(Type_Proxy["Proxy"].value)(row);
                            var key = Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value);
                            var insert = Record_Unsafe.unsafeSet(key);
                            return insert(tt1)(tail);
                        };
                    }
                };
            };
        };
    };
};
export {
    tt,
    ff,
    implies,
    conj,
    disj,
    not,
    ffRecord,
    ttRecord,
    impliesRecord,
    conjRecord,
    disjRecord,
    notRecord,
    heytingAlgebraBoolean,
    heytingAlgebraUnit,
    heytingAlgebraFunction,
    heytingAlgebraProxy,
    heytingAlgebraRecord,
    heytingAlgebraRecordNil,
    heytingAlgebraRecordCons
};
