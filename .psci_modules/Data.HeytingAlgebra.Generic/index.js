import * as Data_Generic_Rep from "../Data.Generic.Rep/index.js";
import * as Data_HeytingAlgebra from "../Data.HeytingAlgebra/index.js";
var genericTT$prime = function (dict) {
    return dict["genericTT'"];
};

// | A `Generic` implementation of the `tt` member from the `HeytingAlgebra` type class.
var genericTT = function (dictGeneric) {
    return function (dictGenericHeytingAlgebra) {
        return Data_Generic_Rep.to(dictGeneric)(genericTT$prime(dictGenericHeytingAlgebra));
    };
};
var genericNot$prime = function (dict) {
    return dict["genericNot'"];
};

// | A `Generic` implementation of the `not` member from the `HeytingAlgebra` type class.
var genericNot = function (dictGeneric) {
    var to = Data_Generic_Rep.to(dictGeneric);
    return function (dictGenericHeytingAlgebra) {
        return function (x) {
            return to(genericNot$prime(dictGenericHeytingAlgebra)(Data_Generic_Rep.from(dictGeneric)(x)));
        };
    };
};
var genericImplies$prime = function (dict) {
    return dict["genericImplies'"];
};

// | A `Generic` implementation of the `implies` member from the `HeytingAlgebra` type class.
var genericImplies = function (dictGeneric) {
    var to = Data_Generic_Rep.to(dictGeneric);
    return function (dictGenericHeytingAlgebra) {
        return function (x) {
            return function (y) {
                return to(genericImplies$prime(dictGenericHeytingAlgebra)(Data_Generic_Rep.from(dictGeneric)(x))(Data_Generic_Rep.from(dictGeneric)(y)));
            };
        };
    };
};
var genericHeytingAlgebraNoArguments = /* #__PURE__ */ (function () {
    return {
        "genericFF'": Data_Generic_Rep.NoArguments.value,
        "genericTT'": Data_Generic_Rep.NoArguments.value,
        "genericImplies'": function (v) {
            return function (v1) {
                return Data_Generic_Rep.NoArguments.value;
            };
        },
        "genericConj'": function (v) {
            return function (v1) {
                return Data_Generic_Rep.NoArguments.value;
            };
        },
        "genericDisj'": function (v) {
            return function (v1) {
                return Data_Generic_Rep.NoArguments.value;
            };
        },
        "genericNot'": function (v) {
            return Data_Generic_Rep.NoArguments.value;
        }
    };
})();
var genericHeytingAlgebraArgument = function (dictHeytingAlgebra) {
    return {
        "genericFF'": Data_HeytingAlgebra.ff(dictHeytingAlgebra),
        "genericTT'": Data_HeytingAlgebra.tt(dictHeytingAlgebra),
        "genericImplies'": function (v) {
            return function (v1) {
                return Data_HeytingAlgebra.implies(dictHeytingAlgebra)(v)(v1);
            };
        },
        "genericConj'": function (v) {
            return function (v1) {
                return Data_HeytingAlgebra.conj(dictHeytingAlgebra)(v)(v1);
            };
        },
        "genericDisj'": function (v) {
            return function (v1) {
                return Data_HeytingAlgebra.disj(dictHeytingAlgebra)(v)(v1);
            };
        },
        "genericNot'": function (v) {
            return Data_HeytingAlgebra.not(dictHeytingAlgebra)(v);
        }
    };
};
var genericFF$prime = function (dict) {
    return dict["genericFF'"];
};

// | A `Generic` implementation of the `ff` member from the `HeytingAlgebra` type class.
var genericFF = function (dictGeneric) {
    return function (dictGenericHeytingAlgebra) {
        return Data_Generic_Rep.to(dictGeneric)(genericFF$prime(dictGenericHeytingAlgebra));
    };
};
var genericDisj$prime = function (dict) {
    return dict["genericDisj'"];
};

