// | This module defines a data type and various functions for creating and
// | manipulating JSON values. The README contains additional documentation
// | for this module.
import * as $foreign from "./foreign.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Foreign_Object from "../Foreign.Object/index.js";
var verbJsonType = function (def) {
    return function (f) {
        return function (g) {
            return g(def)(f);
        };
    };
};

// Decoding
var toJsonType = /* #__PURE__ */ (function () {
    return verbJsonType(Data_Maybe.Nothing.value)(Data_Maybe.Just.create);
})();

// | The number zero represented as `Json`
var jsonZero = /* #__PURE__ */ $foreign.fromNumber(0.0);

// | The true boolean value represented as `Json`
var jsonTrue = /* #__PURE__ */ $foreign.fromBoolean(true);

// | Constructs a `Json` object value containing only the provided key and value
var jsonSingletonObject = function (key) {
    return function (val) {
        return $foreign.fromObject(Foreign_Object.singleton(key)(val));
    };
};

// | Constructs a `Json` array value containing only the provided value
var jsonSingletonArray = function (j) {
    return $foreign.fromArray([ j ]);
};

// | The false boolean value represented as `Json`
var jsonFalse = /* #__PURE__ */ $foreign.fromBoolean(false);

// | An empty string represented as `Json`
var jsonEmptyString = /* #__PURE__ */ $foreign.fromString("");

// | An empty object represented as `Json`
var jsonEmptyObject = /* #__PURE__ */ $foreign.fromObject(Foreign_Object.empty);

// | An empty array represented as `Json`
var jsonEmptyArray = /* #__PURE__ */ $foreign.fromArray([  ]);

// Tests
var isJsonType = /* #__PURE__ */ verbJsonType(false)(/* #__PURE__ */ Data_Function["const"](true));
var ordJson = {
    compare: function (a) {
        return function (b) {
            return $foreign["_compare"](Data_Ordering.EQ.value, Data_Ordering.GT.value, Data_Ordering.LT.value, a, b);
        };
    },
    Eq0: function () {
        return eqJson;
    }
};
var eqJson = {
    eq: function (j1) {
        return function (j2) {
            return Data_Eq.eq(Data_Ordering.eqOrdering)(Data_Ord.compare(ordJson)(j1)(j2))(Data_Ordering.EQ.value);
        };
    }
};
var eqJNull = {
    eq: function (v) {
        return function (v1) {
            return true;
        };
    }
};
var ordJNull = {
    compare: function (v) {
        return function (v1) {
            return Data_Ordering.EQ.value;
        };
    },
    Eq0: function () {
        return eqJNull;
    }
};

// | A simpler version of `caseJson` which accepts a callback for when the
// | `Json` argument was a `String`, and a default value for all other cases.
var caseJsonString = function (d) {
    return function (f) {
        return function (j) {
            return $foreign["_caseJson"](Data_Function["const"](d), Data_Function["const"](d), Data_Function["const"](d), f, Data_Function["const"](d), Data_Function["const"](d), j);
        };
    };
};

// | Check if the provided `Json` is a `String`
var isString = /* #__PURE__ */ isJsonType(caseJsonString);

// | Convert `Json` to a `String` value, if the `Json` is a string. To write a
// | `Json` value to a JSON string, see `stringify`.
var toString = /* #__PURE__ */ toJsonType(caseJsonString);

// | A simpler version of `caseJson` which accepts a callback for when the
// | `Json` argument was an `Object`, and a default value for all other cases.
var caseJsonObject = function (d) {
    return function (f) {
        return function (j) {
            return $foreign["_caseJson"](Data_Function["const"](d), Data_Function["const"](d), Data_Function["const"](d), Data_Function["const"](d), Data_Function["const"](d), f, j);
        };
    };
};

// | Check if the provided `Json` is an `Object`
var isObject = /* #__PURE__ */ isJsonType(caseJsonObject);

// | Convert `Json` to an `Object` of `Json` values, if the `Json` is an object.
var toObject = /* #__PURE__ */ toJsonType(caseJsonObject);

