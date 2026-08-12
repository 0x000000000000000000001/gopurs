import * as Data_Generic_Rep from "../Data.Generic.Rep/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
var genericZero$prime = function (dict) {
    return dict["genericZero'"];
};

// | A `Generic` implementation of the `zero` member from the `Semiring` type class.
var genericZero = function (dictGeneric) {
    return function (dictGenericSemiring) {
        return Data_Generic_Rep.to(dictGeneric)(genericZero$prime(dictGenericSemiring));
    };
};
var genericSemiringNoArguments = /* #__PURE__ */ (function () {
    return {
        "genericAdd'": function (v) {
            return function (v1) {
                return Data_Generic_Rep.NoArguments.value;
            };
        },
        "genericZero'": Data_Generic_Rep.NoArguments.value,
        "genericMul'": function (v) {
            return function (v1) {
                return Data_Generic_Rep.NoArguments.value;
            };
        },
        "genericOne'": Data_Generic_Rep.NoArguments.value
    };
})();
var genericSemiringArgument = function (dictSemiring) {
    return {
        "genericAdd'": function (v) {
            return function (v1) {
                return Data_Semiring.add(dictSemiring)(v)(v1);
            };
        },
        "genericZero'": Data_Semiring.zero(dictSemiring),
        "genericMul'": function (v) {
            return function (v1) {
                return Data_Semiring.mul(dictSemiring)(v)(v1);
            };
        },
        "genericOne'": Data_Semiring.one(dictSemiring)
    };
};
var genericOne$prime = function (dict) {
    return dict["genericOne'"];
};

// | A `Generic` implementation of the `one` member from the `Semiring` type class.
var genericOne = function (dictGeneric) {
    return function (dictGenericSemiring) {
        return Data_Generic_Rep.to(dictGeneric)(genericOne$prime(dictGenericSemiring));
    };
};
var genericMul$prime = function (dict) {
    return dict["genericMul'"];
};

// | A `Generic` implementation of the `mul` member from the `Semiring` type class.
var genericMul = function (dictGeneric) {
    var to = Data_Generic_Rep.to(dictGeneric);
    return function (dictGenericSemiring) {
        return function (x) {
            return function (y) {
                return to(genericMul$prime(dictGenericSemiring)(Data_Generic_Rep.from(dictGeneric)(x))(Data_Generic_Rep.from(dictGeneric)(y)));
            };
        };
    };
};
var genericAdd$prime = function (dict) {
    return dict["genericAdd'"];
};
var genericSemiringConstructor = function (dictGenericSemiring) {
    return {
        "genericAdd'": function (v) {
            return function (v1) {
                return genericAdd$prime(dictGenericSemiring)(v)(v1);
            };
        },
        "genericZero'": genericZero$prime(dictGenericSemiring),
        "genericMul'": function (v) {
            return function (v1) {
                return genericMul$prime(dictGenericSemiring)(v)(v1);
            };
        },
        "genericOne'": genericOne$prime(dictGenericSemiring)
    };
};
var genericSemiringProduct = function (dictGenericSemiring) {
    return function (dictGenericSemiring1) {
        return {
            "genericAdd'": function (v) {
                return function (v1) {
                    return new Data_Generic_Rep.Product(genericAdd$prime(dictGenericSemiring)(v.value0)(v1.value0), genericAdd$prime(dictGenericSemiring1)(v.value1)(v1.value1));
                };
            },
            "genericZero'": new Data_Generic_Rep.Product(genericZero$prime(dictGenericSemiring), genericZero$prime(dictGenericSemiring1)),
            "genericMul'": function (v) {
                return function (v1) {
                    return new Data_Generic_Rep.Product(genericMul$prime(dictGenericSemiring)(v.value0)(v1.value0), genericMul$prime(dictGenericSemiring1)(v.value1)(v1.value1));
                };
            },
            "genericOne'": new Data_Generic_Rep.Product(genericOne$prime(dictGenericSemiring), genericOne$prime(dictGenericSemiring1))
        };
    };
};

// | A `Generic` implementation of the `add` member from the `Semiring` type class.
var genericAdd = function (dictGeneric) {
    var to = Data_Generic_Rep.to(dictGeneric);
    return function (dictGenericSemiring) {
        return function (x) {
            return function (y) {
                return to(genericAdd$prime(dictGenericSemiring)(Data_Generic_Rep.from(dictGeneric)(x))(Data_Generic_Rep.from(dictGeneric)(y)));
            };
        };
    };
};
export {
    genericAdd$prime,
    genericMul$prime,
    genericOne$prime,
    genericZero$prime,
    genericZero,
    genericOne,
    genericAdd,
    genericMul,
    genericSemiringNoArguments,
    genericSemiringArgument,
    genericSemiringProduct,
    genericSemiringConstructor
};
