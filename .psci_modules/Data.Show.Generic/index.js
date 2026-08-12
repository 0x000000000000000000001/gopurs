import * as $foreign from "./foreign.js";
import * as Data_Generic_Rep from "../Data.Generic.Rep/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Symbol from "../Data.Symbol/index.js";
import * as Type_Proxy from "../Type.Proxy/index.js";
var genericShowArgsNoArguments = {
    genericShowArgs: function (v) {
        return [  ];
    }
};
var genericShowArgsArgument = function (dictShow) {
    return {
        genericShowArgs: function (v) {
            return [ Data_Show.show(dictShow)(v) ];
        }
    };
};
var genericShowArgs = function (dict) {
    return dict.genericShowArgs;
};
var genericShowArgsProduct = function (dictGenericShowArgs) {
    return function (dictGenericShowArgs1) {
        return {
            genericShowArgs: function (v) {
                return Data_Semigroup.append(Data_Semigroup.semigroupArray)(genericShowArgs(dictGenericShowArgs)(v.value0))(genericShowArgs(dictGenericShowArgs1)(v.value1));
            }
        };
    };
};
var genericShowConstructor = function (dictGenericShowArgs) {
    return function (dictIsSymbol) {
        return {
            "genericShow'": function (v) {
                var ctor = Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value);
                var v1 = genericShowArgs(dictGenericShowArgs)(v);
                if (v1.length === 0) {
                    return ctor;
                };
                return "(" + ($foreign.intercalate(" ")(Data_Semigroup.append(Data_Semigroup.semigroupArray)([ ctor ])(v1)) + ")");
            }
        };
    };
};
var genericShow$prime = function (dict) {
    return dict["genericShow'"];
};
var genericShowNoConstructors = {
    "genericShow'": function (a) {
        return genericShow$prime(genericShowNoConstructors)(a);
    }
};
var genericShowSum = function (dictGenericShow) {
    return function (dictGenericShow1) {
        return {
            "genericShow'": function (v) {
                if (v instanceof Data_Generic_Rep.Inl) {
                    return genericShow$prime(dictGenericShow)(v.value0);
                };
                if (v instanceof Data_Generic_Rep.Inr) {
                    return genericShow$prime(dictGenericShow1)(v.value0);
                };
                throw new Error("Failed pattern match at Data.Show.Generic (line 26, column 1 - line 28, column 40): " + [ v.constructor.name ]);
            }
        };
    };
};

// | A `Generic` implementation of the `show` member from the `Show` type class.
var genericShow = function (dictGeneric) {
    return function (dictGenericShow) {
        return function (x) {
            return genericShow$prime(dictGenericShow)(Data_Generic_Rep.from(dictGeneric)(x));
        };
    };
};
export {
    genericShow$prime,
    genericShow,
    genericShowArgs,
    genericShowNoConstructors,
    genericShowArgsNoArguments,
    genericShowSum,
    genericShowArgsProduct,
    genericShowConstructor,
    genericShowArgsArgument
};