// | A simpler version of `caseJson` which accepts a callback for when the
// | `Json` argument was a `Number`, and a default value for all other cases.
var caseJsonNumber = function (d) {
    return function (f) {
        return function (j) {
            return $foreign["_caseJson"](Data_Function["const"](d), Data_Function["const"](d), f, Data_Function["const"](d), Data_Function["const"](d), Data_Function["const"](d), j);
        };
    };
};

// | Check if the provided `Json` is a `Number`
var isNumber = /* #__PURE__ */ isJsonType(caseJsonNumber);

// | Convert `Json` to a `Number` value, if the `Json` is a number.
var toNumber = /* #__PURE__ */ toJsonType(caseJsonNumber);

// | A simpler version of `caseJson` which accepts a callback for when the
// | `Json` argument was null, and a default value for all other cases.
var caseJsonNull = function (d) {
    return function (f) {
        return function (j) {
            return $foreign["_caseJson"](f, Data_Function["const"](d), Data_Function["const"](d), Data_Function["const"](d), Data_Function["const"](d), Data_Function["const"](d), j);
        };
    };
};

// | Check if the provided `Json` is the `null` value
var isNull = /* #__PURE__ */ isJsonType(caseJsonNull);

// | Convert `Json` to the `Unit` value if the `Json` is the null value
var toNull = /* #__PURE__ */ toJsonType(caseJsonNull);

// | A simpler version of `caseJson` which accepts a callback for when the
// | `Json` argument was a `Boolean`, and a default value for all other cases.
var caseJsonBoolean = function (d) {
    return function (f) {
        return function (j) {
            return $foreign["_caseJson"](Data_Function["const"](d), f, Data_Function["const"](d), Data_Function["const"](d), Data_Function["const"](d), Data_Function["const"](d), j);
        };
    };
};

// | Check if the provided `Json` is a `Boolean`
var isBoolean = /* #__PURE__ */ isJsonType(caseJsonBoolean);

// | Convert `Json` to a `Boolean` value, if the `Json` is a boolean.
var toBoolean = /* #__PURE__ */ toJsonType(caseJsonBoolean);

// | A simpler version of `caseJson` which accepts a callback for when the
// | `Json` argument was a `Array Json`, and a default value for all other cases.
var caseJsonArray = function (d) {
    return function (f) {
        return function (j) {
            return $foreign["_caseJson"](Data_Function["const"](d), Data_Function["const"](d), Data_Function["const"](d), Data_Function["const"](d), f, Data_Function["const"](d), j);
        };
    };
};

// | Check if the provided `Json` is an `Array`
var isArray = /* #__PURE__ */ isJsonType(caseJsonArray);

// | Convert `Json` to an `Array` of `Json` values, if the `Json` is an array.
var toArray = /* #__PURE__ */ toJsonType(caseJsonArray);

// | Case analysis for `Json` values. See the README for more information.
var caseJson = function (a) {
    return function (b) {
        return function (c) {
            return function (d) {
                return function (e) {
                    return function (f) {
                        return function (json) {
                            return $foreign["_caseJson"](a, b, c, d, e, f, json);
                        };
                    };
                };
            };
        };
    };
};
export {
    fromBoolean,
    fromNumber,
    fromString,
    fromArray,
    fromObject,
    jsonNull,
    stringify,
    stringifyWithIndent
} from "./foreign.js";
export {
    caseJson,
    caseJsonNull,
    caseJsonBoolean,
    caseJsonNumber,
    caseJsonString,
    caseJsonArray,
    caseJsonObject,
    isNull,
    isBoolean,
    isNumber,
    isString,
    isArray,
    isObject,
    toNull,
    toBoolean,
    toNumber,
    toString,
    toArray,
    toObject,
    jsonTrue,
    jsonFalse,
    jsonZero,
    jsonEmptyString,
    jsonEmptyArray,
    jsonSingletonArray,
    jsonEmptyObject,
    jsonSingletonObject,
    eqJson,
    ordJson
};