// | A `Generic` implementation of the `disj` member from the `HeytingAlgebra` type class.
var genericDisj = function (dictGeneric) {
    var to = Data_Generic_Rep.to(dictGeneric);
    return function (dictGenericHeytingAlgebra) {
        return function (x) {
            return function (y) {
                return to(genericDisj$prime(dictGenericHeytingAlgebra)(Data_Generic_Rep.from(dictGeneric)(x))(Data_Generic_Rep.from(dictGeneric)(y)));
            };
        };
    };
};
var genericConj$prime = function (dict) {
    return dict["genericConj'"];
};
var genericHeytingAlgebraConstructor = function (dictGenericHeytingAlgebra) {
    return {
        "genericFF'": genericFF$prime(dictGenericHeytingAlgebra),
        "genericTT'": genericTT$prime(dictGenericHeytingAlgebra),
        "genericImplies'": function (v) {
            return function (v1) {
                return genericImplies$prime(dictGenericHeytingAlgebra)(v)(v1);
            };
        },
        "genericConj'": function (v) {
            return function (v1) {
                return genericConj$prime(dictGenericHeytingAlgebra)(v)(v1);
            };
        },
        "genericDisj'": function (v) {
            return function (v1) {
                return genericDisj$prime(dictGenericHeytingAlgebra)(v)(v1);
            };
        },
        "genericNot'": function (v) {
            return genericNot$prime(dictGenericHeytingAlgebra)(v);
        }
    };
};
var genericHeytingAlgebraProduct = function (dictGenericHeytingAlgebra) {
    return function (dictGenericHeytingAlgebra1) {
        return {
            "genericFF'": new Data_Generic_Rep.Product(genericFF$prime(dictGenericHeytingAlgebra), genericFF$prime(dictGenericHeytingAlgebra1)),
            "genericTT'": new Data_Generic_Rep.Product(genericTT$prime(dictGenericHeytingAlgebra), genericTT$prime(dictGenericHeytingAlgebra1)),
            "genericImplies'": function (v) {
                return function (v1) {
                    return new Data_Generic_Rep.Product(genericImplies$prime(dictGenericHeytingAlgebra)(v.value0)(v1.value0), genericImplies$prime(dictGenericHeytingAlgebra1)(v.value1)(v1.value1));
                };
            },
            "genericConj'": function (v) {
                return function (v1) {
                    return new Data_Generic_Rep.Product(genericConj$prime(dictGenericHeytingAlgebra)(v.value0)(v1.value0), genericConj$prime(dictGenericHeytingAlgebra1)(v.value1)(v1.value1));
                };
            },
            "genericDisj'": function (v) {
                return function (v1) {
                    return new Data_Generic_Rep.Product(genericDisj$prime(dictGenericHeytingAlgebra)(v.value0)(v1.value0), genericDisj$prime(dictGenericHeytingAlgebra1)(v.value1)(v1.value1));
                };
            },
            "genericNot'": function (v) {
                return new Data_Generic_Rep.Product(genericNot$prime(dictGenericHeytingAlgebra)(v.value0), genericNot$prime(dictGenericHeytingAlgebra1)(v.value1));
            }
        };
    };
};

// | A `Generic` implementation of the `conj` member from the `HeytingAlgebra` type class.
var genericConj = function (dictGeneric) {
    var to = Data_Generic_Rep.to(dictGeneric);
    return function (dictGenericHeytingAlgebra) {
        return function (x) {
            return function (y) {
                return to(genericConj$prime(dictGenericHeytingAlgebra)(Data_Generic_Rep.from(dictGeneric)(x))(Data_Generic_Rep.from(dictGeneric)(y)));
            };
        };
    };
};
export {
    genericConj$prime,
    genericDisj$prime,
    genericFF$prime,
    genericImplies$prime,
    genericNot$prime,
    genericTT$prime,
    genericFF,
    genericTT,
    genericImplies,
    genericConj,
    genericDisj,
    genericNot,
    genericHeytingAlgebraNoArguments,
    genericHeytingAlgebraArgument,
    genericHeytingAlgebraProduct,
    genericHeytingAlgebraConstructor
};
