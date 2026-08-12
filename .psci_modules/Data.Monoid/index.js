import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_EuclideanRing from "../Data.EuclideanRing/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Symbol from "../Data.Symbol/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Record_Unsafe from "../Record.Unsafe/index.js";
import * as Type_Proxy from "../Type.Proxy/index.js";
var semigroupRecord = /* #__PURE__ */ Data_Semigroup.semigroupRecord();
var monoidUnit = {
    mempty: Data_Unit.unit,
    Semigroup0: function () {
        return Data_Semigroup.semigroupUnit;
    }
};
var monoidString = {
    mempty: "",
    Semigroup0: function () {
        return Data_Semigroup.semigroupString;
    }
};
var monoidRecordNil = {
    memptyRecord: function (v) {
        return {};
    },
    SemigroupRecord0: function () {
        return Data_Semigroup.semigroupRecordNil;
    }
};
var monoidOrdering = /* #__PURE__ */ (function () {
    return {
        mempty: Data_Ordering.EQ.value,
        Semigroup0: function () {
            return Data_Ordering.semigroupOrdering;
        }
    };
})();
var monoidArray = {
    mempty: [  ],
    Semigroup0: function () {
        return Data_Semigroup.semigroupArray;
    }
};
var memptyRecord = function (dict) {
    return dict.memptyRecord;
};
var monoidRecord = function () {
    return function (dictMonoidRecord) {
        var semigroupRecord1 = semigroupRecord(dictMonoidRecord.SemigroupRecord0());
        return {
            mempty: memptyRecord(dictMonoidRecord)(Type_Proxy["Proxy"].value),
            Semigroup0: function () {
                return semigroupRecord1;
            }
        };
    };
};
var mempty = function (dict) {
    return dict.mempty;
};
var monoidFn = function (dictMonoid) {
    var semigroupFn = Data_Semigroup.semigroupFn(dictMonoid.Semigroup0());
    return {
        mempty: function (v) {
            return mempty(dictMonoid);
        },
        Semigroup0: function () {
            return semigroupFn;
        }
    };
};
var monoidRecordCons = function (dictIsSymbol) {
    var semigroupRecordCons = Data_Semigroup.semigroupRecordCons(dictIsSymbol)();
    return function (dictMonoid) {
        var mempty1 = mempty(dictMonoid);
        var Semigroup0 = dictMonoid.Semigroup0();
        return function () {
            return function (dictMonoidRecord) {
                var semigroupRecordCons1 = semigroupRecordCons(dictMonoidRecord.SemigroupRecord0())(Semigroup0);
                return {
                    memptyRecord: function (v) {
                        var tail = memptyRecord(dictMonoidRecord)(Type_Proxy["Proxy"].value);
                        var key = Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value);
                        var insert = Record_Unsafe.unsafeSet(key);
                        return insert(mempty1)(tail);
                    },
                    SemigroupRecord0: function () {
                        return semigroupRecordCons1;
                    }
                };
            };
        };
    };
};

// | Append a value to itself a certain number of times. For the
// | `Multiplicative` type, and for a non-negative power, this is the same as
// | normal number exponentiation.
// |
// | If the second argument is negative this function will return `mempty`
// | (*unlike* normal number exponentiation). The `Monoid` constraint alone
// | is not enough to write a `power` function with the property that `power x
// | n` cancels with `power x (-n)`, i.e. `power x n <> power x (-n) = mempty`.
// | For that, we would additionally need the ability to invert elements, i.e.
// | a Group.
// |
// | ```purescript
// | power [1,2] 3    == [1,2,1,2,1,2]
// | power [1,2] 1    == [1,2]
// | power [1,2] 0    == []
// | power [1,2] (-3) == []
// | ```
// |
var power = function (dictMonoid) {
    var mempty1 = mempty(dictMonoid);
    var Semigroup0 = dictMonoid.Semigroup0();
    return function (x) {
        var go = function (p) {
            if (p <= 0) {
                return mempty1;
            };
            if (p === 1) {
                return x;
            };
            if (Data_EuclideanRing.mod(Data_EuclideanRing.euclideanRingInt)(p)(2) === 0) {
                var x$prime = go(Data_EuclideanRing.div(Data_EuclideanRing.euclideanRingInt)(p)(2));
                return Data_Semigroup.append(Semigroup0)(x$prime)(x$prime);
            };
            if (Data_Boolean.otherwise) {
                var x$prime = go(Data_EuclideanRing.div(Data_EuclideanRing.euclideanRingInt)(p)(2));
                return Data_Semigroup.append(Semigroup0)(x$prime)(Data_Semigroup.append(Semigroup0)(x$prime)(x));
            };
            throw new Error("Failed pattern match at Data.Monoid (line 88, column 3 - line 88, column 17): " + [ p.constructor.name ]);
        };
        return go;
    };
};

// | Allow or "truncate" a Monoid to its `mempty` value based on a condition.
var guard = function (dictMonoid) {
    var mempty1 = mempty(dictMonoid);
    return function (v) {
        return function (v1) {
            if (v) {
                return v1;
            };
            if (!v) {
                return mempty1;
            };
            throw new Error("Failed pattern match at Data.Monoid (line 96, column 1 - line 96, column 49): " + [ v.constructor.name, v1.constructor.name ]);
        };
    };
};
export {
    mempty,
    power,
    guard,
    memptyRecord,
    monoidUnit,
    monoidOrdering,
    monoidFn,
    monoidString,
    monoidArray,
    monoidRecord,
    monoidRecordNil,
    monoidRecordCons
};
